// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package shared

import (
	"reflect"
	"time"

	"github.com/orbcorp/orb-go/internal/apijson"
	"github.com/orbcorp/orb-go/internal/param"
	"github.com/tidwall/gjson"
)

type Address struct {
	City       string      `json:"city,required,nullable"`
	Country    string      `json:"country,required,nullable"`
	Line1      string      `json:"line1,required,nullable"`
	Line2      string      `json:"line2,required,nullable"`
	PostalCode string      `json:"postal_code,required,nullable"`
	State      string      `json:"state,required,nullable"`
	JSON       addressJSON `json:"-"`
}

// addressJSON contains the JSON metadata for the struct [Address]
type addressJSON struct {
	City        apijson.Field
	Country     apijson.Field
	Line1       apijson.Field
	Line2       apijson.Field
	PostalCode  apijson.Field
	State       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *Address) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r addressJSON) RawJSON() string {
	return r.raw
}

type AdjustmentInterval struct {
	ID         string                       `json:"id,required"`
	Adjustment AdjustmentIntervalAdjustment `json:"adjustment,required"`
	// The price interval IDs that this adjustment applies to.
	AppliesToPriceIntervalIDs []string `json:"applies_to_price_interval_ids,required"`
	// The end date of the adjustment interval.
	EndDate time.Time `json:"end_date,required,nullable" format:"date-time"`
	// The start date of the adjustment interval.
	StartDate time.Time              `json:"start_date,required" format:"date-time"`
	JSON      adjustmentIntervalJSON `json:"-"`
}

// adjustmentIntervalJSON contains the JSON metadata for the struct
// [AdjustmentInterval]
type adjustmentIntervalJSON struct {
	ID                        apijson.Field
	Adjustment                apijson.Field
	AppliesToPriceIntervalIDs apijson.Field
	EndDate                   apijson.Field
	StartDate                 apijson.Field
	raw                       string
	ExtraFields               map[string]apijson.Field
}

func (r *AdjustmentInterval) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r adjustmentIntervalJSON) RawJSON() string {
	return r.raw
}

type AdjustmentIntervalAdjustment struct {
	ID             string                                     `json:"id,required"`
	AdjustmentType AdjustmentIntervalAdjustmentAdjustmentType `json:"adjustment_type,required"`
	// This field can have the runtime type of [[]string].
	AppliesToPriceIDs interface{} `json:"applies_to_price_ids,required"`
	// This field can have the runtime type of [[]TransformPriceFilter].
	Filters interface{} `json:"filters,required"`
	// True for adjustments that apply to an entire invocice, false for adjustments
	// that apply to only one price.
	IsInvoiceLevel bool `json:"is_invoice_level,required"`
	// The plan phase in which this adjustment is active.
	PlanPhaseOrder int64 `json:"plan_phase_order,required,nullable"`
	// The reason for the adjustment.
	Reason string `json:"reason,required,nullable"`
	// The adjustment id this adjustment replaces. This adjustment will take the place
	// of the replaced adjustment in plan version migrations.
	ReplacesAdjustmentID string `json:"replaces_adjustment_id,required,nullable"`
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
	UsageDiscount float64                          `json:"usage_discount"`
	JSON          adjustmentIntervalAdjustmentJSON `json:"-"`
	union         AdjustmentIntervalAdjustmentUnion
}

// adjustmentIntervalAdjustmentJSON contains the JSON metadata for the struct
// [AdjustmentIntervalAdjustment]
type adjustmentIntervalAdjustmentJSON struct {
	ID                   apijson.Field
	AdjustmentType       apijson.Field
	AppliesToPriceIDs    apijson.Field
	Filters              apijson.Field
	IsInvoiceLevel       apijson.Field
	PlanPhaseOrder       apijson.Field
	Reason               apijson.Field
	ReplacesAdjustmentID apijson.Field
	AmountDiscount       apijson.Field
	ItemID               apijson.Field
	MaximumAmount        apijson.Field
	MinimumAmount        apijson.Field
	PercentageDiscount   apijson.Field
	UsageDiscount        apijson.Field
	raw                  string
	ExtraFields          map[string]apijson.Field
}

func (r adjustmentIntervalAdjustmentJSON) RawJSON() string {
	return r.raw
}

func (r *AdjustmentIntervalAdjustment) UnmarshalJSON(data []byte) (err error) {
	*r = AdjustmentIntervalAdjustment{}
	err = apijson.UnmarshalRoot(data, &r.union)
	if err != nil {
		return err
	}
	return apijson.Port(r.union, &r)
}

// AsUnion returns a [AdjustmentIntervalAdjustmentUnion] interface which you can
// cast to the specific types for more type safety.
//
// Possible runtime types of the union are [PlanPhaseUsageDiscountAdjustment],
// [PlanPhaseAmountDiscountAdjustment], [PlanPhasePercentageDiscountAdjustment],
// [PlanPhaseMinimumAdjustment], [PlanPhaseMaximumAdjustment].
func (r AdjustmentIntervalAdjustment) AsUnion() AdjustmentIntervalAdjustmentUnion {
	return r.union
}

// Union satisfied by [PlanPhaseUsageDiscountAdjustment],
// [PlanPhaseAmountDiscountAdjustment], [PlanPhasePercentageDiscountAdjustment],
// [PlanPhaseMinimumAdjustment] or [PlanPhaseMaximumAdjustment].
type AdjustmentIntervalAdjustmentUnion interface {
	ImplementsAdjustmentIntervalAdjustment()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*AdjustmentIntervalAdjustmentUnion)(nil)).Elem(),
		"adjustment_type",
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(PlanPhaseUsageDiscountAdjustment{}),
			DiscriminatorValue: "usage_discount",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(PlanPhaseAmountDiscountAdjustment{}),
			DiscriminatorValue: "amount_discount",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(PlanPhasePercentageDiscountAdjustment{}),
			DiscriminatorValue: "percentage_discount",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(PlanPhaseMinimumAdjustment{}),
			DiscriminatorValue: "minimum",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(PlanPhaseMaximumAdjustment{}),
			DiscriminatorValue: "maximum",
		},
	)
}

type AdjustmentIntervalAdjustmentAdjustmentType string

const (
	AdjustmentIntervalAdjustmentAdjustmentTypeUsageDiscount      AdjustmentIntervalAdjustmentAdjustmentType = "usage_discount"
	AdjustmentIntervalAdjustmentAdjustmentTypeAmountDiscount     AdjustmentIntervalAdjustmentAdjustmentType = "amount_discount"
	AdjustmentIntervalAdjustmentAdjustmentTypePercentageDiscount AdjustmentIntervalAdjustmentAdjustmentType = "percentage_discount"
	AdjustmentIntervalAdjustmentAdjustmentTypeMinimum            AdjustmentIntervalAdjustmentAdjustmentType = "minimum"
	AdjustmentIntervalAdjustmentAdjustmentTypeMaximum            AdjustmentIntervalAdjustmentAdjustmentType = "maximum"
)

func (r AdjustmentIntervalAdjustmentAdjustmentType) IsKnown() bool {
	switch r {
	case AdjustmentIntervalAdjustmentAdjustmentTypeUsageDiscount, AdjustmentIntervalAdjustmentAdjustmentTypeAmountDiscount, AdjustmentIntervalAdjustmentAdjustmentTypePercentageDiscount, AdjustmentIntervalAdjustmentAdjustmentTypeMinimum, AdjustmentIntervalAdjustmentAdjustmentTypeMaximum:
		return true
	}
	return false
}

type AggregatedCost struct {
	PerPriceCosts []PerPriceCost `json:"per_price_costs,required"`
	// Total costs for the timeframe, excluding any minimums and discounts.
	Subtotal       string    `json:"subtotal,required"`
	TimeframeEnd   time.Time `json:"timeframe_end,required" format:"date-time"`
	TimeframeStart time.Time `json:"timeframe_start,required" format:"date-time"`
	// Total costs for the timeframe, including any minimums and discounts.
	Total string             `json:"total,required"`
	JSON  aggregatedCostJSON `json:"-"`
}

// aggregatedCostJSON contains the JSON metadata for the struct [AggregatedCost]
type aggregatedCostJSON struct {
	PerPriceCosts  apijson.Field
	Subtotal       apijson.Field
	TimeframeEnd   apijson.Field
	TimeframeStart apijson.Field
	Total          apijson.Field
	raw            string
	ExtraFields    map[string]apijson.Field
}

func (r *AggregatedCost) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r aggregatedCostJSON) RawJSON() string {
	return r.raw
}

type Allocation struct {
	AllowsRollover   bool             `json:"allows_rollover,required"`
	Currency         string           `json:"currency,required"`
	CustomExpiration CustomExpiration `json:"custom_expiration,required,nullable"`
	JSON             allocationJSON   `json:"-"`
}

// allocationJSON contains the JSON metadata for the struct [Allocation]
type allocationJSON struct {
	AllowsRollover   apijson.Field
	Currency         apijson.Field
	CustomExpiration apijson.Field
	raw              string
	ExtraFields      map[string]apijson.Field
}

func (r *Allocation) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r allocationJSON) RawJSON() string {
	return r.raw
}

type AmountDiscount struct {
	// Only available if discount_type is `amount`.
	AmountDiscount string                     `json:"amount_discount,required"`
	DiscountType   AmountDiscountDiscountType `json:"discount_type,required"`
	// List of price_ids that this discount applies to. For plan/plan phase discounts,
	// this can be a subset of prices.
	AppliesToPriceIDs []string `json:"applies_to_price_ids,nullable"`
	// The filters that determine which prices to apply this discount to.
	Filters []TransformPriceFilter `json:"filters,nullable"`
	Reason  string                 `json:"reason,nullable"`
	JSON    amountDiscountJSON     `json:"-"`
}

// amountDiscountJSON contains the JSON metadata for the struct [AmountDiscount]
type amountDiscountJSON struct {
	AmountDiscount    apijson.Field
	DiscountType      apijson.Field
	AppliesToPriceIDs apijson.Field
	Filters           apijson.Field
	Reason            apijson.Field
	raw               string
	ExtraFields       map[string]apijson.Field
}

func (r *AmountDiscount) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r amountDiscountJSON) RawJSON() string {
	return r.raw
}

func (r AmountDiscount) ImplementsDiscount() {}

func (r AmountDiscount) ImplementsInvoiceLevelDiscount() {}

func (r AmountDiscount) ImplementsCouponDiscount() {}

type AmountDiscountDiscountType string

const (
	AmountDiscountDiscountTypeAmount AmountDiscountDiscountType = "amount"
)

func (r AmountDiscountDiscountType) IsKnown() bool {
	switch r {
	case AmountDiscountDiscountTypeAmount:
		return true
	}
	return false
}

type AmountDiscountParam struct {
	// Only available if discount_type is `amount`.
	AmountDiscount param.Field[string]                     `json:"amount_discount,required"`
	DiscountType   param.Field[AmountDiscountDiscountType] `json:"discount_type,required"`
	// List of price_ids that this discount applies to. For plan/plan phase discounts,
	// this can be a subset of prices.
	AppliesToPriceIDs param.Field[[]string] `json:"applies_to_price_ids"`
	// The filters that determine which prices to apply this discount to.
	Filters param.Field[[]TransformPriceFilterParam] `json:"filters"`
	Reason  param.Field[string]                      `json:"reason"`
}

func (r AmountDiscountParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r AmountDiscountParam) ImplementsDiscountUnionParam() {}

type AmountDiscountInterval struct {
	// Only available if discount_type is `amount`.
	AmountDiscount string `json:"amount_discount,required"`
	// The price interval ids that this discount interval applies to.
	AppliesToPriceIntervalIDs []string                           `json:"applies_to_price_interval_ids,required"`
	DiscountType              AmountDiscountIntervalDiscountType `json:"discount_type,required"`
	// The end date of the discount interval.
	EndDate time.Time `json:"end_date,required,nullable" format:"date-time"`
	// The filters that determine which prices this discount interval applies to.
	Filters []TransformPriceFilter `json:"filters,required"`
	// The start date of the discount interval.
	StartDate time.Time                  `json:"start_date,required" format:"date-time"`
	JSON      amountDiscountIntervalJSON `json:"-"`
}

// amountDiscountIntervalJSON contains the JSON metadata for the struct
// [AmountDiscountInterval]
type amountDiscountIntervalJSON struct {
	AmountDiscount            apijson.Field
	AppliesToPriceIntervalIDs apijson.Field
	DiscountType              apijson.Field
	EndDate                   apijson.Field
	Filters                   apijson.Field
	StartDate                 apijson.Field
	raw                       string
	ExtraFields               map[string]apijson.Field
}

func (r *AmountDiscountInterval) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r amountDiscountIntervalJSON) RawJSON() string {
	return r.raw
}

func (r AmountDiscountInterval) ImplementsSubscriptionDiscountInterval() {}

func (r AmountDiscountInterval) ImplementsMutatedSubscriptionDiscountInterval() {}

type AmountDiscountIntervalDiscountType string

const (
	AmountDiscountIntervalDiscountTypeAmount AmountDiscountIntervalDiscountType = "amount"
)

func (r AmountDiscountIntervalDiscountType) IsKnown() bool {
	switch r {
	case AmountDiscountIntervalDiscountTypeAmount:
		return true
	}
	return false
}

type BillableMetricTiny struct {
	ID   string                 `json:"id,required"`
	JSON billableMetricTinyJSON `json:"-"`
}

// billableMetricTinyJSON contains the JSON metadata for the struct
// [BillableMetricTiny]
type billableMetricTinyJSON struct {
	ID          apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *BillableMetricTiny) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r billableMetricTinyJSON) RawJSON() string {
	return r.raw
}

type BillingCycleAnchorConfiguration struct {
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
	Year int64                               `json:"year,nullable"`
	JSON billingCycleAnchorConfigurationJSON `json:"-"`
}

// billingCycleAnchorConfigurationJSON contains the JSON metadata for the struct
// [BillingCycleAnchorConfiguration]
type billingCycleAnchorConfigurationJSON struct {
	Day         apijson.Field
	Month       apijson.Field
	Year        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *BillingCycleAnchorConfiguration) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r billingCycleAnchorConfigurationJSON) RawJSON() string {
	return r.raw
}

type BillingCycleAnchorConfigurationParam struct {
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

func (r BillingCycleAnchorConfigurationParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type BillingCycleConfiguration struct {
	Duration     int64                                 `json:"duration,required"`
	DurationUnit BillingCycleConfigurationDurationUnit `json:"duration_unit,required"`
	JSON         billingCycleConfigurationJSON         `json:"-"`
}

// billingCycleConfigurationJSON contains the JSON metadata for the struct
// [BillingCycleConfiguration]
type billingCycleConfigurationJSON struct {
	Duration     apijson.Field
	DurationUnit apijson.Field
	raw          string
	ExtraFields  map[string]apijson.Field
}

func (r *BillingCycleConfiguration) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r billingCycleConfigurationJSON) RawJSON() string {
	return r.raw
}

type BillingCycleConfigurationDurationUnit string

const (
	BillingCycleConfigurationDurationUnitDay   BillingCycleConfigurationDurationUnit = "day"
	BillingCycleConfigurationDurationUnitMonth BillingCycleConfigurationDurationUnit = "month"
)

func (r BillingCycleConfigurationDurationUnit) IsKnown() bool {
	switch r {
	case BillingCycleConfigurationDurationUnitDay, BillingCycleConfigurationDurationUnitMonth:
		return true
	}
	return false
}

type BillingCycleRelativeDate string

const (
	BillingCycleRelativeDateStartOfTerm BillingCycleRelativeDate = "start_of_term"
	BillingCycleRelativeDateEndOfTerm   BillingCycleRelativeDate = "end_of_term"
)

func (r BillingCycleRelativeDate) IsKnown() bool {
	switch r {
	case BillingCycleRelativeDateStartOfTerm, BillingCycleRelativeDateEndOfTerm:
		return true
	}
	return false
}

func (r BillingCycleRelativeDate) ImplementsSubscriptionPriceIntervalsParamsAddStartDateUnion() {}

func (r BillingCycleRelativeDate) ImplementsSubscriptionPriceIntervalsParamsAddEndDateUnion() {}

func (r BillingCycleRelativeDate) ImplementsSubscriptionPriceIntervalsParamsAddAdjustmentsStartDateUnion() {
}

func (r BillingCycleRelativeDate) ImplementsSubscriptionPriceIntervalsParamsAddAdjustmentsEndDateUnion() {
}

func (r BillingCycleRelativeDate) ImplementsSubscriptionPriceIntervalsParamsEditEndDateUnion() {}

func (r BillingCycleRelativeDate) ImplementsSubscriptionPriceIntervalsParamsEditStartDateUnion() {}

func (r BillingCycleRelativeDate) ImplementsSubscriptionPriceIntervalsParamsEditAdjustmentsEndDateUnion() {
}

func (r BillingCycleRelativeDate) ImplementsSubscriptionPriceIntervalsParamsEditAdjustmentsStartDateUnion() {
}

type BPSConfig struct {
	// Basis point take rate per event
	BPS float64 `json:"bps,required"`
	// Optional currency amount maximum to cap spend per event
	PerUnitMaximum string        `json:"per_unit_maximum,nullable"`
	JSON           bpsConfigJSON `json:"-"`
}

// bpsConfigJSON contains the JSON metadata for the struct [BPSConfig]
type bpsConfigJSON struct {
	BPS            apijson.Field
	PerUnitMaximum apijson.Field
	raw            string
	ExtraFields    map[string]apijson.Field
}

func (r *BPSConfig) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r bpsConfigJSON) RawJSON() string {
	return r.raw
}

type BPSConfigParam struct {
	// Basis point take rate per event
	BPS param.Field[float64] `json:"bps,required"`
	// Optional currency amount maximum to cap spend per event
	PerUnitMaximum param.Field[string] `json:"per_unit_maximum"`
}

func (r BPSConfigParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type BPSTier struct {
	// Per-event basis point rate
	BPS float64 `json:"bps,required"`
	// Exclusive tier starting value
	MinimumAmount string `json:"minimum_amount,required"`
	// Inclusive tier ending value
	MaximumAmount string `json:"maximum_amount,nullable"`
	// Per unit maximum to charge
	PerUnitMaximum string      `json:"per_unit_maximum,nullable"`
	JSON           bpsTierJSON `json:"-"`
}

// bpsTierJSON contains the JSON metadata for the struct [BPSTier]
type bpsTierJSON struct {
	BPS            apijson.Field
	MinimumAmount  apijson.Field
	MaximumAmount  apijson.Field
	PerUnitMaximum apijson.Field
	raw            string
	ExtraFields    map[string]apijson.Field
}

func (r *BPSTier) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r bpsTierJSON) RawJSON() string {
	return r.raw
}

type BPSTierParam struct {
	// Per-event basis point rate
	BPS param.Field[float64] `json:"bps,required"`
	// Exclusive tier starting value
	MinimumAmount param.Field[string] `json:"minimum_amount,required"`
	// Inclusive tier ending value
	MaximumAmount param.Field[string] `json:"maximum_amount"`
	// Per unit maximum to charge
	PerUnitMaximum param.Field[string] `json:"per_unit_maximum"`
}

func (r BPSTierParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type BulkBPSConfig struct {
	// Tiers for a bulk BPS pricing model where all usage is aggregated to a single
	// tier based on total volume
	Tiers []BulkBPSTier     `json:"tiers,required"`
	JSON  bulkBPSConfigJSON `json:"-"`
}

// bulkBPSConfigJSON contains the JSON metadata for the struct [BulkBPSConfig]
type bulkBPSConfigJSON struct {
	Tiers       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *BulkBPSConfig) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r bulkBPSConfigJSON) RawJSON() string {
	return r.raw
}

type BulkBPSConfigParam struct {
	// Tiers for a bulk BPS pricing model where all usage is aggregated to a single
	// tier based on total volume
	Tiers param.Field[[]BulkBPSTierParam] `json:"tiers,required"`
}

func (r BulkBPSConfigParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type BulkBPSTier struct {
	// Basis points to rate on
	BPS float64 `json:"bps,required"`
	// Upper bound for tier
	MaximumAmount string `json:"maximum_amount,nullable"`
	// The maximum amount to charge for any one event
	PerUnitMaximum string          `json:"per_unit_maximum,nullable"`
	JSON           bulkBPSTierJSON `json:"-"`
}

// bulkBPSTierJSON contains the JSON metadata for the struct [BulkBPSTier]
type bulkBPSTierJSON struct {
	BPS            apijson.Field
	MaximumAmount  apijson.Field
	PerUnitMaximum apijson.Field
	raw            string
	ExtraFields    map[string]apijson.Field
}

func (r *BulkBPSTier) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r bulkBPSTierJSON) RawJSON() string {
	return r.raw
}

type BulkBPSTierParam struct {
	// Basis points to rate on
	BPS param.Field[float64] `json:"bps,required"`
	// Upper bound for tier
	MaximumAmount param.Field[string] `json:"maximum_amount"`
	// The maximum amount to charge for any one event
	PerUnitMaximum param.Field[string] `json:"per_unit_maximum"`
}

func (r BulkBPSTierParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type BulkConfig struct {
	// Bulk tiers for rating based on total usage volume
	Tiers []BulkTier     `json:"tiers,required"`
	JSON  bulkConfigJSON `json:"-"`
}

// bulkConfigJSON contains the JSON metadata for the struct [BulkConfig]
type bulkConfigJSON struct {
	Tiers       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *BulkConfig) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r bulkConfigJSON) RawJSON() string {
	return r.raw
}

type BulkConfigParam struct {
	// Bulk tiers for rating based on total usage volume
	Tiers param.Field[[]BulkTierParam] `json:"tiers,required"`
}

func (r BulkConfigParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type BulkTier struct {
	// Amount per unit
	UnitAmount string `json:"unit_amount,required"`
	// Upper bound for this tier
	MaximumUnits float64      `json:"maximum_units,nullable"`
	JSON         bulkTierJSON `json:"-"`
}

// bulkTierJSON contains the JSON metadata for the struct [BulkTier]
type bulkTierJSON struct {
	UnitAmount   apijson.Field
	MaximumUnits apijson.Field
	raw          string
	ExtraFields  map[string]apijson.Field
}

func (r *BulkTier) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r bulkTierJSON) RawJSON() string {
	return r.raw
}

type BulkTierParam struct {
	// Amount per unit
	UnitAmount param.Field[string] `json:"unit_amount,required"`
	// Upper bound for this tier
	MaximumUnits param.Field[float64] `json:"maximum_units"`
}

func (r BulkTierParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type ChangedSubscriptionResources struct {
	// The credit notes that were created as part of this operation.
	CreatedCreditNotes []CreditNote `json:"created_credit_notes,required"`
	// The invoices that were created as part of this operation.
	CreatedInvoices []Invoice `json:"created_invoices,required"`
	// The credit notes that were voided as part of this operation.
	VoidedCreditNotes []CreditNote `json:"voided_credit_notes,required"`
	// The invoices that were voided as part of this operation.
	VoidedInvoices []Invoice                        `json:"voided_invoices,required"`
	JSON           changedSubscriptionResourcesJSON `json:"-"`
}

// changedSubscriptionResourcesJSON contains the JSON metadata for the struct
// [ChangedSubscriptionResources]
type changedSubscriptionResourcesJSON struct {
	CreatedCreditNotes apijson.Field
	CreatedInvoices    apijson.Field
	VoidedCreditNotes  apijson.Field
	VoidedInvoices     apijson.Field
	raw                string
	ExtraFields        map[string]apijson.Field
}

func (r *ChangedSubscriptionResources) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r changedSubscriptionResourcesJSON) RawJSON() string {
	return r.raw
}

type ConversionRateTier struct {
	// Exclusive tier starting value
	FirstUnit float64 `json:"first_unit,required"`
	// Amount per unit of overage
	UnitAmount string `json:"unit_amount,required"`
	// Inclusive tier ending value. If null, this is treated as the last tier
	LastUnit float64                `json:"last_unit,nullable"`
	JSON     conversionRateTierJSON `json:"-"`
}

// conversionRateTierJSON contains the JSON metadata for the struct
// [ConversionRateTier]
type conversionRateTierJSON struct {
	FirstUnit   apijson.Field
	UnitAmount  apijson.Field
	LastUnit    apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ConversionRateTier) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r conversionRateTierJSON) RawJSON() string {
	return r.raw
}

type ConversionRateTierParam struct {
	// Exclusive tier starting value
	FirstUnit param.Field[float64] `json:"first_unit,required"`
	// Amount per unit of overage
	UnitAmount param.Field[string] `json:"unit_amount,required"`
	// Inclusive tier ending value. If null, this is treated as the last tier
	LastUnit param.Field[float64] `json:"last_unit"`
}

func (r ConversionRateTierParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type ConversionRateTieredConfig struct {
	// Tiers for rating based on total usage quantities into the specified tier
	Tiers []ConversionRateTier           `json:"tiers,required"`
	JSON  conversionRateTieredConfigJSON `json:"-"`
}

// conversionRateTieredConfigJSON contains the JSON metadata for the struct
// [ConversionRateTieredConfig]
type conversionRateTieredConfigJSON struct {
	Tiers       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ConversionRateTieredConfig) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r conversionRateTieredConfigJSON) RawJSON() string {
	return r.raw
}

type ConversionRateTieredConfigParam struct {
	// Tiers for rating based on total usage quantities into the specified tier
	Tiers param.Field[[]ConversionRateTierParam] `json:"tiers,required"`
}

func (r ConversionRateTieredConfigParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type ConversionRateUnitConfig struct {
	// Amount per unit of overage
	UnitAmount string                       `json:"unit_amount,required"`
	JSON       conversionRateUnitConfigJSON `json:"-"`
}

// conversionRateUnitConfigJSON contains the JSON metadata for the struct
// [ConversionRateUnitConfig]
type conversionRateUnitConfigJSON struct {
	UnitAmount  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ConversionRateUnitConfig) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r conversionRateUnitConfigJSON) RawJSON() string {
	return r.raw
}

type ConversionRateUnitConfigParam struct {
	// Amount per unit of overage
	UnitAmount param.Field[string] `json:"unit_amount,required"`
}

func (r ConversionRateUnitConfigParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type CouponRedemption struct {
	CouponID  string               `json:"coupon_id,required"`
	EndDate   time.Time            `json:"end_date,required,nullable" format:"date-time"`
	StartDate time.Time            `json:"start_date,required" format:"date-time"`
	JSON      couponRedemptionJSON `json:"-"`
}

// couponRedemptionJSON contains the JSON metadata for the struct
// [CouponRedemption]
type couponRedemptionJSON struct {
	CouponID    apijson.Field
	EndDate     apijson.Field
	StartDate   apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *CouponRedemption) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r couponRedemptionJSON) RawJSON() string {
	return r.raw
}

// The [Credit Note](/invoicing/credit-notes) resource represents a credit that has
// been applied to a particular invoice.
type CreditNote struct {
	// The Orb id of this credit note.
	ID string `json:"id,required"`
	// The creation time of the resource in Orb.
	CreatedAt time.Time `json:"created_at,required" format:"date-time"`
	// The unique identifier for credit notes.
	CreditNoteNumber string `json:"credit_note_number,required"`
	// A URL to a PDF of the credit note.
	CreditNotePdf string           `json:"credit_note_pdf,required,nullable"`
	Customer      CustomerMinified `json:"customer,required"`
	// The id of the invoice resource that this credit note is applied to.
	InvoiceID string `json:"invoice_id,required"`
	// All of the line items associated with this credit note.
	LineItems []CreditNoteLineItem `json:"line_items,required"`
	// The maximum amount applied on the original invoice
	MaximumAmountAdjustment CreditNoteMaximumAmountAdjustment `json:"maximum_amount_adjustment,required,nullable"`
	// An optional memo supplied on the credit note.
	Memo string `json:"memo,required,nullable"`
	// Any credited amount from the applied minimum on the invoice.
	MinimumAmountRefunded string           `json:"minimum_amount_refunded,required,nullable"`
	Reason                CreditNoteReason `json:"reason,required,nullable"`
	// The total prior to any creditable invoice-level discounts or minimums.
	Subtotal string `json:"subtotal,required"`
	// The total including creditable invoice-level discounts or minimums, and tax.
	Total string         `json:"total,required"`
	Type  CreditNoteType `json:"type,required"`
	// The time at which the credit note was voided in Orb, if applicable.
	VoidedAt time.Time `json:"voided_at,required,nullable" format:"date-time"`
	// Any discounts applied on the original invoice.
	Discounts []CreditNoteDiscount `json:"discounts"`
	JSON      creditNoteJSON       `json:"-"`
}

// creditNoteJSON contains the JSON metadata for the struct [CreditNote]
type creditNoteJSON struct {
	ID                      apijson.Field
	CreatedAt               apijson.Field
	CreditNoteNumber        apijson.Field
	CreditNotePdf           apijson.Field
	Customer                apijson.Field
	InvoiceID               apijson.Field
	LineItems               apijson.Field
	MaximumAmountAdjustment apijson.Field
	Memo                    apijson.Field
	MinimumAmountRefunded   apijson.Field
	Reason                  apijson.Field
	Subtotal                apijson.Field
	Total                   apijson.Field
	Type                    apijson.Field
	VoidedAt                apijson.Field
	Discounts               apijson.Field
	raw                     string
	ExtraFields             map[string]apijson.Field
}

func (r *CreditNote) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r creditNoteJSON) RawJSON() string {
	return r.raw
}

type CreditNoteLineItem struct {
	// The Orb id of this resource.
	ID string `json:"id,required"`
	// The amount of the line item, including any line item minimums and discounts.
	Amount string `json:"amount,required"`
	// The id of the item associated with this line item.
	ItemID string `json:"item_id,required"`
	// The name of the corresponding invoice line item.
	Name string `json:"name,required"`
	// An optional quantity credited.
	Quantity float64 `json:"quantity,required,nullable"`
	// The amount of the line item, excluding any line item minimums and discounts.
	Subtotal string `json:"subtotal,required"`
	// Any tax amounts applied onto the line item.
	TaxAmounts []TaxAmount `json:"tax_amounts,required"`
	// Any line item discounts from the invoice's line item.
	Discounts []CreditNoteLineItemsDiscount `json:"discounts"`
	// The end time of the service period for this credit note line item.
	EndTimeExclusive time.Time `json:"end_time_exclusive,nullable" format:"date-time"`
	// The start time of the service period for this credit note line item.
	StartTimeInclusive time.Time              `json:"start_time_inclusive,nullable" format:"date-time"`
	JSON               creditNoteLineItemJSON `json:"-"`
}

// creditNoteLineItemJSON contains the JSON metadata for the struct
// [CreditNoteLineItem]
type creditNoteLineItemJSON struct {
	ID                 apijson.Field
	Amount             apijson.Field
	ItemID             apijson.Field
	Name               apijson.Field
	Quantity           apijson.Field
	Subtotal           apijson.Field
	TaxAmounts         apijson.Field
	Discounts          apijson.Field
	EndTimeExclusive   apijson.Field
	StartTimeInclusive apijson.Field
	raw                string
	ExtraFields        map[string]apijson.Field
}

func (r *CreditNoteLineItem) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r creditNoteLineItemJSON) RawJSON() string {
	return r.raw
}

type CreditNoteLineItemsDiscount struct {
	ID                 string                                   `json:"id,required"`
	AmountApplied      string                                   `json:"amount_applied,required"`
	AppliesToPriceIDs  []string                                 `json:"applies_to_price_ids,required"`
	DiscountType       CreditNoteLineItemsDiscountsDiscountType `json:"discount_type,required"`
	PercentageDiscount float64                                  `json:"percentage_discount,required"`
	AmountDiscount     string                                   `json:"amount_discount,nullable"`
	Reason             string                                   `json:"reason,nullable"`
	JSON               creditNoteLineItemsDiscountJSON          `json:"-"`
}

// creditNoteLineItemsDiscountJSON contains the JSON metadata for the struct
// [CreditNoteLineItemsDiscount]
type creditNoteLineItemsDiscountJSON struct {
	ID                 apijson.Field
	AmountApplied      apijson.Field
	AppliesToPriceIDs  apijson.Field
	DiscountType       apijson.Field
	PercentageDiscount apijson.Field
	AmountDiscount     apijson.Field
	Reason             apijson.Field
	raw                string
	ExtraFields        map[string]apijson.Field
}

func (r *CreditNoteLineItemsDiscount) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r creditNoteLineItemsDiscountJSON) RawJSON() string {
	return r.raw
}

type CreditNoteLineItemsDiscountsDiscountType string

const (
	CreditNoteLineItemsDiscountsDiscountTypePercentage CreditNoteLineItemsDiscountsDiscountType = "percentage"
	CreditNoteLineItemsDiscountsDiscountTypeAmount     CreditNoteLineItemsDiscountsDiscountType = "amount"
)

func (r CreditNoteLineItemsDiscountsDiscountType) IsKnown() bool {
	switch r {
	case CreditNoteLineItemsDiscountsDiscountTypePercentage, CreditNoteLineItemsDiscountsDiscountTypeAmount:
		return true
	}
	return false
}

// The maximum amount applied on the original invoice
type CreditNoteMaximumAmountAdjustment struct {
	AmountApplied      string                                            `json:"amount_applied,required"`
	DiscountType       CreditNoteMaximumAmountAdjustmentDiscountType     `json:"discount_type,required"`
	PercentageDiscount float64                                           `json:"percentage_discount,required"`
	AppliesToPrices    []CreditNoteMaximumAmountAdjustmentAppliesToPrice `json:"applies_to_prices,nullable"`
	Reason             string                                            `json:"reason,nullable"`
	JSON               creditNoteMaximumAmountAdjustmentJSON             `json:"-"`
}

// creditNoteMaximumAmountAdjustmentJSON contains the JSON metadata for the struct
// [CreditNoteMaximumAmountAdjustment]
type creditNoteMaximumAmountAdjustmentJSON struct {
	AmountApplied      apijson.Field
	DiscountType       apijson.Field
	PercentageDiscount apijson.Field
	AppliesToPrices    apijson.Field
	Reason             apijson.Field
	raw                string
	ExtraFields        map[string]apijson.Field
}

func (r *CreditNoteMaximumAmountAdjustment) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r creditNoteMaximumAmountAdjustmentJSON) RawJSON() string {
	return r.raw
}

type CreditNoteMaximumAmountAdjustmentDiscountType string

const (
	CreditNoteMaximumAmountAdjustmentDiscountTypePercentage CreditNoteMaximumAmountAdjustmentDiscountType = "percentage"
)

func (r CreditNoteMaximumAmountAdjustmentDiscountType) IsKnown() bool {
	switch r {
	case CreditNoteMaximumAmountAdjustmentDiscountTypePercentage:
		return true
	}
	return false
}

type CreditNoteMaximumAmountAdjustmentAppliesToPrice struct {
	ID   string                                              `json:"id,required"`
	Name string                                              `json:"name,required"`
	JSON creditNoteMaximumAmountAdjustmentAppliesToPriceJSON `json:"-"`
}

// creditNoteMaximumAmountAdjustmentAppliesToPriceJSON contains the JSON metadata
// for the struct [CreditNoteMaximumAmountAdjustmentAppliesToPrice]
type creditNoteMaximumAmountAdjustmentAppliesToPriceJSON struct {
	ID          apijson.Field
	Name        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *CreditNoteMaximumAmountAdjustmentAppliesToPrice) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r creditNoteMaximumAmountAdjustmentAppliesToPriceJSON) RawJSON() string {
	return r.raw
}

type CreditNoteReason string

const (
	CreditNoteReasonDuplicate             CreditNoteReason = "Duplicate"
	CreditNoteReasonFraudulent            CreditNoteReason = "Fraudulent"
	CreditNoteReasonOrderChange           CreditNoteReason = "Order change"
	CreditNoteReasonProductUnsatisfactory CreditNoteReason = "Product unsatisfactory"
)

func (r CreditNoteReason) IsKnown() bool {
	switch r {
	case CreditNoteReasonDuplicate, CreditNoteReasonFraudulent, CreditNoteReasonOrderChange, CreditNoteReasonProductUnsatisfactory:
		return true
	}
	return false
}

type CreditNoteType string

const (
	CreditNoteTypeRefund     CreditNoteType = "refund"
	CreditNoteTypeAdjustment CreditNoteType = "adjustment"
)

func (r CreditNoteType) IsKnown() bool {
	switch r {
	case CreditNoteTypeRefund, CreditNoteTypeAdjustment:
		return true
	}
	return false
}

type CreditNoteDiscount struct {
	AmountApplied      string                              `json:"amount_applied,required"`
	DiscountType       CreditNoteDiscountsDiscountType     `json:"discount_type,required"`
	PercentageDiscount float64                             `json:"percentage_discount,required"`
	AppliesToPrices    []CreditNoteDiscountsAppliesToPrice `json:"applies_to_prices,nullable"`
	Reason             string                              `json:"reason,nullable"`
	JSON               creditNoteDiscountJSON              `json:"-"`
}

// creditNoteDiscountJSON contains the JSON metadata for the struct
// [CreditNoteDiscount]
type creditNoteDiscountJSON struct {
	AmountApplied      apijson.Field
	DiscountType       apijson.Field
	PercentageDiscount apijson.Field
	AppliesToPrices    apijson.Field
	Reason             apijson.Field
	raw                string
	ExtraFields        map[string]apijson.Field
}

func (r *CreditNoteDiscount) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r creditNoteDiscountJSON) RawJSON() string {
	return r.raw
}

type CreditNoteDiscountsDiscountType string

const (
	CreditNoteDiscountsDiscountTypePercentage CreditNoteDiscountsDiscountType = "percentage"
)

func (r CreditNoteDiscountsDiscountType) IsKnown() bool {
	switch r {
	case CreditNoteDiscountsDiscountTypePercentage:
		return true
	}
	return false
}

type CreditNoteDiscountsAppliesToPrice struct {
	ID   string                                `json:"id,required"`
	Name string                                `json:"name,required"`
	JSON creditNoteDiscountsAppliesToPriceJSON `json:"-"`
}

// creditNoteDiscountsAppliesToPriceJSON contains the JSON metadata for the struct
// [CreditNoteDiscountsAppliesToPrice]
type creditNoteDiscountsAppliesToPriceJSON struct {
	ID          apijson.Field
	Name        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *CreditNoteDiscountsAppliesToPrice) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r creditNoteDiscountsAppliesToPriceJSON) RawJSON() string {
	return r.raw
}

type CreditNoteTiny struct {
	// The id of the Credit note
	ID   string             `json:"id,required"`
	JSON creditNoteTinyJSON `json:"-"`
}

// creditNoteTinyJSON contains the JSON metadata for the struct [CreditNoteTiny]
type creditNoteTinyJSON struct {
	ID          apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *CreditNoteTiny) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r creditNoteTinyJSON) RawJSON() string {
	return r.raw
}

type CustomExpiration struct {
	Duration     int64                        `json:"duration,required"`
	DurationUnit CustomExpirationDurationUnit `json:"duration_unit,required"`
	JSON         customExpirationJSON         `json:"-"`
}

// customExpirationJSON contains the JSON metadata for the struct
// [CustomExpiration]
type customExpirationJSON struct {
	Duration     apijson.Field
	DurationUnit apijson.Field
	raw          string
	ExtraFields  map[string]apijson.Field
}

func (r *CustomExpiration) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r customExpirationJSON) RawJSON() string {
	return r.raw
}

type CustomExpirationDurationUnit string

const (
	CustomExpirationDurationUnitDay   CustomExpirationDurationUnit = "day"
	CustomExpirationDurationUnitMonth CustomExpirationDurationUnit = "month"
)

func (r CustomExpirationDurationUnit) IsKnown() bool {
	switch r {
	case CustomExpirationDurationUnitDay, CustomExpirationDurationUnitMonth:
		return true
	}
	return false
}

type CustomExpirationParam struct {
	Duration     param.Field[int64]                        `json:"duration,required"`
	DurationUnit param.Field[CustomExpirationDurationUnit] `json:"duration_unit,required"`
}

func (r CustomExpirationParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type CustomerMinified struct {
	ID                 string               `json:"id,required"`
	ExternalCustomerID string               `json:"external_customer_id,required,nullable"`
	JSON               customerMinifiedJSON `json:"-"`
}

// customerMinifiedJSON contains the JSON metadata for the struct
// [CustomerMinified]
type customerMinifiedJSON struct {
	ID                 apijson.Field
	ExternalCustomerID apijson.Field
	raw                string
	ExtraFields        map[string]apijson.Field
}

func (r *CustomerMinified) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r customerMinifiedJSON) RawJSON() string {
	return r.raw
}

// Tax IDs are commonly required to be displayed on customer invoices, which are
// added to the headers of invoices.
//
// ### Supported Tax ID Countries and Types
//
// | Country                | Type         | Description                                                                                             |
// | ---------------------- | ------------ | ------------------------------------------------------------------------------------------------------- |
// | Albania                | `al_tin`     | Albania Tax Identification Number                                                                       |
// | Andorra                | `ad_nrt`     | Andorran NRT Number                                                                                     |
// | Angola                 | `ao_tin`     | Angola Tax Identification Number                                                                        |
// | Argentina              | `ar_cuit`    | Argentinian Tax ID Number                                                                               |
// | Armenia                | `am_tin`     | Armenia Tax Identification Number                                                                       |
// | Aruba                  | `aw_tin`     | Aruba Tax Identification Number                                                                         |
// | Australia              | `au_abn`     | Australian Business Number (AU ABN)                                                                     |
// | Australia              | `au_arn`     | Australian Taxation Office Reference Number                                                             |
// | Austria                | `eu_vat`     | European VAT Number                                                                                     |
// | Azerbaijan             | `az_tin`     | Azerbaijan Tax Identification Number                                                                    |
// | Bahamas                | `bs_tin`     | Bahamas Tax Identification Number                                                                       |
// | Bahrain                | `bh_vat`     | Bahraini VAT Number                                                                                     |
// | Bangladesh             | `bd_bin`     | Bangladesh Business Identification Number                                                               |
// | Barbados               | `bb_tin`     | Barbados Tax Identification Number                                                                      |
// | Belarus                | `by_tin`     | Belarus TIN Number                                                                                      |
// | Belgium                | `eu_vat`     | European VAT Number                                                                                     |
// | Benin                  | `bj_ifu`     | Benin Tax Identification Number (Identifiant Fiscal Unique)                                             |
// | Bolivia                | `bo_tin`     | Bolivian Tax ID                                                                                         |
// | Bosnia and Herzegovina | `ba_tin`     | Bosnia and Herzegovina Tax Identification Number                                                        |
// | Brazil                 | `br_cnpj`    | Brazilian CNPJ Number                                                                                   |
// | Brazil                 | `br_cpf`     | Brazilian CPF Number                                                                                    |
// | Bulgaria               | `bg_uic`     | Bulgaria Unified Identification Code                                                                    |
// | Bulgaria               | `eu_vat`     | European VAT Number                                                                                     |
// | Burkina Faso           | `bf_ifu`     | Burkina Faso Tax Identification Number (Numéro d'Identifiant Fiscal Unique)                             |
// | Cambodia               | `kh_tin`     | Cambodia Tax Identification Number                                                                      |
// | Cameroon               | `cm_niu`     | Cameroon Tax Identification Number (Numéro d'Identifiant fiscal Unique)                                 |
// | Canada                 | `ca_bn`      | Canadian BN                                                                                             |
// | Canada                 | `ca_gst_hst` | Canadian GST/HST Number                                                                                 |
// | Canada                 | `ca_pst_bc`  | Canadian PST Number (British Columbia)                                                                  |
// | Canada                 | `ca_pst_mb`  | Canadian PST Number (Manitoba)                                                                          |
// | Canada                 | `ca_pst_sk`  | Canadian PST Number (Saskatchewan)                                                                      |
// | Canada                 | `ca_qst`     | Canadian QST Number (Québec)                                                                            |
// | Cape Verde             | `cv_nif`     | Cape Verde Tax Identification Number (Número de Identificação Fiscal)                                   |
// | Chile                  | `cl_tin`     | Chilean TIN                                                                                             |
// | China                  | `cn_tin`     | Chinese Tax ID                                                                                          |
// | Colombia               | `co_nit`     | Colombian NIT Number                                                                                    |
// | Congo-Kinshasa         | `cd_nif`     | Congo (DR) Tax Identification Number (Número de Identificação Fiscal)                                   |
// | Costa Rica             | `cr_tin`     | Costa Rican Tax ID                                                                                      |
// | Croatia                | `eu_vat`     | European VAT Number                                                                                     |
// | Croatia                | `hr_oib`     | Croatian Personal Identification Number (OIB)                                                           |
// | Cyprus                 | `eu_vat`     | European VAT Number                                                                                     |
// | Czech Republic         | `eu_vat`     | European VAT Number                                                                                     |
// | Denmark                | `eu_vat`     | European VAT Number                                                                                     |
// | Dominican Republic     | `do_rcn`     | Dominican RCN Number                                                                                    |
// | Ecuador                | `ec_ruc`     | Ecuadorian RUC Number                                                                                   |
// | Egypt                  | `eg_tin`     | Egyptian Tax Identification Number                                                                      |
// | El Salvador            | `sv_nit`     | El Salvadorian NIT Number                                                                               |
// | Estonia                | `eu_vat`     | European VAT Number                                                                                     |
// | Ethiopia               | `et_tin`     | Ethiopia Tax Identification Number                                                                      |
// | European Union         | `eu_oss_vat` | European One Stop Shop VAT Number for non-Union scheme                                                  |
// | Finland                | `eu_vat`     | European VAT Number                                                                                     |
// | France                 | `eu_vat`     | European VAT Number                                                                                     |
// | Georgia                | `ge_vat`     | Georgian VAT                                                                                            |
// | Germany                | `de_stn`     | German Tax Number (Steuernummer)                                                                        |
// | Germany                | `eu_vat`     | European VAT Number                                                                                     |
// | Greece                 | `eu_vat`     | European VAT Number                                                                                     |
// | Guinea                 | `gn_nif`     | Guinea Tax Identification Number (Número de Identificação Fiscal)                                       |
// | Hong Kong              | `hk_br`      | Hong Kong BR Number                                                                                     |
// | Hungary                | `eu_vat`     | European VAT Number                                                                                     |
// | Hungary                | `hu_tin`     | Hungary Tax Number (adószám)                                                                            |
// | Iceland                | `is_vat`     | Icelandic VAT                                                                                           |
// | India                  | `in_gst`     | Indian GST Number                                                                                       |
// | Indonesia              | `id_npwp`    | Indonesian NPWP Number                                                                                  |
// | Ireland                | `eu_vat`     | European VAT Number                                                                                     |
// | Israel                 | `il_vat`     | Israel VAT                                                                                              |
// | Italy                  | `eu_vat`     | European VAT Number                                                                                     |
// | Japan                  | `jp_cn`      | Japanese Corporate Number (_Hōjin Bangō_)                                                               |
// | Japan                  | `jp_rn`      | Japanese Registered Foreign Businesses' Registration Number (_Tōroku Kokugai Jigyōsha no Tōroku Bangō_) |
// | Japan                  | `jp_trn`     | Japanese Tax Registration Number (_Tōroku Bangō_)                                                       |
// | Kazakhstan             | `kz_bin`     | Kazakhstani Business Identification Number                                                              |
// | Kenya                  | `ke_pin`     | Kenya Revenue Authority Personal Identification Number                                                  |
// | Kyrgyzstan             | `kg_tin`     | Kyrgyzstan Tax Identification Number                                                                    |
// | Laos                   | `la_tin`     | Laos Tax Identification Number                                                                          |
// | Latvia                 | `eu_vat`     | European VAT Number                                                                                     |
// | Liechtenstein          | `li_uid`     | Liechtensteinian UID Number                                                                             |
// | Liechtenstein          | `li_vat`     | Liechtenstein VAT Number                                                                                |
// | Lithuania              | `eu_vat`     | European VAT Number                                                                                     |
// | Luxembourg             | `eu_vat`     | European VAT Number                                                                                     |
// | Malaysia               | `my_frp`     | Malaysian FRP Number                                                                                    |
// | Malaysia               | `my_itn`     | Malaysian ITN                                                                                           |
// | Malaysia               | `my_sst`     | Malaysian SST Number                                                                                    |
// | Malta                  | `eu_vat`     | European VAT Number                                                                                     |
// | Mauritania             | `mr_nif`     | Mauritania Tax Identification Number (Número de Identificação Fiscal)                                   |
// | Mexico                 | `mx_rfc`     | Mexican RFC Number                                                                                      |
// | Moldova                | `md_vat`     | Moldova VAT Number                                                                                      |
// | Montenegro             | `me_pib`     | Montenegro PIB Number                                                                                   |
// | Morocco                | `ma_vat`     | Morocco VAT Number                                                                                      |
// | Nepal                  | `np_pan`     | Nepal PAN Number                                                                                        |
// | Netherlands            | `eu_vat`     | European VAT Number                                                                                     |
// | New Zealand            | `nz_gst`     | New Zealand GST Number                                                                                  |
// | Nigeria                | `ng_tin`     | Nigerian Tax Identification Number                                                                      |
// | North Macedonia        | `mk_vat`     | North Macedonia VAT Number                                                                              |
// | Northern Ireland       | `eu_vat`     | Northern Ireland VAT Number                                                                             |
// | Norway                 | `no_vat`     | Norwegian VAT Number                                                                                    |
// | Norway                 | `no_voec`    | Norwegian VAT on e-commerce Number                                                                      |
// | Oman                   | `om_vat`     | Omani VAT Number                                                                                        |
// | Peru                   | `pe_ruc`     | Peruvian RUC Number                                                                                     |
// | Philippines            | `ph_tin`     | Philippines Tax Identification Number                                                                   |
// | Poland                 | `eu_vat`     | European VAT Number                                                                                     |
// | Portugal               | `eu_vat`     | European VAT Number                                                                                     |
// | Romania                | `eu_vat`     | European VAT Number                                                                                     |
// | Romania                | `ro_tin`     | Romanian Tax ID Number                                                                                  |
// | Russia                 | `ru_inn`     | Russian INN                                                                                             |
// | Russia                 | `ru_kpp`     | Russian KPP                                                                                             |
// | Saudi Arabia           | `sa_vat`     | Saudi Arabia VAT                                                                                        |
// | Senegal                | `sn_ninea`   | Senegal NINEA Number                                                                                    |
// | Serbia                 | `rs_pib`     | Serbian PIB Number                                                                                      |
// | Singapore              | `sg_gst`     | Singaporean GST                                                                                         |
// | Singapore              | `sg_uen`     | Singaporean UEN                                                                                         |
// | Slovakia               | `eu_vat`     | European VAT Number                                                                                     |
// | Slovenia               | `eu_vat`     | European VAT Number                                                                                     |
// | Slovenia               | `si_tin`     | Slovenia Tax Number (davčna številka)                                                                   |
// | South Africa           | `za_vat`     | South African VAT Number                                                                                |
// | South Korea            | `kr_brn`     | Korean BRN                                                                                              |
// | Spain                  | `es_cif`     | Spanish NIF Number (previously Spanish CIF Number)                                                      |
// | Spain                  | `eu_vat`     | European VAT Number                                                                                     |
// | Suriname               | `sr_fin`     | Suriname FIN Number                                                                                     |
// | Sweden                 | `eu_vat`     | European VAT Number                                                                                     |
// | Switzerland            | `ch_uid`     | Switzerland UID Number                                                                                  |
// | Switzerland            | `ch_vat`     | Switzerland VAT Number                                                                                  |
// | Taiwan                 | `tw_vat`     | Taiwanese VAT                                                                                           |
// | Tajikistan             | `tj_tin`     | Tajikistan Tax Identification Number                                                                    |
// | Tanzania               | `tz_vat`     | Tanzania VAT Number                                                                                     |
// | Thailand               | `th_vat`     | Thai VAT                                                                                                |
// | Turkey                 | `tr_tin`     | Turkish Tax Identification Number                                                                       |
// | Uganda                 | `ug_tin`     | Uganda Tax Identification Number                                                                        |
// | Ukraine                | `ua_vat`     | Ukrainian VAT                                                                                           |
// | United Arab Emirates   | `ae_trn`     | United Arab Emirates TRN                                                                                |
// | United Kingdom         | `gb_vat`     | United Kingdom VAT Number                                                                               |
// | United States          | `us_ein`     | United States EIN                                                                                       |
// | Uruguay                | `uy_ruc`     | Uruguayan RUC Number                                                                                    |
// | Uzbekistan             | `uz_tin`     | Uzbekistan TIN Number                                                                                   |
// | Uzbekistan             | `uz_vat`     | Uzbekistan VAT Number                                                                                   |
// | Venezuela              | `ve_rif`     | Venezuelan RIF Number                                                                                   |
// | Vietnam                | `vn_tin`     | Vietnamese Tax ID Number                                                                                |
// | Zambia                 | `zm_tin`     | Zambia Tax Identification Number                                                                        |
// | Zimbabwe               | `zw_tin`     | Zimbabwe Tax Identification Number                                                                      |
type CustomerTaxID struct {
	Country CustomerTaxIDCountry `json:"country,required"`
	Type    CustomerTaxIDType    `json:"type,required"`
	Value   string               `json:"value,required"`
	JSON    customerTaxIDJSON    `json:"-"`
}

// customerTaxIDJSON contains the JSON metadata for the struct [CustomerTaxID]
type customerTaxIDJSON struct {
	Country     apijson.Field
	Type        apijson.Field
	Value       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *CustomerTaxID) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r customerTaxIDJSON) RawJSON() string {
	return r.raw
}

type CustomerTaxIDCountry string

const (
	CustomerTaxIDCountryAd CustomerTaxIDCountry = "AD"
	CustomerTaxIDCountryAe CustomerTaxIDCountry = "AE"
	CustomerTaxIDCountryAl CustomerTaxIDCountry = "AL"
	CustomerTaxIDCountryAm CustomerTaxIDCountry = "AM"
	CustomerTaxIDCountryAo CustomerTaxIDCountry = "AO"
	CustomerTaxIDCountryAr CustomerTaxIDCountry = "AR"
	CustomerTaxIDCountryAt CustomerTaxIDCountry = "AT"
	CustomerTaxIDCountryAu CustomerTaxIDCountry = "AU"
	CustomerTaxIDCountryAw CustomerTaxIDCountry = "AW"
	CustomerTaxIDCountryAz CustomerTaxIDCountry = "AZ"
	CustomerTaxIDCountryBa CustomerTaxIDCountry = "BA"
	CustomerTaxIDCountryBb CustomerTaxIDCountry = "BB"
	CustomerTaxIDCountryBd CustomerTaxIDCountry = "BD"
	CustomerTaxIDCountryBe CustomerTaxIDCountry = "BE"
	CustomerTaxIDCountryBf CustomerTaxIDCountry = "BF"
	CustomerTaxIDCountryBg CustomerTaxIDCountry = "BG"
	CustomerTaxIDCountryBh CustomerTaxIDCountry = "BH"
	CustomerTaxIDCountryBj CustomerTaxIDCountry = "BJ"
	CustomerTaxIDCountryBo CustomerTaxIDCountry = "BO"
	CustomerTaxIDCountryBr CustomerTaxIDCountry = "BR"
	CustomerTaxIDCountryBs CustomerTaxIDCountry = "BS"
	CustomerTaxIDCountryBy CustomerTaxIDCountry = "BY"
	CustomerTaxIDCountryCa CustomerTaxIDCountry = "CA"
	CustomerTaxIDCountryCd CustomerTaxIDCountry = "CD"
	CustomerTaxIDCountryCh CustomerTaxIDCountry = "CH"
	CustomerTaxIDCountryCl CustomerTaxIDCountry = "CL"
	CustomerTaxIDCountryCm CustomerTaxIDCountry = "CM"
	CustomerTaxIDCountryCn CustomerTaxIDCountry = "CN"
	CustomerTaxIDCountryCo CustomerTaxIDCountry = "CO"
	CustomerTaxIDCountryCr CustomerTaxIDCountry = "CR"
	CustomerTaxIDCountryCv CustomerTaxIDCountry = "CV"
	CustomerTaxIDCountryDe CustomerTaxIDCountry = "DE"
	CustomerTaxIDCountryCy CustomerTaxIDCountry = "CY"
	CustomerTaxIDCountryCz CustomerTaxIDCountry = "CZ"
	CustomerTaxIDCountryDk CustomerTaxIDCountry = "DK"
	CustomerTaxIDCountryDo CustomerTaxIDCountry = "DO"
	CustomerTaxIDCountryEc CustomerTaxIDCountry = "EC"
	CustomerTaxIDCountryEe CustomerTaxIDCountry = "EE"
	CustomerTaxIDCountryEg CustomerTaxIDCountry = "EG"
	CustomerTaxIDCountryEs CustomerTaxIDCountry = "ES"
	CustomerTaxIDCountryEt CustomerTaxIDCountry = "ET"
	CustomerTaxIDCountryEu CustomerTaxIDCountry = "EU"
	CustomerTaxIDCountryFi CustomerTaxIDCountry = "FI"
	CustomerTaxIDCountryFr CustomerTaxIDCountry = "FR"
	CustomerTaxIDCountryGB CustomerTaxIDCountry = "GB"
	CustomerTaxIDCountryGe CustomerTaxIDCountry = "GE"
	CustomerTaxIDCountryGn CustomerTaxIDCountry = "GN"
	CustomerTaxIDCountryGr CustomerTaxIDCountry = "GR"
	CustomerTaxIDCountryHk CustomerTaxIDCountry = "HK"
	CustomerTaxIDCountryHr CustomerTaxIDCountry = "HR"
	CustomerTaxIDCountryHu CustomerTaxIDCountry = "HU"
	CustomerTaxIDCountryID CustomerTaxIDCountry = "ID"
	CustomerTaxIDCountryIe CustomerTaxIDCountry = "IE"
	CustomerTaxIDCountryIl CustomerTaxIDCountry = "IL"
	CustomerTaxIDCountryIn CustomerTaxIDCountry = "IN"
	CustomerTaxIDCountryIs CustomerTaxIDCountry = "IS"
	CustomerTaxIDCountryIt CustomerTaxIDCountry = "IT"
	CustomerTaxIDCountryJp CustomerTaxIDCountry = "JP"
	CustomerTaxIDCountryKe CustomerTaxIDCountry = "KE"
	CustomerTaxIDCountryKg CustomerTaxIDCountry = "KG"
	CustomerTaxIDCountryKh CustomerTaxIDCountry = "KH"
	CustomerTaxIDCountryKr CustomerTaxIDCountry = "KR"
	CustomerTaxIDCountryKz CustomerTaxIDCountry = "KZ"
	CustomerTaxIDCountryLa CustomerTaxIDCountry = "LA"
	CustomerTaxIDCountryLi CustomerTaxIDCountry = "LI"
	CustomerTaxIDCountryLt CustomerTaxIDCountry = "LT"
	CustomerTaxIDCountryLu CustomerTaxIDCountry = "LU"
	CustomerTaxIDCountryLv CustomerTaxIDCountry = "LV"
	CustomerTaxIDCountryMa CustomerTaxIDCountry = "MA"
	CustomerTaxIDCountryMd CustomerTaxIDCountry = "MD"
	CustomerTaxIDCountryMe CustomerTaxIDCountry = "ME"
	CustomerTaxIDCountryMk CustomerTaxIDCountry = "MK"
	CustomerTaxIDCountryMr CustomerTaxIDCountry = "MR"
	CustomerTaxIDCountryMt CustomerTaxIDCountry = "MT"
	CustomerTaxIDCountryMx CustomerTaxIDCountry = "MX"
	CustomerTaxIDCountryMy CustomerTaxIDCountry = "MY"
	CustomerTaxIDCountryNg CustomerTaxIDCountry = "NG"
	CustomerTaxIDCountryNl CustomerTaxIDCountry = "NL"
	CustomerTaxIDCountryNo CustomerTaxIDCountry = "NO"
	CustomerTaxIDCountryNp CustomerTaxIDCountry = "NP"
	CustomerTaxIDCountryNz CustomerTaxIDCountry = "NZ"
	CustomerTaxIDCountryOm CustomerTaxIDCountry = "OM"
	CustomerTaxIDCountryPe CustomerTaxIDCountry = "PE"
	CustomerTaxIDCountryPh CustomerTaxIDCountry = "PH"
	CustomerTaxIDCountryPl CustomerTaxIDCountry = "PL"
	CustomerTaxIDCountryPt CustomerTaxIDCountry = "PT"
	CustomerTaxIDCountryRo CustomerTaxIDCountry = "RO"
	CustomerTaxIDCountryRs CustomerTaxIDCountry = "RS"
	CustomerTaxIDCountryRu CustomerTaxIDCountry = "RU"
	CustomerTaxIDCountrySa CustomerTaxIDCountry = "SA"
	CustomerTaxIDCountrySe CustomerTaxIDCountry = "SE"
	CustomerTaxIDCountrySg CustomerTaxIDCountry = "SG"
	CustomerTaxIDCountrySi CustomerTaxIDCountry = "SI"
	CustomerTaxIDCountrySk CustomerTaxIDCountry = "SK"
	CustomerTaxIDCountrySn CustomerTaxIDCountry = "SN"
	CustomerTaxIDCountrySr CustomerTaxIDCountry = "SR"
	CustomerTaxIDCountrySv CustomerTaxIDCountry = "SV"
	CustomerTaxIDCountryTh CustomerTaxIDCountry = "TH"
	CustomerTaxIDCountryTj CustomerTaxIDCountry = "TJ"
	CustomerTaxIDCountryTr CustomerTaxIDCountry = "TR"
	CustomerTaxIDCountryTw CustomerTaxIDCountry = "TW"
	CustomerTaxIDCountryTz CustomerTaxIDCountry = "TZ"
	CustomerTaxIDCountryUa CustomerTaxIDCountry = "UA"
	CustomerTaxIDCountryUg CustomerTaxIDCountry = "UG"
	CustomerTaxIDCountryUs CustomerTaxIDCountry = "US"
	CustomerTaxIDCountryUy CustomerTaxIDCountry = "UY"
	CustomerTaxIDCountryUz CustomerTaxIDCountry = "UZ"
	CustomerTaxIDCountryVe CustomerTaxIDCountry = "VE"
	CustomerTaxIDCountryVn CustomerTaxIDCountry = "VN"
	CustomerTaxIDCountryZa CustomerTaxIDCountry = "ZA"
	CustomerTaxIDCountryZm CustomerTaxIDCountry = "ZM"
	CustomerTaxIDCountryZw CustomerTaxIDCountry = "ZW"
)

func (r CustomerTaxIDCountry) IsKnown() bool {
	switch r {
	case CustomerTaxIDCountryAd, CustomerTaxIDCountryAe, CustomerTaxIDCountryAl, CustomerTaxIDCountryAm, CustomerTaxIDCountryAo, CustomerTaxIDCountryAr, CustomerTaxIDCountryAt, CustomerTaxIDCountryAu, CustomerTaxIDCountryAw, CustomerTaxIDCountryAz, CustomerTaxIDCountryBa, CustomerTaxIDCountryBb, CustomerTaxIDCountryBd, CustomerTaxIDCountryBe, CustomerTaxIDCountryBf, CustomerTaxIDCountryBg, CustomerTaxIDCountryBh, CustomerTaxIDCountryBj, CustomerTaxIDCountryBo, CustomerTaxIDCountryBr, CustomerTaxIDCountryBs, CustomerTaxIDCountryBy, CustomerTaxIDCountryCa, CustomerTaxIDCountryCd, CustomerTaxIDCountryCh, CustomerTaxIDCountryCl, CustomerTaxIDCountryCm, CustomerTaxIDCountryCn, CustomerTaxIDCountryCo, CustomerTaxIDCountryCr, CustomerTaxIDCountryCv, CustomerTaxIDCountryDe, CustomerTaxIDCountryCy, CustomerTaxIDCountryCz, CustomerTaxIDCountryDk, CustomerTaxIDCountryDo, CustomerTaxIDCountryEc, CustomerTaxIDCountryEe, CustomerTaxIDCountryEg, CustomerTaxIDCountryEs, CustomerTaxIDCountryEt, CustomerTaxIDCountryEu, CustomerTaxIDCountryFi, CustomerTaxIDCountryFr, CustomerTaxIDCountryGB, CustomerTaxIDCountryGe, CustomerTaxIDCountryGn, CustomerTaxIDCountryGr, CustomerTaxIDCountryHk, CustomerTaxIDCountryHr, CustomerTaxIDCountryHu, CustomerTaxIDCountryID, CustomerTaxIDCountryIe, CustomerTaxIDCountryIl, CustomerTaxIDCountryIn, CustomerTaxIDCountryIs, CustomerTaxIDCountryIt, CustomerTaxIDCountryJp, CustomerTaxIDCountryKe, CustomerTaxIDCountryKg, CustomerTaxIDCountryKh, CustomerTaxIDCountryKr, CustomerTaxIDCountryKz, CustomerTaxIDCountryLa, CustomerTaxIDCountryLi, CustomerTaxIDCountryLt, CustomerTaxIDCountryLu, CustomerTaxIDCountryLv, CustomerTaxIDCountryMa, CustomerTaxIDCountryMd, CustomerTaxIDCountryMe, CustomerTaxIDCountryMk, CustomerTaxIDCountryMr, CustomerTaxIDCountryMt, CustomerTaxIDCountryMx, CustomerTaxIDCountryMy, CustomerTaxIDCountryNg, CustomerTaxIDCountryNl, CustomerTaxIDCountryNo, CustomerTaxIDCountryNp, CustomerTaxIDCountryNz, CustomerTaxIDCountryOm, CustomerTaxIDCountryPe, CustomerTaxIDCountryPh, CustomerTaxIDCountryPl, CustomerTaxIDCountryPt, CustomerTaxIDCountryRo, CustomerTaxIDCountryRs, CustomerTaxIDCountryRu, CustomerTaxIDCountrySa, CustomerTaxIDCountrySe, CustomerTaxIDCountrySg, CustomerTaxIDCountrySi, CustomerTaxIDCountrySk, CustomerTaxIDCountrySn, CustomerTaxIDCountrySr, CustomerTaxIDCountrySv, CustomerTaxIDCountryTh, CustomerTaxIDCountryTj, CustomerTaxIDCountryTr, CustomerTaxIDCountryTw, CustomerTaxIDCountryTz, CustomerTaxIDCountryUa, CustomerTaxIDCountryUg, CustomerTaxIDCountryUs, CustomerTaxIDCountryUy, CustomerTaxIDCountryUz, CustomerTaxIDCountryVe, CustomerTaxIDCountryVn, CustomerTaxIDCountryZa, CustomerTaxIDCountryZm, CustomerTaxIDCountryZw:
		return true
	}
	return false
}

type CustomerTaxIDType string

const (
	CustomerTaxIDTypeAdNrt    CustomerTaxIDType = "ad_nrt"
	CustomerTaxIDTypeAeTrn    CustomerTaxIDType = "ae_trn"
	CustomerTaxIDTypeAlTin    CustomerTaxIDType = "al_tin"
	CustomerTaxIDTypeAmTin    CustomerTaxIDType = "am_tin"
	CustomerTaxIDTypeAoTin    CustomerTaxIDType = "ao_tin"
	CustomerTaxIDTypeArCuit   CustomerTaxIDType = "ar_cuit"
	CustomerTaxIDTypeEuVat    CustomerTaxIDType = "eu_vat"
	CustomerTaxIDTypeAuAbn    CustomerTaxIDType = "au_abn"
	CustomerTaxIDTypeAuArn    CustomerTaxIDType = "au_arn"
	CustomerTaxIDTypeAwTin    CustomerTaxIDType = "aw_tin"
	CustomerTaxIDTypeAzTin    CustomerTaxIDType = "az_tin"
	CustomerTaxIDTypeBaTin    CustomerTaxIDType = "ba_tin"
	CustomerTaxIDTypeBbTin    CustomerTaxIDType = "bb_tin"
	CustomerTaxIDTypeBdBin    CustomerTaxIDType = "bd_bin"
	CustomerTaxIDTypeBfIfu    CustomerTaxIDType = "bf_ifu"
	CustomerTaxIDTypeBgUic    CustomerTaxIDType = "bg_uic"
	CustomerTaxIDTypeBhVat    CustomerTaxIDType = "bh_vat"
	CustomerTaxIDTypeBjIfu    CustomerTaxIDType = "bj_ifu"
	CustomerTaxIDTypeBoTin    CustomerTaxIDType = "bo_tin"
	CustomerTaxIDTypeBrCnpj   CustomerTaxIDType = "br_cnpj"
	CustomerTaxIDTypeBrCpf    CustomerTaxIDType = "br_cpf"
	CustomerTaxIDTypeBsTin    CustomerTaxIDType = "bs_tin"
	CustomerTaxIDTypeByTin    CustomerTaxIDType = "by_tin"
	CustomerTaxIDTypeCaBn     CustomerTaxIDType = "ca_bn"
	CustomerTaxIDTypeCaGstHst CustomerTaxIDType = "ca_gst_hst"
	CustomerTaxIDTypeCaPstBc  CustomerTaxIDType = "ca_pst_bc"
	CustomerTaxIDTypeCaPstMB  CustomerTaxIDType = "ca_pst_mb"
	CustomerTaxIDTypeCaPstSk  CustomerTaxIDType = "ca_pst_sk"
	CustomerTaxIDTypeCaQst    CustomerTaxIDType = "ca_qst"
	CustomerTaxIDTypeCdNif    CustomerTaxIDType = "cd_nif"
	CustomerTaxIDTypeChUid    CustomerTaxIDType = "ch_uid"
	CustomerTaxIDTypeChVat    CustomerTaxIDType = "ch_vat"
	CustomerTaxIDTypeClTin    CustomerTaxIDType = "cl_tin"
	CustomerTaxIDTypeCmNiu    CustomerTaxIDType = "cm_niu"
	CustomerTaxIDTypeCnTin    CustomerTaxIDType = "cn_tin"
	CustomerTaxIDTypeCoNit    CustomerTaxIDType = "co_nit"
	CustomerTaxIDTypeCrTin    CustomerTaxIDType = "cr_tin"
	CustomerTaxIDTypeCvNif    CustomerTaxIDType = "cv_nif"
	CustomerTaxIDTypeDeStn    CustomerTaxIDType = "de_stn"
	CustomerTaxIDTypeDoRcn    CustomerTaxIDType = "do_rcn"
	CustomerTaxIDTypeEcRuc    CustomerTaxIDType = "ec_ruc"
	CustomerTaxIDTypeEgTin    CustomerTaxIDType = "eg_tin"
	CustomerTaxIDTypeEsCif    CustomerTaxIDType = "es_cif"
	CustomerTaxIDTypeEtTin    CustomerTaxIDType = "et_tin"
	CustomerTaxIDTypeEuOssVat CustomerTaxIDType = "eu_oss_vat"
	CustomerTaxIDTypeGBVat    CustomerTaxIDType = "gb_vat"
	CustomerTaxIDTypeGeVat    CustomerTaxIDType = "ge_vat"
	CustomerTaxIDTypeGnNif    CustomerTaxIDType = "gn_nif"
	CustomerTaxIDTypeHkBr     CustomerTaxIDType = "hk_br"
	CustomerTaxIDTypeHrOib    CustomerTaxIDType = "hr_oib"
	CustomerTaxIDTypeHuTin    CustomerTaxIDType = "hu_tin"
	CustomerTaxIDTypeIDNpwp   CustomerTaxIDType = "id_npwp"
	CustomerTaxIDTypeIlVat    CustomerTaxIDType = "il_vat"
	CustomerTaxIDTypeInGst    CustomerTaxIDType = "in_gst"
	CustomerTaxIDTypeIsVat    CustomerTaxIDType = "is_vat"
	CustomerTaxIDTypeJpCn     CustomerTaxIDType = "jp_cn"
	CustomerTaxIDTypeJpRn     CustomerTaxIDType = "jp_rn"
	CustomerTaxIDTypeJpTrn    CustomerTaxIDType = "jp_trn"
	CustomerTaxIDTypeKePin    CustomerTaxIDType = "ke_pin"
	CustomerTaxIDTypeKgTin    CustomerTaxIDType = "kg_tin"
	CustomerTaxIDTypeKhTin    CustomerTaxIDType = "kh_tin"
	CustomerTaxIDTypeKrBrn    CustomerTaxIDType = "kr_brn"
	CustomerTaxIDTypeKzBin    CustomerTaxIDType = "kz_bin"
	CustomerTaxIDTypeLaTin    CustomerTaxIDType = "la_tin"
	CustomerTaxIDTypeLiUid    CustomerTaxIDType = "li_uid"
	CustomerTaxIDTypeLiVat    CustomerTaxIDType = "li_vat"
	CustomerTaxIDTypeMaVat    CustomerTaxIDType = "ma_vat"
	CustomerTaxIDTypeMdVat    CustomerTaxIDType = "md_vat"
	CustomerTaxIDTypeMePib    CustomerTaxIDType = "me_pib"
	CustomerTaxIDTypeMkVat    CustomerTaxIDType = "mk_vat"
	CustomerTaxIDTypeMrNif    CustomerTaxIDType = "mr_nif"
	CustomerTaxIDTypeMxRfc    CustomerTaxIDType = "mx_rfc"
	CustomerTaxIDTypeMyFrp    CustomerTaxIDType = "my_frp"
	CustomerTaxIDTypeMyItn    CustomerTaxIDType = "my_itn"
	CustomerTaxIDTypeMySst    CustomerTaxIDType = "my_sst"
	CustomerTaxIDTypeNgTin    CustomerTaxIDType = "ng_tin"
	CustomerTaxIDTypeNoVat    CustomerTaxIDType = "no_vat"
	CustomerTaxIDTypeNoVoec   CustomerTaxIDType = "no_voec"
	CustomerTaxIDTypeNpPan    CustomerTaxIDType = "np_pan"
	CustomerTaxIDTypeNzGst    CustomerTaxIDType = "nz_gst"
	CustomerTaxIDTypeOmVat    CustomerTaxIDType = "om_vat"
	CustomerTaxIDTypePeRuc    CustomerTaxIDType = "pe_ruc"
	CustomerTaxIDTypePhTin    CustomerTaxIDType = "ph_tin"
	CustomerTaxIDTypeRoTin    CustomerTaxIDType = "ro_tin"
	CustomerTaxIDTypeRsPib    CustomerTaxIDType = "rs_pib"
	CustomerTaxIDTypeRuInn    CustomerTaxIDType = "ru_inn"
	CustomerTaxIDTypeRuKpp    CustomerTaxIDType = "ru_kpp"
	CustomerTaxIDTypeSaVat    CustomerTaxIDType = "sa_vat"
	CustomerTaxIDTypeSgGst    CustomerTaxIDType = "sg_gst"
	CustomerTaxIDTypeSgUen    CustomerTaxIDType = "sg_uen"
	CustomerTaxIDTypeSiTin    CustomerTaxIDType = "si_tin"
	CustomerTaxIDTypeSnNinea  CustomerTaxIDType = "sn_ninea"
	CustomerTaxIDTypeSrFin    CustomerTaxIDType = "sr_fin"
	CustomerTaxIDTypeSvNit    CustomerTaxIDType = "sv_nit"
	CustomerTaxIDTypeThVat    CustomerTaxIDType = "th_vat"
	CustomerTaxIDTypeTjTin    CustomerTaxIDType = "tj_tin"
	CustomerTaxIDTypeTrTin    CustomerTaxIDType = "tr_tin"
	CustomerTaxIDTypeTwVat    CustomerTaxIDType = "tw_vat"
	CustomerTaxIDTypeTzVat    CustomerTaxIDType = "tz_vat"
	CustomerTaxIDTypeUaVat    CustomerTaxIDType = "ua_vat"
	CustomerTaxIDTypeUgTin    CustomerTaxIDType = "ug_tin"
	CustomerTaxIDTypeUsEin    CustomerTaxIDType = "us_ein"
	CustomerTaxIDTypeUyRuc    CustomerTaxIDType = "uy_ruc"
	CustomerTaxIDTypeUzTin    CustomerTaxIDType = "uz_tin"
	CustomerTaxIDTypeUzVat    CustomerTaxIDType = "uz_vat"
	CustomerTaxIDTypeVeRif    CustomerTaxIDType = "ve_rif"
	CustomerTaxIDTypeVnTin    CustomerTaxIDType = "vn_tin"
	CustomerTaxIDTypeZaVat    CustomerTaxIDType = "za_vat"
	CustomerTaxIDTypeZmTin    CustomerTaxIDType = "zm_tin"
	CustomerTaxIDTypeZwTin    CustomerTaxIDType = "zw_tin"
)

func (r CustomerTaxIDType) IsKnown() bool {
	switch r {
	case CustomerTaxIDTypeAdNrt, CustomerTaxIDTypeAeTrn, CustomerTaxIDTypeAlTin, CustomerTaxIDTypeAmTin, CustomerTaxIDTypeAoTin, CustomerTaxIDTypeArCuit, CustomerTaxIDTypeEuVat, CustomerTaxIDTypeAuAbn, CustomerTaxIDTypeAuArn, CustomerTaxIDTypeAwTin, CustomerTaxIDTypeAzTin, CustomerTaxIDTypeBaTin, CustomerTaxIDTypeBbTin, CustomerTaxIDTypeBdBin, CustomerTaxIDTypeBfIfu, CustomerTaxIDTypeBgUic, CustomerTaxIDTypeBhVat, CustomerTaxIDTypeBjIfu, CustomerTaxIDTypeBoTin, CustomerTaxIDTypeBrCnpj, CustomerTaxIDTypeBrCpf, CustomerTaxIDTypeBsTin, CustomerTaxIDTypeByTin, CustomerTaxIDTypeCaBn, CustomerTaxIDTypeCaGstHst, CustomerTaxIDTypeCaPstBc, CustomerTaxIDTypeCaPstMB, CustomerTaxIDTypeCaPstSk, CustomerTaxIDTypeCaQst, CustomerTaxIDTypeCdNif, CustomerTaxIDTypeChUid, CustomerTaxIDTypeChVat, CustomerTaxIDTypeClTin, CustomerTaxIDTypeCmNiu, CustomerTaxIDTypeCnTin, CustomerTaxIDTypeCoNit, CustomerTaxIDTypeCrTin, CustomerTaxIDTypeCvNif, CustomerTaxIDTypeDeStn, CustomerTaxIDTypeDoRcn, CustomerTaxIDTypeEcRuc, CustomerTaxIDTypeEgTin, CustomerTaxIDTypeEsCif, CustomerTaxIDTypeEtTin, CustomerTaxIDTypeEuOssVat, CustomerTaxIDTypeGBVat, CustomerTaxIDTypeGeVat, CustomerTaxIDTypeGnNif, CustomerTaxIDTypeHkBr, CustomerTaxIDTypeHrOib, CustomerTaxIDTypeHuTin, CustomerTaxIDTypeIDNpwp, CustomerTaxIDTypeIlVat, CustomerTaxIDTypeInGst, CustomerTaxIDTypeIsVat, CustomerTaxIDTypeJpCn, CustomerTaxIDTypeJpRn, CustomerTaxIDTypeJpTrn, CustomerTaxIDTypeKePin, CustomerTaxIDTypeKgTin, CustomerTaxIDTypeKhTin, CustomerTaxIDTypeKrBrn, CustomerTaxIDTypeKzBin, CustomerTaxIDTypeLaTin, CustomerTaxIDTypeLiUid, CustomerTaxIDTypeLiVat, CustomerTaxIDTypeMaVat, CustomerTaxIDTypeMdVat, CustomerTaxIDTypeMePib, CustomerTaxIDTypeMkVat, CustomerTaxIDTypeMrNif, CustomerTaxIDTypeMxRfc, CustomerTaxIDTypeMyFrp, CustomerTaxIDTypeMyItn, CustomerTaxIDTypeMySst, CustomerTaxIDTypeNgTin, CustomerTaxIDTypeNoVat, CustomerTaxIDTypeNoVoec, CustomerTaxIDTypeNpPan, CustomerTaxIDTypeNzGst, CustomerTaxIDTypeOmVat, CustomerTaxIDTypePeRuc, CustomerTaxIDTypePhTin, CustomerTaxIDTypeRoTin, CustomerTaxIDTypeRsPib, CustomerTaxIDTypeRuInn, CustomerTaxIDTypeRuKpp, CustomerTaxIDTypeSaVat, CustomerTaxIDTypeSgGst, CustomerTaxIDTypeSgUen, CustomerTaxIDTypeSiTin, CustomerTaxIDTypeSnNinea, CustomerTaxIDTypeSrFin, CustomerTaxIDTypeSvNit, CustomerTaxIDTypeThVat, CustomerTaxIDTypeTjTin, CustomerTaxIDTypeTrTin, CustomerTaxIDTypeTwVat, CustomerTaxIDTypeTzVat, CustomerTaxIDTypeUaVat, CustomerTaxIDTypeUgTin, CustomerTaxIDTypeUsEin, CustomerTaxIDTypeUyRuc, CustomerTaxIDTypeUzTin, CustomerTaxIDTypeUzVat, CustomerTaxIDTypeVeRif, CustomerTaxIDTypeVnTin, CustomerTaxIDTypeZaVat, CustomerTaxIDTypeZmTin, CustomerTaxIDTypeZwTin:
		return true
	}
	return false
}

// Tax IDs are commonly required to be displayed on customer invoices, which are
// added to the headers of invoices.
//
// ### Supported Tax ID Countries and Types
//
// | Country                | Type         | Description                                                                                             |
// | ---------------------- | ------------ | ------------------------------------------------------------------------------------------------------- |
// | Albania                | `al_tin`     | Albania Tax Identification Number                                                                       |
// | Andorra                | `ad_nrt`     | Andorran NRT Number                                                                                     |
// | Angola                 | `ao_tin`     | Angola Tax Identification Number                                                                        |
// | Argentina              | `ar_cuit`    | Argentinian Tax ID Number                                                                               |
// | Armenia                | `am_tin`     | Armenia Tax Identification Number                                                                       |
// | Aruba                  | `aw_tin`     | Aruba Tax Identification Number                                                                         |
// | Australia              | `au_abn`     | Australian Business Number (AU ABN)                                                                     |
// | Australia              | `au_arn`     | Australian Taxation Office Reference Number                                                             |
// | Austria                | `eu_vat`     | European VAT Number                                                                                     |
// | Azerbaijan             | `az_tin`     | Azerbaijan Tax Identification Number                                                                    |
// | Bahamas                | `bs_tin`     | Bahamas Tax Identification Number                                                                       |
// | Bahrain                | `bh_vat`     | Bahraini VAT Number                                                                                     |
// | Bangladesh             | `bd_bin`     | Bangladesh Business Identification Number                                                               |
// | Barbados               | `bb_tin`     | Barbados Tax Identification Number                                                                      |
// | Belarus                | `by_tin`     | Belarus TIN Number                                                                                      |
// | Belgium                | `eu_vat`     | European VAT Number                                                                                     |
// | Benin                  | `bj_ifu`     | Benin Tax Identification Number (Identifiant Fiscal Unique)                                             |
// | Bolivia                | `bo_tin`     | Bolivian Tax ID                                                                                         |
// | Bosnia and Herzegovina | `ba_tin`     | Bosnia and Herzegovina Tax Identification Number                                                        |
// | Brazil                 | `br_cnpj`    | Brazilian CNPJ Number                                                                                   |
// | Brazil                 | `br_cpf`     | Brazilian CPF Number                                                                                    |
// | Bulgaria               | `bg_uic`     | Bulgaria Unified Identification Code                                                                    |
// | Bulgaria               | `eu_vat`     | European VAT Number                                                                                     |
// | Burkina Faso           | `bf_ifu`     | Burkina Faso Tax Identification Number (Numéro d'Identifiant Fiscal Unique)                             |
// | Cambodia               | `kh_tin`     | Cambodia Tax Identification Number                                                                      |
// | Cameroon               | `cm_niu`     | Cameroon Tax Identification Number (Numéro d'Identifiant fiscal Unique)                                 |
// | Canada                 | `ca_bn`      | Canadian BN                                                                                             |
// | Canada                 | `ca_gst_hst` | Canadian GST/HST Number                                                                                 |
// | Canada                 | `ca_pst_bc`  | Canadian PST Number (British Columbia)                                                                  |
// | Canada                 | `ca_pst_mb`  | Canadian PST Number (Manitoba)                                                                          |
// | Canada                 | `ca_pst_sk`  | Canadian PST Number (Saskatchewan)                                                                      |
// | Canada                 | `ca_qst`     | Canadian QST Number (Québec)                                                                            |
// | Cape Verde             | `cv_nif`     | Cape Verde Tax Identification Number (Número de Identificação Fiscal)                                   |
// | Chile                  | `cl_tin`     | Chilean TIN                                                                                             |
// | China                  | `cn_tin`     | Chinese Tax ID                                                                                          |
// | Colombia               | `co_nit`     | Colombian NIT Number                                                                                    |
// | Congo-Kinshasa         | `cd_nif`     | Congo (DR) Tax Identification Number (Número de Identificação Fiscal)                                   |
// | Costa Rica             | `cr_tin`     | Costa Rican Tax ID                                                                                      |
// | Croatia                | `eu_vat`     | European VAT Number                                                                                     |
// | Croatia                | `hr_oib`     | Croatian Personal Identification Number (OIB)                                                           |
// | Cyprus                 | `eu_vat`     | European VAT Number                                                                                     |
// | Czech Republic         | `eu_vat`     | European VAT Number                                                                                     |
// | Denmark                | `eu_vat`     | European VAT Number                                                                                     |
// | Dominican Republic     | `do_rcn`     | Dominican RCN Number                                                                                    |
// | Ecuador                | `ec_ruc`     | Ecuadorian RUC Number                                                                                   |
// | Egypt                  | `eg_tin`     | Egyptian Tax Identification Number                                                                      |
// | El Salvador            | `sv_nit`     | El Salvadorian NIT Number                                                                               |
// | Estonia                | `eu_vat`     | European VAT Number                                                                                     |
// | Ethiopia               | `et_tin`     | Ethiopia Tax Identification Number                                                                      |
// | European Union         | `eu_oss_vat` | European One Stop Shop VAT Number for non-Union scheme                                                  |
// | Finland                | `eu_vat`     | European VAT Number                                                                                     |
// | France                 | `eu_vat`     | European VAT Number                                                                                     |
// | Georgia                | `ge_vat`     | Georgian VAT                                                                                            |
// | Germany                | `de_stn`     | German Tax Number (Steuernummer)                                                                        |
// | Germany                | `eu_vat`     | European VAT Number                                                                                     |
// | Greece                 | `eu_vat`     | European VAT Number                                                                                     |
// | Guinea                 | `gn_nif`     | Guinea Tax Identification Number (Número de Identificação Fiscal)                                       |
// | Hong Kong              | `hk_br`      | Hong Kong BR Number                                                                                     |
// | Hungary                | `eu_vat`     | European VAT Number                                                                                     |
// | Hungary                | `hu_tin`     | Hungary Tax Number (adószám)                                                                            |
// | Iceland                | `is_vat`     | Icelandic VAT                                                                                           |
// | India                  | `in_gst`     | Indian GST Number                                                                                       |
// | Indonesia              | `id_npwp`    | Indonesian NPWP Number                                                                                  |
// | Ireland                | `eu_vat`     | European VAT Number                                                                                     |
// | Israel                 | `il_vat`     | Israel VAT                                                                                              |
// | Italy                  | `eu_vat`     | European VAT Number                                                                                     |
// | Japan                  | `jp_cn`      | Japanese Corporate Number (_Hōjin Bangō_)                                                               |
// | Japan                  | `jp_rn`      | Japanese Registered Foreign Businesses' Registration Number (_Tōroku Kokugai Jigyōsha no Tōroku Bangō_) |
// | Japan                  | `jp_trn`     | Japanese Tax Registration Number (_Tōroku Bangō_)                                                       |
// | Kazakhstan             | `kz_bin`     | Kazakhstani Business Identification Number                                                              |
// | Kenya                  | `ke_pin`     | Kenya Revenue Authority Personal Identification Number                                                  |
// | Kyrgyzstan             | `kg_tin`     | Kyrgyzstan Tax Identification Number                                                                    |
// | Laos                   | `la_tin`     | Laos Tax Identification Number                                                                          |
// | Latvia                 | `eu_vat`     | European VAT Number                                                                                     |
// | Liechtenstein          | `li_uid`     | Liechtensteinian UID Number                                                                             |
// | Liechtenstein          | `li_vat`     | Liechtenstein VAT Number                                                                                |
// | Lithuania              | `eu_vat`     | European VAT Number                                                                                     |
// | Luxembourg             | `eu_vat`     | European VAT Number                                                                                     |
// | Malaysia               | `my_frp`     | Malaysian FRP Number                                                                                    |
// | Malaysia               | `my_itn`     | Malaysian ITN                                                                                           |
// | Malaysia               | `my_sst`     | Malaysian SST Number                                                                                    |
// | Malta                  | `eu_vat`     | European VAT Number                                                                                     |
// | Mauritania             | `mr_nif`     | Mauritania Tax Identification Number (Número de Identificação Fiscal)                                   |
// | Mexico                 | `mx_rfc`     | Mexican RFC Number                                                                                      |
// | Moldova                | `md_vat`     | Moldova VAT Number                                                                                      |
// | Montenegro             | `me_pib`     | Montenegro PIB Number                                                                                   |
// | Morocco                | `ma_vat`     | Morocco VAT Number                                                                                      |
// | Nepal                  | `np_pan`     | Nepal PAN Number                                                                                        |
// | Netherlands            | `eu_vat`     | European VAT Number                                                                                     |
// | New Zealand            | `nz_gst`     | New Zealand GST Number                                                                                  |
// | Nigeria                | `ng_tin`     | Nigerian Tax Identification Number                                                                      |
// | North Macedonia        | `mk_vat`     | North Macedonia VAT Number                                                                              |
// | Northern Ireland       | `eu_vat`     | Northern Ireland VAT Number                                                                             |
// | Norway                 | `no_vat`     | Norwegian VAT Number                                                                                    |
// | Norway                 | `no_voec`    | Norwegian VAT on e-commerce Number                                                                      |
// | Oman                   | `om_vat`     | Omani VAT Number                                                                                        |
// | Peru                   | `pe_ruc`     | Peruvian RUC Number                                                                                     |
// | Philippines            | `ph_tin`     | Philippines Tax Identification Number                                                                   |
// | Poland                 | `eu_vat`     | European VAT Number                                                                                     |
// | Portugal               | `eu_vat`     | European VAT Number                                                                                     |
// | Romania                | `eu_vat`     | European VAT Number                                                                                     |
// | Romania                | `ro_tin`     | Romanian Tax ID Number                                                                                  |
// | Russia                 | `ru_inn`     | Russian INN                                                                                             |
// | Russia                 | `ru_kpp`     | Russian KPP                                                                                             |
// | Saudi Arabia           | `sa_vat`     | Saudi Arabia VAT                                                                                        |
// | Senegal                | `sn_ninea`   | Senegal NINEA Number                                                                                    |
// | Serbia                 | `rs_pib`     | Serbian PIB Number                                                                                      |
// | Singapore              | `sg_gst`     | Singaporean GST                                                                                         |
// | Singapore              | `sg_uen`     | Singaporean UEN                                                                                         |
// | Slovakia               | `eu_vat`     | European VAT Number                                                                                     |
// | Slovenia               | `eu_vat`     | European VAT Number                                                                                     |
// | Slovenia               | `si_tin`     | Slovenia Tax Number (davčna številka)                                                                   |
// | South Africa           | `za_vat`     | South African VAT Number                                                                                |
// | South Korea            | `kr_brn`     | Korean BRN                                                                                              |
// | Spain                  | `es_cif`     | Spanish NIF Number (previously Spanish CIF Number)                                                      |
// | Spain                  | `eu_vat`     | European VAT Number                                                                                     |
// | Suriname               | `sr_fin`     | Suriname FIN Number                                                                                     |
// | Sweden                 | `eu_vat`     | European VAT Number                                                                                     |
// | Switzerland            | `ch_uid`     | Switzerland UID Number                                                                                  |
// | Switzerland            | `ch_vat`     | Switzerland VAT Number                                                                                  |
// | Taiwan                 | `tw_vat`     | Taiwanese VAT                                                                                           |
// | Tajikistan             | `tj_tin`     | Tajikistan Tax Identification Number                                                                    |
// | Tanzania               | `tz_vat`     | Tanzania VAT Number                                                                                     |
// | Thailand               | `th_vat`     | Thai VAT                                                                                                |
// | Turkey                 | `tr_tin`     | Turkish Tax Identification Number                                                                       |
// | Uganda                 | `ug_tin`     | Uganda Tax Identification Number                                                                        |
// | Ukraine                | `ua_vat`     | Ukrainian VAT                                                                                           |
// | United Arab Emirates   | `ae_trn`     | United Arab Emirates TRN                                                                                |
// | United Kingdom         | `gb_vat`     | United Kingdom VAT Number                                                                               |
// | United States          | `us_ein`     | United States EIN                                                                                       |
// | Uruguay                | `uy_ruc`     | Uruguayan RUC Number                                                                                    |
// | Uzbekistan             | `uz_tin`     | Uzbekistan TIN Number                                                                                   |
// | Uzbekistan             | `uz_vat`     | Uzbekistan VAT Number                                                                                   |
// | Venezuela              | `ve_rif`     | Venezuelan RIF Number                                                                                   |
// | Vietnam                | `vn_tin`     | Vietnamese Tax ID Number                                                                                |
// | Zambia                 | `zm_tin`     | Zambia Tax Identification Number                                                                        |
// | Zimbabwe               | `zw_tin`     | Zimbabwe Tax Identification Number                                                                      |
type CustomerTaxIDParam struct {
	Country param.Field[CustomerTaxIDCountry] `json:"country,required"`
	Type    param.Field[CustomerTaxIDType]    `json:"type,required"`
	Value   param.Field[string]               `json:"value,required"`
}

func (r CustomerTaxIDParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type DimensionalPriceConfiguration struct {
	DimensionValues         []string                          `json:"dimension_values,required"`
	DimensionalPriceGroupID string                            `json:"dimensional_price_group_id,required"`
	JSON                    dimensionalPriceConfigurationJSON `json:"-"`
}

// dimensionalPriceConfigurationJSON contains the JSON metadata for the struct
// [DimensionalPriceConfiguration]
type dimensionalPriceConfigurationJSON struct {
	DimensionValues         apijson.Field
	DimensionalPriceGroupID apijson.Field
	raw                     string
	ExtraFields             map[string]apijson.Field
}

func (r *DimensionalPriceConfiguration) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r dimensionalPriceConfigurationJSON) RawJSON() string {
	return r.raw
}

type Discount struct {
	DiscountType DiscountDiscountType `json:"discount_type,required"`
	// Only available if discount_type is `amount`.
	AmountDiscount string `json:"amount_discount"`
	// This field can have the runtime type of [[]string].
	AppliesToPriceIDs interface{} `json:"applies_to_price_ids"`
	// This field can have the runtime type of [[]TransformPriceFilter].
	Filters interface{} `json:"filters"`
	// Only available if discount_type is `percentage`. This is a number between 0
	// and 1.
	PercentageDiscount float64 `json:"percentage_discount"`
	Reason             string  `json:"reason,nullable"`
	// Only available if discount_type is `trial`
	TrialAmountDiscount string `json:"trial_amount_discount,nullable"`
	// Only available if discount_type is `trial`
	TrialPercentageDiscount float64 `json:"trial_percentage_discount,nullable"`
	// Only available if discount_type is `usage`. Number of usage units that this
	// discount is for
	UsageDiscount float64      `json:"usage_discount"`
	JSON          discountJSON `json:"-"`
	union         DiscountUnion
}

// discountJSON contains the JSON metadata for the struct [Discount]
type discountJSON struct {
	DiscountType            apijson.Field
	AmountDiscount          apijson.Field
	AppliesToPriceIDs       apijson.Field
	Filters                 apijson.Field
	PercentageDiscount      apijson.Field
	Reason                  apijson.Field
	TrialAmountDiscount     apijson.Field
	TrialPercentageDiscount apijson.Field
	UsageDiscount           apijson.Field
	raw                     string
	ExtraFields             map[string]apijson.Field
}

func (r discountJSON) RawJSON() string {
	return r.raw
}

func (r *Discount) UnmarshalJSON(data []byte) (err error) {
	*r = Discount{}
	err = apijson.UnmarshalRoot(data, &r.union)
	if err != nil {
		return err
	}
	return apijson.Port(r.union, &r)
}

// AsUnion returns a [DiscountUnion] interface which you can cast to the specific
// types for more type safety.
//
// Possible runtime types of the union are [PercentageDiscount], [TrialDiscount],
// [UsageDiscount], [AmountDiscount].
func (r Discount) AsUnion() DiscountUnion {
	return r.union
}

// Union satisfied by [PercentageDiscount], [TrialDiscount], [UsageDiscount] or
// [AmountDiscount].
type DiscountUnion interface {
	ImplementsDiscount()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*DiscountUnion)(nil)).Elem(),
		"discount_type",
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(PercentageDiscount{}),
			DiscriminatorValue: "percentage",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(TrialDiscount{}),
			DiscriminatorValue: "trial",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(UsageDiscount{}),
			DiscriminatorValue: "usage",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(AmountDiscount{}),
			DiscriminatorValue: "amount",
		},
	)
}

type DiscountDiscountType string

const (
	DiscountDiscountTypePercentage DiscountDiscountType = "percentage"
	DiscountDiscountTypeTrial      DiscountDiscountType = "trial"
	DiscountDiscountTypeUsage      DiscountDiscountType = "usage"
	DiscountDiscountTypeAmount     DiscountDiscountType = "amount"
)

func (r DiscountDiscountType) IsKnown() bool {
	switch r {
	case DiscountDiscountTypePercentage, DiscountDiscountTypeTrial, DiscountDiscountTypeUsage, DiscountDiscountTypeAmount:
		return true
	}
	return false
}

type DiscountParam struct {
	DiscountType param.Field[DiscountDiscountType] `json:"discount_type,required"`
	// Only available if discount_type is `amount`.
	AmountDiscount    param.Field[string]      `json:"amount_discount"`
	AppliesToPriceIDs param.Field[interface{}] `json:"applies_to_price_ids"`
	Filters           param.Field[interface{}] `json:"filters"`
	// Only available if discount_type is `percentage`. This is a number between 0
	// and 1.
	PercentageDiscount param.Field[float64] `json:"percentage_discount"`
	Reason             param.Field[string]  `json:"reason"`
	// Only available if discount_type is `trial`
	TrialAmountDiscount param.Field[string] `json:"trial_amount_discount"`
	// Only available if discount_type is `trial`
	TrialPercentageDiscount param.Field[float64] `json:"trial_percentage_discount"`
	// Only available if discount_type is `usage`. Number of usage units that this
	// discount is for
	UsageDiscount param.Field[float64] `json:"usage_discount"`
}

func (r DiscountParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r DiscountParam) ImplementsDiscountUnionParam() {}

// Satisfied by [shared.PercentageDiscountParam], [shared.TrialDiscountParam],
// [shared.UsageDiscountParam], [shared.AmountDiscountParam], [DiscountParam].
type DiscountUnionParam interface {
	ImplementsDiscountUnionParam()
}

type FixedFeeQuantityScheduleEntry struct {
	EndDate   time.Time                         `json:"end_date,required,nullable" format:"date-time"`
	PriceID   string                            `json:"price_id,required"`
	Quantity  float64                           `json:"quantity,required"`
	StartDate time.Time                         `json:"start_date,required" format:"date-time"`
	JSON      fixedFeeQuantityScheduleEntryJSON `json:"-"`
}

// fixedFeeQuantityScheduleEntryJSON contains the JSON metadata for the struct
// [FixedFeeQuantityScheduleEntry]
type fixedFeeQuantityScheduleEntryJSON struct {
	EndDate     apijson.Field
	PriceID     apijson.Field
	Quantity    apijson.Field
	StartDate   apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *FixedFeeQuantityScheduleEntry) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r fixedFeeQuantityScheduleEntryJSON) RawJSON() string {
	return r.raw
}

type FixedFeeQuantityTransition struct {
	EffectiveDate time.Time                      `json:"effective_date,required" format:"date-time"`
	PriceID       string                         `json:"price_id,required"`
	Quantity      int64                          `json:"quantity,required"`
	JSON          fixedFeeQuantityTransitionJSON `json:"-"`
}

// fixedFeeQuantityTransitionJSON contains the JSON metadata for the struct
// [FixedFeeQuantityTransition]
type fixedFeeQuantityTransitionJSON struct {
	EffectiveDate apijson.Field
	PriceID       apijson.Field
	Quantity      apijson.Field
	raw           string
	ExtraFields   map[string]apijson.Field
}

func (r *FixedFeeQuantityTransition) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r fixedFeeQuantityTransitionJSON) RawJSON() string {
	return r.raw
}

// An [`Invoice`](/core-concepts#invoice) is a fundamental billing entity,
// representing the request for payment for a single subscription. This includes a
// set of line items, which correspond to prices in the subscription's plan and can
// represent fixed recurring fees or usage-based fees. They are generated at the
// end of a billing period, or as the result of an action, such as a cancellation.
type Invoice struct {
	ID string `json:"id,required"`
	// This is the final amount required to be charged to the customer and reflects the
	// application of the customer balance to the `total` of the invoice.
	AmountDue      string                `json:"amount_due,required"`
	AutoCollection InvoiceAutoCollection `json:"auto_collection,required"`
	BillingAddress Address               `json:"billing_address,required,nullable"`
	// The creation time of the resource in Orb.
	CreatedAt time.Time `json:"created_at,required" format:"date-time"`
	// A list of credit notes associated with the invoice
	CreditNotes []InvoiceCreditNote `json:"credit_notes,required"`
	// An ISO 4217 currency string or `credits`
	Currency                    string                              `json:"currency,required"`
	Customer                    CustomerMinified                    `json:"customer,required"`
	CustomerBalanceTransactions []InvoiceCustomerBalanceTransaction `json:"customer_balance_transactions,required"`
	// Tax IDs are commonly required to be displayed on customer invoices, which are
	// added to the headers of invoices.
	//
	// ### Supported Tax ID Countries and Types
	//
	// | Country                | Type         | Description                                                                                             |
	// | ---------------------- | ------------ | ------------------------------------------------------------------------------------------------------- |
	// | Albania                | `al_tin`     | Albania Tax Identification Number                                                                       |
	// | Andorra                | `ad_nrt`     | Andorran NRT Number                                                                                     |
	// | Angola                 | `ao_tin`     | Angola Tax Identification Number                                                                        |
	// | Argentina              | `ar_cuit`    | Argentinian Tax ID Number                                                                               |
	// | Armenia                | `am_tin`     | Armenia Tax Identification Number                                                                       |
	// | Aruba                  | `aw_tin`     | Aruba Tax Identification Number                                                                         |
	// | Australia              | `au_abn`     | Australian Business Number (AU ABN)                                                                     |
	// | Australia              | `au_arn`     | Australian Taxation Office Reference Number                                                             |
	// | Austria                | `eu_vat`     | European VAT Number                                                                                     |
	// | Azerbaijan             | `az_tin`     | Azerbaijan Tax Identification Number                                                                    |
	// | Bahamas                | `bs_tin`     | Bahamas Tax Identification Number                                                                       |
	// | Bahrain                | `bh_vat`     | Bahraini VAT Number                                                                                     |
	// | Bangladesh             | `bd_bin`     | Bangladesh Business Identification Number                                                               |
	// | Barbados               | `bb_tin`     | Barbados Tax Identification Number                                                                      |
	// | Belarus                | `by_tin`     | Belarus TIN Number                                                                                      |
	// | Belgium                | `eu_vat`     | European VAT Number                                                                                     |
	// | Benin                  | `bj_ifu`     | Benin Tax Identification Number (Identifiant Fiscal Unique)                                             |
	// | Bolivia                | `bo_tin`     | Bolivian Tax ID                                                                                         |
	// | Bosnia and Herzegovina | `ba_tin`     | Bosnia and Herzegovina Tax Identification Number                                                        |
	// | Brazil                 | `br_cnpj`    | Brazilian CNPJ Number                                                                                   |
	// | Brazil                 | `br_cpf`     | Brazilian CPF Number                                                                                    |
	// | Bulgaria               | `bg_uic`     | Bulgaria Unified Identification Code                                                                    |
	// | Bulgaria               | `eu_vat`     | European VAT Number                                                                                     |
	// | Burkina Faso           | `bf_ifu`     | Burkina Faso Tax Identification Number (Numéro d'Identifiant Fiscal Unique)                             |
	// | Cambodia               | `kh_tin`     | Cambodia Tax Identification Number                                                                      |
	// | Cameroon               | `cm_niu`     | Cameroon Tax Identification Number (Numéro d'Identifiant fiscal Unique)                                 |
	// | Canada                 | `ca_bn`      | Canadian BN                                                                                             |
	// | Canada                 | `ca_gst_hst` | Canadian GST/HST Number                                                                                 |
	// | Canada                 | `ca_pst_bc`  | Canadian PST Number (British Columbia)                                                                  |
	// | Canada                 | `ca_pst_mb`  | Canadian PST Number (Manitoba)                                                                          |
	// | Canada                 | `ca_pst_sk`  | Canadian PST Number (Saskatchewan)                                                                      |
	// | Canada                 | `ca_qst`     | Canadian QST Number (Québec)                                                                            |
	// | Cape Verde             | `cv_nif`     | Cape Verde Tax Identification Number (Número de Identificação Fiscal)                                   |
	// | Chile                  | `cl_tin`     | Chilean TIN                                                                                             |
	// | China                  | `cn_tin`     | Chinese Tax ID                                                                                          |
	// | Colombia               | `co_nit`     | Colombian NIT Number                                                                                    |
	// | Congo-Kinshasa         | `cd_nif`     | Congo (DR) Tax Identification Number (Número de Identificação Fiscal)                                   |
	// | Costa Rica             | `cr_tin`     | Costa Rican Tax ID                                                                                      |
	// | Croatia                | `eu_vat`     | European VAT Number                                                                                     |
	// | Croatia                | `hr_oib`     | Croatian Personal Identification Number (OIB)                                                           |
	// | Cyprus                 | `eu_vat`     | European VAT Number                                                                                     |
	// | Czech Republic         | `eu_vat`     | European VAT Number                                                                                     |
	// | Denmark                | `eu_vat`     | European VAT Number                                                                                     |
	// | Dominican Republic     | `do_rcn`     | Dominican RCN Number                                                                                    |
	// | Ecuador                | `ec_ruc`     | Ecuadorian RUC Number                                                                                   |
	// | Egypt                  | `eg_tin`     | Egyptian Tax Identification Number                                                                      |
	// | El Salvador            | `sv_nit`     | El Salvadorian NIT Number                                                                               |
	// | Estonia                | `eu_vat`     | European VAT Number                                                                                     |
	// | Ethiopia               | `et_tin`     | Ethiopia Tax Identification Number                                                                      |
	// | European Union         | `eu_oss_vat` | European One Stop Shop VAT Number for non-Union scheme                                                  |
	// | Finland                | `eu_vat`     | European VAT Number                                                                                     |
	// | France                 | `eu_vat`     | European VAT Number                                                                                     |
	// | Georgia                | `ge_vat`     | Georgian VAT                                                                                            |
	// | Germany                | `de_stn`     | German Tax Number (Steuernummer)                                                                        |
	// | Germany                | `eu_vat`     | European VAT Number                                                                                     |
	// | Greece                 | `eu_vat`     | European VAT Number                                                                                     |
	// | Guinea                 | `gn_nif`     | Guinea Tax Identification Number (Número de Identificação Fiscal)                                       |
	// | Hong Kong              | `hk_br`      | Hong Kong BR Number                                                                                     |
	// | Hungary                | `eu_vat`     | European VAT Number                                                                                     |
	// | Hungary                | `hu_tin`     | Hungary Tax Number (adószám)                                                                            |
	// | Iceland                | `is_vat`     | Icelandic VAT                                                                                           |
	// | India                  | `in_gst`     | Indian GST Number                                                                                       |
	// | Indonesia              | `id_npwp`    | Indonesian NPWP Number                                                                                  |
	// | Ireland                | `eu_vat`     | European VAT Number                                                                                     |
	// | Israel                 | `il_vat`     | Israel VAT                                                                                              |
	// | Italy                  | `eu_vat`     | European VAT Number                                                                                     |
	// | Japan                  | `jp_cn`      | Japanese Corporate Number (_Hōjin Bangō_)                                                               |
	// | Japan                  | `jp_rn`      | Japanese Registered Foreign Businesses' Registration Number (_Tōroku Kokugai Jigyōsha no Tōroku Bangō_) |
	// | Japan                  | `jp_trn`     | Japanese Tax Registration Number (_Tōroku Bangō_)                                                       |
	// | Kazakhstan             | `kz_bin`     | Kazakhstani Business Identification Number                                                              |
	// | Kenya                  | `ke_pin`     | Kenya Revenue Authority Personal Identification Number                                                  |
	// | Kyrgyzstan             | `kg_tin`     | Kyrgyzstan Tax Identification Number                                                                    |
	// | Laos                   | `la_tin`     | Laos Tax Identification Number                                                                          |
	// | Latvia                 | `eu_vat`     | European VAT Number                                                                                     |
	// | Liechtenstein          | `li_uid`     | Liechtensteinian UID Number                                                                             |
	// | Liechtenstein          | `li_vat`     | Liechtenstein VAT Number                                                                                |
	// | Lithuania              | `eu_vat`     | European VAT Number                                                                                     |
	// | Luxembourg             | `eu_vat`     | European VAT Number                                                                                     |
	// | Malaysia               | `my_frp`     | Malaysian FRP Number                                                                                    |
	// | Malaysia               | `my_itn`     | Malaysian ITN                                                                                           |
	// | Malaysia               | `my_sst`     | Malaysian SST Number                                                                                    |
	// | Malta                  | `eu_vat`     | European VAT Number                                                                                     |
	// | Mauritania             | `mr_nif`     | Mauritania Tax Identification Number (Número de Identificação Fiscal)                                   |
	// | Mexico                 | `mx_rfc`     | Mexican RFC Number                                                                                      |
	// | Moldova                | `md_vat`     | Moldova VAT Number                                                                                      |
	// | Montenegro             | `me_pib`     | Montenegro PIB Number                                                                                   |
	// | Morocco                | `ma_vat`     | Morocco VAT Number                                                                                      |
	// | Nepal                  | `np_pan`     | Nepal PAN Number                                                                                        |
	// | Netherlands            | `eu_vat`     | European VAT Number                                                                                     |
	// | New Zealand            | `nz_gst`     | New Zealand GST Number                                                                                  |
	// | Nigeria                | `ng_tin`     | Nigerian Tax Identification Number                                                                      |
	// | North Macedonia        | `mk_vat`     | North Macedonia VAT Number                                                                              |
	// | Northern Ireland       | `eu_vat`     | Northern Ireland VAT Number                                                                             |
	// | Norway                 | `no_vat`     | Norwegian VAT Number                                                                                    |
	// | Norway                 | `no_voec`    | Norwegian VAT on e-commerce Number                                                                      |
	// | Oman                   | `om_vat`     | Omani VAT Number                                                                                        |
	// | Peru                   | `pe_ruc`     | Peruvian RUC Number                                                                                     |
	// | Philippines            | `ph_tin`     | Philippines Tax Identification Number                                                                   |
	// | Poland                 | `eu_vat`     | European VAT Number                                                                                     |
	// | Portugal               | `eu_vat`     | European VAT Number                                                                                     |
	// | Romania                | `eu_vat`     | European VAT Number                                                                                     |
	// | Romania                | `ro_tin`     | Romanian Tax ID Number                                                                                  |
	// | Russia                 | `ru_inn`     | Russian INN                                                                                             |
	// | Russia                 | `ru_kpp`     | Russian KPP                                                                                             |
	// | Saudi Arabia           | `sa_vat`     | Saudi Arabia VAT                                                                                        |
	// | Senegal                | `sn_ninea`   | Senegal NINEA Number                                                                                    |
	// | Serbia                 | `rs_pib`     | Serbian PIB Number                                                                                      |
	// | Singapore              | `sg_gst`     | Singaporean GST                                                                                         |
	// | Singapore              | `sg_uen`     | Singaporean UEN                                                                                         |
	// | Slovakia               | `eu_vat`     | European VAT Number                                                                                     |
	// | Slovenia               | `eu_vat`     | European VAT Number                                                                                     |
	// | Slovenia               | `si_tin`     | Slovenia Tax Number (davčna številka)                                                                   |
	// | South Africa           | `za_vat`     | South African VAT Number                                                                                |
	// | South Korea            | `kr_brn`     | Korean BRN                                                                                              |
	// | Spain                  | `es_cif`     | Spanish NIF Number (previously Spanish CIF Number)                                                      |
	// | Spain                  | `eu_vat`     | European VAT Number                                                                                     |
	// | Suriname               | `sr_fin`     | Suriname FIN Number                                                                                     |
	// | Sweden                 | `eu_vat`     | European VAT Number                                                                                     |
	// | Switzerland            | `ch_uid`     | Switzerland UID Number                                                                                  |
	// | Switzerland            | `ch_vat`     | Switzerland VAT Number                                                                                  |
	// | Taiwan                 | `tw_vat`     | Taiwanese VAT                                                                                           |
	// | Tajikistan             | `tj_tin`     | Tajikistan Tax Identification Number                                                                    |
	// | Tanzania               | `tz_vat`     | Tanzania VAT Number                                                                                     |
	// | Thailand               | `th_vat`     | Thai VAT                                                                                                |
	// | Turkey                 | `tr_tin`     | Turkish Tax Identification Number                                                                       |
	// | Uganda                 | `ug_tin`     | Uganda Tax Identification Number                                                                        |
	// | Ukraine                | `ua_vat`     | Ukrainian VAT                                                                                           |
	// | United Arab Emirates   | `ae_trn`     | United Arab Emirates TRN                                                                                |
	// | United Kingdom         | `gb_vat`     | United Kingdom VAT Number                                                                               |
	// | United States          | `us_ein`     | United States EIN                                                                                       |
	// | Uruguay                | `uy_ruc`     | Uruguayan RUC Number                                                                                    |
	// | Uzbekistan             | `uz_tin`     | Uzbekistan TIN Number                                                                                   |
	// | Uzbekistan             | `uz_vat`     | Uzbekistan VAT Number                                                                                   |
	// | Venezuela              | `ve_rif`     | Venezuelan RIF Number                                                                                   |
	// | Vietnam                | `vn_tin`     | Vietnamese Tax ID Number                                                                                |
	// | Zambia                 | `zm_tin`     | Zambia Tax Identification Number                                                                        |
	// | Zimbabwe               | `zw_tin`     | Zimbabwe Tax Identification Number                                                                      |
	CustomerTaxID CustomerTaxID `json:"customer_tax_id,required,nullable"`
	// This field is deprecated in favor of `discounts`. If a `discounts` list is
	// provided, the first discount in the list will be returned. If the list is empty,
	// `None` will be returned.
	//
	// Deprecated: deprecated
	Discount  interface{}            `json:"discount,required"`
	Discounts []InvoiceLevelDiscount `json:"discounts,required"`
	// When the invoice payment is due. The due date is null if the invoice is not yet
	// finalized.
	DueDate time.Time `json:"due_date,required,nullable" format:"date-time"`
	// If the invoice has a status of `draft`, this will be the time that the invoice
	// will be eligible to be issued, otherwise it will be `null`. If `auto-issue` is
	// true, the invoice will automatically begin issuing at this time.
	EligibleToIssueAt time.Time `json:"eligible_to_issue_at,required,nullable" format:"date-time"`
	// A URL for the customer-facing invoice portal. This URL expires 30 days after the
	// invoice's due date, or 60 days after being re-generated through the UI.
	HostedInvoiceURL string `json:"hosted_invoice_url,required,nullable"`
	// The scheduled date of the invoice
	InvoiceDate time.Time `json:"invoice_date,required" format:"date-time"`
	// Automatically generated invoice number to help track and reconcile invoices.
	// Invoice numbers have a prefix such as `RFOBWG`. These can be sequential per
	// account or customer.
	InvoiceNumber string `json:"invoice_number,required"`
	// The link to download the PDF representation of the `Invoice`.
	InvoicePdf    string               `json:"invoice_pdf,required,nullable"`
	InvoiceSource InvoiceInvoiceSource `json:"invoice_source,required"`
	// If the invoice failed to issue, this will be the last time it failed to issue
	// (even if it is now in a different state.)
	IssueFailedAt time.Time `json:"issue_failed_at,required,nullable" format:"date-time"`
	// If the invoice has been issued, this will be the time it transitioned to
	// `issued` (even if it is now in a different state.)
	IssuedAt time.Time `json:"issued_at,required,nullable" format:"date-time"`
	// The breakdown of prices in this invoice.
	LineItems     []InvoiceLineItem `json:"line_items,required"`
	Maximum       Maximum           `json:"maximum,required,nullable"`
	MaximumAmount string            `json:"maximum_amount,required,nullable"`
	// Free-form text which is available on the invoice PDF and the Orb invoice portal.
	Memo string `json:"memo,required,nullable"`
	// User specified key-value pairs for the resource. If not present, this defaults
	// to an empty dictionary. Individual keys can be removed by setting the value to
	// `null`, and the entire metadata mapping can be cleared by setting `metadata` to
	// `null`.
	Metadata      map[string]string `json:"metadata,required"`
	Minimum       Minimum           `json:"minimum,required,nullable"`
	MinimumAmount string            `json:"minimum_amount,required,nullable"`
	// If the invoice has a status of `paid`, this gives a timestamp when the invoice
	// was paid.
	PaidAt time.Time `json:"paid_at,required,nullable" format:"date-time"`
	// A list of payment attempts associated with the invoice
	PaymentAttempts []InvoicePaymentAttempt `json:"payment_attempts,required"`
	// If payment was attempted on this invoice but failed, this will be the time of
	// the most recent attempt.
	PaymentFailedAt time.Time `json:"payment_failed_at,required,nullable" format:"date-time"`
	// If payment was attempted on this invoice, this will be the start time of the
	// most recent attempt. This field is especially useful for delayed-notification
	// payment mechanisms (like bank transfers), where payment can take 3 days or more.
	PaymentStartedAt time.Time `json:"payment_started_at,required,nullable" format:"date-time"`
	// If the invoice is in draft, this timestamp will reflect when the invoice is
	// scheduled to be issued.
	ScheduledIssueAt time.Time            `json:"scheduled_issue_at,required,nullable" format:"date-time"`
	ShippingAddress  Address              `json:"shipping_address,required,nullable"`
	Status           InvoiceStatus        `json:"status,required"`
	Subscription     SubscriptionMinified `json:"subscription,required,nullable"`
	// The total before any discounts and minimums are applied.
	Subtotal string `json:"subtotal,required"`
	// If the invoice failed to sync, this will be the last time an external invoicing
	// provider sync was attempted. This field will always be `null` for invoices using
	// Orb Invoicing.
	SyncFailedAt time.Time `json:"sync_failed_at,required,nullable" format:"date-time"`
	// The total after any minimums and discounts have been applied.
	Total string `json:"total,required"`
	// If the invoice has a status of `void`, this gives a timestamp when the invoice
	// was voided.
	VoidedAt time.Time `json:"voided_at,required,nullable" format:"date-time"`
	// This is true if the invoice will be automatically issued in the future, and
	// false otherwise.
	WillAutoIssue bool        `json:"will_auto_issue,required"`
	JSON          invoiceJSON `json:"-"`
}

// invoiceJSON contains the JSON metadata for the struct [Invoice]
type invoiceJSON struct {
	ID                          apijson.Field
	AmountDue                   apijson.Field
	AutoCollection              apijson.Field
	BillingAddress              apijson.Field
	CreatedAt                   apijson.Field
	CreditNotes                 apijson.Field
	Currency                    apijson.Field
	Customer                    apijson.Field
	CustomerBalanceTransactions apijson.Field
	CustomerTaxID               apijson.Field
	Discount                    apijson.Field
	Discounts                   apijson.Field
	DueDate                     apijson.Field
	EligibleToIssueAt           apijson.Field
	HostedInvoiceURL            apijson.Field
	InvoiceDate                 apijson.Field
	InvoiceNumber               apijson.Field
	InvoicePdf                  apijson.Field
	InvoiceSource               apijson.Field
	IssueFailedAt               apijson.Field
	IssuedAt                    apijson.Field
	LineItems                   apijson.Field
	Maximum                     apijson.Field
	MaximumAmount               apijson.Field
	Memo                        apijson.Field
	Metadata                    apijson.Field
	Minimum                     apijson.Field
	MinimumAmount               apijson.Field
	PaidAt                      apijson.Field
	PaymentAttempts             apijson.Field
	PaymentFailedAt             apijson.Field
	PaymentStartedAt            apijson.Field
	ScheduledIssueAt            apijson.Field
	ShippingAddress             apijson.Field
	Status                      apijson.Field
	Subscription                apijson.Field
	Subtotal                    apijson.Field
	SyncFailedAt                apijson.Field
	Total                       apijson.Field
	VoidedAt                    apijson.Field
	WillAutoIssue               apijson.Field
	raw                         string
	ExtraFields                 map[string]apijson.Field
}

func (r *Invoice) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r invoiceJSON) RawJSON() string {
	return r.raw
}

type InvoiceAutoCollection struct {
	// True only if auto-collection is enabled for this invoice.
	Enabled bool `json:"enabled,required,nullable"`
	// If the invoice is scheduled for auto-collection, this field will reflect when
	// the next attempt will occur. If dunning has been exhausted, or auto-collection
	// is not enabled for this invoice, this field will be `null`.
	NextAttemptAt time.Time `json:"next_attempt_at,required,nullable" format:"date-time"`
	// Number of auto-collection payment attempts.
	NumAttempts int64 `json:"num_attempts,required,nullable"`
	// If Orb has ever attempted payment auto-collection for this invoice, this field
	// will reflect when that attempt occurred. In conjunction with `next_attempt_at`,
	// this can be used to tell whether the invoice is currently in dunning (that is,
	// `previously_attempted_at` is non-null, and `next_attempt_time` is non-null), or
	// if dunning has been exhausted (`previously_attempted_at` is non-null, but
	// `next_attempt_time` is null).
	PreviouslyAttemptedAt time.Time                 `json:"previously_attempted_at,required,nullable" format:"date-time"`
	JSON                  invoiceAutoCollectionJSON `json:"-"`
}

// invoiceAutoCollectionJSON contains the JSON metadata for the struct
// [InvoiceAutoCollection]
type invoiceAutoCollectionJSON struct {
	Enabled               apijson.Field
	NextAttemptAt         apijson.Field
	NumAttempts           apijson.Field
	PreviouslyAttemptedAt apijson.Field
	raw                   string
	ExtraFields           map[string]apijson.Field
}

func (r *InvoiceAutoCollection) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r invoiceAutoCollectionJSON) RawJSON() string {
	return r.raw
}

type InvoiceCreditNote struct {
	ID               string `json:"id,required"`
	CreditNoteNumber string `json:"credit_note_number,required"`
	// An optional memo supplied on the credit note.
	Memo   string `json:"memo,required,nullable"`
	Reason string `json:"reason,required"`
	Total  string `json:"total,required"`
	Type   string `json:"type,required"`
	// If the credit note has a status of `void`, this gives a timestamp when the
	// credit note was voided.
	VoidedAt time.Time             `json:"voided_at,required,nullable" format:"date-time"`
	JSON     invoiceCreditNoteJSON `json:"-"`
}

// invoiceCreditNoteJSON contains the JSON metadata for the struct
// [InvoiceCreditNote]
type invoiceCreditNoteJSON struct {
	ID               apijson.Field
	CreditNoteNumber apijson.Field
	Memo             apijson.Field
	Reason           apijson.Field
	Total            apijson.Field
	Type             apijson.Field
	VoidedAt         apijson.Field
	raw              string
	ExtraFields      map[string]apijson.Field
}

func (r *InvoiceCreditNote) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r invoiceCreditNoteJSON) RawJSON() string {
	return r.raw
}

type InvoiceCustomerBalanceTransaction struct {
	// A unique id for this transaction.
	ID     string                                   `json:"id,required"`
	Action InvoiceCustomerBalanceTransactionsAction `json:"action,required"`
	// The value of the amount changed in the transaction.
	Amount string `json:"amount,required"`
	// The creation time of this transaction.
	CreatedAt  time.Time      `json:"created_at,required" format:"date-time"`
	CreditNote CreditNoteTiny `json:"credit_note,required,nullable"`
	// An optional description provided for manual customer balance adjustments.
	Description string `json:"description,required,nullable"`
	// The new value of the customer's balance prior to the transaction, in the
	// customer's currency.
	EndingBalance string      `json:"ending_balance,required"`
	Invoice       InvoiceTiny `json:"invoice,required,nullable"`
	// The original value of the customer's balance prior to the transaction, in the
	// customer's currency.
	StartingBalance string                                 `json:"starting_balance,required"`
	Type            InvoiceCustomerBalanceTransactionsType `json:"type,required"`
	JSON            invoiceCustomerBalanceTransactionJSON  `json:"-"`
}

// invoiceCustomerBalanceTransactionJSON contains the JSON metadata for the struct
// [InvoiceCustomerBalanceTransaction]
type invoiceCustomerBalanceTransactionJSON struct {
	ID              apijson.Field
	Action          apijson.Field
	Amount          apijson.Field
	CreatedAt       apijson.Field
	CreditNote      apijson.Field
	Description     apijson.Field
	EndingBalance   apijson.Field
	Invoice         apijson.Field
	StartingBalance apijson.Field
	Type            apijson.Field
	raw             string
	ExtraFields     map[string]apijson.Field
}

func (r *InvoiceCustomerBalanceTransaction) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r invoiceCustomerBalanceTransactionJSON) RawJSON() string {
	return r.raw
}

type InvoiceCustomerBalanceTransactionsAction string

const (
	InvoiceCustomerBalanceTransactionsActionAppliedToInvoice     InvoiceCustomerBalanceTransactionsAction = "applied_to_invoice"
	InvoiceCustomerBalanceTransactionsActionManualAdjustment     InvoiceCustomerBalanceTransactionsAction = "manual_adjustment"
	InvoiceCustomerBalanceTransactionsActionProratedRefund       InvoiceCustomerBalanceTransactionsAction = "prorated_refund"
	InvoiceCustomerBalanceTransactionsActionRevertProratedRefund InvoiceCustomerBalanceTransactionsAction = "revert_prorated_refund"
	InvoiceCustomerBalanceTransactionsActionReturnFromVoiding    InvoiceCustomerBalanceTransactionsAction = "return_from_voiding"
	InvoiceCustomerBalanceTransactionsActionCreditNoteApplied    InvoiceCustomerBalanceTransactionsAction = "credit_note_applied"
	InvoiceCustomerBalanceTransactionsActionCreditNoteVoided     InvoiceCustomerBalanceTransactionsAction = "credit_note_voided"
	InvoiceCustomerBalanceTransactionsActionOverpaymentRefund    InvoiceCustomerBalanceTransactionsAction = "overpayment_refund"
	InvoiceCustomerBalanceTransactionsActionExternalPayment      InvoiceCustomerBalanceTransactionsAction = "external_payment"
)

func (r InvoiceCustomerBalanceTransactionsAction) IsKnown() bool {
	switch r {
	case InvoiceCustomerBalanceTransactionsActionAppliedToInvoice, InvoiceCustomerBalanceTransactionsActionManualAdjustment, InvoiceCustomerBalanceTransactionsActionProratedRefund, InvoiceCustomerBalanceTransactionsActionRevertProratedRefund, InvoiceCustomerBalanceTransactionsActionReturnFromVoiding, InvoiceCustomerBalanceTransactionsActionCreditNoteApplied, InvoiceCustomerBalanceTransactionsActionCreditNoteVoided, InvoiceCustomerBalanceTransactionsActionOverpaymentRefund, InvoiceCustomerBalanceTransactionsActionExternalPayment:
		return true
	}
	return false
}

type InvoiceCustomerBalanceTransactionsType string

const (
	InvoiceCustomerBalanceTransactionsTypeIncrement InvoiceCustomerBalanceTransactionsType = "increment"
	InvoiceCustomerBalanceTransactionsTypeDecrement InvoiceCustomerBalanceTransactionsType = "decrement"
)

func (r InvoiceCustomerBalanceTransactionsType) IsKnown() bool {
	switch r {
	case InvoiceCustomerBalanceTransactionsTypeIncrement, InvoiceCustomerBalanceTransactionsTypeDecrement:
		return true
	}
	return false
}

type InvoiceInvoiceSource string

const (
	InvoiceInvoiceSourceSubscription InvoiceInvoiceSource = "subscription"
	InvoiceInvoiceSourcePartial      InvoiceInvoiceSource = "partial"
	InvoiceInvoiceSourceOneOff       InvoiceInvoiceSource = "one_off"
)

func (r InvoiceInvoiceSource) IsKnown() bool {
	switch r {
	case InvoiceInvoiceSourceSubscription, InvoiceInvoiceSourcePartial, InvoiceInvoiceSourceOneOff:
		return true
	}
	return false
}

type InvoiceLineItem struct {
	// A unique ID for this line item.
	ID string `json:"id,required"`
	// The line amount after any adjustments and before overage conversion, credits and
	// partial invoicing.
	AdjustedSubtotal string `json:"adjusted_subtotal,required"`
	// All adjustments applied to the line item in the order they were applied based on
	// invoice calculations (ie. usage discounts -> amount discounts -> percentage
	// discounts -> minimums -> maximums).
	Adjustments []InvoiceLineItemsAdjustment `json:"adjustments,required"`
	// The final amount for a line item after all adjustments and pre paid credits have
	// been applied.
	Amount string `json:"amount,required"`
	// The number of prepaid credits applied.
	CreditsApplied string `json:"credits_applied,required"`
	// This field is deprecated in favor of `adjustments`
	//
	// Deprecated: deprecated
	Discount Discount `json:"discount,required,nullable"`
	// The end date of the range of time applied for this line item's price.
	EndDate time.Time `json:"end_date,required" format:"date-time"`
	// An additional filter that was used to calculate the usage for this line item.
	Filter string `json:"filter,required,nullable"`
	// [DEPRECATED] For configured prices that are split by a grouping key, this will
	// be populated with the key and a value. The `amount` and `subtotal` will be the
	// values for this particular grouping.
	Grouping string `json:"grouping,required,nullable"`
	// This field is deprecated in favor of `adjustments`.
	//
	// Deprecated: deprecated
	Maximum Maximum `json:"maximum,required,nullable"`
	// This field is deprecated in favor of `adjustments`.
	//
	// Deprecated: deprecated
	MaximumAmount string `json:"maximum_amount,required,nullable"`
	// This field is deprecated in favor of `adjustments`.
	//
	// Deprecated: deprecated
	Minimum Minimum `json:"minimum,required,nullable"`
	// This field is deprecated in favor of `adjustments`.
	//
	// Deprecated: deprecated
	MinimumAmount string `json:"minimum_amount,required,nullable"`
	// The name of the price associated with this line item.
	Name string `json:"name,required"`
	// Any amount applied from a partial invoice
	PartiallyInvoicedAmount string `json:"partially_invoiced_amount,required"`
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
	// Either the fixed fee quantity or the usage during the service period.
	Quantity float64 `json:"quantity,required"`
	// The start date of the range of time applied for this line item's price.
	StartDate time.Time `json:"start_date,required" format:"date-time"`
	// For complex pricing structures, the line item can be broken down further in
	// `sub_line_items`.
	SubLineItems []InvoiceLineItemsSubLineItem `json:"sub_line_items,required"`
	// The line amount before before any adjustments.
	Subtotal string `json:"subtotal,required"`
	// An array of tax rates and their incurred tax amounts. Empty if no tax
	// integration is configured.
	TaxAmounts []TaxAmount `json:"tax_amounts,required"`
	// A list of customer ids that were used to calculate the usage for this line item.
	UsageCustomerIDs []string            `json:"usage_customer_ids,required,nullable"`
	JSON             invoiceLineItemJSON `json:"-"`
}

// invoiceLineItemJSON contains the JSON metadata for the struct [InvoiceLineItem]
type invoiceLineItemJSON struct {
	ID                      apijson.Field
	AdjustedSubtotal        apijson.Field
	Adjustments             apijson.Field
	Amount                  apijson.Field
	CreditsApplied          apijson.Field
	Discount                apijson.Field
	EndDate                 apijson.Field
	Filter                  apijson.Field
	Grouping                apijson.Field
	Maximum                 apijson.Field
	MaximumAmount           apijson.Field
	Minimum                 apijson.Field
	MinimumAmount           apijson.Field
	Name                    apijson.Field
	PartiallyInvoicedAmount apijson.Field
	Price                   apijson.Field
	Quantity                apijson.Field
	StartDate               apijson.Field
	SubLineItems            apijson.Field
	Subtotal                apijson.Field
	TaxAmounts              apijson.Field
	UsageCustomerIDs        apijson.Field
	raw                     string
	ExtraFields             map[string]apijson.Field
}

func (r *InvoiceLineItem) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r invoiceLineItemJSON) RawJSON() string {
	return r.raw
}

type InvoiceLineItemsAdjustment struct {
	ID             string                                    `json:"id,required"`
	AdjustmentType InvoiceLineItemsAdjustmentsAdjustmentType `json:"adjustment_type,required"`
	// The value applied by an adjustment.
	Amount string `json:"amount,required"`
	// This field can have the runtime type of [[]string].
	AppliesToPriceIDs interface{} `json:"applies_to_price_ids,required"`
	// This field can have the runtime type of [[]TransformPriceFilter].
	Filters interface{} `json:"filters,required"`
	// True for adjustments that apply to an entire invocice, false for adjustments
	// that apply to only one price.
	IsInvoiceLevel bool `json:"is_invoice_level,required"`
	// The reason for the adjustment.
	Reason string `json:"reason,required,nullable"`
	// The adjustment id this adjustment replaces. This adjustment will take the place
	// of the replaced adjustment in plan version migrations.
	ReplacesAdjustmentID string `json:"replaces_adjustment_id,required,nullable"`
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
	UsageDiscount float64                        `json:"usage_discount"`
	JSON          invoiceLineItemsAdjustmentJSON `json:"-"`
	union         InvoiceLineItemsAdjustmentsUnion
}

// invoiceLineItemsAdjustmentJSON contains the JSON metadata for the struct
// [InvoiceLineItemsAdjustment]
type invoiceLineItemsAdjustmentJSON struct {
	ID                   apijson.Field
	AdjustmentType       apijson.Field
	Amount               apijson.Field
	AppliesToPriceIDs    apijson.Field
	Filters              apijson.Field
	IsInvoiceLevel       apijson.Field
	Reason               apijson.Field
	ReplacesAdjustmentID apijson.Field
	AmountDiscount       apijson.Field
	ItemID               apijson.Field
	MaximumAmount        apijson.Field
	MinimumAmount        apijson.Field
	PercentageDiscount   apijson.Field
	UsageDiscount        apijson.Field
	raw                  string
	ExtraFields          map[string]apijson.Field
}

func (r invoiceLineItemsAdjustmentJSON) RawJSON() string {
	return r.raw
}

func (r *InvoiceLineItemsAdjustment) UnmarshalJSON(data []byte) (err error) {
	*r = InvoiceLineItemsAdjustment{}
	err = apijson.UnmarshalRoot(data, &r.union)
	if err != nil {
		return err
	}
	return apijson.Port(r.union, &r)
}

// AsUnion returns a [InvoiceLineItemsAdjustmentsUnion] interface which you can
// cast to the specific types for more type safety.
//
// Possible runtime types of the union are [MonetaryUsageDiscountAdjustment],
// [MonetaryAmountDiscountAdjustment], [MonetaryPercentageDiscountAdjustment],
// [MonetaryMinimumAdjustment], [MonetaryMaximumAdjustment].
func (r InvoiceLineItemsAdjustment) AsUnion() InvoiceLineItemsAdjustmentsUnion {
	return r.union
}

// Union satisfied by [MonetaryUsageDiscountAdjustment],
// [MonetaryAmountDiscountAdjustment], [MonetaryPercentageDiscountAdjustment],
// [MonetaryMinimumAdjustment] or [MonetaryMaximumAdjustment].
type InvoiceLineItemsAdjustmentsUnion interface {
	ImplementsInvoiceLineItemsAdjustment()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*InvoiceLineItemsAdjustmentsUnion)(nil)).Elem(),
		"adjustment_type",
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(MonetaryUsageDiscountAdjustment{}),
			DiscriminatorValue: "usage_discount",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(MonetaryAmountDiscountAdjustment{}),
			DiscriminatorValue: "amount_discount",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(MonetaryPercentageDiscountAdjustment{}),
			DiscriminatorValue: "percentage_discount",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(MonetaryMinimumAdjustment{}),
			DiscriminatorValue: "minimum",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(MonetaryMaximumAdjustment{}),
			DiscriminatorValue: "maximum",
		},
	)
}

type InvoiceLineItemsAdjustmentsAdjustmentType string

const (
	InvoiceLineItemsAdjustmentsAdjustmentTypeUsageDiscount      InvoiceLineItemsAdjustmentsAdjustmentType = "usage_discount"
	InvoiceLineItemsAdjustmentsAdjustmentTypeAmountDiscount     InvoiceLineItemsAdjustmentsAdjustmentType = "amount_discount"
	InvoiceLineItemsAdjustmentsAdjustmentTypePercentageDiscount InvoiceLineItemsAdjustmentsAdjustmentType = "percentage_discount"
	InvoiceLineItemsAdjustmentsAdjustmentTypeMinimum            InvoiceLineItemsAdjustmentsAdjustmentType = "minimum"
	InvoiceLineItemsAdjustmentsAdjustmentTypeMaximum            InvoiceLineItemsAdjustmentsAdjustmentType = "maximum"
)

func (r InvoiceLineItemsAdjustmentsAdjustmentType) IsKnown() bool {
	switch r {
	case InvoiceLineItemsAdjustmentsAdjustmentTypeUsageDiscount, InvoiceLineItemsAdjustmentsAdjustmentTypeAmountDiscount, InvoiceLineItemsAdjustmentsAdjustmentTypePercentageDiscount, InvoiceLineItemsAdjustmentsAdjustmentTypeMinimum, InvoiceLineItemsAdjustmentsAdjustmentTypeMaximum:
		return true
	}
	return false
}

type InvoiceLineItemsSubLineItem struct {
	// The total amount for this sub line item.
	Amount       string                           `json:"amount,required"`
	Grouping     SubLineItemGrouping              `json:"grouping,required,nullable"`
	Name         string                           `json:"name,required"`
	Quantity     float64                          `json:"quantity,required"`
	Type         InvoiceLineItemsSubLineItemsType `json:"type,required"`
	MatrixConfig SubLineItemMatrixConfig          `json:"matrix_config"`
	TierConfig   TierConfig                       `json:"tier_config"`
	JSON         invoiceLineItemsSubLineItemJSON  `json:"-"`
	union        InvoiceLineItemsSubLineItemsUnion
}

// invoiceLineItemsSubLineItemJSON contains the JSON metadata for the struct
// [InvoiceLineItemsSubLineItem]
type invoiceLineItemsSubLineItemJSON struct {
	Amount       apijson.Field
	Grouping     apijson.Field
	Name         apijson.Field
	Quantity     apijson.Field
	Type         apijson.Field
	MatrixConfig apijson.Field
	TierConfig   apijson.Field
	raw          string
	ExtraFields  map[string]apijson.Field
}

func (r invoiceLineItemsSubLineItemJSON) RawJSON() string {
	return r.raw
}

func (r *InvoiceLineItemsSubLineItem) UnmarshalJSON(data []byte) (err error) {
	*r = InvoiceLineItemsSubLineItem{}
	err = apijson.UnmarshalRoot(data, &r.union)
	if err != nil {
		return err
	}
	return apijson.Port(r.union, &r)
}

// AsUnion returns a [InvoiceLineItemsSubLineItemsUnion] interface which you can
// cast to the specific types for more type safety.
//
// Possible runtime types of the union are [MatrixSubLineItem], [TierSubLineItem],
// [OtherSubLineItem].
func (r InvoiceLineItemsSubLineItem) AsUnion() InvoiceLineItemsSubLineItemsUnion {
	return r.union
}

// Union satisfied by [MatrixSubLineItem], [TierSubLineItem] or [OtherSubLineItem].
type InvoiceLineItemsSubLineItemsUnion interface {
	ImplementsInvoiceLineItemsSubLineItem()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*InvoiceLineItemsSubLineItemsUnion)(nil)).Elem(),
		"type",
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(MatrixSubLineItem{}),
			DiscriminatorValue: "matrix",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(TierSubLineItem{}),
			DiscriminatorValue: "tier",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(OtherSubLineItem{}),
			DiscriminatorValue: "'null'",
		},
	)
}

type InvoiceLineItemsSubLineItemsType string

const (
	InvoiceLineItemsSubLineItemsTypeMatrix InvoiceLineItemsSubLineItemsType = "matrix"
	InvoiceLineItemsSubLineItemsTypeTier   InvoiceLineItemsSubLineItemsType = "tier"
	InvoiceLineItemsSubLineItemsTypeNull   InvoiceLineItemsSubLineItemsType = "'null'"
)

func (r InvoiceLineItemsSubLineItemsType) IsKnown() bool {
	switch r {
	case InvoiceLineItemsSubLineItemsTypeMatrix, InvoiceLineItemsSubLineItemsTypeTier, InvoiceLineItemsSubLineItemsTypeNull:
		return true
	}
	return false
}

type InvoicePaymentAttempt struct {
	// The ID of the payment attempt.
	ID string `json:"id,required"`
	// The amount of the payment attempt.
	Amount string `json:"amount,required"`
	// The time at which the payment attempt was created.
	CreatedAt time.Time `json:"created_at,required" format:"date-time"`
	// The payment provider that attempted to collect the payment.
	PaymentProvider InvoicePaymentAttemptsPaymentProvider `json:"payment_provider,required,nullable"`
	// The ID of the payment attempt in the payment provider.
	PaymentProviderID string `json:"payment_provider_id,required,nullable"`
	// Whether the payment attempt succeeded.
	Succeeded bool                      `json:"succeeded,required"`
	JSON      invoicePaymentAttemptJSON `json:"-"`
}

// invoicePaymentAttemptJSON contains the JSON metadata for the struct
// [InvoicePaymentAttempt]
type invoicePaymentAttemptJSON struct {
	ID                apijson.Field
	Amount            apijson.Field
	CreatedAt         apijson.Field
	PaymentProvider   apijson.Field
	PaymentProviderID apijson.Field
	Succeeded         apijson.Field
	raw               string
	ExtraFields       map[string]apijson.Field
}

func (r *InvoicePaymentAttempt) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r invoicePaymentAttemptJSON) RawJSON() string {
	return r.raw
}

// The payment provider that attempted to collect the payment.
type InvoicePaymentAttemptsPaymentProvider string

const (
	InvoicePaymentAttemptsPaymentProviderStripe InvoicePaymentAttemptsPaymentProvider = "stripe"
)

func (r InvoicePaymentAttemptsPaymentProvider) IsKnown() bool {
	switch r {
	case InvoicePaymentAttemptsPaymentProviderStripe:
		return true
	}
	return false
}

type InvoiceStatus string

const (
	InvoiceStatusIssued InvoiceStatus = "issued"
	InvoiceStatusPaid   InvoiceStatus = "paid"
	InvoiceStatusSynced InvoiceStatus = "synced"
	InvoiceStatusVoid   InvoiceStatus = "void"
	InvoiceStatusDraft  InvoiceStatus = "draft"
)

func (r InvoiceStatus) IsKnown() bool {
	switch r {
	case InvoiceStatusIssued, InvoiceStatusPaid, InvoiceStatusSynced, InvoiceStatusVoid, InvoiceStatusDraft:
		return true
	}
	return false
}

type InvoiceLevelDiscount struct {
	DiscountType InvoiceLevelDiscountDiscountType `json:"discount_type,required"`
	// Only available if discount_type is `amount`.
	AmountDiscount string `json:"amount_discount"`
	// This field can have the runtime type of [[]string].
	AppliesToPriceIDs interface{} `json:"applies_to_price_ids"`
	// This field can have the runtime type of [[]TransformPriceFilter].
	Filters interface{} `json:"filters"`
	// Only available if discount_type is `percentage`. This is a number between 0
	// and 1.
	PercentageDiscount float64 `json:"percentage_discount"`
	Reason             string  `json:"reason,nullable"`
	// Only available if discount_type is `trial`
	TrialAmountDiscount string `json:"trial_amount_discount,nullable"`
	// Only available if discount_type is `trial`
	TrialPercentageDiscount float64                  `json:"trial_percentage_discount,nullable"`
	JSON                    invoiceLevelDiscountJSON `json:"-"`
	union                   InvoiceLevelDiscountUnion
}

// invoiceLevelDiscountJSON contains the JSON metadata for the struct
// [InvoiceLevelDiscount]
type invoiceLevelDiscountJSON struct {
	DiscountType            apijson.Field
	AmountDiscount          apijson.Field
	AppliesToPriceIDs       apijson.Field
	Filters                 apijson.Field
	PercentageDiscount      apijson.Field
	Reason                  apijson.Field
	TrialAmountDiscount     apijson.Field
	TrialPercentageDiscount apijson.Field
	raw                     string
	ExtraFields             map[string]apijson.Field
}

func (r invoiceLevelDiscountJSON) RawJSON() string {
	return r.raw
}

func (r *InvoiceLevelDiscount) UnmarshalJSON(data []byte) (err error) {
	*r = InvoiceLevelDiscount{}
	err = apijson.UnmarshalRoot(data, &r.union)
	if err != nil {
		return err
	}
	return apijson.Port(r.union, &r)
}

// AsUnion returns a [InvoiceLevelDiscountUnion] interface which you can cast to
// the specific types for more type safety.
//
// Possible runtime types of the union are [PercentageDiscount], [AmountDiscount],
// [TrialDiscount].
func (r InvoiceLevelDiscount) AsUnion() InvoiceLevelDiscountUnion {
	return r.union
}

// Union satisfied by [PercentageDiscount], [AmountDiscount] or [TrialDiscount].
type InvoiceLevelDiscountUnion interface {
	ImplementsInvoiceLevelDiscount()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*InvoiceLevelDiscountUnion)(nil)).Elem(),
		"discount_type",
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(PercentageDiscount{}),
			DiscriminatorValue: "percentage",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(AmountDiscount{}),
			DiscriminatorValue: "amount",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(TrialDiscount{}),
			DiscriminatorValue: "trial",
		},
	)
}

type InvoiceLevelDiscountDiscountType string

const (
	InvoiceLevelDiscountDiscountTypePercentage InvoiceLevelDiscountDiscountType = "percentage"
	InvoiceLevelDiscountDiscountTypeAmount     InvoiceLevelDiscountDiscountType = "amount"
	InvoiceLevelDiscountDiscountTypeTrial      InvoiceLevelDiscountDiscountType = "trial"
)

func (r InvoiceLevelDiscountDiscountType) IsKnown() bool {
	switch r {
	case InvoiceLevelDiscountDiscountTypePercentage, InvoiceLevelDiscountDiscountTypeAmount, InvoiceLevelDiscountDiscountTypeTrial:
		return true
	}
	return false
}

type InvoiceTiny struct {
	// The Invoice id
	ID   string          `json:"id,required"`
	JSON invoiceTinyJSON `json:"-"`
}

// invoiceTinyJSON contains the JSON metadata for the struct [InvoiceTiny]
type invoiceTinyJSON struct {
	ID          apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *InvoiceTiny) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r invoiceTinyJSON) RawJSON() string {
	return r.raw
}

type ItemSlim struct {
	ID   string       `json:"id,required"`
	Name string       `json:"name,required"`
	JSON itemSlimJSON `json:"-"`
}

// itemSlimJSON contains the JSON metadata for the struct [ItemSlim]
type itemSlimJSON struct {
	ID          apijson.Field
	Name        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ItemSlim) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r itemSlimJSON) RawJSON() string {
	return r.raw
}

type MatrixConfig struct {
	// Default per unit rate for any usage not bucketed into a specified matrix_value
	DefaultUnitAmount string `json:"default_unit_amount,required"`
	// One or two event property values to evaluate matrix groups by
	Dimensions []string `json:"dimensions,required"`
	// Matrix values for specified matrix grouping keys
	MatrixValues []MatrixValue    `json:"matrix_values,required"`
	JSON         matrixConfigJSON `json:"-"`
}

// matrixConfigJSON contains the JSON metadata for the struct [MatrixConfig]
type matrixConfigJSON struct {
	DefaultUnitAmount apijson.Field
	Dimensions        apijson.Field
	MatrixValues      apijson.Field
	raw               string
	ExtraFields       map[string]apijson.Field
}

func (r *MatrixConfig) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r matrixConfigJSON) RawJSON() string {
	return r.raw
}

type MatrixConfigParam struct {
	// Default per unit rate for any usage not bucketed into a specified matrix_value
	DefaultUnitAmount param.Field[string] `json:"default_unit_amount,required"`
	// One or two event property values to evaluate matrix groups by
	Dimensions param.Field[[]string] `json:"dimensions,required"`
	// Matrix values for specified matrix grouping keys
	MatrixValues param.Field[[]MatrixValueParam] `json:"matrix_values,required"`
}

func (r MatrixConfigParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type MatrixSubLineItem struct {
	// The total amount for this sub line item.
	Amount       string                  `json:"amount,required"`
	Grouping     SubLineItemGrouping     `json:"grouping,required,nullable"`
	MatrixConfig SubLineItemMatrixConfig `json:"matrix_config,required"`
	Name         string                  `json:"name,required"`
	Quantity     float64                 `json:"quantity,required"`
	Type         MatrixSubLineItemType   `json:"type,required"`
	JSON         matrixSubLineItemJSON   `json:"-"`
}

// matrixSubLineItemJSON contains the JSON metadata for the struct
// [MatrixSubLineItem]
type matrixSubLineItemJSON struct {
	Amount       apijson.Field
	Grouping     apijson.Field
	MatrixConfig apijson.Field
	Name         apijson.Field
	Quantity     apijson.Field
	Type         apijson.Field
	raw          string
	ExtraFields  map[string]apijson.Field
}

func (r *MatrixSubLineItem) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r matrixSubLineItemJSON) RawJSON() string {
	return r.raw
}

func (r MatrixSubLineItem) ImplementsInvoiceLineItemsSubLineItem() {}

func (r MatrixSubLineItem) ImplementsInvoiceLineItemNewResponseSubLineItem() {}

func (r MatrixSubLineItem) ImplementsInvoiceFetchUpcomingResponseLineItemsSubLineItem() {}

type MatrixSubLineItemType string

const (
	MatrixSubLineItemTypeMatrix MatrixSubLineItemType = "matrix"
)

func (r MatrixSubLineItemType) IsKnown() bool {
	switch r {
	case MatrixSubLineItemTypeMatrix:
		return true
	}
	return false
}

type MatrixValue struct {
	// One or two matrix keys to filter usage to this Matrix value by. For example,
	// ["region", "tier"] could be used to filter cloud usage by a cloud region and an
	// instance tier.
	DimensionValues []string `json:"dimension_values,required"`
	// Unit price for the specified dimension_values
	UnitAmount string          `json:"unit_amount,required"`
	JSON       matrixValueJSON `json:"-"`
}

// matrixValueJSON contains the JSON metadata for the struct [MatrixValue]
type matrixValueJSON struct {
	DimensionValues apijson.Field
	UnitAmount      apijson.Field
	raw             string
	ExtraFields     map[string]apijson.Field
}

func (r *MatrixValue) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r matrixValueJSON) RawJSON() string {
	return r.raw
}

type MatrixValueParam struct {
	// One or two matrix keys to filter usage to this Matrix value by. For example,
	// ["region", "tier"] could be used to filter cloud usage by a cloud region and an
	// instance tier.
	DimensionValues param.Field[[]string] `json:"dimension_values,required"`
	// Unit price for the specified dimension_values
	UnitAmount param.Field[string] `json:"unit_amount,required"`
}

func (r MatrixValueParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type MatrixWithAllocationConfig struct {
	// Allocation to be used to calculate the price
	Allocation float64 `json:"allocation,required"`
	// Default per unit rate for any usage not bucketed into a specified matrix_value
	DefaultUnitAmount string `json:"default_unit_amount,required"`
	// One or two event property values to evaluate matrix groups by
	Dimensions []string `json:"dimensions,required"`
	// Matrix values for specified matrix grouping keys
	MatrixValues []MatrixValue                  `json:"matrix_values,required"`
	JSON         matrixWithAllocationConfigJSON `json:"-"`
}

// matrixWithAllocationConfigJSON contains the JSON metadata for the struct
// [MatrixWithAllocationConfig]
type matrixWithAllocationConfigJSON struct {
	Allocation        apijson.Field
	DefaultUnitAmount apijson.Field
	Dimensions        apijson.Field
	MatrixValues      apijson.Field
	raw               string
	ExtraFields       map[string]apijson.Field
}

func (r *MatrixWithAllocationConfig) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r matrixWithAllocationConfigJSON) RawJSON() string {
	return r.raw
}

type MatrixWithAllocationConfigParam struct {
	// Allocation to be used to calculate the price
	Allocation param.Field[float64] `json:"allocation,required"`
	// Default per unit rate for any usage not bucketed into a specified matrix_value
	DefaultUnitAmount param.Field[string] `json:"default_unit_amount,required"`
	// One or two event property values to evaluate matrix groups by
	Dimensions param.Field[[]string] `json:"dimensions,required"`
	// Matrix values for specified matrix grouping keys
	MatrixValues param.Field[[]MatrixValueParam] `json:"matrix_values,required"`
}

func (r MatrixWithAllocationConfigParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type Maximum struct {
	// List of price_ids that this maximum amount applies to. For plan/plan phase
	// maximums, this can be a subset of prices.
	//
	// Deprecated: deprecated
	AppliesToPriceIDs []string `json:"applies_to_price_ids,required"`
	// The filters that determine which prices to apply this maximum to.
	Filters []TransformPriceFilter `json:"filters,required"`
	// Maximum amount applied
	MaximumAmount string      `json:"maximum_amount,required"`
	JSON          maximumJSON `json:"-"`
}

// maximumJSON contains the JSON metadata for the struct [Maximum]
type maximumJSON struct {
	AppliesToPriceIDs apijson.Field
	Filters           apijson.Field
	MaximumAmount     apijson.Field
	raw               string
	ExtraFields       map[string]apijson.Field
}

func (r *Maximum) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r maximumJSON) RawJSON() string {
	return r.raw
}

type MaximumInterval struct {
	// The price interval ids that this maximum interval applies to.
	AppliesToPriceIntervalIDs []string `json:"applies_to_price_interval_ids,required"`
	// The end date of the maximum interval.
	EndDate time.Time `json:"end_date,required,nullable" format:"date-time"`
	// The filters that determine which prices this maximum interval applies to.
	Filters []TransformPriceFilter `json:"filters,required"`
	// The maximum amount to charge in a given billing period for the price intervals
	// this transform applies to.
	MaximumAmount string `json:"maximum_amount,required"`
	// The start date of the maximum interval.
	StartDate time.Time           `json:"start_date,required" format:"date-time"`
	JSON      maximumIntervalJSON `json:"-"`
}

// maximumIntervalJSON contains the JSON metadata for the struct [MaximumInterval]
type maximumIntervalJSON struct {
	AppliesToPriceIntervalIDs apijson.Field
	EndDate                   apijson.Field
	Filters                   apijson.Field
	MaximumAmount             apijson.Field
	StartDate                 apijson.Field
	raw                       string
	ExtraFields               map[string]apijson.Field
}

func (r *MaximumInterval) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r maximumIntervalJSON) RawJSON() string {
	return r.raw
}

type Minimum struct {
	// List of price_ids that this minimum amount applies to. For plan/plan phase
	// minimums, this can be a subset of prices.
	//
	// Deprecated: deprecated
	AppliesToPriceIDs []string `json:"applies_to_price_ids,required"`
	// The filters that determine which prices to apply this minimum to.
	Filters []TransformPriceFilter `json:"filters,required"`
	// Minimum amount applied
	MinimumAmount string      `json:"minimum_amount,required"`
	JSON          minimumJSON `json:"-"`
}

// minimumJSON contains the JSON metadata for the struct [Minimum]
type minimumJSON struct {
	AppliesToPriceIDs apijson.Field
	Filters           apijson.Field
	MinimumAmount     apijson.Field
	raw               string
	ExtraFields       map[string]apijson.Field
}

func (r *Minimum) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r minimumJSON) RawJSON() string {
	return r.raw
}

type MinimumInterval struct {
	// The price interval ids that this minimum interval applies to.
	AppliesToPriceIntervalIDs []string `json:"applies_to_price_interval_ids,required"`
	// The end date of the minimum interval.
	EndDate time.Time `json:"end_date,required,nullable" format:"date-time"`
	// The filters that determine which prices this minimum interval applies to.
	Filters []TransformPriceFilter `json:"filters,required"`
	// The minimum amount to charge in a given billing period for the price intervals
	// this minimum applies to.
	MinimumAmount string `json:"minimum_amount,required"`
	// The start date of the minimum interval.
	StartDate time.Time           `json:"start_date,required" format:"date-time"`
	JSON      minimumIntervalJSON `json:"-"`
}

// minimumIntervalJSON contains the JSON metadata for the struct [MinimumInterval]
type minimumIntervalJSON struct {
	AppliesToPriceIntervalIDs apijson.Field
	EndDate                   apijson.Field
	Filters                   apijson.Field
	MinimumAmount             apijson.Field
	StartDate                 apijson.Field
	raw                       string
	ExtraFields               map[string]apijson.Field
}

func (r *MinimumInterval) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r minimumIntervalJSON) RawJSON() string {
	return r.raw
}

type MonetaryAmountDiscountAdjustment struct {
	ID             string                                         `json:"id,required"`
	AdjustmentType MonetaryAmountDiscountAdjustmentAdjustmentType `json:"adjustment_type,required"`
	// The value applied by an adjustment.
	Amount string `json:"amount,required"`
	// The amount by which to discount the prices this adjustment applies to in a given
	// billing period.
	AmountDiscount string `json:"amount_discount,required"`
	// The price IDs that this adjustment applies to.
	//
	// Deprecated: deprecated
	AppliesToPriceIDs []string `json:"applies_to_price_ids,required"`
	// The filters that determine which prices to apply this adjustment to.
	Filters []TransformPriceFilter `json:"filters,required"`
	// True for adjustments that apply to an entire invocice, false for adjustments
	// that apply to only one price.
	IsInvoiceLevel bool `json:"is_invoice_level,required"`
	// The reason for the adjustment.
	Reason string `json:"reason,required,nullable"`
	// The adjustment id this adjustment replaces. This adjustment will take the place
	// of the replaced adjustment in plan version migrations.
	ReplacesAdjustmentID string                               `json:"replaces_adjustment_id,required,nullable"`
	JSON                 monetaryAmountDiscountAdjustmentJSON `json:"-"`
}

// monetaryAmountDiscountAdjustmentJSON contains the JSON metadata for the struct
// [MonetaryAmountDiscountAdjustment]
type monetaryAmountDiscountAdjustmentJSON struct {
	ID                   apijson.Field
	AdjustmentType       apijson.Field
	Amount               apijson.Field
	AmountDiscount       apijson.Field
	AppliesToPriceIDs    apijson.Field
	Filters              apijson.Field
	IsInvoiceLevel       apijson.Field
	Reason               apijson.Field
	ReplacesAdjustmentID apijson.Field
	raw                  string
	ExtraFields          map[string]apijson.Field
}

func (r *MonetaryAmountDiscountAdjustment) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r monetaryAmountDiscountAdjustmentJSON) RawJSON() string {
	return r.raw
}

func (r MonetaryAmountDiscountAdjustment) ImplementsInvoiceLineItemsAdjustment() {}

func (r MonetaryAmountDiscountAdjustment) ImplementsInvoiceLineItemNewResponseAdjustment() {}

func (r MonetaryAmountDiscountAdjustment) ImplementsInvoiceFetchUpcomingResponseLineItemsAdjustment() {
}

type MonetaryAmountDiscountAdjustmentAdjustmentType string

const (
	MonetaryAmountDiscountAdjustmentAdjustmentTypeAmountDiscount MonetaryAmountDiscountAdjustmentAdjustmentType = "amount_discount"
)

func (r MonetaryAmountDiscountAdjustmentAdjustmentType) IsKnown() bool {
	switch r {
	case MonetaryAmountDiscountAdjustmentAdjustmentTypeAmountDiscount:
		return true
	}
	return false
}

type MonetaryMaximumAdjustment struct {
	ID             string                                  `json:"id,required"`
	AdjustmentType MonetaryMaximumAdjustmentAdjustmentType `json:"adjustment_type,required"`
	// The value applied by an adjustment.
	Amount string `json:"amount,required"`
	// The price IDs that this adjustment applies to.
	//
	// Deprecated: deprecated
	AppliesToPriceIDs []string `json:"applies_to_price_ids,required"`
	// The filters that determine which prices to apply this adjustment to.
	Filters []TransformPriceFilter `json:"filters,required"`
	// True for adjustments that apply to an entire invocice, false for adjustments
	// that apply to only one price.
	IsInvoiceLevel bool `json:"is_invoice_level,required"`
	// The maximum amount to charge in a given billing period for the prices this
	// adjustment applies to.
	MaximumAmount string `json:"maximum_amount,required"`
	// The reason for the adjustment.
	Reason string `json:"reason,required,nullable"`
	// The adjustment id this adjustment replaces. This adjustment will take the place
	// of the replaced adjustment in plan version migrations.
	ReplacesAdjustmentID string                        `json:"replaces_adjustment_id,required,nullable"`
	JSON                 monetaryMaximumAdjustmentJSON `json:"-"`
}

// monetaryMaximumAdjustmentJSON contains the JSON metadata for the struct
// [MonetaryMaximumAdjustment]
type monetaryMaximumAdjustmentJSON struct {
	ID                   apijson.Field
	AdjustmentType       apijson.Field
	Amount               apijson.Field
	AppliesToPriceIDs    apijson.Field
	Filters              apijson.Field
	IsInvoiceLevel       apijson.Field
	MaximumAmount        apijson.Field
	Reason               apijson.Field
	ReplacesAdjustmentID apijson.Field
	raw                  string
	ExtraFields          map[string]apijson.Field
}

func (r *MonetaryMaximumAdjustment) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r monetaryMaximumAdjustmentJSON) RawJSON() string {
	return r.raw
}

func (r MonetaryMaximumAdjustment) ImplementsInvoiceLineItemsAdjustment() {}

func (r MonetaryMaximumAdjustment) ImplementsInvoiceLineItemNewResponseAdjustment() {}

func (r MonetaryMaximumAdjustment) ImplementsInvoiceFetchUpcomingResponseLineItemsAdjustment() {}

type MonetaryMaximumAdjustmentAdjustmentType string

const (
	MonetaryMaximumAdjustmentAdjustmentTypeMaximum MonetaryMaximumAdjustmentAdjustmentType = "maximum"
)

func (r MonetaryMaximumAdjustmentAdjustmentType) IsKnown() bool {
	switch r {
	case MonetaryMaximumAdjustmentAdjustmentTypeMaximum:
		return true
	}
	return false
}

type MonetaryMinimumAdjustment struct {
	ID             string                                  `json:"id,required"`
	AdjustmentType MonetaryMinimumAdjustmentAdjustmentType `json:"adjustment_type,required"`
	// The value applied by an adjustment.
	Amount string `json:"amount,required"`
	// The price IDs that this adjustment applies to.
	//
	// Deprecated: deprecated
	AppliesToPriceIDs []string `json:"applies_to_price_ids,required"`
	// The filters that determine which prices to apply this adjustment to.
	Filters []TransformPriceFilter `json:"filters,required"`
	// True for adjustments that apply to an entire invocice, false for adjustments
	// that apply to only one price.
	IsInvoiceLevel bool `json:"is_invoice_level,required"`
	// The item ID that revenue from this minimum will be attributed to.
	ItemID string `json:"item_id,required"`
	// The minimum amount to charge in a given billing period for the prices this
	// adjustment applies to.
	MinimumAmount string `json:"minimum_amount,required"`
	// The reason for the adjustment.
	Reason string `json:"reason,required,nullable"`
	// The adjustment id this adjustment replaces. This adjustment will take the place
	// of the replaced adjustment in plan version migrations.
	ReplacesAdjustmentID string                        `json:"replaces_adjustment_id,required,nullable"`
	JSON                 monetaryMinimumAdjustmentJSON `json:"-"`
}

// monetaryMinimumAdjustmentJSON contains the JSON metadata for the struct
// [MonetaryMinimumAdjustment]
type monetaryMinimumAdjustmentJSON struct {
	ID                   apijson.Field
	AdjustmentType       apijson.Field
	Amount               apijson.Field
	AppliesToPriceIDs    apijson.Field
	Filters              apijson.Field
	IsInvoiceLevel       apijson.Field
	ItemID               apijson.Field
	MinimumAmount        apijson.Field
	Reason               apijson.Field
	ReplacesAdjustmentID apijson.Field
	raw                  string
	ExtraFields          map[string]apijson.Field
}

func (r *MonetaryMinimumAdjustment) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r monetaryMinimumAdjustmentJSON) RawJSON() string {
	return r.raw
}

func (r MonetaryMinimumAdjustment) ImplementsInvoiceLineItemsAdjustment() {}

func (r MonetaryMinimumAdjustment) ImplementsInvoiceLineItemNewResponseAdjustment() {}

func (r MonetaryMinimumAdjustment) ImplementsInvoiceFetchUpcomingResponseLineItemsAdjustment() {}

type MonetaryMinimumAdjustmentAdjustmentType string

const (
	MonetaryMinimumAdjustmentAdjustmentTypeMinimum MonetaryMinimumAdjustmentAdjustmentType = "minimum"
)

func (r MonetaryMinimumAdjustmentAdjustmentType) IsKnown() bool {
	switch r {
	case MonetaryMinimumAdjustmentAdjustmentTypeMinimum:
		return true
	}
	return false
}

type MonetaryPercentageDiscountAdjustment struct {
	ID             string                                             `json:"id,required"`
	AdjustmentType MonetaryPercentageDiscountAdjustmentAdjustmentType `json:"adjustment_type,required"`
	// The value applied by an adjustment.
	Amount string `json:"amount,required"`
	// The price IDs that this adjustment applies to.
	//
	// Deprecated: deprecated
	AppliesToPriceIDs []string `json:"applies_to_price_ids,required"`
	// The filters that determine which prices to apply this adjustment to.
	Filters []TransformPriceFilter `json:"filters,required"`
	// True for adjustments that apply to an entire invocice, false for adjustments
	// that apply to only one price.
	IsInvoiceLevel bool `json:"is_invoice_level,required"`
	// The percentage (as a value between 0 and 1) by which to discount the price
	// intervals this adjustment applies to in a given billing period.
	PercentageDiscount float64 `json:"percentage_discount,required"`
	// The reason for the adjustment.
	Reason string `json:"reason,required,nullable"`
	// The adjustment id this adjustment replaces. This adjustment will take the place
	// of the replaced adjustment in plan version migrations.
	ReplacesAdjustmentID string                                   `json:"replaces_adjustment_id,required,nullable"`
	JSON                 monetaryPercentageDiscountAdjustmentJSON `json:"-"`
}

// monetaryPercentageDiscountAdjustmentJSON contains the JSON metadata for the
// struct [MonetaryPercentageDiscountAdjustment]
type monetaryPercentageDiscountAdjustmentJSON struct {
	ID                   apijson.Field
	AdjustmentType       apijson.Field
	Amount               apijson.Field
	AppliesToPriceIDs    apijson.Field
	Filters              apijson.Field
	IsInvoiceLevel       apijson.Field
	PercentageDiscount   apijson.Field
	Reason               apijson.Field
	ReplacesAdjustmentID apijson.Field
	raw                  string
	ExtraFields          map[string]apijson.Field
}

func (r *MonetaryPercentageDiscountAdjustment) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r monetaryPercentageDiscountAdjustmentJSON) RawJSON() string {
	return r.raw
}

func (r MonetaryPercentageDiscountAdjustment) ImplementsInvoiceLineItemsAdjustment() {}

func (r MonetaryPercentageDiscountAdjustment) ImplementsInvoiceLineItemNewResponseAdjustment() {}

func (r MonetaryPercentageDiscountAdjustment) ImplementsInvoiceFetchUpcomingResponseLineItemsAdjustment() {
}

type MonetaryPercentageDiscountAdjustmentAdjustmentType string

const (
	MonetaryPercentageDiscountAdjustmentAdjustmentTypePercentageDiscount MonetaryPercentageDiscountAdjustmentAdjustmentType = "percentage_discount"
)

func (r MonetaryPercentageDiscountAdjustmentAdjustmentType) IsKnown() bool {
	switch r {
	case MonetaryPercentageDiscountAdjustmentAdjustmentTypePercentageDiscount:
		return true
	}
	return false
}

type MonetaryUsageDiscountAdjustment struct {
	ID             string                                        `json:"id,required"`
	AdjustmentType MonetaryUsageDiscountAdjustmentAdjustmentType `json:"adjustment_type,required"`
	// The value applied by an adjustment.
	Amount string `json:"amount,required"`
	// The price IDs that this adjustment applies to.
	//
	// Deprecated: deprecated
	AppliesToPriceIDs []string `json:"applies_to_price_ids,required"`
	// The filters that determine which prices to apply this adjustment to.
	Filters []TransformPriceFilter `json:"filters,required"`
	// True for adjustments that apply to an entire invocice, false for adjustments
	// that apply to only one price.
	IsInvoiceLevel bool `json:"is_invoice_level,required"`
	// The reason for the adjustment.
	Reason string `json:"reason,required,nullable"`
	// The adjustment id this adjustment replaces. This adjustment will take the place
	// of the replaced adjustment in plan version migrations.
	ReplacesAdjustmentID string `json:"replaces_adjustment_id,required,nullable"`
	// The number of usage units by which to discount the price this adjustment applies
	// to in a given billing period.
	UsageDiscount float64                             `json:"usage_discount,required"`
	JSON          monetaryUsageDiscountAdjustmentJSON `json:"-"`
}

// monetaryUsageDiscountAdjustmentJSON contains the JSON metadata for the struct
// [MonetaryUsageDiscountAdjustment]
type monetaryUsageDiscountAdjustmentJSON struct {
	ID                   apijson.Field
	AdjustmentType       apijson.Field
	Amount               apijson.Field
	AppliesToPriceIDs    apijson.Field
	Filters              apijson.Field
	IsInvoiceLevel       apijson.Field
	Reason               apijson.Field
	ReplacesAdjustmentID apijson.Field
	UsageDiscount        apijson.Field
	raw                  string
	ExtraFields          map[string]apijson.Field
}

func (r *MonetaryUsageDiscountAdjustment) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r monetaryUsageDiscountAdjustmentJSON) RawJSON() string {
	return r.raw
}

func (r MonetaryUsageDiscountAdjustment) ImplementsInvoiceLineItemsAdjustment() {}

func (r MonetaryUsageDiscountAdjustment) ImplementsInvoiceLineItemNewResponseAdjustment() {}

func (r MonetaryUsageDiscountAdjustment) ImplementsInvoiceFetchUpcomingResponseLineItemsAdjustment() {
}

type MonetaryUsageDiscountAdjustmentAdjustmentType string

const (
	MonetaryUsageDiscountAdjustmentAdjustmentTypeUsageDiscount MonetaryUsageDiscountAdjustmentAdjustmentType = "usage_discount"
)

func (r MonetaryUsageDiscountAdjustmentAdjustmentType) IsKnown() bool {
	switch r {
	case MonetaryUsageDiscountAdjustmentAdjustmentTypeUsageDiscount:
		return true
	}
	return false
}

type NewAllocationPriceParam struct {
	// An amount of the currency to allocate to the customer at the specified cadence.
	Amount param.Field[string] `json:"amount,required"`
	// The cadence at which to allocate the amount to the customer.
	Cadence param.Field[NewAllocationPriceCadence] `json:"cadence,required"`
	// An ISO 4217 currency string or a custom pricing unit identifier in which to bill
	// this price.
	Currency param.Field[string] `json:"currency,required"`
	// The custom expiration for the allocation.
	CustomExpiration param.Field[CustomExpirationParam] `json:"custom_expiration"`
	// Whether the allocated amount should expire at the end of the cadence or roll
	// over to the next period. Set to null if using custom_expiration.
	ExpiresAtEndOfCadence param.Field[bool] `json:"expires_at_end_of_cadence"`
}

func (r NewAllocationPriceParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// The cadence at which to allocate the amount to the customer.
type NewAllocationPriceCadence string

const (
	NewAllocationPriceCadenceOneTime    NewAllocationPriceCadence = "one_time"
	NewAllocationPriceCadenceMonthly    NewAllocationPriceCadence = "monthly"
	NewAllocationPriceCadenceQuarterly  NewAllocationPriceCadence = "quarterly"
	NewAllocationPriceCadenceSemiAnnual NewAllocationPriceCadence = "semi_annual"
	NewAllocationPriceCadenceAnnual     NewAllocationPriceCadence = "annual"
)

func (r NewAllocationPriceCadence) IsKnown() bool {
	switch r {
	case NewAllocationPriceCadenceOneTime, NewAllocationPriceCadenceMonthly, NewAllocationPriceCadenceQuarterly, NewAllocationPriceCadenceSemiAnnual, NewAllocationPriceCadenceAnnual:
		return true
	}
	return false
}

type NewAmountDiscountParam struct {
	AdjustmentType param.Field[NewAmountDiscountAdjustmentType] `json:"adjustment_type,required"`
	AmountDiscount param.Field[string]                          `json:"amount_discount,required"`
	// If set, the adjustment will apply to every price on the subscription.
	AppliesToAll param.Field[NewAmountDiscountAppliesToAll] `json:"applies_to_all"`
	// The set of item IDs to which this adjustment applies.
	AppliesToItemIDs param.Field[[]string] `json:"applies_to_item_ids"`
	// The set of price IDs to which this adjustment applies.
	AppliesToPriceIDs param.Field[[]string] `json:"applies_to_price_ids"`
	// If set, only prices in the specified currency will have the adjustment applied.
	Currency param.Field[string] `json:"currency"`
	// A list of filters that determine which prices this adjustment will apply to.
	Filters param.Field[[]TransformPriceFilterParam] `json:"filters"`
	// When false, this adjustment will be applied to a single price. Otherwise, it
	// will be applied at the invoice level, possibly to multiple prices.
	IsInvoiceLevel param.Field[bool] `json:"is_invoice_level"`
	// If set, only prices of the specified type will have the adjustment applied.
	PriceType param.Field[NewAmountDiscountPriceType] `json:"price_type"`
}

func (r NewAmountDiscountParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r NewAmountDiscountParam) ImplementsBetaNewPlanVersionParamsAddAdjustmentsAdjustmentUnion() {}

func (r NewAmountDiscountParam) ImplementsBetaNewPlanVersionParamsReplaceAdjustmentsAdjustmentUnion() {
}

func (r NewAmountDiscountParam) ImplementsBetaExternalPlanIDNewPlanVersionParamsAddAdjustmentsAdjustmentUnion() {
}

func (r NewAmountDiscountParam) ImplementsBetaExternalPlanIDNewPlanVersionParamsReplaceAdjustmentsAdjustmentUnion() {
}

func (r NewAmountDiscountParam) ImplementsPlanNewParamsAdjustmentsAdjustmentUnion() {}

func (r NewAmountDiscountParam) ImplementsSubscriptionNewParamsAddAdjustmentsAdjustmentUnion() {}

func (r NewAmountDiscountParam) ImplementsSubscriptionNewParamsReplaceAdjustmentsAdjustmentUnion() {}

func (r NewAmountDiscountParam) ImplementsSubscriptionPriceIntervalsParamsAddAdjustmentsAdjustmentUnion() {
}

func (r NewAmountDiscountParam) ImplementsSubscriptionSchedulePlanChangeParamsAddAdjustmentsAdjustmentUnion() {
}

func (r NewAmountDiscountParam) ImplementsSubscriptionSchedulePlanChangeParamsReplaceAdjustmentsAdjustmentUnion() {
}

type NewAmountDiscountAdjustmentType string

const (
	NewAmountDiscountAdjustmentTypeAmountDiscount NewAmountDiscountAdjustmentType = "amount_discount"
)

func (r NewAmountDiscountAdjustmentType) IsKnown() bool {
	switch r {
	case NewAmountDiscountAdjustmentTypeAmountDiscount:
		return true
	}
	return false
}

// If set, the adjustment will apply to every price on the subscription.
type NewAmountDiscountAppliesToAll bool

const (
	NewAmountDiscountAppliesToAllTrue NewAmountDiscountAppliesToAll = true
)

func (r NewAmountDiscountAppliesToAll) IsKnown() bool {
	switch r {
	case NewAmountDiscountAppliesToAllTrue:
		return true
	}
	return false
}

// If set, only prices of the specified type will have the adjustment applied.
type NewAmountDiscountPriceType string

const (
	NewAmountDiscountPriceTypeUsage          NewAmountDiscountPriceType = "usage"
	NewAmountDiscountPriceTypeFixedInAdvance NewAmountDiscountPriceType = "fixed_in_advance"
	NewAmountDiscountPriceTypeFixedInArrears NewAmountDiscountPriceType = "fixed_in_arrears"
	NewAmountDiscountPriceTypeFixed          NewAmountDiscountPriceType = "fixed"
	NewAmountDiscountPriceTypeInArrears      NewAmountDiscountPriceType = "in_arrears"
)

func (r NewAmountDiscountPriceType) IsKnown() bool {
	switch r {
	case NewAmountDiscountPriceTypeUsage, NewAmountDiscountPriceTypeFixedInAdvance, NewAmountDiscountPriceTypeFixedInArrears, NewAmountDiscountPriceTypeFixed, NewAmountDiscountPriceTypeInArrears:
		return true
	}
	return false
}

type NewBillingCycleConfigurationParam struct {
	// The duration of the billing period.
	Duration param.Field[int64] `json:"duration,required"`
	// The unit of billing period duration.
	DurationUnit param.Field[NewBillingCycleConfigurationDurationUnit] `json:"duration_unit,required"`
}

func (r NewBillingCycleConfigurationParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// The unit of billing period duration.
type NewBillingCycleConfigurationDurationUnit string

const (
	NewBillingCycleConfigurationDurationUnitDay   NewBillingCycleConfigurationDurationUnit = "day"
	NewBillingCycleConfigurationDurationUnitMonth NewBillingCycleConfigurationDurationUnit = "month"
)

func (r NewBillingCycleConfigurationDurationUnit) IsKnown() bool {
	switch r {
	case NewBillingCycleConfigurationDurationUnitDay, NewBillingCycleConfigurationDurationUnitMonth:
		return true
	}
	return false
}

type NewDimensionalPriceConfigurationParam struct {
	// The list of dimension values matching (in order) the dimensions of the price
	// group
	DimensionValues param.Field[[]string] `json:"dimension_values,required"`
	// The id of the dimensional price group to include this price in
	DimensionalPriceGroupID param.Field[string] `json:"dimensional_price_group_id"`
	// The external id of the dimensional price group to include this price in
	ExternalDimensionalPriceGroupID param.Field[string] `json:"external_dimensional_price_group_id"`
}

func (r NewDimensionalPriceConfigurationParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type NewFloatingBPSPriceParam struct {
	BPSConfig param.Field[BPSConfigParam] `json:"bps_config,required"`
	// The cadence to bill for this price on.
	Cadence param.Field[NewFloatingBPSPriceCadence] `json:"cadence,required"`
	// An ISO 4217 currency string for which this price is billed in.
	Currency param.Field[string] `json:"currency,required"`
	// The id of the item the price will be associated with.
	ItemID    param.Field[string]                       `json:"item_id,required"`
	ModelType param.Field[NewFloatingBPSPriceModelType] `json:"model_type,required"`
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
	BillingCycleConfiguration param.Field[NewBillingCycleConfigurationParam] `json:"billing_cycle_configuration"`
	// The per unit conversion rate of the price currency to the invoicing currency.
	ConversionRate param.Field[float64] `json:"conversion_rate"`
	// The configuration for the rate of the price currency to the invoicing currency.
	ConversionRateConfig param.Field[NewFloatingBPSPriceUnionParam] `json:"conversion_rate_config"`
	// For dimensional price: specifies a price group and dimension values
	DimensionalPriceConfiguration param.Field[NewDimensionalPriceConfigurationParam] `json:"dimensional_price_configuration"`
	// An alias for the price.
	ExternalPriceID param.Field[string] `json:"external_price_id"`
	// If the Price represents a fixed cost, this represents the quantity of units
	// applied.
	FixedPriceQuantity param.Field[float64] `json:"fixed_price_quantity"`
	// The property used to group this price on an invoice
	InvoiceGroupingKey param.Field[string] `json:"invoice_grouping_key"`
	// Within each billing cycle, specifies the cadence at which invoices are produced.
	// If unspecified, a single invoice is produced per billing cycle.
	InvoicingCycleConfiguration param.Field[NewBillingCycleConfigurationParam] `json:"invoicing_cycle_configuration"`
	// User-specified key/value pairs for the resource. Individual keys can be removed
	// by setting the value to `null`, and the entire metadata mapping can be cleared
	// by setting `metadata` to `null`.
	Metadata param.Field[map[string]string] `json:"metadata"`
}

func (r NewFloatingBPSPriceParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r NewFloatingBPSPriceParam) ImplementsPriceEvaluateMultipleParamsPriceEvaluationsPriceUnion() {}

func (r NewFloatingBPSPriceParam) ImplementsPriceEvaluatePreviewEventsParamsPriceEvaluationsPriceUnion() {
}

func (r NewFloatingBPSPriceParam) ImplementsSubscriptionPriceIntervalsParamsAddPriceUnion() {}

// The cadence to bill for this price on.
type NewFloatingBPSPriceCadence string

const (
	NewFloatingBPSPriceCadenceAnnual     NewFloatingBPSPriceCadence = "annual"
	NewFloatingBPSPriceCadenceSemiAnnual NewFloatingBPSPriceCadence = "semi_annual"
	NewFloatingBPSPriceCadenceMonthly    NewFloatingBPSPriceCadence = "monthly"
	NewFloatingBPSPriceCadenceQuarterly  NewFloatingBPSPriceCadence = "quarterly"
	NewFloatingBPSPriceCadenceOneTime    NewFloatingBPSPriceCadence = "one_time"
	NewFloatingBPSPriceCadenceCustom     NewFloatingBPSPriceCadence = "custom"
)

func (r NewFloatingBPSPriceCadence) IsKnown() bool {
	switch r {
	case NewFloatingBPSPriceCadenceAnnual, NewFloatingBPSPriceCadenceSemiAnnual, NewFloatingBPSPriceCadenceMonthly, NewFloatingBPSPriceCadenceQuarterly, NewFloatingBPSPriceCadenceOneTime, NewFloatingBPSPriceCadenceCustom:
		return true
	}
	return false
}

type NewFloatingBPSPriceModelType string

const (
	NewFloatingBPSPriceModelTypeBPS NewFloatingBPSPriceModelType = "bps"
)

func (r NewFloatingBPSPriceModelType) IsKnown() bool {
	switch r {
	case NewFloatingBPSPriceModelTypeBPS:
		return true
	}
	return false
}

type NewFloatingBPSPriceParam struct {
	ConversionRateType param.Field[NewFloatingBPSPriceConversionRateType] `json:"conversion_rate_type,required"`
	TieredConfig       param.Field[ConversionRateTieredConfigParam]       `json:"tiered_config"`
	UnitConfig         param.Field[ConversionRateUnitConfigParam]         `json:"unit_config"`
}

func (r NewFloatingBPSPriceParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r NewFloatingBPSPriceParam) ImplementsNewFloatingBPSPriceUnionParam() {}

// Satisfied by [shared.UnitConversionRateConfigParam],
// [shared.TieredConversionRateConfigParam], [NewFloatingBPSPriceParam].
type NewFloatingBPSPriceUnionParam interface {
	ImplementsNewFloatingBPSPriceUnionParam()
}

type NewFloatingBPSPriceConversionRateType string

const (
	NewFloatingBPSPriceConversionRateTypeUnit   NewFloatingBPSPriceConversionRateType = "unit"
	NewFloatingBPSPriceConversionRateTypeTiered NewFloatingBPSPriceConversionRateType = "tiered"
)

func (r NewFloatingBPSPriceConversionRateType) IsKnown() bool {
	switch r {
	case NewFloatingBPSPriceConversionRateTypeUnit, NewFloatingBPSPriceConversionRateTypeTiered:
		return true
	}
	return false
}

type NewFloatingBulkBPSPriceParam struct {
	BulkBPSConfig param.Field[BulkBPSConfigParam] `json:"bulk_bps_config,required"`
	// The cadence to bill for this price on.
	Cadence param.Field[NewFloatingBulkBPSPriceCadence] `json:"cadence,required"`
	// An ISO 4217 currency string for which this price is billed in.
	Currency param.Field[string] `json:"currency,required"`
	// The id of the item the price will be associated with.
	ItemID    param.Field[string]                           `json:"item_id,required"`
	ModelType param.Field[NewFloatingBulkBPSPriceModelType] `json:"model_type,required"`
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
	BillingCycleConfiguration param.Field[NewBillingCycleConfigurationParam] `json:"billing_cycle_configuration"`
	// The per unit conversion rate of the price currency to the invoicing currency.
	ConversionRate param.Field[float64] `json:"conversion_rate"`
	// The configuration for the rate of the price currency to the invoicing currency.
	ConversionRateConfig param.Field[NewFloatingBulkBPSPriceUnionParam] `json:"conversion_rate_config"`
	// For dimensional price: specifies a price group and dimension values
	DimensionalPriceConfiguration param.Field[NewDimensionalPriceConfigurationParam] `json:"dimensional_price_configuration"`
	// An alias for the price.
	ExternalPriceID param.Field[string] `json:"external_price_id"`
	// If the Price represents a fixed cost, this represents the quantity of units
	// applied.
	FixedPriceQuantity param.Field[float64] `json:"fixed_price_quantity"`
	// The property used to group this price on an invoice
	InvoiceGroupingKey param.Field[string] `json:"invoice_grouping_key"`
	// Within each billing cycle, specifies the cadence at which invoices are produced.
	// If unspecified, a single invoice is produced per billing cycle.
	InvoicingCycleConfiguration param.Field[NewBillingCycleConfigurationParam] `json:"invoicing_cycle_configuration"`
	// User-specified key/value pairs for the resource. Individual keys can be removed
	// by setting the value to `null`, and the entire metadata mapping can be cleared
	// by setting `metadata` to `null`.
	Metadata param.Field[map[string]string] `json:"metadata"`
}

func (r NewFloatingBulkBPSPriceParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r NewFloatingBulkBPSPriceParam) ImplementsPriceEvaluateMultipleParamsPriceEvaluationsPriceUnion() {
}

func (r NewFloatingBulkBPSPriceParam) ImplementsPriceEvaluatePreviewEventsParamsPriceEvaluationsPriceUnion() {
}

func (r NewFloatingBulkBPSPriceParam) ImplementsSubscriptionPriceIntervalsParamsAddPriceUnion() {}

// The cadence to bill for this price on.
type NewFloatingBulkBPSPriceCadence string

const (
	NewFloatingBulkBPSPriceCadenceAnnual     NewFloatingBulkBPSPriceCadence = "annual"
	NewFloatingBulkBPSPriceCadenceSemiAnnual NewFloatingBulkBPSPriceCadence = "semi_annual"
	NewFloatingBulkBPSPriceCadenceMonthly    NewFloatingBulkBPSPriceCadence = "monthly"
	NewFloatingBulkBPSPriceCadenceQuarterly  NewFloatingBulkBPSPriceCadence = "quarterly"
	NewFloatingBulkBPSPriceCadenceOneTime    NewFloatingBulkBPSPriceCadence = "one_time"
	NewFloatingBulkBPSPriceCadenceCustom     NewFloatingBulkBPSPriceCadence = "custom"
)

func (r NewFloatingBulkBPSPriceCadence) IsKnown() bool {
	switch r {
	case NewFloatingBulkBPSPriceCadenceAnnual, NewFloatingBulkBPSPriceCadenceSemiAnnual, NewFloatingBulkBPSPriceCadenceMonthly, NewFloatingBulkBPSPriceCadenceQuarterly, NewFloatingBulkBPSPriceCadenceOneTime, NewFloatingBulkBPSPriceCadenceCustom:
		return true
	}
	return false
}

type NewFloatingBulkBPSPriceModelType string

const (
	NewFloatingBulkBPSPriceModelTypeBulkBPS NewFloatingBulkBPSPriceModelType = "bulk_bps"
)

func (r NewFloatingBulkBPSPriceModelType) IsKnown() bool {
	switch r {
	case NewFloatingBulkBPSPriceModelTypeBulkBPS:
		return true
	}
	return false
}

type NewFloatingBulkBPSPriceParam struct {
	ConversionRateType param.Field[NewFloatingBulkBPSPriceConversionRateType] `json:"conversion_rate_type,required"`
	TieredConfig       param.Field[ConversionRateTieredConfigParam]           `json:"tiered_config"`
	UnitConfig         param.Field[ConversionRateUnitConfigParam]             `json:"unit_config"`
}

func (r NewFloatingBulkBPSPriceParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r NewFloatingBulkBPSPriceParam) ImplementsNewFloatingBulkBPSPriceUnionParam() {}

// Satisfied by [shared.UnitConversionRateConfigParam],
// [shared.TieredConversionRateConfigParam], [NewFloatingBulkBPSPriceParam].
type NewFloatingBulkBPSPriceUnionParam interface {
	ImplementsNewFloatingBulkBPSPriceUnionParam()
}

type NewFloatingBulkBPSPriceConversionRateType string

const (
	NewFloatingBulkBPSPriceConversionRateTypeUnit   NewFloatingBulkBPSPriceConversionRateType = "unit"
	NewFloatingBulkBPSPriceConversionRateTypeTiered NewFloatingBulkBPSPriceConversionRateType = "tiered"
)

func (r NewFloatingBulkBPSPriceConversionRateType) IsKnown() bool {
	switch r {
	case NewFloatingBulkBPSPriceConversionRateTypeUnit, NewFloatingBulkBPSPriceConversionRateTypeTiered:
		return true
	}
	return false
}

type NewFloatingBulkPriceParam struct {
	BulkConfig param.Field[BulkConfigParam] `json:"bulk_config,required"`
	// The cadence to bill for this price on.
	Cadence param.Field[NewFloatingBulkPriceCadence] `json:"cadence,required"`
	// An ISO 4217 currency string for which this price is billed in.
	Currency param.Field[string] `json:"currency,required"`
	// The id of the item the price will be associated with.
	ItemID    param.Field[string]                        `json:"item_id,required"`
	ModelType param.Field[NewFloatingBulkPriceModelType] `json:"model_type,required"`
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
	BillingCycleConfiguration param.Field[NewBillingCycleConfigurationParam] `json:"billing_cycle_configuration"`
	// The per unit conversion rate of the price currency to the invoicing currency.
	ConversionRate param.Field[float64] `json:"conversion_rate"`
	// The configuration for the rate of the price currency to the invoicing currency.
	ConversionRateConfig param.Field[NewFloatingBulkPriceUnionParam] `json:"conversion_rate_config"`
	// For dimensional price: specifies a price group and dimension values
	DimensionalPriceConfiguration param.Field[NewDimensionalPriceConfigurationParam] `json:"dimensional_price_configuration"`
	// An alias for the price.
	ExternalPriceID param.Field[string] `json:"external_price_id"`
	// If the Price represents a fixed cost, this represents the quantity of units
	// applied.
	FixedPriceQuantity param.Field[float64] `json:"fixed_price_quantity"`
	// The property used to group this price on an invoice
	InvoiceGroupingKey param.Field[string] `json:"invoice_grouping_key"`
	// Within each billing cycle, specifies the cadence at which invoices are produced.
	// If unspecified, a single invoice is produced per billing cycle.
	InvoicingCycleConfiguration param.Field[NewBillingCycleConfigurationParam] `json:"invoicing_cycle_configuration"`
	// User-specified key/value pairs for the resource. Individual keys can be removed
	// by setting the value to `null`, and the entire metadata mapping can be cleared
	// by setting `metadata` to `null`.
	Metadata param.Field[map[string]string] `json:"metadata"`
}

func (r NewFloatingBulkPriceParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r NewFloatingBulkPriceParam) ImplementsPriceEvaluateMultipleParamsPriceEvaluationsPriceUnion() {
}

func (r NewFloatingBulkPriceParam) ImplementsPriceEvaluatePreviewEventsParamsPriceEvaluationsPriceUnion() {
}

func (r NewFloatingBulkPriceParam) ImplementsSubscriptionPriceIntervalsParamsAddPriceUnion() {}

// The cadence to bill for this price on.
type NewFloatingBulkPriceCadence string

const (
	NewFloatingBulkPriceCadenceAnnual     NewFloatingBulkPriceCadence = "annual"
	NewFloatingBulkPriceCadenceSemiAnnual NewFloatingBulkPriceCadence = "semi_annual"
	NewFloatingBulkPriceCadenceMonthly    NewFloatingBulkPriceCadence = "monthly"
	NewFloatingBulkPriceCadenceQuarterly  NewFloatingBulkPriceCadence = "quarterly"
	NewFloatingBulkPriceCadenceOneTime    NewFloatingBulkPriceCadence = "one_time"
	NewFloatingBulkPriceCadenceCustom     NewFloatingBulkPriceCadence = "custom"
)

func (r NewFloatingBulkPriceCadence) IsKnown() bool {
	switch r {
	case NewFloatingBulkPriceCadenceAnnual, NewFloatingBulkPriceCadenceSemiAnnual, NewFloatingBulkPriceCadenceMonthly, NewFloatingBulkPriceCadenceQuarterly, NewFloatingBulkPriceCadenceOneTime, NewFloatingBulkPriceCadenceCustom:
		return true
	}
	return false
}

type NewFloatingBulkPriceModelType string

const (
	NewFloatingBulkPriceModelTypeBulk NewFloatingBulkPriceModelType = "bulk"
)

func (r NewFloatingBulkPriceModelType) IsKnown() bool {
	switch r {
	case NewFloatingBulkPriceModelTypeBulk:
		return true
	}
	return false
}

type NewFloatingBulkPriceParam struct {
	ConversionRateType param.Field[NewFloatingBulkPriceConversionRateType] `json:"conversion_rate_type,required"`
	TieredConfig       param.Field[ConversionRateTieredConfigParam]        `json:"tiered_config"`
	UnitConfig         param.Field[ConversionRateUnitConfigParam]          `json:"unit_config"`
}

func (r NewFloatingBulkPriceParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r NewFloatingBulkPriceParam) ImplementsNewFloatingBulkPriceUnionParam() {}

// Satisfied by [shared.UnitConversionRateConfigParam],
// [shared.TieredConversionRateConfigParam], [NewFloatingBulkPriceParam].
type NewFloatingBulkPriceUnionParam interface {
	ImplementsNewFloatingBulkPriceUnionParam()
}

type NewFloatingBulkPriceConversionRateType string

const (
	NewFloatingBulkPriceConversionRateTypeUnit   NewFloatingBulkPriceConversionRateType = "unit"
	NewFloatingBulkPriceConversionRateTypeTiered NewFloatingBulkPriceConversionRateType = "tiered"
)

func (r NewFloatingBulkPriceConversionRateType) IsKnown() bool {
	switch r {
	case NewFloatingBulkPriceConversionRateTypeUnit, NewFloatingBulkPriceConversionRateTypeTiered:
		return true
	}
	return false
}

type NewFloatingBulkWithProrationPriceParam struct {
	BulkWithProrationConfig param.Field[map[string]interface{}] `json:"bulk_with_proration_config,required"`
	// The cadence to bill for this price on.
	Cadence param.Field[NewFloatingBulkWithProrationPriceCadence] `json:"cadence,required"`
	// An ISO 4217 currency string for which this price is billed in.
	Currency param.Field[string] `json:"currency,required"`
	// The id of the item the price will be associated with.
	ItemID    param.Field[string]                                     `json:"item_id,required"`
	ModelType param.Field[NewFloatingBulkWithProrationPriceModelType] `json:"model_type,required"`
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
	BillingCycleConfiguration param.Field[NewBillingCycleConfigurationParam] `json:"billing_cycle_configuration"`
	// The per unit conversion rate of the price currency to the invoicing currency.
	ConversionRate param.Field[float64] `json:"conversion_rate"`
	// The configuration for the rate of the price currency to the invoicing currency.
	ConversionRateConfig param.Field[NewFloatingBulkWithProrationPriceUnionParam] `json:"conversion_rate_config"`
	// For dimensional price: specifies a price group and dimension values
	DimensionalPriceConfiguration param.Field[NewDimensionalPriceConfigurationParam] `json:"dimensional_price_configuration"`
	// An alias for the price.
	ExternalPriceID param.Field[string] `json:"external_price_id"`
	// If the Price represents a fixed cost, this represents the quantity of units
	// applied.
	FixedPriceQuantity param.Field[float64] `json:"fixed_price_quantity"`
	// The property used to group this price on an invoice
	InvoiceGroupingKey param.Field[string] `json:"invoice_grouping_key"`
	// Within each billing cycle, specifies the cadence at which invoices are produced.
	// If unspecified, a single invoice is produced per billing cycle.
	InvoicingCycleConfiguration param.Field[NewBillingCycleConfigurationParam] `json:"invoicing_cycle_configuration"`
	// User-specified key/value pairs for the resource. Individual keys can be removed
	// by setting the value to `null`, and the entire metadata mapping can be cleared
	// by setting `metadata` to `null`.
	Metadata param.Field[map[string]string] `json:"metadata"`
}

func (r NewFloatingBulkWithProrationPriceParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r NewFloatingBulkWithProrationPriceParam) ImplementsPriceEvaluateMultipleParamsPriceEvaluationsPriceUnion() {
}

func (r NewFloatingBulkWithProrationPriceParam) ImplementsPriceEvaluatePreviewEventsParamsPriceEvaluationsPriceUnion() {
}

func (r NewFloatingBulkWithProrationPriceParam) ImplementsSubscriptionPriceIntervalsParamsAddPriceUnion() {
}

// The cadence to bill for this price on.
type NewFloatingBulkWithProrationPriceCadence string

const (
	NewFloatingBulkWithProrationPriceCadenceAnnual     NewFloatingBulkWithProrationPriceCadence = "annual"
	NewFloatingBulkWithProrationPriceCadenceSemiAnnual NewFloatingBulkWithProrationPriceCadence = "semi_annual"
	NewFloatingBulkWithProrationPriceCadenceMonthly    NewFloatingBulkWithProrationPriceCadence = "monthly"
	NewFloatingBulkWithProrationPriceCadenceQuarterly  NewFloatingBulkWithProrationPriceCadence = "quarterly"
	NewFloatingBulkWithProrationPriceCadenceOneTime    NewFloatingBulkWithProrationPriceCadence = "one_time"
	NewFloatingBulkWithProrationPriceCadenceCustom     NewFloatingBulkWithProrationPriceCadence = "custom"
)

func (r NewFloatingBulkWithProrationPriceCadence) IsKnown() bool {
	switch r {
	case NewFloatingBulkWithProrationPriceCadenceAnnual, NewFloatingBulkWithProrationPriceCadenceSemiAnnual, NewFloatingBulkWithProrationPriceCadenceMonthly, NewFloatingBulkWithProrationPriceCadenceQuarterly, NewFloatingBulkWithProrationPriceCadenceOneTime, NewFloatingBulkWithProrationPriceCadenceCustom:
		return true
	}
	return false
}

type NewFloatingBulkWithProrationPriceModelType string

const (
	NewFloatingBulkWithProrationPriceModelTypeBulkWithProration NewFloatingBulkWithProrationPriceModelType = "bulk_with_proration"
)

func (r NewFloatingBulkWithProrationPriceModelType) IsKnown() bool {
	switch r {
	case NewFloatingBulkWithProrationPriceModelTypeBulkWithProration:
		return true
	}
	return false
}

type NewFloatingBulkWithProrationPriceParam struct {
	ConversionRateType param.Field[NewFloatingBulkWithProrationPriceConversionRateType] `json:"conversion_rate_type,required"`
	TieredConfig       param.Field[ConversionRateTieredConfigParam]                     `json:"tiered_config"`
	UnitConfig         param.Field[ConversionRateUnitConfigParam]                       `json:"unit_config"`
}

func (r NewFloatingBulkWithProrationPriceParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r NewFloatingBulkWithProrationPriceParam) ImplementsNewFloatingBulkWithProrationPriceUnionParam() {
}

// Satisfied by [shared.UnitConversionRateConfigParam],
// [shared.TieredConversionRateConfigParam],
// [NewFloatingBulkWithProrationPriceParam].
type NewFloatingBulkWithProrationPriceUnionParam interface {
	ImplementsNewFloatingBulkWithProrationPriceUnionParam()
}

type NewFloatingBulkWithProrationPriceConversionRateType string

const (
	NewFloatingBulkWithProrationPriceConversionRateTypeUnit   NewFloatingBulkWithProrationPriceConversionRateType = "unit"
	NewFloatingBulkWithProrationPriceConversionRateTypeTiered NewFloatingBulkWithProrationPriceConversionRateType = "tiered"
)

func (r NewFloatingBulkWithProrationPriceConversionRateType) IsKnown() bool {
	switch r {
	case NewFloatingBulkWithProrationPriceConversionRateTypeUnit, NewFloatingBulkWithProrationPriceConversionRateTypeTiered:
		return true
	}
	return false
}

type NewFloatingCumulativeGroupedBulkPriceParam struct {
	// The cadence to bill for this price on.
	Cadence                     param.Field[NewFloatingCumulativeGroupedBulkPriceCadence] `json:"cadence,required"`
	CumulativeGroupedBulkConfig param.Field[map[string]interface{}]                       `json:"cumulative_grouped_bulk_config,required"`
	// An ISO 4217 currency string for which this price is billed in.
	Currency param.Field[string] `json:"currency,required"`
	// The id of the item the price will be associated with.
	ItemID    param.Field[string]                                         `json:"item_id,required"`
	ModelType param.Field[NewFloatingCumulativeGroupedBulkPriceModelType] `json:"model_type,required"`
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
	BillingCycleConfiguration param.Field[NewBillingCycleConfigurationParam] `json:"billing_cycle_configuration"`
	// The per unit conversion rate of the price currency to the invoicing currency.
	ConversionRate param.Field[float64] `json:"conversion_rate"`
	// The configuration for the rate of the price currency to the invoicing currency.
	ConversionRateConfig param.Field[NewFloatingCumulativeGroupedBulkPriceUnionParam] `json:"conversion_rate_config"`
	// For dimensional price: specifies a price group and dimension values
	DimensionalPriceConfiguration param.Field[NewDimensionalPriceConfigurationParam] `json:"dimensional_price_configuration"`
	// An alias for the price.
	ExternalPriceID param.Field[string] `json:"external_price_id"`
	// If the Price represents a fixed cost, this represents the quantity of units
	// applied.
	FixedPriceQuantity param.Field[float64] `json:"fixed_price_quantity"`
	// The property used to group this price on an invoice
	InvoiceGroupingKey param.Field[string] `json:"invoice_grouping_key"`
	// Within each billing cycle, specifies the cadence at which invoices are produced.
	// If unspecified, a single invoice is produced per billing cycle.
	InvoicingCycleConfiguration param.Field[NewBillingCycleConfigurationParam] `json:"invoicing_cycle_configuration"`
	// User-specified key/value pairs for the resource. Individual keys can be removed
	// by setting the value to `null`, and the entire metadata mapping can be cleared
	// by setting `metadata` to `null`.
	Metadata param.Field[map[string]string] `json:"metadata"`
}

func (r NewFloatingCumulativeGroupedBulkPriceParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r NewFloatingCumulativeGroupedBulkPriceParam) ImplementsPriceEvaluateMultipleParamsPriceEvaluationsPriceUnion() {
}

func (r NewFloatingCumulativeGroupedBulkPriceParam) ImplementsPriceEvaluatePreviewEventsParamsPriceEvaluationsPriceUnion() {
}

func (r NewFloatingCumulativeGroupedBulkPriceParam) ImplementsSubscriptionPriceIntervalsParamsAddPriceUnion() {
}

// The cadence to bill for this price on.
type NewFloatingCumulativeGroupedBulkPriceCadence string

const (
	NewFloatingCumulativeGroupedBulkPriceCadenceAnnual     NewFloatingCumulativeGroupedBulkPriceCadence = "annual"
	NewFloatingCumulativeGroupedBulkPriceCadenceSemiAnnual NewFloatingCumulativeGroupedBulkPriceCadence = "semi_annual"
	NewFloatingCumulativeGroupedBulkPriceCadenceMonthly    NewFloatingCumulativeGroupedBulkPriceCadence = "monthly"
	NewFloatingCumulativeGroupedBulkPriceCadenceQuarterly  NewFloatingCumulativeGroupedBulkPriceCadence = "quarterly"
	NewFloatingCumulativeGroupedBulkPriceCadenceOneTime    NewFloatingCumulativeGroupedBulkPriceCadence = "one_time"
	NewFloatingCumulativeGroupedBulkPriceCadenceCustom     NewFloatingCumulativeGroupedBulkPriceCadence = "custom"
)

func (r NewFloatingCumulativeGroupedBulkPriceCadence) IsKnown() bool {
	switch r {
	case NewFloatingCumulativeGroupedBulkPriceCadenceAnnual, NewFloatingCumulativeGroupedBulkPriceCadenceSemiAnnual, NewFloatingCumulativeGroupedBulkPriceCadenceMonthly, NewFloatingCumulativeGroupedBulkPriceCadenceQuarterly, NewFloatingCumulativeGroupedBulkPriceCadenceOneTime, NewFloatingCumulativeGroupedBulkPriceCadenceCustom:
		return true
	}
	return false
}

type NewFloatingCumulativeGroupedBulkPriceModelType string

const (
	NewFloatingCumulativeGroupedBulkPriceModelTypeCumulativeGroupedBulk NewFloatingCumulativeGroupedBulkPriceModelType = "cumulative_grouped_bulk"
)

func (r NewFloatingCumulativeGroupedBulkPriceModelType) IsKnown() bool {
	switch r {
	case NewFloatingCumulativeGroupedBulkPriceModelTypeCumulativeGroupedBulk:
		return true
	}
	return false
}

type NewFloatingCumulativeGroupedBulkPriceParam struct {
	ConversionRateType param.Field[NewFloatingCumulativeGroupedBulkPriceConversionRateType] `json:"conversion_rate_type,required"`
	TieredConfig       param.Field[ConversionRateTieredConfigParam]                         `json:"tiered_config"`
	UnitConfig         param.Field[ConversionRateUnitConfigParam]                           `json:"unit_config"`
}

func (r NewFloatingCumulativeGroupedBulkPriceParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r NewFloatingCumulativeGroupedBulkPriceParam) ImplementsNewFloatingCumulativeGroupedBulkPriceUnionParam() {
}

// Satisfied by [shared.UnitConversionRateConfigParam],
// [shared.TieredConversionRateConfigParam],
// [NewFloatingCumulativeGroupedBulkPriceParam].
type NewFloatingCumulativeGroupedBulkPriceUnionParam interface {
	ImplementsNewFloatingCumulativeGroupedBulkPriceUnionParam()
}

type NewFloatingCumulativeGroupedBulkPriceConversionRateType string

const (
	NewFloatingCumulativeGroupedBulkPriceConversionRateTypeUnit   NewFloatingCumulativeGroupedBulkPriceConversionRateType = "unit"
	NewFloatingCumulativeGroupedBulkPriceConversionRateTypeTiered NewFloatingCumulativeGroupedBulkPriceConversionRateType = "tiered"
)

func (r NewFloatingCumulativeGroupedBulkPriceConversionRateType) IsKnown() bool {
	switch r {
	case NewFloatingCumulativeGroupedBulkPriceConversionRateTypeUnit, NewFloatingCumulativeGroupedBulkPriceConversionRateTypeTiered:
		return true
	}
	return false
}

type NewFloatingGroupedAllocationPriceParam struct {
	// The cadence to bill for this price on.
	Cadence param.Field[NewFloatingGroupedAllocationPriceCadence] `json:"cadence,required"`
	// An ISO 4217 currency string for which this price is billed in.
	Currency                param.Field[string]                 `json:"currency,required"`
	GroupedAllocationConfig param.Field[map[string]interface{}] `json:"grouped_allocation_config,required"`
	// The id of the item the price will be associated with.
	ItemID    param.Field[string]                                     `json:"item_id,required"`
	ModelType param.Field[NewFloatingGroupedAllocationPriceModelType] `json:"model_type,required"`
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
	BillingCycleConfiguration param.Field[NewBillingCycleConfigurationParam] `json:"billing_cycle_configuration"`
	// The per unit conversion rate of the price currency to the invoicing currency.
	ConversionRate param.Field[float64] `json:"conversion_rate"`
	// The configuration for the rate of the price currency to the invoicing currency.
	ConversionRateConfig param.Field[NewFloatingGroupedAllocationPriceUnionParam] `json:"conversion_rate_config"`
	// For dimensional price: specifies a price group and dimension values
	DimensionalPriceConfiguration param.Field[NewDimensionalPriceConfigurationParam] `json:"dimensional_price_configuration"`
	// An alias for the price.
	ExternalPriceID param.Field[string] `json:"external_price_id"`
	// If the Price represents a fixed cost, this represents the quantity of units
	// applied.
	FixedPriceQuantity param.Field[float64] `json:"fixed_price_quantity"`
	// The property used to group this price on an invoice
	InvoiceGroupingKey param.Field[string] `json:"invoice_grouping_key"`
	// Within each billing cycle, specifies the cadence at which invoices are produced.
	// If unspecified, a single invoice is produced per billing cycle.
	InvoicingCycleConfiguration param.Field[NewBillingCycleConfigurationParam] `json:"invoicing_cycle_configuration"`
	// User-specified key/value pairs for the resource. Individual keys can be removed
	// by setting the value to `null`, and the entire metadata mapping can be cleared
	// by setting `metadata` to `null`.
	Metadata param.Field[map[string]string] `json:"metadata"`
}

func (r NewFloatingGroupedAllocationPriceParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r NewFloatingGroupedAllocationPriceParam) ImplementsPriceEvaluateMultipleParamsPriceEvaluationsPriceUnion() {
}

func (r NewFloatingGroupedAllocationPriceParam) ImplementsPriceEvaluatePreviewEventsParamsPriceEvaluationsPriceUnion() {
}

func (r NewFloatingGroupedAllocationPriceParam) ImplementsSubscriptionPriceIntervalsParamsAddPriceUnion() {
}

// The cadence to bill for this price on.
type NewFloatingGroupedAllocationPriceCadence string

const (
	NewFloatingGroupedAllocationPriceCadenceAnnual     NewFloatingGroupedAllocationPriceCadence = "annual"
	NewFloatingGroupedAllocationPriceCadenceSemiAnnual NewFloatingGroupedAllocationPriceCadence = "semi_annual"
	NewFloatingGroupedAllocationPriceCadenceMonthly    NewFloatingGroupedAllocationPriceCadence = "monthly"
	NewFloatingGroupedAllocationPriceCadenceQuarterly  NewFloatingGroupedAllocationPriceCadence = "quarterly"
	NewFloatingGroupedAllocationPriceCadenceOneTime    NewFloatingGroupedAllocationPriceCadence = "one_time"
	NewFloatingGroupedAllocationPriceCadenceCustom     NewFloatingGroupedAllocationPriceCadence = "custom"
)

func (r NewFloatingGroupedAllocationPriceCadence) IsKnown() bool {
	switch r {
	case NewFloatingGroupedAllocationPriceCadenceAnnual, NewFloatingGroupedAllocationPriceCadenceSemiAnnual, NewFloatingGroupedAllocationPriceCadenceMonthly, NewFloatingGroupedAllocationPriceCadenceQuarterly, NewFloatingGroupedAllocationPriceCadenceOneTime, NewFloatingGroupedAllocationPriceCadenceCustom:
		return true
	}
	return false
}

type NewFloatingGroupedAllocationPriceModelType string

const (
	NewFloatingGroupedAllocationPriceModelTypeGroupedAllocation NewFloatingGroupedAllocationPriceModelType = "grouped_allocation"
)

func (r NewFloatingGroupedAllocationPriceModelType) IsKnown() bool {
	switch r {
	case NewFloatingGroupedAllocationPriceModelTypeGroupedAllocation:
		return true
	}
	return false
}

type NewFloatingGroupedAllocationPriceParam struct {
	ConversionRateType param.Field[NewFloatingGroupedAllocationPriceConversionRateType] `json:"conversion_rate_type,required"`
	TieredConfig       param.Field[ConversionRateTieredConfigParam]                     `json:"tiered_config"`
	UnitConfig         param.Field[ConversionRateUnitConfigParam]                       `json:"unit_config"`
}

func (r NewFloatingGroupedAllocationPriceParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r NewFloatingGroupedAllocationPriceParam) ImplementsNewFloatingGroupedAllocationPriceUnionParam() {
}

// Satisfied by [shared.UnitConversionRateConfigParam],
// [shared.TieredConversionRateConfigParam],
// [NewFloatingGroupedAllocationPriceParam].
type NewFloatingGroupedAllocationPriceUnionParam interface {
	ImplementsNewFloatingGroupedAllocationPriceUnionParam()
}

type NewFloatingGroupedAllocationPriceConversionRateType string

const (
	NewFloatingGroupedAllocationPriceConversionRateTypeUnit   NewFloatingGroupedAllocationPriceConversionRateType = "unit"
	NewFloatingGroupedAllocationPriceConversionRateTypeTiered NewFloatingGroupedAllocationPriceConversionRateType = "tiered"
)

func (r NewFloatingGroupedAllocationPriceConversionRateType) IsKnown() bool {
	switch r {
	case NewFloatingGroupedAllocationPriceConversionRateTypeUnit, NewFloatingGroupedAllocationPriceConversionRateTypeTiered:
		return true
	}
	return false
}

type NewFloatingGroupedTieredPackagePriceParam struct {
	// The cadence to bill for this price on.
	Cadence param.Field[NewFloatingGroupedTieredPackagePriceCadence] `json:"cadence,required"`
	// An ISO 4217 currency string for which this price is billed in.
	Currency                   param.Field[string]                 `json:"currency,required"`
	GroupedTieredPackageConfig param.Field[map[string]interface{}] `json:"grouped_tiered_package_config,required"`
	// The id of the item the price will be associated with.
	ItemID    param.Field[string]                                        `json:"item_id,required"`
	ModelType param.Field[NewFloatingGroupedTieredPackagePriceModelType] `json:"model_type,required"`
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
	BillingCycleConfiguration param.Field[NewBillingCycleConfigurationParam] `json:"billing_cycle_configuration"`
	// The per unit conversion rate of the price currency to the invoicing currency.
	ConversionRate param.Field[float64] `json:"conversion_rate"`
	// The configuration for the rate of the price currency to the invoicing currency.
	ConversionRateConfig param.Field[NewFloatingGroupedTieredPackagePriceUnionParam] `json:"conversion_rate_config"`
	// For dimensional price: specifies a price group and dimension values
	DimensionalPriceConfiguration param.Field[NewDimensionalPriceConfigurationParam] `json:"dimensional_price_configuration"`
	// An alias for the price.
	ExternalPriceID param.Field[string] `json:"external_price_id"`
	// If the Price represents a fixed cost, this represents the quantity of units
	// applied.
	FixedPriceQuantity param.Field[float64] `json:"fixed_price_quantity"`
	// The property used to group this price on an invoice
	InvoiceGroupingKey param.Field[string] `json:"invoice_grouping_key"`
	// Within each billing cycle, specifies the cadence at which invoices are produced.
	// If unspecified, a single invoice is produced per billing cycle.
	InvoicingCycleConfiguration param.Field[NewBillingCycleConfigurationParam] `json:"invoicing_cycle_configuration"`
	// User-specified key/value pairs for the resource. Individual keys can be removed
	// by setting the value to `null`, and the entire metadata mapping can be cleared
	// by setting `metadata` to `null`.
	Metadata param.Field[map[string]string] `json:"metadata"`
}

func (r NewFloatingGroupedTieredPackagePriceParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r NewFloatingGroupedTieredPackagePriceParam) ImplementsPriceEvaluateMultipleParamsPriceEvaluationsPriceUnion() {
}

func (r NewFloatingGroupedTieredPackagePriceParam) ImplementsPriceEvaluatePreviewEventsParamsPriceEvaluationsPriceUnion() {
}

func (r NewFloatingGroupedTieredPackagePriceParam) ImplementsSubscriptionPriceIntervalsParamsAddPriceUnion() {
}

// The cadence to bill for this price on.
type NewFloatingGroupedTieredPackagePriceCadence string

const (
	NewFloatingGroupedTieredPackagePriceCadenceAnnual     NewFloatingGroupedTieredPackagePriceCadence = "annual"
	NewFloatingGroupedTieredPackagePriceCadenceSemiAnnual NewFloatingGroupedTieredPackagePriceCadence = "semi_annual"
	NewFloatingGroupedTieredPackagePriceCadenceMonthly    NewFloatingGroupedTieredPackagePriceCadence = "monthly"
	NewFloatingGroupedTieredPackagePriceCadenceQuarterly  NewFloatingGroupedTieredPackagePriceCadence = "quarterly"
	NewFloatingGroupedTieredPackagePriceCadenceOneTime    NewFloatingGroupedTieredPackagePriceCadence = "one_time"
	NewFloatingGroupedTieredPackagePriceCadenceCustom     NewFloatingGroupedTieredPackagePriceCadence = "custom"
)

func (r NewFloatingGroupedTieredPackagePriceCadence) IsKnown() bool {
	switch r {
	case NewFloatingGroupedTieredPackagePriceCadenceAnnual, NewFloatingGroupedTieredPackagePriceCadenceSemiAnnual, NewFloatingGroupedTieredPackagePriceCadenceMonthly, NewFloatingGroupedTieredPackagePriceCadenceQuarterly, NewFloatingGroupedTieredPackagePriceCadenceOneTime, NewFloatingGroupedTieredPackagePriceCadenceCustom:
		return true
	}
	return false
}

type NewFloatingGroupedTieredPackagePriceModelType string

const (
	NewFloatingGroupedTieredPackagePriceModelTypeGroupedTieredPackage NewFloatingGroupedTieredPackagePriceModelType = "grouped_tiered_package"
)

func (r NewFloatingGroupedTieredPackagePriceModelType) IsKnown() bool {
	switch r {
	case NewFloatingGroupedTieredPackagePriceModelTypeGroupedTieredPackage:
		return true
	}
	return false
}

type NewFloatingGroupedTieredPackagePriceParam struct {
	ConversionRateType param.Field[NewFloatingGroupedTieredPackagePriceConversionRateType] `json:"conversion_rate_type,required"`
	TieredConfig       param.Field[ConversionRateTieredConfigParam]                        `json:"tiered_config"`
	UnitConfig         param.Field[ConversionRateUnitConfigParam]                          `json:"unit_config"`
}

func (r NewFloatingGroupedTieredPackagePriceParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r NewFloatingGroupedTieredPackagePriceParam) ImplementsNewFloatingGroupedTieredPackagePriceUnionParam() {
}

// Satisfied by [shared.UnitConversionRateConfigParam],
// [shared.TieredConversionRateConfigParam],
// [NewFloatingGroupedTieredPackagePriceParam].
type NewFloatingGroupedTieredPackagePriceUnionParam interface {
	ImplementsNewFloatingGroupedTieredPackagePriceUnionParam()
}

type NewFloatingGroupedTieredPackagePriceConversionRateType string

const (
	NewFloatingGroupedTieredPackagePriceConversionRateTypeUnit   NewFloatingGroupedTieredPackagePriceConversionRateType = "unit"
	NewFloatingGroupedTieredPackagePriceConversionRateTypeTiered NewFloatingGroupedTieredPackagePriceConversionRateType = "tiered"
)

func (r NewFloatingGroupedTieredPackagePriceConversionRateType) IsKnown() bool {
	switch r {
	case NewFloatingGroupedTieredPackagePriceConversionRateTypeUnit, NewFloatingGroupedTieredPackagePriceConversionRateTypeTiered:
		return true
	}
	return false
}

type NewFloatingGroupedTieredPriceParam struct {
	// The cadence to bill for this price on.
	Cadence param.Field[NewFloatingGroupedTieredPriceCadence] `json:"cadence,required"`
	// An ISO 4217 currency string for which this price is billed in.
	Currency            param.Field[string]                 `json:"currency,required"`
	GroupedTieredConfig param.Field[map[string]interface{}] `json:"grouped_tiered_config,required"`
	// The id of the item the price will be associated with.
	ItemID    param.Field[string]                                 `json:"item_id,required"`
	ModelType param.Field[NewFloatingGroupedTieredPriceModelType] `json:"model_type,required"`
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
	BillingCycleConfiguration param.Field[NewBillingCycleConfigurationParam] `json:"billing_cycle_configuration"`
	// The per unit conversion rate of the price currency to the invoicing currency.
	ConversionRate param.Field[float64] `json:"conversion_rate"`
	// The configuration for the rate of the price currency to the invoicing currency.
	ConversionRateConfig param.Field[NewFloatingGroupedTieredPriceUnionParam] `json:"conversion_rate_config"`
	// For dimensional price: specifies a price group and dimension values
	DimensionalPriceConfiguration param.Field[NewDimensionalPriceConfigurationParam] `json:"dimensional_price_configuration"`
	// An alias for the price.
	ExternalPriceID param.Field[string] `json:"external_price_id"`
	// If the Price represents a fixed cost, this represents the quantity of units
	// applied.
	FixedPriceQuantity param.Field[float64] `json:"fixed_price_quantity"`
	// The property used to group this price on an invoice
	InvoiceGroupingKey param.Field[string] `json:"invoice_grouping_key"`
	// Within each billing cycle, specifies the cadence at which invoices are produced.
	// If unspecified, a single invoice is produced per billing cycle.
	InvoicingCycleConfiguration param.Field[NewBillingCycleConfigurationParam] `json:"invoicing_cycle_configuration"`
	// User-specified key/value pairs for the resource. Individual keys can be removed
	// by setting the value to `null`, and the entire metadata mapping can be cleared
	// by setting `metadata` to `null`.
	Metadata param.Field[map[string]string] `json:"metadata"`
}

func (r NewFloatingGroupedTieredPriceParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r NewFloatingGroupedTieredPriceParam) ImplementsPriceEvaluateMultipleParamsPriceEvaluationsPriceUnion() {
}

func (r NewFloatingGroupedTieredPriceParam) ImplementsPriceEvaluatePreviewEventsParamsPriceEvaluationsPriceUnion() {
}

func (r NewFloatingGroupedTieredPriceParam) ImplementsSubscriptionPriceIntervalsParamsAddPriceUnion() {
}

// The cadence to bill for this price on.
type NewFloatingGroupedTieredPriceCadence string

const (
	NewFloatingGroupedTieredPriceCadenceAnnual     NewFloatingGroupedTieredPriceCadence = "annual"
	NewFloatingGroupedTieredPriceCadenceSemiAnnual NewFloatingGroupedTieredPriceCadence = "semi_annual"
	NewFloatingGroupedTieredPriceCadenceMonthly    NewFloatingGroupedTieredPriceCadence = "monthly"
	NewFloatingGroupedTieredPriceCadenceQuarterly  NewFloatingGroupedTieredPriceCadence = "quarterly"
	NewFloatingGroupedTieredPriceCadenceOneTime    NewFloatingGroupedTieredPriceCadence = "one_time"
	NewFloatingGroupedTieredPriceCadenceCustom     NewFloatingGroupedTieredPriceCadence = "custom"
)

func (r NewFloatingGroupedTieredPriceCadence) IsKnown() bool {
	switch r {
	case NewFloatingGroupedTieredPriceCadenceAnnual, NewFloatingGroupedTieredPriceCadenceSemiAnnual, NewFloatingGroupedTieredPriceCadenceMonthly, NewFloatingGroupedTieredPriceCadenceQuarterly, NewFloatingGroupedTieredPriceCadenceOneTime, NewFloatingGroupedTieredPriceCadenceCustom:
		return true
	}
	return false
}

type NewFloatingGroupedTieredPriceModelType string

const (
	NewFloatingGroupedTieredPriceModelTypeGroupedTiered NewFloatingGroupedTieredPriceModelType = "grouped_tiered"
)

func (r NewFloatingGroupedTieredPriceModelType) IsKnown() bool {
	switch r {
	case NewFloatingGroupedTieredPriceModelTypeGroupedTiered:
		return true
	}
	return false
}

type NewFloatingGroupedTieredPriceParam struct {
	ConversionRateType param.Field[NewFloatingGroupedTieredPriceConversionRateType] `json:"conversion_rate_type,required"`
	TieredConfig       param.Field[ConversionRateTieredConfigParam]                 `json:"tiered_config"`
	UnitConfig         param.Field[ConversionRateUnitConfigParam]                   `json:"unit_config"`
}

func (r NewFloatingGroupedTieredPriceParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r NewFloatingGroupedTieredPriceParam) ImplementsNewFloatingGroupedTieredPriceUnionParam() {}

// Satisfied by [shared.UnitConversionRateConfigParam],
// [shared.TieredConversionRateConfigParam], [NewFloatingGroupedTieredPriceParam].
type NewFloatingGroupedTieredPriceUnionParam interface {
	ImplementsNewFloatingGroupedTieredPriceUnionParam()
}

type NewFloatingGroupedTieredPriceConversionRateType string

const (
	NewFloatingGroupedTieredPriceConversionRateTypeUnit   NewFloatingGroupedTieredPriceConversionRateType = "unit"
	NewFloatingGroupedTieredPriceConversionRateTypeTiered NewFloatingGroupedTieredPriceConversionRateType = "tiered"
)

func (r NewFloatingGroupedTieredPriceConversionRateType) IsKnown() bool {
	switch r {
	case NewFloatingGroupedTieredPriceConversionRateTypeUnit, NewFloatingGroupedTieredPriceConversionRateTypeTiered:
		return true
	}
	return false
}

type NewFloatingGroupedWithMeteredMinimumPriceParam struct {
	// The cadence to bill for this price on.
	Cadence param.Field[NewFloatingGroupedWithMeteredMinimumPriceCadence] `json:"cadence,required"`
	// An ISO 4217 currency string for which this price is billed in.
	Currency                        param.Field[string]                 `json:"currency,required"`
	GroupedWithMeteredMinimumConfig param.Field[map[string]interface{}] `json:"grouped_with_metered_minimum_config,required"`
	// The id of the item the price will be associated with.
	ItemID    param.Field[string]                                             `json:"item_id,required"`
	ModelType param.Field[NewFloatingGroupedWithMeteredMinimumPriceModelType] `json:"model_type,required"`
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
	BillingCycleConfiguration param.Field[NewBillingCycleConfigurationParam] `json:"billing_cycle_configuration"`
	// The per unit conversion rate of the price currency to the invoicing currency.
	ConversionRate param.Field[float64] `json:"conversion_rate"`
	// The configuration for the rate of the price currency to the invoicing currency.
	ConversionRateConfig param.Field[NewFloatingGroupedWithMeteredMinimumPriceUnionParam] `json:"conversion_rate_config"`
	// For dimensional price: specifies a price group and dimension values
	DimensionalPriceConfiguration param.Field[NewDimensionalPriceConfigurationParam] `json:"dimensional_price_configuration"`
	// An alias for the price.
	ExternalPriceID param.Field[string] `json:"external_price_id"`
	// If the Price represents a fixed cost, this represents the quantity of units
	// applied.
	FixedPriceQuantity param.Field[float64] `json:"fixed_price_quantity"`
	// The property used to group this price on an invoice
	InvoiceGroupingKey param.Field[string] `json:"invoice_grouping_key"`
	// Within each billing cycle, specifies the cadence at which invoices are produced.
	// If unspecified, a single invoice is produced per billing cycle.
	InvoicingCycleConfiguration param.Field[NewBillingCycleConfigurationParam] `json:"invoicing_cycle_configuration"`
	// User-specified key/value pairs for the resource. Individual keys can be removed
	// by setting the value to `null`, and the entire metadata mapping can be cleared
	// by setting `metadata` to `null`.
	Metadata param.Field[map[string]string] `json:"metadata"`
}

func (r NewFloatingGroupedWithMeteredMinimumPriceParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r NewFloatingGroupedWithMeteredMinimumPriceParam) ImplementsPriceEvaluateMultipleParamsPriceEvaluationsPriceUnion() {
}

func (r NewFloatingGroupedWithMeteredMinimumPriceParam) ImplementsPriceEvaluatePreviewEventsParamsPriceEvaluationsPriceUnion() {
}

func (r NewFloatingGroupedWithMeteredMinimumPriceParam) ImplementsSubscriptionPriceIntervalsParamsAddPriceUnion() {
}

// The cadence to bill for this price on.
type NewFloatingGroupedWithMeteredMinimumPriceCadence string

const (
	NewFloatingGroupedWithMeteredMinimumPriceCadenceAnnual     NewFloatingGroupedWithMeteredMinimumPriceCadence = "annual"
	NewFloatingGroupedWithMeteredMinimumPriceCadenceSemiAnnual NewFloatingGroupedWithMeteredMinimumPriceCadence = "semi_annual"
	NewFloatingGroupedWithMeteredMinimumPriceCadenceMonthly    NewFloatingGroupedWithMeteredMinimumPriceCadence = "monthly"
	NewFloatingGroupedWithMeteredMinimumPriceCadenceQuarterly  NewFloatingGroupedWithMeteredMinimumPriceCadence = "quarterly"
	NewFloatingGroupedWithMeteredMinimumPriceCadenceOneTime    NewFloatingGroupedWithMeteredMinimumPriceCadence = "one_time"
	NewFloatingGroupedWithMeteredMinimumPriceCadenceCustom     NewFloatingGroupedWithMeteredMinimumPriceCadence = "custom"
)

func (r NewFloatingGroupedWithMeteredMinimumPriceCadence) IsKnown() bool {
	switch r {
	case NewFloatingGroupedWithMeteredMinimumPriceCadenceAnnual, NewFloatingGroupedWithMeteredMinimumPriceCadenceSemiAnnual, NewFloatingGroupedWithMeteredMinimumPriceCadenceMonthly, NewFloatingGroupedWithMeteredMinimumPriceCadenceQuarterly, NewFloatingGroupedWithMeteredMinimumPriceCadenceOneTime, NewFloatingGroupedWithMeteredMinimumPriceCadenceCustom:
		return true
	}
	return false
}

type NewFloatingGroupedWithMeteredMinimumPriceModelType string

const (
	NewFloatingGroupedWithMeteredMinimumPriceModelTypeGroupedWithMeteredMinimum NewFloatingGroupedWithMeteredMinimumPriceModelType = "grouped_with_metered_minimum"
)

func (r NewFloatingGroupedWithMeteredMinimumPriceModelType) IsKnown() bool {
	switch r {
	case NewFloatingGroupedWithMeteredMinimumPriceModelTypeGroupedWithMeteredMinimum:
		return true
	}
	return false
}

type NewFloatingGroupedWithMeteredMinimumPriceParam struct {
	ConversionRateType param.Field[NewFloatingGroupedWithMeteredMinimumPriceConversionRateType] `json:"conversion_rate_type,required"`
	TieredConfig       param.Field[ConversionRateTieredConfigParam]                             `json:"tiered_config"`
	UnitConfig         param.Field[ConversionRateUnitConfigParam]                               `json:"unit_config"`
}

func (r NewFloatingGroupedWithMeteredMinimumPriceParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r NewFloatingGroupedWithMeteredMinimumPriceParam) ImplementsNewFloatingGroupedWithMeteredMinimumPriceUnionParam() {
}

// Satisfied by [shared.UnitConversionRateConfigParam],
// [shared.TieredConversionRateConfigParam],
// [NewFloatingGroupedWithMeteredMinimumPriceParam].
type NewFloatingGroupedWithMeteredMinimumPriceUnionParam interface {
	ImplementsNewFloatingGroupedWithMeteredMinimumPriceUnionParam()
}

type NewFloatingGroupedWithMeteredMinimumPriceConversionRateType string

const (
	NewFloatingGroupedWithMeteredMinimumPriceConversionRateTypeUnit   NewFloatingGroupedWithMeteredMinimumPriceConversionRateType = "unit"
	NewFloatingGroupedWithMeteredMinimumPriceConversionRateTypeTiered NewFloatingGroupedWithMeteredMinimumPriceConversionRateType = "tiered"
)

func (r NewFloatingGroupedWithMeteredMinimumPriceConversionRateType) IsKnown() bool {
	switch r {
	case NewFloatingGroupedWithMeteredMinimumPriceConversionRateTypeUnit, NewFloatingGroupedWithMeteredMinimumPriceConversionRateTypeTiered:
		return true
	}
	return false
}

type NewFloatingGroupedWithProratedMinimumPriceParam struct {
	// The cadence to bill for this price on.
	Cadence param.Field[NewFloatingGroupedWithProratedMinimumPriceCadence] `json:"cadence,required"`
	// An ISO 4217 currency string for which this price is billed in.
	Currency                         param.Field[string]                 `json:"currency,required"`
	GroupedWithProratedMinimumConfig param.Field[map[string]interface{}] `json:"grouped_with_prorated_minimum_config,required"`
	// The id of the item the price will be associated with.
	ItemID    param.Field[string]                                              `json:"item_id,required"`
	ModelType param.Field[NewFloatingGroupedWithProratedMinimumPriceModelType] `json:"model_type,required"`
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
	BillingCycleConfiguration param.Field[NewBillingCycleConfigurationParam] `json:"billing_cycle_configuration"`
	// The per unit conversion rate of the price currency to the invoicing currency.
	ConversionRate param.Field[float64] `json:"conversion_rate"`
	// The configuration for the rate of the price currency to the invoicing currency.
	ConversionRateConfig param.Field[NewFloatingGroupedWithProratedMinimumPriceUnionParam] `json:"conversion_rate_config"`
	// For dimensional price: specifies a price group and dimension values
	DimensionalPriceConfiguration param.Field[NewDimensionalPriceConfigurationParam] `json:"dimensional_price_configuration"`
	// An alias for the price.
	ExternalPriceID param.Field[string] `json:"external_price_id"`
	// If the Price represents a fixed cost, this represents the quantity of units
	// applied.
	FixedPriceQuantity param.Field[float64] `json:"fixed_price_quantity"`
	// The property used to group this price on an invoice
	InvoiceGroupingKey param.Field[string] `json:"invoice_grouping_key"`
	// Within each billing cycle, specifies the cadence at which invoices are produced.
	// If unspecified, a single invoice is produced per billing cycle.
	InvoicingCycleConfiguration param.Field[NewBillingCycleConfigurationParam] `json:"invoicing_cycle_configuration"`
	// User-specified key/value pairs for the resource. Individual keys can be removed
	// by setting the value to `null`, and the entire metadata mapping can be cleared
	// by setting `metadata` to `null`.
	Metadata param.Field[map[string]string] `json:"metadata"`
}

func (r NewFloatingGroupedWithProratedMinimumPriceParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r NewFloatingGroupedWithProratedMinimumPriceParam) ImplementsPriceEvaluateMultipleParamsPriceEvaluationsPriceUnion() {
}

func (r NewFloatingGroupedWithProratedMinimumPriceParam) ImplementsPriceEvaluatePreviewEventsParamsPriceEvaluationsPriceUnion() {
}

func (r NewFloatingGroupedWithProratedMinimumPriceParam) ImplementsSubscriptionPriceIntervalsParamsAddPriceUnion() {
}

// The cadence to bill for this price on.
type NewFloatingGroupedWithProratedMinimumPriceCadence string

const (
	NewFloatingGroupedWithProratedMinimumPriceCadenceAnnual     NewFloatingGroupedWithProratedMinimumPriceCadence = "annual"
	NewFloatingGroupedWithProratedMinimumPriceCadenceSemiAnnual NewFloatingGroupedWithProratedMinimumPriceCadence = "semi_annual"
	NewFloatingGroupedWithProratedMinimumPriceCadenceMonthly    NewFloatingGroupedWithProratedMinimumPriceCadence = "monthly"
	NewFloatingGroupedWithProratedMinimumPriceCadenceQuarterly  NewFloatingGroupedWithProratedMinimumPriceCadence = "quarterly"
	NewFloatingGroupedWithProratedMinimumPriceCadenceOneTime    NewFloatingGroupedWithProratedMinimumPriceCadence = "one_time"
	NewFloatingGroupedWithProratedMinimumPriceCadenceCustom     NewFloatingGroupedWithProratedMinimumPriceCadence = "custom"
)

func (r NewFloatingGroupedWithProratedMinimumPriceCadence) IsKnown() bool {
	switch r {
	case NewFloatingGroupedWithProratedMinimumPriceCadenceAnnual, NewFloatingGroupedWithProratedMinimumPriceCadenceSemiAnnual, NewFloatingGroupedWithProratedMinimumPriceCadenceMonthly, NewFloatingGroupedWithProratedMinimumPriceCadenceQuarterly, NewFloatingGroupedWithProratedMinimumPriceCadenceOneTime, NewFloatingGroupedWithProratedMinimumPriceCadenceCustom:
		return true
	}
	return false
}

type NewFloatingGroupedWithProratedMinimumPriceModelType string

const (
	NewFloatingGroupedWithProratedMinimumPriceModelTypeGroupedWithProratedMinimum NewFloatingGroupedWithProratedMinimumPriceModelType = "grouped_with_prorated_minimum"
)

func (r NewFloatingGroupedWithProratedMinimumPriceModelType) IsKnown() bool {
	switch r {
	case NewFloatingGroupedWithProratedMinimumPriceModelTypeGroupedWithProratedMinimum:
		return true
	}
	return false
}

type NewFloatingGroupedWithProratedMinimumPriceParam struct {
	ConversionRateType param.Field[NewFloatingGroupedWithProratedMinimumPriceConversionRateType] `json:"conversion_rate_type,required"`
	TieredConfig       param.Field[ConversionRateTieredConfigParam]                              `json:"tiered_config"`
	UnitConfig         param.Field[ConversionRateUnitConfigParam]                                `json:"unit_config"`
}

func (r NewFloatingGroupedWithProratedMinimumPriceParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r NewFloatingGroupedWithProratedMinimumPriceParam) ImplementsNewFloatingGroupedWithProratedMinimumPriceUnionParam() {
}

// Satisfied by [shared.UnitConversionRateConfigParam],
// [shared.TieredConversionRateConfigParam],
// [NewFloatingGroupedWithProratedMinimumPriceParam].
type NewFloatingGroupedWithProratedMinimumPriceUnionParam interface {
	ImplementsNewFloatingGroupedWithProratedMinimumPriceUnionParam()
}

type NewFloatingGroupedWithProratedMinimumPriceConversionRateType string

const (
	NewFloatingGroupedWithProratedMinimumPriceConversionRateTypeUnit   NewFloatingGroupedWithProratedMinimumPriceConversionRateType = "unit"
	NewFloatingGroupedWithProratedMinimumPriceConversionRateTypeTiered NewFloatingGroupedWithProratedMinimumPriceConversionRateType = "tiered"
)

func (r NewFloatingGroupedWithProratedMinimumPriceConversionRateType) IsKnown() bool {
	switch r {
	case NewFloatingGroupedWithProratedMinimumPriceConversionRateTypeUnit, NewFloatingGroupedWithProratedMinimumPriceConversionRateTypeTiered:
		return true
	}
	return false
}

type NewFloatingMatrixPriceParam struct {
	// The cadence to bill for this price on.
	Cadence param.Field[NewFloatingMatrixPriceCadence] `json:"cadence,required"`
	// An ISO 4217 currency string for which this price is billed in.
	Currency param.Field[string] `json:"currency,required"`
	// The id of the item the price will be associated with.
	ItemID       param.Field[string]                          `json:"item_id,required"`
	MatrixConfig param.Field[MatrixConfigParam]               `json:"matrix_config,required"`
	ModelType    param.Field[NewFloatingMatrixPriceModelType] `json:"model_type,required"`
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
	BillingCycleConfiguration param.Field[NewBillingCycleConfigurationParam] `json:"billing_cycle_configuration"`
	// The per unit conversion rate of the price currency to the invoicing currency.
	ConversionRate param.Field[float64] `json:"conversion_rate"`
	// The configuration for the rate of the price currency to the invoicing currency.
	ConversionRateConfig param.Field[NewFloatingMatrixPriceUnionParam] `json:"conversion_rate_config"`
	// For dimensional price: specifies a price group and dimension values
	DimensionalPriceConfiguration param.Field[NewDimensionalPriceConfigurationParam] `json:"dimensional_price_configuration"`
	// An alias for the price.
	ExternalPriceID param.Field[string] `json:"external_price_id"`
	// If the Price represents a fixed cost, this represents the quantity of units
	// applied.
	FixedPriceQuantity param.Field[float64] `json:"fixed_price_quantity"`
	// The property used to group this price on an invoice
	InvoiceGroupingKey param.Field[string] `json:"invoice_grouping_key"`
	// Within each billing cycle, specifies the cadence at which invoices are produced.
	// If unspecified, a single invoice is produced per billing cycle.
	InvoicingCycleConfiguration param.Field[NewBillingCycleConfigurationParam] `json:"invoicing_cycle_configuration"`
	// User-specified key/value pairs for the resource. Individual keys can be removed
	// by setting the value to `null`, and the entire metadata mapping can be cleared
	// by setting `metadata` to `null`.
	Metadata param.Field[map[string]string] `json:"metadata"`
}

func (r NewFloatingMatrixPriceParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r NewFloatingMatrixPriceParam) ImplementsPriceEvaluateMultipleParamsPriceEvaluationsPriceUnion() {
}

func (r NewFloatingMatrixPriceParam) ImplementsPriceEvaluatePreviewEventsParamsPriceEvaluationsPriceUnion() {
}

func (r NewFloatingMatrixPriceParam) ImplementsSubscriptionPriceIntervalsParamsAddPriceUnion() {}

// The cadence to bill for this price on.
type NewFloatingMatrixPriceCadence string

const (
	NewFloatingMatrixPriceCadenceAnnual     NewFloatingMatrixPriceCadence = "annual"
	NewFloatingMatrixPriceCadenceSemiAnnual NewFloatingMatrixPriceCadence = "semi_annual"
	NewFloatingMatrixPriceCadenceMonthly    NewFloatingMatrixPriceCadence = "monthly"
	NewFloatingMatrixPriceCadenceQuarterly  NewFloatingMatrixPriceCadence = "quarterly"
	NewFloatingMatrixPriceCadenceOneTime    NewFloatingMatrixPriceCadence = "one_time"
	NewFloatingMatrixPriceCadenceCustom     NewFloatingMatrixPriceCadence = "custom"
)

func (r NewFloatingMatrixPriceCadence) IsKnown() bool {
	switch r {
	case NewFloatingMatrixPriceCadenceAnnual, NewFloatingMatrixPriceCadenceSemiAnnual, NewFloatingMatrixPriceCadenceMonthly, NewFloatingMatrixPriceCadenceQuarterly, NewFloatingMatrixPriceCadenceOneTime, NewFloatingMatrixPriceCadenceCustom:
		return true
	}
	return false
}

type NewFloatingMatrixPriceModelType string

const (
	NewFloatingMatrixPriceModelTypeMatrix NewFloatingMatrixPriceModelType = "matrix"
)

func (r NewFloatingMatrixPriceModelType) IsKnown() bool {
	switch r {
	case NewFloatingMatrixPriceModelTypeMatrix:
		return true
	}
	return false
}

type NewFloatingMatrixPriceParam struct {
	ConversionRateType param.Field[NewFloatingMatrixPriceConversionRateType] `json:"conversion_rate_type,required"`
	TieredConfig       param.Field[ConversionRateTieredConfigParam]          `json:"tiered_config"`
	UnitConfig         param.Field[ConversionRateUnitConfigParam]            `json:"unit_config"`
}

func (r NewFloatingMatrixPriceParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r NewFloatingMatrixPriceParam) ImplementsNewFloatingMatrixPriceUnionParam() {}

// Satisfied by [shared.UnitConversionRateConfigParam],
// [shared.TieredConversionRateConfigParam], [NewFloatingMatrixPriceParam].
type NewFloatingMatrixPriceUnionParam interface {
	ImplementsNewFloatingMatrixPriceUnionParam()
}

type NewFloatingMatrixPriceConversionRateType string

const (
	NewFloatingMatrixPriceConversionRateTypeUnit   NewFloatingMatrixPriceConversionRateType = "unit"
	NewFloatingMatrixPriceConversionRateTypeTiered NewFloatingMatrixPriceConversionRateType = "tiered"
)

func (r NewFloatingMatrixPriceConversionRateType) IsKnown() bool {
	switch r {
	case NewFloatingMatrixPriceConversionRateTypeUnit, NewFloatingMatrixPriceConversionRateTypeTiered:
		return true
	}
	return false
}

type NewFloatingMatrixWithAllocationPriceParam struct {
	// The cadence to bill for this price on.
	Cadence param.Field[NewFloatingMatrixWithAllocationPriceCadence] `json:"cadence,required"`
	// An ISO 4217 currency string for which this price is billed in.
	Currency param.Field[string] `json:"currency,required"`
	// The id of the item the price will be associated with.
	ItemID                     param.Field[string]                                        `json:"item_id,required"`
	MatrixWithAllocationConfig param.Field[MatrixWithAllocationConfigParam]               `json:"matrix_with_allocation_config,required"`
	ModelType                  param.Field[NewFloatingMatrixWithAllocationPriceModelType] `json:"model_type,required"`
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
	BillingCycleConfiguration param.Field[NewBillingCycleConfigurationParam] `json:"billing_cycle_configuration"`
	// The per unit conversion rate of the price currency to the invoicing currency.
	ConversionRate param.Field[float64] `json:"conversion_rate"`
	// The configuration for the rate of the price currency to the invoicing currency.
	ConversionRateConfig param.Field[NewFloatingMatrixWithAllocationPriceUnionParam] `json:"conversion_rate_config"`
	// For dimensional price: specifies a price group and dimension values
	DimensionalPriceConfiguration param.Field[NewDimensionalPriceConfigurationParam] `json:"dimensional_price_configuration"`
	// An alias for the price.
	ExternalPriceID param.Field[string] `json:"external_price_id"`
	// If the Price represents a fixed cost, this represents the quantity of units
	// applied.
	FixedPriceQuantity param.Field[float64] `json:"fixed_price_quantity"`
	// The property used to group this price on an invoice
	InvoiceGroupingKey param.Field[string] `json:"invoice_grouping_key"`
	// Within each billing cycle, specifies the cadence at which invoices are produced.
	// If unspecified, a single invoice is produced per billing cycle.
	InvoicingCycleConfiguration param.Field[NewBillingCycleConfigurationParam] `json:"invoicing_cycle_configuration"`
	// User-specified key/value pairs for the resource. Individual keys can be removed
	// by setting the value to `null`, and the entire metadata mapping can be cleared
	// by setting `metadata` to `null`.
	Metadata param.Field[map[string]string] `json:"metadata"`
}

func (r NewFloatingMatrixWithAllocationPriceParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r NewFloatingMatrixWithAllocationPriceParam) ImplementsPriceEvaluateMultipleParamsPriceEvaluationsPriceUnion() {
}

func (r NewFloatingMatrixWithAllocationPriceParam) ImplementsPriceEvaluatePreviewEventsParamsPriceEvaluationsPriceUnion() {
}

func (r NewFloatingMatrixWithAllocationPriceParam) ImplementsSubscriptionPriceIntervalsParamsAddPriceUnion() {
}

// The cadence to bill for this price on.
type NewFloatingMatrixWithAllocationPriceCadence string

const (
	NewFloatingMatrixWithAllocationPriceCadenceAnnual     NewFloatingMatrixWithAllocationPriceCadence = "annual"
	NewFloatingMatrixWithAllocationPriceCadenceSemiAnnual NewFloatingMatrixWithAllocationPriceCadence = "semi_annual"
	NewFloatingMatrixWithAllocationPriceCadenceMonthly    NewFloatingMatrixWithAllocationPriceCadence = "monthly"
	NewFloatingMatrixWithAllocationPriceCadenceQuarterly  NewFloatingMatrixWithAllocationPriceCadence = "quarterly"
	NewFloatingMatrixWithAllocationPriceCadenceOneTime    NewFloatingMatrixWithAllocationPriceCadence = "one_time"
	NewFloatingMatrixWithAllocationPriceCadenceCustom     NewFloatingMatrixWithAllocationPriceCadence = "custom"
)

func (r NewFloatingMatrixWithAllocationPriceCadence) IsKnown() bool {
	switch r {
	case NewFloatingMatrixWithAllocationPriceCadenceAnnual, NewFloatingMatrixWithAllocationPriceCadenceSemiAnnual, NewFloatingMatrixWithAllocationPriceCadenceMonthly, NewFloatingMatrixWithAllocationPriceCadenceQuarterly, NewFloatingMatrixWithAllocationPriceCadenceOneTime, NewFloatingMatrixWithAllocationPriceCadenceCustom:
		return true
	}
	return false
}

type NewFloatingMatrixWithAllocationPriceModelType string

const (
	NewFloatingMatrixWithAllocationPriceModelTypeMatrixWithAllocation NewFloatingMatrixWithAllocationPriceModelType = "matrix_with_allocation"
)

func (r NewFloatingMatrixWithAllocationPriceModelType) IsKnown() bool {
	switch r {
	case NewFloatingMatrixWithAllocationPriceModelTypeMatrixWithAllocation:
		return true
	}
	return false
}

type NewFloatingMatrixWithAllocationPriceParam struct {
	ConversionRateType param.Field[NewFloatingMatrixWithAllocationPriceConversionRateType] `json:"conversion_rate_type,required"`
	TieredConfig       param.Field[ConversionRateTieredConfigParam]                        `json:"tiered_config"`
	UnitConfig         param.Field[ConversionRateUnitConfigParam]                          `json:"unit_config"`
}

func (r NewFloatingMatrixWithAllocationPriceParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r NewFloatingMatrixWithAllocationPriceParam) ImplementsNewFloatingMatrixWithAllocationPriceUnionParam() {
}

// Satisfied by [shared.UnitConversionRateConfigParam],
// [shared.TieredConversionRateConfigParam],
// [NewFloatingMatrixWithAllocationPriceParam].
type NewFloatingMatrixWithAllocationPriceUnionParam interface {
	ImplementsNewFloatingMatrixWithAllocationPriceUnionParam()
}

type NewFloatingMatrixWithAllocationPriceConversionRateType string

const (
	NewFloatingMatrixWithAllocationPriceConversionRateTypeUnit   NewFloatingMatrixWithAllocationPriceConversionRateType = "unit"
	NewFloatingMatrixWithAllocationPriceConversionRateTypeTiered NewFloatingMatrixWithAllocationPriceConversionRateType = "tiered"
)

func (r NewFloatingMatrixWithAllocationPriceConversionRateType) IsKnown() bool {
	switch r {
	case NewFloatingMatrixWithAllocationPriceConversionRateTypeUnit, NewFloatingMatrixWithAllocationPriceConversionRateTypeTiered:
		return true
	}
	return false
}

type NewFloatingMatrixWithDisplayNamePriceParam struct {
	// The cadence to bill for this price on.
	Cadence param.Field[NewFloatingMatrixWithDisplayNamePriceCadence] `json:"cadence,required"`
	// An ISO 4217 currency string for which this price is billed in.
	Currency param.Field[string] `json:"currency,required"`
	// The id of the item the price will be associated with.
	ItemID                      param.Field[string]                                         `json:"item_id,required"`
	MatrixWithDisplayNameConfig param.Field[map[string]interface{}]                         `json:"matrix_with_display_name_config,required"`
	ModelType                   param.Field[NewFloatingMatrixWithDisplayNamePriceModelType] `json:"model_type,required"`
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
	BillingCycleConfiguration param.Field[NewBillingCycleConfigurationParam] `json:"billing_cycle_configuration"`
	// The per unit conversion rate of the price currency to the invoicing currency.
	ConversionRate param.Field[float64] `json:"conversion_rate"`
	// The configuration for the rate of the price currency to the invoicing currency.
	ConversionRateConfig param.Field[NewFloatingMatrixWithDisplayNamePriceUnionParam] `json:"conversion_rate_config"`
	// For dimensional price: specifies a price group and dimension values
	DimensionalPriceConfiguration param.Field[NewDimensionalPriceConfigurationParam] `json:"dimensional_price_configuration"`
	// An alias for the price.
	ExternalPriceID param.Field[string] `json:"external_price_id"`
	// If the Price represents a fixed cost, this represents the quantity of units
	// applied.
	FixedPriceQuantity param.Field[float64] `json:"fixed_price_quantity"`
	// The property used to group this price on an invoice
	InvoiceGroupingKey param.Field[string] `json:"invoice_grouping_key"`
	// Within each billing cycle, specifies the cadence at which invoices are produced.
	// If unspecified, a single invoice is produced per billing cycle.
	InvoicingCycleConfiguration param.Field[NewBillingCycleConfigurationParam] `json:"invoicing_cycle_configuration"`
	// User-specified key/value pairs for the resource. Individual keys can be removed
	// by setting the value to `null`, and the entire metadata mapping can be cleared
	// by setting `metadata` to `null`.
	Metadata param.Field[map[string]string] `json:"metadata"`
}

func (r NewFloatingMatrixWithDisplayNamePriceParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r NewFloatingMatrixWithDisplayNamePriceParam) ImplementsPriceEvaluateMultipleParamsPriceEvaluationsPriceUnion() {
}

func (r NewFloatingMatrixWithDisplayNamePriceParam) ImplementsPriceEvaluatePreviewEventsParamsPriceEvaluationsPriceUnion() {
}

func (r NewFloatingMatrixWithDisplayNamePriceParam) ImplementsSubscriptionPriceIntervalsParamsAddPriceUnion() {
}

// The cadence to bill for this price on.
type NewFloatingMatrixWithDisplayNamePriceCadence string

const (
	NewFloatingMatrixWithDisplayNamePriceCadenceAnnual     NewFloatingMatrixWithDisplayNamePriceCadence = "annual"
	NewFloatingMatrixWithDisplayNamePriceCadenceSemiAnnual NewFloatingMatrixWithDisplayNamePriceCadence = "semi_annual"
	NewFloatingMatrixWithDisplayNamePriceCadenceMonthly    NewFloatingMatrixWithDisplayNamePriceCadence = "monthly"
	NewFloatingMatrixWithDisplayNamePriceCadenceQuarterly  NewFloatingMatrixWithDisplayNamePriceCadence = "quarterly"
	NewFloatingMatrixWithDisplayNamePriceCadenceOneTime    NewFloatingMatrixWithDisplayNamePriceCadence = "one_time"
	NewFloatingMatrixWithDisplayNamePriceCadenceCustom     NewFloatingMatrixWithDisplayNamePriceCadence = "custom"
)

func (r NewFloatingMatrixWithDisplayNamePriceCadence) IsKnown() bool {
	switch r {
	case NewFloatingMatrixWithDisplayNamePriceCadenceAnnual, NewFloatingMatrixWithDisplayNamePriceCadenceSemiAnnual, NewFloatingMatrixWithDisplayNamePriceCadenceMonthly, NewFloatingMatrixWithDisplayNamePriceCadenceQuarterly, NewFloatingMatrixWithDisplayNamePriceCadenceOneTime, NewFloatingMatrixWithDisplayNamePriceCadenceCustom:
		return true
	}
	return false
}

type NewFloatingMatrixWithDisplayNamePriceModelType string

const (
	NewFloatingMatrixWithDisplayNamePriceModelTypeMatrixWithDisplayName NewFloatingMatrixWithDisplayNamePriceModelType = "matrix_with_display_name"
)

func (r NewFloatingMatrixWithDisplayNamePriceModelType) IsKnown() bool {
	switch r {
	case NewFloatingMatrixWithDisplayNamePriceModelTypeMatrixWithDisplayName:
		return true
	}
	return false
}

type NewFloatingMatrixWithDisplayNamePriceParam struct {
	ConversionRateType param.Field[NewFloatingMatrixWithDisplayNamePriceConversionRateType] `json:"conversion_rate_type,required"`
	TieredConfig       param.Field[ConversionRateTieredConfigParam]                         `json:"tiered_config"`
	UnitConfig         param.Field[ConversionRateUnitConfigParam]                           `json:"unit_config"`
}

func (r NewFloatingMatrixWithDisplayNamePriceParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r NewFloatingMatrixWithDisplayNamePriceParam) ImplementsNewFloatingMatrixWithDisplayNamePriceUnionParam() {
}

// Satisfied by [shared.UnitConversionRateConfigParam],
// [shared.TieredConversionRateConfigParam],
// [NewFloatingMatrixWithDisplayNamePriceParam].
type NewFloatingMatrixWithDisplayNamePriceUnionParam interface {
	ImplementsNewFloatingMatrixWithDisplayNamePriceUnionParam()
}

type NewFloatingMatrixWithDisplayNamePriceConversionRateType string

const (
	NewFloatingMatrixWithDisplayNamePriceConversionRateTypeUnit   NewFloatingMatrixWithDisplayNamePriceConversionRateType = "unit"
	NewFloatingMatrixWithDisplayNamePriceConversionRateTypeTiered NewFloatingMatrixWithDisplayNamePriceConversionRateType = "tiered"
)

func (r NewFloatingMatrixWithDisplayNamePriceConversionRateType) IsKnown() bool {
	switch r {
	case NewFloatingMatrixWithDisplayNamePriceConversionRateTypeUnit, NewFloatingMatrixWithDisplayNamePriceConversionRateTypeTiered:
		return true
	}
	return false
}

type NewFloatingMaxGroupTieredPackagePriceParam struct {
	// The cadence to bill for this price on.
	Cadence param.Field[NewFloatingMaxGroupTieredPackagePriceCadence] `json:"cadence,required"`
	// An ISO 4217 currency string for which this price is billed in.
	Currency param.Field[string] `json:"currency,required"`
	// The id of the item the price will be associated with.
	ItemID                      param.Field[string]                                         `json:"item_id,required"`
	MaxGroupTieredPackageConfig param.Field[map[string]interface{}]                         `json:"max_group_tiered_package_config,required"`
	ModelType                   param.Field[NewFloatingMaxGroupTieredPackagePriceModelType] `json:"model_type,required"`
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
	BillingCycleConfiguration param.Field[NewBillingCycleConfigurationParam] `json:"billing_cycle_configuration"`
	// The per unit conversion rate of the price currency to the invoicing currency.
	ConversionRate param.Field[float64] `json:"conversion_rate"`
	// The configuration for the rate of the price currency to the invoicing currency.
	ConversionRateConfig param.Field[NewFloatingMaxGroupTieredPackagePriceUnionParam] `json:"conversion_rate_config"`
	// For dimensional price: specifies a price group and dimension values
	DimensionalPriceConfiguration param.Field[NewDimensionalPriceConfigurationParam] `json:"dimensional_price_configuration"`
	// An alias for the price.
	ExternalPriceID param.Field[string] `json:"external_price_id"`
	// If the Price represents a fixed cost, this represents the quantity of units
	// applied.
	FixedPriceQuantity param.Field[float64] `json:"fixed_price_quantity"`
	// The property used to group this price on an invoice
	InvoiceGroupingKey param.Field[string] `json:"invoice_grouping_key"`
	// Within each billing cycle, specifies the cadence at which invoices are produced.
	// If unspecified, a single invoice is produced per billing cycle.
	InvoicingCycleConfiguration param.Field[NewBillingCycleConfigurationParam] `json:"invoicing_cycle_configuration"`
	// User-specified key/value pairs for the resource. Individual keys can be removed
	// by setting the value to `null`, and the entire metadata mapping can be cleared
	// by setting `metadata` to `null`.
	Metadata param.Field[map[string]string] `json:"metadata"`
}

func (r NewFloatingMaxGroupTieredPackagePriceParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r NewFloatingMaxGroupTieredPackagePriceParam) ImplementsPriceEvaluateMultipleParamsPriceEvaluationsPriceUnion() {
}

func (r NewFloatingMaxGroupTieredPackagePriceParam) ImplementsPriceEvaluatePreviewEventsParamsPriceEvaluationsPriceUnion() {
}

func (r NewFloatingMaxGroupTieredPackagePriceParam) ImplementsSubscriptionPriceIntervalsParamsAddPriceUnion() {
}

// The cadence to bill for this price on.
type NewFloatingMaxGroupTieredPackagePriceCadence string

const (
	NewFloatingMaxGroupTieredPackagePriceCadenceAnnual     NewFloatingMaxGroupTieredPackagePriceCadence = "annual"
	NewFloatingMaxGroupTieredPackagePriceCadenceSemiAnnual NewFloatingMaxGroupTieredPackagePriceCadence = "semi_annual"
	NewFloatingMaxGroupTieredPackagePriceCadenceMonthly    NewFloatingMaxGroupTieredPackagePriceCadence = "monthly"
	NewFloatingMaxGroupTieredPackagePriceCadenceQuarterly  NewFloatingMaxGroupTieredPackagePriceCadence = "quarterly"
	NewFloatingMaxGroupTieredPackagePriceCadenceOneTime    NewFloatingMaxGroupTieredPackagePriceCadence = "one_time"
	NewFloatingMaxGroupTieredPackagePriceCadenceCustom     NewFloatingMaxGroupTieredPackagePriceCadence = "custom"
)

func (r NewFloatingMaxGroupTieredPackagePriceCadence) IsKnown() bool {
	switch r {
	case NewFloatingMaxGroupTieredPackagePriceCadenceAnnual, NewFloatingMaxGroupTieredPackagePriceCadenceSemiAnnual, NewFloatingMaxGroupTieredPackagePriceCadenceMonthly, NewFloatingMaxGroupTieredPackagePriceCadenceQuarterly, NewFloatingMaxGroupTieredPackagePriceCadenceOneTime, NewFloatingMaxGroupTieredPackagePriceCadenceCustom:
		return true
	}
	return false
}

type NewFloatingMaxGroupTieredPackagePriceModelType string

const (
	NewFloatingMaxGroupTieredPackagePriceModelTypeMaxGroupTieredPackage NewFloatingMaxGroupTieredPackagePriceModelType = "max_group_tiered_package"
)

func (r NewFloatingMaxGroupTieredPackagePriceModelType) IsKnown() bool {
	switch r {
	case NewFloatingMaxGroupTieredPackagePriceModelTypeMaxGroupTieredPackage:
		return true
	}
	return false
}

type NewFloatingMaxGroupTieredPackagePriceParam struct {
	ConversionRateType param.Field[NewFloatingMaxGroupTieredPackagePriceConversionRateType] `json:"conversion_rate_type,required"`
	TieredConfig       param.Field[ConversionRateTieredConfigParam]                         `json:"tiered_config"`
	UnitConfig         param.Field[ConversionRateUnitConfigParam]                           `json:"unit_config"`
}

func (r NewFloatingMaxGroupTieredPackagePriceParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r NewFloatingMaxGroupTieredPackagePriceParam) ImplementsNewFloatingMaxGroupTieredPackagePriceUnionParam() {
}

// Satisfied by [shared.UnitConversionRateConfigParam],
// [shared.TieredConversionRateConfigParam],
// [NewFloatingMaxGroupTieredPackagePriceParam].
type NewFloatingMaxGroupTieredPackagePriceUnionParam interface {
	ImplementsNewFloatingMaxGroupTieredPackagePriceUnionParam()
}

type NewFloatingMaxGroupTieredPackagePriceConversionRateType string

const (
	NewFloatingMaxGroupTieredPackagePriceConversionRateTypeUnit   NewFloatingMaxGroupTieredPackagePriceConversionRateType = "unit"
	NewFloatingMaxGroupTieredPackagePriceConversionRateTypeTiered NewFloatingMaxGroupTieredPackagePriceConversionRateType = "tiered"
)

func (r NewFloatingMaxGroupTieredPackagePriceConversionRateType) IsKnown() bool {
	switch r {
	case NewFloatingMaxGroupTieredPackagePriceConversionRateTypeUnit, NewFloatingMaxGroupTieredPackagePriceConversionRateTypeTiered:
		return true
	}
	return false
}

type NewFloatingPackagePriceParam struct {
	// The cadence to bill for this price on.
	Cadence param.Field[NewFloatingPackagePriceCadence] `json:"cadence,required"`
	// An ISO 4217 currency string for which this price is billed in.
	Currency param.Field[string] `json:"currency,required"`
	// The id of the item the price will be associated with.
	ItemID    param.Field[string]                           `json:"item_id,required"`
	ModelType param.Field[NewFloatingPackagePriceModelType] `json:"model_type,required"`
	// The name of the price.
	Name          param.Field[string]             `json:"name,required"`
	PackageConfig param.Field[PackageConfigParam] `json:"package_config,required"`
	// The id of the billable metric for the price. Only needed if the price is
	// usage-based.
	BillableMetricID param.Field[string] `json:"billable_metric_id"`
	// If the Price represents a fixed cost, the price will be billed in-advance if
	// this is true, and in-arrears if this is false.
	BilledInAdvance param.Field[bool] `json:"billed_in_advance"`
	// For custom cadence: specifies the duration of the billing period in days or
	// months.
	BillingCycleConfiguration param.Field[NewBillingCycleConfigurationParam] `json:"billing_cycle_configuration"`
	// The per unit conversion rate of the price currency to the invoicing currency.
	ConversionRate param.Field[float64] `json:"conversion_rate"`
	// The configuration for the rate of the price currency to the invoicing currency.
	ConversionRateConfig param.Field[NewFloatingPackagePriceUnionParam] `json:"conversion_rate_config"`
	// For dimensional price: specifies a price group and dimension values
	DimensionalPriceConfiguration param.Field[NewDimensionalPriceConfigurationParam] `json:"dimensional_price_configuration"`
	// An alias for the price.
	ExternalPriceID param.Field[string] `json:"external_price_id"`
	// If the Price represents a fixed cost, this represents the quantity of units
	// applied.
	FixedPriceQuantity param.Field[float64] `json:"fixed_price_quantity"`
	// The property used to group this price on an invoice
	InvoiceGroupingKey param.Field[string] `json:"invoice_grouping_key"`
	// Within each billing cycle, specifies the cadence at which invoices are produced.
	// If unspecified, a single invoice is produced per billing cycle.
	InvoicingCycleConfiguration param.Field[NewBillingCycleConfigurationParam] `json:"invoicing_cycle_configuration"`
	// User-specified key/value pairs for the resource. Individual keys can be removed
	// by setting the value to `null`, and the entire metadata mapping can be cleared
	// by setting `metadata` to `null`.
	Metadata param.Field[map[string]string] `json:"metadata"`
}

func (r NewFloatingPackagePriceParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r NewFloatingPackagePriceParam) ImplementsPriceEvaluateMultipleParamsPriceEvaluationsPriceUnion() {
}

func (r NewFloatingPackagePriceParam) ImplementsPriceEvaluatePreviewEventsParamsPriceEvaluationsPriceUnion() {
}

func (r NewFloatingPackagePriceParam) ImplementsSubscriptionPriceIntervalsParamsAddPriceUnion() {}

// The cadence to bill for this price on.
type NewFloatingPackagePriceCadence string

const (
	NewFloatingPackagePriceCadenceAnnual     NewFloatingPackagePriceCadence = "annual"
	NewFloatingPackagePriceCadenceSemiAnnual NewFloatingPackagePriceCadence = "semi_annual"
	NewFloatingPackagePriceCadenceMonthly    NewFloatingPackagePriceCadence = "monthly"
	NewFloatingPackagePriceCadenceQuarterly  NewFloatingPackagePriceCadence = "quarterly"
	NewFloatingPackagePriceCadenceOneTime    NewFloatingPackagePriceCadence = "one_time"
	NewFloatingPackagePriceCadenceCustom     NewFloatingPackagePriceCadence = "custom"
)

func (r NewFloatingPackagePriceCadence) IsKnown() bool {
	switch r {
	case NewFloatingPackagePriceCadenceAnnual, NewFloatingPackagePriceCadenceSemiAnnual, NewFloatingPackagePriceCadenceMonthly, NewFloatingPackagePriceCadenceQuarterly, NewFloatingPackagePriceCadenceOneTime, NewFloatingPackagePriceCadenceCustom:
		return true
	}
	return false
}

type NewFloatingPackagePriceModelType string

const (
	NewFloatingPackagePriceModelTypePackage NewFloatingPackagePriceModelType = "package"
)

func (r NewFloatingPackagePriceModelType) IsKnown() bool {
	switch r {
	case NewFloatingPackagePriceModelTypePackage:
		return true
	}
	return false
}

type NewFloatingPackagePriceParam struct {
	ConversionRateType param.Field[NewFloatingPackagePriceConversionRateType] `json:"conversion_rate_type,required"`
	TieredConfig       param.Field[ConversionRateTieredConfigParam]           `json:"tiered_config"`
	UnitConfig         param.Field[ConversionRateUnitConfigParam]             `json:"unit_config"`
}

func (r NewFloatingPackagePriceParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r NewFloatingPackagePriceParam) ImplementsNewFloatingPackagePriceUnionParam() {}

// Satisfied by [shared.UnitConversionRateConfigParam],
// [shared.TieredConversionRateConfigParam], [NewFloatingPackagePriceParam].
type NewFloatingPackagePriceUnionParam interface {
	ImplementsNewFloatingPackagePriceUnionParam()
}

type NewFloatingPackagePriceConversionRateType string

const (
	NewFloatingPackagePriceConversionRateTypeUnit   NewFloatingPackagePriceConversionRateType = "unit"
	NewFloatingPackagePriceConversionRateTypeTiered NewFloatingPackagePriceConversionRateType = "tiered"
)

func (r NewFloatingPackagePriceConversionRateType) IsKnown() bool {
	switch r {
	case NewFloatingPackagePriceConversionRateTypeUnit, NewFloatingPackagePriceConversionRateTypeTiered:
		return true
	}
	return false
}

type NewFloatingPackageWithAllocationPriceParam struct {
	// The cadence to bill for this price on.
	Cadence param.Field[NewFloatingPackageWithAllocationPriceCadence] `json:"cadence,required"`
	// An ISO 4217 currency string for which this price is billed in.
	Currency param.Field[string] `json:"currency,required"`
	// The id of the item the price will be associated with.
	ItemID    param.Field[string]                                         `json:"item_id,required"`
	ModelType param.Field[NewFloatingPackageWithAllocationPriceModelType] `json:"model_type,required"`
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
	BillingCycleConfiguration param.Field[NewBillingCycleConfigurationParam] `json:"billing_cycle_configuration"`
	// The per unit conversion rate of the price currency to the invoicing currency.
	ConversionRate param.Field[float64] `json:"conversion_rate"`
	// The configuration for the rate of the price currency to the invoicing currency.
	ConversionRateConfig param.Field[NewFloatingPackageWithAllocationPriceUnionParam] `json:"conversion_rate_config"`
	// For dimensional price: specifies a price group and dimension values
	DimensionalPriceConfiguration param.Field[NewDimensionalPriceConfigurationParam] `json:"dimensional_price_configuration"`
	// An alias for the price.
	ExternalPriceID param.Field[string] `json:"external_price_id"`
	// If the Price represents a fixed cost, this represents the quantity of units
	// applied.
	FixedPriceQuantity param.Field[float64] `json:"fixed_price_quantity"`
	// The property used to group this price on an invoice
	InvoiceGroupingKey param.Field[string] `json:"invoice_grouping_key"`
	// Within each billing cycle, specifies the cadence at which invoices are produced.
	// If unspecified, a single invoice is produced per billing cycle.
	InvoicingCycleConfiguration param.Field[NewBillingCycleConfigurationParam] `json:"invoicing_cycle_configuration"`
	// User-specified key/value pairs for the resource. Individual keys can be removed
	// by setting the value to `null`, and the entire metadata mapping can be cleared
	// by setting `metadata` to `null`.
	Metadata param.Field[map[string]string] `json:"metadata"`
}

func (r NewFloatingPackageWithAllocationPriceParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r NewFloatingPackageWithAllocationPriceParam) ImplementsPriceEvaluateMultipleParamsPriceEvaluationsPriceUnion() {
}

func (r NewFloatingPackageWithAllocationPriceParam) ImplementsPriceEvaluatePreviewEventsParamsPriceEvaluationsPriceUnion() {
}

func (r NewFloatingPackageWithAllocationPriceParam) ImplementsSubscriptionPriceIntervalsParamsAddPriceUnion() {
}

// The cadence to bill for this price on.
type NewFloatingPackageWithAllocationPriceCadence string

const (
	NewFloatingPackageWithAllocationPriceCadenceAnnual     NewFloatingPackageWithAllocationPriceCadence = "annual"
	NewFloatingPackageWithAllocationPriceCadenceSemiAnnual NewFloatingPackageWithAllocationPriceCadence = "semi_annual"
	NewFloatingPackageWithAllocationPriceCadenceMonthly    NewFloatingPackageWithAllocationPriceCadence = "monthly"
	NewFloatingPackageWithAllocationPriceCadenceQuarterly  NewFloatingPackageWithAllocationPriceCadence = "quarterly"
	NewFloatingPackageWithAllocationPriceCadenceOneTime    NewFloatingPackageWithAllocationPriceCadence = "one_time"
	NewFloatingPackageWithAllocationPriceCadenceCustom     NewFloatingPackageWithAllocationPriceCadence = "custom"
)

func (r NewFloatingPackageWithAllocationPriceCadence) IsKnown() bool {
	switch r {
	case NewFloatingPackageWithAllocationPriceCadenceAnnual, NewFloatingPackageWithAllocationPriceCadenceSemiAnnual, NewFloatingPackageWithAllocationPriceCadenceMonthly, NewFloatingPackageWithAllocationPriceCadenceQuarterly, NewFloatingPackageWithAllocationPriceCadenceOneTime, NewFloatingPackageWithAllocationPriceCadenceCustom:
		return true
	}
	return false
}

type NewFloatingPackageWithAllocationPriceModelType string

const (
	NewFloatingPackageWithAllocationPriceModelTypePackageWithAllocation NewFloatingPackageWithAllocationPriceModelType = "package_with_allocation"
)

func (r NewFloatingPackageWithAllocationPriceModelType) IsKnown() bool {
	switch r {
	case NewFloatingPackageWithAllocationPriceModelTypePackageWithAllocation:
		return true
	}
	return false
}

type NewFloatingPackageWithAllocationPriceParam struct {
	ConversionRateType param.Field[NewFloatingPackageWithAllocationPriceConversionRateType] `json:"conversion_rate_type,required"`
	TieredConfig       param.Field[ConversionRateTieredConfigParam]                         `json:"tiered_config"`
	UnitConfig         param.Field[ConversionRateUnitConfigParam]                           `json:"unit_config"`
}

func (r NewFloatingPackageWithAllocationPriceParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r NewFloatingPackageWithAllocationPriceParam) ImplementsNewFloatingPackageWithAllocationPriceUnionParam() {
}

// Satisfied by [shared.UnitConversionRateConfigParam],
// [shared.TieredConversionRateConfigParam],
// [NewFloatingPackageWithAllocationPriceParam].
type NewFloatingPackageWithAllocationPriceUnionParam interface {
	ImplementsNewFloatingPackageWithAllocationPriceUnionParam()
}

type NewFloatingPackageWithAllocationPriceConversionRateType string

const (
	NewFloatingPackageWithAllocationPriceConversionRateTypeUnit   NewFloatingPackageWithAllocationPriceConversionRateType = "unit"
	NewFloatingPackageWithAllocationPriceConversionRateTypeTiered NewFloatingPackageWithAllocationPriceConversionRateType = "tiered"
)

func (r NewFloatingPackageWithAllocationPriceConversionRateType) IsKnown() bool {
	switch r {
	case NewFloatingPackageWithAllocationPriceConversionRateTypeUnit, NewFloatingPackageWithAllocationPriceConversionRateTypeTiered:
		return true
	}
	return false
}

type NewFloatingScalableMatrixWithTieredPricingPriceParam struct {
	// The cadence to bill for this price on.
	Cadence param.Field[NewFloatingScalableMatrixWithTieredPricingPriceCadence] `json:"cadence,required"`
	// An ISO 4217 currency string for which this price is billed in.
	Currency param.Field[string] `json:"currency,required"`
	// The id of the item the price will be associated with.
	ItemID    param.Field[string]                                                   `json:"item_id,required"`
	ModelType param.Field[NewFloatingScalableMatrixWithTieredPricingPriceModelType] `json:"model_type,required"`
	// The name of the price.
	Name                                  param.Field[string]                 `json:"name,required"`
	ScalableMatrixWithTieredPricingConfig param.Field[map[string]interface{}] `json:"scalable_matrix_with_tiered_pricing_config,required"`
	// The id of the billable metric for the price. Only needed if the price is
	// usage-based.
	BillableMetricID param.Field[string] `json:"billable_metric_id"`
	// If the Price represents a fixed cost, the price will be billed in-advance if
	// this is true, and in-arrears if this is false.
	BilledInAdvance param.Field[bool] `json:"billed_in_advance"`
	// For custom cadence: specifies the duration of the billing period in days or
	// months.
	BillingCycleConfiguration param.Field[NewBillingCycleConfigurationParam] `json:"billing_cycle_configuration"`
	// The per unit conversion rate of the price currency to the invoicing currency.
	ConversionRate param.Field[float64] `json:"conversion_rate"`
	// The configuration for the rate of the price currency to the invoicing currency.
	ConversionRateConfig param.Field[NewFloatingScalableMatrixWithTieredPricingPriceUnionParam] `json:"conversion_rate_config"`
	// For dimensional price: specifies a price group and dimension values
	DimensionalPriceConfiguration param.Field[NewDimensionalPriceConfigurationParam] `json:"dimensional_price_configuration"`
	// An alias for the price.
	ExternalPriceID param.Field[string] `json:"external_price_id"`
	// If the Price represents a fixed cost, this represents the quantity of units
	// applied.
	FixedPriceQuantity param.Field[float64] `json:"fixed_price_quantity"`
	// The property used to group this price on an invoice
	InvoiceGroupingKey param.Field[string] `json:"invoice_grouping_key"`
	// Within each billing cycle, specifies the cadence at which invoices are produced.
	// If unspecified, a single invoice is produced per billing cycle.
	InvoicingCycleConfiguration param.Field[NewBillingCycleConfigurationParam] `json:"invoicing_cycle_configuration"`
	// User-specified key/value pairs for the resource. Individual keys can be removed
	// by setting the value to `null`, and the entire metadata mapping can be cleared
	// by setting `metadata` to `null`.
	Metadata param.Field[map[string]string] `json:"metadata"`
}

func (r NewFloatingScalableMatrixWithTieredPricingPriceParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r NewFloatingScalableMatrixWithTieredPricingPriceParam) ImplementsPriceEvaluateMultipleParamsPriceEvaluationsPriceUnion() {
}

func (r NewFloatingScalableMatrixWithTieredPricingPriceParam) ImplementsPriceEvaluatePreviewEventsParamsPriceEvaluationsPriceUnion() {
}

func (r NewFloatingScalableMatrixWithTieredPricingPriceParam) ImplementsSubscriptionPriceIntervalsParamsAddPriceUnion() {
}

// The cadence to bill for this price on.
type NewFloatingScalableMatrixWithTieredPricingPriceCadence string

const (
	NewFloatingScalableMatrixWithTieredPricingPriceCadenceAnnual     NewFloatingScalableMatrixWithTieredPricingPriceCadence = "annual"
	NewFloatingScalableMatrixWithTieredPricingPriceCadenceSemiAnnual NewFloatingScalableMatrixWithTieredPricingPriceCadence = "semi_annual"
	NewFloatingScalableMatrixWithTieredPricingPriceCadenceMonthly    NewFloatingScalableMatrixWithTieredPricingPriceCadence = "monthly"
	NewFloatingScalableMatrixWithTieredPricingPriceCadenceQuarterly  NewFloatingScalableMatrixWithTieredPricingPriceCadence = "quarterly"
	NewFloatingScalableMatrixWithTieredPricingPriceCadenceOneTime    NewFloatingScalableMatrixWithTieredPricingPriceCadence = "one_time"
	NewFloatingScalableMatrixWithTieredPricingPriceCadenceCustom     NewFloatingScalableMatrixWithTieredPricingPriceCadence = "custom"
)

func (r NewFloatingScalableMatrixWithTieredPricingPriceCadence) IsKnown() bool {
	switch r {
	case NewFloatingScalableMatrixWithTieredPricingPriceCadenceAnnual, NewFloatingScalableMatrixWithTieredPricingPriceCadenceSemiAnnual, NewFloatingScalableMatrixWithTieredPricingPriceCadenceMonthly, NewFloatingScalableMatrixWithTieredPricingPriceCadenceQuarterly, NewFloatingScalableMatrixWithTieredPricingPriceCadenceOneTime, NewFloatingScalableMatrixWithTieredPricingPriceCadenceCustom:
		return true
	}
	return false
}

type NewFloatingScalableMatrixWithTieredPricingPriceModelType string

const (
	NewFloatingScalableMatrixWithTieredPricingPriceModelTypeScalableMatrixWithTieredPricing NewFloatingScalableMatrixWithTieredPricingPriceModelType = "scalable_matrix_with_tiered_pricing"
)

func (r NewFloatingScalableMatrixWithTieredPricingPriceModelType) IsKnown() bool {
	switch r {
	case NewFloatingScalableMatrixWithTieredPricingPriceModelTypeScalableMatrixWithTieredPricing:
		return true
	}
	return false
}

type NewFloatingScalableMatrixWithTieredPricingPriceParam struct {
	ConversionRateType param.Field[NewFloatingScalableMatrixWithTieredPricingPriceConversionRateType] `json:"conversion_rate_type,required"`
	TieredConfig       param.Field[ConversionRateTieredConfigParam]                                   `json:"tiered_config"`
	UnitConfig         param.Field[ConversionRateUnitConfigParam]                                     `json:"unit_config"`
}

func (r NewFloatingScalableMatrixWithTieredPricingPriceParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r NewFloatingScalableMatrixWithTieredPricingPriceParam) ImplementsNewFloatingScalableMatrixWithTieredPricingPriceUnionParam() {
}

// Satisfied by [shared.UnitConversionRateConfigParam],
// [shared.TieredConversionRateConfigParam],
// [NewFloatingScalableMatrixWithTieredPricingPriceParam].
type NewFloatingScalableMatrixWithTieredPricingPriceUnionParam interface {
	ImplementsNewFloatingScalableMatrixWithTieredPricingPriceUnionParam()
}

type NewFloatingScalableMatrixWithTieredPricingPriceConversionRateType string

const (
	NewFloatingScalableMatrixWithTieredPricingPriceConversionRateTypeUnit   NewFloatingScalableMatrixWithTieredPricingPriceConversionRateType = "unit"
	NewFloatingScalableMatrixWithTieredPricingPriceConversionRateTypeTiered NewFloatingScalableMatrixWithTieredPricingPriceConversionRateType = "tiered"
)

func (r NewFloatingScalableMatrixWithTieredPricingPriceConversionRateType) IsKnown() bool {
	switch r {
	case NewFloatingScalableMatrixWithTieredPricingPriceConversionRateTypeUnit, NewFloatingScalableMatrixWithTieredPricingPriceConversionRateTypeTiered:
		return true
	}
	return false
}

type NewFloatingScalableMatrixWithUnitPricingPriceParam struct {
	// The cadence to bill for this price on.
	Cadence param.Field[NewFloatingScalableMatrixWithUnitPricingPriceCadence] `json:"cadence,required"`
	// An ISO 4217 currency string for which this price is billed in.
	Currency param.Field[string] `json:"currency,required"`
	// The id of the item the price will be associated with.
	ItemID    param.Field[string]                                                 `json:"item_id,required"`
	ModelType param.Field[NewFloatingScalableMatrixWithUnitPricingPriceModelType] `json:"model_type,required"`
	// The name of the price.
	Name                                param.Field[string]                 `json:"name,required"`
	ScalableMatrixWithUnitPricingConfig param.Field[map[string]interface{}] `json:"scalable_matrix_with_unit_pricing_config,required"`
	// The id of the billable metric for the price. Only needed if the price is
	// usage-based.
	BillableMetricID param.Field[string] `json:"billable_metric_id"`
	// If the Price represents a fixed cost, the price will be billed in-advance if
	// this is true, and in-arrears if this is false.
	BilledInAdvance param.Field[bool] `json:"billed_in_advance"`
	// For custom cadence: specifies the duration of the billing period in days or
	// months.
	BillingCycleConfiguration param.Field[NewBillingCycleConfigurationParam] `json:"billing_cycle_configuration"`
	// The per unit conversion rate of the price currency to the invoicing currency.
	ConversionRate param.Field[float64] `json:"conversion_rate"`
	// The configuration for the rate of the price currency to the invoicing currency.
	ConversionRateConfig param.Field[NewFloatingScalableMatrixWithUnitPricingPriceUnionParam] `json:"conversion_rate_config"`
	// For dimensional price: specifies a price group and dimension values
	DimensionalPriceConfiguration param.Field[NewDimensionalPriceConfigurationParam] `json:"dimensional_price_configuration"`
	// An alias for the price.
	ExternalPriceID param.Field[string] `json:"external_price_id"`
	// If the Price represents a fixed cost, this represents the quantity of units
	// applied.
	FixedPriceQuantity param.Field[float64] `json:"fixed_price_quantity"`
	// The property used to group this price on an invoice
	InvoiceGroupingKey param.Field[string] `json:"invoice_grouping_key"`
	// Within each billing cycle, specifies the cadence at which invoices are produced.
	// If unspecified, a single invoice is produced per billing cycle.
	InvoicingCycleConfiguration param.Field[NewBillingCycleConfigurationParam] `json:"invoicing_cycle_configuration"`
	// User-specified key/value pairs for the resource. Individual keys can be removed
	// by setting the value to `null`, and the entire metadata mapping can be cleared
	// by setting `metadata` to `null`.
	Metadata param.Field[map[string]string] `json:"metadata"`
}

func (r NewFloatingScalableMatrixWithUnitPricingPriceParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r NewFloatingScalableMatrixWithUnitPricingPriceParam) ImplementsPriceEvaluateMultipleParamsPriceEvaluationsPriceUnion() {
}

func (r NewFloatingScalableMatrixWithUnitPricingPriceParam) ImplementsPriceEvaluatePreviewEventsParamsPriceEvaluationsPriceUnion() {
}

func (r NewFloatingScalableMatrixWithUnitPricingPriceParam) ImplementsSubscriptionPriceIntervalsParamsAddPriceUnion() {
}

// The cadence to bill for this price on.
type NewFloatingScalableMatrixWithUnitPricingPriceCadence string

const (
	NewFloatingScalableMatrixWithUnitPricingPriceCadenceAnnual     NewFloatingScalableMatrixWithUnitPricingPriceCadence = "annual"
	NewFloatingScalableMatrixWithUnitPricingPriceCadenceSemiAnnual NewFloatingScalableMatrixWithUnitPricingPriceCadence = "semi_annual"
	NewFloatingScalableMatrixWithUnitPricingPriceCadenceMonthly    NewFloatingScalableMatrixWithUnitPricingPriceCadence = "monthly"
	NewFloatingScalableMatrixWithUnitPricingPriceCadenceQuarterly  NewFloatingScalableMatrixWithUnitPricingPriceCadence = "quarterly"
	NewFloatingScalableMatrixWithUnitPricingPriceCadenceOneTime    NewFloatingScalableMatrixWithUnitPricingPriceCadence = "one_time"
	NewFloatingScalableMatrixWithUnitPricingPriceCadenceCustom     NewFloatingScalableMatrixWithUnitPricingPriceCadence = "custom"
)

func (r NewFloatingScalableMatrixWithUnitPricingPriceCadence) IsKnown() bool {
	switch r {
	case NewFloatingScalableMatrixWithUnitPricingPriceCadenceAnnual, NewFloatingScalableMatrixWithUnitPricingPriceCadenceSemiAnnual, NewFloatingScalableMatrixWithUnitPricingPriceCadenceMonthly, NewFloatingScalableMatrixWithUnitPricingPriceCadenceQuarterly, NewFloatingScalableMatrixWithUnitPricingPriceCadenceOneTime, NewFloatingScalableMatrixWithUnitPricingPriceCadenceCustom:
		return true
	}
	return false
}

type NewFloatingScalableMatrixWithUnitPricingPriceModelType string

const (
	NewFloatingScalableMatrixWithUnitPricingPriceModelTypeScalableMatrixWithUnitPricing NewFloatingScalableMatrixWithUnitPricingPriceModelType = "scalable_matrix_with_unit_pricing"
)

func (r NewFloatingScalableMatrixWithUnitPricingPriceModelType) IsKnown() bool {
	switch r {
	case NewFloatingScalableMatrixWithUnitPricingPriceModelTypeScalableMatrixWithUnitPricing:
		return true
	}
	return false
}

type NewFloatingScalableMatrixWithUnitPricingPriceParam struct {
	ConversionRateType param.Field[NewFloatingScalableMatrixWithUnitPricingPriceConversionRateType] `json:"conversion_rate_type,required"`
	TieredConfig       param.Field[ConversionRateTieredConfigParam]                                 `json:"tiered_config"`
	UnitConfig         param.Field[ConversionRateUnitConfigParam]                                   `json:"unit_config"`
}

func (r NewFloatingScalableMatrixWithUnitPricingPriceParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r NewFloatingScalableMatrixWithUnitPricingPriceParam) ImplementsNewFloatingScalableMatrixWithUnitPricingPriceUnionParam() {
}

// Satisfied by [shared.UnitConversionRateConfigParam],
// [shared.TieredConversionRateConfigParam],
// [NewFloatingScalableMatrixWithUnitPricingPriceParam].
type NewFloatingScalableMatrixWithUnitPricingPriceUnionParam interface {
	ImplementsNewFloatingScalableMatrixWithUnitPricingPriceUnionParam()
}

type NewFloatingScalableMatrixWithUnitPricingPriceConversionRateType string

const (
	NewFloatingScalableMatrixWithUnitPricingPriceConversionRateTypeUnit   NewFloatingScalableMatrixWithUnitPricingPriceConversionRateType = "unit"
	NewFloatingScalableMatrixWithUnitPricingPriceConversionRateTypeTiered NewFloatingScalableMatrixWithUnitPricingPriceConversionRateType = "tiered"
)

func (r NewFloatingScalableMatrixWithUnitPricingPriceConversionRateType) IsKnown() bool {
	switch r {
	case NewFloatingScalableMatrixWithUnitPricingPriceConversionRateTypeUnit, NewFloatingScalableMatrixWithUnitPricingPriceConversionRateTypeTiered:
		return true
	}
	return false
}

type NewFloatingThresholdTotalAmountPriceParam struct {
	// The cadence to bill for this price on.
	Cadence param.Field[NewFloatingThresholdTotalAmountPriceCadence] `json:"cadence,required"`
	// An ISO 4217 currency string for which this price is billed in.
	Currency param.Field[string] `json:"currency,required"`
	// The id of the item the price will be associated with.
	ItemID    param.Field[string]                                        `json:"item_id,required"`
	ModelType param.Field[NewFloatingThresholdTotalAmountPriceModelType] `json:"model_type,required"`
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
	BillingCycleConfiguration param.Field[NewBillingCycleConfigurationParam] `json:"billing_cycle_configuration"`
	// The per unit conversion rate of the price currency to the invoicing currency.
	ConversionRate param.Field[float64] `json:"conversion_rate"`
	// The configuration for the rate of the price currency to the invoicing currency.
	ConversionRateConfig param.Field[NewFloatingThresholdTotalAmountPriceUnionParam] `json:"conversion_rate_config"`
	// For dimensional price: specifies a price group and dimension values
	DimensionalPriceConfiguration param.Field[NewDimensionalPriceConfigurationParam] `json:"dimensional_price_configuration"`
	// An alias for the price.
	ExternalPriceID param.Field[string] `json:"external_price_id"`
	// If the Price represents a fixed cost, this represents the quantity of units
	// applied.
	FixedPriceQuantity param.Field[float64] `json:"fixed_price_quantity"`
	// The property used to group this price on an invoice
	InvoiceGroupingKey param.Field[string] `json:"invoice_grouping_key"`
	// Within each billing cycle, specifies the cadence at which invoices are produced.
	// If unspecified, a single invoice is produced per billing cycle.
	InvoicingCycleConfiguration param.Field[NewBillingCycleConfigurationParam] `json:"invoicing_cycle_configuration"`
	// User-specified key/value pairs for the resource. Individual keys can be removed
	// by setting the value to `null`, and the entire metadata mapping can be cleared
	// by setting `metadata` to `null`.
	Metadata param.Field[map[string]string] `json:"metadata"`
}

func (r NewFloatingThresholdTotalAmountPriceParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r NewFloatingThresholdTotalAmountPriceParam) ImplementsPriceEvaluateMultipleParamsPriceEvaluationsPriceUnion() {
}

func (r NewFloatingThresholdTotalAmountPriceParam) ImplementsPriceEvaluatePreviewEventsParamsPriceEvaluationsPriceUnion() {
}

func (r NewFloatingThresholdTotalAmountPriceParam) ImplementsSubscriptionPriceIntervalsParamsAddPriceUnion() {
}

// The cadence to bill for this price on.
type NewFloatingThresholdTotalAmountPriceCadence string

const (
	NewFloatingThresholdTotalAmountPriceCadenceAnnual     NewFloatingThresholdTotalAmountPriceCadence = "annual"
	NewFloatingThresholdTotalAmountPriceCadenceSemiAnnual NewFloatingThresholdTotalAmountPriceCadence = "semi_annual"
	NewFloatingThresholdTotalAmountPriceCadenceMonthly    NewFloatingThresholdTotalAmountPriceCadence = "monthly"
	NewFloatingThresholdTotalAmountPriceCadenceQuarterly  NewFloatingThresholdTotalAmountPriceCadence = "quarterly"
	NewFloatingThresholdTotalAmountPriceCadenceOneTime    NewFloatingThresholdTotalAmountPriceCadence = "one_time"
	NewFloatingThresholdTotalAmountPriceCadenceCustom     NewFloatingThresholdTotalAmountPriceCadence = "custom"
)

func (r NewFloatingThresholdTotalAmountPriceCadence) IsKnown() bool {
	switch r {
	case NewFloatingThresholdTotalAmountPriceCadenceAnnual, NewFloatingThresholdTotalAmountPriceCadenceSemiAnnual, NewFloatingThresholdTotalAmountPriceCadenceMonthly, NewFloatingThresholdTotalAmountPriceCadenceQuarterly, NewFloatingThresholdTotalAmountPriceCadenceOneTime, NewFloatingThresholdTotalAmountPriceCadenceCustom:
		return true
	}
	return false
}

type NewFloatingThresholdTotalAmountPriceModelType string

const (
	NewFloatingThresholdTotalAmountPriceModelTypeThresholdTotalAmount NewFloatingThresholdTotalAmountPriceModelType = "threshold_total_amount"
)

func (r NewFloatingThresholdTotalAmountPriceModelType) IsKnown() bool {
	switch r {
	case NewFloatingThresholdTotalAmountPriceModelTypeThresholdTotalAmount:
		return true
	}
	return false
}

type NewFloatingThresholdTotalAmountPriceParam struct {
	ConversionRateType param.Field[NewFloatingThresholdTotalAmountPriceConversionRateType] `json:"conversion_rate_type,required"`
	TieredConfig       param.Field[ConversionRateTieredConfigParam]                        `json:"tiered_config"`
	UnitConfig         param.Field[ConversionRateUnitConfigParam]                          `json:"unit_config"`
}

func (r NewFloatingThresholdTotalAmountPriceParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r NewFloatingThresholdTotalAmountPriceParam) ImplementsNewFloatingThresholdTotalAmountPriceUnionParam() {
}

// Satisfied by [shared.UnitConversionRateConfigParam],
// [shared.TieredConversionRateConfigParam],
// [NewFloatingThresholdTotalAmountPriceParam].
type NewFloatingThresholdTotalAmountPriceUnionParam interface {
	ImplementsNewFloatingThresholdTotalAmountPriceUnionParam()
}

type NewFloatingThresholdTotalAmountPriceConversionRateType string

const (
	NewFloatingThresholdTotalAmountPriceConversionRateTypeUnit   NewFloatingThresholdTotalAmountPriceConversionRateType = "unit"
	NewFloatingThresholdTotalAmountPriceConversionRateTypeTiered NewFloatingThresholdTotalAmountPriceConversionRateType = "tiered"
)

func (r NewFloatingThresholdTotalAmountPriceConversionRateType) IsKnown() bool {
	switch r {
	case NewFloatingThresholdTotalAmountPriceConversionRateTypeUnit, NewFloatingThresholdTotalAmountPriceConversionRateTypeTiered:
		return true
	}
	return false
}

type NewFloatingTieredBPSPriceParam struct {
	// The cadence to bill for this price on.
	Cadence param.Field[NewFloatingTieredBPSPriceCadence] `json:"cadence,required"`
	// An ISO 4217 currency string for which this price is billed in.
	Currency param.Field[string] `json:"currency,required"`
	// The id of the item the price will be associated with.
	ItemID    param.Field[string]                             `json:"item_id,required"`
	ModelType param.Field[NewFloatingTieredBPSPriceModelType] `json:"model_type,required"`
	// The name of the price.
	Name            param.Field[string]               `json:"name,required"`
	TieredBPSConfig param.Field[TieredBPSConfigParam] `json:"tiered_bps_config,required"`
	// The id of the billable metric for the price. Only needed if the price is
	// usage-based.
	BillableMetricID param.Field[string] `json:"billable_metric_id"`
	// If the Price represents a fixed cost, the price will be billed in-advance if
	// this is true, and in-arrears if this is false.
	BilledInAdvance param.Field[bool] `json:"billed_in_advance"`
	// For custom cadence: specifies the duration of the billing period in days or
	// months.
	BillingCycleConfiguration param.Field[NewBillingCycleConfigurationParam] `json:"billing_cycle_configuration"`
	// The per unit conversion rate of the price currency to the invoicing currency.
	ConversionRate param.Field[float64] `json:"conversion_rate"`
	// The configuration for the rate of the price currency to the invoicing currency.
	ConversionRateConfig param.Field[NewFloatingTieredBPSPriceUnionParam] `json:"conversion_rate_config"`
	// For dimensional price: specifies a price group and dimension values
	DimensionalPriceConfiguration param.Field[NewDimensionalPriceConfigurationParam] `json:"dimensional_price_configuration"`
	// An alias for the price.
	ExternalPriceID param.Field[string] `json:"external_price_id"`
	// If the Price represents a fixed cost, this represents the quantity of units
	// applied.
	FixedPriceQuantity param.Field[float64] `json:"fixed_price_quantity"`
	// The property used to group this price on an invoice
	InvoiceGroupingKey param.Field[string] `json:"invoice_grouping_key"`
	// Within each billing cycle, specifies the cadence at which invoices are produced.
	// If unspecified, a single invoice is produced per billing cycle.
	InvoicingCycleConfiguration param.Field[NewBillingCycleConfigurationParam] `json:"invoicing_cycle_configuration"`
	// User-specified key/value pairs for the resource. Individual keys can be removed
	// by setting the value to `null`, and the entire metadata mapping can be cleared
	// by setting `metadata` to `null`.
	Metadata param.Field[map[string]string] `json:"metadata"`
}

func (r NewFloatingTieredBPSPriceParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r NewFloatingTieredBPSPriceParam) ImplementsPriceEvaluateMultipleParamsPriceEvaluationsPriceUnion() {
}

func (r NewFloatingTieredBPSPriceParam) ImplementsPriceEvaluatePreviewEventsParamsPriceEvaluationsPriceUnion() {
}

func (r NewFloatingTieredBPSPriceParam) ImplementsSubscriptionPriceIntervalsParamsAddPriceUnion() {}

// The cadence to bill for this price on.
type NewFloatingTieredBPSPriceCadence string

const (
	NewFloatingTieredBPSPriceCadenceAnnual     NewFloatingTieredBPSPriceCadence = "annual"
	NewFloatingTieredBPSPriceCadenceSemiAnnual NewFloatingTieredBPSPriceCadence = "semi_annual"
	NewFloatingTieredBPSPriceCadenceMonthly    NewFloatingTieredBPSPriceCadence = "monthly"
	NewFloatingTieredBPSPriceCadenceQuarterly  NewFloatingTieredBPSPriceCadence = "quarterly"
	NewFloatingTieredBPSPriceCadenceOneTime    NewFloatingTieredBPSPriceCadence = "one_time"
	NewFloatingTieredBPSPriceCadenceCustom     NewFloatingTieredBPSPriceCadence = "custom"
)

func (r NewFloatingTieredBPSPriceCadence) IsKnown() bool {
	switch r {
	case NewFloatingTieredBPSPriceCadenceAnnual, NewFloatingTieredBPSPriceCadenceSemiAnnual, NewFloatingTieredBPSPriceCadenceMonthly, NewFloatingTieredBPSPriceCadenceQuarterly, NewFloatingTieredBPSPriceCadenceOneTime, NewFloatingTieredBPSPriceCadenceCustom:
		return true
	}
	return false
}

type NewFloatingTieredBPSPriceModelType string

const (
	NewFloatingTieredBPSPriceModelTypeTieredBPS NewFloatingTieredBPSPriceModelType = "tiered_bps"
)

func (r NewFloatingTieredBPSPriceModelType) IsKnown() bool {
	switch r {
	case NewFloatingTieredBPSPriceModelTypeTieredBPS:
		return true
	}
	return false
}

type NewFloatingTieredBPSPriceParam struct {
	ConversionRateType param.Field[NewFloatingTieredBPSPriceConversionRateType] `json:"conversion_rate_type,required"`
	TieredConfig       param.Field[ConversionRateTieredConfigParam]             `json:"tiered_config"`
	UnitConfig         param.Field[ConversionRateUnitConfigParam]               `json:"unit_config"`
}

func (r NewFloatingTieredBPSPriceParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r NewFloatingTieredBPSPriceParam) ImplementsNewFloatingTieredBPSPriceUnionParam() {}

// Satisfied by [shared.UnitConversionRateConfigParam],
// [shared.TieredConversionRateConfigParam], [NewFloatingTieredBPSPriceParam].
type NewFloatingTieredBPSPriceUnionParam interface {
	ImplementsNewFloatingTieredBPSPriceUnionParam()
}

type NewFloatingTieredBPSPriceConversionRateType string

const (
	NewFloatingTieredBPSPriceConversionRateTypeUnit   NewFloatingTieredBPSPriceConversionRateType = "unit"
	NewFloatingTieredBPSPriceConversionRateTypeTiered NewFloatingTieredBPSPriceConversionRateType = "tiered"
)

func (r NewFloatingTieredBPSPriceConversionRateType) IsKnown() bool {
	switch r {
	case NewFloatingTieredBPSPriceConversionRateTypeUnit, NewFloatingTieredBPSPriceConversionRateTypeTiered:
		return true
	}
	return false
}

type NewFloatingTieredPackagePriceParam struct {
	// The cadence to bill for this price on.
	Cadence param.Field[NewFloatingTieredPackagePriceCadence] `json:"cadence,required"`
	// An ISO 4217 currency string for which this price is billed in.
	Currency param.Field[string] `json:"currency,required"`
	// The id of the item the price will be associated with.
	ItemID    param.Field[string]                                 `json:"item_id,required"`
	ModelType param.Field[NewFloatingTieredPackagePriceModelType] `json:"model_type,required"`
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
	BillingCycleConfiguration param.Field[NewBillingCycleConfigurationParam] `json:"billing_cycle_configuration"`
	// The per unit conversion rate of the price currency to the invoicing currency.
	ConversionRate param.Field[float64] `json:"conversion_rate"`
	// The configuration for the rate of the price currency to the invoicing currency.
	ConversionRateConfig param.Field[NewFloatingTieredPackagePriceUnionParam] `json:"conversion_rate_config"`
	// For dimensional price: specifies a price group and dimension values
	DimensionalPriceConfiguration param.Field[NewDimensionalPriceConfigurationParam] `json:"dimensional_price_configuration"`
	// An alias for the price.
	ExternalPriceID param.Field[string] `json:"external_price_id"`
	// If the Price represents a fixed cost, this represents the quantity of units
	// applied.
	FixedPriceQuantity param.Field[float64] `json:"fixed_price_quantity"`
	// The property used to group this price on an invoice
	InvoiceGroupingKey param.Field[string] `json:"invoice_grouping_key"`
	// Within each billing cycle, specifies the cadence at which invoices are produced.
	// If unspecified, a single invoice is produced per billing cycle.
	InvoicingCycleConfiguration param.Field[NewBillingCycleConfigurationParam] `json:"invoicing_cycle_configuration"`
	// User-specified key/value pairs for the resource. Individual keys can be removed
	// by setting the value to `null`, and the entire metadata mapping can be cleared
	// by setting `metadata` to `null`.
	Metadata param.Field[map[string]string] `json:"metadata"`
}

func (r NewFloatingTieredPackagePriceParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r NewFloatingTieredPackagePriceParam) ImplementsPriceEvaluateMultipleParamsPriceEvaluationsPriceUnion() {
}

func (r NewFloatingTieredPackagePriceParam) ImplementsPriceEvaluatePreviewEventsParamsPriceEvaluationsPriceUnion() {
}

func (r NewFloatingTieredPackagePriceParam) ImplementsSubscriptionPriceIntervalsParamsAddPriceUnion() {
}

// The cadence to bill for this price on.
type NewFloatingTieredPackagePriceCadence string

const (
	NewFloatingTieredPackagePriceCadenceAnnual     NewFloatingTieredPackagePriceCadence = "annual"
	NewFloatingTieredPackagePriceCadenceSemiAnnual NewFloatingTieredPackagePriceCadence = "semi_annual"
	NewFloatingTieredPackagePriceCadenceMonthly    NewFloatingTieredPackagePriceCadence = "monthly"
	NewFloatingTieredPackagePriceCadenceQuarterly  NewFloatingTieredPackagePriceCadence = "quarterly"
	NewFloatingTieredPackagePriceCadenceOneTime    NewFloatingTieredPackagePriceCadence = "one_time"
	NewFloatingTieredPackagePriceCadenceCustom     NewFloatingTieredPackagePriceCadence = "custom"
)

func (r NewFloatingTieredPackagePriceCadence) IsKnown() bool {
	switch r {
	case NewFloatingTieredPackagePriceCadenceAnnual, NewFloatingTieredPackagePriceCadenceSemiAnnual, NewFloatingTieredPackagePriceCadenceMonthly, NewFloatingTieredPackagePriceCadenceQuarterly, NewFloatingTieredPackagePriceCadenceOneTime, NewFloatingTieredPackagePriceCadenceCustom:
		return true
	}
	return false
}

type NewFloatingTieredPackagePriceModelType string

const (
	NewFloatingTieredPackagePriceModelTypeTieredPackage NewFloatingTieredPackagePriceModelType = "tiered_package"
)

func (r NewFloatingTieredPackagePriceModelType) IsKnown() bool {
	switch r {
	case NewFloatingTieredPackagePriceModelTypeTieredPackage:
		return true
	}
	return false
}

type NewFloatingTieredPackagePriceParam struct {
	ConversionRateType param.Field[NewFloatingTieredPackagePriceConversionRateType] `json:"conversion_rate_type,required"`
	TieredConfig       param.Field[ConversionRateTieredConfigParam]                 `json:"tiered_config"`
	UnitConfig         param.Field[ConversionRateUnitConfigParam]                   `json:"unit_config"`
}

func (r NewFloatingTieredPackagePriceParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r NewFloatingTieredPackagePriceParam) ImplementsNewFloatingTieredPackagePriceUnionParam() {}

// Satisfied by [shared.UnitConversionRateConfigParam],
// [shared.TieredConversionRateConfigParam], [NewFloatingTieredPackagePriceParam].
type NewFloatingTieredPackagePriceUnionParam interface {
	ImplementsNewFloatingTieredPackagePriceUnionParam()
}

type NewFloatingTieredPackagePriceConversionRateType string

const (
	NewFloatingTieredPackagePriceConversionRateTypeUnit   NewFloatingTieredPackagePriceConversionRateType = "unit"
	NewFloatingTieredPackagePriceConversionRateTypeTiered NewFloatingTieredPackagePriceConversionRateType = "tiered"
)

func (r NewFloatingTieredPackagePriceConversionRateType) IsKnown() bool {
	switch r {
	case NewFloatingTieredPackagePriceConversionRateTypeUnit, NewFloatingTieredPackagePriceConversionRateTypeTiered:
		return true
	}
	return false
}

type NewFloatingTieredPackageWithMinimumPriceParam struct {
	// The cadence to bill for this price on.
	Cadence param.Field[NewFloatingTieredPackageWithMinimumPriceCadence] `json:"cadence,required"`
	// An ISO 4217 currency string for which this price is billed in.
	Currency param.Field[string] `json:"currency,required"`
	// The id of the item the price will be associated with.
	ItemID    param.Field[string]                                            `json:"item_id,required"`
	ModelType param.Field[NewFloatingTieredPackageWithMinimumPriceModelType] `json:"model_type,required"`
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
	BillingCycleConfiguration param.Field[NewBillingCycleConfigurationParam] `json:"billing_cycle_configuration"`
	// The per unit conversion rate of the price currency to the invoicing currency.
	ConversionRate param.Field[float64] `json:"conversion_rate"`
	// The configuration for the rate of the price currency to the invoicing currency.
	ConversionRateConfig param.Field[NewFloatingTieredPackageWithMinimumPriceUnionParam] `json:"conversion_rate_config"`
	// For dimensional price: specifies a price group and dimension values
	DimensionalPriceConfiguration param.Field[NewDimensionalPriceConfigurationParam] `json:"dimensional_price_configuration"`
	// An alias for the price.
	ExternalPriceID param.Field[string] `json:"external_price_id"`
	// If the Price represents a fixed cost, this represents the quantity of units
	// applied.
	FixedPriceQuantity param.Field[float64] `json:"fixed_price_quantity"`
	// The property used to group this price on an invoice
	InvoiceGroupingKey param.Field[string] `json:"invoice_grouping_key"`
	// Within each billing cycle, specifies the cadence at which invoices are produced.
	// If unspecified, a single invoice is produced per billing cycle.
	InvoicingCycleConfiguration param.Field[NewBillingCycleConfigurationParam] `json:"invoicing_cycle_configuration"`
	// User-specified key/value pairs for the resource. Individual keys can be removed
	// by setting the value to `null`, and the entire metadata mapping can be cleared
	// by setting `metadata` to `null`.
	Metadata param.Field[map[string]string] `json:"metadata"`
}

func (r NewFloatingTieredPackageWithMinimumPriceParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r NewFloatingTieredPackageWithMinimumPriceParam) ImplementsPriceEvaluateMultipleParamsPriceEvaluationsPriceUnion() {
}

func (r NewFloatingTieredPackageWithMinimumPriceParam) ImplementsPriceEvaluatePreviewEventsParamsPriceEvaluationsPriceUnion() {
}

func (r NewFloatingTieredPackageWithMinimumPriceParam) ImplementsSubscriptionPriceIntervalsParamsAddPriceUnion() {
}

// The cadence to bill for this price on.
type NewFloatingTieredPackageWithMinimumPriceCadence string

const (
	NewFloatingTieredPackageWithMinimumPriceCadenceAnnual     NewFloatingTieredPackageWithMinimumPriceCadence = "annual"
	NewFloatingTieredPackageWithMinimumPriceCadenceSemiAnnual NewFloatingTieredPackageWithMinimumPriceCadence = "semi_annual"
	NewFloatingTieredPackageWithMinimumPriceCadenceMonthly    NewFloatingTieredPackageWithMinimumPriceCadence = "monthly"
	NewFloatingTieredPackageWithMinimumPriceCadenceQuarterly  NewFloatingTieredPackageWithMinimumPriceCadence = "quarterly"
	NewFloatingTieredPackageWithMinimumPriceCadenceOneTime    NewFloatingTieredPackageWithMinimumPriceCadence = "one_time"
	NewFloatingTieredPackageWithMinimumPriceCadenceCustom     NewFloatingTieredPackageWithMinimumPriceCadence = "custom"
)

func (r NewFloatingTieredPackageWithMinimumPriceCadence) IsKnown() bool {
	switch r {
	case NewFloatingTieredPackageWithMinimumPriceCadenceAnnual, NewFloatingTieredPackageWithMinimumPriceCadenceSemiAnnual, NewFloatingTieredPackageWithMinimumPriceCadenceMonthly, NewFloatingTieredPackageWithMinimumPriceCadenceQuarterly, NewFloatingTieredPackageWithMinimumPriceCadenceOneTime, NewFloatingTieredPackageWithMinimumPriceCadenceCustom:
		return true
	}
	return false
}

type NewFloatingTieredPackageWithMinimumPriceModelType string

const (
	NewFloatingTieredPackageWithMinimumPriceModelTypeTieredPackageWithMinimum NewFloatingTieredPackageWithMinimumPriceModelType = "tiered_package_with_minimum"
)

func (r NewFloatingTieredPackageWithMinimumPriceModelType) IsKnown() bool {
	switch r {
	case NewFloatingTieredPackageWithMinimumPriceModelTypeTieredPackageWithMinimum:
		return true
	}
	return false
}

type NewFloatingTieredPackageWithMinimumPriceParam struct {
	ConversionRateType param.Field[NewFloatingTieredPackageWithMinimumPriceConversionRateType] `json:"conversion_rate_type,required"`
	TieredConfig       param.Field[ConversionRateTieredConfigParam]                            `json:"tiered_config"`
	UnitConfig         param.Field[ConversionRateUnitConfigParam]                              `json:"unit_config"`
}

func (r NewFloatingTieredPackageWithMinimumPriceParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r NewFloatingTieredPackageWithMinimumPriceParam) ImplementsNewFloatingTieredPackageWithMinimumPriceUnionParam() {
}

// Satisfied by [shared.UnitConversionRateConfigParam],
// [shared.TieredConversionRateConfigParam],
// [NewFloatingTieredPackageWithMinimumPriceParam].
type NewFloatingTieredPackageWithMinimumPriceUnionParam interface {
	ImplementsNewFloatingTieredPackageWithMinimumPriceUnionParam()
}

type NewFloatingTieredPackageWithMinimumPriceConversionRateType string

const (
	NewFloatingTieredPackageWithMinimumPriceConversionRateTypeUnit   NewFloatingTieredPackageWithMinimumPriceConversionRateType = "unit"
	NewFloatingTieredPackageWithMinimumPriceConversionRateTypeTiered NewFloatingTieredPackageWithMinimumPriceConversionRateType = "tiered"
)

func (r NewFloatingTieredPackageWithMinimumPriceConversionRateType) IsKnown() bool {
	switch r {
	case NewFloatingTieredPackageWithMinimumPriceConversionRateTypeUnit, NewFloatingTieredPackageWithMinimumPriceConversionRateTypeTiered:
		return true
	}
	return false
}

type NewFloatingTieredPriceParam struct {
	// The cadence to bill for this price on.
	Cadence param.Field[NewFloatingTieredPriceCadence] `json:"cadence,required"`
	// An ISO 4217 currency string for which this price is billed in.
	Currency param.Field[string] `json:"currency,required"`
	// The id of the item the price will be associated with.
	ItemID    param.Field[string]                          `json:"item_id,required"`
	ModelType param.Field[NewFloatingTieredPriceModelType] `json:"model_type,required"`
	// The name of the price.
	Name         param.Field[string]            `json:"name,required"`
	TieredConfig param.Field[TieredConfigParam] `json:"tiered_config,required"`
	// The id of the billable metric for the price. Only needed if the price is
	// usage-based.
	BillableMetricID param.Field[string] `json:"billable_metric_id"`
	// If the Price represents a fixed cost, the price will be billed in-advance if
	// this is true, and in-arrears if this is false.
	BilledInAdvance param.Field[bool] `json:"billed_in_advance"`
	// For custom cadence: specifies the duration of the billing period in days or
	// months.
	BillingCycleConfiguration param.Field[NewBillingCycleConfigurationParam] `json:"billing_cycle_configuration"`
	// The per unit conversion rate of the price currency to the invoicing currency.
	ConversionRate param.Field[float64] `json:"conversion_rate"`
	// The configuration for the rate of the price currency to the invoicing currency.
	ConversionRateConfig param.Field[NewFloatingTieredPriceUnionParam] `json:"conversion_rate_config"`
	// For dimensional price: specifies a price group and dimension values
	DimensionalPriceConfiguration param.Field[NewDimensionalPriceConfigurationParam] `json:"dimensional_price_configuration"`
	// An alias for the price.
	ExternalPriceID param.Field[string] `json:"external_price_id"`
	// If the Price represents a fixed cost, this represents the quantity of units
	// applied.
	FixedPriceQuantity param.Field[float64] `json:"fixed_price_quantity"`
	// The property used to group this price on an invoice
	InvoiceGroupingKey param.Field[string] `json:"invoice_grouping_key"`
	// Within each billing cycle, specifies the cadence at which invoices are produced.
	// If unspecified, a single invoice is produced per billing cycle.
	InvoicingCycleConfiguration param.Field[NewBillingCycleConfigurationParam] `json:"invoicing_cycle_configuration"`
	// User-specified key/value pairs for the resource. Individual keys can be removed
	// by setting the value to `null`, and the entire metadata mapping can be cleared
	// by setting `metadata` to `null`.
	Metadata param.Field[map[string]string] `json:"metadata"`
}

func (r NewFloatingTieredPriceParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r NewFloatingTieredPriceParam) ImplementsPriceEvaluateMultipleParamsPriceEvaluationsPriceUnion() {
}

func (r NewFloatingTieredPriceParam) ImplementsPriceEvaluatePreviewEventsParamsPriceEvaluationsPriceUnion() {
}

func (r NewFloatingTieredPriceParam) ImplementsSubscriptionPriceIntervalsParamsAddPriceUnion() {}

// The cadence to bill for this price on.
type NewFloatingTieredPriceCadence string

const (
	NewFloatingTieredPriceCadenceAnnual     NewFloatingTieredPriceCadence = "annual"
	NewFloatingTieredPriceCadenceSemiAnnual NewFloatingTieredPriceCadence = "semi_annual"
	NewFloatingTieredPriceCadenceMonthly    NewFloatingTieredPriceCadence = "monthly"
	NewFloatingTieredPriceCadenceQuarterly  NewFloatingTieredPriceCadence = "quarterly"
	NewFloatingTieredPriceCadenceOneTime    NewFloatingTieredPriceCadence = "one_time"
	NewFloatingTieredPriceCadenceCustom     NewFloatingTieredPriceCadence = "custom"
)

func (r NewFloatingTieredPriceCadence) IsKnown() bool {
	switch r {
	case NewFloatingTieredPriceCadenceAnnual, NewFloatingTieredPriceCadenceSemiAnnual, NewFloatingTieredPriceCadenceMonthly, NewFloatingTieredPriceCadenceQuarterly, NewFloatingTieredPriceCadenceOneTime, NewFloatingTieredPriceCadenceCustom:
		return true
	}
	return false
}

type NewFloatingTieredPriceModelType string

const (
	NewFloatingTieredPriceModelTypeTiered NewFloatingTieredPriceModelType = "tiered"
)

func (r NewFloatingTieredPriceModelType) IsKnown() bool {
	switch r {
	case NewFloatingTieredPriceModelTypeTiered:
		return true
	}
	return false
}

type NewFloatingTieredPriceParam struct {
	ConversionRateType param.Field[NewFloatingTieredPriceConversionRateType] `json:"conversion_rate_type,required"`
	TieredConfig       param.Field[ConversionRateTieredConfigParam]          `json:"tiered_config"`
	UnitConfig         param.Field[ConversionRateUnitConfigParam]            `json:"unit_config"`
}

func (r NewFloatingTieredPriceParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r NewFloatingTieredPriceParam) ImplementsNewFloatingTieredPriceUnionParam() {}

// Satisfied by [shared.UnitConversionRateConfigParam],
// [shared.TieredConversionRateConfigParam], [NewFloatingTieredPriceParam].
type NewFloatingTieredPriceUnionParam interface {
	ImplementsNewFloatingTieredPriceUnionParam()
}

type NewFloatingTieredPriceConversionRateType string

const (
	NewFloatingTieredPriceConversionRateTypeUnit   NewFloatingTieredPriceConversionRateType = "unit"
	NewFloatingTieredPriceConversionRateTypeTiered NewFloatingTieredPriceConversionRateType = "tiered"
)

func (r NewFloatingTieredPriceConversionRateType) IsKnown() bool {
	switch r {
	case NewFloatingTieredPriceConversionRateTypeUnit, NewFloatingTieredPriceConversionRateTypeTiered:
		return true
	}
	return false
}

type NewFloatingTieredWithMinimumPriceParam struct {
	// The cadence to bill for this price on.
	Cadence param.Field[NewFloatingTieredWithMinimumPriceCadence] `json:"cadence,required"`
	// An ISO 4217 currency string for which this price is billed in.
	Currency param.Field[string] `json:"currency,required"`
	// The id of the item the price will be associated with.
	ItemID    param.Field[string]                                     `json:"item_id,required"`
	ModelType param.Field[NewFloatingTieredWithMinimumPriceModelType] `json:"model_type,required"`
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
	BillingCycleConfiguration param.Field[NewBillingCycleConfigurationParam] `json:"billing_cycle_configuration"`
	// The per unit conversion rate of the price currency to the invoicing currency.
	ConversionRate param.Field[float64] `json:"conversion_rate"`
	// The configuration for the rate of the price currency to the invoicing currency.
	ConversionRateConfig param.Field[NewFloatingTieredWithMinimumPriceUnionParam] `json:"conversion_rate_config"`
	// For dimensional price: specifies a price group and dimension values
	DimensionalPriceConfiguration param.Field[NewDimensionalPriceConfigurationParam] `json:"dimensional_price_configuration"`
	// An alias for the price.
	ExternalPriceID param.Field[string] `json:"external_price_id"`
	// If the Price represents a fixed cost, this represents the quantity of units
	// applied.
	FixedPriceQuantity param.Field[float64] `json:"fixed_price_quantity"`
	// The property used to group this price on an invoice
	InvoiceGroupingKey param.Field[string] `json:"invoice_grouping_key"`
	// Within each billing cycle, specifies the cadence at which invoices are produced.
	// If unspecified, a single invoice is produced per billing cycle.
	InvoicingCycleConfiguration param.Field[NewBillingCycleConfigurationParam] `json:"invoicing_cycle_configuration"`
	// User-specified key/value pairs for the resource. Individual keys can be removed
	// by setting the value to `null`, and the entire metadata mapping can be cleared
	// by setting `metadata` to `null`.
	Metadata param.Field[map[string]string] `json:"metadata"`
}

func (r NewFloatingTieredWithMinimumPriceParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r NewFloatingTieredWithMinimumPriceParam) ImplementsPriceEvaluateMultipleParamsPriceEvaluationsPriceUnion() {
}

func (r NewFloatingTieredWithMinimumPriceParam) ImplementsPriceEvaluatePreviewEventsParamsPriceEvaluationsPriceUnion() {
}

func (r NewFloatingTieredWithMinimumPriceParam) ImplementsSubscriptionPriceIntervalsParamsAddPriceUnion() {
}

// The cadence to bill for this price on.
type NewFloatingTieredWithMinimumPriceCadence string

const (
	NewFloatingTieredWithMinimumPriceCadenceAnnual     NewFloatingTieredWithMinimumPriceCadence = "annual"
	NewFloatingTieredWithMinimumPriceCadenceSemiAnnual NewFloatingTieredWithMinimumPriceCadence = "semi_annual"
	NewFloatingTieredWithMinimumPriceCadenceMonthly    NewFloatingTieredWithMinimumPriceCadence = "monthly"
	NewFloatingTieredWithMinimumPriceCadenceQuarterly  NewFloatingTieredWithMinimumPriceCadence = "quarterly"
	NewFloatingTieredWithMinimumPriceCadenceOneTime    NewFloatingTieredWithMinimumPriceCadence = "one_time"
	NewFloatingTieredWithMinimumPriceCadenceCustom     NewFloatingTieredWithMinimumPriceCadence = "custom"
)

func (r NewFloatingTieredWithMinimumPriceCadence) IsKnown() bool {
	switch r {
	case NewFloatingTieredWithMinimumPriceCadenceAnnual, NewFloatingTieredWithMinimumPriceCadenceSemiAnnual, NewFloatingTieredWithMinimumPriceCadenceMonthly, NewFloatingTieredWithMinimumPriceCadenceQuarterly, NewFloatingTieredWithMinimumPriceCadenceOneTime, NewFloatingTieredWithMinimumPriceCadenceCustom:
		return true
	}
	return false
}

type NewFloatingTieredWithMinimumPriceModelType string

const (
	NewFloatingTieredWithMinimumPriceModelTypeTieredWithMinimum NewFloatingTieredWithMinimumPriceModelType = "tiered_with_minimum"
)

func (r NewFloatingTieredWithMinimumPriceModelType) IsKnown() bool {
	switch r {
	case NewFloatingTieredWithMinimumPriceModelTypeTieredWithMinimum:
		return true
	}
	return false
}

type NewFloatingTieredWithMinimumPriceParam struct {
	ConversionRateType param.Field[NewFloatingTieredWithMinimumPriceConversionRateType] `json:"conversion_rate_type,required"`
	TieredConfig       param.Field[ConversionRateTieredConfigParam]                     `json:"tiered_config"`
	UnitConfig         param.Field[ConversionRateUnitConfigParam]                       `json:"unit_config"`
}

func (r NewFloatingTieredWithMinimumPriceParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r NewFloatingTieredWithMinimumPriceParam) ImplementsNewFloatingTieredWithMinimumPriceUnionParam() {
}

// Satisfied by [shared.UnitConversionRateConfigParam],
// [shared.TieredConversionRateConfigParam],
// [NewFloatingTieredWithMinimumPriceParam].
type NewFloatingTieredWithMinimumPriceUnionParam interface {
	ImplementsNewFloatingTieredWithMinimumPriceUnionParam()
}

type NewFloatingTieredWithMinimumPriceConversionRateType string

const (
	NewFloatingTieredWithMinimumPriceConversionRateTypeUnit   NewFloatingTieredWithMinimumPriceConversionRateType = "unit"
	NewFloatingTieredWithMinimumPriceConversionRateTypeTiered NewFloatingTieredWithMinimumPriceConversionRateType = "tiered"
)

func (r NewFloatingTieredWithMinimumPriceConversionRateType) IsKnown() bool {
	switch r {
	case NewFloatingTieredWithMinimumPriceConversionRateTypeUnit, NewFloatingTieredWithMinimumPriceConversionRateTypeTiered:
		return true
	}
	return false
}

type NewFloatingTieredWithProrationPriceParam struct {
	// The cadence to bill for this price on.
	Cadence param.Field[NewFloatingTieredWithProrationPriceCadence] `json:"cadence,required"`
	// An ISO 4217 currency string for which this price is billed in.
	Currency param.Field[string] `json:"currency,required"`
	// The id of the item the price will be associated with.
	ItemID    param.Field[string]                                       `json:"item_id,required"`
	ModelType param.Field[NewFloatingTieredWithProrationPriceModelType] `json:"model_type,required"`
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
	BillingCycleConfiguration param.Field[NewBillingCycleConfigurationParam] `json:"billing_cycle_configuration"`
	// The per unit conversion rate of the price currency to the invoicing currency.
	ConversionRate param.Field[float64] `json:"conversion_rate"`
	// The configuration for the rate of the price currency to the invoicing currency.
	ConversionRateConfig param.Field[NewFloatingTieredWithProrationPriceUnionParam] `json:"conversion_rate_config"`
	// For dimensional price: specifies a price group and dimension values
	DimensionalPriceConfiguration param.Field[NewDimensionalPriceConfigurationParam] `json:"dimensional_price_configuration"`
	// An alias for the price.
	ExternalPriceID param.Field[string] `json:"external_price_id"`
	// If the Price represents a fixed cost, this represents the quantity of units
	// applied.
	FixedPriceQuantity param.Field[float64] `json:"fixed_price_quantity"`
	// The property used to group this price on an invoice
	InvoiceGroupingKey param.Field[string] `json:"invoice_grouping_key"`
	// Within each billing cycle, specifies the cadence at which invoices are produced.
	// If unspecified, a single invoice is produced per billing cycle.
	InvoicingCycleConfiguration param.Field[NewBillingCycleConfigurationParam] `json:"invoicing_cycle_configuration"`
	// User-specified key/value pairs for the resource. Individual keys can be removed
	// by setting the value to `null`, and the entire metadata mapping can be cleared
	// by setting `metadata` to `null`.
	Metadata param.Field[map[string]string] `json:"metadata"`
}

func (r NewFloatingTieredWithProrationPriceParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r NewFloatingTieredWithProrationPriceParam) ImplementsPriceEvaluateMultipleParamsPriceEvaluationsPriceUnion() {
}

func (r NewFloatingTieredWithProrationPriceParam) ImplementsPriceEvaluatePreviewEventsParamsPriceEvaluationsPriceUnion() {
}

func (r NewFloatingTieredWithProrationPriceParam) ImplementsSubscriptionPriceIntervalsParamsAddPriceUnion() {
}

// The cadence to bill for this price on.
type NewFloatingTieredWithProrationPriceCadence string

const (
	NewFloatingTieredWithProrationPriceCadenceAnnual     NewFloatingTieredWithProrationPriceCadence = "annual"
	NewFloatingTieredWithProrationPriceCadenceSemiAnnual NewFloatingTieredWithProrationPriceCadence = "semi_annual"
	NewFloatingTieredWithProrationPriceCadenceMonthly    NewFloatingTieredWithProrationPriceCadence = "monthly"
	NewFloatingTieredWithProrationPriceCadenceQuarterly  NewFloatingTieredWithProrationPriceCadence = "quarterly"
	NewFloatingTieredWithProrationPriceCadenceOneTime    NewFloatingTieredWithProrationPriceCadence = "one_time"
	NewFloatingTieredWithProrationPriceCadenceCustom     NewFloatingTieredWithProrationPriceCadence = "custom"
)

func (r NewFloatingTieredWithProrationPriceCadence) IsKnown() bool {
	switch r {
	case NewFloatingTieredWithProrationPriceCadenceAnnual, NewFloatingTieredWithProrationPriceCadenceSemiAnnual, NewFloatingTieredWithProrationPriceCadenceMonthly, NewFloatingTieredWithProrationPriceCadenceQuarterly, NewFloatingTieredWithProrationPriceCadenceOneTime, NewFloatingTieredWithProrationPriceCadenceCustom:
		return true
	}
	return false
}

type NewFloatingTieredWithProrationPriceModelType string

const (
	NewFloatingTieredWithProrationPriceModelTypeTieredWithProration NewFloatingTieredWithProrationPriceModelType = "tiered_with_proration"
)

func (r NewFloatingTieredWithProrationPriceModelType) IsKnown() bool {
	switch r {
	case NewFloatingTieredWithProrationPriceModelTypeTieredWithProration:
		return true
	}
	return false
}

type NewFloatingTieredWithProrationPriceParam struct {
	ConversionRateType param.Field[NewFloatingTieredWithProrationPriceConversionRateType] `json:"conversion_rate_type,required"`
	TieredConfig       param.Field[ConversionRateTieredConfigParam]                       `json:"tiered_config"`
	UnitConfig         param.Field[ConversionRateUnitConfigParam]                         `json:"unit_config"`
}

func (r NewFloatingTieredWithProrationPriceParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r NewFloatingTieredWithProrationPriceParam) ImplementsNewFloatingTieredWithProrationPriceUnionParam() {
}

// Satisfied by [shared.UnitConversionRateConfigParam],
// [shared.TieredConversionRateConfigParam],
// [NewFloatingTieredWithProrationPriceParam].
type NewFloatingTieredWithProrationPriceUnionParam interface {
	ImplementsNewFloatingTieredWithProrationPriceUnionParam()
}

type NewFloatingTieredWithProrationPriceConversionRateType string

const (
	NewFloatingTieredWithProrationPriceConversionRateTypeUnit   NewFloatingTieredWithProrationPriceConversionRateType = "unit"
	NewFloatingTieredWithProrationPriceConversionRateTypeTiered NewFloatingTieredWithProrationPriceConversionRateType = "tiered"
)

func (r NewFloatingTieredWithProrationPriceConversionRateType) IsKnown() bool {
	switch r {
	case NewFloatingTieredWithProrationPriceConversionRateTypeUnit, NewFloatingTieredWithProrationPriceConversionRateTypeTiered:
		return true
	}
	return false
}

type NewFloatingUnitPriceParam struct {
	// The cadence to bill for this price on.
	Cadence param.Field[NewFloatingUnitPriceCadence] `json:"cadence,required"`
	// An ISO 4217 currency string for which this price is billed in.
	Currency param.Field[string] `json:"currency,required"`
	// The id of the item the price will be associated with.
	ItemID    param.Field[string]                        `json:"item_id,required"`
	ModelType param.Field[NewFloatingUnitPriceModelType] `json:"model_type,required"`
	// The name of the price.
	Name       param.Field[string]          `json:"name,required"`
	UnitConfig param.Field[UnitConfigParam] `json:"unit_config,required"`
	// The id of the billable metric for the price. Only needed if the price is
	// usage-based.
	BillableMetricID param.Field[string] `json:"billable_metric_id"`
	// If the Price represents a fixed cost, the price will be billed in-advance if
	// this is true, and in-arrears if this is false.
	BilledInAdvance param.Field[bool] `json:"billed_in_advance"`
	// For custom cadence: specifies the duration of the billing period in days or
	// months.
	BillingCycleConfiguration param.Field[NewBillingCycleConfigurationParam] `json:"billing_cycle_configuration"`
	// The per unit conversion rate of the price currency to the invoicing currency.
	ConversionRate param.Field[float64] `json:"conversion_rate"`
	// The configuration for the rate of the price currency to the invoicing currency.
	ConversionRateConfig param.Field[NewFloatingUnitPriceUnionParam] `json:"conversion_rate_config"`
	// For dimensional price: specifies a price group and dimension values
	DimensionalPriceConfiguration param.Field[NewDimensionalPriceConfigurationParam] `json:"dimensional_price_configuration"`
	// An alias for the price.
	ExternalPriceID param.Field[string] `json:"external_price_id"`
	// If the Price represents a fixed cost, this represents the quantity of units
	// applied.
	FixedPriceQuantity param.Field[float64] `json:"fixed_price_quantity"`
	// The property used to group this price on an invoice
	InvoiceGroupingKey param.Field[string] `json:"invoice_grouping_key"`
	// Within each billing cycle, specifies the cadence at which invoices are produced.
	// If unspecified, a single invoice is produced per billing cycle.
	InvoicingCycleConfiguration param.Field[NewBillingCycleConfigurationParam] `json:"invoicing_cycle_configuration"`
	// User-specified key/value pairs for the resource. Individual keys can be removed
	// by setting the value to `null`, and the entire metadata mapping can be cleared
	// by setting `metadata` to `null`.
	Metadata param.Field[map[string]string] `json:"metadata"`
}

func (r NewFloatingUnitPriceParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r NewFloatingUnitPriceParam) ImplementsPriceEvaluateMultipleParamsPriceEvaluationsPriceUnion() {
}

func (r NewFloatingUnitPriceParam) ImplementsPriceEvaluatePreviewEventsParamsPriceEvaluationsPriceUnion() {
}

func (r NewFloatingUnitPriceParam) ImplementsSubscriptionPriceIntervalsParamsAddPriceUnion() {}

// The cadence to bill for this price on.
type NewFloatingUnitPriceCadence string

const (
	NewFloatingUnitPriceCadenceAnnual     NewFloatingUnitPriceCadence = "annual"
	NewFloatingUnitPriceCadenceSemiAnnual NewFloatingUnitPriceCadence = "semi_annual"
	NewFloatingUnitPriceCadenceMonthly    NewFloatingUnitPriceCadence = "monthly"
	NewFloatingUnitPriceCadenceQuarterly  NewFloatingUnitPriceCadence = "quarterly"
	NewFloatingUnitPriceCadenceOneTime    NewFloatingUnitPriceCadence = "one_time"
	NewFloatingUnitPriceCadenceCustom     NewFloatingUnitPriceCadence = "custom"
)

func (r NewFloatingUnitPriceCadence) IsKnown() bool {
	switch r {
	case NewFloatingUnitPriceCadenceAnnual, NewFloatingUnitPriceCadenceSemiAnnual, NewFloatingUnitPriceCadenceMonthly, NewFloatingUnitPriceCadenceQuarterly, NewFloatingUnitPriceCadenceOneTime, NewFloatingUnitPriceCadenceCustom:
		return true
	}
	return false
}

type NewFloatingUnitPriceModelType string

const (
	NewFloatingUnitPriceModelTypeUnit NewFloatingUnitPriceModelType = "unit"
)

func (r NewFloatingUnitPriceModelType) IsKnown() bool {
	switch r {
	case NewFloatingUnitPriceModelTypeUnit:
		return true
	}
	return false
}

type NewFloatingUnitPriceParam struct {
	ConversionRateType param.Field[NewFloatingUnitPriceConversionRateType] `json:"conversion_rate_type,required"`
	TieredConfig       param.Field[ConversionRateTieredConfigParam]        `json:"tiered_config"`
	UnitConfig         param.Field[ConversionRateUnitConfigParam]          `json:"unit_config"`
}

func (r NewFloatingUnitPriceParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r NewFloatingUnitPriceParam) ImplementsNewFloatingUnitPriceUnionParam() {}

// Satisfied by [shared.UnitConversionRateConfigParam],
// [shared.TieredConversionRateConfigParam], [NewFloatingUnitPriceParam].
type NewFloatingUnitPriceUnionParam interface {
	ImplementsNewFloatingUnitPriceUnionParam()
}

type NewFloatingUnitPriceConversionRateType string

const (
	NewFloatingUnitPriceConversionRateTypeUnit   NewFloatingUnitPriceConversionRateType = "unit"
	NewFloatingUnitPriceConversionRateTypeTiered NewFloatingUnitPriceConversionRateType = "tiered"
)

func (r NewFloatingUnitPriceConversionRateType) IsKnown() bool {
	switch r {
	case NewFloatingUnitPriceConversionRateTypeUnit, NewFloatingUnitPriceConversionRateTypeTiered:
		return true
	}
	return false
}

type NewFloatingUnitWithPercentPriceParam struct {
	// The cadence to bill for this price on.
	Cadence param.Field[NewFloatingUnitWithPercentPriceCadence] `json:"cadence,required"`
	// An ISO 4217 currency string for which this price is billed in.
	Currency param.Field[string] `json:"currency,required"`
	// The id of the item the price will be associated with.
	ItemID    param.Field[string]                                   `json:"item_id,required"`
	ModelType param.Field[NewFloatingUnitWithPercentPriceModelType] `json:"model_type,required"`
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
	BillingCycleConfiguration param.Field[NewBillingCycleConfigurationParam] `json:"billing_cycle_configuration"`
	// The per unit conversion rate of the price currency to the invoicing currency.
	ConversionRate param.Field[float64] `json:"conversion_rate"`
	// The configuration for the rate of the price currency to the invoicing currency.
	ConversionRateConfig param.Field[NewFloatingUnitWithPercentPriceUnionParam] `json:"conversion_rate_config"`
	// For dimensional price: specifies a price group and dimension values
	DimensionalPriceConfiguration param.Field[NewDimensionalPriceConfigurationParam] `json:"dimensional_price_configuration"`
	// An alias for the price.
	ExternalPriceID param.Field[string] `json:"external_price_id"`
	// If the Price represents a fixed cost, this represents the quantity of units
	// applied.
	FixedPriceQuantity param.Field[float64] `json:"fixed_price_quantity"`
	// The property used to group this price on an invoice
	InvoiceGroupingKey param.Field[string] `json:"invoice_grouping_key"`
	// Within each billing cycle, specifies the cadence at which invoices are produced.
	// If unspecified, a single invoice is produced per billing cycle.
	InvoicingCycleConfiguration param.Field[NewBillingCycleConfigurationParam] `json:"invoicing_cycle_configuration"`
	// User-specified key/value pairs for the resource. Individual keys can be removed
	// by setting the value to `null`, and the entire metadata mapping can be cleared
	// by setting `metadata` to `null`.
	Metadata param.Field[map[string]string] `json:"metadata"`
}

func (r NewFloatingUnitWithPercentPriceParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r NewFloatingUnitWithPercentPriceParam) ImplementsPriceEvaluateMultipleParamsPriceEvaluationsPriceUnion() {
}

func (r NewFloatingUnitWithPercentPriceParam) ImplementsPriceEvaluatePreviewEventsParamsPriceEvaluationsPriceUnion() {
}

func (r NewFloatingUnitWithPercentPriceParam) ImplementsSubscriptionPriceIntervalsParamsAddPriceUnion() {
}

// The cadence to bill for this price on.
type NewFloatingUnitWithPercentPriceCadence string

const (
	NewFloatingUnitWithPercentPriceCadenceAnnual     NewFloatingUnitWithPercentPriceCadence = "annual"
	NewFloatingUnitWithPercentPriceCadenceSemiAnnual NewFloatingUnitWithPercentPriceCadence = "semi_annual"
	NewFloatingUnitWithPercentPriceCadenceMonthly    NewFloatingUnitWithPercentPriceCadence = "monthly"
	NewFloatingUnitWithPercentPriceCadenceQuarterly  NewFloatingUnitWithPercentPriceCadence = "quarterly"
	NewFloatingUnitWithPercentPriceCadenceOneTime    NewFloatingUnitWithPercentPriceCadence = "one_time"
	NewFloatingUnitWithPercentPriceCadenceCustom     NewFloatingUnitWithPercentPriceCadence = "custom"
)

func (r NewFloatingUnitWithPercentPriceCadence) IsKnown() bool {
	switch r {
	case NewFloatingUnitWithPercentPriceCadenceAnnual, NewFloatingUnitWithPercentPriceCadenceSemiAnnual, NewFloatingUnitWithPercentPriceCadenceMonthly, NewFloatingUnitWithPercentPriceCadenceQuarterly, NewFloatingUnitWithPercentPriceCadenceOneTime, NewFloatingUnitWithPercentPriceCadenceCustom:
		return true
	}
	return false
}

type NewFloatingUnitWithPercentPriceModelType string

const (
	NewFloatingUnitWithPercentPriceModelTypeUnitWithPercent NewFloatingUnitWithPercentPriceModelType = "unit_with_percent"
)

func (r NewFloatingUnitWithPercentPriceModelType) IsKnown() bool {
	switch r {
	case NewFloatingUnitWithPercentPriceModelTypeUnitWithPercent:
		return true
	}
	return false
}

type NewFloatingUnitWithPercentPriceParam struct {
	ConversionRateType param.Field[NewFloatingUnitWithPercentPriceConversionRateType] `json:"conversion_rate_type,required"`
	TieredConfig       param.Field[ConversionRateTieredConfigParam]                   `json:"tiered_config"`
	UnitConfig         param.Field[ConversionRateUnitConfigParam]                     `json:"unit_config"`
}

func (r NewFloatingUnitWithPercentPriceParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r NewFloatingUnitWithPercentPriceParam) ImplementsNewFloatingUnitWithPercentPriceUnionParam() {}

// Satisfied by [shared.UnitConversionRateConfigParam],
// [shared.TieredConversionRateConfigParam],
// [NewFloatingUnitWithPercentPriceParam].
type NewFloatingUnitWithPercentPriceUnionParam interface {
	ImplementsNewFloatingUnitWithPercentPriceUnionParam()
}

type NewFloatingUnitWithPercentPriceConversionRateType string

const (
	NewFloatingUnitWithPercentPriceConversionRateTypeUnit   NewFloatingUnitWithPercentPriceConversionRateType = "unit"
	NewFloatingUnitWithPercentPriceConversionRateTypeTiered NewFloatingUnitWithPercentPriceConversionRateType = "tiered"
)

func (r NewFloatingUnitWithPercentPriceConversionRateType) IsKnown() bool {
	switch r {
	case NewFloatingUnitWithPercentPriceConversionRateTypeUnit, NewFloatingUnitWithPercentPriceConversionRateTypeTiered:
		return true
	}
	return false
}

type NewFloatingUnitWithProrationPriceParam struct {
	// The cadence to bill for this price on.
	Cadence param.Field[NewFloatingUnitWithProrationPriceCadence] `json:"cadence,required"`
	// An ISO 4217 currency string for which this price is billed in.
	Currency param.Field[string] `json:"currency,required"`
	// The id of the item the price will be associated with.
	ItemID    param.Field[string]                                     `json:"item_id,required"`
	ModelType param.Field[NewFloatingUnitWithProrationPriceModelType] `json:"model_type,required"`
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
	BillingCycleConfiguration param.Field[NewBillingCycleConfigurationParam] `json:"billing_cycle_configuration"`
	// The per unit conversion rate of the price currency to the invoicing currency.
	ConversionRate param.Field[float64] `json:"conversion_rate"`
	// The configuration for the rate of the price currency to the invoicing currency.
	ConversionRateConfig param.Field[NewFloatingUnitWithProrationPriceUnionParam] `json:"conversion_rate_config"`
	// For dimensional price: specifies a price group and dimension values
	DimensionalPriceConfiguration param.Field[NewDimensionalPriceConfigurationParam] `json:"dimensional_price_configuration"`
	// An alias for the price.
	ExternalPriceID param.Field[string] `json:"external_price_id"`
	// If the Price represents a fixed cost, this represents the quantity of units
	// applied.
	FixedPriceQuantity param.Field[float64] `json:"fixed_price_quantity"`
	// The property used to group this price on an invoice
	InvoiceGroupingKey param.Field[string] `json:"invoice_grouping_key"`
	// Within each billing cycle, specifies the cadence at which invoices are produced.
	// If unspecified, a single invoice is produced per billing cycle.
	InvoicingCycleConfiguration param.Field[NewBillingCycleConfigurationParam] `json:"invoicing_cycle_configuration"`
	// User-specified key/value pairs for the resource. Individual keys can be removed
	// by setting the value to `null`, and the entire metadata mapping can be cleared
	// by setting `metadata` to `null`.
	Metadata param.Field[map[string]string] `json:"metadata"`
}

func (r NewFloatingUnitWithProrationPriceParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r NewFloatingUnitWithProrationPriceParam) ImplementsPriceEvaluateMultipleParamsPriceEvaluationsPriceUnion() {
}

func (r NewFloatingUnitWithProrationPriceParam) ImplementsPriceEvaluatePreviewEventsParamsPriceEvaluationsPriceUnion() {
}

func (r NewFloatingUnitWithProrationPriceParam) ImplementsSubscriptionPriceIntervalsParamsAddPriceUnion() {
}

// The cadence to bill for this price on.
type NewFloatingUnitWithProrationPriceCadence string

const (
	NewFloatingUnitWithProrationPriceCadenceAnnual     NewFloatingUnitWithProrationPriceCadence = "annual"
	NewFloatingUnitWithProrationPriceCadenceSemiAnnual NewFloatingUnitWithProrationPriceCadence = "semi_annual"
	NewFloatingUnitWithProrationPriceCadenceMonthly    NewFloatingUnitWithProrationPriceCadence = "monthly"
	NewFloatingUnitWithProrationPriceCadenceQuarterly  NewFloatingUnitWithProrationPriceCadence = "quarterly"
	NewFloatingUnitWithProrationPriceCadenceOneTime    NewFloatingUnitWithProrationPriceCadence = "one_time"
	NewFloatingUnitWithProrationPriceCadenceCustom     NewFloatingUnitWithProrationPriceCadence = "custom"
)

func (r NewFloatingUnitWithProrationPriceCadence) IsKnown() bool {
	switch r {
	case NewFloatingUnitWithProrationPriceCadenceAnnual, NewFloatingUnitWithProrationPriceCadenceSemiAnnual, NewFloatingUnitWithProrationPriceCadenceMonthly, NewFloatingUnitWithProrationPriceCadenceQuarterly, NewFloatingUnitWithProrationPriceCadenceOneTime, NewFloatingUnitWithProrationPriceCadenceCustom:
		return true
	}
	return false
}

type NewFloatingUnitWithProrationPriceModelType string

const (
	NewFloatingUnitWithProrationPriceModelTypeUnitWithProration NewFloatingUnitWithProrationPriceModelType = "unit_with_proration"
)

func (r NewFloatingUnitWithProrationPriceModelType) IsKnown() bool {
	switch r {
	case NewFloatingUnitWithProrationPriceModelTypeUnitWithProration:
		return true
	}
	return false
}

type NewFloatingUnitWithProrationPriceParam struct {
	ConversionRateType param.Field[NewFloatingUnitWithProrationPriceConversionRateType] `json:"conversion_rate_type,required"`
	TieredConfig       param.Field[ConversionRateTieredConfigParam]                     `json:"tiered_config"`
	UnitConfig         param.Field[ConversionRateUnitConfigParam]                       `json:"unit_config"`
}

func (r NewFloatingUnitWithProrationPriceParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r NewFloatingUnitWithProrationPriceParam) ImplementsNewFloatingUnitWithProrationPriceUnionParam() {
}

// Satisfied by [shared.UnitConversionRateConfigParam],
// [shared.TieredConversionRateConfigParam],
// [NewFloatingUnitWithProrationPriceParam].
type NewFloatingUnitWithProrationPriceUnionParam interface {
	ImplementsNewFloatingUnitWithProrationPriceUnionParam()
}

type NewFloatingUnitWithProrationPriceConversionRateType string

const (
	NewFloatingUnitWithProrationPriceConversionRateTypeUnit   NewFloatingUnitWithProrationPriceConversionRateType = "unit"
	NewFloatingUnitWithProrationPriceConversionRateTypeTiered NewFloatingUnitWithProrationPriceConversionRateType = "tiered"
)

func (r NewFloatingUnitWithProrationPriceConversionRateType) IsKnown() bool {
	switch r {
	case NewFloatingUnitWithProrationPriceConversionRateTypeUnit, NewFloatingUnitWithProrationPriceConversionRateTypeTiered:
		return true
	}
	return false
}

type NewMaximumParam struct {
	AdjustmentType param.Field[NewMaximumAdjustmentType] `json:"adjustment_type,required"`
	MaximumAmount  param.Field[string]                   `json:"maximum_amount,required"`
	// If set, the adjustment will apply to every price on the subscription.
	AppliesToAll param.Field[NewMaximumAppliesToAll] `json:"applies_to_all"`
	// The set of item IDs to which this adjustment applies.
	AppliesToItemIDs param.Field[[]string] `json:"applies_to_item_ids"`
	// The set of price IDs to which this adjustment applies.
	AppliesToPriceIDs param.Field[[]string] `json:"applies_to_price_ids"`
	// If set, only prices in the specified currency will have the adjustment applied.
	Currency param.Field[string] `json:"currency"`
	// A list of filters that determine which prices this adjustment will apply to.
	Filters param.Field[[]TransformPriceFilterParam] `json:"filters"`
	// When false, this adjustment will be applied to a single price. Otherwise, it
	// will be applied at the invoice level, possibly to multiple prices.
	IsInvoiceLevel param.Field[bool] `json:"is_invoice_level"`
	// If set, only prices of the specified type will have the adjustment applied.
	PriceType param.Field[NewMaximumPriceType] `json:"price_type"`
}

func (r NewMaximumParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r NewMaximumParam) ImplementsBetaNewPlanVersionParamsAddAdjustmentsAdjustmentUnion() {}

func (r NewMaximumParam) ImplementsBetaNewPlanVersionParamsReplaceAdjustmentsAdjustmentUnion() {}

func (r NewMaximumParam) ImplementsBetaExternalPlanIDNewPlanVersionParamsAddAdjustmentsAdjustmentUnion() {
}

func (r NewMaximumParam) ImplementsBetaExternalPlanIDNewPlanVersionParamsReplaceAdjustmentsAdjustmentUnion() {
}

func (r NewMaximumParam) ImplementsPlanNewParamsAdjustmentsAdjustmentUnion() {}

func (r NewMaximumParam) ImplementsSubscriptionNewParamsAddAdjustmentsAdjustmentUnion() {}

func (r NewMaximumParam) ImplementsSubscriptionNewParamsReplaceAdjustmentsAdjustmentUnion() {}

func (r NewMaximumParam) ImplementsSubscriptionPriceIntervalsParamsAddAdjustmentsAdjustmentUnion() {}

func (r NewMaximumParam) ImplementsSubscriptionSchedulePlanChangeParamsAddAdjustmentsAdjustmentUnion() {
}

func (r NewMaximumParam) ImplementsSubscriptionSchedulePlanChangeParamsReplaceAdjustmentsAdjustmentUnion() {
}

type NewMaximumAdjustmentType string

const (
	NewMaximumAdjustmentTypeMaximum NewMaximumAdjustmentType = "maximum"
)

func (r NewMaximumAdjustmentType) IsKnown() bool {
	switch r {
	case NewMaximumAdjustmentTypeMaximum:
		return true
	}
	return false
}

// If set, the adjustment will apply to every price on the subscription.
type NewMaximumAppliesToAll bool

const (
	NewMaximumAppliesToAllTrue NewMaximumAppliesToAll = true
)

func (r NewMaximumAppliesToAll) IsKnown() bool {
	switch r {
	case NewMaximumAppliesToAllTrue:
		return true
	}
	return false
}

// If set, only prices of the specified type will have the adjustment applied.
type NewMaximumPriceType string

const (
	NewMaximumPriceTypeUsage          NewMaximumPriceType = "usage"
	NewMaximumPriceTypeFixedInAdvance NewMaximumPriceType = "fixed_in_advance"
	NewMaximumPriceTypeFixedInArrears NewMaximumPriceType = "fixed_in_arrears"
	NewMaximumPriceTypeFixed          NewMaximumPriceType = "fixed"
	NewMaximumPriceTypeInArrears      NewMaximumPriceType = "in_arrears"
)

func (r NewMaximumPriceType) IsKnown() bool {
	switch r {
	case NewMaximumPriceTypeUsage, NewMaximumPriceTypeFixedInAdvance, NewMaximumPriceTypeFixedInArrears, NewMaximumPriceTypeFixed, NewMaximumPriceTypeInArrears:
		return true
	}
	return false
}

type NewMinimumParam struct {
	AdjustmentType param.Field[NewMinimumAdjustmentType] `json:"adjustment_type,required"`
	// The item ID that revenue from this minimum will be attributed to.
	ItemID        param.Field[string] `json:"item_id,required"`
	MinimumAmount param.Field[string] `json:"minimum_amount,required"`
	// If set, the adjustment will apply to every price on the subscription.
	AppliesToAll param.Field[NewMinimumAppliesToAll] `json:"applies_to_all"`
	// The set of item IDs to which this adjustment applies.
	AppliesToItemIDs param.Field[[]string] `json:"applies_to_item_ids"`
	// The set of price IDs to which this adjustment applies.
	AppliesToPriceIDs param.Field[[]string] `json:"applies_to_price_ids"`
	// If set, only prices in the specified currency will have the adjustment applied.
	Currency param.Field[string] `json:"currency"`
	// A list of filters that determine which prices this adjustment will apply to.
	Filters param.Field[[]TransformPriceFilterParam] `json:"filters"`
	// When false, this adjustment will be applied to a single price. Otherwise, it
	// will be applied at the invoice level, possibly to multiple prices.
	IsInvoiceLevel param.Field[bool] `json:"is_invoice_level"`
	// If set, only prices of the specified type will have the adjustment applied.
	PriceType param.Field[NewMinimumPriceType] `json:"price_type"`
}

func (r NewMinimumParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r NewMinimumParam) ImplementsBetaNewPlanVersionParamsAddAdjustmentsAdjustmentUnion() {}

func (r NewMinimumParam) ImplementsBetaNewPlanVersionParamsReplaceAdjustmentsAdjustmentUnion() {}

func (r NewMinimumParam) ImplementsBetaExternalPlanIDNewPlanVersionParamsAddAdjustmentsAdjustmentUnion() {
}

func (r NewMinimumParam) ImplementsBetaExternalPlanIDNewPlanVersionParamsReplaceAdjustmentsAdjustmentUnion() {
}

func (r NewMinimumParam) ImplementsPlanNewParamsAdjustmentsAdjustmentUnion() {}

func (r NewMinimumParam) ImplementsSubscriptionNewParamsAddAdjustmentsAdjustmentUnion() {}

func (r NewMinimumParam) ImplementsSubscriptionNewParamsReplaceAdjustmentsAdjustmentUnion() {}

func (r NewMinimumParam) ImplementsSubscriptionPriceIntervalsParamsAddAdjustmentsAdjustmentUnion() {}

func (r NewMinimumParam) ImplementsSubscriptionSchedulePlanChangeParamsAddAdjustmentsAdjustmentUnion() {
}

func (r NewMinimumParam) ImplementsSubscriptionSchedulePlanChangeParamsReplaceAdjustmentsAdjustmentUnion() {
}

type NewMinimumAdjustmentType string

const (
	NewMinimumAdjustmentTypeMinimum NewMinimumAdjustmentType = "minimum"
)

func (r NewMinimumAdjustmentType) IsKnown() bool {
	switch r {
	case NewMinimumAdjustmentTypeMinimum:
		return true
	}
	return false
}

// If set, the adjustment will apply to every price on the subscription.
type NewMinimumAppliesToAll bool

const (
	NewMinimumAppliesToAllTrue NewMinimumAppliesToAll = true
)

func (r NewMinimumAppliesToAll) IsKnown() bool {
	switch r {
	case NewMinimumAppliesToAllTrue:
		return true
	}
	return false
}

// If set, only prices of the specified type will have the adjustment applied.
type NewMinimumPriceType string

const (
	NewMinimumPriceTypeUsage          NewMinimumPriceType = "usage"
	NewMinimumPriceTypeFixedInAdvance NewMinimumPriceType = "fixed_in_advance"
	NewMinimumPriceTypeFixedInArrears NewMinimumPriceType = "fixed_in_arrears"
	NewMinimumPriceTypeFixed          NewMinimumPriceType = "fixed"
	NewMinimumPriceTypeInArrears      NewMinimumPriceType = "in_arrears"
)

func (r NewMinimumPriceType) IsKnown() bool {
	switch r {
	case NewMinimumPriceTypeUsage, NewMinimumPriceTypeFixedInAdvance, NewMinimumPriceTypeFixedInArrears, NewMinimumPriceTypeFixed, NewMinimumPriceTypeInArrears:
		return true
	}
	return false
}

type NewPercentageDiscountParam struct {
	AdjustmentType     param.Field[NewPercentageDiscountAdjustmentType] `json:"adjustment_type,required"`
	PercentageDiscount param.Field[float64]                             `json:"percentage_discount,required"`
	// If set, the adjustment will apply to every price on the subscription.
	AppliesToAll param.Field[NewPercentageDiscountAppliesToAll] `json:"applies_to_all"`
	// The set of item IDs to which this adjustment applies.
	AppliesToItemIDs param.Field[[]string] `json:"applies_to_item_ids"`
	// The set of price IDs to which this adjustment applies.
	AppliesToPriceIDs param.Field[[]string] `json:"applies_to_price_ids"`
	// If set, only prices in the specified currency will have the adjustment applied.
	Currency param.Field[string] `json:"currency"`
	// A list of filters that determine which prices this adjustment will apply to.
	Filters param.Field[[]TransformPriceFilterParam] `json:"filters"`
	// When false, this adjustment will be applied to a single price. Otherwise, it
	// will be applied at the invoice level, possibly to multiple prices.
	IsInvoiceLevel param.Field[bool] `json:"is_invoice_level"`
	// If set, only prices of the specified type will have the adjustment applied.
	PriceType param.Field[NewPercentageDiscountPriceType] `json:"price_type"`
}

func (r NewPercentageDiscountParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r NewPercentageDiscountParam) ImplementsBetaNewPlanVersionParamsAddAdjustmentsAdjustmentUnion() {
}

func (r NewPercentageDiscountParam) ImplementsBetaNewPlanVersionParamsReplaceAdjustmentsAdjustmentUnion() {
}

func (r NewPercentageDiscountParam) ImplementsBetaExternalPlanIDNewPlanVersionParamsAddAdjustmentsAdjustmentUnion() {
}

func (r NewPercentageDiscountParam) ImplementsBetaExternalPlanIDNewPlanVersionParamsReplaceAdjustmentsAdjustmentUnion() {
}

func (r NewPercentageDiscountParam) ImplementsPlanNewParamsAdjustmentsAdjustmentUnion() {}

func (r NewPercentageDiscountParam) ImplementsSubscriptionNewParamsAddAdjustmentsAdjustmentUnion() {}

func (r NewPercentageDiscountParam) ImplementsSubscriptionNewParamsReplaceAdjustmentsAdjustmentUnion() {
}

func (r NewPercentageDiscountParam) ImplementsSubscriptionPriceIntervalsParamsAddAdjustmentsAdjustmentUnion() {
}

func (r NewPercentageDiscountParam) ImplementsSubscriptionSchedulePlanChangeParamsAddAdjustmentsAdjustmentUnion() {
}

func (r NewPercentageDiscountParam) ImplementsSubscriptionSchedulePlanChangeParamsReplaceAdjustmentsAdjustmentUnion() {
}

type NewPercentageDiscountAdjustmentType string

const (
	NewPercentageDiscountAdjustmentTypePercentageDiscount NewPercentageDiscountAdjustmentType = "percentage_discount"
)

func (r NewPercentageDiscountAdjustmentType) IsKnown() bool {
	switch r {
	case NewPercentageDiscountAdjustmentTypePercentageDiscount:
		return true
	}
	return false
}

// If set, the adjustment will apply to every price on the subscription.
type NewPercentageDiscountAppliesToAll bool

const (
	NewPercentageDiscountAppliesToAllTrue NewPercentageDiscountAppliesToAll = true
)

func (r NewPercentageDiscountAppliesToAll) IsKnown() bool {
	switch r {
	case NewPercentageDiscountAppliesToAllTrue:
		return true
	}
	return false
}

// If set, only prices of the specified type will have the adjustment applied.
type NewPercentageDiscountPriceType string

const (
	NewPercentageDiscountPriceTypeUsage          NewPercentageDiscountPriceType = "usage"
	NewPercentageDiscountPriceTypeFixedInAdvance NewPercentageDiscountPriceType = "fixed_in_advance"
	NewPercentageDiscountPriceTypeFixedInArrears NewPercentageDiscountPriceType = "fixed_in_arrears"
	NewPercentageDiscountPriceTypeFixed          NewPercentageDiscountPriceType = "fixed"
	NewPercentageDiscountPriceTypeInArrears      NewPercentageDiscountPriceType = "in_arrears"
)

func (r NewPercentageDiscountPriceType) IsKnown() bool {
	switch r {
	case NewPercentageDiscountPriceTypeUsage, NewPercentageDiscountPriceTypeFixedInAdvance, NewPercentageDiscountPriceTypeFixedInArrears, NewPercentageDiscountPriceTypeFixed, NewPercentageDiscountPriceTypeInArrears:
		return true
	}
	return false
}

type NewPlanBPSPriceParam struct {
	BPSConfig param.Field[BPSConfigParam] `json:"bps_config,required"`
	// The cadence to bill for this price on.
	Cadence param.Field[NewPlanBPSPriceCadence] `json:"cadence,required"`
	// The id of the item the price will be associated with.
	ItemID    param.Field[string]                   `json:"item_id,required"`
	ModelType param.Field[NewPlanBPSPriceModelType] `json:"model_type,required"`
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
	BillingCycleConfiguration param.Field[NewBillingCycleConfigurationParam] `json:"billing_cycle_configuration"`
	// The per unit conversion rate of the price currency to the invoicing currency.
	ConversionRate param.Field[float64] `json:"conversion_rate"`
	// The configuration for the rate of the price currency to the invoicing currency.
	ConversionRateConfig param.Field[NewPlanBPSPriceUnionParam] `json:"conversion_rate_config"`
	// An ISO 4217 currency string, or custom pricing unit identifier, in which this
	// price is billed.
	Currency param.Field[string] `json:"currency"`
	// For dimensional price: specifies a price group and dimension values
	DimensionalPriceConfiguration param.Field[NewDimensionalPriceConfigurationParam] `json:"dimensional_price_configuration"`
	// An alias for the price.
	ExternalPriceID param.Field[string] `json:"external_price_id"`
	// If the Price represents a fixed cost, this represents the quantity of units
	// applied.
	FixedPriceQuantity param.Field[float64] `json:"fixed_price_quantity"`
	// The property used to group this price on an invoice
	InvoiceGroupingKey param.Field[string] `json:"invoice_grouping_key"`
	// Within each billing cycle, specifies the cadence at which invoices are produced.
	// If unspecified, a single invoice is produced per billing cycle.
	InvoicingCycleConfiguration param.Field[NewBillingCycleConfigurationParam] `json:"invoicing_cycle_configuration"`
	// User-specified key/value pairs for the resource. Individual keys can be removed
	// by setting the value to `null`, and the entire metadata mapping can be cleared
	// by setting `metadata` to `null`.
	Metadata param.Field[map[string]string] `json:"metadata"`
	// A transient ID that can be used to reference this price when adding adjustments
	// in the same API call.
	ReferenceID param.Field[string] `json:"reference_id"`
}

func (r NewPlanBPSPriceParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r NewPlanBPSPriceParam) ImplementsBetaNewPlanVersionParamsAddPricesPriceUnion() {}

func (r NewPlanBPSPriceParam) ImplementsBetaNewPlanVersionParamsReplacePricesPriceUnion() {}

func (r NewPlanBPSPriceParam) ImplementsBetaExternalPlanIDNewPlanVersionParamsAddPricesPriceUnion() {}

func (r NewPlanBPSPriceParam) ImplementsBetaExternalPlanIDNewPlanVersionParamsReplacePricesPriceUnion() {
}

func (r NewPlanBPSPriceParam) ImplementsPlanNewParamsPricesPriceUnion() {}

// The cadence to bill for this price on.
type NewPlanBPSPriceCadence string

const (
	NewPlanBPSPriceCadenceAnnual     NewPlanBPSPriceCadence = "annual"
	NewPlanBPSPriceCadenceSemiAnnual NewPlanBPSPriceCadence = "semi_annual"
	NewPlanBPSPriceCadenceMonthly    NewPlanBPSPriceCadence = "monthly"
	NewPlanBPSPriceCadenceQuarterly  NewPlanBPSPriceCadence = "quarterly"
	NewPlanBPSPriceCadenceOneTime    NewPlanBPSPriceCadence = "one_time"
	NewPlanBPSPriceCadenceCustom     NewPlanBPSPriceCadence = "custom"
)

func (r NewPlanBPSPriceCadence) IsKnown() bool {
	switch r {
	case NewPlanBPSPriceCadenceAnnual, NewPlanBPSPriceCadenceSemiAnnual, NewPlanBPSPriceCadenceMonthly, NewPlanBPSPriceCadenceQuarterly, NewPlanBPSPriceCadenceOneTime, NewPlanBPSPriceCadenceCustom:
		return true
	}
	return false
}

type NewPlanBPSPriceModelType string

const (
	NewPlanBPSPriceModelTypeBPS NewPlanBPSPriceModelType = "bps"
)

func (r NewPlanBPSPriceModelType) IsKnown() bool {
	switch r {
	case NewPlanBPSPriceModelTypeBPS:
		return true
	}
	return false
}

type NewPlanBPSPriceParam struct {
	ConversionRateType param.Field[NewPlanBPSPriceConversionRateType] `json:"conversion_rate_type,required"`
	TieredConfig       param.Field[ConversionRateTieredConfigParam]   `json:"tiered_config"`
	UnitConfig         param.Field[ConversionRateUnitConfigParam]     `json:"unit_config"`
}

func (r NewPlanBPSPriceParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r NewPlanBPSPriceParam) ImplementsNewPlanBPSPriceUnionParam() {}

// Satisfied by [shared.UnitConversionRateConfigParam],
// [shared.TieredConversionRateConfigParam], [NewPlanBPSPriceParam].
type NewPlanBPSPriceUnionParam interface {
	ImplementsNewPlanBPSPriceUnionParam()
}

type NewPlanBPSPriceConversionRateType string

const (
	NewPlanBPSPriceConversionRateTypeUnit   NewPlanBPSPriceConversionRateType = "unit"
	NewPlanBPSPriceConversionRateTypeTiered NewPlanBPSPriceConversionRateType = "tiered"
)

func (r NewPlanBPSPriceConversionRateType) IsKnown() bool {
	switch r {
	case NewPlanBPSPriceConversionRateTypeUnit, NewPlanBPSPriceConversionRateTypeTiered:
		return true
	}
	return false
}

type NewPlanBulkBPSPriceParam struct {
	BulkBPSConfig param.Field[BulkBPSConfigParam] `json:"bulk_bps_config,required"`
	// The cadence to bill for this price on.
	Cadence param.Field[NewPlanBulkBPSPriceCadence] `json:"cadence,required"`
	// The id of the item the price will be associated with.
	ItemID    param.Field[string]                       `json:"item_id,required"`
	ModelType param.Field[NewPlanBulkBPSPriceModelType] `json:"model_type,required"`
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
	BillingCycleConfiguration param.Field[NewBillingCycleConfigurationParam] `json:"billing_cycle_configuration"`
	// The per unit conversion rate of the price currency to the invoicing currency.
	ConversionRate param.Field[float64] `json:"conversion_rate"`
	// The configuration for the rate of the price currency to the invoicing currency.
	ConversionRateConfig param.Field[NewPlanBulkBPSPriceUnionParam] `json:"conversion_rate_config"`
	// An ISO 4217 currency string, or custom pricing unit identifier, in which this
	// price is billed.
	Currency param.Field[string] `json:"currency"`
	// For dimensional price: specifies a price group and dimension values
	DimensionalPriceConfiguration param.Field[NewDimensionalPriceConfigurationParam] `json:"dimensional_price_configuration"`
	// An alias for the price.
	ExternalPriceID param.Field[string] `json:"external_price_id"`
	// If the Price represents a fixed cost, this represents the quantity of units
	// applied.
	FixedPriceQuantity param.Field[float64] `json:"fixed_price_quantity"`
	// The property used to group this price on an invoice
	InvoiceGroupingKey param.Field[string] `json:"invoice_grouping_key"`
	// Within each billing cycle, specifies the cadence at which invoices are produced.
	// If unspecified, a single invoice is produced per billing cycle.
	InvoicingCycleConfiguration param.Field[NewBillingCycleConfigurationParam] `json:"invoicing_cycle_configuration"`
	// User-specified key/value pairs for the resource. Individual keys can be removed
	// by setting the value to `null`, and the entire metadata mapping can be cleared
	// by setting `metadata` to `null`.
	Metadata param.Field[map[string]string] `json:"metadata"`
	// A transient ID that can be used to reference this price when adding adjustments
	// in the same API call.
	ReferenceID param.Field[string] `json:"reference_id"`
}

func (r NewPlanBulkBPSPriceParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r NewPlanBulkBPSPriceParam) ImplementsBetaNewPlanVersionParamsAddPricesPriceUnion() {}

func (r NewPlanBulkBPSPriceParam) ImplementsBetaNewPlanVersionParamsReplacePricesPriceUnion() {}

func (r NewPlanBulkBPSPriceParam) ImplementsBetaExternalPlanIDNewPlanVersionParamsAddPricesPriceUnion() {
}

func (r NewPlanBulkBPSPriceParam) ImplementsBetaExternalPlanIDNewPlanVersionParamsReplacePricesPriceUnion() {
}

func (r NewPlanBulkBPSPriceParam) ImplementsPlanNewParamsPricesPriceUnion() {}

// The cadence to bill for this price on.
type NewPlanBulkBPSPriceCadence string

const (
	NewPlanBulkBPSPriceCadenceAnnual     NewPlanBulkBPSPriceCadence = "annual"
	NewPlanBulkBPSPriceCadenceSemiAnnual NewPlanBulkBPSPriceCadence = "semi_annual"
	NewPlanBulkBPSPriceCadenceMonthly    NewPlanBulkBPSPriceCadence = "monthly"
	NewPlanBulkBPSPriceCadenceQuarterly  NewPlanBulkBPSPriceCadence = "quarterly"
	NewPlanBulkBPSPriceCadenceOneTime    NewPlanBulkBPSPriceCadence = "one_time"
	NewPlanBulkBPSPriceCadenceCustom     NewPlanBulkBPSPriceCadence = "custom"
)

func (r NewPlanBulkBPSPriceCadence) IsKnown() bool {
	switch r {
	case NewPlanBulkBPSPriceCadenceAnnual, NewPlanBulkBPSPriceCadenceSemiAnnual, NewPlanBulkBPSPriceCadenceMonthly, NewPlanBulkBPSPriceCadenceQuarterly, NewPlanBulkBPSPriceCadenceOneTime, NewPlanBulkBPSPriceCadenceCustom:
		return true
	}
	return false
}

type NewPlanBulkBPSPriceModelType string

const (
	NewPlanBulkBPSPriceModelTypeBulkBPS NewPlanBulkBPSPriceModelType = "bulk_bps"
)

func (r NewPlanBulkBPSPriceModelType) IsKnown() bool {
	switch r {
	case NewPlanBulkBPSPriceModelTypeBulkBPS:
		return true
	}
	return false
}

type NewPlanBulkBPSPriceParam struct {
	ConversionRateType param.Field[NewPlanBulkBPSPriceConversionRateType] `json:"conversion_rate_type,required"`
	TieredConfig       param.Field[ConversionRateTieredConfigParam]       `json:"tiered_config"`
	UnitConfig         param.Field[ConversionRateUnitConfigParam]         `json:"unit_config"`
}

func (r NewPlanBulkBPSPriceParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r NewPlanBulkBPSPriceParam) ImplementsNewPlanBulkBPSPriceUnionParam() {}

// Satisfied by [shared.UnitConversionRateConfigParam],
// [shared.TieredConversionRateConfigParam], [NewPlanBulkBPSPriceParam].
type NewPlanBulkBPSPriceUnionParam interface {
	ImplementsNewPlanBulkBPSPriceUnionParam()
}

type NewPlanBulkBPSPriceConversionRateType string

const (
	NewPlanBulkBPSPriceConversionRateTypeUnit   NewPlanBulkBPSPriceConversionRateType = "unit"
	NewPlanBulkBPSPriceConversionRateTypeTiered NewPlanBulkBPSPriceConversionRateType = "tiered"
)

func (r NewPlanBulkBPSPriceConversionRateType) IsKnown() bool {
	switch r {
	case NewPlanBulkBPSPriceConversionRateTypeUnit, NewPlanBulkBPSPriceConversionRateTypeTiered:
		return true
	}
	return false
}

type NewPlanBulkPriceParam struct {
	BulkConfig param.Field[BulkConfigParam] `json:"bulk_config,required"`
	// The cadence to bill for this price on.
	Cadence param.Field[NewPlanBulkPriceCadence] `json:"cadence,required"`
	// The id of the item the price will be associated with.
	ItemID    param.Field[string]                    `json:"item_id,required"`
	ModelType param.Field[NewPlanBulkPriceModelType] `json:"model_type,required"`
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
	BillingCycleConfiguration param.Field[NewBillingCycleConfigurationParam] `json:"billing_cycle_configuration"`
	// The per unit conversion rate of the price currency to the invoicing currency.
	ConversionRate param.Field[float64] `json:"conversion_rate"`
	// The configuration for the rate of the price currency to the invoicing currency.
	ConversionRateConfig param.Field[NewPlanBulkPriceUnionParam] `json:"conversion_rate_config"`
	// An ISO 4217 currency string, or custom pricing unit identifier, in which this
	// price is billed.
	Currency param.Field[string] `json:"currency"`
	// For dimensional price: specifies a price group and dimension values
	DimensionalPriceConfiguration param.Field[NewDimensionalPriceConfigurationParam] `json:"dimensional_price_configuration"`
	// An alias for the price.
	ExternalPriceID param.Field[string] `json:"external_price_id"`
	// If the Price represents a fixed cost, this represents the quantity of units
	// applied.
	FixedPriceQuantity param.Field[float64] `json:"fixed_price_quantity"`
	// The property used to group this price on an invoice
	InvoiceGroupingKey param.Field[string] `json:"invoice_grouping_key"`
	// Within each billing cycle, specifies the cadence at which invoices are produced.
	// If unspecified, a single invoice is produced per billing cycle.
	InvoicingCycleConfiguration param.Field[NewBillingCycleConfigurationParam] `json:"invoicing_cycle_configuration"`
	// User-specified key/value pairs for the resource. Individual keys can be removed
	// by setting the value to `null`, and the entire metadata mapping can be cleared
	// by setting `metadata` to `null`.
	Metadata param.Field[map[string]string] `json:"metadata"`
	// A transient ID that can be used to reference this price when adding adjustments
	// in the same API call.
	ReferenceID param.Field[string] `json:"reference_id"`
}

func (r NewPlanBulkPriceParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r NewPlanBulkPriceParam) ImplementsBetaNewPlanVersionParamsAddPricesPriceUnion() {}

func (r NewPlanBulkPriceParam) ImplementsBetaNewPlanVersionParamsReplacePricesPriceUnion() {}

func (r NewPlanBulkPriceParam) ImplementsBetaExternalPlanIDNewPlanVersionParamsAddPricesPriceUnion() {
}

func (r NewPlanBulkPriceParam) ImplementsBetaExternalPlanIDNewPlanVersionParamsReplacePricesPriceUnion() {
}

func (r NewPlanBulkPriceParam) ImplementsPlanNewParamsPricesPriceUnion() {}

// The cadence to bill for this price on.
type NewPlanBulkPriceCadence string

const (
	NewPlanBulkPriceCadenceAnnual     NewPlanBulkPriceCadence = "annual"
	NewPlanBulkPriceCadenceSemiAnnual NewPlanBulkPriceCadence = "semi_annual"
	NewPlanBulkPriceCadenceMonthly    NewPlanBulkPriceCadence = "monthly"
	NewPlanBulkPriceCadenceQuarterly  NewPlanBulkPriceCadence = "quarterly"
	NewPlanBulkPriceCadenceOneTime    NewPlanBulkPriceCadence = "one_time"
	NewPlanBulkPriceCadenceCustom     NewPlanBulkPriceCadence = "custom"
)

func (r NewPlanBulkPriceCadence) IsKnown() bool {
	switch r {
	case NewPlanBulkPriceCadenceAnnual, NewPlanBulkPriceCadenceSemiAnnual, NewPlanBulkPriceCadenceMonthly, NewPlanBulkPriceCadenceQuarterly, NewPlanBulkPriceCadenceOneTime, NewPlanBulkPriceCadenceCustom:
		return true
	}
	return false
}

type NewPlanBulkPriceModelType string

const (
	NewPlanBulkPriceModelTypeBulk NewPlanBulkPriceModelType = "bulk"
)

func (r NewPlanBulkPriceModelType) IsKnown() bool {
	switch r {
	case NewPlanBulkPriceModelTypeBulk:
		return true
	}
	return false
}

type NewPlanBulkPriceParam struct {
	ConversionRateType param.Field[NewPlanBulkPriceConversionRateType] `json:"conversion_rate_type,required"`
	TieredConfig       param.Field[ConversionRateTieredConfigParam]    `json:"tiered_config"`
	UnitConfig         param.Field[ConversionRateUnitConfigParam]      `json:"unit_config"`
}

func (r NewPlanBulkPriceParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r NewPlanBulkPriceParam) ImplementsNewPlanBulkPriceUnionParam() {}

// Satisfied by [shared.UnitConversionRateConfigParam],
// [shared.TieredConversionRateConfigParam], [NewPlanBulkPriceParam].
type NewPlanBulkPriceUnionParam interface {
	ImplementsNewPlanBulkPriceUnionParam()
}

type NewPlanBulkPriceConversionRateType string

const (
	NewPlanBulkPriceConversionRateTypeUnit   NewPlanBulkPriceConversionRateType = "unit"
	NewPlanBulkPriceConversionRateTypeTiered NewPlanBulkPriceConversionRateType = "tiered"
)

func (r NewPlanBulkPriceConversionRateType) IsKnown() bool {
	switch r {
	case NewPlanBulkPriceConversionRateTypeUnit, NewPlanBulkPriceConversionRateTypeTiered:
		return true
	}
	return false
}

type NewPlanBulkWithProrationPriceParam struct {
	BulkWithProrationConfig param.Field[map[string]interface{}] `json:"bulk_with_proration_config,required"`
	// The cadence to bill for this price on.
	Cadence param.Field[NewPlanBulkWithProrationPriceCadence] `json:"cadence,required"`
	// The id of the item the price will be associated with.
	ItemID    param.Field[string]                                 `json:"item_id,required"`
	ModelType param.Field[NewPlanBulkWithProrationPriceModelType] `json:"model_type,required"`
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
	BillingCycleConfiguration param.Field[NewBillingCycleConfigurationParam] `json:"billing_cycle_configuration"`
	// The per unit conversion rate of the price currency to the invoicing currency.
	ConversionRate param.Field[float64] `json:"conversion_rate"`
	// The configuration for the rate of the price currency to the invoicing currency.
	ConversionRateConfig param.Field[NewPlanBulkWithProrationPriceUnionParam] `json:"conversion_rate_config"`
	// An ISO 4217 currency string, or custom pricing unit identifier, in which this
	// price is billed.
	Currency param.Field[string] `json:"currency"`
	// For dimensional price: specifies a price group and dimension values
	DimensionalPriceConfiguration param.Field[NewDimensionalPriceConfigurationParam] `json:"dimensional_price_configuration"`
	// An alias for the price.
	ExternalPriceID param.Field[string] `json:"external_price_id"`
	// If the Price represents a fixed cost, this represents the quantity of units
	// applied.
	FixedPriceQuantity param.Field[float64] `json:"fixed_price_quantity"`
	// The property used to group this price on an invoice
	InvoiceGroupingKey param.Field[string] `json:"invoice_grouping_key"`
	// Within each billing cycle, specifies the cadence at which invoices are produced.
	// If unspecified, a single invoice is produced per billing cycle.
	InvoicingCycleConfiguration param.Field[NewBillingCycleConfigurationParam] `json:"invoicing_cycle_configuration"`
	// User-specified key/value pairs for the resource. Individual keys can be removed
	// by setting the value to `null`, and the entire metadata mapping can be cleared
	// by setting `metadata` to `null`.
	Metadata param.Field[map[string]string] `json:"metadata"`
	// A transient ID that can be used to reference this price when adding adjustments
	// in the same API call.
	ReferenceID param.Field[string] `json:"reference_id"`
}

func (r NewPlanBulkWithProrationPriceParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r NewPlanBulkWithProrationPriceParam) ImplementsBetaNewPlanVersionParamsAddPricesPriceUnion() {}

func (r NewPlanBulkWithProrationPriceParam) ImplementsBetaNewPlanVersionParamsReplacePricesPriceUnion() {
}

func (r NewPlanBulkWithProrationPriceParam) ImplementsBetaExternalPlanIDNewPlanVersionParamsAddPricesPriceUnion() {
}

func (r NewPlanBulkWithProrationPriceParam) ImplementsBetaExternalPlanIDNewPlanVersionParamsReplacePricesPriceUnion() {
}

func (r NewPlanBulkWithProrationPriceParam) ImplementsPlanNewParamsPricesPriceUnion() {}

// The cadence to bill for this price on.
type NewPlanBulkWithProrationPriceCadence string

const (
	NewPlanBulkWithProrationPriceCadenceAnnual     NewPlanBulkWithProrationPriceCadence = "annual"
	NewPlanBulkWithProrationPriceCadenceSemiAnnual NewPlanBulkWithProrationPriceCadence = "semi_annual"
	NewPlanBulkWithProrationPriceCadenceMonthly    NewPlanBulkWithProrationPriceCadence = "monthly"
	NewPlanBulkWithProrationPriceCadenceQuarterly  NewPlanBulkWithProrationPriceCadence = "quarterly"
	NewPlanBulkWithProrationPriceCadenceOneTime    NewPlanBulkWithProrationPriceCadence = "one_time"
	NewPlanBulkWithProrationPriceCadenceCustom     NewPlanBulkWithProrationPriceCadence = "custom"
)

func (r NewPlanBulkWithProrationPriceCadence) IsKnown() bool {
	switch r {
	case NewPlanBulkWithProrationPriceCadenceAnnual, NewPlanBulkWithProrationPriceCadenceSemiAnnual, NewPlanBulkWithProrationPriceCadenceMonthly, NewPlanBulkWithProrationPriceCadenceQuarterly, NewPlanBulkWithProrationPriceCadenceOneTime, NewPlanBulkWithProrationPriceCadenceCustom:
		return true
	}
	return false
}

type NewPlanBulkWithProrationPriceModelType string

const (
	NewPlanBulkWithProrationPriceModelTypeBulkWithProration NewPlanBulkWithProrationPriceModelType = "bulk_with_proration"
)

func (r NewPlanBulkWithProrationPriceModelType) IsKnown() bool {
	switch r {
	case NewPlanBulkWithProrationPriceModelTypeBulkWithProration:
		return true
	}
	return false
}

type NewPlanBulkWithProrationPriceParam struct {
	ConversionRateType param.Field[NewPlanBulkWithProrationPriceConversionRateType] `json:"conversion_rate_type,required"`
	TieredConfig       param.Field[ConversionRateTieredConfigParam]                 `json:"tiered_config"`
	UnitConfig         param.Field[ConversionRateUnitConfigParam]                   `json:"unit_config"`
}

func (r NewPlanBulkWithProrationPriceParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r NewPlanBulkWithProrationPriceParam) ImplementsNewPlanBulkWithProrationPriceUnionParam() {}

// Satisfied by [shared.UnitConversionRateConfigParam],
// [shared.TieredConversionRateConfigParam], [NewPlanBulkWithProrationPriceParam].
type NewPlanBulkWithProrationPriceUnionParam interface {
	ImplementsNewPlanBulkWithProrationPriceUnionParam()
}

type NewPlanBulkWithProrationPriceConversionRateType string

const (
	NewPlanBulkWithProrationPriceConversionRateTypeUnit   NewPlanBulkWithProrationPriceConversionRateType = "unit"
	NewPlanBulkWithProrationPriceConversionRateTypeTiered NewPlanBulkWithProrationPriceConversionRateType = "tiered"
)

func (r NewPlanBulkWithProrationPriceConversionRateType) IsKnown() bool {
	switch r {
	case NewPlanBulkWithProrationPriceConversionRateTypeUnit, NewPlanBulkWithProrationPriceConversionRateTypeTiered:
		return true
	}
	return false
}

type NewPlanCumulativeGroupedBulkPriceParam struct {
	// The cadence to bill for this price on.
	Cadence                     param.Field[NewPlanCumulativeGroupedBulkPriceCadence] `json:"cadence,required"`
	CumulativeGroupedBulkConfig param.Field[map[string]interface{}]                   `json:"cumulative_grouped_bulk_config,required"`
	// The id of the item the price will be associated with.
	ItemID    param.Field[string]                                     `json:"item_id,required"`
	ModelType param.Field[NewPlanCumulativeGroupedBulkPriceModelType] `json:"model_type,required"`
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
	BillingCycleConfiguration param.Field[NewBillingCycleConfigurationParam] `json:"billing_cycle_configuration"`
	// The per unit conversion rate of the price currency to the invoicing currency.
	ConversionRate param.Field[float64] `json:"conversion_rate"`
	// The configuration for the rate of the price currency to the invoicing currency.
	ConversionRateConfig param.Field[NewPlanCumulativeGroupedBulkPriceUnionParam] `json:"conversion_rate_config"`
	// An ISO 4217 currency string, or custom pricing unit identifier, in which this
	// price is billed.
	Currency param.Field[string] `json:"currency"`
	// For dimensional price: specifies a price group and dimension values
	DimensionalPriceConfiguration param.Field[NewDimensionalPriceConfigurationParam] `json:"dimensional_price_configuration"`
	// An alias for the price.
	ExternalPriceID param.Field[string] `json:"external_price_id"`
	// If the Price represents a fixed cost, this represents the quantity of units
	// applied.
	FixedPriceQuantity param.Field[float64] `json:"fixed_price_quantity"`
	// The property used to group this price on an invoice
	InvoiceGroupingKey param.Field[string] `json:"invoice_grouping_key"`
	// Within each billing cycle, specifies the cadence at which invoices are produced.
	// If unspecified, a single invoice is produced per billing cycle.
	InvoicingCycleConfiguration param.Field[NewBillingCycleConfigurationParam] `json:"invoicing_cycle_configuration"`
	// User-specified key/value pairs for the resource. Individual keys can be removed
	// by setting the value to `null`, and the entire metadata mapping can be cleared
	// by setting `metadata` to `null`.
	Metadata param.Field[map[string]string] `json:"metadata"`
	// A transient ID that can be used to reference this price when adding adjustments
	// in the same API call.
	ReferenceID param.Field[string] `json:"reference_id"`
}

func (r NewPlanCumulativeGroupedBulkPriceParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r NewPlanCumulativeGroupedBulkPriceParam) ImplementsBetaNewPlanVersionParamsAddPricesPriceUnion() {
}

func (r NewPlanCumulativeGroupedBulkPriceParam) ImplementsBetaNewPlanVersionParamsReplacePricesPriceUnion() {
}

func (r NewPlanCumulativeGroupedBulkPriceParam) ImplementsBetaExternalPlanIDNewPlanVersionParamsAddPricesPriceUnion() {
}

func (r NewPlanCumulativeGroupedBulkPriceParam) ImplementsBetaExternalPlanIDNewPlanVersionParamsReplacePricesPriceUnion() {
}

func (r NewPlanCumulativeGroupedBulkPriceParam) ImplementsPlanNewParamsPricesPriceUnion() {}

// The cadence to bill for this price on.
type NewPlanCumulativeGroupedBulkPriceCadence string

const (
	NewPlanCumulativeGroupedBulkPriceCadenceAnnual     NewPlanCumulativeGroupedBulkPriceCadence = "annual"
	NewPlanCumulativeGroupedBulkPriceCadenceSemiAnnual NewPlanCumulativeGroupedBulkPriceCadence = "semi_annual"
	NewPlanCumulativeGroupedBulkPriceCadenceMonthly    NewPlanCumulativeGroupedBulkPriceCadence = "monthly"
	NewPlanCumulativeGroupedBulkPriceCadenceQuarterly  NewPlanCumulativeGroupedBulkPriceCadence = "quarterly"
	NewPlanCumulativeGroupedBulkPriceCadenceOneTime    NewPlanCumulativeGroupedBulkPriceCadence = "one_time"
	NewPlanCumulativeGroupedBulkPriceCadenceCustom     NewPlanCumulativeGroupedBulkPriceCadence = "custom"
)

func (r NewPlanCumulativeGroupedBulkPriceCadence) IsKnown() bool {
	switch r {
	case NewPlanCumulativeGroupedBulkPriceCadenceAnnual, NewPlanCumulativeGroupedBulkPriceCadenceSemiAnnual, NewPlanCumulativeGroupedBulkPriceCadenceMonthly, NewPlanCumulativeGroupedBulkPriceCadenceQuarterly, NewPlanCumulativeGroupedBulkPriceCadenceOneTime, NewPlanCumulativeGroupedBulkPriceCadenceCustom:
		return true
	}
	return false
}

type NewPlanCumulativeGroupedBulkPriceModelType string

const (
	NewPlanCumulativeGroupedBulkPriceModelTypeCumulativeGroupedBulk NewPlanCumulativeGroupedBulkPriceModelType = "cumulative_grouped_bulk"
)

func (r NewPlanCumulativeGroupedBulkPriceModelType) IsKnown() bool {
	switch r {
	case NewPlanCumulativeGroupedBulkPriceModelTypeCumulativeGroupedBulk:
		return true
	}
	return false
}

type NewPlanCumulativeGroupedBulkPriceParam struct {
	ConversionRateType param.Field[NewPlanCumulativeGroupedBulkPriceConversionRateType] `json:"conversion_rate_type,required"`
	TieredConfig       param.Field[ConversionRateTieredConfigParam]                     `json:"tiered_config"`
	UnitConfig         param.Field[ConversionRateUnitConfigParam]                       `json:"unit_config"`
}

func (r NewPlanCumulativeGroupedBulkPriceParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r NewPlanCumulativeGroupedBulkPriceParam) ImplementsNewPlanCumulativeGroupedBulkPriceUnionParam() {
}

// Satisfied by [shared.UnitConversionRateConfigParam],
// [shared.TieredConversionRateConfigParam],
// [NewPlanCumulativeGroupedBulkPriceParam].
type NewPlanCumulativeGroupedBulkPriceUnionParam interface {
	ImplementsNewPlanCumulativeGroupedBulkPriceUnionParam()
}

type NewPlanCumulativeGroupedBulkPriceConversionRateType string

const (
	NewPlanCumulativeGroupedBulkPriceConversionRateTypeUnit   NewPlanCumulativeGroupedBulkPriceConversionRateType = "unit"
	NewPlanCumulativeGroupedBulkPriceConversionRateTypeTiered NewPlanCumulativeGroupedBulkPriceConversionRateType = "tiered"
)

func (r NewPlanCumulativeGroupedBulkPriceConversionRateType) IsKnown() bool {
	switch r {
	case NewPlanCumulativeGroupedBulkPriceConversionRateTypeUnit, NewPlanCumulativeGroupedBulkPriceConversionRateTypeTiered:
		return true
	}
	return false
}

type NewPlanGroupedAllocationPriceParam struct {
	// The cadence to bill for this price on.
	Cadence                 param.Field[NewPlanGroupedAllocationPriceCadence] `json:"cadence,required"`
	GroupedAllocationConfig param.Field[map[string]interface{}]               `json:"grouped_allocation_config,required"`
	// The id of the item the price will be associated with.
	ItemID    param.Field[string]                                 `json:"item_id,required"`
	ModelType param.Field[NewPlanGroupedAllocationPriceModelType] `json:"model_type,required"`
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
	BillingCycleConfiguration param.Field[NewBillingCycleConfigurationParam] `json:"billing_cycle_configuration"`
	// The per unit conversion rate of the price currency to the invoicing currency.
	ConversionRate param.Field[float64] `json:"conversion_rate"`
	// The configuration for the rate of the price currency to the invoicing currency.
	ConversionRateConfig param.Field[NewPlanGroupedAllocationPriceUnionParam] `json:"conversion_rate_config"`
	// An ISO 4217 currency string, or custom pricing unit identifier, in which this
	// price is billed.
	Currency param.Field[string] `json:"currency"`
	// For dimensional price: specifies a price group and dimension values
	DimensionalPriceConfiguration param.Field[NewDimensionalPriceConfigurationParam] `json:"dimensional_price_configuration"`
	// An alias for the price.
	ExternalPriceID param.Field[string] `json:"external_price_id"`
	// If the Price represents a fixed cost, this represents the quantity of units
	// applied.
	FixedPriceQuantity param.Field[float64] `json:"fixed_price_quantity"`
	// The property used to group this price on an invoice
	InvoiceGroupingKey param.Field[string] `json:"invoice_grouping_key"`
	// Within each billing cycle, specifies the cadence at which invoices are produced.
	// If unspecified, a single invoice is produced per billing cycle.
	InvoicingCycleConfiguration param.Field[NewBillingCycleConfigurationParam] `json:"invoicing_cycle_configuration"`
	// User-specified key/value pairs for the resource. Individual keys can be removed
	// by setting the value to `null`, and the entire metadata mapping can be cleared
	// by setting `metadata` to `null`.
	Metadata param.Field[map[string]string] `json:"metadata"`
	// A transient ID that can be used to reference this price when adding adjustments
	// in the same API call.
	ReferenceID param.Field[string] `json:"reference_id"`
}

func (r NewPlanGroupedAllocationPriceParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r NewPlanGroupedAllocationPriceParam) ImplementsBetaNewPlanVersionParamsAddPricesPriceUnion() {}

func (r NewPlanGroupedAllocationPriceParam) ImplementsBetaNewPlanVersionParamsReplacePricesPriceUnion() {
}

func (r NewPlanGroupedAllocationPriceParam) ImplementsBetaExternalPlanIDNewPlanVersionParamsAddPricesPriceUnion() {
}

func (r NewPlanGroupedAllocationPriceParam) ImplementsBetaExternalPlanIDNewPlanVersionParamsReplacePricesPriceUnion() {
}

func (r NewPlanGroupedAllocationPriceParam) ImplementsPlanNewParamsPricesPriceUnion() {}

// The cadence to bill for this price on.
type NewPlanGroupedAllocationPriceCadence string

const (
	NewPlanGroupedAllocationPriceCadenceAnnual     NewPlanGroupedAllocationPriceCadence = "annual"
	NewPlanGroupedAllocationPriceCadenceSemiAnnual NewPlanGroupedAllocationPriceCadence = "semi_annual"
	NewPlanGroupedAllocationPriceCadenceMonthly    NewPlanGroupedAllocationPriceCadence = "monthly"
	NewPlanGroupedAllocationPriceCadenceQuarterly  NewPlanGroupedAllocationPriceCadence = "quarterly"
	NewPlanGroupedAllocationPriceCadenceOneTime    NewPlanGroupedAllocationPriceCadence = "one_time"
	NewPlanGroupedAllocationPriceCadenceCustom     NewPlanGroupedAllocationPriceCadence = "custom"
)

func (r NewPlanGroupedAllocationPriceCadence) IsKnown() bool {
	switch r {
	case NewPlanGroupedAllocationPriceCadenceAnnual, NewPlanGroupedAllocationPriceCadenceSemiAnnual, NewPlanGroupedAllocationPriceCadenceMonthly, NewPlanGroupedAllocationPriceCadenceQuarterly, NewPlanGroupedAllocationPriceCadenceOneTime, NewPlanGroupedAllocationPriceCadenceCustom:
		return true
	}
	return false
}

type NewPlanGroupedAllocationPriceModelType string

const (
	NewPlanGroupedAllocationPriceModelTypeGroupedAllocation NewPlanGroupedAllocationPriceModelType = "grouped_allocation"
)

func (r NewPlanGroupedAllocationPriceModelType) IsKnown() bool {
	switch r {
	case NewPlanGroupedAllocationPriceModelTypeGroupedAllocation:
		return true
	}
	return false
}

type NewPlanGroupedAllocationPriceParam struct {
	ConversionRateType param.Field[NewPlanGroupedAllocationPriceConversionRateType] `json:"conversion_rate_type,required"`
	TieredConfig       param.Field[ConversionRateTieredConfigParam]                 `json:"tiered_config"`
	UnitConfig         param.Field[ConversionRateUnitConfigParam]                   `json:"unit_config"`
}

func (r NewPlanGroupedAllocationPriceParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r NewPlanGroupedAllocationPriceParam) ImplementsNewPlanGroupedAllocationPriceUnionParam() {}

// Satisfied by [shared.UnitConversionRateConfigParam],
// [shared.TieredConversionRateConfigParam], [NewPlanGroupedAllocationPriceParam].
type NewPlanGroupedAllocationPriceUnionParam interface {
	ImplementsNewPlanGroupedAllocationPriceUnionParam()
}

type NewPlanGroupedAllocationPriceConversionRateType string

const (
	NewPlanGroupedAllocationPriceConversionRateTypeUnit   NewPlanGroupedAllocationPriceConversionRateType = "unit"
	NewPlanGroupedAllocationPriceConversionRateTypeTiered NewPlanGroupedAllocationPriceConversionRateType = "tiered"
)

func (r NewPlanGroupedAllocationPriceConversionRateType) IsKnown() bool {
	switch r {
	case NewPlanGroupedAllocationPriceConversionRateTypeUnit, NewPlanGroupedAllocationPriceConversionRateTypeTiered:
		return true
	}
	return false
}

type NewPlanGroupedTieredPackagePriceParam struct {
	// The cadence to bill for this price on.
	Cadence                    param.Field[NewPlanGroupedTieredPackagePriceCadence] `json:"cadence,required"`
	GroupedTieredPackageConfig param.Field[map[string]interface{}]                  `json:"grouped_tiered_package_config,required"`
	// The id of the item the price will be associated with.
	ItemID    param.Field[string]                                    `json:"item_id,required"`
	ModelType param.Field[NewPlanGroupedTieredPackagePriceModelType] `json:"model_type,required"`
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
	BillingCycleConfiguration param.Field[NewBillingCycleConfigurationParam] `json:"billing_cycle_configuration"`
	// The per unit conversion rate of the price currency to the invoicing currency.
	ConversionRate param.Field[float64] `json:"conversion_rate"`
	// The configuration for the rate of the price currency to the invoicing currency.
	ConversionRateConfig param.Field[NewPlanGroupedTieredPackagePriceUnionParam] `json:"conversion_rate_config"`
	// An ISO 4217 currency string, or custom pricing unit identifier, in which this
	// price is billed.
	Currency param.Field[string] `json:"currency"`
	// For dimensional price: specifies a price group and dimension values
	DimensionalPriceConfiguration param.Field[NewDimensionalPriceConfigurationParam] `json:"dimensional_price_configuration"`
	// An alias for the price.
	ExternalPriceID param.Field[string] `json:"external_price_id"`
	// If the Price represents a fixed cost, this represents the quantity of units
	// applied.
	FixedPriceQuantity param.Field[float64] `json:"fixed_price_quantity"`
	// The property used to group this price on an invoice
	InvoiceGroupingKey param.Field[string] `json:"invoice_grouping_key"`
	// Within each billing cycle, specifies the cadence at which invoices are produced.
	// If unspecified, a single invoice is produced per billing cycle.
	InvoicingCycleConfiguration param.Field[NewBillingCycleConfigurationParam] `json:"invoicing_cycle_configuration"`
	// User-specified key/value pairs for the resource. Individual keys can be removed
	// by setting the value to `null`, and the entire metadata mapping can be cleared
	// by setting `metadata` to `null`.
	Metadata param.Field[map[string]string] `json:"metadata"`
	// A transient ID that can be used to reference this price when adding adjustments
	// in the same API call.
	ReferenceID param.Field[string] `json:"reference_id"`
}

func (r NewPlanGroupedTieredPackagePriceParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r NewPlanGroupedTieredPackagePriceParam) ImplementsBetaNewPlanVersionParamsAddPricesPriceUnion() {
}

func (r NewPlanGroupedTieredPackagePriceParam) ImplementsBetaNewPlanVersionParamsReplacePricesPriceUnion() {
}

func (r NewPlanGroupedTieredPackagePriceParam) ImplementsBetaExternalPlanIDNewPlanVersionParamsAddPricesPriceUnion() {
}

func (r NewPlanGroupedTieredPackagePriceParam) ImplementsBetaExternalPlanIDNewPlanVersionParamsReplacePricesPriceUnion() {
}

func (r NewPlanGroupedTieredPackagePriceParam) ImplementsPlanNewParamsPricesPriceUnion() {}

// The cadence to bill for this price on.
type NewPlanGroupedTieredPackagePriceCadence string

const (
	NewPlanGroupedTieredPackagePriceCadenceAnnual     NewPlanGroupedTieredPackagePriceCadence = "annual"
	NewPlanGroupedTieredPackagePriceCadenceSemiAnnual NewPlanGroupedTieredPackagePriceCadence = "semi_annual"
	NewPlanGroupedTieredPackagePriceCadenceMonthly    NewPlanGroupedTieredPackagePriceCadence = "monthly"
	NewPlanGroupedTieredPackagePriceCadenceQuarterly  NewPlanGroupedTieredPackagePriceCadence = "quarterly"
	NewPlanGroupedTieredPackagePriceCadenceOneTime    NewPlanGroupedTieredPackagePriceCadence = "one_time"
	NewPlanGroupedTieredPackagePriceCadenceCustom     NewPlanGroupedTieredPackagePriceCadence = "custom"
)

func (r NewPlanGroupedTieredPackagePriceCadence) IsKnown() bool {
	switch r {
	case NewPlanGroupedTieredPackagePriceCadenceAnnual, NewPlanGroupedTieredPackagePriceCadenceSemiAnnual, NewPlanGroupedTieredPackagePriceCadenceMonthly, NewPlanGroupedTieredPackagePriceCadenceQuarterly, NewPlanGroupedTieredPackagePriceCadenceOneTime, NewPlanGroupedTieredPackagePriceCadenceCustom:
		return true
	}
	return false
}

type NewPlanGroupedTieredPackagePriceModelType string

const (
	NewPlanGroupedTieredPackagePriceModelTypeGroupedTieredPackage NewPlanGroupedTieredPackagePriceModelType = "grouped_tiered_package"
)

func (r NewPlanGroupedTieredPackagePriceModelType) IsKnown() bool {
	switch r {
	case NewPlanGroupedTieredPackagePriceModelTypeGroupedTieredPackage:
		return true
	}
	return false
}

type NewPlanGroupedTieredPackagePriceParam struct {
	ConversionRateType param.Field[NewPlanGroupedTieredPackagePriceConversionRateType] `json:"conversion_rate_type,required"`
	TieredConfig       param.Field[ConversionRateTieredConfigParam]                    `json:"tiered_config"`
	UnitConfig         param.Field[ConversionRateUnitConfigParam]                      `json:"unit_config"`
}

func (r NewPlanGroupedTieredPackagePriceParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r NewPlanGroupedTieredPackagePriceParam) ImplementsNewPlanGroupedTieredPackagePriceUnionParam() {
}

// Satisfied by [shared.UnitConversionRateConfigParam],
// [shared.TieredConversionRateConfigParam],
// [NewPlanGroupedTieredPackagePriceParam].
type NewPlanGroupedTieredPackagePriceUnionParam interface {
	ImplementsNewPlanGroupedTieredPackagePriceUnionParam()
}

type NewPlanGroupedTieredPackagePriceConversionRateType string

const (
	NewPlanGroupedTieredPackagePriceConversionRateTypeUnit   NewPlanGroupedTieredPackagePriceConversionRateType = "unit"
	NewPlanGroupedTieredPackagePriceConversionRateTypeTiered NewPlanGroupedTieredPackagePriceConversionRateType = "tiered"
)

func (r NewPlanGroupedTieredPackagePriceConversionRateType) IsKnown() bool {
	switch r {
	case NewPlanGroupedTieredPackagePriceConversionRateTypeUnit, NewPlanGroupedTieredPackagePriceConversionRateTypeTiered:
		return true
	}
	return false
}

type NewPlanGroupedTieredPriceParam struct {
	// The cadence to bill for this price on.
	Cadence             param.Field[NewPlanGroupedTieredPriceCadence] `json:"cadence,required"`
	GroupedTieredConfig param.Field[map[string]interface{}]           `json:"grouped_tiered_config,required"`
	// The id of the item the price will be associated with.
	ItemID    param.Field[string]                             `json:"item_id,required"`
	ModelType param.Field[NewPlanGroupedTieredPriceModelType] `json:"model_type,required"`
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
	BillingCycleConfiguration param.Field[NewBillingCycleConfigurationParam] `json:"billing_cycle_configuration"`
	// The per unit conversion rate of the price currency to the invoicing currency.
	ConversionRate param.Field[float64] `json:"conversion_rate"`
	// The configuration for the rate of the price currency to the invoicing currency.
	ConversionRateConfig param.Field[NewPlanGroupedTieredPriceUnionParam] `json:"conversion_rate_config"`
	// An ISO 4217 currency string, or custom pricing unit identifier, in which this
	// price is billed.
	Currency param.Field[string] `json:"currency"`
	// For dimensional price: specifies a price group and dimension values
	DimensionalPriceConfiguration param.Field[NewDimensionalPriceConfigurationParam] `json:"dimensional_price_configuration"`
	// An alias for the price.
	ExternalPriceID param.Field[string] `json:"external_price_id"`
	// If the Price represents a fixed cost, this represents the quantity of units
	// applied.
	FixedPriceQuantity param.Field[float64] `json:"fixed_price_quantity"`
	// The property used to group this price on an invoice
	InvoiceGroupingKey param.Field[string] `json:"invoice_grouping_key"`
	// Within each billing cycle, specifies the cadence at which invoices are produced.
	// If unspecified, a single invoice is produced per billing cycle.
	InvoicingCycleConfiguration param.Field[NewBillingCycleConfigurationParam] `json:"invoicing_cycle_configuration"`
	// User-specified key/value pairs for the resource. Individual keys can be removed
	// by setting the value to `null`, and the entire metadata mapping can be cleared
	// by setting `metadata` to `null`.
	Metadata param.Field[map[string]string] `json:"metadata"`
	// A transient ID that can be used to reference this price when adding adjustments
	// in the same API call.
	ReferenceID param.Field[string] `json:"reference_id"`
}

func (r NewPlanGroupedTieredPriceParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r NewPlanGroupedTieredPriceParam) ImplementsBetaNewPlanVersionParamsAddPricesPriceUnion() {}

func (r NewPlanGroupedTieredPriceParam) ImplementsBetaNewPlanVersionParamsReplacePricesPriceUnion() {}

func (r NewPlanGroupedTieredPriceParam) ImplementsBetaExternalPlanIDNewPlanVersionParamsAddPricesPriceUnion() {
}

func (r NewPlanGroupedTieredPriceParam) ImplementsBetaExternalPlanIDNewPlanVersionParamsReplacePricesPriceUnion() {
}

func (r NewPlanGroupedTieredPriceParam) ImplementsPlanNewParamsPricesPriceUnion() {}

// The cadence to bill for this price on.
type NewPlanGroupedTieredPriceCadence string

const (
	NewPlanGroupedTieredPriceCadenceAnnual     NewPlanGroupedTieredPriceCadence = "annual"
	NewPlanGroupedTieredPriceCadenceSemiAnnual NewPlanGroupedTieredPriceCadence = "semi_annual"
	NewPlanGroupedTieredPriceCadenceMonthly    NewPlanGroupedTieredPriceCadence = "monthly"
	NewPlanGroupedTieredPriceCadenceQuarterly  NewPlanGroupedTieredPriceCadence = "quarterly"
	NewPlanGroupedTieredPriceCadenceOneTime    NewPlanGroupedTieredPriceCadence = "one_time"
	NewPlanGroupedTieredPriceCadenceCustom     NewPlanGroupedTieredPriceCadence = "custom"
)

func (r NewPlanGroupedTieredPriceCadence) IsKnown() bool {
	switch r {
	case NewPlanGroupedTieredPriceCadenceAnnual, NewPlanGroupedTieredPriceCadenceSemiAnnual, NewPlanGroupedTieredPriceCadenceMonthly, NewPlanGroupedTieredPriceCadenceQuarterly, NewPlanGroupedTieredPriceCadenceOneTime, NewPlanGroupedTieredPriceCadenceCustom:
		return true
	}
	return false
}

type NewPlanGroupedTieredPriceModelType string

const (
	NewPlanGroupedTieredPriceModelTypeGroupedTiered NewPlanGroupedTieredPriceModelType = "grouped_tiered"
)

func (r NewPlanGroupedTieredPriceModelType) IsKnown() bool {
	switch r {
	case NewPlanGroupedTieredPriceModelTypeGroupedTiered:
		return true
	}
	return false
}

type NewPlanGroupedTieredPriceParam struct {
	ConversionRateType param.Field[NewPlanGroupedTieredPriceConversionRateType] `json:"conversion_rate_type,required"`
	TieredConfig       param.Field[ConversionRateTieredConfigParam]             `json:"tiered_config"`
	UnitConfig         param.Field[ConversionRateUnitConfigParam]               `json:"unit_config"`
}

func (r NewPlanGroupedTieredPriceParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r NewPlanGroupedTieredPriceParam) ImplementsNewPlanGroupedTieredPriceUnionParam() {}

// Satisfied by [shared.UnitConversionRateConfigParam],
// [shared.TieredConversionRateConfigParam], [NewPlanGroupedTieredPriceParam].
type NewPlanGroupedTieredPriceUnionParam interface {
	ImplementsNewPlanGroupedTieredPriceUnionParam()
}

type NewPlanGroupedTieredPriceConversionRateType string

const (
	NewPlanGroupedTieredPriceConversionRateTypeUnit   NewPlanGroupedTieredPriceConversionRateType = "unit"
	NewPlanGroupedTieredPriceConversionRateTypeTiered NewPlanGroupedTieredPriceConversionRateType = "tiered"
)

func (r NewPlanGroupedTieredPriceConversionRateType) IsKnown() bool {
	switch r {
	case NewPlanGroupedTieredPriceConversionRateTypeUnit, NewPlanGroupedTieredPriceConversionRateTypeTiered:
		return true
	}
	return false
}

type NewPlanGroupedWithMeteredMinimumPriceParam struct {
	// The cadence to bill for this price on.
	Cadence                         param.Field[NewPlanGroupedWithMeteredMinimumPriceCadence] `json:"cadence,required"`
	GroupedWithMeteredMinimumConfig param.Field[map[string]interface{}]                       `json:"grouped_with_metered_minimum_config,required"`
	// The id of the item the price will be associated with.
	ItemID    param.Field[string]                                         `json:"item_id,required"`
	ModelType param.Field[NewPlanGroupedWithMeteredMinimumPriceModelType] `json:"model_type,required"`
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
	BillingCycleConfiguration param.Field[NewBillingCycleConfigurationParam] `json:"billing_cycle_configuration"`
	// The per unit conversion rate of the price currency to the invoicing currency.
	ConversionRate param.Field[float64] `json:"conversion_rate"`
	// The configuration for the rate of the price currency to the invoicing currency.
	ConversionRateConfig param.Field[NewPlanGroupedWithMeteredMinimumPriceUnionParam] `json:"conversion_rate_config"`
	// An ISO 4217 currency string, or custom pricing unit identifier, in which this
	// price is billed.
	Currency param.Field[string] `json:"currency"`
	// For dimensional price: specifies a price group and dimension values
	DimensionalPriceConfiguration param.Field[NewDimensionalPriceConfigurationParam] `json:"dimensional_price_configuration"`
	// An alias for the price.
	ExternalPriceID param.Field[string] `json:"external_price_id"`
	// If the Price represents a fixed cost, this represents the quantity of units
	// applied.
	FixedPriceQuantity param.Field[float64] `json:"fixed_price_quantity"`
	// The property used to group this price on an invoice
	InvoiceGroupingKey param.Field[string] `json:"invoice_grouping_key"`
	// Within each billing cycle, specifies the cadence at which invoices are produced.
	// If unspecified, a single invoice is produced per billing cycle.
	InvoicingCycleConfiguration param.Field[NewBillingCycleConfigurationParam] `json:"invoicing_cycle_configuration"`
	// User-specified key/value pairs for the resource. Individual keys can be removed
	// by setting the value to `null`, and the entire metadata mapping can be cleared
	// by setting `metadata` to `null`.
	Metadata param.Field[map[string]string] `json:"metadata"`
	// A transient ID that can be used to reference this price when adding adjustments
	// in the same API call.
	ReferenceID param.Field[string] `json:"reference_id"`
}

func (r NewPlanGroupedWithMeteredMinimumPriceParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r NewPlanGroupedWithMeteredMinimumPriceParam) ImplementsBetaNewPlanVersionParamsAddPricesPriceUnion() {
}

func (r NewPlanGroupedWithMeteredMinimumPriceParam) ImplementsBetaNewPlanVersionParamsReplacePricesPriceUnion() {
}

func (r NewPlanGroupedWithMeteredMinimumPriceParam) ImplementsBetaExternalPlanIDNewPlanVersionParamsAddPricesPriceUnion() {
}

func (r NewPlanGroupedWithMeteredMinimumPriceParam) ImplementsBetaExternalPlanIDNewPlanVersionParamsReplacePricesPriceUnion() {
}

func (r NewPlanGroupedWithMeteredMinimumPriceParam) ImplementsPlanNewParamsPricesPriceUnion() {}

// The cadence to bill for this price on.
type NewPlanGroupedWithMeteredMinimumPriceCadence string

const (
	NewPlanGroupedWithMeteredMinimumPriceCadenceAnnual     NewPlanGroupedWithMeteredMinimumPriceCadence = "annual"
	NewPlanGroupedWithMeteredMinimumPriceCadenceSemiAnnual NewPlanGroupedWithMeteredMinimumPriceCadence = "semi_annual"
	NewPlanGroupedWithMeteredMinimumPriceCadenceMonthly    NewPlanGroupedWithMeteredMinimumPriceCadence = "monthly"
	NewPlanGroupedWithMeteredMinimumPriceCadenceQuarterly  NewPlanGroupedWithMeteredMinimumPriceCadence = "quarterly"
	NewPlanGroupedWithMeteredMinimumPriceCadenceOneTime    NewPlanGroupedWithMeteredMinimumPriceCadence = "one_time"
	NewPlanGroupedWithMeteredMinimumPriceCadenceCustom     NewPlanGroupedWithMeteredMinimumPriceCadence = "custom"
)

func (r NewPlanGroupedWithMeteredMinimumPriceCadence) IsKnown() bool {
	switch r {
	case NewPlanGroupedWithMeteredMinimumPriceCadenceAnnual, NewPlanGroupedWithMeteredMinimumPriceCadenceSemiAnnual, NewPlanGroupedWithMeteredMinimumPriceCadenceMonthly, NewPlanGroupedWithMeteredMinimumPriceCadenceQuarterly, NewPlanGroupedWithMeteredMinimumPriceCadenceOneTime, NewPlanGroupedWithMeteredMinimumPriceCadenceCustom:
		return true
	}
	return false
}

type NewPlanGroupedWithMeteredMinimumPriceModelType string

const (
	NewPlanGroupedWithMeteredMinimumPriceModelTypeGroupedWithMeteredMinimum NewPlanGroupedWithMeteredMinimumPriceModelType = "grouped_with_metered_minimum"
)

func (r NewPlanGroupedWithMeteredMinimumPriceModelType) IsKnown() bool {
	switch r {
	case NewPlanGroupedWithMeteredMinimumPriceModelTypeGroupedWithMeteredMinimum:
		return true
	}
	return false
}

type NewPlanGroupedWithMeteredMinimumPriceParam struct {
	ConversionRateType param.Field[NewPlanGroupedWithMeteredMinimumPriceConversionRateType] `json:"conversion_rate_type,required"`
	TieredConfig       param.Field[ConversionRateTieredConfigParam]                         `json:"tiered_config"`
	UnitConfig         param.Field[ConversionRateUnitConfigParam]                           `json:"unit_config"`
}

func (r NewPlanGroupedWithMeteredMinimumPriceParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r NewPlanGroupedWithMeteredMinimumPriceParam) ImplementsNewPlanGroupedWithMeteredMinimumPriceUnionParam() {
}

// Satisfied by [shared.UnitConversionRateConfigParam],
// [shared.TieredConversionRateConfigParam],
// [NewPlanGroupedWithMeteredMinimumPriceParam].
type NewPlanGroupedWithMeteredMinimumPriceUnionParam interface {
	ImplementsNewPlanGroupedWithMeteredMinimumPriceUnionParam()
}

type NewPlanGroupedWithMeteredMinimumPriceConversionRateType string

const (
	NewPlanGroupedWithMeteredMinimumPriceConversionRateTypeUnit   NewPlanGroupedWithMeteredMinimumPriceConversionRateType = "unit"
	NewPlanGroupedWithMeteredMinimumPriceConversionRateTypeTiered NewPlanGroupedWithMeteredMinimumPriceConversionRateType = "tiered"
)

func (r NewPlanGroupedWithMeteredMinimumPriceConversionRateType) IsKnown() bool {
	switch r {
	case NewPlanGroupedWithMeteredMinimumPriceConversionRateTypeUnit, NewPlanGroupedWithMeteredMinimumPriceConversionRateTypeTiered:
		return true
	}
	return false
}

type NewPlanGroupedWithProratedMinimumPriceParam struct {
	// The cadence to bill for this price on.
	Cadence                          param.Field[NewPlanGroupedWithProratedMinimumPriceCadence] `json:"cadence,required"`
	GroupedWithProratedMinimumConfig param.Field[map[string]interface{}]                        `json:"grouped_with_prorated_minimum_config,required"`
	// The id of the item the price will be associated with.
	ItemID    param.Field[string]                                          `json:"item_id,required"`
	ModelType param.Field[NewPlanGroupedWithProratedMinimumPriceModelType] `json:"model_type,required"`
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
	BillingCycleConfiguration param.Field[NewBillingCycleConfigurationParam] `json:"billing_cycle_configuration"`
	// The per unit conversion rate of the price currency to the invoicing currency.
	ConversionRate param.Field[float64] `json:"conversion_rate"`
	// The configuration for the rate of the price currency to the invoicing currency.
	ConversionRateConfig param.Field[NewPlanGroupedWithProratedMinimumPriceUnionParam] `json:"conversion_rate_config"`
	// An ISO 4217 currency string, or custom pricing unit identifier, in which this
	// price is billed.
	Currency param.Field[string] `json:"currency"`
	// For dimensional price: specifies a price group and dimension values
	DimensionalPriceConfiguration param.Field[NewDimensionalPriceConfigurationParam] `json:"dimensional_price_configuration"`
	// An alias for the price.
	ExternalPriceID param.Field[string] `json:"external_price_id"`
	// If the Price represents a fixed cost, this represents the quantity of units
	// applied.
	FixedPriceQuantity param.Field[float64] `json:"fixed_price_quantity"`
	// The property used to group this price on an invoice
	InvoiceGroupingKey param.Field[string] `json:"invoice_grouping_key"`
	// Within each billing cycle, specifies the cadence at which invoices are produced.
	// If unspecified, a single invoice is produced per billing cycle.
	InvoicingCycleConfiguration param.Field[NewBillingCycleConfigurationParam] `json:"invoicing_cycle_configuration"`
	// User-specified key/value pairs for the resource. Individual keys can be removed
	// by setting the value to `null`, and the entire metadata mapping can be cleared
	// by setting `metadata` to `null`.
	Metadata param.Field[map[string]string] `json:"metadata"`
	// A transient ID that can be used to reference this price when adding adjustments
	// in the same API call.
	ReferenceID param.Field[string] `json:"reference_id"`
}

func (r NewPlanGroupedWithProratedMinimumPriceParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r NewPlanGroupedWithProratedMinimumPriceParam) ImplementsBetaNewPlanVersionParamsAddPricesPriceUnion() {
}

func (r NewPlanGroupedWithProratedMinimumPriceParam) ImplementsBetaNewPlanVersionParamsReplacePricesPriceUnion() {
}

func (r NewPlanGroupedWithProratedMinimumPriceParam) ImplementsBetaExternalPlanIDNewPlanVersionParamsAddPricesPriceUnion() {
}

func (r NewPlanGroupedWithProratedMinimumPriceParam) ImplementsBetaExternalPlanIDNewPlanVersionParamsReplacePricesPriceUnion() {
}

func (r NewPlanGroupedWithProratedMinimumPriceParam) ImplementsPlanNewParamsPricesPriceUnion() {}

// The cadence to bill for this price on.
type NewPlanGroupedWithProratedMinimumPriceCadence string

const (
	NewPlanGroupedWithProratedMinimumPriceCadenceAnnual     NewPlanGroupedWithProratedMinimumPriceCadence = "annual"
	NewPlanGroupedWithProratedMinimumPriceCadenceSemiAnnual NewPlanGroupedWithProratedMinimumPriceCadence = "semi_annual"
	NewPlanGroupedWithProratedMinimumPriceCadenceMonthly    NewPlanGroupedWithProratedMinimumPriceCadence = "monthly"
	NewPlanGroupedWithProratedMinimumPriceCadenceQuarterly  NewPlanGroupedWithProratedMinimumPriceCadence = "quarterly"
	NewPlanGroupedWithProratedMinimumPriceCadenceOneTime    NewPlanGroupedWithProratedMinimumPriceCadence = "one_time"
	NewPlanGroupedWithProratedMinimumPriceCadenceCustom     NewPlanGroupedWithProratedMinimumPriceCadence = "custom"
)

func (r NewPlanGroupedWithProratedMinimumPriceCadence) IsKnown() bool {
	switch r {
	case NewPlanGroupedWithProratedMinimumPriceCadenceAnnual, NewPlanGroupedWithProratedMinimumPriceCadenceSemiAnnual, NewPlanGroupedWithProratedMinimumPriceCadenceMonthly, NewPlanGroupedWithProratedMinimumPriceCadenceQuarterly, NewPlanGroupedWithProratedMinimumPriceCadenceOneTime, NewPlanGroupedWithProratedMinimumPriceCadenceCustom:
		return true
	}
	return false
}

type NewPlanGroupedWithProratedMinimumPriceModelType string

const (
	NewPlanGroupedWithProratedMinimumPriceModelTypeGroupedWithProratedMinimum NewPlanGroupedWithProratedMinimumPriceModelType = "grouped_with_prorated_minimum"
)

func (r NewPlanGroupedWithProratedMinimumPriceModelType) IsKnown() bool {
	switch r {
	case NewPlanGroupedWithProratedMinimumPriceModelTypeGroupedWithProratedMinimum:
		return true
	}
	return false
}

type NewPlanGroupedWithProratedMinimumPriceParam struct {
	ConversionRateType param.Field[NewPlanGroupedWithProratedMinimumPriceConversionRateType] `json:"conversion_rate_type,required"`
	TieredConfig       param.Field[ConversionRateTieredConfigParam]                          `json:"tiered_config"`
	UnitConfig         param.Field[ConversionRateUnitConfigParam]                            `json:"unit_config"`
}

func (r NewPlanGroupedWithProratedMinimumPriceParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r NewPlanGroupedWithProratedMinimumPriceParam) ImplementsNewPlanGroupedWithProratedMinimumPriceUnionParam() {
}

// Satisfied by [shared.UnitConversionRateConfigParam],
// [shared.TieredConversionRateConfigParam],
// [NewPlanGroupedWithProratedMinimumPriceParam].
type NewPlanGroupedWithProratedMinimumPriceUnionParam interface {
	ImplementsNewPlanGroupedWithProratedMinimumPriceUnionParam()
}

type NewPlanGroupedWithProratedMinimumPriceConversionRateType string

const (
	NewPlanGroupedWithProratedMinimumPriceConversionRateTypeUnit   NewPlanGroupedWithProratedMinimumPriceConversionRateType = "unit"
	NewPlanGroupedWithProratedMinimumPriceConversionRateTypeTiered NewPlanGroupedWithProratedMinimumPriceConversionRateType = "tiered"
)

func (r NewPlanGroupedWithProratedMinimumPriceConversionRateType) IsKnown() bool {
	switch r {
	case NewPlanGroupedWithProratedMinimumPriceConversionRateTypeUnit, NewPlanGroupedWithProratedMinimumPriceConversionRateTypeTiered:
		return true
	}
	return false
}

type NewPlanMatrixPriceParam struct {
	// The cadence to bill for this price on.
	Cadence param.Field[NewPlanMatrixPriceCadence] `json:"cadence,required"`
	// The id of the item the price will be associated with.
	ItemID       param.Field[string]                      `json:"item_id,required"`
	MatrixConfig param.Field[MatrixConfigParam]           `json:"matrix_config,required"`
	ModelType    param.Field[NewPlanMatrixPriceModelType] `json:"model_type,required"`
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
	BillingCycleConfiguration param.Field[NewBillingCycleConfigurationParam] `json:"billing_cycle_configuration"`
	// The per unit conversion rate of the price currency to the invoicing currency.
	ConversionRate param.Field[float64] `json:"conversion_rate"`
	// The configuration for the rate of the price currency to the invoicing currency.
	ConversionRateConfig param.Field[NewPlanMatrixPriceUnionParam] `json:"conversion_rate_config"`
	// An ISO 4217 currency string, or custom pricing unit identifier, in which this
	// price is billed.
	Currency param.Field[string] `json:"currency"`
	// For dimensional price: specifies a price group and dimension values
	DimensionalPriceConfiguration param.Field[NewDimensionalPriceConfigurationParam] `json:"dimensional_price_configuration"`
	// An alias for the price.
	ExternalPriceID param.Field[string] `json:"external_price_id"`
	// If the Price represents a fixed cost, this represents the quantity of units
	// applied.
	FixedPriceQuantity param.Field[float64] `json:"fixed_price_quantity"`
	// The property used to group this price on an invoice
	InvoiceGroupingKey param.Field[string] `json:"invoice_grouping_key"`
	// Within each billing cycle, specifies the cadence at which invoices are produced.
	// If unspecified, a single invoice is produced per billing cycle.
	InvoicingCycleConfiguration param.Field[NewBillingCycleConfigurationParam] `json:"invoicing_cycle_configuration"`
	// User-specified key/value pairs for the resource. Individual keys can be removed
	// by setting the value to `null`, and the entire metadata mapping can be cleared
	// by setting `metadata` to `null`.
	Metadata param.Field[map[string]string] `json:"metadata"`
	// A transient ID that can be used to reference this price when adding adjustments
	// in the same API call.
	ReferenceID param.Field[string] `json:"reference_id"`
}

func (r NewPlanMatrixPriceParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r NewPlanMatrixPriceParam) ImplementsBetaNewPlanVersionParamsAddPricesPriceUnion() {}

func (r NewPlanMatrixPriceParam) ImplementsBetaNewPlanVersionParamsReplacePricesPriceUnion() {}

func (r NewPlanMatrixPriceParam) ImplementsBetaExternalPlanIDNewPlanVersionParamsAddPricesPriceUnion() {
}

func (r NewPlanMatrixPriceParam) ImplementsBetaExternalPlanIDNewPlanVersionParamsReplacePricesPriceUnion() {
}

func (r NewPlanMatrixPriceParam) ImplementsPlanNewParamsPricesPriceUnion() {}

// The cadence to bill for this price on.
type NewPlanMatrixPriceCadence string

const (
	NewPlanMatrixPriceCadenceAnnual     NewPlanMatrixPriceCadence = "annual"
	NewPlanMatrixPriceCadenceSemiAnnual NewPlanMatrixPriceCadence = "semi_annual"
	NewPlanMatrixPriceCadenceMonthly    NewPlanMatrixPriceCadence = "monthly"
	NewPlanMatrixPriceCadenceQuarterly  NewPlanMatrixPriceCadence = "quarterly"
	NewPlanMatrixPriceCadenceOneTime    NewPlanMatrixPriceCadence = "one_time"
	NewPlanMatrixPriceCadenceCustom     NewPlanMatrixPriceCadence = "custom"
)

func (r NewPlanMatrixPriceCadence) IsKnown() bool {
	switch r {
	case NewPlanMatrixPriceCadenceAnnual, NewPlanMatrixPriceCadenceSemiAnnual, NewPlanMatrixPriceCadenceMonthly, NewPlanMatrixPriceCadenceQuarterly, NewPlanMatrixPriceCadenceOneTime, NewPlanMatrixPriceCadenceCustom:
		return true
	}
	return false
}

type NewPlanMatrixPriceModelType string

const (
	NewPlanMatrixPriceModelTypeMatrix NewPlanMatrixPriceModelType = "matrix"
)

func (r NewPlanMatrixPriceModelType) IsKnown() bool {
	switch r {
	case NewPlanMatrixPriceModelTypeMatrix:
		return true
	}
	return false
}

type NewPlanMatrixPriceParam struct {
	ConversionRateType param.Field[NewPlanMatrixPriceConversionRateType] `json:"conversion_rate_type,required"`
	TieredConfig       param.Field[ConversionRateTieredConfigParam]      `json:"tiered_config"`
	UnitConfig         param.Field[ConversionRateUnitConfigParam]        `json:"unit_config"`
}

func (r NewPlanMatrixPriceParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r NewPlanMatrixPriceParam) ImplementsNewPlanMatrixPriceUnionParam() {}

// Satisfied by [shared.UnitConversionRateConfigParam],
// [shared.TieredConversionRateConfigParam], [NewPlanMatrixPriceParam].
type NewPlanMatrixPriceUnionParam interface {
	ImplementsNewPlanMatrixPriceUnionParam()
}

type NewPlanMatrixPriceConversionRateType string

const (
	NewPlanMatrixPriceConversionRateTypeUnit   NewPlanMatrixPriceConversionRateType = "unit"
	NewPlanMatrixPriceConversionRateTypeTiered NewPlanMatrixPriceConversionRateType = "tiered"
)

func (r NewPlanMatrixPriceConversionRateType) IsKnown() bool {
	switch r {
	case NewPlanMatrixPriceConversionRateTypeUnit, NewPlanMatrixPriceConversionRateTypeTiered:
		return true
	}
	return false
}

type NewPlanMatrixWithAllocationPriceParam struct {
	// The cadence to bill for this price on.
	Cadence param.Field[NewPlanMatrixWithAllocationPriceCadence] `json:"cadence,required"`
	// The id of the item the price will be associated with.
	ItemID                     param.Field[string]                                    `json:"item_id,required"`
	MatrixWithAllocationConfig param.Field[MatrixWithAllocationConfigParam]           `json:"matrix_with_allocation_config,required"`
	ModelType                  param.Field[NewPlanMatrixWithAllocationPriceModelType] `json:"model_type,required"`
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
	BillingCycleConfiguration param.Field[NewBillingCycleConfigurationParam] `json:"billing_cycle_configuration"`
	// The per unit conversion rate of the price currency to the invoicing currency.
	ConversionRate param.Field[float64] `json:"conversion_rate"`
	// The configuration for the rate of the price currency to the invoicing currency.
	ConversionRateConfig param.Field[NewPlanMatrixWithAllocationPriceUnionParam] `json:"conversion_rate_config"`
	// An ISO 4217 currency string, or custom pricing unit identifier, in which this
	// price is billed.
	Currency param.Field[string] `json:"currency"`
	// For dimensional price: specifies a price group and dimension values
	DimensionalPriceConfiguration param.Field[NewDimensionalPriceConfigurationParam] `json:"dimensional_price_configuration"`
	// An alias for the price.
	ExternalPriceID param.Field[string] `json:"external_price_id"`
	// If the Price represents a fixed cost, this represents the quantity of units
	// applied.
	FixedPriceQuantity param.Field[float64] `json:"fixed_price_quantity"`
	// The property used to group this price on an invoice
	InvoiceGroupingKey param.Field[string] `json:"invoice_grouping_key"`
	// Within each billing cycle, specifies the cadence at which invoices are produced.
	// If unspecified, a single invoice is produced per billing cycle.
	InvoicingCycleConfiguration param.Field[NewBillingCycleConfigurationParam] `json:"invoicing_cycle_configuration"`
	// User-specified key/value pairs for the resource. Individual keys can be removed
	// by setting the value to `null`, and the entire metadata mapping can be cleared
	// by setting `metadata` to `null`.
	Metadata param.Field[map[string]string] `json:"metadata"`
	// A transient ID that can be used to reference this price when adding adjustments
	// in the same API call.
	ReferenceID param.Field[string] `json:"reference_id"`
}

func (r NewPlanMatrixWithAllocationPriceParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r NewPlanMatrixWithAllocationPriceParam) ImplementsBetaNewPlanVersionParamsAddPricesPriceUnion() {
}

func (r NewPlanMatrixWithAllocationPriceParam) ImplementsBetaNewPlanVersionParamsReplacePricesPriceUnion() {
}

func (r NewPlanMatrixWithAllocationPriceParam) ImplementsBetaExternalPlanIDNewPlanVersionParamsAddPricesPriceUnion() {
}

func (r NewPlanMatrixWithAllocationPriceParam) ImplementsBetaExternalPlanIDNewPlanVersionParamsReplacePricesPriceUnion() {
}

func (r NewPlanMatrixWithAllocationPriceParam) ImplementsPlanNewParamsPricesPriceUnion() {}

// The cadence to bill for this price on.
type NewPlanMatrixWithAllocationPriceCadence string

const (
	NewPlanMatrixWithAllocationPriceCadenceAnnual     NewPlanMatrixWithAllocationPriceCadence = "annual"
	NewPlanMatrixWithAllocationPriceCadenceSemiAnnual NewPlanMatrixWithAllocationPriceCadence = "semi_annual"
	NewPlanMatrixWithAllocationPriceCadenceMonthly    NewPlanMatrixWithAllocationPriceCadence = "monthly"
	NewPlanMatrixWithAllocationPriceCadenceQuarterly  NewPlanMatrixWithAllocationPriceCadence = "quarterly"
	NewPlanMatrixWithAllocationPriceCadenceOneTime    NewPlanMatrixWithAllocationPriceCadence = "one_time"
	NewPlanMatrixWithAllocationPriceCadenceCustom     NewPlanMatrixWithAllocationPriceCadence = "custom"
)

func (r NewPlanMatrixWithAllocationPriceCadence) IsKnown() bool {
	switch r {
	case NewPlanMatrixWithAllocationPriceCadenceAnnual, NewPlanMatrixWithAllocationPriceCadenceSemiAnnual, NewPlanMatrixWithAllocationPriceCadenceMonthly, NewPlanMatrixWithAllocationPriceCadenceQuarterly, NewPlanMatrixWithAllocationPriceCadenceOneTime, NewPlanMatrixWithAllocationPriceCadenceCustom:
		return true
	}
	return false
}

type NewPlanMatrixWithAllocationPriceModelType string

const (
	NewPlanMatrixWithAllocationPriceModelTypeMatrixWithAllocation NewPlanMatrixWithAllocationPriceModelType = "matrix_with_allocation"
)

func (r NewPlanMatrixWithAllocationPriceModelType) IsKnown() bool {
	switch r {
	case NewPlanMatrixWithAllocationPriceModelTypeMatrixWithAllocation:
		return true
	}
	return false
}

type NewPlanMatrixWithAllocationPriceParam struct {
	ConversionRateType param.Field[NewPlanMatrixWithAllocationPriceConversionRateType] `json:"conversion_rate_type,required"`
	TieredConfig       param.Field[ConversionRateTieredConfigParam]                    `json:"tiered_config"`
	UnitConfig         param.Field[ConversionRateUnitConfigParam]                      `json:"unit_config"`
}

func (r NewPlanMatrixWithAllocationPriceParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r NewPlanMatrixWithAllocationPriceParam) ImplementsNewPlanMatrixWithAllocationPriceUnionParam() {
}

// Satisfied by [shared.UnitConversionRateConfigParam],
// [shared.TieredConversionRateConfigParam],
// [NewPlanMatrixWithAllocationPriceParam].
type NewPlanMatrixWithAllocationPriceUnionParam interface {
	ImplementsNewPlanMatrixWithAllocationPriceUnionParam()
}

type NewPlanMatrixWithAllocationPriceConversionRateType string

const (
	NewPlanMatrixWithAllocationPriceConversionRateTypeUnit   NewPlanMatrixWithAllocationPriceConversionRateType = "unit"
	NewPlanMatrixWithAllocationPriceConversionRateTypeTiered NewPlanMatrixWithAllocationPriceConversionRateType = "tiered"
)

func (r NewPlanMatrixWithAllocationPriceConversionRateType) IsKnown() bool {
	switch r {
	case NewPlanMatrixWithAllocationPriceConversionRateTypeUnit, NewPlanMatrixWithAllocationPriceConversionRateTypeTiered:
		return true
	}
	return false
}

type NewPlanMatrixWithDisplayNamePriceParam struct {
	// The cadence to bill for this price on.
	Cadence param.Field[NewPlanMatrixWithDisplayNamePriceCadence] `json:"cadence,required"`
	// The id of the item the price will be associated with.
	ItemID                      param.Field[string]                                     `json:"item_id,required"`
	MatrixWithDisplayNameConfig param.Field[map[string]interface{}]                     `json:"matrix_with_display_name_config,required"`
	ModelType                   param.Field[NewPlanMatrixWithDisplayNamePriceModelType] `json:"model_type,required"`
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
	BillingCycleConfiguration param.Field[NewBillingCycleConfigurationParam] `json:"billing_cycle_configuration"`
	// The per unit conversion rate of the price currency to the invoicing currency.
	ConversionRate param.Field[float64] `json:"conversion_rate"`
	// The configuration for the rate of the price currency to the invoicing currency.
	ConversionRateConfig param.Field[NewPlanMatrixWithDisplayNamePriceUnionParam] `json:"conversion_rate_config"`
	// An ISO 4217 currency string, or custom pricing unit identifier, in which this
	// price is billed.
	Currency param.Field[string] `json:"currency"`
	// For dimensional price: specifies a price group and dimension values
	DimensionalPriceConfiguration param.Field[NewDimensionalPriceConfigurationParam] `json:"dimensional_price_configuration"`
	// An alias for the price.
	ExternalPriceID param.Field[string] `json:"external_price_id"`
	// If the Price represents a fixed cost, this represents the quantity of units
	// applied.
	FixedPriceQuantity param.Field[float64] `json:"fixed_price_quantity"`
	// The property used to group this price on an invoice
	InvoiceGroupingKey param.Field[string] `json:"invoice_grouping_key"`
	// Within each billing cycle, specifies the cadence at which invoices are produced.
	// If unspecified, a single invoice is produced per billing cycle.
	InvoicingCycleConfiguration param.Field[NewBillingCycleConfigurationParam] `json:"invoicing_cycle_configuration"`
	// User-specified key/value pairs for the resource. Individual keys can be removed
	// by setting the value to `null`, and the entire metadata mapping can be cleared
	// by setting `metadata` to `null`.
	Metadata param.Field[map[string]string] `json:"metadata"`
	// A transient ID that can be used to reference this price when adding adjustments
	// in the same API call.
	ReferenceID param.Field[string] `json:"reference_id"`
}

func (r NewPlanMatrixWithDisplayNamePriceParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r NewPlanMatrixWithDisplayNamePriceParam) ImplementsBetaNewPlanVersionParamsAddPricesPriceUnion() {
}

func (r NewPlanMatrixWithDisplayNamePriceParam) ImplementsBetaNewPlanVersionParamsReplacePricesPriceUnion() {
}

func (r NewPlanMatrixWithDisplayNamePriceParam) ImplementsBetaExternalPlanIDNewPlanVersionParamsAddPricesPriceUnion() {
}

func (r NewPlanMatrixWithDisplayNamePriceParam) ImplementsBetaExternalPlanIDNewPlanVersionParamsReplacePricesPriceUnion() {
}

func (r NewPlanMatrixWithDisplayNamePriceParam) ImplementsPlanNewParamsPricesPriceUnion() {}

// The cadence to bill for this price on.
type NewPlanMatrixWithDisplayNamePriceCadence string

const (
	NewPlanMatrixWithDisplayNamePriceCadenceAnnual     NewPlanMatrixWithDisplayNamePriceCadence = "annual"
	NewPlanMatrixWithDisplayNamePriceCadenceSemiAnnual NewPlanMatrixWithDisplayNamePriceCadence = "semi_annual"
	NewPlanMatrixWithDisplayNamePriceCadenceMonthly    NewPlanMatrixWithDisplayNamePriceCadence = "monthly"
	NewPlanMatrixWithDisplayNamePriceCadenceQuarterly  NewPlanMatrixWithDisplayNamePriceCadence = "quarterly"
	NewPlanMatrixWithDisplayNamePriceCadenceOneTime    NewPlanMatrixWithDisplayNamePriceCadence = "one_time"
	NewPlanMatrixWithDisplayNamePriceCadenceCustom     NewPlanMatrixWithDisplayNamePriceCadence = "custom"
)

func (r NewPlanMatrixWithDisplayNamePriceCadence) IsKnown() bool {
	switch r {
	case NewPlanMatrixWithDisplayNamePriceCadenceAnnual, NewPlanMatrixWithDisplayNamePriceCadenceSemiAnnual, NewPlanMatrixWithDisplayNamePriceCadenceMonthly, NewPlanMatrixWithDisplayNamePriceCadenceQuarterly, NewPlanMatrixWithDisplayNamePriceCadenceOneTime, NewPlanMatrixWithDisplayNamePriceCadenceCustom:
		return true
	}
	return false
}

type NewPlanMatrixWithDisplayNamePriceModelType string

const (
	NewPlanMatrixWithDisplayNamePriceModelTypeMatrixWithDisplayName NewPlanMatrixWithDisplayNamePriceModelType = "matrix_with_display_name"
)

func (r NewPlanMatrixWithDisplayNamePriceModelType) IsKnown() bool {
	switch r {
	case NewPlanMatrixWithDisplayNamePriceModelTypeMatrixWithDisplayName:
		return true
	}
	return false
}

type NewPlanMatrixWithDisplayNamePriceParam struct {
	ConversionRateType param.Field[NewPlanMatrixWithDisplayNamePriceConversionRateType] `json:"conversion_rate_type,required"`
	TieredConfig       param.Field[ConversionRateTieredConfigParam]                     `json:"tiered_config"`
	UnitConfig         param.Field[ConversionRateUnitConfigParam]                       `json:"unit_config"`
}

func (r NewPlanMatrixWithDisplayNamePriceParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r NewPlanMatrixWithDisplayNamePriceParam) ImplementsNewPlanMatrixWithDisplayNamePriceUnionParam() {
}

// Satisfied by [shared.UnitConversionRateConfigParam],
// [shared.TieredConversionRateConfigParam],
// [NewPlanMatrixWithDisplayNamePriceParam].
type NewPlanMatrixWithDisplayNamePriceUnionParam interface {
	ImplementsNewPlanMatrixWithDisplayNamePriceUnionParam()
}

type NewPlanMatrixWithDisplayNamePriceConversionRateType string

const (
	NewPlanMatrixWithDisplayNamePriceConversionRateTypeUnit   NewPlanMatrixWithDisplayNamePriceConversionRateType = "unit"
	NewPlanMatrixWithDisplayNamePriceConversionRateTypeTiered NewPlanMatrixWithDisplayNamePriceConversionRateType = "tiered"
)

func (r NewPlanMatrixWithDisplayNamePriceConversionRateType) IsKnown() bool {
	switch r {
	case NewPlanMatrixWithDisplayNamePriceConversionRateTypeUnit, NewPlanMatrixWithDisplayNamePriceConversionRateTypeTiered:
		return true
	}
	return false
}

type NewPlanMaxGroupTieredPackagePriceParam struct {
	// The cadence to bill for this price on.
	Cadence param.Field[NewPlanMaxGroupTieredPackagePriceCadence] `json:"cadence,required"`
	// The id of the item the price will be associated with.
	ItemID                      param.Field[string]                                     `json:"item_id,required"`
	MaxGroupTieredPackageConfig param.Field[map[string]interface{}]                     `json:"max_group_tiered_package_config,required"`
	ModelType                   param.Field[NewPlanMaxGroupTieredPackagePriceModelType] `json:"model_type,required"`
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
	BillingCycleConfiguration param.Field[NewBillingCycleConfigurationParam] `json:"billing_cycle_configuration"`
	// The per unit conversion rate of the price currency to the invoicing currency.
	ConversionRate param.Field[float64] `json:"conversion_rate"`
	// The configuration for the rate of the price currency to the invoicing currency.
	ConversionRateConfig param.Field[NewPlanMaxGroupTieredPackagePriceUnionParam] `json:"conversion_rate_config"`
	// An ISO 4217 currency string, or custom pricing unit identifier, in which this
	// price is billed.
	Currency param.Field[string] `json:"currency"`
	// For dimensional price: specifies a price group and dimension values
	DimensionalPriceConfiguration param.Field[NewDimensionalPriceConfigurationParam] `json:"dimensional_price_configuration"`
	// An alias for the price.
	ExternalPriceID param.Field[string] `json:"external_price_id"`
	// If the Price represents a fixed cost, this represents the quantity of units
	// applied.
	FixedPriceQuantity param.Field[float64] `json:"fixed_price_quantity"`
	// The property used to group this price on an invoice
	InvoiceGroupingKey param.Field[string] `json:"invoice_grouping_key"`
	// Within each billing cycle, specifies the cadence at which invoices are produced.
	// If unspecified, a single invoice is produced per billing cycle.
	InvoicingCycleConfiguration param.Field[NewBillingCycleConfigurationParam] `json:"invoicing_cycle_configuration"`
	// User-specified key/value pairs for the resource. Individual keys can be removed
	// by setting the value to `null`, and the entire metadata mapping can be cleared
	// by setting `metadata` to `null`.
	Metadata param.Field[map[string]string] `json:"metadata"`
	// A transient ID that can be used to reference this price when adding adjustments
	// in the same API call.
	ReferenceID param.Field[string] `json:"reference_id"`
}

func (r NewPlanMaxGroupTieredPackagePriceParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r NewPlanMaxGroupTieredPackagePriceParam) ImplementsBetaNewPlanVersionParamsAddPricesPriceUnion() {
}

func (r NewPlanMaxGroupTieredPackagePriceParam) ImplementsBetaNewPlanVersionParamsReplacePricesPriceUnion() {
}

func (r NewPlanMaxGroupTieredPackagePriceParam) ImplementsBetaExternalPlanIDNewPlanVersionParamsAddPricesPriceUnion() {
}

func (r NewPlanMaxGroupTieredPackagePriceParam) ImplementsBetaExternalPlanIDNewPlanVersionParamsReplacePricesPriceUnion() {
}

func (r NewPlanMaxGroupTieredPackagePriceParam) ImplementsPlanNewParamsPricesPriceUnion() {}

// The cadence to bill for this price on.
type NewPlanMaxGroupTieredPackagePriceCadence string

const (
	NewPlanMaxGroupTieredPackagePriceCadenceAnnual     NewPlanMaxGroupTieredPackagePriceCadence = "annual"
	NewPlanMaxGroupTieredPackagePriceCadenceSemiAnnual NewPlanMaxGroupTieredPackagePriceCadence = "semi_annual"
	NewPlanMaxGroupTieredPackagePriceCadenceMonthly    NewPlanMaxGroupTieredPackagePriceCadence = "monthly"
	NewPlanMaxGroupTieredPackagePriceCadenceQuarterly  NewPlanMaxGroupTieredPackagePriceCadence = "quarterly"
	NewPlanMaxGroupTieredPackagePriceCadenceOneTime    NewPlanMaxGroupTieredPackagePriceCadence = "one_time"
	NewPlanMaxGroupTieredPackagePriceCadenceCustom     NewPlanMaxGroupTieredPackagePriceCadence = "custom"
)

func (r NewPlanMaxGroupTieredPackagePriceCadence) IsKnown() bool {
	switch r {
	case NewPlanMaxGroupTieredPackagePriceCadenceAnnual, NewPlanMaxGroupTieredPackagePriceCadenceSemiAnnual, NewPlanMaxGroupTieredPackagePriceCadenceMonthly, NewPlanMaxGroupTieredPackagePriceCadenceQuarterly, NewPlanMaxGroupTieredPackagePriceCadenceOneTime, NewPlanMaxGroupTieredPackagePriceCadenceCustom:
		return true
	}
	return false
}

type NewPlanMaxGroupTieredPackagePriceModelType string

const (
	NewPlanMaxGroupTieredPackagePriceModelTypeMaxGroupTieredPackage NewPlanMaxGroupTieredPackagePriceModelType = "max_group_tiered_package"
)

func (r NewPlanMaxGroupTieredPackagePriceModelType) IsKnown() bool {
	switch r {
	case NewPlanMaxGroupTieredPackagePriceModelTypeMaxGroupTieredPackage:
		return true
	}
	return false
}

type NewPlanMaxGroupTieredPackagePriceParam struct {
	ConversionRateType param.Field[NewPlanMaxGroupTieredPackagePriceConversionRateType] `json:"conversion_rate_type,required"`
	TieredConfig       param.Field[ConversionRateTieredConfigParam]                     `json:"tiered_config"`
	UnitConfig         param.Field[ConversionRateUnitConfigParam]                       `json:"unit_config"`
}

func (r NewPlanMaxGroupTieredPackagePriceParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r NewPlanMaxGroupTieredPackagePriceParam) ImplementsNewPlanMaxGroupTieredPackagePriceUnionParam() {
}

// Satisfied by [shared.UnitConversionRateConfigParam],
// [shared.TieredConversionRateConfigParam],
// [NewPlanMaxGroupTieredPackagePriceParam].
type NewPlanMaxGroupTieredPackagePriceUnionParam interface {
	ImplementsNewPlanMaxGroupTieredPackagePriceUnionParam()
}

type NewPlanMaxGroupTieredPackagePriceConversionRateType string

const (
	NewPlanMaxGroupTieredPackagePriceConversionRateTypeUnit   NewPlanMaxGroupTieredPackagePriceConversionRateType = "unit"
	NewPlanMaxGroupTieredPackagePriceConversionRateTypeTiered NewPlanMaxGroupTieredPackagePriceConversionRateType = "tiered"
)

func (r NewPlanMaxGroupTieredPackagePriceConversionRateType) IsKnown() bool {
	switch r {
	case NewPlanMaxGroupTieredPackagePriceConversionRateTypeUnit, NewPlanMaxGroupTieredPackagePriceConversionRateTypeTiered:
		return true
	}
	return false
}

type NewPlanPackagePriceParam struct {
	// The cadence to bill for this price on.
	Cadence param.Field[NewPlanPackagePriceCadence] `json:"cadence,required"`
	// The id of the item the price will be associated with.
	ItemID    param.Field[string]                       `json:"item_id,required"`
	ModelType param.Field[NewPlanPackagePriceModelType] `json:"model_type,required"`
	// The name of the price.
	Name          param.Field[string]             `json:"name,required"`
	PackageConfig param.Field[PackageConfigParam] `json:"package_config,required"`
	// The id of the billable metric for the price. Only needed if the price is
	// usage-based.
	BillableMetricID param.Field[string] `json:"billable_metric_id"`
	// If the Price represents a fixed cost, the price will be billed in-advance if
	// this is true, and in-arrears if this is false.
	BilledInAdvance param.Field[bool] `json:"billed_in_advance"`
	// For custom cadence: specifies the duration of the billing period in days or
	// months.
	BillingCycleConfiguration param.Field[NewBillingCycleConfigurationParam] `json:"billing_cycle_configuration"`
	// The per unit conversion rate of the price currency to the invoicing currency.
	ConversionRate param.Field[float64] `json:"conversion_rate"`
	// The configuration for the rate of the price currency to the invoicing currency.
	ConversionRateConfig param.Field[NewPlanPackagePriceUnionParam] `json:"conversion_rate_config"`
	// An ISO 4217 currency string, or custom pricing unit identifier, in which this
	// price is billed.
	Currency param.Field[string] `json:"currency"`
	// For dimensional price: specifies a price group and dimension values
	DimensionalPriceConfiguration param.Field[NewDimensionalPriceConfigurationParam] `json:"dimensional_price_configuration"`
	// An alias for the price.
	ExternalPriceID param.Field[string] `json:"external_price_id"`
	// If the Price represents a fixed cost, this represents the quantity of units
	// applied.
	FixedPriceQuantity param.Field[float64] `json:"fixed_price_quantity"`
	// The property used to group this price on an invoice
	InvoiceGroupingKey param.Field[string] `json:"invoice_grouping_key"`
	// Within each billing cycle, specifies the cadence at which invoices are produced.
	// If unspecified, a single invoice is produced per billing cycle.
	InvoicingCycleConfiguration param.Field[NewBillingCycleConfigurationParam] `json:"invoicing_cycle_configuration"`
	// User-specified key/value pairs for the resource. Individual keys can be removed
	// by setting the value to `null`, and the entire metadata mapping can be cleared
	// by setting `metadata` to `null`.
	Metadata param.Field[map[string]string] `json:"metadata"`
	// A transient ID that can be used to reference this price when adding adjustments
	// in the same API call.
	ReferenceID param.Field[string] `json:"reference_id"`
}

func (r NewPlanPackagePriceParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r NewPlanPackagePriceParam) ImplementsBetaNewPlanVersionParamsAddPricesPriceUnion() {}

func (r NewPlanPackagePriceParam) ImplementsBetaNewPlanVersionParamsReplacePricesPriceUnion() {}

func (r NewPlanPackagePriceParam) ImplementsBetaExternalPlanIDNewPlanVersionParamsAddPricesPriceUnion() {
}

func (r NewPlanPackagePriceParam) ImplementsBetaExternalPlanIDNewPlanVersionParamsReplacePricesPriceUnion() {
}

func (r NewPlanPackagePriceParam) ImplementsPlanNewParamsPricesPriceUnion() {}

// The cadence to bill for this price on.
type NewPlanPackagePriceCadence string

const (
	NewPlanPackagePriceCadenceAnnual     NewPlanPackagePriceCadence = "annual"
	NewPlanPackagePriceCadenceSemiAnnual NewPlanPackagePriceCadence = "semi_annual"
	NewPlanPackagePriceCadenceMonthly    NewPlanPackagePriceCadence = "monthly"
	NewPlanPackagePriceCadenceQuarterly  NewPlanPackagePriceCadence = "quarterly"
	NewPlanPackagePriceCadenceOneTime    NewPlanPackagePriceCadence = "one_time"
	NewPlanPackagePriceCadenceCustom     NewPlanPackagePriceCadence = "custom"
)

func (r NewPlanPackagePriceCadence) IsKnown() bool {
	switch r {
	case NewPlanPackagePriceCadenceAnnual, NewPlanPackagePriceCadenceSemiAnnual, NewPlanPackagePriceCadenceMonthly, NewPlanPackagePriceCadenceQuarterly, NewPlanPackagePriceCadenceOneTime, NewPlanPackagePriceCadenceCustom:
		return true
	}
	return false
}

type NewPlanPackagePriceModelType string

const (
	NewPlanPackagePriceModelTypePackage NewPlanPackagePriceModelType = "package"
)

func (r NewPlanPackagePriceModelType) IsKnown() bool {
	switch r {
	case NewPlanPackagePriceModelTypePackage:
		return true
	}
	return false
}

type NewPlanPackagePriceParam struct {
	ConversionRateType param.Field[NewPlanPackagePriceConversionRateType] `json:"conversion_rate_type,required"`
	TieredConfig       param.Field[ConversionRateTieredConfigParam]       `json:"tiered_config"`
	UnitConfig         param.Field[ConversionRateUnitConfigParam]         `json:"unit_config"`
}

func (r NewPlanPackagePriceParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r NewPlanPackagePriceParam) ImplementsNewPlanPackagePriceUnionParam() {}

// Satisfied by [shared.UnitConversionRateConfigParam],
// [shared.TieredConversionRateConfigParam], [NewPlanPackagePriceParam].
type NewPlanPackagePriceUnionParam interface {
	ImplementsNewPlanPackagePriceUnionParam()
}

type NewPlanPackagePriceConversionRateType string

const (
	NewPlanPackagePriceConversionRateTypeUnit   NewPlanPackagePriceConversionRateType = "unit"
	NewPlanPackagePriceConversionRateTypeTiered NewPlanPackagePriceConversionRateType = "tiered"
)

func (r NewPlanPackagePriceConversionRateType) IsKnown() bool {
	switch r {
	case NewPlanPackagePriceConversionRateTypeUnit, NewPlanPackagePriceConversionRateTypeTiered:
		return true
	}
	return false
}

type NewPlanPackageWithAllocationPriceParam struct {
	// The cadence to bill for this price on.
	Cadence param.Field[NewPlanPackageWithAllocationPriceCadence] `json:"cadence,required"`
	// The id of the item the price will be associated with.
	ItemID    param.Field[string]                                     `json:"item_id,required"`
	ModelType param.Field[NewPlanPackageWithAllocationPriceModelType] `json:"model_type,required"`
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
	BillingCycleConfiguration param.Field[NewBillingCycleConfigurationParam] `json:"billing_cycle_configuration"`
	// The per unit conversion rate of the price currency to the invoicing currency.
	ConversionRate param.Field[float64] `json:"conversion_rate"`
	// The configuration for the rate of the price currency to the invoicing currency.
	ConversionRateConfig param.Field[NewPlanPackageWithAllocationPriceUnionParam] `json:"conversion_rate_config"`
	// An ISO 4217 currency string, or custom pricing unit identifier, in which this
	// price is billed.
	Currency param.Field[string] `json:"currency"`
	// For dimensional price: specifies a price group and dimension values
	DimensionalPriceConfiguration param.Field[NewDimensionalPriceConfigurationParam] `json:"dimensional_price_configuration"`
	// An alias for the price.
	ExternalPriceID param.Field[string] `json:"external_price_id"`
	// If the Price represents a fixed cost, this represents the quantity of units
	// applied.
	FixedPriceQuantity param.Field[float64] `json:"fixed_price_quantity"`
	// The property used to group this price on an invoice
	InvoiceGroupingKey param.Field[string] `json:"invoice_grouping_key"`
	// Within each billing cycle, specifies the cadence at which invoices are produced.
	// If unspecified, a single invoice is produced per billing cycle.
	InvoicingCycleConfiguration param.Field[NewBillingCycleConfigurationParam] `json:"invoicing_cycle_configuration"`
	// User-specified key/value pairs for the resource. Individual keys can be removed
	// by setting the value to `null`, and the entire metadata mapping can be cleared
	// by setting `metadata` to `null`.
	Metadata param.Field[map[string]string] `json:"metadata"`
	// A transient ID that can be used to reference this price when adding adjustments
	// in the same API call.
	ReferenceID param.Field[string] `json:"reference_id"`
}

func (r NewPlanPackageWithAllocationPriceParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r NewPlanPackageWithAllocationPriceParam) ImplementsBetaNewPlanVersionParamsAddPricesPriceUnion() {
}

func (r NewPlanPackageWithAllocationPriceParam) ImplementsBetaNewPlanVersionParamsReplacePricesPriceUnion() {
}

func (r NewPlanPackageWithAllocationPriceParam) ImplementsBetaExternalPlanIDNewPlanVersionParamsAddPricesPriceUnion() {
}

func (r NewPlanPackageWithAllocationPriceParam) ImplementsBetaExternalPlanIDNewPlanVersionParamsReplacePricesPriceUnion() {
}

func (r NewPlanPackageWithAllocationPriceParam) ImplementsPlanNewParamsPricesPriceUnion() {}

// The cadence to bill for this price on.
type NewPlanPackageWithAllocationPriceCadence string

const (
	NewPlanPackageWithAllocationPriceCadenceAnnual     NewPlanPackageWithAllocationPriceCadence = "annual"
	NewPlanPackageWithAllocationPriceCadenceSemiAnnual NewPlanPackageWithAllocationPriceCadence = "semi_annual"
	NewPlanPackageWithAllocationPriceCadenceMonthly    NewPlanPackageWithAllocationPriceCadence = "monthly"
	NewPlanPackageWithAllocationPriceCadenceQuarterly  NewPlanPackageWithAllocationPriceCadence = "quarterly"
	NewPlanPackageWithAllocationPriceCadenceOneTime    NewPlanPackageWithAllocationPriceCadence = "one_time"
	NewPlanPackageWithAllocationPriceCadenceCustom     NewPlanPackageWithAllocationPriceCadence = "custom"
)

func (r NewPlanPackageWithAllocationPriceCadence) IsKnown() bool {
	switch r {
	case NewPlanPackageWithAllocationPriceCadenceAnnual, NewPlanPackageWithAllocationPriceCadenceSemiAnnual, NewPlanPackageWithAllocationPriceCadenceMonthly, NewPlanPackageWithAllocationPriceCadenceQuarterly, NewPlanPackageWithAllocationPriceCadenceOneTime, NewPlanPackageWithAllocationPriceCadenceCustom:
		return true
	}
	return false
}

type NewPlanPackageWithAllocationPriceModelType string

const (
	NewPlanPackageWithAllocationPriceModelTypePackageWithAllocation NewPlanPackageWithAllocationPriceModelType = "package_with_allocation"
)

func (r NewPlanPackageWithAllocationPriceModelType) IsKnown() bool {
	switch r {
	case NewPlanPackageWithAllocationPriceModelTypePackageWithAllocation:
		return true
	}
	return false
}

type NewPlanPackageWithAllocationPriceParam struct {
	ConversionRateType param.Field[NewPlanPackageWithAllocationPriceConversionRateType] `json:"conversion_rate_type,required"`
	TieredConfig       param.Field[ConversionRateTieredConfigParam]                     `json:"tiered_config"`
	UnitConfig         param.Field[ConversionRateUnitConfigParam]                       `json:"unit_config"`
}

func (r NewPlanPackageWithAllocationPriceParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r NewPlanPackageWithAllocationPriceParam) ImplementsNewPlanPackageWithAllocationPriceUnionParam() {
}

// Satisfied by [shared.UnitConversionRateConfigParam],
// [shared.TieredConversionRateConfigParam],
// [NewPlanPackageWithAllocationPriceParam].
type NewPlanPackageWithAllocationPriceUnionParam interface {
	ImplementsNewPlanPackageWithAllocationPriceUnionParam()
}

type NewPlanPackageWithAllocationPriceConversionRateType string

const (
	NewPlanPackageWithAllocationPriceConversionRateTypeUnit   NewPlanPackageWithAllocationPriceConversionRateType = "unit"
	NewPlanPackageWithAllocationPriceConversionRateTypeTiered NewPlanPackageWithAllocationPriceConversionRateType = "tiered"
)

func (r NewPlanPackageWithAllocationPriceConversionRateType) IsKnown() bool {
	switch r {
	case NewPlanPackageWithAllocationPriceConversionRateTypeUnit, NewPlanPackageWithAllocationPriceConversionRateTypeTiered:
		return true
	}
	return false
}

type NewPlanScalableMatrixWithTieredPricingPriceParam struct {
	// The cadence to bill for this price on.
	Cadence param.Field[NewPlanScalableMatrixWithTieredPricingPriceCadence] `json:"cadence,required"`
	// The id of the item the price will be associated with.
	ItemID    param.Field[string]                                               `json:"item_id,required"`
	ModelType param.Field[NewPlanScalableMatrixWithTieredPricingPriceModelType] `json:"model_type,required"`
	// The name of the price.
	Name                                  param.Field[string]                 `json:"name,required"`
	ScalableMatrixWithTieredPricingConfig param.Field[map[string]interface{}] `json:"scalable_matrix_with_tiered_pricing_config,required"`
	// The id of the billable metric for the price. Only needed if the price is
	// usage-based.
	BillableMetricID param.Field[string] `json:"billable_metric_id"`
	// If the Price represents a fixed cost, the price will be billed in-advance if
	// this is true, and in-arrears if this is false.
	BilledInAdvance param.Field[bool] `json:"billed_in_advance"`
	// For custom cadence: specifies the duration of the billing period in days or
	// months.
	BillingCycleConfiguration param.Field[NewBillingCycleConfigurationParam] `json:"billing_cycle_configuration"`
	// The per unit conversion rate of the price currency to the invoicing currency.
	ConversionRate param.Field[float64] `json:"conversion_rate"`
	// The configuration for the rate of the price currency to the invoicing currency.
	ConversionRateConfig param.Field[NewPlanScalableMatrixWithTieredPricingPriceUnionParam] `json:"conversion_rate_config"`
	// An ISO 4217 currency string, or custom pricing unit identifier, in which this
	// price is billed.
	Currency param.Field[string] `json:"currency"`
	// For dimensional price: specifies a price group and dimension values
	DimensionalPriceConfiguration param.Field[NewDimensionalPriceConfigurationParam] `json:"dimensional_price_configuration"`
	// An alias for the price.
	ExternalPriceID param.Field[string] `json:"external_price_id"`
	// If the Price represents a fixed cost, this represents the quantity of units
	// applied.
	FixedPriceQuantity param.Field[float64] `json:"fixed_price_quantity"`
	// The property used to group this price on an invoice
	InvoiceGroupingKey param.Field[string] `json:"invoice_grouping_key"`
	// Within each billing cycle, specifies the cadence at which invoices are produced.
	// If unspecified, a single invoice is produced per billing cycle.
	InvoicingCycleConfiguration param.Field[NewBillingCycleConfigurationParam] `json:"invoicing_cycle_configuration"`
	// User-specified key/value pairs for the resource. Individual keys can be removed
	// by setting the value to `null`, and the entire metadata mapping can be cleared
	// by setting `metadata` to `null`.
	Metadata param.Field[map[string]string] `json:"metadata"`
	// A transient ID that can be used to reference this price when adding adjustments
	// in the same API call.
	ReferenceID param.Field[string] `json:"reference_id"`
}

func (r NewPlanScalableMatrixWithTieredPricingPriceParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r NewPlanScalableMatrixWithTieredPricingPriceParam) ImplementsBetaNewPlanVersionParamsAddPricesPriceUnion() {
}

func (r NewPlanScalableMatrixWithTieredPricingPriceParam) ImplementsBetaNewPlanVersionParamsReplacePricesPriceUnion() {
}

func (r NewPlanScalableMatrixWithTieredPricingPriceParam) ImplementsBetaExternalPlanIDNewPlanVersionParamsAddPricesPriceUnion() {
}

func (r NewPlanScalableMatrixWithTieredPricingPriceParam) ImplementsBetaExternalPlanIDNewPlanVersionParamsReplacePricesPriceUnion() {
}

func (r NewPlanScalableMatrixWithTieredPricingPriceParam) ImplementsPlanNewParamsPricesPriceUnion() {}

// The cadence to bill for this price on.
type NewPlanScalableMatrixWithTieredPricingPriceCadence string

const (
	NewPlanScalableMatrixWithTieredPricingPriceCadenceAnnual     NewPlanScalableMatrixWithTieredPricingPriceCadence = "annual"
	NewPlanScalableMatrixWithTieredPricingPriceCadenceSemiAnnual NewPlanScalableMatrixWithTieredPricingPriceCadence = "semi_annual"
	NewPlanScalableMatrixWithTieredPricingPriceCadenceMonthly    NewPlanScalableMatrixWithTieredPricingPriceCadence = "monthly"
	NewPlanScalableMatrixWithTieredPricingPriceCadenceQuarterly  NewPlanScalableMatrixWithTieredPricingPriceCadence = "quarterly"
	NewPlanScalableMatrixWithTieredPricingPriceCadenceOneTime    NewPlanScalableMatrixWithTieredPricingPriceCadence = "one_time"
	NewPlanScalableMatrixWithTieredPricingPriceCadenceCustom     NewPlanScalableMatrixWithTieredPricingPriceCadence = "custom"
)

func (r NewPlanScalableMatrixWithTieredPricingPriceCadence) IsKnown() bool {
	switch r {
	case NewPlanScalableMatrixWithTieredPricingPriceCadenceAnnual, NewPlanScalableMatrixWithTieredPricingPriceCadenceSemiAnnual, NewPlanScalableMatrixWithTieredPricingPriceCadenceMonthly, NewPlanScalableMatrixWithTieredPricingPriceCadenceQuarterly, NewPlanScalableMatrixWithTieredPricingPriceCadenceOneTime, NewPlanScalableMatrixWithTieredPricingPriceCadenceCustom:
		return true
	}
	return false
}

type NewPlanScalableMatrixWithTieredPricingPriceModelType string

const (
	NewPlanScalableMatrixWithTieredPricingPriceModelTypeScalableMatrixWithTieredPricing NewPlanScalableMatrixWithTieredPricingPriceModelType = "scalable_matrix_with_tiered_pricing"
)

func (r NewPlanScalableMatrixWithTieredPricingPriceModelType) IsKnown() bool {
	switch r {
	case NewPlanScalableMatrixWithTieredPricingPriceModelTypeScalableMatrixWithTieredPricing:
		return true
	}
	return false
}

type NewPlanScalableMatrixWithTieredPricingPriceParam struct {
	ConversionRateType param.Field[NewPlanScalableMatrixWithTieredPricingPriceConversionRateType] `json:"conversion_rate_type,required"`
	TieredConfig       param.Field[ConversionRateTieredConfigParam]                               `json:"tiered_config"`
	UnitConfig         param.Field[ConversionRateUnitConfigParam]                                 `json:"unit_config"`
}

func (r NewPlanScalableMatrixWithTieredPricingPriceParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r NewPlanScalableMatrixWithTieredPricingPriceParam) ImplementsNewPlanScalableMatrixWithTieredPricingPriceUnionParam() {
}

// Satisfied by [shared.UnitConversionRateConfigParam],
// [shared.TieredConversionRateConfigParam],
// [NewPlanScalableMatrixWithTieredPricingPriceParam].
type NewPlanScalableMatrixWithTieredPricingPriceUnionParam interface {
	ImplementsNewPlanScalableMatrixWithTieredPricingPriceUnionParam()
}

type NewPlanScalableMatrixWithTieredPricingPriceConversionRateType string

const (
	NewPlanScalableMatrixWithTieredPricingPriceConversionRateTypeUnit   NewPlanScalableMatrixWithTieredPricingPriceConversionRateType = "unit"
	NewPlanScalableMatrixWithTieredPricingPriceConversionRateTypeTiered NewPlanScalableMatrixWithTieredPricingPriceConversionRateType = "tiered"
)

func (r NewPlanScalableMatrixWithTieredPricingPriceConversionRateType) IsKnown() bool {
	switch r {
	case NewPlanScalableMatrixWithTieredPricingPriceConversionRateTypeUnit, NewPlanScalableMatrixWithTieredPricingPriceConversionRateTypeTiered:
		return true
	}
	return false
}

type NewPlanScalableMatrixWithUnitPricingPriceParam struct {
	// The cadence to bill for this price on.
	Cadence param.Field[NewPlanScalableMatrixWithUnitPricingPriceCadence] `json:"cadence,required"`
	// The id of the item the price will be associated with.
	ItemID    param.Field[string]                                             `json:"item_id,required"`
	ModelType param.Field[NewPlanScalableMatrixWithUnitPricingPriceModelType] `json:"model_type,required"`
	// The name of the price.
	Name                                param.Field[string]                 `json:"name,required"`
	ScalableMatrixWithUnitPricingConfig param.Field[map[string]interface{}] `json:"scalable_matrix_with_unit_pricing_config,required"`
	// The id of the billable metric for the price. Only needed if the price is
	// usage-based.
	BillableMetricID param.Field[string] `json:"billable_metric_id"`
	// If the Price represents a fixed cost, the price will be billed in-advance if
	// this is true, and in-arrears if this is false.
	BilledInAdvance param.Field[bool] `json:"billed_in_advance"`
	// For custom cadence: specifies the duration of the billing period in days or
	// months.
	BillingCycleConfiguration param.Field[NewBillingCycleConfigurationParam] `json:"billing_cycle_configuration"`
	// The per unit conversion rate of the price currency to the invoicing currency.
	ConversionRate param.Field[float64] `json:"conversion_rate"`
	// The configuration for the rate of the price currency to the invoicing currency.
	ConversionRateConfig param.Field[NewPlanScalableMatrixWithUnitPricingPriceUnionParam] `json:"conversion_rate_config"`
	// An ISO 4217 currency string, or custom pricing unit identifier, in which this
	// price is billed.
	Currency param.Field[string] `json:"currency"`
	// For dimensional price: specifies a price group and dimension values
	DimensionalPriceConfiguration param.Field[NewDimensionalPriceConfigurationParam] `json:"dimensional_price_configuration"`
	// An alias for the price.
	ExternalPriceID param.Field[string] `json:"external_price_id"`
	// If the Price represents a fixed cost, this represents the quantity of units
	// applied.
	FixedPriceQuantity param.Field[float64] `json:"fixed_price_quantity"`
	// The property used to group this price on an invoice
	InvoiceGroupingKey param.Field[string] `json:"invoice_grouping_key"`
	// Within each billing cycle, specifies the cadence at which invoices are produced.
	// If unspecified, a single invoice is produced per billing cycle.
	InvoicingCycleConfiguration param.Field[NewBillingCycleConfigurationParam] `json:"invoicing_cycle_configuration"`
	// User-specified key/value pairs for the resource. Individual keys can be removed
	// by setting the value to `null`, and the entire metadata mapping can be cleared
	// by setting `metadata` to `null`.
	Metadata param.Field[map[string]string] `json:"metadata"`
	// A transient ID that can be used to reference this price when adding adjustments
	// in the same API call.
	ReferenceID param.Field[string] `json:"reference_id"`
}

func (r NewPlanScalableMatrixWithUnitPricingPriceParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r NewPlanScalableMatrixWithUnitPricingPriceParam) ImplementsBetaNewPlanVersionParamsAddPricesPriceUnion() {
}

func (r NewPlanScalableMatrixWithUnitPricingPriceParam) ImplementsBetaNewPlanVersionParamsReplacePricesPriceUnion() {
}

func (r NewPlanScalableMatrixWithUnitPricingPriceParam) ImplementsBetaExternalPlanIDNewPlanVersionParamsAddPricesPriceUnion() {
}

func (r NewPlanScalableMatrixWithUnitPricingPriceParam) ImplementsBetaExternalPlanIDNewPlanVersionParamsReplacePricesPriceUnion() {
}

func (r NewPlanScalableMatrixWithUnitPricingPriceParam) ImplementsPlanNewParamsPricesPriceUnion() {}

// The cadence to bill for this price on.
type NewPlanScalableMatrixWithUnitPricingPriceCadence string

const (
	NewPlanScalableMatrixWithUnitPricingPriceCadenceAnnual     NewPlanScalableMatrixWithUnitPricingPriceCadence = "annual"
	NewPlanScalableMatrixWithUnitPricingPriceCadenceSemiAnnual NewPlanScalableMatrixWithUnitPricingPriceCadence = "semi_annual"
	NewPlanScalableMatrixWithUnitPricingPriceCadenceMonthly    NewPlanScalableMatrixWithUnitPricingPriceCadence = "monthly"
	NewPlanScalableMatrixWithUnitPricingPriceCadenceQuarterly  NewPlanScalableMatrixWithUnitPricingPriceCadence = "quarterly"
	NewPlanScalableMatrixWithUnitPricingPriceCadenceOneTime    NewPlanScalableMatrixWithUnitPricingPriceCadence = "one_time"
	NewPlanScalableMatrixWithUnitPricingPriceCadenceCustom     NewPlanScalableMatrixWithUnitPricingPriceCadence = "custom"
)

func (r NewPlanScalableMatrixWithUnitPricingPriceCadence) IsKnown() bool {
	switch r {
	case NewPlanScalableMatrixWithUnitPricingPriceCadenceAnnual, NewPlanScalableMatrixWithUnitPricingPriceCadenceSemiAnnual, NewPlanScalableMatrixWithUnitPricingPriceCadenceMonthly, NewPlanScalableMatrixWithUnitPricingPriceCadenceQuarterly, NewPlanScalableMatrixWithUnitPricingPriceCadenceOneTime, NewPlanScalableMatrixWithUnitPricingPriceCadenceCustom:
		return true
	}
	return false
}

type NewPlanScalableMatrixWithUnitPricingPriceModelType string

const (
	NewPlanScalableMatrixWithUnitPricingPriceModelTypeScalableMatrixWithUnitPricing NewPlanScalableMatrixWithUnitPricingPriceModelType = "scalable_matrix_with_unit_pricing"
)

func (r NewPlanScalableMatrixWithUnitPricingPriceModelType) IsKnown() bool {
	switch r {
	case NewPlanScalableMatrixWithUnitPricingPriceModelTypeScalableMatrixWithUnitPricing:
		return true
	}
	return false
}

type NewPlanScalableMatrixWithUnitPricingPriceParam struct {
	ConversionRateType param.Field[NewPlanScalableMatrixWithUnitPricingPriceConversionRateType] `json:"conversion_rate_type,required"`
	TieredConfig       param.Field[ConversionRateTieredConfigParam]                             `json:"tiered_config"`
	UnitConfig         param.Field[ConversionRateUnitConfigParam]                               `json:"unit_config"`
}

func (r NewPlanScalableMatrixWithUnitPricingPriceParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r NewPlanScalableMatrixWithUnitPricingPriceParam) ImplementsNewPlanScalableMatrixWithUnitPricingPriceUnionParam() {
}

// Satisfied by [shared.UnitConversionRateConfigParam],
// [shared.TieredConversionRateConfigParam],
// [NewPlanScalableMatrixWithUnitPricingPriceParam].
type NewPlanScalableMatrixWithUnitPricingPriceUnionParam interface {
	ImplementsNewPlanScalableMatrixWithUnitPricingPriceUnionParam()
}

type NewPlanScalableMatrixWithUnitPricingPriceConversionRateType string

const (
	NewPlanScalableMatrixWithUnitPricingPriceConversionRateTypeUnit   NewPlanScalableMatrixWithUnitPricingPriceConversionRateType = "unit"
	NewPlanScalableMatrixWithUnitPricingPriceConversionRateTypeTiered NewPlanScalableMatrixWithUnitPricingPriceConversionRateType = "tiered"
)

func (r NewPlanScalableMatrixWithUnitPricingPriceConversionRateType) IsKnown() bool {
	switch r {
	case NewPlanScalableMatrixWithUnitPricingPriceConversionRateTypeUnit, NewPlanScalableMatrixWithUnitPricingPriceConversionRateTypeTiered:
		return true
	}
	return false
}

type NewPlanThresholdTotalAmountPriceParam struct {
	// The cadence to bill for this price on.
	Cadence param.Field[NewPlanThresholdTotalAmountPriceCadence] `json:"cadence,required"`
	// The id of the item the price will be associated with.
	ItemID    param.Field[string]                                    `json:"item_id,required"`
	ModelType param.Field[NewPlanThresholdTotalAmountPriceModelType] `json:"model_type,required"`
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
	BillingCycleConfiguration param.Field[NewBillingCycleConfigurationParam] `json:"billing_cycle_configuration"`
	// The per unit conversion rate of the price currency to the invoicing currency.
	ConversionRate param.Field[float64] `json:"conversion_rate"`
	// The configuration for the rate of the price currency to the invoicing currency.
	ConversionRateConfig param.Field[NewPlanThresholdTotalAmountPriceUnionParam] `json:"conversion_rate_config"`
	// An ISO 4217 currency string, or custom pricing unit identifier, in which this
	// price is billed.
	Currency param.Field[string] `json:"currency"`
	// For dimensional price: specifies a price group and dimension values
	DimensionalPriceConfiguration param.Field[NewDimensionalPriceConfigurationParam] `json:"dimensional_price_configuration"`
	// An alias for the price.
	ExternalPriceID param.Field[string] `json:"external_price_id"`
	// If the Price represents a fixed cost, this represents the quantity of units
	// applied.
	FixedPriceQuantity param.Field[float64] `json:"fixed_price_quantity"`
	// The property used to group this price on an invoice
	InvoiceGroupingKey param.Field[string] `json:"invoice_grouping_key"`
	// Within each billing cycle, specifies the cadence at which invoices are produced.
	// If unspecified, a single invoice is produced per billing cycle.
	InvoicingCycleConfiguration param.Field[NewBillingCycleConfigurationParam] `json:"invoicing_cycle_configuration"`
	// User-specified key/value pairs for the resource. Individual keys can be removed
	// by setting the value to `null`, and the entire metadata mapping can be cleared
	// by setting `metadata` to `null`.
	Metadata param.Field[map[string]string] `json:"metadata"`
	// A transient ID that can be used to reference this price when adding adjustments
	// in the same API call.
	ReferenceID param.Field[string] `json:"reference_id"`
}

func (r NewPlanThresholdTotalAmountPriceParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r NewPlanThresholdTotalAmountPriceParam) ImplementsBetaNewPlanVersionParamsAddPricesPriceUnion() {
}

func (r NewPlanThresholdTotalAmountPriceParam) ImplementsBetaNewPlanVersionParamsReplacePricesPriceUnion() {
}

func (r NewPlanThresholdTotalAmountPriceParam) ImplementsBetaExternalPlanIDNewPlanVersionParamsAddPricesPriceUnion() {
}

func (r NewPlanThresholdTotalAmountPriceParam) ImplementsBetaExternalPlanIDNewPlanVersionParamsReplacePricesPriceUnion() {
}

func (r NewPlanThresholdTotalAmountPriceParam) ImplementsPlanNewParamsPricesPriceUnion() {}

// The cadence to bill for this price on.
type NewPlanThresholdTotalAmountPriceCadence string

const (
	NewPlanThresholdTotalAmountPriceCadenceAnnual     NewPlanThresholdTotalAmountPriceCadence = "annual"
	NewPlanThresholdTotalAmountPriceCadenceSemiAnnual NewPlanThresholdTotalAmountPriceCadence = "semi_annual"
	NewPlanThresholdTotalAmountPriceCadenceMonthly    NewPlanThresholdTotalAmountPriceCadence = "monthly"
	NewPlanThresholdTotalAmountPriceCadenceQuarterly  NewPlanThresholdTotalAmountPriceCadence = "quarterly"
	NewPlanThresholdTotalAmountPriceCadenceOneTime    NewPlanThresholdTotalAmountPriceCadence = "one_time"
	NewPlanThresholdTotalAmountPriceCadenceCustom     NewPlanThresholdTotalAmountPriceCadence = "custom"
)

func (r NewPlanThresholdTotalAmountPriceCadence) IsKnown() bool {
	switch r {
	case NewPlanThresholdTotalAmountPriceCadenceAnnual, NewPlanThresholdTotalAmountPriceCadenceSemiAnnual, NewPlanThresholdTotalAmountPriceCadenceMonthly, NewPlanThresholdTotalAmountPriceCadenceQuarterly, NewPlanThresholdTotalAmountPriceCadenceOneTime, NewPlanThresholdTotalAmountPriceCadenceCustom:
		return true
	}
	return false
}

type NewPlanThresholdTotalAmountPriceModelType string

const (
	NewPlanThresholdTotalAmountPriceModelTypeThresholdTotalAmount NewPlanThresholdTotalAmountPriceModelType = "threshold_total_amount"
)

func (r NewPlanThresholdTotalAmountPriceModelType) IsKnown() bool {
	switch r {
	case NewPlanThresholdTotalAmountPriceModelTypeThresholdTotalAmount:
		return true
	}
	return false
}

type NewPlanThresholdTotalAmountPriceParam struct {
	ConversionRateType param.Field[NewPlanThresholdTotalAmountPriceConversionRateType] `json:"conversion_rate_type,required"`
	TieredConfig       param.Field[ConversionRateTieredConfigParam]                    `json:"tiered_config"`
	UnitConfig         param.Field[ConversionRateUnitConfigParam]                      `json:"unit_config"`
}

func (r NewPlanThresholdTotalAmountPriceParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r NewPlanThresholdTotalAmountPriceParam) ImplementsNewPlanThresholdTotalAmountPriceUnionParam() {
}

// Satisfied by [shared.UnitConversionRateConfigParam],
// [shared.TieredConversionRateConfigParam],
// [NewPlanThresholdTotalAmountPriceParam].
type NewPlanThresholdTotalAmountPriceUnionParam interface {
	ImplementsNewPlanThresholdTotalAmountPriceUnionParam()
}

type NewPlanThresholdTotalAmountPriceConversionRateType string

const (
	NewPlanThresholdTotalAmountPriceConversionRateTypeUnit   NewPlanThresholdTotalAmountPriceConversionRateType = "unit"
	NewPlanThresholdTotalAmountPriceConversionRateTypeTiered NewPlanThresholdTotalAmountPriceConversionRateType = "tiered"
)

func (r NewPlanThresholdTotalAmountPriceConversionRateType) IsKnown() bool {
	switch r {
	case NewPlanThresholdTotalAmountPriceConversionRateTypeUnit, NewPlanThresholdTotalAmountPriceConversionRateTypeTiered:
		return true
	}
	return false
}

type NewPlanTierWithProrationPriceParam struct {
	// The cadence to bill for this price on.
	Cadence param.Field[NewPlanTierWithProrationPriceCadence] `json:"cadence,required"`
	// The id of the item the price will be associated with.
	ItemID    param.Field[string]                                 `json:"item_id,required"`
	ModelType param.Field[NewPlanTierWithProrationPriceModelType] `json:"model_type,required"`
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
	BillingCycleConfiguration param.Field[NewBillingCycleConfigurationParam] `json:"billing_cycle_configuration"`
	// The per unit conversion rate of the price currency to the invoicing currency.
	ConversionRate param.Field[float64] `json:"conversion_rate"`
	// The configuration for the rate of the price currency to the invoicing currency.
	ConversionRateConfig param.Field[NewPlanTierWithProrationPriceUnionParam] `json:"conversion_rate_config"`
	// An ISO 4217 currency string, or custom pricing unit identifier, in which this
	// price is billed.
	Currency param.Field[string] `json:"currency"`
	// For dimensional price: specifies a price group and dimension values
	DimensionalPriceConfiguration param.Field[NewDimensionalPriceConfigurationParam] `json:"dimensional_price_configuration"`
	// An alias for the price.
	ExternalPriceID param.Field[string] `json:"external_price_id"`
	// If the Price represents a fixed cost, this represents the quantity of units
	// applied.
	FixedPriceQuantity param.Field[float64] `json:"fixed_price_quantity"`
	// The property used to group this price on an invoice
	InvoiceGroupingKey param.Field[string] `json:"invoice_grouping_key"`
	// Within each billing cycle, specifies the cadence at which invoices are produced.
	// If unspecified, a single invoice is produced per billing cycle.
	InvoicingCycleConfiguration param.Field[NewBillingCycleConfigurationParam] `json:"invoicing_cycle_configuration"`
	// User-specified key/value pairs for the resource. Individual keys can be removed
	// by setting the value to `null`, and the entire metadata mapping can be cleared
	// by setting `metadata` to `null`.
	Metadata param.Field[map[string]string] `json:"metadata"`
	// A transient ID that can be used to reference this price when adding adjustments
	// in the same API call.
	ReferenceID param.Field[string] `json:"reference_id"`
}

func (r NewPlanTierWithProrationPriceParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r NewPlanTierWithProrationPriceParam) ImplementsBetaNewPlanVersionParamsAddPricesPriceUnion() {}

func (r NewPlanTierWithProrationPriceParam) ImplementsBetaNewPlanVersionParamsReplacePricesPriceUnion() {
}

func (r NewPlanTierWithProrationPriceParam) ImplementsBetaExternalPlanIDNewPlanVersionParamsAddPricesPriceUnion() {
}

func (r NewPlanTierWithProrationPriceParam) ImplementsBetaExternalPlanIDNewPlanVersionParamsReplacePricesPriceUnion() {
}

func (r NewPlanTierWithProrationPriceParam) ImplementsPlanNewParamsPricesPriceUnion() {}

// The cadence to bill for this price on.
type NewPlanTierWithProrationPriceCadence string

const (
	NewPlanTierWithProrationPriceCadenceAnnual     NewPlanTierWithProrationPriceCadence = "annual"
	NewPlanTierWithProrationPriceCadenceSemiAnnual NewPlanTierWithProrationPriceCadence = "semi_annual"
	NewPlanTierWithProrationPriceCadenceMonthly    NewPlanTierWithProrationPriceCadence = "monthly"
	NewPlanTierWithProrationPriceCadenceQuarterly  NewPlanTierWithProrationPriceCadence = "quarterly"
	NewPlanTierWithProrationPriceCadenceOneTime    NewPlanTierWithProrationPriceCadence = "one_time"
	NewPlanTierWithProrationPriceCadenceCustom     NewPlanTierWithProrationPriceCadence = "custom"
)

func (r NewPlanTierWithProrationPriceCadence) IsKnown() bool {
	switch r {
	case NewPlanTierWithProrationPriceCadenceAnnual, NewPlanTierWithProrationPriceCadenceSemiAnnual, NewPlanTierWithProrationPriceCadenceMonthly, NewPlanTierWithProrationPriceCadenceQuarterly, NewPlanTierWithProrationPriceCadenceOneTime, NewPlanTierWithProrationPriceCadenceCustom:
		return true
	}
	return false
}

type NewPlanTierWithProrationPriceModelType string

const (
	NewPlanTierWithProrationPriceModelTypeTieredWithProration NewPlanTierWithProrationPriceModelType = "tiered_with_proration"
)

func (r NewPlanTierWithProrationPriceModelType) IsKnown() bool {
	switch r {
	case NewPlanTierWithProrationPriceModelTypeTieredWithProration:
		return true
	}
	return false
}

type NewPlanTierWithProrationPriceParam struct {
	ConversionRateType param.Field[NewPlanTierWithProrationPriceConversionRateType] `json:"conversion_rate_type,required"`
	TieredConfig       param.Field[ConversionRateTieredConfigParam]                 `json:"tiered_config"`
	UnitConfig         param.Field[ConversionRateUnitConfigParam]                   `json:"unit_config"`
}

func (r NewPlanTierWithProrationPriceParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r NewPlanTierWithProrationPriceParam) ImplementsNewPlanTierWithProrationPriceUnionParam() {}

// Satisfied by [shared.UnitConversionRateConfigParam],
// [shared.TieredConversionRateConfigParam], [NewPlanTierWithProrationPriceParam].
type NewPlanTierWithProrationPriceUnionParam interface {
	ImplementsNewPlanTierWithProrationPriceUnionParam()
}

type NewPlanTierWithProrationPriceConversionRateType string

const (
	NewPlanTierWithProrationPriceConversionRateTypeUnit   NewPlanTierWithProrationPriceConversionRateType = "unit"
	NewPlanTierWithProrationPriceConversionRateTypeTiered NewPlanTierWithProrationPriceConversionRateType = "tiered"
)

func (r NewPlanTierWithProrationPriceConversionRateType) IsKnown() bool {
	switch r {
	case NewPlanTierWithProrationPriceConversionRateTypeUnit, NewPlanTierWithProrationPriceConversionRateTypeTiered:
		return true
	}
	return false
}

type NewPlanTieredBPSPriceParam struct {
	// The cadence to bill for this price on.
	Cadence param.Field[NewPlanTieredBPSPriceCadence] `json:"cadence,required"`
	// The id of the item the price will be associated with.
	ItemID    param.Field[string]                         `json:"item_id,required"`
	ModelType param.Field[NewPlanTieredBPSPriceModelType] `json:"model_type,required"`
	// The name of the price.
	Name            param.Field[string]               `json:"name,required"`
	TieredBPSConfig param.Field[TieredBPSConfigParam] `json:"tiered_bps_config,required"`
	// The id of the billable metric for the price. Only needed if the price is
	// usage-based.
	BillableMetricID param.Field[string] `json:"billable_metric_id"`
	// If the Price represents a fixed cost, the price will be billed in-advance if
	// this is true, and in-arrears if this is false.
	BilledInAdvance param.Field[bool] `json:"billed_in_advance"`
	// For custom cadence: specifies the duration of the billing period in days or
	// months.
	BillingCycleConfiguration param.Field[NewBillingCycleConfigurationParam] `json:"billing_cycle_configuration"`
	// The per unit conversion rate of the price currency to the invoicing currency.
	ConversionRate param.Field[float64] `json:"conversion_rate"`
	// The configuration for the rate of the price currency to the invoicing currency.
	ConversionRateConfig param.Field[NewPlanTieredBPSPriceUnionParam] `json:"conversion_rate_config"`
	// An ISO 4217 currency string, or custom pricing unit identifier, in which this
	// price is billed.
	Currency param.Field[string] `json:"currency"`
	// For dimensional price: specifies a price group and dimension values
	DimensionalPriceConfiguration param.Field[NewDimensionalPriceConfigurationParam] `json:"dimensional_price_configuration"`
	// An alias for the price.
	ExternalPriceID param.Field[string] `json:"external_price_id"`
	// If the Price represents a fixed cost, this represents the quantity of units
	// applied.
	FixedPriceQuantity param.Field[float64] `json:"fixed_price_quantity"`
	// The property used to group this price on an invoice
	InvoiceGroupingKey param.Field[string] `json:"invoice_grouping_key"`
	// Within each billing cycle, specifies the cadence at which invoices are produced.
	// If unspecified, a single invoice is produced per billing cycle.
	InvoicingCycleConfiguration param.Field[NewBillingCycleConfigurationParam] `json:"invoicing_cycle_configuration"`
	// User-specified key/value pairs for the resource. Individual keys can be removed
	// by setting the value to `null`, and the entire metadata mapping can be cleared
	// by setting `metadata` to `null`.
	Metadata param.Field[map[string]string] `json:"metadata"`
	// A transient ID that can be used to reference this price when adding adjustments
	// in the same API call.
	ReferenceID param.Field[string] `json:"reference_id"`
}

func (r NewPlanTieredBPSPriceParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r NewPlanTieredBPSPriceParam) ImplementsBetaNewPlanVersionParamsAddPricesPriceUnion() {}

func (r NewPlanTieredBPSPriceParam) ImplementsBetaNewPlanVersionParamsReplacePricesPriceUnion() {}

func (r NewPlanTieredBPSPriceParam) ImplementsBetaExternalPlanIDNewPlanVersionParamsAddPricesPriceUnion() {
}

func (r NewPlanTieredBPSPriceParam) ImplementsBetaExternalPlanIDNewPlanVersionParamsReplacePricesPriceUnion() {
}

func (r NewPlanTieredBPSPriceParam) ImplementsPlanNewParamsPricesPriceUnion() {}

// The cadence to bill for this price on.
type NewPlanTieredBPSPriceCadence string

const (
	NewPlanTieredBPSPriceCadenceAnnual     NewPlanTieredBPSPriceCadence = "annual"
	NewPlanTieredBPSPriceCadenceSemiAnnual NewPlanTieredBPSPriceCadence = "semi_annual"
	NewPlanTieredBPSPriceCadenceMonthly    NewPlanTieredBPSPriceCadence = "monthly"
	NewPlanTieredBPSPriceCadenceQuarterly  NewPlanTieredBPSPriceCadence = "quarterly"
	NewPlanTieredBPSPriceCadenceOneTime    NewPlanTieredBPSPriceCadence = "one_time"
	NewPlanTieredBPSPriceCadenceCustom     NewPlanTieredBPSPriceCadence = "custom"
)

func (r NewPlanTieredBPSPriceCadence) IsKnown() bool {
	switch r {
	case NewPlanTieredBPSPriceCadenceAnnual, NewPlanTieredBPSPriceCadenceSemiAnnual, NewPlanTieredBPSPriceCadenceMonthly, NewPlanTieredBPSPriceCadenceQuarterly, NewPlanTieredBPSPriceCadenceOneTime, NewPlanTieredBPSPriceCadenceCustom:
		return true
	}
	return false
}

type NewPlanTieredBPSPriceModelType string

const (
	NewPlanTieredBPSPriceModelTypeTieredBPS NewPlanTieredBPSPriceModelType = "tiered_bps"
)

func (r NewPlanTieredBPSPriceModelType) IsKnown() bool {
	switch r {
	case NewPlanTieredBPSPriceModelTypeTieredBPS:
		return true
	}
	return false
}

type NewPlanTieredBPSPriceParam struct {
	ConversionRateType param.Field[NewPlanTieredBPSPriceConversionRateType] `json:"conversion_rate_type,required"`
	TieredConfig       param.Field[ConversionRateTieredConfigParam]         `json:"tiered_config"`
	UnitConfig         param.Field[ConversionRateUnitConfigParam]           `json:"unit_config"`
}

func (r NewPlanTieredBPSPriceParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r NewPlanTieredBPSPriceParam) ImplementsNewPlanTieredBPSPriceUnionParam() {}

// Satisfied by [shared.UnitConversionRateConfigParam],
// [shared.TieredConversionRateConfigParam], [NewPlanTieredBPSPriceParam].
type NewPlanTieredBPSPriceUnionParam interface {
	ImplementsNewPlanTieredBPSPriceUnionParam()
}

type NewPlanTieredBPSPriceConversionRateType string

const (
	NewPlanTieredBPSPriceConversionRateTypeUnit   NewPlanTieredBPSPriceConversionRateType = "unit"
	NewPlanTieredBPSPriceConversionRateTypeTiered NewPlanTieredBPSPriceConversionRateType = "tiered"
)

func (r NewPlanTieredBPSPriceConversionRateType) IsKnown() bool {
	switch r {
	case NewPlanTieredBPSPriceConversionRateTypeUnit, NewPlanTieredBPSPriceConversionRateTypeTiered:
		return true
	}
	return false
}

type NewPlanTieredPackagePriceParam struct {
	// The cadence to bill for this price on.
	Cadence param.Field[NewPlanTieredPackagePriceCadence] `json:"cadence,required"`
	// The id of the item the price will be associated with.
	ItemID    param.Field[string]                             `json:"item_id,required"`
	ModelType param.Field[NewPlanTieredPackagePriceModelType] `json:"model_type,required"`
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
	BillingCycleConfiguration param.Field[NewBillingCycleConfigurationParam] `json:"billing_cycle_configuration"`
	// The per unit conversion rate of the price currency to the invoicing currency.
	ConversionRate param.Field[float64] `json:"conversion_rate"`
	// The configuration for the rate of the price currency to the invoicing currency.
	ConversionRateConfig param.Field[NewPlanTieredPackagePriceUnionParam] `json:"conversion_rate_config"`
	// An ISO 4217 currency string, or custom pricing unit identifier, in which this
	// price is billed.
	Currency param.Field[string] `json:"currency"`
	// For dimensional price: specifies a price group and dimension values
	DimensionalPriceConfiguration param.Field[NewDimensionalPriceConfigurationParam] `json:"dimensional_price_configuration"`
	// An alias for the price.
	ExternalPriceID param.Field[string] `json:"external_price_id"`
	// If the Price represents a fixed cost, this represents the quantity of units
	// applied.
	FixedPriceQuantity param.Field[float64] `json:"fixed_price_quantity"`
	// The property used to group this price on an invoice
	InvoiceGroupingKey param.Field[string] `json:"invoice_grouping_key"`
	// Within each billing cycle, specifies the cadence at which invoices are produced.
	// If unspecified, a single invoice is produced per billing cycle.
	InvoicingCycleConfiguration param.Field[NewBillingCycleConfigurationParam] `json:"invoicing_cycle_configuration"`
	// User-specified key/value pairs for the resource. Individual keys can be removed
	// by setting the value to `null`, and the entire metadata mapping can be cleared
	// by setting `metadata` to `null`.
	Metadata param.Field[map[string]string] `json:"metadata"`
	// A transient ID that can be used to reference this price when adding adjustments
	// in the same API call.
	ReferenceID param.Field[string] `json:"reference_id"`
}

func (r NewPlanTieredPackagePriceParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r NewPlanTieredPackagePriceParam) ImplementsBetaNewPlanVersionParamsAddPricesPriceUnion() {}

func (r NewPlanTieredPackagePriceParam) ImplementsBetaNewPlanVersionParamsReplacePricesPriceUnion() {}

func (r NewPlanTieredPackagePriceParam) ImplementsBetaExternalPlanIDNewPlanVersionParamsAddPricesPriceUnion() {
}

func (r NewPlanTieredPackagePriceParam) ImplementsBetaExternalPlanIDNewPlanVersionParamsReplacePricesPriceUnion() {
}

func (r NewPlanTieredPackagePriceParam) ImplementsPlanNewParamsPricesPriceUnion() {}

// The cadence to bill for this price on.
type NewPlanTieredPackagePriceCadence string

const (
	NewPlanTieredPackagePriceCadenceAnnual     NewPlanTieredPackagePriceCadence = "annual"
	NewPlanTieredPackagePriceCadenceSemiAnnual NewPlanTieredPackagePriceCadence = "semi_annual"
	NewPlanTieredPackagePriceCadenceMonthly    NewPlanTieredPackagePriceCadence = "monthly"
	NewPlanTieredPackagePriceCadenceQuarterly  NewPlanTieredPackagePriceCadence = "quarterly"
	NewPlanTieredPackagePriceCadenceOneTime    NewPlanTieredPackagePriceCadence = "one_time"
	NewPlanTieredPackagePriceCadenceCustom     NewPlanTieredPackagePriceCadence = "custom"
)

func (r NewPlanTieredPackagePriceCadence) IsKnown() bool {
	switch r {
	case NewPlanTieredPackagePriceCadenceAnnual, NewPlanTieredPackagePriceCadenceSemiAnnual, NewPlanTieredPackagePriceCadenceMonthly, NewPlanTieredPackagePriceCadenceQuarterly, NewPlanTieredPackagePriceCadenceOneTime, NewPlanTieredPackagePriceCadenceCustom:
		return true
	}
	return false
}

type NewPlanTieredPackagePriceModelType string

const (
	NewPlanTieredPackagePriceModelTypeTieredPackage NewPlanTieredPackagePriceModelType = "tiered_package"
)

func (r NewPlanTieredPackagePriceModelType) IsKnown() bool {
	switch r {
	case NewPlanTieredPackagePriceModelTypeTieredPackage:
		return true
	}
	return false
}

type NewPlanTieredPackagePriceParam struct {
	ConversionRateType param.Field[NewPlanTieredPackagePriceConversionRateType] `json:"conversion_rate_type,required"`
	TieredConfig       param.Field[ConversionRateTieredConfigParam]             `json:"tiered_config"`
	UnitConfig         param.Field[ConversionRateUnitConfigParam]               `json:"unit_config"`
}

func (r NewPlanTieredPackagePriceParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r NewPlanTieredPackagePriceParam) ImplementsNewPlanTieredPackagePriceUnionParam() {}

// Satisfied by [shared.UnitConversionRateConfigParam],
// [shared.TieredConversionRateConfigParam], [NewPlanTieredPackagePriceParam].
type NewPlanTieredPackagePriceUnionParam interface {
	ImplementsNewPlanTieredPackagePriceUnionParam()
}

type NewPlanTieredPackagePriceConversionRateType string

const (
	NewPlanTieredPackagePriceConversionRateTypeUnit   NewPlanTieredPackagePriceConversionRateType = "unit"
	NewPlanTieredPackagePriceConversionRateTypeTiered NewPlanTieredPackagePriceConversionRateType = "tiered"
)

func (r NewPlanTieredPackagePriceConversionRateType) IsKnown() bool {
	switch r {
	case NewPlanTieredPackagePriceConversionRateTypeUnit, NewPlanTieredPackagePriceConversionRateTypeTiered:
		return true
	}
	return false
}

type NewPlanTieredPackageWithMinimumPriceParam struct {
	// The cadence to bill for this price on.
	Cadence param.Field[NewPlanTieredPackageWithMinimumPriceCadence] `json:"cadence,required"`
	// The id of the item the price will be associated with.
	ItemID    param.Field[string]                                        `json:"item_id,required"`
	ModelType param.Field[NewPlanTieredPackageWithMinimumPriceModelType] `json:"model_type,required"`
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
	BillingCycleConfiguration param.Field[NewBillingCycleConfigurationParam] `json:"billing_cycle_configuration"`
	// The per unit conversion rate of the price currency to the invoicing currency.
	ConversionRate param.Field[float64] `json:"conversion_rate"`
	// The configuration for the rate of the price currency to the invoicing currency.
	ConversionRateConfig param.Field[NewPlanTieredPackageWithMinimumPriceUnionParam] `json:"conversion_rate_config"`
	// An ISO 4217 currency string, or custom pricing unit identifier, in which this
	// price is billed.
	Currency param.Field[string] `json:"currency"`
	// For dimensional price: specifies a price group and dimension values
	DimensionalPriceConfiguration param.Field[NewDimensionalPriceConfigurationParam] `json:"dimensional_price_configuration"`
	// An alias for the price.
	ExternalPriceID param.Field[string] `json:"external_price_id"`
	// If the Price represents a fixed cost, this represents the quantity of units
	// applied.
	FixedPriceQuantity param.Field[float64] `json:"fixed_price_quantity"`
	// The property used to group this price on an invoice
	InvoiceGroupingKey param.Field[string] `json:"invoice_grouping_key"`
	// Within each billing cycle, specifies the cadence at which invoices are produced.
	// If unspecified, a single invoice is produced per billing cycle.
	InvoicingCycleConfiguration param.Field[NewBillingCycleConfigurationParam] `json:"invoicing_cycle_configuration"`
	// User-specified key/value pairs for the resource. Individual keys can be removed
	// by setting the value to `null`, and the entire metadata mapping can be cleared
	// by setting `metadata` to `null`.
	Metadata param.Field[map[string]string] `json:"metadata"`
	// A transient ID that can be used to reference this price when adding adjustments
	// in the same API call.
	ReferenceID param.Field[string] `json:"reference_id"`
}

func (r NewPlanTieredPackageWithMinimumPriceParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r NewPlanTieredPackageWithMinimumPriceParam) ImplementsBetaNewPlanVersionParamsAddPricesPriceUnion() {
}

func (r NewPlanTieredPackageWithMinimumPriceParam) ImplementsBetaNewPlanVersionParamsReplacePricesPriceUnion() {
}

func (r NewPlanTieredPackageWithMinimumPriceParam) ImplementsBetaExternalPlanIDNewPlanVersionParamsAddPricesPriceUnion() {
}

func (r NewPlanTieredPackageWithMinimumPriceParam) ImplementsBetaExternalPlanIDNewPlanVersionParamsReplacePricesPriceUnion() {
}

func (r NewPlanTieredPackageWithMinimumPriceParam) ImplementsPlanNewParamsPricesPriceUnion() {}

// The cadence to bill for this price on.
type NewPlanTieredPackageWithMinimumPriceCadence string

const (
	NewPlanTieredPackageWithMinimumPriceCadenceAnnual     NewPlanTieredPackageWithMinimumPriceCadence = "annual"
	NewPlanTieredPackageWithMinimumPriceCadenceSemiAnnual NewPlanTieredPackageWithMinimumPriceCadence = "semi_annual"
	NewPlanTieredPackageWithMinimumPriceCadenceMonthly    NewPlanTieredPackageWithMinimumPriceCadence = "monthly"
	NewPlanTieredPackageWithMinimumPriceCadenceQuarterly  NewPlanTieredPackageWithMinimumPriceCadence = "quarterly"
	NewPlanTieredPackageWithMinimumPriceCadenceOneTime    NewPlanTieredPackageWithMinimumPriceCadence = "one_time"
	NewPlanTieredPackageWithMinimumPriceCadenceCustom     NewPlanTieredPackageWithMinimumPriceCadence = "custom"
)

func (r NewPlanTieredPackageWithMinimumPriceCadence) IsKnown() bool {
	switch r {
	case NewPlanTieredPackageWithMinimumPriceCadenceAnnual, NewPlanTieredPackageWithMinimumPriceCadenceSemiAnnual, NewPlanTieredPackageWithMinimumPriceCadenceMonthly, NewPlanTieredPackageWithMinimumPriceCadenceQuarterly, NewPlanTieredPackageWithMinimumPriceCadenceOneTime, NewPlanTieredPackageWithMinimumPriceCadenceCustom:
		return true
	}
	return false
}

type NewPlanTieredPackageWithMinimumPriceModelType string

const (
	NewPlanTieredPackageWithMinimumPriceModelTypeTieredPackageWithMinimum NewPlanTieredPackageWithMinimumPriceModelType = "tiered_package_with_minimum"
)

func (r NewPlanTieredPackageWithMinimumPriceModelType) IsKnown() bool {
	switch r {
	case NewPlanTieredPackageWithMinimumPriceModelTypeTieredPackageWithMinimum:
		return true
	}
	return false
}

type NewPlanTieredPackageWithMinimumPriceParam struct {
	ConversionRateType param.Field[NewPlanTieredPackageWithMinimumPriceConversionRateType] `json:"conversion_rate_type,required"`
	TieredConfig       param.Field[ConversionRateTieredConfigParam]                        `json:"tiered_config"`
	UnitConfig         param.Field[ConversionRateUnitConfigParam]                          `json:"unit_config"`
}

func (r NewPlanTieredPackageWithMinimumPriceParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r NewPlanTieredPackageWithMinimumPriceParam) ImplementsNewPlanTieredPackageWithMinimumPriceUnionParam() {
}

// Satisfied by [shared.UnitConversionRateConfigParam],
// [shared.TieredConversionRateConfigParam],
// [NewPlanTieredPackageWithMinimumPriceParam].
type NewPlanTieredPackageWithMinimumPriceUnionParam interface {
	ImplementsNewPlanTieredPackageWithMinimumPriceUnionParam()
}

type NewPlanTieredPackageWithMinimumPriceConversionRateType string

const (
	NewPlanTieredPackageWithMinimumPriceConversionRateTypeUnit   NewPlanTieredPackageWithMinimumPriceConversionRateType = "unit"
	NewPlanTieredPackageWithMinimumPriceConversionRateTypeTiered NewPlanTieredPackageWithMinimumPriceConversionRateType = "tiered"
)

func (r NewPlanTieredPackageWithMinimumPriceConversionRateType) IsKnown() bool {
	switch r {
	case NewPlanTieredPackageWithMinimumPriceConversionRateTypeUnit, NewPlanTieredPackageWithMinimumPriceConversionRateTypeTiered:
		return true
	}
	return false
}

type NewPlanTieredPriceParam struct {
	// The cadence to bill for this price on.
	Cadence param.Field[NewPlanTieredPriceCadence] `json:"cadence,required"`
	// The id of the item the price will be associated with.
	ItemID    param.Field[string]                      `json:"item_id,required"`
	ModelType param.Field[NewPlanTieredPriceModelType] `json:"model_type,required"`
	// The name of the price.
	Name         param.Field[string]            `json:"name,required"`
	TieredConfig param.Field[TieredConfigParam] `json:"tiered_config,required"`
	// The id of the billable metric for the price. Only needed if the price is
	// usage-based.
	BillableMetricID param.Field[string] `json:"billable_metric_id"`
	// If the Price represents a fixed cost, the price will be billed in-advance if
	// this is true, and in-arrears if this is false.
	BilledInAdvance param.Field[bool] `json:"billed_in_advance"`
	// For custom cadence: specifies the duration of the billing period in days or
	// months.
	BillingCycleConfiguration param.Field[NewBillingCycleConfigurationParam] `json:"billing_cycle_configuration"`
	// The per unit conversion rate of the price currency to the invoicing currency.
	ConversionRate param.Field[float64] `json:"conversion_rate"`
	// The configuration for the rate of the price currency to the invoicing currency.
	ConversionRateConfig param.Field[NewPlanTieredPriceUnionParam] `json:"conversion_rate_config"`
	// An ISO 4217 currency string, or custom pricing unit identifier, in which this
	// price is billed.
	Currency param.Field[string] `json:"currency"`
	// For dimensional price: specifies a price group and dimension values
	DimensionalPriceConfiguration param.Field[NewDimensionalPriceConfigurationParam] `json:"dimensional_price_configuration"`
	// An alias for the price.
	ExternalPriceID param.Field[string] `json:"external_price_id"`
	// If the Price represents a fixed cost, this represents the quantity of units
	// applied.
	FixedPriceQuantity param.Field[float64] `json:"fixed_price_quantity"`
	// The property used to group this price on an invoice
	InvoiceGroupingKey param.Field[string] `json:"invoice_grouping_key"`
	// Within each billing cycle, specifies the cadence at which invoices are produced.
	// If unspecified, a single invoice is produced per billing cycle.
	InvoicingCycleConfiguration param.Field[NewBillingCycleConfigurationParam] `json:"invoicing_cycle_configuration"`
	// User-specified key/value pairs for the resource. Individual keys can be removed
	// by setting the value to `null`, and the entire metadata mapping can be cleared
	// by setting `metadata` to `null`.
	Metadata param.Field[map[string]string] `json:"metadata"`
	// A transient ID that can be used to reference this price when adding adjustments
	// in the same API call.
	ReferenceID param.Field[string] `json:"reference_id"`
}

func (r NewPlanTieredPriceParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r NewPlanTieredPriceParam) ImplementsBetaNewPlanVersionParamsAddPricesPriceUnion() {}

func (r NewPlanTieredPriceParam) ImplementsBetaNewPlanVersionParamsReplacePricesPriceUnion() {}

func (r NewPlanTieredPriceParam) ImplementsBetaExternalPlanIDNewPlanVersionParamsAddPricesPriceUnion() {
}

func (r NewPlanTieredPriceParam) ImplementsBetaExternalPlanIDNewPlanVersionParamsReplacePricesPriceUnion() {
}

func (r NewPlanTieredPriceParam) ImplementsPlanNewParamsPricesPriceUnion() {}

// The cadence to bill for this price on.
type NewPlanTieredPriceCadence string

const (
	NewPlanTieredPriceCadenceAnnual     NewPlanTieredPriceCadence = "annual"
	NewPlanTieredPriceCadenceSemiAnnual NewPlanTieredPriceCadence = "semi_annual"
	NewPlanTieredPriceCadenceMonthly    NewPlanTieredPriceCadence = "monthly"
	NewPlanTieredPriceCadenceQuarterly  NewPlanTieredPriceCadence = "quarterly"
	NewPlanTieredPriceCadenceOneTime    NewPlanTieredPriceCadence = "one_time"
	NewPlanTieredPriceCadenceCustom     NewPlanTieredPriceCadence = "custom"
)

func (r NewPlanTieredPriceCadence) IsKnown() bool {
	switch r {
	case NewPlanTieredPriceCadenceAnnual, NewPlanTieredPriceCadenceSemiAnnual, NewPlanTieredPriceCadenceMonthly, NewPlanTieredPriceCadenceQuarterly, NewPlanTieredPriceCadenceOneTime, NewPlanTieredPriceCadenceCustom:
		return true
	}
	return false
}

type NewPlanTieredPriceModelType string

const (
	NewPlanTieredPriceModelTypeTiered NewPlanTieredPriceModelType = "tiered"
)

func (r NewPlanTieredPriceModelType) IsKnown() bool {
	switch r {
	case NewPlanTieredPriceModelTypeTiered:
		return true
	}
	return false
}

type NewPlanTieredPriceParam struct {
	ConversionRateType param.Field[NewPlanTieredPriceConversionRateType] `json:"conversion_rate_type,required"`
	TieredConfig       param.Field[ConversionRateTieredConfigParam]      `json:"tiered_config"`
	UnitConfig         param.Field[ConversionRateUnitConfigParam]        `json:"unit_config"`
}

func (r NewPlanTieredPriceParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r NewPlanTieredPriceParam) ImplementsNewPlanTieredPriceUnionParam() {}

// Satisfied by [shared.UnitConversionRateConfigParam],
// [shared.TieredConversionRateConfigParam], [NewPlanTieredPriceParam].
type NewPlanTieredPriceUnionParam interface {
	ImplementsNewPlanTieredPriceUnionParam()
}

type NewPlanTieredPriceConversionRateType string

const (
	NewPlanTieredPriceConversionRateTypeUnit   NewPlanTieredPriceConversionRateType = "unit"
	NewPlanTieredPriceConversionRateTypeTiered NewPlanTieredPriceConversionRateType = "tiered"
)

func (r NewPlanTieredPriceConversionRateType) IsKnown() bool {
	switch r {
	case NewPlanTieredPriceConversionRateTypeUnit, NewPlanTieredPriceConversionRateTypeTiered:
		return true
	}
	return false
}

type NewPlanTieredWithMinimumPriceParam struct {
	// The cadence to bill for this price on.
	Cadence param.Field[NewPlanTieredWithMinimumPriceCadence] `json:"cadence,required"`
	// The id of the item the price will be associated with.
	ItemID    param.Field[string]                                 `json:"item_id,required"`
	ModelType param.Field[NewPlanTieredWithMinimumPriceModelType] `json:"model_type,required"`
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
	BillingCycleConfiguration param.Field[NewBillingCycleConfigurationParam] `json:"billing_cycle_configuration"`
	// The per unit conversion rate of the price currency to the invoicing currency.
	ConversionRate param.Field[float64] `json:"conversion_rate"`
	// The configuration for the rate of the price currency to the invoicing currency.
	ConversionRateConfig param.Field[NewPlanTieredWithMinimumPriceUnionParam] `json:"conversion_rate_config"`
	// An ISO 4217 currency string, or custom pricing unit identifier, in which this
	// price is billed.
	Currency param.Field[string] `json:"currency"`
	// For dimensional price: specifies a price group and dimension values
	DimensionalPriceConfiguration param.Field[NewDimensionalPriceConfigurationParam] `json:"dimensional_price_configuration"`
	// An alias for the price.
	ExternalPriceID param.Field[string] `json:"external_price_id"`
	// If the Price represents a fixed cost, this represents the quantity of units
	// applied.
	FixedPriceQuantity param.Field[float64] `json:"fixed_price_quantity"`
	// The property used to group this price on an invoice
	InvoiceGroupingKey param.Field[string] `json:"invoice_grouping_key"`
	// Within each billing cycle, specifies the cadence at which invoices are produced.
	// If unspecified, a single invoice is produced per billing cycle.
	InvoicingCycleConfiguration param.Field[NewBillingCycleConfigurationParam] `json:"invoicing_cycle_configuration"`
	// User-specified key/value pairs for the resource. Individual keys can be removed
	// by setting the value to `null`, and the entire metadata mapping can be cleared
	// by setting `metadata` to `null`.
	Metadata param.Field[map[string]string] `json:"metadata"`
	// A transient ID that can be used to reference this price when adding adjustments
	// in the same API call.
	ReferenceID param.Field[string] `json:"reference_id"`
}

func (r NewPlanTieredWithMinimumPriceParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r NewPlanTieredWithMinimumPriceParam) ImplementsBetaNewPlanVersionParamsAddPricesPriceUnion() {}

func (r NewPlanTieredWithMinimumPriceParam) ImplementsBetaNewPlanVersionParamsReplacePricesPriceUnion() {
}

func (r NewPlanTieredWithMinimumPriceParam) ImplementsBetaExternalPlanIDNewPlanVersionParamsAddPricesPriceUnion() {
}

func (r NewPlanTieredWithMinimumPriceParam) ImplementsBetaExternalPlanIDNewPlanVersionParamsReplacePricesPriceUnion() {
}

func (r NewPlanTieredWithMinimumPriceParam) ImplementsPlanNewParamsPricesPriceUnion() {}

// The cadence to bill for this price on.
type NewPlanTieredWithMinimumPriceCadence string

const (
	NewPlanTieredWithMinimumPriceCadenceAnnual     NewPlanTieredWithMinimumPriceCadence = "annual"
	NewPlanTieredWithMinimumPriceCadenceSemiAnnual NewPlanTieredWithMinimumPriceCadence = "semi_annual"
	NewPlanTieredWithMinimumPriceCadenceMonthly    NewPlanTieredWithMinimumPriceCadence = "monthly"
	NewPlanTieredWithMinimumPriceCadenceQuarterly  NewPlanTieredWithMinimumPriceCadence = "quarterly"
	NewPlanTieredWithMinimumPriceCadenceOneTime    NewPlanTieredWithMinimumPriceCadence = "one_time"
	NewPlanTieredWithMinimumPriceCadenceCustom     NewPlanTieredWithMinimumPriceCadence = "custom"
)

func (r NewPlanTieredWithMinimumPriceCadence) IsKnown() bool {
	switch r {
	case NewPlanTieredWithMinimumPriceCadenceAnnual, NewPlanTieredWithMinimumPriceCadenceSemiAnnual, NewPlanTieredWithMinimumPriceCadenceMonthly, NewPlanTieredWithMinimumPriceCadenceQuarterly, NewPlanTieredWithMinimumPriceCadenceOneTime, NewPlanTieredWithMinimumPriceCadenceCustom:
		return true
	}
	return false
}

type NewPlanTieredWithMinimumPriceModelType string

const (
	NewPlanTieredWithMinimumPriceModelTypeTieredWithMinimum NewPlanTieredWithMinimumPriceModelType = "tiered_with_minimum"
)

func (r NewPlanTieredWithMinimumPriceModelType) IsKnown() bool {
	switch r {
	case NewPlanTieredWithMinimumPriceModelTypeTieredWithMinimum:
		return true
	}
	return false
}

type NewPlanTieredWithMinimumPriceParam struct {
	ConversionRateType param.Field[NewPlanTieredWithMinimumPriceConversionRateType] `json:"conversion_rate_type,required"`
	TieredConfig       param.Field[ConversionRateTieredConfigParam]                 `json:"tiered_config"`
	UnitConfig         param.Field[ConversionRateUnitConfigParam]                   `json:"unit_config"`
}

func (r NewPlanTieredWithMinimumPriceParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r NewPlanTieredWithMinimumPriceParam) ImplementsNewPlanTieredWithMinimumPriceUnionParam() {}

// Satisfied by [shared.UnitConversionRateConfigParam],
// [shared.TieredConversionRateConfigParam], [NewPlanTieredWithMinimumPriceParam].
type NewPlanTieredWithMinimumPriceUnionParam interface {
	ImplementsNewPlanTieredWithMinimumPriceUnionParam()
}

type NewPlanTieredWithMinimumPriceConversionRateType string

const (
	NewPlanTieredWithMinimumPriceConversionRateTypeUnit   NewPlanTieredWithMinimumPriceConversionRateType = "unit"
	NewPlanTieredWithMinimumPriceConversionRateTypeTiered NewPlanTieredWithMinimumPriceConversionRateType = "tiered"
)

func (r NewPlanTieredWithMinimumPriceConversionRateType) IsKnown() bool {
	switch r {
	case NewPlanTieredWithMinimumPriceConversionRateTypeUnit, NewPlanTieredWithMinimumPriceConversionRateTypeTiered:
		return true
	}
	return false
}

type NewPlanUnitPriceParam struct {
	// The cadence to bill for this price on.
	Cadence param.Field[NewPlanUnitPriceCadence] `json:"cadence,required"`
	// The id of the item the price will be associated with.
	ItemID    param.Field[string]                    `json:"item_id,required"`
	ModelType param.Field[NewPlanUnitPriceModelType] `json:"model_type,required"`
	// The name of the price.
	Name       param.Field[string]          `json:"name,required"`
	UnitConfig param.Field[UnitConfigParam] `json:"unit_config,required"`
	// The id of the billable metric for the price. Only needed if the price is
	// usage-based.
	BillableMetricID param.Field[string] `json:"billable_metric_id"`
	// If the Price represents a fixed cost, the price will be billed in-advance if
	// this is true, and in-arrears if this is false.
	BilledInAdvance param.Field[bool] `json:"billed_in_advance"`
	// For custom cadence: specifies the duration of the billing period in days or
	// months.
	BillingCycleConfiguration param.Field[NewBillingCycleConfigurationParam] `json:"billing_cycle_configuration"`
	// The per unit conversion rate of the price currency to the invoicing currency.
	ConversionRate param.Field[float64] `json:"conversion_rate"`
	// The configuration for the rate of the price currency to the invoicing currency.
	ConversionRateConfig param.Field[NewPlanUnitPriceUnionParam] `json:"conversion_rate_config"`
	// An ISO 4217 currency string, or custom pricing unit identifier, in which this
	// price is billed.
	Currency param.Field[string] `json:"currency"`
	// For dimensional price: specifies a price group and dimension values
	DimensionalPriceConfiguration param.Field[NewDimensionalPriceConfigurationParam] `json:"dimensional_price_configuration"`
	// An alias for the price.
	ExternalPriceID param.Field[string] `json:"external_price_id"`
	// If the Price represents a fixed cost, this represents the quantity of units
	// applied.
	FixedPriceQuantity param.Field[float64] `json:"fixed_price_quantity"`
	// The property used to group this price on an invoice
	InvoiceGroupingKey param.Field[string] `json:"invoice_grouping_key"`
	// Within each billing cycle, specifies the cadence at which invoices are produced.
	// If unspecified, a single invoice is produced per billing cycle.
	InvoicingCycleConfiguration param.Field[NewBillingCycleConfigurationParam] `json:"invoicing_cycle_configuration"`
	// User-specified key/value pairs for the resource. Individual keys can be removed
	// by setting the value to `null`, and the entire metadata mapping can be cleared
	// by setting `metadata` to `null`.
	Metadata param.Field[map[string]string] `json:"metadata"`
	// A transient ID that can be used to reference this price when adding adjustments
	// in the same API call.
	ReferenceID param.Field[string] `json:"reference_id"`
}

func (r NewPlanUnitPriceParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r NewPlanUnitPriceParam) ImplementsBetaNewPlanVersionParamsAddPricesPriceUnion() {}

func (r NewPlanUnitPriceParam) ImplementsBetaNewPlanVersionParamsReplacePricesPriceUnion() {}

func (r NewPlanUnitPriceParam) ImplementsBetaExternalPlanIDNewPlanVersionParamsAddPricesPriceUnion() {
}

func (r NewPlanUnitPriceParam) ImplementsBetaExternalPlanIDNewPlanVersionParamsReplacePricesPriceUnion() {
}

func (r NewPlanUnitPriceParam) ImplementsPlanNewParamsPricesPriceUnion() {}

// The cadence to bill for this price on.
type NewPlanUnitPriceCadence string

const (
	NewPlanUnitPriceCadenceAnnual     NewPlanUnitPriceCadence = "annual"
	NewPlanUnitPriceCadenceSemiAnnual NewPlanUnitPriceCadence = "semi_annual"
	NewPlanUnitPriceCadenceMonthly    NewPlanUnitPriceCadence = "monthly"
	NewPlanUnitPriceCadenceQuarterly  NewPlanUnitPriceCadence = "quarterly"
	NewPlanUnitPriceCadenceOneTime    NewPlanUnitPriceCadence = "one_time"
	NewPlanUnitPriceCadenceCustom     NewPlanUnitPriceCadence = "custom"
)

func (r NewPlanUnitPriceCadence) IsKnown() bool {
	switch r {
	case NewPlanUnitPriceCadenceAnnual, NewPlanUnitPriceCadenceSemiAnnual, NewPlanUnitPriceCadenceMonthly, NewPlanUnitPriceCadenceQuarterly, NewPlanUnitPriceCadenceOneTime, NewPlanUnitPriceCadenceCustom:
		return true
	}
	return false
}

type NewPlanUnitPriceModelType string

const (
	NewPlanUnitPriceModelTypeUnit NewPlanUnitPriceModelType = "unit"
)

func (r NewPlanUnitPriceModelType) IsKnown() bool {
	switch r {
	case NewPlanUnitPriceModelTypeUnit:
		return true
	}
	return false
}

type NewPlanUnitPriceParam struct {
	ConversionRateType param.Field[NewPlanUnitPriceConversionRateType] `json:"conversion_rate_type,required"`
	TieredConfig       param.Field[ConversionRateTieredConfigParam]    `json:"tiered_config"`
	UnitConfig         param.Field[ConversionRateUnitConfigParam]      `json:"unit_config"`
}

func (r NewPlanUnitPriceParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r NewPlanUnitPriceParam) ImplementsNewPlanUnitPriceUnionParam() {}

// Satisfied by [shared.UnitConversionRateConfigParam],
// [shared.TieredConversionRateConfigParam], [NewPlanUnitPriceParam].
type NewPlanUnitPriceUnionParam interface {
	ImplementsNewPlanUnitPriceUnionParam()
}

type NewPlanUnitPriceConversionRateType string

const (
	NewPlanUnitPriceConversionRateTypeUnit   NewPlanUnitPriceConversionRateType = "unit"
	NewPlanUnitPriceConversionRateTypeTiered NewPlanUnitPriceConversionRateType = "tiered"
)

func (r NewPlanUnitPriceConversionRateType) IsKnown() bool {
	switch r {
	case NewPlanUnitPriceConversionRateTypeUnit, NewPlanUnitPriceConversionRateTypeTiered:
		return true
	}
	return false
}

type NewPlanUnitWithPercentPriceParam struct {
	// The cadence to bill for this price on.
	Cadence param.Field[NewPlanUnitWithPercentPriceCadence] `json:"cadence,required"`
	// The id of the item the price will be associated with.
	ItemID    param.Field[string]                               `json:"item_id,required"`
	ModelType param.Field[NewPlanUnitWithPercentPriceModelType] `json:"model_type,required"`
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
	BillingCycleConfiguration param.Field[NewBillingCycleConfigurationParam] `json:"billing_cycle_configuration"`
	// The per unit conversion rate of the price currency to the invoicing currency.
	ConversionRate param.Field[float64] `json:"conversion_rate"`
	// The configuration for the rate of the price currency to the invoicing currency.
	ConversionRateConfig param.Field[NewPlanUnitWithPercentPriceUnionParam] `json:"conversion_rate_config"`
	// An ISO 4217 currency string, or custom pricing unit identifier, in which this
	// price is billed.
	Currency param.Field[string] `json:"currency"`
	// For dimensional price: specifies a price group and dimension values
	DimensionalPriceConfiguration param.Field[NewDimensionalPriceConfigurationParam] `json:"dimensional_price_configuration"`
	// An alias for the price.
	ExternalPriceID param.Field[string] `json:"external_price_id"`
	// If the Price represents a fixed cost, this represents the quantity of units
	// applied.
	FixedPriceQuantity param.Field[float64] `json:"fixed_price_quantity"`
	// The property used to group this price on an invoice
	InvoiceGroupingKey param.Field[string] `json:"invoice_grouping_key"`
	// Within each billing cycle, specifies the cadence at which invoices are produced.
	// If unspecified, a single invoice is produced per billing cycle.
	InvoicingCycleConfiguration param.Field[NewBillingCycleConfigurationParam] `json:"invoicing_cycle_configuration"`
	// User-specified key/value pairs for the resource. Individual keys can be removed
	// by setting the value to `null`, and the entire metadata mapping can be cleared
	// by setting `metadata` to `null`.
	Metadata param.Field[map[string]string] `json:"metadata"`
	// A transient ID that can be used to reference this price when adding adjustments
	// in the same API call.
	ReferenceID param.Field[string] `json:"reference_id"`
}

func (r NewPlanUnitWithPercentPriceParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r NewPlanUnitWithPercentPriceParam) ImplementsBetaNewPlanVersionParamsAddPricesPriceUnion() {}

func (r NewPlanUnitWithPercentPriceParam) ImplementsBetaNewPlanVersionParamsReplacePricesPriceUnion() {
}

func (r NewPlanUnitWithPercentPriceParam) ImplementsBetaExternalPlanIDNewPlanVersionParamsAddPricesPriceUnion() {
}

func (r NewPlanUnitWithPercentPriceParam) ImplementsBetaExternalPlanIDNewPlanVersionParamsReplacePricesPriceUnion() {
}

func (r NewPlanUnitWithPercentPriceParam) ImplementsPlanNewParamsPricesPriceUnion() {}

// The cadence to bill for this price on.
type NewPlanUnitWithPercentPriceCadence string

const (
	NewPlanUnitWithPercentPriceCadenceAnnual     NewPlanUnitWithPercentPriceCadence = "annual"
	NewPlanUnitWithPercentPriceCadenceSemiAnnual NewPlanUnitWithPercentPriceCadence = "semi_annual"
	NewPlanUnitWithPercentPriceCadenceMonthly    NewPlanUnitWithPercentPriceCadence = "monthly"
	NewPlanUnitWithPercentPriceCadenceQuarterly  NewPlanUnitWithPercentPriceCadence = "quarterly"
	NewPlanUnitWithPercentPriceCadenceOneTime    NewPlanUnitWithPercentPriceCadence = "one_time"
	NewPlanUnitWithPercentPriceCadenceCustom     NewPlanUnitWithPercentPriceCadence = "custom"
)

func (r NewPlanUnitWithPercentPriceCadence) IsKnown() bool {
	switch r {
	case NewPlanUnitWithPercentPriceCadenceAnnual, NewPlanUnitWithPercentPriceCadenceSemiAnnual, NewPlanUnitWithPercentPriceCadenceMonthly, NewPlanUnitWithPercentPriceCadenceQuarterly, NewPlanUnitWithPercentPriceCadenceOneTime, NewPlanUnitWithPercentPriceCadenceCustom:
		return true
	}
	return false
}

type NewPlanUnitWithPercentPriceModelType string

const (
	NewPlanUnitWithPercentPriceModelTypeUnitWithPercent NewPlanUnitWithPercentPriceModelType = "unit_with_percent"
)

func (r NewPlanUnitWithPercentPriceModelType) IsKnown() bool {
	switch r {
	case NewPlanUnitWithPercentPriceModelTypeUnitWithPercent:
		return true
	}
	return false
}

type NewPlanUnitWithPercentPriceParam struct {
	ConversionRateType param.Field[NewPlanUnitWithPercentPriceConversionRateType] `json:"conversion_rate_type,required"`
	TieredConfig       param.Field[ConversionRateTieredConfigParam]               `json:"tiered_config"`
	UnitConfig         param.Field[ConversionRateUnitConfigParam]                 `json:"unit_config"`
}

func (r NewPlanUnitWithPercentPriceParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r NewPlanUnitWithPercentPriceParam) ImplementsNewPlanUnitWithPercentPriceUnionParam() {}

// Satisfied by [shared.UnitConversionRateConfigParam],
// [shared.TieredConversionRateConfigParam], [NewPlanUnitWithPercentPriceParam].
type NewPlanUnitWithPercentPriceUnionParam interface {
	ImplementsNewPlanUnitWithPercentPriceUnionParam()
}

type NewPlanUnitWithPercentPriceConversionRateType string

const (
	NewPlanUnitWithPercentPriceConversionRateTypeUnit   NewPlanUnitWithPercentPriceConversionRateType = "unit"
	NewPlanUnitWithPercentPriceConversionRateTypeTiered NewPlanUnitWithPercentPriceConversionRateType = "tiered"
)

func (r NewPlanUnitWithPercentPriceConversionRateType) IsKnown() bool {
	switch r {
	case NewPlanUnitWithPercentPriceConversionRateTypeUnit, NewPlanUnitWithPercentPriceConversionRateTypeTiered:
		return true
	}
	return false
}

type NewPlanUnitWithProrationPriceParam struct {
	// The cadence to bill for this price on.
	Cadence param.Field[NewPlanUnitWithProrationPriceCadence] `json:"cadence,required"`
	// The id of the item the price will be associated with.
	ItemID    param.Field[string]                                 `json:"item_id,required"`
	ModelType param.Field[NewPlanUnitWithProrationPriceModelType] `json:"model_type,required"`
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
	BillingCycleConfiguration param.Field[NewBillingCycleConfigurationParam] `json:"billing_cycle_configuration"`
	// The per unit conversion rate of the price currency to the invoicing currency.
	ConversionRate param.Field[float64] `json:"conversion_rate"`
	// The configuration for the rate of the price currency to the invoicing currency.
	ConversionRateConfig param.Field[NewPlanUnitWithProrationPriceUnionParam] `json:"conversion_rate_config"`
	// An ISO 4217 currency string, or custom pricing unit identifier, in which this
	// price is billed.
	Currency param.Field[string] `json:"currency"`
	// For dimensional price: specifies a price group and dimension values
	DimensionalPriceConfiguration param.Field[NewDimensionalPriceConfigurationParam] `json:"dimensional_price_configuration"`
	// An alias for the price.
	ExternalPriceID param.Field[string] `json:"external_price_id"`
	// If the Price represents a fixed cost, this represents the quantity of units
	// applied.
	FixedPriceQuantity param.Field[float64] `json:"fixed_price_quantity"`
	// The property used to group this price on an invoice
	InvoiceGroupingKey param.Field[string] `json:"invoice_grouping_key"`
	// Within each billing cycle, specifies the cadence at which invoices are produced.
	// If unspecified, a single invoice is produced per billing cycle.
	InvoicingCycleConfiguration param.Field[NewBillingCycleConfigurationParam] `json:"invoicing_cycle_configuration"`
	// User-specified key/value pairs for the resource. Individual keys can be removed
	// by setting the value to `null`, and the entire metadata mapping can be cleared
	// by setting `metadata` to `null`.
	Metadata param.Field[map[string]string] `json:"metadata"`
	// A transient ID that can be used to reference this price when adding adjustments
	// in the same API call.
	ReferenceID param.Field[string] `json:"reference_id"`
}

func (r NewPlanUnitWithProrationPriceParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r NewPlanUnitWithProrationPriceParam) ImplementsBetaNewPlanVersionParamsAddPricesPriceUnion() {}

func (r NewPlanUnitWithProrationPriceParam) ImplementsBetaNewPlanVersionParamsReplacePricesPriceUnion() {
}

func (r NewPlanUnitWithProrationPriceParam) ImplementsBetaExternalPlanIDNewPlanVersionParamsAddPricesPriceUnion() {
}

func (r NewPlanUnitWithProrationPriceParam) ImplementsBetaExternalPlanIDNewPlanVersionParamsReplacePricesPriceUnion() {
}

func (r NewPlanUnitWithProrationPriceParam) ImplementsPlanNewParamsPricesPriceUnion() {}

// The cadence to bill for this price on.
type NewPlanUnitWithProrationPriceCadence string

const (
	NewPlanUnitWithProrationPriceCadenceAnnual     NewPlanUnitWithProrationPriceCadence = "annual"
	NewPlanUnitWithProrationPriceCadenceSemiAnnual NewPlanUnitWithProrationPriceCadence = "semi_annual"
	NewPlanUnitWithProrationPriceCadenceMonthly    NewPlanUnitWithProrationPriceCadence = "monthly"
	NewPlanUnitWithProrationPriceCadenceQuarterly  NewPlanUnitWithProrationPriceCadence = "quarterly"
	NewPlanUnitWithProrationPriceCadenceOneTime    NewPlanUnitWithProrationPriceCadence = "one_time"
	NewPlanUnitWithProrationPriceCadenceCustom     NewPlanUnitWithProrationPriceCadence = "custom"
)

func (r NewPlanUnitWithProrationPriceCadence) IsKnown() bool {
	switch r {
	case NewPlanUnitWithProrationPriceCadenceAnnual, NewPlanUnitWithProrationPriceCadenceSemiAnnual, NewPlanUnitWithProrationPriceCadenceMonthly, NewPlanUnitWithProrationPriceCadenceQuarterly, NewPlanUnitWithProrationPriceCadenceOneTime, NewPlanUnitWithProrationPriceCadenceCustom:
		return true
	}
	return false
}

type NewPlanUnitWithProrationPriceModelType string

const (
	NewPlanUnitWithProrationPriceModelTypeUnitWithProration NewPlanUnitWithProrationPriceModelType = "unit_with_proration"
)

func (r NewPlanUnitWithProrationPriceModelType) IsKnown() bool {
	switch r {
	case NewPlanUnitWithProrationPriceModelTypeUnitWithProration:
		return true
	}
	return false
}

type NewPlanUnitWithProrationPriceParam struct {
	ConversionRateType param.Field[NewPlanUnitWithProrationPriceConversionRateType] `json:"conversion_rate_type,required"`
	TieredConfig       param.Field[ConversionRateTieredConfigParam]                 `json:"tiered_config"`
	UnitConfig         param.Field[ConversionRateUnitConfigParam]                   `json:"unit_config"`
}

func (r NewPlanUnitWithProrationPriceParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r NewPlanUnitWithProrationPriceParam) ImplementsNewPlanUnitWithProrationPriceUnionParam() {}

// Satisfied by [shared.UnitConversionRateConfigParam],
// [shared.TieredConversionRateConfigParam], [NewPlanUnitWithProrationPriceParam].
type NewPlanUnitWithProrationPriceUnionParam interface {
	ImplementsNewPlanUnitWithProrationPriceUnionParam()
}

type NewPlanUnitWithProrationPriceConversionRateType string

const (
	NewPlanUnitWithProrationPriceConversionRateTypeUnit   NewPlanUnitWithProrationPriceConversionRateType = "unit"
	NewPlanUnitWithProrationPriceConversionRateTypeTiered NewPlanUnitWithProrationPriceConversionRateType = "tiered"
)

func (r NewPlanUnitWithProrationPriceConversionRateType) IsKnown() bool {
	switch r {
	case NewPlanUnitWithProrationPriceConversionRateTypeUnit, NewPlanUnitWithProrationPriceConversionRateTypeTiered:
		return true
	}
	return false
}

type NewUsageDiscountParam struct {
	AdjustmentType param.Field[NewUsageDiscountAdjustmentType] `json:"adjustment_type,required"`
	UsageDiscount  param.Field[float64]                        `json:"usage_discount,required"`
	// If set, the adjustment will apply to every price on the subscription.
	AppliesToAll param.Field[NewUsageDiscountAppliesToAll] `json:"applies_to_all"`
	// The set of item IDs to which this adjustment applies.
	AppliesToItemIDs param.Field[[]string] `json:"applies_to_item_ids"`
	// The set of price IDs to which this adjustment applies.
	AppliesToPriceIDs param.Field[[]string] `json:"applies_to_price_ids"`
	// If set, only prices in the specified currency will have the adjustment applied.
	Currency param.Field[string] `json:"currency"`
	// A list of filters that determine which prices this adjustment will apply to.
	Filters param.Field[[]TransformPriceFilterParam] `json:"filters"`
	// When false, this adjustment will be applied to a single price. Otherwise, it
	// will be applied at the invoice level, possibly to multiple prices.
	IsInvoiceLevel param.Field[bool] `json:"is_invoice_level"`
	// If set, only prices of the specified type will have the adjustment applied.
	PriceType param.Field[NewUsageDiscountPriceType] `json:"price_type"`
}

func (r NewUsageDiscountParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r NewUsageDiscountParam) ImplementsBetaNewPlanVersionParamsAddAdjustmentsAdjustmentUnion() {}

func (r NewUsageDiscountParam) ImplementsBetaNewPlanVersionParamsReplaceAdjustmentsAdjustmentUnion() {
}

func (r NewUsageDiscountParam) ImplementsBetaExternalPlanIDNewPlanVersionParamsAddAdjustmentsAdjustmentUnion() {
}

func (r NewUsageDiscountParam) ImplementsBetaExternalPlanIDNewPlanVersionParamsReplaceAdjustmentsAdjustmentUnion() {
}

func (r NewUsageDiscountParam) ImplementsPlanNewParamsAdjustmentsAdjustmentUnion() {}

func (r NewUsageDiscountParam) ImplementsSubscriptionNewParamsAddAdjustmentsAdjustmentUnion() {}

func (r NewUsageDiscountParam) ImplementsSubscriptionNewParamsReplaceAdjustmentsAdjustmentUnion() {}

func (r NewUsageDiscountParam) ImplementsSubscriptionPriceIntervalsParamsAddAdjustmentsAdjustmentUnion() {
}

func (r NewUsageDiscountParam) ImplementsSubscriptionSchedulePlanChangeParamsAddAdjustmentsAdjustmentUnion() {
}

func (r NewUsageDiscountParam) ImplementsSubscriptionSchedulePlanChangeParamsReplaceAdjustmentsAdjustmentUnion() {
}

type NewUsageDiscountAdjustmentType string

const (
	NewUsageDiscountAdjustmentTypeUsageDiscount NewUsageDiscountAdjustmentType = "usage_discount"
)

func (r NewUsageDiscountAdjustmentType) IsKnown() bool {
	switch r {
	case NewUsageDiscountAdjustmentTypeUsageDiscount:
		return true
	}
	return false
}

// If set, the adjustment will apply to every price on the subscription.
type NewUsageDiscountAppliesToAll bool

const (
	NewUsageDiscountAppliesToAllTrue NewUsageDiscountAppliesToAll = true
)

func (r NewUsageDiscountAppliesToAll) IsKnown() bool {
	switch r {
	case NewUsageDiscountAppliesToAllTrue:
		return true
	}
	return false
}

// If set, only prices of the specified type will have the adjustment applied.
type NewUsageDiscountPriceType string

const (
	NewUsageDiscountPriceTypeUsage          NewUsageDiscountPriceType = "usage"
	NewUsageDiscountPriceTypeFixedInAdvance NewUsageDiscountPriceType = "fixed_in_advance"
	NewUsageDiscountPriceTypeFixedInArrears NewUsageDiscountPriceType = "fixed_in_arrears"
	NewUsageDiscountPriceTypeFixed          NewUsageDiscountPriceType = "fixed"
	NewUsageDiscountPriceTypeInArrears      NewUsageDiscountPriceType = "in_arrears"
)

func (r NewUsageDiscountPriceType) IsKnown() bool {
	switch r {
	case NewUsageDiscountPriceTypeUsage, NewUsageDiscountPriceTypeFixedInAdvance, NewUsageDiscountPriceTypeFixedInArrears, NewUsageDiscountPriceTypeFixed, NewUsageDiscountPriceTypeInArrears:
		return true
	}
	return false
}

type OtherSubLineItem struct {
	// The total amount for this sub line item.
	Amount   string               `json:"amount,required"`
	Grouping SubLineItemGrouping  `json:"grouping,required,nullable"`
	Name     string               `json:"name,required"`
	Quantity float64              `json:"quantity,required"`
	Type     OtherSubLineItemType `json:"type,required"`
	JSON     otherSubLineItemJSON `json:"-"`
}

// otherSubLineItemJSON contains the JSON metadata for the struct
// [OtherSubLineItem]
type otherSubLineItemJSON struct {
	Amount      apijson.Field
	Grouping    apijson.Field
	Name        apijson.Field
	Quantity    apijson.Field
	Type        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *OtherSubLineItem) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r otherSubLineItemJSON) RawJSON() string {
	return r.raw
}

func (r OtherSubLineItem) ImplementsInvoiceLineItemsSubLineItem() {}

func (r OtherSubLineItem) ImplementsInvoiceLineItemNewResponseSubLineItem() {}

func (r OtherSubLineItem) ImplementsInvoiceFetchUpcomingResponseLineItemsSubLineItem() {}

type OtherSubLineItemType string

const (
	OtherSubLineItemTypeNull OtherSubLineItemType = "'null'"
)

func (r OtherSubLineItemType) IsKnown() bool {
	switch r {
	case OtherSubLineItemTypeNull:
		return true
	}
	return false
}

type PackageConfig struct {
	// A currency amount to rate usage by
	PackageAmount string `json:"package_amount,required"`
	// An integer amount to represent package size. For example, 1000 here would divide
	// usage by 1000 before multiplying by package_amount in rating
	PackageSize int64             `json:"package_size,required"`
	JSON        packageConfigJSON `json:"-"`
}

// packageConfigJSON contains the JSON metadata for the struct [PackageConfig]
type packageConfigJSON struct {
	PackageAmount apijson.Field
	PackageSize   apijson.Field
	raw           string
	ExtraFields   map[string]apijson.Field
}

func (r *PackageConfig) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r packageConfigJSON) RawJSON() string {
	return r.raw
}

type PackageConfigParam struct {
	// A currency amount to rate usage by
	PackageAmount param.Field[string] `json:"package_amount,required"`
	// An integer amount to represent package size. For example, 1000 here would divide
	// usage by 1000 before multiplying by package_amount in rating
	PackageSize param.Field[int64] `json:"package_size,required"`
}

func (r PackageConfigParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type PaginationMetadata struct {
	HasMore    bool                   `json:"has_more,required"`
	NextCursor string                 `json:"next_cursor,required,nullable"`
	JSON       paginationMetadataJSON `json:"-"`
}

// paginationMetadataJSON contains the JSON metadata for the struct
// [PaginationMetadata]
type paginationMetadataJSON struct {
	HasMore     apijson.Field
	NextCursor  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *PaginationMetadata) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r paginationMetadataJSON) RawJSON() string {
	return r.raw
}

type PerPriceCost struct {
	// The price object
	Price Price `json:"price,required"`
	// The price the cost is associated with
	PriceID string `json:"price_id,required"`
	// Price's contributions for the timeframe, excluding any minimums and discounts.
	Subtotal string `json:"subtotal,required"`
	// Price's contributions for the timeframe, including minimums and discounts.
	Total string `json:"total,required"`
	// The price's quantity for the timeframe
	Quantity float64          `json:"quantity,nullable"`
	JSON     perPriceCostJSON `json:"-"`
}

// perPriceCostJSON contains the JSON metadata for the struct [PerPriceCost]
type perPriceCostJSON struct {
	Price       apijson.Field
	PriceID     apijson.Field
	Subtotal    apijson.Field
	Total       apijson.Field
	Quantity    apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *PerPriceCost) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r perPriceCostJSON) RawJSON() string {
	return r.raw
}

type PercentageDiscount struct {
	DiscountType PercentageDiscountDiscountType `json:"discount_type,required"`
	// Only available if discount_type is `percentage`. This is a number between 0
	// and 1.
	PercentageDiscount float64 `json:"percentage_discount,required"`
	// List of price_ids that this discount applies to. For plan/plan phase discounts,
	// this can be a subset of prices.
	AppliesToPriceIDs []string `json:"applies_to_price_ids,nullable"`
	// The filters that determine which prices to apply this discount to.
	Filters []TransformPriceFilter `json:"filters,nullable"`
	Reason  string                 `json:"reason,nullable"`
	JSON    percentageDiscountJSON `json:"-"`
}

// percentageDiscountJSON contains the JSON metadata for the struct
// [PercentageDiscount]
type percentageDiscountJSON struct {
	DiscountType       apijson.Field
	PercentageDiscount apijson.Field
	AppliesToPriceIDs  apijson.Field
	Filters            apijson.Field
	Reason             apijson.Field
	raw                string
	ExtraFields        map[string]apijson.Field
}

func (r *PercentageDiscount) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r percentageDiscountJSON) RawJSON() string {
	return r.raw
}

func (r PercentageDiscount) ImplementsDiscount() {}

func (r PercentageDiscount) ImplementsInvoiceLevelDiscount() {}

func (r PercentageDiscount) ImplementsCouponDiscount() {}

type PercentageDiscountDiscountType string

const (
	PercentageDiscountDiscountTypePercentage PercentageDiscountDiscountType = "percentage"
)

func (r PercentageDiscountDiscountType) IsKnown() bool {
	switch r {
	case PercentageDiscountDiscountTypePercentage:
		return true
	}
	return false
}

type PercentageDiscountParam struct {
	DiscountType param.Field[PercentageDiscountDiscountType] `json:"discount_type,required"`
	// Only available if discount_type is `percentage`. This is a number between 0
	// and 1.
	PercentageDiscount param.Field[float64] `json:"percentage_discount,required"`
	// List of price_ids that this discount applies to. For plan/plan phase discounts,
	// this can be a subset of prices.
	AppliesToPriceIDs param.Field[[]string] `json:"applies_to_price_ids"`
	// The filters that determine which prices to apply this discount to.
	Filters param.Field[[]TransformPriceFilterParam] `json:"filters"`
	Reason  param.Field[string]                      `json:"reason"`
}

func (r PercentageDiscountParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r PercentageDiscountParam) ImplementsDiscountUnionParam() {}

type PercentageDiscountInterval struct {
	// The price interval ids that this discount interval applies to.
	AppliesToPriceIntervalIDs []string                               `json:"applies_to_price_interval_ids,required"`
	DiscountType              PercentageDiscountIntervalDiscountType `json:"discount_type,required"`
	// The end date of the discount interval.
	EndDate time.Time `json:"end_date,required,nullable" format:"date-time"`
	// The filters that determine which prices this discount interval applies to.
	Filters []TransformPriceFilter `json:"filters,required"`
	// Only available if discount_type is `percentage`.This is a number between 0
	// and 1.
	PercentageDiscount float64 `json:"percentage_discount,required"`
	// The start date of the discount interval.
	StartDate time.Time                      `json:"start_date,required" format:"date-time"`
	JSON      percentageDiscountIntervalJSON `json:"-"`
}

// percentageDiscountIntervalJSON contains the JSON metadata for the struct
// [PercentageDiscountInterval]
type percentageDiscountIntervalJSON struct {
	AppliesToPriceIntervalIDs apijson.Field
	DiscountType              apijson.Field
	EndDate                   apijson.Field
	Filters                   apijson.Field
	PercentageDiscount        apijson.Field
	StartDate                 apijson.Field
	raw                       string
	ExtraFields               map[string]apijson.Field
}

func (r *PercentageDiscountInterval) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r percentageDiscountIntervalJSON) RawJSON() string {
	return r.raw
}

func (r PercentageDiscountInterval) ImplementsSubscriptionDiscountInterval() {}

func (r PercentageDiscountInterval) ImplementsMutatedSubscriptionDiscountInterval() {}

type PercentageDiscountIntervalDiscountType string

const (
	PercentageDiscountIntervalDiscountTypePercentage PercentageDiscountIntervalDiscountType = "percentage"
)

func (r PercentageDiscountIntervalDiscountType) IsKnown() bool {
	switch r {
	case PercentageDiscountIntervalDiscountTypePercentage:
		return true
	}
	return false
}

type PlanPhaseAmountDiscountAdjustment struct {
	ID             string                                          `json:"id,required"`
	AdjustmentType PlanPhaseAmountDiscountAdjustmentAdjustmentType `json:"adjustment_type,required"`
	// The amount by which to discount the prices this adjustment applies to in a given
	// billing period.
	AmountDiscount string `json:"amount_discount,required"`
	// The price IDs that this adjustment applies to.
	//
	// Deprecated: deprecated
	AppliesToPriceIDs []string `json:"applies_to_price_ids,required"`
	// The filters that determine which prices to apply this adjustment to.
	Filters []TransformPriceFilter `json:"filters,required"`
	// True for adjustments that apply to an entire invocice, false for adjustments
	// that apply to only one price.
	IsInvoiceLevel bool `json:"is_invoice_level,required"`
	// The plan phase in which this adjustment is active.
	PlanPhaseOrder int64 `json:"plan_phase_order,required,nullable"`
	// The reason for the adjustment.
	Reason string `json:"reason,required,nullable"`
	// The adjustment id this adjustment replaces. This adjustment will take the place
	// of the replaced adjustment in plan version migrations.
	ReplacesAdjustmentID string                                `json:"replaces_adjustment_id,required,nullable"`
	JSON                 planPhaseAmountDiscountAdjustmentJSON `json:"-"`
}

// planPhaseAmountDiscountAdjustmentJSON contains the JSON metadata for the struct
// [PlanPhaseAmountDiscountAdjustment]
type planPhaseAmountDiscountAdjustmentJSON struct {
	ID                   apijson.Field
	AdjustmentType       apijson.Field
	AmountDiscount       apijson.Field
	AppliesToPriceIDs    apijson.Field
	Filters              apijson.Field
	IsInvoiceLevel       apijson.Field
	PlanPhaseOrder       apijson.Field
	Reason               apijson.Field
	ReplacesAdjustmentID apijson.Field
	raw                  string
	ExtraFields          map[string]apijson.Field
}

func (r *PlanPhaseAmountDiscountAdjustment) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r planPhaseAmountDiscountAdjustmentJSON) RawJSON() string {
	return r.raw
}

func (r PlanPhaseAmountDiscountAdjustment) ImplementsAdjustmentIntervalAdjustment() {}

func (r PlanPhaseAmountDiscountAdjustment) ImplementsPlanVersionAdjustment() {}

func (r PlanPhaseAmountDiscountAdjustment) ImplementsPlanAdjustment() {}

type PlanPhaseAmountDiscountAdjustmentAdjustmentType string

const (
	PlanPhaseAmountDiscountAdjustmentAdjustmentTypeAmountDiscount PlanPhaseAmountDiscountAdjustmentAdjustmentType = "amount_discount"
)

func (r PlanPhaseAmountDiscountAdjustmentAdjustmentType) IsKnown() bool {
	switch r {
	case PlanPhaseAmountDiscountAdjustmentAdjustmentTypeAmountDiscount:
		return true
	}
	return false
}

type PlanPhaseMaximumAdjustment struct {
	ID             string                                   `json:"id,required"`
	AdjustmentType PlanPhaseMaximumAdjustmentAdjustmentType `json:"adjustment_type,required"`
	// The price IDs that this adjustment applies to.
	//
	// Deprecated: deprecated
	AppliesToPriceIDs []string `json:"applies_to_price_ids,required"`
	// The filters that determine which prices to apply this adjustment to.
	Filters []TransformPriceFilter `json:"filters,required"`
	// True for adjustments that apply to an entire invocice, false for adjustments
	// that apply to only one price.
	IsInvoiceLevel bool `json:"is_invoice_level,required"`
	// The maximum amount to charge in a given billing period for the prices this
	// adjustment applies to.
	MaximumAmount string `json:"maximum_amount,required"`
	// The plan phase in which this adjustment is active.
	PlanPhaseOrder int64 `json:"plan_phase_order,required,nullable"`
	// The reason for the adjustment.
	Reason string `json:"reason,required,nullable"`
	// The adjustment id this adjustment replaces. This adjustment will take the place
	// of the replaced adjustment in plan version migrations.
	ReplacesAdjustmentID string                         `json:"replaces_adjustment_id,required,nullable"`
	JSON                 planPhaseMaximumAdjustmentJSON `json:"-"`
}

// planPhaseMaximumAdjustmentJSON contains the JSON metadata for the struct
// [PlanPhaseMaximumAdjustment]
type planPhaseMaximumAdjustmentJSON struct {
	ID                   apijson.Field
	AdjustmentType       apijson.Field
	AppliesToPriceIDs    apijson.Field
	Filters              apijson.Field
	IsInvoiceLevel       apijson.Field
	MaximumAmount        apijson.Field
	PlanPhaseOrder       apijson.Field
	Reason               apijson.Field
	ReplacesAdjustmentID apijson.Field
	raw                  string
	ExtraFields          map[string]apijson.Field
}

func (r *PlanPhaseMaximumAdjustment) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r planPhaseMaximumAdjustmentJSON) RawJSON() string {
	return r.raw
}

func (r PlanPhaseMaximumAdjustment) ImplementsAdjustmentIntervalAdjustment() {}

func (r PlanPhaseMaximumAdjustment) ImplementsPlanVersionAdjustment() {}

func (r PlanPhaseMaximumAdjustment) ImplementsPlanAdjustment() {}

type PlanPhaseMaximumAdjustmentAdjustmentType string

const (
	PlanPhaseMaximumAdjustmentAdjustmentTypeMaximum PlanPhaseMaximumAdjustmentAdjustmentType = "maximum"
)

func (r PlanPhaseMaximumAdjustmentAdjustmentType) IsKnown() bool {
	switch r {
	case PlanPhaseMaximumAdjustmentAdjustmentTypeMaximum:
		return true
	}
	return false
}

type PlanPhaseMinimumAdjustment struct {
	ID             string                                   `json:"id,required"`
	AdjustmentType PlanPhaseMinimumAdjustmentAdjustmentType `json:"adjustment_type,required"`
	// The price IDs that this adjustment applies to.
	//
	// Deprecated: deprecated
	AppliesToPriceIDs []string `json:"applies_to_price_ids,required"`
	// The filters that determine which prices to apply this adjustment to.
	Filters []TransformPriceFilter `json:"filters,required"`
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
	Reason string `json:"reason,required,nullable"`
	// The adjustment id this adjustment replaces. This adjustment will take the place
	// of the replaced adjustment in plan version migrations.
	ReplacesAdjustmentID string                         `json:"replaces_adjustment_id,required,nullable"`
	JSON                 planPhaseMinimumAdjustmentJSON `json:"-"`
}

// planPhaseMinimumAdjustmentJSON contains the JSON metadata for the struct
// [PlanPhaseMinimumAdjustment]
type planPhaseMinimumAdjustmentJSON struct {
	ID                   apijson.Field
	AdjustmentType       apijson.Field
	AppliesToPriceIDs    apijson.Field
	Filters              apijson.Field
	IsInvoiceLevel       apijson.Field
	ItemID               apijson.Field
	MinimumAmount        apijson.Field
	PlanPhaseOrder       apijson.Field
	Reason               apijson.Field
	ReplacesAdjustmentID apijson.Field
	raw                  string
	ExtraFields          map[string]apijson.Field
}

func (r *PlanPhaseMinimumAdjustment) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r planPhaseMinimumAdjustmentJSON) RawJSON() string {
	return r.raw
}

func (r PlanPhaseMinimumAdjustment) ImplementsAdjustmentIntervalAdjustment() {}

func (r PlanPhaseMinimumAdjustment) ImplementsPlanVersionAdjustment() {}

func (r PlanPhaseMinimumAdjustment) ImplementsPlanAdjustment() {}

type PlanPhaseMinimumAdjustmentAdjustmentType string

const (
	PlanPhaseMinimumAdjustmentAdjustmentTypeMinimum PlanPhaseMinimumAdjustmentAdjustmentType = "minimum"
)

func (r PlanPhaseMinimumAdjustmentAdjustmentType) IsKnown() bool {
	switch r {
	case PlanPhaseMinimumAdjustmentAdjustmentTypeMinimum:
		return true
	}
	return false
}

type PlanPhasePercentageDiscountAdjustment struct {
	ID             string                                              `json:"id,required"`
	AdjustmentType PlanPhasePercentageDiscountAdjustmentAdjustmentType `json:"adjustment_type,required"`
	// The price IDs that this adjustment applies to.
	//
	// Deprecated: deprecated
	AppliesToPriceIDs []string `json:"applies_to_price_ids,required"`
	// The filters that determine which prices to apply this adjustment to.
	Filters []TransformPriceFilter `json:"filters,required"`
	// True for adjustments that apply to an entire invocice, false for adjustments
	// that apply to only one price.
	IsInvoiceLevel bool `json:"is_invoice_level,required"`
	// The percentage (as a value between 0 and 1) by which to discount the price
	// intervals this adjustment applies to in a given billing period.
	PercentageDiscount float64 `json:"percentage_discount,required"`
	// The plan phase in which this adjustment is active.
	PlanPhaseOrder int64 `json:"plan_phase_order,required,nullable"`
	// The reason for the adjustment.
	Reason string `json:"reason,required,nullable"`
	// The adjustment id this adjustment replaces. This adjustment will take the place
	// of the replaced adjustment in plan version migrations.
	ReplacesAdjustmentID string                                    `json:"replaces_adjustment_id,required,nullable"`
	JSON                 planPhasePercentageDiscountAdjustmentJSON `json:"-"`
}

// planPhasePercentageDiscountAdjustmentJSON contains the JSON metadata for the
// struct [PlanPhasePercentageDiscountAdjustment]
type planPhasePercentageDiscountAdjustmentJSON struct {
	ID                   apijson.Field
	AdjustmentType       apijson.Field
	AppliesToPriceIDs    apijson.Field
	Filters              apijson.Field
	IsInvoiceLevel       apijson.Field
	PercentageDiscount   apijson.Field
	PlanPhaseOrder       apijson.Field
	Reason               apijson.Field
	ReplacesAdjustmentID apijson.Field
	raw                  string
	ExtraFields          map[string]apijson.Field
}

func (r *PlanPhasePercentageDiscountAdjustment) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r planPhasePercentageDiscountAdjustmentJSON) RawJSON() string {
	return r.raw
}

func (r PlanPhasePercentageDiscountAdjustment) ImplementsAdjustmentIntervalAdjustment() {}

func (r PlanPhasePercentageDiscountAdjustment) ImplementsPlanVersionAdjustment() {}

func (r PlanPhasePercentageDiscountAdjustment) ImplementsPlanAdjustment() {}

type PlanPhasePercentageDiscountAdjustmentAdjustmentType string

const (
	PlanPhasePercentageDiscountAdjustmentAdjustmentTypePercentageDiscount PlanPhasePercentageDiscountAdjustmentAdjustmentType = "percentage_discount"
)

func (r PlanPhasePercentageDiscountAdjustmentAdjustmentType) IsKnown() bool {
	switch r {
	case PlanPhasePercentageDiscountAdjustmentAdjustmentTypePercentageDiscount:
		return true
	}
	return false
}

type PlanPhaseUsageDiscountAdjustment struct {
	ID             string                                         `json:"id,required"`
	AdjustmentType PlanPhaseUsageDiscountAdjustmentAdjustmentType `json:"adjustment_type,required"`
	// The price IDs that this adjustment applies to.
	//
	// Deprecated: deprecated
	AppliesToPriceIDs []string `json:"applies_to_price_ids,required"`
	// The filters that determine which prices to apply this adjustment to.
	Filters []TransformPriceFilter `json:"filters,required"`
	// True for adjustments that apply to an entire invocice, false for adjustments
	// that apply to only one price.
	IsInvoiceLevel bool `json:"is_invoice_level,required"`
	// The plan phase in which this adjustment is active.
	PlanPhaseOrder int64 `json:"plan_phase_order,required,nullable"`
	// The reason for the adjustment.
	Reason string `json:"reason,required,nullable"`
	// The adjustment id this adjustment replaces. This adjustment will take the place
	// of the replaced adjustment in plan version migrations.
	ReplacesAdjustmentID string `json:"replaces_adjustment_id,required,nullable"`
	// The number of usage units by which to discount the price this adjustment applies
	// to in a given billing period.
	UsageDiscount float64                              `json:"usage_discount,required"`
	JSON          planPhaseUsageDiscountAdjustmentJSON `json:"-"`
}

// planPhaseUsageDiscountAdjustmentJSON contains the JSON metadata for the struct
// [PlanPhaseUsageDiscountAdjustment]
type planPhaseUsageDiscountAdjustmentJSON struct {
	ID                   apijson.Field
	AdjustmentType       apijson.Field
	AppliesToPriceIDs    apijson.Field
	Filters              apijson.Field
	IsInvoiceLevel       apijson.Field
	PlanPhaseOrder       apijson.Field
	Reason               apijson.Field
	ReplacesAdjustmentID apijson.Field
	UsageDiscount        apijson.Field
	raw                  string
	ExtraFields          map[string]apijson.Field
}

func (r *PlanPhaseUsageDiscountAdjustment) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r planPhaseUsageDiscountAdjustmentJSON) RawJSON() string {
	return r.raw
}

func (r PlanPhaseUsageDiscountAdjustment) ImplementsAdjustmentIntervalAdjustment() {}

func (r PlanPhaseUsageDiscountAdjustment) ImplementsPlanVersionAdjustment() {}

func (r PlanPhaseUsageDiscountAdjustment) ImplementsPlanAdjustment() {}

type PlanPhaseUsageDiscountAdjustmentAdjustmentType string

const (
	PlanPhaseUsageDiscountAdjustmentAdjustmentTypeUsageDiscount PlanPhaseUsageDiscountAdjustmentAdjustmentType = "usage_discount"
)

func (r PlanPhaseUsageDiscountAdjustmentAdjustmentType) IsKnown() bool {
	switch r {
	case PlanPhaseUsageDiscountAdjustmentAdjustmentTypeUsageDiscount:
		return true
	}
	return false
}

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
type Price struct {
	ID                        string                    `json:"id,required"`
	BillableMetric            BillableMetricTiny        `json:"billable_metric,required,nullable"`
	BillingCycleConfiguration BillingCycleConfiguration `json:"billing_cycle_configuration,required"`
	Cadence                   PriceCadence              `json:"cadence,required"`
	ConversionRate            float64                   `json:"conversion_rate,required,nullable"`
	// This field can have the runtime type of [PriceUnitPrice], [PricePackagePrice],
	// [PriceMatrixPrice], [PriceTieredPrice], [PriceTieredBPSPrice], [PriceBPSPrice],
	// [PriceBulkBPSPrice], [PriceBulkPrice], [PriceThresholdTotalAmountPrice],
	// [PriceTieredPackagePrice], [PriceGroupedTieredPrice],
	// [PriceTieredWithMinimumPrice], [PriceTieredPackageWithMinimumPrice],
	// [PricePackageWithAllocationPrice], [PriceUnitWithPercentPrice],
	// [PriceMatrixWithAllocationPrice], [PriceTieredWithProrationPrice],
	// [PriceUnitWithProrationPrice], [PriceGroupedAllocationPrice],
	// [PriceGroupedWithProratedMinimumPrice], [PriceGroupedWithMeteredMinimumPrice],
	// [PriceMatrixWithDisplayNamePrice], [PriceBulkWithProrationPrice],
	// [PriceGroupedTieredPackagePrice], [PriceMaxGroupTieredPackagePrice],
	// [PriceScalableMatrixWithUnitPricingPrice],
	// [PriceScalableMatrixWithTieredPricingPrice], [PriceCumulativeGroupedBulkPrice],
	// [PriceGroupedWithMinMaxThresholdsPrice].
	ConversionRateConfig interface{} `json:"conversion_rate_config,required"`
	CreatedAt            time.Time   `json:"created_at,required" format:"date-time"`
	CreditAllocation     Allocation  `json:"credit_allocation,required,nullable"`
	Currency             string      `json:"currency,required"`
	// Deprecated: deprecated
	Discount                    Discount                  `json:"discount,required,nullable"`
	ExternalPriceID             string                    `json:"external_price_id,required,nullable"`
	FixedPriceQuantity          float64                   `json:"fixed_price_quantity,required,nullable"`
	InvoicingCycleConfiguration BillingCycleConfiguration `json:"invoicing_cycle_configuration,required,nullable"`
	Item                        ItemSlim                  `json:"item,required"`
	// Deprecated: deprecated
	Maximum Maximum `json:"maximum,required,nullable"`
	// Deprecated: deprecated
	MaximumAmount string `json:"maximum_amount,required,nullable"`
	// This field can have the runtime type of [map[string]string].
	Metadata interface{} `json:"metadata,required"`
	// Deprecated: deprecated
	Minimum Minimum `json:"minimum,required,nullable"`
	// Deprecated: deprecated
	MinimumAmount  string         `json:"minimum_amount,required,nullable"`
	ModelType      PriceModelType `json:"model_type,required"`
	Name           string         `json:"name,required"`
	PlanPhaseOrder int64          `json:"plan_phase_order,required,nullable"`
	PriceType      PricePriceType `json:"price_type,required"`
	// The price id this price replaces. This price will take the place of the replaced
	// price in plan version migrations.
	ReplacesPriceID string        `json:"replaces_price_id,required,nullable"`
	BPSConfig       BPSConfig     `json:"bps_config"`
	BulkBPSConfig   BulkBPSConfig `json:"bulk_bps_config"`
	BulkConfig      BulkConfig    `json:"bulk_config"`
	// This field can have the runtime type of [map[string]interface{}].
	BulkWithProrationConfig interface{} `json:"bulk_with_proration_config"`
	// This field can have the runtime type of [map[string]interface{}].
	CumulativeGroupedBulkConfig   interface{}                   `json:"cumulative_grouped_bulk_config"`
	DimensionalPriceConfiguration DimensionalPriceConfiguration `json:"dimensional_price_configuration,nullable"`
	// This field can have the runtime type of [map[string]interface{}].
	GroupedAllocationConfig interface{} `json:"grouped_allocation_config"`
	// This field can have the runtime type of [map[string]interface{}].
	GroupedTieredConfig interface{} `json:"grouped_tiered_config"`
	// This field can have the runtime type of [map[string]interface{}].
	GroupedTieredPackageConfig interface{} `json:"grouped_tiered_package_config"`
	// This field can have the runtime type of [map[string]interface{}].
	GroupedWithMeteredMinimumConfig interface{} `json:"grouped_with_metered_minimum_config"`
	// This field can have the runtime type of [map[string]interface{}].
	GroupedWithMinMaxThresholdsConfig interface{} `json:"grouped_with_min_max_thresholds_config"`
	// This field can have the runtime type of [map[string]interface{}].
	GroupedWithProratedMinimumConfig interface{}                `json:"grouped_with_prorated_minimum_config"`
	MatrixConfig                     MatrixConfig               `json:"matrix_config"`
	MatrixWithAllocationConfig       MatrixWithAllocationConfig `json:"matrix_with_allocation_config"`
	// This field can have the runtime type of [map[string]interface{}].
	MatrixWithDisplayNameConfig interface{} `json:"matrix_with_display_name_config"`
	// This field can have the runtime type of [map[string]interface{}].
	MaxGroupTieredPackageConfig interface{}   `json:"max_group_tiered_package_config"`
	PackageConfig               PackageConfig `json:"package_config"`
	// This field can have the runtime type of [map[string]interface{}].
	PackageWithAllocationConfig interface{} `json:"package_with_allocation_config"`
	// This field can have the runtime type of [map[string]interface{}].
	ScalableMatrixWithTieredPricingConfig interface{} `json:"scalable_matrix_with_tiered_pricing_config"`
	// This field can have the runtime type of [map[string]interface{}].
	ScalableMatrixWithUnitPricingConfig interface{} `json:"scalable_matrix_with_unit_pricing_config"`
	// This field can have the runtime type of [map[string]interface{}].
	ThresholdTotalAmountConfig interface{}     `json:"threshold_total_amount_config"`
	TieredBPSConfig            TieredBPSConfig `json:"tiered_bps_config"`
	TieredConfig               TieredConfig    `json:"tiered_config"`
	// This field can have the runtime type of [map[string]interface{}].
	TieredPackageConfig interface{} `json:"tiered_package_config"`
	// This field can have the runtime type of [map[string]interface{}].
	TieredPackageWithMinimumConfig interface{} `json:"tiered_package_with_minimum_config"`
	// This field can have the runtime type of [map[string]interface{}].
	TieredWithMinimumConfig interface{} `json:"tiered_with_minimum_config"`
	// This field can have the runtime type of [map[string]interface{}].
	TieredWithProrationConfig interface{} `json:"tiered_with_proration_config"`
	UnitConfig                UnitConfig  `json:"unit_config"`
	// This field can have the runtime type of [map[string]interface{}].
	UnitWithPercentConfig interface{} `json:"unit_with_percent_config"`
	// This field can have the runtime type of [map[string]interface{}].
	UnitWithProrationConfig interface{} `json:"unit_with_proration_config"`
	JSON                    priceJSON   `json:"-"`
	union                   PriceUnion
}

// priceJSON contains the JSON metadata for the struct [Price]
type priceJSON struct {
	ID                                    apijson.Field
	BillableMetric                        apijson.Field
	BillingCycleConfiguration             apijson.Field
	Cadence                               apijson.Field
	ConversionRate                        apijson.Field
	ConversionRateConfig                  apijson.Field
	CreatedAt                             apijson.Field
	CreditAllocation                      apijson.Field
	Currency                              apijson.Field
	Discount                              apijson.Field
	ExternalPriceID                       apijson.Field
	FixedPriceQuantity                    apijson.Field
	InvoicingCycleConfiguration           apijson.Field
	Item                                  apijson.Field
	Maximum                               apijson.Field
	MaximumAmount                         apijson.Field
	Metadata                              apijson.Field
	Minimum                               apijson.Field
	MinimumAmount                         apijson.Field
	ModelType                             apijson.Field
	Name                                  apijson.Field
	PlanPhaseOrder                        apijson.Field
	PriceType                             apijson.Field
	ReplacesPriceID                       apijson.Field
	BPSConfig                             apijson.Field
	BulkBPSConfig                         apijson.Field
	BulkConfig                            apijson.Field
	BulkWithProrationConfig               apijson.Field
	CumulativeGroupedBulkConfig           apijson.Field
	DimensionalPriceConfiguration         apijson.Field
	GroupedAllocationConfig               apijson.Field
	GroupedTieredConfig                   apijson.Field
	GroupedTieredPackageConfig            apijson.Field
	GroupedWithMeteredMinimumConfig       apijson.Field
	GroupedWithMinMaxThresholdsConfig     apijson.Field
	GroupedWithProratedMinimumConfig      apijson.Field
	MatrixConfig                          apijson.Field
	MatrixWithAllocationConfig            apijson.Field
	MatrixWithDisplayNameConfig           apijson.Field
	MaxGroupTieredPackageConfig           apijson.Field
	PackageConfig                         apijson.Field
	PackageWithAllocationConfig           apijson.Field
	ScalableMatrixWithTieredPricingConfig apijson.Field
	ScalableMatrixWithUnitPricingConfig   apijson.Field
	ThresholdTotalAmountConfig            apijson.Field
	TieredBPSConfig                       apijson.Field
	TieredConfig                          apijson.Field
	TieredPackageConfig                   apijson.Field
	TieredPackageWithMinimumConfig        apijson.Field
	TieredWithMinimumConfig               apijson.Field
	TieredWithProrationConfig             apijson.Field
	UnitConfig                            apijson.Field
	UnitWithPercentConfig                 apijson.Field
	UnitWithProrationConfig               apijson.Field
	raw                                   string
	ExtraFields                           map[string]apijson.Field
}

func (r priceJSON) RawJSON() string {
	return r.raw
}

func (r *Price) UnmarshalJSON(data []byte) (err error) {
	*r = Price{}
	err = apijson.UnmarshalRoot(data, &r.union)
	if err != nil {
		return err
	}
	return apijson.Port(r.union, &r)
}

// AsUnion returns a [PriceUnion] interface which you can cast to the specific
// types for more type safety.
//
// Possible runtime types of the union are [PriceUnitPrice], [PricePackagePrice],
// [PriceMatrixPrice], [PriceTieredPrice], [PriceTieredBPSPrice], [PriceBPSPrice],
// [PriceBulkBPSPrice], [PriceBulkPrice], [PriceThresholdTotalAmountPrice],
// [PriceTieredPackagePrice], [PriceGroupedTieredPrice],
// [PriceTieredWithMinimumPrice], [PriceTieredPackageWithMinimumPrice],
// [PricePackageWithAllocationPrice], [PriceUnitWithPercentPrice],
// [PriceMatrixWithAllocationPrice], [PriceTieredWithProrationPrice],
// [PriceUnitWithProrationPrice], [PriceGroupedAllocationPrice],
// [PriceGroupedWithProratedMinimumPrice], [PriceGroupedWithMeteredMinimumPrice],
// [PriceMatrixWithDisplayNamePrice], [PriceBulkWithProrationPrice],
// [PriceGroupedTieredPackagePrice], [PriceMaxGroupTieredPackagePrice],
// [PriceScalableMatrixWithUnitPricingPrice],
// [PriceScalableMatrixWithTieredPricingPrice], [PriceCumulativeGroupedBulkPrice],
// [PriceGroupedWithMinMaxThresholdsPrice].
func (r Price) AsUnion() PriceUnion {
	return r.union
}

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
//
// Union satisfied by [PriceUnitPrice], [PricePackagePrice], [PriceMatrixPrice],
// [PriceTieredPrice], [PriceTieredBPSPrice], [PriceBPSPrice], [PriceBulkBPSPrice],
// [PriceBulkPrice], [PriceThresholdTotalAmountPrice], [PriceTieredPackagePrice],
// [PriceGroupedTieredPrice], [PriceTieredWithMinimumPrice],
// [PriceTieredPackageWithMinimumPrice], [PricePackageWithAllocationPrice],
// [PriceUnitWithPercentPrice], [PriceMatrixWithAllocationPrice],
// [PriceTieredWithProrationPrice], [PriceUnitWithProrationPrice],
// [PriceGroupedAllocationPrice], [PriceGroupedWithProratedMinimumPrice],
// [PriceGroupedWithMeteredMinimumPrice], [PriceMatrixWithDisplayNamePrice],
// [PriceBulkWithProrationPrice], [PriceGroupedTieredPackagePrice],
// [PriceMaxGroupTieredPackagePrice], [PriceScalableMatrixWithUnitPricingPrice],
// [PriceScalableMatrixWithTieredPricingPrice], [PriceCumulativeGroupedBulkPrice]
// or [PriceGroupedWithMinMaxThresholdsPrice].
type PriceUnion interface {
	implementsPrice()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*PriceUnion)(nil)).Elem(),
		"model_type",
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(PriceUnitPrice{}),
			DiscriminatorValue: "unit",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(PricePackagePrice{}),
			DiscriminatorValue: "package",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(PriceMatrixPrice{}),
			DiscriminatorValue: "matrix",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(PriceTieredPrice{}),
			DiscriminatorValue: "tiered",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(PriceTieredBPSPrice{}),
			DiscriminatorValue: "tiered_bps",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(PriceBPSPrice{}),
			DiscriminatorValue: "bps",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(PriceBulkBPSPrice{}),
			DiscriminatorValue: "bulk_bps",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(PriceBulkPrice{}),
			DiscriminatorValue: "bulk",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(PriceThresholdTotalAmountPrice{}),
			DiscriminatorValue: "threshold_total_amount",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(PriceTieredPackagePrice{}),
			DiscriminatorValue: "tiered_package",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(PriceGroupedTieredPrice{}),
			DiscriminatorValue: "grouped_tiered",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(PriceTieredWithMinimumPrice{}),
			DiscriminatorValue: "tiered_with_minimum",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(PriceTieredPackageWithMinimumPrice{}),
			DiscriminatorValue: "tiered_package_with_minimum",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(PricePackageWithAllocationPrice{}),
			DiscriminatorValue: "package_with_allocation",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(PriceUnitWithPercentPrice{}),
			DiscriminatorValue: "unit_with_percent",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(PriceMatrixWithAllocationPrice{}),
			DiscriminatorValue: "matrix_with_allocation",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(PriceTieredWithProrationPrice{}),
			DiscriminatorValue: "tiered_with_proration",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(PriceUnitWithProrationPrice{}),
			DiscriminatorValue: "unit_with_proration",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(PriceGroupedAllocationPrice{}),
			DiscriminatorValue: "grouped_allocation",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(PriceGroupedWithProratedMinimumPrice{}),
			DiscriminatorValue: "grouped_with_prorated_minimum",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(PriceGroupedWithMeteredMinimumPrice{}),
			DiscriminatorValue: "grouped_with_metered_minimum",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(PriceMatrixWithDisplayNamePrice{}),
			DiscriminatorValue: "matrix_with_display_name",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(PriceBulkWithProrationPrice{}),
			DiscriminatorValue: "bulk_with_proration",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(PriceGroupedTieredPackagePrice{}),
			DiscriminatorValue: "grouped_tiered_package",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(PriceMaxGroupTieredPackagePrice{}),
			DiscriminatorValue: "max_group_tiered_package",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(PriceScalableMatrixWithUnitPricingPrice{}),
			DiscriminatorValue: "scalable_matrix_with_unit_pricing",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(PriceScalableMatrixWithTieredPricingPrice{}),
			DiscriminatorValue: "scalable_matrix_with_tiered_pricing",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(PriceCumulativeGroupedBulkPrice{}),
			DiscriminatorValue: "cumulative_grouped_bulk",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(PriceGroupedWithMinMaxThresholdsPrice{}),
			DiscriminatorValue: "grouped_with_min_max_thresholds",
		},
	)
}

type PriceUnitPrice struct {
	ID                        string                    `json:"id,required"`
	BillableMetric            BillableMetricTiny        `json:"billable_metric,required,nullable"`
	BillingCycleConfiguration BillingCycleConfiguration `json:"billing_cycle_configuration,required"`
	Cadence                   PriceUnitPriceCadence     `json:"cadence,required"`
	ConversionRate            float64                   `json:"conversion_rate,required,nullable"`
	ConversionRateConfig      PriceUnitPrice            `json:"conversion_rate_config,required,nullable"`
	CreatedAt                 time.Time                 `json:"created_at,required" format:"date-time"`
	CreditAllocation          Allocation                `json:"credit_allocation,required,nullable"`
	Currency                  string                    `json:"currency,required"`
	// Deprecated: deprecated
	Discount                    Discount                  `json:"discount,required,nullable"`
	ExternalPriceID             string                    `json:"external_price_id,required,nullable"`
	FixedPriceQuantity          float64                   `json:"fixed_price_quantity,required,nullable"`
	InvoicingCycleConfiguration BillingCycleConfiguration `json:"invoicing_cycle_configuration,required,nullable"`
	Item                        ItemSlim                  `json:"item,required"`
	// Deprecated: deprecated
	Maximum Maximum `json:"maximum,required,nullable"`
	// Deprecated: deprecated
	MaximumAmount string `json:"maximum_amount,required,nullable"`
	// User specified key-value pairs for the resource. If not present, this defaults
	// to an empty dictionary. Individual keys can be removed by setting the value to
	// `null`, and the entire metadata mapping can be cleared by setting `metadata` to
	// `null`.
	Metadata map[string]string `json:"metadata,required"`
	// Deprecated: deprecated
	Minimum Minimum `json:"minimum,required,nullable"`
	// Deprecated: deprecated
	MinimumAmount  string                  `json:"minimum_amount,required,nullable"`
	ModelType      PriceUnitPriceModelType `json:"model_type,required"`
	Name           string                  `json:"name,required"`
	PlanPhaseOrder int64                   `json:"plan_phase_order,required,nullable"`
	PriceType      PriceUnitPricePriceType `json:"price_type,required"`
	// The price id this price replaces. This price will take the place of the replaced
	// price in plan version migrations.
	ReplacesPriceID               string                        `json:"replaces_price_id,required,nullable"`
	UnitConfig                    UnitConfig                    `json:"unit_config,required"`
	DimensionalPriceConfiguration DimensionalPriceConfiguration `json:"dimensional_price_configuration,nullable"`
	JSON                          priceUnitPriceJSON            `json:"-"`
}

// priceUnitPriceJSON contains the JSON metadata for the struct [PriceUnitPrice]
type priceUnitPriceJSON struct {
	ID                            apijson.Field
	BillableMetric                apijson.Field
	BillingCycleConfiguration     apijson.Field
	Cadence                       apijson.Field
	ConversionRate                apijson.Field
	ConversionRateConfig          apijson.Field
	CreatedAt                     apijson.Field
	CreditAllocation              apijson.Field
	Currency                      apijson.Field
	Discount                      apijson.Field
	ExternalPriceID               apijson.Field
	FixedPriceQuantity            apijson.Field
	InvoicingCycleConfiguration   apijson.Field
	Item                          apijson.Field
	Maximum                       apijson.Field
	MaximumAmount                 apijson.Field
	Metadata                      apijson.Field
	Minimum                       apijson.Field
	MinimumAmount                 apijson.Field
	ModelType                     apijson.Field
	Name                          apijson.Field
	PlanPhaseOrder                apijson.Field
	PriceType                     apijson.Field
	ReplacesPriceID               apijson.Field
	UnitConfig                    apijson.Field
	DimensionalPriceConfiguration apijson.Field
	raw                           string
	ExtraFields                   map[string]apijson.Field
}

func (r *PriceUnitPrice) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r priceUnitPriceJSON) RawJSON() string {
	return r.raw
}

func (r PriceUnitPrice) implementsPrice() {}

type PriceUnitPriceCadence string

const (
	PriceUnitPriceCadenceOneTime    PriceUnitPriceCadence = "one_time"
	PriceUnitPriceCadenceMonthly    PriceUnitPriceCadence = "monthly"
	PriceUnitPriceCadenceQuarterly  PriceUnitPriceCadence = "quarterly"
	PriceUnitPriceCadenceSemiAnnual PriceUnitPriceCadence = "semi_annual"
	PriceUnitPriceCadenceAnnual     PriceUnitPriceCadence = "annual"
	PriceUnitPriceCadenceCustom     PriceUnitPriceCadence = "custom"
)

func (r PriceUnitPriceCadence) IsKnown() bool {
	switch r {
	case PriceUnitPriceCadenceOneTime, PriceUnitPriceCadenceMonthly, PriceUnitPriceCadenceQuarterly, PriceUnitPriceCadenceSemiAnnual, PriceUnitPriceCadenceAnnual, PriceUnitPriceCadenceCustom:
		return true
	}
	return false
}

type PriceUnitPriceModelType string

const (
	PriceUnitPriceModelTypeUnit PriceUnitPriceModelType = "unit"
)

func (r PriceUnitPriceModelType) IsKnown() bool {
	switch r {
	case PriceUnitPriceModelTypeUnit:
		return true
	}
	return false
}

type PriceUnitPricePriceType string

const (
	PriceUnitPricePriceTypeUsagePrice PriceUnitPricePriceType = "usage_price"
	PriceUnitPricePriceTypeFixedPrice PriceUnitPricePriceType = "fixed_price"
)

func (r PriceUnitPricePriceType) IsKnown() bool {
	switch r {
	case PriceUnitPricePriceTypeUsagePrice, PriceUnitPricePriceTypeFixedPrice:
		return true
	}
	return false
}

type PricePackagePrice struct {
	ID                        string                    `json:"id,required"`
	BillableMetric            BillableMetricTiny        `json:"billable_metric,required,nullable"`
	BillingCycleConfiguration BillingCycleConfiguration `json:"billing_cycle_configuration,required"`
	Cadence                   PricePackagePriceCadence  `json:"cadence,required"`
	ConversionRate            float64                   `json:"conversion_rate,required,nullable"`
	ConversionRateConfig      PricePackagePrice         `json:"conversion_rate_config,required,nullable"`
	CreatedAt                 time.Time                 `json:"created_at,required" format:"date-time"`
	CreditAllocation          Allocation                `json:"credit_allocation,required,nullable"`
	Currency                  string                    `json:"currency,required"`
	// Deprecated: deprecated
	Discount                    Discount                  `json:"discount,required,nullable"`
	ExternalPriceID             string                    `json:"external_price_id,required,nullable"`
	FixedPriceQuantity          float64                   `json:"fixed_price_quantity,required,nullable"`
	InvoicingCycleConfiguration BillingCycleConfiguration `json:"invoicing_cycle_configuration,required,nullable"`
	Item                        ItemSlim                  `json:"item,required"`
	// Deprecated: deprecated
	Maximum Maximum `json:"maximum,required,nullable"`
	// Deprecated: deprecated
	MaximumAmount string `json:"maximum_amount,required,nullable"`
	// User specified key-value pairs for the resource. If not present, this defaults
	// to an empty dictionary. Individual keys can be removed by setting the value to
	// `null`, and the entire metadata mapping can be cleared by setting `metadata` to
	// `null`.
	Metadata map[string]string `json:"metadata,required"`
	// Deprecated: deprecated
	Minimum Minimum `json:"minimum,required,nullable"`
	// Deprecated: deprecated
	MinimumAmount  string                     `json:"minimum_amount,required,nullable"`
	ModelType      PricePackagePriceModelType `json:"model_type,required"`
	Name           string                     `json:"name,required"`
	PackageConfig  PackageConfig              `json:"package_config,required"`
	PlanPhaseOrder int64                      `json:"plan_phase_order,required,nullable"`
	PriceType      PricePackagePricePriceType `json:"price_type,required"`
	// The price id this price replaces. This price will take the place of the replaced
	// price in plan version migrations.
	ReplacesPriceID               string                        `json:"replaces_price_id,required,nullable"`
	DimensionalPriceConfiguration DimensionalPriceConfiguration `json:"dimensional_price_configuration,nullable"`
	JSON                          pricePackagePriceJSON         `json:"-"`
}

// pricePackagePriceJSON contains the JSON metadata for the struct
// [PricePackagePrice]
type pricePackagePriceJSON struct {
	ID                            apijson.Field
	BillableMetric                apijson.Field
	BillingCycleConfiguration     apijson.Field
	Cadence                       apijson.Field
	ConversionRate                apijson.Field
	ConversionRateConfig          apijson.Field
	CreatedAt                     apijson.Field
	CreditAllocation              apijson.Field
	Currency                      apijson.Field
	Discount                      apijson.Field
	ExternalPriceID               apijson.Field
	FixedPriceQuantity            apijson.Field
	InvoicingCycleConfiguration   apijson.Field
	Item                          apijson.Field
	Maximum                       apijson.Field
	MaximumAmount                 apijson.Field
	Metadata                      apijson.Field
	Minimum                       apijson.Field
	MinimumAmount                 apijson.Field
	ModelType                     apijson.Field
	Name                          apijson.Field
	PackageConfig                 apijson.Field
	PlanPhaseOrder                apijson.Field
	PriceType                     apijson.Field
	ReplacesPriceID               apijson.Field
	DimensionalPriceConfiguration apijson.Field
	raw                           string
	ExtraFields                   map[string]apijson.Field
}

func (r *PricePackagePrice) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r pricePackagePriceJSON) RawJSON() string {
	return r.raw
}

func (r PricePackagePrice) implementsPrice() {}

type PricePackagePriceCadence string

const (
	PricePackagePriceCadenceOneTime    PricePackagePriceCadence = "one_time"
	PricePackagePriceCadenceMonthly    PricePackagePriceCadence = "monthly"
	PricePackagePriceCadenceQuarterly  PricePackagePriceCadence = "quarterly"
	PricePackagePriceCadenceSemiAnnual PricePackagePriceCadence = "semi_annual"
	PricePackagePriceCadenceAnnual     PricePackagePriceCadence = "annual"
	PricePackagePriceCadenceCustom     PricePackagePriceCadence = "custom"
)

func (r PricePackagePriceCadence) IsKnown() bool {
	switch r {
	case PricePackagePriceCadenceOneTime, PricePackagePriceCadenceMonthly, PricePackagePriceCadenceQuarterly, PricePackagePriceCadenceSemiAnnual, PricePackagePriceCadenceAnnual, PricePackagePriceCadenceCustom:
		return true
	}
	return false
}

type PricePackagePriceModelType string

const (
	PricePackagePriceModelTypePackage PricePackagePriceModelType = "package"
)

func (r PricePackagePriceModelType) IsKnown() bool {
	switch r {
	case PricePackagePriceModelTypePackage:
		return true
	}
	return false
}

type PricePackagePricePriceType string

const (
	PricePackagePricePriceTypeUsagePrice PricePackagePricePriceType = "usage_price"
	PricePackagePricePriceTypeFixedPrice PricePackagePricePriceType = "fixed_price"
)

func (r PricePackagePricePriceType) IsKnown() bool {
	switch r {
	case PricePackagePricePriceTypeUsagePrice, PricePackagePricePriceTypeFixedPrice:
		return true
	}
	return false
}

type PriceMatrixPrice struct {
	ID                        string                    `json:"id,required"`
	BillableMetric            BillableMetricTiny        `json:"billable_metric,required,nullable"`
	BillingCycleConfiguration BillingCycleConfiguration `json:"billing_cycle_configuration,required"`
	Cadence                   PriceMatrixPriceCadence   `json:"cadence,required"`
	ConversionRate            float64                   `json:"conversion_rate,required,nullable"`
	ConversionRateConfig      PriceMatrixPrice          `json:"conversion_rate_config,required,nullable"`
	CreatedAt                 time.Time                 `json:"created_at,required" format:"date-time"`
	CreditAllocation          Allocation                `json:"credit_allocation,required,nullable"`
	Currency                  string                    `json:"currency,required"`
	// Deprecated: deprecated
	Discount                    Discount                  `json:"discount,required,nullable"`
	ExternalPriceID             string                    `json:"external_price_id,required,nullable"`
	FixedPriceQuantity          float64                   `json:"fixed_price_quantity,required,nullable"`
	InvoicingCycleConfiguration BillingCycleConfiguration `json:"invoicing_cycle_configuration,required,nullable"`
	Item                        ItemSlim                  `json:"item,required"`
	MatrixConfig                MatrixConfig              `json:"matrix_config,required"`
	// Deprecated: deprecated
	Maximum Maximum `json:"maximum,required,nullable"`
	// Deprecated: deprecated
	MaximumAmount string `json:"maximum_amount,required,nullable"`
	// User specified key-value pairs for the resource. If not present, this defaults
	// to an empty dictionary. Individual keys can be removed by setting the value to
	// `null`, and the entire metadata mapping can be cleared by setting `metadata` to
	// `null`.
	Metadata map[string]string `json:"metadata,required"`
	// Deprecated: deprecated
	Minimum Minimum `json:"minimum,required,nullable"`
	// Deprecated: deprecated
	MinimumAmount  string                    `json:"minimum_amount,required,nullable"`
	ModelType      PriceMatrixPriceModelType `json:"model_type,required"`
	Name           string                    `json:"name,required"`
	PlanPhaseOrder int64                     `json:"plan_phase_order,required,nullable"`
	PriceType      PriceMatrixPricePriceType `json:"price_type,required"`
	// The price id this price replaces. This price will take the place of the replaced
	// price in plan version migrations.
	ReplacesPriceID               string                        `json:"replaces_price_id,required,nullable"`
	DimensionalPriceConfiguration DimensionalPriceConfiguration `json:"dimensional_price_configuration,nullable"`
	JSON                          priceMatrixPriceJSON          `json:"-"`
}

// priceMatrixPriceJSON contains the JSON metadata for the struct
// [PriceMatrixPrice]
type priceMatrixPriceJSON struct {
	ID                            apijson.Field
	BillableMetric                apijson.Field
	BillingCycleConfiguration     apijson.Field
	Cadence                       apijson.Field
	ConversionRate                apijson.Field
	ConversionRateConfig          apijson.Field
	CreatedAt                     apijson.Field
	CreditAllocation              apijson.Field
	Currency                      apijson.Field
	Discount                      apijson.Field
	ExternalPriceID               apijson.Field
	FixedPriceQuantity            apijson.Field
	InvoicingCycleConfiguration   apijson.Field
	Item                          apijson.Field
	MatrixConfig                  apijson.Field
	Maximum                       apijson.Field
	MaximumAmount                 apijson.Field
	Metadata                      apijson.Field
	Minimum                       apijson.Field
	MinimumAmount                 apijson.Field
	ModelType                     apijson.Field
	Name                          apijson.Field
	PlanPhaseOrder                apijson.Field
	PriceType                     apijson.Field
	ReplacesPriceID               apijson.Field
	DimensionalPriceConfiguration apijson.Field
	raw                           string
	ExtraFields                   map[string]apijson.Field
}

func (r *PriceMatrixPrice) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r priceMatrixPriceJSON) RawJSON() string {
	return r.raw
}

func (r PriceMatrixPrice) implementsPrice() {}

type PriceMatrixPriceCadence string

const (
	PriceMatrixPriceCadenceOneTime    PriceMatrixPriceCadence = "one_time"
	PriceMatrixPriceCadenceMonthly    PriceMatrixPriceCadence = "monthly"
	PriceMatrixPriceCadenceQuarterly  PriceMatrixPriceCadence = "quarterly"
	PriceMatrixPriceCadenceSemiAnnual PriceMatrixPriceCadence = "semi_annual"
	PriceMatrixPriceCadenceAnnual     PriceMatrixPriceCadence = "annual"
	PriceMatrixPriceCadenceCustom     PriceMatrixPriceCadence = "custom"
)

func (r PriceMatrixPriceCadence) IsKnown() bool {
	switch r {
	case PriceMatrixPriceCadenceOneTime, PriceMatrixPriceCadenceMonthly, PriceMatrixPriceCadenceQuarterly, PriceMatrixPriceCadenceSemiAnnual, PriceMatrixPriceCadenceAnnual, PriceMatrixPriceCadenceCustom:
		return true
	}
	return false
}

type PriceMatrixPriceModelType string

const (
	PriceMatrixPriceModelTypeMatrix PriceMatrixPriceModelType = "matrix"
)

func (r PriceMatrixPriceModelType) IsKnown() bool {
	switch r {
	case PriceMatrixPriceModelTypeMatrix:
		return true
	}
	return false
}

type PriceMatrixPricePriceType string

const (
	PriceMatrixPricePriceTypeUsagePrice PriceMatrixPricePriceType = "usage_price"
	PriceMatrixPricePriceTypeFixedPrice PriceMatrixPricePriceType = "fixed_price"
)

func (r PriceMatrixPricePriceType) IsKnown() bool {
	switch r {
	case PriceMatrixPricePriceTypeUsagePrice, PriceMatrixPricePriceTypeFixedPrice:
		return true
	}
	return false
}

type PriceTieredPrice struct {
	ID                        string                    `json:"id,required"`
	BillableMetric            BillableMetricTiny        `json:"billable_metric,required,nullable"`
	BillingCycleConfiguration BillingCycleConfiguration `json:"billing_cycle_configuration,required"`
	Cadence                   PriceTieredPriceCadence   `json:"cadence,required"`
	ConversionRate            float64                   `json:"conversion_rate,required,nullable"`
	ConversionRateConfig      PriceTieredPrice          `json:"conversion_rate_config,required,nullable"`
	CreatedAt                 time.Time                 `json:"created_at,required" format:"date-time"`
	CreditAllocation          Allocation                `json:"credit_allocation,required,nullable"`
	Currency                  string                    `json:"currency,required"`
	// Deprecated: deprecated
	Discount                    Discount                  `json:"discount,required,nullable"`
	ExternalPriceID             string                    `json:"external_price_id,required,nullable"`
	FixedPriceQuantity          float64                   `json:"fixed_price_quantity,required,nullable"`
	InvoicingCycleConfiguration BillingCycleConfiguration `json:"invoicing_cycle_configuration,required,nullable"`
	Item                        ItemSlim                  `json:"item,required"`
	// Deprecated: deprecated
	Maximum Maximum `json:"maximum,required,nullable"`
	// Deprecated: deprecated
	MaximumAmount string `json:"maximum_amount,required,nullable"`
	// User specified key-value pairs for the resource. If not present, this defaults
	// to an empty dictionary. Individual keys can be removed by setting the value to
	// `null`, and the entire metadata mapping can be cleared by setting `metadata` to
	// `null`.
	Metadata map[string]string `json:"metadata,required"`
	// Deprecated: deprecated
	Minimum Minimum `json:"minimum,required,nullable"`
	// Deprecated: deprecated
	MinimumAmount  string                    `json:"minimum_amount,required,nullable"`
	ModelType      PriceTieredPriceModelType `json:"model_type,required"`
	Name           string                    `json:"name,required"`
	PlanPhaseOrder int64                     `json:"plan_phase_order,required,nullable"`
	PriceType      PriceTieredPricePriceType `json:"price_type,required"`
	// The price id this price replaces. This price will take the place of the replaced
	// price in plan version migrations.
	ReplacesPriceID               string                        `json:"replaces_price_id,required,nullable"`
	TieredConfig                  TieredConfig                  `json:"tiered_config,required"`
	DimensionalPriceConfiguration DimensionalPriceConfiguration `json:"dimensional_price_configuration,nullable"`
	JSON                          priceTieredPriceJSON          `json:"-"`
}

// priceTieredPriceJSON contains the JSON metadata for the struct
// [PriceTieredPrice]
type priceTieredPriceJSON struct {
	ID                            apijson.Field
	BillableMetric                apijson.Field
	BillingCycleConfiguration     apijson.Field
	Cadence                       apijson.Field
	ConversionRate                apijson.Field
	ConversionRateConfig          apijson.Field
	CreatedAt                     apijson.Field
	CreditAllocation              apijson.Field
	Currency                      apijson.Field
	Discount                      apijson.Field
	ExternalPriceID               apijson.Field
	FixedPriceQuantity            apijson.Field
	InvoicingCycleConfiguration   apijson.Field
	Item                          apijson.Field
	Maximum                       apijson.Field
	MaximumAmount                 apijson.Field
	Metadata                      apijson.Field
	Minimum                       apijson.Field
	MinimumAmount                 apijson.Field
	ModelType                     apijson.Field
	Name                          apijson.Field
	PlanPhaseOrder                apijson.Field
	PriceType                     apijson.Field
	ReplacesPriceID               apijson.Field
	TieredConfig                  apijson.Field
	DimensionalPriceConfiguration apijson.Field
	raw                           string
	ExtraFields                   map[string]apijson.Field
}

func (r *PriceTieredPrice) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r priceTieredPriceJSON) RawJSON() string {
	return r.raw
}

func (r PriceTieredPrice) implementsPrice() {}

type PriceTieredPriceCadence string

const (
	PriceTieredPriceCadenceOneTime    PriceTieredPriceCadence = "one_time"
	PriceTieredPriceCadenceMonthly    PriceTieredPriceCadence = "monthly"
	PriceTieredPriceCadenceQuarterly  PriceTieredPriceCadence = "quarterly"
	PriceTieredPriceCadenceSemiAnnual PriceTieredPriceCadence = "semi_annual"
	PriceTieredPriceCadenceAnnual     PriceTieredPriceCadence = "annual"
	PriceTieredPriceCadenceCustom     PriceTieredPriceCadence = "custom"
)

func (r PriceTieredPriceCadence) IsKnown() bool {
	switch r {
	case PriceTieredPriceCadenceOneTime, PriceTieredPriceCadenceMonthly, PriceTieredPriceCadenceQuarterly, PriceTieredPriceCadenceSemiAnnual, PriceTieredPriceCadenceAnnual, PriceTieredPriceCadenceCustom:
		return true
	}
	return false
}

type PriceTieredPriceModelType string

const (
	PriceTieredPriceModelTypeTiered PriceTieredPriceModelType = "tiered"
)

func (r PriceTieredPriceModelType) IsKnown() bool {
	switch r {
	case PriceTieredPriceModelTypeTiered:
		return true
	}
	return false
}

type PriceTieredPricePriceType string

const (
	PriceTieredPricePriceTypeUsagePrice PriceTieredPricePriceType = "usage_price"
	PriceTieredPricePriceTypeFixedPrice PriceTieredPricePriceType = "fixed_price"
)

func (r PriceTieredPricePriceType) IsKnown() bool {
	switch r {
	case PriceTieredPricePriceTypeUsagePrice, PriceTieredPricePriceTypeFixedPrice:
		return true
	}
	return false
}

type PriceTieredBPSPrice struct {
	ID                        string                     `json:"id,required"`
	BillableMetric            BillableMetricTiny         `json:"billable_metric,required,nullable"`
	BillingCycleConfiguration BillingCycleConfiguration  `json:"billing_cycle_configuration,required"`
	Cadence                   PriceTieredBPSPriceCadence `json:"cadence,required"`
	ConversionRate            float64                    `json:"conversion_rate,required,nullable"`
	ConversionRateConfig      PriceTieredBPSPrice        `json:"conversion_rate_config,required,nullable"`
	CreatedAt                 time.Time                  `json:"created_at,required" format:"date-time"`
	CreditAllocation          Allocation                 `json:"credit_allocation,required,nullable"`
	Currency                  string                     `json:"currency,required"`
	// Deprecated: deprecated
	Discount                    Discount                  `json:"discount,required,nullable"`
	ExternalPriceID             string                    `json:"external_price_id,required,nullable"`
	FixedPriceQuantity          float64                   `json:"fixed_price_quantity,required,nullable"`
	InvoicingCycleConfiguration BillingCycleConfiguration `json:"invoicing_cycle_configuration,required,nullable"`
	Item                        ItemSlim                  `json:"item,required"`
	// Deprecated: deprecated
	Maximum Maximum `json:"maximum,required,nullable"`
	// Deprecated: deprecated
	MaximumAmount string `json:"maximum_amount,required,nullable"`
	// User specified key-value pairs for the resource. If not present, this defaults
	// to an empty dictionary. Individual keys can be removed by setting the value to
	// `null`, and the entire metadata mapping can be cleared by setting `metadata` to
	// `null`.
	Metadata map[string]string `json:"metadata,required"`
	// Deprecated: deprecated
	Minimum Minimum `json:"minimum,required,nullable"`
	// Deprecated: deprecated
	MinimumAmount  string                       `json:"minimum_amount,required,nullable"`
	ModelType      PriceTieredBPSPriceModelType `json:"model_type,required"`
	Name           string                       `json:"name,required"`
	PlanPhaseOrder int64                        `json:"plan_phase_order,required,nullable"`
	PriceType      PriceTieredBPSPricePriceType `json:"price_type,required"`
	// The price id this price replaces. This price will take the place of the replaced
	// price in plan version migrations.
	ReplacesPriceID               string                        `json:"replaces_price_id,required,nullable"`
	TieredBPSConfig               TieredBPSConfig               `json:"tiered_bps_config,required"`
	DimensionalPriceConfiguration DimensionalPriceConfiguration `json:"dimensional_price_configuration,nullable"`
	JSON                          priceTieredBPSPriceJSON       `json:"-"`
}

// priceTieredBPSPriceJSON contains the JSON metadata for the struct
// [PriceTieredBPSPrice]
type priceTieredBPSPriceJSON struct {
	ID                            apijson.Field
	BillableMetric                apijson.Field
	BillingCycleConfiguration     apijson.Field
	Cadence                       apijson.Field
	ConversionRate                apijson.Field
	ConversionRateConfig          apijson.Field
	CreatedAt                     apijson.Field
	CreditAllocation              apijson.Field
	Currency                      apijson.Field
	Discount                      apijson.Field
	ExternalPriceID               apijson.Field
	FixedPriceQuantity            apijson.Field
	InvoicingCycleConfiguration   apijson.Field
	Item                          apijson.Field
	Maximum                       apijson.Field
	MaximumAmount                 apijson.Field
	Metadata                      apijson.Field
	Minimum                       apijson.Field
	MinimumAmount                 apijson.Field
	ModelType                     apijson.Field
	Name                          apijson.Field
	PlanPhaseOrder                apijson.Field
	PriceType                     apijson.Field
	ReplacesPriceID               apijson.Field
	TieredBPSConfig               apijson.Field
	DimensionalPriceConfiguration apijson.Field
	raw                           string
	ExtraFields                   map[string]apijson.Field
}

func (r *PriceTieredBPSPrice) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r priceTieredBPSPriceJSON) RawJSON() string {
	return r.raw
}

func (r PriceTieredBPSPrice) implementsPrice() {}

type PriceTieredBPSPriceCadence string

const (
	PriceTieredBPSPriceCadenceOneTime    PriceTieredBPSPriceCadence = "one_time"
	PriceTieredBPSPriceCadenceMonthly    PriceTieredBPSPriceCadence = "monthly"
	PriceTieredBPSPriceCadenceQuarterly  PriceTieredBPSPriceCadence = "quarterly"
	PriceTieredBPSPriceCadenceSemiAnnual PriceTieredBPSPriceCadence = "semi_annual"
	PriceTieredBPSPriceCadenceAnnual     PriceTieredBPSPriceCadence = "annual"
	PriceTieredBPSPriceCadenceCustom     PriceTieredBPSPriceCadence = "custom"
)

func (r PriceTieredBPSPriceCadence) IsKnown() bool {
	switch r {
	case PriceTieredBPSPriceCadenceOneTime, PriceTieredBPSPriceCadenceMonthly, PriceTieredBPSPriceCadenceQuarterly, PriceTieredBPSPriceCadenceSemiAnnual, PriceTieredBPSPriceCadenceAnnual, PriceTieredBPSPriceCadenceCustom:
		return true
	}
	return false
}

type PriceTieredBPSPriceModelType string

const (
	PriceTieredBPSPriceModelTypeTieredBPS PriceTieredBPSPriceModelType = "tiered_bps"
)

func (r PriceTieredBPSPriceModelType) IsKnown() bool {
	switch r {
	case PriceTieredBPSPriceModelTypeTieredBPS:
		return true
	}
	return false
}

type PriceTieredBPSPricePriceType string

const (
	PriceTieredBPSPricePriceTypeUsagePrice PriceTieredBPSPricePriceType = "usage_price"
	PriceTieredBPSPricePriceTypeFixedPrice PriceTieredBPSPricePriceType = "fixed_price"
)

func (r PriceTieredBPSPricePriceType) IsKnown() bool {
	switch r {
	case PriceTieredBPSPricePriceTypeUsagePrice, PriceTieredBPSPricePriceTypeFixedPrice:
		return true
	}
	return false
}

type PriceBPSPrice struct {
	ID                        string                    `json:"id,required"`
	BillableMetric            BillableMetricTiny        `json:"billable_metric,required,nullable"`
	BillingCycleConfiguration BillingCycleConfiguration `json:"billing_cycle_configuration,required"`
	BPSConfig                 BPSConfig                 `json:"bps_config,required"`
	Cadence                   PriceBPSPriceCadence      `json:"cadence,required"`
	ConversionRate            float64                   `json:"conversion_rate,required,nullable"`
	ConversionRateConfig      PriceBPSPrice             `json:"conversion_rate_config,required,nullable"`
	CreatedAt                 time.Time                 `json:"created_at,required" format:"date-time"`
	CreditAllocation          Allocation                `json:"credit_allocation,required,nullable"`
	Currency                  string                    `json:"currency,required"`
	// Deprecated: deprecated
	Discount                    Discount                  `json:"discount,required,nullable"`
	ExternalPriceID             string                    `json:"external_price_id,required,nullable"`
	FixedPriceQuantity          float64                   `json:"fixed_price_quantity,required,nullable"`
	InvoicingCycleConfiguration BillingCycleConfiguration `json:"invoicing_cycle_configuration,required,nullable"`
	Item                        ItemSlim                  `json:"item,required"`
	// Deprecated: deprecated
	Maximum Maximum `json:"maximum,required,nullable"`
	// Deprecated: deprecated
	MaximumAmount string `json:"maximum_amount,required,nullable"`
	// User specified key-value pairs for the resource. If not present, this defaults
	// to an empty dictionary. Individual keys can be removed by setting the value to
	// `null`, and the entire metadata mapping can be cleared by setting `metadata` to
	// `null`.
	Metadata map[string]string `json:"metadata,required"`
	// Deprecated: deprecated
	Minimum Minimum `json:"minimum,required,nullable"`
	// Deprecated: deprecated
	MinimumAmount  string                 `json:"minimum_amount,required,nullable"`
	ModelType      PriceBPSPriceModelType `json:"model_type,required"`
	Name           string                 `json:"name,required"`
	PlanPhaseOrder int64                  `json:"plan_phase_order,required,nullable"`
	PriceType      PriceBPSPricePriceType `json:"price_type,required"`
	// The price id this price replaces. This price will take the place of the replaced
	// price in plan version migrations.
	ReplacesPriceID               string                        `json:"replaces_price_id,required,nullable"`
	DimensionalPriceConfiguration DimensionalPriceConfiguration `json:"dimensional_price_configuration,nullable"`
	JSON                          priceBPSPriceJSON             `json:"-"`
}

// priceBPSPriceJSON contains the JSON metadata for the struct [PriceBPSPrice]
type priceBPSPriceJSON struct {
	ID                            apijson.Field
	BillableMetric                apijson.Field
	BillingCycleConfiguration     apijson.Field
	BPSConfig                     apijson.Field
	Cadence                       apijson.Field
	ConversionRate                apijson.Field
	ConversionRateConfig          apijson.Field
	CreatedAt                     apijson.Field
	CreditAllocation              apijson.Field
	Currency                      apijson.Field
	Discount                      apijson.Field
	ExternalPriceID               apijson.Field
	FixedPriceQuantity            apijson.Field
	InvoicingCycleConfiguration   apijson.Field
	Item                          apijson.Field
	Maximum                       apijson.Field
	MaximumAmount                 apijson.Field
	Metadata                      apijson.Field
	Minimum                       apijson.Field
	MinimumAmount                 apijson.Field
	ModelType                     apijson.Field
	Name                          apijson.Field
	PlanPhaseOrder                apijson.Field
	PriceType                     apijson.Field
	ReplacesPriceID               apijson.Field
	DimensionalPriceConfiguration apijson.Field
	raw                           string
	ExtraFields                   map[string]apijson.Field
}

func (r *PriceBPSPrice) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r priceBPSPriceJSON) RawJSON() string {
	return r.raw
}

func (r PriceBPSPrice) implementsPrice() {}

type PriceBPSPriceCadence string

const (
	PriceBPSPriceCadenceOneTime    PriceBPSPriceCadence = "one_time"
	PriceBPSPriceCadenceMonthly    PriceBPSPriceCadence = "monthly"
	PriceBPSPriceCadenceQuarterly  PriceBPSPriceCadence = "quarterly"
	PriceBPSPriceCadenceSemiAnnual PriceBPSPriceCadence = "semi_annual"
	PriceBPSPriceCadenceAnnual     PriceBPSPriceCadence = "annual"
	PriceBPSPriceCadenceCustom     PriceBPSPriceCadence = "custom"
)

func (r PriceBPSPriceCadence) IsKnown() bool {
	switch r {
	case PriceBPSPriceCadenceOneTime, PriceBPSPriceCadenceMonthly, PriceBPSPriceCadenceQuarterly, PriceBPSPriceCadenceSemiAnnual, PriceBPSPriceCadenceAnnual, PriceBPSPriceCadenceCustom:
		return true
	}
	return false
}

type PriceBPSPriceModelType string

const (
	PriceBPSPriceModelTypeBPS PriceBPSPriceModelType = "bps"
)

func (r PriceBPSPriceModelType) IsKnown() bool {
	switch r {
	case PriceBPSPriceModelTypeBPS:
		return true
	}
	return false
}

type PriceBPSPricePriceType string

const (
	PriceBPSPricePriceTypeUsagePrice PriceBPSPricePriceType = "usage_price"
	PriceBPSPricePriceTypeFixedPrice PriceBPSPricePriceType = "fixed_price"
)

func (r PriceBPSPricePriceType) IsKnown() bool {
	switch r {
	case PriceBPSPricePriceTypeUsagePrice, PriceBPSPricePriceTypeFixedPrice:
		return true
	}
	return false
}

type PriceBulkBPSPrice struct {
	ID                        string                    `json:"id,required"`
	BillableMetric            BillableMetricTiny        `json:"billable_metric,required,nullable"`
	BillingCycleConfiguration BillingCycleConfiguration `json:"billing_cycle_configuration,required"`
	BulkBPSConfig             BulkBPSConfig             `json:"bulk_bps_config,required"`
	Cadence                   PriceBulkBPSPriceCadence  `json:"cadence,required"`
	ConversionRate            float64                   `json:"conversion_rate,required,nullable"`
	ConversionRateConfig      PriceBulkBPSPrice         `json:"conversion_rate_config,required,nullable"`
	CreatedAt                 time.Time                 `json:"created_at,required" format:"date-time"`
	CreditAllocation          Allocation                `json:"credit_allocation,required,nullable"`
	Currency                  string                    `json:"currency,required"`
	// Deprecated: deprecated
	Discount                    Discount                  `json:"discount,required,nullable"`
	ExternalPriceID             string                    `json:"external_price_id,required,nullable"`
	FixedPriceQuantity          float64                   `json:"fixed_price_quantity,required,nullable"`
	InvoicingCycleConfiguration BillingCycleConfiguration `json:"invoicing_cycle_configuration,required,nullable"`
	Item                        ItemSlim                  `json:"item,required"`
	// Deprecated: deprecated
	Maximum Maximum `json:"maximum,required,nullable"`
	// Deprecated: deprecated
	MaximumAmount string `json:"maximum_amount,required,nullable"`
	// User specified key-value pairs for the resource. If not present, this defaults
	// to an empty dictionary. Individual keys can be removed by setting the value to
	// `null`, and the entire metadata mapping can be cleared by setting `metadata` to
	// `null`.
	Metadata map[string]string `json:"metadata,required"`
	// Deprecated: deprecated
	Minimum Minimum `json:"minimum,required,nullable"`
	// Deprecated: deprecated
	MinimumAmount  string                     `json:"minimum_amount,required,nullable"`
	ModelType      PriceBulkBPSPriceModelType `json:"model_type,required"`
	Name           string                     `json:"name,required"`
	PlanPhaseOrder int64                      `json:"plan_phase_order,required,nullable"`
	PriceType      PriceBulkBPSPricePriceType `json:"price_type,required"`
	// The price id this price replaces. This price will take the place of the replaced
	// price in plan version migrations.
	ReplacesPriceID               string                        `json:"replaces_price_id,required,nullable"`
	DimensionalPriceConfiguration DimensionalPriceConfiguration `json:"dimensional_price_configuration,nullable"`
	JSON                          priceBulkBPSPriceJSON         `json:"-"`
}

// priceBulkBPSPriceJSON contains the JSON metadata for the struct
// [PriceBulkBPSPrice]
type priceBulkBPSPriceJSON struct {
	ID                            apijson.Field
	BillableMetric                apijson.Field
	BillingCycleConfiguration     apijson.Field
	BulkBPSConfig                 apijson.Field
	Cadence                       apijson.Field
	ConversionRate                apijson.Field
	ConversionRateConfig          apijson.Field
	CreatedAt                     apijson.Field
	CreditAllocation              apijson.Field
	Currency                      apijson.Field
	Discount                      apijson.Field
	ExternalPriceID               apijson.Field
	FixedPriceQuantity            apijson.Field
	InvoicingCycleConfiguration   apijson.Field
	Item                          apijson.Field
	Maximum                       apijson.Field
	MaximumAmount                 apijson.Field
	Metadata                      apijson.Field
	Minimum                       apijson.Field
	MinimumAmount                 apijson.Field
	ModelType                     apijson.Field
	Name                          apijson.Field
	PlanPhaseOrder                apijson.Field
	PriceType                     apijson.Field
	ReplacesPriceID               apijson.Field
	DimensionalPriceConfiguration apijson.Field
	raw                           string
	ExtraFields                   map[string]apijson.Field
}

func (r *PriceBulkBPSPrice) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r priceBulkBPSPriceJSON) RawJSON() string {
	return r.raw
}

func (r PriceBulkBPSPrice) implementsPrice() {}

type PriceBulkBPSPriceCadence string

const (
	PriceBulkBPSPriceCadenceOneTime    PriceBulkBPSPriceCadence = "one_time"
	PriceBulkBPSPriceCadenceMonthly    PriceBulkBPSPriceCadence = "monthly"
	PriceBulkBPSPriceCadenceQuarterly  PriceBulkBPSPriceCadence = "quarterly"
	PriceBulkBPSPriceCadenceSemiAnnual PriceBulkBPSPriceCadence = "semi_annual"
	PriceBulkBPSPriceCadenceAnnual     PriceBulkBPSPriceCadence = "annual"
	PriceBulkBPSPriceCadenceCustom     PriceBulkBPSPriceCadence = "custom"
)

func (r PriceBulkBPSPriceCadence) IsKnown() bool {
	switch r {
	case PriceBulkBPSPriceCadenceOneTime, PriceBulkBPSPriceCadenceMonthly, PriceBulkBPSPriceCadenceQuarterly, PriceBulkBPSPriceCadenceSemiAnnual, PriceBulkBPSPriceCadenceAnnual, PriceBulkBPSPriceCadenceCustom:
		return true
	}
	return false
}

type PriceBulkBPSPriceModelType string

const (
	PriceBulkBPSPriceModelTypeBulkBPS PriceBulkBPSPriceModelType = "bulk_bps"
)

func (r PriceBulkBPSPriceModelType) IsKnown() bool {
	switch r {
	case PriceBulkBPSPriceModelTypeBulkBPS:
		return true
	}
	return false
}

type PriceBulkBPSPricePriceType string

const (
	PriceBulkBPSPricePriceTypeUsagePrice PriceBulkBPSPricePriceType = "usage_price"
	PriceBulkBPSPricePriceTypeFixedPrice PriceBulkBPSPricePriceType = "fixed_price"
)

func (r PriceBulkBPSPricePriceType) IsKnown() bool {
	switch r {
	case PriceBulkBPSPricePriceTypeUsagePrice, PriceBulkBPSPricePriceTypeFixedPrice:
		return true
	}
	return false
}

type PriceBulkPrice struct {
	ID                        string                    `json:"id,required"`
	BillableMetric            BillableMetricTiny        `json:"billable_metric,required,nullable"`
	BillingCycleConfiguration BillingCycleConfiguration `json:"billing_cycle_configuration,required"`
	BulkConfig                BulkConfig                `json:"bulk_config,required"`
	Cadence                   PriceBulkPriceCadence     `json:"cadence,required"`
	ConversionRate            float64                   `json:"conversion_rate,required,nullable"`
	ConversionRateConfig      PriceBulkPrice            `json:"conversion_rate_config,required,nullable"`
	CreatedAt                 time.Time                 `json:"created_at,required" format:"date-time"`
	CreditAllocation          Allocation                `json:"credit_allocation,required,nullable"`
	Currency                  string                    `json:"currency,required"`
	// Deprecated: deprecated
	Discount                    Discount                  `json:"discount,required,nullable"`
	ExternalPriceID             string                    `json:"external_price_id,required,nullable"`
	FixedPriceQuantity          float64                   `json:"fixed_price_quantity,required,nullable"`
	InvoicingCycleConfiguration BillingCycleConfiguration `json:"invoicing_cycle_configuration,required,nullable"`
	Item                        ItemSlim                  `json:"item,required"`
	// Deprecated: deprecated
	Maximum Maximum `json:"maximum,required,nullable"`
	// Deprecated: deprecated
	MaximumAmount string `json:"maximum_amount,required,nullable"`
	// User specified key-value pairs for the resource. If not present, this defaults
	// to an empty dictionary. Individual keys can be removed by setting the value to
	// `null`, and the entire metadata mapping can be cleared by setting `metadata` to
	// `null`.
	Metadata map[string]string `json:"metadata,required"`
	// Deprecated: deprecated
	Minimum Minimum `json:"minimum,required,nullable"`
	// Deprecated: deprecated
	MinimumAmount  string                  `json:"minimum_amount,required,nullable"`
	ModelType      PriceBulkPriceModelType `json:"model_type,required"`
	Name           string                  `json:"name,required"`
	PlanPhaseOrder int64                   `json:"plan_phase_order,required,nullable"`
	PriceType      PriceBulkPricePriceType `json:"price_type,required"`
	// The price id this price replaces. This price will take the place of the replaced
	// price in plan version migrations.
	ReplacesPriceID               string                        `json:"replaces_price_id,required,nullable"`
	DimensionalPriceConfiguration DimensionalPriceConfiguration `json:"dimensional_price_configuration,nullable"`
	JSON                          priceBulkPriceJSON            `json:"-"`
}

// priceBulkPriceJSON contains the JSON metadata for the struct [PriceBulkPrice]
type priceBulkPriceJSON struct {
	ID                            apijson.Field
	BillableMetric                apijson.Field
	BillingCycleConfiguration     apijson.Field
	BulkConfig                    apijson.Field
	Cadence                       apijson.Field
	ConversionRate                apijson.Field
	ConversionRateConfig          apijson.Field
	CreatedAt                     apijson.Field
	CreditAllocation              apijson.Field
	Currency                      apijson.Field
	Discount                      apijson.Field
	ExternalPriceID               apijson.Field
	FixedPriceQuantity            apijson.Field
	InvoicingCycleConfiguration   apijson.Field
	Item                          apijson.Field
	Maximum                       apijson.Field
	MaximumAmount                 apijson.Field
	Metadata                      apijson.Field
	Minimum                       apijson.Field
	MinimumAmount                 apijson.Field
	ModelType                     apijson.Field
	Name                          apijson.Field
	PlanPhaseOrder                apijson.Field
	PriceType                     apijson.Field
	ReplacesPriceID               apijson.Field
	DimensionalPriceConfiguration apijson.Field
	raw                           string
	ExtraFields                   map[string]apijson.Field
}

func (r *PriceBulkPrice) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r priceBulkPriceJSON) RawJSON() string {
	return r.raw
}

func (r PriceBulkPrice) implementsPrice() {}

type PriceBulkPriceCadence string

const (
	PriceBulkPriceCadenceOneTime    PriceBulkPriceCadence = "one_time"
	PriceBulkPriceCadenceMonthly    PriceBulkPriceCadence = "monthly"
	PriceBulkPriceCadenceQuarterly  PriceBulkPriceCadence = "quarterly"
	PriceBulkPriceCadenceSemiAnnual PriceBulkPriceCadence = "semi_annual"
	PriceBulkPriceCadenceAnnual     PriceBulkPriceCadence = "annual"
	PriceBulkPriceCadenceCustom     PriceBulkPriceCadence = "custom"
)

func (r PriceBulkPriceCadence) IsKnown() bool {
	switch r {
	case PriceBulkPriceCadenceOneTime, PriceBulkPriceCadenceMonthly, PriceBulkPriceCadenceQuarterly, PriceBulkPriceCadenceSemiAnnual, PriceBulkPriceCadenceAnnual, PriceBulkPriceCadenceCustom:
		return true
	}
	return false
}

type PriceBulkPriceModelType string

const (
	PriceBulkPriceModelTypeBulk PriceBulkPriceModelType = "bulk"
)

func (r PriceBulkPriceModelType) IsKnown() bool {
	switch r {
	case PriceBulkPriceModelTypeBulk:
		return true
	}
	return false
}

type PriceBulkPricePriceType string

const (
	PriceBulkPricePriceTypeUsagePrice PriceBulkPricePriceType = "usage_price"
	PriceBulkPricePriceTypeFixedPrice PriceBulkPricePriceType = "fixed_price"
)

func (r PriceBulkPricePriceType) IsKnown() bool {
	switch r {
	case PriceBulkPricePriceTypeUsagePrice, PriceBulkPricePriceTypeFixedPrice:
		return true
	}
	return false
}

type PriceThresholdTotalAmountPrice struct {
	ID                        string                                `json:"id,required"`
	BillableMetric            BillableMetricTiny                    `json:"billable_metric,required,nullable"`
	BillingCycleConfiguration BillingCycleConfiguration             `json:"billing_cycle_configuration,required"`
	Cadence                   PriceThresholdTotalAmountPriceCadence `json:"cadence,required"`
	ConversionRate            float64                               `json:"conversion_rate,required,nullable"`
	ConversionRateConfig      PriceThresholdTotalAmountPrice        `json:"conversion_rate_config,required,nullable"`
	CreatedAt                 time.Time                             `json:"created_at,required" format:"date-time"`
	CreditAllocation          Allocation                            `json:"credit_allocation,required,nullable"`
	Currency                  string                                `json:"currency,required"`
	// Deprecated: deprecated
	Discount                    Discount                  `json:"discount,required,nullable"`
	ExternalPriceID             string                    `json:"external_price_id,required,nullable"`
	FixedPriceQuantity          float64                   `json:"fixed_price_quantity,required,nullable"`
	InvoicingCycleConfiguration BillingCycleConfiguration `json:"invoicing_cycle_configuration,required,nullable"`
	Item                        ItemSlim                  `json:"item,required"`
	// Deprecated: deprecated
	Maximum Maximum `json:"maximum,required,nullable"`
	// Deprecated: deprecated
	MaximumAmount string `json:"maximum_amount,required,nullable"`
	// User specified key-value pairs for the resource. If not present, this defaults
	// to an empty dictionary. Individual keys can be removed by setting the value to
	// `null`, and the entire metadata mapping can be cleared by setting `metadata` to
	// `null`.
	Metadata map[string]string `json:"metadata,required"`
	// Deprecated: deprecated
	Minimum Minimum `json:"minimum,required,nullable"`
	// Deprecated: deprecated
	MinimumAmount  string                                  `json:"minimum_amount,required,nullable"`
	ModelType      PriceThresholdTotalAmountPriceModelType `json:"model_type,required"`
	Name           string                                  `json:"name,required"`
	PlanPhaseOrder int64                                   `json:"plan_phase_order,required,nullable"`
	PriceType      PriceThresholdTotalAmountPricePriceType `json:"price_type,required"`
	// The price id this price replaces. This price will take the place of the replaced
	// price in plan version migrations.
	ReplacesPriceID               string                             `json:"replaces_price_id,required,nullable"`
	ThresholdTotalAmountConfig    map[string]interface{}             `json:"threshold_total_amount_config,required"`
	DimensionalPriceConfiguration DimensionalPriceConfiguration      `json:"dimensional_price_configuration,nullable"`
	JSON                          priceThresholdTotalAmountPriceJSON `json:"-"`
}

// priceThresholdTotalAmountPriceJSON contains the JSON metadata for the struct
// [PriceThresholdTotalAmountPrice]
type priceThresholdTotalAmountPriceJSON struct {
	ID                            apijson.Field
	BillableMetric                apijson.Field
	BillingCycleConfiguration     apijson.Field
	Cadence                       apijson.Field
	ConversionRate                apijson.Field
	ConversionRateConfig          apijson.Field
	CreatedAt                     apijson.Field
	CreditAllocation              apijson.Field
	Currency                      apijson.Field
	Discount                      apijson.Field
	ExternalPriceID               apijson.Field
	FixedPriceQuantity            apijson.Field
	InvoicingCycleConfiguration   apijson.Field
	Item                          apijson.Field
	Maximum                       apijson.Field
	MaximumAmount                 apijson.Field
	Metadata                      apijson.Field
	Minimum                       apijson.Field
	MinimumAmount                 apijson.Field
	ModelType                     apijson.Field
	Name                          apijson.Field
	PlanPhaseOrder                apijson.Field
	PriceType                     apijson.Field
	ReplacesPriceID               apijson.Field
	ThresholdTotalAmountConfig    apijson.Field
	DimensionalPriceConfiguration apijson.Field
	raw                           string
	ExtraFields                   map[string]apijson.Field
}

func (r *PriceThresholdTotalAmountPrice) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r priceThresholdTotalAmountPriceJSON) RawJSON() string {
	return r.raw
}

func (r PriceThresholdTotalAmountPrice) implementsPrice() {}

type PriceThresholdTotalAmountPriceCadence string

const (
	PriceThresholdTotalAmountPriceCadenceOneTime    PriceThresholdTotalAmountPriceCadence = "one_time"
	PriceThresholdTotalAmountPriceCadenceMonthly    PriceThresholdTotalAmountPriceCadence = "monthly"
	PriceThresholdTotalAmountPriceCadenceQuarterly  PriceThresholdTotalAmountPriceCadence = "quarterly"
	PriceThresholdTotalAmountPriceCadenceSemiAnnual PriceThresholdTotalAmountPriceCadence = "semi_annual"
	PriceThresholdTotalAmountPriceCadenceAnnual     PriceThresholdTotalAmountPriceCadence = "annual"
	PriceThresholdTotalAmountPriceCadenceCustom     PriceThresholdTotalAmountPriceCadence = "custom"
)

func (r PriceThresholdTotalAmountPriceCadence) IsKnown() bool {
	switch r {
	case PriceThresholdTotalAmountPriceCadenceOneTime, PriceThresholdTotalAmountPriceCadenceMonthly, PriceThresholdTotalAmountPriceCadenceQuarterly, PriceThresholdTotalAmountPriceCadenceSemiAnnual, PriceThresholdTotalAmountPriceCadenceAnnual, PriceThresholdTotalAmountPriceCadenceCustom:
		return true
	}
	return false
}

type PriceThresholdTotalAmountPriceModelType string

const (
	PriceThresholdTotalAmountPriceModelTypeThresholdTotalAmount PriceThresholdTotalAmountPriceModelType = "threshold_total_amount"
)

func (r PriceThresholdTotalAmountPriceModelType) IsKnown() bool {
	switch r {
	case PriceThresholdTotalAmountPriceModelTypeThresholdTotalAmount:
		return true
	}
	return false
}

type PriceThresholdTotalAmountPricePriceType string

const (
	PriceThresholdTotalAmountPricePriceTypeUsagePrice PriceThresholdTotalAmountPricePriceType = "usage_price"
	PriceThresholdTotalAmountPricePriceTypeFixedPrice PriceThresholdTotalAmountPricePriceType = "fixed_price"
)

func (r PriceThresholdTotalAmountPricePriceType) IsKnown() bool {
	switch r {
	case PriceThresholdTotalAmountPricePriceTypeUsagePrice, PriceThresholdTotalAmountPricePriceTypeFixedPrice:
		return true
	}
	return false
}

type PriceTieredPackagePrice struct {
	ID                        string                         `json:"id,required"`
	BillableMetric            BillableMetricTiny             `json:"billable_metric,required,nullable"`
	BillingCycleConfiguration BillingCycleConfiguration      `json:"billing_cycle_configuration,required"`
	Cadence                   PriceTieredPackagePriceCadence `json:"cadence,required"`
	ConversionRate            float64                        `json:"conversion_rate,required,nullable"`
	ConversionRateConfig      PriceTieredPackagePrice        `json:"conversion_rate_config,required,nullable"`
	CreatedAt                 time.Time                      `json:"created_at,required" format:"date-time"`
	CreditAllocation          Allocation                     `json:"credit_allocation,required,nullable"`
	Currency                  string                         `json:"currency,required"`
	// Deprecated: deprecated
	Discount                    Discount                  `json:"discount,required,nullable"`
	ExternalPriceID             string                    `json:"external_price_id,required,nullable"`
	FixedPriceQuantity          float64                   `json:"fixed_price_quantity,required,nullable"`
	InvoicingCycleConfiguration BillingCycleConfiguration `json:"invoicing_cycle_configuration,required,nullable"`
	Item                        ItemSlim                  `json:"item,required"`
	// Deprecated: deprecated
	Maximum Maximum `json:"maximum,required,nullable"`
	// Deprecated: deprecated
	MaximumAmount string `json:"maximum_amount,required,nullable"`
	// User specified key-value pairs for the resource. If not present, this defaults
	// to an empty dictionary. Individual keys can be removed by setting the value to
	// `null`, and the entire metadata mapping can be cleared by setting `metadata` to
	// `null`.
	Metadata map[string]string `json:"metadata,required"`
	// Deprecated: deprecated
	Minimum Minimum `json:"minimum,required,nullable"`
	// Deprecated: deprecated
	MinimumAmount  string                           `json:"minimum_amount,required,nullable"`
	ModelType      PriceTieredPackagePriceModelType `json:"model_type,required"`
	Name           string                           `json:"name,required"`
	PlanPhaseOrder int64                            `json:"plan_phase_order,required,nullable"`
	PriceType      PriceTieredPackagePricePriceType `json:"price_type,required"`
	// The price id this price replaces. This price will take the place of the replaced
	// price in plan version migrations.
	ReplacesPriceID               string                        `json:"replaces_price_id,required,nullable"`
	TieredPackageConfig           map[string]interface{}        `json:"tiered_package_config,required"`
	DimensionalPriceConfiguration DimensionalPriceConfiguration `json:"dimensional_price_configuration,nullable"`
	JSON                          priceTieredPackagePriceJSON   `json:"-"`
}

// priceTieredPackagePriceJSON contains the JSON metadata for the struct
// [PriceTieredPackagePrice]
type priceTieredPackagePriceJSON struct {
	ID                            apijson.Field
	BillableMetric                apijson.Field
	BillingCycleConfiguration     apijson.Field
	Cadence                       apijson.Field
	ConversionRate                apijson.Field
	ConversionRateConfig          apijson.Field
	CreatedAt                     apijson.Field
	CreditAllocation              apijson.Field
	Currency                      apijson.Field
	Discount                      apijson.Field
	ExternalPriceID               apijson.Field
	FixedPriceQuantity            apijson.Field
	InvoicingCycleConfiguration   apijson.Field
	Item                          apijson.Field
	Maximum                       apijson.Field
	MaximumAmount                 apijson.Field
	Metadata                      apijson.Field
	Minimum                       apijson.Field
	MinimumAmount                 apijson.Field
	ModelType                     apijson.Field
	Name                          apijson.Field
	PlanPhaseOrder                apijson.Field
	PriceType                     apijson.Field
	ReplacesPriceID               apijson.Field
	TieredPackageConfig           apijson.Field
	DimensionalPriceConfiguration apijson.Field
	raw                           string
	ExtraFields                   map[string]apijson.Field
}

func (r *PriceTieredPackagePrice) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r priceTieredPackagePriceJSON) RawJSON() string {
	return r.raw
}

func (r PriceTieredPackagePrice) implementsPrice() {}

type PriceTieredPackagePriceCadence string

const (
	PriceTieredPackagePriceCadenceOneTime    PriceTieredPackagePriceCadence = "one_time"
	PriceTieredPackagePriceCadenceMonthly    PriceTieredPackagePriceCadence = "monthly"
	PriceTieredPackagePriceCadenceQuarterly  PriceTieredPackagePriceCadence = "quarterly"
	PriceTieredPackagePriceCadenceSemiAnnual PriceTieredPackagePriceCadence = "semi_annual"
	PriceTieredPackagePriceCadenceAnnual     PriceTieredPackagePriceCadence = "annual"
	PriceTieredPackagePriceCadenceCustom     PriceTieredPackagePriceCadence = "custom"
)

func (r PriceTieredPackagePriceCadence) IsKnown() bool {
	switch r {
	case PriceTieredPackagePriceCadenceOneTime, PriceTieredPackagePriceCadenceMonthly, PriceTieredPackagePriceCadenceQuarterly, PriceTieredPackagePriceCadenceSemiAnnual, PriceTieredPackagePriceCadenceAnnual, PriceTieredPackagePriceCadenceCustom:
		return true
	}
	return false
}

type PriceTieredPackagePriceModelType string

const (
	PriceTieredPackagePriceModelTypeTieredPackage PriceTieredPackagePriceModelType = "tiered_package"
)

func (r PriceTieredPackagePriceModelType) IsKnown() bool {
	switch r {
	case PriceTieredPackagePriceModelTypeTieredPackage:
		return true
	}
	return false
}

type PriceTieredPackagePricePriceType string

const (
	PriceTieredPackagePricePriceTypeUsagePrice PriceTieredPackagePricePriceType = "usage_price"
	PriceTieredPackagePricePriceTypeFixedPrice PriceTieredPackagePricePriceType = "fixed_price"
)

func (r PriceTieredPackagePricePriceType) IsKnown() bool {
	switch r {
	case PriceTieredPackagePricePriceTypeUsagePrice, PriceTieredPackagePricePriceTypeFixedPrice:
		return true
	}
	return false
}

type PriceGroupedTieredPrice struct {
	ID                        string                         `json:"id,required"`
	BillableMetric            BillableMetricTiny             `json:"billable_metric,required,nullable"`
	BillingCycleConfiguration BillingCycleConfiguration      `json:"billing_cycle_configuration,required"`
	Cadence                   PriceGroupedTieredPriceCadence `json:"cadence,required"`
	ConversionRate            float64                        `json:"conversion_rate,required,nullable"`
	ConversionRateConfig      PriceGroupedTieredPrice        `json:"conversion_rate_config,required,nullable"`
	CreatedAt                 time.Time                      `json:"created_at,required" format:"date-time"`
	CreditAllocation          Allocation                     `json:"credit_allocation,required,nullable"`
	Currency                  string                         `json:"currency,required"`
	// Deprecated: deprecated
	Discount                    Discount                  `json:"discount,required,nullable"`
	ExternalPriceID             string                    `json:"external_price_id,required,nullable"`
	FixedPriceQuantity          float64                   `json:"fixed_price_quantity,required,nullable"`
	GroupedTieredConfig         map[string]interface{}    `json:"grouped_tiered_config,required"`
	InvoicingCycleConfiguration BillingCycleConfiguration `json:"invoicing_cycle_configuration,required,nullable"`
	Item                        ItemSlim                  `json:"item,required"`
	// Deprecated: deprecated
	Maximum Maximum `json:"maximum,required,nullable"`
	// Deprecated: deprecated
	MaximumAmount string `json:"maximum_amount,required,nullable"`
	// User specified key-value pairs for the resource. If not present, this defaults
	// to an empty dictionary. Individual keys can be removed by setting the value to
	// `null`, and the entire metadata mapping can be cleared by setting `metadata` to
	// `null`.
	Metadata map[string]string `json:"metadata,required"`
	// Deprecated: deprecated
	Minimum Minimum `json:"minimum,required,nullable"`
	// Deprecated: deprecated
	MinimumAmount  string                           `json:"minimum_amount,required,nullable"`
	ModelType      PriceGroupedTieredPriceModelType `json:"model_type,required"`
	Name           string                           `json:"name,required"`
	PlanPhaseOrder int64                            `json:"plan_phase_order,required,nullable"`
	PriceType      PriceGroupedTieredPricePriceType `json:"price_type,required"`
	// The price id this price replaces. This price will take the place of the replaced
	// price in plan version migrations.
	ReplacesPriceID               string                        `json:"replaces_price_id,required,nullable"`
	DimensionalPriceConfiguration DimensionalPriceConfiguration `json:"dimensional_price_configuration,nullable"`
	JSON                          priceGroupedTieredPriceJSON   `json:"-"`
}

// priceGroupedTieredPriceJSON contains the JSON metadata for the struct
// [PriceGroupedTieredPrice]
type priceGroupedTieredPriceJSON struct {
	ID                            apijson.Field
	BillableMetric                apijson.Field
	BillingCycleConfiguration     apijson.Field
	Cadence                       apijson.Field
	ConversionRate                apijson.Field
	ConversionRateConfig          apijson.Field
	CreatedAt                     apijson.Field
	CreditAllocation              apijson.Field
	Currency                      apijson.Field
	Discount                      apijson.Field
	ExternalPriceID               apijson.Field
	FixedPriceQuantity            apijson.Field
	GroupedTieredConfig           apijson.Field
	InvoicingCycleConfiguration   apijson.Field
	Item                          apijson.Field
	Maximum                       apijson.Field
	MaximumAmount                 apijson.Field
	Metadata                      apijson.Field
	Minimum                       apijson.Field
	MinimumAmount                 apijson.Field
	ModelType                     apijson.Field
	Name                          apijson.Field
	PlanPhaseOrder                apijson.Field
	PriceType                     apijson.Field
	ReplacesPriceID               apijson.Field
	DimensionalPriceConfiguration apijson.Field
	raw                           string
	ExtraFields                   map[string]apijson.Field
}

func (r *PriceGroupedTieredPrice) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r priceGroupedTieredPriceJSON) RawJSON() string {
	return r.raw
}

func (r PriceGroupedTieredPrice) implementsPrice() {}

type PriceGroupedTieredPriceCadence string

const (
	PriceGroupedTieredPriceCadenceOneTime    PriceGroupedTieredPriceCadence = "one_time"
	PriceGroupedTieredPriceCadenceMonthly    PriceGroupedTieredPriceCadence = "monthly"
	PriceGroupedTieredPriceCadenceQuarterly  PriceGroupedTieredPriceCadence = "quarterly"
	PriceGroupedTieredPriceCadenceSemiAnnual PriceGroupedTieredPriceCadence = "semi_annual"
	PriceGroupedTieredPriceCadenceAnnual     PriceGroupedTieredPriceCadence = "annual"
	PriceGroupedTieredPriceCadenceCustom     PriceGroupedTieredPriceCadence = "custom"
)

func (r PriceGroupedTieredPriceCadence) IsKnown() bool {
	switch r {
	case PriceGroupedTieredPriceCadenceOneTime, PriceGroupedTieredPriceCadenceMonthly, PriceGroupedTieredPriceCadenceQuarterly, PriceGroupedTieredPriceCadenceSemiAnnual, PriceGroupedTieredPriceCadenceAnnual, PriceGroupedTieredPriceCadenceCustom:
		return true
	}
	return false
}

type PriceGroupedTieredPriceModelType string

const (
	PriceGroupedTieredPriceModelTypeGroupedTiered PriceGroupedTieredPriceModelType = "grouped_tiered"
)

func (r PriceGroupedTieredPriceModelType) IsKnown() bool {
	switch r {
	case PriceGroupedTieredPriceModelTypeGroupedTiered:
		return true
	}
	return false
}

type PriceGroupedTieredPricePriceType string

const (
	PriceGroupedTieredPricePriceTypeUsagePrice PriceGroupedTieredPricePriceType = "usage_price"
	PriceGroupedTieredPricePriceTypeFixedPrice PriceGroupedTieredPricePriceType = "fixed_price"
)

func (r PriceGroupedTieredPricePriceType) IsKnown() bool {
	switch r {
	case PriceGroupedTieredPricePriceTypeUsagePrice, PriceGroupedTieredPricePriceTypeFixedPrice:
		return true
	}
	return false
}

type PriceTieredWithMinimumPrice struct {
	ID                        string                             `json:"id,required"`
	BillableMetric            BillableMetricTiny                 `json:"billable_metric,required,nullable"`
	BillingCycleConfiguration BillingCycleConfiguration          `json:"billing_cycle_configuration,required"`
	Cadence                   PriceTieredWithMinimumPriceCadence `json:"cadence,required"`
	ConversionRate            float64                            `json:"conversion_rate,required,nullable"`
	ConversionRateConfig      PriceTieredWithMinimumPrice        `json:"conversion_rate_config,required,nullable"`
	CreatedAt                 time.Time                          `json:"created_at,required" format:"date-time"`
	CreditAllocation          Allocation                         `json:"credit_allocation,required,nullable"`
	Currency                  string                             `json:"currency,required"`
	// Deprecated: deprecated
	Discount                    Discount                  `json:"discount,required,nullable"`
	ExternalPriceID             string                    `json:"external_price_id,required,nullable"`
	FixedPriceQuantity          float64                   `json:"fixed_price_quantity,required,nullable"`
	InvoicingCycleConfiguration BillingCycleConfiguration `json:"invoicing_cycle_configuration,required,nullable"`
	Item                        ItemSlim                  `json:"item,required"`
	// Deprecated: deprecated
	Maximum Maximum `json:"maximum,required,nullable"`
	// Deprecated: deprecated
	MaximumAmount string `json:"maximum_amount,required,nullable"`
	// User specified key-value pairs for the resource. If not present, this defaults
	// to an empty dictionary. Individual keys can be removed by setting the value to
	// `null`, and the entire metadata mapping can be cleared by setting `metadata` to
	// `null`.
	Metadata map[string]string `json:"metadata,required"`
	// Deprecated: deprecated
	Minimum Minimum `json:"minimum,required,nullable"`
	// Deprecated: deprecated
	MinimumAmount  string                               `json:"minimum_amount,required,nullable"`
	ModelType      PriceTieredWithMinimumPriceModelType `json:"model_type,required"`
	Name           string                               `json:"name,required"`
	PlanPhaseOrder int64                                `json:"plan_phase_order,required,nullable"`
	PriceType      PriceTieredWithMinimumPricePriceType `json:"price_type,required"`
	// The price id this price replaces. This price will take the place of the replaced
	// price in plan version migrations.
	ReplacesPriceID               string                          `json:"replaces_price_id,required,nullable"`
	TieredWithMinimumConfig       map[string]interface{}          `json:"tiered_with_minimum_config,required"`
	DimensionalPriceConfiguration DimensionalPriceConfiguration   `json:"dimensional_price_configuration,nullable"`
	JSON                          priceTieredWithMinimumPriceJSON `json:"-"`
}

// priceTieredWithMinimumPriceJSON contains the JSON metadata for the struct
// [PriceTieredWithMinimumPrice]
type priceTieredWithMinimumPriceJSON struct {
	ID                            apijson.Field
	BillableMetric                apijson.Field
	BillingCycleConfiguration     apijson.Field
	Cadence                       apijson.Field
	ConversionRate                apijson.Field
	ConversionRateConfig          apijson.Field
	CreatedAt                     apijson.Field
	CreditAllocation              apijson.Field
	Currency                      apijson.Field
	Discount                      apijson.Field
	ExternalPriceID               apijson.Field
	FixedPriceQuantity            apijson.Field
	InvoicingCycleConfiguration   apijson.Field
	Item                          apijson.Field
	Maximum                       apijson.Field
	MaximumAmount                 apijson.Field
	Metadata                      apijson.Field
	Minimum                       apijson.Field
	MinimumAmount                 apijson.Field
	ModelType                     apijson.Field
	Name                          apijson.Field
	PlanPhaseOrder                apijson.Field
	PriceType                     apijson.Field
	ReplacesPriceID               apijson.Field
	TieredWithMinimumConfig       apijson.Field
	DimensionalPriceConfiguration apijson.Field
	raw                           string
	ExtraFields                   map[string]apijson.Field
}

func (r *PriceTieredWithMinimumPrice) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r priceTieredWithMinimumPriceJSON) RawJSON() string {
	return r.raw
}

func (r PriceTieredWithMinimumPrice) implementsPrice() {}

type PriceTieredWithMinimumPriceCadence string

const (
	PriceTieredWithMinimumPriceCadenceOneTime    PriceTieredWithMinimumPriceCadence = "one_time"
	PriceTieredWithMinimumPriceCadenceMonthly    PriceTieredWithMinimumPriceCadence = "monthly"
	PriceTieredWithMinimumPriceCadenceQuarterly  PriceTieredWithMinimumPriceCadence = "quarterly"
	PriceTieredWithMinimumPriceCadenceSemiAnnual PriceTieredWithMinimumPriceCadence = "semi_annual"
	PriceTieredWithMinimumPriceCadenceAnnual     PriceTieredWithMinimumPriceCadence = "annual"
	PriceTieredWithMinimumPriceCadenceCustom     PriceTieredWithMinimumPriceCadence = "custom"
)

func (r PriceTieredWithMinimumPriceCadence) IsKnown() bool {
	switch r {
	case PriceTieredWithMinimumPriceCadenceOneTime, PriceTieredWithMinimumPriceCadenceMonthly, PriceTieredWithMinimumPriceCadenceQuarterly, PriceTieredWithMinimumPriceCadenceSemiAnnual, PriceTieredWithMinimumPriceCadenceAnnual, PriceTieredWithMinimumPriceCadenceCustom:
		return true
	}
	return false
}

type PriceTieredWithMinimumPriceModelType string

const (
	PriceTieredWithMinimumPriceModelTypeTieredWithMinimum PriceTieredWithMinimumPriceModelType = "tiered_with_minimum"
)

func (r PriceTieredWithMinimumPriceModelType) IsKnown() bool {
	switch r {
	case PriceTieredWithMinimumPriceModelTypeTieredWithMinimum:
		return true
	}
	return false
}

type PriceTieredWithMinimumPricePriceType string

const (
	PriceTieredWithMinimumPricePriceTypeUsagePrice PriceTieredWithMinimumPricePriceType = "usage_price"
	PriceTieredWithMinimumPricePriceTypeFixedPrice PriceTieredWithMinimumPricePriceType = "fixed_price"
)

func (r PriceTieredWithMinimumPricePriceType) IsKnown() bool {
	switch r {
	case PriceTieredWithMinimumPricePriceTypeUsagePrice, PriceTieredWithMinimumPricePriceTypeFixedPrice:
		return true
	}
	return false
}

type PriceTieredPackageWithMinimumPrice struct {
	ID                        string                                    `json:"id,required"`
	BillableMetric            BillableMetricTiny                        `json:"billable_metric,required,nullable"`
	BillingCycleConfiguration BillingCycleConfiguration                 `json:"billing_cycle_configuration,required"`
	Cadence                   PriceTieredPackageWithMinimumPriceCadence `json:"cadence,required"`
	ConversionRate            float64                                   `json:"conversion_rate,required,nullable"`
	ConversionRateConfig      PriceTieredPackageWithMinimumPrice        `json:"conversion_rate_config,required,nullable"`
	CreatedAt                 time.Time                                 `json:"created_at,required" format:"date-time"`
	CreditAllocation          Allocation                                `json:"credit_allocation,required,nullable"`
	Currency                  string                                    `json:"currency,required"`
	// Deprecated: deprecated
	Discount                    Discount                  `json:"discount,required,nullable"`
	ExternalPriceID             string                    `json:"external_price_id,required,nullable"`
	FixedPriceQuantity          float64                   `json:"fixed_price_quantity,required,nullable"`
	InvoicingCycleConfiguration BillingCycleConfiguration `json:"invoicing_cycle_configuration,required,nullable"`
	Item                        ItemSlim                  `json:"item,required"`
	// Deprecated: deprecated
	Maximum Maximum `json:"maximum,required,nullable"`
	// Deprecated: deprecated
	MaximumAmount string `json:"maximum_amount,required,nullable"`
	// User specified key-value pairs for the resource. If not present, this defaults
	// to an empty dictionary. Individual keys can be removed by setting the value to
	// `null`, and the entire metadata mapping can be cleared by setting `metadata` to
	// `null`.
	Metadata map[string]string `json:"metadata,required"`
	// Deprecated: deprecated
	Minimum Minimum `json:"minimum,required,nullable"`
	// Deprecated: deprecated
	MinimumAmount  string                                      `json:"minimum_amount,required,nullable"`
	ModelType      PriceTieredPackageWithMinimumPriceModelType `json:"model_type,required"`
	Name           string                                      `json:"name,required"`
	PlanPhaseOrder int64                                       `json:"plan_phase_order,required,nullable"`
	PriceType      PriceTieredPackageWithMinimumPricePriceType `json:"price_type,required"`
	// The price id this price replaces. This price will take the place of the replaced
	// price in plan version migrations.
	ReplacesPriceID                string                                 `json:"replaces_price_id,required,nullable"`
	TieredPackageWithMinimumConfig map[string]interface{}                 `json:"tiered_package_with_minimum_config,required"`
	DimensionalPriceConfiguration  DimensionalPriceConfiguration          `json:"dimensional_price_configuration,nullable"`
	JSON                           priceTieredPackageWithMinimumPriceJSON `json:"-"`
}

// priceTieredPackageWithMinimumPriceJSON contains the JSON metadata for the struct
// [PriceTieredPackageWithMinimumPrice]
type priceTieredPackageWithMinimumPriceJSON struct {
	ID                             apijson.Field
	BillableMetric                 apijson.Field
	BillingCycleConfiguration      apijson.Field
	Cadence                        apijson.Field
	ConversionRate                 apijson.Field
	ConversionRateConfig           apijson.Field
	CreatedAt                      apijson.Field
	CreditAllocation               apijson.Field
	Currency                       apijson.Field
	Discount                       apijson.Field
	ExternalPriceID                apijson.Field
	FixedPriceQuantity             apijson.Field
	InvoicingCycleConfiguration    apijson.Field
	Item                           apijson.Field
	Maximum                        apijson.Field
	MaximumAmount                  apijson.Field
	Metadata                       apijson.Field
	Minimum                        apijson.Field
	MinimumAmount                  apijson.Field
	ModelType                      apijson.Field
	Name                           apijson.Field
	PlanPhaseOrder                 apijson.Field
	PriceType                      apijson.Field
	ReplacesPriceID                apijson.Field
	TieredPackageWithMinimumConfig apijson.Field
	DimensionalPriceConfiguration  apijson.Field
	raw                            string
	ExtraFields                    map[string]apijson.Field
}

func (r *PriceTieredPackageWithMinimumPrice) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r priceTieredPackageWithMinimumPriceJSON) RawJSON() string {
	return r.raw
}

func (r PriceTieredPackageWithMinimumPrice) implementsPrice() {}

type PriceTieredPackageWithMinimumPriceCadence string

const (
	PriceTieredPackageWithMinimumPriceCadenceOneTime    PriceTieredPackageWithMinimumPriceCadence = "one_time"
	PriceTieredPackageWithMinimumPriceCadenceMonthly    PriceTieredPackageWithMinimumPriceCadence = "monthly"
	PriceTieredPackageWithMinimumPriceCadenceQuarterly  PriceTieredPackageWithMinimumPriceCadence = "quarterly"
	PriceTieredPackageWithMinimumPriceCadenceSemiAnnual PriceTieredPackageWithMinimumPriceCadence = "semi_annual"
	PriceTieredPackageWithMinimumPriceCadenceAnnual     PriceTieredPackageWithMinimumPriceCadence = "annual"
	PriceTieredPackageWithMinimumPriceCadenceCustom     PriceTieredPackageWithMinimumPriceCadence = "custom"
)

func (r PriceTieredPackageWithMinimumPriceCadence) IsKnown() bool {
	switch r {
	case PriceTieredPackageWithMinimumPriceCadenceOneTime, PriceTieredPackageWithMinimumPriceCadenceMonthly, PriceTieredPackageWithMinimumPriceCadenceQuarterly, PriceTieredPackageWithMinimumPriceCadenceSemiAnnual, PriceTieredPackageWithMinimumPriceCadenceAnnual, PriceTieredPackageWithMinimumPriceCadenceCustom:
		return true
	}
	return false
}

type PriceTieredPackageWithMinimumPriceModelType string

const (
	PriceTieredPackageWithMinimumPriceModelTypeTieredPackageWithMinimum PriceTieredPackageWithMinimumPriceModelType = "tiered_package_with_minimum"
)

func (r PriceTieredPackageWithMinimumPriceModelType) IsKnown() bool {
	switch r {
	case PriceTieredPackageWithMinimumPriceModelTypeTieredPackageWithMinimum:
		return true
	}
	return false
}

type PriceTieredPackageWithMinimumPricePriceType string

const (
	PriceTieredPackageWithMinimumPricePriceTypeUsagePrice PriceTieredPackageWithMinimumPricePriceType = "usage_price"
	PriceTieredPackageWithMinimumPricePriceTypeFixedPrice PriceTieredPackageWithMinimumPricePriceType = "fixed_price"
)

func (r PriceTieredPackageWithMinimumPricePriceType) IsKnown() bool {
	switch r {
	case PriceTieredPackageWithMinimumPricePriceTypeUsagePrice, PriceTieredPackageWithMinimumPricePriceTypeFixedPrice:
		return true
	}
	return false
}

type PricePackageWithAllocationPrice struct {
	ID                        string                                 `json:"id,required"`
	BillableMetric            BillableMetricTiny                     `json:"billable_metric,required,nullable"`
	BillingCycleConfiguration BillingCycleConfiguration              `json:"billing_cycle_configuration,required"`
	Cadence                   PricePackageWithAllocationPriceCadence `json:"cadence,required"`
	ConversionRate            float64                                `json:"conversion_rate,required,nullable"`
	ConversionRateConfig      PricePackageWithAllocationPrice        `json:"conversion_rate_config,required,nullable"`
	CreatedAt                 time.Time                              `json:"created_at,required" format:"date-time"`
	CreditAllocation          Allocation                             `json:"credit_allocation,required,nullable"`
	Currency                  string                                 `json:"currency,required"`
	// Deprecated: deprecated
	Discount                    Discount                  `json:"discount,required,nullable"`
	ExternalPriceID             string                    `json:"external_price_id,required,nullable"`
	FixedPriceQuantity          float64                   `json:"fixed_price_quantity,required,nullable"`
	InvoicingCycleConfiguration BillingCycleConfiguration `json:"invoicing_cycle_configuration,required,nullable"`
	Item                        ItemSlim                  `json:"item,required"`
	// Deprecated: deprecated
	Maximum Maximum `json:"maximum,required,nullable"`
	// Deprecated: deprecated
	MaximumAmount string `json:"maximum_amount,required,nullable"`
	// User specified key-value pairs for the resource. If not present, this defaults
	// to an empty dictionary. Individual keys can be removed by setting the value to
	// `null`, and the entire metadata mapping can be cleared by setting `metadata` to
	// `null`.
	Metadata map[string]string `json:"metadata,required"`
	// Deprecated: deprecated
	Minimum Minimum `json:"minimum,required,nullable"`
	// Deprecated: deprecated
	MinimumAmount               string                                   `json:"minimum_amount,required,nullable"`
	ModelType                   PricePackageWithAllocationPriceModelType `json:"model_type,required"`
	Name                        string                                   `json:"name,required"`
	PackageWithAllocationConfig map[string]interface{}                   `json:"package_with_allocation_config,required"`
	PlanPhaseOrder              int64                                    `json:"plan_phase_order,required,nullable"`
	PriceType                   PricePackageWithAllocationPricePriceType `json:"price_type,required"`
	// The price id this price replaces. This price will take the place of the replaced
	// price in plan version migrations.
	ReplacesPriceID               string                              `json:"replaces_price_id,required,nullable"`
	DimensionalPriceConfiguration DimensionalPriceConfiguration       `json:"dimensional_price_configuration,nullable"`
	JSON                          pricePackageWithAllocationPriceJSON `json:"-"`
}

// pricePackageWithAllocationPriceJSON contains the JSON metadata for the struct
// [PricePackageWithAllocationPrice]
type pricePackageWithAllocationPriceJSON struct {
	ID                            apijson.Field
	BillableMetric                apijson.Field
	BillingCycleConfiguration     apijson.Field
	Cadence                       apijson.Field
	ConversionRate                apijson.Field
	ConversionRateConfig          apijson.Field
	CreatedAt                     apijson.Field
	CreditAllocation              apijson.Field
	Currency                      apijson.Field
	Discount                      apijson.Field
	ExternalPriceID               apijson.Field
	FixedPriceQuantity            apijson.Field
	InvoicingCycleConfiguration   apijson.Field
	Item                          apijson.Field
	Maximum                       apijson.Field
	MaximumAmount                 apijson.Field
	Metadata                      apijson.Field
	Minimum                       apijson.Field
	MinimumAmount                 apijson.Field
	ModelType                     apijson.Field
	Name                          apijson.Field
	PackageWithAllocationConfig   apijson.Field
	PlanPhaseOrder                apijson.Field
	PriceType                     apijson.Field
	ReplacesPriceID               apijson.Field
	DimensionalPriceConfiguration apijson.Field
	raw                           string
	ExtraFields                   map[string]apijson.Field
}

func (r *PricePackageWithAllocationPrice) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r pricePackageWithAllocationPriceJSON) RawJSON() string {
	return r.raw
}

func (r PricePackageWithAllocationPrice) implementsPrice() {}

type PricePackageWithAllocationPriceCadence string

const (
	PricePackageWithAllocationPriceCadenceOneTime    PricePackageWithAllocationPriceCadence = "one_time"
	PricePackageWithAllocationPriceCadenceMonthly    PricePackageWithAllocationPriceCadence = "monthly"
	PricePackageWithAllocationPriceCadenceQuarterly  PricePackageWithAllocationPriceCadence = "quarterly"
	PricePackageWithAllocationPriceCadenceSemiAnnual PricePackageWithAllocationPriceCadence = "semi_annual"
	PricePackageWithAllocationPriceCadenceAnnual     PricePackageWithAllocationPriceCadence = "annual"
	PricePackageWithAllocationPriceCadenceCustom     PricePackageWithAllocationPriceCadence = "custom"
)

func (r PricePackageWithAllocationPriceCadence) IsKnown() bool {
	switch r {
	case PricePackageWithAllocationPriceCadenceOneTime, PricePackageWithAllocationPriceCadenceMonthly, PricePackageWithAllocationPriceCadenceQuarterly, PricePackageWithAllocationPriceCadenceSemiAnnual, PricePackageWithAllocationPriceCadenceAnnual, PricePackageWithAllocationPriceCadenceCustom:
		return true
	}
	return false
}

type PricePackageWithAllocationPriceModelType string

const (
	PricePackageWithAllocationPriceModelTypePackageWithAllocation PricePackageWithAllocationPriceModelType = "package_with_allocation"
)

func (r PricePackageWithAllocationPriceModelType) IsKnown() bool {
	switch r {
	case PricePackageWithAllocationPriceModelTypePackageWithAllocation:
		return true
	}
	return false
}

type PricePackageWithAllocationPricePriceType string

const (
	PricePackageWithAllocationPricePriceTypeUsagePrice PricePackageWithAllocationPricePriceType = "usage_price"
	PricePackageWithAllocationPricePriceTypeFixedPrice PricePackageWithAllocationPricePriceType = "fixed_price"
)

func (r PricePackageWithAllocationPricePriceType) IsKnown() bool {
	switch r {
	case PricePackageWithAllocationPricePriceTypeUsagePrice, PricePackageWithAllocationPricePriceTypeFixedPrice:
		return true
	}
	return false
}

type PriceUnitWithPercentPrice struct {
	ID                        string                           `json:"id,required"`
	BillableMetric            BillableMetricTiny               `json:"billable_metric,required,nullable"`
	BillingCycleConfiguration BillingCycleConfiguration        `json:"billing_cycle_configuration,required"`
	Cadence                   PriceUnitWithPercentPriceCadence `json:"cadence,required"`
	ConversionRate            float64                          `json:"conversion_rate,required,nullable"`
	ConversionRateConfig      PriceUnitWithPercentPrice        `json:"conversion_rate_config,required,nullable"`
	CreatedAt                 time.Time                        `json:"created_at,required" format:"date-time"`
	CreditAllocation          Allocation                       `json:"credit_allocation,required,nullable"`
	Currency                  string                           `json:"currency,required"`
	// Deprecated: deprecated
	Discount                    Discount                  `json:"discount,required,nullable"`
	ExternalPriceID             string                    `json:"external_price_id,required,nullable"`
	FixedPriceQuantity          float64                   `json:"fixed_price_quantity,required,nullable"`
	InvoicingCycleConfiguration BillingCycleConfiguration `json:"invoicing_cycle_configuration,required,nullable"`
	Item                        ItemSlim                  `json:"item,required"`
	// Deprecated: deprecated
	Maximum Maximum `json:"maximum,required,nullable"`
	// Deprecated: deprecated
	MaximumAmount string `json:"maximum_amount,required,nullable"`
	// User specified key-value pairs for the resource. If not present, this defaults
	// to an empty dictionary. Individual keys can be removed by setting the value to
	// `null`, and the entire metadata mapping can be cleared by setting `metadata` to
	// `null`.
	Metadata map[string]string `json:"metadata,required"`
	// Deprecated: deprecated
	Minimum Minimum `json:"minimum,required,nullable"`
	// Deprecated: deprecated
	MinimumAmount  string                             `json:"minimum_amount,required,nullable"`
	ModelType      PriceUnitWithPercentPriceModelType `json:"model_type,required"`
	Name           string                             `json:"name,required"`
	PlanPhaseOrder int64                              `json:"plan_phase_order,required,nullable"`
	PriceType      PriceUnitWithPercentPricePriceType `json:"price_type,required"`
	// The price id this price replaces. This price will take the place of the replaced
	// price in plan version migrations.
	ReplacesPriceID               string                        `json:"replaces_price_id,required,nullable"`
	UnitWithPercentConfig         map[string]interface{}        `json:"unit_with_percent_config,required"`
	DimensionalPriceConfiguration DimensionalPriceConfiguration `json:"dimensional_price_configuration,nullable"`
	JSON                          priceUnitWithPercentPriceJSON `json:"-"`
}

// priceUnitWithPercentPriceJSON contains the JSON metadata for the struct
// [PriceUnitWithPercentPrice]
type priceUnitWithPercentPriceJSON struct {
	ID                            apijson.Field
	BillableMetric                apijson.Field
	BillingCycleConfiguration     apijson.Field
	Cadence                       apijson.Field
	ConversionRate                apijson.Field
	ConversionRateConfig          apijson.Field
	CreatedAt                     apijson.Field
	CreditAllocation              apijson.Field
	Currency                      apijson.Field
	Discount                      apijson.Field
	ExternalPriceID               apijson.Field
	FixedPriceQuantity            apijson.Field
	InvoicingCycleConfiguration   apijson.Field
	Item                          apijson.Field
	Maximum                       apijson.Field
	MaximumAmount                 apijson.Field
	Metadata                      apijson.Field
	Minimum                       apijson.Field
	MinimumAmount                 apijson.Field
	ModelType                     apijson.Field
	Name                          apijson.Field
	PlanPhaseOrder                apijson.Field
	PriceType                     apijson.Field
	ReplacesPriceID               apijson.Field
	UnitWithPercentConfig         apijson.Field
	DimensionalPriceConfiguration apijson.Field
	raw                           string
	ExtraFields                   map[string]apijson.Field
}

func (r *PriceUnitWithPercentPrice) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r priceUnitWithPercentPriceJSON) RawJSON() string {
	return r.raw
}

func (r PriceUnitWithPercentPrice) implementsPrice() {}

type PriceUnitWithPercentPriceCadence string

const (
	PriceUnitWithPercentPriceCadenceOneTime    PriceUnitWithPercentPriceCadence = "one_time"
	PriceUnitWithPercentPriceCadenceMonthly    PriceUnitWithPercentPriceCadence = "monthly"
	PriceUnitWithPercentPriceCadenceQuarterly  PriceUnitWithPercentPriceCadence = "quarterly"
	PriceUnitWithPercentPriceCadenceSemiAnnual PriceUnitWithPercentPriceCadence = "semi_annual"
	PriceUnitWithPercentPriceCadenceAnnual     PriceUnitWithPercentPriceCadence = "annual"
	PriceUnitWithPercentPriceCadenceCustom     PriceUnitWithPercentPriceCadence = "custom"
)

func (r PriceUnitWithPercentPriceCadence) IsKnown() bool {
	switch r {
	case PriceUnitWithPercentPriceCadenceOneTime, PriceUnitWithPercentPriceCadenceMonthly, PriceUnitWithPercentPriceCadenceQuarterly, PriceUnitWithPercentPriceCadenceSemiAnnual, PriceUnitWithPercentPriceCadenceAnnual, PriceUnitWithPercentPriceCadenceCustom:
		return true
	}
	return false
}

type PriceUnitWithPercentPriceModelType string

const (
	PriceUnitWithPercentPriceModelTypeUnitWithPercent PriceUnitWithPercentPriceModelType = "unit_with_percent"
)

func (r PriceUnitWithPercentPriceModelType) IsKnown() bool {
	switch r {
	case PriceUnitWithPercentPriceModelTypeUnitWithPercent:
		return true
	}
	return false
}

type PriceUnitWithPercentPricePriceType string

const (
	PriceUnitWithPercentPricePriceTypeUsagePrice PriceUnitWithPercentPricePriceType = "usage_price"
	PriceUnitWithPercentPricePriceTypeFixedPrice PriceUnitWithPercentPricePriceType = "fixed_price"
)

func (r PriceUnitWithPercentPricePriceType) IsKnown() bool {
	switch r {
	case PriceUnitWithPercentPricePriceTypeUsagePrice, PriceUnitWithPercentPricePriceTypeFixedPrice:
		return true
	}
	return false
}

type PriceMatrixWithAllocationPrice struct {
	ID                        string                                `json:"id,required"`
	BillableMetric            BillableMetricTiny                    `json:"billable_metric,required,nullable"`
	BillingCycleConfiguration BillingCycleConfiguration             `json:"billing_cycle_configuration,required"`
	Cadence                   PriceMatrixWithAllocationPriceCadence `json:"cadence,required"`
	ConversionRate            float64                               `json:"conversion_rate,required,nullable"`
	ConversionRateConfig      PriceMatrixWithAllocationPrice        `json:"conversion_rate_config,required,nullable"`
	CreatedAt                 time.Time                             `json:"created_at,required" format:"date-time"`
	CreditAllocation          Allocation                            `json:"credit_allocation,required,nullable"`
	Currency                  string                                `json:"currency,required"`
	// Deprecated: deprecated
	Discount                    Discount                   `json:"discount,required,nullable"`
	ExternalPriceID             string                     `json:"external_price_id,required,nullable"`
	FixedPriceQuantity          float64                    `json:"fixed_price_quantity,required,nullable"`
	InvoicingCycleConfiguration BillingCycleConfiguration  `json:"invoicing_cycle_configuration,required,nullable"`
	Item                        ItemSlim                   `json:"item,required"`
	MatrixWithAllocationConfig  MatrixWithAllocationConfig `json:"matrix_with_allocation_config,required"`
	// Deprecated: deprecated
	Maximum Maximum `json:"maximum,required,nullable"`
	// Deprecated: deprecated
	MaximumAmount string `json:"maximum_amount,required,nullable"`
	// User specified key-value pairs for the resource. If not present, this defaults
	// to an empty dictionary. Individual keys can be removed by setting the value to
	// `null`, and the entire metadata mapping can be cleared by setting `metadata` to
	// `null`.
	Metadata map[string]string `json:"metadata,required"`
	// Deprecated: deprecated
	Minimum Minimum `json:"minimum,required,nullable"`
	// Deprecated: deprecated
	MinimumAmount  string                                  `json:"minimum_amount,required,nullable"`
	ModelType      PriceMatrixWithAllocationPriceModelType `json:"model_type,required"`
	Name           string                                  `json:"name,required"`
	PlanPhaseOrder int64                                   `json:"plan_phase_order,required,nullable"`
	PriceType      PriceMatrixWithAllocationPricePriceType `json:"price_type,required"`
	// The price id this price replaces. This price will take the place of the replaced
	// price in plan version migrations.
	ReplacesPriceID               string                             `json:"replaces_price_id,required,nullable"`
	DimensionalPriceConfiguration DimensionalPriceConfiguration      `json:"dimensional_price_configuration,nullable"`
	JSON                          priceMatrixWithAllocationPriceJSON `json:"-"`
}

// priceMatrixWithAllocationPriceJSON contains the JSON metadata for the struct
// [PriceMatrixWithAllocationPrice]
type priceMatrixWithAllocationPriceJSON struct {
	ID                            apijson.Field
	BillableMetric                apijson.Field
	BillingCycleConfiguration     apijson.Field
	Cadence                       apijson.Field
	ConversionRate                apijson.Field
	ConversionRateConfig          apijson.Field
	CreatedAt                     apijson.Field
	CreditAllocation              apijson.Field
	Currency                      apijson.Field
	Discount                      apijson.Field
	ExternalPriceID               apijson.Field
	FixedPriceQuantity            apijson.Field
	InvoicingCycleConfiguration   apijson.Field
	Item                          apijson.Field
	MatrixWithAllocationConfig    apijson.Field
	Maximum                       apijson.Field
	MaximumAmount                 apijson.Field
	Metadata                      apijson.Field
	Minimum                       apijson.Field
	MinimumAmount                 apijson.Field
	ModelType                     apijson.Field
	Name                          apijson.Field
	PlanPhaseOrder                apijson.Field
	PriceType                     apijson.Field
	ReplacesPriceID               apijson.Field
	DimensionalPriceConfiguration apijson.Field
	raw                           string
	ExtraFields                   map[string]apijson.Field
}

func (r *PriceMatrixWithAllocationPrice) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r priceMatrixWithAllocationPriceJSON) RawJSON() string {
	return r.raw
}

func (r PriceMatrixWithAllocationPrice) implementsPrice() {}

type PriceMatrixWithAllocationPriceCadence string

const (
	PriceMatrixWithAllocationPriceCadenceOneTime    PriceMatrixWithAllocationPriceCadence = "one_time"
	PriceMatrixWithAllocationPriceCadenceMonthly    PriceMatrixWithAllocationPriceCadence = "monthly"
	PriceMatrixWithAllocationPriceCadenceQuarterly  PriceMatrixWithAllocationPriceCadence = "quarterly"
	PriceMatrixWithAllocationPriceCadenceSemiAnnual PriceMatrixWithAllocationPriceCadence = "semi_annual"
	PriceMatrixWithAllocationPriceCadenceAnnual     PriceMatrixWithAllocationPriceCadence = "annual"
	PriceMatrixWithAllocationPriceCadenceCustom     PriceMatrixWithAllocationPriceCadence = "custom"
)

func (r PriceMatrixWithAllocationPriceCadence) IsKnown() bool {
	switch r {
	case PriceMatrixWithAllocationPriceCadenceOneTime, PriceMatrixWithAllocationPriceCadenceMonthly, PriceMatrixWithAllocationPriceCadenceQuarterly, PriceMatrixWithAllocationPriceCadenceSemiAnnual, PriceMatrixWithAllocationPriceCadenceAnnual, PriceMatrixWithAllocationPriceCadenceCustom:
		return true
	}
	return false
}

type PriceMatrixWithAllocationPriceModelType string

const (
	PriceMatrixWithAllocationPriceModelTypeMatrixWithAllocation PriceMatrixWithAllocationPriceModelType = "matrix_with_allocation"
)

func (r PriceMatrixWithAllocationPriceModelType) IsKnown() bool {
	switch r {
	case PriceMatrixWithAllocationPriceModelTypeMatrixWithAllocation:
		return true
	}
	return false
}

type PriceMatrixWithAllocationPricePriceType string

const (
	PriceMatrixWithAllocationPricePriceTypeUsagePrice PriceMatrixWithAllocationPricePriceType = "usage_price"
	PriceMatrixWithAllocationPricePriceTypeFixedPrice PriceMatrixWithAllocationPricePriceType = "fixed_price"
)

func (r PriceMatrixWithAllocationPricePriceType) IsKnown() bool {
	switch r {
	case PriceMatrixWithAllocationPricePriceTypeUsagePrice, PriceMatrixWithAllocationPricePriceTypeFixedPrice:
		return true
	}
	return false
}

type PriceTieredWithProrationPrice struct {
	ID                        string                               `json:"id,required"`
	BillableMetric            BillableMetricTiny                   `json:"billable_metric,required,nullable"`
	BillingCycleConfiguration BillingCycleConfiguration            `json:"billing_cycle_configuration,required"`
	Cadence                   PriceTieredWithProrationPriceCadence `json:"cadence,required"`
	ConversionRate            float64                              `json:"conversion_rate,required,nullable"`
	ConversionRateConfig      PriceTieredWithProrationPrice        `json:"conversion_rate_config,required,nullable"`
	CreatedAt                 time.Time                            `json:"created_at,required" format:"date-time"`
	CreditAllocation          Allocation                           `json:"credit_allocation,required,nullable"`
	Currency                  string                               `json:"currency,required"`
	// Deprecated: deprecated
	Discount                    Discount                  `json:"discount,required,nullable"`
	ExternalPriceID             string                    `json:"external_price_id,required,nullable"`
	FixedPriceQuantity          float64                   `json:"fixed_price_quantity,required,nullable"`
	InvoicingCycleConfiguration BillingCycleConfiguration `json:"invoicing_cycle_configuration,required,nullable"`
	Item                        ItemSlim                  `json:"item,required"`
	// Deprecated: deprecated
	Maximum Maximum `json:"maximum,required,nullable"`
	// Deprecated: deprecated
	MaximumAmount string `json:"maximum_amount,required,nullable"`
	// User specified key-value pairs for the resource. If not present, this defaults
	// to an empty dictionary. Individual keys can be removed by setting the value to
	// `null`, and the entire metadata mapping can be cleared by setting `metadata` to
	// `null`.
	Metadata map[string]string `json:"metadata,required"`
	// Deprecated: deprecated
	Minimum Minimum `json:"minimum,required,nullable"`
	// Deprecated: deprecated
	MinimumAmount  string                                 `json:"minimum_amount,required,nullable"`
	ModelType      PriceTieredWithProrationPriceModelType `json:"model_type,required"`
	Name           string                                 `json:"name,required"`
	PlanPhaseOrder int64                                  `json:"plan_phase_order,required,nullable"`
	PriceType      PriceTieredWithProrationPricePriceType `json:"price_type,required"`
	// The price id this price replaces. This price will take the place of the replaced
	// price in plan version migrations.
	ReplacesPriceID               string                            `json:"replaces_price_id,required,nullable"`
	TieredWithProrationConfig     map[string]interface{}            `json:"tiered_with_proration_config,required"`
	DimensionalPriceConfiguration DimensionalPriceConfiguration     `json:"dimensional_price_configuration,nullable"`
	JSON                          priceTieredWithProrationPriceJSON `json:"-"`
}

// priceTieredWithProrationPriceJSON contains the JSON metadata for the struct
// [PriceTieredWithProrationPrice]
type priceTieredWithProrationPriceJSON struct {
	ID                            apijson.Field
	BillableMetric                apijson.Field
	BillingCycleConfiguration     apijson.Field
	Cadence                       apijson.Field
	ConversionRate                apijson.Field
	ConversionRateConfig          apijson.Field
	CreatedAt                     apijson.Field
	CreditAllocation              apijson.Field
	Currency                      apijson.Field
	Discount                      apijson.Field
	ExternalPriceID               apijson.Field
	FixedPriceQuantity            apijson.Field
	InvoicingCycleConfiguration   apijson.Field
	Item                          apijson.Field
	Maximum                       apijson.Field
	MaximumAmount                 apijson.Field
	Metadata                      apijson.Field
	Minimum                       apijson.Field
	MinimumAmount                 apijson.Field
	ModelType                     apijson.Field
	Name                          apijson.Field
	PlanPhaseOrder                apijson.Field
	PriceType                     apijson.Field
	ReplacesPriceID               apijson.Field
	TieredWithProrationConfig     apijson.Field
	DimensionalPriceConfiguration apijson.Field
	raw                           string
	ExtraFields                   map[string]apijson.Field
}

func (r *PriceTieredWithProrationPrice) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r priceTieredWithProrationPriceJSON) RawJSON() string {
	return r.raw
}

func (r PriceTieredWithProrationPrice) implementsPrice() {}

type PriceTieredWithProrationPriceCadence string

const (
	PriceTieredWithProrationPriceCadenceOneTime    PriceTieredWithProrationPriceCadence = "one_time"
	PriceTieredWithProrationPriceCadenceMonthly    PriceTieredWithProrationPriceCadence = "monthly"
	PriceTieredWithProrationPriceCadenceQuarterly  PriceTieredWithProrationPriceCadence = "quarterly"
	PriceTieredWithProrationPriceCadenceSemiAnnual PriceTieredWithProrationPriceCadence = "semi_annual"
	PriceTieredWithProrationPriceCadenceAnnual     PriceTieredWithProrationPriceCadence = "annual"
	PriceTieredWithProrationPriceCadenceCustom     PriceTieredWithProrationPriceCadence = "custom"
)

func (r PriceTieredWithProrationPriceCadence) IsKnown() bool {
	switch r {
	case PriceTieredWithProrationPriceCadenceOneTime, PriceTieredWithProrationPriceCadenceMonthly, PriceTieredWithProrationPriceCadenceQuarterly, PriceTieredWithProrationPriceCadenceSemiAnnual, PriceTieredWithProrationPriceCadenceAnnual, PriceTieredWithProrationPriceCadenceCustom:
		return true
	}
	return false
}

type PriceTieredWithProrationPriceModelType string

const (
	PriceTieredWithProrationPriceModelTypeTieredWithProration PriceTieredWithProrationPriceModelType = "tiered_with_proration"
)

func (r PriceTieredWithProrationPriceModelType) IsKnown() bool {
	switch r {
	case PriceTieredWithProrationPriceModelTypeTieredWithProration:
		return true
	}
	return false
}

type PriceTieredWithProrationPricePriceType string

const (
	PriceTieredWithProrationPricePriceTypeUsagePrice PriceTieredWithProrationPricePriceType = "usage_price"
	PriceTieredWithProrationPricePriceTypeFixedPrice PriceTieredWithProrationPricePriceType = "fixed_price"
)

func (r PriceTieredWithProrationPricePriceType) IsKnown() bool {
	switch r {
	case PriceTieredWithProrationPricePriceTypeUsagePrice, PriceTieredWithProrationPricePriceTypeFixedPrice:
		return true
	}
	return false
}

type PriceUnitWithProrationPrice struct {
	ID                        string                             `json:"id,required"`
	BillableMetric            BillableMetricTiny                 `json:"billable_metric,required,nullable"`
	BillingCycleConfiguration BillingCycleConfiguration          `json:"billing_cycle_configuration,required"`
	Cadence                   PriceUnitWithProrationPriceCadence `json:"cadence,required"`
	ConversionRate            float64                            `json:"conversion_rate,required,nullable"`
	ConversionRateConfig      PriceUnitWithProrationPrice        `json:"conversion_rate_config,required,nullable"`
	CreatedAt                 time.Time                          `json:"created_at,required" format:"date-time"`
	CreditAllocation          Allocation                         `json:"credit_allocation,required,nullable"`
	Currency                  string                             `json:"currency,required"`
	// Deprecated: deprecated
	Discount                    Discount                  `json:"discount,required,nullable"`
	ExternalPriceID             string                    `json:"external_price_id,required,nullable"`
	FixedPriceQuantity          float64                   `json:"fixed_price_quantity,required,nullable"`
	InvoicingCycleConfiguration BillingCycleConfiguration `json:"invoicing_cycle_configuration,required,nullable"`
	Item                        ItemSlim                  `json:"item,required"`
	// Deprecated: deprecated
	Maximum Maximum `json:"maximum,required,nullable"`
	// Deprecated: deprecated
	MaximumAmount string `json:"maximum_amount,required,nullable"`
	// User specified key-value pairs for the resource. If not present, this defaults
	// to an empty dictionary. Individual keys can be removed by setting the value to
	// `null`, and the entire metadata mapping can be cleared by setting `metadata` to
	// `null`.
	Metadata map[string]string `json:"metadata,required"`
	// Deprecated: deprecated
	Minimum Minimum `json:"minimum,required,nullable"`
	// Deprecated: deprecated
	MinimumAmount  string                               `json:"minimum_amount,required,nullable"`
	ModelType      PriceUnitWithProrationPriceModelType `json:"model_type,required"`
	Name           string                               `json:"name,required"`
	PlanPhaseOrder int64                                `json:"plan_phase_order,required,nullable"`
	PriceType      PriceUnitWithProrationPricePriceType `json:"price_type,required"`
	// The price id this price replaces. This price will take the place of the replaced
	// price in plan version migrations.
	ReplacesPriceID               string                          `json:"replaces_price_id,required,nullable"`
	UnitWithProrationConfig       map[string]interface{}          `json:"unit_with_proration_config,required"`
	DimensionalPriceConfiguration DimensionalPriceConfiguration   `json:"dimensional_price_configuration,nullable"`
	JSON                          priceUnitWithProrationPriceJSON `json:"-"`
}

// priceUnitWithProrationPriceJSON contains the JSON metadata for the struct
// [PriceUnitWithProrationPrice]
type priceUnitWithProrationPriceJSON struct {
	ID                            apijson.Field
	BillableMetric                apijson.Field
	BillingCycleConfiguration     apijson.Field
	Cadence                       apijson.Field
	ConversionRate                apijson.Field
	ConversionRateConfig          apijson.Field
	CreatedAt                     apijson.Field
	CreditAllocation              apijson.Field
	Currency                      apijson.Field
	Discount                      apijson.Field
	ExternalPriceID               apijson.Field
	FixedPriceQuantity            apijson.Field
	InvoicingCycleConfiguration   apijson.Field
	Item                          apijson.Field
	Maximum                       apijson.Field
	MaximumAmount                 apijson.Field
	Metadata                      apijson.Field
	Minimum                       apijson.Field
	MinimumAmount                 apijson.Field
	ModelType                     apijson.Field
	Name                          apijson.Field
	PlanPhaseOrder                apijson.Field
	PriceType                     apijson.Field
	ReplacesPriceID               apijson.Field
	UnitWithProrationConfig       apijson.Field
	DimensionalPriceConfiguration apijson.Field
	raw                           string
	ExtraFields                   map[string]apijson.Field
}

func (r *PriceUnitWithProrationPrice) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r priceUnitWithProrationPriceJSON) RawJSON() string {
	return r.raw
}

func (r PriceUnitWithProrationPrice) implementsPrice() {}

type PriceUnitWithProrationPriceCadence string

const (
	PriceUnitWithProrationPriceCadenceOneTime    PriceUnitWithProrationPriceCadence = "one_time"
	PriceUnitWithProrationPriceCadenceMonthly    PriceUnitWithProrationPriceCadence = "monthly"
	PriceUnitWithProrationPriceCadenceQuarterly  PriceUnitWithProrationPriceCadence = "quarterly"
	PriceUnitWithProrationPriceCadenceSemiAnnual PriceUnitWithProrationPriceCadence = "semi_annual"
	PriceUnitWithProrationPriceCadenceAnnual     PriceUnitWithProrationPriceCadence = "annual"
	PriceUnitWithProrationPriceCadenceCustom     PriceUnitWithProrationPriceCadence = "custom"
)

func (r PriceUnitWithProrationPriceCadence) IsKnown() bool {
	switch r {
	case PriceUnitWithProrationPriceCadenceOneTime, PriceUnitWithProrationPriceCadenceMonthly, PriceUnitWithProrationPriceCadenceQuarterly, PriceUnitWithProrationPriceCadenceSemiAnnual, PriceUnitWithProrationPriceCadenceAnnual, PriceUnitWithProrationPriceCadenceCustom:
		return true
	}
	return false
}

type PriceUnitWithProrationPriceModelType string

const (
	PriceUnitWithProrationPriceModelTypeUnitWithProration PriceUnitWithProrationPriceModelType = "unit_with_proration"
)

func (r PriceUnitWithProrationPriceModelType) IsKnown() bool {
	switch r {
	case PriceUnitWithProrationPriceModelTypeUnitWithProration:
		return true
	}
	return false
}

type PriceUnitWithProrationPricePriceType string

const (
	PriceUnitWithProrationPricePriceTypeUsagePrice PriceUnitWithProrationPricePriceType = "usage_price"
	PriceUnitWithProrationPricePriceTypeFixedPrice PriceUnitWithProrationPricePriceType = "fixed_price"
)

func (r PriceUnitWithProrationPricePriceType) IsKnown() bool {
	switch r {
	case PriceUnitWithProrationPricePriceTypeUsagePrice, PriceUnitWithProrationPricePriceTypeFixedPrice:
		return true
	}
	return false
}

type PriceGroupedAllocationPrice struct {
	ID                        string                             `json:"id,required"`
	BillableMetric            BillableMetricTiny                 `json:"billable_metric,required,nullable"`
	BillingCycleConfiguration BillingCycleConfiguration          `json:"billing_cycle_configuration,required"`
	Cadence                   PriceGroupedAllocationPriceCadence `json:"cadence,required"`
	ConversionRate            float64                            `json:"conversion_rate,required,nullable"`
	ConversionRateConfig      PriceGroupedAllocationPrice        `json:"conversion_rate_config,required,nullable"`
	CreatedAt                 time.Time                          `json:"created_at,required" format:"date-time"`
	CreditAllocation          Allocation                         `json:"credit_allocation,required,nullable"`
	Currency                  string                             `json:"currency,required"`
	// Deprecated: deprecated
	Discount                    Discount                  `json:"discount,required,nullable"`
	ExternalPriceID             string                    `json:"external_price_id,required,nullable"`
	FixedPriceQuantity          float64                   `json:"fixed_price_quantity,required,nullable"`
	GroupedAllocationConfig     map[string]interface{}    `json:"grouped_allocation_config,required"`
	InvoicingCycleConfiguration BillingCycleConfiguration `json:"invoicing_cycle_configuration,required,nullable"`
	Item                        ItemSlim                  `json:"item,required"`
	// Deprecated: deprecated
	Maximum Maximum `json:"maximum,required,nullable"`
	// Deprecated: deprecated
	MaximumAmount string `json:"maximum_amount,required,nullable"`
	// User specified key-value pairs for the resource. If not present, this defaults
	// to an empty dictionary. Individual keys can be removed by setting the value to
	// `null`, and the entire metadata mapping can be cleared by setting `metadata` to
	// `null`.
	Metadata map[string]string `json:"metadata,required"`
	// Deprecated: deprecated
	Minimum Minimum `json:"minimum,required,nullable"`
	// Deprecated: deprecated
	MinimumAmount  string                               `json:"minimum_amount,required,nullable"`
	ModelType      PriceGroupedAllocationPriceModelType `json:"model_type,required"`
	Name           string                               `json:"name,required"`
	PlanPhaseOrder int64                                `json:"plan_phase_order,required,nullable"`
	PriceType      PriceGroupedAllocationPricePriceType `json:"price_type,required"`
	// The price id this price replaces. This price will take the place of the replaced
	// price in plan version migrations.
	ReplacesPriceID               string                          `json:"replaces_price_id,required,nullable"`
	DimensionalPriceConfiguration DimensionalPriceConfiguration   `json:"dimensional_price_configuration,nullable"`
	JSON                          priceGroupedAllocationPriceJSON `json:"-"`
}

// priceGroupedAllocationPriceJSON contains the JSON metadata for the struct
// [PriceGroupedAllocationPrice]
type priceGroupedAllocationPriceJSON struct {
	ID                            apijson.Field
	BillableMetric                apijson.Field
	BillingCycleConfiguration     apijson.Field
	Cadence                       apijson.Field
	ConversionRate                apijson.Field
	ConversionRateConfig          apijson.Field
	CreatedAt                     apijson.Field
	CreditAllocation              apijson.Field
	Currency                      apijson.Field
	Discount                      apijson.Field
	ExternalPriceID               apijson.Field
	FixedPriceQuantity            apijson.Field
	GroupedAllocationConfig       apijson.Field
	InvoicingCycleConfiguration   apijson.Field
	Item                          apijson.Field
	Maximum                       apijson.Field
	MaximumAmount                 apijson.Field
	Metadata                      apijson.Field
	Minimum                       apijson.Field
	MinimumAmount                 apijson.Field
	ModelType                     apijson.Field
	Name                          apijson.Field
	PlanPhaseOrder                apijson.Field
	PriceType                     apijson.Field
	ReplacesPriceID               apijson.Field
	DimensionalPriceConfiguration apijson.Field
	raw                           string
	ExtraFields                   map[string]apijson.Field
}

func (r *PriceGroupedAllocationPrice) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r priceGroupedAllocationPriceJSON) RawJSON() string {
	return r.raw
}

func (r PriceGroupedAllocationPrice) implementsPrice() {}

type PriceGroupedAllocationPriceCadence string

const (
	PriceGroupedAllocationPriceCadenceOneTime    PriceGroupedAllocationPriceCadence = "one_time"
	PriceGroupedAllocationPriceCadenceMonthly    PriceGroupedAllocationPriceCadence = "monthly"
	PriceGroupedAllocationPriceCadenceQuarterly  PriceGroupedAllocationPriceCadence = "quarterly"
	PriceGroupedAllocationPriceCadenceSemiAnnual PriceGroupedAllocationPriceCadence = "semi_annual"
	PriceGroupedAllocationPriceCadenceAnnual     PriceGroupedAllocationPriceCadence = "annual"
	PriceGroupedAllocationPriceCadenceCustom     PriceGroupedAllocationPriceCadence = "custom"
)

func (r PriceGroupedAllocationPriceCadence) IsKnown() bool {
	switch r {
	case PriceGroupedAllocationPriceCadenceOneTime, PriceGroupedAllocationPriceCadenceMonthly, PriceGroupedAllocationPriceCadenceQuarterly, PriceGroupedAllocationPriceCadenceSemiAnnual, PriceGroupedAllocationPriceCadenceAnnual, PriceGroupedAllocationPriceCadenceCustom:
		return true
	}
	return false
}

type PriceGroupedAllocationPriceModelType string

const (
	PriceGroupedAllocationPriceModelTypeGroupedAllocation PriceGroupedAllocationPriceModelType = "grouped_allocation"
)

func (r PriceGroupedAllocationPriceModelType) IsKnown() bool {
	switch r {
	case PriceGroupedAllocationPriceModelTypeGroupedAllocation:
		return true
	}
	return false
}

type PriceGroupedAllocationPricePriceType string

const (
	PriceGroupedAllocationPricePriceTypeUsagePrice PriceGroupedAllocationPricePriceType = "usage_price"
	PriceGroupedAllocationPricePriceTypeFixedPrice PriceGroupedAllocationPricePriceType = "fixed_price"
)

func (r PriceGroupedAllocationPricePriceType) IsKnown() bool {
	switch r {
	case PriceGroupedAllocationPricePriceTypeUsagePrice, PriceGroupedAllocationPricePriceTypeFixedPrice:
		return true
	}
	return false
}

type PriceGroupedWithProratedMinimumPrice struct {
	ID                        string                                      `json:"id,required"`
	BillableMetric            BillableMetricTiny                          `json:"billable_metric,required,nullable"`
	BillingCycleConfiguration BillingCycleConfiguration                   `json:"billing_cycle_configuration,required"`
	Cadence                   PriceGroupedWithProratedMinimumPriceCadence `json:"cadence,required"`
	ConversionRate            float64                                     `json:"conversion_rate,required,nullable"`
	ConversionRateConfig      PriceGroupedWithProratedMinimumPrice        `json:"conversion_rate_config,required,nullable"`
	CreatedAt                 time.Time                                   `json:"created_at,required" format:"date-time"`
	CreditAllocation          Allocation                                  `json:"credit_allocation,required,nullable"`
	Currency                  string                                      `json:"currency,required"`
	// Deprecated: deprecated
	Discount                         Discount                  `json:"discount,required,nullable"`
	ExternalPriceID                  string                    `json:"external_price_id,required,nullable"`
	FixedPriceQuantity               float64                   `json:"fixed_price_quantity,required,nullable"`
	GroupedWithProratedMinimumConfig map[string]interface{}    `json:"grouped_with_prorated_minimum_config,required"`
	InvoicingCycleConfiguration      BillingCycleConfiguration `json:"invoicing_cycle_configuration,required,nullable"`
	Item                             ItemSlim                  `json:"item,required"`
	// Deprecated: deprecated
	Maximum Maximum `json:"maximum,required,nullable"`
	// Deprecated: deprecated
	MaximumAmount string `json:"maximum_amount,required,nullable"`
	// User specified key-value pairs for the resource. If not present, this defaults
	// to an empty dictionary. Individual keys can be removed by setting the value to
	// `null`, and the entire metadata mapping can be cleared by setting `metadata` to
	// `null`.
	Metadata map[string]string `json:"metadata,required"`
	// Deprecated: deprecated
	Minimum Minimum `json:"minimum,required,nullable"`
	// Deprecated: deprecated
	MinimumAmount  string                                        `json:"minimum_amount,required,nullable"`
	ModelType      PriceGroupedWithProratedMinimumPriceModelType `json:"model_type,required"`
	Name           string                                        `json:"name,required"`
	PlanPhaseOrder int64                                         `json:"plan_phase_order,required,nullable"`
	PriceType      PriceGroupedWithProratedMinimumPricePriceType `json:"price_type,required"`
	// The price id this price replaces. This price will take the place of the replaced
	// price in plan version migrations.
	ReplacesPriceID               string                                   `json:"replaces_price_id,required,nullable"`
	DimensionalPriceConfiguration DimensionalPriceConfiguration            `json:"dimensional_price_configuration,nullable"`
	JSON                          priceGroupedWithProratedMinimumPriceJSON `json:"-"`
}

// priceGroupedWithProratedMinimumPriceJSON contains the JSON metadata for the
// struct [PriceGroupedWithProratedMinimumPrice]
type priceGroupedWithProratedMinimumPriceJSON struct {
	ID                               apijson.Field
	BillableMetric                   apijson.Field
	BillingCycleConfiguration        apijson.Field
	Cadence                          apijson.Field
	ConversionRate                   apijson.Field
	ConversionRateConfig             apijson.Field
	CreatedAt                        apijson.Field
	CreditAllocation                 apijson.Field
	Currency                         apijson.Field
	Discount                         apijson.Field
	ExternalPriceID                  apijson.Field
	FixedPriceQuantity               apijson.Field
	GroupedWithProratedMinimumConfig apijson.Field
	InvoicingCycleConfiguration      apijson.Field
	Item                             apijson.Field
	Maximum                          apijson.Field
	MaximumAmount                    apijson.Field
	Metadata                         apijson.Field
	Minimum                          apijson.Field
	MinimumAmount                    apijson.Field
	ModelType                        apijson.Field
	Name                             apijson.Field
	PlanPhaseOrder                   apijson.Field
	PriceType                        apijson.Field
	ReplacesPriceID                  apijson.Field
	DimensionalPriceConfiguration    apijson.Field
	raw                              string
	ExtraFields                      map[string]apijson.Field
}

func (r *PriceGroupedWithProratedMinimumPrice) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r priceGroupedWithProratedMinimumPriceJSON) RawJSON() string {
	return r.raw
}

func (r PriceGroupedWithProratedMinimumPrice) implementsPrice() {}

type PriceGroupedWithProratedMinimumPriceCadence string

const (
	PriceGroupedWithProratedMinimumPriceCadenceOneTime    PriceGroupedWithProratedMinimumPriceCadence = "one_time"
	PriceGroupedWithProratedMinimumPriceCadenceMonthly    PriceGroupedWithProratedMinimumPriceCadence = "monthly"
	PriceGroupedWithProratedMinimumPriceCadenceQuarterly  PriceGroupedWithProratedMinimumPriceCadence = "quarterly"
	PriceGroupedWithProratedMinimumPriceCadenceSemiAnnual PriceGroupedWithProratedMinimumPriceCadence = "semi_annual"
	PriceGroupedWithProratedMinimumPriceCadenceAnnual     PriceGroupedWithProratedMinimumPriceCadence = "annual"
	PriceGroupedWithProratedMinimumPriceCadenceCustom     PriceGroupedWithProratedMinimumPriceCadence = "custom"
)

func (r PriceGroupedWithProratedMinimumPriceCadence) IsKnown() bool {
	switch r {
	case PriceGroupedWithProratedMinimumPriceCadenceOneTime, PriceGroupedWithProratedMinimumPriceCadenceMonthly, PriceGroupedWithProratedMinimumPriceCadenceQuarterly, PriceGroupedWithProratedMinimumPriceCadenceSemiAnnual, PriceGroupedWithProratedMinimumPriceCadenceAnnual, PriceGroupedWithProratedMinimumPriceCadenceCustom:
		return true
	}
	return false
}

type PriceGroupedWithProratedMinimumPriceModelType string

const (
	PriceGroupedWithProratedMinimumPriceModelTypeGroupedWithProratedMinimum PriceGroupedWithProratedMinimumPriceModelType = "grouped_with_prorated_minimum"
)

func (r PriceGroupedWithProratedMinimumPriceModelType) IsKnown() bool {
	switch r {
	case PriceGroupedWithProratedMinimumPriceModelTypeGroupedWithProratedMinimum:
		return true
	}
	return false
}

type PriceGroupedWithProratedMinimumPricePriceType string

const (
	PriceGroupedWithProratedMinimumPricePriceTypeUsagePrice PriceGroupedWithProratedMinimumPricePriceType = "usage_price"
	PriceGroupedWithProratedMinimumPricePriceTypeFixedPrice PriceGroupedWithProratedMinimumPricePriceType = "fixed_price"
)

func (r PriceGroupedWithProratedMinimumPricePriceType) IsKnown() bool {
	switch r {
	case PriceGroupedWithProratedMinimumPricePriceTypeUsagePrice, PriceGroupedWithProratedMinimumPricePriceTypeFixedPrice:
		return true
	}
	return false
}

type PriceGroupedWithMeteredMinimumPrice struct {
	ID                        string                                     `json:"id,required"`
	BillableMetric            BillableMetricTiny                         `json:"billable_metric,required,nullable"`
	BillingCycleConfiguration BillingCycleConfiguration                  `json:"billing_cycle_configuration,required"`
	Cadence                   PriceGroupedWithMeteredMinimumPriceCadence `json:"cadence,required"`
	ConversionRate            float64                                    `json:"conversion_rate,required,nullable"`
	ConversionRateConfig      PriceGroupedWithMeteredMinimumPrice        `json:"conversion_rate_config,required,nullable"`
	CreatedAt                 time.Time                                  `json:"created_at,required" format:"date-time"`
	CreditAllocation          Allocation                                 `json:"credit_allocation,required,nullable"`
	Currency                  string                                     `json:"currency,required"`
	// Deprecated: deprecated
	Discount                        Discount                  `json:"discount,required,nullable"`
	ExternalPriceID                 string                    `json:"external_price_id,required,nullable"`
	FixedPriceQuantity              float64                   `json:"fixed_price_quantity,required,nullable"`
	GroupedWithMeteredMinimumConfig map[string]interface{}    `json:"grouped_with_metered_minimum_config,required"`
	InvoicingCycleConfiguration     BillingCycleConfiguration `json:"invoicing_cycle_configuration,required,nullable"`
	Item                            ItemSlim                  `json:"item,required"`
	// Deprecated: deprecated
	Maximum Maximum `json:"maximum,required,nullable"`
	// Deprecated: deprecated
	MaximumAmount string `json:"maximum_amount,required,nullable"`
	// User specified key-value pairs for the resource. If not present, this defaults
	// to an empty dictionary. Individual keys can be removed by setting the value to
	// `null`, and the entire metadata mapping can be cleared by setting `metadata` to
	// `null`.
	Metadata map[string]string `json:"metadata,required"`
	// Deprecated: deprecated
	Minimum Minimum `json:"minimum,required,nullable"`
	// Deprecated: deprecated
	MinimumAmount  string                                       `json:"minimum_amount,required,nullable"`
	ModelType      PriceGroupedWithMeteredMinimumPriceModelType `json:"model_type,required"`
	Name           string                                       `json:"name,required"`
	PlanPhaseOrder int64                                        `json:"plan_phase_order,required,nullable"`
	PriceType      PriceGroupedWithMeteredMinimumPricePriceType `json:"price_type,required"`
	// The price id this price replaces. This price will take the place of the replaced
	// price in plan version migrations.
	ReplacesPriceID               string                                  `json:"replaces_price_id,required,nullable"`
	DimensionalPriceConfiguration DimensionalPriceConfiguration           `json:"dimensional_price_configuration,nullable"`
	JSON                          priceGroupedWithMeteredMinimumPriceJSON `json:"-"`
}

// priceGroupedWithMeteredMinimumPriceJSON contains the JSON metadata for the
// struct [PriceGroupedWithMeteredMinimumPrice]
type priceGroupedWithMeteredMinimumPriceJSON struct {
	ID                              apijson.Field
	BillableMetric                  apijson.Field
	BillingCycleConfiguration       apijson.Field
	Cadence                         apijson.Field
	ConversionRate                  apijson.Field
	ConversionRateConfig            apijson.Field
	CreatedAt                       apijson.Field
	CreditAllocation                apijson.Field
	Currency                        apijson.Field
	Discount                        apijson.Field
	ExternalPriceID                 apijson.Field
	FixedPriceQuantity              apijson.Field
	GroupedWithMeteredMinimumConfig apijson.Field
	InvoicingCycleConfiguration     apijson.Field
	Item                            apijson.Field
	Maximum                         apijson.Field
	MaximumAmount                   apijson.Field
	Metadata                        apijson.Field
	Minimum                         apijson.Field
	MinimumAmount                   apijson.Field
	ModelType                       apijson.Field
	Name                            apijson.Field
	PlanPhaseOrder                  apijson.Field
	PriceType                       apijson.Field
	ReplacesPriceID                 apijson.Field
	DimensionalPriceConfiguration   apijson.Field
	raw                             string
	ExtraFields                     map[string]apijson.Field
}

func (r *PriceGroupedWithMeteredMinimumPrice) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r priceGroupedWithMeteredMinimumPriceJSON) RawJSON() string {
	return r.raw
}

func (r PriceGroupedWithMeteredMinimumPrice) implementsPrice() {}

type PriceGroupedWithMeteredMinimumPriceCadence string

const (
	PriceGroupedWithMeteredMinimumPriceCadenceOneTime    PriceGroupedWithMeteredMinimumPriceCadence = "one_time"
	PriceGroupedWithMeteredMinimumPriceCadenceMonthly    PriceGroupedWithMeteredMinimumPriceCadence = "monthly"
	PriceGroupedWithMeteredMinimumPriceCadenceQuarterly  PriceGroupedWithMeteredMinimumPriceCadence = "quarterly"
	PriceGroupedWithMeteredMinimumPriceCadenceSemiAnnual PriceGroupedWithMeteredMinimumPriceCadence = "semi_annual"
	PriceGroupedWithMeteredMinimumPriceCadenceAnnual     PriceGroupedWithMeteredMinimumPriceCadence = "annual"
	PriceGroupedWithMeteredMinimumPriceCadenceCustom     PriceGroupedWithMeteredMinimumPriceCadence = "custom"
)

func (r PriceGroupedWithMeteredMinimumPriceCadence) IsKnown() bool {
	switch r {
	case PriceGroupedWithMeteredMinimumPriceCadenceOneTime, PriceGroupedWithMeteredMinimumPriceCadenceMonthly, PriceGroupedWithMeteredMinimumPriceCadenceQuarterly, PriceGroupedWithMeteredMinimumPriceCadenceSemiAnnual, PriceGroupedWithMeteredMinimumPriceCadenceAnnual, PriceGroupedWithMeteredMinimumPriceCadenceCustom:
		return true
	}
	return false
}

type PriceGroupedWithMeteredMinimumPriceModelType string

const (
	PriceGroupedWithMeteredMinimumPriceModelTypeGroupedWithMeteredMinimum PriceGroupedWithMeteredMinimumPriceModelType = "grouped_with_metered_minimum"
)

func (r PriceGroupedWithMeteredMinimumPriceModelType) IsKnown() bool {
	switch r {
	case PriceGroupedWithMeteredMinimumPriceModelTypeGroupedWithMeteredMinimum:
		return true
	}
	return false
}

type PriceGroupedWithMeteredMinimumPricePriceType string

const (
	PriceGroupedWithMeteredMinimumPricePriceTypeUsagePrice PriceGroupedWithMeteredMinimumPricePriceType = "usage_price"
	PriceGroupedWithMeteredMinimumPricePriceTypeFixedPrice PriceGroupedWithMeteredMinimumPricePriceType = "fixed_price"
)

func (r PriceGroupedWithMeteredMinimumPricePriceType) IsKnown() bool {
	switch r {
	case PriceGroupedWithMeteredMinimumPricePriceTypeUsagePrice, PriceGroupedWithMeteredMinimumPricePriceTypeFixedPrice:
		return true
	}
	return false
}

type PriceMatrixWithDisplayNamePrice struct {
	ID                        string                                 `json:"id,required"`
	BillableMetric            BillableMetricTiny                     `json:"billable_metric,required,nullable"`
	BillingCycleConfiguration BillingCycleConfiguration              `json:"billing_cycle_configuration,required"`
	Cadence                   PriceMatrixWithDisplayNamePriceCadence `json:"cadence,required"`
	ConversionRate            float64                                `json:"conversion_rate,required,nullable"`
	ConversionRateConfig      PriceMatrixWithDisplayNamePrice        `json:"conversion_rate_config,required,nullable"`
	CreatedAt                 time.Time                              `json:"created_at,required" format:"date-time"`
	CreditAllocation          Allocation                             `json:"credit_allocation,required,nullable"`
	Currency                  string                                 `json:"currency,required"`
	// Deprecated: deprecated
	Discount                    Discount                  `json:"discount,required,nullable"`
	ExternalPriceID             string                    `json:"external_price_id,required,nullable"`
	FixedPriceQuantity          float64                   `json:"fixed_price_quantity,required,nullable"`
	InvoicingCycleConfiguration BillingCycleConfiguration `json:"invoicing_cycle_configuration,required,nullable"`
	Item                        ItemSlim                  `json:"item,required"`
	MatrixWithDisplayNameConfig map[string]interface{}    `json:"matrix_with_display_name_config,required"`
	// Deprecated: deprecated
	Maximum Maximum `json:"maximum,required,nullable"`
	// Deprecated: deprecated
	MaximumAmount string `json:"maximum_amount,required,nullable"`
	// User specified key-value pairs for the resource. If not present, this defaults
	// to an empty dictionary. Individual keys can be removed by setting the value to
	// `null`, and the entire metadata mapping can be cleared by setting `metadata` to
	// `null`.
	Metadata map[string]string `json:"metadata,required"`
	// Deprecated: deprecated
	Minimum Minimum `json:"minimum,required,nullable"`
	// Deprecated: deprecated
	MinimumAmount  string                                   `json:"minimum_amount,required,nullable"`
	ModelType      PriceMatrixWithDisplayNamePriceModelType `json:"model_type,required"`
	Name           string                                   `json:"name,required"`
	PlanPhaseOrder int64                                    `json:"plan_phase_order,required,nullable"`
	PriceType      PriceMatrixWithDisplayNamePricePriceType `json:"price_type,required"`
	// The price id this price replaces. This price will take the place of the replaced
	// price in plan version migrations.
	ReplacesPriceID               string                              `json:"replaces_price_id,required,nullable"`
	DimensionalPriceConfiguration DimensionalPriceConfiguration       `json:"dimensional_price_configuration,nullable"`
	JSON                          priceMatrixWithDisplayNamePriceJSON `json:"-"`
}

// priceMatrixWithDisplayNamePriceJSON contains the JSON metadata for the struct
// [PriceMatrixWithDisplayNamePrice]
type priceMatrixWithDisplayNamePriceJSON struct {
	ID                            apijson.Field
	BillableMetric                apijson.Field
	BillingCycleConfiguration     apijson.Field
	Cadence                       apijson.Field
	ConversionRate                apijson.Field
	ConversionRateConfig          apijson.Field
	CreatedAt                     apijson.Field
	CreditAllocation              apijson.Field
	Currency                      apijson.Field
	Discount                      apijson.Field
	ExternalPriceID               apijson.Field
	FixedPriceQuantity            apijson.Field
	InvoicingCycleConfiguration   apijson.Field
	Item                          apijson.Field
	MatrixWithDisplayNameConfig   apijson.Field
	Maximum                       apijson.Field
	MaximumAmount                 apijson.Field
	Metadata                      apijson.Field
	Minimum                       apijson.Field
	MinimumAmount                 apijson.Field
	ModelType                     apijson.Field
	Name                          apijson.Field
	PlanPhaseOrder                apijson.Field
	PriceType                     apijson.Field
	ReplacesPriceID               apijson.Field
	DimensionalPriceConfiguration apijson.Field
	raw                           string
	ExtraFields                   map[string]apijson.Field
}

func (r *PriceMatrixWithDisplayNamePrice) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r priceMatrixWithDisplayNamePriceJSON) RawJSON() string {
	return r.raw
}

func (r PriceMatrixWithDisplayNamePrice) implementsPrice() {}

type PriceMatrixWithDisplayNamePriceCadence string

const (
	PriceMatrixWithDisplayNamePriceCadenceOneTime    PriceMatrixWithDisplayNamePriceCadence = "one_time"
	PriceMatrixWithDisplayNamePriceCadenceMonthly    PriceMatrixWithDisplayNamePriceCadence = "monthly"
	PriceMatrixWithDisplayNamePriceCadenceQuarterly  PriceMatrixWithDisplayNamePriceCadence = "quarterly"
	PriceMatrixWithDisplayNamePriceCadenceSemiAnnual PriceMatrixWithDisplayNamePriceCadence = "semi_annual"
	PriceMatrixWithDisplayNamePriceCadenceAnnual     PriceMatrixWithDisplayNamePriceCadence = "annual"
	PriceMatrixWithDisplayNamePriceCadenceCustom     PriceMatrixWithDisplayNamePriceCadence = "custom"
)

func (r PriceMatrixWithDisplayNamePriceCadence) IsKnown() bool {
	switch r {
	case PriceMatrixWithDisplayNamePriceCadenceOneTime, PriceMatrixWithDisplayNamePriceCadenceMonthly, PriceMatrixWithDisplayNamePriceCadenceQuarterly, PriceMatrixWithDisplayNamePriceCadenceSemiAnnual, PriceMatrixWithDisplayNamePriceCadenceAnnual, PriceMatrixWithDisplayNamePriceCadenceCustom:
		return true
	}
	return false
}

type PriceMatrixWithDisplayNamePriceModelType string

const (
	PriceMatrixWithDisplayNamePriceModelTypeMatrixWithDisplayName PriceMatrixWithDisplayNamePriceModelType = "matrix_with_display_name"
)

func (r PriceMatrixWithDisplayNamePriceModelType) IsKnown() bool {
	switch r {
	case PriceMatrixWithDisplayNamePriceModelTypeMatrixWithDisplayName:
		return true
	}
	return false
}

type PriceMatrixWithDisplayNamePricePriceType string

const (
	PriceMatrixWithDisplayNamePricePriceTypeUsagePrice PriceMatrixWithDisplayNamePricePriceType = "usage_price"
	PriceMatrixWithDisplayNamePricePriceTypeFixedPrice PriceMatrixWithDisplayNamePricePriceType = "fixed_price"
)

func (r PriceMatrixWithDisplayNamePricePriceType) IsKnown() bool {
	switch r {
	case PriceMatrixWithDisplayNamePricePriceTypeUsagePrice, PriceMatrixWithDisplayNamePricePriceTypeFixedPrice:
		return true
	}
	return false
}

type PriceBulkWithProrationPrice struct {
	ID                        string                             `json:"id,required"`
	BillableMetric            BillableMetricTiny                 `json:"billable_metric,required,nullable"`
	BillingCycleConfiguration BillingCycleConfiguration          `json:"billing_cycle_configuration,required"`
	BulkWithProrationConfig   map[string]interface{}             `json:"bulk_with_proration_config,required"`
	Cadence                   PriceBulkWithProrationPriceCadence `json:"cadence,required"`
	ConversionRate            float64                            `json:"conversion_rate,required,nullable"`
	ConversionRateConfig      PriceBulkWithProrationPrice        `json:"conversion_rate_config,required,nullable"`
	CreatedAt                 time.Time                          `json:"created_at,required" format:"date-time"`
	CreditAllocation          Allocation                         `json:"credit_allocation,required,nullable"`
	Currency                  string                             `json:"currency,required"`
	// Deprecated: deprecated
	Discount                    Discount                  `json:"discount,required,nullable"`
	ExternalPriceID             string                    `json:"external_price_id,required,nullable"`
	FixedPriceQuantity          float64                   `json:"fixed_price_quantity,required,nullable"`
	InvoicingCycleConfiguration BillingCycleConfiguration `json:"invoicing_cycle_configuration,required,nullable"`
	Item                        ItemSlim                  `json:"item,required"`
	// Deprecated: deprecated
	Maximum Maximum `json:"maximum,required,nullable"`
	// Deprecated: deprecated
	MaximumAmount string `json:"maximum_amount,required,nullable"`
	// User specified key-value pairs for the resource. If not present, this defaults
	// to an empty dictionary. Individual keys can be removed by setting the value to
	// `null`, and the entire metadata mapping can be cleared by setting `metadata` to
	// `null`.
	Metadata map[string]string `json:"metadata,required"`
	// Deprecated: deprecated
	Minimum Minimum `json:"minimum,required,nullable"`
	// Deprecated: deprecated
	MinimumAmount  string                               `json:"minimum_amount,required,nullable"`
	ModelType      PriceBulkWithProrationPriceModelType `json:"model_type,required"`
	Name           string                               `json:"name,required"`
	PlanPhaseOrder int64                                `json:"plan_phase_order,required,nullable"`
	PriceType      PriceBulkWithProrationPricePriceType `json:"price_type,required"`
	// The price id this price replaces. This price will take the place of the replaced
	// price in plan version migrations.
	ReplacesPriceID               string                          `json:"replaces_price_id,required,nullable"`
	DimensionalPriceConfiguration DimensionalPriceConfiguration   `json:"dimensional_price_configuration,nullable"`
	JSON                          priceBulkWithProrationPriceJSON `json:"-"`
}

// priceBulkWithProrationPriceJSON contains the JSON metadata for the struct
// [PriceBulkWithProrationPrice]
type priceBulkWithProrationPriceJSON struct {
	ID                            apijson.Field
	BillableMetric                apijson.Field
	BillingCycleConfiguration     apijson.Field
	BulkWithProrationConfig       apijson.Field
	Cadence                       apijson.Field
	ConversionRate                apijson.Field
	ConversionRateConfig          apijson.Field
	CreatedAt                     apijson.Field
	CreditAllocation              apijson.Field
	Currency                      apijson.Field
	Discount                      apijson.Field
	ExternalPriceID               apijson.Field
	FixedPriceQuantity            apijson.Field
	InvoicingCycleConfiguration   apijson.Field
	Item                          apijson.Field
	Maximum                       apijson.Field
	MaximumAmount                 apijson.Field
	Metadata                      apijson.Field
	Minimum                       apijson.Field
	MinimumAmount                 apijson.Field
	ModelType                     apijson.Field
	Name                          apijson.Field
	PlanPhaseOrder                apijson.Field
	PriceType                     apijson.Field
	ReplacesPriceID               apijson.Field
	DimensionalPriceConfiguration apijson.Field
	raw                           string
	ExtraFields                   map[string]apijson.Field
}

func (r *PriceBulkWithProrationPrice) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r priceBulkWithProrationPriceJSON) RawJSON() string {
	return r.raw
}

func (r PriceBulkWithProrationPrice) implementsPrice() {}

type PriceBulkWithProrationPriceCadence string

const (
	PriceBulkWithProrationPriceCadenceOneTime    PriceBulkWithProrationPriceCadence = "one_time"
	PriceBulkWithProrationPriceCadenceMonthly    PriceBulkWithProrationPriceCadence = "monthly"
	PriceBulkWithProrationPriceCadenceQuarterly  PriceBulkWithProrationPriceCadence = "quarterly"
	PriceBulkWithProrationPriceCadenceSemiAnnual PriceBulkWithProrationPriceCadence = "semi_annual"
	PriceBulkWithProrationPriceCadenceAnnual     PriceBulkWithProrationPriceCadence = "annual"
	PriceBulkWithProrationPriceCadenceCustom     PriceBulkWithProrationPriceCadence = "custom"
)

func (r PriceBulkWithProrationPriceCadence) IsKnown() bool {
	switch r {
	case PriceBulkWithProrationPriceCadenceOneTime, PriceBulkWithProrationPriceCadenceMonthly, PriceBulkWithProrationPriceCadenceQuarterly, PriceBulkWithProrationPriceCadenceSemiAnnual, PriceBulkWithProrationPriceCadenceAnnual, PriceBulkWithProrationPriceCadenceCustom:
		return true
	}
	return false
}

type PriceBulkWithProrationPriceModelType string

const (
	PriceBulkWithProrationPriceModelTypeBulkWithProration PriceBulkWithProrationPriceModelType = "bulk_with_proration"
)

func (r PriceBulkWithProrationPriceModelType) IsKnown() bool {
	switch r {
	case PriceBulkWithProrationPriceModelTypeBulkWithProration:
		return true
	}
	return false
}

type PriceBulkWithProrationPricePriceType string

const (
	PriceBulkWithProrationPricePriceTypeUsagePrice PriceBulkWithProrationPricePriceType = "usage_price"
	PriceBulkWithProrationPricePriceTypeFixedPrice PriceBulkWithProrationPricePriceType = "fixed_price"
)

func (r PriceBulkWithProrationPricePriceType) IsKnown() bool {
	switch r {
	case PriceBulkWithProrationPricePriceTypeUsagePrice, PriceBulkWithProrationPricePriceTypeFixedPrice:
		return true
	}
	return false
}

type PriceGroupedTieredPackagePrice struct {
	ID                        string                                `json:"id,required"`
	BillableMetric            BillableMetricTiny                    `json:"billable_metric,required,nullable"`
	BillingCycleConfiguration BillingCycleConfiguration             `json:"billing_cycle_configuration,required"`
	Cadence                   PriceGroupedTieredPackagePriceCadence `json:"cadence,required"`
	ConversionRate            float64                               `json:"conversion_rate,required,nullable"`
	ConversionRateConfig      PriceGroupedTieredPackagePrice        `json:"conversion_rate_config,required,nullable"`
	CreatedAt                 time.Time                             `json:"created_at,required" format:"date-time"`
	CreditAllocation          Allocation                            `json:"credit_allocation,required,nullable"`
	Currency                  string                                `json:"currency,required"`
	// Deprecated: deprecated
	Discount                    Discount                  `json:"discount,required,nullable"`
	ExternalPriceID             string                    `json:"external_price_id,required,nullable"`
	FixedPriceQuantity          float64                   `json:"fixed_price_quantity,required,nullable"`
	GroupedTieredPackageConfig  map[string]interface{}    `json:"grouped_tiered_package_config,required"`
	InvoicingCycleConfiguration BillingCycleConfiguration `json:"invoicing_cycle_configuration,required,nullable"`
	Item                        ItemSlim                  `json:"item,required"`
	// Deprecated: deprecated
	Maximum Maximum `json:"maximum,required,nullable"`
	// Deprecated: deprecated
	MaximumAmount string `json:"maximum_amount,required,nullable"`
	// User specified key-value pairs for the resource. If not present, this defaults
	// to an empty dictionary. Individual keys can be removed by setting the value to
	// `null`, and the entire metadata mapping can be cleared by setting `metadata` to
	// `null`.
	Metadata map[string]string `json:"metadata,required"`
	// Deprecated: deprecated
	Minimum Minimum `json:"minimum,required,nullable"`
	// Deprecated: deprecated
	MinimumAmount  string                                  `json:"minimum_amount,required,nullable"`
	ModelType      PriceGroupedTieredPackagePriceModelType `json:"model_type,required"`
	Name           string                                  `json:"name,required"`
	PlanPhaseOrder int64                                   `json:"plan_phase_order,required,nullable"`
	PriceType      PriceGroupedTieredPackagePricePriceType `json:"price_type,required"`
	// The price id this price replaces. This price will take the place of the replaced
	// price in plan version migrations.
	ReplacesPriceID               string                             `json:"replaces_price_id,required,nullable"`
	DimensionalPriceConfiguration DimensionalPriceConfiguration      `json:"dimensional_price_configuration,nullable"`
	JSON                          priceGroupedTieredPackagePriceJSON `json:"-"`
}

// priceGroupedTieredPackagePriceJSON contains the JSON metadata for the struct
// [PriceGroupedTieredPackagePrice]
type priceGroupedTieredPackagePriceJSON struct {
	ID                            apijson.Field
	BillableMetric                apijson.Field
	BillingCycleConfiguration     apijson.Field
	Cadence                       apijson.Field
	ConversionRate                apijson.Field
	ConversionRateConfig          apijson.Field
	CreatedAt                     apijson.Field
	CreditAllocation              apijson.Field
	Currency                      apijson.Field
	Discount                      apijson.Field
	ExternalPriceID               apijson.Field
	FixedPriceQuantity            apijson.Field
	GroupedTieredPackageConfig    apijson.Field
	InvoicingCycleConfiguration   apijson.Field
	Item                          apijson.Field
	Maximum                       apijson.Field
	MaximumAmount                 apijson.Field
	Metadata                      apijson.Field
	Minimum                       apijson.Field
	MinimumAmount                 apijson.Field
	ModelType                     apijson.Field
	Name                          apijson.Field
	PlanPhaseOrder                apijson.Field
	PriceType                     apijson.Field
	ReplacesPriceID               apijson.Field
	DimensionalPriceConfiguration apijson.Field
	raw                           string
	ExtraFields                   map[string]apijson.Field
}

func (r *PriceGroupedTieredPackagePrice) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r priceGroupedTieredPackagePriceJSON) RawJSON() string {
	return r.raw
}

func (r PriceGroupedTieredPackagePrice) implementsPrice() {}

type PriceGroupedTieredPackagePriceCadence string

const (
	PriceGroupedTieredPackagePriceCadenceOneTime    PriceGroupedTieredPackagePriceCadence = "one_time"
	PriceGroupedTieredPackagePriceCadenceMonthly    PriceGroupedTieredPackagePriceCadence = "monthly"
	PriceGroupedTieredPackagePriceCadenceQuarterly  PriceGroupedTieredPackagePriceCadence = "quarterly"
	PriceGroupedTieredPackagePriceCadenceSemiAnnual PriceGroupedTieredPackagePriceCadence = "semi_annual"
	PriceGroupedTieredPackagePriceCadenceAnnual     PriceGroupedTieredPackagePriceCadence = "annual"
	PriceGroupedTieredPackagePriceCadenceCustom     PriceGroupedTieredPackagePriceCadence = "custom"
)

func (r PriceGroupedTieredPackagePriceCadence) IsKnown() bool {
	switch r {
	case PriceGroupedTieredPackagePriceCadenceOneTime, PriceGroupedTieredPackagePriceCadenceMonthly, PriceGroupedTieredPackagePriceCadenceQuarterly, PriceGroupedTieredPackagePriceCadenceSemiAnnual, PriceGroupedTieredPackagePriceCadenceAnnual, PriceGroupedTieredPackagePriceCadenceCustom:
		return true
	}
	return false
}

type PriceGroupedTieredPackagePriceModelType string

const (
	PriceGroupedTieredPackagePriceModelTypeGroupedTieredPackage PriceGroupedTieredPackagePriceModelType = "grouped_tiered_package"
)

func (r PriceGroupedTieredPackagePriceModelType) IsKnown() bool {
	switch r {
	case PriceGroupedTieredPackagePriceModelTypeGroupedTieredPackage:
		return true
	}
	return false
}

type PriceGroupedTieredPackagePricePriceType string

const (
	PriceGroupedTieredPackagePricePriceTypeUsagePrice PriceGroupedTieredPackagePricePriceType = "usage_price"
	PriceGroupedTieredPackagePricePriceTypeFixedPrice PriceGroupedTieredPackagePricePriceType = "fixed_price"
)

func (r PriceGroupedTieredPackagePricePriceType) IsKnown() bool {
	switch r {
	case PriceGroupedTieredPackagePricePriceTypeUsagePrice, PriceGroupedTieredPackagePricePriceTypeFixedPrice:
		return true
	}
	return false
}

type PriceMaxGroupTieredPackagePrice struct {
	ID                        string                                 `json:"id,required"`
	BillableMetric            BillableMetricTiny                     `json:"billable_metric,required,nullable"`
	BillingCycleConfiguration BillingCycleConfiguration              `json:"billing_cycle_configuration,required"`
	Cadence                   PriceMaxGroupTieredPackagePriceCadence `json:"cadence,required"`
	ConversionRate            float64                                `json:"conversion_rate,required,nullable"`
	ConversionRateConfig      PriceMaxGroupTieredPackagePrice        `json:"conversion_rate_config,required,nullable"`
	CreatedAt                 time.Time                              `json:"created_at,required" format:"date-time"`
	CreditAllocation          Allocation                             `json:"credit_allocation,required,nullable"`
	Currency                  string                                 `json:"currency,required"`
	// Deprecated: deprecated
	Discount                    Discount                  `json:"discount,required,nullable"`
	ExternalPriceID             string                    `json:"external_price_id,required,nullable"`
	FixedPriceQuantity          float64                   `json:"fixed_price_quantity,required,nullable"`
	InvoicingCycleConfiguration BillingCycleConfiguration `json:"invoicing_cycle_configuration,required,nullable"`
	Item                        ItemSlim                  `json:"item,required"`
	MaxGroupTieredPackageConfig map[string]interface{}    `json:"max_group_tiered_package_config,required"`
	// Deprecated: deprecated
	Maximum Maximum `json:"maximum,required,nullable"`
	// Deprecated: deprecated
	MaximumAmount string `json:"maximum_amount,required,nullable"`
	// User specified key-value pairs for the resource. If not present, this defaults
	// to an empty dictionary. Individual keys can be removed by setting the value to
	// `null`, and the entire metadata mapping can be cleared by setting `metadata` to
	// `null`.
	Metadata map[string]string `json:"metadata,required"`
	// Deprecated: deprecated
	Minimum Minimum `json:"minimum,required,nullable"`
	// Deprecated: deprecated
	MinimumAmount  string                                   `json:"minimum_amount,required,nullable"`
	ModelType      PriceMaxGroupTieredPackagePriceModelType `json:"model_type,required"`
	Name           string                                   `json:"name,required"`
	PlanPhaseOrder int64                                    `json:"plan_phase_order,required,nullable"`
	PriceType      PriceMaxGroupTieredPackagePricePriceType `json:"price_type,required"`
	// The price id this price replaces. This price will take the place of the replaced
	// price in plan version migrations.
	ReplacesPriceID               string                              `json:"replaces_price_id,required,nullable"`
	DimensionalPriceConfiguration DimensionalPriceConfiguration       `json:"dimensional_price_configuration,nullable"`
	JSON                          priceMaxGroupTieredPackagePriceJSON `json:"-"`
}

// priceMaxGroupTieredPackagePriceJSON contains the JSON metadata for the struct
// [PriceMaxGroupTieredPackagePrice]
type priceMaxGroupTieredPackagePriceJSON struct {
	ID                            apijson.Field
	BillableMetric                apijson.Field
	BillingCycleConfiguration     apijson.Field
	Cadence                       apijson.Field
	ConversionRate                apijson.Field
	ConversionRateConfig          apijson.Field
	CreatedAt                     apijson.Field
	CreditAllocation              apijson.Field
	Currency                      apijson.Field
	Discount                      apijson.Field
	ExternalPriceID               apijson.Field
	FixedPriceQuantity            apijson.Field
	InvoicingCycleConfiguration   apijson.Field
	Item                          apijson.Field
	MaxGroupTieredPackageConfig   apijson.Field
	Maximum                       apijson.Field
	MaximumAmount                 apijson.Field
	Metadata                      apijson.Field
	Minimum                       apijson.Field
	MinimumAmount                 apijson.Field
	ModelType                     apijson.Field
	Name                          apijson.Field
	PlanPhaseOrder                apijson.Field
	PriceType                     apijson.Field
	ReplacesPriceID               apijson.Field
	DimensionalPriceConfiguration apijson.Field
	raw                           string
	ExtraFields                   map[string]apijson.Field
}

func (r *PriceMaxGroupTieredPackagePrice) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r priceMaxGroupTieredPackagePriceJSON) RawJSON() string {
	return r.raw
}

func (r PriceMaxGroupTieredPackagePrice) implementsPrice() {}

type PriceMaxGroupTieredPackagePriceCadence string

const (
	PriceMaxGroupTieredPackagePriceCadenceOneTime    PriceMaxGroupTieredPackagePriceCadence = "one_time"
	PriceMaxGroupTieredPackagePriceCadenceMonthly    PriceMaxGroupTieredPackagePriceCadence = "monthly"
	PriceMaxGroupTieredPackagePriceCadenceQuarterly  PriceMaxGroupTieredPackagePriceCadence = "quarterly"
	PriceMaxGroupTieredPackagePriceCadenceSemiAnnual PriceMaxGroupTieredPackagePriceCadence = "semi_annual"
	PriceMaxGroupTieredPackagePriceCadenceAnnual     PriceMaxGroupTieredPackagePriceCadence = "annual"
	PriceMaxGroupTieredPackagePriceCadenceCustom     PriceMaxGroupTieredPackagePriceCadence = "custom"
)

func (r PriceMaxGroupTieredPackagePriceCadence) IsKnown() bool {
	switch r {
	case PriceMaxGroupTieredPackagePriceCadenceOneTime, PriceMaxGroupTieredPackagePriceCadenceMonthly, PriceMaxGroupTieredPackagePriceCadenceQuarterly, PriceMaxGroupTieredPackagePriceCadenceSemiAnnual, PriceMaxGroupTieredPackagePriceCadenceAnnual, PriceMaxGroupTieredPackagePriceCadenceCustom:
		return true
	}
	return false
}

type PriceMaxGroupTieredPackagePriceModelType string

const (
	PriceMaxGroupTieredPackagePriceModelTypeMaxGroupTieredPackage PriceMaxGroupTieredPackagePriceModelType = "max_group_tiered_package"
)

func (r PriceMaxGroupTieredPackagePriceModelType) IsKnown() bool {
	switch r {
	case PriceMaxGroupTieredPackagePriceModelTypeMaxGroupTieredPackage:
		return true
	}
	return false
}

type PriceMaxGroupTieredPackagePricePriceType string

const (
	PriceMaxGroupTieredPackagePricePriceTypeUsagePrice PriceMaxGroupTieredPackagePricePriceType = "usage_price"
	PriceMaxGroupTieredPackagePricePriceTypeFixedPrice PriceMaxGroupTieredPackagePricePriceType = "fixed_price"
)

func (r PriceMaxGroupTieredPackagePricePriceType) IsKnown() bool {
	switch r {
	case PriceMaxGroupTieredPackagePricePriceTypeUsagePrice, PriceMaxGroupTieredPackagePricePriceTypeFixedPrice:
		return true
	}
	return false
}

type PriceScalableMatrixWithUnitPricingPrice struct {
	ID                        string                                         `json:"id,required"`
	BillableMetric            BillableMetricTiny                             `json:"billable_metric,required,nullable"`
	BillingCycleConfiguration BillingCycleConfiguration                      `json:"billing_cycle_configuration,required"`
	Cadence                   PriceScalableMatrixWithUnitPricingPriceCadence `json:"cadence,required"`
	ConversionRate            float64                                        `json:"conversion_rate,required,nullable"`
	ConversionRateConfig      PriceScalableMatrixWithUnitPricingPrice        `json:"conversion_rate_config,required,nullable"`
	CreatedAt                 time.Time                                      `json:"created_at,required" format:"date-time"`
	CreditAllocation          Allocation                                     `json:"credit_allocation,required,nullable"`
	Currency                  string                                         `json:"currency,required"`
	// Deprecated: deprecated
	Discount                    Discount                  `json:"discount,required,nullable"`
	ExternalPriceID             string                    `json:"external_price_id,required,nullable"`
	FixedPriceQuantity          float64                   `json:"fixed_price_quantity,required,nullable"`
	InvoicingCycleConfiguration BillingCycleConfiguration `json:"invoicing_cycle_configuration,required,nullable"`
	Item                        ItemSlim                  `json:"item,required"`
	// Deprecated: deprecated
	Maximum Maximum `json:"maximum,required,nullable"`
	// Deprecated: deprecated
	MaximumAmount string `json:"maximum_amount,required,nullable"`
	// User specified key-value pairs for the resource. If not present, this defaults
	// to an empty dictionary. Individual keys can be removed by setting the value to
	// `null`, and the entire metadata mapping can be cleared by setting `metadata` to
	// `null`.
	Metadata map[string]string `json:"metadata,required"`
	// Deprecated: deprecated
	Minimum Minimum `json:"minimum,required,nullable"`
	// Deprecated: deprecated
	MinimumAmount  string                                           `json:"minimum_amount,required,nullable"`
	ModelType      PriceScalableMatrixWithUnitPricingPriceModelType `json:"model_type,required"`
	Name           string                                           `json:"name,required"`
	PlanPhaseOrder int64                                            `json:"plan_phase_order,required,nullable"`
	PriceType      PriceScalableMatrixWithUnitPricingPricePriceType `json:"price_type,required"`
	// The price id this price replaces. This price will take the place of the replaced
	// price in plan version migrations.
	ReplacesPriceID                     string                                      `json:"replaces_price_id,required,nullable"`
	ScalableMatrixWithUnitPricingConfig map[string]interface{}                      `json:"scalable_matrix_with_unit_pricing_config,required"`
	DimensionalPriceConfiguration       DimensionalPriceConfiguration               `json:"dimensional_price_configuration,nullable"`
	JSON                                priceScalableMatrixWithUnitPricingPriceJSON `json:"-"`
}

// priceScalableMatrixWithUnitPricingPriceJSON contains the JSON metadata for the
// struct [PriceScalableMatrixWithUnitPricingPrice]
type priceScalableMatrixWithUnitPricingPriceJSON struct {
	ID                                  apijson.Field
	BillableMetric                      apijson.Field
	BillingCycleConfiguration           apijson.Field
	Cadence                             apijson.Field
	ConversionRate                      apijson.Field
	ConversionRateConfig                apijson.Field
	CreatedAt                           apijson.Field
	CreditAllocation                    apijson.Field
	Currency                            apijson.Field
	Discount                            apijson.Field
	ExternalPriceID                     apijson.Field
	FixedPriceQuantity                  apijson.Field
	InvoicingCycleConfiguration         apijson.Field
	Item                                apijson.Field
	Maximum                             apijson.Field
	MaximumAmount                       apijson.Field
	Metadata                            apijson.Field
	Minimum                             apijson.Field
	MinimumAmount                       apijson.Field
	ModelType                           apijson.Field
	Name                                apijson.Field
	PlanPhaseOrder                      apijson.Field
	PriceType                           apijson.Field
	ReplacesPriceID                     apijson.Field
	ScalableMatrixWithUnitPricingConfig apijson.Field
	DimensionalPriceConfiguration       apijson.Field
	raw                                 string
	ExtraFields                         map[string]apijson.Field
}

func (r *PriceScalableMatrixWithUnitPricingPrice) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r priceScalableMatrixWithUnitPricingPriceJSON) RawJSON() string {
	return r.raw
}

func (r PriceScalableMatrixWithUnitPricingPrice) implementsPrice() {}

type PriceScalableMatrixWithUnitPricingPriceCadence string

const (
	PriceScalableMatrixWithUnitPricingPriceCadenceOneTime    PriceScalableMatrixWithUnitPricingPriceCadence = "one_time"
	PriceScalableMatrixWithUnitPricingPriceCadenceMonthly    PriceScalableMatrixWithUnitPricingPriceCadence = "monthly"
	PriceScalableMatrixWithUnitPricingPriceCadenceQuarterly  PriceScalableMatrixWithUnitPricingPriceCadence = "quarterly"
	PriceScalableMatrixWithUnitPricingPriceCadenceSemiAnnual PriceScalableMatrixWithUnitPricingPriceCadence = "semi_annual"
	PriceScalableMatrixWithUnitPricingPriceCadenceAnnual     PriceScalableMatrixWithUnitPricingPriceCadence = "annual"
	PriceScalableMatrixWithUnitPricingPriceCadenceCustom     PriceScalableMatrixWithUnitPricingPriceCadence = "custom"
)

func (r PriceScalableMatrixWithUnitPricingPriceCadence) IsKnown() bool {
	switch r {
	case PriceScalableMatrixWithUnitPricingPriceCadenceOneTime, PriceScalableMatrixWithUnitPricingPriceCadenceMonthly, PriceScalableMatrixWithUnitPricingPriceCadenceQuarterly, PriceScalableMatrixWithUnitPricingPriceCadenceSemiAnnual, PriceScalableMatrixWithUnitPricingPriceCadenceAnnual, PriceScalableMatrixWithUnitPricingPriceCadenceCustom:
		return true
	}
	return false
}

type PriceScalableMatrixWithUnitPricingPriceModelType string

const (
	PriceScalableMatrixWithUnitPricingPriceModelTypeScalableMatrixWithUnitPricing PriceScalableMatrixWithUnitPricingPriceModelType = "scalable_matrix_with_unit_pricing"
)

func (r PriceScalableMatrixWithUnitPricingPriceModelType) IsKnown() bool {
	switch r {
	case PriceScalableMatrixWithUnitPricingPriceModelTypeScalableMatrixWithUnitPricing:
		return true
	}
	return false
}

type PriceScalableMatrixWithUnitPricingPricePriceType string

const (
	PriceScalableMatrixWithUnitPricingPricePriceTypeUsagePrice PriceScalableMatrixWithUnitPricingPricePriceType = "usage_price"
	PriceScalableMatrixWithUnitPricingPricePriceTypeFixedPrice PriceScalableMatrixWithUnitPricingPricePriceType = "fixed_price"
)

func (r PriceScalableMatrixWithUnitPricingPricePriceType) IsKnown() bool {
	switch r {
	case PriceScalableMatrixWithUnitPricingPricePriceTypeUsagePrice, PriceScalableMatrixWithUnitPricingPricePriceTypeFixedPrice:
		return true
	}
	return false
}

type PriceScalableMatrixWithTieredPricingPrice struct {
	ID                        string                                           `json:"id,required"`
	BillableMetric            BillableMetricTiny                               `json:"billable_metric,required,nullable"`
	BillingCycleConfiguration BillingCycleConfiguration                        `json:"billing_cycle_configuration,required"`
	Cadence                   PriceScalableMatrixWithTieredPricingPriceCadence `json:"cadence,required"`
	ConversionRate            float64                                          `json:"conversion_rate,required,nullable"`
	ConversionRateConfig      PriceScalableMatrixWithTieredPricingPrice        `json:"conversion_rate_config,required,nullable"`
	CreatedAt                 time.Time                                        `json:"created_at,required" format:"date-time"`
	CreditAllocation          Allocation                                       `json:"credit_allocation,required,nullable"`
	Currency                  string                                           `json:"currency,required"`
	// Deprecated: deprecated
	Discount                    Discount                  `json:"discount,required,nullable"`
	ExternalPriceID             string                    `json:"external_price_id,required,nullable"`
	FixedPriceQuantity          float64                   `json:"fixed_price_quantity,required,nullable"`
	InvoicingCycleConfiguration BillingCycleConfiguration `json:"invoicing_cycle_configuration,required,nullable"`
	Item                        ItemSlim                  `json:"item,required"`
	// Deprecated: deprecated
	Maximum Maximum `json:"maximum,required,nullable"`
	// Deprecated: deprecated
	MaximumAmount string `json:"maximum_amount,required,nullable"`
	// User specified key-value pairs for the resource. If not present, this defaults
	// to an empty dictionary. Individual keys can be removed by setting the value to
	// `null`, and the entire metadata mapping can be cleared by setting `metadata` to
	// `null`.
	Metadata map[string]string `json:"metadata,required"`
	// Deprecated: deprecated
	Minimum Minimum `json:"minimum,required,nullable"`
	// Deprecated: deprecated
	MinimumAmount  string                                             `json:"minimum_amount,required,nullable"`
	ModelType      PriceScalableMatrixWithTieredPricingPriceModelType `json:"model_type,required"`
	Name           string                                             `json:"name,required"`
	PlanPhaseOrder int64                                              `json:"plan_phase_order,required,nullable"`
	PriceType      PriceScalableMatrixWithTieredPricingPricePriceType `json:"price_type,required"`
	// The price id this price replaces. This price will take the place of the replaced
	// price in plan version migrations.
	ReplacesPriceID                       string                                        `json:"replaces_price_id,required,nullable"`
	ScalableMatrixWithTieredPricingConfig map[string]interface{}                        `json:"scalable_matrix_with_tiered_pricing_config,required"`
	DimensionalPriceConfiguration         DimensionalPriceConfiguration                 `json:"dimensional_price_configuration,nullable"`
	JSON                                  priceScalableMatrixWithTieredPricingPriceJSON `json:"-"`
}

// priceScalableMatrixWithTieredPricingPriceJSON contains the JSON metadata for the
// struct [PriceScalableMatrixWithTieredPricingPrice]
type priceScalableMatrixWithTieredPricingPriceJSON struct {
	ID                                    apijson.Field
	BillableMetric                        apijson.Field
	BillingCycleConfiguration             apijson.Field
	Cadence                               apijson.Field
	ConversionRate                        apijson.Field
	ConversionRateConfig                  apijson.Field
	CreatedAt                             apijson.Field
	CreditAllocation                      apijson.Field
	Currency                              apijson.Field
	Discount                              apijson.Field
	ExternalPriceID                       apijson.Field
	FixedPriceQuantity                    apijson.Field
	InvoicingCycleConfiguration           apijson.Field
	Item                                  apijson.Field
	Maximum                               apijson.Field
	MaximumAmount                         apijson.Field
	Metadata                              apijson.Field
	Minimum                               apijson.Field
	MinimumAmount                         apijson.Field
	ModelType                             apijson.Field
	Name                                  apijson.Field
	PlanPhaseOrder                        apijson.Field
	PriceType                             apijson.Field
	ReplacesPriceID                       apijson.Field
	ScalableMatrixWithTieredPricingConfig apijson.Field
	DimensionalPriceConfiguration         apijson.Field
	raw                                   string
	ExtraFields                           map[string]apijson.Field
}

func (r *PriceScalableMatrixWithTieredPricingPrice) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r priceScalableMatrixWithTieredPricingPriceJSON) RawJSON() string {
	return r.raw
}

func (r PriceScalableMatrixWithTieredPricingPrice) implementsPrice() {}

type PriceScalableMatrixWithTieredPricingPriceCadence string

const (
	PriceScalableMatrixWithTieredPricingPriceCadenceOneTime    PriceScalableMatrixWithTieredPricingPriceCadence = "one_time"
	PriceScalableMatrixWithTieredPricingPriceCadenceMonthly    PriceScalableMatrixWithTieredPricingPriceCadence = "monthly"
	PriceScalableMatrixWithTieredPricingPriceCadenceQuarterly  PriceScalableMatrixWithTieredPricingPriceCadence = "quarterly"
	PriceScalableMatrixWithTieredPricingPriceCadenceSemiAnnual PriceScalableMatrixWithTieredPricingPriceCadence = "semi_annual"
	PriceScalableMatrixWithTieredPricingPriceCadenceAnnual     PriceScalableMatrixWithTieredPricingPriceCadence = "annual"
	PriceScalableMatrixWithTieredPricingPriceCadenceCustom     PriceScalableMatrixWithTieredPricingPriceCadence = "custom"
)

func (r PriceScalableMatrixWithTieredPricingPriceCadence) IsKnown() bool {
	switch r {
	case PriceScalableMatrixWithTieredPricingPriceCadenceOneTime, PriceScalableMatrixWithTieredPricingPriceCadenceMonthly, PriceScalableMatrixWithTieredPricingPriceCadenceQuarterly, PriceScalableMatrixWithTieredPricingPriceCadenceSemiAnnual, PriceScalableMatrixWithTieredPricingPriceCadenceAnnual, PriceScalableMatrixWithTieredPricingPriceCadenceCustom:
		return true
	}
	return false
}

type PriceScalableMatrixWithTieredPricingPriceModelType string

const (
	PriceScalableMatrixWithTieredPricingPriceModelTypeScalableMatrixWithTieredPricing PriceScalableMatrixWithTieredPricingPriceModelType = "scalable_matrix_with_tiered_pricing"
)

func (r PriceScalableMatrixWithTieredPricingPriceModelType) IsKnown() bool {
	switch r {
	case PriceScalableMatrixWithTieredPricingPriceModelTypeScalableMatrixWithTieredPricing:
		return true
	}
	return false
}

type PriceScalableMatrixWithTieredPricingPricePriceType string

const (
	PriceScalableMatrixWithTieredPricingPricePriceTypeUsagePrice PriceScalableMatrixWithTieredPricingPricePriceType = "usage_price"
	PriceScalableMatrixWithTieredPricingPricePriceTypeFixedPrice PriceScalableMatrixWithTieredPricingPricePriceType = "fixed_price"
)

func (r PriceScalableMatrixWithTieredPricingPricePriceType) IsKnown() bool {
	switch r {
	case PriceScalableMatrixWithTieredPricingPricePriceTypeUsagePrice, PriceScalableMatrixWithTieredPricingPricePriceTypeFixedPrice:
		return true
	}
	return false
}

type PriceCumulativeGroupedBulkPrice struct {
	ID                          string                                 `json:"id,required"`
	BillableMetric              BillableMetricTiny                     `json:"billable_metric,required,nullable"`
	BillingCycleConfiguration   BillingCycleConfiguration              `json:"billing_cycle_configuration,required"`
	Cadence                     PriceCumulativeGroupedBulkPriceCadence `json:"cadence,required"`
	ConversionRate              float64                                `json:"conversion_rate,required,nullable"`
	ConversionRateConfig        PriceCumulativeGroupedBulkPrice        `json:"conversion_rate_config,required,nullable"`
	CreatedAt                   time.Time                              `json:"created_at,required" format:"date-time"`
	CreditAllocation            Allocation                             `json:"credit_allocation,required,nullable"`
	CumulativeGroupedBulkConfig map[string]interface{}                 `json:"cumulative_grouped_bulk_config,required"`
	Currency                    string                                 `json:"currency,required"`
	// Deprecated: deprecated
	Discount                    Discount                  `json:"discount,required,nullable"`
	ExternalPriceID             string                    `json:"external_price_id,required,nullable"`
	FixedPriceQuantity          float64                   `json:"fixed_price_quantity,required,nullable"`
	InvoicingCycleConfiguration BillingCycleConfiguration `json:"invoicing_cycle_configuration,required,nullable"`
	Item                        ItemSlim                  `json:"item,required"`
	// Deprecated: deprecated
	Maximum Maximum `json:"maximum,required,nullable"`
	// Deprecated: deprecated
	MaximumAmount string `json:"maximum_amount,required,nullable"`
	// User specified key-value pairs for the resource. If not present, this defaults
	// to an empty dictionary. Individual keys can be removed by setting the value to
	// `null`, and the entire metadata mapping can be cleared by setting `metadata` to
	// `null`.
	Metadata map[string]string `json:"metadata,required"`
	// Deprecated: deprecated
	Minimum Minimum `json:"minimum,required,nullable"`
	// Deprecated: deprecated
	MinimumAmount  string                                   `json:"minimum_amount,required,nullable"`
	ModelType      PriceCumulativeGroupedBulkPriceModelType `json:"model_type,required"`
	Name           string                                   `json:"name,required"`
	PlanPhaseOrder int64                                    `json:"plan_phase_order,required,nullable"`
	PriceType      PriceCumulativeGroupedBulkPricePriceType `json:"price_type,required"`
	// The price id this price replaces. This price will take the place of the replaced
	// price in plan version migrations.
	ReplacesPriceID               string                              `json:"replaces_price_id,required,nullable"`
	DimensionalPriceConfiguration DimensionalPriceConfiguration       `json:"dimensional_price_configuration,nullable"`
	JSON                          priceCumulativeGroupedBulkPriceJSON `json:"-"`
}

// priceCumulativeGroupedBulkPriceJSON contains the JSON metadata for the struct
// [PriceCumulativeGroupedBulkPrice]
type priceCumulativeGroupedBulkPriceJSON struct {
	ID                            apijson.Field
	BillableMetric                apijson.Field
	BillingCycleConfiguration     apijson.Field
	Cadence                       apijson.Field
	ConversionRate                apijson.Field
	ConversionRateConfig          apijson.Field
	CreatedAt                     apijson.Field
	CreditAllocation              apijson.Field
	CumulativeGroupedBulkConfig   apijson.Field
	Currency                      apijson.Field
	Discount                      apijson.Field
	ExternalPriceID               apijson.Field
	FixedPriceQuantity            apijson.Field
	InvoicingCycleConfiguration   apijson.Field
	Item                          apijson.Field
	Maximum                       apijson.Field
	MaximumAmount                 apijson.Field
	Metadata                      apijson.Field
	Minimum                       apijson.Field
	MinimumAmount                 apijson.Field
	ModelType                     apijson.Field
	Name                          apijson.Field
	PlanPhaseOrder                apijson.Field
	PriceType                     apijson.Field
	ReplacesPriceID               apijson.Field
	DimensionalPriceConfiguration apijson.Field
	raw                           string
	ExtraFields                   map[string]apijson.Field
}

func (r *PriceCumulativeGroupedBulkPrice) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r priceCumulativeGroupedBulkPriceJSON) RawJSON() string {
	return r.raw
}

func (r PriceCumulativeGroupedBulkPrice) implementsPrice() {}

type PriceCumulativeGroupedBulkPriceCadence string

const (
	PriceCumulativeGroupedBulkPriceCadenceOneTime    PriceCumulativeGroupedBulkPriceCadence = "one_time"
	PriceCumulativeGroupedBulkPriceCadenceMonthly    PriceCumulativeGroupedBulkPriceCadence = "monthly"
	PriceCumulativeGroupedBulkPriceCadenceQuarterly  PriceCumulativeGroupedBulkPriceCadence = "quarterly"
	PriceCumulativeGroupedBulkPriceCadenceSemiAnnual PriceCumulativeGroupedBulkPriceCadence = "semi_annual"
	PriceCumulativeGroupedBulkPriceCadenceAnnual     PriceCumulativeGroupedBulkPriceCadence = "annual"
	PriceCumulativeGroupedBulkPriceCadenceCustom     PriceCumulativeGroupedBulkPriceCadence = "custom"
)

func (r PriceCumulativeGroupedBulkPriceCadence) IsKnown() bool {
	switch r {
	case PriceCumulativeGroupedBulkPriceCadenceOneTime, PriceCumulativeGroupedBulkPriceCadenceMonthly, PriceCumulativeGroupedBulkPriceCadenceQuarterly, PriceCumulativeGroupedBulkPriceCadenceSemiAnnual, PriceCumulativeGroupedBulkPriceCadenceAnnual, PriceCumulativeGroupedBulkPriceCadenceCustom:
		return true
	}
	return false
}

type PriceCumulativeGroupedBulkPriceModelType string

const (
	PriceCumulativeGroupedBulkPriceModelTypeCumulativeGroupedBulk PriceCumulativeGroupedBulkPriceModelType = "cumulative_grouped_bulk"
)

func (r PriceCumulativeGroupedBulkPriceModelType) IsKnown() bool {
	switch r {
	case PriceCumulativeGroupedBulkPriceModelTypeCumulativeGroupedBulk:
		return true
	}
	return false
}

type PriceCumulativeGroupedBulkPricePriceType string

const (
	PriceCumulativeGroupedBulkPricePriceTypeUsagePrice PriceCumulativeGroupedBulkPricePriceType = "usage_price"
	PriceCumulativeGroupedBulkPricePriceTypeFixedPrice PriceCumulativeGroupedBulkPricePriceType = "fixed_price"
)

func (r PriceCumulativeGroupedBulkPricePriceType) IsKnown() bool {
	switch r {
	case PriceCumulativeGroupedBulkPricePriceTypeUsagePrice, PriceCumulativeGroupedBulkPricePriceTypeFixedPrice:
		return true
	}
	return false
}

type PriceGroupedWithMinMaxThresholdsPrice struct {
	ID                        string                                       `json:"id,required"`
	BillableMetric            BillableMetricTiny                           `json:"billable_metric,required,nullable"`
	BillingCycleConfiguration BillingCycleConfiguration                    `json:"billing_cycle_configuration,required"`
	Cadence                   PriceGroupedWithMinMaxThresholdsPriceCadence `json:"cadence,required"`
	ConversionRate            float64                                      `json:"conversion_rate,required,nullable"`
	ConversionRateConfig      PriceGroupedWithMinMaxThresholdsPrice        `json:"conversion_rate_config,required,nullable"`
	CreatedAt                 time.Time                                    `json:"created_at,required" format:"date-time"`
	CreditAllocation          Allocation                                   `json:"credit_allocation,required,nullable"`
	Currency                  string                                       `json:"currency,required"`
	// Deprecated: deprecated
	Discount                          Discount                  `json:"discount,required,nullable"`
	ExternalPriceID                   string                    `json:"external_price_id,required,nullable"`
	FixedPriceQuantity                float64                   `json:"fixed_price_quantity,required,nullable"`
	GroupedWithMinMaxThresholdsConfig map[string]interface{}    `json:"grouped_with_min_max_thresholds_config,required"`
	InvoicingCycleConfiguration       BillingCycleConfiguration `json:"invoicing_cycle_configuration,required,nullable"`
	Item                              ItemSlim                  `json:"item,required"`
	// Deprecated: deprecated
	Maximum Maximum `json:"maximum,required,nullable"`
	// Deprecated: deprecated
	MaximumAmount string `json:"maximum_amount,required,nullable"`
	// User specified key-value pairs for the resource. If not present, this defaults
	// to an empty dictionary. Individual keys can be removed by setting the value to
	// `null`, and the entire metadata mapping can be cleared by setting `metadata` to
	// `null`.
	Metadata map[string]string `json:"metadata,required"`
	// Deprecated: deprecated
	Minimum Minimum `json:"minimum,required,nullable"`
	// Deprecated: deprecated
	MinimumAmount  string                                         `json:"minimum_amount,required,nullable"`
	ModelType      PriceGroupedWithMinMaxThresholdsPriceModelType `json:"model_type,required"`
	Name           string                                         `json:"name,required"`
	PlanPhaseOrder int64                                          `json:"plan_phase_order,required,nullable"`
	PriceType      PriceGroupedWithMinMaxThresholdsPricePriceType `json:"price_type,required"`
	// The price id this price replaces. This price will take the place of the replaced
	// price in plan version migrations.
	ReplacesPriceID               string                                    `json:"replaces_price_id,required,nullable"`
	DimensionalPriceConfiguration DimensionalPriceConfiguration             `json:"dimensional_price_configuration,nullable"`
	JSON                          priceGroupedWithMinMaxThresholdsPriceJSON `json:"-"`
}

// priceGroupedWithMinMaxThresholdsPriceJSON contains the JSON metadata for the
// struct [PriceGroupedWithMinMaxThresholdsPrice]
type priceGroupedWithMinMaxThresholdsPriceJSON struct {
	ID                                apijson.Field
	BillableMetric                    apijson.Field
	BillingCycleConfiguration         apijson.Field
	Cadence                           apijson.Field
	ConversionRate                    apijson.Field
	ConversionRateConfig              apijson.Field
	CreatedAt                         apijson.Field
	CreditAllocation                  apijson.Field
	Currency                          apijson.Field
	Discount                          apijson.Field
	ExternalPriceID                   apijson.Field
	FixedPriceQuantity                apijson.Field
	GroupedWithMinMaxThresholdsConfig apijson.Field
	InvoicingCycleConfiguration       apijson.Field
	Item                              apijson.Field
	Maximum                           apijson.Field
	MaximumAmount                     apijson.Field
	Metadata                          apijson.Field
	Minimum                           apijson.Field
	MinimumAmount                     apijson.Field
	ModelType                         apijson.Field
	Name                              apijson.Field
	PlanPhaseOrder                    apijson.Field
	PriceType                         apijson.Field
	ReplacesPriceID                   apijson.Field
	DimensionalPriceConfiguration     apijson.Field
	raw                               string
	ExtraFields                       map[string]apijson.Field
}

func (r *PriceGroupedWithMinMaxThresholdsPrice) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r priceGroupedWithMinMaxThresholdsPriceJSON) RawJSON() string {
	return r.raw
}

func (r PriceGroupedWithMinMaxThresholdsPrice) implementsPrice() {}

type PriceGroupedWithMinMaxThresholdsPriceCadence string

const (
	PriceGroupedWithMinMaxThresholdsPriceCadenceOneTime    PriceGroupedWithMinMaxThresholdsPriceCadence = "one_time"
	PriceGroupedWithMinMaxThresholdsPriceCadenceMonthly    PriceGroupedWithMinMaxThresholdsPriceCadence = "monthly"
	PriceGroupedWithMinMaxThresholdsPriceCadenceQuarterly  PriceGroupedWithMinMaxThresholdsPriceCadence = "quarterly"
	PriceGroupedWithMinMaxThresholdsPriceCadenceSemiAnnual PriceGroupedWithMinMaxThresholdsPriceCadence = "semi_annual"
	PriceGroupedWithMinMaxThresholdsPriceCadenceAnnual     PriceGroupedWithMinMaxThresholdsPriceCadence = "annual"
	PriceGroupedWithMinMaxThresholdsPriceCadenceCustom     PriceGroupedWithMinMaxThresholdsPriceCadence = "custom"
)

func (r PriceGroupedWithMinMaxThresholdsPriceCadence) IsKnown() bool {
	switch r {
	case PriceGroupedWithMinMaxThresholdsPriceCadenceOneTime, PriceGroupedWithMinMaxThresholdsPriceCadenceMonthly, PriceGroupedWithMinMaxThresholdsPriceCadenceQuarterly, PriceGroupedWithMinMaxThresholdsPriceCadenceSemiAnnual, PriceGroupedWithMinMaxThresholdsPriceCadenceAnnual, PriceGroupedWithMinMaxThresholdsPriceCadenceCustom:
		return true
	}
	return false
}

type PriceGroupedWithMinMaxThresholdsPriceModelType string

const (
	PriceGroupedWithMinMaxThresholdsPriceModelTypeGroupedWithMinMaxThresholds PriceGroupedWithMinMaxThresholdsPriceModelType = "grouped_with_min_max_thresholds"
)

func (r PriceGroupedWithMinMaxThresholdsPriceModelType) IsKnown() bool {
	switch r {
	case PriceGroupedWithMinMaxThresholdsPriceModelTypeGroupedWithMinMaxThresholds:
		return true
	}
	return false
}

type PriceGroupedWithMinMaxThresholdsPricePriceType string

const (
	PriceGroupedWithMinMaxThresholdsPricePriceTypeUsagePrice PriceGroupedWithMinMaxThresholdsPricePriceType = "usage_price"
	PriceGroupedWithMinMaxThresholdsPricePriceTypeFixedPrice PriceGroupedWithMinMaxThresholdsPricePriceType = "fixed_price"
)

func (r PriceGroupedWithMinMaxThresholdsPricePriceType) IsKnown() bool {
	switch r {
	case PriceGroupedWithMinMaxThresholdsPricePriceTypeUsagePrice, PriceGroupedWithMinMaxThresholdsPricePriceTypeFixedPrice:
		return true
	}
	return false
}

type PriceCadence string

const (
	PriceCadenceOneTime    PriceCadence = "one_time"
	PriceCadenceMonthly    PriceCadence = "monthly"
	PriceCadenceQuarterly  PriceCadence = "quarterly"
	PriceCadenceSemiAnnual PriceCadence = "semi_annual"
	PriceCadenceAnnual     PriceCadence = "annual"
	PriceCadenceCustom     PriceCadence = "custom"
)

func (r PriceCadence) IsKnown() bool {
	switch r {
	case PriceCadenceOneTime, PriceCadenceMonthly, PriceCadenceQuarterly, PriceCadenceSemiAnnual, PriceCadenceAnnual, PriceCadenceCustom:
		return true
	}
	return false
}

type PriceModelType string

const (
	PriceModelTypeUnit                            PriceModelType = "unit"
	PriceModelTypePackage                         PriceModelType = "package"
	PriceModelTypeMatrix                          PriceModelType = "matrix"
	PriceModelTypeTiered                          PriceModelType = "tiered"
	PriceModelTypeTieredBPS                       PriceModelType = "tiered_bps"
	PriceModelTypeBPS                             PriceModelType = "bps"
	PriceModelTypeBulkBPS                         PriceModelType = "bulk_bps"
	PriceModelTypeBulk                            PriceModelType = "bulk"
	PriceModelTypeThresholdTotalAmount            PriceModelType = "threshold_total_amount"
	PriceModelTypeTieredPackage                   PriceModelType = "tiered_package"
	PriceModelTypeGroupedTiered                   PriceModelType = "grouped_tiered"
	PriceModelTypeTieredWithMinimum               PriceModelType = "tiered_with_minimum"
	PriceModelTypeTieredPackageWithMinimum        PriceModelType = "tiered_package_with_minimum"
	PriceModelTypePackageWithAllocation           PriceModelType = "package_with_allocation"
	PriceModelTypeUnitWithPercent                 PriceModelType = "unit_with_percent"
	PriceModelTypeMatrixWithAllocation            PriceModelType = "matrix_with_allocation"
	PriceModelTypeTieredWithProration             PriceModelType = "tiered_with_proration"
	PriceModelTypeUnitWithProration               PriceModelType = "unit_with_proration"
	PriceModelTypeGroupedAllocation               PriceModelType = "grouped_allocation"
	PriceModelTypeGroupedWithProratedMinimum      PriceModelType = "grouped_with_prorated_minimum"
	PriceModelTypeGroupedWithMeteredMinimum       PriceModelType = "grouped_with_metered_minimum"
	PriceModelTypeMatrixWithDisplayName           PriceModelType = "matrix_with_display_name"
	PriceModelTypeBulkWithProration               PriceModelType = "bulk_with_proration"
	PriceModelTypeGroupedTieredPackage            PriceModelType = "grouped_tiered_package"
	PriceModelTypeMaxGroupTieredPackage           PriceModelType = "max_group_tiered_package"
	PriceModelTypeScalableMatrixWithUnitPricing   PriceModelType = "scalable_matrix_with_unit_pricing"
	PriceModelTypeScalableMatrixWithTieredPricing PriceModelType = "scalable_matrix_with_tiered_pricing"
	PriceModelTypeCumulativeGroupedBulk           PriceModelType = "cumulative_grouped_bulk"
	PriceModelTypeGroupedWithMinMaxThresholds     PriceModelType = "grouped_with_min_max_thresholds"
)

func (r PriceModelType) IsKnown() bool {
	switch r {
	case PriceModelTypeUnit, PriceModelTypePackage, PriceModelTypeMatrix, PriceModelTypeTiered, PriceModelTypeTieredBPS, PriceModelTypeBPS, PriceModelTypeBulkBPS, PriceModelTypeBulk, PriceModelTypeThresholdTotalAmount, PriceModelTypeTieredPackage, PriceModelTypeGroupedTiered, PriceModelTypeTieredWithMinimum, PriceModelTypeTieredPackageWithMinimum, PriceModelTypePackageWithAllocation, PriceModelTypeUnitWithPercent, PriceModelTypeMatrixWithAllocation, PriceModelTypeTieredWithProration, PriceModelTypeUnitWithProration, PriceModelTypeGroupedAllocation, PriceModelTypeGroupedWithProratedMinimum, PriceModelTypeGroupedWithMeteredMinimum, PriceModelTypeMatrixWithDisplayName, PriceModelTypeBulkWithProration, PriceModelTypeGroupedTieredPackage, PriceModelTypeMaxGroupTieredPackage, PriceModelTypeScalableMatrixWithUnitPricing, PriceModelTypeScalableMatrixWithTieredPricing, PriceModelTypeCumulativeGroupedBulk, PriceModelTypeGroupedWithMinMaxThresholds:
		return true
	}
	return false
}

type PricePriceType string

const (
	PricePriceTypeUsagePrice PricePriceType = "usage_price"
	PricePriceTypeFixedPrice PricePriceType = "fixed_price"
)

func (r PricePriceType) IsKnown() bool {
	switch r {
	case PricePriceTypeUsagePrice, PricePriceTypeFixedPrice:
		return true
	}
	return false
}

// The Price Interval resource represents a period of time for which a price will
// bill on a subscription. A subscription’s price intervals define its billing
// behavior.
type PriceInterval struct {
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
	FixedFeeQuantityTransitions []FixedFeeQuantityTransition `json:"fixed_fee_quantity_transitions,required,nullable"`
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
	UsageCustomerIDs []string          `json:"usage_customer_ids,required,nullable"`
	JSON             priceIntervalJSON `json:"-"`
}

// priceIntervalJSON contains the JSON metadata for the struct [PriceInterval]
type priceIntervalJSON struct {
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

func (r *PriceInterval) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r priceIntervalJSON) RawJSON() string {
	return r.raw
}

type SubLineItemGrouping struct {
	Key string `json:"key,required"`
	// No value indicates the default group
	Value string                  `json:"value,required,nullable"`
	JSON  subLineItemGroupingJSON `json:"-"`
}

// subLineItemGroupingJSON contains the JSON metadata for the struct
// [SubLineItemGrouping]
type subLineItemGroupingJSON struct {
	Key         apijson.Field
	Value       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *SubLineItemGrouping) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r subLineItemGroupingJSON) RawJSON() string {
	return r.raw
}

type SubLineItemMatrixConfig struct {
	// The ordered dimension values for this line item.
	DimensionValues []string                    `json:"dimension_values,required"`
	JSON            subLineItemMatrixConfigJSON `json:"-"`
}

// subLineItemMatrixConfigJSON contains the JSON metadata for the struct
// [SubLineItemMatrixConfig]
type subLineItemMatrixConfigJSON struct {
	DimensionValues apijson.Field
	raw             string
	ExtraFields     map[string]apijson.Field
}

func (r *SubLineItemMatrixConfig) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r subLineItemMatrixConfigJSON) RawJSON() string {
	return r.raw
}

type SubscriptionChangeMinified struct {
	ID   string                         `json:"id,required"`
	JSON subscriptionChangeMinifiedJSON `json:"-"`
}

// subscriptionChangeMinifiedJSON contains the JSON metadata for the struct
// [SubscriptionChangeMinified]
type subscriptionChangeMinifiedJSON struct {
	ID          apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *SubscriptionChangeMinified) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r subscriptionChangeMinifiedJSON) RawJSON() string {
	return r.raw
}

type SubscriptionMinified struct {
	ID   string                   `json:"id,required"`
	JSON subscriptionMinifiedJSON `json:"-"`
}

// subscriptionMinifiedJSON contains the JSON metadata for the struct
// [SubscriptionMinified]
type subscriptionMinifiedJSON struct {
	ID          apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *SubscriptionMinified) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r subscriptionMinifiedJSON) RawJSON() string {
	return r.raw
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

type TaxAmount struct {
	// The amount of additional tax incurred by this tax rate.
	Amount string `json:"amount,required"`
	// The human-readable description of the applied tax rate.
	TaxRateDescription string `json:"tax_rate_description,required"`
	// The tax rate percentage, out of 100.
	TaxRatePercentage string        `json:"tax_rate_percentage,required,nullable"`
	JSON              taxAmountJSON `json:"-"`
}

// taxAmountJSON contains the JSON metadata for the struct [TaxAmount]
type taxAmountJSON struct {
	Amount             apijson.Field
	TaxRateDescription apijson.Field
	TaxRatePercentage  apijson.Field
	raw                string
	ExtraFields        map[string]apijson.Field
}

func (r *TaxAmount) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r taxAmountJSON) RawJSON() string {
	return r.raw
}

type Tier struct {
	// Exclusive tier starting value
	FirstUnit float64 `json:"first_unit,required"`
	// Amount per unit
	UnitAmount string `json:"unit_amount,required"`
	// Inclusive tier ending value. If null, this is treated as the last tier
	LastUnit float64  `json:"last_unit,nullable"`
	JSON     tierJSON `json:"-"`
}

// tierJSON contains the JSON metadata for the struct [Tier]
type tierJSON struct {
	FirstUnit   apijson.Field
	UnitAmount  apijson.Field
	LastUnit    apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *Tier) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r tierJSON) RawJSON() string {
	return r.raw
}

type TierParam struct {
	// Exclusive tier starting value
	FirstUnit param.Field[float64] `json:"first_unit,required"`
	// Amount per unit
	UnitAmount param.Field[string] `json:"unit_amount,required"`
	// Inclusive tier ending value. If null, this is treated as the last tier
	LastUnit param.Field[float64] `json:"last_unit"`
}

func (r TierParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type TierConfig struct {
	FirstUnit  float64        `json:"first_unit,required"`
	LastUnit   float64        `json:"last_unit,required,nullable"`
	UnitAmount string         `json:"unit_amount,required"`
	JSON       tierConfigJSON `json:"-"`
}

// tierConfigJSON contains the JSON metadata for the struct [TierConfig]
type tierConfigJSON struct {
	FirstUnit   apijson.Field
	LastUnit    apijson.Field
	UnitAmount  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *TierConfig) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r tierConfigJSON) RawJSON() string {
	return r.raw
}

type TierSubLineItem struct {
	// The total amount for this sub line item.
	Amount     string              `json:"amount,required"`
	Grouping   SubLineItemGrouping `json:"grouping,required,nullable"`
	Name       string              `json:"name,required"`
	Quantity   float64             `json:"quantity,required"`
	TierConfig TierConfig          `json:"tier_config,required"`
	Type       TierSubLineItemType `json:"type,required"`
	JSON       tierSubLineItemJSON `json:"-"`
}

// tierSubLineItemJSON contains the JSON metadata for the struct [TierSubLineItem]
type tierSubLineItemJSON struct {
	Amount      apijson.Field
	Grouping    apijson.Field
	Name        apijson.Field
	Quantity    apijson.Field
	TierConfig  apijson.Field
	Type        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *TierSubLineItem) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r tierSubLineItemJSON) RawJSON() string {
	return r.raw
}

func (r TierSubLineItem) ImplementsInvoiceLineItemsSubLineItem() {}

func (r TierSubLineItem) ImplementsInvoiceLineItemNewResponseSubLineItem() {}

func (r TierSubLineItem) ImplementsInvoiceFetchUpcomingResponseLineItemsSubLineItem() {}

type TierSubLineItemType string

const (
	TierSubLineItemTypeTier TierSubLineItemType = "tier"
)

func (r TierSubLineItemType) IsKnown() bool {
	switch r {
	case TierSubLineItemTypeTier:
		return true
	}
	return false
}

type TieredBPSConfig struct {
	// Tiers for a Graduated BPS pricing model, where usage is bucketed into specified
	// tiers
	Tiers []BPSTier           `json:"tiers,required"`
	JSON  tieredBPSConfigJSON `json:"-"`
}

// tieredBPSConfigJSON contains the JSON metadata for the struct [TieredBPSConfig]
type tieredBPSConfigJSON struct {
	Tiers       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *TieredBPSConfig) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r tieredBPSConfigJSON) RawJSON() string {
	return r.raw
}

type TieredBPSConfigParam struct {
	// Tiers for a Graduated BPS pricing model, where usage is bucketed into specified
	// tiers
	Tiers param.Field[[]BPSTierParam] `json:"tiers,required"`
}

func (r TieredBPSConfigParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type TieredConfig struct {
	// Tiers for rating based on total usage quantities into the specified tier
	Tiers []Tier           `json:"tiers,required"`
	JSON  tieredConfigJSON `json:"-"`
}

// tieredConfigJSON contains the JSON metadata for the struct [TieredConfig]
type tieredConfigJSON struct {
	Tiers       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *TieredConfig) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r tieredConfigJSON) RawJSON() string {
	return r.raw
}

type TieredConfigParam struct {
	// Tiers for rating based on total usage quantities into the specified tier
	Tiers param.Field[[]TierParam] `json:"tiers,required"`
}

func (r TieredConfigParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type TieredConversionRateConfig struct {
	ConversionRateType TieredConversionRateConfigConversionRateType `json:"conversion_rate_type,required"`
	TieredConfig       ConversionRateTieredConfig                   `json:"tiered_config,required"`
	JSON               tieredConversionRateConfigJSON               `json:"-"`
}

// tieredConversionRateConfigJSON contains the JSON metadata for the struct
// [TieredConversionRateConfig]
type tieredConversionRateConfigJSON struct {
	ConversionRateType apijson.Field
	TieredConfig       apijson.Field
	raw                string
	ExtraFields        map[string]apijson.Field
}

func (r *TieredConversionRateConfig) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r tieredConversionRateConfigJSON) RawJSON() string {
	return r.raw
}

type TieredConversionRateConfigConversionRateType string

const (
	TieredConversionRateConfigConversionRateTypeTiered TieredConversionRateConfigConversionRateType = "tiered"
)

func (r TieredConversionRateConfigConversionRateType) IsKnown() bool {
	switch r {
	case TieredConversionRateConfigConversionRateTypeTiered:
		return true
	}
	return false
}

type TieredConversionRateConfigParam struct {
	ConversionRateType param.Field[TieredConversionRateConfigConversionRateType] `json:"conversion_rate_type,required"`
	TieredConfig       param.Field[ConversionRateTieredConfigParam]              `json:"tiered_config,required"`
}

func (r TieredConversionRateConfigParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r TieredConversionRateConfigParam) ImplementsPriceNewParamsNewFloatingUnitPriceConversionRateConfigUnion() {
}

type TransformPriceFilter struct {
	// The property of the price to filter on.
	Field TransformPriceFilterField `json:"field,required"`
	// Should prices that match the filter be included or excluded.
	Operator TransformPriceFilterOperator `json:"operator,required"`
	// The IDs or values that match this filter.
	Values []string                 `json:"values,required"`
	JSON   transformPriceFilterJSON `json:"-"`
}

// transformPriceFilterJSON contains the JSON metadata for the struct
// [TransformPriceFilter]
type transformPriceFilterJSON struct {
	Field       apijson.Field
	Operator    apijson.Field
	Values      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *TransformPriceFilter) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r transformPriceFilterJSON) RawJSON() string {
	return r.raw
}

// The property of the price to filter on.
type TransformPriceFilterField string

const (
	TransformPriceFilterFieldPriceID       TransformPriceFilterField = "price_id"
	TransformPriceFilterFieldItemID        TransformPriceFilterField = "item_id"
	TransformPriceFilterFieldPriceType     TransformPriceFilterField = "price_type"
	TransformPriceFilterFieldCurrency      TransformPriceFilterField = "currency"
	TransformPriceFilterFieldPricingUnitID TransformPriceFilterField = "pricing_unit_id"
)

func (r TransformPriceFilterField) IsKnown() bool {
	switch r {
	case TransformPriceFilterFieldPriceID, TransformPriceFilterFieldItemID, TransformPriceFilterFieldPriceType, TransformPriceFilterFieldCurrency, TransformPriceFilterFieldPricingUnitID:
		return true
	}
	return false
}

// Should prices that match the filter be included or excluded.
type TransformPriceFilterOperator string

const (
	TransformPriceFilterOperatorIncludes TransformPriceFilterOperator = "includes"
	TransformPriceFilterOperatorExcludes TransformPriceFilterOperator = "excludes"
)

func (r TransformPriceFilterOperator) IsKnown() bool {
	switch r {
	case TransformPriceFilterOperatorIncludes, TransformPriceFilterOperatorExcludes:
		return true
	}
	return false
}

type TransformPriceFilterParam struct {
	// The property of the price to filter on.
	Field param.Field[TransformPriceFilterField] `json:"field,required"`
	// Should prices that match the filter be included or excluded.
	Operator param.Field[TransformPriceFilterOperator] `json:"operator,required"`
	// The IDs or values that match this filter.
	Values param.Field[[]string] `json:"values,required"`
}

func (r TransformPriceFilterParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type TrialDiscount struct {
	DiscountType TrialDiscountDiscountType `json:"discount_type,required"`
	// List of price_ids that this discount applies to. For plan/plan phase discounts,
	// this can be a subset of prices.
	AppliesToPriceIDs []string `json:"applies_to_price_ids,nullable"`
	// The filters that determine which prices to apply this discount to.
	Filters []TransformPriceFilter `json:"filters,nullable"`
	Reason  string                 `json:"reason,nullable"`
	// Only available if discount_type is `trial`
	TrialAmountDiscount string `json:"trial_amount_discount,nullable"`
	// Only available if discount_type is `trial`
	TrialPercentageDiscount float64           `json:"trial_percentage_discount,nullable"`
	JSON                    trialDiscountJSON `json:"-"`
}

// trialDiscountJSON contains the JSON metadata for the struct [TrialDiscount]
type trialDiscountJSON struct {
	DiscountType            apijson.Field
	AppliesToPriceIDs       apijson.Field
	Filters                 apijson.Field
	Reason                  apijson.Field
	TrialAmountDiscount     apijson.Field
	TrialPercentageDiscount apijson.Field
	raw                     string
	ExtraFields             map[string]apijson.Field
}

func (r *TrialDiscount) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r trialDiscountJSON) RawJSON() string {
	return r.raw
}

func (r TrialDiscount) ImplementsDiscount() {}

func (r TrialDiscount) ImplementsInvoiceLevelDiscount() {}

type TrialDiscountDiscountType string

const (
	TrialDiscountDiscountTypeTrial TrialDiscountDiscountType = "trial"
)

func (r TrialDiscountDiscountType) IsKnown() bool {
	switch r {
	case TrialDiscountDiscountTypeTrial:
		return true
	}
	return false
}

type TrialDiscountParam struct {
	DiscountType param.Field[TrialDiscountDiscountType] `json:"discount_type,required"`
	// List of price_ids that this discount applies to. For plan/plan phase discounts,
	// this can be a subset of prices.
	AppliesToPriceIDs param.Field[[]string] `json:"applies_to_price_ids"`
	// The filters that determine which prices to apply this discount to.
	Filters param.Field[[]TransformPriceFilterParam] `json:"filters"`
	Reason  param.Field[string]                      `json:"reason"`
	// Only available if discount_type is `trial`
	TrialAmountDiscount param.Field[string] `json:"trial_amount_discount"`
	// Only available if discount_type is `trial`
	TrialPercentageDiscount param.Field[float64] `json:"trial_percentage_discount"`
}

func (r TrialDiscountParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r TrialDiscountParam) ImplementsDiscountUnionParam() {}

type UnitConfig struct {
	// Rate per unit of usage
	UnitAmount string         `json:"unit_amount,required"`
	JSON       unitConfigJSON `json:"-"`
}

// unitConfigJSON contains the JSON metadata for the struct [UnitConfig]
type unitConfigJSON struct {
	UnitAmount  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *UnitConfig) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r unitConfigJSON) RawJSON() string {
	return r.raw
}

type UnitConfigParam struct {
	// Rate per unit of usage
	UnitAmount param.Field[string] `json:"unit_amount,required"`
}

func (r UnitConfigParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type UnitConversionRateConfig struct {
	ConversionRateType UnitConversionRateConfigConversionRateType `json:"conversion_rate_type,required"`
	UnitConfig         ConversionRateUnitConfig                   `json:"unit_config,required"`
	JSON               unitConversionRateConfigJSON               `json:"-"`
}

// unitConversionRateConfigJSON contains the JSON metadata for the struct
// [UnitConversionRateConfig]
type unitConversionRateConfigJSON struct {
	ConversionRateType apijson.Field
	UnitConfig         apijson.Field
	raw                string
	ExtraFields        map[string]apijson.Field
}

func (r *UnitConversionRateConfig) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r unitConversionRateConfigJSON) RawJSON() string {
	return r.raw
}

type UnitConversionRateConfigConversionRateType string

const (
	UnitConversionRateConfigConversionRateTypeUnit UnitConversionRateConfigConversionRateType = "unit"
)

func (r UnitConversionRateConfigConversionRateType) IsKnown() bool {
	switch r {
	case UnitConversionRateConfigConversionRateTypeUnit:
		return true
	}
	return false
}

type UnitConversionRateConfigParam struct {
	ConversionRateType param.Field[UnitConversionRateConfigConversionRateType] `json:"conversion_rate_type,required"`
	UnitConfig         param.Field[ConversionRateUnitConfigParam]              `json:"unit_config,required"`
}

func (r UnitConversionRateConfigParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r UnitConversionRateConfigParam) ImplementsPriceNewParamsNewFloatingUnitPriceConversionRateConfigUnion() {
}

type UsageDiscount struct {
	DiscountType UsageDiscountDiscountType `json:"discount_type,required"`
	// Only available if discount_type is `usage`. Number of usage units that this
	// discount is for
	UsageDiscount float64 `json:"usage_discount,required"`
	// List of price_ids that this discount applies to. For plan/plan phase discounts,
	// this can be a subset of prices.
	AppliesToPriceIDs []string `json:"applies_to_price_ids,nullable"`
	// The filters that determine which prices to apply this discount to.
	Filters []TransformPriceFilter `json:"filters,nullable"`
	Reason  string                 `json:"reason,nullable"`
	JSON    usageDiscountJSON      `json:"-"`
}

// usageDiscountJSON contains the JSON metadata for the struct [UsageDiscount]
type usageDiscountJSON struct {
	DiscountType      apijson.Field
	UsageDiscount     apijson.Field
	AppliesToPriceIDs apijson.Field
	Filters           apijson.Field
	Reason            apijson.Field
	raw               string
	ExtraFields       map[string]apijson.Field
}

func (r *UsageDiscount) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r usageDiscountJSON) RawJSON() string {
	return r.raw
}

func (r UsageDiscount) ImplementsDiscount() {}

type UsageDiscountDiscountType string

const (
	UsageDiscountDiscountTypeUsage UsageDiscountDiscountType = "usage"
)

func (r UsageDiscountDiscountType) IsKnown() bool {
	switch r {
	case UsageDiscountDiscountTypeUsage:
		return true
	}
	return false
}

type UsageDiscountParam struct {
	DiscountType param.Field[UsageDiscountDiscountType] `json:"discount_type,required"`
	// Only available if discount_type is `usage`. Number of usage units that this
	// discount is for
	UsageDiscount param.Field[float64] `json:"usage_discount,required"`
	// List of price_ids that this discount applies to. For plan/plan phase discounts,
	// this can be a subset of prices.
	AppliesToPriceIDs param.Field[[]string] `json:"applies_to_price_ids"`
	// The filters that determine which prices to apply this discount to.
	Filters param.Field[[]TransformPriceFilterParam] `json:"filters"`
	Reason  param.Field[string]                      `json:"reason"`
}

func (r UsageDiscountParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r UsageDiscountParam) ImplementsDiscountUnionParam() {}

type UsageDiscountInterval struct {
	// The price interval ids that this discount interval applies to.
	AppliesToPriceIntervalIDs []string                          `json:"applies_to_price_interval_ids,required"`
	DiscountType              UsageDiscountIntervalDiscountType `json:"discount_type,required"`
	// The end date of the discount interval.
	EndDate time.Time `json:"end_date,required,nullable" format:"date-time"`
	// The filters that determine which prices this discount interval applies to.
	Filters []TransformPriceFilter `json:"filters,required"`
	// The start date of the discount interval.
	StartDate time.Time `json:"start_date,required" format:"date-time"`
	// Only available if discount_type is `usage`. Number of usage units that this
	// discount is for
	UsageDiscount float64                   `json:"usage_discount,required"`
	JSON          usageDiscountIntervalJSON `json:"-"`
}

// usageDiscountIntervalJSON contains the JSON metadata for the struct
// [UsageDiscountInterval]
type usageDiscountIntervalJSON struct {
	AppliesToPriceIntervalIDs apijson.Field
	DiscountType              apijson.Field
	EndDate                   apijson.Field
	Filters                   apijson.Field
	StartDate                 apijson.Field
	UsageDiscount             apijson.Field
	raw                       string
	ExtraFields               map[string]apijson.Field
}

func (r *UsageDiscountInterval) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r usageDiscountIntervalJSON) RawJSON() string {
	return r.raw
}

func (r UsageDiscountInterval) ImplementsSubscriptionDiscountInterval() {}

func (r UsageDiscountInterval) ImplementsMutatedSubscriptionDiscountInterval() {}

type UsageDiscountIntervalDiscountType string

const (
	UsageDiscountIntervalDiscountTypeUsage UsageDiscountIntervalDiscountType = "usage"
)

func (r UsageDiscountIntervalDiscountType) IsKnown() bool {
	switch r {
	case UsageDiscountIntervalDiscountTypeUsage:
		return true
	}
	return false
}
