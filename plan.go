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
	Draft *PlanDraftService
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
	r.Draft = NewPlanDraftService(opts...)
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

// This endpoint can be used to update the `name`, `external_plan_id`,
// `description`, and `metadata` of an existing plan. A plan's `name` cannot be
// empty.
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
// [`pagination_metadata`](/api-reference/pagination) which lets the caller
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
// [`pagination_metadata`](/api-reference/pagination) which lets the caller
// retrieve the next page of results if they exist.
func (r *PlanService) ListAutoPaging(ctx context.Context, query PlanListParams, opts ...option.RequestOption) *pagination.PageAutoPager[Plan] {
	return pagination.NewPageAutoPager(r.List(ctx, query, opts...))
}

// This endpoint allows the creation of a new plan version for an existing plan.
func (r *PlanService) NewPlanVersion(ctx context.Context, planID string, body PlanNewPlanVersionParams, opts ...option.RequestOption) (res *PlanVersion, err error) {
	opts = slices.Concat(r.Options, opts)
	if planID == "" {
		err = errors.New("missing required plan_id parameter")
		return nil, err
	}
	path := fmt.Sprintf("plans/%s/versions", planID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return res, err
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

// This endpoint is used to fetch a plan version. It returns the phases, prices,
// and adjustments present on this version of the plan.
func (r *PlanService) FetchPlanVersion(ctx context.Context, planID string, version string, opts ...option.RequestOption) (res *PlanVersion, err error) {
	opts = slices.Concat(r.Options, opts)
	if planID == "" {
		err = errors.New("missing required plan_id parameter")
		return nil, err
	}
	if version == "" {
		err = errors.New("missing required version parameter")
		return nil, err
	}
	path := fmt.Sprintf("plans/%s/versions/%s", planID, version)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return res, err
}

// This endpoint allows setting the default version of a plan.
func (r *PlanService) SetDefaultPlanVersion(ctx context.Context, planID string, body PlanSetDefaultPlanVersionParams, opts ...option.RequestOption) (res *Plan, err error) {
	opts = slices.Concat(r.Options, opts)
	if planID == "" {
		err = errors.New("missing required plan_id parameter")
		return nil, err
	}
	path := fmt.Sprintf("plans/%s/set_default_version", planID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
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
	// Whether the plan has an unpublished draft version.
	HasDraft bool `json:"has_draft" api:"required"`
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
	HasDraft           apijson.Field
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
	// [[]PlanAdjustmentsPlanPhaseTieredPercentageDiscountAdjustmentFilter],
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
	// This field can have the runtime type of
	// [[]PlanAdjustmentsPlanPhaseTieredPercentageDiscountAdjustmentTier].
	Tiers interface{} `json:"tiers"`
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
	Tiers                apijson.Field
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
// [PlanAdjustmentsPlanPhaseTieredPercentageDiscountAdjustment],
// [shared.PlanPhaseMinimumAdjustment], [shared.PlanPhaseMaximumAdjustment].
func (r PlanAdjustment) AsUnion() PlanAdjustmentsUnion {
	return r.union
}

// Union satisfied by [shared.PlanPhaseUsageDiscountAdjustment],
// [shared.PlanPhaseAmountDiscountAdjustment],
// [shared.PlanPhasePercentageDiscountAdjustment],
// [PlanAdjustmentsPlanPhaseTieredPercentageDiscountAdjustment],
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
			Type:               reflect.TypeOf(PlanAdjustmentsPlanPhaseTieredPercentageDiscountAdjustment{}),
			DiscriminatorValue: "tiered_percentage_discount",
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

type PlanAdjustmentsPlanPhaseTieredPercentageDiscountAdjustment struct {
	ID             string                                                                   `json:"id" api:"required"`
	AdjustmentType PlanAdjustmentsPlanPhaseTieredPercentageDiscountAdjustmentAdjustmentType `json:"adjustment_type" api:"required"`
	// The price IDs that this adjustment applies to.
	//
	// Deprecated: deprecated
	AppliesToPriceIDs []string `json:"applies_to_price_ids" api:"required"`
	// The filters that determine which prices to apply this adjustment to.
	Filters []PlanAdjustmentsPlanPhaseTieredPercentageDiscountAdjustmentFilter `json:"filters" api:"required"`
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
	// The ordered, contiguous bands of cumulative eligible spend, each discounted at
	// its own percentage (progressive fill-a-tier), applied to the prices this
	// adjustment covers in a given billing period.
	Tiers []PlanAdjustmentsPlanPhaseTieredPercentageDiscountAdjustmentTier `json:"tiers" api:"required"`
	JSON  planAdjustmentsPlanPhaseTieredPercentageDiscountAdjustmentJSON   `json:"-"`
}

// planAdjustmentsPlanPhaseTieredPercentageDiscountAdjustmentJSON contains the JSON
// metadata for the struct
// [PlanAdjustmentsPlanPhaseTieredPercentageDiscountAdjustment]
type planAdjustmentsPlanPhaseTieredPercentageDiscountAdjustmentJSON struct {
	ID                   apijson.Field
	AdjustmentType       apijson.Field
	AppliesToPriceIDs    apijson.Field
	Filters              apijson.Field
	IsInvoiceLevel       apijson.Field
	PlanPhaseOrder       apijson.Field
	Reason               apijson.Field
	ReplacesAdjustmentID apijson.Field
	Tiers                apijson.Field
	raw                  string
	ExtraFields          map[string]apijson.Field
}

func (r *PlanAdjustmentsPlanPhaseTieredPercentageDiscountAdjustment) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r planAdjustmentsPlanPhaseTieredPercentageDiscountAdjustmentJSON) RawJSON() string {
	return r.raw
}

func (r PlanAdjustmentsPlanPhaseTieredPercentageDiscountAdjustment) ImplementsPlanAdjustment() {}

type PlanAdjustmentsPlanPhaseTieredPercentageDiscountAdjustmentAdjustmentType string

const (
	PlanAdjustmentsPlanPhaseTieredPercentageDiscountAdjustmentAdjustmentTypeTieredPercentageDiscount PlanAdjustmentsPlanPhaseTieredPercentageDiscountAdjustmentAdjustmentType = "tiered_percentage_discount"
)

func (r PlanAdjustmentsPlanPhaseTieredPercentageDiscountAdjustmentAdjustmentType) IsKnown() bool {
	switch r {
	case PlanAdjustmentsPlanPhaseTieredPercentageDiscountAdjustmentAdjustmentTypeTieredPercentageDiscount:
		return true
	}
	return false
}

type PlanAdjustmentsPlanPhaseTieredPercentageDiscountAdjustmentFilter struct {
	// The property of the price to filter on.
	Field PlanAdjustmentsPlanPhaseTieredPercentageDiscountAdjustmentFiltersField `json:"field" api:"required"`
	// Should prices that match the filter be included or excluded.
	Operator PlanAdjustmentsPlanPhaseTieredPercentageDiscountAdjustmentFiltersOperator `json:"operator" api:"required"`
	// The IDs or values that match this filter.
	Values []string                                                             `json:"values" api:"required"`
	JSON   planAdjustmentsPlanPhaseTieredPercentageDiscountAdjustmentFilterJSON `json:"-"`
}

// planAdjustmentsPlanPhaseTieredPercentageDiscountAdjustmentFilterJSON contains
// the JSON metadata for the struct
// [PlanAdjustmentsPlanPhaseTieredPercentageDiscountAdjustmentFilter]
type planAdjustmentsPlanPhaseTieredPercentageDiscountAdjustmentFilterJSON struct {
	Field       apijson.Field
	Operator    apijson.Field
	Values      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *PlanAdjustmentsPlanPhaseTieredPercentageDiscountAdjustmentFilter) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r planAdjustmentsPlanPhaseTieredPercentageDiscountAdjustmentFilterJSON) RawJSON() string {
	return r.raw
}

// The property of the price to filter on.
type PlanAdjustmentsPlanPhaseTieredPercentageDiscountAdjustmentFiltersField string

const (
	PlanAdjustmentsPlanPhaseTieredPercentageDiscountAdjustmentFiltersFieldPriceID       PlanAdjustmentsPlanPhaseTieredPercentageDiscountAdjustmentFiltersField = "price_id"
	PlanAdjustmentsPlanPhaseTieredPercentageDiscountAdjustmentFiltersFieldItemID        PlanAdjustmentsPlanPhaseTieredPercentageDiscountAdjustmentFiltersField = "item_id"
	PlanAdjustmentsPlanPhaseTieredPercentageDiscountAdjustmentFiltersFieldPriceType     PlanAdjustmentsPlanPhaseTieredPercentageDiscountAdjustmentFiltersField = "price_type"
	PlanAdjustmentsPlanPhaseTieredPercentageDiscountAdjustmentFiltersFieldCurrency      PlanAdjustmentsPlanPhaseTieredPercentageDiscountAdjustmentFiltersField = "currency"
	PlanAdjustmentsPlanPhaseTieredPercentageDiscountAdjustmentFiltersFieldPricingUnitID PlanAdjustmentsPlanPhaseTieredPercentageDiscountAdjustmentFiltersField = "pricing_unit_id"
)

func (r PlanAdjustmentsPlanPhaseTieredPercentageDiscountAdjustmentFiltersField) IsKnown() bool {
	switch r {
	case PlanAdjustmentsPlanPhaseTieredPercentageDiscountAdjustmentFiltersFieldPriceID, PlanAdjustmentsPlanPhaseTieredPercentageDiscountAdjustmentFiltersFieldItemID, PlanAdjustmentsPlanPhaseTieredPercentageDiscountAdjustmentFiltersFieldPriceType, PlanAdjustmentsPlanPhaseTieredPercentageDiscountAdjustmentFiltersFieldCurrency, PlanAdjustmentsPlanPhaseTieredPercentageDiscountAdjustmentFiltersFieldPricingUnitID:
		return true
	}
	return false
}

// Should prices that match the filter be included or excluded.
type PlanAdjustmentsPlanPhaseTieredPercentageDiscountAdjustmentFiltersOperator string

const (
	PlanAdjustmentsPlanPhaseTieredPercentageDiscountAdjustmentFiltersOperatorIncludes PlanAdjustmentsPlanPhaseTieredPercentageDiscountAdjustmentFiltersOperator = "includes"
	PlanAdjustmentsPlanPhaseTieredPercentageDiscountAdjustmentFiltersOperatorExcludes PlanAdjustmentsPlanPhaseTieredPercentageDiscountAdjustmentFiltersOperator = "excludes"
)

func (r PlanAdjustmentsPlanPhaseTieredPercentageDiscountAdjustmentFiltersOperator) IsKnown() bool {
	switch r {
	case PlanAdjustmentsPlanPhaseTieredPercentageDiscountAdjustmentFiltersOperatorIncludes, PlanAdjustmentsPlanPhaseTieredPercentageDiscountAdjustmentFiltersOperatorExcludes:
		return true
	}
	return false
}

// One band of a tiered percentage discount. Bounds are denominated in the
// discount's currency. `lower_bound` is the exclusive start of the band and
// `upper_bound` is the inclusive end; `upper_bound` is null only for the
// open-ended final tier.
type PlanAdjustmentsPlanPhaseTieredPercentageDiscountAdjustmentTier struct {
	// Exclusive lower bound of cumulative spend for this tier.
	LowerBound float64 `json:"lower_bound" api:"required"`
	// The percentage (between 0 and 1) discounted from spend that falls within this
	// tier.
	Percentage float64 `json:"percentage" api:"required"`
	// Inclusive upper bound of cumulative spend for this tier; null for the final
	// open-ended tier.
	UpperBound float64                                                            `json:"upper_bound" api:"nullable"`
	JSON       planAdjustmentsPlanPhaseTieredPercentageDiscountAdjustmentTierJSON `json:"-"`
}

// planAdjustmentsPlanPhaseTieredPercentageDiscountAdjustmentTierJSON contains the
// JSON metadata for the struct
// [PlanAdjustmentsPlanPhaseTieredPercentageDiscountAdjustmentTier]
type planAdjustmentsPlanPhaseTieredPercentageDiscountAdjustmentTierJSON struct {
	LowerBound  apijson.Field
	Percentage  apijson.Field
	UpperBound  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *PlanAdjustmentsPlanPhaseTieredPercentageDiscountAdjustmentTier) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r planAdjustmentsPlanPhaseTieredPercentageDiscountAdjustmentTierJSON) RawJSON() string {
	return r.raw
}

type PlanAdjustmentsAdjustmentType string

const (
	PlanAdjustmentsAdjustmentTypeUsageDiscount            PlanAdjustmentsAdjustmentType = "usage_discount"
	PlanAdjustmentsAdjustmentTypeAmountDiscount           PlanAdjustmentsAdjustmentType = "amount_discount"
	PlanAdjustmentsAdjustmentTypePercentageDiscount       PlanAdjustmentsAdjustmentType = "percentage_discount"
	PlanAdjustmentsAdjustmentTypeTieredPercentageDiscount PlanAdjustmentsAdjustmentType = "tiered_percentage_discount"
	PlanAdjustmentsAdjustmentTypeMinimum                  PlanAdjustmentsAdjustmentType = "minimum"
	PlanAdjustmentsAdjustmentTypeMaximum                  PlanAdjustmentsAdjustmentType = "maximum"
)

func (r PlanAdjustmentsAdjustmentType) IsKnown() bool {
	switch r {
	case PlanAdjustmentsAdjustmentTypeUsageDiscount, PlanAdjustmentsAdjustmentTypeAmountDiscount, PlanAdjustmentsAdjustmentTypePercentageDiscount, PlanAdjustmentsAdjustmentTypeTieredPercentageDiscount, PlanAdjustmentsAdjustmentTypeMinimum, PlanAdjustmentsAdjustmentTypeMaximum:
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
	Order int64 `json:"order" api:"required"`
	// The IDs of the prices active in this phase.
	PriceIDs []string          `json:"price_ids" api:"required"`
	JSON     planPlanPhaseJSON `json:"-"`
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
	PriceIDs      apijson.Field
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

// The PlanVersion resource represents the prices and adjustments present on a
// specific version of a plan.
type PlanVersion struct {
	// Adjustments for this plan. If the plan has phases, this includes adjustments
	// across all phases of the plan.
	Adjustments []PlanVersionAdjustment `json:"adjustments" api:"required"`
	CreatedAt   time.Time               `json:"created_at" api:"required" format:"date-time"`
	PlanPhases  []PlanVersionPhase      `json:"plan_phases" api:"required,nullable"`
	// Prices for this plan. If the plan has phases, this includes prices across all
	// phases of the plan.
	Prices  []shared.Price  `json:"prices" api:"required"`
	Version int64           `json:"version" api:"required"`
	JSON    planVersionJSON `json:"-"`
}

// planVersionJSON contains the JSON metadata for the struct [PlanVersion]
type planVersionJSON struct {
	Adjustments apijson.Field
	CreatedAt   apijson.Field
	PlanPhases  apijson.Field
	Prices      apijson.Field
	Version     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *PlanVersion) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r planVersionJSON) RawJSON() string {
	return r.raw
}

type PlanVersionAdjustment struct {
	ID             string                               `json:"id" api:"required"`
	AdjustmentType PlanVersionAdjustmentsAdjustmentType `json:"adjustment_type" api:"required"`
	// This field can have the runtime type of [[]string].
	AppliesToPriceIDs interface{} `json:"applies_to_price_ids" api:"required"`
	// This field can have the runtime type of
	// [[]shared.PlanPhaseUsageDiscountAdjustmentFilter],
	// [[]shared.PlanPhaseAmountDiscountAdjustmentFilter],
	// [[]shared.PlanPhasePercentageDiscountAdjustmentFilter],
	// [[]PlanVersionAdjustmentsPlanPhaseTieredPercentageDiscountAdjustmentFilter],
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
	// This field can have the runtime type of
	// [[]PlanVersionAdjustmentsPlanPhaseTieredPercentageDiscountAdjustmentTier].
	Tiers interface{} `json:"tiers"`
	// The number of usage units by which to discount the price this adjustment applies
	// to in a given billing period.
	UsageDiscount float64                   `json:"usage_discount"`
	JSON          planVersionAdjustmentJSON `json:"-"`
	union         PlanVersionAdjustmentsUnion
}

// planVersionAdjustmentJSON contains the JSON metadata for the struct
// [PlanVersionAdjustment]
type planVersionAdjustmentJSON struct {
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
	Tiers                apijson.Field
	UsageDiscount        apijson.Field
	raw                  string
	ExtraFields          map[string]apijson.Field
}

func (r planVersionAdjustmentJSON) RawJSON() string {
	return r.raw
}

func (r *PlanVersionAdjustment) UnmarshalJSON(data []byte) (err error) {
	*r = PlanVersionAdjustment{}
	err = apijson.UnmarshalRoot(data, &r.union)
	if err != nil {
		return err
	}
	return apijson.Port(r.union, &r)
}

// AsUnion returns a [PlanVersionAdjustmentsUnion] interface which you can cast to
// the specific types for more type safety.
//
// Possible runtime types of the union are
// [shared.PlanPhaseUsageDiscountAdjustment],
// [shared.PlanPhaseAmountDiscountAdjustment],
// [shared.PlanPhasePercentageDiscountAdjustment],
// [PlanVersionAdjustmentsPlanPhaseTieredPercentageDiscountAdjustment],
// [shared.PlanPhaseMinimumAdjustment], [shared.PlanPhaseMaximumAdjustment].
func (r PlanVersionAdjustment) AsUnion() PlanVersionAdjustmentsUnion {
	return r.union
}

// Union satisfied by [shared.PlanPhaseUsageDiscountAdjustment],
// [shared.PlanPhaseAmountDiscountAdjustment],
// [shared.PlanPhasePercentageDiscountAdjustment],
// [PlanVersionAdjustmentsPlanPhaseTieredPercentageDiscountAdjustment],
// [shared.PlanPhaseMinimumAdjustment] or [shared.PlanPhaseMaximumAdjustment].
type PlanVersionAdjustmentsUnion interface {
	ImplementsPlanVersionAdjustment()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*PlanVersionAdjustmentsUnion)(nil)).Elem(),
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
			Type:               reflect.TypeOf(PlanVersionAdjustmentsPlanPhaseTieredPercentageDiscountAdjustment{}),
			DiscriminatorValue: "tiered_percentage_discount",
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

type PlanVersionAdjustmentsPlanPhaseTieredPercentageDiscountAdjustment struct {
	ID             string                                                                          `json:"id" api:"required"`
	AdjustmentType PlanVersionAdjustmentsPlanPhaseTieredPercentageDiscountAdjustmentAdjustmentType `json:"adjustment_type" api:"required"`
	// The price IDs that this adjustment applies to.
	//
	// Deprecated: deprecated
	AppliesToPriceIDs []string `json:"applies_to_price_ids" api:"required"`
	// The filters that determine which prices to apply this adjustment to.
	Filters []PlanVersionAdjustmentsPlanPhaseTieredPercentageDiscountAdjustmentFilter `json:"filters" api:"required"`
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
	// The ordered, contiguous bands of cumulative eligible spend, each discounted at
	// its own percentage (progressive fill-a-tier), applied to the prices this
	// adjustment covers in a given billing period.
	Tiers []PlanVersionAdjustmentsPlanPhaseTieredPercentageDiscountAdjustmentTier `json:"tiers" api:"required"`
	JSON  planVersionAdjustmentsPlanPhaseTieredPercentageDiscountAdjustmentJSON   `json:"-"`
}

// planVersionAdjustmentsPlanPhaseTieredPercentageDiscountAdjustmentJSON contains
// the JSON metadata for the struct
// [PlanVersionAdjustmentsPlanPhaseTieredPercentageDiscountAdjustment]
type planVersionAdjustmentsPlanPhaseTieredPercentageDiscountAdjustmentJSON struct {
	ID                   apijson.Field
	AdjustmentType       apijson.Field
	AppliesToPriceIDs    apijson.Field
	Filters              apijson.Field
	IsInvoiceLevel       apijson.Field
	PlanPhaseOrder       apijson.Field
	Reason               apijson.Field
	ReplacesAdjustmentID apijson.Field
	Tiers                apijson.Field
	raw                  string
	ExtraFields          map[string]apijson.Field
}

func (r *PlanVersionAdjustmentsPlanPhaseTieredPercentageDiscountAdjustment) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r planVersionAdjustmentsPlanPhaseTieredPercentageDiscountAdjustmentJSON) RawJSON() string {
	return r.raw
}

func (r PlanVersionAdjustmentsPlanPhaseTieredPercentageDiscountAdjustment) ImplementsPlanVersionAdjustment() {
}

type PlanVersionAdjustmentsPlanPhaseTieredPercentageDiscountAdjustmentAdjustmentType string

const (
	PlanVersionAdjustmentsPlanPhaseTieredPercentageDiscountAdjustmentAdjustmentTypeTieredPercentageDiscount PlanVersionAdjustmentsPlanPhaseTieredPercentageDiscountAdjustmentAdjustmentType = "tiered_percentage_discount"
)

func (r PlanVersionAdjustmentsPlanPhaseTieredPercentageDiscountAdjustmentAdjustmentType) IsKnown() bool {
	switch r {
	case PlanVersionAdjustmentsPlanPhaseTieredPercentageDiscountAdjustmentAdjustmentTypeTieredPercentageDiscount:
		return true
	}
	return false
}

type PlanVersionAdjustmentsPlanPhaseTieredPercentageDiscountAdjustmentFilter struct {
	// The property of the price to filter on.
	Field PlanVersionAdjustmentsPlanPhaseTieredPercentageDiscountAdjustmentFiltersField `json:"field" api:"required"`
	// Should prices that match the filter be included or excluded.
	Operator PlanVersionAdjustmentsPlanPhaseTieredPercentageDiscountAdjustmentFiltersOperator `json:"operator" api:"required"`
	// The IDs or values that match this filter.
	Values []string                                                                    `json:"values" api:"required"`
	JSON   planVersionAdjustmentsPlanPhaseTieredPercentageDiscountAdjustmentFilterJSON `json:"-"`
}

// planVersionAdjustmentsPlanPhaseTieredPercentageDiscountAdjustmentFilterJSON
// contains the JSON metadata for the struct
// [PlanVersionAdjustmentsPlanPhaseTieredPercentageDiscountAdjustmentFilter]
type planVersionAdjustmentsPlanPhaseTieredPercentageDiscountAdjustmentFilterJSON struct {
	Field       apijson.Field
	Operator    apijson.Field
	Values      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *PlanVersionAdjustmentsPlanPhaseTieredPercentageDiscountAdjustmentFilter) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r planVersionAdjustmentsPlanPhaseTieredPercentageDiscountAdjustmentFilterJSON) RawJSON() string {
	return r.raw
}

// The property of the price to filter on.
type PlanVersionAdjustmentsPlanPhaseTieredPercentageDiscountAdjustmentFiltersField string

const (
	PlanVersionAdjustmentsPlanPhaseTieredPercentageDiscountAdjustmentFiltersFieldPriceID       PlanVersionAdjustmentsPlanPhaseTieredPercentageDiscountAdjustmentFiltersField = "price_id"
	PlanVersionAdjustmentsPlanPhaseTieredPercentageDiscountAdjustmentFiltersFieldItemID        PlanVersionAdjustmentsPlanPhaseTieredPercentageDiscountAdjustmentFiltersField = "item_id"
	PlanVersionAdjustmentsPlanPhaseTieredPercentageDiscountAdjustmentFiltersFieldPriceType     PlanVersionAdjustmentsPlanPhaseTieredPercentageDiscountAdjustmentFiltersField = "price_type"
	PlanVersionAdjustmentsPlanPhaseTieredPercentageDiscountAdjustmentFiltersFieldCurrency      PlanVersionAdjustmentsPlanPhaseTieredPercentageDiscountAdjustmentFiltersField = "currency"
	PlanVersionAdjustmentsPlanPhaseTieredPercentageDiscountAdjustmentFiltersFieldPricingUnitID PlanVersionAdjustmentsPlanPhaseTieredPercentageDiscountAdjustmentFiltersField = "pricing_unit_id"
)

func (r PlanVersionAdjustmentsPlanPhaseTieredPercentageDiscountAdjustmentFiltersField) IsKnown() bool {
	switch r {
	case PlanVersionAdjustmentsPlanPhaseTieredPercentageDiscountAdjustmentFiltersFieldPriceID, PlanVersionAdjustmentsPlanPhaseTieredPercentageDiscountAdjustmentFiltersFieldItemID, PlanVersionAdjustmentsPlanPhaseTieredPercentageDiscountAdjustmentFiltersFieldPriceType, PlanVersionAdjustmentsPlanPhaseTieredPercentageDiscountAdjustmentFiltersFieldCurrency, PlanVersionAdjustmentsPlanPhaseTieredPercentageDiscountAdjustmentFiltersFieldPricingUnitID:
		return true
	}
	return false
}

// Should prices that match the filter be included or excluded.
type PlanVersionAdjustmentsPlanPhaseTieredPercentageDiscountAdjustmentFiltersOperator string

const (
	PlanVersionAdjustmentsPlanPhaseTieredPercentageDiscountAdjustmentFiltersOperatorIncludes PlanVersionAdjustmentsPlanPhaseTieredPercentageDiscountAdjustmentFiltersOperator = "includes"
	PlanVersionAdjustmentsPlanPhaseTieredPercentageDiscountAdjustmentFiltersOperatorExcludes PlanVersionAdjustmentsPlanPhaseTieredPercentageDiscountAdjustmentFiltersOperator = "excludes"
)

func (r PlanVersionAdjustmentsPlanPhaseTieredPercentageDiscountAdjustmentFiltersOperator) IsKnown() bool {
	switch r {
	case PlanVersionAdjustmentsPlanPhaseTieredPercentageDiscountAdjustmentFiltersOperatorIncludes, PlanVersionAdjustmentsPlanPhaseTieredPercentageDiscountAdjustmentFiltersOperatorExcludes:
		return true
	}
	return false
}

// One band of a tiered percentage discount. Bounds are denominated in the
// discount's currency. `lower_bound` is the exclusive start of the band and
// `upper_bound` is the inclusive end; `upper_bound` is null only for the
// open-ended final tier.
type PlanVersionAdjustmentsPlanPhaseTieredPercentageDiscountAdjustmentTier struct {
	// Exclusive lower bound of cumulative spend for this tier.
	LowerBound float64 `json:"lower_bound" api:"required"`
	// The percentage (between 0 and 1) discounted from spend that falls within this
	// tier.
	Percentage float64 `json:"percentage" api:"required"`
	// Inclusive upper bound of cumulative spend for this tier; null for the final
	// open-ended tier.
	UpperBound float64                                                                   `json:"upper_bound" api:"nullable"`
	JSON       planVersionAdjustmentsPlanPhaseTieredPercentageDiscountAdjustmentTierJSON `json:"-"`
}

// planVersionAdjustmentsPlanPhaseTieredPercentageDiscountAdjustmentTierJSON
// contains the JSON metadata for the struct
// [PlanVersionAdjustmentsPlanPhaseTieredPercentageDiscountAdjustmentTier]
type planVersionAdjustmentsPlanPhaseTieredPercentageDiscountAdjustmentTierJSON struct {
	LowerBound  apijson.Field
	Percentage  apijson.Field
	UpperBound  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *PlanVersionAdjustmentsPlanPhaseTieredPercentageDiscountAdjustmentTier) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r planVersionAdjustmentsPlanPhaseTieredPercentageDiscountAdjustmentTierJSON) RawJSON() string {
	return r.raw
}

type PlanVersionAdjustmentsAdjustmentType string

const (
	PlanVersionAdjustmentsAdjustmentTypeUsageDiscount            PlanVersionAdjustmentsAdjustmentType = "usage_discount"
	PlanVersionAdjustmentsAdjustmentTypeAmountDiscount           PlanVersionAdjustmentsAdjustmentType = "amount_discount"
	PlanVersionAdjustmentsAdjustmentTypePercentageDiscount       PlanVersionAdjustmentsAdjustmentType = "percentage_discount"
	PlanVersionAdjustmentsAdjustmentTypeTieredPercentageDiscount PlanVersionAdjustmentsAdjustmentType = "tiered_percentage_discount"
	PlanVersionAdjustmentsAdjustmentTypeMinimum                  PlanVersionAdjustmentsAdjustmentType = "minimum"
	PlanVersionAdjustmentsAdjustmentTypeMaximum                  PlanVersionAdjustmentsAdjustmentType = "maximum"
)

func (r PlanVersionAdjustmentsAdjustmentType) IsKnown() bool {
	switch r {
	case PlanVersionAdjustmentsAdjustmentTypeUsageDiscount, PlanVersionAdjustmentsAdjustmentTypeAmountDiscount, PlanVersionAdjustmentsAdjustmentTypePercentageDiscount, PlanVersionAdjustmentsAdjustmentTypeTieredPercentageDiscount, PlanVersionAdjustmentsAdjustmentTypeMinimum, PlanVersionAdjustmentsAdjustmentTypeMaximum:
		return true
	}
	return false
}

type PlanVersionPhase struct {
	ID          string `json:"id" api:"required"`
	Description string `json:"description" api:"required,nullable"`
	// How many terms of length `duration_unit` this phase is active for. If null, this
	// phase is evergreen and active indefinitely
	Duration     int64                        `json:"duration" api:"required,nullable"`
	DurationUnit PlanVersionPhaseDurationUnit `json:"duration_unit" api:"required,nullable"`
	Name         string                       `json:"name" api:"required"`
	// Determines the ordering of the phase in a plan's lifecycle. 1 = first phase.
	Order int64 `json:"order" api:"required"`
	// The IDs of the prices active in this phase.
	PriceIDs []string             `json:"price_ids" api:"required"`
	JSON     planVersionPhaseJSON `json:"-"`
}

// planVersionPhaseJSON contains the JSON metadata for the struct
// [PlanVersionPhase]
type planVersionPhaseJSON struct {
	ID           apijson.Field
	Description  apijson.Field
	Duration     apijson.Field
	DurationUnit apijson.Field
	Name         apijson.Field
	Order        apijson.Field
	PriceIDs     apijson.Field
	raw          string
	ExtraFields  map[string]apijson.Field
}

func (r *PlanVersionPhase) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r planVersionPhaseJSON) RawJSON() string {
	return r.raw
}

type PlanVersionPhaseDurationUnit string

const (
	PlanVersionPhaseDurationUnitDaily      PlanVersionPhaseDurationUnit = "daily"
	PlanVersionPhaseDurationUnitMonthly    PlanVersionPhaseDurationUnit = "monthly"
	PlanVersionPhaseDurationUnitQuarterly  PlanVersionPhaseDurationUnit = "quarterly"
	PlanVersionPhaseDurationUnitSemiAnnual PlanVersionPhaseDurationUnit = "semi_annual"
	PlanVersionPhaseDurationUnitAnnual     PlanVersionPhaseDurationUnit = "annual"
)

func (r PlanVersionPhaseDurationUnit) IsKnown() bool {
	switch r {
	case PlanVersionPhaseDurationUnitDaily, PlanVersionPhaseDurationUnitMonthly, PlanVersionPhaseDurationUnitQuarterly, PlanVersionPhaseDurationUnitSemiAnnual, PlanVersionPhaseDurationUnitAnnual:
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
	LicenseAllocationPrice param.Field[PlanNewParamsPricesLicenseAllocationPrice] `json:"license_allocation_price"`
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
	ItemID param.Field[string] `json:"item_id" api:"required"`
	// License allocations to associate with this price. Each entry defines a
	// per-license credit pool granted each cadence. Requires license_type_id or
	// license_type_configuration to be set. Usage prices draw down from these
	// allocations only if they set the same license type.
	LicenseAllocations param.Field[[]PlanNewParamsPricesLicenseAllocationPriceLicenseAllocation] `json:"license_allocations" api:"required"`
	// The pricing model type
	ModelType param.Field[PlanNewParamsPricesLicenseAllocationPriceModelType] `json:"model_type" api:"required"`
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
	ConversionRateConfig param.Field[PlanNewParamsPricesLicenseAllocationPriceConversionRateConfigUnion] `json:"conversion_rate_config"`
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
	// The ID of the license type to associate with this price. On a usage price this
	// also marks the price as eligible to draw down from that license type's
	// allocation; a usage price created without it is billed normally. Usage prices
	// with a license type must use the `unit` model, and only draw down when their
	// currency matches the allocation's.
	LicenseTypeID param.Field[string] `json:"license_type_id"`
	// User-specified key/value pairs for the resource. Individual keys can be removed
	// by setting the value to `null`, and the entire metadata mapping can be cleared
	// by setting `metadata` to `null`.
	Metadata param.Field[map[string]string] `json:"metadata"`
	// A transient ID that can be used to reference this price when adding adjustments
	// in the same API call.
	ReferenceID param.Field[string] `json:"reference_id"`
}

func (r PlanNewParamsPricesLicenseAllocationPrice) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// The cadence to bill for this price on.
type PlanNewParamsPricesLicenseAllocationPriceCadence string

const (
	PlanNewParamsPricesLicenseAllocationPriceCadenceOneTime    PlanNewParamsPricesLicenseAllocationPriceCadence = "one_time"
	PlanNewParamsPricesLicenseAllocationPriceCadenceMonthly    PlanNewParamsPricesLicenseAllocationPriceCadence = "monthly"
	PlanNewParamsPricesLicenseAllocationPriceCadenceQuarterly  PlanNewParamsPricesLicenseAllocationPriceCadence = "quarterly"
	PlanNewParamsPricesLicenseAllocationPriceCadenceSemiAnnual PlanNewParamsPricesLicenseAllocationPriceCadence = "semi_annual"
	PlanNewParamsPricesLicenseAllocationPriceCadenceAnnual     PlanNewParamsPricesLicenseAllocationPriceCadence = "annual"
	PlanNewParamsPricesLicenseAllocationPriceCadenceCustom     PlanNewParamsPricesLicenseAllocationPriceCadence = "custom"
)

func (r PlanNewParamsPricesLicenseAllocationPriceCadence) IsKnown() bool {
	switch r {
	case PlanNewParamsPricesLicenseAllocationPriceCadenceOneTime, PlanNewParamsPricesLicenseAllocationPriceCadenceMonthly, PlanNewParamsPricesLicenseAllocationPriceCadenceQuarterly, PlanNewParamsPricesLicenseAllocationPriceCadenceSemiAnnual, PlanNewParamsPricesLicenseAllocationPriceCadenceAnnual, PlanNewParamsPricesLicenseAllocationPriceCadenceCustom:
		return true
	}
	return false
}

type PlanNewParamsPricesLicenseAllocationPriceLicenseAllocation struct {
	// The amount of credits granted per active license per cadence.
	Amount param.Field[string] `json:"amount" api:"required"`
	// The currency of the license allocation.
	Currency param.Field[string] `json:"currency" api:"required"`
	// When True, overage beyond the allocation is written off.
	WriteOffOverage param.Field[bool] `json:"write_off_overage"`
}

func (r PlanNewParamsPricesLicenseAllocationPriceLicenseAllocation) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// The pricing model type
type PlanNewParamsPricesLicenseAllocationPriceModelType string

const (
	PlanNewParamsPricesLicenseAllocationPriceModelTypeUnit PlanNewParamsPricesLicenseAllocationPriceModelType = "unit"
)

func (r PlanNewParamsPricesLicenseAllocationPriceModelType) IsKnown() bool {
	switch r {
	case PlanNewParamsPricesLicenseAllocationPriceModelTypeUnit:
		return true
	}
	return false
}

type PlanNewParamsPricesLicenseAllocationPriceConversionRateConfig struct {
	ConversionRateType param.Field[PlanNewParamsPricesLicenseAllocationPriceConversionRateConfigConversionRateType] `json:"conversion_rate_type" api:"required"`
	TieredConfig       param.Field[shared.ConversionRateTieredConfigParam]                                          `json:"tiered_config"`
	UnitConfig         param.Field[shared.ConversionRateUnitConfigParam]                                            `json:"unit_config"`
}

func (r PlanNewParamsPricesLicenseAllocationPriceConversionRateConfig) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r PlanNewParamsPricesLicenseAllocationPriceConversionRateConfig) ImplementsPlanNewParamsPricesLicenseAllocationPriceConversionRateConfigUnion() {
}

// Satisfied by [shared.UnitConversionRateConfigParam],
// [shared.TieredConversionRateConfigParam],
// [PlanNewParamsPricesLicenseAllocationPriceConversionRateConfig].
type PlanNewParamsPricesLicenseAllocationPriceConversionRateConfigUnion interface {
	ImplementsPlanNewParamsPricesLicenseAllocationPriceConversionRateConfigUnion()
}

type PlanNewParamsPricesLicenseAllocationPriceConversionRateConfigConversionRateType string

const (
	PlanNewParamsPricesLicenseAllocationPriceConversionRateConfigConversionRateTypeUnit   PlanNewParamsPricesLicenseAllocationPriceConversionRateConfigConversionRateType = "unit"
	PlanNewParamsPricesLicenseAllocationPriceConversionRateConfigConversionRateTypeTiered PlanNewParamsPricesLicenseAllocationPriceConversionRateConfigConversionRateType = "tiered"
)

func (r PlanNewParamsPricesLicenseAllocationPriceConversionRateConfigConversionRateType) IsKnown() bool {
	switch r {
	case PlanNewParamsPricesLicenseAllocationPriceConversionRateConfigConversionRateTypeUnit, PlanNewParamsPricesLicenseAllocationPriceConversionRateConfigConversionRateTypeTiered:
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
	Currency                   param.Field[string]      `json:"currency"`
	DailyCreditAllowanceConfig param.Field[interface{}] `json:"daily_credit_allowance_config"`
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
	GroupedTieredMatrixConfig         param.Field[interface{}] `json:"grouped_tiered_matrix_config"`
	GroupedTieredPackageConfig        param.Field[interface{}] `json:"grouped_tiered_package_config"`
	GroupedWithMeteredMinimumConfig   param.Field[interface{}] `json:"grouped_with_metered_minimum_config"`
	GroupedWithMinMaxThresholdsConfig param.Field[interface{}] `json:"grouped_with_min_max_thresholds_config"`
	GroupedWithProratedMinimumConfig  param.Field[interface{}] `json:"grouped_with_prorated_minimum_config"`
	// The property used to group this price on an invoice
	InvoiceGroupingKey param.Field[string] `json:"invoice_grouping_key"`
	// Within each billing cycle, specifies the cadence at which invoices are produced.
	// If unspecified, a single invoice is produced per billing cycle.
	InvoicingCycleConfiguration param.Field[shared.NewBillingCycleConfigurationParam] `json:"invoicing_cycle_configuration"`
	// The ID of the license type to associate with this price. On a usage price this
	// also marks the price as eligible to draw down from that license type's
	// allocation; a usage price created without it is billed normally. Usage prices
	// with a license type must use the `unit` model, and only draw down when their
	// currency matches the allocation's.
	LicenseTypeID param.Field[string] `json:"license_type_id"`
	// Configuration for matrix pricing
	MatrixConfig param.Field[shared.MatrixConfigParam] `json:"matrix_config"`
	// Configuration for matrix_with_allocation pricing
	MatrixWithAllocationConfig         param.Field[shared.MatrixWithAllocationConfigParam] `json:"matrix_with_allocation_config"`
	MatrixWithDisplayNameConfig        param.Field[interface{}]                            `json:"matrix_with_display_name_config"`
	MatrixWithThresholdDiscountsConfig param.Field[interface{}]                            `json:"matrix_with_threshold_discounts_config"`
	MaxGroupTieredPackageConfig        param.Field[interface{}]                            `json:"max_group_tiered_package_config"`
	Metadata                           param.Field[interface{}]                            `json:"metadata"`
	MeteredAllowanceConfig             param.Field[interface{}]                            `json:"metered_allowance_config"`
	MinimumCompositeConfig             param.Field[interface{}]                            `json:"minimum_composite_config"`
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
	TieredConfig                     param.Field[shared.TieredConfigParam] `json:"tiered_config"`
	TieredMatrixWithAllocationConfig param.Field[interface{}]              `json:"tiered_matrix_with_allocation_config"`
	TieredPackageConfig              param.Field[interface{}]              `json:"tiered_package_config"`
	TieredPackageWithMinimumConfig   param.Field[interface{}]              `json:"tiered_package_with_minimum_config"`
	TieredWithMinimumConfig          param.Field[interface{}]              `json:"tiered_with_minimum_config"`
	TieredWithProrationConfig        param.Field[interface{}]              `json:"tiered_with_proration_config"`
	// Configuration for unit pricing
	UnitConfig               param.Field[shared.UnitConfigParam] `json:"unit_config"`
	UnitWithMultiplierConfig param.Field[interface{}]            `json:"unit_with_multiplier_config"`
	UnitWithPercentConfig    param.Field[interface{}]            `json:"unit_with_percent_config"`
	UnitWithProrationConfig  param.Field[interface{}]            `json:"unit_with_proration_config"`
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
// [PlanNewParamsPricesPriceNewPlanGroupedTieredMatrixPrice],
// [shared.NewPlanTieredPackageWithMinimumPriceParam],
// [shared.NewPlanPackageWithAllocationPriceParam],
// [shared.NewPlanUnitWithPercentPriceParam],
// [shared.NewPlanMatrixWithAllocationPriceParam],
// [PlanNewParamsPricesPriceNewPlanTieredMatrixWithAllocationPrice],
// [PlanNewParamsPricesPriceNewPlanMatrixWithThresholdDiscountsPrice],
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
// [PlanNewParamsPricesPriceNewPlanDailyCreditAllowancePrice],
// [PlanNewParamsPricesPriceNewPlanMeteredAllowancePrice],
// [shared.NewPlanMinimumCompositePriceParam],
// [PlanNewParamsPricesPriceNewPlanPercentCompositePrice],
// [PlanNewParamsPricesPriceNewPlanEventOutputPrice],
// [PlanNewParamsPricesPriceNewPlanUnitWithMultiplierPrice],
// [PlanNewParamsPricesPrice].
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
	// The ID of the license type to associate with this price. On a usage price this
	// also marks the price as eligible to draw down from that license type's
	// allocation; a usage price created without it is billed normally. Usage prices
	// with a license type must use the `unit` model, and only draw down when their
	// currency matches the allocation's.
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
	PlanNewParamsPricesPriceNewPlanBulkWithFiltersPriceCadenceOneTime    PlanNewParamsPricesPriceNewPlanBulkWithFiltersPriceCadence = "one_time"
	PlanNewParamsPricesPriceNewPlanBulkWithFiltersPriceCadenceMonthly    PlanNewParamsPricesPriceNewPlanBulkWithFiltersPriceCadence = "monthly"
	PlanNewParamsPricesPriceNewPlanBulkWithFiltersPriceCadenceQuarterly  PlanNewParamsPricesPriceNewPlanBulkWithFiltersPriceCadence = "quarterly"
	PlanNewParamsPricesPriceNewPlanBulkWithFiltersPriceCadenceSemiAnnual PlanNewParamsPricesPriceNewPlanBulkWithFiltersPriceCadence = "semi_annual"
	PlanNewParamsPricesPriceNewPlanBulkWithFiltersPriceCadenceAnnual     PlanNewParamsPricesPriceNewPlanBulkWithFiltersPriceCadence = "annual"
	PlanNewParamsPricesPriceNewPlanBulkWithFiltersPriceCadenceCustom     PlanNewParamsPricesPriceNewPlanBulkWithFiltersPriceCadence = "custom"
)

func (r PlanNewParamsPricesPriceNewPlanBulkWithFiltersPriceCadence) IsKnown() bool {
	switch r {
	case PlanNewParamsPricesPriceNewPlanBulkWithFiltersPriceCadenceOneTime, PlanNewParamsPricesPriceNewPlanBulkWithFiltersPriceCadenceMonthly, PlanNewParamsPricesPriceNewPlanBulkWithFiltersPriceCadenceQuarterly, PlanNewParamsPricesPriceNewPlanBulkWithFiltersPriceCadenceSemiAnnual, PlanNewParamsPricesPriceNewPlanBulkWithFiltersPriceCadenceAnnual, PlanNewParamsPricesPriceNewPlanBulkWithFiltersPriceCadenceCustom:
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

type PlanNewParamsPricesPriceNewPlanGroupedTieredMatrixPrice struct {
	// The cadence to bill for this price on.
	Cadence param.Field[PlanNewParamsPricesPriceNewPlanGroupedTieredMatrixPriceCadence] `json:"cadence" api:"required"`
	// Configuration for grouped_tiered_matrix pricing
	GroupedTieredMatrixConfig param.Field[PlanNewParamsPricesPriceNewPlanGroupedTieredMatrixPriceGroupedTieredMatrixConfig] `json:"grouped_tiered_matrix_config" api:"required"`
	// The id of the item the price will be associated with.
	ItemID param.Field[string] `json:"item_id" api:"required"`
	// The pricing model type
	ModelType param.Field[PlanNewParamsPricesPriceNewPlanGroupedTieredMatrixPriceModelType] `json:"model_type" api:"required"`
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
	ConversionRateConfig param.Field[PlanNewParamsPricesPriceNewPlanGroupedTieredMatrixPriceConversionRateConfigUnion] `json:"conversion_rate_config"`
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
	// The ID of the license type to associate with this price. On a usage price this
	// also marks the price as eligible to draw down from that license type's
	// allocation; a usage price created without it is billed normally. Usage prices
	// with a license type must use the `unit` model, and only draw down when their
	// currency matches the allocation's.
	LicenseTypeID param.Field[string] `json:"license_type_id"`
	// User-specified key/value pairs for the resource. Individual keys can be removed
	// by setting the value to `null`, and the entire metadata mapping can be cleared
	// by setting `metadata` to `null`.
	Metadata param.Field[map[string]string] `json:"metadata"`
	// A transient ID that can be used to reference this price when adding adjustments
	// in the same API call.
	ReferenceID param.Field[string] `json:"reference_id"`
}

func (r PlanNewParamsPricesPriceNewPlanGroupedTieredMatrixPrice) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r PlanNewParamsPricesPriceNewPlanGroupedTieredMatrixPrice) ImplementsPlanNewParamsPricesPriceUnion() {
}

// The cadence to bill for this price on.
type PlanNewParamsPricesPriceNewPlanGroupedTieredMatrixPriceCadence string

const (
	PlanNewParamsPricesPriceNewPlanGroupedTieredMatrixPriceCadenceOneTime    PlanNewParamsPricesPriceNewPlanGroupedTieredMatrixPriceCadence = "one_time"
	PlanNewParamsPricesPriceNewPlanGroupedTieredMatrixPriceCadenceMonthly    PlanNewParamsPricesPriceNewPlanGroupedTieredMatrixPriceCadence = "monthly"
	PlanNewParamsPricesPriceNewPlanGroupedTieredMatrixPriceCadenceQuarterly  PlanNewParamsPricesPriceNewPlanGroupedTieredMatrixPriceCadence = "quarterly"
	PlanNewParamsPricesPriceNewPlanGroupedTieredMatrixPriceCadenceSemiAnnual PlanNewParamsPricesPriceNewPlanGroupedTieredMatrixPriceCadence = "semi_annual"
	PlanNewParamsPricesPriceNewPlanGroupedTieredMatrixPriceCadenceAnnual     PlanNewParamsPricesPriceNewPlanGroupedTieredMatrixPriceCadence = "annual"
	PlanNewParamsPricesPriceNewPlanGroupedTieredMatrixPriceCadenceCustom     PlanNewParamsPricesPriceNewPlanGroupedTieredMatrixPriceCadence = "custom"
)

func (r PlanNewParamsPricesPriceNewPlanGroupedTieredMatrixPriceCadence) IsKnown() bool {
	switch r {
	case PlanNewParamsPricesPriceNewPlanGroupedTieredMatrixPriceCadenceOneTime, PlanNewParamsPricesPriceNewPlanGroupedTieredMatrixPriceCadenceMonthly, PlanNewParamsPricesPriceNewPlanGroupedTieredMatrixPriceCadenceQuarterly, PlanNewParamsPricesPriceNewPlanGroupedTieredMatrixPriceCadenceSemiAnnual, PlanNewParamsPricesPriceNewPlanGroupedTieredMatrixPriceCadenceAnnual, PlanNewParamsPricesPriceNewPlanGroupedTieredMatrixPriceCadenceCustom:
		return true
	}
	return false
}

// Configuration for grouped_tiered_matrix pricing
type PlanNewParamsPricesPriceNewPlanGroupedTieredMatrixPriceGroupedTieredMatrixConfig struct {
	// Per unit rate for usage whose dimension value has no configured tiers
	DefaultUnitAmount param.Field[string] `json:"default_unit_amount" api:"required"`
	// The billable metric property used to group usage before tiering
	Dimension param.Field[string] `json:"dimension" api:"required"`
	// Graduated tiers keyed by dimension value; usage for a value is tiered only
	// against its own rows
	Tiers param.Field[[]PlanNewParamsPricesPriceNewPlanGroupedTieredMatrixPriceGroupedTieredMatrixConfigTier] `json:"tiers" api:"required"`
}

func (r PlanNewParamsPricesPriceNewPlanGroupedTieredMatrixPriceGroupedTieredMatrixConfig) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Configuration for a single tier scoped to a dimension value
type PlanNewParamsPricesPriceNewPlanGroupedTieredMatrixPriceGroupedTieredMatrixConfigTier struct {
	// The dimension value this tier applies to
	DimensionValue param.Field[string] `json:"dimension_value" api:"required"`
	TierLowerBound param.Field[string] `json:"tier_lower_bound" api:"required"`
	// Per unit amount
	UnitAmount param.Field[string] `json:"unit_amount" api:"required"`
}

func (r PlanNewParamsPricesPriceNewPlanGroupedTieredMatrixPriceGroupedTieredMatrixConfigTier) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// The pricing model type
type PlanNewParamsPricesPriceNewPlanGroupedTieredMatrixPriceModelType string

const (
	PlanNewParamsPricesPriceNewPlanGroupedTieredMatrixPriceModelTypeGroupedTieredMatrix PlanNewParamsPricesPriceNewPlanGroupedTieredMatrixPriceModelType = "grouped_tiered_matrix"
)

func (r PlanNewParamsPricesPriceNewPlanGroupedTieredMatrixPriceModelType) IsKnown() bool {
	switch r {
	case PlanNewParamsPricesPriceNewPlanGroupedTieredMatrixPriceModelTypeGroupedTieredMatrix:
		return true
	}
	return false
}

type PlanNewParamsPricesPriceNewPlanGroupedTieredMatrixPriceConversionRateConfig struct {
	ConversionRateType param.Field[PlanNewParamsPricesPriceNewPlanGroupedTieredMatrixPriceConversionRateConfigConversionRateType] `json:"conversion_rate_type" api:"required"`
	TieredConfig       param.Field[shared.ConversionRateTieredConfigParam]                                                        `json:"tiered_config"`
	UnitConfig         param.Field[shared.ConversionRateUnitConfigParam]                                                          `json:"unit_config"`
}

func (r PlanNewParamsPricesPriceNewPlanGroupedTieredMatrixPriceConversionRateConfig) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r PlanNewParamsPricesPriceNewPlanGroupedTieredMatrixPriceConversionRateConfig) ImplementsPlanNewParamsPricesPriceNewPlanGroupedTieredMatrixPriceConversionRateConfigUnion() {
}

// Satisfied by [shared.UnitConversionRateConfigParam],
// [shared.TieredConversionRateConfigParam],
// [PlanNewParamsPricesPriceNewPlanGroupedTieredMatrixPriceConversionRateConfig].
type PlanNewParamsPricesPriceNewPlanGroupedTieredMatrixPriceConversionRateConfigUnion interface {
	ImplementsPlanNewParamsPricesPriceNewPlanGroupedTieredMatrixPriceConversionRateConfigUnion()
}

type PlanNewParamsPricesPriceNewPlanGroupedTieredMatrixPriceConversionRateConfigConversionRateType string

const (
	PlanNewParamsPricesPriceNewPlanGroupedTieredMatrixPriceConversionRateConfigConversionRateTypeUnit   PlanNewParamsPricesPriceNewPlanGroupedTieredMatrixPriceConversionRateConfigConversionRateType = "unit"
	PlanNewParamsPricesPriceNewPlanGroupedTieredMatrixPriceConversionRateConfigConversionRateTypeTiered PlanNewParamsPricesPriceNewPlanGroupedTieredMatrixPriceConversionRateConfigConversionRateType = "tiered"
)

func (r PlanNewParamsPricesPriceNewPlanGroupedTieredMatrixPriceConversionRateConfigConversionRateType) IsKnown() bool {
	switch r {
	case PlanNewParamsPricesPriceNewPlanGroupedTieredMatrixPriceConversionRateConfigConversionRateTypeUnit, PlanNewParamsPricesPriceNewPlanGroupedTieredMatrixPriceConversionRateConfigConversionRateTypeTiered:
		return true
	}
	return false
}

type PlanNewParamsPricesPriceNewPlanTieredMatrixWithAllocationPrice struct {
	// The cadence to bill for this price on.
	Cadence param.Field[PlanNewParamsPricesPriceNewPlanTieredMatrixWithAllocationPriceCadence] `json:"cadence" api:"required"`
	// The id of the item the price will be associated with.
	ItemID param.Field[string] `json:"item_id" api:"required"`
	// The pricing model type
	ModelType param.Field[PlanNewParamsPricesPriceNewPlanTieredMatrixWithAllocationPriceModelType] `json:"model_type" api:"required"`
	// The name of the price.
	Name param.Field[string] `json:"name" api:"required"`
	// Configuration for tiered_matrix_with_allocation pricing
	TieredMatrixWithAllocationConfig param.Field[PlanNewParamsPricesPriceNewPlanTieredMatrixWithAllocationPriceTieredMatrixWithAllocationConfig] `json:"tiered_matrix_with_allocation_config" api:"required"`
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
	ConversionRateConfig param.Field[PlanNewParamsPricesPriceNewPlanTieredMatrixWithAllocationPriceConversionRateConfigUnion] `json:"conversion_rate_config"`
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
	// The ID of the license type to associate with this price. On a usage price this
	// also marks the price as eligible to draw down from that license type's
	// allocation; a usage price created without it is billed normally. Usage prices
	// with a license type must use the `unit` model, and only draw down when their
	// currency matches the allocation's.
	LicenseTypeID param.Field[string] `json:"license_type_id"`
	// User-specified key/value pairs for the resource. Individual keys can be removed
	// by setting the value to `null`, and the entire metadata mapping can be cleared
	// by setting `metadata` to `null`.
	Metadata param.Field[map[string]string] `json:"metadata"`
	// A transient ID that can be used to reference this price when adding adjustments
	// in the same API call.
	ReferenceID param.Field[string] `json:"reference_id"`
}

func (r PlanNewParamsPricesPriceNewPlanTieredMatrixWithAllocationPrice) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r PlanNewParamsPricesPriceNewPlanTieredMatrixWithAllocationPrice) ImplementsPlanNewParamsPricesPriceUnion() {
}

// The cadence to bill for this price on.
type PlanNewParamsPricesPriceNewPlanTieredMatrixWithAllocationPriceCadence string

const (
	PlanNewParamsPricesPriceNewPlanTieredMatrixWithAllocationPriceCadenceOneTime    PlanNewParamsPricesPriceNewPlanTieredMatrixWithAllocationPriceCadence = "one_time"
	PlanNewParamsPricesPriceNewPlanTieredMatrixWithAllocationPriceCadenceMonthly    PlanNewParamsPricesPriceNewPlanTieredMatrixWithAllocationPriceCadence = "monthly"
	PlanNewParamsPricesPriceNewPlanTieredMatrixWithAllocationPriceCadenceQuarterly  PlanNewParamsPricesPriceNewPlanTieredMatrixWithAllocationPriceCadence = "quarterly"
	PlanNewParamsPricesPriceNewPlanTieredMatrixWithAllocationPriceCadenceSemiAnnual PlanNewParamsPricesPriceNewPlanTieredMatrixWithAllocationPriceCadence = "semi_annual"
	PlanNewParamsPricesPriceNewPlanTieredMatrixWithAllocationPriceCadenceAnnual     PlanNewParamsPricesPriceNewPlanTieredMatrixWithAllocationPriceCadence = "annual"
	PlanNewParamsPricesPriceNewPlanTieredMatrixWithAllocationPriceCadenceCustom     PlanNewParamsPricesPriceNewPlanTieredMatrixWithAllocationPriceCadence = "custom"
)

func (r PlanNewParamsPricesPriceNewPlanTieredMatrixWithAllocationPriceCadence) IsKnown() bool {
	switch r {
	case PlanNewParamsPricesPriceNewPlanTieredMatrixWithAllocationPriceCadenceOneTime, PlanNewParamsPricesPriceNewPlanTieredMatrixWithAllocationPriceCadenceMonthly, PlanNewParamsPricesPriceNewPlanTieredMatrixWithAllocationPriceCadenceQuarterly, PlanNewParamsPricesPriceNewPlanTieredMatrixWithAllocationPriceCadenceSemiAnnual, PlanNewParamsPricesPriceNewPlanTieredMatrixWithAllocationPriceCadenceAnnual, PlanNewParamsPricesPriceNewPlanTieredMatrixWithAllocationPriceCadenceCustom:
		return true
	}
	return false
}

// The pricing model type
type PlanNewParamsPricesPriceNewPlanTieredMatrixWithAllocationPriceModelType string

const (
	PlanNewParamsPricesPriceNewPlanTieredMatrixWithAllocationPriceModelTypeTieredMatrixWithAllocation PlanNewParamsPricesPriceNewPlanTieredMatrixWithAllocationPriceModelType = "tiered_matrix_with_allocation"
)

func (r PlanNewParamsPricesPriceNewPlanTieredMatrixWithAllocationPriceModelType) IsKnown() bool {
	switch r {
	case PlanNewParamsPricesPriceNewPlanTieredMatrixWithAllocationPriceModelTypeTieredMatrixWithAllocation:
		return true
	}
	return false
}

// Configuration for tiered_matrix_with_allocation pricing
type PlanNewParamsPricesPriceNewPlanTieredMatrixWithAllocationPriceTieredMatrixWithAllocationConfig struct {
	// Usage allocation, pooled across all matrix cells
	Allocation param.Field[string] `json:"allocation" api:"required"`
	// Per unit rate for usage whose matrix cell has no configured tiers
	DefaultUnitAmount param.Field[string] `json:"default_unit_amount" api:"required"`
	// One or two event property values to evaluate matrix cells by
	Dimensions param.Field[[]string] `json:"dimensions" api:"required"`
	// Graduated tiers keyed by matrix cell; usage in a cell is tiered only against its
	// own rows
	Tiers param.Field[[]PlanNewParamsPricesPriceNewPlanTieredMatrixWithAllocationPriceTieredMatrixWithAllocationConfigTier] `json:"tiers" api:"required"`
}

func (r PlanNewParamsPricesPriceNewPlanTieredMatrixWithAllocationPriceTieredMatrixWithAllocationConfig) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Configuration for a single tier scoped to one matrix cell
type PlanNewParamsPricesPriceNewPlanTieredMatrixWithAllocationPriceTieredMatrixWithAllocationConfigTier struct {
	// The matrix cell this tier applies to, as one or two dimension values
	DimensionValues param.Field[[]string] `json:"dimension_values" api:"required"`
	// Exclusive tier starting value. The tier runs up to and including the next bound
	// configured for the same matrix cell.
	TierLowerBound param.Field[string] `json:"tier_lower_bound" api:"required"`
	// Per unit amount
	UnitAmount param.Field[string] `json:"unit_amount" api:"required"`
}

func (r PlanNewParamsPricesPriceNewPlanTieredMatrixWithAllocationPriceTieredMatrixWithAllocationConfigTier) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type PlanNewParamsPricesPriceNewPlanTieredMatrixWithAllocationPriceConversionRateConfig struct {
	ConversionRateType param.Field[PlanNewParamsPricesPriceNewPlanTieredMatrixWithAllocationPriceConversionRateConfigConversionRateType] `json:"conversion_rate_type" api:"required"`
	TieredConfig       param.Field[shared.ConversionRateTieredConfigParam]                                                               `json:"tiered_config"`
	UnitConfig         param.Field[shared.ConversionRateUnitConfigParam]                                                                 `json:"unit_config"`
}

func (r PlanNewParamsPricesPriceNewPlanTieredMatrixWithAllocationPriceConversionRateConfig) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r PlanNewParamsPricesPriceNewPlanTieredMatrixWithAllocationPriceConversionRateConfig) ImplementsPlanNewParamsPricesPriceNewPlanTieredMatrixWithAllocationPriceConversionRateConfigUnion() {
}

// Satisfied by [shared.UnitConversionRateConfigParam],
// [shared.TieredConversionRateConfigParam],
// [PlanNewParamsPricesPriceNewPlanTieredMatrixWithAllocationPriceConversionRateConfig].
type PlanNewParamsPricesPriceNewPlanTieredMatrixWithAllocationPriceConversionRateConfigUnion interface {
	ImplementsPlanNewParamsPricesPriceNewPlanTieredMatrixWithAllocationPriceConversionRateConfigUnion()
}

type PlanNewParamsPricesPriceNewPlanTieredMatrixWithAllocationPriceConversionRateConfigConversionRateType string

const (
	PlanNewParamsPricesPriceNewPlanTieredMatrixWithAllocationPriceConversionRateConfigConversionRateTypeUnit   PlanNewParamsPricesPriceNewPlanTieredMatrixWithAllocationPriceConversionRateConfigConversionRateType = "unit"
	PlanNewParamsPricesPriceNewPlanTieredMatrixWithAllocationPriceConversionRateConfigConversionRateTypeTiered PlanNewParamsPricesPriceNewPlanTieredMatrixWithAllocationPriceConversionRateConfigConversionRateType = "tiered"
)

func (r PlanNewParamsPricesPriceNewPlanTieredMatrixWithAllocationPriceConversionRateConfigConversionRateType) IsKnown() bool {
	switch r {
	case PlanNewParamsPricesPriceNewPlanTieredMatrixWithAllocationPriceConversionRateConfigConversionRateTypeUnit, PlanNewParamsPricesPriceNewPlanTieredMatrixWithAllocationPriceConversionRateConfigConversionRateTypeTiered:
		return true
	}
	return false
}

type PlanNewParamsPricesPriceNewPlanMatrixWithThresholdDiscountsPrice struct {
	// The cadence to bill for this price on.
	Cadence param.Field[PlanNewParamsPricesPriceNewPlanMatrixWithThresholdDiscountsPriceCadence] `json:"cadence" api:"required"`
	// The id of the item the price will be associated with.
	ItemID param.Field[string] `json:"item_id" api:"required"`
	// Configuration for matrix_with_threshold_discounts pricing
	MatrixWithThresholdDiscountsConfig param.Field[PlanNewParamsPricesPriceNewPlanMatrixWithThresholdDiscountsPriceMatrixWithThresholdDiscountsConfig] `json:"matrix_with_threshold_discounts_config" api:"required"`
	// The pricing model type
	ModelType param.Field[PlanNewParamsPricesPriceNewPlanMatrixWithThresholdDiscountsPriceModelType] `json:"model_type" api:"required"`
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
	ConversionRateConfig param.Field[PlanNewParamsPricesPriceNewPlanMatrixWithThresholdDiscountsPriceConversionRateConfigUnion] `json:"conversion_rate_config"`
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
	// The ID of the license type to associate with this price. On a usage price this
	// also marks the price as eligible to draw down from that license type's
	// allocation; a usage price created without it is billed normally. Usage prices
	// with a license type must use the `unit` model, and only draw down when their
	// currency matches the allocation's.
	LicenseTypeID param.Field[string] `json:"license_type_id"`
	// User-specified key/value pairs for the resource. Individual keys can be removed
	// by setting the value to `null`, and the entire metadata mapping can be cleared
	// by setting `metadata` to `null`.
	Metadata param.Field[map[string]string] `json:"metadata"`
	// A transient ID that can be used to reference this price when adding adjustments
	// in the same API call.
	ReferenceID param.Field[string] `json:"reference_id"`
}

func (r PlanNewParamsPricesPriceNewPlanMatrixWithThresholdDiscountsPrice) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r PlanNewParamsPricesPriceNewPlanMatrixWithThresholdDiscountsPrice) ImplementsPlanNewParamsPricesPriceUnion() {
}

// The cadence to bill for this price on.
type PlanNewParamsPricesPriceNewPlanMatrixWithThresholdDiscountsPriceCadence string

const (
	PlanNewParamsPricesPriceNewPlanMatrixWithThresholdDiscountsPriceCadenceOneTime    PlanNewParamsPricesPriceNewPlanMatrixWithThresholdDiscountsPriceCadence = "one_time"
	PlanNewParamsPricesPriceNewPlanMatrixWithThresholdDiscountsPriceCadenceMonthly    PlanNewParamsPricesPriceNewPlanMatrixWithThresholdDiscountsPriceCadence = "monthly"
	PlanNewParamsPricesPriceNewPlanMatrixWithThresholdDiscountsPriceCadenceQuarterly  PlanNewParamsPricesPriceNewPlanMatrixWithThresholdDiscountsPriceCadence = "quarterly"
	PlanNewParamsPricesPriceNewPlanMatrixWithThresholdDiscountsPriceCadenceSemiAnnual PlanNewParamsPricesPriceNewPlanMatrixWithThresholdDiscountsPriceCadence = "semi_annual"
	PlanNewParamsPricesPriceNewPlanMatrixWithThresholdDiscountsPriceCadenceAnnual     PlanNewParamsPricesPriceNewPlanMatrixWithThresholdDiscountsPriceCadence = "annual"
	PlanNewParamsPricesPriceNewPlanMatrixWithThresholdDiscountsPriceCadenceCustom     PlanNewParamsPricesPriceNewPlanMatrixWithThresholdDiscountsPriceCadence = "custom"
)

func (r PlanNewParamsPricesPriceNewPlanMatrixWithThresholdDiscountsPriceCadence) IsKnown() bool {
	switch r {
	case PlanNewParamsPricesPriceNewPlanMatrixWithThresholdDiscountsPriceCadenceOneTime, PlanNewParamsPricesPriceNewPlanMatrixWithThresholdDiscountsPriceCadenceMonthly, PlanNewParamsPricesPriceNewPlanMatrixWithThresholdDiscountsPriceCadenceQuarterly, PlanNewParamsPricesPriceNewPlanMatrixWithThresholdDiscountsPriceCadenceSemiAnnual, PlanNewParamsPricesPriceNewPlanMatrixWithThresholdDiscountsPriceCadenceAnnual, PlanNewParamsPricesPriceNewPlanMatrixWithThresholdDiscountsPriceCadenceCustom:
		return true
	}
	return false
}

// Configuration for matrix_with_threshold_discounts pricing
type PlanNewParamsPricesPriceNewPlanMatrixWithThresholdDiscountsPriceMatrixWithThresholdDiscountsConfig struct {
	// Unit price used for usage that does not match any defined matrix cell.
	DefaultUnitAmount param.Field[string] `json:"default_unit_amount" api:"required"`
	// First matrix dimension key.
	FirstDimension param.Field[string] `json:"first_dimension" api:"required"`
	// Per-cell unit prices.
	MatrixValues param.Field[[]PlanNewParamsPricesPriceNewPlanMatrixWithThresholdDiscountsPriceMatrixWithThresholdDiscountsConfigMatrixValue] `json:"matrix_values" api:"required"`
	// Optional second matrix dimension key.
	SecondDimension         param.Field[string]                                                                                                                     `json:"second_dimension"`
	ThresholdDiscountGroups param.Field[[]PlanNewParamsPricesPriceNewPlanMatrixWithThresholdDiscountsPriceMatrixWithThresholdDiscountsConfigThresholdDiscountGroup] `json:"threshold_discount_groups"`
}

func (r PlanNewParamsPricesPriceNewPlanMatrixWithThresholdDiscountsPriceMatrixWithThresholdDiscountsConfig) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type PlanNewParamsPricesPriceNewPlanMatrixWithThresholdDiscountsPriceMatrixWithThresholdDiscountsConfigMatrixValue struct {
	FirstDimensionValue  param.Field[string] `json:"first_dimension_value" api:"required"`
	UnitAmount           param.Field[string] `json:"unit_amount" api:"required"`
	SecondDimensionValue param.Field[string] `json:"second_dimension_value"`
}

func (r PlanNewParamsPricesPriceNewPlanMatrixWithThresholdDiscountsPriceMatrixWithThresholdDiscountsConfigMatrixValue) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type PlanNewParamsPricesPriceNewPlanMatrixWithThresholdDiscountsPriceMatrixWithThresholdDiscountsConfigThresholdDiscountGroup struct {
	// Discount rate applied to spend above the threshold.
	AboveThresholdDiscountPercentage param.Field[string] `json:"above_threshold_discount_percentage" api:"required"`
	// Discount rate applied to spend at or below the threshold. Set to 0 for no
	// baseline discount.
	BelowThresholdDiscountPercentage param.Field[string] `json:"below_threshold_discount_percentage" api:"required"`
	// Semicolon-separated list of matrix cell coordinates targeted by this group. Each
	// coordinate is `first,second` when the matrix has two dimensions, or just `first`
	// for a single-dimension matrix. Example: `blue,circle;green,triangle`.
	CellCoordinates param.Field[string] `json:"cell_coordinates" api:"required"`
	ThresholdAmount param.Field[string] `json:"threshold_amount" api:"required"`
	Description     param.Field[string] `json:"description"`
}

func (r PlanNewParamsPricesPriceNewPlanMatrixWithThresholdDiscountsPriceMatrixWithThresholdDiscountsConfigThresholdDiscountGroup) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// The pricing model type
type PlanNewParamsPricesPriceNewPlanMatrixWithThresholdDiscountsPriceModelType string

const (
	PlanNewParamsPricesPriceNewPlanMatrixWithThresholdDiscountsPriceModelTypeMatrixWithThresholdDiscounts PlanNewParamsPricesPriceNewPlanMatrixWithThresholdDiscountsPriceModelType = "matrix_with_threshold_discounts"
)

func (r PlanNewParamsPricesPriceNewPlanMatrixWithThresholdDiscountsPriceModelType) IsKnown() bool {
	switch r {
	case PlanNewParamsPricesPriceNewPlanMatrixWithThresholdDiscountsPriceModelTypeMatrixWithThresholdDiscounts:
		return true
	}
	return false
}

type PlanNewParamsPricesPriceNewPlanMatrixWithThresholdDiscountsPriceConversionRateConfig struct {
	ConversionRateType param.Field[PlanNewParamsPricesPriceNewPlanMatrixWithThresholdDiscountsPriceConversionRateConfigConversionRateType] `json:"conversion_rate_type" api:"required"`
	TieredConfig       param.Field[shared.ConversionRateTieredConfigParam]                                                                 `json:"tiered_config"`
	UnitConfig         param.Field[shared.ConversionRateUnitConfigParam]                                                                   `json:"unit_config"`
}

func (r PlanNewParamsPricesPriceNewPlanMatrixWithThresholdDiscountsPriceConversionRateConfig) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r PlanNewParamsPricesPriceNewPlanMatrixWithThresholdDiscountsPriceConversionRateConfig) ImplementsPlanNewParamsPricesPriceNewPlanMatrixWithThresholdDiscountsPriceConversionRateConfigUnion() {
}

// Satisfied by [shared.UnitConversionRateConfigParam],
// [shared.TieredConversionRateConfigParam],
// [PlanNewParamsPricesPriceNewPlanMatrixWithThresholdDiscountsPriceConversionRateConfig].
type PlanNewParamsPricesPriceNewPlanMatrixWithThresholdDiscountsPriceConversionRateConfigUnion interface {
	ImplementsPlanNewParamsPricesPriceNewPlanMatrixWithThresholdDiscountsPriceConversionRateConfigUnion()
}

type PlanNewParamsPricesPriceNewPlanMatrixWithThresholdDiscountsPriceConversionRateConfigConversionRateType string

const (
	PlanNewParamsPricesPriceNewPlanMatrixWithThresholdDiscountsPriceConversionRateConfigConversionRateTypeUnit   PlanNewParamsPricesPriceNewPlanMatrixWithThresholdDiscountsPriceConversionRateConfigConversionRateType = "unit"
	PlanNewParamsPricesPriceNewPlanMatrixWithThresholdDiscountsPriceConversionRateConfigConversionRateTypeTiered PlanNewParamsPricesPriceNewPlanMatrixWithThresholdDiscountsPriceConversionRateConfigConversionRateType = "tiered"
)

func (r PlanNewParamsPricesPriceNewPlanMatrixWithThresholdDiscountsPriceConversionRateConfigConversionRateType) IsKnown() bool {
	switch r {
	case PlanNewParamsPricesPriceNewPlanMatrixWithThresholdDiscountsPriceConversionRateConfigConversionRateTypeUnit, PlanNewParamsPricesPriceNewPlanMatrixWithThresholdDiscountsPriceConversionRateConfigConversionRateTypeTiered:
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
	// The ID of the license type to associate with this price. On a usage price this
	// also marks the price as eligible to draw down from that license type's
	// allocation; a usage price created without it is billed normally. Usage prices
	// with a license type must use the `unit` model, and only draw down when their
	// currency matches the allocation's.
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
	PlanNewParamsPricesPriceNewPlanTieredWithProrationPriceCadenceOneTime    PlanNewParamsPricesPriceNewPlanTieredWithProrationPriceCadence = "one_time"
	PlanNewParamsPricesPriceNewPlanTieredWithProrationPriceCadenceMonthly    PlanNewParamsPricesPriceNewPlanTieredWithProrationPriceCadence = "monthly"
	PlanNewParamsPricesPriceNewPlanTieredWithProrationPriceCadenceQuarterly  PlanNewParamsPricesPriceNewPlanTieredWithProrationPriceCadence = "quarterly"
	PlanNewParamsPricesPriceNewPlanTieredWithProrationPriceCadenceSemiAnnual PlanNewParamsPricesPriceNewPlanTieredWithProrationPriceCadence = "semi_annual"
	PlanNewParamsPricesPriceNewPlanTieredWithProrationPriceCadenceAnnual     PlanNewParamsPricesPriceNewPlanTieredWithProrationPriceCadence = "annual"
	PlanNewParamsPricesPriceNewPlanTieredWithProrationPriceCadenceCustom     PlanNewParamsPricesPriceNewPlanTieredWithProrationPriceCadence = "custom"
)

func (r PlanNewParamsPricesPriceNewPlanTieredWithProrationPriceCadence) IsKnown() bool {
	switch r {
	case PlanNewParamsPricesPriceNewPlanTieredWithProrationPriceCadenceOneTime, PlanNewParamsPricesPriceNewPlanTieredWithProrationPriceCadenceMonthly, PlanNewParamsPricesPriceNewPlanTieredWithProrationPriceCadenceQuarterly, PlanNewParamsPricesPriceNewPlanTieredWithProrationPriceCadenceSemiAnnual, PlanNewParamsPricesPriceNewPlanTieredWithProrationPriceCadenceAnnual, PlanNewParamsPricesPriceNewPlanTieredWithProrationPriceCadenceCustom:
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
	// The ID of the license type to associate with this price. On a usage price this
	// also marks the price as eligible to draw down from that license type's
	// allocation; a usage price created without it is billed normally. Usage prices
	// with a license type must use the `unit` model, and only draw down when their
	// currency matches the allocation's.
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
	PlanNewParamsPricesPriceNewPlanGroupedWithMinMaxThresholdsPriceCadenceOneTime    PlanNewParamsPricesPriceNewPlanGroupedWithMinMaxThresholdsPriceCadence = "one_time"
	PlanNewParamsPricesPriceNewPlanGroupedWithMinMaxThresholdsPriceCadenceMonthly    PlanNewParamsPricesPriceNewPlanGroupedWithMinMaxThresholdsPriceCadence = "monthly"
	PlanNewParamsPricesPriceNewPlanGroupedWithMinMaxThresholdsPriceCadenceQuarterly  PlanNewParamsPricesPriceNewPlanGroupedWithMinMaxThresholdsPriceCadence = "quarterly"
	PlanNewParamsPricesPriceNewPlanGroupedWithMinMaxThresholdsPriceCadenceSemiAnnual PlanNewParamsPricesPriceNewPlanGroupedWithMinMaxThresholdsPriceCadence = "semi_annual"
	PlanNewParamsPricesPriceNewPlanGroupedWithMinMaxThresholdsPriceCadenceAnnual     PlanNewParamsPricesPriceNewPlanGroupedWithMinMaxThresholdsPriceCadence = "annual"
	PlanNewParamsPricesPriceNewPlanGroupedWithMinMaxThresholdsPriceCadenceCustom     PlanNewParamsPricesPriceNewPlanGroupedWithMinMaxThresholdsPriceCadence = "custom"
)

func (r PlanNewParamsPricesPriceNewPlanGroupedWithMinMaxThresholdsPriceCadence) IsKnown() bool {
	switch r {
	case PlanNewParamsPricesPriceNewPlanGroupedWithMinMaxThresholdsPriceCadenceOneTime, PlanNewParamsPricesPriceNewPlanGroupedWithMinMaxThresholdsPriceCadenceMonthly, PlanNewParamsPricesPriceNewPlanGroupedWithMinMaxThresholdsPriceCadenceQuarterly, PlanNewParamsPricesPriceNewPlanGroupedWithMinMaxThresholdsPriceCadenceSemiAnnual, PlanNewParamsPricesPriceNewPlanGroupedWithMinMaxThresholdsPriceCadenceAnnual, PlanNewParamsPricesPriceNewPlanGroupedWithMinMaxThresholdsPriceCadenceCustom:
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
	// The ID of the license type to associate with this price. On a usage price this
	// also marks the price as eligible to draw down from that license type's
	// allocation; a usage price created without it is billed normally. Usage prices
	// with a license type must use the `unit` model, and only draw down when their
	// currency matches the allocation's.
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
	PlanNewParamsPricesPriceNewPlanCumulativeGroupedAllocationPriceCadenceOneTime    PlanNewParamsPricesPriceNewPlanCumulativeGroupedAllocationPriceCadence = "one_time"
	PlanNewParamsPricesPriceNewPlanCumulativeGroupedAllocationPriceCadenceMonthly    PlanNewParamsPricesPriceNewPlanCumulativeGroupedAllocationPriceCadence = "monthly"
	PlanNewParamsPricesPriceNewPlanCumulativeGroupedAllocationPriceCadenceQuarterly  PlanNewParamsPricesPriceNewPlanCumulativeGroupedAllocationPriceCadence = "quarterly"
	PlanNewParamsPricesPriceNewPlanCumulativeGroupedAllocationPriceCadenceSemiAnnual PlanNewParamsPricesPriceNewPlanCumulativeGroupedAllocationPriceCadence = "semi_annual"
	PlanNewParamsPricesPriceNewPlanCumulativeGroupedAllocationPriceCadenceAnnual     PlanNewParamsPricesPriceNewPlanCumulativeGroupedAllocationPriceCadence = "annual"
	PlanNewParamsPricesPriceNewPlanCumulativeGroupedAllocationPriceCadenceCustom     PlanNewParamsPricesPriceNewPlanCumulativeGroupedAllocationPriceCadence = "custom"
)

func (r PlanNewParamsPricesPriceNewPlanCumulativeGroupedAllocationPriceCadence) IsKnown() bool {
	switch r {
	case PlanNewParamsPricesPriceNewPlanCumulativeGroupedAllocationPriceCadenceOneTime, PlanNewParamsPricesPriceNewPlanCumulativeGroupedAllocationPriceCadenceMonthly, PlanNewParamsPricesPriceNewPlanCumulativeGroupedAllocationPriceCadenceQuarterly, PlanNewParamsPricesPriceNewPlanCumulativeGroupedAllocationPriceCadenceSemiAnnual, PlanNewParamsPricesPriceNewPlanCumulativeGroupedAllocationPriceCadenceAnnual, PlanNewParamsPricesPriceNewPlanCumulativeGroupedAllocationPriceCadenceCustom:
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

type PlanNewParamsPricesPriceNewPlanDailyCreditAllowancePrice struct {
	// The cadence to bill for this price on.
	Cadence param.Field[PlanNewParamsPricesPriceNewPlanDailyCreditAllowancePriceCadence] `json:"cadence" api:"required"`
	// Configuration for daily_credit_allowance pricing
	DailyCreditAllowanceConfig param.Field[PlanNewParamsPricesPriceNewPlanDailyCreditAllowancePriceDailyCreditAllowanceConfig] `json:"daily_credit_allowance_config" api:"required"`
	// The id of the item the price will be associated with.
	ItemID param.Field[string] `json:"item_id" api:"required"`
	// The pricing model type
	ModelType param.Field[PlanNewParamsPricesPriceNewPlanDailyCreditAllowancePriceModelType] `json:"model_type" api:"required"`
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
	ConversionRateConfig param.Field[PlanNewParamsPricesPriceNewPlanDailyCreditAllowancePriceConversionRateConfigUnion] `json:"conversion_rate_config"`
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
	// The ID of the license type to associate with this price. On a usage price this
	// also marks the price as eligible to draw down from that license type's
	// allocation; a usage price created without it is billed normally. Usage prices
	// with a license type must use the `unit` model, and only draw down when their
	// currency matches the allocation's.
	LicenseTypeID param.Field[string] `json:"license_type_id"`
	// User-specified key/value pairs for the resource. Individual keys can be removed
	// by setting the value to `null`, and the entire metadata mapping can be cleared
	// by setting `metadata` to `null`.
	Metadata param.Field[map[string]string] `json:"metadata"`
	// A transient ID that can be used to reference this price when adding adjustments
	// in the same API call.
	ReferenceID param.Field[string] `json:"reference_id"`
}

func (r PlanNewParamsPricesPriceNewPlanDailyCreditAllowancePrice) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r PlanNewParamsPricesPriceNewPlanDailyCreditAllowancePrice) ImplementsPlanNewParamsPricesPriceUnion() {
}

// The cadence to bill for this price on.
type PlanNewParamsPricesPriceNewPlanDailyCreditAllowancePriceCadence string

const (
	PlanNewParamsPricesPriceNewPlanDailyCreditAllowancePriceCadenceOneTime    PlanNewParamsPricesPriceNewPlanDailyCreditAllowancePriceCadence = "one_time"
	PlanNewParamsPricesPriceNewPlanDailyCreditAllowancePriceCadenceMonthly    PlanNewParamsPricesPriceNewPlanDailyCreditAllowancePriceCadence = "monthly"
	PlanNewParamsPricesPriceNewPlanDailyCreditAllowancePriceCadenceQuarterly  PlanNewParamsPricesPriceNewPlanDailyCreditAllowancePriceCadence = "quarterly"
	PlanNewParamsPricesPriceNewPlanDailyCreditAllowancePriceCadenceSemiAnnual PlanNewParamsPricesPriceNewPlanDailyCreditAllowancePriceCadence = "semi_annual"
	PlanNewParamsPricesPriceNewPlanDailyCreditAllowancePriceCadenceAnnual     PlanNewParamsPricesPriceNewPlanDailyCreditAllowancePriceCadence = "annual"
	PlanNewParamsPricesPriceNewPlanDailyCreditAllowancePriceCadenceCustom     PlanNewParamsPricesPriceNewPlanDailyCreditAllowancePriceCadence = "custom"
)

func (r PlanNewParamsPricesPriceNewPlanDailyCreditAllowancePriceCadence) IsKnown() bool {
	switch r {
	case PlanNewParamsPricesPriceNewPlanDailyCreditAllowancePriceCadenceOneTime, PlanNewParamsPricesPriceNewPlanDailyCreditAllowancePriceCadenceMonthly, PlanNewParamsPricesPriceNewPlanDailyCreditAllowancePriceCadenceQuarterly, PlanNewParamsPricesPriceNewPlanDailyCreditAllowancePriceCadenceSemiAnnual, PlanNewParamsPricesPriceNewPlanDailyCreditAllowancePriceCadenceAnnual, PlanNewParamsPricesPriceNewPlanDailyCreditAllowancePriceCadenceCustom:
		return true
	}
	return false
}

// Configuration for daily_credit_allowance pricing
type PlanNewParamsPricesPriceNewPlanDailyCreditAllowancePriceDailyCreditAllowanceConfig struct {
	// Credits granted per day. Lose-it-or-use-it; does not roll over.
	DailyAllowance param.Field[string] `json:"daily_allowance" api:"required"`
	// Default per-unit credit rate for any usage not bucketed into a specified
	// matrix_value
	DefaultUnitAmount param.Field[string] `json:"default_unit_amount" api:"required"`
	// One or two event property values to evaluate matrix groups by
	Dimensions param.Field[[]string] `json:"dimensions" api:"required"`
	// Event property whose value identifies the day bucket the event belongs to (e.g.
	// 'event_day' set to an ISO date string in the customer's timezone). The allowance
	// resets per distinct value of this property.
	EventDayProperty param.Field[string] `json:"event_day_property" api:"required"`
	// Per-dimension credit rates
	MatrixValues param.Field[[]PlanNewParamsPricesPriceNewPlanDailyCreditAllowancePriceDailyCreditAllowanceConfigMatrixValue] `json:"matrix_values" api:"required"`
}

func (r PlanNewParamsPricesPriceNewPlanDailyCreditAllowancePriceDailyCreditAllowanceConfig) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Per-dimension credit price for the daily credit allowance model.
type PlanNewParamsPricesPriceNewPlanDailyCreditAllowancePriceDailyCreditAllowanceConfigMatrixValue struct {
	// One or two matrix keys to filter usage to this value by. For example, ["model"]
	// could be used to apply a different credit rate to each AI model.
	DimensionValues param.Field[[]string] `json:"dimension_values" api:"required"`
	// Credits charged per unit of usage matching the specified dimension_values
	UnitAmount param.Field[string] `json:"unit_amount" api:"required"`
}

func (r PlanNewParamsPricesPriceNewPlanDailyCreditAllowancePriceDailyCreditAllowanceConfigMatrixValue) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// The pricing model type
type PlanNewParamsPricesPriceNewPlanDailyCreditAllowancePriceModelType string

const (
	PlanNewParamsPricesPriceNewPlanDailyCreditAllowancePriceModelTypeDailyCreditAllowance PlanNewParamsPricesPriceNewPlanDailyCreditAllowancePriceModelType = "daily_credit_allowance"
)

func (r PlanNewParamsPricesPriceNewPlanDailyCreditAllowancePriceModelType) IsKnown() bool {
	switch r {
	case PlanNewParamsPricesPriceNewPlanDailyCreditAllowancePriceModelTypeDailyCreditAllowance:
		return true
	}
	return false
}

type PlanNewParamsPricesPriceNewPlanDailyCreditAllowancePriceConversionRateConfig struct {
	ConversionRateType param.Field[PlanNewParamsPricesPriceNewPlanDailyCreditAllowancePriceConversionRateConfigConversionRateType] `json:"conversion_rate_type" api:"required"`
	TieredConfig       param.Field[shared.ConversionRateTieredConfigParam]                                                         `json:"tiered_config"`
	UnitConfig         param.Field[shared.ConversionRateUnitConfigParam]                                                           `json:"unit_config"`
}

func (r PlanNewParamsPricesPriceNewPlanDailyCreditAllowancePriceConversionRateConfig) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r PlanNewParamsPricesPriceNewPlanDailyCreditAllowancePriceConversionRateConfig) ImplementsPlanNewParamsPricesPriceNewPlanDailyCreditAllowancePriceConversionRateConfigUnion() {
}

// Satisfied by [shared.UnitConversionRateConfigParam],
// [shared.TieredConversionRateConfigParam],
// [PlanNewParamsPricesPriceNewPlanDailyCreditAllowancePriceConversionRateConfig].
type PlanNewParamsPricesPriceNewPlanDailyCreditAllowancePriceConversionRateConfigUnion interface {
	ImplementsPlanNewParamsPricesPriceNewPlanDailyCreditAllowancePriceConversionRateConfigUnion()
}

type PlanNewParamsPricesPriceNewPlanDailyCreditAllowancePriceConversionRateConfigConversionRateType string

const (
	PlanNewParamsPricesPriceNewPlanDailyCreditAllowancePriceConversionRateConfigConversionRateTypeUnit   PlanNewParamsPricesPriceNewPlanDailyCreditAllowancePriceConversionRateConfigConversionRateType = "unit"
	PlanNewParamsPricesPriceNewPlanDailyCreditAllowancePriceConversionRateConfigConversionRateTypeTiered PlanNewParamsPricesPriceNewPlanDailyCreditAllowancePriceConversionRateConfigConversionRateType = "tiered"
)

func (r PlanNewParamsPricesPriceNewPlanDailyCreditAllowancePriceConversionRateConfigConversionRateType) IsKnown() bool {
	switch r {
	case PlanNewParamsPricesPriceNewPlanDailyCreditAllowancePriceConversionRateConfigConversionRateTypeUnit, PlanNewParamsPricesPriceNewPlanDailyCreditAllowancePriceConversionRateConfigConversionRateTypeTiered:
		return true
	}
	return false
}

type PlanNewParamsPricesPriceNewPlanMeteredAllowancePrice struct {
	// The cadence to bill for this price on.
	Cadence param.Field[PlanNewParamsPricesPriceNewPlanMeteredAllowancePriceCadence] `json:"cadence" api:"required"`
	// The id of the item the price will be associated with.
	ItemID param.Field[string] `json:"item_id" api:"required"`
	// Configuration for metered_allowance pricing
	MeteredAllowanceConfig param.Field[PlanNewParamsPricesPriceNewPlanMeteredAllowancePriceMeteredAllowanceConfig] `json:"metered_allowance_config" api:"required"`
	// The pricing model type
	ModelType param.Field[PlanNewParamsPricesPriceNewPlanMeteredAllowancePriceModelType] `json:"model_type" api:"required"`
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
	ConversionRateConfig param.Field[PlanNewParamsPricesPriceNewPlanMeteredAllowancePriceConversionRateConfigUnion] `json:"conversion_rate_config"`
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
	// The ID of the license type to associate with this price. On a usage price this
	// also marks the price as eligible to draw down from that license type's
	// allocation; a usage price created without it is billed normally. Usage prices
	// with a license type must use the `unit` model, and only draw down when their
	// currency matches the allocation's.
	LicenseTypeID param.Field[string] `json:"license_type_id"`
	// User-specified key/value pairs for the resource. Individual keys can be removed
	// by setting the value to `null`, and the entire metadata mapping can be cleared
	// by setting `metadata` to `null`.
	Metadata param.Field[map[string]string] `json:"metadata"`
	// A transient ID that can be used to reference this price when adding adjustments
	// in the same API call.
	ReferenceID param.Field[string] `json:"reference_id"`
}

func (r PlanNewParamsPricesPriceNewPlanMeteredAllowancePrice) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r PlanNewParamsPricesPriceNewPlanMeteredAllowancePrice) ImplementsPlanNewParamsPricesPriceUnion() {
}

// The cadence to bill for this price on.
type PlanNewParamsPricesPriceNewPlanMeteredAllowancePriceCadence string

const (
	PlanNewParamsPricesPriceNewPlanMeteredAllowancePriceCadenceOneTime    PlanNewParamsPricesPriceNewPlanMeteredAllowancePriceCadence = "one_time"
	PlanNewParamsPricesPriceNewPlanMeteredAllowancePriceCadenceMonthly    PlanNewParamsPricesPriceNewPlanMeteredAllowancePriceCadence = "monthly"
	PlanNewParamsPricesPriceNewPlanMeteredAllowancePriceCadenceQuarterly  PlanNewParamsPricesPriceNewPlanMeteredAllowancePriceCadence = "quarterly"
	PlanNewParamsPricesPriceNewPlanMeteredAllowancePriceCadenceSemiAnnual PlanNewParamsPricesPriceNewPlanMeteredAllowancePriceCadence = "semi_annual"
	PlanNewParamsPricesPriceNewPlanMeteredAllowancePriceCadenceAnnual     PlanNewParamsPricesPriceNewPlanMeteredAllowancePriceCadence = "annual"
	PlanNewParamsPricesPriceNewPlanMeteredAllowancePriceCadenceCustom     PlanNewParamsPricesPriceNewPlanMeteredAllowancePriceCadence = "custom"
)

func (r PlanNewParamsPricesPriceNewPlanMeteredAllowancePriceCadence) IsKnown() bool {
	switch r {
	case PlanNewParamsPricesPriceNewPlanMeteredAllowancePriceCadenceOneTime, PlanNewParamsPricesPriceNewPlanMeteredAllowancePriceCadenceMonthly, PlanNewParamsPricesPriceNewPlanMeteredAllowancePriceCadenceQuarterly, PlanNewParamsPricesPriceNewPlanMeteredAllowancePriceCadenceSemiAnnual, PlanNewParamsPricesPriceNewPlanMeteredAllowancePriceCadenceAnnual, PlanNewParamsPricesPriceNewPlanMeteredAllowancePriceCadenceCustom:
		return true
	}
	return false
}

// Configuration for metered_allowance pricing
type PlanNewParamsPricesPriceNewPlanMeteredAllowancePriceMeteredAllowanceConfig struct {
	// The grouping_key value whose summed quantity represents the allowance for this
	// period (e.g. 'storage_snapshot' emitting 3 × avg storage). Capped at consumption
	// — credit can never exceed actual usage.
	AllowanceGroupingValue param.Field[string] `json:"allowance_grouping_value" api:"required"`
	// The grouping_key value whose summed quantity represents consumption (e.g.
	// 'download'). Charged at unit_amount.
	ConsumptionGroupingValue param.Field[string] `json:"consumption_grouping_value" api:"required"`
	// Event property used to partition the metric into consumption and allowance
	// quantities (e.g. 'event_name'). The metric is queried with this key and the two
	// values below select which partition is which.
	GroupingKey param.Field[string] `json:"grouping_key" api:"required"`
	// Per-unit price applied to gross consumption and to the allowance credit.
	UnitAmount param.Field[string] `json:"unit_amount" api:"required"`
	// Sub-line label for the credit row (e.g. 'Up to 3x free egress').
	AllowanceDisplayName param.Field[string] `json:"allowance_display_name"`
	// Sub-line label for the gross consumption row (e.g. 'bytes gotten').
	ConsumptionDisplayName param.Field[string] `json:"consumption_display_name"`
}

func (r PlanNewParamsPricesPriceNewPlanMeteredAllowancePriceMeteredAllowanceConfig) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// The pricing model type
type PlanNewParamsPricesPriceNewPlanMeteredAllowancePriceModelType string

const (
	PlanNewParamsPricesPriceNewPlanMeteredAllowancePriceModelTypeMeteredAllowance PlanNewParamsPricesPriceNewPlanMeteredAllowancePriceModelType = "metered_allowance"
)

func (r PlanNewParamsPricesPriceNewPlanMeteredAllowancePriceModelType) IsKnown() bool {
	switch r {
	case PlanNewParamsPricesPriceNewPlanMeteredAllowancePriceModelTypeMeteredAllowance:
		return true
	}
	return false
}

type PlanNewParamsPricesPriceNewPlanMeteredAllowancePriceConversionRateConfig struct {
	ConversionRateType param.Field[PlanNewParamsPricesPriceNewPlanMeteredAllowancePriceConversionRateConfigConversionRateType] `json:"conversion_rate_type" api:"required"`
	TieredConfig       param.Field[shared.ConversionRateTieredConfigParam]                                                     `json:"tiered_config"`
	UnitConfig         param.Field[shared.ConversionRateUnitConfigParam]                                                       `json:"unit_config"`
}

func (r PlanNewParamsPricesPriceNewPlanMeteredAllowancePriceConversionRateConfig) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r PlanNewParamsPricesPriceNewPlanMeteredAllowancePriceConversionRateConfig) ImplementsPlanNewParamsPricesPriceNewPlanMeteredAllowancePriceConversionRateConfigUnion() {
}

// Satisfied by [shared.UnitConversionRateConfigParam],
// [shared.TieredConversionRateConfigParam],
// [PlanNewParamsPricesPriceNewPlanMeteredAllowancePriceConversionRateConfig].
type PlanNewParamsPricesPriceNewPlanMeteredAllowancePriceConversionRateConfigUnion interface {
	ImplementsPlanNewParamsPricesPriceNewPlanMeteredAllowancePriceConversionRateConfigUnion()
}

type PlanNewParamsPricesPriceNewPlanMeteredAllowancePriceConversionRateConfigConversionRateType string

const (
	PlanNewParamsPricesPriceNewPlanMeteredAllowancePriceConversionRateConfigConversionRateTypeUnit   PlanNewParamsPricesPriceNewPlanMeteredAllowancePriceConversionRateConfigConversionRateType = "unit"
	PlanNewParamsPricesPriceNewPlanMeteredAllowancePriceConversionRateConfigConversionRateTypeTiered PlanNewParamsPricesPriceNewPlanMeteredAllowancePriceConversionRateConfigConversionRateType = "tiered"
)

func (r PlanNewParamsPricesPriceNewPlanMeteredAllowancePriceConversionRateConfigConversionRateType) IsKnown() bool {
	switch r {
	case PlanNewParamsPricesPriceNewPlanMeteredAllowancePriceConversionRateConfigConversionRateTypeUnit, PlanNewParamsPricesPriceNewPlanMeteredAllowancePriceConversionRateConfigConversionRateTypeTiered:
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
	// The ID of the license type to associate with this price. On a usage price this
	// also marks the price as eligible to draw down from that license type's
	// allocation; a usage price created without it is billed normally. Usage prices
	// with a license type must use the `unit` model, and only draw down when their
	// currency matches the allocation's.
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
	PlanNewParamsPricesPriceNewPlanPercentCompositePriceCadenceOneTime    PlanNewParamsPricesPriceNewPlanPercentCompositePriceCadence = "one_time"
	PlanNewParamsPricesPriceNewPlanPercentCompositePriceCadenceMonthly    PlanNewParamsPricesPriceNewPlanPercentCompositePriceCadence = "monthly"
	PlanNewParamsPricesPriceNewPlanPercentCompositePriceCadenceQuarterly  PlanNewParamsPricesPriceNewPlanPercentCompositePriceCadence = "quarterly"
	PlanNewParamsPricesPriceNewPlanPercentCompositePriceCadenceSemiAnnual PlanNewParamsPricesPriceNewPlanPercentCompositePriceCadence = "semi_annual"
	PlanNewParamsPricesPriceNewPlanPercentCompositePriceCadenceAnnual     PlanNewParamsPricesPriceNewPlanPercentCompositePriceCadence = "annual"
	PlanNewParamsPricesPriceNewPlanPercentCompositePriceCadenceCustom     PlanNewParamsPricesPriceNewPlanPercentCompositePriceCadence = "custom"
)

func (r PlanNewParamsPricesPriceNewPlanPercentCompositePriceCadence) IsKnown() bool {
	switch r {
	case PlanNewParamsPricesPriceNewPlanPercentCompositePriceCadenceOneTime, PlanNewParamsPricesPriceNewPlanPercentCompositePriceCadenceMonthly, PlanNewParamsPricesPriceNewPlanPercentCompositePriceCadenceQuarterly, PlanNewParamsPricesPriceNewPlanPercentCompositePriceCadenceSemiAnnual, PlanNewParamsPricesPriceNewPlanPercentCompositePriceCadenceAnnual, PlanNewParamsPricesPriceNewPlanPercentCompositePriceCadenceCustom:
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
	// Fraction of the component subtotals to charge (0 < percent <= 1).
	Percent param.Field[float64] `json:"percent" api:"required"`
	// Maximum amount to charge. If unset, the fee has no upper bound.
	MaximumAmount param.Field[string] `json:"maximum_amount"`
	// Minimum amount to charge. If unset, the fee is bounded below by 0.
	MinimumAmount param.Field[string] `json:"minimum_amount"`
	// If true, the minimum_amount is prorated based on the service period. The
	// maximum_amount is an absolute cap (never prorated), and the percent applied to
	// upstream subtotals is never prorated either.
	Prorated param.Field[bool] `json:"prorated"`
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
	// The ID of the license type to associate with this price. On a usage price this
	// also marks the price as eligible to draw down from that license type's
	// allocation; a usage price created without it is billed normally. Usage prices
	// with a license type must use the `unit` model, and only draw down when their
	// currency matches the allocation's.
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
	PlanNewParamsPricesPriceNewPlanEventOutputPriceCadenceOneTime    PlanNewParamsPricesPriceNewPlanEventOutputPriceCadence = "one_time"
	PlanNewParamsPricesPriceNewPlanEventOutputPriceCadenceMonthly    PlanNewParamsPricesPriceNewPlanEventOutputPriceCadence = "monthly"
	PlanNewParamsPricesPriceNewPlanEventOutputPriceCadenceQuarterly  PlanNewParamsPricesPriceNewPlanEventOutputPriceCadence = "quarterly"
	PlanNewParamsPricesPriceNewPlanEventOutputPriceCadenceSemiAnnual PlanNewParamsPricesPriceNewPlanEventOutputPriceCadence = "semi_annual"
	PlanNewParamsPricesPriceNewPlanEventOutputPriceCadenceAnnual     PlanNewParamsPricesPriceNewPlanEventOutputPriceCadence = "annual"
	PlanNewParamsPricesPriceNewPlanEventOutputPriceCadenceCustom     PlanNewParamsPricesPriceNewPlanEventOutputPriceCadence = "custom"
)

func (r PlanNewParamsPricesPriceNewPlanEventOutputPriceCadence) IsKnown() bool {
	switch r {
	case PlanNewParamsPricesPriceNewPlanEventOutputPriceCadenceOneTime, PlanNewParamsPricesPriceNewPlanEventOutputPriceCadenceMonthly, PlanNewParamsPricesPriceNewPlanEventOutputPriceCadenceQuarterly, PlanNewParamsPricesPriceNewPlanEventOutputPriceCadenceSemiAnnual, PlanNewParamsPricesPriceNewPlanEventOutputPriceCadenceAnnual, PlanNewParamsPricesPriceNewPlanEventOutputPriceCadenceCustom:
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

type PlanNewParamsPricesPriceNewPlanUnitWithMultiplierPrice struct {
	// The cadence to bill for this price on.
	Cadence param.Field[PlanNewParamsPricesPriceNewPlanUnitWithMultiplierPriceCadence] `json:"cadence" api:"required"`
	// The id of the item the price will be associated with.
	ItemID param.Field[string] `json:"item_id" api:"required"`
	// The pricing model type
	ModelType param.Field[PlanNewParamsPricesPriceNewPlanUnitWithMultiplierPriceModelType] `json:"model_type" api:"required"`
	// The name of the price.
	Name param.Field[string] `json:"name" api:"required"`
	// Configuration for unit_with_multiplier pricing
	UnitWithMultiplierConfig param.Field[PlanNewParamsPricesPriceNewPlanUnitWithMultiplierPriceUnitWithMultiplierConfig] `json:"unit_with_multiplier_config" api:"required"`
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
	ConversionRateConfig param.Field[PlanNewParamsPricesPriceNewPlanUnitWithMultiplierPriceConversionRateConfigUnion] `json:"conversion_rate_config"`
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
	// The ID of the license type to associate with this price. On a usage price this
	// also marks the price as eligible to draw down from that license type's
	// allocation; a usage price created without it is billed normally. Usage prices
	// with a license type must use the `unit` model, and only draw down when their
	// currency matches the allocation's.
	LicenseTypeID param.Field[string] `json:"license_type_id"`
	// User-specified key/value pairs for the resource. Individual keys can be removed
	// by setting the value to `null`, and the entire metadata mapping can be cleared
	// by setting `metadata` to `null`.
	Metadata param.Field[map[string]string] `json:"metadata"`
	// A transient ID that can be used to reference this price when adding adjustments
	// in the same API call.
	ReferenceID param.Field[string] `json:"reference_id"`
}

func (r PlanNewParamsPricesPriceNewPlanUnitWithMultiplierPrice) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r PlanNewParamsPricesPriceNewPlanUnitWithMultiplierPrice) ImplementsPlanNewParamsPricesPriceUnion() {
}

// The cadence to bill for this price on.
type PlanNewParamsPricesPriceNewPlanUnitWithMultiplierPriceCadence string

const (
	PlanNewParamsPricesPriceNewPlanUnitWithMultiplierPriceCadenceOneTime    PlanNewParamsPricesPriceNewPlanUnitWithMultiplierPriceCadence = "one_time"
	PlanNewParamsPricesPriceNewPlanUnitWithMultiplierPriceCadenceMonthly    PlanNewParamsPricesPriceNewPlanUnitWithMultiplierPriceCadence = "monthly"
	PlanNewParamsPricesPriceNewPlanUnitWithMultiplierPriceCadenceQuarterly  PlanNewParamsPricesPriceNewPlanUnitWithMultiplierPriceCadence = "quarterly"
	PlanNewParamsPricesPriceNewPlanUnitWithMultiplierPriceCadenceSemiAnnual PlanNewParamsPricesPriceNewPlanUnitWithMultiplierPriceCadence = "semi_annual"
	PlanNewParamsPricesPriceNewPlanUnitWithMultiplierPriceCadenceAnnual     PlanNewParamsPricesPriceNewPlanUnitWithMultiplierPriceCadence = "annual"
	PlanNewParamsPricesPriceNewPlanUnitWithMultiplierPriceCadenceCustom     PlanNewParamsPricesPriceNewPlanUnitWithMultiplierPriceCadence = "custom"
)

func (r PlanNewParamsPricesPriceNewPlanUnitWithMultiplierPriceCadence) IsKnown() bool {
	switch r {
	case PlanNewParamsPricesPriceNewPlanUnitWithMultiplierPriceCadenceOneTime, PlanNewParamsPricesPriceNewPlanUnitWithMultiplierPriceCadenceMonthly, PlanNewParamsPricesPriceNewPlanUnitWithMultiplierPriceCadenceQuarterly, PlanNewParamsPricesPriceNewPlanUnitWithMultiplierPriceCadenceSemiAnnual, PlanNewParamsPricesPriceNewPlanUnitWithMultiplierPriceCadenceAnnual, PlanNewParamsPricesPriceNewPlanUnitWithMultiplierPriceCadenceCustom:
		return true
	}
	return false
}

// The pricing model type
type PlanNewParamsPricesPriceNewPlanUnitWithMultiplierPriceModelType string

const (
	PlanNewParamsPricesPriceNewPlanUnitWithMultiplierPriceModelTypeUnitWithMultiplier PlanNewParamsPricesPriceNewPlanUnitWithMultiplierPriceModelType = "unit_with_multiplier"
)

func (r PlanNewParamsPricesPriceNewPlanUnitWithMultiplierPriceModelType) IsKnown() bool {
	switch r {
	case PlanNewParamsPricesPriceNewPlanUnitWithMultiplierPriceModelTypeUnitWithMultiplier:
		return true
	}
	return false
}

// Configuration for unit_with_multiplier pricing
type PlanNewParamsPricesPriceNewPlanUnitWithMultiplierPriceUnitWithMultiplierConfig struct {
	// Sub-line label for list-price events that don't set display_name_key. Unlabeled
	// events with another multiplier are labeled by their multiplier (e.g. '0.8×
	// rate').
	DefaultDisplayName param.Field[string] `json:"default_display_name" api:"required"`
	// Event property holding the sub-line label for the event's rate (e.g.
	// 'Off-peak').
	DisplayNameKey param.Field[string] `json:"display_name_key" api:"required"`
	// Event property holding the multiplier applied to unit_amount (e.g. 0.8 for a 20%
	// discount). Events without a number between 0 and 1000 with at most 12 decimal
	// places are charged at unit_amount.
	MultiplierKey param.Field[string] `json:"multiplier_key" api:"required"`
	// List price per unit, before the event's multiplier is applied.
	UnitAmount param.Field[string] `json:"unit_amount" api:"required"`
}

func (r PlanNewParamsPricesPriceNewPlanUnitWithMultiplierPriceUnitWithMultiplierConfig) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type PlanNewParamsPricesPriceNewPlanUnitWithMultiplierPriceConversionRateConfig struct {
	ConversionRateType param.Field[PlanNewParamsPricesPriceNewPlanUnitWithMultiplierPriceConversionRateConfigConversionRateType] `json:"conversion_rate_type" api:"required"`
	TieredConfig       param.Field[shared.ConversionRateTieredConfigParam]                                                       `json:"tiered_config"`
	UnitConfig         param.Field[shared.ConversionRateUnitConfigParam]                                                         `json:"unit_config"`
}

func (r PlanNewParamsPricesPriceNewPlanUnitWithMultiplierPriceConversionRateConfig) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r PlanNewParamsPricesPriceNewPlanUnitWithMultiplierPriceConversionRateConfig) ImplementsPlanNewParamsPricesPriceNewPlanUnitWithMultiplierPriceConversionRateConfigUnion() {
}

// Satisfied by [shared.UnitConversionRateConfigParam],
// [shared.TieredConversionRateConfigParam],
// [PlanNewParamsPricesPriceNewPlanUnitWithMultiplierPriceConversionRateConfig].
type PlanNewParamsPricesPriceNewPlanUnitWithMultiplierPriceConversionRateConfigUnion interface {
	ImplementsPlanNewParamsPricesPriceNewPlanUnitWithMultiplierPriceConversionRateConfigUnion()
}

type PlanNewParamsPricesPriceNewPlanUnitWithMultiplierPriceConversionRateConfigConversionRateType string

const (
	PlanNewParamsPricesPriceNewPlanUnitWithMultiplierPriceConversionRateConfigConversionRateTypeUnit   PlanNewParamsPricesPriceNewPlanUnitWithMultiplierPriceConversionRateConfigConversionRateType = "unit"
	PlanNewParamsPricesPriceNewPlanUnitWithMultiplierPriceConversionRateConfigConversionRateTypeTiered PlanNewParamsPricesPriceNewPlanUnitWithMultiplierPriceConversionRateConfigConversionRateType = "tiered"
)

func (r PlanNewParamsPricesPriceNewPlanUnitWithMultiplierPriceConversionRateConfigConversionRateType) IsKnown() bool {
	switch r {
	case PlanNewParamsPricesPriceNewPlanUnitWithMultiplierPriceConversionRateConfigConversionRateTypeUnit, PlanNewParamsPricesPriceNewPlanUnitWithMultiplierPriceConversionRateConfigConversionRateTypeTiered:
		return true
	}
	return false
}

// The cadence to bill for this price on.
type PlanNewParamsPricesPriceCadence string

const (
	PlanNewParamsPricesPriceCadenceOneTime    PlanNewParamsPricesPriceCadence = "one_time"
	PlanNewParamsPricesPriceCadenceMonthly    PlanNewParamsPricesPriceCadence = "monthly"
	PlanNewParamsPricesPriceCadenceQuarterly  PlanNewParamsPricesPriceCadence = "quarterly"
	PlanNewParamsPricesPriceCadenceSemiAnnual PlanNewParamsPricesPriceCadence = "semi_annual"
	PlanNewParamsPricesPriceCadenceAnnual     PlanNewParamsPricesPriceCadence = "annual"
	PlanNewParamsPricesPriceCadenceCustom     PlanNewParamsPricesPriceCadence = "custom"
)

func (r PlanNewParamsPricesPriceCadence) IsKnown() bool {
	switch r {
	case PlanNewParamsPricesPriceCadenceOneTime, PlanNewParamsPricesPriceCadenceMonthly, PlanNewParamsPricesPriceCadenceQuarterly, PlanNewParamsPricesPriceCadenceSemiAnnual, PlanNewParamsPricesPriceCadenceAnnual, PlanNewParamsPricesPriceCadenceCustom:
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
	PlanNewParamsPricesPriceModelTypeGroupedTieredMatrix             PlanNewParamsPricesPriceModelType = "grouped_tiered_matrix"
	PlanNewParamsPricesPriceModelTypeTieredPackageWithMinimum        PlanNewParamsPricesPriceModelType = "tiered_package_with_minimum"
	PlanNewParamsPricesPriceModelTypePackageWithAllocation           PlanNewParamsPricesPriceModelType = "package_with_allocation"
	PlanNewParamsPricesPriceModelTypeUnitWithPercent                 PlanNewParamsPricesPriceModelType = "unit_with_percent"
	PlanNewParamsPricesPriceModelTypeMatrixWithAllocation            PlanNewParamsPricesPriceModelType = "matrix_with_allocation"
	PlanNewParamsPricesPriceModelTypeTieredMatrixWithAllocation      PlanNewParamsPricesPriceModelType = "tiered_matrix_with_allocation"
	PlanNewParamsPricesPriceModelTypeMatrixWithThresholdDiscounts    PlanNewParamsPricesPriceModelType = "matrix_with_threshold_discounts"
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
	PlanNewParamsPricesPriceModelTypeDailyCreditAllowance            PlanNewParamsPricesPriceModelType = "daily_credit_allowance"
	PlanNewParamsPricesPriceModelTypeMeteredAllowance                PlanNewParamsPricesPriceModelType = "metered_allowance"
	PlanNewParamsPricesPriceModelTypeMinimumComposite                PlanNewParamsPricesPriceModelType = "minimum_composite"
	PlanNewParamsPricesPriceModelTypePercent                         PlanNewParamsPricesPriceModelType = "percent"
	PlanNewParamsPricesPriceModelTypeEventOutput                     PlanNewParamsPricesPriceModelType = "event_output"
	PlanNewParamsPricesPriceModelTypeUnitWithMultiplier              PlanNewParamsPricesPriceModelType = "unit_with_multiplier"
)

func (r PlanNewParamsPricesPriceModelType) IsKnown() bool {
	switch r {
	case PlanNewParamsPricesPriceModelTypeUnit, PlanNewParamsPricesPriceModelTypeTiered, PlanNewParamsPricesPriceModelTypeBulk, PlanNewParamsPricesPriceModelTypeBulkWithFilters, PlanNewParamsPricesPriceModelTypePackage, PlanNewParamsPricesPriceModelTypeMatrix, PlanNewParamsPricesPriceModelTypeThresholdTotalAmount, PlanNewParamsPricesPriceModelTypeTieredPackage, PlanNewParamsPricesPriceModelTypeTieredWithMinimum, PlanNewParamsPricesPriceModelTypeGroupedTiered, PlanNewParamsPricesPriceModelTypeGroupedTieredMatrix, PlanNewParamsPricesPriceModelTypeTieredPackageWithMinimum, PlanNewParamsPricesPriceModelTypePackageWithAllocation, PlanNewParamsPricesPriceModelTypeUnitWithPercent, PlanNewParamsPricesPriceModelTypeMatrixWithAllocation, PlanNewParamsPricesPriceModelTypeTieredMatrixWithAllocation, PlanNewParamsPricesPriceModelTypeMatrixWithThresholdDiscounts, PlanNewParamsPricesPriceModelTypeTieredWithProration, PlanNewParamsPricesPriceModelTypeUnitWithProration, PlanNewParamsPricesPriceModelTypeGroupedAllocation, PlanNewParamsPricesPriceModelTypeBulkWithProration, PlanNewParamsPricesPriceModelTypeGroupedWithProratedMinimum, PlanNewParamsPricesPriceModelTypeGroupedWithMeteredMinimum, PlanNewParamsPricesPriceModelTypeGroupedWithMinMaxThresholds, PlanNewParamsPricesPriceModelTypeMatrixWithDisplayName, PlanNewParamsPricesPriceModelTypeGroupedTieredPackage, PlanNewParamsPricesPriceModelTypeMaxGroupTieredPackage, PlanNewParamsPricesPriceModelTypeScalableMatrixWithUnitPricing, PlanNewParamsPricesPriceModelTypeScalableMatrixWithTieredPricing, PlanNewParamsPricesPriceModelTypeCumulativeGroupedBulk, PlanNewParamsPricesPriceModelTypeCumulativeGroupedAllocation, PlanNewParamsPricesPriceModelTypeDailyCreditAllowance, PlanNewParamsPricesPriceModelTypeMeteredAllowance, PlanNewParamsPricesPriceModelTypeMinimumComposite, PlanNewParamsPricesPriceModelTypePercent, PlanNewParamsPricesPriceModelTypeEventOutput, PlanNewParamsPricesPriceModelTypeUnitWithMultiplier:
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
	Tiers         param.Field[interface{}]                                 `json:"tiers"`
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
// [PlanNewParamsAdjustmentsAdjustmentNewTieredPercentageDiscount],
// [PlanNewParamsAdjustmentsAdjustment].
type PlanNewParamsAdjustmentsAdjustmentUnion interface {
	ImplementsPlanNewParamsAdjustmentsAdjustmentUnion()
}

type PlanNewParamsAdjustmentsAdjustmentNewTieredPercentageDiscount struct {
	AdjustmentType param.Field[PlanNewParamsAdjustmentsAdjustmentNewTieredPercentageDiscountAdjustmentType] `json:"adjustment_type" api:"required"`
	Tiers          param.Field[[]PlanNewParamsAdjustmentsAdjustmentNewTieredPercentageDiscountTier]         `json:"tiers" api:"required"`
	// If set, the adjustment will apply to every price on the subscription.
	AppliesToAll param.Field[PlanNewParamsAdjustmentsAdjustmentNewTieredPercentageDiscountAppliesToAll] `json:"applies_to_all"`
	// The set of item IDs to which this adjustment applies.
	AppliesToItemIDs param.Field[[]string] `json:"applies_to_item_ids"`
	// The set of price IDs to which this adjustment applies.
	AppliesToPriceIDs param.Field[[]string] `json:"applies_to_price_ids"`
	// If set, only prices in the specified currency will have the adjustment applied.
	Currency param.Field[string] `json:"currency"`
	// A list of filters that determine which prices this adjustment will apply to.
	Filters param.Field[[]PlanNewParamsAdjustmentsAdjustmentNewTieredPercentageDiscountFilter] `json:"filters"`
	// When false, this adjustment will be applied to a single price. Otherwise, it
	// will be applied at the invoice level, possibly to multiple prices.
	IsInvoiceLevel param.Field[bool] `json:"is_invoice_level"`
	// If set, only prices of the specified type will have the adjustment applied.
	PriceType param.Field[PlanNewParamsAdjustmentsAdjustmentNewTieredPercentageDiscountPriceType] `json:"price_type"`
}

func (r PlanNewParamsAdjustmentsAdjustmentNewTieredPercentageDiscount) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r PlanNewParamsAdjustmentsAdjustmentNewTieredPercentageDiscount) ImplementsPlanNewParamsAdjustmentsAdjustmentUnion() {
}

type PlanNewParamsAdjustmentsAdjustmentNewTieredPercentageDiscountAdjustmentType string

const (
	PlanNewParamsAdjustmentsAdjustmentNewTieredPercentageDiscountAdjustmentTypeTieredPercentageDiscount PlanNewParamsAdjustmentsAdjustmentNewTieredPercentageDiscountAdjustmentType = "tiered_percentage_discount"
)

func (r PlanNewParamsAdjustmentsAdjustmentNewTieredPercentageDiscountAdjustmentType) IsKnown() bool {
	switch r {
	case PlanNewParamsAdjustmentsAdjustmentNewTieredPercentageDiscountAdjustmentTypeTieredPercentageDiscount:
		return true
	}
	return false
}

type PlanNewParamsAdjustmentsAdjustmentNewTieredPercentageDiscountTier struct {
	// Exclusive lower bound of cumulative spend for this tier.
	LowerBound param.Field[float64] `json:"lower_bound" api:"required"`
	// The percentage (0-1) discounted from spend in this tier.
	Percentage param.Field[float64] `json:"percentage" api:"required"`
	// Inclusive upper bound of cumulative spend; null for the final open-ended tier.
	UpperBound param.Field[float64] `json:"upper_bound"`
}

func (r PlanNewParamsAdjustmentsAdjustmentNewTieredPercentageDiscountTier) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// If set, the adjustment will apply to every price on the subscription.
type PlanNewParamsAdjustmentsAdjustmentNewTieredPercentageDiscountAppliesToAll bool

const (
	PlanNewParamsAdjustmentsAdjustmentNewTieredPercentageDiscountAppliesToAllTrue PlanNewParamsAdjustmentsAdjustmentNewTieredPercentageDiscountAppliesToAll = true
)

func (r PlanNewParamsAdjustmentsAdjustmentNewTieredPercentageDiscountAppliesToAll) IsKnown() bool {
	switch r {
	case PlanNewParamsAdjustmentsAdjustmentNewTieredPercentageDiscountAppliesToAllTrue:
		return true
	}
	return false
}

type PlanNewParamsAdjustmentsAdjustmentNewTieredPercentageDiscountFilter struct {
	// The property of the price to filter on.
	Field param.Field[PlanNewParamsAdjustmentsAdjustmentNewTieredPercentageDiscountFiltersField] `json:"field" api:"required"`
	// Should prices that match the filter be included or excluded.
	Operator param.Field[PlanNewParamsAdjustmentsAdjustmentNewTieredPercentageDiscountFiltersOperator] `json:"operator" api:"required"`
	// The IDs or values that match this filter.
	Values param.Field[[]string] `json:"values" api:"required"`
}

func (r PlanNewParamsAdjustmentsAdjustmentNewTieredPercentageDiscountFilter) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// The property of the price to filter on.
type PlanNewParamsAdjustmentsAdjustmentNewTieredPercentageDiscountFiltersField string

const (
	PlanNewParamsAdjustmentsAdjustmentNewTieredPercentageDiscountFiltersFieldPriceID       PlanNewParamsAdjustmentsAdjustmentNewTieredPercentageDiscountFiltersField = "price_id"
	PlanNewParamsAdjustmentsAdjustmentNewTieredPercentageDiscountFiltersFieldItemID        PlanNewParamsAdjustmentsAdjustmentNewTieredPercentageDiscountFiltersField = "item_id"
	PlanNewParamsAdjustmentsAdjustmentNewTieredPercentageDiscountFiltersFieldPriceType     PlanNewParamsAdjustmentsAdjustmentNewTieredPercentageDiscountFiltersField = "price_type"
	PlanNewParamsAdjustmentsAdjustmentNewTieredPercentageDiscountFiltersFieldCurrency      PlanNewParamsAdjustmentsAdjustmentNewTieredPercentageDiscountFiltersField = "currency"
	PlanNewParamsAdjustmentsAdjustmentNewTieredPercentageDiscountFiltersFieldPricingUnitID PlanNewParamsAdjustmentsAdjustmentNewTieredPercentageDiscountFiltersField = "pricing_unit_id"
)

func (r PlanNewParamsAdjustmentsAdjustmentNewTieredPercentageDiscountFiltersField) IsKnown() bool {
	switch r {
	case PlanNewParamsAdjustmentsAdjustmentNewTieredPercentageDiscountFiltersFieldPriceID, PlanNewParamsAdjustmentsAdjustmentNewTieredPercentageDiscountFiltersFieldItemID, PlanNewParamsAdjustmentsAdjustmentNewTieredPercentageDiscountFiltersFieldPriceType, PlanNewParamsAdjustmentsAdjustmentNewTieredPercentageDiscountFiltersFieldCurrency, PlanNewParamsAdjustmentsAdjustmentNewTieredPercentageDiscountFiltersFieldPricingUnitID:
		return true
	}
	return false
}

// Should prices that match the filter be included or excluded.
type PlanNewParamsAdjustmentsAdjustmentNewTieredPercentageDiscountFiltersOperator string

const (
	PlanNewParamsAdjustmentsAdjustmentNewTieredPercentageDiscountFiltersOperatorIncludes PlanNewParamsAdjustmentsAdjustmentNewTieredPercentageDiscountFiltersOperator = "includes"
	PlanNewParamsAdjustmentsAdjustmentNewTieredPercentageDiscountFiltersOperatorExcludes PlanNewParamsAdjustmentsAdjustmentNewTieredPercentageDiscountFiltersOperator = "excludes"
)

func (r PlanNewParamsAdjustmentsAdjustmentNewTieredPercentageDiscountFiltersOperator) IsKnown() bool {
	switch r {
	case PlanNewParamsAdjustmentsAdjustmentNewTieredPercentageDiscountFiltersOperatorIncludes, PlanNewParamsAdjustmentsAdjustmentNewTieredPercentageDiscountFiltersOperatorExcludes:
		return true
	}
	return false
}

// If set, only prices of the specified type will have the adjustment applied.
type PlanNewParamsAdjustmentsAdjustmentNewTieredPercentageDiscountPriceType string

const (
	PlanNewParamsAdjustmentsAdjustmentNewTieredPercentageDiscountPriceTypeUsage          PlanNewParamsAdjustmentsAdjustmentNewTieredPercentageDiscountPriceType = "usage"
	PlanNewParamsAdjustmentsAdjustmentNewTieredPercentageDiscountPriceTypeFixedInAdvance PlanNewParamsAdjustmentsAdjustmentNewTieredPercentageDiscountPriceType = "fixed_in_advance"
	PlanNewParamsAdjustmentsAdjustmentNewTieredPercentageDiscountPriceTypeFixedInArrears PlanNewParamsAdjustmentsAdjustmentNewTieredPercentageDiscountPriceType = "fixed_in_arrears"
	PlanNewParamsAdjustmentsAdjustmentNewTieredPercentageDiscountPriceTypeFixed          PlanNewParamsAdjustmentsAdjustmentNewTieredPercentageDiscountPriceType = "fixed"
	PlanNewParamsAdjustmentsAdjustmentNewTieredPercentageDiscountPriceTypeInArrears      PlanNewParamsAdjustmentsAdjustmentNewTieredPercentageDiscountPriceType = "in_arrears"
)

func (r PlanNewParamsAdjustmentsAdjustmentNewTieredPercentageDiscountPriceType) IsKnown() bool {
	switch r {
	case PlanNewParamsAdjustmentsAdjustmentNewTieredPercentageDiscountPriceTypeUsage, PlanNewParamsAdjustmentsAdjustmentNewTieredPercentageDiscountPriceTypeFixedInAdvance, PlanNewParamsAdjustmentsAdjustmentNewTieredPercentageDiscountPriceTypeFixedInArrears, PlanNewParamsAdjustmentsAdjustmentNewTieredPercentageDiscountPriceTypeFixed, PlanNewParamsAdjustmentsAdjustmentNewTieredPercentageDiscountPriceTypeInArrears:
		return true
	}
	return false
}

type PlanNewParamsAdjustmentsAdjustmentAdjustmentType string

const (
	PlanNewParamsAdjustmentsAdjustmentAdjustmentTypePercentageDiscount       PlanNewParamsAdjustmentsAdjustmentAdjustmentType = "percentage_discount"
	PlanNewParamsAdjustmentsAdjustmentAdjustmentTypeUsageDiscount            PlanNewParamsAdjustmentsAdjustmentAdjustmentType = "usage_discount"
	PlanNewParamsAdjustmentsAdjustmentAdjustmentTypeAmountDiscount           PlanNewParamsAdjustmentsAdjustmentAdjustmentType = "amount_discount"
	PlanNewParamsAdjustmentsAdjustmentAdjustmentTypeMinimum                  PlanNewParamsAdjustmentsAdjustmentAdjustmentType = "minimum"
	PlanNewParamsAdjustmentsAdjustmentAdjustmentTypeMaximum                  PlanNewParamsAdjustmentsAdjustmentAdjustmentType = "maximum"
	PlanNewParamsAdjustmentsAdjustmentAdjustmentTypeTieredPercentageDiscount PlanNewParamsAdjustmentsAdjustmentAdjustmentType = "tiered_percentage_discount"
)

func (r PlanNewParamsAdjustmentsAdjustmentAdjustmentType) IsKnown() bool {
	switch r {
	case PlanNewParamsAdjustmentsAdjustmentAdjustmentTypePercentageDiscount, PlanNewParamsAdjustmentsAdjustmentAdjustmentTypeUsageDiscount, PlanNewParamsAdjustmentsAdjustmentAdjustmentTypeAmountDiscount, PlanNewParamsAdjustmentsAdjustmentAdjustmentTypeMinimum, PlanNewParamsAdjustmentsAdjustmentAdjustmentTypeMaximum, PlanNewParamsAdjustmentsAdjustmentAdjustmentTypeTieredPercentageDiscount:
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
	// The name of the plan.
	Name param.Field[string] `json:"name"`
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

type PlanNewPlanVersionParams struct {
	// New version number.
	Version param.Field[int64] `json:"version" api:"required"`
	// Additional adjustments to be added to the plan.
	AddAdjustments param.Field[[]PlanNewPlanVersionParamsAddAdjustment] `json:"add_adjustments"`
	// Additional prices to be added to the plan.
	AddPrices param.Field[[]PlanNewPlanVersionParamsAddPrice] `json:"add_prices"`
	// Adjustments to be removed from the plan.
	RemoveAdjustments param.Field[[]PlanNewPlanVersionParamsRemoveAdjustment] `json:"remove_adjustments"`
	// Prices to be removed from the plan.
	RemovePrices param.Field[[]PlanNewPlanVersionParamsRemovePrice] `json:"remove_prices"`
	// Adjustments to be replaced with additional adjustments on the plan.
	ReplaceAdjustments param.Field[[]PlanNewPlanVersionParamsReplaceAdjustment] `json:"replace_adjustments"`
	// Prices to be replaced with additional prices on the plan.
	ReplacePrices param.Field[[]PlanNewPlanVersionParamsReplacePrice] `json:"replace_prices"`
	// Set this new plan version as the default
	SetAsDefault param.Field[bool] `json:"set_as_default"`
}

func (r PlanNewPlanVersionParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type PlanNewPlanVersionParamsAddAdjustment struct {
	// The definition of a new adjustment to create and add to the plan.
	Adjustment param.Field[PlanNewPlanVersionParamsAddAdjustmentsAdjustmentUnion] `json:"adjustment" api:"required"`
	// The phase to add this adjustment to.
	PlanPhaseOrder param.Field[int64] `json:"plan_phase_order"`
}

func (r PlanNewPlanVersionParamsAddAdjustment) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// The definition of a new adjustment to create and add to the plan.
type PlanNewPlanVersionParamsAddAdjustmentsAdjustment struct {
	AdjustmentType param.Field[PlanNewPlanVersionParamsAddAdjustmentsAdjustmentAdjustmentType] `json:"adjustment_type" api:"required"`
	AmountDiscount param.Field[string]                                                         `json:"amount_discount"`
	// If set, the adjustment will apply to every price on the subscription.
	AppliesToAll      param.Field[PlanNewPlanVersionParamsAddAdjustmentsAdjustmentAppliesToAll] `json:"applies_to_all"`
	AppliesToItemIDs  param.Field[interface{}]                                                  `json:"applies_to_item_ids"`
	AppliesToPriceIDs param.Field[interface{}]                                                  `json:"applies_to_price_ids"`
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
	PriceType     param.Field[PlanNewPlanVersionParamsAddAdjustmentsAdjustmentPriceType] `json:"price_type"`
	Tiers         param.Field[interface{}]                                               `json:"tiers"`
	UsageDiscount param.Field[float64]                                                   `json:"usage_discount"`
}

func (r PlanNewPlanVersionParamsAddAdjustmentsAdjustment) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r PlanNewPlanVersionParamsAddAdjustmentsAdjustment) ImplementsPlanNewPlanVersionParamsAddAdjustmentsAdjustmentUnion() {
}

// The definition of a new adjustment to create and add to the plan.
//
// Satisfied by [shared.NewPercentageDiscountParam],
// [shared.NewUsageDiscountParam], [shared.NewAmountDiscountParam],
// [shared.NewMinimumParam], [shared.NewMaximumParam],
// [PlanNewPlanVersionParamsAddAdjustmentsAdjustmentNewTieredPercentageDiscount],
// [PlanNewPlanVersionParamsAddAdjustmentsAdjustment].
type PlanNewPlanVersionParamsAddAdjustmentsAdjustmentUnion interface {
	ImplementsPlanNewPlanVersionParamsAddAdjustmentsAdjustmentUnion()
}

type PlanNewPlanVersionParamsAddAdjustmentsAdjustmentNewTieredPercentageDiscount struct {
	AdjustmentType param.Field[PlanNewPlanVersionParamsAddAdjustmentsAdjustmentNewTieredPercentageDiscountAdjustmentType] `json:"adjustment_type" api:"required"`
	Tiers          param.Field[[]PlanNewPlanVersionParamsAddAdjustmentsAdjustmentNewTieredPercentageDiscountTier]         `json:"tiers" api:"required"`
	// If set, the adjustment will apply to every price on the subscription.
	AppliesToAll param.Field[PlanNewPlanVersionParamsAddAdjustmentsAdjustmentNewTieredPercentageDiscountAppliesToAll] `json:"applies_to_all"`
	// The set of item IDs to which this adjustment applies.
	AppliesToItemIDs param.Field[[]string] `json:"applies_to_item_ids"`
	// The set of price IDs to which this adjustment applies.
	AppliesToPriceIDs param.Field[[]string] `json:"applies_to_price_ids"`
	// If set, only prices in the specified currency will have the adjustment applied.
	Currency param.Field[string] `json:"currency"`
	// A list of filters that determine which prices this adjustment will apply to.
	Filters param.Field[[]PlanNewPlanVersionParamsAddAdjustmentsAdjustmentNewTieredPercentageDiscountFilter] `json:"filters"`
	// When false, this adjustment will be applied to a single price. Otherwise, it
	// will be applied at the invoice level, possibly to multiple prices.
	IsInvoiceLevel param.Field[bool] `json:"is_invoice_level"`
	// If set, only prices of the specified type will have the adjustment applied.
	PriceType param.Field[PlanNewPlanVersionParamsAddAdjustmentsAdjustmentNewTieredPercentageDiscountPriceType] `json:"price_type"`
}

func (r PlanNewPlanVersionParamsAddAdjustmentsAdjustmentNewTieredPercentageDiscount) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r PlanNewPlanVersionParamsAddAdjustmentsAdjustmentNewTieredPercentageDiscount) ImplementsPlanNewPlanVersionParamsAddAdjustmentsAdjustmentUnion() {
}

type PlanNewPlanVersionParamsAddAdjustmentsAdjustmentNewTieredPercentageDiscountAdjustmentType string

const (
	PlanNewPlanVersionParamsAddAdjustmentsAdjustmentNewTieredPercentageDiscountAdjustmentTypeTieredPercentageDiscount PlanNewPlanVersionParamsAddAdjustmentsAdjustmentNewTieredPercentageDiscountAdjustmentType = "tiered_percentage_discount"
)

func (r PlanNewPlanVersionParamsAddAdjustmentsAdjustmentNewTieredPercentageDiscountAdjustmentType) IsKnown() bool {
	switch r {
	case PlanNewPlanVersionParamsAddAdjustmentsAdjustmentNewTieredPercentageDiscountAdjustmentTypeTieredPercentageDiscount:
		return true
	}
	return false
}

type PlanNewPlanVersionParamsAddAdjustmentsAdjustmentNewTieredPercentageDiscountTier struct {
	// Exclusive lower bound of cumulative spend for this tier.
	LowerBound param.Field[float64] `json:"lower_bound" api:"required"`
	// The percentage (0-1) discounted from spend in this tier.
	Percentage param.Field[float64] `json:"percentage" api:"required"`
	// Inclusive upper bound of cumulative spend; null for the final open-ended tier.
	UpperBound param.Field[float64] `json:"upper_bound"`
}

func (r PlanNewPlanVersionParamsAddAdjustmentsAdjustmentNewTieredPercentageDiscountTier) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// If set, the adjustment will apply to every price on the subscription.
type PlanNewPlanVersionParamsAddAdjustmentsAdjustmentNewTieredPercentageDiscountAppliesToAll bool

const (
	PlanNewPlanVersionParamsAddAdjustmentsAdjustmentNewTieredPercentageDiscountAppliesToAllTrue PlanNewPlanVersionParamsAddAdjustmentsAdjustmentNewTieredPercentageDiscountAppliesToAll = true
)

func (r PlanNewPlanVersionParamsAddAdjustmentsAdjustmentNewTieredPercentageDiscountAppliesToAll) IsKnown() bool {
	switch r {
	case PlanNewPlanVersionParamsAddAdjustmentsAdjustmentNewTieredPercentageDiscountAppliesToAllTrue:
		return true
	}
	return false
}

type PlanNewPlanVersionParamsAddAdjustmentsAdjustmentNewTieredPercentageDiscountFilter struct {
	// The property of the price to filter on.
	Field param.Field[PlanNewPlanVersionParamsAddAdjustmentsAdjustmentNewTieredPercentageDiscountFiltersField] `json:"field" api:"required"`
	// Should prices that match the filter be included or excluded.
	Operator param.Field[PlanNewPlanVersionParamsAddAdjustmentsAdjustmentNewTieredPercentageDiscountFiltersOperator] `json:"operator" api:"required"`
	// The IDs or values that match this filter.
	Values param.Field[[]string] `json:"values" api:"required"`
}

func (r PlanNewPlanVersionParamsAddAdjustmentsAdjustmentNewTieredPercentageDiscountFilter) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// The property of the price to filter on.
type PlanNewPlanVersionParamsAddAdjustmentsAdjustmentNewTieredPercentageDiscountFiltersField string

const (
	PlanNewPlanVersionParamsAddAdjustmentsAdjustmentNewTieredPercentageDiscountFiltersFieldPriceID       PlanNewPlanVersionParamsAddAdjustmentsAdjustmentNewTieredPercentageDiscountFiltersField = "price_id"
	PlanNewPlanVersionParamsAddAdjustmentsAdjustmentNewTieredPercentageDiscountFiltersFieldItemID        PlanNewPlanVersionParamsAddAdjustmentsAdjustmentNewTieredPercentageDiscountFiltersField = "item_id"
	PlanNewPlanVersionParamsAddAdjustmentsAdjustmentNewTieredPercentageDiscountFiltersFieldPriceType     PlanNewPlanVersionParamsAddAdjustmentsAdjustmentNewTieredPercentageDiscountFiltersField = "price_type"
	PlanNewPlanVersionParamsAddAdjustmentsAdjustmentNewTieredPercentageDiscountFiltersFieldCurrency      PlanNewPlanVersionParamsAddAdjustmentsAdjustmentNewTieredPercentageDiscountFiltersField = "currency"
	PlanNewPlanVersionParamsAddAdjustmentsAdjustmentNewTieredPercentageDiscountFiltersFieldPricingUnitID PlanNewPlanVersionParamsAddAdjustmentsAdjustmentNewTieredPercentageDiscountFiltersField = "pricing_unit_id"
)

func (r PlanNewPlanVersionParamsAddAdjustmentsAdjustmentNewTieredPercentageDiscountFiltersField) IsKnown() bool {
	switch r {
	case PlanNewPlanVersionParamsAddAdjustmentsAdjustmentNewTieredPercentageDiscountFiltersFieldPriceID, PlanNewPlanVersionParamsAddAdjustmentsAdjustmentNewTieredPercentageDiscountFiltersFieldItemID, PlanNewPlanVersionParamsAddAdjustmentsAdjustmentNewTieredPercentageDiscountFiltersFieldPriceType, PlanNewPlanVersionParamsAddAdjustmentsAdjustmentNewTieredPercentageDiscountFiltersFieldCurrency, PlanNewPlanVersionParamsAddAdjustmentsAdjustmentNewTieredPercentageDiscountFiltersFieldPricingUnitID:
		return true
	}
	return false
}

// Should prices that match the filter be included or excluded.
type PlanNewPlanVersionParamsAddAdjustmentsAdjustmentNewTieredPercentageDiscountFiltersOperator string

const (
	PlanNewPlanVersionParamsAddAdjustmentsAdjustmentNewTieredPercentageDiscountFiltersOperatorIncludes PlanNewPlanVersionParamsAddAdjustmentsAdjustmentNewTieredPercentageDiscountFiltersOperator = "includes"
	PlanNewPlanVersionParamsAddAdjustmentsAdjustmentNewTieredPercentageDiscountFiltersOperatorExcludes PlanNewPlanVersionParamsAddAdjustmentsAdjustmentNewTieredPercentageDiscountFiltersOperator = "excludes"
)

func (r PlanNewPlanVersionParamsAddAdjustmentsAdjustmentNewTieredPercentageDiscountFiltersOperator) IsKnown() bool {
	switch r {
	case PlanNewPlanVersionParamsAddAdjustmentsAdjustmentNewTieredPercentageDiscountFiltersOperatorIncludes, PlanNewPlanVersionParamsAddAdjustmentsAdjustmentNewTieredPercentageDiscountFiltersOperatorExcludes:
		return true
	}
	return false
}

// If set, only prices of the specified type will have the adjustment applied.
type PlanNewPlanVersionParamsAddAdjustmentsAdjustmentNewTieredPercentageDiscountPriceType string

const (
	PlanNewPlanVersionParamsAddAdjustmentsAdjustmentNewTieredPercentageDiscountPriceTypeUsage          PlanNewPlanVersionParamsAddAdjustmentsAdjustmentNewTieredPercentageDiscountPriceType = "usage"
	PlanNewPlanVersionParamsAddAdjustmentsAdjustmentNewTieredPercentageDiscountPriceTypeFixedInAdvance PlanNewPlanVersionParamsAddAdjustmentsAdjustmentNewTieredPercentageDiscountPriceType = "fixed_in_advance"
	PlanNewPlanVersionParamsAddAdjustmentsAdjustmentNewTieredPercentageDiscountPriceTypeFixedInArrears PlanNewPlanVersionParamsAddAdjustmentsAdjustmentNewTieredPercentageDiscountPriceType = "fixed_in_arrears"
	PlanNewPlanVersionParamsAddAdjustmentsAdjustmentNewTieredPercentageDiscountPriceTypeFixed          PlanNewPlanVersionParamsAddAdjustmentsAdjustmentNewTieredPercentageDiscountPriceType = "fixed"
	PlanNewPlanVersionParamsAddAdjustmentsAdjustmentNewTieredPercentageDiscountPriceTypeInArrears      PlanNewPlanVersionParamsAddAdjustmentsAdjustmentNewTieredPercentageDiscountPriceType = "in_arrears"
)

func (r PlanNewPlanVersionParamsAddAdjustmentsAdjustmentNewTieredPercentageDiscountPriceType) IsKnown() bool {
	switch r {
	case PlanNewPlanVersionParamsAddAdjustmentsAdjustmentNewTieredPercentageDiscountPriceTypeUsage, PlanNewPlanVersionParamsAddAdjustmentsAdjustmentNewTieredPercentageDiscountPriceTypeFixedInAdvance, PlanNewPlanVersionParamsAddAdjustmentsAdjustmentNewTieredPercentageDiscountPriceTypeFixedInArrears, PlanNewPlanVersionParamsAddAdjustmentsAdjustmentNewTieredPercentageDiscountPriceTypeFixed, PlanNewPlanVersionParamsAddAdjustmentsAdjustmentNewTieredPercentageDiscountPriceTypeInArrears:
		return true
	}
	return false
}

type PlanNewPlanVersionParamsAddAdjustmentsAdjustmentAdjustmentType string

const (
	PlanNewPlanVersionParamsAddAdjustmentsAdjustmentAdjustmentTypePercentageDiscount       PlanNewPlanVersionParamsAddAdjustmentsAdjustmentAdjustmentType = "percentage_discount"
	PlanNewPlanVersionParamsAddAdjustmentsAdjustmentAdjustmentTypeUsageDiscount            PlanNewPlanVersionParamsAddAdjustmentsAdjustmentAdjustmentType = "usage_discount"
	PlanNewPlanVersionParamsAddAdjustmentsAdjustmentAdjustmentTypeAmountDiscount           PlanNewPlanVersionParamsAddAdjustmentsAdjustmentAdjustmentType = "amount_discount"
	PlanNewPlanVersionParamsAddAdjustmentsAdjustmentAdjustmentTypeMinimum                  PlanNewPlanVersionParamsAddAdjustmentsAdjustmentAdjustmentType = "minimum"
	PlanNewPlanVersionParamsAddAdjustmentsAdjustmentAdjustmentTypeMaximum                  PlanNewPlanVersionParamsAddAdjustmentsAdjustmentAdjustmentType = "maximum"
	PlanNewPlanVersionParamsAddAdjustmentsAdjustmentAdjustmentTypeTieredPercentageDiscount PlanNewPlanVersionParamsAddAdjustmentsAdjustmentAdjustmentType = "tiered_percentage_discount"
)

func (r PlanNewPlanVersionParamsAddAdjustmentsAdjustmentAdjustmentType) IsKnown() bool {
	switch r {
	case PlanNewPlanVersionParamsAddAdjustmentsAdjustmentAdjustmentTypePercentageDiscount, PlanNewPlanVersionParamsAddAdjustmentsAdjustmentAdjustmentTypeUsageDiscount, PlanNewPlanVersionParamsAddAdjustmentsAdjustmentAdjustmentTypeAmountDiscount, PlanNewPlanVersionParamsAddAdjustmentsAdjustmentAdjustmentTypeMinimum, PlanNewPlanVersionParamsAddAdjustmentsAdjustmentAdjustmentTypeMaximum, PlanNewPlanVersionParamsAddAdjustmentsAdjustmentAdjustmentTypeTieredPercentageDiscount:
		return true
	}
	return false
}

// If set, the adjustment will apply to every price on the subscription.
type PlanNewPlanVersionParamsAddAdjustmentsAdjustmentAppliesToAll bool

const (
	PlanNewPlanVersionParamsAddAdjustmentsAdjustmentAppliesToAllTrue PlanNewPlanVersionParamsAddAdjustmentsAdjustmentAppliesToAll = true
)

func (r PlanNewPlanVersionParamsAddAdjustmentsAdjustmentAppliesToAll) IsKnown() bool {
	switch r {
	case PlanNewPlanVersionParamsAddAdjustmentsAdjustmentAppliesToAllTrue:
		return true
	}
	return false
}

// If set, only prices of the specified type will have the adjustment applied.
type PlanNewPlanVersionParamsAddAdjustmentsAdjustmentPriceType string

const (
	PlanNewPlanVersionParamsAddAdjustmentsAdjustmentPriceTypeUsage          PlanNewPlanVersionParamsAddAdjustmentsAdjustmentPriceType = "usage"
	PlanNewPlanVersionParamsAddAdjustmentsAdjustmentPriceTypeFixedInAdvance PlanNewPlanVersionParamsAddAdjustmentsAdjustmentPriceType = "fixed_in_advance"
	PlanNewPlanVersionParamsAddAdjustmentsAdjustmentPriceTypeFixedInArrears PlanNewPlanVersionParamsAddAdjustmentsAdjustmentPriceType = "fixed_in_arrears"
	PlanNewPlanVersionParamsAddAdjustmentsAdjustmentPriceTypeFixed          PlanNewPlanVersionParamsAddAdjustmentsAdjustmentPriceType = "fixed"
	PlanNewPlanVersionParamsAddAdjustmentsAdjustmentPriceTypeInArrears      PlanNewPlanVersionParamsAddAdjustmentsAdjustmentPriceType = "in_arrears"
)

func (r PlanNewPlanVersionParamsAddAdjustmentsAdjustmentPriceType) IsKnown() bool {
	switch r {
	case PlanNewPlanVersionParamsAddAdjustmentsAdjustmentPriceTypeUsage, PlanNewPlanVersionParamsAddAdjustmentsAdjustmentPriceTypeFixedInAdvance, PlanNewPlanVersionParamsAddAdjustmentsAdjustmentPriceTypeFixedInArrears, PlanNewPlanVersionParamsAddAdjustmentsAdjustmentPriceTypeFixed, PlanNewPlanVersionParamsAddAdjustmentsAdjustmentPriceTypeInArrears:
		return true
	}
	return false
}

type PlanNewPlanVersionParamsAddPrice struct {
	// The allocation price to add to the plan.
	AllocationPrice param.Field[shared.NewAllocationPriceParam] `json:"allocation_price"`
	// The license allocation price to add to the plan.
	LicenseAllocationPrice param.Field[PlanNewPlanVersionParamsAddPricesLicenseAllocationPrice] `json:"license_allocation_price"`
	// The phase to add this price to.
	PlanPhaseOrder param.Field[int64] `json:"plan_phase_order"`
	// New plan price request body params.
	Price param.Field[PlanNewPlanVersionParamsAddPricesPriceUnion] `json:"price"`
}

func (r PlanNewPlanVersionParamsAddPrice) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// The license allocation price to add to the plan.
type PlanNewPlanVersionParamsAddPricesLicenseAllocationPrice struct {
	// The cadence to bill for this price on.
	Cadence param.Field[PlanNewPlanVersionParamsAddPricesLicenseAllocationPriceCadence] `json:"cadence" api:"required"`
	// The id of the item the price will be associated with.
	ItemID param.Field[string] `json:"item_id" api:"required"`
	// License allocations to associate with this price. Each entry defines a
	// per-license credit pool granted each cadence. Requires license_type_id or
	// license_type_configuration to be set. Usage prices draw down from these
	// allocations only if they set the same license type.
	LicenseAllocations param.Field[[]PlanNewPlanVersionParamsAddPricesLicenseAllocationPriceLicenseAllocation] `json:"license_allocations" api:"required"`
	// The pricing model type
	ModelType param.Field[PlanNewPlanVersionParamsAddPricesLicenseAllocationPriceModelType] `json:"model_type" api:"required"`
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
	ConversionRateConfig param.Field[PlanNewPlanVersionParamsAddPricesLicenseAllocationPriceConversionRateConfigUnion] `json:"conversion_rate_config"`
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
	// The ID of the license type to associate with this price. On a usage price this
	// also marks the price as eligible to draw down from that license type's
	// allocation; a usage price created without it is billed normally. Usage prices
	// with a license type must use the `unit` model, and only draw down when their
	// currency matches the allocation's.
	LicenseTypeID param.Field[string] `json:"license_type_id"`
	// User-specified key/value pairs for the resource. Individual keys can be removed
	// by setting the value to `null`, and the entire metadata mapping can be cleared
	// by setting `metadata` to `null`.
	Metadata param.Field[map[string]string] `json:"metadata"`
	// A transient ID that can be used to reference this price when adding adjustments
	// in the same API call.
	ReferenceID param.Field[string] `json:"reference_id"`
}

func (r PlanNewPlanVersionParamsAddPricesLicenseAllocationPrice) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// The cadence to bill for this price on.
type PlanNewPlanVersionParamsAddPricesLicenseAllocationPriceCadence string

const (
	PlanNewPlanVersionParamsAddPricesLicenseAllocationPriceCadenceOneTime    PlanNewPlanVersionParamsAddPricesLicenseAllocationPriceCadence = "one_time"
	PlanNewPlanVersionParamsAddPricesLicenseAllocationPriceCadenceMonthly    PlanNewPlanVersionParamsAddPricesLicenseAllocationPriceCadence = "monthly"
	PlanNewPlanVersionParamsAddPricesLicenseAllocationPriceCadenceQuarterly  PlanNewPlanVersionParamsAddPricesLicenseAllocationPriceCadence = "quarterly"
	PlanNewPlanVersionParamsAddPricesLicenseAllocationPriceCadenceSemiAnnual PlanNewPlanVersionParamsAddPricesLicenseAllocationPriceCadence = "semi_annual"
	PlanNewPlanVersionParamsAddPricesLicenseAllocationPriceCadenceAnnual     PlanNewPlanVersionParamsAddPricesLicenseAllocationPriceCadence = "annual"
	PlanNewPlanVersionParamsAddPricesLicenseAllocationPriceCadenceCustom     PlanNewPlanVersionParamsAddPricesLicenseAllocationPriceCadence = "custom"
)

func (r PlanNewPlanVersionParamsAddPricesLicenseAllocationPriceCadence) IsKnown() bool {
	switch r {
	case PlanNewPlanVersionParamsAddPricesLicenseAllocationPriceCadenceOneTime, PlanNewPlanVersionParamsAddPricesLicenseAllocationPriceCadenceMonthly, PlanNewPlanVersionParamsAddPricesLicenseAllocationPriceCadenceQuarterly, PlanNewPlanVersionParamsAddPricesLicenseAllocationPriceCadenceSemiAnnual, PlanNewPlanVersionParamsAddPricesLicenseAllocationPriceCadenceAnnual, PlanNewPlanVersionParamsAddPricesLicenseAllocationPriceCadenceCustom:
		return true
	}
	return false
}

type PlanNewPlanVersionParamsAddPricesLicenseAllocationPriceLicenseAllocation struct {
	// The amount of credits granted per active license per cadence.
	Amount param.Field[string] `json:"amount" api:"required"`
	// The currency of the license allocation.
	Currency param.Field[string] `json:"currency" api:"required"`
	// When True, overage beyond the allocation is written off.
	WriteOffOverage param.Field[bool] `json:"write_off_overage"`
}

func (r PlanNewPlanVersionParamsAddPricesLicenseAllocationPriceLicenseAllocation) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// The pricing model type
type PlanNewPlanVersionParamsAddPricesLicenseAllocationPriceModelType string

const (
	PlanNewPlanVersionParamsAddPricesLicenseAllocationPriceModelTypeUnit PlanNewPlanVersionParamsAddPricesLicenseAllocationPriceModelType = "unit"
)

func (r PlanNewPlanVersionParamsAddPricesLicenseAllocationPriceModelType) IsKnown() bool {
	switch r {
	case PlanNewPlanVersionParamsAddPricesLicenseAllocationPriceModelTypeUnit:
		return true
	}
	return false
}

type PlanNewPlanVersionParamsAddPricesLicenseAllocationPriceConversionRateConfig struct {
	ConversionRateType param.Field[PlanNewPlanVersionParamsAddPricesLicenseAllocationPriceConversionRateConfigConversionRateType] `json:"conversion_rate_type" api:"required"`
	TieredConfig       param.Field[shared.ConversionRateTieredConfigParam]                                                        `json:"tiered_config"`
	UnitConfig         param.Field[shared.ConversionRateUnitConfigParam]                                                          `json:"unit_config"`
}

func (r PlanNewPlanVersionParamsAddPricesLicenseAllocationPriceConversionRateConfig) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r PlanNewPlanVersionParamsAddPricesLicenseAllocationPriceConversionRateConfig) ImplementsPlanNewPlanVersionParamsAddPricesLicenseAllocationPriceConversionRateConfigUnion() {
}

// Satisfied by [shared.UnitConversionRateConfigParam],
// [shared.TieredConversionRateConfigParam],
// [PlanNewPlanVersionParamsAddPricesLicenseAllocationPriceConversionRateConfig].
type PlanNewPlanVersionParamsAddPricesLicenseAllocationPriceConversionRateConfigUnion interface {
	ImplementsPlanNewPlanVersionParamsAddPricesLicenseAllocationPriceConversionRateConfigUnion()
}

type PlanNewPlanVersionParamsAddPricesLicenseAllocationPriceConversionRateConfigConversionRateType string

const (
	PlanNewPlanVersionParamsAddPricesLicenseAllocationPriceConversionRateConfigConversionRateTypeUnit   PlanNewPlanVersionParamsAddPricesLicenseAllocationPriceConversionRateConfigConversionRateType = "unit"
	PlanNewPlanVersionParamsAddPricesLicenseAllocationPriceConversionRateConfigConversionRateTypeTiered PlanNewPlanVersionParamsAddPricesLicenseAllocationPriceConversionRateConfigConversionRateType = "tiered"
)

func (r PlanNewPlanVersionParamsAddPricesLicenseAllocationPriceConversionRateConfigConversionRateType) IsKnown() bool {
	switch r {
	case PlanNewPlanVersionParamsAddPricesLicenseAllocationPriceConversionRateConfigConversionRateTypeUnit, PlanNewPlanVersionParamsAddPricesLicenseAllocationPriceConversionRateConfigConversionRateTypeTiered:
		return true
	}
	return false
}

// New plan price request body params.
type PlanNewPlanVersionParamsAddPricesPrice struct {
	// The cadence to bill for this price on.
	Cadence param.Field[PlanNewPlanVersionParamsAddPricesPriceCadence] `json:"cadence" api:"required"`
	// The id of the item the price will be associated with.
	ItemID param.Field[string] `json:"item_id" api:"required"`
	// The pricing model type
	ModelType param.Field[PlanNewPlanVersionParamsAddPricesPriceModelType] `json:"model_type" api:"required"`
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
	Currency                   param.Field[string]      `json:"currency"`
	DailyCreditAllowanceConfig param.Field[interface{}] `json:"daily_credit_allowance_config"`
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
	GroupedTieredMatrixConfig         param.Field[interface{}] `json:"grouped_tiered_matrix_config"`
	GroupedTieredPackageConfig        param.Field[interface{}] `json:"grouped_tiered_package_config"`
	GroupedWithMeteredMinimumConfig   param.Field[interface{}] `json:"grouped_with_metered_minimum_config"`
	GroupedWithMinMaxThresholdsConfig param.Field[interface{}] `json:"grouped_with_min_max_thresholds_config"`
	GroupedWithProratedMinimumConfig  param.Field[interface{}] `json:"grouped_with_prorated_minimum_config"`
	// The property used to group this price on an invoice
	InvoiceGroupingKey param.Field[string] `json:"invoice_grouping_key"`
	// Within each billing cycle, specifies the cadence at which invoices are produced.
	// If unspecified, a single invoice is produced per billing cycle.
	InvoicingCycleConfiguration param.Field[shared.NewBillingCycleConfigurationParam] `json:"invoicing_cycle_configuration"`
	// The ID of the license type to associate with this price. On a usage price this
	// also marks the price as eligible to draw down from that license type's
	// allocation; a usage price created without it is billed normally. Usage prices
	// with a license type must use the `unit` model, and only draw down when their
	// currency matches the allocation's.
	LicenseTypeID param.Field[string] `json:"license_type_id"`
	// Configuration for matrix pricing
	MatrixConfig param.Field[shared.MatrixConfigParam] `json:"matrix_config"`
	// Configuration for matrix_with_allocation pricing
	MatrixWithAllocationConfig         param.Field[shared.MatrixWithAllocationConfigParam] `json:"matrix_with_allocation_config"`
	MatrixWithDisplayNameConfig        param.Field[interface{}]                            `json:"matrix_with_display_name_config"`
	MatrixWithThresholdDiscountsConfig param.Field[interface{}]                            `json:"matrix_with_threshold_discounts_config"`
	MaxGroupTieredPackageConfig        param.Field[interface{}]                            `json:"max_group_tiered_package_config"`
	Metadata                           param.Field[interface{}]                            `json:"metadata"`
	MeteredAllowanceConfig             param.Field[interface{}]                            `json:"metered_allowance_config"`
	MinimumCompositeConfig             param.Field[interface{}]                            `json:"minimum_composite_config"`
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
	TieredConfig                     param.Field[shared.TieredConfigParam] `json:"tiered_config"`
	TieredMatrixWithAllocationConfig param.Field[interface{}]              `json:"tiered_matrix_with_allocation_config"`
	TieredPackageConfig              param.Field[interface{}]              `json:"tiered_package_config"`
	TieredPackageWithMinimumConfig   param.Field[interface{}]              `json:"tiered_package_with_minimum_config"`
	TieredWithMinimumConfig          param.Field[interface{}]              `json:"tiered_with_minimum_config"`
	TieredWithProrationConfig        param.Field[interface{}]              `json:"tiered_with_proration_config"`
	// Configuration for unit pricing
	UnitConfig               param.Field[shared.UnitConfigParam] `json:"unit_config"`
	UnitWithMultiplierConfig param.Field[interface{}]            `json:"unit_with_multiplier_config"`
	UnitWithPercentConfig    param.Field[interface{}]            `json:"unit_with_percent_config"`
	UnitWithProrationConfig  param.Field[interface{}]            `json:"unit_with_proration_config"`
}

func (r PlanNewPlanVersionParamsAddPricesPrice) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r PlanNewPlanVersionParamsAddPricesPrice) ImplementsPlanNewPlanVersionParamsAddPricesPriceUnion() {
}

// New plan price request body params.
//
// Satisfied by [shared.NewPlanUnitPriceParam], [shared.NewPlanTieredPriceParam],
// [shared.NewPlanBulkPriceParam],
// [PlanNewPlanVersionParamsAddPricesPriceNewPlanBulkWithFiltersPrice],
// [shared.NewPlanPackagePriceParam], [shared.NewPlanMatrixPriceParam],
// [shared.NewPlanThresholdTotalAmountPriceParam],
// [shared.NewPlanTieredPackagePriceParam],
// [shared.NewPlanTieredWithMinimumPriceParam],
// [shared.NewPlanGroupedTieredPriceParam],
// [PlanNewPlanVersionParamsAddPricesPriceNewPlanGroupedTieredMatrixPrice],
// [shared.NewPlanTieredPackageWithMinimumPriceParam],
// [shared.NewPlanPackageWithAllocationPriceParam],
// [shared.NewPlanUnitWithPercentPriceParam],
// [shared.NewPlanMatrixWithAllocationPriceParam],
// [PlanNewPlanVersionParamsAddPricesPriceNewPlanTieredMatrixWithAllocationPrice],
// [PlanNewPlanVersionParamsAddPricesPriceNewPlanMatrixWithThresholdDiscountsPrice],
// [PlanNewPlanVersionParamsAddPricesPriceNewPlanTieredWithProrationPrice],
// [shared.NewPlanUnitWithProrationPriceParam],
// [shared.NewPlanGroupedAllocationPriceParam],
// [shared.NewPlanBulkWithProrationPriceParam],
// [shared.NewPlanGroupedWithProratedMinimumPriceParam],
// [shared.NewPlanGroupedWithMeteredMinimumPriceParam],
// [PlanNewPlanVersionParamsAddPricesPriceNewPlanGroupedWithMinMaxThresholdsPrice],
// [shared.NewPlanMatrixWithDisplayNamePriceParam],
// [shared.NewPlanGroupedTieredPackagePriceParam],
// [shared.NewPlanMaxGroupTieredPackagePriceParam],
// [shared.NewPlanScalableMatrixWithUnitPricingPriceParam],
// [shared.NewPlanScalableMatrixWithTieredPricingPriceParam],
// [shared.NewPlanCumulativeGroupedBulkPriceParam],
// [PlanNewPlanVersionParamsAddPricesPriceNewPlanCumulativeGroupedAllocationPrice],
// [PlanNewPlanVersionParamsAddPricesPriceNewPlanDailyCreditAllowancePrice],
// [PlanNewPlanVersionParamsAddPricesPriceNewPlanMeteredAllowancePrice],
// [shared.NewPlanMinimumCompositePriceParam],
// [PlanNewPlanVersionParamsAddPricesPriceNewPlanPercentCompositePrice],
// [PlanNewPlanVersionParamsAddPricesPriceNewPlanEventOutputPrice],
// [PlanNewPlanVersionParamsAddPricesPriceNewPlanUnitWithMultiplierPrice],
// [PlanNewPlanVersionParamsAddPricesPrice].
type PlanNewPlanVersionParamsAddPricesPriceUnion interface {
	ImplementsPlanNewPlanVersionParamsAddPricesPriceUnion()
}

type PlanNewPlanVersionParamsAddPricesPriceNewPlanBulkWithFiltersPrice struct {
	// Configuration for bulk_with_filters pricing
	BulkWithFiltersConfig param.Field[PlanNewPlanVersionParamsAddPricesPriceNewPlanBulkWithFiltersPriceBulkWithFiltersConfig] `json:"bulk_with_filters_config" api:"required"`
	// The cadence to bill for this price on.
	Cadence param.Field[PlanNewPlanVersionParamsAddPricesPriceNewPlanBulkWithFiltersPriceCadence] `json:"cadence" api:"required"`
	// The id of the item the price will be associated with.
	ItemID param.Field[string] `json:"item_id" api:"required"`
	// The pricing model type
	ModelType param.Field[PlanNewPlanVersionParamsAddPricesPriceNewPlanBulkWithFiltersPriceModelType] `json:"model_type" api:"required"`
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
	ConversionRateConfig param.Field[PlanNewPlanVersionParamsAddPricesPriceNewPlanBulkWithFiltersPriceConversionRateConfigUnion] `json:"conversion_rate_config"`
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
	// The ID of the license type to associate with this price. On a usage price this
	// also marks the price as eligible to draw down from that license type's
	// allocation; a usage price created without it is billed normally. Usage prices
	// with a license type must use the `unit` model, and only draw down when their
	// currency matches the allocation's.
	LicenseTypeID param.Field[string] `json:"license_type_id"`
	// User-specified key/value pairs for the resource. Individual keys can be removed
	// by setting the value to `null`, and the entire metadata mapping can be cleared
	// by setting `metadata` to `null`.
	Metadata param.Field[map[string]string] `json:"metadata"`
	// A transient ID that can be used to reference this price when adding adjustments
	// in the same API call.
	ReferenceID param.Field[string] `json:"reference_id"`
}

func (r PlanNewPlanVersionParamsAddPricesPriceNewPlanBulkWithFiltersPrice) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r PlanNewPlanVersionParamsAddPricesPriceNewPlanBulkWithFiltersPrice) ImplementsPlanNewPlanVersionParamsAddPricesPriceUnion() {
}

// Configuration for bulk_with_filters pricing
type PlanNewPlanVersionParamsAddPricesPriceNewPlanBulkWithFiltersPriceBulkWithFiltersConfig struct {
	// Property filters to apply (all must match)
	Filters param.Field[[]PlanNewPlanVersionParamsAddPricesPriceNewPlanBulkWithFiltersPriceBulkWithFiltersConfigFilter] `json:"filters" api:"required"`
	// Bulk tiers for rating based on total usage volume
	Tiers param.Field[[]PlanNewPlanVersionParamsAddPricesPriceNewPlanBulkWithFiltersPriceBulkWithFiltersConfigTier] `json:"tiers" api:"required"`
}

func (r PlanNewPlanVersionParamsAddPricesPriceNewPlanBulkWithFiltersPriceBulkWithFiltersConfig) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Configuration for a single property filter
type PlanNewPlanVersionParamsAddPricesPriceNewPlanBulkWithFiltersPriceBulkWithFiltersConfigFilter struct {
	// Event property key to filter on
	PropertyKey param.Field[string] `json:"property_key" api:"required"`
	// Event property value to match
	PropertyValue param.Field[string] `json:"property_value" api:"required"`
}

func (r PlanNewPlanVersionParamsAddPricesPriceNewPlanBulkWithFiltersPriceBulkWithFiltersConfigFilter) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Configuration for a single bulk pricing tier
type PlanNewPlanVersionParamsAddPricesPriceNewPlanBulkWithFiltersPriceBulkWithFiltersConfigTier struct {
	// Amount per unit
	UnitAmount param.Field[string] `json:"unit_amount" api:"required"`
	// The lower bound for this tier
	TierLowerBound param.Field[string] `json:"tier_lower_bound"`
}

func (r PlanNewPlanVersionParamsAddPricesPriceNewPlanBulkWithFiltersPriceBulkWithFiltersConfigTier) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// The cadence to bill for this price on.
type PlanNewPlanVersionParamsAddPricesPriceNewPlanBulkWithFiltersPriceCadence string

const (
	PlanNewPlanVersionParamsAddPricesPriceNewPlanBulkWithFiltersPriceCadenceOneTime    PlanNewPlanVersionParamsAddPricesPriceNewPlanBulkWithFiltersPriceCadence = "one_time"
	PlanNewPlanVersionParamsAddPricesPriceNewPlanBulkWithFiltersPriceCadenceMonthly    PlanNewPlanVersionParamsAddPricesPriceNewPlanBulkWithFiltersPriceCadence = "monthly"
	PlanNewPlanVersionParamsAddPricesPriceNewPlanBulkWithFiltersPriceCadenceQuarterly  PlanNewPlanVersionParamsAddPricesPriceNewPlanBulkWithFiltersPriceCadence = "quarterly"
	PlanNewPlanVersionParamsAddPricesPriceNewPlanBulkWithFiltersPriceCadenceSemiAnnual PlanNewPlanVersionParamsAddPricesPriceNewPlanBulkWithFiltersPriceCadence = "semi_annual"
	PlanNewPlanVersionParamsAddPricesPriceNewPlanBulkWithFiltersPriceCadenceAnnual     PlanNewPlanVersionParamsAddPricesPriceNewPlanBulkWithFiltersPriceCadence = "annual"
	PlanNewPlanVersionParamsAddPricesPriceNewPlanBulkWithFiltersPriceCadenceCustom     PlanNewPlanVersionParamsAddPricesPriceNewPlanBulkWithFiltersPriceCadence = "custom"
)

func (r PlanNewPlanVersionParamsAddPricesPriceNewPlanBulkWithFiltersPriceCadence) IsKnown() bool {
	switch r {
	case PlanNewPlanVersionParamsAddPricesPriceNewPlanBulkWithFiltersPriceCadenceOneTime, PlanNewPlanVersionParamsAddPricesPriceNewPlanBulkWithFiltersPriceCadenceMonthly, PlanNewPlanVersionParamsAddPricesPriceNewPlanBulkWithFiltersPriceCadenceQuarterly, PlanNewPlanVersionParamsAddPricesPriceNewPlanBulkWithFiltersPriceCadenceSemiAnnual, PlanNewPlanVersionParamsAddPricesPriceNewPlanBulkWithFiltersPriceCadenceAnnual, PlanNewPlanVersionParamsAddPricesPriceNewPlanBulkWithFiltersPriceCadenceCustom:
		return true
	}
	return false
}

// The pricing model type
type PlanNewPlanVersionParamsAddPricesPriceNewPlanBulkWithFiltersPriceModelType string

const (
	PlanNewPlanVersionParamsAddPricesPriceNewPlanBulkWithFiltersPriceModelTypeBulkWithFilters PlanNewPlanVersionParamsAddPricesPriceNewPlanBulkWithFiltersPriceModelType = "bulk_with_filters"
)

func (r PlanNewPlanVersionParamsAddPricesPriceNewPlanBulkWithFiltersPriceModelType) IsKnown() bool {
	switch r {
	case PlanNewPlanVersionParamsAddPricesPriceNewPlanBulkWithFiltersPriceModelTypeBulkWithFilters:
		return true
	}
	return false
}

type PlanNewPlanVersionParamsAddPricesPriceNewPlanBulkWithFiltersPriceConversionRateConfig struct {
	ConversionRateType param.Field[PlanNewPlanVersionParamsAddPricesPriceNewPlanBulkWithFiltersPriceConversionRateConfigConversionRateType] `json:"conversion_rate_type" api:"required"`
	TieredConfig       param.Field[shared.ConversionRateTieredConfigParam]                                                                  `json:"tiered_config"`
	UnitConfig         param.Field[shared.ConversionRateUnitConfigParam]                                                                    `json:"unit_config"`
}

func (r PlanNewPlanVersionParamsAddPricesPriceNewPlanBulkWithFiltersPriceConversionRateConfig) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r PlanNewPlanVersionParamsAddPricesPriceNewPlanBulkWithFiltersPriceConversionRateConfig) ImplementsPlanNewPlanVersionParamsAddPricesPriceNewPlanBulkWithFiltersPriceConversionRateConfigUnion() {
}

// Satisfied by [shared.UnitConversionRateConfigParam],
// [shared.TieredConversionRateConfigParam],
// [PlanNewPlanVersionParamsAddPricesPriceNewPlanBulkWithFiltersPriceConversionRateConfig].
type PlanNewPlanVersionParamsAddPricesPriceNewPlanBulkWithFiltersPriceConversionRateConfigUnion interface {
	ImplementsPlanNewPlanVersionParamsAddPricesPriceNewPlanBulkWithFiltersPriceConversionRateConfigUnion()
}

type PlanNewPlanVersionParamsAddPricesPriceNewPlanBulkWithFiltersPriceConversionRateConfigConversionRateType string

const (
	PlanNewPlanVersionParamsAddPricesPriceNewPlanBulkWithFiltersPriceConversionRateConfigConversionRateTypeUnit   PlanNewPlanVersionParamsAddPricesPriceNewPlanBulkWithFiltersPriceConversionRateConfigConversionRateType = "unit"
	PlanNewPlanVersionParamsAddPricesPriceNewPlanBulkWithFiltersPriceConversionRateConfigConversionRateTypeTiered PlanNewPlanVersionParamsAddPricesPriceNewPlanBulkWithFiltersPriceConversionRateConfigConversionRateType = "tiered"
)

func (r PlanNewPlanVersionParamsAddPricesPriceNewPlanBulkWithFiltersPriceConversionRateConfigConversionRateType) IsKnown() bool {
	switch r {
	case PlanNewPlanVersionParamsAddPricesPriceNewPlanBulkWithFiltersPriceConversionRateConfigConversionRateTypeUnit, PlanNewPlanVersionParamsAddPricesPriceNewPlanBulkWithFiltersPriceConversionRateConfigConversionRateTypeTiered:
		return true
	}
	return false
}

type PlanNewPlanVersionParamsAddPricesPriceNewPlanGroupedTieredMatrixPrice struct {
	// The cadence to bill for this price on.
	Cadence param.Field[PlanNewPlanVersionParamsAddPricesPriceNewPlanGroupedTieredMatrixPriceCadence] `json:"cadence" api:"required"`
	// Configuration for grouped_tiered_matrix pricing
	GroupedTieredMatrixConfig param.Field[PlanNewPlanVersionParamsAddPricesPriceNewPlanGroupedTieredMatrixPriceGroupedTieredMatrixConfig] `json:"grouped_tiered_matrix_config" api:"required"`
	// The id of the item the price will be associated with.
	ItemID param.Field[string] `json:"item_id" api:"required"`
	// The pricing model type
	ModelType param.Field[PlanNewPlanVersionParamsAddPricesPriceNewPlanGroupedTieredMatrixPriceModelType] `json:"model_type" api:"required"`
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
	ConversionRateConfig param.Field[PlanNewPlanVersionParamsAddPricesPriceNewPlanGroupedTieredMatrixPriceConversionRateConfigUnion] `json:"conversion_rate_config"`
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
	// The ID of the license type to associate with this price. On a usage price this
	// also marks the price as eligible to draw down from that license type's
	// allocation; a usage price created without it is billed normally. Usage prices
	// with a license type must use the `unit` model, and only draw down when their
	// currency matches the allocation's.
	LicenseTypeID param.Field[string] `json:"license_type_id"`
	// User-specified key/value pairs for the resource. Individual keys can be removed
	// by setting the value to `null`, and the entire metadata mapping can be cleared
	// by setting `metadata` to `null`.
	Metadata param.Field[map[string]string] `json:"metadata"`
	// A transient ID that can be used to reference this price when adding adjustments
	// in the same API call.
	ReferenceID param.Field[string] `json:"reference_id"`
}

func (r PlanNewPlanVersionParamsAddPricesPriceNewPlanGroupedTieredMatrixPrice) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r PlanNewPlanVersionParamsAddPricesPriceNewPlanGroupedTieredMatrixPrice) ImplementsPlanNewPlanVersionParamsAddPricesPriceUnion() {
}

// The cadence to bill for this price on.
type PlanNewPlanVersionParamsAddPricesPriceNewPlanGroupedTieredMatrixPriceCadence string

const (
	PlanNewPlanVersionParamsAddPricesPriceNewPlanGroupedTieredMatrixPriceCadenceOneTime    PlanNewPlanVersionParamsAddPricesPriceNewPlanGroupedTieredMatrixPriceCadence = "one_time"
	PlanNewPlanVersionParamsAddPricesPriceNewPlanGroupedTieredMatrixPriceCadenceMonthly    PlanNewPlanVersionParamsAddPricesPriceNewPlanGroupedTieredMatrixPriceCadence = "monthly"
	PlanNewPlanVersionParamsAddPricesPriceNewPlanGroupedTieredMatrixPriceCadenceQuarterly  PlanNewPlanVersionParamsAddPricesPriceNewPlanGroupedTieredMatrixPriceCadence = "quarterly"
	PlanNewPlanVersionParamsAddPricesPriceNewPlanGroupedTieredMatrixPriceCadenceSemiAnnual PlanNewPlanVersionParamsAddPricesPriceNewPlanGroupedTieredMatrixPriceCadence = "semi_annual"
	PlanNewPlanVersionParamsAddPricesPriceNewPlanGroupedTieredMatrixPriceCadenceAnnual     PlanNewPlanVersionParamsAddPricesPriceNewPlanGroupedTieredMatrixPriceCadence = "annual"
	PlanNewPlanVersionParamsAddPricesPriceNewPlanGroupedTieredMatrixPriceCadenceCustom     PlanNewPlanVersionParamsAddPricesPriceNewPlanGroupedTieredMatrixPriceCadence = "custom"
)

func (r PlanNewPlanVersionParamsAddPricesPriceNewPlanGroupedTieredMatrixPriceCadence) IsKnown() bool {
	switch r {
	case PlanNewPlanVersionParamsAddPricesPriceNewPlanGroupedTieredMatrixPriceCadenceOneTime, PlanNewPlanVersionParamsAddPricesPriceNewPlanGroupedTieredMatrixPriceCadenceMonthly, PlanNewPlanVersionParamsAddPricesPriceNewPlanGroupedTieredMatrixPriceCadenceQuarterly, PlanNewPlanVersionParamsAddPricesPriceNewPlanGroupedTieredMatrixPriceCadenceSemiAnnual, PlanNewPlanVersionParamsAddPricesPriceNewPlanGroupedTieredMatrixPriceCadenceAnnual, PlanNewPlanVersionParamsAddPricesPriceNewPlanGroupedTieredMatrixPriceCadenceCustom:
		return true
	}
	return false
}

// Configuration for grouped_tiered_matrix pricing
type PlanNewPlanVersionParamsAddPricesPriceNewPlanGroupedTieredMatrixPriceGroupedTieredMatrixConfig struct {
	// Per unit rate for usage whose dimension value has no configured tiers
	DefaultUnitAmount param.Field[string] `json:"default_unit_amount" api:"required"`
	// The billable metric property used to group usage before tiering
	Dimension param.Field[string] `json:"dimension" api:"required"`
	// Graduated tiers keyed by dimension value; usage for a value is tiered only
	// against its own rows
	Tiers param.Field[[]PlanNewPlanVersionParamsAddPricesPriceNewPlanGroupedTieredMatrixPriceGroupedTieredMatrixConfigTier] `json:"tiers" api:"required"`
}

func (r PlanNewPlanVersionParamsAddPricesPriceNewPlanGroupedTieredMatrixPriceGroupedTieredMatrixConfig) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Configuration for a single tier scoped to a dimension value
type PlanNewPlanVersionParamsAddPricesPriceNewPlanGroupedTieredMatrixPriceGroupedTieredMatrixConfigTier struct {
	// The dimension value this tier applies to
	DimensionValue param.Field[string] `json:"dimension_value" api:"required"`
	TierLowerBound param.Field[string] `json:"tier_lower_bound" api:"required"`
	// Per unit amount
	UnitAmount param.Field[string] `json:"unit_amount" api:"required"`
}

func (r PlanNewPlanVersionParamsAddPricesPriceNewPlanGroupedTieredMatrixPriceGroupedTieredMatrixConfigTier) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// The pricing model type
type PlanNewPlanVersionParamsAddPricesPriceNewPlanGroupedTieredMatrixPriceModelType string

const (
	PlanNewPlanVersionParamsAddPricesPriceNewPlanGroupedTieredMatrixPriceModelTypeGroupedTieredMatrix PlanNewPlanVersionParamsAddPricesPriceNewPlanGroupedTieredMatrixPriceModelType = "grouped_tiered_matrix"
)

func (r PlanNewPlanVersionParamsAddPricesPriceNewPlanGroupedTieredMatrixPriceModelType) IsKnown() bool {
	switch r {
	case PlanNewPlanVersionParamsAddPricesPriceNewPlanGroupedTieredMatrixPriceModelTypeGroupedTieredMatrix:
		return true
	}
	return false
}

type PlanNewPlanVersionParamsAddPricesPriceNewPlanGroupedTieredMatrixPriceConversionRateConfig struct {
	ConversionRateType param.Field[PlanNewPlanVersionParamsAddPricesPriceNewPlanGroupedTieredMatrixPriceConversionRateConfigConversionRateType] `json:"conversion_rate_type" api:"required"`
	TieredConfig       param.Field[shared.ConversionRateTieredConfigParam]                                                                      `json:"tiered_config"`
	UnitConfig         param.Field[shared.ConversionRateUnitConfigParam]                                                                        `json:"unit_config"`
}

func (r PlanNewPlanVersionParamsAddPricesPriceNewPlanGroupedTieredMatrixPriceConversionRateConfig) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r PlanNewPlanVersionParamsAddPricesPriceNewPlanGroupedTieredMatrixPriceConversionRateConfig) ImplementsPlanNewPlanVersionParamsAddPricesPriceNewPlanGroupedTieredMatrixPriceConversionRateConfigUnion() {
}

// Satisfied by [shared.UnitConversionRateConfigParam],
// [shared.TieredConversionRateConfigParam],
// [PlanNewPlanVersionParamsAddPricesPriceNewPlanGroupedTieredMatrixPriceConversionRateConfig].
type PlanNewPlanVersionParamsAddPricesPriceNewPlanGroupedTieredMatrixPriceConversionRateConfigUnion interface {
	ImplementsPlanNewPlanVersionParamsAddPricesPriceNewPlanGroupedTieredMatrixPriceConversionRateConfigUnion()
}

type PlanNewPlanVersionParamsAddPricesPriceNewPlanGroupedTieredMatrixPriceConversionRateConfigConversionRateType string

const (
	PlanNewPlanVersionParamsAddPricesPriceNewPlanGroupedTieredMatrixPriceConversionRateConfigConversionRateTypeUnit   PlanNewPlanVersionParamsAddPricesPriceNewPlanGroupedTieredMatrixPriceConversionRateConfigConversionRateType = "unit"
	PlanNewPlanVersionParamsAddPricesPriceNewPlanGroupedTieredMatrixPriceConversionRateConfigConversionRateTypeTiered PlanNewPlanVersionParamsAddPricesPriceNewPlanGroupedTieredMatrixPriceConversionRateConfigConversionRateType = "tiered"
)

func (r PlanNewPlanVersionParamsAddPricesPriceNewPlanGroupedTieredMatrixPriceConversionRateConfigConversionRateType) IsKnown() bool {
	switch r {
	case PlanNewPlanVersionParamsAddPricesPriceNewPlanGroupedTieredMatrixPriceConversionRateConfigConversionRateTypeUnit, PlanNewPlanVersionParamsAddPricesPriceNewPlanGroupedTieredMatrixPriceConversionRateConfigConversionRateTypeTiered:
		return true
	}
	return false
}

type PlanNewPlanVersionParamsAddPricesPriceNewPlanTieredMatrixWithAllocationPrice struct {
	// The cadence to bill for this price on.
	Cadence param.Field[PlanNewPlanVersionParamsAddPricesPriceNewPlanTieredMatrixWithAllocationPriceCadence] `json:"cadence" api:"required"`
	// The id of the item the price will be associated with.
	ItemID param.Field[string] `json:"item_id" api:"required"`
	// The pricing model type
	ModelType param.Field[PlanNewPlanVersionParamsAddPricesPriceNewPlanTieredMatrixWithAllocationPriceModelType] `json:"model_type" api:"required"`
	// The name of the price.
	Name param.Field[string] `json:"name" api:"required"`
	// Configuration for tiered_matrix_with_allocation pricing
	TieredMatrixWithAllocationConfig param.Field[PlanNewPlanVersionParamsAddPricesPriceNewPlanTieredMatrixWithAllocationPriceTieredMatrixWithAllocationConfig] `json:"tiered_matrix_with_allocation_config" api:"required"`
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
	ConversionRateConfig param.Field[PlanNewPlanVersionParamsAddPricesPriceNewPlanTieredMatrixWithAllocationPriceConversionRateConfigUnion] `json:"conversion_rate_config"`
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
	// The ID of the license type to associate with this price. On a usage price this
	// also marks the price as eligible to draw down from that license type's
	// allocation; a usage price created without it is billed normally. Usage prices
	// with a license type must use the `unit` model, and only draw down when their
	// currency matches the allocation's.
	LicenseTypeID param.Field[string] `json:"license_type_id"`
	// User-specified key/value pairs for the resource. Individual keys can be removed
	// by setting the value to `null`, and the entire metadata mapping can be cleared
	// by setting `metadata` to `null`.
	Metadata param.Field[map[string]string] `json:"metadata"`
	// A transient ID that can be used to reference this price when adding adjustments
	// in the same API call.
	ReferenceID param.Field[string] `json:"reference_id"`
}

func (r PlanNewPlanVersionParamsAddPricesPriceNewPlanTieredMatrixWithAllocationPrice) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r PlanNewPlanVersionParamsAddPricesPriceNewPlanTieredMatrixWithAllocationPrice) ImplementsPlanNewPlanVersionParamsAddPricesPriceUnion() {
}

// The cadence to bill for this price on.
type PlanNewPlanVersionParamsAddPricesPriceNewPlanTieredMatrixWithAllocationPriceCadence string

const (
	PlanNewPlanVersionParamsAddPricesPriceNewPlanTieredMatrixWithAllocationPriceCadenceOneTime    PlanNewPlanVersionParamsAddPricesPriceNewPlanTieredMatrixWithAllocationPriceCadence = "one_time"
	PlanNewPlanVersionParamsAddPricesPriceNewPlanTieredMatrixWithAllocationPriceCadenceMonthly    PlanNewPlanVersionParamsAddPricesPriceNewPlanTieredMatrixWithAllocationPriceCadence = "monthly"
	PlanNewPlanVersionParamsAddPricesPriceNewPlanTieredMatrixWithAllocationPriceCadenceQuarterly  PlanNewPlanVersionParamsAddPricesPriceNewPlanTieredMatrixWithAllocationPriceCadence = "quarterly"
	PlanNewPlanVersionParamsAddPricesPriceNewPlanTieredMatrixWithAllocationPriceCadenceSemiAnnual PlanNewPlanVersionParamsAddPricesPriceNewPlanTieredMatrixWithAllocationPriceCadence = "semi_annual"
	PlanNewPlanVersionParamsAddPricesPriceNewPlanTieredMatrixWithAllocationPriceCadenceAnnual     PlanNewPlanVersionParamsAddPricesPriceNewPlanTieredMatrixWithAllocationPriceCadence = "annual"
	PlanNewPlanVersionParamsAddPricesPriceNewPlanTieredMatrixWithAllocationPriceCadenceCustom     PlanNewPlanVersionParamsAddPricesPriceNewPlanTieredMatrixWithAllocationPriceCadence = "custom"
)

func (r PlanNewPlanVersionParamsAddPricesPriceNewPlanTieredMatrixWithAllocationPriceCadence) IsKnown() bool {
	switch r {
	case PlanNewPlanVersionParamsAddPricesPriceNewPlanTieredMatrixWithAllocationPriceCadenceOneTime, PlanNewPlanVersionParamsAddPricesPriceNewPlanTieredMatrixWithAllocationPriceCadenceMonthly, PlanNewPlanVersionParamsAddPricesPriceNewPlanTieredMatrixWithAllocationPriceCadenceQuarterly, PlanNewPlanVersionParamsAddPricesPriceNewPlanTieredMatrixWithAllocationPriceCadenceSemiAnnual, PlanNewPlanVersionParamsAddPricesPriceNewPlanTieredMatrixWithAllocationPriceCadenceAnnual, PlanNewPlanVersionParamsAddPricesPriceNewPlanTieredMatrixWithAllocationPriceCadenceCustom:
		return true
	}
	return false
}

// The pricing model type
type PlanNewPlanVersionParamsAddPricesPriceNewPlanTieredMatrixWithAllocationPriceModelType string

const (
	PlanNewPlanVersionParamsAddPricesPriceNewPlanTieredMatrixWithAllocationPriceModelTypeTieredMatrixWithAllocation PlanNewPlanVersionParamsAddPricesPriceNewPlanTieredMatrixWithAllocationPriceModelType = "tiered_matrix_with_allocation"
)

func (r PlanNewPlanVersionParamsAddPricesPriceNewPlanTieredMatrixWithAllocationPriceModelType) IsKnown() bool {
	switch r {
	case PlanNewPlanVersionParamsAddPricesPriceNewPlanTieredMatrixWithAllocationPriceModelTypeTieredMatrixWithAllocation:
		return true
	}
	return false
}

// Configuration for tiered_matrix_with_allocation pricing
type PlanNewPlanVersionParamsAddPricesPriceNewPlanTieredMatrixWithAllocationPriceTieredMatrixWithAllocationConfig struct {
	// Usage allocation, pooled across all matrix cells
	Allocation param.Field[string] `json:"allocation" api:"required"`
	// Per unit rate for usage whose matrix cell has no configured tiers
	DefaultUnitAmount param.Field[string] `json:"default_unit_amount" api:"required"`
	// One or two event property values to evaluate matrix cells by
	Dimensions param.Field[[]string] `json:"dimensions" api:"required"`
	// Graduated tiers keyed by matrix cell; usage in a cell is tiered only against its
	// own rows
	Tiers param.Field[[]PlanNewPlanVersionParamsAddPricesPriceNewPlanTieredMatrixWithAllocationPriceTieredMatrixWithAllocationConfigTier] `json:"tiers" api:"required"`
}

func (r PlanNewPlanVersionParamsAddPricesPriceNewPlanTieredMatrixWithAllocationPriceTieredMatrixWithAllocationConfig) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Configuration for a single tier scoped to one matrix cell
type PlanNewPlanVersionParamsAddPricesPriceNewPlanTieredMatrixWithAllocationPriceTieredMatrixWithAllocationConfigTier struct {
	// The matrix cell this tier applies to, as one or two dimension values
	DimensionValues param.Field[[]string] `json:"dimension_values" api:"required"`
	// Exclusive tier starting value. The tier runs up to and including the next bound
	// configured for the same matrix cell.
	TierLowerBound param.Field[string] `json:"tier_lower_bound" api:"required"`
	// Per unit amount
	UnitAmount param.Field[string] `json:"unit_amount" api:"required"`
}

func (r PlanNewPlanVersionParamsAddPricesPriceNewPlanTieredMatrixWithAllocationPriceTieredMatrixWithAllocationConfigTier) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type PlanNewPlanVersionParamsAddPricesPriceNewPlanTieredMatrixWithAllocationPriceConversionRateConfig struct {
	ConversionRateType param.Field[PlanNewPlanVersionParamsAddPricesPriceNewPlanTieredMatrixWithAllocationPriceConversionRateConfigConversionRateType] `json:"conversion_rate_type" api:"required"`
	TieredConfig       param.Field[shared.ConversionRateTieredConfigParam]                                                                             `json:"tiered_config"`
	UnitConfig         param.Field[shared.ConversionRateUnitConfigParam]                                                                               `json:"unit_config"`
}

func (r PlanNewPlanVersionParamsAddPricesPriceNewPlanTieredMatrixWithAllocationPriceConversionRateConfig) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r PlanNewPlanVersionParamsAddPricesPriceNewPlanTieredMatrixWithAllocationPriceConversionRateConfig) ImplementsPlanNewPlanVersionParamsAddPricesPriceNewPlanTieredMatrixWithAllocationPriceConversionRateConfigUnion() {
}

// Satisfied by [shared.UnitConversionRateConfigParam],
// [shared.TieredConversionRateConfigParam],
// [PlanNewPlanVersionParamsAddPricesPriceNewPlanTieredMatrixWithAllocationPriceConversionRateConfig].
type PlanNewPlanVersionParamsAddPricesPriceNewPlanTieredMatrixWithAllocationPriceConversionRateConfigUnion interface {
	ImplementsPlanNewPlanVersionParamsAddPricesPriceNewPlanTieredMatrixWithAllocationPriceConversionRateConfigUnion()
}

type PlanNewPlanVersionParamsAddPricesPriceNewPlanTieredMatrixWithAllocationPriceConversionRateConfigConversionRateType string

const (
	PlanNewPlanVersionParamsAddPricesPriceNewPlanTieredMatrixWithAllocationPriceConversionRateConfigConversionRateTypeUnit   PlanNewPlanVersionParamsAddPricesPriceNewPlanTieredMatrixWithAllocationPriceConversionRateConfigConversionRateType = "unit"
	PlanNewPlanVersionParamsAddPricesPriceNewPlanTieredMatrixWithAllocationPriceConversionRateConfigConversionRateTypeTiered PlanNewPlanVersionParamsAddPricesPriceNewPlanTieredMatrixWithAllocationPriceConversionRateConfigConversionRateType = "tiered"
)

func (r PlanNewPlanVersionParamsAddPricesPriceNewPlanTieredMatrixWithAllocationPriceConversionRateConfigConversionRateType) IsKnown() bool {
	switch r {
	case PlanNewPlanVersionParamsAddPricesPriceNewPlanTieredMatrixWithAllocationPriceConversionRateConfigConversionRateTypeUnit, PlanNewPlanVersionParamsAddPricesPriceNewPlanTieredMatrixWithAllocationPriceConversionRateConfigConversionRateTypeTiered:
		return true
	}
	return false
}

type PlanNewPlanVersionParamsAddPricesPriceNewPlanMatrixWithThresholdDiscountsPrice struct {
	// The cadence to bill for this price on.
	Cadence param.Field[PlanNewPlanVersionParamsAddPricesPriceNewPlanMatrixWithThresholdDiscountsPriceCadence] `json:"cadence" api:"required"`
	// The id of the item the price will be associated with.
	ItemID param.Field[string] `json:"item_id" api:"required"`
	// Configuration for matrix_with_threshold_discounts pricing
	MatrixWithThresholdDiscountsConfig param.Field[PlanNewPlanVersionParamsAddPricesPriceNewPlanMatrixWithThresholdDiscountsPriceMatrixWithThresholdDiscountsConfig] `json:"matrix_with_threshold_discounts_config" api:"required"`
	// The pricing model type
	ModelType param.Field[PlanNewPlanVersionParamsAddPricesPriceNewPlanMatrixWithThresholdDiscountsPriceModelType] `json:"model_type" api:"required"`
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
	ConversionRateConfig param.Field[PlanNewPlanVersionParamsAddPricesPriceNewPlanMatrixWithThresholdDiscountsPriceConversionRateConfigUnion] `json:"conversion_rate_config"`
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
	// The ID of the license type to associate with this price. On a usage price this
	// also marks the price as eligible to draw down from that license type's
	// allocation; a usage price created without it is billed normally. Usage prices
	// with a license type must use the `unit` model, and only draw down when their
	// currency matches the allocation's.
	LicenseTypeID param.Field[string] `json:"license_type_id"`
	// User-specified key/value pairs for the resource. Individual keys can be removed
	// by setting the value to `null`, and the entire metadata mapping can be cleared
	// by setting `metadata` to `null`.
	Metadata param.Field[map[string]string] `json:"metadata"`
	// A transient ID that can be used to reference this price when adding adjustments
	// in the same API call.
	ReferenceID param.Field[string] `json:"reference_id"`
}

func (r PlanNewPlanVersionParamsAddPricesPriceNewPlanMatrixWithThresholdDiscountsPrice) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r PlanNewPlanVersionParamsAddPricesPriceNewPlanMatrixWithThresholdDiscountsPrice) ImplementsPlanNewPlanVersionParamsAddPricesPriceUnion() {
}

// The cadence to bill for this price on.
type PlanNewPlanVersionParamsAddPricesPriceNewPlanMatrixWithThresholdDiscountsPriceCadence string

const (
	PlanNewPlanVersionParamsAddPricesPriceNewPlanMatrixWithThresholdDiscountsPriceCadenceOneTime    PlanNewPlanVersionParamsAddPricesPriceNewPlanMatrixWithThresholdDiscountsPriceCadence = "one_time"
	PlanNewPlanVersionParamsAddPricesPriceNewPlanMatrixWithThresholdDiscountsPriceCadenceMonthly    PlanNewPlanVersionParamsAddPricesPriceNewPlanMatrixWithThresholdDiscountsPriceCadence = "monthly"
	PlanNewPlanVersionParamsAddPricesPriceNewPlanMatrixWithThresholdDiscountsPriceCadenceQuarterly  PlanNewPlanVersionParamsAddPricesPriceNewPlanMatrixWithThresholdDiscountsPriceCadence = "quarterly"
	PlanNewPlanVersionParamsAddPricesPriceNewPlanMatrixWithThresholdDiscountsPriceCadenceSemiAnnual PlanNewPlanVersionParamsAddPricesPriceNewPlanMatrixWithThresholdDiscountsPriceCadence = "semi_annual"
	PlanNewPlanVersionParamsAddPricesPriceNewPlanMatrixWithThresholdDiscountsPriceCadenceAnnual     PlanNewPlanVersionParamsAddPricesPriceNewPlanMatrixWithThresholdDiscountsPriceCadence = "annual"
	PlanNewPlanVersionParamsAddPricesPriceNewPlanMatrixWithThresholdDiscountsPriceCadenceCustom     PlanNewPlanVersionParamsAddPricesPriceNewPlanMatrixWithThresholdDiscountsPriceCadence = "custom"
)

func (r PlanNewPlanVersionParamsAddPricesPriceNewPlanMatrixWithThresholdDiscountsPriceCadence) IsKnown() bool {
	switch r {
	case PlanNewPlanVersionParamsAddPricesPriceNewPlanMatrixWithThresholdDiscountsPriceCadenceOneTime, PlanNewPlanVersionParamsAddPricesPriceNewPlanMatrixWithThresholdDiscountsPriceCadenceMonthly, PlanNewPlanVersionParamsAddPricesPriceNewPlanMatrixWithThresholdDiscountsPriceCadenceQuarterly, PlanNewPlanVersionParamsAddPricesPriceNewPlanMatrixWithThresholdDiscountsPriceCadenceSemiAnnual, PlanNewPlanVersionParamsAddPricesPriceNewPlanMatrixWithThresholdDiscountsPriceCadenceAnnual, PlanNewPlanVersionParamsAddPricesPriceNewPlanMatrixWithThresholdDiscountsPriceCadenceCustom:
		return true
	}
	return false
}

// Configuration for matrix_with_threshold_discounts pricing
type PlanNewPlanVersionParamsAddPricesPriceNewPlanMatrixWithThresholdDiscountsPriceMatrixWithThresholdDiscountsConfig struct {
	// Unit price used for usage that does not match any defined matrix cell.
	DefaultUnitAmount param.Field[string] `json:"default_unit_amount" api:"required"`
	// First matrix dimension key.
	FirstDimension param.Field[string] `json:"first_dimension" api:"required"`
	// Per-cell unit prices.
	MatrixValues param.Field[[]PlanNewPlanVersionParamsAddPricesPriceNewPlanMatrixWithThresholdDiscountsPriceMatrixWithThresholdDiscountsConfigMatrixValue] `json:"matrix_values" api:"required"`
	// Optional second matrix dimension key.
	SecondDimension         param.Field[string]                                                                                                                                   `json:"second_dimension"`
	ThresholdDiscountGroups param.Field[[]PlanNewPlanVersionParamsAddPricesPriceNewPlanMatrixWithThresholdDiscountsPriceMatrixWithThresholdDiscountsConfigThresholdDiscountGroup] `json:"threshold_discount_groups"`
}

func (r PlanNewPlanVersionParamsAddPricesPriceNewPlanMatrixWithThresholdDiscountsPriceMatrixWithThresholdDiscountsConfig) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type PlanNewPlanVersionParamsAddPricesPriceNewPlanMatrixWithThresholdDiscountsPriceMatrixWithThresholdDiscountsConfigMatrixValue struct {
	FirstDimensionValue  param.Field[string] `json:"first_dimension_value" api:"required"`
	UnitAmount           param.Field[string] `json:"unit_amount" api:"required"`
	SecondDimensionValue param.Field[string] `json:"second_dimension_value"`
}

func (r PlanNewPlanVersionParamsAddPricesPriceNewPlanMatrixWithThresholdDiscountsPriceMatrixWithThresholdDiscountsConfigMatrixValue) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type PlanNewPlanVersionParamsAddPricesPriceNewPlanMatrixWithThresholdDiscountsPriceMatrixWithThresholdDiscountsConfigThresholdDiscountGroup struct {
	// Discount rate applied to spend above the threshold.
	AboveThresholdDiscountPercentage param.Field[string] `json:"above_threshold_discount_percentage" api:"required"`
	// Discount rate applied to spend at or below the threshold. Set to 0 for no
	// baseline discount.
	BelowThresholdDiscountPercentage param.Field[string] `json:"below_threshold_discount_percentage" api:"required"`
	// Semicolon-separated list of matrix cell coordinates targeted by this group. Each
	// coordinate is `first,second` when the matrix has two dimensions, or just `first`
	// for a single-dimension matrix. Example: `blue,circle;green,triangle`.
	CellCoordinates param.Field[string] `json:"cell_coordinates" api:"required"`
	ThresholdAmount param.Field[string] `json:"threshold_amount" api:"required"`
	Description     param.Field[string] `json:"description"`
}

func (r PlanNewPlanVersionParamsAddPricesPriceNewPlanMatrixWithThresholdDiscountsPriceMatrixWithThresholdDiscountsConfigThresholdDiscountGroup) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// The pricing model type
type PlanNewPlanVersionParamsAddPricesPriceNewPlanMatrixWithThresholdDiscountsPriceModelType string

const (
	PlanNewPlanVersionParamsAddPricesPriceNewPlanMatrixWithThresholdDiscountsPriceModelTypeMatrixWithThresholdDiscounts PlanNewPlanVersionParamsAddPricesPriceNewPlanMatrixWithThresholdDiscountsPriceModelType = "matrix_with_threshold_discounts"
)

func (r PlanNewPlanVersionParamsAddPricesPriceNewPlanMatrixWithThresholdDiscountsPriceModelType) IsKnown() bool {
	switch r {
	case PlanNewPlanVersionParamsAddPricesPriceNewPlanMatrixWithThresholdDiscountsPriceModelTypeMatrixWithThresholdDiscounts:
		return true
	}
	return false
}

type PlanNewPlanVersionParamsAddPricesPriceNewPlanMatrixWithThresholdDiscountsPriceConversionRateConfig struct {
	ConversionRateType param.Field[PlanNewPlanVersionParamsAddPricesPriceNewPlanMatrixWithThresholdDiscountsPriceConversionRateConfigConversionRateType] `json:"conversion_rate_type" api:"required"`
	TieredConfig       param.Field[shared.ConversionRateTieredConfigParam]                                                                               `json:"tiered_config"`
	UnitConfig         param.Field[shared.ConversionRateUnitConfigParam]                                                                                 `json:"unit_config"`
}

func (r PlanNewPlanVersionParamsAddPricesPriceNewPlanMatrixWithThresholdDiscountsPriceConversionRateConfig) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r PlanNewPlanVersionParamsAddPricesPriceNewPlanMatrixWithThresholdDiscountsPriceConversionRateConfig) ImplementsPlanNewPlanVersionParamsAddPricesPriceNewPlanMatrixWithThresholdDiscountsPriceConversionRateConfigUnion() {
}

// Satisfied by [shared.UnitConversionRateConfigParam],
// [shared.TieredConversionRateConfigParam],
// [PlanNewPlanVersionParamsAddPricesPriceNewPlanMatrixWithThresholdDiscountsPriceConversionRateConfig].
type PlanNewPlanVersionParamsAddPricesPriceNewPlanMatrixWithThresholdDiscountsPriceConversionRateConfigUnion interface {
	ImplementsPlanNewPlanVersionParamsAddPricesPriceNewPlanMatrixWithThresholdDiscountsPriceConversionRateConfigUnion()
}

type PlanNewPlanVersionParamsAddPricesPriceNewPlanMatrixWithThresholdDiscountsPriceConversionRateConfigConversionRateType string

const (
	PlanNewPlanVersionParamsAddPricesPriceNewPlanMatrixWithThresholdDiscountsPriceConversionRateConfigConversionRateTypeUnit   PlanNewPlanVersionParamsAddPricesPriceNewPlanMatrixWithThresholdDiscountsPriceConversionRateConfigConversionRateType = "unit"
	PlanNewPlanVersionParamsAddPricesPriceNewPlanMatrixWithThresholdDiscountsPriceConversionRateConfigConversionRateTypeTiered PlanNewPlanVersionParamsAddPricesPriceNewPlanMatrixWithThresholdDiscountsPriceConversionRateConfigConversionRateType = "tiered"
)

func (r PlanNewPlanVersionParamsAddPricesPriceNewPlanMatrixWithThresholdDiscountsPriceConversionRateConfigConversionRateType) IsKnown() bool {
	switch r {
	case PlanNewPlanVersionParamsAddPricesPriceNewPlanMatrixWithThresholdDiscountsPriceConversionRateConfigConversionRateTypeUnit, PlanNewPlanVersionParamsAddPricesPriceNewPlanMatrixWithThresholdDiscountsPriceConversionRateConfigConversionRateTypeTiered:
		return true
	}
	return false
}

type PlanNewPlanVersionParamsAddPricesPriceNewPlanTieredWithProrationPrice struct {
	// The cadence to bill for this price on.
	Cadence param.Field[PlanNewPlanVersionParamsAddPricesPriceNewPlanTieredWithProrationPriceCadence] `json:"cadence" api:"required"`
	// The id of the item the price will be associated with.
	ItemID param.Field[string] `json:"item_id" api:"required"`
	// The pricing model type
	ModelType param.Field[PlanNewPlanVersionParamsAddPricesPriceNewPlanTieredWithProrationPriceModelType] `json:"model_type" api:"required"`
	// The name of the price.
	Name param.Field[string] `json:"name" api:"required"`
	// Configuration for tiered_with_proration pricing
	TieredWithProrationConfig param.Field[PlanNewPlanVersionParamsAddPricesPriceNewPlanTieredWithProrationPriceTieredWithProrationConfig] `json:"tiered_with_proration_config" api:"required"`
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
	ConversionRateConfig param.Field[PlanNewPlanVersionParamsAddPricesPriceNewPlanTieredWithProrationPriceConversionRateConfigUnion] `json:"conversion_rate_config"`
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
	// The ID of the license type to associate with this price. On a usage price this
	// also marks the price as eligible to draw down from that license type's
	// allocation; a usage price created without it is billed normally. Usage prices
	// with a license type must use the `unit` model, and only draw down when their
	// currency matches the allocation's.
	LicenseTypeID param.Field[string] `json:"license_type_id"`
	// User-specified key/value pairs for the resource. Individual keys can be removed
	// by setting the value to `null`, and the entire metadata mapping can be cleared
	// by setting `metadata` to `null`.
	Metadata param.Field[map[string]string] `json:"metadata"`
	// A transient ID that can be used to reference this price when adding adjustments
	// in the same API call.
	ReferenceID param.Field[string] `json:"reference_id"`
}

func (r PlanNewPlanVersionParamsAddPricesPriceNewPlanTieredWithProrationPrice) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r PlanNewPlanVersionParamsAddPricesPriceNewPlanTieredWithProrationPrice) ImplementsPlanNewPlanVersionParamsAddPricesPriceUnion() {
}

// The cadence to bill for this price on.
type PlanNewPlanVersionParamsAddPricesPriceNewPlanTieredWithProrationPriceCadence string

const (
	PlanNewPlanVersionParamsAddPricesPriceNewPlanTieredWithProrationPriceCadenceOneTime    PlanNewPlanVersionParamsAddPricesPriceNewPlanTieredWithProrationPriceCadence = "one_time"
	PlanNewPlanVersionParamsAddPricesPriceNewPlanTieredWithProrationPriceCadenceMonthly    PlanNewPlanVersionParamsAddPricesPriceNewPlanTieredWithProrationPriceCadence = "monthly"
	PlanNewPlanVersionParamsAddPricesPriceNewPlanTieredWithProrationPriceCadenceQuarterly  PlanNewPlanVersionParamsAddPricesPriceNewPlanTieredWithProrationPriceCadence = "quarterly"
	PlanNewPlanVersionParamsAddPricesPriceNewPlanTieredWithProrationPriceCadenceSemiAnnual PlanNewPlanVersionParamsAddPricesPriceNewPlanTieredWithProrationPriceCadence = "semi_annual"
	PlanNewPlanVersionParamsAddPricesPriceNewPlanTieredWithProrationPriceCadenceAnnual     PlanNewPlanVersionParamsAddPricesPriceNewPlanTieredWithProrationPriceCadence = "annual"
	PlanNewPlanVersionParamsAddPricesPriceNewPlanTieredWithProrationPriceCadenceCustom     PlanNewPlanVersionParamsAddPricesPriceNewPlanTieredWithProrationPriceCadence = "custom"
)

func (r PlanNewPlanVersionParamsAddPricesPriceNewPlanTieredWithProrationPriceCadence) IsKnown() bool {
	switch r {
	case PlanNewPlanVersionParamsAddPricesPriceNewPlanTieredWithProrationPriceCadenceOneTime, PlanNewPlanVersionParamsAddPricesPriceNewPlanTieredWithProrationPriceCadenceMonthly, PlanNewPlanVersionParamsAddPricesPriceNewPlanTieredWithProrationPriceCadenceQuarterly, PlanNewPlanVersionParamsAddPricesPriceNewPlanTieredWithProrationPriceCadenceSemiAnnual, PlanNewPlanVersionParamsAddPricesPriceNewPlanTieredWithProrationPriceCadenceAnnual, PlanNewPlanVersionParamsAddPricesPriceNewPlanTieredWithProrationPriceCadenceCustom:
		return true
	}
	return false
}

// The pricing model type
type PlanNewPlanVersionParamsAddPricesPriceNewPlanTieredWithProrationPriceModelType string

const (
	PlanNewPlanVersionParamsAddPricesPriceNewPlanTieredWithProrationPriceModelTypeTieredWithProration PlanNewPlanVersionParamsAddPricesPriceNewPlanTieredWithProrationPriceModelType = "tiered_with_proration"
)

func (r PlanNewPlanVersionParamsAddPricesPriceNewPlanTieredWithProrationPriceModelType) IsKnown() bool {
	switch r {
	case PlanNewPlanVersionParamsAddPricesPriceNewPlanTieredWithProrationPriceModelTypeTieredWithProration:
		return true
	}
	return false
}

// Configuration for tiered_with_proration pricing
type PlanNewPlanVersionParamsAddPricesPriceNewPlanTieredWithProrationPriceTieredWithProrationConfig struct {
	// Tiers for rating based on total usage quantities into the specified tier with
	// proration
	Tiers param.Field[[]PlanNewPlanVersionParamsAddPricesPriceNewPlanTieredWithProrationPriceTieredWithProrationConfigTier] `json:"tiers" api:"required"`
}

func (r PlanNewPlanVersionParamsAddPricesPriceNewPlanTieredWithProrationPriceTieredWithProrationConfig) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Configuration for a single tiered with proration tier
type PlanNewPlanVersionParamsAddPricesPriceNewPlanTieredWithProrationPriceTieredWithProrationConfigTier struct {
	// Inclusive tier starting value
	TierLowerBound param.Field[string] `json:"tier_lower_bound" api:"required"`
	// Amount per unit
	UnitAmount param.Field[string] `json:"unit_amount" api:"required"`
}

func (r PlanNewPlanVersionParamsAddPricesPriceNewPlanTieredWithProrationPriceTieredWithProrationConfigTier) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type PlanNewPlanVersionParamsAddPricesPriceNewPlanTieredWithProrationPriceConversionRateConfig struct {
	ConversionRateType param.Field[PlanNewPlanVersionParamsAddPricesPriceNewPlanTieredWithProrationPriceConversionRateConfigConversionRateType] `json:"conversion_rate_type" api:"required"`
	TieredConfig       param.Field[shared.ConversionRateTieredConfigParam]                                                                      `json:"tiered_config"`
	UnitConfig         param.Field[shared.ConversionRateUnitConfigParam]                                                                        `json:"unit_config"`
}

func (r PlanNewPlanVersionParamsAddPricesPriceNewPlanTieredWithProrationPriceConversionRateConfig) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r PlanNewPlanVersionParamsAddPricesPriceNewPlanTieredWithProrationPriceConversionRateConfig) ImplementsPlanNewPlanVersionParamsAddPricesPriceNewPlanTieredWithProrationPriceConversionRateConfigUnion() {
}

// Satisfied by [shared.UnitConversionRateConfigParam],
// [shared.TieredConversionRateConfigParam],
// [PlanNewPlanVersionParamsAddPricesPriceNewPlanTieredWithProrationPriceConversionRateConfig].
type PlanNewPlanVersionParamsAddPricesPriceNewPlanTieredWithProrationPriceConversionRateConfigUnion interface {
	ImplementsPlanNewPlanVersionParamsAddPricesPriceNewPlanTieredWithProrationPriceConversionRateConfigUnion()
}

type PlanNewPlanVersionParamsAddPricesPriceNewPlanTieredWithProrationPriceConversionRateConfigConversionRateType string

const (
	PlanNewPlanVersionParamsAddPricesPriceNewPlanTieredWithProrationPriceConversionRateConfigConversionRateTypeUnit   PlanNewPlanVersionParamsAddPricesPriceNewPlanTieredWithProrationPriceConversionRateConfigConversionRateType = "unit"
	PlanNewPlanVersionParamsAddPricesPriceNewPlanTieredWithProrationPriceConversionRateConfigConversionRateTypeTiered PlanNewPlanVersionParamsAddPricesPriceNewPlanTieredWithProrationPriceConversionRateConfigConversionRateType = "tiered"
)

func (r PlanNewPlanVersionParamsAddPricesPriceNewPlanTieredWithProrationPriceConversionRateConfigConversionRateType) IsKnown() bool {
	switch r {
	case PlanNewPlanVersionParamsAddPricesPriceNewPlanTieredWithProrationPriceConversionRateConfigConversionRateTypeUnit, PlanNewPlanVersionParamsAddPricesPriceNewPlanTieredWithProrationPriceConversionRateConfigConversionRateTypeTiered:
		return true
	}
	return false
}

type PlanNewPlanVersionParamsAddPricesPriceNewPlanGroupedWithMinMaxThresholdsPrice struct {
	// The cadence to bill for this price on.
	Cadence param.Field[PlanNewPlanVersionParamsAddPricesPriceNewPlanGroupedWithMinMaxThresholdsPriceCadence] `json:"cadence" api:"required"`
	// Configuration for grouped_with_min_max_thresholds pricing
	GroupedWithMinMaxThresholdsConfig param.Field[PlanNewPlanVersionParamsAddPricesPriceNewPlanGroupedWithMinMaxThresholdsPriceGroupedWithMinMaxThresholdsConfig] `json:"grouped_with_min_max_thresholds_config" api:"required"`
	// The id of the item the price will be associated with.
	ItemID param.Field[string] `json:"item_id" api:"required"`
	// The pricing model type
	ModelType param.Field[PlanNewPlanVersionParamsAddPricesPriceNewPlanGroupedWithMinMaxThresholdsPriceModelType] `json:"model_type" api:"required"`
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
	ConversionRateConfig param.Field[PlanNewPlanVersionParamsAddPricesPriceNewPlanGroupedWithMinMaxThresholdsPriceConversionRateConfigUnion] `json:"conversion_rate_config"`
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
	// The ID of the license type to associate with this price. On a usage price this
	// also marks the price as eligible to draw down from that license type's
	// allocation; a usage price created without it is billed normally. Usage prices
	// with a license type must use the `unit` model, and only draw down when their
	// currency matches the allocation's.
	LicenseTypeID param.Field[string] `json:"license_type_id"`
	// User-specified key/value pairs for the resource. Individual keys can be removed
	// by setting the value to `null`, and the entire metadata mapping can be cleared
	// by setting `metadata` to `null`.
	Metadata param.Field[map[string]string] `json:"metadata"`
	// A transient ID that can be used to reference this price when adding adjustments
	// in the same API call.
	ReferenceID param.Field[string] `json:"reference_id"`
}

func (r PlanNewPlanVersionParamsAddPricesPriceNewPlanGroupedWithMinMaxThresholdsPrice) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r PlanNewPlanVersionParamsAddPricesPriceNewPlanGroupedWithMinMaxThresholdsPrice) ImplementsPlanNewPlanVersionParamsAddPricesPriceUnion() {
}

// The cadence to bill for this price on.
type PlanNewPlanVersionParamsAddPricesPriceNewPlanGroupedWithMinMaxThresholdsPriceCadence string

const (
	PlanNewPlanVersionParamsAddPricesPriceNewPlanGroupedWithMinMaxThresholdsPriceCadenceOneTime    PlanNewPlanVersionParamsAddPricesPriceNewPlanGroupedWithMinMaxThresholdsPriceCadence = "one_time"
	PlanNewPlanVersionParamsAddPricesPriceNewPlanGroupedWithMinMaxThresholdsPriceCadenceMonthly    PlanNewPlanVersionParamsAddPricesPriceNewPlanGroupedWithMinMaxThresholdsPriceCadence = "monthly"
	PlanNewPlanVersionParamsAddPricesPriceNewPlanGroupedWithMinMaxThresholdsPriceCadenceQuarterly  PlanNewPlanVersionParamsAddPricesPriceNewPlanGroupedWithMinMaxThresholdsPriceCadence = "quarterly"
	PlanNewPlanVersionParamsAddPricesPriceNewPlanGroupedWithMinMaxThresholdsPriceCadenceSemiAnnual PlanNewPlanVersionParamsAddPricesPriceNewPlanGroupedWithMinMaxThresholdsPriceCadence = "semi_annual"
	PlanNewPlanVersionParamsAddPricesPriceNewPlanGroupedWithMinMaxThresholdsPriceCadenceAnnual     PlanNewPlanVersionParamsAddPricesPriceNewPlanGroupedWithMinMaxThresholdsPriceCadence = "annual"
	PlanNewPlanVersionParamsAddPricesPriceNewPlanGroupedWithMinMaxThresholdsPriceCadenceCustom     PlanNewPlanVersionParamsAddPricesPriceNewPlanGroupedWithMinMaxThresholdsPriceCadence = "custom"
)

func (r PlanNewPlanVersionParamsAddPricesPriceNewPlanGroupedWithMinMaxThresholdsPriceCadence) IsKnown() bool {
	switch r {
	case PlanNewPlanVersionParamsAddPricesPriceNewPlanGroupedWithMinMaxThresholdsPriceCadenceOneTime, PlanNewPlanVersionParamsAddPricesPriceNewPlanGroupedWithMinMaxThresholdsPriceCadenceMonthly, PlanNewPlanVersionParamsAddPricesPriceNewPlanGroupedWithMinMaxThresholdsPriceCadenceQuarterly, PlanNewPlanVersionParamsAddPricesPriceNewPlanGroupedWithMinMaxThresholdsPriceCadenceSemiAnnual, PlanNewPlanVersionParamsAddPricesPriceNewPlanGroupedWithMinMaxThresholdsPriceCadenceAnnual, PlanNewPlanVersionParamsAddPricesPriceNewPlanGroupedWithMinMaxThresholdsPriceCadenceCustom:
		return true
	}
	return false
}

// Configuration for grouped_with_min_max_thresholds pricing
type PlanNewPlanVersionParamsAddPricesPriceNewPlanGroupedWithMinMaxThresholdsPriceGroupedWithMinMaxThresholdsConfig struct {
	// The event property used to group before applying thresholds
	GroupingKey param.Field[string] `json:"grouping_key" api:"required"`
	// The maximum amount to charge each group
	MaximumCharge param.Field[string] `json:"maximum_charge" api:"required"`
	// The minimum amount to charge each group, regardless of usage
	MinimumCharge param.Field[string] `json:"minimum_charge" api:"required"`
	// The base price charged per group
	PerUnitRate param.Field[string] `json:"per_unit_rate" api:"required"`
}

func (r PlanNewPlanVersionParamsAddPricesPriceNewPlanGroupedWithMinMaxThresholdsPriceGroupedWithMinMaxThresholdsConfig) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// The pricing model type
type PlanNewPlanVersionParamsAddPricesPriceNewPlanGroupedWithMinMaxThresholdsPriceModelType string

const (
	PlanNewPlanVersionParamsAddPricesPriceNewPlanGroupedWithMinMaxThresholdsPriceModelTypeGroupedWithMinMaxThresholds PlanNewPlanVersionParamsAddPricesPriceNewPlanGroupedWithMinMaxThresholdsPriceModelType = "grouped_with_min_max_thresholds"
)

func (r PlanNewPlanVersionParamsAddPricesPriceNewPlanGroupedWithMinMaxThresholdsPriceModelType) IsKnown() bool {
	switch r {
	case PlanNewPlanVersionParamsAddPricesPriceNewPlanGroupedWithMinMaxThresholdsPriceModelTypeGroupedWithMinMaxThresholds:
		return true
	}
	return false
}

type PlanNewPlanVersionParamsAddPricesPriceNewPlanGroupedWithMinMaxThresholdsPriceConversionRateConfig struct {
	ConversionRateType param.Field[PlanNewPlanVersionParamsAddPricesPriceNewPlanGroupedWithMinMaxThresholdsPriceConversionRateConfigConversionRateType] `json:"conversion_rate_type" api:"required"`
	TieredConfig       param.Field[shared.ConversionRateTieredConfigParam]                                                                              `json:"tiered_config"`
	UnitConfig         param.Field[shared.ConversionRateUnitConfigParam]                                                                                `json:"unit_config"`
}

func (r PlanNewPlanVersionParamsAddPricesPriceNewPlanGroupedWithMinMaxThresholdsPriceConversionRateConfig) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r PlanNewPlanVersionParamsAddPricesPriceNewPlanGroupedWithMinMaxThresholdsPriceConversionRateConfig) ImplementsPlanNewPlanVersionParamsAddPricesPriceNewPlanGroupedWithMinMaxThresholdsPriceConversionRateConfigUnion() {
}

// Satisfied by [shared.UnitConversionRateConfigParam],
// [shared.TieredConversionRateConfigParam],
// [PlanNewPlanVersionParamsAddPricesPriceNewPlanGroupedWithMinMaxThresholdsPriceConversionRateConfig].
type PlanNewPlanVersionParamsAddPricesPriceNewPlanGroupedWithMinMaxThresholdsPriceConversionRateConfigUnion interface {
	ImplementsPlanNewPlanVersionParamsAddPricesPriceNewPlanGroupedWithMinMaxThresholdsPriceConversionRateConfigUnion()
}

type PlanNewPlanVersionParamsAddPricesPriceNewPlanGroupedWithMinMaxThresholdsPriceConversionRateConfigConversionRateType string

const (
	PlanNewPlanVersionParamsAddPricesPriceNewPlanGroupedWithMinMaxThresholdsPriceConversionRateConfigConversionRateTypeUnit   PlanNewPlanVersionParamsAddPricesPriceNewPlanGroupedWithMinMaxThresholdsPriceConversionRateConfigConversionRateType = "unit"
	PlanNewPlanVersionParamsAddPricesPriceNewPlanGroupedWithMinMaxThresholdsPriceConversionRateConfigConversionRateTypeTiered PlanNewPlanVersionParamsAddPricesPriceNewPlanGroupedWithMinMaxThresholdsPriceConversionRateConfigConversionRateType = "tiered"
)

func (r PlanNewPlanVersionParamsAddPricesPriceNewPlanGroupedWithMinMaxThresholdsPriceConversionRateConfigConversionRateType) IsKnown() bool {
	switch r {
	case PlanNewPlanVersionParamsAddPricesPriceNewPlanGroupedWithMinMaxThresholdsPriceConversionRateConfigConversionRateTypeUnit, PlanNewPlanVersionParamsAddPricesPriceNewPlanGroupedWithMinMaxThresholdsPriceConversionRateConfigConversionRateTypeTiered:
		return true
	}
	return false
}

type PlanNewPlanVersionParamsAddPricesPriceNewPlanCumulativeGroupedAllocationPrice struct {
	// The cadence to bill for this price on.
	Cadence param.Field[PlanNewPlanVersionParamsAddPricesPriceNewPlanCumulativeGroupedAllocationPriceCadence] `json:"cadence" api:"required"`
	// Configuration for cumulative_grouped_allocation pricing
	CumulativeGroupedAllocationConfig param.Field[PlanNewPlanVersionParamsAddPricesPriceNewPlanCumulativeGroupedAllocationPriceCumulativeGroupedAllocationConfig] `json:"cumulative_grouped_allocation_config" api:"required"`
	// The id of the item the price will be associated with.
	ItemID param.Field[string] `json:"item_id" api:"required"`
	// The pricing model type
	ModelType param.Field[PlanNewPlanVersionParamsAddPricesPriceNewPlanCumulativeGroupedAllocationPriceModelType] `json:"model_type" api:"required"`
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
	ConversionRateConfig param.Field[PlanNewPlanVersionParamsAddPricesPriceNewPlanCumulativeGroupedAllocationPriceConversionRateConfigUnion] `json:"conversion_rate_config"`
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
	// The ID of the license type to associate with this price. On a usage price this
	// also marks the price as eligible to draw down from that license type's
	// allocation; a usage price created without it is billed normally. Usage prices
	// with a license type must use the `unit` model, and only draw down when their
	// currency matches the allocation's.
	LicenseTypeID param.Field[string] `json:"license_type_id"`
	// User-specified key/value pairs for the resource. Individual keys can be removed
	// by setting the value to `null`, and the entire metadata mapping can be cleared
	// by setting `metadata` to `null`.
	Metadata param.Field[map[string]string] `json:"metadata"`
	// A transient ID that can be used to reference this price when adding adjustments
	// in the same API call.
	ReferenceID param.Field[string] `json:"reference_id"`
}

func (r PlanNewPlanVersionParamsAddPricesPriceNewPlanCumulativeGroupedAllocationPrice) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r PlanNewPlanVersionParamsAddPricesPriceNewPlanCumulativeGroupedAllocationPrice) ImplementsPlanNewPlanVersionParamsAddPricesPriceUnion() {
}

// The cadence to bill for this price on.
type PlanNewPlanVersionParamsAddPricesPriceNewPlanCumulativeGroupedAllocationPriceCadence string

const (
	PlanNewPlanVersionParamsAddPricesPriceNewPlanCumulativeGroupedAllocationPriceCadenceOneTime    PlanNewPlanVersionParamsAddPricesPriceNewPlanCumulativeGroupedAllocationPriceCadence = "one_time"
	PlanNewPlanVersionParamsAddPricesPriceNewPlanCumulativeGroupedAllocationPriceCadenceMonthly    PlanNewPlanVersionParamsAddPricesPriceNewPlanCumulativeGroupedAllocationPriceCadence = "monthly"
	PlanNewPlanVersionParamsAddPricesPriceNewPlanCumulativeGroupedAllocationPriceCadenceQuarterly  PlanNewPlanVersionParamsAddPricesPriceNewPlanCumulativeGroupedAllocationPriceCadence = "quarterly"
	PlanNewPlanVersionParamsAddPricesPriceNewPlanCumulativeGroupedAllocationPriceCadenceSemiAnnual PlanNewPlanVersionParamsAddPricesPriceNewPlanCumulativeGroupedAllocationPriceCadence = "semi_annual"
	PlanNewPlanVersionParamsAddPricesPriceNewPlanCumulativeGroupedAllocationPriceCadenceAnnual     PlanNewPlanVersionParamsAddPricesPriceNewPlanCumulativeGroupedAllocationPriceCadence = "annual"
	PlanNewPlanVersionParamsAddPricesPriceNewPlanCumulativeGroupedAllocationPriceCadenceCustom     PlanNewPlanVersionParamsAddPricesPriceNewPlanCumulativeGroupedAllocationPriceCadence = "custom"
)

func (r PlanNewPlanVersionParamsAddPricesPriceNewPlanCumulativeGroupedAllocationPriceCadence) IsKnown() bool {
	switch r {
	case PlanNewPlanVersionParamsAddPricesPriceNewPlanCumulativeGroupedAllocationPriceCadenceOneTime, PlanNewPlanVersionParamsAddPricesPriceNewPlanCumulativeGroupedAllocationPriceCadenceMonthly, PlanNewPlanVersionParamsAddPricesPriceNewPlanCumulativeGroupedAllocationPriceCadenceQuarterly, PlanNewPlanVersionParamsAddPricesPriceNewPlanCumulativeGroupedAllocationPriceCadenceSemiAnnual, PlanNewPlanVersionParamsAddPricesPriceNewPlanCumulativeGroupedAllocationPriceCadenceAnnual, PlanNewPlanVersionParamsAddPricesPriceNewPlanCumulativeGroupedAllocationPriceCadenceCustom:
		return true
	}
	return false
}

// Configuration for cumulative_grouped_allocation pricing
type PlanNewPlanVersionParamsAddPricesPriceNewPlanCumulativeGroupedAllocationPriceCumulativeGroupedAllocationConfig struct {
	// The overall allocation across all groups
	CumulativeAllocation param.Field[string] `json:"cumulative_allocation" api:"required"`
	// The allocation per individual group
	GroupAllocation param.Field[string] `json:"group_allocation" api:"required"`
	// The event property used to group usage before applying allocations
	GroupingKey param.Field[string] `json:"grouping_key" api:"required"`
	// The amount to charge for each unit outside of the allocation
	UnitAmount param.Field[string] `json:"unit_amount" api:"required"`
}

func (r PlanNewPlanVersionParamsAddPricesPriceNewPlanCumulativeGroupedAllocationPriceCumulativeGroupedAllocationConfig) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// The pricing model type
type PlanNewPlanVersionParamsAddPricesPriceNewPlanCumulativeGroupedAllocationPriceModelType string

const (
	PlanNewPlanVersionParamsAddPricesPriceNewPlanCumulativeGroupedAllocationPriceModelTypeCumulativeGroupedAllocation PlanNewPlanVersionParamsAddPricesPriceNewPlanCumulativeGroupedAllocationPriceModelType = "cumulative_grouped_allocation"
)

func (r PlanNewPlanVersionParamsAddPricesPriceNewPlanCumulativeGroupedAllocationPriceModelType) IsKnown() bool {
	switch r {
	case PlanNewPlanVersionParamsAddPricesPriceNewPlanCumulativeGroupedAllocationPriceModelTypeCumulativeGroupedAllocation:
		return true
	}
	return false
}

type PlanNewPlanVersionParamsAddPricesPriceNewPlanCumulativeGroupedAllocationPriceConversionRateConfig struct {
	ConversionRateType param.Field[PlanNewPlanVersionParamsAddPricesPriceNewPlanCumulativeGroupedAllocationPriceConversionRateConfigConversionRateType] `json:"conversion_rate_type" api:"required"`
	TieredConfig       param.Field[shared.ConversionRateTieredConfigParam]                                                                              `json:"tiered_config"`
	UnitConfig         param.Field[shared.ConversionRateUnitConfigParam]                                                                                `json:"unit_config"`
}

func (r PlanNewPlanVersionParamsAddPricesPriceNewPlanCumulativeGroupedAllocationPriceConversionRateConfig) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r PlanNewPlanVersionParamsAddPricesPriceNewPlanCumulativeGroupedAllocationPriceConversionRateConfig) ImplementsPlanNewPlanVersionParamsAddPricesPriceNewPlanCumulativeGroupedAllocationPriceConversionRateConfigUnion() {
}

// Satisfied by [shared.UnitConversionRateConfigParam],
// [shared.TieredConversionRateConfigParam],
// [PlanNewPlanVersionParamsAddPricesPriceNewPlanCumulativeGroupedAllocationPriceConversionRateConfig].
type PlanNewPlanVersionParamsAddPricesPriceNewPlanCumulativeGroupedAllocationPriceConversionRateConfigUnion interface {
	ImplementsPlanNewPlanVersionParamsAddPricesPriceNewPlanCumulativeGroupedAllocationPriceConversionRateConfigUnion()
}

type PlanNewPlanVersionParamsAddPricesPriceNewPlanCumulativeGroupedAllocationPriceConversionRateConfigConversionRateType string

const (
	PlanNewPlanVersionParamsAddPricesPriceNewPlanCumulativeGroupedAllocationPriceConversionRateConfigConversionRateTypeUnit   PlanNewPlanVersionParamsAddPricesPriceNewPlanCumulativeGroupedAllocationPriceConversionRateConfigConversionRateType = "unit"
	PlanNewPlanVersionParamsAddPricesPriceNewPlanCumulativeGroupedAllocationPriceConversionRateConfigConversionRateTypeTiered PlanNewPlanVersionParamsAddPricesPriceNewPlanCumulativeGroupedAllocationPriceConversionRateConfigConversionRateType = "tiered"
)

func (r PlanNewPlanVersionParamsAddPricesPriceNewPlanCumulativeGroupedAllocationPriceConversionRateConfigConversionRateType) IsKnown() bool {
	switch r {
	case PlanNewPlanVersionParamsAddPricesPriceNewPlanCumulativeGroupedAllocationPriceConversionRateConfigConversionRateTypeUnit, PlanNewPlanVersionParamsAddPricesPriceNewPlanCumulativeGroupedAllocationPriceConversionRateConfigConversionRateTypeTiered:
		return true
	}
	return false
}

type PlanNewPlanVersionParamsAddPricesPriceNewPlanDailyCreditAllowancePrice struct {
	// The cadence to bill for this price on.
	Cadence param.Field[PlanNewPlanVersionParamsAddPricesPriceNewPlanDailyCreditAllowancePriceCadence] `json:"cadence" api:"required"`
	// Configuration for daily_credit_allowance pricing
	DailyCreditAllowanceConfig param.Field[PlanNewPlanVersionParamsAddPricesPriceNewPlanDailyCreditAllowancePriceDailyCreditAllowanceConfig] `json:"daily_credit_allowance_config" api:"required"`
	// The id of the item the price will be associated with.
	ItemID param.Field[string] `json:"item_id" api:"required"`
	// The pricing model type
	ModelType param.Field[PlanNewPlanVersionParamsAddPricesPriceNewPlanDailyCreditAllowancePriceModelType] `json:"model_type" api:"required"`
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
	ConversionRateConfig param.Field[PlanNewPlanVersionParamsAddPricesPriceNewPlanDailyCreditAllowancePriceConversionRateConfigUnion] `json:"conversion_rate_config"`
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
	// The ID of the license type to associate with this price. On a usage price this
	// also marks the price as eligible to draw down from that license type's
	// allocation; a usage price created without it is billed normally. Usage prices
	// with a license type must use the `unit` model, and only draw down when their
	// currency matches the allocation's.
	LicenseTypeID param.Field[string] `json:"license_type_id"`
	// User-specified key/value pairs for the resource. Individual keys can be removed
	// by setting the value to `null`, and the entire metadata mapping can be cleared
	// by setting `metadata` to `null`.
	Metadata param.Field[map[string]string] `json:"metadata"`
	// A transient ID that can be used to reference this price when adding adjustments
	// in the same API call.
	ReferenceID param.Field[string] `json:"reference_id"`
}

func (r PlanNewPlanVersionParamsAddPricesPriceNewPlanDailyCreditAllowancePrice) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r PlanNewPlanVersionParamsAddPricesPriceNewPlanDailyCreditAllowancePrice) ImplementsPlanNewPlanVersionParamsAddPricesPriceUnion() {
}

// The cadence to bill for this price on.
type PlanNewPlanVersionParamsAddPricesPriceNewPlanDailyCreditAllowancePriceCadence string

const (
	PlanNewPlanVersionParamsAddPricesPriceNewPlanDailyCreditAllowancePriceCadenceOneTime    PlanNewPlanVersionParamsAddPricesPriceNewPlanDailyCreditAllowancePriceCadence = "one_time"
	PlanNewPlanVersionParamsAddPricesPriceNewPlanDailyCreditAllowancePriceCadenceMonthly    PlanNewPlanVersionParamsAddPricesPriceNewPlanDailyCreditAllowancePriceCadence = "monthly"
	PlanNewPlanVersionParamsAddPricesPriceNewPlanDailyCreditAllowancePriceCadenceQuarterly  PlanNewPlanVersionParamsAddPricesPriceNewPlanDailyCreditAllowancePriceCadence = "quarterly"
	PlanNewPlanVersionParamsAddPricesPriceNewPlanDailyCreditAllowancePriceCadenceSemiAnnual PlanNewPlanVersionParamsAddPricesPriceNewPlanDailyCreditAllowancePriceCadence = "semi_annual"
	PlanNewPlanVersionParamsAddPricesPriceNewPlanDailyCreditAllowancePriceCadenceAnnual     PlanNewPlanVersionParamsAddPricesPriceNewPlanDailyCreditAllowancePriceCadence = "annual"
	PlanNewPlanVersionParamsAddPricesPriceNewPlanDailyCreditAllowancePriceCadenceCustom     PlanNewPlanVersionParamsAddPricesPriceNewPlanDailyCreditAllowancePriceCadence = "custom"
)

func (r PlanNewPlanVersionParamsAddPricesPriceNewPlanDailyCreditAllowancePriceCadence) IsKnown() bool {
	switch r {
	case PlanNewPlanVersionParamsAddPricesPriceNewPlanDailyCreditAllowancePriceCadenceOneTime, PlanNewPlanVersionParamsAddPricesPriceNewPlanDailyCreditAllowancePriceCadenceMonthly, PlanNewPlanVersionParamsAddPricesPriceNewPlanDailyCreditAllowancePriceCadenceQuarterly, PlanNewPlanVersionParamsAddPricesPriceNewPlanDailyCreditAllowancePriceCadenceSemiAnnual, PlanNewPlanVersionParamsAddPricesPriceNewPlanDailyCreditAllowancePriceCadenceAnnual, PlanNewPlanVersionParamsAddPricesPriceNewPlanDailyCreditAllowancePriceCadenceCustom:
		return true
	}
	return false
}

// Configuration for daily_credit_allowance pricing
type PlanNewPlanVersionParamsAddPricesPriceNewPlanDailyCreditAllowancePriceDailyCreditAllowanceConfig struct {
	// Credits granted per day. Lose-it-or-use-it; does not roll over.
	DailyAllowance param.Field[string] `json:"daily_allowance" api:"required"`
	// Default per-unit credit rate for any usage not bucketed into a specified
	// matrix_value
	DefaultUnitAmount param.Field[string] `json:"default_unit_amount" api:"required"`
	// One or two event property values to evaluate matrix groups by
	Dimensions param.Field[[]string] `json:"dimensions" api:"required"`
	// Event property whose value identifies the day bucket the event belongs to (e.g.
	// 'event_day' set to an ISO date string in the customer's timezone). The allowance
	// resets per distinct value of this property.
	EventDayProperty param.Field[string] `json:"event_day_property" api:"required"`
	// Per-dimension credit rates
	MatrixValues param.Field[[]PlanNewPlanVersionParamsAddPricesPriceNewPlanDailyCreditAllowancePriceDailyCreditAllowanceConfigMatrixValue] `json:"matrix_values" api:"required"`
}

func (r PlanNewPlanVersionParamsAddPricesPriceNewPlanDailyCreditAllowancePriceDailyCreditAllowanceConfig) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Per-dimension credit price for the daily credit allowance model.
type PlanNewPlanVersionParamsAddPricesPriceNewPlanDailyCreditAllowancePriceDailyCreditAllowanceConfigMatrixValue struct {
	// One or two matrix keys to filter usage to this value by. For example, ["model"]
	// could be used to apply a different credit rate to each AI model.
	DimensionValues param.Field[[]string] `json:"dimension_values" api:"required"`
	// Credits charged per unit of usage matching the specified dimension_values
	UnitAmount param.Field[string] `json:"unit_amount" api:"required"`
}

func (r PlanNewPlanVersionParamsAddPricesPriceNewPlanDailyCreditAllowancePriceDailyCreditAllowanceConfigMatrixValue) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// The pricing model type
type PlanNewPlanVersionParamsAddPricesPriceNewPlanDailyCreditAllowancePriceModelType string

const (
	PlanNewPlanVersionParamsAddPricesPriceNewPlanDailyCreditAllowancePriceModelTypeDailyCreditAllowance PlanNewPlanVersionParamsAddPricesPriceNewPlanDailyCreditAllowancePriceModelType = "daily_credit_allowance"
)

func (r PlanNewPlanVersionParamsAddPricesPriceNewPlanDailyCreditAllowancePriceModelType) IsKnown() bool {
	switch r {
	case PlanNewPlanVersionParamsAddPricesPriceNewPlanDailyCreditAllowancePriceModelTypeDailyCreditAllowance:
		return true
	}
	return false
}

type PlanNewPlanVersionParamsAddPricesPriceNewPlanDailyCreditAllowancePriceConversionRateConfig struct {
	ConversionRateType param.Field[PlanNewPlanVersionParamsAddPricesPriceNewPlanDailyCreditAllowancePriceConversionRateConfigConversionRateType] `json:"conversion_rate_type" api:"required"`
	TieredConfig       param.Field[shared.ConversionRateTieredConfigParam]                                                                       `json:"tiered_config"`
	UnitConfig         param.Field[shared.ConversionRateUnitConfigParam]                                                                         `json:"unit_config"`
}

func (r PlanNewPlanVersionParamsAddPricesPriceNewPlanDailyCreditAllowancePriceConversionRateConfig) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r PlanNewPlanVersionParamsAddPricesPriceNewPlanDailyCreditAllowancePriceConversionRateConfig) ImplementsPlanNewPlanVersionParamsAddPricesPriceNewPlanDailyCreditAllowancePriceConversionRateConfigUnion() {
}

// Satisfied by [shared.UnitConversionRateConfigParam],
// [shared.TieredConversionRateConfigParam],
// [PlanNewPlanVersionParamsAddPricesPriceNewPlanDailyCreditAllowancePriceConversionRateConfig].
type PlanNewPlanVersionParamsAddPricesPriceNewPlanDailyCreditAllowancePriceConversionRateConfigUnion interface {
	ImplementsPlanNewPlanVersionParamsAddPricesPriceNewPlanDailyCreditAllowancePriceConversionRateConfigUnion()
}

type PlanNewPlanVersionParamsAddPricesPriceNewPlanDailyCreditAllowancePriceConversionRateConfigConversionRateType string

const (
	PlanNewPlanVersionParamsAddPricesPriceNewPlanDailyCreditAllowancePriceConversionRateConfigConversionRateTypeUnit   PlanNewPlanVersionParamsAddPricesPriceNewPlanDailyCreditAllowancePriceConversionRateConfigConversionRateType = "unit"
	PlanNewPlanVersionParamsAddPricesPriceNewPlanDailyCreditAllowancePriceConversionRateConfigConversionRateTypeTiered PlanNewPlanVersionParamsAddPricesPriceNewPlanDailyCreditAllowancePriceConversionRateConfigConversionRateType = "tiered"
)

func (r PlanNewPlanVersionParamsAddPricesPriceNewPlanDailyCreditAllowancePriceConversionRateConfigConversionRateType) IsKnown() bool {
	switch r {
	case PlanNewPlanVersionParamsAddPricesPriceNewPlanDailyCreditAllowancePriceConversionRateConfigConversionRateTypeUnit, PlanNewPlanVersionParamsAddPricesPriceNewPlanDailyCreditAllowancePriceConversionRateConfigConversionRateTypeTiered:
		return true
	}
	return false
}

type PlanNewPlanVersionParamsAddPricesPriceNewPlanMeteredAllowancePrice struct {
	// The cadence to bill for this price on.
	Cadence param.Field[PlanNewPlanVersionParamsAddPricesPriceNewPlanMeteredAllowancePriceCadence] `json:"cadence" api:"required"`
	// The id of the item the price will be associated with.
	ItemID param.Field[string] `json:"item_id" api:"required"`
	// Configuration for metered_allowance pricing
	MeteredAllowanceConfig param.Field[PlanNewPlanVersionParamsAddPricesPriceNewPlanMeteredAllowancePriceMeteredAllowanceConfig] `json:"metered_allowance_config" api:"required"`
	// The pricing model type
	ModelType param.Field[PlanNewPlanVersionParamsAddPricesPriceNewPlanMeteredAllowancePriceModelType] `json:"model_type" api:"required"`
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
	ConversionRateConfig param.Field[PlanNewPlanVersionParamsAddPricesPriceNewPlanMeteredAllowancePriceConversionRateConfigUnion] `json:"conversion_rate_config"`
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
	// The ID of the license type to associate with this price. On a usage price this
	// also marks the price as eligible to draw down from that license type's
	// allocation; a usage price created without it is billed normally. Usage prices
	// with a license type must use the `unit` model, and only draw down when their
	// currency matches the allocation's.
	LicenseTypeID param.Field[string] `json:"license_type_id"`
	// User-specified key/value pairs for the resource. Individual keys can be removed
	// by setting the value to `null`, and the entire metadata mapping can be cleared
	// by setting `metadata` to `null`.
	Metadata param.Field[map[string]string] `json:"metadata"`
	// A transient ID that can be used to reference this price when adding adjustments
	// in the same API call.
	ReferenceID param.Field[string] `json:"reference_id"`
}

func (r PlanNewPlanVersionParamsAddPricesPriceNewPlanMeteredAllowancePrice) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r PlanNewPlanVersionParamsAddPricesPriceNewPlanMeteredAllowancePrice) ImplementsPlanNewPlanVersionParamsAddPricesPriceUnion() {
}

// The cadence to bill for this price on.
type PlanNewPlanVersionParamsAddPricesPriceNewPlanMeteredAllowancePriceCadence string

const (
	PlanNewPlanVersionParamsAddPricesPriceNewPlanMeteredAllowancePriceCadenceOneTime    PlanNewPlanVersionParamsAddPricesPriceNewPlanMeteredAllowancePriceCadence = "one_time"
	PlanNewPlanVersionParamsAddPricesPriceNewPlanMeteredAllowancePriceCadenceMonthly    PlanNewPlanVersionParamsAddPricesPriceNewPlanMeteredAllowancePriceCadence = "monthly"
	PlanNewPlanVersionParamsAddPricesPriceNewPlanMeteredAllowancePriceCadenceQuarterly  PlanNewPlanVersionParamsAddPricesPriceNewPlanMeteredAllowancePriceCadence = "quarterly"
	PlanNewPlanVersionParamsAddPricesPriceNewPlanMeteredAllowancePriceCadenceSemiAnnual PlanNewPlanVersionParamsAddPricesPriceNewPlanMeteredAllowancePriceCadence = "semi_annual"
	PlanNewPlanVersionParamsAddPricesPriceNewPlanMeteredAllowancePriceCadenceAnnual     PlanNewPlanVersionParamsAddPricesPriceNewPlanMeteredAllowancePriceCadence = "annual"
	PlanNewPlanVersionParamsAddPricesPriceNewPlanMeteredAllowancePriceCadenceCustom     PlanNewPlanVersionParamsAddPricesPriceNewPlanMeteredAllowancePriceCadence = "custom"
)

func (r PlanNewPlanVersionParamsAddPricesPriceNewPlanMeteredAllowancePriceCadence) IsKnown() bool {
	switch r {
	case PlanNewPlanVersionParamsAddPricesPriceNewPlanMeteredAllowancePriceCadenceOneTime, PlanNewPlanVersionParamsAddPricesPriceNewPlanMeteredAllowancePriceCadenceMonthly, PlanNewPlanVersionParamsAddPricesPriceNewPlanMeteredAllowancePriceCadenceQuarterly, PlanNewPlanVersionParamsAddPricesPriceNewPlanMeteredAllowancePriceCadenceSemiAnnual, PlanNewPlanVersionParamsAddPricesPriceNewPlanMeteredAllowancePriceCadenceAnnual, PlanNewPlanVersionParamsAddPricesPriceNewPlanMeteredAllowancePriceCadenceCustom:
		return true
	}
	return false
}

// Configuration for metered_allowance pricing
type PlanNewPlanVersionParamsAddPricesPriceNewPlanMeteredAllowancePriceMeteredAllowanceConfig struct {
	// The grouping_key value whose summed quantity represents the allowance for this
	// period (e.g. 'storage_snapshot' emitting 3 × avg storage). Capped at consumption
	// — credit can never exceed actual usage.
	AllowanceGroupingValue param.Field[string] `json:"allowance_grouping_value" api:"required"`
	// The grouping_key value whose summed quantity represents consumption (e.g.
	// 'download'). Charged at unit_amount.
	ConsumptionGroupingValue param.Field[string] `json:"consumption_grouping_value" api:"required"`
	// Event property used to partition the metric into consumption and allowance
	// quantities (e.g. 'event_name'). The metric is queried with this key and the two
	// values below select which partition is which.
	GroupingKey param.Field[string] `json:"grouping_key" api:"required"`
	// Per-unit price applied to gross consumption and to the allowance credit.
	UnitAmount param.Field[string] `json:"unit_amount" api:"required"`
	// Sub-line label for the credit row (e.g. 'Up to 3x free egress').
	AllowanceDisplayName param.Field[string] `json:"allowance_display_name"`
	// Sub-line label for the gross consumption row (e.g. 'bytes gotten').
	ConsumptionDisplayName param.Field[string] `json:"consumption_display_name"`
}

func (r PlanNewPlanVersionParamsAddPricesPriceNewPlanMeteredAllowancePriceMeteredAllowanceConfig) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// The pricing model type
type PlanNewPlanVersionParamsAddPricesPriceNewPlanMeteredAllowancePriceModelType string

const (
	PlanNewPlanVersionParamsAddPricesPriceNewPlanMeteredAllowancePriceModelTypeMeteredAllowance PlanNewPlanVersionParamsAddPricesPriceNewPlanMeteredAllowancePriceModelType = "metered_allowance"
)

func (r PlanNewPlanVersionParamsAddPricesPriceNewPlanMeteredAllowancePriceModelType) IsKnown() bool {
	switch r {
	case PlanNewPlanVersionParamsAddPricesPriceNewPlanMeteredAllowancePriceModelTypeMeteredAllowance:
		return true
	}
	return false
}

type PlanNewPlanVersionParamsAddPricesPriceNewPlanMeteredAllowancePriceConversionRateConfig struct {
	ConversionRateType param.Field[PlanNewPlanVersionParamsAddPricesPriceNewPlanMeteredAllowancePriceConversionRateConfigConversionRateType] `json:"conversion_rate_type" api:"required"`
	TieredConfig       param.Field[shared.ConversionRateTieredConfigParam]                                                                   `json:"tiered_config"`
	UnitConfig         param.Field[shared.ConversionRateUnitConfigParam]                                                                     `json:"unit_config"`
}

func (r PlanNewPlanVersionParamsAddPricesPriceNewPlanMeteredAllowancePriceConversionRateConfig) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r PlanNewPlanVersionParamsAddPricesPriceNewPlanMeteredAllowancePriceConversionRateConfig) ImplementsPlanNewPlanVersionParamsAddPricesPriceNewPlanMeteredAllowancePriceConversionRateConfigUnion() {
}

// Satisfied by [shared.UnitConversionRateConfigParam],
// [shared.TieredConversionRateConfigParam],
// [PlanNewPlanVersionParamsAddPricesPriceNewPlanMeteredAllowancePriceConversionRateConfig].
type PlanNewPlanVersionParamsAddPricesPriceNewPlanMeteredAllowancePriceConversionRateConfigUnion interface {
	ImplementsPlanNewPlanVersionParamsAddPricesPriceNewPlanMeteredAllowancePriceConversionRateConfigUnion()
}

type PlanNewPlanVersionParamsAddPricesPriceNewPlanMeteredAllowancePriceConversionRateConfigConversionRateType string

const (
	PlanNewPlanVersionParamsAddPricesPriceNewPlanMeteredAllowancePriceConversionRateConfigConversionRateTypeUnit   PlanNewPlanVersionParamsAddPricesPriceNewPlanMeteredAllowancePriceConversionRateConfigConversionRateType = "unit"
	PlanNewPlanVersionParamsAddPricesPriceNewPlanMeteredAllowancePriceConversionRateConfigConversionRateTypeTiered PlanNewPlanVersionParamsAddPricesPriceNewPlanMeteredAllowancePriceConversionRateConfigConversionRateType = "tiered"
)

func (r PlanNewPlanVersionParamsAddPricesPriceNewPlanMeteredAllowancePriceConversionRateConfigConversionRateType) IsKnown() bool {
	switch r {
	case PlanNewPlanVersionParamsAddPricesPriceNewPlanMeteredAllowancePriceConversionRateConfigConversionRateTypeUnit, PlanNewPlanVersionParamsAddPricesPriceNewPlanMeteredAllowancePriceConversionRateConfigConversionRateTypeTiered:
		return true
	}
	return false
}

type PlanNewPlanVersionParamsAddPricesPriceNewPlanPercentCompositePrice struct {
	// The cadence to bill for this price on.
	Cadence param.Field[PlanNewPlanVersionParamsAddPricesPriceNewPlanPercentCompositePriceCadence] `json:"cadence" api:"required"`
	// The id of the item the price will be associated with.
	ItemID param.Field[string] `json:"item_id" api:"required"`
	// The pricing model type
	ModelType param.Field[PlanNewPlanVersionParamsAddPricesPriceNewPlanPercentCompositePriceModelType] `json:"model_type" api:"required"`
	// The name of the price.
	Name param.Field[string] `json:"name" api:"required"`
	// Configuration for percent pricing
	PercentConfig param.Field[PlanNewPlanVersionParamsAddPricesPriceNewPlanPercentCompositePricePercentConfig] `json:"percent_config" api:"required"`
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
	ConversionRateConfig param.Field[PlanNewPlanVersionParamsAddPricesPriceNewPlanPercentCompositePriceConversionRateConfigUnion] `json:"conversion_rate_config"`
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
	// The ID of the license type to associate with this price. On a usage price this
	// also marks the price as eligible to draw down from that license type's
	// allocation; a usage price created without it is billed normally. Usage prices
	// with a license type must use the `unit` model, and only draw down when their
	// currency matches the allocation's.
	LicenseTypeID param.Field[string] `json:"license_type_id"`
	// User-specified key/value pairs for the resource. Individual keys can be removed
	// by setting the value to `null`, and the entire metadata mapping can be cleared
	// by setting `metadata` to `null`.
	Metadata param.Field[map[string]string] `json:"metadata"`
	// A transient ID that can be used to reference this price when adding adjustments
	// in the same API call.
	ReferenceID param.Field[string] `json:"reference_id"`
}

func (r PlanNewPlanVersionParamsAddPricesPriceNewPlanPercentCompositePrice) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r PlanNewPlanVersionParamsAddPricesPriceNewPlanPercentCompositePrice) ImplementsPlanNewPlanVersionParamsAddPricesPriceUnion() {
}

// The cadence to bill for this price on.
type PlanNewPlanVersionParamsAddPricesPriceNewPlanPercentCompositePriceCadence string

const (
	PlanNewPlanVersionParamsAddPricesPriceNewPlanPercentCompositePriceCadenceOneTime    PlanNewPlanVersionParamsAddPricesPriceNewPlanPercentCompositePriceCadence = "one_time"
	PlanNewPlanVersionParamsAddPricesPriceNewPlanPercentCompositePriceCadenceMonthly    PlanNewPlanVersionParamsAddPricesPriceNewPlanPercentCompositePriceCadence = "monthly"
	PlanNewPlanVersionParamsAddPricesPriceNewPlanPercentCompositePriceCadenceQuarterly  PlanNewPlanVersionParamsAddPricesPriceNewPlanPercentCompositePriceCadence = "quarterly"
	PlanNewPlanVersionParamsAddPricesPriceNewPlanPercentCompositePriceCadenceSemiAnnual PlanNewPlanVersionParamsAddPricesPriceNewPlanPercentCompositePriceCadence = "semi_annual"
	PlanNewPlanVersionParamsAddPricesPriceNewPlanPercentCompositePriceCadenceAnnual     PlanNewPlanVersionParamsAddPricesPriceNewPlanPercentCompositePriceCadence = "annual"
	PlanNewPlanVersionParamsAddPricesPriceNewPlanPercentCompositePriceCadenceCustom     PlanNewPlanVersionParamsAddPricesPriceNewPlanPercentCompositePriceCadence = "custom"
)

func (r PlanNewPlanVersionParamsAddPricesPriceNewPlanPercentCompositePriceCadence) IsKnown() bool {
	switch r {
	case PlanNewPlanVersionParamsAddPricesPriceNewPlanPercentCompositePriceCadenceOneTime, PlanNewPlanVersionParamsAddPricesPriceNewPlanPercentCompositePriceCadenceMonthly, PlanNewPlanVersionParamsAddPricesPriceNewPlanPercentCompositePriceCadenceQuarterly, PlanNewPlanVersionParamsAddPricesPriceNewPlanPercentCompositePriceCadenceSemiAnnual, PlanNewPlanVersionParamsAddPricesPriceNewPlanPercentCompositePriceCadenceAnnual, PlanNewPlanVersionParamsAddPricesPriceNewPlanPercentCompositePriceCadenceCustom:
		return true
	}
	return false
}

// The pricing model type
type PlanNewPlanVersionParamsAddPricesPriceNewPlanPercentCompositePriceModelType string

const (
	PlanNewPlanVersionParamsAddPricesPriceNewPlanPercentCompositePriceModelTypePercent PlanNewPlanVersionParamsAddPricesPriceNewPlanPercentCompositePriceModelType = "percent"
)

func (r PlanNewPlanVersionParamsAddPricesPriceNewPlanPercentCompositePriceModelType) IsKnown() bool {
	switch r {
	case PlanNewPlanVersionParamsAddPricesPriceNewPlanPercentCompositePriceModelTypePercent:
		return true
	}
	return false
}

// Configuration for percent pricing
type PlanNewPlanVersionParamsAddPricesPriceNewPlanPercentCompositePricePercentConfig struct {
	// Fraction of the component subtotals to charge (0 < percent <= 1).
	Percent param.Field[float64] `json:"percent" api:"required"`
	// Maximum amount to charge. If unset, the fee has no upper bound.
	MaximumAmount param.Field[string] `json:"maximum_amount"`
	// Minimum amount to charge. If unset, the fee is bounded below by 0.
	MinimumAmount param.Field[string] `json:"minimum_amount"`
	// If true, the minimum_amount is prorated based on the service period. The
	// maximum_amount is an absolute cap (never prorated), and the percent applied to
	// upstream subtotals is never prorated either.
	Prorated param.Field[bool] `json:"prorated"`
}

func (r PlanNewPlanVersionParamsAddPricesPriceNewPlanPercentCompositePricePercentConfig) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type PlanNewPlanVersionParamsAddPricesPriceNewPlanPercentCompositePriceConversionRateConfig struct {
	ConversionRateType param.Field[PlanNewPlanVersionParamsAddPricesPriceNewPlanPercentCompositePriceConversionRateConfigConversionRateType] `json:"conversion_rate_type" api:"required"`
	TieredConfig       param.Field[shared.ConversionRateTieredConfigParam]                                                                   `json:"tiered_config"`
	UnitConfig         param.Field[shared.ConversionRateUnitConfigParam]                                                                     `json:"unit_config"`
}

func (r PlanNewPlanVersionParamsAddPricesPriceNewPlanPercentCompositePriceConversionRateConfig) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r PlanNewPlanVersionParamsAddPricesPriceNewPlanPercentCompositePriceConversionRateConfig) ImplementsPlanNewPlanVersionParamsAddPricesPriceNewPlanPercentCompositePriceConversionRateConfigUnion() {
}

// Satisfied by [shared.UnitConversionRateConfigParam],
// [shared.TieredConversionRateConfigParam],
// [PlanNewPlanVersionParamsAddPricesPriceNewPlanPercentCompositePriceConversionRateConfig].
type PlanNewPlanVersionParamsAddPricesPriceNewPlanPercentCompositePriceConversionRateConfigUnion interface {
	ImplementsPlanNewPlanVersionParamsAddPricesPriceNewPlanPercentCompositePriceConversionRateConfigUnion()
}

type PlanNewPlanVersionParamsAddPricesPriceNewPlanPercentCompositePriceConversionRateConfigConversionRateType string

const (
	PlanNewPlanVersionParamsAddPricesPriceNewPlanPercentCompositePriceConversionRateConfigConversionRateTypeUnit   PlanNewPlanVersionParamsAddPricesPriceNewPlanPercentCompositePriceConversionRateConfigConversionRateType = "unit"
	PlanNewPlanVersionParamsAddPricesPriceNewPlanPercentCompositePriceConversionRateConfigConversionRateTypeTiered PlanNewPlanVersionParamsAddPricesPriceNewPlanPercentCompositePriceConversionRateConfigConversionRateType = "tiered"
)

func (r PlanNewPlanVersionParamsAddPricesPriceNewPlanPercentCompositePriceConversionRateConfigConversionRateType) IsKnown() bool {
	switch r {
	case PlanNewPlanVersionParamsAddPricesPriceNewPlanPercentCompositePriceConversionRateConfigConversionRateTypeUnit, PlanNewPlanVersionParamsAddPricesPriceNewPlanPercentCompositePriceConversionRateConfigConversionRateTypeTiered:
		return true
	}
	return false
}

type PlanNewPlanVersionParamsAddPricesPriceNewPlanEventOutputPrice struct {
	// The cadence to bill for this price on.
	Cadence param.Field[PlanNewPlanVersionParamsAddPricesPriceNewPlanEventOutputPriceCadence] `json:"cadence" api:"required"`
	// Configuration for event_output pricing
	EventOutputConfig param.Field[PlanNewPlanVersionParamsAddPricesPriceNewPlanEventOutputPriceEventOutputConfig] `json:"event_output_config" api:"required"`
	// The id of the item the price will be associated with.
	ItemID param.Field[string] `json:"item_id" api:"required"`
	// The pricing model type
	ModelType param.Field[PlanNewPlanVersionParamsAddPricesPriceNewPlanEventOutputPriceModelType] `json:"model_type" api:"required"`
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
	ConversionRateConfig param.Field[PlanNewPlanVersionParamsAddPricesPriceNewPlanEventOutputPriceConversionRateConfigUnion] `json:"conversion_rate_config"`
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
	// The ID of the license type to associate with this price. On a usage price this
	// also marks the price as eligible to draw down from that license type's
	// allocation; a usage price created without it is billed normally. Usage prices
	// with a license type must use the `unit` model, and only draw down when their
	// currency matches the allocation's.
	LicenseTypeID param.Field[string] `json:"license_type_id"`
	// User-specified key/value pairs for the resource. Individual keys can be removed
	// by setting the value to `null`, and the entire metadata mapping can be cleared
	// by setting `metadata` to `null`.
	Metadata param.Field[map[string]string] `json:"metadata"`
	// A transient ID that can be used to reference this price when adding adjustments
	// in the same API call.
	ReferenceID param.Field[string] `json:"reference_id"`
}

func (r PlanNewPlanVersionParamsAddPricesPriceNewPlanEventOutputPrice) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r PlanNewPlanVersionParamsAddPricesPriceNewPlanEventOutputPrice) ImplementsPlanNewPlanVersionParamsAddPricesPriceUnion() {
}

// The cadence to bill for this price on.
type PlanNewPlanVersionParamsAddPricesPriceNewPlanEventOutputPriceCadence string

const (
	PlanNewPlanVersionParamsAddPricesPriceNewPlanEventOutputPriceCadenceOneTime    PlanNewPlanVersionParamsAddPricesPriceNewPlanEventOutputPriceCadence = "one_time"
	PlanNewPlanVersionParamsAddPricesPriceNewPlanEventOutputPriceCadenceMonthly    PlanNewPlanVersionParamsAddPricesPriceNewPlanEventOutputPriceCadence = "monthly"
	PlanNewPlanVersionParamsAddPricesPriceNewPlanEventOutputPriceCadenceQuarterly  PlanNewPlanVersionParamsAddPricesPriceNewPlanEventOutputPriceCadence = "quarterly"
	PlanNewPlanVersionParamsAddPricesPriceNewPlanEventOutputPriceCadenceSemiAnnual PlanNewPlanVersionParamsAddPricesPriceNewPlanEventOutputPriceCadence = "semi_annual"
	PlanNewPlanVersionParamsAddPricesPriceNewPlanEventOutputPriceCadenceAnnual     PlanNewPlanVersionParamsAddPricesPriceNewPlanEventOutputPriceCadence = "annual"
	PlanNewPlanVersionParamsAddPricesPriceNewPlanEventOutputPriceCadenceCustom     PlanNewPlanVersionParamsAddPricesPriceNewPlanEventOutputPriceCadence = "custom"
)

func (r PlanNewPlanVersionParamsAddPricesPriceNewPlanEventOutputPriceCadence) IsKnown() bool {
	switch r {
	case PlanNewPlanVersionParamsAddPricesPriceNewPlanEventOutputPriceCadenceOneTime, PlanNewPlanVersionParamsAddPricesPriceNewPlanEventOutputPriceCadenceMonthly, PlanNewPlanVersionParamsAddPricesPriceNewPlanEventOutputPriceCadenceQuarterly, PlanNewPlanVersionParamsAddPricesPriceNewPlanEventOutputPriceCadenceSemiAnnual, PlanNewPlanVersionParamsAddPricesPriceNewPlanEventOutputPriceCadenceAnnual, PlanNewPlanVersionParamsAddPricesPriceNewPlanEventOutputPriceCadenceCustom:
		return true
	}
	return false
}

// Configuration for event_output pricing
type PlanNewPlanVersionParamsAddPricesPriceNewPlanEventOutputPriceEventOutputConfig struct {
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

func (r PlanNewPlanVersionParamsAddPricesPriceNewPlanEventOutputPriceEventOutputConfig) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// The pricing model type
type PlanNewPlanVersionParamsAddPricesPriceNewPlanEventOutputPriceModelType string

const (
	PlanNewPlanVersionParamsAddPricesPriceNewPlanEventOutputPriceModelTypeEventOutput PlanNewPlanVersionParamsAddPricesPriceNewPlanEventOutputPriceModelType = "event_output"
)

func (r PlanNewPlanVersionParamsAddPricesPriceNewPlanEventOutputPriceModelType) IsKnown() bool {
	switch r {
	case PlanNewPlanVersionParamsAddPricesPriceNewPlanEventOutputPriceModelTypeEventOutput:
		return true
	}
	return false
}

type PlanNewPlanVersionParamsAddPricesPriceNewPlanEventOutputPriceConversionRateConfig struct {
	ConversionRateType param.Field[PlanNewPlanVersionParamsAddPricesPriceNewPlanEventOutputPriceConversionRateConfigConversionRateType] `json:"conversion_rate_type" api:"required"`
	TieredConfig       param.Field[shared.ConversionRateTieredConfigParam]                                                              `json:"tiered_config"`
	UnitConfig         param.Field[shared.ConversionRateUnitConfigParam]                                                                `json:"unit_config"`
}

func (r PlanNewPlanVersionParamsAddPricesPriceNewPlanEventOutputPriceConversionRateConfig) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r PlanNewPlanVersionParamsAddPricesPriceNewPlanEventOutputPriceConversionRateConfig) ImplementsPlanNewPlanVersionParamsAddPricesPriceNewPlanEventOutputPriceConversionRateConfigUnion() {
}

// Satisfied by [shared.UnitConversionRateConfigParam],
// [shared.TieredConversionRateConfigParam],
// [PlanNewPlanVersionParamsAddPricesPriceNewPlanEventOutputPriceConversionRateConfig].
type PlanNewPlanVersionParamsAddPricesPriceNewPlanEventOutputPriceConversionRateConfigUnion interface {
	ImplementsPlanNewPlanVersionParamsAddPricesPriceNewPlanEventOutputPriceConversionRateConfigUnion()
}

type PlanNewPlanVersionParamsAddPricesPriceNewPlanEventOutputPriceConversionRateConfigConversionRateType string

const (
	PlanNewPlanVersionParamsAddPricesPriceNewPlanEventOutputPriceConversionRateConfigConversionRateTypeUnit   PlanNewPlanVersionParamsAddPricesPriceNewPlanEventOutputPriceConversionRateConfigConversionRateType = "unit"
	PlanNewPlanVersionParamsAddPricesPriceNewPlanEventOutputPriceConversionRateConfigConversionRateTypeTiered PlanNewPlanVersionParamsAddPricesPriceNewPlanEventOutputPriceConversionRateConfigConversionRateType = "tiered"
)

func (r PlanNewPlanVersionParamsAddPricesPriceNewPlanEventOutputPriceConversionRateConfigConversionRateType) IsKnown() bool {
	switch r {
	case PlanNewPlanVersionParamsAddPricesPriceNewPlanEventOutputPriceConversionRateConfigConversionRateTypeUnit, PlanNewPlanVersionParamsAddPricesPriceNewPlanEventOutputPriceConversionRateConfigConversionRateTypeTiered:
		return true
	}
	return false
}

type PlanNewPlanVersionParamsAddPricesPriceNewPlanUnitWithMultiplierPrice struct {
	// The cadence to bill for this price on.
	Cadence param.Field[PlanNewPlanVersionParamsAddPricesPriceNewPlanUnitWithMultiplierPriceCadence] `json:"cadence" api:"required"`
	// The id of the item the price will be associated with.
	ItemID param.Field[string] `json:"item_id" api:"required"`
	// The pricing model type
	ModelType param.Field[PlanNewPlanVersionParamsAddPricesPriceNewPlanUnitWithMultiplierPriceModelType] `json:"model_type" api:"required"`
	// The name of the price.
	Name param.Field[string] `json:"name" api:"required"`
	// Configuration for unit_with_multiplier pricing
	UnitWithMultiplierConfig param.Field[PlanNewPlanVersionParamsAddPricesPriceNewPlanUnitWithMultiplierPriceUnitWithMultiplierConfig] `json:"unit_with_multiplier_config" api:"required"`
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
	ConversionRateConfig param.Field[PlanNewPlanVersionParamsAddPricesPriceNewPlanUnitWithMultiplierPriceConversionRateConfigUnion] `json:"conversion_rate_config"`
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
	// The ID of the license type to associate with this price. On a usage price this
	// also marks the price as eligible to draw down from that license type's
	// allocation; a usage price created without it is billed normally. Usage prices
	// with a license type must use the `unit` model, and only draw down when their
	// currency matches the allocation's.
	LicenseTypeID param.Field[string] `json:"license_type_id"`
	// User-specified key/value pairs for the resource. Individual keys can be removed
	// by setting the value to `null`, and the entire metadata mapping can be cleared
	// by setting `metadata` to `null`.
	Metadata param.Field[map[string]string] `json:"metadata"`
	// A transient ID that can be used to reference this price when adding adjustments
	// in the same API call.
	ReferenceID param.Field[string] `json:"reference_id"`
}

func (r PlanNewPlanVersionParamsAddPricesPriceNewPlanUnitWithMultiplierPrice) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r PlanNewPlanVersionParamsAddPricesPriceNewPlanUnitWithMultiplierPrice) ImplementsPlanNewPlanVersionParamsAddPricesPriceUnion() {
}

// The cadence to bill for this price on.
type PlanNewPlanVersionParamsAddPricesPriceNewPlanUnitWithMultiplierPriceCadence string

const (
	PlanNewPlanVersionParamsAddPricesPriceNewPlanUnitWithMultiplierPriceCadenceOneTime    PlanNewPlanVersionParamsAddPricesPriceNewPlanUnitWithMultiplierPriceCadence = "one_time"
	PlanNewPlanVersionParamsAddPricesPriceNewPlanUnitWithMultiplierPriceCadenceMonthly    PlanNewPlanVersionParamsAddPricesPriceNewPlanUnitWithMultiplierPriceCadence = "monthly"
	PlanNewPlanVersionParamsAddPricesPriceNewPlanUnitWithMultiplierPriceCadenceQuarterly  PlanNewPlanVersionParamsAddPricesPriceNewPlanUnitWithMultiplierPriceCadence = "quarterly"
	PlanNewPlanVersionParamsAddPricesPriceNewPlanUnitWithMultiplierPriceCadenceSemiAnnual PlanNewPlanVersionParamsAddPricesPriceNewPlanUnitWithMultiplierPriceCadence = "semi_annual"
	PlanNewPlanVersionParamsAddPricesPriceNewPlanUnitWithMultiplierPriceCadenceAnnual     PlanNewPlanVersionParamsAddPricesPriceNewPlanUnitWithMultiplierPriceCadence = "annual"
	PlanNewPlanVersionParamsAddPricesPriceNewPlanUnitWithMultiplierPriceCadenceCustom     PlanNewPlanVersionParamsAddPricesPriceNewPlanUnitWithMultiplierPriceCadence = "custom"
)

func (r PlanNewPlanVersionParamsAddPricesPriceNewPlanUnitWithMultiplierPriceCadence) IsKnown() bool {
	switch r {
	case PlanNewPlanVersionParamsAddPricesPriceNewPlanUnitWithMultiplierPriceCadenceOneTime, PlanNewPlanVersionParamsAddPricesPriceNewPlanUnitWithMultiplierPriceCadenceMonthly, PlanNewPlanVersionParamsAddPricesPriceNewPlanUnitWithMultiplierPriceCadenceQuarterly, PlanNewPlanVersionParamsAddPricesPriceNewPlanUnitWithMultiplierPriceCadenceSemiAnnual, PlanNewPlanVersionParamsAddPricesPriceNewPlanUnitWithMultiplierPriceCadenceAnnual, PlanNewPlanVersionParamsAddPricesPriceNewPlanUnitWithMultiplierPriceCadenceCustom:
		return true
	}
	return false
}

// The pricing model type
type PlanNewPlanVersionParamsAddPricesPriceNewPlanUnitWithMultiplierPriceModelType string

const (
	PlanNewPlanVersionParamsAddPricesPriceNewPlanUnitWithMultiplierPriceModelTypeUnitWithMultiplier PlanNewPlanVersionParamsAddPricesPriceNewPlanUnitWithMultiplierPriceModelType = "unit_with_multiplier"
)

func (r PlanNewPlanVersionParamsAddPricesPriceNewPlanUnitWithMultiplierPriceModelType) IsKnown() bool {
	switch r {
	case PlanNewPlanVersionParamsAddPricesPriceNewPlanUnitWithMultiplierPriceModelTypeUnitWithMultiplier:
		return true
	}
	return false
}

// Configuration for unit_with_multiplier pricing
type PlanNewPlanVersionParamsAddPricesPriceNewPlanUnitWithMultiplierPriceUnitWithMultiplierConfig struct {
	// Sub-line label for list-price events that don't set display_name_key. Unlabeled
	// events with another multiplier are labeled by their multiplier (e.g. '0.8×
	// rate').
	DefaultDisplayName param.Field[string] `json:"default_display_name" api:"required"`
	// Event property holding the sub-line label for the event's rate (e.g.
	// 'Off-peak').
	DisplayNameKey param.Field[string] `json:"display_name_key" api:"required"`
	// Event property holding the multiplier applied to unit_amount (e.g. 0.8 for a 20%
	// discount). Events without a number between 0 and 1000 with at most 12 decimal
	// places are charged at unit_amount.
	MultiplierKey param.Field[string] `json:"multiplier_key" api:"required"`
	// List price per unit, before the event's multiplier is applied.
	UnitAmount param.Field[string] `json:"unit_amount" api:"required"`
}

func (r PlanNewPlanVersionParamsAddPricesPriceNewPlanUnitWithMultiplierPriceUnitWithMultiplierConfig) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type PlanNewPlanVersionParamsAddPricesPriceNewPlanUnitWithMultiplierPriceConversionRateConfig struct {
	ConversionRateType param.Field[PlanNewPlanVersionParamsAddPricesPriceNewPlanUnitWithMultiplierPriceConversionRateConfigConversionRateType] `json:"conversion_rate_type" api:"required"`
	TieredConfig       param.Field[shared.ConversionRateTieredConfigParam]                                                                     `json:"tiered_config"`
	UnitConfig         param.Field[shared.ConversionRateUnitConfigParam]                                                                       `json:"unit_config"`
}

func (r PlanNewPlanVersionParamsAddPricesPriceNewPlanUnitWithMultiplierPriceConversionRateConfig) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r PlanNewPlanVersionParamsAddPricesPriceNewPlanUnitWithMultiplierPriceConversionRateConfig) ImplementsPlanNewPlanVersionParamsAddPricesPriceNewPlanUnitWithMultiplierPriceConversionRateConfigUnion() {
}

// Satisfied by [shared.UnitConversionRateConfigParam],
// [shared.TieredConversionRateConfigParam],
// [PlanNewPlanVersionParamsAddPricesPriceNewPlanUnitWithMultiplierPriceConversionRateConfig].
type PlanNewPlanVersionParamsAddPricesPriceNewPlanUnitWithMultiplierPriceConversionRateConfigUnion interface {
	ImplementsPlanNewPlanVersionParamsAddPricesPriceNewPlanUnitWithMultiplierPriceConversionRateConfigUnion()
}

type PlanNewPlanVersionParamsAddPricesPriceNewPlanUnitWithMultiplierPriceConversionRateConfigConversionRateType string

const (
	PlanNewPlanVersionParamsAddPricesPriceNewPlanUnitWithMultiplierPriceConversionRateConfigConversionRateTypeUnit   PlanNewPlanVersionParamsAddPricesPriceNewPlanUnitWithMultiplierPriceConversionRateConfigConversionRateType = "unit"
	PlanNewPlanVersionParamsAddPricesPriceNewPlanUnitWithMultiplierPriceConversionRateConfigConversionRateTypeTiered PlanNewPlanVersionParamsAddPricesPriceNewPlanUnitWithMultiplierPriceConversionRateConfigConversionRateType = "tiered"
)

func (r PlanNewPlanVersionParamsAddPricesPriceNewPlanUnitWithMultiplierPriceConversionRateConfigConversionRateType) IsKnown() bool {
	switch r {
	case PlanNewPlanVersionParamsAddPricesPriceNewPlanUnitWithMultiplierPriceConversionRateConfigConversionRateTypeUnit, PlanNewPlanVersionParamsAddPricesPriceNewPlanUnitWithMultiplierPriceConversionRateConfigConversionRateTypeTiered:
		return true
	}
	return false
}

// The cadence to bill for this price on.
type PlanNewPlanVersionParamsAddPricesPriceCadence string

const (
	PlanNewPlanVersionParamsAddPricesPriceCadenceOneTime    PlanNewPlanVersionParamsAddPricesPriceCadence = "one_time"
	PlanNewPlanVersionParamsAddPricesPriceCadenceMonthly    PlanNewPlanVersionParamsAddPricesPriceCadence = "monthly"
	PlanNewPlanVersionParamsAddPricesPriceCadenceQuarterly  PlanNewPlanVersionParamsAddPricesPriceCadence = "quarterly"
	PlanNewPlanVersionParamsAddPricesPriceCadenceSemiAnnual PlanNewPlanVersionParamsAddPricesPriceCadence = "semi_annual"
	PlanNewPlanVersionParamsAddPricesPriceCadenceAnnual     PlanNewPlanVersionParamsAddPricesPriceCadence = "annual"
	PlanNewPlanVersionParamsAddPricesPriceCadenceCustom     PlanNewPlanVersionParamsAddPricesPriceCadence = "custom"
)

func (r PlanNewPlanVersionParamsAddPricesPriceCadence) IsKnown() bool {
	switch r {
	case PlanNewPlanVersionParamsAddPricesPriceCadenceOneTime, PlanNewPlanVersionParamsAddPricesPriceCadenceMonthly, PlanNewPlanVersionParamsAddPricesPriceCadenceQuarterly, PlanNewPlanVersionParamsAddPricesPriceCadenceSemiAnnual, PlanNewPlanVersionParamsAddPricesPriceCadenceAnnual, PlanNewPlanVersionParamsAddPricesPriceCadenceCustom:
		return true
	}
	return false
}

// The pricing model type
type PlanNewPlanVersionParamsAddPricesPriceModelType string

const (
	PlanNewPlanVersionParamsAddPricesPriceModelTypeUnit                            PlanNewPlanVersionParamsAddPricesPriceModelType = "unit"
	PlanNewPlanVersionParamsAddPricesPriceModelTypeTiered                          PlanNewPlanVersionParamsAddPricesPriceModelType = "tiered"
	PlanNewPlanVersionParamsAddPricesPriceModelTypeBulk                            PlanNewPlanVersionParamsAddPricesPriceModelType = "bulk"
	PlanNewPlanVersionParamsAddPricesPriceModelTypeBulkWithFilters                 PlanNewPlanVersionParamsAddPricesPriceModelType = "bulk_with_filters"
	PlanNewPlanVersionParamsAddPricesPriceModelTypePackage                         PlanNewPlanVersionParamsAddPricesPriceModelType = "package"
	PlanNewPlanVersionParamsAddPricesPriceModelTypeMatrix                          PlanNewPlanVersionParamsAddPricesPriceModelType = "matrix"
	PlanNewPlanVersionParamsAddPricesPriceModelTypeThresholdTotalAmount            PlanNewPlanVersionParamsAddPricesPriceModelType = "threshold_total_amount"
	PlanNewPlanVersionParamsAddPricesPriceModelTypeTieredPackage                   PlanNewPlanVersionParamsAddPricesPriceModelType = "tiered_package"
	PlanNewPlanVersionParamsAddPricesPriceModelTypeTieredWithMinimum               PlanNewPlanVersionParamsAddPricesPriceModelType = "tiered_with_minimum"
	PlanNewPlanVersionParamsAddPricesPriceModelTypeGroupedTiered                   PlanNewPlanVersionParamsAddPricesPriceModelType = "grouped_tiered"
	PlanNewPlanVersionParamsAddPricesPriceModelTypeGroupedTieredMatrix             PlanNewPlanVersionParamsAddPricesPriceModelType = "grouped_tiered_matrix"
	PlanNewPlanVersionParamsAddPricesPriceModelTypeTieredPackageWithMinimum        PlanNewPlanVersionParamsAddPricesPriceModelType = "tiered_package_with_minimum"
	PlanNewPlanVersionParamsAddPricesPriceModelTypePackageWithAllocation           PlanNewPlanVersionParamsAddPricesPriceModelType = "package_with_allocation"
	PlanNewPlanVersionParamsAddPricesPriceModelTypeUnitWithPercent                 PlanNewPlanVersionParamsAddPricesPriceModelType = "unit_with_percent"
	PlanNewPlanVersionParamsAddPricesPriceModelTypeMatrixWithAllocation            PlanNewPlanVersionParamsAddPricesPriceModelType = "matrix_with_allocation"
	PlanNewPlanVersionParamsAddPricesPriceModelTypeTieredMatrixWithAllocation      PlanNewPlanVersionParamsAddPricesPriceModelType = "tiered_matrix_with_allocation"
	PlanNewPlanVersionParamsAddPricesPriceModelTypeMatrixWithThresholdDiscounts    PlanNewPlanVersionParamsAddPricesPriceModelType = "matrix_with_threshold_discounts"
	PlanNewPlanVersionParamsAddPricesPriceModelTypeTieredWithProration             PlanNewPlanVersionParamsAddPricesPriceModelType = "tiered_with_proration"
	PlanNewPlanVersionParamsAddPricesPriceModelTypeUnitWithProration               PlanNewPlanVersionParamsAddPricesPriceModelType = "unit_with_proration"
	PlanNewPlanVersionParamsAddPricesPriceModelTypeGroupedAllocation               PlanNewPlanVersionParamsAddPricesPriceModelType = "grouped_allocation"
	PlanNewPlanVersionParamsAddPricesPriceModelTypeBulkWithProration               PlanNewPlanVersionParamsAddPricesPriceModelType = "bulk_with_proration"
	PlanNewPlanVersionParamsAddPricesPriceModelTypeGroupedWithProratedMinimum      PlanNewPlanVersionParamsAddPricesPriceModelType = "grouped_with_prorated_minimum"
	PlanNewPlanVersionParamsAddPricesPriceModelTypeGroupedWithMeteredMinimum       PlanNewPlanVersionParamsAddPricesPriceModelType = "grouped_with_metered_minimum"
	PlanNewPlanVersionParamsAddPricesPriceModelTypeGroupedWithMinMaxThresholds     PlanNewPlanVersionParamsAddPricesPriceModelType = "grouped_with_min_max_thresholds"
	PlanNewPlanVersionParamsAddPricesPriceModelTypeMatrixWithDisplayName           PlanNewPlanVersionParamsAddPricesPriceModelType = "matrix_with_display_name"
	PlanNewPlanVersionParamsAddPricesPriceModelTypeGroupedTieredPackage            PlanNewPlanVersionParamsAddPricesPriceModelType = "grouped_tiered_package"
	PlanNewPlanVersionParamsAddPricesPriceModelTypeMaxGroupTieredPackage           PlanNewPlanVersionParamsAddPricesPriceModelType = "max_group_tiered_package"
	PlanNewPlanVersionParamsAddPricesPriceModelTypeScalableMatrixWithUnitPricing   PlanNewPlanVersionParamsAddPricesPriceModelType = "scalable_matrix_with_unit_pricing"
	PlanNewPlanVersionParamsAddPricesPriceModelTypeScalableMatrixWithTieredPricing PlanNewPlanVersionParamsAddPricesPriceModelType = "scalable_matrix_with_tiered_pricing"
	PlanNewPlanVersionParamsAddPricesPriceModelTypeCumulativeGroupedBulk           PlanNewPlanVersionParamsAddPricesPriceModelType = "cumulative_grouped_bulk"
	PlanNewPlanVersionParamsAddPricesPriceModelTypeCumulativeGroupedAllocation     PlanNewPlanVersionParamsAddPricesPriceModelType = "cumulative_grouped_allocation"
	PlanNewPlanVersionParamsAddPricesPriceModelTypeDailyCreditAllowance            PlanNewPlanVersionParamsAddPricesPriceModelType = "daily_credit_allowance"
	PlanNewPlanVersionParamsAddPricesPriceModelTypeMeteredAllowance                PlanNewPlanVersionParamsAddPricesPriceModelType = "metered_allowance"
	PlanNewPlanVersionParamsAddPricesPriceModelTypeMinimumComposite                PlanNewPlanVersionParamsAddPricesPriceModelType = "minimum_composite"
	PlanNewPlanVersionParamsAddPricesPriceModelTypePercent                         PlanNewPlanVersionParamsAddPricesPriceModelType = "percent"
	PlanNewPlanVersionParamsAddPricesPriceModelTypeEventOutput                     PlanNewPlanVersionParamsAddPricesPriceModelType = "event_output"
	PlanNewPlanVersionParamsAddPricesPriceModelTypeUnitWithMultiplier              PlanNewPlanVersionParamsAddPricesPriceModelType = "unit_with_multiplier"
)

func (r PlanNewPlanVersionParamsAddPricesPriceModelType) IsKnown() bool {
	switch r {
	case PlanNewPlanVersionParamsAddPricesPriceModelTypeUnit, PlanNewPlanVersionParamsAddPricesPriceModelTypeTiered, PlanNewPlanVersionParamsAddPricesPriceModelTypeBulk, PlanNewPlanVersionParamsAddPricesPriceModelTypeBulkWithFilters, PlanNewPlanVersionParamsAddPricesPriceModelTypePackage, PlanNewPlanVersionParamsAddPricesPriceModelTypeMatrix, PlanNewPlanVersionParamsAddPricesPriceModelTypeThresholdTotalAmount, PlanNewPlanVersionParamsAddPricesPriceModelTypeTieredPackage, PlanNewPlanVersionParamsAddPricesPriceModelTypeTieredWithMinimum, PlanNewPlanVersionParamsAddPricesPriceModelTypeGroupedTiered, PlanNewPlanVersionParamsAddPricesPriceModelTypeGroupedTieredMatrix, PlanNewPlanVersionParamsAddPricesPriceModelTypeTieredPackageWithMinimum, PlanNewPlanVersionParamsAddPricesPriceModelTypePackageWithAllocation, PlanNewPlanVersionParamsAddPricesPriceModelTypeUnitWithPercent, PlanNewPlanVersionParamsAddPricesPriceModelTypeMatrixWithAllocation, PlanNewPlanVersionParamsAddPricesPriceModelTypeTieredMatrixWithAllocation, PlanNewPlanVersionParamsAddPricesPriceModelTypeMatrixWithThresholdDiscounts, PlanNewPlanVersionParamsAddPricesPriceModelTypeTieredWithProration, PlanNewPlanVersionParamsAddPricesPriceModelTypeUnitWithProration, PlanNewPlanVersionParamsAddPricesPriceModelTypeGroupedAllocation, PlanNewPlanVersionParamsAddPricesPriceModelTypeBulkWithProration, PlanNewPlanVersionParamsAddPricesPriceModelTypeGroupedWithProratedMinimum, PlanNewPlanVersionParamsAddPricesPriceModelTypeGroupedWithMeteredMinimum, PlanNewPlanVersionParamsAddPricesPriceModelTypeGroupedWithMinMaxThresholds, PlanNewPlanVersionParamsAddPricesPriceModelTypeMatrixWithDisplayName, PlanNewPlanVersionParamsAddPricesPriceModelTypeGroupedTieredPackage, PlanNewPlanVersionParamsAddPricesPriceModelTypeMaxGroupTieredPackage, PlanNewPlanVersionParamsAddPricesPriceModelTypeScalableMatrixWithUnitPricing, PlanNewPlanVersionParamsAddPricesPriceModelTypeScalableMatrixWithTieredPricing, PlanNewPlanVersionParamsAddPricesPriceModelTypeCumulativeGroupedBulk, PlanNewPlanVersionParamsAddPricesPriceModelTypeCumulativeGroupedAllocation, PlanNewPlanVersionParamsAddPricesPriceModelTypeDailyCreditAllowance, PlanNewPlanVersionParamsAddPricesPriceModelTypeMeteredAllowance, PlanNewPlanVersionParamsAddPricesPriceModelTypeMinimumComposite, PlanNewPlanVersionParamsAddPricesPriceModelTypePercent, PlanNewPlanVersionParamsAddPricesPriceModelTypeEventOutput, PlanNewPlanVersionParamsAddPricesPriceModelTypeUnitWithMultiplier:
		return true
	}
	return false
}

type PlanNewPlanVersionParamsRemoveAdjustment struct {
	// The id of the adjustment to remove from on the plan.
	AdjustmentID param.Field[string] `json:"adjustment_id" api:"required"`
	// The phase to remove this adjustment from.
	PlanPhaseOrder param.Field[int64] `json:"plan_phase_order"`
}

func (r PlanNewPlanVersionParamsRemoveAdjustment) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type PlanNewPlanVersionParamsRemovePrice struct {
	// The id of the price to remove from the plan.
	PriceID param.Field[string] `json:"price_id" api:"required"`
	// The phase to remove this price from.
	PlanPhaseOrder param.Field[int64] `json:"plan_phase_order"`
}

func (r PlanNewPlanVersionParamsRemovePrice) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type PlanNewPlanVersionParamsReplaceAdjustment struct {
	// The definition of a new adjustment to create and add to the plan.
	Adjustment param.Field[PlanNewPlanVersionParamsReplaceAdjustmentsAdjustmentUnion] `json:"adjustment" api:"required"`
	// The id of the adjustment on the plan to replace in the plan.
	ReplacesAdjustmentID param.Field[string] `json:"replaces_adjustment_id" api:"required"`
	// The phase to replace this adjustment from.
	PlanPhaseOrder param.Field[int64] `json:"plan_phase_order"`
}

func (r PlanNewPlanVersionParamsReplaceAdjustment) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// The definition of a new adjustment to create and add to the plan.
type PlanNewPlanVersionParamsReplaceAdjustmentsAdjustment struct {
	AdjustmentType param.Field[PlanNewPlanVersionParamsReplaceAdjustmentsAdjustmentAdjustmentType] `json:"adjustment_type" api:"required"`
	AmountDiscount param.Field[string]                                                             `json:"amount_discount"`
	// If set, the adjustment will apply to every price on the subscription.
	AppliesToAll      param.Field[PlanNewPlanVersionParamsReplaceAdjustmentsAdjustmentAppliesToAll] `json:"applies_to_all"`
	AppliesToItemIDs  param.Field[interface{}]                                                      `json:"applies_to_item_ids"`
	AppliesToPriceIDs param.Field[interface{}]                                                      `json:"applies_to_price_ids"`
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
	PriceType     param.Field[PlanNewPlanVersionParamsReplaceAdjustmentsAdjustmentPriceType] `json:"price_type"`
	Tiers         param.Field[interface{}]                                                   `json:"tiers"`
	UsageDiscount param.Field[float64]                                                       `json:"usage_discount"`
}

func (r PlanNewPlanVersionParamsReplaceAdjustmentsAdjustment) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r PlanNewPlanVersionParamsReplaceAdjustmentsAdjustment) ImplementsPlanNewPlanVersionParamsReplaceAdjustmentsAdjustmentUnion() {
}

// The definition of a new adjustment to create and add to the plan.
//
// Satisfied by [shared.NewPercentageDiscountParam],
// [shared.NewUsageDiscountParam], [shared.NewAmountDiscountParam],
// [shared.NewMinimumParam], [shared.NewMaximumParam],
// [PlanNewPlanVersionParamsReplaceAdjustmentsAdjustmentNewTieredPercentageDiscount],
// [PlanNewPlanVersionParamsReplaceAdjustmentsAdjustment].
type PlanNewPlanVersionParamsReplaceAdjustmentsAdjustmentUnion interface {
	ImplementsPlanNewPlanVersionParamsReplaceAdjustmentsAdjustmentUnion()
}

type PlanNewPlanVersionParamsReplaceAdjustmentsAdjustmentNewTieredPercentageDiscount struct {
	AdjustmentType param.Field[PlanNewPlanVersionParamsReplaceAdjustmentsAdjustmentNewTieredPercentageDiscountAdjustmentType] `json:"adjustment_type" api:"required"`
	Tiers          param.Field[[]PlanNewPlanVersionParamsReplaceAdjustmentsAdjustmentNewTieredPercentageDiscountTier]         `json:"tiers" api:"required"`
	// If set, the adjustment will apply to every price on the subscription.
	AppliesToAll param.Field[PlanNewPlanVersionParamsReplaceAdjustmentsAdjustmentNewTieredPercentageDiscountAppliesToAll] `json:"applies_to_all"`
	// The set of item IDs to which this adjustment applies.
	AppliesToItemIDs param.Field[[]string] `json:"applies_to_item_ids"`
	// The set of price IDs to which this adjustment applies.
	AppliesToPriceIDs param.Field[[]string] `json:"applies_to_price_ids"`
	// If set, only prices in the specified currency will have the adjustment applied.
	Currency param.Field[string] `json:"currency"`
	// A list of filters that determine which prices this adjustment will apply to.
	Filters param.Field[[]PlanNewPlanVersionParamsReplaceAdjustmentsAdjustmentNewTieredPercentageDiscountFilter] `json:"filters"`
	// When false, this adjustment will be applied to a single price. Otherwise, it
	// will be applied at the invoice level, possibly to multiple prices.
	IsInvoiceLevel param.Field[bool] `json:"is_invoice_level"`
	// If set, only prices of the specified type will have the adjustment applied.
	PriceType param.Field[PlanNewPlanVersionParamsReplaceAdjustmentsAdjustmentNewTieredPercentageDiscountPriceType] `json:"price_type"`
}

func (r PlanNewPlanVersionParamsReplaceAdjustmentsAdjustmentNewTieredPercentageDiscount) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r PlanNewPlanVersionParamsReplaceAdjustmentsAdjustmentNewTieredPercentageDiscount) ImplementsPlanNewPlanVersionParamsReplaceAdjustmentsAdjustmentUnion() {
}

type PlanNewPlanVersionParamsReplaceAdjustmentsAdjustmentNewTieredPercentageDiscountAdjustmentType string

const (
	PlanNewPlanVersionParamsReplaceAdjustmentsAdjustmentNewTieredPercentageDiscountAdjustmentTypeTieredPercentageDiscount PlanNewPlanVersionParamsReplaceAdjustmentsAdjustmentNewTieredPercentageDiscountAdjustmentType = "tiered_percentage_discount"
)

func (r PlanNewPlanVersionParamsReplaceAdjustmentsAdjustmentNewTieredPercentageDiscountAdjustmentType) IsKnown() bool {
	switch r {
	case PlanNewPlanVersionParamsReplaceAdjustmentsAdjustmentNewTieredPercentageDiscountAdjustmentTypeTieredPercentageDiscount:
		return true
	}
	return false
}

type PlanNewPlanVersionParamsReplaceAdjustmentsAdjustmentNewTieredPercentageDiscountTier struct {
	// Exclusive lower bound of cumulative spend for this tier.
	LowerBound param.Field[float64] `json:"lower_bound" api:"required"`
	// The percentage (0-1) discounted from spend in this tier.
	Percentage param.Field[float64] `json:"percentage" api:"required"`
	// Inclusive upper bound of cumulative spend; null for the final open-ended tier.
	UpperBound param.Field[float64] `json:"upper_bound"`
}

func (r PlanNewPlanVersionParamsReplaceAdjustmentsAdjustmentNewTieredPercentageDiscountTier) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// If set, the adjustment will apply to every price on the subscription.
type PlanNewPlanVersionParamsReplaceAdjustmentsAdjustmentNewTieredPercentageDiscountAppliesToAll bool

const (
	PlanNewPlanVersionParamsReplaceAdjustmentsAdjustmentNewTieredPercentageDiscountAppliesToAllTrue PlanNewPlanVersionParamsReplaceAdjustmentsAdjustmentNewTieredPercentageDiscountAppliesToAll = true
)

func (r PlanNewPlanVersionParamsReplaceAdjustmentsAdjustmentNewTieredPercentageDiscountAppliesToAll) IsKnown() bool {
	switch r {
	case PlanNewPlanVersionParamsReplaceAdjustmentsAdjustmentNewTieredPercentageDiscountAppliesToAllTrue:
		return true
	}
	return false
}

type PlanNewPlanVersionParamsReplaceAdjustmentsAdjustmentNewTieredPercentageDiscountFilter struct {
	// The property of the price to filter on.
	Field param.Field[PlanNewPlanVersionParamsReplaceAdjustmentsAdjustmentNewTieredPercentageDiscountFiltersField] `json:"field" api:"required"`
	// Should prices that match the filter be included or excluded.
	Operator param.Field[PlanNewPlanVersionParamsReplaceAdjustmentsAdjustmentNewTieredPercentageDiscountFiltersOperator] `json:"operator" api:"required"`
	// The IDs or values that match this filter.
	Values param.Field[[]string] `json:"values" api:"required"`
}

func (r PlanNewPlanVersionParamsReplaceAdjustmentsAdjustmentNewTieredPercentageDiscountFilter) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// The property of the price to filter on.
type PlanNewPlanVersionParamsReplaceAdjustmentsAdjustmentNewTieredPercentageDiscountFiltersField string

const (
	PlanNewPlanVersionParamsReplaceAdjustmentsAdjustmentNewTieredPercentageDiscountFiltersFieldPriceID       PlanNewPlanVersionParamsReplaceAdjustmentsAdjustmentNewTieredPercentageDiscountFiltersField = "price_id"
	PlanNewPlanVersionParamsReplaceAdjustmentsAdjustmentNewTieredPercentageDiscountFiltersFieldItemID        PlanNewPlanVersionParamsReplaceAdjustmentsAdjustmentNewTieredPercentageDiscountFiltersField = "item_id"
	PlanNewPlanVersionParamsReplaceAdjustmentsAdjustmentNewTieredPercentageDiscountFiltersFieldPriceType     PlanNewPlanVersionParamsReplaceAdjustmentsAdjustmentNewTieredPercentageDiscountFiltersField = "price_type"
	PlanNewPlanVersionParamsReplaceAdjustmentsAdjustmentNewTieredPercentageDiscountFiltersFieldCurrency      PlanNewPlanVersionParamsReplaceAdjustmentsAdjustmentNewTieredPercentageDiscountFiltersField = "currency"
	PlanNewPlanVersionParamsReplaceAdjustmentsAdjustmentNewTieredPercentageDiscountFiltersFieldPricingUnitID PlanNewPlanVersionParamsReplaceAdjustmentsAdjustmentNewTieredPercentageDiscountFiltersField = "pricing_unit_id"
)

func (r PlanNewPlanVersionParamsReplaceAdjustmentsAdjustmentNewTieredPercentageDiscountFiltersField) IsKnown() bool {
	switch r {
	case PlanNewPlanVersionParamsReplaceAdjustmentsAdjustmentNewTieredPercentageDiscountFiltersFieldPriceID, PlanNewPlanVersionParamsReplaceAdjustmentsAdjustmentNewTieredPercentageDiscountFiltersFieldItemID, PlanNewPlanVersionParamsReplaceAdjustmentsAdjustmentNewTieredPercentageDiscountFiltersFieldPriceType, PlanNewPlanVersionParamsReplaceAdjustmentsAdjustmentNewTieredPercentageDiscountFiltersFieldCurrency, PlanNewPlanVersionParamsReplaceAdjustmentsAdjustmentNewTieredPercentageDiscountFiltersFieldPricingUnitID:
		return true
	}
	return false
}

// Should prices that match the filter be included or excluded.
type PlanNewPlanVersionParamsReplaceAdjustmentsAdjustmentNewTieredPercentageDiscountFiltersOperator string

const (
	PlanNewPlanVersionParamsReplaceAdjustmentsAdjustmentNewTieredPercentageDiscountFiltersOperatorIncludes PlanNewPlanVersionParamsReplaceAdjustmentsAdjustmentNewTieredPercentageDiscountFiltersOperator = "includes"
	PlanNewPlanVersionParamsReplaceAdjustmentsAdjustmentNewTieredPercentageDiscountFiltersOperatorExcludes PlanNewPlanVersionParamsReplaceAdjustmentsAdjustmentNewTieredPercentageDiscountFiltersOperator = "excludes"
)

func (r PlanNewPlanVersionParamsReplaceAdjustmentsAdjustmentNewTieredPercentageDiscountFiltersOperator) IsKnown() bool {
	switch r {
	case PlanNewPlanVersionParamsReplaceAdjustmentsAdjustmentNewTieredPercentageDiscountFiltersOperatorIncludes, PlanNewPlanVersionParamsReplaceAdjustmentsAdjustmentNewTieredPercentageDiscountFiltersOperatorExcludes:
		return true
	}
	return false
}

// If set, only prices of the specified type will have the adjustment applied.
type PlanNewPlanVersionParamsReplaceAdjustmentsAdjustmentNewTieredPercentageDiscountPriceType string

const (
	PlanNewPlanVersionParamsReplaceAdjustmentsAdjustmentNewTieredPercentageDiscountPriceTypeUsage          PlanNewPlanVersionParamsReplaceAdjustmentsAdjustmentNewTieredPercentageDiscountPriceType = "usage"
	PlanNewPlanVersionParamsReplaceAdjustmentsAdjustmentNewTieredPercentageDiscountPriceTypeFixedInAdvance PlanNewPlanVersionParamsReplaceAdjustmentsAdjustmentNewTieredPercentageDiscountPriceType = "fixed_in_advance"
	PlanNewPlanVersionParamsReplaceAdjustmentsAdjustmentNewTieredPercentageDiscountPriceTypeFixedInArrears PlanNewPlanVersionParamsReplaceAdjustmentsAdjustmentNewTieredPercentageDiscountPriceType = "fixed_in_arrears"
	PlanNewPlanVersionParamsReplaceAdjustmentsAdjustmentNewTieredPercentageDiscountPriceTypeFixed          PlanNewPlanVersionParamsReplaceAdjustmentsAdjustmentNewTieredPercentageDiscountPriceType = "fixed"
	PlanNewPlanVersionParamsReplaceAdjustmentsAdjustmentNewTieredPercentageDiscountPriceTypeInArrears      PlanNewPlanVersionParamsReplaceAdjustmentsAdjustmentNewTieredPercentageDiscountPriceType = "in_arrears"
)

func (r PlanNewPlanVersionParamsReplaceAdjustmentsAdjustmentNewTieredPercentageDiscountPriceType) IsKnown() bool {
	switch r {
	case PlanNewPlanVersionParamsReplaceAdjustmentsAdjustmentNewTieredPercentageDiscountPriceTypeUsage, PlanNewPlanVersionParamsReplaceAdjustmentsAdjustmentNewTieredPercentageDiscountPriceTypeFixedInAdvance, PlanNewPlanVersionParamsReplaceAdjustmentsAdjustmentNewTieredPercentageDiscountPriceTypeFixedInArrears, PlanNewPlanVersionParamsReplaceAdjustmentsAdjustmentNewTieredPercentageDiscountPriceTypeFixed, PlanNewPlanVersionParamsReplaceAdjustmentsAdjustmentNewTieredPercentageDiscountPriceTypeInArrears:
		return true
	}
	return false
}

type PlanNewPlanVersionParamsReplaceAdjustmentsAdjustmentAdjustmentType string

const (
	PlanNewPlanVersionParamsReplaceAdjustmentsAdjustmentAdjustmentTypePercentageDiscount       PlanNewPlanVersionParamsReplaceAdjustmentsAdjustmentAdjustmentType = "percentage_discount"
	PlanNewPlanVersionParamsReplaceAdjustmentsAdjustmentAdjustmentTypeUsageDiscount            PlanNewPlanVersionParamsReplaceAdjustmentsAdjustmentAdjustmentType = "usage_discount"
	PlanNewPlanVersionParamsReplaceAdjustmentsAdjustmentAdjustmentTypeAmountDiscount           PlanNewPlanVersionParamsReplaceAdjustmentsAdjustmentAdjustmentType = "amount_discount"
	PlanNewPlanVersionParamsReplaceAdjustmentsAdjustmentAdjustmentTypeMinimum                  PlanNewPlanVersionParamsReplaceAdjustmentsAdjustmentAdjustmentType = "minimum"
	PlanNewPlanVersionParamsReplaceAdjustmentsAdjustmentAdjustmentTypeMaximum                  PlanNewPlanVersionParamsReplaceAdjustmentsAdjustmentAdjustmentType = "maximum"
	PlanNewPlanVersionParamsReplaceAdjustmentsAdjustmentAdjustmentTypeTieredPercentageDiscount PlanNewPlanVersionParamsReplaceAdjustmentsAdjustmentAdjustmentType = "tiered_percentage_discount"
)

func (r PlanNewPlanVersionParamsReplaceAdjustmentsAdjustmentAdjustmentType) IsKnown() bool {
	switch r {
	case PlanNewPlanVersionParamsReplaceAdjustmentsAdjustmentAdjustmentTypePercentageDiscount, PlanNewPlanVersionParamsReplaceAdjustmentsAdjustmentAdjustmentTypeUsageDiscount, PlanNewPlanVersionParamsReplaceAdjustmentsAdjustmentAdjustmentTypeAmountDiscount, PlanNewPlanVersionParamsReplaceAdjustmentsAdjustmentAdjustmentTypeMinimum, PlanNewPlanVersionParamsReplaceAdjustmentsAdjustmentAdjustmentTypeMaximum, PlanNewPlanVersionParamsReplaceAdjustmentsAdjustmentAdjustmentTypeTieredPercentageDiscount:
		return true
	}
	return false
}

// If set, the adjustment will apply to every price on the subscription.
type PlanNewPlanVersionParamsReplaceAdjustmentsAdjustmentAppliesToAll bool

const (
	PlanNewPlanVersionParamsReplaceAdjustmentsAdjustmentAppliesToAllTrue PlanNewPlanVersionParamsReplaceAdjustmentsAdjustmentAppliesToAll = true
)

func (r PlanNewPlanVersionParamsReplaceAdjustmentsAdjustmentAppliesToAll) IsKnown() bool {
	switch r {
	case PlanNewPlanVersionParamsReplaceAdjustmentsAdjustmentAppliesToAllTrue:
		return true
	}
	return false
}

// If set, only prices of the specified type will have the adjustment applied.
type PlanNewPlanVersionParamsReplaceAdjustmentsAdjustmentPriceType string

const (
	PlanNewPlanVersionParamsReplaceAdjustmentsAdjustmentPriceTypeUsage          PlanNewPlanVersionParamsReplaceAdjustmentsAdjustmentPriceType = "usage"
	PlanNewPlanVersionParamsReplaceAdjustmentsAdjustmentPriceTypeFixedInAdvance PlanNewPlanVersionParamsReplaceAdjustmentsAdjustmentPriceType = "fixed_in_advance"
	PlanNewPlanVersionParamsReplaceAdjustmentsAdjustmentPriceTypeFixedInArrears PlanNewPlanVersionParamsReplaceAdjustmentsAdjustmentPriceType = "fixed_in_arrears"
	PlanNewPlanVersionParamsReplaceAdjustmentsAdjustmentPriceTypeFixed          PlanNewPlanVersionParamsReplaceAdjustmentsAdjustmentPriceType = "fixed"
	PlanNewPlanVersionParamsReplaceAdjustmentsAdjustmentPriceTypeInArrears      PlanNewPlanVersionParamsReplaceAdjustmentsAdjustmentPriceType = "in_arrears"
)

func (r PlanNewPlanVersionParamsReplaceAdjustmentsAdjustmentPriceType) IsKnown() bool {
	switch r {
	case PlanNewPlanVersionParamsReplaceAdjustmentsAdjustmentPriceTypeUsage, PlanNewPlanVersionParamsReplaceAdjustmentsAdjustmentPriceTypeFixedInAdvance, PlanNewPlanVersionParamsReplaceAdjustmentsAdjustmentPriceTypeFixedInArrears, PlanNewPlanVersionParamsReplaceAdjustmentsAdjustmentPriceTypeFixed, PlanNewPlanVersionParamsReplaceAdjustmentsAdjustmentPriceTypeInArrears:
		return true
	}
	return false
}

type PlanNewPlanVersionParamsReplacePrice struct {
	// The id of the price on the plan to replace in the plan.
	ReplacesPriceID param.Field[string] `json:"replaces_price_id" api:"required"`
	// The allocation price to add to the plan.
	AllocationPrice param.Field[shared.NewAllocationPriceParam] `json:"allocation_price"`
	// The license allocation price to add to the plan.
	LicenseAllocationPrice param.Field[PlanNewPlanVersionParamsReplacePricesLicenseAllocationPrice] `json:"license_allocation_price"`
	// The phase to replace this price from.
	PlanPhaseOrder param.Field[int64] `json:"plan_phase_order"`
	// New plan price request body params.
	Price param.Field[PlanNewPlanVersionParamsReplacePricesPriceUnion] `json:"price"`
}

func (r PlanNewPlanVersionParamsReplacePrice) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// The license allocation price to add to the plan.
type PlanNewPlanVersionParamsReplacePricesLicenseAllocationPrice struct {
	// The cadence to bill for this price on.
	Cadence param.Field[PlanNewPlanVersionParamsReplacePricesLicenseAllocationPriceCadence] `json:"cadence" api:"required"`
	// The id of the item the price will be associated with.
	ItemID param.Field[string] `json:"item_id" api:"required"`
	// License allocations to associate with this price. Each entry defines a
	// per-license credit pool granted each cadence. Requires license_type_id or
	// license_type_configuration to be set. Usage prices draw down from these
	// allocations only if they set the same license type.
	LicenseAllocations param.Field[[]PlanNewPlanVersionParamsReplacePricesLicenseAllocationPriceLicenseAllocation] `json:"license_allocations" api:"required"`
	// The pricing model type
	ModelType param.Field[PlanNewPlanVersionParamsReplacePricesLicenseAllocationPriceModelType] `json:"model_type" api:"required"`
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
	ConversionRateConfig param.Field[PlanNewPlanVersionParamsReplacePricesLicenseAllocationPriceConversionRateConfigUnion] `json:"conversion_rate_config"`
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
	// The ID of the license type to associate with this price. On a usage price this
	// also marks the price as eligible to draw down from that license type's
	// allocation; a usage price created without it is billed normally. Usage prices
	// with a license type must use the `unit` model, and only draw down when their
	// currency matches the allocation's.
	LicenseTypeID param.Field[string] `json:"license_type_id"`
	// User-specified key/value pairs for the resource. Individual keys can be removed
	// by setting the value to `null`, and the entire metadata mapping can be cleared
	// by setting `metadata` to `null`.
	Metadata param.Field[map[string]string] `json:"metadata"`
	// A transient ID that can be used to reference this price when adding adjustments
	// in the same API call.
	ReferenceID param.Field[string] `json:"reference_id"`
}

func (r PlanNewPlanVersionParamsReplacePricesLicenseAllocationPrice) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// The cadence to bill for this price on.
type PlanNewPlanVersionParamsReplacePricesLicenseAllocationPriceCadence string

const (
	PlanNewPlanVersionParamsReplacePricesLicenseAllocationPriceCadenceOneTime    PlanNewPlanVersionParamsReplacePricesLicenseAllocationPriceCadence = "one_time"
	PlanNewPlanVersionParamsReplacePricesLicenseAllocationPriceCadenceMonthly    PlanNewPlanVersionParamsReplacePricesLicenseAllocationPriceCadence = "monthly"
	PlanNewPlanVersionParamsReplacePricesLicenseAllocationPriceCadenceQuarterly  PlanNewPlanVersionParamsReplacePricesLicenseAllocationPriceCadence = "quarterly"
	PlanNewPlanVersionParamsReplacePricesLicenseAllocationPriceCadenceSemiAnnual PlanNewPlanVersionParamsReplacePricesLicenseAllocationPriceCadence = "semi_annual"
	PlanNewPlanVersionParamsReplacePricesLicenseAllocationPriceCadenceAnnual     PlanNewPlanVersionParamsReplacePricesLicenseAllocationPriceCadence = "annual"
	PlanNewPlanVersionParamsReplacePricesLicenseAllocationPriceCadenceCustom     PlanNewPlanVersionParamsReplacePricesLicenseAllocationPriceCadence = "custom"
)

func (r PlanNewPlanVersionParamsReplacePricesLicenseAllocationPriceCadence) IsKnown() bool {
	switch r {
	case PlanNewPlanVersionParamsReplacePricesLicenseAllocationPriceCadenceOneTime, PlanNewPlanVersionParamsReplacePricesLicenseAllocationPriceCadenceMonthly, PlanNewPlanVersionParamsReplacePricesLicenseAllocationPriceCadenceQuarterly, PlanNewPlanVersionParamsReplacePricesLicenseAllocationPriceCadenceSemiAnnual, PlanNewPlanVersionParamsReplacePricesLicenseAllocationPriceCadenceAnnual, PlanNewPlanVersionParamsReplacePricesLicenseAllocationPriceCadenceCustom:
		return true
	}
	return false
}

type PlanNewPlanVersionParamsReplacePricesLicenseAllocationPriceLicenseAllocation struct {
	// The amount of credits granted per active license per cadence.
	Amount param.Field[string] `json:"amount" api:"required"`
	// The currency of the license allocation.
	Currency param.Field[string] `json:"currency" api:"required"`
	// When True, overage beyond the allocation is written off.
	WriteOffOverage param.Field[bool] `json:"write_off_overage"`
}

func (r PlanNewPlanVersionParamsReplacePricesLicenseAllocationPriceLicenseAllocation) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// The pricing model type
type PlanNewPlanVersionParamsReplacePricesLicenseAllocationPriceModelType string

const (
	PlanNewPlanVersionParamsReplacePricesLicenseAllocationPriceModelTypeUnit PlanNewPlanVersionParamsReplacePricesLicenseAllocationPriceModelType = "unit"
)

func (r PlanNewPlanVersionParamsReplacePricesLicenseAllocationPriceModelType) IsKnown() bool {
	switch r {
	case PlanNewPlanVersionParamsReplacePricesLicenseAllocationPriceModelTypeUnit:
		return true
	}
	return false
}

type PlanNewPlanVersionParamsReplacePricesLicenseAllocationPriceConversionRateConfig struct {
	ConversionRateType param.Field[PlanNewPlanVersionParamsReplacePricesLicenseAllocationPriceConversionRateConfigConversionRateType] `json:"conversion_rate_type" api:"required"`
	TieredConfig       param.Field[shared.ConversionRateTieredConfigParam]                                                            `json:"tiered_config"`
	UnitConfig         param.Field[shared.ConversionRateUnitConfigParam]                                                              `json:"unit_config"`
}

func (r PlanNewPlanVersionParamsReplacePricesLicenseAllocationPriceConversionRateConfig) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r PlanNewPlanVersionParamsReplacePricesLicenseAllocationPriceConversionRateConfig) ImplementsPlanNewPlanVersionParamsReplacePricesLicenseAllocationPriceConversionRateConfigUnion() {
}

// Satisfied by [shared.UnitConversionRateConfigParam],
// [shared.TieredConversionRateConfigParam],
// [PlanNewPlanVersionParamsReplacePricesLicenseAllocationPriceConversionRateConfig].
type PlanNewPlanVersionParamsReplacePricesLicenseAllocationPriceConversionRateConfigUnion interface {
	ImplementsPlanNewPlanVersionParamsReplacePricesLicenseAllocationPriceConversionRateConfigUnion()
}

type PlanNewPlanVersionParamsReplacePricesLicenseAllocationPriceConversionRateConfigConversionRateType string

const (
	PlanNewPlanVersionParamsReplacePricesLicenseAllocationPriceConversionRateConfigConversionRateTypeUnit   PlanNewPlanVersionParamsReplacePricesLicenseAllocationPriceConversionRateConfigConversionRateType = "unit"
	PlanNewPlanVersionParamsReplacePricesLicenseAllocationPriceConversionRateConfigConversionRateTypeTiered PlanNewPlanVersionParamsReplacePricesLicenseAllocationPriceConversionRateConfigConversionRateType = "tiered"
)

func (r PlanNewPlanVersionParamsReplacePricesLicenseAllocationPriceConversionRateConfigConversionRateType) IsKnown() bool {
	switch r {
	case PlanNewPlanVersionParamsReplacePricesLicenseAllocationPriceConversionRateConfigConversionRateTypeUnit, PlanNewPlanVersionParamsReplacePricesLicenseAllocationPriceConversionRateConfigConversionRateTypeTiered:
		return true
	}
	return false
}

// New plan price request body params.
type PlanNewPlanVersionParamsReplacePricesPrice struct {
	// The cadence to bill for this price on.
	Cadence param.Field[PlanNewPlanVersionParamsReplacePricesPriceCadence] `json:"cadence" api:"required"`
	// The id of the item the price will be associated with.
	ItemID param.Field[string] `json:"item_id" api:"required"`
	// The pricing model type
	ModelType param.Field[PlanNewPlanVersionParamsReplacePricesPriceModelType] `json:"model_type" api:"required"`
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
	Currency                   param.Field[string]      `json:"currency"`
	DailyCreditAllowanceConfig param.Field[interface{}] `json:"daily_credit_allowance_config"`
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
	GroupedTieredMatrixConfig         param.Field[interface{}] `json:"grouped_tiered_matrix_config"`
	GroupedTieredPackageConfig        param.Field[interface{}] `json:"grouped_tiered_package_config"`
	GroupedWithMeteredMinimumConfig   param.Field[interface{}] `json:"grouped_with_metered_minimum_config"`
	GroupedWithMinMaxThresholdsConfig param.Field[interface{}] `json:"grouped_with_min_max_thresholds_config"`
	GroupedWithProratedMinimumConfig  param.Field[interface{}] `json:"grouped_with_prorated_minimum_config"`
	// The property used to group this price on an invoice
	InvoiceGroupingKey param.Field[string] `json:"invoice_grouping_key"`
	// Within each billing cycle, specifies the cadence at which invoices are produced.
	// If unspecified, a single invoice is produced per billing cycle.
	InvoicingCycleConfiguration param.Field[shared.NewBillingCycleConfigurationParam] `json:"invoicing_cycle_configuration"`
	// The ID of the license type to associate with this price. On a usage price this
	// also marks the price as eligible to draw down from that license type's
	// allocation; a usage price created without it is billed normally. Usage prices
	// with a license type must use the `unit` model, and only draw down when their
	// currency matches the allocation's.
	LicenseTypeID param.Field[string] `json:"license_type_id"`
	// Configuration for matrix pricing
	MatrixConfig param.Field[shared.MatrixConfigParam] `json:"matrix_config"`
	// Configuration for matrix_with_allocation pricing
	MatrixWithAllocationConfig         param.Field[shared.MatrixWithAllocationConfigParam] `json:"matrix_with_allocation_config"`
	MatrixWithDisplayNameConfig        param.Field[interface{}]                            `json:"matrix_with_display_name_config"`
	MatrixWithThresholdDiscountsConfig param.Field[interface{}]                            `json:"matrix_with_threshold_discounts_config"`
	MaxGroupTieredPackageConfig        param.Field[interface{}]                            `json:"max_group_tiered_package_config"`
	Metadata                           param.Field[interface{}]                            `json:"metadata"`
	MeteredAllowanceConfig             param.Field[interface{}]                            `json:"metered_allowance_config"`
	MinimumCompositeConfig             param.Field[interface{}]                            `json:"minimum_composite_config"`
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
	TieredConfig                     param.Field[shared.TieredConfigParam] `json:"tiered_config"`
	TieredMatrixWithAllocationConfig param.Field[interface{}]              `json:"tiered_matrix_with_allocation_config"`
	TieredPackageConfig              param.Field[interface{}]              `json:"tiered_package_config"`
	TieredPackageWithMinimumConfig   param.Field[interface{}]              `json:"tiered_package_with_minimum_config"`
	TieredWithMinimumConfig          param.Field[interface{}]              `json:"tiered_with_minimum_config"`
	TieredWithProrationConfig        param.Field[interface{}]              `json:"tiered_with_proration_config"`
	// Configuration for unit pricing
	UnitConfig               param.Field[shared.UnitConfigParam] `json:"unit_config"`
	UnitWithMultiplierConfig param.Field[interface{}]            `json:"unit_with_multiplier_config"`
	UnitWithPercentConfig    param.Field[interface{}]            `json:"unit_with_percent_config"`
	UnitWithProrationConfig  param.Field[interface{}]            `json:"unit_with_proration_config"`
}

func (r PlanNewPlanVersionParamsReplacePricesPrice) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r PlanNewPlanVersionParamsReplacePricesPrice) ImplementsPlanNewPlanVersionParamsReplacePricesPriceUnion() {
}

// New plan price request body params.
//
// Satisfied by [shared.NewPlanUnitPriceParam], [shared.NewPlanTieredPriceParam],
// [shared.NewPlanBulkPriceParam],
// [PlanNewPlanVersionParamsReplacePricesPriceNewPlanBulkWithFiltersPrice],
// [shared.NewPlanPackagePriceParam], [shared.NewPlanMatrixPriceParam],
// [shared.NewPlanThresholdTotalAmountPriceParam],
// [shared.NewPlanTieredPackagePriceParam],
// [shared.NewPlanTieredWithMinimumPriceParam],
// [shared.NewPlanGroupedTieredPriceParam],
// [PlanNewPlanVersionParamsReplacePricesPriceNewPlanGroupedTieredMatrixPrice],
// [shared.NewPlanTieredPackageWithMinimumPriceParam],
// [shared.NewPlanPackageWithAllocationPriceParam],
// [shared.NewPlanUnitWithPercentPriceParam],
// [shared.NewPlanMatrixWithAllocationPriceParam],
// [PlanNewPlanVersionParamsReplacePricesPriceNewPlanTieredMatrixWithAllocationPrice],
// [PlanNewPlanVersionParamsReplacePricesPriceNewPlanMatrixWithThresholdDiscountsPrice],
// [PlanNewPlanVersionParamsReplacePricesPriceNewPlanTieredWithProrationPrice],
// [shared.NewPlanUnitWithProrationPriceParam],
// [shared.NewPlanGroupedAllocationPriceParam],
// [shared.NewPlanBulkWithProrationPriceParam],
// [shared.NewPlanGroupedWithProratedMinimumPriceParam],
// [shared.NewPlanGroupedWithMeteredMinimumPriceParam],
// [PlanNewPlanVersionParamsReplacePricesPriceNewPlanGroupedWithMinMaxThresholdsPrice],
// [shared.NewPlanMatrixWithDisplayNamePriceParam],
// [shared.NewPlanGroupedTieredPackagePriceParam],
// [shared.NewPlanMaxGroupTieredPackagePriceParam],
// [shared.NewPlanScalableMatrixWithUnitPricingPriceParam],
// [shared.NewPlanScalableMatrixWithTieredPricingPriceParam],
// [shared.NewPlanCumulativeGroupedBulkPriceParam],
// [PlanNewPlanVersionParamsReplacePricesPriceNewPlanCumulativeGroupedAllocationPrice],
// [PlanNewPlanVersionParamsReplacePricesPriceNewPlanDailyCreditAllowancePrice],
// [PlanNewPlanVersionParamsReplacePricesPriceNewPlanMeteredAllowancePrice],
// [shared.NewPlanMinimumCompositePriceParam],
// [PlanNewPlanVersionParamsReplacePricesPriceNewPlanPercentCompositePrice],
// [PlanNewPlanVersionParamsReplacePricesPriceNewPlanEventOutputPrice],
// [PlanNewPlanVersionParamsReplacePricesPriceNewPlanUnitWithMultiplierPrice],
// [PlanNewPlanVersionParamsReplacePricesPrice].
type PlanNewPlanVersionParamsReplacePricesPriceUnion interface {
	ImplementsPlanNewPlanVersionParamsReplacePricesPriceUnion()
}

type PlanNewPlanVersionParamsReplacePricesPriceNewPlanBulkWithFiltersPrice struct {
	// Configuration for bulk_with_filters pricing
	BulkWithFiltersConfig param.Field[PlanNewPlanVersionParamsReplacePricesPriceNewPlanBulkWithFiltersPriceBulkWithFiltersConfig] `json:"bulk_with_filters_config" api:"required"`
	// The cadence to bill for this price on.
	Cadence param.Field[PlanNewPlanVersionParamsReplacePricesPriceNewPlanBulkWithFiltersPriceCadence] `json:"cadence" api:"required"`
	// The id of the item the price will be associated with.
	ItemID param.Field[string] `json:"item_id" api:"required"`
	// The pricing model type
	ModelType param.Field[PlanNewPlanVersionParamsReplacePricesPriceNewPlanBulkWithFiltersPriceModelType] `json:"model_type" api:"required"`
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
	ConversionRateConfig param.Field[PlanNewPlanVersionParamsReplacePricesPriceNewPlanBulkWithFiltersPriceConversionRateConfigUnion] `json:"conversion_rate_config"`
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
	// The ID of the license type to associate with this price. On a usage price this
	// also marks the price as eligible to draw down from that license type's
	// allocation; a usage price created without it is billed normally. Usage prices
	// with a license type must use the `unit` model, and only draw down when their
	// currency matches the allocation's.
	LicenseTypeID param.Field[string] `json:"license_type_id"`
	// User-specified key/value pairs for the resource. Individual keys can be removed
	// by setting the value to `null`, and the entire metadata mapping can be cleared
	// by setting `metadata` to `null`.
	Metadata param.Field[map[string]string] `json:"metadata"`
	// A transient ID that can be used to reference this price when adding adjustments
	// in the same API call.
	ReferenceID param.Field[string] `json:"reference_id"`
}

func (r PlanNewPlanVersionParamsReplacePricesPriceNewPlanBulkWithFiltersPrice) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r PlanNewPlanVersionParamsReplacePricesPriceNewPlanBulkWithFiltersPrice) ImplementsPlanNewPlanVersionParamsReplacePricesPriceUnion() {
}

// Configuration for bulk_with_filters pricing
type PlanNewPlanVersionParamsReplacePricesPriceNewPlanBulkWithFiltersPriceBulkWithFiltersConfig struct {
	// Property filters to apply (all must match)
	Filters param.Field[[]PlanNewPlanVersionParamsReplacePricesPriceNewPlanBulkWithFiltersPriceBulkWithFiltersConfigFilter] `json:"filters" api:"required"`
	// Bulk tiers for rating based on total usage volume
	Tiers param.Field[[]PlanNewPlanVersionParamsReplacePricesPriceNewPlanBulkWithFiltersPriceBulkWithFiltersConfigTier] `json:"tiers" api:"required"`
}

func (r PlanNewPlanVersionParamsReplacePricesPriceNewPlanBulkWithFiltersPriceBulkWithFiltersConfig) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Configuration for a single property filter
type PlanNewPlanVersionParamsReplacePricesPriceNewPlanBulkWithFiltersPriceBulkWithFiltersConfigFilter struct {
	// Event property key to filter on
	PropertyKey param.Field[string] `json:"property_key" api:"required"`
	// Event property value to match
	PropertyValue param.Field[string] `json:"property_value" api:"required"`
}

func (r PlanNewPlanVersionParamsReplacePricesPriceNewPlanBulkWithFiltersPriceBulkWithFiltersConfigFilter) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Configuration for a single bulk pricing tier
type PlanNewPlanVersionParamsReplacePricesPriceNewPlanBulkWithFiltersPriceBulkWithFiltersConfigTier struct {
	// Amount per unit
	UnitAmount param.Field[string] `json:"unit_amount" api:"required"`
	// The lower bound for this tier
	TierLowerBound param.Field[string] `json:"tier_lower_bound"`
}

func (r PlanNewPlanVersionParamsReplacePricesPriceNewPlanBulkWithFiltersPriceBulkWithFiltersConfigTier) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// The cadence to bill for this price on.
type PlanNewPlanVersionParamsReplacePricesPriceNewPlanBulkWithFiltersPriceCadence string

const (
	PlanNewPlanVersionParamsReplacePricesPriceNewPlanBulkWithFiltersPriceCadenceOneTime    PlanNewPlanVersionParamsReplacePricesPriceNewPlanBulkWithFiltersPriceCadence = "one_time"
	PlanNewPlanVersionParamsReplacePricesPriceNewPlanBulkWithFiltersPriceCadenceMonthly    PlanNewPlanVersionParamsReplacePricesPriceNewPlanBulkWithFiltersPriceCadence = "monthly"
	PlanNewPlanVersionParamsReplacePricesPriceNewPlanBulkWithFiltersPriceCadenceQuarterly  PlanNewPlanVersionParamsReplacePricesPriceNewPlanBulkWithFiltersPriceCadence = "quarterly"
	PlanNewPlanVersionParamsReplacePricesPriceNewPlanBulkWithFiltersPriceCadenceSemiAnnual PlanNewPlanVersionParamsReplacePricesPriceNewPlanBulkWithFiltersPriceCadence = "semi_annual"
	PlanNewPlanVersionParamsReplacePricesPriceNewPlanBulkWithFiltersPriceCadenceAnnual     PlanNewPlanVersionParamsReplacePricesPriceNewPlanBulkWithFiltersPriceCadence = "annual"
	PlanNewPlanVersionParamsReplacePricesPriceNewPlanBulkWithFiltersPriceCadenceCustom     PlanNewPlanVersionParamsReplacePricesPriceNewPlanBulkWithFiltersPriceCadence = "custom"
)

func (r PlanNewPlanVersionParamsReplacePricesPriceNewPlanBulkWithFiltersPriceCadence) IsKnown() bool {
	switch r {
	case PlanNewPlanVersionParamsReplacePricesPriceNewPlanBulkWithFiltersPriceCadenceOneTime, PlanNewPlanVersionParamsReplacePricesPriceNewPlanBulkWithFiltersPriceCadenceMonthly, PlanNewPlanVersionParamsReplacePricesPriceNewPlanBulkWithFiltersPriceCadenceQuarterly, PlanNewPlanVersionParamsReplacePricesPriceNewPlanBulkWithFiltersPriceCadenceSemiAnnual, PlanNewPlanVersionParamsReplacePricesPriceNewPlanBulkWithFiltersPriceCadenceAnnual, PlanNewPlanVersionParamsReplacePricesPriceNewPlanBulkWithFiltersPriceCadenceCustom:
		return true
	}
	return false
}

// The pricing model type
type PlanNewPlanVersionParamsReplacePricesPriceNewPlanBulkWithFiltersPriceModelType string

const (
	PlanNewPlanVersionParamsReplacePricesPriceNewPlanBulkWithFiltersPriceModelTypeBulkWithFilters PlanNewPlanVersionParamsReplacePricesPriceNewPlanBulkWithFiltersPriceModelType = "bulk_with_filters"
)

func (r PlanNewPlanVersionParamsReplacePricesPriceNewPlanBulkWithFiltersPriceModelType) IsKnown() bool {
	switch r {
	case PlanNewPlanVersionParamsReplacePricesPriceNewPlanBulkWithFiltersPriceModelTypeBulkWithFilters:
		return true
	}
	return false
}

type PlanNewPlanVersionParamsReplacePricesPriceNewPlanBulkWithFiltersPriceConversionRateConfig struct {
	ConversionRateType param.Field[PlanNewPlanVersionParamsReplacePricesPriceNewPlanBulkWithFiltersPriceConversionRateConfigConversionRateType] `json:"conversion_rate_type" api:"required"`
	TieredConfig       param.Field[shared.ConversionRateTieredConfigParam]                                                                      `json:"tiered_config"`
	UnitConfig         param.Field[shared.ConversionRateUnitConfigParam]                                                                        `json:"unit_config"`
}

func (r PlanNewPlanVersionParamsReplacePricesPriceNewPlanBulkWithFiltersPriceConversionRateConfig) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r PlanNewPlanVersionParamsReplacePricesPriceNewPlanBulkWithFiltersPriceConversionRateConfig) ImplementsPlanNewPlanVersionParamsReplacePricesPriceNewPlanBulkWithFiltersPriceConversionRateConfigUnion() {
}

// Satisfied by [shared.UnitConversionRateConfigParam],
// [shared.TieredConversionRateConfigParam],
// [PlanNewPlanVersionParamsReplacePricesPriceNewPlanBulkWithFiltersPriceConversionRateConfig].
type PlanNewPlanVersionParamsReplacePricesPriceNewPlanBulkWithFiltersPriceConversionRateConfigUnion interface {
	ImplementsPlanNewPlanVersionParamsReplacePricesPriceNewPlanBulkWithFiltersPriceConversionRateConfigUnion()
}

type PlanNewPlanVersionParamsReplacePricesPriceNewPlanBulkWithFiltersPriceConversionRateConfigConversionRateType string

const (
	PlanNewPlanVersionParamsReplacePricesPriceNewPlanBulkWithFiltersPriceConversionRateConfigConversionRateTypeUnit   PlanNewPlanVersionParamsReplacePricesPriceNewPlanBulkWithFiltersPriceConversionRateConfigConversionRateType = "unit"
	PlanNewPlanVersionParamsReplacePricesPriceNewPlanBulkWithFiltersPriceConversionRateConfigConversionRateTypeTiered PlanNewPlanVersionParamsReplacePricesPriceNewPlanBulkWithFiltersPriceConversionRateConfigConversionRateType = "tiered"
)

func (r PlanNewPlanVersionParamsReplacePricesPriceNewPlanBulkWithFiltersPriceConversionRateConfigConversionRateType) IsKnown() bool {
	switch r {
	case PlanNewPlanVersionParamsReplacePricesPriceNewPlanBulkWithFiltersPriceConversionRateConfigConversionRateTypeUnit, PlanNewPlanVersionParamsReplacePricesPriceNewPlanBulkWithFiltersPriceConversionRateConfigConversionRateTypeTiered:
		return true
	}
	return false
}

type PlanNewPlanVersionParamsReplacePricesPriceNewPlanGroupedTieredMatrixPrice struct {
	// The cadence to bill for this price on.
	Cadence param.Field[PlanNewPlanVersionParamsReplacePricesPriceNewPlanGroupedTieredMatrixPriceCadence] `json:"cadence" api:"required"`
	// Configuration for grouped_tiered_matrix pricing
	GroupedTieredMatrixConfig param.Field[PlanNewPlanVersionParamsReplacePricesPriceNewPlanGroupedTieredMatrixPriceGroupedTieredMatrixConfig] `json:"grouped_tiered_matrix_config" api:"required"`
	// The id of the item the price will be associated with.
	ItemID param.Field[string] `json:"item_id" api:"required"`
	// The pricing model type
	ModelType param.Field[PlanNewPlanVersionParamsReplacePricesPriceNewPlanGroupedTieredMatrixPriceModelType] `json:"model_type" api:"required"`
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
	ConversionRateConfig param.Field[PlanNewPlanVersionParamsReplacePricesPriceNewPlanGroupedTieredMatrixPriceConversionRateConfigUnion] `json:"conversion_rate_config"`
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
	// The ID of the license type to associate with this price. On a usage price this
	// also marks the price as eligible to draw down from that license type's
	// allocation; a usage price created without it is billed normally. Usage prices
	// with a license type must use the `unit` model, and only draw down when their
	// currency matches the allocation's.
	LicenseTypeID param.Field[string] `json:"license_type_id"`
	// User-specified key/value pairs for the resource. Individual keys can be removed
	// by setting the value to `null`, and the entire metadata mapping can be cleared
	// by setting `metadata` to `null`.
	Metadata param.Field[map[string]string] `json:"metadata"`
	// A transient ID that can be used to reference this price when adding adjustments
	// in the same API call.
	ReferenceID param.Field[string] `json:"reference_id"`
}

func (r PlanNewPlanVersionParamsReplacePricesPriceNewPlanGroupedTieredMatrixPrice) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r PlanNewPlanVersionParamsReplacePricesPriceNewPlanGroupedTieredMatrixPrice) ImplementsPlanNewPlanVersionParamsReplacePricesPriceUnion() {
}

// The cadence to bill for this price on.
type PlanNewPlanVersionParamsReplacePricesPriceNewPlanGroupedTieredMatrixPriceCadence string

const (
	PlanNewPlanVersionParamsReplacePricesPriceNewPlanGroupedTieredMatrixPriceCadenceOneTime    PlanNewPlanVersionParamsReplacePricesPriceNewPlanGroupedTieredMatrixPriceCadence = "one_time"
	PlanNewPlanVersionParamsReplacePricesPriceNewPlanGroupedTieredMatrixPriceCadenceMonthly    PlanNewPlanVersionParamsReplacePricesPriceNewPlanGroupedTieredMatrixPriceCadence = "monthly"
	PlanNewPlanVersionParamsReplacePricesPriceNewPlanGroupedTieredMatrixPriceCadenceQuarterly  PlanNewPlanVersionParamsReplacePricesPriceNewPlanGroupedTieredMatrixPriceCadence = "quarterly"
	PlanNewPlanVersionParamsReplacePricesPriceNewPlanGroupedTieredMatrixPriceCadenceSemiAnnual PlanNewPlanVersionParamsReplacePricesPriceNewPlanGroupedTieredMatrixPriceCadence = "semi_annual"
	PlanNewPlanVersionParamsReplacePricesPriceNewPlanGroupedTieredMatrixPriceCadenceAnnual     PlanNewPlanVersionParamsReplacePricesPriceNewPlanGroupedTieredMatrixPriceCadence = "annual"
	PlanNewPlanVersionParamsReplacePricesPriceNewPlanGroupedTieredMatrixPriceCadenceCustom     PlanNewPlanVersionParamsReplacePricesPriceNewPlanGroupedTieredMatrixPriceCadence = "custom"
)

func (r PlanNewPlanVersionParamsReplacePricesPriceNewPlanGroupedTieredMatrixPriceCadence) IsKnown() bool {
	switch r {
	case PlanNewPlanVersionParamsReplacePricesPriceNewPlanGroupedTieredMatrixPriceCadenceOneTime, PlanNewPlanVersionParamsReplacePricesPriceNewPlanGroupedTieredMatrixPriceCadenceMonthly, PlanNewPlanVersionParamsReplacePricesPriceNewPlanGroupedTieredMatrixPriceCadenceQuarterly, PlanNewPlanVersionParamsReplacePricesPriceNewPlanGroupedTieredMatrixPriceCadenceSemiAnnual, PlanNewPlanVersionParamsReplacePricesPriceNewPlanGroupedTieredMatrixPriceCadenceAnnual, PlanNewPlanVersionParamsReplacePricesPriceNewPlanGroupedTieredMatrixPriceCadenceCustom:
		return true
	}
	return false
}

// Configuration for grouped_tiered_matrix pricing
type PlanNewPlanVersionParamsReplacePricesPriceNewPlanGroupedTieredMatrixPriceGroupedTieredMatrixConfig struct {
	// Per unit rate for usage whose dimension value has no configured tiers
	DefaultUnitAmount param.Field[string] `json:"default_unit_amount" api:"required"`
	// The billable metric property used to group usage before tiering
	Dimension param.Field[string] `json:"dimension" api:"required"`
	// Graduated tiers keyed by dimension value; usage for a value is tiered only
	// against its own rows
	Tiers param.Field[[]PlanNewPlanVersionParamsReplacePricesPriceNewPlanGroupedTieredMatrixPriceGroupedTieredMatrixConfigTier] `json:"tiers" api:"required"`
}

func (r PlanNewPlanVersionParamsReplacePricesPriceNewPlanGroupedTieredMatrixPriceGroupedTieredMatrixConfig) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Configuration for a single tier scoped to a dimension value
type PlanNewPlanVersionParamsReplacePricesPriceNewPlanGroupedTieredMatrixPriceGroupedTieredMatrixConfigTier struct {
	// The dimension value this tier applies to
	DimensionValue param.Field[string] `json:"dimension_value" api:"required"`
	TierLowerBound param.Field[string] `json:"tier_lower_bound" api:"required"`
	// Per unit amount
	UnitAmount param.Field[string] `json:"unit_amount" api:"required"`
}

func (r PlanNewPlanVersionParamsReplacePricesPriceNewPlanGroupedTieredMatrixPriceGroupedTieredMatrixConfigTier) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// The pricing model type
type PlanNewPlanVersionParamsReplacePricesPriceNewPlanGroupedTieredMatrixPriceModelType string

const (
	PlanNewPlanVersionParamsReplacePricesPriceNewPlanGroupedTieredMatrixPriceModelTypeGroupedTieredMatrix PlanNewPlanVersionParamsReplacePricesPriceNewPlanGroupedTieredMatrixPriceModelType = "grouped_tiered_matrix"
)

func (r PlanNewPlanVersionParamsReplacePricesPriceNewPlanGroupedTieredMatrixPriceModelType) IsKnown() bool {
	switch r {
	case PlanNewPlanVersionParamsReplacePricesPriceNewPlanGroupedTieredMatrixPriceModelTypeGroupedTieredMatrix:
		return true
	}
	return false
}

type PlanNewPlanVersionParamsReplacePricesPriceNewPlanGroupedTieredMatrixPriceConversionRateConfig struct {
	ConversionRateType param.Field[PlanNewPlanVersionParamsReplacePricesPriceNewPlanGroupedTieredMatrixPriceConversionRateConfigConversionRateType] `json:"conversion_rate_type" api:"required"`
	TieredConfig       param.Field[shared.ConversionRateTieredConfigParam]                                                                          `json:"tiered_config"`
	UnitConfig         param.Field[shared.ConversionRateUnitConfigParam]                                                                            `json:"unit_config"`
}

func (r PlanNewPlanVersionParamsReplacePricesPriceNewPlanGroupedTieredMatrixPriceConversionRateConfig) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r PlanNewPlanVersionParamsReplacePricesPriceNewPlanGroupedTieredMatrixPriceConversionRateConfig) ImplementsPlanNewPlanVersionParamsReplacePricesPriceNewPlanGroupedTieredMatrixPriceConversionRateConfigUnion() {
}

// Satisfied by [shared.UnitConversionRateConfigParam],
// [shared.TieredConversionRateConfigParam],
// [PlanNewPlanVersionParamsReplacePricesPriceNewPlanGroupedTieredMatrixPriceConversionRateConfig].
type PlanNewPlanVersionParamsReplacePricesPriceNewPlanGroupedTieredMatrixPriceConversionRateConfigUnion interface {
	ImplementsPlanNewPlanVersionParamsReplacePricesPriceNewPlanGroupedTieredMatrixPriceConversionRateConfigUnion()
}

type PlanNewPlanVersionParamsReplacePricesPriceNewPlanGroupedTieredMatrixPriceConversionRateConfigConversionRateType string

const (
	PlanNewPlanVersionParamsReplacePricesPriceNewPlanGroupedTieredMatrixPriceConversionRateConfigConversionRateTypeUnit   PlanNewPlanVersionParamsReplacePricesPriceNewPlanGroupedTieredMatrixPriceConversionRateConfigConversionRateType = "unit"
	PlanNewPlanVersionParamsReplacePricesPriceNewPlanGroupedTieredMatrixPriceConversionRateConfigConversionRateTypeTiered PlanNewPlanVersionParamsReplacePricesPriceNewPlanGroupedTieredMatrixPriceConversionRateConfigConversionRateType = "tiered"
)

func (r PlanNewPlanVersionParamsReplacePricesPriceNewPlanGroupedTieredMatrixPriceConversionRateConfigConversionRateType) IsKnown() bool {
	switch r {
	case PlanNewPlanVersionParamsReplacePricesPriceNewPlanGroupedTieredMatrixPriceConversionRateConfigConversionRateTypeUnit, PlanNewPlanVersionParamsReplacePricesPriceNewPlanGroupedTieredMatrixPriceConversionRateConfigConversionRateTypeTiered:
		return true
	}
	return false
}

type PlanNewPlanVersionParamsReplacePricesPriceNewPlanTieredMatrixWithAllocationPrice struct {
	// The cadence to bill for this price on.
	Cadence param.Field[PlanNewPlanVersionParamsReplacePricesPriceNewPlanTieredMatrixWithAllocationPriceCadence] `json:"cadence" api:"required"`
	// The id of the item the price will be associated with.
	ItemID param.Field[string] `json:"item_id" api:"required"`
	// The pricing model type
	ModelType param.Field[PlanNewPlanVersionParamsReplacePricesPriceNewPlanTieredMatrixWithAllocationPriceModelType] `json:"model_type" api:"required"`
	// The name of the price.
	Name param.Field[string] `json:"name" api:"required"`
	// Configuration for tiered_matrix_with_allocation pricing
	TieredMatrixWithAllocationConfig param.Field[PlanNewPlanVersionParamsReplacePricesPriceNewPlanTieredMatrixWithAllocationPriceTieredMatrixWithAllocationConfig] `json:"tiered_matrix_with_allocation_config" api:"required"`
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
	ConversionRateConfig param.Field[PlanNewPlanVersionParamsReplacePricesPriceNewPlanTieredMatrixWithAllocationPriceConversionRateConfigUnion] `json:"conversion_rate_config"`
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
	// The ID of the license type to associate with this price. On a usage price this
	// also marks the price as eligible to draw down from that license type's
	// allocation; a usage price created without it is billed normally. Usage prices
	// with a license type must use the `unit` model, and only draw down when their
	// currency matches the allocation's.
	LicenseTypeID param.Field[string] `json:"license_type_id"`
	// User-specified key/value pairs for the resource. Individual keys can be removed
	// by setting the value to `null`, and the entire metadata mapping can be cleared
	// by setting `metadata` to `null`.
	Metadata param.Field[map[string]string] `json:"metadata"`
	// A transient ID that can be used to reference this price when adding adjustments
	// in the same API call.
	ReferenceID param.Field[string] `json:"reference_id"`
}

func (r PlanNewPlanVersionParamsReplacePricesPriceNewPlanTieredMatrixWithAllocationPrice) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r PlanNewPlanVersionParamsReplacePricesPriceNewPlanTieredMatrixWithAllocationPrice) ImplementsPlanNewPlanVersionParamsReplacePricesPriceUnion() {
}

// The cadence to bill for this price on.
type PlanNewPlanVersionParamsReplacePricesPriceNewPlanTieredMatrixWithAllocationPriceCadence string

const (
	PlanNewPlanVersionParamsReplacePricesPriceNewPlanTieredMatrixWithAllocationPriceCadenceOneTime    PlanNewPlanVersionParamsReplacePricesPriceNewPlanTieredMatrixWithAllocationPriceCadence = "one_time"
	PlanNewPlanVersionParamsReplacePricesPriceNewPlanTieredMatrixWithAllocationPriceCadenceMonthly    PlanNewPlanVersionParamsReplacePricesPriceNewPlanTieredMatrixWithAllocationPriceCadence = "monthly"
	PlanNewPlanVersionParamsReplacePricesPriceNewPlanTieredMatrixWithAllocationPriceCadenceQuarterly  PlanNewPlanVersionParamsReplacePricesPriceNewPlanTieredMatrixWithAllocationPriceCadence = "quarterly"
	PlanNewPlanVersionParamsReplacePricesPriceNewPlanTieredMatrixWithAllocationPriceCadenceSemiAnnual PlanNewPlanVersionParamsReplacePricesPriceNewPlanTieredMatrixWithAllocationPriceCadence = "semi_annual"
	PlanNewPlanVersionParamsReplacePricesPriceNewPlanTieredMatrixWithAllocationPriceCadenceAnnual     PlanNewPlanVersionParamsReplacePricesPriceNewPlanTieredMatrixWithAllocationPriceCadence = "annual"
	PlanNewPlanVersionParamsReplacePricesPriceNewPlanTieredMatrixWithAllocationPriceCadenceCustom     PlanNewPlanVersionParamsReplacePricesPriceNewPlanTieredMatrixWithAllocationPriceCadence = "custom"
)

func (r PlanNewPlanVersionParamsReplacePricesPriceNewPlanTieredMatrixWithAllocationPriceCadence) IsKnown() bool {
	switch r {
	case PlanNewPlanVersionParamsReplacePricesPriceNewPlanTieredMatrixWithAllocationPriceCadenceOneTime, PlanNewPlanVersionParamsReplacePricesPriceNewPlanTieredMatrixWithAllocationPriceCadenceMonthly, PlanNewPlanVersionParamsReplacePricesPriceNewPlanTieredMatrixWithAllocationPriceCadenceQuarterly, PlanNewPlanVersionParamsReplacePricesPriceNewPlanTieredMatrixWithAllocationPriceCadenceSemiAnnual, PlanNewPlanVersionParamsReplacePricesPriceNewPlanTieredMatrixWithAllocationPriceCadenceAnnual, PlanNewPlanVersionParamsReplacePricesPriceNewPlanTieredMatrixWithAllocationPriceCadenceCustom:
		return true
	}
	return false
}

// The pricing model type
type PlanNewPlanVersionParamsReplacePricesPriceNewPlanTieredMatrixWithAllocationPriceModelType string

const (
	PlanNewPlanVersionParamsReplacePricesPriceNewPlanTieredMatrixWithAllocationPriceModelTypeTieredMatrixWithAllocation PlanNewPlanVersionParamsReplacePricesPriceNewPlanTieredMatrixWithAllocationPriceModelType = "tiered_matrix_with_allocation"
)

func (r PlanNewPlanVersionParamsReplacePricesPriceNewPlanTieredMatrixWithAllocationPriceModelType) IsKnown() bool {
	switch r {
	case PlanNewPlanVersionParamsReplacePricesPriceNewPlanTieredMatrixWithAllocationPriceModelTypeTieredMatrixWithAllocation:
		return true
	}
	return false
}

// Configuration for tiered_matrix_with_allocation pricing
type PlanNewPlanVersionParamsReplacePricesPriceNewPlanTieredMatrixWithAllocationPriceTieredMatrixWithAllocationConfig struct {
	// Usage allocation, pooled across all matrix cells
	Allocation param.Field[string] `json:"allocation" api:"required"`
	// Per unit rate for usage whose matrix cell has no configured tiers
	DefaultUnitAmount param.Field[string] `json:"default_unit_amount" api:"required"`
	// One or two event property values to evaluate matrix cells by
	Dimensions param.Field[[]string] `json:"dimensions" api:"required"`
	// Graduated tiers keyed by matrix cell; usage in a cell is tiered only against its
	// own rows
	Tiers param.Field[[]PlanNewPlanVersionParamsReplacePricesPriceNewPlanTieredMatrixWithAllocationPriceTieredMatrixWithAllocationConfigTier] `json:"tiers" api:"required"`
}

func (r PlanNewPlanVersionParamsReplacePricesPriceNewPlanTieredMatrixWithAllocationPriceTieredMatrixWithAllocationConfig) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Configuration for a single tier scoped to one matrix cell
type PlanNewPlanVersionParamsReplacePricesPriceNewPlanTieredMatrixWithAllocationPriceTieredMatrixWithAllocationConfigTier struct {
	// The matrix cell this tier applies to, as one or two dimension values
	DimensionValues param.Field[[]string] `json:"dimension_values" api:"required"`
	// Exclusive tier starting value. The tier runs up to and including the next bound
	// configured for the same matrix cell.
	TierLowerBound param.Field[string] `json:"tier_lower_bound" api:"required"`
	// Per unit amount
	UnitAmount param.Field[string] `json:"unit_amount" api:"required"`
}

func (r PlanNewPlanVersionParamsReplacePricesPriceNewPlanTieredMatrixWithAllocationPriceTieredMatrixWithAllocationConfigTier) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type PlanNewPlanVersionParamsReplacePricesPriceNewPlanTieredMatrixWithAllocationPriceConversionRateConfig struct {
	ConversionRateType param.Field[PlanNewPlanVersionParamsReplacePricesPriceNewPlanTieredMatrixWithAllocationPriceConversionRateConfigConversionRateType] `json:"conversion_rate_type" api:"required"`
	TieredConfig       param.Field[shared.ConversionRateTieredConfigParam]                                                                                 `json:"tiered_config"`
	UnitConfig         param.Field[shared.ConversionRateUnitConfigParam]                                                                                   `json:"unit_config"`
}

func (r PlanNewPlanVersionParamsReplacePricesPriceNewPlanTieredMatrixWithAllocationPriceConversionRateConfig) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r PlanNewPlanVersionParamsReplacePricesPriceNewPlanTieredMatrixWithAllocationPriceConversionRateConfig) ImplementsPlanNewPlanVersionParamsReplacePricesPriceNewPlanTieredMatrixWithAllocationPriceConversionRateConfigUnion() {
}

// Satisfied by [shared.UnitConversionRateConfigParam],
// [shared.TieredConversionRateConfigParam],
// [PlanNewPlanVersionParamsReplacePricesPriceNewPlanTieredMatrixWithAllocationPriceConversionRateConfig].
type PlanNewPlanVersionParamsReplacePricesPriceNewPlanTieredMatrixWithAllocationPriceConversionRateConfigUnion interface {
	ImplementsPlanNewPlanVersionParamsReplacePricesPriceNewPlanTieredMatrixWithAllocationPriceConversionRateConfigUnion()
}

type PlanNewPlanVersionParamsReplacePricesPriceNewPlanTieredMatrixWithAllocationPriceConversionRateConfigConversionRateType string

const (
	PlanNewPlanVersionParamsReplacePricesPriceNewPlanTieredMatrixWithAllocationPriceConversionRateConfigConversionRateTypeUnit   PlanNewPlanVersionParamsReplacePricesPriceNewPlanTieredMatrixWithAllocationPriceConversionRateConfigConversionRateType = "unit"
	PlanNewPlanVersionParamsReplacePricesPriceNewPlanTieredMatrixWithAllocationPriceConversionRateConfigConversionRateTypeTiered PlanNewPlanVersionParamsReplacePricesPriceNewPlanTieredMatrixWithAllocationPriceConversionRateConfigConversionRateType = "tiered"
)

func (r PlanNewPlanVersionParamsReplacePricesPriceNewPlanTieredMatrixWithAllocationPriceConversionRateConfigConversionRateType) IsKnown() bool {
	switch r {
	case PlanNewPlanVersionParamsReplacePricesPriceNewPlanTieredMatrixWithAllocationPriceConversionRateConfigConversionRateTypeUnit, PlanNewPlanVersionParamsReplacePricesPriceNewPlanTieredMatrixWithAllocationPriceConversionRateConfigConversionRateTypeTiered:
		return true
	}
	return false
}

type PlanNewPlanVersionParamsReplacePricesPriceNewPlanMatrixWithThresholdDiscountsPrice struct {
	// The cadence to bill for this price on.
	Cadence param.Field[PlanNewPlanVersionParamsReplacePricesPriceNewPlanMatrixWithThresholdDiscountsPriceCadence] `json:"cadence" api:"required"`
	// The id of the item the price will be associated with.
	ItemID param.Field[string] `json:"item_id" api:"required"`
	// Configuration for matrix_with_threshold_discounts pricing
	MatrixWithThresholdDiscountsConfig param.Field[PlanNewPlanVersionParamsReplacePricesPriceNewPlanMatrixWithThresholdDiscountsPriceMatrixWithThresholdDiscountsConfig] `json:"matrix_with_threshold_discounts_config" api:"required"`
	// The pricing model type
	ModelType param.Field[PlanNewPlanVersionParamsReplacePricesPriceNewPlanMatrixWithThresholdDiscountsPriceModelType] `json:"model_type" api:"required"`
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
	ConversionRateConfig param.Field[PlanNewPlanVersionParamsReplacePricesPriceNewPlanMatrixWithThresholdDiscountsPriceConversionRateConfigUnion] `json:"conversion_rate_config"`
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
	// The ID of the license type to associate with this price. On a usage price this
	// also marks the price as eligible to draw down from that license type's
	// allocation; a usage price created without it is billed normally. Usage prices
	// with a license type must use the `unit` model, and only draw down when their
	// currency matches the allocation's.
	LicenseTypeID param.Field[string] `json:"license_type_id"`
	// User-specified key/value pairs for the resource. Individual keys can be removed
	// by setting the value to `null`, and the entire metadata mapping can be cleared
	// by setting `metadata` to `null`.
	Metadata param.Field[map[string]string] `json:"metadata"`
	// A transient ID that can be used to reference this price when adding adjustments
	// in the same API call.
	ReferenceID param.Field[string] `json:"reference_id"`
}

func (r PlanNewPlanVersionParamsReplacePricesPriceNewPlanMatrixWithThresholdDiscountsPrice) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r PlanNewPlanVersionParamsReplacePricesPriceNewPlanMatrixWithThresholdDiscountsPrice) ImplementsPlanNewPlanVersionParamsReplacePricesPriceUnion() {
}

// The cadence to bill for this price on.
type PlanNewPlanVersionParamsReplacePricesPriceNewPlanMatrixWithThresholdDiscountsPriceCadence string

const (
	PlanNewPlanVersionParamsReplacePricesPriceNewPlanMatrixWithThresholdDiscountsPriceCadenceOneTime    PlanNewPlanVersionParamsReplacePricesPriceNewPlanMatrixWithThresholdDiscountsPriceCadence = "one_time"
	PlanNewPlanVersionParamsReplacePricesPriceNewPlanMatrixWithThresholdDiscountsPriceCadenceMonthly    PlanNewPlanVersionParamsReplacePricesPriceNewPlanMatrixWithThresholdDiscountsPriceCadence = "monthly"
	PlanNewPlanVersionParamsReplacePricesPriceNewPlanMatrixWithThresholdDiscountsPriceCadenceQuarterly  PlanNewPlanVersionParamsReplacePricesPriceNewPlanMatrixWithThresholdDiscountsPriceCadence = "quarterly"
	PlanNewPlanVersionParamsReplacePricesPriceNewPlanMatrixWithThresholdDiscountsPriceCadenceSemiAnnual PlanNewPlanVersionParamsReplacePricesPriceNewPlanMatrixWithThresholdDiscountsPriceCadence = "semi_annual"
	PlanNewPlanVersionParamsReplacePricesPriceNewPlanMatrixWithThresholdDiscountsPriceCadenceAnnual     PlanNewPlanVersionParamsReplacePricesPriceNewPlanMatrixWithThresholdDiscountsPriceCadence = "annual"
	PlanNewPlanVersionParamsReplacePricesPriceNewPlanMatrixWithThresholdDiscountsPriceCadenceCustom     PlanNewPlanVersionParamsReplacePricesPriceNewPlanMatrixWithThresholdDiscountsPriceCadence = "custom"
)

func (r PlanNewPlanVersionParamsReplacePricesPriceNewPlanMatrixWithThresholdDiscountsPriceCadence) IsKnown() bool {
	switch r {
	case PlanNewPlanVersionParamsReplacePricesPriceNewPlanMatrixWithThresholdDiscountsPriceCadenceOneTime, PlanNewPlanVersionParamsReplacePricesPriceNewPlanMatrixWithThresholdDiscountsPriceCadenceMonthly, PlanNewPlanVersionParamsReplacePricesPriceNewPlanMatrixWithThresholdDiscountsPriceCadenceQuarterly, PlanNewPlanVersionParamsReplacePricesPriceNewPlanMatrixWithThresholdDiscountsPriceCadenceSemiAnnual, PlanNewPlanVersionParamsReplacePricesPriceNewPlanMatrixWithThresholdDiscountsPriceCadenceAnnual, PlanNewPlanVersionParamsReplacePricesPriceNewPlanMatrixWithThresholdDiscountsPriceCadenceCustom:
		return true
	}
	return false
}

// Configuration for matrix_with_threshold_discounts pricing
type PlanNewPlanVersionParamsReplacePricesPriceNewPlanMatrixWithThresholdDiscountsPriceMatrixWithThresholdDiscountsConfig struct {
	// Unit price used for usage that does not match any defined matrix cell.
	DefaultUnitAmount param.Field[string] `json:"default_unit_amount" api:"required"`
	// First matrix dimension key.
	FirstDimension param.Field[string] `json:"first_dimension" api:"required"`
	// Per-cell unit prices.
	MatrixValues param.Field[[]PlanNewPlanVersionParamsReplacePricesPriceNewPlanMatrixWithThresholdDiscountsPriceMatrixWithThresholdDiscountsConfigMatrixValue] `json:"matrix_values" api:"required"`
	// Optional second matrix dimension key.
	SecondDimension         param.Field[string]                                                                                                                                       `json:"second_dimension"`
	ThresholdDiscountGroups param.Field[[]PlanNewPlanVersionParamsReplacePricesPriceNewPlanMatrixWithThresholdDiscountsPriceMatrixWithThresholdDiscountsConfigThresholdDiscountGroup] `json:"threshold_discount_groups"`
}

func (r PlanNewPlanVersionParamsReplacePricesPriceNewPlanMatrixWithThresholdDiscountsPriceMatrixWithThresholdDiscountsConfig) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type PlanNewPlanVersionParamsReplacePricesPriceNewPlanMatrixWithThresholdDiscountsPriceMatrixWithThresholdDiscountsConfigMatrixValue struct {
	FirstDimensionValue  param.Field[string] `json:"first_dimension_value" api:"required"`
	UnitAmount           param.Field[string] `json:"unit_amount" api:"required"`
	SecondDimensionValue param.Field[string] `json:"second_dimension_value"`
}

func (r PlanNewPlanVersionParamsReplacePricesPriceNewPlanMatrixWithThresholdDiscountsPriceMatrixWithThresholdDiscountsConfigMatrixValue) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type PlanNewPlanVersionParamsReplacePricesPriceNewPlanMatrixWithThresholdDiscountsPriceMatrixWithThresholdDiscountsConfigThresholdDiscountGroup struct {
	// Discount rate applied to spend above the threshold.
	AboveThresholdDiscountPercentage param.Field[string] `json:"above_threshold_discount_percentage" api:"required"`
	// Discount rate applied to spend at or below the threshold. Set to 0 for no
	// baseline discount.
	BelowThresholdDiscountPercentage param.Field[string] `json:"below_threshold_discount_percentage" api:"required"`
	// Semicolon-separated list of matrix cell coordinates targeted by this group. Each
	// coordinate is `first,second` when the matrix has two dimensions, or just `first`
	// for a single-dimension matrix. Example: `blue,circle;green,triangle`.
	CellCoordinates param.Field[string] `json:"cell_coordinates" api:"required"`
	ThresholdAmount param.Field[string] `json:"threshold_amount" api:"required"`
	Description     param.Field[string] `json:"description"`
}

func (r PlanNewPlanVersionParamsReplacePricesPriceNewPlanMatrixWithThresholdDiscountsPriceMatrixWithThresholdDiscountsConfigThresholdDiscountGroup) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// The pricing model type
type PlanNewPlanVersionParamsReplacePricesPriceNewPlanMatrixWithThresholdDiscountsPriceModelType string

const (
	PlanNewPlanVersionParamsReplacePricesPriceNewPlanMatrixWithThresholdDiscountsPriceModelTypeMatrixWithThresholdDiscounts PlanNewPlanVersionParamsReplacePricesPriceNewPlanMatrixWithThresholdDiscountsPriceModelType = "matrix_with_threshold_discounts"
)

func (r PlanNewPlanVersionParamsReplacePricesPriceNewPlanMatrixWithThresholdDiscountsPriceModelType) IsKnown() bool {
	switch r {
	case PlanNewPlanVersionParamsReplacePricesPriceNewPlanMatrixWithThresholdDiscountsPriceModelTypeMatrixWithThresholdDiscounts:
		return true
	}
	return false
}

type PlanNewPlanVersionParamsReplacePricesPriceNewPlanMatrixWithThresholdDiscountsPriceConversionRateConfig struct {
	ConversionRateType param.Field[PlanNewPlanVersionParamsReplacePricesPriceNewPlanMatrixWithThresholdDiscountsPriceConversionRateConfigConversionRateType] `json:"conversion_rate_type" api:"required"`
	TieredConfig       param.Field[shared.ConversionRateTieredConfigParam]                                                                                   `json:"tiered_config"`
	UnitConfig         param.Field[shared.ConversionRateUnitConfigParam]                                                                                     `json:"unit_config"`
}

func (r PlanNewPlanVersionParamsReplacePricesPriceNewPlanMatrixWithThresholdDiscountsPriceConversionRateConfig) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r PlanNewPlanVersionParamsReplacePricesPriceNewPlanMatrixWithThresholdDiscountsPriceConversionRateConfig) ImplementsPlanNewPlanVersionParamsReplacePricesPriceNewPlanMatrixWithThresholdDiscountsPriceConversionRateConfigUnion() {
}

// Satisfied by [shared.UnitConversionRateConfigParam],
// [shared.TieredConversionRateConfigParam],
// [PlanNewPlanVersionParamsReplacePricesPriceNewPlanMatrixWithThresholdDiscountsPriceConversionRateConfig].
type PlanNewPlanVersionParamsReplacePricesPriceNewPlanMatrixWithThresholdDiscountsPriceConversionRateConfigUnion interface {
	ImplementsPlanNewPlanVersionParamsReplacePricesPriceNewPlanMatrixWithThresholdDiscountsPriceConversionRateConfigUnion()
}

type PlanNewPlanVersionParamsReplacePricesPriceNewPlanMatrixWithThresholdDiscountsPriceConversionRateConfigConversionRateType string

const (
	PlanNewPlanVersionParamsReplacePricesPriceNewPlanMatrixWithThresholdDiscountsPriceConversionRateConfigConversionRateTypeUnit   PlanNewPlanVersionParamsReplacePricesPriceNewPlanMatrixWithThresholdDiscountsPriceConversionRateConfigConversionRateType = "unit"
	PlanNewPlanVersionParamsReplacePricesPriceNewPlanMatrixWithThresholdDiscountsPriceConversionRateConfigConversionRateTypeTiered PlanNewPlanVersionParamsReplacePricesPriceNewPlanMatrixWithThresholdDiscountsPriceConversionRateConfigConversionRateType = "tiered"
)

func (r PlanNewPlanVersionParamsReplacePricesPriceNewPlanMatrixWithThresholdDiscountsPriceConversionRateConfigConversionRateType) IsKnown() bool {
	switch r {
	case PlanNewPlanVersionParamsReplacePricesPriceNewPlanMatrixWithThresholdDiscountsPriceConversionRateConfigConversionRateTypeUnit, PlanNewPlanVersionParamsReplacePricesPriceNewPlanMatrixWithThresholdDiscountsPriceConversionRateConfigConversionRateTypeTiered:
		return true
	}
	return false
}

type PlanNewPlanVersionParamsReplacePricesPriceNewPlanTieredWithProrationPrice struct {
	// The cadence to bill for this price on.
	Cadence param.Field[PlanNewPlanVersionParamsReplacePricesPriceNewPlanTieredWithProrationPriceCadence] `json:"cadence" api:"required"`
	// The id of the item the price will be associated with.
	ItemID param.Field[string] `json:"item_id" api:"required"`
	// The pricing model type
	ModelType param.Field[PlanNewPlanVersionParamsReplacePricesPriceNewPlanTieredWithProrationPriceModelType] `json:"model_type" api:"required"`
	// The name of the price.
	Name param.Field[string] `json:"name" api:"required"`
	// Configuration for tiered_with_proration pricing
	TieredWithProrationConfig param.Field[PlanNewPlanVersionParamsReplacePricesPriceNewPlanTieredWithProrationPriceTieredWithProrationConfig] `json:"tiered_with_proration_config" api:"required"`
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
	ConversionRateConfig param.Field[PlanNewPlanVersionParamsReplacePricesPriceNewPlanTieredWithProrationPriceConversionRateConfigUnion] `json:"conversion_rate_config"`
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
	// The ID of the license type to associate with this price. On a usage price this
	// also marks the price as eligible to draw down from that license type's
	// allocation; a usage price created without it is billed normally. Usage prices
	// with a license type must use the `unit` model, and only draw down when their
	// currency matches the allocation's.
	LicenseTypeID param.Field[string] `json:"license_type_id"`
	// User-specified key/value pairs for the resource. Individual keys can be removed
	// by setting the value to `null`, and the entire metadata mapping can be cleared
	// by setting `metadata` to `null`.
	Metadata param.Field[map[string]string] `json:"metadata"`
	// A transient ID that can be used to reference this price when adding adjustments
	// in the same API call.
	ReferenceID param.Field[string] `json:"reference_id"`
}

func (r PlanNewPlanVersionParamsReplacePricesPriceNewPlanTieredWithProrationPrice) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r PlanNewPlanVersionParamsReplacePricesPriceNewPlanTieredWithProrationPrice) ImplementsPlanNewPlanVersionParamsReplacePricesPriceUnion() {
}

// The cadence to bill for this price on.
type PlanNewPlanVersionParamsReplacePricesPriceNewPlanTieredWithProrationPriceCadence string

const (
	PlanNewPlanVersionParamsReplacePricesPriceNewPlanTieredWithProrationPriceCadenceOneTime    PlanNewPlanVersionParamsReplacePricesPriceNewPlanTieredWithProrationPriceCadence = "one_time"
	PlanNewPlanVersionParamsReplacePricesPriceNewPlanTieredWithProrationPriceCadenceMonthly    PlanNewPlanVersionParamsReplacePricesPriceNewPlanTieredWithProrationPriceCadence = "monthly"
	PlanNewPlanVersionParamsReplacePricesPriceNewPlanTieredWithProrationPriceCadenceQuarterly  PlanNewPlanVersionParamsReplacePricesPriceNewPlanTieredWithProrationPriceCadence = "quarterly"
	PlanNewPlanVersionParamsReplacePricesPriceNewPlanTieredWithProrationPriceCadenceSemiAnnual PlanNewPlanVersionParamsReplacePricesPriceNewPlanTieredWithProrationPriceCadence = "semi_annual"
	PlanNewPlanVersionParamsReplacePricesPriceNewPlanTieredWithProrationPriceCadenceAnnual     PlanNewPlanVersionParamsReplacePricesPriceNewPlanTieredWithProrationPriceCadence = "annual"
	PlanNewPlanVersionParamsReplacePricesPriceNewPlanTieredWithProrationPriceCadenceCustom     PlanNewPlanVersionParamsReplacePricesPriceNewPlanTieredWithProrationPriceCadence = "custom"
)

func (r PlanNewPlanVersionParamsReplacePricesPriceNewPlanTieredWithProrationPriceCadence) IsKnown() bool {
	switch r {
	case PlanNewPlanVersionParamsReplacePricesPriceNewPlanTieredWithProrationPriceCadenceOneTime, PlanNewPlanVersionParamsReplacePricesPriceNewPlanTieredWithProrationPriceCadenceMonthly, PlanNewPlanVersionParamsReplacePricesPriceNewPlanTieredWithProrationPriceCadenceQuarterly, PlanNewPlanVersionParamsReplacePricesPriceNewPlanTieredWithProrationPriceCadenceSemiAnnual, PlanNewPlanVersionParamsReplacePricesPriceNewPlanTieredWithProrationPriceCadenceAnnual, PlanNewPlanVersionParamsReplacePricesPriceNewPlanTieredWithProrationPriceCadenceCustom:
		return true
	}
	return false
}

// The pricing model type
type PlanNewPlanVersionParamsReplacePricesPriceNewPlanTieredWithProrationPriceModelType string

const (
	PlanNewPlanVersionParamsReplacePricesPriceNewPlanTieredWithProrationPriceModelTypeTieredWithProration PlanNewPlanVersionParamsReplacePricesPriceNewPlanTieredWithProrationPriceModelType = "tiered_with_proration"
)

func (r PlanNewPlanVersionParamsReplacePricesPriceNewPlanTieredWithProrationPriceModelType) IsKnown() bool {
	switch r {
	case PlanNewPlanVersionParamsReplacePricesPriceNewPlanTieredWithProrationPriceModelTypeTieredWithProration:
		return true
	}
	return false
}

// Configuration for tiered_with_proration pricing
type PlanNewPlanVersionParamsReplacePricesPriceNewPlanTieredWithProrationPriceTieredWithProrationConfig struct {
	// Tiers for rating based on total usage quantities into the specified tier with
	// proration
	Tiers param.Field[[]PlanNewPlanVersionParamsReplacePricesPriceNewPlanTieredWithProrationPriceTieredWithProrationConfigTier] `json:"tiers" api:"required"`
}

func (r PlanNewPlanVersionParamsReplacePricesPriceNewPlanTieredWithProrationPriceTieredWithProrationConfig) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Configuration for a single tiered with proration tier
type PlanNewPlanVersionParamsReplacePricesPriceNewPlanTieredWithProrationPriceTieredWithProrationConfigTier struct {
	// Inclusive tier starting value
	TierLowerBound param.Field[string] `json:"tier_lower_bound" api:"required"`
	// Amount per unit
	UnitAmount param.Field[string] `json:"unit_amount" api:"required"`
}

func (r PlanNewPlanVersionParamsReplacePricesPriceNewPlanTieredWithProrationPriceTieredWithProrationConfigTier) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type PlanNewPlanVersionParamsReplacePricesPriceNewPlanTieredWithProrationPriceConversionRateConfig struct {
	ConversionRateType param.Field[PlanNewPlanVersionParamsReplacePricesPriceNewPlanTieredWithProrationPriceConversionRateConfigConversionRateType] `json:"conversion_rate_type" api:"required"`
	TieredConfig       param.Field[shared.ConversionRateTieredConfigParam]                                                                          `json:"tiered_config"`
	UnitConfig         param.Field[shared.ConversionRateUnitConfigParam]                                                                            `json:"unit_config"`
}

func (r PlanNewPlanVersionParamsReplacePricesPriceNewPlanTieredWithProrationPriceConversionRateConfig) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r PlanNewPlanVersionParamsReplacePricesPriceNewPlanTieredWithProrationPriceConversionRateConfig) ImplementsPlanNewPlanVersionParamsReplacePricesPriceNewPlanTieredWithProrationPriceConversionRateConfigUnion() {
}

// Satisfied by [shared.UnitConversionRateConfigParam],
// [shared.TieredConversionRateConfigParam],
// [PlanNewPlanVersionParamsReplacePricesPriceNewPlanTieredWithProrationPriceConversionRateConfig].
type PlanNewPlanVersionParamsReplacePricesPriceNewPlanTieredWithProrationPriceConversionRateConfigUnion interface {
	ImplementsPlanNewPlanVersionParamsReplacePricesPriceNewPlanTieredWithProrationPriceConversionRateConfigUnion()
}

type PlanNewPlanVersionParamsReplacePricesPriceNewPlanTieredWithProrationPriceConversionRateConfigConversionRateType string

const (
	PlanNewPlanVersionParamsReplacePricesPriceNewPlanTieredWithProrationPriceConversionRateConfigConversionRateTypeUnit   PlanNewPlanVersionParamsReplacePricesPriceNewPlanTieredWithProrationPriceConversionRateConfigConversionRateType = "unit"
	PlanNewPlanVersionParamsReplacePricesPriceNewPlanTieredWithProrationPriceConversionRateConfigConversionRateTypeTiered PlanNewPlanVersionParamsReplacePricesPriceNewPlanTieredWithProrationPriceConversionRateConfigConversionRateType = "tiered"
)

func (r PlanNewPlanVersionParamsReplacePricesPriceNewPlanTieredWithProrationPriceConversionRateConfigConversionRateType) IsKnown() bool {
	switch r {
	case PlanNewPlanVersionParamsReplacePricesPriceNewPlanTieredWithProrationPriceConversionRateConfigConversionRateTypeUnit, PlanNewPlanVersionParamsReplacePricesPriceNewPlanTieredWithProrationPriceConversionRateConfigConversionRateTypeTiered:
		return true
	}
	return false
}

type PlanNewPlanVersionParamsReplacePricesPriceNewPlanGroupedWithMinMaxThresholdsPrice struct {
	// The cadence to bill for this price on.
	Cadence param.Field[PlanNewPlanVersionParamsReplacePricesPriceNewPlanGroupedWithMinMaxThresholdsPriceCadence] `json:"cadence" api:"required"`
	// Configuration for grouped_with_min_max_thresholds pricing
	GroupedWithMinMaxThresholdsConfig param.Field[PlanNewPlanVersionParamsReplacePricesPriceNewPlanGroupedWithMinMaxThresholdsPriceGroupedWithMinMaxThresholdsConfig] `json:"grouped_with_min_max_thresholds_config" api:"required"`
	// The id of the item the price will be associated with.
	ItemID param.Field[string] `json:"item_id" api:"required"`
	// The pricing model type
	ModelType param.Field[PlanNewPlanVersionParamsReplacePricesPriceNewPlanGroupedWithMinMaxThresholdsPriceModelType] `json:"model_type" api:"required"`
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
	ConversionRateConfig param.Field[PlanNewPlanVersionParamsReplacePricesPriceNewPlanGroupedWithMinMaxThresholdsPriceConversionRateConfigUnion] `json:"conversion_rate_config"`
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
	// The ID of the license type to associate with this price. On a usage price this
	// also marks the price as eligible to draw down from that license type's
	// allocation; a usage price created without it is billed normally. Usage prices
	// with a license type must use the `unit` model, and only draw down when their
	// currency matches the allocation's.
	LicenseTypeID param.Field[string] `json:"license_type_id"`
	// User-specified key/value pairs for the resource. Individual keys can be removed
	// by setting the value to `null`, and the entire metadata mapping can be cleared
	// by setting `metadata` to `null`.
	Metadata param.Field[map[string]string] `json:"metadata"`
	// A transient ID that can be used to reference this price when adding adjustments
	// in the same API call.
	ReferenceID param.Field[string] `json:"reference_id"`
}

func (r PlanNewPlanVersionParamsReplacePricesPriceNewPlanGroupedWithMinMaxThresholdsPrice) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r PlanNewPlanVersionParamsReplacePricesPriceNewPlanGroupedWithMinMaxThresholdsPrice) ImplementsPlanNewPlanVersionParamsReplacePricesPriceUnion() {
}

// The cadence to bill for this price on.
type PlanNewPlanVersionParamsReplacePricesPriceNewPlanGroupedWithMinMaxThresholdsPriceCadence string

const (
	PlanNewPlanVersionParamsReplacePricesPriceNewPlanGroupedWithMinMaxThresholdsPriceCadenceOneTime    PlanNewPlanVersionParamsReplacePricesPriceNewPlanGroupedWithMinMaxThresholdsPriceCadence = "one_time"
	PlanNewPlanVersionParamsReplacePricesPriceNewPlanGroupedWithMinMaxThresholdsPriceCadenceMonthly    PlanNewPlanVersionParamsReplacePricesPriceNewPlanGroupedWithMinMaxThresholdsPriceCadence = "monthly"
	PlanNewPlanVersionParamsReplacePricesPriceNewPlanGroupedWithMinMaxThresholdsPriceCadenceQuarterly  PlanNewPlanVersionParamsReplacePricesPriceNewPlanGroupedWithMinMaxThresholdsPriceCadence = "quarterly"
	PlanNewPlanVersionParamsReplacePricesPriceNewPlanGroupedWithMinMaxThresholdsPriceCadenceSemiAnnual PlanNewPlanVersionParamsReplacePricesPriceNewPlanGroupedWithMinMaxThresholdsPriceCadence = "semi_annual"
	PlanNewPlanVersionParamsReplacePricesPriceNewPlanGroupedWithMinMaxThresholdsPriceCadenceAnnual     PlanNewPlanVersionParamsReplacePricesPriceNewPlanGroupedWithMinMaxThresholdsPriceCadence = "annual"
	PlanNewPlanVersionParamsReplacePricesPriceNewPlanGroupedWithMinMaxThresholdsPriceCadenceCustom     PlanNewPlanVersionParamsReplacePricesPriceNewPlanGroupedWithMinMaxThresholdsPriceCadence = "custom"
)

func (r PlanNewPlanVersionParamsReplacePricesPriceNewPlanGroupedWithMinMaxThresholdsPriceCadence) IsKnown() bool {
	switch r {
	case PlanNewPlanVersionParamsReplacePricesPriceNewPlanGroupedWithMinMaxThresholdsPriceCadenceOneTime, PlanNewPlanVersionParamsReplacePricesPriceNewPlanGroupedWithMinMaxThresholdsPriceCadenceMonthly, PlanNewPlanVersionParamsReplacePricesPriceNewPlanGroupedWithMinMaxThresholdsPriceCadenceQuarterly, PlanNewPlanVersionParamsReplacePricesPriceNewPlanGroupedWithMinMaxThresholdsPriceCadenceSemiAnnual, PlanNewPlanVersionParamsReplacePricesPriceNewPlanGroupedWithMinMaxThresholdsPriceCadenceAnnual, PlanNewPlanVersionParamsReplacePricesPriceNewPlanGroupedWithMinMaxThresholdsPriceCadenceCustom:
		return true
	}
	return false
}

// Configuration for grouped_with_min_max_thresholds pricing
type PlanNewPlanVersionParamsReplacePricesPriceNewPlanGroupedWithMinMaxThresholdsPriceGroupedWithMinMaxThresholdsConfig struct {
	// The event property used to group before applying thresholds
	GroupingKey param.Field[string] `json:"grouping_key" api:"required"`
	// The maximum amount to charge each group
	MaximumCharge param.Field[string] `json:"maximum_charge" api:"required"`
	// The minimum amount to charge each group, regardless of usage
	MinimumCharge param.Field[string] `json:"minimum_charge" api:"required"`
	// The base price charged per group
	PerUnitRate param.Field[string] `json:"per_unit_rate" api:"required"`
}

func (r PlanNewPlanVersionParamsReplacePricesPriceNewPlanGroupedWithMinMaxThresholdsPriceGroupedWithMinMaxThresholdsConfig) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// The pricing model type
type PlanNewPlanVersionParamsReplacePricesPriceNewPlanGroupedWithMinMaxThresholdsPriceModelType string

const (
	PlanNewPlanVersionParamsReplacePricesPriceNewPlanGroupedWithMinMaxThresholdsPriceModelTypeGroupedWithMinMaxThresholds PlanNewPlanVersionParamsReplacePricesPriceNewPlanGroupedWithMinMaxThresholdsPriceModelType = "grouped_with_min_max_thresholds"
)

func (r PlanNewPlanVersionParamsReplacePricesPriceNewPlanGroupedWithMinMaxThresholdsPriceModelType) IsKnown() bool {
	switch r {
	case PlanNewPlanVersionParamsReplacePricesPriceNewPlanGroupedWithMinMaxThresholdsPriceModelTypeGroupedWithMinMaxThresholds:
		return true
	}
	return false
}

type PlanNewPlanVersionParamsReplacePricesPriceNewPlanGroupedWithMinMaxThresholdsPriceConversionRateConfig struct {
	ConversionRateType param.Field[PlanNewPlanVersionParamsReplacePricesPriceNewPlanGroupedWithMinMaxThresholdsPriceConversionRateConfigConversionRateType] `json:"conversion_rate_type" api:"required"`
	TieredConfig       param.Field[shared.ConversionRateTieredConfigParam]                                                                                  `json:"tiered_config"`
	UnitConfig         param.Field[shared.ConversionRateUnitConfigParam]                                                                                    `json:"unit_config"`
}

func (r PlanNewPlanVersionParamsReplacePricesPriceNewPlanGroupedWithMinMaxThresholdsPriceConversionRateConfig) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r PlanNewPlanVersionParamsReplacePricesPriceNewPlanGroupedWithMinMaxThresholdsPriceConversionRateConfig) ImplementsPlanNewPlanVersionParamsReplacePricesPriceNewPlanGroupedWithMinMaxThresholdsPriceConversionRateConfigUnion() {
}

// Satisfied by [shared.UnitConversionRateConfigParam],
// [shared.TieredConversionRateConfigParam],
// [PlanNewPlanVersionParamsReplacePricesPriceNewPlanGroupedWithMinMaxThresholdsPriceConversionRateConfig].
type PlanNewPlanVersionParamsReplacePricesPriceNewPlanGroupedWithMinMaxThresholdsPriceConversionRateConfigUnion interface {
	ImplementsPlanNewPlanVersionParamsReplacePricesPriceNewPlanGroupedWithMinMaxThresholdsPriceConversionRateConfigUnion()
}

type PlanNewPlanVersionParamsReplacePricesPriceNewPlanGroupedWithMinMaxThresholdsPriceConversionRateConfigConversionRateType string

const (
	PlanNewPlanVersionParamsReplacePricesPriceNewPlanGroupedWithMinMaxThresholdsPriceConversionRateConfigConversionRateTypeUnit   PlanNewPlanVersionParamsReplacePricesPriceNewPlanGroupedWithMinMaxThresholdsPriceConversionRateConfigConversionRateType = "unit"
	PlanNewPlanVersionParamsReplacePricesPriceNewPlanGroupedWithMinMaxThresholdsPriceConversionRateConfigConversionRateTypeTiered PlanNewPlanVersionParamsReplacePricesPriceNewPlanGroupedWithMinMaxThresholdsPriceConversionRateConfigConversionRateType = "tiered"
)

func (r PlanNewPlanVersionParamsReplacePricesPriceNewPlanGroupedWithMinMaxThresholdsPriceConversionRateConfigConversionRateType) IsKnown() bool {
	switch r {
	case PlanNewPlanVersionParamsReplacePricesPriceNewPlanGroupedWithMinMaxThresholdsPriceConversionRateConfigConversionRateTypeUnit, PlanNewPlanVersionParamsReplacePricesPriceNewPlanGroupedWithMinMaxThresholdsPriceConversionRateConfigConversionRateTypeTiered:
		return true
	}
	return false
}

type PlanNewPlanVersionParamsReplacePricesPriceNewPlanCumulativeGroupedAllocationPrice struct {
	// The cadence to bill for this price on.
	Cadence param.Field[PlanNewPlanVersionParamsReplacePricesPriceNewPlanCumulativeGroupedAllocationPriceCadence] `json:"cadence" api:"required"`
	// Configuration for cumulative_grouped_allocation pricing
	CumulativeGroupedAllocationConfig param.Field[PlanNewPlanVersionParamsReplacePricesPriceNewPlanCumulativeGroupedAllocationPriceCumulativeGroupedAllocationConfig] `json:"cumulative_grouped_allocation_config" api:"required"`
	// The id of the item the price will be associated with.
	ItemID param.Field[string] `json:"item_id" api:"required"`
	// The pricing model type
	ModelType param.Field[PlanNewPlanVersionParamsReplacePricesPriceNewPlanCumulativeGroupedAllocationPriceModelType] `json:"model_type" api:"required"`
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
	ConversionRateConfig param.Field[PlanNewPlanVersionParamsReplacePricesPriceNewPlanCumulativeGroupedAllocationPriceConversionRateConfigUnion] `json:"conversion_rate_config"`
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
	// The ID of the license type to associate with this price. On a usage price this
	// also marks the price as eligible to draw down from that license type's
	// allocation; a usage price created without it is billed normally. Usage prices
	// with a license type must use the `unit` model, and only draw down when their
	// currency matches the allocation's.
	LicenseTypeID param.Field[string] `json:"license_type_id"`
	// User-specified key/value pairs for the resource. Individual keys can be removed
	// by setting the value to `null`, and the entire metadata mapping can be cleared
	// by setting `metadata` to `null`.
	Metadata param.Field[map[string]string] `json:"metadata"`
	// A transient ID that can be used to reference this price when adding adjustments
	// in the same API call.
	ReferenceID param.Field[string] `json:"reference_id"`
}

func (r PlanNewPlanVersionParamsReplacePricesPriceNewPlanCumulativeGroupedAllocationPrice) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r PlanNewPlanVersionParamsReplacePricesPriceNewPlanCumulativeGroupedAllocationPrice) ImplementsPlanNewPlanVersionParamsReplacePricesPriceUnion() {
}

// The cadence to bill for this price on.
type PlanNewPlanVersionParamsReplacePricesPriceNewPlanCumulativeGroupedAllocationPriceCadence string

const (
	PlanNewPlanVersionParamsReplacePricesPriceNewPlanCumulativeGroupedAllocationPriceCadenceOneTime    PlanNewPlanVersionParamsReplacePricesPriceNewPlanCumulativeGroupedAllocationPriceCadence = "one_time"
	PlanNewPlanVersionParamsReplacePricesPriceNewPlanCumulativeGroupedAllocationPriceCadenceMonthly    PlanNewPlanVersionParamsReplacePricesPriceNewPlanCumulativeGroupedAllocationPriceCadence = "monthly"
	PlanNewPlanVersionParamsReplacePricesPriceNewPlanCumulativeGroupedAllocationPriceCadenceQuarterly  PlanNewPlanVersionParamsReplacePricesPriceNewPlanCumulativeGroupedAllocationPriceCadence = "quarterly"
	PlanNewPlanVersionParamsReplacePricesPriceNewPlanCumulativeGroupedAllocationPriceCadenceSemiAnnual PlanNewPlanVersionParamsReplacePricesPriceNewPlanCumulativeGroupedAllocationPriceCadence = "semi_annual"
	PlanNewPlanVersionParamsReplacePricesPriceNewPlanCumulativeGroupedAllocationPriceCadenceAnnual     PlanNewPlanVersionParamsReplacePricesPriceNewPlanCumulativeGroupedAllocationPriceCadence = "annual"
	PlanNewPlanVersionParamsReplacePricesPriceNewPlanCumulativeGroupedAllocationPriceCadenceCustom     PlanNewPlanVersionParamsReplacePricesPriceNewPlanCumulativeGroupedAllocationPriceCadence = "custom"
)

func (r PlanNewPlanVersionParamsReplacePricesPriceNewPlanCumulativeGroupedAllocationPriceCadence) IsKnown() bool {
	switch r {
	case PlanNewPlanVersionParamsReplacePricesPriceNewPlanCumulativeGroupedAllocationPriceCadenceOneTime, PlanNewPlanVersionParamsReplacePricesPriceNewPlanCumulativeGroupedAllocationPriceCadenceMonthly, PlanNewPlanVersionParamsReplacePricesPriceNewPlanCumulativeGroupedAllocationPriceCadenceQuarterly, PlanNewPlanVersionParamsReplacePricesPriceNewPlanCumulativeGroupedAllocationPriceCadenceSemiAnnual, PlanNewPlanVersionParamsReplacePricesPriceNewPlanCumulativeGroupedAllocationPriceCadenceAnnual, PlanNewPlanVersionParamsReplacePricesPriceNewPlanCumulativeGroupedAllocationPriceCadenceCustom:
		return true
	}
	return false
}

// Configuration for cumulative_grouped_allocation pricing
type PlanNewPlanVersionParamsReplacePricesPriceNewPlanCumulativeGroupedAllocationPriceCumulativeGroupedAllocationConfig struct {
	// The overall allocation across all groups
	CumulativeAllocation param.Field[string] `json:"cumulative_allocation" api:"required"`
	// The allocation per individual group
	GroupAllocation param.Field[string] `json:"group_allocation" api:"required"`
	// The event property used to group usage before applying allocations
	GroupingKey param.Field[string] `json:"grouping_key" api:"required"`
	// The amount to charge for each unit outside of the allocation
	UnitAmount param.Field[string] `json:"unit_amount" api:"required"`
}

func (r PlanNewPlanVersionParamsReplacePricesPriceNewPlanCumulativeGroupedAllocationPriceCumulativeGroupedAllocationConfig) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// The pricing model type
type PlanNewPlanVersionParamsReplacePricesPriceNewPlanCumulativeGroupedAllocationPriceModelType string

const (
	PlanNewPlanVersionParamsReplacePricesPriceNewPlanCumulativeGroupedAllocationPriceModelTypeCumulativeGroupedAllocation PlanNewPlanVersionParamsReplacePricesPriceNewPlanCumulativeGroupedAllocationPriceModelType = "cumulative_grouped_allocation"
)

func (r PlanNewPlanVersionParamsReplacePricesPriceNewPlanCumulativeGroupedAllocationPriceModelType) IsKnown() bool {
	switch r {
	case PlanNewPlanVersionParamsReplacePricesPriceNewPlanCumulativeGroupedAllocationPriceModelTypeCumulativeGroupedAllocation:
		return true
	}
	return false
}

type PlanNewPlanVersionParamsReplacePricesPriceNewPlanCumulativeGroupedAllocationPriceConversionRateConfig struct {
	ConversionRateType param.Field[PlanNewPlanVersionParamsReplacePricesPriceNewPlanCumulativeGroupedAllocationPriceConversionRateConfigConversionRateType] `json:"conversion_rate_type" api:"required"`
	TieredConfig       param.Field[shared.ConversionRateTieredConfigParam]                                                                                  `json:"tiered_config"`
	UnitConfig         param.Field[shared.ConversionRateUnitConfigParam]                                                                                    `json:"unit_config"`
}

func (r PlanNewPlanVersionParamsReplacePricesPriceNewPlanCumulativeGroupedAllocationPriceConversionRateConfig) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r PlanNewPlanVersionParamsReplacePricesPriceNewPlanCumulativeGroupedAllocationPriceConversionRateConfig) ImplementsPlanNewPlanVersionParamsReplacePricesPriceNewPlanCumulativeGroupedAllocationPriceConversionRateConfigUnion() {
}

// Satisfied by [shared.UnitConversionRateConfigParam],
// [shared.TieredConversionRateConfigParam],
// [PlanNewPlanVersionParamsReplacePricesPriceNewPlanCumulativeGroupedAllocationPriceConversionRateConfig].
type PlanNewPlanVersionParamsReplacePricesPriceNewPlanCumulativeGroupedAllocationPriceConversionRateConfigUnion interface {
	ImplementsPlanNewPlanVersionParamsReplacePricesPriceNewPlanCumulativeGroupedAllocationPriceConversionRateConfigUnion()
}

type PlanNewPlanVersionParamsReplacePricesPriceNewPlanCumulativeGroupedAllocationPriceConversionRateConfigConversionRateType string

const (
	PlanNewPlanVersionParamsReplacePricesPriceNewPlanCumulativeGroupedAllocationPriceConversionRateConfigConversionRateTypeUnit   PlanNewPlanVersionParamsReplacePricesPriceNewPlanCumulativeGroupedAllocationPriceConversionRateConfigConversionRateType = "unit"
	PlanNewPlanVersionParamsReplacePricesPriceNewPlanCumulativeGroupedAllocationPriceConversionRateConfigConversionRateTypeTiered PlanNewPlanVersionParamsReplacePricesPriceNewPlanCumulativeGroupedAllocationPriceConversionRateConfigConversionRateType = "tiered"
)

func (r PlanNewPlanVersionParamsReplacePricesPriceNewPlanCumulativeGroupedAllocationPriceConversionRateConfigConversionRateType) IsKnown() bool {
	switch r {
	case PlanNewPlanVersionParamsReplacePricesPriceNewPlanCumulativeGroupedAllocationPriceConversionRateConfigConversionRateTypeUnit, PlanNewPlanVersionParamsReplacePricesPriceNewPlanCumulativeGroupedAllocationPriceConversionRateConfigConversionRateTypeTiered:
		return true
	}
	return false
}

type PlanNewPlanVersionParamsReplacePricesPriceNewPlanDailyCreditAllowancePrice struct {
	// The cadence to bill for this price on.
	Cadence param.Field[PlanNewPlanVersionParamsReplacePricesPriceNewPlanDailyCreditAllowancePriceCadence] `json:"cadence" api:"required"`
	// Configuration for daily_credit_allowance pricing
	DailyCreditAllowanceConfig param.Field[PlanNewPlanVersionParamsReplacePricesPriceNewPlanDailyCreditAllowancePriceDailyCreditAllowanceConfig] `json:"daily_credit_allowance_config" api:"required"`
	// The id of the item the price will be associated with.
	ItemID param.Field[string] `json:"item_id" api:"required"`
	// The pricing model type
	ModelType param.Field[PlanNewPlanVersionParamsReplacePricesPriceNewPlanDailyCreditAllowancePriceModelType] `json:"model_type" api:"required"`
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
	ConversionRateConfig param.Field[PlanNewPlanVersionParamsReplacePricesPriceNewPlanDailyCreditAllowancePriceConversionRateConfigUnion] `json:"conversion_rate_config"`
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
	// The ID of the license type to associate with this price. On a usage price this
	// also marks the price as eligible to draw down from that license type's
	// allocation; a usage price created without it is billed normally. Usage prices
	// with a license type must use the `unit` model, and only draw down when their
	// currency matches the allocation's.
	LicenseTypeID param.Field[string] `json:"license_type_id"`
	// User-specified key/value pairs for the resource. Individual keys can be removed
	// by setting the value to `null`, and the entire metadata mapping can be cleared
	// by setting `metadata` to `null`.
	Metadata param.Field[map[string]string] `json:"metadata"`
	// A transient ID that can be used to reference this price when adding adjustments
	// in the same API call.
	ReferenceID param.Field[string] `json:"reference_id"`
}

func (r PlanNewPlanVersionParamsReplacePricesPriceNewPlanDailyCreditAllowancePrice) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r PlanNewPlanVersionParamsReplacePricesPriceNewPlanDailyCreditAllowancePrice) ImplementsPlanNewPlanVersionParamsReplacePricesPriceUnion() {
}

// The cadence to bill for this price on.
type PlanNewPlanVersionParamsReplacePricesPriceNewPlanDailyCreditAllowancePriceCadence string

const (
	PlanNewPlanVersionParamsReplacePricesPriceNewPlanDailyCreditAllowancePriceCadenceOneTime    PlanNewPlanVersionParamsReplacePricesPriceNewPlanDailyCreditAllowancePriceCadence = "one_time"
	PlanNewPlanVersionParamsReplacePricesPriceNewPlanDailyCreditAllowancePriceCadenceMonthly    PlanNewPlanVersionParamsReplacePricesPriceNewPlanDailyCreditAllowancePriceCadence = "monthly"
	PlanNewPlanVersionParamsReplacePricesPriceNewPlanDailyCreditAllowancePriceCadenceQuarterly  PlanNewPlanVersionParamsReplacePricesPriceNewPlanDailyCreditAllowancePriceCadence = "quarterly"
	PlanNewPlanVersionParamsReplacePricesPriceNewPlanDailyCreditAllowancePriceCadenceSemiAnnual PlanNewPlanVersionParamsReplacePricesPriceNewPlanDailyCreditAllowancePriceCadence = "semi_annual"
	PlanNewPlanVersionParamsReplacePricesPriceNewPlanDailyCreditAllowancePriceCadenceAnnual     PlanNewPlanVersionParamsReplacePricesPriceNewPlanDailyCreditAllowancePriceCadence = "annual"
	PlanNewPlanVersionParamsReplacePricesPriceNewPlanDailyCreditAllowancePriceCadenceCustom     PlanNewPlanVersionParamsReplacePricesPriceNewPlanDailyCreditAllowancePriceCadence = "custom"
)

func (r PlanNewPlanVersionParamsReplacePricesPriceNewPlanDailyCreditAllowancePriceCadence) IsKnown() bool {
	switch r {
	case PlanNewPlanVersionParamsReplacePricesPriceNewPlanDailyCreditAllowancePriceCadenceOneTime, PlanNewPlanVersionParamsReplacePricesPriceNewPlanDailyCreditAllowancePriceCadenceMonthly, PlanNewPlanVersionParamsReplacePricesPriceNewPlanDailyCreditAllowancePriceCadenceQuarterly, PlanNewPlanVersionParamsReplacePricesPriceNewPlanDailyCreditAllowancePriceCadenceSemiAnnual, PlanNewPlanVersionParamsReplacePricesPriceNewPlanDailyCreditAllowancePriceCadenceAnnual, PlanNewPlanVersionParamsReplacePricesPriceNewPlanDailyCreditAllowancePriceCadenceCustom:
		return true
	}
	return false
}

// Configuration for daily_credit_allowance pricing
type PlanNewPlanVersionParamsReplacePricesPriceNewPlanDailyCreditAllowancePriceDailyCreditAllowanceConfig struct {
	// Credits granted per day. Lose-it-or-use-it; does not roll over.
	DailyAllowance param.Field[string] `json:"daily_allowance" api:"required"`
	// Default per-unit credit rate for any usage not bucketed into a specified
	// matrix_value
	DefaultUnitAmount param.Field[string] `json:"default_unit_amount" api:"required"`
	// One or two event property values to evaluate matrix groups by
	Dimensions param.Field[[]string] `json:"dimensions" api:"required"`
	// Event property whose value identifies the day bucket the event belongs to (e.g.
	// 'event_day' set to an ISO date string in the customer's timezone). The allowance
	// resets per distinct value of this property.
	EventDayProperty param.Field[string] `json:"event_day_property" api:"required"`
	// Per-dimension credit rates
	MatrixValues param.Field[[]PlanNewPlanVersionParamsReplacePricesPriceNewPlanDailyCreditAllowancePriceDailyCreditAllowanceConfigMatrixValue] `json:"matrix_values" api:"required"`
}

func (r PlanNewPlanVersionParamsReplacePricesPriceNewPlanDailyCreditAllowancePriceDailyCreditAllowanceConfig) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Per-dimension credit price for the daily credit allowance model.
type PlanNewPlanVersionParamsReplacePricesPriceNewPlanDailyCreditAllowancePriceDailyCreditAllowanceConfigMatrixValue struct {
	// One or two matrix keys to filter usage to this value by. For example, ["model"]
	// could be used to apply a different credit rate to each AI model.
	DimensionValues param.Field[[]string] `json:"dimension_values" api:"required"`
	// Credits charged per unit of usage matching the specified dimension_values
	UnitAmount param.Field[string] `json:"unit_amount" api:"required"`
}

func (r PlanNewPlanVersionParamsReplacePricesPriceNewPlanDailyCreditAllowancePriceDailyCreditAllowanceConfigMatrixValue) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// The pricing model type
type PlanNewPlanVersionParamsReplacePricesPriceNewPlanDailyCreditAllowancePriceModelType string

const (
	PlanNewPlanVersionParamsReplacePricesPriceNewPlanDailyCreditAllowancePriceModelTypeDailyCreditAllowance PlanNewPlanVersionParamsReplacePricesPriceNewPlanDailyCreditAllowancePriceModelType = "daily_credit_allowance"
)

func (r PlanNewPlanVersionParamsReplacePricesPriceNewPlanDailyCreditAllowancePriceModelType) IsKnown() bool {
	switch r {
	case PlanNewPlanVersionParamsReplacePricesPriceNewPlanDailyCreditAllowancePriceModelTypeDailyCreditAllowance:
		return true
	}
	return false
}

type PlanNewPlanVersionParamsReplacePricesPriceNewPlanDailyCreditAllowancePriceConversionRateConfig struct {
	ConversionRateType param.Field[PlanNewPlanVersionParamsReplacePricesPriceNewPlanDailyCreditAllowancePriceConversionRateConfigConversionRateType] `json:"conversion_rate_type" api:"required"`
	TieredConfig       param.Field[shared.ConversionRateTieredConfigParam]                                                                           `json:"tiered_config"`
	UnitConfig         param.Field[shared.ConversionRateUnitConfigParam]                                                                             `json:"unit_config"`
}

func (r PlanNewPlanVersionParamsReplacePricesPriceNewPlanDailyCreditAllowancePriceConversionRateConfig) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r PlanNewPlanVersionParamsReplacePricesPriceNewPlanDailyCreditAllowancePriceConversionRateConfig) ImplementsPlanNewPlanVersionParamsReplacePricesPriceNewPlanDailyCreditAllowancePriceConversionRateConfigUnion() {
}

// Satisfied by [shared.UnitConversionRateConfigParam],
// [shared.TieredConversionRateConfigParam],
// [PlanNewPlanVersionParamsReplacePricesPriceNewPlanDailyCreditAllowancePriceConversionRateConfig].
type PlanNewPlanVersionParamsReplacePricesPriceNewPlanDailyCreditAllowancePriceConversionRateConfigUnion interface {
	ImplementsPlanNewPlanVersionParamsReplacePricesPriceNewPlanDailyCreditAllowancePriceConversionRateConfigUnion()
}

type PlanNewPlanVersionParamsReplacePricesPriceNewPlanDailyCreditAllowancePriceConversionRateConfigConversionRateType string

const (
	PlanNewPlanVersionParamsReplacePricesPriceNewPlanDailyCreditAllowancePriceConversionRateConfigConversionRateTypeUnit   PlanNewPlanVersionParamsReplacePricesPriceNewPlanDailyCreditAllowancePriceConversionRateConfigConversionRateType = "unit"
	PlanNewPlanVersionParamsReplacePricesPriceNewPlanDailyCreditAllowancePriceConversionRateConfigConversionRateTypeTiered PlanNewPlanVersionParamsReplacePricesPriceNewPlanDailyCreditAllowancePriceConversionRateConfigConversionRateType = "tiered"
)

func (r PlanNewPlanVersionParamsReplacePricesPriceNewPlanDailyCreditAllowancePriceConversionRateConfigConversionRateType) IsKnown() bool {
	switch r {
	case PlanNewPlanVersionParamsReplacePricesPriceNewPlanDailyCreditAllowancePriceConversionRateConfigConversionRateTypeUnit, PlanNewPlanVersionParamsReplacePricesPriceNewPlanDailyCreditAllowancePriceConversionRateConfigConversionRateTypeTiered:
		return true
	}
	return false
}

type PlanNewPlanVersionParamsReplacePricesPriceNewPlanMeteredAllowancePrice struct {
	// The cadence to bill for this price on.
	Cadence param.Field[PlanNewPlanVersionParamsReplacePricesPriceNewPlanMeteredAllowancePriceCadence] `json:"cadence" api:"required"`
	// The id of the item the price will be associated with.
	ItemID param.Field[string] `json:"item_id" api:"required"`
	// Configuration for metered_allowance pricing
	MeteredAllowanceConfig param.Field[PlanNewPlanVersionParamsReplacePricesPriceNewPlanMeteredAllowancePriceMeteredAllowanceConfig] `json:"metered_allowance_config" api:"required"`
	// The pricing model type
	ModelType param.Field[PlanNewPlanVersionParamsReplacePricesPriceNewPlanMeteredAllowancePriceModelType] `json:"model_type" api:"required"`
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
	ConversionRateConfig param.Field[PlanNewPlanVersionParamsReplacePricesPriceNewPlanMeteredAllowancePriceConversionRateConfigUnion] `json:"conversion_rate_config"`
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
	// The ID of the license type to associate with this price. On a usage price this
	// also marks the price as eligible to draw down from that license type's
	// allocation; a usage price created without it is billed normally. Usage prices
	// with a license type must use the `unit` model, and only draw down when their
	// currency matches the allocation's.
	LicenseTypeID param.Field[string] `json:"license_type_id"`
	// User-specified key/value pairs for the resource. Individual keys can be removed
	// by setting the value to `null`, and the entire metadata mapping can be cleared
	// by setting `metadata` to `null`.
	Metadata param.Field[map[string]string] `json:"metadata"`
	// A transient ID that can be used to reference this price when adding adjustments
	// in the same API call.
	ReferenceID param.Field[string] `json:"reference_id"`
}

func (r PlanNewPlanVersionParamsReplacePricesPriceNewPlanMeteredAllowancePrice) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r PlanNewPlanVersionParamsReplacePricesPriceNewPlanMeteredAllowancePrice) ImplementsPlanNewPlanVersionParamsReplacePricesPriceUnion() {
}

// The cadence to bill for this price on.
type PlanNewPlanVersionParamsReplacePricesPriceNewPlanMeteredAllowancePriceCadence string

const (
	PlanNewPlanVersionParamsReplacePricesPriceNewPlanMeteredAllowancePriceCadenceOneTime    PlanNewPlanVersionParamsReplacePricesPriceNewPlanMeteredAllowancePriceCadence = "one_time"
	PlanNewPlanVersionParamsReplacePricesPriceNewPlanMeteredAllowancePriceCadenceMonthly    PlanNewPlanVersionParamsReplacePricesPriceNewPlanMeteredAllowancePriceCadence = "monthly"
	PlanNewPlanVersionParamsReplacePricesPriceNewPlanMeteredAllowancePriceCadenceQuarterly  PlanNewPlanVersionParamsReplacePricesPriceNewPlanMeteredAllowancePriceCadence = "quarterly"
	PlanNewPlanVersionParamsReplacePricesPriceNewPlanMeteredAllowancePriceCadenceSemiAnnual PlanNewPlanVersionParamsReplacePricesPriceNewPlanMeteredAllowancePriceCadence = "semi_annual"
	PlanNewPlanVersionParamsReplacePricesPriceNewPlanMeteredAllowancePriceCadenceAnnual     PlanNewPlanVersionParamsReplacePricesPriceNewPlanMeteredAllowancePriceCadence = "annual"
	PlanNewPlanVersionParamsReplacePricesPriceNewPlanMeteredAllowancePriceCadenceCustom     PlanNewPlanVersionParamsReplacePricesPriceNewPlanMeteredAllowancePriceCadence = "custom"
)

func (r PlanNewPlanVersionParamsReplacePricesPriceNewPlanMeteredAllowancePriceCadence) IsKnown() bool {
	switch r {
	case PlanNewPlanVersionParamsReplacePricesPriceNewPlanMeteredAllowancePriceCadenceOneTime, PlanNewPlanVersionParamsReplacePricesPriceNewPlanMeteredAllowancePriceCadenceMonthly, PlanNewPlanVersionParamsReplacePricesPriceNewPlanMeteredAllowancePriceCadenceQuarterly, PlanNewPlanVersionParamsReplacePricesPriceNewPlanMeteredAllowancePriceCadenceSemiAnnual, PlanNewPlanVersionParamsReplacePricesPriceNewPlanMeteredAllowancePriceCadenceAnnual, PlanNewPlanVersionParamsReplacePricesPriceNewPlanMeteredAllowancePriceCadenceCustom:
		return true
	}
	return false
}

// Configuration for metered_allowance pricing
type PlanNewPlanVersionParamsReplacePricesPriceNewPlanMeteredAllowancePriceMeteredAllowanceConfig struct {
	// The grouping_key value whose summed quantity represents the allowance for this
	// period (e.g. 'storage_snapshot' emitting 3 × avg storage). Capped at consumption
	// — credit can never exceed actual usage.
	AllowanceGroupingValue param.Field[string] `json:"allowance_grouping_value" api:"required"`
	// The grouping_key value whose summed quantity represents consumption (e.g.
	// 'download'). Charged at unit_amount.
	ConsumptionGroupingValue param.Field[string] `json:"consumption_grouping_value" api:"required"`
	// Event property used to partition the metric into consumption and allowance
	// quantities (e.g. 'event_name'). The metric is queried with this key and the two
	// values below select which partition is which.
	GroupingKey param.Field[string] `json:"grouping_key" api:"required"`
	// Per-unit price applied to gross consumption and to the allowance credit.
	UnitAmount param.Field[string] `json:"unit_amount" api:"required"`
	// Sub-line label for the credit row (e.g. 'Up to 3x free egress').
	AllowanceDisplayName param.Field[string] `json:"allowance_display_name"`
	// Sub-line label for the gross consumption row (e.g. 'bytes gotten').
	ConsumptionDisplayName param.Field[string] `json:"consumption_display_name"`
}

func (r PlanNewPlanVersionParamsReplacePricesPriceNewPlanMeteredAllowancePriceMeteredAllowanceConfig) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// The pricing model type
type PlanNewPlanVersionParamsReplacePricesPriceNewPlanMeteredAllowancePriceModelType string

const (
	PlanNewPlanVersionParamsReplacePricesPriceNewPlanMeteredAllowancePriceModelTypeMeteredAllowance PlanNewPlanVersionParamsReplacePricesPriceNewPlanMeteredAllowancePriceModelType = "metered_allowance"
)

func (r PlanNewPlanVersionParamsReplacePricesPriceNewPlanMeteredAllowancePriceModelType) IsKnown() bool {
	switch r {
	case PlanNewPlanVersionParamsReplacePricesPriceNewPlanMeteredAllowancePriceModelTypeMeteredAllowance:
		return true
	}
	return false
}

type PlanNewPlanVersionParamsReplacePricesPriceNewPlanMeteredAllowancePriceConversionRateConfig struct {
	ConversionRateType param.Field[PlanNewPlanVersionParamsReplacePricesPriceNewPlanMeteredAllowancePriceConversionRateConfigConversionRateType] `json:"conversion_rate_type" api:"required"`
	TieredConfig       param.Field[shared.ConversionRateTieredConfigParam]                                                                       `json:"tiered_config"`
	UnitConfig         param.Field[shared.ConversionRateUnitConfigParam]                                                                         `json:"unit_config"`
}

func (r PlanNewPlanVersionParamsReplacePricesPriceNewPlanMeteredAllowancePriceConversionRateConfig) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r PlanNewPlanVersionParamsReplacePricesPriceNewPlanMeteredAllowancePriceConversionRateConfig) ImplementsPlanNewPlanVersionParamsReplacePricesPriceNewPlanMeteredAllowancePriceConversionRateConfigUnion() {
}

// Satisfied by [shared.UnitConversionRateConfigParam],
// [shared.TieredConversionRateConfigParam],
// [PlanNewPlanVersionParamsReplacePricesPriceNewPlanMeteredAllowancePriceConversionRateConfig].
type PlanNewPlanVersionParamsReplacePricesPriceNewPlanMeteredAllowancePriceConversionRateConfigUnion interface {
	ImplementsPlanNewPlanVersionParamsReplacePricesPriceNewPlanMeteredAllowancePriceConversionRateConfigUnion()
}

type PlanNewPlanVersionParamsReplacePricesPriceNewPlanMeteredAllowancePriceConversionRateConfigConversionRateType string

const (
	PlanNewPlanVersionParamsReplacePricesPriceNewPlanMeteredAllowancePriceConversionRateConfigConversionRateTypeUnit   PlanNewPlanVersionParamsReplacePricesPriceNewPlanMeteredAllowancePriceConversionRateConfigConversionRateType = "unit"
	PlanNewPlanVersionParamsReplacePricesPriceNewPlanMeteredAllowancePriceConversionRateConfigConversionRateTypeTiered PlanNewPlanVersionParamsReplacePricesPriceNewPlanMeteredAllowancePriceConversionRateConfigConversionRateType = "tiered"
)

func (r PlanNewPlanVersionParamsReplacePricesPriceNewPlanMeteredAllowancePriceConversionRateConfigConversionRateType) IsKnown() bool {
	switch r {
	case PlanNewPlanVersionParamsReplacePricesPriceNewPlanMeteredAllowancePriceConversionRateConfigConversionRateTypeUnit, PlanNewPlanVersionParamsReplacePricesPriceNewPlanMeteredAllowancePriceConversionRateConfigConversionRateTypeTiered:
		return true
	}
	return false
}

type PlanNewPlanVersionParamsReplacePricesPriceNewPlanPercentCompositePrice struct {
	// The cadence to bill for this price on.
	Cadence param.Field[PlanNewPlanVersionParamsReplacePricesPriceNewPlanPercentCompositePriceCadence] `json:"cadence" api:"required"`
	// The id of the item the price will be associated with.
	ItemID param.Field[string] `json:"item_id" api:"required"`
	// The pricing model type
	ModelType param.Field[PlanNewPlanVersionParamsReplacePricesPriceNewPlanPercentCompositePriceModelType] `json:"model_type" api:"required"`
	// The name of the price.
	Name param.Field[string] `json:"name" api:"required"`
	// Configuration for percent pricing
	PercentConfig param.Field[PlanNewPlanVersionParamsReplacePricesPriceNewPlanPercentCompositePricePercentConfig] `json:"percent_config" api:"required"`
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
	ConversionRateConfig param.Field[PlanNewPlanVersionParamsReplacePricesPriceNewPlanPercentCompositePriceConversionRateConfigUnion] `json:"conversion_rate_config"`
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
	// The ID of the license type to associate with this price. On a usage price this
	// also marks the price as eligible to draw down from that license type's
	// allocation; a usage price created without it is billed normally. Usage prices
	// with a license type must use the `unit` model, and only draw down when their
	// currency matches the allocation's.
	LicenseTypeID param.Field[string] `json:"license_type_id"`
	// User-specified key/value pairs for the resource. Individual keys can be removed
	// by setting the value to `null`, and the entire metadata mapping can be cleared
	// by setting `metadata` to `null`.
	Metadata param.Field[map[string]string] `json:"metadata"`
	// A transient ID that can be used to reference this price when adding adjustments
	// in the same API call.
	ReferenceID param.Field[string] `json:"reference_id"`
}

func (r PlanNewPlanVersionParamsReplacePricesPriceNewPlanPercentCompositePrice) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r PlanNewPlanVersionParamsReplacePricesPriceNewPlanPercentCompositePrice) ImplementsPlanNewPlanVersionParamsReplacePricesPriceUnion() {
}

// The cadence to bill for this price on.
type PlanNewPlanVersionParamsReplacePricesPriceNewPlanPercentCompositePriceCadence string

const (
	PlanNewPlanVersionParamsReplacePricesPriceNewPlanPercentCompositePriceCadenceOneTime    PlanNewPlanVersionParamsReplacePricesPriceNewPlanPercentCompositePriceCadence = "one_time"
	PlanNewPlanVersionParamsReplacePricesPriceNewPlanPercentCompositePriceCadenceMonthly    PlanNewPlanVersionParamsReplacePricesPriceNewPlanPercentCompositePriceCadence = "monthly"
	PlanNewPlanVersionParamsReplacePricesPriceNewPlanPercentCompositePriceCadenceQuarterly  PlanNewPlanVersionParamsReplacePricesPriceNewPlanPercentCompositePriceCadence = "quarterly"
	PlanNewPlanVersionParamsReplacePricesPriceNewPlanPercentCompositePriceCadenceSemiAnnual PlanNewPlanVersionParamsReplacePricesPriceNewPlanPercentCompositePriceCadence = "semi_annual"
	PlanNewPlanVersionParamsReplacePricesPriceNewPlanPercentCompositePriceCadenceAnnual     PlanNewPlanVersionParamsReplacePricesPriceNewPlanPercentCompositePriceCadence = "annual"
	PlanNewPlanVersionParamsReplacePricesPriceNewPlanPercentCompositePriceCadenceCustom     PlanNewPlanVersionParamsReplacePricesPriceNewPlanPercentCompositePriceCadence = "custom"
)

func (r PlanNewPlanVersionParamsReplacePricesPriceNewPlanPercentCompositePriceCadence) IsKnown() bool {
	switch r {
	case PlanNewPlanVersionParamsReplacePricesPriceNewPlanPercentCompositePriceCadenceOneTime, PlanNewPlanVersionParamsReplacePricesPriceNewPlanPercentCompositePriceCadenceMonthly, PlanNewPlanVersionParamsReplacePricesPriceNewPlanPercentCompositePriceCadenceQuarterly, PlanNewPlanVersionParamsReplacePricesPriceNewPlanPercentCompositePriceCadenceSemiAnnual, PlanNewPlanVersionParamsReplacePricesPriceNewPlanPercentCompositePriceCadenceAnnual, PlanNewPlanVersionParamsReplacePricesPriceNewPlanPercentCompositePriceCadenceCustom:
		return true
	}
	return false
}

// The pricing model type
type PlanNewPlanVersionParamsReplacePricesPriceNewPlanPercentCompositePriceModelType string

const (
	PlanNewPlanVersionParamsReplacePricesPriceNewPlanPercentCompositePriceModelTypePercent PlanNewPlanVersionParamsReplacePricesPriceNewPlanPercentCompositePriceModelType = "percent"
)

func (r PlanNewPlanVersionParamsReplacePricesPriceNewPlanPercentCompositePriceModelType) IsKnown() bool {
	switch r {
	case PlanNewPlanVersionParamsReplacePricesPriceNewPlanPercentCompositePriceModelTypePercent:
		return true
	}
	return false
}

// Configuration for percent pricing
type PlanNewPlanVersionParamsReplacePricesPriceNewPlanPercentCompositePricePercentConfig struct {
	// Fraction of the component subtotals to charge (0 < percent <= 1).
	Percent param.Field[float64] `json:"percent" api:"required"`
	// Maximum amount to charge. If unset, the fee has no upper bound.
	MaximumAmount param.Field[string] `json:"maximum_amount"`
	// Minimum amount to charge. If unset, the fee is bounded below by 0.
	MinimumAmount param.Field[string] `json:"minimum_amount"`
	// If true, the minimum_amount is prorated based on the service period. The
	// maximum_amount is an absolute cap (never prorated), and the percent applied to
	// upstream subtotals is never prorated either.
	Prorated param.Field[bool] `json:"prorated"`
}

func (r PlanNewPlanVersionParamsReplacePricesPriceNewPlanPercentCompositePricePercentConfig) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type PlanNewPlanVersionParamsReplacePricesPriceNewPlanPercentCompositePriceConversionRateConfig struct {
	ConversionRateType param.Field[PlanNewPlanVersionParamsReplacePricesPriceNewPlanPercentCompositePriceConversionRateConfigConversionRateType] `json:"conversion_rate_type" api:"required"`
	TieredConfig       param.Field[shared.ConversionRateTieredConfigParam]                                                                       `json:"tiered_config"`
	UnitConfig         param.Field[shared.ConversionRateUnitConfigParam]                                                                         `json:"unit_config"`
}

func (r PlanNewPlanVersionParamsReplacePricesPriceNewPlanPercentCompositePriceConversionRateConfig) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r PlanNewPlanVersionParamsReplacePricesPriceNewPlanPercentCompositePriceConversionRateConfig) ImplementsPlanNewPlanVersionParamsReplacePricesPriceNewPlanPercentCompositePriceConversionRateConfigUnion() {
}

// Satisfied by [shared.UnitConversionRateConfigParam],
// [shared.TieredConversionRateConfigParam],
// [PlanNewPlanVersionParamsReplacePricesPriceNewPlanPercentCompositePriceConversionRateConfig].
type PlanNewPlanVersionParamsReplacePricesPriceNewPlanPercentCompositePriceConversionRateConfigUnion interface {
	ImplementsPlanNewPlanVersionParamsReplacePricesPriceNewPlanPercentCompositePriceConversionRateConfigUnion()
}

type PlanNewPlanVersionParamsReplacePricesPriceNewPlanPercentCompositePriceConversionRateConfigConversionRateType string

const (
	PlanNewPlanVersionParamsReplacePricesPriceNewPlanPercentCompositePriceConversionRateConfigConversionRateTypeUnit   PlanNewPlanVersionParamsReplacePricesPriceNewPlanPercentCompositePriceConversionRateConfigConversionRateType = "unit"
	PlanNewPlanVersionParamsReplacePricesPriceNewPlanPercentCompositePriceConversionRateConfigConversionRateTypeTiered PlanNewPlanVersionParamsReplacePricesPriceNewPlanPercentCompositePriceConversionRateConfigConversionRateType = "tiered"
)

func (r PlanNewPlanVersionParamsReplacePricesPriceNewPlanPercentCompositePriceConversionRateConfigConversionRateType) IsKnown() bool {
	switch r {
	case PlanNewPlanVersionParamsReplacePricesPriceNewPlanPercentCompositePriceConversionRateConfigConversionRateTypeUnit, PlanNewPlanVersionParamsReplacePricesPriceNewPlanPercentCompositePriceConversionRateConfigConversionRateTypeTiered:
		return true
	}
	return false
}

type PlanNewPlanVersionParamsReplacePricesPriceNewPlanEventOutputPrice struct {
	// The cadence to bill for this price on.
	Cadence param.Field[PlanNewPlanVersionParamsReplacePricesPriceNewPlanEventOutputPriceCadence] `json:"cadence" api:"required"`
	// Configuration for event_output pricing
	EventOutputConfig param.Field[PlanNewPlanVersionParamsReplacePricesPriceNewPlanEventOutputPriceEventOutputConfig] `json:"event_output_config" api:"required"`
	// The id of the item the price will be associated with.
	ItemID param.Field[string] `json:"item_id" api:"required"`
	// The pricing model type
	ModelType param.Field[PlanNewPlanVersionParamsReplacePricesPriceNewPlanEventOutputPriceModelType] `json:"model_type" api:"required"`
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
	ConversionRateConfig param.Field[PlanNewPlanVersionParamsReplacePricesPriceNewPlanEventOutputPriceConversionRateConfigUnion] `json:"conversion_rate_config"`
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
	// The ID of the license type to associate with this price. On a usage price this
	// also marks the price as eligible to draw down from that license type's
	// allocation; a usage price created without it is billed normally. Usage prices
	// with a license type must use the `unit` model, and only draw down when their
	// currency matches the allocation's.
	LicenseTypeID param.Field[string] `json:"license_type_id"`
	// User-specified key/value pairs for the resource. Individual keys can be removed
	// by setting the value to `null`, and the entire metadata mapping can be cleared
	// by setting `metadata` to `null`.
	Metadata param.Field[map[string]string] `json:"metadata"`
	// A transient ID that can be used to reference this price when adding adjustments
	// in the same API call.
	ReferenceID param.Field[string] `json:"reference_id"`
}

func (r PlanNewPlanVersionParamsReplacePricesPriceNewPlanEventOutputPrice) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r PlanNewPlanVersionParamsReplacePricesPriceNewPlanEventOutputPrice) ImplementsPlanNewPlanVersionParamsReplacePricesPriceUnion() {
}

// The cadence to bill for this price on.
type PlanNewPlanVersionParamsReplacePricesPriceNewPlanEventOutputPriceCadence string

const (
	PlanNewPlanVersionParamsReplacePricesPriceNewPlanEventOutputPriceCadenceOneTime    PlanNewPlanVersionParamsReplacePricesPriceNewPlanEventOutputPriceCadence = "one_time"
	PlanNewPlanVersionParamsReplacePricesPriceNewPlanEventOutputPriceCadenceMonthly    PlanNewPlanVersionParamsReplacePricesPriceNewPlanEventOutputPriceCadence = "monthly"
	PlanNewPlanVersionParamsReplacePricesPriceNewPlanEventOutputPriceCadenceQuarterly  PlanNewPlanVersionParamsReplacePricesPriceNewPlanEventOutputPriceCadence = "quarterly"
	PlanNewPlanVersionParamsReplacePricesPriceNewPlanEventOutputPriceCadenceSemiAnnual PlanNewPlanVersionParamsReplacePricesPriceNewPlanEventOutputPriceCadence = "semi_annual"
	PlanNewPlanVersionParamsReplacePricesPriceNewPlanEventOutputPriceCadenceAnnual     PlanNewPlanVersionParamsReplacePricesPriceNewPlanEventOutputPriceCadence = "annual"
	PlanNewPlanVersionParamsReplacePricesPriceNewPlanEventOutputPriceCadenceCustom     PlanNewPlanVersionParamsReplacePricesPriceNewPlanEventOutputPriceCadence = "custom"
)

func (r PlanNewPlanVersionParamsReplacePricesPriceNewPlanEventOutputPriceCadence) IsKnown() bool {
	switch r {
	case PlanNewPlanVersionParamsReplacePricesPriceNewPlanEventOutputPriceCadenceOneTime, PlanNewPlanVersionParamsReplacePricesPriceNewPlanEventOutputPriceCadenceMonthly, PlanNewPlanVersionParamsReplacePricesPriceNewPlanEventOutputPriceCadenceQuarterly, PlanNewPlanVersionParamsReplacePricesPriceNewPlanEventOutputPriceCadenceSemiAnnual, PlanNewPlanVersionParamsReplacePricesPriceNewPlanEventOutputPriceCadenceAnnual, PlanNewPlanVersionParamsReplacePricesPriceNewPlanEventOutputPriceCadenceCustom:
		return true
	}
	return false
}

// Configuration for event_output pricing
type PlanNewPlanVersionParamsReplacePricesPriceNewPlanEventOutputPriceEventOutputConfig struct {
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

func (r PlanNewPlanVersionParamsReplacePricesPriceNewPlanEventOutputPriceEventOutputConfig) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// The pricing model type
type PlanNewPlanVersionParamsReplacePricesPriceNewPlanEventOutputPriceModelType string

const (
	PlanNewPlanVersionParamsReplacePricesPriceNewPlanEventOutputPriceModelTypeEventOutput PlanNewPlanVersionParamsReplacePricesPriceNewPlanEventOutputPriceModelType = "event_output"
)

func (r PlanNewPlanVersionParamsReplacePricesPriceNewPlanEventOutputPriceModelType) IsKnown() bool {
	switch r {
	case PlanNewPlanVersionParamsReplacePricesPriceNewPlanEventOutputPriceModelTypeEventOutput:
		return true
	}
	return false
}

type PlanNewPlanVersionParamsReplacePricesPriceNewPlanEventOutputPriceConversionRateConfig struct {
	ConversionRateType param.Field[PlanNewPlanVersionParamsReplacePricesPriceNewPlanEventOutputPriceConversionRateConfigConversionRateType] `json:"conversion_rate_type" api:"required"`
	TieredConfig       param.Field[shared.ConversionRateTieredConfigParam]                                                                  `json:"tiered_config"`
	UnitConfig         param.Field[shared.ConversionRateUnitConfigParam]                                                                    `json:"unit_config"`
}

func (r PlanNewPlanVersionParamsReplacePricesPriceNewPlanEventOutputPriceConversionRateConfig) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r PlanNewPlanVersionParamsReplacePricesPriceNewPlanEventOutputPriceConversionRateConfig) ImplementsPlanNewPlanVersionParamsReplacePricesPriceNewPlanEventOutputPriceConversionRateConfigUnion() {
}

// Satisfied by [shared.UnitConversionRateConfigParam],
// [shared.TieredConversionRateConfigParam],
// [PlanNewPlanVersionParamsReplacePricesPriceNewPlanEventOutputPriceConversionRateConfig].
type PlanNewPlanVersionParamsReplacePricesPriceNewPlanEventOutputPriceConversionRateConfigUnion interface {
	ImplementsPlanNewPlanVersionParamsReplacePricesPriceNewPlanEventOutputPriceConversionRateConfigUnion()
}

type PlanNewPlanVersionParamsReplacePricesPriceNewPlanEventOutputPriceConversionRateConfigConversionRateType string

const (
	PlanNewPlanVersionParamsReplacePricesPriceNewPlanEventOutputPriceConversionRateConfigConversionRateTypeUnit   PlanNewPlanVersionParamsReplacePricesPriceNewPlanEventOutputPriceConversionRateConfigConversionRateType = "unit"
	PlanNewPlanVersionParamsReplacePricesPriceNewPlanEventOutputPriceConversionRateConfigConversionRateTypeTiered PlanNewPlanVersionParamsReplacePricesPriceNewPlanEventOutputPriceConversionRateConfigConversionRateType = "tiered"
)

func (r PlanNewPlanVersionParamsReplacePricesPriceNewPlanEventOutputPriceConversionRateConfigConversionRateType) IsKnown() bool {
	switch r {
	case PlanNewPlanVersionParamsReplacePricesPriceNewPlanEventOutputPriceConversionRateConfigConversionRateTypeUnit, PlanNewPlanVersionParamsReplacePricesPriceNewPlanEventOutputPriceConversionRateConfigConversionRateTypeTiered:
		return true
	}
	return false
}

type PlanNewPlanVersionParamsReplacePricesPriceNewPlanUnitWithMultiplierPrice struct {
	// The cadence to bill for this price on.
	Cadence param.Field[PlanNewPlanVersionParamsReplacePricesPriceNewPlanUnitWithMultiplierPriceCadence] `json:"cadence" api:"required"`
	// The id of the item the price will be associated with.
	ItemID param.Field[string] `json:"item_id" api:"required"`
	// The pricing model type
	ModelType param.Field[PlanNewPlanVersionParamsReplacePricesPriceNewPlanUnitWithMultiplierPriceModelType] `json:"model_type" api:"required"`
	// The name of the price.
	Name param.Field[string] `json:"name" api:"required"`
	// Configuration for unit_with_multiplier pricing
	UnitWithMultiplierConfig param.Field[PlanNewPlanVersionParamsReplacePricesPriceNewPlanUnitWithMultiplierPriceUnitWithMultiplierConfig] `json:"unit_with_multiplier_config" api:"required"`
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
	ConversionRateConfig param.Field[PlanNewPlanVersionParamsReplacePricesPriceNewPlanUnitWithMultiplierPriceConversionRateConfigUnion] `json:"conversion_rate_config"`
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
	// The ID of the license type to associate with this price. On a usage price this
	// also marks the price as eligible to draw down from that license type's
	// allocation; a usage price created without it is billed normally. Usage prices
	// with a license type must use the `unit` model, and only draw down when their
	// currency matches the allocation's.
	LicenseTypeID param.Field[string] `json:"license_type_id"`
	// User-specified key/value pairs for the resource. Individual keys can be removed
	// by setting the value to `null`, and the entire metadata mapping can be cleared
	// by setting `metadata` to `null`.
	Metadata param.Field[map[string]string] `json:"metadata"`
	// A transient ID that can be used to reference this price when adding adjustments
	// in the same API call.
	ReferenceID param.Field[string] `json:"reference_id"`
}

func (r PlanNewPlanVersionParamsReplacePricesPriceNewPlanUnitWithMultiplierPrice) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r PlanNewPlanVersionParamsReplacePricesPriceNewPlanUnitWithMultiplierPrice) ImplementsPlanNewPlanVersionParamsReplacePricesPriceUnion() {
}

// The cadence to bill for this price on.
type PlanNewPlanVersionParamsReplacePricesPriceNewPlanUnitWithMultiplierPriceCadence string

const (
	PlanNewPlanVersionParamsReplacePricesPriceNewPlanUnitWithMultiplierPriceCadenceOneTime    PlanNewPlanVersionParamsReplacePricesPriceNewPlanUnitWithMultiplierPriceCadence = "one_time"
	PlanNewPlanVersionParamsReplacePricesPriceNewPlanUnitWithMultiplierPriceCadenceMonthly    PlanNewPlanVersionParamsReplacePricesPriceNewPlanUnitWithMultiplierPriceCadence = "monthly"
	PlanNewPlanVersionParamsReplacePricesPriceNewPlanUnitWithMultiplierPriceCadenceQuarterly  PlanNewPlanVersionParamsReplacePricesPriceNewPlanUnitWithMultiplierPriceCadence = "quarterly"
	PlanNewPlanVersionParamsReplacePricesPriceNewPlanUnitWithMultiplierPriceCadenceSemiAnnual PlanNewPlanVersionParamsReplacePricesPriceNewPlanUnitWithMultiplierPriceCadence = "semi_annual"
	PlanNewPlanVersionParamsReplacePricesPriceNewPlanUnitWithMultiplierPriceCadenceAnnual     PlanNewPlanVersionParamsReplacePricesPriceNewPlanUnitWithMultiplierPriceCadence = "annual"
	PlanNewPlanVersionParamsReplacePricesPriceNewPlanUnitWithMultiplierPriceCadenceCustom     PlanNewPlanVersionParamsReplacePricesPriceNewPlanUnitWithMultiplierPriceCadence = "custom"
)

func (r PlanNewPlanVersionParamsReplacePricesPriceNewPlanUnitWithMultiplierPriceCadence) IsKnown() bool {
	switch r {
	case PlanNewPlanVersionParamsReplacePricesPriceNewPlanUnitWithMultiplierPriceCadenceOneTime, PlanNewPlanVersionParamsReplacePricesPriceNewPlanUnitWithMultiplierPriceCadenceMonthly, PlanNewPlanVersionParamsReplacePricesPriceNewPlanUnitWithMultiplierPriceCadenceQuarterly, PlanNewPlanVersionParamsReplacePricesPriceNewPlanUnitWithMultiplierPriceCadenceSemiAnnual, PlanNewPlanVersionParamsReplacePricesPriceNewPlanUnitWithMultiplierPriceCadenceAnnual, PlanNewPlanVersionParamsReplacePricesPriceNewPlanUnitWithMultiplierPriceCadenceCustom:
		return true
	}
	return false
}

// The pricing model type
type PlanNewPlanVersionParamsReplacePricesPriceNewPlanUnitWithMultiplierPriceModelType string

const (
	PlanNewPlanVersionParamsReplacePricesPriceNewPlanUnitWithMultiplierPriceModelTypeUnitWithMultiplier PlanNewPlanVersionParamsReplacePricesPriceNewPlanUnitWithMultiplierPriceModelType = "unit_with_multiplier"
)

func (r PlanNewPlanVersionParamsReplacePricesPriceNewPlanUnitWithMultiplierPriceModelType) IsKnown() bool {
	switch r {
	case PlanNewPlanVersionParamsReplacePricesPriceNewPlanUnitWithMultiplierPriceModelTypeUnitWithMultiplier:
		return true
	}
	return false
}

// Configuration for unit_with_multiplier pricing
type PlanNewPlanVersionParamsReplacePricesPriceNewPlanUnitWithMultiplierPriceUnitWithMultiplierConfig struct {
	// Sub-line label for list-price events that don't set display_name_key. Unlabeled
	// events with another multiplier are labeled by their multiplier (e.g. '0.8×
	// rate').
	DefaultDisplayName param.Field[string] `json:"default_display_name" api:"required"`
	// Event property holding the sub-line label for the event's rate (e.g.
	// 'Off-peak').
	DisplayNameKey param.Field[string] `json:"display_name_key" api:"required"`
	// Event property holding the multiplier applied to unit_amount (e.g. 0.8 for a 20%
	// discount). Events without a number between 0 and 1000 with at most 12 decimal
	// places are charged at unit_amount.
	MultiplierKey param.Field[string] `json:"multiplier_key" api:"required"`
	// List price per unit, before the event's multiplier is applied.
	UnitAmount param.Field[string] `json:"unit_amount" api:"required"`
}

func (r PlanNewPlanVersionParamsReplacePricesPriceNewPlanUnitWithMultiplierPriceUnitWithMultiplierConfig) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type PlanNewPlanVersionParamsReplacePricesPriceNewPlanUnitWithMultiplierPriceConversionRateConfig struct {
	ConversionRateType param.Field[PlanNewPlanVersionParamsReplacePricesPriceNewPlanUnitWithMultiplierPriceConversionRateConfigConversionRateType] `json:"conversion_rate_type" api:"required"`
	TieredConfig       param.Field[shared.ConversionRateTieredConfigParam]                                                                         `json:"tiered_config"`
	UnitConfig         param.Field[shared.ConversionRateUnitConfigParam]                                                                           `json:"unit_config"`
}

func (r PlanNewPlanVersionParamsReplacePricesPriceNewPlanUnitWithMultiplierPriceConversionRateConfig) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r PlanNewPlanVersionParamsReplacePricesPriceNewPlanUnitWithMultiplierPriceConversionRateConfig) ImplementsPlanNewPlanVersionParamsReplacePricesPriceNewPlanUnitWithMultiplierPriceConversionRateConfigUnion() {
}

// Satisfied by [shared.UnitConversionRateConfigParam],
// [shared.TieredConversionRateConfigParam],
// [PlanNewPlanVersionParamsReplacePricesPriceNewPlanUnitWithMultiplierPriceConversionRateConfig].
type PlanNewPlanVersionParamsReplacePricesPriceNewPlanUnitWithMultiplierPriceConversionRateConfigUnion interface {
	ImplementsPlanNewPlanVersionParamsReplacePricesPriceNewPlanUnitWithMultiplierPriceConversionRateConfigUnion()
}

type PlanNewPlanVersionParamsReplacePricesPriceNewPlanUnitWithMultiplierPriceConversionRateConfigConversionRateType string

const (
	PlanNewPlanVersionParamsReplacePricesPriceNewPlanUnitWithMultiplierPriceConversionRateConfigConversionRateTypeUnit   PlanNewPlanVersionParamsReplacePricesPriceNewPlanUnitWithMultiplierPriceConversionRateConfigConversionRateType = "unit"
	PlanNewPlanVersionParamsReplacePricesPriceNewPlanUnitWithMultiplierPriceConversionRateConfigConversionRateTypeTiered PlanNewPlanVersionParamsReplacePricesPriceNewPlanUnitWithMultiplierPriceConversionRateConfigConversionRateType = "tiered"
)

func (r PlanNewPlanVersionParamsReplacePricesPriceNewPlanUnitWithMultiplierPriceConversionRateConfigConversionRateType) IsKnown() bool {
	switch r {
	case PlanNewPlanVersionParamsReplacePricesPriceNewPlanUnitWithMultiplierPriceConversionRateConfigConversionRateTypeUnit, PlanNewPlanVersionParamsReplacePricesPriceNewPlanUnitWithMultiplierPriceConversionRateConfigConversionRateTypeTiered:
		return true
	}
	return false
}

// The cadence to bill for this price on.
type PlanNewPlanVersionParamsReplacePricesPriceCadence string

const (
	PlanNewPlanVersionParamsReplacePricesPriceCadenceOneTime    PlanNewPlanVersionParamsReplacePricesPriceCadence = "one_time"
	PlanNewPlanVersionParamsReplacePricesPriceCadenceMonthly    PlanNewPlanVersionParamsReplacePricesPriceCadence = "monthly"
	PlanNewPlanVersionParamsReplacePricesPriceCadenceQuarterly  PlanNewPlanVersionParamsReplacePricesPriceCadence = "quarterly"
	PlanNewPlanVersionParamsReplacePricesPriceCadenceSemiAnnual PlanNewPlanVersionParamsReplacePricesPriceCadence = "semi_annual"
	PlanNewPlanVersionParamsReplacePricesPriceCadenceAnnual     PlanNewPlanVersionParamsReplacePricesPriceCadence = "annual"
	PlanNewPlanVersionParamsReplacePricesPriceCadenceCustom     PlanNewPlanVersionParamsReplacePricesPriceCadence = "custom"
)

func (r PlanNewPlanVersionParamsReplacePricesPriceCadence) IsKnown() bool {
	switch r {
	case PlanNewPlanVersionParamsReplacePricesPriceCadenceOneTime, PlanNewPlanVersionParamsReplacePricesPriceCadenceMonthly, PlanNewPlanVersionParamsReplacePricesPriceCadenceQuarterly, PlanNewPlanVersionParamsReplacePricesPriceCadenceSemiAnnual, PlanNewPlanVersionParamsReplacePricesPriceCadenceAnnual, PlanNewPlanVersionParamsReplacePricesPriceCadenceCustom:
		return true
	}
	return false
}

// The pricing model type
type PlanNewPlanVersionParamsReplacePricesPriceModelType string

const (
	PlanNewPlanVersionParamsReplacePricesPriceModelTypeUnit                            PlanNewPlanVersionParamsReplacePricesPriceModelType = "unit"
	PlanNewPlanVersionParamsReplacePricesPriceModelTypeTiered                          PlanNewPlanVersionParamsReplacePricesPriceModelType = "tiered"
	PlanNewPlanVersionParamsReplacePricesPriceModelTypeBulk                            PlanNewPlanVersionParamsReplacePricesPriceModelType = "bulk"
	PlanNewPlanVersionParamsReplacePricesPriceModelTypeBulkWithFilters                 PlanNewPlanVersionParamsReplacePricesPriceModelType = "bulk_with_filters"
	PlanNewPlanVersionParamsReplacePricesPriceModelTypePackage                         PlanNewPlanVersionParamsReplacePricesPriceModelType = "package"
	PlanNewPlanVersionParamsReplacePricesPriceModelTypeMatrix                          PlanNewPlanVersionParamsReplacePricesPriceModelType = "matrix"
	PlanNewPlanVersionParamsReplacePricesPriceModelTypeThresholdTotalAmount            PlanNewPlanVersionParamsReplacePricesPriceModelType = "threshold_total_amount"
	PlanNewPlanVersionParamsReplacePricesPriceModelTypeTieredPackage                   PlanNewPlanVersionParamsReplacePricesPriceModelType = "tiered_package"
	PlanNewPlanVersionParamsReplacePricesPriceModelTypeTieredWithMinimum               PlanNewPlanVersionParamsReplacePricesPriceModelType = "tiered_with_minimum"
	PlanNewPlanVersionParamsReplacePricesPriceModelTypeGroupedTiered                   PlanNewPlanVersionParamsReplacePricesPriceModelType = "grouped_tiered"
	PlanNewPlanVersionParamsReplacePricesPriceModelTypeGroupedTieredMatrix             PlanNewPlanVersionParamsReplacePricesPriceModelType = "grouped_tiered_matrix"
	PlanNewPlanVersionParamsReplacePricesPriceModelTypeTieredPackageWithMinimum        PlanNewPlanVersionParamsReplacePricesPriceModelType = "tiered_package_with_minimum"
	PlanNewPlanVersionParamsReplacePricesPriceModelTypePackageWithAllocation           PlanNewPlanVersionParamsReplacePricesPriceModelType = "package_with_allocation"
	PlanNewPlanVersionParamsReplacePricesPriceModelTypeUnitWithPercent                 PlanNewPlanVersionParamsReplacePricesPriceModelType = "unit_with_percent"
	PlanNewPlanVersionParamsReplacePricesPriceModelTypeMatrixWithAllocation            PlanNewPlanVersionParamsReplacePricesPriceModelType = "matrix_with_allocation"
	PlanNewPlanVersionParamsReplacePricesPriceModelTypeTieredMatrixWithAllocation      PlanNewPlanVersionParamsReplacePricesPriceModelType = "tiered_matrix_with_allocation"
	PlanNewPlanVersionParamsReplacePricesPriceModelTypeMatrixWithThresholdDiscounts    PlanNewPlanVersionParamsReplacePricesPriceModelType = "matrix_with_threshold_discounts"
	PlanNewPlanVersionParamsReplacePricesPriceModelTypeTieredWithProration             PlanNewPlanVersionParamsReplacePricesPriceModelType = "tiered_with_proration"
	PlanNewPlanVersionParamsReplacePricesPriceModelTypeUnitWithProration               PlanNewPlanVersionParamsReplacePricesPriceModelType = "unit_with_proration"
	PlanNewPlanVersionParamsReplacePricesPriceModelTypeGroupedAllocation               PlanNewPlanVersionParamsReplacePricesPriceModelType = "grouped_allocation"
	PlanNewPlanVersionParamsReplacePricesPriceModelTypeBulkWithProration               PlanNewPlanVersionParamsReplacePricesPriceModelType = "bulk_with_proration"
	PlanNewPlanVersionParamsReplacePricesPriceModelTypeGroupedWithProratedMinimum      PlanNewPlanVersionParamsReplacePricesPriceModelType = "grouped_with_prorated_minimum"
	PlanNewPlanVersionParamsReplacePricesPriceModelTypeGroupedWithMeteredMinimum       PlanNewPlanVersionParamsReplacePricesPriceModelType = "grouped_with_metered_minimum"
	PlanNewPlanVersionParamsReplacePricesPriceModelTypeGroupedWithMinMaxThresholds     PlanNewPlanVersionParamsReplacePricesPriceModelType = "grouped_with_min_max_thresholds"
	PlanNewPlanVersionParamsReplacePricesPriceModelTypeMatrixWithDisplayName           PlanNewPlanVersionParamsReplacePricesPriceModelType = "matrix_with_display_name"
	PlanNewPlanVersionParamsReplacePricesPriceModelTypeGroupedTieredPackage            PlanNewPlanVersionParamsReplacePricesPriceModelType = "grouped_tiered_package"
	PlanNewPlanVersionParamsReplacePricesPriceModelTypeMaxGroupTieredPackage           PlanNewPlanVersionParamsReplacePricesPriceModelType = "max_group_tiered_package"
	PlanNewPlanVersionParamsReplacePricesPriceModelTypeScalableMatrixWithUnitPricing   PlanNewPlanVersionParamsReplacePricesPriceModelType = "scalable_matrix_with_unit_pricing"
	PlanNewPlanVersionParamsReplacePricesPriceModelTypeScalableMatrixWithTieredPricing PlanNewPlanVersionParamsReplacePricesPriceModelType = "scalable_matrix_with_tiered_pricing"
	PlanNewPlanVersionParamsReplacePricesPriceModelTypeCumulativeGroupedBulk           PlanNewPlanVersionParamsReplacePricesPriceModelType = "cumulative_grouped_bulk"
	PlanNewPlanVersionParamsReplacePricesPriceModelTypeCumulativeGroupedAllocation     PlanNewPlanVersionParamsReplacePricesPriceModelType = "cumulative_grouped_allocation"
	PlanNewPlanVersionParamsReplacePricesPriceModelTypeDailyCreditAllowance            PlanNewPlanVersionParamsReplacePricesPriceModelType = "daily_credit_allowance"
	PlanNewPlanVersionParamsReplacePricesPriceModelTypeMeteredAllowance                PlanNewPlanVersionParamsReplacePricesPriceModelType = "metered_allowance"
	PlanNewPlanVersionParamsReplacePricesPriceModelTypeMinimumComposite                PlanNewPlanVersionParamsReplacePricesPriceModelType = "minimum_composite"
	PlanNewPlanVersionParamsReplacePricesPriceModelTypePercent                         PlanNewPlanVersionParamsReplacePricesPriceModelType = "percent"
	PlanNewPlanVersionParamsReplacePricesPriceModelTypeEventOutput                     PlanNewPlanVersionParamsReplacePricesPriceModelType = "event_output"
	PlanNewPlanVersionParamsReplacePricesPriceModelTypeUnitWithMultiplier              PlanNewPlanVersionParamsReplacePricesPriceModelType = "unit_with_multiplier"
)

func (r PlanNewPlanVersionParamsReplacePricesPriceModelType) IsKnown() bool {
	switch r {
	case PlanNewPlanVersionParamsReplacePricesPriceModelTypeUnit, PlanNewPlanVersionParamsReplacePricesPriceModelTypeTiered, PlanNewPlanVersionParamsReplacePricesPriceModelTypeBulk, PlanNewPlanVersionParamsReplacePricesPriceModelTypeBulkWithFilters, PlanNewPlanVersionParamsReplacePricesPriceModelTypePackage, PlanNewPlanVersionParamsReplacePricesPriceModelTypeMatrix, PlanNewPlanVersionParamsReplacePricesPriceModelTypeThresholdTotalAmount, PlanNewPlanVersionParamsReplacePricesPriceModelTypeTieredPackage, PlanNewPlanVersionParamsReplacePricesPriceModelTypeTieredWithMinimum, PlanNewPlanVersionParamsReplacePricesPriceModelTypeGroupedTiered, PlanNewPlanVersionParamsReplacePricesPriceModelTypeGroupedTieredMatrix, PlanNewPlanVersionParamsReplacePricesPriceModelTypeTieredPackageWithMinimum, PlanNewPlanVersionParamsReplacePricesPriceModelTypePackageWithAllocation, PlanNewPlanVersionParamsReplacePricesPriceModelTypeUnitWithPercent, PlanNewPlanVersionParamsReplacePricesPriceModelTypeMatrixWithAllocation, PlanNewPlanVersionParamsReplacePricesPriceModelTypeTieredMatrixWithAllocation, PlanNewPlanVersionParamsReplacePricesPriceModelTypeMatrixWithThresholdDiscounts, PlanNewPlanVersionParamsReplacePricesPriceModelTypeTieredWithProration, PlanNewPlanVersionParamsReplacePricesPriceModelTypeUnitWithProration, PlanNewPlanVersionParamsReplacePricesPriceModelTypeGroupedAllocation, PlanNewPlanVersionParamsReplacePricesPriceModelTypeBulkWithProration, PlanNewPlanVersionParamsReplacePricesPriceModelTypeGroupedWithProratedMinimum, PlanNewPlanVersionParamsReplacePricesPriceModelTypeGroupedWithMeteredMinimum, PlanNewPlanVersionParamsReplacePricesPriceModelTypeGroupedWithMinMaxThresholds, PlanNewPlanVersionParamsReplacePricesPriceModelTypeMatrixWithDisplayName, PlanNewPlanVersionParamsReplacePricesPriceModelTypeGroupedTieredPackage, PlanNewPlanVersionParamsReplacePricesPriceModelTypeMaxGroupTieredPackage, PlanNewPlanVersionParamsReplacePricesPriceModelTypeScalableMatrixWithUnitPricing, PlanNewPlanVersionParamsReplacePricesPriceModelTypeScalableMatrixWithTieredPricing, PlanNewPlanVersionParamsReplacePricesPriceModelTypeCumulativeGroupedBulk, PlanNewPlanVersionParamsReplacePricesPriceModelTypeCumulativeGroupedAllocation, PlanNewPlanVersionParamsReplacePricesPriceModelTypeDailyCreditAllowance, PlanNewPlanVersionParamsReplacePricesPriceModelTypeMeteredAllowance, PlanNewPlanVersionParamsReplacePricesPriceModelTypeMinimumComposite, PlanNewPlanVersionParamsReplacePricesPriceModelTypePercent, PlanNewPlanVersionParamsReplacePricesPriceModelTypeEventOutput, PlanNewPlanVersionParamsReplacePricesPriceModelTypeUnitWithMultiplier:
		return true
	}
	return false
}

type PlanSetDefaultPlanVersionParams struct {
	// Plan version to set as the default.
	Version param.Field[int64] `json:"version" api:"required"`
}

func (r PlanSetDefaultPlanVersionParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}
