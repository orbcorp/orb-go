// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package orb

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"net/url"
	"reflect"
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
	opts = append(r.Options[:], opts...)
	path := "plans"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return
}

// This endpoint can be used to update the `external_plan_id`, and `metadata` of an
// existing plan.
//
// Other fields on a plan are currently immutable.
func (r *PlanService) Update(ctx context.Context, planID string, body PlanUpdateParams, opts ...option.RequestOption) (res *Plan, err error) {
	opts = append(r.Options[:], opts...)
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
	opts = append(r.Options[:], opts...)
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
	opts = append(r.Options[:], opts...)
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
	BasePlan    PlanBasePlan     `json:"base_plan,required,nullable"`
	// The parent plan id if the given plan was created by overriding one or more of
	// the parent's prices
	BasePlanID string    `json:"base_plan_id,required,nullable"`
	CreatedAt  time.Time `json:"created_at,required" format:"date-time"`
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
	ID             string                        `json:"id,required"`
	AdjustmentType PlanAdjustmentsAdjustmentType `json:"adjustment_type,required"`
	// This field can have the runtime type of [[]string].
	AppliesToPriceIDs interface{} `json:"applies_to_price_ids,required"`
	// This field can have the runtime type of [[]shared.TransformPriceFilter].
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
	UsageDiscount float64            `json:"usage_discount"`
	JSON          planAdjustmentJSON `json:"-"`
	union         PlanAdjustmentsUnion
}

// planAdjustmentJSON contains the JSON metadata for the struct [PlanAdjustment]
type planAdjustmentJSON struct {
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

type PlanNewParams struct {
	// An ISO 4217 currency string for invoices generated by subscriptions on this
	// plan.
	Currency param.Field[string] `json:"currency,required"`
	Name     param.Field[string] `json:"name,required"`
	// Prices for this plan. If the plan has phases, this includes prices across all
	// phases of the plan.
	Prices param.Field[[]PlanNewParamsPriceUnion] `json:"prices,required"`
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
	// The status of the plan to create (either active or draft). If not specified,
	// this defaults to active.
	Status param.Field[PlanNewParamsStatus] `json:"status"`
}

func (r PlanNewParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type PlanNewParamsPrice struct {
	// The cadence to bill for this price on.
	Cadence param.Field[PlanNewParamsPricesCadence] `json:"cadence,required"`
	// The id of the item the price will be associated with.
	ItemID    param.Field[string]                       `json:"item_id,required"`
	ModelType param.Field[PlanNewParamsPricesModelType] `json:"model_type,required"`
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
	BPSConfig                 param.Field[shared.BPSConfigParam]                    `json:"bps_config"`
	BulkBPSConfig             param.Field[shared.BulkBPSConfigParam]                `json:"bulk_bps_config"`
	BulkConfig                param.Field[shared.BulkConfigParam]                   `json:"bulk_config"`
	BulkWithProrationConfig   param.Field[interface{}]                              `json:"bulk_with_proration_config"`
	// The per unit conversion rate of the price currency to the invoicing currency.
	ConversionRate              param.Field[float64]     `json:"conversion_rate"`
	CumulativeGroupedBulkConfig param.Field[interface{}] `json:"cumulative_grouped_bulk_config"`
	// An ISO 4217 currency string, or custom pricing unit identifier, in which this
	// price is billed.
	Currency param.Field[string] `json:"currency"`
	// For dimensional price: specifies a price group and dimension values
	DimensionalPriceConfiguration param.Field[shared.NewDimensionalPriceConfigurationParam] `json:"dimensional_price_configuration"`
	// An alias for the price.
	ExternalPriceID param.Field[string] `json:"external_price_id"`
	// If the Price represents a fixed cost, this represents the quantity of units
	// applied.
	FixedPriceQuantity               param.Field[float64]     `json:"fixed_price_quantity"`
	GroupedAllocationConfig          param.Field[interface{}] `json:"grouped_allocation_config"`
	GroupedTieredConfig              param.Field[interface{}] `json:"grouped_tiered_config"`
	GroupedTieredPackageConfig       param.Field[interface{}] `json:"grouped_tiered_package_config"`
	GroupedWithMeteredMinimumConfig  param.Field[interface{}] `json:"grouped_with_metered_minimum_config"`
	GroupedWithProratedMinimumConfig param.Field[interface{}] `json:"grouped_with_prorated_minimum_config"`
	// The property used to group this price on an invoice
	InvoiceGroupingKey param.Field[string] `json:"invoice_grouping_key"`
	// Within each billing cycle, specifies the cadence at which invoices are produced.
	// If unspecified, a single invoice is produced per billing cycle.
	InvoicingCycleConfiguration           param.Field[shared.NewBillingCycleConfigurationParam] `json:"invoicing_cycle_configuration"`
	MatrixConfig                          param.Field[shared.MatrixConfigParam]                 `json:"matrix_config"`
	MatrixWithAllocationConfig            param.Field[shared.MatrixWithAllocationConfigParam]   `json:"matrix_with_allocation_config"`
	MatrixWithDisplayNameConfig           param.Field[interface{}]                              `json:"matrix_with_display_name_config"`
	MaxGroupTieredPackageConfig           param.Field[interface{}]                              `json:"max_group_tiered_package_config"`
	Metadata                              param.Field[interface{}]                              `json:"metadata"`
	PackageConfig                         param.Field[shared.PackageConfigParam]                `json:"package_config"`
	PackageWithAllocationConfig           param.Field[interface{}]                              `json:"package_with_allocation_config"`
	ScalableMatrixWithTieredPricingConfig param.Field[interface{}]                              `json:"scalable_matrix_with_tiered_pricing_config"`
	ScalableMatrixWithUnitPricingConfig   param.Field[interface{}]                              `json:"scalable_matrix_with_unit_pricing_config"`
	ThresholdTotalAmountConfig            param.Field[interface{}]                              `json:"threshold_total_amount_config"`
	TieredBPSConfig                       param.Field[shared.TieredBPSConfigParam]              `json:"tiered_bps_config"`
	TieredConfig                          param.Field[shared.TieredConfigParam]                 `json:"tiered_config"`
	TieredPackageConfig                   param.Field[interface{}]                              `json:"tiered_package_config"`
	TieredPackageWithMinimumConfig        param.Field[interface{}]                              `json:"tiered_package_with_minimum_config"`
	TieredWithMinimumConfig               param.Field[interface{}]                              `json:"tiered_with_minimum_config"`
	TieredWithProrationConfig             param.Field[interface{}]                              `json:"tiered_with_proration_config"`
	UnitConfig                            param.Field[shared.UnitConfigParam]                   `json:"unit_config"`
	UnitWithPercentConfig                 param.Field[interface{}]                              `json:"unit_with_percent_config"`
	UnitWithProrationConfig               param.Field[interface{}]                              `json:"unit_with_proration_config"`
}

func (r PlanNewParamsPrice) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r PlanNewParamsPrice) ImplementsPlanNewParamsPriceUnion() {}

// Satisfied by [shared.NewPlanUnitPriceParam], [shared.NewPlanPackagePriceParam],
// [shared.NewPlanMatrixPriceParam], [shared.NewPlanTieredPriceParam],
// [shared.NewPlanTieredBPSPriceParam], [shared.NewPlanBPSPriceParam],
// [shared.NewPlanBulkBPSPriceParam], [shared.NewPlanBulkPriceParam],
// [shared.NewPlanThresholdTotalAmountPriceParam],
// [shared.NewPlanTieredPackagePriceParam],
// [shared.NewPlanTieredWithMinimumPriceParam],
// [shared.NewPlanUnitWithPercentPriceParam],
// [shared.NewPlanPackageWithAllocationPriceParam],
// [shared.NewPlanTierWithProrationPriceParam],
// [shared.NewPlanUnitWithProrationPriceParam],
// [shared.NewPlanGroupedAllocationPriceParam],
// [shared.NewPlanGroupedWithProratedMinimumPriceParam],
// [shared.NewPlanGroupedWithMeteredMinimumPriceParam],
// [shared.NewPlanMatrixWithDisplayNamePriceParam],
// [shared.NewPlanBulkWithProrationPriceParam],
// [shared.NewPlanGroupedTieredPackagePriceParam],
// [shared.NewPlanMaxGroupTieredPackagePriceParam],
// [shared.NewPlanScalableMatrixWithUnitPricingPriceParam],
// [shared.NewPlanScalableMatrixWithTieredPricingPriceParam],
// [shared.NewPlanCumulativeGroupedBulkPriceParam],
// [shared.NewPlanTieredPackageWithMinimumPriceParam],
// [shared.NewPlanMatrixWithAllocationPriceParam],
// [shared.NewPlanGroupedTieredPriceParam], [PlanNewParamsPrice].
type PlanNewParamsPriceUnion interface {
	ImplementsPlanNewParamsPriceUnion()
}

// The cadence to bill for this price on.
type PlanNewParamsPricesCadence string

const (
	PlanNewParamsPricesCadenceAnnual     PlanNewParamsPricesCadence = "annual"
	PlanNewParamsPricesCadenceSemiAnnual PlanNewParamsPricesCadence = "semi_annual"
	PlanNewParamsPricesCadenceMonthly    PlanNewParamsPricesCadence = "monthly"
	PlanNewParamsPricesCadenceQuarterly  PlanNewParamsPricesCadence = "quarterly"
	PlanNewParamsPricesCadenceOneTime    PlanNewParamsPricesCadence = "one_time"
	PlanNewParamsPricesCadenceCustom     PlanNewParamsPricesCadence = "custom"
)

func (r PlanNewParamsPricesCadence) IsKnown() bool {
	switch r {
	case PlanNewParamsPricesCadenceAnnual, PlanNewParamsPricesCadenceSemiAnnual, PlanNewParamsPricesCadenceMonthly, PlanNewParamsPricesCadenceQuarterly, PlanNewParamsPricesCadenceOneTime, PlanNewParamsPricesCadenceCustom:
		return true
	}
	return false
}

type PlanNewParamsPricesModelType string

const (
	PlanNewParamsPricesModelTypeUnit                            PlanNewParamsPricesModelType = "unit"
	PlanNewParamsPricesModelTypePackage                         PlanNewParamsPricesModelType = "package"
	PlanNewParamsPricesModelTypeMatrix                          PlanNewParamsPricesModelType = "matrix"
	PlanNewParamsPricesModelTypeTiered                          PlanNewParamsPricesModelType = "tiered"
	PlanNewParamsPricesModelTypeTieredBPS                       PlanNewParamsPricesModelType = "tiered_bps"
	PlanNewParamsPricesModelTypeBPS                             PlanNewParamsPricesModelType = "bps"
	PlanNewParamsPricesModelTypeBulkBPS                         PlanNewParamsPricesModelType = "bulk_bps"
	PlanNewParamsPricesModelTypeBulk                            PlanNewParamsPricesModelType = "bulk"
	PlanNewParamsPricesModelTypeThresholdTotalAmount            PlanNewParamsPricesModelType = "threshold_total_amount"
	PlanNewParamsPricesModelTypeTieredPackage                   PlanNewParamsPricesModelType = "tiered_package"
	PlanNewParamsPricesModelTypeTieredWithMinimum               PlanNewParamsPricesModelType = "tiered_with_minimum"
	PlanNewParamsPricesModelTypeUnitWithPercent                 PlanNewParamsPricesModelType = "unit_with_percent"
	PlanNewParamsPricesModelTypePackageWithAllocation           PlanNewParamsPricesModelType = "package_with_allocation"
	PlanNewParamsPricesModelTypeTieredWithProration             PlanNewParamsPricesModelType = "tiered_with_proration"
	PlanNewParamsPricesModelTypeUnitWithProration               PlanNewParamsPricesModelType = "unit_with_proration"
	PlanNewParamsPricesModelTypeGroupedAllocation               PlanNewParamsPricesModelType = "grouped_allocation"
	PlanNewParamsPricesModelTypeGroupedWithProratedMinimum      PlanNewParamsPricesModelType = "grouped_with_prorated_minimum"
	PlanNewParamsPricesModelTypeGroupedWithMeteredMinimum       PlanNewParamsPricesModelType = "grouped_with_metered_minimum"
	PlanNewParamsPricesModelTypeMatrixWithDisplayName           PlanNewParamsPricesModelType = "matrix_with_display_name"
	PlanNewParamsPricesModelTypeBulkWithProration               PlanNewParamsPricesModelType = "bulk_with_proration"
	PlanNewParamsPricesModelTypeGroupedTieredPackage            PlanNewParamsPricesModelType = "grouped_tiered_package"
	PlanNewParamsPricesModelTypeMaxGroupTieredPackage           PlanNewParamsPricesModelType = "max_group_tiered_package"
	PlanNewParamsPricesModelTypeScalableMatrixWithUnitPricing   PlanNewParamsPricesModelType = "scalable_matrix_with_unit_pricing"
	PlanNewParamsPricesModelTypeScalableMatrixWithTieredPricing PlanNewParamsPricesModelType = "scalable_matrix_with_tiered_pricing"
	PlanNewParamsPricesModelTypeCumulativeGroupedBulk           PlanNewParamsPricesModelType = "cumulative_grouped_bulk"
	PlanNewParamsPricesModelTypeTieredPackageWithMinimum        PlanNewParamsPricesModelType = "tiered_package_with_minimum"
	PlanNewParamsPricesModelTypeMatrixWithAllocation            PlanNewParamsPricesModelType = "matrix_with_allocation"
	PlanNewParamsPricesModelTypeGroupedTiered                   PlanNewParamsPricesModelType = "grouped_tiered"
)

func (r PlanNewParamsPricesModelType) IsKnown() bool {
	switch r {
	case PlanNewParamsPricesModelTypeUnit, PlanNewParamsPricesModelTypePackage, PlanNewParamsPricesModelTypeMatrix, PlanNewParamsPricesModelTypeTiered, PlanNewParamsPricesModelTypeTieredBPS, PlanNewParamsPricesModelTypeBPS, PlanNewParamsPricesModelTypeBulkBPS, PlanNewParamsPricesModelTypeBulk, PlanNewParamsPricesModelTypeThresholdTotalAmount, PlanNewParamsPricesModelTypeTieredPackage, PlanNewParamsPricesModelTypeTieredWithMinimum, PlanNewParamsPricesModelTypeUnitWithPercent, PlanNewParamsPricesModelTypePackageWithAllocation, PlanNewParamsPricesModelTypeTieredWithProration, PlanNewParamsPricesModelTypeUnitWithProration, PlanNewParamsPricesModelTypeGroupedAllocation, PlanNewParamsPricesModelTypeGroupedWithProratedMinimum, PlanNewParamsPricesModelTypeGroupedWithMeteredMinimum, PlanNewParamsPricesModelTypeMatrixWithDisplayName, PlanNewParamsPricesModelTypeBulkWithProration, PlanNewParamsPricesModelTypeGroupedTieredPackage, PlanNewParamsPricesModelTypeMaxGroupTieredPackage, PlanNewParamsPricesModelTypeScalableMatrixWithUnitPricing, PlanNewParamsPricesModelTypeScalableMatrixWithTieredPricing, PlanNewParamsPricesModelTypeCumulativeGroupedBulk, PlanNewParamsPricesModelTypeTieredPackageWithMinimum, PlanNewParamsPricesModelTypeMatrixWithAllocation, PlanNewParamsPricesModelTypeGroupedTiered:
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
