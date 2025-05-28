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

// PlanExternalPlanIDVersionService contains methods and other services that help
// with interacting with the orb API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewPlanExternalPlanIDVersionService] method instead.
type PlanExternalPlanIDVersionService struct {
	Options []option.RequestOption
}

// NewPlanExternalPlanIDVersionService generates a new service that applies the
// given options to each request. These options are applied after the parent
// client's options (if there is one), and before any request-specific options.
func NewPlanExternalPlanIDVersionService(opts ...option.RequestOption) (r *PlanExternalPlanIDVersionService) {
	r = &PlanExternalPlanIDVersionService{}
	r.Options = opts
	return
}

// This API endpoint is in beta and its interface may change. It is recommended for
// use only in test mode.
//
// This endpoint allows the creation of a new plan version for an existing plan.
func (r *PlanExternalPlanIDVersionService) New(ctx context.Context, externalPlanID string, body PlanExternalPlanIDVersionNewParams, opts ...option.RequestOption) (res *PlanExternalPlanIDVersionNewResponse, err error) {
	opts = append(r.Options[:], opts...)
	if externalPlanID == "" {
		err = errors.New("missing required external_plan_id parameter")
		return
	}
	path := fmt.Sprintf("plans/external_plan_id/%s/versions", externalPlanID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return
}

// This API endpoint is in beta and its interface may change. It is recommended for
// use only in test mode.
//
// This endpoint is used to fetch a plan version. It returns the phases, prices,
// and adjustments present on this version of the plan.
func (r *PlanExternalPlanIDVersionService) Get(ctx context.Context, externalPlanID string, version string, opts ...option.RequestOption) (res *PlanExternalPlanIDVersionGetResponse, err error) {
	opts = append(r.Options[:], opts...)
	if externalPlanID == "" {
		err = errors.New("missing required external_plan_id parameter")
		return
	}
	if version == "" {
		err = errors.New("missing required version parameter")
		return
	}
	path := fmt.Sprintf("plans/external_plan_id/%s/versions/%s", externalPlanID, version)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

// The PlanVersion resource represents the prices and adjustments present on a
// specific version of a plan.
type PlanExternalPlanIDVersionNewResponse struct {
	// Adjustments for this plan. If the plan has phases, this includes adjustments
	// across all phases of the plan.
	Adjustments []PlanExternalPlanIDVersionNewResponseAdjustment `json:"adjustments,required"`
	CreatedAt   time.Time                                        `json:"created_at,required" format:"date-time"`
	PlanPhases  []PlanExternalPlanIDVersionNewResponsePlanPhase  `json:"plan_phases,required,nullable"`
	// Prices for this plan. If the plan has phases, this includes prices across all
	// phases of the plan.
	Prices  []Price                                  `json:"prices,required"`
	Version int64                                    `json:"version,required"`
	JSON    planExternalPlanIDVersionNewResponseJSON `json:"-"`
}

// planExternalPlanIDVersionNewResponseJSON contains the JSON metadata for the
// struct [PlanExternalPlanIDVersionNewResponse]
type planExternalPlanIDVersionNewResponseJSON struct {
	Adjustments apijson.Field
	CreatedAt   apijson.Field
	PlanPhases  apijson.Field
	Prices      apijson.Field
	Version     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *PlanExternalPlanIDVersionNewResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r planExternalPlanIDVersionNewResponseJSON) RawJSON() string {
	return r.raw
}

type PlanExternalPlanIDVersionNewResponseAdjustment struct {
	ID             string                                                        `json:"id,required"`
	AdjustmentType PlanExternalPlanIDVersionNewResponseAdjustmentsAdjustmentType `json:"adjustment_type,required"`
	// This field can have the runtime type of [[]string].
	AppliesToPriceIDs interface{} `json:"applies_to_price_ids,required"`
	// This field can have the runtime type of
	// [[]PlanExternalPlanIDVersionNewResponseAdjustmentsPlanPhaseUsageDiscountAdjustmentFilter],
	// [[]PlanExternalPlanIDVersionNewResponseAdjustmentsPlanPhaseAmountDiscountAdjustmentFilter],
	// [[]PlanExternalPlanIDVersionNewResponseAdjustmentsPlanPhasePercentageDiscountAdjustmentFilter],
	// [[]PlanExternalPlanIDVersionNewResponseAdjustmentsPlanPhaseMinimumAdjustmentFilter],
	// [[]PlanExternalPlanIDVersionNewResponseAdjustmentsPlanPhaseMaximumAdjustmentFilter].
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
	UsageDiscount float64                                            `json:"usage_discount"`
	JSON          planExternalPlanIDVersionNewResponseAdjustmentJSON `json:"-"`
	union         PlanExternalPlanIDVersionNewResponseAdjustmentsUnion
}

// planExternalPlanIDVersionNewResponseAdjustmentJSON contains the JSON metadata
// for the struct [PlanExternalPlanIDVersionNewResponseAdjustment]
type planExternalPlanIDVersionNewResponseAdjustmentJSON struct {
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

func (r planExternalPlanIDVersionNewResponseAdjustmentJSON) RawJSON() string {
	return r.raw
}

func (r *PlanExternalPlanIDVersionNewResponseAdjustment) UnmarshalJSON(data []byte) (err error) {
	*r = PlanExternalPlanIDVersionNewResponseAdjustment{}
	err = apijson.UnmarshalRoot(data, &r.union)
	if err != nil {
		return err
	}
	return apijson.Port(r.union, &r)
}

// AsUnion returns a [PlanExternalPlanIDVersionNewResponseAdjustmentsUnion]
// interface which you can cast to the specific types for more type safety.
//
// Possible runtime types of the union are
// [PlanExternalPlanIDVersionNewResponseAdjustmentsPlanPhaseUsageDiscountAdjustment],
// [PlanExternalPlanIDVersionNewResponseAdjustmentsPlanPhaseAmountDiscountAdjustment],
// [PlanExternalPlanIDVersionNewResponseAdjustmentsPlanPhasePercentageDiscountAdjustment],
// [PlanExternalPlanIDVersionNewResponseAdjustmentsPlanPhaseMinimumAdjustment],
// [PlanExternalPlanIDVersionNewResponseAdjustmentsPlanPhaseMaximumAdjustment].
func (r PlanExternalPlanIDVersionNewResponseAdjustment) AsUnion() PlanExternalPlanIDVersionNewResponseAdjustmentsUnion {
	return r.union
}

// Union satisfied by
// [PlanExternalPlanIDVersionNewResponseAdjustmentsPlanPhaseUsageDiscountAdjustment],
// [PlanExternalPlanIDVersionNewResponseAdjustmentsPlanPhaseAmountDiscountAdjustment],
// [PlanExternalPlanIDVersionNewResponseAdjustmentsPlanPhasePercentageDiscountAdjustment],
// [PlanExternalPlanIDVersionNewResponseAdjustmentsPlanPhaseMinimumAdjustment] or
// [PlanExternalPlanIDVersionNewResponseAdjustmentsPlanPhaseMaximumAdjustment].
type PlanExternalPlanIDVersionNewResponseAdjustmentsUnion interface {
	implementsPlanExternalPlanIDVersionNewResponseAdjustment()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*PlanExternalPlanIDVersionNewResponseAdjustmentsUnion)(nil)).Elem(),
		"adjustment_type",
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(PlanExternalPlanIDVersionNewResponseAdjustmentsPlanPhaseUsageDiscountAdjustment{}),
			DiscriminatorValue: "usage_discount",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(PlanExternalPlanIDVersionNewResponseAdjustmentsPlanPhaseAmountDiscountAdjustment{}),
			DiscriminatorValue: "amount_discount",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(PlanExternalPlanIDVersionNewResponseAdjustmentsPlanPhasePercentageDiscountAdjustment{}),
			DiscriminatorValue: "percentage_discount",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(PlanExternalPlanIDVersionNewResponseAdjustmentsPlanPhaseMinimumAdjustment{}),
			DiscriminatorValue: "minimum",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(PlanExternalPlanIDVersionNewResponseAdjustmentsPlanPhaseMaximumAdjustment{}),
			DiscriminatorValue: "maximum",
		},
	)
}

type PlanExternalPlanIDVersionNewResponseAdjustmentsPlanPhaseUsageDiscountAdjustment struct {
	ID             string                                                                                        `json:"id,required"`
	AdjustmentType PlanExternalPlanIDVersionNewResponseAdjustmentsPlanPhaseUsageDiscountAdjustmentAdjustmentType `json:"adjustment_type,required"`
	// The price IDs that this adjustment applies to.
	//
	// Deprecated: deprecated
	AppliesToPriceIDs []string `json:"applies_to_price_ids,required"`
	// The filters that determine which prices to apply this adjustment to.
	Filters []PlanExternalPlanIDVersionNewResponseAdjustmentsPlanPhaseUsageDiscountAdjustmentFilter `json:"filters,required"`
	// True for adjustments that apply to an entire invocice, false for adjustments
	// that apply to only one price.
	IsInvoiceLevel bool `json:"is_invoice_level,required"`
	// The plan phase in which this adjustment is active.
	PlanPhaseOrder int64 `json:"plan_phase_order,required,nullable"`
	// The reason for the adjustment.
	Reason string `json:"reason,required,nullable"`
	// The number of usage units by which to discount the price this adjustment applies
	// to in a given billing period.
	UsageDiscount float64                                                                             `json:"usage_discount,required"`
	JSON          planExternalPlanIDVersionNewResponseAdjustmentsPlanPhaseUsageDiscountAdjustmentJSON `json:"-"`
}

// planExternalPlanIDVersionNewResponseAdjustmentsPlanPhaseUsageDiscountAdjustmentJSON
// contains the JSON metadata for the struct
// [PlanExternalPlanIDVersionNewResponseAdjustmentsPlanPhaseUsageDiscountAdjustment]
type planExternalPlanIDVersionNewResponseAdjustmentsPlanPhaseUsageDiscountAdjustmentJSON struct {
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

func (r *PlanExternalPlanIDVersionNewResponseAdjustmentsPlanPhaseUsageDiscountAdjustment) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r planExternalPlanIDVersionNewResponseAdjustmentsPlanPhaseUsageDiscountAdjustmentJSON) RawJSON() string {
	return r.raw
}

func (r PlanExternalPlanIDVersionNewResponseAdjustmentsPlanPhaseUsageDiscountAdjustment) implementsPlanExternalPlanIDVersionNewResponseAdjustment() {
}

type PlanExternalPlanIDVersionNewResponseAdjustmentsPlanPhaseUsageDiscountAdjustmentAdjustmentType string

const (
	PlanExternalPlanIDVersionNewResponseAdjustmentsPlanPhaseUsageDiscountAdjustmentAdjustmentTypeUsageDiscount PlanExternalPlanIDVersionNewResponseAdjustmentsPlanPhaseUsageDiscountAdjustmentAdjustmentType = "usage_discount"
)

func (r PlanExternalPlanIDVersionNewResponseAdjustmentsPlanPhaseUsageDiscountAdjustmentAdjustmentType) IsKnown() bool {
	switch r {
	case PlanExternalPlanIDVersionNewResponseAdjustmentsPlanPhaseUsageDiscountAdjustmentAdjustmentTypeUsageDiscount:
		return true
	}
	return false
}

type PlanExternalPlanIDVersionNewResponseAdjustmentsPlanPhaseUsageDiscountAdjustmentFilter struct {
	// The property of the price to filter on.
	Field PlanExternalPlanIDVersionNewResponseAdjustmentsPlanPhaseUsageDiscountAdjustmentFiltersField `json:"field,required"`
	// Should prices that match the filter be included or excluded.
	Operator PlanExternalPlanIDVersionNewResponseAdjustmentsPlanPhaseUsageDiscountAdjustmentFiltersOperator `json:"operator,required"`
	// The IDs or values that match this filter.
	Values []string                                                                                  `json:"values,required"`
	JSON   planExternalPlanIDVersionNewResponseAdjustmentsPlanPhaseUsageDiscountAdjustmentFilterJSON `json:"-"`
}

// planExternalPlanIDVersionNewResponseAdjustmentsPlanPhaseUsageDiscountAdjustmentFilterJSON
// contains the JSON metadata for the struct
// [PlanExternalPlanIDVersionNewResponseAdjustmentsPlanPhaseUsageDiscountAdjustmentFilter]
type planExternalPlanIDVersionNewResponseAdjustmentsPlanPhaseUsageDiscountAdjustmentFilterJSON struct {
	Field       apijson.Field
	Operator    apijson.Field
	Values      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *PlanExternalPlanIDVersionNewResponseAdjustmentsPlanPhaseUsageDiscountAdjustmentFilter) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r planExternalPlanIDVersionNewResponseAdjustmentsPlanPhaseUsageDiscountAdjustmentFilterJSON) RawJSON() string {
	return r.raw
}

// The property of the price to filter on.
type PlanExternalPlanIDVersionNewResponseAdjustmentsPlanPhaseUsageDiscountAdjustmentFiltersField string

const (
	PlanExternalPlanIDVersionNewResponseAdjustmentsPlanPhaseUsageDiscountAdjustmentFiltersFieldPriceID       PlanExternalPlanIDVersionNewResponseAdjustmentsPlanPhaseUsageDiscountAdjustmentFiltersField = "price_id"
	PlanExternalPlanIDVersionNewResponseAdjustmentsPlanPhaseUsageDiscountAdjustmentFiltersFieldItemID        PlanExternalPlanIDVersionNewResponseAdjustmentsPlanPhaseUsageDiscountAdjustmentFiltersField = "item_id"
	PlanExternalPlanIDVersionNewResponseAdjustmentsPlanPhaseUsageDiscountAdjustmentFiltersFieldPriceType     PlanExternalPlanIDVersionNewResponseAdjustmentsPlanPhaseUsageDiscountAdjustmentFiltersField = "price_type"
	PlanExternalPlanIDVersionNewResponseAdjustmentsPlanPhaseUsageDiscountAdjustmentFiltersFieldCurrency      PlanExternalPlanIDVersionNewResponseAdjustmentsPlanPhaseUsageDiscountAdjustmentFiltersField = "currency"
	PlanExternalPlanIDVersionNewResponseAdjustmentsPlanPhaseUsageDiscountAdjustmentFiltersFieldPricingUnitID PlanExternalPlanIDVersionNewResponseAdjustmentsPlanPhaseUsageDiscountAdjustmentFiltersField = "pricing_unit_id"
)

func (r PlanExternalPlanIDVersionNewResponseAdjustmentsPlanPhaseUsageDiscountAdjustmentFiltersField) IsKnown() bool {
	switch r {
	case PlanExternalPlanIDVersionNewResponseAdjustmentsPlanPhaseUsageDiscountAdjustmentFiltersFieldPriceID, PlanExternalPlanIDVersionNewResponseAdjustmentsPlanPhaseUsageDiscountAdjustmentFiltersFieldItemID, PlanExternalPlanIDVersionNewResponseAdjustmentsPlanPhaseUsageDiscountAdjustmentFiltersFieldPriceType, PlanExternalPlanIDVersionNewResponseAdjustmentsPlanPhaseUsageDiscountAdjustmentFiltersFieldCurrency, PlanExternalPlanIDVersionNewResponseAdjustmentsPlanPhaseUsageDiscountAdjustmentFiltersFieldPricingUnitID:
		return true
	}
	return false
}

// Should prices that match the filter be included or excluded.
type PlanExternalPlanIDVersionNewResponseAdjustmentsPlanPhaseUsageDiscountAdjustmentFiltersOperator string

const (
	PlanExternalPlanIDVersionNewResponseAdjustmentsPlanPhaseUsageDiscountAdjustmentFiltersOperatorIncludes PlanExternalPlanIDVersionNewResponseAdjustmentsPlanPhaseUsageDiscountAdjustmentFiltersOperator = "includes"
	PlanExternalPlanIDVersionNewResponseAdjustmentsPlanPhaseUsageDiscountAdjustmentFiltersOperatorExcludes PlanExternalPlanIDVersionNewResponseAdjustmentsPlanPhaseUsageDiscountAdjustmentFiltersOperator = "excludes"
)

func (r PlanExternalPlanIDVersionNewResponseAdjustmentsPlanPhaseUsageDiscountAdjustmentFiltersOperator) IsKnown() bool {
	switch r {
	case PlanExternalPlanIDVersionNewResponseAdjustmentsPlanPhaseUsageDiscountAdjustmentFiltersOperatorIncludes, PlanExternalPlanIDVersionNewResponseAdjustmentsPlanPhaseUsageDiscountAdjustmentFiltersOperatorExcludes:
		return true
	}
	return false
}

type PlanExternalPlanIDVersionNewResponseAdjustmentsPlanPhaseAmountDiscountAdjustment struct {
	ID             string                                                                                         `json:"id,required"`
	AdjustmentType PlanExternalPlanIDVersionNewResponseAdjustmentsPlanPhaseAmountDiscountAdjustmentAdjustmentType `json:"adjustment_type,required"`
	// The amount by which to discount the prices this adjustment applies to in a given
	// billing period.
	AmountDiscount string `json:"amount_discount,required"`
	// The price IDs that this adjustment applies to.
	//
	// Deprecated: deprecated
	AppliesToPriceIDs []string `json:"applies_to_price_ids,required"`
	// The filters that determine which prices to apply this adjustment to.
	Filters []PlanExternalPlanIDVersionNewResponseAdjustmentsPlanPhaseAmountDiscountAdjustmentFilter `json:"filters,required"`
	// True for adjustments that apply to an entire invocice, false for adjustments
	// that apply to only one price.
	IsInvoiceLevel bool `json:"is_invoice_level,required"`
	// The plan phase in which this adjustment is active.
	PlanPhaseOrder int64 `json:"plan_phase_order,required,nullable"`
	// The reason for the adjustment.
	Reason string                                                                               `json:"reason,required,nullable"`
	JSON   planExternalPlanIDVersionNewResponseAdjustmentsPlanPhaseAmountDiscountAdjustmentJSON `json:"-"`
}

// planExternalPlanIDVersionNewResponseAdjustmentsPlanPhaseAmountDiscountAdjustmentJSON
// contains the JSON metadata for the struct
// [PlanExternalPlanIDVersionNewResponseAdjustmentsPlanPhaseAmountDiscountAdjustment]
type planExternalPlanIDVersionNewResponseAdjustmentsPlanPhaseAmountDiscountAdjustmentJSON struct {
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

func (r *PlanExternalPlanIDVersionNewResponseAdjustmentsPlanPhaseAmountDiscountAdjustment) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r planExternalPlanIDVersionNewResponseAdjustmentsPlanPhaseAmountDiscountAdjustmentJSON) RawJSON() string {
	return r.raw
}

func (r PlanExternalPlanIDVersionNewResponseAdjustmentsPlanPhaseAmountDiscountAdjustment) implementsPlanExternalPlanIDVersionNewResponseAdjustment() {
}

type PlanExternalPlanIDVersionNewResponseAdjustmentsPlanPhaseAmountDiscountAdjustmentAdjustmentType string

const (
	PlanExternalPlanIDVersionNewResponseAdjustmentsPlanPhaseAmountDiscountAdjustmentAdjustmentTypeAmountDiscount PlanExternalPlanIDVersionNewResponseAdjustmentsPlanPhaseAmountDiscountAdjustmentAdjustmentType = "amount_discount"
)

func (r PlanExternalPlanIDVersionNewResponseAdjustmentsPlanPhaseAmountDiscountAdjustmentAdjustmentType) IsKnown() bool {
	switch r {
	case PlanExternalPlanIDVersionNewResponseAdjustmentsPlanPhaseAmountDiscountAdjustmentAdjustmentTypeAmountDiscount:
		return true
	}
	return false
}

type PlanExternalPlanIDVersionNewResponseAdjustmentsPlanPhaseAmountDiscountAdjustmentFilter struct {
	// The property of the price to filter on.
	Field PlanExternalPlanIDVersionNewResponseAdjustmentsPlanPhaseAmountDiscountAdjustmentFiltersField `json:"field,required"`
	// Should prices that match the filter be included or excluded.
	Operator PlanExternalPlanIDVersionNewResponseAdjustmentsPlanPhaseAmountDiscountAdjustmentFiltersOperator `json:"operator,required"`
	// The IDs or values that match this filter.
	Values []string                                                                                   `json:"values,required"`
	JSON   planExternalPlanIDVersionNewResponseAdjustmentsPlanPhaseAmountDiscountAdjustmentFilterJSON `json:"-"`
}

// planExternalPlanIDVersionNewResponseAdjustmentsPlanPhaseAmountDiscountAdjustmentFilterJSON
// contains the JSON metadata for the struct
// [PlanExternalPlanIDVersionNewResponseAdjustmentsPlanPhaseAmountDiscountAdjustmentFilter]
type planExternalPlanIDVersionNewResponseAdjustmentsPlanPhaseAmountDiscountAdjustmentFilterJSON struct {
	Field       apijson.Field
	Operator    apijson.Field
	Values      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *PlanExternalPlanIDVersionNewResponseAdjustmentsPlanPhaseAmountDiscountAdjustmentFilter) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r planExternalPlanIDVersionNewResponseAdjustmentsPlanPhaseAmountDiscountAdjustmentFilterJSON) RawJSON() string {
	return r.raw
}

// The property of the price to filter on.
type PlanExternalPlanIDVersionNewResponseAdjustmentsPlanPhaseAmountDiscountAdjustmentFiltersField string

const (
	PlanExternalPlanIDVersionNewResponseAdjustmentsPlanPhaseAmountDiscountAdjustmentFiltersFieldPriceID       PlanExternalPlanIDVersionNewResponseAdjustmentsPlanPhaseAmountDiscountAdjustmentFiltersField = "price_id"
	PlanExternalPlanIDVersionNewResponseAdjustmentsPlanPhaseAmountDiscountAdjustmentFiltersFieldItemID        PlanExternalPlanIDVersionNewResponseAdjustmentsPlanPhaseAmountDiscountAdjustmentFiltersField = "item_id"
	PlanExternalPlanIDVersionNewResponseAdjustmentsPlanPhaseAmountDiscountAdjustmentFiltersFieldPriceType     PlanExternalPlanIDVersionNewResponseAdjustmentsPlanPhaseAmountDiscountAdjustmentFiltersField = "price_type"
	PlanExternalPlanIDVersionNewResponseAdjustmentsPlanPhaseAmountDiscountAdjustmentFiltersFieldCurrency      PlanExternalPlanIDVersionNewResponseAdjustmentsPlanPhaseAmountDiscountAdjustmentFiltersField = "currency"
	PlanExternalPlanIDVersionNewResponseAdjustmentsPlanPhaseAmountDiscountAdjustmentFiltersFieldPricingUnitID PlanExternalPlanIDVersionNewResponseAdjustmentsPlanPhaseAmountDiscountAdjustmentFiltersField = "pricing_unit_id"
)

func (r PlanExternalPlanIDVersionNewResponseAdjustmentsPlanPhaseAmountDiscountAdjustmentFiltersField) IsKnown() bool {
	switch r {
	case PlanExternalPlanIDVersionNewResponseAdjustmentsPlanPhaseAmountDiscountAdjustmentFiltersFieldPriceID, PlanExternalPlanIDVersionNewResponseAdjustmentsPlanPhaseAmountDiscountAdjustmentFiltersFieldItemID, PlanExternalPlanIDVersionNewResponseAdjustmentsPlanPhaseAmountDiscountAdjustmentFiltersFieldPriceType, PlanExternalPlanIDVersionNewResponseAdjustmentsPlanPhaseAmountDiscountAdjustmentFiltersFieldCurrency, PlanExternalPlanIDVersionNewResponseAdjustmentsPlanPhaseAmountDiscountAdjustmentFiltersFieldPricingUnitID:
		return true
	}
	return false
}

// Should prices that match the filter be included or excluded.
type PlanExternalPlanIDVersionNewResponseAdjustmentsPlanPhaseAmountDiscountAdjustmentFiltersOperator string

const (
	PlanExternalPlanIDVersionNewResponseAdjustmentsPlanPhaseAmountDiscountAdjustmentFiltersOperatorIncludes PlanExternalPlanIDVersionNewResponseAdjustmentsPlanPhaseAmountDiscountAdjustmentFiltersOperator = "includes"
	PlanExternalPlanIDVersionNewResponseAdjustmentsPlanPhaseAmountDiscountAdjustmentFiltersOperatorExcludes PlanExternalPlanIDVersionNewResponseAdjustmentsPlanPhaseAmountDiscountAdjustmentFiltersOperator = "excludes"
)

func (r PlanExternalPlanIDVersionNewResponseAdjustmentsPlanPhaseAmountDiscountAdjustmentFiltersOperator) IsKnown() bool {
	switch r {
	case PlanExternalPlanIDVersionNewResponseAdjustmentsPlanPhaseAmountDiscountAdjustmentFiltersOperatorIncludes, PlanExternalPlanIDVersionNewResponseAdjustmentsPlanPhaseAmountDiscountAdjustmentFiltersOperatorExcludes:
		return true
	}
	return false
}

type PlanExternalPlanIDVersionNewResponseAdjustmentsPlanPhasePercentageDiscountAdjustment struct {
	ID             string                                                                                             `json:"id,required"`
	AdjustmentType PlanExternalPlanIDVersionNewResponseAdjustmentsPlanPhasePercentageDiscountAdjustmentAdjustmentType `json:"adjustment_type,required"`
	// The price IDs that this adjustment applies to.
	//
	// Deprecated: deprecated
	AppliesToPriceIDs []string `json:"applies_to_price_ids,required"`
	// The filters that determine which prices to apply this adjustment to.
	Filters []PlanExternalPlanIDVersionNewResponseAdjustmentsPlanPhasePercentageDiscountAdjustmentFilter `json:"filters,required"`
	// True for adjustments that apply to an entire invocice, false for adjustments
	// that apply to only one price.
	IsInvoiceLevel bool `json:"is_invoice_level,required"`
	// The percentage (as a value between 0 and 1) by which to discount the price
	// intervals this adjustment applies to in a given billing period.
	PercentageDiscount float64 `json:"percentage_discount,required"`
	// The plan phase in which this adjustment is active.
	PlanPhaseOrder int64 `json:"plan_phase_order,required,nullable"`
	// The reason for the adjustment.
	Reason string                                                                                   `json:"reason,required,nullable"`
	JSON   planExternalPlanIDVersionNewResponseAdjustmentsPlanPhasePercentageDiscountAdjustmentJSON `json:"-"`
}

// planExternalPlanIDVersionNewResponseAdjustmentsPlanPhasePercentageDiscountAdjustmentJSON
// contains the JSON metadata for the struct
// [PlanExternalPlanIDVersionNewResponseAdjustmentsPlanPhasePercentageDiscountAdjustment]
type planExternalPlanIDVersionNewResponseAdjustmentsPlanPhasePercentageDiscountAdjustmentJSON struct {
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

func (r *PlanExternalPlanIDVersionNewResponseAdjustmentsPlanPhasePercentageDiscountAdjustment) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r planExternalPlanIDVersionNewResponseAdjustmentsPlanPhasePercentageDiscountAdjustmentJSON) RawJSON() string {
	return r.raw
}

func (r PlanExternalPlanIDVersionNewResponseAdjustmentsPlanPhasePercentageDiscountAdjustment) implementsPlanExternalPlanIDVersionNewResponseAdjustment() {
}

type PlanExternalPlanIDVersionNewResponseAdjustmentsPlanPhasePercentageDiscountAdjustmentAdjustmentType string

const (
	PlanExternalPlanIDVersionNewResponseAdjustmentsPlanPhasePercentageDiscountAdjustmentAdjustmentTypePercentageDiscount PlanExternalPlanIDVersionNewResponseAdjustmentsPlanPhasePercentageDiscountAdjustmentAdjustmentType = "percentage_discount"
)

func (r PlanExternalPlanIDVersionNewResponseAdjustmentsPlanPhasePercentageDiscountAdjustmentAdjustmentType) IsKnown() bool {
	switch r {
	case PlanExternalPlanIDVersionNewResponseAdjustmentsPlanPhasePercentageDiscountAdjustmentAdjustmentTypePercentageDiscount:
		return true
	}
	return false
}

type PlanExternalPlanIDVersionNewResponseAdjustmentsPlanPhasePercentageDiscountAdjustmentFilter struct {
	// The property of the price to filter on.
	Field PlanExternalPlanIDVersionNewResponseAdjustmentsPlanPhasePercentageDiscountAdjustmentFiltersField `json:"field,required"`
	// Should prices that match the filter be included or excluded.
	Operator PlanExternalPlanIDVersionNewResponseAdjustmentsPlanPhasePercentageDiscountAdjustmentFiltersOperator `json:"operator,required"`
	// The IDs or values that match this filter.
	Values []string                                                                                       `json:"values,required"`
	JSON   planExternalPlanIDVersionNewResponseAdjustmentsPlanPhasePercentageDiscountAdjustmentFilterJSON `json:"-"`
}

// planExternalPlanIDVersionNewResponseAdjustmentsPlanPhasePercentageDiscountAdjustmentFilterJSON
// contains the JSON metadata for the struct
// [PlanExternalPlanIDVersionNewResponseAdjustmentsPlanPhasePercentageDiscountAdjustmentFilter]
type planExternalPlanIDVersionNewResponseAdjustmentsPlanPhasePercentageDiscountAdjustmentFilterJSON struct {
	Field       apijson.Field
	Operator    apijson.Field
	Values      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *PlanExternalPlanIDVersionNewResponseAdjustmentsPlanPhasePercentageDiscountAdjustmentFilter) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r planExternalPlanIDVersionNewResponseAdjustmentsPlanPhasePercentageDiscountAdjustmentFilterJSON) RawJSON() string {
	return r.raw
}

// The property of the price to filter on.
type PlanExternalPlanIDVersionNewResponseAdjustmentsPlanPhasePercentageDiscountAdjustmentFiltersField string

const (
	PlanExternalPlanIDVersionNewResponseAdjustmentsPlanPhasePercentageDiscountAdjustmentFiltersFieldPriceID       PlanExternalPlanIDVersionNewResponseAdjustmentsPlanPhasePercentageDiscountAdjustmentFiltersField = "price_id"
	PlanExternalPlanIDVersionNewResponseAdjustmentsPlanPhasePercentageDiscountAdjustmentFiltersFieldItemID        PlanExternalPlanIDVersionNewResponseAdjustmentsPlanPhasePercentageDiscountAdjustmentFiltersField = "item_id"
	PlanExternalPlanIDVersionNewResponseAdjustmentsPlanPhasePercentageDiscountAdjustmentFiltersFieldPriceType     PlanExternalPlanIDVersionNewResponseAdjustmentsPlanPhasePercentageDiscountAdjustmentFiltersField = "price_type"
	PlanExternalPlanIDVersionNewResponseAdjustmentsPlanPhasePercentageDiscountAdjustmentFiltersFieldCurrency      PlanExternalPlanIDVersionNewResponseAdjustmentsPlanPhasePercentageDiscountAdjustmentFiltersField = "currency"
	PlanExternalPlanIDVersionNewResponseAdjustmentsPlanPhasePercentageDiscountAdjustmentFiltersFieldPricingUnitID PlanExternalPlanIDVersionNewResponseAdjustmentsPlanPhasePercentageDiscountAdjustmentFiltersField = "pricing_unit_id"
)

func (r PlanExternalPlanIDVersionNewResponseAdjustmentsPlanPhasePercentageDiscountAdjustmentFiltersField) IsKnown() bool {
	switch r {
	case PlanExternalPlanIDVersionNewResponseAdjustmentsPlanPhasePercentageDiscountAdjustmentFiltersFieldPriceID, PlanExternalPlanIDVersionNewResponseAdjustmentsPlanPhasePercentageDiscountAdjustmentFiltersFieldItemID, PlanExternalPlanIDVersionNewResponseAdjustmentsPlanPhasePercentageDiscountAdjustmentFiltersFieldPriceType, PlanExternalPlanIDVersionNewResponseAdjustmentsPlanPhasePercentageDiscountAdjustmentFiltersFieldCurrency, PlanExternalPlanIDVersionNewResponseAdjustmentsPlanPhasePercentageDiscountAdjustmentFiltersFieldPricingUnitID:
		return true
	}
	return false
}

// Should prices that match the filter be included or excluded.
type PlanExternalPlanIDVersionNewResponseAdjustmentsPlanPhasePercentageDiscountAdjustmentFiltersOperator string

const (
	PlanExternalPlanIDVersionNewResponseAdjustmentsPlanPhasePercentageDiscountAdjustmentFiltersOperatorIncludes PlanExternalPlanIDVersionNewResponseAdjustmentsPlanPhasePercentageDiscountAdjustmentFiltersOperator = "includes"
	PlanExternalPlanIDVersionNewResponseAdjustmentsPlanPhasePercentageDiscountAdjustmentFiltersOperatorExcludes PlanExternalPlanIDVersionNewResponseAdjustmentsPlanPhasePercentageDiscountAdjustmentFiltersOperator = "excludes"
)

func (r PlanExternalPlanIDVersionNewResponseAdjustmentsPlanPhasePercentageDiscountAdjustmentFiltersOperator) IsKnown() bool {
	switch r {
	case PlanExternalPlanIDVersionNewResponseAdjustmentsPlanPhasePercentageDiscountAdjustmentFiltersOperatorIncludes, PlanExternalPlanIDVersionNewResponseAdjustmentsPlanPhasePercentageDiscountAdjustmentFiltersOperatorExcludes:
		return true
	}
	return false
}

type PlanExternalPlanIDVersionNewResponseAdjustmentsPlanPhaseMinimumAdjustment struct {
	ID             string                                                                                  `json:"id,required"`
	AdjustmentType PlanExternalPlanIDVersionNewResponseAdjustmentsPlanPhaseMinimumAdjustmentAdjustmentType `json:"adjustment_type,required"`
	// The price IDs that this adjustment applies to.
	//
	// Deprecated: deprecated
	AppliesToPriceIDs []string `json:"applies_to_price_ids,required"`
	// The filters that determine which prices to apply this adjustment to.
	Filters []PlanExternalPlanIDVersionNewResponseAdjustmentsPlanPhaseMinimumAdjustmentFilter `json:"filters,required"`
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
	Reason string                                                                        `json:"reason,required,nullable"`
	JSON   planExternalPlanIDVersionNewResponseAdjustmentsPlanPhaseMinimumAdjustmentJSON `json:"-"`
}

// planExternalPlanIDVersionNewResponseAdjustmentsPlanPhaseMinimumAdjustmentJSON
// contains the JSON metadata for the struct
// [PlanExternalPlanIDVersionNewResponseAdjustmentsPlanPhaseMinimumAdjustment]
type planExternalPlanIDVersionNewResponseAdjustmentsPlanPhaseMinimumAdjustmentJSON struct {
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

func (r *PlanExternalPlanIDVersionNewResponseAdjustmentsPlanPhaseMinimumAdjustment) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r planExternalPlanIDVersionNewResponseAdjustmentsPlanPhaseMinimumAdjustmentJSON) RawJSON() string {
	return r.raw
}

func (r PlanExternalPlanIDVersionNewResponseAdjustmentsPlanPhaseMinimumAdjustment) implementsPlanExternalPlanIDVersionNewResponseAdjustment() {
}

type PlanExternalPlanIDVersionNewResponseAdjustmentsPlanPhaseMinimumAdjustmentAdjustmentType string

const (
	PlanExternalPlanIDVersionNewResponseAdjustmentsPlanPhaseMinimumAdjustmentAdjustmentTypeMinimum PlanExternalPlanIDVersionNewResponseAdjustmentsPlanPhaseMinimumAdjustmentAdjustmentType = "minimum"
)

func (r PlanExternalPlanIDVersionNewResponseAdjustmentsPlanPhaseMinimumAdjustmentAdjustmentType) IsKnown() bool {
	switch r {
	case PlanExternalPlanIDVersionNewResponseAdjustmentsPlanPhaseMinimumAdjustmentAdjustmentTypeMinimum:
		return true
	}
	return false
}

type PlanExternalPlanIDVersionNewResponseAdjustmentsPlanPhaseMinimumAdjustmentFilter struct {
	// The property of the price to filter on.
	Field PlanExternalPlanIDVersionNewResponseAdjustmentsPlanPhaseMinimumAdjustmentFiltersField `json:"field,required"`
	// Should prices that match the filter be included or excluded.
	Operator PlanExternalPlanIDVersionNewResponseAdjustmentsPlanPhaseMinimumAdjustmentFiltersOperator `json:"operator,required"`
	// The IDs or values that match this filter.
	Values []string                                                                            `json:"values,required"`
	JSON   planExternalPlanIDVersionNewResponseAdjustmentsPlanPhaseMinimumAdjustmentFilterJSON `json:"-"`
}

// planExternalPlanIDVersionNewResponseAdjustmentsPlanPhaseMinimumAdjustmentFilterJSON
// contains the JSON metadata for the struct
// [PlanExternalPlanIDVersionNewResponseAdjustmentsPlanPhaseMinimumAdjustmentFilter]
type planExternalPlanIDVersionNewResponseAdjustmentsPlanPhaseMinimumAdjustmentFilterJSON struct {
	Field       apijson.Field
	Operator    apijson.Field
	Values      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *PlanExternalPlanIDVersionNewResponseAdjustmentsPlanPhaseMinimumAdjustmentFilter) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r planExternalPlanIDVersionNewResponseAdjustmentsPlanPhaseMinimumAdjustmentFilterJSON) RawJSON() string {
	return r.raw
}

// The property of the price to filter on.
type PlanExternalPlanIDVersionNewResponseAdjustmentsPlanPhaseMinimumAdjustmentFiltersField string

const (
	PlanExternalPlanIDVersionNewResponseAdjustmentsPlanPhaseMinimumAdjustmentFiltersFieldPriceID       PlanExternalPlanIDVersionNewResponseAdjustmentsPlanPhaseMinimumAdjustmentFiltersField = "price_id"
	PlanExternalPlanIDVersionNewResponseAdjustmentsPlanPhaseMinimumAdjustmentFiltersFieldItemID        PlanExternalPlanIDVersionNewResponseAdjustmentsPlanPhaseMinimumAdjustmentFiltersField = "item_id"
	PlanExternalPlanIDVersionNewResponseAdjustmentsPlanPhaseMinimumAdjustmentFiltersFieldPriceType     PlanExternalPlanIDVersionNewResponseAdjustmentsPlanPhaseMinimumAdjustmentFiltersField = "price_type"
	PlanExternalPlanIDVersionNewResponseAdjustmentsPlanPhaseMinimumAdjustmentFiltersFieldCurrency      PlanExternalPlanIDVersionNewResponseAdjustmentsPlanPhaseMinimumAdjustmentFiltersField = "currency"
	PlanExternalPlanIDVersionNewResponseAdjustmentsPlanPhaseMinimumAdjustmentFiltersFieldPricingUnitID PlanExternalPlanIDVersionNewResponseAdjustmentsPlanPhaseMinimumAdjustmentFiltersField = "pricing_unit_id"
)

func (r PlanExternalPlanIDVersionNewResponseAdjustmentsPlanPhaseMinimumAdjustmentFiltersField) IsKnown() bool {
	switch r {
	case PlanExternalPlanIDVersionNewResponseAdjustmentsPlanPhaseMinimumAdjustmentFiltersFieldPriceID, PlanExternalPlanIDVersionNewResponseAdjustmentsPlanPhaseMinimumAdjustmentFiltersFieldItemID, PlanExternalPlanIDVersionNewResponseAdjustmentsPlanPhaseMinimumAdjustmentFiltersFieldPriceType, PlanExternalPlanIDVersionNewResponseAdjustmentsPlanPhaseMinimumAdjustmentFiltersFieldCurrency, PlanExternalPlanIDVersionNewResponseAdjustmentsPlanPhaseMinimumAdjustmentFiltersFieldPricingUnitID:
		return true
	}
	return false
}

// Should prices that match the filter be included or excluded.
type PlanExternalPlanIDVersionNewResponseAdjustmentsPlanPhaseMinimumAdjustmentFiltersOperator string

const (
	PlanExternalPlanIDVersionNewResponseAdjustmentsPlanPhaseMinimumAdjustmentFiltersOperatorIncludes PlanExternalPlanIDVersionNewResponseAdjustmentsPlanPhaseMinimumAdjustmentFiltersOperator = "includes"
	PlanExternalPlanIDVersionNewResponseAdjustmentsPlanPhaseMinimumAdjustmentFiltersOperatorExcludes PlanExternalPlanIDVersionNewResponseAdjustmentsPlanPhaseMinimumAdjustmentFiltersOperator = "excludes"
)

func (r PlanExternalPlanIDVersionNewResponseAdjustmentsPlanPhaseMinimumAdjustmentFiltersOperator) IsKnown() bool {
	switch r {
	case PlanExternalPlanIDVersionNewResponseAdjustmentsPlanPhaseMinimumAdjustmentFiltersOperatorIncludes, PlanExternalPlanIDVersionNewResponseAdjustmentsPlanPhaseMinimumAdjustmentFiltersOperatorExcludes:
		return true
	}
	return false
}

type PlanExternalPlanIDVersionNewResponseAdjustmentsPlanPhaseMaximumAdjustment struct {
	ID             string                                                                                  `json:"id,required"`
	AdjustmentType PlanExternalPlanIDVersionNewResponseAdjustmentsPlanPhaseMaximumAdjustmentAdjustmentType `json:"adjustment_type,required"`
	// The price IDs that this adjustment applies to.
	//
	// Deprecated: deprecated
	AppliesToPriceIDs []string `json:"applies_to_price_ids,required"`
	// The filters that determine which prices to apply this adjustment to.
	Filters []PlanExternalPlanIDVersionNewResponseAdjustmentsPlanPhaseMaximumAdjustmentFilter `json:"filters,required"`
	// True for adjustments that apply to an entire invocice, false for adjustments
	// that apply to only one price.
	IsInvoiceLevel bool `json:"is_invoice_level,required"`
	// The maximum amount to charge in a given billing period for the prices this
	// adjustment applies to.
	MaximumAmount string `json:"maximum_amount,required"`
	// The plan phase in which this adjustment is active.
	PlanPhaseOrder int64 `json:"plan_phase_order,required,nullable"`
	// The reason for the adjustment.
	Reason string                                                                        `json:"reason,required,nullable"`
	JSON   planExternalPlanIDVersionNewResponseAdjustmentsPlanPhaseMaximumAdjustmentJSON `json:"-"`
}

// planExternalPlanIDVersionNewResponseAdjustmentsPlanPhaseMaximumAdjustmentJSON
// contains the JSON metadata for the struct
// [PlanExternalPlanIDVersionNewResponseAdjustmentsPlanPhaseMaximumAdjustment]
type planExternalPlanIDVersionNewResponseAdjustmentsPlanPhaseMaximumAdjustmentJSON struct {
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

func (r *PlanExternalPlanIDVersionNewResponseAdjustmentsPlanPhaseMaximumAdjustment) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r planExternalPlanIDVersionNewResponseAdjustmentsPlanPhaseMaximumAdjustmentJSON) RawJSON() string {
	return r.raw
}

func (r PlanExternalPlanIDVersionNewResponseAdjustmentsPlanPhaseMaximumAdjustment) implementsPlanExternalPlanIDVersionNewResponseAdjustment() {
}

type PlanExternalPlanIDVersionNewResponseAdjustmentsPlanPhaseMaximumAdjustmentAdjustmentType string

const (
	PlanExternalPlanIDVersionNewResponseAdjustmentsPlanPhaseMaximumAdjustmentAdjustmentTypeMaximum PlanExternalPlanIDVersionNewResponseAdjustmentsPlanPhaseMaximumAdjustmentAdjustmentType = "maximum"
)

func (r PlanExternalPlanIDVersionNewResponseAdjustmentsPlanPhaseMaximumAdjustmentAdjustmentType) IsKnown() bool {
	switch r {
	case PlanExternalPlanIDVersionNewResponseAdjustmentsPlanPhaseMaximumAdjustmentAdjustmentTypeMaximum:
		return true
	}
	return false
}

type PlanExternalPlanIDVersionNewResponseAdjustmentsPlanPhaseMaximumAdjustmentFilter struct {
	// The property of the price to filter on.
	Field PlanExternalPlanIDVersionNewResponseAdjustmentsPlanPhaseMaximumAdjustmentFiltersField `json:"field,required"`
	// Should prices that match the filter be included or excluded.
	Operator PlanExternalPlanIDVersionNewResponseAdjustmentsPlanPhaseMaximumAdjustmentFiltersOperator `json:"operator,required"`
	// The IDs or values that match this filter.
	Values []string                                                                            `json:"values,required"`
	JSON   planExternalPlanIDVersionNewResponseAdjustmentsPlanPhaseMaximumAdjustmentFilterJSON `json:"-"`
}

// planExternalPlanIDVersionNewResponseAdjustmentsPlanPhaseMaximumAdjustmentFilterJSON
// contains the JSON metadata for the struct
// [PlanExternalPlanIDVersionNewResponseAdjustmentsPlanPhaseMaximumAdjustmentFilter]
type planExternalPlanIDVersionNewResponseAdjustmentsPlanPhaseMaximumAdjustmentFilterJSON struct {
	Field       apijson.Field
	Operator    apijson.Field
	Values      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *PlanExternalPlanIDVersionNewResponseAdjustmentsPlanPhaseMaximumAdjustmentFilter) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r planExternalPlanIDVersionNewResponseAdjustmentsPlanPhaseMaximumAdjustmentFilterJSON) RawJSON() string {
	return r.raw
}

// The property of the price to filter on.
type PlanExternalPlanIDVersionNewResponseAdjustmentsPlanPhaseMaximumAdjustmentFiltersField string

const (
	PlanExternalPlanIDVersionNewResponseAdjustmentsPlanPhaseMaximumAdjustmentFiltersFieldPriceID       PlanExternalPlanIDVersionNewResponseAdjustmentsPlanPhaseMaximumAdjustmentFiltersField = "price_id"
	PlanExternalPlanIDVersionNewResponseAdjustmentsPlanPhaseMaximumAdjustmentFiltersFieldItemID        PlanExternalPlanIDVersionNewResponseAdjustmentsPlanPhaseMaximumAdjustmentFiltersField = "item_id"
	PlanExternalPlanIDVersionNewResponseAdjustmentsPlanPhaseMaximumAdjustmentFiltersFieldPriceType     PlanExternalPlanIDVersionNewResponseAdjustmentsPlanPhaseMaximumAdjustmentFiltersField = "price_type"
	PlanExternalPlanIDVersionNewResponseAdjustmentsPlanPhaseMaximumAdjustmentFiltersFieldCurrency      PlanExternalPlanIDVersionNewResponseAdjustmentsPlanPhaseMaximumAdjustmentFiltersField = "currency"
	PlanExternalPlanIDVersionNewResponseAdjustmentsPlanPhaseMaximumAdjustmentFiltersFieldPricingUnitID PlanExternalPlanIDVersionNewResponseAdjustmentsPlanPhaseMaximumAdjustmentFiltersField = "pricing_unit_id"
)

func (r PlanExternalPlanIDVersionNewResponseAdjustmentsPlanPhaseMaximumAdjustmentFiltersField) IsKnown() bool {
	switch r {
	case PlanExternalPlanIDVersionNewResponseAdjustmentsPlanPhaseMaximumAdjustmentFiltersFieldPriceID, PlanExternalPlanIDVersionNewResponseAdjustmentsPlanPhaseMaximumAdjustmentFiltersFieldItemID, PlanExternalPlanIDVersionNewResponseAdjustmentsPlanPhaseMaximumAdjustmentFiltersFieldPriceType, PlanExternalPlanIDVersionNewResponseAdjustmentsPlanPhaseMaximumAdjustmentFiltersFieldCurrency, PlanExternalPlanIDVersionNewResponseAdjustmentsPlanPhaseMaximumAdjustmentFiltersFieldPricingUnitID:
		return true
	}
	return false
}

// Should prices that match the filter be included or excluded.
type PlanExternalPlanIDVersionNewResponseAdjustmentsPlanPhaseMaximumAdjustmentFiltersOperator string

const (
	PlanExternalPlanIDVersionNewResponseAdjustmentsPlanPhaseMaximumAdjustmentFiltersOperatorIncludes PlanExternalPlanIDVersionNewResponseAdjustmentsPlanPhaseMaximumAdjustmentFiltersOperator = "includes"
	PlanExternalPlanIDVersionNewResponseAdjustmentsPlanPhaseMaximumAdjustmentFiltersOperatorExcludes PlanExternalPlanIDVersionNewResponseAdjustmentsPlanPhaseMaximumAdjustmentFiltersOperator = "excludes"
)

func (r PlanExternalPlanIDVersionNewResponseAdjustmentsPlanPhaseMaximumAdjustmentFiltersOperator) IsKnown() bool {
	switch r {
	case PlanExternalPlanIDVersionNewResponseAdjustmentsPlanPhaseMaximumAdjustmentFiltersOperatorIncludes, PlanExternalPlanIDVersionNewResponseAdjustmentsPlanPhaseMaximumAdjustmentFiltersOperatorExcludes:
		return true
	}
	return false
}

type PlanExternalPlanIDVersionNewResponseAdjustmentsAdjustmentType string

const (
	PlanExternalPlanIDVersionNewResponseAdjustmentsAdjustmentTypeUsageDiscount      PlanExternalPlanIDVersionNewResponseAdjustmentsAdjustmentType = "usage_discount"
	PlanExternalPlanIDVersionNewResponseAdjustmentsAdjustmentTypeAmountDiscount     PlanExternalPlanIDVersionNewResponseAdjustmentsAdjustmentType = "amount_discount"
	PlanExternalPlanIDVersionNewResponseAdjustmentsAdjustmentTypePercentageDiscount PlanExternalPlanIDVersionNewResponseAdjustmentsAdjustmentType = "percentage_discount"
	PlanExternalPlanIDVersionNewResponseAdjustmentsAdjustmentTypeMinimum            PlanExternalPlanIDVersionNewResponseAdjustmentsAdjustmentType = "minimum"
	PlanExternalPlanIDVersionNewResponseAdjustmentsAdjustmentTypeMaximum            PlanExternalPlanIDVersionNewResponseAdjustmentsAdjustmentType = "maximum"
)

func (r PlanExternalPlanIDVersionNewResponseAdjustmentsAdjustmentType) IsKnown() bool {
	switch r {
	case PlanExternalPlanIDVersionNewResponseAdjustmentsAdjustmentTypeUsageDiscount, PlanExternalPlanIDVersionNewResponseAdjustmentsAdjustmentTypeAmountDiscount, PlanExternalPlanIDVersionNewResponseAdjustmentsAdjustmentTypePercentageDiscount, PlanExternalPlanIDVersionNewResponseAdjustmentsAdjustmentTypeMinimum, PlanExternalPlanIDVersionNewResponseAdjustmentsAdjustmentTypeMaximum:
		return true
	}
	return false
}

type PlanExternalPlanIDVersionNewResponsePlanPhase struct {
	ID          string `json:"id,required"`
	Description string `json:"description,required,nullable"`
	// How many terms of length `duration_unit` this phase is active for. If null, this
	// phase is evergreen and active indefinitely
	Duration     int64                                                      `json:"duration,required,nullable"`
	DurationUnit PlanExternalPlanIDVersionNewResponsePlanPhasesDurationUnit `json:"duration_unit,required,nullable"`
	Name         string                                                     `json:"name,required"`
	// Determines the ordering of the phase in a plan's lifecycle. 1 = first phase.
	Order int64                                             `json:"order,required"`
	JSON  planExternalPlanIDVersionNewResponsePlanPhaseJSON `json:"-"`
}

// planExternalPlanIDVersionNewResponsePlanPhaseJSON contains the JSON metadata for
// the struct [PlanExternalPlanIDVersionNewResponsePlanPhase]
type planExternalPlanIDVersionNewResponsePlanPhaseJSON struct {
	ID           apijson.Field
	Description  apijson.Field
	Duration     apijson.Field
	DurationUnit apijson.Field
	Name         apijson.Field
	Order        apijson.Field
	raw          string
	ExtraFields  map[string]apijson.Field
}

func (r *PlanExternalPlanIDVersionNewResponsePlanPhase) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r planExternalPlanIDVersionNewResponsePlanPhaseJSON) RawJSON() string {
	return r.raw
}

type PlanExternalPlanIDVersionNewResponsePlanPhasesDurationUnit string

const (
	PlanExternalPlanIDVersionNewResponsePlanPhasesDurationUnitDaily      PlanExternalPlanIDVersionNewResponsePlanPhasesDurationUnit = "daily"
	PlanExternalPlanIDVersionNewResponsePlanPhasesDurationUnitMonthly    PlanExternalPlanIDVersionNewResponsePlanPhasesDurationUnit = "monthly"
	PlanExternalPlanIDVersionNewResponsePlanPhasesDurationUnitQuarterly  PlanExternalPlanIDVersionNewResponsePlanPhasesDurationUnit = "quarterly"
	PlanExternalPlanIDVersionNewResponsePlanPhasesDurationUnitSemiAnnual PlanExternalPlanIDVersionNewResponsePlanPhasesDurationUnit = "semi_annual"
	PlanExternalPlanIDVersionNewResponsePlanPhasesDurationUnitAnnual     PlanExternalPlanIDVersionNewResponsePlanPhasesDurationUnit = "annual"
)

func (r PlanExternalPlanIDVersionNewResponsePlanPhasesDurationUnit) IsKnown() bool {
	switch r {
	case PlanExternalPlanIDVersionNewResponsePlanPhasesDurationUnitDaily, PlanExternalPlanIDVersionNewResponsePlanPhasesDurationUnitMonthly, PlanExternalPlanIDVersionNewResponsePlanPhasesDurationUnitQuarterly, PlanExternalPlanIDVersionNewResponsePlanPhasesDurationUnitSemiAnnual, PlanExternalPlanIDVersionNewResponsePlanPhasesDurationUnitAnnual:
		return true
	}
	return false
}

// The PlanVersion resource represents the prices and adjustments present on a
// specific version of a plan.
type PlanExternalPlanIDVersionGetResponse struct {
	// Adjustments for this plan. If the plan has phases, this includes adjustments
	// across all phases of the plan.
	Adjustments []PlanExternalPlanIDVersionGetResponseAdjustment `json:"adjustments,required"`
	CreatedAt   time.Time                                        `json:"created_at,required" format:"date-time"`
	PlanPhases  []PlanExternalPlanIDVersionGetResponsePlanPhase  `json:"plan_phases,required,nullable"`
	// Prices for this plan. If the plan has phases, this includes prices across all
	// phases of the plan.
	Prices  []Price                                  `json:"prices,required"`
	Version int64                                    `json:"version,required"`
	JSON    planExternalPlanIDVersionGetResponseJSON `json:"-"`
}

// planExternalPlanIDVersionGetResponseJSON contains the JSON metadata for the
// struct [PlanExternalPlanIDVersionGetResponse]
type planExternalPlanIDVersionGetResponseJSON struct {
	Adjustments apijson.Field
	CreatedAt   apijson.Field
	PlanPhases  apijson.Field
	Prices      apijson.Field
	Version     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *PlanExternalPlanIDVersionGetResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r planExternalPlanIDVersionGetResponseJSON) RawJSON() string {
	return r.raw
}

type PlanExternalPlanIDVersionGetResponseAdjustment struct {
	ID             string                                                        `json:"id,required"`
	AdjustmentType PlanExternalPlanIDVersionGetResponseAdjustmentsAdjustmentType `json:"adjustment_type,required"`
	// This field can have the runtime type of [[]string].
	AppliesToPriceIDs interface{} `json:"applies_to_price_ids,required"`
	// This field can have the runtime type of
	// [[]PlanExternalPlanIDVersionGetResponseAdjustmentsPlanPhaseUsageDiscountAdjustmentFilter],
	// [[]PlanExternalPlanIDVersionGetResponseAdjustmentsPlanPhaseAmountDiscountAdjustmentFilter],
	// [[]PlanExternalPlanIDVersionGetResponseAdjustmentsPlanPhasePercentageDiscountAdjustmentFilter],
	// [[]PlanExternalPlanIDVersionGetResponseAdjustmentsPlanPhaseMinimumAdjustmentFilter],
	// [[]PlanExternalPlanIDVersionGetResponseAdjustmentsPlanPhaseMaximumAdjustmentFilter].
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
	UsageDiscount float64                                            `json:"usage_discount"`
	JSON          planExternalPlanIDVersionGetResponseAdjustmentJSON `json:"-"`
	union         PlanExternalPlanIDVersionGetResponseAdjustmentsUnion
}

// planExternalPlanIDVersionGetResponseAdjustmentJSON contains the JSON metadata
// for the struct [PlanExternalPlanIDVersionGetResponseAdjustment]
type planExternalPlanIDVersionGetResponseAdjustmentJSON struct {
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

func (r planExternalPlanIDVersionGetResponseAdjustmentJSON) RawJSON() string {
	return r.raw
}

func (r *PlanExternalPlanIDVersionGetResponseAdjustment) UnmarshalJSON(data []byte) (err error) {
	*r = PlanExternalPlanIDVersionGetResponseAdjustment{}
	err = apijson.UnmarshalRoot(data, &r.union)
	if err != nil {
		return err
	}
	return apijson.Port(r.union, &r)
}

// AsUnion returns a [PlanExternalPlanIDVersionGetResponseAdjustmentsUnion]
// interface which you can cast to the specific types for more type safety.
//
// Possible runtime types of the union are
// [PlanExternalPlanIDVersionGetResponseAdjustmentsPlanPhaseUsageDiscountAdjustment],
// [PlanExternalPlanIDVersionGetResponseAdjustmentsPlanPhaseAmountDiscountAdjustment],
// [PlanExternalPlanIDVersionGetResponseAdjustmentsPlanPhasePercentageDiscountAdjustment],
// [PlanExternalPlanIDVersionGetResponseAdjustmentsPlanPhaseMinimumAdjustment],
// [PlanExternalPlanIDVersionGetResponseAdjustmentsPlanPhaseMaximumAdjustment].
func (r PlanExternalPlanIDVersionGetResponseAdjustment) AsUnion() PlanExternalPlanIDVersionGetResponseAdjustmentsUnion {
	return r.union
}

// Union satisfied by
// [PlanExternalPlanIDVersionGetResponseAdjustmentsPlanPhaseUsageDiscountAdjustment],
// [PlanExternalPlanIDVersionGetResponseAdjustmentsPlanPhaseAmountDiscountAdjustment],
// [PlanExternalPlanIDVersionGetResponseAdjustmentsPlanPhasePercentageDiscountAdjustment],
// [PlanExternalPlanIDVersionGetResponseAdjustmentsPlanPhaseMinimumAdjustment] or
// [PlanExternalPlanIDVersionGetResponseAdjustmentsPlanPhaseMaximumAdjustment].
type PlanExternalPlanIDVersionGetResponseAdjustmentsUnion interface {
	implementsPlanExternalPlanIDVersionGetResponseAdjustment()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*PlanExternalPlanIDVersionGetResponseAdjustmentsUnion)(nil)).Elem(),
		"adjustment_type",
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(PlanExternalPlanIDVersionGetResponseAdjustmentsPlanPhaseUsageDiscountAdjustment{}),
			DiscriminatorValue: "usage_discount",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(PlanExternalPlanIDVersionGetResponseAdjustmentsPlanPhaseAmountDiscountAdjustment{}),
			DiscriminatorValue: "amount_discount",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(PlanExternalPlanIDVersionGetResponseAdjustmentsPlanPhasePercentageDiscountAdjustment{}),
			DiscriminatorValue: "percentage_discount",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(PlanExternalPlanIDVersionGetResponseAdjustmentsPlanPhaseMinimumAdjustment{}),
			DiscriminatorValue: "minimum",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(PlanExternalPlanIDVersionGetResponseAdjustmentsPlanPhaseMaximumAdjustment{}),
			DiscriminatorValue: "maximum",
		},
	)
}

type PlanExternalPlanIDVersionGetResponseAdjustmentsPlanPhaseUsageDiscountAdjustment struct {
	ID             string                                                                                        `json:"id,required"`
	AdjustmentType PlanExternalPlanIDVersionGetResponseAdjustmentsPlanPhaseUsageDiscountAdjustmentAdjustmentType `json:"adjustment_type,required"`
	// The price IDs that this adjustment applies to.
	//
	// Deprecated: deprecated
	AppliesToPriceIDs []string `json:"applies_to_price_ids,required"`
	// The filters that determine which prices to apply this adjustment to.
	Filters []PlanExternalPlanIDVersionGetResponseAdjustmentsPlanPhaseUsageDiscountAdjustmentFilter `json:"filters,required"`
	// True for adjustments that apply to an entire invocice, false for adjustments
	// that apply to only one price.
	IsInvoiceLevel bool `json:"is_invoice_level,required"`
	// The plan phase in which this adjustment is active.
	PlanPhaseOrder int64 `json:"plan_phase_order,required,nullable"`
	// The reason for the adjustment.
	Reason string `json:"reason,required,nullable"`
	// The number of usage units by which to discount the price this adjustment applies
	// to in a given billing period.
	UsageDiscount float64                                                                             `json:"usage_discount,required"`
	JSON          planExternalPlanIDVersionGetResponseAdjustmentsPlanPhaseUsageDiscountAdjustmentJSON `json:"-"`
}

// planExternalPlanIDVersionGetResponseAdjustmentsPlanPhaseUsageDiscountAdjustmentJSON
// contains the JSON metadata for the struct
// [PlanExternalPlanIDVersionGetResponseAdjustmentsPlanPhaseUsageDiscountAdjustment]
type planExternalPlanIDVersionGetResponseAdjustmentsPlanPhaseUsageDiscountAdjustmentJSON struct {
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

func (r *PlanExternalPlanIDVersionGetResponseAdjustmentsPlanPhaseUsageDiscountAdjustment) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r planExternalPlanIDVersionGetResponseAdjustmentsPlanPhaseUsageDiscountAdjustmentJSON) RawJSON() string {
	return r.raw
}

func (r PlanExternalPlanIDVersionGetResponseAdjustmentsPlanPhaseUsageDiscountAdjustment) implementsPlanExternalPlanIDVersionGetResponseAdjustment() {
}

type PlanExternalPlanIDVersionGetResponseAdjustmentsPlanPhaseUsageDiscountAdjustmentAdjustmentType string

const (
	PlanExternalPlanIDVersionGetResponseAdjustmentsPlanPhaseUsageDiscountAdjustmentAdjustmentTypeUsageDiscount PlanExternalPlanIDVersionGetResponseAdjustmentsPlanPhaseUsageDiscountAdjustmentAdjustmentType = "usage_discount"
)

func (r PlanExternalPlanIDVersionGetResponseAdjustmentsPlanPhaseUsageDiscountAdjustmentAdjustmentType) IsKnown() bool {
	switch r {
	case PlanExternalPlanIDVersionGetResponseAdjustmentsPlanPhaseUsageDiscountAdjustmentAdjustmentTypeUsageDiscount:
		return true
	}
	return false
}

type PlanExternalPlanIDVersionGetResponseAdjustmentsPlanPhaseUsageDiscountAdjustmentFilter struct {
	// The property of the price to filter on.
	Field PlanExternalPlanIDVersionGetResponseAdjustmentsPlanPhaseUsageDiscountAdjustmentFiltersField `json:"field,required"`
	// Should prices that match the filter be included or excluded.
	Operator PlanExternalPlanIDVersionGetResponseAdjustmentsPlanPhaseUsageDiscountAdjustmentFiltersOperator `json:"operator,required"`
	// The IDs or values that match this filter.
	Values []string                                                                                  `json:"values,required"`
	JSON   planExternalPlanIDVersionGetResponseAdjustmentsPlanPhaseUsageDiscountAdjustmentFilterJSON `json:"-"`
}

// planExternalPlanIDVersionGetResponseAdjustmentsPlanPhaseUsageDiscountAdjustmentFilterJSON
// contains the JSON metadata for the struct
// [PlanExternalPlanIDVersionGetResponseAdjustmentsPlanPhaseUsageDiscountAdjustmentFilter]
type planExternalPlanIDVersionGetResponseAdjustmentsPlanPhaseUsageDiscountAdjustmentFilterJSON struct {
	Field       apijson.Field
	Operator    apijson.Field
	Values      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *PlanExternalPlanIDVersionGetResponseAdjustmentsPlanPhaseUsageDiscountAdjustmentFilter) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r planExternalPlanIDVersionGetResponseAdjustmentsPlanPhaseUsageDiscountAdjustmentFilterJSON) RawJSON() string {
	return r.raw
}

// The property of the price to filter on.
type PlanExternalPlanIDVersionGetResponseAdjustmentsPlanPhaseUsageDiscountAdjustmentFiltersField string

const (
	PlanExternalPlanIDVersionGetResponseAdjustmentsPlanPhaseUsageDiscountAdjustmentFiltersFieldPriceID       PlanExternalPlanIDVersionGetResponseAdjustmentsPlanPhaseUsageDiscountAdjustmentFiltersField = "price_id"
	PlanExternalPlanIDVersionGetResponseAdjustmentsPlanPhaseUsageDiscountAdjustmentFiltersFieldItemID        PlanExternalPlanIDVersionGetResponseAdjustmentsPlanPhaseUsageDiscountAdjustmentFiltersField = "item_id"
	PlanExternalPlanIDVersionGetResponseAdjustmentsPlanPhaseUsageDiscountAdjustmentFiltersFieldPriceType     PlanExternalPlanIDVersionGetResponseAdjustmentsPlanPhaseUsageDiscountAdjustmentFiltersField = "price_type"
	PlanExternalPlanIDVersionGetResponseAdjustmentsPlanPhaseUsageDiscountAdjustmentFiltersFieldCurrency      PlanExternalPlanIDVersionGetResponseAdjustmentsPlanPhaseUsageDiscountAdjustmentFiltersField = "currency"
	PlanExternalPlanIDVersionGetResponseAdjustmentsPlanPhaseUsageDiscountAdjustmentFiltersFieldPricingUnitID PlanExternalPlanIDVersionGetResponseAdjustmentsPlanPhaseUsageDiscountAdjustmentFiltersField = "pricing_unit_id"
)

func (r PlanExternalPlanIDVersionGetResponseAdjustmentsPlanPhaseUsageDiscountAdjustmentFiltersField) IsKnown() bool {
	switch r {
	case PlanExternalPlanIDVersionGetResponseAdjustmentsPlanPhaseUsageDiscountAdjustmentFiltersFieldPriceID, PlanExternalPlanIDVersionGetResponseAdjustmentsPlanPhaseUsageDiscountAdjustmentFiltersFieldItemID, PlanExternalPlanIDVersionGetResponseAdjustmentsPlanPhaseUsageDiscountAdjustmentFiltersFieldPriceType, PlanExternalPlanIDVersionGetResponseAdjustmentsPlanPhaseUsageDiscountAdjustmentFiltersFieldCurrency, PlanExternalPlanIDVersionGetResponseAdjustmentsPlanPhaseUsageDiscountAdjustmentFiltersFieldPricingUnitID:
		return true
	}
	return false
}

// Should prices that match the filter be included or excluded.
type PlanExternalPlanIDVersionGetResponseAdjustmentsPlanPhaseUsageDiscountAdjustmentFiltersOperator string

const (
	PlanExternalPlanIDVersionGetResponseAdjustmentsPlanPhaseUsageDiscountAdjustmentFiltersOperatorIncludes PlanExternalPlanIDVersionGetResponseAdjustmentsPlanPhaseUsageDiscountAdjustmentFiltersOperator = "includes"
	PlanExternalPlanIDVersionGetResponseAdjustmentsPlanPhaseUsageDiscountAdjustmentFiltersOperatorExcludes PlanExternalPlanIDVersionGetResponseAdjustmentsPlanPhaseUsageDiscountAdjustmentFiltersOperator = "excludes"
)

func (r PlanExternalPlanIDVersionGetResponseAdjustmentsPlanPhaseUsageDiscountAdjustmentFiltersOperator) IsKnown() bool {
	switch r {
	case PlanExternalPlanIDVersionGetResponseAdjustmentsPlanPhaseUsageDiscountAdjustmentFiltersOperatorIncludes, PlanExternalPlanIDVersionGetResponseAdjustmentsPlanPhaseUsageDiscountAdjustmentFiltersOperatorExcludes:
		return true
	}
	return false
}

type PlanExternalPlanIDVersionGetResponseAdjustmentsPlanPhaseAmountDiscountAdjustment struct {
	ID             string                                                                                         `json:"id,required"`
	AdjustmentType PlanExternalPlanIDVersionGetResponseAdjustmentsPlanPhaseAmountDiscountAdjustmentAdjustmentType `json:"adjustment_type,required"`
	// The amount by which to discount the prices this adjustment applies to in a given
	// billing period.
	AmountDiscount string `json:"amount_discount,required"`
	// The price IDs that this adjustment applies to.
	//
	// Deprecated: deprecated
	AppliesToPriceIDs []string `json:"applies_to_price_ids,required"`
	// The filters that determine which prices to apply this adjustment to.
	Filters []PlanExternalPlanIDVersionGetResponseAdjustmentsPlanPhaseAmountDiscountAdjustmentFilter `json:"filters,required"`
	// True for adjustments that apply to an entire invocice, false for adjustments
	// that apply to only one price.
	IsInvoiceLevel bool `json:"is_invoice_level,required"`
	// The plan phase in which this adjustment is active.
	PlanPhaseOrder int64 `json:"plan_phase_order,required,nullable"`
	// The reason for the adjustment.
	Reason string                                                                               `json:"reason,required,nullable"`
	JSON   planExternalPlanIDVersionGetResponseAdjustmentsPlanPhaseAmountDiscountAdjustmentJSON `json:"-"`
}

// planExternalPlanIDVersionGetResponseAdjustmentsPlanPhaseAmountDiscountAdjustmentJSON
// contains the JSON metadata for the struct
// [PlanExternalPlanIDVersionGetResponseAdjustmentsPlanPhaseAmountDiscountAdjustment]
type planExternalPlanIDVersionGetResponseAdjustmentsPlanPhaseAmountDiscountAdjustmentJSON struct {
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

func (r *PlanExternalPlanIDVersionGetResponseAdjustmentsPlanPhaseAmountDiscountAdjustment) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r planExternalPlanIDVersionGetResponseAdjustmentsPlanPhaseAmountDiscountAdjustmentJSON) RawJSON() string {
	return r.raw
}

func (r PlanExternalPlanIDVersionGetResponseAdjustmentsPlanPhaseAmountDiscountAdjustment) implementsPlanExternalPlanIDVersionGetResponseAdjustment() {
}

type PlanExternalPlanIDVersionGetResponseAdjustmentsPlanPhaseAmountDiscountAdjustmentAdjustmentType string

const (
	PlanExternalPlanIDVersionGetResponseAdjustmentsPlanPhaseAmountDiscountAdjustmentAdjustmentTypeAmountDiscount PlanExternalPlanIDVersionGetResponseAdjustmentsPlanPhaseAmountDiscountAdjustmentAdjustmentType = "amount_discount"
)

func (r PlanExternalPlanIDVersionGetResponseAdjustmentsPlanPhaseAmountDiscountAdjustmentAdjustmentType) IsKnown() bool {
	switch r {
	case PlanExternalPlanIDVersionGetResponseAdjustmentsPlanPhaseAmountDiscountAdjustmentAdjustmentTypeAmountDiscount:
		return true
	}
	return false
}

type PlanExternalPlanIDVersionGetResponseAdjustmentsPlanPhaseAmountDiscountAdjustmentFilter struct {
	// The property of the price to filter on.
	Field PlanExternalPlanIDVersionGetResponseAdjustmentsPlanPhaseAmountDiscountAdjustmentFiltersField `json:"field,required"`
	// Should prices that match the filter be included or excluded.
	Operator PlanExternalPlanIDVersionGetResponseAdjustmentsPlanPhaseAmountDiscountAdjustmentFiltersOperator `json:"operator,required"`
	// The IDs or values that match this filter.
	Values []string                                                                                   `json:"values,required"`
	JSON   planExternalPlanIDVersionGetResponseAdjustmentsPlanPhaseAmountDiscountAdjustmentFilterJSON `json:"-"`
}

// planExternalPlanIDVersionGetResponseAdjustmentsPlanPhaseAmountDiscountAdjustmentFilterJSON
// contains the JSON metadata for the struct
// [PlanExternalPlanIDVersionGetResponseAdjustmentsPlanPhaseAmountDiscountAdjustmentFilter]
type planExternalPlanIDVersionGetResponseAdjustmentsPlanPhaseAmountDiscountAdjustmentFilterJSON struct {
	Field       apijson.Field
	Operator    apijson.Field
	Values      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *PlanExternalPlanIDVersionGetResponseAdjustmentsPlanPhaseAmountDiscountAdjustmentFilter) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r planExternalPlanIDVersionGetResponseAdjustmentsPlanPhaseAmountDiscountAdjustmentFilterJSON) RawJSON() string {
	return r.raw
}

// The property of the price to filter on.
type PlanExternalPlanIDVersionGetResponseAdjustmentsPlanPhaseAmountDiscountAdjustmentFiltersField string

const (
	PlanExternalPlanIDVersionGetResponseAdjustmentsPlanPhaseAmountDiscountAdjustmentFiltersFieldPriceID       PlanExternalPlanIDVersionGetResponseAdjustmentsPlanPhaseAmountDiscountAdjustmentFiltersField = "price_id"
	PlanExternalPlanIDVersionGetResponseAdjustmentsPlanPhaseAmountDiscountAdjustmentFiltersFieldItemID        PlanExternalPlanIDVersionGetResponseAdjustmentsPlanPhaseAmountDiscountAdjustmentFiltersField = "item_id"
	PlanExternalPlanIDVersionGetResponseAdjustmentsPlanPhaseAmountDiscountAdjustmentFiltersFieldPriceType     PlanExternalPlanIDVersionGetResponseAdjustmentsPlanPhaseAmountDiscountAdjustmentFiltersField = "price_type"
	PlanExternalPlanIDVersionGetResponseAdjustmentsPlanPhaseAmountDiscountAdjustmentFiltersFieldCurrency      PlanExternalPlanIDVersionGetResponseAdjustmentsPlanPhaseAmountDiscountAdjustmentFiltersField = "currency"
	PlanExternalPlanIDVersionGetResponseAdjustmentsPlanPhaseAmountDiscountAdjustmentFiltersFieldPricingUnitID PlanExternalPlanIDVersionGetResponseAdjustmentsPlanPhaseAmountDiscountAdjustmentFiltersField = "pricing_unit_id"
)

func (r PlanExternalPlanIDVersionGetResponseAdjustmentsPlanPhaseAmountDiscountAdjustmentFiltersField) IsKnown() bool {
	switch r {
	case PlanExternalPlanIDVersionGetResponseAdjustmentsPlanPhaseAmountDiscountAdjustmentFiltersFieldPriceID, PlanExternalPlanIDVersionGetResponseAdjustmentsPlanPhaseAmountDiscountAdjustmentFiltersFieldItemID, PlanExternalPlanIDVersionGetResponseAdjustmentsPlanPhaseAmountDiscountAdjustmentFiltersFieldPriceType, PlanExternalPlanIDVersionGetResponseAdjustmentsPlanPhaseAmountDiscountAdjustmentFiltersFieldCurrency, PlanExternalPlanIDVersionGetResponseAdjustmentsPlanPhaseAmountDiscountAdjustmentFiltersFieldPricingUnitID:
		return true
	}
	return false
}

// Should prices that match the filter be included or excluded.
type PlanExternalPlanIDVersionGetResponseAdjustmentsPlanPhaseAmountDiscountAdjustmentFiltersOperator string

const (
	PlanExternalPlanIDVersionGetResponseAdjustmentsPlanPhaseAmountDiscountAdjustmentFiltersOperatorIncludes PlanExternalPlanIDVersionGetResponseAdjustmentsPlanPhaseAmountDiscountAdjustmentFiltersOperator = "includes"
	PlanExternalPlanIDVersionGetResponseAdjustmentsPlanPhaseAmountDiscountAdjustmentFiltersOperatorExcludes PlanExternalPlanIDVersionGetResponseAdjustmentsPlanPhaseAmountDiscountAdjustmentFiltersOperator = "excludes"
)

func (r PlanExternalPlanIDVersionGetResponseAdjustmentsPlanPhaseAmountDiscountAdjustmentFiltersOperator) IsKnown() bool {
	switch r {
	case PlanExternalPlanIDVersionGetResponseAdjustmentsPlanPhaseAmountDiscountAdjustmentFiltersOperatorIncludes, PlanExternalPlanIDVersionGetResponseAdjustmentsPlanPhaseAmountDiscountAdjustmentFiltersOperatorExcludes:
		return true
	}
	return false
}

type PlanExternalPlanIDVersionGetResponseAdjustmentsPlanPhasePercentageDiscountAdjustment struct {
	ID             string                                                                                             `json:"id,required"`
	AdjustmentType PlanExternalPlanIDVersionGetResponseAdjustmentsPlanPhasePercentageDiscountAdjustmentAdjustmentType `json:"adjustment_type,required"`
	// The price IDs that this adjustment applies to.
	//
	// Deprecated: deprecated
	AppliesToPriceIDs []string `json:"applies_to_price_ids,required"`
	// The filters that determine which prices to apply this adjustment to.
	Filters []PlanExternalPlanIDVersionGetResponseAdjustmentsPlanPhasePercentageDiscountAdjustmentFilter `json:"filters,required"`
	// True for adjustments that apply to an entire invocice, false for adjustments
	// that apply to only one price.
	IsInvoiceLevel bool `json:"is_invoice_level,required"`
	// The percentage (as a value between 0 and 1) by which to discount the price
	// intervals this adjustment applies to in a given billing period.
	PercentageDiscount float64 `json:"percentage_discount,required"`
	// The plan phase in which this adjustment is active.
	PlanPhaseOrder int64 `json:"plan_phase_order,required,nullable"`
	// The reason for the adjustment.
	Reason string                                                                                   `json:"reason,required,nullable"`
	JSON   planExternalPlanIDVersionGetResponseAdjustmentsPlanPhasePercentageDiscountAdjustmentJSON `json:"-"`
}

// planExternalPlanIDVersionGetResponseAdjustmentsPlanPhasePercentageDiscountAdjustmentJSON
// contains the JSON metadata for the struct
// [PlanExternalPlanIDVersionGetResponseAdjustmentsPlanPhasePercentageDiscountAdjustment]
type planExternalPlanIDVersionGetResponseAdjustmentsPlanPhasePercentageDiscountAdjustmentJSON struct {
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

func (r *PlanExternalPlanIDVersionGetResponseAdjustmentsPlanPhasePercentageDiscountAdjustment) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r planExternalPlanIDVersionGetResponseAdjustmentsPlanPhasePercentageDiscountAdjustmentJSON) RawJSON() string {
	return r.raw
}

func (r PlanExternalPlanIDVersionGetResponseAdjustmentsPlanPhasePercentageDiscountAdjustment) implementsPlanExternalPlanIDVersionGetResponseAdjustment() {
}

type PlanExternalPlanIDVersionGetResponseAdjustmentsPlanPhasePercentageDiscountAdjustmentAdjustmentType string

const (
	PlanExternalPlanIDVersionGetResponseAdjustmentsPlanPhasePercentageDiscountAdjustmentAdjustmentTypePercentageDiscount PlanExternalPlanIDVersionGetResponseAdjustmentsPlanPhasePercentageDiscountAdjustmentAdjustmentType = "percentage_discount"
)

func (r PlanExternalPlanIDVersionGetResponseAdjustmentsPlanPhasePercentageDiscountAdjustmentAdjustmentType) IsKnown() bool {
	switch r {
	case PlanExternalPlanIDVersionGetResponseAdjustmentsPlanPhasePercentageDiscountAdjustmentAdjustmentTypePercentageDiscount:
		return true
	}
	return false
}

type PlanExternalPlanIDVersionGetResponseAdjustmentsPlanPhasePercentageDiscountAdjustmentFilter struct {
	// The property of the price to filter on.
	Field PlanExternalPlanIDVersionGetResponseAdjustmentsPlanPhasePercentageDiscountAdjustmentFiltersField `json:"field,required"`
	// Should prices that match the filter be included or excluded.
	Operator PlanExternalPlanIDVersionGetResponseAdjustmentsPlanPhasePercentageDiscountAdjustmentFiltersOperator `json:"operator,required"`
	// The IDs or values that match this filter.
	Values []string                                                                                       `json:"values,required"`
	JSON   planExternalPlanIDVersionGetResponseAdjustmentsPlanPhasePercentageDiscountAdjustmentFilterJSON `json:"-"`
}

// planExternalPlanIDVersionGetResponseAdjustmentsPlanPhasePercentageDiscountAdjustmentFilterJSON
// contains the JSON metadata for the struct
// [PlanExternalPlanIDVersionGetResponseAdjustmentsPlanPhasePercentageDiscountAdjustmentFilter]
type planExternalPlanIDVersionGetResponseAdjustmentsPlanPhasePercentageDiscountAdjustmentFilterJSON struct {
	Field       apijson.Field
	Operator    apijson.Field
	Values      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *PlanExternalPlanIDVersionGetResponseAdjustmentsPlanPhasePercentageDiscountAdjustmentFilter) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r planExternalPlanIDVersionGetResponseAdjustmentsPlanPhasePercentageDiscountAdjustmentFilterJSON) RawJSON() string {
	return r.raw
}

// The property of the price to filter on.
type PlanExternalPlanIDVersionGetResponseAdjustmentsPlanPhasePercentageDiscountAdjustmentFiltersField string

const (
	PlanExternalPlanIDVersionGetResponseAdjustmentsPlanPhasePercentageDiscountAdjustmentFiltersFieldPriceID       PlanExternalPlanIDVersionGetResponseAdjustmentsPlanPhasePercentageDiscountAdjustmentFiltersField = "price_id"
	PlanExternalPlanIDVersionGetResponseAdjustmentsPlanPhasePercentageDiscountAdjustmentFiltersFieldItemID        PlanExternalPlanIDVersionGetResponseAdjustmentsPlanPhasePercentageDiscountAdjustmentFiltersField = "item_id"
	PlanExternalPlanIDVersionGetResponseAdjustmentsPlanPhasePercentageDiscountAdjustmentFiltersFieldPriceType     PlanExternalPlanIDVersionGetResponseAdjustmentsPlanPhasePercentageDiscountAdjustmentFiltersField = "price_type"
	PlanExternalPlanIDVersionGetResponseAdjustmentsPlanPhasePercentageDiscountAdjustmentFiltersFieldCurrency      PlanExternalPlanIDVersionGetResponseAdjustmentsPlanPhasePercentageDiscountAdjustmentFiltersField = "currency"
	PlanExternalPlanIDVersionGetResponseAdjustmentsPlanPhasePercentageDiscountAdjustmentFiltersFieldPricingUnitID PlanExternalPlanIDVersionGetResponseAdjustmentsPlanPhasePercentageDiscountAdjustmentFiltersField = "pricing_unit_id"
)

func (r PlanExternalPlanIDVersionGetResponseAdjustmentsPlanPhasePercentageDiscountAdjustmentFiltersField) IsKnown() bool {
	switch r {
	case PlanExternalPlanIDVersionGetResponseAdjustmentsPlanPhasePercentageDiscountAdjustmentFiltersFieldPriceID, PlanExternalPlanIDVersionGetResponseAdjustmentsPlanPhasePercentageDiscountAdjustmentFiltersFieldItemID, PlanExternalPlanIDVersionGetResponseAdjustmentsPlanPhasePercentageDiscountAdjustmentFiltersFieldPriceType, PlanExternalPlanIDVersionGetResponseAdjustmentsPlanPhasePercentageDiscountAdjustmentFiltersFieldCurrency, PlanExternalPlanIDVersionGetResponseAdjustmentsPlanPhasePercentageDiscountAdjustmentFiltersFieldPricingUnitID:
		return true
	}
	return false
}

// Should prices that match the filter be included or excluded.
type PlanExternalPlanIDVersionGetResponseAdjustmentsPlanPhasePercentageDiscountAdjustmentFiltersOperator string

const (
	PlanExternalPlanIDVersionGetResponseAdjustmentsPlanPhasePercentageDiscountAdjustmentFiltersOperatorIncludes PlanExternalPlanIDVersionGetResponseAdjustmentsPlanPhasePercentageDiscountAdjustmentFiltersOperator = "includes"
	PlanExternalPlanIDVersionGetResponseAdjustmentsPlanPhasePercentageDiscountAdjustmentFiltersOperatorExcludes PlanExternalPlanIDVersionGetResponseAdjustmentsPlanPhasePercentageDiscountAdjustmentFiltersOperator = "excludes"
)

func (r PlanExternalPlanIDVersionGetResponseAdjustmentsPlanPhasePercentageDiscountAdjustmentFiltersOperator) IsKnown() bool {
	switch r {
	case PlanExternalPlanIDVersionGetResponseAdjustmentsPlanPhasePercentageDiscountAdjustmentFiltersOperatorIncludes, PlanExternalPlanIDVersionGetResponseAdjustmentsPlanPhasePercentageDiscountAdjustmentFiltersOperatorExcludes:
		return true
	}
	return false
}

type PlanExternalPlanIDVersionGetResponseAdjustmentsPlanPhaseMinimumAdjustment struct {
	ID             string                                                                                  `json:"id,required"`
	AdjustmentType PlanExternalPlanIDVersionGetResponseAdjustmentsPlanPhaseMinimumAdjustmentAdjustmentType `json:"adjustment_type,required"`
	// The price IDs that this adjustment applies to.
	//
	// Deprecated: deprecated
	AppliesToPriceIDs []string `json:"applies_to_price_ids,required"`
	// The filters that determine which prices to apply this adjustment to.
	Filters []PlanExternalPlanIDVersionGetResponseAdjustmentsPlanPhaseMinimumAdjustmentFilter `json:"filters,required"`
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
	Reason string                                                                        `json:"reason,required,nullable"`
	JSON   planExternalPlanIDVersionGetResponseAdjustmentsPlanPhaseMinimumAdjustmentJSON `json:"-"`
}

// planExternalPlanIDVersionGetResponseAdjustmentsPlanPhaseMinimumAdjustmentJSON
// contains the JSON metadata for the struct
// [PlanExternalPlanIDVersionGetResponseAdjustmentsPlanPhaseMinimumAdjustment]
type planExternalPlanIDVersionGetResponseAdjustmentsPlanPhaseMinimumAdjustmentJSON struct {
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

func (r *PlanExternalPlanIDVersionGetResponseAdjustmentsPlanPhaseMinimumAdjustment) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r planExternalPlanIDVersionGetResponseAdjustmentsPlanPhaseMinimumAdjustmentJSON) RawJSON() string {
	return r.raw
}

func (r PlanExternalPlanIDVersionGetResponseAdjustmentsPlanPhaseMinimumAdjustment) implementsPlanExternalPlanIDVersionGetResponseAdjustment() {
}

type PlanExternalPlanIDVersionGetResponseAdjustmentsPlanPhaseMinimumAdjustmentAdjustmentType string

const (
	PlanExternalPlanIDVersionGetResponseAdjustmentsPlanPhaseMinimumAdjustmentAdjustmentTypeMinimum PlanExternalPlanIDVersionGetResponseAdjustmentsPlanPhaseMinimumAdjustmentAdjustmentType = "minimum"
)

func (r PlanExternalPlanIDVersionGetResponseAdjustmentsPlanPhaseMinimumAdjustmentAdjustmentType) IsKnown() bool {
	switch r {
	case PlanExternalPlanIDVersionGetResponseAdjustmentsPlanPhaseMinimumAdjustmentAdjustmentTypeMinimum:
		return true
	}
	return false
}

type PlanExternalPlanIDVersionGetResponseAdjustmentsPlanPhaseMinimumAdjustmentFilter struct {
	// The property of the price to filter on.
	Field PlanExternalPlanIDVersionGetResponseAdjustmentsPlanPhaseMinimumAdjustmentFiltersField `json:"field,required"`
	// Should prices that match the filter be included or excluded.
	Operator PlanExternalPlanIDVersionGetResponseAdjustmentsPlanPhaseMinimumAdjustmentFiltersOperator `json:"operator,required"`
	// The IDs or values that match this filter.
	Values []string                                                                            `json:"values,required"`
	JSON   planExternalPlanIDVersionGetResponseAdjustmentsPlanPhaseMinimumAdjustmentFilterJSON `json:"-"`
}

// planExternalPlanIDVersionGetResponseAdjustmentsPlanPhaseMinimumAdjustmentFilterJSON
// contains the JSON metadata for the struct
// [PlanExternalPlanIDVersionGetResponseAdjustmentsPlanPhaseMinimumAdjustmentFilter]
type planExternalPlanIDVersionGetResponseAdjustmentsPlanPhaseMinimumAdjustmentFilterJSON struct {
	Field       apijson.Field
	Operator    apijson.Field
	Values      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *PlanExternalPlanIDVersionGetResponseAdjustmentsPlanPhaseMinimumAdjustmentFilter) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r planExternalPlanIDVersionGetResponseAdjustmentsPlanPhaseMinimumAdjustmentFilterJSON) RawJSON() string {
	return r.raw
}

// The property of the price to filter on.
type PlanExternalPlanIDVersionGetResponseAdjustmentsPlanPhaseMinimumAdjustmentFiltersField string

const (
	PlanExternalPlanIDVersionGetResponseAdjustmentsPlanPhaseMinimumAdjustmentFiltersFieldPriceID       PlanExternalPlanIDVersionGetResponseAdjustmentsPlanPhaseMinimumAdjustmentFiltersField = "price_id"
	PlanExternalPlanIDVersionGetResponseAdjustmentsPlanPhaseMinimumAdjustmentFiltersFieldItemID        PlanExternalPlanIDVersionGetResponseAdjustmentsPlanPhaseMinimumAdjustmentFiltersField = "item_id"
	PlanExternalPlanIDVersionGetResponseAdjustmentsPlanPhaseMinimumAdjustmentFiltersFieldPriceType     PlanExternalPlanIDVersionGetResponseAdjustmentsPlanPhaseMinimumAdjustmentFiltersField = "price_type"
	PlanExternalPlanIDVersionGetResponseAdjustmentsPlanPhaseMinimumAdjustmentFiltersFieldCurrency      PlanExternalPlanIDVersionGetResponseAdjustmentsPlanPhaseMinimumAdjustmentFiltersField = "currency"
	PlanExternalPlanIDVersionGetResponseAdjustmentsPlanPhaseMinimumAdjustmentFiltersFieldPricingUnitID PlanExternalPlanIDVersionGetResponseAdjustmentsPlanPhaseMinimumAdjustmentFiltersField = "pricing_unit_id"
)

func (r PlanExternalPlanIDVersionGetResponseAdjustmentsPlanPhaseMinimumAdjustmentFiltersField) IsKnown() bool {
	switch r {
	case PlanExternalPlanIDVersionGetResponseAdjustmentsPlanPhaseMinimumAdjustmentFiltersFieldPriceID, PlanExternalPlanIDVersionGetResponseAdjustmentsPlanPhaseMinimumAdjustmentFiltersFieldItemID, PlanExternalPlanIDVersionGetResponseAdjustmentsPlanPhaseMinimumAdjustmentFiltersFieldPriceType, PlanExternalPlanIDVersionGetResponseAdjustmentsPlanPhaseMinimumAdjustmentFiltersFieldCurrency, PlanExternalPlanIDVersionGetResponseAdjustmentsPlanPhaseMinimumAdjustmentFiltersFieldPricingUnitID:
		return true
	}
	return false
}

// Should prices that match the filter be included or excluded.
type PlanExternalPlanIDVersionGetResponseAdjustmentsPlanPhaseMinimumAdjustmentFiltersOperator string

const (
	PlanExternalPlanIDVersionGetResponseAdjustmentsPlanPhaseMinimumAdjustmentFiltersOperatorIncludes PlanExternalPlanIDVersionGetResponseAdjustmentsPlanPhaseMinimumAdjustmentFiltersOperator = "includes"
	PlanExternalPlanIDVersionGetResponseAdjustmentsPlanPhaseMinimumAdjustmentFiltersOperatorExcludes PlanExternalPlanIDVersionGetResponseAdjustmentsPlanPhaseMinimumAdjustmentFiltersOperator = "excludes"
)

func (r PlanExternalPlanIDVersionGetResponseAdjustmentsPlanPhaseMinimumAdjustmentFiltersOperator) IsKnown() bool {
	switch r {
	case PlanExternalPlanIDVersionGetResponseAdjustmentsPlanPhaseMinimumAdjustmentFiltersOperatorIncludes, PlanExternalPlanIDVersionGetResponseAdjustmentsPlanPhaseMinimumAdjustmentFiltersOperatorExcludes:
		return true
	}
	return false
}

type PlanExternalPlanIDVersionGetResponseAdjustmentsPlanPhaseMaximumAdjustment struct {
	ID             string                                                                                  `json:"id,required"`
	AdjustmentType PlanExternalPlanIDVersionGetResponseAdjustmentsPlanPhaseMaximumAdjustmentAdjustmentType `json:"adjustment_type,required"`
	// The price IDs that this adjustment applies to.
	//
	// Deprecated: deprecated
	AppliesToPriceIDs []string `json:"applies_to_price_ids,required"`
	// The filters that determine which prices to apply this adjustment to.
	Filters []PlanExternalPlanIDVersionGetResponseAdjustmentsPlanPhaseMaximumAdjustmentFilter `json:"filters,required"`
	// True for adjustments that apply to an entire invocice, false for adjustments
	// that apply to only one price.
	IsInvoiceLevel bool `json:"is_invoice_level,required"`
	// The maximum amount to charge in a given billing period for the prices this
	// adjustment applies to.
	MaximumAmount string `json:"maximum_amount,required"`
	// The plan phase in which this adjustment is active.
	PlanPhaseOrder int64 `json:"plan_phase_order,required,nullable"`
	// The reason for the adjustment.
	Reason string                                                                        `json:"reason,required,nullable"`
	JSON   planExternalPlanIDVersionGetResponseAdjustmentsPlanPhaseMaximumAdjustmentJSON `json:"-"`
}

// planExternalPlanIDVersionGetResponseAdjustmentsPlanPhaseMaximumAdjustmentJSON
// contains the JSON metadata for the struct
// [PlanExternalPlanIDVersionGetResponseAdjustmentsPlanPhaseMaximumAdjustment]
type planExternalPlanIDVersionGetResponseAdjustmentsPlanPhaseMaximumAdjustmentJSON struct {
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

func (r *PlanExternalPlanIDVersionGetResponseAdjustmentsPlanPhaseMaximumAdjustment) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r planExternalPlanIDVersionGetResponseAdjustmentsPlanPhaseMaximumAdjustmentJSON) RawJSON() string {
	return r.raw
}

func (r PlanExternalPlanIDVersionGetResponseAdjustmentsPlanPhaseMaximumAdjustment) implementsPlanExternalPlanIDVersionGetResponseAdjustment() {
}

type PlanExternalPlanIDVersionGetResponseAdjustmentsPlanPhaseMaximumAdjustmentAdjustmentType string

const (
	PlanExternalPlanIDVersionGetResponseAdjustmentsPlanPhaseMaximumAdjustmentAdjustmentTypeMaximum PlanExternalPlanIDVersionGetResponseAdjustmentsPlanPhaseMaximumAdjustmentAdjustmentType = "maximum"
)

func (r PlanExternalPlanIDVersionGetResponseAdjustmentsPlanPhaseMaximumAdjustmentAdjustmentType) IsKnown() bool {
	switch r {
	case PlanExternalPlanIDVersionGetResponseAdjustmentsPlanPhaseMaximumAdjustmentAdjustmentTypeMaximum:
		return true
	}
	return false
}

type PlanExternalPlanIDVersionGetResponseAdjustmentsPlanPhaseMaximumAdjustmentFilter struct {
	// The property of the price to filter on.
	Field PlanExternalPlanIDVersionGetResponseAdjustmentsPlanPhaseMaximumAdjustmentFiltersField `json:"field,required"`
	// Should prices that match the filter be included or excluded.
	Operator PlanExternalPlanIDVersionGetResponseAdjustmentsPlanPhaseMaximumAdjustmentFiltersOperator `json:"operator,required"`
	// The IDs or values that match this filter.
	Values []string                                                                            `json:"values,required"`
	JSON   planExternalPlanIDVersionGetResponseAdjustmentsPlanPhaseMaximumAdjustmentFilterJSON `json:"-"`
}

// planExternalPlanIDVersionGetResponseAdjustmentsPlanPhaseMaximumAdjustmentFilterJSON
// contains the JSON metadata for the struct
// [PlanExternalPlanIDVersionGetResponseAdjustmentsPlanPhaseMaximumAdjustmentFilter]
type planExternalPlanIDVersionGetResponseAdjustmentsPlanPhaseMaximumAdjustmentFilterJSON struct {
	Field       apijson.Field
	Operator    apijson.Field
	Values      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *PlanExternalPlanIDVersionGetResponseAdjustmentsPlanPhaseMaximumAdjustmentFilter) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r planExternalPlanIDVersionGetResponseAdjustmentsPlanPhaseMaximumAdjustmentFilterJSON) RawJSON() string {
	return r.raw
}

// The property of the price to filter on.
type PlanExternalPlanIDVersionGetResponseAdjustmentsPlanPhaseMaximumAdjustmentFiltersField string

const (
	PlanExternalPlanIDVersionGetResponseAdjustmentsPlanPhaseMaximumAdjustmentFiltersFieldPriceID       PlanExternalPlanIDVersionGetResponseAdjustmentsPlanPhaseMaximumAdjustmentFiltersField = "price_id"
	PlanExternalPlanIDVersionGetResponseAdjustmentsPlanPhaseMaximumAdjustmentFiltersFieldItemID        PlanExternalPlanIDVersionGetResponseAdjustmentsPlanPhaseMaximumAdjustmentFiltersField = "item_id"
	PlanExternalPlanIDVersionGetResponseAdjustmentsPlanPhaseMaximumAdjustmentFiltersFieldPriceType     PlanExternalPlanIDVersionGetResponseAdjustmentsPlanPhaseMaximumAdjustmentFiltersField = "price_type"
	PlanExternalPlanIDVersionGetResponseAdjustmentsPlanPhaseMaximumAdjustmentFiltersFieldCurrency      PlanExternalPlanIDVersionGetResponseAdjustmentsPlanPhaseMaximumAdjustmentFiltersField = "currency"
	PlanExternalPlanIDVersionGetResponseAdjustmentsPlanPhaseMaximumAdjustmentFiltersFieldPricingUnitID PlanExternalPlanIDVersionGetResponseAdjustmentsPlanPhaseMaximumAdjustmentFiltersField = "pricing_unit_id"
)

func (r PlanExternalPlanIDVersionGetResponseAdjustmentsPlanPhaseMaximumAdjustmentFiltersField) IsKnown() bool {
	switch r {
	case PlanExternalPlanIDVersionGetResponseAdjustmentsPlanPhaseMaximumAdjustmentFiltersFieldPriceID, PlanExternalPlanIDVersionGetResponseAdjustmentsPlanPhaseMaximumAdjustmentFiltersFieldItemID, PlanExternalPlanIDVersionGetResponseAdjustmentsPlanPhaseMaximumAdjustmentFiltersFieldPriceType, PlanExternalPlanIDVersionGetResponseAdjustmentsPlanPhaseMaximumAdjustmentFiltersFieldCurrency, PlanExternalPlanIDVersionGetResponseAdjustmentsPlanPhaseMaximumAdjustmentFiltersFieldPricingUnitID:
		return true
	}
	return false
}

// Should prices that match the filter be included or excluded.
type PlanExternalPlanIDVersionGetResponseAdjustmentsPlanPhaseMaximumAdjustmentFiltersOperator string

const (
	PlanExternalPlanIDVersionGetResponseAdjustmentsPlanPhaseMaximumAdjustmentFiltersOperatorIncludes PlanExternalPlanIDVersionGetResponseAdjustmentsPlanPhaseMaximumAdjustmentFiltersOperator = "includes"
	PlanExternalPlanIDVersionGetResponseAdjustmentsPlanPhaseMaximumAdjustmentFiltersOperatorExcludes PlanExternalPlanIDVersionGetResponseAdjustmentsPlanPhaseMaximumAdjustmentFiltersOperator = "excludes"
)

func (r PlanExternalPlanIDVersionGetResponseAdjustmentsPlanPhaseMaximumAdjustmentFiltersOperator) IsKnown() bool {
	switch r {
	case PlanExternalPlanIDVersionGetResponseAdjustmentsPlanPhaseMaximumAdjustmentFiltersOperatorIncludes, PlanExternalPlanIDVersionGetResponseAdjustmentsPlanPhaseMaximumAdjustmentFiltersOperatorExcludes:
		return true
	}
	return false
}

type PlanExternalPlanIDVersionGetResponseAdjustmentsAdjustmentType string

const (
	PlanExternalPlanIDVersionGetResponseAdjustmentsAdjustmentTypeUsageDiscount      PlanExternalPlanIDVersionGetResponseAdjustmentsAdjustmentType = "usage_discount"
	PlanExternalPlanIDVersionGetResponseAdjustmentsAdjustmentTypeAmountDiscount     PlanExternalPlanIDVersionGetResponseAdjustmentsAdjustmentType = "amount_discount"
	PlanExternalPlanIDVersionGetResponseAdjustmentsAdjustmentTypePercentageDiscount PlanExternalPlanIDVersionGetResponseAdjustmentsAdjustmentType = "percentage_discount"
	PlanExternalPlanIDVersionGetResponseAdjustmentsAdjustmentTypeMinimum            PlanExternalPlanIDVersionGetResponseAdjustmentsAdjustmentType = "minimum"
	PlanExternalPlanIDVersionGetResponseAdjustmentsAdjustmentTypeMaximum            PlanExternalPlanIDVersionGetResponseAdjustmentsAdjustmentType = "maximum"
)

func (r PlanExternalPlanIDVersionGetResponseAdjustmentsAdjustmentType) IsKnown() bool {
	switch r {
	case PlanExternalPlanIDVersionGetResponseAdjustmentsAdjustmentTypeUsageDiscount, PlanExternalPlanIDVersionGetResponseAdjustmentsAdjustmentTypeAmountDiscount, PlanExternalPlanIDVersionGetResponseAdjustmentsAdjustmentTypePercentageDiscount, PlanExternalPlanIDVersionGetResponseAdjustmentsAdjustmentTypeMinimum, PlanExternalPlanIDVersionGetResponseAdjustmentsAdjustmentTypeMaximum:
		return true
	}
	return false
}

type PlanExternalPlanIDVersionGetResponsePlanPhase struct {
	ID          string `json:"id,required"`
	Description string `json:"description,required,nullable"`
	// How many terms of length `duration_unit` this phase is active for. If null, this
	// phase is evergreen and active indefinitely
	Duration     int64                                                      `json:"duration,required,nullable"`
	DurationUnit PlanExternalPlanIDVersionGetResponsePlanPhasesDurationUnit `json:"duration_unit,required,nullable"`
	Name         string                                                     `json:"name,required"`
	// Determines the ordering of the phase in a plan's lifecycle. 1 = first phase.
	Order int64                                             `json:"order,required"`
	JSON  planExternalPlanIDVersionGetResponsePlanPhaseJSON `json:"-"`
}

// planExternalPlanIDVersionGetResponsePlanPhaseJSON contains the JSON metadata for
// the struct [PlanExternalPlanIDVersionGetResponsePlanPhase]
type planExternalPlanIDVersionGetResponsePlanPhaseJSON struct {
	ID           apijson.Field
	Description  apijson.Field
	Duration     apijson.Field
	DurationUnit apijson.Field
	Name         apijson.Field
	Order        apijson.Field
	raw          string
	ExtraFields  map[string]apijson.Field
}

func (r *PlanExternalPlanIDVersionGetResponsePlanPhase) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r planExternalPlanIDVersionGetResponsePlanPhaseJSON) RawJSON() string {
	return r.raw
}

type PlanExternalPlanIDVersionGetResponsePlanPhasesDurationUnit string

const (
	PlanExternalPlanIDVersionGetResponsePlanPhasesDurationUnitDaily      PlanExternalPlanIDVersionGetResponsePlanPhasesDurationUnit = "daily"
	PlanExternalPlanIDVersionGetResponsePlanPhasesDurationUnitMonthly    PlanExternalPlanIDVersionGetResponsePlanPhasesDurationUnit = "monthly"
	PlanExternalPlanIDVersionGetResponsePlanPhasesDurationUnitQuarterly  PlanExternalPlanIDVersionGetResponsePlanPhasesDurationUnit = "quarterly"
	PlanExternalPlanIDVersionGetResponsePlanPhasesDurationUnitSemiAnnual PlanExternalPlanIDVersionGetResponsePlanPhasesDurationUnit = "semi_annual"
	PlanExternalPlanIDVersionGetResponsePlanPhasesDurationUnitAnnual     PlanExternalPlanIDVersionGetResponsePlanPhasesDurationUnit = "annual"
)

func (r PlanExternalPlanIDVersionGetResponsePlanPhasesDurationUnit) IsKnown() bool {
	switch r {
	case PlanExternalPlanIDVersionGetResponsePlanPhasesDurationUnitDaily, PlanExternalPlanIDVersionGetResponsePlanPhasesDurationUnitMonthly, PlanExternalPlanIDVersionGetResponsePlanPhasesDurationUnitQuarterly, PlanExternalPlanIDVersionGetResponsePlanPhasesDurationUnitSemiAnnual, PlanExternalPlanIDVersionGetResponsePlanPhasesDurationUnitAnnual:
		return true
	}
	return false
}

type PlanExternalPlanIDVersionNewParams struct {
	// New version number.
	Version param.Field[int64] `json:"version,required"`
	// Additional adjustments to be added to the plan.
	AddAdjustments param.Field[[]PlanExternalPlanIDVersionNewParamsAddAdjustment] `json:"add_adjustments"`
	// Additional prices to be added to the plan.
	AddPrices param.Field[[]PlanExternalPlanIDVersionNewParamsAddPrice] `json:"add_prices"`
	// Adjustments to be removed from the plan.
	RemoveAdjustments param.Field[[]PlanExternalPlanIDVersionNewParamsRemoveAdjustment] `json:"remove_adjustments"`
	// Prices to be removed from the plan.
	RemovePrices param.Field[[]PlanExternalPlanIDVersionNewParamsRemovePrice] `json:"remove_prices"`
	// Adjustments to be replaced with additional adjustments on the plan.
	ReplaceAdjustments param.Field[[]PlanExternalPlanIDVersionNewParamsReplaceAdjustment] `json:"replace_adjustments"`
	// Prices to be replaced with additional prices on the plan.
	ReplacePrices param.Field[[]PlanExternalPlanIDVersionNewParamsReplacePrice] `json:"replace_prices"`
	// Set this new plan version as the default
	SetAsDefault param.Field[bool] `json:"set_as_default"`
}

func (r PlanExternalPlanIDVersionNewParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type PlanExternalPlanIDVersionNewParamsAddAdjustment struct {
	// The definition of a new adjustment to create and add to the plan.
	Adjustment param.Field[PlanExternalPlanIDVersionNewParamsAddAdjustmentsAdjustmentUnion] `json:"adjustment,required"`
	// The phase to add this adjustment to.
	PlanPhaseOrder param.Field[int64] `json:"plan_phase_order"`
}

func (r PlanExternalPlanIDVersionNewParamsAddAdjustment) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// The definition of a new adjustment to create and add to the plan.
type PlanExternalPlanIDVersionNewParamsAddAdjustmentsAdjustment struct {
	AdjustmentType    param.Field[PlanExternalPlanIDVersionNewParamsAddAdjustmentsAdjustmentAdjustmentType] `json:"adjustment_type,required"`
	AmountDiscount    param.Field[string]                                                                   `json:"amount_discount"`
	AppliesToPriceIDs param.Field[interface{}]                                                              `json:"applies_to_price_ids"`
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

func (r PlanExternalPlanIDVersionNewParamsAddAdjustmentsAdjustment) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r PlanExternalPlanIDVersionNewParamsAddAdjustmentsAdjustment) implementsPlanExternalPlanIDVersionNewParamsAddAdjustmentsAdjustmentUnion() {
}

// The definition of a new adjustment to create and add to the plan.
//
// Satisfied by
// [PlanExternalPlanIDVersionNewParamsAddAdjustmentsAdjustmentNewPercentageDiscount],
// [PlanExternalPlanIDVersionNewParamsAddAdjustmentsAdjustmentNewUsageDiscount],
// [PlanExternalPlanIDVersionNewParamsAddAdjustmentsAdjustmentNewAmountDiscount],
// [PlanExternalPlanIDVersionNewParamsAddAdjustmentsAdjustmentNewMinimum],
// [PlanExternalPlanIDVersionNewParamsAddAdjustmentsAdjustmentNewMaximum],
// [PlanExternalPlanIDVersionNewParamsAddAdjustmentsAdjustment].
type PlanExternalPlanIDVersionNewParamsAddAdjustmentsAdjustmentUnion interface {
	implementsPlanExternalPlanIDVersionNewParamsAddAdjustmentsAdjustmentUnion()
}

type PlanExternalPlanIDVersionNewParamsAddAdjustmentsAdjustmentNewPercentageDiscount struct {
	AdjustmentType     param.Field[PlanExternalPlanIDVersionNewParamsAddAdjustmentsAdjustmentNewPercentageDiscountAdjustmentType] `json:"adjustment_type,required"`
	PercentageDiscount param.Field[float64]                                                                                       `json:"percentage_discount,required"`
	// The set of price IDs to which this adjustment applies.
	AppliesToPriceIDs param.Field[[]string] `json:"applies_to_price_ids"`
	// When false, this adjustment will be applied to a single price. Otherwise, it
	// will be applied at the invoice level, possibly to multiple prices.
	IsInvoiceLevel param.Field[bool] `json:"is_invoice_level"`
}

func (r PlanExternalPlanIDVersionNewParamsAddAdjustmentsAdjustmentNewPercentageDiscount) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r PlanExternalPlanIDVersionNewParamsAddAdjustmentsAdjustmentNewPercentageDiscount) implementsPlanExternalPlanIDVersionNewParamsAddAdjustmentsAdjustmentUnion() {
}

type PlanExternalPlanIDVersionNewParamsAddAdjustmentsAdjustmentNewPercentageDiscountAdjustmentType string

const (
	PlanExternalPlanIDVersionNewParamsAddAdjustmentsAdjustmentNewPercentageDiscountAdjustmentTypePercentageDiscount PlanExternalPlanIDVersionNewParamsAddAdjustmentsAdjustmentNewPercentageDiscountAdjustmentType = "percentage_discount"
)

func (r PlanExternalPlanIDVersionNewParamsAddAdjustmentsAdjustmentNewPercentageDiscountAdjustmentType) IsKnown() bool {
	switch r {
	case PlanExternalPlanIDVersionNewParamsAddAdjustmentsAdjustmentNewPercentageDiscountAdjustmentTypePercentageDiscount:
		return true
	}
	return false
}

type PlanExternalPlanIDVersionNewParamsAddAdjustmentsAdjustmentNewUsageDiscount struct {
	AdjustmentType param.Field[PlanExternalPlanIDVersionNewParamsAddAdjustmentsAdjustmentNewUsageDiscountAdjustmentType] `json:"adjustment_type,required"`
	UsageDiscount  param.Field[float64]                                                                                  `json:"usage_discount,required"`
	// The set of price IDs to which this adjustment applies.
	AppliesToPriceIDs param.Field[[]string] `json:"applies_to_price_ids"`
	// When false, this adjustment will be applied to a single price. Otherwise, it
	// will be applied at the invoice level, possibly to multiple prices.
	IsInvoiceLevel param.Field[bool] `json:"is_invoice_level"`
}

func (r PlanExternalPlanIDVersionNewParamsAddAdjustmentsAdjustmentNewUsageDiscount) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r PlanExternalPlanIDVersionNewParamsAddAdjustmentsAdjustmentNewUsageDiscount) implementsPlanExternalPlanIDVersionNewParamsAddAdjustmentsAdjustmentUnion() {
}

type PlanExternalPlanIDVersionNewParamsAddAdjustmentsAdjustmentNewUsageDiscountAdjustmentType string

const (
	PlanExternalPlanIDVersionNewParamsAddAdjustmentsAdjustmentNewUsageDiscountAdjustmentTypeUsageDiscount PlanExternalPlanIDVersionNewParamsAddAdjustmentsAdjustmentNewUsageDiscountAdjustmentType = "usage_discount"
)

func (r PlanExternalPlanIDVersionNewParamsAddAdjustmentsAdjustmentNewUsageDiscountAdjustmentType) IsKnown() bool {
	switch r {
	case PlanExternalPlanIDVersionNewParamsAddAdjustmentsAdjustmentNewUsageDiscountAdjustmentTypeUsageDiscount:
		return true
	}
	return false
}

type PlanExternalPlanIDVersionNewParamsAddAdjustmentsAdjustmentNewAmountDiscount struct {
	AdjustmentType param.Field[PlanExternalPlanIDVersionNewParamsAddAdjustmentsAdjustmentNewAmountDiscountAdjustmentType] `json:"adjustment_type,required"`
	AmountDiscount param.Field[string]                                                                                    `json:"amount_discount,required"`
	// The set of price IDs to which this adjustment applies.
	AppliesToPriceIDs param.Field[[]string] `json:"applies_to_price_ids"`
	// When false, this adjustment will be applied to a single price. Otherwise, it
	// will be applied at the invoice level, possibly to multiple prices.
	IsInvoiceLevel param.Field[bool] `json:"is_invoice_level"`
}

func (r PlanExternalPlanIDVersionNewParamsAddAdjustmentsAdjustmentNewAmountDiscount) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r PlanExternalPlanIDVersionNewParamsAddAdjustmentsAdjustmentNewAmountDiscount) implementsPlanExternalPlanIDVersionNewParamsAddAdjustmentsAdjustmentUnion() {
}

type PlanExternalPlanIDVersionNewParamsAddAdjustmentsAdjustmentNewAmountDiscountAdjustmentType string

const (
	PlanExternalPlanIDVersionNewParamsAddAdjustmentsAdjustmentNewAmountDiscountAdjustmentTypeAmountDiscount PlanExternalPlanIDVersionNewParamsAddAdjustmentsAdjustmentNewAmountDiscountAdjustmentType = "amount_discount"
)

func (r PlanExternalPlanIDVersionNewParamsAddAdjustmentsAdjustmentNewAmountDiscountAdjustmentType) IsKnown() bool {
	switch r {
	case PlanExternalPlanIDVersionNewParamsAddAdjustmentsAdjustmentNewAmountDiscountAdjustmentTypeAmountDiscount:
		return true
	}
	return false
}

type PlanExternalPlanIDVersionNewParamsAddAdjustmentsAdjustmentNewMinimum struct {
	AdjustmentType param.Field[PlanExternalPlanIDVersionNewParamsAddAdjustmentsAdjustmentNewMinimumAdjustmentType] `json:"adjustment_type,required"`
	// The item ID that revenue from this minimum will be attributed to.
	ItemID        param.Field[string] `json:"item_id,required"`
	MinimumAmount param.Field[string] `json:"minimum_amount,required"`
	// The set of price IDs to which this adjustment applies.
	AppliesToPriceIDs param.Field[[]string] `json:"applies_to_price_ids"`
	// When false, this adjustment will be applied to a single price. Otherwise, it
	// will be applied at the invoice level, possibly to multiple prices.
	IsInvoiceLevel param.Field[bool] `json:"is_invoice_level"`
}

func (r PlanExternalPlanIDVersionNewParamsAddAdjustmentsAdjustmentNewMinimum) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r PlanExternalPlanIDVersionNewParamsAddAdjustmentsAdjustmentNewMinimum) implementsPlanExternalPlanIDVersionNewParamsAddAdjustmentsAdjustmentUnion() {
}

type PlanExternalPlanIDVersionNewParamsAddAdjustmentsAdjustmentNewMinimumAdjustmentType string

const (
	PlanExternalPlanIDVersionNewParamsAddAdjustmentsAdjustmentNewMinimumAdjustmentTypeMinimum PlanExternalPlanIDVersionNewParamsAddAdjustmentsAdjustmentNewMinimumAdjustmentType = "minimum"
)

func (r PlanExternalPlanIDVersionNewParamsAddAdjustmentsAdjustmentNewMinimumAdjustmentType) IsKnown() bool {
	switch r {
	case PlanExternalPlanIDVersionNewParamsAddAdjustmentsAdjustmentNewMinimumAdjustmentTypeMinimum:
		return true
	}
	return false
}

type PlanExternalPlanIDVersionNewParamsAddAdjustmentsAdjustmentNewMaximum struct {
	AdjustmentType param.Field[PlanExternalPlanIDVersionNewParamsAddAdjustmentsAdjustmentNewMaximumAdjustmentType] `json:"adjustment_type,required"`
	MaximumAmount  param.Field[string]                                                                             `json:"maximum_amount,required"`
	// The set of price IDs to which this adjustment applies.
	AppliesToPriceIDs param.Field[[]string] `json:"applies_to_price_ids"`
	// When false, this adjustment will be applied to a single price. Otherwise, it
	// will be applied at the invoice level, possibly to multiple prices.
	IsInvoiceLevel param.Field[bool] `json:"is_invoice_level"`
}

func (r PlanExternalPlanIDVersionNewParamsAddAdjustmentsAdjustmentNewMaximum) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r PlanExternalPlanIDVersionNewParamsAddAdjustmentsAdjustmentNewMaximum) implementsPlanExternalPlanIDVersionNewParamsAddAdjustmentsAdjustmentUnion() {
}

type PlanExternalPlanIDVersionNewParamsAddAdjustmentsAdjustmentNewMaximumAdjustmentType string

const (
	PlanExternalPlanIDVersionNewParamsAddAdjustmentsAdjustmentNewMaximumAdjustmentTypeMaximum PlanExternalPlanIDVersionNewParamsAddAdjustmentsAdjustmentNewMaximumAdjustmentType = "maximum"
)

func (r PlanExternalPlanIDVersionNewParamsAddAdjustmentsAdjustmentNewMaximumAdjustmentType) IsKnown() bool {
	switch r {
	case PlanExternalPlanIDVersionNewParamsAddAdjustmentsAdjustmentNewMaximumAdjustmentTypeMaximum:
		return true
	}
	return false
}

type PlanExternalPlanIDVersionNewParamsAddAdjustmentsAdjustmentAdjustmentType string

const (
	PlanExternalPlanIDVersionNewParamsAddAdjustmentsAdjustmentAdjustmentTypePercentageDiscount PlanExternalPlanIDVersionNewParamsAddAdjustmentsAdjustmentAdjustmentType = "percentage_discount"
	PlanExternalPlanIDVersionNewParamsAddAdjustmentsAdjustmentAdjustmentTypeUsageDiscount      PlanExternalPlanIDVersionNewParamsAddAdjustmentsAdjustmentAdjustmentType = "usage_discount"
	PlanExternalPlanIDVersionNewParamsAddAdjustmentsAdjustmentAdjustmentTypeAmountDiscount     PlanExternalPlanIDVersionNewParamsAddAdjustmentsAdjustmentAdjustmentType = "amount_discount"
	PlanExternalPlanIDVersionNewParamsAddAdjustmentsAdjustmentAdjustmentTypeMinimum            PlanExternalPlanIDVersionNewParamsAddAdjustmentsAdjustmentAdjustmentType = "minimum"
	PlanExternalPlanIDVersionNewParamsAddAdjustmentsAdjustmentAdjustmentTypeMaximum            PlanExternalPlanIDVersionNewParamsAddAdjustmentsAdjustmentAdjustmentType = "maximum"
)

func (r PlanExternalPlanIDVersionNewParamsAddAdjustmentsAdjustmentAdjustmentType) IsKnown() bool {
	switch r {
	case PlanExternalPlanIDVersionNewParamsAddAdjustmentsAdjustmentAdjustmentTypePercentageDiscount, PlanExternalPlanIDVersionNewParamsAddAdjustmentsAdjustmentAdjustmentTypeUsageDiscount, PlanExternalPlanIDVersionNewParamsAddAdjustmentsAdjustmentAdjustmentTypeAmountDiscount, PlanExternalPlanIDVersionNewParamsAddAdjustmentsAdjustmentAdjustmentTypeMinimum, PlanExternalPlanIDVersionNewParamsAddAdjustmentsAdjustmentAdjustmentTypeMaximum:
		return true
	}
	return false
}

type PlanExternalPlanIDVersionNewParamsAddPrice struct {
	// The allocation price to add to the plan.
	AllocationPrice param.Field[PlanExternalPlanIDVersionNewParamsAddPricesAllocationPrice] `json:"allocation_price"`
	// The phase to add this price to.
	PlanPhaseOrder param.Field[int64] `json:"plan_phase_order"`
	// The price to add to the plan
	Price param.Field[PlanExternalPlanIDVersionNewParamsAddPricesPriceUnion] `json:"price"`
}

func (r PlanExternalPlanIDVersionNewParamsAddPrice) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// The allocation price to add to the plan.
type PlanExternalPlanIDVersionNewParamsAddPricesAllocationPrice struct {
	// An amount of the currency to allocate to the customer at the specified cadence.
	Amount param.Field[string] `json:"amount,required"`
	// The cadence at which to allocate the amount to the customer.
	Cadence param.Field[PlanExternalPlanIDVersionNewParamsAddPricesAllocationPriceCadence] `json:"cadence,required"`
	// An ISO 4217 currency string or a custom pricing unit identifier in which to bill
	// this price.
	Currency param.Field[string] `json:"currency,required"`
	// Whether the allocated amount should expire at the end of the cadence or roll
	// over to the next period.
	ExpiresAtEndOfCadence param.Field[bool] `json:"expires_at_end_of_cadence,required"`
}

func (r PlanExternalPlanIDVersionNewParamsAddPricesAllocationPrice) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// The cadence at which to allocate the amount to the customer.
type PlanExternalPlanIDVersionNewParamsAddPricesAllocationPriceCadence string

const (
	PlanExternalPlanIDVersionNewParamsAddPricesAllocationPriceCadenceOneTime    PlanExternalPlanIDVersionNewParamsAddPricesAllocationPriceCadence = "one_time"
	PlanExternalPlanIDVersionNewParamsAddPricesAllocationPriceCadenceMonthly    PlanExternalPlanIDVersionNewParamsAddPricesAllocationPriceCadence = "monthly"
	PlanExternalPlanIDVersionNewParamsAddPricesAllocationPriceCadenceQuarterly  PlanExternalPlanIDVersionNewParamsAddPricesAllocationPriceCadence = "quarterly"
	PlanExternalPlanIDVersionNewParamsAddPricesAllocationPriceCadenceSemiAnnual PlanExternalPlanIDVersionNewParamsAddPricesAllocationPriceCadence = "semi_annual"
	PlanExternalPlanIDVersionNewParamsAddPricesAllocationPriceCadenceAnnual     PlanExternalPlanIDVersionNewParamsAddPricesAllocationPriceCadence = "annual"
	PlanExternalPlanIDVersionNewParamsAddPricesAllocationPriceCadenceCustom     PlanExternalPlanIDVersionNewParamsAddPricesAllocationPriceCadence = "custom"
)

func (r PlanExternalPlanIDVersionNewParamsAddPricesAllocationPriceCadence) IsKnown() bool {
	switch r {
	case PlanExternalPlanIDVersionNewParamsAddPricesAllocationPriceCadenceOneTime, PlanExternalPlanIDVersionNewParamsAddPricesAllocationPriceCadenceMonthly, PlanExternalPlanIDVersionNewParamsAddPricesAllocationPriceCadenceQuarterly, PlanExternalPlanIDVersionNewParamsAddPricesAllocationPriceCadenceSemiAnnual, PlanExternalPlanIDVersionNewParamsAddPricesAllocationPriceCadenceAnnual, PlanExternalPlanIDVersionNewParamsAddPricesAllocationPriceCadenceCustom:
		return true
	}
	return false
}

// The price to add to the plan
type PlanExternalPlanIDVersionNewParamsAddPricesPrice struct {
	// The cadence to bill for this price on.
	Cadence param.Field[PlanExternalPlanIDVersionNewParamsAddPricesPriceCadence] `json:"cadence,required"`
	// The id of the item the price will be associated with.
	ItemID    param.Field[string]                                                    `json:"item_id,required"`
	ModelType param.Field[PlanExternalPlanIDVersionNewParamsAddPricesPriceModelType] `json:"model_type,required"`
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

func (r PlanExternalPlanIDVersionNewParamsAddPricesPrice) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r PlanExternalPlanIDVersionNewParamsAddPricesPrice) implementsPlanExternalPlanIDVersionNewParamsAddPricesPriceUnion() {
}

// The price to add to the plan
//
// Satisfied by [PlanExternalPlanIDVersionNewParamsAddPricesPriceNewPlanUnitPrice],
// [PlanExternalPlanIDVersionNewParamsAddPricesPriceNewPlanPackagePrice],
// [PlanExternalPlanIDVersionNewParamsAddPricesPriceNewPlanMatrixPrice],
// [PlanExternalPlanIDVersionNewParamsAddPricesPriceNewPlanTieredPrice],
// [PlanExternalPlanIDVersionNewParamsAddPricesPriceNewPlanTieredBpsPrice],
// [PlanExternalPlanIDVersionNewParamsAddPricesPriceNewPlanBpsPrice],
// [PlanExternalPlanIDVersionNewParamsAddPricesPriceNewPlanBulkBpsPrice],
// [PlanExternalPlanIDVersionNewParamsAddPricesPriceNewPlanBulkPrice],
// [PlanExternalPlanIDVersionNewParamsAddPricesPriceNewPlanThresholdTotalAmountPrice],
// [PlanExternalPlanIDVersionNewParamsAddPricesPriceNewPlanTieredPackagePrice],
// [PlanExternalPlanIDVersionNewParamsAddPricesPriceNewPlanTieredWithMinimumPrice],
// [PlanExternalPlanIDVersionNewParamsAddPricesPriceNewPlanUnitWithPercentPrice],
// [PlanExternalPlanIDVersionNewParamsAddPricesPriceNewPlanPackageWithAllocationPrice],
// [PlanExternalPlanIDVersionNewParamsAddPricesPriceNewPlanTierWithProrationPrice],
// [PlanExternalPlanIDVersionNewParamsAddPricesPriceNewPlanUnitWithProrationPrice],
// [PlanExternalPlanIDVersionNewParamsAddPricesPriceNewPlanGroupedAllocationPrice],
// [PlanExternalPlanIDVersionNewParamsAddPricesPriceNewPlanGroupedWithProratedMinimumPrice],
// [PlanExternalPlanIDVersionNewParamsAddPricesPriceNewPlanGroupedWithMeteredMinimumPrice],
// [PlanExternalPlanIDVersionNewParamsAddPricesPriceNewPlanMatrixWithDisplayNamePrice],
// [PlanExternalPlanIDVersionNewParamsAddPricesPriceNewPlanBulkWithProrationPrice],
// [PlanExternalPlanIDVersionNewParamsAddPricesPriceNewPlanGroupedTieredPackagePrice],
// [PlanExternalPlanIDVersionNewParamsAddPricesPriceNewPlanMaxGroupTieredPackagePrice],
// [PlanExternalPlanIDVersionNewParamsAddPricesPriceNewPlanScalableMatrixWithUnitPricingPrice],
// [PlanExternalPlanIDVersionNewParamsAddPricesPriceNewPlanScalableMatrixWithTieredPricingPrice],
// [PlanExternalPlanIDVersionNewParamsAddPricesPriceNewPlanCumulativeGroupedBulkPrice],
// [PlanExternalPlanIDVersionNewParamsAddPricesPriceNewPlanTieredPackageWithMinimumPrice],
// [PlanExternalPlanIDVersionNewParamsAddPricesPriceNewPlanMatrixWithAllocationPrice],
// [PlanExternalPlanIDVersionNewParamsAddPricesPriceNewPlanGroupedTieredPrice],
// [PlanExternalPlanIDVersionNewParamsAddPricesPrice].
type PlanExternalPlanIDVersionNewParamsAddPricesPriceUnion interface {
	implementsPlanExternalPlanIDVersionNewParamsAddPricesPriceUnion()
}

type PlanExternalPlanIDVersionNewParamsAddPricesPriceNewPlanUnitPrice struct {
	// The cadence to bill for this price on.
	Cadence param.Field[PlanExternalPlanIDVersionNewParamsAddPricesPriceNewPlanUnitPriceCadence] `json:"cadence,required"`
	// The id of the item the price will be associated with.
	ItemID    param.Field[string]                                                                    `json:"item_id,required"`
	ModelType param.Field[PlanExternalPlanIDVersionNewParamsAddPricesPriceNewPlanUnitPriceModelType] `json:"model_type,required"`
	// The name of the price.
	Name       param.Field[string]                                                                     `json:"name,required"`
	UnitConfig param.Field[PlanExternalPlanIDVersionNewParamsAddPricesPriceNewPlanUnitPriceUnitConfig] `json:"unit_config,required"`
	// The id of the billable metric for the price. Only needed if the price is
	// usage-based.
	BillableMetricID param.Field[string] `json:"billable_metric_id"`
	// If the Price represents a fixed cost, the price will be billed in-advance if
	// this is true, and in-arrears if this is false.
	BilledInAdvance param.Field[bool] `json:"billed_in_advance"`
	// For custom cadence: specifies the duration of the billing period in days or
	// months.
	BillingCycleConfiguration param.Field[PlanExternalPlanIDVersionNewParamsAddPricesPriceNewPlanUnitPriceBillingCycleConfiguration] `json:"billing_cycle_configuration"`
	// The per unit conversion rate of the price currency to the invoicing currency.
	ConversionRate param.Field[float64] `json:"conversion_rate"`
	// An ISO 4217 currency string, or custom pricing unit identifier, in which this
	// price is billed.
	Currency param.Field[string] `json:"currency"`
	// For dimensional price: specifies a price group and dimension values
	DimensionalPriceConfiguration param.Field[PlanExternalPlanIDVersionNewParamsAddPricesPriceNewPlanUnitPriceDimensionalPriceConfiguration] `json:"dimensional_price_configuration"`
	// An alias for the price.
	ExternalPriceID param.Field[string] `json:"external_price_id"`
	// If the Price represents a fixed cost, this represents the quantity of units
	// applied.
	FixedPriceQuantity param.Field[float64] `json:"fixed_price_quantity"`
	// The property used to group this price on an invoice
	InvoiceGroupingKey param.Field[string] `json:"invoice_grouping_key"`
	// Within each billing cycle, specifies the cadence at which invoices are produced.
	// If unspecified, a single invoice is produced per billing cycle.
	InvoicingCycleConfiguration param.Field[PlanExternalPlanIDVersionNewParamsAddPricesPriceNewPlanUnitPriceInvoicingCycleConfiguration] `json:"invoicing_cycle_configuration"`
	// User-specified key/value pairs for the resource. Individual keys can be removed
	// by setting the value to `null`, and the entire metadata mapping can be cleared
	// by setting `metadata` to `null`.
	Metadata param.Field[map[string]string] `json:"metadata"`
}

func (r PlanExternalPlanIDVersionNewParamsAddPricesPriceNewPlanUnitPrice) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r PlanExternalPlanIDVersionNewParamsAddPricesPriceNewPlanUnitPrice) implementsPlanExternalPlanIDVersionNewParamsAddPricesPriceUnion() {
}

// The cadence to bill for this price on.
type PlanExternalPlanIDVersionNewParamsAddPricesPriceNewPlanUnitPriceCadence string

const (
	PlanExternalPlanIDVersionNewParamsAddPricesPriceNewPlanUnitPriceCadenceAnnual     PlanExternalPlanIDVersionNewParamsAddPricesPriceNewPlanUnitPriceCadence = "annual"
	PlanExternalPlanIDVersionNewParamsAddPricesPriceNewPlanUnitPriceCadenceSemiAnnual PlanExternalPlanIDVersionNewParamsAddPricesPriceNewPlanUnitPriceCadence = "semi_annual"
	PlanExternalPlanIDVersionNewParamsAddPricesPriceNewPlanUnitPriceCadenceMonthly    PlanExternalPlanIDVersionNewParamsAddPricesPriceNewPlanUnitPriceCadence = "monthly"
	PlanExternalPlanIDVersionNewParamsAddPricesPriceNewPlanUnitPriceCadenceQuarterly  PlanExternalPlanIDVersionNewParamsAddPricesPriceNewPlanUnitPriceCadence = "quarterly"
	PlanExternalPlanIDVersionNewParamsAddPricesPriceNewPlanUnitPriceCadenceOneTime    PlanExternalPlanIDVersionNewParamsAddPricesPriceNewPlanUnitPriceCadence = "one_time"
	PlanExternalPlanIDVersionNewParamsAddPricesPriceNewPlanUnitPriceCadenceCustom     PlanExternalPlanIDVersionNewParamsAddPricesPriceNewPlanUnitPriceCadence = "custom"
)

func (r PlanExternalPlanIDVersionNewParamsAddPricesPriceNewPlanUnitPriceCadence) IsKnown() bool {
	switch r {
	case PlanExternalPlanIDVersionNewParamsAddPricesPriceNewPlanUnitPriceCadenceAnnual, PlanExternalPlanIDVersionNewParamsAddPricesPriceNewPlanUnitPriceCadenceSemiAnnual, PlanExternalPlanIDVersionNewParamsAddPricesPriceNewPlanUnitPriceCadenceMonthly, PlanExternalPlanIDVersionNewParamsAddPricesPriceNewPlanUnitPriceCadenceQuarterly, PlanExternalPlanIDVersionNewParamsAddPricesPriceNewPlanUnitPriceCadenceOneTime, PlanExternalPlanIDVersionNewParamsAddPricesPriceNewPlanUnitPriceCadenceCustom:
		return true
	}
	return false
}

type PlanExternalPlanIDVersionNewParamsAddPricesPriceNewPlanUnitPriceModelType string

const (
	PlanExternalPlanIDVersionNewParamsAddPricesPriceNewPlanUnitPriceModelTypeUnit PlanExternalPlanIDVersionNewParamsAddPricesPriceNewPlanUnitPriceModelType = "unit"
)

func (r PlanExternalPlanIDVersionNewParamsAddPricesPriceNewPlanUnitPriceModelType) IsKnown() bool {
	switch r {
	case PlanExternalPlanIDVersionNewParamsAddPricesPriceNewPlanUnitPriceModelTypeUnit:
		return true
	}
	return false
}

type PlanExternalPlanIDVersionNewParamsAddPricesPriceNewPlanUnitPriceUnitConfig struct {
	// Rate per unit of usage
	UnitAmount param.Field[string] `json:"unit_amount,required"`
}

func (r PlanExternalPlanIDVersionNewParamsAddPricesPriceNewPlanUnitPriceUnitConfig) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// For custom cadence: specifies the duration of the billing period in days or
// months.
type PlanExternalPlanIDVersionNewParamsAddPricesPriceNewPlanUnitPriceBillingCycleConfiguration struct {
	// The duration of the billing period.
	Duration param.Field[int64] `json:"duration,required"`
	// The unit of billing period duration.
	DurationUnit param.Field[PlanExternalPlanIDVersionNewParamsAddPricesPriceNewPlanUnitPriceBillingCycleConfigurationDurationUnit] `json:"duration_unit,required"`
}

func (r PlanExternalPlanIDVersionNewParamsAddPricesPriceNewPlanUnitPriceBillingCycleConfiguration) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// The unit of billing period duration.
type PlanExternalPlanIDVersionNewParamsAddPricesPriceNewPlanUnitPriceBillingCycleConfigurationDurationUnit string

const (
	PlanExternalPlanIDVersionNewParamsAddPricesPriceNewPlanUnitPriceBillingCycleConfigurationDurationUnitDay   PlanExternalPlanIDVersionNewParamsAddPricesPriceNewPlanUnitPriceBillingCycleConfigurationDurationUnit = "day"
	PlanExternalPlanIDVersionNewParamsAddPricesPriceNewPlanUnitPriceBillingCycleConfigurationDurationUnitMonth PlanExternalPlanIDVersionNewParamsAddPricesPriceNewPlanUnitPriceBillingCycleConfigurationDurationUnit = "month"
)

func (r PlanExternalPlanIDVersionNewParamsAddPricesPriceNewPlanUnitPriceBillingCycleConfigurationDurationUnit) IsKnown() bool {
	switch r {
	case PlanExternalPlanIDVersionNewParamsAddPricesPriceNewPlanUnitPriceBillingCycleConfigurationDurationUnitDay, PlanExternalPlanIDVersionNewParamsAddPricesPriceNewPlanUnitPriceBillingCycleConfigurationDurationUnitMonth:
		return true
	}
	return false
}

// For dimensional price: specifies a price group and dimension values
type PlanExternalPlanIDVersionNewParamsAddPricesPriceNewPlanUnitPriceDimensionalPriceConfiguration struct {
	// The list of dimension values matching (in order) the dimensions of the price
	// group
	DimensionValues param.Field[[]string] `json:"dimension_values,required"`
	// The id of the dimensional price group to include this price in
	DimensionalPriceGroupID param.Field[string] `json:"dimensional_price_group_id"`
	// The external id of the dimensional price group to include this price in
	ExternalDimensionalPriceGroupID param.Field[string] `json:"external_dimensional_price_group_id"`
}

func (r PlanExternalPlanIDVersionNewParamsAddPricesPriceNewPlanUnitPriceDimensionalPriceConfiguration) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Within each billing cycle, specifies the cadence at which invoices are produced.
// If unspecified, a single invoice is produced per billing cycle.
type PlanExternalPlanIDVersionNewParamsAddPricesPriceNewPlanUnitPriceInvoicingCycleConfiguration struct {
	// The duration of the billing period.
	Duration param.Field[int64] `json:"duration,required"`
	// The unit of billing period duration.
	DurationUnit param.Field[PlanExternalPlanIDVersionNewParamsAddPricesPriceNewPlanUnitPriceInvoicingCycleConfigurationDurationUnit] `json:"duration_unit,required"`
}

func (r PlanExternalPlanIDVersionNewParamsAddPricesPriceNewPlanUnitPriceInvoicingCycleConfiguration) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// The unit of billing period duration.
type PlanExternalPlanIDVersionNewParamsAddPricesPriceNewPlanUnitPriceInvoicingCycleConfigurationDurationUnit string

const (
	PlanExternalPlanIDVersionNewParamsAddPricesPriceNewPlanUnitPriceInvoicingCycleConfigurationDurationUnitDay   PlanExternalPlanIDVersionNewParamsAddPricesPriceNewPlanUnitPriceInvoicingCycleConfigurationDurationUnit = "day"
	PlanExternalPlanIDVersionNewParamsAddPricesPriceNewPlanUnitPriceInvoicingCycleConfigurationDurationUnitMonth PlanExternalPlanIDVersionNewParamsAddPricesPriceNewPlanUnitPriceInvoicingCycleConfigurationDurationUnit = "month"
)

func (r PlanExternalPlanIDVersionNewParamsAddPricesPriceNewPlanUnitPriceInvoicingCycleConfigurationDurationUnit) IsKnown() bool {
	switch r {
	case PlanExternalPlanIDVersionNewParamsAddPricesPriceNewPlanUnitPriceInvoicingCycleConfigurationDurationUnitDay, PlanExternalPlanIDVersionNewParamsAddPricesPriceNewPlanUnitPriceInvoicingCycleConfigurationDurationUnitMonth:
		return true
	}
	return false
}

type PlanExternalPlanIDVersionNewParamsAddPricesPriceNewPlanPackagePrice struct {
	// The cadence to bill for this price on.
	Cadence param.Field[PlanExternalPlanIDVersionNewParamsAddPricesPriceNewPlanPackagePriceCadence] `json:"cadence,required"`
	// The id of the item the price will be associated with.
	ItemID    param.Field[string]                                                                       `json:"item_id,required"`
	ModelType param.Field[PlanExternalPlanIDVersionNewParamsAddPricesPriceNewPlanPackagePriceModelType] `json:"model_type,required"`
	// The name of the price.
	Name          param.Field[string]                                                                           `json:"name,required"`
	PackageConfig param.Field[PlanExternalPlanIDVersionNewParamsAddPricesPriceNewPlanPackagePricePackageConfig] `json:"package_config,required"`
	// The id of the billable metric for the price. Only needed if the price is
	// usage-based.
	BillableMetricID param.Field[string] `json:"billable_metric_id"`
	// If the Price represents a fixed cost, the price will be billed in-advance if
	// this is true, and in-arrears if this is false.
	BilledInAdvance param.Field[bool] `json:"billed_in_advance"`
	// For custom cadence: specifies the duration of the billing period in days or
	// months.
	BillingCycleConfiguration param.Field[PlanExternalPlanIDVersionNewParamsAddPricesPriceNewPlanPackagePriceBillingCycleConfiguration] `json:"billing_cycle_configuration"`
	// The per unit conversion rate of the price currency to the invoicing currency.
	ConversionRate param.Field[float64] `json:"conversion_rate"`
	// An ISO 4217 currency string, or custom pricing unit identifier, in which this
	// price is billed.
	Currency param.Field[string] `json:"currency"`
	// For dimensional price: specifies a price group and dimension values
	DimensionalPriceConfiguration param.Field[PlanExternalPlanIDVersionNewParamsAddPricesPriceNewPlanPackagePriceDimensionalPriceConfiguration] `json:"dimensional_price_configuration"`
	// An alias for the price.
	ExternalPriceID param.Field[string] `json:"external_price_id"`
	// If the Price represents a fixed cost, this represents the quantity of units
	// applied.
	FixedPriceQuantity param.Field[float64] `json:"fixed_price_quantity"`
	// The property used to group this price on an invoice
	InvoiceGroupingKey param.Field[string] `json:"invoice_grouping_key"`
	// Within each billing cycle, specifies the cadence at which invoices are produced.
	// If unspecified, a single invoice is produced per billing cycle.
	InvoicingCycleConfiguration param.Field[PlanExternalPlanIDVersionNewParamsAddPricesPriceNewPlanPackagePriceInvoicingCycleConfiguration] `json:"invoicing_cycle_configuration"`
	// User-specified key/value pairs for the resource. Individual keys can be removed
	// by setting the value to `null`, and the entire metadata mapping can be cleared
	// by setting `metadata` to `null`.
	Metadata param.Field[map[string]string] `json:"metadata"`
}

func (r PlanExternalPlanIDVersionNewParamsAddPricesPriceNewPlanPackagePrice) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r PlanExternalPlanIDVersionNewParamsAddPricesPriceNewPlanPackagePrice) implementsPlanExternalPlanIDVersionNewParamsAddPricesPriceUnion() {
}

// The cadence to bill for this price on.
type PlanExternalPlanIDVersionNewParamsAddPricesPriceNewPlanPackagePriceCadence string

const (
	PlanExternalPlanIDVersionNewParamsAddPricesPriceNewPlanPackagePriceCadenceAnnual     PlanExternalPlanIDVersionNewParamsAddPricesPriceNewPlanPackagePriceCadence = "annual"
	PlanExternalPlanIDVersionNewParamsAddPricesPriceNewPlanPackagePriceCadenceSemiAnnual PlanExternalPlanIDVersionNewParamsAddPricesPriceNewPlanPackagePriceCadence = "semi_annual"
	PlanExternalPlanIDVersionNewParamsAddPricesPriceNewPlanPackagePriceCadenceMonthly    PlanExternalPlanIDVersionNewParamsAddPricesPriceNewPlanPackagePriceCadence = "monthly"
	PlanExternalPlanIDVersionNewParamsAddPricesPriceNewPlanPackagePriceCadenceQuarterly  PlanExternalPlanIDVersionNewParamsAddPricesPriceNewPlanPackagePriceCadence = "quarterly"
	PlanExternalPlanIDVersionNewParamsAddPricesPriceNewPlanPackagePriceCadenceOneTime    PlanExternalPlanIDVersionNewParamsAddPricesPriceNewPlanPackagePriceCadence = "one_time"
	PlanExternalPlanIDVersionNewParamsAddPricesPriceNewPlanPackagePriceCadenceCustom     PlanExternalPlanIDVersionNewParamsAddPricesPriceNewPlanPackagePriceCadence = "custom"
)

func (r PlanExternalPlanIDVersionNewParamsAddPricesPriceNewPlanPackagePriceCadence) IsKnown() bool {
	switch r {
	case PlanExternalPlanIDVersionNewParamsAddPricesPriceNewPlanPackagePriceCadenceAnnual, PlanExternalPlanIDVersionNewParamsAddPricesPriceNewPlanPackagePriceCadenceSemiAnnual, PlanExternalPlanIDVersionNewParamsAddPricesPriceNewPlanPackagePriceCadenceMonthly, PlanExternalPlanIDVersionNewParamsAddPricesPriceNewPlanPackagePriceCadenceQuarterly, PlanExternalPlanIDVersionNewParamsAddPricesPriceNewPlanPackagePriceCadenceOneTime, PlanExternalPlanIDVersionNewParamsAddPricesPriceNewPlanPackagePriceCadenceCustom:
		return true
	}
	return false
}

type PlanExternalPlanIDVersionNewParamsAddPricesPriceNewPlanPackagePriceModelType string

const (
	PlanExternalPlanIDVersionNewParamsAddPricesPriceNewPlanPackagePriceModelTypePackage PlanExternalPlanIDVersionNewParamsAddPricesPriceNewPlanPackagePriceModelType = "package"
)

func (r PlanExternalPlanIDVersionNewParamsAddPricesPriceNewPlanPackagePriceModelType) IsKnown() bool {
	switch r {
	case PlanExternalPlanIDVersionNewParamsAddPricesPriceNewPlanPackagePriceModelTypePackage:
		return true
	}
	return false
}

type PlanExternalPlanIDVersionNewParamsAddPricesPriceNewPlanPackagePricePackageConfig struct {
	// A currency amount to rate usage by
	PackageAmount param.Field[string] `json:"package_amount,required"`
	// An integer amount to represent package size. For example, 1000 here would divide
	// usage by 1000 before multiplying by package_amount in rating
	PackageSize param.Field[int64] `json:"package_size,required"`
}

func (r PlanExternalPlanIDVersionNewParamsAddPricesPriceNewPlanPackagePricePackageConfig) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// For custom cadence: specifies the duration of the billing period in days or
// months.
type PlanExternalPlanIDVersionNewParamsAddPricesPriceNewPlanPackagePriceBillingCycleConfiguration struct {
	// The duration of the billing period.
	Duration param.Field[int64] `json:"duration,required"`
	// The unit of billing period duration.
	DurationUnit param.Field[PlanExternalPlanIDVersionNewParamsAddPricesPriceNewPlanPackagePriceBillingCycleConfigurationDurationUnit] `json:"duration_unit,required"`
}

func (r PlanExternalPlanIDVersionNewParamsAddPricesPriceNewPlanPackagePriceBillingCycleConfiguration) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// The unit of billing period duration.
type PlanExternalPlanIDVersionNewParamsAddPricesPriceNewPlanPackagePriceBillingCycleConfigurationDurationUnit string

const (
	PlanExternalPlanIDVersionNewParamsAddPricesPriceNewPlanPackagePriceBillingCycleConfigurationDurationUnitDay   PlanExternalPlanIDVersionNewParamsAddPricesPriceNewPlanPackagePriceBillingCycleConfigurationDurationUnit = "day"
	PlanExternalPlanIDVersionNewParamsAddPricesPriceNewPlanPackagePriceBillingCycleConfigurationDurationUnitMonth PlanExternalPlanIDVersionNewParamsAddPricesPriceNewPlanPackagePriceBillingCycleConfigurationDurationUnit = "month"
)

func (r PlanExternalPlanIDVersionNewParamsAddPricesPriceNewPlanPackagePriceBillingCycleConfigurationDurationUnit) IsKnown() bool {
	switch r {
	case PlanExternalPlanIDVersionNewParamsAddPricesPriceNewPlanPackagePriceBillingCycleConfigurationDurationUnitDay, PlanExternalPlanIDVersionNewParamsAddPricesPriceNewPlanPackagePriceBillingCycleConfigurationDurationUnitMonth:
		return true
	}
	return false
}

// For dimensional price: specifies a price group and dimension values
type PlanExternalPlanIDVersionNewParamsAddPricesPriceNewPlanPackagePriceDimensionalPriceConfiguration struct {
	// The list of dimension values matching (in order) the dimensions of the price
	// group
	DimensionValues param.Field[[]string] `json:"dimension_values,required"`
	// The id of the dimensional price group to include this price in
	DimensionalPriceGroupID param.Field[string] `json:"dimensional_price_group_id"`
	// The external id of the dimensional price group to include this price in
	ExternalDimensionalPriceGroupID param.Field[string] `json:"external_dimensional_price_group_id"`
}

func (r PlanExternalPlanIDVersionNewParamsAddPricesPriceNewPlanPackagePriceDimensionalPriceConfiguration) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Within each billing cycle, specifies the cadence at which invoices are produced.
// If unspecified, a single invoice is produced per billing cycle.
type PlanExternalPlanIDVersionNewParamsAddPricesPriceNewPlanPackagePriceInvoicingCycleConfiguration struct {
	// The duration of the billing period.
	Duration param.Field[int64] `json:"duration,required"`
	// The unit of billing period duration.
	DurationUnit param.Field[PlanExternalPlanIDVersionNewParamsAddPricesPriceNewPlanPackagePriceInvoicingCycleConfigurationDurationUnit] `json:"duration_unit,required"`
}

func (r PlanExternalPlanIDVersionNewParamsAddPricesPriceNewPlanPackagePriceInvoicingCycleConfiguration) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// The unit of billing period duration.
type PlanExternalPlanIDVersionNewParamsAddPricesPriceNewPlanPackagePriceInvoicingCycleConfigurationDurationUnit string

const (
	PlanExternalPlanIDVersionNewParamsAddPricesPriceNewPlanPackagePriceInvoicingCycleConfigurationDurationUnitDay   PlanExternalPlanIDVersionNewParamsAddPricesPriceNewPlanPackagePriceInvoicingCycleConfigurationDurationUnit = "day"
	PlanExternalPlanIDVersionNewParamsAddPricesPriceNewPlanPackagePriceInvoicingCycleConfigurationDurationUnitMonth PlanExternalPlanIDVersionNewParamsAddPricesPriceNewPlanPackagePriceInvoicingCycleConfigurationDurationUnit = "month"
)

func (r PlanExternalPlanIDVersionNewParamsAddPricesPriceNewPlanPackagePriceInvoicingCycleConfigurationDurationUnit) IsKnown() bool {
	switch r {
	case PlanExternalPlanIDVersionNewParamsAddPricesPriceNewPlanPackagePriceInvoicingCycleConfigurationDurationUnitDay, PlanExternalPlanIDVersionNewParamsAddPricesPriceNewPlanPackagePriceInvoicingCycleConfigurationDurationUnitMonth:
		return true
	}
	return false
}

type PlanExternalPlanIDVersionNewParamsAddPricesPriceNewPlanMatrixPrice struct {
	// The cadence to bill for this price on.
	Cadence param.Field[PlanExternalPlanIDVersionNewParamsAddPricesPriceNewPlanMatrixPriceCadence] `json:"cadence,required"`
	// The id of the item the price will be associated with.
	ItemID       param.Field[string]                                                                         `json:"item_id,required"`
	MatrixConfig param.Field[PlanExternalPlanIDVersionNewParamsAddPricesPriceNewPlanMatrixPriceMatrixConfig] `json:"matrix_config,required"`
	ModelType    param.Field[PlanExternalPlanIDVersionNewParamsAddPricesPriceNewPlanMatrixPriceModelType]    `json:"model_type,required"`
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
	BillingCycleConfiguration param.Field[PlanExternalPlanIDVersionNewParamsAddPricesPriceNewPlanMatrixPriceBillingCycleConfiguration] `json:"billing_cycle_configuration"`
	// The per unit conversion rate of the price currency to the invoicing currency.
	ConversionRate param.Field[float64] `json:"conversion_rate"`
	// An ISO 4217 currency string, or custom pricing unit identifier, in which this
	// price is billed.
	Currency param.Field[string] `json:"currency"`
	// For dimensional price: specifies a price group and dimension values
	DimensionalPriceConfiguration param.Field[PlanExternalPlanIDVersionNewParamsAddPricesPriceNewPlanMatrixPriceDimensionalPriceConfiguration] `json:"dimensional_price_configuration"`
	// An alias for the price.
	ExternalPriceID param.Field[string] `json:"external_price_id"`
	// If the Price represents a fixed cost, this represents the quantity of units
	// applied.
	FixedPriceQuantity param.Field[float64] `json:"fixed_price_quantity"`
	// The property used to group this price on an invoice
	InvoiceGroupingKey param.Field[string] `json:"invoice_grouping_key"`
	// Within each billing cycle, specifies the cadence at which invoices are produced.
	// If unspecified, a single invoice is produced per billing cycle.
	InvoicingCycleConfiguration param.Field[PlanExternalPlanIDVersionNewParamsAddPricesPriceNewPlanMatrixPriceInvoicingCycleConfiguration] `json:"invoicing_cycle_configuration"`
	// User-specified key/value pairs for the resource. Individual keys can be removed
	// by setting the value to `null`, and the entire metadata mapping can be cleared
	// by setting `metadata` to `null`.
	Metadata param.Field[map[string]string] `json:"metadata"`
}

func (r PlanExternalPlanIDVersionNewParamsAddPricesPriceNewPlanMatrixPrice) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r PlanExternalPlanIDVersionNewParamsAddPricesPriceNewPlanMatrixPrice) implementsPlanExternalPlanIDVersionNewParamsAddPricesPriceUnion() {
}

// The cadence to bill for this price on.
type PlanExternalPlanIDVersionNewParamsAddPricesPriceNewPlanMatrixPriceCadence string

const (
	PlanExternalPlanIDVersionNewParamsAddPricesPriceNewPlanMatrixPriceCadenceAnnual     PlanExternalPlanIDVersionNewParamsAddPricesPriceNewPlanMatrixPriceCadence = "annual"
	PlanExternalPlanIDVersionNewParamsAddPricesPriceNewPlanMatrixPriceCadenceSemiAnnual PlanExternalPlanIDVersionNewParamsAddPricesPriceNewPlanMatrixPriceCadence = "semi_annual"
	PlanExternalPlanIDVersionNewParamsAddPricesPriceNewPlanMatrixPriceCadenceMonthly    PlanExternalPlanIDVersionNewParamsAddPricesPriceNewPlanMatrixPriceCadence = "monthly"
	PlanExternalPlanIDVersionNewParamsAddPricesPriceNewPlanMatrixPriceCadenceQuarterly  PlanExternalPlanIDVersionNewParamsAddPricesPriceNewPlanMatrixPriceCadence = "quarterly"
	PlanExternalPlanIDVersionNewParamsAddPricesPriceNewPlanMatrixPriceCadenceOneTime    PlanExternalPlanIDVersionNewParamsAddPricesPriceNewPlanMatrixPriceCadence = "one_time"
	PlanExternalPlanIDVersionNewParamsAddPricesPriceNewPlanMatrixPriceCadenceCustom     PlanExternalPlanIDVersionNewParamsAddPricesPriceNewPlanMatrixPriceCadence = "custom"
)

func (r PlanExternalPlanIDVersionNewParamsAddPricesPriceNewPlanMatrixPriceCadence) IsKnown() bool {
	switch r {
	case PlanExternalPlanIDVersionNewParamsAddPricesPriceNewPlanMatrixPriceCadenceAnnual, PlanExternalPlanIDVersionNewParamsAddPricesPriceNewPlanMatrixPriceCadenceSemiAnnual, PlanExternalPlanIDVersionNewParamsAddPricesPriceNewPlanMatrixPriceCadenceMonthly, PlanExternalPlanIDVersionNewParamsAddPricesPriceNewPlanMatrixPriceCadenceQuarterly, PlanExternalPlanIDVersionNewParamsAddPricesPriceNewPlanMatrixPriceCadenceOneTime, PlanExternalPlanIDVersionNewParamsAddPricesPriceNewPlanMatrixPriceCadenceCustom:
		return true
	}
	return false
}

type PlanExternalPlanIDVersionNewParamsAddPricesPriceNewPlanMatrixPriceMatrixConfig struct {
	// Default per unit rate for any usage not bucketed into a specified matrix_value
	DefaultUnitAmount param.Field[string] `json:"default_unit_amount,required"`
	// One or two event property values to evaluate matrix groups by
	Dimensions param.Field[[]string] `json:"dimensions,required"`
	// Matrix values for specified matrix grouping keys
	MatrixValues param.Field[[]PlanExternalPlanIDVersionNewParamsAddPricesPriceNewPlanMatrixPriceMatrixConfigMatrixValue] `json:"matrix_values,required"`
}

func (r PlanExternalPlanIDVersionNewParamsAddPricesPriceNewPlanMatrixPriceMatrixConfig) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type PlanExternalPlanIDVersionNewParamsAddPricesPriceNewPlanMatrixPriceMatrixConfigMatrixValue struct {
	// One or two matrix keys to filter usage to this Matrix value by. For example,
	// ["region", "tier"] could be used to filter cloud usage by a cloud region and an
	// instance tier.
	DimensionValues param.Field[[]string] `json:"dimension_values,required"`
	// Unit price for the specified dimension_values
	UnitAmount param.Field[string] `json:"unit_amount,required"`
}

func (r PlanExternalPlanIDVersionNewParamsAddPricesPriceNewPlanMatrixPriceMatrixConfigMatrixValue) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type PlanExternalPlanIDVersionNewParamsAddPricesPriceNewPlanMatrixPriceModelType string

const (
	PlanExternalPlanIDVersionNewParamsAddPricesPriceNewPlanMatrixPriceModelTypeMatrix PlanExternalPlanIDVersionNewParamsAddPricesPriceNewPlanMatrixPriceModelType = "matrix"
)

func (r PlanExternalPlanIDVersionNewParamsAddPricesPriceNewPlanMatrixPriceModelType) IsKnown() bool {
	switch r {
	case PlanExternalPlanIDVersionNewParamsAddPricesPriceNewPlanMatrixPriceModelTypeMatrix:
		return true
	}
	return false
}

// For custom cadence: specifies the duration of the billing period in days or
// months.
type PlanExternalPlanIDVersionNewParamsAddPricesPriceNewPlanMatrixPriceBillingCycleConfiguration struct {
	// The duration of the billing period.
	Duration param.Field[int64] `json:"duration,required"`
	// The unit of billing period duration.
	DurationUnit param.Field[PlanExternalPlanIDVersionNewParamsAddPricesPriceNewPlanMatrixPriceBillingCycleConfigurationDurationUnit] `json:"duration_unit,required"`
}

func (r PlanExternalPlanIDVersionNewParamsAddPricesPriceNewPlanMatrixPriceBillingCycleConfiguration) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// The unit of billing period duration.
type PlanExternalPlanIDVersionNewParamsAddPricesPriceNewPlanMatrixPriceBillingCycleConfigurationDurationUnit string

const (
	PlanExternalPlanIDVersionNewParamsAddPricesPriceNewPlanMatrixPriceBillingCycleConfigurationDurationUnitDay   PlanExternalPlanIDVersionNewParamsAddPricesPriceNewPlanMatrixPriceBillingCycleConfigurationDurationUnit = "day"
	PlanExternalPlanIDVersionNewParamsAddPricesPriceNewPlanMatrixPriceBillingCycleConfigurationDurationUnitMonth PlanExternalPlanIDVersionNewParamsAddPricesPriceNewPlanMatrixPriceBillingCycleConfigurationDurationUnit = "month"
)

func (r PlanExternalPlanIDVersionNewParamsAddPricesPriceNewPlanMatrixPriceBillingCycleConfigurationDurationUnit) IsKnown() bool {
	switch r {
	case PlanExternalPlanIDVersionNewParamsAddPricesPriceNewPlanMatrixPriceBillingCycleConfigurationDurationUnitDay, PlanExternalPlanIDVersionNewParamsAddPricesPriceNewPlanMatrixPriceBillingCycleConfigurationDurationUnitMonth:
		return true
	}
	return false
}

// For dimensional price: specifies a price group and dimension values
type PlanExternalPlanIDVersionNewParamsAddPricesPriceNewPlanMatrixPriceDimensionalPriceConfiguration struct {
	// The list of dimension values matching (in order) the dimensions of the price
	// group
	DimensionValues param.Field[[]string] `json:"dimension_values,required"`
	// The id of the dimensional price group to include this price in
	DimensionalPriceGroupID param.Field[string] `json:"dimensional_price_group_id"`
	// The external id of the dimensional price group to include this price in
	ExternalDimensionalPriceGroupID param.Field[string] `json:"external_dimensional_price_group_id"`
}

func (r PlanExternalPlanIDVersionNewParamsAddPricesPriceNewPlanMatrixPriceDimensionalPriceConfiguration) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Within each billing cycle, specifies the cadence at which invoices are produced.
// If unspecified, a single invoice is produced per billing cycle.
type PlanExternalPlanIDVersionNewParamsAddPricesPriceNewPlanMatrixPriceInvoicingCycleConfiguration struct {
	// The duration of the billing period.
	Duration param.Field[int64] `json:"duration,required"`
	// The unit of billing period duration.
	DurationUnit param.Field[PlanExternalPlanIDVersionNewParamsAddPricesPriceNewPlanMatrixPriceInvoicingCycleConfigurationDurationUnit] `json:"duration_unit,required"`
}

func (r PlanExternalPlanIDVersionNewParamsAddPricesPriceNewPlanMatrixPriceInvoicingCycleConfiguration) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// The unit of billing period duration.
type PlanExternalPlanIDVersionNewParamsAddPricesPriceNewPlanMatrixPriceInvoicingCycleConfigurationDurationUnit string

const (
	PlanExternalPlanIDVersionNewParamsAddPricesPriceNewPlanMatrixPriceInvoicingCycleConfigurationDurationUnitDay   PlanExternalPlanIDVersionNewParamsAddPricesPriceNewPlanMatrixPriceInvoicingCycleConfigurationDurationUnit = "day"
	PlanExternalPlanIDVersionNewParamsAddPricesPriceNewPlanMatrixPriceInvoicingCycleConfigurationDurationUnitMonth PlanExternalPlanIDVersionNewParamsAddPricesPriceNewPlanMatrixPriceInvoicingCycleConfigurationDurationUnit = "month"
)

func (r PlanExternalPlanIDVersionNewParamsAddPricesPriceNewPlanMatrixPriceInvoicingCycleConfigurationDurationUnit) IsKnown() bool {
	switch r {
	case PlanExternalPlanIDVersionNewParamsAddPricesPriceNewPlanMatrixPriceInvoicingCycleConfigurationDurationUnitDay, PlanExternalPlanIDVersionNewParamsAddPricesPriceNewPlanMatrixPriceInvoicingCycleConfigurationDurationUnitMonth:
		return true
	}
	return false
}

type PlanExternalPlanIDVersionNewParamsAddPricesPriceNewPlanTieredPrice struct {
	// The cadence to bill for this price on.
	Cadence param.Field[PlanExternalPlanIDVersionNewParamsAddPricesPriceNewPlanTieredPriceCadence] `json:"cadence,required"`
	// The id of the item the price will be associated with.
	ItemID    param.Field[string]                                                                      `json:"item_id,required"`
	ModelType param.Field[PlanExternalPlanIDVersionNewParamsAddPricesPriceNewPlanTieredPriceModelType] `json:"model_type,required"`
	// The name of the price.
	Name         param.Field[string]                                                                         `json:"name,required"`
	TieredConfig param.Field[PlanExternalPlanIDVersionNewParamsAddPricesPriceNewPlanTieredPriceTieredConfig] `json:"tiered_config,required"`
	// The id of the billable metric for the price. Only needed if the price is
	// usage-based.
	BillableMetricID param.Field[string] `json:"billable_metric_id"`
	// If the Price represents a fixed cost, the price will be billed in-advance if
	// this is true, and in-arrears if this is false.
	BilledInAdvance param.Field[bool] `json:"billed_in_advance"`
	// For custom cadence: specifies the duration of the billing period in days or
	// months.
	BillingCycleConfiguration param.Field[PlanExternalPlanIDVersionNewParamsAddPricesPriceNewPlanTieredPriceBillingCycleConfiguration] `json:"billing_cycle_configuration"`
	// The per unit conversion rate of the price currency to the invoicing currency.
	ConversionRate param.Field[float64] `json:"conversion_rate"`
	// An ISO 4217 currency string, or custom pricing unit identifier, in which this
	// price is billed.
	Currency param.Field[string] `json:"currency"`
	// For dimensional price: specifies a price group and dimension values
	DimensionalPriceConfiguration param.Field[PlanExternalPlanIDVersionNewParamsAddPricesPriceNewPlanTieredPriceDimensionalPriceConfiguration] `json:"dimensional_price_configuration"`
	// An alias for the price.
	ExternalPriceID param.Field[string] `json:"external_price_id"`
	// If the Price represents a fixed cost, this represents the quantity of units
	// applied.
	FixedPriceQuantity param.Field[float64] `json:"fixed_price_quantity"`
	// The property used to group this price on an invoice
	InvoiceGroupingKey param.Field[string] `json:"invoice_grouping_key"`
	// Within each billing cycle, specifies the cadence at which invoices are produced.
	// If unspecified, a single invoice is produced per billing cycle.
	InvoicingCycleConfiguration param.Field[PlanExternalPlanIDVersionNewParamsAddPricesPriceNewPlanTieredPriceInvoicingCycleConfiguration] `json:"invoicing_cycle_configuration"`
	// User-specified key/value pairs for the resource. Individual keys can be removed
	// by setting the value to `null`, and the entire metadata mapping can be cleared
	// by setting `metadata` to `null`.
	Metadata param.Field[map[string]string] `json:"metadata"`
}

func (r PlanExternalPlanIDVersionNewParamsAddPricesPriceNewPlanTieredPrice) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r PlanExternalPlanIDVersionNewParamsAddPricesPriceNewPlanTieredPrice) implementsPlanExternalPlanIDVersionNewParamsAddPricesPriceUnion() {
}

// The cadence to bill for this price on.
type PlanExternalPlanIDVersionNewParamsAddPricesPriceNewPlanTieredPriceCadence string

const (
	PlanExternalPlanIDVersionNewParamsAddPricesPriceNewPlanTieredPriceCadenceAnnual     PlanExternalPlanIDVersionNewParamsAddPricesPriceNewPlanTieredPriceCadence = "annual"
	PlanExternalPlanIDVersionNewParamsAddPricesPriceNewPlanTieredPriceCadenceSemiAnnual PlanExternalPlanIDVersionNewParamsAddPricesPriceNewPlanTieredPriceCadence = "semi_annual"
	PlanExternalPlanIDVersionNewParamsAddPricesPriceNewPlanTieredPriceCadenceMonthly    PlanExternalPlanIDVersionNewParamsAddPricesPriceNewPlanTieredPriceCadence = "monthly"
	PlanExternalPlanIDVersionNewParamsAddPricesPriceNewPlanTieredPriceCadenceQuarterly  PlanExternalPlanIDVersionNewParamsAddPricesPriceNewPlanTieredPriceCadence = "quarterly"
	PlanExternalPlanIDVersionNewParamsAddPricesPriceNewPlanTieredPriceCadenceOneTime    PlanExternalPlanIDVersionNewParamsAddPricesPriceNewPlanTieredPriceCadence = "one_time"
	PlanExternalPlanIDVersionNewParamsAddPricesPriceNewPlanTieredPriceCadenceCustom     PlanExternalPlanIDVersionNewParamsAddPricesPriceNewPlanTieredPriceCadence = "custom"
)

func (r PlanExternalPlanIDVersionNewParamsAddPricesPriceNewPlanTieredPriceCadence) IsKnown() bool {
	switch r {
	case PlanExternalPlanIDVersionNewParamsAddPricesPriceNewPlanTieredPriceCadenceAnnual, PlanExternalPlanIDVersionNewParamsAddPricesPriceNewPlanTieredPriceCadenceSemiAnnual, PlanExternalPlanIDVersionNewParamsAddPricesPriceNewPlanTieredPriceCadenceMonthly, PlanExternalPlanIDVersionNewParamsAddPricesPriceNewPlanTieredPriceCadenceQuarterly, PlanExternalPlanIDVersionNewParamsAddPricesPriceNewPlanTieredPriceCadenceOneTime, PlanExternalPlanIDVersionNewParamsAddPricesPriceNewPlanTieredPriceCadenceCustom:
		return true
	}
	return false
}

type PlanExternalPlanIDVersionNewParamsAddPricesPriceNewPlanTieredPriceModelType string

const (
	PlanExternalPlanIDVersionNewParamsAddPricesPriceNewPlanTieredPriceModelTypeTiered PlanExternalPlanIDVersionNewParamsAddPricesPriceNewPlanTieredPriceModelType = "tiered"
)

func (r PlanExternalPlanIDVersionNewParamsAddPricesPriceNewPlanTieredPriceModelType) IsKnown() bool {
	switch r {
	case PlanExternalPlanIDVersionNewParamsAddPricesPriceNewPlanTieredPriceModelTypeTiered:
		return true
	}
	return false
}

type PlanExternalPlanIDVersionNewParamsAddPricesPriceNewPlanTieredPriceTieredConfig struct {
	// Tiers for rating based on total usage quantities into the specified tier
	Tiers param.Field[[]PlanExternalPlanIDVersionNewParamsAddPricesPriceNewPlanTieredPriceTieredConfigTier] `json:"tiers,required"`
}

func (r PlanExternalPlanIDVersionNewParamsAddPricesPriceNewPlanTieredPriceTieredConfig) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type PlanExternalPlanIDVersionNewParamsAddPricesPriceNewPlanTieredPriceTieredConfigTier struct {
	// Exclusive tier starting value
	FirstUnit param.Field[float64] `json:"first_unit,required"`
	// Amount per unit
	UnitAmount param.Field[string] `json:"unit_amount,required"`
	// Inclusive tier ending value. If null, this is treated as the last tier
	LastUnit param.Field[float64] `json:"last_unit"`
}

func (r PlanExternalPlanIDVersionNewParamsAddPricesPriceNewPlanTieredPriceTieredConfigTier) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// For custom cadence: specifies the duration of the billing period in days or
// months.
type PlanExternalPlanIDVersionNewParamsAddPricesPriceNewPlanTieredPriceBillingCycleConfiguration struct {
	// The duration of the billing period.
	Duration param.Field[int64] `json:"duration,required"`
	// The unit of billing period duration.
	DurationUnit param.Field[PlanExternalPlanIDVersionNewParamsAddPricesPriceNewPlanTieredPriceBillingCycleConfigurationDurationUnit] `json:"duration_unit,required"`
}

func (r PlanExternalPlanIDVersionNewParamsAddPricesPriceNewPlanTieredPriceBillingCycleConfiguration) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// The unit of billing period duration.
type PlanExternalPlanIDVersionNewParamsAddPricesPriceNewPlanTieredPriceBillingCycleConfigurationDurationUnit string

const (
	PlanExternalPlanIDVersionNewParamsAddPricesPriceNewPlanTieredPriceBillingCycleConfigurationDurationUnitDay   PlanExternalPlanIDVersionNewParamsAddPricesPriceNewPlanTieredPriceBillingCycleConfigurationDurationUnit = "day"
	PlanExternalPlanIDVersionNewParamsAddPricesPriceNewPlanTieredPriceBillingCycleConfigurationDurationUnitMonth PlanExternalPlanIDVersionNewParamsAddPricesPriceNewPlanTieredPriceBillingCycleConfigurationDurationUnit = "month"
)

func (r PlanExternalPlanIDVersionNewParamsAddPricesPriceNewPlanTieredPriceBillingCycleConfigurationDurationUnit) IsKnown() bool {
	switch r {
	case PlanExternalPlanIDVersionNewParamsAddPricesPriceNewPlanTieredPriceBillingCycleConfigurationDurationUnitDay, PlanExternalPlanIDVersionNewParamsAddPricesPriceNewPlanTieredPriceBillingCycleConfigurationDurationUnitMonth:
		return true
	}
	return false
}

// For dimensional price: specifies a price group and dimension values
type PlanExternalPlanIDVersionNewParamsAddPricesPriceNewPlanTieredPriceDimensionalPriceConfiguration struct {
	// The list of dimension values matching (in order) the dimensions of the price
	// group
	DimensionValues param.Field[[]string] `json:"dimension_values,required"`
	// The id of the dimensional price group to include this price in
	DimensionalPriceGroupID param.Field[string] `json:"dimensional_price_group_id"`
	// The external id of the dimensional price group to include this price in
	ExternalDimensionalPriceGroupID param.Field[string] `json:"external_dimensional_price_group_id"`
}

func (r PlanExternalPlanIDVersionNewParamsAddPricesPriceNewPlanTieredPriceDimensionalPriceConfiguration) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Within each billing cycle, specifies the cadence at which invoices are produced.
// If unspecified, a single invoice is produced per billing cycle.
type PlanExternalPlanIDVersionNewParamsAddPricesPriceNewPlanTieredPriceInvoicingCycleConfiguration struct {
	// The duration of the billing period.
	Duration param.Field[int64] `json:"duration,required"`
	// The unit of billing period duration.
	DurationUnit param.Field[PlanExternalPlanIDVersionNewParamsAddPricesPriceNewPlanTieredPriceInvoicingCycleConfigurationDurationUnit] `json:"duration_unit,required"`
}

func (r PlanExternalPlanIDVersionNewParamsAddPricesPriceNewPlanTieredPriceInvoicingCycleConfiguration) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// The unit of billing period duration.
type PlanExternalPlanIDVersionNewParamsAddPricesPriceNewPlanTieredPriceInvoicingCycleConfigurationDurationUnit string

const (
	PlanExternalPlanIDVersionNewParamsAddPricesPriceNewPlanTieredPriceInvoicingCycleConfigurationDurationUnitDay   PlanExternalPlanIDVersionNewParamsAddPricesPriceNewPlanTieredPriceInvoicingCycleConfigurationDurationUnit = "day"
	PlanExternalPlanIDVersionNewParamsAddPricesPriceNewPlanTieredPriceInvoicingCycleConfigurationDurationUnitMonth PlanExternalPlanIDVersionNewParamsAddPricesPriceNewPlanTieredPriceInvoicingCycleConfigurationDurationUnit = "month"
)

func (r PlanExternalPlanIDVersionNewParamsAddPricesPriceNewPlanTieredPriceInvoicingCycleConfigurationDurationUnit) IsKnown() bool {
	switch r {
	case PlanExternalPlanIDVersionNewParamsAddPricesPriceNewPlanTieredPriceInvoicingCycleConfigurationDurationUnitDay, PlanExternalPlanIDVersionNewParamsAddPricesPriceNewPlanTieredPriceInvoicingCycleConfigurationDurationUnitMonth:
		return true
	}
	return false
}

type PlanExternalPlanIDVersionNewParamsAddPricesPriceNewPlanTieredBpsPrice struct {
	// The cadence to bill for this price on.
	Cadence param.Field[PlanExternalPlanIDVersionNewParamsAddPricesPriceNewPlanTieredBpsPriceCadence] `json:"cadence,required"`
	// The id of the item the price will be associated with.
	ItemID    param.Field[string]                                                                         `json:"item_id,required"`
	ModelType param.Field[PlanExternalPlanIDVersionNewParamsAddPricesPriceNewPlanTieredBpsPriceModelType] `json:"model_type,required"`
	// The name of the price.
	Name            param.Field[string]                                                                               `json:"name,required"`
	TieredBpsConfig param.Field[PlanExternalPlanIDVersionNewParamsAddPricesPriceNewPlanTieredBpsPriceTieredBpsConfig] `json:"tiered_bps_config,required"`
	// The id of the billable metric for the price. Only needed if the price is
	// usage-based.
	BillableMetricID param.Field[string] `json:"billable_metric_id"`
	// If the Price represents a fixed cost, the price will be billed in-advance if
	// this is true, and in-arrears if this is false.
	BilledInAdvance param.Field[bool] `json:"billed_in_advance"`
	// For custom cadence: specifies the duration of the billing period in days or
	// months.
	BillingCycleConfiguration param.Field[PlanExternalPlanIDVersionNewParamsAddPricesPriceNewPlanTieredBpsPriceBillingCycleConfiguration] `json:"billing_cycle_configuration"`
	// The per unit conversion rate of the price currency to the invoicing currency.
	ConversionRate param.Field[float64] `json:"conversion_rate"`
	// An ISO 4217 currency string, or custom pricing unit identifier, in which this
	// price is billed.
	Currency param.Field[string] `json:"currency"`
	// For dimensional price: specifies a price group and dimension values
	DimensionalPriceConfiguration param.Field[PlanExternalPlanIDVersionNewParamsAddPricesPriceNewPlanTieredBpsPriceDimensionalPriceConfiguration] `json:"dimensional_price_configuration"`
	// An alias for the price.
	ExternalPriceID param.Field[string] `json:"external_price_id"`
	// If the Price represents a fixed cost, this represents the quantity of units
	// applied.
	FixedPriceQuantity param.Field[float64] `json:"fixed_price_quantity"`
	// The property used to group this price on an invoice
	InvoiceGroupingKey param.Field[string] `json:"invoice_grouping_key"`
	// Within each billing cycle, specifies the cadence at which invoices are produced.
	// If unspecified, a single invoice is produced per billing cycle.
	InvoicingCycleConfiguration param.Field[PlanExternalPlanIDVersionNewParamsAddPricesPriceNewPlanTieredBpsPriceInvoicingCycleConfiguration] `json:"invoicing_cycle_configuration"`
	// User-specified key/value pairs for the resource. Individual keys can be removed
	// by setting the value to `null`, and the entire metadata mapping can be cleared
	// by setting `metadata` to `null`.
	Metadata param.Field[map[string]string] `json:"metadata"`
}

func (r PlanExternalPlanIDVersionNewParamsAddPricesPriceNewPlanTieredBpsPrice) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r PlanExternalPlanIDVersionNewParamsAddPricesPriceNewPlanTieredBpsPrice) implementsPlanExternalPlanIDVersionNewParamsAddPricesPriceUnion() {
}

// The cadence to bill for this price on.
type PlanExternalPlanIDVersionNewParamsAddPricesPriceNewPlanTieredBpsPriceCadence string

const (
	PlanExternalPlanIDVersionNewParamsAddPricesPriceNewPlanTieredBpsPriceCadenceAnnual     PlanExternalPlanIDVersionNewParamsAddPricesPriceNewPlanTieredBpsPriceCadence = "annual"
	PlanExternalPlanIDVersionNewParamsAddPricesPriceNewPlanTieredBpsPriceCadenceSemiAnnual PlanExternalPlanIDVersionNewParamsAddPricesPriceNewPlanTieredBpsPriceCadence = "semi_annual"
	PlanExternalPlanIDVersionNewParamsAddPricesPriceNewPlanTieredBpsPriceCadenceMonthly    PlanExternalPlanIDVersionNewParamsAddPricesPriceNewPlanTieredBpsPriceCadence = "monthly"
	PlanExternalPlanIDVersionNewParamsAddPricesPriceNewPlanTieredBpsPriceCadenceQuarterly  PlanExternalPlanIDVersionNewParamsAddPricesPriceNewPlanTieredBpsPriceCadence = "quarterly"
	PlanExternalPlanIDVersionNewParamsAddPricesPriceNewPlanTieredBpsPriceCadenceOneTime    PlanExternalPlanIDVersionNewParamsAddPricesPriceNewPlanTieredBpsPriceCadence = "one_time"
	PlanExternalPlanIDVersionNewParamsAddPricesPriceNewPlanTieredBpsPriceCadenceCustom     PlanExternalPlanIDVersionNewParamsAddPricesPriceNewPlanTieredBpsPriceCadence = "custom"
)

func (r PlanExternalPlanIDVersionNewParamsAddPricesPriceNewPlanTieredBpsPriceCadence) IsKnown() bool {
	switch r {
	case PlanExternalPlanIDVersionNewParamsAddPricesPriceNewPlanTieredBpsPriceCadenceAnnual, PlanExternalPlanIDVersionNewParamsAddPricesPriceNewPlanTieredBpsPriceCadenceSemiAnnual, PlanExternalPlanIDVersionNewParamsAddPricesPriceNewPlanTieredBpsPriceCadenceMonthly, PlanExternalPlanIDVersionNewParamsAddPricesPriceNewPlanTieredBpsPriceCadenceQuarterly, PlanExternalPlanIDVersionNewParamsAddPricesPriceNewPlanTieredBpsPriceCadenceOneTime, PlanExternalPlanIDVersionNewParamsAddPricesPriceNewPlanTieredBpsPriceCadenceCustom:
		return true
	}
	return false
}

type PlanExternalPlanIDVersionNewParamsAddPricesPriceNewPlanTieredBpsPriceModelType string

const (
	PlanExternalPlanIDVersionNewParamsAddPricesPriceNewPlanTieredBpsPriceModelTypeTieredBps PlanExternalPlanIDVersionNewParamsAddPricesPriceNewPlanTieredBpsPriceModelType = "tiered_bps"
)

func (r PlanExternalPlanIDVersionNewParamsAddPricesPriceNewPlanTieredBpsPriceModelType) IsKnown() bool {
	switch r {
	case PlanExternalPlanIDVersionNewParamsAddPricesPriceNewPlanTieredBpsPriceModelTypeTieredBps:
		return true
	}
	return false
}

type PlanExternalPlanIDVersionNewParamsAddPricesPriceNewPlanTieredBpsPriceTieredBpsConfig struct {
	// Tiers for a Graduated BPS pricing model, where usage is bucketed into specified
	// tiers
	Tiers param.Field[[]PlanExternalPlanIDVersionNewParamsAddPricesPriceNewPlanTieredBpsPriceTieredBpsConfigTier] `json:"tiers,required"`
}

func (r PlanExternalPlanIDVersionNewParamsAddPricesPriceNewPlanTieredBpsPriceTieredBpsConfig) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type PlanExternalPlanIDVersionNewParamsAddPricesPriceNewPlanTieredBpsPriceTieredBpsConfigTier struct {
	// Per-event basis point rate
	Bps param.Field[float64] `json:"bps,required"`
	// Exclusive tier starting value
	MinimumAmount param.Field[string] `json:"minimum_amount,required"`
	// Inclusive tier ending value
	MaximumAmount param.Field[string] `json:"maximum_amount"`
	// Per unit maximum to charge
	PerUnitMaximum param.Field[string] `json:"per_unit_maximum"`
}

func (r PlanExternalPlanIDVersionNewParamsAddPricesPriceNewPlanTieredBpsPriceTieredBpsConfigTier) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// For custom cadence: specifies the duration of the billing period in days or
// months.
type PlanExternalPlanIDVersionNewParamsAddPricesPriceNewPlanTieredBpsPriceBillingCycleConfiguration struct {
	// The duration of the billing period.
	Duration param.Field[int64] `json:"duration,required"`
	// The unit of billing period duration.
	DurationUnit param.Field[PlanExternalPlanIDVersionNewParamsAddPricesPriceNewPlanTieredBpsPriceBillingCycleConfigurationDurationUnit] `json:"duration_unit,required"`
}

func (r PlanExternalPlanIDVersionNewParamsAddPricesPriceNewPlanTieredBpsPriceBillingCycleConfiguration) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// The unit of billing period duration.
type PlanExternalPlanIDVersionNewParamsAddPricesPriceNewPlanTieredBpsPriceBillingCycleConfigurationDurationUnit string

const (
	PlanExternalPlanIDVersionNewParamsAddPricesPriceNewPlanTieredBpsPriceBillingCycleConfigurationDurationUnitDay   PlanExternalPlanIDVersionNewParamsAddPricesPriceNewPlanTieredBpsPriceBillingCycleConfigurationDurationUnit = "day"
	PlanExternalPlanIDVersionNewParamsAddPricesPriceNewPlanTieredBpsPriceBillingCycleConfigurationDurationUnitMonth PlanExternalPlanIDVersionNewParamsAddPricesPriceNewPlanTieredBpsPriceBillingCycleConfigurationDurationUnit = "month"
)

func (r PlanExternalPlanIDVersionNewParamsAddPricesPriceNewPlanTieredBpsPriceBillingCycleConfigurationDurationUnit) IsKnown() bool {
	switch r {
	case PlanExternalPlanIDVersionNewParamsAddPricesPriceNewPlanTieredBpsPriceBillingCycleConfigurationDurationUnitDay, PlanExternalPlanIDVersionNewParamsAddPricesPriceNewPlanTieredBpsPriceBillingCycleConfigurationDurationUnitMonth:
		return true
	}
	return false
}

// For dimensional price: specifies a price group and dimension values
type PlanExternalPlanIDVersionNewParamsAddPricesPriceNewPlanTieredBpsPriceDimensionalPriceConfiguration struct {
	// The list of dimension values matching (in order) the dimensions of the price
	// group
	DimensionValues param.Field[[]string] `json:"dimension_values,required"`
	// The id of the dimensional price group to include this price in
	DimensionalPriceGroupID param.Field[string] `json:"dimensional_price_group_id"`
	// The external id of the dimensional price group to include this price in
	ExternalDimensionalPriceGroupID param.Field[string] `json:"external_dimensional_price_group_id"`
}

func (r PlanExternalPlanIDVersionNewParamsAddPricesPriceNewPlanTieredBpsPriceDimensionalPriceConfiguration) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Within each billing cycle, specifies the cadence at which invoices are produced.
// If unspecified, a single invoice is produced per billing cycle.
type PlanExternalPlanIDVersionNewParamsAddPricesPriceNewPlanTieredBpsPriceInvoicingCycleConfiguration struct {
	// The duration of the billing period.
	Duration param.Field[int64] `json:"duration,required"`
	// The unit of billing period duration.
	DurationUnit param.Field[PlanExternalPlanIDVersionNewParamsAddPricesPriceNewPlanTieredBpsPriceInvoicingCycleConfigurationDurationUnit] `json:"duration_unit,required"`
}

func (r PlanExternalPlanIDVersionNewParamsAddPricesPriceNewPlanTieredBpsPriceInvoicingCycleConfiguration) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// The unit of billing period duration.
type PlanExternalPlanIDVersionNewParamsAddPricesPriceNewPlanTieredBpsPriceInvoicingCycleConfigurationDurationUnit string

const (
	PlanExternalPlanIDVersionNewParamsAddPricesPriceNewPlanTieredBpsPriceInvoicingCycleConfigurationDurationUnitDay   PlanExternalPlanIDVersionNewParamsAddPricesPriceNewPlanTieredBpsPriceInvoicingCycleConfigurationDurationUnit = "day"
	PlanExternalPlanIDVersionNewParamsAddPricesPriceNewPlanTieredBpsPriceInvoicingCycleConfigurationDurationUnitMonth PlanExternalPlanIDVersionNewParamsAddPricesPriceNewPlanTieredBpsPriceInvoicingCycleConfigurationDurationUnit = "month"
)

func (r PlanExternalPlanIDVersionNewParamsAddPricesPriceNewPlanTieredBpsPriceInvoicingCycleConfigurationDurationUnit) IsKnown() bool {
	switch r {
	case PlanExternalPlanIDVersionNewParamsAddPricesPriceNewPlanTieredBpsPriceInvoicingCycleConfigurationDurationUnitDay, PlanExternalPlanIDVersionNewParamsAddPricesPriceNewPlanTieredBpsPriceInvoicingCycleConfigurationDurationUnitMonth:
		return true
	}
	return false
}

type PlanExternalPlanIDVersionNewParamsAddPricesPriceNewPlanBpsPrice struct {
	BpsConfig param.Field[PlanExternalPlanIDVersionNewParamsAddPricesPriceNewPlanBpsPriceBpsConfig] `json:"bps_config,required"`
	// The cadence to bill for this price on.
	Cadence param.Field[PlanExternalPlanIDVersionNewParamsAddPricesPriceNewPlanBpsPriceCadence] `json:"cadence,required"`
	// The id of the item the price will be associated with.
	ItemID    param.Field[string]                                                                   `json:"item_id,required"`
	ModelType param.Field[PlanExternalPlanIDVersionNewParamsAddPricesPriceNewPlanBpsPriceModelType] `json:"model_type,required"`
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
	BillingCycleConfiguration param.Field[PlanExternalPlanIDVersionNewParamsAddPricesPriceNewPlanBpsPriceBillingCycleConfiguration] `json:"billing_cycle_configuration"`
	// The per unit conversion rate of the price currency to the invoicing currency.
	ConversionRate param.Field[float64] `json:"conversion_rate"`
	// An ISO 4217 currency string, or custom pricing unit identifier, in which this
	// price is billed.
	Currency param.Field[string] `json:"currency"`
	// For dimensional price: specifies a price group and dimension values
	DimensionalPriceConfiguration param.Field[PlanExternalPlanIDVersionNewParamsAddPricesPriceNewPlanBpsPriceDimensionalPriceConfiguration] `json:"dimensional_price_configuration"`
	// An alias for the price.
	ExternalPriceID param.Field[string] `json:"external_price_id"`
	// If the Price represents a fixed cost, this represents the quantity of units
	// applied.
	FixedPriceQuantity param.Field[float64] `json:"fixed_price_quantity"`
	// The property used to group this price on an invoice
	InvoiceGroupingKey param.Field[string] `json:"invoice_grouping_key"`
	// Within each billing cycle, specifies the cadence at which invoices are produced.
	// If unspecified, a single invoice is produced per billing cycle.
	InvoicingCycleConfiguration param.Field[PlanExternalPlanIDVersionNewParamsAddPricesPriceNewPlanBpsPriceInvoicingCycleConfiguration] `json:"invoicing_cycle_configuration"`
	// User-specified key/value pairs for the resource. Individual keys can be removed
	// by setting the value to `null`, and the entire metadata mapping can be cleared
	// by setting `metadata` to `null`.
	Metadata param.Field[map[string]string] `json:"metadata"`
}

func (r PlanExternalPlanIDVersionNewParamsAddPricesPriceNewPlanBpsPrice) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r PlanExternalPlanIDVersionNewParamsAddPricesPriceNewPlanBpsPrice) implementsPlanExternalPlanIDVersionNewParamsAddPricesPriceUnion() {
}

type PlanExternalPlanIDVersionNewParamsAddPricesPriceNewPlanBpsPriceBpsConfig struct {
	// Basis point take rate per event
	Bps param.Field[float64] `json:"bps,required"`
	// Optional currency amount maximum to cap spend per event
	PerUnitMaximum param.Field[string] `json:"per_unit_maximum"`
}

func (r PlanExternalPlanIDVersionNewParamsAddPricesPriceNewPlanBpsPriceBpsConfig) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// The cadence to bill for this price on.
type PlanExternalPlanIDVersionNewParamsAddPricesPriceNewPlanBpsPriceCadence string

const (
	PlanExternalPlanIDVersionNewParamsAddPricesPriceNewPlanBpsPriceCadenceAnnual     PlanExternalPlanIDVersionNewParamsAddPricesPriceNewPlanBpsPriceCadence = "annual"
	PlanExternalPlanIDVersionNewParamsAddPricesPriceNewPlanBpsPriceCadenceSemiAnnual PlanExternalPlanIDVersionNewParamsAddPricesPriceNewPlanBpsPriceCadence = "semi_annual"
	PlanExternalPlanIDVersionNewParamsAddPricesPriceNewPlanBpsPriceCadenceMonthly    PlanExternalPlanIDVersionNewParamsAddPricesPriceNewPlanBpsPriceCadence = "monthly"
	PlanExternalPlanIDVersionNewParamsAddPricesPriceNewPlanBpsPriceCadenceQuarterly  PlanExternalPlanIDVersionNewParamsAddPricesPriceNewPlanBpsPriceCadence = "quarterly"
	PlanExternalPlanIDVersionNewParamsAddPricesPriceNewPlanBpsPriceCadenceOneTime    PlanExternalPlanIDVersionNewParamsAddPricesPriceNewPlanBpsPriceCadence = "one_time"
	PlanExternalPlanIDVersionNewParamsAddPricesPriceNewPlanBpsPriceCadenceCustom     PlanExternalPlanIDVersionNewParamsAddPricesPriceNewPlanBpsPriceCadence = "custom"
)

func (r PlanExternalPlanIDVersionNewParamsAddPricesPriceNewPlanBpsPriceCadence) IsKnown() bool {
	switch r {
	case PlanExternalPlanIDVersionNewParamsAddPricesPriceNewPlanBpsPriceCadenceAnnual, PlanExternalPlanIDVersionNewParamsAddPricesPriceNewPlanBpsPriceCadenceSemiAnnual, PlanExternalPlanIDVersionNewParamsAddPricesPriceNewPlanBpsPriceCadenceMonthly, PlanExternalPlanIDVersionNewParamsAddPricesPriceNewPlanBpsPriceCadenceQuarterly, PlanExternalPlanIDVersionNewParamsAddPricesPriceNewPlanBpsPriceCadenceOneTime, PlanExternalPlanIDVersionNewParamsAddPricesPriceNewPlanBpsPriceCadenceCustom:
		return true
	}
	return false
}

type PlanExternalPlanIDVersionNewParamsAddPricesPriceNewPlanBpsPriceModelType string

const (
	PlanExternalPlanIDVersionNewParamsAddPricesPriceNewPlanBpsPriceModelTypeBps PlanExternalPlanIDVersionNewParamsAddPricesPriceNewPlanBpsPriceModelType = "bps"
)

func (r PlanExternalPlanIDVersionNewParamsAddPricesPriceNewPlanBpsPriceModelType) IsKnown() bool {
	switch r {
	case PlanExternalPlanIDVersionNewParamsAddPricesPriceNewPlanBpsPriceModelTypeBps:
		return true
	}
	return false
}

// For custom cadence: specifies the duration of the billing period in days or
// months.
type PlanExternalPlanIDVersionNewParamsAddPricesPriceNewPlanBpsPriceBillingCycleConfiguration struct {
	// The duration of the billing period.
	Duration param.Field[int64] `json:"duration,required"`
	// The unit of billing period duration.
	DurationUnit param.Field[PlanExternalPlanIDVersionNewParamsAddPricesPriceNewPlanBpsPriceBillingCycleConfigurationDurationUnit] `json:"duration_unit,required"`
}

func (r PlanExternalPlanIDVersionNewParamsAddPricesPriceNewPlanBpsPriceBillingCycleConfiguration) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// The unit of billing period duration.
type PlanExternalPlanIDVersionNewParamsAddPricesPriceNewPlanBpsPriceBillingCycleConfigurationDurationUnit string

const (
	PlanExternalPlanIDVersionNewParamsAddPricesPriceNewPlanBpsPriceBillingCycleConfigurationDurationUnitDay   PlanExternalPlanIDVersionNewParamsAddPricesPriceNewPlanBpsPriceBillingCycleConfigurationDurationUnit = "day"
	PlanExternalPlanIDVersionNewParamsAddPricesPriceNewPlanBpsPriceBillingCycleConfigurationDurationUnitMonth PlanExternalPlanIDVersionNewParamsAddPricesPriceNewPlanBpsPriceBillingCycleConfigurationDurationUnit = "month"
)

func (r PlanExternalPlanIDVersionNewParamsAddPricesPriceNewPlanBpsPriceBillingCycleConfigurationDurationUnit) IsKnown() bool {
	switch r {
	case PlanExternalPlanIDVersionNewParamsAddPricesPriceNewPlanBpsPriceBillingCycleConfigurationDurationUnitDay, PlanExternalPlanIDVersionNewParamsAddPricesPriceNewPlanBpsPriceBillingCycleConfigurationDurationUnitMonth:
		return true
	}
	return false
}

// For dimensional price: specifies a price group and dimension values
type PlanExternalPlanIDVersionNewParamsAddPricesPriceNewPlanBpsPriceDimensionalPriceConfiguration struct {
	// The list of dimension values matching (in order) the dimensions of the price
	// group
	DimensionValues param.Field[[]string] `json:"dimension_values,required"`
	// The id of the dimensional price group to include this price in
	DimensionalPriceGroupID param.Field[string] `json:"dimensional_price_group_id"`
	// The external id of the dimensional price group to include this price in
	ExternalDimensionalPriceGroupID param.Field[string] `json:"external_dimensional_price_group_id"`
}

func (r PlanExternalPlanIDVersionNewParamsAddPricesPriceNewPlanBpsPriceDimensionalPriceConfiguration) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Within each billing cycle, specifies the cadence at which invoices are produced.
// If unspecified, a single invoice is produced per billing cycle.
type PlanExternalPlanIDVersionNewParamsAddPricesPriceNewPlanBpsPriceInvoicingCycleConfiguration struct {
	// The duration of the billing period.
	Duration param.Field[int64] `json:"duration,required"`
	// The unit of billing period duration.
	DurationUnit param.Field[PlanExternalPlanIDVersionNewParamsAddPricesPriceNewPlanBpsPriceInvoicingCycleConfigurationDurationUnit] `json:"duration_unit,required"`
}

func (r PlanExternalPlanIDVersionNewParamsAddPricesPriceNewPlanBpsPriceInvoicingCycleConfiguration) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// The unit of billing period duration.
type PlanExternalPlanIDVersionNewParamsAddPricesPriceNewPlanBpsPriceInvoicingCycleConfigurationDurationUnit string

const (
	PlanExternalPlanIDVersionNewParamsAddPricesPriceNewPlanBpsPriceInvoicingCycleConfigurationDurationUnitDay   PlanExternalPlanIDVersionNewParamsAddPricesPriceNewPlanBpsPriceInvoicingCycleConfigurationDurationUnit = "day"
	PlanExternalPlanIDVersionNewParamsAddPricesPriceNewPlanBpsPriceInvoicingCycleConfigurationDurationUnitMonth PlanExternalPlanIDVersionNewParamsAddPricesPriceNewPlanBpsPriceInvoicingCycleConfigurationDurationUnit = "month"
)

func (r PlanExternalPlanIDVersionNewParamsAddPricesPriceNewPlanBpsPriceInvoicingCycleConfigurationDurationUnit) IsKnown() bool {
	switch r {
	case PlanExternalPlanIDVersionNewParamsAddPricesPriceNewPlanBpsPriceInvoicingCycleConfigurationDurationUnitDay, PlanExternalPlanIDVersionNewParamsAddPricesPriceNewPlanBpsPriceInvoicingCycleConfigurationDurationUnitMonth:
		return true
	}
	return false
}

type PlanExternalPlanIDVersionNewParamsAddPricesPriceNewPlanBulkBpsPrice struct {
	BulkBpsConfig param.Field[PlanExternalPlanIDVersionNewParamsAddPricesPriceNewPlanBulkBpsPriceBulkBpsConfig] `json:"bulk_bps_config,required"`
	// The cadence to bill for this price on.
	Cadence param.Field[PlanExternalPlanIDVersionNewParamsAddPricesPriceNewPlanBulkBpsPriceCadence] `json:"cadence,required"`
	// The id of the item the price will be associated with.
	ItemID    param.Field[string]                                                                       `json:"item_id,required"`
	ModelType param.Field[PlanExternalPlanIDVersionNewParamsAddPricesPriceNewPlanBulkBpsPriceModelType] `json:"model_type,required"`
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
	BillingCycleConfiguration param.Field[PlanExternalPlanIDVersionNewParamsAddPricesPriceNewPlanBulkBpsPriceBillingCycleConfiguration] `json:"billing_cycle_configuration"`
	// The per unit conversion rate of the price currency to the invoicing currency.
	ConversionRate param.Field[float64] `json:"conversion_rate"`
	// An ISO 4217 currency string, or custom pricing unit identifier, in which this
	// price is billed.
	Currency param.Field[string] `json:"currency"`
	// For dimensional price: specifies a price group and dimension values
	DimensionalPriceConfiguration param.Field[PlanExternalPlanIDVersionNewParamsAddPricesPriceNewPlanBulkBpsPriceDimensionalPriceConfiguration] `json:"dimensional_price_configuration"`
	// An alias for the price.
	ExternalPriceID param.Field[string] `json:"external_price_id"`
	// If the Price represents a fixed cost, this represents the quantity of units
	// applied.
	FixedPriceQuantity param.Field[float64] `json:"fixed_price_quantity"`
	// The property used to group this price on an invoice
	InvoiceGroupingKey param.Field[string] `json:"invoice_grouping_key"`
	// Within each billing cycle, specifies the cadence at which invoices are produced.
	// If unspecified, a single invoice is produced per billing cycle.
	InvoicingCycleConfiguration param.Field[PlanExternalPlanIDVersionNewParamsAddPricesPriceNewPlanBulkBpsPriceInvoicingCycleConfiguration] `json:"invoicing_cycle_configuration"`
	// User-specified key/value pairs for the resource. Individual keys can be removed
	// by setting the value to `null`, and the entire metadata mapping can be cleared
	// by setting `metadata` to `null`.
	Metadata param.Field[map[string]string] `json:"metadata"`
}

func (r PlanExternalPlanIDVersionNewParamsAddPricesPriceNewPlanBulkBpsPrice) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r PlanExternalPlanIDVersionNewParamsAddPricesPriceNewPlanBulkBpsPrice) implementsPlanExternalPlanIDVersionNewParamsAddPricesPriceUnion() {
}

type PlanExternalPlanIDVersionNewParamsAddPricesPriceNewPlanBulkBpsPriceBulkBpsConfig struct {
	// Tiers for a bulk BPS pricing model where all usage is aggregated to a single
	// tier based on total volume
	Tiers param.Field[[]PlanExternalPlanIDVersionNewParamsAddPricesPriceNewPlanBulkBpsPriceBulkBpsConfigTier] `json:"tiers,required"`
}

func (r PlanExternalPlanIDVersionNewParamsAddPricesPriceNewPlanBulkBpsPriceBulkBpsConfig) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type PlanExternalPlanIDVersionNewParamsAddPricesPriceNewPlanBulkBpsPriceBulkBpsConfigTier struct {
	// Basis points to rate on
	Bps param.Field[float64] `json:"bps,required"`
	// Upper bound for tier
	MaximumAmount param.Field[string] `json:"maximum_amount"`
	// The maximum amount to charge for any one event
	PerUnitMaximum param.Field[string] `json:"per_unit_maximum"`
}

func (r PlanExternalPlanIDVersionNewParamsAddPricesPriceNewPlanBulkBpsPriceBulkBpsConfigTier) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// The cadence to bill for this price on.
type PlanExternalPlanIDVersionNewParamsAddPricesPriceNewPlanBulkBpsPriceCadence string

const (
	PlanExternalPlanIDVersionNewParamsAddPricesPriceNewPlanBulkBpsPriceCadenceAnnual     PlanExternalPlanIDVersionNewParamsAddPricesPriceNewPlanBulkBpsPriceCadence = "annual"
	PlanExternalPlanIDVersionNewParamsAddPricesPriceNewPlanBulkBpsPriceCadenceSemiAnnual PlanExternalPlanIDVersionNewParamsAddPricesPriceNewPlanBulkBpsPriceCadence = "semi_annual"
	PlanExternalPlanIDVersionNewParamsAddPricesPriceNewPlanBulkBpsPriceCadenceMonthly    PlanExternalPlanIDVersionNewParamsAddPricesPriceNewPlanBulkBpsPriceCadence = "monthly"
	PlanExternalPlanIDVersionNewParamsAddPricesPriceNewPlanBulkBpsPriceCadenceQuarterly  PlanExternalPlanIDVersionNewParamsAddPricesPriceNewPlanBulkBpsPriceCadence = "quarterly"
	PlanExternalPlanIDVersionNewParamsAddPricesPriceNewPlanBulkBpsPriceCadenceOneTime    PlanExternalPlanIDVersionNewParamsAddPricesPriceNewPlanBulkBpsPriceCadence = "one_time"
	PlanExternalPlanIDVersionNewParamsAddPricesPriceNewPlanBulkBpsPriceCadenceCustom     PlanExternalPlanIDVersionNewParamsAddPricesPriceNewPlanBulkBpsPriceCadence = "custom"
)

func (r PlanExternalPlanIDVersionNewParamsAddPricesPriceNewPlanBulkBpsPriceCadence) IsKnown() bool {
	switch r {
	case PlanExternalPlanIDVersionNewParamsAddPricesPriceNewPlanBulkBpsPriceCadenceAnnual, PlanExternalPlanIDVersionNewParamsAddPricesPriceNewPlanBulkBpsPriceCadenceSemiAnnual, PlanExternalPlanIDVersionNewParamsAddPricesPriceNewPlanBulkBpsPriceCadenceMonthly, PlanExternalPlanIDVersionNewParamsAddPricesPriceNewPlanBulkBpsPriceCadenceQuarterly, PlanExternalPlanIDVersionNewParamsAddPricesPriceNewPlanBulkBpsPriceCadenceOneTime, PlanExternalPlanIDVersionNewParamsAddPricesPriceNewPlanBulkBpsPriceCadenceCustom:
		return true
	}
	return false
}

type PlanExternalPlanIDVersionNewParamsAddPricesPriceNewPlanBulkBpsPriceModelType string

const (
	PlanExternalPlanIDVersionNewParamsAddPricesPriceNewPlanBulkBpsPriceModelTypeBulkBps PlanExternalPlanIDVersionNewParamsAddPricesPriceNewPlanBulkBpsPriceModelType = "bulk_bps"
)

func (r PlanExternalPlanIDVersionNewParamsAddPricesPriceNewPlanBulkBpsPriceModelType) IsKnown() bool {
	switch r {
	case PlanExternalPlanIDVersionNewParamsAddPricesPriceNewPlanBulkBpsPriceModelTypeBulkBps:
		return true
	}
	return false
}

// For custom cadence: specifies the duration of the billing period in days or
// months.
type PlanExternalPlanIDVersionNewParamsAddPricesPriceNewPlanBulkBpsPriceBillingCycleConfiguration struct {
	// The duration of the billing period.
	Duration param.Field[int64] `json:"duration,required"`
	// The unit of billing period duration.
	DurationUnit param.Field[PlanExternalPlanIDVersionNewParamsAddPricesPriceNewPlanBulkBpsPriceBillingCycleConfigurationDurationUnit] `json:"duration_unit,required"`
}

func (r PlanExternalPlanIDVersionNewParamsAddPricesPriceNewPlanBulkBpsPriceBillingCycleConfiguration) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// The unit of billing period duration.
type PlanExternalPlanIDVersionNewParamsAddPricesPriceNewPlanBulkBpsPriceBillingCycleConfigurationDurationUnit string

const (
	PlanExternalPlanIDVersionNewParamsAddPricesPriceNewPlanBulkBpsPriceBillingCycleConfigurationDurationUnitDay   PlanExternalPlanIDVersionNewParamsAddPricesPriceNewPlanBulkBpsPriceBillingCycleConfigurationDurationUnit = "day"
	PlanExternalPlanIDVersionNewParamsAddPricesPriceNewPlanBulkBpsPriceBillingCycleConfigurationDurationUnitMonth PlanExternalPlanIDVersionNewParamsAddPricesPriceNewPlanBulkBpsPriceBillingCycleConfigurationDurationUnit = "month"
)

func (r PlanExternalPlanIDVersionNewParamsAddPricesPriceNewPlanBulkBpsPriceBillingCycleConfigurationDurationUnit) IsKnown() bool {
	switch r {
	case PlanExternalPlanIDVersionNewParamsAddPricesPriceNewPlanBulkBpsPriceBillingCycleConfigurationDurationUnitDay, PlanExternalPlanIDVersionNewParamsAddPricesPriceNewPlanBulkBpsPriceBillingCycleConfigurationDurationUnitMonth:
		return true
	}
	return false
}

// For dimensional price: specifies a price group and dimension values
type PlanExternalPlanIDVersionNewParamsAddPricesPriceNewPlanBulkBpsPriceDimensionalPriceConfiguration struct {
	// The list of dimension values matching (in order) the dimensions of the price
	// group
	DimensionValues param.Field[[]string] `json:"dimension_values,required"`
	// The id of the dimensional price group to include this price in
	DimensionalPriceGroupID param.Field[string] `json:"dimensional_price_group_id"`
	// The external id of the dimensional price group to include this price in
	ExternalDimensionalPriceGroupID param.Field[string] `json:"external_dimensional_price_group_id"`
}

func (r PlanExternalPlanIDVersionNewParamsAddPricesPriceNewPlanBulkBpsPriceDimensionalPriceConfiguration) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Within each billing cycle, specifies the cadence at which invoices are produced.
// If unspecified, a single invoice is produced per billing cycle.
type PlanExternalPlanIDVersionNewParamsAddPricesPriceNewPlanBulkBpsPriceInvoicingCycleConfiguration struct {
	// The duration of the billing period.
	Duration param.Field[int64] `json:"duration,required"`
	// The unit of billing period duration.
	DurationUnit param.Field[PlanExternalPlanIDVersionNewParamsAddPricesPriceNewPlanBulkBpsPriceInvoicingCycleConfigurationDurationUnit] `json:"duration_unit,required"`
}

func (r PlanExternalPlanIDVersionNewParamsAddPricesPriceNewPlanBulkBpsPriceInvoicingCycleConfiguration) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// The unit of billing period duration.
type PlanExternalPlanIDVersionNewParamsAddPricesPriceNewPlanBulkBpsPriceInvoicingCycleConfigurationDurationUnit string

const (
	PlanExternalPlanIDVersionNewParamsAddPricesPriceNewPlanBulkBpsPriceInvoicingCycleConfigurationDurationUnitDay   PlanExternalPlanIDVersionNewParamsAddPricesPriceNewPlanBulkBpsPriceInvoicingCycleConfigurationDurationUnit = "day"
	PlanExternalPlanIDVersionNewParamsAddPricesPriceNewPlanBulkBpsPriceInvoicingCycleConfigurationDurationUnitMonth PlanExternalPlanIDVersionNewParamsAddPricesPriceNewPlanBulkBpsPriceInvoicingCycleConfigurationDurationUnit = "month"
)

func (r PlanExternalPlanIDVersionNewParamsAddPricesPriceNewPlanBulkBpsPriceInvoicingCycleConfigurationDurationUnit) IsKnown() bool {
	switch r {
	case PlanExternalPlanIDVersionNewParamsAddPricesPriceNewPlanBulkBpsPriceInvoicingCycleConfigurationDurationUnitDay, PlanExternalPlanIDVersionNewParamsAddPricesPriceNewPlanBulkBpsPriceInvoicingCycleConfigurationDurationUnitMonth:
		return true
	}
	return false
}

type PlanExternalPlanIDVersionNewParamsAddPricesPriceNewPlanBulkPrice struct {
	BulkConfig param.Field[PlanExternalPlanIDVersionNewParamsAddPricesPriceNewPlanBulkPriceBulkConfig] `json:"bulk_config,required"`
	// The cadence to bill for this price on.
	Cadence param.Field[PlanExternalPlanIDVersionNewParamsAddPricesPriceNewPlanBulkPriceCadence] `json:"cadence,required"`
	// The id of the item the price will be associated with.
	ItemID    param.Field[string]                                                                    `json:"item_id,required"`
	ModelType param.Field[PlanExternalPlanIDVersionNewParamsAddPricesPriceNewPlanBulkPriceModelType] `json:"model_type,required"`
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
	BillingCycleConfiguration param.Field[PlanExternalPlanIDVersionNewParamsAddPricesPriceNewPlanBulkPriceBillingCycleConfiguration] `json:"billing_cycle_configuration"`
	// The per unit conversion rate of the price currency to the invoicing currency.
	ConversionRate param.Field[float64] `json:"conversion_rate"`
	// An ISO 4217 currency string, or custom pricing unit identifier, in which this
	// price is billed.
	Currency param.Field[string] `json:"currency"`
	// For dimensional price: specifies a price group and dimension values
	DimensionalPriceConfiguration param.Field[PlanExternalPlanIDVersionNewParamsAddPricesPriceNewPlanBulkPriceDimensionalPriceConfiguration] `json:"dimensional_price_configuration"`
	// An alias for the price.
	ExternalPriceID param.Field[string] `json:"external_price_id"`
	// If the Price represents a fixed cost, this represents the quantity of units
	// applied.
	FixedPriceQuantity param.Field[float64] `json:"fixed_price_quantity"`
	// The property used to group this price on an invoice
	InvoiceGroupingKey param.Field[string] `json:"invoice_grouping_key"`
	// Within each billing cycle, specifies the cadence at which invoices are produced.
	// If unspecified, a single invoice is produced per billing cycle.
	InvoicingCycleConfiguration param.Field[PlanExternalPlanIDVersionNewParamsAddPricesPriceNewPlanBulkPriceInvoicingCycleConfiguration] `json:"invoicing_cycle_configuration"`
	// User-specified key/value pairs for the resource. Individual keys can be removed
	// by setting the value to `null`, and the entire metadata mapping can be cleared
	// by setting `metadata` to `null`.
	Metadata param.Field[map[string]string] `json:"metadata"`
}

func (r PlanExternalPlanIDVersionNewParamsAddPricesPriceNewPlanBulkPrice) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r PlanExternalPlanIDVersionNewParamsAddPricesPriceNewPlanBulkPrice) implementsPlanExternalPlanIDVersionNewParamsAddPricesPriceUnion() {
}

type PlanExternalPlanIDVersionNewParamsAddPricesPriceNewPlanBulkPriceBulkConfig struct {
	// Bulk tiers for rating based on total usage volume
	Tiers param.Field[[]PlanExternalPlanIDVersionNewParamsAddPricesPriceNewPlanBulkPriceBulkConfigTier] `json:"tiers,required"`
}

func (r PlanExternalPlanIDVersionNewParamsAddPricesPriceNewPlanBulkPriceBulkConfig) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type PlanExternalPlanIDVersionNewParamsAddPricesPriceNewPlanBulkPriceBulkConfigTier struct {
	// Amount per unit
	UnitAmount param.Field[string] `json:"unit_amount,required"`
	// Upper bound for this tier
	MaximumUnits param.Field[float64] `json:"maximum_units"`
}

func (r PlanExternalPlanIDVersionNewParamsAddPricesPriceNewPlanBulkPriceBulkConfigTier) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// The cadence to bill for this price on.
type PlanExternalPlanIDVersionNewParamsAddPricesPriceNewPlanBulkPriceCadence string

const (
	PlanExternalPlanIDVersionNewParamsAddPricesPriceNewPlanBulkPriceCadenceAnnual     PlanExternalPlanIDVersionNewParamsAddPricesPriceNewPlanBulkPriceCadence = "annual"
	PlanExternalPlanIDVersionNewParamsAddPricesPriceNewPlanBulkPriceCadenceSemiAnnual PlanExternalPlanIDVersionNewParamsAddPricesPriceNewPlanBulkPriceCadence = "semi_annual"
	PlanExternalPlanIDVersionNewParamsAddPricesPriceNewPlanBulkPriceCadenceMonthly    PlanExternalPlanIDVersionNewParamsAddPricesPriceNewPlanBulkPriceCadence = "monthly"
	PlanExternalPlanIDVersionNewParamsAddPricesPriceNewPlanBulkPriceCadenceQuarterly  PlanExternalPlanIDVersionNewParamsAddPricesPriceNewPlanBulkPriceCadence = "quarterly"
	PlanExternalPlanIDVersionNewParamsAddPricesPriceNewPlanBulkPriceCadenceOneTime    PlanExternalPlanIDVersionNewParamsAddPricesPriceNewPlanBulkPriceCadence = "one_time"
	PlanExternalPlanIDVersionNewParamsAddPricesPriceNewPlanBulkPriceCadenceCustom     PlanExternalPlanIDVersionNewParamsAddPricesPriceNewPlanBulkPriceCadence = "custom"
)

func (r PlanExternalPlanIDVersionNewParamsAddPricesPriceNewPlanBulkPriceCadence) IsKnown() bool {
	switch r {
	case PlanExternalPlanIDVersionNewParamsAddPricesPriceNewPlanBulkPriceCadenceAnnual, PlanExternalPlanIDVersionNewParamsAddPricesPriceNewPlanBulkPriceCadenceSemiAnnual, PlanExternalPlanIDVersionNewParamsAddPricesPriceNewPlanBulkPriceCadenceMonthly, PlanExternalPlanIDVersionNewParamsAddPricesPriceNewPlanBulkPriceCadenceQuarterly, PlanExternalPlanIDVersionNewParamsAddPricesPriceNewPlanBulkPriceCadenceOneTime, PlanExternalPlanIDVersionNewParamsAddPricesPriceNewPlanBulkPriceCadenceCustom:
		return true
	}
	return false
}

type PlanExternalPlanIDVersionNewParamsAddPricesPriceNewPlanBulkPriceModelType string

const (
	PlanExternalPlanIDVersionNewParamsAddPricesPriceNewPlanBulkPriceModelTypeBulk PlanExternalPlanIDVersionNewParamsAddPricesPriceNewPlanBulkPriceModelType = "bulk"
)

func (r PlanExternalPlanIDVersionNewParamsAddPricesPriceNewPlanBulkPriceModelType) IsKnown() bool {
	switch r {
	case PlanExternalPlanIDVersionNewParamsAddPricesPriceNewPlanBulkPriceModelTypeBulk:
		return true
	}
	return false
}

// For custom cadence: specifies the duration of the billing period in days or
// months.
type PlanExternalPlanIDVersionNewParamsAddPricesPriceNewPlanBulkPriceBillingCycleConfiguration struct {
	// The duration of the billing period.
	Duration param.Field[int64] `json:"duration,required"`
	// The unit of billing period duration.
	DurationUnit param.Field[PlanExternalPlanIDVersionNewParamsAddPricesPriceNewPlanBulkPriceBillingCycleConfigurationDurationUnit] `json:"duration_unit,required"`
}

func (r PlanExternalPlanIDVersionNewParamsAddPricesPriceNewPlanBulkPriceBillingCycleConfiguration) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// The unit of billing period duration.
type PlanExternalPlanIDVersionNewParamsAddPricesPriceNewPlanBulkPriceBillingCycleConfigurationDurationUnit string

const (
	PlanExternalPlanIDVersionNewParamsAddPricesPriceNewPlanBulkPriceBillingCycleConfigurationDurationUnitDay   PlanExternalPlanIDVersionNewParamsAddPricesPriceNewPlanBulkPriceBillingCycleConfigurationDurationUnit = "day"
	PlanExternalPlanIDVersionNewParamsAddPricesPriceNewPlanBulkPriceBillingCycleConfigurationDurationUnitMonth PlanExternalPlanIDVersionNewParamsAddPricesPriceNewPlanBulkPriceBillingCycleConfigurationDurationUnit = "month"
)

func (r PlanExternalPlanIDVersionNewParamsAddPricesPriceNewPlanBulkPriceBillingCycleConfigurationDurationUnit) IsKnown() bool {
	switch r {
	case PlanExternalPlanIDVersionNewParamsAddPricesPriceNewPlanBulkPriceBillingCycleConfigurationDurationUnitDay, PlanExternalPlanIDVersionNewParamsAddPricesPriceNewPlanBulkPriceBillingCycleConfigurationDurationUnitMonth:
		return true
	}
	return false
}

// For dimensional price: specifies a price group and dimension values
type PlanExternalPlanIDVersionNewParamsAddPricesPriceNewPlanBulkPriceDimensionalPriceConfiguration struct {
	// The list of dimension values matching (in order) the dimensions of the price
	// group
	DimensionValues param.Field[[]string] `json:"dimension_values,required"`
	// The id of the dimensional price group to include this price in
	DimensionalPriceGroupID param.Field[string] `json:"dimensional_price_group_id"`
	// The external id of the dimensional price group to include this price in
	ExternalDimensionalPriceGroupID param.Field[string] `json:"external_dimensional_price_group_id"`
}

func (r PlanExternalPlanIDVersionNewParamsAddPricesPriceNewPlanBulkPriceDimensionalPriceConfiguration) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Within each billing cycle, specifies the cadence at which invoices are produced.
// If unspecified, a single invoice is produced per billing cycle.
type PlanExternalPlanIDVersionNewParamsAddPricesPriceNewPlanBulkPriceInvoicingCycleConfiguration struct {
	// The duration of the billing period.
	Duration param.Field[int64] `json:"duration,required"`
	// The unit of billing period duration.
	DurationUnit param.Field[PlanExternalPlanIDVersionNewParamsAddPricesPriceNewPlanBulkPriceInvoicingCycleConfigurationDurationUnit] `json:"duration_unit,required"`
}

func (r PlanExternalPlanIDVersionNewParamsAddPricesPriceNewPlanBulkPriceInvoicingCycleConfiguration) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// The unit of billing period duration.
type PlanExternalPlanIDVersionNewParamsAddPricesPriceNewPlanBulkPriceInvoicingCycleConfigurationDurationUnit string

const (
	PlanExternalPlanIDVersionNewParamsAddPricesPriceNewPlanBulkPriceInvoicingCycleConfigurationDurationUnitDay   PlanExternalPlanIDVersionNewParamsAddPricesPriceNewPlanBulkPriceInvoicingCycleConfigurationDurationUnit = "day"
	PlanExternalPlanIDVersionNewParamsAddPricesPriceNewPlanBulkPriceInvoicingCycleConfigurationDurationUnitMonth PlanExternalPlanIDVersionNewParamsAddPricesPriceNewPlanBulkPriceInvoicingCycleConfigurationDurationUnit = "month"
)

func (r PlanExternalPlanIDVersionNewParamsAddPricesPriceNewPlanBulkPriceInvoicingCycleConfigurationDurationUnit) IsKnown() bool {
	switch r {
	case PlanExternalPlanIDVersionNewParamsAddPricesPriceNewPlanBulkPriceInvoicingCycleConfigurationDurationUnitDay, PlanExternalPlanIDVersionNewParamsAddPricesPriceNewPlanBulkPriceInvoicingCycleConfigurationDurationUnitMonth:
		return true
	}
	return false
}

type PlanExternalPlanIDVersionNewParamsAddPricesPriceNewPlanThresholdTotalAmountPrice struct {
	// The cadence to bill for this price on.
	Cadence param.Field[PlanExternalPlanIDVersionNewParamsAddPricesPriceNewPlanThresholdTotalAmountPriceCadence] `json:"cadence,required"`
	// The id of the item the price will be associated with.
	ItemID    param.Field[string]                                                                                    `json:"item_id,required"`
	ModelType param.Field[PlanExternalPlanIDVersionNewParamsAddPricesPriceNewPlanThresholdTotalAmountPriceModelType] `json:"model_type,required"`
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
	BillingCycleConfiguration param.Field[PlanExternalPlanIDVersionNewParamsAddPricesPriceNewPlanThresholdTotalAmountPriceBillingCycleConfiguration] `json:"billing_cycle_configuration"`
	// The per unit conversion rate of the price currency to the invoicing currency.
	ConversionRate param.Field[float64] `json:"conversion_rate"`
	// An ISO 4217 currency string, or custom pricing unit identifier, in which this
	// price is billed.
	Currency param.Field[string] `json:"currency"`
	// For dimensional price: specifies a price group and dimension values
	DimensionalPriceConfiguration param.Field[PlanExternalPlanIDVersionNewParamsAddPricesPriceNewPlanThresholdTotalAmountPriceDimensionalPriceConfiguration] `json:"dimensional_price_configuration"`
	// An alias for the price.
	ExternalPriceID param.Field[string] `json:"external_price_id"`
	// If the Price represents a fixed cost, this represents the quantity of units
	// applied.
	FixedPriceQuantity param.Field[float64] `json:"fixed_price_quantity"`
	// The property used to group this price on an invoice
	InvoiceGroupingKey param.Field[string] `json:"invoice_grouping_key"`
	// Within each billing cycle, specifies the cadence at which invoices are produced.
	// If unspecified, a single invoice is produced per billing cycle.
	InvoicingCycleConfiguration param.Field[PlanExternalPlanIDVersionNewParamsAddPricesPriceNewPlanThresholdTotalAmountPriceInvoicingCycleConfiguration] `json:"invoicing_cycle_configuration"`
	// User-specified key/value pairs for the resource. Individual keys can be removed
	// by setting the value to `null`, and the entire metadata mapping can be cleared
	// by setting `metadata` to `null`.
	Metadata param.Field[map[string]string] `json:"metadata"`
}

func (r PlanExternalPlanIDVersionNewParamsAddPricesPriceNewPlanThresholdTotalAmountPrice) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r PlanExternalPlanIDVersionNewParamsAddPricesPriceNewPlanThresholdTotalAmountPrice) implementsPlanExternalPlanIDVersionNewParamsAddPricesPriceUnion() {
}

// The cadence to bill for this price on.
type PlanExternalPlanIDVersionNewParamsAddPricesPriceNewPlanThresholdTotalAmountPriceCadence string

const (
	PlanExternalPlanIDVersionNewParamsAddPricesPriceNewPlanThresholdTotalAmountPriceCadenceAnnual     PlanExternalPlanIDVersionNewParamsAddPricesPriceNewPlanThresholdTotalAmountPriceCadence = "annual"
	PlanExternalPlanIDVersionNewParamsAddPricesPriceNewPlanThresholdTotalAmountPriceCadenceSemiAnnual PlanExternalPlanIDVersionNewParamsAddPricesPriceNewPlanThresholdTotalAmountPriceCadence = "semi_annual"
	PlanExternalPlanIDVersionNewParamsAddPricesPriceNewPlanThresholdTotalAmountPriceCadenceMonthly    PlanExternalPlanIDVersionNewParamsAddPricesPriceNewPlanThresholdTotalAmountPriceCadence = "monthly"
	PlanExternalPlanIDVersionNewParamsAddPricesPriceNewPlanThresholdTotalAmountPriceCadenceQuarterly  PlanExternalPlanIDVersionNewParamsAddPricesPriceNewPlanThresholdTotalAmountPriceCadence = "quarterly"
	PlanExternalPlanIDVersionNewParamsAddPricesPriceNewPlanThresholdTotalAmountPriceCadenceOneTime    PlanExternalPlanIDVersionNewParamsAddPricesPriceNewPlanThresholdTotalAmountPriceCadence = "one_time"
	PlanExternalPlanIDVersionNewParamsAddPricesPriceNewPlanThresholdTotalAmountPriceCadenceCustom     PlanExternalPlanIDVersionNewParamsAddPricesPriceNewPlanThresholdTotalAmountPriceCadence = "custom"
)

func (r PlanExternalPlanIDVersionNewParamsAddPricesPriceNewPlanThresholdTotalAmountPriceCadence) IsKnown() bool {
	switch r {
	case PlanExternalPlanIDVersionNewParamsAddPricesPriceNewPlanThresholdTotalAmountPriceCadenceAnnual, PlanExternalPlanIDVersionNewParamsAddPricesPriceNewPlanThresholdTotalAmountPriceCadenceSemiAnnual, PlanExternalPlanIDVersionNewParamsAddPricesPriceNewPlanThresholdTotalAmountPriceCadenceMonthly, PlanExternalPlanIDVersionNewParamsAddPricesPriceNewPlanThresholdTotalAmountPriceCadenceQuarterly, PlanExternalPlanIDVersionNewParamsAddPricesPriceNewPlanThresholdTotalAmountPriceCadenceOneTime, PlanExternalPlanIDVersionNewParamsAddPricesPriceNewPlanThresholdTotalAmountPriceCadenceCustom:
		return true
	}
	return false
}

type PlanExternalPlanIDVersionNewParamsAddPricesPriceNewPlanThresholdTotalAmountPriceModelType string

const (
	PlanExternalPlanIDVersionNewParamsAddPricesPriceNewPlanThresholdTotalAmountPriceModelTypeThresholdTotalAmount PlanExternalPlanIDVersionNewParamsAddPricesPriceNewPlanThresholdTotalAmountPriceModelType = "threshold_total_amount"
)

func (r PlanExternalPlanIDVersionNewParamsAddPricesPriceNewPlanThresholdTotalAmountPriceModelType) IsKnown() bool {
	switch r {
	case PlanExternalPlanIDVersionNewParamsAddPricesPriceNewPlanThresholdTotalAmountPriceModelTypeThresholdTotalAmount:
		return true
	}
	return false
}

// For custom cadence: specifies the duration of the billing period in days or
// months.
type PlanExternalPlanIDVersionNewParamsAddPricesPriceNewPlanThresholdTotalAmountPriceBillingCycleConfiguration struct {
	// The duration of the billing period.
	Duration param.Field[int64] `json:"duration,required"`
	// The unit of billing period duration.
	DurationUnit param.Field[PlanExternalPlanIDVersionNewParamsAddPricesPriceNewPlanThresholdTotalAmountPriceBillingCycleConfigurationDurationUnit] `json:"duration_unit,required"`
}

func (r PlanExternalPlanIDVersionNewParamsAddPricesPriceNewPlanThresholdTotalAmountPriceBillingCycleConfiguration) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// The unit of billing period duration.
type PlanExternalPlanIDVersionNewParamsAddPricesPriceNewPlanThresholdTotalAmountPriceBillingCycleConfigurationDurationUnit string

const (
	PlanExternalPlanIDVersionNewParamsAddPricesPriceNewPlanThresholdTotalAmountPriceBillingCycleConfigurationDurationUnitDay   PlanExternalPlanIDVersionNewParamsAddPricesPriceNewPlanThresholdTotalAmountPriceBillingCycleConfigurationDurationUnit = "day"
	PlanExternalPlanIDVersionNewParamsAddPricesPriceNewPlanThresholdTotalAmountPriceBillingCycleConfigurationDurationUnitMonth PlanExternalPlanIDVersionNewParamsAddPricesPriceNewPlanThresholdTotalAmountPriceBillingCycleConfigurationDurationUnit = "month"
)

func (r PlanExternalPlanIDVersionNewParamsAddPricesPriceNewPlanThresholdTotalAmountPriceBillingCycleConfigurationDurationUnit) IsKnown() bool {
	switch r {
	case PlanExternalPlanIDVersionNewParamsAddPricesPriceNewPlanThresholdTotalAmountPriceBillingCycleConfigurationDurationUnitDay, PlanExternalPlanIDVersionNewParamsAddPricesPriceNewPlanThresholdTotalAmountPriceBillingCycleConfigurationDurationUnitMonth:
		return true
	}
	return false
}

// For dimensional price: specifies a price group and dimension values
type PlanExternalPlanIDVersionNewParamsAddPricesPriceNewPlanThresholdTotalAmountPriceDimensionalPriceConfiguration struct {
	// The list of dimension values matching (in order) the dimensions of the price
	// group
	DimensionValues param.Field[[]string] `json:"dimension_values,required"`
	// The id of the dimensional price group to include this price in
	DimensionalPriceGroupID param.Field[string] `json:"dimensional_price_group_id"`
	// The external id of the dimensional price group to include this price in
	ExternalDimensionalPriceGroupID param.Field[string] `json:"external_dimensional_price_group_id"`
}

func (r PlanExternalPlanIDVersionNewParamsAddPricesPriceNewPlanThresholdTotalAmountPriceDimensionalPriceConfiguration) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Within each billing cycle, specifies the cadence at which invoices are produced.
// If unspecified, a single invoice is produced per billing cycle.
type PlanExternalPlanIDVersionNewParamsAddPricesPriceNewPlanThresholdTotalAmountPriceInvoicingCycleConfiguration struct {
	// The duration of the billing period.
	Duration param.Field[int64] `json:"duration,required"`
	// The unit of billing period duration.
	DurationUnit param.Field[PlanExternalPlanIDVersionNewParamsAddPricesPriceNewPlanThresholdTotalAmountPriceInvoicingCycleConfigurationDurationUnit] `json:"duration_unit,required"`
}

func (r PlanExternalPlanIDVersionNewParamsAddPricesPriceNewPlanThresholdTotalAmountPriceInvoicingCycleConfiguration) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// The unit of billing period duration.
type PlanExternalPlanIDVersionNewParamsAddPricesPriceNewPlanThresholdTotalAmountPriceInvoicingCycleConfigurationDurationUnit string

const (
	PlanExternalPlanIDVersionNewParamsAddPricesPriceNewPlanThresholdTotalAmountPriceInvoicingCycleConfigurationDurationUnitDay   PlanExternalPlanIDVersionNewParamsAddPricesPriceNewPlanThresholdTotalAmountPriceInvoicingCycleConfigurationDurationUnit = "day"
	PlanExternalPlanIDVersionNewParamsAddPricesPriceNewPlanThresholdTotalAmountPriceInvoicingCycleConfigurationDurationUnitMonth PlanExternalPlanIDVersionNewParamsAddPricesPriceNewPlanThresholdTotalAmountPriceInvoicingCycleConfigurationDurationUnit = "month"
)

func (r PlanExternalPlanIDVersionNewParamsAddPricesPriceNewPlanThresholdTotalAmountPriceInvoicingCycleConfigurationDurationUnit) IsKnown() bool {
	switch r {
	case PlanExternalPlanIDVersionNewParamsAddPricesPriceNewPlanThresholdTotalAmountPriceInvoicingCycleConfigurationDurationUnitDay, PlanExternalPlanIDVersionNewParamsAddPricesPriceNewPlanThresholdTotalAmountPriceInvoicingCycleConfigurationDurationUnitMonth:
		return true
	}
	return false
}

type PlanExternalPlanIDVersionNewParamsAddPricesPriceNewPlanTieredPackagePrice struct {
	// The cadence to bill for this price on.
	Cadence param.Field[PlanExternalPlanIDVersionNewParamsAddPricesPriceNewPlanTieredPackagePriceCadence] `json:"cadence,required"`
	// The id of the item the price will be associated with.
	ItemID    param.Field[string]                                                                             `json:"item_id,required"`
	ModelType param.Field[PlanExternalPlanIDVersionNewParamsAddPricesPriceNewPlanTieredPackagePriceModelType] `json:"model_type,required"`
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
	BillingCycleConfiguration param.Field[PlanExternalPlanIDVersionNewParamsAddPricesPriceNewPlanTieredPackagePriceBillingCycleConfiguration] `json:"billing_cycle_configuration"`
	// The per unit conversion rate of the price currency to the invoicing currency.
	ConversionRate param.Field[float64] `json:"conversion_rate"`
	// An ISO 4217 currency string, or custom pricing unit identifier, in which this
	// price is billed.
	Currency param.Field[string] `json:"currency"`
	// For dimensional price: specifies a price group and dimension values
	DimensionalPriceConfiguration param.Field[PlanExternalPlanIDVersionNewParamsAddPricesPriceNewPlanTieredPackagePriceDimensionalPriceConfiguration] `json:"dimensional_price_configuration"`
	// An alias for the price.
	ExternalPriceID param.Field[string] `json:"external_price_id"`
	// If the Price represents a fixed cost, this represents the quantity of units
	// applied.
	FixedPriceQuantity param.Field[float64] `json:"fixed_price_quantity"`
	// The property used to group this price on an invoice
	InvoiceGroupingKey param.Field[string] `json:"invoice_grouping_key"`
	// Within each billing cycle, specifies the cadence at which invoices are produced.
	// If unspecified, a single invoice is produced per billing cycle.
	InvoicingCycleConfiguration param.Field[PlanExternalPlanIDVersionNewParamsAddPricesPriceNewPlanTieredPackagePriceInvoicingCycleConfiguration] `json:"invoicing_cycle_configuration"`
	// User-specified key/value pairs for the resource. Individual keys can be removed
	// by setting the value to `null`, and the entire metadata mapping can be cleared
	// by setting `metadata` to `null`.
	Metadata param.Field[map[string]string] `json:"metadata"`
}

func (r PlanExternalPlanIDVersionNewParamsAddPricesPriceNewPlanTieredPackagePrice) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r PlanExternalPlanIDVersionNewParamsAddPricesPriceNewPlanTieredPackagePrice) implementsPlanExternalPlanIDVersionNewParamsAddPricesPriceUnion() {
}

// The cadence to bill for this price on.
type PlanExternalPlanIDVersionNewParamsAddPricesPriceNewPlanTieredPackagePriceCadence string

const (
	PlanExternalPlanIDVersionNewParamsAddPricesPriceNewPlanTieredPackagePriceCadenceAnnual     PlanExternalPlanIDVersionNewParamsAddPricesPriceNewPlanTieredPackagePriceCadence = "annual"
	PlanExternalPlanIDVersionNewParamsAddPricesPriceNewPlanTieredPackagePriceCadenceSemiAnnual PlanExternalPlanIDVersionNewParamsAddPricesPriceNewPlanTieredPackagePriceCadence = "semi_annual"
	PlanExternalPlanIDVersionNewParamsAddPricesPriceNewPlanTieredPackagePriceCadenceMonthly    PlanExternalPlanIDVersionNewParamsAddPricesPriceNewPlanTieredPackagePriceCadence = "monthly"
	PlanExternalPlanIDVersionNewParamsAddPricesPriceNewPlanTieredPackagePriceCadenceQuarterly  PlanExternalPlanIDVersionNewParamsAddPricesPriceNewPlanTieredPackagePriceCadence = "quarterly"
	PlanExternalPlanIDVersionNewParamsAddPricesPriceNewPlanTieredPackagePriceCadenceOneTime    PlanExternalPlanIDVersionNewParamsAddPricesPriceNewPlanTieredPackagePriceCadence = "one_time"
	PlanExternalPlanIDVersionNewParamsAddPricesPriceNewPlanTieredPackagePriceCadenceCustom     PlanExternalPlanIDVersionNewParamsAddPricesPriceNewPlanTieredPackagePriceCadence = "custom"
)

func (r PlanExternalPlanIDVersionNewParamsAddPricesPriceNewPlanTieredPackagePriceCadence) IsKnown() bool {
	switch r {
	case PlanExternalPlanIDVersionNewParamsAddPricesPriceNewPlanTieredPackagePriceCadenceAnnual, PlanExternalPlanIDVersionNewParamsAddPricesPriceNewPlanTieredPackagePriceCadenceSemiAnnual, PlanExternalPlanIDVersionNewParamsAddPricesPriceNewPlanTieredPackagePriceCadenceMonthly, PlanExternalPlanIDVersionNewParamsAddPricesPriceNewPlanTieredPackagePriceCadenceQuarterly, PlanExternalPlanIDVersionNewParamsAddPricesPriceNewPlanTieredPackagePriceCadenceOneTime, PlanExternalPlanIDVersionNewParamsAddPricesPriceNewPlanTieredPackagePriceCadenceCustom:
		return true
	}
	return false
}

type PlanExternalPlanIDVersionNewParamsAddPricesPriceNewPlanTieredPackagePriceModelType string

const (
	PlanExternalPlanIDVersionNewParamsAddPricesPriceNewPlanTieredPackagePriceModelTypeTieredPackage PlanExternalPlanIDVersionNewParamsAddPricesPriceNewPlanTieredPackagePriceModelType = "tiered_package"
)

func (r PlanExternalPlanIDVersionNewParamsAddPricesPriceNewPlanTieredPackagePriceModelType) IsKnown() bool {
	switch r {
	case PlanExternalPlanIDVersionNewParamsAddPricesPriceNewPlanTieredPackagePriceModelTypeTieredPackage:
		return true
	}
	return false
}

// For custom cadence: specifies the duration of the billing period in days or
// months.
type PlanExternalPlanIDVersionNewParamsAddPricesPriceNewPlanTieredPackagePriceBillingCycleConfiguration struct {
	// The duration of the billing period.
	Duration param.Field[int64] `json:"duration,required"`
	// The unit of billing period duration.
	DurationUnit param.Field[PlanExternalPlanIDVersionNewParamsAddPricesPriceNewPlanTieredPackagePriceBillingCycleConfigurationDurationUnit] `json:"duration_unit,required"`
}

func (r PlanExternalPlanIDVersionNewParamsAddPricesPriceNewPlanTieredPackagePriceBillingCycleConfiguration) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// The unit of billing period duration.
type PlanExternalPlanIDVersionNewParamsAddPricesPriceNewPlanTieredPackagePriceBillingCycleConfigurationDurationUnit string

const (
	PlanExternalPlanIDVersionNewParamsAddPricesPriceNewPlanTieredPackagePriceBillingCycleConfigurationDurationUnitDay   PlanExternalPlanIDVersionNewParamsAddPricesPriceNewPlanTieredPackagePriceBillingCycleConfigurationDurationUnit = "day"
	PlanExternalPlanIDVersionNewParamsAddPricesPriceNewPlanTieredPackagePriceBillingCycleConfigurationDurationUnitMonth PlanExternalPlanIDVersionNewParamsAddPricesPriceNewPlanTieredPackagePriceBillingCycleConfigurationDurationUnit = "month"
)

func (r PlanExternalPlanIDVersionNewParamsAddPricesPriceNewPlanTieredPackagePriceBillingCycleConfigurationDurationUnit) IsKnown() bool {
	switch r {
	case PlanExternalPlanIDVersionNewParamsAddPricesPriceNewPlanTieredPackagePriceBillingCycleConfigurationDurationUnitDay, PlanExternalPlanIDVersionNewParamsAddPricesPriceNewPlanTieredPackagePriceBillingCycleConfigurationDurationUnitMonth:
		return true
	}
	return false
}

// For dimensional price: specifies a price group and dimension values
type PlanExternalPlanIDVersionNewParamsAddPricesPriceNewPlanTieredPackagePriceDimensionalPriceConfiguration struct {
	// The list of dimension values matching (in order) the dimensions of the price
	// group
	DimensionValues param.Field[[]string] `json:"dimension_values,required"`
	// The id of the dimensional price group to include this price in
	DimensionalPriceGroupID param.Field[string] `json:"dimensional_price_group_id"`
	// The external id of the dimensional price group to include this price in
	ExternalDimensionalPriceGroupID param.Field[string] `json:"external_dimensional_price_group_id"`
}

func (r PlanExternalPlanIDVersionNewParamsAddPricesPriceNewPlanTieredPackagePriceDimensionalPriceConfiguration) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Within each billing cycle, specifies the cadence at which invoices are produced.
// If unspecified, a single invoice is produced per billing cycle.
type PlanExternalPlanIDVersionNewParamsAddPricesPriceNewPlanTieredPackagePriceInvoicingCycleConfiguration struct {
	// The duration of the billing period.
	Duration param.Field[int64] `json:"duration,required"`
	// The unit of billing period duration.
	DurationUnit param.Field[PlanExternalPlanIDVersionNewParamsAddPricesPriceNewPlanTieredPackagePriceInvoicingCycleConfigurationDurationUnit] `json:"duration_unit,required"`
}

func (r PlanExternalPlanIDVersionNewParamsAddPricesPriceNewPlanTieredPackagePriceInvoicingCycleConfiguration) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// The unit of billing period duration.
type PlanExternalPlanIDVersionNewParamsAddPricesPriceNewPlanTieredPackagePriceInvoicingCycleConfigurationDurationUnit string

const (
	PlanExternalPlanIDVersionNewParamsAddPricesPriceNewPlanTieredPackagePriceInvoicingCycleConfigurationDurationUnitDay   PlanExternalPlanIDVersionNewParamsAddPricesPriceNewPlanTieredPackagePriceInvoicingCycleConfigurationDurationUnit = "day"
	PlanExternalPlanIDVersionNewParamsAddPricesPriceNewPlanTieredPackagePriceInvoicingCycleConfigurationDurationUnitMonth PlanExternalPlanIDVersionNewParamsAddPricesPriceNewPlanTieredPackagePriceInvoicingCycleConfigurationDurationUnit = "month"
)

func (r PlanExternalPlanIDVersionNewParamsAddPricesPriceNewPlanTieredPackagePriceInvoicingCycleConfigurationDurationUnit) IsKnown() bool {
	switch r {
	case PlanExternalPlanIDVersionNewParamsAddPricesPriceNewPlanTieredPackagePriceInvoicingCycleConfigurationDurationUnitDay, PlanExternalPlanIDVersionNewParamsAddPricesPriceNewPlanTieredPackagePriceInvoicingCycleConfigurationDurationUnitMonth:
		return true
	}
	return false
}

type PlanExternalPlanIDVersionNewParamsAddPricesPriceNewPlanTieredWithMinimumPrice struct {
	// The cadence to bill for this price on.
	Cadence param.Field[PlanExternalPlanIDVersionNewParamsAddPricesPriceNewPlanTieredWithMinimumPriceCadence] `json:"cadence,required"`
	// The id of the item the price will be associated with.
	ItemID    param.Field[string]                                                                                 `json:"item_id,required"`
	ModelType param.Field[PlanExternalPlanIDVersionNewParamsAddPricesPriceNewPlanTieredWithMinimumPriceModelType] `json:"model_type,required"`
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
	BillingCycleConfiguration param.Field[PlanExternalPlanIDVersionNewParamsAddPricesPriceNewPlanTieredWithMinimumPriceBillingCycleConfiguration] `json:"billing_cycle_configuration"`
	// The per unit conversion rate of the price currency to the invoicing currency.
	ConversionRate param.Field[float64] `json:"conversion_rate"`
	// An ISO 4217 currency string, or custom pricing unit identifier, in which this
	// price is billed.
	Currency param.Field[string] `json:"currency"`
	// For dimensional price: specifies a price group and dimension values
	DimensionalPriceConfiguration param.Field[PlanExternalPlanIDVersionNewParamsAddPricesPriceNewPlanTieredWithMinimumPriceDimensionalPriceConfiguration] `json:"dimensional_price_configuration"`
	// An alias for the price.
	ExternalPriceID param.Field[string] `json:"external_price_id"`
	// If the Price represents a fixed cost, this represents the quantity of units
	// applied.
	FixedPriceQuantity param.Field[float64] `json:"fixed_price_quantity"`
	// The property used to group this price on an invoice
	InvoiceGroupingKey param.Field[string] `json:"invoice_grouping_key"`
	// Within each billing cycle, specifies the cadence at which invoices are produced.
	// If unspecified, a single invoice is produced per billing cycle.
	InvoicingCycleConfiguration param.Field[PlanExternalPlanIDVersionNewParamsAddPricesPriceNewPlanTieredWithMinimumPriceInvoicingCycleConfiguration] `json:"invoicing_cycle_configuration"`
	// User-specified key/value pairs for the resource. Individual keys can be removed
	// by setting the value to `null`, and the entire metadata mapping can be cleared
	// by setting `metadata` to `null`.
	Metadata param.Field[map[string]string] `json:"metadata"`
}

func (r PlanExternalPlanIDVersionNewParamsAddPricesPriceNewPlanTieredWithMinimumPrice) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r PlanExternalPlanIDVersionNewParamsAddPricesPriceNewPlanTieredWithMinimumPrice) implementsPlanExternalPlanIDVersionNewParamsAddPricesPriceUnion() {
}

// The cadence to bill for this price on.
type PlanExternalPlanIDVersionNewParamsAddPricesPriceNewPlanTieredWithMinimumPriceCadence string

const (
	PlanExternalPlanIDVersionNewParamsAddPricesPriceNewPlanTieredWithMinimumPriceCadenceAnnual     PlanExternalPlanIDVersionNewParamsAddPricesPriceNewPlanTieredWithMinimumPriceCadence = "annual"
	PlanExternalPlanIDVersionNewParamsAddPricesPriceNewPlanTieredWithMinimumPriceCadenceSemiAnnual PlanExternalPlanIDVersionNewParamsAddPricesPriceNewPlanTieredWithMinimumPriceCadence = "semi_annual"
	PlanExternalPlanIDVersionNewParamsAddPricesPriceNewPlanTieredWithMinimumPriceCadenceMonthly    PlanExternalPlanIDVersionNewParamsAddPricesPriceNewPlanTieredWithMinimumPriceCadence = "monthly"
	PlanExternalPlanIDVersionNewParamsAddPricesPriceNewPlanTieredWithMinimumPriceCadenceQuarterly  PlanExternalPlanIDVersionNewParamsAddPricesPriceNewPlanTieredWithMinimumPriceCadence = "quarterly"
	PlanExternalPlanIDVersionNewParamsAddPricesPriceNewPlanTieredWithMinimumPriceCadenceOneTime    PlanExternalPlanIDVersionNewParamsAddPricesPriceNewPlanTieredWithMinimumPriceCadence = "one_time"
	PlanExternalPlanIDVersionNewParamsAddPricesPriceNewPlanTieredWithMinimumPriceCadenceCustom     PlanExternalPlanIDVersionNewParamsAddPricesPriceNewPlanTieredWithMinimumPriceCadence = "custom"
)

func (r PlanExternalPlanIDVersionNewParamsAddPricesPriceNewPlanTieredWithMinimumPriceCadence) IsKnown() bool {
	switch r {
	case PlanExternalPlanIDVersionNewParamsAddPricesPriceNewPlanTieredWithMinimumPriceCadenceAnnual, PlanExternalPlanIDVersionNewParamsAddPricesPriceNewPlanTieredWithMinimumPriceCadenceSemiAnnual, PlanExternalPlanIDVersionNewParamsAddPricesPriceNewPlanTieredWithMinimumPriceCadenceMonthly, PlanExternalPlanIDVersionNewParamsAddPricesPriceNewPlanTieredWithMinimumPriceCadenceQuarterly, PlanExternalPlanIDVersionNewParamsAddPricesPriceNewPlanTieredWithMinimumPriceCadenceOneTime, PlanExternalPlanIDVersionNewParamsAddPricesPriceNewPlanTieredWithMinimumPriceCadenceCustom:
		return true
	}
	return false
}

type PlanExternalPlanIDVersionNewParamsAddPricesPriceNewPlanTieredWithMinimumPriceModelType string

const (
	PlanExternalPlanIDVersionNewParamsAddPricesPriceNewPlanTieredWithMinimumPriceModelTypeTieredWithMinimum PlanExternalPlanIDVersionNewParamsAddPricesPriceNewPlanTieredWithMinimumPriceModelType = "tiered_with_minimum"
)

func (r PlanExternalPlanIDVersionNewParamsAddPricesPriceNewPlanTieredWithMinimumPriceModelType) IsKnown() bool {
	switch r {
	case PlanExternalPlanIDVersionNewParamsAddPricesPriceNewPlanTieredWithMinimumPriceModelTypeTieredWithMinimum:
		return true
	}
	return false
}

// For custom cadence: specifies the duration of the billing period in days or
// months.
type PlanExternalPlanIDVersionNewParamsAddPricesPriceNewPlanTieredWithMinimumPriceBillingCycleConfiguration struct {
	// The duration of the billing period.
	Duration param.Field[int64] `json:"duration,required"`
	// The unit of billing period duration.
	DurationUnit param.Field[PlanExternalPlanIDVersionNewParamsAddPricesPriceNewPlanTieredWithMinimumPriceBillingCycleConfigurationDurationUnit] `json:"duration_unit,required"`
}

func (r PlanExternalPlanIDVersionNewParamsAddPricesPriceNewPlanTieredWithMinimumPriceBillingCycleConfiguration) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// The unit of billing period duration.
type PlanExternalPlanIDVersionNewParamsAddPricesPriceNewPlanTieredWithMinimumPriceBillingCycleConfigurationDurationUnit string

const (
	PlanExternalPlanIDVersionNewParamsAddPricesPriceNewPlanTieredWithMinimumPriceBillingCycleConfigurationDurationUnitDay   PlanExternalPlanIDVersionNewParamsAddPricesPriceNewPlanTieredWithMinimumPriceBillingCycleConfigurationDurationUnit = "day"
	PlanExternalPlanIDVersionNewParamsAddPricesPriceNewPlanTieredWithMinimumPriceBillingCycleConfigurationDurationUnitMonth PlanExternalPlanIDVersionNewParamsAddPricesPriceNewPlanTieredWithMinimumPriceBillingCycleConfigurationDurationUnit = "month"
)

func (r PlanExternalPlanIDVersionNewParamsAddPricesPriceNewPlanTieredWithMinimumPriceBillingCycleConfigurationDurationUnit) IsKnown() bool {
	switch r {
	case PlanExternalPlanIDVersionNewParamsAddPricesPriceNewPlanTieredWithMinimumPriceBillingCycleConfigurationDurationUnitDay, PlanExternalPlanIDVersionNewParamsAddPricesPriceNewPlanTieredWithMinimumPriceBillingCycleConfigurationDurationUnitMonth:
		return true
	}
	return false
}

// For dimensional price: specifies a price group and dimension values
type PlanExternalPlanIDVersionNewParamsAddPricesPriceNewPlanTieredWithMinimumPriceDimensionalPriceConfiguration struct {
	// The list of dimension values matching (in order) the dimensions of the price
	// group
	DimensionValues param.Field[[]string] `json:"dimension_values,required"`
	// The id of the dimensional price group to include this price in
	DimensionalPriceGroupID param.Field[string] `json:"dimensional_price_group_id"`
	// The external id of the dimensional price group to include this price in
	ExternalDimensionalPriceGroupID param.Field[string] `json:"external_dimensional_price_group_id"`
}

func (r PlanExternalPlanIDVersionNewParamsAddPricesPriceNewPlanTieredWithMinimumPriceDimensionalPriceConfiguration) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Within each billing cycle, specifies the cadence at which invoices are produced.
// If unspecified, a single invoice is produced per billing cycle.
type PlanExternalPlanIDVersionNewParamsAddPricesPriceNewPlanTieredWithMinimumPriceInvoicingCycleConfiguration struct {
	// The duration of the billing period.
	Duration param.Field[int64] `json:"duration,required"`
	// The unit of billing period duration.
	DurationUnit param.Field[PlanExternalPlanIDVersionNewParamsAddPricesPriceNewPlanTieredWithMinimumPriceInvoicingCycleConfigurationDurationUnit] `json:"duration_unit,required"`
}

func (r PlanExternalPlanIDVersionNewParamsAddPricesPriceNewPlanTieredWithMinimumPriceInvoicingCycleConfiguration) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// The unit of billing period duration.
type PlanExternalPlanIDVersionNewParamsAddPricesPriceNewPlanTieredWithMinimumPriceInvoicingCycleConfigurationDurationUnit string

const (
	PlanExternalPlanIDVersionNewParamsAddPricesPriceNewPlanTieredWithMinimumPriceInvoicingCycleConfigurationDurationUnitDay   PlanExternalPlanIDVersionNewParamsAddPricesPriceNewPlanTieredWithMinimumPriceInvoicingCycleConfigurationDurationUnit = "day"
	PlanExternalPlanIDVersionNewParamsAddPricesPriceNewPlanTieredWithMinimumPriceInvoicingCycleConfigurationDurationUnitMonth PlanExternalPlanIDVersionNewParamsAddPricesPriceNewPlanTieredWithMinimumPriceInvoicingCycleConfigurationDurationUnit = "month"
)

func (r PlanExternalPlanIDVersionNewParamsAddPricesPriceNewPlanTieredWithMinimumPriceInvoicingCycleConfigurationDurationUnit) IsKnown() bool {
	switch r {
	case PlanExternalPlanIDVersionNewParamsAddPricesPriceNewPlanTieredWithMinimumPriceInvoicingCycleConfigurationDurationUnitDay, PlanExternalPlanIDVersionNewParamsAddPricesPriceNewPlanTieredWithMinimumPriceInvoicingCycleConfigurationDurationUnitMonth:
		return true
	}
	return false
}

type PlanExternalPlanIDVersionNewParamsAddPricesPriceNewPlanUnitWithPercentPrice struct {
	// The cadence to bill for this price on.
	Cadence param.Field[PlanExternalPlanIDVersionNewParamsAddPricesPriceNewPlanUnitWithPercentPriceCadence] `json:"cadence,required"`
	// The id of the item the price will be associated with.
	ItemID    param.Field[string]                                                                               `json:"item_id,required"`
	ModelType param.Field[PlanExternalPlanIDVersionNewParamsAddPricesPriceNewPlanUnitWithPercentPriceModelType] `json:"model_type,required"`
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
	BillingCycleConfiguration param.Field[PlanExternalPlanIDVersionNewParamsAddPricesPriceNewPlanUnitWithPercentPriceBillingCycleConfiguration] `json:"billing_cycle_configuration"`
	// The per unit conversion rate of the price currency to the invoicing currency.
	ConversionRate param.Field[float64] `json:"conversion_rate"`
	// An ISO 4217 currency string, or custom pricing unit identifier, in which this
	// price is billed.
	Currency param.Field[string] `json:"currency"`
	// For dimensional price: specifies a price group and dimension values
	DimensionalPriceConfiguration param.Field[PlanExternalPlanIDVersionNewParamsAddPricesPriceNewPlanUnitWithPercentPriceDimensionalPriceConfiguration] `json:"dimensional_price_configuration"`
	// An alias for the price.
	ExternalPriceID param.Field[string] `json:"external_price_id"`
	// If the Price represents a fixed cost, this represents the quantity of units
	// applied.
	FixedPriceQuantity param.Field[float64] `json:"fixed_price_quantity"`
	// The property used to group this price on an invoice
	InvoiceGroupingKey param.Field[string] `json:"invoice_grouping_key"`
	// Within each billing cycle, specifies the cadence at which invoices are produced.
	// If unspecified, a single invoice is produced per billing cycle.
	InvoicingCycleConfiguration param.Field[PlanExternalPlanIDVersionNewParamsAddPricesPriceNewPlanUnitWithPercentPriceInvoicingCycleConfiguration] `json:"invoicing_cycle_configuration"`
	// User-specified key/value pairs for the resource. Individual keys can be removed
	// by setting the value to `null`, and the entire metadata mapping can be cleared
	// by setting `metadata` to `null`.
	Metadata param.Field[map[string]string] `json:"metadata"`
}

func (r PlanExternalPlanIDVersionNewParamsAddPricesPriceNewPlanUnitWithPercentPrice) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r PlanExternalPlanIDVersionNewParamsAddPricesPriceNewPlanUnitWithPercentPrice) implementsPlanExternalPlanIDVersionNewParamsAddPricesPriceUnion() {
}

// The cadence to bill for this price on.
type PlanExternalPlanIDVersionNewParamsAddPricesPriceNewPlanUnitWithPercentPriceCadence string

const (
	PlanExternalPlanIDVersionNewParamsAddPricesPriceNewPlanUnitWithPercentPriceCadenceAnnual     PlanExternalPlanIDVersionNewParamsAddPricesPriceNewPlanUnitWithPercentPriceCadence = "annual"
	PlanExternalPlanIDVersionNewParamsAddPricesPriceNewPlanUnitWithPercentPriceCadenceSemiAnnual PlanExternalPlanIDVersionNewParamsAddPricesPriceNewPlanUnitWithPercentPriceCadence = "semi_annual"
	PlanExternalPlanIDVersionNewParamsAddPricesPriceNewPlanUnitWithPercentPriceCadenceMonthly    PlanExternalPlanIDVersionNewParamsAddPricesPriceNewPlanUnitWithPercentPriceCadence = "monthly"
	PlanExternalPlanIDVersionNewParamsAddPricesPriceNewPlanUnitWithPercentPriceCadenceQuarterly  PlanExternalPlanIDVersionNewParamsAddPricesPriceNewPlanUnitWithPercentPriceCadence = "quarterly"
	PlanExternalPlanIDVersionNewParamsAddPricesPriceNewPlanUnitWithPercentPriceCadenceOneTime    PlanExternalPlanIDVersionNewParamsAddPricesPriceNewPlanUnitWithPercentPriceCadence = "one_time"
	PlanExternalPlanIDVersionNewParamsAddPricesPriceNewPlanUnitWithPercentPriceCadenceCustom     PlanExternalPlanIDVersionNewParamsAddPricesPriceNewPlanUnitWithPercentPriceCadence = "custom"
)

func (r PlanExternalPlanIDVersionNewParamsAddPricesPriceNewPlanUnitWithPercentPriceCadence) IsKnown() bool {
	switch r {
	case PlanExternalPlanIDVersionNewParamsAddPricesPriceNewPlanUnitWithPercentPriceCadenceAnnual, PlanExternalPlanIDVersionNewParamsAddPricesPriceNewPlanUnitWithPercentPriceCadenceSemiAnnual, PlanExternalPlanIDVersionNewParamsAddPricesPriceNewPlanUnitWithPercentPriceCadenceMonthly, PlanExternalPlanIDVersionNewParamsAddPricesPriceNewPlanUnitWithPercentPriceCadenceQuarterly, PlanExternalPlanIDVersionNewParamsAddPricesPriceNewPlanUnitWithPercentPriceCadenceOneTime, PlanExternalPlanIDVersionNewParamsAddPricesPriceNewPlanUnitWithPercentPriceCadenceCustom:
		return true
	}
	return false
}

type PlanExternalPlanIDVersionNewParamsAddPricesPriceNewPlanUnitWithPercentPriceModelType string

const (
	PlanExternalPlanIDVersionNewParamsAddPricesPriceNewPlanUnitWithPercentPriceModelTypeUnitWithPercent PlanExternalPlanIDVersionNewParamsAddPricesPriceNewPlanUnitWithPercentPriceModelType = "unit_with_percent"
)

func (r PlanExternalPlanIDVersionNewParamsAddPricesPriceNewPlanUnitWithPercentPriceModelType) IsKnown() bool {
	switch r {
	case PlanExternalPlanIDVersionNewParamsAddPricesPriceNewPlanUnitWithPercentPriceModelTypeUnitWithPercent:
		return true
	}
	return false
}

// For custom cadence: specifies the duration of the billing period in days or
// months.
type PlanExternalPlanIDVersionNewParamsAddPricesPriceNewPlanUnitWithPercentPriceBillingCycleConfiguration struct {
	// The duration of the billing period.
	Duration param.Field[int64] `json:"duration,required"`
	// The unit of billing period duration.
	DurationUnit param.Field[PlanExternalPlanIDVersionNewParamsAddPricesPriceNewPlanUnitWithPercentPriceBillingCycleConfigurationDurationUnit] `json:"duration_unit,required"`
}

func (r PlanExternalPlanIDVersionNewParamsAddPricesPriceNewPlanUnitWithPercentPriceBillingCycleConfiguration) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// The unit of billing period duration.
type PlanExternalPlanIDVersionNewParamsAddPricesPriceNewPlanUnitWithPercentPriceBillingCycleConfigurationDurationUnit string

const (
	PlanExternalPlanIDVersionNewParamsAddPricesPriceNewPlanUnitWithPercentPriceBillingCycleConfigurationDurationUnitDay   PlanExternalPlanIDVersionNewParamsAddPricesPriceNewPlanUnitWithPercentPriceBillingCycleConfigurationDurationUnit = "day"
	PlanExternalPlanIDVersionNewParamsAddPricesPriceNewPlanUnitWithPercentPriceBillingCycleConfigurationDurationUnitMonth PlanExternalPlanIDVersionNewParamsAddPricesPriceNewPlanUnitWithPercentPriceBillingCycleConfigurationDurationUnit = "month"
)

func (r PlanExternalPlanIDVersionNewParamsAddPricesPriceNewPlanUnitWithPercentPriceBillingCycleConfigurationDurationUnit) IsKnown() bool {
	switch r {
	case PlanExternalPlanIDVersionNewParamsAddPricesPriceNewPlanUnitWithPercentPriceBillingCycleConfigurationDurationUnitDay, PlanExternalPlanIDVersionNewParamsAddPricesPriceNewPlanUnitWithPercentPriceBillingCycleConfigurationDurationUnitMonth:
		return true
	}
	return false
}

// For dimensional price: specifies a price group and dimension values
type PlanExternalPlanIDVersionNewParamsAddPricesPriceNewPlanUnitWithPercentPriceDimensionalPriceConfiguration struct {
	// The list of dimension values matching (in order) the dimensions of the price
	// group
	DimensionValues param.Field[[]string] `json:"dimension_values,required"`
	// The id of the dimensional price group to include this price in
	DimensionalPriceGroupID param.Field[string] `json:"dimensional_price_group_id"`
	// The external id of the dimensional price group to include this price in
	ExternalDimensionalPriceGroupID param.Field[string] `json:"external_dimensional_price_group_id"`
}

func (r PlanExternalPlanIDVersionNewParamsAddPricesPriceNewPlanUnitWithPercentPriceDimensionalPriceConfiguration) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Within each billing cycle, specifies the cadence at which invoices are produced.
// If unspecified, a single invoice is produced per billing cycle.
type PlanExternalPlanIDVersionNewParamsAddPricesPriceNewPlanUnitWithPercentPriceInvoicingCycleConfiguration struct {
	// The duration of the billing period.
	Duration param.Field[int64] `json:"duration,required"`
	// The unit of billing period duration.
	DurationUnit param.Field[PlanExternalPlanIDVersionNewParamsAddPricesPriceNewPlanUnitWithPercentPriceInvoicingCycleConfigurationDurationUnit] `json:"duration_unit,required"`
}

func (r PlanExternalPlanIDVersionNewParamsAddPricesPriceNewPlanUnitWithPercentPriceInvoicingCycleConfiguration) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// The unit of billing period duration.
type PlanExternalPlanIDVersionNewParamsAddPricesPriceNewPlanUnitWithPercentPriceInvoicingCycleConfigurationDurationUnit string

const (
	PlanExternalPlanIDVersionNewParamsAddPricesPriceNewPlanUnitWithPercentPriceInvoicingCycleConfigurationDurationUnitDay   PlanExternalPlanIDVersionNewParamsAddPricesPriceNewPlanUnitWithPercentPriceInvoicingCycleConfigurationDurationUnit = "day"
	PlanExternalPlanIDVersionNewParamsAddPricesPriceNewPlanUnitWithPercentPriceInvoicingCycleConfigurationDurationUnitMonth PlanExternalPlanIDVersionNewParamsAddPricesPriceNewPlanUnitWithPercentPriceInvoicingCycleConfigurationDurationUnit = "month"
)

func (r PlanExternalPlanIDVersionNewParamsAddPricesPriceNewPlanUnitWithPercentPriceInvoicingCycleConfigurationDurationUnit) IsKnown() bool {
	switch r {
	case PlanExternalPlanIDVersionNewParamsAddPricesPriceNewPlanUnitWithPercentPriceInvoicingCycleConfigurationDurationUnitDay, PlanExternalPlanIDVersionNewParamsAddPricesPriceNewPlanUnitWithPercentPriceInvoicingCycleConfigurationDurationUnitMonth:
		return true
	}
	return false
}

type PlanExternalPlanIDVersionNewParamsAddPricesPriceNewPlanPackageWithAllocationPrice struct {
	// The cadence to bill for this price on.
	Cadence param.Field[PlanExternalPlanIDVersionNewParamsAddPricesPriceNewPlanPackageWithAllocationPriceCadence] `json:"cadence,required"`
	// The id of the item the price will be associated with.
	ItemID    param.Field[string]                                                                                     `json:"item_id,required"`
	ModelType param.Field[PlanExternalPlanIDVersionNewParamsAddPricesPriceNewPlanPackageWithAllocationPriceModelType] `json:"model_type,required"`
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
	BillingCycleConfiguration param.Field[PlanExternalPlanIDVersionNewParamsAddPricesPriceNewPlanPackageWithAllocationPriceBillingCycleConfiguration] `json:"billing_cycle_configuration"`
	// The per unit conversion rate of the price currency to the invoicing currency.
	ConversionRate param.Field[float64] `json:"conversion_rate"`
	// An ISO 4217 currency string, or custom pricing unit identifier, in which this
	// price is billed.
	Currency param.Field[string] `json:"currency"`
	// For dimensional price: specifies a price group and dimension values
	DimensionalPriceConfiguration param.Field[PlanExternalPlanIDVersionNewParamsAddPricesPriceNewPlanPackageWithAllocationPriceDimensionalPriceConfiguration] `json:"dimensional_price_configuration"`
	// An alias for the price.
	ExternalPriceID param.Field[string] `json:"external_price_id"`
	// If the Price represents a fixed cost, this represents the quantity of units
	// applied.
	FixedPriceQuantity param.Field[float64] `json:"fixed_price_quantity"`
	// The property used to group this price on an invoice
	InvoiceGroupingKey param.Field[string] `json:"invoice_grouping_key"`
	// Within each billing cycle, specifies the cadence at which invoices are produced.
	// If unspecified, a single invoice is produced per billing cycle.
	InvoicingCycleConfiguration param.Field[PlanExternalPlanIDVersionNewParamsAddPricesPriceNewPlanPackageWithAllocationPriceInvoicingCycleConfiguration] `json:"invoicing_cycle_configuration"`
	// User-specified key/value pairs for the resource. Individual keys can be removed
	// by setting the value to `null`, and the entire metadata mapping can be cleared
	// by setting `metadata` to `null`.
	Metadata param.Field[map[string]string] `json:"metadata"`
}

func (r PlanExternalPlanIDVersionNewParamsAddPricesPriceNewPlanPackageWithAllocationPrice) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r PlanExternalPlanIDVersionNewParamsAddPricesPriceNewPlanPackageWithAllocationPrice) implementsPlanExternalPlanIDVersionNewParamsAddPricesPriceUnion() {
}

// The cadence to bill for this price on.
type PlanExternalPlanIDVersionNewParamsAddPricesPriceNewPlanPackageWithAllocationPriceCadence string

const (
	PlanExternalPlanIDVersionNewParamsAddPricesPriceNewPlanPackageWithAllocationPriceCadenceAnnual     PlanExternalPlanIDVersionNewParamsAddPricesPriceNewPlanPackageWithAllocationPriceCadence = "annual"
	PlanExternalPlanIDVersionNewParamsAddPricesPriceNewPlanPackageWithAllocationPriceCadenceSemiAnnual PlanExternalPlanIDVersionNewParamsAddPricesPriceNewPlanPackageWithAllocationPriceCadence = "semi_annual"
	PlanExternalPlanIDVersionNewParamsAddPricesPriceNewPlanPackageWithAllocationPriceCadenceMonthly    PlanExternalPlanIDVersionNewParamsAddPricesPriceNewPlanPackageWithAllocationPriceCadence = "monthly"
	PlanExternalPlanIDVersionNewParamsAddPricesPriceNewPlanPackageWithAllocationPriceCadenceQuarterly  PlanExternalPlanIDVersionNewParamsAddPricesPriceNewPlanPackageWithAllocationPriceCadence = "quarterly"
	PlanExternalPlanIDVersionNewParamsAddPricesPriceNewPlanPackageWithAllocationPriceCadenceOneTime    PlanExternalPlanIDVersionNewParamsAddPricesPriceNewPlanPackageWithAllocationPriceCadence = "one_time"
	PlanExternalPlanIDVersionNewParamsAddPricesPriceNewPlanPackageWithAllocationPriceCadenceCustom     PlanExternalPlanIDVersionNewParamsAddPricesPriceNewPlanPackageWithAllocationPriceCadence = "custom"
)

func (r PlanExternalPlanIDVersionNewParamsAddPricesPriceNewPlanPackageWithAllocationPriceCadence) IsKnown() bool {
	switch r {
	case PlanExternalPlanIDVersionNewParamsAddPricesPriceNewPlanPackageWithAllocationPriceCadenceAnnual, PlanExternalPlanIDVersionNewParamsAddPricesPriceNewPlanPackageWithAllocationPriceCadenceSemiAnnual, PlanExternalPlanIDVersionNewParamsAddPricesPriceNewPlanPackageWithAllocationPriceCadenceMonthly, PlanExternalPlanIDVersionNewParamsAddPricesPriceNewPlanPackageWithAllocationPriceCadenceQuarterly, PlanExternalPlanIDVersionNewParamsAddPricesPriceNewPlanPackageWithAllocationPriceCadenceOneTime, PlanExternalPlanIDVersionNewParamsAddPricesPriceNewPlanPackageWithAllocationPriceCadenceCustom:
		return true
	}
	return false
}

type PlanExternalPlanIDVersionNewParamsAddPricesPriceNewPlanPackageWithAllocationPriceModelType string

const (
	PlanExternalPlanIDVersionNewParamsAddPricesPriceNewPlanPackageWithAllocationPriceModelTypePackageWithAllocation PlanExternalPlanIDVersionNewParamsAddPricesPriceNewPlanPackageWithAllocationPriceModelType = "package_with_allocation"
)

func (r PlanExternalPlanIDVersionNewParamsAddPricesPriceNewPlanPackageWithAllocationPriceModelType) IsKnown() bool {
	switch r {
	case PlanExternalPlanIDVersionNewParamsAddPricesPriceNewPlanPackageWithAllocationPriceModelTypePackageWithAllocation:
		return true
	}
	return false
}

// For custom cadence: specifies the duration of the billing period in days or
// months.
type PlanExternalPlanIDVersionNewParamsAddPricesPriceNewPlanPackageWithAllocationPriceBillingCycleConfiguration struct {
	// The duration of the billing period.
	Duration param.Field[int64] `json:"duration,required"`
	// The unit of billing period duration.
	DurationUnit param.Field[PlanExternalPlanIDVersionNewParamsAddPricesPriceNewPlanPackageWithAllocationPriceBillingCycleConfigurationDurationUnit] `json:"duration_unit,required"`
}

func (r PlanExternalPlanIDVersionNewParamsAddPricesPriceNewPlanPackageWithAllocationPriceBillingCycleConfiguration) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// The unit of billing period duration.
type PlanExternalPlanIDVersionNewParamsAddPricesPriceNewPlanPackageWithAllocationPriceBillingCycleConfigurationDurationUnit string

const (
	PlanExternalPlanIDVersionNewParamsAddPricesPriceNewPlanPackageWithAllocationPriceBillingCycleConfigurationDurationUnitDay   PlanExternalPlanIDVersionNewParamsAddPricesPriceNewPlanPackageWithAllocationPriceBillingCycleConfigurationDurationUnit = "day"
	PlanExternalPlanIDVersionNewParamsAddPricesPriceNewPlanPackageWithAllocationPriceBillingCycleConfigurationDurationUnitMonth PlanExternalPlanIDVersionNewParamsAddPricesPriceNewPlanPackageWithAllocationPriceBillingCycleConfigurationDurationUnit = "month"
)

func (r PlanExternalPlanIDVersionNewParamsAddPricesPriceNewPlanPackageWithAllocationPriceBillingCycleConfigurationDurationUnit) IsKnown() bool {
	switch r {
	case PlanExternalPlanIDVersionNewParamsAddPricesPriceNewPlanPackageWithAllocationPriceBillingCycleConfigurationDurationUnitDay, PlanExternalPlanIDVersionNewParamsAddPricesPriceNewPlanPackageWithAllocationPriceBillingCycleConfigurationDurationUnitMonth:
		return true
	}
	return false
}

// For dimensional price: specifies a price group and dimension values
type PlanExternalPlanIDVersionNewParamsAddPricesPriceNewPlanPackageWithAllocationPriceDimensionalPriceConfiguration struct {
	// The list of dimension values matching (in order) the dimensions of the price
	// group
	DimensionValues param.Field[[]string] `json:"dimension_values,required"`
	// The id of the dimensional price group to include this price in
	DimensionalPriceGroupID param.Field[string] `json:"dimensional_price_group_id"`
	// The external id of the dimensional price group to include this price in
	ExternalDimensionalPriceGroupID param.Field[string] `json:"external_dimensional_price_group_id"`
}

func (r PlanExternalPlanIDVersionNewParamsAddPricesPriceNewPlanPackageWithAllocationPriceDimensionalPriceConfiguration) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Within each billing cycle, specifies the cadence at which invoices are produced.
// If unspecified, a single invoice is produced per billing cycle.
type PlanExternalPlanIDVersionNewParamsAddPricesPriceNewPlanPackageWithAllocationPriceInvoicingCycleConfiguration struct {
	// The duration of the billing period.
	Duration param.Field[int64] `json:"duration,required"`
	// The unit of billing period duration.
	DurationUnit param.Field[PlanExternalPlanIDVersionNewParamsAddPricesPriceNewPlanPackageWithAllocationPriceInvoicingCycleConfigurationDurationUnit] `json:"duration_unit,required"`
}

func (r PlanExternalPlanIDVersionNewParamsAddPricesPriceNewPlanPackageWithAllocationPriceInvoicingCycleConfiguration) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// The unit of billing period duration.
type PlanExternalPlanIDVersionNewParamsAddPricesPriceNewPlanPackageWithAllocationPriceInvoicingCycleConfigurationDurationUnit string

const (
	PlanExternalPlanIDVersionNewParamsAddPricesPriceNewPlanPackageWithAllocationPriceInvoicingCycleConfigurationDurationUnitDay   PlanExternalPlanIDVersionNewParamsAddPricesPriceNewPlanPackageWithAllocationPriceInvoicingCycleConfigurationDurationUnit = "day"
	PlanExternalPlanIDVersionNewParamsAddPricesPriceNewPlanPackageWithAllocationPriceInvoicingCycleConfigurationDurationUnitMonth PlanExternalPlanIDVersionNewParamsAddPricesPriceNewPlanPackageWithAllocationPriceInvoicingCycleConfigurationDurationUnit = "month"
)

func (r PlanExternalPlanIDVersionNewParamsAddPricesPriceNewPlanPackageWithAllocationPriceInvoicingCycleConfigurationDurationUnit) IsKnown() bool {
	switch r {
	case PlanExternalPlanIDVersionNewParamsAddPricesPriceNewPlanPackageWithAllocationPriceInvoicingCycleConfigurationDurationUnitDay, PlanExternalPlanIDVersionNewParamsAddPricesPriceNewPlanPackageWithAllocationPriceInvoicingCycleConfigurationDurationUnitMonth:
		return true
	}
	return false
}

type PlanExternalPlanIDVersionNewParamsAddPricesPriceNewPlanTierWithProrationPrice struct {
	// The cadence to bill for this price on.
	Cadence param.Field[PlanExternalPlanIDVersionNewParamsAddPricesPriceNewPlanTierWithProrationPriceCadence] `json:"cadence,required"`
	// The id of the item the price will be associated with.
	ItemID    param.Field[string]                                                                                 `json:"item_id,required"`
	ModelType param.Field[PlanExternalPlanIDVersionNewParamsAddPricesPriceNewPlanTierWithProrationPriceModelType] `json:"model_type,required"`
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
	BillingCycleConfiguration param.Field[PlanExternalPlanIDVersionNewParamsAddPricesPriceNewPlanTierWithProrationPriceBillingCycleConfiguration] `json:"billing_cycle_configuration"`
	// The per unit conversion rate of the price currency to the invoicing currency.
	ConversionRate param.Field[float64] `json:"conversion_rate"`
	// An ISO 4217 currency string, or custom pricing unit identifier, in which this
	// price is billed.
	Currency param.Field[string] `json:"currency"`
	// For dimensional price: specifies a price group and dimension values
	DimensionalPriceConfiguration param.Field[PlanExternalPlanIDVersionNewParamsAddPricesPriceNewPlanTierWithProrationPriceDimensionalPriceConfiguration] `json:"dimensional_price_configuration"`
	// An alias for the price.
	ExternalPriceID param.Field[string] `json:"external_price_id"`
	// If the Price represents a fixed cost, this represents the quantity of units
	// applied.
	FixedPriceQuantity param.Field[float64] `json:"fixed_price_quantity"`
	// The property used to group this price on an invoice
	InvoiceGroupingKey param.Field[string] `json:"invoice_grouping_key"`
	// Within each billing cycle, specifies the cadence at which invoices are produced.
	// If unspecified, a single invoice is produced per billing cycle.
	InvoicingCycleConfiguration param.Field[PlanExternalPlanIDVersionNewParamsAddPricesPriceNewPlanTierWithProrationPriceInvoicingCycleConfiguration] `json:"invoicing_cycle_configuration"`
	// User-specified key/value pairs for the resource. Individual keys can be removed
	// by setting the value to `null`, and the entire metadata mapping can be cleared
	// by setting `metadata` to `null`.
	Metadata param.Field[map[string]string] `json:"metadata"`
}

func (r PlanExternalPlanIDVersionNewParamsAddPricesPriceNewPlanTierWithProrationPrice) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r PlanExternalPlanIDVersionNewParamsAddPricesPriceNewPlanTierWithProrationPrice) implementsPlanExternalPlanIDVersionNewParamsAddPricesPriceUnion() {
}

// The cadence to bill for this price on.
type PlanExternalPlanIDVersionNewParamsAddPricesPriceNewPlanTierWithProrationPriceCadence string

const (
	PlanExternalPlanIDVersionNewParamsAddPricesPriceNewPlanTierWithProrationPriceCadenceAnnual     PlanExternalPlanIDVersionNewParamsAddPricesPriceNewPlanTierWithProrationPriceCadence = "annual"
	PlanExternalPlanIDVersionNewParamsAddPricesPriceNewPlanTierWithProrationPriceCadenceSemiAnnual PlanExternalPlanIDVersionNewParamsAddPricesPriceNewPlanTierWithProrationPriceCadence = "semi_annual"
	PlanExternalPlanIDVersionNewParamsAddPricesPriceNewPlanTierWithProrationPriceCadenceMonthly    PlanExternalPlanIDVersionNewParamsAddPricesPriceNewPlanTierWithProrationPriceCadence = "monthly"
	PlanExternalPlanIDVersionNewParamsAddPricesPriceNewPlanTierWithProrationPriceCadenceQuarterly  PlanExternalPlanIDVersionNewParamsAddPricesPriceNewPlanTierWithProrationPriceCadence = "quarterly"
	PlanExternalPlanIDVersionNewParamsAddPricesPriceNewPlanTierWithProrationPriceCadenceOneTime    PlanExternalPlanIDVersionNewParamsAddPricesPriceNewPlanTierWithProrationPriceCadence = "one_time"
	PlanExternalPlanIDVersionNewParamsAddPricesPriceNewPlanTierWithProrationPriceCadenceCustom     PlanExternalPlanIDVersionNewParamsAddPricesPriceNewPlanTierWithProrationPriceCadence = "custom"
)

func (r PlanExternalPlanIDVersionNewParamsAddPricesPriceNewPlanTierWithProrationPriceCadence) IsKnown() bool {
	switch r {
	case PlanExternalPlanIDVersionNewParamsAddPricesPriceNewPlanTierWithProrationPriceCadenceAnnual, PlanExternalPlanIDVersionNewParamsAddPricesPriceNewPlanTierWithProrationPriceCadenceSemiAnnual, PlanExternalPlanIDVersionNewParamsAddPricesPriceNewPlanTierWithProrationPriceCadenceMonthly, PlanExternalPlanIDVersionNewParamsAddPricesPriceNewPlanTierWithProrationPriceCadenceQuarterly, PlanExternalPlanIDVersionNewParamsAddPricesPriceNewPlanTierWithProrationPriceCadenceOneTime, PlanExternalPlanIDVersionNewParamsAddPricesPriceNewPlanTierWithProrationPriceCadenceCustom:
		return true
	}
	return false
}

type PlanExternalPlanIDVersionNewParamsAddPricesPriceNewPlanTierWithProrationPriceModelType string

const (
	PlanExternalPlanIDVersionNewParamsAddPricesPriceNewPlanTierWithProrationPriceModelTypeTieredWithProration PlanExternalPlanIDVersionNewParamsAddPricesPriceNewPlanTierWithProrationPriceModelType = "tiered_with_proration"
)

func (r PlanExternalPlanIDVersionNewParamsAddPricesPriceNewPlanTierWithProrationPriceModelType) IsKnown() bool {
	switch r {
	case PlanExternalPlanIDVersionNewParamsAddPricesPriceNewPlanTierWithProrationPriceModelTypeTieredWithProration:
		return true
	}
	return false
}

// For custom cadence: specifies the duration of the billing period in days or
// months.
type PlanExternalPlanIDVersionNewParamsAddPricesPriceNewPlanTierWithProrationPriceBillingCycleConfiguration struct {
	// The duration of the billing period.
	Duration param.Field[int64] `json:"duration,required"`
	// The unit of billing period duration.
	DurationUnit param.Field[PlanExternalPlanIDVersionNewParamsAddPricesPriceNewPlanTierWithProrationPriceBillingCycleConfigurationDurationUnit] `json:"duration_unit,required"`
}

func (r PlanExternalPlanIDVersionNewParamsAddPricesPriceNewPlanTierWithProrationPriceBillingCycleConfiguration) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// The unit of billing period duration.
type PlanExternalPlanIDVersionNewParamsAddPricesPriceNewPlanTierWithProrationPriceBillingCycleConfigurationDurationUnit string

const (
	PlanExternalPlanIDVersionNewParamsAddPricesPriceNewPlanTierWithProrationPriceBillingCycleConfigurationDurationUnitDay   PlanExternalPlanIDVersionNewParamsAddPricesPriceNewPlanTierWithProrationPriceBillingCycleConfigurationDurationUnit = "day"
	PlanExternalPlanIDVersionNewParamsAddPricesPriceNewPlanTierWithProrationPriceBillingCycleConfigurationDurationUnitMonth PlanExternalPlanIDVersionNewParamsAddPricesPriceNewPlanTierWithProrationPriceBillingCycleConfigurationDurationUnit = "month"
)

func (r PlanExternalPlanIDVersionNewParamsAddPricesPriceNewPlanTierWithProrationPriceBillingCycleConfigurationDurationUnit) IsKnown() bool {
	switch r {
	case PlanExternalPlanIDVersionNewParamsAddPricesPriceNewPlanTierWithProrationPriceBillingCycleConfigurationDurationUnitDay, PlanExternalPlanIDVersionNewParamsAddPricesPriceNewPlanTierWithProrationPriceBillingCycleConfigurationDurationUnitMonth:
		return true
	}
	return false
}

// For dimensional price: specifies a price group and dimension values
type PlanExternalPlanIDVersionNewParamsAddPricesPriceNewPlanTierWithProrationPriceDimensionalPriceConfiguration struct {
	// The list of dimension values matching (in order) the dimensions of the price
	// group
	DimensionValues param.Field[[]string] `json:"dimension_values,required"`
	// The id of the dimensional price group to include this price in
	DimensionalPriceGroupID param.Field[string] `json:"dimensional_price_group_id"`
	// The external id of the dimensional price group to include this price in
	ExternalDimensionalPriceGroupID param.Field[string] `json:"external_dimensional_price_group_id"`
}

func (r PlanExternalPlanIDVersionNewParamsAddPricesPriceNewPlanTierWithProrationPriceDimensionalPriceConfiguration) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Within each billing cycle, specifies the cadence at which invoices are produced.
// If unspecified, a single invoice is produced per billing cycle.
type PlanExternalPlanIDVersionNewParamsAddPricesPriceNewPlanTierWithProrationPriceInvoicingCycleConfiguration struct {
	// The duration of the billing period.
	Duration param.Field[int64] `json:"duration,required"`
	// The unit of billing period duration.
	DurationUnit param.Field[PlanExternalPlanIDVersionNewParamsAddPricesPriceNewPlanTierWithProrationPriceInvoicingCycleConfigurationDurationUnit] `json:"duration_unit,required"`
}

func (r PlanExternalPlanIDVersionNewParamsAddPricesPriceNewPlanTierWithProrationPriceInvoicingCycleConfiguration) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// The unit of billing period duration.
type PlanExternalPlanIDVersionNewParamsAddPricesPriceNewPlanTierWithProrationPriceInvoicingCycleConfigurationDurationUnit string

const (
	PlanExternalPlanIDVersionNewParamsAddPricesPriceNewPlanTierWithProrationPriceInvoicingCycleConfigurationDurationUnitDay   PlanExternalPlanIDVersionNewParamsAddPricesPriceNewPlanTierWithProrationPriceInvoicingCycleConfigurationDurationUnit = "day"
	PlanExternalPlanIDVersionNewParamsAddPricesPriceNewPlanTierWithProrationPriceInvoicingCycleConfigurationDurationUnitMonth PlanExternalPlanIDVersionNewParamsAddPricesPriceNewPlanTierWithProrationPriceInvoicingCycleConfigurationDurationUnit = "month"
)

func (r PlanExternalPlanIDVersionNewParamsAddPricesPriceNewPlanTierWithProrationPriceInvoicingCycleConfigurationDurationUnit) IsKnown() bool {
	switch r {
	case PlanExternalPlanIDVersionNewParamsAddPricesPriceNewPlanTierWithProrationPriceInvoicingCycleConfigurationDurationUnitDay, PlanExternalPlanIDVersionNewParamsAddPricesPriceNewPlanTierWithProrationPriceInvoicingCycleConfigurationDurationUnitMonth:
		return true
	}
	return false
}

type PlanExternalPlanIDVersionNewParamsAddPricesPriceNewPlanUnitWithProrationPrice struct {
	// The cadence to bill for this price on.
	Cadence param.Field[PlanExternalPlanIDVersionNewParamsAddPricesPriceNewPlanUnitWithProrationPriceCadence] `json:"cadence,required"`
	// The id of the item the price will be associated with.
	ItemID    param.Field[string]                                                                                 `json:"item_id,required"`
	ModelType param.Field[PlanExternalPlanIDVersionNewParamsAddPricesPriceNewPlanUnitWithProrationPriceModelType] `json:"model_type,required"`
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
	BillingCycleConfiguration param.Field[PlanExternalPlanIDVersionNewParamsAddPricesPriceNewPlanUnitWithProrationPriceBillingCycleConfiguration] `json:"billing_cycle_configuration"`
	// The per unit conversion rate of the price currency to the invoicing currency.
	ConversionRate param.Field[float64] `json:"conversion_rate"`
	// An ISO 4217 currency string, or custom pricing unit identifier, in which this
	// price is billed.
	Currency param.Field[string] `json:"currency"`
	// For dimensional price: specifies a price group and dimension values
	DimensionalPriceConfiguration param.Field[PlanExternalPlanIDVersionNewParamsAddPricesPriceNewPlanUnitWithProrationPriceDimensionalPriceConfiguration] `json:"dimensional_price_configuration"`
	// An alias for the price.
	ExternalPriceID param.Field[string] `json:"external_price_id"`
	// If the Price represents a fixed cost, this represents the quantity of units
	// applied.
	FixedPriceQuantity param.Field[float64] `json:"fixed_price_quantity"`
	// The property used to group this price on an invoice
	InvoiceGroupingKey param.Field[string] `json:"invoice_grouping_key"`
	// Within each billing cycle, specifies the cadence at which invoices are produced.
	// If unspecified, a single invoice is produced per billing cycle.
	InvoicingCycleConfiguration param.Field[PlanExternalPlanIDVersionNewParamsAddPricesPriceNewPlanUnitWithProrationPriceInvoicingCycleConfiguration] `json:"invoicing_cycle_configuration"`
	// User-specified key/value pairs for the resource. Individual keys can be removed
	// by setting the value to `null`, and the entire metadata mapping can be cleared
	// by setting `metadata` to `null`.
	Metadata param.Field[map[string]string] `json:"metadata"`
}

func (r PlanExternalPlanIDVersionNewParamsAddPricesPriceNewPlanUnitWithProrationPrice) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r PlanExternalPlanIDVersionNewParamsAddPricesPriceNewPlanUnitWithProrationPrice) implementsPlanExternalPlanIDVersionNewParamsAddPricesPriceUnion() {
}

// The cadence to bill for this price on.
type PlanExternalPlanIDVersionNewParamsAddPricesPriceNewPlanUnitWithProrationPriceCadence string

const (
	PlanExternalPlanIDVersionNewParamsAddPricesPriceNewPlanUnitWithProrationPriceCadenceAnnual     PlanExternalPlanIDVersionNewParamsAddPricesPriceNewPlanUnitWithProrationPriceCadence = "annual"
	PlanExternalPlanIDVersionNewParamsAddPricesPriceNewPlanUnitWithProrationPriceCadenceSemiAnnual PlanExternalPlanIDVersionNewParamsAddPricesPriceNewPlanUnitWithProrationPriceCadence = "semi_annual"
	PlanExternalPlanIDVersionNewParamsAddPricesPriceNewPlanUnitWithProrationPriceCadenceMonthly    PlanExternalPlanIDVersionNewParamsAddPricesPriceNewPlanUnitWithProrationPriceCadence = "monthly"
	PlanExternalPlanIDVersionNewParamsAddPricesPriceNewPlanUnitWithProrationPriceCadenceQuarterly  PlanExternalPlanIDVersionNewParamsAddPricesPriceNewPlanUnitWithProrationPriceCadence = "quarterly"
	PlanExternalPlanIDVersionNewParamsAddPricesPriceNewPlanUnitWithProrationPriceCadenceOneTime    PlanExternalPlanIDVersionNewParamsAddPricesPriceNewPlanUnitWithProrationPriceCadence = "one_time"
	PlanExternalPlanIDVersionNewParamsAddPricesPriceNewPlanUnitWithProrationPriceCadenceCustom     PlanExternalPlanIDVersionNewParamsAddPricesPriceNewPlanUnitWithProrationPriceCadence = "custom"
)

func (r PlanExternalPlanIDVersionNewParamsAddPricesPriceNewPlanUnitWithProrationPriceCadence) IsKnown() bool {
	switch r {
	case PlanExternalPlanIDVersionNewParamsAddPricesPriceNewPlanUnitWithProrationPriceCadenceAnnual, PlanExternalPlanIDVersionNewParamsAddPricesPriceNewPlanUnitWithProrationPriceCadenceSemiAnnual, PlanExternalPlanIDVersionNewParamsAddPricesPriceNewPlanUnitWithProrationPriceCadenceMonthly, PlanExternalPlanIDVersionNewParamsAddPricesPriceNewPlanUnitWithProrationPriceCadenceQuarterly, PlanExternalPlanIDVersionNewParamsAddPricesPriceNewPlanUnitWithProrationPriceCadenceOneTime, PlanExternalPlanIDVersionNewParamsAddPricesPriceNewPlanUnitWithProrationPriceCadenceCustom:
		return true
	}
	return false
}

type PlanExternalPlanIDVersionNewParamsAddPricesPriceNewPlanUnitWithProrationPriceModelType string

const (
	PlanExternalPlanIDVersionNewParamsAddPricesPriceNewPlanUnitWithProrationPriceModelTypeUnitWithProration PlanExternalPlanIDVersionNewParamsAddPricesPriceNewPlanUnitWithProrationPriceModelType = "unit_with_proration"
)

func (r PlanExternalPlanIDVersionNewParamsAddPricesPriceNewPlanUnitWithProrationPriceModelType) IsKnown() bool {
	switch r {
	case PlanExternalPlanIDVersionNewParamsAddPricesPriceNewPlanUnitWithProrationPriceModelTypeUnitWithProration:
		return true
	}
	return false
}

// For custom cadence: specifies the duration of the billing period in days or
// months.
type PlanExternalPlanIDVersionNewParamsAddPricesPriceNewPlanUnitWithProrationPriceBillingCycleConfiguration struct {
	// The duration of the billing period.
	Duration param.Field[int64] `json:"duration,required"`
	// The unit of billing period duration.
	DurationUnit param.Field[PlanExternalPlanIDVersionNewParamsAddPricesPriceNewPlanUnitWithProrationPriceBillingCycleConfigurationDurationUnit] `json:"duration_unit,required"`
}

func (r PlanExternalPlanIDVersionNewParamsAddPricesPriceNewPlanUnitWithProrationPriceBillingCycleConfiguration) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// The unit of billing period duration.
type PlanExternalPlanIDVersionNewParamsAddPricesPriceNewPlanUnitWithProrationPriceBillingCycleConfigurationDurationUnit string

const (
	PlanExternalPlanIDVersionNewParamsAddPricesPriceNewPlanUnitWithProrationPriceBillingCycleConfigurationDurationUnitDay   PlanExternalPlanIDVersionNewParamsAddPricesPriceNewPlanUnitWithProrationPriceBillingCycleConfigurationDurationUnit = "day"
	PlanExternalPlanIDVersionNewParamsAddPricesPriceNewPlanUnitWithProrationPriceBillingCycleConfigurationDurationUnitMonth PlanExternalPlanIDVersionNewParamsAddPricesPriceNewPlanUnitWithProrationPriceBillingCycleConfigurationDurationUnit = "month"
)

func (r PlanExternalPlanIDVersionNewParamsAddPricesPriceNewPlanUnitWithProrationPriceBillingCycleConfigurationDurationUnit) IsKnown() bool {
	switch r {
	case PlanExternalPlanIDVersionNewParamsAddPricesPriceNewPlanUnitWithProrationPriceBillingCycleConfigurationDurationUnitDay, PlanExternalPlanIDVersionNewParamsAddPricesPriceNewPlanUnitWithProrationPriceBillingCycleConfigurationDurationUnitMonth:
		return true
	}
	return false
}

// For dimensional price: specifies a price group and dimension values
type PlanExternalPlanIDVersionNewParamsAddPricesPriceNewPlanUnitWithProrationPriceDimensionalPriceConfiguration struct {
	// The list of dimension values matching (in order) the dimensions of the price
	// group
	DimensionValues param.Field[[]string] `json:"dimension_values,required"`
	// The id of the dimensional price group to include this price in
	DimensionalPriceGroupID param.Field[string] `json:"dimensional_price_group_id"`
	// The external id of the dimensional price group to include this price in
	ExternalDimensionalPriceGroupID param.Field[string] `json:"external_dimensional_price_group_id"`
}

func (r PlanExternalPlanIDVersionNewParamsAddPricesPriceNewPlanUnitWithProrationPriceDimensionalPriceConfiguration) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Within each billing cycle, specifies the cadence at which invoices are produced.
// If unspecified, a single invoice is produced per billing cycle.
type PlanExternalPlanIDVersionNewParamsAddPricesPriceNewPlanUnitWithProrationPriceInvoicingCycleConfiguration struct {
	// The duration of the billing period.
	Duration param.Field[int64] `json:"duration,required"`
	// The unit of billing period duration.
	DurationUnit param.Field[PlanExternalPlanIDVersionNewParamsAddPricesPriceNewPlanUnitWithProrationPriceInvoicingCycleConfigurationDurationUnit] `json:"duration_unit,required"`
}

func (r PlanExternalPlanIDVersionNewParamsAddPricesPriceNewPlanUnitWithProrationPriceInvoicingCycleConfiguration) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// The unit of billing period duration.
type PlanExternalPlanIDVersionNewParamsAddPricesPriceNewPlanUnitWithProrationPriceInvoicingCycleConfigurationDurationUnit string

const (
	PlanExternalPlanIDVersionNewParamsAddPricesPriceNewPlanUnitWithProrationPriceInvoicingCycleConfigurationDurationUnitDay   PlanExternalPlanIDVersionNewParamsAddPricesPriceNewPlanUnitWithProrationPriceInvoicingCycleConfigurationDurationUnit = "day"
	PlanExternalPlanIDVersionNewParamsAddPricesPriceNewPlanUnitWithProrationPriceInvoicingCycleConfigurationDurationUnitMonth PlanExternalPlanIDVersionNewParamsAddPricesPriceNewPlanUnitWithProrationPriceInvoicingCycleConfigurationDurationUnit = "month"
)

func (r PlanExternalPlanIDVersionNewParamsAddPricesPriceNewPlanUnitWithProrationPriceInvoicingCycleConfigurationDurationUnit) IsKnown() bool {
	switch r {
	case PlanExternalPlanIDVersionNewParamsAddPricesPriceNewPlanUnitWithProrationPriceInvoicingCycleConfigurationDurationUnitDay, PlanExternalPlanIDVersionNewParamsAddPricesPriceNewPlanUnitWithProrationPriceInvoicingCycleConfigurationDurationUnitMonth:
		return true
	}
	return false
}

type PlanExternalPlanIDVersionNewParamsAddPricesPriceNewPlanGroupedAllocationPrice struct {
	// The cadence to bill for this price on.
	Cadence                 param.Field[PlanExternalPlanIDVersionNewParamsAddPricesPriceNewPlanGroupedAllocationPriceCadence] `json:"cadence,required"`
	GroupedAllocationConfig param.Field[map[string]interface{}]                                                               `json:"grouped_allocation_config,required"`
	// The id of the item the price will be associated with.
	ItemID    param.Field[string]                                                                                 `json:"item_id,required"`
	ModelType param.Field[PlanExternalPlanIDVersionNewParamsAddPricesPriceNewPlanGroupedAllocationPriceModelType] `json:"model_type,required"`
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
	BillingCycleConfiguration param.Field[PlanExternalPlanIDVersionNewParamsAddPricesPriceNewPlanGroupedAllocationPriceBillingCycleConfiguration] `json:"billing_cycle_configuration"`
	// The per unit conversion rate of the price currency to the invoicing currency.
	ConversionRate param.Field[float64] `json:"conversion_rate"`
	// An ISO 4217 currency string, or custom pricing unit identifier, in which this
	// price is billed.
	Currency param.Field[string] `json:"currency"`
	// For dimensional price: specifies a price group and dimension values
	DimensionalPriceConfiguration param.Field[PlanExternalPlanIDVersionNewParamsAddPricesPriceNewPlanGroupedAllocationPriceDimensionalPriceConfiguration] `json:"dimensional_price_configuration"`
	// An alias for the price.
	ExternalPriceID param.Field[string] `json:"external_price_id"`
	// If the Price represents a fixed cost, this represents the quantity of units
	// applied.
	FixedPriceQuantity param.Field[float64] `json:"fixed_price_quantity"`
	// The property used to group this price on an invoice
	InvoiceGroupingKey param.Field[string] `json:"invoice_grouping_key"`
	// Within each billing cycle, specifies the cadence at which invoices are produced.
	// If unspecified, a single invoice is produced per billing cycle.
	InvoicingCycleConfiguration param.Field[PlanExternalPlanIDVersionNewParamsAddPricesPriceNewPlanGroupedAllocationPriceInvoicingCycleConfiguration] `json:"invoicing_cycle_configuration"`
	// User-specified key/value pairs for the resource. Individual keys can be removed
	// by setting the value to `null`, and the entire metadata mapping can be cleared
	// by setting `metadata` to `null`.
	Metadata param.Field[map[string]string] `json:"metadata"`
}

func (r PlanExternalPlanIDVersionNewParamsAddPricesPriceNewPlanGroupedAllocationPrice) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r PlanExternalPlanIDVersionNewParamsAddPricesPriceNewPlanGroupedAllocationPrice) implementsPlanExternalPlanIDVersionNewParamsAddPricesPriceUnion() {
}

// The cadence to bill for this price on.
type PlanExternalPlanIDVersionNewParamsAddPricesPriceNewPlanGroupedAllocationPriceCadence string

const (
	PlanExternalPlanIDVersionNewParamsAddPricesPriceNewPlanGroupedAllocationPriceCadenceAnnual     PlanExternalPlanIDVersionNewParamsAddPricesPriceNewPlanGroupedAllocationPriceCadence = "annual"
	PlanExternalPlanIDVersionNewParamsAddPricesPriceNewPlanGroupedAllocationPriceCadenceSemiAnnual PlanExternalPlanIDVersionNewParamsAddPricesPriceNewPlanGroupedAllocationPriceCadence = "semi_annual"
	PlanExternalPlanIDVersionNewParamsAddPricesPriceNewPlanGroupedAllocationPriceCadenceMonthly    PlanExternalPlanIDVersionNewParamsAddPricesPriceNewPlanGroupedAllocationPriceCadence = "monthly"
	PlanExternalPlanIDVersionNewParamsAddPricesPriceNewPlanGroupedAllocationPriceCadenceQuarterly  PlanExternalPlanIDVersionNewParamsAddPricesPriceNewPlanGroupedAllocationPriceCadence = "quarterly"
	PlanExternalPlanIDVersionNewParamsAddPricesPriceNewPlanGroupedAllocationPriceCadenceOneTime    PlanExternalPlanIDVersionNewParamsAddPricesPriceNewPlanGroupedAllocationPriceCadence = "one_time"
	PlanExternalPlanIDVersionNewParamsAddPricesPriceNewPlanGroupedAllocationPriceCadenceCustom     PlanExternalPlanIDVersionNewParamsAddPricesPriceNewPlanGroupedAllocationPriceCadence = "custom"
)

func (r PlanExternalPlanIDVersionNewParamsAddPricesPriceNewPlanGroupedAllocationPriceCadence) IsKnown() bool {
	switch r {
	case PlanExternalPlanIDVersionNewParamsAddPricesPriceNewPlanGroupedAllocationPriceCadenceAnnual, PlanExternalPlanIDVersionNewParamsAddPricesPriceNewPlanGroupedAllocationPriceCadenceSemiAnnual, PlanExternalPlanIDVersionNewParamsAddPricesPriceNewPlanGroupedAllocationPriceCadenceMonthly, PlanExternalPlanIDVersionNewParamsAddPricesPriceNewPlanGroupedAllocationPriceCadenceQuarterly, PlanExternalPlanIDVersionNewParamsAddPricesPriceNewPlanGroupedAllocationPriceCadenceOneTime, PlanExternalPlanIDVersionNewParamsAddPricesPriceNewPlanGroupedAllocationPriceCadenceCustom:
		return true
	}
	return false
}

type PlanExternalPlanIDVersionNewParamsAddPricesPriceNewPlanGroupedAllocationPriceModelType string

const (
	PlanExternalPlanIDVersionNewParamsAddPricesPriceNewPlanGroupedAllocationPriceModelTypeGroupedAllocation PlanExternalPlanIDVersionNewParamsAddPricesPriceNewPlanGroupedAllocationPriceModelType = "grouped_allocation"
)

func (r PlanExternalPlanIDVersionNewParamsAddPricesPriceNewPlanGroupedAllocationPriceModelType) IsKnown() bool {
	switch r {
	case PlanExternalPlanIDVersionNewParamsAddPricesPriceNewPlanGroupedAllocationPriceModelTypeGroupedAllocation:
		return true
	}
	return false
}

// For custom cadence: specifies the duration of the billing period in days or
// months.
type PlanExternalPlanIDVersionNewParamsAddPricesPriceNewPlanGroupedAllocationPriceBillingCycleConfiguration struct {
	// The duration of the billing period.
	Duration param.Field[int64] `json:"duration,required"`
	// The unit of billing period duration.
	DurationUnit param.Field[PlanExternalPlanIDVersionNewParamsAddPricesPriceNewPlanGroupedAllocationPriceBillingCycleConfigurationDurationUnit] `json:"duration_unit,required"`
}

func (r PlanExternalPlanIDVersionNewParamsAddPricesPriceNewPlanGroupedAllocationPriceBillingCycleConfiguration) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// The unit of billing period duration.
type PlanExternalPlanIDVersionNewParamsAddPricesPriceNewPlanGroupedAllocationPriceBillingCycleConfigurationDurationUnit string

const (
	PlanExternalPlanIDVersionNewParamsAddPricesPriceNewPlanGroupedAllocationPriceBillingCycleConfigurationDurationUnitDay   PlanExternalPlanIDVersionNewParamsAddPricesPriceNewPlanGroupedAllocationPriceBillingCycleConfigurationDurationUnit = "day"
	PlanExternalPlanIDVersionNewParamsAddPricesPriceNewPlanGroupedAllocationPriceBillingCycleConfigurationDurationUnitMonth PlanExternalPlanIDVersionNewParamsAddPricesPriceNewPlanGroupedAllocationPriceBillingCycleConfigurationDurationUnit = "month"
)

func (r PlanExternalPlanIDVersionNewParamsAddPricesPriceNewPlanGroupedAllocationPriceBillingCycleConfigurationDurationUnit) IsKnown() bool {
	switch r {
	case PlanExternalPlanIDVersionNewParamsAddPricesPriceNewPlanGroupedAllocationPriceBillingCycleConfigurationDurationUnitDay, PlanExternalPlanIDVersionNewParamsAddPricesPriceNewPlanGroupedAllocationPriceBillingCycleConfigurationDurationUnitMonth:
		return true
	}
	return false
}

// For dimensional price: specifies a price group and dimension values
type PlanExternalPlanIDVersionNewParamsAddPricesPriceNewPlanGroupedAllocationPriceDimensionalPriceConfiguration struct {
	// The list of dimension values matching (in order) the dimensions of the price
	// group
	DimensionValues param.Field[[]string] `json:"dimension_values,required"`
	// The id of the dimensional price group to include this price in
	DimensionalPriceGroupID param.Field[string] `json:"dimensional_price_group_id"`
	// The external id of the dimensional price group to include this price in
	ExternalDimensionalPriceGroupID param.Field[string] `json:"external_dimensional_price_group_id"`
}

func (r PlanExternalPlanIDVersionNewParamsAddPricesPriceNewPlanGroupedAllocationPriceDimensionalPriceConfiguration) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Within each billing cycle, specifies the cadence at which invoices are produced.
// If unspecified, a single invoice is produced per billing cycle.
type PlanExternalPlanIDVersionNewParamsAddPricesPriceNewPlanGroupedAllocationPriceInvoicingCycleConfiguration struct {
	// The duration of the billing period.
	Duration param.Field[int64] `json:"duration,required"`
	// The unit of billing period duration.
	DurationUnit param.Field[PlanExternalPlanIDVersionNewParamsAddPricesPriceNewPlanGroupedAllocationPriceInvoicingCycleConfigurationDurationUnit] `json:"duration_unit,required"`
}

func (r PlanExternalPlanIDVersionNewParamsAddPricesPriceNewPlanGroupedAllocationPriceInvoicingCycleConfiguration) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// The unit of billing period duration.
type PlanExternalPlanIDVersionNewParamsAddPricesPriceNewPlanGroupedAllocationPriceInvoicingCycleConfigurationDurationUnit string

const (
	PlanExternalPlanIDVersionNewParamsAddPricesPriceNewPlanGroupedAllocationPriceInvoicingCycleConfigurationDurationUnitDay   PlanExternalPlanIDVersionNewParamsAddPricesPriceNewPlanGroupedAllocationPriceInvoicingCycleConfigurationDurationUnit = "day"
	PlanExternalPlanIDVersionNewParamsAddPricesPriceNewPlanGroupedAllocationPriceInvoicingCycleConfigurationDurationUnitMonth PlanExternalPlanIDVersionNewParamsAddPricesPriceNewPlanGroupedAllocationPriceInvoicingCycleConfigurationDurationUnit = "month"
)

func (r PlanExternalPlanIDVersionNewParamsAddPricesPriceNewPlanGroupedAllocationPriceInvoicingCycleConfigurationDurationUnit) IsKnown() bool {
	switch r {
	case PlanExternalPlanIDVersionNewParamsAddPricesPriceNewPlanGroupedAllocationPriceInvoicingCycleConfigurationDurationUnitDay, PlanExternalPlanIDVersionNewParamsAddPricesPriceNewPlanGroupedAllocationPriceInvoicingCycleConfigurationDurationUnitMonth:
		return true
	}
	return false
}

type PlanExternalPlanIDVersionNewParamsAddPricesPriceNewPlanGroupedWithProratedMinimumPrice struct {
	// The cadence to bill for this price on.
	Cadence                          param.Field[PlanExternalPlanIDVersionNewParamsAddPricesPriceNewPlanGroupedWithProratedMinimumPriceCadence] `json:"cadence,required"`
	GroupedWithProratedMinimumConfig param.Field[map[string]interface{}]                                                                        `json:"grouped_with_prorated_minimum_config,required"`
	// The id of the item the price will be associated with.
	ItemID    param.Field[string]                                                                                          `json:"item_id,required"`
	ModelType param.Field[PlanExternalPlanIDVersionNewParamsAddPricesPriceNewPlanGroupedWithProratedMinimumPriceModelType] `json:"model_type,required"`
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
	BillingCycleConfiguration param.Field[PlanExternalPlanIDVersionNewParamsAddPricesPriceNewPlanGroupedWithProratedMinimumPriceBillingCycleConfiguration] `json:"billing_cycle_configuration"`
	// The per unit conversion rate of the price currency to the invoicing currency.
	ConversionRate param.Field[float64] `json:"conversion_rate"`
	// An ISO 4217 currency string, or custom pricing unit identifier, in which this
	// price is billed.
	Currency param.Field[string] `json:"currency"`
	// For dimensional price: specifies a price group and dimension values
	DimensionalPriceConfiguration param.Field[PlanExternalPlanIDVersionNewParamsAddPricesPriceNewPlanGroupedWithProratedMinimumPriceDimensionalPriceConfiguration] `json:"dimensional_price_configuration"`
	// An alias for the price.
	ExternalPriceID param.Field[string] `json:"external_price_id"`
	// If the Price represents a fixed cost, this represents the quantity of units
	// applied.
	FixedPriceQuantity param.Field[float64] `json:"fixed_price_quantity"`
	// The property used to group this price on an invoice
	InvoiceGroupingKey param.Field[string] `json:"invoice_grouping_key"`
	// Within each billing cycle, specifies the cadence at which invoices are produced.
	// If unspecified, a single invoice is produced per billing cycle.
	InvoicingCycleConfiguration param.Field[PlanExternalPlanIDVersionNewParamsAddPricesPriceNewPlanGroupedWithProratedMinimumPriceInvoicingCycleConfiguration] `json:"invoicing_cycle_configuration"`
	// User-specified key/value pairs for the resource. Individual keys can be removed
	// by setting the value to `null`, and the entire metadata mapping can be cleared
	// by setting `metadata` to `null`.
	Metadata param.Field[map[string]string] `json:"metadata"`
}

func (r PlanExternalPlanIDVersionNewParamsAddPricesPriceNewPlanGroupedWithProratedMinimumPrice) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r PlanExternalPlanIDVersionNewParamsAddPricesPriceNewPlanGroupedWithProratedMinimumPrice) implementsPlanExternalPlanIDVersionNewParamsAddPricesPriceUnion() {
}

// The cadence to bill for this price on.
type PlanExternalPlanIDVersionNewParamsAddPricesPriceNewPlanGroupedWithProratedMinimumPriceCadence string

const (
	PlanExternalPlanIDVersionNewParamsAddPricesPriceNewPlanGroupedWithProratedMinimumPriceCadenceAnnual     PlanExternalPlanIDVersionNewParamsAddPricesPriceNewPlanGroupedWithProratedMinimumPriceCadence = "annual"
	PlanExternalPlanIDVersionNewParamsAddPricesPriceNewPlanGroupedWithProratedMinimumPriceCadenceSemiAnnual PlanExternalPlanIDVersionNewParamsAddPricesPriceNewPlanGroupedWithProratedMinimumPriceCadence = "semi_annual"
	PlanExternalPlanIDVersionNewParamsAddPricesPriceNewPlanGroupedWithProratedMinimumPriceCadenceMonthly    PlanExternalPlanIDVersionNewParamsAddPricesPriceNewPlanGroupedWithProratedMinimumPriceCadence = "monthly"
	PlanExternalPlanIDVersionNewParamsAddPricesPriceNewPlanGroupedWithProratedMinimumPriceCadenceQuarterly  PlanExternalPlanIDVersionNewParamsAddPricesPriceNewPlanGroupedWithProratedMinimumPriceCadence = "quarterly"
	PlanExternalPlanIDVersionNewParamsAddPricesPriceNewPlanGroupedWithProratedMinimumPriceCadenceOneTime    PlanExternalPlanIDVersionNewParamsAddPricesPriceNewPlanGroupedWithProratedMinimumPriceCadence = "one_time"
	PlanExternalPlanIDVersionNewParamsAddPricesPriceNewPlanGroupedWithProratedMinimumPriceCadenceCustom     PlanExternalPlanIDVersionNewParamsAddPricesPriceNewPlanGroupedWithProratedMinimumPriceCadence = "custom"
)

func (r PlanExternalPlanIDVersionNewParamsAddPricesPriceNewPlanGroupedWithProratedMinimumPriceCadence) IsKnown() bool {
	switch r {
	case PlanExternalPlanIDVersionNewParamsAddPricesPriceNewPlanGroupedWithProratedMinimumPriceCadenceAnnual, PlanExternalPlanIDVersionNewParamsAddPricesPriceNewPlanGroupedWithProratedMinimumPriceCadenceSemiAnnual, PlanExternalPlanIDVersionNewParamsAddPricesPriceNewPlanGroupedWithProratedMinimumPriceCadenceMonthly, PlanExternalPlanIDVersionNewParamsAddPricesPriceNewPlanGroupedWithProratedMinimumPriceCadenceQuarterly, PlanExternalPlanIDVersionNewParamsAddPricesPriceNewPlanGroupedWithProratedMinimumPriceCadenceOneTime, PlanExternalPlanIDVersionNewParamsAddPricesPriceNewPlanGroupedWithProratedMinimumPriceCadenceCustom:
		return true
	}
	return false
}

type PlanExternalPlanIDVersionNewParamsAddPricesPriceNewPlanGroupedWithProratedMinimumPriceModelType string

const (
	PlanExternalPlanIDVersionNewParamsAddPricesPriceNewPlanGroupedWithProratedMinimumPriceModelTypeGroupedWithProratedMinimum PlanExternalPlanIDVersionNewParamsAddPricesPriceNewPlanGroupedWithProratedMinimumPriceModelType = "grouped_with_prorated_minimum"
)

func (r PlanExternalPlanIDVersionNewParamsAddPricesPriceNewPlanGroupedWithProratedMinimumPriceModelType) IsKnown() bool {
	switch r {
	case PlanExternalPlanIDVersionNewParamsAddPricesPriceNewPlanGroupedWithProratedMinimumPriceModelTypeGroupedWithProratedMinimum:
		return true
	}
	return false
}

// For custom cadence: specifies the duration of the billing period in days or
// months.
type PlanExternalPlanIDVersionNewParamsAddPricesPriceNewPlanGroupedWithProratedMinimumPriceBillingCycleConfiguration struct {
	// The duration of the billing period.
	Duration param.Field[int64] `json:"duration,required"`
	// The unit of billing period duration.
	DurationUnit param.Field[PlanExternalPlanIDVersionNewParamsAddPricesPriceNewPlanGroupedWithProratedMinimumPriceBillingCycleConfigurationDurationUnit] `json:"duration_unit,required"`
}

func (r PlanExternalPlanIDVersionNewParamsAddPricesPriceNewPlanGroupedWithProratedMinimumPriceBillingCycleConfiguration) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// The unit of billing period duration.
type PlanExternalPlanIDVersionNewParamsAddPricesPriceNewPlanGroupedWithProratedMinimumPriceBillingCycleConfigurationDurationUnit string

const (
	PlanExternalPlanIDVersionNewParamsAddPricesPriceNewPlanGroupedWithProratedMinimumPriceBillingCycleConfigurationDurationUnitDay   PlanExternalPlanIDVersionNewParamsAddPricesPriceNewPlanGroupedWithProratedMinimumPriceBillingCycleConfigurationDurationUnit = "day"
	PlanExternalPlanIDVersionNewParamsAddPricesPriceNewPlanGroupedWithProratedMinimumPriceBillingCycleConfigurationDurationUnitMonth PlanExternalPlanIDVersionNewParamsAddPricesPriceNewPlanGroupedWithProratedMinimumPriceBillingCycleConfigurationDurationUnit = "month"
)

func (r PlanExternalPlanIDVersionNewParamsAddPricesPriceNewPlanGroupedWithProratedMinimumPriceBillingCycleConfigurationDurationUnit) IsKnown() bool {
	switch r {
	case PlanExternalPlanIDVersionNewParamsAddPricesPriceNewPlanGroupedWithProratedMinimumPriceBillingCycleConfigurationDurationUnitDay, PlanExternalPlanIDVersionNewParamsAddPricesPriceNewPlanGroupedWithProratedMinimumPriceBillingCycleConfigurationDurationUnitMonth:
		return true
	}
	return false
}

// For dimensional price: specifies a price group and dimension values
type PlanExternalPlanIDVersionNewParamsAddPricesPriceNewPlanGroupedWithProratedMinimumPriceDimensionalPriceConfiguration struct {
	// The list of dimension values matching (in order) the dimensions of the price
	// group
	DimensionValues param.Field[[]string] `json:"dimension_values,required"`
	// The id of the dimensional price group to include this price in
	DimensionalPriceGroupID param.Field[string] `json:"dimensional_price_group_id"`
	// The external id of the dimensional price group to include this price in
	ExternalDimensionalPriceGroupID param.Field[string] `json:"external_dimensional_price_group_id"`
}

func (r PlanExternalPlanIDVersionNewParamsAddPricesPriceNewPlanGroupedWithProratedMinimumPriceDimensionalPriceConfiguration) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Within each billing cycle, specifies the cadence at which invoices are produced.
// If unspecified, a single invoice is produced per billing cycle.
type PlanExternalPlanIDVersionNewParamsAddPricesPriceNewPlanGroupedWithProratedMinimumPriceInvoicingCycleConfiguration struct {
	// The duration of the billing period.
	Duration param.Field[int64] `json:"duration,required"`
	// The unit of billing period duration.
	DurationUnit param.Field[PlanExternalPlanIDVersionNewParamsAddPricesPriceNewPlanGroupedWithProratedMinimumPriceInvoicingCycleConfigurationDurationUnit] `json:"duration_unit,required"`
}

func (r PlanExternalPlanIDVersionNewParamsAddPricesPriceNewPlanGroupedWithProratedMinimumPriceInvoicingCycleConfiguration) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// The unit of billing period duration.
type PlanExternalPlanIDVersionNewParamsAddPricesPriceNewPlanGroupedWithProratedMinimumPriceInvoicingCycleConfigurationDurationUnit string

const (
	PlanExternalPlanIDVersionNewParamsAddPricesPriceNewPlanGroupedWithProratedMinimumPriceInvoicingCycleConfigurationDurationUnitDay   PlanExternalPlanIDVersionNewParamsAddPricesPriceNewPlanGroupedWithProratedMinimumPriceInvoicingCycleConfigurationDurationUnit = "day"
	PlanExternalPlanIDVersionNewParamsAddPricesPriceNewPlanGroupedWithProratedMinimumPriceInvoicingCycleConfigurationDurationUnitMonth PlanExternalPlanIDVersionNewParamsAddPricesPriceNewPlanGroupedWithProratedMinimumPriceInvoicingCycleConfigurationDurationUnit = "month"
)

func (r PlanExternalPlanIDVersionNewParamsAddPricesPriceNewPlanGroupedWithProratedMinimumPriceInvoicingCycleConfigurationDurationUnit) IsKnown() bool {
	switch r {
	case PlanExternalPlanIDVersionNewParamsAddPricesPriceNewPlanGroupedWithProratedMinimumPriceInvoicingCycleConfigurationDurationUnitDay, PlanExternalPlanIDVersionNewParamsAddPricesPriceNewPlanGroupedWithProratedMinimumPriceInvoicingCycleConfigurationDurationUnitMonth:
		return true
	}
	return false
}

type PlanExternalPlanIDVersionNewParamsAddPricesPriceNewPlanGroupedWithMeteredMinimumPrice struct {
	// The cadence to bill for this price on.
	Cadence                         param.Field[PlanExternalPlanIDVersionNewParamsAddPricesPriceNewPlanGroupedWithMeteredMinimumPriceCadence] `json:"cadence,required"`
	GroupedWithMeteredMinimumConfig param.Field[map[string]interface{}]                                                                       `json:"grouped_with_metered_minimum_config,required"`
	// The id of the item the price will be associated with.
	ItemID    param.Field[string]                                                                                         `json:"item_id,required"`
	ModelType param.Field[PlanExternalPlanIDVersionNewParamsAddPricesPriceNewPlanGroupedWithMeteredMinimumPriceModelType] `json:"model_type,required"`
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
	BillingCycleConfiguration param.Field[PlanExternalPlanIDVersionNewParamsAddPricesPriceNewPlanGroupedWithMeteredMinimumPriceBillingCycleConfiguration] `json:"billing_cycle_configuration"`
	// The per unit conversion rate of the price currency to the invoicing currency.
	ConversionRate param.Field[float64] `json:"conversion_rate"`
	// An ISO 4217 currency string, or custom pricing unit identifier, in which this
	// price is billed.
	Currency param.Field[string] `json:"currency"`
	// For dimensional price: specifies a price group and dimension values
	DimensionalPriceConfiguration param.Field[PlanExternalPlanIDVersionNewParamsAddPricesPriceNewPlanGroupedWithMeteredMinimumPriceDimensionalPriceConfiguration] `json:"dimensional_price_configuration"`
	// An alias for the price.
	ExternalPriceID param.Field[string] `json:"external_price_id"`
	// If the Price represents a fixed cost, this represents the quantity of units
	// applied.
	FixedPriceQuantity param.Field[float64] `json:"fixed_price_quantity"`
	// The property used to group this price on an invoice
	InvoiceGroupingKey param.Field[string] `json:"invoice_grouping_key"`
	// Within each billing cycle, specifies the cadence at which invoices are produced.
	// If unspecified, a single invoice is produced per billing cycle.
	InvoicingCycleConfiguration param.Field[PlanExternalPlanIDVersionNewParamsAddPricesPriceNewPlanGroupedWithMeteredMinimumPriceInvoicingCycleConfiguration] `json:"invoicing_cycle_configuration"`
	// User-specified key/value pairs for the resource. Individual keys can be removed
	// by setting the value to `null`, and the entire metadata mapping can be cleared
	// by setting `metadata` to `null`.
	Metadata param.Field[map[string]string] `json:"metadata"`
}

func (r PlanExternalPlanIDVersionNewParamsAddPricesPriceNewPlanGroupedWithMeteredMinimumPrice) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r PlanExternalPlanIDVersionNewParamsAddPricesPriceNewPlanGroupedWithMeteredMinimumPrice) implementsPlanExternalPlanIDVersionNewParamsAddPricesPriceUnion() {
}

// The cadence to bill for this price on.
type PlanExternalPlanIDVersionNewParamsAddPricesPriceNewPlanGroupedWithMeteredMinimumPriceCadence string

const (
	PlanExternalPlanIDVersionNewParamsAddPricesPriceNewPlanGroupedWithMeteredMinimumPriceCadenceAnnual     PlanExternalPlanIDVersionNewParamsAddPricesPriceNewPlanGroupedWithMeteredMinimumPriceCadence = "annual"
	PlanExternalPlanIDVersionNewParamsAddPricesPriceNewPlanGroupedWithMeteredMinimumPriceCadenceSemiAnnual PlanExternalPlanIDVersionNewParamsAddPricesPriceNewPlanGroupedWithMeteredMinimumPriceCadence = "semi_annual"
	PlanExternalPlanIDVersionNewParamsAddPricesPriceNewPlanGroupedWithMeteredMinimumPriceCadenceMonthly    PlanExternalPlanIDVersionNewParamsAddPricesPriceNewPlanGroupedWithMeteredMinimumPriceCadence = "monthly"
	PlanExternalPlanIDVersionNewParamsAddPricesPriceNewPlanGroupedWithMeteredMinimumPriceCadenceQuarterly  PlanExternalPlanIDVersionNewParamsAddPricesPriceNewPlanGroupedWithMeteredMinimumPriceCadence = "quarterly"
	PlanExternalPlanIDVersionNewParamsAddPricesPriceNewPlanGroupedWithMeteredMinimumPriceCadenceOneTime    PlanExternalPlanIDVersionNewParamsAddPricesPriceNewPlanGroupedWithMeteredMinimumPriceCadence = "one_time"
	PlanExternalPlanIDVersionNewParamsAddPricesPriceNewPlanGroupedWithMeteredMinimumPriceCadenceCustom     PlanExternalPlanIDVersionNewParamsAddPricesPriceNewPlanGroupedWithMeteredMinimumPriceCadence = "custom"
)

func (r PlanExternalPlanIDVersionNewParamsAddPricesPriceNewPlanGroupedWithMeteredMinimumPriceCadence) IsKnown() bool {
	switch r {
	case PlanExternalPlanIDVersionNewParamsAddPricesPriceNewPlanGroupedWithMeteredMinimumPriceCadenceAnnual, PlanExternalPlanIDVersionNewParamsAddPricesPriceNewPlanGroupedWithMeteredMinimumPriceCadenceSemiAnnual, PlanExternalPlanIDVersionNewParamsAddPricesPriceNewPlanGroupedWithMeteredMinimumPriceCadenceMonthly, PlanExternalPlanIDVersionNewParamsAddPricesPriceNewPlanGroupedWithMeteredMinimumPriceCadenceQuarterly, PlanExternalPlanIDVersionNewParamsAddPricesPriceNewPlanGroupedWithMeteredMinimumPriceCadenceOneTime, PlanExternalPlanIDVersionNewParamsAddPricesPriceNewPlanGroupedWithMeteredMinimumPriceCadenceCustom:
		return true
	}
	return false
}

type PlanExternalPlanIDVersionNewParamsAddPricesPriceNewPlanGroupedWithMeteredMinimumPriceModelType string

const (
	PlanExternalPlanIDVersionNewParamsAddPricesPriceNewPlanGroupedWithMeteredMinimumPriceModelTypeGroupedWithMeteredMinimum PlanExternalPlanIDVersionNewParamsAddPricesPriceNewPlanGroupedWithMeteredMinimumPriceModelType = "grouped_with_metered_minimum"
)

func (r PlanExternalPlanIDVersionNewParamsAddPricesPriceNewPlanGroupedWithMeteredMinimumPriceModelType) IsKnown() bool {
	switch r {
	case PlanExternalPlanIDVersionNewParamsAddPricesPriceNewPlanGroupedWithMeteredMinimumPriceModelTypeGroupedWithMeteredMinimum:
		return true
	}
	return false
}

// For custom cadence: specifies the duration of the billing period in days or
// months.
type PlanExternalPlanIDVersionNewParamsAddPricesPriceNewPlanGroupedWithMeteredMinimumPriceBillingCycleConfiguration struct {
	// The duration of the billing period.
	Duration param.Field[int64] `json:"duration,required"`
	// The unit of billing period duration.
	DurationUnit param.Field[PlanExternalPlanIDVersionNewParamsAddPricesPriceNewPlanGroupedWithMeteredMinimumPriceBillingCycleConfigurationDurationUnit] `json:"duration_unit,required"`
}

func (r PlanExternalPlanIDVersionNewParamsAddPricesPriceNewPlanGroupedWithMeteredMinimumPriceBillingCycleConfiguration) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// The unit of billing period duration.
type PlanExternalPlanIDVersionNewParamsAddPricesPriceNewPlanGroupedWithMeteredMinimumPriceBillingCycleConfigurationDurationUnit string

const (
	PlanExternalPlanIDVersionNewParamsAddPricesPriceNewPlanGroupedWithMeteredMinimumPriceBillingCycleConfigurationDurationUnitDay   PlanExternalPlanIDVersionNewParamsAddPricesPriceNewPlanGroupedWithMeteredMinimumPriceBillingCycleConfigurationDurationUnit = "day"
	PlanExternalPlanIDVersionNewParamsAddPricesPriceNewPlanGroupedWithMeteredMinimumPriceBillingCycleConfigurationDurationUnitMonth PlanExternalPlanIDVersionNewParamsAddPricesPriceNewPlanGroupedWithMeteredMinimumPriceBillingCycleConfigurationDurationUnit = "month"
)

func (r PlanExternalPlanIDVersionNewParamsAddPricesPriceNewPlanGroupedWithMeteredMinimumPriceBillingCycleConfigurationDurationUnit) IsKnown() bool {
	switch r {
	case PlanExternalPlanIDVersionNewParamsAddPricesPriceNewPlanGroupedWithMeteredMinimumPriceBillingCycleConfigurationDurationUnitDay, PlanExternalPlanIDVersionNewParamsAddPricesPriceNewPlanGroupedWithMeteredMinimumPriceBillingCycleConfigurationDurationUnitMonth:
		return true
	}
	return false
}

// For dimensional price: specifies a price group and dimension values
type PlanExternalPlanIDVersionNewParamsAddPricesPriceNewPlanGroupedWithMeteredMinimumPriceDimensionalPriceConfiguration struct {
	// The list of dimension values matching (in order) the dimensions of the price
	// group
	DimensionValues param.Field[[]string] `json:"dimension_values,required"`
	// The id of the dimensional price group to include this price in
	DimensionalPriceGroupID param.Field[string] `json:"dimensional_price_group_id"`
	// The external id of the dimensional price group to include this price in
	ExternalDimensionalPriceGroupID param.Field[string] `json:"external_dimensional_price_group_id"`
}

func (r PlanExternalPlanIDVersionNewParamsAddPricesPriceNewPlanGroupedWithMeteredMinimumPriceDimensionalPriceConfiguration) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Within each billing cycle, specifies the cadence at which invoices are produced.
// If unspecified, a single invoice is produced per billing cycle.
type PlanExternalPlanIDVersionNewParamsAddPricesPriceNewPlanGroupedWithMeteredMinimumPriceInvoicingCycleConfiguration struct {
	// The duration of the billing period.
	Duration param.Field[int64] `json:"duration,required"`
	// The unit of billing period duration.
	DurationUnit param.Field[PlanExternalPlanIDVersionNewParamsAddPricesPriceNewPlanGroupedWithMeteredMinimumPriceInvoicingCycleConfigurationDurationUnit] `json:"duration_unit,required"`
}

func (r PlanExternalPlanIDVersionNewParamsAddPricesPriceNewPlanGroupedWithMeteredMinimumPriceInvoicingCycleConfiguration) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// The unit of billing period duration.
type PlanExternalPlanIDVersionNewParamsAddPricesPriceNewPlanGroupedWithMeteredMinimumPriceInvoicingCycleConfigurationDurationUnit string

const (
	PlanExternalPlanIDVersionNewParamsAddPricesPriceNewPlanGroupedWithMeteredMinimumPriceInvoicingCycleConfigurationDurationUnitDay   PlanExternalPlanIDVersionNewParamsAddPricesPriceNewPlanGroupedWithMeteredMinimumPriceInvoicingCycleConfigurationDurationUnit = "day"
	PlanExternalPlanIDVersionNewParamsAddPricesPriceNewPlanGroupedWithMeteredMinimumPriceInvoicingCycleConfigurationDurationUnitMonth PlanExternalPlanIDVersionNewParamsAddPricesPriceNewPlanGroupedWithMeteredMinimumPriceInvoicingCycleConfigurationDurationUnit = "month"
)

func (r PlanExternalPlanIDVersionNewParamsAddPricesPriceNewPlanGroupedWithMeteredMinimumPriceInvoicingCycleConfigurationDurationUnit) IsKnown() bool {
	switch r {
	case PlanExternalPlanIDVersionNewParamsAddPricesPriceNewPlanGroupedWithMeteredMinimumPriceInvoicingCycleConfigurationDurationUnitDay, PlanExternalPlanIDVersionNewParamsAddPricesPriceNewPlanGroupedWithMeteredMinimumPriceInvoicingCycleConfigurationDurationUnitMonth:
		return true
	}
	return false
}

type PlanExternalPlanIDVersionNewParamsAddPricesPriceNewPlanMatrixWithDisplayNamePrice struct {
	// The cadence to bill for this price on.
	Cadence param.Field[PlanExternalPlanIDVersionNewParamsAddPricesPriceNewPlanMatrixWithDisplayNamePriceCadence] `json:"cadence,required"`
	// The id of the item the price will be associated with.
	ItemID                      param.Field[string]                                                                                     `json:"item_id,required"`
	MatrixWithDisplayNameConfig param.Field[map[string]interface{}]                                                                     `json:"matrix_with_display_name_config,required"`
	ModelType                   param.Field[PlanExternalPlanIDVersionNewParamsAddPricesPriceNewPlanMatrixWithDisplayNamePriceModelType] `json:"model_type,required"`
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
	BillingCycleConfiguration param.Field[PlanExternalPlanIDVersionNewParamsAddPricesPriceNewPlanMatrixWithDisplayNamePriceBillingCycleConfiguration] `json:"billing_cycle_configuration"`
	// The per unit conversion rate of the price currency to the invoicing currency.
	ConversionRate param.Field[float64] `json:"conversion_rate"`
	// An ISO 4217 currency string, or custom pricing unit identifier, in which this
	// price is billed.
	Currency param.Field[string] `json:"currency"`
	// For dimensional price: specifies a price group and dimension values
	DimensionalPriceConfiguration param.Field[PlanExternalPlanIDVersionNewParamsAddPricesPriceNewPlanMatrixWithDisplayNamePriceDimensionalPriceConfiguration] `json:"dimensional_price_configuration"`
	// An alias for the price.
	ExternalPriceID param.Field[string] `json:"external_price_id"`
	// If the Price represents a fixed cost, this represents the quantity of units
	// applied.
	FixedPriceQuantity param.Field[float64] `json:"fixed_price_quantity"`
	// The property used to group this price on an invoice
	InvoiceGroupingKey param.Field[string] `json:"invoice_grouping_key"`
	// Within each billing cycle, specifies the cadence at which invoices are produced.
	// If unspecified, a single invoice is produced per billing cycle.
	InvoicingCycleConfiguration param.Field[PlanExternalPlanIDVersionNewParamsAddPricesPriceNewPlanMatrixWithDisplayNamePriceInvoicingCycleConfiguration] `json:"invoicing_cycle_configuration"`
	// User-specified key/value pairs for the resource. Individual keys can be removed
	// by setting the value to `null`, and the entire metadata mapping can be cleared
	// by setting `metadata` to `null`.
	Metadata param.Field[map[string]string] `json:"metadata"`
}

func (r PlanExternalPlanIDVersionNewParamsAddPricesPriceNewPlanMatrixWithDisplayNamePrice) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r PlanExternalPlanIDVersionNewParamsAddPricesPriceNewPlanMatrixWithDisplayNamePrice) implementsPlanExternalPlanIDVersionNewParamsAddPricesPriceUnion() {
}

// The cadence to bill for this price on.
type PlanExternalPlanIDVersionNewParamsAddPricesPriceNewPlanMatrixWithDisplayNamePriceCadence string

const (
	PlanExternalPlanIDVersionNewParamsAddPricesPriceNewPlanMatrixWithDisplayNamePriceCadenceAnnual     PlanExternalPlanIDVersionNewParamsAddPricesPriceNewPlanMatrixWithDisplayNamePriceCadence = "annual"
	PlanExternalPlanIDVersionNewParamsAddPricesPriceNewPlanMatrixWithDisplayNamePriceCadenceSemiAnnual PlanExternalPlanIDVersionNewParamsAddPricesPriceNewPlanMatrixWithDisplayNamePriceCadence = "semi_annual"
	PlanExternalPlanIDVersionNewParamsAddPricesPriceNewPlanMatrixWithDisplayNamePriceCadenceMonthly    PlanExternalPlanIDVersionNewParamsAddPricesPriceNewPlanMatrixWithDisplayNamePriceCadence = "monthly"
	PlanExternalPlanIDVersionNewParamsAddPricesPriceNewPlanMatrixWithDisplayNamePriceCadenceQuarterly  PlanExternalPlanIDVersionNewParamsAddPricesPriceNewPlanMatrixWithDisplayNamePriceCadence = "quarterly"
	PlanExternalPlanIDVersionNewParamsAddPricesPriceNewPlanMatrixWithDisplayNamePriceCadenceOneTime    PlanExternalPlanIDVersionNewParamsAddPricesPriceNewPlanMatrixWithDisplayNamePriceCadence = "one_time"
	PlanExternalPlanIDVersionNewParamsAddPricesPriceNewPlanMatrixWithDisplayNamePriceCadenceCustom     PlanExternalPlanIDVersionNewParamsAddPricesPriceNewPlanMatrixWithDisplayNamePriceCadence = "custom"
)

func (r PlanExternalPlanIDVersionNewParamsAddPricesPriceNewPlanMatrixWithDisplayNamePriceCadence) IsKnown() bool {
	switch r {
	case PlanExternalPlanIDVersionNewParamsAddPricesPriceNewPlanMatrixWithDisplayNamePriceCadenceAnnual, PlanExternalPlanIDVersionNewParamsAddPricesPriceNewPlanMatrixWithDisplayNamePriceCadenceSemiAnnual, PlanExternalPlanIDVersionNewParamsAddPricesPriceNewPlanMatrixWithDisplayNamePriceCadenceMonthly, PlanExternalPlanIDVersionNewParamsAddPricesPriceNewPlanMatrixWithDisplayNamePriceCadenceQuarterly, PlanExternalPlanIDVersionNewParamsAddPricesPriceNewPlanMatrixWithDisplayNamePriceCadenceOneTime, PlanExternalPlanIDVersionNewParamsAddPricesPriceNewPlanMatrixWithDisplayNamePriceCadenceCustom:
		return true
	}
	return false
}

type PlanExternalPlanIDVersionNewParamsAddPricesPriceNewPlanMatrixWithDisplayNamePriceModelType string

const (
	PlanExternalPlanIDVersionNewParamsAddPricesPriceNewPlanMatrixWithDisplayNamePriceModelTypeMatrixWithDisplayName PlanExternalPlanIDVersionNewParamsAddPricesPriceNewPlanMatrixWithDisplayNamePriceModelType = "matrix_with_display_name"
)

func (r PlanExternalPlanIDVersionNewParamsAddPricesPriceNewPlanMatrixWithDisplayNamePriceModelType) IsKnown() bool {
	switch r {
	case PlanExternalPlanIDVersionNewParamsAddPricesPriceNewPlanMatrixWithDisplayNamePriceModelTypeMatrixWithDisplayName:
		return true
	}
	return false
}

// For custom cadence: specifies the duration of the billing period in days or
// months.
type PlanExternalPlanIDVersionNewParamsAddPricesPriceNewPlanMatrixWithDisplayNamePriceBillingCycleConfiguration struct {
	// The duration of the billing period.
	Duration param.Field[int64] `json:"duration,required"`
	// The unit of billing period duration.
	DurationUnit param.Field[PlanExternalPlanIDVersionNewParamsAddPricesPriceNewPlanMatrixWithDisplayNamePriceBillingCycleConfigurationDurationUnit] `json:"duration_unit,required"`
}

func (r PlanExternalPlanIDVersionNewParamsAddPricesPriceNewPlanMatrixWithDisplayNamePriceBillingCycleConfiguration) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// The unit of billing period duration.
type PlanExternalPlanIDVersionNewParamsAddPricesPriceNewPlanMatrixWithDisplayNamePriceBillingCycleConfigurationDurationUnit string

const (
	PlanExternalPlanIDVersionNewParamsAddPricesPriceNewPlanMatrixWithDisplayNamePriceBillingCycleConfigurationDurationUnitDay   PlanExternalPlanIDVersionNewParamsAddPricesPriceNewPlanMatrixWithDisplayNamePriceBillingCycleConfigurationDurationUnit = "day"
	PlanExternalPlanIDVersionNewParamsAddPricesPriceNewPlanMatrixWithDisplayNamePriceBillingCycleConfigurationDurationUnitMonth PlanExternalPlanIDVersionNewParamsAddPricesPriceNewPlanMatrixWithDisplayNamePriceBillingCycleConfigurationDurationUnit = "month"
)

func (r PlanExternalPlanIDVersionNewParamsAddPricesPriceNewPlanMatrixWithDisplayNamePriceBillingCycleConfigurationDurationUnit) IsKnown() bool {
	switch r {
	case PlanExternalPlanIDVersionNewParamsAddPricesPriceNewPlanMatrixWithDisplayNamePriceBillingCycleConfigurationDurationUnitDay, PlanExternalPlanIDVersionNewParamsAddPricesPriceNewPlanMatrixWithDisplayNamePriceBillingCycleConfigurationDurationUnitMonth:
		return true
	}
	return false
}

// For dimensional price: specifies a price group and dimension values
type PlanExternalPlanIDVersionNewParamsAddPricesPriceNewPlanMatrixWithDisplayNamePriceDimensionalPriceConfiguration struct {
	// The list of dimension values matching (in order) the dimensions of the price
	// group
	DimensionValues param.Field[[]string] `json:"dimension_values,required"`
	// The id of the dimensional price group to include this price in
	DimensionalPriceGroupID param.Field[string] `json:"dimensional_price_group_id"`
	// The external id of the dimensional price group to include this price in
	ExternalDimensionalPriceGroupID param.Field[string] `json:"external_dimensional_price_group_id"`
}

func (r PlanExternalPlanIDVersionNewParamsAddPricesPriceNewPlanMatrixWithDisplayNamePriceDimensionalPriceConfiguration) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Within each billing cycle, specifies the cadence at which invoices are produced.
// If unspecified, a single invoice is produced per billing cycle.
type PlanExternalPlanIDVersionNewParamsAddPricesPriceNewPlanMatrixWithDisplayNamePriceInvoicingCycleConfiguration struct {
	// The duration of the billing period.
	Duration param.Field[int64] `json:"duration,required"`
	// The unit of billing period duration.
	DurationUnit param.Field[PlanExternalPlanIDVersionNewParamsAddPricesPriceNewPlanMatrixWithDisplayNamePriceInvoicingCycleConfigurationDurationUnit] `json:"duration_unit,required"`
}

func (r PlanExternalPlanIDVersionNewParamsAddPricesPriceNewPlanMatrixWithDisplayNamePriceInvoicingCycleConfiguration) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// The unit of billing period duration.
type PlanExternalPlanIDVersionNewParamsAddPricesPriceNewPlanMatrixWithDisplayNamePriceInvoicingCycleConfigurationDurationUnit string

const (
	PlanExternalPlanIDVersionNewParamsAddPricesPriceNewPlanMatrixWithDisplayNamePriceInvoicingCycleConfigurationDurationUnitDay   PlanExternalPlanIDVersionNewParamsAddPricesPriceNewPlanMatrixWithDisplayNamePriceInvoicingCycleConfigurationDurationUnit = "day"
	PlanExternalPlanIDVersionNewParamsAddPricesPriceNewPlanMatrixWithDisplayNamePriceInvoicingCycleConfigurationDurationUnitMonth PlanExternalPlanIDVersionNewParamsAddPricesPriceNewPlanMatrixWithDisplayNamePriceInvoicingCycleConfigurationDurationUnit = "month"
)

func (r PlanExternalPlanIDVersionNewParamsAddPricesPriceNewPlanMatrixWithDisplayNamePriceInvoicingCycleConfigurationDurationUnit) IsKnown() bool {
	switch r {
	case PlanExternalPlanIDVersionNewParamsAddPricesPriceNewPlanMatrixWithDisplayNamePriceInvoicingCycleConfigurationDurationUnitDay, PlanExternalPlanIDVersionNewParamsAddPricesPriceNewPlanMatrixWithDisplayNamePriceInvoicingCycleConfigurationDurationUnitMonth:
		return true
	}
	return false
}

type PlanExternalPlanIDVersionNewParamsAddPricesPriceNewPlanBulkWithProrationPrice struct {
	BulkWithProrationConfig param.Field[map[string]interface{}] `json:"bulk_with_proration_config,required"`
	// The cadence to bill for this price on.
	Cadence param.Field[PlanExternalPlanIDVersionNewParamsAddPricesPriceNewPlanBulkWithProrationPriceCadence] `json:"cadence,required"`
	// The id of the item the price will be associated with.
	ItemID    param.Field[string]                                                                                 `json:"item_id,required"`
	ModelType param.Field[PlanExternalPlanIDVersionNewParamsAddPricesPriceNewPlanBulkWithProrationPriceModelType] `json:"model_type,required"`
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
	BillingCycleConfiguration param.Field[PlanExternalPlanIDVersionNewParamsAddPricesPriceNewPlanBulkWithProrationPriceBillingCycleConfiguration] `json:"billing_cycle_configuration"`
	// The per unit conversion rate of the price currency to the invoicing currency.
	ConversionRate param.Field[float64] `json:"conversion_rate"`
	// An ISO 4217 currency string, or custom pricing unit identifier, in which this
	// price is billed.
	Currency param.Field[string] `json:"currency"`
	// For dimensional price: specifies a price group and dimension values
	DimensionalPriceConfiguration param.Field[PlanExternalPlanIDVersionNewParamsAddPricesPriceNewPlanBulkWithProrationPriceDimensionalPriceConfiguration] `json:"dimensional_price_configuration"`
	// An alias for the price.
	ExternalPriceID param.Field[string] `json:"external_price_id"`
	// If the Price represents a fixed cost, this represents the quantity of units
	// applied.
	FixedPriceQuantity param.Field[float64] `json:"fixed_price_quantity"`
	// The property used to group this price on an invoice
	InvoiceGroupingKey param.Field[string] `json:"invoice_grouping_key"`
	// Within each billing cycle, specifies the cadence at which invoices are produced.
	// If unspecified, a single invoice is produced per billing cycle.
	InvoicingCycleConfiguration param.Field[PlanExternalPlanIDVersionNewParamsAddPricesPriceNewPlanBulkWithProrationPriceInvoicingCycleConfiguration] `json:"invoicing_cycle_configuration"`
	// User-specified key/value pairs for the resource. Individual keys can be removed
	// by setting the value to `null`, and the entire metadata mapping can be cleared
	// by setting `metadata` to `null`.
	Metadata param.Field[map[string]string] `json:"metadata"`
}

func (r PlanExternalPlanIDVersionNewParamsAddPricesPriceNewPlanBulkWithProrationPrice) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r PlanExternalPlanIDVersionNewParamsAddPricesPriceNewPlanBulkWithProrationPrice) implementsPlanExternalPlanIDVersionNewParamsAddPricesPriceUnion() {
}

// The cadence to bill for this price on.
type PlanExternalPlanIDVersionNewParamsAddPricesPriceNewPlanBulkWithProrationPriceCadence string

const (
	PlanExternalPlanIDVersionNewParamsAddPricesPriceNewPlanBulkWithProrationPriceCadenceAnnual     PlanExternalPlanIDVersionNewParamsAddPricesPriceNewPlanBulkWithProrationPriceCadence = "annual"
	PlanExternalPlanIDVersionNewParamsAddPricesPriceNewPlanBulkWithProrationPriceCadenceSemiAnnual PlanExternalPlanIDVersionNewParamsAddPricesPriceNewPlanBulkWithProrationPriceCadence = "semi_annual"
	PlanExternalPlanIDVersionNewParamsAddPricesPriceNewPlanBulkWithProrationPriceCadenceMonthly    PlanExternalPlanIDVersionNewParamsAddPricesPriceNewPlanBulkWithProrationPriceCadence = "monthly"
	PlanExternalPlanIDVersionNewParamsAddPricesPriceNewPlanBulkWithProrationPriceCadenceQuarterly  PlanExternalPlanIDVersionNewParamsAddPricesPriceNewPlanBulkWithProrationPriceCadence = "quarterly"
	PlanExternalPlanIDVersionNewParamsAddPricesPriceNewPlanBulkWithProrationPriceCadenceOneTime    PlanExternalPlanIDVersionNewParamsAddPricesPriceNewPlanBulkWithProrationPriceCadence = "one_time"
	PlanExternalPlanIDVersionNewParamsAddPricesPriceNewPlanBulkWithProrationPriceCadenceCustom     PlanExternalPlanIDVersionNewParamsAddPricesPriceNewPlanBulkWithProrationPriceCadence = "custom"
)

func (r PlanExternalPlanIDVersionNewParamsAddPricesPriceNewPlanBulkWithProrationPriceCadence) IsKnown() bool {
	switch r {
	case PlanExternalPlanIDVersionNewParamsAddPricesPriceNewPlanBulkWithProrationPriceCadenceAnnual, PlanExternalPlanIDVersionNewParamsAddPricesPriceNewPlanBulkWithProrationPriceCadenceSemiAnnual, PlanExternalPlanIDVersionNewParamsAddPricesPriceNewPlanBulkWithProrationPriceCadenceMonthly, PlanExternalPlanIDVersionNewParamsAddPricesPriceNewPlanBulkWithProrationPriceCadenceQuarterly, PlanExternalPlanIDVersionNewParamsAddPricesPriceNewPlanBulkWithProrationPriceCadenceOneTime, PlanExternalPlanIDVersionNewParamsAddPricesPriceNewPlanBulkWithProrationPriceCadenceCustom:
		return true
	}
	return false
}

type PlanExternalPlanIDVersionNewParamsAddPricesPriceNewPlanBulkWithProrationPriceModelType string

const (
	PlanExternalPlanIDVersionNewParamsAddPricesPriceNewPlanBulkWithProrationPriceModelTypeBulkWithProration PlanExternalPlanIDVersionNewParamsAddPricesPriceNewPlanBulkWithProrationPriceModelType = "bulk_with_proration"
)

func (r PlanExternalPlanIDVersionNewParamsAddPricesPriceNewPlanBulkWithProrationPriceModelType) IsKnown() bool {
	switch r {
	case PlanExternalPlanIDVersionNewParamsAddPricesPriceNewPlanBulkWithProrationPriceModelTypeBulkWithProration:
		return true
	}
	return false
}

// For custom cadence: specifies the duration of the billing period in days or
// months.
type PlanExternalPlanIDVersionNewParamsAddPricesPriceNewPlanBulkWithProrationPriceBillingCycleConfiguration struct {
	// The duration of the billing period.
	Duration param.Field[int64] `json:"duration,required"`
	// The unit of billing period duration.
	DurationUnit param.Field[PlanExternalPlanIDVersionNewParamsAddPricesPriceNewPlanBulkWithProrationPriceBillingCycleConfigurationDurationUnit] `json:"duration_unit,required"`
}

func (r PlanExternalPlanIDVersionNewParamsAddPricesPriceNewPlanBulkWithProrationPriceBillingCycleConfiguration) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// The unit of billing period duration.
type PlanExternalPlanIDVersionNewParamsAddPricesPriceNewPlanBulkWithProrationPriceBillingCycleConfigurationDurationUnit string

const (
	PlanExternalPlanIDVersionNewParamsAddPricesPriceNewPlanBulkWithProrationPriceBillingCycleConfigurationDurationUnitDay   PlanExternalPlanIDVersionNewParamsAddPricesPriceNewPlanBulkWithProrationPriceBillingCycleConfigurationDurationUnit = "day"
	PlanExternalPlanIDVersionNewParamsAddPricesPriceNewPlanBulkWithProrationPriceBillingCycleConfigurationDurationUnitMonth PlanExternalPlanIDVersionNewParamsAddPricesPriceNewPlanBulkWithProrationPriceBillingCycleConfigurationDurationUnit = "month"
)

func (r PlanExternalPlanIDVersionNewParamsAddPricesPriceNewPlanBulkWithProrationPriceBillingCycleConfigurationDurationUnit) IsKnown() bool {
	switch r {
	case PlanExternalPlanIDVersionNewParamsAddPricesPriceNewPlanBulkWithProrationPriceBillingCycleConfigurationDurationUnitDay, PlanExternalPlanIDVersionNewParamsAddPricesPriceNewPlanBulkWithProrationPriceBillingCycleConfigurationDurationUnitMonth:
		return true
	}
	return false
}

// For dimensional price: specifies a price group and dimension values
type PlanExternalPlanIDVersionNewParamsAddPricesPriceNewPlanBulkWithProrationPriceDimensionalPriceConfiguration struct {
	// The list of dimension values matching (in order) the dimensions of the price
	// group
	DimensionValues param.Field[[]string] `json:"dimension_values,required"`
	// The id of the dimensional price group to include this price in
	DimensionalPriceGroupID param.Field[string] `json:"dimensional_price_group_id"`
	// The external id of the dimensional price group to include this price in
	ExternalDimensionalPriceGroupID param.Field[string] `json:"external_dimensional_price_group_id"`
}

func (r PlanExternalPlanIDVersionNewParamsAddPricesPriceNewPlanBulkWithProrationPriceDimensionalPriceConfiguration) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Within each billing cycle, specifies the cadence at which invoices are produced.
// If unspecified, a single invoice is produced per billing cycle.
type PlanExternalPlanIDVersionNewParamsAddPricesPriceNewPlanBulkWithProrationPriceInvoicingCycleConfiguration struct {
	// The duration of the billing period.
	Duration param.Field[int64] `json:"duration,required"`
	// The unit of billing period duration.
	DurationUnit param.Field[PlanExternalPlanIDVersionNewParamsAddPricesPriceNewPlanBulkWithProrationPriceInvoicingCycleConfigurationDurationUnit] `json:"duration_unit,required"`
}

func (r PlanExternalPlanIDVersionNewParamsAddPricesPriceNewPlanBulkWithProrationPriceInvoicingCycleConfiguration) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// The unit of billing period duration.
type PlanExternalPlanIDVersionNewParamsAddPricesPriceNewPlanBulkWithProrationPriceInvoicingCycleConfigurationDurationUnit string

const (
	PlanExternalPlanIDVersionNewParamsAddPricesPriceNewPlanBulkWithProrationPriceInvoicingCycleConfigurationDurationUnitDay   PlanExternalPlanIDVersionNewParamsAddPricesPriceNewPlanBulkWithProrationPriceInvoicingCycleConfigurationDurationUnit = "day"
	PlanExternalPlanIDVersionNewParamsAddPricesPriceNewPlanBulkWithProrationPriceInvoicingCycleConfigurationDurationUnitMonth PlanExternalPlanIDVersionNewParamsAddPricesPriceNewPlanBulkWithProrationPriceInvoicingCycleConfigurationDurationUnit = "month"
)

func (r PlanExternalPlanIDVersionNewParamsAddPricesPriceNewPlanBulkWithProrationPriceInvoicingCycleConfigurationDurationUnit) IsKnown() bool {
	switch r {
	case PlanExternalPlanIDVersionNewParamsAddPricesPriceNewPlanBulkWithProrationPriceInvoicingCycleConfigurationDurationUnitDay, PlanExternalPlanIDVersionNewParamsAddPricesPriceNewPlanBulkWithProrationPriceInvoicingCycleConfigurationDurationUnitMonth:
		return true
	}
	return false
}

type PlanExternalPlanIDVersionNewParamsAddPricesPriceNewPlanGroupedTieredPackagePrice struct {
	// The cadence to bill for this price on.
	Cadence                    param.Field[PlanExternalPlanIDVersionNewParamsAddPricesPriceNewPlanGroupedTieredPackagePriceCadence] `json:"cadence,required"`
	GroupedTieredPackageConfig param.Field[map[string]interface{}]                                                                  `json:"grouped_tiered_package_config,required"`
	// The id of the item the price will be associated with.
	ItemID    param.Field[string]                                                                                    `json:"item_id,required"`
	ModelType param.Field[PlanExternalPlanIDVersionNewParamsAddPricesPriceNewPlanGroupedTieredPackagePriceModelType] `json:"model_type,required"`
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
	BillingCycleConfiguration param.Field[PlanExternalPlanIDVersionNewParamsAddPricesPriceNewPlanGroupedTieredPackagePriceBillingCycleConfiguration] `json:"billing_cycle_configuration"`
	// The per unit conversion rate of the price currency to the invoicing currency.
	ConversionRate param.Field[float64] `json:"conversion_rate"`
	// An ISO 4217 currency string, or custom pricing unit identifier, in which this
	// price is billed.
	Currency param.Field[string] `json:"currency"`
	// For dimensional price: specifies a price group and dimension values
	DimensionalPriceConfiguration param.Field[PlanExternalPlanIDVersionNewParamsAddPricesPriceNewPlanGroupedTieredPackagePriceDimensionalPriceConfiguration] `json:"dimensional_price_configuration"`
	// An alias for the price.
	ExternalPriceID param.Field[string] `json:"external_price_id"`
	// If the Price represents a fixed cost, this represents the quantity of units
	// applied.
	FixedPriceQuantity param.Field[float64] `json:"fixed_price_quantity"`
	// The property used to group this price on an invoice
	InvoiceGroupingKey param.Field[string] `json:"invoice_grouping_key"`
	// Within each billing cycle, specifies the cadence at which invoices are produced.
	// If unspecified, a single invoice is produced per billing cycle.
	InvoicingCycleConfiguration param.Field[PlanExternalPlanIDVersionNewParamsAddPricesPriceNewPlanGroupedTieredPackagePriceInvoicingCycleConfiguration] `json:"invoicing_cycle_configuration"`
	// User-specified key/value pairs for the resource. Individual keys can be removed
	// by setting the value to `null`, and the entire metadata mapping can be cleared
	// by setting `metadata` to `null`.
	Metadata param.Field[map[string]string] `json:"metadata"`
}

func (r PlanExternalPlanIDVersionNewParamsAddPricesPriceNewPlanGroupedTieredPackagePrice) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r PlanExternalPlanIDVersionNewParamsAddPricesPriceNewPlanGroupedTieredPackagePrice) implementsPlanExternalPlanIDVersionNewParamsAddPricesPriceUnion() {
}

// The cadence to bill for this price on.
type PlanExternalPlanIDVersionNewParamsAddPricesPriceNewPlanGroupedTieredPackagePriceCadence string

const (
	PlanExternalPlanIDVersionNewParamsAddPricesPriceNewPlanGroupedTieredPackagePriceCadenceAnnual     PlanExternalPlanIDVersionNewParamsAddPricesPriceNewPlanGroupedTieredPackagePriceCadence = "annual"
	PlanExternalPlanIDVersionNewParamsAddPricesPriceNewPlanGroupedTieredPackagePriceCadenceSemiAnnual PlanExternalPlanIDVersionNewParamsAddPricesPriceNewPlanGroupedTieredPackagePriceCadence = "semi_annual"
	PlanExternalPlanIDVersionNewParamsAddPricesPriceNewPlanGroupedTieredPackagePriceCadenceMonthly    PlanExternalPlanIDVersionNewParamsAddPricesPriceNewPlanGroupedTieredPackagePriceCadence = "monthly"
	PlanExternalPlanIDVersionNewParamsAddPricesPriceNewPlanGroupedTieredPackagePriceCadenceQuarterly  PlanExternalPlanIDVersionNewParamsAddPricesPriceNewPlanGroupedTieredPackagePriceCadence = "quarterly"
	PlanExternalPlanIDVersionNewParamsAddPricesPriceNewPlanGroupedTieredPackagePriceCadenceOneTime    PlanExternalPlanIDVersionNewParamsAddPricesPriceNewPlanGroupedTieredPackagePriceCadence = "one_time"
	PlanExternalPlanIDVersionNewParamsAddPricesPriceNewPlanGroupedTieredPackagePriceCadenceCustom     PlanExternalPlanIDVersionNewParamsAddPricesPriceNewPlanGroupedTieredPackagePriceCadence = "custom"
)

func (r PlanExternalPlanIDVersionNewParamsAddPricesPriceNewPlanGroupedTieredPackagePriceCadence) IsKnown() bool {
	switch r {
	case PlanExternalPlanIDVersionNewParamsAddPricesPriceNewPlanGroupedTieredPackagePriceCadenceAnnual, PlanExternalPlanIDVersionNewParamsAddPricesPriceNewPlanGroupedTieredPackagePriceCadenceSemiAnnual, PlanExternalPlanIDVersionNewParamsAddPricesPriceNewPlanGroupedTieredPackagePriceCadenceMonthly, PlanExternalPlanIDVersionNewParamsAddPricesPriceNewPlanGroupedTieredPackagePriceCadenceQuarterly, PlanExternalPlanIDVersionNewParamsAddPricesPriceNewPlanGroupedTieredPackagePriceCadenceOneTime, PlanExternalPlanIDVersionNewParamsAddPricesPriceNewPlanGroupedTieredPackagePriceCadenceCustom:
		return true
	}
	return false
}

type PlanExternalPlanIDVersionNewParamsAddPricesPriceNewPlanGroupedTieredPackagePriceModelType string

const (
	PlanExternalPlanIDVersionNewParamsAddPricesPriceNewPlanGroupedTieredPackagePriceModelTypeGroupedTieredPackage PlanExternalPlanIDVersionNewParamsAddPricesPriceNewPlanGroupedTieredPackagePriceModelType = "grouped_tiered_package"
)

func (r PlanExternalPlanIDVersionNewParamsAddPricesPriceNewPlanGroupedTieredPackagePriceModelType) IsKnown() bool {
	switch r {
	case PlanExternalPlanIDVersionNewParamsAddPricesPriceNewPlanGroupedTieredPackagePriceModelTypeGroupedTieredPackage:
		return true
	}
	return false
}

// For custom cadence: specifies the duration of the billing period in days or
// months.
type PlanExternalPlanIDVersionNewParamsAddPricesPriceNewPlanGroupedTieredPackagePriceBillingCycleConfiguration struct {
	// The duration of the billing period.
	Duration param.Field[int64] `json:"duration,required"`
	// The unit of billing period duration.
	DurationUnit param.Field[PlanExternalPlanIDVersionNewParamsAddPricesPriceNewPlanGroupedTieredPackagePriceBillingCycleConfigurationDurationUnit] `json:"duration_unit,required"`
}

func (r PlanExternalPlanIDVersionNewParamsAddPricesPriceNewPlanGroupedTieredPackagePriceBillingCycleConfiguration) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// The unit of billing period duration.
type PlanExternalPlanIDVersionNewParamsAddPricesPriceNewPlanGroupedTieredPackagePriceBillingCycleConfigurationDurationUnit string

const (
	PlanExternalPlanIDVersionNewParamsAddPricesPriceNewPlanGroupedTieredPackagePriceBillingCycleConfigurationDurationUnitDay   PlanExternalPlanIDVersionNewParamsAddPricesPriceNewPlanGroupedTieredPackagePriceBillingCycleConfigurationDurationUnit = "day"
	PlanExternalPlanIDVersionNewParamsAddPricesPriceNewPlanGroupedTieredPackagePriceBillingCycleConfigurationDurationUnitMonth PlanExternalPlanIDVersionNewParamsAddPricesPriceNewPlanGroupedTieredPackagePriceBillingCycleConfigurationDurationUnit = "month"
)

func (r PlanExternalPlanIDVersionNewParamsAddPricesPriceNewPlanGroupedTieredPackagePriceBillingCycleConfigurationDurationUnit) IsKnown() bool {
	switch r {
	case PlanExternalPlanIDVersionNewParamsAddPricesPriceNewPlanGroupedTieredPackagePriceBillingCycleConfigurationDurationUnitDay, PlanExternalPlanIDVersionNewParamsAddPricesPriceNewPlanGroupedTieredPackagePriceBillingCycleConfigurationDurationUnitMonth:
		return true
	}
	return false
}

// For dimensional price: specifies a price group and dimension values
type PlanExternalPlanIDVersionNewParamsAddPricesPriceNewPlanGroupedTieredPackagePriceDimensionalPriceConfiguration struct {
	// The list of dimension values matching (in order) the dimensions of the price
	// group
	DimensionValues param.Field[[]string] `json:"dimension_values,required"`
	// The id of the dimensional price group to include this price in
	DimensionalPriceGroupID param.Field[string] `json:"dimensional_price_group_id"`
	// The external id of the dimensional price group to include this price in
	ExternalDimensionalPriceGroupID param.Field[string] `json:"external_dimensional_price_group_id"`
}

func (r PlanExternalPlanIDVersionNewParamsAddPricesPriceNewPlanGroupedTieredPackagePriceDimensionalPriceConfiguration) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Within each billing cycle, specifies the cadence at which invoices are produced.
// If unspecified, a single invoice is produced per billing cycle.
type PlanExternalPlanIDVersionNewParamsAddPricesPriceNewPlanGroupedTieredPackagePriceInvoicingCycleConfiguration struct {
	// The duration of the billing period.
	Duration param.Field[int64] `json:"duration,required"`
	// The unit of billing period duration.
	DurationUnit param.Field[PlanExternalPlanIDVersionNewParamsAddPricesPriceNewPlanGroupedTieredPackagePriceInvoicingCycleConfigurationDurationUnit] `json:"duration_unit,required"`
}

func (r PlanExternalPlanIDVersionNewParamsAddPricesPriceNewPlanGroupedTieredPackagePriceInvoicingCycleConfiguration) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// The unit of billing period duration.
type PlanExternalPlanIDVersionNewParamsAddPricesPriceNewPlanGroupedTieredPackagePriceInvoicingCycleConfigurationDurationUnit string

const (
	PlanExternalPlanIDVersionNewParamsAddPricesPriceNewPlanGroupedTieredPackagePriceInvoicingCycleConfigurationDurationUnitDay   PlanExternalPlanIDVersionNewParamsAddPricesPriceNewPlanGroupedTieredPackagePriceInvoicingCycleConfigurationDurationUnit = "day"
	PlanExternalPlanIDVersionNewParamsAddPricesPriceNewPlanGroupedTieredPackagePriceInvoicingCycleConfigurationDurationUnitMonth PlanExternalPlanIDVersionNewParamsAddPricesPriceNewPlanGroupedTieredPackagePriceInvoicingCycleConfigurationDurationUnit = "month"
)

func (r PlanExternalPlanIDVersionNewParamsAddPricesPriceNewPlanGroupedTieredPackagePriceInvoicingCycleConfigurationDurationUnit) IsKnown() bool {
	switch r {
	case PlanExternalPlanIDVersionNewParamsAddPricesPriceNewPlanGroupedTieredPackagePriceInvoicingCycleConfigurationDurationUnitDay, PlanExternalPlanIDVersionNewParamsAddPricesPriceNewPlanGroupedTieredPackagePriceInvoicingCycleConfigurationDurationUnitMonth:
		return true
	}
	return false
}

type PlanExternalPlanIDVersionNewParamsAddPricesPriceNewPlanMaxGroupTieredPackagePrice struct {
	// The cadence to bill for this price on.
	Cadence param.Field[PlanExternalPlanIDVersionNewParamsAddPricesPriceNewPlanMaxGroupTieredPackagePriceCadence] `json:"cadence,required"`
	// The id of the item the price will be associated with.
	ItemID                      param.Field[string]                                                                                     `json:"item_id,required"`
	MaxGroupTieredPackageConfig param.Field[map[string]interface{}]                                                                     `json:"max_group_tiered_package_config,required"`
	ModelType                   param.Field[PlanExternalPlanIDVersionNewParamsAddPricesPriceNewPlanMaxGroupTieredPackagePriceModelType] `json:"model_type,required"`
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
	BillingCycleConfiguration param.Field[PlanExternalPlanIDVersionNewParamsAddPricesPriceNewPlanMaxGroupTieredPackagePriceBillingCycleConfiguration] `json:"billing_cycle_configuration"`
	// The per unit conversion rate of the price currency to the invoicing currency.
	ConversionRate param.Field[float64] `json:"conversion_rate"`
	// An ISO 4217 currency string, or custom pricing unit identifier, in which this
	// price is billed.
	Currency param.Field[string] `json:"currency"`
	// For dimensional price: specifies a price group and dimension values
	DimensionalPriceConfiguration param.Field[PlanExternalPlanIDVersionNewParamsAddPricesPriceNewPlanMaxGroupTieredPackagePriceDimensionalPriceConfiguration] `json:"dimensional_price_configuration"`
	// An alias for the price.
	ExternalPriceID param.Field[string] `json:"external_price_id"`
	// If the Price represents a fixed cost, this represents the quantity of units
	// applied.
	FixedPriceQuantity param.Field[float64] `json:"fixed_price_quantity"`
	// The property used to group this price on an invoice
	InvoiceGroupingKey param.Field[string] `json:"invoice_grouping_key"`
	// Within each billing cycle, specifies the cadence at which invoices are produced.
	// If unspecified, a single invoice is produced per billing cycle.
	InvoicingCycleConfiguration param.Field[PlanExternalPlanIDVersionNewParamsAddPricesPriceNewPlanMaxGroupTieredPackagePriceInvoicingCycleConfiguration] `json:"invoicing_cycle_configuration"`
	// User-specified key/value pairs for the resource. Individual keys can be removed
	// by setting the value to `null`, and the entire metadata mapping can be cleared
	// by setting `metadata` to `null`.
	Metadata param.Field[map[string]string] `json:"metadata"`
}

func (r PlanExternalPlanIDVersionNewParamsAddPricesPriceNewPlanMaxGroupTieredPackagePrice) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r PlanExternalPlanIDVersionNewParamsAddPricesPriceNewPlanMaxGroupTieredPackagePrice) implementsPlanExternalPlanIDVersionNewParamsAddPricesPriceUnion() {
}

// The cadence to bill for this price on.
type PlanExternalPlanIDVersionNewParamsAddPricesPriceNewPlanMaxGroupTieredPackagePriceCadence string

const (
	PlanExternalPlanIDVersionNewParamsAddPricesPriceNewPlanMaxGroupTieredPackagePriceCadenceAnnual     PlanExternalPlanIDVersionNewParamsAddPricesPriceNewPlanMaxGroupTieredPackagePriceCadence = "annual"
	PlanExternalPlanIDVersionNewParamsAddPricesPriceNewPlanMaxGroupTieredPackagePriceCadenceSemiAnnual PlanExternalPlanIDVersionNewParamsAddPricesPriceNewPlanMaxGroupTieredPackagePriceCadence = "semi_annual"
	PlanExternalPlanIDVersionNewParamsAddPricesPriceNewPlanMaxGroupTieredPackagePriceCadenceMonthly    PlanExternalPlanIDVersionNewParamsAddPricesPriceNewPlanMaxGroupTieredPackagePriceCadence = "monthly"
	PlanExternalPlanIDVersionNewParamsAddPricesPriceNewPlanMaxGroupTieredPackagePriceCadenceQuarterly  PlanExternalPlanIDVersionNewParamsAddPricesPriceNewPlanMaxGroupTieredPackagePriceCadence = "quarterly"
	PlanExternalPlanIDVersionNewParamsAddPricesPriceNewPlanMaxGroupTieredPackagePriceCadenceOneTime    PlanExternalPlanIDVersionNewParamsAddPricesPriceNewPlanMaxGroupTieredPackagePriceCadence = "one_time"
	PlanExternalPlanIDVersionNewParamsAddPricesPriceNewPlanMaxGroupTieredPackagePriceCadenceCustom     PlanExternalPlanIDVersionNewParamsAddPricesPriceNewPlanMaxGroupTieredPackagePriceCadence = "custom"
)

func (r PlanExternalPlanIDVersionNewParamsAddPricesPriceNewPlanMaxGroupTieredPackagePriceCadence) IsKnown() bool {
	switch r {
	case PlanExternalPlanIDVersionNewParamsAddPricesPriceNewPlanMaxGroupTieredPackagePriceCadenceAnnual, PlanExternalPlanIDVersionNewParamsAddPricesPriceNewPlanMaxGroupTieredPackagePriceCadenceSemiAnnual, PlanExternalPlanIDVersionNewParamsAddPricesPriceNewPlanMaxGroupTieredPackagePriceCadenceMonthly, PlanExternalPlanIDVersionNewParamsAddPricesPriceNewPlanMaxGroupTieredPackagePriceCadenceQuarterly, PlanExternalPlanIDVersionNewParamsAddPricesPriceNewPlanMaxGroupTieredPackagePriceCadenceOneTime, PlanExternalPlanIDVersionNewParamsAddPricesPriceNewPlanMaxGroupTieredPackagePriceCadenceCustom:
		return true
	}
	return false
}

type PlanExternalPlanIDVersionNewParamsAddPricesPriceNewPlanMaxGroupTieredPackagePriceModelType string

const (
	PlanExternalPlanIDVersionNewParamsAddPricesPriceNewPlanMaxGroupTieredPackagePriceModelTypeMaxGroupTieredPackage PlanExternalPlanIDVersionNewParamsAddPricesPriceNewPlanMaxGroupTieredPackagePriceModelType = "max_group_tiered_package"
)

func (r PlanExternalPlanIDVersionNewParamsAddPricesPriceNewPlanMaxGroupTieredPackagePriceModelType) IsKnown() bool {
	switch r {
	case PlanExternalPlanIDVersionNewParamsAddPricesPriceNewPlanMaxGroupTieredPackagePriceModelTypeMaxGroupTieredPackage:
		return true
	}
	return false
}

// For custom cadence: specifies the duration of the billing period in days or
// months.
type PlanExternalPlanIDVersionNewParamsAddPricesPriceNewPlanMaxGroupTieredPackagePriceBillingCycleConfiguration struct {
	// The duration of the billing period.
	Duration param.Field[int64] `json:"duration,required"`
	// The unit of billing period duration.
	DurationUnit param.Field[PlanExternalPlanIDVersionNewParamsAddPricesPriceNewPlanMaxGroupTieredPackagePriceBillingCycleConfigurationDurationUnit] `json:"duration_unit,required"`
}

func (r PlanExternalPlanIDVersionNewParamsAddPricesPriceNewPlanMaxGroupTieredPackagePriceBillingCycleConfiguration) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// The unit of billing period duration.
type PlanExternalPlanIDVersionNewParamsAddPricesPriceNewPlanMaxGroupTieredPackagePriceBillingCycleConfigurationDurationUnit string

const (
	PlanExternalPlanIDVersionNewParamsAddPricesPriceNewPlanMaxGroupTieredPackagePriceBillingCycleConfigurationDurationUnitDay   PlanExternalPlanIDVersionNewParamsAddPricesPriceNewPlanMaxGroupTieredPackagePriceBillingCycleConfigurationDurationUnit = "day"
	PlanExternalPlanIDVersionNewParamsAddPricesPriceNewPlanMaxGroupTieredPackagePriceBillingCycleConfigurationDurationUnitMonth PlanExternalPlanIDVersionNewParamsAddPricesPriceNewPlanMaxGroupTieredPackagePriceBillingCycleConfigurationDurationUnit = "month"
)

func (r PlanExternalPlanIDVersionNewParamsAddPricesPriceNewPlanMaxGroupTieredPackagePriceBillingCycleConfigurationDurationUnit) IsKnown() bool {
	switch r {
	case PlanExternalPlanIDVersionNewParamsAddPricesPriceNewPlanMaxGroupTieredPackagePriceBillingCycleConfigurationDurationUnitDay, PlanExternalPlanIDVersionNewParamsAddPricesPriceNewPlanMaxGroupTieredPackagePriceBillingCycleConfigurationDurationUnitMonth:
		return true
	}
	return false
}

// For dimensional price: specifies a price group and dimension values
type PlanExternalPlanIDVersionNewParamsAddPricesPriceNewPlanMaxGroupTieredPackagePriceDimensionalPriceConfiguration struct {
	// The list of dimension values matching (in order) the dimensions of the price
	// group
	DimensionValues param.Field[[]string] `json:"dimension_values,required"`
	// The id of the dimensional price group to include this price in
	DimensionalPriceGroupID param.Field[string] `json:"dimensional_price_group_id"`
	// The external id of the dimensional price group to include this price in
	ExternalDimensionalPriceGroupID param.Field[string] `json:"external_dimensional_price_group_id"`
}

func (r PlanExternalPlanIDVersionNewParamsAddPricesPriceNewPlanMaxGroupTieredPackagePriceDimensionalPriceConfiguration) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Within each billing cycle, specifies the cadence at which invoices are produced.
// If unspecified, a single invoice is produced per billing cycle.
type PlanExternalPlanIDVersionNewParamsAddPricesPriceNewPlanMaxGroupTieredPackagePriceInvoicingCycleConfiguration struct {
	// The duration of the billing period.
	Duration param.Field[int64] `json:"duration,required"`
	// The unit of billing period duration.
	DurationUnit param.Field[PlanExternalPlanIDVersionNewParamsAddPricesPriceNewPlanMaxGroupTieredPackagePriceInvoicingCycleConfigurationDurationUnit] `json:"duration_unit,required"`
}

func (r PlanExternalPlanIDVersionNewParamsAddPricesPriceNewPlanMaxGroupTieredPackagePriceInvoicingCycleConfiguration) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// The unit of billing period duration.
type PlanExternalPlanIDVersionNewParamsAddPricesPriceNewPlanMaxGroupTieredPackagePriceInvoicingCycleConfigurationDurationUnit string

const (
	PlanExternalPlanIDVersionNewParamsAddPricesPriceNewPlanMaxGroupTieredPackagePriceInvoicingCycleConfigurationDurationUnitDay   PlanExternalPlanIDVersionNewParamsAddPricesPriceNewPlanMaxGroupTieredPackagePriceInvoicingCycleConfigurationDurationUnit = "day"
	PlanExternalPlanIDVersionNewParamsAddPricesPriceNewPlanMaxGroupTieredPackagePriceInvoicingCycleConfigurationDurationUnitMonth PlanExternalPlanIDVersionNewParamsAddPricesPriceNewPlanMaxGroupTieredPackagePriceInvoicingCycleConfigurationDurationUnit = "month"
)

func (r PlanExternalPlanIDVersionNewParamsAddPricesPriceNewPlanMaxGroupTieredPackagePriceInvoicingCycleConfigurationDurationUnit) IsKnown() bool {
	switch r {
	case PlanExternalPlanIDVersionNewParamsAddPricesPriceNewPlanMaxGroupTieredPackagePriceInvoicingCycleConfigurationDurationUnitDay, PlanExternalPlanIDVersionNewParamsAddPricesPriceNewPlanMaxGroupTieredPackagePriceInvoicingCycleConfigurationDurationUnitMonth:
		return true
	}
	return false
}

type PlanExternalPlanIDVersionNewParamsAddPricesPriceNewPlanScalableMatrixWithUnitPricingPrice struct {
	// The cadence to bill for this price on.
	Cadence param.Field[PlanExternalPlanIDVersionNewParamsAddPricesPriceNewPlanScalableMatrixWithUnitPricingPriceCadence] `json:"cadence,required"`
	// The id of the item the price will be associated with.
	ItemID    param.Field[string]                                                                                             `json:"item_id,required"`
	ModelType param.Field[PlanExternalPlanIDVersionNewParamsAddPricesPriceNewPlanScalableMatrixWithUnitPricingPriceModelType] `json:"model_type,required"`
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
	BillingCycleConfiguration param.Field[PlanExternalPlanIDVersionNewParamsAddPricesPriceNewPlanScalableMatrixWithUnitPricingPriceBillingCycleConfiguration] `json:"billing_cycle_configuration"`
	// The per unit conversion rate of the price currency to the invoicing currency.
	ConversionRate param.Field[float64] `json:"conversion_rate"`
	// An ISO 4217 currency string, or custom pricing unit identifier, in which this
	// price is billed.
	Currency param.Field[string] `json:"currency"`
	// For dimensional price: specifies a price group and dimension values
	DimensionalPriceConfiguration param.Field[PlanExternalPlanIDVersionNewParamsAddPricesPriceNewPlanScalableMatrixWithUnitPricingPriceDimensionalPriceConfiguration] `json:"dimensional_price_configuration"`
	// An alias for the price.
	ExternalPriceID param.Field[string] `json:"external_price_id"`
	// If the Price represents a fixed cost, this represents the quantity of units
	// applied.
	FixedPriceQuantity param.Field[float64] `json:"fixed_price_quantity"`
	// The property used to group this price on an invoice
	InvoiceGroupingKey param.Field[string] `json:"invoice_grouping_key"`
	// Within each billing cycle, specifies the cadence at which invoices are produced.
	// If unspecified, a single invoice is produced per billing cycle.
	InvoicingCycleConfiguration param.Field[PlanExternalPlanIDVersionNewParamsAddPricesPriceNewPlanScalableMatrixWithUnitPricingPriceInvoicingCycleConfiguration] `json:"invoicing_cycle_configuration"`
	// User-specified key/value pairs for the resource. Individual keys can be removed
	// by setting the value to `null`, and the entire metadata mapping can be cleared
	// by setting `metadata` to `null`.
	Metadata param.Field[map[string]string] `json:"metadata"`
}

func (r PlanExternalPlanIDVersionNewParamsAddPricesPriceNewPlanScalableMatrixWithUnitPricingPrice) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r PlanExternalPlanIDVersionNewParamsAddPricesPriceNewPlanScalableMatrixWithUnitPricingPrice) implementsPlanExternalPlanIDVersionNewParamsAddPricesPriceUnion() {
}

// The cadence to bill for this price on.
type PlanExternalPlanIDVersionNewParamsAddPricesPriceNewPlanScalableMatrixWithUnitPricingPriceCadence string

const (
	PlanExternalPlanIDVersionNewParamsAddPricesPriceNewPlanScalableMatrixWithUnitPricingPriceCadenceAnnual     PlanExternalPlanIDVersionNewParamsAddPricesPriceNewPlanScalableMatrixWithUnitPricingPriceCadence = "annual"
	PlanExternalPlanIDVersionNewParamsAddPricesPriceNewPlanScalableMatrixWithUnitPricingPriceCadenceSemiAnnual PlanExternalPlanIDVersionNewParamsAddPricesPriceNewPlanScalableMatrixWithUnitPricingPriceCadence = "semi_annual"
	PlanExternalPlanIDVersionNewParamsAddPricesPriceNewPlanScalableMatrixWithUnitPricingPriceCadenceMonthly    PlanExternalPlanIDVersionNewParamsAddPricesPriceNewPlanScalableMatrixWithUnitPricingPriceCadence = "monthly"
	PlanExternalPlanIDVersionNewParamsAddPricesPriceNewPlanScalableMatrixWithUnitPricingPriceCadenceQuarterly  PlanExternalPlanIDVersionNewParamsAddPricesPriceNewPlanScalableMatrixWithUnitPricingPriceCadence = "quarterly"
	PlanExternalPlanIDVersionNewParamsAddPricesPriceNewPlanScalableMatrixWithUnitPricingPriceCadenceOneTime    PlanExternalPlanIDVersionNewParamsAddPricesPriceNewPlanScalableMatrixWithUnitPricingPriceCadence = "one_time"
	PlanExternalPlanIDVersionNewParamsAddPricesPriceNewPlanScalableMatrixWithUnitPricingPriceCadenceCustom     PlanExternalPlanIDVersionNewParamsAddPricesPriceNewPlanScalableMatrixWithUnitPricingPriceCadence = "custom"
)

func (r PlanExternalPlanIDVersionNewParamsAddPricesPriceNewPlanScalableMatrixWithUnitPricingPriceCadence) IsKnown() bool {
	switch r {
	case PlanExternalPlanIDVersionNewParamsAddPricesPriceNewPlanScalableMatrixWithUnitPricingPriceCadenceAnnual, PlanExternalPlanIDVersionNewParamsAddPricesPriceNewPlanScalableMatrixWithUnitPricingPriceCadenceSemiAnnual, PlanExternalPlanIDVersionNewParamsAddPricesPriceNewPlanScalableMatrixWithUnitPricingPriceCadenceMonthly, PlanExternalPlanIDVersionNewParamsAddPricesPriceNewPlanScalableMatrixWithUnitPricingPriceCadenceQuarterly, PlanExternalPlanIDVersionNewParamsAddPricesPriceNewPlanScalableMatrixWithUnitPricingPriceCadenceOneTime, PlanExternalPlanIDVersionNewParamsAddPricesPriceNewPlanScalableMatrixWithUnitPricingPriceCadenceCustom:
		return true
	}
	return false
}

type PlanExternalPlanIDVersionNewParamsAddPricesPriceNewPlanScalableMatrixWithUnitPricingPriceModelType string

const (
	PlanExternalPlanIDVersionNewParamsAddPricesPriceNewPlanScalableMatrixWithUnitPricingPriceModelTypeScalableMatrixWithUnitPricing PlanExternalPlanIDVersionNewParamsAddPricesPriceNewPlanScalableMatrixWithUnitPricingPriceModelType = "scalable_matrix_with_unit_pricing"
)

func (r PlanExternalPlanIDVersionNewParamsAddPricesPriceNewPlanScalableMatrixWithUnitPricingPriceModelType) IsKnown() bool {
	switch r {
	case PlanExternalPlanIDVersionNewParamsAddPricesPriceNewPlanScalableMatrixWithUnitPricingPriceModelTypeScalableMatrixWithUnitPricing:
		return true
	}
	return false
}

// For custom cadence: specifies the duration of the billing period in days or
// months.
type PlanExternalPlanIDVersionNewParamsAddPricesPriceNewPlanScalableMatrixWithUnitPricingPriceBillingCycleConfiguration struct {
	// The duration of the billing period.
	Duration param.Field[int64] `json:"duration,required"`
	// The unit of billing period duration.
	DurationUnit param.Field[PlanExternalPlanIDVersionNewParamsAddPricesPriceNewPlanScalableMatrixWithUnitPricingPriceBillingCycleConfigurationDurationUnit] `json:"duration_unit,required"`
}

func (r PlanExternalPlanIDVersionNewParamsAddPricesPriceNewPlanScalableMatrixWithUnitPricingPriceBillingCycleConfiguration) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// The unit of billing period duration.
type PlanExternalPlanIDVersionNewParamsAddPricesPriceNewPlanScalableMatrixWithUnitPricingPriceBillingCycleConfigurationDurationUnit string

const (
	PlanExternalPlanIDVersionNewParamsAddPricesPriceNewPlanScalableMatrixWithUnitPricingPriceBillingCycleConfigurationDurationUnitDay   PlanExternalPlanIDVersionNewParamsAddPricesPriceNewPlanScalableMatrixWithUnitPricingPriceBillingCycleConfigurationDurationUnit = "day"
	PlanExternalPlanIDVersionNewParamsAddPricesPriceNewPlanScalableMatrixWithUnitPricingPriceBillingCycleConfigurationDurationUnitMonth PlanExternalPlanIDVersionNewParamsAddPricesPriceNewPlanScalableMatrixWithUnitPricingPriceBillingCycleConfigurationDurationUnit = "month"
)

func (r PlanExternalPlanIDVersionNewParamsAddPricesPriceNewPlanScalableMatrixWithUnitPricingPriceBillingCycleConfigurationDurationUnit) IsKnown() bool {
	switch r {
	case PlanExternalPlanIDVersionNewParamsAddPricesPriceNewPlanScalableMatrixWithUnitPricingPriceBillingCycleConfigurationDurationUnitDay, PlanExternalPlanIDVersionNewParamsAddPricesPriceNewPlanScalableMatrixWithUnitPricingPriceBillingCycleConfigurationDurationUnitMonth:
		return true
	}
	return false
}

// For dimensional price: specifies a price group and dimension values
type PlanExternalPlanIDVersionNewParamsAddPricesPriceNewPlanScalableMatrixWithUnitPricingPriceDimensionalPriceConfiguration struct {
	// The list of dimension values matching (in order) the dimensions of the price
	// group
	DimensionValues param.Field[[]string] `json:"dimension_values,required"`
	// The id of the dimensional price group to include this price in
	DimensionalPriceGroupID param.Field[string] `json:"dimensional_price_group_id"`
	// The external id of the dimensional price group to include this price in
	ExternalDimensionalPriceGroupID param.Field[string] `json:"external_dimensional_price_group_id"`
}

func (r PlanExternalPlanIDVersionNewParamsAddPricesPriceNewPlanScalableMatrixWithUnitPricingPriceDimensionalPriceConfiguration) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Within each billing cycle, specifies the cadence at which invoices are produced.
// If unspecified, a single invoice is produced per billing cycle.
type PlanExternalPlanIDVersionNewParamsAddPricesPriceNewPlanScalableMatrixWithUnitPricingPriceInvoicingCycleConfiguration struct {
	// The duration of the billing period.
	Duration param.Field[int64] `json:"duration,required"`
	// The unit of billing period duration.
	DurationUnit param.Field[PlanExternalPlanIDVersionNewParamsAddPricesPriceNewPlanScalableMatrixWithUnitPricingPriceInvoicingCycleConfigurationDurationUnit] `json:"duration_unit,required"`
}

func (r PlanExternalPlanIDVersionNewParamsAddPricesPriceNewPlanScalableMatrixWithUnitPricingPriceInvoicingCycleConfiguration) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// The unit of billing period duration.
type PlanExternalPlanIDVersionNewParamsAddPricesPriceNewPlanScalableMatrixWithUnitPricingPriceInvoicingCycleConfigurationDurationUnit string

const (
	PlanExternalPlanIDVersionNewParamsAddPricesPriceNewPlanScalableMatrixWithUnitPricingPriceInvoicingCycleConfigurationDurationUnitDay   PlanExternalPlanIDVersionNewParamsAddPricesPriceNewPlanScalableMatrixWithUnitPricingPriceInvoicingCycleConfigurationDurationUnit = "day"
	PlanExternalPlanIDVersionNewParamsAddPricesPriceNewPlanScalableMatrixWithUnitPricingPriceInvoicingCycleConfigurationDurationUnitMonth PlanExternalPlanIDVersionNewParamsAddPricesPriceNewPlanScalableMatrixWithUnitPricingPriceInvoicingCycleConfigurationDurationUnit = "month"
)

func (r PlanExternalPlanIDVersionNewParamsAddPricesPriceNewPlanScalableMatrixWithUnitPricingPriceInvoicingCycleConfigurationDurationUnit) IsKnown() bool {
	switch r {
	case PlanExternalPlanIDVersionNewParamsAddPricesPriceNewPlanScalableMatrixWithUnitPricingPriceInvoicingCycleConfigurationDurationUnitDay, PlanExternalPlanIDVersionNewParamsAddPricesPriceNewPlanScalableMatrixWithUnitPricingPriceInvoicingCycleConfigurationDurationUnitMonth:
		return true
	}
	return false
}

type PlanExternalPlanIDVersionNewParamsAddPricesPriceNewPlanScalableMatrixWithTieredPricingPrice struct {
	// The cadence to bill for this price on.
	Cadence param.Field[PlanExternalPlanIDVersionNewParamsAddPricesPriceNewPlanScalableMatrixWithTieredPricingPriceCadence] `json:"cadence,required"`
	// The id of the item the price will be associated with.
	ItemID    param.Field[string]                                                                                               `json:"item_id,required"`
	ModelType param.Field[PlanExternalPlanIDVersionNewParamsAddPricesPriceNewPlanScalableMatrixWithTieredPricingPriceModelType] `json:"model_type,required"`
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
	BillingCycleConfiguration param.Field[PlanExternalPlanIDVersionNewParamsAddPricesPriceNewPlanScalableMatrixWithTieredPricingPriceBillingCycleConfiguration] `json:"billing_cycle_configuration"`
	// The per unit conversion rate of the price currency to the invoicing currency.
	ConversionRate param.Field[float64] `json:"conversion_rate"`
	// An ISO 4217 currency string, or custom pricing unit identifier, in which this
	// price is billed.
	Currency param.Field[string] `json:"currency"`
	// For dimensional price: specifies a price group and dimension values
	DimensionalPriceConfiguration param.Field[PlanExternalPlanIDVersionNewParamsAddPricesPriceNewPlanScalableMatrixWithTieredPricingPriceDimensionalPriceConfiguration] `json:"dimensional_price_configuration"`
	// An alias for the price.
	ExternalPriceID param.Field[string] `json:"external_price_id"`
	// If the Price represents a fixed cost, this represents the quantity of units
	// applied.
	FixedPriceQuantity param.Field[float64] `json:"fixed_price_quantity"`
	// The property used to group this price on an invoice
	InvoiceGroupingKey param.Field[string] `json:"invoice_grouping_key"`
	// Within each billing cycle, specifies the cadence at which invoices are produced.
	// If unspecified, a single invoice is produced per billing cycle.
	InvoicingCycleConfiguration param.Field[PlanExternalPlanIDVersionNewParamsAddPricesPriceNewPlanScalableMatrixWithTieredPricingPriceInvoicingCycleConfiguration] `json:"invoicing_cycle_configuration"`
	// User-specified key/value pairs for the resource. Individual keys can be removed
	// by setting the value to `null`, and the entire metadata mapping can be cleared
	// by setting `metadata` to `null`.
	Metadata param.Field[map[string]string] `json:"metadata"`
}

func (r PlanExternalPlanIDVersionNewParamsAddPricesPriceNewPlanScalableMatrixWithTieredPricingPrice) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r PlanExternalPlanIDVersionNewParamsAddPricesPriceNewPlanScalableMatrixWithTieredPricingPrice) implementsPlanExternalPlanIDVersionNewParamsAddPricesPriceUnion() {
}

// The cadence to bill for this price on.
type PlanExternalPlanIDVersionNewParamsAddPricesPriceNewPlanScalableMatrixWithTieredPricingPriceCadence string

const (
	PlanExternalPlanIDVersionNewParamsAddPricesPriceNewPlanScalableMatrixWithTieredPricingPriceCadenceAnnual     PlanExternalPlanIDVersionNewParamsAddPricesPriceNewPlanScalableMatrixWithTieredPricingPriceCadence = "annual"
	PlanExternalPlanIDVersionNewParamsAddPricesPriceNewPlanScalableMatrixWithTieredPricingPriceCadenceSemiAnnual PlanExternalPlanIDVersionNewParamsAddPricesPriceNewPlanScalableMatrixWithTieredPricingPriceCadence = "semi_annual"
	PlanExternalPlanIDVersionNewParamsAddPricesPriceNewPlanScalableMatrixWithTieredPricingPriceCadenceMonthly    PlanExternalPlanIDVersionNewParamsAddPricesPriceNewPlanScalableMatrixWithTieredPricingPriceCadence = "monthly"
	PlanExternalPlanIDVersionNewParamsAddPricesPriceNewPlanScalableMatrixWithTieredPricingPriceCadenceQuarterly  PlanExternalPlanIDVersionNewParamsAddPricesPriceNewPlanScalableMatrixWithTieredPricingPriceCadence = "quarterly"
	PlanExternalPlanIDVersionNewParamsAddPricesPriceNewPlanScalableMatrixWithTieredPricingPriceCadenceOneTime    PlanExternalPlanIDVersionNewParamsAddPricesPriceNewPlanScalableMatrixWithTieredPricingPriceCadence = "one_time"
	PlanExternalPlanIDVersionNewParamsAddPricesPriceNewPlanScalableMatrixWithTieredPricingPriceCadenceCustom     PlanExternalPlanIDVersionNewParamsAddPricesPriceNewPlanScalableMatrixWithTieredPricingPriceCadence = "custom"
)

func (r PlanExternalPlanIDVersionNewParamsAddPricesPriceNewPlanScalableMatrixWithTieredPricingPriceCadence) IsKnown() bool {
	switch r {
	case PlanExternalPlanIDVersionNewParamsAddPricesPriceNewPlanScalableMatrixWithTieredPricingPriceCadenceAnnual, PlanExternalPlanIDVersionNewParamsAddPricesPriceNewPlanScalableMatrixWithTieredPricingPriceCadenceSemiAnnual, PlanExternalPlanIDVersionNewParamsAddPricesPriceNewPlanScalableMatrixWithTieredPricingPriceCadenceMonthly, PlanExternalPlanIDVersionNewParamsAddPricesPriceNewPlanScalableMatrixWithTieredPricingPriceCadenceQuarterly, PlanExternalPlanIDVersionNewParamsAddPricesPriceNewPlanScalableMatrixWithTieredPricingPriceCadenceOneTime, PlanExternalPlanIDVersionNewParamsAddPricesPriceNewPlanScalableMatrixWithTieredPricingPriceCadenceCustom:
		return true
	}
	return false
}

type PlanExternalPlanIDVersionNewParamsAddPricesPriceNewPlanScalableMatrixWithTieredPricingPriceModelType string

const (
	PlanExternalPlanIDVersionNewParamsAddPricesPriceNewPlanScalableMatrixWithTieredPricingPriceModelTypeScalableMatrixWithTieredPricing PlanExternalPlanIDVersionNewParamsAddPricesPriceNewPlanScalableMatrixWithTieredPricingPriceModelType = "scalable_matrix_with_tiered_pricing"
)

func (r PlanExternalPlanIDVersionNewParamsAddPricesPriceNewPlanScalableMatrixWithTieredPricingPriceModelType) IsKnown() bool {
	switch r {
	case PlanExternalPlanIDVersionNewParamsAddPricesPriceNewPlanScalableMatrixWithTieredPricingPriceModelTypeScalableMatrixWithTieredPricing:
		return true
	}
	return false
}

// For custom cadence: specifies the duration of the billing period in days or
// months.
type PlanExternalPlanIDVersionNewParamsAddPricesPriceNewPlanScalableMatrixWithTieredPricingPriceBillingCycleConfiguration struct {
	// The duration of the billing period.
	Duration param.Field[int64] `json:"duration,required"`
	// The unit of billing period duration.
	DurationUnit param.Field[PlanExternalPlanIDVersionNewParamsAddPricesPriceNewPlanScalableMatrixWithTieredPricingPriceBillingCycleConfigurationDurationUnit] `json:"duration_unit,required"`
}

func (r PlanExternalPlanIDVersionNewParamsAddPricesPriceNewPlanScalableMatrixWithTieredPricingPriceBillingCycleConfiguration) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// The unit of billing period duration.
type PlanExternalPlanIDVersionNewParamsAddPricesPriceNewPlanScalableMatrixWithTieredPricingPriceBillingCycleConfigurationDurationUnit string

const (
	PlanExternalPlanIDVersionNewParamsAddPricesPriceNewPlanScalableMatrixWithTieredPricingPriceBillingCycleConfigurationDurationUnitDay   PlanExternalPlanIDVersionNewParamsAddPricesPriceNewPlanScalableMatrixWithTieredPricingPriceBillingCycleConfigurationDurationUnit = "day"
	PlanExternalPlanIDVersionNewParamsAddPricesPriceNewPlanScalableMatrixWithTieredPricingPriceBillingCycleConfigurationDurationUnitMonth PlanExternalPlanIDVersionNewParamsAddPricesPriceNewPlanScalableMatrixWithTieredPricingPriceBillingCycleConfigurationDurationUnit = "month"
)

func (r PlanExternalPlanIDVersionNewParamsAddPricesPriceNewPlanScalableMatrixWithTieredPricingPriceBillingCycleConfigurationDurationUnit) IsKnown() bool {
	switch r {
	case PlanExternalPlanIDVersionNewParamsAddPricesPriceNewPlanScalableMatrixWithTieredPricingPriceBillingCycleConfigurationDurationUnitDay, PlanExternalPlanIDVersionNewParamsAddPricesPriceNewPlanScalableMatrixWithTieredPricingPriceBillingCycleConfigurationDurationUnitMonth:
		return true
	}
	return false
}

// For dimensional price: specifies a price group and dimension values
type PlanExternalPlanIDVersionNewParamsAddPricesPriceNewPlanScalableMatrixWithTieredPricingPriceDimensionalPriceConfiguration struct {
	// The list of dimension values matching (in order) the dimensions of the price
	// group
	DimensionValues param.Field[[]string] `json:"dimension_values,required"`
	// The id of the dimensional price group to include this price in
	DimensionalPriceGroupID param.Field[string] `json:"dimensional_price_group_id"`
	// The external id of the dimensional price group to include this price in
	ExternalDimensionalPriceGroupID param.Field[string] `json:"external_dimensional_price_group_id"`
}

func (r PlanExternalPlanIDVersionNewParamsAddPricesPriceNewPlanScalableMatrixWithTieredPricingPriceDimensionalPriceConfiguration) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Within each billing cycle, specifies the cadence at which invoices are produced.
// If unspecified, a single invoice is produced per billing cycle.
type PlanExternalPlanIDVersionNewParamsAddPricesPriceNewPlanScalableMatrixWithTieredPricingPriceInvoicingCycleConfiguration struct {
	// The duration of the billing period.
	Duration param.Field[int64] `json:"duration,required"`
	// The unit of billing period duration.
	DurationUnit param.Field[PlanExternalPlanIDVersionNewParamsAddPricesPriceNewPlanScalableMatrixWithTieredPricingPriceInvoicingCycleConfigurationDurationUnit] `json:"duration_unit,required"`
}

func (r PlanExternalPlanIDVersionNewParamsAddPricesPriceNewPlanScalableMatrixWithTieredPricingPriceInvoicingCycleConfiguration) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// The unit of billing period duration.
type PlanExternalPlanIDVersionNewParamsAddPricesPriceNewPlanScalableMatrixWithTieredPricingPriceInvoicingCycleConfigurationDurationUnit string

const (
	PlanExternalPlanIDVersionNewParamsAddPricesPriceNewPlanScalableMatrixWithTieredPricingPriceInvoicingCycleConfigurationDurationUnitDay   PlanExternalPlanIDVersionNewParamsAddPricesPriceNewPlanScalableMatrixWithTieredPricingPriceInvoicingCycleConfigurationDurationUnit = "day"
	PlanExternalPlanIDVersionNewParamsAddPricesPriceNewPlanScalableMatrixWithTieredPricingPriceInvoicingCycleConfigurationDurationUnitMonth PlanExternalPlanIDVersionNewParamsAddPricesPriceNewPlanScalableMatrixWithTieredPricingPriceInvoicingCycleConfigurationDurationUnit = "month"
)

func (r PlanExternalPlanIDVersionNewParamsAddPricesPriceNewPlanScalableMatrixWithTieredPricingPriceInvoicingCycleConfigurationDurationUnit) IsKnown() bool {
	switch r {
	case PlanExternalPlanIDVersionNewParamsAddPricesPriceNewPlanScalableMatrixWithTieredPricingPriceInvoicingCycleConfigurationDurationUnitDay, PlanExternalPlanIDVersionNewParamsAddPricesPriceNewPlanScalableMatrixWithTieredPricingPriceInvoicingCycleConfigurationDurationUnitMonth:
		return true
	}
	return false
}

type PlanExternalPlanIDVersionNewParamsAddPricesPriceNewPlanCumulativeGroupedBulkPrice struct {
	// The cadence to bill for this price on.
	Cadence                     param.Field[PlanExternalPlanIDVersionNewParamsAddPricesPriceNewPlanCumulativeGroupedBulkPriceCadence] `json:"cadence,required"`
	CumulativeGroupedBulkConfig param.Field[map[string]interface{}]                                                                   `json:"cumulative_grouped_bulk_config,required"`
	// The id of the item the price will be associated with.
	ItemID    param.Field[string]                                                                                     `json:"item_id,required"`
	ModelType param.Field[PlanExternalPlanIDVersionNewParamsAddPricesPriceNewPlanCumulativeGroupedBulkPriceModelType] `json:"model_type,required"`
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
	BillingCycleConfiguration param.Field[PlanExternalPlanIDVersionNewParamsAddPricesPriceNewPlanCumulativeGroupedBulkPriceBillingCycleConfiguration] `json:"billing_cycle_configuration"`
	// The per unit conversion rate of the price currency to the invoicing currency.
	ConversionRate param.Field[float64] `json:"conversion_rate"`
	// An ISO 4217 currency string, or custom pricing unit identifier, in which this
	// price is billed.
	Currency param.Field[string] `json:"currency"`
	// For dimensional price: specifies a price group and dimension values
	DimensionalPriceConfiguration param.Field[PlanExternalPlanIDVersionNewParamsAddPricesPriceNewPlanCumulativeGroupedBulkPriceDimensionalPriceConfiguration] `json:"dimensional_price_configuration"`
	// An alias for the price.
	ExternalPriceID param.Field[string] `json:"external_price_id"`
	// If the Price represents a fixed cost, this represents the quantity of units
	// applied.
	FixedPriceQuantity param.Field[float64] `json:"fixed_price_quantity"`
	// The property used to group this price on an invoice
	InvoiceGroupingKey param.Field[string] `json:"invoice_grouping_key"`
	// Within each billing cycle, specifies the cadence at which invoices are produced.
	// If unspecified, a single invoice is produced per billing cycle.
	InvoicingCycleConfiguration param.Field[PlanExternalPlanIDVersionNewParamsAddPricesPriceNewPlanCumulativeGroupedBulkPriceInvoicingCycleConfiguration] `json:"invoicing_cycle_configuration"`
	// User-specified key/value pairs for the resource. Individual keys can be removed
	// by setting the value to `null`, and the entire metadata mapping can be cleared
	// by setting `metadata` to `null`.
	Metadata param.Field[map[string]string] `json:"metadata"`
}

func (r PlanExternalPlanIDVersionNewParamsAddPricesPriceNewPlanCumulativeGroupedBulkPrice) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r PlanExternalPlanIDVersionNewParamsAddPricesPriceNewPlanCumulativeGroupedBulkPrice) implementsPlanExternalPlanIDVersionNewParamsAddPricesPriceUnion() {
}

// The cadence to bill for this price on.
type PlanExternalPlanIDVersionNewParamsAddPricesPriceNewPlanCumulativeGroupedBulkPriceCadence string

const (
	PlanExternalPlanIDVersionNewParamsAddPricesPriceNewPlanCumulativeGroupedBulkPriceCadenceAnnual     PlanExternalPlanIDVersionNewParamsAddPricesPriceNewPlanCumulativeGroupedBulkPriceCadence = "annual"
	PlanExternalPlanIDVersionNewParamsAddPricesPriceNewPlanCumulativeGroupedBulkPriceCadenceSemiAnnual PlanExternalPlanIDVersionNewParamsAddPricesPriceNewPlanCumulativeGroupedBulkPriceCadence = "semi_annual"
	PlanExternalPlanIDVersionNewParamsAddPricesPriceNewPlanCumulativeGroupedBulkPriceCadenceMonthly    PlanExternalPlanIDVersionNewParamsAddPricesPriceNewPlanCumulativeGroupedBulkPriceCadence = "monthly"
	PlanExternalPlanIDVersionNewParamsAddPricesPriceNewPlanCumulativeGroupedBulkPriceCadenceQuarterly  PlanExternalPlanIDVersionNewParamsAddPricesPriceNewPlanCumulativeGroupedBulkPriceCadence = "quarterly"
	PlanExternalPlanIDVersionNewParamsAddPricesPriceNewPlanCumulativeGroupedBulkPriceCadenceOneTime    PlanExternalPlanIDVersionNewParamsAddPricesPriceNewPlanCumulativeGroupedBulkPriceCadence = "one_time"
	PlanExternalPlanIDVersionNewParamsAddPricesPriceNewPlanCumulativeGroupedBulkPriceCadenceCustom     PlanExternalPlanIDVersionNewParamsAddPricesPriceNewPlanCumulativeGroupedBulkPriceCadence = "custom"
)

func (r PlanExternalPlanIDVersionNewParamsAddPricesPriceNewPlanCumulativeGroupedBulkPriceCadence) IsKnown() bool {
	switch r {
	case PlanExternalPlanIDVersionNewParamsAddPricesPriceNewPlanCumulativeGroupedBulkPriceCadenceAnnual, PlanExternalPlanIDVersionNewParamsAddPricesPriceNewPlanCumulativeGroupedBulkPriceCadenceSemiAnnual, PlanExternalPlanIDVersionNewParamsAddPricesPriceNewPlanCumulativeGroupedBulkPriceCadenceMonthly, PlanExternalPlanIDVersionNewParamsAddPricesPriceNewPlanCumulativeGroupedBulkPriceCadenceQuarterly, PlanExternalPlanIDVersionNewParamsAddPricesPriceNewPlanCumulativeGroupedBulkPriceCadenceOneTime, PlanExternalPlanIDVersionNewParamsAddPricesPriceNewPlanCumulativeGroupedBulkPriceCadenceCustom:
		return true
	}
	return false
}

type PlanExternalPlanIDVersionNewParamsAddPricesPriceNewPlanCumulativeGroupedBulkPriceModelType string

const (
	PlanExternalPlanIDVersionNewParamsAddPricesPriceNewPlanCumulativeGroupedBulkPriceModelTypeCumulativeGroupedBulk PlanExternalPlanIDVersionNewParamsAddPricesPriceNewPlanCumulativeGroupedBulkPriceModelType = "cumulative_grouped_bulk"
)

func (r PlanExternalPlanIDVersionNewParamsAddPricesPriceNewPlanCumulativeGroupedBulkPriceModelType) IsKnown() bool {
	switch r {
	case PlanExternalPlanIDVersionNewParamsAddPricesPriceNewPlanCumulativeGroupedBulkPriceModelTypeCumulativeGroupedBulk:
		return true
	}
	return false
}

// For custom cadence: specifies the duration of the billing period in days or
// months.
type PlanExternalPlanIDVersionNewParamsAddPricesPriceNewPlanCumulativeGroupedBulkPriceBillingCycleConfiguration struct {
	// The duration of the billing period.
	Duration param.Field[int64] `json:"duration,required"`
	// The unit of billing period duration.
	DurationUnit param.Field[PlanExternalPlanIDVersionNewParamsAddPricesPriceNewPlanCumulativeGroupedBulkPriceBillingCycleConfigurationDurationUnit] `json:"duration_unit,required"`
}

func (r PlanExternalPlanIDVersionNewParamsAddPricesPriceNewPlanCumulativeGroupedBulkPriceBillingCycleConfiguration) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// The unit of billing period duration.
type PlanExternalPlanIDVersionNewParamsAddPricesPriceNewPlanCumulativeGroupedBulkPriceBillingCycleConfigurationDurationUnit string

const (
	PlanExternalPlanIDVersionNewParamsAddPricesPriceNewPlanCumulativeGroupedBulkPriceBillingCycleConfigurationDurationUnitDay   PlanExternalPlanIDVersionNewParamsAddPricesPriceNewPlanCumulativeGroupedBulkPriceBillingCycleConfigurationDurationUnit = "day"
	PlanExternalPlanIDVersionNewParamsAddPricesPriceNewPlanCumulativeGroupedBulkPriceBillingCycleConfigurationDurationUnitMonth PlanExternalPlanIDVersionNewParamsAddPricesPriceNewPlanCumulativeGroupedBulkPriceBillingCycleConfigurationDurationUnit = "month"
)

func (r PlanExternalPlanIDVersionNewParamsAddPricesPriceNewPlanCumulativeGroupedBulkPriceBillingCycleConfigurationDurationUnit) IsKnown() bool {
	switch r {
	case PlanExternalPlanIDVersionNewParamsAddPricesPriceNewPlanCumulativeGroupedBulkPriceBillingCycleConfigurationDurationUnitDay, PlanExternalPlanIDVersionNewParamsAddPricesPriceNewPlanCumulativeGroupedBulkPriceBillingCycleConfigurationDurationUnitMonth:
		return true
	}
	return false
}

// For dimensional price: specifies a price group and dimension values
type PlanExternalPlanIDVersionNewParamsAddPricesPriceNewPlanCumulativeGroupedBulkPriceDimensionalPriceConfiguration struct {
	// The list of dimension values matching (in order) the dimensions of the price
	// group
	DimensionValues param.Field[[]string] `json:"dimension_values,required"`
	// The id of the dimensional price group to include this price in
	DimensionalPriceGroupID param.Field[string] `json:"dimensional_price_group_id"`
	// The external id of the dimensional price group to include this price in
	ExternalDimensionalPriceGroupID param.Field[string] `json:"external_dimensional_price_group_id"`
}

func (r PlanExternalPlanIDVersionNewParamsAddPricesPriceNewPlanCumulativeGroupedBulkPriceDimensionalPriceConfiguration) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Within each billing cycle, specifies the cadence at which invoices are produced.
// If unspecified, a single invoice is produced per billing cycle.
type PlanExternalPlanIDVersionNewParamsAddPricesPriceNewPlanCumulativeGroupedBulkPriceInvoicingCycleConfiguration struct {
	// The duration of the billing period.
	Duration param.Field[int64] `json:"duration,required"`
	// The unit of billing period duration.
	DurationUnit param.Field[PlanExternalPlanIDVersionNewParamsAddPricesPriceNewPlanCumulativeGroupedBulkPriceInvoicingCycleConfigurationDurationUnit] `json:"duration_unit,required"`
}

func (r PlanExternalPlanIDVersionNewParamsAddPricesPriceNewPlanCumulativeGroupedBulkPriceInvoicingCycleConfiguration) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// The unit of billing period duration.
type PlanExternalPlanIDVersionNewParamsAddPricesPriceNewPlanCumulativeGroupedBulkPriceInvoicingCycleConfigurationDurationUnit string

const (
	PlanExternalPlanIDVersionNewParamsAddPricesPriceNewPlanCumulativeGroupedBulkPriceInvoicingCycleConfigurationDurationUnitDay   PlanExternalPlanIDVersionNewParamsAddPricesPriceNewPlanCumulativeGroupedBulkPriceInvoicingCycleConfigurationDurationUnit = "day"
	PlanExternalPlanIDVersionNewParamsAddPricesPriceNewPlanCumulativeGroupedBulkPriceInvoicingCycleConfigurationDurationUnitMonth PlanExternalPlanIDVersionNewParamsAddPricesPriceNewPlanCumulativeGroupedBulkPriceInvoicingCycleConfigurationDurationUnit = "month"
)

func (r PlanExternalPlanIDVersionNewParamsAddPricesPriceNewPlanCumulativeGroupedBulkPriceInvoicingCycleConfigurationDurationUnit) IsKnown() bool {
	switch r {
	case PlanExternalPlanIDVersionNewParamsAddPricesPriceNewPlanCumulativeGroupedBulkPriceInvoicingCycleConfigurationDurationUnitDay, PlanExternalPlanIDVersionNewParamsAddPricesPriceNewPlanCumulativeGroupedBulkPriceInvoicingCycleConfigurationDurationUnitMonth:
		return true
	}
	return false
}

type PlanExternalPlanIDVersionNewParamsAddPricesPriceNewPlanTieredPackageWithMinimumPrice struct {
	// The cadence to bill for this price on.
	Cadence param.Field[PlanExternalPlanIDVersionNewParamsAddPricesPriceNewPlanTieredPackageWithMinimumPriceCadence] `json:"cadence,required"`
	// The id of the item the price will be associated with.
	ItemID    param.Field[string]                                                                                        `json:"item_id,required"`
	ModelType param.Field[PlanExternalPlanIDVersionNewParamsAddPricesPriceNewPlanTieredPackageWithMinimumPriceModelType] `json:"model_type,required"`
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
	BillingCycleConfiguration param.Field[PlanExternalPlanIDVersionNewParamsAddPricesPriceNewPlanTieredPackageWithMinimumPriceBillingCycleConfiguration] `json:"billing_cycle_configuration"`
	// The per unit conversion rate of the price currency to the invoicing currency.
	ConversionRate param.Field[float64] `json:"conversion_rate"`
	// An ISO 4217 currency string, or custom pricing unit identifier, in which this
	// price is billed.
	Currency param.Field[string] `json:"currency"`
	// For dimensional price: specifies a price group and dimension values
	DimensionalPriceConfiguration param.Field[PlanExternalPlanIDVersionNewParamsAddPricesPriceNewPlanTieredPackageWithMinimumPriceDimensionalPriceConfiguration] `json:"dimensional_price_configuration"`
	// An alias for the price.
	ExternalPriceID param.Field[string] `json:"external_price_id"`
	// If the Price represents a fixed cost, this represents the quantity of units
	// applied.
	FixedPriceQuantity param.Field[float64] `json:"fixed_price_quantity"`
	// The property used to group this price on an invoice
	InvoiceGroupingKey param.Field[string] `json:"invoice_grouping_key"`
	// Within each billing cycle, specifies the cadence at which invoices are produced.
	// If unspecified, a single invoice is produced per billing cycle.
	InvoicingCycleConfiguration param.Field[PlanExternalPlanIDVersionNewParamsAddPricesPriceNewPlanTieredPackageWithMinimumPriceInvoicingCycleConfiguration] `json:"invoicing_cycle_configuration"`
	// User-specified key/value pairs for the resource. Individual keys can be removed
	// by setting the value to `null`, and the entire metadata mapping can be cleared
	// by setting `metadata` to `null`.
	Metadata param.Field[map[string]string] `json:"metadata"`
}

func (r PlanExternalPlanIDVersionNewParamsAddPricesPriceNewPlanTieredPackageWithMinimumPrice) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r PlanExternalPlanIDVersionNewParamsAddPricesPriceNewPlanTieredPackageWithMinimumPrice) implementsPlanExternalPlanIDVersionNewParamsAddPricesPriceUnion() {
}

// The cadence to bill for this price on.
type PlanExternalPlanIDVersionNewParamsAddPricesPriceNewPlanTieredPackageWithMinimumPriceCadence string

const (
	PlanExternalPlanIDVersionNewParamsAddPricesPriceNewPlanTieredPackageWithMinimumPriceCadenceAnnual     PlanExternalPlanIDVersionNewParamsAddPricesPriceNewPlanTieredPackageWithMinimumPriceCadence = "annual"
	PlanExternalPlanIDVersionNewParamsAddPricesPriceNewPlanTieredPackageWithMinimumPriceCadenceSemiAnnual PlanExternalPlanIDVersionNewParamsAddPricesPriceNewPlanTieredPackageWithMinimumPriceCadence = "semi_annual"
	PlanExternalPlanIDVersionNewParamsAddPricesPriceNewPlanTieredPackageWithMinimumPriceCadenceMonthly    PlanExternalPlanIDVersionNewParamsAddPricesPriceNewPlanTieredPackageWithMinimumPriceCadence = "monthly"
	PlanExternalPlanIDVersionNewParamsAddPricesPriceNewPlanTieredPackageWithMinimumPriceCadenceQuarterly  PlanExternalPlanIDVersionNewParamsAddPricesPriceNewPlanTieredPackageWithMinimumPriceCadence = "quarterly"
	PlanExternalPlanIDVersionNewParamsAddPricesPriceNewPlanTieredPackageWithMinimumPriceCadenceOneTime    PlanExternalPlanIDVersionNewParamsAddPricesPriceNewPlanTieredPackageWithMinimumPriceCadence = "one_time"
	PlanExternalPlanIDVersionNewParamsAddPricesPriceNewPlanTieredPackageWithMinimumPriceCadenceCustom     PlanExternalPlanIDVersionNewParamsAddPricesPriceNewPlanTieredPackageWithMinimumPriceCadence = "custom"
)

func (r PlanExternalPlanIDVersionNewParamsAddPricesPriceNewPlanTieredPackageWithMinimumPriceCadence) IsKnown() bool {
	switch r {
	case PlanExternalPlanIDVersionNewParamsAddPricesPriceNewPlanTieredPackageWithMinimumPriceCadenceAnnual, PlanExternalPlanIDVersionNewParamsAddPricesPriceNewPlanTieredPackageWithMinimumPriceCadenceSemiAnnual, PlanExternalPlanIDVersionNewParamsAddPricesPriceNewPlanTieredPackageWithMinimumPriceCadenceMonthly, PlanExternalPlanIDVersionNewParamsAddPricesPriceNewPlanTieredPackageWithMinimumPriceCadenceQuarterly, PlanExternalPlanIDVersionNewParamsAddPricesPriceNewPlanTieredPackageWithMinimumPriceCadenceOneTime, PlanExternalPlanIDVersionNewParamsAddPricesPriceNewPlanTieredPackageWithMinimumPriceCadenceCustom:
		return true
	}
	return false
}

type PlanExternalPlanIDVersionNewParamsAddPricesPriceNewPlanTieredPackageWithMinimumPriceModelType string

const (
	PlanExternalPlanIDVersionNewParamsAddPricesPriceNewPlanTieredPackageWithMinimumPriceModelTypeTieredPackageWithMinimum PlanExternalPlanIDVersionNewParamsAddPricesPriceNewPlanTieredPackageWithMinimumPriceModelType = "tiered_package_with_minimum"
)

func (r PlanExternalPlanIDVersionNewParamsAddPricesPriceNewPlanTieredPackageWithMinimumPriceModelType) IsKnown() bool {
	switch r {
	case PlanExternalPlanIDVersionNewParamsAddPricesPriceNewPlanTieredPackageWithMinimumPriceModelTypeTieredPackageWithMinimum:
		return true
	}
	return false
}

// For custom cadence: specifies the duration of the billing period in days or
// months.
type PlanExternalPlanIDVersionNewParamsAddPricesPriceNewPlanTieredPackageWithMinimumPriceBillingCycleConfiguration struct {
	// The duration of the billing period.
	Duration param.Field[int64] `json:"duration,required"`
	// The unit of billing period duration.
	DurationUnit param.Field[PlanExternalPlanIDVersionNewParamsAddPricesPriceNewPlanTieredPackageWithMinimumPriceBillingCycleConfigurationDurationUnit] `json:"duration_unit,required"`
}

func (r PlanExternalPlanIDVersionNewParamsAddPricesPriceNewPlanTieredPackageWithMinimumPriceBillingCycleConfiguration) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// The unit of billing period duration.
type PlanExternalPlanIDVersionNewParamsAddPricesPriceNewPlanTieredPackageWithMinimumPriceBillingCycleConfigurationDurationUnit string

const (
	PlanExternalPlanIDVersionNewParamsAddPricesPriceNewPlanTieredPackageWithMinimumPriceBillingCycleConfigurationDurationUnitDay   PlanExternalPlanIDVersionNewParamsAddPricesPriceNewPlanTieredPackageWithMinimumPriceBillingCycleConfigurationDurationUnit = "day"
	PlanExternalPlanIDVersionNewParamsAddPricesPriceNewPlanTieredPackageWithMinimumPriceBillingCycleConfigurationDurationUnitMonth PlanExternalPlanIDVersionNewParamsAddPricesPriceNewPlanTieredPackageWithMinimumPriceBillingCycleConfigurationDurationUnit = "month"
)

func (r PlanExternalPlanIDVersionNewParamsAddPricesPriceNewPlanTieredPackageWithMinimumPriceBillingCycleConfigurationDurationUnit) IsKnown() bool {
	switch r {
	case PlanExternalPlanIDVersionNewParamsAddPricesPriceNewPlanTieredPackageWithMinimumPriceBillingCycleConfigurationDurationUnitDay, PlanExternalPlanIDVersionNewParamsAddPricesPriceNewPlanTieredPackageWithMinimumPriceBillingCycleConfigurationDurationUnitMonth:
		return true
	}
	return false
}

// For dimensional price: specifies a price group and dimension values
type PlanExternalPlanIDVersionNewParamsAddPricesPriceNewPlanTieredPackageWithMinimumPriceDimensionalPriceConfiguration struct {
	// The list of dimension values matching (in order) the dimensions of the price
	// group
	DimensionValues param.Field[[]string] `json:"dimension_values,required"`
	// The id of the dimensional price group to include this price in
	DimensionalPriceGroupID param.Field[string] `json:"dimensional_price_group_id"`
	// The external id of the dimensional price group to include this price in
	ExternalDimensionalPriceGroupID param.Field[string] `json:"external_dimensional_price_group_id"`
}

func (r PlanExternalPlanIDVersionNewParamsAddPricesPriceNewPlanTieredPackageWithMinimumPriceDimensionalPriceConfiguration) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Within each billing cycle, specifies the cadence at which invoices are produced.
// If unspecified, a single invoice is produced per billing cycle.
type PlanExternalPlanIDVersionNewParamsAddPricesPriceNewPlanTieredPackageWithMinimumPriceInvoicingCycleConfiguration struct {
	// The duration of the billing period.
	Duration param.Field[int64] `json:"duration,required"`
	// The unit of billing period duration.
	DurationUnit param.Field[PlanExternalPlanIDVersionNewParamsAddPricesPriceNewPlanTieredPackageWithMinimumPriceInvoicingCycleConfigurationDurationUnit] `json:"duration_unit,required"`
}

func (r PlanExternalPlanIDVersionNewParamsAddPricesPriceNewPlanTieredPackageWithMinimumPriceInvoicingCycleConfiguration) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// The unit of billing period duration.
type PlanExternalPlanIDVersionNewParamsAddPricesPriceNewPlanTieredPackageWithMinimumPriceInvoicingCycleConfigurationDurationUnit string

const (
	PlanExternalPlanIDVersionNewParamsAddPricesPriceNewPlanTieredPackageWithMinimumPriceInvoicingCycleConfigurationDurationUnitDay   PlanExternalPlanIDVersionNewParamsAddPricesPriceNewPlanTieredPackageWithMinimumPriceInvoicingCycleConfigurationDurationUnit = "day"
	PlanExternalPlanIDVersionNewParamsAddPricesPriceNewPlanTieredPackageWithMinimumPriceInvoicingCycleConfigurationDurationUnitMonth PlanExternalPlanIDVersionNewParamsAddPricesPriceNewPlanTieredPackageWithMinimumPriceInvoicingCycleConfigurationDurationUnit = "month"
)

func (r PlanExternalPlanIDVersionNewParamsAddPricesPriceNewPlanTieredPackageWithMinimumPriceInvoicingCycleConfigurationDurationUnit) IsKnown() bool {
	switch r {
	case PlanExternalPlanIDVersionNewParamsAddPricesPriceNewPlanTieredPackageWithMinimumPriceInvoicingCycleConfigurationDurationUnitDay, PlanExternalPlanIDVersionNewParamsAddPricesPriceNewPlanTieredPackageWithMinimumPriceInvoicingCycleConfigurationDurationUnitMonth:
		return true
	}
	return false
}

type PlanExternalPlanIDVersionNewParamsAddPricesPriceNewPlanMatrixWithAllocationPrice struct {
	// The cadence to bill for this price on.
	Cadence param.Field[PlanExternalPlanIDVersionNewParamsAddPricesPriceNewPlanMatrixWithAllocationPriceCadence] `json:"cadence,required"`
	// The id of the item the price will be associated with.
	ItemID                     param.Field[string]                                                                                                     `json:"item_id,required"`
	MatrixWithAllocationConfig param.Field[PlanExternalPlanIDVersionNewParamsAddPricesPriceNewPlanMatrixWithAllocationPriceMatrixWithAllocationConfig] `json:"matrix_with_allocation_config,required"`
	ModelType                  param.Field[PlanExternalPlanIDVersionNewParamsAddPricesPriceNewPlanMatrixWithAllocationPriceModelType]                  `json:"model_type,required"`
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
	BillingCycleConfiguration param.Field[PlanExternalPlanIDVersionNewParamsAddPricesPriceNewPlanMatrixWithAllocationPriceBillingCycleConfiguration] `json:"billing_cycle_configuration"`
	// The per unit conversion rate of the price currency to the invoicing currency.
	ConversionRate param.Field[float64] `json:"conversion_rate"`
	// An ISO 4217 currency string, or custom pricing unit identifier, in which this
	// price is billed.
	Currency param.Field[string] `json:"currency"`
	// For dimensional price: specifies a price group and dimension values
	DimensionalPriceConfiguration param.Field[PlanExternalPlanIDVersionNewParamsAddPricesPriceNewPlanMatrixWithAllocationPriceDimensionalPriceConfiguration] `json:"dimensional_price_configuration"`
	// An alias for the price.
	ExternalPriceID param.Field[string] `json:"external_price_id"`
	// If the Price represents a fixed cost, this represents the quantity of units
	// applied.
	FixedPriceQuantity param.Field[float64] `json:"fixed_price_quantity"`
	// The property used to group this price on an invoice
	InvoiceGroupingKey param.Field[string] `json:"invoice_grouping_key"`
	// Within each billing cycle, specifies the cadence at which invoices are produced.
	// If unspecified, a single invoice is produced per billing cycle.
	InvoicingCycleConfiguration param.Field[PlanExternalPlanIDVersionNewParamsAddPricesPriceNewPlanMatrixWithAllocationPriceInvoicingCycleConfiguration] `json:"invoicing_cycle_configuration"`
	// User-specified key/value pairs for the resource. Individual keys can be removed
	// by setting the value to `null`, and the entire metadata mapping can be cleared
	// by setting `metadata` to `null`.
	Metadata param.Field[map[string]string] `json:"metadata"`
}

func (r PlanExternalPlanIDVersionNewParamsAddPricesPriceNewPlanMatrixWithAllocationPrice) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r PlanExternalPlanIDVersionNewParamsAddPricesPriceNewPlanMatrixWithAllocationPrice) implementsPlanExternalPlanIDVersionNewParamsAddPricesPriceUnion() {
}

// The cadence to bill for this price on.
type PlanExternalPlanIDVersionNewParamsAddPricesPriceNewPlanMatrixWithAllocationPriceCadence string

const (
	PlanExternalPlanIDVersionNewParamsAddPricesPriceNewPlanMatrixWithAllocationPriceCadenceAnnual     PlanExternalPlanIDVersionNewParamsAddPricesPriceNewPlanMatrixWithAllocationPriceCadence = "annual"
	PlanExternalPlanIDVersionNewParamsAddPricesPriceNewPlanMatrixWithAllocationPriceCadenceSemiAnnual PlanExternalPlanIDVersionNewParamsAddPricesPriceNewPlanMatrixWithAllocationPriceCadence = "semi_annual"
	PlanExternalPlanIDVersionNewParamsAddPricesPriceNewPlanMatrixWithAllocationPriceCadenceMonthly    PlanExternalPlanIDVersionNewParamsAddPricesPriceNewPlanMatrixWithAllocationPriceCadence = "monthly"
	PlanExternalPlanIDVersionNewParamsAddPricesPriceNewPlanMatrixWithAllocationPriceCadenceQuarterly  PlanExternalPlanIDVersionNewParamsAddPricesPriceNewPlanMatrixWithAllocationPriceCadence = "quarterly"
	PlanExternalPlanIDVersionNewParamsAddPricesPriceNewPlanMatrixWithAllocationPriceCadenceOneTime    PlanExternalPlanIDVersionNewParamsAddPricesPriceNewPlanMatrixWithAllocationPriceCadence = "one_time"
	PlanExternalPlanIDVersionNewParamsAddPricesPriceNewPlanMatrixWithAllocationPriceCadenceCustom     PlanExternalPlanIDVersionNewParamsAddPricesPriceNewPlanMatrixWithAllocationPriceCadence = "custom"
)

func (r PlanExternalPlanIDVersionNewParamsAddPricesPriceNewPlanMatrixWithAllocationPriceCadence) IsKnown() bool {
	switch r {
	case PlanExternalPlanIDVersionNewParamsAddPricesPriceNewPlanMatrixWithAllocationPriceCadenceAnnual, PlanExternalPlanIDVersionNewParamsAddPricesPriceNewPlanMatrixWithAllocationPriceCadenceSemiAnnual, PlanExternalPlanIDVersionNewParamsAddPricesPriceNewPlanMatrixWithAllocationPriceCadenceMonthly, PlanExternalPlanIDVersionNewParamsAddPricesPriceNewPlanMatrixWithAllocationPriceCadenceQuarterly, PlanExternalPlanIDVersionNewParamsAddPricesPriceNewPlanMatrixWithAllocationPriceCadenceOneTime, PlanExternalPlanIDVersionNewParamsAddPricesPriceNewPlanMatrixWithAllocationPriceCadenceCustom:
		return true
	}
	return false
}

type PlanExternalPlanIDVersionNewParamsAddPricesPriceNewPlanMatrixWithAllocationPriceMatrixWithAllocationConfig struct {
	// Allocation to be used to calculate the price
	Allocation param.Field[float64] `json:"allocation,required"`
	// Default per unit rate for any usage not bucketed into a specified matrix_value
	DefaultUnitAmount param.Field[string] `json:"default_unit_amount,required"`
	// One or two event property values to evaluate matrix groups by
	Dimensions param.Field[[]string] `json:"dimensions,required"`
	// Matrix values for specified matrix grouping keys
	MatrixValues param.Field[[]PlanExternalPlanIDVersionNewParamsAddPricesPriceNewPlanMatrixWithAllocationPriceMatrixWithAllocationConfigMatrixValue] `json:"matrix_values,required"`
}

func (r PlanExternalPlanIDVersionNewParamsAddPricesPriceNewPlanMatrixWithAllocationPriceMatrixWithAllocationConfig) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type PlanExternalPlanIDVersionNewParamsAddPricesPriceNewPlanMatrixWithAllocationPriceMatrixWithAllocationConfigMatrixValue struct {
	// One or two matrix keys to filter usage to this Matrix value by. For example,
	// ["region", "tier"] could be used to filter cloud usage by a cloud region and an
	// instance tier.
	DimensionValues param.Field[[]string] `json:"dimension_values,required"`
	// Unit price for the specified dimension_values
	UnitAmount param.Field[string] `json:"unit_amount,required"`
}

func (r PlanExternalPlanIDVersionNewParamsAddPricesPriceNewPlanMatrixWithAllocationPriceMatrixWithAllocationConfigMatrixValue) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type PlanExternalPlanIDVersionNewParamsAddPricesPriceNewPlanMatrixWithAllocationPriceModelType string

const (
	PlanExternalPlanIDVersionNewParamsAddPricesPriceNewPlanMatrixWithAllocationPriceModelTypeMatrixWithAllocation PlanExternalPlanIDVersionNewParamsAddPricesPriceNewPlanMatrixWithAllocationPriceModelType = "matrix_with_allocation"
)

func (r PlanExternalPlanIDVersionNewParamsAddPricesPriceNewPlanMatrixWithAllocationPriceModelType) IsKnown() bool {
	switch r {
	case PlanExternalPlanIDVersionNewParamsAddPricesPriceNewPlanMatrixWithAllocationPriceModelTypeMatrixWithAllocation:
		return true
	}
	return false
}

// For custom cadence: specifies the duration of the billing period in days or
// months.
type PlanExternalPlanIDVersionNewParamsAddPricesPriceNewPlanMatrixWithAllocationPriceBillingCycleConfiguration struct {
	// The duration of the billing period.
	Duration param.Field[int64] `json:"duration,required"`
	// The unit of billing period duration.
	DurationUnit param.Field[PlanExternalPlanIDVersionNewParamsAddPricesPriceNewPlanMatrixWithAllocationPriceBillingCycleConfigurationDurationUnit] `json:"duration_unit,required"`
}

func (r PlanExternalPlanIDVersionNewParamsAddPricesPriceNewPlanMatrixWithAllocationPriceBillingCycleConfiguration) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// The unit of billing period duration.
type PlanExternalPlanIDVersionNewParamsAddPricesPriceNewPlanMatrixWithAllocationPriceBillingCycleConfigurationDurationUnit string

const (
	PlanExternalPlanIDVersionNewParamsAddPricesPriceNewPlanMatrixWithAllocationPriceBillingCycleConfigurationDurationUnitDay   PlanExternalPlanIDVersionNewParamsAddPricesPriceNewPlanMatrixWithAllocationPriceBillingCycleConfigurationDurationUnit = "day"
	PlanExternalPlanIDVersionNewParamsAddPricesPriceNewPlanMatrixWithAllocationPriceBillingCycleConfigurationDurationUnitMonth PlanExternalPlanIDVersionNewParamsAddPricesPriceNewPlanMatrixWithAllocationPriceBillingCycleConfigurationDurationUnit = "month"
)

func (r PlanExternalPlanIDVersionNewParamsAddPricesPriceNewPlanMatrixWithAllocationPriceBillingCycleConfigurationDurationUnit) IsKnown() bool {
	switch r {
	case PlanExternalPlanIDVersionNewParamsAddPricesPriceNewPlanMatrixWithAllocationPriceBillingCycleConfigurationDurationUnitDay, PlanExternalPlanIDVersionNewParamsAddPricesPriceNewPlanMatrixWithAllocationPriceBillingCycleConfigurationDurationUnitMonth:
		return true
	}
	return false
}

// For dimensional price: specifies a price group and dimension values
type PlanExternalPlanIDVersionNewParamsAddPricesPriceNewPlanMatrixWithAllocationPriceDimensionalPriceConfiguration struct {
	// The list of dimension values matching (in order) the dimensions of the price
	// group
	DimensionValues param.Field[[]string] `json:"dimension_values,required"`
	// The id of the dimensional price group to include this price in
	DimensionalPriceGroupID param.Field[string] `json:"dimensional_price_group_id"`
	// The external id of the dimensional price group to include this price in
	ExternalDimensionalPriceGroupID param.Field[string] `json:"external_dimensional_price_group_id"`
}

func (r PlanExternalPlanIDVersionNewParamsAddPricesPriceNewPlanMatrixWithAllocationPriceDimensionalPriceConfiguration) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Within each billing cycle, specifies the cadence at which invoices are produced.
// If unspecified, a single invoice is produced per billing cycle.
type PlanExternalPlanIDVersionNewParamsAddPricesPriceNewPlanMatrixWithAllocationPriceInvoicingCycleConfiguration struct {
	// The duration of the billing period.
	Duration param.Field[int64] `json:"duration,required"`
	// The unit of billing period duration.
	DurationUnit param.Field[PlanExternalPlanIDVersionNewParamsAddPricesPriceNewPlanMatrixWithAllocationPriceInvoicingCycleConfigurationDurationUnit] `json:"duration_unit,required"`
}

func (r PlanExternalPlanIDVersionNewParamsAddPricesPriceNewPlanMatrixWithAllocationPriceInvoicingCycleConfiguration) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// The unit of billing period duration.
type PlanExternalPlanIDVersionNewParamsAddPricesPriceNewPlanMatrixWithAllocationPriceInvoicingCycleConfigurationDurationUnit string

const (
	PlanExternalPlanIDVersionNewParamsAddPricesPriceNewPlanMatrixWithAllocationPriceInvoicingCycleConfigurationDurationUnitDay   PlanExternalPlanIDVersionNewParamsAddPricesPriceNewPlanMatrixWithAllocationPriceInvoicingCycleConfigurationDurationUnit = "day"
	PlanExternalPlanIDVersionNewParamsAddPricesPriceNewPlanMatrixWithAllocationPriceInvoicingCycleConfigurationDurationUnitMonth PlanExternalPlanIDVersionNewParamsAddPricesPriceNewPlanMatrixWithAllocationPriceInvoicingCycleConfigurationDurationUnit = "month"
)

func (r PlanExternalPlanIDVersionNewParamsAddPricesPriceNewPlanMatrixWithAllocationPriceInvoicingCycleConfigurationDurationUnit) IsKnown() bool {
	switch r {
	case PlanExternalPlanIDVersionNewParamsAddPricesPriceNewPlanMatrixWithAllocationPriceInvoicingCycleConfigurationDurationUnitDay, PlanExternalPlanIDVersionNewParamsAddPricesPriceNewPlanMatrixWithAllocationPriceInvoicingCycleConfigurationDurationUnitMonth:
		return true
	}
	return false
}

type PlanExternalPlanIDVersionNewParamsAddPricesPriceNewPlanGroupedTieredPrice struct {
	// The cadence to bill for this price on.
	Cadence             param.Field[PlanExternalPlanIDVersionNewParamsAddPricesPriceNewPlanGroupedTieredPriceCadence] `json:"cadence,required"`
	GroupedTieredConfig param.Field[map[string]interface{}]                                                           `json:"grouped_tiered_config,required"`
	// The id of the item the price will be associated with.
	ItemID    param.Field[string]                                                                             `json:"item_id,required"`
	ModelType param.Field[PlanExternalPlanIDVersionNewParamsAddPricesPriceNewPlanGroupedTieredPriceModelType] `json:"model_type,required"`
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
	BillingCycleConfiguration param.Field[PlanExternalPlanIDVersionNewParamsAddPricesPriceNewPlanGroupedTieredPriceBillingCycleConfiguration] `json:"billing_cycle_configuration"`
	// The per unit conversion rate of the price currency to the invoicing currency.
	ConversionRate param.Field[float64] `json:"conversion_rate"`
	// An ISO 4217 currency string, or custom pricing unit identifier, in which this
	// price is billed.
	Currency param.Field[string] `json:"currency"`
	// For dimensional price: specifies a price group and dimension values
	DimensionalPriceConfiguration param.Field[PlanExternalPlanIDVersionNewParamsAddPricesPriceNewPlanGroupedTieredPriceDimensionalPriceConfiguration] `json:"dimensional_price_configuration"`
	// An alias for the price.
	ExternalPriceID param.Field[string] `json:"external_price_id"`
	// If the Price represents a fixed cost, this represents the quantity of units
	// applied.
	FixedPriceQuantity param.Field[float64] `json:"fixed_price_quantity"`
	// The property used to group this price on an invoice
	InvoiceGroupingKey param.Field[string] `json:"invoice_grouping_key"`
	// Within each billing cycle, specifies the cadence at which invoices are produced.
	// If unspecified, a single invoice is produced per billing cycle.
	InvoicingCycleConfiguration param.Field[PlanExternalPlanIDVersionNewParamsAddPricesPriceNewPlanGroupedTieredPriceInvoicingCycleConfiguration] `json:"invoicing_cycle_configuration"`
	// User-specified key/value pairs for the resource. Individual keys can be removed
	// by setting the value to `null`, and the entire metadata mapping can be cleared
	// by setting `metadata` to `null`.
	Metadata param.Field[map[string]string] `json:"metadata"`
}

func (r PlanExternalPlanIDVersionNewParamsAddPricesPriceNewPlanGroupedTieredPrice) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r PlanExternalPlanIDVersionNewParamsAddPricesPriceNewPlanGroupedTieredPrice) implementsPlanExternalPlanIDVersionNewParamsAddPricesPriceUnion() {
}

// The cadence to bill for this price on.
type PlanExternalPlanIDVersionNewParamsAddPricesPriceNewPlanGroupedTieredPriceCadence string

const (
	PlanExternalPlanIDVersionNewParamsAddPricesPriceNewPlanGroupedTieredPriceCadenceAnnual     PlanExternalPlanIDVersionNewParamsAddPricesPriceNewPlanGroupedTieredPriceCadence = "annual"
	PlanExternalPlanIDVersionNewParamsAddPricesPriceNewPlanGroupedTieredPriceCadenceSemiAnnual PlanExternalPlanIDVersionNewParamsAddPricesPriceNewPlanGroupedTieredPriceCadence = "semi_annual"
	PlanExternalPlanIDVersionNewParamsAddPricesPriceNewPlanGroupedTieredPriceCadenceMonthly    PlanExternalPlanIDVersionNewParamsAddPricesPriceNewPlanGroupedTieredPriceCadence = "monthly"
	PlanExternalPlanIDVersionNewParamsAddPricesPriceNewPlanGroupedTieredPriceCadenceQuarterly  PlanExternalPlanIDVersionNewParamsAddPricesPriceNewPlanGroupedTieredPriceCadence = "quarterly"
	PlanExternalPlanIDVersionNewParamsAddPricesPriceNewPlanGroupedTieredPriceCadenceOneTime    PlanExternalPlanIDVersionNewParamsAddPricesPriceNewPlanGroupedTieredPriceCadence = "one_time"
	PlanExternalPlanIDVersionNewParamsAddPricesPriceNewPlanGroupedTieredPriceCadenceCustom     PlanExternalPlanIDVersionNewParamsAddPricesPriceNewPlanGroupedTieredPriceCadence = "custom"
)

func (r PlanExternalPlanIDVersionNewParamsAddPricesPriceNewPlanGroupedTieredPriceCadence) IsKnown() bool {
	switch r {
	case PlanExternalPlanIDVersionNewParamsAddPricesPriceNewPlanGroupedTieredPriceCadenceAnnual, PlanExternalPlanIDVersionNewParamsAddPricesPriceNewPlanGroupedTieredPriceCadenceSemiAnnual, PlanExternalPlanIDVersionNewParamsAddPricesPriceNewPlanGroupedTieredPriceCadenceMonthly, PlanExternalPlanIDVersionNewParamsAddPricesPriceNewPlanGroupedTieredPriceCadenceQuarterly, PlanExternalPlanIDVersionNewParamsAddPricesPriceNewPlanGroupedTieredPriceCadenceOneTime, PlanExternalPlanIDVersionNewParamsAddPricesPriceNewPlanGroupedTieredPriceCadenceCustom:
		return true
	}
	return false
}

type PlanExternalPlanIDVersionNewParamsAddPricesPriceNewPlanGroupedTieredPriceModelType string

const (
	PlanExternalPlanIDVersionNewParamsAddPricesPriceNewPlanGroupedTieredPriceModelTypeGroupedTiered PlanExternalPlanIDVersionNewParamsAddPricesPriceNewPlanGroupedTieredPriceModelType = "grouped_tiered"
)

func (r PlanExternalPlanIDVersionNewParamsAddPricesPriceNewPlanGroupedTieredPriceModelType) IsKnown() bool {
	switch r {
	case PlanExternalPlanIDVersionNewParamsAddPricesPriceNewPlanGroupedTieredPriceModelTypeGroupedTiered:
		return true
	}
	return false
}

// For custom cadence: specifies the duration of the billing period in days or
// months.
type PlanExternalPlanIDVersionNewParamsAddPricesPriceNewPlanGroupedTieredPriceBillingCycleConfiguration struct {
	// The duration of the billing period.
	Duration param.Field[int64] `json:"duration,required"`
	// The unit of billing period duration.
	DurationUnit param.Field[PlanExternalPlanIDVersionNewParamsAddPricesPriceNewPlanGroupedTieredPriceBillingCycleConfigurationDurationUnit] `json:"duration_unit,required"`
}

func (r PlanExternalPlanIDVersionNewParamsAddPricesPriceNewPlanGroupedTieredPriceBillingCycleConfiguration) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// The unit of billing period duration.
type PlanExternalPlanIDVersionNewParamsAddPricesPriceNewPlanGroupedTieredPriceBillingCycleConfigurationDurationUnit string

const (
	PlanExternalPlanIDVersionNewParamsAddPricesPriceNewPlanGroupedTieredPriceBillingCycleConfigurationDurationUnitDay   PlanExternalPlanIDVersionNewParamsAddPricesPriceNewPlanGroupedTieredPriceBillingCycleConfigurationDurationUnit = "day"
	PlanExternalPlanIDVersionNewParamsAddPricesPriceNewPlanGroupedTieredPriceBillingCycleConfigurationDurationUnitMonth PlanExternalPlanIDVersionNewParamsAddPricesPriceNewPlanGroupedTieredPriceBillingCycleConfigurationDurationUnit = "month"
)

func (r PlanExternalPlanIDVersionNewParamsAddPricesPriceNewPlanGroupedTieredPriceBillingCycleConfigurationDurationUnit) IsKnown() bool {
	switch r {
	case PlanExternalPlanIDVersionNewParamsAddPricesPriceNewPlanGroupedTieredPriceBillingCycleConfigurationDurationUnitDay, PlanExternalPlanIDVersionNewParamsAddPricesPriceNewPlanGroupedTieredPriceBillingCycleConfigurationDurationUnitMonth:
		return true
	}
	return false
}

// For dimensional price: specifies a price group and dimension values
type PlanExternalPlanIDVersionNewParamsAddPricesPriceNewPlanGroupedTieredPriceDimensionalPriceConfiguration struct {
	// The list of dimension values matching (in order) the dimensions of the price
	// group
	DimensionValues param.Field[[]string] `json:"dimension_values,required"`
	// The id of the dimensional price group to include this price in
	DimensionalPriceGroupID param.Field[string] `json:"dimensional_price_group_id"`
	// The external id of the dimensional price group to include this price in
	ExternalDimensionalPriceGroupID param.Field[string] `json:"external_dimensional_price_group_id"`
}

func (r PlanExternalPlanIDVersionNewParamsAddPricesPriceNewPlanGroupedTieredPriceDimensionalPriceConfiguration) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Within each billing cycle, specifies the cadence at which invoices are produced.
// If unspecified, a single invoice is produced per billing cycle.
type PlanExternalPlanIDVersionNewParamsAddPricesPriceNewPlanGroupedTieredPriceInvoicingCycleConfiguration struct {
	// The duration of the billing period.
	Duration param.Field[int64] `json:"duration,required"`
	// The unit of billing period duration.
	DurationUnit param.Field[PlanExternalPlanIDVersionNewParamsAddPricesPriceNewPlanGroupedTieredPriceInvoicingCycleConfigurationDurationUnit] `json:"duration_unit,required"`
}

func (r PlanExternalPlanIDVersionNewParamsAddPricesPriceNewPlanGroupedTieredPriceInvoicingCycleConfiguration) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// The unit of billing period duration.
type PlanExternalPlanIDVersionNewParamsAddPricesPriceNewPlanGroupedTieredPriceInvoicingCycleConfigurationDurationUnit string

const (
	PlanExternalPlanIDVersionNewParamsAddPricesPriceNewPlanGroupedTieredPriceInvoicingCycleConfigurationDurationUnitDay   PlanExternalPlanIDVersionNewParamsAddPricesPriceNewPlanGroupedTieredPriceInvoicingCycleConfigurationDurationUnit = "day"
	PlanExternalPlanIDVersionNewParamsAddPricesPriceNewPlanGroupedTieredPriceInvoicingCycleConfigurationDurationUnitMonth PlanExternalPlanIDVersionNewParamsAddPricesPriceNewPlanGroupedTieredPriceInvoicingCycleConfigurationDurationUnit = "month"
)

func (r PlanExternalPlanIDVersionNewParamsAddPricesPriceNewPlanGroupedTieredPriceInvoicingCycleConfigurationDurationUnit) IsKnown() bool {
	switch r {
	case PlanExternalPlanIDVersionNewParamsAddPricesPriceNewPlanGroupedTieredPriceInvoicingCycleConfigurationDurationUnitDay, PlanExternalPlanIDVersionNewParamsAddPricesPriceNewPlanGroupedTieredPriceInvoicingCycleConfigurationDurationUnitMonth:
		return true
	}
	return false
}

// The cadence to bill for this price on.
type PlanExternalPlanIDVersionNewParamsAddPricesPriceCadence string

const (
	PlanExternalPlanIDVersionNewParamsAddPricesPriceCadenceAnnual     PlanExternalPlanIDVersionNewParamsAddPricesPriceCadence = "annual"
	PlanExternalPlanIDVersionNewParamsAddPricesPriceCadenceSemiAnnual PlanExternalPlanIDVersionNewParamsAddPricesPriceCadence = "semi_annual"
	PlanExternalPlanIDVersionNewParamsAddPricesPriceCadenceMonthly    PlanExternalPlanIDVersionNewParamsAddPricesPriceCadence = "monthly"
	PlanExternalPlanIDVersionNewParamsAddPricesPriceCadenceQuarterly  PlanExternalPlanIDVersionNewParamsAddPricesPriceCadence = "quarterly"
	PlanExternalPlanIDVersionNewParamsAddPricesPriceCadenceOneTime    PlanExternalPlanIDVersionNewParamsAddPricesPriceCadence = "one_time"
	PlanExternalPlanIDVersionNewParamsAddPricesPriceCadenceCustom     PlanExternalPlanIDVersionNewParamsAddPricesPriceCadence = "custom"
)

func (r PlanExternalPlanIDVersionNewParamsAddPricesPriceCadence) IsKnown() bool {
	switch r {
	case PlanExternalPlanIDVersionNewParamsAddPricesPriceCadenceAnnual, PlanExternalPlanIDVersionNewParamsAddPricesPriceCadenceSemiAnnual, PlanExternalPlanIDVersionNewParamsAddPricesPriceCadenceMonthly, PlanExternalPlanIDVersionNewParamsAddPricesPriceCadenceQuarterly, PlanExternalPlanIDVersionNewParamsAddPricesPriceCadenceOneTime, PlanExternalPlanIDVersionNewParamsAddPricesPriceCadenceCustom:
		return true
	}
	return false
}

type PlanExternalPlanIDVersionNewParamsAddPricesPriceModelType string

const (
	PlanExternalPlanIDVersionNewParamsAddPricesPriceModelTypeUnit                            PlanExternalPlanIDVersionNewParamsAddPricesPriceModelType = "unit"
	PlanExternalPlanIDVersionNewParamsAddPricesPriceModelTypePackage                         PlanExternalPlanIDVersionNewParamsAddPricesPriceModelType = "package"
	PlanExternalPlanIDVersionNewParamsAddPricesPriceModelTypeMatrix                          PlanExternalPlanIDVersionNewParamsAddPricesPriceModelType = "matrix"
	PlanExternalPlanIDVersionNewParamsAddPricesPriceModelTypeTiered                          PlanExternalPlanIDVersionNewParamsAddPricesPriceModelType = "tiered"
	PlanExternalPlanIDVersionNewParamsAddPricesPriceModelTypeTieredBps                       PlanExternalPlanIDVersionNewParamsAddPricesPriceModelType = "tiered_bps"
	PlanExternalPlanIDVersionNewParamsAddPricesPriceModelTypeBps                             PlanExternalPlanIDVersionNewParamsAddPricesPriceModelType = "bps"
	PlanExternalPlanIDVersionNewParamsAddPricesPriceModelTypeBulkBps                         PlanExternalPlanIDVersionNewParamsAddPricesPriceModelType = "bulk_bps"
	PlanExternalPlanIDVersionNewParamsAddPricesPriceModelTypeBulk                            PlanExternalPlanIDVersionNewParamsAddPricesPriceModelType = "bulk"
	PlanExternalPlanIDVersionNewParamsAddPricesPriceModelTypeThresholdTotalAmount            PlanExternalPlanIDVersionNewParamsAddPricesPriceModelType = "threshold_total_amount"
	PlanExternalPlanIDVersionNewParamsAddPricesPriceModelTypeTieredPackage                   PlanExternalPlanIDVersionNewParamsAddPricesPriceModelType = "tiered_package"
	PlanExternalPlanIDVersionNewParamsAddPricesPriceModelTypeTieredWithMinimum               PlanExternalPlanIDVersionNewParamsAddPricesPriceModelType = "tiered_with_minimum"
	PlanExternalPlanIDVersionNewParamsAddPricesPriceModelTypeUnitWithPercent                 PlanExternalPlanIDVersionNewParamsAddPricesPriceModelType = "unit_with_percent"
	PlanExternalPlanIDVersionNewParamsAddPricesPriceModelTypePackageWithAllocation           PlanExternalPlanIDVersionNewParamsAddPricesPriceModelType = "package_with_allocation"
	PlanExternalPlanIDVersionNewParamsAddPricesPriceModelTypeTieredWithProration             PlanExternalPlanIDVersionNewParamsAddPricesPriceModelType = "tiered_with_proration"
	PlanExternalPlanIDVersionNewParamsAddPricesPriceModelTypeUnitWithProration               PlanExternalPlanIDVersionNewParamsAddPricesPriceModelType = "unit_with_proration"
	PlanExternalPlanIDVersionNewParamsAddPricesPriceModelTypeGroupedAllocation               PlanExternalPlanIDVersionNewParamsAddPricesPriceModelType = "grouped_allocation"
	PlanExternalPlanIDVersionNewParamsAddPricesPriceModelTypeGroupedWithProratedMinimum      PlanExternalPlanIDVersionNewParamsAddPricesPriceModelType = "grouped_with_prorated_minimum"
	PlanExternalPlanIDVersionNewParamsAddPricesPriceModelTypeGroupedWithMeteredMinimum       PlanExternalPlanIDVersionNewParamsAddPricesPriceModelType = "grouped_with_metered_minimum"
	PlanExternalPlanIDVersionNewParamsAddPricesPriceModelTypeMatrixWithDisplayName           PlanExternalPlanIDVersionNewParamsAddPricesPriceModelType = "matrix_with_display_name"
	PlanExternalPlanIDVersionNewParamsAddPricesPriceModelTypeBulkWithProration               PlanExternalPlanIDVersionNewParamsAddPricesPriceModelType = "bulk_with_proration"
	PlanExternalPlanIDVersionNewParamsAddPricesPriceModelTypeGroupedTieredPackage            PlanExternalPlanIDVersionNewParamsAddPricesPriceModelType = "grouped_tiered_package"
	PlanExternalPlanIDVersionNewParamsAddPricesPriceModelTypeMaxGroupTieredPackage           PlanExternalPlanIDVersionNewParamsAddPricesPriceModelType = "max_group_tiered_package"
	PlanExternalPlanIDVersionNewParamsAddPricesPriceModelTypeScalableMatrixWithUnitPricing   PlanExternalPlanIDVersionNewParamsAddPricesPriceModelType = "scalable_matrix_with_unit_pricing"
	PlanExternalPlanIDVersionNewParamsAddPricesPriceModelTypeScalableMatrixWithTieredPricing PlanExternalPlanIDVersionNewParamsAddPricesPriceModelType = "scalable_matrix_with_tiered_pricing"
	PlanExternalPlanIDVersionNewParamsAddPricesPriceModelTypeCumulativeGroupedBulk           PlanExternalPlanIDVersionNewParamsAddPricesPriceModelType = "cumulative_grouped_bulk"
	PlanExternalPlanIDVersionNewParamsAddPricesPriceModelTypeTieredPackageWithMinimum        PlanExternalPlanIDVersionNewParamsAddPricesPriceModelType = "tiered_package_with_minimum"
	PlanExternalPlanIDVersionNewParamsAddPricesPriceModelTypeMatrixWithAllocation            PlanExternalPlanIDVersionNewParamsAddPricesPriceModelType = "matrix_with_allocation"
	PlanExternalPlanIDVersionNewParamsAddPricesPriceModelTypeGroupedTiered                   PlanExternalPlanIDVersionNewParamsAddPricesPriceModelType = "grouped_tiered"
)

func (r PlanExternalPlanIDVersionNewParamsAddPricesPriceModelType) IsKnown() bool {
	switch r {
	case PlanExternalPlanIDVersionNewParamsAddPricesPriceModelTypeUnit, PlanExternalPlanIDVersionNewParamsAddPricesPriceModelTypePackage, PlanExternalPlanIDVersionNewParamsAddPricesPriceModelTypeMatrix, PlanExternalPlanIDVersionNewParamsAddPricesPriceModelTypeTiered, PlanExternalPlanIDVersionNewParamsAddPricesPriceModelTypeTieredBps, PlanExternalPlanIDVersionNewParamsAddPricesPriceModelTypeBps, PlanExternalPlanIDVersionNewParamsAddPricesPriceModelTypeBulkBps, PlanExternalPlanIDVersionNewParamsAddPricesPriceModelTypeBulk, PlanExternalPlanIDVersionNewParamsAddPricesPriceModelTypeThresholdTotalAmount, PlanExternalPlanIDVersionNewParamsAddPricesPriceModelTypeTieredPackage, PlanExternalPlanIDVersionNewParamsAddPricesPriceModelTypeTieredWithMinimum, PlanExternalPlanIDVersionNewParamsAddPricesPriceModelTypeUnitWithPercent, PlanExternalPlanIDVersionNewParamsAddPricesPriceModelTypePackageWithAllocation, PlanExternalPlanIDVersionNewParamsAddPricesPriceModelTypeTieredWithProration, PlanExternalPlanIDVersionNewParamsAddPricesPriceModelTypeUnitWithProration, PlanExternalPlanIDVersionNewParamsAddPricesPriceModelTypeGroupedAllocation, PlanExternalPlanIDVersionNewParamsAddPricesPriceModelTypeGroupedWithProratedMinimum, PlanExternalPlanIDVersionNewParamsAddPricesPriceModelTypeGroupedWithMeteredMinimum, PlanExternalPlanIDVersionNewParamsAddPricesPriceModelTypeMatrixWithDisplayName, PlanExternalPlanIDVersionNewParamsAddPricesPriceModelTypeBulkWithProration, PlanExternalPlanIDVersionNewParamsAddPricesPriceModelTypeGroupedTieredPackage, PlanExternalPlanIDVersionNewParamsAddPricesPriceModelTypeMaxGroupTieredPackage, PlanExternalPlanIDVersionNewParamsAddPricesPriceModelTypeScalableMatrixWithUnitPricing, PlanExternalPlanIDVersionNewParamsAddPricesPriceModelTypeScalableMatrixWithTieredPricing, PlanExternalPlanIDVersionNewParamsAddPricesPriceModelTypeCumulativeGroupedBulk, PlanExternalPlanIDVersionNewParamsAddPricesPriceModelTypeTieredPackageWithMinimum, PlanExternalPlanIDVersionNewParamsAddPricesPriceModelTypeMatrixWithAllocation, PlanExternalPlanIDVersionNewParamsAddPricesPriceModelTypeGroupedTiered:
		return true
	}
	return false
}

type PlanExternalPlanIDVersionNewParamsRemoveAdjustment struct {
	// The id of the adjustment to remove from on the plan.
	AdjustmentID param.Field[string] `json:"adjustment_id,required"`
	// The phase to remove this adjustment from.
	PlanPhaseOrder param.Field[int64] `json:"plan_phase_order"`
}

func (r PlanExternalPlanIDVersionNewParamsRemoveAdjustment) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type PlanExternalPlanIDVersionNewParamsRemovePrice struct {
	// The id of the price to remove from the plan.
	PriceID param.Field[string] `json:"price_id,required"`
	// The phase to remove this price from.
	PlanPhaseOrder param.Field[int64] `json:"plan_phase_order"`
}

func (r PlanExternalPlanIDVersionNewParamsRemovePrice) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type PlanExternalPlanIDVersionNewParamsReplaceAdjustment struct {
	// The definition of a new adjustment to create and add to the plan.
	Adjustment param.Field[PlanExternalPlanIDVersionNewParamsReplaceAdjustmentsAdjustmentUnion] `json:"adjustment,required"`
	// The id of the adjustment on the plan to replace in the plan.
	ReplacesAdjustmentID param.Field[string] `json:"replaces_adjustment_id,required"`
	// The phase to replace this adjustment from.
	PlanPhaseOrder param.Field[int64] `json:"plan_phase_order"`
}

func (r PlanExternalPlanIDVersionNewParamsReplaceAdjustment) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// The definition of a new adjustment to create and add to the plan.
type PlanExternalPlanIDVersionNewParamsReplaceAdjustmentsAdjustment struct {
	AdjustmentType    param.Field[PlanExternalPlanIDVersionNewParamsReplaceAdjustmentsAdjustmentAdjustmentType] `json:"adjustment_type,required"`
	AmountDiscount    param.Field[string]                                                                       `json:"amount_discount"`
	AppliesToPriceIDs param.Field[interface{}]                                                                  `json:"applies_to_price_ids"`
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

func (r PlanExternalPlanIDVersionNewParamsReplaceAdjustmentsAdjustment) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r PlanExternalPlanIDVersionNewParamsReplaceAdjustmentsAdjustment) implementsPlanExternalPlanIDVersionNewParamsReplaceAdjustmentsAdjustmentUnion() {
}

// The definition of a new adjustment to create and add to the plan.
//
// Satisfied by
// [PlanExternalPlanIDVersionNewParamsReplaceAdjustmentsAdjustmentNewPercentageDiscount],
// [PlanExternalPlanIDVersionNewParamsReplaceAdjustmentsAdjustmentNewUsageDiscount],
// [PlanExternalPlanIDVersionNewParamsReplaceAdjustmentsAdjustmentNewAmountDiscount],
// [PlanExternalPlanIDVersionNewParamsReplaceAdjustmentsAdjustmentNewMinimum],
// [PlanExternalPlanIDVersionNewParamsReplaceAdjustmentsAdjustmentNewMaximum],
// [PlanExternalPlanIDVersionNewParamsReplaceAdjustmentsAdjustment].
type PlanExternalPlanIDVersionNewParamsReplaceAdjustmentsAdjustmentUnion interface {
	implementsPlanExternalPlanIDVersionNewParamsReplaceAdjustmentsAdjustmentUnion()
}

type PlanExternalPlanIDVersionNewParamsReplaceAdjustmentsAdjustmentNewPercentageDiscount struct {
	AdjustmentType     param.Field[PlanExternalPlanIDVersionNewParamsReplaceAdjustmentsAdjustmentNewPercentageDiscountAdjustmentType] `json:"adjustment_type,required"`
	PercentageDiscount param.Field[float64]                                                                                           `json:"percentage_discount,required"`
	// The set of price IDs to which this adjustment applies.
	AppliesToPriceIDs param.Field[[]string] `json:"applies_to_price_ids"`
	// When false, this adjustment will be applied to a single price. Otherwise, it
	// will be applied at the invoice level, possibly to multiple prices.
	IsInvoiceLevel param.Field[bool] `json:"is_invoice_level"`
}

func (r PlanExternalPlanIDVersionNewParamsReplaceAdjustmentsAdjustmentNewPercentageDiscount) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r PlanExternalPlanIDVersionNewParamsReplaceAdjustmentsAdjustmentNewPercentageDiscount) implementsPlanExternalPlanIDVersionNewParamsReplaceAdjustmentsAdjustmentUnion() {
}

type PlanExternalPlanIDVersionNewParamsReplaceAdjustmentsAdjustmentNewPercentageDiscountAdjustmentType string

const (
	PlanExternalPlanIDVersionNewParamsReplaceAdjustmentsAdjustmentNewPercentageDiscountAdjustmentTypePercentageDiscount PlanExternalPlanIDVersionNewParamsReplaceAdjustmentsAdjustmentNewPercentageDiscountAdjustmentType = "percentage_discount"
)

func (r PlanExternalPlanIDVersionNewParamsReplaceAdjustmentsAdjustmentNewPercentageDiscountAdjustmentType) IsKnown() bool {
	switch r {
	case PlanExternalPlanIDVersionNewParamsReplaceAdjustmentsAdjustmentNewPercentageDiscountAdjustmentTypePercentageDiscount:
		return true
	}
	return false
}

type PlanExternalPlanIDVersionNewParamsReplaceAdjustmentsAdjustmentNewUsageDiscount struct {
	AdjustmentType param.Field[PlanExternalPlanIDVersionNewParamsReplaceAdjustmentsAdjustmentNewUsageDiscountAdjustmentType] `json:"adjustment_type,required"`
	UsageDiscount  param.Field[float64]                                                                                      `json:"usage_discount,required"`
	// The set of price IDs to which this adjustment applies.
	AppliesToPriceIDs param.Field[[]string] `json:"applies_to_price_ids"`
	// When false, this adjustment will be applied to a single price. Otherwise, it
	// will be applied at the invoice level, possibly to multiple prices.
	IsInvoiceLevel param.Field[bool] `json:"is_invoice_level"`
}

func (r PlanExternalPlanIDVersionNewParamsReplaceAdjustmentsAdjustmentNewUsageDiscount) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r PlanExternalPlanIDVersionNewParamsReplaceAdjustmentsAdjustmentNewUsageDiscount) implementsPlanExternalPlanIDVersionNewParamsReplaceAdjustmentsAdjustmentUnion() {
}

type PlanExternalPlanIDVersionNewParamsReplaceAdjustmentsAdjustmentNewUsageDiscountAdjustmentType string

const (
	PlanExternalPlanIDVersionNewParamsReplaceAdjustmentsAdjustmentNewUsageDiscountAdjustmentTypeUsageDiscount PlanExternalPlanIDVersionNewParamsReplaceAdjustmentsAdjustmentNewUsageDiscountAdjustmentType = "usage_discount"
)

func (r PlanExternalPlanIDVersionNewParamsReplaceAdjustmentsAdjustmentNewUsageDiscountAdjustmentType) IsKnown() bool {
	switch r {
	case PlanExternalPlanIDVersionNewParamsReplaceAdjustmentsAdjustmentNewUsageDiscountAdjustmentTypeUsageDiscount:
		return true
	}
	return false
}

type PlanExternalPlanIDVersionNewParamsReplaceAdjustmentsAdjustmentNewAmountDiscount struct {
	AdjustmentType param.Field[PlanExternalPlanIDVersionNewParamsReplaceAdjustmentsAdjustmentNewAmountDiscountAdjustmentType] `json:"adjustment_type,required"`
	AmountDiscount param.Field[string]                                                                                        `json:"amount_discount,required"`
	// The set of price IDs to which this adjustment applies.
	AppliesToPriceIDs param.Field[[]string] `json:"applies_to_price_ids"`
	// When false, this adjustment will be applied to a single price. Otherwise, it
	// will be applied at the invoice level, possibly to multiple prices.
	IsInvoiceLevel param.Field[bool] `json:"is_invoice_level"`
}

func (r PlanExternalPlanIDVersionNewParamsReplaceAdjustmentsAdjustmentNewAmountDiscount) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r PlanExternalPlanIDVersionNewParamsReplaceAdjustmentsAdjustmentNewAmountDiscount) implementsPlanExternalPlanIDVersionNewParamsReplaceAdjustmentsAdjustmentUnion() {
}

type PlanExternalPlanIDVersionNewParamsReplaceAdjustmentsAdjustmentNewAmountDiscountAdjustmentType string

const (
	PlanExternalPlanIDVersionNewParamsReplaceAdjustmentsAdjustmentNewAmountDiscountAdjustmentTypeAmountDiscount PlanExternalPlanIDVersionNewParamsReplaceAdjustmentsAdjustmentNewAmountDiscountAdjustmentType = "amount_discount"
)

func (r PlanExternalPlanIDVersionNewParamsReplaceAdjustmentsAdjustmentNewAmountDiscountAdjustmentType) IsKnown() bool {
	switch r {
	case PlanExternalPlanIDVersionNewParamsReplaceAdjustmentsAdjustmentNewAmountDiscountAdjustmentTypeAmountDiscount:
		return true
	}
	return false
}

type PlanExternalPlanIDVersionNewParamsReplaceAdjustmentsAdjustmentNewMinimum struct {
	AdjustmentType param.Field[PlanExternalPlanIDVersionNewParamsReplaceAdjustmentsAdjustmentNewMinimumAdjustmentType] `json:"adjustment_type,required"`
	// The item ID that revenue from this minimum will be attributed to.
	ItemID        param.Field[string] `json:"item_id,required"`
	MinimumAmount param.Field[string] `json:"minimum_amount,required"`
	// The set of price IDs to which this adjustment applies.
	AppliesToPriceIDs param.Field[[]string] `json:"applies_to_price_ids"`
	// When false, this adjustment will be applied to a single price. Otherwise, it
	// will be applied at the invoice level, possibly to multiple prices.
	IsInvoiceLevel param.Field[bool] `json:"is_invoice_level"`
}

func (r PlanExternalPlanIDVersionNewParamsReplaceAdjustmentsAdjustmentNewMinimum) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r PlanExternalPlanIDVersionNewParamsReplaceAdjustmentsAdjustmentNewMinimum) implementsPlanExternalPlanIDVersionNewParamsReplaceAdjustmentsAdjustmentUnion() {
}

type PlanExternalPlanIDVersionNewParamsReplaceAdjustmentsAdjustmentNewMinimumAdjustmentType string

const (
	PlanExternalPlanIDVersionNewParamsReplaceAdjustmentsAdjustmentNewMinimumAdjustmentTypeMinimum PlanExternalPlanIDVersionNewParamsReplaceAdjustmentsAdjustmentNewMinimumAdjustmentType = "minimum"
)

func (r PlanExternalPlanIDVersionNewParamsReplaceAdjustmentsAdjustmentNewMinimumAdjustmentType) IsKnown() bool {
	switch r {
	case PlanExternalPlanIDVersionNewParamsReplaceAdjustmentsAdjustmentNewMinimumAdjustmentTypeMinimum:
		return true
	}
	return false
}

type PlanExternalPlanIDVersionNewParamsReplaceAdjustmentsAdjustmentNewMaximum struct {
	AdjustmentType param.Field[PlanExternalPlanIDVersionNewParamsReplaceAdjustmentsAdjustmentNewMaximumAdjustmentType] `json:"adjustment_type,required"`
	MaximumAmount  param.Field[string]                                                                                 `json:"maximum_amount,required"`
	// The set of price IDs to which this adjustment applies.
	AppliesToPriceIDs param.Field[[]string] `json:"applies_to_price_ids"`
	// When false, this adjustment will be applied to a single price. Otherwise, it
	// will be applied at the invoice level, possibly to multiple prices.
	IsInvoiceLevel param.Field[bool] `json:"is_invoice_level"`
}

func (r PlanExternalPlanIDVersionNewParamsReplaceAdjustmentsAdjustmentNewMaximum) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r PlanExternalPlanIDVersionNewParamsReplaceAdjustmentsAdjustmentNewMaximum) implementsPlanExternalPlanIDVersionNewParamsReplaceAdjustmentsAdjustmentUnion() {
}

type PlanExternalPlanIDVersionNewParamsReplaceAdjustmentsAdjustmentNewMaximumAdjustmentType string

const (
	PlanExternalPlanIDVersionNewParamsReplaceAdjustmentsAdjustmentNewMaximumAdjustmentTypeMaximum PlanExternalPlanIDVersionNewParamsReplaceAdjustmentsAdjustmentNewMaximumAdjustmentType = "maximum"
)

func (r PlanExternalPlanIDVersionNewParamsReplaceAdjustmentsAdjustmentNewMaximumAdjustmentType) IsKnown() bool {
	switch r {
	case PlanExternalPlanIDVersionNewParamsReplaceAdjustmentsAdjustmentNewMaximumAdjustmentTypeMaximum:
		return true
	}
	return false
}

type PlanExternalPlanIDVersionNewParamsReplaceAdjustmentsAdjustmentAdjustmentType string

const (
	PlanExternalPlanIDVersionNewParamsReplaceAdjustmentsAdjustmentAdjustmentTypePercentageDiscount PlanExternalPlanIDVersionNewParamsReplaceAdjustmentsAdjustmentAdjustmentType = "percentage_discount"
	PlanExternalPlanIDVersionNewParamsReplaceAdjustmentsAdjustmentAdjustmentTypeUsageDiscount      PlanExternalPlanIDVersionNewParamsReplaceAdjustmentsAdjustmentAdjustmentType = "usage_discount"
	PlanExternalPlanIDVersionNewParamsReplaceAdjustmentsAdjustmentAdjustmentTypeAmountDiscount     PlanExternalPlanIDVersionNewParamsReplaceAdjustmentsAdjustmentAdjustmentType = "amount_discount"
	PlanExternalPlanIDVersionNewParamsReplaceAdjustmentsAdjustmentAdjustmentTypeMinimum            PlanExternalPlanIDVersionNewParamsReplaceAdjustmentsAdjustmentAdjustmentType = "minimum"
	PlanExternalPlanIDVersionNewParamsReplaceAdjustmentsAdjustmentAdjustmentTypeMaximum            PlanExternalPlanIDVersionNewParamsReplaceAdjustmentsAdjustmentAdjustmentType = "maximum"
)

func (r PlanExternalPlanIDVersionNewParamsReplaceAdjustmentsAdjustmentAdjustmentType) IsKnown() bool {
	switch r {
	case PlanExternalPlanIDVersionNewParamsReplaceAdjustmentsAdjustmentAdjustmentTypePercentageDiscount, PlanExternalPlanIDVersionNewParamsReplaceAdjustmentsAdjustmentAdjustmentTypeUsageDiscount, PlanExternalPlanIDVersionNewParamsReplaceAdjustmentsAdjustmentAdjustmentTypeAmountDiscount, PlanExternalPlanIDVersionNewParamsReplaceAdjustmentsAdjustmentAdjustmentTypeMinimum, PlanExternalPlanIDVersionNewParamsReplaceAdjustmentsAdjustmentAdjustmentTypeMaximum:
		return true
	}
	return false
}

type PlanExternalPlanIDVersionNewParamsReplacePrice struct {
	// The id of the price on the plan to replace in the plan.
	ReplacesPriceID param.Field[string] `json:"replaces_price_id,required"`
	// The allocation price to add to the plan.
	AllocationPrice param.Field[PlanExternalPlanIDVersionNewParamsReplacePricesAllocationPrice] `json:"allocation_price"`
	// The phase to replace this price from.
	PlanPhaseOrder param.Field[int64] `json:"plan_phase_order"`
	// The price to add to the plan
	Price param.Field[PlanExternalPlanIDVersionNewParamsReplacePricesPriceUnion] `json:"price"`
}

func (r PlanExternalPlanIDVersionNewParamsReplacePrice) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// The allocation price to add to the plan.
type PlanExternalPlanIDVersionNewParamsReplacePricesAllocationPrice struct {
	// An amount of the currency to allocate to the customer at the specified cadence.
	Amount param.Field[string] `json:"amount,required"`
	// The cadence at which to allocate the amount to the customer.
	Cadence param.Field[PlanExternalPlanIDVersionNewParamsReplacePricesAllocationPriceCadence] `json:"cadence,required"`
	// An ISO 4217 currency string or a custom pricing unit identifier in which to bill
	// this price.
	Currency param.Field[string] `json:"currency,required"`
	// Whether the allocated amount should expire at the end of the cadence or roll
	// over to the next period.
	ExpiresAtEndOfCadence param.Field[bool] `json:"expires_at_end_of_cadence,required"`
}

func (r PlanExternalPlanIDVersionNewParamsReplacePricesAllocationPrice) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// The cadence at which to allocate the amount to the customer.
type PlanExternalPlanIDVersionNewParamsReplacePricesAllocationPriceCadence string

const (
	PlanExternalPlanIDVersionNewParamsReplacePricesAllocationPriceCadenceOneTime    PlanExternalPlanIDVersionNewParamsReplacePricesAllocationPriceCadence = "one_time"
	PlanExternalPlanIDVersionNewParamsReplacePricesAllocationPriceCadenceMonthly    PlanExternalPlanIDVersionNewParamsReplacePricesAllocationPriceCadence = "monthly"
	PlanExternalPlanIDVersionNewParamsReplacePricesAllocationPriceCadenceQuarterly  PlanExternalPlanIDVersionNewParamsReplacePricesAllocationPriceCadence = "quarterly"
	PlanExternalPlanIDVersionNewParamsReplacePricesAllocationPriceCadenceSemiAnnual PlanExternalPlanIDVersionNewParamsReplacePricesAllocationPriceCadence = "semi_annual"
	PlanExternalPlanIDVersionNewParamsReplacePricesAllocationPriceCadenceAnnual     PlanExternalPlanIDVersionNewParamsReplacePricesAllocationPriceCadence = "annual"
	PlanExternalPlanIDVersionNewParamsReplacePricesAllocationPriceCadenceCustom     PlanExternalPlanIDVersionNewParamsReplacePricesAllocationPriceCadence = "custom"
)

func (r PlanExternalPlanIDVersionNewParamsReplacePricesAllocationPriceCadence) IsKnown() bool {
	switch r {
	case PlanExternalPlanIDVersionNewParamsReplacePricesAllocationPriceCadenceOneTime, PlanExternalPlanIDVersionNewParamsReplacePricesAllocationPriceCadenceMonthly, PlanExternalPlanIDVersionNewParamsReplacePricesAllocationPriceCadenceQuarterly, PlanExternalPlanIDVersionNewParamsReplacePricesAllocationPriceCadenceSemiAnnual, PlanExternalPlanIDVersionNewParamsReplacePricesAllocationPriceCadenceAnnual, PlanExternalPlanIDVersionNewParamsReplacePricesAllocationPriceCadenceCustom:
		return true
	}
	return false
}

// The price to add to the plan
type PlanExternalPlanIDVersionNewParamsReplacePricesPrice struct {
	// The cadence to bill for this price on.
	Cadence param.Field[PlanExternalPlanIDVersionNewParamsReplacePricesPriceCadence] `json:"cadence,required"`
	// The id of the item the price will be associated with.
	ItemID    param.Field[string]                                                        `json:"item_id,required"`
	ModelType param.Field[PlanExternalPlanIDVersionNewParamsReplacePricesPriceModelType] `json:"model_type,required"`
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

func (r PlanExternalPlanIDVersionNewParamsReplacePricesPrice) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r PlanExternalPlanIDVersionNewParamsReplacePricesPrice) implementsPlanExternalPlanIDVersionNewParamsReplacePricesPriceUnion() {
}

// The price to add to the plan
//
// Satisfied by
// [PlanExternalPlanIDVersionNewParamsReplacePricesPriceNewPlanUnitPrice],
// [PlanExternalPlanIDVersionNewParamsReplacePricesPriceNewPlanPackagePrice],
// [PlanExternalPlanIDVersionNewParamsReplacePricesPriceNewPlanMatrixPrice],
// [PlanExternalPlanIDVersionNewParamsReplacePricesPriceNewPlanTieredPrice],
// [PlanExternalPlanIDVersionNewParamsReplacePricesPriceNewPlanTieredBpsPrice],
// [PlanExternalPlanIDVersionNewParamsReplacePricesPriceNewPlanBpsPrice],
// [PlanExternalPlanIDVersionNewParamsReplacePricesPriceNewPlanBulkBpsPrice],
// [PlanExternalPlanIDVersionNewParamsReplacePricesPriceNewPlanBulkPrice],
// [PlanExternalPlanIDVersionNewParamsReplacePricesPriceNewPlanThresholdTotalAmountPrice],
// [PlanExternalPlanIDVersionNewParamsReplacePricesPriceNewPlanTieredPackagePrice],
// [PlanExternalPlanIDVersionNewParamsReplacePricesPriceNewPlanTieredWithMinimumPrice],
// [PlanExternalPlanIDVersionNewParamsReplacePricesPriceNewPlanUnitWithPercentPrice],
// [PlanExternalPlanIDVersionNewParamsReplacePricesPriceNewPlanPackageWithAllocationPrice],
// [PlanExternalPlanIDVersionNewParamsReplacePricesPriceNewPlanTierWithProrationPrice],
// [PlanExternalPlanIDVersionNewParamsReplacePricesPriceNewPlanUnitWithProrationPrice],
// [PlanExternalPlanIDVersionNewParamsReplacePricesPriceNewPlanGroupedAllocationPrice],
// [PlanExternalPlanIDVersionNewParamsReplacePricesPriceNewPlanGroupedWithProratedMinimumPrice],
// [PlanExternalPlanIDVersionNewParamsReplacePricesPriceNewPlanGroupedWithMeteredMinimumPrice],
// [PlanExternalPlanIDVersionNewParamsReplacePricesPriceNewPlanMatrixWithDisplayNamePrice],
// [PlanExternalPlanIDVersionNewParamsReplacePricesPriceNewPlanBulkWithProrationPrice],
// [PlanExternalPlanIDVersionNewParamsReplacePricesPriceNewPlanGroupedTieredPackagePrice],
// [PlanExternalPlanIDVersionNewParamsReplacePricesPriceNewPlanMaxGroupTieredPackagePrice],
// [PlanExternalPlanIDVersionNewParamsReplacePricesPriceNewPlanScalableMatrixWithUnitPricingPrice],
// [PlanExternalPlanIDVersionNewParamsReplacePricesPriceNewPlanScalableMatrixWithTieredPricingPrice],
// [PlanExternalPlanIDVersionNewParamsReplacePricesPriceNewPlanCumulativeGroupedBulkPrice],
// [PlanExternalPlanIDVersionNewParamsReplacePricesPriceNewPlanTieredPackageWithMinimumPrice],
// [PlanExternalPlanIDVersionNewParamsReplacePricesPriceNewPlanMatrixWithAllocationPrice],
// [PlanExternalPlanIDVersionNewParamsReplacePricesPriceNewPlanGroupedTieredPrice],
// [PlanExternalPlanIDVersionNewParamsReplacePricesPrice].
type PlanExternalPlanIDVersionNewParamsReplacePricesPriceUnion interface {
	implementsPlanExternalPlanIDVersionNewParamsReplacePricesPriceUnion()
}

type PlanExternalPlanIDVersionNewParamsReplacePricesPriceNewPlanUnitPrice struct {
	// The cadence to bill for this price on.
	Cadence param.Field[PlanExternalPlanIDVersionNewParamsReplacePricesPriceNewPlanUnitPriceCadence] `json:"cadence,required"`
	// The id of the item the price will be associated with.
	ItemID    param.Field[string]                                                                        `json:"item_id,required"`
	ModelType param.Field[PlanExternalPlanIDVersionNewParamsReplacePricesPriceNewPlanUnitPriceModelType] `json:"model_type,required"`
	// The name of the price.
	Name       param.Field[string]                                                                         `json:"name,required"`
	UnitConfig param.Field[PlanExternalPlanIDVersionNewParamsReplacePricesPriceNewPlanUnitPriceUnitConfig] `json:"unit_config,required"`
	// The id of the billable metric for the price. Only needed if the price is
	// usage-based.
	BillableMetricID param.Field[string] `json:"billable_metric_id"`
	// If the Price represents a fixed cost, the price will be billed in-advance if
	// this is true, and in-arrears if this is false.
	BilledInAdvance param.Field[bool] `json:"billed_in_advance"`
	// For custom cadence: specifies the duration of the billing period in days or
	// months.
	BillingCycleConfiguration param.Field[PlanExternalPlanIDVersionNewParamsReplacePricesPriceNewPlanUnitPriceBillingCycleConfiguration] `json:"billing_cycle_configuration"`
	// The per unit conversion rate of the price currency to the invoicing currency.
	ConversionRate param.Field[float64] `json:"conversion_rate"`
	// An ISO 4217 currency string, or custom pricing unit identifier, in which this
	// price is billed.
	Currency param.Field[string] `json:"currency"`
	// For dimensional price: specifies a price group and dimension values
	DimensionalPriceConfiguration param.Field[PlanExternalPlanIDVersionNewParamsReplacePricesPriceNewPlanUnitPriceDimensionalPriceConfiguration] `json:"dimensional_price_configuration"`
	// An alias for the price.
	ExternalPriceID param.Field[string] `json:"external_price_id"`
	// If the Price represents a fixed cost, this represents the quantity of units
	// applied.
	FixedPriceQuantity param.Field[float64] `json:"fixed_price_quantity"`
	// The property used to group this price on an invoice
	InvoiceGroupingKey param.Field[string] `json:"invoice_grouping_key"`
	// Within each billing cycle, specifies the cadence at which invoices are produced.
	// If unspecified, a single invoice is produced per billing cycle.
	InvoicingCycleConfiguration param.Field[PlanExternalPlanIDVersionNewParamsReplacePricesPriceNewPlanUnitPriceInvoicingCycleConfiguration] `json:"invoicing_cycle_configuration"`
	// User-specified key/value pairs for the resource. Individual keys can be removed
	// by setting the value to `null`, and the entire metadata mapping can be cleared
	// by setting `metadata` to `null`.
	Metadata param.Field[map[string]string] `json:"metadata"`
}

func (r PlanExternalPlanIDVersionNewParamsReplacePricesPriceNewPlanUnitPrice) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r PlanExternalPlanIDVersionNewParamsReplacePricesPriceNewPlanUnitPrice) implementsPlanExternalPlanIDVersionNewParamsReplacePricesPriceUnion() {
}

// The cadence to bill for this price on.
type PlanExternalPlanIDVersionNewParamsReplacePricesPriceNewPlanUnitPriceCadence string

const (
	PlanExternalPlanIDVersionNewParamsReplacePricesPriceNewPlanUnitPriceCadenceAnnual     PlanExternalPlanIDVersionNewParamsReplacePricesPriceNewPlanUnitPriceCadence = "annual"
	PlanExternalPlanIDVersionNewParamsReplacePricesPriceNewPlanUnitPriceCadenceSemiAnnual PlanExternalPlanIDVersionNewParamsReplacePricesPriceNewPlanUnitPriceCadence = "semi_annual"
	PlanExternalPlanIDVersionNewParamsReplacePricesPriceNewPlanUnitPriceCadenceMonthly    PlanExternalPlanIDVersionNewParamsReplacePricesPriceNewPlanUnitPriceCadence = "monthly"
	PlanExternalPlanIDVersionNewParamsReplacePricesPriceNewPlanUnitPriceCadenceQuarterly  PlanExternalPlanIDVersionNewParamsReplacePricesPriceNewPlanUnitPriceCadence = "quarterly"
	PlanExternalPlanIDVersionNewParamsReplacePricesPriceNewPlanUnitPriceCadenceOneTime    PlanExternalPlanIDVersionNewParamsReplacePricesPriceNewPlanUnitPriceCadence = "one_time"
	PlanExternalPlanIDVersionNewParamsReplacePricesPriceNewPlanUnitPriceCadenceCustom     PlanExternalPlanIDVersionNewParamsReplacePricesPriceNewPlanUnitPriceCadence = "custom"
)

func (r PlanExternalPlanIDVersionNewParamsReplacePricesPriceNewPlanUnitPriceCadence) IsKnown() bool {
	switch r {
	case PlanExternalPlanIDVersionNewParamsReplacePricesPriceNewPlanUnitPriceCadenceAnnual, PlanExternalPlanIDVersionNewParamsReplacePricesPriceNewPlanUnitPriceCadenceSemiAnnual, PlanExternalPlanIDVersionNewParamsReplacePricesPriceNewPlanUnitPriceCadenceMonthly, PlanExternalPlanIDVersionNewParamsReplacePricesPriceNewPlanUnitPriceCadenceQuarterly, PlanExternalPlanIDVersionNewParamsReplacePricesPriceNewPlanUnitPriceCadenceOneTime, PlanExternalPlanIDVersionNewParamsReplacePricesPriceNewPlanUnitPriceCadenceCustom:
		return true
	}
	return false
}

type PlanExternalPlanIDVersionNewParamsReplacePricesPriceNewPlanUnitPriceModelType string

const (
	PlanExternalPlanIDVersionNewParamsReplacePricesPriceNewPlanUnitPriceModelTypeUnit PlanExternalPlanIDVersionNewParamsReplacePricesPriceNewPlanUnitPriceModelType = "unit"
)

func (r PlanExternalPlanIDVersionNewParamsReplacePricesPriceNewPlanUnitPriceModelType) IsKnown() bool {
	switch r {
	case PlanExternalPlanIDVersionNewParamsReplacePricesPriceNewPlanUnitPriceModelTypeUnit:
		return true
	}
	return false
}

type PlanExternalPlanIDVersionNewParamsReplacePricesPriceNewPlanUnitPriceUnitConfig struct {
	// Rate per unit of usage
	UnitAmount param.Field[string] `json:"unit_amount,required"`
}

func (r PlanExternalPlanIDVersionNewParamsReplacePricesPriceNewPlanUnitPriceUnitConfig) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// For custom cadence: specifies the duration of the billing period in days or
// months.
type PlanExternalPlanIDVersionNewParamsReplacePricesPriceNewPlanUnitPriceBillingCycleConfiguration struct {
	// The duration of the billing period.
	Duration param.Field[int64] `json:"duration,required"`
	// The unit of billing period duration.
	DurationUnit param.Field[PlanExternalPlanIDVersionNewParamsReplacePricesPriceNewPlanUnitPriceBillingCycleConfigurationDurationUnit] `json:"duration_unit,required"`
}

func (r PlanExternalPlanIDVersionNewParamsReplacePricesPriceNewPlanUnitPriceBillingCycleConfiguration) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// The unit of billing period duration.
type PlanExternalPlanIDVersionNewParamsReplacePricesPriceNewPlanUnitPriceBillingCycleConfigurationDurationUnit string

const (
	PlanExternalPlanIDVersionNewParamsReplacePricesPriceNewPlanUnitPriceBillingCycleConfigurationDurationUnitDay   PlanExternalPlanIDVersionNewParamsReplacePricesPriceNewPlanUnitPriceBillingCycleConfigurationDurationUnit = "day"
	PlanExternalPlanIDVersionNewParamsReplacePricesPriceNewPlanUnitPriceBillingCycleConfigurationDurationUnitMonth PlanExternalPlanIDVersionNewParamsReplacePricesPriceNewPlanUnitPriceBillingCycleConfigurationDurationUnit = "month"
)

func (r PlanExternalPlanIDVersionNewParamsReplacePricesPriceNewPlanUnitPriceBillingCycleConfigurationDurationUnit) IsKnown() bool {
	switch r {
	case PlanExternalPlanIDVersionNewParamsReplacePricesPriceNewPlanUnitPriceBillingCycleConfigurationDurationUnitDay, PlanExternalPlanIDVersionNewParamsReplacePricesPriceNewPlanUnitPriceBillingCycleConfigurationDurationUnitMonth:
		return true
	}
	return false
}

// For dimensional price: specifies a price group and dimension values
type PlanExternalPlanIDVersionNewParamsReplacePricesPriceNewPlanUnitPriceDimensionalPriceConfiguration struct {
	// The list of dimension values matching (in order) the dimensions of the price
	// group
	DimensionValues param.Field[[]string] `json:"dimension_values,required"`
	// The id of the dimensional price group to include this price in
	DimensionalPriceGroupID param.Field[string] `json:"dimensional_price_group_id"`
	// The external id of the dimensional price group to include this price in
	ExternalDimensionalPriceGroupID param.Field[string] `json:"external_dimensional_price_group_id"`
}

func (r PlanExternalPlanIDVersionNewParamsReplacePricesPriceNewPlanUnitPriceDimensionalPriceConfiguration) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Within each billing cycle, specifies the cadence at which invoices are produced.
// If unspecified, a single invoice is produced per billing cycle.
type PlanExternalPlanIDVersionNewParamsReplacePricesPriceNewPlanUnitPriceInvoicingCycleConfiguration struct {
	// The duration of the billing period.
	Duration param.Field[int64] `json:"duration,required"`
	// The unit of billing period duration.
	DurationUnit param.Field[PlanExternalPlanIDVersionNewParamsReplacePricesPriceNewPlanUnitPriceInvoicingCycleConfigurationDurationUnit] `json:"duration_unit,required"`
}

func (r PlanExternalPlanIDVersionNewParamsReplacePricesPriceNewPlanUnitPriceInvoicingCycleConfiguration) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// The unit of billing period duration.
type PlanExternalPlanIDVersionNewParamsReplacePricesPriceNewPlanUnitPriceInvoicingCycleConfigurationDurationUnit string

const (
	PlanExternalPlanIDVersionNewParamsReplacePricesPriceNewPlanUnitPriceInvoicingCycleConfigurationDurationUnitDay   PlanExternalPlanIDVersionNewParamsReplacePricesPriceNewPlanUnitPriceInvoicingCycleConfigurationDurationUnit = "day"
	PlanExternalPlanIDVersionNewParamsReplacePricesPriceNewPlanUnitPriceInvoicingCycleConfigurationDurationUnitMonth PlanExternalPlanIDVersionNewParamsReplacePricesPriceNewPlanUnitPriceInvoicingCycleConfigurationDurationUnit = "month"
)

func (r PlanExternalPlanIDVersionNewParamsReplacePricesPriceNewPlanUnitPriceInvoicingCycleConfigurationDurationUnit) IsKnown() bool {
	switch r {
	case PlanExternalPlanIDVersionNewParamsReplacePricesPriceNewPlanUnitPriceInvoicingCycleConfigurationDurationUnitDay, PlanExternalPlanIDVersionNewParamsReplacePricesPriceNewPlanUnitPriceInvoicingCycleConfigurationDurationUnitMonth:
		return true
	}
	return false
}

type PlanExternalPlanIDVersionNewParamsReplacePricesPriceNewPlanPackagePrice struct {
	// The cadence to bill for this price on.
	Cadence param.Field[PlanExternalPlanIDVersionNewParamsReplacePricesPriceNewPlanPackagePriceCadence] `json:"cadence,required"`
	// The id of the item the price will be associated with.
	ItemID    param.Field[string]                                                                           `json:"item_id,required"`
	ModelType param.Field[PlanExternalPlanIDVersionNewParamsReplacePricesPriceNewPlanPackagePriceModelType] `json:"model_type,required"`
	// The name of the price.
	Name          param.Field[string]                                                                               `json:"name,required"`
	PackageConfig param.Field[PlanExternalPlanIDVersionNewParamsReplacePricesPriceNewPlanPackagePricePackageConfig] `json:"package_config,required"`
	// The id of the billable metric for the price. Only needed if the price is
	// usage-based.
	BillableMetricID param.Field[string] `json:"billable_metric_id"`
	// If the Price represents a fixed cost, the price will be billed in-advance if
	// this is true, and in-arrears if this is false.
	BilledInAdvance param.Field[bool] `json:"billed_in_advance"`
	// For custom cadence: specifies the duration of the billing period in days or
	// months.
	BillingCycleConfiguration param.Field[PlanExternalPlanIDVersionNewParamsReplacePricesPriceNewPlanPackagePriceBillingCycleConfiguration] `json:"billing_cycle_configuration"`
	// The per unit conversion rate of the price currency to the invoicing currency.
	ConversionRate param.Field[float64] `json:"conversion_rate"`
	// An ISO 4217 currency string, or custom pricing unit identifier, in which this
	// price is billed.
	Currency param.Field[string] `json:"currency"`
	// For dimensional price: specifies a price group and dimension values
	DimensionalPriceConfiguration param.Field[PlanExternalPlanIDVersionNewParamsReplacePricesPriceNewPlanPackagePriceDimensionalPriceConfiguration] `json:"dimensional_price_configuration"`
	// An alias for the price.
	ExternalPriceID param.Field[string] `json:"external_price_id"`
	// If the Price represents a fixed cost, this represents the quantity of units
	// applied.
	FixedPriceQuantity param.Field[float64] `json:"fixed_price_quantity"`
	// The property used to group this price on an invoice
	InvoiceGroupingKey param.Field[string] `json:"invoice_grouping_key"`
	// Within each billing cycle, specifies the cadence at which invoices are produced.
	// If unspecified, a single invoice is produced per billing cycle.
	InvoicingCycleConfiguration param.Field[PlanExternalPlanIDVersionNewParamsReplacePricesPriceNewPlanPackagePriceInvoicingCycleConfiguration] `json:"invoicing_cycle_configuration"`
	// User-specified key/value pairs for the resource. Individual keys can be removed
	// by setting the value to `null`, and the entire metadata mapping can be cleared
	// by setting `metadata` to `null`.
	Metadata param.Field[map[string]string] `json:"metadata"`
}

func (r PlanExternalPlanIDVersionNewParamsReplacePricesPriceNewPlanPackagePrice) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r PlanExternalPlanIDVersionNewParamsReplacePricesPriceNewPlanPackagePrice) implementsPlanExternalPlanIDVersionNewParamsReplacePricesPriceUnion() {
}

// The cadence to bill for this price on.
type PlanExternalPlanIDVersionNewParamsReplacePricesPriceNewPlanPackagePriceCadence string

const (
	PlanExternalPlanIDVersionNewParamsReplacePricesPriceNewPlanPackagePriceCadenceAnnual     PlanExternalPlanIDVersionNewParamsReplacePricesPriceNewPlanPackagePriceCadence = "annual"
	PlanExternalPlanIDVersionNewParamsReplacePricesPriceNewPlanPackagePriceCadenceSemiAnnual PlanExternalPlanIDVersionNewParamsReplacePricesPriceNewPlanPackagePriceCadence = "semi_annual"
	PlanExternalPlanIDVersionNewParamsReplacePricesPriceNewPlanPackagePriceCadenceMonthly    PlanExternalPlanIDVersionNewParamsReplacePricesPriceNewPlanPackagePriceCadence = "monthly"
	PlanExternalPlanIDVersionNewParamsReplacePricesPriceNewPlanPackagePriceCadenceQuarterly  PlanExternalPlanIDVersionNewParamsReplacePricesPriceNewPlanPackagePriceCadence = "quarterly"
	PlanExternalPlanIDVersionNewParamsReplacePricesPriceNewPlanPackagePriceCadenceOneTime    PlanExternalPlanIDVersionNewParamsReplacePricesPriceNewPlanPackagePriceCadence = "one_time"
	PlanExternalPlanIDVersionNewParamsReplacePricesPriceNewPlanPackagePriceCadenceCustom     PlanExternalPlanIDVersionNewParamsReplacePricesPriceNewPlanPackagePriceCadence = "custom"
)

func (r PlanExternalPlanIDVersionNewParamsReplacePricesPriceNewPlanPackagePriceCadence) IsKnown() bool {
	switch r {
	case PlanExternalPlanIDVersionNewParamsReplacePricesPriceNewPlanPackagePriceCadenceAnnual, PlanExternalPlanIDVersionNewParamsReplacePricesPriceNewPlanPackagePriceCadenceSemiAnnual, PlanExternalPlanIDVersionNewParamsReplacePricesPriceNewPlanPackagePriceCadenceMonthly, PlanExternalPlanIDVersionNewParamsReplacePricesPriceNewPlanPackagePriceCadenceQuarterly, PlanExternalPlanIDVersionNewParamsReplacePricesPriceNewPlanPackagePriceCadenceOneTime, PlanExternalPlanIDVersionNewParamsReplacePricesPriceNewPlanPackagePriceCadenceCustom:
		return true
	}
	return false
}

type PlanExternalPlanIDVersionNewParamsReplacePricesPriceNewPlanPackagePriceModelType string

const (
	PlanExternalPlanIDVersionNewParamsReplacePricesPriceNewPlanPackagePriceModelTypePackage PlanExternalPlanIDVersionNewParamsReplacePricesPriceNewPlanPackagePriceModelType = "package"
)

func (r PlanExternalPlanIDVersionNewParamsReplacePricesPriceNewPlanPackagePriceModelType) IsKnown() bool {
	switch r {
	case PlanExternalPlanIDVersionNewParamsReplacePricesPriceNewPlanPackagePriceModelTypePackage:
		return true
	}
	return false
}

type PlanExternalPlanIDVersionNewParamsReplacePricesPriceNewPlanPackagePricePackageConfig struct {
	// A currency amount to rate usage by
	PackageAmount param.Field[string] `json:"package_amount,required"`
	// An integer amount to represent package size. For example, 1000 here would divide
	// usage by 1000 before multiplying by package_amount in rating
	PackageSize param.Field[int64] `json:"package_size,required"`
}

func (r PlanExternalPlanIDVersionNewParamsReplacePricesPriceNewPlanPackagePricePackageConfig) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// For custom cadence: specifies the duration of the billing period in days or
// months.
type PlanExternalPlanIDVersionNewParamsReplacePricesPriceNewPlanPackagePriceBillingCycleConfiguration struct {
	// The duration of the billing period.
	Duration param.Field[int64] `json:"duration,required"`
	// The unit of billing period duration.
	DurationUnit param.Field[PlanExternalPlanIDVersionNewParamsReplacePricesPriceNewPlanPackagePriceBillingCycleConfigurationDurationUnit] `json:"duration_unit,required"`
}

func (r PlanExternalPlanIDVersionNewParamsReplacePricesPriceNewPlanPackagePriceBillingCycleConfiguration) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// The unit of billing period duration.
type PlanExternalPlanIDVersionNewParamsReplacePricesPriceNewPlanPackagePriceBillingCycleConfigurationDurationUnit string

const (
	PlanExternalPlanIDVersionNewParamsReplacePricesPriceNewPlanPackagePriceBillingCycleConfigurationDurationUnitDay   PlanExternalPlanIDVersionNewParamsReplacePricesPriceNewPlanPackagePriceBillingCycleConfigurationDurationUnit = "day"
	PlanExternalPlanIDVersionNewParamsReplacePricesPriceNewPlanPackagePriceBillingCycleConfigurationDurationUnitMonth PlanExternalPlanIDVersionNewParamsReplacePricesPriceNewPlanPackagePriceBillingCycleConfigurationDurationUnit = "month"
)

func (r PlanExternalPlanIDVersionNewParamsReplacePricesPriceNewPlanPackagePriceBillingCycleConfigurationDurationUnit) IsKnown() bool {
	switch r {
	case PlanExternalPlanIDVersionNewParamsReplacePricesPriceNewPlanPackagePriceBillingCycleConfigurationDurationUnitDay, PlanExternalPlanIDVersionNewParamsReplacePricesPriceNewPlanPackagePriceBillingCycleConfigurationDurationUnitMonth:
		return true
	}
	return false
}

// For dimensional price: specifies a price group and dimension values
type PlanExternalPlanIDVersionNewParamsReplacePricesPriceNewPlanPackagePriceDimensionalPriceConfiguration struct {
	// The list of dimension values matching (in order) the dimensions of the price
	// group
	DimensionValues param.Field[[]string] `json:"dimension_values,required"`
	// The id of the dimensional price group to include this price in
	DimensionalPriceGroupID param.Field[string] `json:"dimensional_price_group_id"`
	// The external id of the dimensional price group to include this price in
	ExternalDimensionalPriceGroupID param.Field[string] `json:"external_dimensional_price_group_id"`
}

func (r PlanExternalPlanIDVersionNewParamsReplacePricesPriceNewPlanPackagePriceDimensionalPriceConfiguration) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Within each billing cycle, specifies the cadence at which invoices are produced.
// If unspecified, a single invoice is produced per billing cycle.
type PlanExternalPlanIDVersionNewParamsReplacePricesPriceNewPlanPackagePriceInvoicingCycleConfiguration struct {
	// The duration of the billing period.
	Duration param.Field[int64] `json:"duration,required"`
	// The unit of billing period duration.
	DurationUnit param.Field[PlanExternalPlanIDVersionNewParamsReplacePricesPriceNewPlanPackagePriceInvoicingCycleConfigurationDurationUnit] `json:"duration_unit,required"`
}

func (r PlanExternalPlanIDVersionNewParamsReplacePricesPriceNewPlanPackagePriceInvoicingCycleConfiguration) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// The unit of billing period duration.
type PlanExternalPlanIDVersionNewParamsReplacePricesPriceNewPlanPackagePriceInvoicingCycleConfigurationDurationUnit string

const (
	PlanExternalPlanIDVersionNewParamsReplacePricesPriceNewPlanPackagePriceInvoicingCycleConfigurationDurationUnitDay   PlanExternalPlanIDVersionNewParamsReplacePricesPriceNewPlanPackagePriceInvoicingCycleConfigurationDurationUnit = "day"
	PlanExternalPlanIDVersionNewParamsReplacePricesPriceNewPlanPackagePriceInvoicingCycleConfigurationDurationUnitMonth PlanExternalPlanIDVersionNewParamsReplacePricesPriceNewPlanPackagePriceInvoicingCycleConfigurationDurationUnit = "month"
)

func (r PlanExternalPlanIDVersionNewParamsReplacePricesPriceNewPlanPackagePriceInvoicingCycleConfigurationDurationUnit) IsKnown() bool {
	switch r {
	case PlanExternalPlanIDVersionNewParamsReplacePricesPriceNewPlanPackagePriceInvoicingCycleConfigurationDurationUnitDay, PlanExternalPlanIDVersionNewParamsReplacePricesPriceNewPlanPackagePriceInvoicingCycleConfigurationDurationUnitMonth:
		return true
	}
	return false
}

type PlanExternalPlanIDVersionNewParamsReplacePricesPriceNewPlanMatrixPrice struct {
	// The cadence to bill for this price on.
	Cadence param.Field[PlanExternalPlanIDVersionNewParamsReplacePricesPriceNewPlanMatrixPriceCadence] `json:"cadence,required"`
	// The id of the item the price will be associated with.
	ItemID       param.Field[string]                                                                             `json:"item_id,required"`
	MatrixConfig param.Field[PlanExternalPlanIDVersionNewParamsReplacePricesPriceNewPlanMatrixPriceMatrixConfig] `json:"matrix_config,required"`
	ModelType    param.Field[PlanExternalPlanIDVersionNewParamsReplacePricesPriceNewPlanMatrixPriceModelType]    `json:"model_type,required"`
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
	BillingCycleConfiguration param.Field[PlanExternalPlanIDVersionNewParamsReplacePricesPriceNewPlanMatrixPriceBillingCycleConfiguration] `json:"billing_cycle_configuration"`
	// The per unit conversion rate of the price currency to the invoicing currency.
	ConversionRate param.Field[float64] `json:"conversion_rate"`
	// An ISO 4217 currency string, or custom pricing unit identifier, in which this
	// price is billed.
	Currency param.Field[string] `json:"currency"`
	// For dimensional price: specifies a price group and dimension values
	DimensionalPriceConfiguration param.Field[PlanExternalPlanIDVersionNewParamsReplacePricesPriceNewPlanMatrixPriceDimensionalPriceConfiguration] `json:"dimensional_price_configuration"`
	// An alias for the price.
	ExternalPriceID param.Field[string] `json:"external_price_id"`
	// If the Price represents a fixed cost, this represents the quantity of units
	// applied.
	FixedPriceQuantity param.Field[float64] `json:"fixed_price_quantity"`
	// The property used to group this price on an invoice
	InvoiceGroupingKey param.Field[string] `json:"invoice_grouping_key"`
	// Within each billing cycle, specifies the cadence at which invoices are produced.
	// If unspecified, a single invoice is produced per billing cycle.
	InvoicingCycleConfiguration param.Field[PlanExternalPlanIDVersionNewParamsReplacePricesPriceNewPlanMatrixPriceInvoicingCycleConfiguration] `json:"invoicing_cycle_configuration"`
	// User-specified key/value pairs for the resource. Individual keys can be removed
	// by setting the value to `null`, and the entire metadata mapping can be cleared
	// by setting `metadata` to `null`.
	Metadata param.Field[map[string]string] `json:"metadata"`
}

func (r PlanExternalPlanIDVersionNewParamsReplacePricesPriceNewPlanMatrixPrice) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r PlanExternalPlanIDVersionNewParamsReplacePricesPriceNewPlanMatrixPrice) implementsPlanExternalPlanIDVersionNewParamsReplacePricesPriceUnion() {
}

// The cadence to bill for this price on.
type PlanExternalPlanIDVersionNewParamsReplacePricesPriceNewPlanMatrixPriceCadence string

const (
	PlanExternalPlanIDVersionNewParamsReplacePricesPriceNewPlanMatrixPriceCadenceAnnual     PlanExternalPlanIDVersionNewParamsReplacePricesPriceNewPlanMatrixPriceCadence = "annual"
	PlanExternalPlanIDVersionNewParamsReplacePricesPriceNewPlanMatrixPriceCadenceSemiAnnual PlanExternalPlanIDVersionNewParamsReplacePricesPriceNewPlanMatrixPriceCadence = "semi_annual"
	PlanExternalPlanIDVersionNewParamsReplacePricesPriceNewPlanMatrixPriceCadenceMonthly    PlanExternalPlanIDVersionNewParamsReplacePricesPriceNewPlanMatrixPriceCadence = "monthly"
	PlanExternalPlanIDVersionNewParamsReplacePricesPriceNewPlanMatrixPriceCadenceQuarterly  PlanExternalPlanIDVersionNewParamsReplacePricesPriceNewPlanMatrixPriceCadence = "quarterly"
	PlanExternalPlanIDVersionNewParamsReplacePricesPriceNewPlanMatrixPriceCadenceOneTime    PlanExternalPlanIDVersionNewParamsReplacePricesPriceNewPlanMatrixPriceCadence = "one_time"
	PlanExternalPlanIDVersionNewParamsReplacePricesPriceNewPlanMatrixPriceCadenceCustom     PlanExternalPlanIDVersionNewParamsReplacePricesPriceNewPlanMatrixPriceCadence = "custom"
)

func (r PlanExternalPlanIDVersionNewParamsReplacePricesPriceNewPlanMatrixPriceCadence) IsKnown() bool {
	switch r {
	case PlanExternalPlanIDVersionNewParamsReplacePricesPriceNewPlanMatrixPriceCadenceAnnual, PlanExternalPlanIDVersionNewParamsReplacePricesPriceNewPlanMatrixPriceCadenceSemiAnnual, PlanExternalPlanIDVersionNewParamsReplacePricesPriceNewPlanMatrixPriceCadenceMonthly, PlanExternalPlanIDVersionNewParamsReplacePricesPriceNewPlanMatrixPriceCadenceQuarterly, PlanExternalPlanIDVersionNewParamsReplacePricesPriceNewPlanMatrixPriceCadenceOneTime, PlanExternalPlanIDVersionNewParamsReplacePricesPriceNewPlanMatrixPriceCadenceCustom:
		return true
	}
	return false
}

type PlanExternalPlanIDVersionNewParamsReplacePricesPriceNewPlanMatrixPriceMatrixConfig struct {
	// Default per unit rate for any usage not bucketed into a specified matrix_value
	DefaultUnitAmount param.Field[string] `json:"default_unit_amount,required"`
	// One or two event property values to evaluate matrix groups by
	Dimensions param.Field[[]string] `json:"dimensions,required"`
	// Matrix values for specified matrix grouping keys
	MatrixValues param.Field[[]PlanExternalPlanIDVersionNewParamsReplacePricesPriceNewPlanMatrixPriceMatrixConfigMatrixValue] `json:"matrix_values,required"`
}

func (r PlanExternalPlanIDVersionNewParamsReplacePricesPriceNewPlanMatrixPriceMatrixConfig) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type PlanExternalPlanIDVersionNewParamsReplacePricesPriceNewPlanMatrixPriceMatrixConfigMatrixValue struct {
	// One or two matrix keys to filter usage to this Matrix value by. For example,
	// ["region", "tier"] could be used to filter cloud usage by a cloud region and an
	// instance tier.
	DimensionValues param.Field[[]string] `json:"dimension_values,required"`
	// Unit price for the specified dimension_values
	UnitAmount param.Field[string] `json:"unit_amount,required"`
}

func (r PlanExternalPlanIDVersionNewParamsReplacePricesPriceNewPlanMatrixPriceMatrixConfigMatrixValue) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type PlanExternalPlanIDVersionNewParamsReplacePricesPriceNewPlanMatrixPriceModelType string

const (
	PlanExternalPlanIDVersionNewParamsReplacePricesPriceNewPlanMatrixPriceModelTypeMatrix PlanExternalPlanIDVersionNewParamsReplacePricesPriceNewPlanMatrixPriceModelType = "matrix"
)

func (r PlanExternalPlanIDVersionNewParamsReplacePricesPriceNewPlanMatrixPriceModelType) IsKnown() bool {
	switch r {
	case PlanExternalPlanIDVersionNewParamsReplacePricesPriceNewPlanMatrixPriceModelTypeMatrix:
		return true
	}
	return false
}

// For custom cadence: specifies the duration of the billing period in days or
// months.
type PlanExternalPlanIDVersionNewParamsReplacePricesPriceNewPlanMatrixPriceBillingCycleConfiguration struct {
	// The duration of the billing period.
	Duration param.Field[int64] `json:"duration,required"`
	// The unit of billing period duration.
	DurationUnit param.Field[PlanExternalPlanIDVersionNewParamsReplacePricesPriceNewPlanMatrixPriceBillingCycleConfigurationDurationUnit] `json:"duration_unit,required"`
}

func (r PlanExternalPlanIDVersionNewParamsReplacePricesPriceNewPlanMatrixPriceBillingCycleConfiguration) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// The unit of billing period duration.
type PlanExternalPlanIDVersionNewParamsReplacePricesPriceNewPlanMatrixPriceBillingCycleConfigurationDurationUnit string

const (
	PlanExternalPlanIDVersionNewParamsReplacePricesPriceNewPlanMatrixPriceBillingCycleConfigurationDurationUnitDay   PlanExternalPlanIDVersionNewParamsReplacePricesPriceNewPlanMatrixPriceBillingCycleConfigurationDurationUnit = "day"
	PlanExternalPlanIDVersionNewParamsReplacePricesPriceNewPlanMatrixPriceBillingCycleConfigurationDurationUnitMonth PlanExternalPlanIDVersionNewParamsReplacePricesPriceNewPlanMatrixPriceBillingCycleConfigurationDurationUnit = "month"
)

func (r PlanExternalPlanIDVersionNewParamsReplacePricesPriceNewPlanMatrixPriceBillingCycleConfigurationDurationUnit) IsKnown() bool {
	switch r {
	case PlanExternalPlanIDVersionNewParamsReplacePricesPriceNewPlanMatrixPriceBillingCycleConfigurationDurationUnitDay, PlanExternalPlanIDVersionNewParamsReplacePricesPriceNewPlanMatrixPriceBillingCycleConfigurationDurationUnitMonth:
		return true
	}
	return false
}

// For dimensional price: specifies a price group and dimension values
type PlanExternalPlanIDVersionNewParamsReplacePricesPriceNewPlanMatrixPriceDimensionalPriceConfiguration struct {
	// The list of dimension values matching (in order) the dimensions of the price
	// group
	DimensionValues param.Field[[]string] `json:"dimension_values,required"`
	// The id of the dimensional price group to include this price in
	DimensionalPriceGroupID param.Field[string] `json:"dimensional_price_group_id"`
	// The external id of the dimensional price group to include this price in
	ExternalDimensionalPriceGroupID param.Field[string] `json:"external_dimensional_price_group_id"`
}

func (r PlanExternalPlanIDVersionNewParamsReplacePricesPriceNewPlanMatrixPriceDimensionalPriceConfiguration) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Within each billing cycle, specifies the cadence at which invoices are produced.
// If unspecified, a single invoice is produced per billing cycle.
type PlanExternalPlanIDVersionNewParamsReplacePricesPriceNewPlanMatrixPriceInvoicingCycleConfiguration struct {
	// The duration of the billing period.
	Duration param.Field[int64] `json:"duration,required"`
	// The unit of billing period duration.
	DurationUnit param.Field[PlanExternalPlanIDVersionNewParamsReplacePricesPriceNewPlanMatrixPriceInvoicingCycleConfigurationDurationUnit] `json:"duration_unit,required"`
}

func (r PlanExternalPlanIDVersionNewParamsReplacePricesPriceNewPlanMatrixPriceInvoicingCycleConfiguration) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// The unit of billing period duration.
type PlanExternalPlanIDVersionNewParamsReplacePricesPriceNewPlanMatrixPriceInvoicingCycleConfigurationDurationUnit string

const (
	PlanExternalPlanIDVersionNewParamsReplacePricesPriceNewPlanMatrixPriceInvoicingCycleConfigurationDurationUnitDay   PlanExternalPlanIDVersionNewParamsReplacePricesPriceNewPlanMatrixPriceInvoicingCycleConfigurationDurationUnit = "day"
	PlanExternalPlanIDVersionNewParamsReplacePricesPriceNewPlanMatrixPriceInvoicingCycleConfigurationDurationUnitMonth PlanExternalPlanIDVersionNewParamsReplacePricesPriceNewPlanMatrixPriceInvoicingCycleConfigurationDurationUnit = "month"
)

func (r PlanExternalPlanIDVersionNewParamsReplacePricesPriceNewPlanMatrixPriceInvoicingCycleConfigurationDurationUnit) IsKnown() bool {
	switch r {
	case PlanExternalPlanIDVersionNewParamsReplacePricesPriceNewPlanMatrixPriceInvoicingCycleConfigurationDurationUnitDay, PlanExternalPlanIDVersionNewParamsReplacePricesPriceNewPlanMatrixPriceInvoicingCycleConfigurationDurationUnitMonth:
		return true
	}
	return false
}

type PlanExternalPlanIDVersionNewParamsReplacePricesPriceNewPlanTieredPrice struct {
	// The cadence to bill for this price on.
	Cadence param.Field[PlanExternalPlanIDVersionNewParamsReplacePricesPriceNewPlanTieredPriceCadence] `json:"cadence,required"`
	// The id of the item the price will be associated with.
	ItemID    param.Field[string]                                                                          `json:"item_id,required"`
	ModelType param.Field[PlanExternalPlanIDVersionNewParamsReplacePricesPriceNewPlanTieredPriceModelType] `json:"model_type,required"`
	// The name of the price.
	Name         param.Field[string]                                                                             `json:"name,required"`
	TieredConfig param.Field[PlanExternalPlanIDVersionNewParamsReplacePricesPriceNewPlanTieredPriceTieredConfig] `json:"tiered_config,required"`
	// The id of the billable metric for the price. Only needed if the price is
	// usage-based.
	BillableMetricID param.Field[string] `json:"billable_metric_id"`
	// If the Price represents a fixed cost, the price will be billed in-advance if
	// this is true, and in-arrears if this is false.
	BilledInAdvance param.Field[bool] `json:"billed_in_advance"`
	// For custom cadence: specifies the duration of the billing period in days or
	// months.
	BillingCycleConfiguration param.Field[PlanExternalPlanIDVersionNewParamsReplacePricesPriceNewPlanTieredPriceBillingCycleConfiguration] `json:"billing_cycle_configuration"`
	// The per unit conversion rate of the price currency to the invoicing currency.
	ConversionRate param.Field[float64] `json:"conversion_rate"`
	// An ISO 4217 currency string, or custom pricing unit identifier, in which this
	// price is billed.
	Currency param.Field[string] `json:"currency"`
	// For dimensional price: specifies a price group and dimension values
	DimensionalPriceConfiguration param.Field[PlanExternalPlanIDVersionNewParamsReplacePricesPriceNewPlanTieredPriceDimensionalPriceConfiguration] `json:"dimensional_price_configuration"`
	// An alias for the price.
	ExternalPriceID param.Field[string] `json:"external_price_id"`
	// If the Price represents a fixed cost, this represents the quantity of units
	// applied.
	FixedPriceQuantity param.Field[float64] `json:"fixed_price_quantity"`
	// The property used to group this price on an invoice
	InvoiceGroupingKey param.Field[string] `json:"invoice_grouping_key"`
	// Within each billing cycle, specifies the cadence at which invoices are produced.
	// If unspecified, a single invoice is produced per billing cycle.
	InvoicingCycleConfiguration param.Field[PlanExternalPlanIDVersionNewParamsReplacePricesPriceNewPlanTieredPriceInvoicingCycleConfiguration] `json:"invoicing_cycle_configuration"`
	// User-specified key/value pairs for the resource. Individual keys can be removed
	// by setting the value to `null`, and the entire metadata mapping can be cleared
	// by setting `metadata` to `null`.
	Metadata param.Field[map[string]string] `json:"metadata"`
}

func (r PlanExternalPlanIDVersionNewParamsReplacePricesPriceNewPlanTieredPrice) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r PlanExternalPlanIDVersionNewParamsReplacePricesPriceNewPlanTieredPrice) implementsPlanExternalPlanIDVersionNewParamsReplacePricesPriceUnion() {
}

// The cadence to bill for this price on.
type PlanExternalPlanIDVersionNewParamsReplacePricesPriceNewPlanTieredPriceCadence string

const (
	PlanExternalPlanIDVersionNewParamsReplacePricesPriceNewPlanTieredPriceCadenceAnnual     PlanExternalPlanIDVersionNewParamsReplacePricesPriceNewPlanTieredPriceCadence = "annual"
	PlanExternalPlanIDVersionNewParamsReplacePricesPriceNewPlanTieredPriceCadenceSemiAnnual PlanExternalPlanIDVersionNewParamsReplacePricesPriceNewPlanTieredPriceCadence = "semi_annual"
	PlanExternalPlanIDVersionNewParamsReplacePricesPriceNewPlanTieredPriceCadenceMonthly    PlanExternalPlanIDVersionNewParamsReplacePricesPriceNewPlanTieredPriceCadence = "monthly"
	PlanExternalPlanIDVersionNewParamsReplacePricesPriceNewPlanTieredPriceCadenceQuarterly  PlanExternalPlanIDVersionNewParamsReplacePricesPriceNewPlanTieredPriceCadence = "quarterly"
	PlanExternalPlanIDVersionNewParamsReplacePricesPriceNewPlanTieredPriceCadenceOneTime    PlanExternalPlanIDVersionNewParamsReplacePricesPriceNewPlanTieredPriceCadence = "one_time"
	PlanExternalPlanIDVersionNewParamsReplacePricesPriceNewPlanTieredPriceCadenceCustom     PlanExternalPlanIDVersionNewParamsReplacePricesPriceNewPlanTieredPriceCadence = "custom"
)

func (r PlanExternalPlanIDVersionNewParamsReplacePricesPriceNewPlanTieredPriceCadence) IsKnown() bool {
	switch r {
	case PlanExternalPlanIDVersionNewParamsReplacePricesPriceNewPlanTieredPriceCadenceAnnual, PlanExternalPlanIDVersionNewParamsReplacePricesPriceNewPlanTieredPriceCadenceSemiAnnual, PlanExternalPlanIDVersionNewParamsReplacePricesPriceNewPlanTieredPriceCadenceMonthly, PlanExternalPlanIDVersionNewParamsReplacePricesPriceNewPlanTieredPriceCadenceQuarterly, PlanExternalPlanIDVersionNewParamsReplacePricesPriceNewPlanTieredPriceCadenceOneTime, PlanExternalPlanIDVersionNewParamsReplacePricesPriceNewPlanTieredPriceCadenceCustom:
		return true
	}
	return false
}

type PlanExternalPlanIDVersionNewParamsReplacePricesPriceNewPlanTieredPriceModelType string

const (
	PlanExternalPlanIDVersionNewParamsReplacePricesPriceNewPlanTieredPriceModelTypeTiered PlanExternalPlanIDVersionNewParamsReplacePricesPriceNewPlanTieredPriceModelType = "tiered"
)

func (r PlanExternalPlanIDVersionNewParamsReplacePricesPriceNewPlanTieredPriceModelType) IsKnown() bool {
	switch r {
	case PlanExternalPlanIDVersionNewParamsReplacePricesPriceNewPlanTieredPriceModelTypeTiered:
		return true
	}
	return false
}

type PlanExternalPlanIDVersionNewParamsReplacePricesPriceNewPlanTieredPriceTieredConfig struct {
	// Tiers for rating based on total usage quantities into the specified tier
	Tiers param.Field[[]PlanExternalPlanIDVersionNewParamsReplacePricesPriceNewPlanTieredPriceTieredConfigTier] `json:"tiers,required"`
}

func (r PlanExternalPlanIDVersionNewParamsReplacePricesPriceNewPlanTieredPriceTieredConfig) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type PlanExternalPlanIDVersionNewParamsReplacePricesPriceNewPlanTieredPriceTieredConfigTier struct {
	// Exclusive tier starting value
	FirstUnit param.Field[float64] `json:"first_unit,required"`
	// Amount per unit
	UnitAmount param.Field[string] `json:"unit_amount,required"`
	// Inclusive tier ending value. If null, this is treated as the last tier
	LastUnit param.Field[float64] `json:"last_unit"`
}

func (r PlanExternalPlanIDVersionNewParamsReplacePricesPriceNewPlanTieredPriceTieredConfigTier) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// For custom cadence: specifies the duration of the billing period in days or
// months.
type PlanExternalPlanIDVersionNewParamsReplacePricesPriceNewPlanTieredPriceBillingCycleConfiguration struct {
	// The duration of the billing period.
	Duration param.Field[int64] `json:"duration,required"`
	// The unit of billing period duration.
	DurationUnit param.Field[PlanExternalPlanIDVersionNewParamsReplacePricesPriceNewPlanTieredPriceBillingCycleConfigurationDurationUnit] `json:"duration_unit,required"`
}

func (r PlanExternalPlanIDVersionNewParamsReplacePricesPriceNewPlanTieredPriceBillingCycleConfiguration) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// The unit of billing period duration.
type PlanExternalPlanIDVersionNewParamsReplacePricesPriceNewPlanTieredPriceBillingCycleConfigurationDurationUnit string

const (
	PlanExternalPlanIDVersionNewParamsReplacePricesPriceNewPlanTieredPriceBillingCycleConfigurationDurationUnitDay   PlanExternalPlanIDVersionNewParamsReplacePricesPriceNewPlanTieredPriceBillingCycleConfigurationDurationUnit = "day"
	PlanExternalPlanIDVersionNewParamsReplacePricesPriceNewPlanTieredPriceBillingCycleConfigurationDurationUnitMonth PlanExternalPlanIDVersionNewParamsReplacePricesPriceNewPlanTieredPriceBillingCycleConfigurationDurationUnit = "month"
)

func (r PlanExternalPlanIDVersionNewParamsReplacePricesPriceNewPlanTieredPriceBillingCycleConfigurationDurationUnit) IsKnown() bool {
	switch r {
	case PlanExternalPlanIDVersionNewParamsReplacePricesPriceNewPlanTieredPriceBillingCycleConfigurationDurationUnitDay, PlanExternalPlanIDVersionNewParamsReplacePricesPriceNewPlanTieredPriceBillingCycleConfigurationDurationUnitMonth:
		return true
	}
	return false
}

// For dimensional price: specifies a price group and dimension values
type PlanExternalPlanIDVersionNewParamsReplacePricesPriceNewPlanTieredPriceDimensionalPriceConfiguration struct {
	// The list of dimension values matching (in order) the dimensions of the price
	// group
	DimensionValues param.Field[[]string] `json:"dimension_values,required"`
	// The id of the dimensional price group to include this price in
	DimensionalPriceGroupID param.Field[string] `json:"dimensional_price_group_id"`
	// The external id of the dimensional price group to include this price in
	ExternalDimensionalPriceGroupID param.Field[string] `json:"external_dimensional_price_group_id"`
}

func (r PlanExternalPlanIDVersionNewParamsReplacePricesPriceNewPlanTieredPriceDimensionalPriceConfiguration) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Within each billing cycle, specifies the cadence at which invoices are produced.
// If unspecified, a single invoice is produced per billing cycle.
type PlanExternalPlanIDVersionNewParamsReplacePricesPriceNewPlanTieredPriceInvoicingCycleConfiguration struct {
	// The duration of the billing period.
	Duration param.Field[int64] `json:"duration,required"`
	// The unit of billing period duration.
	DurationUnit param.Field[PlanExternalPlanIDVersionNewParamsReplacePricesPriceNewPlanTieredPriceInvoicingCycleConfigurationDurationUnit] `json:"duration_unit,required"`
}

func (r PlanExternalPlanIDVersionNewParamsReplacePricesPriceNewPlanTieredPriceInvoicingCycleConfiguration) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// The unit of billing period duration.
type PlanExternalPlanIDVersionNewParamsReplacePricesPriceNewPlanTieredPriceInvoicingCycleConfigurationDurationUnit string

const (
	PlanExternalPlanIDVersionNewParamsReplacePricesPriceNewPlanTieredPriceInvoicingCycleConfigurationDurationUnitDay   PlanExternalPlanIDVersionNewParamsReplacePricesPriceNewPlanTieredPriceInvoicingCycleConfigurationDurationUnit = "day"
	PlanExternalPlanIDVersionNewParamsReplacePricesPriceNewPlanTieredPriceInvoicingCycleConfigurationDurationUnitMonth PlanExternalPlanIDVersionNewParamsReplacePricesPriceNewPlanTieredPriceInvoicingCycleConfigurationDurationUnit = "month"
)

func (r PlanExternalPlanIDVersionNewParamsReplacePricesPriceNewPlanTieredPriceInvoicingCycleConfigurationDurationUnit) IsKnown() bool {
	switch r {
	case PlanExternalPlanIDVersionNewParamsReplacePricesPriceNewPlanTieredPriceInvoicingCycleConfigurationDurationUnitDay, PlanExternalPlanIDVersionNewParamsReplacePricesPriceNewPlanTieredPriceInvoicingCycleConfigurationDurationUnitMonth:
		return true
	}
	return false
}

type PlanExternalPlanIDVersionNewParamsReplacePricesPriceNewPlanTieredBpsPrice struct {
	// The cadence to bill for this price on.
	Cadence param.Field[PlanExternalPlanIDVersionNewParamsReplacePricesPriceNewPlanTieredBpsPriceCadence] `json:"cadence,required"`
	// The id of the item the price will be associated with.
	ItemID    param.Field[string]                                                                             `json:"item_id,required"`
	ModelType param.Field[PlanExternalPlanIDVersionNewParamsReplacePricesPriceNewPlanTieredBpsPriceModelType] `json:"model_type,required"`
	// The name of the price.
	Name            param.Field[string]                                                                                   `json:"name,required"`
	TieredBpsConfig param.Field[PlanExternalPlanIDVersionNewParamsReplacePricesPriceNewPlanTieredBpsPriceTieredBpsConfig] `json:"tiered_bps_config,required"`
	// The id of the billable metric for the price. Only needed if the price is
	// usage-based.
	BillableMetricID param.Field[string] `json:"billable_metric_id"`
	// If the Price represents a fixed cost, the price will be billed in-advance if
	// this is true, and in-arrears if this is false.
	BilledInAdvance param.Field[bool] `json:"billed_in_advance"`
	// For custom cadence: specifies the duration of the billing period in days or
	// months.
	BillingCycleConfiguration param.Field[PlanExternalPlanIDVersionNewParamsReplacePricesPriceNewPlanTieredBpsPriceBillingCycleConfiguration] `json:"billing_cycle_configuration"`
	// The per unit conversion rate of the price currency to the invoicing currency.
	ConversionRate param.Field[float64] `json:"conversion_rate"`
	// An ISO 4217 currency string, or custom pricing unit identifier, in which this
	// price is billed.
	Currency param.Field[string] `json:"currency"`
	// For dimensional price: specifies a price group and dimension values
	DimensionalPriceConfiguration param.Field[PlanExternalPlanIDVersionNewParamsReplacePricesPriceNewPlanTieredBpsPriceDimensionalPriceConfiguration] `json:"dimensional_price_configuration"`
	// An alias for the price.
	ExternalPriceID param.Field[string] `json:"external_price_id"`
	// If the Price represents a fixed cost, this represents the quantity of units
	// applied.
	FixedPriceQuantity param.Field[float64] `json:"fixed_price_quantity"`
	// The property used to group this price on an invoice
	InvoiceGroupingKey param.Field[string] `json:"invoice_grouping_key"`
	// Within each billing cycle, specifies the cadence at which invoices are produced.
	// If unspecified, a single invoice is produced per billing cycle.
	InvoicingCycleConfiguration param.Field[PlanExternalPlanIDVersionNewParamsReplacePricesPriceNewPlanTieredBpsPriceInvoicingCycleConfiguration] `json:"invoicing_cycle_configuration"`
	// User-specified key/value pairs for the resource. Individual keys can be removed
	// by setting the value to `null`, and the entire metadata mapping can be cleared
	// by setting `metadata` to `null`.
	Metadata param.Field[map[string]string] `json:"metadata"`
}

func (r PlanExternalPlanIDVersionNewParamsReplacePricesPriceNewPlanTieredBpsPrice) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r PlanExternalPlanIDVersionNewParamsReplacePricesPriceNewPlanTieredBpsPrice) implementsPlanExternalPlanIDVersionNewParamsReplacePricesPriceUnion() {
}

// The cadence to bill for this price on.
type PlanExternalPlanIDVersionNewParamsReplacePricesPriceNewPlanTieredBpsPriceCadence string

const (
	PlanExternalPlanIDVersionNewParamsReplacePricesPriceNewPlanTieredBpsPriceCadenceAnnual     PlanExternalPlanIDVersionNewParamsReplacePricesPriceNewPlanTieredBpsPriceCadence = "annual"
	PlanExternalPlanIDVersionNewParamsReplacePricesPriceNewPlanTieredBpsPriceCadenceSemiAnnual PlanExternalPlanIDVersionNewParamsReplacePricesPriceNewPlanTieredBpsPriceCadence = "semi_annual"
	PlanExternalPlanIDVersionNewParamsReplacePricesPriceNewPlanTieredBpsPriceCadenceMonthly    PlanExternalPlanIDVersionNewParamsReplacePricesPriceNewPlanTieredBpsPriceCadence = "monthly"
	PlanExternalPlanIDVersionNewParamsReplacePricesPriceNewPlanTieredBpsPriceCadenceQuarterly  PlanExternalPlanIDVersionNewParamsReplacePricesPriceNewPlanTieredBpsPriceCadence = "quarterly"
	PlanExternalPlanIDVersionNewParamsReplacePricesPriceNewPlanTieredBpsPriceCadenceOneTime    PlanExternalPlanIDVersionNewParamsReplacePricesPriceNewPlanTieredBpsPriceCadence = "one_time"
	PlanExternalPlanIDVersionNewParamsReplacePricesPriceNewPlanTieredBpsPriceCadenceCustom     PlanExternalPlanIDVersionNewParamsReplacePricesPriceNewPlanTieredBpsPriceCadence = "custom"
)

func (r PlanExternalPlanIDVersionNewParamsReplacePricesPriceNewPlanTieredBpsPriceCadence) IsKnown() bool {
	switch r {
	case PlanExternalPlanIDVersionNewParamsReplacePricesPriceNewPlanTieredBpsPriceCadenceAnnual, PlanExternalPlanIDVersionNewParamsReplacePricesPriceNewPlanTieredBpsPriceCadenceSemiAnnual, PlanExternalPlanIDVersionNewParamsReplacePricesPriceNewPlanTieredBpsPriceCadenceMonthly, PlanExternalPlanIDVersionNewParamsReplacePricesPriceNewPlanTieredBpsPriceCadenceQuarterly, PlanExternalPlanIDVersionNewParamsReplacePricesPriceNewPlanTieredBpsPriceCadenceOneTime, PlanExternalPlanIDVersionNewParamsReplacePricesPriceNewPlanTieredBpsPriceCadenceCustom:
		return true
	}
	return false
}

type PlanExternalPlanIDVersionNewParamsReplacePricesPriceNewPlanTieredBpsPriceModelType string

const (
	PlanExternalPlanIDVersionNewParamsReplacePricesPriceNewPlanTieredBpsPriceModelTypeTieredBps PlanExternalPlanIDVersionNewParamsReplacePricesPriceNewPlanTieredBpsPriceModelType = "tiered_bps"
)

func (r PlanExternalPlanIDVersionNewParamsReplacePricesPriceNewPlanTieredBpsPriceModelType) IsKnown() bool {
	switch r {
	case PlanExternalPlanIDVersionNewParamsReplacePricesPriceNewPlanTieredBpsPriceModelTypeTieredBps:
		return true
	}
	return false
}

type PlanExternalPlanIDVersionNewParamsReplacePricesPriceNewPlanTieredBpsPriceTieredBpsConfig struct {
	// Tiers for a Graduated BPS pricing model, where usage is bucketed into specified
	// tiers
	Tiers param.Field[[]PlanExternalPlanIDVersionNewParamsReplacePricesPriceNewPlanTieredBpsPriceTieredBpsConfigTier] `json:"tiers,required"`
}

func (r PlanExternalPlanIDVersionNewParamsReplacePricesPriceNewPlanTieredBpsPriceTieredBpsConfig) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type PlanExternalPlanIDVersionNewParamsReplacePricesPriceNewPlanTieredBpsPriceTieredBpsConfigTier struct {
	// Per-event basis point rate
	Bps param.Field[float64] `json:"bps,required"`
	// Exclusive tier starting value
	MinimumAmount param.Field[string] `json:"minimum_amount,required"`
	// Inclusive tier ending value
	MaximumAmount param.Field[string] `json:"maximum_amount"`
	// Per unit maximum to charge
	PerUnitMaximum param.Field[string] `json:"per_unit_maximum"`
}

func (r PlanExternalPlanIDVersionNewParamsReplacePricesPriceNewPlanTieredBpsPriceTieredBpsConfigTier) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// For custom cadence: specifies the duration of the billing period in days or
// months.
type PlanExternalPlanIDVersionNewParamsReplacePricesPriceNewPlanTieredBpsPriceBillingCycleConfiguration struct {
	// The duration of the billing period.
	Duration param.Field[int64] `json:"duration,required"`
	// The unit of billing period duration.
	DurationUnit param.Field[PlanExternalPlanIDVersionNewParamsReplacePricesPriceNewPlanTieredBpsPriceBillingCycleConfigurationDurationUnit] `json:"duration_unit,required"`
}

func (r PlanExternalPlanIDVersionNewParamsReplacePricesPriceNewPlanTieredBpsPriceBillingCycleConfiguration) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// The unit of billing period duration.
type PlanExternalPlanIDVersionNewParamsReplacePricesPriceNewPlanTieredBpsPriceBillingCycleConfigurationDurationUnit string

const (
	PlanExternalPlanIDVersionNewParamsReplacePricesPriceNewPlanTieredBpsPriceBillingCycleConfigurationDurationUnitDay   PlanExternalPlanIDVersionNewParamsReplacePricesPriceNewPlanTieredBpsPriceBillingCycleConfigurationDurationUnit = "day"
	PlanExternalPlanIDVersionNewParamsReplacePricesPriceNewPlanTieredBpsPriceBillingCycleConfigurationDurationUnitMonth PlanExternalPlanIDVersionNewParamsReplacePricesPriceNewPlanTieredBpsPriceBillingCycleConfigurationDurationUnit = "month"
)

func (r PlanExternalPlanIDVersionNewParamsReplacePricesPriceNewPlanTieredBpsPriceBillingCycleConfigurationDurationUnit) IsKnown() bool {
	switch r {
	case PlanExternalPlanIDVersionNewParamsReplacePricesPriceNewPlanTieredBpsPriceBillingCycleConfigurationDurationUnitDay, PlanExternalPlanIDVersionNewParamsReplacePricesPriceNewPlanTieredBpsPriceBillingCycleConfigurationDurationUnitMonth:
		return true
	}
	return false
}

// For dimensional price: specifies a price group and dimension values
type PlanExternalPlanIDVersionNewParamsReplacePricesPriceNewPlanTieredBpsPriceDimensionalPriceConfiguration struct {
	// The list of dimension values matching (in order) the dimensions of the price
	// group
	DimensionValues param.Field[[]string] `json:"dimension_values,required"`
	// The id of the dimensional price group to include this price in
	DimensionalPriceGroupID param.Field[string] `json:"dimensional_price_group_id"`
	// The external id of the dimensional price group to include this price in
	ExternalDimensionalPriceGroupID param.Field[string] `json:"external_dimensional_price_group_id"`
}

func (r PlanExternalPlanIDVersionNewParamsReplacePricesPriceNewPlanTieredBpsPriceDimensionalPriceConfiguration) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Within each billing cycle, specifies the cadence at which invoices are produced.
// If unspecified, a single invoice is produced per billing cycle.
type PlanExternalPlanIDVersionNewParamsReplacePricesPriceNewPlanTieredBpsPriceInvoicingCycleConfiguration struct {
	// The duration of the billing period.
	Duration param.Field[int64] `json:"duration,required"`
	// The unit of billing period duration.
	DurationUnit param.Field[PlanExternalPlanIDVersionNewParamsReplacePricesPriceNewPlanTieredBpsPriceInvoicingCycleConfigurationDurationUnit] `json:"duration_unit,required"`
}

func (r PlanExternalPlanIDVersionNewParamsReplacePricesPriceNewPlanTieredBpsPriceInvoicingCycleConfiguration) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// The unit of billing period duration.
type PlanExternalPlanIDVersionNewParamsReplacePricesPriceNewPlanTieredBpsPriceInvoicingCycleConfigurationDurationUnit string

const (
	PlanExternalPlanIDVersionNewParamsReplacePricesPriceNewPlanTieredBpsPriceInvoicingCycleConfigurationDurationUnitDay   PlanExternalPlanIDVersionNewParamsReplacePricesPriceNewPlanTieredBpsPriceInvoicingCycleConfigurationDurationUnit = "day"
	PlanExternalPlanIDVersionNewParamsReplacePricesPriceNewPlanTieredBpsPriceInvoicingCycleConfigurationDurationUnitMonth PlanExternalPlanIDVersionNewParamsReplacePricesPriceNewPlanTieredBpsPriceInvoicingCycleConfigurationDurationUnit = "month"
)

func (r PlanExternalPlanIDVersionNewParamsReplacePricesPriceNewPlanTieredBpsPriceInvoicingCycleConfigurationDurationUnit) IsKnown() bool {
	switch r {
	case PlanExternalPlanIDVersionNewParamsReplacePricesPriceNewPlanTieredBpsPriceInvoicingCycleConfigurationDurationUnitDay, PlanExternalPlanIDVersionNewParamsReplacePricesPriceNewPlanTieredBpsPriceInvoicingCycleConfigurationDurationUnitMonth:
		return true
	}
	return false
}

type PlanExternalPlanIDVersionNewParamsReplacePricesPriceNewPlanBpsPrice struct {
	BpsConfig param.Field[PlanExternalPlanIDVersionNewParamsReplacePricesPriceNewPlanBpsPriceBpsConfig] `json:"bps_config,required"`
	// The cadence to bill for this price on.
	Cadence param.Field[PlanExternalPlanIDVersionNewParamsReplacePricesPriceNewPlanBpsPriceCadence] `json:"cadence,required"`
	// The id of the item the price will be associated with.
	ItemID    param.Field[string]                                                                       `json:"item_id,required"`
	ModelType param.Field[PlanExternalPlanIDVersionNewParamsReplacePricesPriceNewPlanBpsPriceModelType] `json:"model_type,required"`
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
	BillingCycleConfiguration param.Field[PlanExternalPlanIDVersionNewParamsReplacePricesPriceNewPlanBpsPriceBillingCycleConfiguration] `json:"billing_cycle_configuration"`
	// The per unit conversion rate of the price currency to the invoicing currency.
	ConversionRate param.Field[float64] `json:"conversion_rate"`
	// An ISO 4217 currency string, or custom pricing unit identifier, in which this
	// price is billed.
	Currency param.Field[string] `json:"currency"`
	// For dimensional price: specifies a price group and dimension values
	DimensionalPriceConfiguration param.Field[PlanExternalPlanIDVersionNewParamsReplacePricesPriceNewPlanBpsPriceDimensionalPriceConfiguration] `json:"dimensional_price_configuration"`
	// An alias for the price.
	ExternalPriceID param.Field[string] `json:"external_price_id"`
	// If the Price represents a fixed cost, this represents the quantity of units
	// applied.
	FixedPriceQuantity param.Field[float64] `json:"fixed_price_quantity"`
	// The property used to group this price on an invoice
	InvoiceGroupingKey param.Field[string] `json:"invoice_grouping_key"`
	// Within each billing cycle, specifies the cadence at which invoices are produced.
	// If unspecified, a single invoice is produced per billing cycle.
	InvoicingCycleConfiguration param.Field[PlanExternalPlanIDVersionNewParamsReplacePricesPriceNewPlanBpsPriceInvoicingCycleConfiguration] `json:"invoicing_cycle_configuration"`
	// User-specified key/value pairs for the resource. Individual keys can be removed
	// by setting the value to `null`, and the entire metadata mapping can be cleared
	// by setting `metadata` to `null`.
	Metadata param.Field[map[string]string] `json:"metadata"`
}

func (r PlanExternalPlanIDVersionNewParamsReplacePricesPriceNewPlanBpsPrice) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r PlanExternalPlanIDVersionNewParamsReplacePricesPriceNewPlanBpsPrice) implementsPlanExternalPlanIDVersionNewParamsReplacePricesPriceUnion() {
}

type PlanExternalPlanIDVersionNewParamsReplacePricesPriceNewPlanBpsPriceBpsConfig struct {
	// Basis point take rate per event
	Bps param.Field[float64] `json:"bps,required"`
	// Optional currency amount maximum to cap spend per event
	PerUnitMaximum param.Field[string] `json:"per_unit_maximum"`
}

func (r PlanExternalPlanIDVersionNewParamsReplacePricesPriceNewPlanBpsPriceBpsConfig) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// The cadence to bill for this price on.
type PlanExternalPlanIDVersionNewParamsReplacePricesPriceNewPlanBpsPriceCadence string

const (
	PlanExternalPlanIDVersionNewParamsReplacePricesPriceNewPlanBpsPriceCadenceAnnual     PlanExternalPlanIDVersionNewParamsReplacePricesPriceNewPlanBpsPriceCadence = "annual"
	PlanExternalPlanIDVersionNewParamsReplacePricesPriceNewPlanBpsPriceCadenceSemiAnnual PlanExternalPlanIDVersionNewParamsReplacePricesPriceNewPlanBpsPriceCadence = "semi_annual"
	PlanExternalPlanIDVersionNewParamsReplacePricesPriceNewPlanBpsPriceCadenceMonthly    PlanExternalPlanIDVersionNewParamsReplacePricesPriceNewPlanBpsPriceCadence = "monthly"
	PlanExternalPlanIDVersionNewParamsReplacePricesPriceNewPlanBpsPriceCadenceQuarterly  PlanExternalPlanIDVersionNewParamsReplacePricesPriceNewPlanBpsPriceCadence = "quarterly"
	PlanExternalPlanIDVersionNewParamsReplacePricesPriceNewPlanBpsPriceCadenceOneTime    PlanExternalPlanIDVersionNewParamsReplacePricesPriceNewPlanBpsPriceCadence = "one_time"
	PlanExternalPlanIDVersionNewParamsReplacePricesPriceNewPlanBpsPriceCadenceCustom     PlanExternalPlanIDVersionNewParamsReplacePricesPriceNewPlanBpsPriceCadence = "custom"
)

func (r PlanExternalPlanIDVersionNewParamsReplacePricesPriceNewPlanBpsPriceCadence) IsKnown() bool {
	switch r {
	case PlanExternalPlanIDVersionNewParamsReplacePricesPriceNewPlanBpsPriceCadenceAnnual, PlanExternalPlanIDVersionNewParamsReplacePricesPriceNewPlanBpsPriceCadenceSemiAnnual, PlanExternalPlanIDVersionNewParamsReplacePricesPriceNewPlanBpsPriceCadenceMonthly, PlanExternalPlanIDVersionNewParamsReplacePricesPriceNewPlanBpsPriceCadenceQuarterly, PlanExternalPlanIDVersionNewParamsReplacePricesPriceNewPlanBpsPriceCadenceOneTime, PlanExternalPlanIDVersionNewParamsReplacePricesPriceNewPlanBpsPriceCadenceCustom:
		return true
	}
	return false
}

type PlanExternalPlanIDVersionNewParamsReplacePricesPriceNewPlanBpsPriceModelType string

const (
	PlanExternalPlanIDVersionNewParamsReplacePricesPriceNewPlanBpsPriceModelTypeBps PlanExternalPlanIDVersionNewParamsReplacePricesPriceNewPlanBpsPriceModelType = "bps"
)

func (r PlanExternalPlanIDVersionNewParamsReplacePricesPriceNewPlanBpsPriceModelType) IsKnown() bool {
	switch r {
	case PlanExternalPlanIDVersionNewParamsReplacePricesPriceNewPlanBpsPriceModelTypeBps:
		return true
	}
	return false
}

// For custom cadence: specifies the duration of the billing period in days or
// months.
type PlanExternalPlanIDVersionNewParamsReplacePricesPriceNewPlanBpsPriceBillingCycleConfiguration struct {
	// The duration of the billing period.
	Duration param.Field[int64] `json:"duration,required"`
	// The unit of billing period duration.
	DurationUnit param.Field[PlanExternalPlanIDVersionNewParamsReplacePricesPriceNewPlanBpsPriceBillingCycleConfigurationDurationUnit] `json:"duration_unit,required"`
}

func (r PlanExternalPlanIDVersionNewParamsReplacePricesPriceNewPlanBpsPriceBillingCycleConfiguration) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// The unit of billing period duration.
type PlanExternalPlanIDVersionNewParamsReplacePricesPriceNewPlanBpsPriceBillingCycleConfigurationDurationUnit string

const (
	PlanExternalPlanIDVersionNewParamsReplacePricesPriceNewPlanBpsPriceBillingCycleConfigurationDurationUnitDay   PlanExternalPlanIDVersionNewParamsReplacePricesPriceNewPlanBpsPriceBillingCycleConfigurationDurationUnit = "day"
	PlanExternalPlanIDVersionNewParamsReplacePricesPriceNewPlanBpsPriceBillingCycleConfigurationDurationUnitMonth PlanExternalPlanIDVersionNewParamsReplacePricesPriceNewPlanBpsPriceBillingCycleConfigurationDurationUnit = "month"
)

func (r PlanExternalPlanIDVersionNewParamsReplacePricesPriceNewPlanBpsPriceBillingCycleConfigurationDurationUnit) IsKnown() bool {
	switch r {
	case PlanExternalPlanIDVersionNewParamsReplacePricesPriceNewPlanBpsPriceBillingCycleConfigurationDurationUnitDay, PlanExternalPlanIDVersionNewParamsReplacePricesPriceNewPlanBpsPriceBillingCycleConfigurationDurationUnitMonth:
		return true
	}
	return false
}

// For dimensional price: specifies a price group and dimension values
type PlanExternalPlanIDVersionNewParamsReplacePricesPriceNewPlanBpsPriceDimensionalPriceConfiguration struct {
	// The list of dimension values matching (in order) the dimensions of the price
	// group
	DimensionValues param.Field[[]string] `json:"dimension_values,required"`
	// The id of the dimensional price group to include this price in
	DimensionalPriceGroupID param.Field[string] `json:"dimensional_price_group_id"`
	// The external id of the dimensional price group to include this price in
	ExternalDimensionalPriceGroupID param.Field[string] `json:"external_dimensional_price_group_id"`
}

func (r PlanExternalPlanIDVersionNewParamsReplacePricesPriceNewPlanBpsPriceDimensionalPriceConfiguration) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Within each billing cycle, specifies the cadence at which invoices are produced.
// If unspecified, a single invoice is produced per billing cycle.
type PlanExternalPlanIDVersionNewParamsReplacePricesPriceNewPlanBpsPriceInvoicingCycleConfiguration struct {
	// The duration of the billing period.
	Duration param.Field[int64] `json:"duration,required"`
	// The unit of billing period duration.
	DurationUnit param.Field[PlanExternalPlanIDVersionNewParamsReplacePricesPriceNewPlanBpsPriceInvoicingCycleConfigurationDurationUnit] `json:"duration_unit,required"`
}

func (r PlanExternalPlanIDVersionNewParamsReplacePricesPriceNewPlanBpsPriceInvoicingCycleConfiguration) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// The unit of billing period duration.
type PlanExternalPlanIDVersionNewParamsReplacePricesPriceNewPlanBpsPriceInvoicingCycleConfigurationDurationUnit string

const (
	PlanExternalPlanIDVersionNewParamsReplacePricesPriceNewPlanBpsPriceInvoicingCycleConfigurationDurationUnitDay   PlanExternalPlanIDVersionNewParamsReplacePricesPriceNewPlanBpsPriceInvoicingCycleConfigurationDurationUnit = "day"
	PlanExternalPlanIDVersionNewParamsReplacePricesPriceNewPlanBpsPriceInvoicingCycleConfigurationDurationUnitMonth PlanExternalPlanIDVersionNewParamsReplacePricesPriceNewPlanBpsPriceInvoicingCycleConfigurationDurationUnit = "month"
)

func (r PlanExternalPlanIDVersionNewParamsReplacePricesPriceNewPlanBpsPriceInvoicingCycleConfigurationDurationUnit) IsKnown() bool {
	switch r {
	case PlanExternalPlanIDVersionNewParamsReplacePricesPriceNewPlanBpsPriceInvoicingCycleConfigurationDurationUnitDay, PlanExternalPlanIDVersionNewParamsReplacePricesPriceNewPlanBpsPriceInvoicingCycleConfigurationDurationUnitMonth:
		return true
	}
	return false
}

type PlanExternalPlanIDVersionNewParamsReplacePricesPriceNewPlanBulkBpsPrice struct {
	BulkBpsConfig param.Field[PlanExternalPlanIDVersionNewParamsReplacePricesPriceNewPlanBulkBpsPriceBulkBpsConfig] `json:"bulk_bps_config,required"`
	// The cadence to bill for this price on.
	Cadence param.Field[PlanExternalPlanIDVersionNewParamsReplacePricesPriceNewPlanBulkBpsPriceCadence] `json:"cadence,required"`
	// The id of the item the price will be associated with.
	ItemID    param.Field[string]                                                                           `json:"item_id,required"`
	ModelType param.Field[PlanExternalPlanIDVersionNewParamsReplacePricesPriceNewPlanBulkBpsPriceModelType] `json:"model_type,required"`
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
	BillingCycleConfiguration param.Field[PlanExternalPlanIDVersionNewParamsReplacePricesPriceNewPlanBulkBpsPriceBillingCycleConfiguration] `json:"billing_cycle_configuration"`
	// The per unit conversion rate of the price currency to the invoicing currency.
	ConversionRate param.Field[float64] `json:"conversion_rate"`
	// An ISO 4217 currency string, or custom pricing unit identifier, in which this
	// price is billed.
	Currency param.Field[string] `json:"currency"`
	// For dimensional price: specifies a price group and dimension values
	DimensionalPriceConfiguration param.Field[PlanExternalPlanIDVersionNewParamsReplacePricesPriceNewPlanBulkBpsPriceDimensionalPriceConfiguration] `json:"dimensional_price_configuration"`
	// An alias for the price.
	ExternalPriceID param.Field[string] `json:"external_price_id"`
	// If the Price represents a fixed cost, this represents the quantity of units
	// applied.
	FixedPriceQuantity param.Field[float64] `json:"fixed_price_quantity"`
	// The property used to group this price on an invoice
	InvoiceGroupingKey param.Field[string] `json:"invoice_grouping_key"`
	// Within each billing cycle, specifies the cadence at which invoices are produced.
	// If unspecified, a single invoice is produced per billing cycle.
	InvoicingCycleConfiguration param.Field[PlanExternalPlanIDVersionNewParamsReplacePricesPriceNewPlanBulkBpsPriceInvoicingCycleConfiguration] `json:"invoicing_cycle_configuration"`
	// User-specified key/value pairs for the resource. Individual keys can be removed
	// by setting the value to `null`, and the entire metadata mapping can be cleared
	// by setting `metadata` to `null`.
	Metadata param.Field[map[string]string] `json:"metadata"`
}

func (r PlanExternalPlanIDVersionNewParamsReplacePricesPriceNewPlanBulkBpsPrice) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r PlanExternalPlanIDVersionNewParamsReplacePricesPriceNewPlanBulkBpsPrice) implementsPlanExternalPlanIDVersionNewParamsReplacePricesPriceUnion() {
}

type PlanExternalPlanIDVersionNewParamsReplacePricesPriceNewPlanBulkBpsPriceBulkBpsConfig struct {
	// Tiers for a bulk BPS pricing model where all usage is aggregated to a single
	// tier based on total volume
	Tiers param.Field[[]PlanExternalPlanIDVersionNewParamsReplacePricesPriceNewPlanBulkBpsPriceBulkBpsConfigTier] `json:"tiers,required"`
}

func (r PlanExternalPlanIDVersionNewParamsReplacePricesPriceNewPlanBulkBpsPriceBulkBpsConfig) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type PlanExternalPlanIDVersionNewParamsReplacePricesPriceNewPlanBulkBpsPriceBulkBpsConfigTier struct {
	// Basis points to rate on
	Bps param.Field[float64] `json:"bps,required"`
	// Upper bound for tier
	MaximumAmount param.Field[string] `json:"maximum_amount"`
	// The maximum amount to charge for any one event
	PerUnitMaximum param.Field[string] `json:"per_unit_maximum"`
}

func (r PlanExternalPlanIDVersionNewParamsReplacePricesPriceNewPlanBulkBpsPriceBulkBpsConfigTier) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// The cadence to bill for this price on.
type PlanExternalPlanIDVersionNewParamsReplacePricesPriceNewPlanBulkBpsPriceCadence string

const (
	PlanExternalPlanIDVersionNewParamsReplacePricesPriceNewPlanBulkBpsPriceCadenceAnnual     PlanExternalPlanIDVersionNewParamsReplacePricesPriceNewPlanBulkBpsPriceCadence = "annual"
	PlanExternalPlanIDVersionNewParamsReplacePricesPriceNewPlanBulkBpsPriceCadenceSemiAnnual PlanExternalPlanIDVersionNewParamsReplacePricesPriceNewPlanBulkBpsPriceCadence = "semi_annual"
	PlanExternalPlanIDVersionNewParamsReplacePricesPriceNewPlanBulkBpsPriceCadenceMonthly    PlanExternalPlanIDVersionNewParamsReplacePricesPriceNewPlanBulkBpsPriceCadence = "monthly"
	PlanExternalPlanIDVersionNewParamsReplacePricesPriceNewPlanBulkBpsPriceCadenceQuarterly  PlanExternalPlanIDVersionNewParamsReplacePricesPriceNewPlanBulkBpsPriceCadence = "quarterly"
	PlanExternalPlanIDVersionNewParamsReplacePricesPriceNewPlanBulkBpsPriceCadenceOneTime    PlanExternalPlanIDVersionNewParamsReplacePricesPriceNewPlanBulkBpsPriceCadence = "one_time"
	PlanExternalPlanIDVersionNewParamsReplacePricesPriceNewPlanBulkBpsPriceCadenceCustom     PlanExternalPlanIDVersionNewParamsReplacePricesPriceNewPlanBulkBpsPriceCadence = "custom"
)

func (r PlanExternalPlanIDVersionNewParamsReplacePricesPriceNewPlanBulkBpsPriceCadence) IsKnown() bool {
	switch r {
	case PlanExternalPlanIDVersionNewParamsReplacePricesPriceNewPlanBulkBpsPriceCadenceAnnual, PlanExternalPlanIDVersionNewParamsReplacePricesPriceNewPlanBulkBpsPriceCadenceSemiAnnual, PlanExternalPlanIDVersionNewParamsReplacePricesPriceNewPlanBulkBpsPriceCadenceMonthly, PlanExternalPlanIDVersionNewParamsReplacePricesPriceNewPlanBulkBpsPriceCadenceQuarterly, PlanExternalPlanIDVersionNewParamsReplacePricesPriceNewPlanBulkBpsPriceCadenceOneTime, PlanExternalPlanIDVersionNewParamsReplacePricesPriceNewPlanBulkBpsPriceCadenceCustom:
		return true
	}
	return false
}

type PlanExternalPlanIDVersionNewParamsReplacePricesPriceNewPlanBulkBpsPriceModelType string

const (
	PlanExternalPlanIDVersionNewParamsReplacePricesPriceNewPlanBulkBpsPriceModelTypeBulkBps PlanExternalPlanIDVersionNewParamsReplacePricesPriceNewPlanBulkBpsPriceModelType = "bulk_bps"
)

func (r PlanExternalPlanIDVersionNewParamsReplacePricesPriceNewPlanBulkBpsPriceModelType) IsKnown() bool {
	switch r {
	case PlanExternalPlanIDVersionNewParamsReplacePricesPriceNewPlanBulkBpsPriceModelTypeBulkBps:
		return true
	}
	return false
}

// For custom cadence: specifies the duration of the billing period in days or
// months.
type PlanExternalPlanIDVersionNewParamsReplacePricesPriceNewPlanBulkBpsPriceBillingCycleConfiguration struct {
	// The duration of the billing period.
	Duration param.Field[int64] `json:"duration,required"`
	// The unit of billing period duration.
	DurationUnit param.Field[PlanExternalPlanIDVersionNewParamsReplacePricesPriceNewPlanBulkBpsPriceBillingCycleConfigurationDurationUnit] `json:"duration_unit,required"`
}

func (r PlanExternalPlanIDVersionNewParamsReplacePricesPriceNewPlanBulkBpsPriceBillingCycleConfiguration) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// The unit of billing period duration.
type PlanExternalPlanIDVersionNewParamsReplacePricesPriceNewPlanBulkBpsPriceBillingCycleConfigurationDurationUnit string

const (
	PlanExternalPlanIDVersionNewParamsReplacePricesPriceNewPlanBulkBpsPriceBillingCycleConfigurationDurationUnitDay   PlanExternalPlanIDVersionNewParamsReplacePricesPriceNewPlanBulkBpsPriceBillingCycleConfigurationDurationUnit = "day"
	PlanExternalPlanIDVersionNewParamsReplacePricesPriceNewPlanBulkBpsPriceBillingCycleConfigurationDurationUnitMonth PlanExternalPlanIDVersionNewParamsReplacePricesPriceNewPlanBulkBpsPriceBillingCycleConfigurationDurationUnit = "month"
)

func (r PlanExternalPlanIDVersionNewParamsReplacePricesPriceNewPlanBulkBpsPriceBillingCycleConfigurationDurationUnit) IsKnown() bool {
	switch r {
	case PlanExternalPlanIDVersionNewParamsReplacePricesPriceNewPlanBulkBpsPriceBillingCycleConfigurationDurationUnitDay, PlanExternalPlanIDVersionNewParamsReplacePricesPriceNewPlanBulkBpsPriceBillingCycleConfigurationDurationUnitMonth:
		return true
	}
	return false
}

// For dimensional price: specifies a price group and dimension values
type PlanExternalPlanIDVersionNewParamsReplacePricesPriceNewPlanBulkBpsPriceDimensionalPriceConfiguration struct {
	// The list of dimension values matching (in order) the dimensions of the price
	// group
	DimensionValues param.Field[[]string] `json:"dimension_values,required"`
	// The id of the dimensional price group to include this price in
	DimensionalPriceGroupID param.Field[string] `json:"dimensional_price_group_id"`
	// The external id of the dimensional price group to include this price in
	ExternalDimensionalPriceGroupID param.Field[string] `json:"external_dimensional_price_group_id"`
}

func (r PlanExternalPlanIDVersionNewParamsReplacePricesPriceNewPlanBulkBpsPriceDimensionalPriceConfiguration) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Within each billing cycle, specifies the cadence at which invoices are produced.
// If unspecified, a single invoice is produced per billing cycle.
type PlanExternalPlanIDVersionNewParamsReplacePricesPriceNewPlanBulkBpsPriceInvoicingCycleConfiguration struct {
	// The duration of the billing period.
	Duration param.Field[int64] `json:"duration,required"`
	// The unit of billing period duration.
	DurationUnit param.Field[PlanExternalPlanIDVersionNewParamsReplacePricesPriceNewPlanBulkBpsPriceInvoicingCycleConfigurationDurationUnit] `json:"duration_unit,required"`
}

func (r PlanExternalPlanIDVersionNewParamsReplacePricesPriceNewPlanBulkBpsPriceInvoicingCycleConfiguration) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// The unit of billing period duration.
type PlanExternalPlanIDVersionNewParamsReplacePricesPriceNewPlanBulkBpsPriceInvoicingCycleConfigurationDurationUnit string

const (
	PlanExternalPlanIDVersionNewParamsReplacePricesPriceNewPlanBulkBpsPriceInvoicingCycleConfigurationDurationUnitDay   PlanExternalPlanIDVersionNewParamsReplacePricesPriceNewPlanBulkBpsPriceInvoicingCycleConfigurationDurationUnit = "day"
	PlanExternalPlanIDVersionNewParamsReplacePricesPriceNewPlanBulkBpsPriceInvoicingCycleConfigurationDurationUnitMonth PlanExternalPlanIDVersionNewParamsReplacePricesPriceNewPlanBulkBpsPriceInvoicingCycleConfigurationDurationUnit = "month"
)

func (r PlanExternalPlanIDVersionNewParamsReplacePricesPriceNewPlanBulkBpsPriceInvoicingCycleConfigurationDurationUnit) IsKnown() bool {
	switch r {
	case PlanExternalPlanIDVersionNewParamsReplacePricesPriceNewPlanBulkBpsPriceInvoicingCycleConfigurationDurationUnitDay, PlanExternalPlanIDVersionNewParamsReplacePricesPriceNewPlanBulkBpsPriceInvoicingCycleConfigurationDurationUnitMonth:
		return true
	}
	return false
}

type PlanExternalPlanIDVersionNewParamsReplacePricesPriceNewPlanBulkPrice struct {
	BulkConfig param.Field[PlanExternalPlanIDVersionNewParamsReplacePricesPriceNewPlanBulkPriceBulkConfig] `json:"bulk_config,required"`
	// The cadence to bill for this price on.
	Cadence param.Field[PlanExternalPlanIDVersionNewParamsReplacePricesPriceNewPlanBulkPriceCadence] `json:"cadence,required"`
	// The id of the item the price will be associated with.
	ItemID    param.Field[string]                                                                        `json:"item_id,required"`
	ModelType param.Field[PlanExternalPlanIDVersionNewParamsReplacePricesPriceNewPlanBulkPriceModelType] `json:"model_type,required"`
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
	BillingCycleConfiguration param.Field[PlanExternalPlanIDVersionNewParamsReplacePricesPriceNewPlanBulkPriceBillingCycleConfiguration] `json:"billing_cycle_configuration"`
	// The per unit conversion rate of the price currency to the invoicing currency.
	ConversionRate param.Field[float64] `json:"conversion_rate"`
	// An ISO 4217 currency string, or custom pricing unit identifier, in which this
	// price is billed.
	Currency param.Field[string] `json:"currency"`
	// For dimensional price: specifies a price group and dimension values
	DimensionalPriceConfiguration param.Field[PlanExternalPlanIDVersionNewParamsReplacePricesPriceNewPlanBulkPriceDimensionalPriceConfiguration] `json:"dimensional_price_configuration"`
	// An alias for the price.
	ExternalPriceID param.Field[string] `json:"external_price_id"`
	// If the Price represents a fixed cost, this represents the quantity of units
	// applied.
	FixedPriceQuantity param.Field[float64] `json:"fixed_price_quantity"`
	// The property used to group this price on an invoice
	InvoiceGroupingKey param.Field[string] `json:"invoice_grouping_key"`
	// Within each billing cycle, specifies the cadence at which invoices are produced.
	// If unspecified, a single invoice is produced per billing cycle.
	InvoicingCycleConfiguration param.Field[PlanExternalPlanIDVersionNewParamsReplacePricesPriceNewPlanBulkPriceInvoicingCycleConfiguration] `json:"invoicing_cycle_configuration"`
	// User-specified key/value pairs for the resource. Individual keys can be removed
	// by setting the value to `null`, and the entire metadata mapping can be cleared
	// by setting `metadata` to `null`.
	Metadata param.Field[map[string]string] `json:"metadata"`
}

func (r PlanExternalPlanIDVersionNewParamsReplacePricesPriceNewPlanBulkPrice) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r PlanExternalPlanIDVersionNewParamsReplacePricesPriceNewPlanBulkPrice) implementsPlanExternalPlanIDVersionNewParamsReplacePricesPriceUnion() {
}

type PlanExternalPlanIDVersionNewParamsReplacePricesPriceNewPlanBulkPriceBulkConfig struct {
	// Bulk tiers for rating based on total usage volume
	Tiers param.Field[[]PlanExternalPlanIDVersionNewParamsReplacePricesPriceNewPlanBulkPriceBulkConfigTier] `json:"tiers,required"`
}

func (r PlanExternalPlanIDVersionNewParamsReplacePricesPriceNewPlanBulkPriceBulkConfig) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type PlanExternalPlanIDVersionNewParamsReplacePricesPriceNewPlanBulkPriceBulkConfigTier struct {
	// Amount per unit
	UnitAmount param.Field[string] `json:"unit_amount,required"`
	// Upper bound for this tier
	MaximumUnits param.Field[float64] `json:"maximum_units"`
}

func (r PlanExternalPlanIDVersionNewParamsReplacePricesPriceNewPlanBulkPriceBulkConfigTier) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// The cadence to bill for this price on.
type PlanExternalPlanIDVersionNewParamsReplacePricesPriceNewPlanBulkPriceCadence string

const (
	PlanExternalPlanIDVersionNewParamsReplacePricesPriceNewPlanBulkPriceCadenceAnnual     PlanExternalPlanIDVersionNewParamsReplacePricesPriceNewPlanBulkPriceCadence = "annual"
	PlanExternalPlanIDVersionNewParamsReplacePricesPriceNewPlanBulkPriceCadenceSemiAnnual PlanExternalPlanIDVersionNewParamsReplacePricesPriceNewPlanBulkPriceCadence = "semi_annual"
	PlanExternalPlanIDVersionNewParamsReplacePricesPriceNewPlanBulkPriceCadenceMonthly    PlanExternalPlanIDVersionNewParamsReplacePricesPriceNewPlanBulkPriceCadence = "monthly"
	PlanExternalPlanIDVersionNewParamsReplacePricesPriceNewPlanBulkPriceCadenceQuarterly  PlanExternalPlanIDVersionNewParamsReplacePricesPriceNewPlanBulkPriceCadence = "quarterly"
	PlanExternalPlanIDVersionNewParamsReplacePricesPriceNewPlanBulkPriceCadenceOneTime    PlanExternalPlanIDVersionNewParamsReplacePricesPriceNewPlanBulkPriceCadence = "one_time"
	PlanExternalPlanIDVersionNewParamsReplacePricesPriceNewPlanBulkPriceCadenceCustom     PlanExternalPlanIDVersionNewParamsReplacePricesPriceNewPlanBulkPriceCadence = "custom"
)

func (r PlanExternalPlanIDVersionNewParamsReplacePricesPriceNewPlanBulkPriceCadence) IsKnown() bool {
	switch r {
	case PlanExternalPlanIDVersionNewParamsReplacePricesPriceNewPlanBulkPriceCadenceAnnual, PlanExternalPlanIDVersionNewParamsReplacePricesPriceNewPlanBulkPriceCadenceSemiAnnual, PlanExternalPlanIDVersionNewParamsReplacePricesPriceNewPlanBulkPriceCadenceMonthly, PlanExternalPlanIDVersionNewParamsReplacePricesPriceNewPlanBulkPriceCadenceQuarterly, PlanExternalPlanIDVersionNewParamsReplacePricesPriceNewPlanBulkPriceCadenceOneTime, PlanExternalPlanIDVersionNewParamsReplacePricesPriceNewPlanBulkPriceCadenceCustom:
		return true
	}
	return false
}

type PlanExternalPlanIDVersionNewParamsReplacePricesPriceNewPlanBulkPriceModelType string

const (
	PlanExternalPlanIDVersionNewParamsReplacePricesPriceNewPlanBulkPriceModelTypeBulk PlanExternalPlanIDVersionNewParamsReplacePricesPriceNewPlanBulkPriceModelType = "bulk"
)

func (r PlanExternalPlanIDVersionNewParamsReplacePricesPriceNewPlanBulkPriceModelType) IsKnown() bool {
	switch r {
	case PlanExternalPlanIDVersionNewParamsReplacePricesPriceNewPlanBulkPriceModelTypeBulk:
		return true
	}
	return false
}

// For custom cadence: specifies the duration of the billing period in days or
// months.
type PlanExternalPlanIDVersionNewParamsReplacePricesPriceNewPlanBulkPriceBillingCycleConfiguration struct {
	// The duration of the billing period.
	Duration param.Field[int64] `json:"duration,required"`
	// The unit of billing period duration.
	DurationUnit param.Field[PlanExternalPlanIDVersionNewParamsReplacePricesPriceNewPlanBulkPriceBillingCycleConfigurationDurationUnit] `json:"duration_unit,required"`
}

func (r PlanExternalPlanIDVersionNewParamsReplacePricesPriceNewPlanBulkPriceBillingCycleConfiguration) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// The unit of billing period duration.
type PlanExternalPlanIDVersionNewParamsReplacePricesPriceNewPlanBulkPriceBillingCycleConfigurationDurationUnit string

const (
	PlanExternalPlanIDVersionNewParamsReplacePricesPriceNewPlanBulkPriceBillingCycleConfigurationDurationUnitDay   PlanExternalPlanIDVersionNewParamsReplacePricesPriceNewPlanBulkPriceBillingCycleConfigurationDurationUnit = "day"
	PlanExternalPlanIDVersionNewParamsReplacePricesPriceNewPlanBulkPriceBillingCycleConfigurationDurationUnitMonth PlanExternalPlanIDVersionNewParamsReplacePricesPriceNewPlanBulkPriceBillingCycleConfigurationDurationUnit = "month"
)

func (r PlanExternalPlanIDVersionNewParamsReplacePricesPriceNewPlanBulkPriceBillingCycleConfigurationDurationUnit) IsKnown() bool {
	switch r {
	case PlanExternalPlanIDVersionNewParamsReplacePricesPriceNewPlanBulkPriceBillingCycleConfigurationDurationUnitDay, PlanExternalPlanIDVersionNewParamsReplacePricesPriceNewPlanBulkPriceBillingCycleConfigurationDurationUnitMonth:
		return true
	}
	return false
}

// For dimensional price: specifies a price group and dimension values
type PlanExternalPlanIDVersionNewParamsReplacePricesPriceNewPlanBulkPriceDimensionalPriceConfiguration struct {
	// The list of dimension values matching (in order) the dimensions of the price
	// group
	DimensionValues param.Field[[]string] `json:"dimension_values,required"`
	// The id of the dimensional price group to include this price in
	DimensionalPriceGroupID param.Field[string] `json:"dimensional_price_group_id"`
	// The external id of the dimensional price group to include this price in
	ExternalDimensionalPriceGroupID param.Field[string] `json:"external_dimensional_price_group_id"`
}

func (r PlanExternalPlanIDVersionNewParamsReplacePricesPriceNewPlanBulkPriceDimensionalPriceConfiguration) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Within each billing cycle, specifies the cadence at which invoices are produced.
// If unspecified, a single invoice is produced per billing cycle.
type PlanExternalPlanIDVersionNewParamsReplacePricesPriceNewPlanBulkPriceInvoicingCycleConfiguration struct {
	// The duration of the billing period.
	Duration param.Field[int64] `json:"duration,required"`
	// The unit of billing period duration.
	DurationUnit param.Field[PlanExternalPlanIDVersionNewParamsReplacePricesPriceNewPlanBulkPriceInvoicingCycleConfigurationDurationUnit] `json:"duration_unit,required"`
}

func (r PlanExternalPlanIDVersionNewParamsReplacePricesPriceNewPlanBulkPriceInvoicingCycleConfiguration) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// The unit of billing period duration.
type PlanExternalPlanIDVersionNewParamsReplacePricesPriceNewPlanBulkPriceInvoicingCycleConfigurationDurationUnit string

const (
	PlanExternalPlanIDVersionNewParamsReplacePricesPriceNewPlanBulkPriceInvoicingCycleConfigurationDurationUnitDay   PlanExternalPlanIDVersionNewParamsReplacePricesPriceNewPlanBulkPriceInvoicingCycleConfigurationDurationUnit = "day"
	PlanExternalPlanIDVersionNewParamsReplacePricesPriceNewPlanBulkPriceInvoicingCycleConfigurationDurationUnitMonth PlanExternalPlanIDVersionNewParamsReplacePricesPriceNewPlanBulkPriceInvoicingCycleConfigurationDurationUnit = "month"
)

func (r PlanExternalPlanIDVersionNewParamsReplacePricesPriceNewPlanBulkPriceInvoicingCycleConfigurationDurationUnit) IsKnown() bool {
	switch r {
	case PlanExternalPlanIDVersionNewParamsReplacePricesPriceNewPlanBulkPriceInvoicingCycleConfigurationDurationUnitDay, PlanExternalPlanIDVersionNewParamsReplacePricesPriceNewPlanBulkPriceInvoicingCycleConfigurationDurationUnitMonth:
		return true
	}
	return false
}

type PlanExternalPlanIDVersionNewParamsReplacePricesPriceNewPlanThresholdTotalAmountPrice struct {
	// The cadence to bill for this price on.
	Cadence param.Field[PlanExternalPlanIDVersionNewParamsReplacePricesPriceNewPlanThresholdTotalAmountPriceCadence] `json:"cadence,required"`
	// The id of the item the price will be associated with.
	ItemID    param.Field[string]                                                                                        `json:"item_id,required"`
	ModelType param.Field[PlanExternalPlanIDVersionNewParamsReplacePricesPriceNewPlanThresholdTotalAmountPriceModelType] `json:"model_type,required"`
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
	BillingCycleConfiguration param.Field[PlanExternalPlanIDVersionNewParamsReplacePricesPriceNewPlanThresholdTotalAmountPriceBillingCycleConfiguration] `json:"billing_cycle_configuration"`
	// The per unit conversion rate of the price currency to the invoicing currency.
	ConversionRate param.Field[float64] `json:"conversion_rate"`
	// An ISO 4217 currency string, or custom pricing unit identifier, in which this
	// price is billed.
	Currency param.Field[string] `json:"currency"`
	// For dimensional price: specifies a price group and dimension values
	DimensionalPriceConfiguration param.Field[PlanExternalPlanIDVersionNewParamsReplacePricesPriceNewPlanThresholdTotalAmountPriceDimensionalPriceConfiguration] `json:"dimensional_price_configuration"`
	// An alias for the price.
	ExternalPriceID param.Field[string] `json:"external_price_id"`
	// If the Price represents a fixed cost, this represents the quantity of units
	// applied.
	FixedPriceQuantity param.Field[float64] `json:"fixed_price_quantity"`
	// The property used to group this price on an invoice
	InvoiceGroupingKey param.Field[string] `json:"invoice_grouping_key"`
	// Within each billing cycle, specifies the cadence at which invoices are produced.
	// If unspecified, a single invoice is produced per billing cycle.
	InvoicingCycleConfiguration param.Field[PlanExternalPlanIDVersionNewParamsReplacePricesPriceNewPlanThresholdTotalAmountPriceInvoicingCycleConfiguration] `json:"invoicing_cycle_configuration"`
	// User-specified key/value pairs for the resource. Individual keys can be removed
	// by setting the value to `null`, and the entire metadata mapping can be cleared
	// by setting `metadata` to `null`.
	Metadata param.Field[map[string]string] `json:"metadata"`
}

func (r PlanExternalPlanIDVersionNewParamsReplacePricesPriceNewPlanThresholdTotalAmountPrice) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r PlanExternalPlanIDVersionNewParamsReplacePricesPriceNewPlanThresholdTotalAmountPrice) implementsPlanExternalPlanIDVersionNewParamsReplacePricesPriceUnion() {
}

// The cadence to bill for this price on.
type PlanExternalPlanIDVersionNewParamsReplacePricesPriceNewPlanThresholdTotalAmountPriceCadence string

const (
	PlanExternalPlanIDVersionNewParamsReplacePricesPriceNewPlanThresholdTotalAmountPriceCadenceAnnual     PlanExternalPlanIDVersionNewParamsReplacePricesPriceNewPlanThresholdTotalAmountPriceCadence = "annual"
	PlanExternalPlanIDVersionNewParamsReplacePricesPriceNewPlanThresholdTotalAmountPriceCadenceSemiAnnual PlanExternalPlanIDVersionNewParamsReplacePricesPriceNewPlanThresholdTotalAmountPriceCadence = "semi_annual"
	PlanExternalPlanIDVersionNewParamsReplacePricesPriceNewPlanThresholdTotalAmountPriceCadenceMonthly    PlanExternalPlanIDVersionNewParamsReplacePricesPriceNewPlanThresholdTotalAmountPriceCadence = "monthly"
	PlanExternalPlanIDVersionNewParamsReplacePricesPriceNewPlanThresholdTotalAmountPriceCadenceQuarterly  PlanExternalPlanIDVersionNewParamsReplacePricesPriceNewPlanThresholdTotalAmountPriceCadence = "quarterly"
	PlanExternalPlanIDVersionNewParamsReplacePricesPriceNewPlanThresholdTotalAmountPriceCadenceOneTime    PlanExternalPlanIDVersionNewParamsReplacePricesPriceNewPlanThresholdTotalAmountPriceCadence = "one_time"
	PlanExternalPlanIDVersionNewParamsReplacePricesPriceNewPlanThresholdTotalAmountPriceCadenceCustom     PlanExternalPlanIDVersionNewParamsReplacePricesPriceNewPlanThresholdTotalAmountPriceCadence = "custom"
)

func (r PlanExternalPlanIDVersionNewParamsReplacePricesPriceNewPlanThresholdTotalAmountPriceCadence) IsKnown() bool {
	switch r {
	case PlanExternalPlanIDVersionNewParamsReplacePricesPriceNewPlanThresholdTotalAmountPriceCadenceAnnual, PlanExternalPlanIDVersionNewParamsReplacePricesPriceNewPlanThresholdTotalAmountPriceCadenceSemiAnnual, PlanExternalPlanIDVersionNewParamsReplacePricesPriceNewPlanThresholdTotalAmountPriceCadenceMonthly, PlanExternalPlanIDVersionNewParamsReplacePricesPriceNewPlanThresholdTotalAmountPriceCadenceQuarterly, PlanExternalPlanIDVersionNewParamsReplacePricesPriceNewPlanThresholdTotalAmountPriceCadenceOneTime, PlanExternalPlanIDVersionNewParamsReplacePricesPriceNewPlanThresholdTotalAmountPriceCadenceCustom:
		return true
	}
	return false
}

type PlanExternalPlanIDVersionNewParamsReplacePricesPriceNewPlanThresholdTotalAmountPriceModelType string

const (
	PlanExternalPlanIDVersionNewParamsReplacePricesPriceNewPlanThresholdTotalAmountPriceModelTypeThresholdTotalAmount PlanExternalPlanIDVersionNewParamsReplacePricesPriceNewPlanThresholdTotalAmountPriceModelType = "threshold_total_amount"
)

func (r PlanExternalPlanIDVersionNewParamsReplacePricesPriceNewPlanThresholdTotalAmountPriceModelType) IsKnown() bool {
	switch r {
	case PlanExternalPlanIDVersionNewParamsReplacePricesPriceNewPlanThresholdTotalAmountPriceModelTypeThresholdTotalAmount:
		return true
	}
	return false
}

// For custom cadence: specifies the duration of the billing period in days or
// months.
type PlanExternalPlanIDVersionNewParamsReplacePricesPriceNewPlanThresholdTotalAmountPriceBillingCycleConfiguration struct {
	// The duration of the billing period.
	Duration param.Field[int64] `json:"duration,required"`
	// The unit of billing period duration.
	DurationUnit param.Field[PlanExternalPlanIDVersionNewParamsReplacePricesPriceNewPlanThresholdTotalAmountPriceBillingCycleConfigurationDurationUnit] `json:"duration_unit,required"`
}

func (r PlanExternalPlanIDVersionNewParamsReplacePricesPriceNewPlanThresholdTotalAmountPriceBillingCycleConfiguration) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// The unit of billing period duration.
type PlanExternalPlanIDVersionNewParamsReplacePricesPriceNewPlanThresholdTotalAmountPriceBillingCycleConfigurationDurationUnit string

const (
	PlanExternalPlanIDVersionNewParamsReplacePricesPriceNewPlanThresholdTotalAmountPriceBillingCycleConfigurationDurationUnitDay   PlanExternalPlanIDVersionNewParamsReplacePricesPriceNewPlanThresholdTotalAmountPriceBillingCycleConfigurationDurationUnit = "day"
	PlanExternalPlanIDVersionNewParamsReplacePricesPriceNewPlanThresholdTotalAmountPriceBillingCycleConfigurationDurationUnitMonth PlanExternalPlanIDVersionNewParamsReplacePricesPriceNewPlanThresholdTotalAmountPriceBillingCycleConfigurationDurationUnit = "month"
)

func (r PlanExternalPlanIDVersionNewParamsReplacePricesPriceNewPlanThresholdTotalAmountPriceBillingCycleConfigurationDurationUnit) IsKnown() bool {
	switch r {
	case PlanExternalPlanIDVersionNewParamsReplacePricesPriceNewPlanThresholdTotalAmountPriceBillingCycleConfigurationDurationUnitDay, PlanExternalPlanIDVersionNewParamsReplacePricesPriceNewPlanThresholdTotalAmountPriceBillingCycleConfigurationDurationUnitMonth:
		return true
	}
	return false
}

// For dimensional price: specifies a price group and dimension values
type PlanExternalPlanIDVersionNewParamsReplacePricesPriceNewPlanThresholdTotalAmountPriceDimensionalPriceConfiguration struct {
	// The list of dimension values matching (in order) the dimensions of the price
	// group
	DimensionValues param.Field[[]string] `json:"dimension_values,required"`
	// The id of the dimensional price group to include this price in
	DimensionalPriceGroupID param.Field[string] `json:"dimensional_price_group_id"`
	// The external id of the dimensional price group to include this price in
	ExternalDimensionalPriceGroupID param.Field[string] `json:"external_dimensional_price_group_id"`
}

func (r PlanExternalPlanIDVersionNewParamsReplacePricesPriceNewPlanThresholdTotalAmountPriceDimensionalPriceConfiguration) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Within each billing cycle, specifies the cadence at which invoices are produced.
// If unspecified, a single invoice is produced per billing cycle.
type PlanExternalPlanIDVersionNewParamsReplacePricesPriceNewPlanThresholdTotalAmountPriceInvoicingCycleConfiguration struct {
	// The duration of the billing period.
	Duration param.Field[int64] `json:"duration,required"`
	// The unit of billing period duration.
	DurationUnit param.Field[PlanExternalPlanIDVersionNewParamsReplacePricesPriceNewPlanThresholdTotalAmountPriceInvoicingCycleConfigurationDurationUnit] `json:"duration_unit,required"`
}

func (r PlanExternalPlanIDVersionNewParamsReplacePricesPriceNewPlanThresholdTotalAmountPriceInvoicingCycleConfiguration) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// The unit of billing period duration.
type PlanExternalPlanIDVersionNewParamsReplacePricesPriceNewPlanThresholdTotalAmountPriceInvoicingCycleConfigurationDurationUnit string

const (
	PlanExternalPlanIDVersionNewParamsReplacePricesPriceNewPlanThresholdTotalAmountPriceInvoicingCycleConfigurationDurationUnitDay   PlanExternalPlanIDVersionNewParamsReplacePricesPriceNewPlanThresholdTotalAmountPriceInvoicingCycleConfigurationDurationUnit = "day"
	PlanExternalPlanIDVersionNewParamsReplacePricesPriceNewPlanThresholdTotalAmountPriceInvoicingCycleConfigurationDurationUnitMonth PlanExternalPlanIDVersionNewParamsReplacePricesPriceNewPlanThresholdTotalAmountPriceInvoicingCycleConfigurationDurationUnit = "month"
)

func (r PlanExternalPlanIDVersionNewParamsReplacePricesPriceNewPlanThresholdTotalAmountPriceInvoicingCycleConfigurationDurationUnit) IsKnown() bool {
	switch r {
	case PlanExternalPlanIDVersionNewParamsReplacePricesPriceNewPlanThresholdTotalAmountPriceInvoicingCycleConfigurationDurationUnitDay, PlanExternalPlanIDVersionNewParamsReplacePricesPriceNewPlanThresholdTotalAmountPriceInvoicingCycleConfigurationDurationUnitMonth:
		return true
	}
	return false
}

type PlanExternalPlanIDVersionNewParamsReplacePricesPriceNewPlanTieredPackagePrice struct {
	// The cadence to bill for this price on.
	Cadence param.Field[PlanExternalPlanIDVersionNewParamsReplacePricesPriceNewPlanTieredPackagePriceCadence] `json:"cadence,required"`
	// The id of the item the price will be associated with.
	ItemID    param.Field[string]                                                                                 `json:"item_id,required"`
	ModelType param.Field[PlanExternalPlanIDVersionNewParamsReplacePricesPriceNewPlanTieredPackagePriceModelType] `json:"model_type,required"`
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
	BillingCycleConfiguration param.Field[PlanExternalPlanIDVersionNewParamsReplacePricesPriceNewPlanTieredPackagePriceBillingCycleConfiguration] `json:"billing_cycle_configuration"`
	// The per unit conversion rate of the price currency to the invoicing currency.
	ConversionRate param.Field[float64] `json:"conversion_rate"`
	// An ISO 4217 currency string, or custom pricing unit identifier, in which this
	// price is billed.
	Currency param.Field[string] `json:"currency"`
	// For dimensional price: specifies a price group and dimension values
	DimensionalPriceConfiguration param.Field[PlanExternalPlanIDVersionNewParamsReplacePricesPriceNewPlanTieredPackagePriceDimensionalPriceConfiguration] `json:"dimensional_price_configuration"`
	// An alias for the price.
	ExternalPriceID param.Field[string] `json:"external_price_id"`
	// If the Price represents a fixed cost, this represents the quantity of units
	// applied.
	FixedPriceQuantity param.Field[float64] `json:"fixed_price_quantity"`
	// The property used to group this price on an invoice
	InvoiceGroupingKey param.Field[string] `json:"invoice_grouping_key"`
	// Within each billing cycle, specifies the cadence at which invoices are produced.
	// If unspecified, a single invoice is produced per billing cycle.
	InvoicingCycleConfiguration param.Field[PlanExternalPlanIDVersionNewParamsReplacePricesPriceNewPlanTieredPackagePriceInvoicingCycleConfiguration] `json:"invoicing_cycle_configuration"`
	// User-specified key/value pairs for the resource. Individual keys can be removed
	// by setting the value to `null`, and the entire metadata mapping can be cleared
	// by setting `metadata` to `null`.
	Metadata param.Field[map[string]string] `json:"metadata"`
}

func (r PlanExternalPlanIDVersionNewParamsReplacePricesPriceNewPlanTieredPackagePrice) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r PlanExternalPlanIDVersionNewParamsReplacePricesPriceNewPlanTieredPackagePrice) implementsPlanExternalPlanIDVersionNewParamsReplacePricesPriceUnion() {
}

// The cadence to bill for this price on.
type PlanExternalPlanIDVersionNewParamsReplacePricesPriceNewPlanTieredPackagePriceCadence string

const (
	PlanExternalPlanIDVersionNewParamsReplacePricesPriceNewPlanTieredPackagePriceCadenceAnnual     PlanExternalPlanIDVersionNewParamsReplacePricesPriceNewPlanTieredPackagePriceCadence = "annual"
	PlanExternalPlanIDVersionNewParamsReplacePricesPriceNewPlanTieredPackagePriceCadenceSemiAnnual PlanExternalPlanIDVersionNewParamsReplacePricesPriceNewPlanTieredPackagePriceCadence = "semi_annual"
	PlanExternalPlanIDVersionNewParamsReplacePricesPriceNewPlanTieredPackagePriceCadenceMonthly    PlanExternalPlanIDVersionNewParamsReplacePricesPriceNewPlanTieredPackagePriceCadence = "monthly"
	PlanExternalPlanIDVersionNewParamsReplacePricesPriceNewPlanTieredPackagePriceCadenceQuarterly  PlanExternalPlanIDVersionNewParamsReplacePricesPriceNewPlanTieredPackagePriceCadence = "quarterly"
	PlanExternalPlanIDVersionNewParamsReplacePricesPriceNewPlanTieredPackagePriceCadenceOneTime    PlanExternalPlanIDVersionNewParamsReplacePricesPriceNewPlanTieredPackagePriceCadence = "one_time"
	PlanExternalPlanIDVersionNewParamsReplacePricesPriceNewPlanTieredPackagePriceCadenceCustom     PlanExternalPlanIDVersionNewParamsReplacePricesPriceNewPlanTieredPackagePriceCadence = "custom"
)

func (r PlanExternalPlanIDVersionNewParamsReplacePricesPriceNewPlanTieredPackagePriceCadence) IsKnown() bool {
	switch r {
	case PlanExternalPlanIDVersionNewParamsReplacePricesPriceNewPlanTieredPackagePriceCadenceAnnual, PlanExternalPlanIDVersionNewParamsReplacePricesPriceNewPlanTieredPackagePriceCadenceSemiAnnual, PlanExternalPlanIDVersionNewParamsReplacePricesPriceNewPlanTieredPackagePriceCadenceMonthly, PlanExternalPlanIDVersionNewParamsReplacePricesPriceNewPlanTieredPackagePriceCadenceQuarterly, PlanExternalPlanIDVersionNewParamsReplacePricesPriceNewPlanTieredPackagePriceCadenceOneTime, PlanExternalPlanIDVersionNewParamsReplacePricesPriceNewPlanTieredPackagePriceCadenceCustom:
		return true
	}
	return false
}

type PlanExternalPlanIDVersionNewParamsReplacePricesPriceNewPlanTieredPackagePriceModelType string

const (
	PlanExternalPlanIDVersionNewParamsReplacePricesPriceNewPlanTieredPackagePriceModelTypeTieredPackage PlanExternalPlanIDVersionNewParamsReplacePricesPriceNewPlanTieredPackagePriceModelType = "tiered_package"
)

func (r PlanExternalPlanIDVersionNewParamsReplacePricesPriceNewPlanTieredPackagePriceModelType) IsKnown() bool {
	switch r {
	case PlanExternalPlanIDVersionNewParamsReplacePricesPriceNewPlanTieredPackagePriceModelTypeTieredPackage:
		return true
	}
	return false
}

// For custom cadence: specifies the duration of the billing period in days or
// months.
type PlanExternalPlanIDVersionNewParamsReplacePricesPriceNewPlanTieredPackagePriceBillingCycleConfiguration struct {
	// The duration of the billing period.
	Duration param.Field[int64] `json:"duration,required"`
	// The unit of billing period duration.
	DurationUnit param.Field[PlanExternalPlanIDVersionNewParamsReplacePricesPriceNewPlanTieredPackagePriceBillingCycleConfigurationDurationUnit] `json:"duration_unit,required"`
}

func (r PlanExternalPlanIDVersionNewParamsReplacePricesPriceNewPlanTieredPackagePriceBillingCycleConfiguration) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// The unit of billing period duration.
type PlanExternalPlanIDVersionNewParamsReplacePricesPriceNewPlanTieredPackagePriceBillingCycleConfigurationDurationUnit string

const (
	PlanExternalPlanIDVersionNewParamsReplacePricesPriceNewPlanTieredPackagePriceBillingCycleConfigurationDurationUnitDay   PlanExternalPlanIDVersionNewParamsReplacePricesPriceNewPlanTieredPackagePriceBillingCycleConfigurationDurationUnit = "day"
	PlanExternalPlanIDVersionNewParamsReplacePricesPriceNewPlanTieredPackagePriceBillingCycleConfigurationDurationUnitMonth PlanExternalPlanIDVersionNewParamsReplacePricesPriceNewPlanTieredPackagePriceBillingCycleConfigurationDurationUnit = "month"
)

func (r PlanExternalPlanIDVersionNewParamsReplacePricesPriceNewPlanTieredPackagePriceBillingCycleConfigurationDurationUnit) IsKnown() bool {
	switch r {
	case PlanExternalPlanIDVersionNewParamsReplacePricesPriceNewPlanTieredPackagePriceBillingCycleConfigurationDurationUnitDay, PlanExternalPlanIDVersionNewParamsReplacePricesPriceNewPlanTieredPackagePriceBillingCycleConfigurationDurationUnitMonth:
		return true
	}
	return false
}

// For dimensional price: specifies a price group and dimension values
type PlanExternalPlanIDVersionNewParamsReplacePricesPriceNewPlanTieredPackagePriceDimensionalPriceConfiguration struct {
	// The list of dimension values matching (in order) the dimensions of the price
	// group
	DimensionValues param.Field[[]string] `json:"dimension_values,required"`
	// The id of the dimensional price group to include this price in
	DimensionalPriceGroupID param.Field[string] `json:"dimensional_price_group_id"`
	// The external id of the dimensional price group to include this price in
	ExternalDimensionalPriceGroupID param.Field[string] `json:"external_dimensional_price_group_id"`
}

func (r PlanExternalPlanIDVersionNewParamsReplacePricesPriceNewPlanTieredPackagePriceDimensionalPriceConfiguration) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Within each billing cycle, specifies the cadence at which invoices are produced.
// If unspecified, a single invoice is produced per billing cycle.
type PlanExternalPlanIDVersionNewParamsReplacePricesPriceNewPlanTieredPackagePriceInvoicingCycleConfiguration struct {
	// The duration of the billing period.
	Duration param.Field[int64] `json:"duration,required"`
	// The unit of billing period duration.
	DurationUnit param.Field[PlanExternalPlanIDVersionNewParamsReplacePricesPriceNewPlanTieredPackagePriceInvoicingCycleConfigurationDurationUnit] `json:"duration_unit,required"`
}

func (r PlanExternalPlanIDVersionNewParamsReplacePricesPriceNewPlanTieredPackagePriceInvoicingCycleConfiguration) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// The unit of billing period duration.
type PlanExternalPlanIDVersionNewParamsReplacePricesPriceNewPlanTieredPackagePriceInvoicingCycleConfigurationDurationUnit string

const (
	PlanExternalPlanIDVersionNewParamsReplacePricesPriceNewPlanTieredPackagePriceInvoicingCycleConfigurationDurationUnitDay   PlanExternalPlanIDVersionNewParamsReplacePricesPriceNewPlanTieredPackagePriceInvoicingCycleConfigurationDurationUnit = "day"
	PlanExternalPlanIDVersionNewParamsReplacePricesPriceNewPlanTieredPackagePriceInvoicingCycleConfigurationDurationUnitMonth PlanExternalPlanIDVersionNewParamsReplacePricesPriceNewPlanTieredPackagePriceInvoicingCycleConfigurationDurationUnit = "month"
)

func (r PlanExternalPlanIDVersionNewParamsReplacePricesPriceNewPlanTieredPackagePriceInvoicingCycleConfigurationDurationUnit) IsKnown() bool {
	switch r {
	case PlanExternalPlanIDVersionNewParamsReplacePricesPriceNewPlanTieredPackagePriceInvoicingCycleConfigurationDurationUnitDay, PlanExternalPlanIDVersionNewParamsReplacePricesPriceNewPlanTieredPackagePriceInvoicingCycleConfigurationDurationUnitMonth:
		return true
	}
	return false
}

type PlanExternalPlanIDVersionNewParamsReplacePricesPriceNewPlanTieredWithMinimumPrice struct {
	// The cadence to bill for this price on.
	Cadence param.Field[PlanExternalPlanIDVersionNewParamsReplacePricesPriceNewPlanTieredWithMinimumPriceCadence] `json:"cadence,required"`
	// The id of the item the price will be associated with.
	ItemID    param.Field[string]                                                                                     `json:"item_id,required"`
	ModelType param.Field[PlanExternalPlanIDVersionNewParamsReplacePricesPriceNewPlanTieredWithMinimumPriceModelType] `json:"model_type,required"`
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
	BillingCycleConfiguration param.Field[PlanExternalPlanIDVersionNewParamsReplacePricesPriceNewPlanTieredWithMinimumPriceBillingCycleConfiguration] `json:"billing_cycle_configuration"`
	// The per unit conversion rate of the price currency to the invoicing currency.
	ConversionRate param.Field[float64] `json:"conversion_rate"`
	// An ISO 4217 currency string, or custom pricing unit identifier, in which this
	// price is billed.
	Currency param.Field[string] `json:"currency"`
	// For dimensional price: specifies a price group and dimension values
	DimensionalPriceConfiguration param.Field[PlanExternalPlanIDVersionNewParamsReplacePricesPriceNewPlanTieredWithMinimumPriceDimensionalPriceConfiguration] `json:"dimensional_price_configuration"`
	// An alias for the price.
	ExternalPriceID param.Field[string] `json:"external_price_id"`
	// If the Price represents a fixed cost, this represents the quantity of units
	// applied.
	FixedPriceQuantity param.Field[float64] `json:"fixed_price_quantity"`
	// The property used to group this price on an invoice
	InvoiceGroupingKey param.Field[string] `json:"invoice_grouping_key"`
	// Within each billing cycle, specifies the cadence at which invoices are produced.
	// If unspecified, a single invoice is produced per billing cycle.
	InvoicingCycleConfiguration param.Field[PlanExternalPlanIDVersionNewParamsReplacePricesPriceNewPlanTieredWithMinimumPriceInvoicingCycleConfiguration] `json:"invoicing_cycle_configuration"`
	// User-specified key/value pairs for the resource. Individual keys can be removed
	// by setting the value to `null`, and the entire metadata mapping can be cleared
	// by setting `metadata` to `null`.
	Metadata param.Field[map[string]string] `json:"metadata"`
}

func (r PlanExternalPlanIDVersionNewParamsReplacePricesPriceNewPlanTieredWithMinimumPrice) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r PlanExternalPlanIDVersionNewParamsReplacePricesPriceNewPlanTieredWithMinimumPrice) implementsPlanExternalPlanIDVersionNewParamsReplacePricesPriceUnion() {
}

// The cadence to bill for this price on.
type PlanExternalPlanIDVersionNewParamsReplacePricesPriceNewPlanTieredWithMinimumPriceCadence string

const (
	PlanExternalPlanIDVersionNewParamsReplacePricesPriceNewPlanTieredWithMinimumPriceCadenceAnnual     PlanExternalPlanIDVersionNewParamsReplacePricesPriceNewPlanTieredWithMinimumPriceCadence = "annual"
	PlanExternalPlanIDVersionNewParamsReplacePricesPriceNewPlanTieredWithMinimumPriceCadenceSemiAnnual PlanExternalPlanIDVersionNewParamsReplacePricesPriceNewPlanTieredWithMinimumPriceCadence = "semi_annual"
	PlanExternalPlanIDVersionNewParamsReplacePricesPriceNewPlanTieredWithMinimumPriceCadenceMonthly    PlanExternalPlanIDVersionNewParamsReplacePricesPriceNewPlanTieredWithMinimumPriceCadence = "monthly"
	PlanExternalPlanIDVersionNewParamsReplacePricesPriceNewPlanTieredWithMinimumPriceCadenceQuarterly  PlanExternalPlanIDVersionNewParamsReplacePricesPriceNewPlanTieredWithMinimumPriceCadence = "quarterly"
	PlanExternalPlanIDVersionNewParamsReplacePricesPriceNewPlanTieredWithMinimumPriceCadenceOneTime    PlanExternalPlanIDVersionNewParamsReplacePricesPriceNewPlanTieredWithMinimumPriceCadence = "one_time"
	PlanExternalPlanIDVersionNewParamsReplacePricesPriceNewPlanTieredWithMinimumPriceCadenceCustom     PlanExternalPlanIDVersionNewParamsReplacePricesPriceNewPlanTieredWithMinimumPriceCadence = "custom"
)

func (r PlanExternalPlanIDVersionNewParamsReplacePricesPriceNewPlanTieredWithMinimumPriceCadence) IsKnown() bool {
	switch r {
	case PlanExternalPlanIDVersionNewParamsReplacePricesPriceNewPlanTieredWithMinimumPriceCadenceAnnual, PlanExternalPlanIDVersionNewParamsReplacePricesPriceNewPlanTieredWithMinimumPriceCadenceSemiAnnual, PlanExternalPlanIDVersionNewParamsReplacePricesPriceNewPlanTieredWithMinimumPriceCadenceMonthly, PlanExternalPlanIDVersionNewParamsReplacePricesPriceNewPlanTieredWithMinimumPriceCadenceQuarterly, PlanExternalPlanIDVersionNewParamsReplacePricesPriceNewPlanTieredWithMinimumPriceCadenceOneTime, PlanExternalPlanIDVersionNewParamsReplacePricesPriceNewPlanTieredWithMinimumPriceCadenceCustom:
		return true
	}
	return false
}

type PlanExternalPlanIDVersionNewParamsReplacePricesPriceNewPlanTieredWithMinimumPriceModelType string

const (
	PlanExternalPlanIDVersionNewParamsReplacePricesPriceNewPlanTieredWithMinimumPriceModelTypeTieredWithMinimum PlanExternalPlanIDVersionNewParamsReplacePricesPriceNewPlanTieredWithMinimumPriceModelType = "tiered_with_minimum"
)

func (r PlanExternalPlanIDVersionNewParamsReplacePricesPriceNewPlanTieredWithMinimumPriceModelType) IsKnown() bool {
	switch r {
	case PlanExternalPlanIDVersionNewParamsReplacePricesPriceNewPlanTieredWithMinimumPriceModelTypeTieredWithMinimum:
		return true
	}
	return false
}

// For custom cadence: specifies the duration of the billing period in days or
// months.
type PlanExternalPlanIDVersionNewParamsReplacePricesPriceNewPlanTieredWithMinimumPriceBillingCycleConfiguration struct {
	// The duration of the billing period.
	Duration param.Field[int64] `json:"duration,required"`
	// The unit of billing period duration.
	DurationUnit param.Field[PlanExternalPlanIDVersionNewParamsReplacePricesPriceNewPlanTieredWithMinimumPriceBillingCycleConfigurationDurationUnit] `json:"duration_unit,required"`
}

func (r PlanExternalPlanIDVersionNewParamsReplacePricesPriceNewPlanTieredWithMinimumPriceBillingCycleConfiguration) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// The unit of billing period duration.
type PlanExternalPlanIDVersionNewParamsReplacePricesPriceNewPlanTieredWithMinimumPriceBillingCycleConfigurationDurationUnit string

const (
	PlanExternalPlanIDVersionNewParamsReplacePricesPriceNewPlanTieredWithMinimumPriceBillingCycleConfigurationDurationUnitDay   PlanExternalPlanIDVersionNewParamsReplacePricesPriceNewPlanTieredWithMinimumPriceBillingCycleConfigurationDurationUnit = "day"
	PlanExternalPlanIDVersionNewParamsReplacePricesPriceNewPlanTieredWithMinimumPriceBillingCycleConfigurationDurationUnitMonth PlanExternalPlanIDVersionNewParamsReplacePricesPriceNewPlanTieredWithMinimumPriceBillingCycleConfigurationDurationUnit = "month"
)

func (r PlanExternalPlanIDVersionNewParamsReplacePricesPriceNewPlanTieredWithMinimumPriceBillingCycleConfigurationDurationUnit) IsKnown() bool {
	switch r {
	case PlanExternalPlanIDVersionNewParamsReplacePricesPriceNewPlanTieredWithMinimumPriceBillingCycleConfigurationDurationUnitDay, PlanExternalPlanIDVersionNewParamsReplacePricesPriceNewPlanTieredWithMinimumPriceBillingCycleConfigurationDurationUnitMonth:
		return true
	}
	return false
}

// For dimensional price: specifies a price group and dimension values
type PlanExternalPlanIDVersionNewParamsReplacePricesPriceNewPlanTieredWithMinimumPriceDimensionalPriceConfiguration struct {
	// The list of dimension values matching (in order) the dimensions of the price
	// group
	DimensionValues param.Field[[]string] `json:"dimension_values,required"`
	// The id of the dimensional price group to include this price in
	DimensionalPriceGroupID param.Field[string] `json:"dimensional_price_group_id"`
	// The external id of the dimensional price group to include this price in
	ExternalDimensionalPriceGroupID param.Field[string] `json:"external_dimensional_price_group_id"`
}

func (r PlanExternalPlanIDVersionNewParamsReplacePricesPriceNewPlanTieredWithMinimumPriceDimensionalPriceConfiguration) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Within each billing cycle, specifies the cadence at which invoices are produced.
// If unspecified, a single invoice is produced per billing cycle.
type PlanExternalPlanIDVersionNewParamsReplacePricesPriceNewPlanTieredWithMinimumPriceInvoicingCycleConfiguration struct {
	// The duration of the billing period.
	Duration param.Field[int64] `json:"duration,required"`
	// The unit of billing period duration.
	DurationUnit param.Field[PlanExternalPlanIDVersionNewParamsReplacePricesPriceNewPlanTieredWithMinimumPriceInvoicingCycleConfigurationDurationUnit] `json:"duration_unit,required"`
}

func (r PlanExternalPlanIDVersionNewParamsReplacePricesPriceNewPlanTieredWithMinimumPriceInvoicingCycleConfiguration) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// The unit of billing period duration.
type PlanExternalPlanIDVersionNewParamsReplacePricesPriceNewPlanTieredWithMinimumPriceInvoicingCycleConfigurationDurationUnit string

const (
	PlanExternalPlanIDVersionNewParamsReplacePricesPriceNewPlanTieredWithMinimumPriceInvoicingCycleConfigurationDurationUnitDay   PlanExternalPlanIDVersionNewParamsReplacePricesPriceNewPlanTieredWithMinimumPriceInvoicingCycleConfigurationDurationUnit = "day"
	PlanExternalPlanIDVersionNewParamsReplacePricesPriceNewPlanTieredWithMinimumPriceInvoicingCycleConfigurationDurationUnitMonth PlanExternalPlanIDVersionNewParamsReplacePricesPriceNewPlanTieredWithMinimumPriceInvoicingCycleConfigurationDurationUnit = "month"
)

func (r PlanExternalPlanIDVersionNewParamsReplacePricesPriceNewPlanTieredWithMinimumPriceInvoicingCycleConfigurationDurationUnit) IsKnown() bool {
	switch r {
	case PlanExternalPlanIDVersionNewParamsReplacePricesPriceNewPlanTieredWithMinimumPriceInvoicingCycleConfigurationDurationUnitDay, PlanExternalPlanIDVersionNewParamsReplacePricesPriceNewPlanTieredWithMinimumPriceInvoicingCycleConfigurationDurationUnitMonth:
		return true
	}
	return false
}

type PlanExternalPlanIDVersionNewParamsReplacePricesPriceNewPlanUnitWithPercentPrice struct {
	// The cadence to bill for this price on.
	Cadence param.Field[PlanExternalPlanIDVersionNewParamsReplacePricesPriceNewPlanUnitWithPercentPriceCadence] `json:"cadence,required"`
	// The id of the item the price will be associated with.
	ItemID    param.Field[string]                                                                                   `json:"item_id,required"`
	ModelType param.Field[PlanExternalPlanIDVersionNewParamsReplacePricesPriceNewPlanUnitWithPercentPriceModelType] `json:"model_type,required"`
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
	BillingCycleConfiguration param.Field[PlanExternalPlanIDVersionNewParamsReplacePricesPriceNewPlanUnitWithPercentPriceBillingCycleConfiguration] `json:"billing_cycle_configuration"`
	// The per unit conversion rate of the price currency to the invoicing currency.
	ConversionRate param.Field[float64] `json:"conversion_rate"`
	// An ISO 4217 currency string, or custom pricing unit identifier, in which this
	// price is billed.
	Currency param.Field[string] `json:"currency"`
	// For dimensional price: specifies a price group and dimension values
	DimensionalPriceConfiguration param.Field[PlanExternalPlanIDVersionNewParamsReplacePricesPriceNewPlanUnitWithPercentPriceDimensionalPriceConfiguration] `json:"dimensional_price_configuration"`
	// An alias for the price.
	ExternalPriceID param.Field[string] `json:"external_price_id"`
	// If the Price represents a fixed cost, this represents the quantity of units
	// applied.
	FixedPriceQuantity param.Field[float64] `json:"fixed_price_quantity"`
	// The property used to group this price on an invoice
	InvoiceGroupingKey param.Field[string] `json:"invoice_grouping_key"`
	// Within each billing cycle, specifies the cadence at which invoices are produced.
	// If unspecified, a single invoice is produced per billing cycle.
	InvoicingCycleConfiguration param.Field[PlanExternalPlanIDVersionNewParamsReplacePricesPriceNewPlanUnitWithPercentPriceInvoicingCycleConfiguration] `json:"invoicing_cycle_configuration"`
	// User-specified key/value pairs for the resource. Individual keys can be removed
	// by setting the value to `null`, and the entire metadata mapping can be cleared
	// by setting `metadata` to `null`.
	Metadata param.Field[map[string]string] `json:"metadata"`
}

func (r PlanExternalPlanIDVersionNewParamsReplacePricesPriceNewPlanUnitWithPercentPrice) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r PlanExternalPlanIDVersionNewParamsReplacePricesPriceNewPlanUnitWithPercentPrice) implementsPlanExternalPlanIDVersionNewParamsReplacePricesPriceUnion() {
}

// The cadence to bill for this price on.
type PlanExternalPlanIDVersionNewParamsReplacePricesPriceNewPlanUnitWithPercentPriceCadence string

const (
	PlanExternalPlanIDVersionNewParamsReplacePricesPriceNewPlanUnitWithPercentPriceCadenceAnnual     PlanExternalPlanIDVersionNewParamsReplacePricesPriceNewPlanUnitWithPercentPriceCadence = "annual"
	PlanExternalPlanIDVersionNewParamsReplacePricesPriceNewPlanUnitWithPercentPriceCadenceSemiAnnual PlanExternalPlanIDVersionNewParamsReplacePricesPriceNewPlanUnitWithPercentPriceCadence = "semi_annual"
	PlanExternalPlanIDVersionNewParamsReplacePricesPriceNewPlanUnitWithPercentPriceCadenceMonthly    PlanExternalPlanIDVersionNewParamsReplacePricesPriceNewPlanUnitWithPercentPriceCadence = "monthly"
	PlanExternalPlanIDVersionNewParamsReplacePricesPriceNewPlanUnitWithPercentPriceCadenceQuarterly  PlanExternalPlanIDVersionNewParamsReplacePricesPriceNewPlanUnitWithPercentPriceCadence = "quarterly"
	PlanExternalPlanIDVersionNewParamsReplacePricesPriceNewPlanUnitWithPercentPriceCadenceOneTime    PlanExternalPlanIDVersionNewParamsReplacePricesPriceNewPlanUnitWithPercentPriceCadence = "one_time"
	PlanExternalPlanIDVersionNewParamsReplacePricesPriceNewPlanUnitWithPercentPriceCadenceCustom     PlanExternalPlanIDVersionNewParamsReplacePricesPriceNewPlanUnitWithPercentPriceCadence = "custom"
)

func (r PlanExternalPlanIDVersionNewParamsReplacePricesPriceNewPlanUnitWithPercentPriceCadence) IsKnown() bool {
	switch r {
	case PlanExternalPlanIDVersionNewParamsReplacePricesPriceNewPlanUnitWithPercentPriceCadenceAnnual, PlanExternalPlanIDVersionNewParamsReplacePricesPriceNewPlanUnitWithPercentPriceCadenceSemiAnnual, PlanExternalPlanIDVersionNewParamsReplacePricesPriceNewPlanUnitWithPercentPriceCadenceMonthly, PlanExternalPlanIDVersionNewParamsReplacePricesPriceNewPlanUnitWithPercentPriceCadenceQuarterly, PlanExternalPlanIDVersionNewParamsReplacePricesPriceNewPlanUnitWithPercentPriceCadenceOneTime, PlanExternalPlanIDVersionNewParamsReplacePricesPriceNewPlanUnitWithPercentPriceCadenceCustom:
		return true
	}
	return false
}

type PlanExternalPlanIDVersionNewParamsReplacePricesPriceNewPlanUnitWithPercentPriceModelType string

const (
	PlanExternalPlanIDVersionNewParamsReplacePricesPriceNewPlanUnitWithPercentPriceModelTypeUnitWithPercent PlanExternalPlanIDVersionNewParamsReplacePricesPriceNewPlanUnitWithPercentPriceModelType = "unit_with_percent"
)

func (r PlanExternalPlanIDVersionNewParamsReplacePricesPriceNewPlanUnitWithPercentPriceModelType) IsKnown() bool {
	switch r {
	case PlanExternalPlanIDVersionNewParamsReplacePricesPriceNewPlanUnitWithPercentPriceModelTypeUnitWithPercent:
		return true
	}
	return false
}

// For custom cadence: specifies the duration of the billing period in days or
// months.
type PlanExternalPlanIDVersionNewParamsReplacePricesPriceNewPlanUnitWithPercentPriceBillingCycleConfiguration struct {
	// The duration of the billing period.
	Duration param.Field[int64] `json:"duration,required"`
	// The unit of billing period duration.
	DurationUnit param.Field[PlanExternalPlanIDVersionNewParamsReplacePricesPriceNewPlanUnitWithPercentPriceBillingCycleConfigurationDurationUnit] `json:"duration_unit,required"`
}

func (r PlanExternalPlanIDVersionNewParamsReplacePricesPriceNewPlanUnitWithPercentPriceBillingCycleConfiguration) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// The unit of billing period duration.
type PlanExternalPlanIDVersionNewParamsReplacePricesPriceNewPlanUnitWithPercentPriceBillingCycleConfigurationDurationUnit string

const (
	PlanExternalPlanIDVersionNewParamsReplacePricesPriceNewPlanUnitWithPercentPriceBillingCycleConfigurationDurationUnitDay   PlanExternalPlanIDVersionNewParamsReplacePricesPriceNewPlanUnitWithPercentPriceBillingCycleConfigurationDurationUnit = "day"
	PlanExternalPlanIDVersionNewParamsReplacePricesPriceNewPlanUnitWithPercentPriceBillingCycleConfigurationDurationUnitMonth PlanExternalPlanIDVersionNewParamsReplacePricesPriceNewPlanUnitWithPercentPriceBillingCycleConfigurationDurationUnit = "month"
)

func (r PlanExternalPlanIDVersionNewParamsReplacePricesPriceNewPlanUnitWithPercentPriceBillingCycleConfigurationDurationUnit) IsKnown() bool {
	switch r {
	case PlanExternalPlanIDVersionNewParamsReplacePricesPriceNewPlanUnitWithPercentPriceBillingCycleConfigurationDurationUnitDay, PlanExternalPlanIDVersionNewParamsReplacePricesPriceNewPlanUnitWithPercentPriceBillingCycleConfigurationDurationUnitMonth:
		return true
	}
	return false
}

// For dimensional price: specifies a price group and dimension values
type PlanExternalPlanIDVersionNewParamsReplacePricesPriceNewPlanUnitWithPercentPriceDimensionalPriceConfiguration struct {
	// The list of dimension values matching (in order) the dimensions of the price
	// group
	DimensionValues param.Field[[]string] `json:"dimension_values,required"`
	// The id of the dimensional price group to include this price in
	DimensionalPriceGroupID param.Field[string] `json:"dimensional_price_group_id"`
	// The external id of the dimensional price group to include this price in
	ExternalDimensionalPriceGroupID param.Field[string] `json:"external_dimensional_price_group_id"`
}

func (r PlanExternalPlanIDVersionNewParamsReplacePricesPriceNewPlanUnitWithPercentPriceDimensionalPriceConfiguration) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Within each billing cycle, specifies the cadence at which invoices are produced.
// If unspecified, a single invoice is produced per billing cycle.
type PlanExternalPlanIDVersionNewParamsReplacePricesPriceNewPlanUnitWithPercentPriceInvoicingCycleConfiguration struct {
	// The duration of the billing period.
	Duration param.Field[int64] `json:"duration,required"`
	// The unit of billing period duration.
	DurationUnit param.Field[PlanExternalPlanIDVersionNewParamsReplacePricesPriceNewPlanUnitWithPercentPriceInvoicingCycleConfigurationDurationUnit] `json:"duration_unit,required"`
}

func (r PlanExternalPlanIDVersionNewParamsReplacePricesPriceNewPlanUnitWithPercentPriceInvoicingCycleConfiguration) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// The unit of billing period duration.
type PlanExternalPlanIDVersionNewParamsReplacePricesPriceNewPlanUnitWithPercentPriceInvoicingCycleConfigurationDurationUnit string

const (
	PlanExternalPlanIDVersionNewParamsReplacePricesPriceNewPlanUnitWithPercentPriceInvoicingCycleConfigurationDurationUnitDay   PlanExternalPlanIDVersionNewParamsReplacePricesPriceNewPlanUnitWithPercentPriceInvoicingCycleConfigurationDurationUnit = "day"
	PlanExternalPlanIDVersionNewParamsReplacePricesPriceNewPlanUnitWithPercentPriceInvoicingCycleConfigurationDurationUnitMonth PlanExternalPlanIDVersionNewParamsReplacePricesPriceNewPlanUnitWithPercentPriceInvoicingCycleConfigurationDurationUnit = "month"
)

func (r PlanExternalPlanIDVersionNewParamsReplacePricesPriceNewPlanUnitWithPercentPriceInvoicingCycleConfigurationDurationUnit) IsKnown() bool {
	switch r {
	case PlanExternalPlanIDVersionNewParamsReplacePricesPriceNewPlanUnitWithPercentPriceInvoicingCycleConfigurationDurationUnitDay, PlanExternalPlanIDVersionNewParamsReplacePricesPriceNewPlanUnitWithPercentPriceInvoicingCycleConfigurationDurationUnitMonth:
		return true
	}
	return false
}

type PlanExternalPlanIDVersionNewParamsReplacePricesPriceNewPlanPackageWithAllocationPrice struct {
	// The cadence to bill for this price on.
	Cadence param.Field[PlanExternalPlanIDVersionNewParamsReplacePricesPriceNewPlanPackageWithAllocationPriceCadence] `json:"cadence,required"`
	// The id of the item the price will be associated with.
	ItemID    param.Field[string]                                                                                         `json:"item_id,required"`
	ModelType param.Field[PlanExternalPlanIDVersionNewParamsReplacePricesPriceNewPlanPackageWithAllocationPriceModelType] `json:"model_type,required"`
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
	BillingCycleConfiguration param.Field[PlanExternalPlanIDVersionNewParamsReplacePricesPriceNewPlanPackageWithAllocationPriceBillingCycleConfiguration] `json:"billing_cycle_configuration"`
	// The per unit conversion rate of the price currency to the invoicing currency.
	ConversionRate param.Field[float64] `json:"conversion_rate"`
	// An ISO 4217 currency string, or custom pricing unit identifier, in which this
	// price is billed.
	Currency param.Field[string] `json:"currency"`
	// For dimensional price: specifies a price group and dimension values
	DimensionalPriceConfiguration param.Field[PlanExternalPlanIDVersionNewParamsReplacePricesPriceNewPlanPackageWithAllocationPriceDimensionalPriceConfiguration] `json:"dimensional_price_configuration"`
	// An alias for the price.
	ExternalPriceID param.Field[string] `json:"external_price_id"`
	// If the Price represents a fixed cost, this represents the quantity of units
	// applied.
	FixedPriceQuantity param.Field[float64] `json:"fixed_price_quantity"`
	// The property used to group this price on an invoice
	InvoiceGroupingKey param.Field[string] `json:"invoice_grouping_key"`
	// Within each billing cycle, specifies the cadence at which invoices are produced.
	// If unspecified, a single invoice is produced per billing cycle.
	InvoicingCycleConfiguration param.Field[PlanExternalPlanIDVersionNewParamsReplacePricesPriceNewPlanPackageWithAllocationPriceInvoicingCycleConfiguration] `json:"invoicing_cycle_configuration"`
	// User-specified key/value pairs for the resource. Individual keys can be removed
	// by setting the value to `null`, and the entire metadata mapping can be cleared
	// by setting `metadata` to `null`.
	Metadata param.Field[map[string]string] `json:"metadata"`
}

func (r PlanExternalPlanIDVersionNewParamsReplacePricesPriceNewPlanPackageWithAllocationPrice) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r PlanExternalPlanIDVersionNewParamsReplacePricesPriceNewPlanPackageWithAllocationPrice) implementsPlanExternalPlanIDVersionNewParamsReplacePricesPriceUnion() {
}

// The cadence to bill for this price on.
type PlanExternalPlanIDVersionNewParamsReplacePricesPriceNewPlanPackageWithAllocationPriceCadence string

const (
	PlanExternalPlanIDVersionNewParamsReplacePricesPriceNewPlanPackageWithAllocationPriceCadenceAnnual     PlanExternalPlanIDVersionNewParamsReplacePricesPriceNewPlanPackageWithAllocationPriceCadence = "annual"
	PlanExternalPlanIDVersionNewParamsReplacePricesPriceNewPlanPackageWithAllocationPriceCadenceSemiAnnual PlanExternalPlanIDVersionNewParamsReplacePricesPriceNewPlanPackageWithAllocationPriceCadence = "semi_annual"
	PlanExternalPlanIDVersionNewParamsReplacePricesPriceNewPlanPackageWithAllocationPriceCadenceMonthly    PlanExternalPlanIDVersionNewParamsReplacePricesPriceNewPlanPackageWithAllocationPriceCadence = "monthly"
	PlanExternalPlanIDVersionNewParamsReplacePricesPriceNewPlanPackageWithAllocationPriceCadenceQuarterly  PlanExternalPlanIDVersionNewParamsReplacePricesPriceNewPlanPackageWithAllocationPriceCadence = "quarterly"
	PlanExternalPlanIDVersionNewParamsReplacePricesPriceNewPlanPackageWithAllocationPriceCadenceOneTime    PlanExternalPlanIDVersionNewParamsReplacePricesPriceNewPlanPackageWithAllocationPriceCadence = "one_time"
	PlanExternalPlanIDVersionNewParamsReplacePricesPriceNewPlanPackageWithAllocationPriceCadenceCustom     PlanExternalPlanIDVersionNewParamsReplacePricesPriceNewPlanPackageWithAllocationPriceCadence = "custom"
)

func (r PlanExternalPlanIDVersionNewParamsReplacePricesPriceNewPlanPackageWithAllocationPriceCadence) IsKnown() bool {
	switch r {
	case PlanExternalPlanIDVersionNewParamsReplacePricesPriceNewPlanPackageWithAllocationPriceCadenceAnnual, PlanExternalPlanIDVersionNewParamsReplacePricesPriceNewPlanPackageWithAllocationPriceCadenceSemiAnnual, PlanExternalPlanIDVersionNewParamsReplacePricesPriceNewPlanPackageWithAllocationPriceCadenceMonthly, PlanExternalPlanIDVersionNewParamsReplacePricesPriceNewPlanPackageWithAllocationPriceCadenceQuarterly, PlanExternalPlanIDVersionNewParamsReplacePricesPriceNewPlanPackageWithAllocationPriceCadenceOneTime, PlanExternalPlanIDVersionNewParamsReplacePricesPriceNewPlanPackageWithAllocationPriceCadenceCustom:
		return true
	}
	return false
}

type PlanExternalPlanIDVersionNewParamsReplacePricesPriceNewPlanPackageWithAllocationPriceModelType string

const (
	PlanExternalPlanIDVersionNewParamsReplacePricesPriceNewPlanPackageWithAllocationPriceModelTypePackageWithAllocation PlanExternalPlanIDVersionNewParamsReplacePricesPriceNewPlanPackageWithAllocationPriceModelType = "package_with_allocation"
)

func (r PlanExternalPlanIDVersionNewParamsReplacePricesPriceNewPlanPackageWithAllocationPriceModelType) IsKnown() bool {
	switch r {
	case PlanExternalPlanIDVersionNewParamsReplacePricesPriceNewPlanPackageWithAllocationPriceModelTypePackageWithAllocation:
		return true
	}
	return false
}

// For custom cadence: specifies the duration of the billing period in days or
// months.
type PlanExternalPlanIDVersionNewParamsReplacePricesPriceNewPlanPackageWithAllocationPriceBillingCycleConfiguration struct {
	// The duration of the billing period.
	Duration param.Field[int64] `json:"duration,required"`
	// The unit of billing period duration.
	DurationUnit param.Field[PlanExternalPlanIDVersionNewParamsReplacePricesPriceNewPlanPackageWithAllocationPriceBillingCycleConfigurationDurationUnit] `json:"duration_unit,required"`
}

func (r PlanExternalPlanIDVersionNewParamsReplacePricesPriceNewPlanPackageWithAllocationPriceBillingCycleConfiguration) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// The unit of billing period duration.
type PlanExternalPlanIDVersionNewParamsReplacePricesPriceNewPlanPackageWithAllocationPriceBillingCycleConfigurationDurationUnit string

const (
	PlanExternalPlanIDVersionNewParamsReplacePricesPriceNewPlanPackageWithAllocationPriceBillingCycleConfigurationDurationUnitDay   PlanExternalPlanIDVersionNewParamsReplacePricesPriceNewPlanPackageWithAllocationPriceBillingCycleConfigurationDurationUnit = "day"
	PlanExternalPlanIDVersionNewParamsReplacePricesPriceNewPlanPackageWithAllocationPriceBillingCycleConfigurationDurationUnitMonth PlanExternalPlanIDVersionNewParamsReplacePricesPriceNewPlanPackageWithAllocationPriceBillingCycleConfigurationDurationUnit = "month"
)

func (r PlanExternalPlanIDVersionNewParamsReplacePricesPriceNewPlanPackageWithAllocationPriceBillingCycleConfigurationDurationUnit) IsKnown() bool {
	switch r {
	case PlanExternalPlanIDVersionNewParamsReplacePricesPriceNewPlanPackageWithAllocationPriceBillingCycleConfigurationDurationUnitDay, PlanExternalPlanIDVersionNewParamsReplacePricesPriceNewPlanPackageWithAllocationPriceBillingCycleConfigurationDurationUnitMonth:
		return true
	}
	return false
}

// For dimensional price: specifies a price group and dimension values
type PlanExternalPlanIDVersionNewParamsReplacePricesPriceNewPlanPackageWithAllocationPriceDimensionalPriceConfiguration struct {
	// The list of dimension values matching (in order) the dimensions of the price
	// group
	DimensionValues param.Field[[]string] `json:"dimension_values,required"`
	// The id of the dimensional price group to include this price in
	DimensionalPriceGroupID param.Field[string] `json:"dimensional_price_group_id"`
	// The external id of the dimensional price group to include this price in
	ExternalDimensionalPriceGroupID param.Field[string] `json:"external_dimensional_price_group_id"`
}

func (r PlanExternalPlanIDVersionNewParamsReplacePricesPriceNewPlanPackageWithAllocationPriceDimensionalPriceConfiguration) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Within each billing cycle, specifies the cadence at which invoices are produced.
// If unspecified, a single invoice is produced per billing cycle.
type PlanExternalPlanIDVersionNewParamsReplacePricesPriceNewPlanPackageWithAllocationPriceInvoicingCycleConfiguration struct {
	// The duration of the billing period.
	Duration param.Field[int64] `json:"duration,required"`
	// The unit of billing period duration.
	DurationUnit param.Field[PlanExternalPlanIDVersionNewParamsReplacePricesPriceNewPlanPackageWithAllocationPriceInvoicingCycleConfigurationDurationUnit] `json:"duration_unit,required"`
}

func (r PlanExternalPlanIDVersionNewParamsReplacePricesPriceNewPlanPackageWithAllocationPriceInvoicingCycleConfiguration) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// The unit of billing period duration.
type PlanExternalPlanIDVersionNewParamsReplacePricesPriceNewPlanPackageWithAllocationPriceInvoicingCycleConfigurationDurationUnit string

const (
	PlanExternalPlanIDVersionNewParamsReplacePricesPriceNewPlanPackageWithAllocationPriceInvoicingCycleConfigurationDurationUnitDay   PlanExternalPlanIDVersionNewParamsReplacePricesPriceNewPlanPackageWithAllocationPriceInvoicingCycleConfigurationDurationUnit = "day"
	PlanExternalPlanIDVersionNewParamsReplacePricesPriceNewPlanPackageWithAllocationPriceInvoicingCycleConfigurationDurationUnitMonth PlanExternalPlanIDVersionNewParamsReplacePricesPriceNewPlanPackageWithAllocationPriceInvoicingCycleConfigurationDurationUnit = "month"
)

func (r PlanExternalPlanIDVersionNewParamsReplacePricesPriceNewPlanPackageWithAllocationPriceInvoicingCycleConfigurationDurationUnit) IsKnown() bool {
	switch r {
	case PlanExternalPlanIDVersionNewParamsReplacePricesPriceNewPlanPackageWithAllocationPriceInvoicingCycleConfigurationDurationUnitDay, PlanExternalPlanIDVersionNewParamsReplacePricesPriceNewPlanPackageWithAllocationPriceInvoicingCycleConfigurationDurationUnitMonth:
		return true
	}
	return false
}

type PlanExternalPlanIDVersionNewParamsReplacePricesPriceNewPlanTierWithProrationPrice struct {
	// The cadence to bill for this price on.
	Cadence param.Field[PlanExternalPlanIDVersionNewParamsReplacePricesPriceNewPlanTierWithProrationPriceCadence] `json:"cadence,required"`
	// The id of the item the price will be associated with.
	ItemID    param.Field[string]                                                                                     `json:"item_id,required"`
	ModelType param.Field[PlanExternalPlanIDVersionNewParamsReplacePricesPriceNewPlanTierWithProrationPriceModelType] `json:"model_type,required"`
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
	BillingCycleConfiguration param.Field[PlanExternalPlanIDVersionNewParamsReplacePricesPriceNewPlanTierWithProrationPriceBillingCycleConfiguration] `json:"billing_cycle_configuration"`
	// The per unit conversion rate of the price currency to the invoicing currency.
	ConversionRate param.Field[float64] `json:"conversion_rate"`
	// An ISO 4217 currency string, or custom pricing unit identifier, in which this
	// price is billed.
	Currency param.Field[string] `json:"currency"`
	// For dimensional price: specifies a price group and dimension values
	DimensionalPriceConfiguration param.Field[PlanExternalPlanIDVersionNewParamsReplacePricesPriceNewPlanTierWithProrationPriceDimensionalPriceConfiguration] `json:"dimensional_price_configuration"`
	// An alias for the price.
	ExternalPriceID param.Field[string] `json:"external_price_id"`
	// If the Price represents a fixed cost, this represents the quantity of units
	// applied.
	FixedPriceQuantity param.Field[float64] `json:"fixed_price_quantity"`
	// The property used to group this price on an invoice
	InvoiceGroupingKey param.Field[string] `json:"invoice_grouping_key"`
	// Within each billing cycle, specifies the cadence at which invoices are produced.
	// If unspecified, a single invoice is produced per billing cycle.
	InvoicingCycleConfiguration param.Field[PlanExternalPlanIDVersionNewParamsReplacePricesPriceNewPlanTierWithProrationPriceInvoicingCycleConfiguration] `json:"invoicing_cycle_configuration"`
	// User-specified key/value pairs for the resource. Individual keys can be removed
	// by setting the value to `null`, and the entire metadata mapping can be cleared
	// by setting `metadata` to `null`.
	Metadata param.Field[map[string]string] `json:"metadata"`
}

func (r PlanExternalPlanIDVersionNewParamsReplacePricesPriceNewPlanTierWithProrationPrice) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r PlanExternalPlanIDVersionNewParamsReplacePricesPriceNewPlanTierWithProrationPrice) implementsPlanExternalPlanIDVersionNewParamsReplacePricesPriceUnion() {
}

// The cadence to bill for this price on.
type PlanExternalPlanIDVersionNewParamsReplacePricesPriceNewPlanTierWithProrationPriceCadence string

const (
	PlanExternalPlanIDVersionNewParamsReplacePricesPriceNewPlanTierWithProrationPriceCadenceAnnual     PlanExternalPlanIDVersionNewParamsReplacePricesPriceNewPlanTierWithProrationPriceCadence = "annual"
	PlanExternalPlanIDVersionNewParamsReplacePricesPriceNewPlanTierWithProrationPriceCadenceSemiAnnual PlanExternalPlanIDVersionNewParamsReplacePricesPriceNewPlanTierWithProrationPriceCadence = "semi_annual"
	PlanExternalPlanIDVersionNewParamsReplacePricesPriceNewPlanTierWithProrationPriceCadenceMonthly    PlanExternalPlanIDVersionNewParamsReplacePricesPriceNewPlanTierWithProrationPriceCadence = "monthly"
	PlanExternalPlanIDVersionNewParamsReplacePricesPriceNewPlanTierWithProrationPriceCadenceQuarterly  PlanExternalPlanIDVersionNewParamsReplacePricesPriceNewPlanTierWithProrationPriceCadence = "quarterly"
	PlanExternalPlanIDVersionNewParamsReplacePricesPriceNewPlanTierWithProrationPriceCadenceOneTime    PlanExternalPlanIDVersionNewParamsReplacePricesPriceNewPlanTierWithProrationPriceCadence = "one_time"
	PlanExternalPlanIDVersionNewParamsReplacePricesPriceNewPlanTierWithProrationPriceCadenceCustom     PlanExternalPlanIDVersionNewParamsReplacePricesPriceNewPlanTierWithProrationPriceCadence = "custom"
)

func (r PlanExternalPlanIDVersionNewParamsReplacePricesPriceNewPlanTierWithProrationPriceCadence) IsKnown() bool {
	switch r {
	case PlanExternalPlanIDVersionNewParamsReplacePricesPriceNewPlanTierWithProrationPriceCadenceAnnual, PlanExternalPlanIDVersionNewParamsReplacePricesPriceNewPlanTierWithProrationPriceCadenceSemiAnnual, PlanExternalPlanIDVersionNewParamsReplacePricesPriceNewPlanTierWithProrationPriceCadenceMonthly, PlanExternalPlanIDVersionNewParamsReplacePricesPriceNewPlanTierWithProrationPriceCadenceQuarterly, PlanExternalPlanIDVersionNewParamsReplacePricesPriceNewPlanTierWithProrationPriceCadenceOneTime, PlanExternalPlanIDVersionNewParamsReplacePricesPriceNewPlanTierWithProrationPriceCadenceCustom:
		return true
	}
	return false
}

type PlanExternalPlanIDVersionNewParamsReplacePricesPriceNewPlanTierWithProrationPriceModelType string

const (
	PlanExternalPlanIDVersionNewParamsReplacePricesPriceNewPlanTierWithProrationPriceModelTypeTieredWithProration PlanExternalPlanIDVersionNewParamsReplacePricesPriceNewPlanTierWithProrationPriceModelType = "tiered_with_proration"
)

func (r PlanExternalPlanIDVersionNewParamsReplacePricesPriceNewPlanTierWithProrationPriceModelType) IsKnown() bool {
	switch r {
	case PlanExternalPlanIDVersionNewParamsReplacePricesPriceNewPlanTierWithProrationPriceModelTypeTieredWithProration:
		return true
	}
	return false
}

// For custom cadence: specifies the duration of the billing period in days or
// months.
type PlanExternalPlanIDVersionNewParamsReplacePricesPriceNewPlanTierWithProrationPriceBillingCycleConfiguration struct {
	// The duration of the billing period.
	Duration param.Field[int64] `json:"duration,required"`
	// The unit of billing period duration.
	DurationUnit param.Field[PlanExternalPlanIDVersionNewParamsReplacePricesPriceNewPlanTierWithProrationPriceBillingCycleConfigurationDurationUnit] `json:"duration_unit,required"`
}

func (r PlanExternalPlanIDVersionNewParamsReplacePricesPriceNewPlanTierWithProrationPriceBillingCycleConfiguration) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// The unit of billing period duration.
type PlanExternalPlanIDVersionNewParamsReplacePricesPriceNewPlanTierWithProrationPriceBillingCycleConfigurationDurationUnit string

const (
	PlanExternalPlanIDVersionNewParamsReplacePricesPriceNewPlanTierWithProrationPriceBillingCycleConfigurationDurationUnitDay   PlanExternalPlanIDVersionNewParamsReplacePricesPriceNewPlanTierWithProrationPriceBillingCycleConfigurationDurationUnit = "day"
	PlanExternalPlanIDVersionNewParamsReplacePricesPriceNewPlanTierWithProrationPriceBillingCycleConfigurationDurationUnitMonth PlanExternalPlanIDVersionNewParamsReplacePricesPriceNewPlanTierWithProrationPriceBillingCycleConfigurationDurationUnit = "month"
)

func (r PlanExternalPlanIDVersionNewParamsReplacePricesPriceNewPlanTierWithProrationPriceBillingCycleConfigurationDurationUnit) IsKnown() bool {
	switch r {
	case PlanExternalPlanIDVersionNewParamsReplacePricesPriceNewPlanTierWithProrationPriceBillingCycleConfigurationDurationUnitDay, PlanExternalPlanIDVersionNewParamsReplacePricesPriceNewPlanTierWithProrationPriceBillingCycleConfigurationDurationUnitMonth:
		return true
	}
	return false
}

// For dimensional price: specifies a price group and dimension values
type PlanExternalPlanIDVersionNewParamsReplacePricesPriceNewPlanTierWithProrationPriceDimensionalPriceConfiguration struct {
	// The list of dimension values matching (in order) the dimensions of the price
	// group
	DimensionValues param.Field[[]string] `json:"dimension_values,required"`
	// The id of the dimensional price group to include this price in
	DimensionalPriceGroupID param.Field[string] `json:"dimensional_price_group_id"`
	// The external id of the dimensional price group to include this price in
	ExternalDimensionalPriceGroupID param.Field[string] `json:"external_dimensional_price_group_id"`
}

func (r PlanExternalPlanIDVersionNewParamsReplacePricesPriceNewPlanTierWithProrationPriceDimensionalPriceConfiguration) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Within each billing cycle, specifies the cadence at which invoices are produced.
// If unspecified, a single invoice is produced per billing cycle.
type PlanExternalPlanIDVersionNewParamsReplacePricesPriceNewPlanTierWithProrationPriceInvoicingCycleConfiguration struct {
	// The duration of the billing period.
	Duration param.Field[int64] `json:"duration,required"`
	// The unit of billing period duration.
	DurationUnit param.Field[PlanExternalPlanIDVersionNewParamsReplacePricesPriceNewPlanTierWithProrationPriceInvoicingCycleConfigurationDurationUnit] `json:"duration_unit,required"`
}

func (r PlanExternalPlanIDVersionNewParamsReplacePricesPriceNewPlanTierWithProrationPriceInvoicingCycleConfiguration) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// The unit of billing period duration.
type PlanExternalPlanIDVersionNewParamsReplacePricesPriceNewPlanTierWithProrationPriceInvoicingCycleConfigurationDurationUnit string

const (
	PlanExternalPlanIDVersionNewParamsReplacePricesPriceNewPlanTierWithProrationPriceInvoicingCycleConfigurationDurationUnitDay   PlanExternalPlanIDVersionNewParamsReplacePricesPriceNewPlanTierWithProrationPriceInvoicingCycleConfigurationDurationUnit = "day"
	PlanExternalPlanIDVersionNewParamsReplacePricesPriceNewPlanTierWithProrationPriceInvoicingCycleConfigurationDurationUnitMonth PlanExternalPlanIDVersionNewParamsReplacePricesPriceNewPlanTierWithProrationPriceInvoicingCycleConfigurationDurationUnit = "month"
)

func (r PlanExternalPlanIDVersionNewParamsReplacePricesPriceNewPlanTierWithProrationPriceInvoicingCycleConfigurationDurationUnit) IsKnown() bool {
	switch r {
	case PlanExternalPlanIDVersionNewParamsReplacePricesPriceNewPlanTierWithProrationPriceInvoicingCycleConfigurationDurationUnitDay, PlanExternalPlanIDVersionNewParamsReplacePricesPriceNewPlanTierWithProrationPriceInvoicingCycleConfigurationDurationUnitMonth:
		return true
	}
	return false
}

type PlanExternalPlanIDVersionNewParamsReplacePricesPriceNewPlanUnitWithProrationPrice struct {
	// The cadence to bill for this price on.
	Cadence param.Field[PlanExternalPlanIDVersionNewParamsReplacePricesPriceNewPlanUnitWithProrationPriceCadence] `json:"cadence,required"`
	// The id of the item the price will be associated with.
	ItemID    param.Field[string]                                                                                     `json:"item_id,required"`
	ModelType param.Field[PlanExternalPlanIDVersionNewParamsReplacePricesPriceNewPlanUnitWithProrationPriceModelType] `json:"model_type,required"`
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
	BillingCycleConfiguration param.Field[PlanExternalPlanIDVersionNewParamsReplacePricesPriceNewPlanUnitWithProrationPriceBillingCycleConfiguration] `json:"billing_cycle_configuration"`
	// The per unit conversion rate of the price currency to the invoicing currency.
	ConversionRate param.Field[float64] `json:"conversion_rate"`
	// An ISO 4217 currency string, or custom pricing unit identifier, in which this
	// price is billed.
	Currency param.Field[string] `json:"currency"`
	// For dimensional price: specifies a price group and dimension values
	DimensionalPriceConfiguration param.Field[PlanExternalPlanIDVersionNewParamsReplacePricesPriceNewPlanUnitWithProrationPriceDimensionalPriceConfiguration] `json:"dimensional_price_configuration"`
	// An alias for the price.
	ExternalPriceID param.Field[string] `json:"external_price_id"`
	// If the Price represents a fixed cost, this represents the quantity of units
	// applied.
	FixedPriceQuantity param.Field[float64] `json:"fixed_price_quantity"`
	// The property used to group this price on an invoice
	InvoiceGroupingKey param.Field[string] `json:"invoice_grouping_key"`
	// Within each billing cycle, specifies the cadence at which invoices are produced.
	// If unspecified, a single invoice is produced per billing cycle.
	InvoicingCycleConfiguration param.Field[PlanExternalPlanIDVersionNewParamsReplacePricesPriceNewPlanUnitWithProrationPriceInvoicingCycleConfiguration] `json:"invoicing_cycle_configuration"`
	// User-specified key/value pairs for the resource. Individual keys can be removed
	// by setting the value to `null`, and the entire metadata mapping can be cleared
	// by setting `metadata` to `null`.
	Metadata param.Field[map[string]string] `json:"metadata"`
}

func (r PlanExternalPlanIDVersionNewParamsReplacePricesPriceNewPlanUnitWithProrationPrice) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r PlanExternalPlanIDVersionNewParamsReplacePricesPriceNewPlanUnitWithProrationPrice) implementsPlanExternalPlanIDVersionNewParamsReplacePricesPriceUnion() {
}

// The cadence to bill for this price on.
type PlanExternalPlanIDVersionNewParamsReplacePricesPriceNewPlanUnitWithProrationPriceCadence string

const (
	PlanExternalPlanIDVersionNewParamsReplacePricesPriceNewPlanUnitWithProrationPriceCadenceAnnual     PlanExternalPlanIDVersionNewParamsReplacePricesPriceNewPlanUnitWithProrationPriceCadence = "annual"
	PlanExternalPlanIDVersionNewParamsReplacePricesPriceNewPlanUnitWithProrationPriceCadenceSemiAnnual PlanExternalPlanIDVersionNewParamsReplacePricesPriceNewPlanUnitWithProrationPriceCadence = "semi_annual"
	PlanExternalPlanIDVersionNewParamsReplacePricesPriceNewPlanUnitWithProrationPriceCadenceMonthly    PlanExternalPlanIDVersionNewParamsReplacePricesPriceNewPlanUnitWithProrationPriceCadence = "monthly"
	PlanExternalPlanIDVersionNewParamsReplacePricesPriceNewPlanUnitWithProrationPriceCadenceQuarterly  PlanExternalPlanIDVersionNewParamsReplacePricesPriceNewPlanUnitWithProrationPriceCadence = "quarterly"
	PlanExternalPlanIDVersionNewParamsReplacePricesPriceNewPlanUnitWithProrationPriceCadenceOneTime    PlanExternalPlanIDVersionNewParamsReplacePricesPriceNewPlanUnitWithProrationPriceCadence = "one_time"
	PlanExternalPlanIDVersionNewParamsReplacePricesPriceNewPlanUnitWithProrationPriceCadenceCustom     PlanExternalPlanIDVersionNewParamsReplacePricesPriceNewPlanUnitWithProrationPriceCadence = "custom"
)

func (r PlanExternalPlanIDVersionNewParamsReplacePricesPriceNewPlanUnitWithProrationPriceCadence) IsKnown() bool {
	switch r {
	case PlanExternalPlanIDVersionNewParamsReplacePricesPriceNewPlanUnitWithProrationPriceCadenceAnnual, PlanExternalPlanIDVersionNewParamsReplacePricesPriceNewPlanUnitWithProrationPriceCadenceSemiAnnual, PlanExternalPlanIDVersionNewParamsReplacePricesPriceNewPlanUnitWithProrationPriceCadenceMonthly, PlanExternalPlanIDVersionNewParamsReplacePricesPriceNewPlanUnitWithProrationPriceCadenceQuarterly, PlanExternalPlanIDVersionNewParamsReplacePricesPriceNewPlanUnitWithProrationPriceCadenceOneTime, PlanExternalPlanIDVersionNewParamsReplacePricesPriceNewPlanUnitWithProrationPriceCadenceCustom:
		return true
	}
	return false
}

type PlanExternalPlanIDVersionNewParamsReplacePricesPriceNewPlanUnitWithProrationPriceModelType string

const (
	PlanExternalPlanIDVersionNewParamsReplacePricesPriceNewPlanUnitWithProrationPriceModelTypeUnitWithProration PlanExternalPlanIDVersionNewParamsReplacePricesPriceNewPlanUnitWithProrationPriceModelType = "unit_with_proration"
)

func (r PlanExternalPlanIDVersionNewParamsReplacePricesPriceNewPlanUnitWithProrationPriceModelType) IsKnown() bool {
	switch r {
	case PlanExternalPlanIDVersionNewParamsReplacePricesPriceNewPlanUnitWithProrationPriceModelTypeUnitWithProration:
		return true
	}
	return false
}

// For custom cadence: specifies the duration of the billing period in days or
// months.
type PlanExternalPlanIDVersionNewParamsReplacePricesPriceNewPlanUnitWithProrationPriceBillingCycleConfiguration struct {
	// The duration of the billing period.
	Duration param.Field[int64] `json:"duration,required"`
	// The unit of billing period duration.
	DurationUnit param.Field[PlanExternalPlanIDVersionNewParamsReplacePricesPriceNewPlanUnitWithProrationPriceBillingCycleConfigurationDurationUnit] `json:"duration_unit,required"`
}

func (r PlanExternalPlanIDVersionNewParamsReplacePricesPriceNewPlanUnitWithProrationPriceBillingCycleConfiguration) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// The unit of billing period duration.
type PlanExternalPlanIDVersionNewParamsReplacePricesPriceNewPlanUnitWithProrationPriceBillingCycleConfigurationDurationUnit string

const (
	PlanExternalPlanIDVersionNewParamsReplacePricesPriceNewPlanUnitWithProrationPriceBillingCycleConfigurationDurationUnitDay   PlanExternalPlanIDVersionNewParamsReplacePricesPriceNewPlanUnitWithProrationPriceBillingCycleConfigurationDurationUnit = "day"
	PlanExternalPlanIDVersionNewParamsReplacePricesPriceNewPlanUnitWithProrationPriceBillingCycleConfigurationDurationUnitMonth PlanExternalPlanIDVersionNewParamsReplacePricesPriceNewPlanUnitWithProrationPriceBillingCycleConfigurationDurationUnit = "month"
)

func (r PlanExternalPlanIDVersionNewParamsReplacePricesPriceNewPlanUnitWithProrationPriceBillingCycleConfigurationDurationUnit) IsKnown() bool {
	switch r {
	case PlanExternalPlanIDVersionNewParamsReplacePricesPriceNewPlanUnitWithProrationPriceBillingCycleConfigurationDurationUnitDay, PlanExternalPlanIDVersionNewParamsReplacePricesPriceNewPlanUnitWithProrationPriceBillingCycleConfigurationDurationUnitMonth:
		return true
	}
	return false
}

// For dimensional price: specifies a price group and dimension values
type PlanExternalPlanIDVersionNewParamsReplacePricesPriceNewPlanUnitWithProrationPriceDimensionalPriceConfiguration struct {
	// The list of dimension values matching (in order) the dimensions of the price
	// group
	DimensionValues param.Field[[]string] `json:"dimension_values,required"`
	// The id of the dimensional price group to include this price in
	DimensionalPriceGroupID param.Field[string] `json:"dimensional_price_group_id"`
	// The external id of the dimensional price group to include this price in
	ExternalDimensionalPriceGroupID param.Field[string] `json:"external_dimensional_price_group_id"`
}

func (r PlanExternalPlanIDVersionNewParamsReplacePricesPriceNewPlanUnitWithProrationPriceDimensionalPriceConfiguration) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Within each billing cycle, specifies the cadence at which invoices are produced.
// If unspecified, a single invoice is produced per billing cycle.
type PlanExternalPlanIDVersionNewParamsReplacePricesPriceNewPlanUnitWithProrationPriceInvoicingCycleConfiguration struct {
	// The duration of the billing period.
	Duration param.Field[int64] `json:"duration,required"`
	// The unit of billing period duration.
	DurationUnit param.Field[PlanExternalPlanIDVersionNewParamsReplacePricesPriceNewPlanUnitWithProrationPriceInvoicingCycleConfigurationDurationUnit] `json:"duration_unit,required"`
}

func (r PlanExternalPlanIDVersionNewParamsReplacePricesPriceNewPlanUnitWithProrationPriceInvoicingCycleConfiguration) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// The unit of billing period duration.
type PlanExternalPlanIDVersionNewParamsReplacePricesPriceNewPlanUnitWithProrationPriceInvoicingCycleConfigurationDurationUnit string

const (
	PlanExternalPlanIDVersionNewParamsReplacePricesPriceNewPlanUnitWithProrationPriceInvoicingCycleConfigurationDurationUnitDay   PlanExternalPlanIDVersionNewParamsReplacePricesPriceNewPlanUnitWithProrationPriceInvoicingCycleConfigurationDurationUnit = "day"
	PlanExternalPlanIDVersionNewParamsReplacePricesPriceNewPlanUnitWithProrationPriceInvoicingCycleConfigurationDurationUnitMonth PlanExternalPlanIDVersionNewParamsReplacePricesPriceNewPlanUnitWithProrationPriceInvoicingCycleConfigurationDurationUnit = "month"
)

func (r PlanExternalPlanIDVersionNewParamsReplacePricesPriceNewPlanUnitWithProrationPriceInvoicingCycleConfigurationDurationUnit) IsKnown() bool {
	switch r {
	case PlanExternalPlanIDVersionNewParamsReplacePricesPriceNewPlanUnitWithProrationPriceInvoicingCycleConfigurationDurationUnitDay, PlanExternalPlanIDVersionNewParamsReplacePricesPriceNewPlanUnitWithProrationPriceInvoicingCycleConfigurationDurationUnitMonth:
		return true
	}
	return false
}

type PlanExternalPlanIDVersionNewParamsReplacePricesPriceNewPlanGroupedAllocationPrice struct {
	// The cadence to bill for this price on.
	Cadence                 param.Field[PlanExternalPlanIDVersionNewParamsReplacePricesPriceNewPlanGroupedAllocationPriceCadence] `json:"cadence,required"`
	GroupedAllocationConfig param.Field[map[string]interface{}]                                                                   `json:"grouped_allocation_config,required"`
	// The id of the item the price will be associated with.
	ItemID    param.Field[string]                                                                                     `json:"item_id,required"`
	ModelType param.Field[PlanExternalPlanIDVersionNewParamsReplacePricesPriceNewPlanGroupedAllocationPriceModelType] `json:"model_type,required"`
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
	BillingCycleConfiguration param.Field[PlanExternalPlanIDVersionNewParamsReplacePricesPriceNewPlanGroupedAllocationPriceBillingCycleConfiguration] `json:"billing_cycle_configuration"`
	// The per unit conversion rate of the price currency to the invoicing currency.
	ConversionRate param.Field[float64] `json:"conversion_rate"`
	// An ISO 4217 currency string, or custom pricing unit identifier, in which this
	// price is billed.
	Currency param.Field[string] `json:"currency"`
	// For dimensional price: specifies a price group and dimension values
	DimensionalPriceConfiguration param.Field[PlanExternalPlanIDVersionNewParamsReplacePricesPriceNewPlanGroupedAllocationPriceDimensionalPriceConfiguration] `json:"dimensional_price_configuration"`
	// An alias for the price.
	ExternalPriceID param.Field[string] `json:"external_price_id"`
	// If the Price represents a fixed cost, this represents the quantity of units
	// applied.
	FixedPriceQuantity param.Field[float64] `json:"fixed_price_quantity"`
	// The property used to group this price on an invoice
	InvoiceGroupingKey param.Field[string] `json:"invoice_grouping_key"`
	// Within each billing cycle, specifies the cadence at which invoices are produced.
	// If unspecified, a single invoice is produced per billing cycle.
	InvoicingCycleConfiguration param.Field[PlanExternalPlanIDVersionNewParamsReplacePricesPriceNewPlanGroupedAllocationPriceInvoicingCycleConfiguration] `json:"invoicing_cycle_configuration"`
	// User-specified key/value pairs for the resource. Individual keys can be removed
	// by setting the value to `null`, and the entire metadata mapping can be cleared
	// by setting `metadata` to `null`.
	Metadata param.Field[map[string]string] `json:"metadata"`
}

func (r PlanExternalPlanIDVersionNewParamsReplacePricesPriceNewPlanGroupedAllocationPrice) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r PlanExternalPlanIDVersionNewParamsReplacePricesPriceNewPlanGroupedAllocationPrice) implementsPlanExternalPlanIDVersionNewParamsReplacePricesPriceUnion() {
}

// The cadence to bill for this price on.
type PlanExternalPlanIDVersionNewParamsReplacePricesPriceNewPlanGroupedAllocationPriceCadence string

const (
	PlanExternalPlanIDVersionNewParamsReplacePricesPriceNewPlanGroupedAllocationPriceCadenceAnnual     PlanExternalPlanIDVersionNewParamsReplacePricesPriceNewPlanGroupedAllocationPriceCadence = "annual"
	PlanExternalPlanIDVersionNewParamsReplacePricesPriceNewPlanGroupedAllocationPriceCadenceSemiAnnual PlanExternalPlanIDVersionNewParamsReplacePricesPriceNewPlanGroupedAllocationPriceCadence = "semi_annual"
	PlanExternalPlanIDVersionNewParamsReplacePricesPriceNewPlanGroupedAllocationPriceCadenceMonthly    PlanExternalPlanIDVersionNewParamsReplacePricesPriceNewPlanGroupedAllocationPriceCadence = "monthly"
	PlanExternalPlanIDVersionNewParamsReplacePricesPriceNewPlanGroupedAllocationPriceCadenceQuarterly  PlanExternalPlanIDVersionNewParamsReplacePricesPriceNewPlanGroupedAllocationPriceCadence = "quarterly"
	PlanExternalPlanIDVersionNewParamsReplacePricesPriceNewPlanGroupedAllocationPriceCadenceOneTime    PlanExternalPlanIDVersionNewParamsReplacePricesPriceNewPlanGroupedAllocationPriceCadence = "one_time"
	PlanExternalPlanIDVersionNewParamsReplacePricesPriceNewPlanGroupedAllocationPriceCadenceCustom     PlanExternalPlanIDVersionNewParamsReplacePricesPriceNewPlanGroupedAllocationPriceCadence = "custom"
)

func (r PlanExternalPlanIDVersionNewParamsReplacePricesPriceNewPlanGroupedAllocationPriceCadence) IsKnown() bool {
	switch r {
	case PlanExternalPlanIDVersionNewParamsReplacePricesPriceNewPlanGroupedAllocationPriceCadenceAnnual, PlanExternalPlanIDVersionNewParamsReplacePricesPriceNewPlanGroupedAllocationPriceCadenceSemiAnnual, PlanExternalPlanIDVersionNewParamsReplacePricesPriceNewPlanGroupedAllocationPriceCadenceMonthly, PlanExternalPlanIDVersionNewParamsReplacePricesPriceNewPlanGroupedAllocationPriceCadenceQuarterly, PlanExternalPlanIDVersionNewParamsReplacePricesPriceNewPlanGroupedAllocationPriceCadenceOneTime, PlanExternalPlanIDVersionNewParamsReplacePricesPriceNewPlanGroupedAllocationPriceCadenceCustom:
		return true
	}
	return false
}

type PlanExternalPlanIDVersionNewParamsReplacePricesPriceNewPlanGroupedAllocationPriceModelType string

const (
	PlanExternalPlanIDVersionNewParamsReplacePricesPriceNewPlanGroupedAllocationPriceModelTypeGroupedAllocation PlanExternalPlanIDVersionNewParamsReplacePricesPriceNewPlanGroupedAllocationPriceModelType = "grouped_allocation"
)

func (r PlanExternalPlanIDVersionNewParamsReplacePricesPriceNewPlanGroupedAllocationPriceModelType) IsKnown() bool {
	switch r {
	case PlanExternalPlanIDVersionNewParamsReplacePricesPriceNewPlanGroupedAllocationPriceModelTypeGroupedAllocation:
		return true
	}
	return false
}

// For custom cadence: specifies the duration of the billing period in days or
// months.
type PlanExternalPlanIDVersionNewParamsReplacePricesPriceNewPlanGroupedAllocationPriceBillingCycleConfiguration struct {
	// The duration of the billing period.
	Duration param.Field[int64] `json:"duration,required"`
	// The unit of billing period duration.
	DurationUnit param.Field[PlanExternalPlanIDVersionNewParamsReplacePricesPriceNewPlanGroupedAllocationPriceBillingCycleConfigurationDurationUnit] `json:"duration_unit,required"`
}

func (r PlanExternalPlanIDVersionNewParamsReplacePricesPriceNewPlanGroupedAllocationPriceBillingCycleConfiguration) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// The unit of billing period duration.
type PlanExternalPlanIDVersionNewParamsReplacePricesPriceNewPlanGroupedAllocationPriceBillingCycleConfigurationDurationUnit string

const (
	PlanExternalPlanIDVersionNewParamsReplacePricesPriceNewPlanGroupedAllocationPriceBillingCycleConfigurationDurationUnitDay   PlanExternalPlanIDVersionNewParamsReplacePricesPriceNewPlanGroupedAllocationPriceBillingCycleConfigurationDurationUnit = "day"
	PlanExternalPlanIDVersionNewParamsReplacePricesPriceNewPlanGroupedAllocationPriceBillingCycleConfigurationDurationUnitMonth PlanExternalPlanIDVersionNewParamsReplacePricesPriceNewPlanGroupedAllocationPriceBillingCycleConfigurationDurationUnit = "month"
)

func (r PlanExternalPlanIDVersionNewParamsReplacePricesPriceNewPlanGroupedAllocationPriceBillingCycleConfigurationDurationUnit) IsKnown() bool {
	switch r {
	case PlanExternalPlanIDVersionNewParamsReplacePricesPriceNewPlanGroupedAllocationPriceBillingCycleConfigurationDurationUnitDay, PlanExternalPlanIDVersionNewParamsReplacePricesPriceNewPlanGroupedAllocationPriceBillingCycleConfigurationDurationUnitMonth:
		return true
	}
	return false
}

// For dimensional price: specifies a price group and dimension values
type PlanExternalPlanIDVersionNewParamsReplacePricesPriceNewPlanGroupedAllocationPriceDimensionalPriceConfiguration struct {
	// The list of dimension values matching (in order) the dimensions of the price
	// group
	DimensionValues param.Field[[]string] `json:"dimension_values,required"`
	// The id of the dimensional price group to include this price in
	DimensionalPriceGroupID param.Field[string] `json:"dimensional_price_group_id"`
	// The external id of the dimensional price group to include this price in
	ExternalDimensionalPriceGroupID param.Field[string] `json:"external_dimensional_price_group_id"`
}

func (r PlanExternalPlanIDVersionNewParamsReplacePricesPriceNewPlanGroupedAllocationPriceDimensionalPriceConfiguration) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Within each billing cycle, specifies the cadence at which invoices are produced.
// If unspecified, a single invoice is produced per billing cycle.
type PlanExternalPlanIDVersionNewParamsReplacePricesPriceNewPlanGroupedAllocationPriceInvoicingCycleConfiguration struct {
	// The duration of the billing period.
	Duration param.Field[int64] `json:"duration,required"`
	// The unit of billing period duration.
	DurationUnit param.Field[PlanExternalPlanIDVersionNewParamsReplacePricesPriceNewPlanGroupedAllocationPriceInvoicingCycleConfigurationDurationUnit] `json:"duration_unit,required"`
}

func (r PlanExternalPlanIDVersionNewParamsReplacePricesPriceNewPlanGroupedAllocationPriceInvoicingCycleConfiguration) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// The unit of billing period duration.
type PlanExternalPlanIDVersionNewParamsReplacePricesPriceNewPlanGroupedAllocationPriceInvoicingCycleConfigurationDurationUnit string

const (
	PlanExternalPlanIDVersionNewParamsReplacePricesPriceNewPlanGroupedAllocationPriceInvoicingCycleConfigurationDurationUnitDay   PlanExternalPlanIDVersionNewParamsReplacePricesPriceNewPlanGroupedAllocationPriceInvoicingCycleConfigurationDurationUnit = "day"
	PlanExternalPlanIDVersionNewParamsReplacePricesPriceNewPlanGroupedAllocationPriceInvoicingCycleConfigurationDurationUnitMonth PlanExternalPlanIDVersionNewParamsReplacePricesPriceNewPlanGroupedAllocationPriceInvoicingCycleConfigurationDurationUnit = "month"
)

func (r PlanExternalPlanIDVersionNewParamsReplacePricesPriceNewPlanGroupedAllocationPriceInvoicingCycleConfigurationDurationUnit) IsKnown() bool {
	switch r {
	case PlanExternalPlanIDVersionNewParamsReplacePricesPriceNewPlanGroupedAllocationPriceInvoicingCycleConfigurationDurationUnitDay, PlanExternalPlanIDVersionNewParamsReplacePricesPriceNewPlanGroupedAllocationPriceInvoicingCycleConfigurationDurationUnitMonth:
		return true
	}
	return false
}

type PlanExternalPlanIDVersionNewParamsReplacePricesPriceNewPlanGroupedWithProratedMinimumPrice struct {
	// The cadence to bill for this price on.
	Cadence                          param.Field[PlanExternalPlanIDVersionNewParamsReplacePricesPriceNewPlanGroupedWithProratedMinimumPriceCadence] `json:"cadence,required"`
	GroupedWithProratedMinimumConfig param.Field[map[string]interface{}]                                                                            `json:"grouped_with_prorated_minimum_config,required"`
	// The id of the item the price will be associated with.
	ItemID    param.Field[string]                                                                                              `json:"item_id,required"`
	ModelType param.Field[PlanExternalPlanIDVersionNewParamsReplacePricesPriceNewPlanGroupedWithProratedMinimumPriceModelType] `json:"model_type,required"`
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
	BillingCycleConfiguration param.Field[PlanExternalPlanIDVersionNewParamsReplacePricesPriceNewPlanGroupedWithProratedMinimumPriceBillingCycleConfiguration] `json:"billing_cycle_configuration"`
	// The per unit conversion rate of the price currency to the invoicing currency.
	ConversionRate param.Field[float64] `json:"conversion_rate"`
	// An ISO 4217 currency string, or custom pricing unit identifier, in which this
	// price is billed.
	Currency param.Field[string] `json:"currency"`
	// For dimensional price: specifies a price group and dimension values
	DimensionalPriceConfiguration param.Field[PlanExternalPlanIDVersionNewParamsReplacePricesPriceNewPlanGroupedWithProratedMinimumPriceDimensionalPriceConfiguration] `json:"dimensional_price_configuration"`
	// An alias for the price.
	ExternalPriceID param.Field[string] `json:"external_price_id"`
	// If the Price represents a fixed cost, this represents the quantity of units
	// applied.
	FixedPriceQuantity param.Field[float64] `json:"fixed_price_quantity"`
	// The property used to group this price on an invoice
	InvoiceGroupingKey param.Field[string] `json:"invoice_grouping_key"`
	// Within each billing cycle, specifies the cadence at which invoices are produced.
	// If unspecified, a single invoice is produced per billing cycle.
	InvoicingCycleConfiguration param.Field[PlanExternalPlanIDVersionNewParamsReplacePricesPriceNewPlanGroupedWithProratedMinimumPriceInvoicingCycleConfiguration] `json:"invoicing_cycle_configuration"`
	// User-specified key/value pairs for the resource. Individual keys can be removed
	// by setting the value to `null`, and the entire metadata mapping can be cleared
	// by setting `metadata` to `null`.
	Metadata param.Field[map[string]string] `json:"metadata"`
}

func (r PlanExternalPlanIDVersionNewParamsReplacePricesPriceNewPlanGroupedWithProratedMinimumPrice) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r PlanExternalPlanIDVersionNewParamsReplacePricesPriceNewPlanGroupedWithProratedMinimumPrice) implementsPlanExternalPlanIDVersionNewParamsReplacePricesPriceUnion() {
}

// The cadence to bill for this price on.
type PlanExternalPlanIDVersionNewParamsReplacePricesPriceNewPlanGroupedWithProratedMinimumPriceCadence string

const (
	PlanExternalPlanIDVersionNewParamsReplacePricesPriceNewPlanGroupedWithProratedMinimumPriceCadenceAnnual     PlanExternalPlanIDVersionNewParamsReplacePricesPriceNewPlanGroupedWithProratedMinimumPriceCadence = "annual"
	PlanExternalPlanIDVersionNewParamsReplacePricesPriceNewPlanGroupedWithProratedMinimumPriceCadenceSemiAnnual PlanExternalPlanIDVersionNewParamsReplacePricesPriceNewPlanGroupedWithProratedMinimumPriceCadence = "semi_annual"
	PlanExternalPlanIDVersionNewParamsReplacePricesPriceNewPlanGroupedWithProratedMinimumPriceCadenceMonthly    PlanExternalPlanIDVersionNewParamsReplacePricesPriceNewPlanGroupedWithProratedMinimumPriceCadence = "monthly"
	PlanExternalPlanIDVersionNewParamsReplacePricesPriceNewPlanGroupedWithProratedMinimumPriceCadenceQuarterly  PlanExternalPlanIDVersionNewParamsReplacePricesPriceNewPlanGroupedWithProratedMinimumPriceCadence = "quarterly"
	PlanExternalPlanIDVersionNewParamsReplacePricesPriceNewPlanGroupedWithProratedMinimumPriceCadenceOneTime    PlanExternalPlanIDVersionNewParamsReplacePricesPriceNewPlanGroupedWithProratedMinimumPriceCadence = "one_time"
	PlanExternalPlanIDVersionNewParamsReplacePricesPriceNewPlanGroupedWithProratedMinimumPriceCadenceCustom     PlanExternalPlanIDVersionNewParamsReplacePricesPriceNewPlanGroupedWithProratedMinimumPriceCadence = "custom"
)

func (r PlanExternalPlanIDVersionNewParamsReplacePricesPriceNewPlanGroupedWithProratedMinimumPriceCadence) IsKnown() bool {
	switch r {
	case PlanExternalPlanIDVersionNewParamsReplacePricesPriceNewPlanGroupedWithProratedMinimumPriceCadenceAnnual, PlanExternalPlanIDVersionNewParamsReplacePricesPriceNewPlanGroupedWithProratedMinimumPriceCadenceSemiAnnual, PlanExternalPlanIDVersionNewParamsReplacePricesPriceNewPlanGroupedWithProratedMinimumPriceCadenceMonthly, PlanExternalPlanIDVersionNewParamsReplacePricesPriceNewPlanGroupedWithProratedMinimumPriceCadenceQuarterly, PlanExternalPlanIDVersionNewParamsReplacePricesPriceNewPlanGroupedWithProratedMinimumPriceCadenceOneTime, PlanExternalPlanIDVersionNewParamsReplacePricesPriceNewPlanGroupedWithProratedMinimumPriceCadenceCustom:
		return true
	}
	return false
}

type PlanExternalPlanIDVersionNewParamsReplacePricesPriceNewPlanGroupedWithProratedMinimumPriceModelType string

const (
	PlanExternalPlanIDVersionNewParamsReplacePricesPriceNewPlanGroupedWithProratedMinimumPriceModelTypeGroupedWithProratedMinimum PlanExternalPlanIDVersionNewParamsReplacePricesPriceNewPlanGroupedWithProratedMinimumPriceModelType = "grouped_with_prorated_minimum"
)

func (r PlanExternalPlanIDVersionNewParamsReplacePricesPriceNewPlanGroupedWithProratedMinimumPriceModelType) IsKnown() bool {
	switch r {
	case PlanExternalPlanIDVersionNewParamsReplacePricesPriceNewPlanGroupedWithProratedMinimumPriceModelTypeGroupedWithProratedMinimum:
		return true
	}
	return false
}

// For custom cadence: specifies the duration of the billing period in days or
// months.
type PlanExternalPlanIDVersionNewParamsReplacePricesPriceNewPlanGroupedWithProratedMinimumPriceBillingCycleConfiguration struct {
	// The duration of the billing period.
	Duration param.Field[int64] `json:"duration,required"`
	// The unit of billing period duration.
	DurationUnit param.Field[PlanExternalPlanIDVersionNewParamsReplacePricesPriceNewPlanGroupedWithProratedMinimumPriceBillingCycleConfigurationDurationUnit] `json:"duration_unit,required"`
}

func (r PlanExternalPlanIDVersionNewParamsReplacePricesPriceNewPlanGroupedWithProratedMinimumPriceBillingCycleConfiguration) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// The unit of billing period duration.
type PlanExternalPlanIDVersionNewParamsReplacePricesPriceNewPlanGroupedWithProratedMinimumPriceBillingCycleConfigurationDurationUnit string

const (
	PlanExternalPlanIDVersionNewParamsReplacePricesPriceNewPlanGroupedWithProratedMinimumPriceBillingCycleConfigurationDurationUnitDay   PlanExternalPlanIDVersionNewParamsReplacePricesPriceNewPlanGroupedWithProratedMinimumPriceBillingCycleConfigurationDurationUnit = "day"
	PlanExternalPlanIDVersionNewParamsReplacePricesPriceNewPlanGroupedWithProratedMinimumPriceBillingCycleConfigurationDurationUnitMonth PlanExternalPlanIDVersionNewParamsReplacePricesPriceNewPlanGroupedWithProratedMinimumPriceBillingCycleConfigurationDurationUnit = "month"
)

func (r PlanExternalPlanIDVersionNewParamsReplacePricesPriceNewPlanGroupedWithProratedMinimumPriceBillingCycleConfigurationDurationUnit) IsKnown() bool {
	switch r {
	case PlanExternalPlanIDVersionNewParamsReplacePricesPriceNewPlanGroupedWithProratedMinimumPriceBillingCycleConfigurationDurationUnitDay, PlanExternalPlanIDVersionNewParamsReplacePricesPriceNewPlanGroupedWithProratedMinimumPriceBillingCycleConfigurationDurationUnitMonth:
		return true
	}
	return false
}

// For dimensional price: specifies a price group and dimension values
type PlanExternalPlanIDVersionNewParamsReplacePricesPriceNewPlanGroupedWithProratedMinimumPriceDimensionalPriceConfiguration struct {
	// The list of dimension values matching (in order) the dimensions of the price
	// group
	DimensionValues param.Field[[]string] `json:"dimension_values,required"`
	// The id of the dimensional price group to include this price in
	DimensionalPriceGroupID param.Field[string] `json:"dimensional_price_group_id"`
	// The external id of the dimensional price group to include this price in
	ExternalDimensionalPriceGroupID param.Field[string] `json:"external_dimensional_price_group_id"`
}

func (r PlanExternalPlanIDVersionNewParamsReplacePricesPriceNewPlanGroupedWithProratedMinimumPriceDimensionalPriceConfiguration) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Within each billing cycle, specifies the cadence at which invoices are produced.
// If unspecified, a single invoice is produced per billing cycle.
type PlanExternalPlanIDVersionNewParamsReplacePricesPriceNewPlanGroupedWithProratedMinimumPriceInvoicingCycleConfiguration struct {
	// The duration of the billing period.
	Duration param.Field[int64] `json:"duration,required"`
	// The unit of billing period duration.
	DurationUnit param.Field[PlanExternalPlanIDVersionNewParamsReplacePricesPriceNewPlanGroupedWithProratedMinimumPriceInvoicingCycleConfigurationDurationUnit] `json:"duration_unit,required"`
}

func (r PlanExternalPlanIDVersionNewParamsReplacePricesPriceNewPlanGroupedWithProratedMinimumPriceInvoicingCycleConfiguration) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// The unit of billing period duration.
type PlanExternalPlanIDVersionNewParamsReplacePricesPriceNewPlanGroupedWithProratedMinimumPriceInvoicingCycleConfigurationDurationUnit string

const (
	PlanExternalPlanIDVersionNewParamsReplacePricesPriceNewPlanGroupedWithProratedMinimumPriceInvoicingCycleConfigurationDurationUnitDay   PlanExternalPlanIDVersionNewParamsReplacePricesPriceNewPlanGroupedWithProratedMinimumPriceInvoicingCycleConfigurationDurationUnit = "day"
	PlanExternalPlanIDVersionNewParamsReplacePricesPriceNewPlanGroupedWithProratedMinimumPriceInvoicingCycleConfigurationDurationUnitMonth PlanExternalPlanIDVersionNewParamsReplacePricesPriceNewPlanGroupedWithProratedMinimumPriceInvoicingCycleConfigurationDurationUnit = "month"
)

func (r PlanExternalPlanIDVersionNewParamsReplacePricesPriceNewPlanGroupedWithProratedMinimumPriceInvoicingCycleConfigurationDurationUnit) IsKnown() bool {
	switch r {
	case PlanExternalPlanIDVersionNewParamsReplacePricesPriceNewPlanGroupedWithProratedMinimumPriceInvoicingCycleConfigurationDurationUnitDay, PlanExternalPlanIDVersionNewParamsReplacePricesPriceNewPlanGroupedWithProratedMinimumPriceInvoicingCycleConfigurationDurationUnitMonth:
		return true
	}
	return false
}

type PlanExternalPlanIDVersionNewParamsReplacePricesPriceNewPlanGroupedWithMeteredMinimumPrice struct {
	// The cadence to bill for this price on.
	Cadence                         param.Field[PlanExternalPlanIDVersionNewParamsReplacePricesPriceNewPlanGroupedWithMeteredMinimumPriceCadence] `json:"cadence,required"`
	GroupedWithMeteredMinimumConfig param.Field[map[string]interface{}]                                                                           `json:"grouped_with_metered_minimum_config,required"`
	// The id of the item the price will be associated with.
	ItemID    param.Field[string]                                                                                             `json:"item_id,required"`
	ModelType param.Field[PlanExternalPlanIDVersionNewParamsReplacePricesPriceNewPlanGroupedWithMeteredMinimumPriceModelType] `json:"model_type,required"`
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
	BillingCycleConfiguration param.Field[PlanExternalPlanIDVersionNewParamsReplacePricesPriceNewPlanGroupedWithMeteredMinimumPriceBillingCycleConfiguration] `json:"billing_cycle_configuration"`
	// The per unit conversion rate of the price currency to the invoicing currency.
	ConversionRate param.Field[float64] `json:"conversion_rate"`
	// An ISO 4217 currency string, or custom pricing unit identifier, in which this
	// price is billed.
	Currency param.Field[string] `json:"currency"`
	// For dimensional price: specifies a price group and dimension values
	DimensionalPriceConfiguration param.Field[PlanExternalPlanIDVersionNewParamsReplacePricesPriceNewPlanGroupedWithMeteredMinimumPriceDimensionalPriceConfiguration] `json:"dimensional_price_configuration"`
	// An alias for the price.
	ExternalPriceID param.Field[string] `json:"external_price_id"`
	// If the Price represents a fixed cost, this represents the quantity of units
	// applied.
	FixedPriceQuantity param.Field[float64] `json:"fixed_price_quantity"`
	// The property used to group this price on an invoice
	InvoiceGroupingKey param.Field[string] `json:"invoice_grouping_key"`
	// Within each billing cycle, specifies the cadence at which invoices are produced.
	// If unspecified, a single invoice is produced per billing cycle.
	InvoicingCycleConfiguration param.Field[PlanExternalPlanIDVersionNewParamsReplacePricesPriceNewPlanGroupedWithMeteredMinimumPriceInvoicingCycleConfiguration] `json:"invoicing_cycle_configuration"`
	// User-specified key/value pairs for the resource. Individual keys can be removed
	// by setting the value to `null`, and the entire metadata mapping can be cleared
	// by setting `metadata` to `null`.
	Metadata param.Field[map[string]string] `json:"metadata"`
}

func (r PlanExternalPlanIDVersionNewParamsReplacePricesPriceNewPlanGroupedWithMeteredMinimumPrice) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r PlanExternalPlanIDVersionNewParamsReplacePricesPriceNewPlanGroupedWithMeteredMinimumPrice) implementsPlanExternalPlanIDVersionNewParamsReplacePricesPriceUnion() {
}

// The cadence to bill for this price on.
type PlanExternalPlanIDVersionNewParamsReplacePricesPriceNewPlanGroupedWithMeteredMinimumPriceCadence string

const (
	PlanExternalPlanIDVersionNewParamsReplacePricesPriceNewPlanGroupedWithMeteredMinimumPriceCadenceAnnual     PlanExternalPlanIDVersionNewParamsReplacePricesPriceNewPlanGroupedWithMeteredMinimumPriceCadence = "annual"
	PlanExternalPlanIDVersionNewParamsReplacePricesPriceNewPlanGroupedWithMeteredMinimumPriceCadenceSemiAnnual PlanExternalPlanIDVersionNewParamsReplacePricesPriceNewPlanGroupedWithMeteredMinimumPriceCadence = "semi_annual"
	PlanExternalPlanIDVersionNewParamsReplacePricesPriceNewPlanGroupedWithMeteredMinimumPriceCadenceMonthly    PlanExternalPlanIDVersionNewParamsReplacePricesPriceNewPlanGroupedWithMeteredMinimumPriceCadence = "monthly"
	PlanExternalPlanIDVersionNewParamsReplacePricesPriceNewPlanGroupedWithMeteredMinimumPriceCadenceQuarterly  PlanExternalPlanIDVersionNewParamsReplacePricesPriceNewPlanGroupedWithMeteredMinimumPriceCadence = "quarterly"
	PlanExternalPlanIDVersionNewParamsReplacePricesPriceNewPlanGroupedWithMeteredMinimumPriceCadenceOneTime    PlanExternalPlanIDVersionNewParamsReplacePricesPriceNewPlanGroupedWithMeteredMinimumPriceCadence = "one_time"
	PlanExternalPlanIDVersionNewParamsReplacePricesPriceNewPlanGroupedWithMeteredMinimumPriceCadenceCustom     PlanExternalPlanIDVersionNewParamsReplacePricesPriceNewPlanGroupedWithMeteredMinimumPriceCadence = "custom"
)

func (r PlanExternalPlanIDVersionNewParamsReplacePricesPriceNewPlanGroupedWithMeteredMinimumPriceCadence) IsKnown() bool {
	switch r {
	case PlanExternalPlanIDVersionNewParamsReplacePricesPriceNewPlanGroupedWithMeteredMinimumPriceCadenceAnnual, PlanExternalPlanIDVersionNewParamsReplacePricesPriceNewPlanGroupedWithMeteredMinimumPriceCadenceSemiAnnual, PlanExternalPlanIDVersionNewParamsReplacePricesPriceNewPlanGroupedWithMeteredMinimumPriceCadenceMonthly, PlanExternalPlanIDVersionNewParamsReplacePricesPriceNewPlanGroupedWithMeteredMinimumPriceCadenceQuarterly, PlanExternalPlanIDVersionNewParamsReplacePricesPriceNewPlanGroupedWithMeteredMinimumPriceCadenceOneTime, PlanExternalPlanIDVersionNewParamsReplacePricesPriceNewPlanGroupedWithMeteredMinimumPriceCadenceCustom:
		return true
	}
	return false
}

type PlanExternalPlanIDVersionNewParamsReplacePricesPriceNewPlanGroupedWithMeteredMinimumPriceModelType string

const (
	PlanExternalPlanIDVersionNewParamsReplacePricesPriceNewPlanGroupedWithMeteredMinimumPriceModelTypeGroupedWithMeteredMinimum PlanExternalPlanIDVersionNewParamsReplacePricesPriceNewPlanGroupedWithMeteredMinimumPriceModelType = "grouped_with_metered_minimum"
)

func (r PlanExternalPlanIDVersionNewParamsReplacePricesPriceNewPlanGroupedWithMeteredMinimumPriceModelType) IsKnown() bool {
	switch r {
	case PlanExternalPlanIDVersionNewParamsReplacePricesPriceNewPlanGroupedWithMeteredMinimumPriceModelTypeGroupedWithMeteredMinimum:
		return true
	}
	return false
}

// For custom cadence: specifies the duration of the billing period in days or
// months.
type PlanExternalPlanIDVersionNewParamsReplacePricesPriceNewPlanGroupedWithMeteredMinimumPriceBillingCycleConfiguration struct {
	// The duration of the billing period.
	Duration param.Field[int64] `json:"duration,required"`
	// The unit of billing period duration.
	DurationUnit param.Field[PlanExternalPlanIDVersionNewParamsReplacePricesPriceNewPlanGroupedWithMeteredMinimumPriceBillingCycleConfigurationDurationUnit] `json:"duration_unit,required"`
}

func (r PlanExternalPlanIDVersionNewParamsReplacePricesPriceNewPlanGroupedWithMeteredMinimumPriceBillingCycleConfiguration) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// The unit of billing period duration.
type PlanExternalPlanIDVersionNewParamsReplacePricesPriceNewPlanGroupedWithMeteredMinimumPriceBillingCycleConfigurationDurationUnit string

const (
	PlanExternalPlanIDVersionNewParamsReplacePricesPriceNewPlanGroupedWithMeteredMinimumPriceBillingCycleConfigurationDurationUnitDay   PlanExternalPlanIDVersionNewParamsReplacePricesPriceNewPlanGroupedWithMeteredMinimumPriceBillingCycleConfigurationDurationUnit = "day"
	PlanExternalPlanIDVersionNewParamsReplacePricesPriceNewPlanGroupedWithMeteredMinimumPriceBillingCycleConfigurationDurationUnitMonth PlanExternalPlanIDVersionNewParamsReplacePricesPriceNewPlanGroupedWithMeteredMinimumPriceBillingCycleConfigurationDurationUnit = "month"
)

func (r PlanExternalPlanIDVersionNewParamsReplacePricesPriceNewPlanGroupedWithMeteredMinimumPriceBillingCycleConfigurationDurationUnit) IsKnown() bool {
	switch r {
	case PlanExternalPlanIDVersionNewParamsReplacePricesPriceNewPlanGroupedWithMeteredMinimumPriceBillingCycleConfigurationDurationUnitDay, PlanExternalPlanIDVersionNewParamsReplacePricesPriceNewPlanGroupedWithMeteredMinimumPriceBillingCycleConfigurationDurationUnitMonth:
		return true
	}
	return false
}

// For dimensional price: specifies a price group and dimension values
type PlanExternalPlanIDVersionNewParamsReplacePricesPriceNewPlanGroupedWithMeteredMinimumPriceDimensionalPriceConfiguration struct {
	// The list of dimension values matching (in order) the dimensions of the price
	// group
	DimensionValues param.Field[[]string] `json:"dimension_values,required"`
	// The id of the dimensional price group to include this price in
	DimensionalPriceGroupID param.Field[string] `json:"dimensional_price_group_id"`
	// The external id of the dimensional price group to include this price in
	ExternalDimensionalPriceGroupID param.Field[string] `json:"external_dimensional_price_group_id"`
}

func (r PlanExternalPlanIDVersionNewParamsReplacePricesPriceNewPlanGroupedWithMeteredMinimumPriceDimensionalPriceConfiguration) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Within each billing cycle, specifies the cadence at which invoices are produced.
// If unspecified, a single invoice is produced per billing cycle.
type PlanExternalPlanIDVersionNewParamsReplacePricesPriceNewPlanGroupedWithMeteredMinimumPriceInvoicingCycleConfiguration struct {
	// The duration of the billing period.
	Duration param.Field[int64] `json:"duration,required"`
	// The unit of billing period duration.
	DurationUnit param.Field[PlanExternalPlanIDVersionNewParamsReplacePricesPriceNewPlanGroupedWithMeteredMinimumPriceInvoicingCycleConfigurationDurationUnit] `json:"duration_unit,required"`
}

func (r PlanExternalPlanIDVersionNewParamsReplacePricesPriceNewPlanGroupedWithMeteredMinimumPriceInvoicingCycleConfiguration) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// The unit of billing period duration.
type PlanExternalPlanIDVersionNewParamsReplacePricesPriceNewPlanGroupedWithMeteredMinimumPriceInvoicingCycleConfigurationDurationUnit string

const (
	PlanExternalPlanIDVersionNewParamsReplacePricesPriceNewPlanGroupedWithMeteredMinimumPriceInvoicingCycleConfigurationDurationUnitDay   PlanExternalPlanIDVersionNewParamsReplacePricesPriceNewPlanGroupedWithMeteredMinimumPriceInvoicingCycleConfigurationDurationUnit = "day"
	PlanExternalPlanIDVersionNewParamsReplacePricesPriceNewPlanGroupedWithMeteredMinimumPriceInvoicingCycleConfigurationDurationUnitMonth PlanExternalPlanIDVersionNewParamsReplacePricesPriceNewPlanGroupedWithMeteredMinimumPriceInvoicingCycleConfigurationDurationUnit = "month"
)

func (r PlanExternalPlanIDVersionNewParamsReplacePricesPriceNewPlanGroupedWithMeteredMinimumPriceInvoicingCycleConfigurationDurationUnit) IsKnown() bool {
	switch r {
	case PlanExternalPlanIDVersionNewParamsReplacePricesPriceNewPlanGroupedWithMeteredMinimumPriceInvoicingCycleConfigurationDurationUnitDay, PlanExternalPlanIDVersionNewParamsReplacePricesPriceNewPlanGroupedWithMeteredMinimumPriceInvoicingCycleConfigurationDurationUnitMonth:
		return true
	}
	return false
}

type PlanExternalPlanIDVersionNewParamsReplacePricesPriceNewPlanMatrixWithDisplayNamePrice struct {
	// The cadence to bill for this price on.
	Cadence param.Field[PlanExternalPlanIDVersionNewParamsReplacePricesPriceNewPlanMatrixWithDisplayNamePriceCadence] `json:"cadence,required"`
	// The id of the item the price will be associated with.
	ItemID                      param.Field[string]                                                                                         `json:"item_id,required"`
	MatrixWithDisplayNameConfig param.Field[map[string]interface{}]                                                                         `json:"matrix_with_display_name_config,required"`
	ModelType                   param.Field[PlanExternalPlanIDVersionNewParamsReplacePricesPriceNewPlanMatrixWithDisplayNamePriceModelType] `json:"model_type,required"`
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
	BillingCycleConfiguration param.Field[PlanExternalPlanIDVersionNewParamsReplacePricesPriceNewPlanMatrixWithDisplayNamePriceBillingCycleConfiguration] `json:"billing_cycle_configuration"`
	// The per unit conversion rate of the price currency to the invoicing currency.
	ConversionRate param.Field[float64] `json:"conversion_rate"`
	// An ISO 4217 currency string, or custom pricing unit identifier, in which this
	// price is billed.
	Currency param.Field[string] `json:"currency"`
	// For dimensional price: specifies a price group and dimension values
	DimensionalPriceConfiguration param.Field[PlanExternalPlanIDVersionNewParamsReplacePricesPriceNewPlanMatrixWithDisplayNamePriceDimensionalPriceConfiguration] `json:"dimensional_price_configuration"`
	// An alias for the price.
	ExternalPriceID param.Field[string] `json:"external_price_id"`
	// If the Price represents a fixed cost, this represents the quantity of units
	// applied.
	FixedPriceQuantity param.Field[float64] `json:"fixed_price_quantity"`
	// The property used to group this price on an invoice
	InvoiceGroupingKey param.Field[string] `json:"invoice_grouping_key"`
	// Within each billing cycle, specifies the cadence at which invoices are produced.
	// If unspecified, a single invoice is produced per billing cycle.
	InvoicingCycleConfiguration param.Field[PlanExternalPlanIDVersionNewParamsReplacePricesPriceNewPlanMatrixWithDisplayNamePriceInvoicingCycleConfiguration] `json:"invoicing_cycle_configuration"`
	// User-specified key/value pairs for the resource. Individual keys can be removed
	// by setting the value to `null`, and the entire metadata mapping can be cleared
	// by setting `metadata` to `null`.
	Metadata param.Field[map[string]string] `json:"metadata"`
}

func (r PlanExternalPlanIDVersionNewParamsReplacePricesPriceNewPlanMatrixWithDisplayNamePrice) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r PlanExternalPlanIDVersionNewParamsReplacePricesPriceNewPlanMatrixWithDisplayNamePrice) implementsPlanExternalPlanIDVersionNewParamsReplacePricesPriceUnion() {
}

// The cadence to bill for this price on.
type PlanExternalPlanIDVersionNewParamsReplacePricesPriceNewPlanMatrixWithDisplayNamePriceCadence string

const (
	PlanExternalPlanIDVersionNewParamsReplacePricesPriceNewPlanMatrixWithDisplayNamePriceCadenceAnnual     PlanExternalPlanIDVersionNewParamsReplacePricesPriceNewPlanMatrixWithDisplayNamePriceCadence = "annual"
	PlanExternalPlanIDVersionNewParamsReplacePricesPriceNewPlanMatrixWithDisplayNamePriceCadenceSemiAnnual PlanExternalPlanIDVersionNewParamsReplacePricesPriceNewPlanMatrixWithDisplayNamePriceCadence = "semi_annual"
	PlanExternalPlanIDVersionNewParamsReplacePricesPriceNewPlanMatrixWithDisplayNamePriceCadenceMonthly    PlanExternalPlanIDVersionNewParamsReplacePricesPriceNewPlanMatrixWithDisplayNamePriceCadence = "monthly"
	PlanExternalPlanIDVersionNewParamsReplacePricesPriceNewPlanMatrixWithDisplayNamePriceCadenceQuarterly  PlanExternalPlanIDVersionNewParamsReplacePricesPriceNewPlanMatrixWithDisplayNamePriceCadence = "quarterly"
	PlanExternalPlanIDVersionNewParamsReplacePricesPriceNewPlanMatrixWithDisplayNamePriceCadenceOneTime    PlanExternalPlanIDVersionNewParamsReplacePricesPriceNewPlanMatrixWithDisplayNamePriceCadence = "one_time"
	PlanExternalPlanIDVersionNewParamsReplacePricesPriceNewPlanMatrixWithDisplayNamePriceCadenceCustom     PlanExternalPlanIDVersionNewParamsReplacePricesPriceNewPlanMatrixWithDisplayNamePriceCadence = "custom"
)

func (r PlanExternalPlanIDVersionNewParamsReplacePricesPriceNewPlanMatrixWithDisplayNamePriceCadence) IsKnown() bool {
	switch r {
	case PlanExternalPlanIDVersionNewParamsReplacePricesPriceNewPlanMatrixWithDisplayNamePriceCadenceAnnual, PlanExternalPlanIDVersionNewParamsReplacePricesPriceNewPlanMatrixWithDisplayNamePriceCadenceSemiAnnual, PlanExternalPlanIDVersionNewParamsReplacePricesPriceNewPlanMatrixWithDisplayNamePriceCadenceMonthly, PlanExternalPlanIDVersionNewParamsReplacePricesPriceNewPlanMatrixWithDisplayNamePriceCadenceQuarterly, PlanExternalPlanIDVersionNewParamsReplacePricesPriceNewPlanMatrixWithDisplayNamePriceCadenceOneTime, PlanExternalPlanIDVersionNewParamsReplacePricesPriceNewPlanMatrixWithDisplayNamePriceCadenceCustom:
		return true
	}
	return false
}

type PlanExternalPlanIDVersionNewParamsReplacePricesPriceNewPlanMatrixWithDisplayNamePriceModelType string

const (
	PlanExternalPlanIDVersionNewParamsReplacePricesPriceNewPlanMatrixWithDisplayNamePriceModelTypeMatrixWithDisplayName PlanExternalPlanIDVersionNewParamsReplacePricesPriceNewPlanMatrixWithDisplayNamePriceModelType = "matrix_with_display_name"
)

func (r PlanExternalPlanIDVersionNewParamsReplacePricesPriceNewPlanMatrixWithDisplayNamePriceModelType) IsKnown() bool {
	switch r {
	case PlanExternalPlanIDVersionNewParamsReplacePricesPriceNewPlanMatrixWithDisplayNamePriceModelTypeMatrixWithDisplayName:
		return true
	}
	return false
}

// For custom cadence: specifies the duration of the billing period in days or
// months.
type PlanExternalPlanIDVersionNewParamsReplacePricesPriceNewPlanMatrixWithDisplayNamePriceBillingCycleConfiguration struct {
	// The duration of the billing period.
	Duration param.Field[int64] `json:"duration,required"`
	// The unit of billing period duration.
	DurationUnit param.Field[PlanExternalPlanIDVersionNewParamsReplacePricesPriceNewPlanMatrixWithDisplayNamePriceBillingCycleConfigurationDurationUnit] `json:"duration_unit,required"`
}

func (r PlanExternalPlanIDVersionNewParamsReplacePricesPriceNewPlanMatrixWithDisplayNamePriceBillingCycleConfiguration) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// The unit of billing period duration.
type PlanExternalPlanIDVersionNewParamsReplacePricesPriceNewPlanMatrixWithDisplayNamePriceBillingCycleConfigurationDurationUnit string

const (
	PlanExternalPlanIDVersionNewParamsReplacePricesPriceNewPlanMatrixWithDisplayNamePriceBillingCycleConfigurationDurationUnitDay   PlanExternalPlanIDVersionNewParamsReplacePricesPriceNewPlanMatrixWithDisplayNamePriceBillingCycleConfigurationDurationUnit = "day"
	PlanExternalPlanIDVersionNewParamsReplacePricesPriceNewPlanMatrixWithDisplayNamePriceBillingCycleConfigurationDurationUnitMonth PlanExternalPlanIDVersionNewParamsReplacePricesPriceNewPlanMatrixWithDisplayNamePriceBillingCycleConfigurationDurationUnit = "month"
)

func (r PlanExternalPlanIDVersionNewParamsReplacePricesPriceNewPlanMatrixWithDisplayNamePriceBillingCycleConfigurationDurationUnit) IsKnown() bool {
	switch r {
	case PlanExternalPlanIDVersionNewParamsReplacePricesPriceNewPlanMatrixWithDisplayNamePriceBillingCycleConfigurationDurationUnitDay, PlanExternalPlanIDVersionNewParamsReplacePricesPriceNewPlanMatrixWithDisplayNamePriceBillingCycleConfigurationDurationUnitMonth:
		return true
	}
	return false
}

// For dimensional price: specifies a price group and dimension values
type PlanExternalPlanIDVersionNewParamsReplacePricesPriceNewPlanMatrixWithDisplayNamePriceDimensionalPriceConfiguration struct {
	// The list of dimension values matching (in order) the dimensions of the price
	// group
	DimensionValues param.Field[[]string] `json:"dimension_values,required"`
	// The id of the dimensional price group to include this price in
	DimensionalPriceGroupID param.Field[string] `json:"dimensional_price_group_id"`
	// The external id of the dimensional price group to include this price in
	ExternalDimensionalPriceGroupID param.Field[string] `json:"external_dimensional_price_group_id"`
}

func (r PlanExternalPlanIDVersionNewParamsReplacePricesPriceNewPlanMatrixWithDisplayNamePriceDimensionalPriceConfiguration) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Within each billing cycle, specifies the cadence at which invoices are produced.
// If unspecified, a single invoice is produced per billing cycle.
type PlanExternalPlanIDVersionNewParamsReplacePricesPriceNewPlanMatrixWithDisplayNamePriceInvoicingCycleConfiguration struct {
	// The duration of the billing period.
	Duration param.Field[int64] `json:"duration,required"`
	// The unit of billing period duration.
	DurationUnit param.Field[PlanExternalPlanIDVersionNewParamsReplacePricesPriceNewPlanMatrixWithDisplayNamePriceInvoicingCycleConfigurationDurationUnit] `json:"duration_unit,required"`
}

func (r PlanExternalPlanIDVersionNewParamsReplacePricesPriceNewPlanMatrixWithDisplayNamePriceInvoicingCycleConfiguration) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// The unit of billing period duration.
type PlanExternalPlanIDVersionNewParamsReplacePricesPriceNewPlanMatrixWithDisplayNamePriceInvoicingCycleConfigurationDurationUnit string

const (
	PlanExternalPlanIDVersionNewParamsReplacePricesPriceNewPlanMatrixWithDisplayNamePriceInvoicingCycleConfigurationDurationUnitDay   PlanExternalPlanIDVersionNewParamsReplacePricesPriceNewPlanMatrixWithDisplayNamePriceInvoicingCycleConfigurationDurationUnit = "day"
	PlanExternalPlanIDVersionNewParamsReplacePricesPriceNewPlanMatrixWithDisplayNamePriceInvoicingCycleConfigurationDurationUnitMonth PlanExternalPlanIDVersionNewParamsReplacePricesPriceNewPlanMatrixWithDisplayNamePriceInvoicingCycleConfigurationDurationUnit = "month"
)

func (r PlanExternalPlanIDVersionNewParamsReplacePricesPriceNewPlanMatrixWithDisplayNamePriceInvoicingCycleConfigurationDurationUnit) IsKnown() bool {
	switch r {
	case PlanExternalPlanIDVersionNewParamsReplacePricesPriceNewPlanMatrixWithDisplayNamePriceInvoicingCycleConfigurationDurationUnitDay, PlanExternalPlanIDVersionNewParamsReplacePricesPriceNewPlanMatrixWithDisplayNamePriceInvoicingCycleConfigurationDurationUnitMonth:
		return true
	}
	return false
}

type PlanExternalPlanIDVersionNewParamsReplacePricesPriceNewPlanBulkWithProrationPrice struct {
	BulkWithProrationConfig param.Field[map[string]interface{}] `json:"bulk_with_proration_config,required"`
	// The cadence to bill for this price on.
	Cadence param.Field[PlanExternalPlanIDVersionNewParamsReplacePricesPriceNewPlanBulkWithProrationPriceCadence] `json:"cadence,required"`
	// The id of the item the price will be associated with.
	ItemID    param.Field[string]                                                                                     `json:"item_id,required"`
	ModelType param.Field[PlanExternalPlanIDVersionNewParamsReplacePricesPriceNewPlanBulkWithProrationPriceModelType] `json:"model_type,required"`
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
	BillingCycleConfiguration param.Field[PlanExternalPlanIDVersionNewParamsReplacePricesPriceNewPlanBulkWithProrationPriceBillingCycleConfiguration] `json:"billing_cycle_configuration"`
	// The per unit conversion rate of the price currency to the invoicing currency.
	ConversionRate param.Field[float64] `json:"conversion_rate"`
	// An ISO 4217 currency string, or custom pricing unit identifier, in which this
	// price is billed.
	Currency param.Field[string] `json:"currency"`
	// For dimensional price: specifies a price group and dimension values
	DimensionalPriceConfiguration param.Field[PlanExternalPlanIDVersionNewParamsReplacePricesPriceNewPlanBulkWithProrationPriceDimensionalPriceConfiguration] `json:"dimensional_price_configuration"`
	// An alias for the price.
	ExternalPriceID param.Field[string] `json:"external_price_id"`
	// If the Price represents a fixed cost, this represents the quantity of units
	// applied.
	FixedPriceQuantity param.Field[float64] `json:"fixed_price_quantity"`
	// The property used to group this price on an invoice
	InvoiceGroupingKey param.Field[string] `json:"invoice_grouping_key"`
	// Within each billing cycle, specifies the cadence at which invoices are produced.
	// If unspecified, a single invoice is produced per billing cycle.
	InvoicingCycleConfiguration param.Field[PlanExternalPlanIDVersionNewParamsReplacePricesPriceNewPlanBulkWithProrationPriceInvoicingCycleConfiguration] `json:"invoicing_cycle_configuration"`
	// User-specified key/value pairs for the resource. Individual keys can be removed
	// by setting the value to `null`, and the entire metadata mapping can be cleared
	// by setting `metadata` to `null`.
	Metadata param.Field[map[string]string] `json:"metadata"`
}

func (r PlanExternalPlanIDVersionNewParamsReplacePricesPriceNewPlanBulkWithProrationPrice) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r PlanExternalPlanIDVersionNewParamsReplacePricesPriceNewPlanBulkWithProrationPrice) implementsPlanExternalPlanIDVersionNewParamsReplacePricesPriceUnion() {
}

// The cadence to bill for this price on.
type PlanExternalPlanIDVersionNewParamsReplacePricesPriceNewPlanBulkWithProrationPriceCadence string

const (
	PlanExternalPlanIDVersionNewParamsReplacePricesPriceNewPlanBulkWithProrationPriceCadenceAnnual     PlanExternalPlanIDVersionNewParamsReplacePricesPriceNewPlanBulkWithProrationPriceCadence = "annual"
	PlanExternalPlanIDVersionNewParamsReplacePricesPriceNewPlanBulkWithProrationPriceCadenceSemiAnnual PlanExternalPlanIDVersionNewParamsReplacePricesPriceNewPlanBulkWithProrationPriceCadence = "semi_annual"
	PlanExternalPlanIDVersionNewParamsReplacePricesPriceNewPlanBulkWithProrationPriceCadenceMonthly    PlanExternalPlanIDVersionNewParamsReplacePricesPriceNewPlanBulkWithProrationPriceCadence = "monthly"
	PlanExternalPlanIDVersionNewParamsReplacePricesPriceNewPlanBulkWithProrationPriceCadenceQuarterly  PlanExternalPlanIDVersionNewParamsReplacePricesPriceNewPlanBulkWithProrationPriceCadence = "quarterly"
	PlanExternalPlanIDVersionNewParamsReplacePricesPriceNewPlanBulkWithProrationPriceCadenceOneTime    PlanExternalPlanIDVersionNewParamsReplacePricesPriceNewPlanBulkWithProrationPriceCadence = "one_time"
	PlanExternalPlanIDVersionNewParamsReplacePricesPriceNewPlanBulkWithProrationPriceCadenceCustom     PlanExternalPlanIDVersionNewParamsReplacePricesPriceNewPlanBulkWithProrationPriceCadence = "custom"
)

func (r PlanExternalPlanIDVersionNewParamsReplacePricesPriceNewPlanBulkWithProrationPriceCadence) IsKnown() bool {
	switch r {
	case PlanExternalPlanIDVersionNewParamsReplacePricesPriceNewPlanBulkWithProrationPriceCadenceAnnual, PlanExternalPlanIDVersionNewParamsReplacePricesPriceNewPlanBulkWithProrationPriceCadenceSemiAnnual, PlanExternalPlanIDVersionNewParamsReplacePricesPriceNewPlanBulkWithProrationPriceCadenceMonthly, PlanExternalPlanIDVersionNewParamsReplacePricesPriceNewPlanBulkWithProrationPriceCadenceQuarterly, PlanExternalPlanIDVersionNewParamsReplacePricesPriceNewPlanBulkWithProrationPriceCadenceOneTime, PlanExternalPlanIDVersionNewParamsReplacePricesPriceNewPlanBulkWithProrationPriceCadenceCustom:
		return true
	}
	return false
}

type PlanExternalPlanIDVersionNewParamsReplacePricesPriceNewPlanBulkWithProrationPriceModelType string

const (
	PlanExternalPlanIDVersionNewParamsReplacePricesPriceNewPlanBulkWithProrationPriceModelTypeBulkWithProration PlanExternalPlanIDVersionNewParamsReplacePricesPriceNewPlanBulkWithProrationPriceModelType = "bulk_with_proration"
)

func (r PlanExternalPlanIDVersionNewParamsReplacePricesPriceNewPlanBulkWithProrationPriceModelType) IsKnown() bool {
	switch r {
	case PlanExternalPlanIDVersionNewParamsReplacePricesPriceNewPlanBulkWithProrationPriceModelTypeBulkWithProration:
		return true
	}
	return false
}

// For custom cadence: specifies the duration of the billing period in days or
// months.
type PlanExternalPlanIDVersionNewParamsReplacePricesPriceNewPlanBulkWithProrationPriceBillingCycleConfiguration struct {
	// The duration of the billing period.
	Duration param.Field[int64] `json:"duration,required"`
	// The unit of billing period duration.
	DurationUnit param.Field[PlanExternalPlanIDVersionNewParamsReplacePricesPriceNewPlanBulkWithProrationPriceBillingCycleConfigurationDurationUnit] `json:"duration_unit,required"`
}

func (r PlanExternalPlanIDVersionNewParamsReplacePricesPriceNewPlanBulkWithProrationPriceBillingCycleConfiguration) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// The unit of billing period duration.
type PlanExternalPlanIDVersionNewParamsReplacePricesPriceNewPlanBulkWithProrationPriceBillingCycleConfigurationDurationUnit string

const (
	PlanExternalPlanIDVersionNewParamsReplacePricesPriceNewPlanBulkWithProrationPriceBillingCycleConfigurationDurationUnitDay   PlanExternalPlanIDVersionNewParamsReplacePricesPriceNewPlanBulkWithProrationPriceBillingCycleConfigurationDurationUnit = "day"
	PlanExternalPlanIDVersionNewParamsReplacePricesPriceNewPlanBulkWithProrationPriceBillingCycleConfigurationDurationUnitMonth PlanExternalPlanIDVersionNewParamsReplacePricesPriceNewPlanBulkWithProrationPriceBillingCycleConfigurationDurationUnit = "month"
)

func (r PlanExternalPlanIDVersionNewParamsReplacePricesPriceNewPlanBulkWithProrationPriceBillingCycleConfigurationDurationUnit) IsKnown() bool {
	switch r {
	case PlanExternalPlanIDVersionNewParamsReplacePricesPriceNewPlanBulkWithProrationPriceBillingCycleConfigurationDurationUnitDay, PlanExternalPlanIDVersionNewParamsReplacePricesPriceNewPlanBulkWithProrationPriceBillingCycleConfigurationDurationUnitMonth:
		return true
	}
	return false
}

// For dimensional price: specifies a price group and dimension values
type PlanExternalPlanIDVersionNewParamsReplacePricesPriceNewPlanBulkWithProrationPriceDimensionalPriceConfiguration struct {
	// The list of dimension values matching (in order) the dimensions of the price
	// group
	DimensionValues param.Field[[]string] `json:"dimension_values,required"`
	// The id of the dimensional price group to include this price in
	DimensionalPriceGroupID param.Field[string] `json:"dimensional_price_group_id"`
	// The external id of the dimensional price group to include this price in
	ExternalDimensionalPriceGroupID param.Field[string] `json:"external_dimensional_price_group_id"`
}

func (r PlanExternalPlanIDVersionNewParamsReplacePricesPriceNewPlanBulkWithProrationPriceDimensionalPriceConfiguration) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Within each billing cycle, specifies the cadence at which invoices are produced.
// If unspecified, a single invoice is produced per billing cycle.
type PlanExternalPlanIDVersionNewParamsReplacePricesPriceNewPlanBulkWithProrationPriceInvoicingCycleConfiguration struct {
	// The duration of the billing period.
	Duration param.Field[int64] `json:"duration,required"`
	// The unit of billing period duration.
	DurationUnit param.Field[PlanExternalPlanIDVersionNewParamsReplacePricesPriceNewPlanBulkWithProrationPriceInvoicingCycleConfigurationDurationUnit] `json:"duration_unit,required"`
}

func (r PlanExternalPlanIDVersionNewParamsReplacePricesPriceNewPlanBulkWithProrationPriceInvoicingCycleConfiguration) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// The unit of billing period duration.
type PlanExternalPlanIDVersionNewParamsReplacePricesPriceNewPlanBulkWithProrationPriceInvoicingCycleConfigurationDurationUnit string

const (
	PlanExternalPlanIDVersionNewParamsReplacePricesPriceNewPlanBulkWithProrationPriceInvoicingCycleConfigurationDurationUnitDay   PlanExternalPlanIDVersionNewParamsReplacePricesPriceNewPlanBulkWithProrationPriceInvoicingCycleConfigurationDurationUnit = "day"
	PlanExternalPlanIDVersionNewParamsReplacePricesPriceNewPlanBulkWithProrationPriceInvoicingCycleConfigurationDurationUnitMonth PlanExternalPlanIDVersionNewParamsReplacePricesPriceNewPlanBulkWithProrationPriceInvoicingCycleConfigurationDurationUnit = "month"
)

func (r PlanExternalPlanIDVersionNewParamsReplacePricesPriceNewPlanBulkWithProrationPriceInvoicingCycleConfigurationDurationUnit) IsKnown() bool {
	switch r {
	case PlanExternalPlanIDVersionNewParamsReplacePricesPriceNewPlanBulkWithProrationPriceInvoicingCycleConfigurationDurationUnitDay, PlanExternalPlanIDVersionNewParamsReplacePricesPriceNewPlanBulkWithProrationPriceInvoicingCycleConfigurationDurationUnitMonth:
		return true
	}
	return false
}

type PlanExternalPlanIDVersionNewParamsReplacePricesPriceNewPlanGroupedTieredPackagePrice struct {
	// The cadence to bill for this price on.
	Cadence                    param.Field[PlanExternalPlanIDVersionNewParamsReplacePricesPriceNewPlanGroupedTieredPackagePriceCadence] `json:"cadence,required"`
	GroupedTieredPackageConfig param.Field[map[string]interface{}]                                                                      `json:"grouped_tiered_package_config,required"`
	// The id of the item the price will be associated with.
	ItemID    param.Field[string]                                                                                        `json:"item_id,required"`
	ModelType param.Field[PlanExternalPlanIDVersionNewParamsReplacePricesPriceNewPlanGroupedTieredPackagePriceModelType] `json:"model_type,required"`
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
	BillingCycleConfiguration param.Field[PlanExternalPlanIDVersionNewParamsReplacePricesPriceNewPlanGroupedTieredPackagePriceBillingCycleConfiguration] `json:"billing_cycle_configuration"`
	// The per unit conversion rate of the price currency to the invoicing currency.
	ConversionRate param.Field[float64] `json:"conversion_rate"`
	// An ISO 4217 currency string, or custom pricing unit identifier, in which this
	// price is billed.
	Currency param.Field[string] `json:"currency"`
	// For dimensional price: specifies a price group and dimension values
	DimensionalPriceConfiguration param.Field[PlanExternalPlanIDVersionNewParamsReplacePricesPriceNewPlanGroupedTieredPackagePriceDimensionalPriceConfiguration] `json:"dimensional_price_configuration"`
	// An alias for the price.
	ExternalPriceID param.Field[string] `json:"external_price_id"`
	// If the Price represents a fixed cost, this represents the quantity of units
	// applied.
	FixedPriceQuantity param.Field[float64] `json:"fixed_price_quantity"`
	// The property used to group this price on an invoice
	InvoiceGroupingKey param.Field[string] `json:"invoice_grouping_key"`
	// Within each billing cycle, specifies the cadence at which invoices are produced.
	// If unspecified, a single invoice is produced per billing cycle.
	InvoicingCycleConfiguration param.Field[PlanExternalPlanIDVersionNewParamsReplacePricesPriceNewPlanGroupedTieredPackagePriceInvoicingCycleConfiguration] `json:"invoicing_cycle_configuration"`
	// User-specified key/value pairs for the resource. Individual keys can be removed
	// by setting the value to `null`, and the entire metadata mapping can be cleared
	// by setting `metadata` to `null`.
	Metadata param.Field[map[string]string] `json:"metadata"`
}

func (r PlanExternalPlanIDVersionNewParamsReplacePricesPriceNewPlanGroupedTieredPackagePrice) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r PlanExternalPlanIDVersionNewParamsReplacePricesPriceNewPlanGroupedTieredPackagePrice) implementsPlanExternalPlanIDVersionNewParamsReplacePricesPriceUnion() {
}

// The cadence to bill for this price on.
type PlanExternalPlanIDVersionNewParamsReplacePricesPriceNewPlanGroupedTieredPackagePriceCadence string

const (
	PlanExternalPlanIDVersionNewParamsReplacePricesPriceNewPlanGroupedTieredPackagePriceCadenceAnnual     PlanExternalPlanIDVersionNewParamsReplacePricesPriceNewPlanGroupedTieredPackagePriceCadence = "annual"
	PlanExternalPlanIDVersionNewParamsReplacePricesPriceNewPlanGroupedTieredPackagePriceCadenceSemiAnnual PlanExternalPlanIDVersionNewParamsReplacePricesPriceNewPlanGroupedTieredPackagePriceCadence = "semi_annual"
	PlanExternalPlanIDVersionNewParamsReplacePricesPriceNewPlanGroupedTieredPackagePriceCadenceMonthly    PlanExternalPlanIDVersionNewParamsReplacePricesPriceNewPlanGroupedTieredPackagePriceCadence = "monthly"
	PlanExternalPlanIDVersionNewParamsReplacePricesPriceNewPlanGroupedTieredPackagePriceCadenceQuarterly  PlanExternalPlanIDVersionNewParamsReplacePricesPriceNewPlanGroupedTieredPackagePriceCadence = "quarterly"
	PlanExternalPlanIDVersionNewParamsReplacePricesPriceNewPlanGroupedTieredPackagePriceCadenceOneTime    PlanExternalPlanIDVersionNewParamsReplacePricesPriceNewPlanGroupedTieredPackagePriceCadence = "one_time"
	PlanExternalPlanIDVersionNewParamsReplacePricesPriceNewPlanGroupedTieredPackagePriceCadenceCustom     PlanExternalPlanIDVersionNewParamsReplacePricesPriceNewPlanGroupedTieredPackagePriceCadence = "custom"
)

func (r PlanExternalPlanIDVersionNewParamsReplacePricesPriceNewPlanGroupedTieredPackagePriceCadence) IsKnown() bool {
	switch r {
	case PlanExternalPlanIDVersionNewParamsReplacePricesPriceNewPlanGroupedTieredPackagePriceCadenceAnnual, PlanExternalPlanIDVersionNewParamsReplacePricesPriceNewPlanGroupedTieredPackagePriceCadenceSemiAnnual, PlanExternalPlanIDVersionNewParamsReplacePricesPriceNewPlanGroupedTieredPackagePriceCadenceMonthly, PlanExternalPlanIDVersionNewParamsReplacePricesPriceNewPlanGroupedTieredPackagePriceCadenceQuarterly, PlanExternalPlanIDVersionNewParamsReplacePricesPriceNewPlanGroupedTieredPackagePriceCadenceOneTime, PlanExternalPlanIDVersionNewParamsReplacePricesPriceNewPlanGroupedTieredPackagePriceCadenceCustom:
		return true
	}
	return false
}

type PlanExternalPlanIDVersionNewParamsReplacePricesPriceNewPlanGroupedTieredPackagePriceModelType string

const (
	PlanExternalPlanIDVersionNewParamsReplacePricesPriceNewPlanGroupedTieredPackagePriceModelTypeGroupedTieredPackage PlanExternalPlanIDVersionNewParamsReplacePricesPriceNewPlanGroupedTieredPackagePriceModelType = "grouped_tiered_package"
)

func (r PlanExternalPlanIDVersionNewParamsReplacePricesPriceNewPlanGroupedTieredPackagePriceModelType) IsKnown() bool {
	switch r {
	case PlanExternalPlanIDVersionNewParamsReplacePricesPriceNewPlanGroupedTieredPackagePriceModelTypeGroupedTieredPackage:
		return true
	}
	return false
}

// For custom cadence: specifies the duration of the billing period in days or
// months.
type PlanExternalPlanIDVersionNewParamsReplacePricesPriceNewPlanGroupedTieredPackagePriceBillingCycleConfiguration struct {
	// The duration of the billing period.
	Duration param.Field[int64] `json:"duration,required"`
	// The unit of billing period duration.
	DurationUnit param.Field[PlanExternalPlanIDVersionNewParamsReplacePricesPriceNewPlanGroupedTieredPackagePriceBillingCycleConfigurationDurationUnit] `json:"duration_unit,required"`
}

func (r PlanExternalPlanIDVersionNewParamsReplacePricesPriceNewPlanGroupedTieredPackagePriceBillingCycleConfiguration) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// The unit of billing period duration.
type PlanExternalPlanIDVersionNewParamsReplacePricesPriceNewPlanGroupedTieredPackagePriceBillingCycleConfigurationDurationUnit string

const (
	PlanExternalPlanIDVersionNewParamsReplacePricesPriceNewPlanGroupedTieredPackagePriceBillingCycleConfigurationDurationUnitDay   PlanExternalPlanIDVersionNewParamsReplacePricesPriceNewPlanGroupedTieredPackagePriceBillingCycleConfigurationDurationUnit = "day"
	PlanExternalPlanIDVersionNewParamsReplacePricesPriceNewPlanGroupedTieredPackagePriceBillingCycleConfigurationDurationUnitMonth PlanExternalPlanIDVersionNewParamsReplacePricesPriceNewPlanGroupedTieredPackagePriceBillingCycleConfigurationDurationUnit = "month"
)

func (r PlanExternalPlanIDVersionNewParamsReplacePricesPriceNewPlanGroupedTieredPackagePriceBillingCycleConfigurationDurationUnit) IsKnown() bool {
	switch r {
	case PlanExternalPlanIDVersionNewParamsReplacePricesPriceNewPlanGroupedTieredPackagePriceBillingCycleConfigurationDurationUnitDay, PlanExternalPlanIDVersionNewParamsReplacePricesPriceNewPlanGroupedTieredPackagePriceBillingCycleConfigurationDurationUnitMonth:
		return true
	}
	return false
}

// For dimensional price: specifies a price group and dimension values
type PlanExternalPlanIDVersionNewParamsReplacePricesPriceNewPlanGroupedTieredPackagePriceDimensionalPriceConfiguration struct {
	// The list of dimension values matching (in order) the dimensions of the price
	// group
	DimensionValues param.Field[[]string] `json:"dimension_values,required"`
	// The id of the dimensional price group to include this price in
	DimensionalPriceGroupID param.Field[string] `json:"dimensional_price_group_id"`
	// The external id of the dimensional price group to include this price in
	ExternalDimensionalPriceGroupID param.Field[string] `json:"external_dimensional_price_group_id"`
}

func (r PlanExternalPlanIDVersionNewParamsReplacePricesPriceNewPlanGroupedTieredPackagePriceDimensionalPriceConfiguration) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Within each billing cycle, specifies the cadence at which invoices are produced.
// If unspecified, a single invoice is produced per billing cycle.
type PlanExternalPlanIDVersionNewParamsReplacePricesPriceNewPlanGroupedTieredPackagePriceInvoicingCycleConfiguration struct {
	// The duration of the billing period.
	Duration param.Field[int64] `json:"duration,required"`
	// The unit of billing period duration.
	DurationUnit param.Field[PlanExternalPlanIDVersionNewParamsReplacePricesPriceNewPlanGroupedTieredPackagePriceInvoicingCycleConfigurationDurationUnit] `json:"duration_unit,required"`
}

func (r PlanExternalPlanIDVersionNewParamsReplacePricesPriceNewPlanGroupedTieredPackagePriceInvoicingCycleConfiguration) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// The unit of billing period duration.
type PlanExternalPlanIDVersionNewParamsReplacePricesPriceNewPlanGroupedTieredPackagePriceInvoicingCycleConfigurationDurationUnit string

const (
	PlanExternalPlanIDVersionNewParamsReplacePricesPriceNewPlanGroupedTieredPackagePriceInvoicingCycleConfigurationDurationUnitDay   PlanExternalPlanIDVersionNewParamsReplacePricesPriceNewPlanGroupedTieredPackagePriceInvoicingCycleConfigurationDurationUnit = "day"
	PlanExternalPlanIDVersionNewParamsReplacePricesPriceNewPlanGroupedTieredPackagePriceInvoicingCycleConfigurationDurationUnitMonth PlanExternalPlanIDVersionNewParamsReplacePricesPriceNewPlanGroupedTieredPackagePriceInvoicingCycleConfigurationDurationUnit = "month"
)

func (r PlanExternalPlanIDVersionNewParamsReplacePricesPriceNewPlanGroupedTieredPackagePriceInvoicingCycleConfigurationDurationUnit) IsKnown() bool {
	switch r {
	case PlanExternalPlanIDVersionNewParamsReplacePricesPriceNewPlanGroupedTieredPackagePriceInvoicingCycleConfigurationDurationUnitDay, PlanExternalPlanIDVersionNewParamsReplacePricesPriceNewPlanGroupedTieredPackagePriceInvoicingCycleConfigurationDurationUnitMonth:
		return true
	}
	return false
}

type PlanExternalPlanIDVersionNewParamsReplacePricesPriceNewPlanMaxGroupTieredPackagePrice struct {
	// The cadence to bill for this price on.
	Cadence param.Field[PlanExternalPlanIDVersionNewParamsReplacePricesPriceNewPlanMaxGroupTieredPackagePriceCadence] `json:"cadence,required"`
	// The id of the item the price will be associated with.
	ItemID                      param.Field[string]                                                                                         `json:"item_id,required"`
	MaxGroupTieredPackageConfig param.Field[map[string]interface{}]                                                                         `json:"max_group_tiered_package_config,required"`
	ModelType                   param.Field[PlanExternalPlanIDVersionNewParamsReplacePricesPriceNewPlanMaxGroupTieredPackagePriceModelType] `json:"model_type,required"`
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
	BillingCycleConfiguration param.Field[PlanExternalPlanIDVersionNewParamsReplacePricesPriceNewPlanMaxGroupTieredPackagePriceBillingCycleConfiguration] `json:"billing_cycle_configuration"`
	// The per unit conversion rate of the price currency to the invoicing currency.
	ConversionRate param.Field[float64] `json:"conversion_rate"`
	// An ISO 4217 currency string, or custom pricing unit identifier, in which this
	// price is billed.
	Currency param.Field[string] `json:"currency"`
	// For dimensional price: specifies a price group and dimension values
	DimensionalPriceConfiguration param.Field[PlanExternalPlanIDVersionNewParamsReplacePricesPriceNewPlanMaxGroupTieredPackagePriceDimensionalPriceConfiguration] `json:"dimensional_price_configuration"`
	// An alias for the price.
	ExternalPriceID param.Field[string] `json:"external_price_id"`
	// If the Price represents a fixed cost, this represents the quantity of units
	// applied.
	FixedPriceQuantity param.Field[float64] `json:"fixed_price_quantity"`
	// The property used to group this price on an invoice
	InvoiceGroupingKey param.Field[string] `json:"invoice_grouping_key"`
	// Within each billing cycle, specifies the cadence at which invoices are produced.
	// If unspecified, a single invoice is produced per billing cycle.
	InvoicingCycleConfiguration param.Field[PlanExternalPlanIDVersionNewParamsReplacePricesPriceNewPlanMaxGroupTieredPackagePriceInvoicingCycleConfiguration] `json:"invoicing_cycle_configuration"`
	// User-specified key/value pairs for the resource. Individual keys can be removed
	// by setting the value to `null`, and the entire metadata mapping can be cleared
	// by setting `metadata` to `null`.
	Metadata param.Field[map[string]string] `json:"metadata"`
}

func (r PlanExternalPlanIDVersionNewParamsReplacePricesPriceNewPlanMaxGroupTieredPackagePrice) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r PlanExternalPlanIDVersionNewParamsReplacePricesPriceNewPlanMaxGroupTieredPackagePrice) implementsPlanExternalPlanIDVersionNewParamsReplacePricesPriceUnion() {
}

// The cadence to bill for this price on.
type PlanExternalPlanIDVersionNewParamsReplacePricesPriceNewPlanMaxGroupTieredPackagePriceCadence string

const (
	PlanExternalPlanIDVersionNewParamsReplacePricesPriceNewPlanMaxGroupTieredPackagePriceCadenceAnnual     PlanExternalPlanIDVersionNewParamsReplacePricesPriceNewPlanMaxGroupTieredPackagePriceCadence = "annual"
	PlanExternalPlanIDVersionNewParamsReplacePricesPriceNewPlanMaxGroupTieredPackagePriceCadenceSemiAnnual PlanExternalPlanIDVersionNewParamsReplacePricesPriceNewPlanMaxGroupTieredPackagePriceCadence = "semi_annual"
	PlanExternalPlanIDVersionNewParamsReplacePricesPriceNewPlanMaxGroupTieredPackagePriceCadenceMonthly    PlanExternalPlanIDVersionNewParamsReplacePricesPriceNewPlanMaxGroupTieredPackagePriceCadence = "monthly"
	PlanExternalPlanIDVersionNewParamsReplacePricesPriceNewPlanMaxGroupTieredPackagePriceCadenceQuarterly  PlanExternalPlanIDVersionNewParamsReplacePricesPriceNewPlanMaxGroupTieredPackagePriceCadence = "quarterly"
	PlanExternalPlanIDVersionNewParamsReplacePricesPriceNewPlanMaxGroupTieredPackagePriceCadenceOneTime    PlanExternalPlanIDVersionNewParamsReplacePricesPriceNewPlanMaxGroupTieredPackagePriceCadence = "one_time"
	PlanExternalPlanIDVersionNewParamsReplacePricesPriceNewPlanMaxGroupTieredPackagePriceCadenceCustom     PlanExternalPlanIDVersionNewParamsReplacePricesPriceNewPlanMaxGroupTieredPackagePriceCadence = "custom"
)

func (r PlanExternalPlanIDVersionNewParamsReplacePricesPriceNewPlanMaxGroupTieredPackagePriceCadence) IsKnown() bool {
	switch r {
	case PlanExternalPlanIDVersionNewParamsReplacePricesPriceNewPlanMaxGroupTieredPackagePriceCadenceAnnual, PlanExternalPlanIDVersionNewParamsReplacePricesPriceNewPlanMaxGroupTieredPackagePriceCadenceSemiAnnual, PlanExternalPlanIDVersionNewParamsReplacePricesPriceNewPlanMaxGroupTieredPackagePriceCadenceMonthly, PlanExternalPlanIDVersionNewParamsReplacePricesPriceNewPlanMaxGroupTieredPackagePriceCadenceQuarterly, PlanExternalPlanIDVersionNewParamsReplacePricesPriceNewPlanMaxGroupTieredPackagePriceCadenceOneTime, PlanExternalPlanIDVersionNewParamsReplacePricesPriceNewPlanMaxGroupTieredPackagePriceCadenceCustom:
		return true
	}
	return false
}

type PlanExternalPlanIDVersionNewParamsReplacePricesPriceNewPlanMaxGroupTieredPackagePriceModelType string

const (
	PlanExternalPlanIDVersionNewParamsReplacePricesPriceNewPlanMaxGroupTieredPackagePriceModelTypeMaxGroupTieredPackage PlanExternalPlanIDVersionNewParamsReplacePricesPriceNewPlanMaxGroupTieredPackagePriceModelType = "max_group_tiered_package"
)

func (r PlanExternalPlanIDVersionNewParamsReplacePricesPriceNewPlanMaxGroupTieredPackagePriceModelType) IsKnown() bool {
	switch r {
	case PlanExternalPlanIDVersionNewParamsReplacePricesPriceNewPlanMaxGroupTieredPackagePriceModelTypeMaxGroupTieredPackage:
		return true
	}
	return false
}

// For custom cadence: specifies the duration of the billing period in days or
// months.
type PlanExternalPlanIDVersionNewParamsReplacePricesPriceNewPlanMaxGroupTieredPackagePriceBillingCycleConfiguration struct {
	// The duration of the billing period.
	Duration param.Field[int64] `json:"duration,required"`
	// The unit of billing period duration.
	DurationUnit param.Field[PlanExternalPlanIDVersionNewParamsReplacePricesPriceNewPlanMaxGroupTieredPackagePriceBillingCycleConfigurationDurationUnit] `json:"duration_unit,required"`
}

func (r PlanExternalPlanIDVersionNewParamsReplacePricesPriceNewPlanMaxGroupTieredPackagePriceBillingCycleConfiguration) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// The unit of billing period duration.
type PlanExternalPlanIDVersionNewParamsReplacePricesPriceNewPlanMaxGroupTieredPackagePriceBillingCycleConfigurationDurationUnit string

const (
	PlanExternalPlanIDVersionNewParamsReplacePricesPriceNewPlanMaxGroupTieredPackagePriceBillingCycleConfigurationDurationUnitDay   PlanExternalPlanIDVersionNewParamsReplacePricesPriceNewPlanMaxGroupTieredPackagePriceBillingCycleConfigurationDurationUnit = "day"
	PlanExternalPlanIDVersionNewParamsReplacePricesPriceNewPlanMaxGroupTieredPackagePriceBillingCycleConfigurationDurationUnitMonth PlanExternalPlanIDVersionNewParamsReplacePricesPriceNewPlanMaxGroupTieredPackagePriceBillingCycleConfigurationDurationUnit = "month"
)

func (r PlanExternalPlanIDVersionNewParamsReplacePricesPriceNewPlanMaxGroupTieredPackagePriceBillingCycleConfigurationDurationUnit) IsKnown() bool {
	switch r {
	case PlanExternalPlanIDVersionNewParamsReplacePricesPriceNewPlanMaxGroupTieredPackagePriceBillingCycleConfigurationDurationUnitDay, PlanExternalPlanIDVersionNewParamsReplacePricesPriceNewPlanMaxGroupTieredPackagePriceBillingCycleConfigurationDurationUnitMonth:
		return true
	}
	return false
}

// For dimensional price: specifies a price group and dimension values
type PlanExternalPlanIDVersionNewParamsReplacePricesPriceNewPlanMaxGroupTieredPackagePriceDimensionalPriceConfiguration struct {
	// The list of dimension values matching (in order) the dimensions of the price
	// group
	DimensionValues param.Field[[]string] `json:"dimension_values,required"`
	// The id of the dimensional price group to include this price in
	DimensionalPriceGroupID param.Field[string] `json:"dimensional_price_group_id"`
	// The external id of the dimensional price group to include this price in
	ExternalDimensionalPriceGroupID param.Field[string] `json:"external_dimensional_price_group_id"`
}

func (r PlanExternalPlanIDVersionNewParamsReplacePricesPriceNewPlanMaxGroupTieredPackagePriceDimensionalPriceConfiguration) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Within each billing cycle, specifies the cadence at which invoices are produced.
// If unspecified, a single invoice is produced per billing cycle.
type PlanExternalPlanIDVersionNewParamsReplacePricesPriceNewPlanMaxGroupTieredPackagePriceInvoicingCycleConfiguration struct {
	// The duration of the billing period.
	Duration param.Field[int64] `json:"duration,required"`
	// The unit of billing period duration.
	DurationUnit param.Field[PlanExternalPlanIDVersionNewParamsReplacePricesPriceNewPlanMaxGroupTieredPackagePriceInvoicingCycleConfigurationDurationUnit] `json:"duration_unit,required"`
}

func (r PlanExternalPlanIDVersionNewParamsReplacePricesPriceNewPlanMaxGroupTieredPackagePriceInvoicingCycleConfiguration) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// The unit of billing period duration.
type PlanExternalPlanIDVersionNewParamsReplacePricesPriceNewPlanMaxGroupTieredPackagePriceInvoicingCycleConfigurationDurationUnit string

const (
	PlanExternalPlanIDVersionNewParamsReplacePricesPriceNewPlanMaxGroupTieredPackagePriceInvoicingCycleConfigurationDurationUnitDay   PlanExternalPlanIDVersionNewParamsReplacePricesPriceNewPlanMaxGroupTieredPackagePriceInvoicingCycleConfigurationDurationUnit = "day"
	PlanExternalPlanIDVersionNewParamsReplacePricesPriceNewPlanMaxGroupTieredPackagePriceInvoicingCycleConfigurationDurationUnitMonth PlanExternalPlanIDVersionNewParamsReplacePricesPriceNewPlanMaxGroupTieredPackagePriceInvoicingCycleConfigurationDurationUnit = "month"
)

func (r PlanExternalPlanIDVersionNewParamsReplacePricesPriceNewPlanMaxGroupTieredPackagePriceInvoicingCycleConfigurationDurationUnit) IsKnown() bool {
	switch r {
	case PlanExternalPlanIDVersionNewParamsReplacePricesPriceNewPlanMaxGroupTieredPackagePriceInvoicingCycleConfigurationDurationUnitDay, PlanExternalPlanIDVersionNewParamsReplacePricesPriceNewPlanMaxGroupTieredPackagePriceInvoicingCycleConfigurationDurationUnitMonth:
		return true
	}
	return false
}

type PlanExternalPlanIDVersionNewParamsReplacePricesPriceNewPlanScalableMatrixWithUnitPricingPrice struct {
	// The cadence to bill for this price on.
	Cadence param.Field[PlanExternalPlanIDVersionNewParamsReplacePricesPriceNewPlanScalableMatrixWithUnitPricingPriceCadence] `json:"cadence,required"`
	// The id of the item the price will be associated with.
	ItemID    param.Field[string]                                                                                                 `json:"item_id,required"`
	ModelType param.Field[PlanExternalPlanIDVersionNewParamsReplacePricesPriceNewPlanScalableMatrixWithUnitPricingPriceModelType] `json:"model_type,required"`
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
	BillingCycleConfiguration param.Field[PlanExternalPlanIDVersionNewParamsReplacePricesPriceNewPlanScalableMatrixWithUnitPricingPriceBillingCycleConfiguration] `json:"billing_cycle_configuration"`
	// The per unit conversion rate of the price currency to the invoicing currency.
	ConversionRate param.Field[float64] `json:"conversion_rate"`
	// An ISO 4217 currency string, or custom pricing unit identifier, in which this
	// price is billed.
	Currency param.Field[string] `json:"currency"`
	// For dimensional price: specifies a price group and dimension values
	DimensionalPriceConfiguration param.Field[PlanExternalPlanIDVersionNewParamsReplacePricesPriceNewPlanScalableMatrixWithUnitPricingPriceDimensionalPriceConfiguration] `json:"dimensional_price_configuration"`
	// An alias for the price.
	ExternalPriceID param.Field[string] `json:"external_price_id"`
	// If the Price represents a fixed cost, this represents the quantity of units
	// applied.
	FixedPriceQuantity param.Field[float64] `json:"fixed_price_quantity"`
	// The property used to group this price on an invoice
	InvoiceGroupingKey param.Field[string] `json:"invoice_grouping_key"`
	// Within each billing cycle, specifies the cadence at which invoices are produced.
	// If unspecified, a single invoice is produced per billing cycle.
	InvoicingCycleConfiguration param.Field[PlanExternalPlanIDVersionNewParamsReplacePricesPriceNewPlanScalableMatrixWithUnitPricingPriceInvoicingCycleConfiguration] `json:"invoicing_cycle_configuration"`
	// User-specified key/value pairs for the resource. Individual keys can be removed
	// by setting the value to `null`, and the entire metadata mapping can be cleared
	// by setting `metadata` to `null`.
	Metadata param.Field[map[string]string] `json:"metadata"`
}

func (r PlanExternalPlanIDVersionNewParamsReplacePricesPriceNewPlanScalableMatrixWithUnitPricingPrice) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r PlanExternalPlanIDVersionNewParamsReplacePricesPriceNewPlanScalableMatrixWithUnitPricingPrice) implementsPlanExternalPlanIDVersionNewParamsReplacePricesPriceUnion() {
}

// The cadence to bill for this price on.
type PlanExternalPlanIDVersionNewParamsReplacePricesPriceNewPlanScalableMatrixWithUnitPricingPriceCadence string

const (
	PlanExternalPlanIDVersionNewParamsReplacePricesPriceNewPlanScalableMatrixWithUnitPricingPriceCadenceAnnual     PlanExternalPlanIDVersionNewParamsReplacePricesPriceNewPlanScalableMatrixWithUnitPricingPriceCadence = "annual"
	PlanExternalPlanIDVersionNewParamsReplacePricesPriceNewPlanScalableMatrixWithUnitPricingPriceCadenceSemiAnnual PlanExternalPlanIDVersionNewParamsReplacePricesPriceNewPlanScalableMatrixWithUnitPricingPriceCadence = "semi_annual"
	PlanExternalPlanIDVersionNewParamsReplacePricesPriceNewPlanScalableMatrixWithUnitPricingPriceCadenceMonthly    PlanExternalPlanIDVersionNewParamsReplacePricesPriceNewPlanScalableMatrixWithUnitPricingPriceCadence = "monthly"
	PlanExternalPlanIDVersionNewParamsReplacePricesPriceNewPlanScalableMatrixWithUnitPricingPriceCadenceQuarterly  PlanExternalPlanIDVersionNewParamsReplacePricesPriceNewPlanScalableMatrixWithUnitPricingPriceCadence = "quarterly"
	PlanExternalPlanIDVersionNewParamsReplacePricesPriceNewPlanScalableMatrixWithUnitPricingPriceCadenceOneTime    PlanExternalPlanIDVersionNewParamsReplacePricesPriceNewPlanScalableMatrixWithUnitPricingPriceCadence = "one_time"
	PlanExternalPlanIDVersionNewParamsReplacePricesPriceNewPlanScalableMatrixWithUnitPricingPriceCadenceCustom     PlanExternalPlanIDVersionNewParamsReplacePricesPriceNewPlanScalableMatrixWithUnitPricingPriceCadence = "custom"
)

func (r PlanExternalPlanIDVersionNewParamsReplacePricesPriceNewPlanScalableMatrixWithUnitPricingPriceCadence) IsKnown() bool {
	switch r {
	case PlanExternalPlanIDVersionNewParamsReplacePricesPriceNewPlanScalableMatrixWithUnitPricingPriceCadenceAnnual, PlanExternalPlanIDVersionNewParamsReplacePricesPriceNewPlanScalableMatrixWithUnitPricingPriceCadenceSemiAnnual, PlanExternalPlanIDVersionNewParamsReplacePricesPriceNewPlanScalableMatrixWithUnitPricingPriceCadenceMonthly, PlanExternalPlanIDVersionNewParamsReplacePricesPriceNewPlanScalableMatrixWithUnitPricingPriceCadenceQuarterly, PlanExternalPlanIDVersionNewParamsReplacePricesPriceNewPlanScalableMatrixWithUnitPricingPriceCadenceOneTime, PlanExternalPlanIDVersionNewParamsReplacePricesPriceNewPlanScalableMatrixWithUnitPricingPriceCadenceCustom:
		return true
	}
	return false
}

type PlanExternalPlanIDVersionNewParamsReplacePricesPriceNewPlanScalableMatrixWithUnitPricingPriceModelType string

const (
	PlanExternalPlanIDVersionNewParamsReplacePricesPriceNewPlanScalableMatrixWithUnitPricingPriceModelTypeScalableMatrixWithUnitPricing PlanExternalPlanIDVersionNewParamsReplacePricesPriceNewPlanScalableMatrixWithUnitPricingPriceModelType = "scalable_matrix_with_unit_pricing"
)

func (r PlanExternalPlanIDVersionNewParamsReplacePricesPriceNewPlanScalableMatrixWithUnitPricingPriceModelType) IsKnown() bool {
	switch r {
	case PlanExternalPlanIDVersionNewParamsReplacePricesPriceNewPlanScalableMatrixWithUnitPricingPriceModelTypeScalableMatrixWithUnitPricing:
		return true
	}
	return false
}

// For custom cadence: specifies the duration of the billing period in days or
// months.
type PlanExternalPlanIDVersionNewParamsReplacePricesPriceNewPlanScalableMatrixWithUnitPricingPriceBillingCycleConfiguration struct {
	// The duration of the billing period.
	Duration param.Field[int64] `json:"duration,required"`
	// The unit of billing period duration.
	DurationUnit param.Field[PlanExternalPlanIDVersionNewParamsReplacePricesPriceNewPlanScalableMatrixWithUnitPricingPriceBillingCycleConfigurationDurationUnit] `json:"duration_unit,required"`
}

func (r PlanExternalPlanIDVersionNewParamsReplacePricesPriceNewPlanScalableMatrixWithUnitPricingPriceBillingCycleConfiguration) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// The unit of billing period duration.
type PlanExternalPlanIDVersionNewParamsReplacePricesPriceNewPlanScalableMatrixWithUnitPricingPriceBillingCycleConfigurationDurationUnit string

const (
	PlanExternalPlanIDVersionNewParamsReplacePricesPriceNewPlanScalableMatrixWithUnitPricingPriceBillingCycleConfigurationDurationUnitDay   PlanExternalPlanIDVersionNewParamsReplacePricesPriceNewPlanScalableMatrixWithUnitPricingPriceBillingCycleConfigurationDurationUnit = "day"
	PlanExternalPlanIDVersionNewParamsReplacePricesPriceNewPlanScalableMatrixWithUnitPricingPriceBillingCycleConfigurationDurationUnitMonth PlanExternalPlanIDVersionNewParamsReplacePricesPriceNewPlanScalableMatrixWithUnitPricingPriceBillingCycleConfigurationDurationUnit = "month"
)

func (r PlanExternalPlanIDVersionNewParamsReplacePricesPriceNewPlanScalableMatrixWithUnitPricingPriceBillingCycleConfigurationDurationUnit) IsKnown() bool {
	switch r {
	case PlanExternalPlanIDVersionNewParamsReplacePricesPriceNewPlanScalableMatrixWithUnitPricingPriceBillingCycleConfigurationDurationUnitDay, PlanExternalPlanIDVersionNewParamsReplacePricesPriceNewPlanScalableMatrixWithUnitPricingPriceBillingCycleConfigurationDurationUnitMonth:
		return true
	}
	return false
}

// For dimensional price: specifies a price group and dimension values
type PlanExternalPlanIDVersionNewParamsReplacePricesPriceNewPlanScalableMatrixWithUnitPricingPriceDimensionalPriceConfiguration struct {
	// The list of dimension values matching (in order) the dimensions of the price
	// group
	DimensionValues param.Field[[]string] `json:"dimension_values,required"`
	// The id of the dimensional price group to include this price in
	DimensionalPriceGroupID param.Field[string] `json:"dimensional_price_group_id"`
	// The external id of the dimensional price group to include this price in
	ExternalDimensionalPriceGroupID param.Field[string] `json:"external_dimensional_price_group_id"`
}

func (r PlanExternalPlanIDVersionNewParamsReplacePricesPriceNewPlanScalableMatrixWithUnitPricingPriceDimensionalPriceConfiguration) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Within each billing cycle, specifies the cadence at which invoices are produced.
// If unspecified, a single invoice is produced per billing cycle.
type PlanExternalPlanIDVersionNewParamsReplacePricesPriceNewPlanScalableMatrixWithUnitPricingPriceInvoicingCycleConfiguration struct {
	// The duration of the billing period.
	Duration param.Field[int64] `json:"duration,required"`
	// The unit of billing period duration.
	DurationUnit param.Field[PlanExternalPlanIDVersionNewParamsReplacePricesPriceNewPlanScalableMatrixWithUnitPricingPriceInvoicingCycleConfigurationDurationUnit] `json:"duration_unit,required"`
}

func (r PlanExternalPlanIDVersionNewParamsReplacePricesPriceNewPlanScalableMatrixWithUnitPricingPriceInvoicingCycleConfiguration) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// The unit of billing period duration.
type PlanExternalPlanIDVersionNewParamsReplacePricesPriceNewPlanScalableMatrixWithUnitPricingPriceInvoicingCycleConfigurationDurationUnit string

const (
	PlanExternalPlanIDVersionNewParamsReplacePricesPriceNewPlanScalableMatrixWithUnitPricingPriceInvoicingCycleConfigurationDurationUnitDay   PlanExternalPlanIDVersionNewParamsReplacePricesPriceNewPlanScalableMatrixWithUnitPricingPriceInvoicingCycleConfigurationDurationUnit = "day"
	PlanExternalPlanIDVersionNewParamsReplacePricesPriceNewPlanScalableMatrixWithUnitPricingPriceInvoicingCycleConfigurationDurationUnitMonth PlanExternalPlanIDVersionNewParamsReplacePricesPriceNewPlanScalableMatrixWithUnitPricingPriceInvoicingCycleConfigurationDurationUnit = "month"
)

func (r PlanExternalPlanIDVersionNewParamsReplacePricesPriceNewPlanScalableMatrixWithUnitPricingPriceInvoicingCycleConfigurationDurationUnit) IsKnown() bool {
	switch r {
	case PlanExternalPlanIDVersionNewParamsReplacePricesPriceNewPlanScalableMatrixWithUnitPricingPriceInvoicingCycleConfigurationDurationUnitDay, PlanExternalPlanIDVersionNewParamsReplacePricesPriceNewPlanScalableMatrixWithUnitPricingPriceInvoicingCycleConfigurationDurationUnitMonth:
		return true
	}
	return false
}

type PlanExternalPlanIDVersionNewParamsReplacePricesPriceNewPlanScalableMatrixWithTieredPricingPrice struct {
	// The cadence to bill for this price on.
	Cadence param.Field[PlanExternalPlanIDVersionNewParamsReplacePricesPriceNewPlanScalableMatrixWithTieredPricingPriceCadence] `json:"cadence,required"`
	// The id of the item the price will be associated with.
	ItemID    param.Field[string]                                                                                                   `json:"item_id,required"`
	ModelType param.Field[PlanExternalPlanIDVersionNewParamsReplacePricesPriceNewPlanScalableMatrixWithTieredPricingPriceModelType] `json:"model_type,required"`
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
	BillingCycleConfiguration param.Field[PlanExternalPlanIDVersionNewParamsReplacePricesPriceNewPlanScalableMatrixWithTieredPricingPriceBillingCycleConfiguration] `json:"billing_cycle_configuration"`
	// The per unit conversion rate of the price currency to the invoicing currency.
	ConversionRate param.Field[float64] `json:"conversion_rate"`
	// An ISO 4217 currency string, or custom pricing unit identifier, in which this
	// price is billed.
	Currency param.Field[string] `json:"currency"`
	// For dimensional price: specifies a price group and dimension values
	DimensionalPriceConfiguration param.Field[PlanExternalPlanIDVersionNewParamsReplacePricesPriceNewPlanScalableMatrixWithTieredPricingPriceDimensionalPriceConfiguration] `json:"dimensional_price_configuration"`
	// An alias for the price.
	ExternalPriceID param.Field[string] `json:"external_price_id"`
	// If the Price represents a fixed cost, this represents the quantity of units
	// applied.
	FixedPriceQuantity param.Field[float64] `json:"fixed_price_quantity"`
	// The property used to group this price on an invoice
	InvoiceGroupingKey param.Field[string] `json:"invoice_grouping_key"`
	// Within each billing cycle, specifies the cadence at which invoices are produced.
	// If unspecified, a single invoice is produced per billing cycle.
	InvoicingCycleConfiguration param.Field[PlanExternalPlanIDVersionNewParamsReplacePricesPriceNewPlanScalableMatrixWithTieredPricingPriceInvoicingCycleConfiguration] `json:"invoicing_cycle_configuration"`
	// User-specified key/value pairs for the resource. Individual keys can be removed
	// by setting the value to `null`, and the entire metadata mapping can be cleared
	// by setting `metadata` to `null`.
	Metadata param.Field[map[string]string] `json:"metadata"`
}

func (r PlanExternalPlanIDVersionNewParamsReplacePricesPriceNewPlanScalableMatrixWithTieredPricingPrice) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r PlanExternalPlanIDVersionNewParamsReplacePricesPriceNewPlanScalableMatrixWithTieredPricingPrice) implementsPlanExternalPlanIDVersionNewParamsReplacePricesPriceUnion() {
}

// The cadence to bill for this price on.
type PlanExternalPlanIDVersionNewParamsReplacePricesPriceNewPlanScalableMatrixWithTieredPricingPriceCadence string

const (
	PlanExternalPlanIDVersionNewParamsReplacePricesPriceNewPlanScalableMatrixWithTieredPricingPriceCadenceAnnual     PlanExternalPlanIDVersionNewParamsReplacePricesPriceNewPlanScalableMatrixWithTieredPricingPriceCadence = "annual"
	PlanExternalPlanIDVersionNewParamsReplacePricesPriceNewPlanScalableMatrixWithTieredPricingPriceCadenceSemiAnnual PlanExternalPlanIDVersionNewParamsReplacePricesPriceNewPlanScalableMatrixWithTieredPricingPriceCadence = "semi_annual"
	PlanExternalPlanIDVersionNewParamsReplacePricesPriceNewPlanScalableMatrixWithTieredPricingPriceCadenceMonthly    PlanExternalPlanIDVersionNewParamsReplacePricesPriceNewPlanScalableMatrixWithTieredPricingPriceCadence = "monthly"
	PlanExternalPlanIDVersionNewParamsReplacePricesPriceNewPlanScalableMatrixWithTieredPricingPriceCadenceQuarterly  PlanExternalPlanIDVersionNewParamsReplacePricesPriceNewPlanScalableMatrixWithTieredPricingPriceCadence = "quarterly"
	PlanExternalPlanIDVersionNewParamsReplacePricesPriceNewPlanScalableMatrixWithTieredPricingPriceCadenceOneTime    PlanExternalPlanIDVersionNewParamsReplacePricesPriceNewPlanScalableMatrixWithTieredPricingPriceCadence = "one_time"
	PlanExternalPlanIDVersionNewParamsReplacePricesPriceNewPlanScalableMatrixWithTieredPricingPriceCadenceCustom     PlanExternalPlanIDVersionNewParamsReplacePricesPriceNewPlanScalableMatrixWithTieredPricingPriceCadence = "custom"
)

func (r PlanExternalPlanIDVersionNewParamsReplacePricesPriceNewPlanScalableMatrixWithTieredPricingPriceCadence) IsKnown() bool {
	switch r {
	case PlanExternalPlanIDVersionNewParamsReplacePricesPriceNewPlanScalableMatrixWithTieredPricingPriceCadenceAnnual, PlanExternalPlanIDVersionNewParamsReplacePricesPriceNewPlanScalableMatrixWithTieredPricingPriceCadenceSemiAnnual, PlanExternalPlanIDVersionNewParamsReplacePricesPriceNewPlanScalableMatrixWithTieredPricingPriceCadenceMonthly, PlanExternalPlanIDVersionNewParamsReplacePricesPriceNewPlanScalableMatrixWithTieredPricingPriceCadenceQuarterly, PlanExternalPlanIDVersionNewParamsReplacePricesPriceNewPlanScalableMatrixWithTieredPricingPriceCadenceOneTime, PlanExternalPlanIDVersionNewParamsReplacePricesPriceNewPlanScalableMatrixWithTieredPricingPriceCadenceCustom:
		return true
	}
	return false
}

type PlanExternalPlanIDVersionNewParamsReplacePricesPriceNewPlanScalableMatrixWithTieredPricingPriceModelType string

const (
	PlanExternalPlanIDVersionNewParamsReplacePricesPriceNewPlanScalableMatrixWithTieredPricingPriceModelTypeScalableMatrixWithTieredPricing PlanExternalPlanIDVersionNewParamsReplacePricesPriceNewPlanScalableMatrixWithTieredPricingPriceModelType = "scalable_matrix_with_tiered_pricing"
)

func (r PlanExternalPlanIDVersionNewParamsReplacePricesPriceNewPlanScalableMatrixWithTieredPricingPriceModelType) IsKnown() bool {
	switch r {
	case PlanExternalPlanIDVersionNewParamsReplacePricesPriceNewPlanScalableMatrixWithTieredPricingPriceModelTypeScalableMatrixWithTieredPricing:
		return true
	}
	return false
}

// For custom cadence: specifies the duration of the billing period in days or
// months.
type PlanExternalPlanIDVersionNewParamsReplacePricesPriceNewPlanScalableMatrixWithTieredPricingPriceBillingCycleConfiguration struct {
	// The duration of the billing period.
	Duration param.Field[int64] `json:"duration,required"`
	// The unit of billing period duration.
	DurationUnit param.Field[PlanExternalPlanIDVersionNewParamsReplacePricesPriceNewPlanScalableMatrixWithTieredPricingPriceBillingCycleConfigurationDurationUnit] `json:"duration_unit,required"`
}

func (r PlanExternalPlanIDVersionNewParamsReplacePricesPriceNewPlanScalableMatrixWithTieredPricingPriceBillingCycleConfiguration) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// The unit of billing period duration.
type PlanExternalPlanIDVersionNewParamsReplacePricesPriceNewPlanScalableMatrixWithTieredPricingPriceBillingCycleConfigurationDurationUnit string

const (
	PlanExternalPlanIDVersionNewParamsReplacePricesPriceNewPlanScalableMatrixWithTieredPricingPriceBillingCycleConfigurationDurationUnitDay   PlanExternalPlanIDVersionNewParamsReplacePricesPriceNewPlanScalableMatrixWithTieredPricingPriceBillingCycleConfigurationDurationUnit = "day"
	PlanExternalPlanIDVersionNewParamsReplacePricesPriceNewPlanScalableMatrixWithTieredPricingPriceBillingCycleConfigurationDurationUnitMonth PlanExternalPlanIDVersionNewParamsReplacePricesPriceNewPlanScalableMatrixWithTieredPricingPriceBillingCycleConfigurationDurationUnit = "month"
)

func (r PlanExternalPlanIDVersionNewParamsReplacePricesPriceNewPlanScalableMatrixWithTieredPricingPriceBillingCycleConfigurationDurationUnit) IsKnown() bool {
	switch r {
	case PlanExternalPlanIDVersionNewParamsReplacePricesPriceNewPlanScalableMatrixWithTieredPricingPriceBillingCycleConfigurationDurationUnitDay, PlanExternalPlanIDVersionNewParamsReplacePricesPriceNewPlanScalableMatrixWithTieredPricingPriceBillingCycleConfigurationDurationUnitMonth:
		return true
	}
	return false
}

// For dimensional price: specifies a price group and dimension values
type PlanExternalPlanIDVersionNewParamsReplacePricesPriceNewPlanScalableMatrixWithTieredPricingPriceDimensionalPriceConfiguration struct {
	// The list of dimension values matching (in order) the dimensions of the price
	// group
	DimensionValues param.Field[[]string] `json:"dimension_values,required"`
	// The id of the dimensional price group to include this price in
	DimensionalPriceGroupID param.Field[string] `json:"dimensional_price_group_id"`
	// The external id of the dimensional price group to include this price in
	ExternalDimensionalPriceGroupID param.Field[string] `json:"external_dimensional_price_group_id"`
}

func (r PlanExternalPlanIDVersionNewParamsReplacePricesPriceNewPlanScalableMatrixWithTieredPricingPriceDimensionalPriceConfiguration) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Within each billing cycle, specifies the cadence at which invoices are produced.
// If unspecified, a single invoice is produced per billing cycle.
type PlanExternalPlanIDVersionNewParamsReplacePricesPriceNewPlanScalableMatrixWithTieredPricingPriceInvoicingCycleConfiguration struct {
	// The duration of the billing period.
	Duration param.Field[int64] `json:"duration,required"`
	// The unit of billing period duration.
	DurationUnit param.Field[PlanExternalPlanIDVersionNewParamsReplacePricesPriceNewPlanScalableMatrixWithTieredPricingPriceInvoicingCycleConfigurationDurationUnit] `json:"duration_unit,required"`
}

func (r PlanExternalPlanIDVersionNewParamsReplacePricesPriceNewPlanScalableMatrixWithTieredPricingPriceInvoicingCycleConfiguration) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// The unit of billing period duration.
type PlanExternalPlanIDVersionNewParamsReplacePricesPriceNewPlanScalableMatrixWithTieredPricingPriceInvoicingCycleConfigurationDurationUnit string

const (
	PlanExternalPlanIDVersionNewParamsReplacePricesPriceNewPlanScalableMatrixWithTieredPricingPriceInvoicingCycleConfigurationDurationUnitDay   PlanExternalPlanIDVersionNewParamsReplacePricesPriceNewPlanScalableMatrixWithTieredPricingPriceInvoicingCycleConfigurationDurationUnit = "day"
	PlanExternalPlanIDVersionNewParamsReplacePricesPriceNewPlanScalableMatrixWithTieredPricingPriceInvoicingCycleConfigurationDurationUnitMonth PlanExternalPlanIDVersionNewParamsReplacePricesPriceNewPlanScalableMatrixWithTieredPricingPriceInvoicingCycleConfigurationDurationUnit = "month"
)

func (r PlanExternalPlanIDVersionNewParamsReplacePricesPriceNewPlanScalableMatrixWithTieredPricingPriceInvoicingCycleConfigurationDurationUnit) IsKnown() bool {
	switch r {
	case PlanExternalPlanIDVersionNewParamsReplacePricesPriceNewPlanScalableMatrixWithTieredPricingPriceInvoicingCycleConfigurationDurationUnitDay, PlanExternalPlanIDVersionNewParamsReplacePricesPriceNewPlanScalableMatrixWithTieredPricingPriceInvoicingCycleConfigurationDurationUnitMonth:
		return true
	}
	return false
}

type PlanExternalPlanIDVersionNewParamsReplacePricesPriceNewPlanCumulativeGroupedBulkPrice struct {
	// The cadence to bill for this price on.
	Cadence                     param.Field[PlanExternalPlanIDVersionNewParamsReplacePricesPriceNewPlanCumulativeGroupedBulkPriceCadence] `json:"cadence,required"`
	CumulativeGroupedBulkConfig param.Field[map[string]interface{}]                                                                       `json:"cumulative_grouped_bulk_config,required"`
	// The id of the item the price will be associated with.
	ItemID    param.Field[string]                                                                                         `json:"item_id,required"`
	ModelType param.Field[PlanExternalPlanIDVersionNewParamsReplacePricesPriceNewPlanCumulativeGroupedBulkPriceModelType] `json:"model_type,required"`
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
	BillingCycleConfiguration param.Field[PlanExternalPlanIDVersionNewParamsReplacePricesPriceNewPlanCumulativeGroupedBulkPriceBillingCycleConfiguration] `json:"billing_cycle_configuration"`
	// The per unit conversion rate of the price currency to the invoicing currency.
	ConversionRate param.Field[float64] `json:"conversion_rate"`
	// An ISO 4217 currency string, or custom pricing unit identifier, in which this
	// price is billed.
	Currency param.Field[string] `json:"currency"`
	// For dimensional price: specifies a price group and dimension values
	DimensionalPriceConfiguration param.Field[PlanExternalPlanIDVersionNewParamsReplacePricesPriceNewPlanCumulativeGroupedBulkPriceDimensionalPriceConfiguration] `json:"dimensional_price_configuration"`
	// An alias for the price.
	ExternalPriceID param.Field[string] `json:"external_price_id"`
	// If the Price represents a fixed cost, this represents the quantity of units
	// applied.
	FixedPriceQuantity param.Field[float64] `json:"fixed_price_quantity"`
	// The property used to group this price on an invoice
	InvoiceGroupingKey param.Field[string] `json:"invoice_grouping_key"`
	// Within each billing cycle, specifies the cadence at which invoices are produced.
	// If unspecified, a single invoice is produced per billing cycle.
	InvoicingCycleConfiguration param.Field[PlanExternalPlanIDVersionNewParamsReplacePricesPriceNewPlanCumulativeGroupedBulkPriceInvoicingCycleConfiguration] `json:"invoicing_cycle_configuration"`
	// User-specified key/value pairs for the resource. Individual keys can be removed
	// by setting the value to `null`, and the entire metadata mapping can be cleared
	// by setting `metadata` to `null`.
	Metadata param.Field[map[string]string] `json:"metadata"`
}

func (r PlanExternalPlanIDVersionNewParamsReplacePricesPriceNewPlanCumulativeGroupedBulkPrice) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r PlanExternalPlanIDVersionNewParamsReplacePricesPriceNewPlanCumulativeGroupedBulkPrice) implementsPlanExternalPlanIDVersionNewParamsReplacePricesPriceUnion() {
}

// The cadence to bill for this price on.
type PlanExternalPlanIDVersionNewParamsReplacePricesPriceNewPlanCumulativeGroupedBulkPriceCadence string

const (
	PlanExternalPlanIDVersionNewParamsReplacePricesPriceNewPlanCumulativeGroupedBulkPriceCadenceAnnual     PlanExternalPlanIDVersionNewParamsReplacePricesPriceNewPlanCumulativeGroupedBulkPriceCadence = "annual"
	PlanExternalPlanIDVersionNewParamsReplacePricesPriceNewPlanCumulativeGroupedBulkPriceCadenceSemiAnnual PlanExternalPlanIDVersionNewParamsReplacePricesPriceNewPlanCumulativeGroupedBulkPriceCadence = "semi_annual"
	PlanExternalPlanIDVersionNewParamsReplacePricesPriceNewPlanCumulativeGroupedBulkPriceCadenceMonthly    PlanExternalPlanIDVersionNewParamsReplacePricesPriceNewPlanCumulativeGroupedBulkPriceCadence = "monthly"
	PlanExternalPlanIDVersionNewParamsReplacePricesPriceNewPlanCumulativeGroupedBulkPriceCadenceQuarterly  PlanExternalPlanIDVersionNewParamsReplacePricesPriceNewPlanCumulativeGroupedBulkPriceCadence = "quarterly"
	PlanExternalPlanIDVersionNewParamsReplacePricesPriceNewPlanCumulativeGroupedBulkPriceCadenceOneTime    PlanExternalPlanIDVersionNewParamsReplacePricesPriceNewPlanCumulativeGroupedBulkPriceCadence = "one_time"
	PlanExternalPlanIDVersionNewParamsReplacePricesPriceNewPlanCumulativeGroupedBulkPriceCadenceCustom     PlanExternalPlanIDVersionNewParamsReplacePricesPriceNewPlanCumulativeGroupedBulkPriceCadence = "custom"
)

func (r PlanExternalPlanIDVersionNewParamsReplacePricesPriceNewPlanCumulativeGroupedBulkPriceCadence) IsKnown() bool {
	switch r {
	case PlanExternalPlanIDVersionNewParamsReplacePricesPriceNewPlanCumulativeGroupedBulkPriceCadenceAnnual, PlanExternalPlanIDVersionNewParamsReplacePricesPriceNewPlanCumulativeGroupedBulkPriceCadenceSemiAnnual, PlanExternalPlanIDVersionNewParamsReplacePricesPriceNewPlanCumulativeGroupedBulkPriceCadenceMonthly, PlanExternalPlanIDVersionNewParamsReplacePricesPriceNewPlanCumulativeGroupedBulkPriceCadenceQuarterly, PlanExternalPlanIDVersionNewParamsReplacePricesPriceNewPlanCumulativeGroupedBulkPriceCadenceOneTime, PlanExternalPlanIDVersionNewParamsReplacePricesPriceNewPlanCumulativeGroupedBulkPriceCadenceCustom:
		return true
	}
	return false
}

type PlanExternalPlanIDVersionNewParamsReplacePricesPriceNewPlanCumulativeGroupedBulkPriceModelType string

const (
	PlanExternalPlanIDVersionNewParamsReplacePricesPriceNewPlanCumulativeGroupedBulkPriceModelTypeCumulativeGroupedBulk PlanExternalPlanIDVersionNewParamsReplacePricesPriceNewPlanCumulativeGroupedBulkPriceModelType = "cumulative_grouped_bulk"
)

func (r PlanExternalPlanIDVersionNewParamsReplacePricesPriceNewPlanCumulativeGroupedBulkPriceModelType) IsKnown() bool {
	switch r {
	case PlanExternalPlanIDVersionNewParamsReplacePricesPriceNewPlanCumulativeGroupedBulkPriceModelTypeCumulativeGroupedBulk:
		return true
	}
	return false
}

// For custom cadence: specifies the duration of the billing period in days or
// months.
type PlanExternalPlanIDVersionNewParamsReplacePricesPriceNewPlanCumulativeGroupedBulkPriceBillingCycleConfiguration struct {
	// The duration of the billing period.
	Duration param.Field[int64] `json:"duration,required"`
	// The unit of billing period duration.
	DurationUnit param.Field[PlanExternalPlanIDVersionNewParamsReplacePricesPriceNewPlanCumulativeGroupedBulkPriceBillingCycleConfigurationDurationUnit] `json:"duration_unit,required"`
}

func (r PlanExternalPlanIDVersionNewParamsReplacePricesPriceNewPlanCumulativeGroupedBulkPriceBillingCycleConfiguration) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// The unit of billing period duration.
type PlanExternalPlanIDVersionNewParamsReplacePricesPriceNewPlanCumulativeGroupedBulkPriceBillingCycleConfigurationDurationUnit string

const (
	PlanExternalPlanIDVersionNewParamsReplacePricesPriceNewPlanCumulativeGroupedBulkPriceBillingCycleConfigurationDurationUnitDay   PlanExternalPlanIDVersionNewParamsReplacePricesPriceNewPlanCumulativeGroupedBulkPriceBillingCycleConfigurationDurationUnit = "day"
	PlanExternalPlanIDVersionNewParamsReplacePricesPriceNewPlanCumulativeGroupedBulkPriceBillingCycleConfigurationDurationUnitMonth PlanExternalPlanIDVersionNewParamsReplacePricesPriceNewPlanCumulativeGroupedBulkPriceBillingCycleConfigurationDurationUnit = "month"
)

func (r PlanExternalPlanIDVersionNewParamsReplacePricesPriceNewPlanCumulativeGroupedBulkPriceBillingCycleConfigurationDurationUnit) IsKnown() bool {
	switch r {
	case PlanExternalPlanIDVersionNewParamsReplacePricesPriceNewPlanCumulativeGroupedBulkPriceBillingCycleConfigurationDurationUnitDay, PlanExternalPlanIDVersionNewParamsReplacePricesPriceNewPlanCumulativeGroupedBulkPriceBillingCycleConfigurationDurationUnitMonth:
		return true
	}
	return false
}

// For dimensional price: specifies a price group and dimension values
type PlanExternalPlanIDVersionNewParamsReplacePricesPriceNewPlanCumulativeGroupedBulkPriceDimensionalPriceConfiguration struct {
	// The list of dimension values matching (in order) the dimensions of the price
	// group
	DimensionValues param.Field[[]string] `json:"dimension_values,required"`
	// The id of the dimensional price group to include this price in
	DimensionalPriceGroupID param.Field[string] `json:"dimensional_price_group_id"`
	// The external id of the dimensional price group to include this price in
	ExternalDimensionalPriceGroupID param.Field[string] `json:"external_dimensional_price_group_id"`
}

func (r PlanExternalPlanIDVersionNewParamsReplacePricesPriceNewPlanCumulativeGroupedBulkPriceDimensionalPriceConfiguration) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Within each billing cycle, specifies the cadence at which invoices are produced.
// If unspecified, a single invoice is produced per billing cycle.
type PlanExternalPlanIDVersionNewParamsReplacePricesPriceNewPlanCumulativeGroupedBulkPriceInvoicingCycleConfiguration struct {
	// The duration of the billing period.
	Duration param.Field[int64] `json:"duration,required"`
	// The unit of billing period duration.
	DurationUnit param.Field[PlanExternalPlanIDVersionNewParamsReplacePricesPriceNewPlanCumulativeGroupedBulkPriceInvoicingCycleConfigurationDurationUnit] `json:"duration_unit,required"`
}

func (r PlanExternalPlanIDVersionNewParamsReplacePricesPriceNewPlanCumulativeGroupedBulkPriceInvoicingCycleConfiguration) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// The unit of billing period duration.
type PlanExternalPlanIDVersionNewParamsReplacePricesPriceNewPlanCumulativeGroupedBulkPriceInvoicingCycleConfigurationDurationUnit string

const (
	PlanExternalPlanIDVersionNewParamsReplacePricesPriceNewPlanCumulativeGroupedBulkPriceInvoicingCycleConfigurationDurationUnitDay   PlanExternalPlanIDVersionNewParamsReplacePricesPriceNewPlanCumulativeGroupedBulkPriceInvoicingCycleConfigurationDurationUnit = "day"
	PlanExternalPlanIDVersionNewParamsReplacePricesPriceNewPlanCumulativeGroupedBulkPriceInvoicingCycleConfigurationDurationUnitMonth PlanExternalPlanIDVersionNewParamsReplacePricesPriceNewPlanCumulativeGroupedBulkPriceInvoicingCycleConfigurationDurationUnit = "month"
)

func (r PlanExternalPlanIDVersionNewParamsReplacePricesPriceNewPlanCumulativeGroupedBulkPriceInvoicingCycleConfigurationDurationUnit) IsKnown() bool {
	switch r {
	case PlanExternalPlanIDVersionNewParamsReplacePricesPriceNewPlanCumulativeGroupedBulkPriceInvoicingCycleConfigurationDurationUnitDay, PlanExternalPlanIDVersionNewParamsReplacePricesPriceNewPlanCumulativeGroupedBulkPriceInvoicingCycleConfigurationDurationUnitMonth:
		return true
	}
	return false
}

type PlanExternalPlanIDVersionNewParamsReplacePricesPriceNewPlanTieredPackageWithMinimumPrice struct {
	// The cadence to bill for this price on.
	Cadence param.Field[PlanExternalPlanIDVersionNewParamsReplacePricesPriceNewPlanTieredPackageWithMinimumPriceCadence] `json:"cadence,required"`
	// The id of the item the price will be associated with.
	ItemID    param.Field[string]                                                                                            `json:"item_id,required"`
	ModelType param.Field[PlanExternalPlanIDVersionNewParamsReplacePricesPriceNewPlanTieredPackageWithMinimumPriceModelType] `json:"model_type,required"`
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
	BillingCycleConfiguration param.Field[PlanExternalPlanIDVersionNewParamsReplacePricesPriceNewPlanTieredPackageWithMinimumPriceBillingCycleConfiguration] `json:"billing_cycle_configuration"`
	// The per unit conversion rate of the price currency to the invoicing currency.
	ConversionRate param.Field[float64] `json:"conversion_rate"`
	// An ISO 4217 currency string, or custom pricing unit identifier, in which this
	// price is billed.
	Currency param.Field[string] `json:"currency"`
	// For dimensional price: specifies a price group and dimension values
	DimensionalPriceConfiguration param.Field[PlanExternalPlanIDVersionNewParamsReplacePricesPriceNewPlanTieredPackageWithMinimumPriceDimensionalPriceConfiguration] `json:"dimensional_price_configuration"`
	// An alias for the price.
	ExternalPriceID param.Field[string] `json:"external_price_id"`
	// If the Price represents a fixed cost, this represents the quantity of units
	// applied.
	FixedPriceQuantity param.Field[float64] `json:"fixed_price_quantity"`
	// The property used to group this price on an invoice
	InvoiceGroupingKey param.Field[string] `json:"invoice_grouping_key"`
	// Within each billing cycle, specifies the cadence at which invoices are produced.
	// If unspecified, a single invoice is produced per billing cycle.
	InvoicingCycleConfiguration param.Field[PlanExternalPlanIDVersionNewParamsReplacePricesPriceNewPlanTieredPackageWithMinimumPriceInvoicingCycleConfiguration] `json:"invoicing_cycle_configuration"`
	// User-specified key/value pairs for the resource. Individual keys can be removed
	// by setting the value to `null`, and the entire metadata mapping can be cleared
	// by setting `metadata` to `null`.
	Metadata param.Field[map[string]string] `json:"metadata"`
}

func (r PlanExternalPlanIDVersionNewParamsReplacePricesPriceNewPlanTieredPackageWithMinimumPrice) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r PlanExternalPlanIDVersionNewParamsReplacePricesPriceNewPlanTieredPackageWithMinimumPrice) implementsPlanExternalPlanIDVersionNewParamsReplacePricesPriceUnion() {
}

// The cadence to bill for this price on.
type PlanExternalPlanIDVersionNewParamsReplacePricesPriceNewPlanTieredPackageWithMinimumPriceCadence string

const (
	PlanExternalPlanIDVersionNewParamsReplacePricesPriceNewPlanTieredPackageWithMinimumPriceCadenceAnnual     PlanExternalPlanIDVersionNewParamsReplacePricesPriceNewPlanTieredPackageWithMinimumPriceCadence = "annual"
	PlanExternalPlanIDVersionNewParamsReplacePricesPriceNewPlanTieredPackageWithMinimumPriceCadenceSemiAnnual PlanExternalPlanIDVersionNewParamsReplacePricesPriceNewPlanTieredPackageWithMinimumPriceCadence = "semi_annual"
	PlanExternalPlanIDVersionNewParamsReplacePricesPriceNewPlanTieredPackageWithMinimumPriceCadenceMonthly    PlanExternalPlanIDVersionNewParamsReplacePricesPriceNewPlanTieredPackageWithMinimumPriceCadence = "monthly"
	PlanExternalPlanIDVersionNewParamsReplacePricesPriceNewPlanTieredPackageWithMinimumPriceCadenceQuarterly  PlanExternalPlanIDVersionNewParamsReplacePricesPriceNewPlanTieredPackageWithMinimumPriceCadence = "quarterly"
	PlanExternalPlanIDVersionNewParamsReplacePricesPriceNewPlanTieredPackageWithMinimumPriceCadenceOneTime    PlanExternalPlanIDVersionNewParamsReplacePricesPriceNewPlanTieredPackageWithMinimumPriceCadence = "one_time"
	PlanExternalPlanIDVersionNewParamsReplacePricesPriceNewPlanTieredPackageWithMinimumPriceCadenceCustom     PlanExternalPlanIDVersionNewParamsReplacePricesPriceNewPlanTieredPackageWithMinimumPriceCadence = "custom"
)

func (r PlanExternalPlanIDVersionNewParamsReplacePricesPriceNewPlanTieredPackageWithMinimumPriceCadence) IsKnown() bool {
	switch r {
	case PlanExternalPlanIDVersionNewParamsReplacePricesPriceNewPlanTieredPackageWithMinimumPriceCadenceAnnual, PlanExternalPlanIDVersionNewParamsReplacePricesPriceNewPlanTieredPackageWithMinimumPriceCadenceSemiAnnual, PlanExternalPlanIDVersionNewParamsReplacePricesPriceNewPlanTieredPackageWithMinimumPriceCadenceMonthly, PlanExternalPlanIDVersionNewParamsReplacePricesPriceNewPlanTieredPackageWithMinimumPriceCadenceQuarterly, PlanExternalPlanIDVersionNewParamsReplacePricesPriceNewPlanTieredPackageWithMinimumPriceCadenceOneTime, PlanExternalPlanIDVersionNewParamsReplacePricesPriceNewPlanTieredPackageWithMinimumPriceCadenceCustom:
		return true
	}
	return false
}

type PlanExternalPlanIDVersionNewParamsReplacePricesPriceNewPlanTieredPackageWithMinimumPriceModelType string

const (
	PlanExternalPlanIDVersionNewParamsReplacePricesPriceNewPlanTieredPackageWithMinimumPriceModelTypeTieredPackageWithMinimum PlanExternalPlanIDVersionNewParamsReplacePricesPriceNewPlanTieredPackageWithMinimumPriceModelType = "tiered_package_with_minimum"
)

func (r PlanExternalPlanIDVersionNewParamsReplacePricesPriceNewPlanTieredPackageWithMinimumPriceModelType) IsKnown() bool {
	switch r {
	case PlanExternalPlanIDVersionNewParamsReplacePricesPriceNewPlanTieredPackageWithMinimumPriceModelTypeTieredPackageWithMinimum:
		return true
	}
	return false
}

// For custom cadence: specifies the duration of the billing period in days or
// months.
type PlanExternalPlanIDVersionNewParamsReplacePricesPriceNewPlanTieredPackageWithMinimumPriceBillingCycleConfiguration struct {
	// The duration of the billing period.
	Duration param.Field[int64] `json:"duration,required"`
	// The unit of billing period duration.
	DurationUnit param.Field[PlanExternalPlanIDVersionNewParamsReplacePricesPriceNewPlanTieredPackageWithMinimumPriceBillingCycleConfigurationDurationUnit] `json:"duration_unit,required"`
}

func (r PlanExternalPlanIDVersionNewParamsReplacePricesPriceNewPlanTieredPackageWithMinimumPriceBillingCycleConfiguration) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// The unit of billing period duration.
type PlanExternalPlanIDVersionNewParamsReplacePricesPriceNewPlanTieredPackageWithMinimumPriceBillingCycleConfigurationDurationUnit string

const (
	PlanExternalPlanIDVersionNewParamsReplacePricesPriceNewPlanTieredPackageWithMinimumPriceBillingCycleConfigurationDurationUnitDay   PlanExternalPlanIDVersionNewParamsReplacePricesPriceNewPlanTieredPackageWithMinimumPriceBillingCycleConfigurationDurationUnit = "day"
	PlanExternalPlanIDVersionNewParamsReplacePricesPriceNewPlanTieredPackageWithMinimumPriceBillingCycleConfigurationDurationUnitMonth PlanExternalPlanIDVersionNewParamsReplacePricesPriceNewPlanTieredPackageWithMinimumPriceBillingCycleConfigurationDurationUnit = "month"
)

func (r PlanExternalPlanIDVersionNewParamsReplacePricesPriceNewPlanTieredPackageWithMinimumPriceBillingCycleConfigurationDurationUnit) IsKnown() bool {
	switch r {
	case PlanExternalPlanIDVersionNewParamsReplacePricesPriceNewPlanTieredPackageWithMinimumPriceBillingCycleConfigurationDurationUnitDay, PlanExternalPlanIDVersionNewParamsReplacePricesPriceNewPlanTieredPackageWithMinimumPriceBillingCycleConfigurationDurationUnitMonth:
		return true
	}
	return false
}

// For dimensional price: specifies a price group and dimension values
type PlanExternalPlanIDVersionNewParamsReplacePricesPriceNewPlanTieredPackageWithMinimumPriceDimensionalPriceConfiguration struct {
	// The list of dimension values matching (in order) the dimensions of the price
	// group
	DimensionValues param.Field[[]string] `json:"dimension_values,required"`
	// The id of the dimensional price group to include this price in
	DimensionalPriceGroupID param.Field[string] `json:"dimensional_price_group_id"`
	// The external id of the dimensional price group to include this price in
	ExternalDimensionalPriceGroupID param.Field[string] `json:"external_dimensional_price_group_id"`
}

func (r PlanExternalPlanIDVersionNewParamsReplacePricesPriceNewPlanTieredPackageWithMinimumPriceDimensionalPriceConfiguration) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Within each billing cycle, specifies the cadence at which invoices are produced.
// If unspecified, a single invoice is produced per billing cycle.
type PlanExternalPlanIDVersionNewParamsReplacePricesPriceNewPlanTieredPackageWithMinimumPriceInvoicingCycleConfiguration struct {
	// The duration of the billing period.
	Duration param.Field[int64] `json:"duration,required"`
	// The unit of billing period duration.
	DurationUnit param.Field[PlanExternalPlanIDVersionNewParamsReplacePricesPriceNewPlanTieredPackageWithMinimumPriceInvoicingCycleConfigurationDurationUnit] `json:"duration_unit,required"`
}

func (r PlanExternalPlanIDVersionNewParamsReplacePricesPriceNewPlanTieredPackageWithMinimumPriceInvoicingCycleConfiguration) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// The unit of billing period duration.
type PlanExternalPlanIDVersionNewParamsReplacePricesPriceNewPlanTieredPackageWithMinimumPriceInvoicingCycleConfigurationDurationUnit string

const (
	PlanExternalPlanIDVersionNewParamsReplacePricesPriceNewPlanTieredPackageWithMinimumPriceInvoicingCycleConfigurationDurationUnitDay   PlanExternalPlanIDVersionNewParamsReplacePricesPriceNewPlanTieredPackageWithMinimumPriceInvoicingCycleConfigurationDurationUnit = "day"
	PlanExternalPlanIDVersionNewParamsReplacePricesPriceNewPlanTieredPackageWithMinimumPriceInvoicingCycleConfigurationDurationUnitMonth PlanExternalPlanIDVersionNewParamsReplacePricesPriceNewPlanTieredPackageWithMinimumPriceInvoicingCycleConfigurationDurationUnit = "month"
)

func (r PlanExternalPlanIDVersionNewParamsReplacePricesPriceNewPlanTieredPackageWithMinimumPriceInvoicingCycleConfigurationDurationUnit) IsKnown() bool {
	switch r {
	case PlanExternalPlanIDVersionNewParamsReplacePricesPriceNewPlanTieredPackageWithMinimumPriceInvoicingCycleConfigurationDurationUnitDay, PlanExternalPlanIDVersionNewParamsReplacePricesPriceNewPlanTieredPackageWithMinimumPriceInvoicingCycleConfigurationDurationUnitMonth:
		return true
	}
	return false
}

type PlanExternalPlanIDVersionNewParamsReplacePricesPriceNewPlanMatrixWithAllocationPrice struct {
	// The cadence to bill for this price on.
	Cadence param.Field[PlanExternalPlanIDVersionNewParamsReplacePricesPriceNewPlanMatrixWithAllocationPriceCadence] `json:"cadence,required"`
	// The id of the item the price will be associated with.
	ItemID                     param.Field[string]                                                                                                         `json:"item_id,required"`
	MatrixWithAllocationConfig param.Field[PlanExternalPlanIDVersionNewParamsReplacePricesPriceNewPlanMatrixWithAllocationPriceMatrixWithAllocationConfig] `json:"matrix_with_allocation_config,required"`
	ModelType                  param.Field[PlanExternalPlanIDVersionNewParamsReplacePricesPriceNewPlanMatrixWithAllocationPriceModelType]                  `json:"model_type,required"`
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
	BillingCycleConfiguration param.Field[PlanExternalPlanIDVersionNewParamsReplacePricesPriceNewPlanMatrixWithAllocationPriceBillingCycleConfiguration] `json:"billing_cycle_configuration"`
	// The per unit conversion rate of the price currency to the invoicing currency.
	ConversionRate param.Field[float64] `json:"conversion_rate"`
	// An ISO 4217 currency string, or custom pricing unit identifier, in which this
	// price is billed.
	Currency param.Field[string] `json:"currency"`
	// For dimensional price: specifies a price group and dimension values
	DimensionalPriceConfiguration param.Field[PlanExternalPlanIDVersionNewParamsReplacePricesPriceNewPlanMatrixWithAllocationPriceDimensionalPriceConfiguration] `json:"dimensional_price_configuration"`
	// An alias for the price.
	ExternalPriceID param.Field[string] `json:"external_price_id"`
	// If the Price represents a fixed cost, this represents the quantity of units
	// applied.
	FixedPriceQuantity param.Field[float64] `json:"fixed_price_quantity"`
	// The property used to group this price on an invoice
	InvoiceGroupingKey param.Field[string] `json:"invoice_grouping_key"`
	// Within each billing cycle, specifies the cadence at which invoices are produced.
	// If unspecified, a single invoice is produced per billing cycle.
	InvoicingCycleConfiguration param.Field[PlanExternalPlanIDVersionNewParamsReplacePricesPriceNewPlanMatrixWithAllocationPriceInvoicingCycleConfiguration] `json:"invoicing_cycle_configuration"`
	// User-specified key/value pairs for the resource. Individual keys can be removed
	// by setting the value to `null`, and the entire metadata mapping can be cleared
	// by setting `metadata` to `null`.
	Metadata param.Field[map[string]string] `json:"metadata"`
}

func (r PlanExternalPlanIDVersionNewParamsReplacePricesPriceNewPlanMatrixWithAllocationPrice) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r PlanExternalPlanIDVersionNewParamsReplacePricesPriceNewPlanMatrixWithAllocationPrice) implementsPlanExternalPlanIDVersionNewParamsReplacePricesPriceUnion() {
}

// The cadence to bill for this price on.
type PlanExternalPlanIDVersionNewParamsReplacePricesPriceNewPlanMatrixWithAllocationPriceCadence string

const (
	PlanExternalPlanIDVersionNewParamsReplacePricesPriceNewPlanMatrixWithAllocationPriceCadenceAnnual     PlanExternalPlanIDVersionNewParamsReplacePricesPriceNewPlanMatrixWithAllocationPriceCadence = "annual"
	PlanExternalPlanIDVersionNewParamsReplacePricesPriceNewPlanMatrixWithAllocationPriceCadenceSemiAnnual PlanExternalPlanIDVersionNewParamsReplacePricesPriceNewPlanMatrixWithAllocationPriceCadence = "semi_annual"
	PlanExternalPlanIDVersionNewParamsReplacePricesPriceNewPlanMatrixWithAllocationPriceCadenceMonthly    PlanExternalPlanIDVersionNewParamsReplacePricesPriceNewPlanMatrixWithAllocationPriceCadence = "monthly"
	PlanExternalPlanIDVersionNewParamsReplacePricesPriceNewPlanMatrixWithAllocationPriceCadenceQuarterly  PlanExternalPlanIDVersionNewParamsReplacePricesPriceNewPlanMatrixWithAllocationPriceCadence = "quarterly"
	PlanExternalPlanIDVersionNewParamsReplacePricesPriceNewPlanMatrixWithAllocationPriceCadenceOneTime    PlanExternalPlanIDVersionNewParamsReplacePricesPriceNewPlanMatrixWithAllocationPriceCadence = "one_time"
	PlanExternalPlanIDVersionNewParamsReplacePricesPriceNewPlanMatrixWithAllocationPriceCadenceCustom     PlanExternalPlanIDVersionNewParamsReplacePricesPriceNewPlanMatrixWithAllocationPriceCadence = "custom"
)

func (r PlanExternalPlanIDVersionNewParamsReplacePricesPriceNewPlanMatrixWithAllocationPriceCadence) IsKnown() bool {
	switch r {
	case PlanExternalPlanIDVersionNewParamsReplacePricesPriceNewPlanMatrixWithAllocationPriceCadenceAnnual, PlanExternalPlanIDVersionNewParamsReplacePricesPriceNewPlanMatrixWithAllocationPriceCadenceSemiAnnual, PlanExternalPlanIDVersionNewParamsReplacePricesPriceNewPlanMatrixWithAllocationPriceCadenceMonthly, PlanExternalPlanIDVersionNewParamsReplacePricesPriceNewPlanMatrixWithAllocationPriceCadenceQuarterly, PlanExternalPlanIDVersionNewParamsReplacePricesPriceNewPlanMatrixWithAllocationPriceCadenceOneTime, PlanExternalPlanIDVersionNewParamsReplacePricesPriceNewPlanMatrixWithAllocationPriceCadenceCustom:
		return true
	}
	return false
}

type PlanExternalPlanIDVersionNewParamsReplacePricesPriceNewPlanMatrixWithAllocationPriceMatrixWithAllocationConfig struct {
	// Allocation to be used to calculate the price
	Allocation param.Field[float64] `json:"allocation,required"`
	// Default per unit rate for any usage not bucketed into a specified matrix_value
	DefaultUnitAmount param.Field[string] `json:"default_unit_amount,required"`
	// One or two event property values to evaluate matrix groups by
	Dimensions param.Field[[]string] `json:"dimensions,required"`
	// Matrix values for specified matrix grouping keys
	MatrixValues param.Field[[]PlanExternalPlanIDVersionNewParamsReplacePricesPriceNewPlanMatrixWithAllocationPriceMatrixWithAllocationConfigMatrixValue] `json:"matrix_values,required"`
}

func (r PlanExternalPlanIDVersionNewParamsReplacePricesPriceNewPlanMatrixWithAllocationPriceMatrixWithAllocationConfig) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type PlanExternalPlanIDVersionNewParamsReplacePricesPriceNewPlanMatrixWithAllocationPriceMatrixWithAllocationConfigMatrixValue struct {
	// One or two matrix keys to filter usage to this Matrix value by. For example,
	// ["region", "tier"] could be used to filter cloud usage by a cloud region and an
	// instance tier.
	DimensionValues param.Field[[]string] `json:"dimension_values,required"`
	// Unit price for the specified dimension_values
	UnitAmount param.Field[string] `json:"unit_amount,required"`
}

func (r PlanExternalPlanIDVersionNewParamsReplacePricesPriceNewPlanMatrixWithAllocationPriceMatrixWithAllocationConfigMatrixValue) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type PlanExternalPlanIDVersionNewParamsReplacePricesPriceNewPlanMatrixWithAllocationPriceModelType string

const (
	PlanExternalPlanIDVersionNewParamsReplacePricesPriceNewPlanMatrixWithAllocationPriceModelTypeMatrixWithAllocation PlanExternalPlanIDVersionNewParamsReplacePricesPriceNewPlanMatrixWithAllocationPriceModelType = "matrix_with_allocation"
)

func (r PlanExternalPlanIDVersionNewParamsReplacePricesPriceNewPlanMatrixWithAllocationPriceModelType) IsKnown() bool {
	switch r {
	case PlanExternalPlanIDVersionNewParamsReplacePricesPriceNewPlanMatrixWithAllocationPriceModelTypeMatrixWithAllocation:
		return true
	}
	return false
}

// For custom cadence: specifies the duration of the billing period in days or
// months.
type PlanExternalPlanIDVersionNewParamsReplacePricesPriceNewPlanMatrixWithAllocationPriceBillingCycleConfiguration struct {
	// The duration of the billing period.
	Duration param.Field[int64] `json:"duration,required"`
	// The unit of billing period duration.
	DurationUnit param.Field[PlanExternalPlanIDVersionNewParamsReplacePricesPriceNewPlanMatrixWithAllocationPriceBillingCycleConfigurationDurationUnit] `json:"duration_unit,required"`
}

func (r PlanExternalPlanIDVersionNewParamsReplacePricesPriceNewPlanMatrixWithAllocationPriceBillingCycleConfiguration) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// The unit of billing period duration.
type PlanExternalPlanIDVersionNewParamsReplacePricesPriceNewPlanMatrixWithAllocationPriceBillingCycleConfigurationDurationUnit string

const (
	PlanExternalPlanIDVersionNewParamsReplacePricesPriceNewPlanMatrixWithAllocationPriceBillingCycleConfigurationDurationUnitDay   PlanExternalPlanIDVersionNewParamsReplacePricesPriceNewPlanMatrixWithAllocationPriceBillingCycleConfigurationDurationUnit = "day"
	PlanExternalPlanIDVersionNewParamsReplacePricesPriceNewPlanMatrixWithAllocationPriceBillingCycleConfigurationDurationUnitMonth PlanExternalPlanIDVersionNewParamsReplacePricesPriceNewPlanMatrixWithAllocationPriceBillingCycleConfigurationDurationUnit = "month"
)

func (r PlanExternalPlanIDVersionNewParamsReplacePricesPriceNewPlanMatrixWithAllocationPriceBillingCycleConfigurationDurationUnit) IsKnown() bool {
	switch r {
	case PlanExternalPlanIDVersionNewParamsReplacePricesPriceNewPlanMatrixWithAllocationPriceBillingCycleConfigurationDurationUnitDay, PlanExternalPlanIDVersionNewParamsReplacePricesPriceNewPlanMatrixWithAllocationPriceBillingCycleConfigurationDurationUnitMonth:
		return true
	}
	return false
}

// For dimensional price: specifies a price group and dimension values
type PlanExternalPlanIDVersionNewParamsReplacePricesPriceNewPlanMatrixWithAllocationPriceDimensionalPriceConfiguration struct {
	// The list of dimension values matching (in order) the dimensions of the price
	// group
	DimensionValues param.Field[[]string] `json:"dimension_values,required"`
	// The id of the dimensional price group to include this price in
	DimensionalPriceGroupID param.Field[string] `json:"dimensional_price_group_id"`
	// The external id of the dimensional price group to include this price in
	ExternalDimensionalPriceGroupID param.Field[string] `json:"external_dimensional_price_group_id"`
}

func (r PlanExternalPlanIDVersionNewParamsReplacePricesPriceNewPlanMatrixWithAllocationPriceDimensionalPriceConfiguration) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Within each billing cycle, specifies the cadence at which invoices are produced.
// If unspecified, a single invoice is produced per billing cycle.
type PlanExternalPlanIDVersionNewParamsReplacePricesPriceNewPlanMatrixWithAllocationPriceInvoicingCycleConfiguration struct {
	// The duration of the billing period.
	Duration param.Field[int64] `json:"duration,required"`
	// The unit of billing period duration.
	DurationUnit param.Field[PlanExternalPlanIDVersionNewParamsReplacePricesPriceNewPlanMatrixWithAllocationPriceInvoicingCycleConfigurationDurationUnit] `json:"duration_unit,required"`
}

func (r PlanExternalPlanIDVersionNewParamsReplacePricesPriceNewPlanMatrixWithAllocationPriceInvoicingCycleConfiguration) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// The unit of billing period duration.
type PlanExternalPlanIDVersionNewParamsReplacePricesPriceNewPlanMatrixWithAllocationPriceInvoicingCycleConfigurationDurationUnit string

const (
	PlanExternalPlanIDVersionNewParamsReplacePricesPriceNewPlanMatrixWithAllocationPriceInvoicingCycleConfigurationDurationUnitDay   PlanExternalPlanIDVersionNewParamsReplacePricesPriceNewPlanMatrixWithAllocationPriceInvoicingCycleConfigurationDurationUnit = "day"
	PlanExternalPlanIDVersionNewParamsReplacePricesPriceNewPlanMatrixWithAllocationPriceInvoicingCycleConfigurationDurationUnitMonth PlanExternalPlanIDVersionNewParamsReplacePricesPriceNewPlanMatrixWithAllocationPriceInvoicingCycleConfigurationDurationUnit = "month"
)

func (r PlanExternalPlanIDVersionNewParamsReplacePricesPriceNewPlanMatrixWithAllocationPriceInvoicingCycleConfigurationDurationUnit) IsKnown() bool {
	switch r {
	case PlanExternalPlanIDVersionNewParamsReplacePricesPriceNewPlanMatrixWithAllocationPriceInvoicingCycleConfigurationDurationUnitDay, PlanExternalPlanIDVersionNewParamsReplacePricesPriceNewPlanMatrixWithAllocationPriceInvoicingCycleConfigurationDurationUnitMonth:
		return true
	}
	return false
}

type PlanExternalPlanIDVersionNewParamsReplacePricesPriceNewPlanGroupedTieredPrice struct {
	// The cadence to bill for this price on.
	Cadence             param.Field[PlanExternalPlanIDVersionNewParamsReplacePricesPriceNewPlanGroupedTieredPriceCadence] `json:"cadence,required"`
	GroupedTieredConfig param.Field[map[string]interface{}]                                                               `json:"grouped_tiered_config,required"`
	// The id of the item the price will be associated with.
	ItemID    param.Field[string]                                                                                 `json:"item_id,required"`
	ModelType param.Field[PlanExternalPlanIDVersionNewParamsReplacePricesPriceNewPlanGroupedTieredPriceModelType] `json:"model_type,required"`
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
	BillingCycleConfiguration param.Field[PlanExternalPlanIDVersionNewParamsReplacePricesPriceNewPlanGroupedTieredPriceBillingCycleConfiguration] `json:"billing_cycle_configuration"`
	// The per unit conversion rate of the price currency to the invoicing currency.
	ConversionRate param.Field[float64] `json:"conversion_rate"`
	// An ISO 4217 currency string, or custom pricing unit identifier, in which this
	// price is billed.
	Currency param.Field[string] `json:"currency"`
	// For dimensional price: specifies a price group and dimension values
	DimensionalPriceConfiguration param.Field[PlanExternalPlanIDVersionNewParamsReplacePricesPriceNewPlanGroupedTieredPriceDimensionalPriceConfiguration] `json:"dimensional_price_configuration"`
	// An alias for the price.
	ExternalPriceID param.Field[string] `json:"external_price_id"`
	// If the Price represents a fixed cost, this represents the quantity of units
	// applied.
	FixedPriceQuantity param.Field[float64] `json:"fixed_price_quantity"`
	// The property used to group this price on an invoice
	InvoiceGroupingKey param.Field[string] `json:"invoice_grouping_key"`
	// Within each billing cycle, specifies the cadence at which invoices are produced.
	// If unspecified, a single invoice is produced per billing cycle.
	InvoicingCycleConfiguration param.Field[PlanExternalPlanIDVersionNewParamsReplacePricesPriceNewPlanGroupedTieredPriceInvoicingCycleConfiguration] `json:"invoicing_cycle_configuration"`
	// User-specified key/value pairs for the resource. Individual keys can be removed
	// by setting the value to `null`, and the entire metadata mapping can be cleared
	// by setting `metadata` to `null`.
	Metadata param.Field[map[string]string] `json:"metadata"`
}

func (r PlanExternalPlanIDVersionNewParamsReplacePricesPriceNewPlanGroupedTieredPrice) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r PlanExternalPlanIDVersionNewParamsReplacePricesPriceNewPlanGroupedTieredPrice) implementsPlanExternalPlanIDVersionNewParamsReplacePricesPriceUnion() {
}

// The cadence to bill for this price on.
type PlanExternalPlanIDVersionNewParamsReplacePricesPriceNewPlanGroupedTieredPriceCadence string

const (
	PlanExternalPlanIDVersionNewParamsReplacePricesPriceNewPlanGroupedTieredPriceCadenceAnnual     PlanExternalPlanIDVersionNewParamsReplacePricesPriceNewPlanGroupedTieredPriceCadence = "annual"
	PlanExternalPlanIDVersionNewParamsReplacePricesPriceNewPlanGroupedTieredPriceCadenceSemiAnnual PlanExternalPlanIDVersionNewParamsReplacePricesPriceNewPlanGroupedTieredPriceCadence = "semi_annual"
	PlanExternalPlanIDVersionNewParamsReplacePricesPriceNewPlanGroupedTieredPriceCadenceMonthly    PlanExternalPlanIDVersionNewParamsReplacePricesPriceNewPlanGroupedTieredPriceCadence = "monthly"
	PlanExternalPlanIDVersionNewParamsReplacePricesPriceNewPlanGroupedTieredPriceCadenceQuarterly  PlanExternalPlanIDVersionNewParamsReplacePricesPriceNewPlanGroupedTieredPriceCadence = "quarterly"
	PlanExternalPlanIDVersionNewParamsReplacePricesPriceNewPlanGroupedTieredPriceCadenceOneTime    PlanExternalPlanIDVersionNewParamsReplacePricesPriceNewPlanGroupedTieredPriceCadence = "one_time"
	PlanExternalPlanIDVersionNewParamsReplacePricesPriceNewPlanGroupedTieredPriceCadenceCustom     PlanExternalPlanIDVersionNewParamsReplacePricesPriceNewPlanGroupedTieredPriceCadence = "custom"
)

func (r PlanExternalPlanIDVersionNewParamsReplacePricesPriceNewPlanGroupedTieredPriceCadence) IsKnown() bool {
	switch r {
	case PlanExternalPlanIDVersionNewParamsReplacePricesPriceNewPlanGroupedTieredPriceCadenceAnnual, PlanExternalPlanIDVersionNewParamsReplacePricesPriceNewPlanGroupedTieredPriceCadenceSemiAnnual, PlanExternalPlanIDVersionNewParamsReplacePricesPriceNewPlanGroupedTieredPriceCadenceMonthly, PlanExternalPlanIDVersionNewParamsReplacePricesPriceNewPlanGroupedTieredPriceCadenceQuarterly, PlanExternalPlanIDVersionNewParamsReplacePricesPriceNewPlanGroupedTieredPriceCadenceOneTime, PlanExternalPlanIDVersionNewParamsReplacePricesPriceNewPlanGroupedTieredPriceCadenceCustom:
		return true
	}
	return false
}

type PlanExternalPlanIDVersionNewParamsReplacePricesPriceNewPlanGroupedTieredPriceModelType string

const (
	PlanExternalPlanIDVersionNewParamsReplacePricesPriceNewPlanGroupedTieredPriceModelTypeGroupedTiered PlanExternalPlanIDVersionNewParamsReplacePricesPriceNewPlanGroupedTieredPriceModelType = "grouped_tiered"
)

func (r PlanExternalPlanIDVersionNewParamsReplacePricesPriceNewPlanGroupedTieredPriceModelType) IsKnown() bool {
	switch r {
	case PlanExternalPlanIDVersionNewParamsReplacePricesPriceNewPlanGroupedTieredPriceModelTypeGroupedTiered:
		return true
	}
	return false
}

// For custom cadence: specifies the duration of the billing period in days or
// months.
type PlanExternalPlanIDVersionNewParamsReplacePricesPriceNewPlanGroupedTieredPriceBillingCycleConfiguration struct {
	// The duration of the billing period.
	Duration param.Field[int64] `json:"duration,required"`
	// The unit of billing period duration.
	DurationUnit param.Field[PlanExternalPlanIDVersionNewParamsReplacePricesPriceNewPlanGroupedTieredPriceBillingCycleConfigurationDurationUnit] `json:"duration_unit,required"`
}

func (r PlanExternalPlanIDVersionNewParamsReplacePricesPriceNewPlanGroupedTieredPriceBillingCycleConfiguration) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// The unit of billing period duration.
type PlanExternalPlanIDVersionNewParamsReplacePricesPriceNewPlanGroupedTieredPriceBillingCycleConfigurationDurationUnit string

const (
	PlanExternalPlanIDVersionNewParamsReplacePricesPriceNewPlanGroupedTieredPriceBillingCycleConfigurationDurationUnitDay   PlanExternalPlanIDVersionNewParamsReplacePricesPriceNewPlanGroupedTieredPriceBillingCycleConfigurationDurationUnit = "day"
	PlanExternalPlanIDVersionNewParamsReplacePricesPriceNewPlanGroupedTieredPriceBillingCycleConfigurationDurationUnitMonth PlanExternalPlanIDVersionNewParamsReplacePricesPriceNewPlanGroupedTieredPriceBillingCycleConfigurationDurationUnit = "month"
)

func (r PlanExternalPlanIDVersionNewParamsReplacePricesPriceNewPlanGroupedTieredPriceBillingCycleConfigurationDurationUnit) IsKnown() bool {
	switch r {
	case PlanExternalPlanIDVersionNewParamsReplacePricesPriceNewPlanGroupedTieredPriceBillingCycleConfigurationDurationUnitDay, PlanExternalPlanIDVersionNewParamsReplacePricesPriceNewPlanGroupedTieredPriceBillingCycleConfigurationDurationUnitMonth:
		return true
	}
	return false
}

// For dimensional price: specifies a price group and dimension values
type PlanExternalPlanIDVersionNewParamsReplacePricesPriceNewPlanGroupedTieredPriceDimensionalPriceConfiguration struct {
	// The list of dimension values matching (in order) the dimensions of the price
	// group
	DimensionValues param.Field[[]string] `json:"dimension_values,required"`
	// The id of the dimensional price group to include this price in
	DimensionalPriceGroupID param.Field[string] `json:"dimensional_price_group_id"`
	// The external id of the dimensional price group to include this price in
	ExternalDimensionalPriceGroupID param.Field[string] `json:"external_dimensional_price_group_id"`
}

func (r PlanExternalPlanIDVersionNewParamsReplacePricesPriceNewPlanGroupedTieredPriceDimensionalPriceConfiguration) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Within each billing cycle, specifies the cadence at which invoices are produced.
// If unspecified, a single invoice is produced per billing cycle.
type PlanExternalPlanIDVersionNewParamsReplacePricesPriceNewPlanGroupedTieredPriceInvoicingCycleConfiguration struct {
	// The duration of the billing period.
	Duration param.Field[int64] `json:"duration,required"`
	// The unit of billing period duration.
	DurationUnit param.Field[PlanExternalPlanIDVersionNewParamsReplacePricesPriceNewPlanGroupedTieredPriceInvoicingCycleConfigurationDurationUnit] `json:"duration_unit,required"`
}

func (r PlanExternalPlanIDVersionNewParamsReplacePricesPriceNewPlanGroupedTieredPriceInvoicingCycleConfiguration) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// The unit of billing period duration.
type PlanExternalPlanIDVersionNewParamsReplacePricesPriceNewPlanGroupedTieredPriceInvoicingCycleConfigurationDurationUnit string

const (
	PlanExternalPlanIDVersionNewParamsReplacePricesPriceNewPlanGroupedTieredPriceInvoicingCycleConfigurationDurationUnitDay   PlanExternalPlanIDVersionNewParamsReplacePricesPriceNewPlanGroupedTieredPriceInvoicingCycleConfigurationDurationUnit = "day"
	PlanExternalPlanIDVersionNewParamsReplacePricesPriceNewPlanGroupedTieredPriceInvoicingCycleConfigurationDurationUnitMonth PlanExternalPlanIDVersionNewParamsReplacePricesPriceNewPlanGroupedTieredPriceInvoicingCycleConfigurationDurationUnit = "month"
)

func (r PlanExternalPlanIDVersionNewParamsReplacePricesPriceNewPlanGroupedTieredPriceInvoicingCycleConfigurationDurationUnit) IsKnown() bool {
	switch r {
	case PlanExternalPlanIDVersionNewParamsReplacePricesPriceNewPlanGroupedTieredPriceInvoicingCycleConfigurationDurationUnitDay, PlanExternalPlanIDVersionNewParamsReplacePricesPriceNewPlanGroupedTieredPriceInvoicingCycleConfigurationDurationUnitMonth:
		return true
	}
	return false
}

// The cadence to bill for this price on.
type PlanExternalPlanIDVersionNewParamsReplacePricesPriceCadence string

const (
	PlanExternalPlanIDVersionNewParamsReplacePricesPriceCadenceAnnual     PlanExternalPlanIDVersionNewParamsReplacePricesPriceCadence = "annual"
	PlanExternalPlanIDVersionNewParamsReplacePricesPriceCadenceSemiAnnual PlanExternalPlanIDVersionNewParamsReplacePricesPriceCadence = "semi_annual"
	PlanExternalPlanIDVersionNewParamsReplacePricesPriceCadenceMonthly    PlanExternalPlanIDVersionNewParamsReplacePricesPriceCadence = "monthly"
	PlanExternalPlanIDVersionNewParamsReplacePricesPriceCadenceQuarterly  PlanExternalPlanIDVersionNewParamsReplacePricesPriceCadence = "quarterly"
	PlanExternalPlanIDVersionNewParamsReplacePricesPriceCadenceOneTime    PlanExternalPlanIDVersionNewParamsReplacePricesPriceCadence = "one_time"
	PlanExternalPlanIDVersionNewParamsReplacePricesPriceCadenceCustom     PlanExternalPlanIDVersionNewParamsReplacePricesPriceCadence = "custom"
)

func (r PlanExternalPlanIDVersionNewParamsReplacePricesPriceCadence) IsKnown() bool {
	switch r {
	case PlanExternalPlanIDVersionNewParamsReplacePricesPriceCadenceAnnual, PlanExternalPlanIDVersionNewParamsReplacePricesPriceCadenceSemiAnnual, PlanExternalPlanIDVersionNewParamsReplacePricesPriceCadenceMonthly, PlanExternalPlanIDVersionNewParamsReplacePricesPriceCadenceQuarterly, PlanExternalPlanIDVersionNewParamsReplacePricesPriceCadenceOneTime, PlanExternalPlanIDVersionNewParamsReplacePricesPriceCadenceCustom:
		return true
	}
	return false
}

type PlanExternalPlanIDVersionNewParamsReplacePricesPriceModelType string

const (
	PlanExternalPlanIDVersionNewParamsReplacePricesPriceModelTypeUnit                            PlanExternalPlanIDVersionNewParamsReplacePricesPriceModelType = "unit"
	PlanExternalPlanIDVersionNewParamsReplacePricesPriceModelTypePackage                         PlanExternalPlanIDVersionNewParamsReplacePricesPriceModelType = "package"
	PlanExternalPlanIDVersionNewParamsReplacePricesPriceModelTypeMatrix                          PlanExternalPlanIDVersionNewParamsReplacePricesPriceModelType = "matrix"
	PlanExternalPlanIDVersionNewParamsReplacePricesPriceModelTypeTiered                          PlanExternalPlanIDVersionNewParamsReplacePricesPriceModelType = "tiered"
	PlanExternalPlanIDVersionNewParamsReplacePricesPriceModelTypeTieredBps                       PlanExternalPlanIDVersionNewParamsReplacePricesPriceModelType = "tiered_bps"
	PlanExternalPlanIDVersionNewParamsReplacePricesPriceModelTypeBps                             PlanExternalPlanIDVersionNewParamsReplacePricesPriceModelType = "bps"
	PlanExternalPlanIDVersionNewParamsReplacePricesPriceModelTypeBulkBps                         PlanExternalPlanIDVersionNewParamsReplacePricesPriceModelType = "bulk_bps"
	PlanExternalPlanIDVersionNewParamsReplacePricesPriceModelTypeBulk                            PlanExternalPlanIDVersionNewParamsReplacePricesPriceModelType = "bulk"
	PlanExternalPlanIDVersionNewParamsReplacePricesPriceModelTypeThresholdTotalAmount            PlanExternalPlanIDVersionNewParamsReplacePricesPriceModelType = "threshold_total_amount"
	PlanExternalPlanIDVersionNewParamsReplacePricesPriceModelTypeTieredPackage                   PlanExternalPlanIDVersionNewParamsReplacePricesPriceModelType = "tiered_package"
	PlanExternalPlanIDVersionNewParamsReplacePricesPriceModelTypeTieredWithMinimum               PlanExternalPlanIDVersionNewParamsReplacePricesPriceModelType = "tiered_with_minimum"
	PlanExternalPlanIDVersionNewParamsReplacePricesPriceModelTypeUnitWithPercent                 PlanExternalPlanIDVersionNewParamsReplacePricesPriceModelType = "unit_with_percent"
	PlanExternalPlanIDVersionNewParamsReplacePricesPriceModelTypePackageWithAllocation           PlanExternalPlanIDVersionNewParamsReplacePricesPriceModelType = "package_with_allocation"
	PlanExternalPlanIDVersionNewParamsReplacePricesPriceModelTypeTieredWithProration             PlanExternalPlanIDVersionNewParamsReplacePricesPriceModelType = "tiered_with_proration"
	PlanExternalPlanIDVersionNewParamsReplacePricesPriceModelTypeUnitWithProration               PlanExternalPlanIDVersionNewParamsReplacePricesPriceModelType = "unit_with_proration"
	PlanExternalPlanIDVersionNewParamsReplacePricesPriceModelTypeGroupedAllocation               PlanExternalPlanIDVersionNewParamsReplacePricesPriceModelType = "grouped_allocation"
	PlanExternalPlanIDVersionNewParamsReplacePricesPriceModelTypeGroupedWithProratedMinimum      PlanExternalPlanIDVersionNewParamsReplacePricesPriceModelType = "grouped_with_prorated_minimum"
	PlanExternalPlanIDVersionNewParamsReplacePricesPriceModelTypeGroupedWithMeteredMinimum       PlanExternalPlanIDVersionNewParamsReplacePricesPriceModelType = "grouped_with_metered_minimum"
	PlanExternalPlanIDVersionNewParamsReplacePricesPriceModelTypeMatrixWithDisplayName           PlanExternalPlanIDVersionNewParamsReplacePricesPriceModelType = "matrix_with_display_name"
	PlanExternalPlanIDVersionNewParamsReplacePricesPriceModelTypeBulkWithProration               PlanExternalPlanIDVersionNewParamsReplacePricesPriceModelType = "bulk_with_proration"
	PlanExternalPlanIDVersionNewParamsReplacePricesPriceModelTypeGroupedTieredPackage            PlanExternalPlanIDVersionNewParamsReplacePricesPriceModelType = "grouped_tiered_package"
	PlanExternalPlanIDVersionNewParamsReplacePricesPriceModelTypeMaxGroupTieredPackage           PlanExternalPlanIDVersionNewParamsReplacePricesPriceModelType = "max_group_tiered_package"
	PlanExternalPlanIDVersionNewParamsReplacePricesPriceModelTypeScalableMatrixWithUnitPricing   PlanExternalPlanIDVersionNewParamsReplacePricesPriceModelType = "scalable_matrix_with_unit_pricing"
	PlanExternalPlanIDVersionNewParamsReplacePricesPriceModelTypeScalableMatrixWithTieredPricing PlanExternalPlanIDVersionNewParamsReplacePricesPriceModelType = "scalable_matrix_with_tiered_pricing"
	PlanExternalPlanIDVersionNewParamsReplacePricesPriceModelTypeCumulativeGroupedBulk           PlanExternalPlanIDVersionNewParamsReplacePricesPriceModelType = "cumulative_grouped_bulk"
	PlanExternalPlanIDVersionNewParamsReplacePricesPriceModelTypeTieredPackageWithMinimum        PlanExternalPlanIDVersionNewParamsReplacePricesPriceModelType = "tiered_package_with_minimum"
	PlanExternalPlanIDVersionNewParamsReplacePricesPriceModelTypeMatrixWithAllocation            PlanExternalPlanIDVersionNewParamsReplacePricesPriceModelType = "matrix_with_allocation"
	PlanExternalPlanIDVersionNewParamsReplacePricesPriceModelTypeGroupedTiered                   PlanExternalPlanIDVersionNewParamsReplacePricesPriceModelType = "grouped_tiered"
)

func (r PlanExternalPlanIDVersionNewParamsReplacePricesPriceModelType) IsKnown() bool {
	switch r {
	case PlanExternalPlanIDVersionNewParamsReplacePricesPriceModelTypeUnit, PlanExternalPlanIDVersionNewParamsReplacePricesPriceModelTypePackage, PlanExternalPlanIDVersionNewParamsReplacePricesPriceModelTypeMatrix, PlanExternalPlanIDVersionNewParamsReplacePricesPriceModelTypeTiered, PlanExternalPlanIDVersionNewParamsReplacePricesPriceModelTypeTieredBps, PlanExternalPlanIDVersionNewParamsReplacePricesPriceModelTypeBps, PlanExternalPlanIDVersionNewParamsReplacePricesPriceModelTypeBulkBps, PlanExternalPlanIDVersionNewParamsReplacePricesPriceModelTypeBulk, PlanExternalPlanIDVersionNewParamsReplacePricesPriceModelTypeThresholdTotalAmount, PlanExternalPlanIDVersionNewParamsReplacePricesPriceModelTypeTieredPackage, PlanExternalPlanIDVersionNewParamsReplacePricesPriceModelTypeTieredWithMinimum, PlanExternalPlanIDVersionNewParamsReplacePricesPriceModelTypeUnitWithPercent, PlanExternalPlanIDVersionNewParamsReplacePricesPriceModelTypePackageWithAllocation, PlanExternalPlanIDVersionNewParamsReplacePricesPriceModelTypeTieredWithProration, PlanExternalPlanIDVersionNewParamsReplacePricesPriceModelTypeUnitWithProration, PlanExternalPlanIDVersionNewParamsReplacePricesPriceModelTypeGroupedAllocation, PlanExternalPlanIDVersionNewParamsReplacePricesPriceModelTypeGroupedWithProratedMinimum, PlanExternalPlanIDVersionNewParamsReplacePricesPriceModelTypeGroupedWithMeteredMinimum, PlanExternalPlanIDVersionNewParamsReplacePricesPriceModelTypeMatrixWithDisplayName, PlanExternalPlanIDVersionNewParamsReplacePricesPriceModelTypeBulkWithProration, PlanExternalPlanIDVersionNewParamsReplacePricesPriceModelTypeGroupedTieredPackage, PlanExternalPlanIDVersionNewParamsReplacePricesPriceModelTypeMaxGroupTieredPackage, PlanExternalPlanIDVersionNewParamsReplacePricesPriceModelTypeScalableMatrixWithUnitPricing, PlanExternalPlanIDVersionNewParamsReplacePricesPriceModelTypeScalableMatrixWithTieredPricing, PlanExternalPlanIDVersionNewParamsReplacePricesPriceModelTypeCumulativeGroupedBulk, PlanExternalPlanIDVersionNewParamsReplacePricesPriceModelTypeTieredPackageWithMinimum, PlanExternalPlanIDVersionNewParamsReplacePricesPriceModelTypeMatrixWithAllocation, PlanExternalPlanIDVersionNewParamsReplacePricesPriceModelTypeGroupedTiered:
		return true
	}
	return false
}
