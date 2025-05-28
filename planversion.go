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

// PlanVersionService contains methods and other services that help with
// interacting with the orb API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewPlanVersionService] method instead.
type PlanVersionService struct {
	Options []option.RequestOption
}

// NewPlanVersionService generates a new service that applies the given options to
// each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewPlanVersionService(opts ...option.RequestOption) (r *PlanVersionService) {
	r = &PlanVersionService{}
	r.Options = opts
	return
}

// This API endpoint is in beta and its interface may change. It is recommended for
// use only in test mode.
//
// This endpoint allows the creation of a new plan version for an existing plan.
func (r *PlanVersionService) New(ctx context.Context, planID string, body PlanVersionNewParams, opts ...option.RequestOption) (res *PlanVersionNewResponse, err error) {
	opts = append(r.Options[:], opts...)
	if planID == "" {
		err = errors.New("missing required plan_id parameter")
		return
	}
	path := fmt.Sprintf("plans/%s/versions", planID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return
}

// This API endpoint is in beta and its interface may change. It is recommended for
// use only in test mode.
//
// This endpoint is used to fetch a plan version. It returns the phases, prices,
// and adjustments present on this version of the plan.
func (r *PlanVersionService) Get(ctx context.Context, planID string, version string, opts ...option.RequestOption) (res *PlanVersionGetResponse, err error) {
	opts = append(r.Options[:], opts...)
	if planID == "" {
		err = errors.New("missing required plan_id parameter")
		return
	}
	if version == "" {
		err = errors.New("missing required version parameter")
		return
	}
	path := fmt.Sprintf("plans/%s/versions/%s", planID, version)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

// The PlanVersion resource represents the prices and adjustments present on a
// specific version of a plan.
type PlanVersionNewResponse struct {
	// Adjustments for this plan. If the plan has phases, this includes adjustments
	// across all phases of the plan.
	Adjustments []PlanVersionNewResponseAdjustment `json:"adjustments,required"`
	CreatedAt   time.Time                          `json:"created_at,required" format:"date-time"`
	PlanPhases  []PlanVersionNewResponsePlanPhase  `json:"plan_phases,required,nullable"`
	// Prices for this plan. If the plan has phases, this includes prices across all
	// phases of the plan.
	Prices  []Price                    `json:"prices,required"`
	Version int64                      `json:"version,required"`
	JSON    planVersionNewResponseJSON `json:"-"`
}

// planVersionNewResponseJSON contains the JSON metadata for the struct
// [PlanVersionNewResponse]
type planVersionNewResponseJSON struct {
	Adjustments apijson.Field
	CreatedAt   apijson.Field
	PlanPhases  apijson.Field
	Prices      apijson.Field
	Version     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *PlanVersionNewResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r planVersionNewResponseJSON) RawJSON() string {
	return r.raw
}

type PlanVersionNewResponseAdjustment struct {
	ID             string                                          `json:"id,required"`
	AdjustmentType PlanVersionNewResponseAdjustmentsAdjustmentType `json:"adjustment_type,required"`
	// This field can have the runtime type of [[]string].
	AppliesToPriceIDs interface{} `json:"applies_to_price_ids,required"`
	// This field can have the runtime type of
	// [[]PlanVersionNewResponseAdjustmentsPlanPhaseUsageDiscountAdjustmentFilter],
	// [[]PlanVersionNewResponseAdjustmentsPlanPhaseAmountDiscountAdjustmentFilter],
	// [[]PlanVersionNewResponseAdjustmentsPlanPhasePercentageDiscountAdjustmentFilter],
	// [[]PlanVersionNewResponseAdjustmentsPlanPhaseMinimumAdjustmentFilter],
	// [[]PlanVersionNewResponseAdjustmentsPlanPhaseMaximumAdjustmentFilter].
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
	UsageDiscount float64                              `json:"usage_discount"`
	JSON          planVersionNewResponseAdjustmentJSON `json:"-"`
	union         PlanVersionNewResponseAdjustmentsUnion
}

// planVersionNewResponseAdjustmentJSON contains the JSON metadata for the struct
// [PlanVersionNewResponseAdjustment]
type planVersionNewResponseAdjustmentJSON struct {
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

func (r planVersionNewResponseAdjustmentJSON) RawJSON() string {
	return r.raw
}

func (r *PlanVersionNewResponseAdjustment) UnmarshalJSON(data []byte) (err error) {
	*r = PlanVersionNewResponseAdjustment{}
	err = apijson.UnmarshalRoot(data, &r.union)
	if err != nil {
		return err
	}
	return apijson.Port(r.union, &r)
}

// AsUnion returns a [PlanVersionNewResponseAdjustmentsUnion] interface which you
// can cast to the specific types for more type safety.
//
// Possible runtime types of the union are
// [PlanVersionNewResponseAdjustmentsPlanPhaseUsageDiscountAdjustment],
// [PlanVersionNewResponseAdjustmentsPlanPhaseAmountDiscountAdjustment],
// [PlanVersionNewResponseAdjustmentsPlanPhasePercentageDiscountAdjustment],
// [PlanVersionNewResponseAdjustmentsPlanPhaseMinimumAdjustment],
// [PlanVersionNewResponseAdjustmentsPlanPhaseMaximumAdjustment].
func (r PlanVersionNewResponseAdjustment) AsUnion() PlanVersionNewResponseAdjustmentsUnion {
	return r.union
}

// Union satisfied by
// [PlanVersionNewResponseAdjustmentsPlanPhaseUsageDiscountAdjustment],
// [PlanVersionNewResponseAdjustmentsPlanPhaseAmountDiscountAdjustment],
// [PlanVersionNewResponseAdjustmentsPlanPhasePercentageDiscountAdjustment],
// [PlanVersionNewResponseAdjustmentsPlanPhaseMinimumAdjustment] or
// [PlanVersionNewResponseAdjustmentsPlanPhaseMaximumAdjustment].
type PlanVersionNewResponseAdjustmentsUnion interface {
	implementsPlanVersionNewResponseAdjustment()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*PlanVersionNewResponseAdjustmentsUnion)(nil)).Elem(),
		"adjustment_type",
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(PlanVersionNewResponseAdjustmentsPlanPhaseUsageDiscountAdjustment{}),
			DiscriminatorValue: "usage_discount",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(PlanVersionNewResponseAdjustmentsPlanPhaseAmountDiscountAdjustment{}),
			DiscriminatorValue: "amount_discount",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(PlanVersionNewResponseAdjustmentsPlanPhasePercentageDiscountAdjustment{}),
			DiscriminatorValue: "percentage_discount",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(PlanVersionNewResponseAdjustmentsPlanPhaseMinimumAdjustment{}),
			DiscriminatorValue: "minimum",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(PlanVersionNewResponseAdjustmentsPlanPhaseMaximumAdjustment{}),
			DiscriminatorValue: "maximum",
		},
	)
}

type PlanVersionNewResponseAdjustmentsPlanPhaseUsageDiscountAdjustment struct {
	ID             string                                                                          `json:"id,required"`
	AdjustmentType PlanVersionNewResponseAdjustmentsPlanPhaseUsageDiscountAdjustmentAdjustmentType `json:"adjustment_type,required"`
	// The price IDs that this adjustment applies to.
	//
	// Deprecated: deprecated
	AppliesToPriceIDs []string `json:"applies_to_price_ids,required"`
	// The filters that determine which prices to apply this adjustment to.
	Filters []PlanVersionNewResponseAdjustmentsPlanPhaseUsageDiscountAdjustmentFilter `json:"filters,required"`
	// True for adjustments that apply to an entire invocice, false for adjustments
	// that apply to only one price.
	IsInvoiceLevel bool `json:"is_invoice_level,required"`
	// The plan phase in which this adjustment is active.
	PlanPhaseOrder int64 `json:"plan_phase_order,required,nullable"`
	// The reason for the adjustment.
	Reason string `json:"reason,required,nullable"`
	// The number of usage units by which to discount the price this adjustment applies
	// to in a given billing period.
	UsageDiscount float64                                                               `json:"usage_discount,required"`
	JSON          planVersionNewResponseAdjustmentsPlanPhaseUsageDiscountAdjustmentJSON `json:"-"`
}

// planVersionNewResponseAdjustmentsPlanPhaseUsageDiscountAdjustmentJSON contains
// the JSON metadata for the struct
// [PlanVersionNewResponseAdjustmentsPlanPhaseUsageDiscountAdjustment]
type planVersionNewResponseAdjustmentsPlanPhaseUsageDiscountAdjustmentJSON struct {
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

func (r *PlanVersionNewResponseAdjustmentsPlanPhaseUsageDiscountAdjustment) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r planVersionNewResponseAdjustmentsPlanPhaseUsageDiscountAdjustmentJSON) RawJSON() string {
	return r.raw
}

func (r PlanVersionNewResponseAdjustmentsPlanPhaseUsageDiscountAdjustment) implementsPlanVersionNewResponseAdjustment() {
}

type PlanVersionNewResponseAdjustmentsPlanPhaseUsageDiscountAdjustmentAdjustmentType string

const (
	PlanVersionNewResponseAdjustmentsPlanPhaseUsageDiscountAdjustmentAdjustmentTypeUsageDiscount PlanVersionNewResponseAdjustmentsPlanPhaseUsageDiscountAdjustmentAdjustmentType = "usage_discount"
)

func (r PlanVersionNewResponseAdjustmentsPlanPhaseUsageDiscountAdjustmentAdjustmentType) IsKnown() bool {
	switch r {
	case PlanVersionNewResponseAdjustmentsPlanPhaseUsageDiscountAdjustmentAdjustmentTypeUsageDiscount:
		return true
	}
	return false
}

type PlanVersionNewResponseAdjustmentsPlanPhaseUsageDiscountAdjustmentFilter struct {
	// The property of the price to filter on.
	Field PlanVersionNewResponseAdjustmentsPlanPhaseUsageDiscountAdjustmentFiltersField `json:"field,required"`
	// Should prices that match the filter be included or excluded.
	Operator PlanVersionNewResponseAdjustmentsPlanPhaseUsageDiscountAdjustmentFiltersOperator `json:"operator,required"`
	// The IDs or values that match this filter.
	Values []string                                                                    `json:"values,required"`
	JSON   planVersionNewResponseAdjustmentsPlanPhaseUsageDiscountAdjustmentFilterJSON `json:"-"`
}

// planVersionNewResponseAdjustmentsPlanPhaseUsageDiscountAdjustmentFilterJSON
// contains the JSON metadata for the struct
// [PlanVersionNewResponseAdjustmentsPlanPhaseUsageDiscountAdjustmentFilter]
type planVersionNewResponseAdjustmentsPlanPhaseUsageDiscountAdjustmentFilterJSON struct {
	Field       apijson.Field
	Operator    apijson.Field
	Values      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *PlanVersionNewResponseAdjustmentsPlanPhaseUsageDiscountAdjustmentFilter) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r planVersionNewResponseAdjustmentsPlanPhaseUsageDiscountAdjustmentFilterJSON) RawJSON() string {
	return r.raw
}

// The property of the price to filter on.
type PlanVersionNewResponseAdjustmentsPlanPhaseUsageDiscountAdjustmentFiltersField string

const (
	PlanVersionNewResponseAdjustmentsPlanPhaseUsageDiscountAdjustmentFiltersFieldPriceID       PlanVersionNewResponseAdjustmentsPlanPhaseUsageDiscountAdjustmentFiltersField = "price_id"
	PlanVersionNewResponseAdjustmentsPlanPhaseUsageDiscountAdjustmentFiltersFieldItemID        PlanVersionNewResponseAdjustmentsPlanPhaseUsageDiscountAdjustmentFiltersField = "item_id"
	PlanVersionNewResponseAdjustmentsPlanPhaseUsageDiscountAdjustmentFiltersFieldPriceType     PlanVersionNewResponseAdjustmentsPlanPhaseUsageDiscountAdjustmentFiltersField = "price_type"
	PlanVersionNewResponseAdjustmentsPlanPhaseUsageDiscountAdjustmentFiltersFieldCurrency      PlanVersionNewResponseAdjustmentsPlanPhaseUsageDiscountAdjustmentFiltersField = "currency"
	PlanVersionNewResponseAdjustmentsPlanPhaseUsageDiscountAdjustmentFiltersFieldPricingUnitID PlanVersionNewResponseAdjustmentsPlanPhaseUsageDiscountAdjustmentFiltersField = "pricing_unit_id"
)

func (r PlanVersionNewResponseAdjustmentsPlanPhaseUsageDiscountAdjustmentFiltersField) IsKnown() bool {
	switch r {
	case PlanVersionNewResponseAdjustmentsPlanPhaseUsageDiscountAdjustmentFiltersFieldPriceID, PlanVersionNewResponseAdjustmentsPlanPhaseUsageDiscountAdjustmentFiltersFieldItemID, PlanVersionNewResponseAdjustmentsPlanPhaseUsageDiscountAdjustmentFiltersFieldPriceType, PlanVersionNewResponseAdjustmentsPlanPhaseUsageDiscountAdjustmentFiltersFieldCurrency, PlanVersionNewResponseAdjustmentsPlanPhaseUsageDiscountAdjustmentFiltersFieldPricingUnitID:
		return true
	}
	return false
}

// Should prices that match the filter be included or excluded.
type PlanVersionNewResponseAdjustmentsPlanPhaseUsageDiscountAdjustmentFiltersOperator string

const (
	PlanVersionNewResponseAdjustmentsPlanPhaseUsageDiscountAdjustmentFiltersOperatorIncludes PlanVersionNewResponseAdjustmentsPlanPhaseUsageDiscountAdjustmentFiltersOperator = "includes"
	PlanVersionNewResponseAdjustmentsPlanPhaseUsageDiscountAdjustmentFiltersOperatorExcludes PlanVersionNewResponseAdjustmentsPlanPhaseUsageDiscountAdjustmentFiltersOperator = "excludes"
)

func (r PlanVersionNewResponseAdjustmentsPlanPhaseUsageDiscountAdjustmentFiltersOperator) IsKnown() bool {
	switch r {
	case PlanVersionNewResponseAdjustmentsPlanPhaseUsageDiscountAdjustmentFiltersOperatorIncludes, PlanVersionNewResponseAdjustmentsPlanPhaseUsageDiscountAdjustmentFiltersOperatorExcludes:
		return true
	}
	return false
}

type PlanVersionNewResponseAdjustmentsPlanPhaseAmountDiscountAdjustment struct {
	ID             string                                                                           `json:"id,required"`
	AdjustmentType PlanVersionNewResponseAdjustmentsPlanPhaseAmountDiscountAdjustmentAdjustmentType `json:"adjustment_type,required"`
	// The amount by which to discount the prices this adjustment applies to in a given
	// billing period.
	AmountDiscount string `json:"amount_discount,required"`
	// The price IDs that this adjustment applies to.
	//
	// Deprecated: deprecated
	AppliesToPriceIDs []string `json:"applies_to_price_ids,required"`
	// The filters that determine which prices to apply this adjustment to.
	Filters []PlanVersionNewResponseAdjustmentsPlanPhaseAmountDiscountAdjustmentFilter `json:"filters,required"`
	// True for adjustments that apply to an entire invocice, false for adjustments
	// that apply to only one price.
	IsInvoiceLevel bool `json:"is_invoice_level,required"`
	// The plan phase in which this adjustment is active.
	PlanPhaseOrder int64 `json:"plan_phase_order,required,nullable"`
	// The reason for the adjustment.
	Reason string                                                                 `json:"reason,required,nullable"`
	JSON   planVersionNewResponseAdjustmentsPlanPhaseAmountDiscountAdjustmentJSON `json:"-"`
}

// planVersionNewResponseAdjustmentsPlanPhaseAmountDiscountAdjustmentJSON contains
// the JSON metadata for the struct
// [PlanVersionNewResponseAdjustmentsPlanPhaseAmountDiscountAdjustment]
type planVersionNewResponseAdjustmentsPlanPhaseAmountDiscountAdjustmentJSON struct {
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

func (r *PlanVersionNewResponseAdjustmentsPlanPhaseAmountDiscountAdjustment) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r planVersionNewResponseAdjustmentsPlanPhaseAmountDiscountAdjustmentJSON) RawJSON() string {
	return r.raw
}

func (r PlanVersionNewResponseAdjustmentsPlanPhaseAmountDiscountAdjustment) implementsPlanVersionNewResponseAdjustment() {
}

type PlanVersionNewResponseAdjustmentsPlanPhaseAmountDiscountAdjustmentAdjustmentType string

const (
	PlanVersionNewResponseAdjustmentsPlanPhaseAmountDiscountAdjustmentAdjustmentTypeAmountDiscount PlanVersionNewResponseAdjustmentsPlanPhaseAmountDiscountAdjustmentAdjustmentType = "amount_discount"
)

func (r PlanVersionNewResponseAdjustmentsPlanPhaseAmountDiscountAdjustmentAdjustmentType) IsKnown() bool {
	switch r {
	case PlanVersionNewResponseAdjustmentsPlanPhaseAmountDiscountAdjustmentAdjustmentTypeAmountDiscount:
		return true
	}
	return false
}

type PlanVersionNewResponseAdjustmentsPlanPhaseAmountDiscountAdjustmentFilter struct {
	// The property of the price to filter on.
	Field PlanVersionNewResponseAdjustmentsPlanPhaseAmountDiscountAdjustmentFiltersField `json:"field,required"`
	// Should prices that match the filter be included or excluded.
	Operator PlanVersionNewResponseAdjustmentsPlanPhaseAmountDiscountAdjustmentFiltersOperator `json:"operator,required"`
	// The IDs or values that match this filter.
	Values []string                                                                     `json:"values,required"`
	JSON   planVersionNewResponseAdjustmentsPlanPhaseAmountDiscountAdjustmentFilterJSON `json:"-"`
}

// planVersionNewResponseAdjustmentsPlanPhaseAmountDiscountAdjustmentFilterJSON
// contains the JSON metadata for the struct
// [PlanVersionNewResponseAdjustmentsPlanPhaseAmountDiscountAdjustmentFilter]
type planVersionNewResponseAdjustmentsPlanPhaseAmountDiscountAdjustmentFilterJSON struct {
	Field       apijson.Field
	Operator    apijson.Field
	Values      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *PlanVersionNewResponseAdjustmentsPlanPhaseAmountDiscountAdjustmentFilter) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r planVersionNewResponseAdjustmentsPlanPhaseAmountDiscountAdjustmentFilterJSON) RawJSON() string {
	return r.raw
}

// The property of the price to filter on.
type PlanVersionNewResponseAdjustmentsPlanPhaseAmountDiscountAdjustmentFiltersField string

const (
	PlanVersionNewResponseAdjustmentsPlanPhaseAmountDiscountAdjustmentFiltersFieldPriceID       PlanVersionNewResponseAdjustmentsPlanPhaseAmountDiscountAdjustmentFiltersField = "price_id"
	PlanVersionNewResponseAdjustmentsPlanPhaseAmountDiscountAdjustmentFiltersFieldItemID        PlanVersionNewResponseAdjustmentsPlanPhaseAmountDiscountAdjustmentFiltersField = "item_id"
	PlanVersionNewResponseAdjustmentsPlanPhaseAmountDiscountAdjustmentFiltersFieldPriceType     PlanVersionNewResponseAdjustmentsPlanPhaseAmountDiscountAdjustmentFiltersField = "price_type"
	PlanVersionNewResponseAdjustmentsPlanPhaseAmountDiscountAdjustmentFiltersFieldCurrency      PlanVersionNewResponseAdjustmentsPlanPhaseAmountDiscountAdjustmentFiltersField = "currency"
	PlanVersionNewResponseAdjustmentsPlanPhaseAmountDiscountAdjustmentFiltersFieldPricingUnitID PlanVersionNewResponseAdjustmentsPlanPhaseAmountDiscountAdjustmentFiltersField = "pricing_unit_id"
)

func (r PlanVersionNewResponseAdjustmentsPlanPhaseAmountDiscountAdjustmentFiltersField) IsKnown() bool {
	switch r {
	case PlanVersionNewResponseAdjustmentsPlanPhaseAmountDiscountAdjustmentFiltersFieldPriceID, PlanVersionNewResponseAdjustmentsPlanPhaseAmountDiscountAdjustmentFiltersFieldItemID, PlanVersionNewResponseAdjustmentsPlanPhaseAmountDiscountAdjustmentFiltersFieldPriceType, PlanVersionNewResponseAdjustmentsPlanPhaseAmountDiscountAdjustmentFiltersFieldCurrency, PlanVersionNewResponseAdjustmentsPlanPhaseAmountDiscountAdjustmentFiltersFieldPricingUnitID:
		return true
	}
	return false
}

// Should prices that match the filter be included or excluded.
type PlanVersionNewResponseAdjustmentsPlanPhaseAmountDiscountAdjustmentFiltersOperator string

const (
	PlanVersionNewResponseAdjustmentsPlanPhaseAmountDiscountAdjustmentFiltersOperatorIncludes PlanVersionNewResponseAdjustmentsPlanPhaseAmountDiscountAdjustmentFiltersOperator = "includes"
	PlanVersionNewResponseAdjustmentsPlanPhaseAmountDiscountAdjustmentFiltersOperatorExcludes PlanVersionNewResponseAdjustmentsPlanPhaseAmountDiscountAdjustmentFiltersOperator = "excludes"
)

func (r PlanVersionNewResponseAdjustmentsPlanPhaseAmountDiscountAdjustmentFiltersOperator) IsKnown() bool {
	switch r {
	case PlanVersionNewResponseAdjustmentsPlanPhaseAmountDiscountAdjustmentFiltersOperatorIncludes, PlanVersionNewResponseAdjustmentsPlanPhaseAmountDiscountAdjustmentFiltersOperatorExcludes:
		return true
	}
	return false
}

type PlanVersionNewResponseAdjustmentsPlanPhasePercentageDiscountAdjustment struct {
	ID             string                                                                               `json:"id,required"`
	AdjustmentType PlanVersionNewResponseAdjustmentsPlanPhasePercentageDiscountAdjustmentAdjustmentType `json:"adjustment_type,required"`
	// The price IDs that this adjustment applies to.
	//
	// Deprecated: deprecated
	AppliesToPriceIDs []string `json:"applies_to_price_ids,required"`
	// The filters that determine which prices to apply this adjustment to.
	Filters []PlanVersionNewResponseAdjustmentsPlanPhasePercentageDiscountAdjustmentFilter `json:"filters,required"`
	// True for adjustments that apply to an entire invocice, false for adjustments
	// that apply to only one price.
	IsInvoiceLevel bool `json:"is_invoice_level,required"`
	// The percentage (as a value between 0 and 1) by which to discount the price
	// intervals this adjustment applies to in a given billing period.
	PercentageDiscount float64 `json:"percentage_discount,required"`
	// The plan phase in which this adjustment is active.
	PlanPhaseOrder int64 `json:"plan_phase_order,required,nullable"`
	// The reason for the adjustment.
	Reason string                                                                     `json:"reason,required,nullable"`
	JSON   planVersionNewResponseAdjustmentsPlanPhasePercentageDiscountAdjustmentJSON `json:"-"`
}

// planVersionNewResponseAdjustmentsPlanPhasePercentageDiscountAdjustmentJSON
// contains the JSON metadata for the struct
// [PlanVersionNewResponseAdjustmentsPlanPhasePercentageDiscountAdjustment]
type planVersionNewResponseAdjustmentsPlanPhasePercentageDiscountAdjustmentJSON struct {
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

func (r *PlanVersionNewResponseAdjustmentsPlanPhasePercentageDiscountAdjustment) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r planVersionNewResponseAdjustmentsPlanPhasePercentageDiscountAdjustmentJSON) RawJSON() string {
	return r.raw
}

func (r PlanVersionNewResponseAdjustmentsPlanPhasePercentageDiscountAdjustment) implementsPlanVersionNewResponseAdjustment() {
}

type PlanVersionNewResponseAdjustmentsPlanPhasePercentageDiscountAdjustmentAdjustmentType string

const (
	PlanVersionNewResponseAdjustmentsPlanPhasePercentageDiscountAdjustmentAdjustmentTypePercentageDiscount PlanVersionNewResponseAdjustmentsPlanPhasePercentageDiscountAdjustmentAdjustmentType = "percentage_discount"
)

func (r PlanVersionNewResponseAdjustmentsPlanPhasePercentageDiscountAdjustmentAdjustmentType) IsKnown() bool {
	switch r {
	case PlanVersionNewResponseAdjustmentsPlanPhasePercentageDiscountAdjustmentAdjustmentTypePercentageDiscount:
		return true
	}
	return false
}

type PlanVersionNewResponseAdjustmentsPlanPhasePercentageDiscountAdjustmentFilter struct {
	// The property of the price to filter on.
	Field PlanVersionNewResponseAdjustmentsPlanPhasePercentageDiscountAdjustmentFiltersField `json:"field,required"`
	// Should prices that match the filter be included or excluded.
	Operator PlanVersionNewResponseAdjustmentsPlanPhasePercentageDiscountAdjustmentFiltersOperator `json:"operator,required"`
	// The IDs or values that match this filter.
	Values []string                                                                         `json:"values,required"`
	JSON   planVersionNewResponseAdjustmentsPlanPhasePercentageDiscountAdjustmentFilterJSON `json:"-"`
}

// planVersionNewResponseAdjustmentsPlanPhasePercentageDiscountAdjustmentFilterJSON
// contains the JSON metadata for the struct
// [PlanVersionNewResponseAdjustmentsPlanPhasePercentageDiscountAdjustmentFilter]
type planVersionNewResponseAdjustmentsPlanPhasePercentageDiscountAdjustmentFilterJSON struct {
	Field       apijson.Field
	Operator    apijson.Field
	Values      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *PlanVersionNewResponseAdjustmentsPlanPhasePercentageDiscountAdjustmentFilter) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r planVersionNewResponseAdjustmentsPlanPhasePercentageDiscountAdjustmentFilterJSON) RawJSON() string {
	return r.raw
}

// The property of the price to filter on.
type PlanVersionNewResponseAdjustmentsPlanPhasePercentageDiscountAdjustmentFiltersField string

const (
	PlanVersionNewResponseAdjustmentsPlanPhasePercentageDiscountAdjustmentFiltersFieldPriceID       PlanVersionNewResponseAdjustmentsPlanPhasePercentageDiscountAdjustmentFiltersField = "price_id"
	PlanVersionNewResponseAdjustmentsPlanPhasePercentageDiscountAdjustmentFiltersFieldItemID        PlanVersionNewResponseAdjustmentsPlanPhasePercentageDiscountAdjustmentFiltersField = "item_id"
	PlanVersionNewResponseAdjustmentsPlanPhasePercentageDiscountAdjustmentFiltersFieldPriceType     PlanVersionNewResponseAdjustmentsPlanPhasePercentageDiscountAdjustmentFiltersField = "price_type"
	PlanVersionNewResponseAdjustmentsPlanPhasePercentageDiscountAdjustmentFiltersFieldCurrency      PlanVersionNewResponseAdjustmentsPlanPhasePercentageDiscountAdjustmentFiltersField = "currency"
	PlanVersionNewResponseAdjustmentsPlanPhasePercentageDiscountAdjustmentFiltersFieldPricingUnitID PlanVersionNewResponseAdjustmentsPlanPhasePercentageDiscountAdjustmentFiltersField = "pricing_unit_id"
)

func (r PlanVersionNewResponseAdjustmentsPlanPhasePercentageDiscountAdjustmentFiltersField) IsKnown() bool {
	switch r {
	case PlanVersionNewResponseAdjustmentsPlanPhasePercentageDiscountAdjustmentFiltersFieldPriceID, PlanVersionNewResponseAdjustmentsPlanPhasePercentageDiscountAdjustmentFiltersFieldItemID, PlanVersionNewResponseAdjustmentsPlanPhasePercentageDiscountAdjustmentFiltersFieldPriceType, PlanVersionNewResponseAdjustmentsPlanPhasePercentageDiscountAdjustmentFiltersFieldCurrency, PlanVersionNewResponseAdjustmentsPlanPhasePercentageDiscountAdjustmentFiltersFieldPricingUnitID:
		return true
	}
	return false
}

// Should prices that match the filter be included or excluded.
type PlanVersionNewResponseAdjustmentsPlanPhasePercentageDiscountAdjustmentFiltersOperator string

const (
	PlanVersionNewResponseAdjustmentsPlanPhasePercentageDiscountAdjustmentFiltersOperatorIncludes PlanVersionNewResponseAdjustmentsPlanPhasePercentageDiscountAdjustmentFiltersOperator = "includes"
	PlanVersionNewResponseAdjustmentsPlanPhasePercentageDiscountAdjustmentFiltersOperatorExcludes PlanVersionNewResponseAdjustmentsPlanPhasePercentageDiscountAdjustmentFiltersOperator = "excludes"
)

func (r PlanVersionNewResponseAdjustmentsPlanPhasePercentageDiscountAdjustmentFiltersOperator) IsKnown() bool {
	switch r {
	case PlanVersionNewResponseAdjustmentsPlanPhasePercentageDiscountAdjustmentFiltersOperatorIncludes, PlanVersionNewResponseAdjustmentsPlanPhasePercentageDiscountAdjustmentFiltersOperatorExcludes:
		return true
	}
	return false
}

type PlanVersionNewResponseAdjustmentsPlanPhaseMinimumAdjustment struct {
	ID             string                                                                    `json:"id,required"`
	AdjustmentType PlanVersionNewResponseAdjustmentsPlanPhaseMinimumAdjustmentAdjustmentType `json:"adjustment_type,required"`
	// The price IDs that this adjustment applies to.
	//
	// Deprecated: deprecated
	AppliesToPriceIDs []string `json:"applies_to_price_ids,required"`
	// The filters that determine which prices to apply this adjustment to.
	Filters []PlanVersionNewResponseAdjustmentsPlanPhaseMinimumAdjustmentFilter `json:"filters,required"`
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
	Reason string                                                          `json:"reason,required,nullable"`
	JSON   planVersionNewResponseAdjustmentsPlanPhaseMinimumAdjustmentJSON `json:"-"`
}

// planVersionNewResponseAdjustmentsPlanPhaseMinimumAdjustmentJSON contains the
// JSON metadata for the struct
// [PlanVersionNewResponseAdjustmentsPlanPhaseMinimumAdjustment]
type planVersionNewResponseAdjustmentsPlanPhaseMinimumAdjustmentJSON struct {
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

func (r *PlanVersionNewResponseAdjustmentsPlanPhaseMinimumAdjustment) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r planVersionNewResponseAdjustmentsPlanPhaseMinimumAdjustmentJSON) RawJSON() string {
	return r.raw
}

func (r PlanVersionNewResponseAdjustmentsPlanPhaseMinimumAdjustment) implementsPlanVersionNewResponseAdjustment() {
}

type PlanVersionNewResponseAdjustmentsPlanPhaseMinimumAdjustmentAdjustmentType string

const (
	PlanVersionNewResponseAdjustmentsPlanPhaseMinimumAdjustmentAdjustmentTypeMinimum PlanVersionNewResponseAdjustmentsPlanPhaseMinimumAdjustmentAdjustmentType = "minimum"
)

func (r PlanVersionNewResponseAdjustmentsPlanPhaseMinimumAdjustmentAdjustmentType) IsKnown() bool {
	switch r {
	case PlanVersionNewResponseAdjustmentsPlanPhaseMinimumAdjustmentAdjustmentTypeMinimum:
		return true
	}
	return false
}

type PlanVersionNewResponseAdjustmentsPlanPhaseMinimumAdjustmentFilter struct {
	// The property of the price to filter on.
	Field PlanVersionNewResponseAdjustmentsPlanPhaseMinimumAdjustmentFiltersField `json:"field,required"`
	// Should prices that match the filter be included or excluded.
	Operator PlanVersionNewResponseAdjustmentsPlanPhaseMinimumAdjustmentFiltersOperator `json:"operator,required"`
	// The IDs or values that match this filter.
	Values []string                                                              `json:"values,required"`
	JSON   planVersionNewResponseAdjustmentsPlanPhaseMinimumAdjustmentFilterJSON `json:"-"`
}

// planVersionNewResponseAdjustmentsPlanPhaseMinimumAdjustmentFilterJSON contains
// the JSON metadata for the struct
// [PlanVersionNewResponseAdjustmentsPlanPhaseMinimumAdjustmentFilter]
type planVersionNewResponseAdjustmentsPlanPhaseMinimumAdjustmentFilterJSON struct {
	Field       apijson.Field
	Operator    apijson.Field
	Values      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *PlanVersionNewResponseAdjustmentsPlanPhaseMinimumAdjustmentFilter) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r planVersionNewResponseAdjustmentsPlanPhaseMinimumAdjustmentFilterJSON) RawJSON() string {
	return r.raw
}

// The property of the price to filter on.
type PlanVersionNewResponseAdjustmentsPlanPhaseMinimumAdjustmentFiltersField string

const (
	PlanVersionNewResponseAdjustmentsPlanPhaseMinimumAdjustmentFiltersFieldPriceID       PlanVersionNewResponseAdjustmentsPlanPhaseMinimumAdjustmentFiltersField = "price_id"
	PlanVersionNewResponseAdjustmentsPlanPhaseMinimumAdjustmentFiltersFieldItemID        PlanVersionNewResponseAdjustmentsPlanPhaseMinimumAdjustmentFiltersField = "item_id"
	PlanVersionNewResponseAdjustmentsPlanPhaseMinimumAdjustmentFiltersFieldPriceType     PlanVersionNewResponseAdjustmentsPlanPhaseMinimumAdjustmentFiltersField = "price_type"
	PlanVersionNewResponseAdjustmentsPlanPhaseMinimumAdjustmentFiltersFieldCurrency      PlanVersionNewResponseAdjustmentsPlanPhaseMinimumAdjustmentFiltersField = "currency"
	PlanVersionNewResponseAdjustmentsPlanPhaseMinimumAdjustmentFiltersFieldPricingUnitID PlanVersionNewResponseAdjustmentsPlanPhaseMinimumAdjustmentFiltersField = "pricing_unit_id"
)

func (r PlanVersionNewResponseAdjustmentsPlanPhaseMinimumAdjustmentFiltersField) IsKnown() bool {
	switch r {
	case PlanVersionNewResponseAdjustmentsPlanPhaseMinimumAdjustmentFiltersFieldPriceID, PlanVersionNewResponseAdjustmentsPlanPhaseMinimumAdjustmentFiltersFieldItemID, PlanVersionNewResponseAdjustmentsPlanPhaseMinimumAdjustmentFiltersFieldPriceType, PlanVersionNewResponseAdjustmentsPlanPhaseMinimumAdjustmentFiltersFieldCurrency, PlanVersionNewResponseAdjustmentsPlanPhaseMinimumAdjustmentFiltersFieldPricingUnitID:
		return true
	}
	return false
}

// Should prices that match the filter be included or excluded.
type PlanVersionNewResponseAdjustmentsPlanPhaseMinimumAdjustmentFiltersOperator string

const (
	PlanVersionNewResponseAdjustmentsPlanPhaseMinimumAdjustmentFiltersOperatorIncludes PlanVersionNewResponseAdjustmentsPlanPhaseMinimumAdjustmentFiltersOperator = "includes"
	PlanVersionNewResponseAdjustmentsPlanPhaseMinimumAdjustmentFiltersOperatorExcludes PlanVersionNewResponseAdjustmentsPlanPhaseMinimumAdjustmentFiltersOperator = "excludes"
)

func (r PlanVersionNewResponseAdjustmentsPlanPhaseMinimumAdjustmentFiltersOperator) IsKnown() bool {
	switch r {
	case PlanVersionNewResponseAdjustmentsPlanPhaseMinimumAdjustmentFiltersOperatorIncludes, PlanVersionNewResponseAdjustmentsPlanPhaseMinimumAdjustmentFiltersOperatorExcludes:
		return true
	}
	return false
}

type PlanVersionNewResponseAdjustmentsPlanPhaseMaximumAdjustment struct {
	ID             string                                                                    `json:"id,required"`
	AdjustmentType PlanVersionNewResponseAdjustmentsPlanPhaseMaximumAdjustmentAdjustmentType `json:"adjustment_type,required"`
	// The price IDs that this adjustment applies to.
	//
	// Deprecated: deprecated
	AppliesToPriceIDs []string `json:"applies_to_price_ids,required"`
	// The filters that determine which prices to apply this adjustment to.
	Filters []PlanVersionNewResponseAdjustmentsPlanPhaseMaximumAdjustmentFilter `json:"filters,required"`
	// True for adjustments that apply to an entire invocice, false for adjustments
	// that apply to only one price.
	IsInvoiceLevel bool `json:"is_invoice_level,required"`
	// The maximum amount to charge in a given billing period for the prices this
	// adjustment applies to.
	MaximumAmount string `json:"maximum_amount,required"`
	// The plan phase in which this adjustment is active.
	PlanPhaseOrder int64 `json:"plan_phase_order,required,nullable"`
	// The reason for the adjustment.
	Reason string                                                          `json:"reason,required,nullable"`
	JSON   planVersionNewResponseAdjustmentsPlanPhaseMaximumAdjustmentJSON `json:"-"`
}

// planVersionNewResponseAdjustmentsPlanPhaseMaximumAdjustmentJSON contains the
// JSON metadata for the struct
// [PlanVersionNewResponseAdjustmentsPlanPhaseMaximumAdjustment]
type planVersionNewResponseAdjustmentsPlanPhaseMaximumAdjustmentJSON struct {
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

func (r *PlanVersionNewResponseAdjustmentsPlanPhaseMaximumAdjustment) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r planVersionNewResponseAdjustmentsPlanPhaseMaximumAdjustmentJSON) RawJSON() string {
	return r.raw
}

func (r PlanVersionNewResponseAdjustmentsPlanPhaseMaximumAdjustment) implementsPlanVersionNewResponseAdjustment() {
}

type PlanVersionNewResponseAdjustmentsPlanPhaseMaximumAdjustmentAdjustmentType string

const (
	PlanVersionNewResponseAdjustmentsPlanPhaseMaximumAdjustmentAdjustmentTypeMaximum PlanVersionNewResponseAdjustmentsPlanPhaseMaximumAdjustmentAdjustmentType = "maximum"
)

func (r PlanVersionNewResponseAdjustmentsPlanPhaseMaximumAdjustmentAdjustmentType) IsKnown() bool {
	switch r {
	case PlanVersionNewResponseAdjustmentsPlanPhaseMaximumAdjustmentAdjustmentTypeMaximum:
		return true
	}
	return false
}

type PlanVersionNewResponseAdjustmentsPlanPhaseMaximumAdjustmentFilter struct {
	// The property of the price to filter on.
	Field PlanVersionNewResponseAdjustmentsPlanPhaseMaximumAdjustmentFiltersField `json:"field,required"`
	// Should prices that match the filter be included or excluded.
	Operator PlanVersionNewResponseAdjustmentsPlanPhaseMaximumAdjustmentFiltersOperator `json:"operator,required"`
	// The IDs or values that match this filter.
	Values []string                                                              `json:"values,required"`
	JSON   planVersionNewResponseAdjustmentsPlanPhaseMaximumAdjustmentFilterJSON `json:"-"`
}

// planVersionNewResponseAdjustmentsPlanPhaseMaximumAdjustmentFilterJSON contains
// the JSON metadata for the struct
// [PlanVersionNewResponseAdjustmentsPlanPhaseMaximumAdjustmentFilter]
type planVersionNewResponseAdjustmentsPlanPhaseMaximumAdjustmentFilterJSON struct {
	Field       apijson.Field
	Operator    apijson.Field
	Values      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *PlanVersionNewResponseAdjustmentsPlanPhaseMaximumAdjustmentFilter) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r planVersionNewResponseAdjustmentsPlanPhaseMaximumAdjustmentFilterJSON) RawJSON() string {
	return r.raw
}

// The property of the price to filter on.
type PlanVersionNewResponseAdjustmentsPlanPhaseMaximumAdjustmentFiltersField string

const (
	PlanVersionNewResponseAdjustmentsPlanPhaseMaximumAdjustmentFiltersFieldPriceID       PlanVersionNewResponseAdjustmentsPlanPhaseMaximumAdjustmentFiltersField = "price_id"
	PlanVersionNewResponseAdjustmentsPlanPhaseMaximumAdjustmentFiltersFieldItemID        PlanVersionNewResponseAdjustmentsPlanPhaseMaximumAdjustmentFiltersField = "item_id"
	PlanVersionNewResponseAdjustmentsPlanPhaseMaximumAdjustmentFiltersFieldPriceType     PlanVersionNewResponseAdjustmentsPlanPhaseMaximumAdjustmentFiltersField = "price_type"
	PlanVersionNewResponseAdjustmentsPlanPhaseMaximumAdjustmentFiltersFieldCurrency      PlanVersionNewResponseAdjustmentsPlanPhaseMaximumAdjustmentFiltersField = "currency"
	PlanVersionNewResponseAdjustmentsPlanPhaseMaximumAdjustmentFiltersFieldPricingUnitID PlanVersionNewResponseAdjustmentsPlanPhaseMaximumAdjustmentFiltersField = "pricing_unit_id"
)

func (r PlanVersionNewResponseAdjustmentsPlanPhaseMaximumAdjustmentFiltersField) IsKnown() bool {
	switch r {
	case PlanVersionNewResponseAdjustmentsPlanPhaseMaximumAdjustmentFiltersFieldPriceID, PlanVersionNewResponseAdjustmentsPlanPhaseMaximumAdjustmentFiltersFieldItemID, PlanVersionNewResponseAdjustmentsPlanPhaseMaximumAdjustmentFiltersFieldPriceType, PlanVersionNewResponseAdjustmentsPlanPhaseMaximumAdjustmentFiltersFieldCurrency, PlanVersionNewResponseAdjustmentsPlanPhaseMaximumAdjustmentFiltersFieldPricingUnitID:
		return true
	}
	return false
}

// Should prices that match the filter be included or excluded.
type PlanVersionNewResponseAdjustmentsPlanPhaseMaximumAdjustmentFiltersOperator string

const (
	PlanVersionNewResponseAdjustmentsPlanPhaseMaximumAdjustmentFiltersOperatorIncludes PlanVersionNewResponseAdjustmentsPlanPhaseMaximumAdjustmentFiltersOperator = "includes"
	PlanVersionNewResponseAdjustmentsPlanPhaseMaximumAdjustmentFiltersOperatorExcludes PlanVersionNewResponseAdjustmentsPlanPhaseMaximumAdjustmentFiltersOperator = "excludes"
)

func (r PlanVersionNewResponseAdjustmentsPlanPhaseMaximumAdjustmentFiltersOperator) IsKnown() bool {
	switch r {
	case PlanVersionNewResponseAdjustmentsPlanPhaseMaximumAdjustmentFiltersOperatorIncludes, PlanVersionNewResponseAdjustmentsPlanPhaseMaximumAdjustmentFiltersOperatorExcludes:
		return true
	}
	return false
}

type PlanVersionNewResponseAdjustmentsAdjustmentType string

const (
	PlanVersionNewResponseAdjustmentsAdjustmentTypeUsageDiscount      PlanVersionNewResponseAdjustmentsAdjustmentType = "usage_discount"
	PlanVersionNewResponseAdjustmentsAdjustmentTypeAmountDiscount     PlanVersionNewResponseAdjustmentsAdjustmentType = "amount_discount"
	PlanVersionNewResponseAdjustmentsAdjustmentTypePercentageDiscount PlanVersionNewResponseAdjustmentsAdjustmentType = "percentage_discount"
	PlanVersionNewResponseAdjustmentsAdjustmentTypeMinimum            PlanVersionNewResponseAdjustmentsAdjustmentType = "minimum"
	PlanVersionNewResponseAdjustmentsAdjustmentTypeMaximum            PlanVersionNewResponseAdjustmentsAdjustmentType = "maximum"
)

func (r PlanVersionNewResponseAdjustmentsAdjustmentType) IsKnown() bool {
	switch r {
	case PlanVersionNewResponseAdjustmentsAdjustmentTypeUsageDiscount, PlanVersionNewResponseAdjustmentsAdjustmentTypeAmountDiscount, PlanVersionNewResponseAdjustmentsAdjustmentTypePercentageDiscount, PlanVersionNewResponseAdjustmentsAdjustmentTypeMinimum, PlanVersionNewResponseAdjustmentsAdjustmentTypeMaximum:
		return true
	}
	return false
}

type PlanVersionNewResponsePlanPhase struct {
	ID          string `json:"id,required"`
	Description string `json:"description,required,nullable"`
	// How many terms of length `duration_unit` this phase is active for. If null, this
	// phase is evergreen and active indefinitely
	Duration     int64                                        `json:"duration,required,nullable"`
	DurationUnit PlanVersionNewResponsePlanPhasesDurationUnit `json:"duration_unit,required,nullable"`
	Name         string                                       `json:"name,required"`
	// Determines the ordering of the phase in a plan's lifecycle. 1 = first phase.
	Order int64                               `json:"order,required"`
	JSON  planVersionNewResponsePlanPhaseJSON `json:"-"`
}

// planVersionNewResponsePlanPhaseJSON contains the JSON metadata for the struct
// [PlanVersionNewResponsePlanPhase]
type planVersionNewResponsePlanPhaseJSON struct {
	ID           apijson.Field
	Description  apijson.Field
	Duration     apijson.Field
	DurationUnit apijson.Field
	Name         apijson.Field
	Order        apijson.Field
	raw          string
	ExtraFields  map[string]apijson.Field
}

func (r *PlanVersionNewResponsePlanPhase) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r planVersionNewResponsePlanPhaseJSON) RawJSON() string {
	return r.raw
}

type PlanVersionNewResponsePlanPhasesDurationUnit string

const (
	PlanVersionNewResponsePlanPhasesDurationUnitDaily      PlanVersionNewResponsePlanPhasesDurationUnit = "daily"
	PlanVersionNewResponsePlanPhasesDurationUnitMonthly    PlanVersionNewResponsePlanPhasesDurationUnit = "monthly"
	PlanVersionNewResponsePlanPhasesDurationUnitQuarterly  PlanVersionNewResponsePlanPhasesDurationUnit = "quarterly"
	PlanVersionNewResponsePlanPhasesDurationUnitSemiAnnual PlanVersionNewResponsePlanPhasesDurationUnit = "semi_annual"
	PlanVersionNewResponsePlanPhasesDurationUnitAnnual     PlanVersionNewResponsePlanPhasesDurationUnit = "annual"
)

func (r PlanVersionNewResponsePlanPhasesDurationUnit) IsKnown() bool {
	switch r {
	case PlanVersionNewResponsePlanPhasesDurationUnitDaily, PlanVersionNewResponsePlanPhasesDurationUnitMonthly, PlanVersionNewResponsePlanPhasesDurationUnitQuarterly, PlanVersionNewResponsePlanPhasesDurationUnitSemiAnnual, PlanVersionNewResponsePlanPhasesDurationUnitAnnual:
		return true
	}
	return false
}

// The PlanVersion resource represents the prices and adjustments present on a
// specific version of a plan.
type PlanVersionGetResponse struct {
	// Adjustments for this plan. If the plan has phases, this includes adjustments
	// across all phases of the plan.
	Adjustments []PlanVersionGetResponseAdjustment `json:"adjustments,required"`
	CreatedAt   time.Time                          `json:"created_at,required" format:"date-time"`
	PlanPhases  []PlanVersionGetResponsePlanPhase  `json:"plan_phases,required,nullable"`
	// Prices for this plan. If the plan has phases, this includes prices across all
	// phases of the plan.
	Prices  []Price                    `json:"prices,required"`
	Version int64                      `json:"version,required"`
	JSON    planVersionGetResponseJSON `json:"-"`
}

// planVersionGetResponseJSON contains the JSON metadata for the struct
// [PlanVersionGetResponse]
type planVersionGetResponseJSON struct {
	Adjustments apijson.Field
	CreatedAt   apijson.Field
	PlanPhases  apijson.Field
	Prices      apijson.Field
	Version     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *PlanVersionGetResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r planVersionGetResponseJSON) RawJSON() string {
	return r.raw
}

type PlanVersionGetResponseAdjustment struct {
	ID             string                                          `json:"id,required"`
	AdjustmentType PlanVersionGetResponseAdjustmentsAdjustmentType `json:"adjustment_type,required"`
	// This field can have the runtime type of [[]string].
	AppliesToPriceIDs interface{} `json:"applies_to_price_ids,required"`
	// This field can have the runtime type of
	// [[]PlanVersionGetResponseAdjustmentsPlanPhaseUsageDiscountAdjustmentFilter],
	// [[]PlanVersionGetResponseAdjustmentsPlanPhaseAmountDiscountAdjustmentFilter],
	// [[]PlanVersionGetResponseAdjustmentsPlanPhasePercentageDiscountAdjustmentFilter],
	// [[]PlanVersionGetResponseAdjustmentsPlanPhaseMinimumAdjustmentFilter],
	// [[]PlanVersionGetResponseAdjustmentsPlanPhaseMaximumAdjustmentFilter].
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
	UsageDiscount float64                              `json:"usage_discount"`
	JSON          planVersionGetResponseAdjustmentJSON `json:"-"`
	union         PlanVersionGetResponseAdjustmentsUnion
}

// planVersionGetResponseAdjustmentJSON contains the JSON metadata for the struct
// [PlanVersionGetResponseAdjustment]
type planVersionGetResponseAdjustmentJSON struct {
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

func (r planVersionGetResponseAdjustmentJSON) RawJSON() string {
	return r.raw
}

func (r *PlanVersionGetResponseAdjustment) UnmarshalJSON(data []byte) (err error) {
	*r = PlanVersionGetResponseAdjustment{}
	err = apijson.UnmarshalRoot(data, &r.union)
	if err != nil {
		return err
	}
	return apijson.Port(r.union, &r)
}

// AsUnion returns a [PlanVersionGetResponseAdjustmentsUnion] interface which you
// can cast to the specific types for more type safety.
//
// Possible runtime types of the union are
// [PlanVersionGetResponseAdjustmentsPlanPhaseUsageDiscountAdjustment],
// [PlanVersionGetResponseAdjustmentsPlanPhaseAmountDiscountAdjustment],
// [PlanVersionGetResponseAdjustmentsPlanPhasePercentageDiscountAdjustment],
// [PlanVersionGetResponseAdjustmentsPlanPhaseMinimumAdjustment],
// [PlanVersionGetResponseAdjustmentsPlanPhaseMaximumAdjustment].
func (r PlanVersionGetResponseAdjustment) AsUnion() PlanVersionGetResponseAdjustmentsUnion {
	return r.union
}

// Union satisfied by
// [PlanVersionGetResponseAdjustmentsPlanPhaseUsageDiscountAdjustment],
// [PlanVersionGetResponseAdjustmentsPlanPhaseAmountDiscountAdjustment],
// [PlanVersionGetResponseAdjustmentsPlanPhasePercentageDiscountAdjustment],
// [PlanVersionGetResponseAdjustmentsPlanPhaseMinimumAdjustment] or
// [PlanVersionGetResponseAdjustmentsPlanPhaseMaximumAdjustment].
type PlanVersionGetResponseAdjustmentsUnion interface {
	implementsPlanVersionGetResponseAdjustment()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*PlanVersionGetResponseAdjustmentsUnion)(nil)).Elem(),
		"adjustment_type",
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(PlanVersionGetResponseAdjustmentsPlanPhaseUsageDiscountAdjustment{}),
			DiscriminatorValue: "usage_discount",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(PlanVersionGetResponseAdjustmentsPlanPhaseAmountDiscountAdjustment{}),
			DiscriminatorValue: "amount_discount",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(PlanVersionGetResponseAdjustmentsPlanPhasePercentageDiscountAdjustment{}),
			DiscriminatorValue: "percentage_discount",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(PlanVersionGetResponseAdjustmentsPlanPhaseMinimumAdjustment{}),
			DiscriminatorValue: "minimum",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(PlanVersionGetResponseAdjustmentsPlanPhaseMaximumAdjustment{}),
			DiscriminatorValue: "maximum",
		},
	)
}

type PlanVersionGetResponseAdjustmentsPlanPhaseUsageDiscountAdjustment struct {
	ID             string                                                                          `json:"id,required"`
	AdjustmentType PlanVersionGetResponseAdjustmentsPlanPhaseUsageDiscountAdjustmentAdjustmentType `json:"adjustment_type,required"`
	// The price IDs that this adjustment applies to.
	//
	// Deprecated: deprecated
	AppliesToPriceIDs []string `json:"applies_to_price_ids,required"`
	// The filters that determine which prices to apply this adjustment to.
	Filters []PlanVersionGetResponseAdjustmentsPlanPhaseUsageDiscountAdjustmentFilter `json:"filters,required"`
	// True for adjustments that apply to an entire invocice, false for adjustments
	// that apply to only one price.
	IsInvoiceLevel bool `json:"is_invoice_level,required"`
	// The plan phase in which this adjustment is active.
	PlanPhaseOrder int64 `json:"plan_phase_order,required,nullable"`
	// The reason for the adjustment.
	Reason string `json:"reason,required,nullable"`
	// The number of usage units by which to discount the price this adjustment applies
	// to in a given billing period.
	UsageDiscount float64                                                               `json:"usage_discount,required"`
	JSON          planVersionGetResponseAdjustmentsPlanPhaseUsageDiscountAdjustmentJSON `json:"-"`
}

// planVersionGetResponseAdjustmentsPlanPhaseUsageDiscountAdjustmentJSON contains
// the JSON metadata for the struct
// [PlanVersionGetResponseAdjustmentsPlanPhaseUsageDiscountAdjustment]
type planVersionGetResponseAdjustmentsPlanPhaseUsageDiscountAdjustmentJSON struct {
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

func (r *PlanVersionGetResponseAdjustmentsPlanPhaseUsageDiscountAdjustment) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r planVersionGetResponseAdjustmentsPlanPhaseUsageDiscountAdjustmentJSON) RawJSON() string {
	return r.raw
}

func (r PlanVersionGetResponseAdjustmentsPlanPhaseUsageDiscountAdjustment) implementsPlanVersionGetResponseAdjustment() {
}

type PlanVersionGetResponseAdjustmentsPlanPhaseUsageDiscountAdjustmentAdjustmentType string

const (
	PlanVersionGetResponseAdjustmentsPlanPhaseUsageDiscountAdjustmentAdjustmentTypeUsageDiscount PlanVersionGetResponseAdjustmentsPlanPhaseUsageDiscountAdjustmentAdjustmentType = "usage_discount"
)

func (r PlanVersionGetResponseAdjustmentsPlanPhaseUsageDiscountAdjustmentAdjustmentType) IsKnown() bool {
	switch r {
	case PlanVersionGetResponseAdjustmentsPlanPhaseUsageDiscountAdjustmentAdjustmentTypeUsageDiscount:
		return true
	}
	return false
}

type PlanVersionGetResponseAdjustmentsPlanPhaseUsageDiscountAdjustmentFilter struct {
	// The property of the price to filter on.
	Field PlanVersionGetResponseAdjustmentsPlanPhaseUsageDiscountAdjustmentFiltersField `json:"field,required"`
	// Should prices that match the filter be included or excluded.
	Operator PlanVersionGetResponseAdjustmentsPlanPhaseUsageDiscountAdjustmentFiltersOperator `json:"operator,required"`
	// The IDs or values that match this filter.
	Values []string                                                                    `json:"values,required"`
	JSON   planVersionGetResponseAdjustmentsPlanPhaseUsageDiscountAdjustmentFilterJSON `json:"-"`
}

// planVersionGetResponseAdjustmentsPlanPhaseUsageDiscountAdjustmentFilterJSON
// contains the JSON metadata for the struct
// [PlanVersionGetResponseAdjustmentsPlanPhaseUsageDiscountAdjustmentFilter]
type planVersionGetResponseAdjustmentsPlanPhaseUsageDiscountAdjustmentFilterJSON struct {
	Field       apijson.Field
	Operator    apijson.Field
	Values      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *PlanVersionGetResponseAdjustmentsPlanPhaseUsageDiscountAdjustmentFilter) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r planVersionGetResponseAdjustmentsPlanPhaseUsageDiscountAdjustmentFilterJSON) RawJSON() string {
	return r.raw
}

// The property of the price to filter on.
type PlanVersionGetResponseAdjustmentsPlanPhaseUsageDiscountAdjustmentFiltersField string

const (
	PlanVersionGetResponseAdjustmentsPlanPhaseUsageDiscountAdjustmentFiltersFieldPriceID       PlanVersionGetResponseAdjustmentsPlanPhaseUsageDiscountAdjustmentFiltersField = "price_id"
	PlanVersionGetResponseAdjustmentsPlanPhaseUsageDiscountAdjustmentFiltersFieldItemID        PlanVersionGetResponseAdjustmentsPlanPhaseUsageDiscountAdjustmentFiltersField = "item_id"
	PlanVersionGetResponseAdjustmentsPlanPhaseUsageDiscountAdjustmentFiltersFieldPriceType     PlanVersionGetResponseAdjustmentsPlanPhaseUsageDiscountAdjustmentFiltersField = "price_type"
	PlanVersionGetResponseAdjustmentsPlanPhaseUsageDiscountAdjustmentFiltersFieldCurrency      PlanVersionGetResponseAdjustmentsPlanPhaseUsageDiscountAdjustmentFiltersField = "currency"
	PlanVersionGetResponseAdjustmentsPlanPhaseUsageDiscountAdjustmentFiltersFieldPricingUnitID PlanVersionGetResponseAdjustmentsPlanPhaseUsageDiscountAdjustmentFiltersField = "pricing_unit_id"
)

func (r PlanVersionGetResponseAdjustmentsPlanPhaseUsageDiscountAdjustmentFiltersField) IsKnown() bool {
	switch r {
	case PlanVersionGetResponseAdjustmentsPlanPhaseUsageDiscountAdjustmentFiltersFieldPriceID, PlanVersionGetResponseAdjustmentsPlanPhaseUsageDiscountAdjustmentFiltersFieldItemID, PlanVersionGetResponseAdjustmentsPlanPhaseUsageDiscountAdjustmentFiltersFieldPriceType, PlanVersionGetResponseAdjustmentsPlanPhaseUsageDiscountAdjustmentFiltersFieldCurrency, PlanVersionGetResponseAdjustmentsPlanPhaseUsageDiscountAdjustmentFiltersFieldPricingUnitID:
		return true
	}
	return false
}

// Should prices that match the filter be included or excluded.
type PlanVersionGetResponseAdjustmentsPlanPhaseUsageDiscountAdjustmentFiltersOperator string

const (
	PlanVersionGetResponseAdjustmentsPlanPhaseUsageDiscountAdjustmentFiltersOperatorIncludes PlanVersionGetResponseAdjustmentsPlanPhaseUsageDiscountAdjustmentFiltersOperator = "includes"
	PlanVersionGetResponseAdjustmentsPlanPhaseUsageDiscountAdjustmentFiltersOperatorExcludes PlanVersionGetResponseAdjustmentsPlanPhaseUsageDiscountAdjustmentFiltersOperator = "excludes"
)

func (r PlanVersionGetResponseAdjustmentsPlanPhaseUsageDiscountAdjustmentFiltersOperator) IsKnown() bool {
	switch r {
	case PlanVersionGetResponseAdjustmentsPlanPhaseUsageDiscountAdjustmentFiltersOperatorIncludes, PlanVersionGetResponseAdjustmentsPlanPhaseUsageDiscountAdjustmentFiltersOperatorExcludes:
		return true
	}
	return false
}

type PlanVersionGetResponseAdjustmentsPlanPhaseAmountDiscountAdjustment struct {
	ID             string                                                                           `json:"id,required"`
	AdjustmentType PlanVersionGetResponseAdjustmentsPlanPhaseAmountDiscountAdjustmentAdjustmentType `json:"adjustment_type,required"`
	// The amount by which to discount the prices this adjustment applies to in a given
	// billing period.
	AmountDiscount string `json:"amount_discount,required"`
	// The price IDs that this adjustment applies to.
	//
	// Deprecated: deprecated
	AppliesToPriceIDs []string `json:"applies_to_price_ids,required"`
	// The filters that determine which prices to apply this adjustment to.
	Filters []PlanVersionGetResponseAdjustmentsPlanPhaseAmountDiscountAdjustmentFilter `json:"filters,required"`
	// True for adjustments that apply to an entire invocice, false for adjustments
	// that apply to only one price.
	IsInvoiceLevel bool `json:"is_invoice_level,required"`
	// The plan phase in which this adjustment is active.
	PlanPhaseOrder int64 `json:"plan_phase_order,required,nullable"`
	// The reason for the adjustment.
	Reason string                                                                 `json:"reason,required,nullable"`
	JSON   planVersionGetResponseAdjustmentsPlanPhaseAmountDiscountAdjustmentJSON `json:"-"`
}

// planVersionGetResponseAdjustmentsPlanPhaseAmountDiscountAdjustmentJSON contains
// the JSON metadata for the struct
// [PlanVersionGetResponseAdjustmentsPlanPhaseAmountDiscountAdjustment]
type planVersionGetResponseAdjustmentsPlanPhaseAmountDiscountAdjustmentJSON struct {
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

func (r *PlanVersionGetResponseAdjustmentsPlanPhaseAmountDiscountAdjustment) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r planVersionGetResponseAdjustmentsPlanPhaseAmountDiscountAdjustmentJSON) RawJSON() string {
	return r.raw
}

func (r PlanVersionGetResponseAdjustmentsPlanPhaseAmountDiscountAdjustment) implementsPlanVersionGetResponseAdjustment() {
}

type PlanVersionGetResponseAdjustmentsPlanPhaseAmountDiscountAdjustmentAdjustmentType string

const (
	PlanVersionGetResponseAdjustmentsPlanPhaseAmountDiscountAdjustmentAdjustmentTypeAmountDiscount PlanVersionGetResponseAdjustmentsPlanPhaseAmountDiscountAdjustmentAdjustmentType = "amount_discount"
)

func (r PlanVersionGetResponseAdjustmentsPlanPhaseAmountDiscountAdjustmentAdjustmentType) IsKnown() bool {
	switch r {
	case PlanVersionGetResponseAdjustmentsPlanPhaseAmountDiscountAdjustmentAdjustmentTypeAmountDiscount:
		return true
	}
	return false
}

type PlanVersionGetResponseAdjustmentsPlanPhaseAmountDiscountAdjustmentFilter struct {
	// The property of the price to filter on.
	Field PlanVersionGetResponseAdjustmentsPlanPhaseAmountDiscountAdjustmentFiltersField `json:"field,required"`
	// Should prices that match the filter be included or excluded.
	Operator PlanVersionGetResponseAdjustmentsPlanPhaseAmountDiscountAdjustmentFiltersOperator `json:"operator,required"`
	// The IDs or values that match this filter.
	Values []string                                                                     `json:"values,required"`
	JSON   planVersionGetResponseAdjustmentsPlanPhaseAmountDiscountAdjustmentFilterJSON `json:"-"`
}

// planVersionGetResponseAdjustmentsPlanPhaseAmountDiscountAdjustmentFilterJSON
// contains the JSON metadata for the struct
// [PlanVersionGetResponseAdjustmentsPlanPhaseAmountDiscountAdjustmentFilter]
type planVersionGetResponseAdjustmentsPlanPhaseAmountDiscountAdjustmentFilterJSON struct {
	Field       apijson.Field
	Operator    apijson.Field
	Values      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *PlanVersionGetResponseAdjustmentsPlanPhaseAmountDiscountAdjustmentFilter) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r planVersionGetResponseAdjustmentsPlanPhaseAmountDiscountAdjustmentFilterJSON) RawJSON() string {
	return r.raw
}

// The property of the price to filter on.
type PlanVersionGetResponseAdjustmentsPlanPhaseAmountDiscountAdjustmentFiltersField string

const (
	PlanVersionGetResponseAdjustmentsPlanPhaseAmountDiscountAdjustmentFiltersFieldPriceID       PlanVersionGetResponseAdjustmentsPlanPhaseAmountDiscountAdjustmentFiltersField = "price_id"
	PlanVersionGetResponseAdjustmentsPlanPhaseAmountDiscountAdjustmentFiltersFieldItemID        PlanVersionGetResponseAdjustmentsPlanPhaseAmountDiscountAdjustmentFiltersField = "item_id"
	PlanVersionGetResponseAdjustmentsPlanPhaseAmountDiscountAdjustmentFiltersFieldPriceType     PlanVersionGetResponseAdjustmentsPlanPhaseAmountDiscountAdjustmentFiltersField = "price_type"
	PlanVersionGetResponseAdjustmentsPlanPhaseAmountDiscountAdjustmentFiltersFieldCurrency      PlanVersionGetResponseAdjustmentsPlanPhaseAmountDiscountAdjustmentFiltersField = "currency"
	PlanVersionGetResponseAdjustmentsPlanPhaseAmountDiscountAdjustmentFiltersFieldPricingUnitID PlanVersionGetResponseAdjustmentsPlanPhaseAmountDiscountAdjustmentFiltersField = "pricing_unit_id"
)

func (r PlanVersionGetResponseAdjustmentsPlanPhaseAmountDiscountAdjustmentFiltersField) IsKnown() bool {
	switch r {
	case PlanVersionGetResponseAdjustmentsPlanPhaseAmountDiscountAdjustmentFiltersFieldPriceID, PlanVersionGetResponseAdjustmentsPlanPhaseAmountDiscountAdjustmentFiltersFieldItemID, PlanVersionGetResponseAdjustmentsPlanPhaseAmountDiscountAdjustmentFiltersFieldPriceType, PlanVersionGetResponseAdjustmentsPlanPhaseAmountDiscountAdjustmentFiltersFieldCurrency, PlanVersionGetResponseAdjustmentsPlanPhaseAmountDiscountAdjustmentFiltersFieldPricingUnitID:
		return true
	}
	return false
}

// Should prices that match the filter be included or excluded.
type PlanVersionGetResponseAdjustmentsPlanPhaseAmountDiscountAdjustmentFiltersOperator string

const (
	PlanVersionGetResponseAdjustmentsPlanPhaseAmountDiscountAdjustmentFiltersOperatorIncludes PlanVersionGetResponseAdjustmentsPlanPhaseAmountDiscountAdjustmentFiltersOperator = "includes"
	PlanVersionGetResponseAdjustmentsPlanPhaseAmountDiscountAdjustmentFiltersOperatorExcludes PlanVersionGetResponseAdjustmentsPlanPhaseAmountDiscountAdjustmentFiltersOperator = "excludes"
)

func (r PlanVersionGetResponseAdjustmentsPlanPhaseAmountDiscountAdjustmentFiltersOperator) IsKnown() bool {
	switch r {
	case PlanVersionGetResponseAdjustmentsPlanPhaseAmountDiscountAdjustmentFiltersOperatorIncludes, PlanVersionGetResponseAdjustmentsPlanPhaseAmountDiscountAdjustmentFiltersOperatorExcludes:
		return true
	}
	return false
}

type PlanVersionGetResponseAdjustmentsPlanPhasePercentageDiscountAdjustment struct {
	ID             string                                                                               `json:"id,required"`
	AdjustmentType PlanVersionGetResponseAdjustmentsPlanPhasePercentageDiscountAdjustmentAdjustmentType `json:"adjustment_type,required"`
	// The price IDs that this adjustment applies to.
	//
	// Deprecated: deprecated
	AppliesToPriceIDs []string `json:"applies_to_price_ids,required"`
	// The filters that determine which prices to apply this adjustment to.
	Filters []PlanVersionGetResponseAdjustmentsPlanPhasePercentageDiscountAdjustmentFilter `json:"filters,required"`
	// True for adjustments that apply to an entire invocice, false for adjustments
	// that apply to only one price.
	IsInvoiceLevel bool `json:"is_invoice_level,required"`
	// The percentage (as a value between 0 and 1) by which to discount the price
	// intervals this adjustment applies to in a given billing period.
	PercentageDiscount float64 `json:"percentage_discount,required"`
	// The plan phase in which this adjustment is active.
	PlanPhaseOrder int64 `json:"plan_phase_order,required,nullable"`
	// The reason for the adjustment.
	Reason string                                                                     `json:"reason,required,nullable"`
	JSON   planVersionGetResponseAdjustmentsPlanPhasePercentageDiscountAdjustmentJSON `json:"-"`
}

// planVersionGetResponseAdjustmentsPlanPhasePercentageDiscountAdjustmentJSON
// contains the JSON metadata for the struct
// [PlanVersionGetResponseAdjustmentsPlanPhasePercentageDiscountAdjustment]
type planVersionGetResponseAdjustmentsPlanPhasePercentageDiscountAdjustmentJSON struct {
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

func (r *PlanVersionGetResponseAdjustmentsPlanPhasePercentageDiscountAdjustment) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r planVersionGetResponseAdjustmentsPlanPhasePercentageDiscountAdjustmentJSON) RawJSON() string {
	return r.raw
}

func (r PlanVersionGetResponseAdjustmentsPlanPhasePercentageDiscountAdjustment) implementsPlanVersionGetResponseAdjustment() {
}

type PlanVersionGetResponseAdjustmentsPlanPhasePercentageDiscountAdjustmentAdjustmentType string

const (
	PlanVersionGetResponseAdjustmentsPlanPhasePercentageDiscountAdjustmentAdjustmentTypePercentageDiscount PlanVersionGetResponseAdjustmentsPlanPhasePercentageDiscountAdjustmentAdjustmentType = "percentage_discount"
)

func (r PlanVersionGetResponseAdjustmentsPlanPhasePercentageDiscountAdjustmentAdjustmentType) IsKnown() bool {
	switch r {
	case PlanVersionGetResponseAdjustmentsPlanPhasePercentageDiscountAdjustmentAdjustmentTypePercentageDiscount:
		return true
	}
	return false
}

type PlanVersionGetResponseAdjustmentsPlanPhasePercentageDiscountAdjustmentFilter struct {
	// The property of the price to filter on.
	Field PlanVersionGetResponseAdjustmentsPlanPhasePercentageDiscountAdjustmentFiltersField `json:"field,required"`
	// Should prices that match the filter be included or excluded.
	Operator PlanVersionGetResponseAdjustmentsPlanPhasePercentageDiscountAdjustmentFiltersOperator `json:"operator,required"`
	// The IDs or values that match this filter.
	Values []string                                                                         `json:"values,required"`
	JSON   planVersionGetResponseAdjustmentsPlanPhasePercentageDiscountAdjustmentFilterJSON `json:"-"`
}

// planVersionGetResponseAdjustmentsPlanPhasePercentageDiscountAdjustmentFilterJSON
// contains the JSON metadata for the struct
// [PlanVersionGetResponseAdjustmentsPlanPhasePercentageDiscountAdjustmentFilter]
type planVersionGetResponseAdjustmentsPlanPhasePercentageDiscountAdjustmentFilterJSON struct {
	Field       apijson.Field
	Operator    apijson.Field
	Values      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *PlanVersionGetResponseAdjustmentsPlanPhasePercentageDiscountAdjustmentFilter) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r planVersionGetResponseAdjustmentsPlanPhasePercentageDiscountAdjustmentFilterJSON) RawJSON() string {
	return r.raw
}

// The property of the price to filter on.
type PlanVersionGetResponseAdjustmentsPlanPhasePercentageDiscountAdjustmentFiltersField string

const (
	PlanVersionGetResponseAdjustmentsPlanPhasePercentageDiscountAdjustmentFiltersFieldPriceID       PlanVersionGetResponseAdjustmentsPlanPhasePercentageDiscountAdjustmentFiltersField = "price_id"
	PlanVersionGetResponseAdjustmentsPlanPhasePercentageDiscountAdjustmentFiltersFieldItemID        PlanVersionGetResponseAdjustmentsPlanPhasePercentageDiscountAdjustmentFiltersField = "item_id"
	PlanVersionGetResponseAdjustmentsPlanPhasePercentageDiscountAdjustmentFiltersFieldPriceType     PlanVersionGetResponseAdjustmentsPlanPhasePercentageDiscountAdjustmentFiltersField = "price_type"
	PlanVersionGetResponseAdjustmentsPlanPhasePercentageDiscountAdjustmentFiltersFieldCurrency      PlanVersionGetResponseAdjustmentsPlanPhasePercentageDiscountAdjustmentFiltersField = "currency"
	PlanVersionGetResponseAdjustmentsPlanPhasePercentageDiscountAdjustmentFiltersFieldPricingUnitID PlanVersionGetResponseAdjustmentsPlanPhasePercentageDiscountAdjustmentFiltersField = "pricing_unit_id"
)

func (r PlanVersionGetResponseAdjustmentsPlanPhasePercentageDiscountAdjustmentFiltersField) IsKnown() bool {
	switch r {
	case PlanVersionGetResponseAdjustmentsPlanPhasePercentageDiscountAdjustmentFiltersFieldPriceID, PlanVersionGetResponseAdjustmentsPlanPhasePercentageDiscountAdjustmentFiltersFieldItemID, PlanVersionGetResponseAdjustmentsPlanPhasePercentageDiscountAdjustmentFiltersFieldPriceType, PlanVersionGetResponseAdjustmentsPlanPhasePercentageDiscountAdjustmentFiltersFieldCurrency, PlanVersionGetResponseAdjustmentsPlanPhasePercentageDiscountAdjustmentFiltersFieldPricingUnitID:
		return true
	}
	return false
}

// Should prices that match the filter be included or excluded.
type PlanVersionGetResponseAdjustmentsPlanPhasePercentageDiscountAdjustmentFiltersOperator string

const (
	PlanVersionGetResponseAdjustmentsPlanPhasePercentageDiscountAdjustmentFiltersOperatorIncludes PlanVersionGetResponseAdjustmentsPlanPhasePercentageDiscountAdjustmentFiltersOperator = "includes"
	PlanVersionGetResponseAdjustmentsPlanPhasePercentageDiscountAdjustmentFiltersOperatorExcludes PlanVersionGetResponseAdjustmentsPlanPhasePercentageDiscountAdjustmentFiltersOperator = "excludes"
)

func (r PlanVersionGetResponseAdjustmentsPlanPhasePercentageDiscountAdjustmentFiltersOperator) IsKnown() bool {
	switch r {
	case PlanVersionGetResponseAdjustmentsPlanPhasePercentageDiscountAdjustmentFiltersOperatorIncludes, PlanVersionGetResponseAdjustmentsPlanPhasePercentageDiscountAdjustmentFiltersOperatorExcludes:
		return true
	}
	return false
}

type PlanVersionGetResponseAdjustmentsPlanPhaseMinimumAdjustment struct {
	ID             string                                                                    `json:"id,required"`
	AdjustmentType PlanVersionGetResponseAdjustmentsPlanPhaseMinimumAdjustmentAdjustmentType `json:"adjustment_type,required"`
	// The price IDs that this adjustment applies to.
	//
	// Deprecated: deprecated
	AppliesToPriceIDs []string `json:"applies_to_price_ids,required"`
	// The filters that determine which prices to apply this adjustment to.
	Filters []PlanVersionGetResponseAdjustmentsPlanPhaseMinimumAdjustmentFilter `json:"filters,required"`
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
	Reason string                                                          `json:"reason,required,nullable"`
	JSON   planVersionGetResponseAdjustmentsPlanPhaseMinimumAdjustmentJSON `json:"-"`
}

// planVersionGetResponseAdjustmentsPlanPhaseMinimumAdjustmentJSON contains the
// JSON metadata for the struct
// [PlanVersionGetResponseAdjustmentsPlanPhaseMinimumAdjustment]
type planVersionGetResponseAdjustmentsPlanPhaseMinimumAdjustmentJSON struct {
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

func (r *PlanVersionGetResponseAdjustmentsPlanPhaseMinimumAdjustment) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r planVersionGetResponseAdjustmentsPlanPhaseMinimumAdjustmentJSON) RawJSON() string {
	return r.raw
}

func (r PlanVersionGetResponseAdjustmentsPlanPhaseMinimumAdjustment) implementsPlanVersionGetResponseAdjustment() {
}

type PlanVersionGetResponseAdjustmentsPlanPhaseMinimumAdjustmentAdjustmentType string

const (
	PlanVersionGetResponseAdjustmentsPlanPhaseMinimumAdjustmentAdjustmentTypeMinimum PlanVersionGetResponseAdjustmentsPlanPhaseMinimumAdjustmentAdjustmentType = "minimum"
)

func (r PlanVersionGetResponseAdjustmentsPlanPhaseMinimumAdjustmentAdjustmentType) IsKnown() bool {
	switch r {
	case PlanVersionGetResponseAdjustmentsPlanPhaseMinimumAdjustmentAdjustmentTypeMinimum:
		return true
	}
	return false
}

type PlanVersionGetResponseAdjustmentsPlanPhaseMinimumAdjustmentFilter struct {
	// The property of the price to filter on.
	Field PlanVersionGetResponseAdjustmentsPlanPhaseMinimumAdjustmentFiltersField `json:"field,required"`
	// Should prices that match the filter be included or excluded.
	Operator PlanVersionGetResponseAdjustmentsPlanPhaseMinimumAdjustmentFiltersOperator `json:"operator,required"`
	// The IDs or values that match this filter.
	Values []string                                                              `json:"values,required"`
	JSON   planVersionGetResponseAdjustmentsPlanPhaseMinimumAdjustmentFilterJSON `json:"-"`
}

// planVersionGetResponseAdjustmentsPlanPhaseMinimumAdjustmentFilterJSON contains
// the JSON metadata for the struct
// [PlanVersionGetResponseAdjustmentsPlanPhaseMinimumAdjustmentFilter]
type planVersionGetResponseAdjustmentsPlanPhaseMinimumAdjustmentFilterJSON struct {
	Field       apijson.Field
	Operator    apijson.Field
	Values      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *PlanVersionGetResponseAdjustmentsPlanPhaseMinimumAdjustmentFilter) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r planVersionGetResponseAdjustmentsPlanPhaseMinimumAdjustmentFilterJSON) RawJSON() string {
	return r.raw
}

// The property of the price to filter on.
type PlanVersionGetResponseAdjustmentsPlanPhaseMinimumAdjustmentFiltersField string

const (
	PlanVersionGetResponseAdjustmentsPlanPhaseMinimumAdjustmentFiltersFieldPriceID       PlanVersionGetResponseAdjustmentsPlanPhaseMinimumAdjustmentFiltersField = "price_id"
	PlanVersionGetResponseAdjustmentsPlanPhaseMinimumAdjustmentFiltersFieldItemID        PlanVersionGetResponseAdjustmentsPlanPhaseMinimumAdjustmentFiltersField = "item_id"
	PlanVersionGetResponseAdjustmentsPlanPhaseMinimumAdjustmentFiltersFieldPriceType     PlanVersionGetResponseAdjustmentsPlanPhaseMinimumAdjustmentFiltersField = "price_type"
	PlanVersionGetResponseAdjustmentsPlanPhaseMinimumAdjustmentFiltersFieldCurrency      PlanVersionGetResponseAdjustmentsPlanPhaseMinimumAdjustmentFiltersField = "currency"
	PlanVersionGetResponseAdjustmentsPlanPhaseMinimumAdjustmentFiltersFieldPricingUnitID PlanVersionGetResponseAdjustmentsPlanPhaseMinimumAdjustmentFiltersField = "pricing_unit_id"
)

func (r PlanVersionGetResponseAdjustmentsPlanPhaseMinimumAdjustmentFiltersField) IsKnown() bool {
	switch r {
	case PlanVersionGetResponseAdjustmentsPlanPhaseMinimumAdjustmentFiltersFieldPriceID, PlanVersionGetResponseAdjustmentsPlanPhaseMinimumAdjustmentFiltersFieldItemID, PlanVersionGetResponseAdjustmentsPlanPhaseMinimumAdjustmentFiltersFieldPriceType, PlanVersionGetResponseAdjustmentsPlanPhaseMinimumAdjustmentFiltersFieldCurrency, PlanVersionGetResponseAdjustmentsPlanPhaseMinimumAdjustmentFiltersFieldPricingUnitID:
		return true
	}
	return false
}

// Should prices that match the filter be included or excluded.
type PlanVersionGetResponseAdjustmentsPlanPhaseMinimumAdjustmentFiltersOperator string

const (
	PlanVersionGetResponseAdjustmentsPlanPhaseMinimumAdjustmentFiltersOperatorIncludes PlanVersionGetResponseAdjustmentsPlanPhaseMinimumAdjustmentFiltersOperator = "includes"
	PlanVersionGetResponseAdjustmentsPlanPhaseMinimumAdjustmentFiltersOperatorExcludes PlanVersionGetResponseAdjustmentsPlanPhaseMinimumAdjustmentFiltersOperator = "excludes"
)

func (r PlanVersionGetResponseAdjustmentsPlanPhaseMinimumAdjustmentFiltersOperator) IsKnown() bool {
	switch r {
	case PlanVersionGetResponseAdjustmentsPlanPhaseMinimumAdjustmentFiltersOperatorIncludes, PlanVersionGetResponseAdjustmentsPlanPhaseMinimumAdjustmentFiltersOperatorExcludes:
		return true
	}
	return false
}

type PlanVersionGetResponseAdjustmentsPlanPhaseMaximumAdjustment struct {
	ID             string                                                                    `json:"id,required"`
	AdjustmentType PlanVersionGetResponseAdjustmentsPlanPhaseMaximumAdjustmentAdjustmentType `json:"adjustment_type,required"`
	// The price IDs that this adjustment applies to.
	//
	// Deprecated: deprecated
	AppliesToPriceIDs []string `json:"applies_to_price_ids,required"`
	// The filters that determine which prices to apply this adjustment to.
	Filters []PlanVersionGetResponseAdjustmentsPlanPhaseMaximumAdjustmentFilter `json:"filters,required"`
	// True for adjustments that apply to an entire invocice, false for adjustments
	// that apply to only one price.
	IsInvoiceLevel bool `json:"is_invoice_level,required"`
	// The maximum amount to charge in a given billing period for the prices this
	// adjustment applies to.
	MaximumAmount string `json:"maximum_amount,required"`
	// The plan phase in which this adjustment is active.
	PlanPhaseOrder int64 `json:"plan_phase_order,required,nullable"`
	// The reason for the adjustment.
	Reason string                                                          `json:"reason,required,nullable"`
	JSON   planVersionGetResponseAdjustmentsPlanPhaseMaximumAdjustmentJSON `json:"-"`
}

// planVersionGetResponseAdjustmentsPlanPhaseMaximumAdjustmentJSON contains the
// JSON metadata for the struct
// [PlanVersionGetResponseAdjustmentsPlanPhaseMaximumAdjustment]
type planVersionGetResponseAdjustmentsPlanPhaseMaximumAdjustmentJSON struct {
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

func (r *PlanVersionGetResponseAdjustmentsPlanPhaseMaximumAdjustment) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r planVersionGetResponseAdjustmentsPlanPhaseMaximumAdjustmentJSON) RawJSON() string {
	return r.raw
}

func (r PlanVersionGetResponseAdjustmentsPlanPhaseMaximumAdjustment) implementsPlanVersionGetResponseAdjustment() {
}

type PlanVersionGetResponseAdjustmentsPlanPhaseMaximumAdjustmentAdjustmentType string

const (
	PlanVersionGetResponseAdjustmentsPlanPhaseMaximumAdjustmentAdjustmentTypeMaximum PlanVersionGetResponseAdjustmentsPlanPhaseMaximumAdjustmentAdjustmentType = "maximum"
)

func (r PlanVersionGetResponseAdjustmentsPlanPhaseMaximumAdjustmentAdjustmentType) IsKnown() bool {
	switch r {
	case PlanVersionGetResponseAdjustmentsPlanPhaseMaximumAdjustmentAdjustmentTypeMaximum:
		return true
	}
	return false
}

type PlanVersionGetResponseAdjustmentsPlanPhaseMaximumAdjustmentFilter struct {
	// The property of the price to filter on.
	Field PlanVersionGetResponseAdjustmentsPlanPhaseMaximumAdjustmentFiltersField `json:"field,required"`
	// Should prices that match the filter be included or excluded.
	Operator PlanVersionGetResponseAdjustmentsPlanPhaseMaximumAdjustmentFiltersOperator `json:"operator,required"`
	// The IDs or values that match this filter.
	Values []string                                                              `json:"values,required"`
	JSON   planVersionGetResponseAdjustmentsPlanPhaseMaximumAdjustmentFilterJSON `json:"-"`
}

// planVersionGetResponseAdjustmentsPlanPhaseMaximumAdjustmentFilterJSON contains
// the JSON metadata for the struct
// [PlanVersionGetResponseAdjustmentsPlanPhaseMaximumAdjustmentFilter]
type planVersionGetResponseAdjustmentsPlanPhaseMaximumAdjustmentFilterJSON struct {
	Field       apijson.Field
	Operator    apijson.Field
	Values      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *PlanVersionGetResponseAdjustmentsPlanPhaseMaximumAdjustmentFilter) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r planVersionGetResponseAdjustmentsPlanPhaseMaximumAdjustmentFilterJSON) RawJSON() string {
	return r.raw
}

// The property of the price to filter on.
type PlanVersionGetResponseAdjustmentsPlanPhaseMaximumAdjustmentFiltersField string

const (
	PlanVersionGetResponseAdjustmentsPlanPhaseMaximumAdjustmentFiltersFieldPriceID       PlanVersionGetResponseAdjustmentsPlanPhaseMaximumAdjustmentFiltersField = "price_id"
	PlanVersionGetResponseAdjustmentsPlanPhaseMaximumAdjustmentFiltersFieldItemID        PlanVersionGetResponseAdjustmentsPlanPhaseMaximumAdjustmentFiltersField = "item_id"
	PlanVersionGetResponseAdjustmentsPlanPhaseMaximumAdjustmentFiltersFieldPriceType     PlanVersionGetResponseAdjustmentsPlanPhaseMaximumAdjustmentFiltersField = "price_type"
	PlanVersionGetResponseAdjustmentsPlanPhaseMaximumAdjustmentFiltersFieldCurrency      PlanVersionGetResponseAdjustmentsPlanPhaseMaximumAdjustmentFiltersField = "currency"
	PlanVersionGetResponseAdjustmentsPlanPhaseMaximumAdjustmentFiltersFieldPricingUnitID PlanVersionGetResponseAdjustmentsPlanPhaseMaximumAdjustmentFiltersField = "pricing_unit_id"
)

func (r PlanVersionGetResponseAdjustmentsPlanPhaseMaximumAdjustmentFiltersField) IsKnown() bool {
	switch r {
	case PlanVersionGetResponseAdjustmentsPlanPhaseMaximumAdjustmentFiltersFieldPriceID, PlanVersionGetResponseAdjustmentsPlanPhaseMaximumAdjustmentFiltersFieldItemID, PlanVersionGetResponseAdjustmentsPlanPhaseMaximumAdjustmentFiltersFieldPriceType, PlanVersionGetResponseAdjustmentsPlanPhaseMaximumAdjustmentFiltersFieldCurrency, PlanVersionGetResponseAdjustmentsPlanPhaseMaximumAdjustmentFiltersFieldPricingUnitID:
		return true
	}
	return false
}

// Should prices that match the filter be included or excluded.
type PlanVersionGetResponseAdjustmentsPlanPhaseMaximumAdjustmentFiltersOperator string

const (
	PlanVersionGetResponseAdjustmentsPlanPhaseMaximumAdjustmentFiltersOperatorIncludes PlanVersionGetResponseAdjustmentsPlanPhaseMaximumAdjustmentFiltersOperator = "includes"
	PlanVersionGetResponseAdjustmentsPlanPhaseMaximumAdjustmentFiltersOperatorExcludes PlanVersionGetResponseAdjustmentsPlanPhaseMaximumAdjustmentFiltersOperator = "excludes"
)

func (r PlanVersionGetResponseAdjustmentsPlanPhaseMaximumAdjustmentFiltersOperator) IsKnown() bool {
	switch r {
	case PlanVersionGetResponseAdjustmentsPlanPhaseMaximumAdjustmentFiltersOperatorIncludes, PlanVersionGetResponseAdjustmentsPlanPhaseMaximumAdjustmentFiltersOperatorExcludes:
		return true
	}
	return false
}

type PlanVersionGetResponseAdjustmentsAdjustmentType string

const (
	PlanVersionGetResponseAdjustmentsAdjustmentTypeUsageDiscount      PlanVersionGetResponseAdjustmentsAdjustmentType = "usage_discount"
	PlanVersionGetResponseAdjustmentsAdjustmentTypeAmountDiscount     PlanVersionGetResponseAdjustmentsAdjustmentType = "amount_discount"
	PlanVersionGetResponseAdjustmentsAdjustmentTypePercentageDiscount PlanVersionGetResponseAdjustmentsAdjustmentType = "percentage_discount"
	PlanVersionGetResponseAdjustmentsAdjustmentTypeMinimum            PlanVersionGetResponseAdjustmentsAdjustmentType = "minimum"
	PlanVersionGetResponseAdjustmentsAdjustmentTypeMaximum            PlanVersionGetResponseAdjustmentsAdjustmentType = "maximum"
)

func (r PlanVersionGetResponseAdjustmentsAdjustmentType) IsKnown() bool {
	switch r {
	case PlanVersionGetResponseAdjustmentsAdjustmentTypeUsageDiscount, PlanVersionGetResponseAdjustmentsAdjustmentTypeAmountDiscount, PlanVersionGetResponseAdjustmentsAdjustmentTypePercentageDiscount, PlanVersionGetResponseAdjustmentsAdjustmentTypeMinimum, PlanVersionGetResponseAdjustmentsAdjustmentTypeMaximum:
		return true
	}
	return false
}

type PlanVersionGetResponsePlanPhase struct {
	ID          string `json:"id,required"`
	Description string `json:"description,required,nullable"`
	// How many terms of length `duration_unit` this phase is active for. If null, this
	// phase is evergreen and active indefinitely
	Duration     int64                                        `json:"duration,required,nullable"`
	DurationUnit PlanVersionGetResponsePlanPhasesDurationUnit `json:"duration_unit,required,nullable"`
	Name         string                                       `json:"name,required"`
	// Determines the ordering of the phase in a plan's lifecycle. 1 = first phase.
	Order int64                               `json:"order,required"`
	JSON  planVersionGetResponsePlanPhaseJSON `json:"-"`
}

// planVersionGetResponsePlanPhaseJSON contains the JSON metadata for the struct
// [PlanVersionGetResponsePlanPhase]
type planVersionGetResponsePlanPhaseJSON struct {
	ID           apijson.Field
	Description  apijson.Field
	Duration     apijson.Field
	DurationUnit apijson.Field
	Name         apijson.Field
	Order        apijson.Field
	raw          string
	ExtraFields  map[string]apijson.Field
}

func (r *PlanVersionGetResponsePlanPhase) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r planVersionGetResponsePlanPhaseJSON) RawJSON() string {
	return r.raw
}

type PlanVersionGetResponsePlanPhasesDurationUnit string

const (
	PlanVersionGetResponsePlanPhasesDurationUnitDaily      PlanVersionGetResponsePlanPhasesDurationUnit = "daily"
	PlanVersionGetResponsePlanPhasesDurationUnitMonthly    PlanVersionGetResponsePlanPhasesDurationUnit = "monthly"
	PlanVersionGetResponsePlanPhasesDurationUnitQuarterly  PlanVersionGetResponsePlanPhasesDurationUnit = "quarterly"
	PlanVersionGetResponsePlanPhasesDurationUnitSemiAnnual PlanVersionGetResponsePlanPhasesDurationUnit = "semi_annual"
	PlanVersionGetResponsePlanPhasesDurationUnitAnnual     PlanVersionGetResponsePlanPhasesDurationUnit = "annual"
)

func (r PlanVersionGetResponsePlanPhasesDurationUnit) IsKnown() bool {
	switch r {
	case PlanVersionGetResponsePlanPhasesDurationUnitDaily, PlanVersionGetResponsePlanPhasesDurationUnitMonthly, PlanVersionGetResponsePlanPhasesDurationUnitQuarterly, PlanVersionGetResponsePlanPhasesDurationUnitSemiAnnual, PlanVersionGetResponsePlanPhasesDurationUnitAnnual:
		return true
	}
	return false
}

type PlanVersionNewParams struct {
	// New version number.
	Version param.Field[int64] `json:"version,required"`
	// Additional adjustments to be added to the plan.
	AddAdjustments param.Field[[]PlanVersionNewParamsAddAdjustment] `json:"add_adjustments"`
	// Additional prices to be added to the plan.
	AddPrices param.Field[[]PlanVersionNewParamsAddPrice] `json:"add_prices"`
	// Adjustments to be removed from the plan.
	RemoveAdjustments param.Field[[]PlanVersionNewParamsRemoveAdjustment] `json:"remove_adjustments"`
	// Prices to be removed from the plan.
	RemovePrices param.Field[[]PlanVersionNewParamsRemovePrice] `json:"remove_prices"`
	// Adjustments to be replaced with additional adjustments on the plan.
	ReplaceAdjustments param.Field[[]PlanVersionNewParamsReplaceAdjustment] `json:"replace_adjustments"`
	// Prices to be replaced with additional prices on the plan.
	ReplacePrices param.Field[[]PlanVersionNewParamsReplacePrice] `json:"replace_prices"`
	// Set this new plan version as the default
	SetAsDefault param.Field[bool] `json:"set_as_default"`
}

func (r PlanVersionNewParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type PlanVersionNewParamsAddAdjustment struct {
	// The definition of a new adjustment to create and add to the plan.
	Adjustment param.Field[PlanVersionNewParamsAddAdjustmentsAdjustmentUnion] `json:"adjustment,required"`
	// The phase to add this adjustment to.
	PlanPhaseOrder param.Field[int64] `json:"plan_phase_order"`
}

func (r PlanVersionNewParamsAddAdjustment) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// The definition of a new adjustment to create and add to the plan.
type PlanVersionNewParamsAddAdjustmentsAdjustment struct {
	AdjustmentType    param.Field[PlanVersionNewParamsAddAdjustmentsAdjustmentAdjustmentType] `json:"adjustment_type,required"`
	AmountDiscount    param.Field[string]                                                     `json:"amount_discount"`
	AppliesToPriceIDs param.Field[interface{}]                                                `json:"applies_to_price_ids"`
	// When false, this adjustment will be applied to a single price. Otherwise, it
	// will be applied at the invoice level, possibly to multiple prices.
	IsInvoiceLevel param.Field[bool] `json:"is_invoice_level"`
	// The item ID that revenue from this minimum will be attributed to.
	ItemID             param.Field[string]  `json:"item_id"`
	MaximumAmount      param.Field[string]  `json:"maximum_amount"`
	MinimumAmount      param.Field[string]  `json:"minimum_amount"`
	PercentageDiscount param.Field[float64] `json:"percentage_discount"`
	UsageDiscount      param.Field[float64] `json:"usage_discount"`
}

func (r PlanVersionNewParamsAddAdjustmentsAdjustment) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r PlanVersionNewParamsAddAdjustmentsAdjustment) implementsPlanVersionNewParamsAddAdjustmentsAdjustmentUnion() {
}

// The definition of a new adjustment to create and add to the plan.
//
// Satisfied by
// [PlanVersionNewParamsAddAdjustmentsAdjustmentNewPercentageDiscount],
// [PlanVersionNewParamsAddAdjustmentsAdjustmentNewUsageDiscount],
// [PlanVersionNewParamsAddAdjustmentsAdjustmentNewAmountDiscount],
// [PlanVersionNewParamsAddAdjustmentsAdjustmentNewMinimum],
// [PlanVersionNewParamsAddAdjustmentsAdjustmentNewMaximum],
// [PlanVersionNewParamsAddAdjustmentsAdjustment].
type PlanVersionNewParamsAddAdjustmentsAdjustmentUnion interface {
	implementsPlanVersionNewParamsAddAdjustmentsAdjustmentUnion()
}

type PlanVersionNewParamsAddAdjustmentsAdjustmentNewPercentageDiscount struct {
	AdjustmentType     param.Field[PlanVersionNewParamsAddAdjustmentsAdjustmentNewPercentageDiscountAdjustmentType] `json:"adjustment_type,required"`
	PercentageDiscount param.Field[float64]                                                                         `json:"percentage_discount,required"`
	// The set of price IDs to which this adjustment applies.
	AppliesToPriceIDs param.Field[[]string] `json:"applies_to_price_ids"`
	// When false, this adjustment will be applied to a single price. Otherwise, it
	// will be applied at the invoice level, possibly to multiple prices.
	IsInvoiceLevel param.Field[bool] `json:"is_invoice_level"`
}

func (r PlanVersionNewParamsAddAdjustmentsAdjustmentNewPercentageDiscount) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r PlanVersionNewParamsAddAdjustmentsAdjustmentNewPercentageDiscount) implementsPlanVersionNewParamsAddAdjustmentsAdjustmentUnion() {
}

type PlanVersionNewParamsAddAdjustmentsAdjustmentNewPercentageDiscountAdjustmentType string

const (
	PlanVersionNewParamsAddAdjustmentsAdjustmentNewPercentageDiscountAdjustmentTypePercentageDiscount PlanVersionNewParamsAddAdjustmentsAdjustmentNewPercentageDiscountAdjustmentType = "percentage_discount"
)

func (r PlanVersionNewParamsAddAdjustmentsAdjustmentNewPercentageDiscountAdjustmentType) IsKnown() bool {
	switch r {
	case PlanVersionNewParamsAddAdjustmentsAdjustmentNewPercentageDiscountAdjustmentTypePercentageDiscount:
		return true
	}
	return false
}

type PlanVersionNewParamsAddAdjustmentsAdjustmentNewUsageDiscount struct {
	AdjustmentType param.Field[PlanVersionNewParamsAddAdjustmentsAdjustmentNewUsageDiscountAdjustmentType] `json:"adjustment_type,required"`
	UsageDiscount  param.Field[float64]                                                                    `json:"usage_discount,required"`
	// The set of price IDs to which this adjustment applies.
	AppliesToPriceIDs param.Field[[]string] `json:"applies_to_price_ids"`
	// When false, this adjustment will be applied to a single price. Otherwise, it
	// will be applied at the invoice level, possibly to multiple prices.
	IsInvoiceLevel param.Field[bool] `json:"is_invoice_level"`
}

func (r PlanVersionNewParamsAddAdjustmentsAdjustmentNewUsageDiscount) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r PlanVersionNewParamsAddAdjustmentsAdjustmentNewUsageDiscount) implementsPlanVersionNewParamsAddAdjustmentsAdjustmentUnion() {
}

type PlanVersionNewParamsAddAdjustmentsAdjustmentNewUsageDiscountAdjustmentType string

const (
	PlanVersionNewParamsAddAdjustmentsAdjustmentNewUsageDiscountAdjustmentTypeUsageDiscount PlanVersionNewParamsAddAdjustmentsAdjustmentNewUsageDiscountAdjustmentType = "usage_discount"
)

func (r PlanVersionNewParamsAddAdjustmentsAdjustmentNewUsageDiscountAdjustmentType) IsKnown() bool {
	switch r {
	case PlanVersionNewParamsAddAdjustmentsAdjustmentNewUsageDiscountAdjustmentTypeUsageDiscount:
		return true
	}
	return false
}

type PlanVersionNewParamsAddAdjustmentsAdjustmentNewAmountDiscount struct {
	AdjustmentType param.Field[PlanVersionNewParamsAddAdjustmentsAdjustmentNewAmountDiscountAdjustmentType] `json:"adjustment_type,required"`
	AmountDiscount param.Field[string]                                                                      `json:"amount_discount,required"`
	// The set of price IDs to which this adjustment applies.
	AppliesToPriceIDs param.Field[[]string] `json:"applies_to_price_ids"`
	// When false, this adjustment will be applied to a single price. Otherwise, it
	// will be applied at the invoice level, possibly to multiple prices.
	IsInvoiceLevel param.Field[bool] `json:"is_invoice_level"`
}

func (r PlanVersionNewParamsAddAdjustmentsAdjustmentNewAmountDiscount) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r PlanVersionNewParamsAddAdjustmentsAdjustmentNewAmountDiscount) implementsPlanVersionNewParamsAddAdjustmentsAdjustmentUnion() {
}

type PlanVersionNewParamsAddAdjustmentsAdjustmentNewAmountDiscountAdjustmentType string

const (
	PlanVersionNewParamsAddAdjustmentsAdjustmentNewAmountDiscountAdjustmentTypeAmountDiscount PlanVersionNewParamsAddAdjustmentsAdjustmentNewAmountDiscountAdjustmentType = "amount_discount"
)

func (r PlanVersionNewParamsAddAdjustmentsAdjustmentNewAmountDiscountAdjustmentType) IsKnown() bool {
	switch r {
	case PlanVersionNewParamsAddAdjustmentsAdjustmentNewAmountDiscountAdjustmentTypeAmountDiscount:
		return true
	}
	return false
}

type PlanVersionNewParamsAddAdjustmentsAdjustmentNewMinimum struct {
	AdjustmentType param.Field[PlanVersionNewParamsAddAdjustmentsAdjustmentNewMinimumAdjustmentType] `json:"adjustment_type,required"`
	// The item ID that revenue from this minimum will be attributed to.
	ItemID        param.Field[string] `json:"item_id,required"`
	MinimumAmount param.Field[string] `json:"minimum_amount,required"`
	// The set of price IDs to which this adjustment applies.
	AppliesToPriceIDs param.Field[[]string] `json:"applies_to_price_ids"`
	// When false, this adjustment will be applied to a single price. Otherwise, it
	// will be applied at the invoice level, possibly to multiple prices.
	IsInvoiceLevel param.Field[bool] `json:"is_invoice_level"`
}

func (r PlanVersionNewParamsAddAdjustmentsAdjustmentNewMinimum) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r PlanVersionNewParamsAddAdjustmentsAdjustmentNewMinimum) implementsPlanVersionNewParamsAddAdjustmentsAdjustmentUnion() {
}

type PlanVersionNewParamsAddAdjustmentsAdjustmentNewMinimumAdjustmentType string

const (
	PlanVersionNewParamsAddAdjustmentsAdjustmentNewMinimumAdjustmentTypeMinimum PlanVersionNewParamsAddAdjustmentsAdjustmentNewMinimumAdjustmentType = "minimum"
)

func (r PlanVersionNewParamsAddAdjustmentsAdjustmentNewMinimumAdjustmentType) IsKnown() bool {
	switch r {
	case PlanVersionNewParamsAddAdjustmentsAdjustmentNewMinimumAdjustmentTypeMinimum:
		return true
	}
	return false
}

type PlanVersionNewParamsAddAdjustmentsAdjustmentNewMaximum struct {
	AdjustmentType param.Field[PlanVersionNewParamsAddAdjustmentsAdjustmentNewMaximumAdjustmentType] `json:"adjustment_type,required"`
	MaximumAmount  param.Field[string]                                                               `json:"maximum_amount,required"`
	// The set of price IDs to which this adjustment applies.
	AppliesToPriceIDs param.Field[[]string] `json:"applies_to_price_ids"`
	// When false, this adjustment will be applied to a single price. Otherwise, it
	// will be applied at the invoice level, possibly to multiple prices.
	IsInvoiceLevel param.Field[bool] `json:"is_invoice_level"`
}

func (r PlanVersionNewParamsAddAdjustmentsAdjustmentNewMaximum) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r PlanVersionNewParamsAddAdjustmentsAdjustmentNewMaximum) implementsPlanVersionNewParamsAddAdjustmentsAdjustmentUnion() {
}

type PlanVersionNewParamsAddAdjustmentsAdjustmentNewMaximumAdjustmentType string

const (
	PlanVersionNewParamsAddAdjustmentsAdjustmentNewMaximumAdjustmentTypeMaximum PlanVersionNewParamsAddAdjustmentsAdjustmentNewMaximumAdjustmentType = "maximum"
)

func (r PlanVersionNewParamsAddAdjustmentsAdjustmentNewMaximumAdjustmentType) IsKnown() bool {
	switch r {
	case PlanVersionNewParamsAddAdjustmentsAdjustmentNewMaximumAdjustmentTypeMaximum:
		return true
	}
	return false
}

type PlanVersionNewParamsAddAdjustmentsAdjustmentAdjustmentType string

const (
	PlanVersionNewParamsAddAdjustmentsAdjustmentAdjustmentTypePercentageDiscount PlanVersionNewParamsAddAdjustmentsAdjustmentAdjustmentType = "percentage_discount"
	PlanVersionNewParamsAddAdjustmentsAdjustmentAdjustmentTypeUsageDiscount      PlanVersionNewParamsAddAdjustmentsAdjustmentAdjustmentType = "usage_discount"
	PlanVersionNewParamsAddAdjustmentsAdjustmentAdjustmentTypeAmountDiscount     PlanVersionNewParamsAddAdjustmentsAdjustmentAdjustmentType = "amount_discount"
	PlanVersionNewParamsAddAdjustmentsAdjustmentAdjustmentTypeMinimum            PlanVersionNewParamsAddAdjustmentsAdjustmentAdjustmentType = "minimum"
	PlanVersionNewParamsAddAdjustmentsAdjustmentAdjustmentTypeMaximum            PlanVersionNewParamsAddAdjustmentsAdjustmentAdjustmentType = "maximum"
)

func (r PlanVersionNewParamsAddAdjustmentsAdjustmentAdjustmentType) IsKnown() bool {
	switch r {
	case PlanVersionNewParamsAddAdjustmentsAdjustmentAdjustmentTypePercentageDiscount, PlanVersionNewParamsAddAdjustmentsAdjustmentAdjustmentTypeUsageDiscount, PlanVersionNewParamsAddAdjustmentsAdjustmentAdjustmentTypeAmountDiscount, PlanVersionNewParamsAddAdjustmentsAdjustmentAdjustmentTypeMinimum, PlanVersionNewParamsAddAdjustmentsAdjustmentAdjustmentTypeMaximum:
		return true
	}
	return false
}

type PlanVersionNewParamsAddPrice struct {
	// The allocation price to add to the plan.
	AllocationPrice param.Field[PlanVersionNewParamsAddPricesAllocationPrice] `json:"allocation_price"`
	// The phase to add this price to.
	PlanPhaseOrder param.Field[int64] `json:"plan_phase_order"`
	// The price to add to the plan
	Price param.Field[PlanVersionNewParamsAddPricesPriceUnion] `json:"price"`
}

func (r PlanVersionNewParamsAddPrice) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// The allocation price to add to the plan.
type PlanVersionNewParamsAddPricesAllocationPrice struct {
	// An amount of the currency to allocate to the customer at the specified cadence.
	Amount param.Field[string] `json:"amount,required"`
	// The cadence at which to allocate the amount to the customer.
	Cadence param.Field[PlanVersionNewParamsAddPricesAllocationPriceCadence] `json:"cadence,required"`
	// An ISO 4217 currency string or a custom pricing unit identifier in which to bill
	// this price.
	Currency param.Field[string] `json:"currency,required"`
	// Whether the allocated amount should expire at the end of the cadence or roll
	// over to the next period.
	ExpiresAtEndOfCadence param.Field[bool] `json:"expires_at_end_of_cadence,required"`
}

func (r PlanVersionNewParamsAddPricesAllocationPrice) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// The cadence at which to allocate the amount to the customer.
type PlanVersionNewParamsAddPricesAllocationPriceCadence string

const (
	PlanVersionNewParamsAddPricesAllocationPriceCadenceOneTime    PlanVersionNewParamsAddPricesAllocationPriceCadence = "one_time"
	PlanVersionNewParamsAddPricesAllocationPriceCadenceMonthly    PlanVersionNewParamsAddPricesAllocationPriceCadence = "monthly"
	PlanVersionNewParamsAddPricesAllocationPriceCadenceQuarterly  PlanVersionNewParamsAddPricesAllocationPriceCadence = "quarterly"
	PlanVersionNewParamsAddPricesAllocationPriceCadenceSemiAnnual PlanVersionNewParamsAddPricesAllocationPriceCadence = "semi_annual"
	PlanVersionNewParamsAddPricesAllocationPriceCadenceAnnual     PlanVersionNewParamsAddPricesAllocationPriceCadence = "annual"
	PlanVersionNewParamsAddPricesAllocationPriceCadenceCustom     PlanVersionNewParamsAddPricesAllocationPriceCadence = "custom"
)

func (r PlanVersionNewParamsAddPricesAllocationPriceCadence) IsKnown() bool {
	switch r {
	case PlanVersionNewParamsAddPricesAllocationPriceCadenceOneTime, PlanVersionNewParamsAddPricesAllocationPriceCadenceMonthly, PlanVersionNewParamsAddPricesAllocationPriceCadenceQuarterly, PlanVersionNewParamsAddPricesAllocationPriceCadenceSemiAnnual, PlanVersionNewParamsAddPricesAllocationPriceCadenceAnnual, PlanVersionNewParamsAddPricesAllocationPriceCadenceCustom:
		return true
	}
	return false
}

// The price to add to the plan
type PlanVersionNewParamsAddPricesPrice struct {
	// The cadence to bill for this price on.
	Cadence param.Field[PlanVersionNewParamsAddPricesPriceCadence] `json:"cadence,required"`
	// The id of the item the price will be associated with.
	ItemID    param.Field[string]                                      `json:"item_id,required"`
	ModelType param.Field[PlanVersionNewParamsAddPricesPriceModelType] `json:"model_type,required"`
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
	ConversionRate              param.Field[float64]     `json:"conversion_rate"`
	CumulativeGroupedBulkConfig param.Field[interface{}] `json:"cumulative_grouped_bulk_config"`
	// An ISO 4217 currency string, or custom pricing unit identifier, in which this
	// price is billed.
	Currency                      param.Field[string]      `json:"currency"`
	DimensionalPriceConfiguration param.Field[interface{}] `json:"dimensional_price_configuration"`
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
	InvoiceGroupingKey                    param.Field[string]      `json:"invoice_grouping_key"`
	InvoicingCycleConfiguration           param.Field[interface{}] `json:"invoicing_cycle_configuration"`
	MatrixConfig                          param.Field[interface{}] `json:"matrix_config"`
	MatrixWithAllocationConfig            param.Field[interface{}] `json:"matrix_with_allocation_config"`
	MatrixWithDisplayNameConfig           param.Field[interface{}] `json:"matrix_with_display_name_config"`
	MaxGroupTieredPackageConfig           param.Field[interface{}] `json:"max_group_tiered_package_config"`
	Metadata                              param.Field[interface{}] `json:"metadata"`
	PackageConfig                         param.Field[interface{}] `json:"package_config"`
	PackageWithAllocationConfig           param.Field[interface{}] `json:"package_with_allocation_config"`
	ScalableMatrixWithTieredPricingConfig param.Field[interface{}] `json:"scalable_matrix_with_tiered_pricing_config"`
	ScalableMatrixWithUnitPricingConfig   param.Field[interface{}] `json:"scalable_matrix_with_unit_pricing_config"`
	ThresholdTotalAmountConfig            param.Field[interface{}] `json:"threshold_total_amount_config"`
	TieredBpsConfig                       param.Field[interface{}] `json:"tiered_bps_config"`
	TieredConfig                          param.Field[interface{}] `json:"tiered_config"`
	TieredPackageConfig                   param.Field[interface{}] `json:"tiered_package_config"`
	TieredPackageWithMinimumConfig        param.Field[interface{}] `json:"tiered_package_with_minimum_config"`
	TieredWithMinimumConfig               param.Field[interface{}] `json:"tiered_with_minimum_config"`
	TieredWithProrationConfig             param.Field[interface{}] `json:"tiered_with_proration_config"`
	UnitConfig                            param.Field[interface{}] `json:"unit_config"`
	UnitWithPercentConfig                 param.Field[interface{}] `json:"unit_with_percent_config"`
	UnitWithProrationConfig               param.Field[interface{}] `json:"unit_with_proration_config"`
}

func (r PlanVersionNewParamsAddPricesPrice) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r PlanVersionNewParamsAddPricesPrice) implementsPlanVersionNewParamsAddPricesPriceUnion() {}

// The price to add to the plan
//
// Satisfied by [PlanVersionNewParamsAddPricesPriceNewPlanUnitPrice],
// [PlanVersionNewParamsAddPricesPriceNewPlanPackagePrice],
// [PlanVersionNewParamsAddPricesPriceNewPlanMatrixPrice],
// [PlanVersionNewParamsAddPricesPriceNewPlanTieredPrice],
// [PlanVersionNewParamsAddPricesPriceNewPlanTieredBpsPrice],
// [PlanVersionNewParamsAddPricesPriceNewPlanBpsPrice],
// [PlanVersionNewParamsAddPricesPriceNewPlanBulkBpsPrice],
// [PlanVersionNewParamsAddPricesPriceNewPlanBulkPrice],
// [PlanVersionNewParamsAddPricesPriceNewPlanThresholdTotalAmountPrice],
// [PlanVersionNewParamsAddPricesPriceNewPlanTieredPackagePrice],
// [PlanVersionNewParamsAddPricesPriceNewPlanTieredWithMinimumPrice],
// [PlanVersionNewParamsAddPricesPriceNewPlanUnitWithPercentPrice],
// [PlanVersionNewParamsAddPricesPriceNewPlanPackageWithAllocationPrice],
// [PlanVersionNewParamsAddPricesPriceNewPlanTierWithProrationPrice],
// [PlanVersionNewParamsAddPricesPriceNewPlanUnitWithProrationPrice],
// [PlanVersionNewParamsAddPricesPriceNewPlanGroupedAllocationPrice],
// [PlanVersionNewParamsAddPricesPriceNewPlanGroupedWithProratedMinimumPrice],
// [PlanVersionNewParamsAddPricesPriceNewPlanGroupedWithMeteredMinimumPrice],
// [PlanVersionNewParamsAddPricesPriceNewPlanMatrixWithDisplayNamePrice],
// [PlanVersionNewParamsAddPricesPriceNewPlanBulkWithProrationPrice],
// [PlanVersionNewParamsAddPricesPriceNewPlanGroupedTieredPackagePrice],
// [PlanVersionNewParamsAddPricesPriceNewPlanMaxGroupTieredPackagePrice],
// [PlanVersionNewParamsAddPricesPriceNewPlanScalableMatrixWithUnitPricingPrice],
// [PlanVersionNewParamsAddPricesPriceNewPlanScalableMatrixWithTieredPricingPrice],
// [PlanVersionNewParamsAddPricesPriceNewPlanCumulativeGroupedBulkPrice],
// [PlanVersionNewParamsAddPricesPriceNewPlanTieredPackageWithMinimumPrice],
// [PlanVersionNewParamsAddPricesPriceNewPlanMatrixWithAllocationPrice],
// [PlanVersionNewParamsAddPricesPriceNewPlanGroupedTieredPrice],
// [PlanVersionNewParamsAddPricesPrice].
type PlanVersionNewParamsAddPricesPriceUnion interface {
	implementsPlanVersionNewParamsAddPricesPriceUnion()
}

type PlanVersionNewParamsAddPricesPriceNewPlanUnitPrice struct {
	// The cadence to bill for this price on.
	Cadence param.Field[PlanVersionNewParamsAddPricesPriceNewPlanUnitPriceCadence] `json:"cadence,required"`
	// The id of the item the price will be associated with.
	ItemID    param.Field[string]                                                      `json:"item_id,required"`
	ModelType param.Field[PlanVersionNewParamsAddPricesPriceNewPlanUnitPriceModelType] `json:"model_type,required"`
	// The name of the price.
	Name       param.Field[string]                                                       `json:"name,required"`
	UnitConfig param.Field[PlanVersionNewParamsAddPricesPriceNewPlanUnitPriceUnitConfig] `json:"unit_config,required"`
	// The id of the billable metric for the price. Only needed if the price is
	// usage-based.
	BillableMetricID param.Field[string] `json:"billable_metric_id"`
	// If the Price represents a fixed cost, the price will be billed in-advance if
	// this is true, and in-arrears if this is false.
	BilledInAdvance param.Field[bool] `json:"billed_in_advance"`
	// For custom cadence: specifies the duration of the billing period in days or
	// months.
	BillingCycleConfiguration param.Field[PlanVersionNewParamsAddPricesPriceNewPlanUnitPriceBillingCycleConfiguration] `json:"billing_cycle_configuration"`
	// The per unit conversion rate of the price currency to the invoicing currency.
	ConversionRate param.Field[float64] `json:"conversion_rate"`
	// An ISO 4217 currency string, or custom pricing unit identifier, in which this
	// price is billed.
	Currency param.Field[string] `json:"currency"`
	// For dimensional price: specifies a price group and dimension values
	DimensionalPriceConfiguration param.Field[PlanVersionNewParamsAddPricesPriceNewPlanUnitPriceDimensionalPriceConfiguration] `json:"dimensional_price_configuration"`
	// An alias for the price.
	ExternalPriceID param.Field[string] `json:"external_price_id"`
	// If the Price represents a fixed cost, this represents the quantity of units
	// applied.
	FixedPriceQuantity param.Field[float64] `json:"fixed_price_quantity"`
	// The property used to group this price on an invoice
	InvoiceGroupingKey param.Field[string] `json:"invoice_grouping_key"`
	// Within each billing cycle, specifies the cadence at which invoices are produced.
	// If unspecified, a single invoice is produced per billing cycle.
	InvoicingCycleConfiguration param.Field[PlanVersionNewParamsAddPricesPriceNewPlanUnitPriceInvoicingCycleConfiguration] `json:"invoicing_cycle_configuration"`
	// User-specified key/value pairs for the resource. Individual keys can be removed
	// by setting the value to `null`, and the entire metadata mapping can be cleared
	// by setting `metadata` to `null`.
	Metadata param.Field[map[string]string] `json:"metadata"`
}

func (r PlanVersionNewParamsAddPricesPriceNewPlanUnitPrice) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r PlanVersionNewParamsAddPricesPriceNewPlanUnitPrice) implementsPlanVersionNewParamsAddPricesPriceUnion() {
}

// The cadence to bill for this price on.
type PlanVersionNewParamsAddPricesPriceNewPlanUnitPriceCadence string

const (
	PlanVersionNewParamsAddPricesPriceNewPlanUnitPriceCadenceAnnual     PlanVersionNewParamsAddPricesPriceNewPlanUnitPriceCadence = "annual"
	PlanVersionNewParamsAddPricesPriceNewPlanUnitPriceCadenceSemiAnnual PlanVersionNewParamsAddPricesPriceNewPlanUnitPriceCadence = "semi_annual"
	PlanVersionNewParamsAddPricesPriceNewPlanUnitPriceCadenceMonthly    PlanVersionNewParamsAddPricesPriceNewPlanUnitPriceCadence = "monthly"
	PlanVersionNewParamsAddPricesPriceNewPlanUnitPriceCadenceQuarterly  PlanVersionNewParamsAddPricesPriceNewPlanUnitPriceCadence = "quarterly"
	PlanVersionNewParamsAddPricesPriceNewPlanUnitPriceCadenceOneTime    PlanVersionNewParamsAddPricesPriceNewPlanUnitPriceCadence = "one_time"
	PlanVersionNewParamsAddPricesPriceNewPlanUnitPriceCadenceCustom     PlanVersionNewParamsAddPricesPriceNewPlanUnitPriceCadence = "custom"
)

func (r PlanVersionNewParamsAddPricesPriceNewPlanUnitPriceCadence) IsKnown() bool {
	switch r {
	case PlanVersionNewParamsAddPricesPriceNewPlanUnitPriceCadenceAnnual, PlanVersionNewParamsAddPricesPriceNewPlanUnitPriceCadenceSemiAnnual, PlanVersionNewParamsAddPricesPriceNewPlanUnitPriceCadenceMonthly, PlanVersionNewParamsAddPricesPriceNewPlanUnitPriceCadenceQuarterly, PlanVersionNewParamsAddPricesPriceNewPlanUnitPriceCadenceOneTime, PlanVersionNewParamsAddPricesPriceNewPlanUnitPriceCadenceCustom:
		return true
	}
	return false
}

type PlanVersionNewParamsAddPricesPriceNewPlanUnitPriceModelType string

const (
	PlanVersionNewParamsAddPricesPriceNewPlanUnitPriceModelTypeUnit PlanVersionNewParamsAddPricesPriceNewPlanUnitPriceModelType = "unit"
)

func (r PlanVersionNewParamsAddPricesPriceNewPlanUnitPriceModelType) IsKnown() bool {
	switch r {
	case PlanVersionNewParamsAddPricesPriceNewPlanUnitPriceModelTypeUnit:
		return true
	}
	return false
}

type PlanVersionNewParamsAddPricesPriceNewPlanUnitPriceUnitConfig struct {
	// Rate per unit of usage
	UnitAmount param.Field[string] `json:"unit_amount,required"`
}

func (r PlanVersionNewParamsAddPricesPriceNewPlanUnitPriceUnitConfig) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// For custom cadence: specifies the duration of the billing period in days or
// months.
type PlanVersionNewParamsAddPricesPriceNewPlanUnitPriceBillingCycleConfiguration struct {
	// The duration of the billing period.
	Duration param.Field[int64] `json:"duration,required"`
	// The unit of billing period duration.
	DurationUnit param.Field[PlanVersionNewParamsAddPricesPriceNewPlanUnitPriceBillingCycleConfigurationDurationUnit] `json:"duration_unit,required"`
}

func (r PlanVersionNewParamsAddPricesPriceNewPlanUnitPriceBillingCycleConfiguration) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// The unit of billing period duration.
type PlanVersionNewParamsAddPricesPriceNewPlanUnitPriceBillingCycleConfigurationDurationUnit string

const (
	PlanVersionNewParamsAddPricesPriceNewPlanUnitPriceBillingCycleConfigurationDurationUnitDay   PlanVersionNewParamsAddPricesPriceNewPlanUnitPriceBillingCycleConfigurationDurationUnit = "day"
	PlanVersionNewParamsAddPricesPriceNewPlanUnitPriceBillingCycleConfigurationDurationUnitMonth PlanVersionNewParamsAddPricesPriceNewPlanUnitPriceBillingCycleConfigurationDurationUnit = "month"
)

func (r PlanVersionNewParamsAddPricesPriceNewPlanUnitPriceBillingCycleConfigurationDurationUnit) IsKnown() bool {
	switch r {
	case PlanVersionNewParamsAddPricesPriceNewPlanUnitPriceBillingCycleConfigurationDurationUnitDay, PlanVersionNewParamsAddPricesPriceNewPlanUnitPriceBillingCycleConfigurationDurationUnitMonth:
		return true
	}
	return false
}

// For dimensional price: specifies a price group and dimension values
type PlanVersionNewParamsAddPricesPriceNewPlanUnitPriceDimensionalPriceConfiguration struct {
	// The list of dimension values matching (in order) the dimensions of the price
	// group
	DimensionValues param.Field[[]string] `json:"dimension_values,required"`
	// The id of the dimensional price group to include this price in
	DimensionalPriceGroupID param.Field[string] `json:"dimensional_price_group_id"`
	// The external id of the dimensional price group to include this price in
	ExternalDimensionalPriceGroupID param.Field[string] `json:"external_dimensional_price_group_id"`
}

func (r PlanVersionNewParamsAddPricesPriceNewPlanUnitPriceDimensionalPriceConfiguration) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Within each billing cycle, specifies the cadence at which invoices are produced.
// If unspecified, a single invoice is produced per billing cycle.
type PlanVersionNewParamsAddPricesPriceNewPlanUnitPriceInvoicingCycleConfiguration struct {
	// The duration of the billing period.
	Duration param.Field[int64] `json:"duration,required"`
	// The unit of billing period duration.
	DurationUnit param.Field[PlanVersionNewParamsAddPricesPriceNewPlanUnitPriceInvoicingCycleConfigurationDurationUnit] `json:"duration_unit,required"`
}

func (r PlanVersionNewParamsAddPricesPriceNewPlanUnitPriceInvoicingCycleConfiguration) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// The unit of billing period duration.
type PlanVersionNewParamsAddPricesPriceNewPlanUnitPriceInvoicingCycleConfigurationDurationUnit string

const (
	PlanVersionNewParamsAddPricesPriceNewPlanUnitPriceInvoicingCycleConfigurationDurationUnitDay   PlanVersionNewParamsAddPricesPriceNewPlanUnitPriceInvoicingCycleConfigurationDurationUnit = "day"
	PlanVersionNewParamsAddPricesPriceNewPlanUnitPriceInvoicingCycleConfigurationDurationUnitMonth PlanVersionNewParamsAddPricesPriceNewPlanUnitPriceInvoicingCycleConfigurationDurationUnit = "month"
)

func (r PlanVersionNewParamsAddPricesPriceNewPlanUnitPriceInvoicingCycleConfigurationDurationUnit) IsKnown() bool {
	switch r {
	case PlanVersionNewParamsAddPricesPriceNewPlanUnitPriceInvoicingCycleConfigurationDurationUnitDay, PlanVersionNewParamsAddPricesPriceNewPlanUnitPriceInvoicingCycleConfigurationDurationUnitMonth:
		return true
	}
	return false
}

type PlanVersionNewParamsAddPricesPriceNewPlanPackagePrice struct {
	// The cadence to bill for this price on.
	Cadence param.Field[PlanVersionNewParamsAddPricesPriceNewPlanPackagePriceCadence] `json:"cadence,required"`
	// The id of the item the price will be associated with.
	ItemID    param.Field[string]                                                         `json:"item_id,required"`
	ModelType param.Field[PlanVersionNewParamsAddPricesPriceNewPlanPackagePriceModelType] `json:"model_type,required"`
	// The name of the price.
	Name          param.Field[string]                                                             `json:"name,required"`
	PackageConfig param.Field[PlanVersionNewParamsAddPricesPriceNewPlanPackagePricePackageConfig] `json:"package_config,required"`
	// The id of the billable metric for the price. Only needed if the price is
	// usage-based.
	BillableMetricID param.Field[string] `json:"billable_metric_id"`
	// If the Price represents a fixed cost, the price will be billed in-advance if
	// this is true, and in-arrears if this is false.
	BilledInAdvance param.Field[bool] `json:"billed_in_advance"`
	// For custom cadence: specifies the duration of the billing period in days or
	// months.
	BillingCycleConfiguration param.Field[PlanVersionNewParamsAddPricesPriceNewPlanPackagePriceBillingCycleConfiguration] `json:"billing_cycle_configuration"`
	// The per unit conversion rate of the price currency to the invoicing currency.
	ConversionRate param.Field[float64] `json:"conversion_rate"`
	// An ISO 4217 currency string, or custom pricing unit identifier, in which this
	// price is billed.
	Currency param.Field[string] `json:"currency"`
	// For dimensional price: specifies a price group and dimension values
	DimensionalPriceConfiguration param.Field[PlanVersionNewParamsAddPricesPriceNewPlanPackagePriceDimensionalPriceConfiguration] `json:"dimensional_price_configuration"`
	// An alias for the price.
	ExternalPriceID param.Field[string] `json:"external_price_id"`
	// If the Price represents a fixed cost, this represents the quantity of units
	// applied.
	FixedPriceQuantity param.Field[float64] `json:"fixed_price_quantity"`
	// The property used to group this price on an invoice
	InvoiceGroupingKey param.Field[string] `json:"invoice_grouping_key"`
	// Within each billing cycle, specifies the cadence at which invoices are produced.
	// If unspecified, a single invoice is produced per billing cycle.
	InvoicingCycleConfiguration param.Field[PlanVersionNewParamsAddPricesPriceNewPlanPackagePriceInvoicingCycleConfiguration] `json:"invoicing_cycle_configuration"`
	// User-specified key/value pairs for the resource. Individual keys can be removed
	// by setting the value to `null`, and the entire metadata mapping can be cleared
	// by setting `metadata` to `null`.
	Metadata param.Field[map[string]string] `json:"metadata"`
}

func (r PlanVersionNewParamsAddPricesPriceNewPlanPackagePrice) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r PlanVersionNewParamsAddPricesPriceNewPlanPackagePrice) implementsPlanVersionNewParamsAddPricesPriceUnion() {
}

// The cadence to bill for this price on.
type PlanVersionNewParamsAddPricesPriceNewPlanPackagePriceCadence string

const (
	PlanVersionNewParamsAddPricesPriceNewPlanPackagePriceCadenceAnnual     PlanVersionNewParamsAddPricesPriceNewPlanPackagePriceCadence = "annual"
	PlanVersionNewParamsAddPricesPriceNewPlanPackagePriceCadenceSemiAnnual PlanVersionNewParamsAddPricesPriceNewPlanPackagePriceCadence = "semi_annual"
	PlanVersionNewParamsAddPricesPriceNewPlanPackagePriceCadenceMonthly    PlanVersionNewParamsAddPricesPriceNewPlanPackagePriceCadence = "monthly"
	PlanVersionNewParamsAddPricesPriceNewPlanPackagePriceCadenceQuarterly  PlanVersionNewParamsAddPricesPriceNewPlanPackagePriceCadence = "quarterly"
	PlanVersionNewParamsAddPricesPriceNewPlanPackagePriceCadenceOneTime    PlanVersionNewParamsAddPricesPriceNewPlanPackagePriceCadence = "one_time"
	PlanVersionNewParamsAddPricesPriceNewPlanPackagePriceCadenceCustom     PlanVersionNewParamsAddPricesPriceNewPlanPackagePriceCadence = "custom"
)

func (r PlanVersionNewParamsAddPricesPriceNewPlanPackagePriceCadence) IsKnown() bool {
	switch r {
	case PlanVersionNewParamsAddPricesPriceNewPlanPackagePriceCadenceAnnual, PlanVersionNewParamsAddPricesPriceNewPlanPackagePriceCadenceSemiAnnual, PlanVersionNewParamsAddPricesPriceNewPlanPackagePriceCadenceMonthly, PlanVersionNewParamsAddPricesPriceNewPlanPackagePriceCadenceQuarterly, PlanVersionNewParamsAddPricesPriceNewPlanPackagePriceCadenceOneTime, PlanVersionNewParamsAddPricesPriceNewPlanPackagePriceCadenceCustom:
		return true
	}
	return false
}

type PlanVersionNewParamsAddPricesPriceNewPlanPackagePriceModelType string

const (
	PlanVersionNewParamsAddPricesPriceNewPlanPackagePriceModelTypePackage PlanVersionNewParamsAddPricesPriceNewPlanPackagePriceModelType = "package"
)

func (r PlanVersionNewParamsAddPricesPriceNewPlanPackagePriceModelType) IsKnown() bool {
	switch r {
	case PlanVersionNewParamsAddPricesPriceNewPlanPackagePriceModelTypePackage:
		return true
	}
	return false
}

type PlanVersionNewParamsAddPricesPriceNewPlanPackagePricePackageConfig struct {
	// A currency amount to rate usage by
	PackageAmount param.Field[string] `json:"package_amount,required"`
	// An integer amount to represent package size. For example, 1000 here would divide
	// usage by 1000 before multiplying by package_amount in rating
	PackageSize param.Field[int64] `json:"package_size,required"`
}

func (r PlanVersionNewParamsAddPricesPriceNewPlanPackagePricePackageConfig) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// For custom cadence: specifies the duration of the billing period in days or
// months.
type PlanVersionNewParamsAddPricesPriceNewPlanPackagePriceBillingCycleConfiguration struct {
	// The duration of the billing period.
	Duration param.Field[int64] `json:"duration,required"`
	// The unit of billing period duration.
	DurationUnit param.Field[PlanVersionNewParamsAddPricesPriceNewPlanPackagePriceBillingCycleConfigurationDurationUnit] `json:"duration_unit,required"`
}

func (r PlanVersionNewParamsAddPricesPriceNewPlanPackagePriceBillingCycleConfiguration) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// The unit of billing period duration.
type PlanVersionNewParamsAddPricesPriceNewPlanPackagePriceBillingCycleConfigurationDurationUnit string

const (
	PlanVersionNewParamsAddPricesPriceNewPlanPackagePriceBillingCycleConfigurationDurationUnitDay   PlanVersionNewParamsAddPricesPriceNewPlanPackagePriceBillingCycleConfigurationDurationUnit = "day"
	PlanVersionNewParamsAddPricesPriceNewPlanPackagePriceBillingCycleConfigurationDurationUnitMonth PlanVersionNewParamsAddPricesPriceNewPlanPackagePriceBillingCycleConfigurationDurationUnit = "month"
)

func (r PlanVersionNewParamsAddPricesPriceNewPlanPackagePriceBillingCycleConfigurationDurationUnit) IsKnown() bool {
	switch r {
	case PlanVersionNewParamsAddPricesPriceNewPlanPackagePriceBillingCycleConfigurationDurationUnitDay, PlanVersionNewParamsAddPricesPriceNewPlanPackagePriceBillingCycleConfigurationDurationUnitMonth:
		return true
	}
	return false
}

// For dimensional price: specifies a price group and dimension values
type PlanVersionNewParamsAddPricesPriceNewPlanPackagePriceDimensionalPriceConfiguration struct {
	// The list of dimension values matching (in order) the dimensions of the price
	// group
	DimensionValues param.Field[[]string] `json:"dimension_values,required"`
	// The id of the dimensional price group to include this price in
	DimensionalPriceGroupID param.Field[string] `json:"dimensional_price_group_id"`
	// The external id of the dimensional price group to include this price in
	ExternalDimensionalPriceGroupID param.Field[string] `json:"external_dimensional_price_group_id"`
}

func (r PlanVersionNewParamsAddPricesPriceNewPlanPackagePriceDimensionalPriceConfiguration) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Within each billing cycle, specifies the cadence at which invoices are produced.
// If unspecified, a single invoice is produced per billing cycle.
type PlanVersionNewParamsAddPricesPriceNewPlanPackagePriceInvoicingCycleConfiguration struct {
	// The duration of the billing period.
	Duration param.Field[int64] `json:"duration,required"`
	// The unit of billing period duration.
	DurationUnit param.Field[PlanVersionNewParamsAddPricesPriceNewPlanPackagePriceInvoicingCycleConfigurationDurationUnit] `json:"duration_unit,required"`
}

func (r PlanVersionNewParamsAddPricesPriceNewPlanPackagePriceInvoicingCycleConfiguration) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// The unit of billing period duration.
type PlanVersionNewParamsAddPricesPriceNewPlanPackagePriceInvoicingCycleConfigurationDurationUnit string

const (
	PlanVersionNewParamsAddPricesPriceNewPlanPackagePriceInvoicingCycleConfigurationDurationUnitDay   PlanVersionNewParamsAddPricesPriceNewPlanPackagePriceInvoicingCycleConfigurationDurationUnit = "day"
	PlanVersionNewParamsAddPricesPriceNewPlanPackagePriceInvoicingCycleConfigurationDurationUnitMonth PlanVersionNewParamsAddPricesPriceNewPlanPackagePriceInvoicingCycleConfigurationDurationUnit = "month"
)

func (r PlanVersionNewParamsAddPricesPriceNewPlanPackagePriceInvoicingCycleConfigurationDurationUnit) IsKnown() bool {
	switch r {
	case PlanVersionNewParamsAddPricesPriceNewPlanPackagePriceInvoicingCycleConfigurationDurationUnitDay, PlanVersionNewParamsAddPricesPriceNewPlanPackagePriceInvoicingCycleConfigurationDurationUnitMonth:
		return true
	}
	return false
}

type PlanVersionNewParamsAddPricesPriceNewPlanMatrixPrice struct {
	// The cadence to bill for this price on.
	Cadence param.Field[PlanVersionNewParamsAddPricesPriceNewPlanMatrixPriceCadence] `json:"cadence,required"`
	// The id of the item the price will be associated with.
	ItemID       param.Field[string]                                                           `json:"item_id,required"`
	MatrixConfig param.Field[PlanVersionNewParamsAddPricesPriceNewPlanMatrixPriceMatrixConfig] `json:"matrix_config,required"`
	ModelType    param.Field[PlanVersionNewParamsAddPricesPriceNewPlanMatrixPriceModelType]    `json:"model_type,required"`
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
	BillingCycleConfiguration param.Field[PlanVersionNewParamsAddPricesPriceNewPlanMatrixPriceBillingCycleConfiguration] `json:"billing_cycle_configuration"`
	// The per unit conversion rate of the price currency to the invoicing currency.
	ConversionRate param.Field[float64] `json:"conversion_rate"`
	// An ISO 4217 currency string, or custom pricing unit identifier, in which this
	// price is billed.
	Currency param.Field[string] `json:"currency"`
	// For dimensional price: specifies a price group and dimension values
	DimensionalPriceConfiguration param.Field[PlanVersionNewParamsAddPricesPriceNewPlanMatrixPriceDimensionalPriceConfiguration] `json:"dimensional_price_configuration"`
	// An alias for the price.
	ExternalPriceID param.Field[string] `json:"external_price_id"`
	// If the Price represents a fixed cost, this represents the quantity of units
	// applied.
	FixedPriceQuantity param.Field[float64] `json:"fixed_price_quantity"`
	// The property used to group this price on an invoice
	InvoiceGroupingKey param.Field[string] `json:"invoice_grouping_key"`
	// Within each billing cycle, specifies the cadence at which invoices are produced.
	// If unspecified, a single invoice is produced per billing cycle.
	InvoicingCycleConfiguration param.Field[PlanVersionNewParamsAddPricesPriceNewPlanMatrixPriceInvoicingCycleConfiguration] `json:"invoicing_cycle_configuration"`
	// User-specified key/value pairs for the resource. Individual keys can be removed
	// by setting the value to `null`, and the entire metadata mapping can be cleared
	// by setting `metadata` to `null`.
	Metadata param.Field[map[string]string] `json:"metadata"`
}

func (r PlanVersionNewParamsAddPricesPriceNewPlanMatrixPrice) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r PlanVersionNewParamsAddPricesPriceNewPlanMatrixPrice) implementsPlanVersionNewParamsAddPricesPriceUnion() {
}

// The cadence to bill for this price on.
type PlanVersionNewParamsAddPricesPriceNewPlanMatrixPriceCadence string

const (
	PlanVersionNewParamsAddPricesPriceNewPlanMatrixPriceCadenceAnnual     PlanVersionNewParamsAddPricesPriceNewPlanMatrixPriceCadence = "annual"
	PlanVersionNewParamsAddPricesPriceNewPlanMatrixPriceCadenceSemiAnnual PlanVersionNewParamsAddPricesPriceNewPlanMatrixPriceCadence = "semi_annual"
	PlanVersionNewParamsAddPricesPriceNewPlanMatrixPriceCadenceMonthly    PlanVersionNewParamsAddPricesPriceNewPlanMatrixPriceCadence = "monthly"
	PlanVersionNewParamsAddPricesPriceNewPlanMatrixPriceCadenceQuarterly  PlanVersionNewParamsAddPricesPriceNewPlanMatrixPriceCadence = "quarterly"
	PlanVersionNewParamsAddPricesPriceNewPlanMatrixPriceCadenceOneTime    PlanVersionNewParamsAddPricesPriceNewPlanMatrixPriceCadence = "one_time"
	PlanVersionNewParamsAddPricesPriceNewPlanMatrixPriceCadenceCustom     PlanVersionNewParamsAddPricesPriceNewPlanMatrixPriceCadence = "custom"
)

func (r PlanVersionNewParamsAddPricesPriceNewPlanMatrixPriceCadence) IsKnown() bool {
	switch r {
	case PlanVersionNewParamsAddPricesPriceNewPlanMatrixPriceCadenceAnnual, PlanVersionNewParamsAddPricesPriceNewPlanMatrixPriceCadenceSemiAnnual, PlanVersionNewParamsAddPricesPriceNewPlanMatrixPriceCadenceMonthly, PlanVersionNewParamsAddPricesPriceNewPlanMatrixPriceCadenceQuarterly, PlanVersionNewParamsAddPricesPriceNewPlanMatrixPriceCadenceOneTime, PlanVersionNewParamsAddPricesPriceNewPlanMatrixPriceCadenceCustom:
		return true
	}
	return false
}

type PlanVersionNewParamsAddPricesPriceNewPlanMatrixPriceMatrixConfig struct {
	// Default per unit rate for any usage not bucketed into a specified matrix_value
	DefaultUnitAmount param.Field[string] `json:"default_unit_amount,required"`
	// One or two event property values to evaluate matrix groups by
	Dimensions param.Field[[]string] `json:"dimensions,required"`
	// Matrix values for specified matrix grouping keys
	MatrixValues param.Field[[]PlanVersionNewParamsAddPricesPriceNewPlanMatrixPriceMatrixConfigMatrixValue] `json:"matrix_values,required"`
}

func (r PlanVersionNewParamsAddPricesPriceNewPlanMatrixPriceMatrixConfig) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type PlanVersionNewParamsAddPricesPriceNewPlanMatrixPriceMatrixConfigMatrixValue struct {
	// One or two matrix keys to filter usage to this Matrix value by. For example,
	// ["region", "tier"] could be used to filter cloud usage by a cloud region and an
	// instance tier.
	DimensionValues param.Field[[]string] `json:"dimension_values,required"`
	// Unit price for the specified dimension_values
	UnitAmount param.Field[string] `json:"unit_amount,required"`
}

func (r PlanVersionNewParamsAddPricesPriceNewPlanMatrixPriceMatrixConfigMatrixValue) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type PlanVersionNewParamsAddPricesPriceNewPlanMatrixPriceModelType string

const (
	PlanVersionNewParamsAddPricesPriceNewPlanMatrixPriceModelTypeMatrix PlanVersionNewParamsAddPricesPriceNewPlanMatrixPriceModelType = "matrix"
)

func (r PlanVersionNewParamsAddPricesPriceNewPlanMatrixPriceModelType) IsKnown() bool {
	switch r {
	case PlanVersionNewParamsAddPricesPriceNewPlanMatrixPriceModelTypeMatrix:
		return true
	}
	return false
}

// For custom cadence: specifies the duration of the billing period in days or
// months.
type PlanVersionNewParamsAddPricesPriceNewPlanMatrixPriceBillingCycleConfiguration struct {
	// The duration of the billing period.
	Duration param.Field[int64] `json:"duration,required"`
	// The unit of billing period duration.
	DurationUnit param.Field[PlanVersionNewParamsAddPricesPriceNewPlanMatrixPriceBillingCycleConfigurationDurationUnit] `json:"duration_unit,required"`
}

func (r PlanVersionNewParamsAddPricesPriceNewPlanMatrixPriceBillingCycleConfiguration) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// The unit of billing period duration.
type PlanVersionNewParamsAddPricesPriceNewPlanMatrixPriceBillingCycleConfigurationDurationUnit string

const (
	PlanVersionNewParamsAddPricesPriceNewPlanMatrixPriceBillingCycleConfigurationDurationUnitDay   PlanVersionNewParamsAddPricesPriceNewPlanMatrixPriceBillingCycleConfigurationDurationUnit = "day"
	PlanVersionNewParamsAddPricesPriceNewPlanMatrixPriceBillingCycleConfigurationDurationUnitMonth PlanVersionNewParamsAddPricesPriceNewPlanMatrixPriceBillingCycleConfigurationDurationUnit = "month"
)

func (r PlanVersionNewParamsAddPricesPriceNewPlanMatrixPriceBillingCycleConfigurationDurationUnit) IsKnown() bool {
	switch r {
	case PlanVersionNewParamsAddPricesPriceNewPlanMatrixPriceBillingCycleConfigurationDurationUnitDay, PlanVersionNewParamsAddPricesPriceNewPlanMatrixPriceBillingCycleConfigurationDurationUnitMonth:
		return true
	}
	return false
}

// For dimensional price: specifies a price group and dimension values
type PlanVersionNewParamsAddPricesPriceNewPlanMatrixPriceDimensionalPriceConfiguration struct {
	// The list of dimension values matching (in order) the dimensions of the price
	// group
	DimensionValues param.Field[[]string] `json:"dimension_values,required"`
	// The id of the dimensional price group to include this price in
	DimensionalPriceGroupID param.Field[string] `json:"dimensional_price_group_id"`
	// The external id of the dimensional price group to include this price in
	ExternalDimensionalPriceGroupID param.Field[string] `json:"external_dimensional_price_group_id"`
}

func (r PlanVersionNewParamsAddPricesPriceNewPlanMatrixPriceDimensionalPriceConfiguration) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Within each billing cycle, specifies the cadence at which invoices are produced.
// If unspecified, a single invoice is produced per billing cycle.
type PlanVersionNewParamsAddPricesPriceNewPlanMatrixPriceInvoicingCycleConfiguration struct {
	// The duration of the billing period.
	Duration param.Field[int64] `json:"duration,required"`
	// The unit of billing period duration.
	DurationUnit param.Field[PlanVersionNewParamsAddPricesPriceNewPlanMatrixPriceInvoicingCycleConfigurationDurationUnit] `json:"duration_unit,required"`
}

func (r PlanVersionNewParamsAddPricesPriceNewPlanMatrixPriceInvoicingCycleConfiguration) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// The unit of billing period duration.
type PlanVersionNewParamsAddPricesPriceNewPlanMatrixPriceInvoicingCycleConfigurationDurationUnit string

const (
	PlanVersionNewParamsAddPricesPriceNewPlanMatrixPriceInvoicingCycleConfigurationDurationUnitDay   PlanVersionNewParamsAddPricesPriceNewPlanMatrixPriceInvoicingCycleConfigurationDurationUnit = "day"
	PlanVersionNewParamsAddPricesPriceNewPlanMatrixPriceInvoicingCycleConfigurationDurationUnitMonth PlanVersionNewParamsAddPricesPriceNewPlanMatrixPriceInvoicingCycleConfigurationDurationUnit = "month"
)

func (r PlanVersionNewParamsAddPricesPriceNewPlanMatrixPriceInvoicingCycleConfigurationDurationUnit) IsKnown() bool {
	switch r {
	case PlanVersionNewParamsAddPricesPriceNewPlanMatrixPriceInvoicingCycleConfigurationDurationUnitDay, PlanVersionNewParamsAddPricesPriceNewPlanMatrixPriceInvoicingCycleConfigurationDurationUnitMonth:
		return true
	}
	return false
}

type PlanVersionNewParamsAddPricesPriceNewPlanTieredPrice struct {
	// The cadence to bill for this price on.
	Cadence param.Field[PlanVersionNewParamsAddPricesPriceNewPlanTieredPriceCadence] `json:"cadence,required"`
	// The id of the item the price will be associated with.
	ItemID    param.Field[string]                                                        `json:"item_id,required"`
	ModelType param.Field[PlanVersionNewParamsAddPricesPriceNewPlanTieredPriceModelType] `json:"model_type,required"`
	// The name of the price.
	Name         param.Field[string]                                                           `json:"name,required"`
	TieredConfig param.Field[PlanVersionNewParamsAddPricesPriceNewPlanTieredPriceTieredConfig] `json:"tiered_config,required"`
	// The id of the billable metric for the price. Only needed if the price is
	// usage-based.
	BillableMetricID param.Field[string] `json:"billable_metric_id"`
	// If the Price represents a fixed cost, the price will be billed in-advance if
	// this is true, and in-arrears if this is false.
	BilledInAdvance param.Field[bool] `json:"billed_in_advance"`
	// For custom cadence: specifies the duration of the billing period in days or
	// months.
	BillingCycleConfiguration param.Field[PlanVersionNewParamsAddPricesPriceNewPlanTieredPriceBillingCycleConfiguration] `json:"billing_cycle_configuration"`
	// The per unit conversion rate of the price currency to the invoicing currency.
	ConversionRate param.Field[float64] `json:"conversion_rate"`
	// An ISO 4217 currency string, or custom pricing unit identifier, in which this
	// price is billed.
	Currency param.Field[string] `json:"currency"`
	// For dimensional price: specifies a price group and dimension values
	DimensionalPriceConfiguration param.Field[PlanVersionNewParamsAddPricesPriceNewPlanTieredPriceDimensionalPriceConfiguration] `json:"dimensional_price_configuration"`
	// An alias for the price.
	ExternalPriceID param.Field[string] `json:"external_price_id"`
	// If the Price represents a fixed cost, this represents the quantity of units
	// applied.
	FixedPriceQuantity param.Field[float64] `json:"fixed_price_quantity"`
	// The property used to group this price on an invoice
	InvoiceGroupingKey param.Field[string] `json:"invoice_grouping_key"`
	// Within each billing cycle, specifies the cadence at which invoices are produced.
	// If unspecified, a single invoice is produced per billing cycle.
	InvoicingCycleConfiguration param.Field[PlanVersionNewParamsAddPricesPriceNewPlanTieredPriceInvoicingCycleConfiguration] `json:"invoicing_cycle_configuration"`
	// User-specified key/value pairs for the resource. Individual keys can be removed
	// by setting the value to `null`, and the entire metadata mapping can be cleared
	// by setting `metadata` to `null`.
	Metadata param.Field[map[string]string] `json:"metadata"`
}

func (r PlanVersionNewParamsAddPricesPriceNewPlanTieredPrice) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r PlanVersionNewParamsAddPricesPriceNewPlanTieredPrice) implementsPlanVersionNewParamsAddPricesPriceUnion() {
}

// The cadence to bill for this price on.
type PlanVersionNewParamsAddPricesPriceNewPlanTieredPriceCadence string

const (
	PlanVersionNewParamsAddPricesPriceNewPlanTieredPriceCadenceAnnual     PlanVersionNewParamsAddPricesPriceNewPlanTieredPriceCadence = "annual"
	PlanVersionNewParamsAddPricesPriceNewPlanTieredPriceCadenceSemiAnnual PlanVersionNewParamsAddPricesPriceNewPlanTieredPriceCadence = "semi_annual"
	PlanVersionNewParamsAddPricesPriceNewPlanTieredPriceCadenceMonthly    PlanVersionNewParamsAddPricesPriceNewPlanTieredPriceCadence = "monthly"
	PlanVersionNewParamsAddPricesPriceNewPlanTieredPriceCadenceQuarterly  PlanVersionNewParamsAddPricesPriceNewPlanTieredPriceCadence = "quarterly"
	PlanVersionNewParamsAddPricesPriceNewPlanTieredPriceCadenceOneTime    PlanVersionNewParamsAddPricesPriceNewPlanTieredPriceCadence = "one_time"
	PlanVersionNewParamsAddPricesPriceNewPlanTieredPriceCadenceCustom     PlanVersionNewParamsAddPricesPriceNewPlanTieredPriceCadence = "custom"
)

func (r PlanVersionNewParamsAddPricesPriceNewPlanTieredPriceCadence) IsKnown() bool {
	switch r {
	case PlanVersionNewParamsAddPricesPriceNewPlanTieredPriceCadenceAnnual, PlanVersionNewParamsAddPricesPriceNewPlanTieredPriceCadenceSemiAnnual, PlanVersionNewParamsAddPricesPriceNewPlanTieredPriceCadenceMonthly, PlanVersionNewParamsAddPricesPriceNewPlanTieredPriceCadenceQuarterly, PlanVersionNewParamsAddPricesPriceNewPlanTieredPriceCadenceOneTime, PlanVersionNewParamsAddPricesPriceNewPlanTieredPriceCadenceCustom:
		return true
	}
	return false
}

type PlanVersionNewParamsAddPricesPriceNewPlanTieredPriceModelType string

const (
	PlanVersionNewParamsAddPricesPriceNewPlanTieredPriceModelTypeTiered PlanVersionNewParamsAddPricesPriceNewPlanTieredPriceModelType = "tiered"
)

func (r PlanVersionNewParamsAddPricesPriceNewPlanTieredPriceModelType) IsKnown() bool {
	switch r {
	case PlanVersionNewParamsAddPricesPriceNewPlanTieredPriceModelTypeTiered:
		return true
	}
	return false
}

type PlanVersionNewParamsAddPricesPriceNewPlanTieredPriceTieredConfig struct {
	// Tiers for rating based on total usage quantities into the specified tier
	Tiers param.Field[[]PlanVersionNewParamsAddPricesPriceNewPlanTieredPriceTieredConfigTier] `json:"tiers,required"`
}

func (r PlanVersionNewParamsAddPricesPriceNewPlanTieredPriceTieredConfig) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type PlanVersionNewParamsAddPricesPriceNewPlanTieredPriceTieredConfigTier struct {
	// Exclusive tier starting value
	FirstUnit param.Field[float64] `json:"first_unit,required"`
	// Amount per unit
	UnitAmount param.Field[string] `json:"unit_amount,required"`
	// Inclusive tier ending value. If null, this is treated as the last tier
	LastUnit param.Field[float64] `json:"last_unit"`
}

func (r PlanVersionNewParamsAddPricesPriceNewPlanTieredPriceTieredConfigTier) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// For custom cadence: specifies the duration of the billing period in days or
// months.
type PlanVersionNewParamsAddPricesPriceNewPlanTieredPriceBillingCycleConfiguration struct {
	// The duration of the billing period.
	Duration param.Field[int64] `json:"duration,required"`
	// The unit of billing period duration.
	DurationUnit param.Field[PlanVersionNewParamsAddPricesPriceNewPlanTieredPriceBillingCycleConfigurationDurationUnit] `json:"duration_unit,required"`
}

func (r PlanVersionNewParamsAddPricesPriceNewPlanTieredPriceBillingCycleConfiguration) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// The unit of billing period duration.
type PlanVersionNewParamsAddPricesPriceNewPlanTieredPriceBillingCycleConfigurationDurationUnit string

const (
	PlanVersionNewParamsAddPricesPriceNewPlanTieredPriceBillingCycleConfigurationDurationUnitDay   PlanVersionNewParamsAddPricesPriceNewPlanTieredPriceBillingCycleConfigurationDurationUnit = "day"
	PlanVersionNewParamsAddPricesPriceNewPlanTieredPriceBillingCycleConfigurationDurationUnitMonth PlanVersionNewParamsAddPricesPriceNewPlanTieredPriceBillingCycleConfigurationDurationUnit = "month"
)

func (r PlanVersionNewParamsAddPricesPriceNewPlanTieredPriceBillingCycleConfigurationDurationUnit) IsKnown() bool {
	switch r {
	case PlanVersionNewParamsAddPricesPriceNewPlanTieredPriceBillingCycleConfigurationDurationUnitDay, PlanVersionNewParamsAddPricesPriceNewPlanTieredPriceBillingCycleConfigurationDurationUnitMonth:
		return true
	}
	return false
}

// For dimensional price: specifies a price group and dimension values
type PlanVersionNewParamsAddPricesPriceNewPlanTieredPriceDimensionalPriceConfiguration struct {
	// The list of dimension values matching (in order) the dimensions of the price
	// group
	DimensionValues param.Field[[]string] `json:"dimension_values,required"`
	// The id of the dimensional price group to include this price in
	DimensionalPriceGroupID param.Field[string] `json:"dimensional_price_group_id"`
	// The external id of the dimensional price group to include this price in
	ExternalDimensionalPriceGroupID param.Field[string] `json:"external_dimensional_price_group_id"`
}

func (r PlanVersionNewParamsAddPricesPriceNewPlanTieredPriceDimensionalPriceConfiguration) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Within each billing cycle, specifies the cadence at which invoices are produced.
// If unspecified, a single invoice is produced per billing cycle.
type PlanVersionNewParamsAddPricesPriceNewPlanTieredPriceInvoicingCycleConfiguration struct {
	// The duration of the billing period.
	Duration param.Field[int64] `json:"duration,required"`
	// The unit of billing period duration.
	DurationUnit param.Field[PlanVersionNewParamsAddPricesPriceNewPlanTieredPriceInvoicingCycleConfigurationDurationUnit] `json:"duration_unit,required"`
}

func (r PlanVersionNewParamsAddPricesPriceNewPlanTieredPriceInvoicingCycleConfiguration) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// The unit of billing period duration.
type PlanVersionNewParamsAddPricesPriceNewPlanTieredPriceInvoicingCycleConfigurationDurationUnit string

const (
	PlanVersionNewParamsAddPricesPriceNewPlanTieredPriceInvoicingCycleConfigurationDurationUnitDay   PlanVersionNewParamsAddPricesPriceNewPlanTieredPriceInvoicingCycleConfigurationDurationUnit = "day"
	PlanVersionNewParamsAddPricesPriceNewPlanTieredPriceInvoicingCycleConfigurationDurationUnitMonth PlanVersionNewParamsAddPricesPriceNewPlanTieredPriceInvoicingCycleConfigurationDurationUnit = "month"
)

func (r PlanVersionNewParamsAddPricesPriceNewPlanTieredPriceInvoicingCycleConfigurationDurationUnit) IsKnown() bool {
	switch r {
	case PlanVersionNewParamsAddPricesPriceNewPlanTieredPriceInvoicingCycleConfigurationDurationUnitDay, PlanVersionNewParamsAddPricesPriceNewPlanTieredPriceInvoicingCycleConfigurationDurationUnitMonth:
		return true
	}
	return false
}

type PlanVersionNewParamsAddPricesPriceNewPlanTieredBpsPrice struct {
	// The cadence to bill for this price on.
	Cadence param.Field[PlanVersionNewParamsAddPricesPriceNewPlanTieredBpsPriceCadence] `json:"cadence,required"`
	// The id of the item the price will be associated with.
	ItemID    param.Field[string]                                                           `json:"item_id,required"`
	ModelType param.Field[PlanVersionNewParamsAddPricesPriceNewPlanTieredBpsPriceModelType] `json:"model_type,required"`
	// The name of the price.
	Name            param.Field[string]                                                                 `json:"name,required"`
	TieredBpsConfig param.Field[PlanVersionNewParamsAddPricesPriceNewPlanTieredBpsPriceTieredBpsConfig] `json:"tiered_bps_config,required"`
	// The id of the billable metric for the price. Only needed if the price is
	// usage-based.
	BillableMetricID param.Field[string] `json:"billable_metric_id"`
	// If the Price represents a fixed cost, the price will be billed in-advance if
	// this is true, and in-arrears if this is false.
	BilledInAdvance param.Field[bool] `json:"billed_in_advance"`
	// For custom cadence: specifies the duration of the billing period in days or
	// months.
	BillingCycleConfiguration param.Field[PlanVersionNewParamsAddPricesPriceNewPlanTieredBpsPriceBillingCycleConfiguration] `json:"billing_cycle_configuration"`
	// The per unit conversion rate of the price currency to the invoicing currency.
	ConversionRate param.Field[float64] `json:"conversion_rate"`
	// An ISO 4217 currency string, or custom pricing unit identifier, in which this
	// price is billed.
	Currency param.Field[string] `json:"currency"`
	// For dimensional price: specifies a price group and dimension values
	DimensionalPriceConfiguration param.Field[PlanVersionNewParamsAddPricesPriceNewPlanTieredBpsPriceDimensionalPriceConfiguration] `json:"dimensional_price_configuration"`
	// An alias for the price.
	ExternalPriceID param.Field[string] `json:"external_price_id"`
	// If the Price represents a fixed cost, this represents the quantity of units
	// applied.
	FixedPriceQuantity param.Field[float64] `json:"fixed_price_quantity"`
	// The property used to group this price on an invoice
	InvoiceGroupingKey param.Field[string] `json:"invoice_grouping_key"`
	// Within each billing cycle, specifies the cadence at which invoices are produced.
	// If unspecified, a single invoice is produced per billing cycle.
	InvoicingCycleConfiguration param.Field[PlanVersionNewParamsAddPricesPriceNewPlanTieredBpsPriceInvoicingCycleConfiguration] `json:"invoicing_cycle_configuration"`
	// User-specified key/value pairs for the resource. Individual keys can be removed
	// by setting the value to `null`, and the entire metadata mapping can be cleared
	// by setting `metadata` to `null`.
	Metadata param.Field[map[string]string] `json:"metadata"`
}

func (r PlanVersionNewParamsAddPricesPriceNewPlanTieredBpsPrice) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r PlanVersionNewParamsAddPricesPriceNewPlanTieredBpsPrice) implementsPlanVersionNewParamsAddPricesPriceUnion() {
}

// The cadence to bill for this price on.
type PlanVersionNewParamsAddPricesPriceNewPlanTieredBpsPriceCadence string

const (
	PlanVersionNewParamsAddPricesPriceNewPlanTieredBpsPriceCadenceAnnual     PlanVersionNewParamsAddPricesPriceNewPlanTieredBpsPriceCadence = "annual"
	PlanVersionNewParamsAddPricesPriceNewPlanTieredBpsPriceCadenceSemiAnnual PlanVersionNewParamsAddPricesPriceNewPlanTieredBpsPriceCadence = "semi_annual"
	PlanVersionNewParamsAddPricesPriceNewPlanTieredBpsPriceCadenceMonthly    PlanVersionNewParamsAddPricesPriceNewPlanTieredBpsPriceCadence = "monthly"
	PlanVersionNewParamsAddPricesPriceNewPlanTieredBpsPriceCadenceQuarterly  PlanVersionNewParamsAddPricesPriceNewPlanTieredBpsPriceCadence = "quarterly"
	PlanVersionNewParamsAddPricesPriceNewPlanTieredBpsPriceCadenceOneTime    PlanVersionNewParamsAddPricesPriceNewPlanTieredBpsPriceCadence = "one_time"
	PlanVersionNewParamsAddPricesPriceNewPlanTieredBpsPriceCadenceCustom     PlanVersionNewParamsAddPricesPriceNewPlanTieredBpsPriceCadence = "custom"
)

func (r PlanVersionNewParamsAddPricesPriceNewPlanTieredBpsPriceCadence) IsKnown() bool {
	switch r {
	case PlanVersionNewParamsAddPricesPriceNewPlanTieredBpsPriceCadenceAnnual, PlanVersionNewParamsAddPricesPriceNewPlanTieredBpsPriceCadenceSemiAnnual, PlanVersionNewParamsAddPricesPriceNewPlanTieredBpsPriceCadenceMonthly, PlanVersionNewParamsAddPricesPriceNewPlanTieredBpsPriceCadenceQuarterly, PlanVersionNewParamsAddPricesPriceNewPlanTieredBpsPriceCadenceOneTime, PlanVersionNewParamsAddPricesPriceNewPlanTieredBpsPriceCadenceCustom:
		return true
	}
	return false
}

type PlanVersionNewParamsAddPricesPriceNewPlanTieredBpsPriceModelType string

const (
	PlanVersionNewParamsAddPricesPriceNewPlanTieredBpsPriceModelTypeTieredBps PlanVersionNewParamsAddPricesPriceNewPlanTieredBpsPriceModelType = "tiered_bps"
)

func (r PlanVersionNewParamsAddPricesPriceNewPlanTieredBpsPriceModelType) IsKnown() bool {
	switch r {
	case PlanVersionNewParamsAddPricesPriceNewPlanTieredBpsPriceModelTypeTieredBps:
		return true
	}
	return false
}

type PlanVersionNewParamsAddPricesPriceNewPlanTieredBpsPriceTieredBpsConfig struct {
	// Tiers for a Graduated BPS pricing model, where usage is bucketed into specified
	// tiers
	Tiers param.Field[[]PlanVersionNewParamsAddPricesPriceNewPlanTieredBpsPriceTieredBpsConfigTier] `json:"tiers,required"`
}

func (r PlanVersionNewParamsAddPricesPriceNewPlanTieredBpsPriceTieredBpsConfig) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type PlanVersionNewParamsAddPricesPriceNewPlanTieredBpsPriceTieredBpsConfigTier struct {
	// Per-event basis point rate
	Bps param.Field[float64] `json:"bps,required"`
	// Exclusive tier starting value
	MinimumAmount param.Field[string] `json:"minimum_amount,required"`
	// Inclusive tier ending value
	MaximumAmount param.Field[string] `json:"maximum_amount"`
	// Per unit maximum to charge
	PerUnitMaximum param.Field[string] `json:"per_unit_maximum"`
}

func (r PlanVersionNewParamsAddPricesPriceNewPlanTieredBpsPriceTieredBpsConfigTier) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// For custom cadence: specifies the duration of the billing period in days or
// months.
type PlanVersionNewParamsAddPricesPriceNewPlanTieredBpsPriceBillingCycleConfiguration struct {
	// The duration of the billing period.
	Duration param.Field[int64] `json:"duration,required"`
	// The unit of billing period duration.
	DurationUnit param.Field[PlanVersionNewParamsAddPricesPriceNewPlanTieredBpsPriceBillingCycleConfigurationDurationUnit] `json:"duration_unit,required"`
}

func (r PlanVersionNewParamsAddPricesPriceNewPlanTieredBpsPriceBillingCycleConfiguration) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// The unit of billing period duration.
type PlanVersionNewParamsAddPricesPriceNewPlanTieredBpsPriceBillingCycleConfigurationDurationUnit string

const (
	PlanVersionNewParamsAddPricesPriceNewPlanTieredBpsPriceBillingCycleConfigurationDurationUnitDay   PlanVersionNewParamsAddPricesPriceNewPlanTieredBpsPriceBillingCycleConfigurationDurationUnit = "day"
	PlanVersionNewParamsAddPricesPriceNewPlanTieredBpsPriceBillingCycleConfigurationDurationUnitMonth PlanVersionNewParamsAddPricesPriceNewPlanTieredBpsPriceBillingCycleConfigurationDurationUnit = "month"
)

func (r PlanVersionNewParamsAddPricesPriceNewPlanTieredBpsPriceBillingCycleConfigurationDurationUnit) IsKnown() bool {
	switch r {
	case PlanVersionNewParamsAddPricesPriceNewPlanTieredBpsPriceBillingCycleConfigurationDurationUnitDay, PlanVersionNewParamsAddPricesPriceNewPlanTieredBpsPriceBillingCycleConfigurationDurationUnitMonth:
		return true
	}
	return false
}

// For dimensional price: specifies a price group and dimension values
type PlanVersionNewParamsAddPricesPriceNewPlanTieredBpsPriceDimensionalPriceConfiguration struct {
	// The list of dimension values matching (in order) the dimensions of the price
	// group
	DimensionValues param.Field[[]string] `json:"dimension_values,required"`
	// The id of the dimensional price group to include this price in
	DimensionalPriceGroupID param.Field[string] `json:"dimensional_price_group_id"`
	// The external id of the dimensional price group to include this price in
	ExternalDimensionalPriceGroupID param.Field[string] `json:"external_dimensional_price_group_id"`
}

func (r PlanVersionNewParamsAddPricesPriceNewPlanTieredBpsPriceDimensionalPriceConfiguration) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Within each billing cycle, specifies the cadence at which invoices are produced.
// If unspecified, a single invoice is produced per billing cycle.
type PlanVersionNewParamsAddPricesPriceNewPlanTieredBpsPriceInvoicingCycleConfiguration struct {
	// The duration of the billing period.
	Duration param.Field[int64] `json:"duration,required"`
	// The unit of billing period duration.
	DurationUnit param.Field[PlanVersionNewParamsAddPricesPriceNewPlanTieredBpsPriceInvoicingCycleConfigurationDurationUnit] `json:"duration_unit,required"`
}

func (r PlanVersionNewParamsAddPricesPriceNewPlanTieredBpsPriceInvoicingCycleConfiguration) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// The unit of billing period duration.
type PlanVersionNewParamsAddPricesPriceNewPlanTieredBpsPriceInvoicingCycleConfigurationDurationUnit string

const (
	PlanVersionNewParamsAddPricesPriceNewPlanTieredBpsPriceInvoicingCycleConfigurationDurationUnitDay   PlanVersionNewParamsAddPricesPriceNewPlanTieredBpsPriceInvoicingCycleConfigurationDurationUnit = "day"
	PlanVersionNewParamsAddPricesPriceNewPlanTieredBpsPriceInvoicingCycleConfigurationDurationUnitMonth PlanVersionNewParamsAddPricesPriceNewPlanTieredBpsPriceInvoicingCycleConfigurationDurationUnit = "month"
)

func (r PlanVersionNewParamsAddPricesPriceNewPlanTieredBpsPriceInvoicingCycleConfigurationDurationUnit) IsKnown() bool {
	switch r {
	case PlanVersionNewParamsAddPricesPriceNewPlanTieredBpsPriceInvoicingCycleConfigurationDurationUnitDay, PlanVersionNewParamsAddPricesPriceNewPlanTieredBpsPriceInvoicingCycleConfigurationDurationUnitMonth:
		return true
	}
	return false
}

type PlanVersionNewParamsAddPricesPriceNewPlanBpsPrice struct {
	BpsConfig param.Field[PlanVersionNewParamsAddPricesPriceNewPlanBpsPriceBpsConfig] `json:"bps_config,required"`
	// The cadence to bill for this price on.
	Cadence param.Field[PlanVersionNewParamsAddPricesPriceNewPlanBpsPriceCadence] `json:"cadence,required"`
	// The id of the item the price will be associated with.
	ItemID    param.Field[string]                                                     `json:"item_id,required"`
	ModelType param.Field[PlanVersionNewParamsAddPricesPriceNewPlanBpsPriceModelType] `json:"model_type,required"`
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
	BillingCycleConfiguration param.Field[PlanVersionNewParamsAddPricesPriceNewPlanBpsPriceBillingCycleConfiguration] `json:"billing_cycle_configuration"`
	// The per unit conversion rate of the price currency to the invoicing currency.
	ConversionRate param.Field[float64] `json:"conversion_rate"`
	// An ISO 4217 currency string, or custom pricing unit identifier, in which this
	// price is billed.
	Currency param.Field[string] `json:"currency"`
	// For dimensional price: specifies a price group and dimension values
	DimensionalPriceConfiguration param.Field[PlanVersionNewParamsAddPricesPriceNewPlanBpsPriceDimensionalPriceConfiguration] `json:"dimensional_price_configuration"`
	// An alias for the price.
	ExternalPriceID param.Field[string] `json:"external_price_id"`
	// If the Price represents a fixed cost, this represents the quantity of units
	// applied.
	FixedPriceQuantity param.Field[float64] `json:"fixed_price_quantity"`
	// The property used to group this price on an invoice
	InvoiceGroupingKey param.Field[string] `json:"invoice_grouping_key"`
	// Within each billing cycle, specifies the cadence at which invoices are produced.
	// If unspecified, a single invoice is produced per billing cycle.
	InvoicingCycleConfiguration param.Field[PlanVersionNewParamsAddPricesPriceNewPlanBpsPriceInvoicingCycleConfiguration] `json:"invoicing_cycle_configuration"`
	// User-specified key/value pairs for the resource. Individual keys can be removed
	// by setting the value to `null`, and the entire metadata mapping can be cleared
	// by setting `metadata` to `null`.
	Metadata param.Field[map[string]string] `json:"metadata"`
}

func (r PlanVersionNewParamsAddPricesPriceNewPlanBpsPrice) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r PlanVersionNewParamsAddPricesPriceNewPlanBpsPrice) implementsPlanVersionNewParamsAddPricesPriceUnion() {
}

type PlanVersionNewParamsAddPricesPriceNewPlanBpsPriceBpsConfig struct {
	// Basis point take rate per event
	Bps param.Field[float64] `json:"bps,required"`
	// Optional currency amount maximum to cap spend per event
	PerUnitMaximum param.Field[string] `json:"per_unit_maximum"`
}

func (r PlanVersionNewParamsAddPricesPriceNewPlanBpsPriceBpsConfig) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// The cadence to bill for this price on.
type PlanVersionNewParamsAddPricesPriceNewPlanBpsPriceCadence string

const (
	PlanVersionNewParamsAddPricesPriceNewPlanBpsPriceCadenceAnnual     PlanVersionNewParamsAddPricesPriceNewPlanBpsPriceCadence = "annual"
	PlanVersionNewParamsAddPricesPriceNewPlanBpsPriceCadenceSemiAnnual PlanVersionNewParamsAddPricesPriceNewPlanBpsPriceCadence = "semi_annual"
	PlanVersionNewParamsAddPricesPriceNewPlanBpsPriceCadenceMonthly    PlanVersionNewParamsAddPricesPriceNewPlanBpsPriceCadence = "monthly"
	PlanVersionNewParamsAddPricesPriceNewPlanBpsPriceCadenceQuarterly  PlanVersionNewParamsAddPricesPriceNewPlanBpsPriceCadence = "quarterly"
	PlanVersionNewParamsAddPricesPriceNewPlanBpsPriceCadenceOneTime    PlanVersionNewParamsAddPricesPriceNewPlanBpsPriceCadence = "one_time"
	PlanVersionNewParamsAddPricesPriceNewPlanBpsPriceCadenceCustom     PlanVersionNewParamsAddPricesPriceNewPlanBpsPriceCadence = "custom"
)

func (r PlanVersionNewParamsAddPricesPriceNewPlanBpsPriceCadence) IsKnown() bool {
	switch r {
	case PlanVersionNewParamsAddPricesPriceNewPlanBpsPriceCadenceAnnual, PlanVersionNewParamsAddPricesPriceNewPlanBpsPriceCadenceSemiAnnual, PlanVersionNewParamsAddPricesPriceNewPlanBpsPriceCadenceMonthly, PlanVersionNewParamsAddPricesPriceNewPlanBpsPriceCadenceQuarterly, PlanVersionNewParamsAddPricesPriceNewPlanBpsPriceCadenceOneTime, PlanVersionNewParamsAddPricesPriceNewPlanBpsPriceCadenceCustom:
		return true
	}
	return false
}

type PlanVersionNewParamsAddPricesPriceNewPlanBpsPriceModelType string

const (
	PlanVersionNewParamsAddPricesPriceNewPlanBpsPriceModelTypeBps PlanVersionNewParamsAddPricesPriceNewPlanBpsPriceModelType = "bps"
)

func (r PlanVersionNewParamsAddPricesPriceNewPlanBpsPriceModelType) IsKnown() bool {
	switch r {
	case PlanVersionNewParamsAddPricesPriceNewPlanBpsPriceModelTypeBps:
		return true
	}
	return false
}

// For custom cadence: specifies the duration of the billing period in days or
// months.
type PlanVersionNewParamsAddPricesPriceNewPlanBpsPriceBillingCycleConfiguration struct {
	// The duration of the billing period.
	Duration param.Field[int64] `json:"duration,required"`
	// The unit of billing period duration.
	DurationUnit param.Field[PlanVersionNewParamsAddPricesPriceNewPlanBpsPriceBillingCycleConfigurationDurationUnit] `json:"duration_unit,required"`
}

func (r PlanVersionNewParamsAddPricesPriceNewPlanBpsPriceBillingCycleConfiguration) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// The unit of billing period duration.
type PlanVersionNewParamsAddPricesPriceNewPlanBpsPriceBillingCycleConfigurationDurationUnit string

const (
	PlanVersionNewParamsAddPricesPriceNewPlanBpsPriceBillingCycleConfigurationDurationUnitDay   PlanVersionNewParamsAddPricesPriceNewPlanBpsPriceBillingCycleConfigurationDurationUnit = "day"
	PlanVersionNewParamsAddPricesPriceNewPlanBpsPriceBillingCycleConfigurationDurationUnitMonth PlanVersionNewParamsAddPricesPriceNewPlanBpsPriceBillingCycleConfigurationDurationUnit = "month"
)

func (r PlanVersionNewParamsAddPricesPriceNewPlanBpsPriceBillingCycleConfigurationDurationUnit) IsKnown() bool {
	switch r {
	case PlanVersionNewParamsAddPricesPriceNewPlanBpsPriceBillingCycleConfigurationDurationUnitDay, PlanVersionNewParamsAddPricesPriceNewPlanBpsPriceBillingCycleConfigurationDurationUnitMonth:
		return true
	}
	return false
}

// For dimensional price: specifies a price group and dimension values
type PlanVersionNewParamsAddPricesPriceNewPlanBpsPriceDimensionalPriceConfiguration struct {
	// The list of dimension values matching (in order) the dimensions of the price
	// group
	DimensionValues param.Field[[]string] `json:"dimension_values,required"`
	// The id of the dimensional price group to include this price in
	DimensionalPriceGroupID param.Field[string] `json:"dimensional_price_group_id"`
	// The external id of the dimensional price group to include this price in
	ExternalDimensionalPriceGroupID param.Field[string] `json:"external_dimensional_price_group_id"`
}

func (r PlanVersionNewParamsAddPricesPriceNewPlanBpsPriceDimensionalPriceConfiguration) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Within each billing cycle, specifies the cadence at which invoices are produced.
// If unspecified, a single invoice is produced per billing cycle.
type PlanVersionNewParamsAddPricesPriceNewPlanBpsPriceInvoicingCycleConfiguration struct {
	// The duration of the billing period.
	Duration param.Field[int64] `json:"duration,required"`
	// The unit of billing period duration.
	DurationUnit param.Field[PlanVersionNewParamsAddPricesPriceNewPlanBpsPriceInvoicingCycleConfigurationDurationUnit] `json:"duration_unit,required"`
}

func (r PlanVersionNewParamsAddPricesPriceNewPlanBpsPriceInvoicingCycleConfiguration) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// The unit of billing period duration.
type PlanVersionNewParamsAddPricesPriceNewPlanBpsPriceInvoicingCycleConfigurationDurationUnit string

const (
	PlanVersionNewParamsAddPricesPriceNewPlanBpsPriceInvoicingCycleConfigurationDurationUnitDay   PlanVersionNewParamsAddPricesPriceNewPlanBpsPriceInvoicingCycleConfigurationDurationUnit = "day"
	PlanVersionNewParamsAddPricesPriceNewPlanBpsPriceInvoicingCycleConfigurationDurationUnitMonth PlanVersionNewParamsAddPricesPriceNewPlanBpsPriceInvoicingCycleConfigurationDurationUnit = "month"
)

func (r PlanVersionNewParamsAddPricesPriceNewPlanBpsPriceInvoicingCycleConfigurationDurationUnit) IsKnown() bool {
	switch r {
	case PlanVersionNewParamsAddPricesPriceNewPlanBpsPriceInvoicingCycleConfigurationDurationUnitDay, PlanVersionNewParamsAddPricesPriceNewPlanBpsPriceInvoicingCycleConfigurationDurationUnitMonth:
		return true
	}
	return false
}

type PlanVersionNewParamsAddPricesPriceNewPlanBulkBpsPrice struct {
	BulkBpsConfig param.Field[PlanVersionNewParamsAddPricesPriceNewPlanBulkBpsPriceBulkBpsConfig] `json:"bulk_bps_config,required"`
	// The cadence to bill for this price on.
	Cadence param.Field[PlanVersionNewParamsAddPricesPriceNewPlanBulkBpsPriceCadence] `json:"cadence,required"`
	// The id of the item the price will be associated with.
	ItemID    param.Field[string]                                                         `json:"item_id,required"`
	ModelType param.Field[PlanVersionNewParamsAddPricesPriceNewPlanBulkBpsPriceModelType] `json:"model_type,required"`
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
	BillingCycleConfiguration param.Field[PlanVersionNewParamsAddPricesPriceNewPlanBulkBpsPriceBillingCycleConfiguration] `json:"billing_cycle_configuration"`
	// The per unit conversion rate of the price currency to the invoicing currency.
	ConversionRate param.Field[float64] `json:"conversion_rate"`
	// An ISO 4217 currency string, or custom pricing unit identifier, in which this
	// price is billed.
	Currency param.Field[string] `json:"currency"`
	// For dimensional price: specifies a price group and dimension values
	DimensionalPriceConfiguration param.Field[PlanVersionNewParamsAddPricesPriceNewPlanBulkBpsPriceDimensionalPriceConfiguration] `json:"dimensional_price_configuration"`
	// An alias for the price.
	ExternalPriceID param.Field[string] `json:"external_price_id"`
	// If the Price represents a fixed cost, this represents the quantity of units
	// applied.
	FixedPriceQuantity param.Field[float64] `json:"fixed_price_quantity"`
	// The property used to group this price on an invoice
	InvoiceGroupingKey param.Field[string] `json:"invoice_grouping_key"`
	// Within each billing cycle, specifies the cadence at which invoices are produced.
	// If unspecified, a single invoice is produced per billing cycle.
	InvoicingCycleConfiguration param.Field[PlanVersionNewParamsAddPricesPriceNewPlanBulkBpsPriceInvoicingCycleConfiguration] `json:"invoicing_cycle_configuration"`
	// User-specified key/value pairs for the resource. Individual keys can be removed
	// by setting the value to `null`, and the entire metadata mapping can be cleared
	// by setting `metadata` to `null`.
	Metadata param.Field[map[string]string] `json:"metadata"`
}

func (r PlanVersionNewParamsAddPricesPriceNewPlanBulkBpsPrice) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r PlanVersionNewParamsAddPricesPriceNewPlanBulkBpsPrice) implementsPlanVersionNewParamsAddPricesPriceUnion() {
}

type PlanVersionNewParamsAddPricesPriceNewPlanBulkBpsPriceBulkBpsConfig struct {
	// Tiers for a bulk BPS pricing model where all usage is aggregated to a single
	// tier based on total volume
	Tiers param.Field[[]PlanVersionNewParamsAddPricesPriceNewPlanBulkBpsPriceBulkBpsConfigTier] `json:"tiers,required"`
}

func (r PlanVersionNewParamsAddPricesPriceNewPlanBulkBpsPriceBulkBpsConfig) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type PlanVersionNewParamsAddPricesPriceNewPlanBulkBpsPriceBulkBpsConfigTier struct {
	// Basis points to rate on
	Bps param.Field[float64] `json:"bps,required"`
	// Upper bound for tier
	MaximumAmount param.Field[string] `json:"maximum_amount"`
	// The maximum amount to charge for any one event
	PerUnitMaximum param.Field[string] `json:"per_unit_maximum"`
}

func (r PlanVersionNewParamsAddPricesPriceNewPlanBulkBpsPriceBulkBpsConfigTier) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// The cadence to bill for this price on.
type PlanVersionNewParamsAddPricesPriceNewPlanBulkBpsPriceCadence string

const (
	PlanVersionNewParamsAddPricesPriceNewPlanBulkBpsPriceCadenceAnnual     PlanVersionNewParamsAddPricesPriceNewPlanBulkBpsPriceCadence = "annual"
	PlanVersionNewParamsAddPricesPriceNewPlanBulkBpsPriceCadenceSemiAnnual PlanVersionNewParamsAddPricesPriceNewPlanBulkBpsPriceCadence = "semi_annual"
	PlanVersionNewParamsAddPricesPriceNewPlanBulkBpsPriceCadenceMonthly    PlanVersionNewParamsAddPricesPriceNewPlanBulkBpsPriceCadence = "monthly"
	PlanVersionNewParamsAddPricesPriceNewPlanBulkBpsPriceCadenceQuarterly  PlanVersionNewParamsAddPricesPriceNewPlanBulkBpsPriceCadence = "quarterly"
	PlanVersionNewParamsAddPricesPriceNewPlanBulkBpsPriceCadenceOneTime    PlanVersionNewParamsAddPricesPriceNewPlanBulkBpsPriceCadence = "one_time"
	PlanVersionNewParamsAddPricesPriceNewPlanBulkBpsPriceCadenceCustom     PlanVersionNewParamsAddPricesPriceNewPlanBulkBpsPriceCadence = "custom"
)

func (r PlanVersionNewParamsAddPricesPriceNewPlanBulkBpsPriceCadence) IsKnown() bool {
	switch r {
	case PlanVersionNewParamsAddPricesPriceNewPlanBulkBpsPriceCadenceAnnual, PlanVersionNewParamsAddPricesPriceNewPlanBulkBpsPriceCadenceSemiAnnual, PlanVersionNewParamsAddPricesPriceNewPlanBulkBpsPriceCadenceMonthly, PlanVersionNewParamsAddPricesPriceNewPlanBulkBpsPriceCadenceQuarterly, PlanVersionNewParamsAddPricesPriceNewPlanBulkBpsPriceCadenceOneTime, PlanVersionNewParamsAddPricesPriceNewPlanBulkBpsPriceCadenceCustom:
		return true
	}
	return false
}

type PlanVersionNewParamsAddPricesPriceNewPlanBulkBpsPriceModelType string

const (
	PlanVersionNewParamsAddPricesPriceNewPlanBulkBpsPriceModelTypeBulkBps PlanVersionNewParamsAddPricesPriceNewPlanBulkBpsPriceModelType = "bulk_bps"
)

func (r PlanVersionNewParamsAddPricesPriceNewPlanBulkBpsPriceModelType) IsKnown() bool {
	switch r {
	case PlanVersionNewParamsAddPricesPriceNewPlanBulkBpsPriceModelTypeBulkBps:
		return true
	}
	return false
}

// For custom cadence: specifies the duration of the billing period in days or
// months.
type PlanVersionNewParamsAddPricesPriceNewPlanBulkBpsPriceBillingCycleConfiguration struct {
	// The duration of the billing period.
	Duration param.Field[int64] `json:"duration,required"`
	// The unit of billing period duration.
	DurationUnit param.Field[PlanVersionNewParamsAddPricesPriceNewPlanBulkBpsPriceBillingCycleConfigurationDurationUnit] `json:"duration_unit,required"`
}

func (r PlanVersionNewParamsAddPricesPriceNewPlanBulkBpsPriceBillingCycleConfiguration) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// The unit of billing period duration.
type PlanVersionNewParamsAddPricesPriceNewPlanBulkBpsPriceBillingCycleConfigurationDurationUnit string

const (
	PlanVersionNewParamsAddPricesPriceNewPlanBulkBpsPriceBillingCycleConfigurationDurationUnitDay   PlanVersionNewParamsAddPricesPriceNewPlanBulkBpsPriceBillingCycleConfigurationDurationUnit = "day"
	PlanVersionNewParamsAddPricesPriceNewPlanBulkBpsPriceBillingCycleConfigurationDurationUnitMonth PlanVersionNewParamsAddPricesPriceNewPlanBulkBpsPriceBillingCycleConfigurationDurationUnit = "month"
)

func (r PlanVersionNewParamsAddPricesPriceNewPlanBulkBpsPriceBillingCycleConfigurationDurationUnit) IsKnown() bool {
	switch r {
	case PlanVersionNewParamsAddPricesPriceNewPlanBulkBpsPriceBillingCycleConfigurationDurationUnitDay, PlanVersionNewParamsAddPricesPriceNewPlanBulkBpsPriceBillingCycleConfigurationDurationUnitMonth:
		return true
	}
	return false
}

// For dimensional price: specifies a price group and dimension values
type PlanVersionNewParamsAddPricesPriceNewPlanBulkBpsPriceDimensionalPriceConfiguration struct {
	// The list of dimension values matching (in order) the dimensions of the price
	// group
	DimensionValues param.Field[[]string] `json:"dimension_values,required"`
	// The id of the dimensional price group to include this price in
	DimensionalPriceGroupID param.Field[string] `json:"dimensional_price_group_id"`
	// The external id of the dimensional price group to include this price in
	ExternalDimensionalPriceGroupID param.Field[string] `json:"external_dimensional_price_group_id"`
}

func (r PlanVersionNewParamsAddPricesPriceNewPlanBulkBpsPriceDimensionalPriceConfiguration) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Within each billing cycle, specifies the cadence at which invoices are produced.
// If unspecified, a single invoice is produced per billing cycle.
type PlanVersionNewParamsAddPricesPriceNewPlanBulkBpsPriceInvoicingCycleConfiguration struct {
	// The duration of the billing period.
	Duration param.Field[int64] `json:"duration,required"`
	// The unit of billing period duration.
	DurationUnit param.Field[PlanVersionNewParamsAddPricesPriceNewPlanBulkBpsPriceInvoicingCycleConfigurationDurationUnit] `json:"duration_unit,required"`
}

func (r PlanVersionNewParamsAddPricesPriceNewPlanBulkBpsPriceInvoicingCycleConfiguration) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// The unit of billing period duration.
type PlanVersionNewParamsAddPricesPriceNewPlanBulkBpsPriceInvoicingCycleConfigurationDurationUnit string

const (
	PlanVersionNewParamsAddPricesPriceNewPlanBulkBpsPriceInvoicingCycleConfigurationDurationUnitDay   PlanVersionNewParamsAddPricesPriceNewPlanBulkBpsPriceInvoicingCycleConfigurationDurationUnit = "day"
	PlanVersionNewParamsAddPricesPriceNewPlanBulkBpsPriceInvoicingCycleConfigurationDurationUnitMonth PlanVersionNewParamsAddPricesPriceNewPlanBulkBpsPriceInvoicingCycleConfigurationDurationUnit = "month"
)

func (r PlanVersionNewParamsAddPricesPriceNewPlanBulkBpsPriceInvoicingCycleConfigurationDurationUnit) IsKnown() bool {
	switch r {
	case PlanVersionNewParamsAddPricesPriceNewPlanBulkBpsPriceInvoicingCycleConfigurationDurationUnitDay, PlanVersionNewParamsAddPricesPriceNewPlanBulkBpsPriceInvoicingCycleConfigurationDurationUnitMonth:
		return true
	}
	return false
}

type PlanVersionNewParamsAddPricesPriceNewPlanBulkPrice struct {
	BulkConfig param.Field[PlanVersionNewParamsAddPricesPriceNewPlanBulkPriceBulkConfig] `json:"bulk_config,required"`
	// The cadence to bill for this price on.
	Cadence param.Field[PlanVersionNewParamsAddPricesPriceNewPlanBulkPriceCadence] `json:"cadence,required"`
	// The id of the item the price will be associated with.
	ItemID    param.Field[string]                                                      `json:"item_id,required"`
	ModelType param.Field[PlanVersionNewParamsAddPricesPriceNewPlanBulkPriceModelType] `json:"model_type,required"`
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
	BillingCycleConfiguration param.Field[PlanVersionNewParamsAddPricesPriceNewPlanBulkPriceBillingCycleConfiguration] `json:"billing_cycle_configuration"`
	// The per unit conversion rate of the price currency to the invoicing currency.
	ConversionRate param.Field[float64] `json:"conversion_rate"`
	// An ISO 4217 currency string, or custom pricing unit identifier, in which this
	// price is billed.
	Currency param.Field[string] `json:"currency"`
	// For dimensional price: specifies a price group and dimension values
	DimensionalPriceConfiguration param.Field[PlanVersionNewParamsAddPricesPriceNewPlanBulkPriceDimensionalPriceConfiguration] `json:"dimensional_price_configuration"`
	// An alias for the price.
	ExternalPriceID param.Field[string] `json:"external_price_id"`
	// If the Price represents a fixed cost, this represents the quantity of units
	// applied.
	FixedPriceQuantity param.Field[float64] `json:"fixed_price_quantity"`
	// The property used to group this price on an invoice
	InvoiceGroupingKey param.Field[string] `json:"invoice_grouping_key"`
	// Within each billing cycle, specifies the cadence at which invoices are produced.
	// If unspecified, a single invoice is produced per billing cycle.
	InvoicingCycleConfiguration param.Field[PlanVersionNewParamsAddPricesPriceNewPlanBulkPriceInvoicingCycleConfiguration] `json:"invoicing_cycle_configuration"`
	// User-specified key/value pairs for the resource. Individual keys can be removed
	// by setting the value to `null`, and the entire metadata mapping can be cleared
	// by setting `metadata` to `null`.
	Metadata param.Field[map[string]string] `json:"metadata"`
}

func (r PlanVersionNewParamsAddPricesPriceNewPlanBulkPrice) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r PlanVersionNewParamsAddPricesPriceNewPlanBulkPrice) implementsPlanVersionNewParamsAddPricesPriceUnion() {
}

type PlanVersionNewParamsAddPricesPriceNewPlanBulkPriceBulkConfig struct {
	// Bulk tiers for rating based on total usage volume
	Tiers param.Field[[]PlanVersionNewParamsAddPricesPriceNewPlanBulkPriceBulkConfigTier] `json:"tiers,required"`
}

func (r PlanVersionNewParamsAddPricesPriceNewPlanBulkPriceBulkConfig) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type PlanVersionNewParamsAddPricesPriceNewPlanBulkPriceBulkConfigTier struct {
	// Amount per unit
	UnitAmount param.Field[string] `json:"unit_amount,required"`
	// Upper bound for this tier
	MaximumUnits param.Field[float64] `json:"maximum_units"`
}

func (r PlanVersionNewParamsAddPricesPriceNewPlanBulkPriceBulkConfigTier) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// The cadence to bill for this price on.
type PlanVersionNewParamsAddPricesPriceNewPlanBulkPriceCadence string

const (
	PlanVersionNewParamsAddPricesPriceNewPlanBulkPriceCadenceAnnual     PlanVersionNewParamsAddPricesPriceNewPlanBulkPriceCadence = "annual"
	PlanVersionNewParamsAddPricesPriceNewPlanBulkPriceCadenceSemiAnnual PlanVersionNewParamsAddPricesPriceNewPlanBulkPriceCadence = "semi_annual"
	PlanVersionNewParamsAddPricesPriceNewPlanBulkPriceCadenceMonthly    PlanVersionNewParamsAddPricesPriceNewPlanBulkPriceCadence = "monthly"
	PlanVersionNewParamsAddPricesPriceNewPlanBulkPriceCadenceQuarterly  PlanVersionNewParamsAddPricesPriceNewPlanBulkPriceCadence = "quarterly"
	PlanVersionNewParamsAddPricesPriceNewPlanBulkPriceCadenceOneTime    PlanVersionNewParamsAddPricesPriceNewPlanBulkPriceCadence = "one_time"
	PlanVersionNewParamsAddPricesPriceNewPlanBulkPriceCadenceCustom     PlanVersionNewParamsAddPricesPriceNewPlanBulkPriceCadence = "custom"
)

func (r PlanVersionNewParamsAddPricesPriceNewPlanBulkPriceCadence) IsKnown() bool {
	switch r {
	case PlanVersionNewParamsAddPricesPriceNewPlanBulkPriceCadenceAnnual, PlanVersionNewParamsAddPricesPriceNewPlanBulkPriceCadenceSemiAnnual, PlanVersionNewParamsAddPricesPriceNewPlanBulkPriceCadenceMonthly, PlanVersionNewParamsAddPricesPriceNewPlanBulkPriceCadenceQuarterly, PlanVersionNewParamsAddPricesPriceNewPlanBulkPriceCadenceOneTime, PlanVersionNewParamsAddPricesPriceNewPlanBulkPriceCadenceCustom:
		return true
	}
	return false
}

type PlanVersionNewParamsAddPricesPriceNewPlanBulkPriceModelType string

const (
	PlanVersionNewParamsAddPricesPriceNewPlanBulkPriceModelTypeBulk PlanVersionNewParamsAddPricesPriceNewPlanBulkPriceModelType = "bulk"
)

func (r PlanVersionNewParamsAddPricesPriceNewPlanBulkPriceModelType) IsKnown() bool {
	switch r {
	case PlanVersionNewParamsAddPricesPriceNewPlanBulkPriceModelTypeBulk:
		return true
	}
	return false
}

// For custom cadence: specifies the duration of the billing period in days or
// months.
type PlanVersionNewParamsAddPricesPriceNewPlanBulkPriceBillingCycleConfiguration struct {
	// The duration of the billing period.
	Duration param.Field[int64] `json:"duration,required"`
	// The unit of billing period duration.
	DurationUnit param.Field[PlanVersionNewParamsAddPricesPriceNewPlanBulkPriceBillingCycleConfigurationDurationUnit] `json:"duration_unit,required"`
}

func (r PlanVersionNewParamsAddPricesPriceNewPlanBulkPriceBillingCycleConfiguration) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// The unit of billing period duration.
type PlanVersionNewParamsAddPricesPriceNewPlanBulkPriceBillingCycleConfigurationDurationUnit string

const (
	PlanVersionNewParamsAddPricesPriceNewPlanBulkPriceBillingCycleConfigurationDurationUnitDay   PlanVersionNewParamsAddPricesPriceNewPlanBulkPriceBillingCycleConfigurationDurationUnit = "day"
	PlanVersionNewParamsAddPricesPriceNewPlanBulkPriceBillingCycleConfigurationDurationUnitMonth PlanVersionNewParamsAddPricesPriceNewPlanBulkPriceBillingCycleConfigurationDurationUnit = "month"
)

func (r PlanVersionNewParamsAddPricesPriceNewPlanBulkPriceBillingCycleConfigurationDurationUnit) IsKnown() bool {
	switch r {
	case PlanVersionNewParamsAddPricesPriceNewPlanBulkPriceBillingCycleConfigurationDurationUnitDay, PlanVersionNewParamsAddPricesPriceNewPlanBulkPriceBillingCycleConfigurationDurationUnitMonth:
		return true
	}
	return false
}

// For dimensional price: specifies a price group and dimension values
type PlanVersionNewParamsAddPricesPriceNewPlanBulkPriceDimensionalPriceConfiguration struct {
	// The list of dimension values matching (in order) the dimensions of the price
	// group
	DimensionValues param.Field[[]string] `json:"dimension_values,required"`
	// The id of the dimensional price group to include this price in
	DimensionalPriceGroupID param.Field[string] `json:"dimensional_price_group_id"`
	// The external id of the dimensional price group to include this price in
	ExternalDimensionalPriceGroupID param.Field[string] `json:"external_dimensional_price_group_id"`
}

func (r PlanVersionNewParamsAddPricesPriceNewPlanBulkPriceDimensionalPriceConfiguration) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Within each billing cycle, specifies the cadence at which invoices are produced.
// If unspecified, a single invoice is produced per billing cycle.
type PlanVersionNewParamsAddPricesPriceNewPlanBulkPriceInvoicingCycleConfiguration struct {
	// The duration of the billing period.
	Duration param.Field[int64] `json:"duration,required"`
	// The unit of billing period duration.
	DurationUnit param.Field[PlanVersionNewParamsAddPricesPriceNewPlanBulkPriceInvoicingCycleConfigurationDurationUnit] `json:"duration_unit,required"`
}

func (r PlanVersionNewParamsAddPricesPriceNewPlanBulkPriceInvoicingCycleConfiguration) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// The unit of billing period duration.
type PlanVersionNewParamsAddPricesPriceNewPlanBulkPriceInvoicingCycleConfigurationDurationUnit string

const (
	PlanVersionNewParamsAddPricesPriceNewPlanBulkPriceInvoicingCycleConfigurationDurationUnitDay   PlanVersionNewParamsAddPricesPriceNewPlanBulkPriceInvoicingCycleConfigurationDurationUnit = "day"
	PlanVersionNewParamsAddPricesPriceNewPlanBulkPriceInvoicingCycleConfigurationDurationUnitMonth PlanVersionNewParamsAddPricesPriceNewPlanBulkPriceInvoicingCycleConfigurationDurationUnit = "month"
)

func (r PlanVersionNewParamsAddPricesPriceNewPlanBulkPriceInvoicingCycleConfigurationDurationUnit) IsKnown() bool {
	switch r {
	case PlanVersionNewParamsAddPricesPriceNewPlanBulkPriceInvoicingCycleConfigurationDurationUnitDay, PlanVersionNewParamsAddPricesPriceNewPlanBulkPriceInvoicingCycleConfigurationDurationUnitMonth:
		return true
	}
	return false
}

type PlanVersionNewParamsAddPricesPriceNewPlanThresholdTotalAmountPrice struct {
	// The cadence to bill for this price on.
	Cadence param.Field[PlanVersionNewParamsAddPricesPriceNewPlanThresholdTotalAmountPriceCadence] `json:"cadence,required"`
	// The id of the item the price will be associated with.
	ItemID    param.Field[string]                                                                      `json:"item_id,required"`
	ModelType param.Field[PlanVersionNewParamsAddPricesPriceNewPlanThresholdTotalAmountPriceModelType] `json:"model_type,required"`
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
	BillingCycleConfiguration param.Field[PlanVersionNewParamsAddPricesPriceNewPlanThresholdTotalAmountPriceBillingCycleConfiguration] `json:"billing_cycle_configuration"`
	// The per unit conversion rate of the price currency to the invoicing currency.
	ConversionRate param.Field[float64] `json:"conversion_rate"`
	// An ISO 4217 currency string, or custom pricing unit identifier, in which this
	// price is billed.
	Currency param.Field[string] `json:"currency"`
	// For dimensional price: specifies a price group and dimension values
	DimensionalPriceConfiguration param.Field[PlanVersionNewParamsAddPricesPriceNewPlanThresholdTotalAmountPriceDimensionalPriceConfiguration] `json:"dimensional_price_configuration"`
	// An alias for the price.
	ExternalPriceID param.Field[string] `json:"external_price_id"`
	// If the Price represents a fixed cost, this represents the quantity of units
	// applied.
	FixedPriceQuantity param.Field[float64] `json:"fixed_price_quantity"`
	// The property used to group this price on an invoice
	InvoiceGroupingKey param.Field[string] `json:"invoice_grouping_key"`
	// Within each billing cycle, specifies the cadence at which invoices are produced.
	// If unspecified, a single invoice is produced per billing cycle.
	InvoicingCycleConfiguration param.Field[PlanVersionNewParamsAddPricesPriceNewPlanThresholdTotalAmountPriceInvoicingCycleConfiguration] `json:"invoicing_cycle_configuration"`
	// User-specified key/value pairs for the resource. Individual keys can be removed
	// by setting the value to `null`, and the entire metadata mapping can be cleared
	// by setting `metadata` to `null`.
	Metadata param.Field[map[string]string] `json:"metadata"`
}

func (r PlanVersionNewParamsAddPricesPriceNewPlanThresholdTotalAmountPrice) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r PlanVersionNewParamsAddPricesPriceNewPlanThresholdTotalAmountPrice) implementsPlanVersionNewParamsAddPricesPriceUnion() {
}

// The cadence to bill for this price on.
type PlanVersionNewParamsAddPricesPriceNewPlanThresholdTotalAmountPriceCadence string

const (
	PlanVersionNewParamsAddPricesPriceNewPlanThresholdTotalAmountPriceCadenceAnnual     PlanVersionNewParamsAddPricesPriceNewPlanThresholdTotalAmountPriceCadence = "annual"
	PlanVersionNewParamsAddPricesPriceNewPlanThresholdTotalAmountPriceCadenceSemiAnnual PlanVersionNewParamsAddPricesPriceNewPlanThresholdTotalAmountPriceCadence = "semi_annual"
	PlanVersionNewParamsAddPricesPriceNewPlanThresholdTotalAmountPriceCadenceMonthly    PlanVersionNewParamsAddPricesPriceNewPlanThresholdTotalAmountPriceCadence = "monthly"
	PlanVersionNewParamsAddPricesPriceNewPlanThresholdTotalAmountPriceCadenceQuarterly  PlanVersionNewParamsAddPricesPriceNewPlanThresholdTotalAmountPriceCadence = "quarterly"
	PlanVersionNewParamsAddPricesPriceNewPlanThresholdTotalAmountPriceCadenceOneTime    PlanVersionNewParamsAddPricesPriceNewPlanThresholdTotalAmountPriceCadence = "one_time"
	PlanVersionNewParamsAddPricesPriceNewPlanThresholdTotalAmountPriceCadenceCustom     PlanVersionNewParamsAddPricesPriceNewPlanThresholdTotalAmountPriceCadence = "custom"
)

func (r PlanVersionNewParamsAddPricesPriceNewPlanThresholdTotalAmountPriceCadence) IsKnown() bool {
	switch r {
	case PlanVersionNewParamsAddPricesPriceNewPlanThresholdTotalAmountPriceCadenceAnnual, PlanVersionNewParamsAddPricesPriceNewPlanThresholdTotalAmountPriceCadenceSemiAnnual, PlanVersionNewParamsAddPricesPriceNewPlanThresholdTotalAmountPriceCadenceMonthly, PlanVersionNewParamsAddPricesPriceNewPlanThresholdTotalAmountPriceCadenceQuarterly, PlanVersionNewParamsAddPricesPriceNewPlanThresholdTotalAmountPriceCadenceOneTime, PlanVersionNewParamsAddPricesPriceNewPlanThresholdTotalAmountPriceCadenceCustom:
		return true
	}
	return false
}

type PlanVersionNewParamsAddPricesPriceNewPlanThresholdTotalAmountPriceModelType string

const (
	PlanVersionNewParamsAddPricesPriceNewPlanThresholdTotalAmountPriceModelTypeThresholdTotalAmount PlanVersionNewParamsAddPricesPriceNewPlanThresholdTotalAmountPriceModelType = "threshold_total_amount"
)

func (r PlanVersionNewParamsAddPricesPriceNewPlanThresholdTotalAmountPriceModelType) IsKnown() bool {
	switch r {
	case PlanVersionNewParamsAddPricesPriceNewPlanThresholdTotalAmountPriceModelTypeThresholdTotalAmount:
		return true
	}
	return false
}

// For custom cadence: specifies the duration of the billing period in days or
// months.
type PlanVersionNewParamsAddPricesPriceNewPlanThresholdTotalAmountPriceBillingCycleConfiguration struct {
	// The duration of the billing period.
	Duration param.Field[int64] `json:"duration,required"`
	// The unit of billing period duration.
	DurationUnit param.Field[PlanVersionNewParamsAddPricesPriceNewPlanThresholdTotalAmountPriceBillingCycleConfigurationDurationUnit] `json:"duration_unit,required"`
}

func (r PlanVersionNewParamsAddPricesPriceNewPlanThresholdTotalAmountPriceBillingCycleConfiguration) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// The unit of billing period duration.
type PlanVersionNewParamsAddPricesPriceNewPlanThresholdTotalAmountPriceBillingCycleConfigurationDurationUnit string

const (
	PlanVersionNewParamsAddPricesPriceNewPlanThresholdTotalAmountPriceBillingCycleConfigurationDurationUnitDay   PlanVersionNewParamsAddPricesPriceNewPlanThresholdTotalAmountPriceBillingCycleConfigurationDurationUnit = "day"
	PlanVersionNewParamsAddPricesPriceNewPlanThresholdTotalAmountPriceBillingCycleConfigurationDurationUnitMonth PlanVersionNewParamsAddPricesPriceNewPlanThresholdTotalAmountPriceBillingCycleConfigurationDurationUnit = "month"
)

func (r PlanVersionNewParamsAddPricesPriceNewPlanThresholdTotalAmountPriceBillingCycleConfigurationDurationUnit) IsKnown() bool {
	switch r {
	case PlanVersionNewParamsAddPricesPriceNewPlanThresholdTotalAmountPriceBillingCycleConfigurationDurationUnitDay, PlanVersionNewParamsAddPricesPriceNewPlanThresholdTotalAmountPriceBillingCycleConfigurationDurationUnitMonth:
		return true
	}
	return false
}

// For dimensional price: specifies a price group and dimension values
type PlanVersionNewParamsAddPricesPriceNewPlanThresholdTotalAmountPriceDimensionalPriceConfiguration struct {
	// The list of dimension values matching (in order) the dimensions of the price
	// group
	DimensionValues param.Field[[]string] `json:"dimension_values,required"`
	// The id of the dimensional price group to include this price in
	DimensionalPriceGroupID param.Field[string] `json:"dimensional_price_group_id"`
	// The external id of the dimensional price group to include this price in
	ExternalDimensionalPriceGroupID param.Field[string] `json:"external_dimensional_price_group_id"`
}

func (r PlanVersionNewParamsAddPricesPriceNewPlanThresholdTotalAmountPriceDimensionalPriceConfiguration) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Within each billing cycle, specifies the cadence at which invoices are produced.
// If unspecified, a single invoice is produced per billing cycle.
type PlanVersionNewParamsAddPricesPriceNewPlanThresholdTotalAmountPriceInvoicingCycleConfiguration struct {
	// The duration of the billing period.
	Duration param.Field[int64] `json:"duration,required"`
	// The unit of billing period duration.
	DurationUnit param.Field[PlanVersionNewParamsAddPricesPriceNewPlanThresholdTotalAmountPriceInvoicingCycleConfigurationDurationUnit] `json:"duration_unit,required"`
}

func (r PlanVersionNewParamsAddPricesPriceNewPlanThresholdTotalAmountPriceInvoicingCycleConfiguration) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// The unit of billing period duration.
type PlanVersionNewParamsAddPricesPriceNewPlanThresholdTotalAmountPriceInvoicingCycleConfigurationDurationUnit string

const (
	PlanVersionNewParamsAddPricesPriceNewPlanThresholdTotalAmountPriceInvoicingCycleConfigurationDurationUnitDay   PlanVersionNewParamsAddPricesPriceNewPlanThresholdTotalAmountPriceInvoicingCycleConfigurationDurationUnit = "day"
	PlanVersionNewParamsAddPricesPriceNewPlanThresholdTotalAmountPriceInvoicingCycleConfigurationDurationUnitMonth PlanVersionNewParamsAddPricesPriceNewPlanThresholdTotalAmountPriceInvoicingCycleConfigurationDurationUnit = "month"
)

func (r PlanVersionNewParamsAddPricesPriceNewPlanThresholdTotalAmountPriceInvoicingCycleConfigurationDurationUnit) IsKnown() bool {
	switch r {
	case PlanVersionNewParamsAddPricesPriceNewPlanThresholdTotalAmountPriceInvoicingCycleConfigurationDurationUnitDay, PlanVersionNewParamsAddPricesPriceNewPlanThresholdTotalAmountPriceInvoicingCycleConfigurationDurationUnitMonth:
		return true
	}
	return false
}

type PlanVersionNewParamsAddPricesPriceNewPlanTieredPackagePrice struct {
	// The cadence to bill for this price on.
	Cadence param.Field[PlanVersionNewParamsAddPricesPriceNewPlanTieredPackagePriceCadence] `json:"cadence,required"`
	// The id of the item the price will be associated with.
	ItemID    param.Field[string]                                                               `json:"item_id,required"`
	ModelType param.Field[PlanVersionNewParamsAddPricesPriceNewPlanTieredPackagePriceModelType] `json:"model_type,required"`
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
	BillingCycleConfiguration param.Field[PlanVersionNewParamsAddPricesPriceNewPlanTieredPackagePriceBillingCycleConfiguration] `json:"billing_cycle_configuration"`
	// The per unit conversion rate of the price currency to the invoicing currency.
	ConversionRate param.Field[float64] `json:"conversion_rate"`
	// An ISO 4217 currency string, or custom pricing unit identifier, in which this
	// price is billed.
	Currency param.Field[string] `json:"currency"`
	// For dimensional price: specifies a price group and dimension values
	DimensionalPriceConfiguration param.Field[PlanVersionNewParamsAddPricesPriceNewPlanTieredPackagePriceDimensionalPriceConfiguration] `json:"dimensional_price_configuration"`
	// An alias for the price.
	ExternalPriceID param.Field[string] `json:"external_price_id"`
	// If the Price represents a fixed cost, this represents the quantity of units
	// applied.
	FixedPriceQuantity param.Field[float64] `json:"fixed_price_quantity"`
	// The property used to group this price on an invoice
	InvoiceGroupingKey param.Field[string] `json:"invoice_grouping_key"`
	// Within each billing cycle, specifies the cadence at which invoices are produced.
	// If unspecified, a single invoice is produced per billing cycle.
	InvoicingCycleConfiguration param.Field[PlanVersionNewParamsAddPricesPriceNewPlanTieredPackagePriceInvoicingCycleConfiguration] `json:"invoicing_cycle_configuration"`
	// User-specified key/value pairs for the resource. Individual keys can be removed
	// by setting the value to `null`, and the entire metadata mapping can be cleared
	// by setting `metadata` to `null`.
	Metadata param.Field[map[string]string] `json:"metadata"`
}

func (r PlanVersionNewParamsAddPricesPriceNewPlanTieredPackagePrice) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r PlanVersionNewParamsAddPricesPriceNewPlanTieredPackagePrice) implementsPlanVersionNewParamsAddPricesPriceUnion() {
}

// The cadence to bill for this price on.
type PlanVersionNewParamsAddPricesPriceNewPlanTieredPackagePriceCadence string

const (
	PlanVersionNewParamsAddPricesPriceNewPlanTieredPackagePriceCadenceAnnual     PlanVersionNewParamsAddPricesPriceNewPlanTieredPackagePriceCadence = "annual"
	PlanVersionNewParamsAddPricesPriceNewPlanTieredPackagePriceCadenceSemiAnnual PlanVersionNewParamsAddPricesPriceNewPlanTieredPackagePriceCadence = "semi_annual"
	PlanVersionNewParamsAddPricesPriceNewPlanTieredPackagePriceCadenceMonthly    PlanVersionNewParamsAddPricesPriceNewPlanTieredPackagePriceCadence = "monthly"
	PlanVersionNewParamsAddPricesPriceNewPlanTieredPackagePriceCadenceQuarterly  PlanVersionNewParamsAddPricesPriceNewPlanTieredPackagePriceCadence = "quarterly"
	PlanVersionNewParamsAddPricesPriceNewPlanTieredPackagePriceCadenceOneTime    PlanVersionNewParamsAddPricesPriceNewPlanTieredPackagePriceCadence = "one_time"
	PlanVersionNewParamsAddPricesPriceNewPlanTieredPackagePriceCadenceCustom     PlanVersionNewParamsAddPricesPriceNewPlanTieredPackagePriceCadence = "custom"
)

func (r PlanVersionNewParamsAddPricesPriceNewPlanTieredPackagePriceCadence) IsKnown() bool {
	switch r {
	case PlanVersionNewParamsAddPricesPriceNewPlanTieredPackagePriceCadenceAnnual, PlanVersionNewParamsAddPricesPriceNewPlanTieredPackagePriceCadenceSemiAnnual, PlanVersionNewParamsAddPricesPriceNewPlanTieredPackagePriceCadenceMonthly, PlanVersionNewParamsAddPricesPriceNewPlanTieredPackagePriceCadenceQuarterly, PlanVersionNewParamsAddPricesPriceNewPlanTieredPackagePriceCadenceOneTime, PlanVersionNewParamsAddPricesPriceNewPlanTieredPackagePriceCadenceCustom:
		return true
	}
	return false
}

type PlanVersionNewParamsAddPricesPriceNewPlanTieredPackagePriceModelType string

const (
	PlanVersionNewParamsAddPricesPriceNewPlanTieredPackagePriceModelTypeTieredPackage PlanVersionNewParamsAddPricesPriceNewPlanTieredPackagePriceModelType = "tiered_package"
)

func (r PlanVersionNewParamsAddPricesPriceNewPlanTieredPackagePriceModelType) IsKnown() bool {
	switch r {
	case PlanVersionNewParamsAddPricesPriceNewPlanTieredPackagePriceModelTypeTieredPackage:
		return true
	}
	return false
}

// For custom cadence: specifies the duration of the billing period in days or
// months.
type PlanVersionNewParamsAddPricesPriceNewPlanTieredPackagePriceBillingCycleConfiguration struct {
	// The duration of the billing period.
	Duration param.Field[int64] `json:"duration,required"`
	// The unit of billing period duration.
	DurationUnit param.Field[PlanVersionNewParamsAddPricesPriceNewPlanTieredPackagePriceBillingCycleConfigurationDurationUnit] `json:"duration_unit,required"`
}

func (r PlanVersionNewParamsAddPricesPriceNewPlanTieredPackagePriceBillingCycleConfiguration) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// The unit of billing period duration.
type PlanVersionNewParamsAddPricesPriceNewPlanTieredPackagePriceBillingCycleConfigurationDurationUnit string

const (
	PlanVersionNewParamsAddPricesPriceNewPlanTieredPackagePriceBillingCycleConfigurationDurationUnitDay   PlanVersionNewParamsAddPricesPriceNewPlanTieredPackagePriceBillingCycleConfigurationDurationUnit = "day"
	PlanVersionNewParamsAddPricesPriceNewPlanTieredPackagePriceBillingCycleConfigurationDurationUnitMonth PlanVersionNewParamsAddPricesPriceNewPlanTieredPackagePriceBillingCycleConfigurationDurationUnit = "month"
)

func (r PlanVersionNewParamsAddPricesPriceNewPlanTieredPackagePriceBillingCycleConfigurationDurationUnit) IsKnown() bool {
	switch r {
	case PlanVersionNewParamsAddPricesPriceNewPlanTieredPackagePriceBillingCycleConfigurationDurationUnitDay, PlanVersionNewParamsAddPricesPriceNewPlanTieredPackagePriceBillingCycleConfigurationDurationUnitMonth:
		return true
	}
	return false
}

// For dimensional price: specifies a price group and dimension values
type PlanVersionNewParamsAddPricesPriceNewPlanTieredPackagePriceDimensionalPriceConfiguration struct {
	// The list of dimension values matching (in order) the dimensions of the price
	// group
	DimensionValues param.Field[[]string] `json:"dimension_values,required"`
	// The id of the dimensional price group to include this price in
	DimensionalPriceGroupID param.Field[string] `json:"dimensional_price_group_id"`
	// The external id of the dimensional price group to include this price in
	ExternalDimensionalPriceGroupID param.Field[string] `json:"external_dimensional_price_group_id"`
}

func (r PlanVersionNewParamsAddPricesPriceNewPlanTieredPackagePriceDimensionalPriceConfiguration) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Within each billing cycle, specifies the cadence at which invoices are produced.
// If unspecified, a single invoice is produced per billing cycle.
type PlanVersionNewParamsAddPricesPriceNewPlanTieredPackagePriceInvoicingCycleConfiguration struct {
	// The duration of the billing period.
	Duration param.Field[int64] `json:"duration,required"`
	// The unit of billing period duration.
	DurationUnit param.Field[PlanVersionNewParamsAddPricesPriceNewPlanTieredPackagePriceInvoicingCycleConfigurationDurationUnit] `json:"duration_unit,required"`
}

func (r PlanVersionNewParamsAddPricesPriceNewPlanTieredPackagePriceInvoicingCycleConfiguration) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// The unit of billing period duration.
type PlanVersionNewParamsAddPricesPriceNewPlanTieredPackagePriceInvoicingCycleConfigurationDurationUnit string

const (
	PlanVersionNewParamsAddPricesPriceNewPlanTieredPackagePriceInvoicingCycleConfigurationDurationUnitDay   PlanVersionNewParamsAddPricesPriceNewPlanTieredPackagePriceInvoicingCycleConfigurationDurationUnit = "day"
	PlanVersionNewParamsAddPricesPriceNewPlanTieredPackagePriceInvoicingCycleConfigurationDurationUnitMonth PlanVersionNewParamsAddPricesPriceNewPlanTieredPackagePriceInvoicingCycleConfigurationDurationUnit = "month"
)

func (r PlanVersionNewParamsAddPricesPriceNewPlanTieredPackagePriceInvoicingCycleConfigurationDurationUnit) IsKnown() bool {
	switch r {
	case PlanVersionNewParamsAddPricesPriceNewPlanTieredPackagePriceInvoicingCycleConfigurationDurationUnitDay, PlanVersionNewParamsAddPricesPriceNewPlanTieredPackagePriceInvoicingCycleConfigurationDurationUnitMonth:
		return true
	}
	return false
}

type PlanVersionNewParamsAddPricesPriceNewPlanTieredWithMinimumPrice struct {
	// The cadence to bill for this price on.
	Cadence param.Field[PlanVersionNewParamsAddPricesPriceNewPlanTieredWithMinimumPriceCadence] `json:"cadence,required"`
	// The id of the item the price will be associated with.
	ItemID    param.Field[string]                                                                   `json:"item_id,required"`
	ModelType param.Field[PlanVersionNewParamsAddPricesPriceNewPlanTieredWithMinimumPriceModelType] `json:"model_type,required"`
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
	BillingCycleConfiguration param.Field[PlanVersionNewParamsAddPricesPriceNewPlanTieredWithMinimumPriceBillingCycleConfiguration] `json:"billing_cycle_configuration"`
	// The per unit conversion rate of the price currency to the invoicing currency.
	ConversionRate param.Field[float64] `json:"conversion_rate"`
	// An ISO 4217 currency string, or custom pricing unit identifier, in which this
	// price is billed.
	Currency param.Field[string] `json:"currency"`
	// For dimensional price: specifies a price group and dimension values
	DimensionalPriceConfiguration param.Field[PlanVersionNewParamsAddPricesPriceNewPlanTieredWithMinimumPriceDimensionalPriceConfiguration] `json:"dimensional_price_configuration"`
	// An alias for the price.
	ExternalPriceID param.Field[string] `json:"external_price_id"`
	// If the Price represents a fixed cost, this represents the quantity of units
	// applied.
	FixedPriceQuantity param.Field[float64] `json:"fixed_price_quantity"`
	// The property used to group this price on an invoice
	InvoiceGroupingKey param.Field[string] `json:"invoice_grouping_key"`
	// Within each billing cycle, specifies the cadence at which invoices are produced.
	// If unspecified, a single invoice is produced per billing cycle.
	InvoicingCycleConfiguration param.Field[PlanVersionNewParamsAddPricesPriceNewPlanTieredWithMinimumPriceInvoicingCycleConfiguration] `json:"invoicing_cycle_configuration"`
	// User-specified key/value pairs for the resource. Individual keys can be removed
	// by setting the value to `null`, and the entire metadata mapping can be cleared
	// by setting `metadata` to `null`.
	Metadata param.Field[map[string]string] `json:"metadata"`
}

func (r PlanVersionNewParamsAddPricesPriceNewPlanTieredWithMinimumPrice) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r PlanVersionNewParamsAddPricesPriceNewPlanTieredWithMinimumPrice) implementsPlanVersionNewParamsAddPricesPriceUnion() {
}

// The cadence to bill for this price on.
type PlanVersionNewParamsAddPricesPriceNewPlanTieredWithMinimumPriceCadence string

const (
	PlanVersionNewParamsAddPricesPriceNewPlanTieredWithMinimumPriceCadenceAnnual     PlanVersionNewParamsAddPricesPriceNewPlanTieredWithMinimumPriceCadence = "annual"
	PlanVersionNewParamsAddPricesPriceNewPlanTieredWithMinimumPriceCadenceSemiAnnual PlanVersionNewParamsAddPricesPriceNewPlanTieredWithMinimumPriceCadence = "semi_annual"
	PlanVersionNewParamsAddPricesPriceNewPlanTieredWithMinimumPriceCadenceMonthly    PlanVersionNewParamsAddPricesPriceNewPlanTieredWithMinimumPriceCadence = "monthly"
	PlanVersionNewParamsAddPricesPriceNewPlanTieredWithMinimumPriceCadenceQuarterly  PlanVersionNewParamsAddPricesPriceNewPlanTieredWithMinimumPriceCadence = "quarterly"
	PlanVersionNewParamsAddPricesPriceNewPlanTieredWithMinimumPriceCadenceOneTime    PlanVersionNewParamsAddPricesPriceNewPlanTieredWithMinimumPriceCadence = "one_time"
	PlanVersionNewParamsAddPricesPriceNewPlanTieredWithMinimumPriceCadenceCustom     PlanVersionNewParamsAddPricesPriceNewPlanTieredWithMinimumPriceCadence = "custom"
)

func (r PlanVersionNewParamsAddPricesPriceNewPlanTieredWithMinimumPriceCadence) IsKnown() bool {
	switch r {
	case PlanVersionNewParamsAddPricesPriceNewPlanTieredWithMinimumPriceCadenceAnnual, PlanVersionNewParamsAddPricesPriceNewPlanTieredWithMinimumPriceCadenceSemiAnnual, PlanVersionNewParamsAddPricesPriceNewPlanTieredWithMinimumPriceCadenceMonthly, PlanVersionNewParamsAddPricesPriceNewPlanTieredWithMinimumPriceCadenceQuarterly, PlanVersionNewParamsAddPricesPriceNewPlanTieredWithMinimumPriceCadenceOneTime, PlanVersionNewParamsAddPricesPriceNewPlanTieredWithMinimumPriceCadenceCustom:
		return true
	}
	return false
}

type PlanVersionNewParamsAddPricesPriceNewPlanTieredWithMinimumPriceModelType string

const (
	PlanVersionNewParamsAddPricesPriceNewPlanTieredWithMinimumPriceModelTypeTieredWithMinimum PlanVersionNewParamsAddPricesPriceNewPlanTieredWithMinimumPriceModelType = "tiered_with_minimum"
)

func (r PlanVersionNewParamsAddPricesPriceNewPlanTieredWithMinimumPriceModelType) IsKnown() bool {
	switch r {
	case PlanVersionNewParamsAddPricesPriceNewPlanTieredWithMinimumPriceModelTypeTieredWithMinimum:
		return true
	}
	return false
}

// For custom cadence: specifies the duration of the billing period in days or
// months.
type PlanVersionNewParamsAddPricesPriceNewPlanTieredWithMinimumPriceBillingCycleConfiguration struct {
	// The duration of the billing period.
	Duration param.Field[int64] `json:"duration,required"`
	// The unit of billing period duration.
	DurationUnit param.Field[PlanVersionNewParamsAddPricesPriceNewPlanTieredWithMinimumPriceBillingCycleConfigurationDurationUnit] `json:"duration_unit,required"`
}

func (r PlanVersionNewParamsAddPricesPriceNewPlanTieredWithMinimumPriceBillingCycleConfiguration) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// The unit of billing period duration.
type PlanVersionNewParamsAddPricesPriceNewPlanTieredWithMinimumPriceBillingCycleConfigurationDurationUnit string

const (
	PlanVersionNewParamsAddPricesPriceNewPlanTieredWithMinimumPriceBillingCycleConfigurationDurationUnitDay   PlanVersionNewParamsAddPricesPriceNewPlanTieredWithMinimumPriceBillingCycleConfigurationDurationUnit = "day"
	PlanVersionNewParamsAddPricesPriceNewPlanTieredWithMinimumPriceBillingCycleConfigurationDurationUnitMonth PlanVersionNewParamsAddPricesPriceNewPlanTieredWithMinimumPriceBillingCycleConfigurationDurationUnit = "month"
)

func (r PlanVersionNewParamsAddPricesPriceNewPlanTieredWithMinimumPriceBillingCycleConfigurationDurationUnit) IsKnown() bool {
	switch r {
	case PlanVersionNewParamsAddPricesPriceNewPlanTieredWithMinimumPriceBillingCycleConfigurationDurationUnitDay, PlanVersionNewParamsAddPricesPriceNewPlanTieredWithMinimumPriceBillingCycleConfigurationDurationUnitMonth:
		return true
	}
	return false
}

// For dimensional price: specifies a price group and dimension values
type PlanVersionNewParamsAddPricesPriceNewPlanTieredWithMinimumPriceDimensionalPriceConfiguration struct {
	// The list of dimension values matching (in order) the dimensions of the price
	// group
	DimensionValues param.Field[[]string] `json:"dimension_values,required"`
	// The id of the dimensional price group to include this price in
	DimensionalPriceGroupID param.Field[string] `json:"dimensional_price_group_id"`
	// The external id of the dimensional price group to include this price in
	ExternalDimensionalPriceGroupID param.Field[string] `json:"external_dimensional_price_group_id"`
}

func (r PlanVersionNewParamsAddPricesPriceNewPlanTieredWithMinimumPriceDimensionalPriceConfiguration) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Within each billing cycle, specifies the cadence at which invoices are produced.
// If unspecified, a single invoice is produced per billing cycle.
type PlanVersionNewParamsAddPricesPriceNewPlanTieredWithMinimumPriceInvoicingCycleConfiguration struct {
	// The duration of the billing period.
	Duration param.Field[int64] `json:"duration,required"`
	// The unit of billing period duration.
	DurationUnit param.Field[PlanVersionNewParamsAddPricesPriceNewPlanTieredWithMinimumPriceInvoicingCycleConfigurationDurationUnit] `json:"duration_unit,required"`
}

func (r PlanVersionNewParamsAddPricesPriceNewPlanTieredWithMinimumPriceInvoicingCycleConfiguration) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// The unit of billing period duration.
type PlanVersionNewParamsAddPricesPriceNewPlanTieredWithMinimumPriceInvoicingCycleConfigurationDurationUnit string

const (
	PlanVersionNewParamsAddPricesPriceNewPlanTieredWithMinimumPriceInvoicingCycleConfigurationDurationUnitDay   PlanVersionNewParamsAddPricesPriceNewPlanTieredWithMinimumPriceInvoicingCycleConfigurationDurationUnit = "day"
	PlanVersionNewParamsAddPricesPriceNewPlanTieredWithMinimumPriceInvoicingCycleConfigurationDurationUnitMonth PlanVersionNewParamsAddPricesPriceNewPlanTieredWithMinimumPriceInvoicingCycleConfigurationDurationUnit = "month"
)

func (r PlanVersionNewParamsAddPricesPriceNewPlanTieredWithMinimumPriceInvoicingCycleConfigurationDurationUnit) IsKnown() bool {
	switch r {
	case PlanVersionNewParamsAddPricesPriceNewPlanTieredWithMinimumPriceInvoicingCycleConfigurationDurationUnitDay, PlanVersionNewParamsAddPricesPriceNewPlanTieredWithMinimumPriceInvoicingCycleConfigurationDurationUnitMonth:
		return true
	}
	return false
}

type PlanVersionNewParamsAddPricesPriceNewPlanUnitWithPercentPrice struct {
	// The cadence to bill for this price on.
	Cadence param.Field[PlanVersionNewParamsAddPricesPriceNewPlanUnitWithPercentPriceCadence] `json:"cadence,required"`
	// The id of the item the price will be associated with.
	ItemID    param.Field[string]                                                                 `json:"item_id,required"`
	ModelType param.Field[PlanVersionNewParamsAddPricesPriceNewPlanUnitWithPercentPriceModelType] `json:"model_type,required"`
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
	BillingCycleConfiguration param.Field[PlanVersionNewParamsAddPricesPriceNewPlanUnitWithPercentPriceBillingCycleConfiguration] `json:"billing_cycle_configuration"`
	// The per unit conversion rate of the price currency to the invoicing currency.
	ConversionRate param.Field[float64] `json:"conversion_rate"`
	// An ISO 4217 currency string, or custom pricing unit identifier, in which this
	// price is billed.
	Currency param.Field[string] `json:"currency"`
	// For dimensional price: specifies a price group and dimension values
	DimensionalPriceConfiguration param.Field[PlanVersionNewParamsAddPricesPriceNewPlanUnitWithPercentPriceDimensionalPriceConfiguration] `json:"dimensional_price_configuration"`
	// An alias for the price.
	ExternalPriceID param.Field[string] `json:"external_price_id"`
	// If the Price represents a fixed cost, this represents the quantity of units
	// applied.
	FixedPriceQuantity param.Field[float64] `json:"fixed_price_quantity"`
	// The property used to group this price on an invoice
	InvoiceGroupingKey param.Field[string] `json:"invoice_grouping_key"`
	// Within each billing cycle, specifies the cadence at which invoices are produced.
	// If unspecified, a single invoice is produced per billing cycle.
	InvoicingCycleConfiguration param.Field[PlanVersionNewParamsAddPricesPriceNewPlanUnitWithPercentPriceInvoicingCycleConfiguration] `json:"invoicing_cycle_configuration"`
	// User-specified key/value pairs for the resource. Individual keys can be removed
	// by setting the value to `null`, and the entire metadata mapping can be cleared
	// by setting `metadata` to `null`.
	Metadata param.Field[map[string]string] `json:"metadata"`
}

func (r PlanVersionNewParamsAddPricesPriceNewPlanUnitWithPercentPrice) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r PlanVersionNewParamsAddPricesPriceNewPlanUnitWithPercentPrice) implementsPlanVersionNewParamsAddPricesPriceUnion() {
}

// The cadence to bill for this price on.
type PlanVersionNewParamsAddPricesPriceNewPlanUnitWithPercentPriceCadence string

const (
	PlanVersionNewParamsAddPricesPriceNewPlanUnitWithPercentPriceCadenceAnnual     PlanVersionNewParamsAddPricesPriceNewPlanUnitWithPercentPriceCadence = "annual"
	PlanVersionNewParamsAddPricesPriceNewPlanUnitWithPercentPriceCadenceSemiAnnual PlanVersionNewParamsAddPricesPriceNewPlanUnitWithPercentPriceCadence = "semi_annual"
	PlanVersionNewParamsAddPricesPriceNewPlanUnitWithPercentPriceCadenceMonthly    PlanVersionNewParamsAddPricesPriceNewPlanUnitWithPercentPriceCadence = "monthly"
	PlanVersionNewParamsAddPricesPriceNewPlanUnitWithPercentPriceCadenceQuarterly  PlanVersionNewParamsAddPricesPriceNewPlanUnitWithPercentPriceCadence = "quarterly"
	PlanVersionNewParamsAddPricesPriceNewPlanUnitWithPercentPriceCadenceOneTime    PlanVersionNewParamsAddPricesPriceNewPlanUnitWithPercentPriceCadence = "one_time"
	PlanVersionNewParamsAddPricesPriceNewPlanUnitWithPercentPriceCadenceCustom     PlanVersionNewParamsAddPricesPriceNewPlanUnitWithPercentPriceCadence = "custom"
)

func (r PlanVersionNewParamsAddPricesPriceNewPlanUnitWithPercentPriceCadence) IsKnown() bool {
	switch r {
	case PlanVersionNewParamsAddPricesPriceNewPlanUnitWithPercentPriceCadenceAnnual, PlanVersionNewParamsAddPricesPriceNewPlanUnitWithPercentPriceCadenceSemiAnnual, PlanVersionNewParamsAddPricesPriceNewPlanUnitWithPercentPriceCadenceMonthly, PlanVersionNewParamsAddPricesPriceNewPlanUnitWithPercentPriceCadenceQuarterly, PlanVersionNewParamsAddPricesPriceNewPlanUnitWithPercentPriceCadenceOneTime, PlanVersionNewParamsAddPricesPriceNewPlanUnitWithPercentPriceCadenceCustom:
		return true
	}
	return false
}

type PlanVersionNewParamsAddPricesPriceNewPlanUnitWithPercentPriceModelType string

const (
	PlanVersionNewParamsAddPricesPriceNewPlanUnitWithPercentPriceModelTypeUnitWithPercent PlanVersionNewParamsAddPricesPriceNewPlanUnitWithPercentPriceModelType = "unit_with_percent"
)

func (r PlanVersionNewParamsAddPricesPriceNewPlanUnitWithPercentPriceModelType) IsKnown() bool {
	switch r {
	case PlanVersionNewParamsAddPricesPriceNewPlanUnitWithPercentPriceModelTypeUnitWithPercent:
		return true
	}
	return false
}

// For custom cadence: specifies the duration of the billing period in days or
// months.
type PlanVersionNewParamsAddPricesPriceNewPlanUnitWithPercentPriceBillingCycleConfiguration struct {
	// The duration of the billing period.
	Duration param.Field[int64] `json:"duration,required"`
	// The unit of billing period duration.
	DurationUnit param.Field[PlanVersionNewParamsAddPricesPriceNewPlanUnitWithPercentPriceBillingCycleConfigurationDurationUnit] `json:"duration_unit,required"`
}

func (r PlanVersionNewParamsAddPricesPriceNewPlanUnitWithPercentPriceBillingCycleConfiguration) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// The unit of billing period duration.
type PlanVersionNewParamsAddPricesPriceNewPlanUnitWithPercentPriceBillingCycleConfigurationDurationUnit string

const (
	PlanVersionNewParamsAddPricesPriceNewPlanUnitWithPercentPriceBillingCycleConfigurationDurationUnitDay   PlanVersionNewParamsAddPricesPriceNewPlanUnitWithPercentPriceBillingCycleConfigurationDurationUnit = "day"
	PlanVersionNewParamsAddPricesPriceNewPlanUnitWithPercentPriceBillingCycleConfigurationDurationUnitMonth PlanVersionNewParamsAddPricesPriceNewPlanUnitWithPercentPriceBillingCycleConfigurationDurationUnit = "month"
)

func (r PlanVersionNewParamsAddPricesPriceNewPlanUnitWithPercentPriceBillingCycleConfigurationDurationUnit) IsKnown() bool {
	switch r {
	case PlanVersionNewParamsAddPricesPriceNewPlanUnitWithPercentPriceBillingCycleConfigurationDurationUnitDay, PlanVersionNewParamsAddPricesPriceNewPlanUnitWithPercentPriceBillingCycleConfigurationDurationUnitMonth:
		return true
	}
	return false
}

// For dimensional price: specifies a price group and dimension values
type PlanVersionNewParamsAddPricesPriceNewPlanUnitWithPercentPriceDimensionalPriceConfiguration struct {
	// The list of dimension values matching (in order) the dimensions of the price
	// group
	DimensionValues param.Field[[]string] `json:"dimension_values,required"`
	// The id of the dimensional price group to include this price in
	DimensionalPriceGroupID param.Field[string] `json:"dimensional_price_group_id"`
	// The external id of the dimensional price group to include this price in
	ExternalDimensionalPriceGroupID param.Field[string] `json:"external_dimensional_price_group_id"`
}

func (r PlanVersionNewParamsAddPricesPriceNewPlanUnitWithPercentPriceDimensionalPriceConfiguration) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Within each billing cycle, specifies the cadence at which invoices are produced.
// If unspecified, a single invoice is produced per billing cycle.
type PlanVersionNewParamsAddPricesPriceNewPlanUnitWithPercentPriceInvoicingCycleConfiguration struct {
	// The duration of the billing period.
	Duration param.Field[int64] `json:"duration,required"`
	// The unit of billing period duration.
	DurationUnit param.Field[PlanVersionNewParamsAddPricesPriceNewPlanUnitWithPercentPriceInvoicingCycleConfigurationDurationUnit] `json:"duration_unit,required"`
}

func (r PlanVersionNewParamsAddPricesPriceNewPlanUnitWithPercentPriceInvoicingCycleConfiguration) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// The unit of billing period duration.
type PlanVersionNewParamsAddPricesPriceNewPlanUnitWithPercentPriceInvoicingCycleConfigurationDurationUnit string

const (
	PlanVersionNewParamsAddPricesPriceNewPlanUnitWithPercentPriceInvoicingCycleConfigurationDurationUnitDay   PlanVersionNewParamsAddPricesPriceNewPlanUnitWithPercentPriceInvoicingCycleConfigurationDurationUnit = "day"
	PlanVersionNewParamsAddPricesPriceNewPlanUnitWithPercentPriceInvoicingCycleConfigurationDurationUnitMonth PlanVersionNewParamsAddPricesPriceNewPlanUnitWithPercentPriceInvoicingCycleConfigurationDurationUnit = "month"
)

func (r PlanVersionNewParamsAddPricesPriceNewPlanUnitWithPercentPriceInvoicingCycleConfigurationDurationUnit) IsKnown() bool {
	switch r {
	case PlanVersionNewParamsAddPricesPriceNewPlanUnitWithPercentPriceInvoicingCycleConfigurationDurationUnitDay, PlanVersionNewParamsAddPricesPriceNewPlanUnitWithPercentPriceInvoicingCycleConfigurationDurationUnitMonth:
		return true
	}
	return false
}

type PlanVersionNewParamsAddPricesPriceNewPlanPackageWithAllocationPrice struct {
	// The cadence to bill for this price on.
	Cadence param.Field[PlanVersionNewParamsAddPricesPriceNewPlanPackageWithAllocationPriceCadence] `json:"cadence,required"`
	// The id of the item the price will be associated with.
	ItemID    param.Field[string]                                                                       `json:"item_id,required"`
	ModelType param.Field[PlanVersionNewParamsAddPricesPriceNewPlanPackageWithAllocationPriceModelType] `json:"model_type,required"`
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
	BillingCycleConfiguration param.Field[PlanVersionNewParamsAddPricesPriceNewPlanPackageWithAllocationPriceBillingCycleConfiguration] `json:"billing_cycle_configuration"`
	// The per unit conversion rate of the price currency to the invoicing currency.
	ConversionRate param.Field[float64] `json:"conversion_rate"`
	// An ISO 4217 currency string, or custom pricing unit identifier, in which this
	// price is billed.
	Currency param.Field[string] `json:"currency"`
	// For dimensional price: specifies a price group and dimension values
	DimensionalPriceConfiguration param.Field[PlanVersionNewParamsAddPricesPriceNewPlanPackageWithAllocationPriceDimensionalPriceConfiguration] `json:"dimensional_price_configuration"`
	// An alias for the price.
	ExternalPriceID param.Field[string] `json:"external_price_id"`
	// If the Price represents a fixed cost, this represents the quantity of units
	// applied.
	FixedPriceQuantity param.Field[float64] `json:"fixed_price_quantity"`
	// The property used to group this price on an invoice
	InvoiceGroupingKey param.Field[string] `json:"invoice_grouping_key"`
	// Within each billing cycle, specifies the cadence at which invoices are produced.
	// If unspecified, a single invoice is produced per billing cycle.
	InvoicingCycleConfiguration param.Field[PlanVersionNewParamsAddPricesPriceNewPlanPackageWithAllocationPriceInvoicingCycleConfiguration] `json:"invoicing_cycle_configuration"`
	// User-specified key/value pairs for the resource. Individual keys can be removed
	// by setting the value to `null`, and the entire metadata mapping can be cleared
	// by setting `metadata` to `null`.
	Metadata param.Field[map[string]string] `json:"metadata"`
}

func (r PlanVersionNewParamsAddPricesPriceNewPlanPackageWithAllocationPrice) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r PlanVersionNewParamsAddPricesPriceNewPlanPackageWithAllocationPrice) implementsPlanVersionNewParamsAddPricesPriceUnion() {
}

// The cadence to bill for this price on.
type PlanVersionNewParamsAddPricesPriceNewPlanPackageWithAllocationPriceCadence string

const (
	PlanVersionNewParamsAddPricesPriceNewPlanPackageWithAllocationPriceCadenceAnnual     PlanVersionNewParamsAddPricesPriceNewPlanPackageWithAllocationPriceCadence = "annual"
	PlanVersionNewParamsAddPricesPriceNewPlanPackageWithAllocationPriceCadenceSemiAnnual PlanVersionNewParamsAddPricesPriceNewPlanPackageWithAllocationPriceCadence = "semi_annual"
	PlanVersionNewParamsAddPricesPriceNewPlanPackageWithAllocationPriceCadenceMonthly    PlanVersionNewParamsAddPricesPriceNewPlanPackageWithAllocationPriceCadence = "monthly"
	PlanVersionNewParamsAddPricesPriceNewPlanPackageWithAllocationPriceCadenceQuarterly  PlanVersionNewParamsAddPricesPriceNewPlanPackageWithAllocationPriceCadence = "quarterly"
	PlanVersionNewParamsAddPricesPriceNewPlanPackageWithAllocationPriceCadenceOneTime    PlanVersionNewParamsAddPricesPriceNewPlanPackageWithAllocationPriceCadence = "one_time"
	PlanVersionNewParamsAddPricesPriceNewPlanPackageWithAllocationPriceCadenceCustom     PlanVersionNewParamsAddPricesPriceNewPlanPackageWithAllocationPriceCadence = "custom"
)

func (r PlanVersionNewParamsAddPricesPriceNewPlanPackageWithAllocationPriceCadence) IsKnown() bool {
	switch r {
	case PlanVersionNewParamsAddPricesPriceNewPlanPackageWithAllocationPriceCadenceAnnual, PlanVersionNewParamsAddPricesPriceNewPlanPackageWithAllocationPriceCadenceSemiAnnual, PlanVersionNewParamsAddPricesPriceNewPlanPackageWithAllocationPriceCadenceMonthly, PlanVersionNewParamsAddPricesPriceNewPlanPackageWithAllocationPriceCadenceQuarterly, PlanVersionNewParamsAddPricesPriceNewPlanPackageWithAllocationPriceCadenceOneTime, PlanVersionNewParamsAddPricesPriceNewPlanPackageWithAllocationPriceCadenceCustom:
		return true
	}
	return false
}

type PlanVersionNewParamsAddPricesPriceNewPlanPackageWithAllocationPriceModelType string

const (
	PlanVersionNewParamsAddPricesPriceNewPlanPackageWithAllocationPriceModelTypePackageWithAllocation PlanVersionNewParamsAddPricesPriceNewPlanPackageWithAllocationPriceModelType = "package_with_allocation"
)

func (r PlanVersionNewParamsAddPricesPriceNewPlanPackageWithAllocationPriceModelType) IsKnown() bool {
	switch r {
	case PlanVersionNewParamsAddPricesPriceNewPlanPackageWithAllocationPriceModelTypePackageWithAllocation:
		return true
	}
	return false
}

// For custom cadence: specifies the duration of the billing period in days or
// months.
type PlanVersionNewParamsAddPricesPriceNewPlanPackageWithAllocationPriceBillingCycleConfiguration struct {
	// The duration of the billing period.
	Duration param.Field[int64] `json:"duration,required"`
	// The unit of billing period duration.
	DurationUnit param.Field[PlanVersionNewParamsAddPricesPriceNewPlanPackageWithAllocationPriceBillingCycleConfigurationDurationUnit] `json:"duration_unit,required"`
}

func (r PlanVersionNewParamsAddPricesPriceNewPlanPackageWithAllocationPriceBillingCycleConfiguration) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// The unit of billing period duration.
type PlanVersionNewParamsAddPricesPriceNewPlanPackageWithAllocationPriceBillingCycleConfigurationDurationUnit string

const (
	PlanVersionNewParamsAddPricesPriceNewPlanPackageWithAllocationPriceBillingCycleConfigurationDurationUnitDay   PlanVersionNewParamsAddPricesPriceNewPlanPackageWithAllocationPriceBillingCycleConfigurationDurationUnit = "day"
	PlanVersionNewParamsAddPricesPriceNewPlanPackageWithAllocationPriceBillingCycleConfigurationDurationUnitMonth PlanVersionNewParamsAddPricesPriceNewPlanPackageWithAllocationPriceBillingCycleConfigurationDurationUnit = "month"
)

func (r PlanVersionNewParamsAddPricesPriceNewPlanPackageWithAllocationPriceBillingCycleConfigurationDurationUnit) IsKnown() bool {
	switch r {
	case PlanVersionNewParamsAddPricesPriceNewPlanPackageWithAllocationPriceBillingCycleConfigurationDurationUnitDay, PlanVersionNewParamsAddPricesPriceNewPlanPackageWithAllocationPriceBillingCycleConfigurationDurationUnitMonth:
		return true
	}
	return false
}

// For dimensional price: specifies a price group and dimension values
type PlanVersionNewParamsAddPricesPriceNewPlanPackageWithAllocationPriceDimensionalPriceConfiguration struct {
	// The list of dimension values matching (in order) the dimensions of the price
	// group
	DimensionValues param.Field[[]string] `json:"dimension_values,required"`
	// The id of the dimensional price group to include this price in
	DimensionalPriceGroupID param.Field[string] `json:"dimensional_price_group_id"`
	// The external id of the dimensional price group to include this price in
	ExternalDimensionalPriceGroupID param.Field[string] `json:"external_dimensional_price_group_id"`
}

func (r PlanVersionNewParamsAddPricesPriceNewPlanPackageWithAllocationPriceDimensionalPriceConfiguration) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Within each billing cycle, specifies the cadence at which invoices are produced.
// If unspecified, a single invoice is produced per billing cycle.
type PlanVersionNewParamsAddPricesPriceNewPlanPackageWithAllocationPriceInvoicingCycleConfiguration struct {
	// The duration of the billing period.
	Duration param.Field[int64] `json:"duration,required"`
	// The unit of billing period duration.
	DurationUnit param.Field[PlanVersionNewParamsAddPricesPriceNewPlanPackageWithAllocationPriceInvoicingCycleConfigurationDurationUnit] `json:"duration_unit,required"`
}

func (r PlanVersionNewParamsAddPricesPriceNewPlanPackageWithAllocationPriceInvoicingCycleConfiguration) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// The unit of billing period duration.
type PlanVersionNewParamsAddPricesPriceNewPlanPackageWithAllocationPriceInvoicingCycleConfigurationDurationUnit string

const (
	PlanVersionNewParamsAddPricesPriceNewPlanPackageWithAllocationPriceInvoicingCycleConfigurationDurationUnitDay   PlanVersionNewParamsAddPricesPriceNewPlanPackageWithAllocationPriceInvoicingCycleConfigurationDurationUnit = "day"
	PlanVersionNewParamsAddPricesPriceNewPlanPackageWithAllocationPriceInvoicingCycleConfigurationDurationUnitMonth PlanVersionNewParamsAddPricesPriceNewPlanPackageWithAllocationPriceInvoicingCycleConfigurationDurationUnit = "month"
)

func (r PlanVersionNewParamsAddPricesPriceNewPlanPackageWithAllocationPriceInvoicingCycleConfigurationDurationUnit) IsKnown() bool {
	switch r {
	case PlanVersionNewParamsAddPricesPriceNewPlanPackageWithAllocationPriceInvoicingCycleConfigurationDurationUnitDay, PlanVersionNewParamsAddPricesPriceNewPlanPackageWithAllocationPriceInvoicingCycleConfigurationDurationUnitMonth:
		return true
	}
	return false
}

type PlanVersionNewParamsAddPricesPriceNewPlanTierWithProrationPrice struct {
	// The cadence to bill for this price on.
	Cadence param.Field[PlanVersionNewParamsAddPricesPriceNewPlanTierWithProrationPriceCadence] `json:"cadence,required"`
	// The id of the item the price will be associated with.
	ItemID    param.Field[string]                                                                   `json:"item_id,required"`
	ModelType param.Field[PlanVersionNewParamsAddPricesPriceNewPlanTierWithProrationPriceModelType] `json:"model_type,required"`
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
	BillingCycleConfiguration param.Field[PlanVersionNewParamsAddPricesPriceNewPlanTierWithProrationPriceBillingCycleConfiguration] `json:"billing_cycle_configuration"`
	// The per unit conversion rate of the price currency to the invoicing currency.
	ConversionRate param.Field[float64] `json:"conversion_rate"`
	// An ISO 4217 currency string, or custom pricing unit identifier, in which this
	// price is billed.
	Currency param.Field[string] `json:"currency"`
	// For dimensional price: specifies a price group and dimension values
	DimensionalPriceConfiguration param.Field[PlanVersionNewParamsAddPricesPriceNewPlanTierWithProrationPriceDimensionalPriceConfiguration] `json:"dimensional_price_configuration"`
	// An alias for the price.
	ExternalPriceID param.Field[string] `json:"external_price_id"`
	// If the Price represents a fixed cost, this represents the quantity of units
	// applied.
	FixedPriceQuantity param.Field[float64] `json:"fixed_price_quantity"`
	// The property used to group this price on an invoice
	InvoiceGroupingKey param.Field[string] `json:"invoice_grouping_key"`
	// Within each billing cycle, specifies the cadence at which invoices are produced.
	// If unspecified, a single invoice is produced per billing cycle.
	InvoicingCycleConfiguration param.Field[PlanVersionNewParamsAddPricesPriceNewPlanTierWithProrationPriceInvoicingCycleConfiguration] `json:"invoicing_cycle_configuration"`
	// User-specified key/value pairs for the resource. Individual keys can be removed
	// by setting the value to `null`, and the entire metadata mapping can be cleared
	// by setting `metadata` to `null`.
	Metadata param.Field[map[string]string] `json:"metadata"`
}

func (r PlanVersionNewParamsAddPricesPriceNewPlanTierWithProrationPrice) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r PlanVersionNewParamsAddPricesPriceNewPlanTierWithProrationPrice) implementsPlanVersionNewParamsAddPricesPriceUnion() {
}

// The cadence to bill for this price on.
type PlanVersionNewParamsAddPricesPriceNewPlanTierWithProrationPriceCadence string

const (
	PlanVersionNewParamsAddPricesPriceNewPlanTierWithProrationPriceCadenceAnnual     PlanVersionNewParamsAddPricesPriceNewPlanTierWithProrationPriceCadence = "annual"
	PlanVersionNewParamsAddPricesPriceNewPlanTierWithProrationPriceCadenceSemiAnnual PlanVersionNewParamsAddPricesPriceNewPlanTierWithProrationPriceCadence = "semi_annual"
	PlanVersionNewParamsAddPricesPriceNewPlanTierWithProrationPriceCadenceMonthly    PlanVersionNewParamsAddPricesPriceNewPlanTierWithProrationPriceCadence = "monthly"
	PlanVersionNewParamsAddPricesPriceNewPlanTierWithProrationPriceCadenceQuarterly  PlanVersionNewParamsAddPricesPriceNewPlanTierWithProrationPriceCadence = "quarterly"
	PlanVersionNewParamsAddPricesPriceNewPlanTierWithProrationPriceCadenceOneTime    PlanVersionNewParamsAddPricesPriceNewPlanTierWithProrationPriceCadence = "one_time"
	PlanVersionNewParamsAddPricesPriceNewPlanTierWithProrationPriceCadenceCustom     PlanVersionNewParamsAddPricesPriceNewPlanTierWithProrationPriceCadence = "custom"
)

func (r PlanVersionNewParamsAddPricesPriceNewPlanTierWithProrationPriceCadence) IsKnown() bool {
	switch r {
	case PlanVersionNewParamsAddPricesPriceNewPlanTierWithProrationPriceCadenceAnnual, PlanVersionNewParamsAddPricesPriceNewPlanTierWithProrationPriceCadenceSemiAnnual, PlanVersionNewParamsAddPricesPriceNewPlanTierWithProrationPriceCadenceMonthly, PlanVersionNewParamsAddPricesPriceNewPlanTierWithProrationPriceCadenceQuarterly, PlanVersionNewParamsAddPricesPriceNewPlanTierWithProrationPriceCadenceOneTime, PlanVersionNewParamsAddPricesPriceNewPlanTierWithProrationPriceCadenceCustom:
		return true
	}
	return false
}

type PlanVersionNewParamsAddPricesPriceNewPlanTierWithProrationPriceModelType string

const (
	PlanVersionNewParamsAddPricesPriceNewPlanTierWithProrationPriceModelTypeTieredWithProration PlanVersionNewParamsAddPricesPriceNewPlanTierWithProrationPriceModelType = "tiered_with_proration"
)

func (r PlanVersionNewParamsAddPricesPriceNewPlanTierWithProrationPriceModelType) IsKnown() bool {
	switch r {
	case PlanVersionNewParamsAddPricesPriceNewPlanTierWithProrationPriceModelTypeTieredWithProration:
		return true
	}
	return false
}

// For custom cadence: specifies the duration of the billing period in days or
// months.
type PlanVersionNewParamsAddPricesPriceNewPlanTierWithProrationPriceBillingCycleConfiguration struct {
	// The duration of the billing period.
	Duration param.Field[int64] `json:"duration,required"`
	// The unit of billing period duration.
	DurationUnit param.Field[PlanVersionNewParamsAddPricesPriceNewPlanTierWithProrationPriceBillingCycleConfigurationDurationUnit] `json:"duration_unit,required"`
}

func (r PlanVersionNewParamsAddPricesPriceNewPlanTierWithProrationPriceBillingCycleConfiguration) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// The unit of billing period duration.
type PlanVersionNewParamsAddPricesPriceNewPlanTierWithProrationPriceBillingCycleConfigurationDurationUnit string

const (
	PlanVersionNewParamsAddPricesPriceNewPlanTierWithProrationPriceBillingCycleConfigurationDurationUnitDay   PlanVersionNewParamsAddPricesPriceNewPlanTierWithProrationPriceBillingCycleConfigurationDurationUnit = "day"
	PlanVersionNewParamsAddPricesPriceNewPlanTierWithProrationPriceBillingCycleConfigurationDurationUnitMonth PlanVersionNewParamsAddPricesPriceNewPlanTierWithProrationPriceBillingCycleConfigurationDurationUnit = "month"
)

func (r PlanVersionNewParamsAddPricesPriceNewPlanTierWithProrationPriceBillingCycleConfigurationDurationUnit) IsKnown() bool {
	switch r {
	case PlanVersionNewParamsAddPricesPriceNewPlanTierWithProrationPriceBillingCycleConfigurationDurationUnitDay, PlanVersionNewParamsAddPricesPriceNewPlanTierWithProrationPriceBillingCycleConfigurationDurationUnitMonth:
		return true
	}
	return false
}

// For dimensional price: specifies a price group and dimension values
type PlanVersionNewParamsAddPricesPriceNewPlanTierWithProrationPriceDimensionalPriceConfiguration struct {
	// The list of dimension values matching (in order) the dimensions of the price
	// group
	DimensionValues param.Field[[]string] `json:"dimension_values,required"`
	// The id of the dimensional price group to include this price in
	DimensionalPriceGroupID param.Field[string] `json:"dimensional_price_group_id"`
	// The external id of the dimensional price group to include this price in
	ExternalDimensionalPriceGroupID param.Field[string] `json:"external_dimensional_price_group_id"`
}

func (r PlanVersionNewParamsAddPricesPriceNewPlanTierWithProrationPriceDimensionalPriceConfiguration) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Within each billing cycle, specifies the cadence at which invoices are produced.
// If unspecified, a single invoice is produced per billing cycle.
type PlanVersionNewParamsAddPricesPriceNewPlanTierWithProrationPriceInvoicingCycleConfiguration struct {
	// The duration of the billing period.
	Duration param.Field[int64] `json:"duration,required"`
	// The unit of billing period duration.
	DurationUnit param.Field[PlanVersionNewParamsAddPricesPriceNewPlanTierWithProrationPriceInvoicingCycleConfigurationDurationUnit] `json:"duration_unit,required"`
}

func (r PlanVersionNewParamsAddPricesPriceNewPlanTierWithProrationPriceInvoicingCycleConfiguration) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// The unit of billing period duration.
type PlanVersionNewParamsAddPricesPriceNewPlanTierWithProrationPriceInvoicingCycleConfigurationDurationUnit string

const (
	PlanVersionNewParamsAddPricesPriceNewPlanTierWithProrationPriceInvoicingCycleConfigurationDurationUnitDay   PlanVersionNewParamsAddPricesPriceNewPlanTierWithProrationPriceInvoicingCycleConfigurationDurationUnit = "day"
	PlanVersionNewParamsAddPricesPriceNewPlanTierWithProrationPriceInvoicingCycleConfigurationDurationUnitMonth PlanVersionNewParamsAddPricesPriceNewPlanTierWithProrationPriceInvoicingCycleConfigurationDurationUnit = "month"
)

func (r PlanVersionNewParamsAddPricesPriceNewPlanTierWithProrationPriceInvoicingCycleConfigurationDurationUnit) IsKnown() bool {
	switch r {
	case PlanVersionNewParamsAddPricesPriceNewPlanTierWithProrationPriceInvoicingCycleConfigurationDurationUnitDay, PlanVersionNewParamsAddPricesPriceNewPlanTierWithProrationPriceInvoicingCycleConfigurationDurationUnitMonth:
		return true
	}
	return false
}

type PlanVersionNewParamsAddPricesPriceNewPlanUnitWithProrationPrice struct {
	// The cadence to bill for this price on.
	Cadence param.Field[PlanVersionNewParamsAddPricesPriceNewPlanUnitWithProrationPriceCadence] `json:"cadence,required"`
	// The id of the item the price will be associated with.
	ItemID    param.Field[string]                                                                   `json:"item_id,required"`
	ModelType param.Field[PlanVersionNewParamsAddPricesPriceNewPlanUnitWithProrationPriceModelType] `json:"model_type,required"`
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
	BillingCycleConfiguration param.Field[PlanVersionNewParamsAddPricesPriceNewPlanUnitWithProrationPriceBillingCycleConfiguration] `json:"billing_cycle_configuration"`
	// The per unit conversion rate of the price currency to the invoicing currency.
	ConversionRate param.Field[float64] `json:"conversion_rate"`
	// An ISO 4217 currency string, or custom pricing unit identifier, in which this
	// price is billed.
	Currency param.Field[string] `json:"currency"`
	// For dimensional price: specifies a price group and dimension values
	DimensionalPriceConfiguration param.Field[PlanVersionNewParamsAddPricesPriceNewPlanUnitWithProrationPriceDimensionalPriceConfiguration] `json:"dimensional_price_configuration"`
	// An alias for the price.
	ExternalPriceID param.Field[string] `json:"external_price_id"`
	// If the Price represents a fixed cost, this represents the quantity of units
	// applied.
	FixedPriceQuantity param.Field[float64] `json:"fixed_price_quantity"`
	// The property used to group this price on an invoice
	InvoiceGroupingKey param.Field[string] `json:"invoice_grouping_key"`
	// Within each billing cycle, specifies the cadence at which invoices are produced.
	// If unspecified, a single invoice is produced per billing cycle.
	InvoicingCycleConfiguration param.Field[PlanVersionNewParamsAddPricesPriceNewPlanUnitWithProrationPriceInvoicingCycleConfiguration] `json:"invoicing_cycle_configuration"`
	// User-specified key/value pairs for the resource. Individual keys can be removed
	// by setting the value to `null`, and the entire metadata mapping can be cleared
	// by setting `metadata` to `null`.
	Metadata param.Field[map[string]string] `json:"metadata"`
}

func (r PlanVersionNewParamsAddPricesPriceNewPlanUnitWithProrationPrice) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r PlanVersionNewParamsAddPricesPriceNewPlanUnitWithProrationPrice) implementsPlanVersionNewParamsAddPricesPriceUnion() {
}

// The cadence to bill for this price on.
type PlanVersionNewParamsAddPricesPriceNewPlanUnitWithProrationPriceCadence string

const (
	PlanVersionNewParamsAddPricesPriceNewPlanUnitWithProrationPriceCadenceAnnual     PlanVersionNewParamsAddPricesPriceNewPlanUnitWithProrationPriceCadence = "annual"
	PlanVersionNewParamsAddPricesPriceNewPlanUnitWithProrationPriceCadenceSemiAnnual PlanVersionNewParamsAddPricesPriceNewPlanUnitWithProrationPriceCadence = "semi_annual"
	PlanVersionNewParamsAddPricesPriceNewPlanUnitWithProrationPriceCadenceMonthly    PlanVersionNewParamsAddPricesPriceNewPlanUnitWithProrationPriceCadence = "monthly"
	PlanVersionNewParamsAddPricesPriceNewPlanUnitWithProrationPriceCadenceQuarterly  PlanVersionNewParamsAddPricesPriceNewPlanUnitWithProrationPriceCadence = "quarterly"
	PlanVersionNewParamsAddPricesPriceNewPlanUnitWithProrationPriceCadenceOneTime    PlanVersionNewParamsAddPricesPriceNewPlanUnitWithProrationPriceCadence = "one_time"
	PlanVersionNewParamsAddPricesPriceNewPlanUnitWithProrationPriceCadenceCustom     PlanVersionNewParamsAddPricesPriceNewPlanUnitWithProrationPriceCadence = "custom"
)

func (r PlanVersionNewParamsAddPricesPriceNewPlanUnitWithProrationPriceCadence) IsKnown() bool {
	switch r {
	case PlanVersionNewParamsAddPricesPriceNewPlanUnitWithProrationPriceCadenceAnnual, PlanVersionNewParamsAddPricesPriceNewPlanUnitWithProrationPriceCadenceSemiAnnual, PlanVersionNewParamsAddPricesPriceNewPlanUnitWithProrationPriceCadenceMonthly, PlanVersionNewParamsAddPricesPriceNewPlanUnitWithProrationPriceCadenceQuarterly, PlanVersionNewParamsAddPricesPriceNewPlanUnitWithProrationPriceCadenceOneTime, PlanVersionNewParamsAddPricesPriceNewPlanUnitWithProrationPriceCadenceCustom:
		return true
	}
	return false
}

type PlanVersionNewParamsAddPricesPriceNewPlanUnitWithProrationPriceModelType string

const (
	PlanVersionNewParamsAddPricesPriceNewPlanUnitWithProrationPriceModelTypeUnitWithProration PlanVersionNewParamsAddPricesPriceNewPlanUnitWithProrationPriceModelType = "unit_with_proration"
)

func (r PlanVersionNewParamsAddPricesPriceNewPlanUnitWithProrationPriceModelType) IsKnown() bool {
	switch r {
	case PlanVersionNewParamsAddPricesPriceNewPlanUnitWithProrationPriceModelTypeUnitWithProration:
		return true
	}
	return false
}

// For custom cadence: specifies the duration of the billing period in days or
// months.
type PlanVersionNewParamsAddPricesPriceNewPlanUnitWithProrationPriceBillingCycleConfiguration struct {
	// The duration of the billing period.
	Duration param.Field[int64] `json:"duration,required"`
	// The unit of billing period duration.
	DurationUnit param.Field[PlanVersionNewParamsAddPricesPriceNewPlanUnitWithProrationPriceBillingCycleConfigurationDurationUnit] `json:"duration_unit,required"`
}

func (r PlanVersionNewParamsAddPricesPriceNewPlanUnitWithProrationPriceBillingCycleConfiguration) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// The unit of billing period duration.
type PlanVersionNewParamsAddPricesPriceNewPlanUnitWithProrationPriceBillingCycleConfigurationDurationUnit string

const (
	PlanVersionNewParamsAddPricesPriceNewPlanUnitWithProrationPriceBillingCycleConfigurationDurationUnitDay   PlanVersionNewParamsAddPricesPriceNewPlanUnitWithProrationPriceBillingCycleConfigurationDurationUnit = "day"
	PlanVersionNewParamsAddPricesPriceNewPlanUnitWithProrationPriceBillingCycleConfigurationDurationUnitMonth PlanVersionNewParamsAddPricesPriceNewPlanUnitWithProrationPriceBillingCycleConfigurationDurationUnit = "month"
)

func (r PlanVersionNewParamsAddPricesPriceNewPlanUnitWithProrationPriceBillingCycleConfigurationDurationUnit) IsKnown() bool {
	switch r {
	case PlanVersionNewParamsAddPricesPriceNewPlanUnitWithProrationPriceBillingCycleConfigurationDurationUnitDay, PlanVersionNewParamsAddPricesPriceNewPlanUnitWithProrationPriceBillingCycleConfigurationDurationUnitMonth:
		return true
	}
	return false
}

// For dimensional price: specifies a price group and dimension values
type PlanVersionNewParamsAddPricesPriceNewPlanUnitWithProrationPriceDimensionalPriceConfiguration struct {
	// The list of dimension values matching (in order) the dimensions of the price
	// group
	DimensionValues param.Field[[]string] `json:"dimension_values,required"`
	// The id of the dimensional price group to include this price in
	DimensionalPriceGroupID param.Field[string] `json:"dimensional_price_group_id"`
	// The external id of the dimensional price group to include this price in
	ExternalDimensionalPriceGroupID param.Field[string] `json:"external_dimensional_price_group_id"`
}

func (r PlanVersionNewParamsAddPricesPriceNewPlanUnitWithProrationPriceDimensionalPriceConfiguration) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Within each billing cycle, specifies the cadence at which invoices are produced.
// If unspecified, a single invoice is produced per billing cycle.
type PlanVersionNewParamsAddPricesPriceNewPlanUnitWithProrationPriceInvoicingCycleConfiguration struct {
	// The duration of the billing period.
	Duration param.Field[int64] `json:"duration,required"`
	// The unit of billing period duration.
	DurationUnit param.Field[PlanVersionNewParamsAddPricesPriceNewPlanUnitWithProrationPriceInvoicingCycleConfigurationDurationUnit] `json:"duration_unit,required"`
}

func (r PlanVersionNewParamsAddPricesPriceNewPlanUnitWithProrationPriceInvoicingCycleConfiguration) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// The unit of billing period duration.
type PlanVersionNewParamsAddPricesPriceNewPlanUnitWithProrationPriceInvoicingCycleConfigurationDurationUnit string

const (
	PlanVersionNewParamsAddPricesPriceNewPlanUnitWithProrationPriceInvoicingCycleConfigurationDurationUnitDay   PlanVersionNewParamsAddPricesPriceNewPlanUnitWithProrationPriceInvoicingCycleConfigurationDurationUnit = "day"
	PlanVersionNewParamsAddPricesPriceNewPlanUnitWithProrationPriceInvoicingCycleConfigurationDurationUnitMonth PlanVersionNewParamsAddPricesPriceNewPlanUnitWithProrationPriceInvoicingCycleConfigurationDurationUnit = "month"
)

func (r PlanVersionNewParamsAddPricesPriceNewPlanUnitWithProrationPriceInvoicingCycleConfigurationDurationUnit) IsKnown() bool {
	switch r {
	case PlanVersionNewParamsAddPricesPriceNewPlanUnitWithProrationPriceInvoicingCycleConfigurationDurationUnitDay, PlanVersionNewParamsAddPricesPriceNewPlanUnitWithProrationPriceInvoicingCycleConfigurationDurationUnitMonth:
		return true
	}
	return false
}

type PlanVersionNewParamsAddPricesPriceNewPlanGroupedAllocationPrice struct {
	// The cadence to bill for this price on.
	Cadence                 param.Field[PlanVersionNewParamsAddPricesPriceNewPlanGroupedAllocationPriceCadence] `json:"cadence,required"`
	GroupedAllocationConfig param.Field[map[string]interface{}]                                                 `json:"grouped_allocation_config,required"`
	// The id of the item the price will be associated with.
	ItemID    param.Field[string]                                                                   `json:"item_id,required"`
	ModelType param.Field[PlanVersionNewParamsAddPricesPriceNewPlanGroupedAllocationPriceModelType] `json:"model_type,required"`
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
	BillingCycleConfiguration param.Field[PlanVersionNewParamsAddPricesPriceNewPlanGroupedAllocationPriceBillingCycleConfiguration] `json:"billing_cycle_configuration"`
	// The per unit conversion rate of the price currency to the invoicing currency.
	ConversionRate param.Field[float64] `json:"conversion_rate"`
	// An ISO 4217 currency string, or custom pricing unit identifier, in which this
	// price is billed.
	Currency param.Field[string] `json:"currency"`
	// For dimensional price: specifies a price group and dimension values
	DimensionalPriceConfiguration param.Field[PlanVersionNewParamsAddPricesPriceNewPlanGroupedAllocationPriceDimensionalPriceConfiguration] `json:"dimensional_price_configuration"`
	// An alias for the price.
	ExternalPriceID param.Field[string] `json:"external_price_id"`
	// If the Price represents a fixed cost, this represents the quantity of units
	// applied.
	FixedPriceQuantity param.Field[float64] `json:"fixed_price_quantity"`
	// The property used to group this price on an invoice
	InvoiceGroupingKey param.Field[string] `json:"invoice_grouping_key"`
	// Within each billing cycle, specifies the cadence at which invoices are produced.
	// If unspecified, a single invoice is produced per billing cycle.
	InvoicingCycleConfiguration param.Field[PlanVersionNewParamsAddPricesPriceNewPlanGroupedAllocationPriceInvoicingCycleConfiguration] `json:"invoicing_cycle_configuration"`
	// User-specified key/value pairs for the resource. Individual keys can be removed
	// by setting the value to `null`, and the entire metadata mapping can be cleared
	// by setting `metadata` to `null`.
	Metadata param.Field[map[string]string] `json:"metadata"`
}

func (r PlanVersionNewParamsAddPricesPriceNewPlanGroupedAllocationPrice) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r PlanVersionNewParamsAddPricesPriceNewPlanGroupedAllocationPrice) implementsPlanVersionNewParamsAddPricesPriceUnion() {
}

// The cadence to bill for this price on.
type PlanVersionNewParamsAddPricesPriceNewPlanGroupedAllocationPriceCadence string

const (
	PlanVersionNewParamsAddPricesPriceNewPlanGroupedAllocationPriceCadenceAnnual     PlanVersionNewParamsAddPricesPriceNewPlanGroupedAllocationPriceCadence = "annual"
	PlanVersionNewParamsAddPricesPriceNewPlanGroupedAllocationPriceCadenceSemiAnnual PlanVersionNewParamsAddPricesPriceNewPlanGroupedAllocationPriceCadence = "semi_annual"
	PlanVersionNewParamsAddPricesPriceNewPlanGroupedAllocationPriceCadenceMonthly    PlanVersionNewParamsAddPricesPriceNewPlanGroupedAllocationPriceCadence = "monthly"
	PlanVersionNewParamsAddPricesPriceNewPlanGroupedAllocationPriceCadenceQuarterly  PlanVersionNewParamsAddPricesPriceNewPlanGroupedAllocationPriceCadence = "quarterly"
	PlanVersionNewParamsAddPricesPriceNewPlanGroupedAllocationPriceCadenceOneTime    PlanVersionNewParamsAddPricesPriceNewPlanGroupedAllocationPriceCadence = "one_time"
	PlanVersionNewParamsAddPricesPriceNewPlanGroupedAllocationPriceCadenceCustom     PlanVersionNewParamsAddPricesPriceNewPlanGroupedAllocationPriceCadence = "custom"
)

func (r PlanVersionNewParamsAddPricesPriceNewPlanGroupedAllocationPriceCadence) IsKnown() bool {
	switch r {
	case PlanVersionNewParamsAddPricesPriceNewPlanGroupedAllocationPriceCadenceAnnual, PlanVersionNewParamsAddPricesPriceNewPlanGroupedAllocationPriceCadenceSemiAnnual, PlanVersionNewParamsAddPricesPriceNewPlanGroupedAllocationPriceCadenceMonthly, PlanVersionNewParamsAddPricesPriceNewPlanGroupedAllocationPriceCadenceQuarterly, PlanVersionNewParamsAddPricesPriceNewPlanGroupedAllocationPriceCadenceOneTime, PlanVersionNewParamsAddPricesPriceNewPlanGroupedAllocationPriceCadenceCustom:
		return true
	}
	return false
}

type PlanVersionNewParamsAddPricesPriceNewPlanGroupedAllocationPriceModelType string

const (
	PlanVersionNewParamsAddPricesPriceNewPlanGroupedAllocationPriceModelTypeGroupedAllocation PlanVersionNewParamsAddPricesPriceNewPlanGroupedAllocationPriceModelType = "grouped_allocation"
)

func (r PlanVersionNewParamsAddPricesPriceNewPlanGroupedAllocationPriceModelType) IsKnown() bool {
	switch r {
	case PlanVersionNewParamsAddPricesPriceNewPlanGroupedAllocationPriceModelTypeGroupedAllocation:
		return true
	}
	return false
}

// For custom cadence: specifies the duration of the billing period in days or
// months.
type PlanVersionNewParamsAddPricesPriceNewPlanGroupedAllocationPriceBillingCycleConfiguration struct {
	// The duration of the billing period.
	Duration param.Field[int64] `json:"duration,required"`
	// The unit of billing period duration.
	DurationUnit param.Field[PlanVersionNewParamsAddPricesPriceNewPlanGroupedAllocationPriceBillingCycleConfigurationDurationUnit] `json:"duration_unit,required"`
}

func (r PlanVersionNewParamsAddPricesPriceNewPlanGroupedAllocationPriceBillingCycleConfiguration) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// The unit of billing period duration.
type PlanVersionNewParamsAddPricesPriceNewPlanGroupedAllocationPriceBillingCycleConfigurationDurationUnit string

const (
	PlanVersionNewParamsAddPricesPriceNewPlanGroupedAllocationPriceBillingCycleConfigurationDurationUnitDay   PlanVersionNewParamsAddPricesPriceNewPlanGroupedAllocationPriceBillingCycleConfigurationDurationUnit = "day"
	PlanVersionNewParamsAddPricesPriceNewPlanGroupedAllocationPriceBillingCycleConfigurationDurationUnitMonth PlanVersionNewParamsAddPricesPriceNewPlanGroupedAllocationPriceBillingCycleConfigurationDurationUnit = "month"
)

func (r PlanVersionNewParamsAddPricesPriceNewPlanGroupedAllocationPriceBillingCycleConfigurationDurationUnit) IsKnown() bool {
	switch r {
	case PlanVersionNewParamsAddPricesPriceNewPlanGroupedAllocationPriceBillingCycleConfigurationDurationUnitDay, PlanVersionNewParamsAddPricesPriceNewPlanGroupedAllocationPriceBillingCycleConfigurationDurationUnitMonth:
		return true
	}
	return false
}

// For dimensional price: specifies a price group and dimension values
type PlanVersionNewParamsAddPricesPriceNewPlanGroupedAllocationPriceDimensionalPriceConfiguration struct {
	// The list of dimension values matching (in order) the dimensions of the price
	// group
	DimensionValues param.Field[[]string] `json:"dimension_values,required"`
	// The id of the dimensional price group to include this price in
	DimensionalPriceGroupID param.Field[string] `json:"dimensional_price_group_id"`
	// The external id of the dimensional price group to include this price in
	ExternalDimensionalPriceGroupID param.Field[string] `json:"external_dimensional_price_group_id"`
}

func (r PlanVersionNewParamsAddPricesPriceNewPlanGroupedAllocationPriceDimensionalPriceConfiguration) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Within each billing cycle, specifies the cadence at which invoices are produced.
// If unspecified, a single invoice is produced per billing cycle.
type PlanVersionNewParamsAddPricesPriceNewPlanGroupedAllocationPriceInvoicingCycleConfiguration struct {
	// The duration of the billing period.
	Duration param.Field[int64] `json:"duration,required"`
	// The unit of billing period duration.
	DurationUnit param.Field[PlanVersionNewParamsAddPricesPriceNewPlanGroupedAllocationPriceInvoicingCycleConfigurationDurationUnit] `json:"duration_unit,required"`
}

func (r PlanVersionNewParamsAddPricesPriceNewPlanGroupedAllocationPriceInvoicingCycleConfiguration) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// The unit of billing period duration.
type PlanVersionNewParamsAddPricesPriceNewPlanGroupedAllocationPriceInvoicingCycleConfigurationDurationUnit string

const (
	PlanVersionNewParamsAddPricesPriceNewPlanGroupedAllocationPriceInvoicingCycleConfigurationDurationUnitDay   PlanVersionNewParamsAddPricesPriceNewPlanGroupedAllocationPriceInvoicingCycleConfigurationDurationUnit = "day"
	PlanVersionNewParamsAddPricesPriceNewPlanGroupedAllocationPriceInvoicingCycleConfigurationDurationUnitMonth PlanVersionNewParamsAddPricesPriceNewPlanGroupedAllocationPriceInvoicingCycleConfigurationDurationUnit = "month"
)

func (r PlanVersionNewParamsAddPricesPriceNewPlanGroupedAllocationPriceInvoicingCycleConfigurationDurationUnit) IsKnown() bool {
	switch r {
	case PlanVersionNewParamsAddPricesPriceNewPlanGroupedAllocationPriceInvoicingCycleConfigurationDurationUnitDay, PlanVersionNewParamsAddPricesPriceNewPlanGroupedAllocationPriceInvoicingCycleConfigurationDurationUnitMonth:
		return true
	}
	return false
}

type PlanVersionNewParamsAddPricesPriceNewPlanGroupedWithProratedMinimumPrice struct {
	// The cadence to bill for this price on.
	Cadence                          param.Field[PlanVersionNewParamsAddPricesPriceNewPlanGroupedWithProratedMinimumPriceCadence] `json:"cadence,required"`
	GroupedWithProratedMinimumConfig param.Field[map[string]interface{}]                                                          `json:"grouped_with_prorated_minimum_config,required"`
	// The id of the item the price will be associated with.
	ItemID    param.Field[string]                                                                            `json:"item_id,required"`
	ModelType param.Field[PlanVersionNewParamsAddPricesPriceNewPlanGroupedWithProratedMinimumPriceModelType] `json:"model_type,required"`
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
	BillingCycleConfiguration param.Field[PlanVersionNewParamsAddPricesPriceNewPlanGroupedWithProratedMinimumPriceBillingCycleConfiguration] `json:"billing_cycle_configuration"`
	// The per unit conversion rate of the price currency to the invoicing currency.
	ConversionRate param.Field[float64] `json:"conversion_rate"`
	// An ISO 4217 currency string, or custom pricing unit identifier, in which this
	// price is billed.
	Currency param.Field[string] `json:"currency"`
	// For dimensional price: specifies a price group and dimension values
	DimensionalPriceConfiguration param.Field[PlanVersionNewParamsAddPricesPriceNewPlanGroupedWithProratedMinimumPriceDimensionalPriceConfiguration] `json:"dimensional_price_configuration"`
	// An alias for the price.
	ExternalPriceID param.Field[string] `json:"external_price_id"`
	// If the Price represents a fixed cost, this represents the quantity of units
	// applied.
	FixedPriceQuantity param.Field[float64] `json:"fixed_price_quantity"`
	// The property used to group this price on an invoice
	InvoiceGroupingKey param.Field[string] `json:"invoice_grouping_key"`
	// Within each billing cycle, specifies the cadence at which invoices are produced.
	// If unspecified, a single invoice is produced per billing cycle.
	InvoicingCycleConfiguration param.Field[PlanVersionNewParamsAddPricesPriceNewPlanGroupedWithProratedMinimumPriceInvoicingCycleConfiguration] `json:"invoicing_cycle_configuration"`
	// User-specified key/value pairs for the resource. Individual keys can be removed
	// by setting the value to `null`, and the entire metadata mapping can be cleared
	// by setting `metadata` to `null`.
	Metadata param.Field[map[string]string] `json:"metadata"`
}

func (r PlanVersionNewParamsAddPricesPriceNewPlanGroupedWithProratedMinimumPrice) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r PlanVersionNewParamsAddPricesPriceNewPlanGroupedWithProratedMinimumPrice) implementsPlanVersionNewParamsAddPricesPriceUnion() {
}

// The cadence to bill for this price on.
type PlanVersionNewParamsAddPricesPriceNewPlanGroupedWithProratedMinimumPriceCadence string

const (
	PlanVersionNewParamsAddPricesPriceNewPlanGroupedWithProratedMinimumPriceCadenceAnnual     PlanVersionNewParamsAddPricesPriceNewPlanGroupedWithProratedMinimumPriceCadence = "annual"
	PlanVersionNewParamsAddPricesPriceNewPlanGroupedWithProratedMinimumPriceCadenceSemiAnnual PlanVersionNewParamsAddPricesPriceNewPlanGroupedWithProratedMinimumPriceCadence = "semi_annual"
	PlanVersionNewParamsAddPricesPriceNewPlanGroupedWithProratedMinimumPriceCadenceMonthly    PlanVersionNewParamsAddPricesPriceNewPlanGroupedWithProratedMinimumPriceCadence = "monthly"
	PlanVersionNewParamsAddPricesPriceNewPlanGroupedWithProratedMinimumPriceCadenceQuarterly  PlanVersionNewParamsAddPricesPriceNewPlanGroupedWithProratedMinimumPriceCadence = "quarterly"
	PlanVersionNewParamsAddPricesPriceNewPlanGroupedWithProratedMinimumPriceCadenceOneTime    PlanVersionNewParamsAddPricesPriceNewPlanGroupedWithProratedMinimumPriceCadence = "one_time"
	PlanVersionNewParamsAddPricesPriceNewPlanGroupedWithProratedMinimumPriceCadenceCustom     PlanVersionNewParamsAddPricesPriceNewPlanGroupedWithProratedMinimumPriceCadence = "custom"
)

func (r PlanVersionNewParamsAddPricesPriceNewPlanGroupedWithProratedMinimumPriceCadence) IsKnown() bool {
	switch r {
	case PlanVersionNewParamsAddPricesPriceNewPlanGroupedWithProratedMinimumPriceCadenceAnnual, PlanVersionNewParamsAddPricesPriceNewPlanGroupedWithProratedMinimumPriceCadenceSemiAnnual, PlanVersionNewParamsAddPricesPriceNewPlanGroupedWithProratedMinimumPriceCadenceMonthly, PlanVersionNewParamsAddPricesPriceNewPlanGroupedWithProratedMinimumPriceCadenceQuarterly, PlanVersionNewParamsAddPricesPriceNewPlanGroupedWithProratedMinimumPriceCadenceOneTime, PlanVersionNewParamsAddPricesPriceNewPlanGroupedWithProratedMinimumPriceCadenceCustom:
		return true
	}
	return false
}

type PlanVersionNewParamsAddPricesPriceNewPlanGroupedWithProratedMinimumPriceModelType string

const (
	PlanVersionNewParamsAddPricesPriceNewPlanGroupedWithProratedMinimumPriceModelTypeGroupedWithProratedMinimum PlanVersionNewParamsAddPricesPriceNewPlanGroupedWithProratedMinimumPriceModelType = "grouped_with_prorated_minimum"
)

func (r PlanVersionNewParamsAddPricesPriceNewPlanGroupedWithProratedMinimumPriceModelType) IsKnown() bool {
	switch r {
	case PlanVersionNewParamsAddPricesPriceNewPlanGroupedWithProratedMinimumPriceModelTypeGroupedWithProratedMinimum:
		return true
	}
	return false
}

// For custom cadence: specifies the duration of the billing period in days or
// months.
type PlanVersionNewParamsAddPricesPriceNewPlanGroupedWithProratedMinimumPriceBillingCycleConfiguration struct {
	// The duration of the billing period.
	Duration param.Field[int64] `json:"duration,required"`
	// The unit of billing period duration.
	DurationUnit param.Field[PlanVersionNewParamsAddPricesPriceNewPlanGroupedWithProratedMinimumPriceBillingCycleConfigurationDurationUnit] `json:"duration_unit,required"`
}

func (r PlanVersionNewParamsAddPricesPriceNewPlanGroupedWithProratedMinimumPriceBillingCycleConfiguration) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// The unit of billing period duration.
type PlanVersionNewParamsAddPricesPriceNewPlanGroupedWithProratedMinimumPriceBillingCycleConfigurationDurationUnit string

const (
	PlanVersionNewParamsAddPricesPriceNewPlanGroupedWithProratedMinimumPriceBillingCycleConfigurationDurationUnitDay   PlanVersionNewParamsAddPricesPriceNewPlanGroupedWithProratedMinimumPriceBillingCycleConfigurationDurationUnit = "day"
	PlanVersionNewParamsAddPricesPriceNewPlanGroupedWithProratedMinimumPriceBillingCycleConfigurationDurationUnitMonth PlanVersionNewParamsAddPricesPriceNewPlanGroupedWithProratedMinimumPriceBillingCycleConfigurationDurationUnit = "month"
)

func (r PlanVersionNewParamsAddPricesPriceNewPlanGroupedWithProratedMinimumPriceBillingCycleConfigurationDurationUnit) IsKnown() bool {
	switch r {
	case PlanVersionNewParamsAddPricesPriceNewPlanGroupedWithProratedMinimumPriceBillingCycleConfigurationDurationUnitDay, PlanVersionNewParamsAddPricesPriceNewPlanGroupedWithProratedMinimumPriceBillingCycleConfigurationDurationUnitMonth:
		return true
	}
	return false
}

// For dimensional price: specifies a price group and dimension values
type PlanVersionNewParamsAddPricesPriceNewPlanGroupedWithProratedMinimumPriceDimensionalPriceConfiguration struct {
	// The list of dimension values matching (in order) the dimensions of the price
	// group
	DimensionValues param.Field[[]string] `json:"dimension_values,required"`
	// The id of the dimensional price group to include this price in
	DimensionalPriceGroupID param.Field[string] `json:"dimensional_price_group_id"`
	// The external id of the dimensional price group to include this price in
	ExternalDimensionalPriceGroupID param.Field[string] `json:"external_dimensional_price_group_id"`
}

func (r PlanVersionNewParamsAddPricesPriceNewPlanGroupedWithProratedMinimumPriceDimensionalPriceConfiguration) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Within each billing cycle, specifies the cadence at which invoices are produced.
// If unspecified, a single invoice is produced per billing cycle.
type PlanVersionNewParamsAddPricesPriceNewPlanGroupedWithProratedMinimumPriceInvoicingCycleConfiguration struct {
	// The duration of the billing period.
	Duration param.Field[int64] `json:"duration,required"`
	// The unit of billing period duration.
	DurationUnit param.Field[PlanVersionNewParamsAddPricesPriceNewPlanGroupedWithProratedMinimumPriceInvoicingCycleConfigurationDurationUnit] `json:"duration_unit,required"`
}

func (r PlanVersionNewParamsAddPricesPriceNewPlanGroupedWithProratedMinimumPriceInvoicingCycleConfiguration) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// The unit of billing period duration.
type PlanVersionNewParamsAddPricesPriceNewPlanGroupedWithProratedMinimumPriceInvoicingCycleConfigurationDurationUnit string

const (
	PlanVersionNewParamsAddPricesPriceNewPlanGroupedWithProratedMinimumPriceInvoicingCycleConfigurationDurationUnitDay   PlanVersionNewParamsAddPricesPriceNewPlanGroupedWithProratedMinimumPriceInvoicingCycleConfigurationDurationUnit = "day"
	PlanVersionNewParamsAddPricesPriceNewPlanGroupedWithProratedMinimumPriceInvoicingCycleConfigurationDurationUnitMonth PlanVersionNewParamsAddPricesPriceNewPlanGroupedWithProratedMinimumPriceInvoicingCycleConfigurationDurationUnit = "month"
)

func (r PlanVersionNewParamsAddPricesPriceNewPlanGroupedWithProratedMinimumPriceInvoicingCycleConfigurationDurationUnit) IsKnown() bool {
	switch r {
	case PlanVersionNewParamsAddPricesPriceNewPlanGroupedWithProratedMinimumPriceInvoicingCycleConfigurationDurationUnitDay, PlanVersionNewParamsAddPricesPriceNewPlanGroupedWithProratedMinimumPriceInvoicingCycleConfigurationDurationUnitMonth:
		return true
	}
	return false
}

type PlanVersionNewParamsAddPricesPriceNewPlanGroupedWithMeteredMinimumPrice struct {
	// The cadence to bill for this price on.
	Cadence                         param.Field[PlanVersionNewParamsAddPricesPriceNewPlanGroupedWithMeteredMinimumPriceCadence] `json:"cadence,required"`
	GroupedWithMeteredMinimumConfig param.Field[map[string]interface{}]                                                         `json:"grouped_with_metered_minimum_config,required"`
	// The id of the item the price will be associated with.
	ItemID    param.Field[string]                                                                           `json:"item_id,required"`
	ModelType param.Field[PlanVersionNewParamsAddPricesPriceNewPlanGroupedWithMeteredMinimumPriceModelType] `json:"model_type,required"`
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
	BillingCycleConfiguration param.Field[PlanVersionNewParamsAddPricesPriceNewPlanGroupedWithMeteredMinimumPriceBillingCycleConfiguration] `json:"billing_cycle_configuration"`
	// The per unit conversion rate of the price currency to the invoicing currency.
	ConversionRate param.Field[float64] `json:"conversion_rate"`
	// An ISO 4217 currency string, or custom pricing unit identifier, in which this
	// price is billed.
	Currency param.Field[string] `json:"currency"`
	// For dimensional price: specifies a price group and dimension values
	DimensionalPriceConfiguration param.Field[PlanVersionNewParamsAddPricesPriceNewPlanGroupedWithMeteredMinimumPriceDimensionalPriceConfiguration] `json:"dimensional_price_configuration"`
	// An alias for the price.
	ExternalPriceID param.Field[string] `json:"external_price_id"`
	// If the Price represents a fixed cost, this represents the quantity of units
	// applied.
	FixedPriceQuantity param.Field[float64] `json:"fixed_price_quantity"`
	// The property used to group this price on an invoice
	InvoiceGroupingKey param.Field[string] `json:"invoice_grouping_key"`
	// Within each billing cycle, specifies the cadence at which invoices are produced.
	// If unspecified, a single invoice is produced per billing cycle.
	InvoicingCycleConfiguration param.Field[PlanVersionNewParamsAddPricesPriceNewPlanGroupedWithMeteredMinimumPriceInvoicingCycleConfiguration] `json:"invoicing_cycle_configuration"`
	// User-specified key/value pairs for the resource. Individual keys can be removed
	// by setting the value to `null`, and the entire metadata mapping can be cleared
	// by setting `metadata` to `null`.
	Metadata param.Field[map[string]string] `json:"metadata"`
}

func (r PlanVersionNewParamsAddPricesPriceNewPlanGroupedWithMeteredMinimumPrice) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r PlanVersionNewParamsAddPricesPriceNewPlanGroupedWithMeteredMinimumPrice) implementsPlanVersionNewParamsAddPricesPriceUnion() {
}

// The cadence to bill for this price on.
type PlanVersionNewParamsAddPricesPriceNewPlanGroupedWithMeteredMinimumPriceCadence string

const (
	PlanVersionNewParamsAddPricesPriceNewPlanGroupedWithMeteredMinimumPriceCadenceAnnual     PlanVersionNewParamsAddPricesPriceNewPlanGroupedWithMeteredMinimumPriceCadence = "annual"
	PlanVersionNewParamsAddPricesPriceNewPlanGroupedWithMeteredMinimumPriceCadenceSemiAnnual PlanVersionNewParamsAddPricesPriceNewPlanGroupedWithMeteredMinimumPriceCadence = "semi_annual"
	PlanVersionNewParamsAddPricesPriceNewPlanGroupedWithMeteredMinimumPriceCadenceMonthly    PlanVersionNewParamsAddPricesPriceNewPlanGroupedWithMeteredMinimumPriceCadence = "monthly"
	PlanVersionNewParamsAddPricesPriceNewPlanGroupedWithMeteredMinimumPriceCadenceQuarterly  PlanVersionNewParamsAddPricesPriceNewPlanGroupedWithMeteredMinimumPriceCadence = "quarterly"
	PlanVersionNewParamsAddPricesPriceNewPlanGroupedWithMeteredMinimumPriceCadenceOneTime    PlanVersionNewParamsAddPricesPriceNewPlanGroupedWithMeteredMinimumPriceCadence = "one_time"
	PlanVersionNewParamsAddPricesPriceNewPlanGroupedWithMeteredMinimumPriceCadenceCustom     PlanVersionNewParamsAddPricesPriceNewPlanGroupedWithMeteredMinimumPriceCadence = "custom"
)

func (r PlanVersionNewParamsAddPricesPriceNewPlanGroupedWithMeteredMinimumPriceCadence) IsKnown() bool {
	switch r {
	case PlanVersionNewParamsAddPricesPriceNewPlanGroupedWithMeteredMinimumPriceCadenceAnnual, PlanVersionNewParamsAddPricesPriceNewPlanGroupedWithMeteredMinimumPriceCadenceSemiAnnual, PlanVersionNewParamsAddPricesPriceNewPlanGroupedWithMeteredMinimumPriceCadenceMonthly, PlanVersionNewParamsAddPricesPriceNewPlanGroupedWithMeteredMinimumPriceCadenceQuarterly, PlanVersionNewParamsAddPricesPriceNewPlanGroupedWithMeteredMinimumPriceCadenceOneTime, PlanVersionNewParamsAddPricesPriceNewPlanGroupedWithMeteredMinimumPriceCadenceCustom:
		return true
	}
	return false
}

type PlanVersionNewParamsAddPricesPriceNewPlanGroupedWithMeteredMinimumPriceModelType string

const (
	PlanVersionNewParamsAddPricesPriceNewPlanGroupedWithMeteredMinimumPriceModelTypeGroupedWithMeteredMinimum PlanVersionNewParamsAddPricesPriceNewPlanGroupedWithMeteredMinimumPriceModelType = "grouped_with_metered_minimum"
)

func (r PlanVersionNewParamsAddPricesPriceNewPlanGroupedWithMeteredMinimumPriceModelType) IsKnown() bool {
	switch r {
	case PlanVersionNewParamsAddPricesPriceNewPlanGroupedWithMeteredMinimumPriceModelTypeGroupedWithMeteredMinimum:
		return true
	}
	return false
}

// For custom cadence: specifies the duration of the billing period in days or
// months.
type PlanVersionNewParamsAddPricesPriceNewPlanGroupedWithMeteredMinimumPriceBillingCycleConfiguration struct {
	// The duration of the billing period.
	Duration param.Field[int64] `json:"duration,required"`
	// The unit of billing period duration.
	DurationUnit param.Field[PlanVersionNewParamsAddPricesPriceNewPlanGroupedWithMeteredMinimumPriceBillingCycleConfigurationDurationUnit] `json:"duration_unit,required"`
}

func (r PlanVersionNewParamsAddPricesPriceNewPlanGroupedWithMeteredMinimumPriceBillingCycleConfiguration) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// The unit of billing period duration.
type PlanVersionNewParamsAddPricesPriceNewPlanGroupedWithMeteredMinimumPriceBillingCycleConfigurationDurationUnit string

const (
	PlanVersionNewParamsAddPricesPriceNewPlanGroupedWithMeteredMinimumPriceBillingCycleConfigurationDurationUnitDay   PlanVersionNewParamsAddPricesPriceNewPlanGroupedWithMeteredMinimumPriceBillingCycleConfigurationDurationUnit = "day"
	PlanVersionNewParamsAddPricesPriceNewPlanGroupedWithMeteredMinimumPriceBillingCycleConfigurationDurationUnitMonth PlanVersionNewParamsAddPricesPriceNewPlanGroupedWithMeteredMinimumPriceBillingCycleConfigurationDurationUnit = "month"
)

func (r PlanVersionNewParamsAddPricesPriceNewPlanGroupedWithMeteredMinimumPriceBillingCycleConfigurationDurationUnit) IsKnown() bool {
	switch r {
	case PlanVersionNewParamsAddPricesPriceNewPlanGroupedWithMeteredMinimumPriceBillingCycleConfigurationDurationUnitDay, PlanVersionNewParamsAddPricesPriceNewPlanGroupedWithMeteredMinimumPriceBillingCycleConfigurationDurationUnitMonth:
		return true
	}
	return false
}

// For dimensional price: specifies a price group and dimension values
type PlanVersionNewParamsAddPricesPriceNewPlanGroupedWithMeteredMinimumPriceDimensionalPriceConfiguration struct {
	// The list of dimension values matching (in order) the dimensions of the price
	// group
	DimensionValues param.Field[[]string] `json:"dimension_values,required"`
	// The id of the dimensional price group to include this price in
	DimensionalPriceGroupID param.Field[string] `json:"dimensional_price_group_id"`
	// The external id of the dimensional price group to include this price in
	ExternalDimensionalPriceGroupID param.Field[string] `json:"external_dimensional_price_group_id"`
}

func (r PlanVersionNewParamsAddPricesPriceNewPlanGroupedWithMeteredMinimumPriceDimensionalPriceConfiguration) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Within each billing cycle, specifies the cadence at which invoices are produced.
// If unspecified, a single invoice is produced per billing cycle.
type PlanVersionNewParamsAddPricesPriceNewPlanGroupedWithMeteredMinimumPriceInvoicingCycleConfiguration struct {
	// The duration of the billing period.
	Duration param.Field[int64] `json:"duration,required"`
	// The unit of billing period duration.
	DurationUnit param.Field[PlanVersionNewParamsAddPricesPriceNewPlanGroupedWithMeteredMinimumPriceInvoicingCycleConfigurationDurationUnit] `json:"duration_unit,required"`
}

func (r PlanVersionNewParamsAddPricesPriceNewPlanGroupedWithMeteredMinimumPriceInvoicingCycleConfiguration) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// The unit of billing period duration.
type PlanVersionNewParamsAddPricesPriceNewPlanGroupedWithMeteredMinimumPriceInvoicingCycleConfigurationDurationUnit string

const (
	PlanVersionNewParamsAddPricesPriceNewPlanGroupedWithMeteredMinimumPriceInvoicingCycleConfigurationDurationUnitDay   PlanVersionNewParamsAddPricesPriceNewPlanGroupedWithMeteredMinimumPriceInvoicingCycleConfigurationDurationUnit = "day"
	PlanVersionNewParamsAddPricesPriceNewPlanGroupedWithMeteredMinimumPriceInvoicingCycleConfigurationDurationUnitMonth PlanVersionNewParamsAddPricesPriceNewPlanGroupedWithMeteredMinimumPriceInvoicingCycleConfigurationDurationUnit = "month"
)

func (r PlanVersionNewParamsAddPricesPriceNewPlanGroupedWithMeteredMinimumPriceInvoicingCycleConfigurationDurationUnit) IsKnown() bool {
	switch r {
	case PlanVersionNewParamsAddPricesPriceNewPlanGroupedWithMeteredMinimumPriceInvoicingCycleConfigurationDurationUnitDay, PlanVersionNewParamsAddPricesPriceNewPlanGroupedWithMeteredMinimumPriceInvoicingCycleConfigurationDurationUnitMonth:
		return true
	}
	return false
}

type PlanVersionNewParamsAddPricesPriceNewPlanMatrixWithDisplayNamePrice struct {
	// The cadence to bill for this price on.
	Cadence param.Field[PlanVersionNewParamsAddPricesPriceNewPlanMatrixWithDisplayNamePriceCadence] `json:"cadence,required"`
	// The id of the item the price will be associated with.
	ItemID                      param.Field[string]                                                                       `json:"item_id,required"`
	MatrixWithDisplayNameConfig param.Field[map[string]interface{}]                                                       `json:"matrix_with_display_name_config,required"`
	ModelType                   param.Field[PlanVersionNewParamsAddPricesPriceNewPlanMatrixWithDisplayNamePriceModelType] `json:"model_type,required"`
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
	BillingCycleConfiguration param.Field[PlanVersionNewParamsAddPricesPriceNewPlanMatrixWithDisplayNamePriceBillingCycleConfiguration] `json:"billing_cycle_configuration"`
	// The per unit conversion rate of the price currency to the invoicing currency.
	ConversionRate param.Field[float64] `json:"conversion_rate"`
	// An ISO 4217 currency string, or custom pricing unit identifier, in which this
	// price is billed.
	Currency param.Field[string] `json:"currency"`
	// For dimensional price: specifies a price group and dimension values
	DimensionalPriceConfiguration param.Field[PlanVersionNewParamsAddPricesPriceNewPlanMatrixWithDisplayNamePriceDimensionalPriceConfiguration] `json:"dimensional_price_configuration"`
	// An alias for the price.
	ExternalPriceID param.Field[string] `json:"external_price_id"`
	// If the Price represents a fixed cost, this represents the quantity of units
	// applied.
	FixedPriceQuantity param.Field[float64] `json:"fixed_price_quantity"`
	// The property used to group this price on an invoice
	InvoiceGroupingKey param.Field[string] `json:"invoice_grouping_key"`
	// Within each billing cycle, specifies the cadence at which invoices are produced.
	// If unspecified, a single invoice is produced per billing cycle.
	InvoicingCycleConfiguration param.Field[PlanVersionNewParamsAddPricesPriceNewPlanMatrixWithDisplayNamePriceInvoicingCycleConfiguration] `json:"invoicing_cycle_configuration"`
	// User-specified key/value pairs for the resource. Individual keys can be removed
	// by setting the value to `null`, and the entire metadata mapping can be cleared
	// by setting `metadata` to `null`.
	Metadata param.Field[map[string]string] `json:"metadata"`
}

func (r PlanVersionNewParamsAddPricesPriceNewPlanMatrixWithDisplayNamePrice) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r PlanVersionNewParamsAddPricesPriceNewPlanMatrixWithDisplayNamePrice) implementsPlanVersionNewParamsAddPricesPriceUnion() {
}

// The cadence to bill for this price on.
type PlanVersionNewParamsAddPricesPriceNewPlanMatrixWithDisplayNamePriceCadence string

const (
	PlanVersionNewParamsAddPricesPriceNewPlanMatrixWithDisplayNamePriceCadenceAnnual     PlanVersionNewParamsAddPricesPriceNewPlanMatrixWithDisplayNamePriceCadence = "annual"
	PlanVersionNewParamsAddPricesPriceNewPlanMatrixWithDisplayNamePriceCadenceSemiAnnual PlanVersionNewParamsAddPricesPriceNewPlanMatrixWithDisplayNamePriceCadence = "semi_annual"
	PlanVersionNewParamsAddPricesPriceNewPlanMatrixWithDisplayNamePriceCadenceMonthly    PlanVersionNewParamsAddPricesPriceNewPlanMatrixWithDisplayNamePriceCadence = "monthly"
	PlanVersionNewParamsAddPricesPriceNewPlanMatrixWithDisplayNamePriceCadenceQuarterly  PlanVersionNewParamsAddPricesPriceNewPlanMatrixWithDisplayNamePriceCadence = "quarterly"
	PlanVersionNewParamsAddPricesPriceNewPlanMatrixWithDisplayNamePriceCadenceOneTime    PlanVersionNewParamsAddPricesPriceNewPlanMatrixWithDisplayNamePriceCadence = "one_time"
	PlanVersionNewParamsAddPricesPriceNewPlanMatrixWithDisplayNamePriceCadenceCustom     PlanVersionNewParamsAddPricesPriceNewPlanMatrixWithDisplayNamePriceCadence = "custom"
)

func (r PlanVersionNewParamsAddPricesPriceNewPlanMatrixWithDisplayNamePriceCadence) IsKnown() bool {
	switch r {
	case PlanVersionNewParamsAddPricesPriceNewPlanMatrixWithDisplayNamePriceCadenceAnnual, PlanVersionNewParamsAddPricesPriceNewPlanMatrixWithDisplayNamePriceCadenceSemiAnnual, PlanVersionNewParamsAddPricesPriceNewPlanMatrixWithDisplayNamePriceCadenceMonthly, PlanVersionNewParamsAddPricesPriceNewPlanMatrixWithDisplayNamePriceCadenceQuarterly, PlanVersionNewParamsAddPricesPriceNewPlanMatrixWithDisplayNamePriceCadenceOneTime, PlanVersionNewParamsAddPricesPriceNewPlanMatrixWithDisplayNamePriceCadenceCustom:
		return true
	}
	return false
}

type PlanVersionNewParamsAddPricesPriceNewPlanMatrixWithDisplayNamePriceModelType string

const (
	PlanVersionNewParamsAddPricesPriceNewPlanMatrixWithDisplayNamePriceModelTypeMatrixWithDisplayName PlanVersionNewParamsAddPricesPriceNewPlanMatrixWithDisplayNamePriceModelType = "matrix_with_display_name"
)

func (r PlanVersionNewParamsAddPricesPriceNewPlanMatrixWithDisplayNamePriceModelType) IsKnown() bool {
	switch r {
	case PlanVersionNewParamsAddPricesPriceNewPlanMatrixWithDisplayNamePriceModelTypeMatrixWithDisplayName:
		return true
	}
	return false
}

// For custom cadence: specifies the duration of the billing period in days or
// months.
type PlanVersionNewParamsAddPricesPriceNewPlanMatrixWithDisplayNamePriceBillingCycleConfiguration struct {
	// The duration of the billing period.
	Duration param.Field[int64] `json:"duration,required"`
	// The unit of billing period duration.
	DurationUnit param.Field[PlanVersionNewParamsAddPricesPriceNewPlanMatrixWithDisplayNamePriceBillingCycleConfigurationDurationUnit] `json:"duration_unit,required"`
}

func (r PlanVersionNewParamsAddPricesPriceNewPlanMatrixWithDisplayNamePriceBillingCycleConfiguration) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// The unit of billing period duration.
type PlanVersionNewParamsAddPricesPriceNewPlanMatrixWithDisplayNamePriceBillingCycleConfigurationDurationUnit string

const (
	PlanVersionNewParamsAddPricesPriceNewPlanMatrixWithDisplayNamePriceBillingCycleConfigurationDurationUnitDay   PlanVersionNewParamsAddPricesPriceNewPlanMatrixWithDisplayNamePriceBillingCycleConfigurationDurationUnit = "day"
	PlanVersionNewParamsAddPricesPriceNewPlanMatrixWithDisplayNamePriceBillingCycleConfigurationDurationUnitMonth PlanVersionNewParamsAddPricesPriceNewPlanMatrixWithDisplayNamePriceBillingCycleConfigurationDurationUnit = "month"
)

func (r PlanVersionNewParamsAddPricesPriceNewPlanMatrixWithDisplayNamePriceBillingCycleConfigurationDurationUnit) IsKnown() bool {
	switch r {
	case PlanVersionNewParamsAddPricesPriceNewPlanMatrixWithDisplayNamePriceBillingCycleConfigurationDurationUnitDay, PlanVersionNewParamsAddPricesPriceNewPlanMatrixWithDisplayNamePriceBillingCycleConfigurationDurationUnitMonth:
		return true
	}
	return false
}

// For dimensional price: specifies a price group and dimension values
type PlanVersionNewParamsAddPricesPriceNewPlanMatrixWithDisplayNamePriceDimensionalPriceConfiguration struct {
	// The list of dimension values matching (in order) the dimensions of the price
	// group
	DimensionValues param.Field[[]string] `json:"dimension_values,required"`
	// The id of the dimensional price group to include this price in
	DimensionalPriceGroupID param.Field[string] `json:"dimensional_price_group_id"`
	// The external id of the dimensional price group to include this price in
	ExternalDimensionalPriceGroupID param.Field[string] `json:"external_dimensional_price_group_id"`
}

func (r PlanVersionNewParamsAddPricesPriceNewPlanMatrixWithDisplayNamePriceDimensionalPriceConfiguration) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Within each billing cycle, specifies the cadence at which invoices are produced.
// If unspecified, a single invoice is produced per billing cycle.
type PlanVersionNewParamsAddPricesPriceNewPlanMatrixWithDisplayNamePriceInvoicingCycleConfiguration struct {
	// The duration of the billing period.
	Duration param.Field[int64] `json:"duration,required"`
	// The unit of billing period duration.
	DurationUnit param.Field[PlanVersionNewParamsAddPricesPriceNewPlanMatrixWithDisplayNamePriceInvoicingCycleConfigurationDurationUnit] `json:"duration_unit,required"`
}

func (r PlanVersionNewParamsAddPricesPriceNewPlanMatrixWithDisplayNamePriceInvoicingCycleConfiguration) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// The unit of billing period duration.
type PlanVersionNewParamsAddPricesPriceNewPlanMatrixWithDisplayNamePriceInvoicingCycleConfigurationDurationUnit string

const (
	PlanVersionNewParamsAddPricesPriceNewPlanMatrixWithDisplayNamePriceInvoicingCycleConfigurationDurationUnitDay   PlanVersionNewParamsAddPricesPriceNewPlanMatrixWithDisplayNamePriceInvoicingCycleConfigurationDurationUnit = "day"
	PlanVersionNewParamsAddPricesPriceNewPlanMatrixWithDisplayNamePriceInvoicingCycleConfigurationDurationUnitMonth PlanVersionNewParamsAddPricesPriceNewPlanMatrixWithDisplayNamePriceInvoicingCycleConfigurationDurationUnit = "month"
)

func (r PlanVersionNewParamsAddPricesPriceNewPlanMatrixWithDisplayNamePriceInvoicingCycleConfigurationDurationUnit) IsKnown() bool {
	switch r {
	case PlanVersionNewParamsAddPricesPriceNewPlanMatrixWithDisplayNamePriceInvoicingCycleConfigurationDurationUnitDay, PlanVersionNewParamsAddPricesPriceNewPlanMatrixWithDisplayNamePriceInvoicingCycleConfigurationDurationUnitMonth:
		return true
	}
	return false
}

type PlanVersionNewParamsAddPricesPriceNewPlanBulkWithProrationPrice struct {
	BulkWithProrationConfig param.Field[map[string]interface{}] `json:"bulk_with_proration_config,required"`
	// The cadence to bill for this price on.
	Cadence param.Field[PlanVersionNewParamsAddPricesPriceNewPlanBulkWithProrationPriceCadence] `json:"cadence,required"`
	// The id of the item the price will be associated with.
	ItemID    param.Field[string]                                                                   `json:"item_id,required"`
	ModelType param.Field[PlanVersionNewParamsAddPricesPriceNewPlanBulkWithProrationPriceModelType] `json:"model_type,required"`
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
	BillingCycleConfiguration param.Field[PlanVersionNewParamsAddPricesPriceNewPlanBulkWithProrationPriceBillingCycleConfiguration] `json:"billing_cycle_configuration"`
	// The per unit conversion rate of the price currency to the invoicing currency.
	ConversionRate param.Field[float64] `json:"conversion_rate"`
	// An ISO 4217 currency string, or custom pricing unit identifier, in which this
	// price is billed.
	Currency param.Field[string] `json:"currency"`
	// For dimensional price: specifies a price group and dimension values
	DimensionalPriceConfiguration param.Field[PlanVersionNewParamsAddPricesPriceNewPlanBulkWithProrationPriceDimensionalPriceConfiguration] `json:"dimensional_price_configuration"`
	// An alias for the price.
	ExternalPriceID param.Field[string] `json:"external_price_id"`
	// If the Price represents a fixed cost, this represents the quantity of units
	// applied.
	FixedPriceQuantity param.Field[float64] `json:"fixed_price_quantity"`
	// The property used to group this price on an invoice
	InvoiceGroupingKey param.Field[string] `json:"invoice_grouping_key"`
	// Within each billing cycle, specifies the cadence at which invoices are produced.
	// If unspecified, a single invoice is produced per billing cycle.
	InvoicingCycleConfiguration param.Field[PlanVersionNewParamsAddPricesPriceNewPlanBulkWithProrationPriceInvoicingCycleConfiguration] `json:"invoicing_cycle_configuration"`
	// User-specified key/value pairs for the resource. Individual keys can be removed
	// by setting the value to `null`, and the entire metadata mapping can be cleared
	// by setting `metadata` to `null`.
	Metadata param.Field[map[string]string] `json:"metadata"`
}

func (r PlanVersionNewParamsAddPricesPriceNewPlanBulkWithProrationPrice) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r PlanVersionNewParamsAddPricesPriceNewPlanBulkWithProrationPrice) implementsPlanVersionNewParamsAddPricesPriceUnion() {
}

// The cadence to bill for this price on.
type PlanVersionNewParamsAddPricesPriceNewPlanBulkWithProrationPriceCadence string

const (
	PlanVersionNewParamsAddPricesPriceNewPlanBulkWithProrationPriceCadenceAnnual     PlanVersionNewParamsAddPricesPriceNewPlanBulkWithProrationPriceCadence = "annual"
	PlanVersionNewParamsAddPricesPriceNewPlanBulkWithProrationPriceCadenceSemiAnnual PlanVersionNewParamsAddPricesPriceNewPlanBulkWithProrationPriceCadence = "semi_annual"
	PlanVersionNewParamsAddPricesPriceNewPlanBulkWithProrationPriceCadenceMonthly    PlanVersionNewParamsAddPricesPriceNewPlanBulkWithProrationPriceCadence = "monthly"
	PlanVersionNewParamsAddPricesPriceNewPlanBulkWithProrationPriceCadenceQuarterly  PlanVersionNewParamsAddPricesPriceNewPlanBulkWithProrationPriceCadence = "quarterly"
	PlanVersionNewParamsAddPricesPriceNewPlanBulkWithProrationPriceCadenceOneTime    PlanVersionNewParamsAddPricesPriceNewPlanBulkWithProrationPriceCadence = "one_time"
	PlanVersionNewParamsAddPricesPriceNewPlanBulkWithProrationPriceCadenceCustom     PlanVersionNewParamsAddPricesPriceNewPlanBulkWithProrationPriceCadence = "custom"
)

func (r PlanVersionNewParamsAddPricesPriceNewPlanBulkWithProrationPriceCadence) IsKnown() bool {
	switch r {
	case PlanVersionNewParamsAddPricesPriceNewPlanBulkWithProrationPriceCadenceAnnual, PlanVersionNewParamsAddPricesPriceNewPlanBulkWithProrationPriceCadenceSemiAnnual, PlanVersionNewParamsAddPricesPriceNewPlanBulkWithProrationPriceCadenceMonthly, PlanVersionNewParamsAddPricesPriceNewPlanBulkWithProrationPriceCadenceQuarterly, PlanVersionNewParamsAddPricesPriceNewPlanBulkWithProrationPriceCadenceOneTime, PlanVersionNewParamsAddPricesPriceNewPlanBulkWithProrationPriceCadenceCustom:
		return true
	}
	return false
}

type PlanVersionNewParamsAddPricesPriceNewPlanBulkWithProrationPriceModelType string

const (
	PlanVersionNewParamsAddPricesPriceNewPlanBulkWithProrationPriceModelTypeBulkWithProration PlanVersionNewParamsAddPricesPriceNewPlanBulkWithProrationPriceModelType = "bulk_with_proration"
)

func (r PlanVersionNewParamsAddPricesPriceNewPlanBulkWithProrationPriceModelType) IsKnown() bool {
	switch r {
	case PlanVersionNewParamsAddPricesPriceNewPlanBulkWithProrationPriceModelTypeBulkWithProration:
		return true
	}
	return false
}

// For custom cadence: specifies the duration of the billing period in days or
// months.
type PlanVersionNewParamsAddPricesPriceNewPlanBulkWithProrationPriceBillingCycleConfiguration struct {
	// The duration of the billing period.
	Duration param.Field[int64] `json:"duration,required"`
	// The unit of billing period duration.
	DurationUnit param.Field[PlanVersionNewParamsAddPricesPriceNewPlanBulkWithProrationPriceBillingCycleConfigurationDurationUnit] `json:"duration_unit,required"`
}

func (r PlanVersionNewParamsAddPricesPriceNewPlanBulkWithProrationPriceBillingCycleConfiguration) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// The unit of billing period duration.
type PlanVersionNewParamsAddPricesPriceNewPlanBulkWithProrationPriceBillingCycleConfigurationDurationUnit string

const (
	PlanVersionNewParamsAddPricesPriceNewPlanBulkWithProrationPriceBillingCycleConfigurationDurationUnitDay   PlanVersionNewParamsAddPricesPriceNewPlanBulkWithProrationPriceBillingCycleConfigurationDurationUnit = "day"
	PlanVersionNewParamsAddPricesPriceNewPlanBulkWithProrationPriceBillingCycleConfigurationDurationUnitMonth PlanVersionNewParamsAddPricesPriceNewPlanBulkWithProrationPriceBillingCycleConfigurationDurationUnit = "month"
)

func (r PlanVersionNewParamsAddPricesPriceNewPlanBulkWithProrationPriceBillingCycleConfigurationDurationUnit) IsKnown() bool {
	switch r {
	case PlanVersionNewParamsAddPricesPriceNewPlanBulkWithProrationPriceBillingCycleConfigurationDurationUnitDay, PlanVersionNewParamsAddPricesPriceNewPlanBulkWithProrationPriceBillingCycleConfigurationDurationUnitMonth:
		return true
	}
	return false
}

// For dimensional price: specifies a price group and dimension values
type PlanVersionNewParamsAddPricesPriceNewPlanBulkWithProrationPriceDimensionalPriceConfiguration struct {
	// The list of dimension values matching (in order) the dimensions of the price
	// group
	DimensionValues param.Field[[]string] `json:"dimension_values,required"`
	// The id of the dimensional price group to include this price in
	DimensionalPriceGroupID param.Field[string] `json:"dimensional_price_group_id"`
	// The external id of the dimensional price group to include this price in
	ExternalDimensionalPriceGroupID param.Field[string] `json:"external_dimensional_price_group_id"`
}

func (r PlanVersionNewParamsAddPricesPriceNewPlanBulkWithProrationPriceDimensionalPriceConfiguration) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Within each billing cycle, specifies the cadence at which invoices are produced.
// If unspecified, a single invoice is produced per billing cycle.
type PlanVersionNewParamsAddPricesPriceNewPlanBulkWithProrationPriceInvoicingCycleConfiguration struct {
	// The duration of the billing period.
	Duration param.Field[int64] `json:"duration,required"`
	// The unit of billing period duration.
	DurationUnit param.Field[PlanVersionNewParamsAddPricesPriceNewPlanBulkWithProrationPriceInvoicingCycleConfigurationDurationUnit] `json:"duration_unit,required"`
}

func (r PlanVersionNewParamsAddPricesPriceNewPlanBulkWithProrationPriceInvoicingCycleConfiguration) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// The unit of billing period duration.
type PlanVersionNewParamsAddPricesPriceNewPlanBulkWithProrationPriceInvoicingCycleConfigurationDurationUnit string

const (
	PlanVersionNewParamsAddPricesPriceNewPlanBulkWithProrationPriceInvoicingCycleConfigurationDurationUnitDay   PlanVersionNewParamsAddPricesPriceNewPlanBulkWithProrationPriceInvoicingCycleConfigurationDurationUnit = "day"
	PlanVersionNewParamsAddPricesPriceNewPlanBulkWithProrationPriceInvoicingCycleConfigurationDurationUnitMonth PlanVersionNewParamsAddPricesPriceNewPlanBulkWithProrationPriceInvoicingCycleConfigurationDurationUnit = "month"
)

func (r PlanVersionNewParamsAddPricesPriceNewPlanBulkWithProrationPriceInvoicingCycleConfigurationDurationUnit) IsKnown() bool {
	switch r {
	case PlanVersionNewParamsAddPricesPriceNewPlanBulkWithProrationPriceInvoicingCycleConfigurationDurationUnitDay, PlanVersionNewParamsAddPricesPriceNewPlanBulkWithProrationPriceInvoicingCycleConfigurationDurationUnitMonth:
		return true
	}
	return false
}

type PlanVersionNewParamsAddPricesPriceNewPlanGroupedTieredPackagePrice struct {
	// The cadence to bill for this price on.
	Cadence                    param.Field[PlanVersionNewParamsAddPricesPriceNewPlanGroupedTieredPackagePriceCadence] `json:"cadence,required"`
	GroupedTieredPackageConfig param.Field[map[string]interface{}]                                                    `json:"grouped_tiered_package_config,required"`
	// The id of the item the price will be associated with.
	ItemID    param.Field[string]                                                                      `json:"item_id,required"`
	ModelType param.Field[PlanVersionNewParamsAddPricesPriceNewPlanGroupedTieredPackagePriceModelType] `json:"model_type,required"`
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
	BillingCycleConfiguration param.Field[PlanVersionNewParamsAddPricesPriceNewPlanGroupedTieredPackagePriceBillingCycleConfiguration] `json:"billing_cycle_configuration"`
	// The per unit conversion rate of the price currency to the invoicing currency.
	ConversionRate param.Field[float64] `json:"conversion_rate"`
	// An ISO 4217 currency string, or custom pricing unit identifier, in which this
	// price is billed.
	Currency param.Field[string] `json:"currency"`
	// For dimensional price: specifies a price group and dimension values
	DimensionalPriceConfiguration param.Field[PlanVersionNewParamsAddPricesPriceNewPlanGroupedTieredPackagePriceDimensionalPriceConfiguration] `json:"dimensional_price_configuration"`
	// An alias for the price.
	ExternalPriceID param.Field[string] `json:"external_price_id"`
	// If the Price represents a fixed cost, this represents the quantity of units
	// applied.
	FixedPriceQuantity param.Field[float64] `json:"fixed_price_quantity"`
	// The property used to group this price on an invoice
	InvoiceGroupingKey param.Field[string] `json:"invoice_grouping_key"`
	// Within each billing cycle, specifies the cadence at which invoices are produced.
	// If unspecified, a single invoice is produced per billing cycle.
	InvoicingCycleConfiguration param.Field[PlanVersionNewParamsAddPricesPriceNewPlanGroupedTieredPackagePriceInvoicingCycleConfiguration] `json:"invoicing_cycle_configuration"`
	// User-specified key/value pairs for the resource. Individual keys can be removed
	// by setting the value to `null`, and the entire metadata mapping can be cleared
	// by setting `metadata` to `null`.
	Metadata param.Field[map[string]string] `json:"metadata"`
}

func (r PlanVersionNewParamsAddPricesPriceNewPlanGroupedTieredPackagePrice) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r PlanVersionNewParamsAddPricesPriceNewPlanGroupedTieredPackagePrice) implementsPlanVersionNewParamsAddPricesPriceUnion() {
}

// The cadence to bill for this price on.
type PlanVersionNewParamsAddPricesPriceNewPlanGroupedTieredPackagePriceCadence string

const (
	PlanVersionNewParamsAddPricesPriceNewPlanGroupedTieredPackagePriceCadenceAnnual     PlanVersionNewParamsAddPricesPriceNewPlanGroupedTieredPackagePriceCadence = "annual"
	PlanVersionNewParamsAddPricesPriceNewPlanGroupedTieredPackagePriceCadenceSemiAnnual PlanVersionNewParamsAddPricesPriceNewPlanGroupedTieredPackagePriceCadence = "semi_annual"
	PlanVersionNewParamsAddPricesPriceNewPlanGroupedTieredPackagePriceCadenceMonthly    PlanVersionNewParamsAddPricesPriceNewPlanGroupedTieredPackagePriceCadence = "monthly"
	PlanVersionNewParamsAddPricesPriceNewPlanGroupedTieredPackagePriceCadenceQuarterly  PlanVersionNewParamsAddPricesPriceNewPlanGroupedTieredPackagePriceCadence = "quarterly"
	PlanVersionNewParamsAddPricesPriceNewPlanGroupedTieredPackagePriceCadenceOneTime    PlanVersionNewParamsAddPricesPriceNewPlanGroupedTieredPackagePriceCadence = "one_time"
	PlanVersionNewParamsAddPricesPriceNewPlanGroupedTieredPackagePriceCadenceCustom     PlanVersionNewParamsAddPricesPriceNewPlanGroupedTieredPackagePriceCadence = "custom"
)

func (r PlanVersionNewParamsAddPricesPriceNewPlanGroupedTieredPackagePriceCadence) IsKnown() bool {
	switch r {
	case PlanVersionNewParamsAddPricesPriceNewPlanGroupedTieredPackagePriceCadenceAnnual, PlanVersionNewParamsAddPricesPriceNewPlanGroupedTieredPackagePriceCadenceSemiAnnual, PlanVersionNewParamsAddPricesPriceNewPlanGroupedTieredPackagePriceCadenceMonthly, PlanVersionNewParamsAddPricesPriceNewPlanGroupedTieredPackagePriceCadenceQuarterly, PlanVersionNewParamsAddPricesPriceNewPlanGroupedTieredPackagePriceCadenceOneTime, PlanVersionNewParamsAddPricesPriceNewPlanGroupedTieredPackagePriceCadenceCustom:
		return true
	}
	return false
}

type PlanVersionNewParamsAddPricesPriceNewPlanGroupedTieredPackagePriceModelType string

const (
	PlanVersionNewParamsAddPricesPriceNewPlanGroupedTieredPackagePriceModelTypeGroupedTieredPackage PlanVersionNewParamsAddPricesPriceNewPlanGroupedTieredPackagePriceModelType = "grouped_tiered_package"
)

func (r PlanVersionNewParamsAddPricesPriceNewPlanGroupedTieredPackagePriceModelType) IsKnown() bool {
	switch r {
	case PlanVersionNewParamsAddPricesPriceNewPlanGroupedTieredPackagePriceModelTypeGroupedTieredPackage:
		return true
	}
	return false
}

// For custom cadence: specifies the duration of the billing period in days or
// months.
type PlanVersionNewParamsAddPricesPriceNewPlanGroupedTieredPackagePriceBillingCycleConfiguration struct {
	// The duration of the billing period.
	Duration param.Field[int64] `json:"duration,required"`
	// The unit of billing period duration.
	DurationUnit param.Field[PlanVersionNewParamsAddPricesPriceNewPlanGroupedTieredPackagePriceBillingCycleConfigurationDurationUnit] `json:"duration_unit,required"`
}

func (r PlanVersionNewParamsAddPricesPriceNewPlanGroupedTieredPackagePriceBillingCycleConfiguration) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// The unit of billing period duration.
type PlanVersionNewParamsAddPricesPriceNewPlanGroupedTieredPackagePriceBillingCycleConfigurationDurationUnit string

const (
	PlanVersionNewParamsAddPricesPriceNewPlanGroupedTieredPackagePriceBillingCycleConfigurationDurationUnitDay   PlanVersionNewParamsAddPricesPriceNewPlanGroupedTieredPackagePriceBillingCycleConfigurationDurationUnit = "day"
	PlanVersionNewParamsAddPricesPriceNewPlanGroupedTieredPackagePriceBillingCycleConfigurationDurationUnitMonth PlanVersionNewParamsAddPricesPriceNewPlanGroupedTieredPackagePriceBillingCycleConfigurationDurationUnit = "month"
)

func (r PlanVersionNewParamsAddPricesPriceNewPlanGroupedTieredPackagePriceBillingCycleConfigurationDurationUnit) IsKnown() bool {
	switch r {
	case PlanVersionNewParamsAddPricesPriceNewPlanGroupedTieredPackagePriceBillingCycleConfigurationDurationUnitDay, PlanVersionNewParamsAddPricesPriceNewPlanGroupedTieredPackagePriceBillingCycleConfigurationDurationUnitMonth:
		return true
	}
	return false
}

// For dimensional price: specifies a price group and dimension values
type PlanVersionNewParamsAddPricesPriceNewPlanGroupedTieredPackagePriceDimensionalPriceConfiguration struct {
	// The list of dimension values matching (in order) the dimensions of the price
	// group
	DimensionValues param.Field[[]string] `json:"dimension_values,required"`
	// The id of the dimensional price group to include this price in
	DimensionalPriceGroupID param.Field[string] `json:"dimensional_price_group_id"`
	// The external id of the dimensional price group to include this price in
	ExternalDimensionalPriceGroupID param.Field[string] `json:"external_dimensional_price_group_id"`
}

func (r PlanVersionNewParamsAddPricesPriceNewPlanGroupedTieredPackagePriceDimensionalPriceConfiguration) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Within each billing cycle, specifies the cadence at which invoices are produced.
// If unspecified, a single invoice is produced per billing cycle.
type PlanVersionNewParamsAddPricesPriceNewPlanGroupedTieredPackagePriceInvoicingCycleConfiguration struct {
	// The duration of the billing period.
	Duration param.Field[int64] `json:"duration,required"`
	// The unit of billing period duration.
	DurationUnit param.Field[PlanVersionNewParamsAddPricesPriceNewPlanGroupedTieredPackagePriceInvoicingCycleConfigurationDurationUnit] `json:"duration_unit,required"`
}

func (r PlanVersionNewParamsAddPricesPriceNewPlanGroupedTieredPackagePriceInvoicingCycleConfiguration) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// The unit of billing period duration.
type PlanVersionNewParamsAddPricesPriceNewPlanGroupedTieredPackagePriceInvoicingCycleConfigurationDurationUnit string

const (
	PlanVersionNewParamsAddPricesPriceNewPlanGroupedTieredPackagePriceInvoicingCycleConfigurationDurationUnitDay   PlanVersionNewParamsAddPricesPriceNewPlanGroupedTieredPackagePriceInvoicingCycleConfigurationDurationUnit = "day"
	PlanVersionNewParamsAddPricesPriceNewPlanGroupedTieredPackagePriceInvoicingCycleConfigurationDurationUnitMonth PlanVersionNewParamsAddPricesPriceNewPlanGroupedTieredPackagePriceInvoicingCycleConfigurationDurationUnit = "month"
)

func (r PlanVersionNewParamsAddPricesPriceNewPlanGroupedTieredPackagePriceInvoicingCycleConfigurationDurationUnit) IsKnown() bool {
	switch r {
	case PlanVersionNewParamsAddPricesPriceNewPlanGroupedTieredPackagePriceInvoicingCycleConfigurationDurationUnitDay, PlanVersionNewParamsAddPricesPriceNewPlanGroupedTieredPackagePriceInvoicingCycleConfigurationDurationUnitMonth:
		return true
	}
	return false
}

type PlanVersionNewParamsAddPricesPriceNewPlanMaxGroupTieredPackagePrice struct {
	// The cadence to bill for this price on.
	Cadence param.Field[PlanVersionNewParamsAddPricesPriceNewPlanMaxGroupTieredPackagePriceCadence] `json:"cadence,required"`
	// The id of the item the price will be associated with.
	ItemID                      param.Field[string]                                                                       `json:"item_id,required"`
	MaxGroupTieredPackageConfig param.Field[map[string]interface{}]                                                       `json:"max_group_tiered_package_config,required"`
	ModelType                   param.Field[PlanVersionNewParamsAddPricesPriceNewPlanMaxGroupTieredPackagePriceModelType] `json:"model_type,required"`
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
	BillingCycleConfiguration param.Field[PlanVersionNewParamsAddPricesPriceNewPlanMaxGroupTieredPackagePriceBillingCycleConfiguration] `json:"billing_cycle_configuration"`
	// The per unit conversion rate of the price currency to the invoicing currency.
	ConversionRate param.Field[float64] `json:"conversion_rate"`
	// An ISO 4217 currency string, or custom pricing unit identifier, in which this
	// price is billed.
	Currency param.Field[string] `json:"currency"`
	// For dimensional price: specifies a price group and dimension values
	DimensionalPriceConfiguration param.Field[PlanVersionNewParamsAddPricesPriceNewPlanMaxGroupTieredPackagePriceDimensionalPriceConfiguration] `json:"dimensional_price_configuration"`
	// An alias for the price.
	ExternalPriceID param.Field[string] `json:"external_price_id"`
	// If the Price represents a fixed cost, this represents the quantity of units
	// applied.
	FixedPriceQuantity param.Field[float64] `json:"fixed_price_quantity"`
	// The property used to group this price on an invoice
	InvoiceGroupingKey param.Field[string] `json:"invoice_grouping_key"`
	// Within each billing cycle, specifies the cadence at which invoices are produced.
	// If unspecified, a single invoice is produced per billing cycle.
	InvoicingCycleConfiguration param.Field[PlanVersionNewParamsAddPricesPriceNewPlanMaxGroupTieredPackagePriceInvoicingCycleConfiguration] `json:"invoicing_cycle_configuration"`
	// User-specified key/value pairs for the resource. Individual keys can be removed
	// by setting the value to `null`, and the entire metadata mapping can be cleared
	// by setting `metadata` to `null`.
	Metadata param.Field[map[string]string] `json:"metadata"`
}

func (r PlanVersionNewParamsAddPricesPriceNewPlanMaxGroupTieredPackagePrice) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r PlanVersionNewParamsAddPricesPriceNewPlanMaxGroupTieredPackagePrice) implementsPlanVersionNewParamsAddPricesPriceUnion() {
}

// The cadence to bill for this price on.
type PlanVersionNewParamsAddPricesPriceNewPlanMaxGroupTieredPackagePriceCadence string

const (
	PlanVersionNewParamsAddPricesPriceNewPlanMaxGroupTieredPackagePriceCadenceAnnual     PlanVersionNewParamsAddPricesPriceNewPlanMaxGroupTieredPackagePriceCadence = "annual"
	PlanVersionNewParamsAddPricesPriceNewPlanMaxGroupTieredPackagePriceCadenceSemiAnnual PlanVersionNewParamsAddPricesPriceNewPlanMaxGroupTieredPackagePriceCadence = "semi_annual"
	PlanVersionNewParamsAddPricesPriceNewPlanMaxGroupTieredPackagePriceCadenceMonthly    PlanVersionNewParamsAddPricesPriceNewPlanMaxGroupTieredPackagePriceCadence = "monthly"
	PlanVersionNewParamsAddPricesPriceNewPlanMaxGroupTieredPackagePriceCadenceQuarterly  PlanVersionNewParamsAddPricesPriceNewPlanMaxGroupTieredPackagePriceCadence = "quarterly"
	PlanVersionNewParamsAddPricesPriceNewPlanMaxGroupTieredPackagePriceCadenceOneTime    PlanVersionNewParamsAddPricesPriceNewPlanMaxGroupTieredPackagePriceCadence = "one_time"
	PlanVersionNewParamsAddPricesPriceNewPlanMaxGroupTieredPackagePriceCadenceCustom     PlanVersionNewParamsAddPricesPriceNewPlanMaxGroupTieredPackagePriceCadence = "custom"
)

func (r PlanVersionNewParamsAddPricesPriceNewPlanMaxGroupTieredPackagePriceCadence) IsKnown() bool {
	switch r {
	case PlanVersionNewParamsAddPricesPriceNewPlanMaxGroupTieredPackagePriceCadenceAnnual, PlanVersionNewParamsAddPricesPriceNewPlanMaxGroupTieredPackagePriceCadenceSemiAnnual, PlanVersionNewParamsAddPricesPriceNewPlanMaxGroupTieredPackagePriceCadenceMonthly, PlanVersionNewParamsAddPricesPriceNewPlanMaxGroupTieredPackagePriceCadenceQuarterly, PlanVersionNewParamsAddPricesPriceNewPlanMaxGroupTieredPackagePriceCadenceOneTime, PlanVersionNewParamsAddPricesPriceNewPlanMaxGroupTieredPackagePriceCadenceCustom:
		return true
	}
	return false
}

type PlanVersionNewParamsAddPricesPriceNewPlanMaxGroupTieredPackagePriceModelType string

const (
	PlanVersionNewParamsAddPricesPriceNewPlanMaxGroupTieredPackagePriceModelTypeMaxGroupTieredPackage PlanVersionNewParamsAddPricesPriceNewPlanMaxGroupTieredPackagePriceModelType = "max_group_tiered_package"
)

func (r PlanVersionNewParamsAddPricesPriceNewPlanMaxGroupTieredPackagePriceModelType) IsKnown() bool {
	switch r {
	case PlanVersionNewParamsAddPricesPriceNewPlanMaxGroupTieredPackagePriceModelTypeMaxGroupTieredPackage:
		return true
	}
	return false
}

// For custom cadence: specifies the duration of the billing period in days or
// months.
type PlanVersionNewParamsAddPricesPriceNewPlanMaxGroupTieredPackagePriceBillingCycleConfiguration struct {
	// The duration of the billing period.
	Duration param.Field[int64] `json:"duration,required"`
	// The unit of billing period duration.
	DurationUnit param.Field[PlanVersionNewParamsAddPricesPriceNewPlanMaxGroupTieredPackagePriceBillingCycleConfigurationDurationUnit] `json:"duration_unit,required"`
}

func (r PlanVersionNewParamsAddPricesPriceNewPlanMaxGroupTieredPackagePriceBillingCycleConfiguration) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// The unit of billing period duration.
type PlanVersionNewParamsAddPricesPriceNewPlanMaxGroupTieredPackagePriceBillingCycleConfigurationDurationUnit string

const (
	PlanVersionNewParamsAddPricesPriceNewPlanMaxGroupTieredPackagePriceBillingCycleConfigurationDurationUnitDay   PlanVersionNewParamsAddPricesPriceNewPlanMaxGroupTieredPackagePriceBillingCycleConfigurationDurationUnit = "day"
	PlanVersionNewParamsAddPricesPriceNewPlanMaxGroupTieredPackagePriceBillingCycleConfigurationDurationUnitMonth PlanVersionNewParamsAddPricesPriceNewPlanMaxGroupTieredPackagePriceBillingCycleConfigurationDurationUnit = "month"
)

func (r PlanVersionNewParamsAddPricesPriceNewPlanMaxGroupTieredPackagePriceBillingCycleConfigurationDurationUnit) IsKnown() bool {
	switch r {
	case PlanVersionNewParamsAddPricesPriceNewPlanMaxGroupTieredPackagePriceBillingCycleConfigurationDurationUnitDay, PlanVersionNewParamsAddPricesPriceNewPlanMaxGroupTieredPackagePriceBillingCycleConfigurationDurationUnitMonth:
		return true
	}
	return false
}

// For dimensional price: specifies a price group and dimension values
type PlanVersionNewParamsAddPricesPriceNewPlanMaxGroupTieredPackagePriceDimensionalPriceConfiguration struct {
	// The list of dimension values matching (in order) the dimensions of the price
	// group
	DimensionValues param.Field[[]string] `json:"dimension_values,required"`
	// The id of the dimensional price group to include this price in
	DimensionalPriceGroupID param.Field[string] `json:"dimensional_price_group_id"`
	// The external id of the dimensional price group to include this price in
	ExternalDimensionalPriceGroupID param.Field[string] `json:"external_dimensional_price_group_id"`
}

func (r PlanVersionNewParamsAddPricesPriceNewPlanMaxGroupTieredPackagePriceDimensionalPriceConfiguration) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Within each billing cycle, specifies the cadence at which invoices are produced.
// If unspecified, a single invoice is produced per billing cycle.
type PlanVersionNewParamsAddPricesPriceNewPlanMaxGroupTieredPackagePriceInvoicingCycleConfiguration struct {
	// The duration of the billing period.
	Duration param.Field[int64] `json:"duration,required"`
	// The unit of billing period duration.
	DurationUnit param.Field[PlanVersionNewParamsAddPricesPriceNewPlanMaxGroupTieredPackagePriceInvoicingCycleConfigurationDurationUnit] `json:"duration_unit,required"`
}

func (r PlanVersionNewParamsAddPricesPriceNewPlanMaxGroupTieredPackagePriceInvoicingCycleConfiguration) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// The unit of billing period duration.
type PlanVersionNewParamsAddPricesPriceNewPlanMaxGroupTieredPackagePriceInvoicingCycleConfigurationDurationUnit string

const (
	PlanVersionNewParamsAddPricesPriceNewPlanMaxGroupTieredPackagePriceInvoicingCycleConfigurationDurationUnitDay   PlanVersionNewParamsAddPricesPriceNewPlanMaxGroupTieredPackagePriceInvoicingCycleConfigurationDurationUnit = "day"
	PlanVersionNewParamsAddPricesPriceNewPlanMaxGroupTieredPackagePriceInvoicingCycleConfigurationDurationUnitMonth PlanVersionNewParamsAddPricesPriceNewPlanMaxGroupTieredPackagePriceInvoicingCycleConfigurationDurationUnit = "month"
)

func (r PlanVersionNewParamsAddPricesPriceNewPlanMaxGroupTieredPackagePriceInvoicingCycleConfigurationDurationUnit) IsKnown() bool {
	switch r {
	case PlanVersionNewParamsAddPricesPriceNewPlanMaxGroupTieredPackagePriceInvoicingCycleConfigurationDurationUnitDay, PlanVersionNewParamsAddPricesPriceNewPlanMaxGroupTieredPackagePriceInvoicingCycleConfigurationDurationUnitMonth:
		return true
	}
	return false
}

type PlanVersionNewParamsAddPricesPriceNewPlanScalableMatrixWithUnitPricingPrice struct {
	// The cadence to bill for this price on.
	Cadence param.Field[PlanVersionNewParamsAddPricesPriceNewPlanScalableMatrixWithUnitPricingPriceCadence] `json:"cadence,required"`
	// The id of the item the price will be associated with.
	ItemID    param.Field[string]                                                                               `json:"item_id,required"`
	ModelType param.Field[PlanVersionNewParamsAddPricesPriceNewPlanScalableMatrixWithUnitPricingPriceModelType] `json:"model_type,required"`
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
	BillingCycleConfiguration param.Field[PlanVersionNewParamsAddPricesPriceNewPlanScalableMatrixWithUnitPricingPriceBillingCycleConfiguration] `json:"billing_cycle_configuration"`
	// The per unit conversion rate of the price currency to the invoicing currency.
	ConversionRate param.Field[float64] `json:"conversion_rate"`
	// An ISO 4217 currency string, or custom pricing unit identifier, in which this
	// price is billed.
	Currency param.Field[string] `json:"currency"`
	// For dimensional price: specifies a price group and dimension values
	DimensionalPriceConfiguration param.Field[PlanVersionNewParamsAddPricesPriceNewPlanScalableMatrixWithUnitPricingPriceDimensionalPriceConfiguration] `json:"dimensional_price_configuration"`
	// An alias for the price.
	ExternalPriceID param.Field[string] `json:"external_price_id"`
	// If the Price represents a fixed cost, this represents the quantity of units
	// applied.
	FixedPriceQuantity param.Field[float64] `json:"fixed_price_quantity"`
	// The property used to group this price on an invoice
	InvoiceGroupingKey param.Field[string] `json:"invoice_grouping_key"`
	// Within each billing cycle, specifies the cadence at which invoices are produced.
	// If unspecified, a single invoice is produced per billing cycle.
	InvoicingCycleConfiguration param.Field[PlanVersionNewParamsAddPricesPriceNewPlanScalableMatrixWithUnitPricingPriceInvoicingCycleConfiguration] `json:"invoicing_cycle_configuration"`
	// User-specified key/value pairs for the resource. Individual keys can be removed
	// by setting the value to `null`, and the entire metadata mapping can be cleared
	// by setting `metadata` to `null`.
	Metadata param.Field[map[string]string] `json:"metadata"`
}

func (r PlanVersionNewParamsAddPricesPriceNewPlanScalableMatrixWithUnitPricingPrice) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r PlanVersionNewParamsAddPricesPriceNewPlanScalableMatrixWithUnitPricingPrice) implementsPlanVersionNewParamsAddPricesPriceUnion() {
}

// The cadence to bill for this price on.
type PlanVersionNewParamsAddPricesPriceNewPlanScalableMatrixWithUnitPricingPriceCadence string

const (
	PlanVersionNewParamsAddPricesPriceNewPlanScalableMatrixWithUnitPricingPriceCadenceAnnual     PlanVersionNewParamsAddPricesPriceNewPlanScalableMatrixWithUnitPricingPriceCadence = "annual"
	PlanVersionNewParamsAddPricesPriceNewPlanScalableMatrixWithUnitPricingPriceCadenceSemiAnnual PlanVersionNewParamsAddPricesPriceNewPlanScalableMatrixWithUnitPricingPriceCadence = "semi_annual"
	PlanVersionNewParamsAddPricesPriceNewPlanScalableMatrixWithUnitPricingPriceCadenceMonthly    PlanVersionNewParamsAddPricesPriceNewPlanScalableMatrixWithUnitPricingPriceCadence = "monthly"
	PlanVersionNewParamsAddPricesPriceNewPlanScalableMatrixWithUnitPricingPriceCadenceQuarterly  PlanVersionNewParamsAddPricesPriceNewPlanScalableMatrixWithUnitPricingPriceCadence = "quarterly"
	PlanVersionNewParamsAddPricesPriceNewPlanScalableMatrixWithUnitPricingPriceCadenceOneTime    PlanVersionNewParamsAddPricesPriceNewPlanScalableMatrixWithUnitPricingPriceCadence = "one_time"
	PlanVersionNewParamsAddPricesPriceNewPlanScalableMatrixWithUnitPricingPriceCadenceCustom     PlanVersionNewParamsAddPricesPriceNewPlanScalableMatrixWithUnitPricingPriceCadence = "custom"
)

func (r PlanVersionNewParamsAddPricesPriceNewPlanScalableMatrixWithUnitPricingPriceCadence) IsKnown() bool {
	switch r {
	case PlanVersionNewParamsAddPricesPriceNewPlanScalableMatrixWithUnitPricingPriceCadenceAnnual, PlanVersionNewParamsAddPricesPriceNewPlanScalableMatrixWithUnitPricingPriceCadenceSemiAnnual, PlanVersionNewParamsAddPricesPriceNewPlanScalableMatrixWithUnitPricingPriceCadenceMonthly, PlanVersionNewParamsAddPricesPriceNewPlanScalableMatrixWithUnitPricingPriceCadenceQuarterly, PlanVersionNewParamsAddPricesPriceNewPlanScalableMatrixWithUnitPricingPriceCadenceOneTime, PlanVersionNewParamsAddPricesPriceNewPlanScalableMatrixWithUnitPricingPriceCadenceCustom:
		return true
	}
	return false
}

type PlanVersionNewParamsAddPricesPriceNewPlanScalableMatrixWithUnitPricingPriceModelType string

const (
	PlanVersionNewParamsAddPricesPriceNewPlanScalableMatrixWithUnitPricingPriceModelTypeScalableMatrixWithUnitPricing PlanVersionNewParamsAddPricesPriceNewPlanScalableMatrixWithUnitPricingPriceModelType = "scalable_matrix_with_unit_pricing"
)

func (r PlanVersionNewParamsAddPricesPriceNewPlanScalableMatrixWithUnitPricingPriceModelType) IsKnown() bool {
	switch r {
	case PlanVersionNewParamsAddPricesPriceNewPlanScalableMatrixWithUnitPricingPriceModelTypeScalableMatrixWithUnitPricing:
		return true
	}
	return false
}

// For custom cadence: specifies the duration of the billing period in days or
// months.
type PlanVersionNewParamsAddPricesPriceNewPlanScalableMatrixWithUnitPricingPriceBillingCycleConfiguration struct {
	// The duration of the billing period.
	Duration param.Field[int64] `json:"duration,required"`
	// The unit of billing period duration.
	DurationUnit param.Field[PlanVersionNewParamsAddPricesPriceNewPlanScalableMatrixWithUnitPricingPriceBillingCycleConfigurationDurationUnit] `json:"duration_unit,required"`
}

func (r PlanVersionNewParamsAddPricesPriceNewPlanScalableMatrixWithUnitPricingPriceBillingCycleConfiguration) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// The unit of billing period duration.
type PlanVersionNewParamsAddPricesPriceNewPlanScalableMatrixWithUnitPricingPriceBillingCycleConfigurationDurationUnit string

const (
	PlanVersionNewParamsAddPricesPriceNewPlanScalableMatrixWithUnitPricingPriceBillingCycleConfigurationDurationUnitDay   PlanVersionNewParamsAddPricesPriceNewPlanScalableMatrixWithUnitPricingPriceBillingCycleConfigurationDurationUnit = "day"
	PlanVersionNewParamsAddPricesPriceNewPlanScalableMatrixWithUnitPricingPriceBillingCycleConfigurationDurationUnitMonth PlanVersionNewParamsAddPricesPriceNewPlanScalableMatrixWithUnitPricingPriceBillingCycleConfigurationDurationUnit = "month"
)

func (r PlanVersionNewParamsAddPricesPriceNewPlanScalableMatrixWithUnitPricingPriceBillingCycleConfigurationDurationUnit) IsKnown() bool {
	switch r {
	case PlanVersionNewParamsAddPricesPriceNewPlanScalableMatrixWithUnitPricingPriceBillingCycleConfigurationDurationUnitDay, PlanVersionNewParamsAddPricesPriceNewPlanScalableMatrixWithUnitPricingPriceBillingCycleConfigurationDurationUnitMonth:
		return true
	}
	return false
}

// For dimensional price: specifies a price group and dimension values
type PlanVersionNewParamsAddPricesPriceNewPlanScalableMatrixWithUnitPricingPriceDimensionalPriceConfiguration struct {
	// The list of dimension values matching (in order) the dimensions of the price
	// group
	DimensionValues param.Field[[]string] `json:"dimension_values,required"`
	// The id of the dimensional price group to include this price in
	DimensionalPriceGroupID param.Field[string] `json:"dimensional_price_group_id"`
	// The external id of the dimensional price group to include this price in
	ExternalDimensionalPriceGroupID param.Field[string] `json:"external_dimensional_price_group_id"`
}

func (r PlanVersionNewParamsAddPricesPriceNewPlanScalableMatrixWithUnitPricingPriceDimensionalPriceConfiguration) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Within each billing cycle, specifies the cadence at which invoices are produced.
// If unspecified, a single invoice is produced per billing cycle.
type PlanVersionNewParamsAddPricesPriceNewPlanScalableMatrixWithUnitPricingPriceInvoicingCycleConfiguration struct {
	// The duration of the billing period.
	Duration param.Field[int64] `json:"duration,required"`
	// The unit of billing period duration.
	DurationUnit param.Field[PlanVersionNewParamsAddPricesPriceNewPlanScalableMatrixWithUnitPricingPriceInvoicingCycleConfigurationDurationUnit] `json:"duration_unit,required"`
}

func (r PlanVersionNewParamsAddPricesPriceNewPlanScalableMatrixWithUnitPricingPriceInvoicingCycleConfiguration) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// The unit of billing period duration.
type PlanVersionNewParamsAddPricesPriceNewPlanScalableMatrixWithUnitPricingPriceInvoicingCycleConfigurationDurationUnit string

const (
	PlanVersionNewParamsAddPricesPriceNewPlanScalableMatrixWithUnitPricingPriceInvoicingCycleConfigurationDurationUnitDay   PlanVersionNewParamsAddPricesPriceNewPlanScalableMatrixWithUnitPricingPriceInvoicingCycleConfigurationDurationUnit = "day"
	PlanVersionNewParamsAddPricesPriceNewPlanScalableMatrixWithUnitPricingPriceInvoicingCycleConfigurationDurationUnitMonth PlanVersionNewParamsAddPricesPriceNewPlanScalableMatrixWithUnitPricingPriceInvoicingCycleConfigurationDurationUnit = "month"
)

func (r PlanVersionNewParamsAddPricesPriceNewPlanScalableMatrixWithUnitPricingPriceInvoicingCycleConfigurationDurationUnit) IsKnown() bool {
	switch r {
	case PlanVersionNewParamsAddPricesPriceNewPlanScalableMatrixWithUnitPricingPriceInvoicingCycleConfigurationDurationUnitDay, PlanVersionNewParamsAddPricesPriceNewPlanScalableMatrixWithUnitPricingPriceInvoicingCycleConfigurationDurationUnitMonth:
		return true
	}
	return false
}

type PlanVersionNewParamsAddPricesPriceNewPlanScalableMatrixWithTieredPricingPrice struct {
	// The cadence to bill for this price on.
	Cadence param.Field[PlanVersionNewParamsAddPricesPriceNewPlanScalableMatrixWithTieredPricingPriceCadence] `json:"cadence,required"`
	// The id of the item the price will be associated with.
	ItemID    param.Field[string]                                                                                 `json:"item_id,required"`
	ModelType param.Field[PlanVersionNewParamsAddPricesPriceNewPlanScalableMatrixWithTieredPricingPriceModelType] `json:"model_type,required"`
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
	BillingCycleConfiguration param.Field[PlanVersionNewParamsAddPricesPriceNewPlanScalableMatrixWithTieredPricingPriceBillingCycleConfiguration] `json:"billing_cycle_configuration"`
	// The per unit conversion rate of the price currency to the invoicing currency.
	ConversionRate param.Field[float64] `json:"conversion_rate"`
	// An ISO 4217 currency string, or custom pricing unit identifier, in which this
	// price is billed.
	Currency param.Field[string] `json:"currency"`
	// For dimensional price: specifies a price group and dimension values
	DimensionalPriceConfiguration param.Field[PlanVersionNewParamsAddPricesPriceNewPlanScalableMatrixWithTieredPricingPriceDimensionalPriceConfiguration] `json:"dimensional_price_configuration"`
	// An alias for the price.
	ExternalPriceID param.Field[string] `json:"external_price_id"`
	// If the Price represents a fixed cost, this represents the quantity of units
	// applied.
	FixedPriceQuantity param.Field[float64] `json:"fixed_price_quantity"`
	// The property used to group this price on an invoice
	InvoiceGroupingKey param.Field[string] `json:"invoice_grouping_key"`
	// Within each billing cycle, specifies the cadence at which invoices are produced.
	// If unspecified, a single invoice is produced per billing cycle.
	InvoicingCycleConfiguration param.Field[PlanVersionNewParamsAddPricesPriceNewPlanScalableMatrixWithTieredPricingPriceInvoicingCycleConfiguration] `json:"invoicing_cycle_configuration"`
	// User-specified key/value pairs for the resource. Individual keys can be removed
	// by setting the value to `null`, and the entire metadata mapping can be cleared
	// by setting `metadata` to `null`.
	Metadata param.Field[map[string]string] `json:"metadata"`
}

func (r PlanVersionNewParamsAddPricesPriceNewPlanScalableMatrixWithTieredPricingPrice) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r PlanVersionNewParamsAddPricesPriceNewPlanScalableMatrixWithTieredPricingPrice) implementsPlanVersionNewParamsAddPricesPriceUnion() {
}

// The cadence to bill for this price on.
type PlanVersionNewParamsAddPricesPriceNewPlanScalableMatrixWithTieredPricingPriceCadence string

const (
	PlanVersionNewParamsAddPricesPriceNewPlanScalableMatrixWithTieredPricingPriceCadenceAnnual     PlanVersionNewParamsAddPricesPriceNewPlanScalableMatrixWithTieredPricingPriceCadence = "annual"
	PlanVersionNewParamsAddPricesPriceNewPlanScalableMatrixWithTieredPricingPriceCadenceSemiAnnual PlanVersionNewParamsAddPricesPriceNewPlanScalableMatrixWithTieredPricingPriceCadence = "semi_annual"
	PlanVersionNewParamsAddPricesPriceNewPlanScalableMatrixWithTieredPricingPriceCadenceMonthly    PlanVersionNewParamsAddPricesPriceNewPlanScalableMatrixWithTieredPricingPriceCadence = "monthly"
	PlanVersionNewParamsAddPricesPriceNewPlanScalableMatrixWithTieredPricingPriceCadenceQuarterly  PlanVersionNewParamsAddPricesPriceNewPlanScalableMatrixWithTieredPricingPriceCadence = "quarterly"
	PlanVersionNewParamsAddPricesPriceNewPlanScalableMatrixWithTieredPricingPriceCadenceOneTime    PlanVersionNewParamsAddPricesPriceNewPlanScalableMatrixWithTieredPricingPriceCadence = "one_time"
	PlanVersionNewParamsAddPricesPriceNewPlanScalableMatrixWithTieredPricingPriceCadenceCustom     PlanVersionNewParamsAddPricesPriceNewPlanScalableMatrixWithTieredPricingPriceCadence = "custom"
)

func (r PlanVersionNewParamsAddPricesPriceNewPlanScalableMatrixWithTieredPricingPriceCadence) IsKnown() bool {
	switch r {
	case PlanVersionNewParamsAddPricesPriceNewPlanScalableMatrixWithTieredPricingPriceCadenceAnnual, PlanVersionNewParamsAddPricesPriceNewPlanScalableMatrixWithTieredPricingPriceCadenceSemiAnnual, PlanVersionNewParamsAddPricesPriceNewPlanScalableMatrixWithTieredPricingPriceCadenceMonthly, PlanVersionNewParamsAddPricesPriceNewPlanScalableMatrixWithTieredPricingPriceCadenceQuarterly, PlanVersionNewParamsAddPricesPriceNewPlanScalableMatrixWithTieredPricingPriceCadenceOneTime, PlanVersionNewParamsAddPricesPriceNewPlanScalableMatrixWithTieredPricingPriceCadenceCustom:
		return true
	}
	return false
}

type PlanVersionNewParamsAddPricesPriceNewPlanScalableMatrixWithTieredPricingPriceModelType string

const (
	PlanVersionNewParamsAddPricesPriceNewPlanScalableMatrixWithTieredPricingPriceModelTypeScalableMatrixWithTieredPricing PlanVersionNewParamsAddPricesPriceNewPlanScalableMatrixWithTieredPricingPriceModelType = "scalable_matrix_with_tiered_pricing"
)

func (r PlanVersionNewParamsAddPricesPriceNewPlanScalableMatrixWithTieredPricingPriceModelType) IsKnown() bool {
	switch r {
	case PlanVersionNewParamsAddPricesPriceNewPlanScalableMatrixWithTieredPricingPriceModelTypeScalableMatrixWithTieredPricing:
		return true
	}
	return false
}

// For custom cadence: specifies the duration of the billing period in days or
// months.
type PlanVersionNewParamsAddPricesPriceNewPlanScalableMatrixWithTieredPricingPriceBillingCycleConfiguration struct {
	// The duration of the billing period.
	Duration param.Field[int64] `json:"duration,required"`
	// The unit of billing period duration.
	DurationUnit param.Field[PlanVersionNewParamsAddPricesPriceNewPlanScalableMatrixWithTieredPricingPriceBillingCycleConfigurationDurationUnit] `json:"duration_unit,required"`
}

func (r PlanVersionNewParamsAddPricesPriceNewPlanScalableMatrixWithTieredPricingPriceBillingCycleConfiguration) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// The unit of billing period duration.
type PlanVersionNewParamsAddPricesPriceNewPlanScalableMatrixWithTieredPricingPriceBillingCycleConfigurationDurationUnit string

const (
	PlanVersionNewParamsAddPricesPriceNewPlanScalableMatrixWithTieredPricingPriceBillingCycleConfigurationDurationUnitDay   PlanVersionNewParamsAddPricesPriceNewPlanScalableMatrixWithTieredPricingPriceBillingCycleConfigurationDurationUnit = "day"
	PlanVersionNewParamsAddPricesPriceNewPlanScalableMatrixWithTieredPricingPriceBillingCycleConfigurationDurationUnitMonth PlanVersionNewParamsAddPricesPriceNewPlanScalableMatrixWithTieredPricingPriceBillingCycleConfigurationDurationUnit = "month"
)

func (r PlanVersionNewParamsAddPricesPriceNewPlanScalableMatrixWithTieredPricingPriceBillingCycleConfigurationDurationUnit) IsKnown() bool {
	switch r {
	case PlanVersionNewParamsAddPricesPriceNewPlanScalableMatrixWithTieredPricingPriceBillingCycleConfigurationDurationUnitDay, PlanVersionNewParamsAddPricesPriceNewPlanScalableMatrixWithTieredPricingPriceBillingCycleConfigurationDurationUnitMonth:
		return true
	}
	return false
}

// For dimensional price: specifies a price group and dimension values
type PlanVersionNewParamsAddPricesPriceNewPlanScalableMatrixWithTieredPricingPriceDimensionalPriceConfiguration struct {
	// The list of dimension values matching (in order) the dimensions of the price
	// group
	DimensionValues param.Field[[]string] `json:"dimension_values,required"`
	// The id of the dimensional price group to include this price in
	DimensionalPriceGroupID param.Field[string] `json:"dimensional_price_group_id"`
	// The external id of the dimensional price group to include this price in
	ExternalDimensionalPriceGroupID param.Field[string] `json:"external_dimensional_price_group_id"`
}

func (r PlanVersionNewParamsAddPricesPriceNewPlanScalableMatrixWithTieredPricingPriceDimensionalPriceConfiguration) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Within each billing cycle, specifies the cadence at which invoices are produced.
// If unspecified, a single invoice is produced per billing cycle.
type PlanVersionNewParamsAddPricesPriceNewPlanScalableMatrixWithTieredPricingPriceInvoicingCycleConfiguration struct {
	// The duration of the billing period.
	Duration param.Field[int64] `json:"duration,required"`
	// The unit of billing period duration.
	DurationUnit param.Field[PlanVersionNewParamsAddPricesPriceNewPlanScalableMatrixWithTieredPricingPriceInvoicingCycleConfigurationDurationUnit] `json:"duration_unit,required"`
}

func (r PlanVersionNewParamsAddPricesPriceNewPlanScalableMatrixWithTieredPricingPriceInvoicingCycleConfiguration) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// The unit of billing period duration.
type PlanVersionNewParamsAddPricesPriceNewPlanScalableMatrixWithTieredPricingPriceInvoicingCycleConfigurationDurationUnit string

const (
	PlanVersionNewParamsAddPricesPriceNewPlanScalableMatrixWithTieredPricingPriceInvoicingCycleConfigurationDurationUnitDay   PlanVersionNewParamsAddPricesPriceNewPlanScalableMatrixWithTieredPricingPriceInvoicingCycleConfigurationDurationUnit = "day"
	PlanVersionNewParamsAddPricesPriceNewPlanScalableMatrixWithTieredPricingPriceInvoicingCycleConfigurationDurationUnitMonth PlanVersionNewParamsAddPricesPriceNewPlanScalableMatrixWithTieredPricingPriceInvoicingCycleConfigurationDurationUnit = "month"
)

func (r PlanVersionNewParamsAddPricesPriceNewPlanScalableMatrixWithTieredPricingPriceInvoicingCycleConfigurationDurationUnit) IsKnown() bool {
	switch r {
	case PlanVersionNewParamsAddPricesPriceNewPlanScalableMatrixWithTieredPricingPriceInvoicingCycleConfigurationDurationUnitDay, PlanVersionNewParamsAddPricesPriceNewPlanScalableMatrixWithTieredPricingPriceInvoicingCycleConfigurationDurationUnitMonth:
		return true
	}
	return false
}

type PlanVersionNewParamsAddPricesPriceNewPlanCumulativeGroupedBulkPrice struct {
	// The cadence to bill for this price on.
	Cadence                     param.Field[PlanVersionNewParamsAddPricesPriceNewPlanCumulativeGroupedBulkPriceCadence] `json:"cadence,required"`
	CumulativeGroupedBulkConfig param.Field[map[string]interface{}]                                                     `json:"cumulative_grouped_bulk_config,required"`
	// The id of the item the price will be associated with.
	ItemID    param.Field[string]                                                                       `json:"item_id,required"`
	ModelType param.Field[PlanVersionNewParamsAddPricesPriceNewPlanCumulativeGroupedBulkPriceModelType] `json:"model_type,required"`
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
	BillingCycleConfiguration param.Field[PlanVersionNewParamsAddPricesPriceNewPlanCumulativeGroupedBulkPriceBillingCycleConfiguration] `json:"billing_cycle_configuration"`
	// The per unit conversion rate of the price currency to the invoicing currency.
	ConversionRate param.Field[float64] `json:"conversion_rate"`
	// An ISO 4217 currency string, or custom pricing unit identifier, in which this
	// price is billed.
	Currency param.Field[string] `json:"currency"`
	// For dimensional price: specifies a price group and dimension values
	DimensionalPriceConfiguration param.Field[PlanVersionNewParamsAddPricesPriceNewPlanCumulativeGroupedBulkPriceDimensionalPriceConfiguration] `json:"dimensional_price_configuration"`
	// An alias for the price.
	ExternalPriceID param.Field[string] `json:"external_price_id"`
	// If the Price represents a fixed cost, this represents the quantity of units
	// applied.
	FixedPriceQuantity param.Field[float64] `json:"fixed_price_quantity"`
	// The property used to group this price on an invoice
	InvoiceGroupingKey param.Field[string] `json:"invoice_grouping_key"`
	// Within each billing cycle, specifies the cadence at which invoices are produced.
	// If unspecified, a single invoice is produced per billing cycle.
	InvoicingCycleConfiguration param.Field[PlanVersionNewParamsAddPricesPriceNewPlanCumulativeGroupedBulkPriceInvoicingCycleConfiguration] `json:"invoicing_cycle_configuration"`
	// User-specified key/value pairs for the resource. Individual keys can be removed
	// by setting the value to `null`, and the entire metadata mapping can be cleared
	// by setting `metadata` to `null`.
	Metadata param.Field[map[string]string] `json:"metadata"`
}

func (r PlanVersionNewParamsAddPricesPriceNewPlanCumulativeGroupedBulkPrice) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r PlanVersionNewParamsAddPricesPriceNewPlanCumulativeGroupedBulkPrice) implementsPlanVersionNewParamsAddPricesPriceUnion() {
}

// The cadence to bill for this price on.
type PlanVersionNewParamsAddPricesPriceNewPlanCumulativeGroupedBulkPriceCadence string

const (
	PlanVersionNewParamsAddPricesPriceNewPlanCumulativeGroupedBulkPriceCadenceAnnual     PlanVersionNewParamsAddPricesPriceNewPlanCumulativeGroupedBulkPriceCadence = "annual"
	PlanVersionNewParamsAddPricesPriceNewPlanCumulativeGroupedBulkPriceCadenceSemiAnnual PlanVersionNewParamsAddPricesPriceNewPlanCumulativeGroupedBulkPriceCadence = "semi_annual"
	PlanVersionNewParamsAddPricesPriceNewPlanCumulativeGroupedBulkPriceCadenceMonthly    PlanVersionNewParamsAddPricesPriceNewPlanCumulativeGroupedBulkPriceCadence = "monthly"
	PlanVersionNewParamsAddPricesPriceNewPlanCumulativeGroupedBulkPriceCadenceQuarterly  PlanVersionNewParamsAddPricesPriceNewPlanCumulativeGroupedBulkPriceCadence = "quarterly"
	PlanVersionNewParamsAddPricesPriceNewPlanCumulativeGroupedBulkPriceCadenceOneTime    PlanVersionNewParamsAddPricesPriceNewPlanCumulativeGroupedBulkPriceCadence = "one_time"
	PlanVersionNewParamsAddPricesPriceNewPlanCumulativeGroupedBulkPriceCadenceCustom     PlanVersionNewParamsAddPricesPriceNewPlanCumulativeGroupedBulkPriceCadence = "custom"
)

func (r PlanVersionNewParamsAddPricesPriceNewPlanCumulativeGroupedBulkPriceCadence) IsKnown() bool {
	switch r {
	case PlanVersionNewParamsAddPricesPriceNewPlanCumulativeGroupedBulkPriceCadenceAnnual, PlanVersionNewParamsAddPricesPriceNewPlanCumulativeGroupedBulkPriceCadenceSemiAnnual, PlanVersionNewParamsAddPricesPriceNewPlanCumulativeGroupedBulkPriceCadenceMonthly, PlanVersionNewParamsAddPricesPriceNewPlanCumulativeGroupedBulkPriceCadenceQuarterly, PlanVersionNewParamsAddPricesPriceNewPlanCumulativeGroupedBulkPriceCadenceOneTime, PlanVersionNewParamsAddPricesPriceNewPlanCumulativeGroupedBulkPriceCadenceCustom:
		return true
	}
	return false
}

type PlanVersionNewParamsAddPricesPriceNewPlanCumulativeGroupedBulkPriceModelType string

const (
	PlanVersionNewParamsAddPricesPriceNewPlanCumulativeGroupedBulkPriceModelTypeCumulativeGroupedBulk PlanVersionNewParamsAddPricesPriceNewPlanCumulativeGroupedBulkPriceModelType = "cumulative_grouped_bulk"
)

func (r PlanVersionNewParamsAddPricesPriceNewPlanCumulativeGroupedBulkPriceModelType) IsKnown() bool {
	switch r {
	case PlanVersionNewParamsAddPricesPriceNewPlanCumulativeGroupedBulkPriceModelTypeCumulativeGroupedBulk:
		return true
	}
	return false
}

// For custom cadence: specifies the duration of the billing period in days or
// months.
type PlanVersionNewParamsAddPricesPriceNewPlanCumulativeGroupedBulkPriceBillingCycleConfiguration struct {
	// The duration of the billing period.
	Duration param.Field[int64] `json:"duration,required"`
	// The unit of billing period duration.
	DurationUnit param.Field[PlanVersionNewParamsAddPricesPriceNewPlanCumulativeGroupedBulkPriceBillingCycleConfigurationDurationUnit] `json:"duration_unit,required"`
}

func (r PlanVersionNewParamsAddPricesPriceNewPlanCumulativeGroupedBulkPriceBillingCycleConfiguration) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// The unit of billing period duration.
type PlanVersionNewParamsAddPricesPriceNewPlanCumulativeGroupedBulkPriceBillingCycleConfigurationDurationUnit string

const (
	PlanVersionNewParamsAddPricesPriceNewPlanCumulativeGroupedBulkPriceBillingCycleConfigurationDurationUnitDay   PlanVersionNewParamsAddPricesPriceNewPlanCumulativeGroupedBulkPriceBillingCycleConfigurationDurationUnit = "day"
	PlanVersionNewParamsAddPricesPriceNewPlanCumulativeGroupedBulkPriceBillingCycleConfigurationDurationUnitMonth PlanVersionNewParamsAddPricesPriceNewPlanCumulativeGroupedBulkPriceBillingCycleConfigurationDurationUnit = "month"
)

func (r PlanVersionNewParamsAddPricesPriceNewPlanCumulativeGroupedBulkPriceBillingCycleConfigurationDurationUnit) IsKnown() bool {
	switch r {
	case PlanVersionNewParamsAddPricesPriceNewPlanCumulativeGroupedBulkPriceBillingCycleConfigurationDurationUnitDay, PlanVersionNewParamsAddPricesPriceNewPlanCumulativeGroupedBulkPriceBillingCycleConfigurationDurationUnitMonth:
		return true
	}
	return false
}

// For dimensional price: specifies a price group and dimension values
type PlanVersionNewParamsAddPricesPriceNewPlanCumulativeGroupedBulkPriceDimensionalPriceConfiguration struct {
	// The list of dimension values matching (in order) the dimensions of the price
	// group
	DimensionValues param.Field[[]string] `json:"dimension_values,required"`
	// The id of the dimensional price group to include this price in
	DimensionalPriceGroupID param.Field[string] `json:"dimensional_price_group_id"`
	// The external id of the dimensional price group to include this price in
	ExternalDimensionalPriceGroupID param.Field[string] `json:"external_dimensional_price_group_id"`
}

func (r PlanVersionNewParamsAddPricesPriceNewPlanCumulativeGroupedBulkPriceDimensionalPriceConfiguration) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Within each billing cycle, specifies the cadence at which invoices are produced.
// If unspecified, a single invoice is produced per billing cycle.
type PlanVersionNewParamsAddPricesPriceNewPlanCumulativeGroupedBulkPriceInvoicingCycleConfiguration struct {
	// The duration of the billing period.
	Duration param.Field[int64] `json:"duration,required"`
	// The unit of billing period duration.
	DurationUnit param.Field[PlanVersionNewParamsAddPricesPriceNewPlanCumulativeGroupedBulkPriceInvoicingCycleConfigurationDurationUnit] `json:"duration_unit,required"`
}

func (r PlanVersionNewParamsAddPricesPriceNewPlanCumulativeGroupedBulkPriceInvoicingCycleConfiguration) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// The unit of billing period duration.
type PlanVersionNewParamsAddPricesPriceNewPlanCumulativeGroupedBulkPriceInvoicingCycleConfigurationDurationUnit string

const (
	PlanVersionNewParamsAddPricesPriceNewPlanCumulativeGroupedBulkPriceInvoicingCycleConfigurationDurationUnitDay   PlanVersionNewParamsAddPricesPriceNewPlanCumulativeGroupedBulkPriceInvoicingCycleConfigurationDurationUnit = "day"
	PlanVersionNewParamsAddPricesPriceNewPlanCumulativeGroupedBulkPriceInvoicingCycleConfigurationDurationUnitMonth PlanVersionNewParamsAddPricesPriceNewPlanCumulativeGroupedBulkPriceInvoicingCycleConfigurationDurationUnit = "month"
)

func (r PlanVersionNewParamsAddPricesPriceNewPlanCumulativeGroupedBulkPriceInvoicingCycleConfigurationDurationUnit) IsKnown() bool {
	switch r {
	case PlanVersionNewParamsAddPricesPriceNewPlanCumulativeGroupedBulkPriceInvoicingCycleConfigurationDurationUnitDay, PlanVersionNewParamsAddPricesPriceNewPlanCumulativeGroupedBulkPriceInvoicingCycleConfigurationDurationUnitMonth:
		return true
	}
	return false
}

type PlanVersionNewParamsAddPricesPriceNewPlanTieredPackageWithMinimumPrice struct {
	// The cadence to bill for this price on.
	Cadence param.Field[PlanVersionNewParamsAddPricesPriceNewPlanTieredPackageWithMinimumPriceCadence] `json:"cadence,required"`
	// The id of the item the price will be associated with.
	ItemID    param.Field[string]                                                                          `json:"item_id,required"`
	ModelType param.Field[PlanVersionNewParamsAddPricesPriceNewPlanTieredPackageWithMinimumPriceModelType] `json:"model_type,required"`
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
	BillingCycleConfiguration param.Field[PlanVersionNewParamsAddPricesPriceNewPlanTieredPackageWithMinimumPriceBillingCycleConfiguration] `json:"billing_cycle_configuration"`
	// The per unit conversion rate of the price currency to the invoicing currency.
	ConversionRate param.Field[float64] `json:"conversion_rate"`
	// An ISO 4217 currency string, or custom pricing unit identifier, in which this
	// price is billed.
	Currency param.Field[string] `json:"currency"`
	// For dimensional price: specifies a price group and dimension values
	DimensionalPriceConfiguration param.Field[PlanVersionNewParamsAddPricesPriceNewPlanTieredPackageWithMinimumPriceDimensionalPriceConfiguration] `json:"dimensional_price_configuration"`
	// An alias for the price.
	ExternalPriceID param.Field[string] `json:"external_price_id"`
	// If the Price represents a fixed cost, this represents the quantity of units
	// applied.
	FixedPriceQuantity param.Field[float64] `json:"fixed_price_quantity"`
	// The property used to group this price on an invoice
	InvoiceGroupingKey param.Field[string] `json:"invoice_grouping_key"`
	// Within each billing cycle, specifies the cadence at which invoices are produced.
	// If unspecified, a single invoice is produced per billing cycle.
	InvoicingCycleConfiguration param.Field[PlanVersionNewParamsAddPricesPriceNewPlanTieredPackageWithMinimumPriceInvoicingCycleConfiguration] `json:"invoicing_cycle_configuration"`
	// User-specified key/value pairs for the resource. Individual keys can be removed
	// by setting the value to `null`, and the entire metadata mapping can be cleared
	// by setting `metadata` to `null`.
	Metadata param.Field[map[string]string] `json:"metadata"`
}

func (r PlanVersionNewParamsAddPricesPriceNewPlanTieredPackageWithMinimumPrice) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r PlanVersionNewParamsAddPricesPriceNewPlanTieredPackageWithMinimumPrice) implementsPlanVersionNewParamsAddPricesPriceUnion() {
}

// The cadence to bill for this price on.
type PlanVersionNewParamsAddPricesPriceNewPlanTieredPackageWithMinimumPriceCadence string

const (
	PlanVersionNewParamsAddPricesPriceNewPlanTieredPackageWithMinimumPriceCadenceAnnual     PlanVersionNewParamsAddPricesPriceNewPlanTieredPackageWithMinimumPriceCadence = "annual"
	PlanVersionNewParamsAddPricesPriceNewPlanTieredPackageWithMinimumPriceCadenceSemiAnnual PlanVersionNewParamsAddPricesPriceNewPlanTieredPackageWithMinimumPriceCadence = "semi_annual"
	PlanVersionNewParamsAddPricesPriceNewPlanTieredPackageWithMinimumPriceCadenceMonthly    PlanVersionNewParamsAddPricesPriceNewPlanTieredPackageWithMinimumPriceCadence = "monthly"
	PlanVersionNewParamsAddPricesPriceNewPlanTieredPackageWithMinimumPriceCadenceQuarterly  PlanVersionNewParamsAddPricesPriceNewPlanTieredPackageWithMinimumPriceCadence = "quarterly"
	PlanVersionNewParamsAddPricesPriceNewPlanTieredPackageWithMinimumPriceCadenceOneTime    PlanVersionNewParamsAddPricesPriceNewPlanTieredPackageWithMinimumPriceCadence = "one_time"
	PlanVersionNewParamsAddPricesPriceNewPlanTieredPackageWithMinimumPriceCadenceCustom     PlanVersionNewParamsAddPricesPriceNewPlanTieredPackageWithMinimumPriceCadence = "custom"
)

func (r PlanVersionNewParamsAddPricesPriceNewPlanTieredPackageWithMinimumPriceCadence) IsKnown() bool {
	switch r {
	case PlanVersionNewParamsAddPricesPriceNewPlanTieredPackageWithMinimumPriceCadenceAnnual, PlanVersionNewParamsAddPricesPriceNewPlanTieredPackageWithMinimumPriceCadenceSemiAnnual, PlanVersionNewParamsAddPricesPriceNewPlanTieredPackageWithMinimumPriceCadenceMonthly, PlanVersionNewParamsAddPricesPriceNewPlanTieredPackageWithMinimumPriceCadenceQuarterly, PlanVersionNewParamsAddPricesPriceNewPlanTieredPackageWithMinimumPriceCadenceOneTime, PlanVersionNewParamsAddPricesPriceNewPlanTieredPackageWithMinimumPriceCadenceCustom:
		return true
	}
	return false
}

type PlanVersionNewParamsAddPricesPriceNewPlanTieredPackageWithMinimumPriceModelType string

const (
	PlanVersionNewParamsAddPricesPriceNewPlanTieredPackageWithMinimumPriceModelTypeTieredPackageWithMinimum PlanVersionNewParamsAddPricesPriceNewPlanTieredPackageWithMinimumPriceModelType = "tiered_package_with_minimum"
)

func (r PlanVersionNewParamsAddPricesPriceNewPlanTieredPackageWithMinimumPriceModelType) IsKnown() bool {
	switch r {
	case PlanVersionNewParamsAddPricesPriceNewPlanTieredPackageWithMinimumPriceModelTypeTieredPackageWithMinimum:
		return true
	}
	return false
}

// For custom cadence: specifies the duration of the billing period in days or
// months.
type PlanVersionNewParamsAddPricesPriceNewPlanTieredPackageWithMinimumPriceBillingCycleConfiguration struct {
	// The duration of the billing period.
	Duration param.Field[int64] `json:"duration,required"`
	// The unit of billing period duration.
	DurationUnit param.Field[PlanVersionNewParamsAddPricesPriceNewPlanTieredPackageWithMinimumPriceBillingCycleConfigurationDurationUnit] `json:"duration_unit,required"`
}

func (r PlanVersionNewParamsAddPricesPriceNewPlanTieredPackageWithMinimumPriceBillingCycleConfiguration) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// The unit of billing period duration.
type PlanVersionNewParamsAddPricesPriceNewPlanTieredPackageWithMinimumPriceBillingCycleConfigurationDurationUnit string

const (
	PlanVersionNewParamsAddPricesPriceNewPlanTieredPackageWithMinimumPriceBillingCycleConfigurationDurationUnitDay   PlanVersionNewParamsAddPricesPriceNewPlanTieredPackageWithMinimumPriceBillingCycleConfigurationDurationUnit = "day"
	PlanVersionNewParamsAddPricesPriceNewPlanTieredPackageWithMinimumPriceBillingCycleConfigurationDurationUnitMonth PlanVersionNewParamsAddPricesPriceNewPlanTieredPackageWithMinimumPriceBillingCycleConfigurationDurationUnit = "month"
)

func (r PlanVersionNewParamsAddPricesPriceNewPlanTieredPackageWithMinimumPriceBillingCycleConfigurationDurationUnit) IsKnown() bool {
	switch r {
	case PlanVersionNewParamsAddPricesPriceNewPlanTieredPackageWithMinimumPriceBillingCycleConfigurationDurationUnitDay, PlanVersionNewParamsAddPricesPriceNewPlanTieredPackageWithMinimumPriceBillingCycleConfigurationDurationUnitMonth:
		return true
	}
	return false
}

// For dimensional price: specifies a price group and dimension values
type PlanVersionNewParamsAddPricesPriceNewPlanTieredPackageWithMinimumPriceDimensionalPriceConfiguration struct {
	// The list of dimension values matching (in order) the dimensions of the price
	// group
	DimensionValues param.Field[[]string] `json:"dimension_values,required"`
	// The id of the dimensional price group to include this price in
	DimensionalPriceGroupID param.Field[string] `json:"dimensional_price_group_id"`
	// The external id of the dimensional price group to include this price in
	ExternalDimensionalPriceGroupID param.Field[string] `json:"external_dimensional_price_group_id"`
}

func (r PlanVersionNewParamsAddPricesPriceNewPlanTieredPackageWithMinimumPriceDimensionalPriceConfiguration) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Within each billing cycle, specifies the cadence at which invoices are produced.
// If unspecified, a single invoice is produced per billing cycle.
type PlanVersionNewParamsAddPricesPriceNewPlanTieredPackageWithMinimumPriceInvoicingCycleConfiguration struct {
	// The duration of the billing period.
	Duration param.Field[int64] `json:"duration,required"`
	// The unit of billing period duration.
	DurationUnit param.Field[PlanVersionNewParamsAddPricesPriceNewPlanTieredPackageWithMinimumPriceInvoicingCycleConfigurationDurationUnit] `json:"duration_unit,required"`
}

func (r PlanVersionNewParamsAddPricesPriceNewPlanTieredPackageWithMinimumPriceInvoicingCycleConfiguration) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// The unit of billing period duration.
type PlanVersionNewParamsAddPricesPriceNewPlanTieredPackageWithMinimumPriceInvoicingCycleConfigurationDurationUnit string

const (
	PlanVersionNewParamsAddPricesPriceNewPlanTieredPackageWithMinimumPriceInvoicingCycleConfigurationDurationUnitDay   PlanVersionNewParamsAddPricesPriceNewPlanTieredPackageWithMinimumPriceInvoicingCycleConfigurationDurationUnit = "day"
	PlanVersionNewParamsAddPricesPriceNewPlanTieredPackageWithMinimumPriceInvoicingCycleConfigurationDurationUnitMonth PlanVersionNewParamsAddPricesPriceNewPlanTieredPackageWithMinimumPriceInvoicingCycleConfigurationDurationUnit = "month"
)

func (r PlanVersionNewParamsAddPricesPriceNewPlanTieredPackageWithMinimumPriceInvoicingCycleConfigurationDurationUnit) IsKnown() bool {
	switch r {
	case PlanVersionNewParamsAddPricesPriceNewPlanTieredPackageWithMinimumPriceInvoicingCycleConfigurationDurationUnitDay, PlanVersionNewParamsAddPricesPriceNewPlanTieredPackageWithMinimumPriceInvoicingCycleConfigurationDurationUnitMonth:
		return true
	}
	return false
}

type PlanVersionNewParamsAddPricesPriceNewPlanMatrixWithAllocationPrice struct {
	// The cadence to bill for this price on.
	Cadence param.Field[PlanVersionNewParamsAddPricesPriceNewPlanMatrixWithAllocationPriceCadence] `json:"cadence,required"`
	// The id of the item the price will be associated with.
	ItemID                     param.Field[string]                                                                                       `json:"item_id,required"`
	MatrixWithAllocationConfig param.Field[PlanVersionNewParamsAddPricesPriceNewPlanMatrixWithAllocationPriceMatrixWithAllocationConfig] `json:"matrix_with_allocation_config,required"`
	ModelType                  param.Field[PlanVersionNewParamsAddPricesPriceNewPlanMatrixWithAllocationPriceModelType]                  `json:"model_type,required"`
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
	BillingCycleConfiguration param.Field[PlanVersionNewParamsAddPricesPriceNewPlanMatrixWithAllocationPriceBillingCycleConfiguration] `json:"billing_cycle_configuration"`
	// The per unit conversion rate of the price currency to the invoicing currency.
	ConversionRate param.Field[float64] `json:"conversion_rate"`
	// An ISO 4217 currency string, or custom pricing unit identifier, in which this
	// price is billed.
	Currency param.Field[string] `json:"currency"`
	// For dimensional price: specifies a price group and dimension values
	DimensionalPriceConfiguration param.Field[PlanVersionNewParamsAddPricesPriceNewPlanMatrixWithAllocationPriceDimensionalPriceConfiguration] `json:"dimensional_price_configuration"`
	// An alias for the price.
	ExternalPriceID param.Field[string] `json:"external_price_id"`
	// If the Price represents a fixed cost, this represents the quantity of units
	// applied.
	FixedPriceQuantity param.Field[float64] `json:"fixed_price_quantity"`
	// The property used to group this price on an invoice
	InvoiceGroupingKey param.Field[string] `json:"invoice_grouping_key"`
	// Within each billing cycle, specifies the cadence at which invoices are produced.
	// If unspecified, a single invoice is produced per billing cycle.
	InvoicingCycleConfiguration param.Field[PlanVersionNewParamsAddPricesPriceNewPlanMatrixWithAllocationPriceInvoicingCycleConfiguration] `json:"invoicing_cycle_configuration"`
	// User-specified key/value pairs for the resource. Individual keys can be removed
	// by setting the value to `null`, and the entire metadata mapping can be cleared
	// by setting `metadata` to `null`.
	Metadata param.Field[map[string]string] `json:"metadata"`
}

func (r PlanVersionNewParamsAddPricesPriceNewPlanMatrixWithAllocationPrice) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r PlanVersionNewParamsAddPricesPriceNewPlanMatrixWithAllocationPrice) implementsPlanVersionNewParamsAddPricesPriceUnion() {
}

// The cadence to bill for this price on.
type PlanVersionNewParamsAddPricesPriceNewPlanMatrixWithAllocationPriceCadence string

const (
	PlanVersionNewParamsAddPricesPriceNewPlanMatrixWithAllocationPriceCadenceAnnual     PlanVersionNewParamsAddPricesPriceNewPlanMatrixWithAllocationPriceCadence = "annual"
	PlanVersionNewParamsAddPricesPriceNewPlanMatrixWithAllocationPriceCadenceSemiAnnual PlanVersionNewParamsAddPricesPriceNewPlanMatrixWithAllocationPriceCadence = "semi_annual"
	PlanVersionNewParamsAddPricesPriceNewPlanMatrixWithAllocationPriceCadenceMonthly    PlanVersionNewParamsAddPricesPriceNewPlanMatrixWithAllocationPriceCadence = "monthly"
	PlanVersionNewParamsAddPricesPriceNewPlanMatrixWithAllocationPriceCadenceQuarterly  PlanVersionNewParamsAddPricesPriceNewPlanMatrixWithAllocationPriceCadence = "quarterly"
	PlanVersionNewParamsAddPricesPriceNewPlanMatrixWithAllocationPriceCadenceOneTime    PlanVersionNewParamsAddPricesPriceNewPlanMatrixWithAllocationPriceCadence = "one_time"
	PlanVersionNewParamsAddPricesPriceNewPlanMatrixWithAllocationPriceCadenceCustom     PlanVersionNewParamsAddPricesPriceNewPlanMatrixWithAllocationPriceCadence = "custom"
)

func (r PlanVersionNewParamsAddPricesPriceNewPlanMatrixWithAllocationPriceCadence) IsKnown() bool {
	switch r {
	case PlanVersionNewParamsAddPricesPriceNewPlanMatrixWithAllocationPriceCadenceAnnual, PlanVersionNewParamsAddPricesPriceNewPlanMatrixWithAllocationPriceCadenceSemiAnnual, PlanVersionNewParamsAddPricesPriceNewPlanMatrixWithAllocationPriceCadenceMonthly, PlanVersionNewParamsAddPricesPriceNewPlanMatrixWithAllocationPriceCadenceQuarterly, PlanVersionNewParamsAddPricesPriceNewPlanMatrixWithAllocationPriceCadenceOneTime, PlanVersionNewParamsAddPricesPriceNewPlanMatrixWithAllocationPriceCadenceCustom:
		return true
	}
	return false
}

type PlanVersionNewParamsAddPricesPriceNewPlanMatrixWithAllocationPriceMatrixWithAllocationConfig struct {
	// Allocation to be used to calculate the price
	Allocation param.Field[float64] `json:"allocation,required"`
	// Default per unit rate for any usage not bucketed into a specified matrix_value
	DefaultUnitAmount param.Field[string] `json:"default_unit_amount,required"`
	// One or two event property values to evaluate matrix groups by
	Dimensions param.Field[[]string] `json:"dimensions,required"`
	// Matrix values for specified matrix grouping keys
	MatrixValues param.Field[[]PlanVersionNewParamsAddPricesPriceNewPlanMatrixWithAllocationPriceMatrixWithAllocationConfigMatrixValue] `json:"matrix_values,required"`
}

func (r PlanVersionNewParamsAddPricesPriceNewPlanMatrixWithAllocationPriceMatrixWithAllocationConfig) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type PlanVersionNewParamsAddPricesPriceNewPlanMatrixWithAllocationPriceMatrixWithAllocationConfigMatrixValue struct {
	// One or two matrix keys to filter usage to this Matrix value by. For example,
	// ["region", "tier"] could be used to filter cloud usage by a cloud region and an
	// instance tier.
	DimensionValues param.Field[[]string] `json:"dimension_values,required"`
	// Unit price for the specified dimension_values
	UnitAmount param.Field[string] `json:"unit_amount,required"`
}

func (r PlanVersionNewParamsAddPricesPriceNewPlanMatrixWithAllocationPriceMatrixWithAllocationConfigMatrixValue) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type PlanVersionNewParamsAddPricesPriceNewPlanMatrixWithAllocationPriceModelType string

const (
	PlanVersionNewParamsAddPricesPriceNewPlanMatrixWithAllocationPriceModelTypeMatrixWithAllocation PlanVersionNewParamsAddPricesPriceNewPlanMatrixWithAllocationPriceModelType = "matrix_with_allocation"
)

func (r PlanVersionNewParamsAddPricesPriceNewPlanMatrixWithAllocationPriceModelType) IsKnown() bool {
	switch r {
	case PlanVersionNewParamsAddPricesPriceNewPlanMatrixWithAllocationPriceModelTypeMatrixWithAllocation:
		return true
	}
	return false
}

// For custom cadence: specifies the duration of the billing period in days or
// months.
type PlanVersionNewParamsAddPricesPriceNewPlanMatrixWithAllocationPriceBillingCycleConfiguration struct {
	// The duration of the billing period.
	Duration param.Field[int64] `json:"duration,required"`
	// The unit of billing period duration.
	DurationUnit param.Field[PlanVersionNewParamsAddPricesPriceNewPlanMatrixWithAllocationPriceBillingCycleConfigurationDurationUnit] `json:"duration_unit,required"`
}

func (r PlanVersionNewParamsAddPricesPriceNewPlanMatrixWithAllocationPriceBillingCycleConfiguration) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// The unit of billing period duration.
type PlanVersionNewParamsAddPricesPriceNewPlanMatrixWithAllocationPriceBillingCycleConfigurationDurationUnit string

const (
	PlanVersionNewParamsAddPricesPriceNewPlanMatrixWithAllocationPriceBillingCycleConfigurationDurationUnitDay   PlanVersionNewParamsAddPricesPriceNewPlanMatrixWithAllocationPriceBillingCycleConfigurationDurationUnit = "day"
	PlanVersionNewParamsAddPricesPriceNewPlanMatrixWithAllocationPriceBillingCycleConfigurationDurationUnitMonth PlanVersionNewParamsAddPricesPriceNewPlanMatrixWithAllocationPriceBillingCycleConfigurationDurationUnit = "month"
)

func (r PlanVersionNewParamsAddPricesPriceNewPlanMatrixWithAllocationPriceBillingCycleConfigurationDurationUnit) IsKnown() bool {
	switch r {
	case PlanVersionNewParamsAddPricesPriceNewPlanMatrixWithAllocationPriceBillingCycleConfigurationDurationUnitDay, PlanVersionNewParamsAddPricesPriceNewPlanMatrixWithAllocationPriceBillingCycleConfigurationDurationUnitMonth:
		return true
	}
	return false
}

// For dimensional price: specifies a price group and dimension values
type PlanVersionNewParamsAddPricesPriceNewPlanMatrixWithAllocationPriceDimensionalPriceConfiguration struct {
	// The list of dimension values matching (in order) the dimensions of the price
	// group
	DimensionValues param.Field[[]string] `json:"dimension_values,required"`
	// The id of the dimensional price group to include this price in
	DimensionalPriceGroupID param.Field[string] `json:"dimensional_price_group_id"`
	// The external id of the dimensional price group to include this price in
	ExternalDimensionalPriceGroupID param.Field[string] `json:"external_dimensional_price_group_id"`
}

func (r PlanVersionNewParamsAddPricesPriceNewPlanMatrixWithAllocationPriceDimensionalPriceConfiguration) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Within each billing cycle, specifies the cadence at which invoices are produced.
// If unspecified, a single invoice is produced per billing cycle.
type PlanVersionNewParamsAddPricesPriceNewPlanMatrixWithAllocationPriceInvoicingCycleConfiguration struct {
	// The duration of the billing period.
	Duration param.Field[int64] `json:"duration,required"`
	// The unit of billing period duration.
	DurationUnit param.Field[PlanVersionNewParamsAddPricesPriceNewPlanMatrixWithAllocationPriceInvoicingCycleConfigurationDurationUnit] `json:"duration_unit,required"`
}

func (r PlanVersionNewParamsAddPricesPriceNewPlanMatrixWithAllocationPriceInvoicingCycleConfiguration) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// The unit of billing period duration.
type PlanVersionNewParamsAddPricesPriceNewPlanMatrixWithAllocationPriceInvoicingCycleConfigurationDurationUnit string

const (
	PlanVersionNewParamsAddPricesPriceNewPlanMatrixWithAllocationPriceInvoicingCycleConfigurationDurationUnitDay   PlanVersionNewParamsAddPricesPriceNewPlanMatrixWithAllocationPriceInvoicingCycleConfigurationDurationUnit = "day"
	PlanVersionNewParamsAddPricesPriceNewPlanMatrixWithAllocationPriceInvoicingCycleConfigurationDurationUnitMonth PlanVersionNewParamsAddPricesPriceNewPlanMatrixWithAllocationPriceInvoicingCycleConfigurationDurationUnit = "month"
)

func (r PlanVersionNewParamsAddPricesPriceNewPlanMatrixWithAllocationPriceInvoicingCycleConfigurationDurationUnit) IsKnown() bool {
	switch r {
	case PlanVersionNewParamsAddPricesPriceNewPlanMatrixWithAllocationPriceInvoicingCycleConfigurationDurationUnitDay, PlanVersionNewParamsAddPricesPriceNewPlanMatrixWithAllocationPriceInvoicingCycleConfigurationDurationUnitMonth:
		return true
	}
	return false
}

type PlanVersionNewParamsAddPricesPriceNewPlanGroupedTieredPrice struct {
	// The cadence to bill for this price on.
	Cadence             param.Field[PlanVersionNewParamsAddPricesPriceNewPlanGroupedTieredPriceCadence] `json:"cadence,required"`
	GroupedTieredConfig param.Field[map[string]interface{}]                                             `json:"grouped_tiered_config,required"`
	// The id of the item the price will be associated with.
	ItemID    param.Field[string]                                                               `json:"item_id,required"`
	ModelType param.Field[PlanVersionNewParamsAddPricesPriceNewPlanGroupedTieredPriceModelType] `json:"model_type,required"`
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
	BillingCycleConfiguration param.Field[PlanVersionNewParamsAddPricesPriceNewPlanGroupedTieredPriceBillingCycleConfiguration] `json:"billing_cycle_configuration"`
	// The per unit conversion rate of the price currency to the invoicing currency.
	ConversionRate param.Field[float64] `json:"conversion_rate"`
	// An ISO 4217 currency string, or custom pricing unit identifier, in which this
	// price is billed.
	Currency param.Field[string] `json:"currency"`
	// For dimensional price: specifies a price group and dimension values
	DimensionalPriceConfiguration param.Field[PlanVersionNewParamsAddPricesPriceNewPlanGroupedTieredPriceDimensionalPriceConfiguration] `json:"dimensional_price_configuration"`
	// An alias for the price.
	ExternalPriceID param.Field[string] `json:"external_price_id"`
	// If the Price represents a fixed cost, this represents the quantity of units
	// applied.
	FixedPriceQuantity param.Field[float64] `json:"fixed_price_quantity"`
	// The property used to group this price on an invoice
	InvoiceGroupingKey param.Field[string] `json:"invoice_grouping_key"`
	// Within each billing cycle, specifies the cadence at which invoices are produced.
	// If unspecified, a single invoice is produced per billing cycle.
	InvoicingCycleConfiguration param.Field[PlanVersionNewParamsAddPricesPriceNewPlanGroupedTieredPriceInvoicingCycleConfiguration] `json:"invoicing_cycle_configuration"`
	// User-specified key/value pairs for the resource. Individual keys can be removed
	// by setting the value to `null`, and the entire metadata mapping can be cleared
	// by setting `metadata` to `null`.
	Metadata param.Field[map[string]string] `json:"metadata"`
}

func (r PlanVersionNewParamsAddPricesPriceNewPlanGroupedTieredPrice) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r PlanVersionNewParamsAddPricesPriceNewPlanGroupedTieredPrice) implementsPlanVersionNewParamsAddPricesPriceUnion() {
}

// The cadence to bill for this price on.
type PlanVersionNewParamsAddPricesPriceNewPlanGroupedTieredPriceCadence string

const (
	PlanVersionNewParamsAddPricesPriceNewPlanGroupedTieredPriceCadenceAnnual     PlanVersionNewParamsAddPricesPriceNewPlanGroupedTieredPriceCadence = "annual"
	PlanVersionNewParamsAddPricesPriceNewPlanGroupedTieredPriceCadenceSemiAnnual PlanVersionNewParamsAddPricesPriceNewPlanGroupedTieredPriceCadence = "semi_annual"
	PlanVersionNewParamsAddPricesPriceNewPlanGroupedTieredPriceCadenceMonthly    PlanVersionNewParamsAddPricesPriceNewPlanGroupedTieredPriceCadence = "monthly"
	PlanVersionNewParamsAddPricesPriceNewPlanGroupedTieredPriceCadenceQuarterly  PlanVersionNewParamsAddPricesPriceNewPlanGroupedTieredPriceCadence = "quarterly"
	PlanVersionNewParamsAddPricesPriceNewPlanGroupedTieredPriceCadenceOneTime    PlanVersionNewParamsAddPricesPriceNewPlanGroupedTieredPriceCadence = "one_time"
	PlanVersionNewParamsAddPricesPriceNewPlanGroupedTieredPriceCadenceCustom     PlanVersionNewParamsAddPricesPriceNewPlanGroupedTieredPriceCadence = "custom"
)

func (r PlanVersionNewParamsAddPricesPriceNewPlanGroupedTieredPriceCadence) IsKnown() bool {
	switch r {
	case PlanVersionNewParamsAddPricesPriceNewPlanGroupedTieredPriceCadenceAnnual, PlanVersionNewParamsAddPricesPriceNewPlanGroupedTieredPriceCadenceSemiAnnual, PlanVersionNewParamsAddPricesPriceNewPlanGroupedTieredPriceCadenceMonthly, PlanVersionNewParamsAddPricesPriceNewPlanGroupedTieredPriceCadenceQuarterly, PlanVersionNewParamsAddPricesPriceNewPlanGroupedTieredPriceCadenceOneTime, PlanVersionNewParamsAddPricesPriceNewPlanGroupedTieredPriceCadenceCustom:
		return true
	}
	return false
}

type PlanVersionNewParamsAddPricesPriceNewPlanGroupedTieredPriceModelType string

const (
	PlanVersionNewParamsAddPricesPriceNewPlanGroupedTieredPriceModelTypeGroupedTiered PlanVersionNewParamsAddPricesPriceNewPlanGroupedTieredPriceModelType = "grouped_tiered"
)

func (r PlanVersionNewParamsAddPricesPriceNewPlanGroupedTieredPriceModelType) IsKnown() bool {
	switch r {
	case PlanVersionNewParamsAddPricesPriceNewPlanGroupedTieredPriceModelTypeGroupedTiered:
		return true
	}
	return false
}

// For custom cadence: specifies the duration of the billing period in days or
// months.
type PlanVersionNewParamsAddPricesPriceNewPlanGroupedTieredPriceBillingCycleConfiguration struct {
	// The duration of the billing period.
	Duration param.Field[int64] `json:"duration,required"`
	// The unit of billing period duration.
	DurationUnit param.Field[PlanVersionNewParamsAddPricesPriceNewPlanGroupedTieredPriceBillingCycleConfigurationDurationUnit] `json:"duration_unit,required"`
}

func (r PlanVersionNewParamsAddPricesPriceNewPlanGroupedTieredPriceBillingCycleConfiguration) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// The unit of billing period duration.
type PlanVersionNewParamsAddPricesPriceNewPlanGroupedTieredPriceBillingCycleConfigurationDurationUnit string

const (
	PlanVersionNewParamsAddPricesPriceNewPlanGroupedTieredPriceBillingCycleConfigurationDurationUnitDay   PlanVersionNewParamsAddPricesPriceNewPlanGroupedTieredPriceBillingCycleConfigurationDurationUnit = "day"
	PlanVersionNewParamsAddPricesPriceNewPlanGroupedTieredPriceBillingCycleConfigurationDurationUnitMonth PlanVersionNewParamsAddPricesPriceNewPlanGroupedTieredPriceBillingCycleConfigurationDurationUnit = "month"
)

func (r PlanVersionNewParamsAddPricesPriceNewPlanGroupedTieredPriceBillingCycleConfigurationDurationUnit) IsKnown() bool {
	switch r {
	case PlanVersionNewParamsAddPricesPriceNewPlanGroupedTieredPriceBillingCycleConfigurationDurationUnitDay, PlanVersionNewParamsAddPricesPriceNewPlanGroupedTieredPriceBillingCycleConfigurationDurationUnitMonth:
		return true
	}
	return false
}

// For dimensional price: specifies a price group and dimension values
type PlanVersionNewParamsAddPricesPriceNewPlanGroupedTieredPriceDimensionalPriceConfiguration struct {
	// The list of dimension values matching (in order) the dimensions of the price
	// group
	DimensionValues param.Field[[]string] `json:"dimension_values,required"`
	// The id of the dimensional price group to include this price in
	DimensionalPriceGroupID param.Field[string] `json:"dimensional_price_group_id"`
	// The external id of the dimensional price group to include this price in
	ExternalDimensionalPriceGroupID param.Field[string] `json:"external_dimensional_price_group_id"`
}

func (r PlanVersionNewParamsAddPricesPriceNewPlanGroupedTieredPriceDimensionalPriceConfiguration) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Within each billing cycle, specifies the cadence at which invoices are produced.
// If unspecified, a single invoice is produced per billing cycle.
type PlanVersionNewParamsAddPricesPriceNewPlanGroupedTieredPriceInvoicingCycleConfiguration struct {
	// The duration of the billing period.
	Duration param.Field[int64] `json:"duration,required"`
	// The unit of billing period duration.
	DurationUnit param.Field[PlanVersionNewParamsAddPricesPriceNewPlanGroupedTieredPriceInvoicingCycleConfigurationDurationUnit] `json:"duration_unit,required"`
}

func (r PlanVersionNewParamsAddPricesPriceNewPlanGroupedTieredPriceInvoicingCycleConfiguration) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// The unit of billing period duration.
type PlanVersionNewParamsAddPricesPriceNewPlanGroupedTieredPriceInvoicingCycleConfigurationDurationUnit string

const (
	PlanVersionNewParamsAddPricesPriceNewPlanGroupedTieredPriceInvoicingCycleConfigurationDurationUnitDay   PlanVersionNewParamsAddPricesPriceNewPlanGroupedTieredPriceInvoicingCycleConfigurationDurationUnit = "day"
	PlanVersionNewParamsAddPricesPriceNewPlanGroupedTieredPriceInvoicingCycleConfigurationDurationUnitMonth PlanVersionNewParamsAddPricesPriceNewPlanGroupedTieredPriceInvoicingCycleConfigurationDurationUnit = "month"
)

func (r PlanVersionNewParamsAddPricesPriceNewPlanGroupedTieredPriceInvoicingCycleConfigurationDurationUnit) IsKnown() bool {
	switch r {
	case PlanVersionNewParamsAddPricesPriceNewPlanGroupedTieredPriceInvoicingCycleConfigurationDurationUnitDay, PlanVersionNewParamsAddPricesPriceNewPlanGroupedTieredPriceInvoicingCycleConfigurationDurationUnitMonth:
		return true
	}
	return false
}

// The cadence to bill for this price on.
type PlanVersionNewParamsAddPricesPriceCadence string

const (
	PlanVersionNewParamsAddPricesPriceCadenceAnnual     PlanVersionNewParamsAddPricesPriceCadence = "annual"
	PlanVersionNewParamsAddPricesPriceCadenceSemiAnnual PlanVersionNewParamsAddPricesPriceCadence = "semi_annual"
	PlanVersionNewParamsAddPricesPriceCadenceMonthly    PlanVersionNewParamsAddPricesPriceCadence = "monthly"
	PlanVersionNewParamsAddPricesPriceCadenceQuarterly  PlanVersionNewParamsAddPricesPriceCadence = "quarterly"
	PlanVersionNewParamsAddPricesPriceCadenceOneTime    PlanVersionNewParamsAddPricesPriceCadence = "one_time"
	PlanVersionNewParamsAddPricesPriceCadenceCustom     PlanVersionNewParamsAddPricesPriceCadence = "custom"
)

func (r PlanVersionNewParamsAddPricesPriceCadence) IsKnown() bool {
	switch r {
	case PlanVersionNewParamsAddPricesPriceCadenceAnnual, PlanVersionNewParamsAddPricesPriceCadenceSemiAnnual, PlanVersionNewParamsAddPricesPriceCadenceMonthly, PlanVersionNewParamsAddPricesPriceCadenceQuarterly, PlanVersionNewParamsAddPricesPriceCadenceOneTime, PlanVersionNewParamsAddPricesPriceCadenceCustom:
		return true
	}
	return false
}

type PlanVersionNewParamsAddPricesPriceModelType string

const (
	PlanVersionNewParamsAddPricesPriceModelTypeUnit                            PlanVersionNewParamsAddPricesPriceModelType = "unit"
	PlanVersionNewParamsAddPricesPriceModelTypePackage                         PlanVersionNewParamsAddPricesPriceModelType = "package"
	PlanVersionNewParamsAddPricesPriceModelTypeMatrix                          PlanVersionNewParamsAddPricesPriceModelType = "matrix"
	PlanVersionNewParamsAddPricesPriceModelTypeTiered                          PlanVersionNewParamsAddPricesPriceModelType = "tiered"
	PlanVersionNewParamsAddPricesPriceModelTypeTieredBps                       PlanVersionNewParamsAddPricesPriceModelType = "tiered_bps"
	PlanVersionNewParamsAddPricesPriceModelTypeBps                             PlanVersionNewParamsAddPricesPriceModelType = "bps"
	PlanVersionNewParamsAddPricesPriceModelTypeBulkBps                         PlanVersionNewParamsAddPricesPriceModelType = "bulk_bps"
	PlanVersionNewParamsAddPricesPriceModelTypeBulk                            PlanVersionNewParamsAddPricesPriceModelType = "bulk"
	PlanVersionNewParamsAddPricesPriceModelTypeThresholdTotalAmount            PlanVersionNewParamsAddPricesPriceModelType = "threshold_total_amount"
	PlanVersionNewParamsAddPricesPriceModelTypeTieredPackage                   PlanVersionNewParamsAddPricesPriceModelType = "tiered_package"
	PlanVersionNewParamsAddPricesPriceModelTypeTieredWithMinimum               PlanVersionNewParamsAddPricesPriceModelType = "tiered_with_minimum"
	PlanVersionNewParamsAddPricesPriceModelTypeUnitWithPercent                 PlanVersionNewParamsAddPricesPriceModelType = "unit_with_percent"
	PlanVersionNewParamsAddPricesPriceModelTypePackageWithAllocation           PlanVersionNewParamsAddPricesPriceModelType = "package_with_allocation"
	PlanVersionNewParamsAddPricesPriceModelTypeTieredWithProration             PlanVersionNewParamsAddPricesPriceModelType = "tiered_with_proration"
	PlanVersionNewParamsAddPricesPriceModelTypeUnitWithProration               PlanVersionNewParamsAddPricesPriceModelType = "unit_with_proration"
	PlanVersionNewParamsAddPricesPriceModelTypeGroupedAllocation               PlanVersionNewParamsAddPricesPriceModelType = "grouped_allocation"
	PlanVersionNewParamsAddPricesPriceModelTypeGroupedWithProratedMinimum      PlanVersionNewParamsAddPricesPriceModelType = "grouped_with_prorated_minimum"
	PlanVersionNewParamsAddPricesPriceModelTypeGroupedWithMeteredMinimum       PlanVersionNewParamsAddPricesPriceModelType = "grouped_with_metered_minimum"
	PlanVersionNewParamsAddPricesPriceModelTypeMatrixWithDisplayName           PlanVersionNewParamsAddPricesPriceModelType = "matrix_with_display_name"
	PlanVersionNewParamsAddPricesPriceModelTypeBulkWithProration               PlanVersionNewParamsAddPricesPriceModelType = "bulk_with_proration"
	PlanVersionNewParamsAddPricesPriceModelTypeGroupedTieredPackage            PlanVersionNewParamsAddPricesPriceModelType = "grouped_tiered_package"
	PlanVersionNewParamsAddPricesPriceModelTypeMaxGroupTieredPackage           PlanVersionNewParamsAddPricesPriceModelType = "max_group_tiered_package"
	PlanVersionNewParamsAddPricesPriceModelTypeScalableMatrixWithUnitPricing   PlanVersionNewParamsAddPricesPriceModelType = "scalable_matrix_with_unit_pricing"
	PlanVersionNewParamsAddPricesPriceModelTypeScalableMatrixWithTieredPricing PlanVersionNewParamsAddPricesPriceModelType = "scalable_matrix_with_tiered_pricing"
	PlanVersionNewParamsAddPricesPriceModelTypeCumulativeGroupedBulk           PlanVersionNewParamsAddPricesPriceModelType = "cumulative_grouped_bulk"
	PlanVersionNewParamsAddPricesPriceModelTypeTieredPackageWithMinimum        PlanVersionNewParamsAddPricesPriceModelType = "tiered_package_with_minimum"
	PlanVersionNewParamsAddPricesPriceModelTypeMatrixWithAllocation            PlanVersionNewParamsAddPricesPriceModelType = "matrix_with_allocation"
	PlanVersionNewParamsAddPricesPriceModelTypeGroupedTiered                   PlanVersionNewParamsAddPricesPriceModelType = "grouped_tiered"
)

func (r PlanVersionNewParamsAddPricesPriceModelType) IsKnown() bool {
	switch r {
	case PlanVersionNewParamsAddPricesPriceModelTypeUnit, PlanVersionNewParamsAddPricesPriceModelTypePackage, PlanVersionNewParamsAddPricesPriceModelTypeMatrix, PlanVersionNewParamsAddPricesPriceModelTypeTiered, PlanVersionNewParamsAddPricesPriceModelTypeTieredBps, PlanVersionNewParamsAddPricesPriceModelTypeBps, PlanVersionNewParamsAddPricesPriceModelTypeBulkBps, PlanVersionNewParamsAddPricesPriceModelTypeBulk, PlanVersionNewParamsAddPricesPriceModelTypeThresholdTotalAmount, PlanVersionNewParamsAddPricesPriceModelTypeTieredPackage, PlanVersionNewParamsAddPricesPriceModelTypeTieredWithMinimum, PlanVersionNewParamsAddPricesPriceModelTypeUnitWithPercent, PlanVersionNewParamsAddPricesPriceModelTypePackageWithAllocation, PlanVersionNewParamsAddPricesPriceModelTypeTieredWithProration, PlanVersionNewParamsAddPricesPriceModelTypeUnitWithProration, PlanVersionNewParamsAddPricesPriceModelTypeGroupedAllocation, PlanVersionNewParamsAddPricesPriceModelTypeGroupedWithProratedMinimum, PlanVersionNewParamsAddPricesPriceModelTypeGroupedWithMeteredMinimum, PlanVersionNewParamsAddPricesPriceModelTypeMatrixWithDisplayName, PlanVersionNewParamsAddPricesPriceModelTypeBulkWithProration, PlanVersionNewParamsAddPricesPriceModelTypeGroupedTieredPackage, PlanVersionNewParamsAddPricesPriceModelTypeMaxGroupTieredPackage, PlanVersionNewParamsAddPricesPriceModelTypeScalableMatrixWithUnitPricing, PlanVersionNewParamsAddPricesPriceModelTypeScalableMatrixWithTieredPricing, PlanVersionNewParamsAddPricesPriceModelTypeCumulativeGroupedBulk, PlanVersionNewParamsAddPricesPriceModelTypeTieredPackageWithMinimum, PlanVersionNewParamsAddPricesPriceModelTypeMatrixWithAllocation, PlanVersionNewParamsAddPricesPriceModelTypeGroupedTiered:
		return true
	}
	return false
}

type PlanVersionNewParamsRemoveAdjustment struct {
	// The id of the adjustment to remove from on the plan.
	AdjustmentID param.Field[string] `json:"adjustment_id,required"`
	// The phase to remove this adjustment from.
	PlanPhaseOrder param.Field[int64] `json:"plan_phase_order"`
}

func (r PlanVersionNewParamsRemoveAdjustment) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type PlanVersionNewParamsRemovePrice struct {
	// The id of the price to remove from the plan.
	PriceID param.Field[string] `json:"price_id,required"`
	// The phase to remove this price from.
	PlanPhaseOrder param.Field[int64] `json:"plan_phase_order"`
}

func (r PlanVersionNewParamsRemovePrice) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type PlanVersionNewParamsReplaceAdjustment struct {
	// The definition of a new adjustment to create and add to the plan.
	Adjustment param.Field[PlanVersionNewParamsReplaceAdjustmentsAdjustmentUnion] `json:"adjustment,required"`
	// The id of the adjustment on the plan to replace in the plan.
	ReplacesAdjustmentID param.Field[string] `json:"replaces_adjustment_id,required"`
	// The phase to replace this adjustment from.
	PlanPhaseOrder param.Field[int64] `json:"plan_phase_order"`
}

func (r PlanVersionNewParamsReplaceAdjustment) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// The definition of a new adjustment to create and add to the plan.
type PlanVersionNewParamsReplaceAdjustmentsAdjustment struct {
	AdjustmentType    param.Field[PlanVersionNewParamsReplaceAdjustmentsAdjustmentAdjustmentType] `json:"adjustment_type,required"`
	AmountDiscount    param.Field[string]                                                         `json:"amount_discount"`
	AppliesToPriceIDs param.Field[interface{}]                                                    `json:"applies_to_price_ids"`
	// When false, this adjustment will be applied to a single price. Otherwise, it
	// will be applied at the invoice level, possibly to multiple prices.
	IsInvoiceLevel param.Field[bool] `json:"is_invoice_level"`
	// The item ID that revenue from this minimum will be attributed to.
	ItemID             param.Field[string]  `json:"item_id"`
	MaximumAmount      param.Field[string]  `json:"maximum_amount"`
	MinimumAmount      param.Field[string]  `json:"minimum_amount"`
	PercentageDiscount param.Field[float64] `json:"percentage_discount"`
	UsageDiscount      param.Field[float64] `json:"usage_discount"`
}

func (r PlanVersionNewParamsReplaceAdjustmentsAdjustment) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r PlanVersionNewParamsReplaceAdjustmentsAdjustment) implementsPlanVersionNewParamsReplaceAdjustmentsAdjustmentUnion() {
}

// The definition of a new adjustment to create and add to the plan.
//
// Satisfied by
// [PlanVersionNewParamsReplaceAdjustmentsAdjustmentNewPercentageDiscount],
// [PlanVersionNewParamsReplaceAdjustmentsAdjustmentNewUsageDiscount],
// [PlanVersionNewParamsReplaceAdjustmentsAdjustmentNewAmountDiscount],
// [PlanVersionNewParamsReplaceAdjustmentsAdjustmentNewMinimum],
// [PlanVersionNewParamsReplaceAdjustmentsAdjustmentNewMaximum],
// [PlanVersionNewParamsReplaceAdjustmentsAdjustment].
type PlanVersionNewParamsReplaceAdjustmentsAdjustmentUnion interface {
	implementsPlanVersionNewParamsReplaceAdjustmentsAdjustmentUnion()
}

type PlanVersionNewParamsReplaceAdjustmentsAdjustmentNewPercentageDiscount struct {
	AdjustmentType     param.Field[PlanVersionNewParamsReplaceAdjustmentsAdjustmentNewPercentageDiscountAdjustmentType] `json:"adjustment_type,required"`
	PercentageDiscount param.Field[float64]                                                                             `json:"percentage_discount,required"`
	// The set of price IDs to which this adjustment applies.
	AppliesToPriceIDs param.Field[[]string] `json:"applies_to_price_ids"`
	// When false, this adjustment will be applied to a single price. Otherwise, it
	// will be applied at the invoice level, possibly to multiple prices.
	IsInvoiceLevel param.Field[bool] `json:"is_invoice_level"`
}

func (r PlanVersionNewParamsReplaceAdjustmentsAdjustmentNewPercentageDiscount) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r PlanVersionNewParamsReplaceAdjustmentsAdjustmentNewPercentageDiscount) implementsPlanVersionNewParamsReplaceAdjustmentsAdjustmentUnion() {
}

type PlanVersionNewParamsReplaceAdjustmentsAdjustmentNewPercentageDiscountAdjustmentType string

const (
	PlanVersionNewParamsReplaceAdjustmentsAdjustmentNewPercentageDiscountAdjustmentTypePercentageDiscount PlanVersionNewParamsReplaceAdjustmentsAdjustmentNewPercentageDiscountAdjustmentType = "percentage_discount"
)

func (r PlanVersionNewParamsReplaceAdjustmentsAdjustmentNewPercentageDiscountAdjustmentType) IsKnown() bool {
	switch r {
	case PlanVersionNewParamsReplaceAdjustmentsAdjustmentNewPercentageDiscountAdjustmentTypePercentageDiscount:
		return true
	}
	return false
}

type PlanVersionNewParamsReplaceAdjustmentsAdjustmentNewUsageDiscount struct {
	AdjustmentType param.Field[PlanVersionNewParamsReplaceAdjustmentsAdjustmentNewUsageDiscountAdjustmentType] `json:"adjustment_type,required"`
	UsageDiscount  param.Field[float64]                                                                        `json:"usage_discount,required"`
	// The set of price IDs to which this adjustment applies.
	AppliesToPriceIDs param.Field[[]string] `json:"applies_to_price_ids"`
	// When false, this adjustment will be applied to a single price. Otherwise, it
	// will be applied at the invoice level, possibly to multiple prices.
	IsInvoiceLevel param.Field[bool] `json:"is_invoice_level"`
}

func (r PlanVersionNewParamsReplaceAdjustmentsAdjustmentNewUsageDiscount) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r PlanVersionNewParamsReplaceAdjustmentsAdjustmentNewUsageDiscount) implementsPlanVersionNewParamsReplaceAdjustmentsAdjustmentUnion() {
}

type PlanVersionNewParamsReplaceAdjustmentsAdjustmentNewUsageDiscountAdjustmentType string

const (
	PlanVersionNewParamsReplaceAdjustmentsAdjustmentNewUsageDiscountAdjustmentTypeUsageDiscount PlanVersionNewParamsReplaceAdjustmentsAdjustmentNewUsageDiscountAdjustmentType = "usage_discount"
)

func (r PlanVersionNewParamsReplaceAdjustmentsAdjustmentNewUsageDiscountAdjustmentType) IsKnown() bool {
	switch r {
	case PlanVersionNewParamsReplaceAdjustmentsAdjustmentNewUsageDiscountAdjustmentTypeUsageDiscount:
		return true
	}
	return false
}

type PlanVersionNewParamsReplaceAdjustmentsAdjustmentNewAmountDiscount struct {
	AdjustmentType param.Field[PlanVersionNewParamsReplaceAdjustmentsAdjustmentNewAmountDiscountAdjustmentType] `json:"adjustment_type,required"`
	AmountDiscount param.Field[string]                                                                          `json:"amount_discount,required"`
	// The set of price IDs to which this adjustment applies.
	AppliesToPriceIDs param.Field[[]string] `json:"applies_to_price_ids"`
	// When false, this adjustment will be applied to a single price. Otherwise, it
	// will be applied at the invoice level, possibly to multiple prices.
	IsInvoiceLevel param.Field[bool] `json:"is_invoice_level"`
}

func (r PlanVersionNewParamsReplaceAdjustmentsAdjustmentNewAmountDiscount) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r PlanVersionNewParamsReplaceAdjustmentsAdjustmentNewAmountDiscount) implementsPlanVersionNewParamsReplaceAdjustmentsAdjustmentUnion() {
}

type PlanVersionNewParamsReplaceAdjustmentsAdjustmentNewAmountDiscountAdjustmentType string

const (
	PlanVersionNewParamsReplaceAdjustmentsAdjustmentNewAmountDiscountAdjustmentTypeAmountDiscount PlanVersionNewParamsReplaceAdjustmentsAdjustmentNewAmountDiscountAdjustmentType = "amount_discount"
)

func (r PlanVersionNewParamsReplaceAdjustmentsAdjustmentNewAmountDiscountAdjustmentType) IsKnown() bool {
	switch r {
	case PlanVersionNewParamsReplaceAdjustmentsAdjustmentNewAmountDiscountAdjustmentTypeAmountDiscount:
		return true
	}
	return false
}

type PlanVersionNewParamsReplaceAdjustmentsAdjustmentNewMinimum struct {
	AdjustmentType param.Field[PlanVersionNewParamsReplaceAdjustmentsAdjustmentNewMinimumAdjustmentType] `json:"adjustment_type,required"`
	// The item ID that revenue from this minimum will be attributed to.
	ItemID        param.Field[string] `json:"item_id,required"`
	MinimumAmount param.Field[string] `json:"minimum_amount,required"`
	// The set of price IDs to which this adjustment applies.
	AppliesToPriceIDs param.Field[[]string] `json:"applies_to_price_ids"`
	// When false, this adjustment will be applied to a single price. Otherwise, it
	// will be applied at the invoice level, possibly to multiple prices.
	IsInvoiceLevel param.Field[bool] `json:"is_invoice_level"`
}

func (r PlanVersionNewParamsReplaceAdjustmentsAdjustmentNewMinimum) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r PlanVersionNewParamsReplaceAdjustmentsAdjustmentNewMinimum) implementsPlanVersionNewParamsReplaceAdjustmentsAdjustmentUnion() {
}

type PlanVersionNewParamsReplaceAdjustmentsAdjustmentNewMinimumAdjustmentType string

const (
	PlanVersionNewParamsReplaceAdjustmentsAdjustmentNewMinimumAdjustmentTypeMinimum PlanVersionNewParamsReplaceAdjustmentsAdjustmentNewMinimumAdjustmentType = "minimum"
)

func (r PlanVersionNewParamsReplaceAdjustmentsAdjustmentNewMinimumAdjustmentType) IsKnown() bool {
	switch r {
	case PlanVersionNewParamsReplaceAdjustmentsAdjustmentNewMinimumAdjustmentTypeMinimum:
		return true
	}
	return false
}

type PlanVersionNewParamsReplaceAdjustmentsAdjustmentNewMaximum struct {
	AdjustmentType param.Field[PlanVersionNewParamsReplaceAdjustmentsAdjustmentNewMaximumAdjustmentType] `json:"adjustment_type,required"`
	MaximumAmount  param.Field[string]                                                                   `json:"maximum_amount,required"`
	// The set of price IDs to which this adjustment applies.
	AppliesToPriceIDs param.Field[[]string] `json:"applies_to_price_ids"`
	// When false, this adjustment will be applied to a single price. Otherwise, it
	// will be applied at the invoice level, possibly to multiple prices.
	IsInvoiceLevel param.Field[bool] `json:"is_invoice_level"`
}

func (r PlanVersionNewParamsReplaceAdjustmentsAdjustmentNewMaximum) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r PlanVersionNewParamsReplaceAdjustmentsAdjustmentNewMaximum) implementsPlanVersionNewParamsReplaceAdjustmentsAdjustmentUnion() {
}

type PlanVersionNewParamsReplaceAdjustmentsAdjustmentNewMaximumAdjustmentType string

const (
	PlanVersionNewParamsReplaceAdjustmentsAdjustmentNewMaximumAdjustmentTypeMaximum PlanVersionNewParamsReplaceAdjustmentsAdjustmentNewMaximumAdjustmentType = "maximum"
)

func (r PlanVersionNewParamsReplaceAdjustmentsAdjustmentNewMaximumAdjustmentType) IsKnown() bool {
	switch r {
	case PlanVersionNewParamsReplaceAdjustmentsAdjustmentNewMaximumAdjustmentTypeMaximum:
		return true
	}
	return false
}

type PlanVersionNewParamsReplaceAdjustmentsAdjustmentAdjustmentType string

const (
	PlanVersionNewParamsReplaceAdjustmentsAdjustmentAdjustmentTypePercentageDiscount PlanVersionNewParamsReplaceAdjustmentsAdjustmentAdjustmentType = "percentage_discount"
	PlanVersionNewParamsReplaceAdjustmentsAdjustmentAdjustmentTypeUsageDiscount      PlanVersionNewParamsReplaceAdjustmentsAdjustmentAdjustmentType = "usage_discount"
	PlanVersionNewParamsReplaceAdjustmentsAdjustmentAdjustmentTypeAmountDiscount     PlanVersionNewParamsReplaceAdjustmentsAdjustmentAdjustmentType = "amount_discount"
	PlanVersionNewParamsReplaceAdjustmentsAdjustmentAdjustmentTypeMinimum            PlanVersionNewParamsReplaceAdjustmentsAdjustmentAdjustmentType = "minimum"
	PlanVersionNewParamsReplaceAdjustmentsAdjustmentAdjustmentTypeMaximum            PlanVersionNewParamsReplaceAdjustmentsAdjustmentAdjustmentType = "maximum"
)

func (r PlanVersionNewParamsReplaceAdjustmentsAdjustmentAdjustmentType) IsKnown() bool {
	switch r {
	case PlanVersionNewParamsReplaceAdjustmentsAdjustmentAdjustmentTypePercentageDiscount, PlanVersionNewParamsReplaceAdjustmentsAdjustmentAdjustmentTypeUsageDiscount, PlanVersionNewParamsReplaceAdjustmentsAdjustmentAdjustmentTypeAmountDiscount, PlanVersionNewParamsReplaceAdjustmentsAdjustmentAdjustmentTypeMinimum, PlanVersionNewParamsReplaceAdjustmentsAdjustmentAdjustmentTypeMaximum:
		return true
	}
	return false
}

type PlanVersionNewParamsReplacePrice struct {
	// The id of the price on the plan to replace in the plan.
	ReplacesPriceID param.Field[string] `json:"replaces_price_id,required"`
	// The allocation price to add to the plan.
	AllocationPrice param.Field[PlanVersionNewParamsReplacePricesAllocationPrice] `json:"allocation_price"`
	// The phase to replace this price from.
	PlanPhaseOrder param.Field[int64] `json:"plan_phase_order"`
	// The price to add to the plan
	Price param.Field[PlanVersionNewParamsReplacePricesPriceUnion] `json:"price"`
}

func (r PlanVersionNewParamsReplacePrice) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// The allocation price to add to the plan.
type PlanVersionNewParamsReplacePricesAllocationPrice struct {
	// An amount of the currency to allocate to the customer at the specified cadence.
	Amount param.Field[string] `json:"amount,required"`
	// The cadence at which to allocate the amount to the customer.
	Cadence param.Field[PlanVersionNewParamsReplacePricesAllocationPriceCadence] `json:"cadence,required"`
	// An ISO 4217 currency string or a custom pricing unit identifier in which to bill
	// this price.
	Currency param.Field[string] `json:"currency,required"`
	// Whether the allocated amount should expire at the end of the cadence or roll
	// over to the next period.
	ExpiresAtEndOfCadence param.Field[bool] `json:"expires_at_end_of_cadence,required"`
}

func (r PlanVersionNewParamsReplacePricesAllocationPrice) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// The cadence at which to allocate the amount to the customer.
type PlanVersionNewParamsReplacePricesAllocationPriceCadence string

const (
	PlanVersionNewParamsReplacePricesAllocationPriceCadenceOneTime    PlanVersionNewParamsReplacePricesAllocationPriceCadence = "one_time"
	PlanVersionNewParamsReplacePricesAllocationPriceCadenceMonthly    PlanVersionNewParamsReplacePricesAllocationPriceCadence = "monthly"
	PlanVersionNewParamsReplacePricesAllocationPriceCadenceQuarterly  PlanVersionNewParamsReplacePricesAllocationPriceCadence = "quarterly"
	PlanVersionNewParamsReplacePricesAllocationPriceCadenceSemiAnnual PlanVersionNewParamsReplacePricesAllocationPriceCadence = "semi_annual"
	PlanVersionNewParamsReplacePricesAllocationPriceCadenceAnnual     PlanVersionNewParamsReplacePricesAllocationPriceCadence = "annual"
	PlanVersionNewParamsReplacePricesAllocationPriceCadenceCustom     PlanVersionNewParamsReplacePricesAllocationPriceCadence = "custom"
)

func (r PlanVersionNewParamsReplacePricesAllocationPriceCadence) IsKnown() bool {
	switch r {
	case PlanVersionNewParamsReplacePricesAllocationPriceCadenceOneTime, PlanVersionNewParamsReplacePricesAllocationPriceCadenceMonthly, PlanVersionNewParamsReplacePricesAllocationPriceCadenceQuarterly, PlanVersionNewParamsReplacePricesAllocationPriceCadenceSemiAnnual, PlanVersionNewParamsReplacePricesAllocationPriceCadenceAnnual, PlanVersionNewParamsReplacePricesAllocationPriceCadenceCustom:
		return true
	}
	return false
}

// The price to add to the plan
type PlanVersionNewParamsReplacePricesPrice struct {
	// The cadence to bill for this price on.
	Cadence param.Field[PlanVersionNewParamsReplacePricesPriceCadence] `json:"cadence,required"`
	// The id of the item the price will be associated with.
	ItemID    param.Field[string]                                          `json:"item_id,required"`
	ModelType param.Field[PlanVersionNewParamsReplacePricesPriceModelType] `json:"model_type,required"`
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
	ConversionRate              param.Field[float64]     `json:"conversion_rate"`
	CumulativeGroupedBulkConfig param.Field[interface{}] `json:"cumulative_grouped_bulk_config"`
	// An ISO 4217 currency string, or custom pricing unit identifier, in which this
	// price is billed.
	Currency                      param.Field[string]      `json:"currency"`
	DimensionalPriceConfiguration param.Field[interface{}] `json:"dimensional_price_configuration"`
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
	InvoiceGroupingKey                    param.Field[string]      `json:"invoice_grouping_key"`
	InvoicingCycleConfiguration           param.Field[interface{}] `json:"invoicing_cycle_configuration"`
	MatrixConfig                          param.Field[interface{}] `json:"matrix_config"`
	MatrixWithAllocationConfig            param.Field[interface{}] `json:"matrix_with_allocation_config"`
	MatrixWithDisplayNameConfig           param.Field[interface{}] `json:"matrix_with_display_name_config"`
	MaxGroupTieredPackageConfig           param.Field[interface{}] `json:"max_group_tiered_package_config"`
	Metadata                              param.Field[interface{}] `json:"metadata"`
	PackageConfig                         param.Field[interface{}] `json:"package_config"`
	PackageWithAllocationConfig           param.Field[interface{}] `json:"package_with_allocation_config"`
	ScalableMatrixWithTieredPricingConfig param.Field[interface{}] `json:"scalable_matrix_with_tiered_pricing_config"`
	ScalableMatrixWithUnitPricingConfig   param.Field[interface{}] `json:"scalable_matrix_with_unit_pricing_config"`
	ThresholdTotalAmountConfig            param.Field[interface{}] `json:"threshold_total_amount_config"`
	TieredBpsConfig                       param.Field[interface{}] `json:"tiered_bps_config"`
	TieredConfig                          param.Field[interface{}] `json:"tiered_config"`
	TieredPackageConfig                   param.Field[interface{}] `json:"tiered_package_config"`
	TieredPackageWithMinimumConfig        param.Field[interface{}] `json:"tiered_package_with_minimum_config"`
	TieredWithMinimumConfig               param.Field[interface{}] `json:"tiered_with_minimum_config"`
	TieredWithProrationConfig             param.Field[interface{}] `json:"tiered_with_proration_config"`
	UnitConfig                            param.Field[interface{}] `json:"unit_config"`
	UnitWithPercentConfig                 param.Field[interface{}] `json:"unit_with_percent_config"`
	UnitWithProrationConfig               param.Field[interface{}] `json:"unit_with_proration_config"`
}

func (r PlanVersionNewParamsReplacePricesPrice) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r PlanVersionNewParamsReplacePricesPrice) implementsPlanVersionNewParamsReplacePricesPriceUnion() {
}

// The price to add to the plan
//
// Satisfied by [PlanVersionNewParamsReplacePricesPriceNewPlanUnitPrice],
// [PlanVersionNewParamsReplacePricesPriceNewPlanPackagePrice],
// [PlanVersionNewParamsReplacePricesPriceNewPlanMatrixPrice],
// [PlanVersionNewParamsReplacePricesPriceNewPlanTieredPrice],
// [PlanVersionNewParamsReplacePricesPriceNewPlanTieredBpsPrice],
// [PlanVersionNewParamsReplacePricesPriceNewPlanBpsPrice],
// [PlanVersionNewParamsReplacePricesPriceNewPlanBulkBpsPrice],
// [PlanVersionNewParamsReplacePricesPriceNewPlanBulkPrice],
// [PlanVersionNewParamsReplacePricesPriceNewPlanThresholdTotalAmountPrice],
// [PlanVersionNewParamsReplacePricesPriceNewPlanTieredPackagePrice],
// [PlanVersionNewParamsReplacePricesPriceNewPlanTieredWithMinimumPrice],
// [PlanVersionNewParamsReplacePricesPriceNewPlanUnitWithPercentPrice],
// [PlanVersionNewParamsReplacePricesPriceNewPlanPackageWithAllocationPrice],
// [PlanVersionNewParamsReplacePricesPriceNewPlanTierWithProrationPrice],
// [PlanVersionNewParamsReplacePricesPriceNewPlanUnitWithProrationPrice],
// [PlanVersionNewParamsReplacePricesPriceNewPlanGroupedAllocationPrice],
// [PlanVersionNewParamsReplacePricesPriceNewPlanGroupedWithProratedMinimumPrice],
// [PlanVersionNewParamsReplacePricesPriceNewPlanGroupedWithMeteredMinimumPrice],
// [PlanVersionNewParamsReplacePricesPriceNewPlanMatrixWithDisplayNamePrice],
// [PlanVersionNewParamsReplacePricesPriceNewPlanBulkWithProrationPrice],
// [PlanVersionNewParamsReplacePricesPriceNewPlanGroupedTieredPackagePrice],
// [PlanVersionNewParamsReplacePricesPriceNewPlanMaxGroupTieredPackagePrice],
// [PlanVersionNewParamsReplacePricesPriceNewPlanScalableMatrixWithUnitPricingPrice],
// [PlanVersionNewParamsReplacePricesPriceNewPlanScalableMatrixWithTieredPricingPrice],
// [PlanVersionNewParamsReplacePricesPriceNewPlanCumulativeGroupedBulkPrice],
// [PlanVersionNewParamsReplacePricesPriceNewPlanTieredPackageWithMinimumPrice],
// [PlanVersionNewParamsReplacePricesPriceNewPlanMatrixWithAllocationPrice],
// [PlanVersionNewParamsReplacePricesPriceNewPlanGroupedTieredPrice],
// [PlanVersionNewParamsReplacePricesPrice].
type PlanVersionNewParamsReplacePricesPriceUnion interface {
	implementsPlanVersionNewParamsReplacePricesPriceUnion()
}

type PlanVersionNewParamsReplacePricesPriceNewPlanUnitPrice struct {
	// The cadence to bill for this price on.
	Cadence param.Field[PlanVersionNewParamsReplacePricesPriceNewPlanUnitPriceCadence] `json:"cadence,required"`
	// The id of the item the price will be associated with.
	ItemID    param.Field[string]                                                          `json:"item_id,required"`
	ModelType param.Field[PlanVersionNewParamsReplacePricesPriceNewPlanUnitPriceModelType] `json:"model_type,required"`
	// The name of the price.
	Name       param.Field[string]                                                           `json:"name,required"`
	UnitConfig param.Field[PlanVersionNewParamsReplacePricesPriceNewPlanUnitPriceUnitConfig] `json:"unit_config,required"`
	// The id of the billable metric for the price. Only needed if the price is
	// usage-based.
	BillableMetricID param.Field[string] `json:"billable_metric_id"`
	// If the Price represents a fixed cost, the price will be billed in-advance if
	// this is true, and in-arrears if this is false.
	BilledInAdvance param.Field[bool] `json:"billed_in_advance"`
	// For custom cadence: specifies the duration of the billing period in days or
	// months.
	BillingCycleConfiguration param.Field[PlanVersionNewParamsReplacePricesPriceNewPlanUnitPriceBillingCycleConfiguration] `json:"billing_cycle_configuration"`
	// The per unit conversion rate of the price currency to the invoicing currency.
	ConversionRate param.Field[float64] `json:"conversion_rate"`
	// An ISO 4217 currency string, or custom pricing unit identifier, in which this
	// price is billed.
	Currency param.Field[string] `json:"currency"`
	// For dimensional price: specifies a price group and dimension values
	DimensionalPriceConfiguration param.Field[PlanVersionNewParamsReplacePricesPriceNewPlanUnitPriceDimensionalPriceConfiguration] `json:"dimensional_price_configuration"`
	// An alias for the price.
	ExternalPriceID param.Field[string] `json:"external_price_id"`
	// If the Price represents a fixed cost, this represents the quantity of units
	// applied.
	FixedPriceQuantity param.Field[float64] `json:"fixed_price_quantity"`
	// The property used to group this price on an invoice
	InvoiceGroupingKey param.Field[string] `json:"invoice_grouping_key"`
	// Within each billing cycle, specifies the cadence at which invoices are produced.
	// If unspecified, a single invoice is produced per billing cycle.
	InvoicingCycleConfiguration param.Field[PlanVersionNewParamsReplacePricesPriceNewPlanUnitPriceInvoicingCycleConfiguration] `json:"invoicing_cycle_configuration"`
	// User-specified key/value pairs for the resource. Individual keys can be removed
	// by setting the value to `null`, and the entire metadata mapping can be cleared
	// by setting `metadata` to `null`.
	Metadata param.Field[map[string]string] `json:"metadata"`
}

func (r PlanVersionNewParamsReplacePricesPriceNewPlanUnitPrice) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r PlanVersionNewParamsReplacePricesPriceNewPlanUnitPrice) implementsPlanVersionNewParamsReplacePricesPriceUnion() {
}

// The cadence to bill for this price on.
type PlanVersionNewParamsReplacePricesPriceNewPlanUnitPriceCadence string

const (
	PlanVersionNewParamsReplacePricesPriceNewPlanUnitPriceCadenceAnnual     PlanVersionNewParamsReplacePricesPriceNewPlanUnitPriceCadence = "annual"
	PlanVersionNewParamsReplacePricesPriceNewPlanUnitPriceCadenceSemiAnnual PlanVersionNewParamsReplacePricesPriceNewPlanUnitPriceCadence = "semi_annual"
	PlanVersionNewParamsReplacePricesPriceNewPlanUnitPriceCadenceMonthly    PlanVersionNewParamsReplacePricesPriceNewPlanUnitPriceCadence = "monthly"
	PlanVersionNewParamsReplacePricesPriceNewPlanUnitPriceCadenceQuarterly  PlanVersionNewParamsReplacePricesPriceNewPlanUnitPriceCadence = "quarterly"
	PlanVersionNewParamsReplacePricesPriceNewPlanUnitPriceCadenceOneTime    PlanVersionNewParamsReplacePricesPriceNewPlanUnitPriceCadence = "one_time"
	PlanVersionNewParamsReplacePricesPriceNewPlanUnitPriceCadenceCustom     PlanVersionNewParamsReplacePricesPriceNewPlanUnitPriceCadence = "custom"
)

func (r PlanVersionNewParamsReplacePricesPriceNewPlanUnitPriceCadence) IsKnown() bool {
	switch r {
	case PlanVersionNewParamsReplacePricesPriceNewPlanUnitPriceCadenceAnnual, PlanVersionNewParamsReplacePricesPriceNewPlanUnitPriceCadenceSemiAnnual, PlanVersionNewParamsReplacePricesPriceNewPlanUnitPriceCadenceMonthly, PlanVersionNewParamsReplacePricesPriceNewPlanUnitPriceCadenceQuarterly, PlanVersionNewParamsReplacePricesPriceNewPlanUnitPriceCadenceOneTime, PlanVersionNewParamsReplacePricesPriceNewPlanUnitPriceCadenceCustom:
		return true
	}
	return false
}

type PlanVersionNewParamsReplacePricesPriceNewPlanUnitPriceModelType string

const (
	PlanVersionNewParamsReplacePricesPriceNewPlanUnitPriceModelTypeUnit PlanVersionNewParamsReplacePricesPriceNewPlanUnitPriceModelType = "unit"
)

func (r PlanVersionNewParamsReplacePricesPriceNewPlanUnitPriceModelType) IsKnown() bool {
	switch r {
	case PlanVersionNewParamsReplacePricesPriceNewPlanUnitPriceModelTypeUnit:
		return true
	}
	return false
}

type PlanVersionNewParamsReplacePricesPriceNewPlanUnitPriceUnitConfig struct {
	// Rate per unit of usage
	UnitAmount param.Field[string] `json:"unit_amount,required"`
}

func (r PlanVersionNewParamsReplacePricesPriceNewPlanUnitPriceUnitConfig) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// For custom cadence: specifies the duration of the billing period in days or
// months.
type PlanVersionNewParamsReplacePricesPriceNewPlanUnitPriceBillingCycleConfiguration struct {
	// The duration of the billing period.
	Duration param.Field[int64] `json:"duration,required"`
	// The unit of billing period duration.
	DurationUnit param.Field[PlanVersionNewParamsReplacePricesPriceNewPlanUnitPriceBillingCycleConfigurationDurationUnit] `json:"duration_unit,required"`
}

func (r PlanVersionNewParamsReplacePricesPriceNewPlanUnitPriceBillingCycleConfiguration) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// The unit of billing period duration.
type PlanVersionNewParamsReplacePricesPriceNewPlanUnitPriceBillingCycleConfigurationDurationUnit string

const (
	PlanVersionNewParamsReplacePricesPriceNewPlanUnitPriceBillingCycleConfigurationDurationUnitDay   PlanVersionNewParamsReplacePricesPriceNewPlanUnitPriceBillingCycleConfigurationDurationUnit = "day"
	PlanVersionNewParamsReplacePricesPriceNewPlanUnitPriceBillingCycleConfigurationDurationUnitMonth PlanVersionNewParamsReplacePricesPriceNewPlanUnitPriceBillingCycleConfigurationDurationUnit = "month"
)

func (r PlanVersionNewParamsReplacePricesPriceNewPlanUnitPriceBillingCycleConfigurationDurationUnit) IsKnown() bool {
	switch r {
	case PlanVersionNewParamsReplacePricesPriceNewPlanUnitPriceBillingCycleConfigurationDurationUnitDay, PlanVersionNewParamsReplacePricesPriceNewPlanUnitPriceBillingCycleConfigurationDurationUnitMonth:
		return true
	}
	return false
}

// For dimensional price: specifies a price group and dimension values
type PlanVersionNewParamsReplacePricesPriceNewPlanUnitPriceDimensionalPriceConfiguration struct {
	// The list of dimension values matching (in order) the dimensions of the price
	// group
	DimensionValues param.Field[[]string] `json:"dimension_values,required"`
	// The id of the dimensional price group to include this price in
	DimensionalPriceGroupID param.Field[string] `json:"dimensional_price_group_id"`
	// The external id of the dimensional price group to include this price in
	ExternalDimensionalPriceGroupID param.Field[string] `json:"external_dimensional_price_group_id"`
}

func (r PlanVersionNewParamsReplacePricesPriceNewPlanUnitPriceDimensionalPriceConfiguration) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Within each billing cycle, specifies the cadence at which invoices are produced.
// If unspecified, a single invoice is produced per billing cycle.
type PlanVersionNewParamsReplacePricesPriceNewPlanUnitPriceInvoicingCycleConfiguration struct {
	// The duration of the billing period.
	Duration param.Field[int64] `json:"duration,required"`
	// The unit of billing period duration.
	DurationUnit param.Field[PlanVersionNewParamsReplacePricesPriceNewPlanUnitPriceInvoicingCycleConfigurationDurationUnit] `json:"duration_unit,required"`
}

func (r PlanVersionNewParamsReplacePricesPriceNewPlanUnitPriceInvoicingCycleConfiguration) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// The unit of billing period duration.
type PlanVersionNewParamsReplacePricesPriceNewPlanUnitPriceInvoicingCycleConfigurationDurationUnit string

const (
	PlanVersionNewParamsReplacePricesPriceNewPlanUnitPriceInvoicingCycleConfigurationDurationUnitDay   PlanVersionNewParamsReplacePricesPriceNewPlanUnitPriceInvoicingCycleConfigurationDurationUnit = "day"
	PlanVersionNewParamsReplacePricesPriceNewPlanUnitPriceInvoicingCycleConfigurationDurationUnitMonth PlanVersionNewParamsReplacePricesPriceNewPlanUnitPriceInvoicingCycleConfigurationDurationUnit = "month"
)

func (r PlanVersionNewParamsReplacePricesPriceNewPlanUnitPriceInvoicingCycleConfigurationDurationUnit) IsKnown() bool {
	switch r {
	case PlanVersionNewParamsReplacePricesPriceNewPlanUnitPriceInvoicingCycleConfigurationDurationUnitDay, PlanVersionNewParamsReplacePricesPriceNewPlanUnitPriceInvoicingCycleConfigurationDurationUnitMonth:
		return true
	}
	return false
}

type PlanVersionNewParamsReplacePricesPriceNewPlanPackagePrice struct {
	// The cadence to bill for this price on.
	Cadence param.Field[PlanVersionNewParamsReplacePricesPriceNewPlanPackagePriceCadence] `json:"cadence,required"`
	// The id of the item the price will be associated with.
	ItemID    param.Field[string]                                                             `json:"item_id,required"`
	ModelType param.Field[PlanVersionNewParamsReplacePricesPriceNewPlanPackagePriceModelType] `json:"model_type,required"`
	// The name of the price.
	Name          param.Field[string]                                                                 `json:"name,required"`
	PackageConfig param.Field[PlanVersionNewParamsReplacePricesPriceNewPlanPackagePricePackageConfig] `json:"package_config,required"`
	// The id of the billable metric for the price. Only needed if the price is
	// usage-based.
	BillableMetricID param.Field[string] `json:"billable_metric_id"`
	// If the Price represents a fixed cost, the price will be billed in-advance if
	// this is true, and in-arrears if this is false.
	BilledInAdvance param.Field[bool] `json:"billed_in_advance"`
	// For custom cadence: specifies the duration of the billing period in days or
	// months.
	BillingCycleConfiguration param.Field[PlanVersionNewParamsReplacePricesPriceNewPlanPackagePriceBillingCycleConfiguration] `json:"billing_cycle_configuration"`
	// The per unit conversion rate of the price currency to the invoicing currency.
	ConversionRate param.Field[float64] `json:"conversion_rate"`
	// An ISO 4217 currency string, or custom pricing unit identifier, in which this
	// price is billed.
	Currency param.Field[string] `json:"currency"`
	// For dimensional price: specifies a price group and dimension values
	DimensionalPriceConfiguration param.Field[PlanVersionNewParamsReplacePricesPriceNewPlanPackagePriceDimensionalPriceConfiguration] `json:"dimensional_price_configuration"`
	// An alias for the price.
	ExternalPriceID param.Field[string] `json:"external_price_id"`
	// If the Price represents a fixed cost, this represents the quantity of units
	// applied.
	FixedPriceQuantity param.Field[float64] `json:"fixed_price_quantity"`
	// The property used to group this price on an invoice
	InvoiceGroupingKey param.Field[string] `json:"invoice_grouping_key"`
	// Within each billing cycle, specifies the cadence at which invoices are produced.
	// If unspecified, a single invoice is produced per billing cycle.
	InvoicingCycleConfiguration param.Field[PlanVersionNewParamsReplacePricesPriceNewPlanPackagePriceInvoicingCycleConfiguration] `json:"invoicing_cycle_configuration"`
	// User-specified key/value pairs for the resource. Individual keys can be removed
	// by setting the value to `null`, and the entire metadata mapping can be cleared
	// by setting `metadata` to `null`.
	Metadata param.Field[map[string]string] `json:"metadata"`
}

func (r PlanVersionNewParamsReplacePricesPriceNewPlanPackagePrice) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r PlanVersionNewParamsReplacePricesPriceNewPlanPackagePrice) implementsPlanVersionNewParamsReplacePricesPriceUnion() {
}

// The cadence to bill for this price on.
type PlanVersionNewParamsReplacePricesPriceNewPlanPackagePriceCadence string

const (
	PlanVersionNewParamsReplacePricesPriceNewPlanPackagePriceCadenceAnnual     PlanVersionNewParamsReplacePricesPriceNewPlanPackagePriceCadence = "annual"
	PlanVersionNewParamsReplacePricesPriceNewPlanPackagePriceCadenceSemiAnnual PlanVersionNewParamsReplacePricesPriceNewPlanPackagePriceCadence = "semi_annual"
	PlanVersionNewParamsReplacePricesPriceNewPlanPackagePriceCadenceMonthly    PlanVersionNewParamsReplacePricesPriceNewPlanPackagePriceCadence = "monthly"
	PlanVersionNewParamsReplacePricesPriceNewPlanPackagePriceCadenceQuarterly  PlanVersionNewParamsReplacePricesPriceNewPlanPackagePriceCadence = "quarterly"
	PlanVersionNewParamsReplacePricesPriceNewPlanPackagePriceCadenceOneTime    PlanVersionNewParamsReplacePricesPriceNewPlanPackagePriceCadence = "one_time"
	PlanVersionNewParamsReplacePricesPriceNewPlanPackagePriceCadenceCustom     PlanVersionNewParamsReplacePricesPriceNewPlanPackagePriceCadence = "custom"
)

func (r PlanVersionNewParamsReplacePricesPriceNewPlanPackagePriceCadence) IsKnown() bool {
	switch r {
	case PlanVersionNewParamsReplacePricesPriceNewPlanPackagePriceCadenceAnnual, PlanVersionNewParamsReplacePricesPriceNewPlanPackagePriceCadenceSemiAnnual, PlanVersionNewParamsReplacePricesPriceNewPlanPackagePriceCadenceMonthly, PlanVersionNewParamsReplacePricesPriceNewPlanPackagePriceCadenceQuarterly, PlanVersionNewParamsReplacePricesPriceNewPlanPackagePriceCadenceOneTime, PlanVersionNewParamsReplacePricesPriceNewPlanPackagePriceCadenceCustom:
		return true
	}
	return false
}

type PlanVersionNewParamsReplacePricesPriceNewPlanPackagePriceModelType string

const (
	PlanVersionNewParamsReplacePricesPriceNewPlanPackagePriceModelTypePackage PlanVersionNewParamsReplacePricesPriceNewPlanPackagePriceModelType = "package"
)

func (r PlanVersionNewParamsReplacePricesPriceNewPlanPackagePriceModelType) IsKnown() bool {
	switch r {
	case PlanVersionNewParamsReplacePricesPriceNewPlanPackagePriceModelTypePackage:
		return true
	}
	return false
}

type PlanVersionNewParamsReplacePricesPriceNewPlanPackagePricePackageConfig struct {
	// A currency amount to rate usage by
	PackageAmount param.Field[string] `json:"package_amount,required"`
	// An integer amount to represent package size. For example, 1000 here would divide
	// usage by 1000 before multiplying by package_amount in rating
	PackageSize param.Field[int64] `json:"package_size,required"`
}

func (r PlanVersionNewParamsReplacePricesPriceNewPlanPackagePricePackageConfig) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// For custom cadence: specifies the duration of the billing period in days or
// months.
type PlanVersionNewParamsReplacePricesPriceNewPlanPackagePriceBillingCycleConfiguration struct {
	// The duration of the billing period.
	Duration param.Field[int64] `json:"duration,required"`
	// The unit of billing period duration.
	DurationUnit param.Field[PlanVersionNewParamsReplacePricesPriceNewPlanPackagePriceBillingCycleConfigurationDurationUnit] `json:"duration_unit,required"`
}

func (r PlanVersionNewParamsReplacePricesPriceNewPlanPackagePriceBillingCycleConfiguration) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// The unit of billing period duration.
type PlanVersionNewParamsReplacePricesPriceNewPlanPackagePriceBillingCycleConfigurationDurationUnit string

const (
	PlanVersionNewParamsReplacePricesPriceNewPlanPackagePriceBillingCycleConfigurationDurationUnitDay   PlanVersionNewParamsReplacePricesPriceNewPlanPackagePriceBillingCycleConfigurationDurationUnit = "day"
	PlanVersionNewParamsReplacePricesPriceNewPlanPackagePriceBillingCycleConfigurationDurationUnitMonth PlanVersionNewParamsReplacePricesPriceNewPlanPackagePriceBillingCycleConfigurationDurationUnit = "month"
)

func (r PlanVersionNewParamsReplacePricesPriceNewPlanPackagePriceBillingCycleConfigurationDurationUnit) IsKnown() bool {
	switch r {
	case PlanVersionNewParamsReplacePricesPriceNewPlanPackagePriceBillingCycleConfigurationDurationUnitDay, PlanVersionNewParamsReplacePricesPriceNewPlanPackagePriceBillingCycleConfigurationDurationUnitMonth:
		return true
	}
	return false
}

// For dimensional price: specifies a price group and dimension values
type PlanVersionNewParamsReplacePricesPriceNewPlanPackagePriceDimensionalPriceConfiguration struct {
	// The list of dimension values matching (in order) the dimensions of the price
	// group
	DimensionValues param.Field[[]string] `json:"dimension_values,required"`
	// The id of the dimensional price group to include this price in
	DimensionalPriceGroupID param.Field[string] `json:"dimensional_price_group_id"`
	// The external id of the dimensional price group to include this price in
	ExternalDimensionalPriceGroupID param.Field[string] `json:"external_dimensional_price_group_id"`
}

func (r PlanVersionNewParamsReplacePricesPriceNewPlanPackagePriceDimensionalPriceConfiguration) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Within each billing cycle, specifies the cadence at which invoices are produced.
// If unspecified, a single invoice is produced per billing cycle.
type PlanVersionNewParamsReplacePricesPriceNewPlanPackagePriceInvoicingCycleConfiguration struct {
	// The duration of the billing period.
	Duration param.Field[int64] `json:"duration,required"`
	// The unit of billing period duration.
	DurationUnit param.Field[PlanVersionNewParamsReplacePricesPriceNewPlanPackagePriceInvoicingCycleConfigurationDurationUnit] `json:"duration_unit,required"`
}

func (r PlanVersionNewParamsReplacePricesPriceNewPlanPackagePriceInvoicingCycleConfiguration) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// The unit of billing period duration.
type PlanVersionNewParamsReplacePricesPriceNewPlanPackagePriceInvoicingCycleConfigurationDurationUnit string

const (
	PlanVersionNewParamsReplacePricesPriceNewPlanPackagePriceInvoicingCycleConfigurationDurationUnitDay   PlanVersionNewParamsReplacePricesPriceNewPlanPackagePriceInvoicingCycleConfigurationDurationUnit = "day"
	PlanVersionNewParamsReplacePricesPriceNewPlanPackagePriceInvoicingCycleConfigurationDurationUnitMonth PlanVersionNewParamsReplacePricesPriceNewPlanPackagePriceInvoicingCycleConfigurationDurationUnit = "month"
)

func (r PlanVersionNewParamsReplacePricesPriceNewPlanPackagePriceInvoicingCycleConfigurationDurationUnit) IsKnown() bool {
	switch r {
	case PlanVersionNewParamsReplacePricesPriceNewPlanPackagePriceInvoicingCycleConfigurationDurationUnitDay, PlanVersionNewParamsReplacePricesPriceNewPlanPackagePriceInvoicingCycleConfigurationDurationUnitMonth:
		return true
	}
	return false
}

type PlanVersionNewParamsReplacePricesPriceNewPlanMatrixPrice struct {
	// The cadence to bill for this price on.
	Cadence param.Field[PlanVersionNewParamsReplacePricesPriceNewPlanMatrixPriceCadence] `json:"cadence,required"`
	// The id of the item the price will be associated with.
	ItemID       param.Field[string]                                                               `json:"item_id,required"`
	MatrixConfig param.Field[PlanVersionNewParamsReplacePricesPriceNewPlanMatrixPriceMatrixConfig] `json:"matrix_config,required"`
	ModelType    param.Field[PlanVersionNewParamsReplacePricesPriceNewPlanMatrixPriceModelType]    `json:"model_type,required"`
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
	BillingCycleConfiguration param.Field[PlanVersionNewParamsReplacePricesPriceNewPlanMatrixPriceBillingCycleConfiguration] `json:"billing_cycle_configuration"`
	// The per unit conversion rate of the price currency to the invoicing currency.
	ConversionRate param.Field[float64] `json:"conversion_rate"`
	// An ISO 4217 currency string, or custom pricing unit identifier, in which this
	// price is billed.
	Currency param.Field[string] `json:"currency"`
	// For dimensional price: specifies a price group and dimension values
	DimensionalPriceConfiguration param.Field[PlanVersionNewParamsReplacePricesPriceNewPlanMatrixPriceDimensionalPriceConfiguration] `json:"dimensional_price_configuration"`
	// An alias for the price.
	ExternalPriceID param.Field[string] `json:"external_price_id"`
	// If the Price represents a fixed cost, this represents the quantity of units
	// applied.
	FixedPriceQuantity param.Field[float64] `json:"fixed_price_quantity"`
	// The property used to group this price on an invoice
	InvoiceGroupingKey param.Field[string] `json:"invoice_grouping_key"`
	// Within each billing cycle, specifies the cadence at which invoices are produced.
	// If unspecified, a single invoice is produced per billing cycle.
	InvoicingCycleConfiguration param.Field[PlanVersionNewParamsReplacePricesPriceNewPlanMatrixPriceInvoicingCycleConfiguration] `json:"invoicing_cycle_configuration"`
	// User-specified key/value pairs for the resource. Individual keys can be removed
	// by setting the value to `null`, and the entire metadata mapping can be cleared
	// by setting `metadata` to `null`.
	Metadata param.Field[map[string]string] `json:"metadata"`
}

func (r PlanVersionNewParamsReplacePricesPriceNewPlanMatrixPrice) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r PlanVersionNewParamsReplacePricesPriceNewPlanMatrixPrice) implementsPlanVersionNewParamsReplacePricesPriceUnion() {
}

// The cadence to bill for this price on.
type PlanVersionNewParamsReplacePricesPriceNewPlanMatrixPriceCadence string

const (
	PlanVersionNewParamsReplacePricesPriceNewPlanMatrixPriceCadenceAnnual     PlanVersionNewParamsReplacePricesPriceNewPlanMatrixPriceCadence = "annual"
	PlanVersionNewParamsReplacePricesPriceNewPlanMatrixPriceCadenceSemiAnnual PlanVersionNewParamsReplacePricesPriceNewPlanMatrixPriceCadence = "semi_annual"
	PlanVersionNewParamsReplacePricesPriceNewPlanMatrixPriceCadenceMonthly    PlanVersionNewParamsReplacePricesPriceNewPlanMatrixPriceCadence = "monthly"
	PlanVersionNewParamsReplacePricesPriceNewPlanMatrixPriceCadenceQuarterly  PlanVersionNewParamsReplacePricesPriceNewPlanMatrixPriceCadence = "quarterly"
	PlanVersionNewParamsReplacePricesPriceNewPlanMatrixPriceCadenceOneTime    PlanVersionNewParamsReplacePricesPriceNewPlanMatrixPriceCadence = "one_time"
	PlanVersionNewParamsReplacePricesPriceNewPlanMatrixPriceCadenceCustom     PlanVersionNewParamsReplacePricesPriceNewPlanMatrixPriceCadence = "custom"
)

func (r PlanVersionNewParamsReplacePricesPriceNewPlanMatrixPriceCadence) IsKnown() bool {
	switch r {
	case PlanVersionNewParamsReplacePricesPriceNewPlanMatrixPriceCadenceAnnual, PlanVersionNewParamsReplacePricesPriceNewPlanMatrixPriceCadenceSemiAnnual, PlanVersionNewParamsReplacePricesPriceNewPlanMatrixPriceCadenceMonthly, PlanVersionNewParamsReplacePricesPriceNewPlanMatrixPriceCadenceQuarterly, PlanVersionNewParamsReplacePricesPriceNewPlanMatrixPriceCadenceOneTime, PlanVersionNewParamsReplacePricesPriceNewPlanMatrixPriceCadenceCustom:
		return true
	}
	return false
}

type PlanVersionNewParamsReplacePricesPriceNewPlanMatrixPriceMatrixConfig struct {
	// Default per unit rate for any usage not bucketed into a specified matrix_value
	DefaultUnitAmount param.Field[string] `json:"default_unit_amount,required"`
	// One or two event property values to evaluate matrix groups by
	Dimensions param.Field[[]string] `json:"dimensions,required"`
	// Matrix values for specified matrix grouping keys
	MatrixValues param.Field[[]PlanVersionNewParamsReplacePricesPriceNewPlanMatrixPriceMatrixConfigMatrixValue] `json:"matrix_values,required"`
}

func (r PlanVersionNewParamsReplacePricesPriceNewPlanMatrixPriceMatrixConfig) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type PlanVersionNewParamsReplacePricesPriceNewPlanMatrixPriceMatrixConfigMatrixValue struct {
	// One or two matrix keys to filter usage to this Matrix value by. For example,
	// ["region", "tier"] could be used to filter cloud usage by a cloud region and an
	// instance tier.
	DimensionValues param.Field[[]string] `json:"dimension_values,required"`
	// Unit price for the specified dimension_values
	UnitAmount param.Field[string] `json:"unit_amount,required"`
}

func (r PlanVersionNewParamsReplacePricesPriceNewPlanMatrixPriceMatrixConfigMatrixValue) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type PlanVersionNewParamsReplacePricesPriceNewPlanMatrixPriceModelType string

const (
	PlanVersionNewParamsReplacePricesPriceNewPlanMatrixPriceModelTypeMatrix PlanVersionNewParamsReplacePricesPriceNewPlanMatrixPriceModelType = "matrix"
)

func (r PlanVersionNewParamsReplacePricesPriceNewPlanMatrixPriceModelType) IsKnown() bool {
	switch r {
	case PlanVersionNewParamsReplacePricesPriceNewPlanMatrixPriceModelTypeMatrix:
		return true
	}
	return false
}

// For custom cadence: specifies the duration of the billing period in days or
// months.
type PlanVersionNewParamsReplacePricesPriceNewPlanMatrixPriceBillingCycleConfiguration struct {
	// The duration of the billing period.
	Duration param.Field[int64] `json:"duration,required"`
	// The unit of billing period duration.
	DurationUnit param.Field[PlanVersionNewParamsReplacePricesPriceNewPlanMatrixPriceBillingCycleConfigurationDurationUnit] `json:"duration_unit,required"`
}

func (r PlanVersionNewParamsReplacePricesPriceNewPlanMatrixPriceBillingCycleConfiguration) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// The unit of billing period duration.
type PlanVersionNewParamsReplacePricesPriceNewPlanMatrixPriceBillingCycleConfigurationDurationUnit string

const (
	PlanVersionNewParamsReplacePricesPriceNewPlanMatrixPriceBillingCycleConfigurationDurationUnitDay   PlanVersionNewParamsReplacePricesPriceNewPlanMatrixPriceBillingCycleConfigurationDurationUnit = "day"
	PlanVersionNewParamsReplacePricesPriceNewPlanMatrixPriceBillingCycleConfigurationDurationUnitMonth PlanVersionNewParamsReplacePricesPriceNewPlanMatrixPriceBillingCycleConfigurationDurationUnit = "month"
)

func (r PlanVersionNewParamsReplacePricesPriceNewPlanMatrixPriceBillingCycleConfigurationDurationUnit) IsKnown() bool {
	switch r {
	case PlanVersionNewParamsReplacePricesPriceNewPlanMatrixPriceBillingCycleConfigurationDurationUnitDay, PlanVersionNewParamsReplacePricesPriceNewPlanMatrixPriceBillingCycleConfigurationDurationUnitMonth:
		return true
	}
	return false
}

// For dimensional price: specifies a price group and dimension values
type PlanVersionNewParamsReplacePricesPriceNewPlanMatrixPriceDimensionalPriceConfiguration struct {
	// The list of dimension values matching (in order) the dimensions of the price
	// group
	DimensionValues param.Field[[]string] `json:"dimension_values,required"`
	// The id of the dimensional price group to include this price in
	DimensionalPriceGroupID param.Field[string] `json:"dimensional_price_group_id"`
	// The external id of the dimensional price group to include this price in
	ExternalDimensionalPriceGroupID param.Field[string] `json:"external_dimensional_price_group_id"`
}

func (r PlanVersionNewParamsReplacePricesPriceNewPlanMatrixPriceDimensionalPriceConfiguration) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Within each billing cycle, specifies the cadence at which invoices are produced.
// If unspecified, a single invoice is produced per billing cycle.
type PlanVersionNewParamsReplacePricesPriceNewPlanMatrixPriceInvoicingCycleConfiguration struct {
	// The duration of the billing period.
	Duration param.Field[int64] `json:"duration,required"`
	// The unit of billing period duration.
	DurationUnit param.Field[PlanVersionNewParamsReplacePricesPriceNewPlanMatrixPriceInvoicingCycleConfigurationDurationUnit] `json:"duration_unit,required"`
}

func (r PlanVersionNewParamsReplacePricesPriceNewPlanMatrixPriceInvoicingCycleConfiguration) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// The unit of billing period duration.
type PlanVersionNewParamsReplacePricesPriceNewPlanMatrixPriceInvoicingCycleConfigurationDurationUnit string

const (
	PlanVersionNewParamsReplacePricesPriceNewPlanMatrixPriceInvoicingCycleConfigurationDurationUnitDay   PlanVersionNewParamsReplacePricesPriceNewPlanMatrixPriceInvoicingCycleConfigurationDurationUnit = "day"
	PlanVersionNewParamsReplacePricesPriceNewPlanMatrixPriceInvoicingCycleConfigurationDurationUnitMonth PlanVersionNewParamsReplacePricesPriceNewPlanMatrixPriceInvoicingCycleConfigurationDurationUnit = "month"
)

func (r PlanVersionNewParamsReplacePricesPriceNewPlanMatrixPriceInvoicingCycleConfigurationDurationUnit) IsKnown() bool {
	switch r {
	case PlanVersionNewParamsReplacePricesPriceNewPlanMatrixPriceInvoicingCycleConfigurationDurationUnitDay, PlanVersionNewParamsReplacePricesPriceNewPlanMatrixPriceInvoicingCycleConfigurationDurationUnitMonth:
		return true
	}
	return false
}

type PlanVersionNewParamsReplacePricesPriceNewPlanTieredPrice struct {
	// The cadence to bill for this price on.
	Cadence param.Field[PlanVersionNewParamsReplacePricesPriceNewPlanTieredPriceCadence] `json:"cadence,required"`
	// The id of the item the price will be associated with.
	ItemID    param.Field[string]                                                            `json:"item_id,required"`
	ModelType param.Field[PlanVersionNewParamsReplacePricesPriceNewPlanTieredPriceModelType] `json:"model_type,required"`
	// The name of the price.
	Name         param.Field[string]                                                               `json:"name,required"`
	TieredConfig param.Field[PlanVersionNewParamsReplacePricesPriceNewPlanTieredPriceTieredConfig] `json:"tiered_config,required"`
	// The id of the billable metric for the price. Only needed if the price is
	// usage-based.
	BillableMetricID param.Field[string] `json:"billable_metric_id"`
	// If the Price represents a fixed cost, the price will be billed in-advance if
	// this is true, and in-arrears if this is false.
	BilledInAdvance param.Field[bool] `json:"billed_in_advance"`
	// For custom cadence: specifies the duration of the billing period in days or
	// months.
	BillingCycleConfiguration param.Field[PlanVersionNewParamsReplacePricesPriceNewPlanTieredPriceBillingCycleConfiguration] `json:"billing_cycle_configuration"`
	// The per unit conversion rate of the price currency to the invoicing currency.
	ConversionRate param.Field[float64] `json:"conversion_rate"`
	// An ISO 4217 currency string, or custom pricing unit identifier, in which this
	// price is billed.
	Currency param.Field[string] `json:"currency"`
	// For dimensional price: specifies a price group and dimension values
	DimensionalPriceConfiguration param.Field[PlanVersionNewParamsReplacePricesPriceNewPlanTieredPriceDimensionalPriceConfiguration] `json:"dimensional_price_configuration"`
	// An alias for the price.
	ExternalPriceID param.Field[string] `json:"external_price_id"`
	// If the Price represents a fixed cost, this represents the quantity of units
	// applied.
	FixedPriceQuantity param.Field[float64] `json:"fixed_price_quantity"`
	// The property used to group this price on an invoice
	InvoiceGroupingKey param.Field[string] `json:"invoice_grouping_key"`
	// Within each billing cycle, specifies the cadence at which invoices are produced.
	// If unspecified, a single invoice is produced per billing cycle.
	InvoicingCycleConfiguration param.Field[PlanVersionNewParamsReplacePricesPriceNewPlanTieredPriceInvoicingCycleConfiguration] `json:"invoicing_cycle_configuration"`
	// User-specified key/value pairs for the resource. Individual keys can be removed
	// by setting the value to `null`, and the entire metadata mapping can be cleared
	// by setting `metadata` to `null`.
	Metadata param.Field[map[string]string] `json:"metadata"`
}

func (r PlanVersionNewParamsReplacePricesPriceNewPlanTieredPrice) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r PlanVersionNewParamsReplacePricesPriceNewPlanTieredPrice) implementsPlanVersionNewParamsReplacePricesPriceUnion() {
}

// The cadence to bill for this price on.
type PlanVersionNewParamsReplacePricesPriceNewPlanTieredPriceCadence string

const (
	PlanVersionNewParamsReplacePricesPriceNewPlanTieredPriceCadenceAnnual     PlanVersionNewParamsReplacePricesPriceNewPlanTieredPriceCadence = "annual"
	PlanVersionNewParamsReplacePricesPriceNewPlanTieredPriceCadenceSemiAnnual PlanVersionNewParamsReplacePricesPriceNewPlanTieredPriceCadence = "semi_annual"
	PlanVersionNewParamsReplacePricesPriceNewPlanTieredPriceCadenceMonthly    PlanVersionNewParamsReplacePricesPriceNewPlanTieredPriceCadence = "monthly"
	PlanVersionNewParamsReplacePricesPriceNewPlanTieredPriceCadenceQuarterly  PlanVersionNewParamsReplacePricesPriceNewPlanTieredPriceCadence = "quarterly"
	PlanVersionNewParamsReplacePricesPriceNewPlanTieredPriceCadenceOneTime    PlanVersionNewParamsReplacePricesPriceNewPlanTieredPriceCadence = "one_time"
	PlanVersionNewParamsReplacePricesPriceNewPlanTieredPriceCadenceCustom     PlanVersionNewParamsReplacePricesPriceNewPlanTieredPriceCadence = "custom"
)

func (r PlanVersionNewParamsReplacePricesPriceNewPlanTieredPriceCadence) IsKnown() bool {
	switch r {
	case PlanVersionNewParamsReplacePricesPriceNewPlanTieredPriceCadenceAnnual, PlanVersionNewParamsReplacePricesPriceNewPlanTieredPriceCadenceSemiAnnual, PlanVersionNewParamsReplacePricesPriceNewPlanTieredPriceCadenceMonthly, PlanVersionNewParamsReplacePricesPriceNewPlanTieredPriceCadenceQuarterly, PlanVersionNewParamsReplacePricesPriceNewPlanTieredPriceCadenceOneTime, PlanVersionNewParamsReplacePricesPriceNewPlanTieredPriceCadenceCustom:
		return true
	}
	return false
}

type PlanVersionNewParamsReplacePricesPriceNewPlanTieredPriceModelType string

const (
	PlanVersionNewParamsReplacePricesPriceNewPlanTieredPriceModelTypeTiered PlanVersionNewParamsReplacePricesPriceNewPlanTieredPriceModelType = "tiered"
)

func (r PlanVersionNewParamsReplacePricesPriceNewPlanTieredPriceModelType) IsKnown() bool {
	switch r {
	case PlanVersionNewParamsReplacePricesPriceNewPlanTieredPriceModelTypeTiered:
		return true
	}
	return false
}

type PlanVersionNewParamsReplacePricesPriceNewPlanTieredPriceTieredConfig struct {
	// Tiers for rating based on total usage quantities into the specified tier
	Tiers param.Field[[]PlanVersionNewParamsReplacePricesPriceNewPlanTieredPriceTieredConfigTier] `json:"tiers,required"`
}

func (r PlanVersionNewParamsReplacePricesPriceNewPlanTieredPriceTieredConfig) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type PlanVersionNewParamsReplacePricesPriceNewPlanTieredPriceTieredConfigTier struct {
	// Exclusive tier starting value
	FirstUnit param.Field[float64] `json:"first_unit,required"`
	// Amount per unit
	UnitAmount param.Field[string] `json:"unit_amount,required"`
	// Inclusive tier ending value. If null, this is treated as the last tier
	LastUnit param.Field[float64] `json:"last_unit"`
}

func (r PlanVersionNewParamsReplacePricesPriceNewPlanTieredPriceTieredConfigTier) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// For custom cadence: specifies the duration of the billing period in days or
// months.
type PlanVersionNewParamsReplacePricesPriceNewPlanTieredPriceBillingCycleConfiguration struct {
	// The duration of the billing period.
	Duration param.Field[int64] `json:"duration,required"`
	// The unit of billing period duration.
	DurationUnit param.Field[PlanVersionNewParamsReplacePricesPriceNewPlanTieredPriceBillingCycleConfigurationDurationUnit] `json:"duration_unit,required"`
}

func (r PlanVersionNewParamsReplacePricesPriceNewPlanTieredPriceBillingCycleConfiguration) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// The unit of billing period duration.
type PlanVersionNewParamsReplacePricesPriceNewPlanTieredPriceBillingCycleConfigurationDurationUnit string

const (
	PlanVersionNewParamsReplacePricesPriceNewPlanTieredPriceBillingCycleConfigurationDurationUnitDay   PlanVersionNewParamsReplacePricesPriceNewPlanTieredPriceBillingCycleConfigurationDurationUnit = "day"
	PlanVersionNewParamsReplacePricesPriceNewPlanTieredPriceBillingCycleConfigurationDurationUnitMonth PlanVersionNewParamsReplacePricesPriceNewPlanTieredPriceBillingCycleConfigurationDurationUnit = "month"
)

func (r PlanVersionNewParamsReplacePricesPriceNewPlanTieredPriceBillingCycleConfigurationDurationUnit) IsKnown() bool {
	switch r {
	case PlanVersionNewParamsReplacePricesPriceNewPlanTieredPriceBillingCycleConfigurationDurationUnitDay, PlanVersionNewParamsReplacePricesPriceNewPlanTieredPriceBillingCycleConfigurationDurationUnitMonth:
		return true
	}
	return false
}

// For dimensional price: specifies a price group and dimension values
type PlanVersionNewParamsReplacePricesPriceNewPlanTieredPriceDimensionalPriceConfiguration struct {
	// The list of dimension values matching (in order) the dimensions of the price
	// group
	DimensionValues param.Field[[]string] `json:"dimension_values,required"`
	// The id of the dimensional price group to include this price in
	DimensionalPriceGroupID param.Field[string] `json:"dimensional_price_group_id"`
	// The external id of the dimensional price group to include this price in
	ExternalDimensionalPriceGroupID param.Field[string] `json:"external_dimensional_price_group_id"`
}

func (r PlanVersionNewParamsReplacePricesPriceNewPlanTieredPriceDimensionalPriceConfiguration) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Within each billing cycle, specifies the cadence at which invoices are produced.
// If unspecified, a single invoice is produced per billing cycle.
type PlanVersionNewParamsReplacePricesPriceNewPlanTieredPriceInvoicingCycleConfiguration struct {
	// The duration of the billing period.
	Duration param.Field[int64] `json:"duration,required"`
	// The unit of billing period duration.
	DurationUnit param.Field[PlanVersionNewParamsReplacePricesPriceNewPlanTieredPriceInvoicingCycleConfigurationDurationUnit] `json:"duration_unit,required"`
}

func (r PlanVersionNewParamsReplacePricesPriceNewPlanTieredPriceInvoicingCycleConfiguration) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// The unit of billing period duration.
type PlanVersionNewParamsReplacePricesPriceNewPlanTieredPriceInvoicingCycleConfigurationDurationUnit string

const (
	PlanVersionNewParamsReplacePricesPriceNewPlanTieredPriceInvoicingCycleConfigurationDurationUnitDay   PlanVersionNewParamsReplacePricesPriceNewPlanTieredPriceInvoicingCycleConfigurationDurationUnit = "day"
	PlanVersionNewParamsReplacePricesPriceNewPlanTieredPriceInvoicingCycleConfigurationDurationUnitMonth PlanVersionNewParamsReplacePricesPriceNewPlanTieredPriceInvoicingCycleConfigurationDurationUnit = "month"
)

func (r PlanVersionNewParamsReplacePricesPriceNewPlanTieredPriceInvoicingCycleConfigurationDurationUnit) IsKnown() bool {
	switch r {
	case PlanVersionNewParamsReplacePricesPriceNewPlanTieredPriceInvoicingCycleConfigurationDurationUnitDay, PlanVersionNewParamsReplacePricesPriceNewPlanTieredPriceInvoicingCycleConfigurationDurationUnitMonth:
		return true
	}
	return false
}

type PlanVersionNewParamsReplacePricesPriceNewPlanTieredBpsPrice struct {
	// The cadence to bill for this price on.
	Cadence param.Field[PlanVersionNewParamsReplacePricesPriceNewPlanTieredBpsPriceCadence] `json:"cadence,required"`
	// The id of the item the price will be associated with.
	ItemID    param.Field[string]                                                               `json:"item_id,required"`
	ModelType param.Field[PlanVersionNewParamsReplacePricesPriceNewPlanTieredBpsPriceModelType] `json:"model_type,required"`
	// The name of the price.
	Name            param.Field[string]                                                                     `json:"name,required"`
	TieredBpsConfig param.Field[PlanVersionNewParamsReplacePricesPriceNewPlanTieredBpsPriceTieredBpsConfig] `json:"tiered_bps_config,required"`
	// The id of the billable metric for the price. Only needed if the price is
	// usage-based.
	BillableMetricID param.Field[string] `json:"billable_metric_id"`
	// If the Price represents a fixed cost, the price will be billed in-advance if
	// this is true, and in-arrears if this is false.
	BilledInAdvance param.Field[bool] `json:"billed_in_advance"`
	// For custom cadence: specifies the duration of the billing period in days or
	// months.
	BillingCycleConfiguration param.Field[PlanVersionNewParamsReplacePricesPriceNewPlanTieredBpsPriceBillingCycleConfiguration] `json:"billing_cycle_configuration"`
	// The per unit conversion rate of the price currency to the invoicing currency.
	ConversionRate param.Field[float64] `json:"conversion_rate"`
	// An ISO 4217 currency string, or custom pricing unit identifier, in which this
	// price is billed.
	Currency param.Field[string] `json:"currency"`
	// For dimensional price: specifies a price group and dimension values
	DimensionalPriceConfiguration param.Field[PlanVersionNewParamsReplacePricesPriceNewPlanTieredBpsPriceDimensionalPriceConfiguration] `json:"dimensional_price_configuration"`
	// An alias for the price.
	ExternalPriceID param.Field[string] `json:"external_price_id"`
	// If the Price represents a fixed cost, this represents the quantity of units
	// applied.
	FixedPriceQuantity param.Field[float64] `json:"fixed_price_quantity"`
	// The property used to group this price on an invoice
	InvoiceGroupingKey param.Field[string] `json:"invoice_grouping_key"`
	// Within each billing cycle, specifies the cadence at which invoices are produced.
	// If unspecified, a single invoice is produced per billing cycle.
	InvoicingCycleConfiguration param.Field[PlanVersionNewParamsReplacePricesPriceNewPlanTieredBpsPriceInvoicingCycleConfiguration] `json:"invoicing_cycle_configuration"`
	// User-specified key/value pairs for the resource. Individual keys can be removed
	// by setting the value to `null`, and the entire metadata mapping can be cleared
	// by setting `metadata` to `null`.
	Metadata param.Field[map[string]string] `json:"metadata"`
}

func (r PlanVersionNewParamsReplacePricesPriceNewPlanTieredBpsPrice) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r PlanVersionNewParamsReplacePricesPriceNewPlanTieredBpsPrice) implementsPlanVersionNewParamsReplacePricesPriceUnion() {
}

// The cadence to bill for this price on.
type PlanVersionNewParamsReplacePricesPriceNewPlanTieredBpsPriceCadence string

const (
	PlanVersionNewParamsReplacePricesPriceNewPlanTieredBpsPriceCadenceAnnual     PlanVersionNewParamsReplacePricesPriceNewPlanTieredBpsPriceCadence = "annual"
	PlanVersionNewParamsReplacePricesPriceNewPlanTieredBpsPriceCadenceSemiAnnual PlanVersionNewParamsReplacePricesPriceNewPlanTieredBpsPriceCadence = "semi_annual"
	PlanVersionNewParamsReplacePricesPriceNewPlanTieredBpsPriceCadenceMonthly    PlanVersionNewParamsReplacePricesPriceNewPlanTieredBpsPriceCadence = "monthly"
	PlanVersionNewParamsReplacePricesPriceNewPlanTieredBpsPriceCadenceQuarterly  PlanVersionNewParamsReplacePricesPriceNewPlanTieredBpsPriceCadence = "quarterly"
	PlanVersionNewParamsReplacePricesPriceNewPlanTieredBpsPriceCadenceOneTime    PlanVersionNewParamsReplacePricesPriceNewPlanTieredBpsPriceCadence = "one_time"
	PlanVersionNewParamsReplacePricesPriceNewPlanTieredBpsPriceCadenceCustom     PlanVersionNewParamsReplacePricesPriceNewPlanTieredBpsPriceCadence = "custom"
)

func (r PlanVersionNewParamsReplacePricesPriceNewPlanTieredBpsPriceCadence) IsKnown() bool {
	switch r {
	case PlanVersionNewParamsReplacePricesPriceNewPlanTieredBpsPriceCadenceAnnual, PlanVersionNewParamsReplacePricesPriceNewPlanTieredBpsPriceCadenceSemiAnnual, PlanVersionNewParamsReplacePricesPriceNewPlanTieredBpsPriceCadenceMonthly, PlanVersionNewParamsReplacePricesPriceNewPlanTieredBpsPriceCadenceQuarterly, PlanVersionNewParamsReplacePricesPriceNewPlanTieredBpsPriceCadenceOneTime, PlanVersionNewParamsReplacePricesPriceNewPlanTieredBpsPriceCadenceCustom:
		return true
	}
	return false
}

type PlanVersionNewParamsReplacePricesPriceNewPlanTieredBpsPriceModelType string

const (
	PlanVersionNewParamsReplacePricesPriceNewPlanTieredBpsPriceModelTypeTieredBps PlanVersionNewParamsReplacePricesPriceNewPlanTieredBpsPriceModelType = "tiered_bps"
)

func (r PlanVersionNewParamsReplacePricesPriceNewPlanTieredBpsPriceModelType) IsKnown() bool {
	switch r {
	case PlanVersionNewParamsReplacePricesPriceNewPlanTieredBpsPriceModelTypeTieredBps:
		return true
	}
	return false
}

type PlanVersionNewParamsReplacePricesPriceNewPlanTieredBpsPriceTieredBpsConfig struct {
	// Tiers for a Graduated BPS pricing model, where usage is bucketed into specified
	// tiers
	Tiers param.Field[[]PlanVersionNewParamsReplacePricesPriceNewPlanTieredBpsPriceTieredBpsConfigTier] `json:"tiers,required"`
}

func (r PlanVersionNewParamsReplacePricesPriceNewPlanTieredBpsPriceTieredBpsConfig) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type PlanVersionNewParamsReplacePricesPriceNewPlanTieredBpsPriceTieredBpsConfigTier struct {
	// Per-event basis point rate
	Bps param.Field[float64] `json:"bps,required"`
	// Exclusive tier starting value
	MinimumAmount param.Field[string] `json:"minimum_amount,required"`
	// Inclusive tier ending value
	MaximumAmount param.Field[string] `json:"maximum_amount"`
	// Per unit maximum to charge
	PerUnitMaximum param.Field[string] `json:"per_unit_maximum"`
}

func (r PlanVersionNewParamsReplacePricesPriceNewPlanTieredBpsPriceTieredBpsConfigTier) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// For custom cadence: specifies the duration of the billing period in days or
// months.
type PlanVersionNewParamsReplacePricesPriceNewPlanTieredBpsPriceBillingCycleConfiguration struct {
	// The duration of the billing period.
	Duration param.Field[int64] `json:"duration,required"`
	// The unit of billing period duration.
	DurationUnit param.Field[PlanVersionNewParamsReplacePricesPriceNewPlanTieredBpsPriceBillingCycleConfigurationDurationUnit] `json:"duration_unit,required"`
}

func (r PlanVersionNewParamsReplacePricesPriceNewPlanTieredBpsPriceBillingCycleConfiguration) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// The unit of billing period duration.
type PlanVersionNewParamsReplacePricesPriceNewPlanTieredBpsPriceBillingCycleConfigurationDurationUnit string

const (
	PlanVersionNewParamsReplacePricesPriceNewPlanTieredBpsPriceBillingCycleConfigurationDurationUnitDay   PlanVersionNewParamsReplacePricesPriceNewPlanTieredBpsPriceBillingCycleConfigurationDurationUnit = "day"
	PlanVersionNewParamsReplacePricesPriceNewPlanTieredBpsPriceBillingCycleConfigurationDurationUnitMonth PlanVersionNewParamsReplacePricesPriceNewPlanTieredBpsPriceBillingCycleConfigurationDurationUnit = "month"
)

func (r PlanVersionNewParamsReplacePricesPriceNewPlanTieredBpsPriceBillingCycleConfigurationDurationUnit) IsKnown() bool {
	switch r {
	case PlanVersionNewParamsReplacePricesPriceNewPlanTieredBpsPriceBillingCycleConfigurationDurationUnitDay, PlanVersionNewParamsReplacePricesPriceNewPlanTieredBpsPriceBillingCycleConfigurationDurationUnitMonth:
		return true
	}
	return false
}

// For dimensional price: specifies a price group and dimension values
type PlanVersionNewParamsReplacePricesPriceNewPlanTieredBpsPriceDimensionalPriceConfiguration struct {
	// The list of dimension values matching (in order) the dimensions of the price
	// group
	DimensionValues param.Field[[]string] `json:"dimension_values,required"`
	// The id of the dimensional price group to include this price in
	DimensionalPriceGroupID param.Field[string] `json:"dimensional_price_group_id"`
	// The external id of the dimensional price group to include this price in
	ExternalDimensionalPriceGroupID param.Field[string] `json:"external_dimensional_price_group_id"`
}

func (r PlanVersionNewParamsReplacePricesPriceNewPlanTieredBpsPriceDimensionalPriceConfiguration) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Within each billing cycle, specifies the cadence at which invoices are produced.
// If unspecified, a single invoice is produced per billing cycle.
type PlanVersionNewParamsReplacePricesPriceNewPlanTieredBpsPriceInvoicingCycleConfiguration struct {
	// The duration of the billing period.
	Duration param.Field[int64] `json:"duration,required"`
	// The unit of billing period duration.
	DurationUnit param.Field[PlanVersionNewParamsReplacePricesPriceNewPlanTieredBpsPriceInvoicingCycleConfigurationDurationUnit] `json:"duration_unit,required"`
}

func (r PlanVersionNewParamsReplacePricesPriceNewPlanTieredBpsPriceInvoicingCycleConfiguration) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// The unit of billing period duration.
type PlanVersionNewParamsReplacePricesPriceNewPlanTieredBpsPriceInvoicingCycleConfigurationDurationUnit string

const (
	PlanVersionNewParamsReplacePricesPriceNewPlanTieredBpsPriceInvoicingCycleConfigurationDurationUnitDay   PlanVersionNewParamsReplacePricesPriceNewPlanTieredBpsPriceInvoicingCycleConfigurationDurationUnit = "day"
	PlanVersionNewParamsReplacePricesPriceNewPlanTieredBpsPriceInvoicingCycleConfigurationDurationUnitMonth PlanVersionNewParamsReplacePricesPriceNewPlanTieredBpsPriceInvoicingCycleConfigurationDurationUnit = "month"
)

func (r PlanVersionNewParamsReplacePricesPriceNewPlanTieredBpsPriceInvoicingCycleConfigurationDurationUnit) IsKnown() bool {
	switch r {
	case PlanVersionNewParamsReplacePricesPriceNewPlanTieredBpsPriceInvoicingCycleConfigurationDurationUnitDay, PlanVersionNewParamsReplacePricesPriceNewPlanTieredBpsPriceInvoicingCycleConfigurationDurationUnitMonth:
		return true
	}
	return false
}

type PlanVersionNewParamsReplacePricesPriceNewPlanBpsPrice struct {
	BpsConfig param.Field[PlanVersionNewParamsReplacePricesPriceNewPlanBpsPriceBpsConfig] `json:"bps_config,required"`
	// The cadence to bill for this price on.
	Cadence param.Field[PlanVersionNewParamsReplacePricesPriceNewPlanBpsPriceCadence] `json:"cadence,required"`
	// The id of the item the price will be associated with.
	ItemID    param.Field[string]                                                         `json:"item_id,required"`
	ModelType param.Field[PlanVersionNewParamsReplacePricesPriceNewPlanBpsPriceModelType] `json:"model_type,required"`
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
	BillingCycleConfiguration param.Field[PlanVersionNewParamsReplacePricesPriceNewPlanBpsPriceBillingCycleConfiguration] `json:"billing_cycle_configuration"`
	// The per unit conversion rate of the price currency to the invoicing currency.
	ConversionRate param.Field[float64] `json:"conversion_rate"`
	// An ISO 4217 currency string, or custom pricing unit identifier, in which this
	// price is billed.
	Currency param.Field[string] `json:"currency"`
	// For dimensional price: specifies a price group and dimension values
	DimensionalPriceConfiguration param.Field[PlanVersionNewParamsReplacePricesPriceNewPlanBpsPriceDimensionalPriceConfiguration] `json:"dimensional_price_configuration"`
	// An alias for the price.
	ExternalPriceID param.Field[string] `json:"external_price_id"`
	// If the Price represents a fixed cost, this represents the quantity of units
	// applied.
	FixedPriceQuantity param.Field[float64] `json:"fixed_price_quantity"`
	// The property used to group this price on an invoice
	InvoiceGroupingKey param.Field[string] `json:"invoice_grouping_key"`
	// Within each billing cycle, specifies the cadence at which invoices are produced.
	// If unspecified, a single invoice is produced per billing cycle.
	InvoicingCycleConfiguration param.Field[PlanVersionNewParamsReplacePricesPriceNewPlanBpsPriceInvoicingCycleConfiguration] `json:"invoicing_cycle_configuration"`
	// User-specified key/value pairs for the resource. Individual keys can be removed
	// by setting the value to `null`, and the entire metadata mapping can be cleared
	// by setting `metadata` to `null`.
	Metadata param.Field[map[string]string] `json:"metadata"`
}

func (r PlanVersionNewParamsReplacePricesPriceNewPlanBpsPrice) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r PlanVersionNewParamsReplacePricesPriceNewPlanBpsPrice) implementsPlanVersionNewParamsReplacePricesPriceUnion() {
}

type PlanVersionNewParamsReplacePricesPriceNewPlanBpsPriceBpsConfig struct {
	// Basis point take rate per event
	Bps param.Field[float64] `json:"bps,required"`
	// Optional currency amount maximum to cap spend per event
	PerUnitMaximum param.Field[string] `json:"per_unit_maximum"`
}

func (r PlanVersionNewParamsReplacePricesPriceNewPlanBpsPriceBpsConfig) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// The cadence to bill for this price on.
type PlanVersionNewParamsReplacePricesPriceNewPlanBpsPriceCadence string

const (
	PlanVersionNewParamsReplacePricesPriceNewPlanBpsPriceCadenceAnnual     PlanVersionNewParamsReplacePricesPriceNewPlanBpsPriceCadence = "annual"
	PlanVersionNewParamsReplacePricesPriceNewPlanBpsPriceCadenceSemiAnnual PlanVersionNewParamsReplacePricesPriceNewPlanBpsPriceCadence = "semi_annual"
	PlanVersionNewParamsReplacePricesPriceNewPlanBpsPriceCadenceMonthly    PlanVersionNewParamsReplacePricesPriceNewPlanBpsPriceCadence = "monthly"
	PlanVersionNewParamsReplacePricesPriceNewPlanBpsPriceCadenceQuarterly  PlanVersionNewParamsReplacePricesPriceNewPlanBpsPriceCadence = "quarterly"
	PlanVersionNewParamsReplacePricesPriceNewPlanBpsPriceCadenceOneTime    PlanVersionNewParamsReplacePricesPriceNewPlanBpsPriceCadence = "one_time"
	PlanVersionNewParamsReplacePricesPriceNewPlanBpsPriceCadenceCustom     PlanVersionNewParamsReplacePricesPriceNewPlanBpsPriceCadence = "custom"
)

func (r PlanVersionNewParamsReplacePricesPriceNewPlanBpsPriceCadence) IsKnown() bool {
	switch r {
	case PlanVersionNewParamsReplacePricesPriceNewPlanBpsPriceCadenceAnnual, PlanVersionNewParamsReplacePricesPriceNewPlanBpsPriceCadenceSemiAnnual, PlanVersionNewParamsReplacePricesPriceNewPlanBpsPriceCadenceMonthly, PlanVersionNewParamsReplacePricesPriceNewPlanBpsPriceCadenceQuarterly, PlanVersionNewParamsReplacePricesPriceNewPlanBpsPriceCadenceOneTime, PlanVersionNewParamsReplacePricesPriceNewPlanBpsPriceCadenceCustom:
		return true
	}
	return false
}

type PlanVersionNewParamsReplacePricesPriceNewPlanBpsPriceModelType string

const (
	PlanVersionNewParamsReplacePricesPriceNewPlanBpsPriceModelTypeBps PlanVersionNewParamsReplacePricesPriceNewPlanBpsPriceModelType = "bps"
)

func (r PlanVersionNewParamsReplacePricesPriceNewPlanBpsPriceModelType) IsKnown() bool {
	switch r {
	case PlanVersionNewParamsReplacePricesPriceNewPlanBpsPriceModelTypeBps:
		return true
	}
	return false
}

// For custom cadence: specifies the duration of the billing period in days or
// months.
type PlanVersionNewParamsReplacePricesPriceNewPlanBpsPriceBillingCycleConfiguration struct {
	// The duration of the billing period.
	Duration param.Field[int64] `json:"duration,required"`
	// The unit of billing period duration.
	DurationUnit param.Field[PlanVersionNewParamsReplacePricesPriceNewPlanBpsPriceBillingCycleConfigurationDurationUnit] `json:"duration_unit,required"`
}

func (r PlanVersionNewParamsReplacePricesPriceNewPlanBpsPriceBillingCycleConfiguration) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// The unit of billing period duration.
type PlanVersionNewParamsReplacePricesPriceNewPlanBpsPriceBillingCycleConfigurationDurationUnit string

const (
	PlanVersionNewParamsReplacePricesPriceNewPlanBpsPriceBillingCycleConfigurationDurationUnitDay   PlanVersionNewParamsReplacePricesPriceNewPlanBpsPriceBillingCycleConfigurationDurationUnit = "day"
	PlanVersionNewParamsReplacePricesPriceNewPlanBpsPriceBillingCycleConfigurationDurationUnitMonth PlanVersionNewParamsReplacePricesPriceNewPlanBpsPriceBillingCycleConfigurationDurationUnit = "month"
)

func (r PlanVersionNewParamsReplacePricesPriceNewPlanBpsPriceBillingCycleConfigurationDurationUnit) IsKnown() bool {
	switch r {
	case PlanVersionNewParamsReplacePricesPriceNewPlanBpsPriceBillingCycleConfigurationDurationUnitDay, PlanVersionNewParamsReplacePricesPriceNewPlanBpsPriceBillingCycleConfigurationDurationUnitMonth:
		return true
	}
	return false
}

// For dimensional price: specifies a price group and dimension values
type PlanVersionNewParamsReplacePricesPriceNewPlanBpsPriceDimensionalPriceConfiguration struct {
	// The list of dimension values matching (in order) the dimensions of the price
	// group
	DimensionValues param.Field[[]string] `json:"dimension_values,required"`
	// The id of the dimensional price group to include this price in
	DimensionalPriceGroupID param.Field[string] `json:"dimensional_price_group_id"`
	// The external id of the dimensional price group to include this price in
	ExternalDimensionalPriceGroupID param.Field[string] `json:"external_dimensional_price_group_id"`
}

func (r PlanVersionNewParamsReplacePricesPriceNewPlanBpsPriceDimensionalPriceConfiguration) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Within each billing cycle, specifies the cadence at which invoices are produced.
// If unspecified, a single invoice is produced per billing cycle.
type PlanVersionNewParamsReplacePricesPriceNewPlanBpsPriceInvoicingCycleConfiguration struct {
	// The duration of the billing period.
	Duration param.Field[int64] `json:"duration,required"`
	// The unit of billing period duration.
	DurationUnit param.Field[PlanVersionNewParamsReplacePricesPriceNewPlanBpsPriceInvoicingCycleConfigurationDurationUnit] `json:"duration_unit,required"`
}

func (r PlanVersionNewParamsReplacePricesPriceNewPlanBpsPriceInvoicingCycleConfiguration) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// The unit of billing period duration.
type PlanVersionNewParamsReplacePricesPriceNewPlanBpsPriceInvoicingCycleConfigurationDurationUnit string

const (
	PlanVersionNewParamsReplacePricesPriceNewPlanBpsPriceInvoicingCycleConfigurationDurationUnitDay   PlanVersionNewParamsReplacePricesPriceNewPlanBpsPriceInvoicingCycleConfigurationDurationUnit = "day"
	PlanVersionNewParamsReplacePricesPriceNewPlanBpsPriceInvoicingCycleConfigurationDurationUnitMonth PlanVersionNewParamsReplacePricesPriceNewPlanBpsPriceInvoicingCycleConfigurationDurationUnit = "month"
)

func (r PlanVersionNewParamsReplacePricesPriceNewPlanBpsPriceInvoicingCycleConfigurationDurationUnit) IsKnown() bool {
	switch r {
	case PlanVersionNewParamsReplacePricesPriceNewPlanBpsPriceInvoicingCycleConfigurationDurationUnitDay, PlanVersionNewParamsReplacePricesPriceNewPlanBpsPriceInvoicingCycleConfigurationDurationUnitMonth:
		return true
	}
	return false
}

type PlanVersionNewParamsReplacePricesPriceNewPlanBulkBpsPrice struct {
	BulkBpsConfig param.Field[PlanVersionNewParamsReplacePricesPriceNewPlanBulkBpsPriceBulkBpsConfig] `json:"bulk_bps_config,required"`
	// The cadence to bill for this price on.
	Cadence param.Field[PlanVersionNewParamsReplacePricesPriceNewPlanBulkBpsPriceCadence] `json:"cadence,required"`
	// The id of the item the price will be associated with.
	ItemID    param.Field[string]                                                             `json:"item_id,required"`
	ModelType param.Field[PlanVersionNewParamsReplacePricesPriceNewPlanBulkBpsPriceModelType] `json:"model_type,required"`
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
	BillingCycleConfiguration param.Field[PlanVersionNewParamsReplacePricesPriceNewPlanBulkBpsPriceBillingCycleConfiguration] `json:"billing_cycle_configuration"`
	// The per unit conversion rate of the price currency to the invoicing currency.
	ConversionRate param.Field[float64] `json:"conversion_rate"`
	// An ISO 4217 currency string, or custom pricing unit identifier, in which this
	// price is billed.
	Currency param.Field[string] `json:"currency"`
	// For dimensional price: specifies a price group and dimension values
	DimensionalPriceConfiguration param.Field[PlanVersionNewParamsReplacePricesPriceNewPlanBulkBpsPriceDimensionalPriceConfiguration] `json:"dimensional_price_configuration"`
	// An alias for the price.
	ExternalPriceID param.Field[string] `json:"external_price_id"`
	// If the Price represents a fixed cost, this represents the quantity of units
	// applied.
	FixedPriceQuantity param.Field[float64] `json:"fixed_price_quantity"`
	// The property used to group this price on an invoice
	InvoiceGroupingKey param.Field[string] `json:"invoice_grouping_key"`
	// Within each billing cycle, specifies the cadence at which invoices are produced.
	// If unspecified, a single invoice is produced per billing cycle.
	InvoicingCycleConfiguration param.Field[PlanVersionNewParamsReplacePricesPriceNewPlanBulkBpsPriceInvoicingCycleConfiguration] `json:"invoicing_cycle_configuration"`
	// User-specified key/value pairs for the resource. Individual keys can be removed
	// by setting the value to `null`, and the entire metadata mapping can be cleared
	// by setting `metadata` to `null`.
	Metadata param.Field[map[string]string] `json:"metadata"`
}

func (r PlanVersionNewParamsReplacePricesPriceNewPlanBulkBpsPrice) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r PlanVersionNewParamsReplacePricesPriceNewPlanBulkBpsPrice) implementsPlanVersionNewParamsReplacePricesPriceUnion() {
}

type PlanVersionNewParamsReplacePricesPriceNewPlanBulkBpsPriceBulkBpsConfig struct {
	// Tiers for a bulk BPS pricing model where all usage is aggregated to a single
	// tier based on total volume
	Tiers param.Field[[]PlanVersionNewParamsReplacePricesPriceNewPlanBulkBpsPriceBulkBpsConfigTier] `json:"tiers,required"`
}

func (r PlanVersionNewParamsReplacePricesPriceNewPlanBulkBpsPriceBulkBpsConfig) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type PlanVersionNewParamsReplacePricesPriceNewPlanBulkBpsPriceBulkBpsConfigTier struct {
	// Basis points to rate on
	Bps param.Field[float64] `json:"bps,required"`
	// Upper bound for tier
	MaximumAmount param.Field[string] `json:"maximum_amount"`
	// The maximum amount to charge for any one event
	PerUnitMaximum param.Field[string] `json:"per_unit_maximum"`
}

func (r PlanVersionNewParamsReplacePricesPriceNewPlanBulkBpsPriceBulkBpsConfigTier) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// The cadence to bill for this price on.
type PlanVersionNewParamsReplacePricesPriceNewPlanBulkBpsPriceCadence string

const (
	PlanVersionNewParamsReplacePricesPriceNewPlanBulkBpsPriceCadenceAnnual     PlanVersionNewParamsReplacePricesPriceNewPlanBulkBpsPriceCadence = "annual"
	PlanVersionNewParamsReplacePricesPriceNewPlanBulkBpsPriceCadenceSemiAnnual PlanVersionNewParamsReplacePricesPriceNewPlanBulkBpsPriceCadence = "semi_annual"
	PlanVersionNewParamsReplacePricesPriceNewPlanBulkBpsPriceCadenceMonthly    PlanVersionNewParamsReplacePricesPriceNewPlanBulkBpsPriceCadence = "monthly"
	PlanVersionNewParamsReplacePricesPriceNewPlanBulkBpsPriceCadenceQuarterly  PlanVersionNewParamsReplacePricesPriceNewPlanBulkBpsPriceCadence = "quarterly"
	PlanVersionNewParamsReplacePricesPriceNewPlanBulkBpsPriceCadenceOneTime    PlanVersionNewParamsReplacePricesPriceNewPlanBulkBpsPriceCadence = "one_time"
	PlanVersionNewParamsReplacePricesPriceNewPlanBulkBpsPriceCadenceCustom     PlanVersionNewParamsReplacePricesPriceNewPlanBulkBpsPriceCadence = "custom"
)

func (r PlanVersionNewParamsReplacePricesPriceNewPlanBulkBpsPriceCadence) IsKnown() bool {
	switch r {
	case PlanVersionNewParamsReplacePricesPriceNewPlanBulkBpsPriceCadenceAnnual, PlanVersionNewParamsReplacePricesPriceNewPlanBulkBpsPriceCadenceSemiAnnual, PlanVersionNewParamsReplacePricesPriceNewPlanBulkBpsPriceCadenceMonthly, PlanVersionNewParamsReplacePricesPriceNewPlanBulkBpsPriceCadenceQuarterly, PlanVersionNewParamsReplacePricesPriceNewPlanBulkBpsPriceCadenceOneTime, PlanVersionNewParamsReplacePricesPriceNewPlanBulkBpsPriceCadenceCustom:
		return true
	}
	return false
}

type PlanVersionNewParamsReplacePricesPriceNewPlanBulkBpsPriceModelType string

const (
	PlanVersionNewParamsReplacePricesPriceNewPlanBulkBpsPriceModelTypeBulkBps PlanVersionNewParamsReplacePricesPriceNewPlanBulkBpsPriceModelType = "bulk_bps"
)

func (r PlanVersionNewParamsReplacePricesPriceNewPlanBulkBpsPriceModelType) IsKnown() bool {
	switch r {
	case PlanVersionNewParamsReplacePricesPriceNewPlanBulkBpsPriceModelTypeBulkBps:
		return true
	}
	return false
}

// For custom cadence: specifies the duration of the billing period in days or
// months.
type PlanVersionNewParamsReplacePricesPriceNewPlanBulkBpsPriceBillingCycleConfiguration struct {
	// The duration of the billing period.
	Duration param.Field[int64] `json:"duration,required"`
	// The unit of billing period duration.
	DurationUnit param.Field[PlanVersionNewParamsReplacePricesPriceNewPlanBulkBpsPriceBillingCycleConfigurationDurationUnit] `json:"duration_unit,required"`
}

func (r PlanVersionNewParamsReplacePricesPriceNewPlanBulkBpsPriceBillingCycleConfiguration) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// The unit of billing period duration.
type PlanVersionNewParamsReplacePricesPriceNewPlanBulkBpsPriceBillingCycleConfigurationDurationUnit string

const (
	PlanVersionNewParamsReplacePricesPriceNewPlanBulkBpsPriceBillingCycleConfigurationDurationUnitDay   PlanVersionNewParamsReplacePricesPriceNewPlanBulkBpsPriceBillingCycleConfigurationDurationUnit = "day"
	PlanVersionNewParamsReplacePricesPriceNewPlanBulkBpsPriceBillingCycleConfigurationDurationUnitMonth PlanVersionNewParamsReplacePricesPriceNewPlanBulkBpsPriceBillingCycleConfigurationDurationUnit = "month"
)

func (r PlanVersionNewParamsReplacePricesPriceNewPlanBulkBpsPriceBillingCycleConfigurationDurationUnit) IsKnown() bool {
	switch r {
	case PlanVersionNewParamsReplacePricesPriceNewPlanBulkBpsPriceBillingCycleConfigurationDurationUnitDay, PlanVersionNewParamsReplacePricesPriceNewPlanBulkBpsPriceBillingCycleConfigurationDurationUnitMonth:
		return true
	}
	return false
}

// For dimensional price: specifies a price group and dimension values
type PlanVersionNewParamsReplacePricesPriceNewPlanBulkBpsPriceDimensionalPriceConfiguration struct {
	// The list of dimension values matching (in order) the dimensions of the price
	// group
	DimensionValues param.Field[[]string] `json:"dimension_values,required"`
	// The id of the dimensional price group to include this price in
	DimensionalPriceGroupID param.Field[string] `json:"dimensional_price_group_id"`
	// The external id of the dimensional price group to include this price in
	ExternalDimensionalPriceGroupID param.Field[string] `json:"external_dimensional_price_group_id"`
}

func (r PlanVersionNewParamsReplacePricesPriceNewPlanBulkBpsPriceDimensionalPriceConfiguration) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Within each billing cycle, specifies the cadence at which invoices are produced.
// If unspecified, a single invoice is produced per billing cycle.
type PlanVersionNewParamsReplacePricesPriceNewPlanBulkBpsPriceInvoicingCycleConfiguration struct {
	// The duration of the billing period.
	Duration param.Field[int64] `json:"duration,required"`
	// The unit of billing period duration.
	DurationUnit param.Field[PlanVersionNewParamsReplacePricesPriceNewPlanBulkBpsPriceInvoicingCycleConfigurationDurationUnit] `json:"duration_unit,required"`
}

func (r PlanVersionNewParamsReplacePricesPriceNewPlanBulkBpsPriceInvoicingCycleConfiguration) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// The unit of billing period duration.
type PlanVersionNewParamsReplacePricesPriceNewPlanBulkBpsPriceInvoicingCycleConfigurationDurationUnit string

const (
	PlanVersionNewParamsReplacePricesPriceNewPlanBulkBpsPriceInvoicingCycleConfigurationDurationUnitDay   PlanVersionNewParamsReplacePricesPriceNewPlanBulkBpsPriceInvoicingCycleConfigurationDurationUnit = "day"
	PlanVersionNewParamsReplacePricesPriceNewPlanBulkBpsPriceInvoicingCycleConfigurationDurationUnitMonth PlanVersionNewParamsReplacePricesPriceNewPlanBulkBpsPriceInvoicingCycleConfigurationDurationUnit = "month"
)

func (r PlanVersionNewParamsReplacePricesPriceNewPlanBulkBpsPriceInvoicingCycleConfigurationDurationUnit) IsKnown() bool {
	switch r {
	case PlanVersionNewParamsReplacePricesPriceNewPlanBulkBpsPriceInvoicingCycleConfigurationDurationUnitDay, PlanVersionNewParamsReplacePricesPriceNewPlanBulkBpsPriceInvoicingCycleConfigurationDurationUnitMonth:
		return true
	}
	return false
}

type PlanVersionNewParamsReplacePricesPriceNewPlanBulkPrice struct {
	BulkConfig param.Field[PlanVersionNewParamsReplacePricesPriceNewPlanBulkPriceBulkConfig] `json:"bulk_config,required"`
	// The cadence to bill for this price on.
	Cadence param.Field[PlanVersionNewParamsReplacePricesPriceNewPlanBulkPriceCadence] `json:"cadence,required"`
	// The id of the item the price will be associated with.
	ItemID    param.Field[string]                                                          `json:"item_id,required"`
	ModelType param.Field[PlanVersionNewParamsReplacePricesPriceNewPlanBulkPriceModelType] `json:"model_type,required"`
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
	BillingCycleConfiguration param.Field[PlanVersionNewParamsReplacePricesPriceNewPlanBulkPriceBillingCycleConfiguration] `json:"billing_cycle_configuration"`
	// The per unit conversion rate of the price currency to the invoicing currency.
	ConversionRate param.Field[float64] `json:"conversion_rate"`
	// An ISO 4217 currency string, or custom pricing unit identifier, in which this
	// price is billed.
	Currency param.Field[string] `json:"currency"`
	// For dimensional price: specifies a price group and dimension values
	DimensionalPriceConfiguration param.Field[PlanVersionNewParamsReplacePricesPriceNewPlanBulkPriceDimensionalPriceConfiguration] `json:"dimensional_price_configuration"`
	// An alias for the price.
	ExternalPriceID param.Field[string] `json:"external_price_id"`
	// If the Price represents a fixed cost, this represents the quantity of units
	// applied.
	FixedPriceQuantity param.Field[float64] `json:"fixed_price_quantity"`
	// The property used to group this price on an invoice
	InvoiceGroupingKey param.Field[string] `json:"invoice_grouping_key"`
	// Within each billing cycle, specifies the cadence at which invoices are produced.
	// If unspecified, a single invoice is produced per billing cycle.
	InvoicingCycleConfiguration param.Field[PlanVersionNewParamsReplacePricesPriceNewPlanBulkPriceInvoicingCycleConfiguration] `json:"invoicing_cycle_configuration"`
	// User-specified key/value pairs for the resource. Individual keys can be removed
	// by setting the value to `null`, and the entire metadata mapping can be cleared
	// by setting `metadata` to `null`.
	Metadata param.Field[map[string]string] `json:"metadata"`
}

func (r PlanVersionNewParamsReplacePricesPriceNewPlanBulkPrice) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r PlanVersionNewParamsReplacePricesPriceNewPlanBulkPrice) implementsPlanVersionNewParamsReplacePricesPriceUnion() {
}

type PlanVersionNewParamsReplacePricesPriceNewPlanBulkPriceBulkConfig struct {
	// Bulk tiers for rating based on total usage volume
	Tiers param.Field[[]PlanVersionNewParamsReplacePricesPriceNewPlanBulkPriceBulkConfigTier] `json:"tiers,required"`
}

func (r PlanVersionNewParamsReplacePricesPriceNewPlanBulkPriceBulkConfig) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type PlanVersionNewParamsReplacePricesPriceNewPlanBulkPriceBulkConfigTier struct {
	// Amount per unit
	UnitAmount param.Field[string] `json:"unit_amount,required"`
	// Upper bound for this tier
	MaximumUnits param.Field[float64] `json:"maximum_units"`
}

func (r PlanVersionNewParamsReplacePricesPriceNewPlanBulkPriceBulkConfigTier) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// The cadence to bill for this price on.
type PlanVersionNewParamsReplacePricesPriceNewPlanBulkPriceCadence string

const (
	PlanVersionNewParamsReplacePricesPriceNewPlanBulkPriceCadenceAnnual     PlanVersionNewParamsReplacePricesPriceNewPlanBulkPriceCadence = "annual"
	PlanVersionNewParamsReplacePricesPriceNewPlanBulkPriceCadenceSemiAnnual PlanVersionNewParamsReplacePricesPriceNewPlanBulkPriceCadence = "semi_annual"
	PlanVersionNewParamsReplacePricesPriceNewPlanBulkPriceCadenceMonthly    PlanVersionNewParamsReplacePricesPriceNewPlanBulkPriceCadence = "monthly"
	PlanVersionNewParamsReplacePricesPriceNewPlanBulkPriceCadenceQuarterly  PlanVersionNewParamsReplacePricesPriceNewPlanBulkPriceCadence = "quarterly"
	PlanVersionNewParamsReplacePricesPriceNewPlanBulkPriceCadenceOneTime    PlanVersionNewParamsReplacePricesPriceNewPlanBulkPriceCadence = "one_time"
	PlanVersionNewParamsReplacePricesPriceNewPlanBulkPriceCadenceCustom     PlanVersionNewParamsReplacePricesPriceNewPlanBulkPriceCadence = "custom"
)

func (r PlanVersionNewParamsReplacePricesPriceNewPlanBulkPriceCadence) IsKnown() bool {
	switch r {
	case PlanVersionNewParamsReplacePricesPriceNewPlanBulkPriceCadenceAnnual, PlanVersionNewParamsReplacePricesPriceNewPlanBulkPriceCadenceSemiAnnual, PlanVersionNewParamsReplacePricesPriceNewPlanBulkPriceCadenceMonthly, PlanVersionNewParamsReplacePricesPriceNewPlanBulkPriceCadenceQuarterly, PlanVersionNewParamsReplacePricesPriceNewPlanBulkPriceCadenceOneTime, PlanVersionNewParamsReplacePricesPriceNewPlanBulkPriceCadenceCustom:
		return true
	}
	return false
}

type PlanVersionNewParamsReplacePricesPriceNewPlanBulkPriceModelType string

const (
	PlanVersionNewParamsReplacePricesPriceNewPlanBulkPriceModelTypeBulk PlanVersionNewParamsReplacePricesPriceNewPlanBulkPriceModelType = "bulk"
)

func (r PlanVersionNewParamsReplacePricesPriceNewPlanBulkPriceModelType) IsKnown() bool {
	switch r {
	case PlanVersionNewParamsReplacePricesPriceNewPlanBulkPriceModelTypeBulk:
		return true
	}
	return false
}

// For custom cadence: specifies the duration of the billing period in days or
// months.
type PlanVersionNewParamsReplacePricesPriceNewPlanBulkPriceBillingCycleConfiguration struct {
	// The duration of the billing period.
	Duration param.Field[int64] `json:"duration,required"`
	// The unit of billing period duration.
	DurationUnit param.Field[PlanVersionNewParamsReplacePricesPriceNewPlanBulkPriceBillingCycleConfigurationDurationUnit] `json:"duration_unit,required"`
}

func (r PlanVersionNewParamsReplacePricesPriceNewPlanBulkPriceBillingCycleConfiguration) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// The unit of billing period duration.
type PlanVersionNewParamsReplacePricesPriceNewPlanBulkPriceBillingCycleConfigurationDurationUnit string

const (
	PlanVersionNewParamsReplacePricesPriceNewPlanBulkPriceBillingCycleConfigurationDurationUnitDay   PlanVersionNewParamsReplacePricesPriceNewPlanBulkPriceBillingCycleConfigurationDurationUnit = "day"
	PlanVersionNewParamsReplacePricesPriceNewPlanBulkPriceBillingCycleConfigurationDurationUnitMonth PlanVersionNewParamsReplacePricesPriceNewPlanBulkPriceBillingCycleConfigurationDurationUnit = "month"
)

func (r PlanVersionNewParamsReplacePricesPriceNewPlanBulkPriceBillingCycleConfigurationDurationUnit) IsKnown() bool {
	switch r {
	case PlanVersionNewParamsReplacePricesPriceNewPlanBulkPriceBillingCycleConfigurationDurationUnitDay, PlanVersionNewParamsReplacePricesPriceNewPlanBulkPriceBillingCycleConfigurationDurationUnitMonth:
		return true
	}
	return false
}

// For dimensional price: specifies a price group and dimension values
type PlanVersionNewParamsReplacePricesPriceNewPlanBulkPriceDimensionalPriceConfiguration struct {
	// The list of dimension values matching (in order) the dimensions of the price
	// group
	DimensionValues param.Field[[]string] `json:"dimension_values,required"`
	// The id of the dimensional price group to include this price in
	DimensionalPriceGroupID param.Field[string] `json:"dimensional_price_group_id"`
	// The external id of the dimensional price group to include this price in
	ExternalDimensionalPriceGroupID param.Field[string] `json:"external_dimensional_price_group_id"`
}

func (r PlanVersionNewParamsReplacePricesPriceNewPlanBulkPriceDimensionalPriceConfiguration) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Within each billing cycle, specifies the cadence at which invoices are produced.
// If unspecified, a single invoice is produced per billing cycle.
type PlanVersionNewParamsReplacePricesPriceNewPlanBulkPriceInvoicingCycleConfiguration struct {
	// The duration of the billing period.
	Duration param.Field[int64] `json:"duration,required"`
	// The unit of billing period duration.
	DurationUnit param.Field[PlanVersionNewParamsReplacePricesPriceNewPlanBulkPriceInvoicingCycleConfigurationDurationUnit] `json:"duration_unit,required"`
}

func (r PlanVersionNewParamsReplacePricesPriceNewPlanBulkPriceInvoicingCycleConfiguration) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// The unit of billing period duration.
type PlanVersionNewParamsReplacePricesPriceNewPlanBulkPriceInvoicingCycleConfigurationDurationUnit string

const (
	PlanVersionNewParamsReplacePricesPriceNewPlanBulkPriceInvoicingCycleConfigurationDurationUnitDay   PlanVersionNewParamsReplacePricesPriceNewPlanBulkPriceInvoicingCycleConfigurationDurationUnit = "day"
	PlanVersionNewParamsReplacePricesPriceNewPlanBulkPriceInvoicingCycleConfigurationDurationUnitMonth PlanVersionNewParamsReplacePricesPriceNewPlanBulkPriceInvoicingCycleConfigurationDurationUnit = "month"
)

func (r PlanVersionNewParamsReplacePricesPriceNewPlanBulkPriceInvoicingCycleConfigurationDurationUnit) IsKnown() bool {
	switch r {
	case PlanVersionNewParamsReplacePricesPriceNewPlanBulkPriceInvoicingCycleConfigurationDurationUnitDay, PlanVersionNewParamsReplacePricesPriceNewPlanBulkPriceInvoicingCycleConfigurationDurationUnitMonth:
		return true
	}
	return false
}

type PlanVersionNewParamsReplacePricesPriceNewPlanThresholdTotalAmountPrice struct {
	// The cadence to bill for this price on.
	Cadence param.Field[PlanVersionNewParamsReplacePricesPriceNewPlanThresholdTotalAmountPriceCadence] `json:"cadence,required"`
	// The id of the item the price will be associated with.
	ItemID    param.Field[string]                                                                          `json:"item_id,required"`
	ModelType param.Field[PlanVersionNewParamsReplacePricesPriceNewPlanThresholdTotalAmountPriceModelType] `json:"model_type,required"`
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
	BillingCycleConfiguration param.Field[PlanVersionNewParamsReplacePricesPriceNewPlanThresholdTotalAmountPriceBillingCycleConfiguration] `json:"billing_cycle_configuration"`
	// The per unit conversion rate of the price currency to the invoicing currency.
	ConversionRate param.Field[float64] `json:"conversion_rate"`
	// An ISO 4217 currency string, or custom pricing unit identifier, in which this
	// price is billed.
	Currency param.Field[string] `json:"currency"`
	// For dimensional price: specifies a price group and dimension values
	DimensionalPriceConfiguration param.Field[PlanVersionNewParamsReplacePricesPriceNewPlanThresholdTotalAmountPriceDimensionalPriceConfiguration] `json:"dimensional_price_configuration"`
	// An alias for the price.
	ExternalPriceID param.Field[string] `json:"external_price_id"`
	// If the Price represents a fixed cost, this represents the quantity of units
	// applied.
	FixedPriceQuantity param.Field[float64] `json:"fixed_price_quantity"`
	// The property used to group this price on an invoice
	InvoiceGroupingKey param.Field[string] `json:"invoice_grouping_key"`
	// Within each billing cycle, specifies the cadence at which invoices are produced.
	// If unspecified, a single invoice is produced per billing cycle.
	InvoicingCycleConfiguration param.Field[PlanVersionNewParamsReplacePricesPriceNewPlanThresholdTotalAmountPriceInvoicingCycleConfiguration] `json:"invoicing_cycle_configuration"`
	// User-specified key/value pairs for the resource. Individual keys can be removed
	// by setting the value to `null`, and the entire metadata mapping can be cleared
	// by setting `metadata` to `null`.
	Metadata param.Field[map[string]string] `json:"metadata"`
}

func (r PlanVersionNewParamsReplacePricesPriceNewPlanThresholdTotalAmountPrice) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r PlanVersionNewParamsReplacePricesPriceNewPlanThresholdTotalAmountPrice) implementsPlanVersionNewParamsReplacePricesPriceUnion() {
}

// The cadence to bill for this price on.
type PlanVersionNewParamsReplacePricesPriceNewPlanThresholdTotalAmountPriceCadence string

const (
	PlanVersionNewParamsReplacePricesPriceNewPlanThresholdTotalAmountPriceCadenceAnnual     PlanVersionNewParamsReplacePricesPriceNewPlanThresholdTotalAmountPriceCadence = "annual"
	PlanVersionNewParamsReplacePricesPriceNewPlanThresholdTotalAmountPriceCadenceSemiAnnual PlanVersionNewParamsReplacePricesPriceNewPlanThresholdTotalAmountPriceCadence = "semi_annual"
	PlanVersionNewParamsReplacePricesPriceNewPlanThresholdTotalAmountPriceCadenceMonthly    PlanVersionNewParamsReplacePricesPriceNewPlanThresholdTotalAmountPriceCadence = "monthly"
	PlanVersionNewParamsReplacePricesPriceNewPlanThresholdTotalAmountPriceCadenceQuarterly  PlanVersionNewParamsReplacePricesPriceNewPlanThresholdTotalAmountPriceCadence = "quarterly"
	PlanVersionNewParamsReplacePricesPriceNewPlanThresholdTotalAmountPriceCadenceOneTime    PlanVersionNewParamsReplacePricesPriceNewPlanThresholdTotalAmountPriceCadence = "one_time"
	PlanVersionNewParamsReplacePricesPriceNewPlanThresholdTotalAmountPriceCadenceCustom     PlanVersionNewParamsReplacePricesPriceNewPlanThresholdTotalAmountPriceCadence = "custom"
)

func (r PlanVersionNewParamsReplacePricesPriceNewPlanThresholdTotalAmountPriceCadence) IsKnown() bool {
	switch r {
	case PlanVersionNewParamsReplacePricesPriceNewPlanThresholdTotalAmountPriceCadenceAnnual, PlanVersionNewParamsReplacePricesPriceNewPlanThresholdTotalAmountPriceCadenceSemiAnnual, PlanVersionNewParamsReplacePricesPriceNewPlanThresholdTotalAmountPriceCadenceMonthly, PlanVersionNewParamsReplacePricesPriceNewPlanThresholdTotalAmountPriceCadenceQuarterly, PlanVersionNewParamsReplacePricesPriceNewPlanThresholdTotalAmountPriceCadenceOneTime, PlanVersionNewParamsReplacePricesPriceNewPlanThresholdTotalAmountPriceCadenceCustom:
		return true
	}
	return false
}

type PlanVersionNewParamsReplacePricesPriceNewPlanThresholdTotalAmountPriceModelType string

const (
	PlanVersionNewParamsReplacePricesPriceNewPlanThresholdTotalAmountPriceModelTypeThresholdTotalAmount PlanVersionNewParamsReplacePricesPriceNewPlanThresholdTotalAmountPriceModelType = "threshold_total_amount"
)

func (r PlanVersionNewParamsReplacePricesPriceNewPlanThresholdTotalAmountPriceModelType) IsKnown() bool {
	switch r {
	case PlanVersionNewParamsReplacePricesPriceNewPlanThresholdTotalAmountPriceModelTypeThresholdTotalAmount:
		return true
	}
	return false
}

// For custom cadence: specifies the duration of the billing period in days or
// months.
type PlanVersionNewParamsReplacePricesPriceNewPlanThresholdTotalAmountPriceBillingCycleConfiguration struct {
	// The duration of the billing period.
	Duration param.Field[int64] `json:"duration,required"`
	// The unit of billing period duration.
	DurationUnit param.Field[PlanVersionNewParamsReplacePricesPriceNewPlanThresholdTotalAmountPriceBillingCycleConfigurationDurationUnit] `json:"duration_unit,required"`
}

func (r PlanVersionNewParamsReplacePricesPriceNewPlanThresholdTotalAmountPriceBillingCycleConfiguration) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// The unit of billing period duration.
type PlanVersionNewParamsReplacePricesPriceNewPlanThresholdTotalAmountPriceBillingCycleConfigurationDurationUnit string

const (
	PlanVersionNewParamsReplacePricesPriceNewPlanThresholdTotalAmountPriceBillingCycleConfigurationDurationUnitDay   PlanVersionNewParamsReplacePricesPriceNewPlanThresholdTotalAmountPriceBillingCycleConfigurationDurationUnit = "day"
	PlanVersionNewParamsReplacePricesPriceNewPlanThresholdTotalAmountPriceBillingCycleConfigurationDurationUnitMonth PlanVersionNewParamsReplacePricesPriceNewPlanThresholdTotalAmountPriceBillingCycleConfigurationDurationUnit = "month"
)

func (r PlanVersionNewParamsReplacePricesPriceNewPlanThresholdTotalAmountPriceBillingCycleConfigurationDurationUnit) IsKnown() bool {
	switch r {
	case PlanVersionNewParamsReplacePricesPriceNewPlanThresholdTotalAmountPriceBillingCycleConfigurationDurationUnitDay, PlanVersionNewParamsReplacePricesPriceNewPlanThresholdTotalAmountPriceBillingCycleConfigurationDurationUnitMonth:
		return true
	}
	return false
}

// For dimensional price: specifies a price group and dimension values
type PlanVersionNewParamsReplacePricesPriceNewPlanThresholdTotalAmountPriceDimensionalPriceConfiguration struct {
	// The list of dimension values matching (in order) the dimensions of the price
	// group
	DimensionValues param.Field[[]string] `json:"dimension_values,required"`
	// The id of the dimensional price group to include this price in
	DimensionalPriceGroupID param.Field[string] `json:"dimensional_price_group_id"`
	// The external id of the dimensional price group to include this price in
	ExternalDimensionalPriceGroupID param.Field[string] `json:"external_dimensional_price_group_id"`
}

func (r PlanVersionNewParamsReplacePricesPriceNewPlanThresholdTotalAmountPriceDimensionalPriceConfiguration) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Within each billing cycle, specifies the cadence at which invoices are produced.
// If unspecified, a single invoice is produced per billing cycle.
type PlanVersionNewParamsReplacePricesPriceNewPlanThresholdTotalAmountPriceInvoicingCycleConfiguration struct {
	// The duration of the billing period.
	Duration param.Field[int64] `json:"duration,required"`
	// The unit of billing period duration.
	DurationUnit param.Field[PlanVersionNewParamsReplacePricesPriceNewPlanThresholdTotalAmountPriceInvoicingCycleConfigurationDurationUnit] `json:"duration_unit,required"`
}

func (r PlanVersionNewParamsReplacePricesPriceNewPlanThresholdTotalAmountPriceInvoicingCycleConfiguration) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// The unit of billing period duration.
type PlanVersionNewParamsReplacePricesPriceNewPlanThresholdTotalAmountPriceInvoicingCycleConfigurationDurationUnit string

const (
	PlanVersionNewParamsReplacePricesPriceNewPlanThresholdTotalAmountPriceInvoicingCycleConfigurationDurationUnitDay   PlanVersionNewParamsReplacePricesPriceNewPlanThresholdTotalAmountPriceInvoicingCycleConfigurationDurationUnit = "day"
	PlanVersionNewParamsReplacePricesPriceNewPlanThresholdTotalAmountPriceInvoicingCycleConfigurationDurationUnitMonth PlanVersionNewParamsReplacePricesPriceNewPlanThresholdTotalAmountPriceInvoicingCycleConfigurationDurationUnit = "month"
)

func (r PlanVersionNewParamsReplacePricesPriceNewPlanThresholdTotalAmountPriceInvoicingCycleConfigurationDurationUnit) IsKnown() bool {
	switch r {
	case PlanVersionNewParamsReplacePricesPriceNewPlanThresholdTotalAmountPriceInvoicingCycleConfigurationDurationUnitDay, PlanVersionNewParamsReplacePricesPriceNewPlanThresholdTotalAmountPriceInvoicingCycleConfigurationDurationUnitMonth:
		return true
	}
	return false
}

type PlanVersionNewParamsReplacePricesPriceNewPlanTieredPackagePrice struct {
	// The cadence to bill for this price on.
	Cadence param.Field[PlanVersionNewParamsReplacePricesPriceNewPlanTieredPackagePriceCadence] `json:"cadence,required"`
	// The id of the item the price will be associated with.
	ItemID    param.Field[string]                                                                   `json:"item_id,required"`
	ModelType param.Field[PlanVersionNewParamsReplacePricesPriceNewPlanTieredPackagePriceModelType] `json:"model_type,required"`
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
	BillingCycleConfiguration param.Field[PlanVersionNewParamsReplacePricesPriceNewPlanTieredPackagePriceBillingCycleConfiguration] `json:"billing_cycle_configuration"`
	// The per unit conversion rate of the price currency to the invoicing currency.
	ConversionRate param.Field[float64] `json:"conversion_rate"`
	// An ISO 4217 currency string, or custom pricing unit identifier, in which this
	// price is billed.
	Currency param.Field[string] `json:"currency"`
	// For dimensional price: specifies a price group and dimension values
	DimensionalPriceConfiguration param.Field[PlanVersionNewParamsReplacePricesPriceNewPlanTieredPackagePriceDimensionalPriceConfiguration] `json:"dimensional_price_configuration"`
	// An alias for the price.
	ExternalPriceID param.Field[string] `json:"external_price_id"`
	// If the Price represents a fixed cost, this represents the quantity of units
	// applied.
	FixedPriceQuantity param.Field[float64] `json:"fixed_price_quantity"`
	// The property used to group this price on an invoice
	InvoiceGroupingKey param.Field[string] `json:"invoice_grouping_key"`
	// Within each billing cycle, specifies the cadence at which invoices are produced.
	// If unspecified, a single invoice is produced per billing cycle.
	InvoicingCycleConfiguration param.Field[PlanVersionNewParamsReplacePricesPriceNewPlanTieredPackagePriceInvoicingCycleConfiguration] `json:"invoicing_cycle_configuration"`
	// User-specified key/value pairs for the resource. Individual keys can be removed
	// by setting the value to `null`, and the entire metadata mapping can be cleared
	// by setting `metadata` to `null`.
	Metadata param.Field[map[string]string] `json:"metadata"`
}

func (r PlanVersionNewParamsReplacePricesPriceNewPlanTieredPackagePrice) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r PlanVersionNewParamsReplacePricesPriceNewPlanTieredPackagePrice) implementsPlanVersionNewParamsReplacePricesPriceUnion() {
}

// The cadence to bill for this price on.
type PlanVersionNewParamsReplacePricesPriceNewPlanTieredPackagePriceCadence string

const (
	PlanVersionNewParamsReplacePricesPriceNewPlanTieredPackagePriceCadenceAnnual     PlanVersionNewParamsReplacePricesPriceNewPlanTieredPackagePriceCadence = "annual"
	PlanVersionNewParamsReplacePricesPriceNewPlanTieredPackagePriceCadenceSemiAnnual PlanVersionNewParamsReplacePricesPriceNewPlanTieredPackagePriceCadence = "semi_annual"
	PlanVersionNewParamsReplacePricesPriceNewPlanTieredPackagePriceCadenceMonthly    PlanVersionNewParamsReplacePricesPriceNewPlanTieredPackagePriceCadence = "monthly"
	PlanVersionNewParamsReplacePricesPriceNewPlanTieredPackagePriceCadenceQuarterly  PlanVersionNewParamsReplacePricesPriceNewPlanTieredPackagePriceCadence = "quarterly"
	PlanVersionNewParamsReplacePricesPriceNewPlanTieredPackagePriceCadenceOneTime    PlanVersionNewParamsReplacePricesPriceNewPlanTieredPackagePriceCadence = "one_time"
	PlanVersionNewParamsReplacePricesPriceNewPlanTieredPackagePriceCadenceCustom     PlanVersionNewParamsReplacePricesPriceNewPlanTieredPackagePriceCadence = "custom"
)

func (r PlanVersionNewParamsReplacePricesPriceNewPlanTieredPackagePriceCadence) IsKnown() bool {
	switch r {
	case PlanVersionNewParamsReplacePricesPriceNewPlanTieredPackagePriceCadenceAnnual, PlanVersionNewParamsReplacePricesPriceNewPlanTieredPackagePriceCadenceSemiAnnual, PlanVersionNewParamsReplacePricesPriceNewPlanTieredPackagePriceCadenceMonthly, PlanVersionNewParamsReplacePricesPriceNewPlanTieredPackagePriceCadenceQuarterly, PlanVersionNewParamsReplacePricesPriceNewPlanTieredPackagePriceCadenceOneTime, PlanVersionNewParamsReplacePricesPriceNewPlanTieredPackagePriceCadenceCustom:
		return true
	}
	return false
}

type PlanVersionNewParamsReplacePricesPriceNewPlanTieredPackagePriceModelType string

const (
	PlanVersionNewParamsReplacePricesPriceNewPlanTieredPackagePriceModelTypeTieredPackage PlanVersionNewParamsReplacePricesPriceNewPlanTieredPackagePriceModelType = "tiered_package"
)

func (r PlanVersionNewParamsReplacePricesPriceNewPlanTieredPackagePriceModelType) IsKnown() bool {
	switch r {
	case PlanVersionNewParamsReplacePricesPriceNewPlanTieredPackagePriceModelTypeTieredPackage:
		return true
	}
	return false
}

// For custom cadence: specifies the duration of the billing period in days or
// months.
type PlanVersionNewParamsReplacePricesPriceNewPlanTieredPackagePriceBillingCycleConfiguration struct {
	// The duration of the billing period.
	Duration param.Field[int64] `json:"duration,required"`
	// The unit of billing period duration.
	DurationUnit param.Field[PlanVersionNewParamsReplacePricesPriceNewPlanTieredPackagePriceBillingCycleConfigurationDurationUnit] `json:"duration_unit,required"`
}

func (r PlanVersionNewParamsReplacePricesPriceNewPlanTieredPackagePriceBillingCycleConfiguration) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// The unit of billing period duration.
type PlanVersionNewParamsReplacePricesPriceNewPlanTieredPackagePriceBillingCycleConfigurationDurationUnit string

const (
	PlanVersionNewParamsReplacePricesPriceNewPlanTieredPackagePriceBillingCycleConfigurationDurationUnitDay   PlanVersionNewParamsReplacePricesPriceNewPlanTieredPackagePriceBillingCycleConfigurationDurationUnit = "day"
	PlanVersionNewParamsReplacePricesPriceNewPlanTieredPackagePriceBillingCycleConfigurationDurationUnitMonth PlanVersionNewParamsReplacePricesPriceNewPlanTieredPackagePriceBillingCycleConfigurationDurationUnit = "month"
)

func (r PlanVersionNewParamsReplacePricesPriceNewPlanTieredPackagePriceBillingCycleConfigurationDurationUnit) IsKnown() bool {
	switch r {
	case PlanVersionNewParamsReplacePricesPriceNewPlanTieredPackagePriceBillingCycleConfigurationDurationUnitDay, PlanVersionNewParamsReplacePricesPriceNewPlanTieredPackagePriceBillingCycleConfigurationDurationUnitMonth:
		return true
	}
	return false
}

// For dimensional price: specifies a price group and dimension values
type PlanVersionNewParamsReplacePricesPriceNewPlanTieredPackagePriceDimensionalPriceConfiguration struct {
	// The list of dimension values matching (in order) the dimensions of the price
	// group
	DimensionValues param.Field[[]string] `json:"dimension_values,required"`
	// The id of the dimensional price group to include this price in
	DimensionalPriceGroupID param.Field[string] `json:"dimensional_price_group_id"`
	// The external id of the dimensional price group to include this price in
	ExternalDimensionalPriceGroupID param.Field[string] `json:"external_dimensional_price_group_id"`
}

func (r PlanVersionNewParamsReplacePricesPriceNewPlanTieredPackagePriceDimensionalPriceConfiguration) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Within each billing cycle, specifies the cadence at which invoices are produced.
// If unspecified, a single invoice is produced per billing cycle.
type PlanVersionNewParamsReplacePricesPriceNewPlanTieredPackagePriceInvoicingCycleConfiguration struct {
	// The duration of the billing period.
	Duration param.Field[int64] `json:"duration,required"`
	// The unit of billing period duration.
	DurationUnit param.Field[PlanVersionNewParamsReplacePricesPriceNewPlanTieredPackagePriceInvoicingCycleConfigurationDurationUnit] `json:"duration_unit,required"`
}

func (r PlanVersionNewParamsReplacePricesPriceNewPlanTieredPackagePriceInvoicingCycleConfiguration) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// The unit of billing period duration.
type PlanVersionNewParamsReplacePricesPriceNewPlanTieredPackagePriceInvoicingCycleConfigurationDurationUnit string

const (
	PlanVersionNewParamsReplacePricesPriceNewPlanTieredPackagePriceInvoicingCycleConfigurationDurationUnitDay   PlanVersionNewParamsReplacePricesPriceNewPlanTieredPackagePriceInvoicingCycleConfigurationDurationUnit = "day"
	PlanVersionNewParamsReplacePricesPriceNewPlanTieredPackagePriceInvoicingCycleConfigurationDurationUnitMonth PlanVersionNewParamsReplacePricesPriceNewPlanTieredPackagePriceInvoicingCycleConfigurationDurationUnit = "month"
)

func (r PlanVersionNewParamsReplacePricesPriceNewPlanTieredPackagePriceInvoicingCycleConfigurationDurationUnit) IsKnown() bool {
	switch r {
	case PlanVersionNewParamsReplacePricesPriceNewPlanTieredPackagePriceInvoicingCycleConfigurationDurationUnitDay, PlanVersionNewParamsReplacePricesPriceNewPlanTieredPackagePriceInvoicingCycleConfigurationDurationUnitMonth:
		return true
	}
	return false
}

type PlanVersionNewParamsReplacePricesPriceNewPlanTieredWithMinimumPrice struct {
	// The cadence to bill for this price on.
	Cadence param.Field[PlanVersionNewParamsReplacePricesPriceNewPlanTieredWithMinimumPriceCadence] `json:"cadence,required"`
	// The id of the item the price will be associated with.
	ItemID    param.Field[string]                                                                       `json:"item_id,required"`
	ModelType param.Field[PlanVersionNewParamsReplacePricesPriceNewPlanTieredWithMinimumPriceModelType] `json:"model_type,required"`
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
	BillingCycleConfiguration param.Field[PlanVersionNewParamsReplacePricesPriceNewPlanTieredWithMinimumPriceBillingCycleConfiguration] `json:"billing_cycle_configuration"`
	// The per unit conversion rate of the price currency to the invoicing currency.
	ConversionRate param.Field[float64] `json:"conversion_rate"`
	// An ISO 4217 currency string, or custom pricing unit identifier, in which this
	// price is billed.
	Currency param.Field[string] `json:"currency"`
	// For dimensional price: specifies a price group and dimension values
	DimensionalPriceConfiguration param.Field[PlanVersionNewParamsReplacePricesPriceNewPlanTieredWithMinimumPriceDimensionalPriceConfiguration] `json:"dimensional_price_configuration"`
	// An alias for the price.
	ExternalPriceID param.Field[string] `json:"external_price_id"`
	// If the Price represents a fixed cost, this represents the quantity of units
	// applied.
	FixedPriceQuantity param.Field[float64] `json:"fixed_price_quantity"`
	// The property used to group this price on an invoice
	InvoiceGroupingKey param.Field[string] `json:"invoice_grouping_key"`
	// Within each billing cycle, specifies the cadence at which invoices are produced.
	// If unspecified, a single invoice is produced per billing cycle.
	InvoicingCycleConfiguration param.Field[PlanVersionNewParamsReplacePricesPriceNewPlanTieredWithMinimumPriceInvoicingCycleConfiguration] `json:"invoicing_cycle_configuration"`
	// User-specified key/value pairs for the resource. Individual keys can be removed
	// by setting the value to `null`, and the entire metadata mapping can be cleared
	// by setting `metadata` to `null`.
	Metadata param.Field[map[string]string] `json:"metadata"`
}

func (r PlanVersionNewParamsReplacePricesPriceNewPlanTieredWithMinimumPrice) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r PlanVersionNewParamsReplacePricesPriceNewPlanTieredWithMinimumPrice) implementsPlanVersionNewParamsReplacePricesPriceUnion() {
}

// The cadence to bill for this price on.
type PlanVersionNewParamsReplacePricesPriceNewPlanTieredWithMinimumPriceCadence string

const (
	PlanVersionNewParamsReplacePricesPriceNewPlanTieredWithMinimumPriceCadenceAnnual     PlanVersionNewParamsReplacePricesPriceNewPlanTieredWithMinimumPriceCadence = "annual"
	PlanVersionNewParamsReplacePricesPriceNewPlanTieredWithMinimumPriceCadenceSemiAnnual PlanVersionNewParamsReplacePricesPriceNewPlanTieredWithMinimumPriceCadence = "semi_annual"
	PlanVersionNewParamsReplacePricesPriceNewPlanTieredWithMinimumPriceCadenceMonthly    PlanVersionNewParamsReplacePricesPriceNewPlanTieredWithMinimumPriceCadence = "monthly"
	PlanVersionNewParamsReplacePricesPriceNewPlanTieredWithMinimumPriceCadenceQuarterly  PlanVersionNewParamsReplacePricesPriceNewPlanTieredWithMinimumPriceCadence = "quarterly"
	PlanVersionNewParamsReplacePricesPriceNewPlanTieredWithMinimumPriceCadenceOneTime    PlanVersionNewParamsReplacePricesPriceNewPlanTieredWithMinimumPriceCadence = "one_time"
	PlanVersionNewParamsReplacePricesPriceNewPlanTieredWithMinimumPriceCadenceCustom     PlanVersionNewParamsReplacePricesPriceNewPlanTieredWithMinimumPriceCadence = "custom"
)

func (r PlanVersionNewParamsReplacePricesPriceNewPlanTieredWithMinimumPriceCadence) IsKnown() bool {
	switch r {
	case PlanVersionNewParamsReplacePricesPriceNewPlanTieredWithMinimumPriceCadenceAnnual, PlanVersionNewParamsReplacePricesPriceNewPlanTieredWithMinimumPriceCadenceSemiAnnual, PlanVersionNewParamsReplacePricesPriceNewPlanTieredWithMinimumPriceCadenceMonthly, PlanVersionNewParamsReplacePricesPriceNewPlanTieredWithMinimumPriceCadenceQuarterly, PlanVersionNewParamsReplacePricesPriceNewPlanTieredWithMinimumPriceCadenceOneTime, PlanVersionNewParamsReplacePricesPriceNewPlanTieredWithMinimumPriceCadenceCustom:
		return true
	}
	return false
}

type PlanVersionNewParamsReplacePricesPriceNewPlanTieredWithMinimumPriceModelType string

const (
	PlanVersionNewParamsReplacePricesPriceNewPlanTieredWithMinimumPriceModelTypeTieredWithMinimum PlanVersionNewParamsReplacePricesPriceNewPlanTieredWithMinimumPriceModelType = "tiered_with_minimum"
)

func (r PlanVersionNewParamsReplacePricesPriceNewPlanTieredWithMinimumPriceModelType) IsKnown() bool {
	switch r {
	case PlanVersionNewParamsReplacePricesPriceNewPlanTieredWithMinimumPriceModelTypeTieredWithMinimum:
		return true
	}
	return false
}

// For custom cadence: specifies the duration of the billing period in days or
// months.
type PlanVersionNewParamsReplacePricesPriceNewPlanTieredWithMinimumPriceBillingCycleConfiguration struct {
	// The duration of the billing period.
	Duration param.Field[int64] `json:"duration,required"`
	// The unit of billing period duration.
	DurationUnit param.Field[PlanVersionNewParamsReplacePricesPriceNewPlanTieredWithMinimumPriceBillingCycleConfigurationDurationUnit] `json:"duration_unit,required"`
}

func (r PlanVersionNewParamsReplacePricesPriceNewPlanTieredWithMinimumPriceBillingCycleConfiguration) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// The unit of billing period duration.
type PlanVersionNewParamsReplacePricesPriceNewPlanTieredWithMinimumPriceBillingCycleConfigurationDurationUnit string

const (
	PlanVersionNewParamsReplacePricesPriceNewPlanTieredWithMinimumPriceBillingCycleConfigurationDurationUnitDay   PlanVersionNewParamsReplacePricesPriceNewPlanTieredWithMinimumPriceBillingCycleConfigurationDurationUnit = "day"
	PlanVersionNewParamsReplacePricesPriceNewPlanTieredWithMinimumPriceBillingCycleConfigurationDurationUnitMonth PlanVersionNewParamsReplacePricesPriceNewPlanTieredWithMinimumPriceBillingCycleConfigurationDurationUnit = "month"
)

func (r PlanVersionNewParamsReplacePricesPriceNewPlanTieredWithMinimumPriceBillingCycleConfigurationDurationUnit) IsKnown() bool {
	switch r {
	case PlanVersionNewParamsReplacePricesPriceNewPlanTieredWithMinimumPriceBillingCycleConfigurationDurationUnitDay, PlanVersionNewParamsReplacePricesPriceNewPlanTieredWithMinimumPriceBillingCycleConfigurationDurationUnitMonth:
		return true
	}
	return false
}

// For dimensional price: specifies a price group and dimension values
type PlanVersionNewParamsReplacePricesPriceNewPlanTieredWithMinimumPriceDimensionalPriceConfiguration struct {
	// The list of dimension values matching (in order) the dimensions of the price
	// group
	DimensionValues param.Field[[]string] `json:"dimension_values,required"`
	// The id of the dimensional price group to include this price in
	DimensionalPriceGroupID param.Field[string] `json:"dimensional_price_group_id"`
	// The external id of the dimensional price group to include this price in
	ExternalDimensionalPriceGroupID param.Field[string] `json:"external_dimensional_price_group_id"`
}

func (r PlanVersionNewParamsReplacePricesPriceNewPlanTieredWithMinimumPriceDimensionalPriceConfiguration) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Within each billing cycle, specifies the cadence at which invoices are produced.
// If unspecified, a single invoice is produced per billing cycle.
type PlanVersionNewParamsReplacePricesPriceNewPlanTieredWithMinimumPriceInvoicingCycleConfiguration struct {
	// The duration of the billing period.
	Duration param.Field[int64] `json:"duration,required"`
	// The unit of billing period duration.
	DurationUnit param.Field[PlanVersionNewParamsReplacePricesPriceNewPlanTieredWithMinimumPriceInvoicingCycleConfigurationDurationUnit] `json:"duration_unit,required"`
}

func (r PlanVersionNewParamsReplacePricesPriceNewPlanTieredWithMinimumPriceInvoicingCycleConfiguration) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// The unit of billing period duration.
type PlanVersionNewParamsReplacePricesPriceNewPlanTieredWithMinimumPriceInvoicingCycleConfigurationDurationUnit string

const (
	PlanVersionNewParamsReplacePricesPriceNewPlanTieredWithMinimumPriceInvoicingCycleConfigurationDurationUnitDay   PlanVersionNewParamsReplacePricesPriceNewPlanTieredWithMinimumPriceInvoicingCycleConfigurationDurationUnit = "day"
	PlanVersionNewParamsReplacePricesPriceNewPlanTieredWithMinimumPriceInvoicingCycleConfigurationDurationUnitMonth PlanVersionNewParamsReplacePricesPriceNewPlanTieredWithMinimumPriceInvoicingCycleConfigurationDurationUnit = "month"
)

func (r PlanVersionNewParamsReplacePricesPriceNewPlanTieredWithMinimumPriceInvoicingCycleConfigurationDurationUnit) IsKnown() bool {
	switch r {
	case PlanVersionNewParamsReplacePricesPriceNewPlanTieredWithMinimumPriceInvoicingCycleConfigurationDurationUnitDay, PlanVersionNewParamsReplacePricesPriceNewPlanTieredWithMinimumPriceInvoicingCycleConfigurationDurationUnitMonth:
		return true
	}
	return false
}

type PlanVersionNewParamsReplacePricesPriceNewPlanUnitWithPercentPrice struct {
	// The cadence to bill for this price on.
	Cadence param.Field[PlanVersionNewParamsReplacePricesPriceNewPlanUnitWithPercentPriceCadence] `json:"cadence,required"`
	// The id of the item the price will be associated with.
	ItemID    param.Field[string]                                                                     `json:"item_id,required"`
	ModelType param.Field[PlanVersionNewParamsReplacePricesPriceNewPlanUnitWithPercentPriceModelType] `json:"model_type,required"`
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
	BillingCycleConfiguration param.Field[PlanVersionNewParamsReplacePricesPriceNewPlanUnitWithPercentPriceBillingCycleConfiguration] `json:"billing_cycle_configuration"`
	// The per unit conversion rate of the price currency to the invoicing currency.
	ConversionRate param.Field[float64] `json:"conversion_rate"`
	// An ISO 4217 currency string, or custom pricing unit identifier, in which this
	// price is billed.
	Currency param.Field[string] `json:"currency"`
	// For dimensional price: specifies a price group and dimension values
	DimensionalPriceConfiguration param.Field[PlanVersionNewParamsReplacePricesPriceNewPlanUnitWithPercentPriceDimensionalPriceConfiguration] `json:"dimensional_price_configuration"`
	// An alias for the price.
	ExternalPriceID param.Field[string] `json:"external_price_id"`
	// If the Price represents a fixed cost, this represents the quantity of units
	// applied.
	FixedPriceQuantity param.Field[float64] `json:"fixed_price_quantity"`
	// The property used to group this price on an invoice
	InvoiceGroupingKey param.Field[string] `json:"invoice_grouping_key"`
	// Within each billing cycle, specifies the cadence at which invoices are produced.
	// If unspecified, a single invoice is produced per billing cycle.
	InvoicingCycleConfiguration param.Field[PlanVersionNewParamsReplacePricesPriceNewPlanUnitWithPercentPriceInvoicingCycleConfiguration] `json:"invoicing_cycle_configuration"`
	// User-specified key/value pairs for the resource. Individual keys can be removed
	// by setting the value to `null`, and the entire metadata mapping can be cleared
	// by setting `metadata` to `null`.
	Metadata param.Field[map[string]string] `json:"metadata"`
}

func (r PlanVersionNewParamsReplacePricesPriceNewPlanUnitWithPercentPrice) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r PlanVersionNewParamsReplacePricesPriceNewPlanUnitWithPercentPrice) implementsPlanVersionNewParamsReplacePricesPriceUnion() {
}

// The cadence to bill for this price on.
type PlanVersionNewParamsReplacePricesPriceNewPlanUnitWithPercentPriceCadence string

const (
	PlanVersionNewParamsReplacePricesPriceNewPlanUnitWithPercentPriceCadenceAnnual     PlanVersionNewParamsReplacePricesPriceNewPlanUnitWithPercentPriceCadence = "annual"
	PlanVersionNewParamsReplacePricesPriceNewPlanUnitWithPercentPriceCadenceSemiAnnual PlanVersionNewParamsReplacePricesPriceNewPlanUnitWithPercentPriceCadence = "semi_annual"
	PlanVersionNewParamsReplacePricesPriceNewPlanUnitWithPercentPriceCadenceMonthly    PlanVersionNewParamsReplacePricesPriceNewPlanUnitWithPercentPriceCadence = "monthly"
	PlanVersionNewParamsReplacePricesPriceNewPlanUnitWithPercentPriceCadenceQuarterly  PlanVersionNewParamsReplacePricesPriceNewPlanUnitWithPercentPriceCadence = "quarterly"
	PlanVersionNewParamsReplacePricesPriceNewPlanUnitWithPercentPriceCadenceOneTime    PlanVersionNewParamsReplacePricesPriceNewPlanUnitWithPercentPriceCadence = "one_time"
	PlanVersionNewParamsReplacePricesPriceNewPlanUnitWithPercentPriceCadenceCustom     PlanVersionNewParamsReplacePricesPriceNewPlanUnitWithPercentPriceCadence = "custom"
)

func (r PlanVersionNewParamsReplacePricesPriceNewPlanUnitWithPercentPriceCadence) IsKnown() bool {
	switch r {
	case PlanVersionNewParamsReplacePricesPriceNewPlanUnitWithPercentPriceCadenceAnnual, PlanVersionNewParamsReplacePricesPriceNewPlanUnitWithPercentPriceCadenceSemiAnnual, PlanVersionNewParamsReplacePricesPriceNewPlanUnitWithPercentPriceCadenceMonthly, PlanVersionNewParamsReplacePricesPriceNewPlanUnitWithPercentPriceCadenceQuarterly, PlanVersionNewParamsReplacePricesPriceNewPlanUnitWithPercentPriceCadenceOneTime, PlanVersionNewParamsReplacePricesPriceNewPlanUnitWithPercentPriceCadenceCustom:
		return true
	}
	return false
}

type PlanVersionNewParamsReplacePricesPriceNewPlanUnitWithPercentPriceModelType string

const (
	PlanVersionNewParamsReplacePricesPriceNewPlanUnitWithPercentPriceModelTypeUnitWithPercent PlanVersionNewParamsReplacePricesPriceNewPlanUnitWithPercentPriceModelType = "unit_with_percent"
)

func (r PlanVersionNewParamsReplacePricesPriceNewPlanUnitWithPercentPriceModelType) IsKnown() bool {
	switch r {
	case PlanVersionNewParamsReplacePricesPriceNewPlanUnitWithPercentPriceModelTypeUnitWithPercent:
		return true
	}
	return false
}

// For custom cadence: specifies the duration of the billing period in days or
// months.
type PlanVersionNewParamsReplacePricesPriceNewPlanUnitWithPercentPriceBillingCycleConfiguration struct {
	// The duration of the billing period.
	Duration param.Field[int64] `json:"duration,required"`
	// The unit of billing period duration.
	DurationUnit param.Field[PlanVersionNewParamsReplacePricesPriceNewPlanUnitWithPercentPriceBillingCycleConfigurationDurationUnit] `json:"duration_unit,required"`
}

func (r PlanVersionNewParamsReplacePricesPriceNewPlanUnitWithPercentPriceBillingCycleConfiguration) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// The unit of billing period duration.
type PlanVersionNewParamsReplacePricesPriceNewPlanUnitWithPercentPriceBillingCycleConfigurationDurationUnit string

const (
	PlanVersionNewParamsReplacePricesPriceNewPlanUnitWithPercentPriceBillingCycleConfigurationDurationUnitDay   PlanVersionNewParamsReplacePricesPriceNewPlanUnitWithPercentPriceBillingCycleConfigurationDurationUnit = "day"
	PlanVersionNewParamsReplacePricesPriceNewPlanUnitWithPercentPriceBillingCycleConfigurationDurationUnitMonth PlanVersionNewParamsReplacePricesPriceNewPlanUnitWithPercentPriceBillingCycleConfigurationDurationUnit = "month"
)

func (r PlanVersionNewParamsReplacePricesPriceNewPlanUnitWithPercentPriceBillingCycleConfigurationDurationUnit) IsKnown() bool {
	switch r {
	case PlanVersionNewParamsReplacePricesPriceNewPlanUnitWithPercentPriceBillingCycleConfigurationDurationUnitDay, PlanVersionNewParamsReplacePricesPriceNewPlanUnitWithPercentPriceBillingCycleConfigurationDurationUnitMonth:
		return true
	}
	return false
}

// For dimensional price: specifies a price group and dimension values
type PlanVersionNewParamsReplacePricesPriceNewPlanUnitWithPercentPriceDimensionalPriceConfiguration struct {
	// The list of dimension values matching (in order) the dimensions of the price
	// group
	DimensionValues param.Field[[]string] `json:"dimension_values,required"`
	// The id of the dimensional price group to include this price in
	DimensionalPriceGroupID param.Field[string] `json:"dimensional_price_group_id"`
	// The external id of the dimensional price group to include this price in
	ExternalDimensionalPriceGroupID param.Field[string] `json:"external_dimensional_price_group_id"`
}

func (r PlanVersionNewParamsReplacePricesPriceNewPlanUnitWithPercentPriceDimensionalPriceConfiguration) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Within each billing cycle, specifies the cadence at which invoices are produced.
// If unspecified, a single invoice is produced per billing cycle.
type PlanVersionNewParamsReplacePricesPriceNewPlanUnitWithPercentPriceInvoicingCycleConfiguration struct {
	// The duration of the billing period.
	Duration param.Field[int64] `json:"duration,required"`
	// The unit of billing period duration.
	DurationUnit param.Field[PlanVersionNewParamsReplacePricesPriceNewPlanUnitWithPercentPriceInvoicingCycleConfigurationDurationUnit] `json:"duration_unit,required"`
}

func (r PlanVersionNewParamsReplacePricesPriceNewPlanUnitWithPercentPriceInvoicingCycleConfiguration) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// The unit of billing period duration.
type PlanVersionNewParamsReplacePricesPriceNewPlanUnitWithPercentPriceInvoicingCycleConfigurationDurationUnit string

const (
	PlanVersionNewParamsReplacePricesPriceNewPlanUnitWithPercentPriceInvoicingCycleConfigurationDurationUnitDay   PlanVersionNewParamsReplacePricesPriceNewPlanUnitWithPercentPriceInvoicingCycleConfigurationDurationUnit = "day"
	PlanVersionNewParamsReplacePricesPriceNewPlanUnitWithPercentPriceInvoicingCycleConfigurationDurationUnitMonth PlanVersionNewParamsReplacePricesPriceNewPlanUnitWithPercentPriceInvoicingCycleConfigurationDurationUnit = "month"
)

func (r PlanVersionNewParamsReplacePricesPriceNewPlanUnitWithPercentPriceInvoicingCycleConfigurationDurationUnit) IsKnown() bool {
	switch r {
	case PlanVersionNewParamsReplacePricesPriceNewPlanUnitWithPercentPriceInvoicingCycleConfigurationDurationUnitDay, PlanVersionNewParamsReplacePricesPriceNewPlanUnitWithPercentPriceInvoicingCycleConfigurationDurationUnitMonth:
		return true
	}
	return false
}

type PlanVersionNewParamsReplacePricesPriceNewPlanPackageWithAllocationPrice struct {
	// The cadence to bill for this price on.
	Cadence param.Field[PlanVersionNewParamsReplacePricesPriceNewPlanPackageWithAllocationPriceCadence] `json:"cadence,required"`
	// The id of the item the price will be associated with.
	ItemID    param.Field[string]                                                                           `json:"item_id,required"`
	ModelType param.Field[PlanVersionNewParamsReplacePricesPriceNewPlanPackageWithAllocationPriceModelType] `json:"model_type,required"`
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
	BillingCycleConfiguration param.Field[PlanVersionNewParamsReplacePricesPriceNewPlanPackageWithAllocationPriceBillingCycleConfiguration] `json:"billing_cycle_configuration"`
	// The per unit conversion rate of the price currency to the invoicing currency.
	ConversionRate param.Field[float64] `json:"conversion_rate"`
	// An ISO 4217 currency string, or custom pricing unit identifier, in which this
	// price is billed.
	Currency param.Field[string] `json:"currency"`
	// For dimensional price: specifies a price group and dimension values
	DimensionalPriceConfiguration param.Field[PlanVersionNewParamsReplacePricesPriceNewPlanPackageWithAllocationPriceDimensionalPriceConfiguration] `json:"dimensional_price_configuration"`
	// An alias for the price.
	ExternalPriceID param.Field[string] `json:"external_price_id"`
	// If the Price represents a fixed cost, this represents the quantity of units
	// applied.
	FixedPriceQuantity param.Field[float64] `json:"fixed_price_quantity"`
	// The property used to group this price on an invoice
	InvoiceGroupingKey param.Field[string] `json:"invoice_grouping_key"`
	// Within each billing cycle, specifies the cadence at which invoices are produced.
	// If unspecified, a single invoice is produced per billing cycle.
	InvoicingCycleConfiguration param.Field[PlanVersionNewParamsReplacePricesPriceNewPlanPackageWithAllocationPriceInvoicingCycleConfiguration] `json:"invoicing_cycle_configuration"`
	// User-specified key/value pairs for the resource. Individual keys can be removed
	// by setting the value to `null`, and the entire metadata mapping can be cleared
	// by setting `metadata` to `null`.
	Metadata param.Field[map[string]string] `json:"metadata"`
}

func (r PlanVersionNewParamsReplacePricesPriceNewPlanPackageWithAllocationPrice) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r PlanVersionNewParamsReplacePricesPriceNewPlanPackageWithAllocationPrice) implementsPlanVersionNewParamsReplacePricesPriceUnion() {
}

// The cadence to bill for this price on.
type PlanVersionNewParamsReplacePricesPriceNewPlanPackageWithAllocationPriceCadence string

const (
	PlanVersionNewParamsReplacePricesPriceNewPlanPackageWithAllocationPriceCadenceAnnual     PlanVersionNewParamsReplacePricesPriceNewPlanPackageWithAllocationPriceCadence = "annual"
	PlanVersionNewParamsReplacePricesPriceNewPlanPackageWithAllocationPriceCadenceSemiAnnual PlanVersionNewParamsReplacePricesPriceNewPlanPackageWithAllocationPriceCadence = "semi_annual"
	PlanVersionNewParamsReplacePricesPriceNewPlanPackageWithAllocationPriceCadenceMonthly    PlanVersionNewParamsReplacePricesPriceNewPlanPackageWithAllocationPriceCadence = "monthly"
	PlanVersionNewParamsReplacePricesPriceNewPlanPackageWithAllocationPriceCadenceQuarterly  PlanVersionNewParamsReplacePricesPriceNewPlanPackageWithAllocationPriceCadence = "quarterly"
	PlanVersionNewParamsReplacePricesPriceNewPlanPackageWithAllocationPriceCadenceOneTime    PlanVersionNewParamsReplacePricesPriceNewPlanPackageWithAllocationPriceCadence = "one_time"
	PlanVersionNewParamsReplacePricesPriceNewPlanPackageWithAllocationPriceCadenceCustom     PlanVersionNewParamsReplacePricesPriceNewPlanPackageWithAllocationPriceCadence = "custom"
)

func (r PlanVersionNewParamsReplacePricesPriceNewPlanPackageWithAllocationPriceCadence) IsKnown() bool {
	switch r {
	case PlanVersionNewParamsReplacePricesPriceNewPlanPackageWithAllocationPriceCadenceAnnual, PlanVersionNewParamsReplacePricesPriceNewPlanPackageWithAllocationPriceCadenceSemiAnnual, PlanVersionNewParamsReplacePricesPriceNewPlanPackageWithAllocationPriceCadenceMonthly, PlanVersionNewParamsReplacePricesPriceNewPlanPackageWithAllocationPriceCadenceQuarterly, PlanVersionNewParamsReplacePricesPriceNewPlanPackageWithAllocationPriceCadenceOneTime, PlanVersionNewParamsReplacePricesPriceNewPlanPackageWithAllocationPriceCadenceCustom:
		return true
	}
	return false
}

type PlanVersionNewParamsReplacePricesPriceNewPlanPackageWithAllocationPriceModelType string

const (
	PlanVersionNewParamsReplacePricesPriceNewPlanPackageWithAllocationPriceModelTypePackageWithAllocation PlanVersionNewParamsReplacePricesPriceNewPlanPackageWithAllocationPriceModelType = "package_with_allocation"
)

func (r PlanVersionNewParamsReplacePricesPriceNewPlanPackageWithAllocationPriceModelType) IsKnown() bool {
	switch r {
	case PlanVersionNewParamsReplacePricesPriceNewPlanPackageWithAllocationPriceModelTypePackageWithAllocation:
		return true
	}
	return false
}

// For custom cadence: specifies the duration of the billing period in days or
// months.
type PlanVersionNewParamsReplacePricesPriceNewPlanPackageWithAllocationPriceBillingCycleConfiguration struct {
	// The duration of the billing period.
	Duration param.Field[int64] `json:"duration,required"`
	// The unit of billing period duration.
	DurationUnit param.Field[PlanVersionNewParamsReplacePricesPriceNewPlanPackageWithAllocationPriceBillingCycleConfigurationDurationUnit] `json:"duration_unit,required"`
}

func (r PlanVersionNewParamsReplacePricesPriceNewPlanPackageWithAllocationPriceBillingCycleConfiguration) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// The unit of billing period duration.
type PlanVersionNewParamsReplacePricesPriceNewPlanPackageWithAllocationPriceBillingCycleConfigurationDurationUnit string

const (
	PlanVersionNewParamsReplacePricesPriceNewPlanPackageWithAllocationPriceBillingCycleConfigurationDurationUnitDay   PlanVersionNewParamsReplacePricesPriceNewPlanPackageWithAllocationPriceBillingCycleConfigurationDurationUnit = "day"
	PlanVersionNewParamsReplacePricesPriceNewPlanPackageWithAllocationPriceBillingCycleConfigurationDurationUnitMonth PlanVersionNewParamsReplacePricesPriceNewPlanPackageWithAllocationPriceBillingCycleConfigurationDurationUnit = "month"
)

func (r PlanVersionNewParamsReplacePricesPriceNewPlanPackageWithAllocationPriceBillingCycleConfigurationDurationUnit) IsKnown() bool {
	switch r {
	case PlanVersionNewParamsReplacePricesPriceNewPlanPackageWithAllocationPriceBillingCycleConfigurationDurationUnitDay, PlanVersionNewParamsReplacePricesPriceNewPlanPackageWithAllocationPriceBillingCycleConfigurationDurationUnitMonth:
		return true
	}
	return false
}

// For dimensional price: specifies a price group and dimension values
type PlanVersionNewParamsReplacePricesPriceNewPlanPackageWithAllocationPriceDimensionalPriceConfiguration struct {
	// The list of dimension values matching (in order) the dimensions of the price
	// group
	DimensionValues param.Field[[]string] `json:"dimension_values,required"`
	// The id of the dimensional price group to include this price in
	DimensionalPriceGroupID param.Field[string] `json:"dimensional_price_group_id"`
	// The external id of the dimensional price group to include this price in
	ExternalDimensionalPriceGroupID param.Field[string] `json:"external_dimensional_price_group_id"`
}

func (r PlanVersionNewParamsReplacePricesPriceNewPlanPackageWithAllocationPriceDimensionalPriceConfiguration) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Within each billing cycle, specifies the cadence at which invoices are produced.
// If unspecified, a single invoice is produced per billing cycle.
type PlanVersionNewParamsReplacePricesPriceNewPlanPackageWithAllocationPriceInvoicingCycleConfiguration struct {
	// The duration of the billing period.
	Duration param.Field[int64] `json:"duration,required"`
	// The unit of billing period duration.
	DurationUnit param.Field[PlanVersionNewParamsReplacePricesPriceNewPlanPackageWithAllocationPriceInvoicingCycleConfigurationDurationUnit] `json:"duration_unit,required"`
}

func (r PlanVersionNewParamsReplacePricesPriceNewPlanPackageWithAllocationPriceInvoicingCycleConfiguration) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// The unit of billing period duration.
type PlanVersionNewParamsReplacePricesPriceNewPlanPackageWithAllocationPriceInvoicingCycleConfigurationDurationUnit string

const (
	PlanVersionNewParamsReplacePricesPriceNewPlanPackageWithAllocationPriceInvoicingCycleConfigurationDurationUnitDay   PlanVersionNewParamsReplacePricesPriceNewPlanPackageWithAllocationPriceInvoicingCycleConfigurationDurationUnit = "day"
	PlanVersionNewParamsReplacePricesPriceNewPlanPackageWithAllocationPriceInvoicingCycleConfigurationDurationUnitMonth PlanVersionNewParamsReplacePricesPriceNewPlanPackageWithAllocationPriceInvoicingCycleConfigurationDurationUnit = "month"
)

func (r PlanVersionNewParamsReplacePricesPriceNewPlanPackageWithAllocationPriceInvoicingCycleConfigurationDurationUnit) IsKnown() bool {
	switch r {
	case PlanVersionNewParamsReplacePricesPriceNewPlanPackageWithAllocationPriceInvoicingCycleConfigurationDurationUnitDay, PlanVersionNewParamsReplacePricesPriceNewPlanPackageWithAllocationPriceInvoicingCycleConfigurationDurationUnitMonth:
		return true
	}
	return false
}

type PlanVersionNewParamsReplacePricesPriceNewPlanTierWithProrationPrice struct {
	// The cadence to bill for this price on.
	Cadence param.Field[PlanVersionNewParamsReplacePricesPriceNewPlanTierWithProrationPriceCadence] `json:"cadence,required"`
	// The id of the item the price will be associated with.
	ItemID    param.Field[string]                                                                       `json:"item_id,required"`
	ModelType param.Field[PlanVersionNewParamsReplacePricesPriceNewPlanTierWithProrationPriceModelType] `json:"model_type,required"`
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
	BillingCycleConfiguration param.Field[PlanVersionNewParamsReplacePricesPriceNewPlanTierWithProrationPriceBillingCycleConfiguration] `json:"billing_cycle_configuration"`
	// The per unit conversion rate of the price currency to the invoicing currency.
	ConversionRate param.Field[float64] `json:"conversion_rate"`
	// An ISO 4217 currency string, or custom pricing unit identifier, in which this
	// price is billed.
	Currency param.Field[string] `json:"currency"`
	// For dimensional price: specifies a price group and dimension values
	DimensionalPriceConfiguration param.Field[PlanVersionNewParamsReplacePricesPriceNewPlanTierWithProrationPriceDimensionalPriceConfiguration] `json:"dimensional_price_configuration"`
	// An alias for the price.
	ExternalPriceID param.Field[string] `json:"external_price_id"`
	// If the Price represents a fixed cost, this represents the quantity of units
	// applied.
	FixedPriceQuantity param.Field[float64] `json:"fixed_price_quantity"`
	// The property used to group this price on an invoice
	InvoiceGroupingKey param.Field[string] `json:"invoice_grouping_key"`
	// Within each billing cycle, specifies the cadence at which invoices are produced.
	// If unspecified, a single invoice is produced per billing cycle.
	InvoicingCycleConfiguration param.Field[PlanVersionNewParamsReplacePricesPriceNewPlanTierWithProrationPriceInvoicingCycleConfiguration] `json:"invoicing_cycle_configuration"`
	// User-specified key/value pairs for the resource. Individual keys can be removed
	// by setting the value to `null`, and the entire metadata mapping can be cleared
	// by setting `metadata` to `null`.
	Metadata param.Field[map[string]string] `json:"metadata"`
}

func (r PlanVersionNewParamsReplacePricesPriceNewPlanTierWithProrationPrice) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r PlanVersionNewParamsReplacePricesPriceNewPlanTierWithProrationPrice) implementsPlanVersionNewParamsReplacePricesPriceUnion() {
}

// The cadence to bill for this price on.
type PlanVersionNewParamsReplacePricesPriceNewPlanTierWithProrationPriceCadence string

const (
	PlanVersionNewParamsReplacePricesPriceNewPlanTierWithProrationPriceCadenceAnnual     PlanVersionNewParamsReplacePricesPriceNewPlanTierWithProrationPriceCadence = "annual"
	PlanVersionNewParamsReplacePricesPriceNewPlanTierWithProrationPriceCadenceSemiAnnual PlanVersionNewParamsReplacePricesPriceNewPlanTierWithProrationPriceCadence = "semi_annual"
	PlanVersionNewParamsReplacePricesPriceNewPlanTierWithProrationPriceCadenceMonthly    PlanVersionNewParamsReplacePricesPriceNewPlanTierWithProrationPriceCadence = "monthly"
	PlanVersionNewParamsReplacePricesPriceNewPlanTierWithProrationPriceCadenceQuarterly  PlanVersionNewParamsReplacePricesPriceNewPlanTierWithProrationPriceCadence = "quarterly"
	PlanVersionNewParamsReplacePricesPriceNewPlanTierWithProrationPriceCadenceOneTime    PlanVersionNewParamsReplacePricesPriceNewPlanTierWithProrationPriceCadence = "one_time"
	PlanVersionNewParamsReplacePricesPriceNewPlanTierWithProrationPriceCadenceCustom     PlanVersionNewParamsReplacePricesPriceNewPlanTierWithProrationPriceCadence = "custom"
)

func (r PlanVersionNewParamsReplacePricesPriceNewPlanTierWithProrationPriceCadence) IsKnown() bool {
	switch r {
	case PlanVersionNewParamsReplacePricesPriceNewPlanTierWithProrationPriceCadenceAnnual, PlanVersionNewParamsReplacePricesPriceNewPlanTierWithProrationPriceCadenceSemiAnnual, PlanVersionNewParamsReplacePricesPriceNewPlanTierWithProrationPriceCadenceMonthly, PlanVersionNewParamsReplacePricesPriceNewPlanTierWithProrationPriceCadenceQuarterly, PlanVersionNewParamsReplacePricesPriceNewPlanTierWithProrationPriceCadenceOneTime, PlanVersionNewParamsReplacePricesPriceNewPlanTierWithProrationPriceCadenceCustom:
		return true
	}
	return false
}

type PlanVersionNewParamsReplacePricesPriceNewPlanTierWithProrationPriceModelType string

const (
	PlanVersionNewParamsReplacePricesPriceNewPlanTierWithProrationPriceModelTypeTieredWithProration PlanVersionNewParamsReplacePricesPriceNewPlanTierWithProrationPriceModelType = "tiered_with_proration"
)

func (r PlanVersionNewParamsReplacePricesPriceNewPlanTierWithProrationPriceModelType) IsKnown() bool {
	switch r {
	case PlanVersionNewParamsReplacePricesPriceNewPlanTierWithProrationPriceModelTypeTieredWithProration:
		return true
	}
	return false
}

// For custom cadence: specifies the duration of the billing period in days or
// months.
type PlanVersionNewParamsReplacePricesPriceNewPlanTierWithProrationPriceBillingCycleConfiguration struct {
	// The duration of the billing period.
	Duration param.Field[int64] `json:"duration,required"`
	// The unit of billing period duration.
	DurationUnit param.Field[PlanVersionNewParamsReplacePricesPriceNewPlanTierWithProrationPriceBillingCycleConfigurationDurationUnit] `json:"duration_unit,required"`
}

func (r PlanVersionNewParamsReplacePricesPriceNewPlanTierWithProrationPriceBillingCycleConfiguration) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// The unit of billing period duration.
type PlanVersionNewParamsReplacePricesPriceNewPlanTierWithProrationPriceBillingCycleConfigurationDurationUnit string

const (
	PlanVersionNewParamsReplacePricesPriceNewPlanTierWithProrationPriceBillingCycleConfigurationDurationUnitDay   PlanVersionNewParamsReplacePricesPriceNewPlanTierWithProrationPriceBillingCycleConfigurationDurationUnit = "day"
	PlanVersionNewParamsReplacePricesPriceNewPlanTierWithProrationPriceBillingCycleConfigurationDurationUnitMonth PlanVersionNewParamsReplacePricesPriceNewPlanTierWithProrationPriceBillingCycleConfigurationDurationUnit = "month"
)

func (r PlanVersionNewParamsReplacePricesPriceNewPlanTierWithProrationPriceBillingCycleConfigurationDurationUnit) IsKnown() bool {
	switch r {
	case PlanVersionNewParamsReplacePricesPriceNewPlanTierWithProrationPriceBillingCycleConfigurationDurationUnitDay, PlanVersionNewParamsReplacePricesPriceNewPlanTierWithProrationPriceBillingCycleConfigurationDurationUnitMonth:
		return true
	}
	return false
}

// For dimensional price: specifies a price group and dimension values
type PlanVersionNewParamsReplacePricesPriceNewPlanTierWithProrationPriceDimensionalPriceConfiguration struct {
	// The list of dimension values matching (in order) the dimensions of the price
	// group
	DimensionValues param.Field[[]string] `json:"dimension_values,required"`
	// The id of the dimensional price group to include this price in
	DimensionalPriceGroupID param.Field[string] `json:"dimensional_price_group_id"`
	// The external id of the dimensional price group to include this price in
	ExternalDimensionalPriceGroupID param.Field[string] `json:"external_dimensional_price_group_id"`
}

func (r PlanVersionNewParamsReplacePricesPriceNewPlanTierWithProrationPriceDimensionalPriceConfiguration) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Within each billing cycle, specifies the cadence at which invoices are produced.
// If unspecified, a single invoice is produced per billing cycle.
type PlanVersionNewParamsReplacePricesPriceNewPlanTierWithProrationPriceInvoicingCycleConfiguration struct {
	// The duration of the billing period.
	Duration param.Field[int64] `json:"duration,required"`
	// The unit of billing period duration.
	DurationUnit param.Field[PlanVersionNewParamsReplacePricesPriceNewPlanTierWithProrationPriceInvoicingCycleConfigurationDurationUnit] `json:"duration_unit,required"`
}

func (r PlanVersionNewParamsReplacePricesPriceNewPlanTierWithProrationPriceInvoicingCycleConfiguration) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// The unit of billing period duration.
type PlanVersionNewParamsReplacePricesPriceNewPlanTierWithProrationPriceInvoicingCycleConfigurationDurationUnit string

const (
	PlanVersionNewParamsReplacePricesPriceNewPlanTierWithProrationPriceInvoicingCycleConfigurationDurationUnitDay   PlanVersionNewParamsReplacePricesPriceNewPlanTierWithProrationPriceInvoicingCycleConfigurationDurationUnit = "day"
	PlanVersionNewParamsReplacePricesPriceNewPlanTierWithProrationPriceInvoicingCycleConfigurationDurationUnitMonth PlanVersionNewParamsReplacePricesPriceNewPlanTierWithProrationPriceInvoicingCycleConfigurationDurationUnit = "month"
)

func (r PlanVersionNewParamsReplacePricesPriceNewPlanTierWithProrationPriceInvoicingCycleConfigurationDurationUnit) IsKnown() bool {
	switch r {
	case PlanVersionNewParamsReplacePricesPriceNewPlanTierWithProrationPriceInvoicingCycleConfigurationDurationUnitDay, PlanVersionNewParamsReplacePricesPriceNewPlanTierWithProrationPriceInvoicingCycleConfigurationDurationUnitMonth:
		return true
	}
	return false
}

type PlanVersionNewParamsReplacePricesPriceNewPlanUnitWithProrationPrice struct {
	// The cadence to bill for this price on.
	Cadence param.Field[PlanVersionNewParamsReplacePricesPriceNewPlanUnitWithProrationPriceCadence] `json:"cadence,required"`
	// The id of the item the price will be associated with.
	ItemID    param.Field[string]                                                                       `json:"item_id,required"`
	ModelType param.Field[PlanVersionNewParamsReplacePricesPriceNewPlanUnitWithProrationPriceModelType] `json:"model_type,required"`
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
	BillingCycleConfiguration param.Field[PlanVersionNewParamsReplacePricesPriceNewPlanUnitWithProrationPriceBillingCycleConfiguration] `json:"billing_cycle_configuration"`
	// The per unit conversion rate of the price currency to the invoicing currency.
	ConversionRate param.Field[float64] `json:"conversion_rate"`
	// An ISO 4217 currency string, or custom pricing unit identifier, in which this
	// price is billed.
	Currency param.Field[string] `json:"currency"`
	// For dimensional price: specifies a price group and dimension values
	DimensionalPriceConfiguration param.Field[PlanVersionNewParamsReplacePricesPriceNewPlanUnitWithProrationPriceDimensionalPriceConfiguration] `json:"dimensional_price_configuration"`
	// An alias for the price.
	ExternalPriceID param.Field[string] `json:"external_price_id"`
	// If the Price represents a fixed cost, this represents the quantity of units
	// applied.
	FixedPriceQuantity param.Field[float64] `json:"fixed_price_quantity"`
	// The property used to group this price on an invoice
	InvoiceGroupingKey param.Field[string] `json:"invoice_grouping_key"`
	// Within each billing cycle, specifies the cadence at which invoices are produced.
	// If unspecified, a single invoice is produced per billing cycle.
	InvoicingCycleConfiguration param.Field[PlanVersionNewParamsReplacePricesPriceNewPlanUnitWithProrationPriceInvoicingCycleConfiguration] `json:"invoicing_cycle_configuration"`
	// User-specified key/value pairs for the resource. Individual keys can be removed
	// by setting the value to `null`, and the entire metadata mapping can be cleared
	// by setting `metadata` to `null`.
	Metadata param.Field[map[string]string] `json:"metadata"`
}

func (r PlanVersionNewParamsReplacePricesPriceNewPlanUnitWithProrationPrice) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r PlanVersionNewParamsReplacePricesPriceNewPlanUnitWithProrationPrice) implementsPlanVersionNewParamsReplacePricesPriceUnion() {
}

// The cadence to bill for this price on.
type PlanVersionNewParamsReplacePricesPriceNewPlanUnitWithProrationPriceCadence string

const (
	PlanVersionNewParamsReplacePricesPriceNewPlanUnitWithProrationPriceCadenceAnnual     PlanVersionNewParamsReplacePricesPriceNewPlanUnitWithProrationPriceCadence = "annual"
	PlanVersionNewParamsReplacePricesPriceNewPlanUnitWithProrationPriceCadenceSemiAnnual PlanVersionNewParamsReplacePricesPriceNewPlanUnitWithProrationPriceCadence = "semi_annual"
	PlanVersionNewParamsReplacePricesPriceNewPlanUnitWithProrationPriceCadenceMonthly    PlanVersionNewParamsReplacePricesPriceNewPlanUnitWithProrationPriceCadence = "monthly"
	PlanVersionNewParamsReplacePricesPriceNewPlanUnitWithProrationPriceCadenceQuarterly  PlanVersionNewParamsReplacePricesPriceNewPlanUnitWithProrationPriceCadence = "quarterly"
	PlanVersionNewParamsReplacePricesPriceNewPlanUnitWithProrationPriceCadenceOneTime    PlanVersionNewParamsReplacePricesPriceNewPlanUnitWithProrationPriceCadence = "one_time"
	PlanVersionNewParamsReplacePricesPriceNewPlanUnitWithProrationPriceCadenceCustom     PlanVersionNewParamsReplacePricesPriceNewPlanUnitWithProrationPriceCadence = "custom"
)

func (r PlanVersionNewParamsReplacePricesPriceNewPlanUnitWithProrationPriceCadence) IsKnown() bool {
	switch r {
	case PlanVersionNewParamsReplacePricesPriceNewPlanUnitWithProrationPriceCadenceAnnual, PlanVersionNewParamsReplacePricesPriceNewPlanUnitWithProrationPriceCadenceSemiAnnual, PlanVersionNewParamsReplacePricesPriceNewPlanUnitWithProrationPriceCadenceMonthly, PlanVersionNewParamsReplacePricesPriceNewPlanUnitWithProrationPriceCadenceQuarterly, PlanVersionNewParamsReplacePricesPriceNewPlanUnitWithProrationPriceCadenceOneTime, PlanVersionNewParamsReplacePricesPriceNewPlanUnitWithProrationPriceCadenceCustom:
		return true
	}
	return false
}

type PlanVersionNewParamsReplacePricesPriceNewPlanUnitWithProrationPriceModelType string

const (
	PlanVersionNewParamsReplacePricesPriceNewPlanUnitWithProrationPriceModelTypeUnitWithProration PlanVersionNewParamsReplacePricesPriceNewPlanUnitWithProrationPriceModelType = "unit_with_proration"
)

func (r PlanVersionNewParamsReplacePricesPriceNewPlanUnitWithProrationPriceModelType) IsKnown() bool {
	switch r {
	case PlanVersionNewParamsReplacePricesPriceNewPlanUnitWithProrationPriceModelTypeUnitWithProration:
		return true
	}
	return false
}

// For custom cadence: specifies the duration of the billing period in days or
// months.
type PlanVersionNewParamsReplacePricesPriceNewPlanUnitWithProrationPriceBillingCycleConfiguration struct {
	// The duration of the billing period.
	Duration param.Field[int64] `json:"duration,required"`
	// The unit of billing period duration.
	DurationUnit param.Field[PlanVersionNewParamsReplacePricesPriceNewPlanUnitWithProrationPriceBillingCycleConfigurationDurationUnit] `json:"duration_unit,required"`
}

func (r PlanVersionNewParamsReplacePricesPriceNewPlanUnitWithProrationPriceBillingCycleConfiguration) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// The unit of billing period duration.
type PlanVersionNewParamsReplacePricesPriceNewPlanUnitWithProrationPriceBillingCycleConfigurationDurationUnit string

const (
	PlanVersionNewParamsReplacePricesPriceNewPlanUnitWithProrationPriceBillingCycleConfigurationDurationUnitDay   PlanVersionNewParamsReplacePricesPriceNewPlanUnitWithProrationPriceBillingCycleConfigurationDurationUnit = "day"
	PlanVersionNewParamsReplacePricesPriceNewPlanUnitWithProrationPriceBillingCycleConfigurationDurationUnitMonth PlanVersionNewParamsReplacePricesPriceNewPlanUnitWithProrationPriceBillingCycleConfigurationDurationUnit = "month"
)

func (r PlanVersionNewParamsReplacePricesPriceNewPlanUnitWithProrationPriceBillingCycleConfigurationDurationUnit) IsKnown() bool {
	switch r {
	case PlanVersionNewParamsReplacePricesPriceNewPlanUnitWithProrationPriceBillingCycleConfigurationDurationUnitDay, PlanVersionNewParamsReplacePricesPriceNewPlanUnitWithProrationPriceBillingCycleConfigurationDurationUnitMonth:
		return true
	}
	return false
}

// For dimensional price: specifies a price group and dimension values
type PlanVersionNewParamsReplacePricesPriceNewPlanUnitWithProrationPriceDimensionalPriceConfiguration struct {
	// The list of dimension values matching (in order) the dimensions of the price
	// group
	DimensionValues param.Field[[]string] `json:"dimension_values,required"`
	// The id of the dimensional price group to include this price in
	DimensionalPriceGroupID param.Field[string] `json:"dimensional_price_group_id"`
	// The external id of the dimensional price group to include this price in
	ExternalDimensionalPriceGroupID param.Field[string] `json:"external_dimensional_price_group_id"`
}

func (r PlanVersionNewParamsReplacePricesPriceNewPlanUnitWithProrationPriceDimensionalPriceConfiguration) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Within each billing cycle, specifies the cadence at which invoices are produced.
// If unspecified, a single invoice is produced per billing cycle.
type PlanVersionNewParamsReplacePricesPriceNewPlanUnitWithProrationPriceInvoicingCycleConfiguration struct {
	// The duration of the billing period.
	Duration param.Field[int64] `json:"duration,required"`
	// The unit of billing period duration.
	DurationUnit param.Field[PlanVersionNewParamsReplacePricesPriceNewPlanUnitWithProrationPriceInvoicingCycleConfigurationDurationUnit] `json:"duration_unit,required"`
}

func (r PlanVersionNewParamsReplacePricesPriceNewPlanUnitWithProrationPriceInvoicingCycleConfiguration) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// The unit of billing period duration.
type PlanVersionNewParamsReplacePricesPriceNewPlanUnitWithProrationPriceInvoicingCycleConfigurationDurationUnit string

const (
	PlanVersionNewParamsReplacePricesPriceNewPlanUnitWithProrationPriceInvoicingCycleConfigurationDurationUnitDay   PlanVersionNewParamsReplacePricesPriceNewPlanUnitWithProrationPriceInvoicingCycleConfigurationDurationUnit = "day"
	PlanVersionNewParamsReplacePricesPriceNewPlanUnitWithProrationPriceInvoicingCycleConfigurationDurationUnitMonth PlanVersionNewParamsReplacePricesPriceNewPlanUnitWithProrationPriceInvoicingCycleConfigurationDurationUnit = "month"
)

func (r PlanVersionNewParamsReplacePricesPriceNewPlanUnitWithProrationPriceInvoicingCycleConfigurationDurationUnit) IsKnown() bool {
	switch r {
	case PlanVersionNewParamsReplacePricesPriceNewPlanUnitWithProrationPriceInvoicingCycleConfigurationDurationUnitDay, PlanVersionNewParamsReplacePricesPriceNewPlanUnitWithProrationPriceInvoicingCycleConfigurationDurationUnitMonth:
		return true
	}
	return false
}

type PlanVersionNewParamsReplacePricesPriceNewPlanGroupedAllocationPrice struct {
	// The cadence to bill for this price on.
	Cadence                 param.Field[PlanVersionNewParamsReplacePricesPriceNewPlanGroupedAllocationPriceCadence] `json:"cadence,required"`
	GroupedAllocationConfig param.Field[map[string]interface{}]                                                     `json:"grouped_allocation_config,required"`
	// The id of the item the price will be associated with.
	ItemID    param.Field[string]                                                                       `json:"item_id,required"`
	ModelType param.Field[PlanVersionNewParamsReplacePricesPriceNewPlanGroupedAllocationPriceModelType] `json:"model_type,required"`
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
	BillingCycleConfiguration param.Field[PlanVersionNewParamsReplacePricesPriceNewPlanGroupedAllocationPriceBillingCycleConfiguration] `json:"billing_cycle_configuration"`
	// The per unit conversion rate of the price currency to the invoicing currency.
	ConversionRate param.Field[float64] `json:"conversion_rate"`
	// An ISO 4217 currency string, or custom pricing unit identifier, in which this
	// price is billed.
	Currency param.Field[string] `json:"currency"`
	// For dimensional price: specifies a price group and dimension values
	DimensionalPriceConfiguration param.Field[PlanVersionNewParamsReplacePricesPriceNewPlanGroupedAllocationPriceDimensionalPriceConfiguration] `json:"dimensional_price_configuration"`
	// An alias for the price.
	ExternalPriceID param.Field[string] `json:"external_price_id"`
	// If the Price represents a fixed cost, this represents the quantity of units
	// applied.
	FixedPriceQuantity param.Field[float64] `json:"fixed_price_quantity"`
	// The property used to group this price on an invoice
	InvoiceGroupingKey param.Field[string] `json:"invoice_grouping_key"`
	// Within each billing cycle, specifies the cadence at which invoices are produced.
	// If unspecified, a single invoice is produced per billing cycle.
	InvoicingCycleConfiguration param.Field[PlanVersionNewParamsReplacePricesPriceNewPlanGroupedAllocationPriceInvoicingCycleConfiguration] `json:"invoicing_cycle_configuration"`
	// User-specified key/value pairs for the resource. Individual keys can be removed
	// by setting the value to `null`, and the entire metadata mapping can be cleared
	// by setting `metadata` to `null`.
	Metadata param.Field[map[string]string] `json:"metadata"`
}

func (r PlanVersionNewParamsReplacePricesPriceNewPlanGroupedAllocationPrice) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r PlanVersionNewParamsReplacePricesPriceNewPlanGroupedAllocationPrice) implementsPlanVersionNewParamsReplacePricesPriceUnion() {
}

// The cadence to bill for this price on.
type PlanVersionNewParamsReplacePricesPriceNewPlanGroupedAllocationPriceCadence string

const (
	PlanVersionNewParamsReplacePricesPriceNewPlanGroupedAllocationPriceCadenceAnnual     PlanVersionNewParamsReplacePricesPriceNewPlanGroupedAllocationPriceCadence = "annual"
	PlanVersionNewParamsReplacePricesPriceNewPlanGroupedAllocationPriceCadenceSemiAnnual PlanVersionNewParamsReplacePricesPriceNewPlanGroupedAllocationPriceCadence = "semi_annual"
	PlanVersionNewParamsReplacePricesPriceNewPlanGroupedAllocationPriceCadenceMonthly    PlanVersionNewParamsReplacePricesPriceNewPlanGroupedAllocationPriceCadence = "monthly"
	PlanVersionNewParamsReplacePricesPriceNewPlanGroupedAllocationPriceCadenceQuarterly  PlanVersionNewParamsReplacePricesPriceNewPlanGroupedAllocationPriceCadence = "quarterly"
	PlanVersionNewParamsReplacePricesPriceNewPlanGroupedAllocationPriceCadenceOneTime    PlanVersionNewParamsReplacePricesPriceNewPlanGroupedAllocationPriceCadence = "one_time"
	PlanVersionNewParamsReplacePricesPriceNewPlanGroupedAllocationPriceCadenceCustom     PlanVersionNewParamsReplacePricesPriceNewPlanGroupedAllocationPriceCadence = "custom"
)

func (r PlanVersionNewParamsReplacePricesPriceNewPlanGroupedAllocationPriceCadence) IsKnown() bool {
	switch r {
	case PlanVersionNewParamsReplacePricesPriceNewPlanGroupedAllocationPriceCadenceAnnual, PlanVersionNewParamsReplacePricesPriceNewPlanGroupedAllocationPriceCadenceSemiAnnual, PlanVersionNewParamsReplacePricesPriceNewPlanGroupedAllocationPriceCadenceMonthly, PlanVersionNewParamsReplacePricesPriceNewPlanGroupedAllocationPriceCadenceQuarterly, PlanVersionNewParamsReplacePricesPriceNewPlanGroupedAllocationPriceCadenceOneTime, PlanVersionNewParamsReplacePricesPriceNewPlanGroupedAllocationPriceCadenceCustom:
		return true
	}
	return false
}

type PlanVersionNewParamsReplacePricesPriceNewPlanGroupedAllocationPriceModelType string

const (
	PlanVersionNewParamsReplacePricesPriceNewPlanGroupedAllocationPriceModelTypeGroupedAllocation PlanVersionNewParamsReplacePricesPriceNewPlanGroupedAllocationPriceModelType = "grouped_allocation"
)

func (r PlanVersionNewParamsReplacePricesPriceNewPlanGroupedAllocationPriceModelType) IsKnown() bool {
	switch r {
	case PlanVersionNewParamsReplacePricesPriceNewPlanGroupedAllocationPriceModelTypeGroupedAllocation:
		return true
	}
	return false
}

// For custom cadence: specifies the duration of the billing period in days or
// months.
type PlanVersionNewParamsReplacePricesPriceNewPlanGroupedAllocationPriceBillingCycleConfiguration struct {
	// The duration of the billing period.
	Duration param.Field[int64] `json:"duration,required"`
	// The unit of billing period duration.
	DurationUnit param.Field[PlanVersionNewParamsReplacePricesPriceNewPlanGroupedAllocationPriceBillingCycleConfigurationDurationUnit] `json:"duration_unit,required"`
}

func (r PlanVersionNewParamsReplacePricesPriceNewPlanGroupedAllocationPriceBillingCycleConfiguration) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// The unit of billing period duration.
type PlanVersionNewParamsReplacePricesPriceNewPlanGroupedAllocationPriceBillingCycleConfigurationDurationUnit string

const (
	PlanVersionNewParamsReplacePricesPriceNewPlanGroupedAllocationPriceBillingCycleConfigurationDurationUnitDay   PlanVersionNewParamsReplacePricesPriceNewPlanGroupedAllocationPriceBillingCycleConfigurationDurationUnit = "day"
	PlanVersionNewParamsReplacePricesPriceNewPlanGroupedAllocationPriceBillingCycleConfigurationDurationUnitMonth PlanVersionNewParamsReplacePricesPriceNewPlanGroupedAllocationPriceBillingCycleConfigurationDurationUnit = "month"
)

func (r PlanVersionNewParamsReplacePricesPriceNewPlanGroupedAllocationPriceBillingCycleConfigurationDurationUnit) IsKnown() bool {
	switch r {
	case PlanVersionNewParamsReplacePricesPriceNewPlanGroupedAllocationPriceBillingCycleConfigurationDurationUnitDay, PlanVersionNewParamsReplacePricesPriceNewPlanGroupedAllocationPriceBillingCycleConfigurationDurationUnitMonth:
		return true
	}
	return false
}

// For dimensional price: specifies a price group and dimension values
type PlanVersionNewParamsReplacePricesPriceNewPlanGroupedAllocationPriceDimensionalPriceConfiguration struct {
	// The list of dimension values matching (in order) the dimensions of the price
	// group
	DimensionValues param.Field[[]string] `json:"dimension_values,required"`
	// The id of the dimensional price group to include this price in
	DimensionalPriceGroupID param.Field[string] `json:"dimensional_price_group_id"`
	// The external id of the dimensional price group to include this price in
	ExternalDimensionalPriceGroupID param.Field[string] `json:"external_dimensional_price_group_id"`
}

func (r PlanVersionNewParamsReplacePricesPriceNewPlanGroupedAllocationPriceDimensionalPriceConfiguration) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Within each billing cycle, specifies the cadence at which invoices are produced.
// If unspecified, a single invoice is produced per billing cycle.
type PlanVersionNewParamsReplacePricesPriceNewPlanGroupedAllocationPriceInvoicingCycleConfiguration struct {
	// The duration of the billing period.
	Duration param.Field[int64] `json:"duration,required"`
	// The unit of billing period duration.
	DurationUnit param.Field[PlanVersionNewParamsReplacePricesPriceNewPlanGroupedAllocationPriceInvoicingCycleConfigurationDurationUnit] `json:"duration_unit,required"`
}

func (r PlanVersionNewParamsReplacePricesPriceNewPlanGroupedAllocationPriceInvoicingCycleConfiguration) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// The unit of billing period duration.
type PlanVersionNewParamsReplacePricesPriceNewPlanGroupedAllocationPriceInvoicingCycleConfigurationDurationUnit string

const (
	PlanVersionNewParamsReplacePricesPriceNewPlanGroupedAllocationPriceInvoicingCycleConfigurationDurationUnitDay   PlanVersionNewParamsReplacePricesPriceNewPlanGroupedAllocationPriceInvoicingCycleConfigurationDurationUnit = "day"
	PlanVersionNewParamsReplacePricesPriceNewPlanGroupedAllocationPriceInvoicingCycleConfigurationDurationUnitMonth PlanVersionNewParamsReplacePricesPriceNewPlanGroupedAllocationPriceInvoicingCycleConfigurationDurationUnit = "month"
)

func (r PlanVersionNewParamsReplacePricesPriceNewPlanGroupedAllocationPriceInvoicingCycleConfigurationDurationUnit) IsKnown() bool {
	switch r {
	case PlanVersionNewParamsReplacePricesPriceNewPlanGroupedAllocationPriceInvoicingCycleConfigurationDurationUnitDay, PlanVersionNewParamsReplacePricesPriceNewPlanGroupedAllocationPriceInvoicingCycleConfigurationDurationUnitMonth:
		return true
	}
	return false
}

type PlanVersionNewParamsReplacePricesPriceNewPlanGroupedWithProratedMinimumPrice struct {
	// The cadence to bill for this price on.
	Cadence                          param.Field[PlanVersionNewParamsReplacePricesPriceNewPlanGroupedWithProratedMinimumPriceCadence] `json:"cadence,required"`
	GroupedWithProratedMinimumConfig param.Field[map[string]interface{}]                                                              `json:"grouped_with_prorated_minimum_config,required"`
	// The id of the item the price will be associated with.
	ItemID    param.Field[string]                                                                                `json:"item_id,required"`
	ModelType param.Field[PlanVersionNewParamsReplacePricesPriceNewPlanGroupedWithProratedMinimumPriceModelType] `json:"model_type,required"`
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
	BillingCycleConfiguration param.Field[PlanVersionNewParamsReplacePricesPriceNewPlanGroupedWithProratedMinimumPriceBillingCycleConfiguration] `json:"billing_cycle_configuration"`
	// The per unit conversion rate of the price currency to the invoicing currency.
	ConversionRate param.Field[float64] `json:"conversion_rate"`
	// An ISO 4217 currency string, or custom pricing unit identifier, in which this
	// price is billed.
	Currency param.Field[string] `json:"currency"`
	// For dimensional price: specifies a price group and dimension values
	DimensionalPriceConfiguration param.Field[PlanVersionNewParamsReplacePricesPriceNewPlanGroupedWithProratedMinimumPriceDimensionalPriceConfiguration] `json:"dimensional_price_configuration"`
	// An alias for the price.
	ExternalPriceID param.Field[string] `json:"external_price_id"`
	// If the Price represents a fixed cost, this represents the quantity of units
	// applied.
	FixedPriceQuantity param.Field[float64] `json:"fixed_price_quantity"`
	// The property used to group this price on an invoice
	InvoiceGroupingKey param.Field[string] `json:"invoice_grouping_key"`
	// Within each billing cycle, specifies the cadence at which invoices are produced.
	// If unspecified, a single invoice is produced per billing cycle.
	InvoicingCycleConfiguration param.Field[PlanVersionNewParamsReplacePricesPriceNewPlanGroupedWithProratedMinimumPriceInvoicingCycleConfiguration] `json:"invoicing_cycle_configuration"`
	// User-specified key/value pairs for the resource. Individual keys can be removed
	// by setting the value to `null`, and the entire metadata mapping can be cleared
	// by setting `metadata` to `null`.
	Metadata param.Field[map[string]string] `json:"metadata"`
}

func (r PlanVersionNewParamsReplacePricesPriceNewPlanGroupedWithProratedMinimumPrice) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r PlanVersionNewParamsReplacePricesPriceNewPlanGroupedWithProratedMinimumPrice) implementsPlanVersionNewParamsReplacePricesPriceUnion() {
}

// The cadence to bill for this price on.
type PlanVersionNewParamsReplacePricesPriceNewPlanGroupedWithProratedMinimumPriceCadence string

const (
	PlanVersionNewParamsReplacePricesPriceNewPlanGroupedWithProratedMinimumPriceCadenceAnnual     PlanVersionNewParamsReplacePricesPriceNewPlanGroupedWithProratedMinimumPriceCadence = "annual"
	PlanVersionNewParamsReplacePricesPriceNewPlanGroupedWithProratedMinimumPriceCadenceSemiAnnual PlanVersionNewParamsReplacePricesPriceNewPlanGroupedWithProratedMinimumPriceCadence = "semi_annual"
	PlanVersionNewParamsReplacePricesPriceNewPlanGroupedWithProratedMinimumPriceCadenceMonthly    PlanVersionNewParamsReplacePricesPriceNewPlanGroupedWithProratedMinimumPriceCadence = "monthly"
	PlanVersionNewParamsReplacePricesPriceNewPlanGroupedWithProratedMinimumPriceCadenceQuarterly  PlanVersionNewParamsReplacePricesPriceNewPlanGroupedWithProratedMinimumPriceCadence = "quarterly"
	PlanVersionNewParamsReplacePricesPriceNewPlanGroupedWithProratedMinimumPriceCadenceOneTime    PlanVersionNewParamsReplacePricesPriceNewPlanGroupedWithProratedMinimumPriceCadence = "one_time"
	PlanVersionNewParamsReplacePricesPriceNewPlanGroupedWithProratedMinimumPriceCadenceCustom     PlanVersionNewParamsReplacePricesPriceNewPlanGroupedWithProratedMinimumPriceCadence = "custom"
)

func (r PlanVersionNewParamsReplacePricesPriceNewPlanGroupedWithProratedMinimumPriceCadence) IsKnown() bool {
	switch r {
	case PlanVersionNewParamsReplacePricesPriceNewPlanGroupedWithProratedMinimumPriceCadenceAnnual, PlanVersionNewParamsReplacePricesPriceNewPlanGroupedWithProratedMinimumPriceCadenceSemiAnnual, PlanVersionNewParamsReplacePricesPriceNewPlanGroupedWithProratedMinimumPriceCadenceMonthly, PlanVersionNewParamsReplacePricesPriceNewPlanGroupedWithProratedMinimumPriceCadenceQuarterly, PlanVersionNewParamsReplacePricesPriceNewPlanGroupedWithProratedMinimumPriceCadenceOneTime, PlanVersionNewParamsReplacePricesPriceNewPlanGroupedWithProratedMinimumPriceCadenceCustom:
		return true
	}
	return false
}

type PlanVersionNewParamsReplacePricesPriceNewPlanGroupedWithProratedMinimumPriceModelType string

const (
	PlanVersionNewParamsReplacePricesPriceNewPlanGroupedWithProratedMinimumPriceModelTypeGroupedWithProratedMinimum PlanVersionNewParamsReplacePricesPriceNewPlanGroupedWithProratedMinimumPriceModelType = "grouped_with_prorated_minimum"
)

func (r PlanVersionNewParamsReplacePricesPriceNewPlanGroupedWithProratedMinimumPriceModelType) IsKnown() bool {
	switch r {
	case PlanVersionNewParamsReplacePricesPriceNewPlanGroupedWithProratedMinimumPriceModelTypeGroupedWithProratedMinimum:
		return true
	}
	return false
}

// For custom cadence: specifies the duration of the billing period in days or
// months.
type PlanVersionNewParamsReplacePricesPriceNewPlanGroupedWithProratedMinimumPriceBillingCycleConfiguration struct {
	// The duration of the billing period.
	Duration param.Field[int64] `json:"duration,required"`
	// The unit of billing period duration.
	DurationUnit param.Field[PlanVersionNewParamsReplacePricesPriceNewPlanGroupedWithProratedMinimumPriceBillingCycleConfigurationDurationUnit] `json:"duration_unit,required"`
}

func (r PlanVersionNewParamsReplacePricesPriceNewPlanGroupedWithProratedMinimumPriceBillingCycleConfiguration) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// The unit of billing period duration.
type PlanVersionNewParamsReplacePricesPriceNewPlanGroupedWithProratedMinimumPriceBillingCycleConfigurationDurationUnit string

const (
	PlanVersionNewParamsReplacePricesPriceNewPlanGroupedWithProratedMinimumPriceBillingCycleConfigurationDurationUnitDay   PlanVersionNewParamsReplacePricesPriceNewPlanGroupedWithProratedMinimumPriceBillingCycleConfigurationDurationUnit = "day"
	PlanVersionNewParamsReplacePricesPriceNewPlanGroupedWithProratedMinimumPriceBillingCycleConfigurationDurationUnitMonth PlanVersionNewParamsReplacePricesPriceNewPlanGroupedWithProratedMinimumPriceBillingCycleConfigurationDurationUnit = "month"
)

func (r PlanVersionNewParamsReplacePricesPriceNewPlanGroupedWithProratedMinimumPriceBillingCycleConfigurationDurationUnit) IsKnown() bool {
	switch r {
	case PlanVersionNewParamsReplacePricesPriceNewPlanGroupedWithProratedMinimumPriceBillingCycleConfigurationDurationUnitDay, PlanVersionNewParamsReplacePricesPriceNewPlanGroupedWithProratedMinimumPriceBillingCycleConfigurationDurationUnitMonth:
		return true
	}
	return false
}

// For dimensional price: specifies a price group and dimension values
type PlanVersionNewParamsReplacePricesPriceNewPlanGroupedWithProratedMinimumPriceDimensionalPriceConfiguration struct {
	// The list of dimension values matching (in order) the dimensions of the price
	// group
	DimensionValues param.Field[[]string] `json:"dimension_values,required"`
	// The id of the dimensional price group to include this price in
	DimensionalPriceGroupID param.Field[string] `json:"dimensional_price_group_id"`
	// The external id of the dimensional price group to include this price in
	ExternalDimensionalPriceGroupID param.Field[string] `json:"external_dimensional_price_group_id"`
}

func (r PlanVersionNewParamsReplacePricesPriceNewPlanGroupedWithProratedMinimumPriceDimensionalPriceConfiguration) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Within each billing cycle, specifies the cadence at which invoices are produced.
// If unspecified, a single invoice is produced per billing cycle.
type PlanVersionNewParamsReplacePricesPriceNewPlanGroupedWithProratedMinimumPriceInvoicingCycleConfiguration struct {
	// The duration of the billing period.
	Duration param.Field[int64] `json:"duration,required"`
	// The unit of billing period duration.
	DurationUnit param.Field[PlanVersionNewParamsReplacePricesPriceNewPlanGroupedWithProratedMinimumPriceInvoicingCycleConfigurationDurationUnit] `json:"duration_unit,required"`
}

func (r PlanVersionNewParamsReplacePricesPriceNewPlanGroupedWithProratedMinimumPriceInvoicingCycleConfiguration) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// The unit of billing period duration.
type PlanVersionNewParamsReplacePricesPriceNewPlanGroupedWithProratedMinimumPriceInvoicingCycleConfigurationDurationUnit string

const (
	PlanVersionNewParamsReplacePricesPriceNewPlanGroupedWithProratedMinimumPriceInvoicingCycleConfigurationDurationUnitDay   PlanVersionNewParamsReplacePricesPriceNewPlanGroupedWithProratedMinimumPriceInvoicingCycleConfigurationDurationUnit = "day"
	PlanVersionNewParamsReplacePricesPriceNewPlanGroupedWithProratedMinimumPriceInvoicingCycleConfigurationDurationUnitMonth PlanVersionNewParamsReplacePricesPriceNewPlanGroupedWithProratedMinimumPriceInvoicingCycleConfigurationDurationUnit = "month"
)

func (r PlanVersionNewParamsReplacePricesPriceNewPlanGroupedWithProratedMinimumPriceInvoicingCycleConfigurationDurationUnit) IsKnown() bool {
	switch r {
	case PlanVersionNewParamsReplacePricesPriceNewPlanGroupedWithProratedMinimumPriceInvoicingCycleConfigurationDurationUnitDay, PlanVersionNewParamsReplacePricesPriceNewPlanGroupedWithProratedMinimumPriceInvoicingCycleConfigurationDurationUnitMonth:
		return true
	}
	return false
}

type PlanVersionNewParamsReplacePricesPriceNewPlanGroupedWithMeteredMinimumPrice struct {
	// The cadence to bill for this price on.
	Cadence                         param.Field[PlanVersionNewParamsReplacePricesPriceNewPlanGroupedWithMeteredMinimumPriceCadence] `json:"cadence,required"`
	GroupedWithMeteredMinimumConfig param.Field[map[string]interface{}]                                                             `json:"grouped_with_metered_minimum_config,required"`
	// The id of the item the price will be associated with.
	ItemID    param.Field[string]                                                                               `json:"item_id,required"`
	ModelType param.Field[PlanVersionNewParamsReplacePricesPriceNewPlanGroupedWithMeteredMinimumPriceModelType] `json:"model_type,required"`
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
	BillingCycleConfiguration param.Field[PlanVersionNewParamsReplacePricesPriceNewPlanGroupedWithMeteredMinimumPriceBillingCycleConfiguration] `json:"billing_cycle_configuration"`
	// The per unit conversion rate of the price currency to the invoicing currency.
	ConversionRate param.Field[float64] `json:"conversion_rate"`
	// An ISO 4217 currency string, or custom pricing unit identifier, in which this
	// price is billed.
	Currency param.Field[string] `json:"currency"`
	// For dimensional price: specifies a price group and dimension values
	DimensionalPriceConfiguration param.Field[PlanVersionNewParamsReplacePricesPriceNewPlanGroupedWithMeteredMinimumPriceDimensionalPriceConfiguration] `json:"dimensional_price_configuration"`
	// An alias for the price.
	ExternalPriceID param.Field[string] `json:"external_price_id"`
	// If the Price represents a fixed cost, this represents the quantity of units
	// applied.
	FixedPriceQuantity param.Field[float64] `json:"fixed_price_quantity"`
	// The property used to group this price on an invoice
	InvoiceGroupingKey param.Field[string] `json:"invoice_grouping_key"`
	// Within each billing cycle, specifies the cadence at which invoices are produced.
	// If unspecified, a single invoice is produced per billing cycle.
	InvoicingCycleConfiguration param.Field[PlanVersionNewParamsReplacePricesPriceNewPlanGroupedWithMeteredMinimumPriceInvoicingCycleConfiguration] `json:"invoicing_cycle_configuration"`
	// User-specified key/value pairs for the resource. Individual keys can be removed
	// by setting the value to `null`, and the entire metadata mapping can be cleared
	// by setting `metadata` to `null`.
	Metadata param.Field[map[string]string] `json:"metadata"`
}

func (r PlanVersionNewParamsReplacePricesPriceNewPlanGroupedWithMeteredMinimumPrice) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r PlanVersionNewParamsReplacePricesPriceNewPlanGroupedWithMeteredMinimumPrice) implementsPlanVersionNewParamsReplacePricesPriceUnion() {
}

// The cadence to bill for this price on.
type PlanVersionNewParamsReplacePricesPriceNewPlanGroupedWithMeteredMinimumPriceCadence string

const (
	PlanVersionNewParamsReplacePricesPriceNewPlanGroupedWithMeteredMinimumPriceCadenceAnnual     PlanVersionNewParamsReplacePricesPriceNewPlanGroupedWithMeteredMinimumPriceCadence = "annual"
	PlanVersionNewParamsReplacePricesPriceNewPlanGroupedWithMeteredMinimumPriceCadenceSemiAnnual PlanVersionNewParamsReplacePricesPriceNewPlanGroupedWithMeteredMinimumPriceCadence = "semi_annual"
	PlanVersionNewParamsReplacePricesPriceNewPlanGroupedWithMeteredMinimumPriceCadenceMonthly    PlanVersionNewParamsReplacePricesPriceNewPlanGroupedWithMeteredMinimumPriceCadence = "monthly"
	PlanVersionNewParamsReplacePricesPriceNewPlanGroupedWithMeteredMinimumPriceCadenceQuarterly  PlanVersionNewParamsReplacePricesPriceNewPlanGroupedWithMeteredMinimumPriceCadence = "quarterly"
	PlanVersionNewParamsReplacePricesPriceNewPlanGroupedWithMeteredMinimumPriceCadenceOneTime    PlanVersionNewParamsReplacePricesPriceNewPlanGroupedWithMeteredMinimumPriceCadence = "one_time"
	PlanVersionNewParamsReplacePricesPriceNewPlanGroupedWithMeteredMinimumPriceCadenceCustom     PlanVersionNewParamsReplacePricesPriceNewPlanGroupedWithMeteredMinimumPriceCadence = "custom"
)

func (r PlanVersionNewParamsReplacePricesPriceNewPlanGroupedWithMeteredMinimumPriceCadence) IsKnown() bool {
	switch r {
	case PlanVersionNewParamsReplacePricesPriceNewPlanGroupedWithMeteredMinimumPriceCadenceAnnual, PlanVersionNewParamsReplacePricesPriceNewPlanGroupedWithMeteredMinimumPriceCadenceSemiAnnual, PlanVersionNewParamsReplacePricesPriceNewPlanGroupedWithMeteredMinimumPriceCadenceMonthly, PlanVersionNewParamsReplacePricesPriceNewPlanGroupedWithMeteredMinimumPriceCadenceQuarterly, PlanVersionNewParamsReplacePricesPriceNewPlanGroupedWithMeteredMinimumPriceCadenceOneTime, PlanVersionNewParamsReplacePricesPriceNewPlanGroupedWithMeteredMinimumPriceCadenceCustom:
		return true
	}
	return false
}

type PlanVersionNewParamsReplacePricesPriceNewPlanGroupedWithMeteredMinimumPriceModelType string

const (
	PlanVersionNewParamsReplacePricesPriceNewPlanGroupedWithMeteredMinimumPriceModelTypeGroupedWithMeteredMinimum PlanVersionNewParamsReplacePricesPriceNewPlanGroupedWithMeteredMinimumPriceModelType = "grouped_with_metered_minimum"
)

func (r PlanVersionNewParamsReplacePricesPriceNewPlanGroupedWithMeteredMinimumPriceModelType) IsKnown() bool {
	switch r {
	case PlanVersionNewParamsReplacePricesPriceNewPlanGroupedWithMeteredMinimumPriceModelTypeGroupedWithMeteredMinimum:
		return true
	}
	return false
}

// For custom cadence: specifies the duration of the billing period in days or
// months.
type PlanVersionNewParamsReplacePricesPriceNewPlanGroupedWithMeteredMinimumPriceBillingCycleConfiguration struct {
	// The duration of the billing period.
	Duration param.Field[int64] `json:"duration,required"`
	// The unit of billing period duration.
	DurationUnit param.Field[PlanVersionNewParamsReplacePricesPriceNewPlanGroupedWithMeteredMinimumPriceBillingCycleConfigurationDurationUnit] `json:"duration_unit,required"`
}

func (r PlanVersionNewParamsReplacePricesPriceNewPlanGroupedWithMeteredMinimumPriceBillingCycleConfiguration) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// The unit of billing period duration.
type PlanVersionNewParamsReplacePricesPriceNewPlanGroupedWithMeteredMinimumPriceBillingCycleConfigurationDurationUnit string

const (
	PlanVersionNewParamsReplacePricesPriceNewPlanGroupedWithMeteredMinimumPriceBillingCycleConfigurationDurationUnitDay   PlanVersionNewParamsReplacePricesPriceNewPlanGroupedWithMeteredMinimumPriceBillingCycleConfigurationDurationUnit = "day"
	PlanVersionNewParamsReplacePricesPriceNewPlanGroupedWithMeteredMinimumPriceBillingCycleConfigurationDurationUnitMonth PlanVersionNewParamsReplacePricesPriceNewPlanGroupedWithMeteredMinimumPriceBillingCycleConfigurationDurationUnit = "month"
)

func (r PlanVersionNewParamsReplacePricesPriceNewPlanGroupedWithMeteredMinimumPriceBillingCycleConfigurationDurationUnit) IsKnown() bool {
	switch r {
	case PlanVersionNewParamsReplacePricesPriceNewPlanGroupedWithMeteredMinimumPriceBillingCycleConfigurationDurationUnitDay, PlanVersionNewParamsReplacePricesPriceNewPlanGroupedWithMeteredMinimumPriceBillingCycleConfigurationDurationUnitMonth:
		return true
	}
	return false
}

// For dimensional price: specifies a price group and dimension values
type PlanVersionNewParamsReplacePricesPriceNewPlanGroupedWithMeteredMinimumPriceDimensionalPriceConfiguration struct {
	// The list of dimension values matching (in order) the dimensions of the price
	// group
	DimensionValues param.Field[[]string] `json:"dimension_values,required"`
	// The id of the dimensional price group to include this price in
	DimensionalPriceGroupID param.Field[string] `json:"dimensional_price_group_id"`
	// The external id of the dimensional price group to include this price in
	ExternalDimensionalPriceGroupID param.Field[string] `json:"external_dimensional_price_group_id"`
}

func (r PlanVersionNewParamsReplacePricesPriceNewPlanGroupedWithMeteredMinimumPriceDimensionalPriceConfiguration) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Within each billing cycle, specifies the cadence at which invoices are produced.
// If unspecified, a single invoice is produced per billing cycle.
type PlanVersionNewParamsReplacePricesPriceNewPlanGroupedWithMeteredMinimumPriceInvoicingCycleConfiguration struct {
	// The duration of the billing period.
	Duration param.Field[int64] `json:"duration,required"`
	// The unit of billing period duration.
	DurationUnit param.Field[PlanVersionNewParamsReplacePricesPriceNewPlanGroupedWithMeteredMinimumPriceInvoicingCycleConfigurationDurationUnit] `json:"duration_unit,required"`
}

func (r PlanVersionNewParamsReplacePricesPriceNewPlanGroupedWithMeteredMinimumPriceInvoicingCycleConfiguration) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// The unit of billing period duration.
type PlanVersionNewParamsReplacePricesPriceNewPlanGroupedWithMeteredMinimumPriceInvoicingCycleConfigurationDurationUnit string

const (
	PlanVersionNewParamsReplacePricesPriceNewPlanGroupedWithMeteredMinimumPriceInvoicingCycleConfigurationDurationUnitDay   PlanVersionNewParamsReplacePricesPriceNewPlanGroupedWithMeteredMinimumPriceInvoicingCycleConfigurationDurationUnit = "day"
	PlanVersionNewParamsReplacePricesPriceNewPlanGroupedWithMeteredMinimumPriceInvoicingCycleConfigurationDurationUnitMonth PlanVersionNewParamsReplacePricesPriceNewPlanGroupedWithMeteredMinimumPriceInvoicingCycleConfigurationDurationUnit = "month"
)

func (r PlanVersionNewParamsReplacePricesPriceNewPlanGroupedWithMeteredMinimumPriceInvoicingCycleConfigurationDurationUnit) IsKnown() bool {
	switch r {
	case PlanVersionNewParamsReplacePricesPriceNewPlanGroupedWithMeteredMinimumPriceInvoicingCycleConfigurationDurationUnitDay, PlanVersionNewParamsReplacePricesPriceNewPlanGroupedWithMeteredMinimumPriceInvoicingCycleConfigurationDurationUnitMonth:
		return true
	}
	return false
}

type PlanVersionNewParamsReplacePricesPriceNewPlanMatrixWithDisplayNamePrice struct {
	// The cadence to bill for this price on.
	Cadence param.Field[PlanVersionNewParamsReplacePricesPriceNewPlanMatrixWithDisplayNamePriceCadence] `json:"cadence,required"`
	// The id of the item the price will be associated with.
	ItemID                      param.Field[string]                                                                           `json:"item_id,required"`
	MatrixWithDisplayNameConfig param.Field[map[string]interface{}]                                                           `json:"matrix_with_display_name_config,required"`
	ModelType                   param.Field[PlanVersionNewParamsReplacePricesPriceNewPlanMatrixWithDisplayNamePriceModelType] `json:"model_type,required"`
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
	BillingCycleConfiguration param.Field[PlanVersionNewParamsReplacePricesPriceNewPlanMatrixWithDisplayNamePriceBillingCycleConfiguration] `json:"billing_cycle_configuration"`
	// The per unit conversion rate of the price currency to the invoicing currency.
	ConversionRate param.Field[float64] `json:"conversion_rate"`
	// An ISO 4217 currency string, or custom pricing unit identifier, in which this
	// price is billed.
	Currency param.Field[string] `json:"currency"`
	// For dimensional price: specifies a price group and dimension values
	DimensionalPriceConfiguration param.Field[PlanVersionNewParamsReplacePricesPriceNewPlanMatrixWithDisplayNamePriceDimensionalPriceConfiguration] `json:"dimensional_price_configuration"`
	// An alias for the price.
	ExternalPriceID param.Field[string] `json:"external_price_id"`
	// If the Price represents a fixed cost, this represents the quantity of units
	// applied.
	FixedPriceQuantity param.Field[float64] `json:"fixed_price_quantity"`
	// The property used to group this price on an invoice
	InvoiceGroupingKey param.Field[string] `json:"invoice_grouping_key"`
	// Within each billing cycle, specifies the cadence at which invoices are produced.
	// If unspecified, a single invoice is produced per billing cycle.
	InvoicingCycleConfiguration param.Field[PlanVersionNewParamsReplacePricesPriceNewPlanMatrixWithDisplayNamePriceInvoicingCycleConfiguration] `json:"invoicing_cycle_configuration"`
	// User-specified key/value pairs for the resource. Individual keys can be removed
	// by setting the value to `null`, and the entire metadata mapping can be cleared
	// by setting `metadata` to `null`.
	Metadata param.Field[map[string]string] `json:"metadata"`
}

func (r PlanVersionNewParamsReplacePricesPriceNewPlanMatrixWithDisplayNamePrice) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r PlanVersionNewParamsReplacePricesPriceNewPlanMatrixWithDisplayNamePrice) implementsPlanVersionNewParamsReplacePricesPriceUnion() {
}

// The cadence to bill for this price on.
type PlanVersionNewParamsReplacePricesPriceNewPlanMatrixWithDisplayNamePriceCadence string

const (
	PlanVersionNewParamsReplacePricesPriceNewPlanMatrixWithDisplayNamePriceCadenceAnnual     PlanVersionNewParamsReplacePricesPriceNewPlanMatrixWithDisplayNamePriceCadence = "annual"
	PlanVersionNewParamsReplacePricesPriceNewPlanMatrixWithDisplayNamePriceCadenceSemiAnnual PlanVersionNewParamsReplacePricesPriceNewPlanMatrixWithDisplayNamePriceCadence = "semi_annual"
	PlanVersionNewParamsReplacePricesPriceNewPlanMatrixWithDisplayNamePriceCadenceMonthly    PlanVersionNewParamsReplacePricesPriceNewPlanMatrixWithDisplayNamePriceCadence = "monthly"
	PlanVersionNewParamsReplacePricesPriceNewPlanMatrixWithDisplayNamePriceCadenceQuarterly  PlanVersionNewParamsReplacePricesPriceNewPlanMatrixWithDisplayNamePriceCadence = "quarterly"
	PlanVersionNewParamsReplacePricesPriceNewPlanMatrixWithDisplayNamePriceCadenceOneTime    PlanVersionNewParamsReplacePricesPriceNewPlanMatrixWithDisplayNamePriceCadence = "one_time"
	PlanVersionNewParamsReplacePricesPriceNewPlanMatrixWithDisplayNamePriceCadenceCustom     PlanVersionNewParamsReplacePricesPriceNewPlanMatrixWithDisplayNamePriceCadence = "custom"
)

func (r PlanVersionNewParamsReplacePricesPriceNewPlanMatrixWithDisplayNamePriceCadence) IsKnown() bool {
	switch r {
	case PlanVersionNewParamsReplacePricesPriceNewPlanMatrixWithDisplayNamePriceCadenceAnnual, PlanVersionNewParamsReplacePricesPriceNewPlanMatrixWithDisplayNamePriceCadenceSemiAnnual, PlanVersionNewParamsReplacePricesPriceNewPlanMatrixWithDisplayNamePriceCadenceMonthly, PlanVersionNewParamsReplacePricesPriceNewPlanMatrixWithDisplayNamePriceCadenceQuarterly, PlanVersionNewParamsReplacePricesPriceNewPlanMatrixWithDisplayNamePriceCadenceOneTime, PlanVersionNewParamsReplacePricesPriceNewPlanMatrixWithDisplayNamePriceCadenceCustom:
		return true
	}
	return false
}

type PlanVersionNewParamsReplacePricesPriceNewPlanMatrixWithDisplayNamePriceModelType string

const (
	PlanVersionNewParamsReplacePricesPriceNewPlanMatrixWithDisplayNamePriceModelTypeMatrixWithDisplayName PlanVersionNewParamsReplacePricesPriceNewPlanMatrixWithDisplayNamePriceModelType = "matrix_with_display_name"
)

func (r PlanVersionNewParamsReplacePricesPriceNewPlanMatrixWithDisplayNamePriceModelType) IsKnown() bool {
	switch r {
	case PlanVersionNewParamsReplacePricesPriceNewPlanMatrixWithDisplayNamePriceModelTypeMatrixWithDisplayName:
		return true
	}
	return false
}

// For custom cadence: specifies the duration of the billing period in days or
// months.
type PlanVersionNewParamsReplacePricesPriceNewPlanMatrixWithDisplayNamePriceBillingCycleConfiguration struct {
	// The duration of the billing period.
	Duration param.Field[int64] `json:"duration,required"`
	// The unit of billing period duration.
	DurationUnit param.Field[PlanVersionNewParamsReplacePricesPriceNewPlanMatrixWithDisplayNamePriceBillingCycleConfigurationDurationUnit] `json:"duration_unit,required"`
}

func (r PlanVersionNewParamsReplacePricesPriceNewPlanMatrixWithDisplayNamePriceBillingCycleConfiguration) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// The unit of billing period duration.
type PlanVersionNewParamsReplacePricesPriceNewPlanMatrixWithDisplayNamePriceBillingCycleConfigurationDurationUnit string

const (
	PlanVersionNewParamsReplacePricesPriceNewPlanMatrixWithDisplayNamePriceBillingCycleConfigurationDurationUnitDay   PlanVersionNewParamsReplacePricesPriceNewPlanMatrixWithDisplayNamePriceBillingCycleConfigurationDurationUnit = "day"
	PlanVersionNewParamsReplacePricesPriceNewPlanMatrixWithDisplayNamePriceBillingCycleConfigurationDurationUnitMonth PlanVersionNewParamsReplacePricesPriceNewPlanMatrixWithDisplayNamePriceBillingCycleConfigurationDurationUnit = "month"
)

func (r PlanVersionNewParamsReplacePricesPriceNewPlanMatrixWithDisplayNamePriceBillingCycleConfigurationDurationUnit) IsKnown() bool {
	switch r {
	case PlanVersionNewParamsReplacePricesPriceNewPlanMatrixWithDisplayNamePriceBillingCycleConfigurationDurationUnitDay, PlanVersionNewParamsReplacePricesPriceNewPlanMatrixWithDisplayNamePriceBillingCycleConfigurationDurationUnitMonth:
		return true
	}
	return false
}

// For dimensional price: specifies a price group and dimension values
type PlanVersionNewParamsReplacePricesPriceNewPlanMatrixWithDisplayNamePriceDimensionalPriceConfiguration struct {
	// The list of dimension values matching (in order) the dimensions of the price
	// group
	DimensionValues param.Field[[]string] `json:"dimension_values,required"`
	// The id of the dimensional price group to include this price in
	DimensionalPriceGroupID param.Field[string] `json:"dimensional_price_group_id"`
	// The external id of the dimensional price group to include this price in
	ExternalDimensionalPriceGroupID param.Field[string] `json:"external_dimensional_price_group_id"`
}

func (r PlanVersionNewParamsReplacePricesPriceNewPlanMatrixWithDisplayNamePriceDimensionalPriceConfiguration) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Within each billing cycle, specifies the cadence at which invoices are produced.
// If unspecified, a single invoice is produced per billing cycle.
type PlanVersionNewParamsReplacePricesPriceNewPlanMatrixWithDisplayNamePriceInvoicingCycleConfiguration struct {
	// The duration of the billing period.
	Duration param.Field[int64] `json:"duration,required"`
	// The unit of billing period duration.
	DurationUnit param.Field[PlanVersionNewParamsReplacePricesPriceNewPlanMatrixWithDisplayNamePriceInvoicingCycleConfigurationDurationUnit] `json:"duration_unit,required"`
}

func (r PlanVersionNewParamsReplacePricesPriceNewPlanMatrixWithDisplayNamePriceInvoicingCycleConfiguration) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// The unit of billing period duration.
type PlanVersionNewParamsReplacePricesPriceNewPlanMatrixWithDisplayNamePriceInvoicingCycleConfigurationDurationUnit string

const (
	PlanVersionNewParamsReplacePricesPriceNewPlanMatrixWithDisplayNamePriceInvoicingCycleConfigurationDurationUnitDay   PlanVersionNewParamsReplacePricesPriceNewPlanMatrixWithDisplayNamePriceInvoicingCycleConfigurationDurationUnit = "day"
	PlanVersionNewParamsReplacePricesPriceNewPlanMatrixWithDisplayNamePriceInvoicingCycleConfigurationDurationUnitMonth PlanVersionNewParamsReplacePricesPriceNewPlanMatrixWithDisplayNamePriceInvoicingCycleConfigurationDurationUnit = "month"
)

func (r PlanVersionNewParamsReplacePricesPriceNewPlanMatrixWithDisplayNamePriceInvoicingCycleConfigurationDurationUnit) IsKnown() bool {
	switch r {
	case PlanVersionNewParamsReplacePricesPriceNewPlanMatrixWithDisplayNamePriceInvoicingCycleConfigurationDurationUnitDay, PlanVersionNewParamsReplacePricesPriceNewPlanMatrixWithDisplayNamePriceInvoicingCycleConfigurationDurationUnitMonth:
		return true
	}
	return false
}

type PlanVersionNewParamsReplacePricesPriceNewPlanBulkWithProrationPrice struct {
	BulkWithProrationConfig param.Field[map[string]interface{}] `json:"bulk_with_proration_config,required"`
	// The cadence to bill for this price on.
	Cadence param.Field[PlanVersionNewParamsReplacePricesPriceNewPlanBulkWithProrationPriceCadence] `json:"cadence,required"`
	// The id of the item the price will be associated with.
	ItemID    param.Field[string]                                                                       `json:"item_id,required"`
	ModelType param.Field[PlanVersionNewParamsReplacePricesPriceNewPlanBulkWithProrationPriceModelType] `json:"model_type,required"`
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
	BillingCycleConfiguration param.Field[PlanVersionNewParamsReplacePricesPriceNewPlanBulkWithProrationPriceBillingCycleConfiguration] `json:"billing_cycle_configuration"`
	// The per unit conversion rate of the price currency to the invoicing currency.
	ConversionRate param.Field[float64] `json:"conversion_rate"`
	// An ISO 4217 currency string, or custom pricing unit identifier, in which this
	// price is billed.
	Currency param.Field[string] `json:"currency"`
	// For dimensional price: specifies a price group and dimension values
	DimensionalPriceConfiguration param.Field[PlanVersionNewParamsReplacePricesPriceNewPlanBulkWithProrationPriceDimensionalPriceConfiguration] `json:"dimensional_price_configuration"`
	// An alias for the price.
	ExternalPriceID param.Field[string] `json:"external_price_id"`
	// If the Price represents a fixed cost, this represents the quantity of units
	// applied.
	FixedPriceQuantity param.Field[float64] `json:"fixed_price_quantity"`
	// The property used to group this price on an invoice
	InvoiceGroupingKey param.Field[string] `json:"invoice_grouping_key"`
	// Within each billing cycle, specifies the cadence at which invoices are produced.
	// If unspecified, a single invoice is produced per billing cycle.
	InvoicingCycleConfiguration param.Field[PlanVersionNewParamsReplacePricesPriceNewPlanBulkWithProrationPriceInvoicingCycleConfiguration] `json:"invoicing_cycle_configuration"`
	// User-specified key/value pairs for the resource. Individual keys can be removed
	// by setting the value to `null`, and the entire metadata mapping can be cleared
	// by setting `metadata` to `null`.
	Metadata param.Field[map[string]string] `json:"metadata"`
}

func (r PlanVersionNewParamsReplacePricesPriceNewPlanBulkWithProrationPrice) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r PlanVersionNewParamsReplacePricesPriceNewPlanBulkWithProrationPrice) implementsPlanVersionNewParamsReplacePricesPriceUnion() {
}

// The cadence to bill for this price on.
type PlanVersionNewParamsReplacePricesPriceNewPlanBulkWithProrationPriceCadence string

const (
	PlanVersionNewParamsReplacePricesPriceNewPlanBulkWithProrationPriceCadenceAnnual     PlanVersionNewParamsReplacePricesPriceNewPlanBulkWithProrationPriceCadence = "annual"
	PlanVersionNewParamsReplacePricesPriceNewPlanBulkWithProrationPriceCadenceSemiAnnual PlanVersionNewParamsReplacePricesPriceNewPlanBulkWithProrationPriceCadence = "semi_annual"
	PlanVersionNewParamsReplacePricesPriceNewPlanBulkWithProrationPriceCadenceMonthly    PlanVersionNewParamsReplacePricesPriceNewPlanBulkWithProrationPriceCadence = "monthly"
	PlanVersionNewParamsReplacePricesPriceNewPlanBulkWithProrationPriceCadenceQuarterly  PlanVersionNewParamsReplacePricesPriceNewPlanBulkWithProrationPriceCadence = "quarterly"
	PlanVersionNewParamsReplacePricesPriceNewPlanBulkWithProrationPriceCadenceOneTime    PlanVersionNewParamsReplacePricesPriceNewPlanBulkWithProrationPriceCadence = "one_time"
	PlanVersionNewParamsReplacePricesPriceNewPlanBulkWithProrationPriceCadenceCustom     PlanVersionNewParamsReplacePricesPriceNewPlanBulkWithProrationPriceCadence = "custom"
)

func (r PlanVersionNewParamsReplacePricesPriceNewPlanBulkWithProrationPriceCadence) IsKnown() bool {
	switch r {
	case PlanVersionNewParamsReplacePricesPriceNewPlanBulkWithProrationPriceCadenceAnnual, PlanVersionNewParamsReplacePricesPriceNewPlanBulkWithProrationPriceCadenceSemiAnnual, PlanVersionNewParamsReplacePricesPriceNewPlanBulkWithProrationPriceCadenceMonthly, PlanVersionNewParamsReplacePricesPriceNewPlanBulkWithProrationPriceCadenceQuarterly, PlanVersionNewParamsReplacePricesPriceNewPlanBulkWithProrationPriceCadenceOneTime, PlanVersionNewParamsReplacePricesPriceNewPlanBulkWithProrationPriceCadenceCustom:
		return true
	}
	return false
}

type PlanVersionNewParamsReplacePricesPriceNewPlanBulkWithProrationPriceModelType string

const (
	PlanVersionNewParamsReplacePricesPriceNewPlanBulkWithProrationPriceModelTypeBulkWithProration PlanVersionNewParamsReplacePricesPriceNewPlanBulkWithProrationPriceModelType = "bulk_with_proration"
)

func (r PlanVersionNewParamsReplacePricesPriceNewPlanBulkWithProrationPriceModelType) IsKnown() bool {
	switch r {
	case PlanVersionNewParamsReplacePricesPriceNewPlanBulkWithProrationPriceModelTypeBulkWithProration:
		return true
	}
	return false
}

// For custom cadence: specifies the duration of the billing period in days or
// months.
type PlanVersionNewParamsReplacePricesPriceNewPlanBulkWithProrationPriceBillingCycleConfiguration struct {
	// The duration of the billing period.
	Duration param.Field[int64] `json:"duration,required"`
	// The unit of billing period duration.
	DurationUnit param.Field[PlanVersionNewParamsReplacePricesPriceNewPlanBulkWithProrationPriceBillingCycleConfigurationDurationUnit] `json:"duration_unit,required"`
}

func (r PlanVersionNewParamsReplacePricesPriceNewPlanBulkWithProrationPriceBillingCycleConfiguration) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// The unit of billing period duration.
type PlanVersionNewParamsReplacePricesPriceNewPlanBulkWithProrationPriceBillingCycleConfigurationDurationUnit string

const (
	PlanVersionNewParamsReplacePricesPriceNewPlanBulkWithProrationPriceBillingCycleConfigurationDurationUnitDay   PlanVersionNewParamsReplacePricesPriceNewPlanBulkWithProrationPriceBillingCycleConfigurationDurationUnit = "day"
	PlanVersionNewParamsReplacePricesPriceNewPlanBulkWithProrationPriceBillingCycleConfigurationDurationUnitMonth PlanVersionNewParamsReplacePricesPriceNewPlanBulkWithProrationPriceBillingCycleConfigurationDurationUnit = "month"
)

func (r PlanVersionNewParamsReplacePricesPriceNewPlanBulkWithProrationPriceBillingCycleConfigurationDurationUnit) IsKnown() bool {
	switch r {
	case PlanVersionNewParamsReplacePricesPriceNewPlanBulkWithProrationPriceBillingCycleConfigurationDurationUnitDay, PlanVersionNewParamsReplacePricesPriceNewPlanBulkWithProrationPriceBillingCycleConfigurationDurationUnitMonth:
		return true
	}
	return false
}

// For dimensional price: specifies a price group and dimension values
type PlanVersionNewParamsReplacePricesPriceNewPlanBulkWithProrationPriceDimensionalPriceConfiguration struct {
	// The list of dimension values matching (in order) the dimensions of the price
	// group
	DimensionValues param.Field[[]string] `json:"dimension_values,required"`
	// The id of the dimensional price group to include this price in
	DimensionalPriceGroupID param.Field[string] `json:"dimensional_price_group_id"`
	// The external id of the dimensional price group to include this price in
	ExternalDimensionalPriceGroupID param.Field[string] `json:"external_dimensional_price_group_id"`
}

func (r PlanVersionNewParamsReplacePricesPriceNewPlanBulkWithProrationPriceDimensionalPriceConfiguration) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Within each billing cycle, specifies the cadence at which invoices are produced.
// If unspecified, a single invoice is produced per billing cycle.
type PlanVersionNewParamsReplacePricesPriceNewPlanBulkWithProrationPriceInvoicingCycleConfiguration struct {
	// The duration of the billing period.
	Duration param.Field[int64] `json:"duration,required"`
	// The unit of billing period duration.
	DurationUnit param.Field[PlanVersionNewParamsReplacePricesPriceNewPlanBulkWithProrationPriceInvoicingCycleConfigurationDurationUnit] `json:"duration_unit,required"`
}

func (r PlanVersionNewParamsReplacePricesPriceNewPlanBulkWithProrationPriceInvoicingCycleConfiguration) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// The unit of billing period duration.
type PlanVersionNewParamsReplacePricesPriceNewPlanBulkWithProrationPriceInvoicingCycleConfigurationDurationUnit string

const (
	PlanVersionNewParamsReplacePricesPriceNewPlanBulkWithProrationPriceInvoicingCycleConfigurationDurationUnitDay   PlanVersionNewParamsReplacePricesPriceNewPlanBulkWithProrationPriceInvoicingCycleConfigurationDurationUnit = "day"
	PlanVersionNewParamsReplacePricesPriceNewPlanBulkWithProrationPriceInvoicingCycleConfigurationDurationUnitMonth PlanVersionNewParamsReplacePricesPriceNewPlanBulkWithProrationPriceInvoicingCycleConfigurationDurationUnit = "month"
)

func (r PlanVersionNewParamsReplacePricesPriceNewPlanBulkWithProrationPriceInvoicingCycleConfigurationDurationUnit) IsKnown() bool {
	switch r {
	case PlanVersionNewParamsReplacePricesPriceNewPlanBulkWithProrationPriceInvoicingCycleConfigurationDurationUnitDay, PlanVersionNewParamsReplacePricesPriceNewPlanBulkWithProrationPriceInvoicingCycleConfigurationDurationUnitMonth:
		return true
	}
	return false
}

type PlanVersionNewParamsReplacePricesPriceNewPlanGroupedTieredPackagePrice struct {
	// The cadence to bill for this price on.
	Cadence                    param.Field[PlanVersionNewParamsReplacePricesPriceNewPlanGroupedTieredPackagePriceCadence] `json:"cadence,required"`
	GroupedTieredPackageConfig param.Field[map[string]interface{}]                                                        `json:"grouped_tiered_package_config,required"`
	// The id of the item the price will be associated with.
	ItemID    param.Field[string]                                                                          `json:"item_id,required"`
	ModelType param.Field[PlanVersionNewParamsReplacePricesPriceNewPlanGroupedTieredPackagePriceModelType] `json:"model_type,required"`
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
	BillingCycleConfiguration param.Field[PlanVersionNewParamsReplacePricesPriceNewPlanGroupedTieredPackagePriceBillingCycleConfiguration] `json:"billing_cycle_configuration"`
	// The per unit conversion rate of the price currency to the invoicing currency.
	ConversionRate param.Field[float64] `json:"conversion_rate"`
	// An ISO 4217 currency string, or custom pricing unit identifier, in which this
	// price is billed.
	Currency param.Field[string] `json:"currency"`
	// For dimensional price: specifies a price group and dimension values
	DimensionalPriceConfiguration param.Field[PlanVersionNewParamsReplacePricesPriceNewPlanGroupedTieredPackagePriceDimensionalPriceConfiguration] `json:"dimensional_price_configuration"`
	// An alias for the price.
	ExternalPriceID param.Field[string] `json:"external_price_id"`
	// If the Price represents a fixed cost, this represents the quantity of units
	// applied.
	FixedPriceQuantity param.Field[float64] `json:"fixed_price_quantity"`
	// The property used to group this price on an invoice
	InvoiceGroupingKey param.Field[string] `json:"invoice_grouping_key"`
	// Within each billing cycle, specifies the cadence at which invoices are produced.
	// If unspecified, a single invoice is produced per billing cycle.
	InvoicingCycleConfiguration param.Field[PlanVersionNewParamsReplacePricesPriceNewPlanGroupedTieredPackagePriceInvoicingCycleConfiguration] `json:"invoicing_cycle_configuration"`
	// User-specified key/value pairs for the resource. Individual keys can be removed
	// by setting the value to `null`, and the entire metadata mapping can be cleared
	// by setting `metadata` to `null`.
	Metadata param.Field[map[string]string] `json:"metadata"`
}

func (r PlanVersionNewParamsReplacePricesPriceNewPlanGroupedTieredPackagePrice) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r PlanVersionNewParamsReplacePricesPriceNewPlanGroupedTieredPackagePrice) implementsPlanVersionNewParamsReplacePricesPriceUnion() {
}

// The cadence to bill for this price on.
type PlanVersionNewParamsReplacePricesPriceNewPlanGroupedTieredPackagePriceCadence string

const (
	PlanVersionNewParamsReplacePricesPriceNewPlanGroupedTieredPackagePriceCadenceAnnual     PlanVersionNewParamsReplacePricesPriceNewPlanGroupedTieredPackagePriceCadence = "annual"
	PlanVersionNewParamsReplacePricesPriceNewPlanGroupedTieredPackagePriceCadenceSemiAnnual PlanVersionNewParamsReplacePricesPriceNewPlanGroupedTieredPackagePriceCadence = "semi_annual"
	PlanVersionNewParamsReplacePricesPriceNewPlanGroupedTieredPackagePriceCadenceMonthly    PlanVersionNewParamsReplacePricesPriceNewPlanGroupedTieredPackagePriceCadence = "monthly"
	PlanVersionNewParamsReplacePricesPriceNewPlanGroupedTieredPackagePriceCadenceQuarterly  PlanVersionNewParamsReplacePricesPriceNewPlanGroupedTieredPackagePriceCadence = "quarterly"
	PlanVersionNewParamsReplacePricesPriceNewPlanGroupedTieredPackagePriceCadenceOneTime    PlanVersionNewParamsReplacePricesPriceNewPlanGroupedTieredPackagePriceCadence = "one_time"
	PlanVersionNewParamsReplacePricesPriceNewPlanGroupedTieredPackagePriceCadenceCustom     PlanVersionNewParamsReplacePricesPriceNewPlanGroupedTieredPackagePriceCadence = "custom"
)

func (r PlanVersionNewParamsReplacePricesPriceNewPlanGroupedTieredPackagePriceCadence) IsKnown() bool {
	switch r {
	case PlanVersionNewParamsReplacePricesPriceNewPlanGroupedTieredPackagePriceCadenceAnnual, PlanVersionNewParamsReplacePricesPriceNewPlanGroupedTieredPackagePriceCadenceSemiAnnual, PlanVersionNewParamsReplacePricesPriceNewPlanGroupedTieredPackagePriceCadenceMonthly, PlanVersionNewParamsReplacePricesPriceNewPlanGroupedTieredPackagePriceCadenceQuarterly, PlanVersionNewParamsReplacePricesPriceNewPlanGroupedTieredPackagePriceCadenceOneTime, PlanVersionNewParamsReplacePricesPriceNewPlanGroupedTieredPackagePriceCadenceCustom:
		return true
	}
	return false
}

type PlanVersionNewParamsReplacePricesPriceNewPlanGroupedTieredPackagePriceModelType string

const (
	PlanVersionNewParamsReplacePricesPriceNewPlanGroupedTieredPackagePriceModelTypeGroupedTieredPackage PlanVersionNewParamsReplacePricesPriceNewPlanGroupedTieredPackagePriceModelType = "grouped_tiered_package"
)

func (r PlanVersionNewParamsReplacePricesPriceNewPlanGroupedTieredPackagePriceModelType) IsKnown() bool {
	switch r {
	case PlanVersionNewParamsReplacePricesPriceNewPlanGroupedTieredPackagePriceModelTypeGroupedTieredPackage:
		return true
	}
	return false
}

// For custom cadence: specifies the duration of the billing period in days or
// months.
type PlanVersionNewParamsReplacePricesPriceNewPlanGroupedTieredPackagePriceBillingCycleConfiguration struct {
	// The duration of the billing period.
	Duration param.Field[int64] `json:"duration,required"`
	// The unit of billing period duration.
	DurationUnit param.Field[PlanVersionNewParamsReplacePricesPriceNewPlanGroupedTieredPackagePriceBillingCycleConfigurationDurationUnit] `json:"duration_unit,required"`
}

func (r PlanVersionNewParamsReplacePricesPriceNewPlanGroupedTieredPackagePriceBillingCycleConfiguration) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// The unit of billing period duration.
type PlanVersionNewParamsReplacePricesPriceNewPlanGroupedTieredPackagePriceBillingCycleConfigurationDurationUnit string

const (
	PlanVersionNewParamsReplacePricesPriceNewPlanGroupedTieredPackagePriceBillingCycleConfigurationDurationUnitDay   PlanVersionNewParamsReplacePricesPriceNewPlanGroupedTieredPackagePriceBillingCycleConfigurationDurationUnit = "day"
	PlanVersionNewParamsReplacePricesPriceNewPlanGroupedTieredPackagePriceBillingCycleConfigurationDurationUnitMonth PlanVersionNewParamsReplacePricesPriceNewPlanGroupedTieredPackagePriceBillingCycleConfigurationDurationUnit = "month"
)

func (r PlanVersionNewParamsReplacePricesPriceNewPlanGroupedTieredPackagePriceBillingCycleConfigurationDurationUnit) IsKnown() bool {
	switch r {
	case PlanVersionNewParamsReplacePricesPriceNewPlanGroupedTieredPackagePriceBillingCycleConfigurationDurationUnitDay, PlanVersionNewParamsReplacePricesPriceNewPlanGroupedTieredPackagePriceBillingCycleConfigurationDurationUnitMonth:
		return true
	}
	return false
}

// For dimensional price: specifies a price group and dimension values
type PlanVersionNewParamsReplacePricesPriceNewPlanGroupedTieredPackagePriceDimensionalPriceConfiguration struct {
	// The list of dimension values matching (in order) the dimensions of the price
	// group
	DimensionValues param.Field[[]string] `json:"dimension_values,required"`
	// The id of the dimensional price group to include this price in
	DimensionalPriceGroupID param.Field[string] `json:"dimensional_price_group_id"`
	// The external id of the dimensional price group to include this price in
	ExternalDimensionalPriceGroupID param.Field[string] `json:"external_dimensional_price_group_id"`
}

func (r PlanVersionNewParamsReplacePricesPriceNewPlanGroupedTieredPackagePriceDimensionalPriceConfiguration) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Within each billing cycle, specifies the cadence at which invoices are produced.
// If unspecified, a single invoice is produced per billing cycle.
type PlanVersionNewParamsReplacePricesPriceNewPlanGroupedTieredPackagePriceInvoicingCycleConfiguration struct {
	// The duration of the billing period.
	Duration param.Field[int64] `json:"duration,required"`
	// The unit of billing period duration.
	DurationUnit param.Field[PlanVersionNewParamsReplacePricesPriceNewPlanGroupedTieredPackagePriceInvoicingCycleConfigurationDurationUnit] `json:"duration_unit,required"`
}

func (r PlanVersionNewParamsReplacePricesPriceNewPlanGroupedTieredPackagePriceInvoicingCycleConfiguration) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// The unit of billing period duration.
type PlanVersionNewParamsReplacePricesPriceNewPlanGroupedTieredPackagePriceInvoicingCycleConfigurationDurationUnit string

const (
	PlanVersionNewParamsReplacePricesPriceNewPlanGroupedTieredPackagePriceInvoicingCycleConfigurationDurationUnitDay   PlanVersionNewParamsReplacePricesPriceNewPlanGroupedTieredPackagePriceInvoicingCycleConfigurationDurationUnit = "day"
	PlanVersionNewParamsReplacePricesPriceNewPlanGroupedTieredPackagePriceInvoicingCycleConfigurationDurationUnitMonth PlanVersionNewParamsReplacePricesPriceNewPlanGroupedTieredPackagePriceInvoicingCycleConfigurationDurationUnit = "month"
)

func (r PlanVersionNewParamsReplacePricesPriceNewPlanGroupedTieredPackagePriceInvoicingCycleConfigurationDurationUnit) IsKnown() bool {
	switch r {
	case PlanVersionNewParamsReplacePricesPriceNewPlanGroupedTieredPackagePriceInvoicingCycleConfigurationDurationUnitDay, PlanVersionNewParamsReplacePricesPriceNewPlanGroupedTieredPackagePriceInvoicingCycleConfigurationDurationUnitMonth:
		return true
	}
	return false
}

type PlanVersionNewParamsReplacePricesPriceNewPlanMaxGroupTieredPackagePrice struct {
	// The cadence to bill for this price on.
	Cadence param.Field[PlanVersionNewParamsReplacePricesPriceNewPlanMaxGroupTieredPackagePriceCadence] `json:"cadence,required"`
	// The id of the item the price will be associated with.
	ItemID                      param.Field[string]                                                                           `json:"item_id,required"`
	MaxGroupTieredPackageConfig param.Field[map[string]interface{}]                                                           `json:"max_group_tiered_package_config,required"`
	ModelType                   param.Field[PlanVersionNewParamsReplacePricesPriceNewPlanMaxGroupTieredPackagePriceModelType] `json:"model_type,required"`
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
	BillingCycleConfiguration param.Field[PlanVersionNewParamsReplacePricesPriceNewPlanMaxGroupTieredPackagePriceBillingCycleConfiguration] `json:"billing_cycle_configuration"`
	// The per unit conversion rate of the price currency to the invoicing currency.
	ConversionRate param.Field[float64] `json:"conversion_rate"`
	// An ISO 4217 currency string, or custom pricing unit identifier, in which this
	// price is billed.
	Currency param.Field[string] `json:"currency"`
	// For dimensional price: specifies a price group and dimension values
	DimensionalPriceConfiguration param.Field[PlanVersionNewParamsReplacePricesPriceNewPlanMaxGroupTieredPackagePriceDimensionalPriceConfiguration] `json:"dimensional_price_configuration"`
	// An alias for the price.
	ExternalPriceID param.Field[string] `json:"external_price_id"`
	// If the Price represents a fixed cost, this represents the quantity of units
	// applied.
	FixedPriceQuantity param.Field[float64] `json:"fixed_price_quantity"`
	// The property used to group this price on an invoice
	InvoiceGroupingKey param.Field[string] `json:"invoice_grouping_key"`
	// Within each billing cycle, specifies the cadence at which invoices are produced.
	// If unspecified, a single invoice is produced per billing cycle.
	InvoicingCycleConfiguration param.Field[PlanVersionNewParamsReplacePricesPriceNewPlanMaxGroupTieredPackagePriceInvoicingCycleConfiguration] `json:"invoicing_cycle_configuration"`
	// User-specified key/value pairs for the resource. Individual keys can be removed
	// by setting the value to `null`, and the entire metadata mapping can be cleared
	// by setting `metadata` to `null`.
	Metadata param.Field[map[string]string] `json:"metadata"`
}

func (r PlanVersionNewParamsReplacePricesPriceNewPlanMaxGroupTieredPackagePrice) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r PlanVersionNewParamsReplacePricesPriceNewPlanMaxGroupTieredPackagePrice) implementsPlanVersionNewParamsReplacePricesPriceUnion() {
}

// The cadence to bill for this price on.
type PlanVersionNewParamsReplacePricesPriceNewPlanMaxGroupTieredPackagePriceCadence string

const (
	PlanVersionNewParamsReplacePricesPriceNewPlanMaxGroupTieredPackagePriceCadenceAnnual     PlanVersionNewParamsReplacePricesPriceNewPlanMaxGroupTieredPackagePriceCadence = "annual"
	PlanVersionNewParamsReplacePricesPriceNewPlanMaxGroupTieredPackagePriceCadenceSemiAnnual PlanVersionNewParamsReplacePricesPriceNewPlanMaxGroupTieredPackagePriceCadence = "semi_annual"
	PlanVersionNewParamsReplacePricesPriceNewPlanMaxGroupTieredPackagePriceCadenceMonthly    PlanVersionNewParamsReplacePricesPriceNewPlanMaxGroupTieredPackagePriceCadence = "monthly"
	PlanVersionNewParamsReplacePricesPriceNewPlanMaxGroupTieredPackagePriceCadenceQuarterly  PlanVersionNewParamsReplacePricesPriceNewPlanMaxGroupTieredPackagePriceCadence = "quarterly"
	PlanVersionNewParamsReplacePricesPriceNewPlanMaxGroupTieredPackagePriceCadenceOneTime    PlanVersionNewParamsReplacePricesPriceNewPlanMaxGroupTieredPackagePriceCadence = "one_time"
	PlanVersionNewParamsReplacePricesPriceNewPlanMaxGroupTieredPackagePriceCadenceCustom     PlanVersionNewParamsReplacePricesPriceNewPlanMaxGroupTieredPackagePriceCadence = "custom"
)

func (r PlanVersionNewParamsReplacePricesPriceNewPlanMaxGroupTieredPackagePriceCadence) IsKnown() bool {
	switch r {
	case PlanVersionNewParamsReplacePricesPriceNewPlanMaxGroupTieredPackagePriceCadenceAnnual, PlanVersionNewParamsReplacePricesPriceNewPlanMaxGroupTieredPackagePriceCadenceSemiAnnual, PlanVersionNewParamsReplacePricesPriceNewPlanMaxGroupTieredPackagePriceCadenceMonthly, PlanVersionNewParamsReplacePricesPriceNewPlanMaxGroupTieredPackagePriceCadenceQuarterly, PlanVersionNewParamsReplacePricesPriceNewPlanMaxGroupTieredPackagePriceCadenceOneTime, PlanVersionNewParamsReplacePricesPriceNewPlanMaxGroupTieredPackagePriceCadenceCustom:
		return true
	}
	return false
}

type PlanVersionNewParamsReplacePricesPriceNewPlanMaxGroupTieredPackagePriceModelType string

const (
	PlanVersionNewParamsReplacePricesPriceNewPlanMaxGroupTieredPackagePriceModelTypeMaxGroupTieredPackage PlanVersionNewParamsReplacePricesPriceNewPlanMaxGroupTieredPackagePriceModelType = "max_group_tiered_package"
)

func (r PlanVersionNewParamsReplacePricesPriceNewPlanMaxGroupTieredPackagePriceModelType) IsKnown() bool {
	switch r {
	case PlanVersionNewParamsReplacePricesPriceNewPlanMaxGroupTieredPackagePriceModelTypeMaxGroupTieredPackage:
		return true
	}
	return false
}

// For custom cadence: specifies the duration of the billing period in days or
// months.
type PlanVersionNewParamsReplacePricesPriceNewPlanMaxGroupTieredPackagePriceBillingCycleConfiguration struct {
	// The duration of the billing period.
	Duration param.Field[int64] `json:"duration,required"`
	// The unit of billing period duration.
	DurationUnit param.Field[PlanVersionNewParamsReplacePricesPriceNewPlanMaxGroupTieredPackagePriceBillingCycleConfigurationDurationUnit] `json:"duration_unit,required"`
}

func (r PlanVersionNewParamsReplacePricesPriceNewPlanMaxGroupTieredPackagePriceBillingCycleConfiguration) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// The unit of billing period duration.
type PlanVersionNewParamsReplacePricesPriceNewPlanMaxGroupTieredPackagePriceBillingCycleConfigurationDurationUnit string

const (
	PlanVersionNewParamsReplacePricesPriceNewPlanMaxGroupTieredPackagePriceBillingCycleConfigurationDurationUnitDay   PlanVersionNewParamsReplacePricesPriceNewPlanMaxGroupTieredPackagePriceBillingCycleConfigurationDurationUnit = "day"
	PlanVersionNewParamsReplacePricesPriceNewPlanMaxGroupTieredPackagePriceBillingCycleConfigurationDurationUnitMonth PlanVersionNewParamsReplacePricesPriceNewPlanMaxGroupTieredPackagePriceBillingCycleConfigurationDurationUnit = "month"
)

func (r PlanVersionNewParamsReplacePricesPriceNewPlanMaxGroupTieredPackagePriceBillingCycleConfigurationDurationUnit) IsKnown() bool {
	switch r {
	case PlanVersionNewParamsReplacePricesPriceNewPlanMaxGroupTieredPackagePriceBillingCycleConfigurationDurationUnitDay, PlanVersionNewParamsReplacePricesPriceNewPlanMaxGroupTieredPackagePriceBillingCycleConfigurationDurationUnitMonth:
		return true
	}
	return false
}

// For dimensional price: specifies a price group and dimension values
type PlanVersionNewParamsReplacePricesPriceNewPlanMaxGroupTieredPackagePriceDimensionalPriceConfiguration struct {
	// The list of dimension values matching (in order) the dimensions of the price
	// group
	DimensionValues param.Field[[]string] `json:"dimension_values,required"`
	// The id of the dimensional price group to include this price in
	DimensionalPriceGroupID param.Field[string] `json:"dimensional_price_group_id"`
	// The external id of the dimensional price group to include this price in
	ExternalDimensionalPriceGroupID param.Field[string] `json:"external_dimensional_price_group_id"`
}

func (r PlanVersionNewParamsReplacePricesPriceNewPlanMaxGroupTieredPackagePriceDimensionalPriceConfiguration) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Within each billing cycle, specifies the cadence at which invoices are produced.
// If unspecified, a single invoice is produced per billing cycle.
type PlanVersionNewParamsReplacePricesPriceNewPlanMaxGroupTieredPackagePriceInvoicingCycleConfiguration struct {
	// The duration of the billing period.
	Duration param.Field[int64] `json:"duration,required"`
	// The unit of billing period duration.
	DurationUnit param.Field[PlanVersionNewParamsReplacePricesPriceNewPlanMaxGroupTieredPackagePriceInvoicingCycleConfigurationDurationUnit] `json:"duration_unit,required"`
}

func (r PlanVersionNewParamsReplacePricesPriceNewPlanMaxGroupTieredPackagePriceInvoicingCycleConfiguration) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// The unit of billing period duration.
type PlanVersionNewParamsReplacePricesPriceNewPlanMaxGroupTieredPackagePriceInvoicingCycleConfigurationDurationUnit string

const (
	PlanVersionNewParamsReplacePricesPriceNewPlanMaxGroupTieredPackagePriceInvoicingCycleConfigurationDurationUnitDay   PlanVersionNewParamsReplacePricesPriceNewPlanMaxGroupTieredPackagePriceInvoicingCycleConfigurationDurationUnit = "day"
	PlanVersionNewParamsReplacePricesPriceNewPlanMaxGroupTieredPackagePriceInvoicingCycleConfigurationDurationUnitMonth PlanVersionNewParamsReplacePricesPriceNewPlanMaxGroupTieredPackagePriceInvoicingCycleConfigurationDurationUnit = "month"
)

func (r PlanVersionNewParamsReplacePricesPriceNewPlanMaxGroupTieredPackagePriceInvoicingCycleConfigurationDurationUnit) IsKnown() bool {
	switch r {
	case PlanVersionNewParamsReplacePricesPriceNewPlanMaxGroupTieredPackagePriceInvoicingCycleConfigurationDurationUnitDay, PlanVersionNewParamsReplacePricesPriceNewPlanMaxGroupTieredPackagePriceInvoicingCycleConfigurationDurationUnitMonth:
		return true
	}
	return false
}

type PlanVersionNewParamsReplacePricesPriceNewPlanScalableMatrixWithUnitPricingPrice struct {
	// The cadence to bill for this price on.
	Cadence param.Field[PlanVersionNewParamsReplacePricesPriceNewPlanScalableMatrixWithUnitPricingPriceCadence] `json:"cadence,required"`
	// The id of the item the price will be associated with.
	ItemID    param.Field[string]                                                                                   `json:"item_id,required"`
	ModelType param.Field[PlanVersionNewParamsReplacePricesPriceNewPlanScalableMatrixWithUnitPricingPriceModelType] `json:"model_type,required"`
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
	BillingCycleConfiguration param.Field[PlanVersionNewParamsReplacePricesPriceNewPlanScalableMatrixWithUnitPricingPriceBillingCycleConfiguration] `json:"billing_cycle_configuration"`
	// The per unit conversion rate of the price currency to the invoicing currency.
	ConversionRate param.Field[float64] `json:"conversion_rate"`
	// An ISO 4217 currency string, or custom pricing unit identifier, in which this
	// price is billed.
	Currency param.Field[string] `json:"currency"`
	// For dimensional price: specifies a price group and dimension values
	DimensionalPriceConfiguration param.Field[PlanVersionNewParamsReplacePricesPriceNewPlanScalableMatrixWithUnitPricingPriceDimensionalPriceConfiguration] `json:"dimensional_price_configuration"`
	// An alias for the price.
	ExternalPriceID param.Field[string] `json:"external_price_id"`
	// If the Price represents a fixed cost, this represents the quantity of units
	// applied.
	FixedPriceQuantity param.Field[float64] `json:"fixed_price_quantity"`
	// The property used to group this price on an invoice
	InvoiceGroupingKey param.Field[string] `json:"invoice_grouping_key"`
	// Within each billing cycle, specifies the cadence at which invoices are produced.
	// If unspecified, a single invoice is produced per billing cycle.
	InvoicingCycleConfiguration param.Field[PlanVersionNewParamsReplacePricesPriceNewPlanScalableMatrixWithUnitPricingPriceInvoicingCycleConfiguration] `json:"invoicing_cycle_configuration"`
	// User-specified key/value pairs for the resource. Individual keys can be removed
	// by setting the value to `null`, and the entire metadata mapping can be cleared
	// by setting `metadata` to `null`.
	Metadata param.Field[map[string]string] `json:"metadata"`
}

func (r PlanVersionNewParamsReplacePricesPriceNewPlanScalableMatrixWithUnitPricingPrice) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r PlanVersionNewParamsReplacePricesPriceNewPlanScalableMatrixWithUnitPricingPrice) implementsPlanVersionNewParamsReplacePricesPriceUnion() {
}

// The cadence to bill for this price on.
type PlanVersionNewParamsReplacePricesPriceNewPlanScalableMatrixWithUnitPricingPriceCadence string

const (
	PlanVersionNewParamsReplacePricesPriceNewPlanScalableMatrixWithUnitPricingPriceCadenceAnnual     PlanVersionNewParamsReplacePricesPriceNewPlanScalableMatrixWithUnitPricingPriceCadence = "annual"
	PlanVersionNewParamsReplacePricesPriceNewPlanScalableMatrixWithUnitPricingPriceCadenceSemiAnnual PlanVersionNewParamsReplacePricesPriceNewPlanScalableMatrixWithUnitPricingPriceCadence = "semi_annual"
	PlanVersionNewParamsReplacePricesPriceNewPlanScalableMatrixWithUnitPricingPriceCadenceMonthly    PlanVersionNewParamsReplacePricesPriceNewPlanScalableMatrixWithUnitPricingPriceCadence = "monthly"
	PlanVersionNewParamsReplacePricesPriceNewPlanScalableMatrixWithUnitPricingPriceCadenceQuarterly  PlanVersionNewParamsReplacePricesPriceNewPlanScalableMatrixWithUnitPricingPriceCadence = "quarterly"
	PlanVersionNewParamsReplacePricesPriceNewPlanScalableMatrixWithUnitPricingPriceCadenceOneTime    PlanVersionNewParamsReplacePricesPriceNewPlanScalableMatrixWithUnitPricingPriceCadence = "one_time"
	PlanVersionNewParamsReplacePricesPriceNewPlanScalableMatrixWithUnitPricingPriceCadenceCustom     PlanVersionNewParamsReplacePricesPriceNewPlanScalableMatrixWithUnitPricingPriceCadence = "custom"
)

func (r PlanVersionNewParamsReplacePricesPriceNewPlanScalableMatrixWithUnitPricingPriceCadence) IsKnown() bool {
	switch r {
	case PlanVersionNewParamsReplacePricesPriceNewPlanScalableMatrixWithUnitPricingPriceCadenceAnnual, PlanVersionNewParamsReplacePricesPriceNewPlanScalableMatrixWithUnitPricingPriceCadenceSemiAnnual, PlanVersionNewParamsReplacePricesPriceNewPlanScalableMatrixWithUnitPricingPriceCadenceMonthly, PlanVersionNewParamsReplacePricesPriceNewPlanScalableMatrixWithUnitPricingPriceCadenceQuarterly, PlanVersionNewParamsReplacePricesPriceNewPlanScalableMatrixWithUnitPricingPriceCadenceOneTime, PlanVersionNewParamsReplacePricesPriceNewPlanScalableMatrixWithUnitPricingPriceCadenceCustom:
		return true
	}
	return false
}

type PlanVersionNewParamsReplacePricesPriceNewPlanScalableMatrixWithUnitPricingPriceModelType string

const (
	PlanVersionNewParamsReplacePricesPriceNewPlanScalableMatrixWithUnitPricingPriceModelTypeScalableMatrixWithUnitPricing PlanVersionNewParamsReplacePricesPriceNewPlanScalableMatrixWithUnitPricingPriceModelType = "scalable_matrix_with_unit_pricing"
)

func (r PlanVersionNewParamsReplacePricesPriceNewPlanScalableMatrixWithUnitPricingPriceModelType) IsKnown() bool {
	switch r {
	case PlanVersionNewParamsReplacePricesPriceNewPlanScalableMatrixWithUnitPricingPriceModelTypeScalableMatrixWithUnitPricing:
		return true
	}
	return false
}

// For custom cadence: specifies the duration of the billing period in days or
// months.
type PlanVersionNewParamsReplacePricesPriceNewPlanScalableMatrixWithUnitPricingPriceBillingCycleConfiguration struct {
	// The duration of the billing period.
	Duration param.Field[int64] `json:"duration,required"`
	// The unit of billing period duration.
	DurationUnit param.Field[PlanVersionNewParamsReplacePricesPriceNewPlanScalableMatrixWithUnitPricingPriceBillingCycleConfigurationDurationUnit] `json:"duration_unit,required"`
}

func (r PlanVersionNewParamsReplacePricesPriceNewPlanScalableMatrixWithUnitPricingPriceBillingCycleConfiguration) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// The unit of billing period duration.
type PlanVersionNewParamsReplacePricesPriceNewPlanScalableMatrixWithUnitPricingPriceBillingCycleConfigurationDurationUnit string

const (
	PlanVersionNewParamsReplacePricesPriceNewPlanScalableMatrixWithUnitPricingPriceBillingCycleConfigurationDurationUnitDay   PlanVersionNewParamsReplacePricesPriceNewPlanScalableMatrixWithUnitPricingPriceBillingCycleConfigurationDurationUnit = "day"
	PlanVersionNewParamsReplacePricesPriceNewPlanScalableMatrixWithUnitPricingPriceBillingCycleConfigurationDurationUnitMonth PlanVersionNewParamsReplacePricesPriceNewPlanScalableMatrixWithUnitPricingPriceBillingCycleConfigurationDurationUnit = "month"
)

func (r PlanVersionNewParamsReplacePricesPriceNewPlanScalableMatrixWithUnitPricingPriceBillingCycleConfigurationDurationUnit) IsKnown() bool {
	switch r {
	case PlanVersionNewParamsReplacePricesPriceNewPlanScalableMatrixWithUnitPricingPriceBillingCycleConfigurationDurationUnitDay, PlanVersionNewParamsReplacePricesPriceNewPlanScalableMatrixWithUnitPricingPriceBillingCycleConfigurationDurationUnitMonth:
		return true
	}
	return false
}

// For dimensional price: specifies a price group and dimension values
type PlanVersionNewParamsReplacePricesPriceNewPlanScalableMatrixWithUnitPricingPriceDimensionalPriceConfiguration struct {
	// The list of dimension values matching (in order) the dimensions of the price
	// group
	DimensionValues param.Field[[]string] `json:"dimension_values,required"`
	// The id of the dimensional price group to include this price in
	DimensionalPriceGroupID param.Field[string] `json:"dimensional_price_group_id"`
	// The external id of the dimensional price group to include this price in
	ExternalDimensionalPriceGroupID param.Field[string] `json:"external_dimensional_price_group_id"`
}

func (r PlanVersionNewParamsReplacePricesPriceNewPlanScalableMatrixWithUnitPricingPriceDimensionalPriceConfiguration) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Within each billing cycle, specifies the cadence at which invoices are produced.
// If unspecified, a single invoice is produced per billing cycle.
type PlanVersionNewParamsReplacePricesPriceNewPlanScalableMatrixWithUnitPricingPriceInvoicingCycleConfiguration struct {
	// The duration of the billing period.
	Duration param.Field[int64] `json:"duration,required"`
	// The unit of billing period duration.
	DurationUnit param.Field[PlanVersionNewParamsReplacePricesPriceNewPlanScalableMatrixWithUnitPricingPriceInvoicingCycleConfigurationDurationUnit] `json:"duration_unit,required"`
}

func (r PlanVersionNewParamsReplacePricesPriceNewPlanScalableMatrixWithUnitPricingPriceInvoicingCycleConfiguration) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// The unit of billing period duration.
type PlanVersionNewParamsReplacePricesPriceNewPlanScalableMatrixWithUnitPricingPriceInvoicingCycleConfigurationDurationUnit string

const (
	PlanVersionNewParamsReplacePricesPriceNewPlanScalableMatrixWithUnitPricingPriceInvoicingCycleConfigurationDurationUnitDay   PlanVersionNewParamsReplacePricesPriceNewPlanScalableMatrixWithUnitPricingPriceInvoicingCycleConfigurationDurationUnit = "day"
	PlanVersionNewParamsReplacePricesPriceNewPlanScalableMatrixWithUnitPricingPriceInvoicingCycleConfigurationDurationUnitMonth PlanVersionNewParamsReplacePricesPriceNewPlanScalableMatrixWithUnitPricingPriceInvoicingCycleConfigurationDurationUnit = "month"
)

func (r PlanVersionNewParamsReplacePricesPriceNewPlanScalableMatrixWithUnitPricingPriceInvoicingCycleConfigurationDurationUnit) IsKnown() bool {
	switch r {
	case PlanVersionNewParamsReplacePricesPriceNewPlanScalableMatrixWithUnitPricingPriceInvoicingCycleConfigurationDurationUnitDay, PlanVersionNewParamsReplacePricesPriceNewPlanScalableMatrixWithUnitPricingPriceInvoicingCycleConfigurationDurationUnitMonth:
		return true
	}
	return false
}

type PlanVersionNewParamsReplacePricesPriceNewPlanScalableMatrixWithTieredPricingPrice struct {
	// The cadence to bill for this price on.
	Cadence param.Field[PlanVersionNewParamsReplacePricesPriceNewPlanScalableMatrixWithTieredPricingPriceCadence] `json:"cadence,required"`
	// The id of the item the price will be associated with.
	ItemID    param.Field[string]                                                                                     `json:"item_id,required"`
	ModelType param.Field[PlanVersionNewParamsReplacePricesPriceNewPlanScalableMatrixWithTieredPricingPriceModelType] `json:"model_type,required"`
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
	BillingCycleConfiguration param.Field[PlanVersionNewParamsReplacePricesPriceNewPlanScalableMatrixWithTieredPricingPriceBillingCycleConfiguration] `json:"billing_cycle_configuration"`
	// The per unit conversion rate of the price currency to the invoicing currency.
	ConversionRate param.Field[float64] `json:"conversion_rate"`
	// An ISO 4217 currency string, or custom pricing unit identifier, in which this
	// price is billed.
	Currency param.Field[string] `json:"currency"`
	// For dimensional price: specifies a price group and dimension values
	DimensionalPriceConfiguration param.Field[PlanVersionNewParamsReplacePricesPriceNewPlanScalableMatrixWithTieredPricingPriceDimensionalPriceConfiguration] `json:"dimensional_price_configuration"`
	// An alias for the price.
	ExternalPriceID param.Field[string] `json:"external_price_id"`
	// If the Price represents a fixed cost, this represents the quantity of units
	// applied.
	FixedPriceQuantity param.Field[float64] `json:"fixed_price_quantity"`
	// The property used to group this price on an invoice
	InvoiceGroupingKey param.Field[string] `json:"invoice_grouping_key"`
	// Within each billing cycle, specifies the cadence at which invoices are produced.
	// If unspecified, a single invoice is produced per billing cycle.
	InvoicingCycleConfiguration param.Field[PlanVersionNewParamsReplacePricesPriceNewPlanScalableMatrixWithTieredPricingPriceInvoicingCycleConfiguration] `json:"invoicing_cycle_configuration"`
	// User-specified key/value pairs for the resource. Individual keys can be removed
	// by setting the value to `null`, and the entire metadata mapping can be cleared
	// by setting `metadata` to `null`.
	Metadata param.Field[map[string]string] `json:"metadata"`
}

func (r PlanVersionNewParamsReplacePricesPriceNewPlanScalableMatrixWithTieredPricingPrice) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r PlanVersionNewParamsReplacePricesPriceNewPlanScalableMatrixWithTieredPricingPrice) implementsPlanVersionNewParamsReplacePricesPriceUnion() {
}

// The cadence to bill for this price on.
type PlanVersionNewParamsReplacePricesPriceNewPlanScalableMatrixWithTieredPricingPriceCadence string

const (
	PlanVersionNewParamsReplacePricesPriceNewPlanScalableMatrixWithTieredPricingPriceCadenceAnnual     PlanVersionNewParamsReplacePricesPriceNewPlanScalableMatrixWithTieredPricingPriceCadence = "annual"
	PlanVersionNewParamsReplacePricesPriceNewPlanScalableMatrixWithTieredPricingPriceCadenceSemiAnnual PlanVersionNewParamsReplacePricesPriceNewPlanScalableMatrixWithTieredPricingPriceCadence = "semi_annual"
	PlanVersionNewParamsReplacePricesPriceNewPlanScalableMatrixWithTieredPricingPriceCadenceMonthly    PlanVersionNewParamsReplacePricesPriceNewPlanScalableMatrixWithTieredPricingPriceCadence = "monthly"
	PlanVersionNewParamsReplacePricesPriceNewPlanScalableMatrixWithTieredPricingPriceCadenceQuarterly  PlanVersionNewParamsReplacePricesPriceNewPlanScalableMatrixWithTieredPricingPriceCadence = "quarterly"
	PlanVersionNewParamsReplacePricesPriceNewPlanScalableMatrixWithTieredPricingPriceCadenceOneTime    PlanVersionNewParamsReplacePricesPriceNewPlanScalableMatrixWithTieredPricingPriceCadence = "one_time"
	PlanVersionNewParamsReplacePricesPriceNewPlanScalableMatrixWithTieredPricingPriceCadenceCustom     PlanVersionNewParamsReplacePricesPriceNewPlanScalableMatrixWithTieredPricingPriceCadence = "custom"
)

func (r PlanVersionNewParamsReplacePricesPriceNewPlanScalableMatrixWithTieredPricingPriceCadence) IsKnown() bool {
	switch r {
	case PlanVersionNewParamsReplacePricesPriceNewPlanScalableMatrixWithTieredPricingPriceCadenceAnnual, PlanVersionNewParamsReplacePricesPriceNewPlanScalableMatrixWithTieredPricingPriceCadenceSemiAnnual, PlanVersionNewParamsReplacePricesPriceNewPlanScalableMatrixWithTieredPricingPriceCadenceMonthly, PlanVersionNewParamsReplacePricesPriceNewPlanScalableMatrixWithTieredPricingPriceCadenceQuarterly, PlanVersionNewParamsReplacePricesPriceNewPlanScalableMatrixWithTieredPricingPriceCadenceOneTime, PlanVersionNewParamsReplacePricesPriceNewPlanScalableMatrixWithTieredPricingPriceCadenceCustom:
		return true
	}
	return false
}

type PlanVersionNewParamsReplacePricesPriceNewPlanScalableMatrixWithTieredPricingPriceModelType string

const (
	PlanVersionNewParamsReplacePricesPriceNewPlanScalableMatrixWithTieredPricingPriceModelTypeScalableMatrixWithTieredPricing PlanVersionNewParamsReplacePricesPriceNewPlanScalableMatrixWithTieredPricingPriceModelType = "scalable_matrix_with_tiered_pricing"
)

func (r PlanVersionNewParamsReplacePricesPriceNewPlanScalableMatrixWithTieredPricingPriceModelType) IsKnown() bool {
	switch r {
	case PlanVersionNewParamsReplacePricesPriceNewPlanScalableMatrixWithTieredPricingPriceModelTypeScalableMatrixWithTieredPricing:
		return true
	}
	return false
}

// For custom cadence: specifies the duration of the billing period in days or
// months.
type PlanVersionNewParamsReplacePricesPriceNewPlanScalableMatrixWithTieredPricingPriceBillingCycleConfiguration struct {
	// The duration of the billing period.
	Duration param.Field[int64] `json:"duration,required"`
	// The unit of billing period duration.
	DurationUnit param.Field[PlanVersionNewParamsReplacePricesPriceNewPlanScalableMatrixWithTieredPricingPriceBillingCycleConfigurationDurationUnit] `json:"duration_unit,required"`
}

func (r PlanVersionNewParamsReplacePricesPriceNewPlanScalableMatrixWithTieredPricingPriceBillingCycleConfiguration) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// The unit of billing period duration.
type PlanVersionNewParamsReplacePricesPriceNewPlanScalableMatrixWithTieredPricingPriceBillingCycleConfigurationDurationUnit string

const (
	PlanVersionNewParamsReplacePricesPriceNewPlanScalableMatrixWithTieredPricingPriceBillingCycleConfigurationDurationUnitDay   PlanVersionNewParamsReplacePricesPriceNewPlanScalableMatrixWithTieredPricingPriceBillingCycleConfigurationDurationUnit = "day"
	PlanVersionNewParamsReplacePricesPriceNewPlanScalableMatrixWithTieredPricingPriceBillingCycleConfigurationDurationUnitMonth PlanVersionNewParamsReplacePricesPriceNewPlanScalableMatrixWithTieredPricingPriceBillingCycleConfigurationDurationUnit = "month"
)

func (r PlanVersionNewParamsReplacePricesPriceNewPlanScalableMatrixWithTieredPricingPriceBillingCycleConfigurationDurationUnit) IsKnown() bool {
	switch r {
	case PlanVersionNewParamsReplacePricesPriceNewPlanScalableMatrixWithTieredPricingPriceBillingCycleConfigurationDurationUnitDay, PlanVersionNewParamsReplacePricesPriceNewPlanScalableMatrixWithTieredPricingPriceBillingCycleConfigurationDurationUnitMonth:
		return true
	}
	return false
}

// For dimensional price: specifies a price group and dimension values
type PlanVersionNewParamsReplacePricesPriceNewPlanScalableMatrixWithTieredPricingPriceDimensionalPriceConfiguration struct {
	// The list of dimension values matching (in order) the dimensions of the price
	// group
	DimensionValues param.Field[[]string] `json:"dimension_values,required"`
	// The id of the dimensional price group to include this price in
	DimensionalPriceGroupID param.Field[string] `json:"dimensional_price_group_id"`
	// The external id of the dimensional price group to include this price in
	ExternalDimensionalPriceGroupID param.Field[string] `json:"external_dimensional_price_group_id"`
}

func (r PlanVersionNewParamsReplacePricesPriceNewPlanScalableMatrixWithTieredPricingPriceDimensionalPriceConfiguration) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Within each billing cycle, specifies the cadence at which invoices are produced.
// If unspecified, a single invoice is produced per billing cycle.
type PlanVersionNewParamsReplacePricesPriceNewPlanScalableMatrixWithTieredPricingPriceInvoicingCycleConfiguration struct {
	// The duration of the billing period.
	Duration param.Field[int64] `json:"duration,required"`
	// The unit of billing period duration.
	DurationUnit param.Field[PlanVersionNewParamsReplacePricesPriceNewPlanScalableMatrixWithTieredPricingPriceInvoicingCycleConfigurationDurationUnit] `json:"duration_unit,required"`
}

func (r PlanVersionNewParamsReplacePricesPriceNewPlanScalableMatrixWithTieredPricingPriceInvoicingCycleConfiguration) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// The unit of billing period duration.
type PlanVersionNewParamsReplacePricesPriceNewPlanScalableMatrixWithTieredPricingPriceInvoicingCycleConfigurationDurationUnit string

const (
	PlanVersionNewParamsReplacePricesPriceNewPlanScalableMatrixWithTieredPricingPriceInvoicingCycleConfigurationDurationUnitDay   PlanVersionNewParamsReplacePricesPriceNewPlanScalableMatrixWithTieredPricingPriceInvoicingCycleConfigurationDurationUnit = "day"
	PlanVersionNewParamsReplacePricesPriceNewPlanScalableMatrixWithTieredPricingPriceInvoicingCycleConfigurationDurationUnitMonth PlanVersionNewParamsReplacePricesPriceNewPlanScalableMatrixWithTieredPricingPriceInvoicingCycleConfigurationDurationUnit = "month"
)

func (r PlanVersionNewParamsReplacePricesPriceNewPlanScalableMatrixWithTieredPricingPriceInvoicingCycleConfigurationDurationUnit) IsKnown() bool {
	switch r {
	case PlanVersionNewParamsReplacePricesPriceNewPlanScalableMatrixWithTieredPricingPriceInvoicingCycleConfigurationDurationUnitDay, PlanVersionNewParamsReplacePricesPriceNewPlanScalableMatrixWithTieredPricingPriceInvoicingCycleConfigurationDurationUnitMonth:
		return true
	}
	return false
}

type PlanVersionNewParamsReplacePricesPriceNewPlanCumulativeGroupedBulkPrice struct {
	// The cadence to bill for this price on.
	Cadence                     param.Field[PlanVersionNewParamsReplacePricesPriceNewPlanCumulativeGroupedBulkPriceCadence] `json:"cadence,required"`
	CumulativeGroupedBulkConfig param.Field[map[string]interface{}]                                                         `json:"cumulative_grouped_bulk_config,required"`
	// The id of the item the price will be associated with.
	ItemID    param.Field[string]                                                                           `json:"item_id,required"`
	ModelType param.Field[PlanVersionNewParamsReplacePricesPriceNewPlanCumulativeGroupedBulkPriceModelType] `json:"model_type,required"`
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
	BillingCycleConfiguration param.Field[PlanVersionNewParamsReplacePricesPriceNewPlanCumulativeGroupedBulkPriceBillingCycleConfiguration] `json:"billing_cycle_configuration"`
	// The per unit conversion rate of the price currency to the invoicing currency.
	ConversionRate param.Field[float64] `json:"conversion_rate"`
	// An ISO 4217 currency string, or custom pricing unit identifier, in which this
	// price is billed.
	Currency param.Field[string] `json:"currency"`
	// For dimensional price: specifies a price group and dimension values
	DimensionalPriceConfiguration param.Field[PlanVersionNewParamsReplacePricesPriceNewPlanCumulativeGroupedBulkPriceDimensionalPriceConfiguration] `json:"dimensional_price_configuration"`
	// An alias for the price.
	ExternalPriceID param.Field[string] `json:"external_price_id"`
	// If the Price represents a fixed cost, this represents the quantity of units
	// applied.
	FixedPriceQuantity param.Field[float64] `json:"fixed_price_quantity"`
	// The property used to group this price on an invoice
	InvoiceGroupingKey param.Field[string] `json:"invoice_grouping_key"`
	// Within each billing cycle, specifies the cadence at which invoices are produced.
	// If unspecified, a single invoice is produced per billing cycle.
	InvoicingCycleConfiguration param.Field[PlanVersionNewParamsReplacePricesPriceNewPlanCumulativeGroupedBulkPriceInvoicingCycleConfiguration] `json:"invoicing_cycle_configuration"`
	// User-specified key/value pairs for the resource. Individual keys can be removed
	// by setting the value to `null`, and the entire metadata mapping can be cleared
	// by setting `metadata` to `null`.
	Metadata param.Field[map[string]string] `json:"metadata"`
}

func (r PlanVersionNewParamsReplacePricesPriceNewPlanCumulativeGroupedBulkPrice) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r PlanVersionNewParamsReplacePricesPriceNewPlanCumulativeGroupedBulkPrice) implementsPlanVersionNewParamsReplacePricesPriceUnion() {
}

// The cadence to bill for this price on.
type PlanVersionNewParamsReplacePricesPriceNewPlanCumulativeGroupedBulkPriceCadence string

const (
	PlanVersionNewParamsReplacePricesPriceNewPlanCumulativeGroupedBulkPriceCadenceAnnual     PlanVersionNewParamsReplacePricesPriceNewPlanCumulativeGroupedBulkPriceCadence = "annual"
	PlanVersionNewParamsReplacePricesPriceNewPlanCumulativeGroupedBulkPriceCadenceSemiAnnual PlanVersionNewParamsReplacePricesPriceNewPlanCumulativeGroupedBulkPriceCadence = "semi_annual"
	PlanVersionNewParamsReplacePricesPriceNewPlanCumulativeGroupedBulkPriceCadenceMonthly    PlanVersionNewParamsReplacePricesPriceNewPlanCumulativeGroupedBulkPriceCadence = "monthly"
	PlanVersionNewParamsReplacePricesPriceNewPlanCumulativeGroupedBulkPriceCadenceQuarterly  PlanVersionNewParamsReplacePricesPriceNewPlanCumulativeGroupedBulkPriceCadence = "quarterly"
	PlanVersionNewParamsReplacePricesPriceNewPlanCumulativeGroupedBulkPriceCadenceOneTime    PlanVersionNewParamsReplacePricesPriceNewPlanCumulativeGroupedBulkPriceCadence = "one_time"
	PlanVersionNewParamsReplacePricesPriceNewPlanCumulativeGroupedBulkPriceCadenceCustom     PlanVersionNewParamsReplacePricesPriceNewPlanCumulativeGroupedBulkPriceCadence = "custom"
)

func (r PlanVersionNewParamsReplacePricesPriceNewPlanCumulativeGroupedBulkPriceCadence) IsKnown() bool {
	switch r {
	case PlanVersionNewParamsReplacePricesPriceNewPlanCumulativeGroupedBulkPriceCadenceAnnual, PlanVersionNewParamsReplacePricesPriceNewPlanCumulativeGroupedBulkPriceCadenceSemiAnnual, PlanVersionNewParamsReplacePricesPriceNewPlanCumulativeGroupedBulkPriceCadenceMonthly, PlanVersionNewParamsReplacePricesPriceNewPlanCumulativeGroupedBulkPriceCadenceQuarterly, PlanVersionNewParamsReplacePricesPriceNewPlanCumulativeGroupedBulkPriceCadenceOneTime, PlanVersionNewParamsReplacePricesPriceNewPlanCumulativeGroupedBulkPriceCadenceCustom:
		return true
	}
	return false
}

type PlanVersionNewParamsReplacePricesPriceNewPlanCumulativeGroupedBulkPriceModelType string

const (
	PlanVersionNewParamsReplacePricesPriceNewPlanCumulativeGroupedBulkPriceModelTypeCumulativeGroupedBulk PlanVersionNewParamsReplacePricesPriceNewPlanCumulativeGroupedBulkPriceModelType = "cumulative_grouped_bulk"
)

func (r PlanVersionNewParamsReplacePricesPriceNewPlanCumulativeGroupedBulkPriceModelType) IsKnown() bool {
	switch r {
	case PlanVersionNewParamsReplacePricesPriceNewPlanCumulativeGroupedBulkPriceModelTypeCumulativeGroupedBulk:
		return true
	}
	return false
}

// For custom cadence: specifies the duration of the billing period in days or
// months.
type PlanVersionNewParamsReplacePricesPriceNewPlanCumulativeGroupedBulkPriceBillingCycleConfiguration struct {
	// The duration of the billing period.
	Duration param.Field[int64] `json:"duration,required"`
	// The unit of billing period duration.
	DurationUnit param.Field[PlanVersionNewParamsReplacePricesPriceNewPlanCumulativeGroupedBulkPriceBillingCycleConfigurationDurationUnit] `json:"duration_unit,required"`
}

func (r PlanVersionNewParamsReplacePricesPriceNewPlanCumulativeGroupedBulkPriceBillingCycleConfiguration) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// The unit of billing period duration.
type PlanVersionNewParamsReplacePricesPriceNewPlanCumulativeGroupedBulkPriceBillingCycleConfigurationDurationUnit string

const (
	PlanVersionNewParamsReplacePricesPriceNewPlanCumulativeGroupedBulkPriceBillingCycleConfigurationDurationUnitDay   PlanVersionNewParamsReplacePricesPriceNewPlanCumulativeGroupedBulkPriceBillingCycleConfigurationDurationUnit = "day"
	PlanVersionNewParamsReplacePricesPriceNewPlanCumulativeGroupedBulkPriceBillingCycleConfigurationDurationUnitMonth PlanVersionNewParamsReplacePricesPriceNewPlanCumulativeGroupedBulkPriceBillingCycleConfigurationDurationUnit = "month"
)

func (r PlanVersionNewParamsReplacePricesPriceNewPlanCumulativeGroupedBulkPriceBillingCycleConfigurationDurationUnit) IsKnown() bool {
	switch r {
	case PlanVersionNewParamsReplacePricesPriceNewPlanCumulativeGroupedBulkPriceBillingCycleConfigurationDurationUnitDay, PlanVersionNewParamsReplacePricesPriceNewPlanCumulativeGroupedBulkPriceBillingCycleConfigurationDurationUnitMonth:
		return true
	}
	return false
}

// For dimensional price: specifies a price group and dimension values
type PlanVersionNewParamsReplacePricesPriceNewPlanCumulativeGroupedBulkPriceDimensionalPriceConfiguration struct {
	// The list of dimension values matching (in order) the dimensions of the price
	// group
	DimensionValues param.Field[[]string] `json:"dimension_values,required"`
	// The id of the dimensional price group to include this price in
	DimensionalPriceGroupID param.Field[string] `json:"dimensional_price_group_id"`
	// The external id of the dimensional price group to include this price in
	ExternalDimensionalPriceGroupID param.Field[string] `json:"external_dimensional_price_group_id"`
}

func (r PlanVersionNewParamsReplacePricesPriceNewPlanCumulativeGroupedBulkPriceDimensionalPriceConfiguration) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Within each billing cycle, specifies the cadence at which invoices are produced.
// If unspecified, a single invoice is produced per billing cycle.
type PlanVersionNewParamsReplacePricesPriceNewPlanCumulativeGroupedBulkPriceInvoicingCycleConfiguration struct {
	// The duration of the billing period.
	Duration param.Field[int64] `json:"duration,required"`
	// The unit of billing period duration.
	DurationUnit param.Field[PlanVersionNewParamsReplacePricesPriceNewPlanCumulativeGroupedBulkPriceInvoicingCycleConfigurationDurationUnit] `json:"duration_unit,required"`
}

func (r PlanVersionNewParamsReplacePricesPriceNewPlanCumulativeGroupedBulkPriceInvoicingCycleConfiguration) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// The unit of billing period duration.
type PlanVersionNewParamsReplacePricesPriceNewPlanCumulativeGroupedBulkPriceInvoicingCycleConfigurationDurationUnit string

const (
	PlanVersionNewParamsReplacePricesPriceNewPlanCumulativeGroupedBulkPriceInvoicingCycleConfigurationDurationUnitDay   PlanVersionNewParamsReplacePricesPriceNewPlanCumulativeGroupedBulkPriceInvoicingCycleConfigurationDurationUnit = "day"
	PlanVersionNewParamsReplacePricesPriceNewPlanCumulativeGroupedBulkPriceInvoicingCycleConfigurationDurationUnitMonth PlanVersionNewParamsReplacePricesPriceNewPlanCumulativeGroupedBulkPriceInvoicingCycleConfigurationDurationUnit = "month"
)

func (r PlanVersionNewParamsReplacePricesPriceNewPlanCumulativeGroupedBulkPriceInvoicingCycleConfigurationDurationUnit) IsKnown() bool {
	switch r {
	case PlanVersionNewParamsReplacePricesPriceNewPlanCumulativeGroupedBulkPriceInvoicingCycleConfigurationDurationUnitDay, PlanVersionNewParamsReplacePricesPriceNewPlanCumulativeGroupedBulkPriceInvoicingCycleConfigurationDurationUnitMonth:
		return true
	}
	return false
}

type PlanVersionNewParamsReplacePricesPriceNewPlanTieredPackageWithMinimumPrice struct {
	// The cadence to bill for this price on.
	Cadence param.Field[PlanVersionNewParamsReplacePricesPriceNewPlanTieredPackageWithMinimumPriceCadence] `json:"cadence,required"`
	// The id of the item the price will be associated with.
	ItemID    param.Field[string]                                                                              `json:"item_id,required"`
	ModelType param.Field[PlanVersionNewParamsReplacePricesPriceNewPlanTieredPackageWithMinimumPriceModelType] `json:"model_type,required"`
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
	BillingCycleConfiguration param.Field[PlanVersionNewParamsReplacePricesPriceNewPlanTieredPackageWithMinimumPriceBillingCycleConfiguration] `json:"billing_cycle_configuration"`
	// The per unit conversion rate of the price currency to the invoicing currency.
	ConversionRate param.Field[float64] `json:"conversion_rate"`
	// An ISO 4217 currency string, or custom pricing unit identifier, in which this
	// price is billed.
	Currency param.Field[string] `json:"currency"`
	// For dimensional price: specifies a price group and dimension values
	DimensionalPriceConfiguration param.Field[PlanVersionNewParamsReplacePricesPriceNewPlanTieredPackageWithMinimumPriceDimensionalPriceConfiguration] `json:"dimensional_price_configuration"`
	// An alias for the price.
	ExternalPriceID param.Field[string] `json:"external_price_id"`
	// If the Price represents a fixed cost, this represents the quantity of units
	// applied.
	FixedPriceQuantity param.Field[float64] `json:"fixed_price_quantity"`
	// The property used to group this price on an invoice
	InvoiceGroupingKey param.Field[string] `json:"invoice_grouping_key"`
	// Within each billing cycle, specifies the cadence at which invoices are produced.
	// If unspecified, a single invoice is produced per billing cycle.
	InvoicingCycleConfiguration param.Field[PlanVersionNewParamsReplacePricesPriceNewPlanTieredPackageWithMinimumPriceInvoicingCycleConfiguration] `json:"invoicing_cycle_configuration"`
	// User-specified key/value pairs for the resource. Individual keys can be removed
	// by setting the value to `null`, and the entire metadata mapping can be cleared
	// by setting `metadata` to `null`.
	Metadata param.Field[map[string]string] `json:"metadata"`
}

func (r PlanVersionNewParamsReplacePricesPriceNewPlanTieredPackageWithMinimumPrice) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r PlanVersionNewParamsReplacePricesPriceNewPlanTieredPackageWithMinimumPrice) implementsPlanVersionNewParamsReplacePricesPriceUnion() {
}

// The cadence to bill for this price on.
type PlanVersionNewParamsReplacePricesPriceNewPlanTieredPackageWithMinimumPriceCadence string

const (
	PlanVersionNewParamsReplacePricesPriceNewPlanTieredPackageWithMinimumPriceCadenceAnnual     PlanVersionNewParamsReplacePricesPriceNewPlanTieredPackageWithMinimumPriceCadence = "annual"
	PlanVersionNewParamsReplacePricesPriceNewPlanTieredPackageWithMinimumPriceCadenceSemiAnnual PlanVersionNewParamsReplacePricesPriceNewPlanTieredPackageWithMinimumPriceCadence = "semi_annual"
	PlanVersionNewParamsReplacePricesPriceNewPlanTieredPackageWithMinimumPriceCadenceMonthly    PlanVersionNewParamsReplacePricesPriceNewPlanTieredPackageWithMinimumPriceCadence = "monthly"
	PlanVersionNewParamsReplacePricesPriceNewPlanTieredPackageWithMinimumPriceCadenceQuarterly  PlanVersionNewParamsReplacePricesPriceNewPlanTieredPackageWithMinimumPriceCadence = "quarterly"
	PlanVersionNewParamsReplacePricesPriceNewPlanTieredPackageWithMinimumPriceCadenceOneTime    PlanVersionNewParamsReplacePricesPriceNewPlanTieredPackageWithMinimumPriceCadence = "one_time"
	PlanVersionNewParamsReplacePricesPriceNewPlanTieredPackageWithMinimumPriceCadenceCustom     PlanVersionNewParamsReplacePricesPriceNewPlanTieredPackageWithMinimumPriceCadence = "custom"
)

func (r PlanVersionNewParamsReplacePricesPriceNewPlanTieredPackageWithMinimumPriceCadence) IsKnown() bool {
	switch r {
	case PlanVersionNewParamsReplacePricesPriceNewPlanTieredPackageWithMinimumPriceCadenceAnnual, PlanVersionNewParamsReplacePricesPriceNewPlanTieredPackageWithMinimumPriceCadenceSemiAnnual, PlanVersionNewParamsReplacePricesPriceNewPlanTieredPackageWithMinimumPriceCadenceMonthly, PlanVersionNewParamsReplacePricesPriceNewPlanTieredPackageWithMinimumPriceCadenceQuarterly, PlanVersionNewParamsReplacePricesPriceNewPlanTieredPackageWithMinimumPriceCadenceOneTime, PlanVersionNewParamsReplacePricesPriceNewPlanTieredPackageWithMinimumPriceCadenceCustom:
		return true
	}
	return false
}

type PlanVersionNewParamsReplacePricesPriceNewPlanTieredPackageWithMinimumPriceModelType string

const (
	PlanVersionNewParamsReplacePricesPriceNewPlanTieredPackageWithMinimumPriceModelTypeTieredPackageWithMinimum PlanVersionNewParamsReplacePricesPriceNewPlanTieredPackageWithMinimumPriceModelType = "tiered_package_with_minimum"
)

func (r PlanVersionNewParamsReplacePricesPriceNewPlanTieredPackageWithMinimumPriceModelType) IsKnown() bool {
	switch r {
	case PlanVersionNewParamsReplacePricesPriceNewPlanTieredPackageWithMinimumPriceModelTypeTieredPackageWithMinimum:
		return true
	}
	return false
}

// For custom cadence: specifies the duration of the billing period in days or
// months.
type PlanVersionNewParamsReplacePricesPriceNewPlanTieredPackageWithMinimumPriceBillingCycleConfiguration struct {
	// The duration of the billing period.
	Duration param.Field[int64] `json:"duration,required"`
	// The unit of billing period duration.
	DurationUnit param.Field[PlanVersionNewParamsReplacePricesPriceNewPlanTieredPackageWithMinimumPriceBillingCycleConfigurationDurationUnit] `json:"duration_unit,required"`
}

func (r PlanVersionNewParamsReplacePricesPriceNewPlanTieredPackageWithMinimumPriceBillingCycleConfiguration) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// The unit of billing period duration.
type PlanVersionNewParamsReplacePricesPriceNewPlanTieredPackageWithMinimumPriceBillingCycleConfigurationDurationUnit string

const (
	PlanVersionNewParamsReplacePricesPriceNewPlanTieredPackageWithMinimumPriceBillingCycleConfigurationDurationUnitDay   PlanVersionNewParamsReplacePricesPriceNewPlanTieredPackageWithMinimumPriceBillingCycleConfigurationDurationUnit = "day"
	PlanVersionNewParamsReplacePricesPriceNewPlanTieredPackageWithMinimumPriceBillingCycleConfigurationDurationUnitMonth PlanVersionNewParamsReplacePricesPriceNewPlanTieredPackageWithMinimumPriceBillingCycleConfigurationDurationUnit = "month"
)

func (r PlanVersionNewParamsReplacePricesPriceNewPlanTieredPackageWithMinimumPriceBillingCycleConfigurationDurationUnit) IsKnown() bool {
	switch r {
	case PlanVersionNewParamsReplacePricesPriceNewPlanTieredPackageWithMinimumPriceBillingCycleConfigurationDurationUnitDay, PlanVersionNewParamsReplacePricesPriceNewPlanTieredPackageWithMinimumPriceBillingCycleConfigurationDurationUnitMonth:
		return true
	}
	return false
}

// For dimensional price: specifies a price group and dimension values
type PlanVersionNewParamsReplacePricesPriceNewPlanTieredPackageWithMinimumPriceDimensionalPriceConfiguration struct {
	// The list of dimension values matching (in order) the dimensions of the price
	// group
	DimensionValues param.Field[[]string] `json:"dimension_values,required"`
	// The id of the dimensional price group to include this price in
	DimensionalPriceGroupID param.Field[string] `json:"dimensional_price_group_id"`
	// The external id of the dimensional price group to include this price in
	ExternalDimensionalPriceGroupID param.Field[string] `json:"external_dimensional_price_group_id"`
}

func (r PlanVersionNewParamsReplacePricesPriceNewPlanTieredPackageWithMinimumPriceDimensionalPriceConfiguration) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Within each billing cycle, specifies the cadence at which invoices are produced.
// If unspecified, a single invoice is produced per billing cycle.
type PlanVersionNewParamsReplacePricesPriceNewPlanTieredPackageWithMinimumPriceInvoicingCycleConfiguration struct {
	// The duration of the billing period.
	Duration param.Field[int64] `json:"duration,required"`
	// The unit of billing period duration.
	DurationUnit param.Field[PlanVersionNewParamsReplacePricesPriceNewPlanTieredPackageWithMinimumPriceInvoicingCycleConfigurationDurationUnit] `json:"duration_unit,required"`
}

func (r PlanVersionNewParamsReplacePricesPriceNewPlanTieredPackageWithMinimumPriceInvoicingCycleConfiguration) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// The unit of billing period duration.
type PlanVersionNewParamsReplacePricesPriceNewPlanTieredPackageWithMinimumPriceInvoicingCycleConfigurationDurationUnit string

const (
	PlanVersionNewParamsReplacePricesPriceNewPlanTieredPackageWithMinimumPriceInvoicingCycleConfigurationDurationUnitDay   PlanVersionNewParamsReplacePricesPriceNewPlanTieredPackageWithMinimumPriceInvoicingCycleConfigurationDurationUnit = "day"
	PlanVersionNewParamsReplacePricesPriceNewPlanTieredPackageWithMinimumPriceInvoicingCycleConfigurationDurationUnitMonth PlanVersionNewParamsReplacePricesPriceNewPlanTieredPackageWithMinimumPriceInvoicingCycleConfigurationDurationUnit = "month"
)

func (r PlanVersionNewParamsReplacePricesPriceNewPlanTieredPackageWithMinimumPriceInvoicingCycleConfigurationDurationUnit) IsKnown() bool {
	switch r {
	case PlanVersionNewParamsReplacePricesPriceNewPlanTieredPackageWithMinimumPriceInvoicingCycleConfigurationDurationUnitDay, PlanVersionNewParamsReplacePricesPriceNewPlanTieredPackageWithMinimumPriceInvoicingCycleConfigurationDurationUnitMonth:
		return true
	}
	return false
}

type PlanVersionNewParamsReplacePricesPriceNewPlanMatrixWithAllocationPrice struct {
	// The cadence to bill for this price on.
	Cadence param.Field[PlanVersionNewParamsReplacePricesPriceNewPlanMatrixWithAllocationPriceCadence] `json:"cadence,required"`
	// The id of the item the price will be associated with.
	ItemID                     param.Field[string]                                                                                           `json:"item_id,required"`
	MatrixWithAllocationConfig param.Field[PlanVersionNewParamsReplacePricesPriceNewPlanMatrixWithAllocationPriceMatrixWithAllocationConfig] `json:"matrix_with_allocation_config,required"`
	ModelType                  param.Field[PlanVersionNewParamsReplacePricesPriceNewPlanMatrixWithAllocationPriceModelType]                  `json:"model_type,required"`
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
	BillingCycleConfiguration param.Field[PlanVersionNewParamsReplacePricesPriceNewPlanMatrixWithAllocationPriceBillingCycleConfiguration] `json:"billing_cycle_configuration"`
	// The per unit conversion rate of the price currency to the invoicing currency.
	ConversionRate param.Field[float64] `json:"conversion_rate"`
	// An ISO 4217 currency string, or custom pricing unit identifier, in which this
	// price is billed.
	Currency param.Field[string] `json:"currency"`
	// For dimensional price: specifies a price group and dimension values
	DimensionalPriceConfiguration param.Field[PlanVersionNewParamsReplacePricesPriceNewPlanMatrixWithAllocationPriceDimensionalPriceConfiguration] `json:"dimensional_price_configuration"`
	// An alias for the price.
	ExternalPriceID param.Field[string] `json:"external_price_id"`
	// If the Price represents a fixed cost, this represents the quantity of units
	// applied.
	FixedPriceQuantity param.Field[float64] `json:"fixed_price_quantity"`
	// The property used to group this price on an invoice
	InvoiceGroupingKey param.Field[string] `json:"invoice_grouping_key"`
	// Within each billing cycle, specifies the cadence at which invoices are produced.
	// If unspecified, a single invoice is produced per billing cycle.
	InvoicingCycleConfiguration param.Field[PlanVersionNewParamsReplacePricesPriceNewPlanMatrixWithAllocationPriceInvoicingCycleConfiguration] `json:"invoicing_cycle_configuration"`
	// User-specified key/value pairs for the resource. Individual keys can be removed
	// by setting the value to `null`, and the entire metadata mapping can be cleared
	// by setting `metadata` to `null`.
	Metadata param.Field[map[string]string] `json:"metadata"`
}

func (r PlanVersionNewParamsReplacePricesPriceNewPlanMatrixWithAllocationPrice) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r PlanVersionNewParamsReplacePricesPriceNewPlanMatrixWithAllocationPrice) implementsPlanVersionNewParamsReplacePricesPriceUnion() {
}

// The cadence to bill for this price on.
type PlanVersionNewParamsReplacePricesPriceNewPlanMatrixWithAllocationPriceCadence string

const (
	PlanVersionNewParamsReplacePricesPriceNewPlanMatrixWithAllocationPriceCadenceAnnual     PlanVersionNewParamsReplacePricesPriceNewPlanMatrixWithAllocationPriceCadence = "annual"
	PlanVersionNewParamsReplacePricesPriceNewPlanMatrixWithAllocationPriceCadenceSemiAnnual PlanVersionNewParamsReplacePricesPriceNewPlanMatrixWithAllocationPriceCadence = "semi_annual"
	PlanVersionNewParamsReplacePricesPriceNewPlanMatrixWithAllocationPriceCadenceMonthly    PlanVersionNewParamsReplacePricesPriceNewPlanMatrixWithAllocationPriceCadence = "monthly"
	PlanVersionNewParamsReplacePricesPriceNewPlanMatrixWithAllocationPriceCadenceQuarterly  PlanVersionNewParamsReplacePricesPriceNewPlanMatrixWithAllocationPriceCadence = "quarterly"
	PlanVersionNewParamsReplacePricesPriceNewPlanMatrixWithAllocationPriceCadenceOneTime    PlanVersionNewParamsReplacePricesPriceNewPlanMatrixWithAllocationPriceCadence = "one_time"
	PlanVersionNewParamsReplacePricesPriceNewPlanMatrixWithAllocationPriceCadenceCustom     PlanVersionNewParamsReplacePricesPriceNewPlanMatrixWithAllocationPriceCadence = "custom"
)

func (r PlanVersionNewParamsReplacePricesPriceNewPlanMatrixWithAllocationPriceCadence) IsKnown() bool {
	switch r {
	case PlanVersionNewParamsReplacePricesPriceNewPlanMatrixWithAllocationPriceCadenceAnnual, PlanVersionNewParamsReplacePricesPriceNewPlanMatrixWithAllocationPriceCadenceSemiAnnual, PlanVersionNewParamsReplacePricesPriceNewPlanMatrixWithAllocationPriceCadenceMonthly, PlanVersionNewParamsReplacePricesPriceNewPlanMatrixWithAllocationPriceCadenceQuarterly, PlanVersionNewParamsReplacePricesPriceNewPlanMatrixWithAllocationPriceCadenceOneTime, PlanVersionNewParamsReplacePricesPriceNewPlanMatrixWithAllocationPriceCadenceCustom:
		return true
	}
	return false
}

type PlanVersionNewParamsReplacePricesPriceNewPlanMatrixWithAllocationPriceMatrixWithAllocationConfig struct {
	// Allocation to be used to calculate the price
	Allocation param.Field[float64] `json:"allocation,required"`
	// Default per unit rate for any usage not bucketed into a specified matrix_value
	DefaultUnitAmount param.Field[string] `json:"default_unit_amount,required"`
	// One or two event property values to evaluate matrix groups by
	Dimensions param.Field[[]string] `json:"dimensions,required"`
	// Matrix values for specified matrix grouping keys
	MatrixValues param.Field[[]PlanVersionNewParamsReplacePricesPriceNewPlanMatrixWithAllocationPriceMatrixWithAllocationConfigMatrixValue] `json:"matrix_values,required"`
}

func (r PlanVersionNewParamsReplacePricesPriceNewPlanMatrixWithAllocationPriceMatrixWithAllocationConfig) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type PlanVersionNewParamsReplacePricesPriceNewPlanMatrixWithAllocationPriceMatrixWithAllocationConfigMatrixValue struct {
	// One or two matrix keys to filter usage to this Matrix value by. For example,
	// ["region", "tier"] could be used to filter cloud usage by a cloud region and an
	// instance tier.
	DimensionValues param.Field[[]string] `json:"dimension_values,required"`
	// Unit price for the specified dimension_values
	UnitAmount param.Field[string] `json:"unit_amount,required"`
}

func (r PlanVersionNewParamsReplacePricesPriceNewPlanMatrixWithAllocationPriceMatrixWithAllocationConfigMatrixValue) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type PlanVersionNewParamsReplacePricesPriceNewPlanMatrixWithAllocationPriceModelType string

const (
	PlanVersionNewParamsReplacePricesPriceNewPlanMatrixWithAllocationPriceModelTypeMatrixWithAllocation PlanVersionNewParamsReplacePricesPriceNewPlanMatrixWithAllocationPriceModelType = "matrix_with_allocation"
)

func (r PlanVersionNewParamsReplacePricesPriceNewPlanMatrixWithAllocationPriceModelType) IsKnown() bool {
	switch r {
	case PlanVersionNewParamsReplacePricesPriceNewPlanMatrixWithAllocationPriceModelTypeMatrixWithAllocation:
		return true
	}
	return false
}

// For custom cadence: specifies the duration of the billing period in days or
// months.
type PlanVersionNewParamsReplacePricesPriceNewPlanMatrixWithAllocationPriceBillingCycleConfiguration struct {
	// The duration of the billing period.
	Duration param.Field[int64] `json:"duration,required"`
	// The unit of billing period duration.
	DurationUnit param.Field[PlanVersionNewParamsReplacePricesPriceNewPlanMatrixWithAllocationPriceBillingCycleConfigurationDurationUnit] `json:"duration_unit,required"`
}

func (r PlanVersionNewParamsReplacePricesPriceNewPlanMatrixWithAllocationPriceBillingCycleConfiguration) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// The unit of billing period duration.
type PlanVersionNewParamsReplacePricesPriceNewPlanMatrixWithAllocationPriceBillingCycleConfigurationDurationUnit string

const (
	PlanVersionNewParamsReplacePricesPriceNewPlanMatrixWithAllocationPriceBillingCycleConfigurationDurationUnitDay   PlanVersionNewParamsReplacePricesPriceNewPlanMatrixWithAllocationPriceBillingCycleConfigurationDurationUnit = "day"
	PlanVersionNewParamsReplacePricesPriceNewPlanMatrixWithAllocationPriceBillingCycleConfigurationDurationUnitMonth PlanVersionNewParamsReplacePricesPriceNewPlanMatrixWithAllocationPriceBillingCycleConfigurationDurationUnit = "month"
)

func (r PlanVersionNewParamsReplacePricesPriceNewPlanMatrixWithAllocationPriceBillingCycleConfigurationDurationUnit) IsKnown() bool {
	switch r {
	case PlanVersionNewParamsReplacePricesPriceNewPlanMatrixWithAllocationPriceBillingCycleConfigurationDurationUnitDay, PlanVersionNewParamsReplacePricesPriceNewPlanMatrixWithAllocationPriceBillingCycleConfigurationDurationUnitMonth:
		return true
	}
	return false
}

// For dimensional price: specifies a price group and dimension values
type PlanVersionNewParamsReplacePricesPriceNewPlanMatrixWithAllocationPriceDimensionalPriceConfiguration struct {
	// The list of dimension values matching (in order) the dimensions of the price
	// group
	DimensionValues param.Field[[]string] `json:"dimension_values,required"`
	// The id of the dimensional price group to include this price in
	DimensionalPriceGroupID param.Field[string] `json:"dimensional_price_group_id"`
	// The external id of the dimensional price group to include this price in
	ExternalDimensionalPriceGroupID param.Field[string] `json:"external_dimensional_price_group_id"`
}

func (r PlanVersionNewParamsReplacePricesPriceNewPlanMatrixWithAllocationPriceDimensionalPriceConfiguration) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Within each billing cycle, specifies the cadence at which invoices are produced.
// If unspecified, a single invoice is produced per billing cycle.
type PlanVersionNewParamsReplacePricesPriceNewPlanMatrixWithAllocationPriceInvoicingCycleConfiguration struct {
	// The duration of the billing period.
	Duration param.Field[int64] `json:"duration,required"`
	// The unit of billing period duration.
	DurationUnit param.Field[PlanVersionNewParamsReplacePricesPriceNewPlanMatrixWithAllocationPriceInvoicingCycleConfigurationDurationUnit] `json:"duration_unit,required"`
}

func (r PlanVersionNewParamsReplacePricesPriceNewPlanMatrixWithAllocationPriceInvoicingCycleConfiguration) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// The unit of billing period duration.
type PlanVersionNewParamsReplacePricesPriceNewPlanMatrixWithAllocationPriceInvoicingCycleConfigurationDurationUnit string

const (
	PlanVersionNewParamsReplacePricesPriceNewPlanMatrixWithAllocationPriceInvoicingCycleConfigurationDurationUnitDay   PlanVersionNewParamsReplacePricesPriceNewPlanMatrixWithAllocationPriceInvoicingCycleConfigurationDurationUnit = "day"
	PlanVersionNewParamsReplacePricesPriceNewPlanMatrixWithAllocationPriceInvoicingCycleConfigurationDurationUnitMonth PlanVersionNewParamsReplacePricesPriceNewPlanMatrixWithAllocationPriceInvoicingCycleConfigurationDurationUnit = "month"
)

func (r PlanVersionNewParamsReplacePricesPriceNewPlanMatrixWithAllocationPriceInvoicingCycleConfigurationDurationUnit) IsKnown() bool {
	switch r {
	case PlanVersionNewParamsReplacePricesPriceNewPlanMatrixWithAllocationPriceInvoicingCycleConfigurationDurationUnitDay, PlanVersionNewParamsReplacePricesPriceNewPlanMatrixWithAllocationPriceInvoicingCycleConfigurationDurationUnitMonth:
		return true
	}
	return false
}

type PlanVersionNewParamsReplacePricesPriceNewPlanGroupedTieredPrice struct {
	// The cadence to bill for this price on.
	Cadence             param.Field[PlanVersionNewParamsReplacePricesPriceNewPlanGroupedTieredPriceCadence] `json:"cadence,required"`
	GroupedTieredConfig param.Field[map[string]interface{}]                                                 `json:"grouped_tiered_config,required"`
	// The id of the item the price will be associated with.
	ItemID    param.Field[string]                                                                   `json:"item_id,required"`
	ModelType param.Field[PlanVersionNewParamsReplacePricesPriceNewPlanGroupedTieredPriceModelType] `json:"model_type,required"`
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
	BillingCycleConfiguration param.Field[PlanVersionNewParamsReplacePricesPriceNewPlanGroupedTieredPriceBillingCycleConfiguration] `json:"billing_cycle_configuration"`
	// The per unit conversion rate of the price currency to the invoicing currency.
	ConversionRate param.Field[float64] `json:"conversion_rate"`
	// An ISO 4217 currency string, or custom pricing unit identifier, in which this
	// price is billed.
	Currency param.Field[string] `json:"currency"`
	// For dimensional price: specifies a price group and dimension values
	DimensionalPriceConfiguration param.Field[PlanVersionNewParamsReplacePricesPriceNewPlanGroupedTieredPriceDimensionalPriceConfiguration] `json:"dimensional_price_configuration"`
	// An alias for the price.
	ExternalPriceID param.Field[string] `json:"external_price_id"`
	// If the Price represents a fixed cost, this represents the quantity of units
	// applied.
	FixedPriceQuantity param.Field[float64] `json:"fixed_price_quantity"`
	// The property used to group this price on an invoice
	InvoiceGroupingKey param.Field[string] `json:"invoice_grouping_key"`
	// Within each billing cycle, specifies the cadence at which invoices are produced.
	// If unspecified, a single invoice is produced per billing cycle.
	InvoicingCycleConfiguration param.Field[PlanVersionNewParamsReplacePricesPriceNewPlanGroupedTieredPriceInvoicingCycleConfiguration] `json:"invoicing_cycle_configuration"`
	// User-specified key/value pairs for the resource. Individual keys can be removed
	// by setting the value to `null`, and the entire metadata mapping can be cleared
	// by setting `metadata` to `null`.
	Metadata param.Field[map[string]string] `json:"metadata"`
}

func (r PlanVersionNewParamsReplacePricesPriceNewPlanGroupedTieredPrice) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r PlanVersionNewParamsReplacePricesPriceNewPlanGroupedTieredPrice) implementsPlanVersionNewParamsReplacePricesPriceUnion() {
}

// The cadence to bill for this price on.
type PlanVersionNewParamsReplacePricesPriceNewPlanGroupedTieredPriceCadence string

const (
	PlanVersionNewParamsReplacePricesPriceNewPlanGroupedTieredPriceCadenceAnnual     PlanVersionNewParamsReplacePricesPriceNewPlanGroupedTieredPriceCadence = "annual"
	PlanVersionNewParamsReplacePricesPriceNewPlanGroupedTieredPriceCadenceSemiAnnual PlanVersionNewParamsReplacePricesPriceNewPlanGroupedTieredPriceCadence = "semi_annual"
	PlanVersionNewParamsReplacePricesPriceNewPlanGroupedTieredPriceCadenceMonthly    PlanVersionNewParamsReplacePricesPriceNewPlanGroupedTieredPriceCadence = "monthly"
	PlanVersionNewParamsReplacePricesPriceNewPlanGroupedTieredPriceCadenceQuarterly  PlanVersionNewParamsReplacePricesPriceNewPlanGroupedTieredPriceCadence = "quarterly"
	PlanVersionNewParamsReplacePricesPriceNewPlanGroupedTieredPriceCadenceOneTime    PlanVersionNewParamsReplacePricesPriceNewPlanGroupedTieredPriceCadence = "one_time"
	PlanVersionNewParamsReplacePricesPriceNewPlanGroupedTieredPriceCadenceCustom     PlanVersionNewParamsReplacePricesPriceNewPlanGroupedTieredPriceCadence = "custom"
)

func (r PlanVersionNewParamsReplacePricesPriceNewPlanGroupedTieredPriceCadence) IsKnown() bool {
	switch r {
	case PlanVersionNewParamsReplacePricesPriceNewPlanGroupedTieredPriceCadenceAnnual, PlanVersionNewParamsReplacePricesPriceNewPlanGroupedTieredPriceCadenceSemiAnnual, PlanVersionNewParamsReplacePricesPriceNewPlanGroupedTieredPriceCadenceMonthly, PlanVersionNewParamsReplacePricesPriceNewPlanGroupedTieredPriceCadenceQuarterly, PlanVersionNewParamsReplacePricesPriceNewPlanGroupedTieredPriceCadenceOneTime, PlanVersionNewParamsReplacePricesPriceNewPlanGroupedTieredPriceCadenceCustom:
		return true
	}
	return false
}

type PlanVersionNewParamsReplacePricesPriceNewPlanGroupedTieredPriceModelType string

const (
	PlanVersionNewParamsReplacePricesPriceNewPlanGroupedTieredPriceModelTypeGroupedTiered PlanVersionNewParamsReplacePricesPriceNewPlanGroupedTieredPriceModelType = "grouped_tiered"
)

func (r PlanVersionNewParamsReplacePricesPriceNewPlanGroupedTieredPriceModelType) IsKnown() bool {
	switch r {
	case PlanVersionNewParamsReplacePricesPriceNewPlanGroupedTieredPriceModelTypeGroupedTiered:
		return true
	}
	return false
}

// For custom cadence: specifies the duration of the billing period in days or
// months.
type PlanVersionNewParamsReplacePricesPriceNewPlanGroupedTieredPriceBillingCycleConfiguration struct {
	// The duration of the billing period.
	Duration param.Field[int64] `json:"duration,required"`
	// The unit of billing period duration.
	DurationUnit param.Field[PlanVersionNewParamsReplacePricesPriceNewPlanGroupedTieredPriceBillingCycleConfigurationDurationUnit] `json:"duration_unit,required"`
}

func (r PlanVersionNewParamsReplacePricesPriceNewPlanGroupedTieredPriceBillingCycleConfiguration) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// The unit of billing period duration.
type PlanVersionNewParamsReplacePricesPriceNewPlanGroupedTieredPriceBillingCycleConfigurationDurationUnit string

const (
	PlanVersionNewParamsReplacePricesPriceNewPlanGroupedTieredPriceBillingCycleConfigurationDurationUnitDay   PlanVersionNewParamsReplacePricesPriceNewPlanGroupedTieredPriceBillingCycleConfigurationDurationUnit = "day"
	PlanVersionNewParamsReplacePricesPriceNewPlanGroupedTieredPriceBillingCycleConfigurationDurationUnitMonth PlanVersionNewParamsReplacePricesPriceNewPlanGroupedTieredPriceBillingCycleConfigurationDurationUnit = "month"
)

func (r PlanVersionNewParamsReplacePricesPriceNewPlanGroupedTieredPriceBillingCycleConfigurationDurationUnit) IsKnown() bool {
	switch r {
	case PlanVersionNewParamsReplacePricesPriceNewPlanGroupedTieredPriceBillingCycleConfigurationDurationUnitDay, PlanVersionNewParamsReplacePricesPriceNewPlanGroupedTieredPriceBillingCycleConfigurationDurationUnitMonth:
		return true
	}
	return false
}

// For dimensional price: specifies a price group and dimension values
type PlanVersionNewParamsReplacePricesPriceNewPlanGroupedTieredPriceDimensionalPriceConfiguration struct {
	// The list of dimension values matching (in order) the dimensions of the price
	// group
	DimensionValues param.Field[[]string] `json:"dimension_values,required"`
	// The id of the dimensional price group to include this price in
	DimensionalPriceGroupID param.Field[string] `json:"dimensional_price_group_id"`
	// The external id of the dimensional price group to include this price in
	ExternalDimensionalPriceGroupID param.Field[string] `json:"external_dimensional_price_group_id"`
}

func (r PlanVersionNewParamsReplacePricesPriceNewPlanGroupedTieredPriceDimensionalPriceConfiguration) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Within each billing cycle, specifies the cadence at which invoices are produced.
// If unspecified, a single invoice is produced per billing cycle.
type PlanVersionNewParamsReplacePricesPriceNewPlanGroupedTieredPriceInvoicingCycleConfiguration struct {
	// The duration of the billing period.
	Duration param.Field[int64] `json:"duration,required"`
	// The unit of billing period duration.
	DurationUnit param.Field[PlanVersionNewParamsReplacePricesPriceNewPlanGroupedTieredPriceInvoicingCycleConfigurationDurationUnit] `json:"duration_unit,required"`
}

func (r PlanVersionNewParamsReplacePricesPriceNewPlanGroupedTieredPriceInvoicingCycleConfiguration) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// The unit of billing period duration.
type PlanVersionNewParamsReplacePricesPriceNewPlanGroupedTieredPriceInvoicingCycleConfigurationDurationUnit string

const (
	PlanVersionNewParamsReplacePricesPriceNewPlanGroupedTieredPriceInvoicingCycleConfigurationDurationUnitDay   PlanVersionNewParamsReplacePricesPriceNewPlanGroupedTieredPriceInvoicingCycleConfigurationDurationUnit = "day"
	PlanVersionNewParamsReplacePricesPriceNewPlanGroupedTieredPriceInvoicingCycleConfigurationDurationUnitMonth PlanVersionNewParamsReplacePricesPriceNewPlanGroupedTieredPriceInvoicingCycleConfigurationDurationUnit = "month"
)

func (r PlanVersionNewParamsReplacePricesPriceNewPlanGroupedTieredPriceInvoicingCycleConfigurationDurationUnit) IsKnown() bool {
	switch r {
	case PlanVersionNewParamsReplacePricesPriceNewPlanGroupedTieredPriceInvoicingCycleConfigurationDurationUnitDay, PlanVersionNewParamsReplacePricesPriceNewPlanGroupedTieredPriceInvoicingCycleConfigurationDurationUnitMonth:
		return true
	}
	return false
}

// The cadence to bill for this price on.
type PlanVersionNewParamsReplacePricesPriceCadence string

const (
	PlanVersionNewParamsReplacePricesPriceCadenceAnnual     PlanVersionNewParamsReplacePricesPriceCadence = "annual"
	PlanVersionNewParamsReplacePricesPriceCadenceSemiAnnual PlanVersionNewParamsReplacePricesPriceCadence = "semi_annual"
	PlanVersionNewParamsReplacePricesPriceCadenceMonthly    PlanVersionNewParamsReplacePricesPriceCadence = "monthly"
	PlanVersionNewParamsReplacePricesPriceCadenceQuarterly  PlanVersionNewParamsReplacePricesPriceCadence = "quarterly"
	PlanVersionNewParamsReplacePricesPriceCadenceOneTime    PlanVersionNewParamsReplacePricesPriceCadence = "one_time"
	PlanVersionNewParamsReplacePricesPriceCadenceCustom     PlanVersionNewParamsReplacePricesPriceCadence = "custom"
)

func (r PlanVersionNewParamsReplacePricesPriceCadence) IsKnown() bool {
	switch r {
	case PlanVersionNewParamsReplacePricesPriceCadenceAnnual, PlanVersionNewParamsReplacePricesPriceCadenceSemiAnnual, PlanVersionNewParamsReplacePricesPriceCadenceMonthly, PlanVersionNewParamsReplacePricesPriceCadenceQuarterly, PlanVersionNewParamsReplacePricesPriceCadenceOneTime, PlanVersionNewParamsReplacePricesPriceCadenceCustom:
		return true
	}
	return false
}

type PlanVersionNewParamsReplacePricesPriceModelType string

const (
	PlanVersionNewParamsReplacePricesPriceModelTypeUnit                            PlanVersionNewParamsReplacePricesPriceModelType = "unit"
	PlanVersionNewParamsReplacePricesPriceModelTypePackage                         PlanVersionNewParamsReplacePricesPriceModelType = "package"
	PlanVersionNewParamsReplacePricesPriceModelTypeMatrix                          PlanVersionNewParamsReplacePricesPriceModelType = "matrix"
	PlanVersionNewParamsReplacePricesPriceModelTypeTiered                          PlanVersionNewParamsReplacePricesPriceModelType = "tiered"
	PlanVersionNewParamsReplacePricesPriceModelTypeTieredBps                       PlanVersionNewParamsReplacePricesPriceModelType = "tiered_bps"
	PlanVersionNewParamsReplacePricesPriceModelTypeBps                             PlanVersionNewParamsReplacePricesPriceModelType = "bps"
	PlanVersionNewParamsReplacePricesPriceModelTypeBulkBps                         PlanVersionNewParamsReplacePricesPriceModelType = "bulk_bps"
	PlanVersionNewParamsReplacePricesPriceModelTypeBulk                            PlanVersionNewParamsReplacePricesPriceModelType = "bulk"
	PlanVersionNewParamsReplacePricesPriceModelTypeThresholdTotalAmount            PlanVersionNewParamsReplacePricesPriceModelType = "threshold_total_amount"
	PlanVersionNewParamsReplacePricesPriceModelTypeTieredPackage                   PlanVersionNewParamsReplacePricesPriceModelType = "tiered_package"
	PlanVersionNewParamsReplacePricesPriceModelTypeTieredWithMinimum               PlanVersionNewParamsReplacePricesPriceModelType = "tiered_with_minimum"
	PlanVersionNewParamsReplacePricesPriceModelTypeUnitWithPercent                 PlanVersionNewParamsReplacePricesPriceModelType = "unit_with_percent"
	PlanVersionNewParamsReplacePricesPriceModelTypePackageWithAllocation           PlanVersionNewParamsReplacePricesPriceModelType = "package_with_allocation"
	PlanVersionNewParamsReplacePricesPriceModelTypeTieredWithProration             PlanVersionNewParamsReplacePricesPriceModelType = "tiered_with_proration"
	PlanVersionNewParamsReplacePricesPriceModelTypeUnitWithProration               PlanVersionNewParamsReplacePricesPriceModelType = "unit_with_proration"
	PlanVersionNewParamsReplacePricesPriceModelTypeGroupedAllocation               PlanVersionNewParamsReplacePricesPriceModelType = "grouped_allocation"
	PlanVersionNewParamsReplacePricesPriceModelTypeGroupedWithProratedMinimum      PlanVersionNewParamsReplacePricesPriceModelType = "grouped_with_prorated_minimum"
	PlanVersionNewParamsReplacePricesPriceModelTypeGroupedWithMeteredMinimum       PlanVersionNewParamsReplacePricesPriceModelType = "grouped_with_metered_minimum"
	PlanVersionNewParamsReplacePricesPriceModelTypeMatrixWithDisplayName           PlanVersionNewParamsReplacePricesPriceModelType = "matrix_with_display_name"
	PlanVersionNewParamsReplacePricesPriceModelTypeBulkWithProration               PlanVersionNewParamsReplacePricesPriceModelType = "bulk_with_proration"
	PlanVersionNewParamsReplacePricesPriceModelTypeGroupedTieredPackage            PlanVersionNewParamsReplacePricesPriceModelType = "grouped_tiered_package"
	PlanVersionNewParamsReplacePricesPriceModelTypeMaxGroupTieredPackage           PlanVersionNewParamsReplacePricesPriceModelType = "max_group_tiered_package"
	PlanVersionNewParamsReplacePricesPriceModelTypeScalableMatrixWithUnitPricing   PlanVersionNewParamsReplacePricesPriceModelType = "scalable_matrix_with_unit_pricing"
	PlanVersionNewParamsReplacePricesPriceModelTypeScalableMatrixWithTieredPricing PlanVersionNewParamsReplacePricesPriceModelType = "scalable_matrix_with_tiered_pricing"
	PlanVersionNewParamsReplacePricesPriceModelTypeCumulativeGroupedBulk           PlanVersionNewParamsReplacePricesPriceModelType = "cumulative_grouped_bulk"
	PlanVersionNewParamsReplacePricesPriceModelTypeTieredPackageWithMinimum        PlanVersionNewParamsReplacePricesPriceModelType = "tiered_package_with_minimum"
	PlanVersionNewParamsReplacePricesPriceModelTypeMatrixWithAllocation            PlanVersionNewParamsReplacePricesPriceModelType = "matrix_with_allocation"
	PlanVersionNewParamsReplacePricesPriceModelTypeGroupedTiered                   PlanVersionNewParamsReplacePricesPriceModelType = "grouped_tiered"
)

func (r PlanVersionNewParamsReplacePricesPriceModelType) IsKnown() bool {
	switch r {
	case PlanVersionNewParamsReplacePricesPriceModelTypeUnit, PlanVersionNewParamsReplacePricesPriceModelTypePackage, PlanVersionNewParamsReplacePricesPriceModelTypeMatrix, PlanVersionNewParamsReplacePricesPriceModelTypeTiered, PlanVersionNewParamsReplacePricesPriceModelTypeTieredBps, PlanVersionNewParamsReplacePricesPriceModelTypeBps, PlanVersionNewParamsReplacePricesPriceModelTypeBulkBps, PlanVersionNewParamsReplacePricesPriceModelTypeBulk, PlanVersionNewParamsReplacePricesPriceModelTypeThresholdTotalAmount, PlanVersionNewParamsReplacePricesPriceModelTypeTieredPackage, PlanVersionNewParamsReplacePricesPriceModelTypeTieredWithMinimum, PlanVersionNewParamsReplacePricesPriceModelTypeUnitWithPercent, PlanVersionNewParamsReplacePricesPriceModelTypePackageWithAllocation, PlanVersionNewParamsReplacePricesPriceModelTypeTieredWithProration, PlanVersionNewParamsReplacePricesPriceModelTypeUnitWithProration, PlanVersionNewParamsReplacePricesPriceModelTypeGroupedAllocation, PlanVersionNewParamsReplacePricesPriceModelTypeGroupedWithProratedMinimum, PlanVersionNewParamsReplacePricesPriceModelTypeGroupedWithMeteredMinimum, PlanVersionNewParamsReplacePricesPriceModelTypeMatrixWithDisplayName, PlanVersionNewParamsReplacePricesPriceModelTypeBulkWithProration, PlanVersionNewParamsReplacePricesPriceModelTypeGroupedTieredPackage, PlanVersionNewParamsReplacePricesPriceModelTypeMaxGroupTieredPackage, PlanVersionNewParamsReplacePricesPriceModelTypeScalableMatrixWithUnitPricing, PlanVersionNewParamsReplacePricesPriceModelTypeScalableMatrixWithTieredPricing, PlanVersionNewParamsReplacePricesPriceModelTypeCumulativeGroupedBulk, PlanVersionNewParamsReplacePricesPriceModelTypeTieredPackageWithMinimum, PlanVersionNewParamsReplacePricesPriceModelTypeMatrixWithAllocation, PlanVersionNewParamsReplacePricesPriceModelTypeGroupedTiered:
		return true
	}
	return false
}
