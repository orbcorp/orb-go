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

// PlanService contains methods and other services that help with interacting with
// the orb API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewPlanService] method instead.
type PlanService struct {
	Options        []option.RequestOption
	ExternalPlanID *PlanExternalPlanIDService
}

// NewPlanService generates a new service that applies the given options to each
// request. These options are applied after the parent client's options (if there
// is one), and before any request-specific options.
func NewPlanService(opts ...option.RequestOption) (r *PlanService) {
	r = &PlanService{}
	r.Options = opts
	r.ExternalPlanID = NewPlanExternalPlanIDService(opts...)
	return
}

// This endpoint allows creation of plans including their prices.
func (r *PlanService) New(ctx context.Context, body PlanNewParams, opts ...option.RequestOption) (res *Plan, err error) {
	opts = slices.Concat(r.Options, opts)
	path := "plans"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return
}

// This endpoint can be used to update the `external_plan_id`, and `metadata` of an
// existing plan.
//
// Other fields on a plan are currently immutable.
func (r *PlanService) Update(ctx context.Context, planID string, body PlanUpdateParams, opts ...option.RequestOption) (res *Plan, err error) {
	opts = slices.Concat(r.Options, opts)
	if planID == "" {
		err = errors.New("missing required plan_id parameter")
		return
	}
	path := fmt.Sprintf("plans/%s", planID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPut, path, body, &res, opts...)
	return
}

// This endpoint returns a list of all [plans](/core-concepts#plan-and-price) for
// an account in a list format. The list of plans is ordered starting from the most
// recently created plan. The response also includes
// [`pagination_metadata`](/api-reference/pagination), which lets the caller
// retrieve the next page of results if they exist.
func (r *PlanService) List(ctx context.Context, query PlanListParams, opts ...option.RequestOption) (res *pagination.Page[Plan], err error) {
	var raw *http.Response
	opts = slices.Concat(r.Options, opts)
	opts = append([]option.RequestOption{option.WithResponseInto(&raw)}, opts...)
	path := "plans"
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

// This endpoint returns a list of all [plans](/core-concepts#plan-and-price) for
// an account in a list format. The list of plans is ordered starting from the most
// recently created plan. The response also includes
// [`pagination_metadata`](/api-reference/pagination), which lets the caller
// retrieve the next page of results if they exist.
func (r *PlanService) ListAutoPaging(ctx context.Context, query PlanListParams, opts ...option.RequestOption) *pagination.PageAutoPager[Plan] {
	return pagination.NewPageAutoPager(r.List(ctx, query, opts...))
}

// This endpoint is used to fetch [plan](/core-concepts#plan-and-price) details
// given a plan identifier. It returns information about the prices included in the
// plan and their configuration, as well as the product that the plan is attached
// to.
//
// ## Serialized prices
//
// Orb supports a few different pricing models out of the box. Each of these models
// is serialized differently in a given [Price](/core-concepts#plan-and-price)
// object. The `model_type` field determines the key for the configuration object
// that is present. A detailed explanation of price types can be found in the
// [Price schema](/core-concepts#plan-and-price).
//
// ## Phases
//
// Orb supports plan phases, also known as contract ramps. For plans with phases,
// the serialized prices refer to all prices across all phases.
func (r *PlanService) Fetch(ctx context.Context, planID string, opts ...option.RequestOption) (res *Plan, err error) {
	opts = slices.Concat(r.Options, opts)
	if planID == "" {
		err = errors.New("missing required plan_id parameter")
		return
	}
	path := fmt.Sprintf("plans/%s", planID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

// The [Plan](/core-concepts#plan-and-price) resource represents a plan that can be
// subscribed to by a customer. Plans define the billing behavior of the
// subscription. You can see more about how to configure prices in the
// [Price resource](/reference/price).
type Plan struct {
	ID string `json:"id,required"`
	// Adjustments for this plan. If the plan has phases, this includes adjustments
	// across all phases of the plan.
	Adjustments []PlanAdjustment `json:"adjustments,required"`
	CreatedAt   time.Time        `json:"created_at,required" format:"date-time"`
	// An ISO 4217 currency string or custom pricing unit (`credits`) for this plan's
	// prices.
	//
	// Deprecated: deprecated
	Currency string `json:"currency,required"`
	// The default memo text on the invoices corresponding to subscriptions on this
	// plan. Note that each subscription may configure its own memo.
	DefaultInvoiceMemo string `json:"default_invoice_memo,required,nullable"`
	Description        string `json:"description,required"`
	// Deprecated: deprecated
	Discount shared.Discount `json:"discount,required,nullable"`
	// An optional user-defined ID for this plan resource, used throughout the system
	// as an alias for this Plan. Use this field to identify a plan by an existing
	// identifier in your system.
	ExternalPlanID string `json:"external_plan_id,required,nullable"`
	// An ISO 4217 currency string for which this plan is billed in. Matches `currency`
	// unless `currency` is a custom pricing unit.
	InvoicingCurrency string `json:"invoicing_currency,required"`
	// Deprecated: deprecated
	Maximum shared.Maximum `json:"maximum,required,nullable"`
	// Deprecated: deprecated
	MaximumAmount string `json:"maximum_amount,required,nullable"`
	// User specified key-value pairs for the resource. If not present, this defaults
	// to an empty dictionary. Individual keys can be removed by setting the value to
	// `null`, and the entire metadata mapping can be cleared by setting `metadata` to
	// `null`.
	Metadata map[string]string `json:"metadata,required"`
	// Deprecated: deprecated
	Minimum shared.Minimum `json:"minimum,required,nullable"`
	// Deprecated: deprecated
	MinimumAmount string `json:"minimum_amount,required,nullable"`
	Name          string `json:"name,required"`
	// Determines the difference between the invoice issue date and the due date. A
	// value of "0" here signifies that invoices are due on issue, whereas a value of
	// "30" means that the customer has a month to pay the invoice before its overdue.
	// Note that individual subscriptions or invoices may set a different net terms
	// configuration.
	NetTerms   int64           `json:"net_terms,required,nullable"`
	PlanPhases []PlanPlanPhase `json:"plan_phases,required,nullable"`
	// Prices for this plan. If the plan has phases, this includes prices across all
	// phases of the plan.
	Prices      []shared.Price  `json:"prices,required"`
	Product     PlanProduct     `json:"product,required"`
	Status      PlanStatus      `json:"status,required"`
	TrialConfig PlanTrialConfig `json:"trial_config,required"`
	Version     int64           `json:"version,required"`
	BasePlan    PlanBasePlan    `json:"base_plan,nullable"`
	// The parent plan id if the given plan was created by overriding one or more of
	// the parent's prices
	BasePlanID string   `json:"base_plan_id,nullable"`
	JSON       planJSON `json:"-"`
}

// planJSON contains the JSON metadata for the struct [Plan]
type planJSON struct {
	ID                 apijson.Field
	Adjustments        apijson.Field
	CreatedAt          apijson.Field
	Currency           apijson.Field
	DefaultInvoiceMemo apijson.Field
	Description        apijson.Field
	Discount           apijson.Field
	ExternalPlanID     apijson.Field
	InvoicingCurrency  apijson.Field
	Maximum            apijson.Field
	MaximumAmount      apijson.Field
	Metadata           apijson.Field
	Minimum            apijson.Field
	MinimumAmount      apijson.Field
	Name               apijson.Field
	NetTerms           apijson.Field
	PlanPhases         apijson.Field
	Prices             apijson.Field
	Product            apijson.Field
	Status             apijson.Field
	TrialConfig        apijson.Field
	Version            apijson.Field
	BasePlan           apijson.Field
	BasePlanID         apijson.Field
	raw                string
	ExtraFields        map[string]apijson.Field
}

func (r *Plan) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r planJSON) RawJSON() string {
	return r.raw
}

type PlanAdjustment struct {
	ID             string                        `json:"id,required"`
	AdjustmentType PlanAdjustmentsAdjustmentType `json:"adjustment_type,required"`
	// This field can have the runtime type of [[]string].
	AppliesToPriceIDs interface{} `json:"applies_to_price_ids,required"`
	// This field can have the runtime type of
	// [[]shared.PlanPhaseUsageDiscountAdjustmentFilter],
	// [[]shared.PlanPhaseAmountDiscountAdjustmentFilter],
	// [[]shared.PlanPhasePercentageDiscountAdjustmentFilter],
	// [[]shared.PlanPhaseMinimumAdjustmentFilter],
	// [[]shared.PlanPhaseMaximumAdjustmentFilter].
	Filters interface{} `json:"filters,required"`
	// True for adjustments that apply to an entire invoice, false for adjustments that
	// apply to only one price.
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
	UsageDiscount float64            `json:"usage_discount"`
	JSON          planAdjustmentJSON `json:"-"`
	union         PlanAdjustmentsUnion
}

// planAdjustmentJSON contains the JSON metadata for the struct [PlanAdjustment]
type planAdjustmentJSON struct {
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

func (r planAdjustmentJSON) RawJSON() string {
	return r.raw
}

func (r *PlanAdjustment) UnmarshalJSON(data []byte) (err error) {
	*r = PlanAdjustment{}
	err = apijson.UnmarshalRoot(data, &r.union)
	if err != nil {
		return err
	}
	return apijson.Port(r.union, &r)
}

// AsUnion returns a [PlanAdjustmentsUnion] interface which you can cast to the
// specific types for more type safety.
//
// Possible runtime types of the union are
// [shared.PlanPhaseUsageDiscountAdjustment],
// [shared.PlanPhaseAmountDiscountAdjustment],
// [shared.PlanPhasePercentageDiscountAdjustment],
// [shared.PlanPhaseMinimumAdjustment], [shared.PlanPhaseMaximumAdjustment].
func (r PlanAdjustment) AsUnion() PlanAdjustmentsUnion {
	return r.union
}

// Union satisfied by [shared.PlanPhaseUsageDiscountAdjustment],
// [shared.PlanPhaseAmountDiscountAdjustment],
// [shared.PlanPhasePercentageDiscountAdjustment],
// [shared.PlanPhaseMinimumAdjustment] or [shared.PlanPhaseMaximumAdjustment].
type PlanAdjustmentsUnion interface {
	ImplementsPlanAdjustment()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*PlanAdjustmentsUnion)(nil)).Elem(),
		"adjustment_type",
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(shared.PlanPhaseUsageDiscountAdjustment{}),
			DiscriminatorValue: "usage_discount",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(shared.PlanPhaseAmountDiscountAdjustment{}),
			DiscriminatorValue: "amount_discount",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(shared.PlanPhasePercentageDiscountAdjustment{}),
			DiscriminatorValue: "percentage_discount",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(shared.PlanPhaseMinimumAdjustment{}),
			DiscriminatorValue: "minimum",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(shared.PlanPhaseMaximumAdjustment{}),
			DiscriminatorValue: "maximum",
		},
	)
}

type PlanAdjustmentsAdjustmentType string

const (
	PlanAdjustmentsAdjustmentTypeUsageDiscount      PlanAdjustmentsAdjustmentType = "usage_discount"
	PlanAdjustmentsAdjustmentTypeAmountDiscount     PlanAdjustmentsAdjustmentType = "amount_discount"
	PlanAdjustmentsAdjustmentTypePercentageDiscount PlanAdjustmentsAdjustmentType = "percentage_discount"
	PlanAdjustmentsAdjustmentTypeMinimum            PlanAdjustmentsAdjustmentType = "minimum"
	PlanAdjustmentsAdjustmentTypeMaximum            PlanAdjustmentsAdjustmentType = "maximum"
)

func (r PlanAdjustmentsAdjustmentType) IsKnown() bool {
	switch r {
	case PlanAdjustmentsAdjustmentTypeUsageDiscount, PlanAdjustmentsAdjustmentTypeAmountDiscount, PlanAdjustmentsAdjustmentTypePercentageDiscount, PlanAdjustmentsAdjustmentTypeMinimum, PlanAdjustmentsAdjustmentTypeMaximum:
		return true
	}
	return false
}

type PlanPlanPhase struct {
	ID          string          `json:"id,required"`
	Description string          `json:"description,required,nullable"`
	Discount    shared.Discount `json:"discount,required,nullable"`
	// How many terms of length `duration_unit` this phase is active for. If null, this
	// phase is evergreen and active indefinitely
	Duration      int64                      `json:"duration,required,nullable"`
	DurationUnit  PlanPlanPhasesDurationUnit `json:"duration_unit,required,nullable"`
	Maximum       shared.Maximum             `json:"maximum,required,nullable"`
	MaximumAmount string                     `json:"maximum_amount,required,nullable"`
	Minimum       shared.Minimum             `json:"minimum,required,nullable"`
	MinimumAmount string                     `json:"minimum_amount,required,nullable"`
	Name          string                     `json:"name,required"`
	// Determines the ordering of the phase in a plan's lifecycle. 1 = first phase.
	Order int64             `json:"order,required"`
	JSON  planPlanPhaseJSON `json:"-"`
}

// planPlanPhaseJSON contains the JSON metadata for the struct [PlanPlanPhase]
type planPlanPhaseJSON struct {
	ID            apijson.Field
	Description   apijson.Field
	Discount      apijson.Field
	Duration      apijson.Field
	DurationUnit  apijson.Field
	Maximum       apijson.Field
	MaximumAmount apijson.Field
	Minimum       apijson.Field
	MinimumAmount apijson.Field
	Name          apijson.Field
	Order         apijson.Field
	raw           string
	ExtraFields   map[string]apijson.Field
}

func (r *PlanPlanPhase) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r planPlanPhaseJSON) RawJSON() string {
	return r.raw
}

type PlanPlanPhasesDurationUnit string

const (
	PlanPlanPhasesDurationUnitDaily      PlanPlanPhasesDurationUnit = "daily"
	PlanPlanPhasesDurationUnitMonthly    PlanPlanPhasesDurationUnit = "monthly"
	PlanPlanPhasesDurationUnitQuarterly  PlanPlanPhasesDurationUnit = "quarterly"
	PlanPlanPhasesDurationUnitSemiAnnual PlanPlanPhasesDurationUnit = "semi_annual"
	PlanPlanPhasesDurationUnitAnnual     PlanPlanPhasesDurationUnit = "annual"
)

func (r PlanPlanPhasesDurationUnit) IsKnown() bool {
	switch r {
	case PlanPlanPhasesDurationUnitDaily, PlanPlanPhasesDurationUnitMonthly, PlanPlanPhasesDurationUnitQuarterly, PlanPlanPhasesDurationUnitSemiAnnual, PlanPlanPhasesDurationUnitAnnual:
		return true
	}
	return false
}

type PlanProduct struct {
	ID        string          `json:"id,required"`
	CreatedAt time.Time       `json:"created_at,required" format:"date-time"`
	Name      string          `json:"name,required"`
	JSON      planProductJSON `json:"-"`
}

// planProductJSON contains the JSON metadata for the struct [PlanProduct]
type planProductJSON struct {
	ID          apijson.Field
	CreatedAt   apijson.Field
	Name        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *PlanProduct) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r planProductJSON) RawJSON() string {
	return r.raw
}

type PlanStatus string

const (
	PlanStatusActive   PlanStatus = "active"
	PlanStatusArchived PlanStatus = "archived"
	PlanStatusDraft    PlanStatus = "draft"
)

func (r PlanStatus) IsKnown() bool {
	switch r {
	case PlanStatusActive, PlanStatusArchived, PlanStatusDraft:
		return true
	}
	return false
}

type PlanTrialConfig struct {
	TrialPeriod     int64                          `json:"trial_period,required,nullable"`
	TrialPeriodUnit PlanTrialConfigTrialPeriodUnit `json:"trial_period_unit,required"`
	JSON            planTrialConfigJSON            `json:"-"`
}

// planTrialConfigJSON contains the JSON metadata for the struct [PlanTrialConfig]
type planTrialConfigJSON struct {
	TrialPeriod     apijson.Field
	TrialPeriodUnit apijson.Field
	raw             string
	ExtraFields     map[string]apijson.Field
}

func (r *PlanTrialConfig) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r planTrialConfigJSON) RawJSON() string {
	return r.raw
}

type PlanTrialConfigTrialPeriodUnit string

const (
	PlanTrialConfigTrialPeriodUnitDays PlanTrialConfigTrialPeriodUnit = "days"
)

func (r PlanTrialConfigTrialPeriodUnit) IsKnown() bool {
	switch r {
	case PlanTrialConfigTrialPeriodUnitDays:
		return true
	}
	return false
}

type PlanBasePlan struct {
	ID string `json:"id,required,nullable"`
	// An optional user-defined ID for this plan resource, used throughout the system
	// as an alias for this Plan. Use this field to identify a plan by an existing
	// identifier in your system.
	ExternalPlanID string           `json:"external_plan_id,required,nullable"`
	Name           string           `json:"name,required,nullable"`
	JSON           planBasePlanJSON `json:"-"`
}

// planBasePlanJSON contains the JSON metadata for the struct [PlanBasePlan]
type planBasePlanJSON struct {
	ID             apijson.Field
	ExternalPlanID apijson.Field
	Name           apijson.Field
	raw            string
	ExtraFields    map[string]apijson.Field
}

func (r *PlanBasePlan) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r planBasePlanJSON) RawJSON() string {
	return r.raw
}

type PlanNewParams struct {
	// An ISO 4217 currency string for invoices generated by subscriptions on this
	// plan.
	Currency param.Field[string] `json:"currency,required"`
	Name     param.Field[string] `json:"name,required"`
	// Prices for this plan. If the plan has phases, this includes prices across all
	// phases of the plan.
	Prices param.Field[[]PlanNewParamsPrice] `json:"prices,required"`
	// Adjustments for this plan. If the plan has phases, this includes adjustments
	// across all phases of the plan.
	Adjustments param.Field[[]PlanNewParamsAdjustment] `json:"adjustments"`
	// Free-form text which is available on the invoice PDF and the Orb invoice portal.
	DefaultInvoiceMemo param.Field[string] `json:"default_invoice_memo"`
	ExternalPlanID     param.Field[string] `json:"external_plan_id"`
	// User-specified key/value pairs for the resource. Individual keys can be removed
	// by setting the value to `null`, and the entire metadata mapping can be cleared
	// by setting `metadata` to `null`.
	Metadata param.Field[map[string]string] `json:"metadata"`
	// The net terms determines the difference between the invoice date and the issue
	// date for the invoice. If you intend the invoice to be due on issue, set this
	// to 0.
	NetTerms param.Field[int64] `json:"net_terms"`
	// Configuration of pre-defined phases, each with their own prices and adjustments.
	// Leave unspecified for plans with a single phase.
	PlanPhases param.Field[[]PlanNewParamsPlanPhase] `json:"plan_phases"`
	// The status of the plan to create (either active or draft). If not specified,
	// this defaults to active.
	Status param.Field[PlanNewParamsStatus] `json:"status"`
}

func (r PlanNewParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type PlanNewParamsPrice struct {
	// The allocation price to add to the plan.
	AllocationPrice param.Field[shared.NewAllocationPriceParam] `json:"allocation_price"`
	// The phase to add this price to.
	PlanPhaseOrder param.Field[int64] `json:"plan_phase_order"`
	// New plan price request body params.
	Price param.Field[PlanNewParamsPricesPriceUnion] `json:"price"`
}

func (r PlanNewParamsPrice) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// New plan price request body params.
type PlanNewParamsPricesPrice struct {
	// The cadence to bill for this price on.
	Cadence param.Field[PlanNewParamsPricesPriceCadence] `json:"cadence,required"`
	// The id of the item the price will be associated with.
	ItemID param.Field[string] `json:"item_id,required"`
	// The pricing model type
	ModelType param.Field[PlanNewParamsPricesPriceModelType] `json:"model_type,required"`
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
	ConversionRate                    param.Field[float64]     `json:"conversion_rate"`
	ConversionRateConfig              param.Field[interface{}] `json:"conversion_rate_config"`
	CumulativeGroupedAllocationConfig param.Field[interface{}] `json:"cumulative_grouped_allocation_config"`
	CumulativeGroupedBulkConfig       param.Field[interface{}] `json:"cumulative_grouped_bulk_config"`
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

func (r PlanNewParamsPricesPrice) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r PlanNewParamsPricesPrice) ImplementsPlanNewParamsPricesPriceUnion() {}

// New plan price request body params.
//
// Satisfied by [shared.NewPlanUnitPriceParam], [shared.NewPlanTieredPriceParam],
// [shared.NewPlanBulkPriceParam],
// [PlanNewParamsPricesPriceNewPlanBulkWithFiltersPrice],
// [shared.NewPlanPackagePriceParam], [shared.NewPlanMatrixPriceParam],
// [shared.NewPlanThresholdTotalAmountPriceParam],
// [shared.NewPlanTieredPackagePriceParam],
// [shared.NewPlanTieredWithMinimumPriceParam],
// [shared.NewPlanGroupedTieredPriceParam],
// [shared.NewPlanTieredPackageWithMinimumPriceParam],
// [shared.NewPlanPackageWithAllocationPriceParam],
// [shared.NewPlanUnitWithPercentPriceParam],
// [shared.NewPlanMatrixWithAllocationPriceParam],
// [PlanNewParamsPricesPriceNewPlanTieredWithProrationPrice],
// [shared.NewPlanUnitWithProrationPriceParam],
// [shared.NewPlanGroupedAllocationPriceParam],
// [shared.NewPlanBulkWithProrationPriceParam],
// [shared.NewPlanGroupedWithProratedMinimumPriceParam],
// [shared.NewPlanGroupedWithMeteredMinimumPriceParam],
// [PlanNewParamsPricesPriceNewPlanGroupedWithMinMaxThresholdsPrice],
// [shared.NewPlanMatrixWithDisplayNamePriceParam],
// [shared.NewPlanGroupedTieredPackagePriceParam],
// [shared.NewPlanMaxGroupTieredPackagePriceParam],
// [shared.NewPlanScalableMatrixWithUnitPricingPriceParam],
// [shared.NewPlanScalableMatrixWithTieredPricingPriceParam],
// [shared.NewPlanCumulativeGroupedBulkPriceParam],
// [PlanNewParamsPricesPriceNewPlanCumulativeGroupedAllocationPrice],
// [shared.NewPlanMinimumCompositePriceParam],
// [PlanNewParamsPricesPriceNewPlanPercentCompositePrice],
// [PlanNewParamsPricesPriceNewPlanEventOutputPrice], [PlanNewParamsPricesPrice].
type PlanNewParamsPricesPriceUnion interface {
	ImplementsPlanNewParamsPricesPriceUnion()
}

type PlanNewParamsPricesPriceNewPlanBulkWithFiltersPrice struct {
	// Configuration for bulk_with_filters pricing
	BulkWithFiltersConfig param.Field[PlanNewParamsPricesPriceNewPlanBulkWithFiltersPriceBulkWithFiltersConfig] `json:"bulk_with_filters_config,required"`
	// The cadence to bill for this price on.
	Cadence param.Field[PlanNewParamsPricesPriceNewPlanBulkWithFiltersPriceCadence] `json:"cadence,required"`
	// The id of the item the price will be associated with.
	ItemID param.Field[string] `json:"item_id,required"`
	// The pricing model type
	ModelType param.Field[PlanNewParamsPricesPriceNewPlanBulkWithFiltersPriceModelType] `json:"model_type,required"`
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
	ConversionRateConfig param.Field[PlanNewParamsPricesPriceNewPlanBulkWithFiltersPriceConversionRateConfigUnion] `json:"conversion_rate_config"`
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

func (r PlanNewParamsPricesPriceNewPlanBulkWithFiltersPrice) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r PlanNewParamsPricesPriceNewPlanBulkWithFiltersPrice) ImplementsPlanNewParamsPricesPriceUnion() {
}

// Configuration for bulk_with_filters pricing
type PlanNewParamsPricesPriceNewPlanBulkWithFiltersPriceBulkWithFiltersConfig struct {
	// Property filters to apply (all must match)
	Filters param.Field[[]PlanNewParamsPricesPriceNewPlanBulkWithFiltersPriceBulkWithFiltersConfigFilter] `json:"filters,required"`
	// Bulk tiers for rating based on total usage volume
	Tiers param.Field[[]PlanNewParamsPricesPriceNewPlanBulkWithFiltersPriceBulkWithFiltersConfigTier] `json:"tiers,required"`
}

func (r PlanNewParamsPricesPriceNewPlanBulkWithFiltersPriceBulkWithFiltersConfig) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Configuration for a single property filter
type PlanNewParamsPricesPriceNewPlanBulkWithFiltersPriceBulkWithFiltersConfigFilter struct {
	// Event property key to filter on
	PropertyKey param.Field[string] `json:"property_key,required"`
	// Event property value to match
	PropertyValue param.Field[string] `json:"property_value,required"`
}

func (r PlanNewParamsPricesPriceNewPlanBulkWithFiltersPriceBulkWithFiltersConfigFilter) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Configuration for a single bulk pricing tier
type PlanNewParamsPricesPriceNewPlanBulkWithFiltersPriceBulkWithFiltersConfigTier struct {
	// Amount per unit
	UnitAmount param.Field[string] `json:"unit_amount,required"`
	// The lower bound for this tier
	TierLowerBound param.Field[string] `json:"tier_lower_bound"`
}

func (r PlanNewParamsPricesPriceNewPlanBulkWithFiltersPriceBulkWithFiltersConfigTier) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// The cadence to bill for this price on.
type PlanNewParamsPricesPriceNewPlanBulkWithFiltersPriceCadence string

const (
	PlanNewParamsPricesPriceNewPlanBulkWithFiltersPriceCadenceAnnual     PlanNewParamsPricesPriceNewPlanBulkWithFiltersPriceCadence = "annual"
	PlanNewParamsPricesPriceNewPlanBulkWithFiltersPriceCadenceSemiAnnual PlanNewParamsPricesPriceNewPlanBulkWithFiltersPriceCadence = "semi_annual"
	PlanNewParamsPricesPriceNewPlanBulkWithFiltersPriceCadenceMonthly    PlanNewParamsPricesPriceNewPlanBulkWithFiltersPriceCadence = "monthly"
	PlanNewParamsPricesPriceNewPlanBulkWithFiltersPriceCadenceQuarterly  PlanNewParamsPricesPriceNewPlanBulkWithFiltersPriceCadence = "quarterly"
	PlanNewParamsPricesPriceNewPlanBulkWithFiltersPriceCadenceOneTime    PlanNewParamsPricesPriceNewPlanBulkWithFiltersPriceCadence = "one_time"
	PlanNewParamsPricesPriceNewPlanBulkWithFiltersPriceCadenceCustom     PlanNewParamsPricesPriceNewPlanBulkWithFiltersPriceCadence = "custom"
)

func (r PlanNewParamsPricesPriceNewPlanBulkWithFiltersPriceCadence) IsKnown() bool {
	switch r {
	case PlanNewParamsPricesPriceNewPlanBulkWithFiltersPriceCadenceAnnual, PlanNewParamsPricesPriceNewPlanBulkWithFiltersPriceCadenceSemiAnnual, PlanNewParamsPricesPriceNewPlanBulkWithFiltersPriceCadenceMonthly, PlanNewParamsPricesPriceNewPlanBulkWithFiltersPriceCadenceQuarterly, PlanNewParamsPricesPriceNewPlanBulkWithFiltersPriceCadenceOneTime, PlanNewParamsPricesPriceNewPlanBulkWithFiltersPriceCadenceCustom:
		return true
	}
	return false
}

// The pricing model type
type PlanNewParamsPricesPriceNewPlanBulkWithFiltersPriceModelType string

const (
	PlanNewParamsPricesPriceNewPlanBulkWithFiltersPriceModelTypeBulkWithFilters PlanNewParamsPricesPriceNewPlanBulkWithFiltersPriceModelType = "bulk_with_filters"
)

func (r PlanNewParamsPricesPriceNewPlanBulkWithFiltersPriceModelType) IsKnown() bool {
	switch r {
	case PlanNewParamsPricesPriceNewPlanBulkWithFiltersPriceModelTypeBulkWithFilters:
		return true
	}
	return false
}

type PlanNewParamsPricesPriceNewPlanBulkWithFiltersPriceConversionRateConfig struct {
	ConversionRateType param.Field[PlanNewParamsPricesPriceNewPlanBulkWithFiltersPriceConversionRateConfigConversionRateType] `json:"conversion_rate_type,required"`
	TieredConfig       param.Field[shared.ConversionRateTieredConfigParam]                                                    `json:"tiered_config"`
	UnitConfig         param.Field[shared.ConversionRateUnitConfigParam]                                                      `json:"unit_config"`
}

func (r PlanNewParamsPricesPriceNewPlanBulkWithFiltersPriceConversionRateConfig) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r PlanNewParamsPricesPriceNewPlanBulkWithFiltersPriceConversionRateConfig) ImplementsPlanNewParamsPricesPriceNewPlanBulkWithFiltersPriceConversionRateConfigUnion() {
}

// Satisfied by [shared.UnitConversionRateConfigParam],
// [shared.TieredConversionRateConfigParam],
// [PlanNewParamsPricesPriceNewPlanBulkWithFiltersPriceConversionRateConfig].
type PlanNewParamsPricesPriceNewPlanBulkWithFiltersPriceConversionRateConfigUnion interface {
	ImplementsPlanNewParamsPricesPriceNewPlanBulkWithFiltersPriceConversionRateConfigUnion()
}

type PlanNewParamsPricesPriceNewPlanBulkWithFiltersPriceConversionRateConfigConversionRateType string

const (
	PlanNewParamsPricesPriceNewPlanBulkWithFiltersPriceConversionRateConfigConversionRateTypeUnit   PlanNewParamsPricesPriceNewPlanBulkWithFiltersPriceConversionRateConfigConversionRateType = "unit"
	PlanNewParamsPricesPriceNewPlanBulkWithFiltersPriceConversionRateConfigConversionRateTypeTiered PlanNewParamsPricesPriceNewPlanBulkWithFiltersPriceConversionRateConfigConversionRateType = "tiered"
)

func (r PlanNewParamsPricesPriceNewPlanBulkWithFiltersPriceConversionRateConfigConversionRateType) IsKnown() bool {
	switch r {
	case PlanNewParamsPricesPriceNewPlanBulkWithFiltersPriceConversionRateConfigConversionRateTypeUnit, PlanNewParamsPricesPriceNewPlanBulkWithFiltersPriceConversionRateConfigConversionRateTypeTiered:
		return true
	}
	return false
}

type PlanNewParamsPricesPriceNewPlanTieredWithProrationPrice struct {
	// The cadence to bill for this price on.
	Cadence param.Field[PlanNewParamsPricesPriceNewPlanTieredWithProrationPriceCadence] `json:"cadence,required"`
	// The id of the item the price will be associated with.
	ItemID param.Field[string] `json:"item_id,required"`
	// The pricing model type
	ModelType param.Field[PlanNewParamsPricesPriceNewPlanTieredWithProrationPriceModelType] `json:"model_type,required"`
	// The name of the price.
	Name param.Field[string] `json:"name,required"`
	// Configuration for tiered_with_proration pricing
	TieredWithProrationConfig param.Field[PlanNewParamsPricesPriceNewPlanTieredWithProrationPriceTieredWithProrationConfig] `json:"tiered_with_proration_config,required"`
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
	ConversionRateConfig param.Field[PlanNewParamsPricesPriceNewPlanTieredWithProrationPriceConversionRateConfigUnion] `json:"conversion_rate_config"`
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

func (r PlanNewParamsPricesPriceNewPlanTieredWithProrationPrice) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r PlanNewParamsPricesPriceNewPlanTieredWithProrationPrice) ImplementsPlanNewParamsPricesPriceUnion() {
}

// The cadence to bill for this price on.
type PlanNewParamsPricesPriceNewPlanTieredWithProrationPriceCadence string

const (
	PlanNewParamsPricesPriceNewPlanTieredWithProrationPriceCadenceAnnual     PlanNewParamsPricesPriceNewPlanTieredWithProrationPriceCadence = "annual"
	PlanNewParamsPricesPriceNewPlanTieredWithProrationPriceCadenceSemiAnnual PlanNewParamsPricesPriceNewPlanTieredWithProrationPriceCadence = "semi_annual"
	PlanNewParamsPricesPriceNewPlanTieredWithProrationPriceCadenceMonthly    PlanNewParamsPricesPriceNewPlanTieredWithProrationPriceCadence = "monthly"
	PlanNewParamsPricesPriceNewPlanTieredWithProrationPriceCadenceQuarterly  PlanNewParamsPricesPriceNewPlanTieredWithProrationPriceCadence = "quarterly"
	PlanNewParamsPricesPriceNewPlanTieredWithProrationPriceCadenceOneTime    PlanNewParamsPricesPriceNewPlanTieredWithProrationPriceCadence = "one_time"
	PlanNewParamsPricesPriceNewPlanTieredWithProrationPriceCadenceCustom     PlanNewParamsPricesPriceNewPlanTieredWithProrationPriceCadence = "custom"
)

func (r PlanNewParamsPricesPriceNewPlanTieredWithProrationPriceCadence) IsKnown() bool {
	switch r {
	case PlanNewParamsPricesPriceNewPlanTieredWithProrationPriceCadenceAnnual, PlanNewParamsPricesPriceNewPlanTieredWithProrationPriceCadenceSemiAnnual, PlanNewParamsPricesPriceNewPlanTieredWithProrationPriceCadenceMonthly, PlanNewParamsPricesPriceNewPlanTieredWithProrationPriceCadenceQuarterly, PlanNewParamsPricesPriceNewPlanTieredWithProrationPriceCadenceOneTime, PlanNewParamsPricesPriceNewPlanTieredWithProrationPriceCadenceCustom:
		return true
	}
	return false
}

// The pricing model type
type PlanNewParamsPricesPriceNewPlanTieredWithProrationPriceModelType string

const (
	PlanNewParamsPricesPriceNewPlanTieredWithProrationPriceModelTypeTieredWithProration PlanNewParamsPricesPriceNewPlanTieredWithProrationPriceModelType = "tiered_with_proration"
)

func (r PlanNewParamsPricesPriceNewPlanTieredWithProrationPriceModelType) IsKnown() bool {
	switch r {
	case PlanNewParamsPricesPriceNewPlanTieredWithProrationPriceModelTypeTieredWithProration:
		return true
	}
	return false
}

// Configuration for tiered_with_proration pricing
type PlanNewParamsPricesPriceNewPlanTieredWithProrationPriceTieredWithProrationConfig struct {
	// Tiers for rating based on total usage quantities into the specified tier with
	// proration
	Tiers param.Field[[]PlanNewParamsPricesPriceNewPlanTieredWithProrationPriceTieredWithProrationConfigTier] `json:"tiers,required"`
}

func (r PlanNewParamsPricesPriceNewPlanTieredWithProrationPriceTieredWithProrationConfig) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Configuration for a single tiered with proration tier
type PlanNewParamsPricesPriceNewPlanTieredWithProrationPriceTieredWithProrationConfigTier struct {
	// Inclusive tier starting value
	TierLowerBound param.Field[string] `json:"tier_lower_bound,required"`
	// Amount per unit
	UnitAmount param.Field[string] `json:"unit_amount,required"`
}

func (r PlanNewParamsPricesPriceNewPlanTieredWithProrationPriceTieredWithProrationConfigTier) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type PlanNewParamsPricesPriceNewPlanTieredWithProrationPriceConversionRateConfig struct {
	ConversionRateType param.Field[PlanNewParamsPricesPriceNewPlanTieredWithProrationPriceConversionRateConfigConversionRateType] `json:"conversion_rate_type,required"`
	TieredConfig       param.Field[shared.ConversionRateTieredConfigParam]                                                        `json:"tiered_config"`
	UnitConfig         param.Field[shared.ConversionRateUnitConfigParam]                                                          `json:"unit_config"`
}

func (r PlanNewParamsPricesPriceNewPlanTieredWithProrationPriceConversionRateConfig) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r PlanNewParamsPricesPriceNewPlanTieredWithProrationPriceConversionRateConfig) ImplementsPlanNewParamsPricesPriceNewPlanTieredWithProrationPriceConversionRateConfigUnion() {
}

// Satisfied by [shared.UnitConversionRateConfigParam],
// [shared.TieredConversionRateConfigParam],
// [PlanNewParamsPricesPriceNewPlanTieredWithProrationPriceConversionRateConfig].
type PlanNewParamsPricesPriceNewPlanTieredWithProrationPriceConversionRateConfigUnion interface {
	ImplementsPlanNewParamsPricesPriceNewPlanTieredWithProrationPriceConversionRateConfigUnion()
}

type PlanNewParamsPricesPriceNewPlanTieredWithProrationPriceConversionRateConfigConversionRateType string

const (
	PlanNewParamsPricesPriceNewPlanTieredWithProrationPriceConversionRateConfigConversionRateTypeUnit   PlanNewParamsPricesPriceNewPlanTieredWithProrationPriceConversionRateConfigConversionRateType = "unit"
	PlanNewParamsPricesPriceNewPlanTieredWithProrationPriceConversionRateConfigConversionRateTypeTiered PlanNewParamsPricesPriceNewPlanTieredWithProrationPriceConversionRateConfigConversionRateType = "tiered"
)

func (r PlanNewParamsPricesPriceNewPlanTieredWithProrationPriceConversionRateConfigConversionRateType) IsKnown() bool {
	switch r {
	case PlanNewParamsPricesPriceNewPlanTieredWithProrationPriceConversionRateConfigConversionRateTypeUnit, PlanNewParamsPricesPriceNewPlanTieredWithProrationPriceConversionRateConfigConversionRateTypeTiered:
		return true
	}
	return false
}

type PlanNewParamsPricesPriceNewPlanGroupedWithMinMaxThresholdsPrice struct {
	// The cadence to bill for this price on.
	Cadence param.Field[PlanNewParamsPricesPriceNewPlanGroupedWithMinMaxThresholdsPriceCadence] `json:"cadence,required"`
	// Configuration for grouped_with_min_max_thresholds pricing
	GroupedWithMinMaxThresholdsConfig param.Field[PlanNewParamsPricesPriceNewPlanGroupedWithMinMaxThresholdsPriceGroupedWithMinMaxThresholdsConfig] `json:"grouped_with_min_max_thresholds_config,required"`
	// The id of the item the price will be associated with.
	ItemID param.Field[string] `json:"item_id,required"`
	// The pricing model type
	ModelType param.Field[PlanNewParamsPricesPriceNewPlanGroupedWithMinMaxThresholdsPriceModelType] `json:"model_type,required"`
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
	ConversionRateConfig param.Field[PlanNewParamsPricesPriceNewPlanGroupedWithMinMaxThresholdsPriceConversionRateConfigUnion] `json:"conversion_rate_config"`
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

func (r PlanNewParamsPricesPriceNewPlanGroupedWithMinMaxThresholdsPrice) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r PlanNewParamsPricesPriceNewPlanGroupedWithMinMaxThresholdsPrice) ImplementsPlanNewParamsPricesPriceUnion() {
}

// The cadence to bill for this price on.
type PlanNewParamsPricesPriceNewPlanGroupedWithMinMaxThresholdsPriceCadence string

const (
	PlanNewParamsPricesPriceNewPlanGroupedWithMinMaxThresholdsPriceCadenceAnnual     PlanNewParamsPricesPriceNewPlanGroupedWithMinMaxThresholdsPriceCadence = "annual"
	PlanNewParamsPricesPriceNewPlanGroupedWithMinMaxThresholdsPriceCadenceSemiAnnual PlanNewParamsPricesPriceNewPlanGroupedWithMinMaxThresholdsPriceCadence = "semi_annual"
	PlanNewParamsPricesPriceNewPlanGroupedWithMinMaxThresholdsPriceCadenceMonthly    PlanNewParamsPricesPriceNewPlanGroupedWithMinMaxThresholdsPriceCadence = "monthly"
	PlanNewParamsPricesPriceNewPlanGroupedWithMinMaxThresholdsPriceCadenceQuarterly  PlanNewParamsPricesPriceNewPlanGroupedWithMinMaxThresholdsPriceCadence = "quarterly"
	PlanNewParamsPricesPriceNewPlanGroupedWithMinMaxThresholdsPriceCadenceOneTime    PlanNewParamsPricesPriceNewPlanGroupedWithMinMaxThresholdsPriceCadence = "one_time"
	PlanNewParamsPricesPriceNewPlanGroupedWithMinMaxThresholdsPriceCadenceCustom     PlanNewParamsPricesPriceNewPlanGroupedWithMinMaxThresholdsPriceCadence = "custom"
)

func (r PlanNewParamsPricesPriceNewPlanGroupedWithMinMaxThresholdsPriceCadence) IsKnown() bool {
	switch r {
	case PlanNewParamsPricesPriceNewPlanGroupedWithMinMaxThresholdsPriceCadenceAnnual, PlanNewParamsPricesPriceNewPlanGroupedWithMinMaxThresholdsPriceCadenceSemiAnnual, PlanNewParamsPricesPriceNewPlanGroupedWithMinMaxThresholdsPriceCadenceMonthly, PlanNewParamsPricesPriceNewPlanGroupedWithMinMaxThresholdsPriceCadenceQuarterly, PlanNewParamsPricesPriceNewPlanGroupedWithMinMaxThresholdsPriceCadenceOneTime, PlanNewParamsPricesPriceNewPlanGroupedWithMinMaxThresholdsPriceCadenceCustom:
		return true
	}
	return false
}

// Configuration for grouped_with_min_max_thresholds pricing
type PlanNewParamsPricesPriceNewPlanGroupedWithMinMaxThresholdsPriceGroupedWithMinMaxThresholdsConfig struct {
	// The event property used to group before applying thresholds
	GroupingKey param.Field[string] `json:"grouping_key,required"`
	// The maximum amount to charge each group
	MaximumCharge param.Field[string] `json:"maximum_charge,required"`
	// The minimum amount to charge each group, regardless of usage
	MinimumCharge param.Field[string] `json:"minimum_charge,required"`
	// The base price charged per group
	PerUnitRate param.Field[string] `json:"per_unit_rate,required"`
}

func (r PlanNewParamsPricesPriceNewPlanGroupedWithMinMaxThresholdsPriceGroupedWithMinMaxThresholdsConfig) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// The pricing model type
type PlanNewParamsPricesPriceNewPlanGroupedWithMinMaxThresholdsPriceModelType string

const (
	PlanNewParamsPricesPriceNewPlanGroupedWithMinMaxThresholdsPriceModelTypeGroupedWithMinMaxThresholds PlanNewParamsPricesPriceNewPlanGroupedWithMinMaxThresholdsPriceModelType = "grouped_with_min_max_thresholds"
)

func (r PlanNewParamsPricesPriceNewPlanGroupedWithMinMaxThresholdsPriceModelType) IsKnown() bool {
	switch r {
	case PlanNewParamsPricesPriceNewPlanGroupedWithMinMaxThresholdsPriceModelTypeGroupedWithMinMaxThresholds:
		return true
	}
	return false
}

type PlanNewParamsPricesPriceNewPlanGroupedWithMinMaxThresholdsPriceConversionRateConfig struct {
	ConversionRateType param.Field[PlanNewParamsPricesPriceNewPlanGroupedWithMinMaxThresholdsPriceConversionRateConfigConversionRateType] `json:"conversion_rate_type,required"`
	TieredConfig       param.Field[shared.ConversionRateTieredConfigParam]                                                                `json:"tiered_config"`
	UnitConfig         param.Field[shared.ConversionRateUnitConfigParam]                                                                  `json:"unit_config"`
}

func (r PlanNewParamsPricesPriceNewPlanGroupedWithMinMaxThresholdsPriceConversionRateConfig) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r PlanNewParamsPricesPriceNewPlanGroupedWithMinMaxThresholdsPriceConversionRateConfig) ImplementsPlanNewParamsPricesPriceNewPlanGroupedWithMinMaxThresholdsPriceConversionRateConfigUnion() {
}

// Satisfied by [shared.UnitConversionRateConfigParam],
// [shared.TieredConversionRateConfigParam],
// [PlanNewParamsPricesPriceNewPlanGroupedWithMinMaxThresholdsPriceConversionRateConfig].
type PlanNewParamsPricesPriceNewPlanGroupedWithMinMaxThresholdsPriceConversionRateConfigUnion interface {
	ImplementsPlanNewParamsPricesPriceNewPlanGroupedWithMinMaxThresholdsPriceConversionRateConfigUnion()
}

type PlanNewParamsPricesPriceNewPlanGroupedWithMinMaxThresholdsPriceConversionRateConfigConversionRateType string

const (
	PlanNewParamsPricesPriceNewPlanGroupedWithMinMaxThresholdsPriceConversionRateConfigConversionRateTypeUnit   PlanNewParamsPricesPriceNewPlanGroupedWithMinMaxThresholdsPriceConversionRateConfigConversionRateType = "unit"
	PlanNewParamsPricesPriceNewPlanGroupedWithMinMaxThresholdsPriceConversionRateConfigConversionRateTypeTiered PlanNewParamsPricesPriceNewPlanGroupedWithMinMaxThresholdsPriceConversionRateConfigConversionRateType = "tiered"
)

func (r PlanNewParamsPricesPriceNewPlanGroupedWithMinMaxThresholdsPriceConversionRateConfigConversionRateType) IsKnown() bool {
	switch r {
	case PlanNewParamsPricesPriceNewPlanGroupedWithMinMaxThresholdsPriceConversionRateConfigConversionRateTypeUnit, PlanNewParamsPricesPriceNewPlanGroupedWithMinMaxThresholdsPriceConversionRateConfigConversionRateTypeTiered:
		return true
	}
	return false
}

type PlanNewParamsPricesPriceNewPlanCumulativeGroupedAllocationPrice struct {
	// The cadence to bill for this price on.
	Cadence param.Field[PlanNewParamsPricesPriceNewPlanCumulativeGroupedAllocationPriceCadence] `json:"cadence,required"`
	// Configuration for cumulative_grouped_allocation pricing
	CumulativeGroupedAllocationConfig param.Field[PlanNewParamsPricesPriceNewPlanCumulativeGroupedAllocationPriceCumulativeGroupedAllocationConfig] `json:"cumulative_grouped_allocation_config,required"`
	// The id of the item the price will be associated with.
	ItemID param.Field[string] `json:"item_id,required"`
	// The pricing model type
	ModelType param.Field[PlanNewParamsPricesPriceNewPlanCumulativeGroupedAllocationPriceModelType] `json:"model_type,required"`
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
	ConversionRateConfig param.Field[PlanNewParamsPricesPriceNewPlanCumulativeGroupedAllocationPriceConversionRateConfigUnion] `json:"conversion_rate_config"`
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

func (r PlanNewParamsPricesPriceNewPlanCumulativeGroupedAllocationPrice) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r PlanNewParamsPricesPriceNewPlanCumulativeGroupedAllocationPrice) ImplementsPlanNewParamsPricesPriceUnion() {
}

// The cadence to bill for this price on.
type PlanNewParamsPricesPriceNewPlanCumulativeGroupedAllocationPriceCadence string

const (
	PlanNewParamsPricesPriceNewPlanCumulativeGroupedAllocationPriceCadenceAnnual     PlanNewParamsPricesPriceNewPlanCumulativeGroupedAllocationPriceCadence = "annual"
	PlanNewParamsPricesPriceNewPlanCumulativeGroupedAllocationPriceCadenceSemiAnnual PlanNewParamsPricesPriceNewPlanCumulativeGroupedAllocationPriceCadence = "semi_annual"
	PlanNewParamsPricesPriceNewPlanCumulativeGroupedAllocationPriceCadenceMonthly    PlanNewParamsPricesPriceNewPlanCumulativeGroupedAllocationPriceCadence = "monthly"
	PlanNewParamsPricesPriceNewPlanCumulativeGroupedAllocationPriceCadenceQuarterly  PlanNewParamsPricesPriceNewPlanCumulativeGroupedAllocationPriceCadence = "quarterly"
	PlanNewParamsPricesPriceNewPlanCumulativeGroupedAllocationPriceCadenceOneTime    PlanNewParamsPricesPriceNewPlanCumulativeGroupedAllocationPriceCadence = "one_time"
	PlanNewParamsPricesPriceNewPlanCumulativeGroupedAllocationPriceCadenceCustom     PlanNewParamsPricesPriceNewPlanCumulativeGroupedAllocationPriceCadence = "custom"
)

func (r PlanNewParamsPricesPriceNewPlanCumulativeGroupedAllocationPriceCadence) IsKnown() bool {
	switch r {
	case PlanNewParamsPricesPriceNewPlanCumulativeGroupedAllocationPriceCadenceAnnual, PlanNewParamsPricesPriceNewPlanCumulativeGroupedAllocationPriceCadenceSemiAnnual, PlanNewParamsPricesPriceNewPlanCumulativeGroupedAllocationPriceCadenceMonthly, PlanNewParamsPricesPriceNewPlanCumulativeGroupedAllocationPriceCadenceQuarterly, PlanNewParamsPricesPriceNewPlanCumulativeGroupedAllocationPriceCadenceOneTime, PlanNewParamsPricesPriceNewPlanCumulativeGroupedAllocationPriceCadenceCustom:
		return true
	}
	return false
}

// Configuration for cumulative_grouped_allocation pricing
type PlanNewParamsPricesPriceNewPlanCumulativeGroupedAllocationPriceCumulativeGroupedAllocationConfig struct {
	// The overall allocation across all groups
	CumulativeAllocation param.Field[string] `json:"cumulative_allocation,required"`
	// The allocation per individual group
	GroupAllocation param.Field[string] `json:"group_allocation,required"`
	// The event property used to group usage before applying allocations
	GroupingKey param.Field[string] `json:"grouping_key,required"`
	// The amount to charge for each unit outside of the allocation
	UnitAmount param.Field[string] `json:"unit_amount,required"`
}

func (r PlanNewParamsPricesPriceNewPlanCumulativeGroupedAllocationPriceCumulativeGroupedAllocationConfig) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// The pricing model type
type PlanNewParamsPricesPriceNewPlanCumulativeGroupedAllocationPriceModelType string

const (
	PlanNewParamsPricesPriceNewPlanCumulativeGroupedAllocationPriceModelTypeCumulativeGroupedAllocation PlanNewParamsPricesPriceNewPlanCumulativeGroupedAllocationPriceModelType = "cumulative_grouped_allocation"
)

func (r PlanNewParamsPricesPriceNewPlanCumulativeGroupedAllocationPriceModelType) IsKnown() bool {
	switch r {
	case PlanNewParamsPricesPriceNewPlanCumulativeGroupedAllocationPriceModelTypeCumulativeGroupedAllocation:
		return true
	}
	return false
}

type PlanNewParamsPricesPriceNewPlanCumulativeGroupedAllocationPriceConversionRateConfig struct {
	ConversionRateType param.Field[PlanNewParamsPricesPriceNewPlanCumulativeGroupedAllocationPriceConversionRateConfigConversionRateType] `json:"conversion_rate_type,required"`
	TieredConfig       param.Field[shared.ConversionRateTieredConfigParam]                                                                `json:"tiered_config"`
	UnitConfig         param.Field[shared.ConversionRateUnitConfigParam]                                                                  `json:"unit_config"`
}

func (r PlanNewParamsPricesPriceNewPlanCumulativeGroupedAllocationPriceConversionRateConfig) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r PlanNewParamsPricesPriceNewPlanCumulativeGroupedAllocationPriceConversionRateConfig) ImplementsPlanNewParamsPricesPriceNewPlanCumulativeGroupedAllocationPriceConversionRateConfigUnion() {
}

// Satisfied by [shared.UnitConversionRateConfigParam],
// [shared.TieredConversionRateConfigParam],
// [PlanNewParamsPricesPriceNewPlanCumulativeGroupedAllocationPriceConversionRateConfig].
type PlanNewParamsPricesPriceNewPlanCumulativeGroupedAllocationPriceConversionRateConfigUnion interface {
	ImplementsPlanNewParamsPricesPriceNewPlanCumulativeGroupedAllocationPriceConversionRateConfigUnion()
}

type PlanNewParamsPricesPriceNewPlanCumulativeGroupedAllocationPriceConversionRateConfigConversionRateType string

const (
	PlanNewParamsPricesPriceNewPlanCumulativeGroupedAllocationPriceConversionRateConfigConversionRateTypeUnit   PlanNewParamsPricesPriceNewPlanCumulativeGroupedAllocationPriceConversionRateConfigConversionRateType = "unit"
	PlanNewParamsPricesPriceNewPlanCumulativeGroupedAllocationPriceConversionRateConfigConversionRateTypeTiered PlanNewParamsPricesPriceNewPlanCumulativeGroupedAllocationPriceConversionRateConfigConversionRateType = "tiered"
)

func (r PlanNewParamsPricesPriceNewPlanCumulativeGroupedAllocationPriceConversionRateConfigConversionRateType) IsKnown() bool {
	switch r {
	case PlanNewParamsPricesPriceNewPlanCumulativeGroupedAllocationPriceConversionRateConfigConversionRateTypeUnit, PlanNewParamsPricesPriceNewPlanCumulativeGroupedAllocationPriceConversionRateConfigConversionRateTypeTiered:
		return true
	}
	return false
}

type PlanNewParamsPricesPriceNewPlanPercentCompositePrice struct {
	// The cadence to bill for this price on.
	Cadence param.Field[PlanNewParamsPricesPriceNewPlanPercentCompositePriceCadence] `json:"cadence,required"`
	// The id of the item the price will be associated with.
	ItemID param.Field[string] `json:"item_id,required"`
	// The pricing model type
	ModelType param.Field[PlanNewParamsPricesPriceNewPlanPercentCompositePriceModelType] `json:"model_type,required"`
	// The name of the price.
	Name param.Field[string] `json:"name,required"`
	// Configuration for percent pricing
	PercentConfig param.Field[PlanNewParamsPricesPriceNewPlanPercentCompositePricePercentConfig] `json:"percent_config,required"`
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
	ConversionRateConfig param.Field[PlanNewParamsPricesPriceNewPlanPercentCompositePriceConversionRateConfigUnion] `json:"conversion_rate_config"`
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

func (r PlanNewParamsPricesPriceNewPlanPercentCompositePrice) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r PlanNewParamsPricesPriceNewPlanPercentCompositePrice) ImplementsPlanNewParamsPricesPriceUnion() {
}

// The cadence to bill for this price on.
type PlanNewParamsPricesPriceNewPlanPercentCompositePriceCadence string

const (
	PlanNewParamsPricesPriceNewPlanPercentCompositePriceCadenceAnnual     PlanNewParamsPricesPriceNewPlanPercentCompositePriceCadence = "annual"
	PlanNewParamsPricesPriceNewPlanPercentCompositePriceCadenceSemiAnnual PlanNewParamsPricesPriceNewPlanPercentCompositePriceCadence = "semi_annual"
	PlanNewParamsPricesPriceNewPlanPercentCompositePriceCadenceMonthly    PlanNewParamsPricesPriceNewPlanPercentCompositePriceCadence = "monthly"
	PlanNewParamsPricesPriceNewPlanPercentCompositePriceCadenceQuarterly  PlanNewParamsPricesPriceNewPlanPercentCompositePriceCadence = "quarterly"
	PlanNewParamsPricesPriceNewPlanPercentCompositePriceCadenceOneTime    PlanNewParamsPricesPriceNewPlanPercentCompositePriceCadence = "one_time"
	PlanNewParamsPricesPriceNewPlanPercentCompositePriceCadenceCustom     PlanNewParamsPricesPriceNewPlanPercentCompositePriceCadence = "custom"
)

func (r PlanNewParamsPricesPriceNewPlanPercentCompositePriceCadence) IsKnown() bool {
	switch r {
	case PlanNewParamsPricesPriceNewPlanPercentCompositePriceCadenceAnnual, PlanNewParamsPricesPriceNewPlanPercentCompositePriceCadenceSemiAnnual, PlanNewParamsPricesPriceNewPlanPercentCompositePriceCadenceMonthly, PlanNewParamsPricesPriceNewPlanPercentCompositePriceCadenceQuarterly, PlanNewParamsPricesPriceNewPlanPercentCompositePriceCadenceOneTime, PlanNewParamsPricesPriceNewPlanPercentCompositePriceCadenceCustom:
		return true
	}
	return false
}

// The pricing model type
type PlanNewParamsPricesPriceNewPlanPercentCompositePriceModelType string

const (
	PlanNewParamsPricesPriceNewPlanPercentCompositePriceModelTypePercent PlanNewParamsPricesPriceNewPlanPercentCompositePriceModelType = "percent"
)

func (r PlanNewParamsPricesPriceNewPlanPercentCompositePriceModelType) IsKnown() bool {
	switch r {
	case PlanNewParamsPricesPriceNewPlanPercentCompositePriceModelTypePercent:
		return true
	}
	return false
}

// Configuration for percent pricing
type PlanNewParamsPricesPriceNewPlanPercentCompositePricePercentConfig struct {
	// What percent of the component subtotals to charge
	Percent param.Field[float64] `json:"percent,required"`
}

func (r PlanNewParamsPricesPriceNewPlanPercentCompositePricePercentConfig) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type PlanNewParamsPricesPriceNewPlanPercentCompositePriceConversionRateConfig struct {
	ConversionRateType param.Field[PlanNewParamsPricesPriceNewPlanPercentCompositePriceConversionRateConfigConversionRateType] `json:"conversion_rate_type,required"`
	TieredConfig       param.Field[shared.ConversionRateTieredConfigParam]                                                     `json:"tiered_config"`
	UnitConfig         param.Field[shared.ConversionRateUnitConfigParam]                                                       `json:"unit_config"`
}

func (r PlanNewParamsPricesPriceNewPlanPercentCompositePriceConversionRateConfig) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r PlanNewParamsPricesPriceNewPlanPercentCompositePriceConversionRateConfig) ImplementsPlanNewParamsPricesPriceNewPlanPercentCompositePriceConversionRateConfigUnion() {
}

// Satisfied by [shared.UnitConversionRateConfigParam],
// [shared.TieredConversionRateConfigParam],
// [PlanNewParamsPricesPriceNewPlanPercentCompositePriceConversionRateConfig].
type PlanNewParamsPricesPriceNewPlanPercentCompositePriceConversionRateConfigUnion interface {
	ImplementsPlanNewParamsPricesPriceNewPlanPercentCompositePriceConversionRateConfigUnion()
}

type PlanNewParamsPricesPriceNewPlanPercentCompositePriceConversionRateConfigConversionRateType string

const (
	PlanNewParamsPricesPriceNewPlanPercentCompositePriceConversionRateConfigConversionRateTypeUnit   PlanNewParamsPricesPriceNewPlanPercentCompositePriceConversionRateConfigConversionRateType = "unit"
	PlanNewParamsPricesPriceNewPlanPercentCompositePriceConversionRateConfigConversionRateTypeTiered PlanNewParamsPricesPriceNewPlanPercentCompositePriceConversionRateConfigConversionRateType = "tiered"
)

func (r PlanNewParamsPricesPriceNewPlanPercentCompositePriceConversionRateConfigConversionRateType) IsKnown() bool {
	switch r {
	case PlanNewParamsPricesPriceNewPlanPercentCompositePriceConversionRateConfigConversionRateTypeUnit, PlanNewParamsPricesPriceNewPlanPercentCompositePriceConversionRateConfigConversionRateTypeTiered:
		return true
	}
	return false
}

type PlanNewParamsPricesPriceNewPlanEventOutputPrice struct {
	// The cadence to bill for this price on.
	Cadence param.Field[PlanNewParamsPricesPriceNewPlanEventOutputPriceCadence] `json:"cadence,required"`
	// Configuration for event_output pricing
	EventOutputConfig param.Field[PlanNewParamsPricesPriceNewPlanEventOutputPriceEventOutputConfig] `json:"event_output_config,required"`
	// The id of the item the price will be associated with.
	ItemID param.Field[string] `json:"item_id,required"`
	// The pricing model type
	ModelType param.Field[PlanNewParamsPricesPriceNewPlanEventOutputPriceModelType] `json:"model_type,required"`
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
	ConversionRateConfig param.Field[PlanNewParamsPricesPriceNewPlanEventOutputPriceConversionRateConfigUnion] `json:"conversion_rate_config"`
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

func (r PlanNewParamsPricesPriceNewPlanEventOutputPrice) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r PlanNewParamsPricesPriceNewPlanEventOutputPrice) ImplementsPlanNewParamsPricesPriceUnion() {}

// The cadence to bill for this price on.
type PlanNewParamsPricesPriceNewPlanEventOutputPriceCadence string

const (
	PlanNewParamsPricesPriceNewPlanEventOutputPriceCadenceAnnual     PlanNewParamsPricesPriceNewPlanEventOutputPriceCadence = "annual"
	PlanNewParamsPricesPriceNewPlanEventOutputPriceCadenceSemiAnnual PlanNewParamsPricesPriceNewPlanEventOutputPriceCadence = "semi_annual"
	PlanNewParamsPricesPriceNewPlanEventOutputPriceCadenceMonthly    PlanNewParamsPricesPriceNewPlanEventOutputPriceCadence = "monthly"
	PlanNewParamsPricesPriceNewPlanEventOutputPriceCadenceQuarterly  PlanNewParamsPricesPriceNewPlanEventOutputPriceCadence = "quarterly"
	PlanNewParamsPricesPriceNewPlanEventOutputPriceCadenceOneTime    PlanNewParamsPricesPriceNewPlanEventOutputPriceCadence = "one_time"
	PlanNewParamsPricesPriceNewPlanEventOutputPriceCadenceCustom     PlanNewParamsPricesPriceNewPlanEventOutputPriceCadence = "custom"
)

func (r PlanNewParamsPricesPriceNewPlanEventOutputPriceCadence) IsKnown() bool {
	switch r {
	case PlanNewParamsPricesPriceNewPlanEventOutputPriceCadenceAnnual, PlanNewParamsPricesPriceNewPlanEventOutputPriceCadenceSemiAnnual, PlanNewParamsPricesPriceNewPlanEventOutputPriceCadenceMonthly, PlanNewParamsPricesPriceNewPlanEventOutputPriceCadenceQuarterly, PlanNewParamsPricesPriceNewPlanEventOutputPriceCadenceOneTime, PlanNewParamsPricesPriceNewPlanEventOutputPriceCadenceCustom:
		return true
	}
	return false
}

// Configuration for event_output pricing
type PlanNewParamsPricesPriceNewPlanEventOutputPriceEventOutputConfig struct {
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

func (r PlanNewParamsPricesPriceNewPlanEventOutputPriceEventOutputConfig) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// The pricing model type
type PlanNewParamsPricesPriceNewPlanEventOutputPriceModelType string

const (
	PlanNewParamsPricesPriceNewPlanEventOutputPriceModelTypeEventOutput PlanNewParamsPricesPriceNewPlanEventOutputPriceModelType = "event_output"
)

func (r PlanNewParamsPricesPriceNewPlanEventOutputPriceModelType) IsKnown() bool {
	switch r {
	case PlanNewParamsPricesPriceNewPlanEventOutputPriceModelTypeEventOutput:
		return true
	}
	return false
}

type PlanNewParamsPricesPriceNewPlanEventOutputPriceConversionRateConfig struct {
	ConversionRateType param.Field[PlanNewParamsPricesPriceNewPlanEventOutputPriceConversionRateConfigConversionRateType] `json:"conversion_rate_type,required"`
	TieredConfig       param.Field[shared.ConversionRateTieredConfigParam]                                                `json:"tiered_config"`
	UnitConfig         param.Field[shared.ConversionRateUnitConfigParam]                                                  `json:"unit_config"`
}

func (r PlanNewParamsPricesPriceNewPlanEventOutputPriceConversionRateConfig) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r PlanNewParamsPricesPriceNewPlanEventOutputPriceConversionRateConfig) ImplementsPlanNewParamsPricesPriceNewPlanEventOutputPriceConversionRateConfigUnion() {
}

// Satisfied by [shared.UnitConversionRateConfigParam],
// [shared.TieredConversionRateConfigParam],
// [PlanNewParamsPricesPriceNewPlanEventOutputPriceConversionRateConfig].
type PlanNewParamsPricesPriceNewPlanEventOutputPriceConversionRateConfigUnion interface {
	ImplementsPlanNewParamsPricesPriceNewPlanEventOutputPriceConversionRateConfigUnion()
}

type PlanNewParamsPricesPriceNewPlanEventOutputPriceConversionRateConfigConversionRateType string

const (
	PlanNewParamsPricesPriceNewPlanEventOutputPriceConversionRateConfigConversionRateTypeUnit   PlanNewParamsPricesPriceNewPlanEventOutputPriceConversionRateConfigConversionRateType = "unit"
	PlanNewParamsPricesPriceNewPlanEventOutputPriceConversionRateConfigConversionRateTypeTiered PlanNewParamsPricesPriceNewPlanEventOutputPriceConversionRateConfigConversionRateType = "tiered"
)

func (r PlanNewParamsPricesPriceNewPlanEventOutputPriceConversionRateConfigConversionRateType) IsKnown() bool {
	switch r {
	case PlanNewParamsPricesPriceNewPlanEventOutputPriceConversionRateConfigConversionRateTypeUnit, PlanNewParamsPricesPriceNewPlanEventOutputPriceConversionRateConfigConversionRateTypeTiered:
		return true
	}
	return false
}

// The cadence to bill for this price on.
type PlanNewParamsPricesPriceCadence string

const (
	PlanNewParamsPricesPriceCadenceAnnual     PlanNewParamsPricesPriceCadence = "annual"
	PlanNewParamsPricesPriceCadenceSemiAnnual PlanNewParamsPricesPriceCadence = "semi_annual"
	PlanNewParamsPricesPriceCadenceMonthly    PlanNewParamsPricesPriceCadence = "monthly"
	PlanNewParamsPricesPriceCadenceQuarterly  PlanNewParamsPricesPriceCadence = "quarterly"
	PlanNewParamsPricesPriceCadenceOneTime    PlanNewParamsPricesPriceCadence = "one_time"
	PlanNewParamsPricesPriceCadenceCustom     PlanNewParamsPricesPriceCadence = "custom"
)

func (r PlanNewParamsPricesPriceCadence) IsKnown() bool {
	switch r {
	case PlanNewParamsPricesPriceCadenceAnnual, PlanNewParamsPricesPriceCadenceSemiAnnual, PlanNewParamsPricesPriceCadenceMonthly, PlanNewParamsPricesPriceCadenceQuarterly, PlanNewParamsPricesPriceCadenceOneTime, PlanNewParamsPricesPriceCadenceCustom:
		return true
	}
	return false
}

// The pricing model type
type PlanNewParamsPricesPriceModelType string

const (
	PlanNewParamsPricesPriceModelTypeUnit                            PlanNewParamsPricesPriceModelType = "unit"
	PlanNewParamsPricesPriceModelTypeTiered                          PlanNewParamsPricesPriceModelType = "tiered"
	PlanNewParamsPricesPriceModelTypeBulk                            PlanNewParamsPricesPriceModelType = "bulk"
	PlanNewParamsPricesPriceModelTypeBulkWithFilters                 PlanNewParamsPricesPriceModelType = "bulk_with_filters"
	PlanNewParamsPricesPriceModelTypePackage                         PlanNewParamsPricesPriceModelType = "package"
	PlanNewParamsPricesPriceModelTypeMatrix                          PlanNewParamsPricesPriceModelType = "matrix"
	PlanNewParamsPricesPriceModelTypeThresholdTotalAmount            PlanNewParamsPricesPriceModelType = "threshold_total_amount"
	PlanNewParamsPricesPriceModelTypeTieredPackage                   PlanNewParamsPricesPriceModelType = "tiered_package"
	PlanNewParamsPricesPriceModelTypeTieredWithMinimum               PlanNewParamsPricesPriceModelType = "tiered_with_minimum"
	PlanNewParamsPricesPriceModelTypeGroupedTiered                   PlanNewParamsPricesPriceModelType = "grouped_tiered"
	PlanNewParamsPricesPriceModelTypeTieredPackageWithMinimum        PlanNewParamsPricesPriceModelType = "tiered_package_with_minimum"
	PlanNewParamsPricesPriceModelTypePackageWithAllocation           PlanNewParamsPricesPriceModelType = "package_with_allocation"
	PlanNewParamsPricesPriceModelTypeUnitWithPercent                 PlanNewParamsPricesPriceModelType = "unit_with_percent"
	PlanNewParamsPricesPriceModelTypeMatrixWithAllocation            PlanNewParamsPricesPriceModelType = "matrix_with_allocation"
	PlanNewParamsPricesPriceModelTypeTieredWithProration             PlanNewParamsPricesPriceModelType = "tiered_with_proration"
	PlanNewParamsPricesPriceModelTypeUnitWithProration               PlanNewParamsPricesPriceModelType = "unit_with_proration"
	PlanNewParamsPricesPriceModelTypeGroupedAllocation               PlanNewParamsPricesPriceModelType = "grouped_allocation"
	PlanNewParamsPricesPriceModelTypeBulkWithProration               PlanNewParamsPricesPriceModelType = "bulk_with_proration"
	PlanNewParamsPricesPriceModelTypeGroupedWithProratedMinimum      PlanNewParamsPricesPriceModelType = "grouped_with_prorated_minimum"
	PlanNewParamsPricesPriceModelTypeGroupedWithMeteredMinimum       PlanNewParamsPricesPriceModelType = "grouped_with_metered_minimum"
	PlanNewParamsPricesPriceModelTypeGroupedWithMinMaxThresholds     PlanNewParamsPricesPriceModelType = "grouped_with_min_max_thresholds"
	PlanNewParamsPricesPriceModelTypeMatrixWithDisplayName           PlanNewParamsPricesPriceModelType = "matrix_with_display_name"
	PlanNewParamsPricesPriceModelTypeGroupedTieredPackage            PlanNewParamsPricesPriceModelType = "grouped_tiered_package"
	PlanNewParamsPricesPriceModelTypeMaxGroupTieredPackage           PlanNewParamsPricesPriceModelType = "max_group_tiered_package"
	PlanNewParamsPricesPriceModelTypeScalableMatrixWithUnitPricing   PlanNewParamsPricesPriceModelType = "scalable_matrix_with_unit_pricing"
	PlanNewParamsPricesPriceModelTypeScalableMatrixWithTieredPricing PlanNewParamsPricesPriceModelType = "scalable_matrix_with_tiered_pricing"
	PlanNewParamsPricesPriceModelTypeCumulativeGroupedBulk           PlanNewParamsPricesPriceModelType = "cumulative_grouped_bulk"
	PlanNewParamsPricesPriceModelTypeCumulativeGroupedAllocation     PlanNewParamsPricesPriceModelType = "cumulative_grouped_allocation"
	PlanNewParamsPricesPriceModelTypeMinimum                         PlanNewParamsPricesPriceModelType = "minimum"
	PlanNewParamsPricesPriceModelTypePercent                         PlanNewParamsPricesPriceModelType = "percent"
	PlanNewParamsPricesPriceModelTypeEventOutput                     PlanNewParamsPricesPriceModelType = "event_output"
)

func (r PlanNewParamsPricesPriceModelType) IsKnown() bool {
	switch r {
	case PlanNewParamsPricesPriceModelTypeUnit, PlanNewParamsPricesPriceModelTypeTiered, PlanNewParamsPricesPriceModelTypeBulk, PlanNewParamsPricesPriceModelTypeBulkWithFilters, PlanNewParamsPricesPriceModelTypePackage, PlanNewParamsPricesPriceModelTypeMatrix, PlanNewParamsPricesPriceModelTypeThresholdTotalAmount, PlanNewParamsPricesPriceModelTypeTieredPackage, PlanNewParamsPricesPriceModelTypeTieredWithMinimum, PlanNewParamsPricesPriceModelTypeGroupedTiered, PlanNewParamsPricesPriceModelTypeTieredPackageWithMinimum, PlanNewParamsPricesPriceModelTypePackageWithAllocation, PlanNewParamsPricesPriceModelTypeUnitWithPercent, PlanNewParamsPricesPriceModelTypeMatrixWithAllocation, PlanNewParamsPricesPriceModelTypeTieredWithProration, PlanNewParamsPricesPriceModelTypeUnitWithProration, PlanNewParamsPricesPriceModelTypeGroupedAllocation, PlanNewParamsPricesPriceModelTypeBulkWithProration, PlanNewParamsPricesPriceModelTypeGroupedWithProratedMinimum, PlanNewParamsPricesPriceModelTypeGroupedWithMeteredMinimum, PlanNewParamsPricesPriceModelTypeGroupedWithMinMaxThresholds, PlanNewParamsPricesPriceModelTypeMatrixWithDisplayName, PlanNewParamsPricesPriceModelTypeGroupedTieredPackage, PlanNewParamsPricesPriceModelTypeMaxGroupTieredPackage, PlanNewParamsPricesPriceModelTypeScalableMatrixWithUnitPricing, PlanNewParamsPricesPriceModelTypeScalableMatrixWithTieredPricing, PlanNewParamsPricesPriceModelTypeCumulativeGroupedBulk, PlanNewParamsPricesPriceModelTypeCumulativeGroupedAllocation, PlanNewParamsPricesPriceModelTypeMinimum, PlanNewParamsPricesPriceModelTypePercent, PlanNewParamsPricesPriceModelTypeEventOutput:
		return true
	}
	return false
}

type PlanNewParamsAdjustment struct {
	// The definition of a new adjustment to create and add to the plan.
	Adjustment param.Field[PlanNewParamsAdjustmentsAdjustmentUnion] `json:"adjustment,required"`
	// The phase to add this adjustment to.
	PlanPhaseOrder param.Field[int64] `json:"plan_phase_order"`
}

func (r PlanNewParamsAdjustment) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// The definition of a new adjustment to create and add to the plan.
type PlanNewParamsAdjustmentsAdjustment struct {
	AdjustmentType param.Field[PlanNewParamsAdjustmentsAdjustmentAdjustmentType] `json:"adjustment_type,required"`
	AmountDiscount param.Field[string]                                           `json:"amount_discount"`
	// If set, the adjustment will apply to every price on the subscription.
	AppliesToAll      param.Field[PlanNewParamsAdjustmentsAdjustmentAppliesToAll] `json:"applies_to_all"`
	AppliesToItemIDs  param.Field[interface{}]                                    `json:"applies_to_item_ids"`
	AppliesToPriceIDs param.Field[interface{}]                                    `json:"applies_to_price_ids"`
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
	PriceType     param.Field[PlanNewParamsAdjustmentsAdjustmentPriceType] `json:"price_type"`
	UsageDiscount param.Field[float64]                                     `json:"usage_discount"`
}

func (r PlanNewParamsAdjustmentsAdjustment) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r PlanNewParamsAdjustmentsAdjustment) ImplementsPlanNewParamsAdjustmentsAdjustmentUnion() {}

// The definition of a new adjustment to create and add to the plan.
//
// Satisfied by [shared.NewPercentageDiscountParam],
// [shared.NewUsageDiscountParam], [shared.NewAmountDiscountParam],
// [shared.NewMinimumParam], [shared.NewMaximumParam],
// [PlanNewParamsAdjustmentsAdjustment].
type PlanNewParamsAdjustmentsAdjustmentUnion interface {
	ImplementsPlanNewParamsAdjustmentsAdjustmentUnion()
}

type PlanNewParamsAdjustmentsAdjustmentAdjustmentType string

const (
	PlanNewParamsAdjustmentsAdjustmentAdjustmentTypePercentageDiscount PlanNewParamsAdjustmentsAdjustmentAdjustmentType = "percentage_discount"
	PlanNewParamsAdjustmentsAdjustmentAdjustmentTypeUsageDiscount      PlanNewParamsAdjustmentsAdjustmentAdjustmentType = "usage_discount"
	PlanNewParamsAdjustmentsAdjustmentAdjustmentTypeAmountDiscount     PlanNewParamsAdjustmentsAdjustmentAdjustmentType = "amount_discount"
	PlanNewParamsAdjustmentsAdjustmentAdjustmentTypeMinimum            PlanNewParamsAdjustmentsAdjustmentAdjustmentType = "minimum"
	PlanNewParamsAdjustmentsAdjustmentAdjustmentTypeMaximum            PlanNewParamsAdjustmentsAdjustmentAdjustmentType = "maximum"
)

func (r PlanNewParamsAdjustmentsAdjustmentAdjustmentType) IsKnown() bool {
	switch r {
	case PlanNewParamsAdjustmentsAdjustmentAdjustmentTypePercentageDiscount, PlanNewParamsAdjustmentsAdjustmentAdjustmentTypeUsageDiscount, PlanNewParamsAdjustmentsAdjustmentAdjustmentTypeAmountDiscount, PlanNewParamsAdjustmentsAdjustmentAdjustmentTypeMinimum, PlanNewParamsAdjustmentsAdjustmentAdjustmentTypeMaximum:
		return true
	}
	return false
}

// If set, the adjustment will apply to every price on the subscription.
type PlanNewParamsAdjustmentsAdjustmentAppliesToAll bool

const (
	PlanNewParamsAdjustmentsAdjustmentAppliesToAllTrue PlanNewParamsAdjustmentsAdjustmentAppliesToAll = true
)

func (r PlanNewParamsAdjustmentsAdjustmentAppliesToAll) IsKnown() bool {
	switch r {
	case PlanNewParamsAdjustmentsAdjustmentAppliesToAllTrue:
		return true
	}
	return false
}

// If set, only prices of the specified type will have the adjustment applied.
type PlanNewParamsAdjustmentsAdjustmentPriceType string

const (
	PlanNewParamsAdjustmentsAdjustmentPriceTypeUsage          PlanNewParamsAdjustmentsAdjustmentPriceType = "usage"
	PlanNewParamsAdjustmentsAdjustmentPriceTypeFixedInAdvance PlanNewParamsAdjustmentsAdjustmentPriceType = "fixed_in_advance"
	PlanNewParamsAdjustmentsAdjustmentPriceTypeFixedInArrears PlanNewParamsAdjustmentsAdjustmentPriceType = "fixed_in_arrears"
	PlanNewParamsAdjustmentsAdjustmentPriceTypeFixed          PlanNewParamsAdjustmentsAdjustmentPriceType = "fixed"
	PlanNewParamsAdjustmentsAdjustmentPriceTypeInArrears      PlanNewParamsAdjustmentsAdjustmentPriceType = "in_arrears"
)

func (r PlanNewParamsAdjustmentsAdjustmentPriceType) IsKnown() bool {
	switch r {
	case PlanNewParamsAdjustmentsAdjustmentPriceTypeUsage, PlanNewParamsAdjustmentsAdjustmentPriceTypeFixedInAdvance, PlanNewParamsAdjustmentsAdjustmentPriceTypeFixedInArrears, PlanNewParamsAdjustmentsAdjustmentPriceTypeFixed, PlanNewParamsAdjustmentsAdjustmentPriceTypeInArrears:
		return true
	}
	return false
}

type PlanNewParamsPlanPhase struct {
	// Determines the ordering of the phase in a plan's lifecycle. 1 = first phase.
	Order param.Field[int64] `json:"order,required"`
	// Align billing cycle day with phase start date.
	AlignBillingWithPhaseStartDate param.Field[bool] `json:"align_billing_with_phase_start_date"`
	// How many terms of length `duration_unit` this phase is active for. If null, this
	// phase is evergreen and active indefinitely
	Duration     param.Field[int64]                               `json:"duration"`
	DurationUnit param.Field[PlanNewParamsPlanPhasesDurationUnit] `json:"duration_unit"`
}

func (r PlanNewParamsPlanPhase) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type PlanNewParamsPlanPhasesDurationUnit string

const (
	PlanNewParamsPlanPhasesDurationUnitDaily      PlanNewParamsPlanPhasesDurationUnit = "daily"
	PlanNewParamsPlanPhasesDurationUnitMonthly    PlanNewParamsPlanPhasesDurationUnit = "monthly"
	PlanNewParamsPlanPhasesDurationUnitQuarterly  PlanNewParamsPlanPhasesDurationUnit = "quarterly"
	PlanNewParamsPlanPhasesDurationUnitSemiAnnual PlanNewParamsPlanPhasesDurationUnit = "semi_annual"
	PlanNewParamsPlanPhasesDurationUnitAnnual     PlanNewParamsPlanPhasesDurationUnit = "annual"
)

func (r PlanNewParamsPlanPhasesDurationUnit) IsKnown() bool {
	switch r {
	case PlanNewParamsPlanPhasesDurationUnitDaily, PlanNewParamsPlanPhasesDurationUnitMonthly, PlanNewParamsPlanPhasesDurationUnitQuarterly, PlanNewParamsPlanPhasesDurationUnitSemiAnnual, PlanNewParamsPlanPhasesDurationUnitAnnual:
		return true
	}
	return false
}

// The status of the plan to create (either active or draft). If not specified,
// this defaults to active.
type PlanNewParamsStatus string

const (
	PlanNewParamsStatusActive PlanNewParamsStatus = "active"
	PlanNewParamsStatusDraft  PlanNewParamsStatus = "draft"
)

func (r PlanNewParamsStatus) IsKnown() bool {
	switch r {
	case PlanNewParamsStatusActive, PlanNewParamsStatusDraft:
		return true
	}
	return false
}

type PlanUpdateParams struct {
	// An optional user-defined ID for this plan resource, used throughout the system
	// as an alias for this Plan. Use this field to identify a plan by an existing
	// identifier in your system.
	ExternalPlanID param.Field[string] `json:"external_plan_id"`
	// User-specified key/value pairs for the resource. Individual keys can be removed
	// by setting the value to `null`, and the entire metadata mapping can be cleared
	// by setting `metadata` to `null`.
	Metadata param.Field[map[string]string] `json:"metadata"`
}

func (r PlanUpdateParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type PlanListParams struct {
	CreatedAtGt  param.Field[time.Time] `query:"created_at[gt]" format:"date-time"`
	CreatedAtGte param.Field[time.Time] `query:"created_at[gte]" format:"date-time"`
	CreatedAtLt  param.Field[time.Time] `query:"created_at[lt]" format:"date-time"`
	CreatedAtLte param.Field[time.Time] `query:"created_at[lte]" format:"date-time"`
	// Cursor for pagination. This can be populated by the `next_cursor` value returned
	// from the initial request.
	Cursor param.Field[string] `query:"cursor"`
	// The number of items to fetch. Defaults to 20.
	Limit param.Field[int64] `query:"limit"`
	// The plan status to filter to ('active', 'archived', or 'draft').
	Status param.Field[PlanListParamsStatus] `query:"status"`
}

// URLQuery serializes [PlanListParams]'s query parameters as `url.Values`.
func (r PlanListParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatBrackets,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

// The plan status to filter to ('active', 'archived', or 'draft').
type PlanListParamsStatus string

const (
	PlanListParamsStatusActive   PlanListParamsStatus = "active"
	PlanListParamsStatusArchived PlanListParamsStatus = "archived"
	PlanListParamsStatusDraft    PlanListParamsStatus = "draft"
)

func (r PlanListParamsStatus) IsKnown() bool {
	switch r {
	case PlanListParamsStatusActive, PlanListParamsStatusArchived, PlanListParamsStatusDraft:
		return true
	}
	return false
}
