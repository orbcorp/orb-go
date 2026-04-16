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

// The [Plan](/core-concepts#plan-and-price) resource represents a plan that can be
// subscribed to by a customer. Plans define the billing behavior of the
// subscription. You can see more about how to configure prices in the
// [Price resource](/reference/price).
//
// PlanService contains methods and other services that help with interacting with
// the orb API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewPlanService] method instead.
type PlanService struct {
	Options []option.RequestOption
	// The [Plan](/core-concepts#plan-and-price) resource represents a plan that can be
	// subscribed to by a customer. Plans define the billing behavior of the
	// subscription. You can see more about how to configure prices in the
	// [Price resource](/reference/price).
	ExternalPlanID *PlanExternalPlanIDService
	// The [Plan](/core-concepts#plan-and-price) resource represents a plan that can be
	// subscribed to by a customer. Plans define the billing behavior of the
	// subscription. You can see more about how to configure prices in the
	// [Price resource](/reference/price).
	Migrations *PlanMigrationService
}

// NewPlanService generates a new service that applies the given options to each
// request. These options are applied after the parent client's options (if there
// is one), and before any request-specific options.
func NewPlanService(opts ...option.RequestOption) (r *PlanService) {
	r = &PlanService{}
	r.Options = opts
	r.ExternalPlanID = NewPlanExternalPlanIDService(opts...)
	r.Migrations = NewPlanMigrationService(opts...)
	return
}

// This endpoint allows creation of plans including their prices.
func (r *PlanService) New(ctx context.Context, body PlanNewParams, opts ...option.RequestOption) (res *Plan, err error) {
	opts = slices.Concat(r.Options, opts)
	path := "plans"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return res, err
}

// This endpoint can be used to update the `external_plan_id`, `description`, and
// `metadata` of an existing plan.
//
// Other fields on a plan are currently immutable.
func (r *PlanService) Update(ctx context.Context, planID string, body PlanUpdateParams, opts ...option.RequestOption) (res *Plan, err error) {
	opts = slices.Concat(r.Options, opts)
	if planID == "" {
		err = errors.New("missing required plan_id parameter")
		return nil, err
	}
	path := fmt.Sprintf("plans/%s", planID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPut, path, body, &res, opts...)
	return res, err
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
		return nil, err
	}
	path := fmt.Sprintf("plans/%s", planID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return res, err
}

// The [Plan](/core-concepts#plan-and-price) resource represents a plan that can be
// subscribed to by a customer. Plans define the billing behavior of the
// subscription. You can see more about how to configure prices in the
// [Price resource](/reference/price).
type Plan struct {
	ID string `json:"id" api:"required"`
	// Adjustments for this plan. If the plan has phases, this includes adjustments
	// across all phases of the plan.
	Adjustments []PlanAdjustment `json:"adjustments" api:"required"`
	// Legacy field representing the parent plan if the current plan is a 'child plan',
	// overriding prices from the parent.
	//
	// Deprecated: deprecated
	BasePlan PlanBasePlan `json:"base_plan" api:"required,nullable"`
	// Legacy field representing the parent plan ID if the current plan is a 'child
	// plan', overriding prices from the parent.
	//
	// Deprecated: deprecated
	BasePlanID string    `json:"base_plan_id" api:"required,nullable"`
	CreatedAt  time.Time `json:"created_at" api:"required" format:"date-time"`
	// An ISO 4217 currency string or custom pricing unit (`credits`) for this plan's
	// prices.
	//
	// Deprecated: deprecated
	Currency string `json:"currency" api:"required"`
	// The default memo text on the invoices corresponding to subscriptions on this
	// plan. Note that each subscription may configure its own memo.
	DefaultInvoiceMemo string `json:"default_invoice_memo" api:"required,nullable"`
	Description        string `json:"description" api:"required"`
	// Deprecated: deprecated
	Discount shared.Discount `json:"discount" api:"required,nullable"`
	// An optional user-defined ID for this plan resource, used throughout the system
	// as an alias for this Plan. Use this field to identify a plan by an existing
	// identifier in your system.
	ExternalPlanID string `json:"external_plan_id" api:"required,nullable"`
	// An ISO 4217 currency string for which this plan is billed in. Matches `currency`
	// unless `currency` is a custom pricing unit.
	InvoicingCurrency string `json:"invoicing_currency" api:"required"`
	// Deprecated: deprecated
	Maximum shared.Maximum `json:"maximum" api:"required,nullable"`
	// Deprecated: deprecated
	MaximumAmount string `json:"maximum_amount" api:"required,nullable"`
	// User specified key-value pairs for the resource. If not present, this defaults
	// to an empty dictionary. Individual keys can be removed by setting the value to
	// `null`, and the entire metadata mapping can be cleared by setting `metadata` to
	// `null`.
	Metadata map[string]string `json:"metadata" api:"required"`
	// Deprecated: deprecated
	Minimum shared.Minimum `json:"minimum" api:"required,nullable"`
	// Deprecated: deprecated
	MinimumAmount string `json:"minimum_amount" api:"required,nullable"`
	Name          string `json:"name" api:"required"`
	// Determines the difference between the invoice issue date and the due date. A
	// value of "0" here signifies that invoices are due on issue, whereas a value of
	// "30" means that the customer has a month to pay the invoice before its overdue.
	// Note that individual subscriptions or invoices may set a different net terms
	// configuration.
	NetTerms   int64           `json:"net_terms" api:"required,nullable"`
	PlanPhases []PlanPlanPhase `json:"plan_phases" api:"required,nullable"`
	// Prices for this plan. If the plan has phases, this includes prices across all
	// phases of the plan.
	Prices      []shared.Price  `json:"prices" api:"required"`
	Product     PlanProduct     `json:"product" api:"required"`
	Status      PlanStatus      `json:"status" api:"required"`
	TrialConfig PlanTrialConfig `json:"trial_config" api:"required"`
	Version     int64           `json:"version" api:"required"`
	JSON        planJSON        `json:"-"`
}

// planJSON contains the JSON metadata for the struct [Plan]
type planJSON struct {
	ID                 apijson.Field
	Adjustments        apijson.Field
	BasePlan           apijson.Field
	BasePlanID         apijson.Field
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
	ID             string                        `json:"id" api:"required"`
	AdjustmentType PlanAdjustmentsAdjustmentType `json:"adjustment_type" api:"required"`
	// This field can have the runtime type of [[]string].
	AppliesToPriceIDs interface{} `json:"applies_to_price_ids" api:"required"`
	// This field can have the runtime type of
	// [[]shared.PlanPhaseUsageDiscountAdjustmentFilter],
	// [[]shared.PlanPhaseAmountDiscountAdjustmentFilter],
	// [[]shared.PlanPhasePercentageDiscountAdjustmentFilter],
	// [[]shared.PlanPhaseMinimumAdjustmentFilter],
	// [[]shared.PlanPhaseMaximumAdjustmentFilter].
	Filters interface{} `json:"filters" api:"required"`
	// True for adjustments that apply to an entire invoice, false for adjustments that
	// apply to only one price.
	IsInvoiceLevel bool `json:"is_invoice_level" api:"required"`
	// The plan phase in which this adjustment is active.
	PlanPhaseOrder int64 `json:"plan_phase_order" api:"required,nullable"`
	// The reason for the adjustment.
	Reason string `json:"reason" api:"required,nullable"`
	// The adjustment id this adjustment replaces. This adjustment will take the place
	// of the replaced adjustment in plan version migrations.
	ReplacesAdjustmentID string `json:"replaces_adjustment_id" api:"required,nullable"`
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

// Legacy field representing the parent plan if the current plan is a 'child plan',
// overriding prices from the parent.
//
// Deprecated: deprecated
type PlanBasePlan struct {
	ID string `json:"id" api:"required,nullable"`
	// An optional user-defined ID for this plan resource, used throughout the system
	// as an alias for this Plan. Use this field to identify a plan by an existing
	// identifier in your system.
	ExternalPlanID string           `json:"external_plan_id" api:"required,nullable"`
	Name           string           `json:"name" api:"required,nullable"`
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

type PlanPlanPhase struct {
	ID          string          `json:"id" api:"required"`
	Description string          `json:"description" api:"required,nullable"`
	Discount    shared.Discount `json:"discount" api:"required,nullable"`
	// How many terms of length `duration_unit` this phase is active for. If null, this
	// phase is evergreen and active indefinitely
	Duration      int64                      `json:"duration" api:"required,nullable"`
	DurationUnit  PlanPlanPhasesDurationUnit `json:"duration_unit" api:"required,nullable"`
	Maximum       shared.Maximum             `json:"maximum" api:"required,nullable"`
	MaximumAmount string                     `json:"maximum_amount" api:"required,nullable"`
	Minimum       shared.Minimum             `json:"minimum" api:"required,nullable"`
	MinimumAmount string                     `json:"minimum_amount" api:"required,nullable"`
	Name          string                     `json:"name" api:"required"`
	// Determines the ordering of the phase in a plan's lifecycle. 1 = first phase.
	Order int64             `json:"order" api:"required"`
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
	ID        string          `json:"id" api:"required"`
	CreatedAt time.Time       `json:"created_at" api:"required" format:"date-time"`
	Name      string          `json:"name" api:"required"`
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
	TrialPeriod     int64                          `json:"trial_period" api:"required,nullable"`
	TrialPeriodUnit PlanTrialConfigTrialPeriodUnit `json:"trial_period_unit" api:"required"`
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

type PlanNewParams struct {
	// An ISO 4217 currency string for invoices generated by subscriptions on this
	// plan.
	Currency param.Field[string] `json:"currency" api:"required"`
	Name     param.Field[string] `json:"name" api:"required"`
	// Prices for this plan. If the plan has phases, this includes prices across all
	// phases of the plan.
	Prices param.Field[[]PlanNewParamsPrice] `json:"prices" api:"required"`
	// Adjustments for this plan. If the plan has phases, this includes adjustments
	// across all phases of the plan.
	Adjustments param.Field[[]PlanNewParamsAdjustment] `json:"adjustments"`
	// Free-form text which is available on the invoice PDF and the Orb invoice portal.
	DefaultInvoiceMemo param.Field[string] `json:"default_invoice_memo"`
	// An optional user-defined description of the plan.
	Description    param.Field[string] `json:"description"`
	ExternalPlanID param.Field[string] `json:"external_plan_id"`
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
	// The license allocation price to add to the plan.
	LicenseAllocationPrice param.Field[PlanNewParamsPricesLicenseAllocationPriceUnion] `json:"license_allocation_price"`
	// The phase to add this price to.
	PlanPhaseOrder param.Field[int64] `json:"plan_phase_order"`
	// New plan price request body params.
	Price param.Field[PlanNewParamsPricesPriceUnion] `json:"price"`
}

func (r PlanNewParamsPrice) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// The license allocation price to add to the plan.
type PlanNewParamsPricesLicenseAllocationPrice struct {
	// The cadence to bill for this price on.
	Cadence param.Field[PlanNewParamsPricesLicenseAllocationPriceCadence] `json:"cadence" api:"required"`
	// The id of the item the price will be associated with.
	ItemID             param.Field[string]      `json:"item_id" api:"required"`
	LicenseAllocations param.Field[interface{}] `json:"license_allocations" api:"required"`
	// The pricing model type
	ModelType param.Field[PlanNewParamsPricesLicenseAllocationPriceModelType] `json:"model_type" api:"required"`
	// The name of the price.
	Name param.Field[string] `json:"name" api:"required"`
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
	// The ID of the license type to associate with this price.
	LicenseTypeID param.Field[string] `json:"license_type_id"`
	// Configuration for matrix pricing
	MatrixConfig param.Field[shared.MatrixConfigParam] `json:"matrix_config"`
	// Configuration for matrix_with_allocation pricing
	MatrixWithAllocationConfig  param.Field[shared.MatrixWithAllocationConfigParam] `json:"matrix_with_allocation_config"`
	MatrixWithDisplayNameConfig param.Field[interface{}]                            `json:"matrix_with_display_name_config"`
	MaxGroupTieredPackageConfig param.Field[interface{}]                            `json:"max_group_tiered_package_config"`
	Metadata                    param.Field[interface{}]                            `json:"metadata"`
	MinimumCompositeConfig      param.Field[interface{}]                            `json:"minimum_composite_config"`
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

func (r PlanNewParamsPricesLicenseAllocationPrice) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r PlanNewParamsPricesLicenseAllocationPrice) implementsPlanNewParamsPricesLicenseAllocationPriceUnion() {
}

// The license allocation price to add to the plan.
//
// Satisfied by
// [PlanNewParamsPricesLicenseAllocationPriceNewLicenseAllocationUnitPrice],
// [PlanNewParamsPricesLicenseAllocationPriceNewLicenseAllocationTieredPrice],
// [PlanNewParamsPricesLicenseAllocationPriceNewLicenseAllocationBulkPrice],
// [PlanNewParamsPricesLicenseAllocationPriceNewLicenseAllocationBulkWithFiltersPrice],
// [PlanNewParamsPricesLicenseAllocationPriceNewLicenseAllocationPackagePrice],
// [PlanNewParamsPricesLicenseAllocationPriceNewLicenseAllocationMatrixPrice],
// [PlanNewParamsPricesLicenseAllocationPriceNewLicenseAllocationThresholdTotalAmountPrice],
// [PlanNewParamsPricesLicenseAllocationPriceNewLicenseAllocationTieredPackagePrice],
// [PlanNewParamsPricesLicenseAllocationPriceNewLicenseAllocationTieredWithMinimumPrice],
// [PlanNewParamsPricesLicenseAllocationPriceNewLicenseAllocationGroupedTieredPrice],
// [PlanNewParamsPricesLicenseAllocationPriceNewLicenseAllocationTieredPackageWithMinimumPrice],
// [PlanNewParamsPricesLicenseAllocationPriceNewLicenseAllocationPackageWithAllocationPrice],
// [PlanNewParamsPricesLicenseAllocationPriceNewLicenseAllocationUnitWithPercentPrice],
// [PlanNewParamsPricesLicenseAllocationPriceNewLicenseAllocationMatrixWithAllocationPrice],
// [PlanNewParamsPricesLicenseAllocationPriceNewLicenseAllocationTieredWithProrationPrice],
// [PlanNewParamsPricesLicenseAllocationPriceNewLicenseAllocationUnitWithProrationPrice],
// [PlanNewParamsPricesLicenseAllocationPriceNewLicenseAllocationGroupedAllocationPrice],
// [PlanNewParamsPricesLicenseAllocationPriceNewLicenseAllocationBulkWithProrationPrice],
// [PlanNewParamsPricesLicenseAllocationPriceNewLicenseAllocationGroupedWithProratedMinimumPrice],
// [PlanNewParamsPricesLicenseAllocationPriceNewLicenseAllocationGroupedWithMeteredMinimumPrice],
// [PlanNewParamsPricesLicenseAllocationPriceNewLicenseAllocationGroupedWithMinMaxThresholdsPrice],
// [PlanNewParamsPricesLicenseAllocationPriceNewLicenseAllocationMatrixWithDisplayNamePrice],
// [PlanNewParamsPricesLicenseAllocationPriceNewLicenseAllocationGroupedTieredPackagePrice],
// [PlanNewParamsPricesLicenseAllocationPriceNewLicenseAllocationMaxGroupTieredPackagePrice],
// [PlanNewParamsPricesLicenseAllocationPriceNewLicenseAllocationScalableMatrixWithUnitPricingPrice],
// [PlanNewParamsPricesLicenseAllocationPriceNewLicenseAllocationScalableMatrixWithTieredPricingPrice],
// [PlanNewParamsPricesLicenseAllocationPriceNewLicenseAllocationCumulativeGroupedBulkPrice],
// [PlanNewParamsPricesLicenseAllocationPriceNewLicenseAllocationCumulativeGroupedAllocationPrice],
// [PlanNewParamsPricesLicenseAllocationPriceNewLicenseAllocationMinimumCompositePrice],
// [PlanNewParamsPricesLicenseAllocationPriceNewLicenseAllocationPercentCompositePrice],
// [PlanNewParamsPricesLicenseAllocationPriceNewLicenseAllocationEventOutputPrice],
// [PlanNewParamsPricesLicenseAllocationPrice].
type PlanNewParamsPricesLicenseAllocationPriceUnion interface {
	implementsPlanNewParamsPricesLicenseAllocationPriceUnion()
}

type PlanNewParamsPricesLicenseAllocationPriceNewLicenseAllocationUnitPrice struct {
	// The cadence to bill for this price on.
	Cadence param.Field[PlanNewParamsPricesLicenseAllocationPriceNewLicenseAllocationUnitPriceCadence] `json:"cadence" api:"required"`
	// The id of the item the price will be associated with.
	ItemID param.Field[string] `json:"item_id" api:"required"`
	// License allocations to associate with this price. Each entry defines a
	// per-license credit pool granted each cadence. Requires license_type_id or
	// license_type_configuration to be set.
	LicenseAllocations param.Field[[]PlanNewParamsPricesLicenseAllocationPriceNewLicenseAllocationUnitPriceLicenseAllocation] `json:"license_allocations" api:"required"`
	// The pricing model type
	ModelType param.Field[PlanNewParamsPricesLicenseAllocationPriceNewLicenseAllocationUnitPriceModelType] `json:"model_type" api:"required"`
	// The name of the price.
	Name param.Field[string] `json:"name" api:"required"`
	// Configuration for unit pricing
	UnitConfig param.Field[shared.UnitConfigParam] `json:"unit_config" api:"required"`
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
	ConversionRateConfig param.Field[PlanNewParamsPricesLicenseAllocationPriceNewLicenseAllocationUnitPriceConversionRateConfigUnion] `json:"conversion_rate_config"`
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
	// The ID of the license type to associate with this price.
	LicenseTypeID param.Field[string] `json:"license_type_id"`
	// User-specified key/value pairs for the resource. Individual keys can be removed
	// by setting the value to `null`, and the entire metadata mapping can be cleared
	// by setting `metadata` to `null`.
	Metadata param.Field[map[string]string] `json:"metadata"`
	// A transient ID that can be used to reference this price when adding adjustments
	// in the same API call.
	ReferenceID param.Field[string] `json:"reference_id"`
}

func (r PlanNewParamsPricesLicenseAllocationPriceNewLicenseAllocationUnitPrice) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r PlanNewParamsPricesLicenseAllocationPriceNewLicenseAllocationUnitPrice) implementsPlanNewParamsPricesLicenseAllocationPriceUnion() {
}

// The cadence to bill for this price on.
type PlanNewParamsPricesLicenseAllocationPriceNewLicenseAllocationUnitPriceCadence string

const (
	PlanNewParamsPricesLicenseAllocationPriceNewLicenseAllocationUnitPriceCadenceAnnual     PlanNewParamsPricesLicenseAllocationPriceNewLicenseAllocationUnitPriceCadence = "annual"
	PlanNewParamsPricesLicenseAllocationPriceNewLicenseAllocationUnitPriceCadenceSemiAnnual PlanNewParamsPricesLicenseAllocationPriceNewLicenseAllocationUnitPriceCadence = "semi_annual"
	PlanNewParamsPricesLicenseAllocationPriceNewLicenseAllocationUnitPriceCadenceMonthly    PlanNewParamsPricesLicenseAllocationPriceNewLicenseAllocationUnitPriceCadence = "monthly"
	PlanNewParamsPricesLicenseAllocationPriceNewLicenseAllocationUnitPriceCadenceQuarterly  PlanNewParamsPricesLicenseAllocationPriceNewLicenseAllocationUnitPriceCadence = "quarterly"
	PlanNewParamsPricesLicenseAllocationPriceNewLicenseAllocationUnitPriceCadenceOneTime    PlanNewParamsPricesLicenseAllocationPriceNewLicenseAllocationUnitPriceCadence = "one_time"
	PlanNewParamsPricesLicenseAllocationPriceNewLicenseAllocationUnitPriceCadenceCustom     PlanNewParamsPricesLicenseAllocationPriceNewLicenseAllocationUnitPriceCadence = "custom"
)

func (r PlanNewParamsPricesLicenseAllocationPriceNewLicenseAllocationUnitPriceCadence) IsKnown() bool {
	switch r {
	case PlanNewParamsPricesLicenseAllocationPriceNewLicenseAllocationUnitPriceCadenceAnnual, PlanNewParamsPricesLicenseAllocationPriceNewLicenseAllocationUnitPriceCadenceSemiAnnual, PlanNewParamsPricesLicenseAllocationPriceNewLicenseAllocationUnitPriceCadenceMonthly, PlanNewParamsPricesLicenseAllocationPriceNewLicenseAllocationUnitPriceCadenceQuarterly, PlanNewParamsPricesLicenseAllocationPriceNewLicenseAllocationUnitPriceCadenceOneTime, PlanNewParamsPricesLicenseAllocationPriceNewLicenseAllocationUnitPriceCadenceCustom:
		return true
	}
	return false
}

type PlanNewParamsPricesLicenseAllocationPriceNewLicenseAllocationUnitPriceLicenseAllocation struct {
	// The amount of credits granted per active license per cadence.
	Amount param.Field[string] `json:"amount" api:"required"`
	// The currency of the license allocation.
	Currency param.Field[string] `json:"currency" api:"required"`
	// When True, overage beyond the allocation is written off.
	WriteOffOverage param.Field[bool] `json:"write_off_overage"`
}

func (r PlanNewParamsPricesLicenseAllocationPriceNewLicenseAllocationUnitPriceLicenseAllocation) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// The pricing model type
type PlanNewParamsPricesLicenseAllocationPriceNewLicenseAllocationUnitPriceModelType string

const (
	PlanNewParamsPricesLicenseAllocationPriceNewLicenseAllocationUnitPriceModelTypeUnit PlanNewParamsPricesLicenseAllocationPriceNewLicenseAllocationUnitPriceModelType = "unit"
)

func (r PlanNewParamsPricesLicenseAllocationPriceNewLicenseAllocationUnitPriceModelType) IsKnown() bool {
	switch r {
	case PlanNewParamsPricesLicenseAllocationPriceNewLicenseAllocationUnitPriceModelTypeUnit:
		return true
	}
	return false
}

type PlanNewParamsPricesLicenseAllocationPriceNewLicenseAllocationUnitPriceConversionRateConfig struct {
	ConversionRateType param.Field[PlanNewParamsPricesLicenseAllocationPriceNewLicenseAllocationUnitPriceConversionRateConfigConversionRateType] `json:"conversion_rate_type" api:"required"`
	TieredConfig       param.Field[shared.ConversionRateTieredConfigParam]                                                                       `json:"tiered_config"`
	UnitConfig         param.Field[shared.ConversionRateUnitConfigParam]                                                                         `json:"unit_config"`
}

func (r PlanNewParamsPricesLicenseAllocationPriceNewLicenseAllocationUnitPriceConversionRateConfig) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r PlanNewParamsPricesLicenseAllocationPriceNewLicenseAllocationUnitPriceConversionRateConfig) ImplementsPlanNewParamsPricesLicenseAllocationPriceNewLicenseAllocationUnitPriceConversionRateConfigUnion() {
}

// Satisfied by [shared.UnitConversionRateConfigParam],
// [shared.TieredConversionRateConfigParam],
// [PlanNewParamsPricesLicenseAllocationPriceNewLicenseAllocationUnitPriceConversionRateConfig].
type PlanNewParamsPricesLicenseAllocationPriceNewLicenseAllocationUnitPriceConversionRateConfigUnion interface {
	ImplementsPlanNewParamsPricesLicenseAllocationPriceNewLicenseAllocationUnitPriceConversionRateConfigUnion()
}

type PlanNewParamsPricesLicenseAllocationPriceNewLicenseAllocationUnitPriceConversionRateConfigConversionRateType string

const (
	PlanNewParamsPricesLicenseAllocationPriceNewLicenseAllocationUnitPriceConversionRateConfigConversionRateTypeUnit   PlanNewParamsPricesLicenseAllocationPriceNewLicenseAllocationUnitPriceConversionRateConfigConversionRateType = "unit"
	PlanNewParamsPricesLicenseAllocationPriceNewLicenseAllocationUnitPriceConversionRateConfigConversionRateTypeTiered PlanNewParamsPricesLicenseAllocationPriceNewLicenseAllocationUnitPriceConversionRateConfigConversionRateType = "tiered"
)

func (r PlanNewParamsPricesLicenseAllocationPriceNewLicenseAllocationUnitPriceConversionRateConfigConversionRateType) IsKnown() bool {
	switch r {
	case PlanNewParamsPricesLicenseAllocationPriceNewLicenseAllocationUnitPriceConversionRateConfigConversionRateTypeUnit, PlanNewParamsPricesLicenseAllocationPriceNewLicenseAllocationUnitPriceConversionRateConfigConversionRateTypeTiered:
		return true
	}
	return false
}

type PlanNewParamsPricesLicenseAllocationPriceNewLicenseAllocationTieredPrice struct {
	// The cadence to bill for this price on.
	Cadence param.Field[PlanNewParamsPricesLicenseAllocationPriceNewLicenseAllocationTieredPriceCadence] `json:"cadence" api:"required"`
	// The id of the item the price will be associated with.
	ItemID param.Field[string] `json:"item_id" api:"required"`
	// License allocations to associate with this price. Each entry defines a
	// per-license credit pool granted each cadence. Requires license_type_id or
	// license_type_configuration to be set.
	LicenseAllocations param.Field[[]PlanNewParamsPricesLicenseAllocationPriceNewLicenseAllocationTieredPriceLicenseAllocation] `json:"license_allocations" api:"required"`
	// The pricing model type
	ModelType param.Field[PlanNewParamsPricesLicenseAllocationPriceNewLicenseAllocationTieredPriceModelType] `json:"model_type" api:"required"`
	// The name of the price.
	Name param.Field[string] `json:"name" api:"required"`
	// Configuration for tiered pricing
	TieredConfig param.Field[shared.TieredConfigParam] `json:"tiered_config" api:"required"`
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
	ConversionRateConfig param.Field[PlanNewParamsPricesLicenseAllocationPriceNewLicenseAllocationTieredPriceConversionRateConfigUnion] `json:"conversion_rate_config"`
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
	// The ID of the license type to associate with this price.
	LicenseTypeID param.Field[string] `json:"license_type_id"`
	// User-specified key/value pairs for the resource. Individual keys can be removed
	// by setting the value to `null`, and the entire metadata mapping can be cleared
	// by setting `metadata` to `null`.
	Metadata param.Field[map[string]string] `json:"metadata"`
	// A transient ID that can be used to reference this price when adding adjustments
	// in the same API call.
	ReferenceID param.Field[string] `json:"reference_id"`
}

func (r PlanNewParamsPricesLicenseAllocationPriceNewLicenseAllocationTieredPrice) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r PlanNewParamsPricesLicenseAllocationPriceNewLicenseAllocationTieredPrice) implementsPlanNewParamsPricesLicenseAllocationPriceUnion() {
}

// The cadence to bill for this price on.
type PlanNewParamsPricesLicenseAllocationPriceNewLicenseAllocationTieredPriceCadence string

const (
	PlanNewParamsPricesLicenseAllocationPriceNewLicenseAllocationTieredPriceCadenceAnnual     PlanNewParamsPricesLicenseAllocationPriceNewLicenseAllocationTieredPriceCadence = "annual"
	PlanNewParamsPricesLicenseAllocationPriceNewLicenseAllocationTieredPriceCadenceSemiAnnual PlanNewParamsPricesLicenseAllocationPriceNewLicenseAllocationTieredPriceCadence = "semi_annual"
	PlanNewParamsPricesLicenseAllocationPriceNewLicenseAllocationTieredPriceCadenceMonthly    PlanNewParamsPricesLicenseAllocationPriceNewLicenseAllocationTieredPriceCadence = "monthly"
	PlanNewParamsPricesLicenseAllocationPriceNewLicenseAllocationTieredPriceCadenceQuarterly  PlanNewParamsPricesLicenseAllocationPriceNewLicenseAllocationTieredPriceCadence = "quarterly"
	PlanNewParamsPricesLicenseAllocationPriceNewLicenseAllocationTieredPriceCadenceOneTime    PlanNewParamsPricesLicenseAllocationPriceNewLicenseAllocationTieredPriceCadence = "one_time"
	PlanNewParamsPricesLicenseAllocationPriceNewLicenseAllocationTieredPriceCadenceCustom     PlanNewParamsPricesLicenseAllocationPriceNewLicenseAllocationTieredPriceCadence = "custom"
)

func (r PlanNewParamsPricesLicenseAllocationPriceNewLicenseAllocationTieredPriceCadence) IsKnown() bool {
	switch r {
	case PlanNewParamsPricesLicenseAllocationPriceNewLicenseAllocationTieredPriceCadenceAnnual, PlanNewParamsPricesLicenseAllocationPriceNewLicenseAllocationTieredPriceCadenceSemiAnnual, PlanNewParamsPricesLicenseAllocationPriceNewLicenseAllocationTieredPriceCadenceMonthly, PlanNewParamsPricesLicenseAllocationPriceNewLicenseAllocationTieredPriceCadenceQuarterly, PlanNewParamsPricesLicenseAllocationPriceNewLicenseAllocationTieredPriceCadenceOneTime, PlanNewParamsPricesLicenseAllocationPriceNewLicenseAllocationTieredPriceCadenceCustom:
		return true
	}
	return false
}

type PlanNewParamsPricesLicenseAllocationPriceNewLicenseAllocationTieredPriceLicenseAllocation struct {
	// The amount of credits granted per active license per cadence.
	Amount param.Field[string] `json:"amount" api:"required"`
	// The currency of the license allocation.
	Currency param.Field[string] `json:"currency" api:"required"`
	// When True, overage beyond the allocation is written off.
	WriteOffOverage param.Field[bool] `json:"write_off_overage"`
}

func (r PlanNewParamsPricesLicenseAllocationPriceNewLicenseAllocationTieredPriceLicenseAllocation) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// The pricing model type
type PlanNewParamsPricesLicenseAllocationPriceNewLicenseAllocationTieredPriceModelType string

const (
	PlanNewParamsPricesLicenseAllocationPriceNewLicenseAllocationTieredPriceModelTypeTiered PlanNewParamsPricesLicenseAllocationPriceNewLicenseAllocationTieredPriceModelType = "tiered"
)

func (r PlanNewParamsPricesLicenseAllocationPriceNewLicenseAllocationTieredPriceModelType) IsKnown() bool {
	switch r {
	case PlanNewParamsPricesLicenseAllocationPriceNewLicenseAllocationTieredPriceModelTypeTiered:
		return true
	}
	return false
}

type PlanNewParamsPricesLicenseAllocationPriceNewLicenseAllocationTieredPriceConversionRateConfig struct {
	ConversionRateType param.Field[PlanNewParamsPricesLicenseAllocationPriceNewLicenseAllocationTieredPriceConversionRateConfigConversionRateType] `json:"conversion_rate_type" api:"required"`
	TieredConfig       param.Field[shared.ConversionRateTieredConfigParam]                                                                         `json:"tiered_config"`
	UnitConfig         param.Field[shared.ConversionRateUnitConfigParam]                                                                           `json:"unit_config"`
}

func (r PlanNewParamsPricesLicenseAllocationPriceNewLicenseAllocationTieredPriceConversionRateConfig) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r PlanNewParamsPricesLicenseAllocationPriceNewLicenseAllocationTieredPriceConversionRateConfig) ImplementsPlanNewParamsPricesLicenseAllocationPriceNewLicenseAllocationTieredPriceConversionRateConfigUnion() {
}

// Satisfied by [shared.UnitConversionRateConfigParam],
// [shared.TieredConversionRateConfigParam],
// [PlanNewParamsPricesLicenseAllocationPriceNewLicenseAllocationTieredPriceConversionRateConfig].
type PlanNewParamsPricesLicenseAllocationPriceNewLicenseAllocationTieredPriceConversionRateConfigUnion interface {
	ImplementsPlanNewParamsPricesLicenseAllocationPriceNewLicenseAllocationTieredPriceConversionRateConfigUnion()
}

type PlanNewParamsPricesLicenseAllocationPriceNewLicenseAllocationTieredPriceConversionRateConfigConversionRateType string

const (
	PlanNewParamsPricesLicenseAllocationPriceNewLicenseAllocationTieredPriceConversionRateConfigConversionRateTypeUnit   PlanNewParamsPricesLicenseAllocationPriceNewLicenseAllocationTieredPriceConversionRateConfigConversionRateType = "unit"
	PlanNewParamsPricesLicenseAllocationPriceNewLicenseAllocationTieredPriceConversionRateConfigConversionRateTypeTiered PlanNewParamsPricesLicenseAllocationPriceNewLicenseAllocationTieredPriceConversionRateConfigConversionRateType = "tiered"
)

func (r PlanNewParamsPricesLicenseAllocationPriceNewLicenseAllocationTieredPriceConversionRateConfigConversionRateType) IsKnown() bool {
	switch r {
	case PlanNewParamsPricesLicenseAllocationPriceNewLicenseAllocationTieredPriceConversionRateConfigConversionRateTypeUnit, PlanNewParamsPricesLicenseAllocationPriceNewLicenseAllocationTieredPriceConversionRateConfigConversionRateTypeTiered:
		return true
	}
	return false
}

type PlanNewParamsPricesLicenseAllocationPriceNewLicenseAllocationBulkPrice struct {
	// Configuration for bulk pricing
	BulkConfig param.Field[shared.BulkConfigParam] `json:"bulk_config" api:"required"`
	// The cadence to bill for this price on.
	Cadence param.Field[PlanNewParamsPricesLicenseAllocationPriceNewLicenseAllocationBulkPriceCadence] `json:"cadence" api:"required"`
	// The id of the item the price will be associated with.
	ItemID param.Field[string] `json:"item_id" api:"required"`
	// License allocations to associate with this price. Each entry defines a
	// per-license credit pool granted each cadence. Requires license_type_id or
	// license_type_configuration to be set.
	LicenseAllocations param.Field[[]PlanNewParamsPricesLicenseAllocationPriceNewLicenseAllocationBulkPriceLicenseAllocation] `json:"license_allocations" api:"required"`
	// The pricing model type
	ModelType param.Field[PlanNewParamsPricesLicenseAllocationPriceNewLicenseAllocationBulkPriceModelType] `json:"model_type" api:"required"`
	// The name of the price.
	Name param.Field[string] `json:"name" api:"required"`
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
	ConversionRateConfig param.Field[PlanNewParamsPricesLicenseAllocationPriceNewLicenseAllocationBulkPriceConversionRateConfigUnion] `json:"conversion_rate_config"`
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
	// The ID of the license type to associate with this price.
	LicenseTypeID param.Field[string] `json:"license_type_id"`
	// User-specified key/value pairs for the resource. Individual keys can be removed
	// by setting the value to `null`, and the entire metadata mapping can be cleared
	// by setting `metadata` to `null`.
	Metadata param.Field[map[string]string] `json:"metadata"`
	// A transient ID that can be used to reference this price when adding adjustments
	// in the same API call.
	ReferenceID param.Field[string] `json:"reference_id"`
}

func (r PlanNewParamsPricesLicenseAllocationPriceNewLicenseAllocationBulkPrice) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r PlanNewParamsPricesLicenseAllocationPriceNewLicenseAllocationBulkPrice) implementsPlanNewParamsPricesLicenseAllocationPriceUnion() {
}

// The cadence to bill for this price on.
type PlanNewParamsPricesLicenseAllocationPriceNewLicenseAllocationBulkPriceCadence string

const (
	PlanNewParamsPricesLicenseAllocationPriceNewLicenseAllocationBulkPriceCadenceAnnual     PlanNewParamsPricesLicenseAllocationPriceNewLicenseAllocationBulkPriceCadence = "annual"
	PlanNewParamsPricesLicenseAllocationPriceNewLicenseAllocationBulkPriceCadenceSemiAnnual PlanNewParamsPricesLicenseAllocationPriceNewLicenseAllocationBulkPriceCadence = "semi_annual"
	PlanNewParamsPricesLicenseAllocationPriceNewLicenseAllocationBulkPriceCadenceMonthly    PlanNewParamsPricesLicenseAllocationPriceNewLicenseAllocationBulkPriceCadence = "monthly"
	PlanNewParamsPricesLicenseAllocationPriceNewLicenseAllocationBulkPriceCadenceQuarterly  PlanNewParamsPricesLicenseAllocationPriceNewLicenseAllocationBulkPriceCadence = "quarterly"
	PlanNewParamsPricesLicenseAllocationPriceNewLicenseAllocationBulkPriceCadenceOneTime    PlanNewParamsPricesLicenseAllocationPriceNewLicenseAllocationBulkPriceCadence = "one_time"
	PlanNewParamsPricesLicenseAllocationPriceNewLicenseAllocationBulkPriceCadenceCustom     PlanNewParamsPricesLicenseAllocationPriceNewLicenseAllocationBulkPriceCadence = "custom"
)

func (r PlanNewParamsPricesLicenseAllocationPriceNewLicenseAllocationBulkPriceCadence) IsKnown() bool {
	switch r {
	case PlanNewParamsPricesLicenseAllocationPriceNewLicenseAllocationBulkPriceCadenceAnnual, PlanNewParamsPricesLicenseAllocationPriceNewLicenseAllocationBulkPriceCadenceSemiAnnual, PlanNewParamsPricesLicenseAllocationPriceNewLicenseAllocationBulkPriceCadenceMonthly, PlanNewParamsPricesLicenseAllocationPriceNewLicenseAllocationBulkPriceCadenceQuarterly, PlanNewParamsPricesLicenseAllocationPriceNewLicenseAllocationBulkPriceCadenceOneTime, PlanNewParamsPricesLicenseAllocationPriceNewLicenseAllocationBulkPriceCadenceCustom:
		return true
	}
	return false
}

type PlanNewParamsPricesLicenseAllocationPriceNewLicenseAllocationBulkPriceLicenseAllocation struct {
	// The amount of credits granted per active license per cadence.
	Amount param.Field[string] `json:"amount" api:"required"`
	// The currency of the license allocation.
	Currency param.Field[string] `json:"currency" api:"required"`
	// When True, overage beyond the allocation is written off.
	WriteOffOverage param.Field[bool] `json:"write_off_overage"`
}

func (r PlanNewParamsPricesLicenseAllocationPriceNewLicenseAllocationBulkPriceLicenseAllocation) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// The pricing model type
type PlanNewParamsPricesLicenseAllocationPriceNewLicenseAllocationBulkPriceModelType string

const (
	PlanNewParamsPricesLicenseAllocationPriceNewLicenseAllocationBulkPriceModelTypeBulk PlanNewParamsPricesLicenseAllocationPriceNewLicenseAllocationBulkPriceModelType = "bulk"
)

func (r PlanNewParamsPricesLicenseAllocationPriceNewLicenseAllocationBulkPriceModelType) IsKnown() bool {
	switch r {
	case PlanNewParamsPricesLicenseAllocationPriceNewLicenseAllocationBulkPriceModelTypeBulk:
		return true
	}
	return false
}

type PlanNewParamsPricesLicenseAllocationPriceNewLicenseAllocationBulkPriceConversionRateConfig struct {
	ConversionRateType param.Field[PlanNewParamsPricesLicenseAllocationPriceNewLicenseAllocationBulkPriceConversionRateConfigConversionRateType] `json:"conversion_rate_type" api:"required"`
	TieredConfig       param.Field[shared.ConversionRateTieredConfigParam]                                                                       `json:"tiered_config"`
	UnitConfig         param.Field[shared.ConversionRateUnitConfigParam]                                                                         `json:"unit_config"`
}

func (r PlanNewParamsPricesLicenseAllocationPriceNewLicenseAllocationBulkPriceConversionRateConfig) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r PlanNewParamsPricesLicenseAllocationPriceNewLicenseAllocationBulkPriceConversionRateConfig) ImplementsPlanNewParamsPricesLicenseAllocationPriceNewLicenseAllocationBulkPriceConversionRateConfigUnion() {
}

// Satisfied by [shared.UnitConversionRateConfigParam],
// [shared.TieredConversionRateConfigParam],
// [PlanNewParamsPricesLicenseAllocationPriceNewLicenseAllocationBulkPriceConversionRateConfig].
type PlanNewParamsPricesLicenseAllocationPriceNewLicenseAllocationBulkPriceConversionRateConfigUnion interface {
	ImplementsPlanNewParamsPricesLicenseAllocationPriceNewLicenseAllocationBulkPriceConversionRateConfigUnion()
}

type PlanNewParamsPricesLicenseAllocationPriceNewLicenseAllocationBulkPriceConversionRateConfigConversionRateType string

const (
	PlanNewParamsPricesLicenseAllocationPriceNewLicenseAllocationBulkPriceConversionRateConfigConversionRateTypeUnit   PlanNewParamsPricesLicenseAllocationPriceNewLicenseAllocationBulkPriceConversionRateConfigConversionRateType = "unit"
	PlanNewParamsPricesLicenseAllocationPriceNewLicenseAllocationBulkPriceConversionRateConfigConversionRateTypeTiered PlanNewParamsPricesLicenseAllocationPriceNewLicenseAllocationBulkPriceConversionRateConfigConversionRateType = "tiered"
)

func (r PlanNewParamsPricesLicenseAllocationPriceNewLicenseAllocationBulkPriceConversionRateConfigConversionRateType) IsKnown() bool {
	switch r {
	case PlanNewParamsPricesLicenseAllocationPriceNewLicenseAllocationBulkPriceConversionRateConfigConversionRateTypeUnit, PlanNewParamsPricesLicenseAllocationPriceNewLicenseAllocationBulkPriceConversionRateConfigConversionRateTypeTiered:
		return true
	}
	return false
}

type PlanNewParamsPricesLicenseAllocationPriceNewLicenseAllocationBulkWithFiltersPrice struct {
	// Configuration for bulk_with_filters pricing
	BulkWithFiltersConfig param.Field[PlanNewParamsPricesLicenseAllocationPriceNewLicenseAllocationBulkWithFiltersPriceBulkWithFiltersConfig] `json:"bulk_with_filters_config" api:"required"`
	// The cadence to bill for this price on.
	Cadence param.Field[PlanNewParamsPricesLicenseAllocationPriceNewLicenseAllocationBulkWithFiltersPriceCadence] `json:"cadence" api:"required"`
	// The id of the item the price will be associated with.
	ItemID param.Field[string] `json:"item_id" api:"required"`
	// License allocations to associate with this price. Each entry defines a
	// per-license credit pool granted each cadence. Requires license_type_id or
	// license_type_configuration to be set.
	LicenseAllocations param.Field[[]PlanNewParamsPricesLicenseAllocationPriceNewLicenseAllocationBulkWithFiltersPriceLicenseAllocation] `json:"license_allocations" api:"required"`
	// The pricing model type
	ModelType param.Field[PlanNewParamsPricesLicenseAllocationPriceNewLicenseAllocationBulkWithFiltersPriceModelType] `json:"model_type" api:"required"`
	// The name of the price.
	Name param.Field[string] `json:"name" api:"required"`
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
	ConversionRateConfig param.Field[PlanNewParamsPricesLicenseAllocationPriceNewLicenseAllocationBulkWithFiltersPriceConversionRateConfigUnion] `json:"conversion_rate_config"`
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
	// The ID of the license type to associate with this price.
	LicenseTypeID param.Field[string] `json:"license_type_id"`
	// User-specified key/value pairs for the resource. Individual keys can be removed
	// by setting the value to `null`, and the entire metadata mapping can be cleared
	// by setting `metadata` to `null`.
	Metadata param.Field[map[string]string] `json:"metadata"`
	// A transient ID that can be used to reference this price when adding adjustments
	// in the same API call.
	ReferenceID param.Field[string] `json:"reference_id"`
}

func (r PlanNewParamsPricesLicenseAllocationPriceNewLicenseAllocationBulkWithFiltersPrice) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r PlanNewParamsPricesLicenseAllocationPriceNewLicenseAllocationBulkWithFiltersPrice) implementsPlanNewParamsPricesLicenseAllocationPriceUnion() {
}

// Configuration for bulk_with_filters pricing
type PlanNewParamsPricesLicenseAllocationPriceNewLicenseAllocationBulkWithFiltersPriceBulkWithFiltersConfig struct {
	// Property filters to apply (all must match)
	Filters param.Field[[]PlanNewParamsPricesLicenseAllocationPriceNewLicenseAllocationBulkWithFiltersPriceBulkWithFiltersConfigFilter] `json:"filters" api:"required"`
	// Bulk tiers for rating based on total usage volume
	Tiers param.Field[[]PlanNewParamsPricesLicenseAllocationPriceNewLicenseAllocationBulkWithFiltersPriceBulkWithFiltersConfigTier] `json:"tiers" api:"required"`
}

func (r PlanNewParamsPricesLicenseAllocationPriceNewLicenseAllocationBulkWithFiltersPriceBulkWithFiltersConfig) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Configuration for a single property filter
type PlanNewParamsPricesLicenseAllocationPriceNewLicenseAllocationBulkWithFiltersPriceBulkWithFiltersConfigFilter struct {
	// Event property key to filter on
	PropertyKey param.Field[string] `json:"property_key" api:"required"`
	// Event property value to match
	PropertyValue param.Field[string] `json:"property_value" api:"required"`
}

func (r PlanNewParamsPricesLicenseAllocationPriceNewLicenseAllocationBulkWithFiltersPriceBulkWithFiltersConfigFilter) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Configuration for a single bulk pricing tier
type PlanNewParamsPricesLicenseAllocationPriceNewLicenseAllocationBulkWithFiltersPriceBulkWithFiltersConfigTier struct {
	// Amount per unit
	UnitAmount param.Field[string] `json:"unit_amount" api:"required"`
	// The lower bound for this tier
	TierLowerBound param.Field[string] `json:"tier_lower_bound"`
}

func (r PlanNewParamsPricesLicenseAllocationPriceNewLicenseAllocationBulkWithFiltersPriceBulkWithFiltersConfigTier) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// The cadence to bill for this price on.
type PlanNewParamsPricesLicenseAllocationPriceNewLicenseAllocationBulkWithFiltersPriceCadence string

const (
	PlanNewParamsPricesLicenseAllocationPriceNewLicenseAllocationBulkWithFiltersPriceCadenceAnnual     PlanNewParamsPricesLicenseAllocationPriceNewLicenseAllocationBulkWithFiltersPriceCadence = "annual"
	PlanNewParamsPricesLicenseAllocationPriceNewLicenseAllocationBulkWithFiltersPriceCadenceSemiAnnual PlanNewParamsPricesLicenseAllocationPriceNewLicenseAllocationBulkWithFiltersPriceCadence = "semi_annual"
	PlanNewParamsPricesLicenseAllocationPriceNewLicenseAllocationBulkWithFiltersPriceCadenceMonthly    PlanNewParamsPricesLicenseAllocationPriceNewLicenseAllocationBulkWithFiltersPriceCadence = "monthly"
	PlanNewParamsPricesLicenseAllocationPriceNewLicenseAllocationBulkWithFiltersPriceCadenceQuarterly  PlanNewParamsPricesLicenseAllocationPriceNewLicenseAllocationBulkWithFiltersPriceCadence = "quarterly"
	PlanNewParamsPricesLicenseAllocationPriceNewLicenseAllocationBulkWithFiltersPriceCadenceOneTime    PlanNewParamsPricesLicenseAllocationPriceNewLicenseAllocationBulkWithFiltersPriceCadence = "one_time"
	PlanNewParamsPricesLicenseAllocationPriceNewLicenseAllocationBulkWithFiltersPriceCadenceCustom     PlanNewParamsPricesLicenseAllocationPriceNewLicenseAllocationBulkWithFiltersPriceCadence = "custom"
)

func (r PlanNewParamsPricesLicenseAllocationPriceNewLicenseAllocationBulkWithFiltersPriceCadence) IsKnown() bool {
	switch r {
	case PlanNewParamsPricesLicenseAllocationPriceNewLicenseAllocationBulkWithFiltersPriceCadenceAnnual, PlanNewParamsPricesLicenseAllocationPriceNewLicenseAllocationBulkWithFiltersPriceCadenceSemiAnnual, PlanNewParamsPricesLicenseAllocationPriceNewLicenseAllocationBulkWithFiltersPriceCadenceMonthly, PlanNewParamsPricesLicenseAllocationPriceNewLicenseAllocationBulkWithFiltersPriceCadenceQuarterly, PlanNewParamsPricesLicenseAllocationPriceNewLicenseAllocationBulkWithFiltersPriceCadenceOneTime, PlanNewParamsPricesLicenseAllocationPriceNewLicenseAllocationBulkWithFiltersPriceCadenceCustom:
		return true
	}
	return false
}

type PlanNewParamsPricesLicenseAllocationPriceNewLicenseAllocationBulkWithFiltersPriceLicenseAllocation struct {
	// The amount of credits granted per active license per cadence.
	Amount param.Field[string] `json:"amount" api:"required"`
	// The currency of the license allocation.
	Currency param.Field[string] `json:"currency" api:"required"`
	// When True, overage beyond the allocation is written off.
	WriteOffOverage param.Field[bool] `json:"write_off_overage"`
}

func (r PlanNewParamsPricesLicenseAllocationPriceNewLicenseAllocationBulkWithFiltersPriceLicenseAllocation) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// The pricing model type
type PlanNewParamsPricesLicenseAllocationPriceNewLicenseAllocationBulkWithFiltersPriceModelType string

const (
	PlanNewParamsPricesLicenseAllocationPriceNewLicenseAllocationBulkWithFiltersPriceModelTypeBulkWithFilters PlanNewParamsPricesLicenseAllocationPriceNewLicenseAllocationBulkWithFiltersPriceModelType = "bulk_with_filters"
)

func (r PlanNewParamsPricesLicenseAllocationPriceNewLicenseAllocationBulkWithFiltersPriceModelType) IsKnown() bool {
	switch r {
	case PlanNewParamsPricesLicenseAllocationPriceNewLicenseAllocationBulkWithFiltersPriceModelTypeBulkWithFilters:
		return true
	}
	return false
}

type PlanNewParamsPricesLicenseAllocationPriceNewLicenseAllocationBulkWithFiltersPriceConversionRateConfig struct {
	ConversionRateType param.Field[PlanNewParamsPricesLicenseAllocationPriceNewLicenseAllocationBulkWithFiltersPriceConversionRateConfigConversionRateType] `json:"conversion_rate_type" api:"required"`
	TieredConfig       param.Field[shared.ConversionRateTieredConfigParam]                                                                                  `json:"tiered_config"`
	UnitConfig         param.Field[shared.ConversionRateUnitConfigParam]                                                                                    `json:"unit_config"`
}

func (r PlanNewParamsPricesLicenseAllocationPriceNewLicenseAllocationBulkWithFiltersPriceConversionRateConfig) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r PlanNewParamsPricesLicenseAllocationPriceNewLicenseAllocationBulkWithFiltersPriceConversionRateConfig) ImplementsPlanNewParamsPricesLicenseAllocationPriceNewLicenseAllocationBulkWithFiltersPriceConversionRateConfigUnion() {
}

// Satisfied by [shared.UnitConversionRateConfigParam],
// [shared.TieredConversionRateConfigParam],
// [PlanNewParamsPricesLicenseAllocationPriceNewLicenseAllocationBulkWithFiltersPriceConversionRateConfig].
type PlanNewParamsPricesLicenseAllocationPriceNewLicenseAllocationBulkWithFiltersPriceConversionRateConfigUnion interface {
	ImplementsPlanNewParamsPricesLicenseAllocationPriceNewLicenseAllocationBulkWithFiltersPriceConversionRateConfigUnion()
}

type PlanNewParamsPricesLicenseAllocationPriceNewLicenseAllocationBulkWithFiltersPriceConversionRateConfigConversionRateType string

const (
	PlanNewParamsPricesLicenseAllocationPriceNewLicenseAllocationBulkWithFiltersPriceConversionRateConfigConversionRateTypeUnit   PlanNewParamsPricesLicenseAllocationPriceNewLicenseAllocationBulkWithFiltersPriceConversionRateConfigConversionRateType = "unit"
	PlanNewParamsPricesLicenseAllocationPriceNewLicenseAllocationBulkWithFiltersPriceConversionRateConfigConversionRateTypeTiered PlanNewParamsPricesLicenseAllocationPriceNewLicenseAllocationBulkWithFiltersPriceConversionRateConfigConversionRateType = "tiered"
)

func (r PlanNewParamsPricesLicenseAllocationPriceNewLicenseAllocationBulkWithFiltersPriceConversionRateConfigConversionRateType) IsKnown() bool {
	switch r {
	case PlanNewParamsPricesLicenseAllocationPriceNewLicenseAllocationBulkWithFiltersPriceConversionRateConfigConversionRateTypeUnit, PlanNewParamsPricesLicenseAllocationPriceNewLicenseAllocationBulkWithFiltersPriceConversionRateConfigConversionRateTypeTiered:
		return true
	}
	return false
}

type PlanNewParamsPricesLicenseAllocationPriceNewLicenseAllocationPackagePrice struct {
	// The cadence to bill for this price on.
	Cadence param.Field[PlanNewParamsPricesLicenseAllocationPriceNewLicenseAllocationPackagePriceCadence] `json:"cadence" api:"required"`
	// The id of the item the price will be associated with.
	ItemID param.Field[string] `json:"item_id" api:"required"`
	// License allocations to associate with this price. Each entry defines a
	// per-license credit pool granted each cadence. Requires license_type_id or
	// license_type_configuration to be set.
	LicenseAllocations param.Field[[]PlanNewParamsPricesLicenseAllocationPriceNewLicenseAllocationPackagePriceLicenseAllocation] `json:"license_allocations" api:"required"`
	// The pricing model type
	ModelType param.Field[PlanNewParamsPricesLicenseAllocationPriceNewLicenseAllocationPackagePriceModelType] `json:"model_type" api:"required"`
	// The name of the price.
	Name param.Field[string] `json:"name" api:"required"`
	// Configuration for package pricing
	PackageConfig param.Field[shared.PackageConfigParam] `json:"package_config" api:"required"`
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
	ConversionRateConfig param.Field[PlanNewParamsPricesLicenseAllocationPriceNewLicenseAllocationPackagePriceConversionRateConfigUnion] `json:"conversion_rate_config"`
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
	// The ID of the license type to associate with this price.
	LicenseTypeID param.Field[string] `json:"license_type_id"`
	// User-specified key/value pairs for the resource. Individual keys can be removed
	// by setting the value to `null`, and the entire metadata mapping can be cleared
	// by setting `metadata` to `null`.
	Metadata param.Field[map[string]string] `json:"metadata"`
	// A transient ID that can be used to reference this price when adding adjustments
	// in the same API call.
	ReferenceID param.Field[string] `json:"reference_id"`
}

func (r PlanNewParamsPricesLicenseAllocationPriceNewLicenseAllocationPackagePrice) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r PlanNewParamsPricesLicenseAllocationPriceNewLicenseAllocationPackagePrice) implementsPlanNewParamsPricesLicenseAllocationPriceUnion() {
}

// The cadence to bill for this price on.
type PlanNewParamsPricesLicenseAllocationPriceNewLicenseAllocationPackagePriceCadence string

const (
	PlanNewParamsPricesLicenseAllocationPriceNewLicenseAllocationPackagePriceCadenceAnnual     PlanNewParamsPricesLicenseAllocationPriceNewLicenseAllocationPackagePriceCadence = "annual"
	PlanNewParamsPricesLicenseAllocationPriceNewLicenseAllocationPackagePriceCadenceSemiAnnual PlanNewParamsPricesLicenseAllocationPriceNewLicenseAllocationPackagePriceCadence = "semi_annual"
	PlanNewParamsPricesLicenseAllocationPriceNewLicenseAllocationPackagePriceCadenceMonthly    PlanNewParamsPricesLicenseAllocationPriceNewLicenseAllocationPackagePriceCadence = "monthly"
	PlanNewParamsPricesLicenseAllocationPriceNewLicenseAllocationPackagePriceCadenceQuarterly  PlanNewParamsPricesLicenseAllocationPriceNewLicenseAllocationPackagePriceCadence = "quarterly"
	PlanNewParamsPricesLicenseAllocationPriceNewLicenseAllocationPackagePriceCadenceOneTime    PlanNewParamsPricesLicenseAllocationPriceNewLicenseAllocationPackagePriceCadence = "one_time"
	PlanNewParamsPricesLicenseAllocationPriceNewLicenseAllocationPackagePriceCadenceCustom     PlanNewParamsPricesLicenseAllocationPriceNewLicenseAllocationPackagePriceCadence = "custom"
)

func (r PlanNewParamsPricesLicenseAllocationPriceNewLicenseAllocationPackagePriceCadence) IsKnown() bool {
	switch r {
	case PlanNewParamsPricesLicenseAllocationPriceNewLicenseAllocationPackagePriceCadenceAnnual, PlanNewParamsPricesLicenseAllocationPriceNewLicenseAllocationPackagePriceCadenceSemiAnnual, PlanNewParamsPricesLicenseAllocationPriceNewLicenseAllocationPackagePriceCadenceMonthly, PlanNewParamsPricesLicenseAllocationPriceNewLicenseAllocationPackagePriceCadenceQuarterly, PlanNewParamsPricesLicenseAllocationPriceNewLicenseAllocationPackagePriceCadenceOneTime, PlanNewParamsPricesLicenseAllocationPriceNewLicenseAllocationPackagePriceCadenceCustom:
		return true
	}
	return false
}

type PlanNewParamsPricesLicenseAllocationPriceNewLicenseAllocationPackagePriceLicenseAllocation struct {
	// The amount of credits granted per active license per cadence.
	Amount param.Field[string] `json:"amount" api:"required"`
	// The currency of the license allocation.
	Currency param.Field[string] `json:"currency" api:"required"`
	// When True, overage beyond the allocation is written off.
	WriteOffOverage param.Field[bool] `json:"write_off_overage"`
}

func (r PlanNewParamsPricesLicenseAllocationPriceNewLicenseAllocationPackagePriceLicenseAllocation) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// The pricing model type
type PlanNewParamsPricesLicenseAllocationPriceNewLicenseAllocationPackagePriceModelType string

const (
	PlanNewParamsPricesLicenseAllocationPriceNewLicenseAllocationPackagePriceModelTypePackage PlanNewParamsPricesLicenseAllocationPriceNewLicenseAllocationPackagePriceModelType = "package"
)

func (r PlanNewParamsPricesLicenseAllocationPriceNewLicenseAllocationPackagePriceModelType) IsKnown() bool {
	switch r {
	case PlanNewParamsPricesLicenseAllocationPriceNewLicenseAllocationPackagePriceModelTypePackage:
		return true
	}
	return false
}

type PlanNewParamsPricesLicenseAllocationPriceNewLicenseAllocationPackagePriceConversionRateConfig struct {
	ConversionRateType param.Field[PlanNewParamsPricesLicenseAllocationPriceNewLicenseAllocationPackagePriceConversionRateConfigConversionRateType] `json:"conversion_rate_type" api:"required"`
	TieredConfig       param.Field[shared.ConversionRateTieredConfigParam]                                                                          `json:"tiered_config"`
	UnitConfig         param.Field[shared.ConversionRateUnitConfigParam]                                                                            `json:"unit_config"`
}

func (r PlanNewParamsPricesLicenseAllocationPriceNewLicenseAllocationPackagePriceConversionRateConfig) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r PlanNewParamsPricesLicenseAllocationPriceNewLicenseAllocationPackagePriceConversionRateConfig) ImplementsPlanNewParamsPricesLicenseAllocationPriceNewLicenseAllocationPackagePriceConversionRateConfigUnion() {
}

// Satisfied by [shared.UnitConversionRateConfigParam],
// [shared.TieredConversionRateConfigParam],
// [PlanNewParamsPricesLicenseAllocationPriceNewLicenseAllocationPackagePriceConversionRateConfig].
type PlanNewParamsPricesLicenseAllocationPriceNewLicenseAllocationPackagePriceConversionRateConfigUnion interface {
	ImplementsPlanNewParamsPricesLicenseAllocationPriceNewLicenseAllocationPackagePriceConversionRateConfigUnion()
}

type PlanNewParamsPricesLicenseAllocationPriceNewLicenseAllocationPackagePriceConversionRateConfigConversionRateType string

const (
	PlanNewParamsPricesLicenseAllocationPriceNewLicenseAllocationPackagePriceConversionRateConfigConversionRateTypeUnit   PlanNewParamsPricesLicenseAllocationPriceNewLicenseAllocationPackagePriceConversionRateConfigConversionRateType = "unit"
	PlanNewParamsPricesLicenseAllocationPriceNewLicenseAllocationPackagePriceConversionRateConfigConversionRateTypeTiered PlanNewParamsPricesLicenseAllocationPriceNewLicenseAllocationPackagePriceConversionRateConfigConversionRateType = "tiered"
)

func (r PlanNewParamsPricesLicenseAllocationPriceNewLicenseAllocationPackagePriceConversionRateConfigConversionRateType) IsKnown() bool {
	switch r {
	case PlanNewParamsPricesLicenseAllocationPriceNewLicenseAllocationPackagePriceConversionRateConfigConversionRateTypeUnit, PlanNewParamsPricesLicenseAllocationPriceNewLicenseAllocationPackagePriceConversionRateConfigConversionRateTypeTiered:
		return true
	}
	return false
}

type PlanNewParamsPricesLicenseAllocationPriceNewLicenseAllocationMatrixPrice struct {
	// The cadence to bill for this price on.
	Cadence param.Field[PlanNewParamsPricesLicenseAllocationPriceNewLicenseAllocationMatrixPriceCadence] `json:"cadence" api:"required"`
	// The id of the item the price will be associated with.
	ItemID param.Field[string] `json:"item_id" api:"required"`
	// License allocations to associate with this price. Each entry defines a
	// per-license credit pool granted each cadence. Requires license_type_id or
	// license_type_configuration to be set.
	LicenseAllocations param.Field[[]PlanNewParamsPricesLicenseAllocationPriceNewLicenseAllocationMatrixPriceLicenseAllocation] `json:"license_allocations" api:"required"`
	// Configuration for matrix pricing
	MatrixConfig param.Field[shared.MatrixConfigParam] `json:"matrix_config" api:"required"`
	// The pricing model type
	ModelType param.Field[PlanNewParamsPricesLicenseAllocationPriceNewLicenseAllocationMatrixPriceModelType] `json:"model_type" api:"required"`
	// The name of the price.
	Name param.Field[string] `json:"name" api:"required"`
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
	ConversionRateConfig param.Field[PlanNewParamsPricesLicenseAllocationPriceNewLicenseAllocationMatrixPriceConversionRateConfigUnion] `json:"conversion_rate_config"`
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
	// The ID of the license type to associate with this price.
	LicenseTypeID param.Field[string] `json:"license_type_id"`
	// User-specified key/value pairs for the resource. Individual keys can be removed
	// by setting the value to `null`, and the entire metadata mapping can be cleared
	// by setting `metadata` to `null`.
	Metadata param.Field[map[string]string] `json:"metadata"`
	// A transient ID that can be used to reference this price when adding adjustments
	// in the same API call.
	ReferenceID param.Field[string] `json:"reference_id"`
}

func (r PlanNewParamsPricesLicenseAllocationPriceNewLicenseAllocationMatrixPrice) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r PlanNewParamsPricesLicenseAllocationPriceNewLicenseAllocationMatrixPrice) implementsPlanNewParamsPricesLicenseAllocationPriceUnion() {
}

// The cadence to bill for this price on.
type PlanNewParamsPricesLicenseAllocationPriceNewLicenseAllocationMatrixPriceCadence string

const (
	PlanNewParamsPricesLicenseAllocationPriceNewLicenseAllocationMatrixPriceCadenceAnnual     PlanNewParamsPricesLicenseAllocationPriceNewLicenseAllocationMatrixPriceCadence = "annual"
	PlanNewParamsPricesLicenseAllocationPriceNewLicenseAllocationMatrixPriceCadenceSemiAnnual PlanNewParamsPricesLicenseAllocationPriceNewLicenseAllocationMatrixPriceCadence = "semi_annual"
	PlanNewParamsPricesLicenseAllocationPriceNewLicenseAllocationMatrixPriceCadenceMonthly    PlanNewParamsPricesLicenseAllocationPriceNewLicenseAllocationMatrixPriceCadence = "monthly"
	PlanNewParamsPricesLicenseAllocationPriceNewLicenseAllocationMatrixPriceCadenceQuarterly  PlanNewParamsPricesLicenseAllocationPriceNewLicenseAllocationMatrixPriceCadence = "quarterly"
	PlanNewParamsPricesLicenseAllocationPriceNewLicenseAllocationMatrixPriceCadenceOneTime    PlanNewParamsPricesLicenseAllocationPriceNewLicenseAllocationMatrixPriceCadence = "one_time"
	PlanNewParamsPricesLicenseAllocationPriceNewLicenseAllocationMatrixPriceCadenceCustom     PlanNewParamsPricesLicenseAllocationPriceNewLicenseAllocationMatrixPriceCadence = "custom"
)

func (r PlanNewParamsPricesLicenseAllocationPriceNewLicenseAllocationMatrixPriceCadence) IsKnown() bool {
	switch r {
	case PlanNewParamsPricesLicenseAllocationPriceNewLicenseAllocationMatrixPriceCadenceAnnual, PlanNewParamsPricesLicenseAllocationPriceNewLicenseAllocationMatrixPriceCadenceSemiAnnual, PlanNewParamsPricesLicenseAllocationPriceNewLicenseAllocationMatrixPriceCadenceMonthly, PlanNewParamsPricesLicenseAllocationPriceNewLicenseAllocationMatrixPriceCadenceQuarterly, PlanNewParamsPricesLicenseAllocationPriceNewLicenseAllocationMatrixPriceCadenceOneTime, PlanNewParamsPricesLicenseAllocationPriceNewLicenseAllocationMatrixPriceCadenceCustom:
		return true
	}
	return false
}

type PlanNewParamsPricesLicenseAllocationPriceNewLicenseAllocationMatrixPriceLicenseAllocation struct {
	// The amount of credits granted per active license per cadence.
	Amount param.Field[string] `json:"amount" api:"required"`
	// The currency of the license allocation.
	Currency param.Field[string] `json:"currency" api:"required"`
	// When True, overage beyond the allocation is written off.
	WriteOffOverage param.Field[bool] `json:"write_off_overage"`
}

func (r PlanNewParamsPricesLicenseAllocationPriceNewLicenseAllocationMatrixPriceLicenseAllocation) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// The pricing model type
type PlanNewParamsPricesLicenseAllocationPriceNewLicenseAllocationMatrixPriceModelType string

const (
	PlanNewParamsPricesLicenseAllocationPriceNewLicenseAllocationMatrixPriceModelTypeMatrix PlanNewParamsPricesLicenseAllocationPriceNewLicenseAllocationMatrixPriceModelType = "matrix"
)

func (r PlanNewParamsPricesLicenseAllocationPriceNewLicenseAllocationMatrixPriceModelType) IsKnown() bool {
	switch r {
	case PlanNewParamsPricesLicenseAllocationPriceNewLicenseAllocationMatrixPriceModelTypeMatrix:
		return true
	}
	return false
}

type PlanNewParamsPricesLicenseAllocationPriceNewLicenseAllocationMatrixPriceConversionRateConfig struct {
	ConversionRateType param.Field[PlanNewParamsPricesLicenseAllocationPriceNewLicenseAllocationMatrixPriceConversionRateConfigConversionRateType] `json:"conversion_rate_type" api:"required"`
	TieredConfig       param.Field[shared.ConversionRateTieredConfigParam]                                                                         `json:"tiered_config"`
	UnitConfig         param.Field[shared.ConversionRateUnitConfigParam]                                                                           `json:"unit_config"`
}

func (r PlanNewParamsPricesLicenseAllocationPriceNewLicenseAllocationMatrixPriceConversionRateConfig) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r PlanNewParamsPricesLicenseAllocationPriceNewLicenseAllocationMatrixPriceConversionRateConfig) ImplementsPlanNewParamsPricesLicenseAllocationPriceNewLicenseAllocationMatrixPriceConversionRateConfigUnion() {
}

// Satisfied by [shared.UnitConversionRateConfigParam],
// [shared.TieredConversionRateConfigParam],
// [PlanNewParamsPricesLicenseAllocationPriceNewLicenseAllocationMatrixPriceConversionRateConfig].
type PlanNewParamsPricesLicenseAllocationPriceNewLicenseAllocationMatrixPriceConversionRateConfigUnion interface {
	ImplementsPlanNewParamsPricesLicenseAllocationPriceNewLicenseAllocationMatrixPriceConversionRateConfigUnion()
}

type PlanNewParamsPricesLicenseAllocationPriceNewLicenseAllocationMatrixPriceConversionRateConfigConversionRateType string

const (
	PlanNewParamsPricesLicenseAllocationPriceNewLicenseAllocationMatrixPriceConversionRateConfigConversionRateTypeUnit   PlanNewParamsPricesLicenseAllocationPriceNewLicenseAllocationMatrixPriceConversionRateConfigConversionRateType = "unit"
	PlanNewParamsPricesLicenseAllocationPriceNewLicenseAllocationMatrixPriceConversionRateConfigConversionRateTypeTiered PlanNewParamsPricesLicenseAllocationPriceNewLicenseAllocationMatrixPriceConversionRateConfigConversionRateType = "tiered"
)

func (r PlanNewParamsPricesLicenseAllocationPriceNewLicenseAllocationMatrixPriceConversionRateConfigConversionRateType) IsKnown() bool {
	switch r {
	case PlanNewParamsPricesLicenseAllocationPriceNewLicenseAllocationMatrixPriceConversionRateConfigConversionRateTypeUnit, PlanNewParamsPricesLicenseAllocationPriceNewLicenseAllocationMatrixPriceConversionRateConfigConversionRateTypeTiered:
		return true
	}
	return false
}

type PlanNewParamsPricesLicenseAllocationPriceNewLicenseAllocationThresholdTotalAmountPrice struct {
	// The cadence to bill for this price on.
	Cadence param.Field[PlanNewParamsPricesLicenseAllocationPriceNewLicenseAllocationThresholdTotalAmountPriceCadence] `json:"cadence" api:"required"`
	// The id of the item the price will be associated with.
	ItemID param.Field[string] `json:"item_id" api:"required"`
	// License allocations to associate with this price. Each entry defines a
	// per-license credit pool granted each cadence. Requires license_type_id or
	// license_type_configuration to be set.
	LicenseAllocations param.Field[[]PlanNewParamsPricesLicenseAllocationPriceNewLicenseAllocationThresholdTotalAmountPriceLicenseAllocation] `json:"license_allocations" api:"required"`
	// The pricing model type
	ModelType param.Field[PlanNewParamsPricesLicenseAllocationPriceNewLicenseAllocationThresholdTotalAmountPriceModelType] `json:"model_type" api:"required"`
	// The name of the price.
	Name param.Field[string] `json:"name" api:"required"`
	// Configuration for threshold_total_amount pricing
	ThresholdTotalAmountConfig param.Field[PlanNewParamsPricesLicenseAllocationPriceNewLicenseAllocationThresholdTotalAmountPriceThresholdTotalAmountConfig] `json:"threshold_total_amount_config" api:"required"`
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
	ConversionRateConfig param.Field[PlanNewParamsPricesLicenseAllocationPriceNewLicenseAllocationThresholdTotalAmountPriceConversionRateConfigUnion] `json:"conversion_rate_config"`
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
	// The ID of the license type to associate with this price.
	LicenseTypeID param.Field[string] `json:"license_type_id"`
	// User-specified key/value pairs for the resource. Individual keys can be removed
	// by setting the value to `null`, and the entire metadata mapping can be cleared
	// by setting `metadata` to `null`.
	Metadata param.Field[map[string]string] `json:"metadata"`
	// A transient ID that can be used to reference this price when adding adjustments
	// in the same API call.
	ReferenceID param.Field[string] `json:"reference_id"`
}

func (r PlanNewParamsPricesLicenseAllocationPriceNewLicenseAllocationThresholdTotalAmountPrice) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r PlanNewParamsPricesLicenseAllocationPriceNewLicenseAllocationThresholdTotalAmountPrice) implementsPlanNewParamsPricesLicenseAllocationPriceUnion() {
}

// The cadence to bill for this price on.
type PlanNewParamsPricesLicenseAllocationPriceNewLicenseAllocationThresholdTotalAmountPriceCadence string

const (
	PlanNewParamsPricesLicenseAllocationPriceNewLicenseAllocationThresholdTotalAmountPriceCadenceAnnual     PlanNewParamsPricesLicenseAllocationPriceNewLicenseAllocationThresholdTotalAmountPriceCadence = "annual"
	PlanNewParamsPricesLicenseAllocationPriceNewLicenseAllocationThresholdTotalAmountPriceCadenceSemiAnnual PlanNewParamsPricesLicenseAllocationPriceNewLicenseAllocationThresholdTotalAmountPriceCadence = "semi_annual"
	PlanNewParamsPricesLicenseAllocationPriceNewLicenseAllocationThresholdTotalAmountPriceCadenceMonthly    PlanNewParamsPricesLicenseAllocationPriceNewLicenseAllocationThresholdTotalAmountPriceCadence = "monthly"
	PlanNewParamsPricesLicenseAllocationPriceNewLicenseAllocationThresholdTotalAmountPriceCadenceQuarterly  PlanNewParamsPricesLicenseAllocationPriceNewLicenseAllocationThresholdTotalAmountPriceCadence = "quarterly"
	PlanNewParamsPricesLicenseAllocationPriceNewLicenseAllocationThresholdTotalAmountPriceCadenceOneTime    PlanNewParamsPricesLicenseAllocationPriceNewLicenseAllocationThresholdTotalAmountPriceCadence = "one_time"
	PlanNewParamsPricesLicenseAllocationPriceNewLicenseAllocationThresholdTotalAmountPriceCadenceCustom     PlanNewParamsPricesLicenseAllocationPriceNewLicenseAllocationThresholdTotalAmountPriceCadence = "custom"
)

func (r PlanNewParamsPricesLicenseAllocationPriceNewLicenseAllocationThresholdTotalAmountPriceCadence) IsKnown() bool {
	switch r {
	case PlanNewParamsPricesLicenseAllocationPriceNewLicenseAllocationThresholdTotalAmountPriceCadenceAnnual, PlanNewParamsPricesLicenseAllocationPriceNewLicenseAllocationThresholdTotalAmountPriceCadenceSemiAnnual, PlanNewParamsPricesLicenseAllocationPriceNewLicenseAllocationThresholdTotalAmountPriceCadenceMonthly, PlanNewParamsPricesLicenseAllocationPriceNewLicenseAllocationThresholdTotalAmountPriceCadenceQuarterly, PlanNewParamsPricesLicenseAllocationPriceNewLicenseAllocationThresholdTotalAmountPriceCadenceOneTime, PlanNewParamsPricesLicenseAllocationPriceNewLicenseAllocationThresholdTotalAmountPriceCadenceCustom:
		return true
	}
	return false
}

type PlanNewParamsPricesLicenseAllocationPriceNewLicenseAllocationThresholdTotalAmountPriceLicenseAllocation struct {
	// The amount of credits granted per active license per cadence.
	Amount param.Field[string] `json:"amount" api:"required"`
	// The currency of the license allocation.
	Currency param.Field[string] `json:"currency" api:"required"`
	// When True, overage beyond the allocation is written off.
	WriteOffOverage param.Field[bool] `json:"write_off_overage"`
}

func (r PlanNewParamsPricesLicenseAllocationPriceNewLicenseAllocationThresholdTotalAmountPriceLicenseAllocation) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// The pricing model type
type PlanNewParamsPricesLicenseAllocationPriceNewLicenseAllocationThresholdTotalAmountPriceModelType string

const (
	PlanNewParamsPricesLicenseAllocationPriceNewLicenseAllocationThresholdTotalAmountPriceModelTypeThresholdTotalAmount PlanNewParamsPricesLicenseAllocationPriceNewLicenseAllocationThresholdTotalAmountPriceModelType = "threshold_total_amount"
)

func (r PlanNewParamsPricesLicenseAllocationPriceNewLicenseAllocationThresholdTotalAmountPriceModelType) IsKnown() bool {
	switch r {
	case PlanNewParamsPricesLicenseAllocationPriceNewLicenseAllocationThresholdTotalAmountPriceModelTypeThresholdTotalAmount:
		return true
	}
	return false
}

// Configuration for threshold_total_amount pricing
type PlanNewParamsPricesLicenseAllocationPriceNewLicenseAllocationThresholdTotalAmountPriceThresholdTotalAmountConfig struct {
	// When the quantity consumed passes a provided threshold, the configured total
	// will be charged
	ConsumptionTable param.Field[[]PlanNewParamsPricesLicenseAllocationPriceNewLicenseAllocationThresholdTotalAmountPriceThresholdTotalAmountConfigConsumptionTable] `json:"consumption_table" api:"required"`
	// If true, the unit price will be prorated to the billing period
	Prorate param.Field[bool] `json:"prorate"`
}

func (r PlanNewParamsPricesLicenseAllocationPriceNewLicenseAllocationThresholdTotalAmountPriceThresholdTotalAmountConfig) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Configuration for a single threshold
type PlanNewParamsPricesLicenseAllocationPriceNewLicenseAllocationThresholdTotalAmountPriceThresholdTotalAmountConfigConsumptionTable struct {
	Threshold param.Field[string] `json:"threshold" api:"required"`
	// Total amount for this threshold
	TotalAmount param.Field[string] `json:"total_amount" api:"required"`
}

func (r PlanNewParamsPricesLicenseAllocationPriceNewLicenseAllocationThresholdTotalAmountPriceThresholdTotalAmountConfigConsumptionTable) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type PlanNewParamsPricesLicenseAllocationPriceNewLicenseAllocationThresholdTotalAmountPriceConversionRateConfig struct {
	ConversionRateType param.Field[PlanNewParamsPricesLicenseAllocationPriceNewLicenseAllocationThresholdTotalAmountPriceConversionRateConfigConversionRateType] `json:"conversion_rate_type" api:"required"`
	TieredConfig       param.Field[shared.ConversionRateTieredConfigParam]                                                                                       `json:"tiered_config"`
	UnitConfig         param.Field[shared.ConversionRateUnitConfigParam]                                                                                         `json:"unit_config"`
}

func (r PlanNewParamsPricesLicenseAllocationPriceNewLicenseAllocationThresholdTotalAmountPriceConversionRateConfig) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r PlanNewParamsPricesLicenseAllocationPriceNewLicenseAllocationThresholdTotalAmountPriceConversionRateConfig) ImplementsPlanNewParamsPricesLicenseAllocationPriceNewLicenseAllocationThresholdTotalAmountPriceConversionRateConfigUnion() {
}

// Satisfied by [shared.UnitConversionRateConfigParam],
// [shared.TieredConversionRateConfigParam],
// [PlanNewParamsPricesLicenseAllocationPriceNewLicenseAllocationThresholdTotalAmountPriceConversionRateConfig].
type PlanNewParamsPricesLicenseAllocationPriceNewLicenseAllocationThresholdTotalAmountPriceConversionRateConfigUnion interface {
	ImplementsPlanNewParamsPricesLicenseAllocationPriceNewLicenseAllocationThresholdTotalAmountPriceConversionRateConfigUnion()
}

type PlanNewParamsPricesLicenseAllocationPriceNewLicenseAllocationThresholdTotalAmountPriceConversionRateConfigConversionRateType string

const (
	PlanNewParamsPricesLicenseAllocationPriceNewLicenseAllocationThresholdTotalAmountPriceConversionRateConfigConversionRateTypeUnit   PlanNewParamsPricesLicenseAllocationPriceNewLicenseAllocationThresholdTotalAmountPriceConversionRateConfigConversionRateType = "unit"
	PlanNewParamsPricesLicenseAllocationPriceNewLicenseAllocationThresholdTotalAmountPriceConversionRateConfigConversionRateTypeTiered PlanNewParamsPricesLicenseAllocationPriceNewLicenseAllocationThresholdTotalAmountPriceConversionRateConfigConversionRateType = "tiered"
)

func (r PlanNewParamsPricesLicenseAllocationPriceNewLicenseAllocationThresholdTotalAmountPriceConversionRateConfigConversionRateType) IsKnown() bool {
	switch r {
	case PlanNewParamsPricesLicenseAllocationPriceNewLicenseAllocationThresholdTotalAmountPriceConversionRateConfigConversionRateTypeUnit, PlanNewParamsPricesLicenseAllocationPriceNewLicenseAllocationThresholdTotalAmountPriceConversionRateConfigConversionRateTypeTiered:
		return true
	}
	return false
}

type PlanNewParamsPricesLicenseAllocationPriceNewLicenseAllocationTieredPackagePrice struct {
	// The cadence to bill for this price on.
	Cadence param.Field[PlanNewParamsPricesLicenseAllocationPriceNewLicenseAllocationTieredPackagePriceCadence] `json:"cadence" api:"required"`
	// The id of the item the price will be associated with.
	ItemID param.Field[string] `json:"item_id" api:"required"`
	// License allocations to associate with this price. Each entry defines a
	// per-license credit pool granted each cadence. Requires license_type_id or
	// license_type_configuration to be set.
	LicenseAllocations param.Field[[]PlanNewParamsPricesLicenseAllocationPriceNewLicenseAllocationTieredPackagePriceLicenseAllocation] `json:"license_allocations" api:"required"`
	// The pricing model type
	ModelType param.Field[PlanNewParamsPricesLicenseAllocationPriceNewLicenseAllocationTieredPackagePriceModelType] `json:"model_type" api:"required"`
	// The name of the price.
	Name param.Field[string] `json:"name" api:"required"`
	// Configuration for tiered_package pricing
	TieredPackageConfig param.Field[PlanNewParamsPricesLicenseAllocationPriceNewLicenseAllocationTieredPackagePriceTieredPackageConfig] `json:"tiered_package_config" api:"required"`
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
	ConversionRateConfig param.Field[PlanNewParamsPricesLicenseAllocationPriceNewLicenseAllocationTieredPackagePriceConversionRateConfigUnion] `json:"conversion_rate_config"`
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
	// The ID of the license type to associate with this price.
	LicenseTypeID param.Field[string] `json:"license_type_id"`
	// User-specified key/value pairs for the resource. Individual keys can be removed
	// by setting the value to `null`, and the entire metadata mapping can be cleared
	// by setting `metadata` to `null`.
	Metadata param.Field[map[string]string] `json:"metadata"`
	// A transient ID that can be used to reference this price when adding adjustments
	// in the same API call.
	ReferenceID param.Field[string] `json:"reference_id"`
}

func (r PlanNewParamsPricesLicenseAllocationPriceNewLicenseAllocationTieredPackagePrice) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r PlanNewParamsPricesLicenseAllocationPriceNewLicenseAllocationTieredPackagePrice) implementsPlanNewParamsPricesLicenseAllocationPriceUnion() {
}

// The cadence to bill for this price on.
type PlanNewParamsPricesLicenseAllocationPriceNewLicenseAllocationTieredPackagePriceCadence string

const (
	PlanNewParamsPricesLicenseAllocationPriceNewLicenseAllocationTieredPackagePriceCadenceAnnual     PlanNewParamsPricesLicenseAllocationPriceNewLicenseAllocationTieredPackagePriceCadence = "annual"
	PlanNewParamsPricesLicenseAllocationPriceNewLicenseAllocationTieredPackagePriceCadenceSemiAnnual PlanNewParamsPricesLicenseAllocationPriceNewLicenseAllocationTieredPackagePriceCadence = "semi_annual"
	PlanNewParamsPricesLicenseAllocationPriceNewLicenseAllocationTieredPackagePriceCadenceMonthly    PlanNewParamsPricesLicenseAllocationPriceNewLicenseAllocationTieredPackagePriceCadence = "monthly"
	PlanNewParamsPricesLicenseAllocationPriceNewLicenseAllocationTieredPackagePriceCadenceQuarterly  PlanNewParamsPricesLicenseAllocationPriceNewLicenseAllocationTieredPackagePriceCadence = "quarterly"
	PlanNewParamsPricesLicenseAllocationPriceNewLicenseAllocationTieredPackagePriceCadenceOneTime    PlanNewParamsPricesLicenseAllocationPriceNewLicenseAllocationTieredPackagePriceCadence = "one_time"
	PlanNewParamsPricesLicenseAllocationPriceNewLicenseAllocationTieredPackagePriceCadenceCustom     PlanNewParamsPricesLicenseAllocationPriceNewLicenseAllocationTieredPackagePriceCadence = "custom"
)

func (r PlanNewParamsPricesLicenseAllocationPriceNewLicenseAllocationTieredPackagePriceCadence) IsKnown() bool {
	switch r {
	case PlanNewParamsPricesLicenseAllocationPriceNewLicenseAllocationTieredPackagePriceCadenceAnnual, PlanNewParamsPricesLicenseAllocationPriceNewLicenseAllocationTieredPackagePriceCadenceSemiAnnual, PlanNewParamsPricesLicenseAllocationPriceNewLicenseAllocationTieredPackagePriceCadenceMonthly, PlanNewParamsPricesLicenseAllocationPriceNewLicenseAllocationTieredPackagePriceCadenceQuarterly, PlanNewParamsPricesLicenseAllocationPriceNewLicenseAllocationTieredPackagePriceCadenceOneTime, PlanNewParamsPricesLicenseAllocationPriceNewLicenseAllocationTieredPackagePriceCadenceCustom:
		return true
	}
	return false
}

type PlanNewParamsPricesLicenseAllocationPriceNewLicenseAllocationTieredPackagePriceLicenseAllocation struct {
	// The amount of credits granted per active license per cadence.
	Amount param.Field[string] `json:"amount" api:"required"`
	// The currency of the license allocation.
	Currency param.Field[string] `json:"currency" api:"required"`
	// When True, overage beyond the allocation is written off.
	WriteOffOverage param.Field[bool] `json:"write_off_overage"`
}

func (r PlanNewParamsPricesLicenseAllocationPriceNewLicenseAllocationTieredPackagePriceLicenseAllocation) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// The pricing model type
type PlanNewParamsPricesLicenseAllocationPriceNewLicenseAllocationTieredPackagePriceModelType string

const (
	PlanNewParamsPricesLicenseAllocationPriceNewLicenseAllocationTieredPackagePriceModelTypeTieredPackage PlanNewParamsPricesLicenseAllocationPriceNewLicenseAllocationTieredPackagePriceModelType = "tiered_package"
)

func (r PlanNewParamsPricesLicenseAllocationPriceNewLicenseAllocationTieredPackagePriceModelType) IsKnown() bool {
	switch r {
	case PlanNewParamsPricesLicenseAllocationPriceNewLicenseAllocationTieredPackagePriceModelTypeTieredPackage:
		return true
	}
	return false
}

// Configuration for tiered_package pricing
type PlanNewParamsPricesLicenseAllocationPriceNewLicenseAllocationTieredPackagePriceTieredPackageConfig struct {
	PackageSize param.Field[string] `json:"package_size" api:"required"`
	// Apply tiered pricing after rounding up the quantity to the package size. Tiers
	// are defined using exclusive lower bounds. The tier bounds are defined based on
	// the total quantity rather than the number of packages, so they must be multiples
	// of the package size.
	Tiers param.Field[[]PlanNewParamsPricesLicenseAllocationPriceNewLicenseAllocationTieredPackagePriceTieredPackageConfigTier] `json:"tiers" api:"required"`
}

func (r PlanNewParamsPricesLicenseAllocationPriceNewLicenseAllocationTieredPackagePriceTieredPackageConfig) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Configuration for a single tier with business logic
type PlanNewParamsPricesLicenseAllocationPriceNewLicenseAllocationTieredPackagePriceTieredPackageConfigTier struct {
	// Price per package
	PerUnit        param.Field[string] `json:"per_unit" api:"required"`
	TierLowerBound param.Field[string] `json:"tier_lower_bound" api:"required"`
}

func (r PlanNewParamsPricesLicenseAllocationPriceNewLicenseAllocationTieredPackagePriceTieredPackageConfigTier) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type PlanNewParamsPricesLicenseAllocationPriceNewLicenseAllocationTieredPackagePriceConversionRateConfig struct {
	ConversionRateType param.Field[PlanNewParamsPricesLicenseAllocationPriceNewLicenseAllocationTieredPackagePriceConversionRateConfigConversionRateType] `json:"conversion_rate_type" api:"required"`
	TieredConfig       param.Field[shared.ConversionRateTieredConfigParam]                                                                                `json:"tiered_config"`
	UnitConfig         param.Field[shared.ConversionRateUnitConfigParam]                                                                                  `json:"unit_config"`
}

func (r PlanNewParamsPricesLicenseAllocationPriceNewLicenseAllocationTieredPackagePriceConversionRateConfig) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r PlanNewParamsPricesLicenseAllocationPriceNewLicenseAllocationTieredPackagePriceConversionRateConfig) ImplementsPlanNewParamsPricesLicenseAllocationPriceNewLicenseAllocationTieredPackagePriceConversionRateConfigUnion() {
}

// Satisfied by [shared.UnitConversionRateConfigParam],
// [shared.TieredConversionRateConfigParam],
// [PlanNewParamsPricesLicenseAllocationPriceNewLicenseAllocationTieredPackagePriceConversionRateConfig].
type PlanNewParamsPricesLicenseAllocationPriceNewLicenseAllocationTieredPackagePriceConversionRateConfigUnion interface {
	ImplementsPlanNewParamsPricesLicenseAllocationPriceNewLicenseAllocationTieredPackagePriceConversionRateConfigUnion()
}

type PlanNewParamsPricesLicenseAllocationPriceNewLicenseAllocationTieredPackagePriceConversionRateConfigConversionRateType string

const (
	PlanNewParamsPricesLicenseAllocationPriceNewLicenseAllocationTieredPackagePriceConversionRateConfigConversionRateTypeUnit   PlanNewParamsPricesLicenseAllocationPriceNewLicenseAllocationTieredPackagePriceConversionRateConfigConversionRateType = "unit"
	PlanNewParamsPricesLicenseAllocationPriceNewLicenseAllocationTieredPackagePriceConversionRateConfigConversionRateTypeTiered PlanNewParamsPricesLicenseAllocationPriceNewLicenseAllocationTieredPackagePriceConversionRateConfigConversionRateType = "tiered"
)

func (r PlanNewParamsPricesLicenseAllocationPriceNewLicenseAllocationTieredPackagePriceConversionRateConfigConversionRateType) IsKnown() bool {
	switch r {
	case PlanNewParamsPricesLicenseAllocationPriceNewLicenseAllocationTieredPackagePriceConversionRateConfigConversionRateTypeUnit, PlanNewParamsPricesLicenseAllocationPriceNewLicenseAllocationTieredPackagePriceConversionRateConfigConversionRateTypeTiered:
		return true
	}
	return false
}

type PlanNewParamsPricesLicenseAllocationPriceNewLicenseAllocationTieredWithMinimumPrice struct {
	// The cadence to bill for this price on.
	Cadence param.Field[PlanNewParamsPricesLicenseAllocationPriceNewLicenseAllocationTieredWithMinimumPriceCadence] `json:"cadence" api:"required"`
	// The id of the item the price will be associated with.
	ItemID param.Field[string] `json:"item_id" api:"required"`
	// License allocations to associate with this price. Each entry defines a
	// per-license credit pool granted each cadence. Requires license_type_id or
	// license_type_configuration to be set.
	LicenseAllocations param.Field[[]PlanNewParamsPricesLicenseAllocationPriceNewLicenseAllocationTieredWithMinimumPriceLicenseAllocation] `json:"license_allocations" api:"required"`
	// The pricing model type
	ModelType param.Field[PlanNewParamsPricesLicenseAllocationPriceNewLicenseAllocationTieredWithMinimumPriceModelType] `json:"model_type" api:"required"`
	// The name of the price.
	Name param.Field[string] `json:"name" api:"required"`
	// Configuration for tiered_with_minimum pricing
	TieredWithMinimumConfig param.Field[PlanNewParamsPricesLicenseAllocationPriceNewLicenseAllocationTieredWithMinimumPriceTieredWithMinimumConfig] `json:"tiered_with_minimum_config" api:"required"`
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
	ConversionRateConfig param.Field[PlanNewParamsPricesLicenseAllocationPriceNewLicenseAllocationTieredWithMinimumPriceConversionRateConfigUnion] `json:"conversion_rate_config"`
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
	// The ID of the license type to associate with this price.
	LicenseTypeID param.Field[string] `json:"license_type_id"`
	// User-specified key/value pairs for the resource. Individual keys can be removed
	// by setting the value to `null`, and the entire metadata mapping can be cleared
	// by setting `metadata` to `null`.
	Metadata param.Field[map[string]string] `json:"metadata"`
	// A transient ID that can be used to reference this price when adding adjustments
	// in the same API call.
	ReferenceID param.Field[string] `json:"reference_id"`
}

func (r PlanNewParamsPricesLicenseAllocationPriceNewLicenseAllocationTieredWithMinimumPrice) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r PlanNewParamsPricesLicenseAllocationPriceNewLicenseAllocationTieredWithMinimumPrice) implementsPlanNewParamsPricesLicenseAllocationPriceUnion() {
}

// The cadence to bill for this price on.
type PlanNewParamsPricesLicenseAllocationPriceNewLicenseAllocationTieredWithMinimumPriceCadence string

const (
	PlanNewParamsPricesLicenseAllocationPriceNewLicenseAllocationTieredWithMinimumPriceCadenceAnnual     PlanNewParamsPricesLicenseAllocationPriceNewLicenseAllocationTieredWithMinimumPriceCadence = "annual"
	PlanNewParamsPricesLicenseAllocationPriceNewLicenseAllocationTieredWithMinimumPriceCadenceSemiAnnual PlanNewParamsPricesLicenseAllocationPriceNewLicenseAllocationTieredWithMinimumPriceCadence = "semi_annual"
	PlanNewParamsPricesLicenseAllocationPriceNewLicenseAllocationTieredWithMinimumPriceCadenceMonthly    PlanNewParamsPricesLicenseAllocationPriceNewLicenseAllocationTieredWithMinimumPriceCadence = "monthly"
	PlanNewParamsPricesLicenseAllocationPriceNewLicenseAllocationTieredWithMinimumPriceCadenceQuarterly  PlanNewParamsPricesLicenseAllocationPriceNewLicenseAllocationTieredWithMinimumPriceCadence = "quarterly"
	PlanNewParamsPricesLicenseAllocationPriceNewLicenseAllocationTieredWithMinimumPriceCadenceOneTime    PlanNewParamsPricesLicenseAllocationPriceNewLicenseAllocationTieredWithMinimumPriceCadence = "one_time"
	PlanNewParamsPricesLicenseAllocationPriceNewLicenseAllocationTieredWithMinimumPriceCadenceCustom     PlanNewParamsPricesLicenseAllocationPriceNewLicenseAllocationTieredWithMinimumPriceCadence = "custom"
)

func (r PlanNewParamsPricesLicenseAllocationPriceNewLicenseAllocationTieredWithMinimumPriceCadence) IsKnown() bool {
	switch r {
	case PlanNewParamsPricesLicenseAllocationPriceNewLicenseAllocationTieredWithMinimumPriceCadenceAnnual, PlanNewParamsPricesLicenseAllocationPriceNewLicenseAllocationTieredWithMinimumPriceCadenceSemiAnnual, PlanNewParamsPricesLicenseAllocationPriceNewLicenseAllocationTieredWithMinimumPriceCadenceMonthly, PlanNewParamsPricesLicenseAllocationPriceNewLicenseAllocationTieredWithMinimumPriceCadenceQuarterly, PlanNewParamsPricesLicenseAllocationPriceNewLicenseAllocationTieredWithMinimumPriceCadenceOneTime, PlanNewParamsPricesLicenseAllocationPriceNewLicenseAllocationTieredWithMinimumPriceCadenceCustom:
		return true
	}
	return false
}

type PlanNewParamsPricesLicenseAllocationPriceNewLicenseAllocationTieredWithMinimumPriceLicenseAllocation struct {
	// The amount of credits granted per active license per cadence.
	Amount param.Field[string] `json:"amount" api:"required"`
	// The currency of the license allocation.
	Currency param.Field[string] `json:"currency" api:"required"`
	// When True, overage beyond the allocation is written off.
	WriteOffOverage param.Field[bool] `json:"write_off_overage"`
}

func (r PlanNewParamsPricesLicenseAllocationPriceNewLicenseAllocationTieredWithMinimumPriceLicenseAllocation) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// The pricing model type
type PlanNewParamsPricesLicenseAllocationPriceNewLicenseAllocationTieredWithMinimumPriceModelType string

const (
	PlanNewParamsPricesLicenseAllocationPriceNewLicenseAllocationTieredWithMinimumPriceModelTypeTieredWithMinimum PlanNewParamsPricesLicenseAllocationPriceNewLicenseAllocationTieredWithMinimumPriceModelType = "tiered_with_minimum"
)

func (r PlanNewParamsPricesLicenseAllocationPriceNewLicenseAllocationTieredWithMinimumPriceModelType) IsKnown() bool {
	switch r {
	case PlanNewParamsPricesLicenseAllocationPriceNewLicenseAllocationTieredWithMinimumPriceModelTypeTieredWithMinimum:
		return true
	}
	return false
}

// Configuration for tiered_with_minimum pricing
type PlanNewParamsPricesLicenseAllocationPriceNewLicenseAllocationTieredWithMinimumPriceTieredWithMinimumConfig struct {
	// Tiered pricing with a minimum amount dependent on the volume tier. Tiers are
	// defined using exclusive lower bounds.
	Tiers param.Field[[]PlanNewParamsPricesLicenseAllocationPriceNewLicenseAllocationTieredWithMinimumPriceTieredWithMinimumConfigTier] `json:"tiers" api:"required"`
	// If true, tiers with an accrued amount of 0 will not be included in the rating.
	HideZeroAmountTiers param.Field[bool] `json:"hide_zero_amount_tiers"`
	// If true, the unit price will be prorated to the billing period
	Prorate param.Field[bool] `json:"prorate"`
}

func (r PlanNewParamsPricesLicenseAllocationPriceNewLicenseAllocationTieredWithMinimumPriceTieredWithMinimumConfig) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Configuration for a single tier
type PlanNewParamsPricesLicenseAllocationPriceNewLicenseAllocationTieredWithMinimumPriceTieredWithMinimumConfigTier struct {
	MinimumAmount  param.Field[string] `json:"minimum_amount" api:"required"`
	TierLowerBound param.Field[string] `json:"tier_lower_bound" api:"required"`
	// Per unit amount
	UnitAmount param.Field[string] `json:"unit_amount" api:"required"`
}

func (r PlanNewParamsPricesLicenseAllocationPriceNewLicenseAllocationTieredWithMinimumPriceTieredWithMinimumConfigTier) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type PlanNewParamsPricesLicenseAllocationPriceNewLicenseAllocationTieredWithMinimumPriceConversionRateConfig struct {
	ConversionRateType param.Field[PlanNewParamsPricesLicenseAllocationPriceNewLicenseAllocationTieredWithMinimumPriceConversionRateConfigConversionRateType] `json:"conversion_rate_type" api:"required"`
	TieredConfig       param.Field[shared.ConversionRateTieredConfigParam]                                                                                    `json:"tiered_config"`
	UnitConfig         param.Field[shared.ConversionRateUnitConfigParam]                                                                                      `json:"unit_config"`
}

func (r PlanNewParamsPricesLicenseAllocationPriceNewLicenseAllocationTieredWithMinimumPriceConversionRateConfig) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r PlanNewParamsPricesLicenseAllocationPriceNewLicenseAllocationTieredWithMinimumPriceConversionRateConfig) ImplementsPlanNewParamsPricesLicenseAllocationPriceNewLicenseAllocationTieredWithMinimumPriceConversionRateConfigUnion() {
}

// Satisfied by [shared.UnitConversionRateConfigParam],
// [shared.TieredConversionRateConfigParam],
// [PlanNewParamsPricesLicenseAllocationPriceNewLicenseAllocationTieredWithMinimumPriceConversionRateConfig].
type PlanNewParamsPricesLicenseAllocationPriceNewLicenseAllocationTieredWithMinimumPriceConversionRateConfigUnion interface {
	ImplementsPlanNewParamsPricesLicenseAllocationPriceNewLicenseAllocationTieredWithMinimumPriceConversionRateConfigUnion()
}

type PlanNewParamsPricesLicenseAllocationPriceNewLicenseAllocationTieredWithMinimumPriceConversionRateConfigConversionRateType string

const (
	PlanNewParamsPricesLicenseAllocationPriceNewLicenseAllocationTieredWithMinimumPriceConversionRateConfigConversionRateTypeUnit   PlanNewParamsPricesLicenseAllocationPriceNewLicenseAllocationTieredWithMinimumPriceConversionRateConfigConversionRateType = "unit"
	PlanNewParamsPricesLicenseAllocationPriceNewLicenseAllocationTieredWithMinimumPriceConversionRateConfigConversionRateTypeTiered PlanNewParamsPricesLicenseAllocationPriceNewLicenseAllocationTieredWithMinimumPriceConversionRateConfigConversionRateType = "tiered"
)

func (r PlanNewParamsPricesLicenseAllocationPriceNewLicenseAllocationTieredWithMinimumPriceConversionRateConfigConversionRateType) IsKnown() bool {
	switch r {
	case PlanNewParamsPricesLicenseAllocationPriceNewLicenseAllocationTieredWithMinimumPriceConversionRateConfigConversionRateTypeUnit, PlanNewParamsPricesLicenseAllocationPriceNewLicenseAllocationTieredWithMinimumPriceConversionRateConfigConversionRateTypeTiered:
		return true
	}
	return false
}

type PlanNewParamsPricesLicenseAllocationPriceNewLicenseAllocationGroupedTieredPrice struct {
	// The cadence to bill for this price on.
	Cadence param.Field[PlanNewParamsPricesLicenseAllocationPriceNewLicenseAllocationGroupedTieredPriceCadence] `json:"cadence" api:"required"`
	// Configuration for grouped_tiered pricing
	GroupedTieredConfig param.Field[PlanNewParamsPricesLicenseAllocationPriceNewLicenseAllocationGroupedTieredPriceGroupedTieredConfig] `json:"grouped_tiered_config" api:"required"`
	// The id of the item the price will be associated with.
	ItemID param.Field[string] `json:"item_id" api:"required"`
	// License allocations to associate with this price. Each entry defines a
	// per-license credit pool granted each cadence. Requires license_type_id or
	// license_type_configuration to be set.
	LicenseAllocations param.Field[[]PlanNewParamsPricesLicenseAllocationPriceNewLicenseAllocationGroupedTieredPriceLicenseAllocation] `json:"license_allocations" api:"required"`
	// The pricing model type
	ModelType param.Field[PlanNewParamsPricesLicenseAllocationPriceNewLicenseAllocationGroupedTieredPriceModelType] `json:"model_type" api:"required"`
	// The name of the price.
	Name param.Field[string] `json:"name" api:"required"`
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
	ConversionRateConfig param.Field[PlanNewParamsPricesLicenseAllocationPriceNewLicenseAllocationGroupedTieredPriceConversionRateConfigUnion] `json:"conversion_rate_config"`
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
	// The ID of the license type to associate with this price.
	LicenseTypeID param.Field[string] `json:"license_type_id"`
	// User-specified key/value pairs for the resource. Individual keys can be removed
	// by setting the value to `null`, and the entire metadata mapping can be cleared
	// by setting `metadata` to `null`.
	Metadata param.Field[map[string]string] `json:"metadata"`
	// A transient ID that can be used to reference this price when adding adjustments
	// in the same API call.
	ReferenceID param.Field[string] `json:"reference_id"`
}

func (r PlanNewParamsPricesLicenseAllocationPriceNewLicenseAllocationGroupedTieredPrice) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r PlanNewParamsPricesLicenseAllocationPriceNewLicenseAllocationGroupedTieredPrice) implementsPlanNewParamsPricesLicenseAllocationPriceUnion() {
}

// The cadence to bill for this price on.
type PlanNewParamsPricesLicenseAllocationPriceNewLicenseAllocationGroupedTieredPriceCadence string

const (
	PlanNewParamsPricesLicenseAllocationPriceNewLicenseAllocationGroupedTieredPriceCadenceAnnual     PlanNewParamsPricesLicenseAllocationPriceNewLicenseAllocationGroupedTieredPriceCadence = "annual"
	PlanNewParamsPricesLicenseAllocationPriceNewLicenseAllocationGroupedTieredPriceCadenceSemiAnnual PlanNewParamsPricesLicenseAllocationPriceNewLicenseAllocationGroupedTieredPriceCadence = "semi_annual"
	PlanNewParamsPricesLicenseAllocationPriceNewLicenseAllocationGroupedTieredPriceCadenceMonthly    PlanNewParamsPricesLicenseAllocationPriceNewLicenseAllocationGroupedTieredPriceCadence = "monthly"
	PlanNewParamsPricesLicenseAllocationPriceNewLicenseAllocationGroupedTieredPriceCadenceQuarterly  PlanNewParamsPricesLicenseAllocationPriceNewLicenseAllocationGroupedTieredPriceCadence = "quarterly"
	PlanNewParamsPricesLicenseAllocationPriceNewLicenseAllocationGroupedTieredPriceCadenceOneTime    PlanNewParamsPricesLicenseAllocationPriceNewLicenseAllocationGroupedTieredPriceCadence = "one_time"
	PlanNewParamsPricesLicenseAllocationPriceNewLicenseAllocationGroupedTieredPriceCadenceCustom     PlanNewParamsPricesLicenseAllocationPriceNewLicenseAllocationGroupedTieredPriceCadence = "custom"
)

func (r PlanNewParamsPricesLicenseAllocationPriceNewLicenseAllocationGroupedTieredPriceCadence) IsKnown() bool {
	switch r {
	case PlanNewParamsPricesLicenseAllocationPriceNewLicenseAllocationGroupedTieredPriceCadenceAnnual, PlanNewParamsPricesLicenseAllocationPriceNewLicenseAllocationGroupedTieredPriceCadenceSemiAnnual, PlanNewParamsPricesLicenseAllocationPriceNewLicenseAllocationGroupedTieredPriceCadenceMonthly, PlanNewParamsPricesLicenseAllocationPriceNewLicenseAllocationGroupedTieredPriceCadenceQuarterly, PlanNewParamsPricesLicenseAllocationPriceNewLicenseAllocationGroupedTieredPriceCadenceOneTime, PlanNewParamsPricesLicenseAllocationPriceNewLicenseAllocationGroupedTieredPriceCadenceCustom:
		return true
	}
	return false
}

// Configuration for grouped_tiered pricing
type PlanNewParamsPricesLicenseAllocationPriceNewLicenseAllocationGroupedTieredPriceGroupedTieredConfig struct {
	// The billable metric property used to group before tiering
	GroupingKey param.Field[string] `json:"grouping_key" api:"required"`
	// Apply tiered pricing to each segment generated after grouping with the provided
	// key
	Tiers param.Field[[]PlanNewParamsPricesLicenseAllocationPriceNewLicenseAllocationGroupedTieredPriceGroupedTieredConfigTier] `json:"tiers" api:"required"`
}

func (r PlanNewParamsPricesLicenseAllocationPriceNewLicenseAllocationGroupedTieredPriceGroupedTieredConfig) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Configuration for a single tier
type PlanNewParamsPricesLicenseAllocationPriceNewLicenseAllocationGroupedTieredPriceGroupedTieredConfigTier struct {
	TierLowerBound param.Field[string] `json:"tier_lower_bound" api:"required"`
	// Per unit amount
	UnitAmount param.Field[string] `json:"unit_amount" api:"required"`
}

func (r PlanNewParamsPricesLicenseAllocationPriceNewLicenseAllocationGroupedTieredPriceGroupedTieredConfigTier) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type PlanNewParamsPricesLicenseAllocationPriceNewLicenseAllocationGroupedTieredPriceLicenseAllocation struct {
	// The amount of credits granted per active license per cadence.
	Amount param.Field[string] `json:"amount" api:"required"`
	// The currency of the license allocation.
	Currency param.Field[string] `json:"currency" api:"required"`
	// When True, overage beyond the allocation is written off.
	WriteOffOverage param.Field[bool] `json:"write_off_overage"`
}

func (r PlanNewParamsPricesLicenseAllocationPriceNewLicenseAllocationGroupedTieredPriceLicenseAllocation) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// The pricing model type
type PlanNewParamsPricesLicenseAllocationPriceNewLicenseAllocationGroupedTieredPriceModelType string

const (
	PlanNewParamsPricesLicenseAllocationPriceNewLicenseAllocationGroupedTieredPriceModelTypeGroupedTiered PlanNewParamsPricesLicenseAllocationPriceNewLicenseAllocationGroupedTieredPriceModelType = "grouped_tiered"
)

func (r PlanNewParamsPricesLicenseAllocationPriceNewLicenseAllocationGroupedTieredPriceModelType) IsKnown() bool {
	switch r {
	case PlanNewParamsPricesLicenseAllocationPriceNewLicenseAllocationGroupedTieredPriceModelTypeGroupedTiered:
		return true
	}
	return false
}

type PlanNewParamsPricesLicenseAllocationPriceNewLicenseAllocationGroupedTieredPriceConversionRateConfig struct {
	ConversionRateType param.Field[PlanNewParamsPricesLicenseAllocationPriceNewLicenseAllocationGroupedTieredPriceConversionRateConfigConversionRateType] `json:"conversion_rate_type" api:"required"`
	TieredConfig       param.Field[shared.ConversionRateTieredConfigParam]                                                                                `json:"tiered_config"`
	UnitConfig         param.Field[shared.ConversionRateUnitConfigParam]                                                                                  `json:"unit_config"`
}

func (r PlanNewParamsPricesLicenseAllocationPriceNewLicenseAllocationGroupedTieredPriceConversionRateConfig) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r PlanNewParamsPricesLicenseAllocationPriceNewLicenseAllocationGroupedTieredPriceConversionRateConfig) ImplementsPlanNewParamsPricesLicenseAllocationPriceNewLicenseAllocationGroupedTieredPriceConversionRateConfigUnion() {
}

// Satisfied by [shared.UnitConversionRateConfigParam],
// [shared.TieredConversionRateConfigParam],
// [PlanNewParamsPricesLicenseAllocationPriceNewLicenseAllocationGroupedTieredPriceConversionRateConfig].
type PlanNewParamsPricesLicenseAllocationPriceNewLicenseAllocationGroupedTieredPriceConversionRateConfigUnion interface {
	ImplementsPlanNewParamsPricesLicenseAllocationPriceNewLicenseAllocationGroupedTieredPriceConversionRateConfigUnion()
}

type PlanNewParamsPricesLicenseAllocationPriceNewLicenseAllocationGroupedTieredPriceConversionRateConfigConversionRateType string

const (
	PlanNewParamsPricesLicenseAllocationPriceNewLicenseAllocationGroupedTieredPriceConversionRateConfigConversionRateTypeUnit   PlanNewParamsPricesLicenseAllocationPriceNewLicenseAllocationGroupedTieredPriceConversionRateConfigConversionRateType = "unit"
	PlanNewParamsPricesLicenseAllocationPriceNewLicenseAllocationGroupedTieredPriceConversionRateConfigConversionRateTypeTiered PlanNewParamsPricesLicenseAllocationPriceNewLicenseAllocationGroupedTieredPriceConversionRateConfigConversionRateType = "tiered"
)

func (r PlanNewParamsPricesLicenseAllocationPriceNewLicenseAllocationGroupedTieredPriceConversionRateConfigConversionRateType) IsKnown() bool {
	switch r {
	case PlanNewParamsPricesLicenseAllocationPriceNewLicenseAllocationGroupedTieredPriceConversionRateConfigConversionRateTypeUnit, PlanNewParamsPricesLicenseAllocationPriceNewLicenseAllocationGroupedTieredPriceConversionRateConfigConversionRateTypeTiered:
		return true
	}
	return false
}

type PlanNewParamsPricesLicenseAllocationPriceNewLicenseAllocationTieredPackageWithMinimumPrice struct {
	// The cadence to bill for this price on.
	Cadence param.Field[PlanNewParamsPricesLicenseAllocationPriceNewLicenseAllocationTieredPackageWithMinimumPriceCadence] `json:"cadence" api:"required"`
	// The id of the item the price will be associated with.
	ItemID param.Field[string] `json:"item_id" api:"required"`
	// License allocations to associate with this price. Each entry defines a
	// per-license credit pool granted each cadence. Requires license_type_id or
	// license_type_configuration to be set.
	LicenseAllocations param.Field[[]PlanNewParamsPricesLicenseAllocationPriceNewLicenseAllocationTieredPackageWithMinimumPriceLicenseAllocation] `json:"license_allocations" api:"required"`
	// The pricing model type
	ModelType param.Field[PlanNewParamsPricesLicenseAllocationPriceNewLicenseAllocationTieredPackageWithMinimumPriceModelType] `json:"model_type" api:"required"`
	// The name of the price.
	Name param.Field[string] `json:"name" api:"required"`
	// Configuration for tiered_package_with_minimum pricing
	TieredPackageWithMinimumConfig param.Field[PlanNewParamsPricesLicenseAllocationPriceNewLicenseAllocationTieredPackageWithMinimumPriceTieredPackageWithMinimumConfig] `json:"tiered_package_with_minimum_config" api:"required"`
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
	ConversionRateConfig param.Field[PlanNewParamsPricesLicenseAllocationPriceNewLicenseAllocationTieredPackageWithMinimumPriceConversionRateConfigUnion] `json:"conversion_rate_config"`
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
	// The ID of the license type to associate with this price.
	LicenseTypeID param.Field[string] `json:"license_type_id"`
	// User-specified key/value pairs for the resource. Individual keys can be removed
	// by setting the value to `null`, and the entire metadata mapping can be cleared
	// by setting `metadata` to `null`.
	Metadata param.Field[map[string]string] `json:"metadata"`
	// A transient ID that can be used to reference this price when adding adjustments
	// in the same API call.
	ReferenceID param.Field[string] `json:"reference_id"`
}

func (r PlanNewParamsPricesLicenseAllocationPriceNewLicenseAllocationTieredPackageWithMinimumPrice) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r PlanNewParamsPricesLicenseAllocationPriceNewLicenseAllocationTieredPackageWithMinimumPrice) implementsPlanNewParamsPricesLicenseAllocationPriceUnion() {
}

// The cadence to bill for this price on.
type PlanNewParamsPricesLicenseAllocationPriceNewLicenseAllocationTieredPackageWithMinimumPriceCadence string

const (
	PlanNewParamsPricesLicenseAllocationPriceNewLicenseAllocationTieredPackageWithMinimumPriceCadenceAnnual     PlanNewParamsPricesLicenseAllocationPriceNewLicenseAllocationTieredPackageWithMinimumPriceCadence = "annual"
	PlanNewParamsPricesLicenseAllocationPriceNewLicenseAllocationTieredPackageWithMinimumPriceCadenceSemiAnnual PlanNewParamsPricesLicenseAllocationPriceNewLicenseAllocationTieredPackageWithMinimumPriceCadence = "semi_annual"
	PlanNewParamsPricesLicenseAllocationPriceNewLicenseAllocationTieredPackageWithMinimumPriceCadenceMonthly    PlanNewParamsPricesLicenseAllocationPriceNewLicenseAllocationTieredPackageWithMinimumPriceCadence = "monthly"
	PlanNewParamsPricesLicenseAllocationPriceNewLicenseAllocationTieredPackageWithMinimumPriceCadenceQuarterly  PlanNewParamsPricesLicenseAllocationPriceNewLicenseAllocationTieredPackageWithMinimumPriceCadence = "quarterly"
	PlanNewParamsPricesLicenseAllocationPriceNewLicenseAllocationTieredPackageWithMinimumPriceCadenceOneTime    PlanNewParamsPricesLicenseAllocationPriceNewLicenseAllocationTieredPackageWithMinimumPriceCadence = "one_time"
	PlanNewParamsPricesLicenseAllocationPriceNewLicenseAllocationTieredPackageWithMinimumPriceCadenceCustom     PlanNewParamsPricesLicenseAllocationPriceNewLicenseAllocationTieredPackageWithMinimumPriceCadence = "custom"
)

func (r PlanNewParamsPricesLicenseAllocationPriceNewLicenseAllocationTieredPackageWithMinimumPriceCadence) IsKnown() bool {
	switch r {
	case PlanNewParamsPricesLicenseAllocationPriceNewLicenseAllocationTieredPackageWithMinimumPriceCadenceAnnual, PlanNewParamsPricesLicenseAllocationPriceNewLicenseAllocationTieredPackageWithMinimumPriceCadenceSemiAnnual, PlanNewParamsPricesLicenseAllocationPriceNewLicenseAllocationTieredPackageWithMinimumPriceCadenceMonthly, PlanNewParamsPricesLicenseAllocationPriceNewLicenseAllocationTieredPackageWithMinimumPriceCadenceQuarterly, PlanNewParamsPricesLicenseAllocationPriceNewLicenseAllocationTieredPackageWithMinimumPriceCadenceOneTime, PlanNewParamsPricesLicenseAllocationPriceNewLicenseAllocationTieredPackageWithMinimumPriceCadenceCustom:
		return true
	}
	return false
}

type PlanNewParamsPricesLicenseAllocationPriceNewLicenseAllocationTieredPackageWithMinimumPriceLicenseAllocation struct {
	// The amount of credits granted per active license per cadence.
	Amount param.Field[string] `json:"amount" api:"required"`
	// The currency of the license allocation.
	Currency param.Field[string] `json:"currency" api:"required"`
	// When True, overage beyond the allocation is written off.
	WriteOffOverage param.Field[bool] `json:"write_off_overage"`
}

func (r PlanNewParamsPricesLicenseAllocationPriceNewLicenseAllocationTieredPackageWithMinimumPriceLicenseAllocation) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// The pricing model type
type PlanNewParamsPricesLicenseAllocationPriceNewLicenseAllocationTieredPackageWithMinimumPriceModelType string

const (
	PlanNewParamsPricesLicenseAllocationPriceNewLicenseAllocationTieredPackageWithMinimumPriceModelTypeTieredPackageWithMinimum PlanNewParamsPricesLicenseAllocationPriceNewLicenseAllocationTieredPackageWithMinimumPriceModelType = "tiered_package_with_minimum"
)

func (r PlanNewParamsPricesLicenseAllocationPriceNewLicenseAllocationTieredPackageWithMinimumPriceModelType) IsKnown() bool {
	switch r {
	case PlanNewParamsPricesLicenseAllocationPriceNewLicenseAllocationTieredPackageWithMinimumPriceModelTypeTieredPackageWithMinimum:
		return true
	}
	return false
}

// Configuration for tiered_package_with_minimum pricing
type PlanNewParamsPricesLicenseAllocationPriceNewLicenseAllocationTieredPackageWithMinimumPriceTieredPackageWithMinimumConfig struct {
	PackageSize param.Field[float64] `json:"package_size" api:"required"`
	// Apply tiered pricing after rounding up the quantity to the package size. Tiers
	// are defined using exclusive lower bounds.
	Tiers param.Field[[]PlanNewParamsPricesLicenseAllocationPriceNewLicenseAllocationTieredPackageWithMinimumPriceTieredPackageWithMinimumConfigTier] `json:"tiers" api:"required"`
}

func (r PlanNewParamsPricesLicenseAllocationPriceNewLicenseAllocationTieredPackageWithMinimumPriceTieredPackageWithMinimumConfig) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Configuration for a single tier
type PlanNewParamsPricesLicenseAllocationPriceNewLicenseAllocationTieredPackageWithMinimumPriceTieredPackageWithMinimumConfigTier struct {
	MinimumAmount  param.Field[string] `json:"minimum_amount" api:"required"`
	PerUnit        param.Field[string] `json:"per_unit" api:"required"`
	TierLowerBound param.Field[string] `json:"tier_lower_bound" api:"required"`
}

func (r PlanNewParamsPricesLicenseAllocationPriceNewLicenseAllocationTieredPackageWithMinimumPriceTieredPackageWithMinimumConfigTier) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type PlanNewParamsPricesLicenseAllocationPriceNewLicenseAllocationTieredPackageWithMinimumPriceConversionRateConfig struct {
	ConversionRateType param.Field[PlanNewParamsPricesLicenseAllocationPriceNewLicenseAllocationTieredPackageWithMinimumPriceConversionRateConfigConversionRateType] `json:"conversion_rate_type" api:"required"`
	TieredConfig       param.Field[shared.ConversionRateTieredConfigParam]                                                                                           `json:"tiered_config"`
	UnitConfig         param.Field[shared.ConversionRateUnitConfigParam]                                                                                             `json:"unit_config"`
}

func (r PlanNewParamsPricesLicenseAllocationPriceNewLicenseAllocationTieredPackageWithMinimumPriceConversionRateConfig) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r PlanNewParamsPricesLicenseAllocationPriceNewLicenseAllocationTieredPackageWithMinimumPriceConversionRateConfig) ImplementsPlanNewParamsPricesLicenseAllocationPriceNewLicenseAllocationTieredPackageWithMinimumPriceConversionRateConfigUnion() {
}

// Satisfied by [shared.UnitConversionRateConfigParam],
// [shared.TieredConversionRateConfigParam],
// [PlanNewParamsPricesLicenseAllocationPriceNewLicenseAllocationTieredPackageWithMinimumPriceConversionRateConfig].
type PlanNewParamsPricesLicenseAllocationPriceNewLicenseAllocationTieredPackageWithMinimumPriceConversionRateConfigUnion interface {
	ImplementsPlanNewParamsPricesLicenseAllocationPriceNewLicenseAllocationTieredPackageWithMinimumPriceConversionRateConfigUnion()
}

type PlanNewParamsPricesLicenseAllocationPriceNewLicenseAllocationTieredPackageWithMinimumPriceConversionRateConfigConversionRateType string

const (
	PlanNewParamsPricesLicenseAllocationPriceNewLicenseAllocationTieredPackageWithMinimumPriceConversionRateConfigConversionRateTypeUnit   PlanNewParamsPricesLicenseAllocationPriceNewLicenseAllocationTieredPackageWithMinimumPriceConversionRateConfigConversionRateType = "unit"
	PlanNewParamsPricesLicenseAllocationPriceNewLicenseAllocationTieredPackageWithMinimumPriceConversionRateConfigConversionRateTypeTiered PlanNewParamsPricesLicenseAllocationPriceNewLicenseAllocationTieredPackageWithMinimumPriceConversionRateConfigConversionRateType = "tiered"
)

func (r PlanNewParamsPricesLicenseAllocationPriceNewLicenseAllocationTieredPackageWithMinimumPriceConversionRateConfigConversionRateType) IsKnown() bool {
	switch r {
	case PlanNewParamsPricesLicenseAllocationPriceNewLicenseAllocationTieredPackageWithMinimumPriceConversionRateConfigConversionRateTypeUnit, PlanNewParamsPricesLicenseAllocationPriceNewLicenseAllocationTieredPackageWithMinimumPriceConversionRateConfigConversionRateTypeTiered:
		return true
	}
	return false
}

type PlanNewParamsPricesLicenseAllocationPriceNewLicenseAllocationPackageWithAllocationPrice struct {
	// The cadence to bill for this price on.
	Cadence param.Field[PlanNewParamsPricesLicenseAllocationPriceNewLicenseAllocationPackageWithAllocationPriceCadence] `json:"cadence" api:"required"`
	// The id of the item the price will be associated with.
	ItemID param.Field[string] `json:"item_id" api:"required"`
	// License allocations to associate with this price. Each entry defines a
	// per-license credit pool granted each cadence. Requires license_type_id or
	// license_type_configuration to be set.
	LicenseAllocations param.Field[[]PlanNewParamsPricesLicenseAllocationPriceNewLicenseAllocationPackageWithAllocationPriceLicenseAllocation] `json:"license_allocations" api:"required"`
	// The pricing model type
	ModelType param.Field[PlanNewParamsPricesLicenseAllocationPriceNewLicenseAllocationPackageWithAllocationPriceModelType] `json:"model_type" api:"required"`
	// The name of the price.
	Name param.Field[string] `json:"name" api:"required"`
	// Configuration for package_with_allocation pricing
	PackageWithAllocationConfig param.Field[PlanNewParamsPricesLicenseAllocationPriceNewLicenseAllocationPackageWithAllocationPricePackageWithAllocationConfig] `json:"package_with_allocation_config" api:"required"`
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
	ConversionRateConfig param.Field[PlanNewParamsPricesLicenseAllocationPriceNewLicenseAllocationPackageWithAllocationPriceConversionRateConfigUnion] `json:"conversion_rate_config"`
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
	// The ID of the license type to associate with this price.
	LicenseTypeID param.Field[string] `json:"license_type_id"`
	// User-specified key/value pairs for the resource. Individual keys can be removed
	// by setting the value to `null`, and the entire metadata mapping can be cleared
	// by setting `metadata` to `null`.
	Metadata param.Field[map[string]string] `json:"metadata"`
	// A transient ID that can be used to reference this price when adding adjustments
	// in the same API call.
	ReferenceID param.Field[string] `json:"reference_id"`
}

func (r PlanNewParamsPricesLicenseAllocationPriceNewLicenseAllocationPackageWithAllocationPrice) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r PlanNewParamsPricesLicenseAllocationPriceNewLicenseAllocationPackageWithAllocationPrice) implementsPlanNewParamsPricesLicenseAllocationPriceUnion() {
}

// The cadence to bill for this price on.
type PlanNewParamsPricesLicenseAllocationPriceNewLicenseAllocationPackageWithAllocationPriceCadence string

const (
	PlanNewParamsPricesLicenseAllocationPriceNewLicenseAllocationPackageWithAllocationPriceCadenceAnnual     PlanNewParamsPricesLicenseAllocationPriceNewLicenseAllocationPackageWithAllocationPriceCadence = "annual"
	PlanNewParamsPricesLicenseAllocationPriceNewLicenseAllocationPackageWithAllocationPriceCadenceSemiAnnual PlanNewParamsPricesLicenseAllocationPriceNewLicenseAllocationPackageWithAllocationPriceCadence = "semi_annual"
	PlanNewParamsPricesLicenseAllocationPriceNewLicenseAllocationPackageWithAllocationPriceCadenceMonthly    PlanNewParamsPricesLicenseAllocationPriceNewLicenseAllocationPackageWithAllocationPriceCadence = "monthly"
	PlanNewParamsPricesLicenseAllocationPriceNewLicenseAllocationPackageWithAllocationPriceCadenceQuarterly  PlanNewParamsPricesLicenseAllocationPriceNewLicenseAllocationPackageWithAllocationPriceCadence = "quarterly"
	PlanNewParamsPricesLicenseAllocationPriceNewLicenseAllocationPackageWithAllocationPriceCadenceOneTime    PlanNewParamsPricesLicenseAllocationPriceNewLicenseAllocationPackageWithAllocationPriceCadence = "one_time"
	PlanNewParamsPricesLicenseAllocationPriceNewLicenseAllocationPackageWithAllocationPriceCadenceCustom     PlanNewParamsPricesLicenseAllocationPriceNewLicenseAllocationPackageWithAllocationPriceCadence = "custom"
)

func (r PlanNewParamsPricesLicenseAllocationPriceNewLicenseAllocationPackageWithAllocationPriceCadence) IsKnown() bool {
	switch r {
	case PlanNewParamsPricesLicenseAllocationPriceNewLicenseAllocationPackageWithAllocationPriceCadenceAnnual, PlanNewParamsPricesLicenseAllocationPriceNewLicenseAllocationPackageWithAllocationPriceCadenceSemiAnnual, PlanNewParamsPricesLicenseAllocationPriceNewLicenseAllocationPackageWithAllocationPriceCadenceMonthly, PlanNewParamsPricesLicenseAllocationPriceNewLicenseAllocationPackageWithAllocationPriceCadenceQuarterly, PlanNewParamsPricesLicenseAllocationPriceNewLicenseAllocationPackageWithAllocationPriceCadenceOneTime, PlanNewParamsPricesLicenseAllocationPriceNewLicenseAllocationPackageWithAllocationPriceCadenceCustom:
		return true
	}
	return false
}

type PlanNewParamsPricesLicenseAllocationPriceNewLicenseAllocationPackageWithAllocationPriceLicenseAllocation struct {
	// The amount of credits granted per active license per cadence.
	Amount param.Field[string] `json:"amount" api:"required"`
	// The currency of the license allocation.
	Currency param.Field[string] `json:"currency" api:"required"`
	// When True, overage beyond the allocation is written off.
	WriteOffOverage param.Field[bool] `json:"write_off_overage"`
}

func (r PlanNewParamsPricesLicenseAllocationPriceNewLicenseAllocationPackageWithAllocationPriceLicenseAllocation) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// The pricing model type
type PlanNewParamsPricesLicenseAllocationPriceNewLicenseAllocationPackageWithAllocationPriceModelType string

const (
	PlanNewParamsPricesLicenseAllocationPriceNewLicenseAllocationPackageWithAllocationPriceModelTypePackageWithAllocation PlanNewParamsPricesLicenseAllocationPriceNewLicenseAllocationPackageWithAllocationPriceModelType = "package_with_allocation"
)

func (r PlanNewParamsPricesLicenseAllocationPriceNewLicenseAllocationPackageWithAllocationPriceModelType) IsKnown() bool {
	switch r {
	case PlanNewParamsPricesLicenseAllocationPriceNewLicenseAllocationPackageWithAllocationPriceModelTypePackageWithAllocation:
		return true
	}
	return false
}

// Configuration for package_with_allocation pricing
type PlanNewParamsPricesLicenseAllocationPriceNewLicenseAllocationPackageWithAllocationPricePackageWithAllocationConfig struct {
	Allocation    param.Field[string] `json:"allocation" api:"required"`
	PackageAmount param.Field[string] `json:"package_amount" api:"required"`
	PackageSize   param.Field[string] `json:"package_size" api:"required"`
}

func (r PlanNewParamsPricesLicenseAllocationPriceNewLicenseAllocationPackageWithAllocationPricePackageWithAllocationConfig) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type PlanNewParamsPricesLicenseAllocationPriceNewLicenseAllocationPackageWithAllocationPriceConversionRateConfig struct {
	ConversionRateType param.Field[PlanNewParamsPricesLicenseAllocationPriceNewLicenseAllocationPackageWithAllocationPriceConversionRateConfigConversionRateType] `json:"conversion_rate_type" api:"required"`
	TieredConfig       param.Field[shared.ConversionRateTieredConfigParam]                                                                                        `json:"tiered_config"`
	UnitConfig         param.Field[shared.ConversionRateUnitConfigParam]                                                                                          `json:"unit_config"`
}

func (r PlanNewParamsPricesLicenseAllocationPriceNewLicenseAllocationPackageWithAllocationPriceConversionRateConfig) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r PlanNewParamsPricesLicenseAllocationPriceNewLicenseAllocationPackageWithAllocationPriceConversionRateConfig) ImplementsPlanNewParamsPricesLicenseAllocationPriceNewLicenseAllocationPackageWithAllocationPriceConversionRateConfigUnion() {
}

// Satisfied by [shared.UnitConversionRateConfigParam],
// [shared.TieredConversionRateConfigParam],
// [PlanNewParamsPricesLicenseAllocationPriceNewLicenseAllocationPackageWithAllocationPriceConversionRateConfig].
type PlanNewParamsPricesLicenseAllocationPriceNewLicenseAllocationPackageWithAllocationPriceConversionRateConfigUnion interface {
	ImplementsPlanNewParamsPricesLicenseAllocationPriceNewLicenseAllocationPackageWithAllocationPriceConversionRateConfigUnion()
}

type PlanNewParamsPricesLicenseAllocationPriceNewLicenseAllocationPackageWithAllocationPriceConversionRateConfigConversionRateType string

const (
	PlanNewParamsPricesLicenseAllocationPriceNewLicenseAllocationPackageWithAllocationPriceConversionRateConfigConversionRateTypeUnit   PlanNewParamsPricesLicenseAllocationPriceNewLicenseAllocationPackageWithAllocationPriceConversionRateConfigConversionRateType = "unit"
	PlanNewParamsPricesLicenseAllocationPriceNewLicenseAllocationPackageWithAllocationPriceConversionRateConfigConversionRateTypeTiered PlanNewParamsPricesLicenseAllocationPriceNewLicenseAllocationPackageWithAllocationPriceConversionRateConfigConversionRateType = "tiered"
)

func (r PlanNewParamsPricesLicenseAllocationPriceNewLicenseAllocationPackageWithAllocationPriceConversionRateConfigConversionRateType) IsKnown() bool {
	switch r {
	case PlanNewParamsPricesLicenseAllocationPriceNewLicenseAllocationPackageWithAllocationPriceConversionRateConfigConversionRateTypeUnit, PlanNewParamsPricesLicenseAllocationPriceNewLicenseAllocationPackageWithAllocationPriceConversionRateConfigConversionRateTypeTiered:
		return true
	}
	return false
}

type PlanNewParamsPricesLicenseAllocationPriceNewLicenseAllocationUnitWithPercentPrice struct {
	// The cadence to bill for this price on.
	Cadence param.Field[PlanNewParamsPricesLicenseAllocationPriceNewLicenseAllocationUnitWithPercentPriceCadence] `json:"cadence" api:"required"`
	// The id of the item the price will be associated with.
	ItemID param.Field[string] `json:"item_id" api:"required"`
	// License allocations to associate with this price. Each entry defines a
	// per-license credit pool granted each cadence. Requires license_type_id or
	// license_type_configuration to be set.
	LicenseAllocations param.Field[[]PlanNewParamsPricesLicenseAllocationPriceNewLicenseAllocationUnitWithPercentPriceLicenseAllocation] `json:"license_allocations" api:"required"`
	// The pricing model type
	ModelType param.Field[PlanNewParamsPricesLicenseAllocationPriceNewLicenseAllocationUnitWithPercentPriceModelType] `json:"model_type" api:"required"`
	// The name of the price.
	Name param.Field[string] `json:"name" api:"required"`
	// Configuration for unit_with_percent pricing
	UnitWithPercentConfig param.Field[PlanNewParamsPricesLicenseAllocationPriceNewLicenseAllocationUnitWithPercentPriceUnitWithPercentConfig] `json:"unit_with_percent_config" api:"required"`
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
	ConversionRateConfig param.Field[PlanNewParamsPricesLicenseAllocationPriceNewLicenseAllocationUnitWithPercentPriceConversionRateConfigUnion] `json:"conversion_rate_config"`
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
	// The ID of the license type to associate with this price.
	LicenseTypeID param.Field[string] `json:"license_type_id"`
	// User-specified key/value pairs for the resource. Individual keys can be removed
	// by setting the value to `null`, and the entire metadata mapping can be cleared
	// by setting `metadata` to `null`.
	Metadata param.Field[map[string]string] `json:"metadata"`
	// A transient ID that can be used to reference this price when adding adjustments
	// in the same API call.
	ReferenceID param.Field[string] `json:"reference_id"`
}

func (r PlanNewParamsPricesLicenseAllocationPriceNewLicenseAllocationUnitWithPercentPrice) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r PlanNewParamsPricesLicenseAllocationPriceNewLicenseAllocationUnitWithPercentPrice) implementsPlanNewParamsPricesLicenseAllocationPriceUnion() {
}

// The cadence to bill for this price on.
type PlanNewParamsPricesLicenseAllocationPriceNewLicenseAllocationUnitWithPercentPriceCadence string

const (
	PlanNewParamsPricesLicenseAllocationPriceNewLicenseAllocationUnitWithPercentPriceCadenceAnnual     PlanNewParamsPricesLicenseAllocationPriceNewLicenseAllocationUnitWithPercentPriceCadence = "annual"
	PlanNewParamsPricesLicenseAllocationPriceNewLicenseAllocationUnitWithPercentPriceCadenceSemiAnnual PlanNewParamsPricesLicenseAllocationPriceNewLicenseAllocationUnitWithPercentPriceCadence = "semi_annual"
	PlanNewParamsPricesLicenseAllocationPriceNewLicenseAllocationUnitWithPercentPriceCadenceMonthly    PlanNewParamsPricesLicenseAllocationPriceNewLicenseAllocationUnitWithPercentPriceCadence = "monthly"
	PlanNewParamsPricesLicenseAllocationPriceNewLicenseAllocationUnitWithPercentPriceCadenceQuarterly  PlanNewParamsPricesLicenseAllocationPriceNewLicenseAllocationUnitWithPercentPriceCadence = "quarterly"
	PlanNewParamsPricesLicenseAllocationPriceNewLicenseAllocationUnitWithPercentPriceCadenceOneTime    PlanNewParamsPricesLicenseAllocationPriceNewLicenseAllocationUnitWithPercentPriceCadence = "one_time"
	PlanNewParamsPricesLicenseAllocationPriceNewLicenseAllocationUnitWithPercentPriceCadenceCustom     PlanNewParamsPricesLicenseAllocationPriceNewLicenseAllocationUnitWithPercentPriceCadence = "custom"
)

func (r PlanNewParamsPricesLicenseAllocationPriceNewLicenseAllocationUnitWithPercentPriceCadence) IsKnown() bool {
	switch r {
	case PlanNewParamsPricesLicenseAllocationPriceNewLicenseAllocationUnitWithPercentPriceCadenceAnnual, PlanNewParamsPricesLicenseAllocationPriceNewLicenseAllocationUnitWithPercentPriceCadenceSemiAnnual, PlanNewParamsPricesLicenseAllocationPriceNewLicenseAllocationUnitWithPercentPriceCadenceMonthly, PlanNewParamsPricesLicenseAllocationPriceNewLicenseAllocationUnitWithPercentPriceCadenceQuarterly, PlanNewParamsPricesLicenseAllocationPriceNewLicenseAllocationUnitWithPercentPriceCadenceOneTime, PlanNewParamsPricesLicenseAllocationPriceNewLicenseAllocationUnitWithPercentPriceCadenceCustom:
		return true
	}
	return false
}

type PlanNewParamsPricesLicenseAllocationPriceNewLicenseAllocationUnitWithPercentPriceLicenseAllocation struct {
	// The amount of credits granted per active license per cadence.
	Amount param.Field[string] `json:"amount" api:"required"`
	// The currency of the license allocation.
	Currency param.Field[string] `json:"currency" api:"required"`
	// When True, overage beyond the allocation is written off.
	WriteOffOverage param.Field[bool] `json:"write_off_overage"`
}

func (r PlanNewParamsPricesLicenseAllocationPriceNewLicenseAllocationUnitWithPercentPriceLicenseAllocation) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// The pricing model type
type PlanNewParamsPricesLicenseAllocationPriceNewLicenseAllocationUnitWithPercentPriceModelType string

const (
	PlanNewParamsPricesLicenseAllocationPriceNewLicenseAllocationUnitWithPercentPriceModelTypeUnitWithPercent PlanNewParamsPricesLicenseAllocationPriceNewLicenseAllocationUnitWithPercentPriceModelType = "unit_with_percent"
)

func (r PlanNewParamsPricesLicenseAllocationPriceNewLicenseAllocationUnitWithPercentPriceModelType) IsKnown() bool {
	switch r {
	case PlanNewParamsPricesLicenseAllocationPriceNewLicenseAllocationUnitWithPercentPriceModelTypeUnitWithPercent:
		return true
	}
	return false
}

// Configuration for unit_with_percent pricing
type PlanNewParamsPricesLicenseAllocationPriceNewLicenseAllocationUnitWithPercentPriceUnitWithPercentConfig struct {
	// What percent, out of 100, of the calculated total to charge
	Percent param.Field[string] `json:"percent" api:"required"`
	// Rate per unit of usage
	UnitAmount param.Field[string] `json:"unit_amount" api:"required"`
}

func (r PlanNewParamsPricesLicenseAllocationPriceNewLicenseAllocationUnitWithPercentPriceUnitWithPercentConfig) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type PlanNewParamsPricesLicenseAllocationPriceNewLicenseAllocationUnitWithPercentPriceConversionRateConfig struct {
	ConversionRateType param.Field[PlanNewParamsPricesLicenseAllocationPriceNewLicenseAllocationUnitWithPercentPriceConversionRateConfigConversionRateType] `json:"conversion_rate_type" api:"required"`
	TieredConfig       param.Field[shared.ConversionRateTieredConfigParam]                                                                                  `json:"tiered_config"`
	UnitConfig         param.Field[shared.ConversionRateUnitConfigParam]                                                                                    `json:"unit_config"`
}

func (r PlanNewParamsPricesLicenseAllocationPriceNewLicenseAllocationUnitWithPercentPriceConversionRateConfig) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r PlanNewParamsPricesLicenseAllocationPriceNewLicenseAllocationUnitWithPercentPriceConversionRateConfig) ImplementsPlanNewParamsPricesLicenseAllocationPriceNewLicenseAllocationUnitWithPercentPriceConversionRateConfigUnion() {
}

// Satisfied by [shared.UnitConversionRateConfigParam],
// [shared.TieredConversionRateConfigParam],
// [PlanNewParamsPricesLicenseAllocationPriceNewLicenseAllocationUnitWithPercentPriceConversionRateConfig].
type PlanNewParamsPricesLicenseAllocationPriceNewLicenseAllocationUnitWithPercentPriceConversionRateConfigUnion interface {
	ImplementsPlanNewParamsPricesLicenseAllocationPriceNewLicenseAllocationUnitWithPercentPriceConversionRateConfigUnion()
}

type PlanNewParamsPricesLicenseAllocationPriceNewLicenseAllocationUnitWithPercentPriceConversionRateConfigConversionRateType string

const (
	PlanNewParamsPricesLicenseAllocationPriceNewLicenseAllocationUnitWithPercentPriceConversionRateConfigConversionRateTypeUnit   PlanNewParamsPricesLicenseAllocationPriceNewLicenseAllocationUnitWithPercentPriceConversionRateConfigConversionRateType = "unit"
	PlanNewParamsPricesLicenseAllocationPriceNewLicenseAllocationUnitWithPercentPriceConversionRateConfigConversionRateTypeTiered PlanNewParamsPricesLicenseAllocationPriceNewLicenseAllocationUnitWithPercentPriceConversionRateConfigConversionRateType = "tiered"
)

func (r PlanNewParamsPricesLicenseAllocationPriceNewLicenseAllocationUnitWithPercentPriceConversionRateConfigConversionRateType) IsKnown() bool {
	switch r {
	case PlanNewParamsPricesLicenseAllocationPriceNewLicenseAllocationUnitWithPercentPriceConversionRateConfigConversionRateTypeUnit, PlanNewParamsPricesLicenseAllocationPriceNewLicenseAllocationUnitWithPercentPriceConversionRateConfigConversionRateTypeTiered:
		return true
	}
	return false
}

type PlanNewParamsPricesLicenseAllocationPriceNewLicenseAllocationMatrixWithAllocationPrice struct {
	// The cadence to bill for this price on.
	Cadence param.Field[PlanNewParamsPricesLicenseAllocationPriceNewLicenseAllocationMatrixWithAllocationPriceCadence] `json:"cadence" api:"required"`
	// The id of the item the price will be associated with.
	ItemID param.Field[string] `json:"item_id" api:"required"`
	// License allocations to associate with this price. Each entry defines a
	// per-license credit pool granted each cadence. Requires license_type_id or
	// license_type_configuration to be set.
	LicenseAllocations param.Field[[]PlanNewParamsPricesLicenseAllocationPriceNewLicenseAllocationMatrixWithAllocationPriceLicenseAllocation] `json:"license_allocations" api:"required"`
	// Configuration for matrix_with_allocation pricing
	MatrixWithAllocationConfig param.Field[shared.MatrixWithAllocationConfigParam] `json:"matrix_with_allocation_config" api:"required"`
	// The pricing model type
	ModelType param.Field[PlanNewParamsPricesLicenseAllocationPriceNewLicenseAllocationMatrixWithAllocationPriceModelType] `json:"model_type" api:"required"`
	// The name of the price.
	Name param.Field[string] `json:"name" api:"required"`
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
	ConversionRateConfig param.Field[PlanNewParamsPricesLicenseAllocationPriceNewLicenseAllocationMatrixWithAllocationPriceConversionRateConfigUnion] `json:"conversion_rate_config"`
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
	// The ID of the license type to associate with this price.
	LicenseTypeID param.Field[string] `json:"license_type_id"`
	// User-specified key/value pairs for the resource. Individual keys can be removed
	// by setting the value to `null`, and the entire metadata mapping can be cleared
	// by setting `metadata` to `null`.
	Metadata param.Field[map[string]string] `json:"metadata"`
	// A transient ID that can be used to reference this price when adding adjustments
	// in the same API call.
	ReferenceID param.Field[string] `json:"reference_id"`
}

func (r PlanNewParamsPricesLicenseAllocationPriceNewLicenseAllocationMatrixWithAllocationPrice) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r PlanNewParamsPricesLicenseAllocationPriceNewLicenseAllocationMatrixWithAllocationPrice) implementsPlanNewParamsPricesLicenseAllocationPriceUnion() {
}

// The cadence to bill for this price on.
type PlanNewParamsPricesLicenseAllocationPriceNewLicenseAllocationMatrixWithAllocationPriceCadence string

const (
	PlanNewParamsPricesLicenseAllocationPriceNewLicenseAllocationMatrixWithAllocationPriceCadenceAnnual     PlanNewParamsPricesLicenseAllocationPriceNewLicenseAllocationMatrixWithAllocationPriceCadence = "annual"
	PlanNewParamsPricesLicenseAllocationPriceNewLicenseAllocationMatrixWithAllocationPriceCadenceSemiAnnual PlanNewParamsPricesLicenseAllocationPriceNewLicenseAllocationMatrixWithAllocationPriceCadence = "semi_annual"
	PlanNewParamsPricesLicenseAllocationPriceNewLicenseAllocationMatrixWithAllocationPriceCadenceMonthly    PlanNewParamsPricesLicenseAllocationPriceNewLicenseAllocationMatrixWithAllocationPriceCadence = "monthly"
	PlanNewParamsPricesLicenseAllocationPriceNewLicenseAllocationMatrixWithAllocationPriceCadenceQuarterly  PlanNewParamsPricesLicenseAllocationPriceNewLicenseAllocationMatrixWithAllocationPriceCadence = "quarterly"
	PlanNewParamsPricesLicenseAllocationPriceNewLicenseAllocationMatrixWithAllocationPriceCadenceOneTime    PlanNewParamsPricesLicenseAllocationPriceNewLicenseAllocationMatrixWithAllocationPriceCadence = "one_time"
	PlanNewParamsPricesLicenseAllocationPriceNewLicenseAllocationMatrixWithAllocationPriceCadenceCustom     PlanNewParamsPricesLicenseAllocationPriceNewLicenseAllocationMatrixWithAllocationPriceCadence = "custom"
)

func (r PlanNewParamsPricesLicenseAllocationPriceNewLicenseAllocationMatrixWithAllocationPriceCadence) IsKnown() bool {
	switch r {
	case PlanNewParamsPricesLicenseAllocationPriceNewLicenseAllocationMatrixWithAllocationPriceCadenceAnnual, PlanNewParamsPricesLicenseAllocationPriceNewLicenseAllocationMatrixWithAllocationPriceCadenceSemiAnnual, PlanNewParamsPricesLicenseAllocationPriceNewLicenseAllocationMatrixWithAllocationPriceCadenceMonthly, PlanNewParamsPricesLicenseAllocationPriceNewLicenseAllocationMatrixWithAllocationPriceCadenceQuarterly, PlanNewParamsPricesLicenseAllocationPriceNewLicenseAllocationMatrixWithAllocationPriceCadenceOneTime, PlanNewParamsPricesLicenseAllocationPriceNewLicenseAllocationMatrixWithAllocationPriceCadenceCustom:
		return true
	}
	return false
}

type PlanNewParamsPricesLicenseAllocationPriceNewLicenseAllocationMatrixWithAllocationPriceLicenseAllocation struct {
	// The amount of credits granted per active license per cadence.
	Amount param.Field[string] `json:"amount" api:"required"`
	// The currency of the license allocation.
	Currency param.Field[string] `json:"currency" api:"required"`
	// When True, overage beyond the allocation is written off.
	WriteOffOverage param.Field[bool] `json:"write_off_overage"`
}

func (r PlanNewParamsPricesLicenseAllocationPriceNewLicenseAllocationMatrixWithAllocationPriceLicenseAllocation) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// The pricing model type
type PlanNewParamsPricesLicenseAllocationPriceNewLicenseAllocationMatrixWithAllocationPriceModelType string

const (
	PlanNewParamsPricesLicenseAllocationPriceNewLicenseAllocationMatrixWithAllocationPriceModelTypeMatrixWithAllocation PlanNewParamsPricesLicenseAllocationPriceNewLicenseAllocationMatrixWithAllocationPriceModelType = "matrix_with_allocation"
)

func (r PlanNewParamsPricesLicenseAllocationPriceNewLicenseAllocationMatrixWithAllocationPriceModelType) IsKnown() bool {
	switch r {
	case PlanNewParamsPricesLicenseAllocationPriceNewLicenseAllocationMatrixWithAllocationPriceModelTypeMatrixWithAllocation:
		return true
	}
	return false
}

type PlanNewParamsPricesLicenseAllocationPriceNewLicenseAllocationMatrixWithAllocationPriceConversionRateConfig struct {
	ConversionRateType param.Field[PlanNewParamsPricesLicenseAllocationPriceNewLicenseAllocationMatrixWithAllocationPriceConversionRateConfigConversionRateType] `json:"conversion_rate_type" api:"required"`
	TieredConfig       param.Field[shared.ConversionRateTieredConfigParam]                                                                                       `json:"tiered_config"`
	UnitConfig         param.Field[shared.ConversionRateUnitConfigParam]                                                                                         `json:"unit_config"`
}

func (r PlanNewParamsPricesLicenseAllocationPriceNewLicenseAllocationMatrixWithAllocationPriceConversionRateConfig) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r PlanNewParamsPricesLicenseAllocationPriceNewLicenseAllocationMatrixWithAllocationPriceConversionRateConfig) ImplementsPlanNewParamsPricesLicenseAllocationPriceNewLicenseAllocationMatrixWithAllocationPriceConversionRateConfigUnion() {
}

// Satisfied by [shared.UnitConversionRateConfigParam],
// [shared.TieredConversionRateConfigParam],
// [PlanNewParamsPricesLicenseAllocationPriceNewLicenseAllocationMatrixWithAllocationPriceConversionRateConfig].
type PlanNewParamsPricesLicenseAllocationPriceNewLicenseAllocationMatrixWithAllocationPriceConversionRateConfigUnion interface {
	ImplementsPlanNewParamsPricesLicenseAllocationPriceNewLicenseAllocationMatrixWithAllocationPriceConversionRateConfigUnion()
}

type PlanNewParamsPricesLicenseAllocationPriceNewLicenseAllocationMatrixWithAllocationPriceConversionRateConfigConversionRateType string

const (
	PlanNewParamsPricesLicenseAllocationPriceNewLicenseAllocationMatrixWithAllocationPriceConversionRateConfigConversionRateTypeUnit   PlanNewParamsPricesLicenseAllocationPriceNewLicenseAllocationMatrixWithAllocationPriceConversionRateConfigConversionRateType = "unit"
	PlanNewParamsPricesLicenseAllocationPriceNewLicenseAllocationMatrixWithAllocationPriceConversionRateConfigConversionRateTypeTiered PlanNewParamsPricesLicenseAllocationPriceNewLicenseAllocationMatrixWithAllocationPriceConversionRateConfigConversionRateType = "tiered"
)

func (r PlanNewParamsPricesLicenseAllocationPriceNewLicenseAllocationMatrixWithAllocationPriceConversionRateConfigConversionRateType) IsKnown() bool {
	switch r {
	case PlanNewParamsPricesLicenseAllocationPriceNewLicenseAllocationMatrixWithAllocationPriceConversionRateConfigConversionRateTypeUnit, PlanNewParamsPricesLicenseAllocationPriceNewLicenseAllocationMatrixWithAllocationPriceConversionRateConfigConversionRateTypeTiered:
		return true
	}
	return false
}

type PlanNewParamsPricesLicenseAllocationPriceNewLicenseAllocationTieredWithProrationPrice struct {
	// The cadence to bill for this price on.
	Cadence param.Field[PlanNewParamsPricesLicenseAllocationPriceNewLicenseAllocationTieredWithProrationPriceCadence] `json:"cadence" api:"required"`
	// The id of the item the price will be associated with.
	ItemID param.Field[string] `json:"item_id" api:"required"`
	// License allocations to associate with this price. Each entry defines a
	// per-license credit pool granted each cadence. Requires license_type_id or
	// license_type_configuration to be set.
	LicenseAllocations param.Field[[]PlanNewParamsPricesLicenseAllocationPriceNewLicenseAllocationTieredWithProrationPriceLicenseAllocation] `json:"license_allocations" api:"required"`
	// The pricing model type
	ModelType param.Field[PlanNewParamsPricesLicenseAllocationPriceNewLicenseAllocationTieredWithProrationPriceModelType] `json:"model_type" api:"required"`
	// The name of the price.
	Name param.Field[string] `json:"name" api:"required"`
	// Configuration for tiered_with_proration pricing
	TieredWithProrationConfig param.Field[PlanNewParamsPricesLicenseAllocationPriceNewLicenseAllocationTieredWithProrationPriceTieredWithProrationConfig] `json:"tiered_with_proration_config" api:"required"`
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
	ConversionRateConfig param.Field[PlanNewParamsPricesLicenseAllocationPriceNewLicenseAllocationTieredWithProrationPriceConversionRateConfigUnion] `json:"conversion_rate_config"`
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
	// The ID of the license type to associate with this price.
	LicenseTypeID param.Field[string] `json:"license_type_id"`
	// User-specified key/value pairs for the resource. Individual keys can be removed
	// by setting the value to `null`, and the entire metadata mapping can be cleared
	// by setting `metadata` to `null`.
	Metadata param.Field[map[string]string] `json:"metadata"`
	// A transient ID that can be used to reference this price when adding adjustments
	// in the same API call.
	ReferenceID param.Field[string] `json:"reference_id"`
}

func (r PlanNewParamsPricesLicenseAllocationPriceNewLicenseAllocationTieredWithProrationPrice) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r PlanNewParamsPricesLicenseAllocationPriceNewLicenseAllocationTieredWithProrationPrice) implementsPlanNewParamsPricesLicenseAllocationPriceUnion() {
}

// The cadence to bill for this price on.
type PlanNewParamsPricesLicenseAllocationPriceNewLicenseAllocationTieredWithProrationPriceCadence string

const (
	PlanNewParamsPricesLicenseAllocationPriceNewLicenseAllocationTieredWithProrationPriceCadenceAnnual     PlanNewParamsPricesLicenseAllocationPriceNewLicenseAllocationTieredWithProrationPriceCadence = "annual"
	PlanNewParamsPricesLicenseAllocationPriceNewLicenseAllocationTieredWithProrationPriceCadenceSemiAnnual PlanNewParamsPricesLicenseAllocationPriceNewLicenseAllocationTieredWithProrationPriceCadence = "semi_annual"
	PlanNewParamsPricesLicenseAllocationPriceNewLicenseAllocationTieredWithProrationPriceCadenceMonthly    PlanNewParamsPricesLicenseAllocationPriceNewLicenseAllocationTieredWithProrationPriceCadence = "monthly"
	PlanNewParamsPricesLicenseAllocationPriceNewLicenseAllocationTieredWithProrationPriceCadenceQuarterly  PlanNewParamsPricesLicenseAllocationPriceNewLicenseAllocationTieredWithProrationPriceCadence = "quarterly"
	PlanNewParamsPricesLicenseAllocationPriceNewLicenseAllocationTieredWithProrationPriceCadenceOneTime    PlanNewParamsPricesLicenseAllocationPriceNewLicenseAllocationTieredWithProrationPriceCadence = "one_time"
	PlanNewParamsPricesLicenseAllocationPriceNewLicenseAllocationTieredWithProrationPriceCadenceCustom     PlanNewParamsPricesLicenseAllocationPriceNewLicenseAllocationTieredWithProrationPriceCadence = "custom"
)

func (r PlanNewParamsPricesLicenseAllocationPriceNewLicenseAllocationTieredWithProrationPriceCadence) IsKnown() bool {
	switch r {
	case PlanNewParamsPricesLicenseAllocationPriceNewLicenseAllocationTieredWithProrationPriceCadenceAnnual, PlanNewParamsPricesLicenseAllocationPriceNewLicenseAllocationTieredWithProrationPriceCadenceSemiAnnual, PlanNewParamsPricesLicenseAllocationPriceNewLicenseAllocationTieredWithProrationPriceCadenceMonthly, PlanNewParamsPricesLicenseAllocationPriceNewLicenseAllocationTieredWithProrationPriceCadenceQuarterly, PlanNewParamsPricesLicenseAllocationPriceNewLicenseAllocationTieredWithProrationPriceCadenceOneTime, PlanNewParamsPricesLicenseAllocationPriceNewLicenseAllocationTieredWithProrationPriceCadenceCustom:
		return true
	}
	return false
}

type PlanNewParamsPricesLicenseAllocationPriceNewLicenseAllocationTieredWithProrationPriceLicenseAllocation struct {
	// The amount of credits granted per active license per cadence.
	Amount param.Field[string] `json:"amount" api:"required"`
	// The currency of the license allocation.
	Currency param.Field[string] `json:"currency" api:"required"`
	// When True, overage beyond the allocation is written off.
	WriteOffOverage param.Field[bool] `json:"write_off_overage"`
}

func (r PlanNewParamsPricesLicenseAllocationPriceNewLicenseAllocationTieredWithProrationPriceLicenseAllocation) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// The pricing model type
type PlanNewParamsPricesLicenseAllocationPriceNewLicenseAllocationTieredWithProrationPriceModelType string

const (
	PlanNewParamsPricesLicenseAllocationPriceNewLicenseAllocationTieredWithProrationPriceModelTypeTieredWithProration PlanNewParamsPricesLicenseAllocationPriceNewLicenseAllocationTieredWithProrationPriceModelType = "tiered_with_proration"
)

func (r PlanNewParamsPricesLicenseAllocationPriceNewLicenseAllocationTieredWithProrationPriceModelType) IsKnown() bool {
	switch r {
	case PlanNewParamsPricesLicenseAllocationPriceNewLicenseAllocationTieredWithProrationPriceModelTypeTieredWithProration:
		return true
	}
	return false
}

// Configuration for tiered_with_proration pricing
type PlanNewParamsPricesLicenseAllocationPriceNewLicenseAllocationTieredWithProrationPriceTieredWithProrationConfig struct {
	// Tiers for rating based on total usage quantities into the specified tier with
	// proration
	Tiers param.Field[[]PlanNewParamsPricesLicenseAllocationPriceNewLicenseAllocationTieredWithProrationPriceTieredWithProrationConfigTier] `json:"tiers" api:"required"`
}

func (r PlanNewParamsPricesLicenseAllocationPriceNewLicenseAllocationTieredWithProrationPriceTieredWithProrationConfig) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Configuration for a single tiered with proration tier
type PlanNewParamsPricesLicenseAllocationPriceNewLicenseAllocationTieredWithProrationPriceTieredWithProrationConfigTier struct {
	// Inclusive tier starting value
	TierLowerBound param.Field[string] `json:"tier_lower_bound" api:"required"`
	// Amount per unit
	UnitAmount param.Field[string] `json:"unit_amount" api:"required"`
}

func (r PlanNewParamsPricesLicenseAllocationPriceNewLicenseAllocationTieredWithProrationPriceTieredWithProrationConfigTier) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type PlanNewParamsPricesLicenseAllocationPriceNewLicenseAllocationTieredWithProrationPriceConversionRateConfig struct {
	ConversionRateType param.Field[PlanNewParamsPricesLicenseAllocationPriceNewLicenseAllocationTieredWithProrationPriceConversionRateConfigConversionRateType] `json:"conversion_rate_type" api:"required"`
	TieredConfig       param.Field[shared.ConversionRateTieredConfigParam]                                                                                      `json:"tiered_config"`
	UnitConfig         param.Field[shared.ConversionRateUnitConfigParam]                                                                                        `json:"unit_config"`
}

func (r PlanNewParamsPricesLicenseAllocationPriceNewLicenseAllocationTieredWithProrationPriceConversionRateConfig) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r PlanNewParamsPricesLicenseAllocationPriceNewLicenseAllocationTieredWithProrationPriceConversionRateConfig) ImplementsPlanNewParamsPricesLicenseAllocationPriceNewLicenseAllocationTieredWithProrationPriceConversionRateConfigUnion() {
}

// Satisfied by [shared.UnitConversionRateConfigParam],
// [shared.TieredConversionRateConfigParam],
// [PlanNewParamsPricesLicenseAllocationPriceNewLicenseAllocationTieredWithProrationPriceConversionRateConfig].
type PlanNewParamsPricesLicenseAllocationPriceNewLicenseAllocationTieredWithProrationPriceConversionRateConfigUnion interface {
	ImplementsPlanNewParamsPricesLicenseAllocationPriceNewLicenseAllocationTieredWithProrationPriceConversionRateConfigUnion()
}

type PlanNewParamsPricesLicenseAllocationPriceNewLicenseAllocationTieredWithProrationPriceConversionRateConfigConversionRateType string

const (
	PlanNewParamsPricesLicenseAllocationPriceNewLicenseAllocationTieredWithProrationPriceConversionRateConfigConversionRateTypeUnit   PlanNewParamsPricesLicenseAllocationPriceNewLicenseAllocationTieredWithProrationPriceConversionRateConfigConversionRateType = "unit"
	PlanNewParamsPricesLicenseAllocationPriceNewLicenseAllocationTieredWithProrationPriceConversionRateConfigConversionRateTypeTiered PlanNewParamsPricesLicenseAllocationPriceNewLicenseAllocationTieredWithProrationPriceConversionRateConfigConversionRateType = "tiered"
)

func (r PlanNewParamsPricesLicenseAllocationPriceNewLicenseAllocationTieredWithProrationPriceConversionRateConfigConversionRateType) IsKnown() bool {
	switch r {
	case PlanNewParamsPricesLicenseAllocationPriceNewLicenseAllocationTieredWithProrationPriceConversionRateConfigConversionRateTypeUnit, PlanNewParamsPricesLicenseAllocationPriceNewLicenseAllocationTieredWithProrationPriceConversionRateConfigConversionRateTypeTiered:
		return true
	}
	return false
}

type PlanNewParamsPricesLicenseAllocationPriceNewLicenseAllocationUnitWithProrationPrice struct {
	// The cadence to bill for this price on.
	Cadence param.Field[PlanNewParamsPricesLicenseAllocationPriceNewLicenseAllocationUnitWithProrationPriceCadence] `json:"cadence" api:"required"`
	// The id of the item the price will be associated with.
	ItemID param.Field[string] `json:"item_id" api:"required"`
	// License allocations to associate with this price. Each entry defines a
	// per-license credit pool granted each cadence. Requires license_type_id or
	// license_type_configuration to be set.
	LicenseAllocations param.Field[[]PlanNewParamsPricesLicenseAllocationPriceNewLicenseAllocationUnitWithProrationPriceLicenseAllocation] `json:"license_allocations" api:"required"`
	// The pricing model type
	ModelType param.Field[PlanNewParamsPricesLicenseAllocationPriceNewLicenseAllocationUnitWithProrationPriceModelType] `json:"model_type" api:"required"`
	// The name of the price.
	Name param.Field[string] `json:"name" api:"required"`
	// Configuration for unit_with_proration pricing
	UnitWithProrationConfig param.Field[PlanNewParamsPricesLicenseAllocationPriceNewLicenseAllocationUnitWithProrationPriceUnitWithProrationConfig] `json:"unit_with_proration_config" api:"required"`
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
	ConversionRateConfig param.Field[PlanNewParamsPricesLicenseAllocationPriceNewLicenseAllocationUnitWithProrationPriceConversionRateConfigUnion] `json:"conversion_rate_config"`
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
	// The ID of the license type to associate with this price.
	LicenseTypeID param.Field[string] `json:"license_type_id"`
	// User-specified key/value pairs for the resource. Individual keys can be removed
	// by setting the value to `null`, and the entire metadata mapping can be cleared
	// by setting `metadata` to `null`.
	Metadata param.Field[map[string]string] `json:"metadata"`
	// A transient ID that can be used to reference this price when adding adjustments
	// in the same API call.
	ReferenceID param.Field[string] `json:"reference_id"`
}

func (r PlanNewParamsPricesLicenseAllocationPriceNewLicenseAllocationUnitWithProrationPrice) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r PlanNewParamsPricesLicenseAllocationPriceNewLicenseAllocationUnitWithProrationPrice) implementsPlanNewParamsPricesLicenseAllocationPriceUnion() {
}

// The cadence to bill for this price on.
type PlanNewParamsPricesLicenseAllocationPriceNewLicenseAllocationUnitWithProrationPriceCadence string

const (
	PlanNewParamsPricesLicenseAllocationPriceNewLicenseAllocationUnitWithProrationPriceCadenceAnnual     PlanNewParamsPricesLicenseAllocationPriceNewLicenseAllocationUnitWithProrationPriceCadence = "annual"
	PlanNewParamsPricesLicenseAllocationPriceNewLicenseAllocationUnitWithProrationPriceCadenceSemiAnnual PlanNewParamsPricesLicenseAllocationPriceNewLicenseAllocationUnitWithProrationPriceCadence = "semi_annual"
	PlanNewParamsPricesLicenseAllocationPriceNewLicenseAllocationUnitWithProrationPriceCadenceMonthly    PlanNewParamsPricesLicenseAllocationPriceNewLicenseAllocationUnitWithProrationPriceCadence = "monthly"
	PlanNewParamsPricesLicenseAllocationPriceNewLicenseAllocationUnitWithProrationPriceCadenceQuarterly  PlanNewParamsPricesLicenseAllocationPriceNewLicenseAllocationUnitWithProrationPriceCadence = "quarterly"
	PlanNewParamsPricesLicenseAllocationPriceNewLicenseAllocationUnitWithProrationPriceCadenceOneTime    PlanNewParamsPricesLicenseAllocationPriceNewLicenseAllocationUnitWithProrationPriceCadence = "one_time"
	PlanNewParamsPricesLicenseAllocationPriceNewLicenseAllocationUnitWithProrationPriceCadenceCustom     PlanNewParamsPricesLicenseAllocationPriceNewLicenseAllocationUnitWithProrationPriceCadence = "custom"
)

func (r PlanNewParamsPricesLicenseAllocationPriceNewLicenseAllocationUnitWithProrationPriceCadence) IsKnown() bool {
	switch r {
	case PlanNewParamsPricesLicenseAllocationPriceNewLicenseAllocationUnitWithProrationPriceCadenceAnnual, PlanNewParamsPricesLicenseAllocationPriceNewLicenseAllocationUnitWithProrationPriceCadenceSemiAnnual, PlanNewParamsPricesLicenseAllocationPriceNewLicenseAllocationUnitWithProrationPriceCadenceMonthly, PlanNewParamsPricesLicenseAllocationPriceNewLicenseAllocationUnitWithProrationPriceCadenceQuarterly, PlanNewParamsPricesLicenseAllocationPriceNewLicenseAllocationUnitWithProrationPriceCadenceOneTime, PlanNewParamsPricesLicenseAllocationPriceNewLicenseAllocationUnitWithProrationPriceCadenceCustom:
		return true
	}
	return false
}

type PlanNewParamsPricesLicenseAllocationPriceNewLicenseAllocationUnitWithProrationPriceLicenseAllocation struct {
	// The amount of credits granted per active license per cadence.
	Amount param.Field[string] `json:"amount" api:"required"`
	// The currency of the license allocation.
	Currency param.Field[string] `json:"currency" api:"required"`
	// When True, overage beyond the allocation is written off.
	WriteOffOverage param.Field[bool] `json:"write_off_overage"`
}

func (r PlanNewParamsPricesLicenseAllocationPriceNewLicenseAllocationUnitWithProrationPriceLicenseAllocation) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// The pricing model type
type PlanNewParamsPricesLicenseAllocationPriceNewLicenseAllocationUnitWithProrationPriceModelType string

const (
	PlanNewParamsPricesLicenseAllocationPriceNewLicenseAllocationUnitWithProrationPriceModelTypeUnitWithProration PlanNewParamsPricesLicenseAllocationPriceNewLicenseAllocationUnitWithProrationPriceModelType = "unit_with_proration"
)

func (r PlanNewParamsPricesLicenseAllocationPriceNewLicenseAllocationUnitWithProrationPriceModelType) IsKnown() bool {
	switch r {
	case PlanNewParamsPricesLicenseAllocationPriceNewLicenseAllocationUnitWithProrationPriceModelTypeUnitWithProration:
		return true
	}
	return false
}

// Configuration for unit_with_proration pricing
type PlanNewParamsPricesLicenseAllocationPriceNewLicenseAllocationUnitWithProrationPriceUnitWithProrationConfig struct {
	// Rate per unit of usage
	UnitAmount param.Field[string] `json:"unit_amount" api:"required"`
}

func (r PlanNewParamsPricesLicenseAllocationPriceNewLicenseAllocationUnitWithProrationPriceUnitWithProrationConfig) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type PlanNewParamsPricesLicenseAllocationPriceNewLicenseAllocationUnitWithProrationPriceConversionRateConfig struct {
	ConversionRateType param.Field[PlanNewParamsPricesLicenseAllocationPriceNewLicenseAllocationUnitWithProrationPriceConversionRateConfigConversionRateType] `json:"conversion_rate_type" api:"required"`
	TieredConfig       param.Field[shared.ConversionRateTieredConfigParam]                                                                                    `json:"tiered_config"`
	UnitConfig         param.Field[shared.ConversionRateUnitConfigParam]                                                                                      `json:"unit_config"`
}

func (r PlanNewParamsPricesLicenseAllocationPriceNewLicenseAllocationUnitWithProrationPriceConversionRateConfig) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r PlanNewParamsPricesLicenseAllocationPriceNewLicenseAllocationUnitWithProrationPriceConversionRateConfig) ImplementsPlanNewParamsPricesLicenseAllocationPriceNewLicenseAllocationUnitWithProrationPriceConversionRateConfigUnion() {
}

// Satisfied by [shared.UnitConversionRateConfigParam],
// [shared.TieredConversionRateConfigParam],
// [PlanNewParamsPricesLicenseAllocationPriceNewLicenseAllocationUnitWithProrationPriceConversionRateConfig].
type PlanNewParamsPricesLicenseAllocationPriceNewLicenseAllocationUnitWithProrationPriceConversionRateConfigUnion interface {
	ImplementsPlanNewParamsPricesLicenseAllocationPriceNewLicenseAllocationUnitWithProrationPriceConversionRateConfigUnion()
}

type PlanNewParamsPricesLicenseAllocationPriceNewLicenseAllocationUnitWithProrationPriceConversionRateConfigConversionRateType string

const (
	PlanNewParamsPricesLicenseAllocationPriceNewLicenseAllocationUnitWithProrationPriceConversionRateConfigConversionRateTypeUnit   PlanNewParamsPricesLicenseAllocationPriceNewLicenseAllocationUnitWithProrationPriceConversionRateConfigConversionRateType = "unit"
	PlanNewParamsPricesLicenseAllocationPriceNewLicenseAllocationUnitWithProrationPriceConversionRateConfigConversionRateTypeTiered PlanNewParamsPricesLicenseAllocationPriceNewLicenseAllocationUnitWithProrationPriceConversionRateConfigConversionRateType = "tiered"
)

func (r PlanNewParamsPricesLicenseAllocationPriceNewLicenseAllocationUnitWithProrationPriceConversionRateConfigConversionRateType) IsKnown() bool {
	switch r {
	case PlanNewParamsPricesLicenseAllocationPriceNewLicenseAllocationUnitWithProrationPriceConversionRateConfigConversionRateTypeUnit, PlanNewParamsPricesLicenseAllocationPriceNewLicenseAllocationUnitWithProrationPriceConversionRateConfigConversionRateTypeTiered:
		return true
	}
	return false
}

type PlanNewParamsPricesLicenseAllocationPriceNewLicenseAllocationGroupedAllocationPrice struct {
	// The cadence to bill for this price on.
	Cadence param.Field[PlanNewParamsPricesLicenseAllocationPriceNewLicenseAllocationGroupedAllocationPriceCadence] `json:"cadence" api:"required"`
	// Configuration for grouped_allocation pricing
	GroupedAllocationConfig param.Field[PlanNewParamsPricesLicenseAllocationPriceNewLicenseAllocationGroupedAllocationPriceGroupedAllocationConfig] `json:"grouped_allocation_config" api:"required"`
	// The id of the item the price will be associated with.
	ItemID param.Field[string] `json:"item_id" api:"required"`
	// License allocations to associate with this price. Each entry defines a
	// per-license credit pool granted each cadence. Requires license_type_id or
	// license_type_configuration to be set.
	LicenseAllocations param.Field[[]PlanNewParamsPricesLicenseAllocationPriceNewLicenseAllocationGroupedAllocationPriceLicenseAllocation] `json:"license_allocations" api:"required"`
	// The pricing model type
	ModelType param.Field[PlanNewParamsPricesLicenseAllocationPriceNewLicenseAllocationGroupedAllocationPriceModelType] `json:"model_type" api:"required"`
	// The name of the price.
	Name param.Field[string] `json:"name" api:"required"`
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
	ConversionRateConfig param.Field[PlanNewParamsPricesLicenseAllocationPriceNewLicenseAllocationGroupedAllocationPriceConversionRateConfigUnion] `json:"conversion_rate_config"`
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
	// The ID of the license type to associate with this price.
	LicenseTypeID param.Field[string] `json:"license_type_id"`
	// User-specified key/value pairs for the resource. Individual keys can be removed
	// by setting the value to `null`, and the entire metadata mapping can be cleared
	// by setting `metadata` to `null`.
	Metadata param.Field[map[string]string] `json:"metadata"`
	// A transient ID that can be used to reference this price when adding adjustments
	// in the same API call.
	ReferenceID param.Field[string] `json:"reference_id"`
}

func (r PlanNewParamsPricesLicenseAllocationPriceNewLicenseAllocationGroupedAllocationPrice) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r PlanNewParamsPricesLicenseAllocationPriceNewLicenseAllocationGroupedAllocationPrice) implementsPlanNewParamsPricesLicenseAllocationPriceUnion() {
}

// The cadence to bill for this price on.
type PlanNewParamsPricesLicenseAllocationPriceNewLicenseAllocationGroupedAllocationPriceCadence string

const (
	PlanNewParamsPricesLicenseAllocationPriceNewLicenseAllocationGroupedAllocationPriceCadenceAnnual     PlanNewParamsPricesLicenseAllocationPriceNewLicenseAllocationGroupedAllocationPriceCadence = "annual"
	PlanNewParamsPricesLicenseAllocationPriceNewLicenseAllocationGroupedAllocationPriceCadenceSemiAnnual PlanNewParamsPricesLicenseAllocationPriceNewLicenseAllocationGroupedAllocationPriceCadence = "semi_annual"
	PlanNewParamsPricesLicenseAllocationPriceNewLicenseAllocationGroupedAllocationPriceCadenceMonthly    PlanNewParamsPricesLicenseAllocationPriceNewLicenseAllocationGroupedAllocationPriceCadence = "monthly"
	PlanNewParamsPricesLicenseAllocationPriceNewLicenseAllocationGroupedAllocationPriceCadenceQuarterly  PlanNewParamsPricesLicenseAllocationPriceNewLicenseAllocationGroupedAllocationPriceCadence = "quarterly"
	PlanNewParamsPricesLicenseAllocationPriceNewLicenseAllocationGroupedAllocationPriceCadenceOneTime    PlanNewParamsPricesLicenseAllocationPriceNewLicenseAllocationGroupedAllocationPriceCadence = "one_time"
	PlanNewParamsPricesLicenseAllocationPriceNewLicenseAllocationGroupedAllocationPriceCadenceCustom     PlanNewParamsPricesLicenseAllocationPriceNewLicenseAllocationGroupedAllocationPriceCadence = "custom"
)

func (r PlanNewParamsPricesLicenseAllocationPriceNewLicenseAllocationGroupedAllocationPriceCadence) IsKnown() bool {
	switch r {
	case PlanNewParamsPricesLicenseAllocationPriceNewLicenseAllocationGroupedAllocationPriceCadenceAnnual, PlanNewParamsPricesLicenseAllocationPriceNewLicenseAllocationGroupedAllocationPriceCadenceSemiAnnual, PlanNewParamsPricesLicenseAllocationPriceNewLicenseAllocationGroupedAllocationPriceCadenceMonthly, PlanNewParamsPricesLicenseAllocationPriceNewLicenseAllocationGroupedAllocationPriceCadenceQuarterly, PlanNewParamsPricesLicenseAllocationPriceNewLicenseAllocationGroupedAllocationPriceCadenceOneTime, PlanNewParamsPricesLicenseAllocationPriceNewLicenseAllocationGroupedAllocationPriceCadenceCustom:
		return true
	}
	return false
}

// Configuration for grouped_allocation pricing
type PlanNewParamsPricesLicenseAllocationPriceNewLicenseAllocationGroupedAllocationPriceGroupedAllocationConfig struct {
	// Usage allocation per group
	Allocation param.Field[string] `json:"allocation" api:"required"`
	// How to determine the groups that should each be allocated some quantity
	GroupingKey param.Field[string] `json:"grouping_key" api:"required"`
	// Unit rate for post-allocation
	OverageUnitRate param.Field[string] `json:"overage_unit_rate" api:"required"`
}

func (r PlanNewParamsPricesLicenseAllocationPriceNewLicenseAllocationGroupedAllocationPriceGroupedAllocationConfig) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type PlanNewParamsPricesLicenseAllocationPriceNewLicenseAllocationGroupedAllocationPriceLicenseAllocation struct {
	// The amount of credits granted per active license per cadence.
	Amount param.Field[string] `json:"amount" api:"required"`
	// The currency of the license allocation.
	Currency param.Field[string] `json:"currency" api:"required"`
	// When True, overage beyond the allocation is written off.
	WriteOffOverage param.Field[bool] `json:"write_off_overage"`
}

func (r PlanNewParamsPricesLicenseAllocationPriceNewLicenseAllocationGroupedAllocationPriceLicenseAllocation) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// The pricing model type
type PlanNewParamsPricesLicenseAllocationPriceNewLicenseAllocationGroupedAllocationPriceModelType string

const (
	PlanNewParamsPricesLicenseAllocationPriceNewLicenseAllocationGroupedAllocationPriceModelTypeGroupedAllocation PlanNewParamsPricesLicenseAllocationPriceNewLicenseAllocationGroupedAllocationPriceModelType = "grouped_allocation"
)

func (r PlanNewParamsPricesLicenseAllocationPriceNewLicenseAllocationGroupedAllocationPriceModelType) IsKnown() bool {
	switch r {
	case PlanNewParamsPricesLicenseAllocationPriceNewLicenseAllocationGroupedAllocationPriceModelTypeGroupedAllocation:
		return true
	}
	return false
}

type PlanNewParamsPricesLicenseAllocationPriceNewLicenseAllocationGroupedAllocationPriceConversionRateConfig struct {
	ConversionRateType param.Field[PlanNewParamsPricesLicenseAllocationPriceNewLicenseAllocationGroupedAllocationPriceConversionRateConfigConversionRateType] `json:"conversion_rate_type" api:"required"`
	TieredConfig       param.Field[shared.ConversionRateTieredConfigParam]                                                                                    `json:"tiered_config"`
	UnitConfig         param.Field[shared.ConversionRateUnitConfigParam]                                                                                      `json:"unit_config"`
}

func (r PlanNewParamsPricesLicenseAllocationPriceNewLicenseAllocationGroupedAllocationPriceConversionRateConfig) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r PlanNewParamsPricesLicenseAllocationPriceNewLicenseAllocationGroupedAllocationPriceConversionRateConfig) ImplementsPlanNewParamsPricesLicenseAllocationPriceNewLicenseAllocationGroupedAllocationPriceConversionRateConfigUnion() {
}

// Satisfied by [shared.UnitConversionRateConfigParam],
// [shared.TieredConversionRateConfigParam],
// [PlanNewParamsPricesLicenseAllocationPriceNewLicenseAllocationGroupedAllocationPriceConversionRateConfig].
type PlanNewParamsPricesLicenseAllocationPriceNewLicenseAllocationGroupedAllocationPriceConversionRateConfigUnion interface {
	ImplementsPlanNewParamsPricesLicenseAllocationPriceNewLicenseAllocationGroupedAllocationPriceConversionRateConfigUnion()
}

type PlanNewParamsPricesLicenseAllocationPriceNewLicenseAllocationGroupedAllocationPriceConversionRateConfigConversionRateType string

const (
	PlanNewParamsPricesLicenseAllocationPriceNewLicenseAllocationGroupedAllocationPriceConversionRateConfigConversionRateTypeUnit   PlanNewParamsPricesLicenseAllocationPriceNewLicenseAllocationGroupedAllocationPriceConversionRateConfigConversionRateType = "unit"
	PlanNewParamsPricesLicenseAllocationPriceNewLicenseAllocationGroupedAllocationPriceConversionRateConfigConversionRateTypeTiered PlanNewParamsPricesLicenseAllocationPriceNewLicenseAllocationGroupedAllocationPriceConversionRateConfigConversionRateType = "tiered"
)

func (r PlanNewParamsPricesLicenseAllocationPriceNewLicenseAllocationGroupedAllocationPriceConversionRateConfigConversionRateType) IsKnown() bool {
	switch r {
	case PlanNewParamsPricesLicenseAllocationPriceNewLicenseAllocationGroupedAllocationPriceConversionRateConfigConversionRateTypeUnit, PlanNewParamsPricesLicenseAllocationPriceNewLicenseAllocationGroupedAllocationPriceConversionRateConfigConversionRateTypeTiered:
		return true
	}
	return false
}

type PlanNewParamsPricesLicenseAllocationPriceNewLicenseAllocationBulkWithProrationPrice struct {
	// Configuration for bulk_with_proration pricing
	BulkWithProrationConfig param.Field[PlanNewParamsPricesLicenseAllocationPriceNewLicenseAllocationBulkWithProrationPriceBulkWithProrationConfig] `json:"bulk_with_proration_config" api:"required"`
	// The cadence to bill for this price on.
	Cadence param.Field[PlanNewParamsPricesLicenseAllocationPriceNewLicenseAllocationBulkWithProrationPriceCadence] `json:"cadence" api:"required"`
	// The id of the item the price will be associated with.
	ItemID param.Field[string] `json:"item_id" api:"required"`
	// License allocations to associate with this price. Each entry defines a
	// per-license credit pool granted each cadence. Requires license_type_id or
	// license_type_configuration to be set.
	LicenseAllocations param.Field[[]PlanNewParamsPricesLicenseAllocationPriceNewLicenseAllocationBulkWithProrationPriceLicenseAllocation] `json:"license_allocations" api:"required"`
	// The pricing model type
	ModelType param.Field[PlanNewParamsPricesLicenseAllocationPriceNewLicenseAllocationBulkWithProrationPriceModelType] `json:"model_type" api:"required"`
	// The name of the price.
	Name param.Field[string] `json:"name" api:"required"`
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
	ConversionRateConfig param.Field[PlanNewParamsPricesLicenseAllocationPriceNewLicenseAllocationBulkWithProrationPriceConversionRateConfigUnion] `json:"conversion_rate_config"`
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
	// The ID of the license type to associate with this price.
	LicenseTypeID param.Field[string] `json:"license_type_id"`
	// User-specified key/value pairs for the resource. Individual keys can be removed
	// by setting the value to `null`, and the entire metadata mapping can be cleared
	// by setting `metadata` to `null`.
	Metadata param.Field[map[string]string] `json:"metadata"`
	// A transient ID that can be used to reference this price when adding adjustments
	// in the same API call.
	ReferenceID param.Field[string] `json:"reference_id"`
}

func (r PlanNewParamsPricesLicenseAllocationPriceNewLicenseAllocationBulkWithProrationPrice) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r PlanNewParamsPricesLicenseAllocationPriceNewLicenseAllocationBulkWithProrationPrice) implementsPlanNewParamsPricesLicenseAllocationPriceUnion() {
}

// Configuration for bulk_with_proration pricing
type PlanNewParamsPricesLicenseAllocationPriceNewLicenseAllocationBulkWithProrationPriceBulkWithProrationConfig struct {
	// Bulk tiers for rating based on total usage volume
	Tiers param.Field[[]PlanNewParamsPricesLicenseAllocationPriceNewLicenseAllocationBulkWithProrationPriceBulkWithProrationConfigTier] `json:"tiers" api:"required"`
}

func (r PlanNewParamsPricesLicenseAllocationPriceNewLicenseAllocationBulkWithProrationPriceBulkWithProrationConfig) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Configuration for a single bulk pricing tier with proration
type PlanNewParamsPricesLicenseAllocationPriceNewLicenseAllocationBulkWithProrationPriceBulkWithProrationConfigTier struct {
	// Cost per unit
	UnitAmount param.Field[string] `json:"unit_amount" api:"required"`
	// The lower bound for this tier
	TierLowerBound param.Field[string] `json:"tier_lower_bound"`
}

func (r PlanNewParamsPricesLicenseAllocationPriceNewLicenseAllocationBulkWithProrationPriceBulkWithProrationConfigTier) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// The cadence to bill for this price on.
type PlanNewParamsPricesLicenseAllocationPriceNewLicenseAllocationBulkWithProrationPriceCadence string

const (
	PlanNewParamsPricesLicenseAllocationPriceNewLicenseAllocationBulkWithProrationPriceCadenceAnnual     PlanNewParamsPricesLicenseAllocationPriceNewLicenseAllocationBulkWithProrationPriceCadence = "annual"
	PlanNewParamsPricesLicenseAllocationPriceNewLicenseAllocationBulkWithProrationPriceCadenceSemiAnnual PlanNewParamsPricesLicenseAllocationPriceNewLicenseAllocationBulkWithProrationPriceCadence = "semi_annual"
	PlanNewParamsPricesLicenseAllocationPriceNewLicenseAllocationBulkWithProrationPriceCadenceMonthly    PlanNewParamsPricesLicenseAllocationPriceNewLicenseAllocationBulkWithProrationPriceCadence = "monthly"
	PlanNewParamsPricesLicenseAllocationPriceNewLicenseAllocationBulkWithProrationPriceCadenceQuarterly  PlanNewParamsPricesLicenseAllocationPriceNewLicenseAllocationBulkWithProrationPriceCadence = "quarterly"
	PlanNewParamsPricesLicenseAllocationPriceNewLicenseAllocationBulkWithProrationPriceCadenceOneTime    PlanNewParamsPricesLicenseAllocationPriceNewLicenseAllocationBulkWithProrationPriceCadence = "one_time"
	PlanNewParamsPricesLicenseAllocationPriceNewLicenseAllocationBulkWithProrationPriceCadenceCustom     PlanNewParamsPricesLicenseAllocationPriceNewLicenseAllocationBulkWithProrationPriceCadence = "custom"
)

func (r PlanNewParamsPricesLicenseAllocationPriceNewLicenseAllocationBulkWithProrationPriceCadence) IsKnown() bool {
	switch r {
	case PlanNewParamsPricesLicenseAllocationPriceNewLicenseAllocationBulkWithProrationPriceCadenceAnnual, PlanNewParamsPricesLicenseAllocationPriceNewLicenseAllocationBulkWithProrationPriceCadenceSemiAnnual, PlanNewParamsPricesLicenseAllocationPriceNewLicenseAllocationBulkWithProrationPriceCadenceMonthly, PlanNewParamsPricesLicenseAllocationPriceNewLicenseAllocationBulkWithProrationPriceCadenceQuarterly, PlanNewParamsPricesLicenseAllocationPriceNewLicenseAllocationBulkWithProrationPriceCadenceOneTime, PlanNewParamsPricesLicenseAllocationPriceNewLicenseAllocationBulkWithProrationPriceCadenceCustom:
		return true
	}
	return false
}

type PlanNewParamsPricesLicenseAllocationPriceNewLicenseAllocationBulkWithProrationPriceLicenseAllocation struct {
	// The amount of credits granted per active license per cadence.
	Amount param.Field[string] `json:"amount" api:"required"`
	// The currency of the license allocation.
	Currency param.Field[string] `json:"currency" api:"required"`
	// When True, overage beyond the allocation is written off.
	WriteOffOverage param.Field[bool] `json:"write_off_overage"`
}

func (r PlanNewParamsPricesLicenseAllocationPriceNewLicenseAllocationBulkWithProrationPriceLicenseAllocation) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// The pricing model type
type PlanNewParamsPricesLicenseAllocationPriceNewLicenseAllocationBulkWithProrationPriceModelType string

const (
	PlanNewParamsPricesLicenseAllocationPriceNewLicenseAllocationBulkWithProrationPriceModelTypeBulkWithProration PlanNewParamsPricesLicenseAllocationPriceNewLicenseAllocationBulkWithProrationPriceModelType = "bulk_with_proration"
)

func (r PlanNewParamsPricesLicenseAllocationPriceNewLicenseAllocationBulkWithProrationPriceModelType) IsKnown() bool {
	switch r {
	case PlanNewParamsPricesLicenseAllocationPriceNewLicenseAllocationBulkWithProrationPriceModelTypeBulkWithProration:
		return true
	}
	return false
}

type PlanNewParamsPricesLicenseAllocationPriceNewLicenseAllocationBulkWithProrationPriceConversionRateConfig struct {
	ConversionRateType param.Field[PlanNewParamsPricesLicenseAllocationPriceNewLicenseAllocationBulkWithProrationPriceConversionRateConfigConversionRateType] `json:"conversion_rate_type" api:"required"`
	TieredConfig       param.Field[shared.ConversionRateTieredConfigParam]                                                                                    `json:"tiered_config"`
	UnitConfig         param.Field[shared.ConversionRateUnitConfigParam]                                                                                      `json:"unit_config"`
}

func (r PlanNewParamsPricesLicenseAllocationPriceNewLicenseAllocationBulkWithProrationPriceConversionRateConfig) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r PlanNewParamsPricesLicenseAllocationPriceNewLicenseAllocationBulkWithProrationPriceConversionRateConfig) ImplementsPlanNewParamsPricesLicenseAllocationPriceNewLicenseAllocationBulkWithProrationPriceConversionRateConfigUnion() {
}

// Satisfied by [shared.UnitConversionRateConfigParam],
// [shared.TieredConversionRateConfigParam],
// [PlanNewParamsPricesLicenseAllocationPriceNewLicenseAllocationBulkWithProrationPriceConversionRateConfig].
type PlanNewParamsPricesLicenseAllocationPriceNewLicenseAllocationBulkWithProrationPriceConversionRateConfigUnion interface {
	ImplementsPlanNewParamsPricesLicenseAllocationPriceNewLicenseAllocationBulkWithProrationPriceConversionRateConfigUnion()
}

type PlanNewParamsPricesLicenseAllocationPriceNewLicenseAllocationBulkWithProrationPriceConversionRateConfigConversionRateType string

const (
	PlanNewParamsPricesLicenseAllocationPriceNewLicenseAllocationBulkWithProrationPriceConversionRateConfigConversionRateTypeUnit   PlanNewParamsPricesLicenseAllocationPriceNewLicenseAllocationBulkWithProrationPriceConversionRateConfigConversionRateType = "unit"
	PlanNewParamsPricesLicenseAllocationPriceNewLicenseAllocationBulkWithProrationPriceConversionRateConfigConversionRateTypeTiered PlanNewParamsPricesLicenseAllocationPriceNewLicenseAllocationBulkWithProrationPriceConversionRateConfigConversionRateType = "tiered"
)

func (r PlanNewParamsPricesLicenseAllocationPriceNewLicenseAllocationBulkWithProrationPriceConversionRateConfigConversionRateType) IsKnown() bool {
	switch r {
	case PlanNewParamsPricesLicenseAllocationPriceNewLicenseAllocationBulkWithProrationPriceConversionRateConfigConversionRateTypeUnit, PlanNewParamsPricesLicenseAllocationPriceNewLicenseAllocationBulkWithProrationPriceConversionRateConfigConversionRateTypeTiered:
		return true
	}
	return false
}

type PlanNewParamsPricesLicenseAllocationPriceNewLicenseAllocationGroupedWithProratedMinimumPrice struct {
	// The cadence to bill for this price on.
	Cadence param.Field[PlanNewParamsPricesLicenseAllocationPriceNewLicenseAllocationGroupedWithProratedMinimumPriceCadence] `json:"cadence" api:"required"`
	// Configuration for grouped_with_prorated_minimum pricing
	GroupedWithProratedMinimumConfig param.Field[PlanNewParamsPricesLicenseAllocationPriceNewLicenseAllocationGroupedWithProratedMinimumPriceGroupedWithProratedMinimumConfig] `json:"grouped_with_prorated_minimum_config" api:"required"`
	// The id of the item the price will be associated with.
	ItemID param.Field[string] `json:"item_id" api:"required"`
	// License allocations to associate with this price. Each entry defines a
	// per-license credit pool granted each cadence. Requires license_type_id or
	// license_type_configuration to be set.
	LicenseAllocations param.Field[[]PlanNewParamsPricesLicenseAllocationPriceNewLicenseAllocationGroupedWithProratedMinimumPriceLicenseAllocation] `json:"license_allocations" api:"required"`
	// The pricing model type
	ModelType param.Field[PlanNewParamsPricesLicenseAllocationPriceNewLicenseAllocationGroupedWithProratedMinimumPriceModelType] `json:"model_type" api:"required"`
	// The name of the price.
	Name param.Field[string] `json:"name" api:"required"`
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
	ConversionRateConfig param.Field[PlanNewParamsPricesLicenseAllocationPriceNewLicenseAllocationGroupedWithProratedMinimumPriceConversionRateConfigUnion] `json:"conversion_rate_config"`
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
	// The ID of the license type to associate with this price.
	LicenseTypeID param.Field[string] `json:"license_type_id"`
	// User-specified key/value pairs for the resource. Individual keys can be removed
	// by setting the value to `null`, and the entire metadata mapping can be cleared
	// by setting `metadata` to `null`.
	Metadata param.Field[map[string]string] `json:"metadata"`
	// A transient ID that can be used to reference this price when adding adjustments
	// in the same API call.
	ReferenceID param.Field[string] `json:"reference_id"`
}

func (r PlanNewParamsPricesLicenseAllocationPriceNewLicenseAllocationGroupedWithProratedMinimumPrice) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r PlanNewParamsPricesLicenseAllocationPriceNewLicenseAllocationGroupedWithProratedMinimumPrice) implementsPlanNewParamsPricesLicenseAllocationPriceUnion() {
}

// The cadence to bill for this price on.
type PlanNewParamsPricesLicenseAllocationPriceNewLicenseAllocationGroupedWithProratedMinimumPriceCadence string

const (
	PlanNewParamsPricesLicenseAllocationPriceNewLicenseAllocationGroupedWithProratedMinimumPriceCadenceAnnual     PlanNewParamsPricesLicenseAllocationPriceNewLicenseAllocationGroupedWithProratedMinimumPriceCadence = "annual"
	PlanNewParamsPricesLicenseAllocationPriceNewLicenseAllocationGroupedWithProratedMinimumPriceCadenceSemiAnnual PlanNewParamsPricesLicenseAllocationPriceNewLicenseAllocationGroupedWithProratedMinimumPriceCadence = "semi_annual"
	PlanNewParamsPricesLicenseAllocationPriceNewLicenseAllocationGroupedWithProratedMinimumPriceCadenceMonthly    PlanNewParamsPricesLicenseAllocationPriceNewLicenseAllocationGroupedWithProratedMinimumPriceCadence = "monthly"
	PlanNewParamsPricesLicenseAllocationPriceNewLicenseAllocationGroupedWithProratedMinimumPriceCadenceQuarterly  PlanNewParamsPricesLicenseAllocationPriceNewLicenseAllocationGroupedWithProratedMinimumPriceCadence = "quarterly"
	PlanNewParamsPricesLicenseAllocationPriceNewLicenseAllocationGroupedWithProratedMinimumPriceCadenceOneTime    PlanNewParamsPricesLicenseAllocationPriceNewLicenseAllocationGroupedWithProratedMinimumPriceCadence = "one_time"
	PlanNewParamsPricesLicenseAllocationPriceNewLicenseAllocationGroupedWithProratedMinimumPriceCadenceCustom     PlanNewParamsPricesLicenseAllocationPriceNewLicenseAllocationGroupedWithProratedMinimumPriceCadence = "custom"
)

func (r PlanNewParamsPricesLicenseAllocationPriceNewLicenseAllocationGroupedWithProratedMinimumPriceCadence) IsKnown() bool {
	switch r {
	case PlanNewParamsPricesLicenseAllocationPriceNewLicenseAllocationGroupedWithProratedMinimumPriceCadenceAnnual, PlanNewParamsPricesLicenseAllocationPriceNewLicenseAllocationGroupedWithProratedMinimumPriceCadenceSemiAnnual, PlanNewParamsPricesLicenseAllocationPriceNewLicenseAllocationGroupedWithProratedMinimumPriceCadenceMonthly, PlanNewParamsPricesLicenseAllocationPriceNewLicenseAllocationGroupedWithProratedMinimumPriceCadenceQuarterly, PlanNewParamsPricesLicenseAllocationPriceNewLicenseAllocationGroupedWithProratedMinimumPriceCadenceOneTime, PlanNewParamsPricesLicenseAllocationPriceNewLicenseAllocationGroupedWithProratedMinimumPriceCadenceCustom:
		return true
	}
	return false
}

// Configuration for grouped_with_prorated_minimum pricing
type PlanNewParamsPricesLicenseAllocationPriceNewLicenseAllocationGroupedWithProratedMinimumPriceGroupedWithProratedMinimumConfig struct {
	// How to determine the groups that should each have a minimum
	GroupingKey param.Field[string] `json:"grouping_key" api:"required"`
	// The minimum amount to charge per group
	Minimum param.Field[string] `json:"minimum" api:"required"`
	// The amount to charge per unit
	UnitRate param.Field[string] `json:"unit_rate" api:"required"`
}

func (r PlanNewParamsPricesLicenseAllocationPriceNewLicenseAllocationGroupedWithProratedMinimumPriceGroupedWithProratedMinimumConfig) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type PlanNewParamsPricesLicenseAllocationPriceNewLicenseAllocationGroupedWithProratedMinimumPriceLicenseAllocation struct {
	// The amount of credits granted per active license per cadence.
	Amount param.Field[string] `json:"amount" api:"required"`
	// The currency of the license allocation.
	Currency param.Field[string] `json:"currency" api:"required"`
	// When True, overage beyond the allocation is written off.
	WriteOffOverage param.Field[bool] `json:"write_off_overage"`
}

func (r PlanNewParamsPricesLicenseAllocationPriceNewLicenseAllocationGroupedWithProratedMinimumPriceLicenseAllocation) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// The pricing model type
type PlanNewParamsPricesLicenseAllocationPriceNewLicenseAllocationGroupedWithProratedMinimumPriceModelType string

const (
	PlanNewParamsPricesLicenseAllocationPriceNewLicenseAllocationGroupedWithProratedMinimumPriceModelTypeGroupedWithProratedMinimum PlanNewParamsPricesLicenseAllocationPriceNewLicenseAllocationGroupedWithProratedMinimumPriceModelType = "grouped_with_prorated_minimum"
)

func (r PlanNewParamsPricesLicenseAllocationPriceNewLicenseAllocationGroupedWithProratedMinimumPriceModelType) IsKnown() bool {
	switch r {
	case PlanNewParamsPricesLicenseAllocationPriceNewLicenseAllocationGroupedWithProratedMinimumPriceModelTypeGroupedWithProratedMinimum:
		return true
	}
	return false
}

type PlanNewParamsPricesLicenseAllocationPriceNewLicenseAllocationGroupedWithProratedMinimumPriceConversionRateConfig struct {
	ConversionRateType param.Field[PlanNewParamsPricesLicenseAllocationPriceNewLicenseAllocationGroupedWithProratedMinimumPriceConversionRateConfigConversionRateType] `json:"conversion_rate_type" api:"required"`
	TieredConfig       param.Field[shared.ConversionRateTieredConfigParam]                                                                                             `json:"tiered_config"`
	UnitConfig         param.Field[shared.ConversionRateUnitConfigParam]                                                                                               `json:"unit_config"`
}

func (r PlanNewParamsPricesLicenseAllocationPriceNewLicenseAllocationGroupedWithProratedMinimumPriceConversionRateConfig) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r PlanNewParamsPricesLicenseAllocationPriceNewLicenseAllocationGroupedWithProratedMinimumPriceConversionRateConfig) ImplementsPlanNewParamsPricesLicenseAllocationPriceNewLicenseAllocationGroupedWithProratedMinimumPriceConversionRateConfigUnion() {
}

// Satisfied by [shared.UnitConversionRateConfigParam],
// [shared.TieredConversionRateConfigParam],
// [PlanNewParamsPricesLicenseAllocationPriceNewLicenseAllocationGroupedWithProratedMinimumPriceConversionRateConfig].
type PlanNewParamsPricesLicenseAllocationPriceNewLicenseAllocationGroupedWithProratedMinimumPriceConversionRateConfigUnion interface {
	ImplementsPlanNewParamsPricesLicenseAllocationPriceNewLicenseAllocationGroupedWithProratedMinimumPriceConversionRateConfigUnion()
}

type PlanNewParamsPricesLicenseAllocationPriceNewLicenseAllocationGroupedWithProratedMinimumPriceConversionRateConfigConversionRateType string

const (
	PlanNewParamsPricesLicenseAllocationPriceNewLicenseAllocationGroupedWithProratedMinimumPriceConversionRateConfigConversionRateTypeUnit   PlanNewParamsPricesLicenseAllocationPriceNewLicenseAllocationGroupedWithProratedMinimumPriceConversionRateConfigConversionRateType = "unit"
	PlanNewParamsPricesLicenseAllocationPriceNewLicenseAllocationGroupedWithProratedMinimumPriceConversionRateConfigConversionRateTypeTiered PlanNewParamsPricesLicenseAllocationPriceNewLicenseAllocationGroupedWithProratedMinimumPriceConversionRateConfigConversionRateType = "tiered"
)

func (r PlanNewParamsPricesLicenseAllocationPriceNewLicenseAllocationGroupedWithProratedMinimumPriceConversionRateConfigConversionRateType) IsKnown() bool {
	switch r {
	case PlanNewParamsPricesLicenseAllocationPriceNewLicenseAllocationGroupedWithProratedMinimumPriceConversionRateConfigConversionRateTypeUnit, PlanNewParamsPricesLicenseAllocationPriceNewLicenseAllocationGroupedWithProratedMinimumPriceConversionRateConfigConversionRateTypeTiered:
		return true
	}
	return false
}

type PlanNewParamsPricesLicenseAllocationPriceNewLicenseAllocationGroupedWithMeteredMinimumPrice struct {
	// The cadence to bill for this price on.
	Cadence param.Field[PlanNewParamsPricesLicenseAllocationPriceNewLicenseAllocationGroupedWithMeteredMinimumPriceCadence] `json:"cadence" api:"required"`
	// Configuration for grouped_with_metered_minimum pricing
	GroupedWithMeteredMinimumConfig param.Field[PlanNewParamsPricesLicenseAllocationPriceNewLicenseAllocationGroupedWithMeteredMinimumPriceGroupedWithMeteredMinimumConfig] `json:"grouped_with_metered_minimum_config" api:"required"`
	// The id of the item the price will be associated with.
	ItemID param.Field[string] `json:"item_id" api:"required"`
	// License allocations to associate with this price. Each entry defines a
	// per-license credit pool granted each cadence. Requires license_type_id or
	// license_type_configuration to be set.
	LicenseAllocations param.Field[[]PlanNewParamsPricesLicenseAllocationPriceNewLicenseAllocationGroupedWithMeteredMinimumPriceLicenseAllocation] `json:"license_allocations" api:"required"`
	// The pricing model type
	ModelType param.Field[PlanNewParamsPricesLicenseAllocationPriceNewLicenseAllocationGroupedWithMeteredMinimumPriceModelType] `json:"model_type" api:"required"`
	// The name of the price.
	Name param.Field[string] `json:"name" api:"required"`
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
	ConversionRateConfig param.Field[PlanNewParamsPricesLicenseAllocationPriceNewLicenseAllocationGroupedWithMeteredMinimumPriceConversionRateConfigUnion] `json:"conversion_rate_config"`
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
	// The ID of the license type to associate with this price.
	LicenseTypeID param.Field[string] `json:"license_type_id"`
	// User-specified key/value pairs for the resource. Individual keys can be removed
	// by setting the value to `null`, and the entire metadata mapping can be cleared
	// by setting `metadata` to `null`.
	Metadata param.Field[map[string]string] `json:"metadata"`
	// A transient ID that can be used to reference this price when adding adjustments
	// in the same API call.
	ReferenceID param.Field[string] `json:"reference_id"`
}

func (r PlanNewParamsPricesLicenseAllocationPriceNewLicenseAllocationGroupedWithMeteredMinimumPrice) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r PlanNewParamsPricesLicenseAllocationPriceNewLicenseAllocationGroupedWithMeteredMinimumPrice) implementsPlanNewParamsPricesLicenseAllocationPriceUnion() {
}

// The cadence to bill for this price on.
type PlanNewParamsPricesLicenseAllocationPriceNewLicenseAllocationGroupedWithMeteredMinimumPriceCadence string

const (
	PlanNewParamsPricesLicenseAllocationPriceNewLicenseAllocationGroupedWithMeteredMinimumPriceCadenceAnnual     PlanNewParamsPricesLicenseAllocationPriceNewLicenseAllocationGroupedWithMeteredMinimumPriceCadence = "annual"
	PlanNewParamsPricesLicenseAllocationPriceNewLicenseAllocationGroupedWithMeteredMinimumPriceCadenceSemiAnnual PlanNewParamsPricesLicenseAllocationPriceNewLicenseAllocationGroupedWithMeteredMinimumPriceCadence = "semi_annual"
	PlanNewParamsPricesLicenseAllocationPriceNewLicenseAllocationGroupedWithMeteredMinimumPriceCadenceMonthly    PlanNewParamsPricesLicenseAllocationPriceNewLicenseAllocationGroupedWithMeteredMinimumPriceCadence = "monthly"
	PlanNewParamsPricesLicenseAllocationPriceNewLicenseAllocationGroupedWithMeteredMinimumPriceCadenceQuarterly  PlanNewParamsPricesLicenseAllocationPriceNewLicenseAllocationGroupedWithMeteredMinimumPriceCadence = "quarterly"
	PlanNewParamsPricesLicenseAllocationPriceNewLicenseAllocationGroupedWithMeteredMinimumPriceCadenceOneTime    PlanNewParamsPricesLicenseAllocationPriceNewLicenseAllocationGroupedWithMeteredMinimumPriceCadence = "one_time"
	PlanNewParamsPricesLicenseAllocationPriceNewLicenseAllocationGroupedWithMeteredMinimumPriceCadenceCustom     PlanNewParamsPricesLicenseAllocationPriceNewLicenseAllocationGroupedWithMeteredMinimumPriceCadence = "custom"
)

func (r PlanNewParamsPricesLicenseAllocationPriceNewLicenseAllocationGroupedWithMeteredMinimumPriceCadence) IsKnown() bool {
	switch r {
	case PlanNewParamsPricesLicenseAllocationPriceNewLicenseAllocationGroupedWithMeteredMinimumPriceCadenceAnnual, PlanNewParamsPricesLicenseAllocationPriceNewLicenseAllocationGroupedWithMeteredMinimumPriceCadenceSemiAnnual, PlanNewParamsPricesLicenseAllocationPriceNewLicenseAllocationGroupedWithMeteredMinimumPriceCadenceMonthly, PlanNewParamsPricesLicenseAllocationPriceNewLicenseAllocationGroupedWithMeteredMinimumPriceCadenceQuarterly, PlanNewParamsPricesLicenseAllocationPriceNewLicenseAllocationGroupedWithMeteredMinimumPriceCadenceOneTime, PlanNewParamsPricesLicenseAllocationPriceNewLicenseAllocationGroupedWithMeteredMinimumPriceCadenceCustom:
		return true
	}
	return false
}

// Configuration for grouped_with_metered_minimum pricing
type PlanNewParamsPricesLicenseAllocationPriceNewLicenseAllocationGroupedWithMeteredMinimumPriceGroupedWithMeteredMinimumConfig struct {
	// Used to partition the usage into groups. The minimum amount is applied to each
	// group.
	GroupingKey param.Field[string] `json:"grouping_key" api:"required"`
	// The minimum amount to charge per group per unit
	MinimumUnitAmount param.Field[string] `json:"minimum_unit_amount" api:"required"`
	// Used to determine the unit rate
	PricingKey param.Field[string] `json:"pricing_key" api:"required"`
	// Scale the unit rates by the scaling factor.
	ScalingFactors param.Field[[]PlanNewParamsPricesLicenseAllocationPriceNewLicenseAllocationGroupedWithMeteredMinimumPriceGroupedWithMeteredMinimumConfigScalingFactor] `json:"scaling_factors" api:"required"`
	// Used to determine the unit rate scaling factor
	ScalingKey param.Field[string] `json:"scaling_key" api:"required"`
	// Apply per unit pricing to each pricing value. The minimum amount is applied any
	// unmatched usage.
	UnitAmounts param.Field[[]PlanNewParamsPricesLicenseAllocationPriceNewLicenseAllocationGroupedWithMeteredMinimumPriceGroupedWithMeteredMinimumConfigUnitAmount] `json:"unit_amounts" api:"required"`
}

func (r PlanNewParamsPricesLicenseAllocationPriceNewLicenseAllocationGroupedWithMeteredMinimumPriceGroupedWithMeteredMinimumConfig) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Configuration for a scaling factor
type PlanNewParamsPricesLicenseAllocationPriceNewLicenseAllocationGroupedWithMeteredMinimumPriceGroupedWithMeteredMinimumConfigScalingFactor struct {
	ScalingFactor param.Field[string] `json:"scaling_factor" api:"required"`
	ScalingValue  param.Field[string] `json:"scaling_value" api:"required"`
}

func (r PlanNewParamsPricesLicenseAllocationPriceNewLicenseAllocationGroupedWithMeteredMinimumPriceGroupedWithMeteredMinimumConfigScalingFactor) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Configuration for a unit amount
type PlanNewParamsPricesLicenseAllocationPriceNewLicenseAllocationGroupedWithMeteredMinimumPriceGroupedWithMeteredMinimumConfigUnitAmount struct {
	PricingValue param.Field[string] `json:"pricing_value" api:"required"`
	// Per unit amount
	UnitAmount param.Field[string] `json:"unit_amount" api:"required"`
}

func (r PlanNewParamsPricesLicenseAllocationPriceNewLicenseAllocationGroupedWithMeteredMinimumPriceGroupedWithMeteredMinimumConfigUnitAmount) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type PlanNewParamsPricesLicenseAllocationPriceNewLicenseAllocationGroupedWithMeteredMinimumPriceLicenseAllocation struct {
	// The amount of credits granted per active license per cadence.
	Amount param.Field[string] `json:"amount" api:"required"`
	// The currency of the license allocation.
	Currency param.Field[string] `json:"currency" api:"required"`
	// When True, overage beyond the allocation is written off.
	WriteOffOverage param.Field[bool] `json:"write_off_overage"`
}

func (r PlanNewParamsPricesLicenseAllocationPriceNewLicenseAllocationGroupedWithMeteredMinimumPriceLicenseAllocation) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// The pricing model type
type PlanNewParamsPricesLicenseAllocationPriceNewLicenseAllocationGroupedWithMeteredMinimumPriceModelType string

const (
	PlanNewParamsPricesLicenseAllocationPriceNewLicenseAllocationGroupedWithMeteredMinimumPriceModelTypeGroupedWithMeteredMinimum PlanNewParamsPricesLicenseAllocationPriceNewLicenseAllocationGroupedWithMeteredMinimumPriceModelType = "grouped_with_metered_minimum"
)

func (r PlanNewParamsPricesLicenseAllocationPriceNewLicenseAllocationGroupedWithMeteredMinimumPriceModelType) IsKnown() bool {
	switch r {
	case PlanNewParamsPricesLicenseAllocationPriceNewLicenseAllocationGroupedWithMeteredMinimumPriceModelTypeGroupedWithMeteredMinimum:
		return true
	}
	return false
}

type PlanNewParamsPricesLicenseAllocationPriceNewLicenseAllocationGroupedWithMeteredMinimumPriceConversionRateConfig struct {
	ConversionRateType param.Field[PlanNewParamsPricesLicenseAllocationPriceNewLicenseAllocationGroupedWithMeteredMinimumPriceConversionRateConfigConversionRateType] `json:"conversion_rate_type" api:"required"`
	TieredConfig       param.Field[shared.ConversionRateTieredConfigParam]                                                                                            `json:"tiered_config"`
	UnitConfig         param.Field[shared.ConversionRateUnitConfigParam]                                                                                              `json:"unit_config"`
}

func (r PlanNewParamsPricesLicenseAllocationPriceNewLicenseAllocationGroupedWithMeteredMinimumPriceConversionRateConfig) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r PlanNewParamsPricesLicenseAllocationPriceNewLicenseAllocationGroupedWithMeteredMinimumPriceConversionRateConfig) ImplementsPlanNewParamsPricesLicenseAllocationPriceNewLicenseAllocationGroupedWithMeteredMinimumPriceConversionRateConfigUnion() {
}

// Satisfied by [shared.UnitConversionRateConfigParam],
// [shared.TieredConversionRateConfigParam],
// [PlanNewParamsPricesLicenseAllocationPriceNewLicenseAllocationGroupedWithMeteredMinimumPriceConversionRateConfig].
type PlanNewParamsPricesLicenseAllocationPriceNewLicenseAllocationGroupedWithMeteredMinimumPriceConversionRateConfigUnion interface {
	ImplementsPlanNewParamsPricesLicenseAllocationPriceNewLicenseAllocationGroupedWithMeteredMinimumPriceConversionRateConfigUnion()
}

type PlanNewParamsPricesLicenseAllocationPriceNewLicenseAllocationGroupedWithMeteredMinimumPriceConversionRateConfigConversionRateType string

const (
	PlanNewParamsPricesLicenseAllocationPriceNewLicenseAllocationGroupedWithMeteredMinimumPriceConversionRateConfigConversionRateTypeUnit   PlanNewParamsPricesLicenseAllocationPriceNewLicenseAllocationGroupedWithMeteredMinimumPriceConversionRateConfigConversionRateType = "unit"
	PlanNewParamsPricesLicenseAllocationPriceNewLicenseAllocationGroupedWithMeteredMinimumPriceConversionRateConfigConversionRateTypeTiered PlanNewParamsPricesLicenseAllocationPriceNewLicenseAllocationGroupedWithMeteredMinimumPriceConversionRateConfigConversionRateType = "tiered"
)

func (r PlanNewParamsPricesLicenseAllocationPriceNewLicenseAllocationGroupedWithMeteredMinimumPriceConversionRateConfigConversionRateType) IsKnown() bool {
	switch r {
	case PlanNewParamsPricesLicenseAllocationPriceNewLicenseAllocationGroupedWithMeteredMinimumPriceConversionRateConfigConversionRateTypeUnit, PlanNewParamsPricesLicenseAllocationPriceNewLicenseAllocationGroupedWithMeteredMinimumPriceConversionRateConfigConversionRateTypeTiered:
		return true
	}
	return false
}

type PlanNewParamsPricesLicenseAllocationPriceNewLicenseAllocationGroupedWithMinMaxThresholdsPrice struct {
	// The cadence to bill for this price on.
	Cadence param.Field[PlanNewParamsPricesLicenseAllocationPriceNewLicenseAllocationGroupedWithMinMaxThresholdsPriceCadence] `json:"cadence" api:"required"`
	// Configuration for grouped_with_min_max_thresholds pricing
	GroupedWithMinMaxThresholdsConfig param.Field[PlanNewParamsPricesLicenseAllocationPriceNewLicenseAllocationGroupedWithMinMaxThresholdsPriceGroupedWithMinMaxThresholdsConfig] `json:"grouped_with_min_max_thresholds_config" api:"required"`
	// The id of the item the price will be associated with.
	ItemID param.Field[string] `json:"item_id" api:"required"`
	// License allocations to associate with this price. Each entry defines a
	// per-license credit pool granted each cadence. Requires license_type_id or
	// license_type_configuration to be set.
	LicenseAllocations param.Field[[]PlanNewParamsPricesLicenseAllocationPriceNewLicenseAllocationGroupedWithMinMaxThresholdsPriceLicenseAllocation] `json:"license_allocations" api:"required"`
	// The pricing model type
	ModelType param.Field[PlanNewParamsPricesLicenseAllocationPriceNewLicenseAllocationGroupedWithMinMaxThresholdsPriceModelType] `json:"model_type" api:"required"`
	// The name of the price.
	Name param.Field[string] `json:"name" api:"required"`
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
	ConversionRateConfig param.Field[PlanNewParamsPricesLicenseAllocationPriceNewLicenseAllocationGroupedWithMinMaxThresholdsPriceConversionRateConfigUnion] `json:"conversion_rate_config"`
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
	// The ID of the license type to associate with this price.
	LicenseTypeID param.Field[string] `json:"license_type_id"`
	// User-specified key/value pairs for the resource. Individual keys can be removed
	// by setting the value to `null`, and the entire metadata mapping can be cleared
	// by setting `metadata` to `null`.
	Metadata param.Field[map[string]string] `json:"metadata"`
	// A transient ID that can be used to reference this price when adding adjustments
	// in the same API call.
	ReferenceID param.Field[string] `json:"reference_id"`
}

func (r PlanNewParamsPricesLicenseAllocationPriceNewLicenseAllocationGroupedWithMinMaxThresholdsPrice) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r PlanNewParamsPricesLicenseAllocationPriceNewLicenseAllocationGroupedWithMinMaxThresholdsPrice) implementsPlanNewParamsPricesLicenseAllocationPriceUnion() {
}

// The cadence to bill for this price on.
type PlanNewParamsPricesLicenseAllocationPriceNewLicenseAllocationGroupedWithMinMaxThresholdsPriceCadence string

const (
	PlanNewParamsPricesLicenseAllocationPriceNewLicenseAllocationGroupedWithMinMaxThresholdsPriceCadenceAnnual     PlanNewParamsPricesLicenseAllocationPriceNewLicenseAllocationGroupedWithMinMaxThresholdsPriceCadence = "annual"
	PlanNewParamsPricesLicenseAllocationPriceNewLicenseAllocationGroupedWithMinMaxThresholdsPriceCadenceSemiAnnual PlanNewParamsPricesLicenseAllocationPriceNewLicenseAllocationGroupedWithMinMaxThresholdsPriceCadence = "semi_annual"
	PlanNewParamsPricesLicenseAllocationPriceNewLicenseAllocationGroupedWithMinMaxThresholdsPriceCadenceMonthly    PlanNewParamsPricesLicenseAllocationPriceNewLicenseAllocationGroupedWithMinMaxThresholdsPriceCadence = "monthly"
	PlanNewParamsPricesLicenseAllocationPriceNewLicenseAllocationGroupedWithMinMaxThresholdsPriceCadenceQuarterly  PlanNewParamsPricesLicenseAllocationPriceNewLicenseAllocationGroupedWithMinMaxThresholdsPriceCadence = "quarterly"
	PlanNewParamsPricesLicenseAllocationPriceNewLicenseAllocationGroupedWithMinMaxThresholdsPriceCadenceOneTime    PlanNewParamsPricesLicenseAllocationPriceNewLicenseAllocationGroupedWithMinMaxThresholdsPriceCadence = "one_time"
	PlanNewParamsPricesLicenseAllocationPriceNewLicenseAllocationGroupedWithMinMaxThresholdsPriceCadenceCustom     PlanNewParamsPricesLicenseAllocationPriceNewLicenseAllocationGroupedWithMinMaxThresholdsPriceCadence = "custom"
)

func (r PlanNewParamsPricesLicenseAllocationPriceNewLicenseAllocationGroupedWithMinMaxThresholdsPriceCadence) IsKnown() bool {
	switch r {
	case PlanNewParamsPricesLicenseAllocationPriceNewLicenseAllocationGroupedWithMinMaxThresholdsPriceCadenceAnnual, PlanNewParamsPricesLicenseAllocationPriceNewLicenseAllocationGroupedWithMinMaxThresholdsPriceCadenceSemiAnnual, PlanNewParamsPricesLicenseAllocationPriceNewLicenseAllocationGroupedWithMinMaxThresholdsPriceCadenceMonthly, PlanNewParamsPricesLicenseAllocationPriceNewLicenseAllocationGroupedWithMinMaxThresholdsPriceCadenceQuarterly, PlanNewParamsPricesLicenseAllocationPriceNewLicenseAllocationGroupedWithMinMaxThresholdsPriceCadenceOneTime, PlanNewParamsPricesLicenseAllocationPriceNewLicenseAllocationGroupedWithMinMaxThresholdsPriceCadenceCustom:
		return true
	}
	return false
}

// Configuration for grouped_with_min_max_thresholds pricing
type PlanNewParamsPricesLicenseAllocationPriceNewLicenseAllocationGroupedWithMinMaxThresholdsPriceGroupedWithMinMaxThresholdsConfig struct {
	// The event property used to group before applying thresholds
	GroupingKey param.Field[string] `json:"grouping_key" api:"required"`
	// The maximum amount to charge each group
	MaximumCharge param.Field[string] `json:"maximum_charge" api:"required"`
	// The minimum amount to charge each group, regardless of usage
	MinimumCharge param.Field[string] `json:"minimum_charge" api:"required"`
	// The base price charged per group
	PerUnitRate param.Field[string] `json:"per_unit_rate" api:"required"`
}

func (r PlanNewParamsPricesLicenseAllocationPriceNewLicenseAllocationGroupedWithMinMaxThresholdsPriceGroupedWithMinMaxThresholdsConfig) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type PlanNewParamsPricesLicenseAllocationPriceNewLicenseAllocationGroupedWithMinMaxThresholdsPriceLicenseAllocation struct {
	// The amount of credits granted per active license per cadence.
	Amount param.Field[string] `json:"amount" api:"required"`
	// The currency of the license allocation.
	Currency param.Field[string] `json:"currency" api:"required"`
	// When True, overage beyond the allocation is written off.
	WriteOffOverage param.Field[bool] `json:"write_off_overage"`
}

func (r PlanNewParamsPricesLicenseAllocationPriceNewLicenseAllocationGroupedWithMinMaxThresholdsPriceLicenseAllocation) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// The pricing model type
type PlanNewParamsPricesLicenseAllocationPriceNewLicenseAllocationGroupedWithMinMaxThresholdsPriceModelType string

const (
	PlanNewParamsPricesLicenseAllocationPriceNewLicenseAllocationGroupedWithMinMaxThresholdsPriceModelTypeGroupedWithMinMaxThresholds PlanNewParamsPricesLicenseAllocationPriceNewLicenseAllocationGroupedWithMinMaxThresholdsPriceModelType = "grouped_with_min_max_thresholds"
)

func (r PlanNewParamsPricesLicenseAllocationPriceNewLicenseAllocationGroupedWithMinMaxThresholdsPriceModelType) IsKnown() bool {
	switch r {
	case PlanNewParamsPricesLicenseAllocationPriceNewLicenseAllocationGroupedWithMinMaxThresholdsPriceModelTypeGroupedWithMinMaxThresholds:
		return true
	}
	return false
}

type PlanNewParamsPricesLicenseAllocationPriceNewLicenseAllocationGroupedWithMinMaxThresholdsPriceConversionRateConfig struct {
	ConversionRateType param.Field[PlanNewParamsPricesLicenseAllocationPriceNewLicenseAllocationGroupedWithMinMaxThresholdsPriceConversionRateConfigConversionRateType] `json:"conversion_rate_type" api:"required"`
	TieredConfig       param.Field[shared.ConversionRateTieredConfigParam]                                                                                              `json:"tiered_config"`
	UnitConfig         param.Field[shared.ConversionRateUnitConfigParam]                                                                                                `json:"unit_config"`
}

func (r PlanNewParamsPricesLicenseAllocationPriceNewLicenseAllocationGroupedWithMinMaxThresholdsPriceConversionRateConfig) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r PlanNewParamsPricesLicenseAllocationPriceNewLicenseAllocationGroupedWithMinMaxThresholdsPriceConversionRateConfig) ImplementsPlanNewParamsPricesLicenseAllocationPriceNewLicenseAllocationGroupedWithMinMaxThresholdsPriceConversionRateConfigUnion() {
}

// Satisfied by [shared.UnitConversionRateConfigParam],
// [shared.TieredConversionRateConfigParam],
// [PlanNewParamsPricesLicenseAllocationPriceNewLicenseAllocationGroupedWithMinMaxThresholdsPriceConversionRateConfig].
type PlanNewParamsPricesLicenseAllocationPriceNewLicenseAllocationGroupedWithMinMaxThresholdsPriceConversionRateConfigUnion interface {
	ImplementsPlanNewParamsPricesLicenseAllocationPriceNewLicenseAllocationGroupedWithMinMaxThresholdsPriceConversionRateConfigUnion()
}

type PlanNewParamsPricesLicenseAllocationPriceNewLicenseAllocationGroupedWithMinMaxThresholdsPriceConversionRateConfigConversionRateType string

const (
	PlanNewParamsPricesLicenseAllocationPriceNewLicenseAllocationGroupedWithMinMaxThresholdsPriceConversionRateConfigConversionRateTypeUnit   PlanNewParamsPricesLicenseAllocationPriceNewLicenseAllocationGroupedWithMinMaxThresholdsPriceConversionRateConfigConversionRateType = "unit"
	PlanNewParamsPricesLicenseAllocationPriceNewLicenseAllocationGroupedWithMinMaxThresholdsPriceConversionRateConfigConversionRateTypeTiered PlanNewParamsPricesLicenseAllocationPriceNewLicenseAllocationGroupedWithMinMaxThresholdsPriceConversionRateConfigConversionRateType = "tiered"
)

func (r PlanNewParamsPricesLicenseAllocationPriceNewLicenseAllocationGroupedWithMinMaxThresholdsPriceConversionRateConfigConversionRateType) IsKnown() bool {
	switch r {
	case PlanNewParamsPricesLicenseAllocationPriceNewLicenseAllocationGroupedWithMinMaxThresholdsPriceConversionRateConfigConversionRateTypeUnit, PlanNewParamsPricesLicenseAllocationPriceNewLicenseAllocationGroupedWithMinMaxThresholdsPriceConversionRateConfigConversionRateTypeTiered:
		return true
	}
	return false
}

type PlanNewParamsPricesLicenseAllocationPriceNewLicenseAllocationMatrixWithDisplayNamePrice struct {
	// The cadence to bill for this price on.
	Cadence param.Field[PlanNewParamsPricesLicenseAllocationPriceNewLicenseAllocationMatrixWithDisplayNamePriceCadence] `json:"cadence" api:"required"`
	// The id of the item the price will be associated with.
	ItemID param.Field[string] `json:"item_id" api:"required"`
	// License allocations to associate with this price. Each entry defines a
	// per-license credit pool granted each cadence. Requires license_type_id or
	// license_type_configuration to be set.
	LicenseAllocations param.Field[[]PlanNewParamsPricesLicenseAllocationPriceNewLicenseAllocationMatrixWithDisplayNamePriceLicenseAllocation] `json:"license_allocations" api:"required"`
	// Configuration for matrix_with_display_name pricing
	MatrixWithDisplayNameConfig param.Field[PlanNewParamsPricesLicenseAllocationPriceNewLicenseAllocationMatrixWithDisplayNamePriceMatrixWithDisplayNameConfig] `json:"matrix_with_display_name_config" api:"required"`
	// The pricing model type
	ModelType param.Field[PlanNewParamsPricesLicenseAllocationPriceNewLicenseAllocationMatrixWithDisplayNamePriceModelType] `json:"model_type" api:"required"`
	// The name of the price.
	Name param.Field[string] `json:"name" api:"required"`
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
	ConversionRateConfig param.Field[PlanNewParamsPricesLicenseAllocationPriceNewLicenseAllocationMatrixWithDisplayNamePriceConversionRateConfigUnion] `json:"conversion_rate_config"`
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
	// The ID of the license type to associate with this price.
	LicenseTypeID param.Field[string] `json:"license_type_id"`
	// User-specified key/value pairs for the resource. Individual keys can be removed
	// by setting the value to `null`, and the entire metadata mapping can be cleared
	// by setting `metadata` to `null`.
	Metadata param.Field[map[string]string] `json:"metadata"`
	// A transient ID that can be used to reference this price when adding adjustments
	// in the same API call.
	ReferenceID param.Field[string] `json:"reference_id"`
}

func (r PlanNewParamsPricesLicenseAllocationPriceNewLicenseAllocationMatrixWithDisplayNamePrice) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r PlanNewParamsPricesLicenseAllocationPriceNewLicenseAllocationMatrixWithDisplayNamePrice) implementsPlanNewParamsPricesLicenseAllocationPriceUnion() {
}

// The cadence to bill for this price on.
type PlanNewParamsPricesLicenseAllocationPriceNewLicenseAllocationMatrixWithDisplayNamePriceCadence string

const (
	PlanNewParamsPricesLicenseAllocationPriceNewLicenseAllocationMatrixWithDisplayNamePriceCadenceAnnual     PlanNewParamsPricesLicenseAllocationPriceNewLicenseAllocationMatrixWithDisplayNamePriceCadence = "annual"
	PlanNewParamsPricesLicenseAllocationPriceNewLicenseAllocationMatrixWithDisplayNamePriceCadenceSemiAnnual PlanNewParamsPricesLicenseAllocationPriceNewLicenseAllocationMatrixWithDisplayNamePriceCadence = "semi_annual"
	PlanNewParamsPricesLicenseAllocationPriceNewLicenseAllocationMatrixWithDisplayNamePriceCadenceMonthly    PlanNewParamsPricesLicenseAllocationPriceNewLicenseAllocationMatrixWithDisplayNamePriceCadence = "monthly"
	PlanNewParamsPricesLicenseAllocationPriceNewLicenseAllocationMatrixWithDisplayNamePriceCadenceQuarterly  PlanNewParamsPricesLicenseAllocationPriceNewLicenseAllocationMatrixWithDisplayNamePriceCadence = "quarterly"
	PlanNewParamsPricesLicenseAllocationPriceNewLicenseAllocationMatrixWithDisplayNamePriceCadenceOneTime    PlanNewParamsPricesLicenseAllocationPriceNewLicenseAllocationMatrixWithDisplayNamePriceCadence = "one_time"
	PlanNewParamsPricesLicenseAllocationPriceNewLicenseAllocationMatrixWithDisplayNamePriceCadenceCustom     PlanNewParamsPricesLicenseAllocationPriceNewLicenseAllocationMatrixWithDisplayNamePriceCadence = "custom"
)

func (r PlanNewParamsPricesLicenseAllocationPriceNewLicenseAllocationMatrixWithDisplayNamePriceCadence) IsKnown() bool {
	switch r {
	case PlanNewParamsPricesLicenseAllocationPriceNewLicenseAllocationMatrixWithDisplayNamePriceCadenceAnnual, PlanNewParamsPricesLicenseAllocationPriceNewLicenseAllocationMatrixWithDisplayNamePriceCadenceSemiAnnual, PlanNewParamsPricesLicenseAllocationPriceNewLicenseAllocationMatrixWithDisplayNamePriceCadenceMonthly, PlanNewParamsPricesLicenseAllocationPriceNewLicenseAllocationMatrixWithDisplayNamePriceCadenceQuarterly, PlanNewParamsPricesLicenseAllocationPriceNewLicenseAllocationMatrixWithDisplayNamePriceCadenceOneTime, PlanNewParamsPricesLicenseAllocationPriceNewLicenseAllocationMatrixWithDisplayNamePriceCadenceCustom:
		return true
	}
	return false
}

type PlanNewParamsPricesLicenseAllocationPriceNewLicenseAllocationMatrixWithDisplayNamePriceLicenseAllocation struct {
	// The amount of credits granted per active license per cadence.
	Amount param.Field[string] `json:"amount" api:"required"`
	// The currency of the license allocation.
	Currency param.Field[string] `json:"currency" api:"required"`
	// When True, overage beyond the allocation is written off.
	WriteOffOverage param.Field[bool] `json:"write_off_overage"`
}

func (r PlanNewParamsPricesLicenseAllocationPriceNewLicenseAllocationMatrixWithDisplayNamePriceLicenseAllocation) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Configuration for matrix_with_display_name pricing
type PlanNewParamsPricesLicenseAllocationPriceNewLicenseAllocationMatrixWithDisplayNamePriceMatrixWithDisplayNameConfig struct {
	// Used to determine the unit rate
	Dimension param.Field[string] `json:"dimension" api:"required"`
	// Apply per unit pricing to each dimension value
	UnitAmounts param.Field[[]PlanNewParamsPricesLicenseAllocationPriceNewLicenseAllocationMatrixWithDisplayNamePriceMatrixWithDisplayNameConfigUnitAmount] `json:"unit_amounts" api:"required"`
}

func (r PlanNewParamsPricesLicenseAllocationPriceNewLicenseAllocationMatrixWithDisplayNamePriceMatrixWithDisplayNameConfig) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Configuration for a unit amount item
type PlanNewParamsPricesLicenseAllocationPriceNewLicenseAllocationMatrixWithDisplayNamePriceMatrixWithDisplayNameConfigUnitAmount struct {
	// The dimension value
	DimensionValue param.Field[string] `json:"dimension_value" api:"required"`
	// Display name for this dimension value
	DisplayName param.Field[string] `json:"display_name" api:"required"`
	// Per unit amount
	UnitAmount param.Field[string] `json:"unit_amount" api:"required"`
}

func (r PlanNewParamsPricesLicenseAllocationPriceNewLicenseAllocationMatrixWithDisplayNamePriceMatrixWithDisplayNameConfigUnitAmount) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// The pricing model type
type PlanNewParamsPricesLicenseAllocationPriceNewLicenseAllocationMatrixWithDisplayNamePriceModelType string

const (
	PlanNewParamsPricesLicenseAllocationPriceNewLicenseAllocationMatrixWithDisplayNamePriceModelTypeMatrixWithDisplayName PlanNewParamsPricesLicenseAllocationPriceNewLicenseAllocationMatrixWithDisplayNamePriceModelType = "matrix_with_display_name"
)

func (r PlanNewParamsPricesLicenseAllocationPriceNewLicenseAllocationMatrixWithDisplayNamePriceModelType) IsKnown() bool {
	switch r {
	case PlanNewParamsPricesLicenseAllocationPriceNewLicenseAllocationMatrixWithDisplayNamePriceModelTypeMatrixWithDisplayName:
		return true
	}
	return false
}

type PlanNewParamsPricesLicenseAllocationPriceNewLicenseAllocationMatrixWithDisplayNamePriceConversionRateConfig struct {
	ConversionRateType param.Field[PlanNewParamsPricesLicenseAllocationPriceNewLicenseAllocationMatrixWithDisplayNamePriceConversionRateConfigConversionRateType] `json:"conversion_rate_type" api:"required"`
	TieredConfig       param.Field[shared.ConversionRateTieredConfigParam]                                                                                        `json:"tiered_config"`
	UnitConfig         param.Field[shared.ConversionRateUnitConfigParam]                                                                                          `json:"unit_config"`
}

func (r PlanNewParamsPricesLicenseAllocationPriceNewLicenseAllocationMatrixWithDisplayNamePriceConversionRateConfig) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r PlanNewParamsPricesLicenseAllocationPriceNewLicenseAllocationMatrixWithDisplayNamePriceConversionRateConfig) ImplementsPlanNewParamsPricesLicenseAllocationPriceNewLicenseAllocationMatrixWithDisplayNamePriceConversionRateConfigUnion() {
}

// Satisfied by [shared.UnitConversionRateConfigParam],
// [shared.TieredConversionRateConfigParam],
// [PlanNewParamsPricesLicenseAllocationPriceNewLicenseAllocationMatrixWithDisplayNamePriceConversionRateConfig].
type PlanNewParamsPricesLicenseAllocationPriceNewLicenseAllocationMatrixWithDisplayNamePriceConversionRateConfigUnion interface {
	ImplementsPlanNewParamsPricesLicenseAllocationPriceNewLicenseAllocationMatrixWithDisplayNamePriceConversionRateConfigUnion()
}

type PlanNewParamsPricesLicenseAllocationPriceNewLicenseAllocationMatrixWithDisplayNamePriceConversionRateConfigConversionRateType string

const (
	PlanNewParamsPricesLicenseAllocationPriceNewLicenseAllocationMatrixWithDisplayNamePriceConversionRateConfigConversionRateTypeUnit   PlanNewParamsPricesLicenseAllocationPriceNewLicenseAllocationMatrixWithDisplayNamePriceConversionRateConfigConversionRateType = "unit"
	PlanNewParamsPricesLicenseAllocationPriceNewLicenseAllocationMatrixWithDisplayNamePriceConversionRateConfigConversionRateTypeTiered PlanNewParamsPricesLicenseAllocationPriceNewLicenseAllocationMatrixWithDisplayNamePriceConversionRateConfigConversionRateType = "tiered"
)

func (r PlanNewParamsPricesLicenseAllocationPriceNewLicenseAllocationMatrixWithDisplayNamePriceConversionRateConfigConversionRateType) IsKnown() bool {
	switch r {
	case PlanNewParamsPricesLicenseAllocationPriceNewLicenseAllocationMatrixWithDisplayNamePriceConversionRateConfigConversionRateTypeUnit, PlanNewParamsPricesLicenseAllocationPriceNewLicenseAllocationMatrixWithDisplayNamePriceConversionRateConfigConversionRateTypeTiered:
		return true
	}
	return false
}

type PlanNewParamsPricesLicenseAllocationPriceNewLicenseAllocationGroupedTieredPackagePrice struct {
	// The cadence to bill for this price on.
	Cadence param.Field[PlanNewParamsPricesLicenseAllocationPriceNewLicenseAllocationGroupedTieredPackagePriceCadence] `json:"cadence" api:"required"`
	// Configuration for grouped_tiered_package pricing
	GroupedTieredPackageConfig param.Field[PlanNewParamsPricesLicenseAllocationPriceNewLicenseAllocationGroupedTieredPackagePriceGroupedTieredPackageConfig] `json:"grouped_tiered_package_config" api:"required"`
	// The id of the item the price will be associated with.
	ItemID param.Field[string] `json:"item_id" api:"required"`
	// License allocations to associate with this price. Each entry defines a
	// per-license credit pool granted each cadence. Requires license_type_id or
	// license_type_configuration to be set.
	LicenseAllocations param.Field[[]PlanNewParamsPricesLicenseAllocationPriceNewLicenseAllocationGroupedTieredPackagePriceLicenseAllocation] `json:"license_allocations" api:"required"`
	// The pricing model type
	ModelType param.Field[PlanNewParamsPricesLicenseAllocationPriceNewLicenseAllocationGroupedTieredPackagePriceModelType] `json:"model_type" api:"required"`
	// The name of the price.
	Name param.Field[string] `json:"name" api:"required"`
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
	ConversionRateConfig param.Field[PlanNewParamsPricesLicenseAllocationPriceNewLicenseAllocationGroupedTieredPackagePriceConversionRateConfigUnion] `json:"conversion_rate_config"`
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
	// The ID of the license type to associate with this price.
	LicenseTypeID param.Field[string] `json:"license_type_id"`
	// User-specified key/value pairs for the resource. Individual keys can be removed
	// by setting the value to `null`, and the entire metadata mapping can be cleared
	// by setting `metadata` to `null`.
	Metadata param.Field[map[string]string] `json:"metadata"`
	// A transient ID that can be used to reference this price when adding adjustments
	// in the same API call.
	ReferenceID param.Field[string] `json:"reference_id"`
}

func (r PlanNewParamsPricesLicenseAllocationPriceNewLicenseAllocationGroupedTieredPackagePrice) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r PlanNewParamsPricesLicenseAllocationPriceNewLicenseAllocationGroupedTieredPackagePrice) implementsPlanNewParamsPricesLicenseAllocationPriceUnion() {
}

// The cadence to bill for this price on.
type PlanNewParamsPricesLicenseAllocationPriceNewLicenseAllocationGroupedTieredPackagePriceCadence string

const (
	PlanNewParamsPricesLicenseAllocationPriceNewLicenseAllocationGroupedTieredPackagePriceCadenceAnnual     PlanNewParamsPricesLicenseAllocationPriceNewLicenseAllocationGroupedTieredPackagePriceCadence = "annual"
	PlanNewParamsPricesLicenseAllocationPriceNewLicenseAllocationGroupedTieredPackagePriceCadenceSemiAnnual PlanNewParamsPricesLicenseAllocationPriceNewLicenseAllocationGroupedTieredPackagePriceCadence = "semi_annual"
	PlanNewParamsPricesLicenseAllocationPriceNewLicenseAllocationGroupedTieredPackagePriceCadenceMonthly    PlanNewParamsPricesLicenseAllocationPriceNewLicenseAllocationGroupedTieredPackagePriceCadence = "monthly"
	PlanNewParamsPricesLicenseAllocationPriceNewLicenseAllocationGroupedTieredPackagePriceCadenceQuarterly  PlanNewParamsPricesLicenseAllocationPriceNewLicenseAllocationGroupedTieredPackagePriceCadence = "quarterly"
	PlanNewParamsPricesLicenseAllocationPriceNewLicenseAllocationGroupedTieredPackagePriceCadenceOneTime    PlanNewParamsPricesLicenseAllocationPriceNewLicenseAllocationGroupedTieredPackagePriceCadence = "one_time"
	PlanNewParamsPricesLicenseAllocationPriceNewLicenseAllocationGroupedTieredPackagePriceCadenceCustom     PlanNewParamsPricesLicenseAllocationPriceNewLicenseAllocationGroupedTieredPackagePriceCadence = "custom"
)

func (r PlanNewParamsPricesLicenseAllocationPriceNewLicenseAllocationGroupedTieredPackagePriceCadence) IsKnown() bool {
	switch r {
	case PlanNewParamsPricesLicenseAllocationPriceNewLicenseAllocationGroupedTieredPackagePriceCadenceAnnual, PlanNewParamsPricesLicenseAllocationPriceNewLicenseAllocationGroupedTieredPackagePriceCadenceSemiAnnual, PlanNewParamsPricesLicenseAllocationPriceNewLicenseAllocationGroupedTieredPackagePriceCadenceMonthly, PlanNewParamsPricesLicenseAllocationPriceNewLicenseAllocationGroupedTieredPackagePriceCadenceQuarterly, PlanNewParamsPricesLicenseAllocationPriceNewLicenseAllocationGroupedTieredPackagePriceCadenceOneTime, PlanNewParamsPricesLicenseAllocationPriceNewLicenseAllocationGroupedTieredPackagePriceCadenceCustom:
		return true
	}
	return false
}

// Configuration for grouped_tiered_package pricing
type PlanNewParamsPricesLicenseAllocationPriceNewLicenseAllocationGroupedTieredPackagePriceGroupedTieredPackageConfig struct {
	// The event property used to group before tiering
	GroupingKey param.Field[string] `json:"grouping_key" api:"required"`
	PackageSize param.Field[string] `json:"package_size" api:"required"`
	// Apply tiered pricing after rounding up the quantity to the package size. Tiers
	// are defined using exclusive lower bounds.
	Tiers param.Field[[]PlanNewParamsPricesLicenseAllocationPriceNewLicenseAllocationGroupedTieredPackagePriceGroupedTieredPackageConfigTier] `json:"tiers" api:"required"`
}

func (r PlanNewParamsPricesLicenseAllocationPriceNewLicenseAllocationGroupedTieredPackagePriceGroupedTieredPackageConfig) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Configuration for a single tier
type PlanNewParamsPricesLicenseAllocationPriceNewLicenseAllocationGroupedTieredPackagePriceGroupedTieredPackageConfigTier struct {
	// Per package
	PerUnit        param.Field[string] `json:"per_unit" api:"required"`
	TierLowerBound param.Field[string] `json:"tier_lower_bound" api:"required"`
}

func (r PlanNewParamsPricesLicenseAllocationPriceNewLicenseAllocationGroupedTieredPackagePriceGroupedTieredPackageConfigTier) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type PlanNewParamsPricesLicenseAllocationPriceNewLicenseAllocationGroupedTieredPackagePriceLicenseAllocation struct {
	// The amount of credits granted per active license per cadence.
	Amount param.Field[string] `json:"amount" api:"required"`
	// The currency of the license allocation.
	Currency param.Field[string] `json:"currency" api:"required"`
	// When True, overage beyond the allocation is written off.
	WriteOffOverage param.Field[bool] `json:"write_off_overage"`
}

func (r PlanNewParamsPricesLicenseAllocationPriceNewLicenseAllocationGroupedTieredPackagePriceLicenseAllocation) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// The pricing model type
type PlanNewParamsPricesLicenseAllocationPriceNewLicenseAllocationGroupedTieredPackagePriceModelType string

const (
	PlanNewParamsPricesLicenseAllocationPriceNewLicenseAllocationGroupedTieredPackagePriceModelTypeGroupedTieredPackage PlanNewParamsPricesLicenseAllocationPriceNewLicenseAllocationGroupedTieredPackagePriceModelType = "grouped_tiered_package"
)

func (r PlanNewParamsPricesLicenseAllocationPriceNewLicenseAllocationGroupedTieredPackagePriceModelType) IsKnown() bool {
	switch r {
	case PlanNewParamsPricesLicenseAllocationPriceNewLicenseAllocationGroupedTieredPackagePriceModelTypeGroupedTieredPackage:
		return true
	}
	return false
}

type PlanNewParamsPricesLicenseAllocationPriceNewLicenseAllocationGroupedTieredPackagePriceConversionRateConfig struct {
	ConversionRateType param.Field[PlanNewParamsPricesLicenseAllocationPriceNewLicenseAllocationGroupedTieredPackagePriceConversionRateConfigConversionRateType] `json:"conversion_rate_type" api:"required"`
	TieredConfig       param.Field[shared.ConversionRateTieredConfigParam]                                                                                       `json:"tiered_config"`
	UnitConfig         param.Field[shared.ConversionRateUnitConfigParam]                                                                                         `json:"unit_config"`
}

func (r PlanNewParamsPricesLicenseAllocationPriceNewLicenseAllocationGroupedTieredPackagePriceConversionRateConfig) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r PlanNewParamsPricesLicenseAllocationPriceNewLicenseAllocationGroupedTieredPackagePriceConversionRateConfig) ImplementsPlanNewParamsPricesLicenseAllocationPriceNewLicenseAllocationGroupedTieredPackagePriceConversionRateConfigUnion() {
}

// Satisfied by [shared.UnitConversionRateConfigParam],
// [shared.TieredConversionRateConfigParam],
// [PlanNewParamsPricesLicenseAllocationPriceNewLicenseAllocationGroupedTieredPackagePriceConversionRateConfig].
type PlanNewParamsPricesLicenseAllocationPriceNewLicenseAllocationGroupedTieredPackagePriceConversionRateConfigUnion interface {
	ImplementsPlanNewParamsPricesLicenseAllocationPriceNewLicenseAllocationGroupedTieredPackagePriceConversionRateConfigUnion()
}

type PlanNewParamsPricesLicenseAllocationPriceNewLicenseAllocationGroupedTieredPackagePriceConversionRateConfigConversionRateType string

const (
	PlanNewParamsPricesLicenseAllocationPriceNewLicenseAllocationGroupedTieredPackagePriceConversionRateConfigConversionRateTypeUnit   PlanNewParamsPricesLicenseAllocationPriceNewLicenseAllocationGroupedTieredPackagePriceConversionRateConfigConversionRateType = "unit"
	PlanNewParamsPricesLicenseAllocationPriceNewLicenseAllocationGroupedTieredPackagePriceConversionRateConfigConversionRateTypeTiered PlanNewParamsPricesLicenseAllocationPriceNewLicenseAllocationGroupedTieredPackagePriceConversionRateConfigConversionRateType = "tiered"
)

func (r PlanNewParamsPricesLicenseAllocationPriceNewLicenseAllocationGroupedTieredPackagePriceConversionRateConfigConversionRateType) IsKnown() bool {
	switch r {
	case PlanNewParamsPricesLicenseAllocationPriceNewLicenseAllocationGroupedTieredPackagePriceConversionRateConfigConversionRateTypeUnit, PlanNewParamsPricesLicenseAllocationPriceNewLicenseAllocationGroupedTieredPackagePriceConversionRateConfigConversionRateTypeTiered:
		return true
	}
	return false
}

type PlanNewParamsPricesLicenseAllocationPriceNewLicenseAllocationMaxGroupTieredPackagePrice struct {
	// The cadence to bill for this price on.
	Cadence param.Field[PlanNewParamsPricesLicenseAllocationPriceNewLicenseAllocationMaxGroupTieredPackagePriceCadence] `json:"cadence" api:"required"`
	// The id of the item the price will be associated with.
	ItemID param.Field[string] `json:"item_id" api:"required"`
	// License allocations to associate with this price. Each entry defines a
	// per-license credit pool granted each cadence. Requires license_type_id or
	// license_type_configuration to be set.
	LicenseAllocations param.Field[[]PlanNewParamsPricesLicenseAllocationPriceNewLicenseAllocationMaxGroupTieredPackagePriceLicenseAllocation] `json:"license_allocations" api:"required"`
	// Configuration for max_group_tiered_package pricing
	MaxGroupTieredPackageConfig param.Field[PlanNewParamsPricesLicenseAllocationPriceNewLicenseAllocationMaxGroupTieredPackagePriceMaxGroupTieredPackageConfig] `json:"max_group_tiered_package_config" api:"required"`
	// The pricing model type
	ModelType param.Field[PlanNewParamsPricesLicenseAllocationPriceNewLicenseAllocationMaxGroupTieredPackagePriceModelType] `json:"model_type" api:"required"`
	// The name of the price.
	Name param.Field[string] `json:"name" api:"required"`
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
	ConversionRateConfig param.Field[PlanNewParamsPricesLicenseAllocationPriceNewLicenseAllocationMaxGroupTieredPackagePriceConversionRateConfigUnion] `json:"conversion_rate_config"`
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
	// The ID of the license type to associate with this price.
	LicenseTypeID param.Field[string] `json:"license_type_id"`
	// User-specified key/value pairs for the resource. Individual keys can be removed
	// by setting the value to `null`, and the entire metadata mapping can be cleared
	// by setting `metadata` to `null`.
	Metadata param.Field[map[string]string] `json:"metadata"`
	// A transient ID that can be used to reference this price when adding adjustments
	// in the same API call.
	ReferenceID param.Field[string] `json:"reference_id"`
}

func (r PlanNewParamsPricesLicenseAllocationPriceNewLicenseAllocationMaxGroupTieredPackagePrice) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r PlanNewParamsPricesLicenseAllocationPriceNewLicenseAllocationMaxGroupTieredPackagePrice) implementsPlanNewParamsPricesLicenseAllocationPriceUnion() {
}

// The cadence to bill for this price on.
type PlanNewParamsPricesLicenseAllocationPriceNewLicenseAllocationMaxGroupTieredPackagePriceCadence string

const (
	PlanNewParamsPricesLicenseAllocationPriceNewLicenseAllocationMaxGroupTieredPackagePriceCadenceAnnual     PlanNewParamsPricesLicenseAllocationPriceNewLicenseAllocationMaxGroupTieredPackagePriceCadence = "annual"
	PlanNewParamsPricesLicenseAllocationPriceNewLicenseAllocationMaxGroupTieredPackagePriceCadenceSemiAnnual PlanNewParamsPricesLicenseAllocationPriceNewLicenseAllocationMaxGroupTieredPackagePriceCadence = "semi_annual"
	PlanNewParamsPricesLicenseAllocationPriceNewLicenseAllocationMaxGroupTieredPackagePriceCadenceMonthly    PlanNewParamsPricesLicenseAllocationPriceNewLicenseAllocationMaxGroupTieredPackagePriceCadence = "monthly"
	PlanNewParamsPricesLicenseAllocationPriceNewLicenseAllocationMaxGroupTieredPackagePriceCadenceQuarterly  PlanNewParamsPricesLicenseAllocationPriceNewLicenseAllocationMaxGroupTieredPackagePriceCadence = "quarterly"
	PlanNewParamsPricesLicenseAllocationPriceNewLicenseAllocationMaxGroupTieredPackagePriceCadenceOneTime    PlanNewParamsPricesLicenseAllocationPriceNewLicenseAllocationMaxGroupTieredPackagePriceCadence = "one_time"
	PlanNewParamsPricesLicenseAllocationPriceNewLicenseAllocationMaxGroupTieredPackagePriceCadenceCustom     PlanNewParamsPricesLicenseAllocationPriceNewLicenseAllocationMaxGroupTieredPackagePriceCadence = "custom"
)

func (r PlanNewParamsPricesLicenseAllocationPriceNewLicenseAllocationMaxGroupTieredPackagePriceCadence) IsKnown() bool {
	switch r {
	case PlanNewParamsPricesLicenseAllocationPriceNewLicenseAllocationMaxGroupTieredPackagePriceCadenceAnnual, PlanNewParamsPricesLicenseAllocationPriceNewLicenseAllocationMaxGroupTieredPackagePriceCadenceSemiAnnual, PlanNewParamsPricesLicenseAllocationPriceNewLicenseAllocationMaxGroupTieredPackagePriceCadenceMonthly, PlanNewParamsPricesLicenseAllocationPriceNewLicenseAllocationMaxGroupTieredPackagePriceCadenceQuarterly, PlanNewParamsPricesLicenseAllocationPriceNewLicenseAllocationMaxGroupTieredPackagePriceCadenceOneTime, PlanNewParamsPricesLicenseAllocationPriceNewLicenseAllocationMaxGroupTieredPackagePriceCadenceCustom:
		return true
	}
	return false
}

type PlanNewParamsPricesLicenseAllocationPriceNewLicenseAllocationMaxGroupTieredPackagePriceLicenseAllocation struct {
	// The amount of credits granted per active license per cadence.
	Amount param.Field[string] `json:"amount" api:"required"`
	// The currency of the license allocation.
	Currency param.Field[string] `json:"currency" api:"required"`
	// When True, overage beyond the allocation is written off.
	WriteOffOverage param.Field[bool] `json:"write_off_overage"`
}

func (r PlanNewParamsPricesLicenseAllocationPriceNewLicenseAllocationMaxGroupTieredPackagePriceLicenseAllocation) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Configuration for max_group_tiered_package pricing
type PlanNewParamsPricesLicenseAllocationPriceNewLicenseAllocationMaxGroupTieredPackagePriceMaxGroupTieredPackageConfig struct {
	// The event property used to group before tiering the group with the highest value
	GroupingKey param.Field[string] `json:"grouping_key" api:"required"`
	PackageSize param.Field[string] `json:"package_size" api:"required"`
	// Apply tiered pricing to the largest group after grouping with the provided key.
	Tiers param.Field[[]PlanNewParamsPricesLicenseAllocationPriceNewLicenseAllocationMaxGroupTieredPackagePriceMaxGroupTieredPackageConfigTier] `json:"tiers" api:"required"`
}

func (r PlanNewParamsPricesLicenseAllocationPriceNewLicenseAllocationMaxGroupTieredPackagePriceMaxGroupTieredPackageConfig) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Configuration for a single tier
type PlanNewParamsPricesLicenseAllocationPriceNewLicenseAllocationMaxGroupTieredPackagePriceMaxGroupTieredPackageConfigTier struct {
	TierLowerBound param.Field[string] `json:"tier_lower_bound" api:"required"`
	// Per unit amount
	UnitAmount param.Field[string] `json:"unit_amount" api:"required"`
}

func (r PlanNewParamsPricesLicenseAllocationPriceNewLicenseAllocationMaxGroupTieredPackagePriceMaxGroupTieredPackageConfigTier) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// The pricing model type
type PlanNewParamsPricesLicenseAllocationPriceNewLicenseAllocationMaxGroupTieredPackagePriceModelType string

const (
	PlanNewParamsPricesLicenseAllocationPriceNewLicenseAllocationMaxGroupTieredPackagePriceModelTypeMaxGroupTieredPackage PlanNewParamsPricesLicenseAllocationPriceNewLicenseAllocationMaxGroupTieredPackagePriceModelType = "max_group_tiered_package"
)

func (r PlanNewParamsPricesLicenseAllocationPriceNewLicenseAllocationMaxGroupTieredPackagePriceModelType) IsKnown() bool {
	switch r {
	case PlanNewParamsPricesLicenseAllocationPriceNewLicenseAllocationMaxGroupTieredPackagePriceModelTypeMaxGroupTieredPackage:
		return true
	}
	return false
}

type PlanNewParamsPricesLicenseAllocationPriceNewLicenseAllocationMaxGroupTieredPackagePriceConversionRateConfig struct {
	ConversionRateType param.Field[PlanNewParamsPricesLicenseAllocationPriceNewLicenseAllocationMaxGroupTieredPackagePriceConversionRateConfigConversionRateType] `json:"conversion_rate_type" api:"required"`
	TieredConfig       param.Field[shared.ConversionRateTieredConfigParam]                                                                                        `json:"tiered_config"`
	UnitConfig         param.Field[shared.ConversionRateUnitConfigParam]                                                                                          `json:"unit_config"`
}

func (r PlanNewParamsPricesLicenseAllocationPriceNewLicenseAllocationMaxGroupTieredPackagePriceConversionRateConfig) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r PlanNewParamsPricesLicenseAllocationPriceNewLicenseAllocationMaxGroupTieredPackagePriceConversionRateConfig) ImplementsPlanNewParamsPricesLicenseAllocationPriceNewLicenseAllocationMaxGroupTieredPackagePriceConversionRateConfigUnion() {
}

// Satisfied by [shared.UnitConversionRateConfigParam],
// [shared.TieredConversionRateConfigParam],
// [PlanNewParamsPricesLicenseAllocationPriceNewLicenseAllocationMaxGroupTieredPackagePriceConversionRateConfig].
type PlanNewParamsPricesLicenseAllocationPriceNewLicenseAllocationMaxGroupTieredPackagePriceConversionRateConfigUnion interface {
	ImplementsPlanNewParamsPricesLicenseAllocationPriceNewLicenseAllocationMaxGroupTieredPackagePriceConversionRateConfigUnion()
}

type PlanNewParamsPricesLicenseAllocationPriceNewLicenseAllocationMaxGroupTieredPackagePriceConversionRateConfigConversionRateType string

const (
	PlanNewParamsPricesLicenseAllocationPriceNewLicenseAllocationMaxGroupTieredPackagePriceConversionRateConfigConversionRateTypeUnit   PlanNewParamsPricesLicenseAllocationPriceNewLicenseAllocationMaxGroupTieredPackagePriceConversionRateConfigConversionRateType = "unit"
	PlanNewParamsPricesLicenseAllocationPriceNewLicenseAllocationMaxGroupTieredPackagePriceConversionRateConfigConversionRateTypeTiered PlanNewParamsPricesLicenseAllocationPriceNewLicenseAllocationMaxGroupTieredPackagePriceConversionRateConfigConversionRateType = "tiered"
)

func (r PlanNewParamsPricesLicenseAllocationPriceNewLicenseAllocationMaxGroupTieredPackagePriceConversionRateConfigConversionRateType) IsKnown() bool {
	switch r {
	case PlanNewParamsPricesLicenseAllocationPriceNewLicenseAllocationMaxGroupTieredPackagePriceConversionRateConfigConversionRateTypeUnit, PlanNewParamsPricesLicenseAllocationPriceNewLicenseAllocationMaxGroupTieredPackagePriceConversionRateConfigConversionRateTypeTiered:
		return true
	}
	return false
}

type PlanNewParamsPricesLicenseAllocationPriceNewLicenseAllocationScalableMatrixWithUnitPricingPrice struct {
	// The cadence to bill for this price on.
	Cadence param.Field[PlanNewParamsPricesLicenseAllocationPriceNewLicenseAllocationScalableMatrixWithUnitPricingPriceCadence] `json:"cadence" api:"required"`
	// The id of the item the price will be associated with.
	ItemID param.Field[string] `json:"item_id" api:"required"`
	// License allocations to associate with this price. Each entry defines a
	// per-license credit pool granted each cadence. Requires license_type_id or
	// license_type_configuration to be set.
	LicenseAllocations param.Field[[]PlanNewParamsPricesLicenseAllocationPriceNewLicenseAllocationScalableMatrixWithUnitPricingPriceLicenseAllocation] `json:"license_allocations" api:"required"`
	// The pricing model type
	ModelType param.Field[PlanNewParamsPricesLicenseAllocationPriceNewLicenseAllocationScalableMatrixWithUnitPricingPriceModelType] `json:"model_type" api:"required"`
	// The name of the price.
	Name param.Field[string] `json:"name" api:"required"`
	// Configuration for scalable_matrix_with_unit_pricing pricing
	ScalableMatrixWithUnitPricingConfig param.Field[PlanNewParamsPricesLicenseAllocationPriceNewLicenseAllocationScalableMatrixWithUnitPricingPriceScalableMatrixWithUnitPricingConfig] `json:"scalable_matrix_with_unit_pricing_config" api:"required"`
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
	ConversionRateConfig param.Field[PlanNewParamsPricesLicenseAllocationPriceNewLicenseAllocationScalableMatrixWithUnitPricingPriceConversionRateConfigUnion] `json:"conversion_rate_config"`
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
	// The ID of the license type to associate with this price.
	LicenseTypeID param.Field[string] `json:"license_type_id"`
	// User-specified key/value pairs for the resource. Individual keys can be removed
	// by setting the value to `null`, and the entire metadata mapping can be cleared
	// by setting `metadata` to `null`.
	Metadata param.Field[map[string]string] `json:"metadata"`
	// A transient ID that can be used to reference this price when adding adjustments
	// in the same API call.
	ReferenceID param.Field[string] `json:"reference_id"`
}

func (r PlanNewParamsPricesLicenseAllocationPriceNewLicenseAllocationScalableMatrixWithUnitPricingPrice) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r PlanNewParamsPricesLicenseAllocationPriceNewLicenseAllocationScalableMatrixWithUnitPricingPrice) implementsPlanNewParamsPricesLicenseAllocationPriceUnion() {
}

// The cadence to bill for this price on.
type PlanNewParamsPricesLicenseAllocationPriceNewLicenseAllocationScalableMatrixWithUnitPricingPriceCadence string

const (
	PlanNewParamsPricesLicenseAllocationPriceNewLicenseAllocationScalableMatrixWithUnitPricingPriceCadenceAnnual     PlanNewParamsPricesLicenseAllocationPriceNewLicenseAllocationScalableMatrixWithUnitPricingPriceCadence = "annual"
	PlanNewParamsPricesLicenseAllocationPriceNewLicenseAllocationScalableMatrixWithUnitPricingPriceCadenceSemiAnnual PlanNewParamsPricesLicenseAllocationPriceNewLicenseAllocationScalableMatrixWithUnitPricingPriceCadence = "semi_annual"
	PlanNewParamsPricesLicenseAllocationPriceNewLicenseAllocationScalableMatrixWithUnitPricingPriceCadenceMonthly    PlanNewParamsPricesLicenseAllocationPriceNewLicenseAllocationScalableMatrixWithUnitPricingPriceCadence = "monthly"
	PlanNewParamsPricesLicenseAllocationPriceNewLicenseAllocationScalableMatrixWithUnitPricingPriceCadenceQuarterly  PlanNewParamsPricesLicenseAllocationPriceNewLicenseAllocationScalableMatrixWithUnitPricingPriceCadence = "quarterly"
	PlanNewParamsPricesLicenseAllocationPriceNewLicenseAllocationScalableMatrixWithUnitPricingPriceCadenceOneTime    PlanNewParamsPricesLicenseAllocationPriceNewLicenseAllocationScalableMatrixWithUnitPricingPriceCadence = "one_time"
	PlanNewParamsPricesLicenseAllocationPriceNewLicenseAllocationScalableMatrixWithUnitPricingPriceCadenceCustom     PlanNewParamsPricesLicenseAllocationPriceNewLicenseAllocationScalableMatrixWithUnitPricingPriceCadence = "custom"
)

func (r PlanNewParamsPricesLicenseAllocationPriceNewLicenseAllocationScalableMatrixWithUnitPricingPriceCadence) IsKnown() bool {
	switch r {
	case PlanNewParamsPricesLicenseAllocationPriceNewLicenseAllocationScalableMatrixWithUnitPricingPriceCadenceAnnual, PlanNewParamsPricesLicenseAllocationPriceNewLicenseAllocationScalableMatrixWithUnitPricingPriceCadenceSemiAnnual, PlanNewParamsPricesLicenseAllocationPriceNewLicenseAllocationScalableMatrixWithUnitPricingPriceCadenceMonthly, PlanNewParamsPricesLicenseAllocationPriceNewLicenseAllocationScalableMatrixWithUnitPricingPriceCadenceQuarterly, PlanNewParamsPricesLicenseAllocationPriceNewLicenseAllocationScalableMatrixWithUnitPricingPriceCadenceOneTime, PlanNewParamsPricesLicenseAllocationPriceNewLicenseAllocationScalableMatrixWithUnitPricingPriceCadenceCustom:
		return true
	}
	return false
}

type PlanNewParamsPricesLicenseAllocationPriceNewLicenseAllocationScalableMatrixWithUnitPricingPriceLicenseAllocation struct {
	// The amount of credits granted per active license per cadence.
	Amount param.Field[string] `json:"amount" api:"required"`
	// The currency of the license allocation.
	Currency param.Field[string] `json:"currency" api:"required"`
	// When True, overage beyond the allocation is written off.
	WriteOffOverage param.Field[bool] `json:"write_off_overage"`
}

func (r PlanNewParamsPricesLicenseAllocationPriceNewLicenseAllocationScalableMatrixWithUnitPricingPriceLicenseAllocation) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// The pricing model type
type PlanNewParamsPricesLicenseAllocationPriceNewLicenseAllocationScalableMatrixWithUnitPricingPriceModelType string

const (
	PlanNewParamsPricesLicenseAllocationPriceNewLicenseAllocationScalableMatrixWithUnitPricingPriceModelTypeScalableMatrixWithUnitPricing PlanNewParamsPricesLicenseAllocationPriceNewLicenseAllocationScalableMatrixWithUnitPricingPriceModelType = "scalable_matrix_with_unit_pricing"
)

func (r PlanNewParamsPricesLicenseAllocationPriceNewLicenseAllocationScalableMatrixWithUnitPricingPriceModelType) IsKnown() bool {
	switch r {
	case PlanNewParamsPricesLicenseAllocationPriceNewLicenseAllocationScalableMatrixWithUnitPricingPriceModelTypeScalableMatrixWithUnitPricing:
		return true
	}
	return false
}

// Configuration for scalable_matrix_with_unit_pricing pricing
type PlanNewParamsPricesLicenseAllocationPriceNewLicenseAllocationScalableMatrixWithUnitPricingPriceScalableMatrixWithUnitPricingConfig struct {
	// Used to determine the unit rate
	FirstDimension param.Field[string] `json:"first_dimension" api:"required"`
	// Apply a scaling factor to each dimension
	MatrixScalingFactors param.Field[[]PlanNewParamsPricesLicenseAllocationPriceNewLicenseAllocationScalableMatrixWithUnitPricingPriceScalableMatrixWithUnitPricingConfigMatrixScalingFactor] `json:"matrix_scaling_factors" api:"required"`
	// The final unit price to rate against the output of the matrix
	UnitPrice param.Field[string] `json:"unit_price" api:"required"`
	// The property used to group this price
	GroupingKey param.Field[string] `json:"grouping_key"`
	// If true, the unit price will be prorated to the billing period
	Prorate param.Field[bool] `json:"prorate"`
	// Used to determine the unit rate (optional)
	SecondDimension param.Field[string] `json:"second_dimension"`
}

func (r PlanNewParamsPricesLicenseAllocationPriceNewLicenseAllocationScalableMatrixWithUnitPricingPriceScalableMatrixWithUnitPricingConfig) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Configuration for a single matrix scaling factor
type PlanNewParamsPricesLicenseAllocationPriceNewLicenseAllocationScalableMatrixWithUnitPricingPriceScalableMatrixWithUnitPricingConfigMatrixScalingFactor struct {
	FirstDimensionValue  param.Field[string] `json:"first_dimension_value" api:"required"`
	ScalingFactor        param.Field[string] `json:"scaling_factor" api:"required"`
	SecondDimensionValue param.Field[string] `json:"second_dimension_value"`
}

func (r PlanNewParamsPricesLicenseAllocationPriceNewLicenseAllocationScalableMatrixWithUnitPricingPriceScalableMatrixWithUnitPricingConfigMatrixScalingFactor) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type PlanNewParamsPricesLicenseAllocationPriceNewLicenseAllocationScalableMatrixWithUnitPricingPriceConversionRateConfig struct {
	ConversionRateType param.Field[PlanNewParamsPricesLicenseAllocationPriceNewLicenseAllocationScalableMatrixWithUnitPricingPriceConversionRateConfigConversionRateType] `json:"conversion_rate_type" api:"required"`
	TieredConfig       param.Field[shared.ConversionRateTieredConfigParam]                                                                                                `json:"tiered_config"`
	UnitConfig         param.Field[shared.ConversionRateUnitConfigParam]                                                                                                  `json:"unit_config"`
}

func (r PlanNewParamsPricesLicenseAllocationPriceNewLicenseAllocationScalableMatrixWithUnitPricingPriceConversionRateConfig) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r PlanNewParamsPricesLicenseAllocationPriceNewLicenseAllocationScalableMatrixWithUnitPricingPriceConversionRateConfig) ImplementsPlanNewParamsPricesLicenseAllocationPriceNewLicenseAllocationScalableMatrixWithUnitPricingPriceConversionRateConfigUnion() {
}

// Satisfied by [shared.UnitConversionRateConfigParam],
// [shared.TieredConversionRateConfigParam],
// [PlanNewParamsPricesLicenseAllocationPriceNewLicenseAllocationScalableMatrixWithUnitPricingPriceConversionRateConfig].
type PlanNewParamsPricesLicenseAllocationPriceNewLicenseAllocationScalableMatrixWithUnitPricingPriceConversionRateConfigUnion interface {
	ImplementsPlanNewParamsPricesLicenseAllocationPriceNewLicenseAllocationScalableMatrixWithUnitPricingPriceConversionRateConfigUnion()
}

type PlanNewParamsPricesLicenseAllocationPriceNewLicenseAllocationScalableMatrixWithUnitPricingPriceConversionRateConfigConversionRateType string

const (
	PlanNewParamsPricesLicenseAllocationPriceNewLicenseAllocationScalableMatrixWithUnitPricingPriceConversionRateConfigConversionRateTypeUnit   PlanNewParamsPricesLicenseAllocationPriceNewLicenseAllocationScalableMatrixWithUnitPricingPriceConversionRateConfigConversionRateType = "unit"
	PlanNewParamsPricesLicenseAllocationPriceNewLicenseAllocationScalableMatrixWithUnitPricingPriceConversionRateConfigConversionRateTypeTiered PlanNewParamsPricesLicenseAllocationPriceNewLicenseAllocationScalableMatrixWithUnitPricingPriceConversionRateConfigConversionRateType = "tiered"
)

func (r PlanNewParamsPricesLicenseAllocationPriceNewLicenseAllocationScalableMatrixWithUnitPricingPriceConversionRateConfigConversionRateType) IsKnown() bool {
	switch r {
	case PlanNewParamsPricesLicenseAllocationPriceNewLicenseAllocationScalableMatrixWithUnitPricingPriceConversionRateConfigConversionRateTypeUnit, PlanNewParamsPricesLicenseAllocationPriceNewLicenseAllocationScalableMatrixWithUnitPricingPriceConversionRateConfigConversionRateTypeTiered:
		return true
	}
	return false
}

type PlanNewParamsPricesLicenseAllocationPriceNewLicenseAllocationScalableMatrixWithTieredPricingPrice struct {
	// The cadence to bill for this price on.
	Cadence param.Field[PlanNewParamsPricesLicenseAllocationPriceNewLicenseAllocationScalableMatrixWithTieredPricingPriceCadence] `json:"cadence" api:"required"`
	// The id of the item the price will be associated with.
	ItemID param.Field[string] `json:"item_id" api:"required"`
	// License allocations to associate with this price. Each entry defines a
	// per-license credit pool granted each cadence. Requires license_type_id or
	// license_type_configuration to be set.
	LicenseAllocations param.Field[[]PlanNewParamsPricesLicenseAllocationPriceNewLicenseAllocationScalableMatrixWithTieredPricingPriceLicenseAllocation] `json:"license_allocations" api:"required"`
	// The pricing model type
	ModelType param.Field[PlanNewParamsPricesLicenseAllocationPriceNewLicenseAllocationScalableMatrixWithTieredPricingPriceModelType] `json:"model_type" api:"required"`
	// The name of the price.
	Name param.Field[string] `json:"name" api:"required"`
	// Configuration for scalable_matrix_with_tiered_pricing pricing
	ScalableMatrixWithTieredPricingConfig param.Field[PlanNewParamsPricesLicenseAllocationPriceNewLicenseAllocationScalableMatrixWithTieredPricingPriceScalableMatrixWithTieredPricingConfig] `json:"scalable_matrix_with_tiered_pricing_config" api:"required"`
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
	ConversionRateConfig param.Field[PlanNewParamsPricesLicenseAllocationPriceNewLicenseAllocationScalableMatrixWithTieredPricingPriceConversionRateConfigUnion] `json:"conversion_rate_config"`
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
	// The ID of the license type to associate with this price.
	LicenseTypeID param.Field[string] `json:"license_type_id"`
	// User-specified key/value pairs for the resource. Individual keys can be removed
	// by setting the value to `null`, and the entire metadata mapping can be cleared
	// by setting `metadata` to `null`.
	Metadata param.Field[map[string]string] `json:"metadata"`
	// A transient ID that can be used to reference this price when adding adjustments
	// in the same API call.
	ReferenceID param.Field[string] `json:"reference_id"`
}

func (r PlanNewParamsPricesLicenseAllocationPriceNewLicenseAllocationScalableMatrixWithTieredPricingPrice) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r PlanNewParamsPricesLicenseAllocationPriceNewLicenseAllocationScalableMatrixWithTieredPricingPrice) implementsPlanNewParamsPricesLicenseAllocationPriceUnion() {
}

// The cadence to bill for this price on.
type PlanNewParamsPricesLicenseAllocationPriceNewLicenseAllocationScalableMatrixWithTieredPricingPriceCadence string

const (
	PlanNewParamsPricesLicenseAllocationPriceNewLicenseAllocationScalableMatrixWithTieredPricingPriceCadenceAnnual     PlanNewParamsPricesLicenseAllocationPriceNewLicenseAllocationScalableMatrixWithTieredPricingPriceCadence = "annual"
	PlanNewParamsPricesLicenseAllocationPriceNewLicenseAllocationScalableMatrixWithTieredPricingPriceCadenceSemiAnnual PlanNewParamsPricesLicenseAllocationPriceNewLicenseAllocationScalableMatrixWithTieredPricingPriceCadence = "semi_annual"
	PlanNewParamsPricesLicenseAllocationPriceNewLicenseAllocationScalableMatrixWithTieredPricingPriceCadenceMonthly    PlanNewParamsPricesLicenseAllocationPriceNewLicenseAllocationScalableMatrixWithTieredPricingPriceCadence = "monthly"
	PlanNewParamsPricesLicenseAllocationPriceNewLicenseAllocationScalableMatrixWithTieredPricingPriceCadenceQuarterly  PlanNewParamsPricesLicenseAllocationPriceNewLicenseAllocationScalableMatrixWithTieredPricingPriceCadence = "quarterly"
	PlanNewParamsPricesLicenseAllocationPriceNewLicenseAllocationScalableMatrixWithTieredPricingPriceCadenceOneTime    PlanNewParamsPricesLicenseAllocationPriceNewLicenseAllocationScalableMatrixWithTieredPricingPriceCadence = "one_time"
	PlanNewParamsPricesLicenseAllocationPriceNewLicenseAllocationScalableMatrixWithTieredPricingPriceCadenceCustom     PlanNewParamsPricesLicenseAllocationPriceNewLicenseAllocationScalableMatrixWithTieredPricingPriceCadence = "custom"
)

func (r PlanNewParamsPricesLicenseAllocationPriceNewLicenseAllocationScalableMatrixWithTieredPricingPriceCadence) IsKnown() bool {
	switch r {
	case PlanNewParamsPricesLicenseAllocationPriceNewLicenseAllocationScalableMatrixWithTieredPricingPriceCadenceAnnual, PlanNewParamsPricesLicenseAllocationPriceNewLicenseAllocationScalableMatrixWithTieredPricingPriceCadenceSemiAnnual, PlanNewParamsPricesLicenseAllocationPriceNewLicenseAllocationScalableMatrixWithTieredPricingPriceCadenceMonthly, PlanNewParamsPricesLicenseAllocationPriceNewLicenseAllocationScalableMatrixWithTieredPricingPriceCadenceQuarterly, PlanNewParamsPricesLicenseAllocationPriceNewLicenseAllocationScalableMatrixWithTieredPricingPriceCadenceOneTime, PlanNewParamsPricesLicenseAllocationPriceNewLicenseAllocationScalableMatrixWithTieredPricingPriceCadenceCustom:
		return true
	}
	return false
}

type PlanNewParamsPricesLicenseAllocationPriceNewLicenseAllocationScalableMatrixWithTieredPricingPriceLicenseAllocation struct {
	// The amount of credits granted per active license per cadence.
	Amount param.Field[string] `json:"amount" api:"required"`
	// The currency of the license allocation.
	Currency param.Field[string] `json:"currency" api:"required"`
	// When True, overage beyond the allocation is written off.
	WriteOffOverage param.Field[bool] `json:"write_off_overage"`
}

func (r PlanNewParamsPricesLicenseAllocationPriceNewLicenseAllocationScalableMatrixWithTieredPricingPriceLicenseAllocation) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// The pricing model type
type PlanNewParamsPricesLicenseAllocationPriceNewLicenseAllocationScalableMatrixWithTieredPricingPriceModelType string

const (
	PlanNewParamsPricesLicenseAllocationPriceNewLicenseAllocationScalableMatrixWithTieredPricingPriceModelTypeScalableMatrixWithTieredPricing PlanNewParamsPricesLicenseAllocationPriceNewLicenseAllocationScalableMatrixWithTieredPricingPriceModelType = "scalable_matrix_with_tiered_pricing"
)

func (r PlanNewParamsPricesLicenseAllocationPriceNewLicenseAllocationScalableMatrixWithTieredPricingPriceModelType) IsKnown() bool {
	switch r {
	case PlanNewParamsPricesLicenseAllocationPriceNewLicenseAllocationScalableMatrixWithTieredPricingPriceModelTypeScalableMatrixWithTieredPricing:
		return true
	}
	return false
}

// Configuration for scalable_matrix_with_tiered_pricing pricing
type PlanNewParamsPricesLicenseAllocationPriceNewLicenseAllocationScalableMatrixWithTieredPricingPriceScalableMatrixWithTieredPricingConfig struct {
	// Used for the scalable matrix first dimension
	FirstDimension param.Field[string] `json:"first_dimension" api:"required"`
	// Apply a scaling factor to each dimension
	MatrixScalingFactors param.Field[[]PlanNewParamsPricesLicenseAllocationPriceNewLicenseAllocationScalableMatrixWithTieredPricingPriceScalableMatrixWithTieredPricingConfigMatrixScalingFactor] `json:"matrix_scaling_factors" api:"required"`
	Tiers                param.Field[[]PlanNewParamsPricesLicenseAllocationPriceNewLicenseAllocationScalableMatrixWithTieredPricingPriceScalableMatrixWithTieredPricingConfigTier]                `json:"tiers" api:"required"`
	// Used for the scalable matrix second dimension (optional)
	SecondDimension param.Field[string] `json:"second_dimension"`
}

func (r PlanNewParamsPricesLicenseAllocationPriceNewLicenseAllocationScalableMatrixWithTieredPricingPriceScalableMatrixWithTieredPricingConfig) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Configuration for a single matrix scaling factor
type PlanNewParamsPricesLicenseAllocationPriceNewLicenseAllocationScalableMatrixWithTieredPricingPriceScalableMatrixWithTieredPricingConfigMatrixScalingFactor struct {
	FirstDimensionValue  param.Field[string] `json:"first_dimension_value" api:"required"`
	ScalingFactor        param.Field[string] `json:"scaling_factor" api:"required"`
	SecondDimensionValue param.Field[string] `json:"second_dimension_value"`
}

func (r PlanNewParamsPricesLicenseAllocationPriceNewLicenseAllocationScalableMatrixWithTieredPricingPriceScalableMatrixWithTieredPricingConfigMatrixScalingFactor) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Configuration for a single tier entry with business logic
type PlanNewParamsPricesLicenseAllocationPriceNewLicenseAllocationScalableMatrixWithTieredPricingPriceScalableMatrixWithTieredPricingConfigTier struct {
	TierLowerBound param.Field[string] `json:"tier_lower_bound" api:"required"`
	UnitAmount     param.Field[string] `json:"unit_amount" api:"required"`
}

func (r PlanNewParamsPricesLicenseAllocationPriceNewLicenseAllocationScalableMatrixWithTieredPricingPriceScalableMatrixWithTieredPricingConfigTier) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type PlanNewParamsPricesLicenseAllocationPriceNewLicenseAllocationScalableMatrixWithTieredPricingPriceConversionRateConfig struct {
	ConversionRateType param.Field[PlanNewParamsPricesLicenseAllocationPriceNewLicenseAllocationScalableMatrixWithTieredPricingPriceConversionRateConfigConversionRateType] `json:"conversion_rate_type" api:"required"`
	TieredConfig       param.Field[shared.ConversionRateTieredConfigParam]                                                                                                  `json:"tiered_config"`
	UnitConfig         param.Field[shared.ConversionRateUnitConfigParam]                                                                                                    `json:"unit_config"`
}

func (r PlanNewParamsPricesLicenseAllocationPriceNewLicenseAllocationScalableMatrixWithTieredPricingPriceConversionRateConfig) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r PlanNewParamsPricesLicenseAllocationPriceNewLicenseAllocationScalableMatrixWithTieredPricingPriceConversionRateConfig) ImplementsPlanNewParamsPricesLicenseAllocationPriceNewLicenseAllocationScalableMatrixWithTieredPricingPriceConversionRateConfigUnion() {
}

// Satisfied by [shared.UnitConversionRateConfigParam],
// [shared.TieredConversionRateConfigParam],
// [PlanNewParamsPricesLicenseAllocationPriceNewLicenseAllocationScalableMatrixWithTieredPricingPriceConversionRateConfig].
type PlanNewParamsPricesLicenseAllocationPriceNewLicenseAllocationScalableMatrixWithTieredPricingPriceConversionRateConfigUnion interface {
	ImplementsPlanNewParamsPricesLicenseAllocationPriceNewLicenseAllocationScalableMatrixWithTieredPricingPriceConversionRateConfigUnion()
}

type PlanNewParamsPricesLicenseAllocationPriceNewLicenseAllocationScalableMatrixWithTieredPricingPriceConversionRateConfigConversionRateType string

const (
	PlanNewParamsPricesLicenseAllocationPriceNewLicenseAllocationScalableMatrixWithTieredPricingPriceConversionRateConfigConversionRateTypeUnit   PlanNewParamsPricesLicenseAllocationPriceNewLicenseAllocationScalableMatrixWithTieredPricingPriceConversionRateConfigConversionRateType = "unit"
	PlanNewParamsPricesLicenseAllocationPriceNewLicenseAllocationScalableMatrixWithTieredPricingPriceConversionRateConfigConversionRateTypeTiered PlanNewParamsPricesLicenseAllocationPriceNewLicenseAllocationScalableMatrixWithTieredPricingPriceConversionRateConfigConversionRateType = "tiered"
)

func (r PlanNewParamsPricesLicenseAllocationPriceNewLicenseAllocationScalableMatrixWithTieredPricingPriceConversionRateConfigConversionRateType) IsKnown() bool {
	switch r {
	case PlanNewParamsPricesLicenseAllocationPriceNewLicenseAllocationScalableMatrixWithTieredPricingPriceConversionRateConfigConversionRateTypeUnit, PlanNewParamsPricesLicenseAllocationPriceNewLicenseAllocationScalableMatrixWithTieredPricingPriceConversionRateConfigConversionRateTypeTiered:
		return true
	}
	return false
}

type PlanNewParamsPricesLicenseAllocationPriceNewLicenseAllocationCumulativeGroupedBulkPrice struct {
	// The cadence to bill for this price on.
	Cadence param.Field[PlanNewParamsPricesLicenseAllocationPriceNewLicenseAllocationCumulativeGroupedBulkPriceCadence] `json:"cadence" api:"required"`
	// Configuration for cumulative_grouped_bulk pricing
	CumulativeGroupedBulkConfig param.Field[PlanNewParamsPricesLicenseAllocationPriceNewLicenseAllocationCumulativeGroupedBulkPriceCumulativeGroupedBulkConfig] `json:"cumulative_grouped_bulk_config" api:"required"`
	// The id of the item the price will be associated with.
	ItemID param.Field[string] `json:"item_id" api:"required"`
	// License allocations to associate with this price. Each entry defines a
	// per-license credit pool granted each cadence. Requires license_type_id or
	// license_type_configuration to be set.
	LicenseAllocations param.Field[[]PlanNewParamsPricesLicenseAllocationPriceNewLicenseAllocationCumulativeGroupedBulkPriceLicenseAllocation] `json:"license_allocations" api:"required"`
	// The pricing model type
	ModelType param.Field[PlanNewParamsPricesLicenseAllocationPriceNewLicenseAllocationCumulativeGroupedBulkPriceModelType] `json:"model_type" api:"required"`
	// The name of the price.
	Name param.Field[string] `json:"name" api:"required"`
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
	ConversionRateConfig param.Field[PlanNewParamsPricesLicenseAllocationPriceNewLicenseAllocationCumulativeGroupedBulkPriceConversionRateConfigUnion] `json:"conversion_rate_config"`
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
	// The ID of the license type to associate with this price.
	LicenseTypeID param.Field[string] `json:"license_type_id"`
	// User-specified key/value pairs for the resource. Individual keys can be removed
	// by setting the value to `null`, and the entire metadata mapping can be cleared
	// by setting `metadata` to `null`.
	Metadata param.Field[map[string]string] `json:"metadata"`
	// A transient ID that can be used to reference this price when adding adjustments
	// in the same API call.
	ReferenceID param.Field[string] `json:"reference_id"`
}

func (r PlanNewParamsPricesLicenseAllocationPriceNewLicenseAllocationCumulativeGroupedBulkPrice) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r PlanNewParamsPricesLicenseAllocationPriceNewLicenseAllocationCumulativeGroupedBulkPrice) implementsPlanNewParamsPricesLicenseAllocationPriceUnion() {
}

// The cadence to bill for this price on.
type PlanNewParamsPricesLicenseAllocationPriceNewLicenseAllocationCumulativeGroupedBulkPriceCadence string

const (
	PlanNewParamsPricesLicenseAllocationPriceNewLicenseAllocationCumulativeGroupedBulkPriceCadenceAnnual     PlanNewParamsPricesLicenseAllocationPriceNewLicenseAllocationCumulativeGroupedBulkPriceCadence = "annual"
	PlanNewParamsPricesLicenseAllocationPriceNewLicenseAllocationCumulativeGroupedBulkPriceCadenceSemiAnnual PlanNewParamsPricesLicenseAllocationPriceNewLicenseAllocationCumulativeGroupedBulkPriceCadence = "semi_annual"
	PlanNewParamsPricesLicenseAllocationPriceNewLicenseAllocationCumulativeGroupedBulkPriceCadenceMonthly    PlanNewParamsPricesLicenseAllocationPriceNewLicenseAllocationCumulativeGroupedBulkPriceCadence = "monthly"
	PlanNewParamsPricesLicenseAllocationPriceNewLicenseAllocationCumulativeGroupedBulkPriceCadenceQuarterly  PlanNewParamsPricesLicenseAllocationPriceNewLicenseAllocationCumulativeGroupedBulkPriceCadence = "quarterly"
	PlanNewParamsPricesLicenseAllocationPriceNewLicenseAllocationCumulativeGroupedBulkPriceCadenceOneTime    PlanNewParamsPricesLicenseAllocationPriceNewLicenseAllocationCumulativeGroupedBulkPriceCadence = "one_time"
	PlanNewParamsPricesLicenseAllocationPriceNewLicenseAllocationCumulativeGroupedBulkPriceCadenceCustom     PlanNewParamsPricesLicenseAllocationPriceNewLicenseAllocationCumulativeGroupedBulkPriceCadence = "custom"
)

func (r PlanNewParamsPricesLicenseAllocationPriceNewLicenseAllocationCumulativeGroupedBulkPriceCadence) IsKnown() bool {
	switch r {
	case PlanNewParamsPricesLicenseAllocationPriceNewLicenseAllocationCumulativeGroupedBulkPriceCadenceAnnual, PlanNewParamsPricesLicenseAllocationPriceNewLicenseAllocationCumulativeGroupedBulkPriceCadenceSemiAnnual, PlanNewParamsPricesLicenseAllocationPriceNewLicenseAllocationCumulativeGroupedBulkPriceCadenceMonthly, PlanNewParamsPricesLicenseAllocationPriceNewLicenseAllocationCumulativeGroupedBulkPriceCadenceQuarterly, PlanNewParamsPricesLicenseAllocationPriceNewLicenseAllocationCumulativeGroupedBulkPriceCadenceOneTime, PlanNewParamsPricesLicenseAllocationPriceNewLicenseAllocationCumulativeGroupedBulkPriceCadenceCustom:
		return true
	}
	return false
}

// Configuration for cumulative_grouped_bulk pricing
type PlanNewParamsPricesLicenseAllocationPriceNewLicenseAllocationCumulativeGroupedBulkPriceCumulativeGroupedBulkConfig struct {
	// Each tier lower bound must have the same group of values.
	DimensionValues param.Field[[]PlanNewParamsPricesLicenseAllocationPriceNewLicenseAllocationCumulativeGroupedBulkPriceCumulativeGroupedBulkConfigDimensionValue] `json:"dimension_values" api:"required"`
	Group           param.Field[string]                                                                                                                             `json:"group" api:"required"`
}

func (r PlanNewParamsPricesLicenseAllocationPriceNewLicenseAllocationCumulativeGroupedBulkPriceCumulativeGroupedBulkConfig) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Configuration for a dimension value entry
type PlanNewParamsPricesLicenseAllocationPriceNewLicenseAllocationCumulativeGroupedBulkPriceCumulativeGroupedBulkConfigDimensionValue struct {
	// Grouping key value
	GroupingKey param.Field[string] `json:"grouping_key" api:"required"`
	// Tier lower bound
	TierLowerBound param.Field[string] `json:"tier_lower_bound" api:"required"`
	// Unit amount for this combination
	UnitAmount param.Field[string] `json:"unit_amount" api:"required"`
}

func (r PlanNewParamsPricesLicenseAllocationPriceNewLicenseAllocationCumulativeGroupedBulkPriceCumulativeGroupedBulkConfigDimensionValue) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type PlanNewParamsPricesLicenseAllocationPriceNewLicenseAllocationCumulativeGroupedBulkPriceLicenseAllocation struct {
	// The amount of credits granted per active license per cadence.
	Amount param.Field[string] `json:"amount" api:"required"`
	// The currency of the license allocation.
	Currency param.Field[string] `json:"currency" api:"required"`
	// When True, overage beyond the allocation is written off.
	WriteOffOverage param.Field[bool] `json:"write_off_overage"`
}

func (r PlanNewParamsPricesLicenseAllocationPriceNewLicenseAllocationCumulativeGroupedBulkPriceLicenseAllocation) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// The pricing model type
type PlanNewParamsPricesLicenseAllocationPriceNewLicenseAllocationCumulativeGroupedBulkPriceModelType string

const (
	PlanNewParamsPricesLicenseAllocationPriceNewLicenseAllocationCumulativeGroupedBulkPriceModelTypeCumulativeGroupedBulk PlanNewParamsPricesLicenseAllocationPriceNewLicenseAllocationCumulativeGroupedBulkPriceModelType = "cumulative_grouped_bulk"
)

func (r PlanNewParamsPricesLicenseAllocationPriceNewLicenseAllocationCumulativeGroupedBulkPriceModelType) IsKnown() bool {
	switch r {
	case PlanNewParamsPricesLicenseAllocationPriceNewLicenseAllocationCumulativeGroupedBulkPriceModelTypeCumulativeGroupedBulk:
		return true
	}
	return false
}

type PlanNewParamsPricesLicenseAllocationPriceNewLicenseAllocationCumulativeGroupedBulkPriceConversionRateConfig struct {
	ConversionRateType param.Field[PlanNewParamsPricesLicenseAllocationPriceNewLicenseAllocationCumulativeGroupedBulkPriceConversionRateConfigConversionRateType] `json:"conversion_rate_type" api:"required"`
	TieredConfig       param.Field[shared.ConversionRateTieredConfigParam]                                                                                        `json:"tiered_config"`
	UnitConfig         param.Field[shared.ConversionRateUnitConfigParam]                                                                                          `json:"unit_config"`
}

func (r PlanNewParamsPricesLicenseAllocationPriceNewLicenseAllocationCumulativeGroupedBulkPriceConversionRateConfig) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r PlanNewParamsPricesLicenseAllocationPriceNewLicenseAllocationCumulativeGroupedBulkPriceConversionRateConfig) ImplementsPlanNewParamsPricesLicenseAllocationPriceNewLicenseAllocationCumulativeGroupedBulkPriceConversionRateConfigUnion() {
}

// Satisfied by [shared.UnitConversionRateConfigParam],
// [shared.TieredConversionRateConfigParam],
// [PlanNewParamsPricesLicenseAllocationPriceNewLicenseAllocationCumulativeGroupedBulkPriceConversionRateConfig].
type PlanNewParamsPricesLicenseAllocationPriceNewLicenseAllocationCumulativeGroupedBulkPriceConversionRateConfigUnion interface {
	ImplementsPlanNewParamsPricesLicenseAllocationPriceNewLicenseAllocationCumulativeGroupedBulkPriceConversionRateConfigUnion()
}

type PlanNewParamsPricesLicenseAllocationPriceNewLicenseAllocationCumulativeGroupedBulkPriceConversionRateConfigConversionRateType string

const (
	PlanNewParamsPricesLicenseAllocationPriceNewLicenseAllocationCumulativeGroupedBulkPriceConversionRateConfigConversionRateTypeUnit   PlanNewParamsPricesLicenseAllocationPriceNewLicenseAllocationCumulativeGroupedBulkPriceConversionRateConfigConversionRateType = "unit"
	PlanNewParamsPricesLicenseAllocationPriceNewLicenseAllocationCumulativeGroupedBulkPriceConversionRateConfigConversionRateTypeTiered PlanNewParamsPricesLicenseAllocationPriceNewLicenseAllocationCumulativeGroupedBulkPriceConversionRateConfigConversionRateType = "tiered"
)

func (r PlanNewParamsPricesLicenseAllocationPriceNewLicenseAllocationCumulativeGroupedBulkPriceConversionRateConfigConversionRateType) IsKnown() bool {
	switch r {
	case PlanNewParamsPricesLicenseAllocationPriceNewLicenseAllocationCumulativeGroupedBulkPriceConversionRateConfigConversionRateTypeUnit, PlanNewParamsPricesLicenseAllocationPriceNewLicenseAllocationCumulativeGroupedBulkPriceConversionRateConfigConversionRateTypeTiered:
		return true
	}
	return false
}

type PlanNewParamsPricesLicenseAllocationPriceNewLicenseAllocationCumulativeGroupedAllocationPrice struct {
	// The cadence to bill for this price on.
	Cadence param.Field[PlanNewParamsPricesLicenseAllocationPriceNewLicenseAllocationCumulativeGroupedAllocationPriceCadence] `json:"cadence" api:"required"`
	// Configuration for cumulative_grouped_allocation pricing
	CumulativeGroupedAllocationConfig param.Field[PlanNewParamsPricesLicenseAllocationPriceNewLicenseAllocationCumulativeGroupedAllocationPriceCumulativeGroupedAllocationConfig] `json:"cumulative_grouped_allocation_config" api:"required"`
	// The id of the item the price will be associated with.
	ItemID param.Field[string] `json:"item_id" api:"required"`
	// License allocations to associate with this price. Each entry defines a
	// per-license credit pool granted each cadence. Requires license_type_id or
	// license_type_configuration to be set.
	LicenseAllocations param.Field[[]PlanNewParamsPricesLicenseAllocationPriceNewLicenseAllocationCumulativeGroupedAllocationPriceLicenseAllocation] `json:"license_allocations" api:"required"`
	// The pricing model type
	ModelType param.Field[PlanNewParamsPricesLicenseAllocationPriceNewLicenseAllocationCumulativeGroupedAllocationPriceModelType] `json:"model_type" api:"required"`
	// The name of the price.
	Name param.Field[string] `json:"name" api:"required"`
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
	ConversionRateConfig param.Field[PlanNewParamsPricesLicenseAllocationPriceNewLicenseAllocationCumulativeGroupedAllocationPriceConversionRateConfigUnion] `json:"conversion_rate_config"`
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
	// The ID of the license type to associate with this price.
	LicenseTypeID param.Field[string] `json:"license_type_id"`
	// User-specified key/value pairs for the resource. Individual keys can be removed
	// by setting the value to `null`, and the entire metadata mapping can be cleared
	// by setting `metadata` to `null`.
	Metadata param.Field[map[string]string] `json:"metadata"`
	// A transient ID that can be used to reference this price when adding adjustments
	// in the same API call.
	ReferenceID param.Field[string] `json:"reference_id"`
}

func (r PlanNewParamsPricesLicenseAllocationPriceNewLicenseAllocationCumulativeGroupedAllocationPrice) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r PlanNewParamsPricesLicenseAllocationPriceNewLicenseAllocationCumulativeGroupedAllocationPrice) implementsPlanNewParamsPricesLicenseAllocationPriceUnion() {
}

// The cadence to bill for this price on.
type PlanNewParamsPricesLicenseAllocationPriceNewLicenseAllocationCumulativeGroupedAllocationPriceCadence string

const (
	PlanNewParamsPricesLicenseAllocationPriceNewLicenseAllocationCumulativeGroupedAllocationPriceCadenceAnnual     PlanNewParamsPricesLicenseAllocationPriceNewLicenseAllocationCumulativeGroupedAllocationPriceCadence = "annual"
	PlanNewParamsPricesLicenseAllocationPriceNewLicenseAllocationCumulativeGroupedAllocationPriceCadenceSemiAnnual PlanNewParamsPricesLicenseAllocationPriceNewLicenseAllocationCumulativeGroupedAllocationPriceCadence = "semi_annual"
	PlanNewParamsPricesLicenseAllocationPriceNewLicenseAllocationCumulativeGroupedAllocationPriceCadenceMonthly    PlanNewParamsPricesLicenseAllocationPriceNewLicenseAllocationCumulativeGroupedAllocationPriceCadence = "monthly"
	PlanNewParamsPricesLicenseAllocationPriceNewLicenseAllocationCumulativeGroupedAllocationPriceCadenceQuarterly  PlanNewParamsPricesLicenseAllocationPriceNewLicenseAllocationCumulativeGroupedAllocationPriceCadence = "quarterly"
	PlanNewParamsPricesLicenseAllocationPriceNewLicenseAllocationCumulativeGroupedAllocationPriceCadenceOneTime    PlanNewParamsPricesLicenseAllocationPriceNewLicenseAllocationCumulativeGroupedAllocationPriceCadence = "one_time"
	PlanNewParamsPricesLicenseAllocationPriceNewLicenseAllocationCumulativeGroupedAllocationPriceCadenceCustom     PlanNewParamsPricesLicenseAllocationPriceNewLicenseAllocationCumulativeGroupedAllocationPriceCadence = "custom"
)

func (r PlanNewParamsPricesLicenseAllocationPriceNewLicenseAllocationCumulativeGroupedAllocationPriceCadence) IsKnown() bool {
	switch r {
	case PlanNewParamsPricesLicenseAllocationPriceNewLicenseAllocationCumulativeGroupedAllocationPriceCadenceAnnual, PlanNewParamsPricesLicenseAllocationPriceNewLicenseAllocationCumulativeGroupedAllocationPriceCadenceSemiAnnual, PlanNewParamsPricesLicenseAllocationPriceNewLicenseAllocationCumulativeGroupedAllocationPriceCadenceMonthly, PlanNewParamsPricesLicenseAllocationPriceNewLicenseAllocationCumulativeGroupedAllocationPriceCadenceQuarterly, PlanNewParamsPricesLicenseAllocationPriceNewLicenseAllocationCumulativeGroupedAllocationPriceCadenceOneTime, PlanNewParamsPricesLicenseAllocationPriceNewLicenseAllocationCumulativeGroupedAllocationPriceCadenceCustom:
		return true
	}
	return false
}

// Configuration for cumulative_grouped_allocation pricing
type PlanNewParamsPricesLicenseAllocationPriceNewLicenseAllocationCumulativeGroupedAllocationPriceCumulativeGroupedAllocationConfig struct {
	// The overall allocation across all groups
	CumulativeAllocation param.Field[string] `json:"cumulative_allocation" api:"required"`
	// The allocation per individual group
	GroupAllocation param.Field[string] `json:"group_allocation" api:"required"`
	// The event property used to group usage before applying allocations
	GroupingKey param.Field[string] `json:"grouping_key" api:"required"`
	// The amount to charge for each unit outside of the allocation
	UnitAmount param.Field[string] `json:"unit_amount" api:"required"`
}

func (r PlanNewParamsPricesLicenseAllocationPriceNewLicenseAllocationCumulativeGroupedAllocationPriceCumulativeGroupedAllocationConfig) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type PlanNewParamsPricesLicenseAllocationPriceNewLicenseAllocationCumulativeGroupedAllocationPriceLicenseAllocation struct {
	// The amount of credits granted per active license per cadence.
	Amount param.Field[string] `json:"amount" api:"required"`
	// The currency of the license allocation.
	Currency param.Field[string] `json:"currency" api:"required"`
	// When True, overage beyond the allocation is written off.
	WriteOffOverage param.Field[bool] `json:"write_off_overage"`
}

func (r PlanNewParamsPricesLicenseAllocationPriceNewLicenseAllocationCumulativeGroupedAllocationPriceLicenseAllocation) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// The pricing model type
type PlanNewParamsPricesLicenseAllocationPriceNewLicenseAllocationCumulativeGroupedAllocationPriceModelType string

const (
	PlanNewParamsPricesLicenseAllocationPriceNewLicenseAllocationCumulativeGroupedAllocationPriceModelTypeCumulativeGroupedAllocation PlanNewParamsPricesLicenseAllocationPriceNewLicenseAllocationCumulativeGroupedAllocationPriceModelType = "cumulative_grouped_allocation"
)

func (r PlanNewParamsPricesLicenseAllocationPriceNewLicenseAllocationCumulativeGroupedAllocationPriceModelType) IsKnown() bool {
	switch r {
	case PlanNewParamsPricesLicenseAllocationPriceNewLicenseAllocationCumulativeGroupedAllocationPriceModelTypeCumulativeGroupedAllocation:
		return true
	}
	return false
}

type PlanNewParamsPricesLicenseAllocationPriceNewLicenseAllocationCumulativeGroupedAllocationPriceConversionRateConfig struct {
	ConversionRateType param.Field[PlanNewParamsPricesLicenseAllocationPriceNewLicenseAllocationCumulativeGroupedAllocationPriceConversionRateConfigConversionRateType] `json:"conversion_rate_type" api:"required"`
	TieredConfig       param.Field[shared.ConversionRateTieredConfigParam]                                                                                              `json:"tiered_config"`
	UnitConfig         param.Field[shared.ConversionRateUnitConfigParam]                                                                                                `json:"unit_config"`
}

func (r PlanNewParamsPricesLicenseAllocationPriceNewLicenseAllocationCumulativeGroupedAllocationPriceConversionRateConfig) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r PlanNewParamsPricesLicenseAllocationPriceNewLicenseAllocationCumulativeGroupedAllocationPriceConversionRateConfig) ImplementsPlanNewParamsPricesLicenseAllocationPriceNewLicenseAllocationCumulativeGroupedAllocationPriceConversionRateConfigUnion() {
}

// Satisfied by [shared.UnitConversionRateConfigParam],
// [shared.TieredConversionRateConfigParam],
// [PlanNewParamsPricesLicenseAllocationPriceNewLicenseAllocationCumulativeGroupedAllocationPriceConversionRateConfig].
type PlanNewParamsPricesLicenseAllocationPriceNewLicenseAllocationCumulativeGroupedAllocationPriceConversionRateConfigUnion interface {
	ImplementsPlanNewParamsPricesLicenseAllocationPriceNewLicenseAllocationCumulativeGroupedAllocationPriceConversionRateConfigUnion()
}

type PlanNewParamsPricesLicenseAllocationPriceNewLicenseAllocationCumulativeGroupedAllocationPriceConversionRateConfigConversionRateType string

const (
	PlanNewParamsPricesLicenseAllocationPriceNewLicenseAllocationCumulativeGroupedAllocationPriceConversionRateConfigConversionRateTypeUnit   PlanNewParamsPricesLicenseAllocationPriceNewLicenseAllocationCumulativeGroupedAllocationPriceConversionRateConfigConversionRateType = "unit"
	PlanNewParamsPricesLicenseAllocationPriceNewLicenseAllocationCumulativeGroupedAllocationPriceConversionRateConfigConversionRateTypeTiered PlanNewParamsPricesLicenseAllocationPriceNewLicenseAllocationCumulativeGroupedAllocationPriceConversionRateConfigConversionRateType = "tiered"
)

func (r PlanNewParamsPricesLicenseAllocationPriceNewLicenseAllocationCumulativeGroupedAllocationPriceConversionRateConfigConversionRateType) IsKnown() bool {
	switch r {
	case PlanNewParamsPricesLicenseAllocationPriceNewLicenseAllocationCumulativeGroupedAllocationPriceConversionRateConfigConversionRateTypeUnit, PlanNewParamsPricesLicenseAllocationPriceNewLicenseAllocationCumulativeGroupedAllocationPriceConversionRateConfigConversionRateTypeTiered:
		return true
	}
	return false
}

type PlanNewParamsPricesLicenseAllocationPriceNewLicenseAllocationMinimumCompositePrice struct {
	// The cadence to bill for this price on.
	Cadence param.Field[PlanNewParamsPricesLicenseAllocationPriceNewLicenseAllocationMinimumCompositePriceCadence] `json:"cadence" api:"required"`
	// The id of the item the price will be associated with.
	ItemID param.Field[string] `json:"item_id" api:"required"`
	// License allocations to associate with this price. Each entry defines a
	// per-license credit pool granted each cadence. Requires license_type_id or
	// license_type_configuration to be set.
	LicenseAllocations param.Field[[]PlanNewParamsPricesLicenseAllocationPriceNewLicenseAllocationMinimumCompositePriceLicenseAllocation] `json:"license_allocations" api:"required"`
	// Configuration for minimum_composite pricing
	MinimumCompositeConfig param.Field[PlanNewParamsPricesLicenseAllocationPriceNewLicenseAllocationMinimumCompositePriceMinimumCompositeConfig] `json:"minimum_composite_config" api:"required"`
	// The pricing model type
	ModelType param.Field[PlanNewParamsPricesLicenseAllocationPriceNewLicenseAllocationMinimumCompositePriceModelType] `json:"model_type" api:"required"`
	// The name of the price.
	Name param.Field[string] `json:"name" api:"required"`
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
	ConversionRateConfig param.Field[PlanNewParamsPricesLicenseAllocationPriceNewLicenseAllocationMinimumCompositePriceConversionRateConfigUnion] `json:"conversion_rate_config"`
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
	// The ID of the license type to associate with this price.
	LicenseTypeID param.Field[string] `json:"license_type_id"`
	// User-specified key/value pairs for the resource. Individual keys can be removed
	// by setting the value to `null`, and the entire metadata mapping can be cleared
	// by setting `metadata` to `null`.
	Metadata param.Field[map[string]string] `json:"metadata"`
	// A transient ID that can be used to reference this price when adding adjustments
	// in the same API call.
	ReferenceID param.Field[string] `json:"reference_id"`
}

func (r PlanNewParamsPricesLicenseAllocationPriceNewLicenseAllocationMinimumCompositePrice) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r PlanNewParamsPricesLicenseAllocationPriceNewLicenseAllocationMinimumCompositePrice) implementsPlanNewParamsPricesLicenseAllocationPriceUnion() {
}

// The cadence to bill for this price on.
type PlanNewParamsPricesLicenseAllocationPriceNewLicenseAllocationMinimumCompositePriceCadence string

const (
	PlanNewParamsPricesLicenseAllocationPriceNewLicenseAllocationMinimumCompositePriceCadenceAnnual     PlanNewParamsPricesLicenseAllocationPriceNewLicenseAllocationMinimumCompositePriceCadence = "annual"
	PlanNewParamsPricesLicenseAllocationPriceNewLicenseAllocationMinimumCompositePriceCadenceSemiAnnual PlanNewParamsPricesLicenseAllocationPriceNewLicenseAllocationMinimumCompositePriceCadence = "semi_annual"
	PlanNewParamsPricesLicenseAllocationPriceNewLicenseAllocationMinimumCompositePriceCadenceMonthly    PlanNewParamsPricesLicenseAllocationPriceNewLicenseAllocationMinimumCompositePriceCadence = "monthly"
	PlanNewParamsPricesLicenseAllocationPriceNewLicenseAllocationMinimumCompositePriceCadenceQuarterly  PlanNewParamsPricesLicenseAllocationPriceNewLicenseAllocationMinimumCompositePriceCadence = "quarterly"
	PlanNewParamsPricesLicenseAllocationPriceNewLicenseAllocationMinimumCompositePriceCadenceOneTime    PlanNewParamsPricesLicenseAllocationPriceNewLicenseAllocationMinimumCompositePriceCadence = "one_time"
	PlanNewParamsPricesLicenseAllocationPriceNewLicenseAllocationMinimumCompositePriceCadenceCustom     PlanNewParamsPricesLicenseAllocationPriceNewLicenseAllocationMinimumCompositePriceCadence = "custom"
)

func (r PlanNewParamsPricesLicenseAllocationPriceNewLicenseAllocationMinimumCompositePriceCadence) IsKnown() bool {
	switch r {
	case PlanNewParamsPricesLicenseAllocationPriceNewLicenseAllocationMinimumCompositePriceCadenceAnnual, PlanNewParamsPricesLicenseAllocationPriceNewLicenseAllocationMinimumCompositePriceCadenceSemiAnnual, PlanNewParamsPricesLicenseAllocationPriceNewLicenseAllocationMinimumCompositePriceCadenceMonthly, PlanNewParamsPricesLicenseAllocationPriceNewLicenseAllocationMinimumCompositePriceCadenceQuarterly, PlanNewParamsPricesLicenseAllocationPriceNewLicenseAllocationMinimumCompositePriceCadenceOneTime, PlanNewParamsPricesLicenseAllocationPriceNewLicenseAllocationMinimumCompositePriceCadenceCustom:
		return true
	}
	return false
}

type PlanNewParamsPricesLicenseAllocationPriceNewLicenseAllocationMinimumCompositePriceLicenseAllocation struct {
	// The amount of credits granted per active license per cadence.
	Amount param.Field[string] `json:"amount" api:"required"`
	// The currency of the license allocation.
	Currency param.Field[string] `json:"currency" api:"required"`
	// When True, overage beyond the allocation is written off.
	WriteOffOverage param.Field[bool] `json:"write_off_overage"`
}

func (r PlanNewParamsPricesLicenseAllocationPriceNewLicenseAllocationMinimumCompositePriceLicenseAllocation) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Configuration for minimum_composite pricing
type PlanNewParamsPricesLicenseAllocationPriceNewLicenseAllocationMinimumCompositePriceMinimumCompositeConfig struct {
	// The minimum amount to apply
	MinimumAmount param.Field[string] `json:"minimum_amount" api:"required"`
	// If true, subtotals from this price are prorated based on the service period
	Prorated param.Field[bool] `json:"prorated"`
}

func (r PlanNewParamsPricesLicenseAllocationPriceNewLicenseAllocationMinimumCompositePriceMinimumCompositeConfig) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// The pricing model type
type PlanNewParamsPricesLicenseAllocationPriceNewLicenseAllocationMinimumCompositePriceModelType string

const (
	PlanNewParamsPricesLicenseAllocationPriceNewLicenseAllocationMinimumCompositePriceModelTypeMinimumComposite PlanNewParamsPricesLicenseAllocationPriceNewLicenseAllocationMinimumCompositePriceModelType = "minimum_composite"
)

func (r PlanNewParamsPricesLicenseAllocationPriceNewLicenseAllocationMinimumCompositePriceModelType) IsKnown() bool {
	switch r {
	case PlanNewParamsPricesLicenseAllocationPriceNewLicenseAllocationMinimumCompositePriceModelTypeMinimumComposite:
		return true
	}
	return false
}

type PlanNewParamsPricesLicenseAllocationPriceNewLicenseAllocationMinimumCompositePriceConversionRateConfig struct {
	ConversionRateType param.Field[PlanNewParamsPricesLicenseAllocationPriceNewLicenseAllocationMinimumCompositePriceConversionRateConfigConversionRateType] `json:"conversion_rate_type" api:"required"`
	TieredConfig       param.Field[shared.ConversionRateTieredConfigParam]                                                                                   `json:"tiered_config"`
	UnitConfig         param.Field[shared.ConversionRateUnitConfigParam]                                                                                     `json:"unit_config"`
}

func (r PlanNewParamsPricesLicenseAllocationPriceNewLicenseAllocationMinimumCompositePriceConversionRateConfig) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r PlanNewParamsPricesLicenseAllocationPriceNewLicenseAllocationMinimumCompositePriceConversionRateConfig) ImplementsPlanNewParamsPricesLicenseAllocationPriceNewLicenseAllocationMinimumCompositePriceConversionRateConfigUnion() {
}

// Satisfied by [shared.UnitConversionRateConfigParam],
// [shared.TieredConversionRateConfigParam],
// [PlanNewParamsPricesLicenseAllocationPriceNewLicenseAllocationMinimumCompositePriceConversionRateConfig].
type PlanNewParamsPricesLicenseAllocationPriceNewLicenseAllocationMinimumCompositePriceConversionRateConfigUnion interface {
	ImplementsPlanNewParamsPricesLicenseAllocationPriceNewLicenseAllocationMinimumCompositePriceConversionRateConfigUnion()
}

type PlanNewParamsPricesLicenseAllocationPriceNewLicenseAllocationMinimumCompositePriceConversionRateConfigConversionRateType string

const (
	PlanNewParamsPricesLicenseAllocationPriceNewLicenseAllocationMinimumCompositePriceConversionRateConfigConversionRateTypeUnit   PlanNewParamsPricesLicenseAllocationPriceNewLicenseAllocationMinimumCompositePriceConversionRateConfigConversionRateType = "unit"
	PlanNewParamsPricesLicenseAllocationPriceNewLicenseAllocationMinimumCompositePriceConversionRateConfigConversionRateTypeTiered PlanNewParamsPricesLicenseAllocationPriceNewLicenseAllocationMinimumCompositePriceConversionRateConfigConversionRateType = "tiered"
)

func (r PlanNewParamsPricesLicenseAllocationPriceNewLicenseAllocationMinimumCompositePriceConversionRateConfigConversionRateType) IsKnown() bool {
	switch r {
	case PlanNewParamsPricesLicenseAllocationPriceNewLicenseAllocationMinimumCompositePriceConversionRateConfigConversionRateTypeUnit, PlanNewParamsPricesLicenseAllocationPriceNewLicenseAllocationMinimumCompositePriceConversionRateConfigConversionRateTypeTiered:
		return true
	}
	return false
}

type PlanNewParamsPricesLicenseAllocationPriceNewLicenseAllocationPercentCompositePrice struct {
	// The cadence to bill for this price on.
	Cadence param.Field[PlanNewParamsPricesLicenseAllocationPriceNewLicenseAllocationPercentCompositePriceCadence] `json:"cadence" api:"required"`
	// The id of the item the price will be associated with.
	ItemID param.Field[string] `json:"item_id" api:"required"`
	// License allocations to associate with this price. Each entry defines a
	// per-license credit pool granted each cadence. Requires license_type_id or
	// license_type_configuration to be set.
	LicenseAllocations param.Field[[]PlanNewParamsPricesLicenseAllocationPriceNewLicenseAllocationPercentCompositePriceLicenseAllocation] `json:"license_allocations" api:"required"`
	// The pricing model type
	ModelType param.Field[PlanNewParamsPricesLicenseAllocationPriceNewLicenseAllocationPercentCompositePriceModelType] `json:"model_type" api:"required"`
	// The name of the price.
	Name param.Field[string] `json:"name" api:"required"`
	// Configuration for percent pricing
	PercentConfig param.Field[PlanNewParamsPricesLicenseAllocationPriceNewLicenseAllocationPercentCompositePricePercentConfig] `json:"percent_config" api:"required"`
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
	ConversionRateConfig param.Field[PlanNewParamsPricesLicenseAllocationPriceNewLicenseAllocationPercentCompositePriceConversionRateConfigUnion] `json:"conversion_rate_config"`
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
	// The ID of the license type to associate with this price.
	LicenseTypeID param.Field[string] `json:"license_type_id"`
	// User-specified key/value pairs for the resource. Individual keys can be removed
	// by setting the value to `null`, and the entire metadata mapping can be cleared
	// by setting `metadata` to `null`.
	Metadata param.Field[map[string]string] `json:"metadata"`
	// A transient ID that can be used to reference this price when adding adjustments
	// in the same API call.
	ReferenceID param.Field[string] `json:"reference_id"`
}

func (r PlanNewParamsPricesLicenseAllocationPriceNewLicenseAllocationPercentCompositePrice) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r PlanNewParamsPricesLicenseAllocationPriceNewLicenseAllocationPercentCompositePrice) implementsPlanNewParamsPricesLicenseAllocationPriceUnion() {
}

// The cadence to bill for this price on.
type PlanNewParamsPricesLicenseAllocationPriceNewLicenseAllocationPercentCompositePriceCadence string

const (
	PlanNewParamsPricesLicenseAllocationPriceNewLicenseAllocationPercentCompositePriceCadenceAnnual     PlanNewParamsPricesLicenseAllocationPriceNewLicenseAllocationPercentCompositePriceCadence = "annual"
	PlanNewParamsPricesLicenseAllocationPriceNewLicenseAllocationPercentCompositePriceCadenceSemiAnnual PlanNewParamsPricesLicenseAllocationPriceNewLicenseAllocationPercentCompositePriceCadence = "semi_annual"
	PlanNewParamsPricesLicenseAllocationPriceNewLicenseAllocationPercentCompositePriceCadenceMonthly    PlanNewParamsPricesLicenseAllocationPriceNewLicenseAllocationPercentCompositePriceCadence = "monthly"
	PlanNewParamsPricesLicenseAllocationPriceNewLicenseAllocationPercentCompositePriceCadenceQuarterly  PlanNewParamsPricesLicenseAllocationPriceNewLicenseAllocationPercentCompositePriceCadence = "quarterly"
	PlanNewParamsPricesLicenseAllocationPriceNewLicenseAllocationPercentCompositePriceCadenceOneTime    PlanNewParamsPricesLicenseAllocationPriceNewLicenseAllocationPercentCompositePriceCadence = "one_time"
	PlanNewParamsPricesLicenseAllocationPriceNewLicenseAllocationPercentCompositePriceCadenceCustom     PlanNewParamsPricesLicenseAllocationPriceNewLicenseAllocationPercentCompositePriceCadence = "custom"
)

func (r PlanNewParamsPricesLicenseAllocationPriceNewLicenseAllocationPercentCompositePriceCadence) IsKnown() bool {
	switch r {
	case PlanNewParamsPricesLicenseAllocationPriceNewLicenseAllocationPercentCompositePriceCadenceAnnual, PlanNewParamsPricesLicenseAllocationPriceNewLicenseAllocationPercentCompositePriceCadenceSemiAnnual, PlanNewParamsPricesLicenseAllocationPriceNewLicenseAllocationPercentCompositePriceCadenceMonthly, PlanNewParamsPricesLicenseAllocationPriceNewLicenseAllocationPercentCompositePriceCadenceQuarterly, PlanNewParamsPricesLicenseAllocationPriceNewLicenseAllocationPercentCompositePriceCadenceOneTime, PlanNewParamsPricesLicenseAllocationPriceNewLicenseAllocationPercentCompositePriceCadenceCustom:
		return true
	}
	return false
}

type PlanNewParamsPricesLicenseAllocationPriceNewLicenseAllocationPercentCompositePriceLicenseAllocation struct {
	// The amount of credits granted per active license per cadence.
	Amount param.Field[string] `json:"amount" api:"required"`
	// The currency of the license allocation.
	Currency param.Field[string] `json:"currency" api:"required"`
	// When True, overage beyond the allocation is written off.
	WriteOffOverage param.Field[bool] `json:"write_off_overage"`
}

func (r PlanNewParamsPricesLicenseAllocationPriceNewLicenseAllocationPercentCompositePriceLicenseAllocation) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// The pricing model type
type PlanNewParamsPricesLicenseAllocationPriceNewLicenseAllocationPercentCompositePriceModelType string

const (
	PlanNewParamsPricesLicenseAllocationPriceNewLicenseAllocationPercentCompositePriceModelTypePercent PlanNewParamsPricesLicenseAllocationPriceNewLicenseAllocationPercentCompositePriceModelType = "percent"
)

func (r PlanNewParamsPricesLicenseAllocationPriceNewLicenseAllocationPercentCompositePriceModelType) IsKnown() bool {
	switch r {
	case PlanNewParamsPricesLicenseAllocationPriceNewLicenseAllocationPercentCompositePriceModelTypePercent:
		return true
	}
	return false
}

// Configuration for percent pricing
type PlanNewParamsPricesLicenseAllocationPriceNewLicenseAllocationPercentCompositePricePercentConfig struct {
	// What percent of the component subtotals to charge
	Percent param.Field[float64] `json:"percent" api:"required"`
}

func (r PlanNewParamsPricesLicenseAllocationPriceNewLicenseAllocationPercentCompositePricePercentConfig) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type PlanNewParamsPricesLicenseAllocationPriceNewLicenseAllocationPercentCompositePriceConversionRateConfig struct {
	ConversionRateType param.Field[PlanNewParamsPricesLicenseAllocationPriceNewLicenseAllocationPercentCompositePriceConversionRateConfigConversionRateType] `json:"conversion_rate_type" api:"required"`
	TieredConfig       param.Field[shared.ConversionRateTieredConfigParam]                                                                                   `json:"tiered_config"`
	UnitConfig         param.Field[shared.ConversionRateUnitConfigParam]                                                                                     `json:"unit_config"`
}

func (r PlanNewParamsPricesLicenseAllocationPriceNewLicenseAllocationPercentCompositePriceConversionRateConfig) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r PlanNewParamsPricesLicenseAllocationPriceNewLicenseAllocationPercentCompositePriceConversionRateConfig) ImplementsPlanNewParamsPricesLicenseAllocationPriceNewLicenseAllocationPercentCompositePriceConversionRateConfigUnion() {
}

// Satisfied by [shared.UnitConversionRateConfigParam],
// [shared.TieredConversionRateConfigParam],
// [PlanNewParamsPricesLicenseAllocationPriceNewLicenseAllocationPercentCompositePriceConversionRateConfig].
type PlanNewParamsPricesLicenseAllocationPriceNewLicenseAllocationPercentCompositePriceConversionRateConfigUnion interface {
	ImplementsPlanNewParamsPricesLicenseAllocationPriceNewLicenseAllocationPercentCompositePriceConversionRateConfigUnion()
}

type PlanNewParamsPricesLicenseAllocationPriceNewLicenseAllocationPercentCompositePriceConversionRateConfigConversionRateType string

const (
	PlanNewParamsPricesLicenseAllocationPriceNewLicenseAllocationPercentCompositePriceConversionRateConfigConversionRateTypeUnit   PlanNewParamsPricesLicenseAllocationPriceNewLicenseAllocationPercentCompositePriceConversionRateConfigConversionRateType = "unit"
	PlanNewParamsPricesLicenseAllocationPriceNewLicenseAllocationPercentCompositePriceConversionRateConfigConversionRateTypeTiered PlanNewParamsPricesLicenseAllocationPriceNewLicenseAllocationPercentCompositePriceConversionRateConfigConversionRateType = "tiered"
)

func (r PlanNewParamsPricesLicenseAllocationPriceNewLicenseAllocationPercentCompositePriceConversionRateConfigConversionRateType) IsKnown() bool {
	switch r {
	case PlanNewParamsPricesLicenseAllocationPriceNewLicenseAllocationPercentCompositePriceConversionRateConfigConversionRateTypeUnit, PlanNewParamsPricesLicenseAllocationPriceNewLicenseAllocationPercentCompositePriceConversionRateConfigConversionRateTypeTiered:
		return true
	}
	return false
}

type PlanNewParamsPricesLicenseAllocationPriceNewLicenseAllocationEventOutputPrice struct {
	// The cadence to bill for this price on.
	Cadence param.Field[PlanNewParamsPricesLicenseAllocationPriceNewLicenseAllocationEventOutputPriceCadence] `json:"cadence" api:"required"`
	// Configuration for event_output pricing
	EventOutputConfig param.Field[PlanNewParamsPricesLicenseAllocationPriceNewLicenseAllocationEventOutputPriceEventOutputConfig] `json:"event_output_config" api:"required"`
	// The id of the item the price will be associated with.
	ItemID param.Field[string] `json:"item_id" api:"required"`
	// License allocations to associate with this price. Each entry defines a
	// per-license credit pool granted each cadence. Requires license_type_id or
	// license_type_configuration to be set.
	LicenseAllocations param.Field[[]PlanNewParamsPricesLicenseAllocationPriceNewLicenseAllocationEventOutputPriceLicenseAllocation] `json:"license_allocations" api:"required"`
	// The pricing model type
	ModelType param.Field[PlanNewParamsPricesLicenseAllocationPriceNewLicenseAllocationEventOutputPriceModelType] `json:"model_type" api:"required"`
	// The name of the price.
	Name param.Field[string] `json:"name" api:"required"`
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
	ConversionRateConfig param.Field[PlanNewParamsPricesLicenseAllocationPriceNewLicenseAllocationEventOutputPriceConversionRateConfigUnion] `json:"conversion_rate_config"`
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
	// The ID of the license type to associate with this price.
	LicenseTypeID param.Field[string] `json:"license_type_id"`
	// User-specified key/value pairs for the resource. Individual keys can be removed
	// by setting the value to `null`, and the entire metadata mapping can be cleared
	// by setting `metadata` to `null`.
	Metadata param.Field[map[string]string] `json:"metadata"`
	// A transient ID that can be used to reference this price when adding adjustments
	// in the same API call.
	ReferenceID param.Field[string] `json:"reference_id"`
}

func (r PlanNewParamsPricesLicenseAllocationPriceNewLicenseAllocationEventOutputPrice) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r PlanNewParamsPricesLicenseAllocationPriceNewLicenseAllocationEventOutputPrice) implementsPlanNewParamsPricesLicenseAllocationPriceUnion() {
}

// The cadence to bill for this price on.
type PlanNewParamsPricesLicenseAllocationPriceNewLicenseAllocationEventOutputPriceCadence string

const (
	PlanNewParamsPricesLicenseAllocationPriceNewLicenseAllocationEventOutputPriceCadenceAnnual     PlanNewParamsPricesLicenseAllocationPriceNewLicenseAllocationEventOutputPriceCadence = "annual"
	PlanNewParamsPricesLicenseAllocationPriceNewLicenseAllocationEventOutputPriceCadenceSemiAnnual PlanNewParamsPricesLicenseAllocationPriceNewLicenseAllocationEventOutputPriceCadence = "semi_annual"
	PlanNewParamsPricesLicenseAllocationPriceNewLicenseAllocationEventOutputPriceCadenceMonthly    PlanNewParamsPricesLicenseAllocationPriceNewLicenseAllocationEventOutputPriceCadence = "monthly"
	PlanNewParamsPricesLicenseAllocationPriceNewLicenseAllocationEventOutputPriceCadenceQuarterly  PlanNewParamsPricesLicenseAllocationPriceNewLicenseAllocationEventOutputPriceCadence = "quarterly"
	PlanNewParamsPricesLicenseAllocationPriceNewLicenseAllocationEventOutputPriceCadenceOneTime    PlanNewParamsPricesLicenseAllocationPriceNewLicenseAllocationEventOutputPriceCadence = "one_time"
	PlanNewParamsPricesLicenseAllocationPriceNewLicenseAllocationEventOutputPriceCadenceCustom     PlanNewParamsPricesLicenseAllocationPriceNewLicenseAllocationEventOutputPriceCadence = "custom"
)

func (r PlanNewParamsPricesLicenseAllocationPriceNewLicenseAllocationEventOutputPriceCadence) IsKnown() bool {
	switch r {
	case PlanNewParamsPricesLicenseAllocationPriceNewLicenseAllocationEventOutputPriceCadenceAnnual, PlanNewParamsPricesLicenseAllocationPriceNewLicenseAllocationEventOutputPriceCadenceSemiAnnual, PlanNewParamsPricesLicenseAllocationPriceNewLicenseAllocationEventOutputPriceCadenceMonthly, PlanNewParamsPricesLicenseAllocationPriceNewLicenseAllocationEventOutputPriceCadenceQuarterly, PlanNewParamsPricesLicenseAllocationPriceNewLicenseAllocationEventOutputPriceCadenceOneTime, PlanNewParamsPricesLicenseAllocationPriceNewLicenseAllocationEventOutputPriceCadenceCustom:
		return true
	}
	return false
}

// Configuration for event_output pricing
type PlanNewParamsPricesLicenseAllocationPriceNewLicenseAllocationEventOutputPriceEventOutputConfig struct {
	// The key in the event data to extract the unit rate from.
	UnitRatingKey param.Field[string] `json:"unit_rating_key" api:"required"`
	// If provided, this amount will be used as the unit rate when an event does not
	// have a value for the `unit_rating_key`. If not provided, events missing a unit
	// rate will be ignored.
	DefaultUnitRate param.Field[string] `json:"default_unit_rate"`
	// An optional key in the event data to group by (e.g., event ID). All events will
	// also be grouped by their unit rate.
	GroupingKey param.Field[string] `json:"grouping_key"`
}

func (r PlanNewParamsPricesLicenseAllocationPriceNewLicenseAllocationEventOutputPriceEventOutputConfig) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type PlanNewParamsPricesLicenseAllocationPriceNewLicenseAllocationEventOutputPriceLicenseAllocation struct {
	// The amount of credits granted per active license per cadence.
	Amount param.Field[string] `json:"amount" api:"required"`
	// The currency of the license allocation.
	Currency param.Field[string] `json:"currency" api:"required"`
	// When True, overage beyond the allocation is written off.
	WriteOffOverage param.Field[bool] `json:"write_off_overage"`
}

func (r PlanNewParamsPricesLicenseAllocationPriceNewLicenseAllocationEventOutputPriceLicenseAllocation) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// The pricing model type
type PlanNewParamsPricesLicenseAllocationPriceNewLicenseAllocationEventOutputPriceModelType string

const (
	PlanNewParamsPricesLicenseAllocationPriceNewLicenseAllocationEventOutputPriceModelTypeEventOutput PlanNewParamsPricesLicenseAllocationPriceNewLicenseAllocationEventOutputPriceModelType = "event_output"
)

func (r PlanNewParamsPricesLicenseAllocationPriceNewLicenseAllocationEventOutputPriceModelType) IsKnown() bool {
	switch r {
	case PlanNewParamsPricesLicenseAllocationPriceNewLicenseAllocationEventOutputPriceModelTypeEventOutput:
		return true
	}
	return false
}

type PlanNewParamsPricesLicenseAllocationPriceNewLicenseAllocationEventOutputPriceConversionRateConfig struct {
	ConversionRateType param.Field[PlanNewParamsPricesLicenseAllocationPriceNewLicenseAllocationEventOutputPriceConversionRateConfigConversionRateType] `json:"conversion_rate_type" api:"required"`
	TieredConfig       param.Field[shared.ConversionRateTieredConfigParam]                                                                              `json:"tiered_config"`
	UnitConfig         param.Field[shared.ConversionRateUnitConfigParam]                                                                                `json:"unit_config"`
}

func (r PlanNewParamsPricesLicenseAllocationPriceNewLicenseAllocationEventOutputPriceConversionRateConfig) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r PlanNewParamsPricesLicenseAllocationPriceNewLicenseAllocationEventOutputPriceConversionRateConfig) ImplementsPlanNewParamsPricesLicenseAllocationPriceNewLicenseAllocationEventOutputPriceConversionRateConfigUnion() {
}

// Satisfied by [shared.UnitConversionRateConfigParam],
// [shared.TieredConversionRateConfigParam],
// [PlanNewParamsPricesLicenseAllocationPriceNewLicenseAllocationEventOutputPriceConversionRateConfig].
type PlanNewParamsPricesLicenseAllocationPriceNewLicenseAllocationEventOutputPriceConversionRateConfigUnion interface {
	ImplementsPlanNewParamsPricesLicenseAllocationPriceNewLicenseAllocationEventOutputPriceConversionRateConfigUnion()
}

type PlanNewParamsPricesLicenseAllocationPriceNewLicenseAllocationEventOutputPriceConversionRateConfigConversionRateType string

const (
	PlanNewParamsPricesLicenseAllocationPriceNewLicenseAllocationEventOutputPriceConversionRateConfigConversionRateTypeUnit   PlanNewParamsPricesLicenseAllocationPriceNewLicenseAllocationEventOutputPriceConversionRateConfigConversionRateType = "unit"
	PlanNewParamsPricesLicenseAllocationPriceNewLicenseAllocationEventOutputPriceConversionRateConfigConversionRateTypeTiered PlanNewParamsPricesLicenseAllocationPriceNewLicenseAllocationEventOutputPriceConversionRateConfigConversionRateType = "tiered"
)

func (r PlanNewParamsPricesLicenseAllocationPriceNewLicenseAllocationEventOutputPriceConversionRateConfigConversionRateType) IsKnown() bool {
	switch r {
	case PlanNewParamsPricesLicenseAllocationPriceNewLicenseAllocationEventOutputPriceConversionRateConfigConversionRateTypeUnit, PlanNewParamsPricesLicenseAllocationPriceNewLicenseAllocationEventOutputPriceConversionRateConfigConversionRateTypeTiered:
		return true
	}
	return false
}

// The cadence to bill for this price on.
type PlanNewParamsPricesLicenseAllocationPriceCadence string

const (
	PlanNewParamsPricesLicenseAllocationPriceCadenceAnnual     PlanNewParamsPricesLicenseAllocationPriceCadence = "annual"
	PlanNewParamsPricesLicenseAllocationPriceCadenceSemiAnnual PlanNewParamsPricesLicenseAllocationPriceCadence = "semi_annual"
	PlanNewParamsPricesLicenseAllocationPriceCadenceMonthly    PlanNewParamsPricesLicenseAllocationPriceCadence = "monthly"
	PlanNewParamsPricesLicenseAllocationPriceCadenceQuarterly  PlanNewParamsPricesLicenseAllocationPriceCadence = "quarterly"
	PlanNewParamsPricesLicenseAllocationPriceCadenceOneTime    PlanNewParamsPricesLicenseAllocationPriceCadence = "one_time"
	PlanNewParamsPricesLicenseAllocationPriceCadenceCustom     PlanNewParamsPricesLicenseAllocationPriceCadence = "custom"
)

func (r PlanNewParamsPricesLicenseAllocationPriceCadence) IsKnown() bool {
	switch r {
	case PlanNewParamsPricesLicenseAllocationPriceCadenceAnnual, PlanNewParamsPricesLicenseAllocationPriceCadenceSemiAnnual, PlanNewParamsPricesLicenseAllocationPriceCadenceMonthly, PlanNewParamsPricesLicenseAllocationPriceCadenceQuarterly, PlanNewParamsPricesLicenseAllocationPriceCadenceOneTime, PlanNewParamsPricesLicenseAllocationPriceCadenceCustom:
		return true
	}
	return false
}

// The pricing model type
type PlanNewParamsPricesLicenseAllocationPriceModelType string

const (
	PlanNewParamsPricesLicenseAllocationPriceModelTypeUnit                            PlanNewParamsPricesLicenseAllocationPriceModelType = "unit"
	PlanNewParamsPricesLicenseAllocationPriceModelTypeTiered                          PlanNewParamsPricesLicenseAllocationPriceModelType = "tiered"
	PlanNewParamsPricesLicenseAllocationPriceModelTypeBulk                            PlanNewParamsPricesLicenseAllocationPriceModelType = "bulk"
	PlanNewParamsPricesLicenseAllocationPriceModelTypeBulkWithFilters                 PlanNewParamsPricesLicenseAllocationPriceModelType = "bulk_with_filters"
	PlanNewParamsPricesLicenseAllocationPriceModelTypePackage                         PlanNewParamsPricesLicenseAllocationPriceModelType = "package"
	PlanNewParamsPricesLicenseAllocationPriceModelTypeMatrix                          PlanNewParamsPricesLicenseAllocationPriceModelType = "matrix"
	PlanNewParamsPricesLicenseAllocationPriceModelTypeThresholdTotalAmount            PlanNewParamsPricesLicenseAllocationPriceModelType = "threshold_total_amount"
	PlanNewParamsPricesLicenseAllocationPriceModelTypeTieredPackage                   PlanNewParamsPricesLicenseAllocationPriceModelType = "tiered_package"
	PlanNewParamsPricesLicenseAllocationPriceModelTypeTieredWithMinimum               PlanNewParamsPricesLicenseAllocationPriceModelType = "tiered_with_minimum"
	PlanNewParamsPricesLicenseAllocationPriceModelTypeGroupedTiered                   PlanNewParamsPricesLicenseAllocationPriceModelType = "grouped_tiered"
	PlanNewParamsPricesLicenseAllocationPriceModelTypeTieredPackageWithMinimum        PlanNewParamsPricesLicenseAllocationPriceModelType = "tiered_package_with_minimum"
	PlanNewParamsPricesLicenseAllocationPriceModelTypePackageWithAllocation           PlanNewParamsPricesLicenseAllocationPriceModelType = "package_with_allocation"
	PlanNewParamsPricesLicenseAllocationPriceModelTypeUnitWithPercent                 PlanNewParamsPricesLicenseAllocationPriceModelType = "unit_with_percent"
	PlanNewParamsPricesLicenseAllocationPriceModelTypeMatrixWithAllocation            PlanNewParamsPricesLicenseAllocationPriceModelType = "matrix_with_allocation"
	PlanNewParamsPricesLicenseAllocationPriceModelTypeTieredWithProration             PlanNewParamsPricesLicenseAllocationPriceModelType = "tiered_with_proration"
	PlanNewParamsPricesLicenseAllocationPriceModelTypeUnitWithProration               PlanNewParamsPricesLicenseAllocationPriceModelType = "unit_with_proration"
	PlanNewParamsPricesLicenseAllocationPriceModelTypeGroupedAllocation               PlanNewParamsPricesLicenseAllocationPriceModelType = "grouped_allocation"
	PlanNewParamsPricesLicenseAllocationPriceModelTypeBulkWithProration               PlanNewParamsPricesLicenseAllocationPriceModelType = "bulk_with_proration"
	PlanNewParamsPricesLicenseAllocationPriceModelTypeGroupedWithProratedMinimum      PlanNewParamsPricesLicenseAllocationPriceModelType = "grouped_with_prorated_minimum"
	PlanNewParamsPricesLicenseAllocationPriceModelTypeGroupedWithMeteredMinimum       PlanNewParamsPricesLicenseAllocationPriceModelType = "grouped_with_metered_minimum"
	PlanNewParamsPricesLicenseAllocationPriceModelTypeGroupedWithMinMaxThresholds     PlanNewParamsPricesLicenseAllocationPriceModelType = "grouped_with_min_max_thresholds"
	PlanNewParamsPricesLicenseAllocationPriceModelTypeMatrixWithDisplayName           PlanNewParamsPricesLicenseAllocationPriceModelType = "matrix_with_display_name"
	PlanNewParamsPricesLicenseAllocationPriceModelTypeGroupedTieredPackage            PlanNewParamsPricesLicenseAllocationPriceModelType = "grouped_tiered_package"
	PlanNewParamsPricesLicenseAllocationPriceModelTypeMaxGroupTieredPackage           PlanNewParamsPricesLicenseAllocationPriceModelType = "max_group_tiered_package"
	PlanNewParamsPricesLicenseAllocationPriceModelTypeScalableMatrixWithUnitPricing   PlanNewParamsPricesLicenseAllocationPriceModelType = "scalable_matrix_with_unit_pricing"
	PlanNewParamsPricesLicenseAllocationPriceModelTypeScalableMatrixWithTieredPricing PlanNewParamsPricesLicenseAllocationPriceModelType = "scalable_matrix_with_tiered_pricing"
	PlanNewParamsPricesLicenseAllocationPriceModelTypeCumulativeGroupedBulk           PlanNewParamsPricesLicenseAllocationPriceModelType = "cumulative_grouped_bulk"
	PlanNewParamsPricesLicenseAllocationPriceModelTypeCumulativeGroupedAllocation     PlanNewParamsPricesLicenseAllocationPriceModelType = "cumulative_grouped_allocation"
	PlanNewParamsPricesLicenseAllocationPriceModelTypeMinimumComposite                PlanNewParamsPricesLicenseAllocationPriceModelType = "minimum_composite"
	PlanNewParamsPricesLicenseAllocationPriceModelTypePercent                         PlanNewParamsPricesLicenseAllocationPriceModelType = "percent"
	PlanNewParamsPricesLicenseAllocationPriceModelTypeEventOutput                     PlanNewParamsPricesLicenseAllocationPriceModelType = "event_output"
)

func (r PlanNewParamsPricesLicenseAllocationPriceModelType) IsKnown() bool {
	switch r {
	case PlanNewParamsPricesLicenseAllocationPriceModelTypeUnit, PlanNewParamsPricesLicenseAllocationPriceModelTypeTiered, PlanNewParamsPricesLicenseAllocationPriceModelTypeBulk, PlanNewParamsPricesLicenseAllocationPriceModelTypeBulkWithFilters, PlanNewParamsPricesLicenseAllocationPriceModelTypePackage, PlanNewParamsPricesLicenseAllocationPriceModelTypeMatrix, PlanNewParamsPricesLicenseAllocationPriceModelTypeThresholdTotalAmount, PlanNewParamsPricesLicenseAllocationPriceModelTypeTieredPackage, PlanNewParamsPricesLicenseAllocationPriceModelTypeTieredWithMinimum, PlanNewParamsPricesLicenseAllocationPriceModelTypeGroupedTiered, PlanNewParamsPricesLicenseAllocationPriceModelTypeTieredPackageWithMinimum, PlanNewParamsPricesLicenseAllocationPriceModelTypePackageWithAllocation, PlanNewParamsPricesLicenseAllocationPriceModelTypeUnitWithPercent, PlanNewParamsPricesLicenseAllocationPriceModelTypeMatrixWithAllocation, PlanNewParamsPricesLicenseAllocationPriceModelTypeTieredWithProration, PlanNewParamsPricesLicenseAllocationPriceModelTypeUnitWithProration, PlanNewParamsPricesLicenseAllocationPriceModelTypeGroupedAllocation, PlanNewParamsPricesLicenseAllocationPriceModelTypeBulkWithProration, PlanNewParamsPricesLicenseAllocationPriceModelTypeGroupedWithProratedMinimum, PlanNewParamsPricesLicenseAllocationPriceModelTypeGroupedWithMeteredMinimum, PlanNewParamsPricesLicenseAllocationPriceModelTypeGroupedWithMinMaxThresholds, PlanNewParamsPricesLicenseAllocationPriceModelTypeMatrixWithDisplayName, PlanNewParamsPricesLicenseAllocationPriceModelTypeGroupedTieredPackage, PlanNewParamsPricesLicenseAllocationPriceModelTypeMaxGroupTieredPackage, PlanNewParamsPricesLicenseAllocationPriceModelTypeScalableMatrixWithUnitPricing, PlanNewParamsPricesLicenseAllocationPriceModelTypeScalableMatrixWithTieredPricing, PlanNewParamsPricesLicenseAllocationPriceModelTypeCumulativeGroupedBulk, PlanNewParamsPricesLicenseAllocationPriceModelTypeCumulativeGroupedAllocation, PlanNewParamsPricesLicenseAllocationPriceModelTypeMinimumComposite, PlanNewParamsPricesLicenseAllocationPriceModelTypePercent, PlanNewParamsPricesLicenseAllocationPriceModelTypeEventOutput:
		return true
	}
	return false
}

// New plan price request body params.
type PlanNewParamsPricesPrice struct {
	// The cadence to bill for this price on.
	Cadence param.Field[PlanNewParamsPricesPriceCadence] `json:"cadence" api:"required"`
	// The id of the item the price will be associated with.
	ItemID param.Field[string] `json:"item_id" api:"required"`
	// The pricing model type
	ModelType param.Field[PlanNewParamsPricesPriceModelType] `json:"model_type" api:"required"`
	// The name of the price.
	Name param.Field[string] `json:"name" api:"required"`
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
	// The ID of the license type to associate with this price.
	LicenseTypeID param.Field[string] `json:"license_type_id"`
	// Configuration for matrix pricing
	MatrixConfig param.Field[shared.MatrixConfigParam] `json:"matrix_config"`
	// Configuration for matrix_with_allocation pricing
	MatrixWithAllocationConfig  param.Field[shared.MatrixWithAllocationConfigParam] `json:"matrix_with_allocation_config"`
	MatrixWithDisplayNameConfig param.Field[interface{}]                            `json:"matrix_with_display_name_config"`
	MaxGroupTieredPackageConfig param.Field[interface{}]                            `json:"max_group_tiered_package_config"`
	Metadata                    param.Field[interface{}]                            `json:"metadata"`
	MinimumCompositeConfig      param.Field[interface{}]                            `json:"minimum_composite_config"`
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
	BulkWithFiltersConfig param.Field[PlanNewParamsPricesPriceNewPlanBulkWithFiltersPriceBulkWithFiltersConfig] `json:"bulk_with_filters_config" api:"required"`
	// The cadence to bill for this price on.
	Cadence param.Field[PlanNewParamsPricesPriceNewPlanBulkWithFiltersPriceCadence] `json:"cadence" api:"required"`
	// The id of the item the price will be associated with.
	ItemID param.Field[string] `json:"item_id" api:"required"`
	// The pricing model type
	ModelType param.Field[PlanNewParamsPricesPriceNewPlanBulkWithFiltersPriceModelType] `json:"model_type" api:"required"`
	// The name of the price.
	Name param.Field[string] `json:"name" api:"required"`
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
	// The ID of the license type to associate with this price.
	LicenseTypeID param.Field[string] `json:"license_type_id"`
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
	Filters param.Field[[]PlanNewParamsPricesPriceNewPlanBulkWithFiltersPriceBulkWithFiltersConfigFilter] `json:"filters" api:"required"`
	// Bulk tiers for rating based on total usage volume
	Tiers param.Field[[]PlanNewParamsPricesPriceNewPlanBulkWithFiltersPriceBulkWithFiltersConfigTier] `json:"tiers" api:"required"`
}

func (r PlanNewParamsPricesPriceNewPlanBulkWithFiltersPriceBulkWithFiltersConfig) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Configuration for a single property filter
type PlanNewParamsPricesPriceNewPlanBulkWithFiltersPriceBulkWithFiltersConfigFilter struct {
	// Event property key to filter on
	PropertyKey param.Field[string] `json:"property_key" api:"required"`
	// Event property value to match
	PropertyValue param.Field[string] `json:"property_value" api:"required"`
}

func (r PlanNewParamsPricesPriceNewPlanBulkWithFiltersPriceBulkWithFiltersConfigFilter) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Configuration for a single bulk pricing tier
type PlanNewParamsPricesPriceNewPlanBulkWithFiltersPriceBulkWithFiltersConfigTier struct {
	// Amount per unit
	UnitAmount param.Field[string] `json:"unit_amount" api:"required"`
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
	ConversionRateType param.Field[PlanNewParamsPricesPriceNewPlanBulkWithFiltersPriceConversionRateConfigConversionRateType] `json:"conversion_rate_type" api:"required"`
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
	Cadence param.Field[PlanNewParamsPricesPriceNewPlanTieredWithProrationPriceCadence] `json:"cadence" api:"required"`
	// The id of the item the price will be associated with.
	ItemID param.Field[string] `json:"item_id" api:"required"`
	// The pricing model type
	ModelType param.Field[PlanNewParamsPricesPriceNewPlanTieredWithProrationPriceModelType] `json:"model_type" api:"required"`
	// The name of the price.
	Name param.Field[string] `json:"name" api:"required"`
	// Configuration for tiered_with_proration pricing
	TieredWithProrationConfig param.Field[PlanNewParamsPricesPriceNewPlanTieredWithProrationPriceTieredWithProrationConfig] `json:"tiered_with_proration_config" api:"required"`
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
	// The ID of the license type to associate with this price.
	LicenseTypeID param.Field[string] `json:"license_type_id"`
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
	Tiers param.Field[[]PlanNewParamsPricesPriceNewPlanTieredWithProrationPriceTieredWithProrationConfigTier] `json:"tiers" api:"required"`
}

func (r PlanNewParamsPricesPriceNewPlanTieredWithProrationPriceTieredWithProrationConfig) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Configuration for a single tiered with proration tier
type PlanNewParamsPricesPriceNewPlanTieredWithProrationPriceTieredWithProrationConfigTier struct {
	// Inclusive tier starting value
	TierLowerBound param.Field[string] `json:"tier_lower_bound" api:"required"`
	// Amount per unit
	UnitAmount param.Field[string] `json:"unit_amount" api:"required"`
}

func (r PlanNewParamsPricesPriceNewPlanTieredWithProrationPriceTieredWithProrationConfigTier) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type PlanNewParamsPricesPriceNewPlanTieredWithProrationPriceConversionRateConfig struct {
	ConversionRateType param.Field[PlanNewParamsPricesPriceNewPlanTieredWithProrationPriceConversionRateConfigConversionRateType] `json:"conversion_rate_type" api:"required"`
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
	Cadence param.Field[PlanNewParamsPricesPriceNewPlanGroupedWithMinMaxThresholdsPriceCadence] `json:"cadence" api:"required"`
	// Configuration for grouped_with_min_max_thresholds pricing
	GroupedWithMinMaxThresholdsConfig param.Field[PlanNewParamsPricesPriceNewPlanGroupedWithMinMaxThresholdsPriceGroupedWithMinMaxThresholdsConfig] `json:"grouped_with_min_max_thresholds_config" api:"required"`
	// The id of the item the price will be associated with.
	ItemID param.Field[string] `json:"item_id" api:"required"`
	// The pricing model type
	ModelType param.Field[PlanNewParamsPricesPriceNewPlanGroupedWithMinMaxThresholdsPriceModelType] `json:"model_type" api:"required"`
	// The name of the price.
	Name param.Field[string] `json:"name" api:"required"`
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
	// The ID of the license type to associate with this price.
	LicenseTypeID param.Field[string] `json:"license_type_id"`
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
	GroupingKey param.Field[string] `json:"grouping_key" api:"required"`
	// The maximum amount to charge each group
	MaximumCharge param.Field[string] `json:"maximum_charge" api:"required"`
	// The minimum amount to charge each group, regardless of usage
	MinimumCharge param.Field[string] `json:"minimum_charge" api:"required"`
	// The base price charged per group
	PerUnitRate param.Field[string] `json:"per_unit_rate" api:"required"`
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
	ConversionRateType param.Field[PlanNewParamsPricesPriceNewPlanGroupedWithMinMaxThresholdsPriceConversionRateConfigConversionRateType] `json:"conversion_rate_type" api:"required"`
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
	Cadence param.Field[PlanNewParamsPricesPriceNewPlanCumulativeGroupedAllocationPriceCadence] `json:"cadence" api:"required"`
	// Configuration for cumulative_grouped_allocation pricing
	CumulativeGroupedAllocationConfig param.Field[PlanNewParamsPricesPriceNewPlanCumulativeGroupedAllocationPriceCumulativeGroupedAllocationConfig] `json:"cumulative_grouped_allocation_config" api:"required"`
	// The id of the item the price will be associated with.
	ItemID param.Field[string] `json:"item_id" api:"required"`
	// The pricing model type
	ModelType param.Field[PlanNewParamsPricesPriceNewPlanCumulativeGroupedAllocationPriceModelType] `json:"model_type" api:"required"`
	// The name of the price.
	Name param.Field[string] `json:"name" api:"required"`
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
	// The ID of the license type to associate with this price.
	LicenseTypeID param.Field[string] `json:"license_type_id"`
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
	CumulativeAllocation param.Field[string] `json:"cumulative_allocation" api:"required"`
	// The allocation per individual group
	GroupAllocation param.Field[string] `json:"group_allocation" api:"required"`
	// The event property used to group usage before applying allocations
	GroupingKey param.Field[string] `json:"grouping_key" api:"required"`
	// The amount to charge for each unit outside of the allocation
	UnitAmount param.Field[string] `json:"unit_amount" api:"required"`
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
	ConversionRateType param.Field[PlanNewParamsPricesPriceNewPlanCumulativeGroupedAllocationPriceConversionRateConfigConversionRateType] `json:"conversion_rate_type" api:"required"`
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
	Cadence param.Field[PlanNewParamsPricesPriceNewPlanPercentCompositePriceCadence] `json:"cadence" api:"required"`
	// The id of the item the price will be associated with.
	ItemID param.Field[string] `json:"item_id" api:"required"`
	// The pricing model type
	ModelType param.Field[PlanNewParamsPricesPriceNewPlanPercentCompositePriceModelType] `json:"model_type" api:"required"`
	// The name of the price.
	Name param.Field[string] `json:"name" api:"required"`
	// Configuration for percent pricing
	PercentConfig param.Field[PlanNewParamsPricesPriceNewPlanPercentCompositePricePercentConfig] `json:"percent_config" api:"required"`
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
	// The ID of the license type to associate with this price.
	LicenseTypeID param.Field[string] `json:"license_type_id"`
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
	Percent param.Field[float64] `json:"percent" api:"required"`
}

func (r PlanNewParamsPricesPriceNewPlanPercentCompositePricePercentConfig) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type PlanNewParamsPricesPriceNewPlanPercentCompositePriceConversionRateConfig struct {
	ConversionRateType param.Field[PlanNewParamsPricesPriceNewPlanPercentCompositePriceConversionRateConfigConversionRateType] `json:"conversion_rate_type" api:"required"`
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
	Cadence param.Field[PlanNewParamsPricesPriceNewPlanEventOutputPriceCadence] `json:"cadence" api:"required"`
	// Configuration for event_output pricing
	EventOutputConfig param.Field[PlanNewParamsPricesPriceNewPlanEventOutputPriceEventOutputConfig] `json:"event_output_config" api:"required"`
	// The id of the item the price will be associated with.
	ItemID param.Field[string] `json:"item_id" api:"required"`
	// The pricing model type
	ModelType param.Field[PlanNewParamsPricesPriceNewPlanEventOutputPriceModelType] `json:"model_type" api:"required"`
	// The name of the price.
	Name param.Field[string] `json:"name" api:"required"`
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
	// The ID of the license type to associate with this price.
	LicenseTypeID param.Field[string] `json:"license_type_id"`
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
	UnitRatingKey param.Field[string] `json:"unit_rating_key" api:"required"`
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
	ConversionRateType param.Field[PlanNewParamsPricesPriceNewPlanEventOutputPriceConversionRateConfigConversionRateType] `json:"conversion_rate_type" api:"required"`
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
	PlanNewParamsPricesPriceModelTypeMinimumComposite                PlanNewParamsPricesPriceModelType = "minimum_composite"
	PlanNewParamsPricesPriceModelTypePercent                         PlanNewParamsPricesPriceModelType = "percent"
	PlanNewParamsPricesPriceModelTypeEventOutput                     PlanNewParamsPricesPriceModelType = "event_output"
)

func (r PlanNewParamsPricesPriceModelType) IsKnown() bool {
	switch r {
	case PlanNewParamsPricesPriceModelTypeUnit, PlanNewParamsPricesPriceModelTypeTiered, PlanNewParamsPricesPriceModelTypeBulk, PlanNewParamsPricesPriceModelTypeBulkWithFilters, PlanNewParamsPricesPriceModelTypePackage, PlanNewParamsPricesPriceModelTypeMatrix, PlanNewParamsPricesPriceModelTypeThresholdTotalAmount, PlanNewParamsPricesPriceModelTypeTieredPackage, PlanNewParamsPricesPriceModelTypeTieredWithMinimum, PlanNewParamsPricesPriceModelTypeGroupedTiered, PlanNewParamsPricesPriceModelTypeTieredPackageWithMinimum, PlanNewParamsPricesPriceModelTypePackageWithAllocation, PlanNewParamsPricesPriceModelTypeUnitWithPercent, PlanNewParamsPricesPriceModelTypeMatrixWithAllocation, PlanNewParamsPricesPriceModelTypeTieredWithProration, PlanNewParamsPricesPriceModelTypeUnitWithProration, PlanNewParamsPricesPriceModelTypeGroupedAllocation, PlanNewParamsPricesPriceModelTypeBulkWithProration, PlanNewParamsPricesPriceModelTypeGroupedWithProratedMinimum, PlanNewParamsPricesPriceModelTypeGroupedWithMeteredMinimum, PlanNewParamsPricesPriceModelTypeGroupedWithMinMaxThresholds, PlanNewParamsPricesPriceModelTypeMatrixWithDisplayName, PlanNewParamsPricesPriceModelTypeGroupedTieredPackage, PlanNewParamsPricesPriceModelTypeMaxGroupTieredPackage, PlanNewParamsPricesPriceModelTypeScalableMatrixWithUnitPricing, PlanNewParamsPricesPriceModelTypeScalableMatrixWithTieredPricing, PlanNewParamsPricesPriceModelTypeCumulativeGroupedBulk, PlanNewParamsPricesPriceModelTypeCumulativeGroupedAllocation, PlanNewParamsPricesPriceModelTypeMinimumComposite, PlanNewParamsPricesPriceModelTypePercent, PlanNewParamsPricesPriceModelTypeEventOutput:
		return true
	}
	return false
}

type PlanNewParamsAdjustment struct {
	// The definition of a new adjustment to create and add to the plan.
	Adjustment param.Field[PlanNewParamsAdjustmentsAdjustmentUnion] `json:"adjustment" api:"required"`
	// The phase to add this adjustment to.
	PlanPhaseOrder param.Field[int64] `json:"plan_phase_order"`
}

func (r PlanNewParamsAdjustment) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// The definition of a new adjustment to create and add to the plan.
type PlanNewParamsAdjustmentsAdjustment struct {
	AdjustmentType param.Field[PlanNewParamsAdjustmentsAdjustmentAdjustmentType] `json:"adjustment_type" api:"required"`
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
	Order param.Field[int64] `json:"order" api:"required"`
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
	// An optional user-defined description of the plan.
	Description param.Field[string] `json:"description"`
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
