// File generated from our OpenAPI spec by Stainless.

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
	Discount           InvoiceDiscount `json:"discount,required,nullable"`
	// An optional user-defined ID for this plan resource, used throughout the system
	// as an alias for this Plan. Use this field to identify a plan by an existing
	// identifier in your system.
	ExternalPlanID string `json:"external_plan_id,required,nullable"`
	// An ISO 4217 currency string for which this plan is billed in. Matches `currency`
	// unless `currency` is a custom pricing unit.
	InvoicingCurrency string      `json:"invoicing_currency,required"`
	Maximum           PlanMaximum `json:"maximum,required,nullable"`
	MaximumAmount     string      `json:"maximum_amount,required,nullable"`
	Metadata          interface{} `json:"metadata,required"`
	Minimum           PlanMinimum `json:"minimum,required,nullable"`
	MinimumAmount     string      `json:"minimum_amount,required,nullable"`
	Name              string      `json:"name,required"`
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
	JSON        planJSON
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
	raw                string
	ExtraFields        map[string]apijson.Field
}

func (r *Plan) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type PlanBasePlan struct {
	ID string `json:"id,required,nullable"`
	// An optional user-defined ID for this plan resource, used throughout the system
	// as an alias for this Plan. Use this field to identify a plan by an existing
	// identifier in your system.
	ExternalPlanID string `json:"external_plan_id,required,nullable"`
	Name           string `json:"name,required,nullable"`
	JSON           planBasePlanJSON
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

type PlanMaximum struct {
	// List of price_ids that this maximum amount applies to. For plan/plan phase
	// maximums, this can be a subset of prices.
	AppliesToPriceIDs []string `json:"applies_to_price_ids,required"`
	// Maximum amount applied
	MaximumAmount string `json:"maximum_amount,required"`
	JSON          planMaximumJSON
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

type PlanMinimum struct {
	// List of price_ids that this minimum amount applies to. For plan/plan phase
	// minimums, this can be a subset of prices.
	AppliesToPriceIDs []string `json:"applies_to_price_ids,required"`
	// Minimum amount applied
	MinimumAmount string `json:"minimum_amount,required"`
	JSON          planMinimumJSON
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

type PlanPlanPhase struct {
	ID          string          `json:"id,required"`
	Description string          `json:"description,required,nullable"`
	Discount    InvoiceDiscount `json:"discount,required,nullable"`
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
	Order int64 `json:"order,required"`
	JSON  planPlanPhaseJSON
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

type PlanPlanPhasesDurationUnit string

const (
	PlanPlanPhasesDurationUnitDaily     PlanPlanPhasesDurationUnit = "daily"
	PlanPlanPhasesDurationUnitMonthly   PlanPlanPhasesDurationUnit = "monthly"
	PlanPlanPhasesDurationUnitQuarterly PlanPlanPhasesDurationUnit = "quarterly"
	PlanPlanPhasesDurationUnitAnnual    PlanPlanPhasesDurationUnit = "annual"
)

type PlanPlanPhasesMaximum struct {
	// List of price_ids that this maximum amount applies to. For plan/plan phase
	// maximums, this can be a subset of prices.
	AppliesToPriceIDs []string `json:"applies_to_price_ids,required"`
	// Maximum amount applied
	MaximumAmount string `json:"maximum_amount,required"`
	JSON          planPlanPhasesMaximumJSON
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

type PlanPlanPhasesMinimum struct {
	// List of price_ids that this minimum amount applies to. For plan/plan phase
	// minimums, this can be a subset of prices.
	AppliesToPriceIDs []string `json:"applies_to_price_ids,required"`
	// Minimum amount applied
	MinimumAmount string `json:"minimum_amount,required"`
	JSON          planPlanPhasesMinimumJSON
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

type PlanProduct struct {
	ID        string    `json:"id,required"`
	CreatedAt time.Time `json:"created_at,required" format:"date-time"`
	Name      string    `json:"name,required"`
	JSON      planProductJSON
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

type PlanStatus string

const (
	PlanStatusActive   PlanStatus = "active"
	PlanStatusArchived PlanStatus = "archived"
	PlanStatusDraft    PlanStatus = "draft"
)

type PlanTrialConfig struct {
	TrialPeriod     int64                          `json:"trial_period,required,nullable"`
	TrialPeriodUnit PlanTrialConfigTrialPeriodUnit `json:"trial_period_unit,required"`
	JSON            planTrialConfigJSON
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

type PlanTrialConfigTrialPeriodUnit string

const (
	PlanTrialConfigTrialPeriodUnitDays PlanTrialConfigTrialPeriodUnit = "days"
)

type PlanNewParams struct {
	// An ISO 4217 currency string or custom pricing unit (`credits`) for this plan's
	// prices.
	Currency param.Field[string] `json:"currency,required"`
	Name     param.Field[string] `json:"name,required"`
	// Prices for this plan. If the plan has phases, this includes prices across all
	// phases of the plan.
	Prices param.Field[[]interface{}] `json:"prices,required"`
	// Free-form text which is available on the invoice PDF and the Orb invoice portal.
	DefaultInvoiceMemo param.Field[string]      `json:"default_invoice_memo"`
	ExternalPlanID     param.Field[string]      `json:"external_plan_id"`
	Metadata           param.Field[interface{}] `json:"metadata"`
	// The net terms determines the difference between the invoice date and the issue
	// date for the invoice. If you intend the invoice to be due on issue, set this
	// to 0.
	NetTerms param.Field[int64] `json:"net_terms"`
}

func (r PlanNewParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type PlanUpdateParams struct {
	// An optional user-defined ID for this plan resource, used throughout the system
	// as an alias for this Plan. Use this field to identify a plan by an existing
	// identifier in your system.
	ExternalPlanID param.Field[string]      `json:"external_plan_id"`
	Metadata       param.Field[interface{}] `json:"metadata"`
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
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
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
