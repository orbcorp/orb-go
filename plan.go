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
// Other fields on a customer are currently immutable.
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
// [PlanAdjustmentsAmountDiscountAdjustment],
// [PlanAdjustmentsPercentageDiscountAdjustment],
// [PlanAdjustmentsUsageDiscountAdjustment], [PlanAdjustmentsMinimumAdjustment],
// [PlanAdjustmentsMaximumAdjustment].
func (r PlanAdjustment) AsUnion() PlanAdjustmentsUnion {
	return r.union
}

// Union satisfied by [PlanAdjustmentsAmountDiscountAdjustment],
// [PlanAdjustmentsPercentageDiscountAdjustment],
// [PlanAdjustmentsUsageDiscountAdjustment], [PlanAdjustmentsMinimumAdjustment] or
// [PlanAdjustmentsMaximumAdjustment].
type PlanAdjustmentsUnion interface {
	implementsPlanAdjustment()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*PlanAdjustmentsUnion)(nil)).Elem(),
		"adjustment_type",
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(PlanAdjustmentsAmountDiscountAdjustment{}),
			DiscriminatorValue: "amount_discount",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(PlanAdjustmentsPercentageDiscountAdjustment{}),
			DiscriminatorValue: "percentage_discount",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(PlanAdjustmentsUsageDiscountAdjustment{}),
			DiscriminatorValue: "usage_discount",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(PlanAdjustmentsMinimumAdjustment{}),
			DiscriminatorValue: "minimum",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(PlanAdjustmentsMaximumAdjustment{}),
			DiscriminatorValue: "maximum",
		},
	)
}

type PlanAdjustmentsAmountDiscountAdjustment struct {
	ID             string                                                `json:"id,required"`
	AdjustmentType PlanAdjustmentsAmountDiscountAdjustmentAdjustmentType `json:"adjustment_type,required"`
	// The amount by which to discount the prices this adjustment applies to in a given
	// billing period.
	AmountDiscount string `json:"amount_discount,required"`
	// The price IDs that this adjustment applies to.
	AppliesToPriceIDs []string `json:"applies_to_price_ids,required"`
	// True for adjustments that apply to an entire invocice, false for adjustments
	// that apply to only one price.
	IsInvoiceLevel bool `json:"is_invoice_level,required"`
	// The plan phase in which this adjustment is active.
	PlanPhaseOrder int64 `json:"plan_phase_order,required,nullable"`
	// The reason for the adjustment.
	Reason string                                      `json:"reason,required,nullable"`
	JSON   planAdjustmentsAmountDiscountAdjustmentJSON `json:"-"`
}

// planAdjustmentsAmountDiscountAdjustmentJSON contains the JSON metadata for the
// struct [PlanAdjustmentsAmountDiscountAdjustment]
type planAdjustmentsAmountDiscountAdjustmentJSON struct {
	ID                apijson.Field
	AdjustmentType    apijson.Field
	AmountDiscount    apijson.Field
	AppliesToPriceIDs apijson.Field
	IsInvoiceLevel    apijson.Field
	PlanPhaseOrder    apijson.Field
	Reason            apijson.Field
	raw               string
	ExtraFields       map[string]apijson.Field
}

func (r *PlanAdjustmentsAmountDiscountAdjustment) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r planAdjustmentsAmountDiscountAdjustmentJSON) RawJSON() string {
	return r.raw
}

func (r PlanAdjustmentsAmountDiscountAdjustment) implementsPlanAdjustment() {}

type PlanAdjustmentsAmountDiscountAdjustmentAdjustmentType string

const (
	PlanAdjustmentsAmountDiscountAdjustmentAdjustmentTypeAmountDiscount PlanAdjustmentsAmountDiscountAdjustmentAdjustmentType = "amount_discount"
)

func (r PlanAdjustmentsAmountDiscountAdjustmentAdjustmentType) IsKnown() bool {
	switch r {
	case PlanAdjustmentsAmountDiscountAdjustmentAdjustmentTypeAmountDiscount:
		return true
	}
	return false
}

type PlanAdjustmentsPercentageDiscountAdjustment struct {
	ID             string                                                    `json:"id,required"`
	AdjustmentType PlanAdjustmentsPercentageDiscountAdjustmentAdjustmentType `json:"adjustment_type,required"`
	// The price IDs that this adjustment applies to.
	AppliesToPriceIDs []string `json:"applies_to_price_ids,required"`
	// True for adjustments that apply to an entire invocice, false for adjustments
	// that apply to only one price.
	IsInvoiceLevel bool `json:"is_invoice_level,required"`
	// The percentage (as a value between 0 and 1) by which to discount the price
	// intervals this adjustment applies to in a given billing period.
	PercentageDiscount float64 `json:"percentage_discount,required"`
	// The plan phase in which this adjustment is active.
	PlanPhaseOrder int64 `json:"plan_phase_order,required,nullable"`
	// The reason for the adjustment.
	Reason string                                          `json:"reason,required,nullable"`
	JSON   planAdjustmentsPercentageDiscountAdjustmentJSON `json:"-"`
}

// planAdjustmentsPercentageDiscountAdjustmentJSON contains the JSON metadata for
// the struct [PlanAdjustmentsPercentageDiscountAdjustment]
type planAdjustmentsPercentageDiscountAdjustmentJSON struct {
	ID                 apijson.Field
	AdjustmentType     apijson.Field
	AppliesToPriceIDs  apijson.Field
	IsInvoiceLevel     apijson.Field
	PercentageDiscount apijson.Field
	PlanPhaseOrder     apijson.Field
	Reason             apijson.Field
	raw                string
	ExtraFields        map[string]apijson.Field
}

func (r *PlanAdjustmentsPercentageDiscountAdjustment) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r planAdjustmentsPercentageDiscountAdjustmentJSON) RawJSON() string {
	return r.raw
}

func (r PlanAdjustmentsPercentageDiscountAdjustment) implementsPlanAdjustment() {}

type PlanAdjustmentsPercentageDiscountAdjustmentAdjustmentType string

const (
	PlanAdjustmentsPercentageDiscountAdjustmentAdjustmentTypePercentageDiscount PlanAdjustmentsPercentageDiscountAdjustmentAdjustmentType = "percentage_discount"
)

func (r PlanAdjustmentsPercentageDiscountAdjustmentAdjustmentType) IsKnown() bool {
	switch r {
	case PlanAdjustmentsPercentageDiscountAdjustmentAdjustmentTypePercentageDiscount:
		return true
	}
	return false
}

type PlanAdjustmentsUsageDiscountAdjustment struct {
	ID             string                                               `json:"id,required"`
	AdjustmentType PlanAdjustmentsUsageDiscountAdjustmentAdjustmentType `json:"adjustment_type,required"`
	// The price IDs that this adjustment applies to.
	AppliesToPriceIDs []string `json:"applies_to_price_ids,required"`
	// True for adjustments that apply to an entire invocice, false for adjustments
	// that apply to only one price.
	IsInvoiceLevel bool `json:"is_invoice_level,required"`
	// The plan phase in which this adjustment is active.
	PlanPhaseOrder int64 `json:"plan_phase_order,required,nullable"`
	// The reason for the adjustment.
	Reason string `json:"reason,required,nullable"`
	// The number of usage units by which to discount the price this adjustment applies
	// to in a given billing period.
	UsageDiscount float64                                    `json:"usage_discount,required"`
	JSON          planAdjustmentsUsageDiscountAdjustmentJSON `json:"-"`
}

// planAdjustmentsUsageDiscountAdjustmentJSON contains the JSON metadata for the
// struct [PlanAdjustmentsUsageDiscountAdjustment]
type planAdjustmentsUsageDiscountAdjustmentJSON struct {
	ID                apijson.Field
	AdjustmentType    apijson.Field
	AppliesToPriceIDs apijson.Field
	IsInvoiceLevel    apijson.Field
	PlanPhaseOrder    apijson.Field
	Reason            apijson.Field
	UsageDiscount     apijson.Field
	raw               string
	ExtraFields       map[string]apijson.Field
}

func (r *PlanAdjustmentsUsageDiscountAdjustment) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r planAdjustmentsUsageDiscountAdjustmentJSON) RawJSON() string {
	return r.raw
}

func (r PlanAdjustmentsUsageDiscountAdjustment) implementsPlanAdjustment() {}

type PlanAdjustmentsUsageDiscountAdjustmentAdjustmentType string

const (
	PlanAdjustmentsUsageDiscountAdjustmentAdjustmentTypeUsageDiscount PlanAdjustmentsUsageDiscountAdjustmentAdjustmentType = "usage_discount"
)

func (r PlanAdjustmentsUsageDiscountAdjustmentAdjustmentType) IsKnown() bool {
	switch r {
	case PlanAdjustmentsUsageDiscountAdjustmentAdjustmentTypeUsageDiscount:
		return true
	}
	return false
}

type PlanAdjustmentsMinimumAdjustment struct {
	ID             string                                         `json:"id,required"`
	AdjustmentType PlanAdjustmentsMinimumAdjustmentAdjustmentType `json:"adjustment_type,required"`
	// The price IDs that this adjustment applies to.
	AppliesToPriceIDs []string `json:"applies_to_price_ids,required"`
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
	Reason string                               `json:"reason,required,nullable"`
	JSON   planAdjustmentsMinimumAdjustmentJSON `json:"-"`
}

// planAdjustmentsMinimumAdjustmentJSON contains the JSON metadata for the struct
// [PlanAdjustmentsMinimumAdjustment]
type planAdjustmentsMinimumAdjustmentJSON struct {
	ID                apijson.Field
	AdjustmentType    apijson.Field
	AppliesToPriceIDs apijson.Field
	IsInvoiceLevel    apijson.Field
	ItemID            apijson.Field
	MinimumAmount     apijson.Field
	PlanPhaseOrder    apijson.Field
	Reason            apijson.Field
	raw               string
	ExtraFields       map[string]apijson.Field
}

func (r *PlanAdjustmentsMinimumAdjustment) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r planAdjustmentsMinimumAdjustmentJSON) RawJSON() string {
	return r.raw
}

func (r PlanAdjustmentsMinimumAdjustment) implementsPlanAdjustment() {}

type PlanAdjustmentsMinimumAdjustmentAdjustmentType string

const (
	PlanAdjustmentsMinimumAdjustmentAdjustmentTypeMinimum PlanAdjustmentsMinimumAdjustmentAdjustmentType = "minimum"
)

func (r PlanAdjustmentsMinimumAdjustmentAdjustmentType) IsKnown() bool {
	switch r {
	case PlanAdjustmentsMinimumAdjustmentAdjustmentTypeMinimum:
		return true
	}
	return false
}

type PlanAdjustmentsMaximumAdjustment struct {
	ID             string                                         `json:"id,required"`
	AdjustmentType PlanAdjustmentsMaximumAdjustmentAdjustmentType `json:"adjustment_type,required"`
	// The price IDs that this adjustment applies to.
	AppliesToPriceIDs []string `json:"applies_to_price_ids,required"`
	// True for adjustments that apply to an entire invocice, false for adjustments
	// that apply to only one price.
	IsInvoiceLevel bool `json:"is_invoice_level,required"`
	// The maximum amount to charge in a given billing period for the prices this
	// adjustment applies to.
	MaximumAmount string `json:"maximum_amount,required"`
	// The plan phase in which this adjustment is active.
	PlanPhaseOrder int64 `json:"plan_phase_order,required,nullable"`
	// The reason for the adjustment.
	Reason string                               `json:"reason,required,nullable"`
	JSON   planAdjustmentsMaximumAdjustmentJSON `json:"-"`
}

// planAdjustmentsMaximumAdjustmentJSON contains the JSON metadata for the struct
// [PlanAdjustmentsMaximumAdjustment]
type planAdjustmentsMaximumAdjustmentJSON struct {
	ID                apijson.Field
	AdjustmentType    apijson.Field
	AppliesToPriceIDs apijson.Field
	IsInvoiceLevel    apijson.Field
	MaximumAmount     apijson.Field
	PlanPhaseOrder    apijson.Field
	Reason            apijson.Field
	raw               string
	ExtraFields       map[string]apijson.Field
}

func (r *PlanAdjustmentsMaximumAdjustment) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r planAdjustmentsMaximumAdjustmentJSON) RawJSON() string {
	return r.raw
}

func (r PlanAdjustmentsMaximumAdjustment) implementsPlanAdjustment() {}

type PlanAdjustmentsMaximumAdjustmentAdjustmentType string

const (
	PlanAdjustmentsMaximumAdjustmentAdjustmentTypeMaximum PlanAdjustmentsMaximumAdjustmentAdjustmentType = "maximum"
)

func (r PlanAdjustmentsMaximumAdjustmentAdjustmentType) IsKnown() bool {
	switch r {
	case PlanAdjustmentsMaximumAdjustmentAdjustmentTypeMaximum:
		return true
	}
	return false
}

type PlanAdjustmentsAdjustmentType string

const (
	PlanAdjustmentsAdjustmentTypeAmountDiscount     PlanAdjustmentsAdjustmentType = "amount_discount"
	PlanAdjustmentsAdjustmentTypePercentageDiscount PlanAdjustmentsAdjustmentType = "percentage_discount"
	PlanAdjustmentsAdjustmentTypeUsageDiscount      PlanAdjustmentsAdjustmentType = "usage_discount"
	PlanAdjustmentsAdjustmentTypeMinimum            PlanAdjustmentsAdjustmentType = "minimum"
	PlanAdjustmentsAdjustmentTypeMaximum            PlanAdjustmentsAdjustmentType = "maximum"
)

func (r PlanAdjustmentsAdjustmentType) IsKnown() bool {
	switch r {
	case PlanAdjustmentsAdjustmentTypeAmountDiscount, PlanAdjustmentsAdjustmentTypePercentageDiscount, PlanAdjustmentsAdjustmentTypeUsageDiscount, PlanAdjustmentsAdjustmentTypeMinimum, PlanAdjustmentsAdjustmentTypeMaximum:
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
	// The id of the item the plan will be associated with.
	ItemID    param.Field[string]                       `json:"item_id,required"`
	ModelType param.Field[PlanNewParamsPricesModelType] `json:"model_type,required"`
	// The name of the price.
	Name param.Field[string] `json:"name,required"`
	// The id of the billable metric for the price. Only needed if the price is
	// usage-based.
	BillableMetricID param.Field[string] `json:"billable_metric_id"`
	// If the Price represents a fixed cost, the price will be billed in-advance if
	// this is true, and in-arrears if this is false.
	BilledInAdvance           param.Field[bool]        `json:"billed_in_advance"`
	BillingCycleConfiguration param.Field[interface{}] `json:"billing_cycle_configuration"`
	BpsConfig                 param.Field[interface{}] `json:"bps_config"`
	BulkBpsConfig             param.Field[interface{}] `json:"bulk_bps_config"`
	BulkConfig                param.Field[interface{}] `json:"bulk_config"`
	BulkWithProrationConfig   param.Field[interface{}] `json:"bulk_with_proration_config"`
	// The per unit conversion rate of the price currency to the invoicing currency.
	ConversionRate param.Field[float64] `json:"conversion_rate"`
	// An ISO 4217 currency string, or custom pricing unit identifier, in which this
	// price is billed.
	Currency param.Field[string] `json:"currency"`
	// An alias for the price.
	ExternalPriceID param.Field[string] `json:"external_price_id"`
	// If the Price represents a fixed cost, this represents the quantity of units
	// applied.
	FixedPriceQuantity               param.Field[float64]     `json:"fixed_price_quantity"`
	GroupedAllocationConfig          param.Field[interface{}] `json:"grouped_allocation_config"`
	GroupedTieredPackageConfig       param.Field[interface{}] `json:"grouped_tiered_package_config"`
	GroupedWithMeteredMinimumConfig  param.Field[interface{}] `json:"grouped_with_metered_minimum_config"`
	GroupedWithProratedMinimumConfig param.Field[interface{}] `json:"grouped_with_prorated_minimum_config"`
	// The property used to group this price on an invoice
	InvoiceGroupingKey          param.Field[string]      `json:"invoice_grouping_key"`
	InvoicingCycleConfiguration param.Field[interface{}] `json:"invoicing_cycle_configuration"`
	MatrixConfig                param.Field[interface{}] `json:"matrix_config"`
	MatrixWithDisplayNameConfig param.Field[interface{}] `json:"matrix_with_display_name_config"`
	MaxGroupTieredPackageConfig param.Field[interface{}] `json:"max_group_tiered_package_config"`
	Metadata                    param.Field[interface{}] `json:"metadata"`
	PackageConfig               param.Field[interface{}] `json:"package_config"`
	PackageWithAllocationConfig param.Field[interface{}] `json:"package_with_allocation_config"`
	ThresholdTotalAmountConfig  param.Field[interface{}] `json:"threshold_total_amount_config"`
	TieredBpsConfig             param.Field[interface{}] `json:"tiered_bps_config"`
	TieredConfig                param.Field[interface{}] `json:"tiered_config"`
	TieredPackageConfig         param.Field[interface{}] `json:"tiered_package_config"`
	TieredWithMinimumConfig     param.Field[interface{}] `json:"tiered_with_minimum_config"`
	TieredWithProrationConfig   param.Field[interface{}] `json:"tiered_with_proration_config"`
	UnitConfig                  param.Field[interface{}] `json:"unit_config"`
	UnitWithPercentConfig       param.Field[interface{}] `json:"unit_with_percent_config"`
	UnitWithProrationConfig     param.Field[interface{}] `json:"unit_with_proration_config"`
}

func (r PlanNewParamsPrice) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r PlanNewParamsPrice) implementsPlanNewParamsPriceUnion() {}

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
// [PlanNewParamsPricesNewPlanPackageWithAllocationPrice],
// [PlanNewParamsPricesNewPlanTierWithProrationPrice],
// [PlanNewParamsPricesNewPlanUnitWithProrationPrice],
// [PlanNewParamsPricesNewPlanGroupedAllocationPrice],
// [PlanNewParamsPricesNewPlanGroupedWithProratedMinimumPrice],
// [PlanNewParamsPricesNewPlanGroupedWithMeteredMinimumPrice],
// [PlanNewParamsPricesNewPlanMatrixWithDisplayNamePrice],
// [PlanNewParamsPricesNewPlanBulkWithProrationPrice],
// [PlanNewParamsPricesNewPlanGroupedTieredPackagePrice],
// [PlanNewParamsPricesNewPlanMaxGroupTieredPackagePrice], [PlanNewParamsPrice].
type PlanNewParamsPriceUnion interface {
	implementsPlanNewParamsPriceUnion()
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
	// For custom cadence: specifies the duration of the billing period in days or
	// months.
	BillingCycleConfiguration param.Field[PlanNewParamsPricesNewPlanUnitPriceBillingCycleConfiguration] `json:"billing_cycle_configuration"`
	// The per unit conversion rate of the price currency to the invoicing currency.
	ConversionRate param.Field[float64] `json:"conversion_rate"`
	// An ISO 4217 currency string, or custom pricing unit identifier, in which this
	// price is billed.
	Currency param.Field[string] `json:"currency"`
	// An alias for the price.
	ExternalPriceID param.Field[string] `json:"external_price_id"`
	// If the Price represents a fixed cost, this represents the quantity of units
	// applied.
	FixedPriceQuantity param.Field[float64] `json:"fixed_price_quantity"`
	// The property used to group this price on an invoice
	InvoiceGroupingKey param.Field[string] `json:"invoice_grouping_key"`
	// Within each billing cycle, specifies the cadence at which invoices are produced.
	// If unspecified, a single invoice is produced per billing cycle.
	InvoicingCycleConfiguration param.Field[PlanNewParamsPricesNewPlanUnitPriceInvoicingCycleConfiguration] `json:"invoicing_cycle_configuration"`
	// User-specified key/value pairs for the resource. Individual keys can be removed
	// by setting the value to `null`, and the entire metadata mapping can be cleared
	// by setting `metadata` to `null`.
	Metadata param.Field[map[string]string] `json:"metadata"`
}

func (r PlanNewParamsPricesNewPlanUnitPrice) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r PlanNewParamsPricesNewPlanUnitPrice) implementsPlanNewParamsPriceUnion() {}

// The cadence to bill for this price on.
type PlanNewParamsPricesNewPlanUnitPriceCadence string

const (
	PlanNewParamsPricesNewPlanUnitPriceCadenceAnnual     PlanNewParamsPricesNewPlanUnitPriceCadence = "annual"
	PlanNewParamsPricesNewPlanUnitPriceCadenceSemiAnnual PlanNewParamsPricesNewPlanUnitPriceCadence = "semi_annual"
	PlanNewParamsPricesNewPlanUnitPriceCadenceMonthly    PlanNewParamsPricesNewPlanUnitPriceCadence = "monthly"
	PlanNewParamsPricesNewPlanUnitPriceCadenceQuarterly  PlanNewParamsPricesNewPlanUnitPriceCadence = "quarterly"
	PlanNewParamsPricesNewPlanUnitPriceCadenceOneTime    PlanNewParamsPricesNewPlanUnitPriceCadence = "one_time"
	PlanNewParamsPricesNewPlanUnitPriceCadenceCustom     PlanNewParamsPricesNewPlanUnitPriceCadence = "custom"
)

func (r PlanNewParamsPricesNewPlanUnitPriceCadence) IsKnown() bool {
	switch r {
	case PlanNewParamsPricesNewPlanUnitPriceCadenceAnnual, PlanNewParamsPricesNewPlanUnitPriceCadenceSemiAnnual, PlanNewParamsPricesNewPlanUnitPriceCadenceMonthly, PlanNewParamsPricesNewPlanUnitPriceCadenceQuarterly, PlanNewParamsPricesNewPlanUnitPriceCadenceOneTime, PlanNewParamsPricesNewPlanUnitPriceCadenceCustom:
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
}

func (r PlanNewParamsPricesNewPlanUnitPriceUnitConfig) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// For custom cadence: specifies the duration of the billing period in days or
// months.
type PlanNewParamsPricesNewPlanUnitPriceBillingCycleConfiguration struct {
	// The duration of the billing period.
	Duration param.Field[int64] `json:"duration,required"`
	// The unit of billing period duration.
	DurationUnit param.Field[PlanNewParamsPricesNewPlanUnitPriceBillingCycleConfigurationDurationUnit] `json:"duration_unit,required"`
}

func (r PlanNewParamsPricesNewPlanUnitPriceBillingCycleConfiguration) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// The unit of billing period duration.
type PlanNewParamsPricesNewPlanUnitPriceBillingCycleConfigurationDurationUnit string

const (
	PlanNewParamsPricesNewPlanUnitPriceBillingCycleConfigurationDurationUnitDay   PlanNewParamsPricesNewPlanUnitPriceBillingCycleConfigurationDurationUnit = "day"
	PlanNewParamsPricesNewPlanUnitPriceBillingCycleConfigurationDurationUnitMonth PlanNewParamsPricesNewPlanUnitPriceBillingCycleConfigurationDurationUnit = "month"
)

func (r PlanNewParamsPricesNewPlanUnitPriceBillingCycleConfigurationDurationUnit) IsKnown() bool {
	switch r {
	case PlanNewParamsPricesNewPlanUnitPriceBillingCycleConfigurationDurationUnitDay, PlanNewParamsPricesNewPlanUnitPriceBillingCycleConfigurationDurationUnitMonth:
		return true
	}
	return false
}

// Within each billing cycle, specifies the cadence at which invoices are produced.
// If unspecified, a single invoice is produced per billing cycle.
type PlanNewParamsPricesNewPlanUnitPriceInvoicingCycleConfiguration struct {
	// The duration of the billing period.
	Duration param.Field[int64] `json:"duration,required"`
	// The unit of billing period duration.
	DurationUnit param.Field[PlanNewParamsPricesNewPlanUnitPriceInvoicingCycleConfigurationDurationUnit] `json:"duration_unit,required"`
}

func (r PlanNewParamsPricesNewPlanUnitPriceInvoicingCycleConfiguration) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// The unit of billing period duration.
type PlanNewParamsPricesNewPlanUnitPriceInvoicingCycleConfigurationDurationUnit string

const (
	PlanNewParamsPricesNewPlanUnitPriceInvoicingCycleConfigurationDurationUnitDay   PlanNewParamsPricesNewPlanUnitPriceInvoicingCycleConfigurationDurationUnit = "day"
	PlanNewParamsPricesNewPlanUnitPriceInvoicingCycleConfigurationDurationUnitMonth PlanNewParamsPricesNewPlanUnitPriceInvoicingCycleConfigurationDurationUnit = "month"
)

func (r PlanNewParamsPricesNewPlanUnitPriceInvoicingCycleConfigurationDurationUnit) IsKnown() bool {
	switch r {
	case PlanNewParamsPricesNewPlanUnitPriceInvoicingCycleConfigurationDurationUnitDay, PlanNewParamsPricesNewPlanUnitPriceInvoicingCycleConfigurationDurationUnitMonth:
		return true
	}
	return false
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
	// For custom cadence: specifies the duration of the billing period in days or
	// months.
	BillingCycleConfiguration param.Field[PlanNewParamsPricesNewPlanPackagePriceBillingCycleConfiguration] `json:"billing_cycle_configuration"`
	// The per unit conversion rate of the price currency to the invoicing currency.
	ConversionRate param.Field[float64] `json:"conversion_rate"`
	// An ISO 4217 currency string, or custom pricing unit identifier, in which this
	// price is billed.
	Currency param.Field[string] `json:"currency"`
	// An alias for the price.
	ExternalPriceID param.Field[string] `json:"external_price_id"`
	// If the Price represents a fixed cost, this represents the quantity of units
	// applied.
	FixedPriceQuantity param.Field[float64] `json:"fixed_price_quantity"`
	// The property used to group this price on an invoice
	InvoiceGroupingKey param.Field[string] `json:"invoice_grouping_key"`
	// Within each billing cycle, specifies the cadence at which invoices are produced.
	// If unspecified, a single invoice is produced per billing cycle.
	InvoicingCycleConfiguration param.Field[PlanNewParamsPricesNewPlanPackagePriceInvoicingCycleConfiguration] `json:"invoicing_cycle_configuration"`
	// User-specified key/value pairs for the resource. Individual keys can be removed
	// by setting the value to `null`, and the entire metadata mapping can be cleared
	// by setting `metadata` to `null`.
	Metadata param.Field[map[string]string] `json:"metadata"`
}

func (r PlanNewParamsPricesNewPlanPackagePrice) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r PlanNewParamsPricesNewPlanPackagePrice) implementsPlanNewParamsPriceUnion() {}

// The cadence to bill for this price on.
type PlanNewParamsPricesNewPlanPackagePriceCadence string

const (
	PlanNewParamsPricesNewPlanPackagePriceCadenceAnnual     PlanNewParamsPricesNewPlanPackagePriceCadence = "annual"
	PlanNewParamsPricesNewPlanPackagePriceCadenceSemiAnnual PlanNewParamsPricesNewPlanPackagePriceCadence = "semi_annual"
	PlanNewParamsPricesNewPlanPackagePriceCadenceMonthly    PlanNewParamsPricesNewPlanPackagePriceCadence = "monthly"
	PlanNewParamsPricesNewPlanPackagePriceCadenceQuarterly  PlanNewParamsPricesNewPlanPackagePriceCadence = "quarterly"
	PlanNewParamsPricesNewPlanPackagePriceCadenceOneTime    PlanNewParamsPricesNewPlanPackagePriceCadence = "one_time"
	PlanNewParamsPricesNewPlanPackagePriceCadenceCustom     PlanNewParamsPricesNewPlanPackagePriceCadence = "custom"
)

func (r PlanNewParamsPricesNewPlanPackagePriceCadence) IsKnown() bool {
	switch r {
	case PlanNewParamsPricesNewPlanPackagePriceCadenceAnnual, PlanNewParamsPricesNewPlanPackagePriceCadenceSemiAnnual, PlanNewParamsPricesNewPlanPackagePriceCadenceMonthly, PlanNewParamsPricesNewPlanPackagePriceCadenceQuarterly, PlanNewParamsPricesNewPlanPackagePriceCadenceOneTime, PlanNewParamsPricesNewPlanPackagePriceCadenceCustom:
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
	PackageSize param.Field[int64] `json:"package_size,required"`
}

func (r PlanNewParamsPricesNewPlanPackagePricePackageConfig) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// For custom cadence: specifies the duration of the billing period in days or
// months.
type PlanNewParamsPricesNewPlanPackagePriceBillingCycleConfiguration struct {
	// The duration of the billing period.
	Duration param.Field[int64] `json:"duration,required"`
	// The unit of billing period duration.
	DurationUnit param.Field[PlanNewParamsPricesNewPlanPackagePriceBillingCycleConfigurationDurationUnit] `json:"duration_unit,required"`
}

func (r PlanNewParamsPricesNewPlanPackagePriceBillingCycleConfiguration) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// The unit of billing period duration.
type PlanNewParamsPricesNewPlanPackagePriceBillingCycleConfigurationDurationUnit string

const (
	PlanNewParamsPricesNewPlanPackagePriceBillingCycleConfigurationDurationUnitDay   PlanNewParamsPricesNewPlanPackagePriceBillingCycleConfigurationDurationUnit = "day"
	PlanNewParamsPricesNewPlanPackagePriceBillingCycleConfigurationDurationUnitMonth PlanNewParamsPricesNewPlanPackagePriceBillingCycleConfigurationDurationUnit = "month"
)

func (r PlanNewParamsPricesNewPlanPackagePriceBillingCycleConfigurationDurationUnit) IsKnown() bool {
	switch r {
	case PlanNewParamsPricesNewPlanPackagePriceBillingCycleConfigurationDurationUnitDay, PlanNewParamsPricesNewPlanPackagePriceBillingCycleConfigurationDurationUnitMonth:
		return true
	}
	return false
}

// Within each billing cycle, specifies the cadence at which invoices are produced.
// If unspecified, a single invoice is produced per billing cycle.
type PlanNewParamsPricesNewPlanPackagePriceInvoicingCycleConfiguration struct {
	// The duration of the billing period.
	Duration param.Field[int64] `json:"duration,required"`
	// The unit of billing period duration.
	DurationUnit param.Field[PlanNewParamsPricesNewPlanPackagePriceInvoicingCycleConfigurationDurationUnit] `json:"duration_unit,required"`
}

func (r PlanNewParamsPricesNewPlanPackagePriceInvoicingCycleConfiguration) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// The unit of billing period duration.
type PlanNewParamsPricesNewPlanPackagePriceInvoicingCycleConfigurationDurationUnit string

const (
	PlanNewParamsPricesNewPlanPackagePriceInvoicingCycleConfigurationDurationUnitDay   PlanNewParamsPricesNewPlanPackagePriceInvoicingCycleConfigurationDurationUnit = "day"
	PlanNewParamsPricesNewPlanPackagePriceInvoicingCycleConfigurationDurationUnitMonth PlanNewParamsPricesNewPlanPackagePriceInvoicingCycleConfigurationDurationUnit = "month"
)

func (r PlanNewParamsPricesNewPlanPackagePriceInvoicingCycleConfigurationDurationUnit) IsKnown() bool {
	switch r {
	case PlanNewParamsPricesNewPlanPackagePriceInvoicingCycleConfigurationDurationUnitDay, PlanNewParamsPricesNewPlanPackagePriceInvoicingCycleConfigurationDurationUnitMonth:
		return true
	}
	return false
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
	// For custom cadence: specifies the duration of the billing period in days or
	// months.
	BillingCycleConfiguration param.Field[PlanNewParamsPricesNewPlanMatrixPriceBillingCycleConfiguration] `json:"billing_cycle_configuration"`
	// The per unit conversion rate of the price currency to the invoicing currency.
	ConversionRate param.Field[float64] `json:"conversion_rate"`
	// An ISO 4217 currency string, or custom pricing unit identifier, in which this
	// price is billed.
	Currency param.Field[string] `json:"currency"`
	// An alias for the price.
	ExternalPriceID param.Field[string] `json:"external_price_id"`
	// If the Price represents a fixed cost, this represents the quantity of units
	// applied.
	FixedPriceQuantity param.Field[float64] `json:"fixed_price_quantity"`
	// The property used to group this price on an invoice
	InvoiceGroupingKey param.Field[string] `json:"invoice_grouping_key"`
	// Within each billing cycle, specifies the cadence at which invoices are produced.
	// If unspecified, a single invoice is produced per billing cycle.
	InvoicingCycleConfiguration param.Field[PlanNewParamsPricesNewPlanMatrixPriceInvoicingCycleConfiguration] `json:"invoicing_cycle_configuration"`
	// User-specified key/value pairs for the resource. Individual keys can be removed
	// by setting the value to `null`, and the entire metadata mapping can be cleared
	// by setting `metadata` to `null`.
	Metadata param.Field[map[string]string] `json:"metadata"`
}

func (r PlanNewParamsPricesNewPlanMatrixPrice) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r PlanNewParamsPricesNewPlanMatrixPrice) implementsPlanNewParamsPriceUnion() {}

// The cadence to bill for this price on.
type PlanNewParamsPricesNewPlanMatrixPriceCadence string

const (
	PlanNewParamsPricesNewPlanMatrixPriceCadenceAnnual     PlanNewParamsPricesNewPlanMatrixPriceCadence = "annual"
	PlanNewParamsPricesNewPlanMatrixPriceCadenceSemiAnnual PlanNewParamsPricesNewPlanMatrixPriceCadence = "semi_annual"
	PlanNewParamsPricesNewPlanMatrixPriceCadenceMonthly    PlanNewParamsPricesNewPlanMatrixPriceCadence = "monthly"
	PlanNewParamsPricesNewPlanMatrixPriceCadenceQuarterly  PlanNewParamsPricesNewPlanMatrixPriceCadence = "quarterly"
	PlanNewParamsPricesNewPlanMatrixPriceCadenceOneTime    PlanNewParamsPricesNewPlanMatrixPriceCadence = "one_time"
	PlanNewParamsPricesNewPlanMatrixPriceCadenceCustom     PlanNewParamsPricesNewPlanMatrixPriceCadence = "custom"
)

func (r PlanNewParamsPricesNewPlanMatrixPriceCadence) IsKnown() bool {
	switch r {
	case PlanNewParamsPricesNewPlanMatrixPriceCadenceAnnual, PlanNewParamsPricesNewPlanMatrixPriceCadenceSemiAnnual, PlanNewParamsPricesNewPlanMatrixPriceCadenceMonthly, PlanNewParamsPricesNewPlanMatrixPriceCadenceQuarterly, PlanNewParamsPricesNewPlanMatrixPriceCadenceOneTime, PlanNewParamsPricesNewPlanMatrixPriceCadenceCustom:
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

// For custom cadence: specifies the duration of the billing period in days or
// months.
type PlanNewParamsPricesNewPlanMatrixPriceBillingCycleConfiguration struct {
	// The duration of the billing period.
	Duration param.Field[int64] `json:"duration,required"`
	// The unit of billing period duration.
	DurationUnit param.Field[PlanNewParamsPricesNewPlanMatrixPriceBillingCycleConfigurationDurationUnit] `json:"duration_unit,required"`
}

func (r PlanNewParamsPricesNewPlanMatrixPriceBillingCycleConfiguration) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// The unit of billing period duration.
type PlanNewParamsPricesNewPlanMatrixPriceBillingCycleConfigurationDurationUnit string

const (
	PlanNewParamsPricesNewPlanMatrixPriceBillingCycleConfigurationDurationUnitDay   PlanNewParamsPricesNewPlanMatrixPriceBillingCycleConfigurationDurationUnit = "day"
	PlanNewParamsPricesNewPlanMatrixPriceBillingCycleConfigurationDurationUnitMonth PlanNewParamsPricesNewPlanMatrixPriceBillingCycleConfigurationDurationUnit = "month"
)

func (r PlanNewParamsPricesNewPlanMatrixPriceBillingCycleConfigurationDurationUnit) IsKnown() bool {
	switch r {
	case PlanNewParamsPricesNewPlanMatrixPriceBillingCycleConfigurationDurationUnitDay, PlanNewParamsPricesNewPlanMatrixPriceBillingCycleConfigurationDurationUnitMonth:
		return true
	}
	return false
}

// Within each billing cycle, specifies the cadence at which invoices are produced.
// If unspecified, a single invoice is produced per billing cycle.
type PlanNewParamsPricesNewPlanMatrixPriceInvoicingCycleConfiguration struct {
	// The duration of the billing period.
	Duration param.Field[int64] `json:"duration,required"`
	// The unit of billing period duration.
	DurationUnit param.Field[PlanNewParamsPricesNewPlanMatrixPriceInvoicingCycleConfigurationDurationUnit] `json:"duration_unit,required"`
}

func (r PlanNewParamsPricesNewPlanMatrixPriceInvoicingCycleConfiguration) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// The unit of billing period duration.
type PlanNewParamsPricesNewPlanMatrixPriceInvoicingCycleConfigurationDurationUnit string

const (
	PlanNewParamsPricesNewPlanMatrixPriceInvoicingCycleConfigurationDurationUnitDay   PlanNewParamsPricesNewPlanMatrixPriceInvoicingCycleConfigurationDurationUnit = "day"
	PlanNewParamsPricesNewPlanMatrixPriceInvoicingCycleConfigurationDurationUnitMonth PlanNewParamsPricesNewPlanMatrixPriceInvoicingCycleConfigurationDurationUnit = "month"
)

func (r PlanNewParamsPricesNewPlanMatrixPriceInvoicingCycleConfigurationDurationUnit) IsKnown() bool {
	switch r {
	case PlanNewParamsPricesNewPlanMatrixPriceInvoicingCycleConfigurationDurationUnitDay, PlanNewParamsPricesNewPlanMatrixPriceInvoicingCycleConfigurationDurationUnitMonth:
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
	// For custom cadence: specifies the duration of the billing period in days or
	// months.
	BillingCycleConfiguration param.Field[PlanNewParamsPricesNewPlanTieredPriceBillingCycleConfiguration] `json:"billing_cycle_configuration"`
	// The per unit conversion rate of the price currency to the invoicing currency.
	ConversionRate param.Field[float64] `json:"conversion_rate"`
	// An ISO 4217 currency string, or custom pricing unit identifier, in which this
	// price is billed.
	Currency param.Field[string] `json:"currency"`
	// An alias for the price.
	ExternalPriceID param.Field[string] `json:"external_price_id"`
	// If the Price represents a fixed cost, this represents the quantity of units
	// applied.
	FixedPriceQuantity param.Field[float64] `json:"fixed_price_quantity"`
	// The property used to group this price on an invoice
	InvoiceGroupingKey param.Field[string] `json:"invoice_grouping_key"`
	// Within each billing cycle, specifies the cadence at which invoices are produced.
	// If unspecified, a single invoice is produced per billing cycle.
	InvoicingCycleConfiguration param.Field[PlanNewParamsPricesNewPlanTieredPriceInvoicingCycleConfiguration] `json:"invoicing_cycle_configuration"`
	// User-specified key/value pairs for the resource. Individual keys can be removed
	// by setting the value to `null`, and the entire metadata mapping can be cleared
	// by setting `metadata` to `null`.
	Metadata param.Field[map[string]string] `json:"metadata"`
}

func (r PlanNewParamsPricesNewPlanTieredPrice) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r PlanNewParamsPricesNewPlanTieredPrice) implementsPlanNewParamsPriceUnion() {}

// The cadence to bill for this price on.
type PlanNewParamsPricesNewPlanTieredPriceCadence string

const (
	PlanNewParamsPricesNewPlanTieredPriceCadenceAnnual     PlanNewParamsPricesNewPlanTieredPriceCadence = "annual"
	PlanNewParamsPricesNewPlanTieredPriceCadenceSemiAnnual PlanNewParamsPricesNewPlanTieredPriceCadence = "semi_annual"
	PlanNewParamsPricesNewPlanTieredPriceCadenceMonthly    PlanNewParamsPricesNewPlanTieredPriceCadence = "monthly"
	PlanNewParamsPricesNewPlanTieredPriceCadenceQuarterly  PlanNewParamsPricesNewPlanTieredPriceCadence = "quarterly"
	PlanNewParamsPricesNewPlanTieredPriceCadenceOneTime    PlanNewParamsPricesNewPlanTieredPriceCadence = "one_time"
	PlanNewParamsPricesNewPlanTieredPriceCadenceCustom     PlanNewParamsPricesNewPlanTieredPriceCadence = "custom"
)

func (r PlanNewParamsPricesNewPlanTieredPriceCadence) IsKnown() bool {
	switch r {
	case PlanNewParamsPricesNewPlanTieredPriceCadenceAnnual, PlanNewParamsPricesNewPlanTieredPriceCadenceSemiAnnual, PlanNewParamsPricesNewPlanTieredPriceCadenceMonthly, PlanNewParamsPricesNewPlanTieredPriceCadenceQuarterly, PlanNewParamsPricesNewPlanTieredPriceCadenceOneTime, PlanNewParamsPricesNewPlanTieredPriceCadenceCustom:
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

// For custom cadence: specifies the duration of the billing period in days or
// months.
type PlanNewParamsPricesNewPlanTieredPriceBillingCycleConfiguration struct {
	// The duration of the billing period.
	Duration param.Field[int64] `json:"duration,required"`
	// The unit of billing period duration.
	DurationUnit param.Field[PlanNewParamsPricesNewPlanTieredPriceBillingCycleConfigurationDurationUnit] `json:"duration_unit,required"`
}

func (r PlanNewParamsPricesNewPlanTieredPriceBillingCycleConfiguration) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// The unit of billing period duration.
type PlanNewParamsPricesNewPlanTieredPriceBillingCycleConfigurationDurationUnit string

const (
	PlanNewParamsPricesNewPlanTieredPriceBillingCycleConfigurationDurationUnitDay   PlanNewParamsPricesNewPlanTieredPriceBillingCycleConfigurationDurationUnit = "day"
	PlanNewParamsPricesNewPlanTieredPriceBillingCycleConfigurationDurationUnitMonth PlanNewParamsPricesNewPlanTieredPriceBillingCycleConfigurationDurationUnit = "month"
)

func (r PlanNewParamsPricesNewPlanTieredPriceBillingCycleConfigurationDurationUnit) IsKnown() bool {
	switch r {
	case PlanNewParamsPricesNewPlanTieredPriceBillingCycleConfigurationDurationUnitDay, PlanNewParamsPricesNewPlanTieredPriceBillingCycleConfigurationDurationUnitMonth:
		return true
	}
	return false
}

// Within each billing cycle, specifies the cadence at which invoices are produced.
// If unspecified, a single invoice is produced per billing cycle.
type PlanNewParamsPricesNewPlanTieredPriceInvoicingCycleConfiguration struct {
	// The duration of the billing period.
	Duration param.Field[int64] `json:"duration,required"`
	// The unit of billing period duration.
	DurationUnit param.Field[PlanNewParamsPricesNewPlanTieredPriceInvoicingCycleConfigurationDurationUnit] `json:"duration_unit,required"`
}

func (r PlanNewParamsPricesNewPlanTieredPriceInvoicingCycleConfiguration) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// The unit of billing period duration.
type PlanNewParamsPricesNewPlanTieredPriceInvoicingCycleConfigurationDurationUnit string

const (
	PlanNewParamsPricesNewPlanTieredPriceInvoicingCycleConfigurationDurationUnitDay   PlanNewParamsPricesNewPlanTieredPriceInvoicingCycleConfigurationDurationUnit = "day"
	PlanNewParamsPricesNewPlanTieredPriceInvoicingCycleConfigurationDurationUnitMonth PlanNewParamsPricesNewPlanTieredPriceInvoicingCycleConfigurationDurationUnit = "month"
)

func (r PlanNewParamsPricesNewPlanTieredPriceInvoicingCycleConfigurationDurationUnit) IsKnown() bool {
	switch r {
	case PlanNewParamsPricesNewPlanTieredPriceInvoicingCycleConfigurationDurationUnitDay, PlanNewParamsPricesNewPlanTieredPriceInvoicingCycleConfigurationDurationUnitMonth:
		return true
	}
	return false
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
	// For custom cadence: specifies the duration of the billing period in days or
	// months.
	BillingCycleConfiguration param.Field[PlanNewParamsPricesNewPlanTieredBpsPriceBillingCycleConfiguration] `json:"billing_cycle_configuration"`
	// The per unit conversion rate of the price currency to the invoicing currency.
	ConversionRate param.Field[float64] `json:"conversion_rate"`
	// An ISO 4217 currency string, or custom pricing unit identifier, in which this
	// price is billed.
	Currency param.Field[string] `json:"currency"`
	// An alias for the price.
	ExternalPriceID param.Field[string] `json:"external_price_id"`
	// If the Price represents a fixed cost, this represents the quantity of units
	// applied.
	FixedPriceQuantity param.Field[float64] `json:"fixed_price_quantity"`
	// The property used to group this price on an invoice
	InvoiceGroupingKey param.Field[string] `json:"invoice_grouping_key"`
	// Within each billing cycle, specifies the cadence at which invoices are produced.
	// If unspecified, a single invoice is produced per billing cycle.
	InvoicingCycleConfiguration param.Field[PlanNewParamsPricesNewPlanTieredBpsPriceInvoicingCycleConfiguration] `json:"invoicing_cycle_configuration"`
	// User-specified key/value pairs for the resource. Individual keys can be removed
	// by setting the value to `null`, and the entire metadata mapping can be cleared
	// by setting `metadata` to `null`.
	Metadata param.Field[map[string]string] `json:"metadata"`
}

func (r PlanNewParamsPricesNewPlanTieredBpsPrice) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r PlanNewParamsPricesNewPlanTieredBpsPrice) implementsPlanNewParamsPriceUnion() {}

// The cadence to bill for this price on.
type PlanNewParamsPricesNewPlanTieredBpsPriceCadence string

const (
	PlanNewParamsPricesNewPlanTieredBpsPriceCadenceAnnual     PlanNewParamsPricesNewPlanTieredBpsPriceCadence = "annual"
	PlanNewParamsPricesNewPlanTieredBpsPriceCadenceSemiAnnual PlanNewParamsPricesNewPlanTieredBpsPriceCadence = "semi_annual"
	PlanNewParamsPricesNewPlanTieredBpsPriceCadenceMonthly    PlanNewParamsPricesNewPlanTieredBpsPriceCadence = "monthly"
	PlanNewParamsPricesNewPlanTieredBpsPriceCadenceQuarterly  PlanNewParamsPricesNewPlanTieredBpsPriceCadence = "quarterly"
	PlanNewParamsPricesNewPlanTieredBpsPriceCadenceOneTime    PlanNewParamsPricesNewPlanTieredBpsPriceCadence = "one_time"
	PlanNewParamsPricesNewPlanTieredBpsPriceCadenceCustom     PlanNewParamsPricesNewPlanTieredBpsPriceCadence = "custom"
)

func (r PlanNewParamsPricesNewPlanTieredBpsPriceCadence) IsKnown() bool {
	switch r {
	case PlanNewParamsPricesNewPlanTieredBpsPriceCadenceAnnual, PlanNewParamsPricesNewPlanTieredBpsPriceCadenceSemiAnnual, PlanNewParamsPricesNewPlanTieredBpsPriceCadenceMonthly, PlanNewParamsPricesNewPlanTieredBpsPriceCadenceQuarterly, PlanNewParamsPricesNewPlanTieredBpsPriceCadenceOneTime, PlanNewParamsPricesNewPlanTieredBpsPriceCadenceCustom:
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

// For custom cadence: specifies the duration of the billing period in days or
// months.
type PlanNewParamsPricesNewPlanTieredBpsPriceBillingCycleConfiguration struct {
	// The duration of the billing period.
	Duration param.Field[int64] `json:"duration,required"`
	// The unit of billing period duration.
	DurationUnit param.Field[PlanNewParamsPricesNewPlanTieredBpsPriceBillingCycleConfigurationDurationUnit] `json:"duration_unit,required"`
}

func (r PlanNewParamsPricesNewPlanTieredBpsPriceBillingCycleConfiguration) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// The unit of billing period duration.
type PlanNewParamsPricesNewPlanTieredBpsPriceBillingCycleConfigurationDurationUnit string

const (
	PlanNewParamsPricesNewPlanTieredBpsPriceBillingCycleConfigurationDurationUnitDay   PlanNewParamsPricesNewPlanTieredBpsPriceBillingCycleConfigurationDurationUnit = "day"
	PlanNewParamsPricesNewPlanTieredBpsPriceBillingCycleConfigurationDurationUnitMonth PlanNewParamsPricesNewPlanTieredBpsPriceBillingCycleConfigurationDurationUnit = "month"
)

func (r PlanNewParamsPricesNewPlanTieredBpsPriceBillingCycleConfigurationDurationUnit) IsKnown() bool {
	switch r {
	case PlanNewParamsPricesNewPlanTieredBpsPriceBillingCycleConfigurationDurationUnitDay, PlanNewParamsPricesNewPlanTieredBpsPriceBillingCycleConfigurationDurationUnitMonth:
		return true
	}
	return false
}

// Within each billing cycle, specifies the cadence at which invoices are produced.
// If unspecified, a single invoice is produced per billing cycle.
type PlanNewParamsPricesNewPlanTieredBpsPriceInvoicingCycleConfiguration struct {
	// The duration of the billing period.
	Duration param.Field[int64] `json:"duration,required"`
	// The unit of billing period duration.
	DurationUnit param.Field[PlanNewParamsPricesNewPlanTieredBpsPriceInvoicingCycleConfigurationDurationUnit] `json:"duration_unit,required"`
}

func (r PlanNewParamsPricesNewPlanTieredBpsPriceInvoicingCycleConfiguration) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// The unit of billing period duration.
type PlanNewParamsPricesNewPlanTieredBpsPriceInvoicingCycleConfigurationDurationUnit string

const (
	PlanNewParamsPricesNewPlanTieredBpsPriceInvoicingCycleConfigurationDurationUnitDay   PlanNewParamsPricesNewPlanTieredBpsPriceInvoicingCycleConfigurationDurationUnit = "day"
	PlanNewParamsPricesNewPlanTieredBpsPriceInvoicingCycleConfigurationDurationUnitMonth PlanNewParamsPricesNewPlanTieredBpsPriceInvoicingCycleConfigurationDurationUnit = "month"
)

func (r PlanNewParamsPricesNewPlanTieredBpsPriceInvoicingCycleConfigurationDurationUnit) IsKnown() bool {
	switch r {
	case PlanNewParamsPricesNewPlanTieredBpsPriceInvoicingCycleConfigurationDurationUnitDay, PlanNewParamsPricesNewPlanTieredBpsPriceInvoicingCycleConfigurationDurationUnitMonth:
		return true
	}
	return false
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
	// For custom cadence: specifies the duration of the billing period in days or
	// months.
	BillingCycleConfiguration param.Field[PlanNewParamsPricesNewPlanBpsPriceBillingCycleConfiguration] `json:"billing_cycle_configuration"`
	// The per unit conversion rate of the price currency to the invoicing currency.
	ConversionRate param.Field[float64] `json:"conversion_rate"`
	// An ISO 4217 currency string, or custom pricing unit identifier, in which this
	// price is billed.
	Currency param.Field[string] `json:"currency"`
	// An alias for the price.
	ExternalPriceID param.Field[string] `json:"external_price_id"`
	// If the Price represents a fixed cost, this represents the quantity of units
	// applied.
	FixedPriceQuantity param.Field[float64] `json:"fixed_price_quantity"`
	// The property used to group this price on an invoice
	InvoiceGroupingKey param.Field[string] `json:"invoice_grouping_key"`
	// Within each billing cycle, specifies the cadence at which invoices are produced.
	// If unspecified, a single invoice is produced per billing cycle.
	InvoicingCycleConfiguration param.Field[PlanNewParamsPricesNewPlanBpsPriceInvoicingCycleConfiguration] `json:"invoicing_cycle_configuration"`
	// User-specified key/value pairs for the resource. Individual keys can be removed
	// by setting the value to `null`, and the entire metadata mapping can be cleared
	// by setting `metadata` to `null`.
	Metadata param.Field[map[string]string] `json:"metadata"`
}

func (r PlanNewParamsPricesNewPlanBpsPrice) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r PlanNewParamsPricesNewPlanBpsPrice) implementsPlanNewParamsPriceUnion() {}

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
	PlanNewParamsPricesNewPlanBpsPriceCadenceAnnual     PlanNewParamsPricesNewPlanBpsPriceCadence = "annual"
	PlanNewParamsPricesNewPlanBpsPriceCadenceSemiAnnual PlanNewParamsPricesNewPlanBpsPriceCadence = "semi_annual"
	PlanNewParamsPricesNewPlanBpsPriceCadenceMonthly    PlanNewParamsPricesNewPlanBpsPriceCadence = "monthly"
	PlanNewParamsPricesNewPlanBpsPriceCadenceQuarterly  PlanNewParamsPricesNewPlanBpsPriceCadence = "quarterly"
	PlanNewParamsPricesNewPlanBpsPriceCadenceOneTime    PlanNewParamsPricesNewPlanBpsPriceCadence = "one_time"
	PlanNewParamsPricesNewPlanBpsPriceCadenceCustom     PlanNewParamsPricesNewPlanBpsPriceCadence = "custom"
)

func (r PlanNewParamsPricesNewPlanBpsPriceCadence) IsKnown() bool {
	switch r {
	case PlanNewParamsPricesNewPlanBpsPriceCadenceAnnual, PlanNewParamsPricesNewPlanBpsPriceCadenceSemiAnnual, PlanNewParamsPricesNewPlanBpsPriceCadenceMonthly, PlanNewParamsPricesNewPlanBpsPriceCadenceQuarterly, PlanNewParamsPricesNewPlanBpsPriceCadenceOneTime, PlanNewParamsPricesNewPlanBpsPriceCadenceCustom:
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

// For custom cadence: specifies the duration of the billing period in days or
// months.
type PlanNewParamsPricesNewPlanBpsPriceBillingCycleConfiguration struct {
	// The duration of the billing period.
	Duration param.Field[int64] `json:"duration,required"`
	// The unit of billing period duration.
	DurationUnit param.Field[PlanNewParamsPricesNewPlanBpsPriceBillingCycleConfigurationDurationUnit] `json:"duration_unit,required"`
}

func (r PlanNewParamsPricesNewPlanBpsPriceBillingCycleConfiguration) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// The unit of billing period duration.
type PlanNewParamsPricesNewPlanBpsPriceBillingCycleConfigurationDurationUnit string

const (
	PlanNewParamsPricesNewPlanBpsPriceBillingCycleConfigurationDurationUnitDay   PlanNewParamsPricesNewPlanBpsPriceBillingCycleConfigurationDurationUnit = "day"
	PlanNewParamsPricesNewPlanBpsPriceBillingCycleConfigurationDurationUnitMonth PlanNewParamsPricesNewPlanBpsPriceBillingCycleConfigurationDurationUnit = "month"
)

func (r PlanNewParamsPricesNewPlanBpsPriceBillingCycleConfigurationDurationUnit) IsKnown() bool {
	switch r {
	case PlanNewParamsPricesNewPlanBpsPriceBillingCycleConfigurationDurationUnitDay, PlanNewParamsPricesNewPlanBpsPriceBillingCycleConfigurationDurationUnitMonth:
		return true
	}
	return false
}

// Within each billing cycle, specifies the cadence at which invoices are produced.
// If unspecified, a single invoice is produced per billing cycle.
type PlanNewParamsPricesNewPlanBpsPriceInvoicingCycleConfiguration struct {
	// The duration of the billing period.
	Duration param.Field[int64] `json:"duration,required"`
	// The unit of billing period duration.
	DurationUnit param.Field[PlanNewParamsPricesNewPlanBpsPriceInvoicingCycleConfigurationDurationUnit] `json:"duration_unit,required"`
}

func (r PlanNewParamsPricesNewPlanBpsPriceInvoicingCycleConfiguration) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// The unit of billing period duration.
type PlanNewParamsPricesNewPlanBpsPriceInvoicingCycleConfigurationDurationUnit string

const (
	PlanNewParamsPricesNewPlanBpsPriceInvoicingCycleConfigurationDurationUnitDay   PlanNewParamsPricesNewPlanBpsPriceInvoicingCycleConfigurationDurationUnit = "day"
	PlanNewParamsPricesNewPlanBpsPriceInvoicingCycleConfigurationDurationUnitMonth PlanNewParamsPricesNewPlanBpsPriceInvoicingCycleConfigurationDurationUnit = "month"
)

func (r PlanNewParamsPricesNewPlanBpsPriceInvoicingCycleConfigurationDurationUnit) IsKnown() bool {
	switch r {
	case PlanNewParamsPricesNewPlanBpsPriceInvoicingCycleConfigurationDurationUnitDay, PlanNewParamsPricesNewPlanBpsPriceInvoicingCycleConfigurationDurationUnitMonth:
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
	// For custom cadence: specifies the duration of the billing period in days or
	// months.
	BillingCycleConfiguration param.Field[PlanNewParamsPricesNewPlanBulkBpsPriceBillingCycleConfiguration] `json:"billing_cycle_configuration"`
	// The per unit conversion rate of the price currency to the invoicing currency.
	ConversionRate param.Field[float64] `json:"conversion_rate"`
	// An ISO 4217 currency string, or custom pricing unit identifier, in which this
	// price is billed.
	Currency param.Field[string] `json:"currency"`
	// An alias for the price.
	ExternalPriceID param.Field[string] `json:"external_price_id"`
	// If the Price represents a fixed cost, this represents the quantity of units
	// applied.
	FixedPriceQuantity param.Field[float64] `json:"fixed_price_quantity"`
	// The property used to group this price on an invoice
	InvoiceGroupingKey param.Field[string] `json:"invoice_grouping_key"`
	// Within each billing cycle, specifies the cadence at which invoices are produced.
	// If unspecified, a single invoice is produced per billing cycle.
	InvoicingCycleConfiguration param.Field[PlanNewParamsPricesNewPlanBulkBpsPriceInvoicingCycleConfiguration] `json:"invoicing_cycle_configuration"`
	// User-specified key/value pairs for the resource. Individual keys can be removed
	// by setting the value to `null`, and the entire metadata mapping can be cleared
	// by setting `metadata` to `null`.
	Metadata param.Field[map[string]string] `json:"metadata"`
}

func (r PlanNewParamsPricesNewPlanBulkBpsPrice) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r PlanNewParamsPricesNewPlanBulkBpsPrice) implementsPlanNewParamsPriceUnion() {}

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
	PlanNewParamsPricesNewPlanBulkBpsPriceCadenceAnnual     PlanNewParamsPricesNewPlanBulkBpsPriceCadence = "annual"
	PlanNewParamsPricesNewPlanBulkBpsPriceCadenceSemiAnnual PlanNewParamsPricesNewPlanBulkBpsPriceCadence = "semi_annual"
	PlanNewParamsPricesNewPlanBulkBpsPriceCadenceMonthly    PlanNewParamsPricesNewPlanBulkBpsPriceCadence = "monthly"
	PlanNewParamsPricesNewPlanBulkBpsPriceCadenceQuarterly  PlanNewParamsPricesNewPlanBulkBpsPriceCadence = "quarterly"
	PlanNewParamsPricesNewPlanBulkBpsPriceCadenceOneTime    PlanNewParamsPricesNewPlanBulkBpsPriceCadence = "one_time"
	PlanNewParamsPricesNewPlanBulkBpsPriceCadenceCustom     PlanNewParamsPricesNewPlanBulkBpsPriceCadence = "custom"
)

func (r PlanNewParamsPricesNewPlanBulkBpsPriceCadence) IsKnown() bool {
	switch r {
	case PlanNewParamsPricesNewPlanBulkBpsPriceCadenceAnnual, PlanNewParamsPricesNewPlanBulkBpsPriceCadenceSemiAnnual, PlanNewParamsPricesNewPlanBulkBpsPriceCadenceMonthly, PlanNewParamsPricesNewPlanBulkBpsPriceCadenceQuarterly, PlanNewParamsPricesNewPlanBulkBpsPriceCadenceOneTime, PlanNewParamsPricesNewPlanBulkBpsPriceCadenceCustom:
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

// For custom cadence: specifies the duration of the billing period in days or
// months.
type PlanNewParamsPricesNewPlanBulkBpsPriceBillingCycleConfiguration struct {
	// The duration of the billing period.
	Duration param.Field[int64] `json:"duration,required"`
	// The unit of billing period duration.
	DurationUnit param.Field[PlanNewParamsPricesNewPlanBulkBpsPriceBillingCycleConfigurationDurationUnit] `json:"duration_unit,required"`
}

func (r PlanNewParamsPricesNewPlanBulkBpsPriceBillingCycleConfiguration) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// The unit of billing period duration.
type PlanNewParamsPricesNewPlanBulkBpsPriceBillingCycleConfigurationDurationUnit string

const (
	PlanNewParamsPricesNewPlanBulkBpsPriceBillingCycleConfigurationDurationUnitDay   PlanNewParamsPricesNewPlanBulkBpsPriceBillingCycleConfigurationDurationUnit = "day"
	PlanNewParamsPricesNewPlanBulkBpsPriceBillingCycleConfigurationDurationUnitMonth PlanNewParamsPricesNewPlanBulkBpsPriceBillingCycleConfigurationDurationUnit = "month"
)

func (r PlanNewParamsPricesNewPlanBulkBpsPriceBillingCycleConfigurationDurationUnit) IsKnown() bool {
	switch r {
	case PlanNewParamsPricesNewPlanBulkBpsPriceBillingCycleConfigurationDurationUnitDay, PlanNewParamsPricesNewPlanBulkBpsPriceBillingCycleConfigurationDurationUnitMonth:
		return true
	}
	return false
}

// Within each billing cycle, specifies the cadence at which invoices are produced.
// If unspecified, a single invoice is produced per billing cycle.
type PlanNewParamsPricesNewPlanBulkBpsPriceInvoicingCycleConfiguration struct {
	// The duration of the billing period.
	Duration param.Field[int64] `json:"duration,required"`
	// The unit of billing period duration.
	DurationUnit param.Field[PlanNewParamsPricesNewPlanBulkBpsPriceInvoicingCycleConfigurationDurationUnit] `json:"duration_unit,required"`
}

func (r PlanNewParamsPricesNewPlanBulkBpsPriceInvoicingCycleConfiguration) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// The unit of billing period duration.
type PlanNewParamsPricesNewPlanBulkBpsPriceInvoicingCycleConfigurationDurationUnit string

const (
	PlanNewParamsPricesNewPlanBulkBpsPriceInvoicingCycleConfigurationDurationUnitDay   PlanNewParamsPricesNewPlanBulkBpsPriceInvoicingCycleConfigurationDurationUnit = "day"
	PlanNewParamsPricesNewPlanBulkBpsPriceInvoicingCycleConfigurationDurationUnitMonth PlanNewParamsPricesNewPlanBulkBpsPriceInvoicingCycleConfigurationDurationUnit = "month"
)

func (r PlanNewParamsPricesNewPlanBulkBpsPriceInvoicingCycleConfigurationDurationUnit) IsKnown() bool {
	switch r {
	case PlanNewParamsPricesNewPlanBulkBpsPriceInvoicingCycleConfigurationDurationUnitDay, PlanNewParamsPricesNewPlanBulkBpsPriceInvoicingCycleConfigurationDurationUnitMonth:
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
	// For custom cadence: specifies the duration of the billing period in days or
	// months.
	BillingCycleConfiguration param.Field[PlanNewParamsPricesNewPlanBulkPriceBillingCycleConfiguration] `json:"billing_cycle_configuration"`
	// The per unit conversion rate of the price currency to the invoicing currency.
	ConversionRate param.Field[float64] `json:"conversion_rate"`
	// An ISO 4217 currency string, or custom pricing unit identifier, in which this
	// price is billed.
	Currency param.Field[string] `json:"currency"`
	// An alias for the price.
	ExternalPriceID param.Field[string] `json:"external_price_id"`
	// If the Price represents a fixed cost, this represents the quantity of units
	// applied.
	FixedPriceQuantity param.Field[float64] `json:"fixed_price_quantity"`
	// The property used to group this price on an invoice
	InvoiceGroupingKey param.Field[string] `json:"invoice_grouping_key"`
	// Within each billing cycle, specifies the cadence at which invoices are produced.
	// If unspecified, a single invoice is produced per billing cycle.
	InvoicingCycleConfiguration param.Field[PlanNewParamsPricesNewPlanBulkPriceInvoicingCycleConfiguration] `json:"invoicing_cycle_configuration"`
	// User-specified key/value pairs for the resource. Individual keys can be removed
	// by setting the value to `null`, and the entire metadata mapping can be cleared
	// by setting `metadata` to `null`.
	Metadata param.Field[map[string]string] `json:"metadata"`
}

func (r PlanNewParamsPricesNewPlanBulkPrice) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r PlanNewParamsPricesNewPlanBulkPrice) implementsPlanNewParamsPriceUnion() {}

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
	PlanNewParamsPricesNewPlanBulkPriceCadenceAnnual     PlanNewParamsPricesNewPlanBulkPriceCadence = "annual"
	PlanNewParamsPricesNewPlanBulkPriceCadenceSemiAnnual PlanNewParamsPricesNewPlanBulkPriceCadence = "semi_annual"
	PlanNewParamsPricesNewPlanBulkPriceCadenceMonthly    PlanNewParamsPricesNewPlanBulkPriceCadence = "monthly"
	PlanNewParamsPricesNewPlanBulkPriceCadenceQuarterly  PlanNewParamsPricesNewPlanBulkPriceCadence = "quarterly"
	PlanNewParamsPricesNewPlanBulkPriceCadenceOneTime    PlanNewParamsPricesNewPlanBulkPriceCadence = "one_time"
	PlanNewParamsPricesNewPlanBulkPriceCadenceCustom     PlanNewParamsPricesNewPlanBulkPriceCadence = "custom"
)

func (r PlanNewParamsPricesNewPlanBulkPriceCadence) IsKnown() bool {
	switch r {
	case PlanNewParamsPricesNewPlanBulkPriceCadenceAnnual, PlanNewParamsPricesNewPlanBulkPriceCadenceSemiAnnual, PlanNewParamsPricesNewPlanBulkPriceCadenceMonthly, PlanNewParamsPricesNewPlanBulkPriceCadenceQuarterly, PlanNewParamsPricesNewPlanBulkPriceCadenceOneTime, PlanNewParamsPricesNewPlanBulkPriceCadenceCustom:
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

// For custom cadence: specifies the duration of the billing period in days or
// months.
type PlanNewParamsPricesNewPlanBulkPriceBillingCycleConfiguration struct {
	// The duration of the billing period.
	Duration param.Field[int64] `json:"duration,required"`
	// The unit of billing period duration.
	DurationUnit param.Field[PlanNewParamsPricesNewPlanBulkPriceBillingCycleConfigurationDurationUnit] `json:"duration_unit,required"`
}

func (r PlanNewParamsPricesNewPlanBulkPriceBillingCycleConfiguration) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// The unit of billing period duration.
type PlanNewParamsPricesNewPlanBulkPriceBillingCycleConfigurationDurationUnit string

const (
	PlanNewParamsPricesNewPlanBulkPriceBillingCycleConfigurationDurationUnitDay   PlanNewParamsPricesNewPlanBulkPriceBillingCycleConfigurationDurationUnit = "day"
	PlanNewParamsPricesNewPlanBulkPriceBillingCycleConfigurationDurationUnitMonth PlanNewParamsPricesNewPlanBulkPriceBillingCycleConfigurationDurationUnit = "month"
)

func (r PlanNewParamsPricesNewPlanBulkPriceBillingCycleConfigurationDurationUnit) IsKnown() bool {
	switch r {
	case PlanNewParamsPricesNewPlanBulkPriceBillingCycleConfigurationDurationUnitDay, PlanNewParamsPricesNewPlanBulkPriceBillingCycleConfigurationDurationUnitMonth:
		return true
	}
	return false
}

// Within each billing cycle, specifies the cadence at which invoices are produced.
// If unspecified, a single invoice is produced per billing cycle.
type PlanNewParamsPricesNewPlanBulkPriceInvoicingCycleConfiguration struct {
	// The duration of the billing period.
	Duration param.Field[int64] `json:"duration,required"`
	// The unit of billing period duration.
	DurationUnit param.Field[PlanNewParamsPricesNewPlanBulkPriceInvoicingCycleConfigurationDurationUnit] `json:"duration_unit,required"`
}

func (r PlanNewParamsPricesNewPlanBulkPriceInvoicingCycleConfiguration) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// The unit of billing period duration.
type PlanNewParamsPricesNewPlanBulkPriceInvoicingCycleConfigurationDurationUnit string

const (
	PlanNewParamsPricesNewPlanBulkPriceInvoicingCycleConfigurationDurationUnitDay   PlanNewParamsPricesNewPlanBulkPriceInvoicingCycleConfigurationDurationUnit = "day"
	PlanNewParamsPricesNewPlanBulkPriceInvoicingCycleConfigurationDurationUnitMonth PlanNewParamsPricesNewPlanBulkPriceInvoicingCycleConfigurationDurationUnit = "month"
)

func (r PlanNewParamsPricesNewPlanBulkPriceInvoicingCycleConfigurationDurationUnit) IsKnown() bool {
	switch r {
	case PlanNewParamsPricesNewPlanBulkPriceInvoicingCycleConfigurationDurationUnitDay, PlanNewParamsPricesNewPlanBulkPriceInvoicingCycleConfigurationDurationUnitMonth:
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
	// For custom cadence: specifies the duration of the billing period in days or
	// months.
	BillingCycleConfiguration param.Field[PlanNewParamsPricesNewPlanThresholdTotalAmountPriceBillingCycleConfiguration] `json:"billing_cycle_configuration"`
	// The per unit conversion rate of the price currency to the invoicing currency.
	ConversionRate param.Field[float64] `json:"conversion_rate"`
	// An ISO 4217 currency string, or custom pricing unit identifier, in which this
	// price is billed.
	Currency param.Field[string] `json:"currency"`
	// An alias for the price.
	ExternalPriceID param.Field[string] `json:"external_price_id"`
	// If the Price represents a fixed cost, this represents the quantity of units
	// applied.
	FixedPriceQuantity param.Field[float64] `json:"fixed_price_quantity"`
	// The property used to group this price on an invoice
	InvoiceGroupingKey param.Field[string] `json:"invoice_grouping_key"`
	// Within each billing cycle, specifies the cadence at which invoices are produced.
	// If unspecified, a single invoice is produced per billing cycle.
	InvoicingCycleConfiguration param.Field[PlanNewParamsPricesNewPlanThresholdTotalAmountPriceInvoicingCycleConfiguration] `json:"invoicing_cycle_configuration"`
	// User-specified key/value pairs for the resource. Individual keys can be removed
	// by setting the value to `null`, and the entire metadata mapping can be cleared
	// by setting `metadata` to `null`.
	Metadata param.Field[map[string]string] `json:"metadata"`
}

func (r PlanNewParamsPricesNewPlanThresholdTotalAmountPrice) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r PlanNewParamsPricesNewPlanThresholdTotalAmountPrice) implementsPlanNewParamsPriceUnion() {}

// The cadence to bill for this price on.
type PlanNewParamsPricesNewPlanThresholdTotalAmountPriceCadence string

const (
	PlanNewParamsPricesNewPlanThresholdTotalAmountPriceCadenceAnnual     PlanNewParamsPricesNewPlanThresholdTotalAmountPriceCadence = "annual"
	PlanNewParamsPricesNewPlanThresholdTotalAmountPriceCadenceSemiAnnual PlanNewParamsPricesNewPlanThresholdTotalAmountPriceCadence = "semi_annual"
	PlanNewParamsPricesNewPlanThresholdTotalAmountPriceCadenceMonthly    PlanNewParamsPricesNewPlanThresholdTotalAmountPriceCadence = "monthly"
	PlanNewParamsPricesNewPlanThresholdTotalAmountPriceCadenceQuarterly  PlanNewParamsPricesNewPlanThresholdTotalAmountPriceCadence = "quarterly"
	PlanNewParamsPricesNewPlanThresholdTotalAmountPriceCadenceOneTime    PlanNewParamsPricesNewPlanThresholdTotalAmountPriceCadence = "one_time"
	PlanNewParamsPricesNewPlanThresholdTotalAmountPriceCadenceCustom     PlanNewParamsPricesNewPlanThresholdTotalAmountPriceCadence = "custom"
)

func (r PlanNewParamsPricesNewPlanThresholdTotalAmountPriceCadence) IsKnown() bool {
	switch r {
	case PlanNewParamsPricesNewPlanThresholdTotalAmountPriceCadenceAnnual, PlanNewParamsPricesNewPlanThresholdTotalAmountPriceCadenceSemiAnnual, PlanNewParamsPricesNewPlanThresholdTotalAmountPriceCadenceMonthly, PlanNewParamsPricesNewPlanThresholdTotalAmountPriceCadenceQuarterly, PlanNewParamsPricesNewPlanThresholdTotalAmountPriceCadenceOneTime, PlanNewParamsPricesNewPlanThresholdTotalAmountPriceCadenceCustom:
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

// For custom cadence: specifies the duration of the billing period in days or
// months.
type PlanNewParamsPricesNewPlanThresholdTotalAmountPriceBillingCycleConfiguration struct {
	// The duration of the billing period.
	Duration param.Field[int64] `json:"duration,required"`
	// The unit of billing period duration.
	DurationUnit param.Field[PlanNewParamsPricesNewPlanThresholdTotalAmountPriceBillingCycleConfigurationDurationUnit] `json:"duration_unit,required"`
}

func (r PlanNewParamsPricesNewPlanThresholdTotalAmountPriceBillingCycleConfiguration) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// The unit of billing period duration.
type PlanNewParamsPricesNewPlanThresholdTotalAmountPriceBillingCycleConfigurationDurationUnit string

const (
	PlanNewParamsPricesNewPlanThresholdTotalAmountPriceBillingCycleConfigurationDurationUnitDay   PlanNewParamsPricesNewPlanThresholdTotalAmountPriceBillingCycleConfigurationDurationUnit = "day"
	PlanNewParamsPricesNewPlanThresholdTotalAmountPriceBillingCycleConfigurationDurationUnitMonth PlanNewParamsPricesNewPlanThresholdTotalAmountPriceBillingCycleConfigurationDurationUnit = "month"
)

func (r PlanNewParamsPricesNewPlanThresholdTotalAmountPriceBillingCycleConfigurationDurationUnit) IsKnown() bool {
	switch r {
	case PlanNewParamsPricesNewPlanThresholdTotalAmountPriceBillingCycleConfigurationDurationUnitDay, PlanNewParamsPricesNewPlanThresholdTotalAmountPriceBillingCycleConfigurationDurationUnitMonth:
		return true
	}
	return false
}

// Within each billing cycle, specifies the cadence at which invoices are produced.
// If unspecified, a single invoice is produced per billing cycle.
type PlanNewParamsPricesNewPlanThresholdTotalAmountPriceInvoicingCycleConfiguration struct {
	// The duration of the billing period.
	Duration param.Field[int64] `json:"duration,required"`
	// The unit of billing period duration.
	DurationUnit param.Field[PlanNewParamsPricesNewPlanThresholdTotalAmountPriceInvoicingCycleConfigurationDurationUnit] `json:"duration_unit,required"`
}

func (r PlanNewParamsPricesNewPlanThresholdTotalAmountPriceInvoicingCycleConfiguration) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// The unit of billing period duration.
type PlanNewParamsPricesNewPlanThresholdTotalAmountPriceInvoicingCycleConfigurationDurationUnit string

const (
	PlanNewParamsPricesNewPlanThresholdTotalAmountPriceInvoicingCycleConfigurationDurationUnitDay   PlanNewParamsPricesNewPlanThresholdTotalAmountPriceInvoicingCycleConfigurationDurationUnit = "day"
	PlanNewParamsPricesNewPlanThresholdTotalAmountPriceInvoicingCycleConfigurationDurationUnitMonth PlanNewParamsPricesNewPlanThresholdTotalAmountPriceInvoicingCycleConfigurationDurationUnit = "month"
)

func (r PlanNewParamsPricesNewPlanThresholdTotalAmountPriceInvoicingCycleConfigurationDurationUnit) IsKnown() bool {
	switch r {
	case PlanNewParamsPricesNewPlanThresholdTotalAmountPriceInvoicingCycleConfigurationDurationUnitDay, PlanNewParamsPricesNewPlanThresholdTotalAmountPriceInvoicingCycleConfigurationDurationUnitMonth:
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
	// For custom cadence: specifies the duration of the billing period in days or
	// months.
	BillingCycleConfiguration param.Field[PlanNewParamsPricesNewPlanTieredPackagePriceBillingCycleConfiguration] `json:"billing_cycle_configuration"`
	// The per unit conversion rate of the price currency to the invoicing currency.
	ConversionRate param.Field[float64] `json:"conversion_rate"`
	// An ISO 4217 currency string, or custom pricing unit identifier, in which this
	// price is billed.
	Currency param.Field[string] `json:"currency"`
	// An alias for the price.
	ExternalPriceID param.Field[string] `json:"external_price_id"`
	// If the Price represents a fixed cost, this represents the quantity of units
	// applied.
	FixedPriceQuantity param.Field[float64] `json:"fixed_price_quantity"`
	// The property used to group this price on an invoice
	InvoiceGroupingKey param.Field[string] `json:"invoice_grouping_key"`
	// Within each billing cycle, specifies the cadence at which invoices are produced.
	// If unspecified, a single invoice is produced per billing cycle.
	InvoicingCycleConfiguration param.Field[PlanNewParamsPricesNewPlanTieredPackagePriceInvoicingCycleConfiguration] `json:"invoicing_cycle_configuration"`
	// User-specified key/value pairs for the resource. Individual keys can be removed
	// by setting the value to `null`, and the entire metadata mapping can be cleared
	// by setting `metadata` to `null`.
	Metadata param.Field[map[string]string] `json:"metadata"`
}

func (r PlanNewParamsPricesNewPlanTieredPackagePrice) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r PlanNewParamsPricesNewPlanTieredPackagePrice) implementsPlanNewParamsPriceUnion() {}

// The cadence to bill for this price on.
type PlanNewParamsPricesNewPlanTieredPackagePriceCadence string

const (
	PlanNewParamsPricesNewPlanTieredPackagePriceCadenceAnnual     PlanNewParamsPricesNewPlanTieredPackagePriceCadence = "annual"
	PlanNewParamsPricesNewPlanTieredPackagePriceCadenceSemiAnnual PlanNewParamsPricesNewPlanTieredPackagePriceCadence = "semi_annual"
	PlanNewParamsPricesNewPlanTieredPackagePriceCadenceMonthly    PlanNewParamsPricesNewPlanTieredPackagePriceCadence = "monthly"
	PlanNewParamsPricesNewPlanTieredPackagePriceCadenceQuarterly  PlanNewParamsPricesNewPlanTieredPackagePriceCadence = "quarterly"
	PlanNewParamsPricesNewPlanTieredPackagePriceCadenceOneTime    PlanNewParamsPricesNewPlanTieredPackagePriceCadence = "one_time"
	PlanNewParamsPricesNewPlanTieredPackagePriceCadenceCustom     PlanNewParamsPricesNewPlanTieredPackagePriceCadence = "custom"
)

func (r PlanNewParamsPricesNewPlanTieredPackagePriceCadence) IsKnown() bool {
	switch r {
	case PlanNewParamsPricesNewPlanTieredPackagePriceCadenceAnnual, PlanNewParamsPricesNewPlanTieredPackagePriceCadenceSemiAnnual, PlanNewParamsPricesNewPlanTieredPackagePriceCadenceMonthly, PlanNewParamsPricesNewPlanTieredPackagePriceCadenceQuarterly, PlanNewParamsPricesNewPlanTieredPackagePriceCadenceOneTime, PlanNewParamsPricesNewPlanTieredPackagePriceCadenceCustom:
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

// For custom cadence: specifies the duration of the billing period in days or
// months.
type PlanNewParamsPricesNewPlanTieredPackagePriceBillingCycleConfiguration struct {
	// The duration of the billing period.
	Duration param.Field[int64] `json:"duration,required"`
	// The unit of billing period duration.
	DurationUnit param.Field[PlanNewParamsPricesNewPlanTieredPackagePriceBillingCycleConfigurationDurationUnit] `json:"duration_unit,required"`
}

func (r PlanNewParamsPricesNewPlanTieredPackagePriceBillingCycleConfiguration) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// The unit of billing period duration.
type PlanNewParamsPricesNewPlanTieredPackagePriceBillingCycleConfigurationDurationUnit string

const (
	PlanNewParamsPricesNewPlanTieredPackagePriceBillingCycleConfigurationDurationUnitDay   PlanNewParamsPricesNewPlanTieredPackagePriceBillingCycleConfigurationDurationUnit = "day"
	PlanNewParamsPricesNewPlanTieredPackagePriceBillingCycleConfigurationDurationUnitMonth PlanNewParamsPricesNewPlanTieredPackagePriceBillingCycleConfigurationDurationUnit = "month"
)

func (r PlanNewParamsPricesNewPlanTieredPackagePriceBillingCycleConfigurationDurationUnit) IsKnown() bool {
	switch r {
	case PlanNewParamsPricesNewPlanTieredPackagePriceBillingCycleConfigurationDurationUnitDay, PlanNewParamsPricesNewPlanTieredPackagePriceBillingCycleConfigurationDurationUnitMonth:
		return true
	}
	return false
}

// Within each billing cycle, specifies the cadence at which invoices are produced.
// If unspecified, a single invoice is produced per billing cycle.
type PlanNewParamsPricesNewPlanTieredPackagePriceInvoicingCycleConfiguration struct {
	// The duration of the billing period.
	Duration param.Field[int64] `json:"duration,required"`
	// The unit of billing period duration.
	DurationUnit param.Field[PlanNewParamsPricesNewPlanTieredPackagePriceInvoicingCycleConfigurationDurationUnit] `json:"duration_unit,required"`
}

func (r PlanNewParamsPricesNewPlanTieredPackagePriceInvoicingCycleConfiguration) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// The unit of billing period duration.
type PlanNewParamsPricesNewPlanTieredPackagePriceInvoicingCycleConfigurationDurationUnit string

const (
	PlanNewParamsPricesNewPlanTieredPackagePriceInvoicingCycleConfigurationDurationUnitDay   PlanNewParamsPricesNewPlanTieredPackagePriceInvoicingCycleConfigurationDurationUnit = "day"
	PlanNewParamsPricesNewPlanTieredPackagePriceInvoicingCycleConfigurationDurationUnitMonth PlanNewParamsPricesNewPlanTieredPackagePriceInvoicingCycleConfigurationDurationUnit = "month"
)

func (r PlanNewParamsPricesNewPlanTieredPackagePriceInvoicingCycleConfigurationDurationUnit) IsKnown() bool {
	switch r {
	case PlanNewParamsPricesNewPlanTieredPackagePriceInvoicingCycleConfigurationDurationUnitDay, PlanNewParamsPricesNewPlanTieredPackagePriceInvoicingCycleConfigurationDurationUnitMonth:
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
	// For custom cadence: specifies the duration of the billing period in days or
	// months.
	BillingCycleConfiguration param.Field[PlanNewParamsPricesNewPlanTieredWithMinimumPriceBillingCycleConfiguration] `json:"billing_cycle_configuration"`
	// The per unit conversion rate of the price currency to the invoicing currency.
	ConversionRate param.Field[float64] `json:"conversion_rate"`
	// An ISO 4217 currency string, or custom pricing unit identifier, in which this
	// price is billed.
	Currency param.Field[string] `json:"currency"`
	// An alias for the price.
	ExternalPriceID param.Field[string] `json:"external_price_id"`
	// If the Price represents a fixed cost, this represents the quantity of units
	// applied.
	FixedPriceQuantity param.Field[float64] `json:"fixed_price_quantity"`
	// The property used to group this price on an invoice
	InvoiceGroupingKey param.Field[string] `json:"invoice_grouping_key"`
	// Within each billing cycle, specifies the cadence at which invoices are produced.
	// If unspecified, a single invoice is produced per billing cycle.
	InvoicingCycleConfiguration param.Field[PlanNewParamsPricesNewPlanTieredWithMinimumPriceInvoicingCycleConfiguration] `json:"invoicing_cycle_configuration"`
	// User-specified key/value pairs for the resource. Individual keys can be removed
	// by setting the value to `null`, and the entire metadata mapping can be cleared
	// by setting `metadata` to `null`.
	Metadata param.Field[map[string]string] `json:"metadata"`
}

func (r PlanNewParamsPricesNewPlanTieredWithMinimumPrice) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r PlanNewParamsPricesNewPlanTieredWithMinimumPrice) implementsPlanNewParamsPriceUnion() {}

// The cadence to bill for this price on.
type PlanNewParamsPricesNewPlanTieredWithMinimumPriceCadence string

const (
	PlanNewParamsPricesNewPlanTieredWithMinimumPriceCadenceAnnual     PlanNewParamsPricesNewPlanTieredWithMinimumPriceCadence = "annual"
	PlanNewParamsPricesNewPlanTieredWithMinimumPriceCadenceSemiAnnual PlanNewParamsPricesNewPlanTieredWithMinimumPriceCadence = "semi_annual"
	PlanNewParamsPricesNewPlanTieredWithMinimumPriceCadenceMonthly    PlanNewParamsPricesNewPlanTieredWithMinimumPriceCadence = "monthly"
	PlanNewParamsPricesNewPlanTieredWithMinimumPriceCadenceQuarterly  PlanNewParamsPricesNewPlanTieredWithMinimumPriceCadence = "quarterly"
	PlanNewParamsPricesNewPlanTieredWithMinimumPriceCadenceOneTime    PlanNewParamsPricesNewPlanTieredWithMinimumPriceCadence = "one_time"
	PlanNewParamsPricesNewPlanTieredWithMinimumPriceCadenceCustom     PlanNewParamsPricesNewPlanTieredWithMinimumPriceCadence = "custom"
)

func (r PlanNewParamsPricesNewPlanTieredWithMinimumPriceCadence) IsKnown() bool {
	switch r {
	case PlanNewParamsPricesNewPlanTieredWithMinimumPriceCadenceAnnual, PlanNewParamsPricesNewPlanTieredWithMinimumPriceCadenceSemiAnnual, PlanNewParamsPricesNewPlanTieredWithMinimumPriceCadenceMonthly, PlanNewParamsPricesNewPlanTieredWithMinimumPriceCadenceQuarterly, PlanNewParamsPricesNewPlanTieredWithMinimumPriceCadenceOneTime, PlanNewParamsPricesNewPlanTieredWithMinimumPriceCadenceCustom:
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

// For custom cadence: specifies the duration of the billing period in days or
// months.
type PlanNewParamsPricesNewPlanTieredWithMinimumPriceBillingCycleConfiguration struct {
	// The duration of the billing period.
	Duration param.Field[int64] `json:"duration,required"`
	// The unit of billing period duration.
	DurationUnit param.Field[PlanNewParamsPricesNewPlanTieredWithMinimumPriceBillingCycleConfigurationDurationUnit] `json:"duration_unit,required"`
}

func (r PlanNewParamsPricesNewPlanTieredWithMinimumPriceBillingCycleConfiguration) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// The unit of billing period duration.
type PlanNewParamsPricesNewPlanTieredWithMinimumPriceBillingCycleConfigurationDurationUnit string

const (
	PlanNewParamsPricesNewPlanTieredWithMinimumPriceBillingCycleConfigurationDurationUnitDay   PlanNewParamsPricesNewPlanTieredWithMinimumPriceBillingCycleConfigurationDurationUnit = "day"
	PlanNewParamsPricesNewPlanTieredWithMinimumPriceBillingCycleConfigurationDurationUnitMonth PlanNewParamsPricesNewPlanTieredWithMinimumPriceBillingCycleConfigurationDurationUnit = "month"
)

func (r PlanNewParamsPricesNewPlanTieredWithMinimumPriceBillingCycleConfigurationDurationUnit) IsKnown() bool {
	switch r {
	case PlanNewParamsPricesNewPlanTieredWithMinimumPriceBillingCycleConfigurationDurationUnitDay, PlanNewParamsPricesNewPlanTieredWithMinimumPriceBillingCycleConfigurationDurationUnitMonth:
		return true
	}
	return false
}

// Within each billing cycle, specifies the cadence at which invoices are produced.
// If unspecified, a single invoice is produced per billing cycle.
type PlanNewParamsPricesNewPlanTieredWithMinimumPriceInvoicingCycleConfiguration struct {
	// The duration of the billing period.
	Duration param.Field[int64] `json:"duration,required"`
	// The unit of billing period duration.
	DurationUnit param.Field[PlanNewParamsPricesNewPlanTieredWithMinimumPriceInvoicingCycleConfigurationDurationUnit] `json:"duration_unit,required"`
}

func (r PlanNewParamsPricesNewPlanTieredWithMinimumPriceInvoicingCycleConfiguration) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// The unit of billing period duration.
type PlanNewParamsPricesNewPlanTieredWithMinimumPriceInvoicingCycleConfigurationDurationUnit string

const (
	PlanNewParamsPricesNewPlanTieredWithMinimumPriceInvoicingCycleConfigurationDurationUnitDay   PlanNewParamsPricesNewPlanTieredWithMinimumPriceInvoicingCycleConfigurationDurationUnit = "day"
	PlanNewParamsPricesNewPlanTieredWithMinimumPriceInvoicingCycleConfigurationDurationUnitMonth PlanNewParamsPricesNewPlanTieredWithMinimumPriceInvoicingCycleConfigurationDurationUnit = "month"
)

func (r PlanNewParamsPricesNewPlanTieredWithMinimumPriceInvoicingCycleConfigurationDurationUnit) IsKnown() bool {
	switch r {
	case PlanNewParamsPricesNewPlanTieredWithMinimumPriceInvoicingCycleConfigurationDurationUnitDay, PlanNewParamsPricesNewPlanTieredWithMinimumPriceInvoicingCycleConfigurationDurationUnitMonth:
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
	// For custom cadence: specifies the duration of the billing period in days or
	// months.
	BillingCycleConfiguration param.Field[PlanNewParamsPricesNewPlanUnitWithPercentPriceBillingCycleConfiguration] `json:"billing_cycle_configuration"`
	// The per unit conversion rate of the price currency to the invoicing currency.
	ConversionRate param.Field[float64] `json:"conversion_rate"`
	// An ISO 4217 currency string, or custom pricing unit identifier, in which this
	// price is billed.
	Currency param.Field[string] `json:"currency"`
	// An alias for the price.
	ExternalPriceID param.Field[string] `json:"external_price_id"`
	// If the Price represents a fixed cost, this represents the quantity of units
	// applied.
	FixedPriceQuantity param.Field[float64] `json:"fixed_price_quantity"`
	// The property used to group this price on an invoice
	InvoiceGroupingKey param.Field[string] `json:"invoice_grouping_key"`
	// Within each billing cycle, specifies the cadence at which invoices are produced.
	// If unspecified, a single invoice is produced per billing cycle.
	InvoicingCycleConfiguration param.Field[PlanNewParamsPricesNewPlanUnitWithPercentPriceInvoicingCycleConfiguration] `json:"invoicing_cycle_configuration"`
	// User-specified key/value pairs for the resource. Individual keys can be removed
	// by setting the value to `null`, and the entire metadata mapping can be cleared
	// by setting `metadata` to `null`.
	Metadata param.Field[map[string]string] `json:"metadata"`
}

func (r PlanNewParamsPricesNewPlanUnitWithPercentPrice) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r PlanNewParamsPricesNewPlanUnitWithPercentPrice) implementsPlanNewParamsPriceUnion() {}

// The cadence to bill for this price on.
type PlanNewParamsPricesNewPlanUnitWithPercentPriceCadence string

const (
	PlanNewParamsPricesNewPlanUnitWithPercentPriceCadenceAnnual     PlanNewParamsPricesNewPlanUnitWithPercentPriceCadence = "annual"
	PlanNewParamsPricesNewPlanUnitWithPercentPriceCadenceSemiAnnual PlanNewParamsPricesNewPlanUnitWithPercentPriceCadence = "semi_annual"
	PlanNewParamsPricesNewPlanUnitWithPercentPriceCadenceMonthly    PlanNewParamsPricesNewPlanUnitWithPercentPriceCadence = "monthly"
	PlanNewParamsPricesNewPlanUnitWithPercentPriceCadenceQuarterly  PlanNewParamsPricesNewPlanUnitWithPercentPriceCadence = "quarterly"
	PlanNewParamsPricesNewPlanUnitWithPercentPriceCadenceOneTime    PlanNewParamsPricesNewPlanUnitWithPercentPriceCadence = "one_time"
	PlanNewParamsPricesNewPlanUnitWithPercentPriceCadenceCustom     PlanNewParamsPricesNewPlanUnitWithPercentPriceCadence = "custom"
)

func (r PlanNewParamsPricesNewPlanUnitWithPercentPriceCadence) IsKnown() bool {
	switch r {
	case PlanNewParamsPricesNewPlanUnitWithPercentPriceCadenceAnnual, PlanNewParamsPricesNewPlanUnitWithPercentPriceCadenceSemiAnnual, PlanNewParamsPricesNewPlanUnitWithPercentPriceCadenceMonthly, PlanNewParamsPricesNewPlanUnitWithPercentPriceCadenceQuarterly, PlanNewParamsPricesNewPlanUnitWithPercentPriceCadenceOneTime, PlanNewParamsPricesNewPlanUnitWithPercentPriceCadenceCustom:
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

// For custom cadence: specifies the duration of the billing period in days or
// months.
type PlanNewParamsPricesNewPlanUnitWithPercentPriceBillingCycleConfiguration struct {
	// The duration of the billing period.
	Duration param.Field[int64] `json:"duration,required"`
	// The unit of billing period duration.
	DurationUnit param.Field[PlanNewParamsPricesNewPlanUnitWithPercentPriceBillingCycleConfigurationDurationUnit] `json:"duration_unit,required"`
}

func (r PlanNewParamsPricesNewPlanUnitWithPercentPriceBillingCycleConfiguration) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// The unit of billing period duration.
type PlanNewParamsPricesNewPlanUnitWithPercentPriceBillingCycleConfigurationDurationUnit string

const (
	PlanNewParamsPricesNewPlanUnitWithPercentPriceBillingCycleConfigurationDurationUnitDay   PlanNewParamsPricesNewPlanUnitWithPercentPriceBillingCycleConfigurationDurationUnit = "day"
	PlanNewParamsPricesNewPlanUnitWithPercentPriceBillingCycleConfigurationDurationUnitMonth PlanNewParamsPricesNewPlanUnitWithPercentPriceBillingCycleConfigurationDurationUnit = "month"
)

func (r PlanNewParamsPricesNewPlanUnitWithPercentPriceBillingCycleConfigurationDurationUnit) IsKnown() bool {
	switch r {
	case PlanNewParamsPricesNewPlanUnitWithPercentPriceBillingCycleConfigurationDurationUnitDay, PlanNewParamsPricesNewPlanUnitWithPercentPriceBillingCycleConfigurationDurationUnitMonth:
		return true
	}
	return false
}

// Within each billing cycle, specifies the cadence at which invoices are produced.
// If unspecified, a single invoice is produced per billing cycle.
type PlanNewParamsPricesNewPlanUnitWithPercentPriceInvoicingCycleConfiguration struct {
	// The duration of the billing period.
	Duration param.Field[int64] `json:"duration,required"`
	// The unit of billing period duration.
	DurationUnit param.Field[PlanNewParamsPricesNewPlanUnitWithPercentPriceInvoicingCycleConfigurationDurationUnit] `json:"duration_unit,required"`
}

func (r PlanNewParamsPricesNewPlanUnitWithPercentPriceInvoicingCycleConfiguration) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// The unit of billing period duration.
type PlanNewParamsPricesNewPlanUnitWithPercentPriceInvoicingCycleConfigurationDurationUnit string

const (
	PlanNewParamsPricesNewPlanUnitWithPercentPriceInvoicingCycleConfigurationDurationUnitDay   PlanNewParamsPricesNewPlanUnitWithPercentPriceInvoicingCycleConfigurationDurationUnit = "day"
	PlanNewParamsPricesNewPlanUnitWithPercentPriceInvoicingCycleConfigurationDurationUnitMonth PlanNewParamsPricesNewPlanUnitWithPercentPriceInvoicingCycleConfigurationDurationUnit = "month"
)

func (r PlanNewParamsPricesNewPlanUnitWithPercentPriceInvoicingCycleConfigurationDurationUnit) IsKnown() bool {
	switch r {
	case PlanNewParamsPricesNewPlanUnitWithPercentPriceInvoicingCycleConfigurationDurationUnitDay, PlanNewParamsPricesNewPlanUnitWithPercentPriceInvoicingCycleConfigurationDurationUnitMonth:
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
	// For custom cadence: specifies the duration of the billing period in days or
	// months.
	BillingCycleConfiguration param.Field[PlanNewParamsPricesNewPlanPackageWithAllocationPriceBillingCycleConfiguration] `json:"billing_cycle_configuration"`
	// The per unit conversion rate of the price currency to the invoicing currency.
	ConversionRate param.Field[float64] `json:"conversion_rate"`
	// An ISO 4217 currency string, or custom pricing unit identifier, in which this
	// price is billed.
	Currency param.Field[string] `json:"currency"`
	// An alias for the price.
	ExternalPriceID param.Field[string] `json:"external_price_id"`
	// If the Price represents a fixed cost, this represents the quantity of units
	// applied.
	FixedPriceQuantity param.Field[float64] `json:"fixed_price_quantity"`
	// The property used to group this price on an invoice
	InvoiceGroupingKey param.Field[string] `json:"invoice_grouping_key"`
	// Within each billing cycle, specifies the cadence at which invoices are produced.
	// If unspecified, a single invoice is produced per billing cycle.
	InvoicingCycleConfiguration param.Field[PlanNewParamsPricesNewPlanPackageWithAllocationPriceInvoicingCycleConfiguration] `json:"invoicing_cycle_configuration"`
	// User-specified key/value pairs for the resource. Individual keys can be removed
	// by setting the value to `null`, and the entire metadata mapping can be cleared
	// by setting `metadata` to `null`.
	Metadata param.Field[map[string]string] `json:"metadata"`
}

func (r PlanNewParamsPricesNewPlanPackageWithAllocationPrice) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r PlanNewParamsPricesNewPlanPackageWithAllocationPrice) implementsPlanNewParamsPriceUnion() {}

// The cadence to bill for this price on.
type PlanNewParamsPricesNewPlanPackageWithAllocationPriceCadence string

const (
	PlanNewParamsPricesNewPlanPackageWithAllocationPriceCadenceAnnual     PlanNewParamsPricesNewPlanPackageWithAllocationPriceCadence = "annual"
	PlanNewParamsPricesNewPlanPackageWithAllocationPriceCadenceSemiAnnual PlanNewParamsPricesNewPlanPackageWithAllocationPriceCadence = "semi_annual"
	PlanNewParamsPricesNewPlanPackageWithAllocationPriceCadenceMonthly    PlanNewParamsPricesNewPlanPackageWithAllocationPriceCadence = "monthly"
	PlanNewParamsPricesNewPlanPackageWithAllocationPriceCadenceQuarterly  PlanNewParamsPricesNewPlanPackageWithAllocationPriceCadence = "quarterly"
	PlanNewParamsPricesNewPlanPackageWithAllocationPriceCadenceOneTime    PlanNewParamsPricesNewPlanPackageWithAllocationPriceCadence = "one_time"
	PlanNewParamsPricesNewPlanPackageWithAllocationPriceCadenceCustom     PlanNewParamsPricesNewPlanPackageWithAllocationPriceCadence = "custom"
)

func (r PlanNewParamsPricesNewPlanPackageWithAllocationPriceCadence) IsKnown() bool {
	switch r {
	case PlanNewParamsPricesNewPlanPackageWithAllocationPriceCadenceAnnual, PlanNewParamsPricesNewPlanPackageWithAllocationPriceCadenceSemiAnnual, PlanNewParamsPricesNewPlanPackageWithAllocationPriceCadenceMonthly, PlanNewParamsPricesNewPlanPackageWithAllocationPriceCadenceQuarterly, PlanNewParamsPricesNewPlanPackageWithAllocationPriceCadenceOneTime, PlanNewParamsPricesNewPlanPackageWithAllocationPriceCadenceCustom:
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

// For custom cadence: specifies the duration of the billing period in days or
// months.
type PlanNewParamsPricesNewPlanPackageWithAllocationPriceBillingCycleConfiguration struct {
	// The duration of the billing period.
	Duration param.Field[int64] `json:"duration,required"`
	// The unit of billing period duration.
	DurationUnit param.Field[PlanNewParamsPricesNewPlanPackageWithAllocationPriceBillingCycleConfigurationDurationUnit] `json:"duration_unit,required"`
}

func (r PlanNewParamsPricesNewPlanPackageWithAllocationPriceBillingCycleConfiguration) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// The unit of billing period duration.
type PlanNewParamsPricesNewPlanPackageWithAllocationPriceBillingCycleConfigurationDurationUnit string

const (
	PlanNewParamsPricesNewPlanPackageWithAllocationPriceBillingCycleConfigurationDurationUnitDay   PlanNewParamsPricesNewPlanPackageWithAllocationPriceBillingCycleConfigurationDurationUnit = "day"
	PlanNewParamsPricesNewPlanPackageWithAllocationPriceBillingCycleConfigurationDurationUnitMonth PlanNewParamsPricesNewPlanPackageWithAllocationPriceBillingCycleConfigurationDurationUnit = "month"
)

func (r PlanNewParamsPricesNewPlanPackageWithAllocationPriceBillingCycleConfigurationDurationUnit) IsKnown() bool {
	switch r {
	case PlanNewParamsPricesNewPlanPackageWithAllocationPriceBillingCycleConfigurationDurationUnitDay, PlanNewParamsPricesNewPlanPackageWithAllocationPriceBillingCycleConfigurationDurationUnitMonth:
		return true
	}
	return false
}

// Within each billing cycle, specifies the cadence at which invoices are produced.
// If unspecified, a single invoice is produced per billing cycle.
type PlanNewParamsPricesNewPlanPackageWithAllocationPriceInvoicingCycleConfiguration struct {
	// The duration of the billing period.
	Duration param.Field[int64] `json:"duration,required"`
	// The unit of billing period duration.
	DurationUnit param.Field[PlanNewParamsPricesNewPlanPackageWithAllocationPriceInvoicingCycleConfigurationDurationUnit] `json:"duration_unit,required"`
}

func (r PlanNewParamsPricesNewPlanPackageWithAllocationPriceInvoicingCycleConfiguration) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// The unit of billing period duration.
type PlanNewParamsPricesNewPlanPackageWithAllocationPriceInvoicingCycleConfigurationDurationUnit string

const (
	PlanNewParamsPricesNewPlanPackageWithAllocationPriceInvoicingCycleConfigurationDurationUnitDay   PlanNewParamsPricesNewPlanPackageWithAllocationPriceInvoicingCycleConfigurationDurationUnit = "day"
	PlanNewParamsPricesNewPlanPackageWithAllocationPriceInvoicingCycleConfigurationDurationUnitMonth PlanNewParamsPricesNewPlanPackageWithAllocationPriceInvoicingCycleConfigurationDurationUnit = "month"
)

func (r PlanNewParamsPricesNewPlanPackageWithAllocationPriceInvoicingCycleConfigurationDurationUnit) IsKnown() bool {
	switch r {
	case PlanNewParamsPricesNewPlanPackageWithAllocationPriceInvoicingCycleConfigurationDurationUnitDay, PlanNewParamsPricesNewPlanPackageWithAllocationPriceInvoicingCycleConfigurationDurationUnitMonth:
		return true
	}
	return false
}

type PlanNewParamsPricesNewPlanTierWithProrationPrice struct {
	// The cadence to bill for this price on.
	Cadence param.Field[PlanNewParamsPricesNewPlanTierWithProrationPriceCadence] `json:"cadence,required"`
	// The id of the item the plan will be associated with.
	ItemID    param.Field[string]                                                    `json:"item_id,required"`
	ModelType param.Field[PlanNewParamsPricesNewPlanTierWithProrationPriceModelType] `json:"model_type,required"`
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
	BillingCycleConfiguration param.Field[PlanNewParamsPricesNewPlanTierWithProrationPriceBillingCycleConfiguration] `json:"billing_cycle_configuration"`
	// The per unit conversion rate of the price currency to the invoicing currency.
	ConversionRate param.Field[float64] `json:"conversion_rate"`
	// An ISO 4217 currency string, or custom pricing unit identifier, in which this
	// price is billed.
	Currency param.Field[string] `json:"currency"`
	// An alias for the price.
	ExternalPriceID param.Field[string] `json:"external_price_id"`
	// If the Price represents a fixed cost, this represents the quantity of units
	// applied.
	FixedPriceQuantity param.Field[float64] `json:"fixed_price_quantity"`
	// The property used to group this price on an invoice
	InvoiceGroupingKey param.Field[string] `json:"invoice_grouping_key"`
	// Within each billing cycle, specifies the cadence at which invoices are produced.
	// If unspecified, a single invoice is produced per billing cycle.
	InvoicingCycleConfiguration param.Field[PlanNewParamsPricesNewPlanTierWithProrationPriceInvoicingCycleConfiguration] `json:"invoicing_cycle_configuration"`
	// User-specified key/value pairs for the resource. Individual keys can be removed
	// by setting the value to `null`, and the entire metadata mapping can be cleared
	// by setting `metadata` to `null`.
	Metadata param.Field[map[string]string] `json:"metadata"`
}

func (r PlanNewParamsPricesNewPlanTierWithProrationPrice) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r PlanNewParamsPricesNewPlanTierWithProrationPrice) implementsPlanNewParamsPriceUnion() {}

// The cadence to bill for this price on.
type PlanNewParamsPricesNewPlanTierWithProrationPriceCadence string

const (
	PlanNewParamsPricesNewPlanTierWithProrationPriceCadenceAnnual     PlanNewParamsPricesNewPlanTierWithProrationPriceCadence = "annual"
	PlanNewParamsPricesNewPlanTierWithProrationPriceCadenceSemiAnnual PlanNewParamsPricesNewPlanTierWithProrationPriceCadence = "semi_annual"
	PlanNewParamsPricesNewPlanTierWithProrationPriceCadenceMonthly    PlanNewParamsPricesNewPlanTierWithProrationPriceCadence = "monthly"
	PlanNewParamsPricesNewPlanTierWithProrationPriceCadenceQuarterly  PlanNewParamsPricesNewPlanTierWithProrationPriceCadence = "quarterly"
	PlanNewParamsPricesNewPlanTierWithProrationPriceCadenceOneTime    PlanNewParamsPricesNewPlanTierWithProrationPriceCadence = "one_time"
	PlanNewParamsPricesNewPlanTierWithProrationPriceCadenceCustom     PlanNewParamsPricesNewPlanTierWithProrationPriceCadence = "custom"
)

func (r PlanNewParamsPricesNewPlanTierWithProrationPriceCadence) IsKnown() bool {
	switch r {
	case PlanNewParamsPricesNewPlanTierWithProrationPriceCadenceAnnual, PlanNewParamsPricesNewPlanTierWithProrationPriceCadenceSemiAnnual, PlanNewParamsPricesNewPlanTierWithProrationPriceCadenceMonthly, PlanNewParamsPricesNewPlanTierWithProrationPriceCadenceQuarterly, PlanNewParamsPricesNewPlanTierWithProrationPriceCadenceOneTime, PlanNewParamsPricesNewPlanTierWithProrationPriceCadenceCustom:
		return true
	}
	return false
}

type PlanNewParamsPricesNewPlanTierWithProrationPriceModelType string

const (
	PlanNewParamsPricesNewPlanTierWithProrationPriceModelTypeTieredWithProration PlanNewParamsPricesNewPlanTierWithProrationPriceModelType = "tiered_with_proration"
)

func (r PlanNewParamsPricesNewPlanTierWithProrationPriceModelType) IsKnown() bool {
	switch r {
	case PlanNewParamsPricesNewPlanTierWithProrationPriceModelTypeTieredWithProration:
		return true
	}
	return false
}

// For custom cadence: specifies the duration of the billing period in days or
// months.
type PlanNewParamsPricesNewPlanTierWithProrationPriceBillingCycleConfiguration struct {
	// The duration of the billing period.
	Duration param.Field[int64] `json:"duration,required"`
	// The unit of billing period duration.
	DurationUnit param.Field[PlanNewParamsPricesNewPlanTierWithProrationPriceBillingCycleConfigurationDurationUnit] `json:"duration_unit,required"`
}

func (r PlanNewParamsPricesNewPlanTierWithProrationPriceBillingCycleConfiguration) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// The unit of billing period duration.
type PlanNewParamsPricesNewPlanTierWithProrationPriceBillingCycleConfigurationDurationUnit string

const (
	PlanNewParamsPricesNewPlanTierWithProrationPriceBillingCycleConfigurationDurationUnitDay   PlanNewParamsPricesNewPlanTierWithProrationPriceBillingCycleConfigurationDurationUnit = "day"
	PlanNewParamsPricesNewPlanTierWithProrationPriceBillingCycleConfigurationDurationUnitMonth PlanNewParamsPricesNewPlanTierWithProrationPriceBillingCycleConfigurationDurationUnit = "month"
)

func (r PlanNewParamsPricesNewPlanTierWithProrationPriceBillingCycleConfigurationDurationUnit) IsKnown() bool {
	switch r {
	case PlanNewParamsPricesNewPlanTierWithProrationPriceBillingCycleConfigurationDurationUnitDay, PlanNewParamsPricesNewPlanTierWithProrationPriceBillingCycleConfigurationDurationUnitMonth:
		return true
	}
	return false
}

// Within each billing cycle, specifies the cadence at which invoices are produced.
// If unspecified, a single invoice is produced per billing cycle.
type PlanNewParamsPricesNewPlanTierWithProrationPriceInvoicingCycleConfiguration struct {
	// The duration of the billing period.
	Duration param.Field[int64] `json:"duration,required"`
	// The unit of billing period duration.
	DurationUnit param.Field[PlanNewParamsPricesNewPlanTierWithProrationPriceInvoicingCycleConfigurationDurationUnit] `json:"duration_unit,required"`
}

func (r PlanNewParamsPricesNewPlanTierWithProrationPriceInvoicingCycleConfiguration) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// The unit of billing period duration.
type PlanNewParamsPricesNewPlanTierWithProrationPriceInvoicingCycleConfigurationDurationUnit string

const (
	PlanNewParamsPricesNewPlanTierWithProrationPriceInvoicingCycleConfigurationDurationUnitDay   PlanNewParamsPricesNewPlanTierWithProrationPriceInvoicingCycleConfigurationDurationUnit = "day"
	PlanNewParamsPricesNewPlanTierWithProrationPriceInvoicingCycleConfigurationDurationUnitMonth PlanNewParamsPricesNewPlanTierWithProrationPriceInvoicingCycleConfigurationDurationUnit = "month"
)

func (r PlanNewParamsPricesNewPlanTierWithProrationPriceInvoicingCycleConfigurationDurationUnit) IsKnown() bool {
	switch r {
	case PlanNewParamsPricesNewPlanTierWithProrationPriceInvoicingCycleConfigurationDurationUnitDay, PlanNewParamsPricesNewPlanTierWithProrationPriceInvoicingCycleConfigurationDurationUnitMonth:
		return true
	}
	return false
}

type PlanNewParamsPricesNewPlanUnitWithProrationPrice struct {
	// The cadence to bill for this price on.
	Cadence param.Field[PlanNewParamsPricesNewPlanUnitWithProrationPriceCadence] `json:"cadence,required"`
	// The id of the item the plan will be associated with.
	ItemID    param.Field[string]                                                    `json:"item_id,required"`
	ModelType param.Field[PlanNewParamsPricesNewPlanUnitWithProrationPriceModelType] `json:"model_type,required"`
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
	BillingCycleConfiguration param.Field[PlanNewParamsPricesNewPlanUnitWithProrationPriceBillingCycleConfiguration] `json:"billing_cycle_configuration"`
	// The per unit conversion rate of the price currency to the invoicing currency.
	ConversionRate param.Field[float64] `json:"conversion_rate"`
	// An ISO 4217 currency string, or custom pricing unit identifier, in which this
	// price is billed.
	Currency param.Field[string] `json:"currency"`
	// An alias for the price.
	ExternalPriceID param.Field[string] `json:"external_price_id"`
	// If the Price represents a fixed cost, this represents the quantity of units
	// applied.
	FixedPriceQuantity param.Field[float64] `json:"fixed_price_quantity"`
	// The property used to group this price on an invoice
	InvoiceGroupingKey param.Field[string] `json:"invoice_grouping_key"`
	// Within each billing cycle, specifies the cadence at which invoices are produced.
	// If unspecified, a single invoice is produced per billing cycle.
	InvoicingCycleConfiguration param.Field[PlanNewParamsPricesNewPlanUnitWithProrationPriceInvoicingCycleConfiguration] `json:"invoicing_cycle_configuration"`
	// User-specified key/value pairs for the resource. Individual keys can be removed
	// by setting the value to `null`, and the entire metadata mapping can be cleared
	// by setting `metadata` to `null`.
	Metadata param.Field[map[string]string] `json:"metadata"`
}

func (r PlanNewParamsPricesNewPlanUnitWithProrationPrice) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r PlanNewParamsPricesNewPlanUnitWithProrationPrice) implementsPlanNewParamsPriceUnion() {}

// The cadence to bill for this price on.
type PlanNewParamsPricesNewPlanUnitWithProrationPriceCadence string

const (
	PlanNewParamsPricesNewPlanUnitWithProrationPriceCadenceAnnual     PlanNewParamsPricesNewPlanUnitWithProrationPriceCadence = "annual"
	PlanNewParamsPricesNewPlanUnitWithProrationPriceCadenceSemiAnnual PlanNewParamsPricesNewPlanUnitWithProrationPriceCadence = "semi_annual"
	PlanNewParamsPricesNewPlanUnitWithProrationPriceCadenceMonthly    PlanNewParamsPricesNewPlanUnitWithProrationPriceCadence = "monthly"
	PlanNewParamsPricesNewPlanUnitWithProrationPriceCadenceQuarterly  PlanNewParamsPricesNewPlanUnitWithProrationPriceCadence = "quarterly"
	PlanNewParamsPricesNewPlanUnitWithProrationPriceCadenceOneTime    PlanNewParamsPricesNewPlanUnitWithProrationPriceCadence = "one_time"
	PlanNewParamsPricesNewPlanUnitWithProrationPriceCadenceCustom     PlanNewParamsPricesNewPlanUnitWithProrationPriceCadence = "custom"
)

func (r PlanNewParamsPricesNewPlanUnitWithProrationPriceCadence) IsKnown() bool {
	switch r {
	case PlanNewParamsPricesNewPlanUnitWithProrationPriceCadenceAnnual, PlanNewParamsPricesNewPlanUnitWithProrationPriceCadenceSemiAnnual, PlanNewParamsPricesNewPlanUnitWithProrationPriceCadenceMonthly, PlanNewParamsPricesNewPlanUnitWithProrationPriceCadenceQuarterly, PlanNewParamsPricesNewPlanUnitWithProrationPriceCadenceOneTime, PlanNewParamsPricesNewPlanUnitWithProrationPriceCadenceCustom:
		return true
	}
	return false
}

type PlanNewParamsPricesNewPlanUnitWithProrationPriceModelType string

const (
	PlanNewParamsPricesNewPlanUnitWithProrationPriceModelTypeUnitWithProration PlanNewParamsPricesNewPlanUnitWithProrationPriceModelType = "unit_with_proration"
)

func (r PlanNewParamsPricesNewPlanUnitWithProrationPriceModelType) IsKnown() bool {
	switch r {
	case PlanNewParamsPricesNewPlanUnitWithProrationPriceModelTypeUnitWithProration:
		return true
	}
	return false
}

// For custom cadence: specifies the duration of the billing period in days or
// months.
type PlanNewParamsPricesNewPlanUnitWithProrationPriceBillingCycleConfiguration struct {
	// The duration of the billing period.
	Duration param.Field[int64] `json:"duration,required"`
	// The unit of billing period duration.
	DurationUnit param.Field[PlanNewParamsPricesNewPlanUnitWithProrationPriceBillingCycleConfigurationDurationUnit] `json:"duration_unit,required"`
}

func (r PlanNewParamsPricesNewPlanUnitWithProrationPriceBillingCycleConfiguration) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// The unit of billing period duration.
type PlanNewParamsPricesNewPlanUnitWithProrationPriceBillingCycleConfigurationDurationUnit string

const (
	PlanNewParamsPricesNewPlanUnitWithProrationPriceBillingCycleConfigurationDurationUnitDay   PlanNewParamsPricesNewPlanUnitWithProrationPriceBillingCycleConfigurationDurationUnit = "day"
	PlanNewParamsPricesNewPlanUnitWithProrationPriceBillingCycleConfigurationDurationUnitMonth PlanNewParamsPricesNewPlanUnitWithProrationPriceBillingCycleConfigurationDurationUnit = "month"
)

func (r PlanNewParamsPricesNewPlanUnitWithProrationPriceBillingCycleConfigurationDurationUnit) IsKnown() bool {
	switch r {
	case PlanNewParamsPricesNewPlanUnitWithProrationPriceBillingCycleConfigurationDurationUnitDay, PlanNewParamsPricesNewPlanUnitWithProrationPriceBillingCycleConfigurationDurationUnitMonth:
		return true
	}
	return false
}

// Within each billing cycle, specifies the cadence at which invoices are produced.
// If unspecified, a single invoice is produced per billing cycle.
type PlanNewParamsPricesNewPlanUnitWithProrationPriceInvoicingCycleConfiguration struct {
	// The duration of the billing period.
	Duration param.Field[int64] `json:"duration,required"`
	// The unit of billing period duration.
	DurationUnit param.Field[PlanNewParamsPricesNewPlanUnitWithProrationPriceInvoicingCycleConfigurationDurationUnit] `json:"duration_unit,required"`
}

func (r PlanNewParamsPricesNewPlanUnitWithProrationPriceInvoicingCycleConfiguration) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// The unit of billing period duration.
type PlanNewParamsPricesNewPlanUnitWithProrationPriceInvoicingCycleConfigurationDurationUnit string

const (
	PlanNewParamsPricesNewPlanUnitWithProrationPriceInvoicingCycleConfigurationDurationUnitDay   PlanNewParamsPricesNewPlanUnitWithProrationPriceInvoicingCycleConfigurationDurationUnit = "day"
	PlanNewParamsPricesNewPlanUnitWithProrationPriceInvoicingCycleConfigurationDurationUnitMonth PlanNewParamsPricesNewPlanUnitWithProrationPriceInvoicingCycleConfigurationDurationUnit = "month"
)

func (r PlanNewParamsPricesNewPlanUnitWithProrationPriceInvoicingCycleConfigurationDurationUnit) IsKnown() bool {
	switch r {
	case PlanNewParamsPricesNewPlanUnitWithProrationPriceInvoicingCycleConfigurationDurationUnitDay, PlanNewParamsPricesNewPlanUnitWithProrationPriceInvoicingCycleConfigurationDurationUnitMonth:
		return true
	}
	return false
}

type PlanNewParamsPricesNewPlanGroupedAllocationPrice struct {
	// The cadence to bill for this price on.
	Cadence                 param.Field[PlanNewParamsPricesNewPlanGroupedAllocationPriceCadence] `json:"cadence,required"`
	GroupedAllocationConfig param.Field[map[string]interface{}]                                  `json:"grouped_allocation_config,required"`
	// The id of the item the plan will be associated with.
	ItemID    param.Field[string]                                                    `json:"item_id,required"`
	ModelType param.Field[PlanNewParamsPricesNewPlanGroupedAllocationPriceModelType] `json:"model_type,required"`
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
	BillingCycleConfiguration param.Field[PlanNewParamsPricesNewPlanGroupedAllocationPriceBillingCycleConfiguration] `json:"billing_cycle_configuration"`
	// The per unit conversion rate of the price currency to the invoicing currency.
	ConversionRate param.Field[float64] `json:"conversion_rate"`
	// An ISO 4217 currency string, or custom pricing unit identifier, in which this
	// price is billed.
	Currency param.Field[string] `json:"currency"`
	// An alias for the price.
	ExternalPriceID param.Field[string] `json:"external_price_id"`
	// If the Price represents a fixed cost, this represents the quantity of units
	// applied.
	FixedPriceQuantity param.Field[float64] `json:"fixed_price_quantity"`
	// The property used to group this price on an invoice
	InvoiceGroupingKey param.Field[string] `json:"invoice_grouping_key"`
	// Within each billing cycle, specifies the cadence at which invoices are produced.
	// If unspecified, a single invoice is produced per billing cycle.
	InvoicingCycleConfiguration param.Field[PlanNewParamsPricesNewPlanGroupedAllocationPriceInvoicingCycleConfiguration] `json:"invoicing_cycle_configuration"`
	// User-specified key/value pairs for the resource. Individual keys can be removed
	// by setting the value to `null`, and the entire metadata mapping can be cleared
	// by setting `metadata` to `null`.
	Metadata param.Field[map[string]string] `json:"metadata"`
}

func (r PlanNewParamsPricesNewPlanGroupedAllocationPrice) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r PlanNewParamsPricesNewPlanGroupedAllocationPrice) implementsPlanNewParamsPriceUnion() {}

// The cadence to bill for this price on.
type PlanNewParamsPricesNewPlanGroupedAllocationPriceCadence string

const (
	PlanNewParamsPricesNewPlanGroupedAllocationPriceCadenceAnnual     PlanNewParamsPricesNewPlanGroupedAllocationPriceCadence = "annual"
	PlanNewParamsPricesNewPlanGroupedAllocationPriceCadenceSemiAnnual PlanNewParamsPricesNewPlanGroupedAllocationPriceCadence = "semi_annual"
	PlanNewParamsPricesNewPlanGroupedAllocationPriceCadenceMonthly    PlanNewParamsPricesNewPlanGroupedAllocationPriceCadence = "monthly"
	PlanNewParamsPricesNewPlanGroupedAllocationPriceCadenceQuarterly  PlanNewParamsPricesNewPlanGroupedAllocationPriceCadence = "quarterly"
	PlanNewParamsPricesNewPlanGroupedAllocationPriceCadenceOneTime    PlanNewParamsPricesNewPlanGroupedAllocationPriceCadence = "one_time"
	PlanNewParamsPricesNewPlanGroupedAllocationPriceCadenceCustom     PlanNewParamsPricesNewPlanGroupedAllocationPriceCadence = "custom"
)

func (r PlanNewParamsPricesNewPlanGroupedAllocationPriceCadence) IsKnown() bool {
	switch r {
	case PlanNewParamsPricesNewPlanGroupedAllocationPriceCadenceAnnual, PlanNewParamsPricesNewPlanGroupedAllocationPriceCadenceSemiAnnual, PlanNewParamsPricesNewPlanGroupedAllocationPriceCadenceMonthly, PlanNewParamsPricesNewPlanGroupedAllocationPriceCadenceQuarterly, PlanNewParamsPricesNewPlanGroupedAllocationPriceCadenceOneTime, PlanNewParamsPricesNewPlanGroupedAllocationPriceCadenceCustom:
		return true
	}
	return false
}

type PlanNewParamsPricesNewPlanGroupedAllocationPriceModelType string

const (
	PlanNewParamsPricesNewPlanGroupedAllocationPriceModelTypeGroupedAllocation PlanNewParamsPricesNewPlanGroupedAllocationPriceModelType = "grouped_allocation"
)

func (r PlanNewParamsPricesNewPlanGroupedAllocationPriceModelType) IsKnown() bool {
	switch r {
	case PlanNewParamsPricesNewPlanGroupedAllocationPriceModelTypeGroupedAllocation:
		return true
	}
	return false
}

// For custom cadence: specifies the duration of the billing period in days or
// months.
type PlanNewParamsPricesNewPlanGroupedAllocationPriceBillingCycleConfiguration struct {
	// The duration of the billing period.
	Duration param.Field[int64] `json:"duration,required"`
	// The unit of billing period duration.
	DurationUnit param.Field[PlanNewParamsPricesNewPlanGroupedAllocationPriceBillingCycleConfigurationDurationUnit] `json:"duration_unit,required"`
}

func (r PlanNewParamsPricesNewPlanGroupedAllocationPriceBillingCycleConfiguration) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// The unit of billing period duration.
type PlanNewParamsPricesNewPlanGroupedAllocationPriceBillingCycleConfigurationDurationUnit string

const (
	PlanNewParamsPricesNewPlanGroupedAllocationPriceBillingCycleConfigurationDurationUnitDay   PlanNewParamsPricesNewPlanGroupedAllocationPriceBillingCycleConfigurationDurationUnit = "day"
	PlanNewParamsPricesNewPlanGroupedAllocationPriceBillingCycleConfigurationDurationUnitMonth PlanNewParamsPricesNewPlanGroupedAllocationPriceBillingCycleConfigurationDurationUnit = "month"
)

func (r PlanNewParamsPricesNewPlanGroupedAllocationPriceBillingCycleConfigurationDurationUnit) IsKnown() bool {
	switch r {
	case PlanNewParamsPricesNewPlanGroupedAllocationPriceBillingCycleConfigurationDurationUnitDay, PlanNewParamsPricesNewPlanGroupedAllocationPriceBillingCycleConfigurationDurationUnitMonth:
		return true
	}
	return false
}

// Within each billing cycle, specifies the cadence at which invoices are produced.
// If unspecified, a single invoice is produced per billing cycle.
type PlanNewParamsPricesNewPlanGroupedAllocationPriceInvoicingCycleConfiguration struct {
	// The duration of the billing period.
	Duration param.Field[int64] `json:"duration,required"`
	// The unit of billing period duration.
	DurationUnit param.Field[PlanNewParamsPricesNewPlanGroupedAllocationPriceInvoicingCycleConfigurationDurationUnit] `json:"duration_unit,required"`
}

func (r PlanNewParamsPricesNewPlanGroupedAllocationPriceInvoicingCycleConfiguration) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// The unit of billing period duration.
type PlanNewParamsPricesNewPlanGroupedAllocationPriceInvoicingCycleConfigurationDurationUnit string

const (
	PlanNewParamsPricesNewPlanGroupedAllocationPriceInvoicingCycleConfigurationDurationUnitDay   PlanNewParamsPricesNewPlanGroupedAllocationPriceInvoicingCycleConfigurationDurationUnit = "day"
	PlanNewParamsPricesNewPlanGroupedAllocationPriceInvoicingCycleConfigurationDurationUnitMonth PlanNewParamsPricesNewPlanGroupedAllocationPriceInvoicingCycleConfigurationDurationUnit = "month"
)

func (r PlanNewParamsPricesNewPlanGroupedAllocationPriceInvoicingCycleConfigurationDurationUnit) IsKnown() bool {
	switch r {
	case PlanNewParamsPricesNewPlanGroupedAllocationPriceInvoicingCycleConfigurationDurationUnitDay, PlanNewParamsPricesNewPlanGroupedAllocationPriceInvoicingCycleConfigurationDurationUnitMonth:
		return true
	}
	return false
}

type PlanNewParamsPricesNewPlanGroupedWithProratedMinimumPrice struct {
	// The cadence to bill for this price on.
	Cadence                          param.Field[PlanNewParamsPricesNewPlanGroupedWithProratedMinimumPriceCadence] `json:"cadence,required"`
	GroupedWithProratedMinimumConfig param.Field[map[string]interface{}]                                           `json:"grouped_with_prorated_minimum_config,required"`
	// The id of the item the plan will be associated with.
	ItemID    param.Field[string]                                                             `json:"item_id,required"`
	ModelType param.Field[PlanNewParamsPricesNewPlanGroupedWithProratedMinimumPriceModelType] `json:"model_type,required"`
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
	BillingCycleConfiguration param.Field[PlanNewParamsPricesNewPlanGroupedWithProratedMinimumPriceBillingCycleConfiguration] `json:"billing_cycle_configuration"`
	// The per unit conversion rate of the price currency to the invoicing currency.
	ConversionRate param.Field[float64] `json:"conversion_rate"`
	// An ISO 4217 currency string, or custom pricing unit identifier, in which this
	// price is billed.
	Currency param.Field[string] `json:"currency"`
	// An alias for the price.
	ExternalPriceID param.Field[string] `json:"external_price_id"`
	// If the Price represents a fixed cost, this represents the quantity of units
	// applied.
	FixedPriceQuantity param.Field[float64] `json:"fixed_price_quantity"`
	// The property used to group this price on an invoice
	InvoiceGroupingKey param.Field[string] `json:"invoice_grouping_key"`
	// Within each billing cycle, specifies the cadence at which invoices are produced.
	// If unspecified, a single invoice is produced per billing cycle.
	InvoicingCycleConfiguration param.Field[PlanNewParamsPricesNewPlanGroupedWithProratedMinimumPriceInvoicingCycleConfiguration] `json:"invoicing_cycle_configuration"`
	// User-specified key/value pairs for the resource. Individual keys can be removed
	// by setting the value to `null`, and the entire metadata mapping can be cleared
	// by setting `metadata` to `null`.
	Metadata param.Field[map[string]string] `json:"metadata"`
}

func (r PlanNewParamsPricesNewPlanGroupedWithProratedMinimumPrice) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r PlanNewParamsPricesNewPlanGroupedWithProratedMinimumPrice) implementsPlanNewParamsPriceUnion() {
}

// The cadence to bill for this price on.
type PlanNewParamsPricesNewPlanGroupedWithProratedMinimumPriceCadence string

const (
	PlanNewParamsPricesNewPlanGroupedWithProratedMinimumPriceCadenceAnnual     PlanNewParamsPricesNewPlanGroupedWithProratedMinimumPriceCadence = "annual"
	PlanNewParamsPricesNewPlanGroupedWithProratedMinimumPriceCadenceSemiAnnual PlanNewParamsPricesNewPlanGroupedWithProratedMinimumPriceCadence = "semi_annual"
	PlanNewParamsPricesNewPlanGroupedWithProratedMinimumPriceCadenceMonthly    PlanNewParamsPricesNewPlanGroupedWithProratedMinimumPriceCadence = "monthly"
	PlanNewParamsPricesNewPlanGroupedWithProratedMinimumPriceCadenceQuarterly  PlanNewParamsPricesNewPlanGroupedWithProratedMinimumPriceCadence = "quarterly"
	PlanNewParamsPricesNewPlanGroupedWithProratedMinimumPriceCadenceOneTime    PlanNewParamsPricesNewPlanGroupedWithProratedMinimumPriceCadence = "one_time"
	PlanNewParamsPricesNewPlanGroupedWithProratedMinimumPriceCadenceCustom     PlanNewParamsPricesNewPlanGroupedWithProratedMinimumPriceCadence = "custom"
)

func (r PlanNewParamsPricesNewPlanGroupedWithProratedMinimumPriceCadence) IsKnown() bool {
	switch r {
	case PlanNewParamsPricesNewPlanGroupedWithProratedMinimumPriceCadenceAnnual, PlanNewParamsPricesNewPlanGroupedWithProratedMinimumPriceCadenceSemiAnnual, PlanNewParamsPricesNewPlanGroupedWithProratedMinimumPriceCadenceMonthly, PlanNewParamsPricesNewPlanGroupedWithProratedMinimumPriceCadenceQuarterly, PlanNewParamsPricesNewPlanGroupedWithProratedMinimumPriceCadenceOneTime, PlanNewParamsPricesNewPlanGroupedWithProratedMinimumPriceCadenceCustom:
		return true
	}
	return false
}

type PlanNewParamsPricesNewPlanGroupedWithProratedMinimumPriceModelType string

const (
	PlanNewParamsPricesNewPlanGroupedWithProratedMinimumPriceModelTypeGroupedWithProratedMinimum PlanNewParamsPricesNewPlanGroupedWithProratedMinimumPriceModelType = "grouped_with_prorated_minimum"
)

func (r PlanNewParamsPricesNewPlanGroupedWithProratedMinimumPriceModelType) IsKnown() bool {
	switch r {
	case PlanNewParamsPricesNewPlanGroupedWithProratedMinimumPriceModelTypeGroupedWithProratedMinimum:
		return true
	}
	return false
}

// For custom cadence: specifies the duration of the billing period in days or
// months.
type PlanNewParamsPricesNewPlanGroupedWithProratedMinimumPriceBillingCycleConfiguration struct {
	// The duration of the billing period.
	Duration param.Field[int64] `json:"duration,required"`
	// The unit of billing period duration.
	DurationUnit param.Field[PlanNewParamsPricesNewPlanGroupedWithProratedMinimumPriceBillingCycleConfigurationDurationUnit] `json:"duration_unit,required"`
}

func (r PlanNewParamsPricesNewPlanGroupedWithProratedMinimumPriceBillingCycleConfiguration) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// The unit of billing period duration.
type PlanNewParamsPricesNewPlanGroupedWithProratedMinimumPriceBillingCycleConfigurationDurationUnit string

const (
	PlanNewParamsPricesNewPlanGroupedWithProratedMinimumPriceBillingCycleConfigurationDurationUnitDay   PlanNewParamsPricesNewPlanGroupedWithProratedMinimumPriceBillingCycleConfigurationDurationUnit = "day"
	PlanNewParamsPricesNewPlanGroupedWithProratedMinimumPriceBillingCycleConfigurationDurationUnitMonth PlanNewParamsPricesNewPlanGroupedWithProratedMinimumPriceBillingCycleConfigurationDurationUnit = "month"
)

func (r PlanNewParamsPricesNewPlanGroupedWithProratedMinimumPriceBillingCycleConfigurationDurationUnit) IsKnown() bool {
	switch r {
	case PlanNewParamsPricesNewPlanGroupedWithProratedMinimumPriceBillingCycleConfigurationDurationUnitDay, PlanNewParamsPricesNewPlanGroupedWithProratedMinimumPriceBillingCycleConfigurationDurationUnitMonth:
		return true
	}
	return false
}

// Within each billing cycle, specifies the cadence at which invoices are produced.
// If unspecified, a single invoice is produced per billing cycle.
type PlanNewParamsPricesNewPlanGroupedWithProratedMinimumPriceInvoicingCycleConfiguration struct {
	// The duration of the billing period.
	Duration param.Field[int64] `json:"duration,required"`
	// The unit of billing period duration.
	DurationUnit param.Field[PlanNewParamsPricesNewPlanGroupedWithProratedMinimumPriceInvoicingCycleConfigurationDurationUnit] `json:"duration_unit,required"`
}

func (r PlanNewParamsPricesNewPlanGroupedWithProratedMinimumPriceInvoicingCycleConfiguration) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// The unit of billing period duration.
type PlanNewParamsPricesNewPlanGroupedWithProratedMinimumPriceInvoicingCycleConfigurationDurationUnit string

const (
	PlanNewParamsPricesNewPlanGroupedWithProratedMinimumPriceInvoicingCycleConfigurationDurationUnitDay   PlanNewParamsPricesNewPlanGroupedWithProratedMinimumPriceInvoicingCycleConfigurationDurationUnit = "day"
	PlanNewParamsPricesNewPlanGroupedWithProratedMinimumPriceInvoicingCycleConfigurationDurationUnitMonth PlanNewParamsPricesNewPlanGroupedWithProratedMinimumPriceInvoicingCycleConfigurationDurationUnit = "month"
)

func (r PlanNewParamsPricesNewPlanGroupedWithProratedMinimumPriceInvoicingCycleConfigurationDurationUnit) IsKnown() bool {
	switch r {
	case PlanNewParamsPricesNewPlanGroupedWithProratedMinimumPriceInvoicingCycleConfigurationDurationUnitDay, PlanNewParamsPricesNewPlanGroupedWithProratedMinimumPriceInvoicingCycleConfigurationDurationUnitMonth:
		return true
	}
	return false
}

type PlanNewParamsPricesNewPlanGroupedWithMeteredMinimumPrice struct {
	// The cadence to bill for this price on.
	Cadence                         param.Field[PlanNewParamsPricesNewPlanGroupedWithMeteredMinimumPriceCadence] `json:"cadence,required"`
	GroupedWithMeteredMinimumConfig param.Field[map[string]interface{}]                                          `json:"grouped_with_metered_minimum_config,required"`
	// The id of the item the plan will be associated with.
	ItemID    param.Field[string]                                                            `json:"item_id,required"`
	ModelType param.Field[PlanNewParamsPricesNewPlanGroupedWithMeteredMinimumPriceModelType] `json:"model_type,required"`
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
	BillingCycleConfiguration param.Field[PlanNewParamsPricesNewPlanGroupedWithMeteredMinimumPriceBillingCycleConfiguration] `json:"billing_cycle_configuration"`
	// The per unit conversion rate of the price currency to the invoicing currency.
	ConversionRate param.Field[float64] `json:"conversion_rate"`
	// An ISO 4217 currency string, or custom pricing unit identifier, in which this
	// price is billed.
	Currency param.Field[string] `json:"currency"`
	// An alias for the price.
	ExternalPriceID param.Field[string] `json:"external_price_id"`
	// If the Price represents a fixed cost, this represents the quantity of units
	// applied.
	FixedPriceQuantity param.Field[float64] `json:"fixed_price_quantity"`
	// The property used to group this price on an invoice
	InvoiceGroupingKey param.Field[string] `json:"invoice_grouping_key"`
	// Within each billing cycle, specifies the cadence at which invoices are produced.
	// If unspecified, a single invoice is produced per billing cycle.
	InvoicingCycleConfiguration param.Field[PlanNewParamsPricesNewPlanGroupedWithMeteredMinimumPriceInvoicingCycleConfiguration] `json:"invoicing_cycle_configuration"`
	// User-specified key/value pairs for the resource. Individual keys can be removed
	// by setting the value to `null`, and the entire metadata mapping can be cleared
	// by setting `metadata` to `null`.
	Metadata param.Field[map[string]string] `json:"metadata"`
}

func (r PlanNewParamsPricesNewPlanGroupedWithMeteredMinimumPrice) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r PlanNewParamsPricesNewPlanGroupedWithMeteredMinimumPrice) implementsPlanNewParamsPriceUnion() {
}

// The cadence to bill for this price on.
type PlanNewParamsPricesNewPlanGroupedWithMeteredMinimumPriceCadence string

const (
	PlanNewParamsPricesNewPlanGroupedWithMeteredMinimumPriceCadenceAnnual     PlanNewParamsPricesNewPlanGroupedWithMeteredMinimumPriceCadence = "annual"
	PlanNewParamsPricesNewPlanGroupedWithMeteredMinimumPriceCadenceSemiAnnual PlanNewParamsPricesNewPlanGroupedWithMeteredMinimumPriceCadence = "semi_annual"
	PlanNewParamsPricesNewPlanGroupedWithMeteredMinimumPriceCadenceMonthly    PlanNewParamsPricesNewPlanGroupedWithMeteredMinimumPriceCadence = "monthly"
	PlanNewParamsPricesNewPlanGroupedWithMeteredMinimumPriceCadenceQuarterly  PlanNewParamsPricesNewPlanGroupedWithMeteredMinimumPriceCadence = "quarterly"
	PlanNewParamsPricesNewPlanGroupedWithMeteredMinimumPriceCadenceOneTime    PlanNewParamsPricesNewPlanGroupedWithMeteredMinimumPriceCadence = "one_time"
	PlanNewParamsPricesNewPlanGroupedWithMeteredMinimumPriceCadenceCustom     PlanNewParamsPricesNewPlanGroupedWithMeteredMinimumPriceCadence = "custom"
)

func (r PlanNewParamsPricesNewPlanGroupedWithMeteredMinimumPriceCadence) IsKnown() bool {
	switch r {
	case PlanNewParamsPricesNewPlanGroupedWithMeteredMinimumPriceCadenceAnnual, PlanNewParamsPricesNewPlanGroupedWithMeteredMinimumPriceCadenceSemiAnnual, PlanNewParamsPricesNewPlanGroupedWithMeteredMinimumPriceCadenceMonthly, PlanNewParamsPricesNewPlanGroupedWithMeteredMinimumPriceCadenceQuarterly, PlanNewParamsPricesNewPlanGroupedWithMeteredMinimumPriceCadenceOneTime, PlanNewParamsPricesNewPlanGroupedWithMeteredMinimumPriceCadenceCustom:
		return true
	}
	return false
}

type PlanNewParamsPricesNewPlanGroupedWithMeteredMinimumPriceModelType string

const (
	PlanNewParamsPricesNewPlanGroupedWithMeteredMinimumPriceModelTypeGroupedWithMeteredMinimum PlanNewParamsPricesNewPlanGroupedWithMeteredMinimumPriceModelType = "grouped_with_metered_minimum"
)

func (r PlanNewParamsPricesNewPlanGroupedWithMeteredMinimumPriceModelType) IsKnown() bool {
	switch r {
	case PlanNewParamsPricesNewPlanGroupedWithMeteredMinimumPriceModelTypeGroupedWithMeteredMinimum:
		return true
	}
	return false
}

// For custom cadence: specifies the duration of the billing period in days or
// months.
type PlanNewParamsPricesNewPlanGroupedWithMeteredMinimumPriceBillingCycleConfiguration struct {
	// The duration of the billing period.
	Duration param.Field[int64] `json:"duration,required"`
	// The unit of billing period duration.
	DurationUnit param.Field[PlanNewParamsPricesNewPlanGroupedWithMeteredMinimumPriceBillingCycleConfigurationDurationUnit] `json:"duration_unit,required"`
}

func (r PlanNewParamsPricesNewPlanGroupedWithMeteredMinimumPriceBillingCycleConfiguration) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// The unit of billing period duration.
type PlanNewParamsPricesNewPlanGroupedWithMeteredMinimumPriceBillingCycleConfigurationDurationUnit string

const (
	PlanNewParamsPricesNewPlanGroupedWithMeteredMinimumPriceBillingCycleConfigurationDurationUnitDay   PlanNewParamsPricesNewPlanGroupedWithMeteredMinimumPriceBillingCycleConfigurationDurationUnit = "day"
	PlanNewParamsPricesNewPlanGroupedWithMeteredMinimumPriceBillingCycleConfigurationDurationUnitMonth PlanNewParamsPricesNewPlanGroupedWithMeteredMinimumPriceBillingCycleConfigurationDurationUnit = "month"
)

func (r PlanNewParamsPricesNewPlanGroupedWithMeteredMinimumPriceBillingCycleConfigurationDurationUnit) IsKnown() bool {
	switch r {
	case PlanNewParamsPricesNewPlanGroupedWithMeteredMinimumPriceBillingCycleConfigurationDurationUnitDay, PlanNewParamsPricesNewPlanGroupedWithMeteredMinimumPriceBillingCycleConfigurationDurationUnitMonth:
		return true
	}
	return false
}

// Within each billing cycle, specifies the cadence at which invoices are produced.
// If unspecified, a single invoice is produced per billing cycle.
type PlanNewParamsPricesNewPlanGroupedWithMeteredMinimumPriceInvoicingCycleConfiguration struct {
	// The duration of the billing period.
	Duration param.Field[int64] `json:"duration,required"`
	// The unit of billing period duration.
	DurationUnit param.Field[PlanNewParamsPricesNewPlanGroupedWithMeteredMinimumPriceInvoicingCycleConfigurationDurationUnit] `json:"duration_unit,required"`
}

func (r PlanNewParamsPricesNewPlanGroupedWithMeteredMinimumPriceInvoicingCycleConfiguration) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// The unit of billing period duration.
type PlanNewParamsPricesNewPlanGroupedWithMeteredMinimumPriceInvoicingCycleConfigurationDurationUnit string

const (
	PlanNewParamsPricesNewPlanGroupedWithMeteredMinimumPriceInvoicingCycleConfigurationDurationUnitDay   PlanNewParamsPricesNewPlanGroupedWithMeteredMinimumPriceInvoicingCycleConfigurationDurationUnit = "day"
	PlanNewParamsPricesNewPlanGroupedWithMeteredMinimumPriceInvoicingCycleConfigurationDurationUnitMonth PlanNewParamsPricesNewPlanGroupedWithMeteredMinimumPriceInvoicingCycleConfigurationDurationUnit = "month"
)

func (r PlanNewParamsPricesNewPlanGroupedWithMeteredMinimumPriceInvoicingCycleConfigurationDurationUnit) IsKnown() bool {
	switch r {
	case PlanNewParamsPricesNewPlanGroupedWithMeteredMinimumPriceInvoicingCycleConfigurationDurationUnitDay, PlanNewParamsPricesNewPlanGroupedWithMeteredMinimumPriceInvoicingCycleConfigurationDurationUnitMonth:
		return true
	}
	return false
}

type PlanNewParamsPricesNewPlanMatrixWithDisplayNamePrice struct {
	// The cadence to bill for this price on.
	Cadence param.Field[PlanNewParamsPricesNewPlanMatrixWithDisplayNamePriceCadence] `json:"cadence,required"`
	// The id of the item the plan will be associated with.
	ItemID                      param.Field[string]                                                        `json:"item_id,required"`
	MatrixWithDisplayNameConfig param.Field[map[string]interface{}]                                        `json:"matrix_with_display_name_config,required"`
	ModelType                   param.Field[PlanNewParamsPricesNewPlanMatrixWithDisplayNamePriceModelType] `json:"model_type,required"`
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
	BillingCycleConfiguration param.Field[PlanNewParamsPricesNewPlanMatrixWithDisplayNamePriceBillingCycleConfiguration] `json:"billing_cycle_configuration"`
	// The per unit conversion rate of the price currency to the invoicing currency.
	ConversionRate param.Field[float64] `json:"conversion_rate"`
	// An ISO 4217 currency string, or custom pricing unit identifier, in which this
	// price is billed.
	Currency param.Field[string] `json:"currency"`
	// An alias for the price.
	ExternalPriceID param.Field[string] `json:"external_price_id"`
	// If the Price represents a fixed cost, this represents the quantity of units
	// applied.
	FixedPriceQuantity param.Field[float64] `json:"fixed_price_quantity"`
	// The property used to group this price on an invoice
	InvoiceGroupingKey param.Field[string] `json:"invoice_grouping_key"`
	// Within each billing cycle, specifies the cadence at which invoices are produced.
	// If unspecified, a single invoice is produced per billing cycle.
	InvoicingCycleConfiguration param.Field[PlanNewParamsPricesNewPlanMatrixWithDisplayNamePriceInvoicingCycleConfiguration] `json:"invoicing_cycle_configuration"`
	// User-specified key/value pairs for the resource. Individual keys can be removed
	// by setting the value to `null`, and the entire metadata mapping can be cleared
	// by setting `metadata` to `null`.
	Metadata param.Field[map[string]string] `json:"metadata"`
}

func (r PlanNewParamsPricesNewPlanMatrixWithDisplayNamePrice) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r PlanNewParamsPricesNewPlanMatrixWithDisplayNamePrice) implementsPlanNewParamsPriceUnion() {}

// The cadence to bill for this price on.
type PlanNewParamsPricesNewPlanMatrixWithDisplayNamePriceCadence string

const (
	PlanNewParamsPricesNewPlanMatrixWithDisplayNamePriceCadenceAnnual     PlanNewParamsPricesNewPlanMatrixWithDisplayNamePriceCadence = "annual"
	PlanNewParamsPricesNewPlanMatrixWithDisplayNamePriceCadenceSemiAnnual PlanNewParamsPricesNewPlanMatrixWithDisplayNamePriceCadence = "semi_annual"
	PlanNewParamsPricesNewPlanMatrixWithDisplayNamePriceCadenceMonthly    PlanNewParamsPricesNewPlanMatrixWithDisplayNamePriceCadence = "monthly"
	PlanNewParamsPricesNewPlanMatrixWithDisplayNamePriceCadenceQuarterly  PlanNewParamsPricesNewPlanMatrixWithDisplayNamePriceCadence = "quarterly"
	PlanNewParamsPricesNewPlanMatrixWithDisplayNamePriceCadenceOneTime    PlanNewParamsPricesNewPlanMatrixWithDisplayNamePriceCadence = "one_time"
	PlanNewParamsPricesNewPlanMatrixWithDisplayNamePriceCadenceCustom     PlanNewParamsPricesNewPlanMatrixWithDisplayNamePriceCadence = "custom"
)

func (r PlanNewParamsPricesNewPlanMatrixWithDisplayNamePriceCadence) IsKnown() bool {
	switch r {
	case PlanNewParamsPricesNewPlanMatrixWithDisplayNamePriceCadenceAnnual, PlanNewParamsPricesNewPlanMatrixWithDisplayNamePriceCadenceSemiAnnual, PlanNewParamsPricesNewPlanMatrixWithDisplayNamePriceCadenceMonthly, PlanNewParamsPricesNewPlanMatrixWithDisplayNamePriceCadenceQuarterly, PlanNewParamsPricesNewPlanMatrixWithDisplayNamePriceCadenceOneTime, PlanNewParamsPricesNewPlanMatrixWithDisplayNamePriceCadenceCustom:
		return true
	}
	return false
}

type PlanNewParamsPricesNewPlanMatrixWithDisplayNamePriceModelType string

const (
	PlanNewParamsPricesNewPlanMatrixWithDisplayNamePriceModelTypeMatrixWithDisplayName PlanNewParamsPricesNewPlanMatrixWithDisplayNamePriceModelType = "matrix_with_display_name"
)

func (r PlanNewParamsPricesNewPlanMatrixWithDisplayNamePriceModelType) IsKnown() bool {
	switch r {
	case PlanNewParamsPricesNewPlanMatrixWithDisplayNamePriceModelTypeMatrixWithDisplayName:
		return true
	}
	return false
}

// For custom cadence: specifies the duration of the billing period in days or
// months.
type PlanNewParamsPricesNewPlanMatrixWithDisplayNamePriceBillingCycleConfiguration struct {
	// The duration of the billing period.
	Duration param.Field[int64] `json:"duration,required"`
	// The unit of billing period duration.
	DurationUnit param.Field[PlanNewParamsPricesNewPlanMatrixWithDisplayNamePriceBillingCycleConfigurationDurationUnit] `json:"duration_unit,required"`
}

func (r PlanNewParamsPricesNewPlanMatrixWithDisplayNamePriceBillingCycleConfiguration) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// The unit of billing period duration.
type PlanNewParamsPricesNewPlanMatrixWithDisplayNamePriceBillingCycleConfigurationDurationUnit string

const (
	PlanNewParamsPricesNewPlanMatrixWithDisplayNamePriceBillingCycleConfigurationDurationUnitDay   PlanNewParamsPricesNewPlanMatrixWithDisplayNamePriceBillingCycleConfigurationDurationUnit = "day"
	PlanNewParamsPricesNewPlanMatrixWithDisplayNamePriceBillingCycleConfigurationDurationUnitMonth PlanNewParamsPricesNewPlanMatrixWithDisplayNamePriceBillingCycleConfigurationDurationUnit = "month"
)

func (r PlanNewParamsPricesNewPlanMatrixWithDisplayNamePriceBillingCycleConfigurationDurationUnit) IsKnown() bool {
	switch r {
	case PlanNewParamsPricesNewPlanMatrixWithDisplayNamePriceBillingCycleConfigurationDurationUnitDay, PlanNewParamsPricesNewPlanMatrixWithDisplayNamePriceBillingCycleConfigurationDurationUnitMonth:
		return true
	}
	return false
}

// Within each billing cycle, specifies the cadence at which invoices are produced.
// If unspecified, a single invoice is produced per billing cycle.
type PlanNewParamsPricesNewPlanMatrixWithDisplayNamePriceInvoicingCycleConfiguration struct {
	// The duration of the billing period.
	Duration param.Field[int64] `json:"duration,required"`
	// The unit of billing period duration.
	DurationUnit param.Field[PlanNewParamsPricesNewPlanMatrixWithDisplayNamePriceInvoicingCycleConfigurationDurationUnit] `json:"duration_unit,required"`
}

func (r PlanNewParamsPricesNewPlanMatrixWithDisplayNamePriceInvoicingCycleConfiguration) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// The unit of billing period duration.
type PlanNewParamsPricesNewPlanMatrixWithDisplayNamePriceInvoicingCycleConfigurationDurationUnit string

const (
	PlanNewParamsPricesNewPlanMatrixWithDisplayNamePriceInvoicingCycleConfigurationDurationUnitDay   PlanNewParamsPricesNewPlanMatrixWithDisplayNamePriceInvoicingCycleConfigurationDurationUnit = "day"
	PlanNewParamsPricesNewPlanMatrixWithDisplayNamePriceInvoicingCycleConfigurationDurationUnitMonth PlanNewParamsPricesNewPlanMatrixWithDisplayNamePriceInvoicingCycleConfigurationDurationUnit = "month"
)

func (r PlanNewParamsPricesNewPlanMatrixWithDisplayNamePriceInvoicingCycleConfigurationDurationUnit) IsKnown() bool {
	switch r {
	case PlanNewParamsPricesNewPlanMatrixWithDisplayNamePriceInvoicingCycleConfigurationDurationUnitDay, PlanNewParamsPricesNewPlanMatrixWithDisplayNamePriceInvoicingCycleConfigurationDurationUnitMonth:
		return true
	}
	return false
}

type PlanNewParamsPricesNewPlanBulkWithProrationPrice struct {
	BulkWithProrationConfig param.Field[map[string]interface{}] `json:"bulk_with_proration_config,required"`
	// The cadence to bill for this price on.
	Cadence param.Field[PlanNewParamsPricesNewPlanBulkWithProrationPriceCadence] `json:"cadence,required"`
	// The id of the item the plan will be associated with.
	ItemID    param.Field[string]                                                    `json:"item_id,required"`
	ModelType param.Field[PlanNewParamsPricesNewPlanBulkWithProrationPriceModelType] `json:"model_type,required"`
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
	BillingCycleConfiguration param.Field[PlanNewParamsPricesNewPlanBulkWithProrationPriceBillingCycleConfiguration] `json:"billing_cycle_configuration"`
	// The per unit conversion rate of the price currency to the invoicing currency.
	ConversionRate param.Field[float64] `json:"conversion_rate"`
	// An ISO 4217 currency string, or custom pricing unit identifier, in which this
	// price is billed.
	Currency param.Field[string] `json:"currency"`
	// An alias for the price.
	ExternalPriceID param.Field[string] `json:"external_price_id"`
	// If the Price represents a fixed cost, this represents the quantity of units
	// applied.
	FixedPriceQuantity param.Field[float64] `json:"fixed_price_quantity"`
	// The property used to group this price on an invoice
	InvoiceGroupingKey param.Field[string] `json:"invoice_grouping_key"`
	// Within each billing cycle, specifies the cadence at which invoices are produced.
	// If unspecified, a single invoice is produced per billing cycle.
	InvoicingCycleConfiguration param.Field[PlanNewParamsPricesNewPlanBulkWithProrationPriceInvoicingCycleConfiguration] `json:"invoicing_cycle_configuration"`
	// User-specified key/value pairs for the resource. Individual keys can be removed
	// by setting the value to `null`, and the entire metadata mapping can be cleared
	// by setting `metadata` to `null`.
	Metadata param.Field[map[string]string] `json:"metadata"`
}

func (r PlanNewParamsPricesNewPlanBulkWithProrationPrice) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r PlanNewParamsPricesNewPlanBulkWithProrationPrice) implementsPlanNewParamsPriceUnion() {}

// The cadence to bill for this price on.
type PlanNewParamsPricesNewPlanBulkWithProrationPriceCadence string

const (
	PlanNewParamsPricesNewPlanBulkWithProrationPriceCadenceAnnual     PlanNewParamsPricesNewPlanBulkWithProrationPriceCadence = "annual"
	PlanNewParamsPricesNewPlanBulkWithProrationPriceCadenceSemiAnnual PlanNewParamsPricesNewPlanBulkWithProrationPriceCadence = "semi_annual"
	PlanNewParamsPricesNewPlanBulkWithProrationPriceCadenceMonthly    PlanNewParamsPricesNewPlanBulkWithProrationPriceCadence = "monthly"
	PlanNewParamsPricesNewPlanBulkWithProrationPriceCadenceQuarterly  PlanNewParamsPricesNewPlanBulkWithProrationPriceCadence = "quarterly"
	PlanNewParamsPricesNewPlanBulkWithProrationPriceCadenceOneTime    PlanNewParamsPricesNewPlanBulkWithProrationPriceCadence = "one_time"
	PlanNewParamsPricesNewPlanBulkWithProrationPriceCadenceCustom     PlanNewParamsPricesNewPlanBulkWithProrationPriceCadence = "custom"
)

func (r PlanNewParamsPricesNewPlanBulkWithProrationPriceCadence) IsKnown() bool {
	switch r {
	case PlanNewParamsPricesNewPlanBulkWithProrationPriceCadenceAnnual, PlanNewParamsPricesNewPlanBulkWithProrationPriceCadenceSemiAnnual, PlanNewParamsPricesNewPlanBulkWithProrationPriceCadenceMonthly, PlanNewParamsPricesNewPlanBulkWithProrationPriceCadenceQuarterly, PlanNewParamsPricesNewPlanBulkWithProrationPriceCadenceOneTime, PlanNewParamsPricesNewPlanBulkWithProrationPriceCadenceCustom:
		return true
	}
	return false
}

type PlanNewParamsPricesNewPlanBulkWithProrationPriceModelType string

const (
	PlanNewParamsPricesNewPlanBulkWithProrationPriceModelTypeBulkWithProration PlanNewParamsPricesNewPlanBulkWithProrationPriceModelType = "bulk_with_proration"
)

func (r PlanNewParamsPricesNewPlanBulkWithProrationPriceModelType) IsKnown() bool {
	switch r {
	case PlanNewParamsPricesNewPlanBulkWithProrationPriceModelTypeBulkWithProration:
		return true
	}
	return false
}

// For custom cadence: specifies the duration of the billing period in days or
// months.
type PlanNewParamsPricesNewPlanBulkWithProrationPriceBillingCycleConfiguration struct {
	// The duration of the billing period.
	Duration param.Field[int64] `json:"duration,required"`
	// The unit of billing period duration.
	DurationUnit param.Field[PlanNewParamsPricesNewPlanBulkWithProrationPriceBillingCycleConfigurationDurationUnit] `json:"duration_unit,required"`
}

func (r PlanNewParamsPricesNewPlanBulkWithProrationPriceBillingCycleConfiguration) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// The unit of billing period duration.
type PlanNewParamsPricesNewPlanBulkWithProrationPriceBillingCycleConfigurationDurationUnit string

const (
	PlanNewParamsPricesNewPlanBulkWithProrationPriceBillingCycleConfigurationDurationUnitDay   PlanNewParamsPricesNewPlanBulkWithProrationPriceBillingCycleConfigurationDurationUnit = "day"
	PlanNewParamsPricesNewPlanBulkWithProrationPriceBillingCycleConfigurationDurationUnitMonth PlanNewParamsPricesNewPlanBulkWithProrationPriceBillingCycleConfigurationDurationUnit = "month"
)

func (r PlanNewParamsPricesNewPlanBulkWithProrationPriceBillingCycleConfigurationDurationUnit) IsKnown() bool {
	switch r {
	case PlanNewParamsPricesNewPlanBulkWithProrationPriceBillingCycleConfigurationDurationUnitDay, PlanNewParamsPricesNewPlanBulkWithProrationPriceBillingCycleConfigurationDurationUnitMonth:
		return true
	}
	return false
}

// Within each billing cycle, specifies the cadence at which invoices are produced.
// If unspecified, a single invoice is produced per billing cycle.
type PlanNewParamsPricesNewPlanBulkWithProrationPriceInvoicingCycleConfiguration struct {
	// The duration of the billing period.
	Duration param.Field[int64] `json:"duration,required"`
	// The unit of billing period duration.
	DurationUnit param.Field[PlanNewParamsPricesNewPlanBulkWithProrationPriceInvoicingCycleConfigurationDurationUnit] `json:"duration_unit,required"`
}

func (r PlanNewParamsPricesNewPlanBulkWithProrationPriceInvoicingCycleConfiguration) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// The unit of billing period duration.
type PlanNewParamsPricesNewPlanBulkWithProrationPriceInvoicingCycleConfigurationDurationUnit string

const (
	PlanNewParamsPricesNewPlanBulkWithProrationPriceInvoicingCycleConfigurationDurationUnitDay   PlanNewParamsPricesNewPlanBulkWithProrationPriceInvoicingCycleConfigurationDurationUnit = "day"
	PlanNewParamsPricesNewPlanBulkWithProrationPriceInvoicingCycleConfigurationDurationUnitMonth PlanNewParamsPricesNewPlanBulkWithProrationPriceInvoicingCycleConfigurationDurationUnit = "month"
)

func (r PlanNewParamsPricesNewPlanBulkWithProrationPriceInvoicingCycleConfigurationDurationUnit) IsKnown() bool {
	switch r {
	case PlanNewParamsPricesNewPlanBulkWithProrationPriceInvoicingCycleConfigurationDurationUnitDay, PlanNewParamsPricesNewPlanBulkWithProrationPriceInvoicingCycleConfigurationDurationUnitMonth:
		return true
	}
	return false
}

type PlanNewParamsPricesNewPlanGroupedTieredPackagePrice struct {
	// The cadence to bill for this price on.
	Cadence                    param.Field[PlanNewParamsPricesNewPlanGroupedTieredPackagePriceCadence] `json:"cadence,required"`
	GroupedTieredPackageConfig param.Field[map[string]interface{}]                                     `json:"grouped_tiered_package_config,required"`
	// The id of the item the plan will be associated with.
	ItemID    param.Field[string]                                                       `json:"item_id,required"`
	ModelType param.Field[PlanNewParamsPricesNewPlanGroupedTieredPackagePriceModelType] `json:"model_type,required"`
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
	BillingCycleConfiguration param.Field[PlanNewParamsPricesNewPlanGroupedTieredPackagePriceBillingCycleConfiguration] `json:"billing_cycle_configuration"`
	// The per unit conversion rate of the price currency to the invoicing currency.
	ConversionRate param.Field[float64] `json:"conversion_rate"`
	// An ISO 4217 currency string, or custom pricing unit identifier, in which this
	// price is billed.
	Currency param.Field[string] `json:"currency"`
	// An alias for the price.
	ExternalPriceID param.Field[string] `json:"external_price_id"`
	// If the Price represents a fixed cost, this represents the quantity of units
	// applied.
	FixedPriceQuantity param.Field[float64] `json:"fixed_price_quantity"`
	// The property used to group this price on an invoice
	InvoiceGroupingKey param.Field[string] `json:"invoice_grouping_key"`
	// Within each billing cycle, specifies the cadence at which invoices are produced.
	// If unspecified, a single invoice is produced per billing cycle.
	InvoicingCycleConfiguration param.Field[PlanNewParamsPricesNewPlanGroupedTieredPackagePriceInvoicingCycleConfiguration] `json:"invoicing_cycle_configuration"`
	// User-specified key/value pairs for the resource. Individual keys can be removed
	// by setting the value to `null`, and the entire metadata mapping can be cleared
	// by setting `metadata` to `null`.
	Metadata param.Field[map[string]string] `json:"metadata"`
}

func (r PlanNewParamsPricesNewPlanGroupedTieredPackagePrice) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r PlanNewParamsPricesNewPlanGroupedTieredPackagePrice) implementsPlanNewParamsPriceUnion() {}

// The cadence to bill for this price on.
type PlanNewParamsPricesNewPlanGroupedTieredPackagePriceCadence string

const (
	PlanNewParamsPricesNewPlanGroupedTieredPackagePriceCadenceAnnual     PlanNewParamsPricesNewPlanGroupedTieredPackagePriceCadence = "annual"
	PlanNewParamsPricesNewPlanGroupedTieredPackagePriceCadenceSemiAnnual PlanNewParamsPricesNewPlanGroupedTieredPackagePriceCadence = "semi_annual"
	PlanNewParamsPricesNewPlanGroupedTieredPackagePriceCadenceMonthly    PlanNewParamsPricesNewPlanGroupedTieredPackagePriceCadence = "monthly"
	PlanNewParamsPricesNewPlanGroupedTieredPackagePriceCadenceQuarterly  PlanNewParamsPricesNewPlanGroupedTieredPackagePriceCadence = "quarterly"
	PlanNewParamsPricesNewPlanGroupedTieredPackagePriceCadenceOneTime    PlanNewParamsPricesNewPlanGroupedTieredPackagePriceCadence = "one_time"
	PlanNewParamsPricesNewPlanGroupedTieredPackagePriceCadenceCustom     PlanNewParamsPricesNewPlanGroupedTieredPackagePriceCadence = "custom"
)

func (r PlanNewParamsPricesNewPlanGroupedTieredPackagePriceCadence) IsKnown() bool {
	switch r {
	case PlanNewParamsPricesNewPlanGroupedTieredPackagePriceCadenceAnnual, PlanNewParamsPricesNewPlanGroupedTieredPackagePriceCadenceSemiAnnual, PlanNewParamsPricesNewPlanGroupedTieredPackagePriceCadenceMonthly, PlanNewParamsPricesNewPlanGroupedTieredPackagePriceCadenceQuarterly, PlanNewParamsPricesNewPlanGroupedTieredPackagePriceCadenceOneTime, PlanNewParamsPricesNewPlanGroupedTieredPackagePriceCadenceCustom:
		return true
	}
	return false
}

type PlanNewParamsPricesNewPlanGroupedTieredPackagePriceModelType string

const (
	PlanNewParamsPricesNewPlanGroupedTieredPackagePriceModelTypeGroupedTieredPackage PlanNewParamsPricesNewPlanGroupedTieredPackagePriceModelType = "grouped_tiered_package"
)

func (r PlanNewParamsPricesNewPlanGroupedTieredPackagePriceModelType) IsKnown() bool {
	switch r {
	case PlanNewParamsPricesNewPlanGroupedTieredPackagePriceModelTypeGroupedTieredPackage:
		return true
	}
	return false
}

// For custom cadence: specifies the duration of the billing period in days or
// months.
type PlanNewParamsPricesNewPlanGroupedTieredPackagePriceBillingCycleConfiguration struct {
	// The duration of the billing period.
	Duration param.Field[int64] `json:"duration,required"`
	// The unit of billing period duration.
	DurationUnit param.Field[PlanNewParamsPricesNewPlanGroupedTieredPackagePriceBillingCycleConfigurationDurationUnit] `json:"duration_unit,required"`
}

func (r PlanNewParamsPricesNewPlanGroupedTieredPackagePriceBillingCycleConfiguration) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// The unit of billing period duration.
type PlanNewParamsPricesNewPlanGroupedTieredPackagePriceBillingCycleConfigurationDurationUnit string

const (
	PlanNewParamsPricesNewPlanGroupedTieredPackagePriceBillingCycleConfigurationDurationUnitDay   PlanNewParamsPricesNewPlanGroupedTieredPackagePriceBillingCycleConfigurationDurationUnit = "day"
	PlanNewParamsPricesNewPlanGroupedTieredPackagePriceBillingCycleConfigurationDurationUnitMonth PlanNewParamsPricesNewPlanGroupedTieredPackagePriceBillingCycleConfigurationDurationUnit = "month"
)

func (r PlanNewParamsPricesNewPlanGroupedTieredPackagePriceBillingCycleConfigurationDurationUnit) IsKnown() bool {
	switch r {
	case PlanNewParamsPricesNewPlanGroupedTieredPackagePriceBillingCycleConfigurationDurationUnitDay, PlanNewParamsPricesNewPlanGroupedTieredPackagePriceBillingCycleConfigurationDurationUnitMonth:
		return true
	}
	return false
}

// Within each billing cycle, specifies the cadence at which invoices are produced.
// If unspecified, a single invoice is produced per billing cycle.
type PlanNewParamsPricesNewPlanGroupedTieredPackagePriceInvoicingCycleConfiguration struct {
	// The duration of the billing period.
	Duration param.Field[int64] `json:"duration,required"`
	// The unit of billing period duration.
	DurationUnit param.Field[PlanNewParamsPricesNewPlanGroupedTieredPackagePriceInvoicingCycleConfigurationDurationUnit] `json:"duration_unit,required"`
}

func (r PlanNewParamsPricesNewPlanGroupedTieredPackagePriceInvoicingCycleConfiguration) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// The unit of billing period duration.
type PlanNewParamsPricesNewPlanGroupedTieredPackagePriceInvoicingCycleConfigurationDurationUnit string

const (
	PlanNewParamsPricesNewPlanGroupedTieredPackagePriceInvoicingCycleConfigurationDurationUnitDay   PlanNewParamsPricesNewPlanGroupedTieredPackagePriceInvoicingCycleConfigurationDurationUnit = "day"
	PlanNewParamsPricesNewPlanGroupedTieredPackagePriceInvoicingCycleConfigurationDurationUnitMonth PlanNewParamsPricesNewPlanGroupedTieredPackagePriceInvoicingCycleConfigurationDurationUnit = "month"
)

func (r PlanNewParamsPricesNewPlanGroupedTieredPackagePriceInvoicingCycleConfigurationDurationUnit) IsKnown() bool {
	switch r {
	case PlanNewParamsPricesNewPlanGroupedTieredPackagePriceInvoicingCycleConfigurationDurationUnitDay, PlanNewParamsPricesNewPlanGroupedTieredPackagePriceInvoicingCycleConfigurationDurationUnitMonth:
		return true
	}
	return false
}

type PlanNewParamsPricesNewPlanMaxGroupTieredPackagePrice struct {
	// The cadence to bill for this price on.
	Cadence param.Field[PlanNewParamsPricesNewPlanMaxGroupTieredPackagePriceCadence] `json:"cadence,required"`
	// The id of the item the plan will be associated with.
	ItemID                      param.Field[string]                                                        `json:"item_id,required"`
	MaxGroupTieredPackageConfig param.Field[map[string]interface{}]                                        `json:"max_group_tiered_package_config,required"`
	ModelType                   param.Field[PlanNewParamsPricesNewPlanMaxGroupTieredPackagePriceModelType] `json:"model_type,required"`
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
	BillingCycleConfiguration param.Field[PlanNewParamsPricesNewPlanMaxGroupTieredPackagePriceBillingCycleConfiguration] `json:"billing_cycle_configuration"`
	// The per unit conversion rate of the price currency to the invoicing currency.
	ConversionRate param.Field[float64] `json:"conversion_rate"`
	// An ISO 4217 currency string, or custom pricing unit identifier, in which this
	// price is billed.
	Currency param.Field[string] `json:"currency"`
	// An alias for the price.
	ExternalPriceID param.Field[string] `json:"external_price_id"`
	// If the Price represents a fixed cost, this represents the quantity of units
	// applied.
	FixedPriceQuantity param.Field[float64] `json:"fixed_price_quantity"`
	// The property used to group this price on an invoice
	InvoiceGroupingKey param.Field[string] `json:"invoice_grouping_key"`
	// Within each billing cycle, specifies the cadence at which invoices are produced.
	// If unspecified, a single invoice is produced per billing cycle.
	InvoicingCycleConfiguration param.Field[PlanNewParamsPricesNewPlanMaxGroupTieredPackagePriceInvoicingCycleConfiguration] `json:"invoicing_cycle_configuration"`
	// User-specified key/value pairs for the resource. Individual keys can be removed
	// by setting the value to `null`, and the entire metadata mapping can be cleared
	// by setting `metadata` to `null`.
	Metadata param.Field[map[string]string] `json:"metadata"`
}

func (r PlanNewParamsPricesNewPlanMaxGroupTieredPackagePrice) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r PlanNewParamsPricesNewPlanMaxGroupTieredPackagePrice) implementsPlanNewParamsPriceUnion() {}

// The cadence to bill for this price on.
type PlanNewParamsPricesNewPlanMaxGroupTieredPackagePriceCadence string

const (
	PlanNewParamsPricesNewPlanMaxGroupTieredPackagePriceCadenceAnnual     PlanNewParamsPricesNewPlanMaxGroupTieredPackagePriceCadence = "annual"
	PlanNewParamsPricesNewPlanMaxGroupTieredPackagePriceCadenceSemiAnnual PlanNewParamsPricesNewPlanMaxGroupTieredPackagePriceCadence = "semi_annual"
	PlanNewParamsPricesNewPlanMaxGroupTieredPackagePriceCadenceMonthly    PlanNewParamsPricesNewPlanMaxGroupTieredPackagePriceCadence = "monthly"
	PlanNewParamsPricesNewPlanMaxGroupTieredPackagePriceCadenceQuarterly  PlanNewParamsPricesNewPlanMaxGroupTieredPackagePriceCadence = "quarterly"
	PlanNewParamsPricesNewPlanMaxGroupTieredPackagePriceCadenceOneTime    PlanNewParamsPricesNewPlanMaxGroupTieredPackagePriceCadence = "one_time"
	PlanNewParamsPricesNewPlanMaxGroupTieredPackagePriceCadenceCustom     PlanNewParamsPricesNewPlanMaxGroupTieredPackagePriceCadence = "custom"
)

func (r PlanNewParamsPricesNewPlanMaxGroupTieredPackagePriceCadence) IsKnown() bool {
	switch r {
	case PlanNewParamsPricesNewPlanMaxGroupTieredPackagePriceCadenceAnnual, PlanNewParamsPricesNewPlanMaxGroupTieredPackagePriceCadenceSemiAnnual, PlanNewParamsPricesNewPlanMaxGroupTieredPackagePriceCadenceMonthly, PlanNewParamsPricesNewPlanMaxGroupTieredPackagePriceCadenceQuarterly, PlanNewParamsPricesNewPlanMaxGroupTieredPackagePriceCadenceOneTime, PlanNewParamsPricesNewPlanMaxGroupTieredPackagePriceCadenceCustom:
		return true
	}
	return false
}

type PlanNewParamsPricesNewPlanMaxGroupTieredPackagePriceModelType string

const (
	PlanNewParamsPricesNewPlanMaxGroupTieredPackagePriceModelTypeMaxGroupTieredPackage PlanNewParamsPricesNewPlanMaxGroupTieredPackagePriceModelType = "max_group_tiered_package"
)

func (r PlanNewParamsPricesNewPlanMaxGroupTieredPackagePriceModelType) IsKnown() bool {
	switch r {
	case PlanNewParamsPricesNewPlanMaxGroupTieredPackagePriceModelTypeMaxGroupTieredPackage:
		return true
	}
	return false
}

// For custom cadence: specifies the duration of the billing period in days or
// months.
type PlanNewParamsPricesNewPlanMaxGroupTieredPackagePriceBillingCycleConfiguration struct {
	// The duration of the billing period.
	Duration param.Field[int64] `json:"duration,required"`
	// The unit of billing period duration.
	DurationUnit param.Field[PlanNewParamsPricesNewPlanMaxGroupTieredPackagePriceBillingCycleConfigurationDurationUnit] `json:"duration_unit,required"`
}

func (r PlanNewParamsPricesNewPlanMaxGroupTieredPackagePriceBillingCycleConfiguration) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// The unit of billing period duration.
type PlanNewParamsPricesNewPlanMaxGroupTieredPackagePriceBillingCycleConfigurationDurationUnit string

const (
	PlanNewParamsPricesNewPlanMaxGroupTieredPackagePriceBillingCycleConfigurationDurationUnitDay   PlanNewParamsPricesNewPlanMaxGroupTieredPackagePriceBillingCycleConfigurationDurationUnit = "day"
	PlanNewParamsPricesNewPlanMaxGroupTieredPackagePriceBillingCycleConfigurationDurationUnitMonth PlanNewParamsPricesNewPlanMaxGroupTieredPackagePriceBillingCycleConfigurationDurationUnit = "month"
)

func (r PlanNewParamsPricesNewPlanMaxGroupTieredPackagePriceBillingCycleConfigurationDurationUnit) IsKnown() bool {
	switch r {
	case PlanNewParamsPricesNewPlanMaxGroupTieredPackagePriceBillingCycleConfigurationDurationUnitDay, PlanNewParamsPricesNewPlanMaxGroupTieredPackagePriceBillingCycleConfigurationDurationUnitMonth:
		return true
	}
	return false
}

// Within each billing cycle, specifies the cadence at which invoices are produced.
// If unspecified, a single invoice is produced per billing cycle.
type PlanNewParamsPricesNewPlanMaxGroupTieredPackagePriceInvoicingCycleConfiguration struct {
	// The duration of the billing period.
	Duration param.Field[int64] `json:"duration,required"`
	// The unit of billing period duration.
	DurationUnit param.Field[PlanNewParamsPricesNewPlanMaxGroupTieredPackagePriceInvoicingCycleConfigurationDurationUnit] `json:"duration_unit,required"`
}

func (r PlanNewParamsPricesNewPlanMaxGroupTieredPackagePriceInvoicingCycleConfiguration) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// The unit of billing period duration.
type PlanNewParamsPricesNewPlanMaxGroupTieredPackagePriceInvoicingCycleConfigurationDurationUnit string

const (
	PlanNewParamsPricesNewPlanMaxGroupTieredPackagePriceInvoicingCycleConfigurationDurationUnitDay   PlanNewParamsPricesNewPlanMaxGroupTieredPackagePriceInvoicingCycleConfigurationDurationUnit = "day"
	PlanNewParamsPricesNewPlanMaxGroupTieredPackagePriceInvoicingCycleConfigurationDurationUnitMonth PlanNewParamsPricesNewPlanMaxGroupTieredPackagePriceInvoicingCycleConfigurationDurationUnit = "month"
)

func (r PlanNewParamsPricesNewPlanMaxGroupTieredPackagePriceInvoicingCycleConfigurationDurationUnit) IsKnown() bool {
	switch r {
	case PlanNewParamsPricesNewPlanMaxGroupTieredPackagePriceInvoicingCycleConfigurationDurationUnitDay, PlanNewParamsPricesNewPlanMaxGroupTieredPackagePriceInvoicingCycleConfigurationDurationUnitMonth:
		return true
	}
	return false
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
	PlanNewParamsPricesModelTypeUnit                       PlanNewParamsPricesModelType = "unit"
	PlanNewParamsPricesModelTypePackage                    PlanNewParamsPricesModelType = "package"
	PlanNewParamsPricesModelTypeMatrix                     PlanNewParamsPricesModelType = "matrix"
	PlanNewParamsPricesModelTypeTiered                     PlanNewParamsPricesModelType = "tiered"
	PlanNewParamsPricesModelTypeTieredBps                  PlanNewParamsPricesModelType = "tiered_bps"
	PlanNewParamsPricesModelTypeBps                        PlanNewParamsPricesModelType = "bps"
	PlanNewParamsPricesModelTypeBulkBps                    PlanNewParamsPricesModelType = "bulk_bps"
	PlanNewParamsPricesModelTypeBulk                       PlanNewParamsPricesModelType = "bulk"
	PlanNewParamsPricesModelTypeThresholdTotalAmount       PlanNewParamsPricesModelType = "threshold_total_amount"
	PlanNewParamsPricesModelTypeTieredPackage              PlanNewParamsPricesModelType = "tiered_package"
	PlanNewParamsPricesModelTypeTieredWithMinimum          PlanNewParamsPricesModelType = "tiered_with_minimum"
	PlanNewParamsPricesModelTypeUnitWithPercent            PlanNewParamsPricesModelType = "unit_with_percent"
	PlanNewParamsPricesModelTypePackageWithAllocation      PlanNewParamsPricesModelType = "package_with_allocation"
	PlanNewParamsPricesModelTypeTieredWithProration        PlanNewParamsPricesModelType = "tiered_with_proration"
	PlanNewParamsPricesModelTypeUnitWithProration          PlanNewParamsPricesModelType = "unit_with_proration"
	PlanNewParamsPricesModelTypeGroupedAllocation          PlanNewParamsPricesModelType = "grouped_allocation"
	PlanNewParamsPricesModelTypeGroupedWithProratedMinimum PlanNewParamsPricesModelType = "grouped_with_prorated_minimum"
	PlanNewParamsPricesModelTypeGroupedWithMeteredMinimum  PlanNewParamsPricesModelType = "grouped_with_metered_minimum"
	PlanNewParamsPricesModelTypeMatrixWithDisplayName      PlanNewParamsPricesModelType = "matrix_with_display_name"
	PlanNewParamsPricesModelTypeBulkWithProration          PlanNewParamsPricesModelType = "bulk_with_proration"
	PlanNewParamsPricesModelTypeGroupedTieredPackage       PlanNewParamsPricesModelType = "grouped_tiered_package"
	PlanNewParamsPricesModelTypeMaxGroupTieredPackage      PlanNewParamsPricesModelType = "max_group_tiered_package"
)

func (r PlanNewParamsPricesModelType) IsKnown() bool {
	switch r {
	case PlanNewParamsPricesModelTypeUnit, PlanNewParamsPricesModelTypePackage, PlanNewParamsPricesModelTypeMatrix, PlanNewParamsPricesModelTypeTiered, PlanNewParamsPricesModelTypeTieredBps, PlanNewParamsPricesModelTypeBps, PlanNewParamsPricesModelTypeBulkBps, PlanNewParamsPricesModelTypeBulk, PlanNewParamsPricesModelTypeThresholdTotalAmount, PlanNewParamsPricesModelTypeTieredPackage, PlanNewParamsPricesModelTypeTieredWithMinimum, PlanNewParamsPricesModelTypeUnitWithPercent, PlanNewParamsPricesModelTypePackageWithAllocation, PlanNewParamsPricesModelTypeTieredWithProration, PlanNewParamsPricesModelTypeUnitWithProration, PlanNewParamsPricesModelTypeGroupedAllocation, PlanNewParamsPricesModelTypeGroupedWithProratedMinimum, PlanNewParamsPricesModelTypeGroupedWithMeteredMinimum, PlanNewParamsPricesModelTypeMatrixWithDisplayName, PlanNewParamsPricesModelTypeBulkWithProration, PlanNewParamsPricesModelTypeGroupedTieredPackage, PlanNewParamsPricesModelTypeMaxGroupTieredPackage:
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
