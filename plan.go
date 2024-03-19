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
	"github.com/orbcorp/orb-go/internal/param"
	"github.com/orbcorp/orb-go/internal/requestconfig"
	"github.com/orbcorp/orb-go/internal/shared"
	"github.com/orbcorp/orb-go/option"
)

// PlanService contains methods and other services that help with interacting with
// the orb API. Note, unlike clients, this service does not read variables from the
// environment automatically. You should not instantiate this service directly, and
// instead use the [NewPlanService] method instead.
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
// Other fields on a customer are currently immutable.
func (r *PlanService) Update(ctx context.Context, planID string, body PlanUpdateParams, opts ...option.RequestOption) (res *Plan, err error) {
	opts = append(r.Options[:], opts...)
	path := fmt.Sprintf("plans/%s", planID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPut, path, body, &res, opts...)
	return
}

// This endpoint returns a list of all [plans](../guides/concepts##plan-and-price)
// for an account in a list format. The list of plans is ordered starting from the
// most recently created plan. The response also includes
// [`pagination_metadata`](../reference/pagination), which lets the caller retrieve
// the next page of results if they exist.
func (r *PlanService) List(ctx context.Context, query PlanListParams, opts ...option.RequestOption) (res *shared.Page[Plan], err error) {
	var raw *http.Response
	opts = append(r.Options, opts...)
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

// This endpoint returns a list of all [plans](../guides/concepts##plan-and-price)
// for an account in a list format. The list of plans is ordered starting from the
// most recently created plan. The response also includes
// [`pagination_metadata`](../reference/pagination), which lets the caller retrieve
// the next page of results if they exist.
func (r *PlanService) ListAutoPaging(ctx context.Context, query PlanListParams, opts ...option.RequestOption) *shared.PageAutoPager[Plan] {
	return shared.NewPageAutoPager(r.List(ctx, query, opts...))
}

// This endpoint is used to fetch [plan](../guides/concepts##plan-and-price)
// details given a plan identifier. It returns information about the prices
// included in the plan and their configuration, as well as the product that the
// plan is attached to.
//
// ## Serialized prices
//
// Orb supports a few different pricing models out of the box. Each of these models
// is serialized differently in a given [Price](../guides/concepts#plan-and-price)
// object. The `model_type` field determines the key for the configuration object
// that is present. A detailed explanation of price types can be found in the
// [Price schema](../guides/concepts#plan-and-price).
//
// ## Phases
//
// Orb supports plan phases, also known as contract ramps. For plans with phases,
// the serialized prices refer to all prices across all phases.
func (r *PlanService) Fetch(ctx context.Context, planID string, opts ...option.RequestOption) (res *Plan, err error) {
	opts = append(r.Options[:], opts...)
	path := fmt.Sprintf("plans/%s", planID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

// The [Plan](../guides/core-concepts.mdx#plan-and-price) resource represents a
// plan that can be subscribed to by a customer. Plans define the billing behavior
// of the subscription. You can see more about how to configure prices in the
// [Price resource](/reference/price).
type Plan struct {
	ID       string       `json:"id,required"`
	BasePlan PlanBasePlan `json:"base_plan,required,nullable"`
	// The parent plan id if the given plan was created by overriding one or more of
	// the parent's prices
	BasePlanID string    `json:"base_plan_id,required,nullable"`
	CreatedAt  time.Time `json:"created_at,required" format:"date-time"`
	// An ISO 4217 currency string or custom pricing unit (`credits`) for this plan's
	// prices.
	Currency string `json:"currency,required"`
	// The default memo text on the invoices corresponding to subscriptions on this
	// plan. Note that each subscription may configure its own memo.
	DefaultInvoiceMemo string          `json:"default_invoice_memo,required,nullable"`
	Description        string          `json:"description,required"`
	Discount           shared.Discount `json:"discount,required,nullable"`
	// An optional user-defined ID for this plan resource, used throughout the system
	// as an alias for this Plan. Use this field to identify a plan by an existing
	// identifier in your system.
	ExternalPlanID string `json:"external_plan_id,required,nullable"`
	// An ISO 4217 currency string for which this plan is billed in. Matches `currency`
	// unless `currency` is a custom pricing unit.
	InvoicingCurrency string      `json:"invoicing_currency,required"`
	Maximum           PlanMaximum `json:"maximum,required,nullable"`
	MaximumAmount     string      `json:"maximum_amount,required,nullable"`
	// User specified key-value pairs for the resource. If not present, this defaults
	// to an empty dictionary. Individual keys can be removed by setting the value to
	// `null`, and the entire metadata mapping can be cleared by setting `metadata` to
	// `null`.
	Metadata      map[string]string `json:"metadata,required"`
	Minimum       PlanMinimum       `json:"minimum,required,nullable"`
	MinimumAmount string            `json:"minimum_amount,required,nullable"`
	Name          string            `json:"name,required"`
	// Determines the difference between the invoice issue date and the due date. A
	// value of "0" here signifies that invoices are due on issue, whereas a value of
	// "30" means that the customer has a month to pay the invoice before its overdue.
	// Note that individual subscriptions or invoices may set a different net terms
	// configuration.
	NetTerms   int64           `json:"net_terms,required,nullable"`
	PlanPhases []PlanPlanPhase `json:"plan_phases,required,nullable"`
	// Prices for this plan. If the plan has phases, this includes prices across all
	// phases of the plan.
	Prices      []Price         `json:"prices,required"`
	Product     PlanProduct     `json:"product,required"`
	Status      PlanStatus      `json:"status,required"`
	TrialConfig PlanTrialConfig `json:"trial_config,required"`
	Version     int64           `json:"version,required"`
	JSON        planJSON        `json:"-"`
}

// planJSON contains the JSON metadata for the struct [Plan]
type planJSON struct {
	ID                 apijson.Field
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

type PlanMaximum struct {
	// List of price_ids that this maximum amount applies to. For plan/plan phase
	// maximums, this can be a subset of prices.
	AppliesToPriceIDs []string `json:"applies_to_price_ids,required"`
	// Maximum amount applied
	MaximumAmount string          `json:"maximum_amount,required"`
	JSON          planMaximumJSON `json:"-"`
}

// planMaximumJSON contains the JSON metadata for the struct [PlanMaximum]
type planMaximumJSON struct {
	AppliesToPriceIDs apijson.Field
	MaximumAmount     apijson.Field
	raw               string
	ExtraFields       map[string]apijson.Field
}

func (r *PlanMaximum) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r planMaximumJSON) RawJSON() string {
	return r.raw
}

type PlanMinimum struct {
	// List of price_ids that this minimum amount applies to. For plan/plan phase
	// minimums, this can be a subset of prices.
	AppliesToPriceIDs []string `json:"applies_to_price_ids,required"`
	// Minimum amount applied
	MinimumAmount string          `json:"minimum_amount,required"`
	JSON          planMinimumJSON `json:"-"`
}

// planMinimumJSON contains the JSON metadata for the struct [PlanMinimum]
type planMinimumJSON struct {
	AppliesToPriceIDs apijson.Field
	MinimumAmount     apijson.Field
	raw               string
	ExtraFields       map[string]apijson.Field
}

func (r *PlanMinimum) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r planMinimumJSON) RawJSON() string {
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
	Maximum       PlanPlanPhasesMaximum      `json:"maximum,required,nullable"`
	MaximumAmount string                     `json:"maximum_amount,required,nullable"`
	Minimum       PlanPlanPhasesMinimum      `json:"minimum,required,nullable"`
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
	PlanPlanPhasesDurationUnitDaily     PlanPlanPhasesDurationUnit = "daily"
	PlanPlanPhasesDurationUnitMonthly   PlanPlanPhasesDurationUnit = "monthly"
	PlanPlanPhasesDurationUnitQuarterly PlanPlanPhasesDurationUnit = "quarterly"
	PlanPlanPhasesDurationUnitAnnual    PlanPlanPhasesDurationUnit = "annual"
)

func (r PlanPlanPhasesDurationUnit) IsKnown() bool {
	switch r {
	case PlanPlanPhasesDurationUnitDaily, PlanPlanPhasesDurationUnitMonthly, PlanPlanPhasesDurationUnitQuarterly, PlanPlanPhasesDurationUnitAnnual:
		return true
	}
	return false
}

type PlanPlanPhasesMaximum struct {
	// List of price_ids that this maximum amount applies to. For plan/plan phase
	// maximums, this can be a subset of prices.
	AppliesToPriceIDs []string `json:"applies_to_price_ids,required"`
	// Maximum amount applied
	MaximumAmount string                    `json:"maximum_amount,required"`
	JSON          planPlanPhasesMaximumJSON `json:"-"`
}

// planPlanPhasesMaximumJSON contains the JSON metadata for the struct
// [PlanPlanPhasesMaximum]
type planPlanPhasesMaximumJSON struct {
	AppliesToPriceIDs apijson.Field
	MaximumAmount     apijson.Field
	raw               string
	ExtraFields       map[string]apijson.Field
}

func (r *PlanPlanPhasesMaximum) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r planPlanPhasesMaximumJSON) RawJSON() string {
	return r.raw
}

type PlanPlanPhasesMinimum struct {
	// List of price_ids that this minimum amount applies to. For plan/plan phase
	// minimums, this can be a subset of prices.
	AppliesToPriceIDs []string `json:"applies_to_price_ids,required"`
	// Minimum amount applied
	MinimumAmount string                    `json:"minimum_amount,required"`
	JSON          planPlanPhasesMinimumJSON `json:"-"`
}

// planPlanPhasesMinimumJSON contains the JSON metadata for the struct
// [PlanPlanPhasesMinimum]
type planPlanPhasesMinimumJSON struct {
	AppliesToPriceIDs apijson.Field
	MinimumAmount     apijson.Field
	raw               string
	ExtraFields       map[string]apijson.Field
}

func (r *PlanPlanPhasesMinimum) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r planPlanPhasesMinimumJSON) RawJSON() string {
	return r.raw
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
	Prices param.Field[[]PlanNewParamsPrice] `json:"prices,required"`
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
}

func (r PlanNewParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Satisfied by [PlanNewParamsPricesNewPlanUnitPrice],
// [PlanNewParamsPricesNewPlanPackagePrice],
// [PlanNewParamsPricesNewPlanMatrixPrice],
// [PlanNewParamsPricesNewPlanTieredPrice],
// [PlanNewParamsPricesNewPlanTieredBpsPrice],
// [PlanNewParamsPricesNewPlanBpsPrice], [PlanNewParamsPricesNewPlanBulkBpsPrice],
// [PlanNewParamsPricesNewPlanBulkPrice],
// [PlanNewParamsPricesNewPlanThresholdTotalAmountPrice],
// [PlanNewParamsPricesNewPlanTieredPackagePrice],
// [PlanNewParamsPricesNewPlanTieredWithMinimumPrice],
// [PlanNewParamsPricesNewPlanUnitWithPercentPrice],
// [PlanNewParamsPricesNewPlanPackageWithAllocationPrice].
type PlanNewParamsPrice interface {
	implementsPlanNewParamsPrice()
}

type PlanNewParamsPricesNewPlanUnitPrice struct {
	// The cadence to bill for this price on.
	Cadence param.Field[PlanNewParamsPricesNewPlanUnitPriceCadence] `json:"cadence,required"`
	// The id of the item the plan will be associated with.
	ItemID    param.Field[string]                                       `json:"item_id,required"`
	ModelType param.Field[PlanNewParamsPricesNewPlanUnitPriceModelType] `json:"model_type,required"`
	// The name of the price.
	Name       param.Field[string]                                        `json:"name,required"`
	UnitConfig param.Field[PlanNewParamsPricesNewPlanUnitPriceUnitConfig] `json:"unit_config,required"`
	// The id of the billable metric for the price. Only needed if the price is
	// usage-based.
	BillableMetricID param.Field[string] `json:"billable_metric_id"`
	// If the Price represents a fixed cost, the price will be billed in-advance if
	// this is true, and in-arrears if this is false.
	BilledInAdvance param.Field[bool] `json:"billed_in_advance"`
	// An alias for the price.
	ExternalPriceID param.Field[string] `json:"external_price_id"`
	// If the Price represents a fixed cost, this represents the quantity of units
	// applied.
	FixedPriceQuantity param.Field[float64] `json:"fixed_price_quantity"`
	// The property used to group this price on an invoice
	InvoiceGroupingKey param.Field[string] `json:"invoice_grouping_key"`
}

func (r PlanNewParamsPricesNewPlanUnitPrice) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r PlanNewParamsPricesNewPlanUnitPrice) implementsPlanNewParamsPrice() {}

// The cadence to bill for this price on.
type PlanNewParamsPricesNewPlanUnitPriceCadence string

const (
	PlanNewParamsPricesNewPlanUnitPriceCadenceAnnual    PlanNewParamsPricesNewPlanUnitPriceCadence = "annual"
	PlanNewParamsPricesNewPlanUnitPriceCadenceMonthly   PlanNewParamsPricesNewPlanUnitPriceCadence = "monthly"
	PlanNewParamsPricesNewPlanUnitPriceCadenceQuarterly PlanNewParamsPricesNewPlanUnitPriceCadence = "quarterly"
	PlanNewParamsPricesNewPlanUnitPriceCadenceOneTime   PlanNewParamsPricesNewPlanUnitPriceCadence = "one_time"
)

func (r PlanNewParamsPricesNewPlanUnitPriceCadence) IsKnown() bool {
	switch r {
	case PlanNewParamsPricesNewPlanUnitPriceCadenceAnnual, PlanNewParamsPricesNewPlanUnitPriceCadenceMonthly, PlanNewParamsPricesNewPlanUnitPriceCadenceQuarterly, PlanNewParamsPricesNewPlanUnitPriceCadenceOneTime:
		return true
	}
	return false
}

type PlanNewParamsPricesNewPlanUnitPriceModelType string

const (
	PlanNewParamsPricesNewPlanUnitPriceModelTypeUnit PlanNewParamsPricesNewPlanUnitPriceModelType = "unit"
)

func (r PlanNewParamsPricesNewPlanUnitPriceModelType) IsKnown() bool {
	switch r {
	case PlanNewParamsPricesNewPlanUnitPriceModelTypeUnit:
		return true
	}
	return false
}

type PlanNewParamsPricesNewPlanUnitPriceUnitConfig struct {
	// Rate per unit of usage
	UnitAmount param.Field[string] `json:"unit_amount,required"`
	// Multiplier to scale rated quantity by
	ScalingFactor param.Field[float64] `json:"scaling_factor"`
}

func (r PlanNewParamsPricesNewPlanUnitPriceUnitConfig) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type PlanNewParamsPricesNewPlanPackagePrice struct {
	// The cadence to bill for this price on.
	Cadence param.Field[PlanNewParamsPricesNewPlanPackagePriceCadence] `json:"cadence,required"`
	// The id of the item the plan will be associated with.
	ItemID    param.Field[string]                                          `json:"item_id,required"`
	ModelType param.Field[PlanNewParamsPricesNewPlanPackagePriceModelType] `json:"model_type,required"`
	// The name of the price.
	Name          param.Field[string]                                              `json:"name,required"`
	PackageConfig param.Field[PlanNewParamsPricesNewPlanPackagePricePackageConfig] `json:"package_config,required"`
	// The id of the billable metric for the price. Only needed if the price is
	// usage-based.
	BillableMetricID param.Field[string] `json:"billable_metric_id"`
	// If the Price represents a fixed cost, the price will be billed in-advance if
	// this is true, and in-arrears if this is false.
	BilledInAdvance param.Field[bool] `json:"billed_in_advance"`
	// An alias for the price.
	ExternalPriceID param.Field[string] `json:"external_price_id"`
	// If the Price represents a fixed cost, this represents the quantity of units
	// applied.
	FixedPriceQuantity param.Field[float64] `json:"fixed_price_quantity"`
	// The property used to group this price on an invoice
	InvoiceGroupingKey param.Field[string] `json:"invoice_grouping_key"`
}

func (r PlanNewParamsPricesNewPlanPackagePrice) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r PlanNewParamsPricesNewPlanPackagePrice) implementsPlanNewParamsPrice() {}

// The cadence to bill for this price on.
type PlanNewParamsPricesNewPlanPackagePriceCadence string

const (
	PlanNewParamsPricesNewPlanPackagePriceCadenceAnnual    PlanNewParamsPricesNewPlanPackagePriceCadence = "annual"
	PlanNewParamsPricesNewPlanPackagePriceCadenceMonthly   PlanNewParamsPricesNewPlanPackagePriceCadence = "monthly"
	PlanNewParamsPricesNewPlanPackagePriceCadenceQuarterly PlanNewParamsPricesNewPlanPackagePriceCadence = "quarterly"
	PlanNewParamsPricesNewPlanPackagePriceCadenceOneTime   PlanNewParamsPricesNewPlanPackagePriceCadence = "one_time"
)

func (r PlanNewParamsPricesNewPlanPackagePriceCadence) IsKnown() bool {
	switch r {
	case PlanNewParamsPricesNewPlanPackagePriceCadenceAnnual, PlanNewParamsPricesNewPlanPackagePriceCadenceMonthly, PlanNewParamsPricesNewPlanPackagePriceCadenceQuarterly, PlanNewParamsPricesNewPlanPackagePriceCadenceOneTime:
		return true
	}
	return false
}

type PlanNewParamsPricesNewPlanPackagePriceModelType string

const (
	PlanNewParamsPricesNewPlanPackagePriceModelTypePackage PlanNewParamsPricesNewPlanPackagePriceModelType = "package"
)

func (r PlanNewParamsPricesNewPlanPackagePriceModelType) IsKnown() bool {
	switch r {
	case PlanNewParamsPricesNewPlanPackagePriceModelTypePackage:
		return true
	}
	return false
}

type PlanNewParamsPricesNewPlanPackagePricePackageConfig struct {
	// A currency amount to rate usage by
	PackageAmount param.Field[string] `json:"package_amount,required"`
	// An integer amount to represent package size. For example, 1000 here would divide
	// usage by 1000 before multiplying by package_amount in rating
	PackageSize param.Field[int64] `json:"package_size"`
}

func (r PlanNewParamsPricesNewPlanPackagePricePackageConfig) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type PlanNewParamsPricesNewPlanMatrixPrice struct {
	// The cadence to bill for this price on.
	Cadence param.Field[PlanNewParamsPricesNewPlanMatrixPriceCadence] `json:"cadence,required"`
	// The id of the item the plan will be associated with.
	ItemID       param.Field[string]                                            `json:"item_id,required"`
	MatrixConfig param.Field[PlanNewParamsPricesNewPlanMatrixPriceMatrixConfig] `json:"matrix_config,required"`
	ModelType    param.Field[PlanNewParamsPricesNewPlanMatrixPriceModelType]    `json:"model_type,required"`
	// The name of the price.
	Name param.Field[string] `json:"name,required"`
	// The id of the billable metric for the price. Only needed if the price is
	// usage-based.
	BillableMetricID param.Field[string] `json:"billable_metric_id"`
	// If the Price represents a fixed cost, the price will be billed in-advance if
	// this is true, and in-arrears if this is false.
	BilledInAdvance param.Field[bool] `json:"billed_in_advance"`
	// An alias for the price.
	ExternalPriceID param.Field[string] `json:"external_price_id"`
	// If the Price represents a fixed cost, this represents the quantity of units
	// applied.
	FixedPriceQuantity param.Field[float64] `json:"fixed_price_quantity"`
	// The property used to group this price on an invoice
	InvoiceGroupingKey param.Field[string] `json:"invoice_grouping_key"`
}

func (r PlanNewParamsPricesNewPlanMatrixPrice) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r PlanNewParamsPricesNewPlanMatrixPrice) implementsPlanNewParamsPrice() {}

// The cadence to bill for this price on.
type PlanNewParamsPricesNewPlanMatrixPriceCadence string

const (
	PlanNewParamsPricesNewPlanMatrixPriceCadenceAnnual    PlanNewParamsPricesNewPlanMatrixPriceCadence = "annual"
	PlanNewParamsPricesNewPlanMatrixPriceCadenceMonthly   PlanNewParamsPricesNewPlanMatrixPriceCadence = "monthly"
	PlanNewParamsPricesNewPlanMatrixPriceCadenceQuarterly PlanNewParamsPricesNewPlanMatrixPriceCadence = "quarterly"
	PlanNewParamsPricesNewPlanMatrixPriceCadenceOneTime   PlanNewParamsPricesNewPlanMatrixPriceCadence = "one_time"
)

func (r PlanNewParamsPricesNewPlanMatrixPriceCadence) IsKnown() bool {
	switch r {
	case PlanNewParamsPricesNewPlanMatrixPriceCadenceAnnual, PlanNewParamsPricesNewPlanMatrixPriceCadenceMonthly, PlanNewParamsPricesNewPlanMatrixPriceCadenceQuarterly, PlanNewParamsPricesNewPlanMatrixPriceCadenceOneTime:
		return true
	}
	return false
}

type PlanNewParamsPricesNewPlanMatrixPriceMatrixConfig struct {
	// Default per unit rate for any usage not bucketed into a specified matrix_value
	DefaultUnitAmount param.Field[string] `json:"default_unit_amount,required"`
	// One or two event property values to evaluate matrix groups by
	Dimensions param.Field[[]string] `json:"dimensions,required"`
	// Matrix values for specified matrix grouping keys
	MatrixValues param.Field[[]PlanNewParamsPricesNewPlanMatrixPriceMatrixConfigMatrixValue] `json:"matrix_values,required"`
	// Default optional multiplier to scale rated quantities that fall into the default
	// bucket by
	ScalingFactor param.Field[float64] `json:"scaling_factor"`
}

func (r PlanNewParamsPricesNewPlanMatrixPriceMatrixConfig) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type PlanNewParamsPricesNewPlanMatrixPriceMatrixConfigMatrixValue struct {
	// One or two matrix keys to filter usage to this Matrix value by. For example,
	// ["region", "tier"] could be used to filter cloud usage by a cloud region and an
	// instance tier.
	DimensionValues param.Field[[]string] `json:"dimension_values,required"`
	// Unit price for the specified dimension_values
	UnitAmount param.Field[string] `json:"unit_amount,required"`
	// Optional multiplier to scale rated quantities by
	ScalingFactor param.Field[float64] `json:"scaling_factor"`
}

func (r PlanNewParamsPricesNewPlanMatrixPriceMatrixConfigMatrixValue) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type PlanNewParamsPricesNewPlanMatrixPriceModelType string

const (
	PlanNewParamsPricesNewPlanMatrixPriceModelTypeMatrix PlanNewParamsPricesNewPlanMatrixPriceModelType = "matrix"
)

func (r PlanNewParamsPricesNewPlanMatrixPriceModelType) IsKnown() bool {
	switch r {
	case PlanNewParamsPricesNewPlanMatrixPriceModelTypeMatrix:
		return true
	}
	return false
}

type PlanNewParamsPricesNewPlanTieredPrice struct {
	// The cadence to bill for this price on.
	Cadence param.Field[PlanNewParamsPricesNewPlanTieredPriceCadence] `json:"cadence,required"`
	// The id of the item the plan will be associated with.
	ItemID    param.Field[string]                                         `json:"item_id,required"`
	ModelType param.Field[PlanNewParamsPricesNewPlanTieredPriceModelType] `json:"model_type,required"`
	// The name of the price.
	Name         param.Field[string]                                            `json:"name,required"`
	TieredConfig param.Field[PlanNewParamsPricesNewPlanTieredPriceTieredConfig] `json:"tiered_config,required"`
	// The id of the billable metric for the price. Only needed if the price is
	// usage-based.
	BillableMetricID param.Field[string] `json:"billable_metric_id"`
	// If the Price represents a fixed cost, the price will be billed in-advance if
	// this is true, and in-arrears if this is false.
	BilledInAdvance param.Field[bool] `json:"billed_in_advance"`
	// An alias for the price.
	ExternalPriceID param.Field[string] `json:"external_price_id"`
	// If the Price represents a fixed cost, this represents the quantity of units
	// applied.
	FixedPriceQuantity param.Field[float64] `json:"fixed_price_quantity"`
	// The property used to group this price on an invoice
	InvoiceGroupingKey param.Field[string] `json:"invoice_grouping_key"`
}

func (r PlanNewParamsPricesNewPlanTieredPrice) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r PlanNewParamsPricesNewPlanTieredPrice) implementsPlanNewParamsPrice() {}

// The cadence to bill for this price on.
type PlanNewParamsPricesNewPlanTieredPriceCadence string

const (
	PlanNewParamsPricesNewPlanTieredPriceCadenceAnnual    PlanNewParamsPricesNewPlanTieredPriceCadence = "annual"
	PlanNewParamsPricesNewPlanTieredPriceCadenceMonthly   PlanNewParamsPricesNewPlanTieredPriceCadence = "monthly"
	PlanNewParamsPricesNewPlanTieredPriceCadenceQuarterly PlanNewParamsPricesNewPlanTieredPriceCadence = "quarterly"
	PlanNewParamsPricesNewPlanTieredPriceCadenceOneTime   PlanNewParamsPricesNewPlanTieredPriceCadence = "one_time"
)

func (r PlanNewParamsPricesNewPlanTieredPriceCadence) IsKnown() bool {
	switch r {
	case PlanNewParamsPricesNewPlanTieredPriceCadenceAnnual, PlanNewParamsPricesNewPlanTieredPriceCadenceMonthly, PlanNewParamsPricesNewPlanTieredPriceCadenceQuarterly, PlanNewParamsPricesNewPlanTieredPriceCadenceOneTime:
		return true
	}
	return false
}

type PlanNewParamsPricesNewPlanTieredPriceModelType string

const (
	PlanNewParamsPricesNewPlanTieredPriceModelTypeTiered PlanNewParamsPricesNewPlanTieredPriceModelType = "tiered"
)

func (r PlanNewParamsPricesNewPlanTieredPriceModelType) IsKnown() bool {
	switch r {
	case PlanNewParamsPricesNewPlanTieredPriceModelTypeTiered:
		return true
	}
	return false
}

type PlanNewParamsPricesNewPlanTieredPriceTieredConfig struct {
	// Tiers for rating based on total usage quantities into the specified tier
	Tiers param.Field[[]PlanNewParamsPricesNewPlanTieredPriceTieredConfigTier] `json:"tiers,required"`
}

func (r PlanNewParamsPricesNewPlanTieredPriceTieredConfig) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type PlanNewParamsPricesNewPlanTieredPriceTieredConfigTier struct {
	// Inclusive tier starting value
	FirstUnit param.Field[float64] `json:"first_unit,required"`
	// Amount per unit
	UnitAmount param.Field[string] `json:"unit_amount,required"`
	// Exclusive tier ending value. If null, this is treated as the last tier
	LastUnit param.Field[float64] `json:"last_unit"`
}

func (r PlanNewParamsPricesNewPlanTieredPriceTieredConfigTier) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type PlanNewParamsPricesNewPlanTieredBpsPrice struct {
	// The cadence to bill for this price on.
	Cadence param.Field[PlanNewParamsPricesNewPlanTieredBpsPriceCadence] `json:"cadence,required"`
	// The id of the item the plan will be associated with.
	ItemID    param.Field[string]                                            `json:"item_id,required"`
	ModelType param.Field[PlanNewParamsPricesNewPlanTieredBpsPriceModelType] `json:"model_type,required"`
	// The name of the price.
	Name            param.Field[string]                                                  `json:"name,required"`
	TieredBpsConfig param.Field[PlanNewParamsPricesNewPlanTieredBpsPriceTieredBpsConfig] `json:"tiered_bps_config,required"`
	// The id of the billable metric for the price. Only needed if the price is
	// usage-based.
	BillableMetricID param.Field[string] `json:"billable_metric_id"`
	// If the Price represents a fixed cost, the price will be billed in-advance if
	// this is true, and in-arrears if this is false.
	BilledInAdvance param.Field[bool] `json:"billed_in_advance"`
	// An alias for the price.
	ExternalPriceID param.Field[string] `json:"external_price_id"`
	// If the Price represents a fixed cost, this represents the quantity of units
	// applied.
	FixedPriceQuantity param.Field[float64] `json:"fixed_price_quantity"`
	// The property used to group this price on an invoice
	InvoiceGroupingKey param.Field[string] `json:"invoice_grouping_key"`
}

func (r PlanNewParamsPricesNewPlanTieredBpsPrice) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r PlanNewParamsPricesNewPlanTieredBpsPrice) implementsPlanNewParamsPrice() {}

// The cadence to bill for this price on.
type PlanNewParamsPricesNewPlanTieredBpsPriceCadence string

const (
	PlanNewParamsPricesNewPlanTieredBpsPriceCadenceAnnual    PlanNewParamsPricesNewPlanTieredBpsPriceCadence = "annual"
	PlanNewParamsPricesNewPlanTieredBpsPriceCadenceMonthly   PlanNewParamsPricesNewPlanTieredBpsPriceCadence = "monthly"
	PlanNewParamsPricesNewPlanTieredBpsPriceCadenceQuarterly PlanNewParamsPricesNewPlanTieredBpsPriceCadence = "quarterly"
	PlanNewParamsPricesNewPlanTieredBpsPriceCadenceOneTime   PlanNewParamsPricesNewPlanTieredBpsPriceCadence = "one_time"
)

func (r PlanNewParamsPricesNewPlanTieredBpsPriceCadence) IsKnown() bool {
	switch r {
	case PlanNewParamsPricesNewPlanTieredBpsPriceCadenceAnnual, PlanNewParamsPricesNewPlanTieredBpsPriceCadenceMonthly, PlanNewParamsPricesNewPlanTieredBpsPriceCadenceQuarterly, PlanNewParamsPricesNewPlanTieredBpsPriceCadenceOneTime:
		return true
	}
	return false
}

type PlanNewParamsPricesNewPlanTieredBpsPriceModelType string

const (
	PlanNewParamsPricesNewPlanTieredBpsPriceModelTypeTieredBps PlanNewParamsPricesNewPlanTieredBpsPriceModelType = "tiered_bps"
)

func (r PlanNewParamsPricesNewPlanTieredBpsPriceModelType) IsKnown() bool {
	switch r {
	case PlanNewParamsPricesNewPlanTieredBpsPriceModelTypeTieredBps:
		return true
	}
	return false
}

type PlanNewParamsPricesNewPlanTieredBpsPriceTieredBpsConfig struct {
	// Tiers for a Graduated BPS pricing model, where usage is bucketed into specified
	// tiers
	Tiers param.Field[[]PlanNewParamsPricesNewPlanTieredBpsPriceTieredBpsConfigTier] `json:"tiers,required"`
}

func (r PlanNewParamsPricesNewPlanTieredBpsPriceTieredBpsConfig) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type PlanNewParamsPricesNewPlanTieredBpsPriceTieredBpsConfigTier struct {
	// Per-event basis point rate
	Bps param.Field[float64] `json:"bps,required"`
	// Inclusive tier starting value
	MinimumAmount param.Field[string] `json:"minimum_amount,required"`
	// Exclusive tier ending value
	MaximumAmount param.Field[string] `json:"maximum_amount"`
	// Per unit maximum to charge
	PerUnitMaximum param.Field[string] `json:"per_unit_maximum"`
}

func (r PlanNewParamsPricesNewPlanTieredBpsPriceTieredBpsConfigTier) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type PlanNewParamsPricesNewPlanBpsPrice struct {
	BpsConfig param.Field[PlanNewParamsPricesNewPlanBpsPriceBpsConfig] `json:"bps_config,required"`
	// The cadence to bill for this price on.
	Cadence param.Field[PlanNewParamsPricesNewPlanBpsPriceCadence] `json:"cadence,required"`
	// The id of the item the plan will be associated with.
	ItemID    param.Field[string]                                      `json:"item_id,required"`
	ModelType param.Field[PlanNewParamsPricesNewPlanBpsPriceModelType] `json:"model_type,required"`
	// The name of the price.
	Name param.Field[string] `json:"name,required"`
	// The id of the billable metric for the price. Only needed if the price is
	// usage-based.
	BillableMetricID param.Field[string] `json:"billable_metric_id"`
	// If the Price represents a fixed cost, the price will be billed in-advance if
	// this is true, and in-arrears if this is false.
	BilledInAdvance param.Field[bool] `json:"billed_in_advance"`
	// An alias for the price.
	ExternalPriceID param.Field[string] `json:"external_price_id"`
	// If the Price represents a fixed cost, this represents the quantity of units
	// applied.
	FixedPriceQuantity param.Field[float64] `json:"fixed_price_quantity"`
	// The property used to group this price on an invoice
	InvoiceGroupingKey param.Field[string] `json:"invoice_grouping_key"`
}

func (r PlanNewParamsPricesNewPlanBpsPrice) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r PlanNewParamsPricesNewPlanBpsPrice) implementsPlanNewParamsPrice() {}

type PlanNewParamsPricesNewPlanBpsPriceBpsConfig struct {
	// Basis point take rate per event
	Bps param.Field[float64] `json:"bps,required"`
	// Optional currency amount maximum to cap spend per event
	PerUnitMaximum param.Field[string] `json:"per_unit_maximum"`
}

func (r PlanNewParamsPricesNewPlanBpsPriceBpsConfig) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// The cadence to bill for this price on.
type PlanNewParamsPricesNewPlanBpsPriceCadence string

const (
	PlanNewParamsPricesNewPlanBpsPriceCadenceAnnual    PlanNewParamsPricesNewPlanBpsPriceCadence = "annual"
	PlanNewParamsPricesNewPlanBpsPriceCadenceMonthly   PlanNewParamsPricesNewPlanBpsPriceCadence = "monthly"
	PlanNewParamsPricesNewPlanBpsPriceCadenceQuarterly PlanNewParamsPricesNewPlanBpsPriceCadence = "quarterly"
	PlanNewParamsPricesNewPlanBpsPriceCadenceOneTime   PlanNewParamsPricesNewPlanBpsPriceCadence = "one_time"
)

func (r PlanNewParamsPricesNewPlanBpsPriceCadence) IsKnown() bool {
	switch r {
	case PlanNewParamsPricesNewPlanBpsPriceCadenceAnnual, PlanNewParamsPricesNewPlanBpsPriceCadenceMonthly, PlanNewParamsPricesNewPlanBpsPriceCadenceQuarterly, PlanNewParamsPricesNewPlanBpsPriceCadenceOneTime:
		return true
	}
	return false
}

type PlanNewParamsPricesNewPlanBpsPriceModelType string

const (
	PlanNewParamsPricesNewPlanBpsPriceModelTypeBps PlanNewParamsPricesNewPlanBpsPriceModelType = "bps"
)

func (r PlanNewParamsPricesNewPlanBpsPriceModelType) IsKnown() bool {
	switch r {
	case PlanNewParamsPricesNewPlanBpsPriceModelTypeBps:
		return true
	}
	return false
}

type PlanNewParamsPricesNewPlanBulkBpsPrice struct {
	BulkBpsConfig param.Field[PlanNewParamsPricesNewPlanBulkBpsPriceBulkBpsConfig] `json:"bulk_bps_config,required"`
	// The cadence to bill for this price on.
	Cadence param.Field[PlanNewParamsPricesNewPlanBulkBpsPriceCadence] `json:"cadence,required"`
	// The id of the item the plan will be associated with.
	ItemID    param.Field[string]                                          `json:"item_id,required"`
	ModelType param.Field[PlanNewParamsPricesNewPlanBulkBpsPriceModelType] `json:"model_type,required"`
	// The name of the price.
	Name param.Field[string] `json:"name,required"`
	// The id of the billable metric for the price. Only needed if the price is
	// usage-based.
	BillableMetricID param.Field[string] `json:"billable_metric_id"`
	// If the Price represents a fixed cost, the price will be billed in-advance if
	// this is true, and in-arrears if this is false.
	BilledInAdvance param.Field[bool] `json:"billed_in_advance"`
	// An alias for the price.
	ExternalPriceID param.Field[string] `json:"external_price_id"`
	// If the Price represents a fixed cost, this represents the quantity of units
	// applied.
	FixedPriceQuantity param.Field[float64] `json:"fixed_price_quantity"`
	// The property used to group this price on an invoice
	InvoiceGroupingKey param.Field[string] `json:"invoice_grouping_key"`
}

func (r PlanNewParamsPricesNewPlanBulkBpsPrice) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r PlanNewParamsPricesNewPlanBulkBpsPrice) implementsPlanNewParamsPrice() {}

type PlanNewParamsPricesNewPlanBulkBpsPriceBulkBpsConfig struct {
	// Tiers for a bulk BPS pricing model where all usage is aggregated to a single
	// tier based on total volume
	Tiers param.Field[[]PlanNewParamsPricesNewPlanBulkBpsPriceBulkBpsConfigTier] `json:"tiers,required"`
}

func (r PlanNewParamsPricesNewPlanBulkBpsPriceBulkBpsConfig) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type PlanNewParamsPricesNewPlanBulkBpsPriceBulkBpsConfigTier struct {
	// Basis points to rate on
	Bps param.Field[float64] `json:"bps,required"`
	// Upper bound for tier
	MaximumAmount param.Field[string] `json:"maximum_amount"`
	// The maximum amount to charge for any one event
	PerUnitMaximum param.Field[string] `json:"per_unit_maximum"`
}

func (r PlanNewParamsPricesNewPlanBulkBpsPriceBulkBpsConfigTier) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// The cadence to bill for this price on.
type PlanNewParamsPricesNewPlanBulkBpsPriceCadence string

const (
	PlanNewParamsPricesNewPlanBulkBpsPriceCadenceAnnual    PlanNewParamsPricesNewPlanBulkBpsPriceCadence = "annual"
	PlanNewParamsPricesNewPlanBulkBpsPriceCadenceMonthly   PlanNewParamsPricesNewPlanBulkBpsPriceCadence = "monthly"
	PlanNewParamsPricesNewPlanBulkBpsPriceCadenceQuarterly PlanNewParamsPricesNewPlanBulkBpsPriceCadence = "quarterly"
	PlanNewParamsPricesNewPlanBulkBpsPriceCadenceOneTime   PlanNewParamsPricesNewPlanBulkBpsPriceCadence = "one_time"
)

func (r PlanNewParamsPricesNewPlanBulkBpsPriceCadence) IsKnown() bool {
	switch r {
	case PlanNewParamsPricesNewPlanBulkBpsPriceCadenceAnnual, PlanNewParamsPricesNewPlanBulkBpsPriceCadenceMonthly, PlanNewParamsPricesNewPlanBulkBpsPriceCadenceQuarterly, PlanNewParamsPricesNewPlanBulkBpsPriceCadenceOneTime:
		return true
	}
	return false
}

type PlanNewParamsPricesNewPlanBulkBpsPriceModelType string

const (
	PlanNewParamsPricesNewPlanBulkBpsPriceModelTypeBulkBps PlanNewParamsPricesNewPlanBulkBpsPriceModelType = "bulk_bps"
)

func (r PlanNewParamsPricesNewPlanBulkBpsPriceModelType) IsKnown() bool {
	switch r {
	case PlanNewParamsPricesNewPlanBulkBpsPriceModelTypeBulkBps:
		return true
	}
	return false
}

type PlanNewParamsPricesNewPlanBulkPrice struct {
	BulkConfig param.Field[PlanNewParamsPricesNewPlanBulkPriceBulkConfig] `json:"bulk_config,required"`
	// The cadence to bill for this price on.
	Cadence param.Field[PlanNewParamsPricesNewPlanBulkPriceCadence] `json:"cadence,required"`
	// The id of the item the plan will be associated with.
	ItemID    param.Field[string]                                       `json:"item_id,required"`
	ModelType param.Field[PlanNewParamsPricesNewPlanBulkPriceModelType] `json:"model_type,required"`
	// The name of the price.
	Name param.Field[string] `json:"name,required"`
	// The id of the billable metric for the price. Only needed if the price is
	// usage-based.
	BillableMetricID param.Field[string] `json:"billable_metric_id"`
	// If the Price represents a fixed cost, the price will be billed in-advance if
	// this is true, and in-arrears if this is false.
	BilledInAdvance param.Field[bool] `json:"billed_in_advance"`
	// An alias for the price.
	ExternalPriceID param.Field[string] `json:"external_price_id"`
	// If the Price represents a fixed cost, this represents the quantity of units
	// applied.
	FixedPriceQuantity param.Field[float64] `json:"fixed_price_quantity"`
	// The property used to group this price on an invoice
	InvoiceGroupingKey param.Field[string] `json:"invoice_grouping_key"`
}

func (r PlanNewParamsPricesNewPlanBulkPrice) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r PlanNewParamsPricesNewPlanBulkPrice) implementsPlanNewParamsPrice() {}

type PlanNewParamsPricesNewPlanBulkPriceBulkConfig struct {
	// Bulk tiers for rating based on total usage volume
	Tiers param.Field[[]PlanNewParamsPricesNewPlanBulkPriceBulkConfigTier] `json:"tiers,required"`
}

func (r PlanNewParamsPricesNewPlanBulkPriceBulkConfig) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type PlanNewParamsPricesNewPlanBulkPriceBulkConfigTier struct {
	// Amount per unit
	UnitAmount param.Field[string] `json:"unit_amount,required"`
	// Upper bound for this tier
	MaximumUnits param.Field[float64] `json:"maximum_units"`
}

func (r PlanNewParamsPricesNewPlanBulkPriceBulkConfigTier) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// The cadence to bill for this price on.
type PlanNewParamsPricesNewPlanBulkPriceCadence string

const (
	PlanNewParamsPricesNewPlanBulkPriceCadenceAnnual    PlanNewParamsPricesNewPlanBulkPriceCadence = "annual"
	PlanNewParamsPricesNewPlanBulkPriceCadenceMonthly   PlanNewParamsPricesNewPlanBulkPriceCadence = "monthly"
	PlanNewParamsPricesNewPlanBulkPriceCadenceQuarterly PlanNewParamsPricesNewPlanBulkPriceCadence = "quarterly"
	PlanNewParamsPricesNewPlanBulkPriceCadenceOneTime   PlanNewParamsPricesNewPlanBulkPriceCadence = "one_time"
)

func (r PlanNewParamsPricesNewPlanBulkPriceCadence) IsKnown() bool {
	switch r {
	case PlanNewParamsPricesNewPlanBulkPriceCadenceAnnual, PlanNewParamsPricesNewPlanBulkPriceCadenceMonthly, PlanNewParamsPricesNewPlanBulkPriceCadenceQuarterly, PlanNewParamsPricesNewPlanBulkPriceCadenceOneTime:
		return true
	}
	return false
}

type PlanNewParamsPricesNewPlanBulkPriceModelType string

const (
	PlanNewParamsPricesNewPlanBulkPriceModelTypeBulk PlanNewParamsPricesNewPlanBulkPriceModelType = "bulk"
)

func (r PlanNewParamsPricesNewPlanBulkPriceModelType) IsKnown() bool {
	switch r {
	case PlanNewParamsPricesNewPlanBulkPriceModelTypeBulk:
		return true
	}
	return false
}

type PlanNewParamsPricesNewPlanThresholdTotalAmountPrice struct {
	// The cadence to bill for this price on.
	Cadence param.Field[PlanNewParamsPricesNewPlanThresholdTotalAmountPriceCadence] `json:"cadence,required"`
	// The id of the item the plan will be associated with.
	ItemID    param.Field[string]                                                       `json:"item_id,required"`
	ModelType param.Field[PlanNewParamsPricesNewPlanThresholdTotalAmountPriceModelType] `json:"model_type,required"`
	// The name of the price.
	Name                       param.Field[string]                 `json:"name,required"`
	ThresholdTotalAmountConfig param.Field[map[string]interface{}] `json:"threshold_total_amount_config,required"`
	// The id of the billable metric for the price. Only needed if the price is
	// usage-based.
	BillableMetricID param.Field[string] `json:"billable_metric_id"`
	// If the Price represents a fixed cost, the price will be billed in-advance if
	// this is true, and in-arrears if this is false.
	BilledInAdvance param.Field[bool] `json:"billed_in_advance"`
	// An alias for the price.
	ExternalPriceID param.Field[string] `json:"external_price_id"`
	// If the Price represents a fixed cost, this represents the quantity of units
	// applied.
	FixedPriceQuantity param.Field[float64] `json:"fixed_price_quantity"`
	// The property used to group this price on an invoice
	InvoiceGroupingKey param.Field[string] `json:"invoice_grouping_key"`
}

func (r PlanNewParamsPricesNewPlanThresholdTotalAmountPrice) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r PlanNewParamsPricesNewPlanThresholdTotalAmountPrice) implementsPlanNewParamsPrice() {}

// The cadence to bill for this price on.
type PlanNewParamsPricesNewPlanThresholdTotalAmountPriceCadence string

const (
	PlanNewParamsPricesNewPlanThresholdTotalAmountPriceCadenceAnnual    PlanNewParamsPricesNewPlanThresholdTotalAmountPriceCadence = "annual"
	PlanNewParamsPricesNewPlanThresholdTotalAmountPriceCadenceMonthly   PlanNewParamsPricesNewPlanThresholdTotalAmountPriceCadence = "monthly"
	PlanNewParamsPricesNewPlanThresholdTotalAmountPriceCadenceQuarterly PlanNewParamsPricesNewPlanThresholdTotalAmountPriceCadence = "quarterly"
	PlanNewParamsPricesNewPlanThresholdTotalAmountPriceCadenceOneTime   PlanNewParamsPricesNewPlanThresholdTotalAmountPriceCadence = "one_time"
)

func (r PlanNewParamsPricesNewPlanThresholdTotalAmountPriceCadence) IsKnown() bool {
	switch r {
	case PlanNewParamsPricesNewPlanThresholdTotalAmountPriceCadenceAnnual, PlanNewParamsPricesNewPlanThresholdTotalAmountPriceCadenceMonthly, PlanNewParamsPricesNewPlanThresholdTotalAmountPriceCadenceQuarterly, PlanNewParamsPricesNewPlanThresholdTotalAmountPriceCadenceOneTime:
		return true
	}
	return false
}

type PlanNewParamsPricesNewPlanThresholdTotalAmountPriceModelType string

const (
	PlanNewParamsPricesNewPlanThresholdTotalAmountPriceModelTypeThresholdTotalAmount PlanNewParamsPricesNewPlanThresholdTotalAmountPriceModelType = "threshold_total_amount"
)

func (r PlanNewParamsPricesNewPlanThresholdTotalAmountPriceModelType) IsKnown() bool {
	switch r {
	case PlanNewParamsPricesNewPlanThresholdTotalAmountPriceModelTypeThresholdTotalAmount:
		return true
	}
	return false
}

type PlanNewParamsPricesNewPlanTieredPackagePrice struct {
	// The cadence to bill for this price on.
	Cadence param.Field[PlanNewParamsPricesNewPlanTieredPackagePriceCadence] `json:"cadence,required"`
	// The id of the item the plan will be associated with.
	ItemID    param.Field[string]                                                `json:"item_id,required"`
	ModelType param.Field[PlanNewParamsPricesNewPlanTieredPackagePriceModelType] `json:"model_type,required"`
	// The name of the price.
	Name                param.Field[string]                 `json:"name,required"`
	TieredPackageConfig param.Field[map[string]interface{}] `json:"tiered_package_config,required"`
	// The id of the billable metric for the price. Only needed if the price is
	// usage-based.
	BillableMetricID param.Field[string] `json:"billable_metric_id"`
	// If the Price represents a fixed cost, the price will be billed in-advance if
	// this is true, and in-arrears if this is false.
	BilledInAdvance param.Field[bool] `json:"billed_in_advance"`
	// An alias for the price.
	ExternalPriceID param.Field[string] `json:"external_price_id"`
	// If the Price represents a fixed cost, this represents the quantity of units
	// applied.
	FixedPriceQuantity param.Field[float64] `json:"fixed_price_quantity"`
	// The property used to group this price on an invoice
	InvoiceGroupingKey param.Field[string] `json:"invoice_grouping_key"`
}

func (r PlanNewParamsPricesNewPlanTieredPackagePrice) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r PlanNewParamsPricesNewPlanTieredPackagePrice) implementsPlanNewParamsPrice() {}

// The cadence to bill for this price on.
type PlanNewParamsPricesNewPlanTieredPackagePriceCadence string

const (
	PlanNewParamsPricesNewPlanTieredPackagePriceCadenceAnnual    PlanNewParamsPricesNewPlanTieredPackagePriceCadence = "annual"
	PlanNewParamsPricesNewPlanTieredPackagePriceCadenceMonthly   PlanNewParamsPricesNewPlanTieredPackagePriceCadence = "monthly"
	PlanNewParamsPricesNewPlanTieredPackagePriceCadenceQuarterly PlanNewParamsPricesNewPlanTieredPackagePriceCadence = "quarterly"
	PlanNewParamsPricesNewPlanTieredPackagePriceCadenceOneTime   PlanNewParamsPricesNewPlanTieredPackagePriceCadence = "one_time"
)

func (r PlanNewParamsPricesNewPlanTieredPackagePriceCadence) IsKnown() bool {
	switch r {
	case PlanNewParamsPricesNewPlanTieredPackagePriceCadenceAnnual, PlanNewParamsPricesNewPlanTieredPackagePriceCadenceMonthly, PlanNewParamsPricesNewPlanTieredPackagePriceCadenceQuarterly, PlanNewParamsPricesNewPlanTieredPackagePriceCadenceOneTime:
		return true
	}
	return false
}

type PlanNewParamsPricesNewPlanTieredPackagePriceModelType string

const (
	PlanNewParamsPricesNewPlanTieredPackagePriceModelTypeTieredPackage PlanNewParamsPricesNewPlanTieredPackagePriceModelType = "tiered_package"
)

func (r PlanNewParamsPricesNewPlanTieredPackagePriceModelType) IsKnown() bool {
	switch r {
	case PlanNewParamsPricesNewPlanTieredPackagePriceModelTypeTieredPackage:
		return true
	}
	return false
}

type PlanNewParamsPricesNewPlanTieredWithMinimumPrice struct {
	// The cadence to bill for this price on.
	Cadence param.Field[PlanNewParamsPricesNewPlanTieredWithMinimumPriceCadence] `json:"cadence,required"`
	// The id of the item the plan will be associated with.
	ItemID    param.Field[string]                                                    `json:"item_id,required"`
	ModelType param.Field[PlanNewParamsPricesNewPlanTieredWithMinimumPriceModelType] `json:"model_type,required"`
	// The name of the price.
	Name                    param.Field[string]                 `json:"name,required"`
	TieredWithMinimumConfig param.Field[map[string]interface{}] `json:"tiered_with_minimum_config,required"`
	// The id of the billable metric for the price. Only needed if the price is
	// usage-based.
	BillableMetricID param.Field[string] `json:"billable_metric_id"`
	// If the Price represents a fixed cost, the price will be billed in-advance if
	// this is true, and in-arrears if this is false.
	BilledInAdvance param.Field[bool] `json:"billed_in_advance"`
	// An alias for the price.
	ExternalPriceID param.Field[string] `json:"external_price_id"`
	// If the Price represents a fixed cost, this represents the quantity of units
	// applied.
	FixedPriceQuantity param.Field[float64] `json:"fixed_price_quantity"`
	// The property used to group this price on an invoice
	InvoiceGroupingKey param.Field[string] `json:"invoice_grouping_key"`
}

func (r PlanNewParamsPricesNewPlanTieredWithMinimumPrice) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r PlanNewParamsPricesNewPlanTieredWithMinimumPrice) implementsPlanNewParamsPrice() {}

// The cadence to bill for this price on.
type PlanNewParamsPricesNewPlanTieredWithMinimumPriceCadence string

const (
	PlanNewParamsPricesNewPlanTieredWithMinimumPriceCadenceAnnual    PlanNewParamsPricesNewPlanTieredWithMinimumPriceCadence = "annual"
	PlanNewParamsPricesNewPlanTieredWithMinimumPriceCadenceMonthly   PlanNewParamsPricesNewPlanTieredWithMinimumPriceCadence = "monthly"
	PlanNewParamsPricesNewPlanTieredWithMinimumPriceCadenceQuarterly PlanNewParamsPricesNewPlanTieredWithMinimumPriceCadence = "quarterly"
	PlanNewParamsPricesNewPlanTieredWithMinimumPriceCadenceOneTime   PlanNewParamsPricesNewPlanTieredWithMinimumPriceCadence = "one_time"
)

func (r PlanNewParamsPricesNewPlanTieredWithMinimumPriceCadence) IsKnown() bool {
	switch r {
	case PlanNewParamsPricesNewPlanTieredWithMinimumPriceCadenceAnnual, PlanNewParamsPricesNewPlanTieredWithMinimumPriceCadenceMonthly, PlanNewParamsPricesNewPlanTieredWithMinimumPriceCadenceQuarterly, PlanNewParamsPricesNewPlanTieredWithMinimumPriceCadenceOneTime:
		return true
	}
	return false
}

type PlanNewParamsPricesNewPlanTieredWithMinimumPriceModelType string

const (
	PlanNewParamsPricesNewPlanTieredWithMinimumPriceModelTypeTieredWithMinimum PlanNewParamsPricesNewPlanTieredWithMinimumPriceModelType = "tiered_with_minimum"
)

func (r PlanNewParamsPricesNewPlanTieredWithMinimumPriceModelType) IsKnown() bool {
	switch r {
	case PlanNewParamsPricesNewPlanTieredWithMinimumPriceModelTypeTieredWithMinimum:
		return true
	}
	return false
}

type PlanNewParamsPricesNewPlanUnitWithPercentPrice struct {
	// The cadence to bill for this price on.
	Cadence param.Field[PlanNewParamsPricesNewPlanUnitWithPercentPriceCadence] `json:"cadence,required"`
	// The id of the item the plan will be associated with.
	ItemID    param.Field[string]                                                  `json:"item_id,required"`
	ModelType param.Field[PlanNewParamsPricesNewPlanUnitWithPercentPriceModelType] `json:"model_type,required"`
	// The name of the price.
	Name                  param.Field[string]                 `json:"name,required"`
	UnitWithPercentConfig param.Field[map[string]interface{}] `json:"unit_with_percent_config,required"`
	// The id of the billable metric for the price. Only needed if the price is
	// usage-based.
	BillableMetricID param.Field[string] `json:"billable_metric_id"`
	// If the Price represents a fixed cost, the price will be billed in-advance if
	// this is true, and in-arrears if this is false.
	BilledInAdvance param.Field[bool] `json:"billed_in_advance"`
	// An alias for the price.
	ExternalPriceID param.Field[string] `json:"external_price_id"`
	// If the Price represents a fixed cost, this represents the quantity of units
	// applied.
	FixedPriceQuantity param.Field[float64] `json:"fixed_price_quantity"`
	// The property used to group this price on an invoice
	InvoiceGroupingKey param.Field[string] `json:"invoice_grouping_key"`
}

func (r PlanNewParamsPricesNewPlanUnitWithPercentPrice) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r PlanNewParamsPricesNewPlanUnitWithPercentPrice) implementsPlanNewParamsPrice() {}

// The cadence to bill for this price on.
type PlanNewParamsPricesNewPlanUnitWithPercentPriceCadence string

const (
	PlanNewParamsPricesNewPlanUnitWithPercentPriceCadenceAnnual    PlanNewParamsPricesNewPlanUnitWithPercentPriceCadence = "annual"
	PlanNewParamsPricesNewPlanUnitWithPercentPriceCadenceMonthly   PlanNewParamsPricesNewPlanUnitWithPercentPriceCadence = "monthly"
	PlanNewParamsPricesNewPlanUnitWithPercentPriceCadenceQuarterly PlanNewParamsPricesNewPlanUnitWithPercentPriceCadence = "quarterly"
	PlanNewParamsPricesNewPlanUnitWithPercentPriceCadenceOneTime   PlanNewParamsPricesNewPlanUnitWithPercentPriceCadence = "one_time"
)

func (r PlanNewParamsPricesNewPlanUnitWithPercentPriceCadence) IsKnown() bool {
	switch r {
	case PlanNewParamsPricesNewPlanUnitWithPercentPriceCadenceAnnual, PlanNewParamsPricesNewPlanUnitWithPercentPriceCadenceMonthly, PlanNewParamsPricesNewPlanUnitWithPercentPriceCadenceQuarterly, PlanNewParamsPricesNewPlanUnitWithPercentPriceCadenceOneTime:
		return true
	}
	return false
}

type PlanNewParamsPricesNewPlanUnitWithPercentPriceModelType string

const (
	PlanNewParamsPricesNewPlanUnitWithPercentPriceModelTypeUnitWithPercent PlanNewParamsPricesNewPlanUnitWithPercentPriceModelType = "unit_with_percent"
)

func (r PlanNewParamsPricesNewPlanUnitWithPercentPriceModelType) IsKnown() bool {
	switch r {
	case PlanNewParamsPricesNewPlanUnitWithPercentPriceModelTypeUnitWithPercent:
		return true
	}
	return false
}

type PlanNewParamsPricesNewPlanPackageWithAllocationPrice struct {
	// The cadence to bill for this price on.
	Cadence param.Field[PlanNewParamsPricesNewPlanPackageWithAllocationPriceCadence] `json:"cadence,required"`
	// The id of the item the plan will be associated with.
	ItemID    param.Field[string]                                                        `json:"item_id,required"`
	ModelType param.Field[PlanNewParamsPricesNewPlanPackageWithAllocationPriceModelType] `json:"model_type,required"`
	// The name of the price.
	Name                        param.Field[string]                 `json:"name,required"`
	PackageWithAllocationConfig param.Field[map[string]interface{}] `json:"package_with_allocation_config,required"`
	// The id of the billable metric for the price. Only needed if the price is
	// usage-based.
	BillableMetricID param.Field[string] `json:"billable_metric_id"`
	// If the Price represents a fixed cost, the price will be billed in-advance if
	// this is true, and in-arrears if this is false.
	BilledInAdvance param.Field[bool] `json:"billed_in_advance"`
	// An alias for the price.
	ExternalPriceID param.Field[string] `json:"external_price_id"`
	// If the Price represents a fixed cost, this represents the quantity of units
	// applied.
	FixedPriceQuantity param.Field[float64] `json:"fixed_price_quantity"`
	// The property used to group this price on an invoice
	InvoiceGroupingKey param.Field[string] `json:"invoice_grouping_key"`
}

func (r PlanNewParamsPricesNewPlanPackageWithAllocationPrice) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r PlanNewParamsPricesNewPlanPackageWithAllocationPrice) implementsPlanNewParamsPrice() {}

// The cadence to bill for this price on.
type PlanNewParamsPricesNewPlanPackageWithAllocationPriceCadence string

const (
	PlanNewParamsPricesNewPlanPackageWithAllocationPriceCadenceAnnual    PlanNewParamsPricesNewPlanPackageWithAllocationPriceCadence = "annual"
	PlanNewParamsPricesNewPlanPackageWithAllocationPriceCadenceMonthly   PlanNewParamsPricesNewPlanPackageWithAllocationPriceCadence = "monthly"
	PlanNewParamsPricesNewPlanPackageWithAllocationPriceCadenceQuarterly PlanNewParamsPricesNewPlanPackageWithAllocationPriceCadence = "quarterly"
	PlanNewParamsPricesNewPlanPackageWithAllocationPriceCadenceOneTime   PlanNewParamsPricesNewPlanPackageWithAllocationPriceCadence = "one_time"
)

func (r PlanNewParamsPricesNewPlanPackageWithAllocationPriceCadence) IsKnown() bool {
	switch r {
	case PlanNewParamsPricesNewPlanPackageWithAllocationPriceCadenceAnnual, PlanNewParamsPricesNewPlanPackageWithAllocationPriceCadenceMonthly, PlanNewParamsPricesNewPlanPackageWithAllocationPriceCadenceQuarterly, PlanNewParamsPricesNewPlanPackageWithAllocationPriceCadenceOneTime:
		return true
	}
	return false
}

type PlanNewParamsPricesNewPlanPackageWithAllocationPriceModelType string

const (
	PlanNewParamsPricesNewPlanPackageWithAllocationPriceModelTypePackageWithAllocation PlanNewParamsPricesNewPlanPackageWithAllocationPriceModelType = "package_with_allocation"
)

func (r PlanNewParamsPricesNewPlanPackageWithAllocationPriceModelType) IsKnown() bool {
	switch r {
	case PlanNewParamsPricesNewPlanPackageWithAllocationPriceModelTypePackageWithAllocation:
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
