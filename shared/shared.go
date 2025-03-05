// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package shared

import (
	"reflect"
	"time"

	"github.com/orbcorp/orb-go/internal/apijson"
	"github.com/orbcorp/orb-go/internal/param"
	"github.com/tidwall/gjson"
)

type AddCreditLedgerEntryRequestParam struct {
	EntryType param.Field[AddCreditLedgerEntryRequestEntryType] `json:"entry_type,required"`
	// The number of credits to effect. Note that this is required for increment,
	// decrement, void, or undo operations.
	Amount param.Field[float64] `json:"amount"`
	// The ID of the block affected by an expiration_change, used to differentiate
	// between multiple blocks with the same `expiry_date`.
	BlockID param.Field[string] `json:"block_id"`
	// The currency or custom pricing unit to use for this ledger entry. If this is a
	// real-world currency, it must match the customer's invoicing currency.
	Currency param.Field[string] `json:"currency"`
	// Optional metadata that can be specified when adding ledger results via the API.
	// For example, this can be used to note an increment refers to trial credits, or
	// for noting corrections as a result of an incident, etc.
	Description param.Field[string] `json:"description"`
	// An ISO 8601 format date that denotes when this credit balance should become
	// available for use.
	EffectiveDate param.Field[time.Time] `json:"effective_date" format:"date-time"`
	// An ISO 8601 format date that denotes when this credit balance should expire.
	ExpiryDate      param.Field[time.Time]   `json:"expiry_date" format:"date-time"`
	InvoiceSettings param.Field[interface{}] `json:"invoice_settings"`
	Metadata        param.Field[interface{}] `json:"metadata"`
	// Can only be specified when entry_type=increment. How much, in the customer's
	// currency, a customer paid for a single credit in this block
	PerUnitCostBasis param.Field[string] `json:"per_unit_cost_basis"`
	// A future date (specified in YYYY-MM-DD format) used for expiration change,
	// denoting when credits transferred (as part of a partial block expiration) should
	// expire.
	TargetExpiryDate param.Field[time.Time] `json:"target_expiry_date" format:"date"`
	// Can only be specified when `entry_type=void`. The reason for the void.
	VoidReason param.Field[AddCreditLedgerEntryRequestVoidReason] `json:"void_reason"`
}

func (r AddCreditLedgerEntryRequestParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r AddCreditLedgerEntryRequestParam) implementsAddCreditLedgerEntryRequestUnionParam() {}

// Satisfied by
// [shared.AddCreditLedgerEntryRequestAddIncrementCreditLedgerEntryRequestParamsParam],
// [shared.AddCreditLedgerEntryRequestAddDecrementCreditLedgerEntryRequestParamsParam],
// [shared.AddCreditLedgerEntryRequestAddExpirationChangeCreditLedgerEntryRequestParamsParam],
// [shared.AddCreditLedgerEntryRequestAddVoidCreditLedgerEntryRequestParamsParam],
// [shared.AddCreditLedgerEntryRequestAddAmendmentCreditLedgerEntryRequestParamsParam],
// [AddCreditLedgerEntryRequestParam].
type AddCreditLedgerEntryRequestUnionParam interface {
	implementsAddCreditLedgerEntryRequestUnionParam()
}

type AddCreditLedgerEntryRequestAddIncrementCreditLedgerEntryRequestParamsParam struct {
	// The number of credits to effect. Note that this is required for increment,
	// decrement, void, or undo operations.
	Amount    param.Field[float64]                                                                        `json:"amount,required"`
	EntryType param.Field[AddCreditLedgerEntryRequestAddIncrementCreditLedgerEntryRequestParamsEntryType] `json:"entry_type,required"`
	// The currency or custom pricing unit to use for this ledger entry. If this is a
	// real-world currency, it must match the customer's invoicing currency.
	Currency param.Field[string] `json:"currency"`
	// Optional metadata that can be specified when adding ledger results via the API.
	// For example, this can be used to note an increment refers to trial credits, or
	// for noting corrections as a result of an incident, etc.
	Description param.Field[string] `json:"description"`
	// An ISO 8601 format date that denotes when this credit balance should become
	// available for use.
	EffectiveDate param.Field[time.Time] `json:"effective_date" format:"date-time"`
	// An ISO 8601 format date that denotes when this credit balance should expire.
	ExpiryDate param.Field[time.Time] `json:"expiry_date" format:"date-time"`
	// Passing `invoice_settings` automatically generates an invoice for the newly
	// added credits. If `invoice_settings` is passed, you must specify
	// per_unit_cost_basis, as the calculation of the invoice total is done on that
	// basis.
	InvoiceSettings param.Field[AddCreditLedgerEntryRequestAddIncrementCreditLedgerEntryRequestParamsInvoiceSettingsParam] `json:"invoice_settings"`
	// User-specified key/value pairs for the resource. Individual keys can be removed
	// by setting the value to `null`, and the entire metadata mapping can be cleared
	// by setting `metadata` to `null`.
	Metadata param.Field[map[string]string] `json:"metadata"`
	// Can only be specified when entry_type=increment. How much, in the customer's
	// currency, a customer paid for a single credit in this block
	PerUnitCostBasis param.Field[string] `json:"per_unit_cost_basis"`
}

func (r AddCreditLedgerEntryRequestAddIncrementCreditLedgerEntryRequestParamsParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r AddCreditLedgerEntryRequestAddIncrementCreditLedgerEntryRequestParamsParam) implementsAddCreditLedgerEntryRequestUnionParam() {
}

type AddCreditLedgerEntryRequestAddIncrementCreditLedgerEntryRequestParamsEntryType string

const (
	AddCreditLedgerEntryRequestAddIncrementCreditLedgerEntryRequestParamsEntryTypeIncrement AddCreditLedgerEntryRequestAddIncrementCreditLedgerEntryRequestParamsEntryType = "increment"
)

func (r AddCreditLedgerEntryRequestAddIncrementCreditLedgerEntryRequestParamsEntryType) IsKnown() bool {
	switch r {
	case AddCreditLedgerEntryRequestAddIncrementCreditLedgerEntryRequestParamsEntryTypeIncrement:
		return true
	}
	return false
}

// Passing `invoice_settings` automatically generates an invoice for the newly
// added credits. If `invoice_settings` is passed, you must specify
// per_unit_cost_basis, as the calculation of the invoice total is done on that
// basis.
type AddCreditLedgerEntryRequestAddIncrementCreditLedgerEntryRequestParamsInvoiceSettingsParam struct {
	// Whether the credits purchase invoice should auto collect with the customer's
	// saved payment method.
	AutoCollection param.Field[bool] `json:"auto_collection,required"`
	// The net terms determines the difference between the invoice date and the issue
	// date for the invoice. If you intend the invoice to be due on issue, set this
	// to 0.
	NetTerms param.Field[int64] `json:"net_terms,required"`
	// An optional memo to display on the invoice.
	Memo param.Field[string] `json:"memo"`
	// If true, the new credit block will require that the corresponding invoice is
	// paid before it can be drawn down from.
	RequireSuccessfulPayment param.Field[bool] `json:"require_successful_payment"`
}

func (r AddCreditLedgerEntryRequestAddIncrementCreditLedgerEntryRequestParamsInvoiceSettingsParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type AddCreditLedgerEntryRequestAddDecrementCreditLedgerEntryRequestParamsParam struct {
	// The number of credits to effect. Note that this is required for increment,
	// decrement, void, or undo operations.
	Amount    param.Field[float64]                                                                        `json:"amount,required"`
	EntryType param.Field[AddCreditLedgerEntryRequestAddDecrementCreditLedgerEntryRequestParamsEntryType] `json:"entry_type,required"`
	// The currency or custom pricing unit to use for this ledger entry. If this is a
	// real-world currency, it must match the customer's invoicing currency.
	Currency param.Field[string] `json:"currency"`
	// Optional metadata that can be specified when adding ledger results via the API.
	// For example, this can be used to note an increment refers to trial credits, or
	// for noting corrections as a result of an incident, etc.
	Description param.Field[string] `json:"description"`
	// User-specified key/value pairs for the resource. Individual keys can be removed
	// by setting the value to `null`, and the entire metadata mapping can be cleared
	// by setting `metadata` to `null`.
	Metadata param.Field[map[string]string] `json:"metadata"`
}

func (r AddCreditLedgerEntryRequestAddDecrementCreditLedgerEntryRequestParamsParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r AddCreditLedgerEntryRequestAddDecrementCreditLedgerEntryRequestParamsParam) implementsAddCreditLedgerEntryRequestUnionParam() {
}

type AddCreditLedgerEntryRequestAddDecrementCreditLedgerEntryRequestParamsEntryType string

const (
	AddCreditLedgerEntryRequestAddDecrementCreditLedgerEntryRequestParamsEntryTypeDecrement AddCreditLedgerEntryRequestAddDecrementCreditLedgerEntryRequestParamsEntryType = "decrement"
)

func (r AddCreditLedgerEntryRequestAddDecrementCreditLedgerEntryRequestParamsEntryType) IsKnown() bool {
	switch r {
	case AddCreditLedgerEntryRequestAddDecrementCreditLedgerEntryRequestParamsEntryTypeDecrement:
		return true
	}
	return false
}

type AddCreditLedgerEntryRequestAddExpirationChangeCreditLedgerEntryRequestParamsParam struct {
	EntryType param.Field[AddCreditLedgerEntryRequestAddExpirationChangeCreditLedgerEntryRequestParamsEntryType] `json:"entry_type,required"`
	// An ISO 8601 format date that identifies the origination credit block to expire
	ExpiryDate param.Field[time.Time] `json:"expiry_date,required" format:"date-time"`
	// A future date (specified in YYYY-MM-DD format) used for expiration change,
	// denoting when credits transferred (as part of a partial block expiration) should
	// expire.
	TargetExpiryDate param.Field[time.Time] `json:"target_expiry_date,required" format:"date"`
	// The number of credits to effect. Note that this is required for increment,
	// decrement, void, or undo operations.
	Amount param.Field[float64] `json:"amount"`
	// The ID of the block affected by an expiration_change, used to differentiate
	// between multiple blocks with the same `expiry_date`.
	BlockID param.Field[string] `json:"block_id"`
	// The currency or custom pricing unit to use for this ledger entry. If this is a
	// real-world currency, it must match the customer's invoicing currency.
	Currency param.Field[string] `json:"currency"`
	// Optional metadata that can be specified when adding ledger results via the API.
	// For example, this can be used to note an increment refers to trial credits, or
	// for noting corrections as a result of an incident, etc.
	Description param.Field[string] `json:"description"`
	// User-specified key/value pairs for the resource. Individual keys can be removed
	// by setting the value to `null`, and the entire metadata mapping can be cleared
	// by setting `metadata` to `null`.
	Metadata param.Field[map[string]string] `json:"metadata"`
}

func (r AddCreditLedgerEntryRequestAddExpirationChangeCreditLedgerEntryRequestParamsParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r AddCreditLedgerEntryRequestAddExpirationChangeCreditLedgerEntryRequestParamsParam) implementsAddCreditLedgerEntryRequestUnionParam() {
}

type AddCreditLedgerEntryRequestAddExpirationChangeCreditLedgerEntryRequestParamsEntryType string

const (
	AddCreditLedgerEntryRequestAddExpirationChangeCreditLedgerEntryRequestParamsEntryTypeExpirationChange AddCreditLedgerEntryRequestAddExpirationChangeCreditLedgerEntryRequestParamsEntryType = "expiration_change"
)

func (r AddCreditLedgerEntryRequestAddExpirationChangeCreditLedgerEntryRequestParamsEntryType) IsKnown() bool {
	switch r {
	case AddCreditLedgerEntryRequestAddExpirationChangeCreditLedgerEntryRequestParamsEntryTypeExpirationChange:
		return true
	}
	return false
}

type AddCreditLedgerEntryRequestAddVoidCreditLedgerEntryRequestParamsParam struct {
	// The number of credits to effect. Note that this is required for increment,
	// decrement, void, or undo operations.
	Amount param.Field[float64] `json:"amount,required"`
	// The ID of the block to void.
	BlockID   param.Field[string]                                                                    `json:"block_id,required"`
	EntryType param.Field[AddCreditLedgerEntryRequestAddVoidCreditLedgerEntryRequestParamsEntryType] `json:"entry_type,required"`
	// The currency or custom pricing unit to use for this ledger entry. If this is a
	// real-world currency, it must match the customer's invoicing currency.
	Currency param.Field[string] `json:"currency"`
	// Optional metadata that can be specified when adding ledger results via the API.
	// For example, this can be used to note an increment refers to trial credits, or
	// for noting corrections as a result of an incident, etc.
	Description param.Field[string] `json:"description"`
	// User-specified key/value pairs for the resource. Individual keys can be removed
	// by setting the value to `null`, and the entire metadata mapping can be cleared
	// by setting `metadata` to `null`.
	Metadata param.Field[map[string]string] `json:"metadata"`
	// Can only be specified when `entry_type=void`. The reason for the void.
	VoidReason param.Field[AddCreditLedgerEntryRequestAddVoidCreditLedgerEntryRequestParamsVoidReason] `json:"void_reason"`
}

func (r AddCreditLedgerEntryRequestAddVoidCreditLedgerEntryRequestParamsParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r AddCreditLedgerEntryRequestAddVoidCreditLedgerEntryRequestParamsParam) implementsAddCreditLedgerEntryRequestUnionParam() {
}

type AddCreditLedgerEntryRequestAddVoidCreditLedgerEntryRequestParamsEntryType string

const (
	AddCreditLedgerEntryRequestAddVoidCreditLedgerEntryRequestParamsEntryTypeVoid AddCreditLedgerEntryRequestAddVoidCreditLedgerEntryRequestParamsEntryType = "void"
)

func (r AddCreditLedgerEntryRequestAddVoidCreditLedgerEntryRequestParamsEntryType) IsKnown() bool {
	switch r {
	case AddCreditLedgerEntryRequestAddVoidCreditLedgerEntryRequestParamsEntryTypeVoid:
		return true
	}
	return false
}

// Can only be specified when `entry_type=void`. The reason for the void.
type AddCreditLedgerEntryRequestAddVoidCreditLedgerEntryRequestParamsVoidReason string

const (
	AddCreditLedgerEntryRequestAddVoidCreditLedgerEntryRequestParamsVoidReasonRefund AddCreditLedgerEntryRequestAddVoidCreditLedgerEntryRequestParamsVoidReason = "refund"
)

func (r AddCreditLedgerEntryRequestAddVoidCreditLedgerEntryRequestParamsVoidReason) IsKnown() bool {
	switch r {
	case AddCreditLedgerEntryRequestAddVoidCreditLedgerEntryRequestParamsVoidReasonRefund:
		return true
	}
	return false
}

type AddCreditLedgerEntryRequestAddAmendmentCreditLedgerEntryRequestParamsParam struct {
	// The number of credits to effect. Note that this is required for increment,
	// decrement or void operations.
	Amount param.Field[float64] `json:"amount,required"`
	// The ID of the block to reverse a decrement from.
	BlockID   param.Field[string]                                                                         `json:"block_id,required"`
	EntryType param.Field[AddCreditLedgerEntryRequestAddAmendmentCreditLedgerEntryRequestParamsEntryType] `json:"entry_type,required"`
	// The currency or custom pricing unit to use for this ledger entry. If this is a
	// real-world currency, it must match the customer's invoicing currency.
	Currency param.Field[string] `json:"currency"`
	// Optional metadata that can be specified when adding ledger results via the API.
	// For example, this can be used to note an increment refers to trial credits, or
	// for noting corrections as a result of an incident, etc.
	Description param.Field[string] `json:"description"`
	// User-specified key/value pairs for the resource. Individual keys can be removed
	// by setting the value to `null`, and the entire metadata mapping can be cleared
	// by setting `metadata` to `null`.
	Metadata param.Field[map[string]string] `json:"metadata"`
}

func (r AddCreditLedgerEntryRequestAddAmendmentCreditLedgerEntryRequestParamsParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r AddCreditLedgerEntryRequestAddAmendmentCreditLedgerEntryRequestParamsParam) implementsAddCreditLedgerEntryRequestUnionParam() {
}

type AddCreditLedgerEntryRequestAddAmendmentCreditLedgerEntryRequestParamsEntryType string

const (
	AddCreditLedgerEntryRequestAddAmendmentCreditLedgerEntryRequestParamsEntryTypeAmendment AddCreditLedgerEntryRequestAddAmendmentCreditLedgerEntryRequestParamsEntryType = "amendment"
)

func (r AddCreditLedgerEntryRequestAddAmendmentCreditLedgerEntryRequestParamsEntryType) IsKnown() bool {
	switch r {
	case AddCreditLedgerEntryRequestAddAmendmentCreditLedgerEntryRequestParamsEntryTypeAmendment:
		return true
	}
	return false
}

type AddCreditLedgerEntryRequestEntryType string

const (
	AddCreditLedgerEntryRequestEntryTypeIncrement        AddCreditLedgerEntryRequestEntryType = "increment"
	AddCreditLedgerEntryRequestEntryTypeDecrement        AddCreditLedgerEntryRequestEntryType = "decrement"
	AddCreditLedgerEntryRequestEntryTypeExpirationChange AddCreditLedgerEntryRequestEntryType = "expiration_change"
	AddCreditLedgerEntryRequestEntryTypeVoid             AddCreditLedgerEntryRequestEntryType = "void"
	AddCreditLedgerEntryRequestEntryTypeAmendment        AddCreditLedgerEntryRequestEntryType = "amendment"
)

func (r AddCreditLedgerEntryRequestEntryType) IsKnown() bool {
	switch r {
	case AddCreditLedgerEntryRequestEntryTypeIncrement, AddCreditLedgerEntryRequestEntryTypeDecrement, AddCreditLedgerEntryRequestEntryTypeExpirationChange, AddCreditLedgerEntryRequestEntryTypeVoid, AddCreditLedgerEntryRequestEntryTypeAmendment:
		return true
	}
	return false
}

// Can only be specified when `entry_type=void`. The reason for the void.
type AddCreditLedgerEntryRequestVoidReason string

const (
	AddCreditLedgerEntryRequestVoidReasonRefund AddCreditLedgerEntryRequestVoidReason = "refund"
)

func (r AddCreditLedgerEntryRequestVoidReason) IsKnown() bool {
	switch r {
	case AddCreditLedgerEntryRequestVoidReasonRefund:
		return true
	}
	return false
}

type AddSubscriptionAdjustmentParams struct {
	// The definition of a new adjustment to create and add to the subscription.
	Adjustment param.Field[NewAdjustmentModelUnionParam] `json:"adjustment,required"`
	// The end date of the adjustment interval. This is the date that the adjustment
	// will stop affecting prices on the subscription.
	EndDate param.Field[time.Time] `json:"end_date" format:"date-time"`
	// The phase to add this adjustment to.
	PlanPhaseOrder param.Field[int64] `json:"plan_phase_order"`
	// The start date of the adjustment interval. This is the date that the adjustment
	// will start affecting prices on the subscription. If null, the adjustment will
	// start when the phase or subscription starts.
	StartDate param.Field[time.Time] `json:"start_date" format:"date-time"`
}

func (r AddSubscriptionAdjustmentParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type AddSubscriptionPriceParams struct {
	// The definition of a new allocation price to create and add to the subscription.
	AllocationPrice param.Field[NewAllocationPriceModelParam] `json:"allocation_price"`
	// [DEPRECATED] Use add_adjustments instead. The subscription's discounts for this
	// price.
	//
	// Deprecated: deprecated
	Discounts param.Field[[]DiscountOverrideModelParam] `json:"discounts"`
	// The end date of the price interval. This is the date that the price will stop
	// billing on the subscription. If null, billing will end when the phase or
	// subscription ends.
	EndDate param.Field[time.Time] `json:"end_date" format:"date-time"`
	// The external price id of the price to add to the subscription.
	ExternalPriceID param.Field[string] `json:"external_price_id"`
	// [DEPRECATED] Use add_adjustments instead. The subscription's maximum amount for
	// this price.
	//
	// Deprecated: deprecated
	MaximumAmount param.Field[string] `json:"maximum_amount"`
	// [DEPRECATED] Use add_adjustments instead. The subscription's minimum amount for
	// this price.
	//
	// Deprecated: deprecated
	MinimumAmount param.Field[string] `json:"minimum_amount"`
	// The phase to add this price to.
	PlanPhaseOrder param.Field[int64] `json:"plan_phase_order"`
	// The definition of a new price to create and add to the subscription.
	Price param.Field[NewSubscriptionPriceModelUnionParam] `json:"price"`
	// The id of the price to add to the subscription.
	PriceID param.Field[string] `json:"price_id"`
	// The start date of the price interval. This is the date that the price will start
	// billing on the subscription. If null, billing will start when the phase or
	// subscription starts.
	StartDate param.Field[time.Time] `json:"start_date" format:"date-time"`
}

func (r AddSubscriptionPriceParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type AddressInputModelParam struct {
	City       param.Field[string] `json:"city"`
	Country    param.Field[string] `json:"country"`
	Line1      param.Field[string] `json:"line1"`
	Line2      param.Field[string] `json:"line2"`
	PostalCode param.Field[string] `json:"postal_code"`
	State      param.Field[string] `json:"state"`
}

func (r AddressInputModelParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type AddressModel struct {
	City       string           `json:"city,required,nullable"`
	Country    string           `json:"country,required,nullable"`
	Line1      string           `json:"line1,required,nullable"`
	Line2      string           `json:"line2,required,nullable"`
	PostalCode string           `json:"postal_code,required,nullable"`
	State      string           `json:"state,required,nullable"`
	JSON       addressModelJSON `json:"-"`
}

// addressModelJSON contains the JSON metadata for the struct [AddressModel]
type addressModelJSON struct {
	City        apijson.Field
	Country     apijson.Field
	Line1       apijson.Field
	Line2       apijson.Field
	PostalCode  apijson.Field
	State       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AddressModel) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r addressModelJSON) RawJSON() string {
	return r.raw
}

type AdjustmentIntervalModel struct {
	ID         string          `json:"id,required"`
	Adjustment AdjustmentModel `json:"adjustment,required"`
	// The price interval IDs that this adjustment applies to.
	AppliesToPriceIntervalIDs []string `json:"applies_to_price_interval_ids,required"`
	// The end date of the adjustment interval.
	EndDate time.Time `json:"end_date,required,nullable" format:"date-time"`
	// The start date of the adjustment interval.
	StartDate time.Time                   `json:"start_date,required" format:"date-time"`
	JSON      adjustmentIntervalModelJSON `json:"-"`
}

// adjustmentIntervalModelJSON contains the JSON metadata for the struct
// [AdjustmentIntervalModel]
type adjustmentIntervalModelJSON struct {
	ID                        apijson.Field
	Adjustment                apijson.Field
	AppliesToPriceIntervalIDs apijson.Field
	EndDate                   apijson.Field
	StartDate                 apijson.Field
	raw                       string
	ExtraFields               map[string]apijson.Field
}

func (r *AdjustmentIntervalModel) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r adjustmentIntervalModelJSON) RawJSON() string {
	return r.raw
}

type AdjustmentModel struct {
	ID             string                        `json:"id,required"`
	AdjustmentType AdjustmentModelAdjustmentType `json:"adjustment_type,required"`
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
	UsageDiscount float64             `json:"usage_discount"`
	JSON          adjustmentModelJSON `json:"-"`
	union         AdjustmentModelUnion
}

// adjustmentModelJSON contains the JSON metadata for the struct [AdjustmentModel]
type adjustmentModelJSON struct {
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

func (r adjustmentModelJSON) RawJSON() string {
	return r.raw
}

func (r *AdjustmentModel) UnmarshalJSON(data []byte) (err error) {
	*r = AdjustmentModel{}
	err = apijson.UnmarshalRoot(data, &r.union)
	if err != nil {
		return err
	}
	return apijson.Port(r.union, &r)
}

// AsUnion returns a [AdjustmentModelUnion] interface which you can cast to the
// specific types for more type safety.
//
// Possible runtime types of the union are
// [shared.AdjustmentModelPlanPhaseUsageDiscountAdjustment],
// [shared.AdjustmentModelPlanPhaseAmountDiscountAdjustment],
// [shared.AdjustmentModelPlanPhasePercentageDiscountAdjustment],
// [shared.AdjustmentModelPlanPhaseMinimumAdjustment],
// [shared.AdjustmentModelPlanPhaseMaximumAdjustment].
func (r AdjustmentModel) AsUnion() AdjustmentModelUnion {
	return r.union
}

// Union satisfied by [shared.AdjustmentModelPlanPhaseUsageDiscountAdjustment],
// [shared.AdjustmentModelPlanPhaseAmountDiscountAdjustment],
// [shared.AdjustmentModelPlanPhasePercentageDiscountAdjustment],
// [shared.AdjustmentModelPlanPhaseMinimumAdjustment] or
// [shared.AdjustmentModelPlanPhaseMaximumAdjustment].
type AdjustmentModelUnion interface {
	implementsAdjustmentModel()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*AdjustmentModelUnion)(nil)).Elem(),
		"adjustment_type",
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(AdjustmentModelPlanPhaseUsageDiscountAdjustment{}),
			DiscriminatorValue: "usage_discount",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(AdjustmentModelPlanPhaseAmountDiscountAdjustment{}),
			DiscriminatorValue: "amount_discount",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(AdjustmentModelPlanPhasePercentageDiscountAdjustment{}),
			DiscriminatorValue: "percentage_discount",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(AdjustmentModelPlanPhaseMinimumAdjustment{}),
			DiscriminatorValue: "minimum",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(AdjustmentModelPlanPhaseMaximumAdjustment{}),
			DiscriminatorValue: "maximum",
		},
	)
}

type AdjustmentModelPlanPhaseUsageDiscountAdjustment struct {
	ID             string                                                        `json:"id,required"`
	AdjustmentType AdjustmentModelPlanPhaseUsageDiscountAdjustmentAdjustmentType `json:"adjustment_type,required"`
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
	UsageDiscount float64                                             `json:"usage_discount,required"`
	JSON          adjustmentModelPlanPhaseUsageDiscountAdjustmentJSON `json:"-"`
}

// adjustmentModelPlanPhaseUsageDiscountAdjustmentJSON contains the JSON metadata
// for the struct [AdjustmentModelPlanPhaseUsageDiscountAdjustment]
type adjustmentModelPlanPhaseUsageDiscountAdjustmentJSON struct {
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

func (r *AdjustmentModelPlanPhaseUsageDiscountAdjustment) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r adjustmentModelPlanPhaseUsageDiscountAdjustmentJSON) RawJSON() string {
	return r.raw
}

func (r AdjustmentModelPlanPhaseUsageDiscountAdjustment) implementsAdjustmentModel() {}

type AdjustmentModelPlanPhaseUsageDiscountAdjustmentAdjustmentType string

const (
	AdjustmentModelPlanPhaseUsageDiscountAdjustmentAdjustmentTypeUsageDiscount AdjustmentModelPlanPhaseUsageDiscountAdjustmentAdjustmentType = "usage_discount"
)

func (r AdjustmentModelPlanPhaseUsageDiscountAdjustmentAdjustmentType) IsKnown() bool {
	switch r {
	case AdjustmentModelPlanPhaseUsageDiscountAdjustmentAdjustmentTypeUsageDiscount:
		return true
	}
	return false
}

type AdjustmentModelPlanPhaseAmountDiscountAdjustment struct {
	ID             string                                                         `json:"id,required"`
	AdjustmentType AdjustmentModelPlanPhaseAmountDiscountAdjustmentAdjustmentType `json:"adjustment_type,required"`
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
	Reason string                                               `json:"reason,required,nullable"`
	JSON   adjustmentModelPlanPhaseAmountDiscountAdjustmentJSON `json:"-"`
}

// adjustmentModelPlanPhaseAmountDiscountAdjustmentJSON contains the JSON metadata
// for the struct [AdjustmentModelPlanPhaseAmountDiscountAdjustment]
type adjustmentModelPlanPhaseAmountDiscountAdjustmentJSON struct {
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

func (r *AdjustmentModelPlanPhaseAmountDiscountAdjustment) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r adjustmentModelPlanPhaseAmountDiscountAdjustmentJSON) RawJSON() string {
	return r.raw
}

func (r AdjustmentModelPlanPhaseAmountDiscountAdjustment) implementsAdjustmentModel() {}

type AdjustmentModelPlanPhaseAmountDiscountAdjustmentAdjustmentType string

const (
	AdjustmentModelPlanPhaseAmountDiscountAdjustmentAdjustmentTypeAmountDiscount AdjustmentModelPlanPhaseAmountDiscountAdjustmentAdjustmentType = "amount_discount"
)

func (r AdjustmentModelPlanPhaseAmountDiscountAdjustmentAdjustmentType) IsKnown() bool {
	switch r {
	case AdjustmentModelPlanPhaseAmountDiscountAdjustmentAdjustmentTypeAmountDiscount:
		return true
	}
	return false
}

type AdjustmentModelPlanPhasePercentageDiscountAdjustment struct {
	ID             string                                                             `json:"id,required"`
	AdjustmentType AdjustmentModelPlanPhasePercentageDiscountAdjustmentAdjustmentType `json:"adjustment_type,required"`
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
	Reason string                                                   `json:"reason,required,nullable"`
	JSON   adjustmentModelPlanPhasePercentageDiscountAdjustmentJSON `json:"-"`
}

// adjustmentModelPlanPhasePercentageDiscountAdjustmentJSON contains the JSON
// metadata for the struct [AdjustmentModelPlanPhasePercentageDiscountAdjustment]
type adjustmentModelPlanPhasePercentageDiscountAdjustmentJSON struct {
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

func (r *AdjustmentModelPlanPhasePercentageDiscountAdjustment) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r adjustmentModelPlanPhasePercentageDiscountAdjustmentJSON) RawJSON() string {
	return r.raw
}

func (r AdjustmentModelPlanPhasePercentageDiscountAdjustment) implementsAdjustmentModel() {}

type AdjustmentModelPlanPhasePercentageDiscountAdjustmentAdjustmentType string

const (
	AdjustmentModelPlanPhasePercentageDiscountAdjustmentAdjustmentTypePercentageDiscount AdjustmentModelPlanPhasePercentageDiscountAdjustmentAdjustmentType = "percentage_discount"
)

func (r AdjustmentModelPlanPhasePercentageDiscountAdjustmentAdjustmentType) IsKnown() bool {
	switch r {
	case AdjustmentModelPlanPhasePercentageDiscountAdjustmentAdjustmentTypePercentageDiscount:
		return true
	}
	return false
}

type AdjustmentModelPlanPhaseMinimumAdjustment struct {
	ID             string                                                  `json:"id,required"`
	AdjustmentType AdjustmentModelPlanPhaseMinimumAdjustmentAdjustmentType `json:"adjustment_type,required"`
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
	Reason string                                        `json:"reason,required,nullable"`
	JSON   adjustmentModelPlanPhaseMinimumAdjustmentJSON `json:"-"`
}

// adjustmentModelPlanPhaseMinimumAdjustmentJSON contains the JSON metadata for the
// struct [AdjustmentModelPlanPhaseMinimumAdjustment]
type adjustmentModelPlanPhaseMinimumAdjustmentJSON struct {
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

func (r *AdjustmentModelPlanPhaseMinimumAdjustment) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r adjustmentModelPlanPhaseMinimumAdjustmentJSON) RawJSON() string {
	return r.raw
}

func (r AdjustmentModelPlanPhaseMinimumAdjustment) implementsAdjustmentModel() {}

type AdjustmentModelPlanPhaseMinimumAdjustmentAdjustmentType string

const (
	AdjustmentModelPlanPhaseMinimumAdjustmentAdjustmentTypeMinimum AdjustmentModelPlanPhaseMinimumAdjustmentAdjustmentType = "minimum"
)

func (r AdjustmentModelPlanPhaseMinimumAdjustmentAdjustmentType) IsKnown() bool {
	switch r {
	case AdjustmentModelPlanPhaseMinimumAdjustmentAdjustmentTypeMinimum:
		return true
	}
	return false
}

type AdjustmentModelPlanPhaseMaximumAdjustment struct {
	ID             string                                                  `json:"id,required"`
	AdjustmentType AdjustmentModelPlanPhaseMaximumAdjustmentAdjustmentType `json:"adjustment_type,required"`
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
	Reason string                                        `json:"reason,required,nullable"`
	JSON   adjustmentModelPlanPhaseMaximumAdjustmentJSON `json:"-"`
}

// adjustmentModelPlanPhaseMaximumAdjustmentJSON contains the JSON metadata for the
// struct [AdjustmentModelPlanPhaseMaximumAdjustment]
type adjustmentModelPlanPhaseMaximumAdjustmentJSON struct {
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

func (r *AdjustmentModelPlanPhaseMaximumAdjustment) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r adjustmentModelPlanPhaseMaximumAdjustmentJSON) RawJSON() string {
	return r.raw
}

func (r AdjustmentModelPlanPhaseMaximumAdjustment) implementsAdjustmentModel() {}

type AdjustmentModelPlanPhaseMaximumAdjustmentAdjustmentType string

const (
	AdjustmentModelPlanPhaseMaximumAdjustmentAdjustmentTypeMaximum AdjustmentModelPlanPhaseMaximumAdjustmentAdjustmentType = "maximum"
)

func (r AdjustmentModelPlanPhaseMaximumAdjustmentAdjustmentType) IsKnown() bool {
	switch r {
	case AdjustmentModelPlanPhaseMaximumAdjustmentAdjustmentTypeMaximum:
		return true
	}
	return false
}

type AdjustmentModelAdjustmentType string

const (
	AdjustmentModelAdjustmentTypeUsageDiscount      AdjustmentModelAdjustmentType = "usage_discount"
	AdjustmentModelAdjustmentTypeAmountDiscount     AdjustmentModelAdjustmentType = "amount_discount"
	AdjustmentModelAdjustmentTypePercentageDiscount AdjustmentModelAdjustmentType = "percentage_discount"
	AdjustmentModelAdjustmentTypeMinimum            AdjustmentModelAdjustmentType = "minimum"
	AdjustmentModelAdjustmentTypeMaximum            AdjustmentModelAdjustmentType = "maximum"
)

func (r AdjustmentModelAdjustmentType) IsKnown() bool {
	switch r {
	case AdjustmentModelAdjustmentTypeUsageDiscount, AdjustmentModelAdjustmentTypeAmountDiscount, AdjustmentModelAdjustmentTypePercentageDiscount, AdjustmentModelAdjustmentTypeMinimum, AdjustmentModelAdjustmentTypeMaximum:
		return true
	}
	return false
}

type AffectedBlockModel struct {
	ID               string                 `json:"id,required"`
	ExpiryDate       time.Time              `json:"expiry_date,required,nullable" format:"date-time"`
	PerUnitCostBasis string                 `json:"per_unit_cost_basis,required,nullable"`
	JSON             affectedBlockModelJSON `json:"-"`
}

// affectedBlockModelJSON contains the JSON metadata for the struct
// [AffectedBlockModel]
type affectedBlockModelJSON struct {
	ID               apijson.Field
	ExpiryDate       apijson.Field
	PerUnitCostBasis apijson.Field
	raw              string
	ExtraFields      map[string]apijson.Field
}

func (r *AffectedBlockModel) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r affectedBlockModelJSON) RawJSON() string {
	return r.raw
}

type AggregatedCostModel struct {
	PerPriceCosts []AggregatedCostModelPerPriceCost `json:"per_price_costs,required"`
	// Total costs for the timeframe, excluding any minimums and discounts.
	Subtotal       string    `json:"subtotal,required"`
	TimeframeEnd   time.Time `json:"timeframe_end,required" format:"date-time"`
	TimeframeStart time.Time `json:"timeframe_start,required" format:"date-time"`
	// Total costs for the timeframe, including any minimums and discounts.
	Total string                  `json:"total,required"`
	JSON  aggregatedCostModelJSON `json:"-"`
}

// aggregatedCostModelJSON contains the JSON metadata for the struct
// [AggregatedCostModel]
type aggregatedCostModelJSON struct {
	PerPriceCosts  apijson.Field
	Subtotal       apijson.Field
	TimeframeEnd   apijson.Field
	TimeframeStart apijson.Field
	Total          apijson.Field
	raw            string
	ExtraFields    map[string]apijson.Field
}

func (r *AggregatedCostModel) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r aggregatedCostModelJSON) RawJSON() string {
	return r.raw
}

type AggregatedCostModelPerPriceCost struct {
	// The price object
	Price PriceModel `json:"price,required"`
	// The price the cost is associated with
	PriceID string `json:"price_id,required"`
	// Price's contributions for the timeframe, excluding any minimums and discounts.
	Subtotal string `json:"subtotal,required"`
	// Price's contributions for the timeframe, including minimums and discounts.
	Total string `json:"total,required"`
	// The price's quantity for the timeframe
	Quantity float64                             `json:"quantity,nullable"`
	JSON     aggregatedCostModelPerPriceCostJSON `json:"-"`
}

// aggregatedCostModelPerPriceCostJSON contains the JSON metadata for the struct
// [AggregatedCostModelPerPriceCost]
type aggregatedCostModelPerPriceCostJSON struct {
	Price       apijson.Field
	PriceID     apijson.Field
	Subtotal    apijson.Field
	Total       apijson.Field
	Quantity    apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AggregatedCostModelPerPriceCost) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r aggregatedCostModelPerPriceCostJSON) RawJSON() string {
	return r.raw
}

// [Alerts within Orb](/product-catalog/configuring-alerts) monitor spending,
// usage, or credit balance and trigger webhooks when a threshold is exceeded.
//
// Alerts created through the API can be scoped to either customers or
// subscriptions.
type AlertModel struct {
	// Also referred to as alert_id in this documentation.
	ID string `json:"id,required"`
	// The creation time of the resource in Orb.
	CreatedAt time.Time `json:"created_at,required" format:"date-time"`
	// The name of the currency the credit balance or invoice cost is denominated in.
	Currency string `json:"currency,required,nullable"`
	// The customer the alert applies to.
	Customer CustomerMinifiedModel `json:"customer,required,nullable"`
	// Whether the alert is enabled or disabled.
	Enabled bool `json:"enabled,required"`
	// The metric the alert applies to.
	Metric AlertModelMetric `json:"metric,required,nullable"`
	// The plan the alert applies to.
	Plan AlertModelPlan `json:"plan,required,nullable"`
	// The subscription the alert applies to.
	Subscription SubscriptionMinifiedModel `json:"subscription,required,nullable"`
	// The thresholds that define the conditions under which the alert will be
	// triggered.
	Thresholds []ThresholdModel `json:"thresholds,required,nullable"`
	// The type of alert. This must be a valid alert type.
	Type AlertModelType `json:"type,required"`
	JSON alertModelJSON `json:"-"`
}

// alertModelJSON contains the JSON metadata for the struct [AlertModel]
type alertModelJSON struct {
	ID           apijson.Field
	CreatedAt    apijson.Field
	Currency     apijson.Field
	Customer     apijson.Field
	Enabled      apijson.Field
	Metric       apijson.Field
	Plan         apijson.Field
	Subscription apijson.Field
	Thresholds   apijson.Field
	Type         apijson.Field
	raw          string
	ExtraFields  map[string]apijson.Field
}

func (r *AlertModel) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r alertModelJSON) RawJSON() string {
	return r.raw
}

// The metric the alert applies to.
type AlertModelMetric struct {
	ID   string               `json:"id,required"`
	JSON alertModelMetricJSON `json:"-"`
}

// alertModelMetricJSON contains the JSON metadata for the struct
// [AlertModelMetric]
type alertModelMetricJSON struct {
	ID          apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AlertModelMetric) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r alertModelMetricJSON) RawJSON() string {
	return r.raw
}

// The plan the alert applies to.
type AlertModelPlan struct {
	ID string `json:"id,required,nullable"`
	// An optional user-defined ID for this plan resource, used throughout the system
	// as an alias for this Plan. Use this field to identify a plan by an existing
	// identifier in your system.
	ExternalPlanID string             `json:"external_plan_id,required,nullable"`
	Name           string             `json:"name,required,nullable"`
	PlanVersion    string             `json:"plan_version,required"`
	JSON           alertModelPlanJSON `json:"-"`
}

// alertModelPlanJSON contains the JSON metadata for the struct [AlertModelPlan]
type alertModelPlanJSON struct {
	ID             apijson.Field
	ExternalPlanID apijson.Field
	Name           apijson.Field
	PlanVersion    apijson.Field
	raw            string
	ExtraFields    map[string]apijson.Field
}

func (r *AlertModelPlan) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r alertModelPlanJSON) RawJSON() string {
	return r.raw
}

// The type of alert. This must be a valid alert type.
type AlertModelType string

const (
	AlertModelTypeUsageExceeded          AlertModelType = "usage_exceeded"
	AlertModelTypeCostExceeded           AlertModelType = "cost_exceeded"
	AlertModelTypeCreditBalanceDepleted  AlertModelType = "credit_balance_depleted"
	AlertModelTypeCreditBalanceDropped   AlertModelType = "credit_balance_dropped"
	AlertModelTypeCreditBalanceRecovered AlertModelType = "credit_balance_recovered"
)

func (r AlertModelType) IsKnown() bool {
	switch r {
	case AlertModelTypeUsageExceeded, AlertModelTypeCostExceeded, AlertModelTypeCreditBalanceDepleted, AlertModelTypeCreditBalanceDropped, AlertModelTypeCreditBalanceRecovered:
		return true
	}
	return false
}

type AllocationModel struct {
	AllowsRollover bool                `json:"allows_rollover,required"`
	Currency       string              `json:"currency,required"`
	JSON           allocationModelJSON `json:"-"`
}

// allocationModelJSON contains the JSON metadata for the struct [AllocationModel]
type allocationModelJSON struct {
	AllowsRollover apijson.Field
	Currency       apijson.Field
	raw            string
	ExtraFields    map[string]apijson.Field
}

func (r *AllocationModel) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r allocationModelJSON) RawJSON() string {
	return r.raw
}

type AmountDiscount struct {
	// Only available if discount_type is `amount`.
	AmountDiscount string `json:"amount_discount,required"`
	// List of price_ids that this discount applies to. For plan/plan phase discounts,
	// this can be a subset of prices.
	AppliesToPriceIDs []string                   `json:"applies_to_price_ids,required"`
	DiscountType      AmountDiscountDiscountType `json:"discount_type,required"`
	Reason            string                     `json:"reason,nullable"`
	JSON              amountDiscountJSON         `json:"-"`
}

// amountDiscountJSON contains the JSON metadata for the struct [AmountDiscount]
type amountDiscountJSON struct {
	AmountDiscount    apijson.Field
	AppliesToPriceIDs apijson.Field
	DiscountType      apijson.Field
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

func (r AmountDiscount) ImplementsCouponModelDiscount() {}

func (r AmountDiscount) ImplementsDiscount() {}

func (r AmountDiscount) ImplementsInvoiceLevelDiscount() {}

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
	AmountDiscount param.Field[string] `json:"amount_discount,required"`
	// List of price_ids that this discount applies to. For plan/plan phase discounts,
	// this can be a subset of prices.
	AppliesToPriceIDs param.Field[[]string]                   `json:"applies_to_price_ids,required"`
	DiscountType      param.Field[AmountDiscountDiscountType] `json:"discount_type,required"`
	Reason            param.Field[string]                     `json:"reason"`
}

func (r AmountDiscountParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r AmountDiscountParam) ImplementsDiscountUnionParam() {}

type AmountDiscountIntervalModel struct {
	// Only available if discount_type is `amount`.
	AmountDiscount string `json:"amount_discount,required"`
	// The price ids that this discount interval applies to.
	AppliesToPriceIDs []string `json:"applies_to_price_ids,required"`
	// The price interval ids that this discount interval applies to.
	AppliesToPriceIntervalIDs []string                                `json:"applies_to_price_interval_ids,required"`
	DiscountType              AmountDiscountIntervalModelDiscountType `json:"discount_type,required"`
	// The end date of the discount interval.
	EndDate time.Time `json:"end_date,required,nullable" format:"date-time"`
	// The start date of the discount interval.
	StartDate time.Time                       `json:"start_date,required" format:"date-time"`
	JSON      amountDiscountIntervalModelJSON `json:"-"`
}

// amountDiscountIntervalModelJSON contains the JSON metadata for the struct
// [AmountDiscountIntervalModel]
type amountDiscountIntervalModelJSON struct {
	AmountDiscount            apijson.Field
	AppliesToPriceIDs         apijson.Field
	AppliesToPriceIntervalIDs apijson.Field
	DiscountType              apijson.Field
	EndDate                   apijson.Field
	StartDate                 apijson.Field
	raw                       string
	ExtraFields               map[string]apijson.Field
}

func (r *AmountDiscountIntervalModel) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r amountDiscountIntervalModelJSON) RawJSON() string {
	return r.raw
}

func (r AmountDiscountIntervalModel) ImplementsMutatedSubscriptionModelDiscountInterval() {}

func (r AmountDiscountIntervalModel) ImplementsSubscriptionModelDiscountInterval() {}

type AmountDiscountIntervalModelDiscountType string

const (
	AmountDiscountIntervalModelDiscountTypeAmount AmountDiscountIntervalModelDiscountType = "amount"
)

func (r AmountDiscountIntervalModelDiscountType) IsKnown() bool {
	switch r {
	case AmountDiscountIntervalModelDiscountTypeAmount:
		return true
	}
	return false
}

type AutoCollectionModel struct {
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
	PreviouslyAttemptedAt time.Time               `json:"previously_attempted_at,required,nullable" format:"date-time"`
	JSON                  autoCollectionModelJSON `json:"-"`
}

// autoCollectionModelJSON contains the JSON metadata for the struct
// [AutoCollectionModel]
type autoCollectionModelJSON struct {
	Enabled               apijson.Field
	NextAttemptAt         apijson.Field
	NumAttempts           apijson.Field
	PreviouslyAttemptedAt apijson.Field
	raw                   string
	ExtraFields           map[string]apijson.Field
}

func (r *AutoCollectionModel) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r autoCollectionModelJSON) RawJSON() string {
	return r.raw
}

// A backfill represents an update to historical usage data, adding or replacing
// events in a timeframe.
type BackfillModel struct {
	ID string `json:"id,required"`
	// If in the future, the time at which the backfill will automatically close. If in
	// the past, the time at which the backfill was closed.
	CloseTime time.Time `json:"close_time,required,nullable" format:"date-time"`
	CreatedAt time.Time `json:"created_at,required" format:"date-time"`
	// The Orb-generated ID of the customer to which this backfill is scoped. If
	// `null`, this backfill is scoped to all customers.
	CustomerID string `json:"customer_id,required,nullable"`
	// The number of events ingested in this backfill.
	EventsIngested int64 `json:"events_ingested,required"`
	// If `true`, existing events in the backfill's timeframe will be replaced with the
	// newly ingested events associated with the backfill. If `false`, newly ingested
	// events will be added to the existing events.
	ReplaceExistingEvents bool `json:"replace_existing_events,required"`
	// The time at which this backfill was reverted.
	RevertedAt time.Time `json:"reverted_at,required,nullable" format:"date-time"`
	// The status of the backfill.
	Status         BackfillModelStatus `json:"status,required"`
	TimeframeEnd   time.Time           `json:"timeframe_end,required" format:"date-time"`
	TimeframeStart time.Time           `json:"timeframe_start,required" format:"date-time"`
	// A boolean
	// [computed property](/extensibility/advanced-metrics#computed-properties) used to
	// filter the set of events to deprecate
	DeprecationFilter string            `json:"deprecation_filter,nullable"`
	JSON              backfillModelJSON `json:"-"`
}

// backfillModelJSON contains the JSON metadata for the struct [BackfillModel]
type backfillModelJSON struct {
	ID                    apijson.Field
	CloseTime             apijson.Field
	CreatedAt             apijson.Field
	CustomerID            apijson.Field
	EventsIngested        apijson.Field
	ReplaceExistingEvents apijson.Field
	RevertedAt            apijson.Field
	Status                apijson.Field
	TimeframeEnd          apijson.Field
	TimeframeStart        apijson.Field
	DeprecationFilter     apijson.Field
	raw                   string
	ExtraFields           map[string]apijson.Field
}

func (r *BackfillModel) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r backfillModelJSON) RawJSON() string {
	return r.raw
}

// The status of the backfill.
type BackfillModelStatus string

const (
	BackfillModelStatusPending       BackfillModelStatus = "pending"
	BackfillModelStatusReflected     BackfillModelStatus = "reflected"
	BackfillModelStatusPendingRevert BackfillModelStatus = "pending_revert"
	BackfillModelStatusReverted      BackfillModelStatus = "reverted"
)

func (r BackfillModelStatus) IsKnown() bool {
	switch r {
	case BackfillModelStatusPending, BackfillModelStatusReflected, BackfillModelStatusPendingRevert, BackfillModelStatusReverted:
		return true
	}
	return false
}

// The Metric resource represents a calculation of a quantity based on events.
// Metrics are defined by the query that transforms raw usage events into
// meaningful values for your customers.
type BillableMetricModel struct {
	ID          string `json:"id,required"`
	Description string `json:"description,required,nullable"`
	// The Item resource represents a sellable product or good. Items are associated
	// with all line items, billable metrics, and prices and are used for defining
	// external sync behavior for invoices and tax calculation purposes.
	Item ItemModel `json:"item,required"`
	// User specified key-value pairs for the resource. If not present, this defaults
	// to an empty dictionary. Individual keys can be removed by setting the value to
	// `null`, and the entire metadata mapping can be cleared by setting `metadata` to
	// `null`.
	Metadata map[string]string         `json:"metadata,required"`
	Name     string                    `json:"name,required"`
	Status   BillableMetricModelStatus `json:"status,required"`
	JSON     billableMetricModelJSON   `json:"-"`
}

// billableMetricModelJSON contains the JSON metadata for the struct
// [BillableMetricModel]
type billableMetricModelJSON struct {
	ID          apijson.Field
	Description apijson.Field
	Item        apijson.Field
	Metadata    apijson.Field
	Name        apijson.Field
	Status      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *BillableMetricModel) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r billableMetricModelJSON) RawJSON() string {
	return r.raw
}

type BillableMetricModelStatus string

const (
	BillableMetricModelStatusActive   BillableMetricModelStatus = "active"
	BillableMetricModelStatusDraft    BillableMetricModelStatus = "draft"
	BillableMetricModelStatusArchived BillableMetricModelStatus = "archived"
)

func (r BillableMetricModelStatus) IsKnown() bool {
	switch r {
	case BillableMetricModelStatusActive, BillableMetricModelStatusDraft, BillableMetricModelStatusArchived:
		return true
	}
	return false
}

type BillableMetricSimpleModel struct {
	ID   string                        `json:"id,required"`
	Name string                        `json:"name,required"`
	JSON billableMetricSimpleModelJSON `json:"-"`
}

// billableMetricSimpleModelJSON contains the JSON metadata for the struct
// [BillableMetricSimpleModel]
type billableMetricSimpleModelJSON struct {
	ID          apijson.Field
	Name        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *BillableMetricSimpleModel) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r billableMetricSimpleModelJSON) RawJSON() string {
	return r.raw
}

type BillableMetricTinyModel struct {
	ID   string                      `json:"id,required"`
	JSON billableMetricTinyModelJSON `json:"-"`
}

// billableMetricTinyModelJSON contains the JSON metadata for the struct
// [BillableMetricTinyModel]
type billableMetricTinyModelJSON struct {
	ID          apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *BillableMetricTinyModel) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r billableMetricTinyModelJSON) RawJSON() string {
	return r.raw
}

type BillingCycleAnchorConfigurationModel struct {
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
	Year int64                                    `json:"year,nullable"`
	JSON billingCycleAnchorConfigurationModelJSON `json:"-"`
}

// billingCycleAnchorConfigurationModelJSON contains the JSON metadata for the
// struct [BillingCycleAnchorConfigurationModel]
type billingCycleAnchorConfigurationModelJSON struct {
	Day         apijson.Field
	Month       apijson.Field
	Year        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *BillingCycleAnchorConfigurationModel) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r billingCycleAnchorConfigurationModelJSON) RawJSON() string {
	return r.raw
}

type BillingCycleAnchorConfigurationModelParam struct {
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

func (r BillingCycleAnchorConfigurationModelParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type BillingCycleConfigurationModel struct {
	Duration     int64                                      `json:"duration,required"`
	DurationUnit BillingCycleConfigurationModelDurationUnit `json:"duration_unit,required"`
	JSON         billingCycleConfigurationModelJSON         `json:"-"`
}

// billingCycleConfigurationModelJSON contains the JSON metadata for the struct
// [BillingCycleConfigurationModel]
type billingCycleConfigurationModelJSON struct {
	Duration     apijson.Field
	DurationUnit apijson.Field
	raw          string
	ExtraFields  map[string]apijson.Field
}

func (r *BillingCycleConfigurationModel) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r billingCycleConfigurationModelJSON) RawJSON() string {
	return r.raw
}

type BillingCycleConfigurationModelDurationUnit string

const (
	BillingCycleConfigurationModelDurationUnitDay   BillingCycleConfigurationModelDurationUnit = "day"
	BillingCycleConfigurationModelDurationUnitMonth BillingCycleConfigurationModelDurationUnit = "month"
)

func (r BillingCycleConfigurationModelDurationUnit) IsKnown() bool {
	switch r {
	case BillingCycleConfigurationModelDurationUnitDay, BillingCycleConfigurationModelDurationUnitMonth:
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

type BpsConfigModel struct {
	// Basis point take rate per event
	Bps float64 `json:"bps,required"`
	// Optional currency amount maximum to cap spend per event
	PerUnitMaximum string             `json:"per_unit_maximum,nullable"`
	JSON           bpsConfigModelJSON `json:"-"`
}

// bpsConfigModelJSON contains the JSON metadata for the struct [BpsConfigModel]
type bpsConfigModelJSON struct {
	Bps            apijson.Field
	PerUnitMaximum apijson.Field
	raw            string
	ExtraFields    map[string]apijson.Field
}

func (r *BpsConfigModel) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r bpsConfigModelJSON) RawJSON() string {
	return r.raw
}

type BpsConfigModelParam struct {
	// Basis point take rate per event
	Bps param.Field[float64] `json:"bps,required"`
	// Optional currency amount maximum to cap spend per event
	PerUnitMaximum param.Field[string] `json:"per_unit_maximum"`
}

func (r BpsConfigModelParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type BulkBpsConfigModel struct {
	// Tiers for a bulk BPS pricing model where all usage is aggregated to a single
	// tier based on total volume
	Tiers []BulkBpsConfigModelTier `json:"tiers,required"`
	JSON  bulkBpsConfigModelJSON   `json:"-"`
}

// bulkBpsConfigModelJSON contains the JSON metadata for the struct
// [BulkBpsConfigModel]
type bulkBpsConfigModelJSON struct {
	Tiers       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *BulkBpsConfigModel) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r bulkBpsConfigModelJSON) RawJSON() string {
	return r.raw
}

type BulkBpsConfigModelTier struct {
	// Basis points to rate on
	Bps float64 `json:"bps,required"`
	// Upper bound for tier
	MaximumAmount string `json:"maximum_amount,nullable"`
	// The maximum amount to charge for any one event
	PerUnitMaximum string                     `json:"per_unit_maximum,nullable"`
	JSON           bulkBpsConfigModelTierJSON `json:"-"`
}

// bulkBpsConfigModelTierJSON contains the JSON metadata for the struct
// [BulkBpsConfigModelTier]
type bulkBpsConfigModelTierJSON struct {
	Bps            apijson.Field
	MaximumAmount  apijson.Field
	PerUnitMaximum apijson.Field
	raw            string
	ExtraFields    map[string]apijson.Field
}

func (r *BulkBpsConfigModelTier) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r bulkBpsConfigModelTierJSON) RawJSON() string {
	return r.raw
}

type BulkBpsConfigModelParam struct {
	// Tiers for a bulk BPS pricing model where all usage is aggregated to a single
	// tier based on total volume
	Tiers param.Field[[]BulkBpsConfigModelTierParam] `json:"tiers,required"`
}

func (r BulkBpsConfigModelParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type BulkBpsConfigModelTierParam struct {
	// Basis points to rate on
	Bps param.Field[float64] `json:"bps,required"`
	// Upper bound for tier
	MaximumAmount param.Field[string] `json:"maximum_amount"`
	// The maximum amount to charge for any one event
	PerUnitMaximum param.Field[string] `json:"per_unit_maximum"`
}

func (r BulkBpsConfigModelTierParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type BulkConfigModel struct {
	// Bulk tiers for rating based on total usage volume
	Tiers []BulkConfigModelTier `json:"tiers,required"`
	JSON  bulkConfigModelJSON   `json:"-"`
}

// bulkConfigModelJSON contains the JSON metadata for the struct [BulkConfigModel]
type bulkConfigModelJSON struct {
	Tiers       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *BulkConfigModel) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r bulkConfigModelJSON) RawJSON() string {
	return r.raw
}

type BulkConfigModelTier struct {
	// Amount per unit
	UnitAmount string `json:"unit_amount,required"`
	// Upper bound for this tier
	MaximumUnits float64                 `json:"maximum_units,nullable"`
	JSON         bulkConfigModelTierJSON `json:"-"`
}

// bulkConfigModelTierJSON contains the JSON metadata for the struct
// [BulkConfigModelTier]
type bulkConfigModelTierJSON struct {
	UnitAmount   apijson.Field
	MaximumUnits apijson.Field
	raw          string
	ExtraFields  map[string]apijson.Field
}

func (r *BulkConfigModelTier) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r bulkConfigModelTierJSON) RawJSON() string {
	return r.raw
}

type BulkConfigModelParam struct {
	// Bulk tiers for rating based on total usage volume
	Tiers param.Field[[]BulkConfigModelTierParam] `json:"tiers,required"`
}

func (r BulkConfigModelParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type BulkConfigModelTierParam struct {
	// Amount per unit
	UnitAmount param.Field[string] `json:"unit_amount,required"`
	// Upper bound for this tier
	MaximumUnits param.Field[float64] `json:"maximum_units"`
}

func (r BulkConfigModelTierParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// A coupon represents a reusable discount configuration that can be applied either
// as a fixed or percentage amount to an invoice or subscription. Coupons are
// activated using a redemption code, which applies the discount to a subscription
// or invoice. The duration of a coupon determines how long it remains available
// for use by end users.
type CouponModel struct {
	// Also referred to as coupon_id in this documentation.
	ID string `json:"id,required"`
	// An archived coupon can no longer be redeemed. Active coupons will have a value
	// of null for `archived_at`; this field will be non-null for archived coupons.
	ArchivedAt time.Time           `json:"archived_at,required,nullable" format:"date-time"`
	Discount   CouponModelDiscount `json:"discount,required"`
	// This allows for a coupon's discount to apply for a limited time (determined in
	// months); a `null` value here means "unlimited time".
	DurationInMonths int64 `json:"duration_in_months,required,nullable"`
	// The maximum number of redemptions allowed for this coupon before it is
	// exhausted; `null` here means "unlimited".
	MaxRedemptions int64 `json:"max_redemptions,required,nullable"`
	// This string can be used to redeem this coupon for a given subscription.
	RedemptionCode string `json:"redemption_code,required"`
	// The number of times this coupon has been redeemed.
	TimesRedeemed int64           `json:"times_redeemed,required"`
	JSON          couponModelJSON `json:"-"`
}

// couponModelJSON contains the JSON metadata for the struct [CouponModel]
type couponModelJSON struct {
	ID               apijson.Field
	ArchivedAt       apijson.Field
	Discount         apijson.Field
	DurationInMonths apijson.Field
	MaxRedemptions   apijson.Field
	RedemptionCode   apijson.Field
	TimesRedeemed    apijson.Field
	raw              string
	ExtraFields      map[string]apijson.Field
}

func (r *CouponModel) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r couponModelJSON) RawJSON() string {
	return r.raw
}

type CouponModelDiscount struct {
	// This field can have the runtime type of [[]string].
	AppliesToPriceIDs interface{}                     `json:"applies_to_price_ids,required"`
	DiscountType      CouponModelDiscountDiscountType `json:"discount_type,required"`
	// Only available if discount_type is `amount`.
	AmountDiscount string `json:"amount_discount"`
	// Only available if discount_type is `percentage`. This is a number between 0
	// and 1.
	PercentageDiscount float64                 `json:"percentage_discount"`
	Reason             string                  `json:"reason,nullable"`
	JSON               couponModelDiscountJSON `json:"-"`
	union              CouponModelDiscountUnion
}

// couponModelDiscountJSON contains the JSON metadata for the struct
// [CouponModelDiscount]
type couponModelDiscountJSON struct {
	AppliesToPriceIDs  apijson.Field
	DiscountType       apijson.Field
	AmountDiscount     apijson.Field
	PercentageDiscount apijson.Field
	Reason             apijson.Field
	raw                string
	ExtraFields        map[string]apijson.Field
}

func (r couponModelDiscountJSON) RawJSON() string {
	return r.raw
}

func (r *CouponModelDiscount) UnmarshalJSON(data []byte) (err error) {
	*r = CouponModelDiscount{}
	err = apijson.UnmarshalRoot(data, &r.union)
	if err != nil {
		return err
	}
	return apijson.Port(r.union, &r)
}

// AsUnion returns a [CouponModelDiscountUnion] interface which you can cast to the
// specific types for more type safety.
//
// Possible runtime types of the union are [shared.PercentageDiscount],
// [shared.AmountDiscount].
func (r CouponModelDiscount) AsUnion() CouponModelDiscountUnion {
	return r.union
}

// Union satisfied by [shared.PercentageDiscount] or [shared.AmountDiscount].
type CouponModelDiscountUnion interface {
	ImplementsCouponModelDiscount()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*CouponModelDiscountUnion)(nil)).Elem(),
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
	)
}

type CouponModelDiscountDiscountType string

const (
	CouponModelDiscountDiscountTypePercentage CouponModelDiscountDiscountType = "percentage"
	CouponModelDiscountDiscountTypeAmount     CouponModelDiscountDiscountType = "amount"
)

func (r CouponModelDiscountDiscountType) IsKnown() bool {
	switch r {
	case CouponModelDiscountDiscountTypePercentage, CouponModelDiscountDiscountTypeAmount:
		return true
	}
	return false
}

type CouponRedemptionModel struct {
	CouponID  string                    `json:"coupon_id,required"`
	EndDate   time.Time                 `json:"end_date,required,nullable" format:"date-time"`
	StartDate time.Time                 `json:"start_date,required" format:"date-time"`
	JSON      couponRedemptionModelJSON `json:"-"`
}

// couponRedemptionModelJSON contains the JSON metadata for the struct
// [CouponRedemptionModel]
type couponRedemptionModelJSON struct {
	CouponID    apijson.Field
	EndDate     apijson.Field
	StartDate   apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *CouponRedemptionModel) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r couponRedemptionModelJSON) RawJSON() string {
	return r.raw
}

type CreditLedgerEntriesModel struct {
	Data               []CreditLedgerEntryModel     `json:"data,required"`
	PaginationMetadata PaginationMetadata           `json:"pagination_metadata,required"`
	JSON               creditLedgerEntriesModelJSON `json:"-"`
}

// creditLedgerEntriesModelJSON contains the JSON metadata for the struct
// [CreditLedgerEntriesModel]
type creditLedgerEntriesModelJSON struct {
	Data               apijson.Field
	PaginationMetadata apijson.Field
	raw                string
	ExtraFields        map[string]apijson.Field
}

func (r *CreditLedgerEntriesModel) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r creditLedgerEntriesModelJSON) RawJSON() string {
	return r.raw
}

// The [Credit Ledger Entry resource](/product-catalog/prepurchase) models prepaid
// credits within Orb.
type CreditLedgerEntryModel struct {
	ID                   string                            `json:"id,required"`
	Amount               float64                           `json:"amount,required"`
	CreatedAt            time.Time                         `json:"created_at,required" format:"date-time"`
	CreditBlock          AffectedBlockModel                `json:"credit_block,required"`
	Currency             string                            `json:"currency,required"`
	Customer             CustomerMinifiedModel             `json:"customer,required"`
	Description          string                            `json:"description,required,nullable"`
	EndingBalance        float64                           `json:"ending_balance,required"`
	EntryStatus          CreditLedgerEntryModelEntryStatus `json:"entry_status,required"`
	EntryType            CreditLedgerEntryModelEntryType   `json:"entry_type,required"`
	LedgerSequenceNumber int64                             `json:"ledger_sequence_number,required"`
	// This field can have the runtime type of [map[string]string].
	Metadata           interface{}                `json:"metadata,required"`
	StartingBalance    float64                    `json:"starting_balance,required"`
	EventID            string                     `json:"event_id,nullable"`
	InvoiceID          string                     `json:"invoice_id,nullable"`
	NewBlockExpiryDate time.Time                  `json:"new_block_expiry_date,nullable" format:"date-time"`
	PriceID            string                     `json:"price_id,nullable"`
	VoidAmount         float64                    `json:"void_amount"`
	VoidReason         string                     `json:"void_reason,nullable"`
	JSON               creditLedgerEntryModelJSON `json:"-"`
	union              CreditLedgerEntryModelUnion
}

// creditLedgerEntryModelJSON contains the JSON metadata for the struct
// [CreditLedgerEntryModel]
type creditLedgerEntryModelJSON struct {
	ID                   apijson.Field
	Amount               apijson.Field
	CreatedAt            apijson.Field
	CreditBlock          apijson.Field
	Currency             apijson.Field
	Customer             apijson.Field
	Description          apijson.Field
	EndingBalance        apijson.Field
	EntryStatus          apijson.Field
	EntryType            apijson.Field
	LedgerSequenceNumber apijson.Field
	Metadata             apijson.Field
	StartingBalance      apijson.Field
	EventID              apijson.Field
	InvoiceID            apijson.Field
	NewBlockExpiryDate   apijson.Field
	PriceID              apijson.Field
	VoidAmount           apijson.Field
	VoidReason           apijson.Field
	raw                  string
	ExtraFields          map[string]apijson.Field
}

func (r creditLedgerEntryModelJSON) RawJSON() string {
	return r.raw
}

func (r *CreditLedgerEntryModel) UnmarshalJSON(data []byte) (err error) {
	*r = CreditLedgerEntryModel{}
	err = apijson.UnmarshalRoot(data, &r.union)
	if err != nil {
		return err
	}
	return apijson.Port(r.union, &r)
}

// AsUnion returns a [CreditLedgerEntryModelUnion] interface which you can cast to
// the specific types for more type safety.
//
// Possible runtime types of the union are
// [shared.CreditLedgerEntryModelIncrementLedgerEntry],
// [shared.CreditLedgerEntryModelDecrementLedgerEntry],
// [shared.CreditLedgerEntryModelExpirationChangeLedgerEntry],
// [shared.CreditLedgerEntryModelCreditBlockExpiryLedgerEntry],
// [shared.CreditLedgerEntryModelVoidLedgerEntry],
// [shared.CreditLedgerEntryModelVoidInitiatedLedgerEntry],
// [shared.CreditLedgerEntryModelAmendmentLedgerEntry].
func (r CreditLedgerEntryModel) AsUnion() CreditLedgerEntryModelUnion {
	return r.union
}

// The [Credit Ledger Entry resource](/product-catalog/prepurchase) models prepaid
// credits within Orb.
//
// Union satisfied by [shared.CreditLedgerEntryModelIncrementLedgerEntry],
// [shared.CreditLedgerEntryModelDecrementLedgerEntry],
// [shared.CreditLedgerEntryModelExpirationChangeLedgerEntry],
// [shared.CreditLedgerEntryModelCreditBlockExpiryLedgerEntry],
// [shared.CreditLedgerEntryModelVoidLedgerEntry],
// [shared.CreditLedgerEntryModelVoidInitiatedLedgerEntry] or
// [shared.CreditLedgerEntryModelAmendmentLedgerEntry].
type CreditLedgerEntryModelUnion interface {
	implementsCreditLedgerEntryModel()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*CreditLedgerEntryModelUnion)(nil)).Elem(),
		"entry_type",
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(CreditLedgerEntryModelIncrementLedgerEntry{}),
			DiscriminatorValue: "increment",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(CreditLedgerEntryModelDecrementLedgerEntry{}),
			DiscriminatorValue: "decrement",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(CreditLedgerEntryModelExpirationChangeLedgerEntry{}),
			DiscriminatorValue: "expiration_change",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(CreditLedgerEntryModelCreditBlockExpiryLedgerEntry{}),
			DiscriminatorValue: "credit_block_expiry",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(CreditLedgerEntryModelVoidLedgerEntry{}),
			DiscriminatorValue: "void",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(CreditLedgerEntryModelVoidInitiatedLedgerEntry{}),
			DiscriminatorValue: "void_initiated",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(CreditLedgerEntryModelAmendmentLedgerEntry{}),
			DiscriminatorValue: "amendment",
		},
	)
}

type CreditLedgerEntryModelIncrementLedgerEntry struct {
	ID                   string                                                `json:"id,required"`
	Amount               float64                                               `json:"amount,required"`
	CreatedAt            time.Time                                             `json:"created_at,required" format:"date-time"`
	CreditBlock          AffectedBlockModel                                    `json:"credit_block,required"`
	Currency             string                                                `json:"currency,required"`
	Customer             CustomerMinifiedModel                                 `json:"customer,required"`
	Description          string                                                `json:"description,required,nullable"`
	EndingBalance        float64                                               `json:"ending_balance,required"`
	EntryStatus          CreditLedgerEntryModelIncrementLedgerEntryEntryStatus `json:"entry_status,required"`
	EntryType            CreditLedgerEntryModelIncrementLedgerEntryEntryType   `json:"entry_type,required"`
	LedgerSequenceNumber int64                                                 `json:"ledger_sequence_number,required"`
	// User specified key-value pairs for the resource. If not present, this defaults
	// to an empty dictionary. Individual keys can be removed by setting the value to
	// `null`, and the entire metadata mapping can be cleared by setting `metadata` to
	// `null`.
	Metadata        map[string]string                              `json:"metadata,required"`
	StartingBalance float64                                        `json:"starting_balance,required"`
	JSON            creditLedgerEntryModelIncrementLedgerEntryJSON `json:"-"`
}

// creditLedgerEntryModelIncrementLedgerEntryJSON contains the JSON metadata for
// the struct [CreditLedgerEntryModelIncrementLedgerEntry]
type creditLedgerEntryModelIncrementLedgerEntryJSON struct {
	ID                   apijson.Field
	Amount               apijson.Field
	CreatedAt            apijson.Field
	CreditBlock          apijson.Field
	Currency             apijson.Field
	Customer             apijson.Field
	Description          apijson.Field
	EndingBalance        apijson.Field
	EntryStatus          apijson.Field
	EntryType            apijson.Field
	LedgerSequenceNumber apijson.Field
	Metadata             apijson.Field
	StartingBalance      apijson.Field
	raw                  string
	ExtraFields          map[string]apijson.Field
}

func (r *CreditLedgerEntryModelIncrementLedgerEntry) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r creditLedgerEntryModelIncrementLedgerEntryJSON) RawJSON() string {
	return r.raw
}

func (r CreditLedgerEntryModelIncrementLedgerEntry) implementsCreditLedgerEntryModel() {}

type CreditLedgerEntryModelIncrementLedgerEntryEntryStatus string

const (
	CreditLedgerEntryModelIncrementLedgerEntryEntryStatusCommitted CreditLedgerEntryModelIncrementLedgerEntryEntryStatus = "committed"
	CreditLedgerEntryModelIncrementLedgerEntryEntryStatusPending   CreditLedgerEntryModelIncrementLedgerEntryEntryStatus = "pending"
)

func (r CreditLedgerEntryModelIncrementLedgerEntryEntryStatus) IsKnown() bool {
	switch r {
	case CreditLedgerEntryModelIncrementLedgerEntryEntryStatusCommitted, CreditLedgerEntryModelIncrementLedgerEntryEntryStatusPending:
		return true
	}
	return false
}

type CreditLedgerEntryModelIncrementLedgerEntryEntryType string

const (
	CreditLedgerEntryModelIncrementLedgerEntryEntryTypeIncrement CreditLedgerEntryModelIncrementLedgerEntryEntryType = "increment"
)

func (r CreditLedgerEntryModelIncrementLedgerEntryEntryType) IsKnown() bool {
	switch r {
	case CreditLedgerEntryModelIncrementLedgerEntryEntryTypeIncrement:
		return true
	}
	return false
}

type CreditLedgerEntryModelDecrementLedgerEntry struct {
	ID                   string                                                `json:"id,required"`
	Amount               float64                                               `json:"amount,required"`
	CreatedAt            time.Time                                             `json:"created_at,required" format:"date-time"`
	CreditBlock          AffectedBlockModel                                    `json:"credit_block,required"`
	Currency             string                                                `json:"currency,required"`
	Customer             CustomerMinifiedModel                                 `json:"customer,required"`
	Description          string                                                `json:"description,required,nullable"`
	EndingBalance        float64                                               `json:"ending_balance,required"`
	EntryStatus          CreditLedgerEntryModelDecrementLedgerEntryEntryStatus `json:"entry_status,required"`
	EntryType            CreditLedgerEntryModelDecrementLedgerEntryEntryType   `json:"entry_type,required"`
	LedgerSequenceNumber int64                                                 `json:"ledger_sequence_number,required"`
	// User specified key-value pairs for the resource. If not present, this defaults
	// to an empty dictionary. Individual keys can be removed by setting the value to
	// `null`, and the entire metadata mapping can be cleared by setting `metadata` to
	// `null`.
	Metadata        map[string]string                              `json:"metadata,required"`
	StartingBalance float64                                        `json:"starting_balance,required"`
	EventID         string                                         `json:"event_id,nullable"`
	InvoiceID       string                                         `json:"invoice_id,nullable"`
	PriceID         string                                         `json:"price_id,nullable"`
	JSON            creditLedgerEntryModelDecrementLedgerEntryJSON `json:"-"`
}

// creditLedgerEntryModelDecrementLedgerEntryJSON contains the JSON metadata for
// the struct [CreditLedgerEntryModelDecrementLedgerEntry]
type creditLedgerEntryModelDecrementLedgerEntryJSON struct {
	ID                   apijson.Field
	Amount               apijson.Field
	CreatedAt            apijson.Field
	CreditBlock          apijson.Field
	Currency             apijson.Field
	Customer             apijson.Field
	Description          apijson.Field
	EndingBalance        apijson.Field
	EntryStatus          apijson.Field
	EntryType            apijson.Field
	LedgerSequenceNumber apijson.Field
	Metadata             apijson.Field
	StartingBalance      apijson.Field
	EventID              apijson.Field
	InvoiceID            apijson.Field
	PriceID              apijson.Field
	raw                  string
	ExtraFields          map[string]apijson.Field
}

func (r *CreditLedgerEntryModelDecrementLedgerEntry) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r creditLedgerEntryModelDecrementLedgerEntryJSON) RawJSON() string {
	return r.raw
}

func (r CreditLedgerEntryModelDecrementLedgerEntry) implementsCreditLedgerEntryModel() {}

type CreditLedgerEntryModelDecrementLedgerEntryEntryStatus string

const (
	CreditLedgerEntryModelDecrementLedgerEntryEntryStatusCommitted CreditLedgerEntryModelDecrementLedgerEntryEntryStatus = "committed"
	CreditLedgerEntryModelDecrementLedgerEntryEntryStatusPending   CreditLedgerEntryModelDecrementLedgerEntryEntryStatus = "pending"
)

func (r CreditLedgerEntryModelDecrementLedgerEntryEntryStatus) IsKnown() bool {
	switch r {
	case CreditLedgerEntryModelDecrementLedgerEntryEntryStatusCommitted, CreditLedgerEntryModelDecrementLedgerEntryEntryStatusPending:
		return true
	}
	return false
}

type CreditLedgerEntryModelDecrementLedgerEntryEntryType string

const (
	CreditLedgerEntryModelDecrementLedgerEntryEntryTypeDecrement CreditLedgerEntryModelDecrementLedgerEntryEntryType = "decrement"
)

func (r CreditLedgerEntryModelDecrementLedgerEntryEntryType) IsKnown() bool {
	switch r {
	case CreditLedgerEntryModelDecrementLedgerEntryEntryTypeDecrement:
		return true
	}
	return false
}

type CreditLedgerEntryModelExpirationChangeLedgerEntry struct {
	ID                   string                                                       `json:"id,required"`
	Amount               float64                                                      `json:"amount,required"`
	CreatedAt            time.Time                                                    `json:"created_at,required" format:"date-time"`
	CreditBlock          AffectedBlockModel                                           `json:"credit_block,required"`
	Currency             string                                                       `json:"currency,required"`
	Customer             CustomerMinifiedModel                                        `json:"customer,required"`
	Description          string                                                       `json:"description,required,nullable"`
	EndingBalance        float64                                                      `json:"ending_balance,required"`
	EntryStatus          CreditLedgerEntryModelExpirationChangeLedgerEntryEntryStatus `json:"entry_status,required"`
	EntryType            CreditLedgerEntryModelExpirationChangeLedgerEntryEntryType   `json:"entry_type,required"`
	LedgerSequenceNumber int64                                                        `json:"ledger_sequence_number,required"`
	// User specified key-value pairs for the resource. If not present, this defaults
	// to an empty dictionary. Individual keys can be removed by setting the value to
	// `null`, and the entire metadata mapping can be cleared by setting `metadata` to
	// `null`.
	Metadata           map[string]string                                     `json:"metadata,required"`
	NewBlockExpiryDate time.Time                                             `json:"new_block_expiry_date,required,nullable" format:"date-time"`
	StartingBalance    float64                                               `json:"starting_balance,required"`
	JSON               creditLedgerEntryModelExpirationChangeLedgerEntryJSON `json:"-"`
}

// creditLedgerEntryModelExpirationChangeLedgerEntryJSON contains the JSON metadata
// for the struct [CreditLedgerEntryModelExpirationChangeLedgerEntry]
type creditLedgerEntryModelExpirationChangeLedgerEntryJSON struct {
	ID                   apijson.Field
	Amount               apijson.Field
	CreatedAt            apijson.Field
	CreditBlock          apijson.Field
	Currency             apijson.Field
	Customer             apijson.Field
	Description          apijson.Field
	EndingBalance        apijson.Field
	EntryStatus          apijson.Field
	EntryType            apijson.Field
	LedgerSequenceNumber apijson.Field
	Metadata             apijson.Field
	NewBlockExpiryDate   apijson.Field
	StartingBalance      apijson.Field
	raw                  string
	ExtraFields          map[string]apijson.Field
}

func (r *CreditLedgerEntryModelExpirationChangeLedgerEntry) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r creditLedgerEntryModelExpirationChangeLedgerEntryJSON) RawJSON() string {
	return r.raw
}

func (r CreditLedgerEntryModelExpirationChangeLedgerEntry) implementsCreditLedgerEntryModel() {}

type CreditLedgerEntryModelExpirationChangeLedgerEntryEntryStatus string

const (
	CreditLedgerEntryModelExpirationChangeLedgerEntryEntryStatusCommitted CreditLedgerEntryModelExpirationChangeLedgerEntryEntryStatus = "committed"
	CreditLedgerEntryModelExpirationChangeLedgerEntryEntryStatusPending   CreditLedgerEntryModelExpirationChangeLedgerEntryEntryStatus = "pending"
)

func (r CreditLedgerEntryModelExpirationChangeLedgerEntryEntryStatus) IsKnown() bool {
	switch r {
	case CreditLedgerEntryModelExpirationChangeLedgerEntryEntryStatusCommitted, CreditLedgerEntryModelExpirationChangeLedgerEntryEntryStatusPending:
		return true
	}
	return false
}

type CreditLedgerEntryModelExpirationChangeLedgerEntryEntryType string

const (
	CreditLedgerEntryModelExpirationChangeLedgerEntryEntryTypeExpirationChange CreditLedgerEntryModelExpirationChangeLedgerEntryEntryType = "expiration_change"
)

func (r CreditLedgerEntryModelExpirationChangeLedgerEntryEntryType) IsKnown() bool {
	switch r {
	case CreditLedgerEntryModelExpirationChangeLedgerEntryEntryTypeExpirationChange:
		return true
	}
	return false
}

type CreditLedgerEntryModelCreditBlockExpiryLedgerEntry struct {
	ID                   string                                                        `json:"id,required"`
	Amount               float64                                                       `json:"amount,required"`
	CreatedAt            time.Time                                                     `json:"created_at,required" format:"date-time"`
	CreditBlock          AffectedBlockModel                                            `json:"credit_block,required"`
	Currency             string                                                        `json:"currency,required"`
	Customer             CustomerMinifiedModel                                         `json:"customer,required"`
	Description          string                                                        `json:"description,required,nullable"`
	EndingBalance        float64                                                       `json:"ending_balance,required"`
	EntryStatus          CreditLedgerEntryModelCreditBlockExpiryLedgerEntryEntryStatus `json:"entry_status,required"`
	EntryType            CreditLedgerEntryModelCreditBlockExpiryLedgerEntryEntryType   `json:"entry_type,required"`
	LedgerSequenceNumber int64                                                         `json:"ledger_sequence_number,required"`
	// User specified key-value pairs for the resource. If not present, this defaults
	// to an empty dictionary. Individual keys can be removed by setting the value to
	// `null`, and the entire metadata mapping can be cleared by setting `metadata` to
	// `null`.
	Metadata        map[string]string                                      `json:"metadata,required"`
	StartingBalance float64                                                `json:"starting_balance,required"`
	JSON            creditLedgerEntryModelCreditBlockExpiryLedgerEntryJSON `json:"-"`
}

// creditLedgerEntryModelCreditBlockExpiryLedgerEntryJSON contains the JSON
// metadata for the struct [CreditLedgerEntryModelCreditBlockExpiryLedgerEntry]
type creditLedgerEntryModelCreditBlockExpiryLedgerEntryJSON struct {
	ID                   apijson.Field
	Amount               apijson.Field
	CreatedAt            apijson.Field
	CreditBlock          apijson.Field
	Currency             apijson.Field
	Customer             apijson.Field
	Description          apijson.Field
	EndingBalance        apijson.Field
	EntryStatus          apijson.Field
	EntryType            apijson.Field
	LedgerSequenceNumber apijson.Field
	Metadata             apijson.Field
	StartingBalance      apijson.Field
	raw                  string
	ExtraFields          map[string]apijson.Field
}

func (r *CreditLedgerEntryModelCreditBlockExpiryLedgerEntry) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r creditLedgerEntryModelCreditBlockExpiryLedgerEntryJSON) RawJSON() string {
	return r.raw
}

func (r CreditLedgerEntryModelCreditBlockExpiryLedgerEntry) implementsCreditLedgerEntryModel() {}

type CreditLedgerEntryModelCreditBlockExpiryLedgerEntryEntryStatus string

const (
	CreditLedgerEntryModelCreditBlockExpiryLedgerEntryEntryStatusCommitted CreditLedgerEntryModelCreditBlockExpiryLedgerEntryEntryStatus = "committed"
	CreditLedgerEntryModelCreditBlockExpiryLedgerEntryEntryStatusPending   CreditLedgerEntryModelCreditBlockExpiryLedgerEntryEntryStatus = "pending"
)

func (r CreditLedgerEntryModelCreditBlockExpiryLedgerEntryEntryStatus) IsKnown() bool {
	switch r {
	case CreditLedgerEntryModelCreditBlockExpiryLedgerEntryEntryStatusCommitted, CreditLedgerEntryModelCreditBlockExpiryLedgerEntryEntryStatusPending:
		return true
	}
	return false
}

type CreditLedgerEntryModelCreditBlockExpiryLedgerEntryEntryType string

const (
	CreditLedgerEntryModelCreditBlockExpiryLedgerEntryEntryTypeCreditBlockExpiry CreditLedgerEntryModelCreditBlockExpiryLedgerEntryEntryType = "credit_block_expiry"
)

func (r CreditLedgerEntryModelCreditBlockExpiryLedgerEntryEntryType) IsKnown() bool {
	switch r {
	case CreditLedgerEntryModelCreditBlockExpiryLedgerEntryEntryTypeCreditBlockExpiry:
		return true
	}
	return false
}

type CreditLedgerEntryModelVoidLedgerEntry struct {
	ID                   string                                           `json:"id,required"`
	Amount               float64                                          `json:"amount,required"`
	CreatedAt            time.Time                                        `json:"created_at,required" format:"date-time"`
	CreditBlock          AffectedBlockModel                               `json:"credit_block,required"`
	Currency             string                                           `json:"currency,required"`
	Customer             CustomerMinifiedModel                            `json:"customer,required"`
	Description          string                                           `json:"description,required,nullable"`
	EndingBalance        float64                                          `json:"ending_balance,required"`
	EntryStatus          CreditLedgerEntryModelVoidLedgerEntryEntryStatus `json:"entry_status,required"`
	EntryType            CreditLedgerEntryModelVoidLedgerEntryEntryType   `json:"entry_type,required"`
	LedgerSequenceNumber int64                                            `json:"ledger_sequence_number,required"`
	// User specified key-value pairs for the resource. If not present, this defaults
	// to an empty dictionary. Individual keys can be removed by setting the value to
	// `null`, and the entire metadata mapping can be cleared by setting `metadata` to
	// `null`.
	Metadata        map[string]string                         `json:"metadata,required"`
	StartingBalance float64                                   `json:"starting_balance,required"`
	VoidAmount      float64                                   `json:"void_amount,required"`
	VoidReason      string                                    `json:"void_reason,required,nullable"`
	JSON            creditLedgerEntryModelVoidLedgerEntryJSON `json:"-"`
}

// creditLedgerEntryModelVoidLedgerEntryJSON contains the JSON metadata for the
// struct [CreditLedgerEntryModelVoidLedgerEntry]
type creditLedgerEntryModelVoidLedgerEntryJSON struct {
	ID                   apijson.Field
	Amount               apijson.Field
	CreatedAt            apijson.Field
	CreditBlock          apijson.Field
	Currency             apijson.Field
	Customer             apijson.Field
	Description          apijson.Field
	EndingBalance        apijson.Field
	EntryStatus          apijson.Field
	EntryType            apijson.Field
	LedgerSequenceNumber apijson.Field
	Metadata             apijson.Field
	StartingBalance      apijson.Field
	VoidAmount           apijson.Field
	VoidReason           apijson.Field
	raw                  string
	ExtraFields          map[string]apijson.Field
}

func (r *CreditLedgerEntryModelVoidLedgerEntry) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r creditLedgerEntryModelVoidLedgerEntryJSON) RawJSON() string {
	return r.raw
}

func (r CreditLedgerEntryModelVoidLedgerEntry) implementsCreditLedgerEntryModel() {}

type CreditLedgerEntryModelVoidLedgerEntryEntryStatus string

const (
	CreditLedgerEntryModelVoidLedgerEntryEntryStatusCommitted CreditLedgerEntryModelVoidLedgerEntryEntryStatus = "committed"
	CreditLedgerEntryModelVoidLedgerEntryEntryStatusPending   CreditLedgerEntryModelVoidLedgerEntryEntryStatus = "pending"
)

func (r CreditLedgerEntryModelVoidLedgerEntryEntryStatus) IsKnown() bool {
	switch r {
	case CreditLedgerEntryModelVoidLedgerEntryEntryStatusCommitted, CreditLedgerEntryModelVoidLedgerEntryEntryStatusPending:
		return true
	}
	return false
}

type CreditLedgerEntryModelVoidLedgerEntryEntryType string

const (
	CreditLedgerEntryModelVoidLedgerEntryEntryTypeVoid CreditLedgerEntryModelVoidLedgerEntryEntryType = "void"
)

func (r CreditLedgerEntryModelVoidLedgerEntryEntryType) IsKnown() bool {
	switch r {
	case CreditLedgerEntryModelVoidLedgerEntryEntryTypeVoid:
		return true
	}
	return false
}

type CreditLedgerEntryModelVoidInitiatedLedgerEntry struct {
	ID                   string                                                    `json:"id,required"`
	Amount               float64                                                   `json:"amount,required"`
	CreatedAt            time.Time                                                 `json:"created_at,required" format:"date-time"`
	CreditBlock          AffectedBlockModel                                        `json:"credit_block,required"`
	Currency             string                                                    `json:"currency,required"`
	Customer             CustomerMinifiedModel                                     `json:"customer,required"`
	Description          string                                                    `json:"description,required,nullable"`
	EndingBalance        float64                                                   `json:"ending_balance,required"`
	EntryStatus          CreditLedgerEntryModelVoidInitiatedLedgerEntryEntryStatus `json:"entry_status,required"`
	EntryType            CreditLedgerEntryModelVoidInitiatedLedgerEntryEntryType   `json:"entry_type,required"`
	LedgerSequenceNumber int64                                                     `json:"ledger_sequence_number,required"`
	// User specified key-value pairs for the resource. If not present, this defaults
	// to an empty dictionary. Individual keys can be removed by setting the value to
	// `null`, and the entire metadata mapping can be cleared by setting `metadata` to
	// `null`.
	Metadata           map[string]string                                  `json:"metadata,required"`
	NewBlockExpiryDate time.Time                                          `json:"new_block_expiry_date,required" format:"date-time"`
	StartingBalance    float64                                            `json:"starting_balance,required"`
	VoidAmount         float64                                            `json:"void_amount,required"`
	VoidReason         string                                             `json:"void_reason,required,nullable"`
	JSON               creditLedgerEntryModelVoidInitiatedLedgerEntryJSON `json:"-"`
}

// creditLedgerEntryModelVoidInitiatedLedgerEntryJSON contains the JSON metadata
// for the struct [CreditLedgerEntryModelVoidInitiatedLedgerEntry]
type creditLedgerEntryModelVoidInitiatedLedgerEntryJSON struct {
	ID                   apijson.Field
	Amount               apijson.Field
	CreatedAt            apijson.Field
	CreditBlock          apijson.Field
	Currency             apijson.Field
	Customer             apijson.Field
	Description          apijson.Field
	EndingBalance        apijson.Field
	EntryStatus          apijson.Field
	EntryType            apijson.Field
	LedgerSequenceNumber apijson.Field
	Metadata             apijson.Field
	NewBlockExpiryDate   apijson.Field
	StartingBalance      apijson.Field
	VoidAmount           apijson.Field
	VoidReason           apijson.Field
	raw                  string
	ExtraFields          map[string]apijson.Field
}

func (r *CreditLedgerEntryModelVoidInitiatedLedgerEntry) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r creditLedgerEntryModelVoidInitiatedLedgerEntryJSON) RawJSON() string {
	return r.raw
}

func (r CreditLedgerEntryModelVoidInitiatedLedgerEntry) implementsCreditLedgerEntryModel() {}

type CreditLedgerEntryModelVoidInitiatedLedgerEntryEntryStatus string

const (
	CreditLedgerEntryModelVoidInitiatedLedgerEntryEntryStatusCommitted CreditLedgerEntryModelVoidInitiatedLedgerEntryEntryStatus = "committed"
	CreditLedgerEntryModelVoidInitiatedLedgerEntryEntryStatusPending   CreditLedgerEntryModelVoidInitiatedLedgerEntryEntryStatus = "pending"
)

func (r CreditLedgerEntryModelVoidInitiatedLedgerEntryEntryStatus) IsKnown() bool {
	switch r {
	case CreditLedgerEntryModelVoidInitiatedLedgerEntryEntryStatusCommitted, CreditLedgerEntryModelVoidInitiatedLedgerEntryEntryStatusPending:
		return true
	}
	return false
}

type CreditLedgerEntryModelVoidInitiatedLedgerEntryEntryType string

const (
	CreditLedgerEntryModelVoidInitiatedLedgerEntryEntryTypeVoidInitiated CreditLedgerEntryModelVoidInitiatedLedgerEntryEntryType = "void_initiated"
)

func (r CreditLedgerEntryModelVoidInitiatedLedgerEntryEntryType) IsKnown() bool {
	switch r {
	case CreditLedgerEntryModelVoidInitiatedLedgerEntryEntryTypeVoidInitiated:
		return true
	}
	return false
}

type CreditLedgerEntryModelAmendmentLedgerEntry struct {
	ID                   string                                                `json:"id,required"`
	Amount               float64                                               `json:"amount,required"`
	CreatedAt            time.Time                                             `json:"created_at,required" format:"date-time"`
	CreditBlock          AffectedBlockModel                                    `json:"credit_block,required"`
	Currency             string                                                `json:"currency,required"`
	Customer             CustomerMinifiedModel                                 `json:"customer,required"`
	Description          string                                                `json:"description,required,nullable"`
	EndingBalance        float64                                               `json:"ending_balance,required"`
	EntryStatus          CreditLedgerEntryModelAmendmentLedgerEntryEntryStatus `json:"entry_status,required"`
	EntryType            CreditLedgerEntryModelAmendmentLedgerEntryEntryType   `json:"entry_type,required"`
	LedgerSequenceNumber int64                                                 `json:"ledger_sequence_number,required"`
	// User specified key-value pairs for the resource. If not present, this defaults
	// to an empty dictionary. Individual keys can be removed by setting the value to
	// `null`, and the entire metadata mapping can be cleared by setting `metadata` to
	// `null`.
	Metadata        map[string]string                              `json:"metadata,required"`
	StartingBalance float64                                        `json:"starting_balance,required"`
	JSON            creditLedgerEntryModelAmendmentLedgerEntryJSON `json:"-"`
}

// creditLedgerEntryModelAmendmentLedgerEntryJSON contains the JSON metadata for
// the struct [CreditLedgerEntryModelAmendmentLedgerEntry]
type creditLedgerEntryModelAmendmentLedgerEntryJSON struct {
	ID                   apijson.Field
	Amount               apijson.Field
	CreatedAt            apijson.Field
	CreditBlock          apijson.Field
	Currency             apijson.Field
	Customer             apijson.Field
	Description          apijson.Field
	EndingBalance        apijson.Field
	EntryStatus          apijson.Field
	EntryType            apijson.Field
	LedgerSequenceNumber apijson.Field
	Metadata             apijson.Field
	StartingBalance      apijson.Field
	raw                  string
	ExtraFields          map[string]apijson.Field
}

func (r *CreditLedgerEntryModelAmendmentLedgerEntry) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r creditLedgerEntryModelAmendmentLedgerEntryJSON) RawJSON() string {
	return r.raw
}

func (r CreditLedgerEntryModelAmendmentLedgerEntry) implementsCreditLedgerEntryModel() {}

type CreditLedgerEntryModelAmendmentLedgerEntryEntryStatus string

const (
	CreditLedgerEntryModelAmendmentLedgerEntryEntryStatusCommitted CreditLedgerEntryModelAmendmentLedgerEntryEntryStatus = "committed"
	CreditLedgerEntryModelAmendmentLedgerEntryEntryStatusPending   CreditLedgerEntryModelAmendmentLedgerEntryEntryStatus = "pending"
)

func (r CreditLedgerEntryModelAmendmentLedgerEntryEntryStatus) IsKnown() bool {
	switch r {
	case CreditLedgerEntryModelAmendmentLedgerEntryEntryStatusCommitted, CreditLedgerEntryModelAmendmentLedgerEntryEntryStatusPending:
		return true
	}
	return false
}

type CreditLedgerEntryModelAmendmentLedgerEntryEntryType string

const (
	CreditLedgerEntryModelAmendmentLedgerEntryEntryTypeAmendment CreditLedgerEntryModelAmendmentLedgerEntryEntryType = "amendment"
)

func (r CreditLedgerEntryModelAmendmentLedgerEntryEntryType) IsKnown() bool {
	switch r {
	case CreditLedgerEntryModelAmendmentLedgerEntryEntryTypeAmendment:
		return true
	}
	return false
}

type CreditLedgerEntryModelEntryStatus string

const (
	CreditLedgerEntryModelEntryStatusCommitted CreditLedgerEntryModelEntryStatus = "committed"
	CreditLedgerEntryModelEntryStatusPending   CreditLedgerEntryModelEntryStatus = "pending"
)

func (r CreditLedgerEntryModelEntryStatus) IsKnown() bool {
	switch r {
	case CreditLedgerEntryModelEntryStatusCommitted, CreditLedgerEntryModelEntryStatusPending:
		return true
	}
	return false
}

type CreditLedgerEntryModelEntryType string

const (
	CreditLedgerEntryModelEntryTypeIncrement         CreditLedgerEntryModelEntryType = "increment"
	CreditLedgerEntryModelEntryTypeDecrement         CreditLedgerEntryModelEntryType = "decrement"
	CreditLedgerEntryModelEntryTypeExpirationChange  CreditLedgerEntryModelEntryType = "expiration_change"
	CreditLedgerEntryModelEntryTypeCreditBlockExpiry CreditLedgerEntryModelEntryType = "credit_block_expiry"
	CreditLedgerEntryModelEntryTypeVoid              CreditLedgerEntryModelEntryType = "void"
	CreditLedgerEntryModelEntryTypeVoidInitiated     CreditLedgerEntryModelEntryType = "void_initiated"
	CreditLedgerEntryModelEntryTypeAmendment         CreditLedgerEntryModelEntryType = "amendment"
)

func (r CreditLedgerEntryModelEntryType) IsKnown() bool {
	switch r {
	case CreditLedgerEntryModelEntryTypeIncrement, CreditLedgerEntryModelEntryTypeDecrement, CreditLedgerEntryModelEntryTypeExpirationChange, CreditLedgerEntryModelEntryTypeCreditBlockExpiry, CreditLedgerEntryModelEntryTypeVoid, CreditLedgerEntryModelEntryTypeVoidInitiated, CreditLedgerEntryModelEntryTypeAmendment:
		return true
	}
	return false
}

type CreditNoteDiscountModel struct {
	AmountApplied      string                                  `json:"amount_applied,required"`
	DiscountType       CreditNoteDiscountModelDiscountType     `json:"discount_type,required"`
	PercentageDiscount float64                                 `json:"percentage_discount,required"`
	AppliesToPrices    []CreditNoteDiscountModelAppliesToPrice `json:"applies_to_prices,nullable"`
	Reason             string                                  `json:"reason,nullable"`
	JSON               creditNoteDiscountModelJSON             `json:"-"`
}

// creditNoteDiscountModelJSON contains the JSON metadata for the struct
// [CreditNoteDiscountModel]
type creditNoteDiscountModelJSON struct {
	AmountApplied      apijson.Field
	DiscountType       apijson.Field
	PercentageDiscount apijson.Field
	AppliesToPrices    apijson.Field
	Reason             apijson.Field
	raw                string
	ExtraFields        map[string]apijson.Field
}

func (r *CreditNoteDiscountModel) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r creditNoteDiscountModelJSON) RawJSON() string {
	return r.raw
}

type CreditNoteDiscountModelDiscountType string

const (
	CreditNoteDiscountModelDiscountTypePercentage CreditNoteDiscountModelDiscountType = "percentage"
)

func (r CreditNoteDiscountModelDiscountType) IsKnown() bool {
	switch r {
	case CreditNoteDiscountModelDiscountTypePercentage:
		return true
	}
	return false
}

type CreditNoteDiscountModelAppliesToPrice struct {
	ID   string                                    `json:"id,required"`
	Name string                                    `json:"name,required"`
	JSON creditNoteDiscountModelAppliesToPriceJSON `json:"-"`
}

// creditNoteDiscountModelAppliesToPriceJSON contains the JSON metadata for the
// struct [CreditNoteDiscountModelAppliesToPrice]
type creditNoteDiscountModelAppliesToPriceJSON struct {
	ID          apijson.Field
	Name        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *CreditNoteDiscountModelAppliesToPrice) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r creditNoteDiscountModelAppliesToPriceJSON) RawJSON() string {
	return r.raw
}

// The [Credit Note](/invoicing/credit-notes) resource represents a credit that has
// been applied to a particular invoice.
type CreditNoteModel struct {
	// The Orb id of this credit note.
	ID string `json:"id,required"`
	// The creation time of the resource in Orb.
	CreatedAt time.Time `json:"created_at,required" format:"date-time"`
	// The unique identifier for credit notes.
	CreditNoteNumber string `json:"credit_note_number,required"`
	// A URL to a PDF of the credit note.
	CreditNotePdf string                `json:"credit_note_pdf,required,nullable"`
	Customer      CustomerMinifiedModel `json:"customer,required"`
	// The id of the invoice resource that this credit note is applied to.
	InvoiceID string `json:"invoice_id,required"`
	// All of the line items associated with this credit note.
	LineItems []CreditNoteModelLineItem `json:"line_items,required"`
	// The maximum amount applied on the original invoice
	MaximumAmountAdjustment CreditNoteDiscountModel `json:"maximum_amount_adjustment,required,nullable"`
	// An optional memo supplied on the credit note.
	Memo string `json:"memo,required,nullable"`
	// Any credited amount from the applied minimum on the invoice.
	MinimumAmountRefunded string                `json:"minimum_amount_refunded,required,nullable"`
	Reason                CreditNoteModelReason `json:"reason,required,nullable"`
	// The total prior to any creditable invoice-level discounts or minimums.
	Subtotal string `json:"subtotal,required"`
	// The total including creditable invoice-level discounts or minimums, and tax.
	Total string              `json:"total,required"`
	Type  CreditNoteModelType `json:"type,required"`
	// The time at which the credit note was voided in Orb, if applicable.
	VoidedAt time.Time `json:"voided_at,required,nullable" format:"date-time"`
	// Any discounts applied on the original invoice.
	Discounts []CreditNoteDiscountModel `json:"discounts"`
	JSON      creditNoteModelJSON       `json:"-"`
}

// creditNoteModelJSON contains the JSON metadata for the struct [CreditNoteModel]
type creditNoteModelJSON struct {
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

func (r *CreditNoteModel) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r creditNoteModelJSON) RawJSON() string {
	return r.raw
}

type CreditNoteModelLineItem struct {
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
	TaxAmounts []TaxAmountModel `json:"tax_amounts,required"`
	// Any line item discounts from the invoice's line item.
	Discounts []CreditNoteModelLineItemsDiscount `json:"discounts"`
	JSON      creditNoteModelLineItemJSON        `json:"-"`
}

// creditNoteModelLineItemJSON contains the JSON metadata for the struct
// [CreditNoteModelLineItem]
type creditNoteModelLineItemJSON struct {
	ID          apijson.Field
	Amount      apijson.Field
	ItemID      apijson.Field
	Name        apijson.Field
	Quantity    apijson.Field
	Subtotal    apijson.Field
	TaxAmounts  apijson.Field
	Discounts   apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *CreditNoteModelLineItem) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r creditNoteModelLineItemJSON) RawJSON() string {
	return r.raw
}

type CreditNoteModelLineItemsDiscount struct {
	ID                 string                                        `json:"id,required"`
	AmountApplied      string                                        `json:"amount_applied,required"`
	AppliesToPriceIDs  []string                                      `json:"applies_to_price_ids,required"`
	DiscountType       CreditNoteModelLineItemsDiscountsDiscountType `json:"discount_type,required"`
	PercentageDiscount float64                                       `json:"percentage_discount,required"`
	AmountDiscount     string                                        `json:"amount_discount,nullable"`
	Reason             string                                        `json:"reason,nullable"`
	JSON               creditNoteModelLineItemsDiscountJSON          `json:"-"`
}

// creditNoteModelLineItemsDiscountJSON contains the JSON metadata for the struct
// [CreditNoteModelLineItemsDiscount]
type creditNoteModelLineItemsDiscountJSON struct {
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

func (r *CreditNoteModelLineItemsDiscount) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r creditNoteModelLineItemsDiscountJSON) RawJSON() string {
	return r.raw
}

type CreditNoteModelLineItemsDiscountsDiscountType string

const (
	CreditNoteModelLineItemsDiscountsDiscountTypePercentage CreditNoteModelLineItemsDiscountsDiscountType = "percentage"
	CreditNoteModelLineItemsDiscountsDiscountTypeAmount     CreditNoteModelLineItemsDiscountsDiscountType = "amount"
)

func (r CreditNoteModelLineItemsDiscountsDiscountType) IsKnown() bool {
	switch r {
	case CreditNoteModelLineItemsDiscountsDiscountTypePercentage, CreditNoteModelLineItemsDiscountsDiscountTypeAmount:
		return true
	}
	return false
}

type CreditNoteModelReason string

const (
	CreditNoteModelReasonDuplicate             CreditNoteModelReason = "Duplicate"
	CreditNoteModelReasonFraudulent            CreditNoteModelReason = "Fraudulent"
	CreditNoteModelReasonOrderChange           CreditNoteModelReason = "Order change"
	CreditNoteModelReasonProductUnsatisfactory CreditNoteModelReason = "Product unsatisfactory"
)

func (r CreditNoteModelReason) IsKnown() bool {
	switch r {
	case CreditNoteModelReasonDuplicate, CreditNoteModelReasonFraudulent, CreditNoteModelReasonOrderChange, CreditNoteModelReasonProductUnsatisfactory:
		return true
	}
	return false
}

type CreditNoteModelType string

const (
	CreditNoteModelTypeRefund     CreditNoteModelType = "refund"
	CreditNoteModelTypeAdjustment CreditNoteModelType = "adjustment"
)

func (r CreditNoteModelType) IsKnown() bool {
	switch r {
	case CreditNoteModelTypeRefund, CreditNoteModelTypeAdjustment:
		return true
	}
	return false
}

type CreditNoteSummaryModel struct {
	ID               string `json:"id,required"`
	CreditNoteNumber string `json:"credit_note_number,required"`
	// An optional memo supplied on the credit note.
	Memo   string `json:"memo,required,nullable"`
	Reason string `json:"reason,required"`
	Total  string `json:"total,required"`
	Type   string `json:"type,required"`
	// If the credit note has a status of `void`, this gives a timestamp when the
	// credit note was voided.
	VoidedAt time.Time                  `json:"voided_at,required,nullable" format:"date-time"`
	JSON     creditNoteSummaryModelJSON `json:"-"`
}

// creditNoteSummaryModelJSON contains the JSON metadata for the struct
// [CreditNoteSummaryModel]
type creditNoteSummaryModelJSON struct {
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

func (r *CreditNoteSummaryModel) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r creditNoteSummaryModelJSON) RawJSON() string {
	return r.raw
}

type CustomRatingFunctionConfigModel map[string]interface{}

type CustomRatingFunctionConfigModelParam map[string]interface{}

type CustomerBalanceTransactionModel struct {
	// A unique id for this transaction.
	ID     string                                `json:"id,required"`
	Action CustomerBalanceTransactionModelAction `json:"action,required"`
	// The value of the amount changed in the transaction.
	Amount string `json:"amount,required"`
	// The creation time of this transaction.
	CreatedAt  time.Time                                 `json:"created_at,required" format:"date-time"`
	CreditNote CustomerBalanceTransactionModelCreditNote `json:"credit_note,required,nullable"`
	// An optional description provided for manual customer balance adjustments.
	Description string `json:"description,required,nullable"`
	// The new value of the customer's balance prior to the transaction, in the
	// customer's currency.
	EndingBalance string                                 `json:"ending_balance,required"`
	Invoice       CustomerBalanceTransactionModelInvoice `json:"invoice,required,nullable"`
	// The original value of the customer's balance prior to the transaction, in the
	// customer's currency.
	StartingBalance string                              `json:"starting_balance,required"`
	Type            CustomerBalanceTransactionModelType `json:"type,required"`
	JSON            customerBalanceTransactionModelJSON `json:"-"`
}

// customerBalanceTransactionModelJSON contains the JSON metadata for the struct
// [CustomerBalanceTransactionModel]
type customerBalanceTransactionModelJSON struct {
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

func (r *CustomerBalanceTransactionModel) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r customerBalanceTransactionModelJSON) RawJSON() string {
	return r.raw
}

type CustomerBalanceTransactionModelAction string

const (
	CustomerBalanceTransactionModelActionAppliedToInvoice     CustomerBalanceTransactionModelAction = "applied_to_invoice"
	CustomerBalanceTransactionModelActionManualAdjustment     CustomerBalanceTransactionModelAction = "manual_adjustment"
	CustomerBalanceTransactionModelActionProratedRefund       CustomerBalanceTransactionModelAction = "prorated_refund"
	CustomerBalanceTransactionModelActionRevertProratedRefund CustomerBalanceTransactionModelAction = "revert_prorated_refund"
	CustomerBalanceTransactionModelActionReturnFromVoiding    CustomerBalanceTransactionModelAction = "return_from_voiding"
	CustomerBalanceTransactionModelActionCreditNoteApplied    CustomerBalanceTransactionModelAction = "credit_note_applied"
	CustomerBalanceTransactionModelActionCreditNoteVoided     CustomerBalanceTransactionModelAction = "credit_note_voided"
	CustomerBalanceTransactionModelActionOverpaymentRefund    CustomerBalanceTransactionModelAction = "overpayment_refund"
	CustomerBalanceTransactionModelActionExternalPayment      CustomerBalanceTransactionModelAction = "external_payment"
)

func (r CustomerBalanceTransactionModelAction) IsKnown() bool {
	switch r {
	case CustomerBalanceTransactionModelActionAppliedToInvoice, CustomerBalanceTransactionModelActionManualAdjustment, CustomerBalanceTransactionModelActionProratedRefund, CustomerBalanceTransactionModelActionRevertProratedRefund, CustomerBalanceTransactionModelActionReturnFromVoiding, CustomerBalanceTransactionModelActionCreditNoteApplied, CustomerBalanceTransactionModelActionCreditNoteVoided, CustomerBalanceTransactionModelActionOverpaymentRefund, CustomerBalanceTransactionModelActionExternalPayment:
		return true
	}
	return false
}

type CustomerBalanceTransactionModelCreditNote struct {
	// The id of the Credit note
	ID   string                                        `json:"id,required"`
	JSON customerBalanceTransactionModelCreditNoteJSON `json:"-"`
}

// customerBalanceTransactionModelCreditNoteJSON contains the JSON metadata for the
// struct [CustomerBalanceTransactionModelCreditNote]
type customerBalanceTransactionModelCreditNoteJSON struct {
	ID          apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *CustomerBalanceTransactionModelCreditNote) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r customerBalanceTransactionModelCreditNoteJSON) RawJSON() string {
	return r.raw
}

type CustomerBalanceTransactionModelInvoice struct {
	// The Invoice id
	ID   string                                     `json:"id,required"`
	JSON customerBalanceTransactionModelInvoiceJSON `json:"-"`
}

// customerBalanceTransactionModelInvoiceJSON contains the JSON metadata for the
// struct [CustomerBalanceTransactionModelInvoice]
type customerBalanceTransactionModelInvoiceJSON struct {
	ID          apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *CustomerBalanceTransactionModelInvoice) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r customerBalanceTransactionModelInvoiceJSON) RawJSON() string {
	return r.raw
}

type CustomerBalanceTransactionModelType string

const (
	CustomerBalanceTransactionModelTypeIncrement CustomerBalanceTransactionModelType = "increment"
	CustomerBalanceTransactionModelTypeDecrement CustomerBalanceTransactionModelType = "decrement"
)

func (r CustomerBalanceTransactionModelType) IsKnown() bool {
	switch r {
	case CustomerBalanceTransactionModelTypeIncrement, CustomerBalanceTransactionModelTypeDecrement:
		return true
	}
	return false
}

type CustomerCostsModel struct {
	Data []AggregatedCostModel  `json:"data,required"`
	JSON customerCostsModelJSON `json:"-"`
}

// customerCostsModelJSON contains the JSON metadata for the struct
// [CustomerCostsModel]
type customerCostsModelJSON struct {
	Data        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *CustomerCostsModel) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r customerCostsModelJSON) RawJSON() string {
	return r.raw
}

type CustomerCreditBalancesModel struct {
	Data               []CustomerCreditBalancesModelData `json:"data,required"`
	PaginationMetadata PaginationMetadata                `json:"pagination_metadata,required"`
	JSON               customerCreditBalancesModelJSON   `json:"-"`
}

// customerCreditBalancesModelJSON contains the JSON metadata for the struct
// [CustomerCreditBalancesModel]
type customerCreditBalancesModelJSON struct {
	Data               apijson.Field
	PaginationMetadata apijson.Field
	raw                string
	ExtraFields        map[string]apijson.Field
}

func (r *CustomerCreditBalancesModel) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r customerCreditBalancesModelJSON) RawJSON() string {
	return r.raw
}

type CustomerCreditBalancesModelData struct {
	ID                    string                                `json:"id,required"`
	Balance               float64                               `json:"balance,required"`
	EffectiveDate         time.Time                             `json:"effective_date,required,nullable" format:"date-time"`
	ExpiryDate            time.Time                             `json:"expiry_date,required,nullable" format:"date-time"`
	MaximumInitialBalance float64                               `json:"maximum_initial_balance,required,nullable"`
	PerUnitCostBasis      string                                `json:"per_unit_cost_basis,required,nullable"`
	Status                CustomerCreditBalancesModelDataStatus `json:"status,required"`
	JSON                  customerCreditBalancesModelDataJSON   `json:"-"`
}

// customerCreditBalancesModelDataJSON contains the JSON metadata for the struct
// [CustomerCreditBalancesModelData]
type customerCreditBalancesModelDataJSON struct {
	ID                    apijson.Field
	Balance               apijson.Field
	EffectiveDate         apijson.Field
	ExpiryDate            apijson.Field
	MaximumInitialBalance apijson.Field
	PerUnitCostBasis      apijson.Field
	Status                apijson.Field
	raw                   string
	ExtraFields           map[string]apijson.Field
}

func (r *CustomerCreditBalancesModelData) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r customerCreditBalancesModelDataJSON) RawJSON() string {
	return r.raw
}

type CustomerCreditBalancesModelDataStatus string

const (
	CustomerCreditBalancesModelDataStatusActive         CustomerCreditBalancesModelDataStatus = "active"
	CustomerCreditBalancesModelDataStatusPendingPayment CustomerCreditBalancesModelDataStatus = "pending_payment"
)

func (r CustomerCreditBalancesModelDataStatus) IsKnown() bool {
	switch r {
	case CustomerCreditBalancesModelDataStatusActive, CustomerCreditBalancesModelDataStatusPendingPayment:
		return true
	}
	return false
}

type CustomerHierarchyConfigModelParam struct {
	// A list of child customer IDs to add to the hierarchy. The desired child
	// customers must not already be part of another hierarchy.
	ChildCustomerIDs param.Field[[]string] `json:"child_customer_ids"`
	// The ID of the parent customer in the hierarchy. The desired parent customer must
	// not be a child of another customer.
	ParentCustomerID param.Field[string] `json:"parent_customer_id"`
}

func (r CustomerHierarchyConfigModelParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type CustomerMinifiedModel struct {
	ID                 string                    `json:"id,required"`
	ExternalCustomerID string                    `json:"external_customer_id,required,nullable"`
	JSON               customerMinifiedModelJSON `json:"-"`
}

// customerMinifiedModelJSON contains the JSON metadata for the struct
// [CustomerMinifiedModel]
type customerMinifiedModelJSON struct {
	ID                 apijson.Field
	ExternalCustomerID apijson.Field
	raw                string
	ExtraFields        map[string]apijson.Field
}

func (r *CustomerMinifiedModel) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r customerMinifiedModelJSON) RawJSON() string {
	return r.raw
}

// A customer is a buyer of your products, and the other party to the billing
// relationship.
//
// In Orb, customers are assigned system generated identifiers automatically, but
// it's often desirable to have these match existing identifiers in your system. To
// avoid having to denormalize Orb ID information, you can pass in an
// `external_customer_id` with your own identifier. See
// [Customer ID Aliases](/events-and-metrics/customer-aliases) for further
// information about how these aliases work in Orb.
//
// In addition to having an identifier in your system, a customer may exist in a
// payment provider solution like Stripe. Use the `payment_provider_id` and the
// `payment_provider` enum field to express this mapping.
//
// A customer also has a timezone (from the standard
// [IANA timezone database](https://www.iana.org/time-zones)), which defaults to
// your account's timezone. See [Timezone localization](/essentials/timezones) for
// information on what this timezone parameter influences within Orb.
type CustomerModel struct {
	ID               string   `json:"id,required"`
	AdditionalEmails []string `json:"additional_emails,required"`
	AutoCollection   bool     `json:"auto_collection,required"`
	// The customer's current balance in their currency.
	Balance        string       `json:"balance,required"`
	BillingAddress AddressModel `json:"billing_address,required,nullable"`
	CreatedAt      time.Time    `json:"created_at,required" format:"date-time"`
	Currency       string       `json:"currency,required,nullable"`
	// A valid customer email, to be used for notifications. When Orb triggers payment
	// through a payment gateway, this email will be used for any automatically issued
	// receipts.
	Email                  string `json:"email,required"`
	EmailDelivery          bool   `json:"email_delivery,required"`
	ExemptFromAutomatedTax bool   `json:"exempt_from_automated_tax,required,nullable"`
	// An optional user-defined ID for this customer resource, used throughout the
	// system as an alias for this Customer. Use this field to identify a customer by
	// an existing identifier in your system.
	ExternalCustomerID string `json:"external_customer_id,required,nullable"`
	// The hierarchical relationships for this customer.
	Hierarchy CustomerModelHierarchy `json:"hierarchy,required"`
	// User specified key-value pairs for the resource. If not present, this defaults
	// to an empty dictionary. Individual keys can be removed by setting the value to
	// `null`, and the entire metadata mapping can be cleared by setting `metadata` to
	// `null`.
	Metadata map[string]string `json:"metadata,required"`
	// The full name of the customer
	Name string `json:"name,required"`
	// This is used for creating charges or invoices in an external system via Orb.
	// When not in test mode, the connection must first be configured in the Orb
	// webapp.
	PaymentProvider CustomerModelPaymentProvider `json:"payment_provider,required,nullable"`
	// The ID of this customer in an external payments solution, such as Stripe. This
	// is used for creating charges or invoices in the external system via Orb.
	PaymentProviderID string       `json:"payment_provider_id,required,nullable"`
	PortalURL         string       `json:"portal_url,required,nullable"`
	ShippingAddress   AddressModel `json:"shipping_address,required,nullable"`
	// Tax IDs are commonly required to be displayed on customer invoices, which are
	// added to the headers of invoices.
	//
	// ### Supported Tax ID Countries and Types
	//
	// | Country              | Type         | Description                                                                                             |
	// | -------------------- | ------------ | ------------------------------------------------------------------------------------------------------- |
	// | Andorra              | `ad_nrt`     | Andorran NRT Number                                                                                     |
	// | Argentina            | `ar_cuit`    | Argentinian Tax ID Number                                                                               |
	// | Australia            | `au_abn`     | Australian Business Number (AU ABN)                                                                     |
	// | Australia            | `au_arn`     | Australian Taxation Office Reference Number                                                             |
	// | Austria              | `eu_vat`     | European VAT Number                                                                                     |
	// | Bahrain              | `bh_vat`     | Bahraini VAT Number                                                                                     |
	// | Belgium              | `eu_vat`     | European VAT Number                                                                                     |
	// | Bolivia              | `bo_tin`     | Bolivian Tax ID                                                                                         |
	// | Brazil               | `br_cnpj`    | Brazilian CNPJ Number                                                                                   |
	// | Brazil               | `br_cpf`     | Brazilian CPF Number                                                                                    |
	// | Bulgaria             | `bg_uic`     | Bulgaria Unified Identification Code                                                                    |
	// | Bulgaria             | `eu_vat`     | European VAT Number                                                                                     |
	// | Canada               | `ca_bn`      | Canadian BN                                                                                             |
	// | Canada               | `ca_gst_hst` | Canadian GST/HST Number                                                                                 |
	// | Canada               | `ca_pst_bc`  | Canadian PST Number (British Columbia)                                                                  |
	// | Canada               | `ca_pst_mb`  | Canadian PST Number (Manitoba)                                                                          |
	// | Canada               | `ca_pst_sk`  | Canadian PST Number (Saskatchewan)                                                                      |
	// | Canada               | `ca_qst`     | Canadian QST Number (Québec)                                                                            |
	// | Chile                | `cl_tin`     | Chilean TIN                                                                                             |
	// | China                | `cn_tin`     | Chinese Tax ID                                                                                          |
	// | Colombia             | `co_nit`     | Colombian NIT Number                                                                                    |
	// | Costa Rica           | `cr_tin`     | Costa Rican Tax ID                                                                                      |
	// | Croatia              | `eu_vat`     | European VAT Number                                                                                     |
	// | Cyprus               | `eu_vat`     | European VAT Number                                                                                     |
	// | Czech Republic       | `eu_vat`     | European VAT Number                                                                                     |
	// | Denmark              | `eu_vat`     | European VAT Number                                                                                     |
	// | Dominican Republic   | `do_rcn`     | Dominican RCN Number                                                                                    |
	// | Ecuador              | `ec_ruc`     | Ecuadorian RUC Number                                                                                   |
	// | Egypt                | `eg_tin`     | Egyptian Tax Identification Number                                                                      |
	// | El Salvador          | `sv_nit`     | El Salvadorian NIT Number                                                                               |
	// | Estonia              | `eu_vat`     | European VAT Number                                                                                     |
	// | EU                   | `eu_oss_vat` | European One Stop Shop VAT Number for non-Union scheme                                                  |
	// | Finland              | `eu_vat`     | European VAT Number                                                                                     |
	// | France               | `eu_vat`     | European VAT Number                                                                                     |
	// | Georgia              | `ge_vat`     | Georgian VAT                                                                                            |
	// | Germany              | `eu_vat`     | European VAT Number                                                                                     |
	// | Greece               | `eu_vat`     | European VAT Number                                                                                     |
	// | Hong Kong            | `hk_br`      | Hong Kong BR Number                                                                                     |
	// | Hungary              | `eu_vat`     | European VAT Number                                                                                     |
	// | Hungary              | `hu_tin`     | Hungary Tax Number (adószám)                                                                            |
	// | Iceland              | `is_vat`     | Icelandic VAT                                                                                           |
	// | India                | `in_gst`     | Indian GST Number                                                                                       |
	// | Indonesia            | `id_npwp`    | Indonesian NPWP Number                                                                                  |
	// | Ireland              | `eu_vat`     | European VAT Number                                                                                     |
	// | Israel               | `il_vat`     | Israel VAT                                                                                              |
	// | Italy                | `eu_vat`     | European VAT Number                                                                                     |
	// | Japan                | `jp_cn`      | Japanese Corporate Number (_Hōjin Bangō_)                                                               |
	// | Japan                | `jp_rn`      | Japanese Registered Foreign Businesses' Registration Number (_Tōroku Kokugai Jigyōsha no Tōroku Bangō_) |
	// | Japan                | `jp_trn`     | Japanese Tax Registration Number (_Tōroku Bangō_)                                                       |
	// | Kazakhstan           | `kz_bin`     | Kazakhstani Business Identification Number                                                              |
	// | Kenya                | `ke_pin`     | Kenya Revenue Authority Personal Identification Number                                                  |
	// | Latvia               | `eu_vat`     | European VAT Number                                                                                     |
	// | Liechtenstein        | `li_uid`     | Liechtensteinian UID Number                                                                             |
	// | Lithuania            | `eu_vat`     | European VAT Number                                                                                     |
	// | Luxembourg           | `eu_vat`     | European VAT Number                                                                                     |
	// | Malaysia             | `my_frp`     | Malaysian FRP Number                                                                                    |
	// | Malaysia             | `my_itn`     | Malaysian ITN                                                                                           |
	// | Malaysia             | `my_sst`     | Malaysian SST Number                                                                                    |
	// | Malta                | `eu_vat `    | European VAT Number                                                                                     |
	// | Mexico               | `mx_rfc`     | Mexican RFC Number                                                                                      |
	// | Netherlands          | `eu_vat`     | European VAT Number                                                                                     |
	// | New Zealand          | `nz_gst`     | New Zealand GST Number                                                                                  |
	// | Nigeria              | `ng_tin`     | Nigerian Tax Identification Number                                                                      |
	// | Norway               | `no_vat`     | Norwegian VAT Number                                                                                    |
	// | Norway               | `no_voec`    | Norwegian VAT on e-commerce Number                                                                      |
	// | Oman                 | `om_vat`     | Omani VAT Number                                                                                        |
	// | Peru                 | `pe_ruc`     | Peruvian RUC Number                                                                                     |
	// | Philippines          | `ph_tin `    | Philippines Tax Identification Number                                                                   |
	// | Poland               | `eu_vat`     | European VAT Number                                                                                     |
	// | Portugal             | `eu_vat`     | European VAT Number                                                                                     |
	// | Romania              | `eu_vat`     | European VAT Number                                                                                     |
	// | Romania              | `ro_tin`     | Romanian Tax ID Number                                                                                  |
	// | Russia               | `ru_inn`     | Russian INN                                                                                             |
	// | Russia               | `ru_kpp`     | Russian KPP                                                                                             |
	// | Saudi Arabia         | `sa_vat`     | Saudi Arabia VAT                                                                                        |
	// | Serbia               | `rs_pib`     | Serbian PIB Number                                                                                      |
	// | Singapore            | `sg_gst`     | Singaporean GST                                                                                         |
	// | Singapore            | `sg_uen`     | Singaporean UEN                                                                                         |
	// | Slovakia             | `eu_vat`     | European VAT Number                                                                                     |
	// | Slovenia             | `eu_vat`     | European VAT Number                                                                                     |
	// | Slovenia             | `si_tin`     | Slovenia Tax Number (davčna številka)                                                                   |
	// | South Africa         | `za_vat`     | South African VAT Number                                                                                |
	// | South Korea          | `kr_brn`     | Korean BRN                                                                                              |
	// | Spain                | `es_cif`     | Spanish NIF Number (previously Spanish CIF Number)                                                      |
	// | Spain                | `eu_vat`     | European VAT Number                                                                                     |
	// | Sweden               | `eu_vat`     | European VAT Number                                                                                     |
	// | Switzerland          | `ch_vat`     | Switzerland VAT Number                                                                                  |
	// | Taiwan               | `tw_vat`     | Taiwanese VAT                                                                                           |
	// | Thailand             | `th_vat`     | Thai VAT                                                                                                |
	// | Turkey               | `tr_tin`     | Turkish Tax Identification Number                                                                       |
	// | Ukraine              | `ua_vat`     | Ukrainian VAT                                                                                           |
	// | United Arab Emirates | `ae_trn`     | United Arab Emirates TRN                                                                                |
	// | United Kingdom       | `eu_vat`     | Northern Ireland VAT Number                                                                             |
	// | United Kingdom       | `gb_vat`     | United Kingdom VAT Number                                                                               |
	// | United States        | `us_ein`     | United States EIN                                                                                       |
	// | Uruguay              | `uy_ruc`     | Uruguayan RUC Number                                                                                    |
	// | Venezuela            | `ve_rif`     | Venezuelan RIF Number                                                                                   |
	// | Vietnam              | `vn_tin`     | Vietnamese Tax ID Number                                                                                |
	TaxID CustomerTaxIDModel `json:"tax_id,required,nullable"`
	// A timezone identifier from the IANA timezone database, such as
	// "America/Los_Angeles". This "defaults to your account's timezone if not set.
	// This cannot be changed after customer creation.
	Timezone                    string                                   `json:"timezone,required"`
	AccountingSyncConfiguration CustomerModelAccountingSyncConfiguration `json:"accounting_sync_configuration,nullable"`
	ReportingConfiguration      CustomerModelReportingConfiguration      `json:"reporting_configuration,nullable"`
	JSON                        customerModelJSON                        `json:"-"`
}

// customerModelJSON contains the JSON metadata for the struct [CustomerModel]
type customerModelJSON struct {
	ID                          apijson.Field
	AdditionalEmails            apijson.Field
	AutoCollection              apijson.Field
	Balance                     apijson.Field
	BillingAddress              apijson.Field
	CreatedAt                   apijson.Field
	Currency                    apijson.Field
	Email                       apijson.Field
	EmailDelivery               apijson.Field
	ExemptFromAutomatedTax      apijson.Field
	ExternalCustomerID          apijson.Field
	Hierarchy                   apijson.Field
	Metadata                    apijson.Field
	Name                        apijson.Field
	PaymentProvider             apijson.Field
	PaymentProviderID           apijson.Field
	PortalURL                   apijson.Field
	ShippingAddress             apijson.Field
	TaxID                       apijson.Field
	Timezone                    apijson.Field
	AccountingSyncConfiguration apijson.Field
	ReportingConfiguration      apijson.Field
	raw                         string
	ExtraFields                 map[string]apijson.Field
}

func (r *CustomerModel) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r customerModelJSON) RawJSON() string {
	return r.raw
}

// The hierarchical relationships for this customer.
type CustomerModelHierarchy struct {
	Children []CustomerMinifiedModel    `json:"children,required"`
	Parent   CustomerMinifiedModel      `json:"parent,required,nullable"`
	JSON     customerModelHierarchyJSON `json:"-"`
}

// customerModelHierarchyJSON contains the JSON metadata for the struct
// [CustomerModelHierarchy]
type customerModelHierarchyJSON struct {
	Children    apijson.Field
	Parent      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *CustomerModelHierarchy) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r customerModelHierarchyJSON) RawJSON() string {
	return r.raw
}

// This is used for creating charges or invoices in an external system via Orb.
// When not in test mode, the connection must first be configured in the Orb
// webapp.
type CustomerModelPaymentProvider string

const (
	CustomerModelPaymentProviderQuickbooks    CustomerModelPaymentProvider = "quickbooks"
	CustomerModelPaymentProviderBillCom       CustomerModelPaymentProvider = "bill.com"
	CustomerModelPaymentProviderStripeCharge  CustomerModelPaymentProvider = "stripe_charge"
	CustomerModelPaymentProviderStripeInvoice CustomerModelPaymentProvider = "stripe_invoice"
	CustomerModelPaymentProviderNetsuite      CustomerModelPaymentProvider = "netsuite"
)

func (r CustomerModelPaymentProvider) IsKnown() bool {
	switch r {
	case CustomerModelPaymentProviderQuickbooks, CustomerModelPaymentProviderBillCom, CustomerModelPaymentProviderStripeCharge, CustomerModelPaymentProviderStripeInvoice, CustomerModelPaymentProviderNetsuite:
		return true
	}
	return false
}

type CustomerModelAccountingSyncConfiguration struct {
	AccountingProviders []CustomerModelAccountingSyncConfigurationAccountingProvider `json:"accounting_providers,required"`
	Excluded            bool                                                         `json:"excluded,required"`
	JSON                customerModelAccountingSyncConfigurationJSON                 `json:"-"`
}

// customerModelAccountingSyncConfigurationJSON contains the JSON metadata for the
// struct [CustomerModelAccountingSyncConfiguration]
type customerModelAccountingSyncConfigurationJSON struct {
	AccountingProviders apijson.Field
	Excluded            apijson.Field
	raw                 string
	ExtraFields         map[string]apijson.Field
}

func (r *CustomerModelAccountingSyncConfiguration) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r customerModelAccountingSyncConfigurationJSON) RawJSON() string {
	return r.raw
}

type CustomerModelAccountingSyncConfigurationAccountingProvider struct {
	ExternalProviderID string                                                                  `json:"external_provider_id,required,nullable"`
	ProviderType       CustomerModelAccountingSyncConfigurationAccountingProvidersProviderType `json:"provider_type,required"`
	JSON               customerModelAccountingSyncConfigurationAccountingProviderJSON          `json:"-"`
}

// customerModelAccountingSyncConfigurationAccountingProviderJSON contains the JSON
// metadata for the struct
// [CustomerModelAccountingSyncConfigurationAccountingProvider]
type customerModelAccountingSyncConfigurationAccountingProviderJSON struct {
	ExternalProviderID apijson.Field
	ProviderType       apijson.Field
	raw                string
	ExtraFields        map[string]apijson.Field
}

func (r *CustomerModelAccountingSyncConfigurationAccountingProvider) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r customerModelAccountingSyncConfigurationAccountingProviderJSON) RawJSON() string {
	return r.raw
}

type CustomerModelAccountingSyncConfigurationAccountingProvidersProviderType string

const (
	CustomerModelAccountingSyncConfigurationAccountingProvidersProviderTypeQuickbooks CustomerModelAccountingSyncConfigurationAccountingProvidersProviderType = "quickbooks"
	CustomerModelAccountingSyncConfigurationAccountingProvidersProviderTypeNetsuite   CustomerModelAccountingSyncConfigurationAccountingProvidersProviderType = "netsuite"
)

func (r CustomerModelAccountingSyncConfigurationAccountingProvidersProviderType) IsKnown() bool {
	switch r {
	case CustomerModelAccountingSyncConfigurationAccountingProvidersProviderTypeQuickbooks, CustomerModelAccountingSyncConfigurationAccountingProvidersProviderTypeNetsuite:
		return true
	}
	return false
}

type CustomerModelReportingConfiguration struct {
	Exempt bool                                    `json:"exempt,required"`
	JSON   customerModelReportingConfigurationJSON `json:"-"`
}

// customerModelReportingConfigurationJSON contains the JSON metadata for the
// struct [CustomerModelReportingConfiguration]
type customerModelReportingConfigurationJSON struct {
	Exempt      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *CustomerModelReportingConfiguration) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r customerModelReportingConfigurationJSON) RawJSON() string {
	return r.raw
}

// Tax IDs are commonly required to be displayed on customer invoices, which are
// added to the headers of invoices.
//
// ### Supported Tax ID Countries and Types
//
// | Country              | Type         | Description                                                                                             |
// | -------------------- | ------------ | ------------------------------------------------------------------------------------------------------- |
// | Andorra              | `ad_nrt`     | Andorran NRT Number                                                                                     |
// | Argentina            | `ar_cuit`    | Argentinian Tax ID Number                                                                               |
// | Australia            | `au_abn`     | Australian Business Number (AU ABN)                                                                     |
// | Australia            | `au_arn`     | Australian Taxation Office Reference Number                                                             |
// | Austria              | `eu_vat`     | European VAT Number                                                                                     |
// | Bahrain              | `bh_vat`     | Bahraini VAT Number                                                                                     |
// | Belgium              | `eu_vat`     | European VAT Number                                                                                     |
// | Bolivia              | `bo_tin`     | Bolivian Tax ID                                                                                         |
// | Brazil               | `br_cnpj`    | Brazilian CNPJ Number                                                                                   |
// | Brazil               | `br_cpf`     | Brazilian CPF Number                                                                                    |
// | Bulgaria             | `bg_uic`     | Bulgaria Unified Identification Code                                                                    |
// | Bulgaria             | `eu_vat`     | European VAT Number                                                                                     |
// | Canada               | `ca_bn`      | Canadian BN                                                                                             |
// | Canada               | `ca_gst_hst` | Canadian GST/HST Number                                                                                 |
// | Canada               | `ca_pst_bc`  | Canadian PST Number (British Columbia)                                                                  |
// | Canada               | `ca_pst_mb`  | Canadian PST Number (Manitoba)                                                                          |
// | Canada               | `ca_pst_sk`  | Canadian PST Number (Saskatchewan)                                                                      |
// | Canada               | `ca_qst`     | Canadian QST Number (Québec)                                                                            |
// | Chile                | `cl_tin`     | Chilean TIN                                                                                             |
// | China                | `cn_tin`     | Chinese Tax ID                                                                                          |
// | Colombia             | `co_nit`     | Colombian NIT Number                                                                                    |
// | Costa Rica           | `cr_tin`     | Costa Rican Tax ID                                                                                      |
// | Croatia              | `eu_vat`     | European VAT Number                                                                                     |
// | Cyprus               | `eu_vat`     | European VAT Number                                                                                     |
// | Czech Republic       | `eu_vat`     | European VAT Number                                                                                     |
// | Denmark              | `eu_vat`     | European VAT Number                                                                                     |
// | Dominican Republic   | `do_rcn`     | Dominican RCN Number                                                                                    |
// | Ecuador              | `ec_ruc`     | Ecuadorian RUC Number                                                                                   |
// | Egypt                | `eg_tin`     | Egyptian Tax Identification Number                                                                      |
// | El Salvador          | `sv_nit`     | El Salvadorian NIT Number                                                                               |
// | Estonia              | `eu_vat`     | European VAT Number                                                                                     |
// | EU                   | `eu_oss_vat` | European One Stop Shop VAT Number for non-Union scheme                                                  |
// | Finland              | `eu_vat`     | European VAT Number                                                                                     |
// | France               | `eu_vat`     | European VAT Number                                                                                     |
// | Georgia              | `ge_vat`     | Georgian VAT                                                                                            |
// | Germany              | `eu_vat`     | European VAT Number                                                                                     |
// | Greece               | `eu_vat`     | European VAT Number                                                                                     |
// | Hong Kong            | `hk_br`      | Hong Kong BR Number                                                                                     |
// | Hungary              | `eu_vat`     | European VAT Number                                                                                     |
// | Hungary              | `hu_tin`     | Hungary Tax Number (adószám)                                                                            |
// | Iceland              | `is_vat`     | Icelandic VAT                                                                                           |
// | India                | `in_gst`     | Indian GST Number                                                                                       |
// | Indonesia            | `id_npwp`    | Indonesian NPWP Number                                                                                  |
// | Ireland              | `eu_vat`     | European VAT Number                                                                                     |
// | Israel               | `il_vat`     | Israel VAT                                                                                              |
// | Italy                | `eu_vat`     | European VAT Number                                                                                     |
// | Japan                | `jp_cn`      | Japanese Corporate Number (_Hōjin Bangō_)                                                               |
// | Japan                | `jp_rn`      | Japanese Registered Foreign Businesses' Registration Number (_Tōroku Kokugai Jigyōsha no Tōroku Bangō_) |
// | Japan                | `jp_trn`     | Japanese Tax Registration Number (_Tōroku Bangō_)                                                       |
// | Kazakhstan           | `kz_bin`     | Kazakhstani Business Identification Number                                                              |
// | Kenya                | `ke_pin`     | Kenya Revenue Authority Personal Identification Number                                                  |
// | Latvia               | `eu_vat`     | European VAT Number                                                                                     |
// | Liechtenstein        | `li_uid`     | Liechtensteinian UID Number                                                                             |
// | Lithuania            | `eu_vat`     | European VAT Number                                                                                     |
// | Luxembourg           | `eu_vat`     | European VAT Number                                                                                     |
// | Malaysia             | `my_frp`     | Malaysian FRP Number                                                                                    |
// | Malaysia             | `my_itn`     | Malaysian ITN                                                                                           |
// | Malaysia             | `my_sst`     | Malaysian SST Number                                                                                    |
// | Malta                | `eu_vat `    | European VAT Number                                                                                     |
// | Mexico               | `mx_rfc`     | Mexican RFC Number                                                                                      |
// | Netherlands          | `eu_vat`     | European VAT Number                                                                                     |
// | New Zealand          | `nz_gst`     | New Zealand GST Number                                                                                  |
// | Nigeria              | `ng_tin`     | Nigerian Tax Identification Number                                                                      |
// | Norway               | `no_vat`     | Norwegian VAT Number                                                                                    |
// | Norway               | `no_voec`    | Norwegian VAT on e-commerce Number                                                                      |
// | Oman                 | `om_vat`     | Omani VAT Number                                                                                        |
// | Peru                 | `pe_ruc`     | Peruvian RUC Number                                                                                     |
// | Philippines          | `ph_tin `    | Philippines Tax Identification Number                                                                   |
// | Poland               | `eu_vat`     | European VAT Number                                                                                     |
// | Portugal             | `eu_vat`     | European VAT Number                                                                                     |
// | Romania              | `eu_vat`     | European VAT Number                                                                                     |
// | Romania              | `ro_tin`     | Romanian Tax ID Number                                                                                  |
// | Russia               | `ru_inn`     | Russian INN                                                                                             |
// | Russia               | `ru_kpp`     | Russian KPP                                                                                             |
// | Saudi Arabia         | `sa_vat`     | Saudi Arabia VAT                                                                                        |
// | Serbia               | `rs_pib`     | Serbian PIB Number                                                                                      |
// | Singapore            | `sg_gst`     | Singaporean GST                                                                                         |
// | Singapore            | `sg_uen`     | Singaporean UEN                                                                                         |
// | Slovakia             | `eu_vat`     | European VAT Number                                                                                     |
// | Slovenia             | `eu_vat`     | European VAT Number                                                                                     |
// | Slovenia             | `si_tin`     | Slovenia Tax Number (davčna številka)                                                                   |
// | South Africa         | `za_vat`     | South African VAT Number                                                                                |
// | South Korea          | `kr_brn`     | Korean BRN                                                                                              |
// | Spain                | `es_cif`     | Spanish NIF Number (previously Spanish CIF Number)                                                      |
// | Spain                | `eu_vat`     | European VAT Number                                                                                     |
// | Sweden               | `eu_vat`     | European VAT Number                                                                                     |
// | Switzerland          | `ch_vat`     | Switzerland VAT Number                                                                                  |
// | Taiwan               | `tw_vat`     | Taiwanese VAT                                                                                           |
// | Thailand             | `th_vat`     | Thai VAT                                                                                                |
// | Turkey               | `tr_tin`     | Turkish Tax Identification Number                                                                       |
// | Ukraine              | `ua_vat`     | Ukrainian VAT                                                                                           |
// | United Arab Emirates | `ae_trn`     | United Arab Emirates TRN                                                                                |
// | United Kingdom       | `eu_vat`     | Northern Ireland VAT Number                                                                             |
// | United Kingdom       | `gb_vat`     | United Kingdom VAT Number                                                                               |
// | United States        | `us_ein`     | United States EIN                                                                                       |
// | Uruguay              | `uy_ruc`     | Uruguayan RUC Number                                                                                    |
// | Venezuela            | `ve_rif`     | Venezuelan RIF Number                                                                                   |
// | Vietnam              | `vn_tin`     | Vietnamese Tax ID Number                                                                                |
type CustomerTaxIDModel struct {
	Country CustomerTaxIDModelCountry `json:"country,required"`
	Type    CustomerTaxIDModelType    `json:"type,required"`
	Value   string                    `json:"value,required"`
	JSON    customerTaxIDModelJSON    `json:"-"`
}

// customerTaxIDModelJSON contains the JSON metadata for the struct
// [CustomerTaxIDModel]
type customerTaxIDModelJSON struct {
	Country     apijson.Field
	Type        apijson.Field
	Value       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *CustomerTaxIDModel) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r customerTaxIDModelJSON) RawJSON() string {
	return r.raw
}

type CustomerTaxIDModelCountry string

const (
	CustomerTaxIDModelCountryAd CustomerTaxIDModelCountry = "AD"
	CustomerTaxIDModelCountryAe CustomerTaxIDModelCountry = "AE"
	CustomerTaxIDModelCountryAr CustomerTaxIDModelCountry = "AR"
	CustomerTaxIDModelCountryAt CustomerTaxIDModelCountry = "AT"
	CustomerTaxIDModelCountryAu CustomerTaxIDModelCountry = "AU"
	CustomerTaxIDModelCountryBe CustomerTaxIDModelCountry = "BE"
	CustomerTaxIDModelCountryBg CustomerTaxIDModelCountry = "BG"
	CustomerTaxIDModelCountryBh CustomerTaxIDModelCountry = "BH"
	CustomerTaxIDModelCountryBo CustomerTaxIDModelCountry = "BO"
	CustomerTaxIDModelCountryBr CustomerTaxIDModelCountry = "BR"
	CustomerTaxIDModelCountryCa CustomerTaxIDModelCountry = "CA"
	CustomerTaxIDModelCountryCh CustomerTaxIDModelCountry = "CH"
	CustomerTaxIDModelCountryCl CustomerTaxIDModelCountry = "CL"
	CustomerTaxIDModelCountryCn CustomerTaxIDModelCountry = "CN"
	CustomerTaxIDModelCountryCo CustomerTaxIDModelCountry = "CO"
	CustomerTaxIDModelCountryCr CustomerTaxIDModelCountry = "CR"
	CustomerTaxIDModelCountryCy CustomerTaxIDModelCountry = "CY"
	CustomerTaxIDModelCountryCz CustomerTaxIDModelCountry = "CZ"
	CustomerTaxIDModelCountryDe CustomerTaxIDModelCountry = "DE"
	CustomerTaxIDModelCountryDk CustomerTaxIDModelCountry = "DK"
	CustomerTaxIDModelCountryEe CustomerTaxIDModelCountry = "EE"
	CustomerTaxIDModelCountryDo CustomerTaxIDModelCountry = "DO"
	CustomerTaxIDModelCountryEc CustomerTaxIDModelCountry = "EC"
	CustomerTaxIDModelCountryEg CustomerTaxIDModelCountry = "EG"
	CustomerTaxIDModelCountryEs CustomerTaxIDModelCountry = "ES"
	CustomerTaxIDModelCountryEu CustomerTaxIDModelCountry = "EU"
	CustomerTaxIDModelCountryFi CustomerTaxIDModelCountry = "FI"
	CustomerTaxIDModelCountryFr CustomerTaxIDModelCountry = "FR"
	CustomerTaxIDModelCountryGB CustomerTaxIDModelCountry = "GB"
	CustomerTaxIDModelCountryGe CustomerTaxIDModelCountry = "GE"
	CustomerTaxIDModelCountryGr CustomerTaxIDModelCountry = "GR"
	CustomerTaxIDModelCountryHk CustomerTaxIDModelCountry = "HK"
	CustomerTaxIDModelCountryHr CustomerTaxIDModelCountry = "HR"
	CustomerTaxIDModelCountryHu CustomerTaxIDModelCountry = "HU"
	CustomerTaxIDModelCountryID CustomerTaxIDModelCountry = "ID"
	CustomerTaxIDModelCountryIe CustomerTaxIDModelCountry = "IE"
	CustomerTaxIDModelCountryIl CustomerTaxIDModelCountry = "IL"
	CustomerTaxIDModelCountryIn CustomerTaxIDModelCountry = "IN"
	CustomerTaxIDModelCountryIs CustomerTaxIDModelCountry = "IS"
	CustomerTaxIDModelCountryIt CustomerTaxIDModelCountry = "IT"
	CustomerTaxIDModelCountryJp CustomerTaxIDModelCountry = "JP"
	CustomerTaxIDModelCountryKe CustomerTaxIDModelCountry = "KE"
	CustomerTaxIDModelCountryKr CustomerTaxIDModelCountry = "KR"
	CustomerTaxIDModelCountryKz CustomerTaxIDModelCountry = "KZ"
	CustomerTaxIDModelCountryLi CustomerTaxIDModelCountry = "LI"
	CustomerTaxIDModelCountryLt CustomerTaxIDModelCountry = "LT"
	CustomerTaxIDModelCountryLu CustomerTaxIDModelCountry = "LU"
	CustomerTaxIDModelCountryLv CustomerTaxIDModelCountry = "LV"
	CustomerTaxIDModelCountryMt CustomerTaxIDModelCountry = "MT"
	CustomerTaxIDModelCountryMx CustomerTaxIDModelCountry = "MX"
	CustomerTaxIDModelCountryMy CustomerTaxIDModelCountry = "MY"
	CustomerTaxIDModelCountryNg CustomerTaxIDModelCountry = "NG"
	CustomerTaxIDModelCountryNl CustomerTaxIDModelCountry = "NL"
	CustomerTaxIDModelCountryNo CustomerTaxIDModelCountry = "NO"
	CustomerTaxIDModelCountryNz CustomerTaxIDModelCountry = "NZ"
	CustomerTaxIDModelCountryOm CustomerTaxIDModelCountry = "OM"
	CustomerTaxIDModelCountryPe CustomerTaxIDModelCountry = "PE"
	CustomerTaxIDModelCountryPh CustomerTaxIDModelCountry = "PH"
	CustomerTaxIDModelCountryPl CustomerTaxIDModelCountry = "PL"
	CustomerTaxIDModelCountryPt CustomerTaxIDModelCountry = "PT"
	CustomerTaxIDModelCountryRo CustomerTaxIDModelCountry = "RO"
	CustomerTaxIDModelCountryRs CustomerTaxIDModelCountry = "RS"
	CustomerTaxIDModelCountryRu CustomerTaxIDModelCountry = "RU"
	CustomerTaxIDModelCountrySa CustomerTaxIDModelCountry = "SA"
	CustomerTaxIDModelCountrySe CustomerTaxIDModelCountry = "SE"
	CustomerTaxIDModelCountrySg CustomerTaxIDModelCountry = "SG"
	CustomerTaxIDModelCountrySi CustomerTaxIDModelCountry = "SI"
	CustomerTaxIDModelCountrySk CustomerTaxIDModelCountry = "SK"
	CustomerTaxIDModelCountrySv CustomerTaxIDModelCountry = "SV"
	CustomerTaxIDModelCountryTh CustomerTaxIDModelCountry = "TH"
	CustomerTaxIDModelCountryTr CustomerTaxIDModelCountry = "TR"
	CustomerTaxIDModelCountryTw CustomerTaxIDModelCountry = "TW"
	CustomerTaxIDModelCountryUa CustomerTaxIDModelCountry = "UA"
	CustomerTaxIDModelCountryUs CustomerTaxIDModelCountry = "US"
	CustomerTaxIDModelCountryUy CustomerTaxIDModelCountry = "UY"
	CustomerTaxIDModelCountryVe CustomerTaxIDModelCountry = "VE"
	CustomerTaxIDModelCountryVn CustomerTaxIDModelCountry = "VN"
	CustomerTaxIDModelCountryZa CustomerTaxIDModelCountry = "ZA"
)

func (r CustomerTaxIDModelCountry) IsKnown() bool {
	switch r {
	case CustomerTaxIDModelCountryAd, CustomerTaxIDModelCountryAe, CustomerTaxIDModelCountryAr, CustomerTaxIDModelCountryAt, CustomerTaxIDModelCountryAu, CustomerTaxIDModelCountryBe, CustomerTaxIDModelCountryBg, CustomerTaxIDModelCountryBh, CustomerTaxIDModelCountryBo, CustomerTaxIDModelCountryBr, CustomerTaxIDModelCountryCa, CustomerTaxIDModelCountryCh, CustomerTaxIDModelCountryCl, CustomerTaxIDModelCountryCn, CustomerTaxIDModelCountryCo, CustomerTaxIDModelCountryCr, CustomerTaxIDModelCountryCy, CustomerTaxIDModelCountryCz, CustomerTaxIDModelCountryDe, CustomerTaxIDModelCountryDk, CustomerTaxIDModelCountryEe, CustomerTaxIDModelCountryDo, CustomerTaxIDModelCountryEc, CustomerTaxIDModelCountryEg, CustomerTaxIDModelCountryEs, CustomerTaxIDModelCountryEu, CustomerTaxIDModelCountryFi, CustomerTaxIDModelCountryFr, CustomerTaxIDModelCountryGB, CustomerTaxIDModelCountryGe, CustomerTaxIDModelCountryGr, CustomerTaxIDModelCountryHk, CustomerTaxIDModelCountryHr, CustomerTaxIDModelCountryHu, CustomerTaxIDModelCountryID, CustomerTaxIDModelCountryIe, CustomerTaxIDModelCountryIl, CustomerTaxIDModelCountryIn, CustomerTaxIDModelCountryIs, CustomerTaxIDModelCountryIt, CustomerTaxIDModelCountryJp, CustomerTaxIDModelCountryKe, CustomerTaxIDModelCountryKr, CustomerTaxIDModelCountryKz, CustomerTaxIDModelCountryLi, CustomerTaxIDModelCountryLt, CustomerTaxIDModelCountryLu, CustomerTaxIDModelCountryLv, CustomerTaxIDModelCountryMt, CustomerTaxIDModelCountryMx, CustomerTaxIDModelCountryMy, CustomerTaxIDModelCountryNg, CustomerTaxIDModelCountryNl, CustomerTaxIDModelCountryNo, CustomerTaxIDModelCountryNz, CustomerTaxIDModelCountryOm, CustomerTaxIDModelCountryPe, CustomerTaxIDModelCountryPh, CustomerTaxIDModelCountryPl, CustomerTaxIDModelCountryPt, CustomerTaxIDModelCountryRo, CustomerTaxIDModelCountryRs, CustomerTaxIDModelCountryRu, CustomerTaxIDModelCountrySa, CustomerTaxIDModelCountrySe, CustomerTaxIDModelCountrySg, CustomerTaxIDModelCountrySi, CustomerTaxIDModelCountrySk, CustomerTaxIDModelCountrySv, CustomerTaxIDModelCountryTh, CustomerTaxIDModelCountryTr, CustomerTaxIDModelCountryTw, CustomerTaxIDModelCountryUa, CustomerTaxIDModelCountryUs, CustomerTaxIDModelCountryUy, CustomerTaxIDModelCountryVe, CustomerTaxIDModelCountryVn, CustomerTaxIDModelCountryZa:
		return true
	}
	return false
}

type CustomerTaxIDModelType string

const (
	CustomerTaxIDModelTypeAdNrt    CustomerTaxIDModelType = "ad_nrt"
	CustomerTaxIDModelTypeAeTrn    CustomerTaxIDModelType = "ae_trn"
	CustomerTaxIDModelTypeArCuit   CustomerTaxIDModelType = "ar_cuit"
	CustomerTaxIDModelTypeEuVat    CustomerTaxIDModelType = "eu_vat"
	CustomerTaxIDModelTypeAuAbn    CustomerTaxIDModelType = "au_abn"
	CustomerTaxIDModelTypeAuArn    CustomerTaxIDModelType = "au_arn"
	CustomerTaxIDModelTypeBgUic    CustomerTaxIDModelType = "bg_uic"
	CustomerTaxIDModelTypeBhVat    CustomerTaxIDModelType = "bh_vat"
	CustomerTaxIDModelTypeBoTin    CustomerTaxIDModelType = "bo_tin"
	CustomerTaxIDModelTypeBrCnpj   CustomerTaxIDModelType = "br_cnpj"
	CustomerTaxIDModelTypeBrCpf    CustomerTaxIDModelType = "br_cpf"
	CustomerTaxIDModelTypeCaBn     CustomerTaxIDModelType = "ca_bn"
	CustomerTaxIDModelTypeCaGstHst CustomerTaxIDModelType = "ca_gst_hst"
	CustomerTaxIDModelTypeCaPstBc  CustomerTaxIDModelType = "ca_pst_bc"
	CustomerTaxIDModelTypeCaPstMB  CustomerTaxIDModelType = "ca_pst_mb"
	CustomerTaxIDModelTypeCaPstSk  CustomerTaxIDModelType = "ca_pst_sk"
	CustomerTaxIDModelTypeCaQst    CustomerTaxIDModelType = "ca_qst"
	CustomerTaxIDModelTypeChVat    CustomerTaxIDModelType = "ch_vat"
	CustomerTaxIDModelTypeClTin    CustomerTaxIDModelType = "cl_tin"
	CustomerTaxIDModelTypeCnTin    CustomerTaxIDModelType = "cn_tin"
	CustomerTaxIDModelTypeCoNit    CustomerTaxIDModelType = "co_nit"
	CustomerTaxIDModelTypeCrTin    CustomerTaxIDModelType = "cr_tin"
	CustomerTaxIDModelTypeDoRcn    CustomerTaxIDModelType = "do_rcn"
	CustomerTaxIDModelTypeEcRuc    CustomerTaxIDModelType = "ec_ruc"
	CustomerTaxIDModelTypeEgTin    CustomerTaxIDModelType = "eg_tin"
	CustomerTaxIDModelTypeEsCif    CustomerTaxIDModelType = "es_cif"
	CustomerTaxIDModelTypeEuOssVat CustomerTaxIDModelType = "eu_oss_vat"
	CustomerTaxIDModelTypeGBVat    CustomerTaxIDModelType = "gb_vat"
	CustomerTaxIDModelTypeGeVat    CustomerTaxIDModelType = "ge_vat"
	CustomerTaxIDModelTypeHkBr     CustomerTaxIDModelType = "hk_br"
	CustomerTaxIDModelTypeHuTin    CustomerTaxIDModelType = "hu_tin"
	CustomerTaxIDModelTypeIDNpwp   CustomerTaxIDModelType = "id_npwp"
	CustomerTaxIDModelTypeIlVat    CustomerTaxIDModelType = "il_vat"
	CustomerTaxIDModelTypeInGst    CustomerTaxIDModelType = "in_gst"
	CustomerTaxIDModelTypeIsVat    CustomerTaxIDModelType = "is_vat"
	CustomerTaxIDModelTypeJpCn     CustomerTaxIDModelType = "jp_cn"
	CustomerTaxIDModelTypeJpRn     CustomerTaxIDModelType = "jp_rn"
	CustomerTaxIDModelTypeJpTrn    CustomerTaxIDModelType = "jp_trn"
	CustomerTaxIDModelTypeKePin    CustomerTaxIDModelType = "ke_pin"
	CustomerTaxIDModelTypeKrBrn    CustomerTaxIDModelType = "kr_brn"
	CustomerTaxIDModelTypeKzBin    CustomerTaxIDModelType = "kz_bin"
	CustomerTaxIDModelTypeLiUid    CustomerTaxIDModelType = "li_uid"
	CustomerTaxIDModelTypeMxRfc    CustomerTaxIDModelType = "mx_rfc"
	CustomerTaxIDModelTypeMyFrp    CustomerTaxIDModelType = "my_frp"
	CustomerTaxIDModelTypeMyItn    CustomerTaxIDModelType = "my_itn"
	CustomerTaxIDModelTypeMySst    CustomerTaxIDModelType = "my_sst"
	CustomerTaxIDModelTypeNgTin    CustomerTaxIDModelType = "ng_tin"
	CustomerTaxIDModelTypeNoVat    CustomerTaxIDModelType = "no_vat"
	CustomerTaxIDModelTypeNoVoec   CustomerTaxIDModelType = "no_voec"
	CustomerTaxIDModelTypeNzGst    CustomerTaxIDModelType = "nz_gst"
	CustomerTaxIDModelTypeOmVat    CustomerTaxIDModelType = "om_vat"
	CustomerTaxIDModelTypePeRuc    CustomerTaxIDModelType = "pe_ruc"
	CustomerTaxIDModelTypePhTin    CustomerTaxIDModelType = "ph_tin"
	CustomerTaxIDModelTypeRoTin    CustomerTaxIDModelType = "ro_tin"
	CustomerTaxIDModelTypeRsPib    CustomerTaxIDModelType = "rs_pib"
	CustomerTaxIDModelTypeRuInn    CustomerTaxIDModelType = "ru_inn"
	CustomerTaxIDModelTypeRuKpp    CustomerTaxIDModelType = "ru_kpp"
	CustomerTaxIDModelTypeSaVat    CustomerTaxIDModelType = "sa_vat"
	CustomerTaxIDModelTypeSgGst    CustomerTaxIDModelType = "sg_gst"
	CustomerTaxIDModelTypeSgUen    CustomerTaxIDModelType = "sg_uen"
	CustomerTaxIDModelTypeSiTin    CustomerTaxIDModelType = "si_tin"
	CustomerTaxIDModelTypeSvNit    CustomerTaxIDModelType = "sv_nit"
	CustomerTaxIDModelTypeThVat    CustomerTaxIDModelType = "th_vat"
	CustomerTaxIDModelTypeTrTin    CustomerTaxIDModelType = "tr_tin"
	CustomerTaxIDModelTypeTwVat    CustomerTaxIDModelType = "tw_vat"
	CustomerTaxIDModelTypeUaVat    CustomerTaxIDModelType = "ua_vat"
	CustomerTaxIDModelTypeUsEin    CustomerTaxIDModelType = "us_ein"
	CustomerTaxIDModelTypeUyRuc    CustomerTaxIDModelType = "uy_ruc"
	CustomerTaxIDModelTypeVeRif    CustomerTaxIDModelType = "ve_rif"
	CustomerTaxIDModelTypeVnTin    CustomerTaxIDModelType = "vn_tin"
	CustomerTaxIDModelTypeZaVat    CustomerTaxIDModelType = "za_vat"
)

func (r CustomerTaxIDModelType) IsKnown() bool {
	switch r {
	case CustomerTaxIDModelTypeAdNrt, CustomerTaxIDModelTypeAeTrn, CustomerTaxIDModelTypeArCuit, CustomerTaxIDModelTypeEuVat, CustomerTaxIDModelTypeAuAbn, CustomerTaxIDModelTypeAuArn, CustomerTaxIDModelTypeBgUic, CustomerTaxIDModelTypeBhVat, CustomerTaxIDModelTypeBoTin, CustomerTaxIDModelTypeBrCnpj, CustomerTaxIDModelTypeBrCpf, CustomerTaxIDModelTypeCaBn, CustomerTaxIDModelTypeCaGstHst, CustomerTaxIDModelTypeCaPstBc, CustomerTaxIDModelTypeCaPstMB, CustomerTaxIDModelTypeCaPstSk, CustomerTaxIDModelTypeCaQst, CustomerTaxIDModelTypeChVat, CustomerTaxIDModelTypeClTin, CustomerTaxIDModelTypeCnTin, CustomerTaxIDModelTypeCoNit, CustomerTaxIDModelTypeCrTin, CustomerTaxIDModelTypeDoRcn, CustomerTaxIDModelTypeEcRuc, CustomerTaxIDModelTypeEgTin, CustomerTaxIDModelTypeEsCif, CustomerTaxIDModelTypeEuOssVat, CustomerTaxIDModelTypeGBVat, CustomerTaxIDModelTypeGeVat, CustomerTaxIDModelTypeHkBr, CustomerTaxIDModelTypeHuTin, CustomerTaxIDModelTypeIDNpwp, CustomerTaxIDModelTypeIlVat, CustomerTaxIDModelTypeInGst, CustomerTaxIDModelTypeIsVat, CustomerTaxIDModelTypeJpCn, CustomerTaxIDModelTypeJpRn, CustomerTaxIDModelTypeJpTrn, CustomerTaxIDModelTypeKePin, CustomerTaxIDModelTypeKrBrn, CustomerTaxIDModelTypeKzBin, CustomerTaxIDModelTypeLiUid, CustomerTaxIDModelTypeMxRfc, CustomerTaxIDModelTypeMyFrp, CustomerTaxIDModelTypeMyItn, CustomerTaxIDModelTypeMySst, CustomerTaxIDModelTypeNgTin, CustomerTaxIDModelTypeNoVat, CustomerTaxIDModelTypeNoVoec, CustomerTaxIDModelTypeNzGst, CustomerTaxIDModelTypeOmVat, CustomerTaxIDModelTypePeRuc, CustomerTaxIDModelTypePhTin, CustomerTaxIDModelTypeRoTin, CustomerTaxIDModelTypeRsPib, CustomerTaxIDModelTypeRuInn, CustomerTaxIDModelTypeRuKpp, CustomerTaxIDModelTypeSaVat, CustomerTaxIDModelTypeSgGst, CustomerTaxIDModelTypeSgUen, CustomerTaxIDModelTypeSiTin, CustomerTaxIDModelTypeSvNit, CustomerTaxIDModelTypeThVat, CustomerTaxIDModelTypeTrTin, CustomerTaxIDModelTypeTwVat, CustomerTaxIDModelTypeUaVat, CustomerTaxIDModelTypeUsEin, CustomerTaxIDModelTypeUyRuc, CustomerTaxIDModelTypeVeRif, CustomerTaxIDModelTypeVnTin, CustomerTaxIDModelTypeZaVat:
		return true
	}
	return false
}

// Tax IDs are commonly required to be displayed on customer invoices, which are
// added to the headers of invoices.
//
// ### Supported Tax ID Countries and Types
//
// | Country              | Type         | Description                                                                                             |
// | -------------------- | ------------ | ------------------------------------------------------------------------------------------------------- |
// | Andorra              | `ad_nrt`     | Andorran NRT Number                                                                                     |
// | Argentina            | `ar_cuit`    | Argentinian Tax ID Number                                                                               |
// | Australia            | `au_abn`     | Australian Business Number (AU ABN)                                                                     |
// | Australia            | `au_arn`     | Australian Taxation Office Reference Number                                                             |
// | Austria              | `eu_vat`     | European VAT Number                                                                                     |
// | Bahrain              | `bh_vat`     | Bahraini VAT Number                                                                                     |
// | Belgium              | `eu_vat`     | European VAT Number                                                                                     |
// | Bolivia              | `bo_tin`     | Bolivian Tax ID                                                                                         |
// | Brazil               | `br_cnpj`    | Brazilian CNPJ Number                                                                                   |
// | Brazil               | `br_cpf`     | Brazilian CPF Number                                                                                    |
// | Bulgaria             | `bg_uic`     | Bulgaria Unified Identification Code                                                                    |
// | Bulgaria             | `eu_vat`     | European VAT Number                                                                                     |
// | Canada               | `ca_bn`      | Canadian BN                                                                                             |
// | Canada               | `ca_gst_hst` | Canadian GST/HST Number                                                                                 |
// | Canada               | `ca_pst_bc`  | Canadian PST Number (British Columbia)                                                                  |
// | Canada               | `ca_pst_mb`  | Canadian PST Number (Manitoba)                                                                          |
// | Canada               | `ca_pst_sk`  | Canadian PST Number (Saskatchewan)                                                                      |
// | Canada               | `ca_qst`     | Canadian QST Number (Québec)                                                                            |
// | Chile                | `cl_tin`     | Chilean TIN                                                                                             |
// | China                | `cn_tin`     | Chinese Tax ID                                                                                          |
// | Colombia             | `co_nit`     | Colombian NIT Number                                                                                    |
// | Costa Rica           | `cr_tin`     | Costa Rican Tax ID                                                                                      |
// | Croatia              | `eu_vat`     | European VAT Number                                                                                     |
// | Cyprus               | `eu_vat`     | European VAT Number                                                                                     |
// | Czech Republic       | `eu_vat`     | European VAT Number                                                                                     |
// | Denmark              | `eu_vat`     | European VAT Number                                                                                     |
// | Dominican Republic   | `do_rcn`     | Dominican RCN Number                                                                                    |
// | Ecuador              | `ec_ruc`     | Ecuadorian RUC Number                                                                                   |
// | Egypt                | `eg_tin`     | Egyptian Tax Identification Number                                                                      |
// | El Salvador          | `sv_nit`     | El Salvadorian NIT Number                                                                               |
// | Estonia              | `eu_vat`     | European VAT Number                                                                                     |
// | EU                   | `eu_oss_vat` | European One Stop Shop VAT Number for non-Union scheme                                                  |
// | Finland              | `eu_vat`     | European VAT Number                                                                                     |
// | France               | `eu_vat`     | European VAT Number                                                                                     |
// | Georgia              | `ge_vat`     | Georgian VAT                                                                                            |
// | Germany              | `eu_vat`     | European VAT Number                                                                                     |
// | Greece               | `eu_vat`     | European VAT Number                                                                                     |
// | Hong Kong            | `hk_br`      | Hong Kong BR Number                                                                                     |
// | Hungary              | `eu_vat`     | European VAT Number                                                                                     |
// | Hungary              | `hu_tin`     | Hungary Tax Number (adószám)                                                                            |
// | Iceland              | `is_vat`     | Icelandic VAT                                                                                           |
// | India                | `in_gst`     | Indian GST Number                                                                                       |
// | Indonesia            | `id_npwp`    | Indonesian NPWP Number                                                                                  |
// | Ireland              | `eu_vat`     | European VAT Number                                                                                     |
// | Israel               | `il_vat`     | Israel VAT                                                                                              |
// | Italy                | `eu_vat`     | European VAT Number                                                                                     |
// | Japan                | `jp_cn`      | Japanese Corporate Number (_Hōjin Bangō_)                                                               |
// | Japan                | `jp_rn`      | Japanese Registered Foreign Businesses' Registration Number (_Tōroku Kokugai Jigyōsha no Tōroku Bangō_) |
// | Japan                | `jp_trn`     | Japanese Tax Registration Number (_Tōroku Bangō_)                                                       |
// | Kazakhstan           | `kz_bin`     | Kazakhstani Business Identification Number                                                              |
// | Kenya                | `ke_pin`     | Kenya Revenue Authority Personal Identification Number                                                  |
// | Latvia               | `eu_vat`     | European VAT Number                                                                                     |
// | Liechtenstein        | `li_uid`     | Liechtensteinian UID Number                                                                             |
// | Lithuania            | `eu_vat`     | European VAT Number                                                                                     |
// | Luxembourg           | `eu_vat`     | European VAT Number                                                                                     |
// | Malaysia             | `my_frp`     | Malaysian FRP Number                                                                                    |
// | Malaysia             | `my_itn`     | Malaysian ITN                                                                                           |
// | Malaysia             | `my_sst`     | Malaysian SST Number                                                                                    |
// | Malta                | `eu_vat `    | European VAT Number                                                                                     |
// | Mexico               | `mx_rfc`     | Mexican RFC Number                                                                                      |
// | Netherlands          | `eu_vat`     | European VAT Number                                                                                     |
// | New Zealand          | `nz_gst`     | New Zealand GST Number                                                                                  |
// | Nigeria              | `ng_tin`     | Nigerian Tax Identification Number                                                                      |
// | Norway               | `no_vat`     | Norwegian VAT Number                                                                                    |
// | Norway               | `no_voec`    | Norwegian VAT on e-commerce Number                                                                      |
// | Oman                 | `om_vat`     | Omani VAT Number                                                                                        |
// | Peru                 | `pe_ruc`     | Peruvian RUC Number                                                                                     |
// | Philippines          | `ph_tin `    | Philippines Tax Identification Number                                                                   |
// | Poland               | `eu_vat`     | European VAT Number                                                                                     |
// | Portugal             | `eu_vat`     | European VAT Number                                                                                     |
// | Romania              | `eu_vat`     | European VAT Number                                                                                     |
// | Romania              | `ro_tin`     | Romanian Tax ID Number                                                                                  |
// | Russia               | `ru_inn`     | Russian INN                                                                                             |
// | Russia               | `ru_kpp`     | Russian KPP                                                                                             |
// | Saudi Arabia         | `sa_vat`     | Saudi Arabia VAT                                                                                        |
// | Serbia               | `rs_pib`     | Serbian PIB Number                                                                                      |
// | Singapore            | `sg_gst`     | Singaporean GST                                                                                         |
// | Singapore            | `sg_uen`     | Singaporean UEN                                                                                         |
// | Slovakia             | `eu_vat`     | European VAT Number                                                                                     |
// | Slovenia             | `eu_vat`     | European VAT Number                                                                                     |
// | Slovenia             | `si_tin`     | Slovenia Tax Number (davčna številka)                                                                   |
// | South Africa         | `za_vat`     | South African VAT Number                                                                                |
// | South Korea          | `kr_brn`     | Korean BRN                                                                                              |
// | Spain                | `es_cif`     | Spanish NIF Number (previously Spanish CIF Number)                                                      |
// | Spain                | `eu_vat`     | European VAT Number                                                                                     |
// | Sweden               | `eu_vat`     | European VAT Number                                                                                     |
// | Switzerland          | `ch_vat`     | Switzerland VAT Number                                                                                  |
// | Taiwan               | `tw_vat`     | Taiwanese VAT                                                                                           |
// | Thailand             | `th_vat`     | Thai VAT                                                                                                |
// | Turkey               | `tr_tin`     | Turkish Tax Identification Number                                                                       |
// | Ukraine              | `ua_vat`     | Ukrainian VAT                                                                                           |
// | United Arab Emirates | `ae_trn`     | United Arab Emirates TRN                                                                                |
// | United Kingdom       | `eu_vat`     | Northern Ireland VAT Number                                                                             |
// | United Kingdom       | `gb_vat`     | United Kingdom VAT Number                                                                               |
// | United States        | `us_ein`     | United States EIN                                                                                       |
// | Uruguay              | `uy_ruc`     | Uruguayan RUC Number                                                                                    |
// | Venezuela            | `ve_rif`     | Venezuelan RIF Number                                                                                   |
// | Vietnam              | `vn_tin`     | Vietnamese Tax ID Number                                                                                |
type CustomerTaxIDModelParam struct {
	Country param.Field[CustomerTaxIDModelCountry] `json:"country,required"`
	Type    param.Field[CustomerTaxIDModelType]    `json:"type,required"`
	Value   param.Field[string]                    `json:"value,required"`
}

func (r CustomerTaxIDModelParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type DimensionalPriceConfigurationModel struct {
	DimensionValues         []string                               `json:"dimension_values,required"`
	DimensionalPriceGroupID string                                 `json:"dimensional_price_group_id,required"`
	JSON                    dimensionalPriceConfigurationModelJSON `json:"-"`
}

// dimensionalPriceConfigurationModelJSON contains the JSON metadata for the struct
// [DimensionalPriceConfigurationModel]
type dimensionalPriceConfigurationModelJSON struct {
	DimensionValues         apijson.Field
	DimensionalPriceGroupID apijson.Field
	raw                     string
	ExtraFields             map[string]apijson.Field
}

func (r *DimensionalPriceConfigurationModel) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r dimensionalPriceConfigurationModelJSON) RawJSON() string {
	return r.raw
}

// A dimensional price group is used to partition the result of a billable metric
// by a set of dimensions. Prices in a price group must specify the parition used
// to derive their usage.
type DimensionalPriceGroupModel struct {
	ID string `json:"id,required"`
	// The billable metric associated with this dimensional price group. All prices
	// associated with this dimensional price group will be computed using this
	// billable metric.
	BillableMetricID string `json:"billable_metric_id,required"`
	// The dimensions that this dimensional price group is defined over
	Dimensions []string `json:"dimensions,required"`
	// An alias for the dimensional price group
	ExternalDimensionalPriceGroupID string `json:"external_dimensional_price_group_id,required,nullable"`
	// User specified key-value pairs for the resource. If not present, this defaults
	// to an empty dictionary. Individual keys can be removed by setting the value to
	// `null`, and the entire metadata mapping can be cleared by setting `metadata` to
	// `null`.
	Metadata map[string]string `json:"metadata,required"`
	// The name of the dimensional price group
	Name string                         `json:"name,required"`
	JSON dimensionalPriceGroupModelJSON `json:"-"`
}

// dimensionalPriceGroupModelJSON contains the JSON metadata for the struct
// [DimensionalPriceGroupModel]
type dimensionalPriceGroupModelJSON struct {
	ID                              apijson.Field
	BillableMetricID                apijson.Field
	Dimensions                      apijson.Field
	ExternalDimensionalPriceGroupID apijson.Field
	Metadata                        apijson.Field
	Name                            apijson.Field
	raw                             string
	ExtraFields                     map[string]apijson.Field
}

func (r *DimensionalPriceGroupModel) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r dimensionalPriceGroupModelJSON) RawJSON() string {
	return r.raw
}

type Discount struct {
	// This field can have the runtime type of [[]string].
	AppliesToPriceIDs interface{}          `json:"applies_to_price_ids,required"`
	DiscountType      DiscountDiscountType `json:"discount_type,required"`
	// Only available if discount_type is `amount`.
	AmountDiscount string `json:"amount_discount"`
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
	AppliesToPriceIDs       apijson.Field
	DiscountType            apijson.Field
	AmountDiscount          apijson.Field
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
// Possible runtime types of the union are [shared.PercentageDiscount],
// [shared.TrialDiscount], [shared.DiscountUsageDiscount], [shared.AmountDiscount].
func (r Discount) AsUnion() DiscountUnion {
	return r.union
}

// Union satisfied by [shared.PercentageDiscount], [shared.TrialDiscount],
// [shared.DiscountUsageDiscount] or [shared.AmountDiscount].
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
			Type:               reflect.TypeOf(DiscountUsageDiscount{}),
			DiscriminatorValue: "usage",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(AmountDiscount{}),
			DiscriminatorValue: "amount",
		},
	)
}

type DiscountUsageDiscount struct {
	// List of price_ids that this discount applies to. For plan/plan phase discounts,
	// this can be a subset of prices.
	AppliesToPriceIDs []string                          `json:"applies_to_price_ids,required"`
	DiscountType      DiscountUsageDiscountDiscountType `json:"discount_type,required"`
	// Only available if discount_type is `usage`. Number of usage units that this
	// discount is for
	UsageDiscount float64                   `json:"usage_discount,required"`
	Reason        string                    `json:"reason,nullable"`
	JSON          discountUsageDiscountJSON `json:"-"`
}

// discountUsageDiscountJSON contains the JSON metadata for the struct
// [DiscountUsageDiscount]
type discountUsageDiscountJSON struct {
	AppliesToPriceIDs apijson.Field
	DiscountType      apijson.Field
	UsageDiscount     apijson.Field
	Reason            apijson.Field
	raw               string
	ExtraFields       map[string]apijson.Field
}

func (r *DiscountUsageDiscount) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r discountUsageDiscountJSON) RawJSON() string {
	return r.raw
}

func (r DiscountUsageDiscount) ImplementsDiscount() {}

type DiscountUsageDiscountDiscountType string

const (
	DiscountUsageDiscountDiscountTypeUsage DiscountUsageDiscountDiscountType = "usage"
)

func (r DiscountUsageDiscountDiscountType) IsKnown() bool {
	switch r {
	case DiscountUsageDiscountDiscountTypeUsage:
		return true
	}
	return false
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
	AppliesToPriceIDs param.Field[interface{}]          `json:"applies_to_price_ids,required"`
	DiscountType      param.Field[DiscountDiscountType] `json:"discount_type,required"`
	// Only available if discount_type is `amount`.
	AmountDiscount param.Field[string] `json:"amount_discount"`
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
// [shared.DiscountUsageDiscountParam], [shared.AmountDiscountParam],
// [DiscountParam].
type DiscountUnionParam interface {
	ImplementsDiscountUnionParam()
}

type DiscountUsageDiscountParam struct {
	// List of price_ids that this discount applies to. For plan/plan phase discounts,
	// this can be a subset of prices.
	AppliesToPriceIDs param.Field[[]string]                          `json:"applies_to_price_ids,required"`
	DiscountType      param.Field[DiscountUsageDiscountDiscountType] `json:"discount_type,required"`
	// Only available if discount_type is `usage`. Number of usage units that this
	// discount is for
	UsageDiscount param.Field[float64] `json:"usage_discount,required"`
	Reason        param.Field[string]  `json:"reason"`
}

func (r DiscountUsageDiscountParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r DiscountUsageDiscountParam) ImplementsDiscountUnionParam() {}

type DiscountOverrideModelParam struct {
	DiscountType param.Field[DiscountOverrideModelDiscountType] `json:"discount_type,required"`
	// Only available if discount_type is `amount`.
	AmountDiscount param.Field[string] `json:"amount_discount"`
	// Only available if discount_type is `percentage`. This is a number between 0
	// and 1.
	PercentageDiscount param.Field[float64] `json:"percentage_discount"`
	// Only available if discount_type is `usage`. Number of usage units that this
	// discount is for
	UsageDiscount param.Field[float64] `json:"usage_discount"`
}

func (r DiscountOverrideModelParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type DiscountOverrideModelDiscountType string

const (
	DiscountOverrideModelDiscountTypePercentage DiscountOverrideModelDiscountType = "percentage"
	DiscountOverrideModelDiscountTypeUsage      DiscountOverrideModelDiscountType = "usage"
	DiscountOverrideModelDiscountTypeAmount     DiscountOverrideModelDiscountType = "amount"
)

func (r DiscountOverrideModelDiscountType) IsKnown() bool {
	switch r {
	case DiscountOverrideModelDiscountTypePercentage, DiscountOverrideModelDiscountTypeUsage, DiscountOverrideModelDiscountTypeAmount:
		return true
	}
	return false
}

type FixedFeeQuantityScheduleEntryModel struct {
	EndDate   time.Time                              `json:"end_date,required,nullable" format:"date-time"`
	PriceID   string                                 `json:"price_id,required"`
	Quantity  float64                                `json:"quantity,required"`
	StartDate time.Time                              `json:"start_date,required" format:"date-time"`
	JSON      fixedFeeQuantityScheduleEntryModelJSON `json:"-"`
}

// fixedFeeQuantityScheduleEntryModelJSON contains the JSON metadata for the struct
// [FixedFeeQuantityScheduleEntryModel]
type fixedFeeQuantityScheduleEntryModelJSON struct {
	EndDate     apijson.Field
	PriceID     apijson.Field
	Quantity    apijson.Field
	StartDate   apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *FixedFeeQuantityScheduleEntryModel) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r fixedFeeQuantityScheduleEntryModelJSON) RawJSON() string {
	return r.raw
}

type InvoiceLevelDiscount struct {
	// This field can have the runtime type of [[]string].
	AppliesToPriceIDs interface{}                      `json:"applies_to_price_ids,required"`
	DiscountType      InvoiceLevelDiscountDiscountType `json:"discount_type,required"`
	// Only available if discount_type is `amount`.
	AmountDiscount string `json:"amount_discount"`
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
	AppliesToPriceIDs       apijson.Field
	DiscountType            apijson.Field
	AmountDiscount          apijson.Field
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
// Possible runtime types of the union are [shared.PercentageDiscount],
// [shared.AmountDiscount], [shared.TrialDiscount].
func (r InvoiceLevelDiscount) AsUnion() InvoiceLevelDiscountUnion {
	return r.union
}

// Union satisfied by [shared.PercentageDiscount], [shared.AmountDiscount] or
// [shared.TrialDiscount].
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

type InvoiceLineItemModel struct {
	// A unique ID for this line item.
	ID string `json:"id,required"`
	// The line amount after any adjustments and before overage conversion, credits and
	// partial invoicing.
	AdjustedSubtotal string `json:"adjusted_subtotal,required"`
	// All adjustments (ie. maximums, minimums, discounts) applied to the line item.
	Adjustments []InvoiceLineItemModelAdjustment `json:"adjustments,required"`
	// The final amount for a line item after all adjustments and pre paid credits have
	// been applied.
	Amount string `json:"amount,required"`
	// The number of prepaid credits applied.
	CreditsApplied string   `json:"credits_applied,required"`
	Discount       Discount `json:"discount,required,nullable"`
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
	Maximum MaximumModel `json:"maximum,required,nullable"`
	// This field is deprecated in favor of `adjustments`.
	//
	// Deprecated: deprecated
	MaximumAmount string `json:"maximum_amount,required,nullable"`
	// This field is deprecated in favor of `adjustments`.
	//
	// Deprecated: deprecated
	Minimum MinimumModel `json:"minimum,required,nullable"`
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
	Price PriceModel `json:"price,required,nullable"`
	// Either the fixed fee quantity or the usage during the service period.
	Quantity float64 `json:"quantity,required"`
	// The start date of the range of time applied for this line item's price.
	StartDate time.Time `json:"start_date,required" format:"date-time"`
	// For complex pricing structures, the line item can be broken down further in
	// `sub_line_items`.
	SubLineItems []InvoiceLineItemModelSubLineItem `json:"sub_line_items,required"`
	// The line amount before before any adjustments.
	Subtotal string `json:"subtotal,required"`
	// An array of tax rates and their incurred tax amounts. Empty if no tax
	// integration is configured.
	TaxAmounts []TaxAmountModel `json:"tax_amounts,required"`
	// A list of customer ids that were used to calculate the usage for this line item.
	UsageCustomerIDs []string                 `json:"usage_customer_ids,required,nullable"`
	JSON             invoiceLineItemModelJSON `json:"-"`
}

// invoiceLineItemModelJSON contains the JSON metadata for the struct
// [InvoiceLineItemModel]
type invoiceLineItemModelJSON struct {
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

func (r *InvoiceLineItemModel) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r invoiceLineItemModelJSON) RawJSON() string {
	return r.raw
}

type InvoiceLineItemModelAdjustment struct {
	ID             string                                        `json:"id,required"`
	AdjustmentType InvoiceLineItemModelAdjustmentsAdjustmentType `json:"adjustment_type,required"`
	// The value applied by an adjustment.
	Amount string `json:"amount,required"`
	// This field can have the runtime type of [[]string].
	AppliesToPriceIDs interface{} `json:"applies_to_price_ids,required"`
	// True for adjustments that apply to an entire invocice, false for adjustments
	// that apply to only one price.
	IsInvoiceLevel bool `json:"is_invoice_level,required"`
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
	UsageDiscount float64                            `json:"usage_discount"`
	JSON          invoiceLineItemModelAdjustmentJSON `json:"-"`
	union         InvoiceLineItemModelAdjustmentsUnion
}

// invoiceLineItemModelAdjustmentJSON contains the JSON metadata for the struct
// [InvoiceLineItemModelAdjustment]
type invoiceLineItemModelAdjustmentJSON struct {
	ID                 apijson.Field
	AdjustmentType     apijson.Field
	Amount             apijson.Field
	AppliesToPriceIDs  apijson.Field
	IsInvoiceLevel     apijson.Field
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

func (r invoiceLineItemModelAdjustmentJSON) RawJSON() string {
	return r.raw
}

func (r *InvoiceLineItemModelAdjustment) UnmarshalJSON(data []byte) (err error) {
	*r = InvoiceLineItemModelAdjustment{}
	err = apijson.UnmarshalRoot(data, &r.union)
	if err != nil {
		return err
	}
	return apijson.Port(r.union, &r)
}

// AsUnion returns a [InvoiceLineItemModelAdjustmentsUnion] interface which you can
// cast to the specific types for more type safety.
//
// Possible runtime types of the union are
// [shared.InvoiceLineItemModelAdjustmentsMonetaryUsageDiscountAdjustment],
// [shared.InvoiceLineItemModelAdjustmentsMonetaryAmountDiscountAdjustment],
// [shared.InvoiceLineItemModelAdjustmentsMonetaryPercentageDiscountAdjustment],
// [shared.InvoiceLineItemModelAdjustmentsMonetaryMinimumAdjustment],
// [shared.InvoiceLineItemModelAdjustmentsMonetaryMaximumAdjustment].
func (r InvoiceLineItemModelAdjustment) AsUnion() InvoiceLineItemModelAdjustmentsUnion {
	return r.union
}

// Union satisfied by
// [shared.InvoiceLineItemModelAdjustmentsMonetaryUsageDiscountAdjustment],
// [shared.InvoiceLineItemModelAdjustmentsMonetaryAmountDiscountAdjustment],
// [shared.InvoiceLineItemModelAdjustmentsMonetaryPercentageDiscountAdjustment],
// [shared.InvoiceLineItemModelAdjustmentsMonetaryMinimumAdjustment] or
// [shared.InvoiceLineItemModelAdjustmentsMonetaryMaximumAdjustment].
type InvoiceLineItemModelAdjustmentsUnion interface {
	implementsInvoiceLineItemModelAdjustment()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*InvoiceLineItemModelAdjustmentsUnion)(nil)).Elem(),
		"adjustment_type",
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(InvoiceLineItemModelAdjustmentsMonetaryUsageDiscountAdjustment{}),
			DiscriminatorValue: "usage_discount",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(InvoiceLineItemModelAdjustmentsMonetaryAmountDiscountAdjustment{}),
			DiscriminatorValue: "amount_discount",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(InvoiceLineItemModelAdjustmentsMonetaryPercentageDiscountAdjustment{}),
			DiscriminatorValue: "percentage_discount",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(InvoiceLineItemModelAdjustmentsMonetaryMinimumAdjustment{}),
			DiscriminatorValue: "minimum",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(InvoiceLineItemModelAdjustmentsMonetaryMaximumAdjustment{}),
			DiscriminatorValue: "maximum",
		},
	)
}

type InvoiceLineItemModelAdjustmentsMonetaryUsageDiscountAdjustment struct {
	ID             string                                                                       `json:"id,required"`
	AdjustmentType InvoiceLineItemModelAdjustmentsMonetaryUsageDiscountAdjustmentAdjustmentType `json:"adjustment_type,required"`
	// The value applied by an adjustment.
	Amount string `json:"amount,required"`
	// The price IDs that this adjustment applies to.
	AppliesToPriceIDs []string `json:"applies_to_price_ids,required"`
	// True for adjustments that apply to an entire invocice, false for adjustments
	// that apply to only one price.
	IsInvoiceLevel bool `json:"is_invoice_level,required"`
	// The reason for the adjustment.
	Reason string `json:"reason,required,nullable"`
	// The number of usage units by which to discount the price this adjustment applies
	// to in a given billing period.
	UsageDiscount float64                                                            `json:"usage_discount,required"`
	JSON          invoiceLineItemModelAdjustmentsMonetaryUsageDiscountAdjustmentJSON `json:"-"`
}

// invoiceLineItemModelAdjustmentsMonetaryUsageDiscountAdjustmentJSON contains the
// JSON metadata for the struct
// [InvoiceLineItemModelAdjustmentsMonetaryUsageDiscountAdjustment]
type invoiceLineItemModelAdjustmentsMonetaryUsageDiscountAdjustmentJSON struct {
	ID                apijson.Field
	AdjustmentType    apijson.Field
	Amount            apijson.Field
	AppliesToPriceIDs apijson.Field
	IsInvoiceLevel    apijson.Field
	Reason            apijson.Field
	UsageDiscount     apijson.Field
	raw               string
	ExtraFields       map[string]apijson.Field
}

func (r *InvoiceLineItemModelAdjustmentsMonetaryUsageDiscountAdjustment) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r invoiceLineItemModelAdjustmentsMonetaryUsageDiscountAdjustmentJSON) RawJSON() string {
	return r.raw
}

func (r InvoiceLineItemModelAdjustmentsMonetaryUsageDiscountAdjustment) implementsInvoiceLineItemModelAdjustment() {
}

type InvoiceLineItemModelAdjustmentsMonetaryUsageDiscountAdjustmentAdjustmentType string

const (
	InvoiceLineItemModelAdjustmentsMonetaryUsageDiscountAdjustmentAdjustmentTypeUsageDiscount InvoiceLineItemModelAdjustmentsMonetaryUsageDiscountAdjustmentAdjustmentType = "usage_discount"
)

func (r InvoiceLineItemModelAdjustmentsMonetaryUsageDiscountAdjustmentAdjustmentType) IsKnown() bool {
	switch r {
	case InvoiceLineItemModelAdjustmentsMonetaryUsageDiscountAdjustmentAdjustmentTypeUsageDiscount:
		return true
	}
	return false
}

type InvoiceLineItemModelAdjustmentsMonetaryAmountDiscountAdjustment struct {
	ID             string                                                                        `json:"id,required"`
	AdjustmentType InvoiceLineItemModelAdjustmentsMonetaryAmountDiscountAdjustmentAdjustmentType `json:"adjustment_type,required"`
	// The value applied by an adjustment.
	Amount string `json:"amount,required"`
	// The amount by which to discount the prices this adjustment applies to in a given
	// billing period.
	AmountDiscount string `json:"amount_discount,required"`
	// The price IDs that this adjustment applies to.
	AppliesToPriceIDs []string `json:"applies_to_price_ids,required"`
	// True for adjustments that apply to an entire invocice, false for adjustments
	// that apply to only one price.
	IsInvoiceLevel bool `json:"is_invoice_level,required"`
	// The reason for the adjustment.
	Reason string                                                              `json:"reason,required,nullable"`
	JSON   invoiceLineItemModelAdjustmentsMonetaryAmountDiscountAdjustmentJSON `json:"-"`
}

// invoiceLineItemModelAdjustmentsMonetaryAmountDiscountAdjustmentJSON contains the
// JSON metadata for the struct
// [InvoiceLineItemModelAdjustmentsMonetaryAmountDiscountAdjustment]
type invoiceLineItemModelAdjustmentsMonetaryAmountDiscountAdjustmentJSON struct {
	ID                apijson.Field
	AdjustmentType    apijson.Field
	Amount            apijson.Field
	AmountDiscount    apijson.Field
	AppliesToPriceIDs apijson.Field
	IsInvoiceLevel    apijson.Field
	Reason            apijson.Field
	raw               string
	ExtraFields       map[string]apijson.Field
}

func (r *InvoiceLineItemModelAdjustmentsMonetaryAmountDiscountAdjustment) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r invoiceLineItemModelAdjustmentsMonetaryAmountDiscountAdjustmentJSON) RawJSON() string {
	return r.raw
}

func (r InvoiceLineItemModelAdjustmentsMonetaryAmountDiscountAdjustment) implementsInvoiceLineItemModelAdjustment() {
}

type InvoiceLineItemModelAdjustmentsMonetaryAmountDiscountAdjustmentAdjustmentType string

const (
	InvoiceLineItemModelAdjustmentsMonetaryAmountDiscountAdjustmentAdjustmentTypeAmountDiscount InvoiceLineItemModelAdjustmentsMonetaryAmountDiscountAdjustmentAdjustmentType = "amount_discount"
)

func (r InvoiceLineItemModelAdjustmentsMonetaryAmountDiscountAdjustmentAdjustmentType) IsKnown() bool {
	switch r {
	case InvoiceLineItemModelAdjustmentsMonetaryAmountDiscountAdjustmentAdjustmentTypeAmountDiscount:
		return true
	}
	return false
}

type InvoiceLineItemModelAdjustmentsMonetaryPercentageDiscountAdjustment struct {
	ID             string                                                                            `json:"id,required"`
	AdjustmentType InvoiceLineItemModelAdjustmentsMonetaryPercentageDiscountAdjustmentAdjustmentType `json:"adjustment_type,required"`
	// The value applied by an adjustment.
	Amount string `json:"amount,required"`
	// The price IDs that this adjustment applies to.
	AppliesToPriceIDs []string `json:"applies_to_price_ids,required"`
	// True for adjustments that apply to an entire invocice, false for adjustments
	// that apply to only one price.
	IsInvoiceLevel bool `json:"is_invoice_level,required"`
	// The percentage (as a value between 0 and 1) by which to discount the price
	// intervals this adjustment applies to in a given billing period.
	PercentageDiscount float64 `json:"percentage_discount,required"`
	// The reason for the adjustment.
	Reason string                                                                  `json:"reason,required,nullable"`
	JSON   invoiceLineItemModelAdjustmentsMonetaryPercentageDiscountAdjustmentJSON `json:"-"`
}

// invoiceLineItemModelAdjustmentsMonetaryPercentageDiscountAdjustmentJSON contains
// the JSON metadata for the struct
// [InvoiceLineItemModelAdjustmentsMonetaryPercentageDiscountAdjustment]
type invoiceLineItemModelAdjustmentsMonetaryPercentageDiscountAdjustmentJSON struct {
	ID                 apijson.Field
	AdjustmentType     apijson.Field
	Amount             apijson.Field
	AppliesToPriceIDs  apijson.Field
	IsInvoiceLevel     apijson.Field
	PercentageDiscount apijson.Field
	Reason             apijson.Field
	raw                string
	ExtraFields        map[string]apijson.Field
}

func (r *InvoiceLineItemModelAdjustmentsMonetaryPercentageDiscountAdjustment) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r invoiceLineItemModelAdjustmentsMonetaryPercentageDiscountAdjustmentJSON) RawJSON() string {
	return r.raw
}

func (r InvoiceLineItemModelAdjustmentsMonetaryPercentageDiscountAdjustment) implementsInvoiceLineItemModelAdjustment() {
}

type InvoiceLineItemModelAdjustmentsMonetaryPercentageDiscountAdjustmentAdjustmentType string

const (
	InvoiceLineItemModelAdjustmentsMonetaryPercentageDiscountAdjustmentAdjustmentTypePercentageDiscount InvoiceLineItemModelAdjustmentsMonetaryPercentageDiscountAdjustmentAdjustmentType = "percentage_discount"
)

func (r InvoiceLineItemModelAdjustmentsMonetaryPercentageDiscountAdjustmentAdjustmentType) IsKnown() bool {
	switch r {
	case InvoiceLineItemModelAdjustmentsMonetaryPercentageDiscountAdjustmentAdjustmentTypePercentageDiscount:
		return true
	}
	return false
}

type InvoiceLineItemModelAdjustmentsMonetaryMinimumAdjustment struct {
	ID             string                                                                 `json:"id,required"`
	AdjustmentType InvoiceLineItemModelAdjustmentsMonetaryMinimumAdjustmentAdjustmentType `json:"adjustment_type,required"`
	// The value applied by an adjustment.
	Amount string `json:"amount,required"`
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
	// The reason for the adjustment.
	Reason string                                                       `json:"reason,required,nullable"`
	JSON   invoiceLineItemModelAdjustmentsMonetaryMinimumAdjustmentJSON `json:"-"`
}

// invoiceLineItemModelAdjustmentsMonetaryMinimumAdjustmentJSON contains the JSON
// metadata for the struct
// [InvoiceLineItemModelAdjustmentsMonetaryMinimumAdjustment]
type invoiceLineItemModelAdjustmentsMonetaryMinimumAdjustmentJSON struct {
	ID                apijson.Field
	AdjustmentType    apijson.Field
	Amount            apijson.Field
	AppliesToPriceIDs apijson.Field
	IsInvoiceLevel    apijson.Field
	ItemID            apijson.Field
	MinimumAmount     apijson.Field
	Reason            apijson.Field
	raw               string
	ExtraFields       map[string]apijson.Field
}

func (r *InvoiceLineItemModelAdjustmentsMonetaryMinimumAdjustment) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r invoiceLineItemModelAdjustmentsMonetaryMinimumAdjustmentJSON) RawJSON() string {
	return r.raw
}

func (r InvoiceLineItemModelAdjustmentsMonetaryMinimumAdjustment) implementsInvoiceLineItemModelAdjustment() {
}

type InvoiceLineItemModelAdjustmentsMonetaryMinimumAdjustmentAdjustmentType string

const (
	InvoiceLineItemModelAdjustmentsMonetaryMinimumAdjustmentAdjustmentTypeMinimum InvoiceLineItemModelAdjustmentsMonetaryMinimumAdjustmentAdjustmentType = "minimum"
)

func (r InvoiceLineItemModelAdjustmentsMonetaryMinimumAdjustmentAdjustmentType) IsKnown() bool {
	switch r {
	case InvoiceLineItemModelAdjustmentsMonetaryMinimumAdjustmentAdjustmentTypeMinimum:
		return true
	}
	return false
}

type InvoiceLineItemModelAdjustmentsMonetaryMaximumAdjustment struct {
	ID             string                                                                 `json:"id,required"`
	AdjustmentType InvoiceLineItemModelAdjustmentsMonetaryMaximumAdjustmentAdjustmentType `json:"adjustment_type,required"`
	// The value applied by an adjustment.
	Amount string `json:"amount,required"`
	// The price IDs that this adjustment applies to.
	AppliesToPriceIDs []string `json:"applies_to_price_ids,required"`
	// True for adjustments that apply to an entire invocice, false for adjustments
	// that apply to only one price.
	IsInvoiceLevel bool `json:"is_invoice_level,required"`
	// The maximum amount to charge in a given billing period for the prices this
	// adjustment applies to.
	MaximumAmount string `json:"maximum_amount,required"`
	// The reason for the adjustment.
	Reason string                                                       `json:"reason,required,nullable"`
	JSON   invoiceLineItemModelAdjustmentsMonetaryMaximumAdjustmentJSON `json:"-"`
}

// invoiceLineItemModelAdjustmentsMonetaryMaximumAdjustmentJSON contains the JSON
// metadata for the struct
// [InvoiceLineItemModelAdjustmentsMonetaryMaximumAdjustment]
type invoiceLineItemModelAdjustmentsMonetaryMaximumAdjustmentJSON struct {
	ID                apijson.Field
	AdjustmentType    apijson.Field
	Amount            apijson.Field
	AppliesToPriceIDs apijson.Field
	IsInvoiceLevel    apijson.Field
	MaximumAmount     apijson.Field
	Reason            apijson.Field
	raw               string
	ExtraFields       map[string]apijson.Field
}

func (r *InvoiceLineItemModelAdjustmentsMonetaryMaximumAdjustment) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r invoiceLineItemModelAdjustmentsMonetaryMaximumAdjustmentJSON) RawJSON() string {
	return r.raw
}

func (r InvoiceLineItemModelAdjustmentsMonetaryMaximumAdjustment) implementsInvoiceLineItemModelAdjustment() {
}

type InvoiceLineItemModelAdjustmentsMonetaryMaximumAdjustmentAdjustmentType string

const (
	InvoiceLineItemModelAdjustmentsMonetaryMaximumAdjustmentAdjustmentTypeMaximum InvoiceLineItemModelAdjustmentsMonetaryMaximumAdjustmentAdjustmentType = "maximum"
)

func (r InvoiceLineItemModelAdjustmentsMonetaryMaximumAdjustmentAdjustmentType) IsKnown() bool {
	switch r {
	case InvoiceLineItemModelAdjustmentsMonetaryMaximumAdjustmentAdjustmentTypeMaximum:
		return true
	}
	return false
}

type InvoiceLineItemModelAdjustmentsAdjustmentType string

const (
	InvoiceLineItemModelAdjustmentsAdjustmentTypeUsageDiscount      InvoiceLineItemModelAdjustmentsAdjustmentType = "usage_discount"
	InvoiceLineItemModelAdjustmentsAdjustmentTypeAmountDiscount     InvoiceLineItemModelAdjustmentsAdjustmentType = "amount_discount"
	InvoiceLineItemModelAdjustmentsAdjustmentTypePercentageDiscount InvoiceLineItemModelAdjustmentsAdjustmentType = "percentage_discount"
	InvoiceLineItemModelAdjustmentsAdjustmentTypeMinimum            InvoiceLineItemModelAdjustmentsAdjustmentType = "minimum"
	InvoiceLineItemModelAdjustmentsAdjustmentTypeMaximum            InvoiceLineItemModelAdjustmentsAdjustmentType = "maximum"
)

func (r InvoiceLineItemModelAdjustmentsAdjustmentType) IsKnown() bool {
	switch r {
	case InvoiceLineItemModelAdjustmentsAdjustmentTypeUsageDiscount, InvoiceLineItemModelAdjustmentsAdjustmentTypeAmountDiscount, InvoiceLineItemModelAdjustmentsAdjustmentTypePercentageDiscount, InvoiceLineItemModelAdjustmentsAdjustmentTypeMinimum, InvoiceLineItemModelAdjustmentsAdjustmentTypeMaximum:
		return true
	}
	return false
}

type InvoiceLineItemModelSubLineItem struct {
	// The total amount for this sub line item.
	Amount   string                               `json:"amount,required"`
	Grouping SubLineItemGroupingModel             `json:"grouping,required,nullable"`
	Name     string                               `json:"name,required"`
	Quantity float64                              `json:"quantity,required"`
	Type     InvoiceLineItemModelSubLineItemsType `json:"type,required"`
	// This field can have the runtime type of
	// [InvoiceLineItemModelSubLineItemsMatrixSubLineItemMatrixConfig].
	MatrixConfig interface{} `json:"matrix_config"`
	// This field can have the runtime type of
	// [InvoiceLineItemModelSubLineItemsTierSubLineItemTierConfig].
	TierConfig interface{}                         `json:"tier_config"`
	JSON       invoiceLineItemModelSubLineItemJSON `json:"-"`
	union      InvoiceLineItemModelSubLineItemsUnion
}

// invoiceLineItemModelSubLineItemJSON contains the JSON metadata for the struct
// [InvoiceLineItemModelSubLineItem]
type invoiceLineItemModelSubLineItemJSON struct {
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

func (r invoiceLineItemModelSubLineItemJSON) RawJSON() string {
	return r.raw
}

func (r *InvoiceLineItemModelSubLineItem) UnmarshalJSON(data []byte) (err error) {
	*r = InvoiceLineItemModelSubLineItem{}
	err = apijson.UnmarshalRoot(data, &r.union)
	if err != nil {
		return err
	}
	return apijson.Port(r.union, &r)
}

// AsUnion returns a [InvoiceLineItemModelSubLineItemsUnion] interface which you
// can cast to the specific types for more type safety.
//
// Possible runtime types of the union are
// [shared.InvoiceLineItemModelSubLineItemsMatrixSubLineItem],
// [shared.InvoiceLineItemModelSubLineItemsTierSubLineItem],
// [shared.InvoiceLineItemModelSubLineItemsOtherSubLineItem].
func (r InvoiceLineItemModelSubLineItem) AsUnion() InvoiceLineItemModelSubLineItemsUnion {
	return r.union
}

// Union satisfied by [shared.InvoiceLineItemModelSubLineItemsMatrixSubLineItem],
// [shared.InvoiceLineItemModelSubLineItemsTierSubLineItem] or
// [shared.InvoiceLineItemModelSubLineItemsOtherSubLineItem].
type InvoiceLineItemModelSubLineItemsUnion interface {
	implementsInvoiceLineItemModelSubLineItem()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*InvoiceLineItemModelSubLineItemsUnion)(nil)).Elem(),
		"type",
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(InvoiceLineItemModelSubLineItemsMatrixSubLineItem{}),
			DiscriminatorValue: "matrix",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(InvoiceLineItemModelSubLineItemsTierSubLineItem{}),
			DiscriminatorValue: "tier",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(InvoiceLineItemModelSubLineItemsOtherSubLineItem{}),
			DiscriminatorValue: "'null'",
		},
	)
}

type InvoiceLineItemModelSubLineItemsMatrixSubLineItem struct {
	// The total amount for this sub line item.
	Amount       string                                                        `json:"amount,required"`
	Grouping     SubLineItemGroupingModel                                      `json:"grouping,required,nullable"`
	MatrixConfig InvoiceLineItemModelSubLineItemsMatrixSubLineItemMatrixConfig `json:"matrix_config,required"`
	Name         string                                                        `json:"name,required"`
	Quantity     float64                                                       `json:"quantity,required"`
	Type         InvoiceLineItemModelSubLineItemsMatrixSubLineItemType         `json:"type,required"`
	JSON         invoiceLineItemModelSubLineItemsMatrixSubLineItemJSON         `json:"-"`
}

// invoiceLineItemModelSubLineItemsMatrixSubLineItemJSON contains the JSON metadata
// for the struct [InvoiceLineItemModelSubLineItemsMatrixSubLineItem]
type invoiceLineItemModelSubLineItemsMatrixSubLineItemJSON struct {
	Amount       apijson.Field
	Grouping     apijson.Field
	MatrixConfig apijson.Field
	Name         apijson.Field
	Quantity     apijson.Field
	Type         apijson.Field
	raw          string
	ExtraFields  map[string]apijson.Field
}

func (r *InvoiceLineItemModelSubLineItemsMatrixSubLineItem) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r invoiceLineItemModelSubLineItemsMatrixSubLineItemJSON) RawJSON() string {
	return r.raw
}

func (r InvoiceLineItemModelSubLineItemsMatrixSubLineItem) implementsInvoiceLineItemModelSubLineItem() {
}

type InvoiceLineItemModelSubLineItemsMatrixSubLineItemMatrixConfig struct {
	// The ordered dimension values for this line item.
	DimensionValues []string                                                          `json:"dimension_values,required"`
	JSON            invoiceLineItemModelSubLineItemsMatrixSubLineItemMatrixConfigJSON `json:"-"`
}

// invoiceLineItemModelSubLineItemsMatrixSubLineItemMatrixConfigJSON contains the
// JSON metadata for the struct
// [InvoiceLineItemModelSubLineItemsMatrixSubLineItemMatrixConfig]
type invoiceLineItemModelSubLineItemsMatrixSubLineItemMatrixConfigJSON struct {
	DimensionValues apijson.Field
	raw             string
	ExtraFields     map[string]apijson.Field
}

func (r *InvoiceLineItemModelSubLineItemsMatrixSubLineItemMatrixConfig) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r invoiceLineItemModelSubLineItemsMatrixSubLineItemMatrixConfigJSON) RawJSON() string {
	return r.raw
}

type InvoiceLineItemModelSubLineItemsMatrixSubLineItemType string

const (
	InvoiceLineItemModelSubLineItemsMatrixSubLineItemTypeMatrix InvoiceLineItemModelSubLineItemsMatrixSubLineItemType = "matrix"
)

func (r InvoiceLineItemModelSubLineItemsMatrixSubLineItemType) IsKnown() bool {
	switch r {
	case InvoiceLineItemModelSubLineItemsMatrixSubLineItemTypeMatrix:
		return true
	}
	return false
}

type InvoiceLineItemModelSubLineItemsTierSubLineItem struct {
	// The total amount for this sub line item.
	Amount     string                                                    `json:"amount,required"`
	Grouping   SubLineItemGroupingModel                                  `json:"grouping,required,nullable"`
	Name       string                                                    `json:"name,required"`
	Quantity   float64                                                   `json:"quantity,required"`
	TierConfig InvoiceLineItemModelSubLineItemsTierSubLineItemTierConfig `json:"tier_config,required"`
	Type       InvoiceLineItemModelSubLineItemsTierSubLineItemType       `json:"type,required"`
	JSON       invoiceLineItemModelSubLineItemsTierSubLineItemJSON       `json:"-"`
}

// invoiceLineItemModelSubLineItemsTierSubLineItemJSON contains the JSON metadata
// for the struct [InvoiceLineItemModelSubLineItemsTierSubLineItem]
type invoiceLineItemModelSubLineItemsTierSubLineItemJSON struct {
	Amount      apijson.Field
	Grouping    apijson.Field
	Name        apijson.Field
	Quantity    apijson.Field
	TierConfig  apijson.Field
	Type        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *InvoiceLineItemModelSubLineItemsTierSubLineItem) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r invoiceLineItemModelSubLineItemsTierSubLineItemJSON) RawJSON() string {
	return r.raw
}

func (r InvoiceLineItemModelSubLineItemsTierSubLineItem) implementsInvoiceLineItemModelSubLineItem() {
}

type InvoiceLineItemModelSubLineItemsTierSubLineItemTierConfig struct {
	FirstUnit  float64                                                       `json:"first_unit,required"`
	LastUnit   float64                                                       `json:"last_unit,required,nullable"`
	UnitAmount string                                                        `json:"unit_amount,required"`
	JSON       invoiceLineItemModelSubLineItemsTierSubLineItemTierConfigJSON `json:"-"`
}

// invoiceLineItemModelSubLineItemsTierSubLineItemTierConfigJSON contains the JSON
// metadata for the struct
// [InvoiceLineItemModelSubLineItemsTierSubLineItemTierConfig]
type invoiceLineItemModelSubLineItemsTierSubLineItemTierConfigJSON struct {
	FirstUnit   apijson.Field
	LastUnit    apijson.Field
	UnitAmount  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *InvoiceLineItemModelSubLineItemsTierSubLineItemTierConfig) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r invoiceLineItemModelSubLineItemsTierSubLineItemTierConfigJSON) RawJSON() string {
	return r.raw
}

type InvoiceLineItemModelSubLineItemsTierSubLineItemType string

const (
	InvoiceLineItemModelSubLineItemsTierSubLineItemTypeTier InvoiceLineItemModelSubLineItemsTierSubLineItemType = "tier"
)

func (r InvoiceLineItemModelSubLineItemsTierSubLineItemType) IsKnown() bool {
	switch r {
	case InvoiceLineItemModelSubLineItemsTierSubLineItemTypeTier:
		return true
	}
	return false
}

type InvoiceLineItemModelSubLineItemsOtherSubLineItem struct {
	// The total amount for this sub line item.
	Amount   string                                               `json:"amount,required"`
	Grouping SubLineItemGroupingModel                             `json:"grouping,required,nullable"`
	Name     string                                               `json:"name,required"`
	Quantity float64                                              `json:"quantity,required"`
	Type     InvoiceLineItemModelSubLineItemsOtherSubLineItemType `json:"type,required"`
	JSON     invoiceLineItemModelSubLineItemsOtherSubLineItemJSON `json:"-"`
}

// invoiceLineItemModelSubLineItemsOtherSubLineItemJSON contains the JSON metadata
// for the struct [InvoiceLineItemModelSubLineItemsOtherSubLineItem]
type invoiceLineItemModelSubLineItemsOtherSubLineItemJSON struct {
	Amount      apijson.Field
	Grouping    apijson.Field
	Name        apijson.Field
	Quantity    apijson.Field
	Type        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *InvoiceLineItemModelSubLineItemsOtherSubLineItem) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r invoiceLineItemModelSubLineItemsOtherSubLineItemJSON) RawJSON() string {
	return r.raw
}

func (r InvoiceLineItemModelSubLineItemsOtherSubLineItem) implementsInvoiceLineItemModelSubLineItem() {
}

type InvoiceLineItemModelSubLineItemsOtherSubLineItemType string

const (
	InvoiceLineItemModelSubLineItemsOtherSubLineItemTypeNull InvoiceLineItemModelSubLineItemsOtherSubLineItemType = "'null'"
)

func (r InvoiceLineItemModelSubLineItemsOtherSubLineItemType) IsKnown() bool {
	switch r {
	case InvoiceLineItemModelSubLineItemsOtherSubLineItemTypeNull:
		return true
	}
	return false
}

type InvoiceLineItemModelSubLineItemsType string

const (
	InvoiceLineItemModelSubLineItemsTypeMatrix InvoiceLineItemModelSubLineItemsType = "matrix"
	InvoiceLineItemModelSubLineItemsTypeTier   InvoiceLineItemModelSubLineItemsType = "tier"
	InvoiceLineItemModelSubLineItemsTypeNull   InvoiceLineItemModelSubLineItemsType = "'null'"
)

func (r InvoiceLineItemModelSubLineItemsType) IsKnown() bool {
	switch r {
	case InvoiceLineItemModelSubLineItemsTypeMatrix, InvoiceLineItemModelSubLineItemsTypeTier, InvoiceLineItemModelSubLineItemsTypeNull:
		return true
	}
	return false
}

// An [`Invoice`](/core-concepts#invoice) is a fundamental billing entity,
// representing the request for payment for a single subscription. This includes a
// set of line items, which correspond to prices in the subscription's plan and can
// represent fixed recurring fees or usage-based fees. They are generated at the
// end of a billing period, or as the result of an action, such as a cancellation.
type InvoiceModel struct {
	ID string `json:"id,required"`
	// This is the final amount required to be charged to the customer and reflects the
	// application of the customer balance to the `total` of the invoice.
	AmountDue      string              `json:"amount_due,required"`
	AutoCollection AutoCollectionModel `json:"auto_collection,required"`
	BillingAddress AddressModel        `json:"billing_address,required,nullable"`
	// The creation time of the resource in Orb.
	CreatedAt time.Time `json:"created_at,required" format:"date-time"`
	// A list of credit notes associated with the invoice
	CreditNotes []CreditNoteSummaryModel `json:"credit_notes,required"`
	// An ISO 4217 currency string or `credits`
	Currency                    string                            `json:"currency,required"`
	Customer                    CustomerMinifiedModel             `json:"customer,required"`
	CustomerBalanceTransactions []CustomerBalanceTransactionModel `json:"customer_balance_transactions,required"`
	// Tax IDs are commonly required to be displayed on customer invoices, which are
	// added to the headers of invoices.
	//
	// ### Supported Tax ID Countries and Types
	//
	// | Country              | Type         | Description                                                                                             |
	// | -------------------- | ------------ | ------------------------------------------------------------------------------------------------------- |
	// | Andorra              | `ad_nrt`     | Andorran NRT Number                                                                                     |
	// | Argentina            | `ar_cuit`    | Argentinian Tax ID Number                                                                               |
	// | Australia            | `au_abn`     | Australian Business Number (AU ABN)                                                                     |
	// | Australia            | `au_arn`     | Australian Taxation Office Reference Number                                                             |
	// | Austria              | `eu_vat`     | European VAT Number                                                                                     |
	// | Bahrain              | `bh_vat`     | Bahraini VAT Number                                                                                     |
	// | Belgium              | `eu_vat`     | European VAT Number                                                                                     |
	// | Bolivia              | `bo_tin`     | Bolivian Tax ID                                                                                         |
	// | Brazil               | `br_cnpj`    | Brazilian CNPJ Number                                                                                   |
	// | Brazil               | `br_cpf`     | Brazilian CPF Number                                                                                    |
	// | Bulgaria             | `bg_uic`     | Bulgaria Unified Identification Code                                                                    |
	// | Bulgaria             | `eu_vat`     | European VAT Number                                                                                     |
	// | Canada               | `ca_bn`      | Canadian BN                                                                                             |
	// | Canada               | `ca_gst_hst` | Canadian GST/HST Number                                                                                 |
	// | Canada               | `ca_pst_bc`  | Canadian PST Number (British Columbia)                                                                  |
	// | Canada               | `ca_pst_mb`  | Canadian PST Number (Manitoba)                                                                          |
	// | Canada               | `ca_pst_sk`  | Canadian PST Number (Saskatchewan)                                                                      |
	// | Canada               | `ca_qst`     | Canadian QST Number (Québec)                                                                            |
	// | Chile                | `cl_tin`     | Chilean TIN                                                                                             |
	// | China                | `cn_tin`     | Chinese Tax ID                                                                                          |
	// | Colombia             | `co_nit`     | Colombian NIT Number                                                                                    |
	// | Costa Rica           | `cr_tin`     | Costa Rican Tax ID                                                                                      |
	// | Croatia              | `eu_vat`     | European VAT Number                                                                                     |
	// | Cyprus               | `eu_vat`     | European VAT Number                                                                                     |
	// | Czech Republic       | `eu_vat`     | European VAT Number                                                                                     |
	// | Denmark              | `eu_vat`     | European VAT Number                                                                                     |
	// | Dominican Republic   | `do_rcn`     | Dominican RCN Number                                                                                    |
	// | Ecuador              | `ec_ruc`     | Ecuadorian RUC Number                                                                                   |
	// | Egypt                | `eg_tin`     | Egyptian Tax Identification Number                                                                      |
	// | El Salvador          | `sv_nit`     | El Salvadorian NIT Number                                                                               |
	// | Estonia              | `eu_vat`     | European VAT Number                                                                                     |
	// | EU                   | `eu_oss_vat` | European One Stop Shop VAT Number for non-Union scheme                                                  |
	// | Finland              | `eu_vat`     | European VAT Number                                                                                     |
	// | France               | `eu_vat`     | European VAT Number                                                                                     |
	// | Georgia              | `ge_vat`     | Georgian VAT                                                                                            |
	// | Germany              | `eu_vat`     | European VAT Number                                                                                     |
	// | Greece               | `eu_vat`     | European VAT Number                                                                                     |
	// | Hong Kong            | `hk_br`      | Hong Kong BR Number                                                                                     |
	// | Hungary              | `eu_vat`     | European VAT Number                                                                                     |
	// | Hungary              | `hu_tin`     | Hungary Tax Number (adószám)                                                                            |
	// | Iceland              | `is_vat`     | Icelandic VAT                                                                                           |
	// | India                | `in_gst`     | Indian GST Number                                                                                       |
	// | Indonesia            | `id_npwp`    | Indonesian NPWP Number                                                                                  |
	// | Ireland              | `eu_vat`     | European VAT Number                                                                                     |
	// | Israel               | `il_vat`     | Israel VAT                                                                                              |
	// | Italy                | `eu_vat`     | European VAT Number                                                                                     |
	// | Japan                | `jp_cn`      | Japanese Corporate Number (_Hōjin Bangō_)                                                               |
	// | Japan                | `jp_rn`      | Japanese Registered Foreign Businesses' Registration Number (_Tōroku Kokugai Jigyōsha no Tōroku Bangō_) |
	// | Japan                | `jp_trn`     | Japanese Tax Registration Number (_Tōroku Bangō_)                                                       |
	// | Kazakhstan           | `kz_bin`     | Kazakhstani Business Identification Number                                                              |
	// | Kenya                | `ke_pin`     | Kenya Revenue Authority Personal Identification Number                                                  |
	// | Latvia               | `eu_vat`     | European VAT Number                                                                                     |
	// | Liechtenstein        | `li_uid`     | Liechtensteinian UID Number                                                                             |
	// | Lithuania            | `eu_vat`     | European VAT Number                                                                                     |
	// | Luxembourg           | `eu_vat`     | European VAT Number                                                                                     |
	// | Malaysia             | `my_frp`     | Malaysian FRP Number                                                                                    |
	// | Malaysia             | `my_itn`     | Malaysian ITN                                                                                           |
	// | Malaysia             | `my_sst`     | Malaysian SST Number                                                                                    |
	// | Malta                | `eu_vat `    | European VAT Number                                                                                     |
	// | Mexico               | `mx_rfc`     | Mexican RFC Number                                                                                      |
	// | Netherlands          | `eu_vat`     | European VAT Number                                                                                     |
	// | New Zealand          | `nz_gst`     | New Zealand GST Number                                                                                  |
	// | Nigeria              | `ng_tin`     | Nigerian Tax Identification Number                                                                      |
	// | Norway               | `no_vat`     | Norwegian VAT Number                                                                                    |
	// | Norway               | `no_voec`    | Norwegian VAT on e-commerce Number                                                                      |
	// | Oman                 | `om_vat`     | Omani VAT Number                                                                                        |
	// | Peru                 | `pe_ruc`     | Peruvian RUC Number                                                                                     |
	// | Philippines          | `ph_tin `    | Philippines Tax Identification Number                                                                   |
	// | Poland               | `eu_vat`     | European VAT Number                                                                                     |
	// | Portugal             | `eu_vat`     | European VAT Number                                                                                     |
	// | Romania              | `eu_vat`     | European VAT Number                                                                                     |
	// | Romania              | `ro_tin`     | Romanian Tax ID Number                                                                                  |
	// | Russia               | `ru_inn`     | Russian INN                                                                                             |
	// | Russia               | `ru_kpp`     | Russian KPP                                                                                             |
	// | Saudi Arabia         | `sa_vat`     | Saudi Arabia VAT                                                                                        |
	// | Serbia               | `rs_pib`     | Serbian PIB Number                                                                                      |
	// | Singapore            | `sg_gst`     | Singaporean GST                                                                                         |
	// | Singapore            | `sg_uen`     | Singaporean UEN                                                                                         |
	// | Slovakia             | `eu_vat`     | European VAT Number                                                                                     |
	// | Slovenia             | `eu_vat`     | European VAT Number                                                                                     |
	// | Slovenia             | `si_tin`     | Slovenia Tax Number (davčna številka)                                                                   |
	// | South Africa         | `za_vat`     | South African VAT Number                                                                                |
	// | South Korea          | `kr_brn`     | Korean BRN                                                                                              |
	// | Spain                | `es_cif`     | Spanish NIF Number (previously Spanish CIF Number)                                                      |
	// | Spain                | `eu_vat`     | European VAT Number                                                                                     |
	// | Sweden               | `eu_vat`     | European VAT Number                                                                                     |
	// | Switzerland          | `ch_vat`     | Switzerland VAT Number                                                                                  |
	// | Taiwan               | `tw_vat`     | Taiwanese VAT                                                                                           |
	// | Thailand             | `th_vat`     | Thai VAT                                                                                                |
	// | Turkey               | `tr_tin`     | Turkish Tax Identification Number                                                                       |
	// | Ukraine              | `ua_vat`     | Ukrainian VAT                                                                                           |
	// | United Arab Emirates | `ae_trn`     | United Arab Emirates TRN                                                                                |
	// | United Kingdom       | `eu_vat`     | Northern Ireland VAT Number                                                                             |
	// | United Kingdom       | `gb_vat`     | United Kingdom VAT Number                                                                               |
	// | United States        | `us_ein`     | United States EIN                                                                                       |
	// | Uruguay              | `uy_ruc`     | Uruguayan RUC Number                                                                                    |
	// | Venezuela            | `ve_rif`     | Venezuelan RIF Number                                                                                   |
	// | Vietnam              | `vn_tin`     | Vietnamese Tax ID Number                                                                                |
	CustomerTaxID CustomerTaxIDModel `json:"customer_tax_id,required,nullable"`
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
	InvoicePdf    string                    `json:"invoice_pdf,required,nullable"`
	InvoiceSource InvoiceModelInvoiceSource `json:"invoice_source,required"`
	// If the invoice failed to issue, this will be the last time it failed to issue
	// (even if it is now in a different state.)
	IssueFailedAt time.Time `json:"issue_failed_at,required,nullable" format:"date-time"`
	// If the invoice has been issued, this will be the time it transitioned to
	// `issued` (even if it is now in a different state.)
	IssuedAt time.Time `json:"issued_at,required,nullable" format:"date-time"`
	// The breakdown of prices in this invoice.
	LineItems     []InvoiceLineItemModel `json:"line_items,required"`
	Maximum       MaximumModel           `json:"maximum,required,nullable"`
	MaximumAmount string                 `json:"maximum_amount,required,nullable"`
	// Free-form text which is available on the invoice PDF and the Orb invoice portal.
	Memo string `json:"memo,required,nullable"`
	// User specified key-value pairs for the resource. If not present, this defaults
	// to an empty dictionary. Individual keys can be removed by setting the value to
	// `null`, and the entire metadata mapping can be cleared by setting `metadata` to
	// `null`.
	Metadata      map[string]string `json:"metadata,required"`
	Minimum       MinimumModel      `json:"minimum,required,nullable"`
	MinimumAmount string            `json:"minimum_amount,required,nullable"`
	// If the invoice has a status of `paid`, this gives a timestamp when the invoice
	// was paid.
	PaidAt time.Time `json:"paid_at,required,nullable" format:"date-time"`
	// A list of payment attempts associated with the invoice
	PaymentAttempts []PaymentAttemptModel `json:"payment_attempts,required"`
	// If payment was attempted on this invoice but failed, this will be the time of
	// the most recent attempt.
	PaymentFailedAt time.Time `json:"payment_failed_at,required,nullable" format:"date-time"`
	// If payment was attempted on this invoice, this will be the start time of the
	// most recent attempt. This field is especially useful for delayed-notification
	// payment mechanisms (like bank transfers), where payment can take 3 days or more.
	PaymentStartedAt time.Time `json:"payment_started_at,required,nullable" format:"date-time"`
	// If the invoice is in draft, this timestamp will reflect when the invoice is
	// scheduled to be issued.
	ScheduledIssueAt time.Time                 `json:"scheduled_issue_at,required,nullable" format:"date-time"`
	ShippingAddress  AddressModel              `json:"shipping_address,required,nullable"`
	Status           InvoiceModelStatus        `json:"status,required"`
	Subscription     SubscriptionMinifiedModel `json:"subscription,required,nullable"`
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
	WillAutoIssue bool             `json:"will_auto_issue,required"`
	JSON          invoiceModelJSON `json:"-"`
}

// invoiceModelJSON contains the JSON metadata for the struct [InvoiceModel]
type invoiceModelJSON struct {
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

func (r *InvoiceModel) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r invoiceModelJSON) RawJSON() string {
	return r.raw
}

type InvoiceModelInvoiceSource string

const (
	InvoiceModelInvoiceSourceSubscription InvoiceModelInvoiceSource = "subscription"
	InvoiceModelInvoiceSourcePartial      InvoiceModelInvoiceSource = "partial"
	InvoiceModelInvoiceSourceOneOff       InvoiceModelInvoiceSource = "one_off"
)

func (r InvoiceModelInvoiceSource) IsKnown() bool {
	switch r {
	case InvoiceModelInvoiceSourceSubscription, InvoiceModelInvoiceSourcePartial, InvoiceModelInvoiceSourceOneOff:
		return true
	}
	return false
}

type InvoiceModelStatus string

const (
	InvoiceModelStatusIssued InvoiceModelStatus = "issued"
	InvoiceModelStatusPaid   InvoiceModelStatus = "paid"
	InvoiceModelStatusSynced InvoiceModelStatus = "synced"
	InvoiceModelStatusVoid   InvoiceModelStatus = "void"
	InvoiceModelStatusDraft  InvoiceModelStatus = "draft"
)

func (r InvoiceModelStatus) IsKnown() bool {
	switch r {
	case InvoiceModelStatusIssued, InvoiceModelStatusPaid, InvoiceModelStatusSynced, InvoiceModelStatusVoid, InvoiceModelStatusDraft:
		return true
	}
	return false
}

type ItemExternalConnectionModel struct {
	ExternalConnectionName ItemExternalConnectionModelExternalConnectionName `json:"external_connection_name,required"`
	ExternalEntityID       string                                            `json:"external_entity_id,required"`
	JSON                   itemExternalConnectionModelJSON                   `json:"-"`
}

// itemExternalConnectionModelJSON contains the JSON metadata for the struct
// [ItemExternalConnectionModel]
type itemExternalConnectionModelJSON struct {
	ExternalConnectionName apijson.Field
	ExternalEntityID       apijson.Field
	raw                    string
	ExtraFields            map[string]apijson.Field
}

func (r *ItemExternalConnectionModel) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r itemExternalConnectionModelJSON) RawJSON() string {
	return r.raw
}

type ItemExternalConnectionModelExternalConnectionName string

const (
	ItemExternalConnectionModelExternalConnectionNameStripe     ItemExternalConnectionModelExternalConnectionName = "stripe"
	ItemExternalConnectionModelExternalConnectionNameQuickbooks ItemExternalConnectionModelExternalConnectionName = "quickbooks"
	ItemExternalConnectionModelExternalConnectionNameBillCom    ItemExternalConnectionModelExternalConnectionName = "bill.com"
	ItemExternalConnectionModelExternalConnectionNameNetsuite   ItemExternalConnectionModelExternalConnectionName = "netsuite"
	ItemExternalConnectionModelExternalConnectionNameTaxjar     ItemExternalConnectionModelExternalConnectionName = "taxjar"
	ItemExternalConnectionModelExternalConnectionNameAvalara    ItemExternalConnectionModelExternalConnectionName = "avalara"
	ItemExternalConnectionModelExternalConnectionNameAnrok      ItemExternalConnectionModelExternalConnectionName = "anrok"
)

func (r ItemExternalConnectionModelExternalConnectionName) IsKnown() bool {
	switch r {
	case ItemExternalConnectionModelExternalConnectionNameStripe, ItemExternalConnectionModelExternalConnectionNameQuickbooks, ItemExternalConnectionModelExternalConnectionNameBillCom, ItemExternalConnectionModelExternalConnectionNameNetsuite, ItemExternalConnectionModelExternalConnectionNameTaxjar, ItemExternalConnectionModelExternalConnectionNameAvalara, ItemExternalConnectionModelExternalConnectionNameAnrok:
		return true
	}
	return false
}

type ItemExternalConnectionModelParam struct {
	ExternalConnectionName param.Field[ItemExternalConnectionModelExternalConnectionName] `json:"external_connection_name,required"`
	ExternalEntityID       param.Field[string]                                            `json:"external_entity_id,required"`
}

func (r ItemExternalConnectionModelParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// The Item resource represents a sellable product or good. Items are associated
// with all line items, billable metrics, and prices and are used for defining
// external sync behavior for invoices and tax calculation purposes.
type ItemModel struct {
	ID                  string                        `json:"id,required"`
	CreatedAt           time.Time                     `json:"created_at,required" format:"date-time"`
	ExternalConnections []ItemExternalConnectionModel `json:"external_connections,required"`
	Name                string                        `json:"name,required"`
	JSON                itemModelJSON                 `json:"-"`
}

// itemModelJSON contains the JSON metadata for the struct [ItemModel]
type itemModelJSON struct {
	ID                  apijson.Field
	CreatedAt           apijson.Field
	ExternalConnections apijson.Field
	Name                apijson.Field
	raw                 string
	ExtraFields         map[string]apijson.Field
}

func (r *ItemModel) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r itemModelJSON) RawJSON() string {
	return r.raw
}

type ItemSlimModel struct {
	ID   string            `json:"id,required"`
	Name string            `json:"name,required"`
	JSON itemSlimModelJSON `json:"-"`
}

// itemSlimModelJSON contains the JSON metadata for the struct [ItemSlimModel]
type itemSlimModelJSON struct {
	ID          apijson.Field
	Name        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ItemSlimModel) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r itemSlimModelJSON) RawJSON() string {
	return r.raw
}

type MatrixConfigModel struct {
	// Default per unit rate for any usage not bucketed into a specified matrix_value
	DefaultUnitAmount string `json:"default_unit_amount,required"`
	// One or two event property values to evaluate matrix groups by
	Dimensions []string `json:"dimensions,required"`
	// Matrix values for specified matrix grouping keys
	MatrixValues []MatrixValueModel    `json:"matrix_values,required"`
	JSON         matrixConfigModelJSON `json:"-"`
}

// matrixConfigModelJSON contains the JSON metadata for the struct
// [MatrixConfigModel]
type matrixConfigModelJSON struct {
	DefaultUnitAmount apijson.Field
	Dimensions        apijson.Field
	MatrixValues      apijson.Field
	raw               string
	ExtraFields       map[string]apijson.Field
}

func (r *MatrixConfigModel) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r matrixConfigModelJSON) RawJSON() string {
	return r.raw
}

type MatrixConfigModelParam struct {
	// Default per unit rate for any usage not bucketed into a specified matrix_value
	DefaultUnitAmount param.Field[string] `json:"default_unit_amount,required"`
	// One or two event property values to evaluate matrix groups by
	Dimensions param.Field[[]string] `json:"dimensions,required"`
	// Matrix values for specified matrix grouping keys
	MatrixValues param.Field[[]MatrixValueModelParam] `json:"matrix_values,required"`
}

func (r MatrixConfigModelParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type MatrixValueModel struct {
	// One or two matrix keys to filter usage to this Matrix value by. For example,
	// ["region", "tier"] could be used to filter cloud usage by a cloud region and an
	// instance tier.
	DimensionValues []string `json:"dimension_values,required"`
	// Unit price for the specified dimension_values
	UnitAmount string               `json:"unit_amount,required"`
	JSON       matrixValueModelJSON `json:"-"`
}

// matrixValueModelJSON contains the JSON metadata for the struct
// [MatrixValueModel]
type matrixValueModelJSON struct {
	DimensionValues apijson.Field
	UnitAmount      apijson.Field
	raw             string
	ExtraFields     map[string]apijson.Field
}

func (r *MatrixValueModel) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r matrixValueModelJSON) RawJSON() string {
	return r.raw
}

type MatrixValueModelParam struct {
	// One or two matrix keys to filter usage to this Matrix value by. For example,
	// ["region", "tier"] could be used to filter cloud usage by a cloud region and an
	// instance tier.
	DimensionValues param.Field[[]string] `json:"dimension_values,required"`
	// Unit price for the specified dimension_values
	UnitAmount param.Field[string] `json:"unit_amount,required"`
}

func (r MatrixValueModelParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type MatrixWithAllocationConfigModel struct {
	// Allocation to be used to calculate the price
	Allocation float64 `json:"allocation,required"`
	// Default per unit rate for any usage not bucketed into a specified matrix_value
	DefaultUnitAmount string `json:"default_unit_amount,required"`
	// One or two event property values to evaluate matrix groups by
	Dimensions []string `json:"dimensions,required"`
	// Matrix values for specified matrix grouping keys
	MatrixValues []MatrixValueModel                  `json:"matrix_values,required"`
	JSON         matrixWithAllocationConfigModelJSON `json:"-"`
}

// matrixWithAllocationConfigModelJSON contains the JSON metadata for the struct
// [MatrixWithAllocationConfigModel]
type matrixWithAllocationConfigModelJSON struct {
	Allocation        apijson.Field
	DefaultUnitAmount apijson.Field
	Dimensions        apijson.Field
	MatrixValues      apijson.Field
	raw               string
	ExtraFields       map[string]apijson.Field
}

func (r *MatrixWithAllocationConfigModel) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r matrixWithAllocationConfigModelJSON) RawJSON() string {
	return r.raw
}

type MatrixWithAllocationConfigModelParam struct {
	// Allocation to be used to calculate the price
	Allocation param.Field[float64] `json:"allocation,required"`
	// Default per unit rate for any usage not bucketed into a specified matrix_value
	DefaultUnitAmount param.Field[string] `json:"default_unit_amount,required"`
	// One or two event property values to evaluate matrix groups by
	Dimensions param.Field[[]string] `json:"dimensions,required"`
	// Matrix values for specified matrix grouping keys
	MatrixValues param.Field[[]MatrixValueModelParam] `json:"matrix_values,required"`
}

func (r MatrixWithAllocationConfigModelParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type MaximumIntervalModel struct {
	// The price ids that this maximum interval applies to.
	AppliesToPriceIDs []string `json:"applies_to_price_ids,required"`
	// The price interval ids that this maximum interval applies to.
	AppliesToPriceIntervalIDs []string `json:"applies_to_price_interval_ids,required"`
	// The end date of the maximum interval.
	EndDate time.Time `json:"end_date,required,nullable" format:"date-time"`
	// The maximum amount to charge in a given billing period for the price intervals
	// this transform applies to.
	MaximumAmount string `json:"maximum_amount,required"`
	// The start date of the maximum interval.
	StartDate time.Time                `json:"start_date,required" format:"date-time"`
	JSON      maximumIntervalModelJSON `json:"-"`
}

// maximumIntervalModelJSON contains the JSON metadata for the struct
// [MaximumIntervalModel]
type maximumIntervalModelJSON struct {
	AppliesToPriceIDs         apijson.Field
	AppliesToPriceIntervalIDs apijson.Field
	EndDate                   apijson.Field
	MaximumAmount             apijson.Field
	StartDate                 apijson.Field
	raw                       string
	ExtraFields               map[string]apijson.Field
}

func (r *MaximumIntervalModel) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r maximumIntervalModelJSON) RawJSON() string {
	return r.raw
}

type MaximumModel struct {
	// List of price_ids that this maximum amount applies to. For plan/plan phase
	// maximums, this can be a subset of prices.
	AppliesToPriceIDs []string `json:"applies_to_price_ids,required"`
	// Maximum amount applied
	MaximumAmount string           `json:"maximum_amount,required"`
	JSON          maximumModelJSON `json:"-"`
}

// maximumModelJSON contains the JSON metadata for the struct [MaximumModel]
type maximumModelJSON struct {
	AppliesToPriceIDs apijson.Field
	MaximumAmount     apijson.Field
	raw               string
	ExtraFields       map[string]apijson.Field
}

func (r *MaximumModel) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r maximumModelJSON) RawJSON() string {
	return r.raw
}

type MinimumIntervalModel struct {
	// The price ids that this minimum interval applies to.
	AppliesToPriceIDs []string `json:"applies_to_price_ids,required"`
	// The price interval ids that this minimum interval applies to.
	AppliesToPriceIntervalIDs []string `json:"applies_to_price_interval_ids,required"`
	// The end date of the minimum interval.
	EndDate time.Time `json:"end_date,required,nullable" format:"date-time"`
	// The minimum amount to charge in a given billing period for the price intervals
	// this minimum applies to.
	MinimumAmount string `json:"minimum_amount,required"`
	// The start date of the minimum interval.
	StartDate time.Time                `json:"start_date,required" format:"date-time"`
	JSON      minimumIntervalModelJSON `json:"-"`
}

// minimumIntervalModelJSON contains the JSON metadata for the struct
// [MinimumIntervalModel]
type minimumIntervalModelJSON struct {
	AppliesToPriceIDs         apijson.Field
	AppliesToPriceIntervalIDs apijson.Field
	EndDate                   apijson.Field
	MinimumAmount             apijson.Field
	StartDate                 apijson.Field
	raw                       string
	ExtraFields               map[string]apijson.Field
}

func (r *MinimumIntervalModel) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r minimumIntervalModelJSON) RawJSON() string {
	return r.raw
}

type MinimumModel struct {
	// List of price_ids that this minimum amount applies to. For plan/plan phase
	// minimums, this can be a subset of prices.
	AppliesToPriceIDs []string `json:"applies_to_price_ids,required"`
	// Minimum amount applied
	MinimumAmount string           `json:"minimum_amount,required"`
	JSON          minimumModelJSON `json:"-"`
}

// minimumModelJSON contains the JSON metadata for the struct [MinimumModel]
type minimumModelJSON struct {
	AppliesToPriceIDs apijson.Field
	MinimumAmount     apijson.Field
	raw               string
	ExtraFields       map[string]apijson.Field
}

func (r *MinimumModel) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r minimumModelJSON) RawJSON() string {
	return r.raw
}

type MutatedSubscriptionModel struct {
	ID string `json:"id,required"`
	// The current plan phase that is active, only if the subscription's plan has
	// phases.
	ActivePlanPhaseOrder int64 `json:"active_plan_phase_order,required,nullable"`
	// The adjustment intervals for this subscription.
	AdjustmentIntervals []AdjustmentIntervalModel `json:"adjustment_intervals,required"`
	// Determines whether issued invoices for this subscription will automatically be
	// charged with the saved payment method on the due date. This property defaults to
	// the plan's behavior. If null, defaults to the customer's setting.
	AutoCollection                  bool                                 `json:"auto_collection,required,nullable"`
	BillingCycleAnchorConfiguration BillingCycleAnchorConfigurationModel `json:"billing_cycle_anchor_configuration,required"`
	// The day of the month on which the billing cycle is anchored. If the maximum
	// number of days in a month is greater than this value, the last day of the month
	// is the billing cycle day (e.g. billing_cycle_day=31 for April means the billing
	// period begins on the 30th.
	BillingCycleDay int64     `json:"billing_cycle_day,required"`
	CreatedAt       time.Time `json:"created_at,required" format:"date-time"`
	// The end of the current billing period. This is an exclusive timestamp, such that
	// the instant returned is not part of the billing period. Set to null for
	// subscriptions that are not currently active.
	CurrentBillingPeriodEndDate time.Time `json:"current_billing_period_end_date,required,nullable" format:"date-time"`
	// The start date of the current billing period. This is an inclusive timestamp;
	// the instant returned is exactly the beginning of the billing period. Set to null
	// if the subscription is not currently active.
	CurrentBillingPeriodStartDate time.Time `json:"current_billing_period_start_date,required,nullable" format:"date-time"`
	// A customer is a buyer of your products, and the other party to the billing
	// relationship.
	//
	// In Orb, customers are assigned system generated identifiers automatically, but
	// it's often desirable to have these match existing identifiers in your system. To
	// avoid having to denormalize Orb ID information, you can pass in an
	// `external_customer_id` with your own identifier. See
	// [Customer ID Aliases](/events-and-metrics/customer-aliases) for further
	// information about how these aliases work in Orb.
	//
	// In addition to having an identifier in your system, a customer may exist in a
	// payment provider solution like Stripe. Use the `payment_provider_id` and the
	// `payment_provider` enum field to express this mapping.
	//
	// A customer also has a timezone (from the standard
	// [IANA timezone database](https://www.iana.org/time-zones)), which defaults to
	// your account's timezone. See [Timezone localization](/essentials/timezones) for
	// information on what this timezone parameter influences within Orb.
	Customer CustomerModel `json:"customer,required"`
	// Determines the default memo on this subscriptions' invoices. Note that if this
	// is not provided, it is determined by the plan configuration.
	DefaultInvoiceMemo string `json:"default_invoice_memo,required,nullable"`
	// The discount intervals for this subscription.
	DiscountIntervals []MutatedSubscriptionModelDiscountInterval `json:"discount_intervals,required"`
	// The date Orb stops billing for this subscription.
	EndDate                  time.Time                            `json:"end_date,required,nullable" format:"date-time"`
	FixedFeeQuantitySchedule []FixedFeeQuantityScheduleEntryModel `json:"fixed_fee_quantity_schedule,required"`
	InvoicingThreshold       string                               `json:"invoicing_threshold,required,nullable"`
	// The maximum intervals for this subscription.
	MaximumIntervals []MaximumIntervalModel `json:"maximum_intervals,required"`
	// User specified key-value pairs for the resource. If not present, this defaults
	// to an empty dictionary. Individual keys can be removed by setting the value to
	// `null`, and the entire metadata mapping can be cleared by setting `metadata` to
	// `null`.
	Metadata map[string]string `json:"metadata,required"`
	// The minimum intervals for this subscription.
	MinimumIntervals []MinimumIntervalModel `json:"minimum_intervals,required"`
	// Determines the difference between the invoice issue date for subscription
	// invoices as the date that they are due. A value of `0` here represents that the
	// invoice is due on issue, whereas a value of `30` represents that the customer
	// has a month to pay the invoice.
	NetTerms int64 `json:"net_terms,required"`
	// The [Plan](/core-concepts#plan-and-price) resource represents a plan that can be
	// subscribed to by a customer. Plans define the billing behavior of the
	// subscription. You can see more about how to configure prices in the
	// [Price resource](/reference/price).
	Plan PlanModel `json:"plan,required"`
	// The price intervals for this subscription.
	PriceIntervals []PriceIntervalModel  `json:"price_intervals,required"`
	RedeemedCoupon CouponRedemptionModel `json:"redeemed_coupon,required,nullable"`
	// The date Orb starts billing for this subscription.
	StartDate time.Time                      `json:"start_date,required" format:"date-time"`
	Status    MutatedSubscriptionModelStatus `json:"status,required"`
	TrialInfo SubscriptionTrialInfoModel     `json:"trial_info,required"`
	JSON      mutatedSubscriptionModelJSON   `json:"-"`
}

// mutatedSubscriptionModelJSON contains the JSON metadata for the struct
// [MutatedSubscriptionModel]
type mutatedSubscriptionModelJSON struct {
	ID                              apijson.Field
	ActivePlanPhaseOrder            apijson.Field
	AdjustmentIntervals             apijson.Field
	AutoCollection                  apijson.Field
	BillingCycleAnchorConfiguration apijson.Field
	BillingCycleDay                 apijson.Field
	CreatedAt                       apijson.Field
	CurrentBillingPeriodEndDate     apijson.Field
	CurrentBillingPeriodStartDate   apijson.Field
	Customer                        apijson.Field
	DefaultInvoiceMemo              apijson.Field
	DiscountIntervals               apijson.Field
	EndDate                         apijson.Field
	FixedFeeQuantitySchedule        apijson.Field
	InvoicingThreshold              apijson.Field
	MaximumIntervals                apijson.Field
	Metadata                        apijson.Field
	MinimumIntervals                apijson.Field
	NetTerms                        apijson.Field
	Plan                            apijson.Field
	PriceIntervals                  apijson.Field
	RedeemedCoupon                  apijson.Field
	StartDate                       apijson.Field
	Status                          apijson.Field
	TrialInfo                       apijson.Field
	raw                             string
	ExtraFields                     map[string]apijson.Field
}

func (r *MutatedSubscriptionModel) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r mutatedSubscriptionModelJSON) RawJSON() string {
	return r.raw
}

type MutatedSubscriptionModelDiscountInterval struct {
	// This field can have the runtime type of [[]string].
	AppliesToPriceIDs interface{} `json:"applies_to_price_ids,required"`
	// This field can have the runtime type of [[]string].
	AppliesToPriceIntervalIDs interface{}                                           `json:"applies_to_price_interval_ids,required"`
	DiscountType              MutatedSubscriptionModelDiscountIntervalsDiscountType `json:"discount_type,required"`
	// The end date of the discount interval.
	EndDate time.Time `json:"end_date,required,nullable" format:"date-time"`
	// The start date of the discount interval.
	StartDate time.Time `json:"start_date,required" format:"date-time"`
	// Only available if discount_type is `amount`.
	AmountDiscount string `json:"amount_discount"`
	// Only available if discount_type is `percentage`.This is a number between 0
	// and 1.
	PercentageDiscount float64 `json:"percentage_discount"`
	// Only available if discount_type is `usage`. Number of usage units that this
	// discount is for
	UsageDiscount float64                                      `json:"usage_discount"`
	JSON          mutatedSubscriptionModelDiscountIntervalJSON `json:"-"`
	union         MutatedSubscriptionModelDiscountIntervalsUnion
}

// mutatedSubscriptionModelDiscountIntervalJSON contains the JSON metadata for the
// struct [MutatedSubscriptionModelDiscountInterval]
type mutatedSubscriptionModelDiscountIntervalJSON struct {
	AppliesToPriceIDs         apijson.Field
	AppliesToPriceIntervalIDs apijson.Field
	DiscountType              apijson.Field
	EndDate                   apijson.Field
	StartDate                 apijson.Field
	AmountDiscount            apijson.Field
	PercentageDiscount        apijson.Field
	UsageDiscount             apijson.Field
	raw                       string
	ExtraFields               map[string]apijson.Field
}

func (r mutatedSubscriptionModelDiscountIntervalJSON) RawJSON() string {
	return r.raw
}

func (r *MutatedSubscriptionModelDiscountInterval) UnmarshalJSON(data []byte) (err error) {
	*r = MutatedSubscriptionModelDiscountInterval{}
	err = apijson.UnmarshalRoot(data, &r.union)
	if err != nil {
		return err
	}
	return apijson.Port(r.union, &r)
}

// AsUnion returns a [MutatedSubscriptionModelDiscountIntervalsUnion] interface
// which you can cast to the specific types for more type safety.
//
// Possible runtime types of the union are [shared.AmountDiscountIntervalModel],
// [shared.PercentageDiscountIntervalModel], [shared.UsageDiscountIntervalModel].
func (r MutatedSubscriptionModelDiscountInterval) AsUnion() MutatedSubscriptionModelDiscountIntervalsUnion {
	return r.union
}

// Union satisfied by [shared.AmountDiscountIntervalModel],
// [shared.PercentageDiscountIntervalModel] or [shared.UsageDiscountIntervalModel].
type MutatedSubscriptionModelDiscountIntervalsUnion interface {
	ImplementsMutatedSubscriptionModelDiscountInterval()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*MutatedSubscriptionModelDiscountIntervalsUnion)(nil)).Elem(),
		"discount_type",
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(AmountDiscountIntervalModel{}),
			DiscriminatorValue: "amount",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(PercentageDiscountIntervalModel{}),
			DiscriminatorValue: "percentage",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(UsageDiscountIntervalModel{}),
			DiscriminatorValue: "usage",
		},
	)
}

type MutatedSubscriptionModelDiscountIntervalsDiscountType string

const (
	MutatedSubscriptionModelDiscountIntervalsDiscountTypeAmount     MutatedSubscriptionModelDiscountIntervalsDiscountType = "amount"
	MutatedSubscriptionModelDiscountIntervalsDiscountTypePercentage MutatedSubscriptionModelDiscountIntervalsDiscountType = "percentage"
	MutatedSubscriptionModelDiscountIntervalsDiscountTypeUsage      MutatedSubscriptionModelDiscountIntervalsDiscountType = "usage"
)

func (r MutatedSubscriptionModelDiscountIntervalsDiscountType) IsKnown() bool {
	switch r {
	case MutatedSubscriptionModelDiscountIntervalsDiscountTypeAmount, MutatedSubscriptionModelDiscountIntervalsDiscountTypePercentage, MutatedSubscriptionModelDiscountIntervalsDiscountTypeUsage:
		return true
	}
	return false
}

type MutatedSubscriptionModelStatus string

const (
	MutatedSubscriptionModelStatusActive   MutatedSubscriptionModelStatus = "active"
	MutatedSubscriptionModelStatusEnded    MutatedSubscriptionModelStatus = "ended"
	MutatedSubscriptionModelStatusUpcoming MutatedSubscriptionModelStatus = "upcoming"
)

func (r MutatedSubscriptionModelStatus) IsKnown() bool {
	switch r {
	case MutatedSubscriptionModelStatusActive, MutatedSubscriptionModelStatusEnded, MutatedSubscriptionModelStatusUpcoming:
		return true
	}
	return false
}

type NewAccountingSyncConfigurationModelParam struct {
	AccountingProviders param.Field[[]NewAccountingSyncConfigurationModelAccountingProviderParam] `json:"accounting_providers"`
	Excluded            param.Field[bool]                                                         `json:"excluded"`
}

func (r NewAccountingSyncConfigurationModelParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type NewAccountingSyncConfigurationModelAccountingProviderParam struct {
	ExternalProviderID param.Field[string] `json:"external_provider_id,required"`
	ProviderType       param.Field[string] `json:"provider_type,required"`
}

func (r NewAccountingSyncConfigurationModelAccountingProviderParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type NewAdjustmentModelParam struct {
	AdjustmentType    param.Field[NewAdjustmentModelAdjustmentType] `json:"adjustment_type,required"`
	AppliesToPriceIDs param.Field[interface{}]                      `json:"applies_to_price_ids,required"`
	AmountDiscount    param.Field[string]                           `json:"amount_discount"`
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

func (r NewAdjustmentModelParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r NewAdjustmentModelParam) implementsNewAdjustmentModelUnionParam() {}

// Satisfied by [shared.NewAdjustmentModelNewPercentageDiscountParam],
// [shared.NewAdjustmentModelNewUsageDiscountParam],
// [shared.NewAdjustmentModelNewAmountDiscountParam],
// [shared.NewAdjustmentModelNewMinimumParam],
// [shared.NewAdjustmentModelNewMaximumParam], [NewAdjustmentModelParam].
type NewAdjustmentModelUnionParam interface {
	implementsNewAdjustmentModelUnionParam()
}

type NewAdjustmentModelNewPercentageDiscountParam struct {
	AdjustmentType param.Field[NewAdjustmentModelNewPercentageDiscountAdjustmentType] `json:"adjustment_type,required"`
	// The set of price IDs to which this adjustment applies.
	AppliesToPriceIDs  param.Field[[]string] `json:"applies_to_price_ids,required"`
	PercentageDiscount param.Field[float64]  `json:"percentage_discount,required"`
	// When false, this adjustment will be applied to a single price. Otherwise, it
	// will be applied at the invoice level, possibly to multiple prices.
	IsInvoiceLevel param.Field[bool] `json:"is_invoice_level"`
}

func (r NewAdjustmentModelNewPercentageDiscountParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r NewAdjustmentModelNewPercentageDiscountParam) implementsNewAdjustmentModelUnionParam() {}

type NewAdjustmentModelNewPercentageDiscountAdjustmentType string

const (
	NewAdjustmentModelNewPercentageDiscountAdjustmentTypePercentageDiscount NewAdjustmentModelNewPercentageDiscountAdjustmentType = "percentage_discount"
)

func (r NewAdjustmentModelNewPercentageDiscountAdjustmentType) IsKnown() bool {
	switch r {
	case NewAdjustmentModelNewPercentageDiscountAdjustmentTypePercentageDiscount:
		return true
	}
	return false
}

type NewAdjustmentModelNewUsageDiscountParam struct {
	AdjustmentType param.Field[NewAdjustmentModelNewUsageDiscountAdjustmentType] `json:"adjustment_type,required"`
	// The set of price IDs to which this adjustment applies.
	AppliesToPriceIDs param.Field[[]string] `json:"applies_to_price_ids,required"`
	UsageDiscount     param.Field[float64]  `json:"usage_discount,required"`
	// When false, this adjustment will be applied to a single price. Otherwise, it
	// will be applied at the invoice level, possibly to multiple prices.
	IsInvoiceLevel param.Field[bool] `json:"is_invoice_level"`
}

func (r NewAdjustmentModelNewUsageDiscountParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r NewAdjustmentModelNewUsageDiscountParam) implementsNewAdjustmentModelUnionParam() {}

type NewAdjustmentModelNewUsageDiscountAdjustmentType string

const (
	NewAdjustmentModelNewUsageDiscountAdjustmentTypeUsageDiscount NewAdjustmentModelNewUsageDiscountAdjustmentType = "usage_discount"
)

func (r NewAdjustmentModelNewUsageDiscountAdjustmentType) IsKnown() bool {
	switch r {
	case NewAdjustmentModelNewUsageDiscountAdjustmentTypeUsageDiscount:
		return true
	}
	return false
}

type NewAdjustmentModelNewAmountDiscountParam struct {
	AdjustmentType param.Field[NewAdjustmentModelNewAmountDiscountAdjustmentType] `json:"adjustment_type,required"`
	AmountDiscount param.Field[string]                                            `json:"amount_discount,required"`
	// The set of price IDs to which this adjustment applies.
	AppliesToPriceIDs param.Field[[]string] `json:"applies_to_price_ids,required"`
	// When false, this adjustment will be applied to a single price. Otherwise, it
	// will be applied at the invoice level, possibly to multiple prices.
	IsInvoiceLevel param.Field[bool] `json:"is_invoice_level"`
}

func (r NewAdjustmentModelNewAmountDiscountParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r NewAdjustmentModelNewAmountDiscountParam) implementsNewAdjustmentModelUnionParam() {}

type NewAdjustmentModelNewAmountDiscountAdjustmentType string

const (
	NewAdjustmentModelNewAmountDiscountAdjustmentTypeAmountDiscount NewAdjustmentModelNewAmountDiscountAdjustmentType = "amount_discount"
)

func (r NewAdjustmentModelNewAmountDiscountAdjustmentType) IsKnown() bool {
	switch r {
	case NewAdjustmentModelNewAmountDiscountAdjustmentTypeAmountDiscount:
		return true
	}
	return false
}

type NewAdjustmentModelNewMinimumParam struct {
	AdjustmentType param.Field[NewAdjustmentModelNewMinimumAdjustmentType] `json:"adjustment_type,required"`
	// The set of price IDs to which this adjustment applies.
	AppliesToPriceIDs param.Field[[]string] `json:"applies_to_price_ids,required"`
	// The item ID that revenue from this minimum will be attributed to.
	ItemID        param.Field[string] `json:"item_id,required"`
	MinimumAmount param.Field[string] `json:"minimum_amount,required"`
	// When false, this adjustment will be applied to a single price. Otherwise, it
	// will be applied at the invoice level, possibly to multiple prices.
	IsInvoiceLevel param.Field[bool] `json:"is_invoice_level"`
}

func (r NewAdjustmentModelNewMinimumParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r NewAdjustmentModelNewMinimumParam) implementsNewAdjustmentModelUnionParam() {}

type NewAdjustmentModelNewMinimumAdjustmentType string

const (
	NewAdjustmentModelNewMinimumAdjustmentTypeMinimum NewAdjustmentModelNewMinimumAdjustmentType = "minimum"
)

func (r NewAdjustmentModelNewMinimumAdjustmentType) IsKnown() bool {
	switch r {
	case NewAdjustmentModelNewMinimumAdjustmentTypeMinimum:
		return true
	}
	return false
}

type NewAdjustmentModelNewMaximumParam struct {
	AdjustmentType param.Field[NewAdjustmentModelNewMaximumAdjustmentType] `json:"adjustment_type,required"`
	// The set of price IDs to which this adjustment applies.
	AppliesToPriceIDs param.Field[[]string] `json:"applies_to_price_ids,required"`
	MaximumAmount     param.Field[string]   `json:"maximum_amount,required"`
	// When false, this adjustment will be applied to a single price. Otherwise, it
	// will be applied at the invoice level, possibly to multiple prices.
	IsInvoiceLevel param.Field[bool] `json:"is_invoice_level"`
}

func (r NewAdjustmentModelNewMaximumParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r NewAdjustmentModelNewMaximumParam) implementsNewAdjustmentModelUnionParam() {}

type NewAdjustmentModelNewMaximumAdjustmentType string

const (
	NewAdjustmentModelNewMaximumAdjustmentTypeMaximum NewAdjustmentModelNewMaximumAdjustmentType = "maximum"
)

func (r NewAdjustmentModelNewMaximumAdjustmentType) IsKnown() bool {
	switch r {
	case NewAdjustmentModelNewMaximumAdjustmentTypeMaximum:
		return true
	}
	return false
}

type NewAdjustmentModelAdjustmentType string

const (
	NewAdjustmentModelAdjustmentTypePercentageDiscount NewAdjustmentModelAdjustmentType = "percentage_discount"
	NewAdjustmentModelAdjustmentTypeUsageDiscount      NewAdjustmentModelAdjustmentType = "usage_discount"
	NewAdjustmentModelAdjustmentTypeAmountDiscount     NewAdjustmentModelAdjustmentType = "amount_discount"
	NewAdjustmentModelAdjustmentTypeMinimum            NewAdjustmentModelAdjustmentType = "minimum"
	NewAdjustmentModelAdjustmentTypeMaximum            NewAdjustmentModelAdjustmentType = "maximum"
)

func (r NewAdjustmentModelAdjustmentType) IsKnown() bool {
	switch r {
	case NewAdjustmentModelAdjustmentTypePercentageDiscount, NewAdjustmentModelAdjustmentTypeUsageDiscount, NewAdjustmentModelAdjustmentTypeAmountDiscount, NewAdjustmentModelAdjustmentTypeMinimum, NewAdjustmentModelAdjustmentTypeMaximum:
		return true
	}
	return false
}

type NewAllocationPriceModelParam struct {
	// An amount of the currency to allocate to the customer at the specified cadence.
	Amount param.Field[string] `json:"amount,required"`
	// The cadence at which to allocate the amount to the customer.
	Cadence param.Field[NewAllocationPriceModelCadence] `json:"cadence,required"`
	// An ISO 4217 currency string or a custom pricing unit identifier in which to bill
	// this price.
	Currency param.Field[string] `json:"currency,required"`
	// Whether the allocated amount should expire at the end of the cadence or roll
	// over to the next period.
	ExpiresAtEndOfCadence param.Field[bool] `json:"expires_at_end_of_cadence,required"`
}

func (r NewAllocationPriceModelParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// The cadence at which to allocate the amount to the customer.
type NewAllocationPriceModelCadence string

const (
	NewAllocationPriceModelCadenceOneTime    NewAllocationPriceModelCadence = "one_time"
	NewAllocationPriceModelCadenceMonthly    NewAllocationPriceModelCadence = "monthly"
	NewAllocationPriceModelCadenceQuarterly  NewAllocationPriceModelCadence = "quarterly"
	NewAllocationPriceModelCadenceSemiAnnual NewAllocationPriceModelCadence = "semi_annual"
	NewAllocationPriceModelCadenceAnnual     NewAllocationPriceModelCadence = "annual"
	NewAllocationPriceModelCadenceCustom     NewAllocationPriceModelCadence = "custom"
)

func (r NewAllocationPriceModelCadence) IsKnown() bool {
	switch r {
	case NewAllocationPriceModelCadenceOneTime, NewAllocationPriceModelCadenceMonthly, NewAllocationPriceModelCadenceQuarterly, NewAllocationPriceModelCadenceSemiAnnual, NewAllocationPriceModelCadenceAnnual, NewAllocationPriceModelCadenceCustom:
		return true
	}
	return false
}

type NewBillingCycleConfigurationModelParam struct {
	// The duration of the billing period.
	Duration param.Field[int64] `json:"duration,required"`
	// The unit of billing period duration.
	DurationUnit param.Field[NewBillingCycleConfigurationModelDurationUnit] `json:"duration_unit,required"`
}

func (r NewBillingCycleConfigurationModelParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// The unit of billing period duration.
type NewBillingCycleConfigurationModelDurationUnit string

const (
	NewBillingCycleConfigurationModelDurationUnitDay   NewBillingCycleConfigurationModelDurationUnit = "day"
	NewBillingCycleConfigurationModelDurationUnitMonth NewBillingCycleConfigurationModelDurationUnit = "month"
)

func (r NewBillingCycleConfigurationModelDurationUnit) IsKnown() bool {
	switch r {
	case NewBillingCycleConfigurationModelDurationUnitDay, NewBillingCycleConfigurationModelDurationUnitMonth:
		return true
	}
	return false
}

type NewFloatingPriceModelParam struct {
	// The cadence to bill for this price on.
	Cadence param.Field[NewFloatingPriceModelCadence] `json:"cadence,required"`
	// An ISO 4217 currency string for which this price is billed in.
	Currency param.Field[string] `json:"currency,required"`
	// The id of the item the price will be associated with.
	ItemID    param.Field[string]                         `json:"item_id,required"`
	ModelType param.Field[NewFloatingPriceModelModelType] `json:"model_type,required"`
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
	BillingCycleConfiguration param.Field[NewBillingCycleConfigurationModelParam] `json:"billing_cycle_configuration"`
	BpsConfig                 param.Field[BpsConfigModelParam]                    `json:"bps_config"`
	BulkBpsConfig             param.Field[BulkBpsConfigModelParam]                `json:"bulk_bps_config"`
	BulkConfig                param.Field[BulkConfigModelParam]                   `json:"bulk_config"`
	BulkWithProrationConfig   param.Field[CustomRatingFunctionConfigModelParam]   `json:"bulk_with_proration_config"`
	// The per unit conversion rate of the price currency to the invoicing currency.
	ConversionRate              param.Field[float64]                              `json:"conversion_rate"`
	CumulativeGroupedBulkConfig param.Field[CustomRatingFunctionConfigModelParam] `json:"cumulative_grouped_bulk_config"`
	// An alias for the price.
	ExternalPriceID param.Field[string] `json:"external_price_id"`
	// If the Price represents a fixed cost, this represents the quantity of units
	// applied.
	FixedPriceQuantity               param.Field[float64]                              `json:"fixed_price_quantity"`
	GroupedAllocationConfig          param.Field[CustomRatingFunctionConfigModelParam] `json:"grouped_allocation_config"`
	GroupedTieredConfig              param.Field[CustomRatingFunctionConfigModelParam] `json:"grouped_tiered_config"`
	GroupedTieredPackageConfig       param.Field[CustomRatingFunctionConfigModelParam] `json:"grouped_tiered_package_config"`
	GroupedWithMeteredMinimumConfig  param.Field[CustomRatingFunctionConfigModelParam] `json:"grouped_with_metered_minimum_config"`
	GroupedWithProratedMinimumConfig param.Field[CustomRatingFunctionConfigModelParam] `json:"grouped_with_prorated_minimum_config"`
	// The property used to group this price on an invoice
	InvoiceGroupingKey param.Field[string] `json:"invoice_grouping_key"`
	// Within each billing cycle, specifies the cadence at which invoices are produced.
	// If unspecified, a single invoice is produced per billing cycle.
	InvoicingCycleConfiguration           param.Field[NewBillingCycleConfigurationModelParam] `json:"invoicing_cycle_configuration"`
	MatrixConfig                          param.Field[MatrixConfigModelParam]                 `json:"matrix_config"`
	MatrixWithAllocationConfig            param.Field[MatrixWithAllocationConfigModelParam]   `json:"matrix_with_allocation_config"`
	MatrixWithDisplayNameConfig           param.Field[CustomRatingFunctionConfigModelParam]   `json:"matrix_with_display_name_config"`
	MaxGroupTieredPackageConfig           param.Field[CustomRatingFunctionConfigModelParam]   `json:"max_group_tiered_package_config"`
	Metadata                              param.Field[interface{}]                            `json:"metadata"`
	PackageConfig                         param.Field[PackageConfigModelParam]                `json:"package_config"`
	PackageWithAllocationConfig           param.Field[CustomRatingFunctionConfigModelParam]   `json:"package_with_allocation_config"`
	ScalableMatrixWithTieredPricingConfig param.Field[CustomRatingFunctionConfigModelParam]   `json:"scalable_matrix_with_tiered_pricing_config"`
	ScalableMatrixWithUnitPricingConfig   param.Field[CustomRatingFunctionConfigModelParam]   `json:"scalable_matrix_with_unit_pricing_config"`
	ThresholdTotalAmountConfig            param.Field[CustomRatingFunctionConfigModelParam]   `json:"threshold_total_amount_config"`
	TieredBpsConfig                       param.Field[TieredBpsConfigModelParam]              `json:"tiered_bps_config"`
	TieredConfig                          param.Field[TieredConfigModelParam]                 `json:"tiered_config"`
	TieredPackageConfig                   param.Field[CustomRatingFunctionConfigModelParam]   `json:"tiered_package_config"`
	TieredPackageWithMinimumConfig        param.Field[CustomRatingFunctionConfigModelParam]   `json:"tiered_package_with_minimum_config"`
	TieredWithMinimumConfig               param.Field[CustomRatingFunctionConfigModelParam]   `json:"tiered_with_minimum_config"`
	TieredWithProrationConfig             param.Field[CustomRatingFunctionConfigModelParam]   `json:"tiered_with_proration_config"`
	UnitConfig                            param.Field[UnitConfigModelParam]                   `json:"unit_config"`
	UnitWithPercentConfig                 param.Field[CustomRatingFunctionConfigModelParam]   `json:"unit_with_percent_config"`
	UnitWithProrationConfig               param.Field[CustomRatingFunctionConfigModelParam]   `json:"unit_with_proration_config"`
}

func (r NewFloatingPriceModelParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r NewFloatingPriceModelParam) implementsNewFloatingPriceModelUnionParam() {}

// Satisfied by [shared.NewFloatingPriceModelNewFloatingUnitPriceParam],
// [shared.NewFloatingPriceModelNewFloatingPackagePriceParam],
// [shared.NewFloatingPriceModelNewFloatingMatrixPriceParam],
// [shared.NewFloatingPriceModelNewFloatingMatrixWithAllocationPriceParam],
// [shared.NewFloatingPriceModelNewFloatingTieredPriceParam],
// [shared.NewFloatingPriceModelNewFloatingTieredBpsPriceParam],
// [shared.NewFloatingPriceModelNewFloatingBpsPriceParam],
// [shared.NewFloatingPriceModelNewFloatingBulkBpsPriceParam],
// [shared.NewFloatingPriceModelNewFloatingBulkPriceParam],
// [shared.NewFloatingPriceModelNewFloatingThresholdTotalAmountPriceParam],
// [shared.NewFloatingPriceModelNewFloatingTieredPackagePriceParam],
// [shared.NewFloatingPriceModelNewFloatingGroupedTieredPriceParam],
// [shared.NewFloatingPriceModelNewFloatingMaxGroupTieredPackagePriceParam],
// [shared.NewFloatingPriceModelNewFloatingTieredWithMinimumPriceParam],
// [shared.NewFloatingPriceModelNewFloatingPackageWithAllocationPriceParam],
// [shared.NewFloatingPriceModelNewFloatingTieredPackageWithMinimumPriceParam],
// [shared.NewFloatingPriceModelNewFloatingUnitWithPercentPriceParam],
// [shared.NewFloatingPriceModelNewFloatingTieredWithProrationPriceParam],
// [shared.NewFloatingPriceModelNewFloatingUnitWithProrationPriceParam],
// [shared.NewFloatingPriceModelNewFloatingGroupedAllocationPriceParam],
// [shared.NewFloatingPriceModelNewFloatingGroupedWithProratedMinimumPriceParam],
// [shared.NewFloatingPriceModelNewFloatingGroupedWithMeteredMinimumPriceParam],
// [shared.NewFloatingPriceModelNewFloatingMatrixWithDisplayNamePriceParam],
// [shared.NewFloatingPriceModelNewFloatingBulkWithProrationPriceParam],
// [shared.NewFloatingPriceModelNewFloatingGroupedTieredPackagePriceParam],
// [shared.NewFloatingPriceModelNewFloatingScalableMatrixWithUnitPricingPriceParam],
// [shared.NewFloatingPriceModelNewFloatingScalableMatrixWithTieredPricingPriceParam],
// [shared.NewFloatingPriceModelNewFloatingCumulativeGroupedBulkPriceParam],
// [NewFloatingPriceModelParam].
type NewFloatingPriceModelUnionParam interface {
	implementsNewFloatingPriceModelUnionParam()
}

type NewFloatingPriceModelNewFloatingUnitPriceParam struct {
	// The cadence to bill for this price on.
	Cadence param.Field[NewFloatingPriceModelNewFloatingUnitPriceCadence] `json:"cadence,required"`
	// An ISO 4217 currency string for which this price is billed in.
	Currency param.Field[string] `json:"currency,required"`
	// The id of the item the price will be associated with.
	ItemID    param.Field[string]                                             `json:"item_id,required"`
	ModelType param.Field[NewFloatingPriceModelNewFloatingUnitPriceModelType] `json:"model_type,required"`
	// The name of the price.
	Name       param.Field[string]               `json:"name,required"`
	UnitConfig param.Field[UnitConfigModelParam] `json:"unit_config,required"`
	// The id of the billable metric for the price. Only needed if the price is
	// usage-based.
	BillableMetricID param.Field[string] `json:"billable_metric_id"`
	// If the Price represents a fixed cost, the price will be billed in-advance if
	// this is true, and in-arrears if this is false.
	BilledInAdvance param.Field[bool] `json:"billed_in_advance"`
	// For custom cadence: specifies the duration of the billing period in days or
	// months.
	BillingCycleConfiguration param.Field[NewBillingCycleConfigurationModelParam] `json:"billing_cycle_configuration"`
	// The per unit conversion rate of the price currency to the invoicing currency.
	ConversionRate param.Field[float64] `json:"conversion_rate"`
	// An alias for the price.
	ExternalPriceID param.Field[string] `json:"external_price_id"`
	// If the Price represents a fixed cost, this represents the quantity of units
	// applied.
	FixedPriceQuantity param.Field[float64] `json:"fixed_price_quantity"`
	// The property used to group this price on an invoice
	InvoiceGroupingKey param.Field[string] `json:"invoice_grouping_key"`
	// Within each billing cycle, specifies the cadence at which invoices are produced.
	// If unspecified, a single invoice is produced per billing cycle.
	InvoicingCycleConfiguration param.Field[NewBillingCycleConfigurationModelParam] `json:"invoicing_cycle_configuration"`
	// User-specified key/value pairs for the resource. Individual keys can be removed
	// by setting the value to `null`, and the entire metadata mapping can be cleared
	// by setting `metadata` to `null`.
	Metadata param.Field[map[string]string] `json:"metadata"`
}

func (r NewFloatingPriceModelNewFloatingUnitPriceParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r NewFloatingPriceModelNewFloatingUnitPriceParam) implementsNewFloatingPriceModelUnionParam() {}

// The cadence to bill for this price on.
type NewFloatingPriceModelNewFloatingUnitPriceCadence string

const (
	NewFloatingPriceModelNewFloatingUnitPriceCadenceAnnual     NewFloatingPriceModelNewFloatingUnitPriceCadence = "annual"
	NewFloatingPriceModelNewFloatingUnitPriceCadenceSemiAnnual NewFloatingPriceModelNewFloatingUnitPriceCadence = "semi_annual"
	NewFloatingPriceModelNewFloatingUnitPriceCadenceMonthly    NewFloatingPriceModelNewFloatingUnitPriceCadence = "monthly"
	NewFloatingPriceModelNewFloatingUnitPriceCadenceQuarterly  NewFloatingPriceModelNewFloatingUnitPriceCadence = "quarterly"
	NewFloatingPriceModelNewFloatingUnitPriceCadenceOneTime    NewFloatingPriceModelNewFloatingUnitPriceCadence = "one_time"
	NewFloatingPriceModelNewFloatingUnitPriceCadenceCustom     NewFloatingPriceModelNewFloatingUnitPriceCadence = "custom"
)

func (r NewFloatingPriceModelNewFloatingUnitPriceCadence) IsKnown() bool {
	switch r {
	case NewFloatingPriceModelNewFloatingUnitPriceCadenceAnnual, NewFloatingPriceModelNewFloatingUnitPriceCadenceSemiAnnual, NewFloatingPriceModelNewFloatingUnitPriceCadenceMonthly, NewFloatingPriceModelNewFloatingUnitPriceCadenceQuarterly, NewFloatingPriceModelNewFloatingUnitPriceCadenceOneTime, NewFloatingPriceModelNewFloatingUnitPriceCadenceCustom:
		return true
	}
	return false
}

type NewFloatingPriceModelNewFloatingUnitPriceModelType string

const (
	NewFloatingPriceModelNewFloatingUnitPriceModelTypeUnit NewFloatingPriceModelNewFloatingUnitPriceModelType = "unit"
)

func (r NewFloatingPriceModelNewFloatingUnitPriceModelType) IsKnown() bool {
	switch r {
	case NewFloatingPriceModelNewFloatingUnitPriceModelTypeUnit:
		return true
	}
	return false
}

type NewFloatingPriceModelNewFloatingPackagePriceParam struct {
	// The cadence to bill for this price on.
	Cadence param.Field[NewFloatingPriceModelNewFloatingPackagePriceCadence] `json:"cadence,required"`
	// An ISO 4217 currency string for which this price is billed in.
	Currency param.Field[string] `json:"currency,required"`
	// The id of the item the price will be associated with.
	ItemID    param.Field[string]                                                `json:"item_id,required"`
	ModelType param.Field[NewFloatingPriceModelNewFloatingPackagePriceModelType] `json:"model_type,required"`
	// The name of the price.
	Name          param.Field[string]                  `json:"name,required"`
	PackageConfig param.Field[PackageConfigModelParam] `json:"package_config,required"`
	// The id of the billable metric for the price. Only needed if the price is
	// usage-based.
	BillableMetricID param.Field[string] `json:"billable_metric_id"`
	// If the Price represents a fixed cost, the price will be billed in-advance if
	// this is true, and in-arrears if this is false.
	BilledInAdvance param.Field[bool] `json:"billed_in_advance"`
	// For custom cadence: specifies the duration of the billing period in days or
	// months.
	BillingCycleConfiguration param.Field[NewBillingCycleConfigurationModelParam] `json:"billing_cycle_configuration"`
	// The per unit conversion rate of the price currency to the invoicing currency.
	ConversionRate param.Field[float64] `json:"conversion_rate"`
	// An alias for the price.
	ExternalPriceID param.Field[string] `json:"external_price_id"`
	// If the Price represents a fixed cost, this represents the quantity of units
	// applied.
	FixedPriceQuantity param.Field[float64] `json:"fixed_price_quantity"`
	// The property used to group this price on an invoice
	InvoiceGroupingKey param.Field[string] `json:"invoice_grouping_key"`
	// Within each billing cycle, specifies the cadence at which invoices are produced.
	// If unspecified, a single invoice is produced per billing cycle.
	InvoicingCycleConfiguration param.Field[NewBillingCycleConfigurationModelParam] `json:"invoicing_cycle_configuration"`
	// User-specified key/value pairs for the resource. Individual keys can be removed
	// by setting the value to `null`, and the entire metadata mapping can be cleared
	// by setting `metadata` to `null`.
	Metadata param.Field[map[string]string] `json:"metadata"`
}

func (r NewFloatingPriceModelNewFloatingPackagePriceParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r NewFloatingPriceModelNewFloatingPackagePriceParam) implementsNewFloatingPriceModelUnionParam() {
}

// The cadence to bill for this price on.
type NewFloatingPriceModelNewFloatingPackagePriceCadence string

const (
	NewFloatingPriceModelNewFloatingPackagePriceCadenceAnnual     NewFloatingPriceModelNewFloatingPackagePriceCadence = "annual"
	NewFloatingPriceModelNewFloatingPackagePriceCadenceSemiAnnual NewFloatingPriceModelNewFloatingPackagePriceCadence = "semi_annual"
	NewFloatingPriceModelNewFloatingPackagePriceCadenceMonthly    NewFloatingPriceModelNewFloatingPackagePriceCadence = "monthly"
	NewFloatingPriceModelNewFloatingPackagePriceCadenceQuarterly  NewFloatingPriceModelNewFloatingPackagePriceCadence = "quarterly"
	NewFloatingPriceModelNewFloatingPackagePriceCadenceOneTime    NewFloatingPriceModelNewFloatingPackagePriceCadence = "one_time"
	NewFloatingPriceModelNewFloatingPackagePriceCadenceCustom     NewFloatingPriceModelNewFloatingPackagePriceCadence = "custom"
)

func (r NewFloatingPriceModelNewFloatingPackagePriceCadence) IsKnown() bool {
	switch r {
	case NewFloatingPriceModelNewFloatingPackagePriceCadenceAnnual, NewFloatingPriceModelNewFloatingPackagePriceCadenceSemiAnnual, NewFloatingPriceModelNewFloatingPackagePriceCadenceMonthly, NewFloatingPriceModelNewFloatingPackagePriceCadenceQuarterly, NewFloatingPriceModelNewFloatingPackagePriceCadenceOneTime, NewFloatingPriceModelNewFloatingPackagePriceCadenceCustom:
		return true
	}
	return false
}

type NewFloatingPriceModelNewFloatingPackagePriceModelType string

const (
	NewFloatingPriceModelNewFloatingPackagePriceModelTypePackage NewFloatingPriceModelNewFloatingPackagePriceModelType = "package"
)

func (r NewFloatingPriceModelNewFloatingPackagePriceModelType) IsKnown() bool {
	switch r {
	case NewFloatingPriceModelNewFloatingPackagePriceModelTypePackage:
		return true
	}
	return false
}

type NewFloatingPriceModelNewFloatingMatrixPriceParam struct {
	// The cadence to bill for this price on.
	Cadence param.Field[NewFloatingPriceModelNewFloatingMatrixPriceCadence] `json:"cadence,required"`
	// An ISO 4217 currency string for which this price is billed in.
	Currency param.Field[string] `json:"currency,required"`
	// The id of the item the price will be associated with.
	ItemID       param.Field[string]                                               `json:"item_id,required"`
	MatrixConfig param.Field[MatrixConfigModelParam]                               `json:"matrix_config,required"`
	ModelType    param.Field[NewFloatingPriceModelNewFloatingMatrixPriceModelType] `json:"model_type,required"`
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
	BillingCycleConfiguration param.Field[NewBillingCycleConfigurationModelParam] `json:"billing_cycle_configuration"`
	// The per unit conversion rate of the price currency to the invoicing currency.
	ConversionRate param.Field[float64] `json:"conversion_rate"`
	// An alias for the price.
	ExternalPriceID param.Field[string] `json:"external_price_id"`
	// If the Price represents a fixed cost, this represents the quantity of units
	// applied.
	FixedPriceQuantity param.Field[float64] `json:"fixed_price_quantity"`
	// The property used to group this price on an invoice
	InvoiceGroupingKey param.Field[string] `json:"invoice_grouping_key"`
	// Within each billing cycle, specifies the cadence at which invoices are produced.
	// If unspecified, a single invoice is produced per billing cycle.
	InvoicingCycleConfiguration param.Field[NewBillingCycleConfigurationModelParam] `json:"invoicing_cycle_configuration"`
	// User-specified key/value pairs for the resource. Individual keys can be removed
	// by setting the value to `null`, and the entire metadata mapping can be cleared
	// by setting `metadata` to `null`.
	Metadata param.Field[map[string]string] `json:"metadata"`
}

func (r NewFloatingPriceModelNewFloatingMatrixPriceParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r NewFloatingPriceModelNewFloatingMatrixPriceParam) implementsNewFloatingPriceModelUnionParam() {
}

// The cadence to bill for this price on.
type NewFloatingPriceModelNewFloatingMatrixPriceCadence string

const (
	NewFloatingPriceModelNewFloatingMatrixPriceCadenceAnnual     NewFloatingPriceModelNewFloatingMatrixPriceCadence = "annual"
	NewFloatingPriceModelNewFloatingMatrixPriceCadenceSemiAnnual NewFloatingPriceModelNewFloatingMatrixPriceCadence = "semi_annual"
	NewFloatingPriceModelNewFloatingMatrixPriceCadenceMonthly    NewFloatingPriceModelNewFloatingMatrixPriceCadence = "monthly"
	NewFloatingPriceModelNewFloatingMatrixPriceCadenceQuarterly  NewFloatingPriceModelNewFloatingMatrixPriceCadence = "quarterly"
	NewFloatingPriceModelNewFloatingMatrixPriceCadenceOneTime    NewFloatingPriceModelNewFloatingMatrixPriceCadence = "one_time"
	NewFloatingPriceModelNewFloatingMatrixPriceCadenceCustom     NewFloatingPriceModelNewFloatingMatrixPriceCadence = "custom"
)

func (r NewFloatingPriceModelNewFloatingMatrixPriceCadence) IsKnown() bool {
	switch r {
	case NewFloatingPriceModelNewFloatingMatrixPriceCadenceAnnual, NewFloatingPriceModelNewFloatingMatrixPriceCadenceSemiAnnual, NewFloatingPriceModelNewFloatingMatrixPriceCadenceMonthly, NewFloatingPriceModelNewFloatingMatrixPriceCadenceQuarterly, NewFloatingPriceModelNewFloatingMatrixPriceCadenceOneTime, NewFloatingPriceModelNewFloatingMatrixPriceCadenceCustom:
		return true
	}
	return false
}

type NewFloatingPriceModelNewFloatingMatrixPriceModelType string

const (
	NewFloatingPriceModelNewFloatingMatrixPriceModelTypeMatrix NewFloatingPriceModelNewFloatingMatrixPriceModelType = "matrix"
)

func (r NewFloatingPriceModelNewFloatingMatrixPriceModelType) IsKnown() bool {
	switch r {
	case NewFloatingPriceModelNewFloatingMatrixPriceModelTypeMatrix:
		return true
	}
	return false
}

type NewFloatingPriceModelNewFloatingMatrixWithAllocationPriceParam struct {
	// The cadence to bill for this price on.
	Cadence param.Field[NewFloatingPriceModelNewFloatingMatrixWithAllocationPriceCadence] `json:"cadence,required"`
	// An ISO 4217 currency string for which this price is billed in.
	Currency param.Field[string] `json:"currency,required"`
	// The id of the item the price will be associated with.
	ItemID                     param.Field[string]                                                             `json:"item_id,required"`
	MatrixWithAllocationConfig param.Field[MatrixWithAllocationConfigModelParam]                               `json:"matrix_with_allocation_config,required"`
	ModelType                  param.Field[NewFloatingPriceModelNewFloatingMatrixWithAllocationPriceModelType] `json:"model_type,required"`
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
	BillingCycleConfiguration param.Field[NewBillingCycleConfigurationModelParam] `json:"billing_cycle_configuration"`
	// The per unit conversion rate of the price currency to the invoicing currency.
	ConversionRate param.Field[float64] `json:"conversion_rate"`
	// An alias for the price.
	ExternalPriceID param.Field[string] `json:"external_price_id"`
	// If the Price represents a fixed cost, this represents the quantity of units
	// applied.
	FixedPriceQuantity param.Field[float64] `json:"fixed_price_quantity"`
	// The property used to group this price on an invoice
	InvoiceGroupingKey param.Field[string] `json:"invoice_grouping_key"`
	// Within each billing cycle, specifies the cadence at which invoices are produced.
	// If unspecified, a single invoice is produced per billing cycle.
	InvoicingCycleConfiguration param.Field[NewBillingCycleConfigurationModelParam] `json:"invoicing_cycle_configuration"`
	// User-specified key/value pairs for the resource. Individual keys can be removed
	// by setting the value to `null`, and the entire metadata mapping can be cleared
	// by setting `metadata` to `null`.
	Metadata param.Field[map[string]string] `json:"metadata"`
}

func (r NewFloatingPriceModelNewFloatingMatrixWithAllocationPriceParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r NewFloatingPriceModelNewFloatingMatrixWithAllocationPriceParam) implementsNewFloatingPriceModelUnionParam() {
}

// The cadence to bill for this price on.
type NewFloatingPriceModelNewFloatingMatrixWithAllocationPriceCadence string

const (
	NewFloatingPriceModelNewFloatingMatrixWithAllocationPriceCadenceAnnual     NewFloatingPriceModelNewFloatingMatrixWithAllocationPriceCadence = "annual"
	NewFloatingPriceModelNewFloatingMatrixWithAllocationPriceCadenceSemiAnnual NewFloatingPriceModelNewFloatingMatrixWithAllocationPriceCadence = "semi_annual"
	NewFloatingPriceModelNewFloatingMatrixWithAllocationPriceCadenceMonthly    NewFloatingPriceModelNewFloatingMatrixWithAllocationPriceCadence = "monthly"
	NewFloatingPriceModelNewFloatingMatrixWithAllocationPriceCadenceQuarterly  NewFloatingPriceModelNewFloatingMatrixWithAllocationPriceCadence = "quarterly"
	NewFloatingPriceModelNewFloatingMatrixWithAllocationPriceCadenceOneTime    NewFloatingPriceModelNewFloatingMatrixWithAllocationPriceCadence = "one_time"
	NewFloatingPriceModelNewFloatingMatrixWithAllocationPriceCadenceCustom     NewFloatingPriceModelNewFloatingMatrixWithAllocationPriceCadence = "custom"
)

func (r NewFloatingPriceModelNewFloatingMatrixWithAllocationPriceCadence) IsKnown() bool {
	switch r {
	case NewFloatingPriceModelNewFloatingMatrixWithAllocationPriceCadenceAnnual, NewFloatingPriceModelNewFloatingMatrixWithAllocationPriceCadenceSemiAnnual, NewFloatingPriceModelNewFloatingMatrixWithAllocationPriceCadenceMonthly, NewFloatingPriceModelNewFloatingMatrixWithAllocationPriceCadenceQuarterly, NewFloatingPriceModelNewFloatingMatrixWithAllocationPriceCadenceOneTime, NewFloatingPriceModelNewFloatingMatrixWithAllocationPriceCadenceCustom:
		return true
	}
	return false
}

type NewFloatingPriceModelNewFloatingMatrixWithAllocationPriceModelType string

const (
	NewFloatingPriceModelNewFloatingMatrixWithAllocationPriceModelTypeMatrixWithAllocation NewFloatingPriceModelNewFloatingMatrixWithAllocationPriceModelType = "matrix_with_allocation"
)

func (r NewFloatingPriceModelNewFloatingMatrixWithAllocationPriceModelType) IsKnown() bool {
	switch r {
	case NewFloatingPriceModelNewFloatingMatrixWithAllocationPriceModelTypeMatrixWithAllocation:
		return true
	}
	return false
}

type NewFloatingPriceModelNewFloatingTieredPriceParam struct {
	// The cadence to bill for this price on.
	Cadence param.Field[NewFloatingPriceModelNewFloatingTieredPriceCadence] `json:"cadence,required"`
	// An ISO 4217 currency string for which this price is billed in.
	Currency param.Field[string] `json:"currency,required"`
	// The id of the item the price will be associated with.
	ItemID    param.Field[string]                                               `json:"item_id,required"`
	ModelType param.Field[NewFloatingPriceModelNewFloatingTieredPriceModelType] `json:"model_type,required"`
	// The name of the price.
	Name         param.Field[string]                 `json:"name,required"`
	TieredConfig param.Field[TieredConfigModelParam] `json:"tiered_config,required"`
	// The id of the billable metric for the price. Only needed if the price is
	// usage-based.
	BillableMetricID param.Field[string] `json:"billable_metric_id"`
	// If the Price represents a fixed cost, the price will be billed in-advance if
	// this is true, and in-arrears if this is false.
	BilledInAdvance param.Field[bool] `json:"billed_in_advance"`
	// For custom cadence: specifies the duration of the billing period in days or
	// months.
	BillingCycleConfiguration param.Field[NewBillingCycleConfigurationModelParam] `json:"billing_cycle_configuration"`
	// The per unit conversion rate of the price currency to the invoicing currency.
	ConversionRate param.Field[float64] `json:"conversion_rate"`
	// An alias for the price.
	ExternalPriceID param.Field[string] `json:"external_price_id"`
	// If the Price represents a fixed cost, this represents the quantity of units
	// applied.
	FixedPriceQuantity param.Field[float64] `json:"fixed_price_quantity"`
	// The property used to group this price on an invoice
	InvoiceGroupingKey param.Field[string] `json:"invoice_grouping_key"`
	// Within each billing cycle, specifies the cadence at which invoices are produced.
	// If unspecified, a single invoice is produced per billing cycle.
	InvoicingCycleConfiguration param.Field[NewBillingCycleConfigurationModelParam] `json:"invoicing_cycle_configuration"`
	// User-specified key/value pairs for the resource. Individual keys can be removed
	// by setting the value to `null`, and the entire metadata mapping can be cleared
	// by setting `metadata` to `null`.
	Metadata param.Field[map[string]string] `json:"metadata"`
}

func (r NewFloatingPriceModelNewFloatingTieredPriceParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r NewFloatingPriceModelNewFloatingTieredPriceParam) implementsNewFloatingPriceModelUnionParam() {
}

// The cadence to bill for this price on.
type NewFloatingPriceModelNewFloatingTieredPriceCadence string

const (
	NewFloatingPriceModelNewFloatingTieredPriceCadenceAnnual     NewFloatingPriceModelNewFloatingTieredPriceCadence = "annual"
	NewFloatingPriceModelNewFloatingTieredPriceCadenceSemiAnnual NewFloatingPriceModelNewFloatingTieredPriceCadence = "semi_annual"
	NewFloatingPriceModelNewFloatingTieredPriceCadenceMonthly    NewFloatingPriceModelNewFloatingTieredPriceCadence = "monthly"
	NewFloatingPriceModelNewFloatingTieredPriceCadenceQuarterly  NewFloatingPriceModelNewFloatingTieredPriceCadence = "quarterly"
	NewFloatingPriceModelNewFloatingTieredPriceCadenceOneTime    NewFloatingPriceModelNewFloatingTieredPriceCadence = "one_time"
	NewFloatingPriceModelNewFloatingTieredPriceCadenceCustom     NewFloatingPriceModelNewFloatingTieredPriceCadence = "custom"
)

func (r NewFloatingPriceModelNewFloatingTieredPriceCadence) IsKnown() bool {
	switch r {
	case NewFloatingPriceModelNewFloatingTieredPriceCadenceAnnual, NewFloatingPriceModelNewFloatingTieredPriceCadenceSemiAnnual, NewFloatingPriceModelNewFloatingTieredPriceCadenceMonthly, NewFloatingPriceModelNewFloatingTieredPriceCadenceQuarterly, NewFloatingPriceModelNewFloatingTieredPriceCadenceOneTime, NewFloatingPriceModelNewFloatingTieredPriceCadenceCustom:
		return true
	}
	return false
}

type NewFloatingPriceModelNewFloatingTieredPriceModelType string

const (
	NewFloatingPriceModelNewFloatingTieredPriceModelTypeTiered NewFloatingPriceModelNewFloatingTieredPriceModelType = "tiered"
)

func (r NewFloatingPriceModelNewFloatingTieredPriceModelType) IsKnown() bool {
	switch r {
	case NewFloatingPriceModelNewFloatingTieredPriceModelTypeTiered:
		return true
	}
	return false
}

type NewFloatingPriceModelNewFloatingTieredBpsPriceParam struct {
	// The cadence to bill for this price on.
	Cadence param.Field[NewFloatingPriceModelNewFloatingTieredBpsPriceCadence] `json:"cadence,required"`
	// An ISO 4217 currency string for which this price is billed in.
	Currency param.Field[string] `json:"currency,required"`
	// The id of the item the price will be associated with.
	ItemID    param.Field[string]                                                  `json:"item_id,required"`
	ModelType param.Field[NewFloatingPriceModelNewFloatingTieredBpsPriceModelType] `json:"model_type,required"`
	// The name of the price.
	Name            param.Field[string]                    `json:"name,required"`
	TieredBpsConfig param.Field[TieredBpsConfigModelParam] `json:"tiered_bps_config,required"`
	// The id of the billable metric for the price. Only needed if the price is
	// usage-based.
	BillableMetricID param.Field[string] `json:"billable_metric_id"`
	// If the Price represents a fixed cost, the price will be billed in-advance if
	// this is true, and in-arrears if this is false.
	BilledInAdvance param.Field[bool] `json:"billed_in_advance"`
	// For custom cadence: specifies the duration of the billing period in days or
	// months.
	BillingCycleConfiguration param.Field[NewBillingCycleConfigurationModelParam] `json:"billing_cycle_configuration"`
	// The per unit conversion rate of the price currency to the invoicing currency.
	ConversionRate param.Field[float64] `json:"conversion_rate"`
	// An alias for the price.
	ExternalPriceID param.Field[string] `json:"external_price_id"`
	// If the Price represents a fixed cost, this represents the quantity of units
	// applied.
	FixedPriceQuantity param.Field[float64] `json:"fixed_price_quantity"`
	// The property used to group this price on an invoice
	InvoiceGroupingKey param.Field[string] `json:"invoice_grouping_key"`
	// Within each billing cycle, specifies the cadence at which invoices are produced.
	// If unspecified, a single invoice is produced per billing cycle.
	InvoicingCycleConfiguration param.Field[NewBillingCycleConfigurationModelParam] `json:"invoicing_cycle_configuration"`
	// User-specified key/value pairs for the resource. Individual keys can be removed
	// by setting the value to `null`, and the entire metadata mapping can be cleared
	// by setting `metadata` to `null`.
	Metadata param.Field[map[string]string] `json:"metadata"`
}

func (r NewFloatingPriceModelNewFloatingTieredBpsPriceParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r NewFloatingPriceModelNewFloatingTieredBpsPriceParam) implementsNewFloatingPriceModelUnionParam() {
}

// The cadence to bill for this price on.
type NewFloatingPriceModelNewFloatingTieredBpsPriceCadence string

const (
	NewFloatingPriceModelNewFloatingTieredBpsPriceCadenceAnnual     NewFloatingPriceModelNewFloatingTieredBpsPriceCadence = "annual"
	NewFloatingPriceModelNewFloatingTieredBpsPriceCadenceSemiAnnual NewFloatingPriceModelNewFloatingTieredBpsPriceCadence = "semi_annual"
	NewFloatingPriceModelNewFloatingTieredBpsPriceCadenceMonthly    NewFloatingPriceModelNewFloatingTieredBpsPriceCadence = "monthly"
	NewFloatingPriceModelNewFloatingTieredBpsPriceCadenceQuarterly  NewFloatingPriceModelNewFloatingTieredBpsPriceCadence = "quarterly"
	NewFloatingPriceModelNewFloatingTieredBpsPriceCadenceOneTime    NewFloatingPriceModelNewFloatingTieredBpsPriceCadence = "one_time"
	NewFloatingPriceModelNewFloatingTieredBpsPriceCadenceCustom     NewFloatingPriceModelNewFloatingTieredBpsPriceCadence = "custom"
)

func (r NewFloatingPriceModelNewFloatingTieredBpsPriceCadence) IsKnown() bool {
	switch r {
	case NewFloatingPriceModelNewFloatingTieredBpsPriceCadenceAnnual, NewFloatingPriceModelNewFloatingTieredBpsPriceCadenceSemiAnnual, NewFloatingPriceModelNewFloatingTieredBpsPriceCadenceMonthly, NewFloatingPriceModelNewFloatingTieredBpsPriceCadenceQuarterly, NewFloatingPriceModelNewFloatingTieredBpsPriceCadenceOneTime, NewFloatingPriceModelNewFloatingTieredBpsPriceCadenceCustom:
		return true
	}
	return false
}

type NewFloatingPriceModelNewFloatingTieredBpsPriceModelType string

const (
	NewFloatingPriceModelNewFloatingTieredBpsPriceModelTypeTieredBps NewFloatingPriceModelNewFloatingTieredBpsPriceModelType = "tiered_bps"
)

func (r NewFloatingPriceModelNewFloatingTieredBpsPriceModelType) IsKnown() bool {
	switch r {
	case NewFloatingPriceModelNewFloatingTieredBpsPriceModelTypeTieredBps:
		return true
	}
	return false
}

type NewFloatingPriceModelNewFloatingBpsPriceParam struct {
	BpsConfig param.Field[BpsConfigModelParam] `json:"bps_config,required"`
	// The cadence to bill for this price on.
	Cadence param.Field[NewFloatingPriceModelNewFloatingBpsPriceCadence] `json:"cadence,required"`
	// An ISO 4217 currency string for which this price is billed in.
	Currency param.Field[string] `json:"currency,required"`
	// The id of the item the price will be associated with.
	ItemID    param.Field[string]                                            `json:"item_id,required"`
	ModelType param.Field[NewFloatingPriceModelNewFloatingBpsPriceModelType] `json:"model_type,required"`
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
	BillingCycleConfiguration param.Field[NewBillingCycleConfigurationModelParam] `json:"billing_cycle_configuration"`
	// The per unit conversion rate of the price currency to the invoicing currency.
	ConversionRate param.Field[float64] `json:"conversion_rate"`
	// An alias for the price.
	ExternalPriceID param.Field[string] `json:"external_price_id"`
	// If the Price represents a fixed cost, this represents the quantity of units
	// applied.
	FixedPriceQuantity param.Field[float64] `json:"fixed_price_quantity"`
	// The property used to group this price on an invoice
	InvoiceGroupingKey param.Field[string] `json:"invoice_grouping_key"`
	// Within each billing cycle, specifies the cadence at which invoices are produced.
	// If unspecified, a single invoice is produced per billing cycle.
	InvoicingCycleConfiguration param.Field[NewBillingCycleConfigurationModelParam] `json:"invoicing_cycle_configuration"`
	// User-specified key/value pairs for the resource. Individual keys can be removed
	// by setting the value to `null`, and the entire metadata mapping can be cleared
	// by setting `metadata` to `null`.
	Metadata param.Field[map[string]string] `json:"metadata"`
}

func (r NewFloatingPriceModelNewFloatingBpsPriceParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r NewFloatingPriceModelNewFloatingBpsPriceParam) implementsNewFloatingPriceModelUnionParam() {}

// The cadence to bill for this price on.
type NewFloatingPriceModelNewFloatingBpsPriceCadence string

const (
	NewFloatingPriceModelNewFloatingBpsPriceCadenceAnnual     NewFloatingPriceModelNewFloatingBpsPriceCadence = "annual"
	NewFloatingPriceModelNewFloatingBpsPriceCadenceSemiAnnual NewFloatingPriceModelNewFloatingBpsPriceCadence = "semi_annual"
	NewFloatingPriceModelNewFloatingBpsPriceCadenceMonthly    NewFloatingPriceModelNewFloatingBpsPriceCadence = "monthly"
	NewFloatingPriceModelNewFloatingBpsPriceCadenceQuarterly  NewFloatingPriceModelNewFloatingBpsPriceCadence = "quarterly"
	NewFloatingPriceModelNewFloatingBpsPriceCadenceOneTime    NewFloatingPriceModelNewFloatingBpsPriceCadence = "one_time"
	NewFloatingPriceModelNewFloatingBpsPriceCadenceCustom     NewFloatingPriceModelNewFloatingBpsPriceCadence = "custom"
)

func (r NewFloatingPriceModelNewFloatingBpsPriceCadence) IsKnown() bool {
	switch r {
	case NewFloatingPriceModelNewFloatingBpsPriceCadenceAnnual, NewFloatingPriceModelNewFloatingBpsPriceCadenceSemiAnnual, NewFloatingPriceModelNewFloatingBpsPriceCadenceMonthly, NewFloatingPriceModelNewFloatingBpsPriceCadenceQuarterly, NewFloatingPriceModelNewFloatingBpsPriceCadenceOneTime, NewFloatingPriceModelNewFloatingBpsPriceCadenceCustom:
		return true
	}
	return false
}

type NewFloatingPriceModelNewFloatingBpsPriceModelType string

const (
	NewFloatingPriceModelNewFloatingBpsPriceModelTypeBps NewFloatingPriceModelNewFloatingBpsPriceModelType = "bps"
)

func (r NewFloatingPriceModelNewFloatingBpsPriceModelType) IsKnown() bool {
	switch r {
	case NewFloatingPriceModelNewFloatingBpsPriceModelTypeBps:
		return true
	}
	return false
}

type NewFloatingPriceModelNewFloatingBulkBpsPriceParam struct {
	BulkBpsConfig param.Field[BulkBpsConfigModelParam] `json:"bulk_bps_config,required"`
	// The cadence to bill for this price on.
	Cadence param.Field[NewFloatingPriceModelNewFloatingBulkBpsPriceCadence] `json:"cadence,required"`
	// An ISO 4217 currency string for which this price is billed in.
	Currency param.Field[string] `json:"currency,required"`
	// The id of the item the price will be associated with.
	ItemID    param.Field[string]                                                `json:"item_id,required"`
	ModelType param.Field[NewFloatingPriceModelNewFloatingBulkBpsPriceModelType] `json:"model_type,required"`
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
	BillingCycleConfiguration param.Field[NewBillingCycleConfigurationModelParam] `json:"billing_cycle_configuration"`
	// The per unit conversion rate of the price currency to the invoicing currency.
	ConversionRate param.Field[float64] `json:"conversion_rate"`
	// An alias for the price.
	ExternalPriceID param.Field[string] `json:"external_price_id"`
	// If the Price represents a fixed cost, this represents the quantity of units
	// applied.
	FixedPriceQuantity param.Field[float64] `json:"fixed_price_quantity"`
	// The property used to group this price on an invoice
	InvoiceGroupingKey param.Field[string] `json:"invoice_grouping_key"`
	// Within each billing cycle, specifies the cadence at which invoices are produced.
	// If unspecified, a single invoice is produced per billing cycle.
	InvoicingCycleConfiguration param.Field[NewBillingCycleConfigurationModelParam] `json:"invoicing_cycle_configuration"`
	// User-specified key/value pairs for the resource. Individual keys can be removed
	// by setting the value to `null`, and the entire metadata mapping can be cleared
	// by setting `metadata` to `null`.
	Metadata param.Field[map[string]string] `json:"metadata"`
}

func (r NewFloatingPriceModelNewFloatingBulkBpsPriceParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r NewFloatingPriceModelNewFloatingBulkBpsPriceParam) implementsNewFloatingPriceModelUnionParam() {
}

// The cadence to bill for this price on.
type NewFloatingPriceModelNewFloatingBulkBpsPriceCadence string

const (
	NewFloatingPriceModelNewFloatingBulkBpsPriceCadenceAnnual     NewFloatingPriceModelNewFloatingBulkBpsPriceCadence = "annual"
	NewFloatingPriceModelNewFloatingBulkBpsPriceCadenceSemiAnnual NewFloatingPriceModelNewFloatingBulkBpsPriceCadence = "semi_annual"
	NewFloatingPriceModelNewFloatingBulkBpsPriceCadenceMonthly    NewFloatingPriceModelNewFloatingBulkBpsPriceCadence = "monthly"
	NewFloatingPriceModelNewFloatingBulkBpsPriceCadenceQuarterly  NewFloatingPriceModelNewFloatingBulkBpsPriceCadence = "quarterly"
	NewFloatingPriceModelNewFloatingBulkBpsPriceCadenceOneTime    NewFloatingPriceModelNewFloatingBulkBpsPriceCadence = "one_time"
	NewFloatingPriceModelNewFloatingBulkBpsPriceCadenceCustom     NewFloatingPriceModelNewFloatingBulkBpsPriceCadence = "custom"
)

func (r NewFloatingPriceModelNewFloatingBulkBpsPriceCadence) IsKnown() bool {
	switch r {
	case NewFloatingPriceModelNewFloatingBulkBpsPriceCadenceAnnual, NewFloatingPriceModelNewFloatingBulkBpsPriceCadenceSemiAnnual, NewFloatingPriceModelNewFloatingBulkBpsPriceCadenceMonthly, NewFloatingPriceModelNewFloatingBulkBpsPriceCadenceQuarterly, NewFloatingPriceModelNewFloatingBulkBpsPriceCadenceOneTime, NewFloatingPriceModelNewFloatingBulkBpsPriceCadenceCustom:
		return true
	}
	return false
}

type NewFloatingPriceModelNewFloatingBulkBpsPriceModelType string

const (
	NewFloatingPriceModelNewFloatingBulkBpsPriceModelTypeBulkBps NewFloatingPriceModelNewFloatingBulkBpsPriceModelType = "bulk_bps"
)

func (r NewFloatingPriceModelNewFloatingBulkBpsPriceModelType) IsKnown() bool {
	switch r {
	case NewFloatingPriceModelNewFloatingBulkBpsPriceModelTypeBulkBps:
		return true
	}
	return false
}

type NewFloatingPriceModelNewFloatingBulkPriceParam struct {
	BulkConfig param.Field[BulkConfigModelParam] `json:"bulk_config,required"`
	// The cadence to bill for this price on.
	Cadence param.Field[NewFloatingPriceModelNewFloatingBulkPriceCadence] `json:"cadence,required"`
	// An ISO 4217 currency string for which this price is billed in.
	Currency param.Field[string] `json:"currency,required"`
	// The id of the item the price will be associated with.
	ItemID    param.Field[string]                                             `json:"item_id,required"`
	ModelType param.Field[NewFloatingPriceModelNewFloatingBulkPriceModelType] `json:"model_type,required"`
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
	BillingCycleConfiguration param.Field[NewBillingCycleConfigurationModelParam] `json:"billing_cycle_configuration"`
	// The per unit conversion rate of the price currency to the invoicing currency.
	ConversionRate param.Field[float64] `json:"conversion_rate"`
	// An alias for the price.
	ExternalPriceID param.Field[string] `json:"external_price_id"`
	// If the Price represents a fixed cost, this represents the quantity of units
	// applied.
	FixedPriceQuantity param.Field[float64] `json:"fixed_price_quantity"`
	// The property used to group this price on an invoice
	InvoiceGroupingKey param.Field[string] `json:"invoice_grouping_key"`
	// Within each billing cycle, specifies the cadence at which invoices are produced.
	// If unspecified, a single invoice is produced per billing cycle.
	InvoicingCycleConfiguration param.Field[NewBillingCycleConfigurationModelParam] `json:"invoicing_cycle_configuration"`
	// User-specified key/value pairs for the resource. Individual keys can be removed
	// by setting the value to `null`, and the entire metadata mapping can be cleared
	// by setting `metadata` to `null`.
	Metadata param.Field[map[string]string] `json:"metadata"`
}

func (r NewFloatingPriceModelNewFloatingBulkPriceParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r NewFloatingPriceModelNewFloatingBulkPriceParam) implementsNewFloatingPriceModelUnionParam() {}

// The cadence to bill for this price on.
type NewFloatingPriceModelNewFloatingBulkPriceCadence string

const (
	NewFloatingPriceModelNewFloatingBulkPriceCadenceAnnual     NewFloatingPriceModelNewFloatingBulkPriceCadence = "annual"
	NewFloatingPriceModelNewFloatingBulkPriceCadenceSemiAnnual NewFloatingPriceModelNewFloatingBulkPriceCadence = "semi_annual"
	NewFloatingPriceModelNewFloatingBulkPriceCadenceMonthly    NewFloatingPriceModelNewFloatingBulkPriceCadence = "monthly"
	NewFloatingPriceModelNewFloatingBulkPriceCadenceQuarterly  NewFloatingPriceModelNewFloatingBulkPriceCadence = "quarterly"
	NewFloatingPriceModelNewFloatingBulkPriceCadenceOneTime    NewFloatingPriceModelNewFloatingBulkPriceCadence = "one_time"
	NewFloatingPriceModelNewFloatingBulkPriceCadenceCustom     NewFloatingPriceModelNewFloatingBulkPriceCadence = "custom"
)

func (r NewFloatingPriceModelNewFloatingBulkPriceCadence) IsKnown() bool {
	switch r {
	case NewFloatingPriceModelNewFloatingBulkPriceCadenceAnnual, NewFloatingPriceModelNewFloatingBulkPriceCadenceSemiAnnual, NewFloatingPriceModelNewFloatingBulkPriceCadenceMonthly, NewFloatingPriceModelNewFloatingBulkPriceCadenceQuarterly, NewFloatingPriceModelNewFloatingBulkPriceCadenceOneTime, NewFloatingPriceModelNewFloatingBulkPriceCadenceCustom:
		return true
	}
	return false
}

type NewFloatingPriceModelNewFloatingBulkPriceModelType string

const (
	NewFloatingPriceModelNewFloatingBulkPriceModelTypeBulk NewFloatingPriceModelNewFloatingBulkPriceModelType = "bulk"
)

func (r NewFloatingPriceModelNewFloatingBulkPriceModelType) IsKnown() bool {
	switch r {
	case NewFloatingPriceModelNewFloatingBulkPriceModelTypeBulk:
		return true
	}
	return false
}

type NewFloatingPriceModelNewFloatingThresholdTotalAmountPriceParam struct {
	// The cadence to bill for this price on.
	Cadence param.Field[NewFloatingPriceModelNewFloatingThresholdTotalAmountPriceCadence] `json:"cadence,required"`
	// An ISO 4217 currency string for which this price is billed in.
	Currency param.Field[string] `json:"currency,required"`
	// The id of the item the price will be associated with.
	ItemID    param.Field[string]                                                             `json:"item_id,required"`
	ModelType param.Field[NewFloatingPriceModelNewFloatingThresholdTotalAmountPriceModelType] `json:"model_type,required"`
	// The name of the price.
	Name                       param.Field[string]                               `json:"name,required"`
	ThresholdTotalAmountConfig param.Field[CustomRatingFunctionConfigModelParam] `json:"threshold_total_amount_config,required"`
	// The id of the billable metric for the price. Only needed if the price is
	// usage-based.
	BillableMetricID param.Field[string] `json:"billable_metric_id"`
	// If the Price represents a fixed cost, the price will be billed in-advance if
	// this is true, and in-arrears if this is false.
	BilledInAdvance param.Field[bool] `json:"billed_in_advance"`
	// For custom cadence: specifies the duration of the billing period in days or
	// months.
	BillingCycleConfiguration param.Field[NewBillingCycleConfigurationModelParam] `json:"billing_cycle_configuration"`
	// The per unit conversion rate of the price currency to the invoicing currency.
	ConversionRate param.Field[float64] `json:"conversion_rate"`
	// An alias for the price.
	ExternalPriceID param.Field[string] `json:"external_price_id"`
	// If the Price represents a fixed cost, this represents the quantity of units
	// applied.
	FixedPriceQuantity param.Field[float64] `json:"fixed_price_quantity"`
	// The property used to group this price on an invoice
	InvoiceGroupingKey param.Field[string] `json:"invoice_grouping_key"`
	// Within each billing cycle, specifies the cadence at which invoices are produced.
	// If unspecified, a single invoice is produced per billing cycle.
	InvoicingCycleConfiguration param.Field[NewBillingCycleConfigurationModelParam] `json:"invoicing_cycle_configuration"`
	// User-specified key/value pairs for the resource. Individual keys can be removed
	// by setting the value to `null`, and the entire metadata mapping can be cleared
	// by setting `metadata` to `null`.
	Metadata param.Field[map[string]string] `json:"metadata"`
}

func (r NewFloatingPriceModelNewFloatingThresholdTotalAmountPriceParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r NewFloatingPriceModelNewFloatingThresholdTotalAmountPriceParam) implementsNewFloatingPriceModelUnionParam() {
}

// The cadence to bill for this price on.
type NewFloatingPriceModelNewFloatingThresholdTotalAmountPriceCadence string

const (
	NewFloatingPriceModelNewFloatingThresholdTotalAmountPriceCadenceAnnual     NewFloatingPriceModelNewFloatingThresholdTotalAmountPriceCadence = "annual"
	NewFloatingPriceModelNewFloatingThresholdTotalAmountPriceCadenceSemiAnnual NewFloatingPriceModelNewFloatingThresholdTotalAmountPriceCadence = "semi_annual"
	NewFloatingPriceModelNewFloatingThresholdTotalAmountPriceCadenceMonthly    NewFloatingPriceModelNewFloatingThresholdTotalAmountPriceCadence = "monthly"
	NewFloatingPriceModelNewFloatingThresholdTotalAmountPriceCadenceQuarterly  NewFloatingPriceModelNewFloatingThresholdTotalAmountPriceCadence = "quarterly"
	NewFloatingPriceModelNewFloatingThresholdTotalAmountPriceCadenceOneTime    NewFloatingPriceModelNewFloatingThresholdTotalAmountPriceCadence = "one_time"
	NewFloatingPriceModelNewFloatingThresholdTotalAmountPriceCadenceCustom     NewFloatingPriceModelNewFloatingThresholdTotalAmountPriceCadence = "custom"
)

func (r NewFloatingPriceModelNewFloatingThresholdTotalAmountPriceCadence) IsKnown() bool {
	switch r {
	case NewFloatingPriceModelNewFloatingThresholdTotalAmountPriceCadenceAnnual, NewFloatingPriceModelNewFloatingThresholdTotalAmountPriceCadenceSemiAnnual, NewFloatingPriceModelNewFloatingThresholdTotalAmountPriceCadenceMonthly, NewFloatingPriceModelNewFloatingThresholdTotalAmountPriceCadenceQuarterly, NewFloatingPriceModelNewFloatingThresholdTotalAmountPriceCadenceOneTime, NewFloatingPriceModelNewFloatingThresholdTotalAmountPriceCadenceCustom:
		return true
	}
	return false
}

type NewFloatingPriceModelNewFloatingThresholdTotalAmountPriceModelType string

const (
	NewFloatingPriceModelNewFloatingThresholdTotalAmountPriceModelTypeThresholdTotalAmount NewFloatingPriceModelNewFloatingThresholdTotalAmountPriceModelType = "threshold_total_amount"
)

func (r NewFloatingPriceModelNewFloatingThresholdTotalAmountPriceModelType) IsKnown() bool {
	switch r {
	case NewFloatingPriceModelNewFloatingThresholdTotalAmountPriceModelTypeThresholdTotalAmount:
		return true
	}
	return false
}

type NewFloatingPriceModelNewFloatingTieredPackagePriceParam struct {
	// The cadence to bill for this price on.
	Cadence param.Field[NewFloatingPriceModelNewFloatingTieredPackagePriceCadence] `json:"cadence,required"`
	// An ISO 4217 currency string for which this price is billed in.
	Currency param.Field[string] `json:"currency,required"`
	// The id of the item the price will be associated with.
	ItemID    param.Field[string]                                                      `json:"item_id,required"`
	ModelType param.Field[NewFloatingPriceModelNewFloatingTieredPackagePriceModelType] `json:"model_type,required"`
	// The name of the price.
	Name                param.Field[string]                               `json:"name,required"`
	TieredPackageConfig param.Field[CustomRatingFunctionConfigModelParam] `json:"tiered_package_config,required"`
	// The id of the billable metric for the price. Only needed if the price is
	// usage-based.
	BillableMetricID param.Field[string] `json:"billable_metric_id"`
	// If the Price represents a fixed cost, the price will be billed in-advance if
	// this is true, and in-arrears if this is false.
	BilledInAdvance param.Field[bool] `json:"billed_in_advance"`
	// For custom cadence: specifies the duration of the billing period in days or
	// months.
	BillingCycleConfiguration param.Field[NewBillingCycleConfigurationModelParam] `json:"billing_cycle_configuration"`
	// The per unit conversion rate of the price currency to the invoicing currency.
	ConversionRate param.Field[float64] `json:"conversion_rate"`
	// An alias for the price.
	ExternalPriceID param.Field[string] `json:"external_price_id"`
	// If the Price represents a fixed cost, this represents the quantity of units
	// applied.
	FixedPriceQuantity param.Field[float64] `json:"fixed_price_quantity"`
	// The property used to group this price on an invoice
	InvoiceGroupingKey param.Field[string] `json:"invoice_grouping_key"`
	// Within each billing cycle, specifies the cadence at which invoices are produced.
	// If unspecified, a single invoice is produced per billing cycle.
	InvoicingCycleConfiguration param.Field[NewBillingCycleConfigurationModelParam] `json:"invoicing_cycle_configuration"`
	// User-specified key/value pairs for the resource. Individual keys can be removed
	// by setting the value to `null`, and the entire metadata mapping can be cleared
	// by setting `metadata` to `null`.
	Metadata param.Field[map[string]string] `json:"metadata"`
}

func (r NewFloatingPriceModelNewFloatingTieredPackagePriceParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r NewFloatingPriceModelNewFloatingTieredPackagePriceParam) implementsNewFloatingPriceModelUnionParam() {
}

// The cadence to bill for this price on.
type NewFloatingPriceModelNewFloatingTieredPackagePriceCadence string

const (
	NewFloatingPriceModelNewFloatingTieredPackagePriceCadenceAnnual     NewFloatingPriceModelNewFloatingTieredPackagePriceCadence = "annual"
	NewFloatingPriceModelNewFloatingTieredPackagePriceCadenceSemiAnnual NewFloatingPriceModelNewFloatingTieredPackagePriceCadence = "semi_annual"
	NewFloatingPriceModelNewFloatingTieredPackagePriceCadenceMonthly    NewFloatingPriceModelNewFloatingTieredPackagePriceCadence = "monthly"
	NewFloatingPriceModelNewFloatingTieredPackagePriceCadenceQuarterly  NewFloatingPriceModelNewFloatingTieredPackagePriceCadence = "quarterly"
	NewFloatingPriceModelNewFloatingTieredPackagePriceCadenceOneTime    NewFloatingPriceModelNewFloatingTieredPackagePriceCadence = "one_time"
	NewFloatingPriceModelNewFloatingTieredPackagePriceCadenceCustom     NewFloatingPriceModelNewFloatingTieredPackagePriceCadence = "custom"
)

func (r NewFloatingPriceModelNewFloatingTieredPackagePriceCadence) IsKnown() bool {
	switch r {
	case NewFloatingPriceModelNewFloatingTieredPackagePriceCadenceAnnual, NewFloatingPriceModelNewFloatingTieredPackagePriceCadenceSemiAnnual, NewFloatingPriceModelNewFloatingTieredPackagePriceCadenceMonthly, NewFloatingPriceModelNewFloatingTieredPackagePriceCadenceQuarterly, NewFloatingPriceModelNewFloatingTieredPackagePriceCadenceOneTime, NewFloatingPriceModelNewFloatingTieredPackagePriceCadenceCustom:
		return true
	}
	return false
}

type NewFloatingPriceModelNewFloatingTieredPackagePriceModelType string

const (
	NewFloatingPriceModelNewFloatingTieredPackagePriceModelTypeTieredPackage NewFloatingPriceModelNewFloatingTieredPackagePriceModelType = "tiered_package"
)

func (r NewFloatingPriceModelNewFloatingTieredPackagePriceModelType) IsKnown() bool {
	switch r {
	case NewFloatingPriceModelNewFloatingTieredPackagePriceModelTypeTieredPackage:
		return true
	}
	return false
}

type NewFloatingPriceModelNewFloatingGroupedTieredPriceParam struct {
	// The cadence to bill for this price on.
	Cadence param.Field[NewFloatingPriceModelNewFloatingGroupedTieredPriceCadence] `json:"cadence,required"`
	// An ISO 4217 currency string for which this price is billed in.
	Currency            param.Field[string]                               `json:"currency,required"`
	GroupedTieredConfig param.Field[CustomRatingFunctionConfigModelParam] `json:"grouped_tiered_config,required"`
	// The id of the item the price will be associated with.
	ItemID    param.Field[string]                                                      `json:"item_id,required"`
	ModelType param.Field[NewFloatingPriceModelNewFloatingGroupedTieredPriceModelType] `json:"model_type,required"`
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
	BillingCycleConfiguration param.Field[NewBillingCycleConfigurationModelParam] `json:"billing_cycle_configuration"`
	// The per unit conversion rate of the price currency to the invoicing currency.
	ConversionRate param.Field[float64] `json:"conversion_rate"`
	// An alias for the price.
	ExternalPriceID param.Field[string] `json:"external_price_id"`
	// If the Price represents a fixed cost, this represents the quantity of units
	// applied.
	FixedPriceQuantity param.Field[float64] `json:"fixed_price_quantity"`
	// The property used to group this price on an invoice
	InvoiceGroupingKey param.Field[string] `json:"invoice_grouping_key"`
	// Within each billing cycle, specifies the cadence at which invoices are produced.
	// If unspecified, a single invoice is produced per billing cycle.
	InvoicingCycleConfiguration param.Field[NewBillingCycleConfigurationModelParam] `json:"invoicing_cycle_configuration"`
	// User-specified key/value pairs for the resource. Individual keys can be removed
	// by setting the value to `null`, and the entire metadata mapping can be cleared
	// by setting `metadata` to `null`.
	Metadata param.Field[map[string]string] `json:"metadata"`
}

func (r NewFloatingPriceModelNewFloatingGroupedTieredPriceParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r NewFloatingPriceModelNewFloatingGroupedTieredPriceParam) implementsNewFloatingPriceModelUnionParam() {
}

// The cadence to bill for this price on.
type NewFloatingPriceModelNewFloatingGroupedTieredPriceCadence string

const (
	NewFloatingPriceModelNewFloatingGroupedTieredPriceCadenceAnnual     NewFloatingPriceModelNewFloatingGroupedTieredPriceCadence = "annual"
	NewFloatingPriceModelNewFloatingGroupedTieredPriceCadenceSemiAnnual NewFloatingPriceModelNewFloatingGroupedTieredPriceCadence = "semi_annual"
	NewFloatingPriceModelNewFloatingGroupedTieredPriceCadenceMonthly    NewFloatingPriceModelNewFloatingGroupedTieredPriceCadence = "monthly"
	NewFloatingPriceModelNewFloatingGroupedTieredPriceCadenceQuarterly  NewFloatingPriceModelNewFloatingGroupedTieredPriceCadence = "quarterly"
	NewFloatingPriceModelNewFloatingGroupedTieredPriceCadenceOneTime    NewFloatingPriceModelNewFloatingGroupedTieredPriceCadence = "one_time"
	NewFloatingPriceModelNewFloatingGroupedTieredPriceCadenceCustom     NewFloatingPriceModelNewFloatingGroupedTieredPriceCadence = "custom"
)

func (r NewFloatingPriceModelNewFloatingGroupedTieredPriceCadence) IsKnown() bool {
	switch r {
	case NewFloatingPriceModelNewFloatingGroupedTieredPriceCadenceAnnual, NewFloatingPriceModelNewFloatingGroupedTieredPriceCadenceSemiAnnual, NewFloatingPriceModelNewFloatingGroupedTieredPriceCadenceMonthly, NewFloatingPriceModelNewFloatingGroupedTieredPriceCadenceQuarterly, NewFloatingPriceModelNewFloatingGroupedTieredPriceCadenceOneTime, NewFloatingPriceModelNewFloatingGroupedTieredPriceCadenceCustom:
		return true
	}
	return false
}

type NewFloatingPriceModelNewFloatingGroupedTieredPriceModelType string

const (
	NewFloatingPriceModelNewFloatingGroupedTieredPriceModelTypeGroupedTiered NewFloatingPriceModelNewFloatingGroupedTieredPriceModelType = "grouped_tiered"
)

func (r NewFloatingPriceModelNewFloatingGroupedTieredPriceModelType) IsKnown() bool {
	switch r {
	case NewFloatingPriceModelNewFloatingGroupedTieredPriceModelTypeGroupedTiered:
		return true
	}
	return false
}

type NewFloatingPriceModelNewFloatingMaxGroupTieredPackagePriceParam struct {
	// The cadence to bill for this price on.
	Cadence param.Field[NewFloatingPriceModelNewFloatingMaxGroupTieredPackagePriceCadence] `json:"cadence,required"`
	// An ISO 4217 currency string for which this price is billed in.
	Currency param.Field[string] `json:"currency,required"`
	// The id of the item the price will be associated with.
	ItemID                      param.Field[string]                                                              `json:"item_id,required"`
	MaxGroupTieredPackageConfig param.Field[CustomRatingFunctionConfigModelParam]                                `json:"max_group_tiered_package_config,required"`
	ModelType                   param.Field[NewFloatingPriceModelNewFloatingMaxGroupTieredPackagePriceModelType] `json:"model_type,required"`
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
	BillingCycleConfiguration param.Field[NewBillingCycleConfigurationModelParam] `json:"billing_cycle_configuration"`
	// The per unit conversion rate of the price currency to the invoicing currency.
	ConversionRate param.Field[float64] `json:"conversion_rate"`
	// An alias for the price.
	ExternalPriceID param.Field[string] `json:"external_price_id"`
	// If the Price represents a fixed cost, this represents the quantity of units
	// applied.
	FixedPriceQuantity param.Field[float64] `json:"fixed_price_quantity"`
	// The property used to group this price on an invoice
	InvoiceGroupingKey param.Field[string] `json:"invoice_grouping_key"`
	// Within each billing cycle, specifies the cadence at which invoices are produced.
	// If unspecified, a single invoice is produced per billing cycle.
	InvoicingCycleConfiguration param.Field[NewBillingCycleConfigurationModelParam] `json:"invoicing_cycle_configuration"`
	// User-specified key/value pairs for the resource. Individual keys can be removed
	// by setting the value to `null`, and the entire metadata mapping can be cleared
	// by setting `metadata` to `null`.
	Metadata param.Field[map[string]string] `json:"metadata"`
}

func (r NewFloatingPriceModelNewFloatingMaxGroupTieredPackagePriceParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r NewFloatingPriceModelNewFloatingMaxGroupTieredPackagePriceParam) implementsNewFloatingPriceModelUnionParam() {
}

// The cadence to bill for this price on.
type NewFloatingPriceModelNewFloatingMaxGroupTieredPackagePriceCadence string

const (
	NewFloatingPriceModelNewFloatingMaxGroupTieredPackagePriceCadenceAnnual     NewFloatingPriceModelNewFloatingMaxGroupTieredPackagePriceCadence = "annual"
	NewFloatingPriceModelNewFloatingMaxGroupTieredPackagePriceCadenceSemiAnnual NewFloatingPriceModelNewFloatingMaxGroupTieredPackagePriceCadence = "semi_annual"
	NewFloatingPriceModelNewFloatingMaxGroupTieredPackagePriceCadenceMonthly    NewFloatingPriceModelNewFloatingMaxGroupTieredPackagePriceCadence = "monthly"
	NewFloatingPriceModelNewFloatingMaxGroupTieredPackagePriceCadenceQuarterly  NewFloatingPriceModelNewFloatingMaxGroupTieredPackagePriceCadence = "quarterly"
	NewFloatingPriceModelNewFloatingMaxGroupTieredPackagePriceCadenceOneTime    NewFloatingPriceModelNewFloatingMaxGroupTieredPackagePriceCadence = "one_time"
	NewFloatingPriceModelNewFloatingMaxGroupTieredPackagePriceCadenceCustom     NewFloatingPriceModelNewFloatingMaxGroupTieredPackagePriceCadence = "custom"
)

func (r NewFloatingPriceModelNewFloatingMaxGroupTieredPackagePriceCadence) IsKnown() bool {
	switch r {
	case NewFloatingPriceModelNewFloatingMaxGroupTieredPackagePriceCadenceAnnual, NewFloatingPriceModelNewFloatingMaxGroupTieredPackagePriceCadenceSemiAnnual, NewFloatingPriceModelNewFloatingMaxGroupTieredPackagePriceCadenceMonthly, NewFloatingPriceModelNewFloatingMaxGroupTieredPackagePriceCadenceQuarterly, NewFloatingPriceModelNewFloatingMaxGroupTieredPackagePriceCadenceOneTime, NewFloatingPriceModelNewFloatingMaxGroupTieredPackagePriceCadenceCustom:
		return true
	}
	return false
}

type NewFloatingPriceModelNewFloatingMaxGroupTieredPackagePriceModelType string

const (
	NewFloatingPriceModelNewFloatingMaxGroupTieredPackagePriceModelTypeMaxGroupTieredPackage NewFloatingPriceModelNewFloatingMaxGroupTieredPackagePriceModelType = "max_group_tiered_package"
)

func (r NewFloatingPriceModelNewFloatingMaxGroupTieredPackagePriceModelType) IsKnown() bool {
	switch r {
	case NewFloatingPriceModelNewFloatingMaxGroupTieredPackagePriceModelTypeMaxGroupTieredPackage:
		return true
	}
	return false
}

type NewFloatingPriceModelNewFloatingTieredWithMinimumPriceParam struct {
	// The cadence to bill for this price on.
	Cadence param.Field[NewFloatingPriceModelNewFloatingTieredWithMinimumPriceCadence] `json:"cadence,required"`
	// An ISO 4217 currency string for which this price is billed in.
	Currency param.Field[string] `json:"currency,required"`
	// The id of the item the price will be associated with.
	ItemID    param.Field[string]                                                          `json:"item_id,required"`
	ModelType param.Field[NewFloatingPriceModelNewFloatingTieredWithMinimumPriceModelType] `json:"model_type,required"`
	// The name of the price.
	Name                    param.Field[string]                               `json:"name,required"`
	TieredWithMinimumConfig param.Field[CustomRatingFunctionConfigModelParam] `json:"tiered_with_minimum_config,required"`
	// The id of the billable metric for the price. Only needed if the price is
	// usage-based.
	BillableMetricID param.Field[string] `json:"billable_metric_id"`
	// If the Price represents a fixed cost, the price will be billed in-advance if
	// this is true, and in-arrears if this is false.
	BilledInAdvance param.Field[bool] `json:"billed_in_advance"`
	// For custom cadence: specifies the duration of the billing period in days or
	// months.
	BillingCycleConfiguration param.Field[NewBillingCycleConfigurationModelParam] `json:"billing_cycle_configuration"`
	// The per unit conversion rate of the price currency to the invoicing currency.
	ConversionRate param.Field[float64] `json:"conversion_rate"`
	// An alias for the price.
	ExternalPriceID param.Field[string] `json:"external_price_id"`
	// If the Price represents a fixed cost, this represents the quantity of units
	// applied.
	FixedPriceQuantity param.Field[float64] `json:"fixed_price_quantity"`
	// The property used to group this price on an invoice
	InvoiceGroupingKey param.Field[string] `json:"invoice_grouping_key"`
	// Within each billing cycle, specifies the cadence at which invoices are produced.
	// If unspecified, a single invoice is produced per billing cycle.
	InvoicingCycleConfiguration param.Field[NewBillingCycleConfigurationModelParam] `json:"invoicing_cycle_configuration"`
	// User-specified key/value pairs for the resource. Individual keys can be removed
	// by setting the value to `null`, and the entire metadata mapping can be cleared
	// by setting `metadata` to `null`.
	Metadata param.Field[map[string]string] `json:"metadata"`
}

func (r NewFloatingPriceModelNewFloatingTieredWithMinimumPriceParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r NewFloatingPriceModelNewFloatingTieredWithMinimumPriceParam) implementsNewFloatingPriceModelUnionParam() {
}

// The cadence to bill for this price on.
type NewFloatingPriceModelNewFloatingTieredWithMinimumPriceCadence string

const (
	NewFloatingPriceModelNewFloatingTieredWithMinimumPriceCadenceAnnual     NewFloatingPriceModelNewFloatingTieredWithMinimumPriceCadence = "annual"
	NewFloatingPriceModelNewFloatingTieredWithMinimumPriceCadenceSemiAnnual NewFloatingPriceModelNewFloatingTieredWithMinimumPriceCadence = "semi_annual"
	NewFloatingPriceModelNewFloatingTieredWithMinimumPriceCadenceMonthly    NewFloatingPriceModelNewFloatingTieredWithMinimumPriceCadence = "monthly"
	NewFloatingPriceModelNewFloatingTieredWithMinimumPriceCadenceQuarterly  NewFloatingPriceModelNewFloatingTieredWithMinimumPriceCadence = "quarterly"
	NewFloatingPriceModelNewFloatingTieredWithMinimumPriceCadenceOneTime    NewFloatingPriceModelNewFloatingTieredWithMinimumPriceCadence = "one_time"
	NewFloatingPriceModelNewFloatingTieredWithMinimumPriceCadenceCustom     NewFloatingPriceModelNewFloatingTieredWithMinimumPriceCadence = "custom"
)

func (r NewFloatingPriceModelNewFloatingTieredWithMinimumPriceCadence) IsKnown() bool {
	switch r {
	case NewFloatingPriceModelNewFloatingTieredWithMinimumPriceCadenceAnnual, NewFloatingPriceModelNewFloatingTieredWithMinimumPriceCadenceSemiAnnual, NewFloatingPriceModelNewFloatingTieredWithMinimumPriceCadenceMonthly, NewFloatingPriceModelNewFloatingTieredWithMinimumPriceCadenceQuarterly, NewFloatingPriceModelNewFloatingTieredWithMinimumPriceCadenceOneTime, NewFloatingPriceModelNewFloatingTieredWithMinimumPriceCadenceCustom:
		return true
	}
	return false
}

type NewFloatingPriceModelNewFloatingTieredWithMinimumPriceModelType string

const (
	NewFloatingPriceModelNewFloatingTieredWithMinimumPriceModelTypeTieredWithMinimum NewFloatingPriceModelNewFloatingTieredWithMinimumPriceModelType = "tiered_with_minimum"
)

func (r NewFloatingPriceModelNewFloatingTieredWithMinimumPriceModelType) IsKnown() bool {
	switch r {
	case NewFloatingPriceModelNewFloatingTieredWithMinimumPriceModelTypeTieredWithMinimum:
		return true
	}
	return false
}

type NewFloatingPriceModelNewFloatingPackageWithAllocationPriceParam struct {
	// The cadence to bill for this price on.
	Cadence param.Field[NewFloatingPriceModelNewFloatingPackageWithAllocationPriceCadence] `json:"cadence,required"`
	// An ISO 4217 currency string for which this price is billed in.
	Currency param.Field[string] `json:"currency,required"`
	// The id of the item the price will be associated with.
	ItemID    param.Field[string]                                                              `json:"item_id,required"`
	ModelType param.Field[NewFloatingPriceModelNewFloatingPackageWithAllocationPriceModelType] `json:"model_type,required"`
	// The name of the price.
	Name                        param.Field[string]                               `json:"name,required"`
	PackageWithAllocationConfig param.Field[CustomRatingFunctionConfigModelParam] `json:"package_with_allocation_config,required"`
	// The id of the billable metric for the price. Only needed if the price is
	// usage-based.
	BillableMetricID param.Field[string] `json:"billable_metric_id"`
	// If the Price represents a fixed cost, the price will be billed in-advance if
	// this is true, and in-arrears if this is false.
	BilledInAdvance param.Field[bool] `json:"billed_in_advance"`
	// For custom cadence: specifies the duration of the billing period in days or
	// months.
	BillingCycleConfiguration param.Field[NewBillingCycleConfigurationModelParam] `json:"billing_cycle_configuration"`
	// The per unit conversion rate of the price currency to the invoicing currency.
	ConversionRate param.Field[float64] `json:"conversion_rate"`
	// An alias for the price.
	ExternalPriceID param.Field[string] `json:"external_price_id"`
	// If the Price represents a fixed cost, this represents the quantity of units
	// applied.
	FixedPriceQuantity param.Field[float64] `json:"fixed_price_quantity"`
	// The property used to group this price on an invoice
	InvoiceGroupingKey param.Field[string] `json:"invoice_grouping_key"`
	// Within each billing cycle, specifies the cadence at which invoices are produced.
	// If unspecified, a single invoice is produced per billing cycle.
	InvoicingCycleConfiguration param.Field[NewBillingCycleConfigurationModelParam] `json:"invoicing_cycle_configuration"`
	// User-specified key/value pairs for the resource. Individual keys can be removed
	// by setting the value to `null`, and the entire metadata mapping can be cleared
	// by setting `metadata` to `null`.
	Metadata param.Field[map[string]string] `json:"metadata"`
}

func (r NewFloatingPriceModelNewFloatingPackageWithAllocationPriceParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r NewFloatingPriceModelNewFloatingPackageWithAllocationPriceParam) implementsNewFloatingPriceModelUnionParam() {
}

// The cadence to bill for this price on.
type NewFloatingPriceModelNewFloatingPackageWithAllocationPriceCadence string

const (
	NewFloatingPriceModelNewFloatingPackageWithAllocationPriceCadenceAnnual     NewFloatingPriceModelNewFloatingPackageWithAllocationPriceCadence = "annual"
	NewFloatingPriceModelNewFloatingPackageWithAllocationPriceCadenceSemiAnnual NewFloatingPriceModelNewFloatingPackageWithAllocationPriceCadence = "semi_annual"
	NewFloatingPriceModelNewFloatingPackageWithAllocationPriceCadenceMonthly    NewFloatingPriceModelNewFloatingPackageWithAllocationPriceCadence = "monthly"
	NewFloatingPriceModelNewFloatingPackageWithAllocationPriceCadenceQuarterly  NewFloatingPriceModelNewFloatingPackageWithAllocationPriceCadence = "quarterly"
	NewFloatingPriceModelNewFloatingPackageWithAllocationPriceCadenceOneTime    NewFloatingPriceModelNewFloatingPackageWithAllocationPriceCadence = "one_time"
	NewFloatingPriceModelNewFloatingPackageWithAllocationPriceCadenceCustom     NewFloatingPriceModelNewFloatingPackageWithAllocationPriceCadence = "custom"
)

func (r NewFloatingPriceModelNewFloatingPackageWithAllocationPriceCadence) IsKnown() bool {
	switch r {
	case NewFloatingPriceModelNewFloatingPackageWithAllocationPriceCadenceAnnual, NewFloatingPriceModelNewFloatingPackageWithAllocationPriceCadenceSemiAnnual, NewFloatingPriceModelNewFloatingPackageWithAllocationPriceCadenceMonthly, NewFloatingPriceModelNewFloatingPackageWithAllocationPriceCadenceQuarterly, NewFloatingPriceModelNewFloatingPackageWithAllocationPriceCadenceOneTime, NewFloatingPriceModelNewFloatingPackageWithAllocationPriceCadenceCustom:
		return true
	}
	return false
}

type NewFloatingPriceModelNewFloatingPackageWithAllocationPriceModelType string

const (
	NewFloatingPriceModelNewFloatingPackageWithAllocationPriceModelTypePackageWithAllocation NewFloatingPriceModelNewFloatingPackageWithAllocationPriceModelType = "package_with_allocation"
)

func (r NewFloatingPriceModelNewFloatingPackageWithAllocationPriceModelType) IsKnown() bool {
	switch r {
	case NewFloatingPriceModelNewFloatingPackageWithAllocationPriceModelTypePackageWithAllocation:
		return true
	}
	return false
}

type NewFloatingPriceModelNewFloatingTieredPackageWithMinimumPriceParam struct {
	// The cadence to bill for this price on.
	Cadence param.Field[NewFloatingPriceModelNewFloatingTieredPackageWithMinimumPriceCadence] `json:"cadence,required"`
	// An ISO 4217 currency string for which this price is billed in.
	Currency param.Field[string] `json:"currency,required"`
	// The id of the item the price will be associated with.
	ItemID    param.Field[string]                                                                 `json:"item_id,required"`
	ModelType param.Field[NewFloatingPriceModelNewFloatingTieredPackageWithMinimumPriceModelType] `json:"model_type,required"`
	// The name of the price.
	Name                           param.Field[string]                               `json:"name,required"`
	TieredPackageWithMinimumConfig param.Field[CustomRatingFunctionConfigModelParam] `json:"tiered_package_with_minimum_config,required"`
	// The id of the billable metric for the price. Only needed if the price is
	// usage-based.
	BillableMetricID param.Field[string] `json:"billable_metric_id"`
	// If the Price represents a fixed cost, the price will be billed in-advance if
	// this is true, and in-arrears if this is false.
	BilledInAdvance param.Field[bool] `json:"billed_in_advance"`
	// For custom cadence: specifies the duration of the billing period in days or
	// months.
	BillingCycleConfiguration param.Field[NewBillingCycleConfigurationModelParam] `json:"billing_cycle_configuration"`
	// The per unit conversion rate of the price currency to the invoicing currency.
	ConversionRate param.Field[float64] `json:"conversion_rate"`
	// An alias for the price.
	ExternalPriceID param.Field[string] `json:"external_price_id"`
	// If the Price represents a fixed cost, this represents the quantity of units
	// applied.
	FixedPriceQuantity param.Field[float64] `json:"fixed_price_quantity"`
	// The property used to group this price on an invoice
	InvoiceGroupingKey param.Field[string] `json:"invoice_grouping_key"`
	// Within each billing cycle, specifies the cadence at which invoices are produced.
	// If unspecified, a single invoice is produced per billing cycle.
	InvoicingCycleConfiguration param.Field[NewBillingCycleConfigurationModelParam] `json:"invoicing_cycle_configuration"`
	// User-specified key/value pairs for the resource. Individual keys can be removed
	// by setting the value to `null`, and the entire metadata mapping can be cleared
	// by setting `metadata` to `null`.
	Metadata param.Field[map[string]string] `json:"metadata"`
}

func (r NewFloatingPriceModelNewFloatingTieredPackageWithMinimumPriceParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r NewFloatingPriceModelNewFloatingTieredPackageWithMinimumPriceParam) implementsNewFloatingPriceModelUnionParam() {
}

// The cadence to bill for this price on.
type NewFloatingPriceModelNewFloatingTieredPackageWithMinimumPriceCadence string

const (
	NewFloatingPriceModelNewFloatingTieredPackageWithMinimumPriceCadenceAnnual     NewFloatingPriceModelNewFloatingTieredPackageWithMinimumPriceCadence = "annual"
	NewFloatingPriceModelNewFloatingTieredPackageWithMinimumPriceCadenceSemiAnnual NewFloatingPriceModelNewFloatingTieredPackageWithMinimumPriceCadence = "semi_annual"
	NewFloatingPriceModelNewFloatingTieredPackageWithMinimumPriceCadenceMonthly    NewFloatingPriceModelNewFloatingTieredPackageWithMinimumPriceCadence = "monthly"
	NewFloatingPriceModelNewFloatingTieredPackageWithMinimumPriceCadenceQuarterly  NewFloatingPriceModelNewFloatingTieredPackageWithMinimumPriceCadence = "quarterly"
	NewFloatingPriceModelNewFloatingTieredPackageWithMinimumPriceCadenceOneTime    NewFloatingPriceModelNewFloatingTieredPackageWithMinimumPriceCadence = "one_time"
	NewFloatingPriceModelNewFloatingTieredPackageWithMinimumPriceCadenceCustom     NewFloatingPriceModelNewFloatingTieredPackageWithMinimumPriceCadence = "custom"
)

func (r NewFloatingPriceModelNewFloatingTieredPackageWithMinimumPriceCadence) IsKnown() bool {
	switch r {
	case NewFloatingPriceModelNewFloatingTieredPackageWithMinimumPriceCadenceAnnual, NewFloatingPriceModelNewFloatingTieredPackageWithMinimumPriceCadenceSemiAnnual, NewFloatingPriceModelNewFloatingTieredPackageWithMinimumPriceCadenceMonthly, NewFloatingPriceModelNewFloatingTieredPackageWithMinimumPriceCadenceQuarterly, NewFloatingPriceModelNewFloatingTieredPackageWithMinimumPriceCadenceOneTime, NewFloatingPriceModelNewFloatingTieredPackageWithMinimumPriceCadenceCustom:
		return true
	}
	return false
}

type NewFloatingPriceModelNewFloatingTieredPackageWithMinimumPriceModelType string

const (
	NewFloatingPriceModelNewFloatingTieredPackageWithMinimumPriceModelTypeTieredPackageWithMinimum NewFloatingPriceModelNewFloatingTieredPackageWithMinimumPriceModelType = "tiered_package_with_minimum"
)

func (r NewFloatingPriceModelNewFloatingTieredPackageWithMinimumPriceModelType) IsKnown() bool {
	switch r {
	case NewFloatingPriceModelNewFloatingTieredPackageWithMinimumPriceModelTypeTieredPackageWithMinimum:
		return true
	}
	return false
}

type NewFloatingPriceModelNewFloatingUnitWithPercentPriceParam struct {
	// The cadence to bill for this price on.
	Cadence param.Field[NewFloatingPriceModelNewFloatingUnitWithPercentPriceCadence] `json:"cadence,required"`
	// An ISO 4217 currency string for which this price is billed in.
	Currency param.Field[string] `json:"currency,required"`
	// The id of the item the price will be associated with.
	ItemID    param.Field[string]                                                        `json:"item_id,required"`
	ModelType param.Field[NewFloatingPriceModelNewFloatingUnitWithPercentPriceModelType] `json:"model_type,required"`
	// The name of the price.
	Name                  param.Field[string]                               `json:"name,required"`
	UnitWithPercentConfig param.Field[CustomRatingFunctionConfigModelParam] `json:"unit_with_percent_config,required"`
	// The id of the billable metric for the price. Only needed if the price is
	// usage-based.
	BillableMetricID param.Field[string] `json:"billable_metric_id"`
	// If the Price represents a fixed cost, the price will be billed in-advance if
	// this is true, and in-arrears if this is false.
	BilledInAdvance param.Field[bool] `json:"billed_in_advance"`
	// For custom cadence: specifies the duration of the billing period in days or
	// months.
	BillingCycleConfiguration param.Field[NewBillingCycleConfigurationModelParam] `json:"billing_cycle_configuration"`
	// The per unit conversion rate of the price currency to the invoicing currency.
	ConversionRate param.Field[float64] `json:"conversion_rate"`
	// An alias for the price.
	ExternalPriceID param.Field[string] `json:"external_price_id"`
	// If the Price represents a fixed cost, this represents the quantity of units
	// applied.
	FixedPriceQuantity param.Field[float64] `json:"fixed_price_quantity"`
	// The property used to group this price on an invoice
	InvoiceGroupingKey param.Field[string] `json:"invoice_grouping_key"`
	// Within each billing cycle, specifies the cadence at which invoices are produced.
	// If unspecified, a single invoice is produced per billing cycle.
	InvoicingCycleConfiguration param.Field[NewBillingCycleConfigurationModelParam] `json:"invoicing_cycle_configuration"`
	// User-specified key/value pairs for the resource. Individual keys can be removed
	// by setting the value to `null`, and the entire metadata mapping can be cleared
	// by setting `metadata` to `null`.
	Metadata param.Field[map[string]string] `json:"metadata"`
}

func (r NewFloatingPriceModelNewFloatingUnitWithPercentPriceParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r NewFloatingPriceModelNewFloatingUnitWithPercentPriceParam) implementsNewFloatingPriceModelUnionParam() {
}

// The cadence to bill for this price on.
type NewFloatingPriceModelNewFloatingUnitWithPercentPriceCadence string

const (
	NewFloatingPriceModelNewFloatingUnitWithPercentPriceCadenceAnnual     NewFloatingPriceModelNewFloatingUnitWithPercentPriceCadence = "annual"
	NewFloatingPriceModelNewFloatingUnitWithPercentPriceCadenceSemiAnnual NewFloatingPriceModelNewFloatingUnitWithPercentPriceCadence = "semi_annual"
	NewFloatingPriceModelNewFloatingUnitWithPercentPriceCadenceMonthly    NewFloatingPriceModelNewFloatingUnitWithPercentPriceCadence = "monthly"
	NewFloatingPriceModelNewFloatingUnitWithPercentPriceCadenceQuarterly  NewFloatingPriceModelNewFloatingUnitWithPercentPriceCadence = "quarterly"
	NewFloatingPriceModelNewFloatingUnitWithPercentPriceCadenceOneTime    NewFloatingPriceModelNewFloatingUnitWithPercentPriceCadence = "one_time"
	NewFloatingPriceModelNewFloatingUnitWithPercentPriceCadenceCustom     NewFloatingPriceModelNewFloatingUnitWithPercentPriceCadence = "custom"
)

func (r NewFloatingPriceModelNewFloatingUnitWithPercentPriceCadence) IsKnown() bool {
	switch r {
	case NewFloatingPriceModelNewFloatingUnitWithPercentPriceCadenceAnnual, NewFloatingPriceModelNewFloatingUnitWithPercentPriceCadenceSemiAnnual, NewFloatingPriceModelNewFloatingUnitWithPercentPriceCadenceMonthly, NewFloatingPriceModelNewFloatingUnitWithPercentPriceCadenceQuarterly, NewFloatingPriceModelNewFloatingUnitWithPercentPriceCadenceOneTime, NewFloatingPriceModelNewFloatingUnitWithPercentPriceCadenceCustom:
		return true
	}
	return false
}

type NewFloatingPriceModelNewFloatingUnitWithPercentPriceModelType string

const (
	NewFloatingPriceModelNewFloatingUnitWithPercentPriceModelTypeUnitWithPercent NewFloatingPriceModelNewFloatingUnitWithPercentPriceModelType = "unit_with_percent"
)

func (r NewFloatingPriceModelNewFloatingUnitWithPercentPriceModelType) IsKnown() bool {
	switch r {
	case NewFloatingPriceModelNewFloatingUnitWithPercentPriceModelTypeUnitWithPercent:
		return true
	}
	return false
}

type NewFloatingPriceModelNewFloatingTieredWithProrationPriceParam struct {
	// The cadence to bill for this price on.
	Cadence param.Field[NewFloatingPriceModelNewFloatingTieredWithProrationPriceCadence] `json:"cadence,required"`
	// An ISO 4217 currency string for which this price is billed in.
	Currency param.Field[string] `json:"currency,required"`
	// The id of the item the price will be associated with.
	ItemID    param.Field[string]                                                            `json:"item_id,required"`
	ModelType param.Field[NewFloatingPriceModelNewFloatingTieredWithProrationPriceModelType] `json:"model_type,required"`
	// The name of the price.
	Name                      param.Field[string]                               `json:"name,required"`
	TieredWithProrationConfig param.Field[CustomRatingFunctionConfigModelParam] `json:"tiered_with_proration_config,required"`
	// The id of the billable metric for the price. Only needed if the price is
	// usage-based.
	BillableMetricID param.Field[string] `json:"billable_metric_id"`
	// If the Price represents a fixed cost, the price will be billed in-advance if
	// this is true, and in-arrears if this is false.
	BilledInAdvance param.Field[bool] `json:"billed_in_advance"`
	// For custom cadence: specifies the duration of the billing period in days or
	// months.
	BillingCycleConfiguration param.Field[NewBillingCycleConfigurationModelParam] `json:"billing_cycle_configuration"`
	// The per unit conversion rate of the price currency to the invoicing currency.
	ConversionRate param.Field[float64] `json:"conversion_rate"`
	// An alias for the price.
	ExternalPriceID param.Field[string] `json:"external_price_id"`
	// If the Price represents a fixed cost, this represents the quantity of units
	// applied.
	FixedPriceQuantity param.Field[float64] `json:"fixed_price_quantity"`
	// The property used to group this price on an invoice
	InvoiceGroupingKey param.Field[string] `json:"invoice_grouping_key"`
	// Within each billing cycle, specifies the cadence at which invoices are produced.
	// If unspecified, a single invoice is produced per billing cycle.
	InvoicingCycleConfiguration param.Field[NewBillingCycleConfigurationModelParam] `json:"invoicing_cycle_configuration"`
	// User-specified key/value pairs for the resource. Individual keys can be removed
	// by setting the value to `null`, and the entire metadata mapping can be cleared
	// by setting `metadata` to `null`.
	Metadata param.Field[map[string]string] `json:"metadata"`
}

func (r NewFloatingPriceModelNewFloatingTieredWithProrationPriceParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r NewFloatingPriceModelNewFloatingTieredWithProrationPriceParam) implementsNewFloatingPriceModelUnionParam() {
}

// The cadence to bill for this price on.
type NewFloatingPriceModelNewFloatingTieredWithProrationPriceCadence string

const (
	NewFloatingPriceModelNewFloatingTieredWithProrationPriceCadenceAnnual     NewFloatingPriceModelNewFloatingTieredWithProrationPriceCadence = "annual"
	NewFloatingPriceModelNewFloatingTieredWithProrationPriceCadenceSemiAnnual NewFloatingPriceModelNewFloatingTieredWithProrationPriceCadence = "semi_annual"
	NewFloatingPriceModelNewFloatingTieredWithProrationPriceCadenceMonthly    NewFloatingPriceModelNewFloatingTieredWithProrationPriceCadence = "monthly"
	NewFloatingPriceModelNewFloatingTieredWithProrationPriceCadenceQuarterly  NewFloatingPriceModelNewFloatingTieredWithProrationPriceCadence = "quarterly"
	NewFloatingPriceModelNewFloatingTieredWithProrationPriceCadenceOneTime    NewFloatingPriceModelNewFloatingTieredWithProrationPriceCadence = "one_time"
	NewFloatingPriceModelNewFloatingTieredWithProrationPriceCadenceCustom     NewFloatingPriceModelNewFloatingTieredWithProrationPriceCadence = "custom"
)

func (r NewFloatingPriceModelNewFloatingTieredWithProrationPriceCadence) IsKnown() bool {
	switch r {
	case NewFloatingPriceModelNewFloatingTieredWithProrationPriceCadenceAnnual, NewFloatingPriceModelNewFloatingTieredWithProrationPriceCadenceSemiAnnual, NewFloatingPriceModelNewFloatingTieredWithProrationPriceCadenceMonthly, NewFloatingPriceModelNewFloatingTieredWithProrationPriceCadenceQuarterly, NewFloatingPriceModelNewFloatingTieredWithProrationPriceCadenceOneTime, NewFloatingPriceModelNewFloatingTieredWithProrationPriceCadenceCustom:
		return true
	}
	return false
}

type NewFloatingPriceModelNewFloatingTieredWithProrationPriceModelType string

const (
	NewFloatingPriceModelNewFloatingTieredWithProrationPriceModelTypeTieredWithProration NewFloatingPriceModelNewFloatingTieredWithProrationPriceModelType = "tiered_with_proration"
)

func (r NewFloatingPriceModelNewFloatingTieredWithProrationPriceModelType) IsKnown() bool {
	switch r {
	case NewFloatingPriceModelNewFloatingTieredWithProrationPriceModelTypeTieredWithProration:
		return true
	}
	return false
}

type NewFloatingPriceModelNewFloatingUnitWithProrationPriceParam struct {
	// The cadence to bill for this price on.
	Cadence param.Field[NewFloatingPriceModelNewFloatingUnitWithProrationPriceCadence] `json:"cadence,required"`
	// An ISO 4217 currency string for which this price is billed in.
	Currency param.Field[string] `json:"currency,required"`
	// The id of the item the price will be associated with.
	ItemID    param.Field[string]                                                          `json:"item_id,required"`
	ModelType param.Field[NewFloatingPriceModelNewFloatingUnitWithProrationPriceModelType] `json:"model_type,required"`
	// The name of the price.
	Name                    param.Field[string]                               `json:"name,required"`
	UnitWithProrationConfig param.Field[CustomRatingFunctionConfigModelParam] `json:"unit_with_proration_config,required"`
	// The id of the billable metric for the price. Only needed if the price is
	// usage-based.
	BillableMetricID param.Field[string] `json:"billable_metric_id"`
	// If the Price represents a fixed cost, the price will be billed in-advance if
	// this is true, and in-arrears if this is false.
	BilledInAdvance param.Field[bool] `json:"billed_in_advance"`
	// For custom cadence: specifies the duration of the billing period in days or
	// months.
	BillingCycleConfiguration param.Field[NewBillingCycleConfigurationModelParam] `json:"billing_cycle_configuration"`
	// The per unit conversion rate of the price currency to the invoicing currency.
	ConversionRate param.Field[float64] `json:"conversion_rate"`
	// An alias for the price.
	ExternalPriceID param.Field[string] `json:"external_price_id"`
	// If the Price represents a fixed cost, this represents the quantity of units
	// applied.
	FixedPriceQuantity param.Field[float64] `json:"fixed_price_quantity"`
	// The property used to group this price on an invoice
	InvoiceGroupingKey param.Field[string] `json:"invoice_grouping_key"`
	// Within each billing cycle, specifies the cadence at which invoices are produced.
	// If unspecified, a single invoice is produced per billing cycle.
	InvoicingCycleConfiguration param.Field[NewBillingCycleConfigurationModelParam] `json:"invoicing_cycle_configuration"`
	// User-specified key/value pairs for the resource. Individual keys can be removed
	// by setting the value to `null`, and the entire metadata mapping can be cleared
	// by setting `metadata` to `null`.
	Metadata param.Field[map[string]string] `json:"metadata"`
}

func (r NewFloatingPriceModelNewFloatingUnitWithProrationPriceParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r NewFloatingPriceModelNewFloatingUnitWithProrationPriceParam) implementsNewFloatingPriceModelUnionParam() {
}

// The cadence to bill for this price on.
type NewFloatingPriceModelNewFloatingUnitWithProrationPriceCadence string

const (
	NewFloatingPriceModelNewFloatingUnitWithProrationPriceCadenceAnnual     NewFloatingPriceModelNewFloatingUnitWithProrationPriceCadence = "annual"
	NewFloatingPriceModelNewFloatingUnitWithProrationPriceCadenceSemiAnnual NewFloatingPriceModelNewFloatingUnitWithProrationPriceCadence = "semi_annual"
	NewFloatingPriceModelNewFloatingUnitWithProrationPriceCadenceMonthly    NewFloatingPriceModelNewFloatingUnitWithProrationPriceCadence = "monthly"
	NewFloatingPriceModelNewFloatingUnitWithProrationPriceCadenceQuarterly  NewFloatingPriceModelNewFloatingUnitWithProrationPriceCadence = "quarterly"
	NewFloatingPriceModelNewFloatingUnitWithProrationPriceCadenceOneTime    NewFloatingPriceModelNewFloatingUnitWithProrationPriceCadence = "one_time"
	NewFloatingPriceModelNewFloatingUnitWithProrationPriceCadenceCustom     NewFloatingPriceModelNewFloatingUnitWithProrationPriceCadence = "custom"
)

func (r NewFloatingPriceModelNewFloatingUnitWithProrationPriceCadence) IsKnown() bool {
	switch r {
	case NewFloatingPriceModelNewFloatingUnitWithProrationPriceCadenceAnnual, NewFloatingPriceModelNewFloatingUnitWithProrationPriceCadenceSemiAnnual, NewFloatingPriceModelNewFloatingUnitWithProrationPriceCadenceMonthly, NewFloatingPriceModelNewFloatingUnitWithProrationPriceCadenceQuarterly, NewFloatingPriceModelNewFloatingUnitWithProrationPriceCadenceOneTime, NewFloatingPriceModelNewFloatingUnitWithProrationPriceCadenceCustom:
		return true
	}
	return false
}

type NewFloatingPriceModelNewFloatingUnitWithProrationPriceModelType string

const (
	NewFloatingPriceModelNewFloatingUnitWithProrationPriceModelTypeUnitWithProration NewFloatingPriceModelNewFloatingUnitWithProrationPriceModelType = "unit_with_proration"
)

func (r NewFloatingPriceModelNewFloatingUnitWithProrationPriceModelType) IsKnown() bool {
	switch r {
	case NewFloatingPriceModelNewFloatingUnitWithProrationPriceModelTypeUnitWithProration:
		return true
	}
	return false
}

type NewFloatingPriceModelNewFloatingGroupedAllocationPriceParam struct {
	// The cadence to bill for this price on.
	Cadence param.Field[NewFloatingPriceModelNewFloatingGroupedAllocationPriceCadence] `json:"cadence,required"`
	// An ISO 4217 currency string for which this price is billed in.
	Currency                param.Field[string]                               `json:"currency,required"`
	GroupedAllocationConfig param.Field[CustomRatingFunctionConfigModelParam] `json:"grouped_allocation_config,required"`
	// The id of the item the price will be associated with.
	ItemID    param.Field[string]                                                          `json:"item_id,required"`
	ModelType param.Field[NewFloatingPriceModelNewFloatingGroupedAllocationPriceModelType] `json:"model_type,required"`
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
	BillingCycleConfiguration param.Field[NewBillingCycleConfigurationModelParam] `json:"billing_cycle_configuration"`
	// The per unit conversion rate of the price currency to the invoicing currency.
	ConversionRate param.Field[float64] `json:"conversion_rate"`
	// An alias for the price.
	ExternalPriceID param.Field[string] `json:"external_price_id"`
	// If the Price represents a fixed cost, this represents the quantity of units
	// applied.
	FixedPriceQuantity param.Field[float64] `json:"fixed_price_quantity"`
	// The property used to group this price on an invoice
	InvoiceGroupingKey param.Field[string] `json:"invoice_grouping_key"`
	// Within each billing cycle, specifies the cadence at which invoices are produced.
	// If unspecified, a single invoice is produced per billing cycle.
	InvoicingCycleConfiguration param.Field[NewBillingCycleConfigurationModelParam] `json:"invoicing_cycle_configuration"`
	// User-specified key/value pairs for the resource. Individual keys can be removed
	// by setting the value to `null`, and the entire metadata mapping can be cleared
	// by setting `metadata` to `null`.
	Metadata param.Field[map[string]string] `json:"metadata"`
}

func (r NewFloatingPriceModelNewFloatingGroupedAllocationPriceParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r NewFloatingPriceModelNewFloatingGroupedAllocationPriceParam) implementsNewFloatingPriceModelUnionParam() {
}

// The cadence to bill for this price on.
type NewFloatingPriceModelNewFloatingGroupedAllocationPriceCadence string

const (
	NewFloatingPriceModelNewFloatingGroupedAllocationPriceCadenceAnnual     NewFloatingPriceModelNewFloatingGroupedAllocationPriceCadence = "annual"
	NewFloatingPriceModelNewFloatingGroupedAllocationPriceCadenceSemiAnnual NewFloatingPriceModelNewFloatingGroupedAllocationPriceCadence = "semi_annual"
	NewFloatingPriceModelNewFloatingGroupedAllocationPriceCadenceMonthly    NewFloatingPriceModelNewFloatingGroupedAllocationPriceCadence = "monthly"
	NewFloatingPriceModelNewFloatingGroupedAllocationPriceCadenceQuarterly  NewFloatingPriceModelNewFloatingGroupedAllocationPriceCadence = "quarterly"
	NewFloatingPriceModelNewFloatingGroupedAllocationPriceCadenceOneTime    NewFloatingPriceModelNewFloatingGroupedAllocationPriceCadence = "one_time"
	NewFloatingPriceModelNewFloatingGroupedAllocationPriceCadenceCustom     NewFloatingPriceModelNewFloatingGroupedAllocationPriceCadence = "custom"
)

func (r NewFloatingPriceModelNewFloatingGroupedAllocationPriceCadence) IsKnown() bool {
	switch r {
	case NewFloatingPriceModelNewFloatingGroupedAllocationPriceCadenceAnnual, NewFloatingPriceModelNewFloatingGroupedAllocationPriceCadenceSemiAnnual, NewFloatingPriceModelNewFloatingGroupedAllocationPriceCadenceMonthly, NewFloatingPriceModelNewFloatingGroupedAllocationPriceCadenceQuarterly, NewFloatingPriceModelNewFloatingGroupedAllocationPriceCadenceOneTime, NewFloatingPriceModelNewFloatingGroupedAllocationPriceCadenceCustom:
		return true
	}
	return false
}

type NewFloatingPriceModelNewFloatingGroupedAllocationPriceModelType string

const (
	NewFloatingPriceModelNewFloatingGroupedAllocationPriceModelTypeGroupedAllocation NewFloatingPriceModelNewFloatingGroupedAllocationPriceModelType = "grouped_allocation"
)

func (r NewFloatingPriceModelNewFloatingGroupedAllocationPriceModelType) IsKnown() bool {
	switch r {
	case NewFloatingPriceModelNewFloatingGroupedAllocationPriceModelTypeGroupedAllocation:
		return true
	}
	return false
}

type NewFloatingPriceModelNewFloatingGroupedWithProratedMinimumPriceParam struct {
	// The cadence to bill for this price on.
	Cadence param.Field[NewFloatingPriceModelNewFloatingGroupedWithProratedMinimumPriceCadence] `json:"cadence,required"`
	// An ISO 4217 currency string for which this price is billed in.
	Currency                         param.Field[string]                               `json:"currency,required"`
	GroupedWithProratedMinimumConfig param.Field[CustomRatingFunctionConfigModelParam] `json:"grouped_with_prorated_minimum_config,required"`
	// The id of the item the price will be associated with.
	ItemID    param.Field[string]                                                                   `json:"item_id,required"`
	ModelType param.Field[NewFloatingPriceModelNewFloatingGroupedWithProratedMinimumPriceModelType] `json:"model_type,required"`
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
	BillingCycleConfiguration param.Field[NewBillingCycleConfigurationModelParam] `json:"billing_cycle_configuration"`
	// The per unit conversion rate of the price currency to the invoicing currency.
	ConversionRate param.Field[float64] `json:"conversion_rate"`
	// An alias for the price.
	ExternalPriceID param.Field[string] `json:"external_price_id"`
	// If the Price represents a fixed cost, this represents the quantity of units
	// applied.
	FixedPriceQuantity param.Field[float64] `json:"fixed_price_quantity"`
	// The property used to group this price on an invoice
	InvoiceGroupingKey param.Field[string] `json:"invoice_grouping_key"`
	// Within each billing cycle, specifies the cadence at which invoices are produced.
	// If unspecified, a single invoice is produced per billing cycle.
	InvoicingCycleConfiguration param.Field[NewBillingCycleConfigurationModelParam] `json:"invoicing_cycle_configuration"`
	// User-specified key/value pairs for the resource. Individual keys can be removed
	// by setting the value to `null`, and the entire metadata mapping can be cleared
	// by setting `metadata` to `null`.
	Metadata param.Field[map[string]string] `json:"metadata"`
}

func (r NewFloatingPriceModelNewFloatingGroupedWithProratedMinimumPriceParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r NewFloatingPriceModelNewFloatingGroupedWithProratedMinimumPriceParam) implementsNewFloatingPriceModelUnionParam() {
}

// The cadence to bill for this price on.
type NewFloatingPriceModelNewFloatingGroupedWithProratedMinimumPriceCadence string

const (
	NewFloatingPriceModelNewFloatingGroupedWithProratedMinimumPriceCadenceAnnual     NewFloatingPriceModelNewFloatingGroupedWithProratedMinimumPriceCadence = "annual"
	NewFloatingPriceModelNewFloatingGroupedWithProratedMinimumPriceCadenceSemiAnnual NewFloatingPriceModelNewFloatingGroupedWithProratedMinimumPriceCadence = "semi_annual"
	NewFloatingPriceModelNewFloatingGroupedWithProratedMinimumPriceCadenceMonthly    NewFloatingPriceModelNewFloatingGroupedWithProratedMinimumPriceCadence = "monthly"
	NewFloatingPriceModelNewFloatingGroupedWithProratedMinimumPriceCadenceQuarterly  NewFloatingPriceModelNewFloatingGroupedWithProratedMinimumPriceCadence = "quarterly"
	NewFloatingPriceModelNewFloatingGroupedWithProratedMinimumPriceCadenceOneTime    NewFloatingPriceModelNewFloatingGroupedWithProratedMinimumPriceCadence = "one_time"
	NewFloatingPriceModelNewFloatingGroupedWithProratedMinimumPriceCadenceCustom     NewFloatingPriceModelNewFloatingGroupedWithProratedMinimumPriceCadence = "custom"
)

func (r NewFloatingPriceModelNewFloatingGroupedWithProratedMinimumPriceCadence) IsKnown() bool {
	switch r {
	case NewFloatingPriceModelNewFloatingGroupedWithProratedMinimumPriceCadenceAnnual, NewFloatingPriceModelNewFloatingGroupedWithProratedMinimumPriceCadenceSemiAnnual, NewFloatingPriceModelNewFloatingGroupedWithProratedMinimumPriceCadenceMonthly, NewFloatingPriceModelNewFloatingGroupedWithProratedMinimumPriceCadenceQuarterly, NewFloatingPriceModelNewFloatingGroupedWithProratedMinimumPriceCadenceOneTime, NewFloatingPriceModelNewFloatingGroupedWithProratedMinimumPriceCadenceCustom:
		return true
	}
	return false
}

type NewFloatingPriceModelNewFloatingGroupedWithProratedMinimumPriceModelType string

const (
	NewFloatingPriceModelNewFloatingGroupedWithProratedMinimumPriceModelTypeGroupedWithProratedMinimum NewFloatingPriceModelNewFloatingGroupedWithProratedMinimumPriceModelType = "grouped_with_prorated_minimum"
)

func (r NewFloatingPriceModelNewFloatingGroupedWithProratedMinimumPriceModelType) IsKnown() bool {
	switch r {
	case NewFloatingPriceModelNewFloatingGroupedWithProratedMinimumPriceModelTypeGroupedWithProratedMinimum:
		return true
	}
	return false
}

type NewFloatingPriceModelNewFloatingGroupedWithMeteredMinimumPriceParam struct {
	// The cadence to bill for this price on.
	Cadence param.Field[NewFloatingPriceModelNewFloatingGroupedWithMeteredMinimumPriceCadence] `json:"cadence,required"`
	// An ISO 4217 currency string for which this price is billed in.
	Currency                        param.Field[string]                               `json:"currency,required"`
	GroupedWithMeteredMinimumConfig param.Field[CustomRatingFunctionConfigModelParam] `json:"grouped_with_metered_minimum_config,required"`
	// The id of the item the price will be associated with.
	ItemID    param.Field[string]                                                                  `json:"item_id,required"`
	ModelType param.Field[NewFloatingPriceModelNewFloatingGroupedWithMeteredMinimumPriceModelType] `json:"model_type,required"`
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
	BillingCycleConfiguration param.Field[NewBillingCycleConfigurationModelParam] `json:"billing_cycle_configuration"`
	// The per unit conversion rate of the price currency to the invoicing currency.
	ConversionRate param.Field[float64] `json:"conversion_rate"`
	// An alias for the price.
	ExternalPriceID param.Field[string] `json:"external_price_id"`
	// If the Price represents a fixed cost, this represents the quantity of units
	// applied.
	FixedPriceQuantity param.Field[float64] `json:"fixed_price_quantity"`
	// The property used to group this price on an invoice
	InvoiceGroupingKey param.Field[string] `json:"invoice_grouping_key"`
	// Within each billing cycle, specifies the cadence at which invoices are produced.
	// If unspecified, a single invoice is produced per billing cycle.
	InvoicingCycleConfiguration param.Field[NewBillingCycleConfigurationModelParam] `json:"invoicing_cycle_configuration"`
	// User-specified key/value pairs for the resource. Individual keys can be removed
	// by setting the value to `null`, and the entire metadata mapping can be cleared
	// by setting `metadata` to `null`.
	Metadata param.Field[map[string]string] `json:"metadata"`
}

func (r NewFloatingPriceModelNewFloatingGroupedWithMeteredMinimumPriceParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r NewFloatingPriceModelNewFloatingGroupedWithMeteredMinimumPriceParam) implementsNewFloatingPriceModelUnionParam() {
}

// The cadence to bill for this price on.
type NewFloatingPriceModelNewFloatingGroupedWithMeteredMinimumPriceCadence string

const (
	NewFloatingPriceModelNewFloatingGroupedWithMeteredMinimumPriceCadenceAnnual     NewFloatingPriceModelNewFloatingGroupedWithMeteredMinimumPriceCadence = "annual"
	NewFloatingPriceModelNewFloatingGroupedWithMeteredMinimumPriceCadenceSemiAnnual NewFloatingPriceModelNewFloatingGroupedWithMeteredMinimumPriceCadence = "semi_annual"
	NewFloatingPriceModelNewFloatingGroupedWithMeteredMinimumPriceCadenceMonthly    NewFloatingPriceModelNewFloatingGroupedWithMeteredMinimumPriceCadence = "monthly"
	NewFloatingPriceModelNewFloatingGroupedWithMeteredMinimumPriceCadenceQuarterly  NewFloatingPriceModelNewFloatingGroupedWithMeteredMinimumPriceCadence = "quarterly"
	NewFloatingPriceModelNewFloatingGroupedWithMeteredMinimumPriceCadenceOneTime    NewFloatingPriceModelNewFloatingGroupedWithMeteredMinimumPriceCadence = "one_time"
	NewFloatingPriceModelNewFloatingGroupedWithMeteredMinimumPriceCadenceCustom     NewFloatingPriceModelNewFloatingGroupedWithMeteredMinimumPriceCadence = "custom"
)

func (r NewFloatingPriceModelNewFloatingGroupedWithMeteredMinimumPriceCadence) IsKnown() bool {
	switch r {
	case NewFloatingPriceModelNewFloatingGroupedWithMeteredMinimumPriceCadenceAnnual, NewFloatingPriceModelNewFloatingGroupedWithMeteredMinimumPriceCadenceSemiAnnual, NewFloatingPriceModelNewFloatingGroupedWithMeteredMinimumPriceCadenceMonthly, NewFloatingPriceModelNewFloatingGroupedWithMeteredMinimumPriceCadenceQuarterly, NewFloatingPriceModelNewFloatingGroupedWithMeteredMinimumPriceCadenceOneTime, NewFloatingPriceModelNewFloatingGroupedWithMeteredMinimumPriceCadenceCustom:
		return true
	}
	return false
}

type NewFloatingPriceModelNewFloatingGroupedWithMeteredMinimumPriceModelType string

const (
	NewFloatingPriceModelNewFloatingGroupedWithMeteredMinimumPriceModelTypeGroupedWithMeteredMinimum NewFloatingPriceModelNewFloatingGroupedWithMeteredMinimumPriceModelType = "grouped_with_metered_minimum"
)

func (r NewFloatingPriceModelNewFloatingGroupedWithMeteredMinimumPriceModelType) IsKnown() bool {
	switch r {
	case NewFloatingPriceModelNewFloatingGroupedWithMeteredMinimumPriceModelTypeGroupedWithMeteredMinimum:
		return true
	}
	return false
}

type NewFloatingPriceModelNewFloatingMatrixWithDisplayNamePriceParam struct {
	// The cadence to bill for this price on.
	Cadence param.Field[NewFloatingPriceModelNewFloatingMatrixWithDisplayNamePriceCadence] `json:"cadence,required"`
	// An ISO 4217 currency string for which this price is billed in.
	Currency param.Field[string] `json:"currency,required"`
	// The id of the item the price will be associated with.
	ItemID                      param.Field[string]                                                              `json:"item_id,required"`
	MatrixWithDisplayNameConfig param.Field[CustomRatingFunctionConfigModelParam]                                `json:"matrix_with_display_name_config,required"`
	ModelType                   param.Field[NewFloatingPriceModelNewFloatingMatrixWithDisplayNamePriceModelType] `json:"model_type,required"`
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
	BillingCycleConfiguration param.Field[NewBillingCycleConfigurationModelParam] `json:"billing_cycle_configuration"`
	// The per unit conversion rate of the price currency to the invoicing currency.
	ConversionRate param.Field[float64] `json:"conversion_rate"`
	// An alias for the price.
	ExternalPriceID param.Field[string] `json:"external_price_id"`
	// If the Price represents a fixed cost, this represents the quantity of units
	// applied.
	FixedPriceQuantity param.Field[float64] `json:"fixed_price_quantity"`
	// The property used to group this price on an invoice
	InvoiceGroupingKey param.Field[string] `json:"invoice_grouping_key"`
	// Within each billing cycle, specifies the cadence at which invoices are produced.
	// If unspecified, a single invoice is produced per billing cycle.
	InvoicingCycleConfiguration param.Field[NewBillingCycleConfigurationModelParam] `json:"invoicing_cycle_configuration"`
	// User-specified key/value pairs for the resource. Individual keys can be removed
	// by setting the value to `null`, and the entire metadata mapping can be cleared
	// by setting `metadata` to `null`.
	Metadata param.Field[map[string]string] `json:"metadata"`
}

func (r NewFloatingPriceModelNewFloatingMatrixWithDisplayNamePriceParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r NewFloatingPriceModelNewFloatingMatrixWithDisplayNamePriceParam) implementsNewFloatingPriceModelUnionParam() {
}

// The cadence to bill for this price on.
type NewFloatingPriceModelNewFloatingMatrixWithDisplayNamePriceCadence string

const (
	NewFloatingPriceModelNewFloatingMatrixWithDisplayNamePriceCadenceAnnual     NewFloatingPriceModelNewFloatingMatrixWithDisplayNamePriceCadence = "annual"
	NewFloatingPriceModelNewFloatingMatrixWithDisplayNamePriceCadenceSemiAnnual NewFloatingPriceModelNewFloatingMatrixWithDisplayNamePriceCadence = "semi_annual"
	NewFloatingPriceModelNewFloatingMatrixWithDisplayNamePriceCadenceMonthly    NewFloatingPriceModelNewFloatingMatrixWithDisplayNamePriceCadence = "monthly"
	NewFloatingPriceModelNewFloatingMatrixWithDisplayNamePriceCadenceQuarterly  NewFloatingPriceModelNewFloatingMatrixWithDisplayNamePriceCadence = "quarterly"
	NewFloatingPriceModelNewFloatingMatrixWithDisplayNamePriceCadenceOneTime    NewFloatingPriceModelNewFloatingMatrixWithDisplayNamePriceCadence = "one_time"
	NewFloatingPriceModelNewFloatingMatrixWithDisplayNamePriceCadenceCustom     NewFloatingPriceModelNewFloatingMatrixWithDisplayNamePriceCadence = "custom"
)

func (r NewFloatingPriceModelNewFloatingMatrixWithDisplayNamePriceCadence) IsKnown() bool {
	switch r {
	case NewFloatingPriceModelNewFloatingMatrixWithDisplayNamePriceCadenceAnnual, NewFloatingPriceModelNewFloatingMatrixWithDisplayNamePriceCadenceSemiAnnual, NewFloatingPriceModelNewFloatingMatrixWithDisplayNamePriceCadenceMonthly, NewFloatingPriceModelNewFloatingMatrixWithDisplayNamePriceCadenceQuarterly, NewFloatingPriceModelNewFloatingMatrixWithDisplayNamePriceCadenceOneTime, NewFloatingPriceModelNewFloatingMatrixWithDisplayNamePriceCadenceCustom:
		return true
	}
	return false
}

type NewFloatingPriceModelNewFloatingMatrixWithDisplayNamePriceModelType string

const (
	NewFloatingPriceModelNewFloatingMatrixWithDisplayNamePriceModelTypeMatrixWithDisplayName NewFloatingPriceModelNewFloatingMatrixWithDisplayNamePriceModelType = "matrix_with_display_name"
)

func (r NewFloatingPriceModelNewFloatingMatrixWithDisplayNamePriceModelType) IsKnown() bool {
	switch r {
	case NewFloatingPriceModelNewFloatingMatrixWithDisplayNamePriceModelTypeMatrixWithDisplayName:
		return true
	}
	return false
}

type NewFloatingPriceModelNewFloatingBulkWithProrationPriceParam struct {
	BulkWithProrationConfig param.Field[CustomRatingFunctionConfigModelParam] `json:"bulk_with_proration_config,required"`
	// The cadence to bill for this price on.
	Cadence param.Field[NewFloatingPriceModelNewFloatingBulkWithProrationPriceCadence] `json:"cadence,required"`
	// An ISO 4217 currency string for which this price is billed in.
	Currency param.Field[string] `json:"currency,required"`
	// The id of the item the price will be associated with.
	ItemID    param.Field[string]                                                          `json:"item_id,required"`
	ModelType param.Field[NewFloatingPriceModelNewFloatingBulkWithProrationPriceModelType] `json:"model_type,required"`
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
	BillingCycleConfiguration param.Field[NewBillingCycleConfigurationModelParam] `json:"billing_cycle_configuration"`
	// The per unit conversion rate of the price currency to the invoicing currency.
	ConversionRate param.Field[float64] `json:"conversion_rate"`
	// An alias for the price.
	ExternalPriceID param.Field[string] `json:"external_price_id"`
	// If the Price represents a fixed cost, this represents the quantity of units
	// applied.
	FixedPriceQuantity param.Field[float64] `json:"fixed_price_quantity"`
	// The property used to group this price on an invoice
	InvoiceGroupingKey param.Field[string] `json:"invoice_grouping_key"`
	// Within each billing cycle, specifies the cadence at which invoices are produced.
	// If unspecified, a single invoice is produced per billing cycle.
	InvoicingCycleConfiguration param.Field[NewBillingCycleConfigurationModelParam] `json:"invoicing_cycle_configuration"`
	// User-specified key/value pairs for the resource. Individual keys can be removed
	// by setting the value to `null`, and the entire metadata mapping can be cleared
	// by setting `metadata` to `null`.
	Metadata param.Field[map[string]string] `json:"metadata"`
}

func (r NewFloatingPriceModelNewFloatingBulkWithProrationPriceParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r NewFloatingPriceModelNewFloatingBulkWithProrationPriceParam) implementsNewFloatingPriceModelUnionParam() {
}

// The cadence to bill for this price on.
type NewFloatingPriceModelNewFloatingBulkWithProrationPriceCadence string

const (
	NewFloatingPriceModelNewFloatingBulkWithProrationPriceCadenceAnnual     NewFloatingPriceModelNewFloatingBulkWithProrationPriceCadence = "annual"
	NewFloatingPriceModelNewFloatingBulkWithProrationPriceCadenceSemiAnnual NewFloatingPriceModelNewFloatingBulkWithProrationPriceCadence = "semi_annual"
	NewFloatingPriceModelNewFloatingBulkWithProrationPriceCadenceMonthly    NewFloatingPriceModelNewFloatingBulkWithProrationPriceCadence = "monthly"
	NewFloatingPriceModelNewFloatingBulkWithProrationPriceCadenceQuarterly  NewFloatingPriceModelNewFloatingBulkWithProrationPriceCadence = "quarterly"
	NewFloatingPriceModelNewFloatingBulkWithProrationPriceCadenceOneTime    NewFloatingPriceModelNewFloatingBulkWithProrationPriceCadence = "one_time"
	NewFloatingPriceModelNewFloatingBulkWithProrationPriceCadenceCustom     NewFloatingPriceModelNewFloatingBulkWithProrationPriceCadence = "custom"
)

func (r NewFloatingPriceModelNewFloatingBulkWithProrationPriceCadence) IsKnown() bool {
	switch r {
	case NewFloatingPriceModelNewFloatingBulkWithProrationPriceCadenceAnnual, NewFloatingPriceModelNewFloatingBulkWithProrationPriceCadenceSemiAnnual, NewFloatingPriceModelNewFloatingBulkWithProrationPriceCadenceMonthly, NewFloatingPriceModelNewFloatingBulkWithProrationPriceCadenceQuarterly, NewFloatingPriceModelNewFloatingBulkWithProrationPriceCadenceOneTime, NewFloatingPriceModelNewFloatingBulkWithProrationPriceCadenceCustom:
		return true
	}
	return false
}

type NewFloatingPriceModelNewFloatingBulkWithProrationPriceModelType string

const (
	NewFloatingPriceModelNewFloatingBulkWithProrationPriceModelTypeBulkWithProration NewFloatingPriceModelNewFloatingBulkWithProrationPriceModelType = "bulk_with_proration"
)

func (r NewFloatingPriceModelNewFloatingBulkWithProrationPriceModelType) IsKnown() bool {
	switch r {
	case NewFloatingPriceModelNewFloatingBulkWithProrationPriceModelTypeBulkWithProration:
		return true
	}
	return false
}

type NewFloatingPriceModelNewFloatingGroupedTieredPackagePriceParam struct {
	// The cadence to bill for this price on.
	Cadence param.Field[NewFloatingPriceModelNewFloatingGroupedTieredPackagePriceCadence] `json:"cadence,required"`
	// An ISO 4217 currency string for which this price is billed in.
	Currency                   param.Field[string]                               `json:"currency,required"`
	GroupedTieredPackageConfig param.Field[CustomRatingFunctionConfigModelParam] `json:"grouped_tiered_package_config,required"`
	// The id of the item the price will be associated with.
	ItemID    param.Field[string]                                                             `json:"item_id,required"`
	ModelType param.Field[NewFloatingPriceModelNewFloatingGroupedTieredPackagePriceModelType] `json:"model_type,required"`
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
	BillingCycleConfiguration param.Field[NewBillingCycleConfigurationModelParam] `json:"billing_cycle_configuration"`
	// The per unit conversion rate of the price currency to the invoicing currency.
	ConversionRate param.Field[float64] `json:"conversion_rate"`
	// An alias for the price.
	ExternalPriceID param.Field[string] `json:"external_price_id"`
	// If the Price represents a fixed cost, this represents the quantity of units
	// applied.
	FixedPriceQuantity param.Field[float64] `json:"fixed_price_quantity"`
	// The property used to group this price on an invoice
	InvoiceGroupingKey param.Field[string] `json:"invoice_grouping_key"`
	// Within each billing cycle, specifies the cadence at which invoices are produced.
	// If unspecified, a single invoice is produced per billing cycle.
	InvoicingCycleConfiguration param.Field[NewBillingCycleConfigurationModelParam] `json:"invoicing_cycle_configuration"`
	// User-specified key/value pairs for the resource. Individual keys can be removed
	// by setting the value to `null`, and the entire metadata mapping can be cleared
	// by setting `metadata` to `null`.
	Metadata param.Field[map[string]string] `json:"metadata"`
}

func (r NewFloatingPriceModelNewFloatingGroupedTieredPackagePriceParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r NewFloatingPriceModelNewFloatingGroupedTieredPackagePriceParam) implementsNewFloatingPriceModelUnionParam() {
}

// The cadence to bill for this price on.
type NewFloatingPriceModelNewFloatingGroupedTieredPackagePriceCadence string

const (
	NewFloatingPriceModelNewFloatingGroupedTieredPackagePriceCadenceAnnual     NewFloatingPriceModelNewFloatingGroupedTieredPackagePriceCadence = "annual"
	NewFloatingPriceModelNewFloatingGroupedTieredPackagePriceCadenceSemiAnnual NewFloatingPriceModelNewFloatingGroupedTieredPackagePriceCadence = "semi_annual"
	NewFloatingPriceModelNewFloatingGroupedTieredPackagePriceCadenceMonthly    NewFloatingPriceModelNewFloatingGroupedTieredPackagePriceCadence = "monthly"
	NewFloatingPriceModelNewFloatingGroupedTieredPackagePriceCadenceQuarterly  NewFloatingPriceModelNewFloatingGroupedTieredPackagePriceCadence = "quarterly"
	NewFloatingPriceModelNewFloatingGroupedTieredPackagePriceCadenceOneTime    NewFloatingPriceModelNewFloatingGroupedTieredPackagePriceCadence = "one_time"
	NewFloatingPriceModelNewFloatingGroupedTieredPackagePriceCadenceCustom     NewFloatingPriceModelNewFloatingGroupedTieredPackagePriceCadence = "custom"
)

func (r NewFloatingPriceModelNewFloatingGroupedTieredPackagePriceCadence) IsKnown() bool {
	switch r {
	case NewFloatingPriceModelNewFloatingGroupedTieredPackagePriceCadenceAnnual, NewFloatingPriceModelNewFloatingGroupedTieredPackagePriceCadenceSemiAnnual, NewFloatingPriceModelNewFloatingGroupedTieredPackagePriceCadenceMonthly, NewFloatingPriceModelNewFloatingGroupedTieredPackagePriceCadenceQuarterly, NewFloatingPriceModelNewFloatingGroupedTieredPackagePriceCadenceOneTime, NewFloatingPriceModelNewFloatingGroupedTieredPackagePriceCadenceCustom:
		return true
	}
	return false
}

type NewFloatingPriceModelNewFloatingGroupedTieredPackagePriceModelType string

const (
	NewFloatingPriceModelNewFloatingGroupedTieredPackagePriceModelTypeGroupedTieredPackage NewFloatingPriceModelNewFloatingGroupedTieredPackagePriceModelType = "grouped_tiered_package"
)

func (r NewFloatingPriceModelNewFloatingGroupedTieredPackagePriceModelType) IsKnown() bool {
	switch r {
	case NewFloatingPriceModelNewFloatingGroupedTieredPackagePriceModelTypeGroupedTieredPackage:
		return true
	}
	return false
}

type NewFloatingPriceModelNewFloatingScalableMatrixWithUnitPricingPriceParam struct {
	// The cadence to bill for this price on.
	Cadence param.Field[NewFloatingPriceModelNewFloatingScalableMatrixWithUnitPricingPriceCadence] `json:"cadence,required"`
	// An ISO 4217 currency string for which this price is billed in.
	Currency param.Field[string] `json:"currency,required"`
	// The id of the item the price will be associated with.
	ItemID    param.Field[string]                                                                      `json:"item_id,required"`
	ModelType param.Field[NewFloatingPriceModelNewFloatingScalableMatrixWithUnitPricingPriceModelType] `json:"model_type,required"`
	// The name of the price.
	Name                                param.Field[string]                               `json:"name,required"`
	ScalableMatrixWithUnitPricingConfig param.Field[CustomRatingFunctionConfigModelParam] `json:"scalable_matrix_with_unit_pricing_config,required"`
	// The id of the billable metric for the price. Only needed if the price is
	// usage-based.
	BillableMetricID param.Field[string] `json:"billable_metric_id"`
	// If the Price represents a fixed cost, the price will be billed in-advance if
	// this is true, and in-arrears if this is false.
	BilledInAdvance param.Field[bool] `json:"billed_in_advance"`
	// For custom cadence: specifies the duration of the billing period in days or
	// months.
	BillingCycleConfiguration param.Field[NewBillingCycleConfigurationModelParam] `json:"billing_cycle_configuration"`
	// The per unit conversion rate of the price currency to the invoicing currency.
	ConversionRate param.Field[float64] `json:"conversion_rate"`
	// An alias for the price.
	ExternalPriceID param.Field[string] `json:"external_price_id"`
	// If the Price represents a fixed cost, this represents the quantity of units
	// applied.
	FixedPriceQuantity param.Field[float64] `json:"fixed_price_quantity"`
	// The property used to group this price on an invoice
	InvoiceGroupingKey param.Field[string] `json:"invoice_grouping_key"`
	// Within each billing cycle, specifies the cadence at which invoices are produced.
	// If unspecified, a single invoice is produced per billing cycle.
	InvoicingCycleConfiguration param.Field[NewBillingCycleConfigurationModelParam] `json:"invoicing_cycle_configuration"`
	// User-specified key/value pairs for the resource. Individual keys can be removed
	// by setting the value to `null`, and the entire metadata mapping can be cleared
	// by setting `metadata` to `null`.
	Metadata param.Field[map[string]string] `json:"metadata"`
}

func (r NewFloatingPriceModelNewFloatingScalableMatrixWithUnitPricingPriceParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r NewFloatingPriceModelNewFloatingScalableMatrixWithUnitPricingPriceParam) implementsNewFloatingPriceModelUnionParam() {
}

// The cadence to bill for this price on.
type NewFloatingPriceModelNewFloatingScalableMatrixWithUnitPricingPriceCadence string

const (
	NewFloatingPriceModelNewFloatingScalableMatrixWithUnitPricingPriceCadenceAnnual     NewFloatingPriceModelNewFloatingScalableMatrixWithUnitPricingPriceCadence = "annual"
	NewFloatingPriceModelNewFloatingScalableMatrixWithUnitPricingPriceCadenceSemiAnnual NewFloatingPriceModelNewFloatingScalableMatrixWithUnitPricingPriceCadence = "semi_annual"
	NewFloatingPriceModelNewFloatingScalableMatrixWithUnitPricingPriceCadenceMonthly    NewFloatingPriceModelNewFloatingScalableMatrixWithUnitPricingPriceCadence = "monthly"
	NewFloatingPriceModelNewFloatingScalableMatrixWithUnitPricingPriceCadenceQuarterly  NewFloatingPriceModelNewFloatingScalableMatrixWithUnitPricingPriceCadence = "quarterly"
	NewFloatingPriceModelNewFloatingScalableMatrixWithUnitPricingPriceCadenceOneTime    NewFloatingPriceModelNewFloatingScalableMatrixWithUnitPricingPriceCadence = "one_time"
	NewFloatingPriceModelNewFloatingScalableMatrixWithUnitPricingPriceCadenceCustom     NewFloatingPriceModelNewFloatingScalableMatrixWithUnitPricingPriceCadence = "custom"
)

func (r NewFloatingPriceModelNewFloatingScalableMatrixWithUnitPricingPriceCadence) IsKnown() bool {
	switch r {
	case NewFloatingPriceModelNewFloatingScalableMatrixWithUnitPricingPriceCadenceAnnual, NewFloatingPriceModelNewFloatingScalableMatrixWithUnitPricingPriceCadenceSemiAnnual, NewFloatingPriceModelNewFloatingScalableMatrixWithUnitPricingPriceCadenceMonthly, NewFloatingPriceModelNewFloatingScalableMatrixWithUnitPricingPriceCadenceQuarterly, NewFloatingPriceModelNewFloatingScalableMatrixWithUnitPricingPriceCadenceOneTime, NewFloatingPriceModelNewFloatingScalableMatrixWithUnitPricingPriceCadenceCustom:
		return true
	}
	return false
}

type NewFloatingPriceModelNewFloatingScalableMatrixWithUnitPricingPriceModelType string

const (
	NewFloatingPriceModelNewFloatingScalableMatrixWithUnitPricingPriceModelTypeScalableMatrixWithUnitPricing NewFloatingPriceModelNewFloatingScalableMatrixWithUnitPricingPriceModelType = "scalable_matrix_with_unit_pricing"
)

func (r NewFloatingPriceModelNewFloatingScalableMatrixWithUnitPricingPriceModelType) IsKnown() bool {
	switch r {
	case NewFloatingPriceModelNewFloatingScalableMatrixWithUnitPricingPriceModelTypeScalableMatrixWithUnitPricing:
		return true
	}
	return false
}

type NewFloatingPriceModelNewFloatingScalableMatrixWithTieredPricingPriceParam struct {
	// The cadence to bill for this price on.
	Cadence param.Field[NewFloatingPriceModelNewFloatingScalableMatrixWithTieredPricingPriceCadence] `json:"cadence,required"`
	// An ISO 4217 currency string for which this price is billed in.
	Currency param.Field[string] `json:"currency,required"`
	// The id of the item the price will be associated with.
	ItemID    param.Field[string]                                                                        `json:"item_id,required"`
	ModelType param.Field[NewFloatingPriceModelNewFloatingScalableMatrixWithTieredPricingPriceModelType] `json:"model_type,required"`
	// The name of the price.
	Name                                  param.Field[string]                               `json:"name,required"`
	ScalableMatrixWithTieredPricingConfig param.Field[CustomRatingFunctionConfigModelParam] `json:"scalable_matrix_with_tiered_pricing_config,required"`
	// The id of the billable metric for the price. Only needed if the price is
	// usage-based.
	BillableMetricID param.Field[string] `json:"billable_metric_id"`
	// If the Price represents a fixed cost, the price will be billed in-advance if
	// this is true, and in-arrears if this is false.
	BilledInAdvance param.Field[bool] `json:"billed_in_advance"`
	// For custom cadence: specifies the duration of the billing period in days or
	// months.
	BillingCycleConfiguration param.Field[NewBillingCycleConfigurationModelParam] `json:"billing_cycle_configuration"`
	// The per unit conversion rate of the price currency to the invoicing currency.
	ConversionRate param.Field[float64] `json:"conversion_rate"`
	// An alias for the price.
	ExternalPriceID param.Field[string] `json:"external_price_id"`
	// If the Price represents a fixed cost, this represents the quantity of units
	// applied.
	FixedPriceQuantity param.Field[float64] `json:"fixed_price_quantity"`
	// The property used to group this price on an invoice
	InvoiceGroupingKey param.Field[string] `json:"invoice_grouping_key"`
	// Within each billing cycle, specifies the cadence at which invoices are produced.
	// If unspecified, a single invoice is produced per billing cycle.
	InvoicingCycleConfiguration param.Field[NewBillingCycleConfigurationModelParam] `json:"invoicing_cycle_configuration"`
	// User-specified key/value pairs for the resource. Individual keys can be removed
	// by setting the value to `null`, and the entire metadata mapping can be cleared
	// by setting `metadata` to `null`.
	Metadata param.Field[map[string]string] `json:"metadata"`
}

func (r NewFloatingPriceModelNewFloatingScalableMatrixWithTieredPricingPriceParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r NewFloatingPriceModelNewFloatingScalableMatrixWithTieredPricingPriceParam) implementsNewFloatingPriceModelUnionParam() {
}

// The cadence to bill for this price on.
type NewFloatingPriceModelNewFloatingScalableMatrixWithTieredPricingPriceCadence string

const (
	NewFloatingPriceModelNewFloatingScalableMatrixWithTieredPricingPriceCadenceAnnual     NewFloatingPriceModelNewFloatingScalableMatrixWithTieredPricingPriceCadence = "annual"
	NewFloatingPriceModelNewFloatingScalableMatrixWithTieredPricingPriceCadenceSemiAnnual NewFloatingPriceModelNewFloatingScalableMatrixWithTieredPricingPriceCadence = "semi_annual"
	NewFloatingPriceModelNewFloatingScalableMatrixWithTieredPricingPriceCadenceMonthly    NewFloatingPriceModelNewFloatingScalableMatrixWithTieredPricingPriceCadence = "monthly"
	NewFloatingPriceModelNewFloatingScalableMatrixWithTieredPricingPriceCadenceQuarterly  NewFloatingPriceModelNewFloatingScalableMatrixWithTieredPricingPriceCadence = "quarterly"
	NewFloatingPriceModelNewFloatingScalableMatrixWithTieredPricingPriceCadenceOneTime    NewFloatingPriceModelNewFloatingScalableMatrixWithTieredPricingPriceCadence = "one_time"
	NewFloatingPriceModelNewFloatingScalableMatrixWithTieredPricingPriceCadenceCustom     NewFloatingPriceModelNewFloatingScalableMatrixWithTieredPricingPriceCadence = "custom"
)

func (r NewFloatingPriceModelNewFloatingScalableMatrixWithTieredPricingPriceCadence) IsKnown() bool {
	switch r {
	case NewFloatingPriceModelNewFloatingScalableMatrixWithTieredPricingPriceCadenceAnnual, NewFloatingPriceModelNewFloatingScalableMatrixWithTieredPricingPriceCadenceSemiAnnual, NewFloatingPriceModelNewFloatingScalableMatrixWithTieredPricingPriceCadenceMonthly, NewFloatingPriceModelNewFloatingScalableMatrixWithTieredPricingPriceCadenceQuarterly, NewFloatingPriceModelNewFloatingScalableMatrixWithTieredPricingPriceCadenceOneTime, NewFloatingPriceModelNewFloatingScalableMatrixWithTieredPricingPriceCadenceCustom:
		return true
	}
	return false
}

type NewFloatingPriceModelNewFloatingScalableMatrixWithTieredPricingPriceModelType string

const (
	NewFloatingPriceModelNewFloatingScalableMatrixWithTieredPricingPriceModelTypeScalableMatrixWithTieredPricing NewFloatingPriceModelNewFloatingScalableMatrixWithTieredPricingPriceModelType = "scalable_matrix_with_tiered_pricing"
)

func (r NewFloatingPriceModelNewFloatingScalableMatrixWithTieredPricingPriceModelType) IsKnown() bool {
	switch r {
	case NewFloatingPriceModelNewFloatingScalableMatrixWithTieredPricingPriceModelTypeScalableMatrixWithTieredPricing:
		return true
	}
	return false
}

type NewFloatingPriceModelNewFloatingCumulativeGroupedBulkPriceParam struct {
	// The cadence to bill for this price on.
	Cadence                     param.Field[NewFloatingPriceModelNewFloatingCumulativeGroupedBulkPriceCadence] `json:"cadence,required"`
	CumulativeGroupedBulkConfig param.Field[CustomRatingFunctionConfigModelParam]                              `json:"cumulative_grouped_bulk_config,required"`
	// An ISO 4217 currency string for which this price is billed in.
	Currency param.Field[string] `json:"currency,required"`
	// The id of the item the price will be associated with.
	ItemID    param.Field[string]                                                              `json:"item_id,required"`
	ModelType param.Field[NewFloatingPriceModelNewFloatingCumulativeGroupedBulkPriceModelType] `json:"model_type,required"`
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
	BillingCycleConfiguration param.Field[NewBillingCycleConfigurationModelParam] `json:"billing_cycle_configuration"`
	// The per unit conversion rate of the price currency to the invoicing currency.
	ConversionRate param.Field[float64] `json:"conversion_rate"`
	// An alias for the price.
	ExternalPriceID param.Field[string] `json:"external_price_id"`
	// If the Price represents a fixed cost, this represents the quantity of units
	// applied.
	FixedPriceQuantity param.Field[float64] `json:"fixed_price_quantity"`
	// The property used to group this price on an invoice
	InvoiceGroupingKey param.Field[string] `json:"invoice_grouping_key"`
	// Within each billing cycle, specifies the cadence at which invoices are produced.
	// If unspecified, a single invoice is produced per billing cycle.
	InvoicingCycleConfiguration param.Field[NewBillingCycleConfigurationModelParam] `json:"invoicing_cycle_configuration"`
	// User-specified key/value pairs for the resource. Individual keys can be removed
	// by setting the value to `null`, and the entire metadata mapping can be cleared
	// by setting `metadata` to `null`.
	Metadata param.Field[map[string]string] `json:"metadata"`
}

func (r NewFloatingPriceModelNewFloatingCumulativeGroupedBulkPriceParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r NewFloatingPriceModelNewFloatingCumulativeGroupedBulkPriceParam) implementsNewFloatingPriceModelUnionParam() {
}

// The cadence to bill for this price on.
type NewFloatingPriceModelNewFloatingCumulativeGroupedBulkPriceCadence string

const (
	NewFloatingPriceModelNewFloatingCumulativeGroupedBulkPriceCadenceAnnual     NewFloatingPriceModelNewFloatingCumulativeGroupedBulkPriceCadence = "annual"
	NewFloatingPriceModelNewFloatingCumulativeGroupedBulkPriceCadenceSemiAnnual NewFloatingPriceModelNewFloatingCumulativeGroupedBulkPriceCadence = "semi_annual"
	NewFloatingPriceModelNewFloatingCumulativeGroupedBulkPriceCadenceMonthly    NewFloatingPriceModelNewFloatingCumulativeGroupedBulkPriceCadence = "monthly"
	NewFloatingPriceModelNewFloatingCumulativeGroupedBulkPriceCadenceQuarterly  NewFloatingPriceModelNewFloatingCumulativeGroupedBulkPriceCadence = "quarterly"
	NewFloatingPriceModelNewFloatingCumulativeGroupedBulkPriceCadenceOneTime    NewFloatingPriceModelNewFloatingCumulativeGroupedBulkPriceCadence = "one_time"
	NewFloatingPriceModelNewFloatingCumulativeGroupedBulkPriceCadenceCustom     NewFloatingPriceModelNewFloatingCumulativeGroupedBulkPriceCadence = "custom"
)

func (r NewFloatingPriceModelNewFloatingCumulativeGroupedBulkPriceCadence) IsKnown() bool {
	switch r {
	case NewFloatingPriceModelNewFloatingCumulativeGroupedBulkPriceCadenceAnnual, NewFloatingPriceModelNewFloatingCumulativeGroupedBulkPriceCadenceSemiAnnual, NewFloatingPriceModelNewFloatingCumulativeGroupedBulkPriceCadenceMonthly, NewFloatingPriceModelNewFloatingCumulativeGroupedBulkPriceCadenceQuarterly, NewFloatingPriceModelNewFloatingCumulativeGroupedBulkPriceCadenceOneTime, NewFloatingPriceModelNewFloatingCumulativeGroupedBulkPriceCadenceCustom:
		return true
	}
	return false
}

type NewFloatingPriceModelNewFloatingCumulativeGroupedBulkPriceModelType string

const (
	NewFloatingPriceModelNewFloatingCumulativeGroupedBulkPriceModelTypeCumulativeGroupedBulk NewFloatingPriceModelNewFloatingCumulativeGroupedBulkPriceModelType = "cumulative_grouped_bulk"
)

func (r NewFloatingPriceModelNewFloatingCumulativeGroupedBulkPriceModelType) IsKnown() bool {
	switch r {
	case NewFloatingPriceModelNewFloatingCumulativeGroupedBulkPriceModelTypeCumulativeGroupedBulk:
		return true
	}
	return false
}

// The cadence to bill for this price on.
type NewFloatingPriceModelCadence string

const (
	NewFloatingPriceModelCadenceAnnual     NewFloatingPriceModelCadence = "annual"
	NewFloatingPriceModelCadenceSemiAnnual NewFloatingPriceModelCadence = "semi_annual"
	NewFloatingPriceModelCadenceMonthly    NewFloatingPriceModelCadence = "monthly"
	NewFloatingPriceModelCadenceQuarterly  NewFloatingPriceModelCadence = "quarterly"
	NewFloatingPriceModelCadenceOneTime    NewFloatingPriceModelCadence = "one_time"
	NewFloatingPriceModelCadenceCustom     NewFloatingPriceModelCadence = "custom"
)

func (r NewFloatingPriceModelCadence) IsKnown() bool {
	switch r {
	case NewFloatingPriceModelCadenceAnnual, NewFloatingPriceModelCadenceSemiAnnual, NewFloatingPriceModelCadenceMonthly, NewFloatingPriceModelCadenceQuarterly, NewFloatingPriceModelCadenceOneTime, NewFloatingPriceModelCadenceCustom:
		return true
	}
	return false
}

type NewFloatingPriceModelModelType string

const (
	NewFloatingPriceModelModelTypeUnit                            NewFloatingPriceModelModelType = "unit"
	NewFloatingPriceModelModelTypePackage                         NewFloatingPriceModelModelType = "package"
	NewFloatingPriceModelModelTypeMatrix                          NewFloatingPriceModelModelType = "matrix"
	NewFloatingPriceModelModelTypeMatrixWithAllocation            NewFloatingPriceModelModelType = "matrix_with_allocation"
	NewFloatingPriceModelModelTypeTiered                          NewFloatingPriceModelModelType = "tiered"
	NewFloatingPriceModelModelTypeTieredBps                       NewFloatingPriceModelModelType = "tiered_bps"
	NewFloatingPriceModelModelTypeBps                             NewFloatingPriceModelModelType = "bps"
	NewFloatingPriceModelModelTypeBulkBps                         NewFloatingPriceModelModelType = "bulk_bps"
	NewFloatingPriceModelModelTypeBulk                            NewFloatingPriceModelModelType = "bulk"
	NewFloatingPriceModelModelTypeThresholdTotalAmount            NewFloatingPriceModelModelType = "threshold_total_amount"
	NewFloatingPriceModelModelTypeTieredPackage                   NewFloatingPriceModelModelType = "tiered_package"
	NewFloatingPriceModelModelTypeGroupedTiered                   NewFloatingPriceModelModelType = "grouped_tiered"
	NewFloatingPriceModelModelTypeMaxGroupTieredPackage           NewFloatingPriceModelModelType = "max_group_tiered_package"
	NewFloatingPriceModelModelTypeTieredWithMinimum               NewFloatingPriceModelModelType = "tiered_with_minimum"
	NewFloatingPriceModelModelTypePackageWithAllocation           NewFloatingPriceModelModelType = "package_with_allocation"
	NewFloatingPriceModelModelTypeTieredPackageWithMinimum        NewFloatingPriceModelModelType = "tiered_package_with_minimum"
	NewFloatingPriceModelModelTypeUnitWithPercent                 NewFloatingPriceModelModelType = "unit_with_percent"
	NewFloatingPriceModelModelTypeTieredWithProration             NewFloatingPriceModelModelType = "tiered_with_proration"
	NewFloatingPriceModelModelTypeUnitWithProration               NewFloatingPriceModelModelType = "unit_with_proration"
	NewFloatingPriceModelModelTypeGroupedAllocation               NewFloatingPriceModelModelType = "grouped_allocation"
	NewFloatingPriceModelModelTypeGroupedWithProratedMinimum      NewFloatingPriceModelModelType = "grouped_with_prorated_minimum"
	NewFloatingPriceModelModelTypeGroupedWithMeteredMinimum       NewFloatingPriceModelModelType = "grouped_with_metered_minimum"
	NewFloatingPriceModelModelTypeMatrixWithDisplayName           NewFloatingPriceModelModelType = "matrix_with_display_name"
	NewFloatingPriceModelModelTypeBulkWithProration               NewFloatingPriceModelModelType = "bulk_with_proration"
	NewFloatingPriceModelModelTypeGroupedTieredPackage            NewFloatingPriceModelModelType = "grouped_tiered_package"
	NewFloatingPriceModelModelTypeScalableMatrixWithUnitPricing   NewFloatingPriceModelModelType = "scalable_matrix_with_unit_pricing"
	NewFloatingPriceModelModelTypeScalableMatrixWithTieredPricing NewFloatingPriceModelModelType = "scalable_matrix_with_tiered_pricing"
	NewFloatingPriceModelModelTypeCumulativeGroupedBulk           NewFloatingPriceModelModelType = "cumulative_grouped_bulk"
)

func (r NewFloatingPriceModelModelType) IsKnown() bool {
	switch r {
	case NewFloatingPriceModelModelTypeUnit, NewFloatingPriceModelModelTypePackage, NewFloatingPriceModelModelTypeMatrix, NewFloatingPriceModelModelTypeMatrixWithAllocation, NewFloatingPriceModelModelTypeTiered, NewFloatingPriceModelModelTypeTieredBps, NewFloatingPriceModelModelTypeBps, NewFloatingPriceModelModelTypeBulkBps, NewFloatingPriceModelModelTypeBulk, NewFloatingPriceModelModelTypeThresholdTotalAmount, NewFloatingPriceModelModelTypeTieredPackage, NewFloatingPriceModelModelTypeGroupedTiered, NewFloatingPriceModelModelTypeMaxGroupTieredPackage, NewFloatingPriceModelModelTypeTieredWithMinimum, NewFloatingPriceModelModelTypePackageWithAllocation, NewFloatingPriceModelModelTypeTieredPackageWithMinimum, NewFloatingPriceModelModelTypeUnitWithPercent, NewFloatingPriceModelModelTypeTieredWithProration, NewFloatingPriceModelModelTypeUnitWithProration, NewFloatingPriceModelModelTypeGroupedAllocation, NewFloatingPriceModelModelTypeGroupedWithProratedMinimum, NewFloatingPriceModelModelTypeGroupedWithMeteredMinimum, NewFloatingPriceModelModelTypeMatrixWithDisplayName, NewFloatingPriceModelModelTypeBulkWithProration, NewFloatingPriceModelModelTypeGroupedTieredPackage, NewFloatingPriceModelModelTypeScalableMatrixWithUnitPricing, NewFloatingPriceModelModelTypeScalableMatrixWithTieredPricing, NewFloatingPriceModelModelTypeCumulativeGroupedBulk:
		return true
	}
	return false
}

type NewReportingConfigurationModelParam struct {
	Exempt param.Field[bool] `json:"exempt,required"`
}

func (r NewReportingConfigurationModelParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type NewSubscriptionPriceModelParam struct {
	// The cadence to bill for this price on.
	Cadence param.Field[NewSubscriptionPriceModelCadence] `json:"cadence,required"`
	// The id of the item the price will be associated with.
	ItemID    param.Field[string]                             `json:"item_id,required"`
	ModelType param.Field[NewSubscriptionPriceModelModelType] `json:"model_type,required"`
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
	BillingCycleConfiguration param.Field[NewBillingCycleConfigurationModelParam] `json:"billing_cycle_configuration"`
	BpsConfig                 param.Field[BpsConfigModelParam]                    `json:"bps_config"`
	BulkBpsConfig             param.Field[BulkBpsConfigModelParam]                `json:"bulk_bps_config"`
	BulkConfig                param.Field[BulkConfigModelParam]                   `json:"bulk_config"`
	BulkWithProrationConfig   param.Field[CustomRatingFunctionConfigModelParam]   `json:"bulk_with_proration_config"`
	// The per unit conversion rate of the price currency to the invoicing currency.
	ConversionRate              param.Field[float64]                              `json:"conversion_rate"`
	CumulativeGroupedBulkConfig param.Field[CustomRatingFunctionConfigModelParam] `json:"cumulative_grouped_bulk_config"`
	// An ISO 4217 currency string, or custom pricing unit identifier, in which this
	// price is billed.
	Currency param.Field[string] `json:"currency"`
	// An alias for the price.
	ExternalPriceID param.Field[string] `json:"external_price_id"`
	// If the Price represents a fixed cost, this represents the quantity of units
	// applied.
	FixedPriceQuantity               param.Field[float64]                              `json:"fixed_price_quantity"`
	GroupedAllocationConfig          param.Field[CustomRatingFunctionConfigModelParam] `json:"grouped_allocation_config"`
	GroupedTieredPackageConfig       param.Field[CustomRatingFunctionConfigModelParam] `json:"grouped_tiered_package_config"`
	GroupedWithMeteredMinimumConfig  param.Field[CustomRatingFunctionConfigModelParam] `json:"grouped_with_metered_minimum_config"`
	GroupedWithProratedMinimumConfig param.Field[CustomRatingFunctionConfigModelParam] `json:"grouped_with_prorated_minimum_config"`
	// The property used to group this price on an invoice
	InvoiceGroupingKey param.Field[string] `json:"invoice_grouping_key"`
	// Within each billing cycle, specifies the cadence at which invoices are produced.
	// If unspecified, a single invoice is produced per billing cycle.
	InvoicingCycleConfiguration param.Field[NewBillingCycleConfigurationModelParam] `json:"invoicing_cycle_configuration"`
	MatrixConfig                param.Field[MatrixConfigModelParam]                 `json:"matrix_config"`
	MatrixWithDisplayNameConfig param.Field[CustomRatingFunctionConfigModelParam]   `json:"matrix_with_display_name_config"`
	MaxGroupTieredPackageConfig param.Field[CustomRatingFunctionConfigModelParam]   `json:"max_group_tiered_package_config"`
	Metadata                    param.Field[interface{}]                            `json:"metadata"`
	PackageConfig               param.Field[PackageConfigModelParam]                `json:"package_config"`
	PackageWithAllocationConfig param.Field[CustomRatingFunctionConfigModelParam]   `json:"package_with_allocation_config"`
	// A transient ID that can be used to reference this price when adding adjustments
	// in the same API call.
	ReferenceID                           param.Field[string]                               `json:"reference_id"`
	ScalableMatrixWithTieredPricingConfig param.Field[CustomRatingFunctionConfigModelParam] `json:"scalable_matrix_with_tiered_pricing_config"`
	ScalableMatrixWithUnitPricingConfig   param.Field[CustomRatingFunctionConfigModelParam] `json:"scalable_matrix_with_unit_pricing_config"`
	ThresholdTotalAmountConfig            param.Field[CustomRatingFunctionConfigModelParam] `json:"threshold_total_amount_config"`
	TieredBpsConfig                       param.Field[TieredBpsConfigModelParam]            `json:"tiered_bps_config"`
	TieredConfig                          param.Field[TieredConfigModelParam]               `json:"tiered_config"`
	TieredPackageConfig                   param.Field[CustomRatingFunctionConfigModelParam] `json:"tiered_package_config"`
	TieredWithMinimumConfig               param.Field[CustomRatingFunctionConfigModelParam] `json:"tiered_with_minimum_config"`
	TieredWithProrationConfig             param.Field[CustomRatingFunctionConfigModelParam] `json:"tiered_with_proration_config"`
	UnitConfig                            param.Field[UnitConfigModelParam]                 `json:"unit_config"`
	UnitWithPercentConfig                 param.Field[CustomRatingFunctionConfigModelParam] `json:"unit_with_percent_config"`
	UnitWithProrationConfig               param.Field[CustomRatingFunctionConfigModelParam] `json:"unit_with_proration_config"`
}

func (r NewSubscriptionPriceModelParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r NewSubscriptionPriceModelParam) implementsNewSubscriptionPriceModelUnionParam() {}

// Satisfied by [shared.NewSubscriptionPriceModelNewSubscriptionUnitPriceParam],
// [shared.NewSubscriptionPriceModelNewSubscriptionPackagePriceParam],
// [shared.NewSubscriptionPriceModelNewSubscriptionMatrixPriceParam],
// [shared.NewSubscriptionPriceModelNewSubscriptionTieredPriceParam],
// [shared.NewSubscriptionPriceModelNewSubscriptionTieredBpsPriceParam],
// [shared.NewSubscriptionPriceModelNewSubscriptionBpsPriceParam],
// [shared.NewSubscriptionPriceModelNewSubscriptionBulkBpsPriceParam],
// [shared.NewSubscriptionPriceModelNewSubscriptionBulkPriceParam],
// [shared.NewSubscriptionPriceModelNewSubscriptionThresholdTotalAmountPriceParam],
// [shared.NewSubscriptionPriceModelNewSubscriptionTieredPackagePriceParam],
// [shared.NewSubscriptionPriceModelNewSubscriptionTieredWithMinimumPriceParam],
// [shared.NewSubscriptionPriceModelNewSubscriptionUnitWithPercentPriceParam],
// [shared.NewSubscriptionPriceModelNewSubscriptionPackageWithAllocationPriceParam],
// [shared.NewSubscriptionPriceModelNewSubscriptionTierWithProrationPriceParam],
// [shared.NewSubscriptionPriceModelNewSubscriptionUnitWithProrationPriceParam],
// [shared.NewSubscriptionPriceModelNewSubscriptionGroupedAllocationPriceParam],
// [shared.NewSubscriptionPriceModelNewSubscriptionGroupedWithProratedMinimumPriceParam],
// [shared.NewSubscriptionPriceModelNewSubscriptionBulkWithProrationPriceParam],
// [shared.NewSubscriptionPriceModelNewSubscriptionScalableMatrixWithUnitPricingPriceParam],
// [shared.NewSubscriptionPriceModelNewSubscriptionScalableMatrixWithTieredPricingPriceParam],
// [shared.NewSubscriptionPriceModelNewSubscriptionCumulativeGroupedBulkPriceParam],
// [shared.NewSubscriptionPriceModelNewSubscriptionMaxGroupTieredPackagePriceParam],
// [shared.NewSubscriptionPriceModelNewSubscriptionGroupedWithMeteredMinimumPriceParam],
// [shared.NewSubscriptionPriceModelNewSubscriptionMatrixWithDisplayNamePriceParam],
// [shared.NewSubscriptionPriceModelNewSubscriptionGroupedTieredPackagePriceParam],
// [NewSubscriptionPriceModelParam].
type NewSubscriptionPriceModelUnionParam interface {
	implementsNewSubscriptionPriceModelUnionParam()
}

type NewSubscriptionPriceModelNewSubscriptionUnitPriceParam struct {
	// The cadence to bill for this price on.
	Cadence param.Field[NewSubscriptionPriceModelNewSubscriptionUnitPriceCadence] `json:"cadence,required"`
	// The id of the item the price will be associated with.
	ItemID    param.Field[string]                                                     `json:"item_id,required"`
	ModelType param.Field[NewSubscriptionPriceModelNewSubscriptionUnitPriceModelType] `json:"model_type,required"`
	// The name of the price.
	Name       param.Field[string]               `json:"name,required"`
	UnitConfig param.Field[UnitConfigModelParam] `json:"unit_config,required"`
	// The id of the billable metric for the price. Only needed if the price is
	// usage-based.
	BillableMetricID param.Field[string] `json:"billable_metric_id"`
	// If the Price represents a fixed cost, the price will be billed in-advance if
	// this is true, and in-arrears if this is false.
	BilledInAdvance param.Field[bool] `json:"billed_in_advance"`
	// For custom cadence: specifies the duration of the billing period in days or
	// months.
	BillingCycleConfiguration param.Field[NewBillingCycleConfigurationModelParam] `json:"billing_cycle_configuration"`
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
	InvoicingCycleConfiguration param.Field[NewBillingCycleConfigurationModelParam] `json:"invoicing_cycle_configuration"`
	// User-specified key/value pairs for the resource. Individual keys can be removed
	// by setting the value to `null`, and the entire metadata mapping can be cleared
	// by setting `metadata` to `null`.
	Metadata param.Field[map[string]string] `json:"metadata"`
	// A transient ID that can be used to reference this price when adding adjustments
	// in the same API call.
	ReferenceID param.Field[string] `json:"reference_id"`
}

func (r NewSubscriptionPriceModelNewSubscriptionUnitPriceParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r NewSubscriptionPriceModelNewSubscriptionUnitPriceParam) implementsNewSubscriptionPriceModelUnionParam() {
}

// The cadence to bill for this price on.
type NewSubscriptionPriceModelNewSubscriptionUnitPriceCadence string

const (
	NewSubscriptionPriceModelNewSubscriptionUnitPriceCadenceAnnual     NewSubscriptionPriceModelNewSubscriptionUnitPriceCadence = "annual"
	NewSubscriptionPriceModelNewSubscriptionUnitPriceCadenceSemiAnnual NewSubscriptionPriceModelNewSubscriptionUnitPriceCadence = "semi_annual"
	NewSubscriptionPriceModelNewSubscriptionUnitPriceCadenceMonthly    NewSubscriptionPriceModelNewSubscriptionUnitPriceCadence = "monthly"
	NewSubscriptionPriceModelNewSubscriptionUnitPriceCadenceQuarterly  NewSubscriptionPriceModelNewSubscriptionUnitPriceCadence = "quarterly"
	NewSubscriptionPriceModelNewSubscriptionUnitPriceCadenceOneTime    NewSubscriptionPriceModelNewSubscriptionUnitPriceCadence = "one_time"
	NewSubscriptionPriceModelNewSubscriptionUnitPriceCadenceCustom     NewSubscriptionPriceModelNewSubscriptionUnitPriceCadence = "custom"
)

func (r NewSubscriptionPriceModelNewSubscriptionUnitPriceCadence) IsKnown() bool {
	switch r {
	case NewSubscriptionPriceModelNewSubscriptionUnitPriceCadenceAnnual, NewSubscriptionPriceModelNewSubscriptionUnitPriceCadenceSemiAnnual, NewSubscriptionPriceModelNewSubscriptionUnitPriceCadenceMonthly, NewSubscriptionPriceModelNewSubscriptionUnitPriceCadenceQuarterly, NewSubscriptionPriceModelNewSubscriptionUnitPriceCadenceOneTime, NewSubscriptionPriceModelNewSubscriptionUnitPriceCadenceCustom:
		return true
	}
	return false
}

type NewSubscriptionPriceModelNewSubscriptionUnitPriceModelType string

const (
	NewSubscriptionPriceModelNewSubscriptionUnitPriceModelTypeUnit NewSubscriptionPriceModelNewSubscriptionUnitPriceModelType = "unit"
)

func (r NewSubscriptionPriceModelNewSubscriptionUnitPriceModelType) IsKnown() bool {
	switch r {
	case NewSubscriptionPriceModelNewSubscriptionUnitPriceModelTypeUnit:
		return true
	}
	return false
}

type NewSubscriptionPriceModelNewSubscriptionPackagePriceParam struct {
	// The cadence to bill for this price on.
	Cadence param.Field[NewSubscriptionPriceModelNewSubscriptionPackagePriceCadence] `json:"cadence,required"`
	// The id of the item the price will be associated with.
	ItemID    param.Field[string]                                                        `json:"item_id,required"`
	ModelType param.Field[NewSubscriptionPriceModelNewSubscriptionPackagePriceModelType] `json:"model_type,required"`
	// The name of the price.
	Name          param.Field[string]                  `json:"name,required"`
	PackageConfig param.Field[PackageConfigModelParam] `json:"package_config,required"`
	// The id of the billable metric for the price. Only needed if the price is
	// usage-based.
	BillableMetricID param.Field[string] `json:"billable_metric_id"`
	// If the Price represents a fixed cost, the price will be billed in-advance if
	// this is true, and in-arrears if this is false.
	BilledInAdvance param.Field[bool] `json:"billed_in_advance"`
	// For custom cadence: specifies the duration of the billing period in days or
	// months.
	BillingCycleConfiguration param.Field[NewBillingCycleConfigurationModelParam] `json:"billing_cycle_configuration"`
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
	InvoicingCycleConfiguration param.Field[NewBillingCycleConfigurationModelParam] `json:"invoicing_cycle_configuration"`
	// User-specified key/value pairs for the resource. Individual keys can be removed
	// by setting the value to `null`, and the entire metadata mapping can be cleared
	// by setting `metadata` to `null`.
	Metadata param.Field[map[string]string] `json:"metadata"`
	// A transient ID that can be used to reference this price when adding adjustments
	// in the same API call.
	ReferenceID param.Field[string] `json:"reference_id"`
}

func (r NewSubscriptionPriceModelNewSubscriptionPackagePriceParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r NewSubscriptionPriceModelNewSubscriptionPackagePriceParam) implementsNewSubscriptionPriceModelUnionParam() {
}

// The cadence to bill for this price on.
type NewSubscriptionPriceModelNewSubscriptionPackagePriceCadence string

const (
	NewSubscriptionPriceModelNewSubscriptionPackagePriceCadenceAnnual     NewSubscriptionPriceModelNewSubscriptionPackagePriceCadence = "annual"
	NewSubscriptionPriceModelNewSubscriptionPackagePriceCadenceSemiAnnual NewSubscriptionPriceModelNewSubscriptionPackagePriceCadence = "semi_annual"
	NewSubscriptionPriceModelNewSubscriptionPackagePriceCadenceMonthly    NewSubscriptionPriceModelNewSubscriptionPackagePriceCadence = "monthly"
	NewSubscriptionPriceModelNewSubscriptionPackagePriceCadenceQuarterly  NewSubscriptionPriceModelNewSubscriptionPackagePriceCadence = "quarterly"
	NewSubscriptionPriceModelNewSubscriptionPackagePriceCadenceOneTime    NewSubscriptionPriceModelNewSubscriptionPackagePriceCadence = "one_time"
	NewSubscriptionPriceModelNewSubscriptionPackagePriceCadenceCustom     NewSubscriptionPriceModelNewSubscriptionPackagePriceCadence = "custom"
)

func (r NewSubscriptionPriceModelNewSubscriptionPackagePriceCadence) IsKnown() bool {
	switch r {
	case NewSubscriptionPriceModelNewSubscriptionPackagePriceCadenceAnnual, NewSubscriptionPriceModelNewSubscriptionPackagePriceCadenceSemiAnnual, NewSubscriptionPriceModelNewSubscriptionPackagePriceCadenceMonthly, NewSubscriptionPriceModelNewSubscriptionPackagePriceCadenceQuarterly, NewSubscriptionPriceModelNewSubscriptionPackagePriceCadenceOneTime, NewSubscriptionPriceModelNewSubscriptionPackagePriceCadenceCustom:
		return true
	}
	return false
}

type NewSubscriptionPriceModelNewSubscriptionPackagePriceModelType string

const (
	NewSubscriptionPriceModelNewSubscriptionPackagePriceModelTypePackage NewSubscriptionPriceModelNewSubscriptionPackagePriceModelType = "package"
)

func (r NewSubscriptionPriceModelNewSubscriptionPackagePriceModelType) IsKnown() bool {
	switch r {
	case NewSubscriptionPriceModelNewSubscriptionPackagePriceModelTypePackage:
		return true
	}
	return false
}

type NewSubscriptionPriceModelNewSubscriptionMatrixPriceParam struct {
	// The cadence to bill for this price on.
	Cadence param.Field[NewSubscriptionPriceModelNewSubscriptionMatrixPriceCadence] `json:"cadence,required"`
	// The id of the item the price will be associated with.
	ItemID       param.Field[string]                                                       `json:"item_id,required"`
	MatrixConfig param.Field[MatrixConfigModelParam]                                       `json:"matrix_config,required"`
	ModelType    param.Field[NewSubscriptionPriceModelNewSubscriptionMatrixPriceModelType] `json:"model_type,required"`
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
	BillingCycleConfiguration param.Field[NewBillingCycleConfigurationModelParam] `json:"billing_cycle_configuration"`
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
	InvoicingCycleConfiguration param.Field[NewBillingCycleConfigurationModelParam] `json:"invoicing_cycle_configuration"`
	// User-specified key/value pairs for the resource. Individual keys can be removed
	// by setting the value to `null`, and the entire metadata mapping can be cleared
	// by setting `metadata` to `null`.
	Metadata param.Field[map[string]string] `json:"metadata"`
	// A transient ID that can be used to reference this price when adding adjustments
	// in the same API call.
	ReferenceID param.Field[string] `json:"reference_id"`
}

func (r NewSubscriptionPriceModelNewSubscriptionMatrixPriceParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r NewSubscriptionPriceModelNewSubscriptionMatrixPriceParam) implementsNewSubscriptionPriceModelUnionParam() {
}

// The cadence to bill for this price on.
type NewSubscriptionPriceModelNewSubscriptionMatrixPriceCadence string

const (
	NewSubscriptionPriceModelNewSubscriptionMatrixPriceCadenceAnnual     NewSubscriptionPriceModelNewSubscriptionMatrixPriceCadence = "annual"
	NewSubscriptionPriceModelNewSubscriptionMatrixPriceCadenceSemiAnnual NewSubscriptionPriceModelNewSubscriptionMatrixPriceCadence = "semi_annual"
	NewSubscriptionPriceModelNewSubscriptionMatrixPriceCadenceMonthly    NewSubscriptionPriceModelNewSubscriptionMatrixPriceCadence = "monthly"
	NewSubscriptionPriceModelNewSubscriptionMatrixPriceCadenceQuarterly  NewSubscriptionPriceModelNewSubscriptionMatrixPriceCadence = "quarterly"
	NewSubscriptionPriceModelNewSubscriptionMatrixPriceCadenceOneTime    NewSubscriptionPriceModelNewSubscriptionMatrixPriceCadence = "one_time"
	NewSubscriptionPriceModelNewSubscriptionMatrixPriceCadenceCustom     NewSubscriptionPriceModelNewSubscriptionMatrixPriceCadence = "custom"
)

func (r NewSubscriptionPriceModelNewSubscriptionMatrixPriceCadence) IsKnown() bool {
	switch r {
	case NewSubscriptionPriceModelNewSubscriptionMatrixPriceCadenceAnnual, NewSubscriptionPriceModelNewSubscriptionMatrixPriceCadenceSemiAnnual, NewSubscriptionPriceModelNewSubscriptionMatrixPriceCadenceMonthly, NewSubscriptionPriceModelNewSubscriptionMatrixPriceCadenceQuarterly, NewSubscriptionPriceModelNewSubscriptionMatrixPriceCadenceOneTime, NewSubscriptionPriceModelNewSubscriptionMatrixPriceCadenceCustom:
		return true
	}
	return false
}

type NewSubscriptionPriceModelNewSubscriptionMatrixPriceModelType string

const (
	NewSubscriptionPriceModelNewSubscriptionMatrixPriceModelTypeMatrix NewSubscriptionPriceModelNewSubscriptionMatrixPriceModelType = "matrix"
)

func (r NewSubscriptionPriceModelNewSubscriptionMatrixPriceModelType) IsKnown() bool {
	switch r {
	case NewSubscriptionPriceModelNewSubscriptionMatrixPriceModelTypeMatrix:
		return true
	}
	return false
}

type NewSubscriptionPriceModelNewSubscriptionTieredPriceParam struct {
	// The cadence to bill for this price on.
	Cadence param.Field[NewSubscriptionPriceModelNewSubscriptionTieredPriceCadence] `json:"cadence,required"`
	// The id of the item the price will be associated with.
	ItemID    param.Field[string]                                                       `json:"item_id,required"`
	ModelType param.Field[NewSubscriptionPriceModelNewSubscriptionTieredPriceModelType] `json:"model_type,required"`
	// The name of the price.
	Name         param.Field[string]                 `json:"name,required"`
	TieredConfig param.Field[TieredConfigModelParam] `json:"tiered_config,required"`
	// The id of the billable metric for the price. Only needed if the price is
	// usage-based.
	BillableMetricID param.Field[string] `json:"billable_metric_id"`
	// If the Price represents a fixed cost, the price will be billed in-advance if
	// this is true, and in-arrears if this is false.
	BilledInAdvance param.Field[bool] `json:"billed_in_advance"`
	// For custom cadence: specifies the duration of the billing period in days or
	// months.
	BillingCycleConfiguration param.Field[NewBillingCycleConfigurationModelParam] `json:"billing_cycle_configuration"`
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
	InvoicingCycleConfiguration param.Field[NewBillingCycleConfigurationModelParam] `json:"invoicing_cycle_configuration"`
	// User-specified key/value pairs for the resource. Individual keys can be removed
	// by setting the value to `null`, and the entire metadata mapping can be cleared
	// by setting `metadata` to `null`.
	Metadata param.Field[map[string]string] `json:"metadata"`
	// A transient ID that can be used to reference this price when adding adjustments
	// in the same API call.
	ReferenceID param.Field[string] `json:"reference_id"`
}

func (r NewSubscriptionPriceModelNewSubscriptionTieredPriceParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r NewSubscriptionPriceModelNewSubscriptionTieredPriceParam) implementsNewSubscriptionPriceModelUnionParam() {
}

// The cadence to bill for this price on.
type NewSubscriptionPriceModelNewSubscriptionTieredPriceCadence string

const (
	NewSubscriptionPriceModelNewSubscriptionTieredPriceCadenceAnnual     NewSubscriptionPriceModelNewSubscriptionTieredPriceCadence = "annual"
	NewSubscriptionPriceModelNewSubscriptionTieredPriceCadenceSemiAnnual NewSubscriptionPriceModelNewSubscriptionTieredPriceCadence = "semi_annual"
	NewSubscriptionPriceModelNewSubscriptionTieredPriceCadenceMonthly    NewSubscriptionPriceModelNewSubscriptionTieredPriceCadence = "monthly"
	NewSubscriptionPriceModelNewSubscriptionTieredPriceCadenceQuarterly  NewSubscriptionPriceModelNewSubscriptionTieredPriceCadence = "quarterly"
	NewSubscriptionPriceModelNewSubscriptionTieredPriceCadenceOneTime    NewSubscriptionPriceModelNewSubscriptionTieredPriceCadence = "one_time"
	NewSubscriptionPriceModelNewSubscriptionTieredPriceCadenceCustom     NewSubscriptionPriceModelNewSubscriptionTieredPriceCadence = "custom"
)

func (r NewSubscriptionPriceModelNewSubscriptionTieredPriceCadence) IsKnown() bool {
	switch r {
	case NewSubscriptionPriceModelNewSubscriptionTieredPriceCadenceAnnual, NewSubscriptionPriceModelNewSubscriptionTieredPriceCadenceSemiAnnual, NewSubscriptionPriceModelNewSubscriptionTieredPriceCadenceMonthly, NewSubscriptionPriceModelNewSubscriptionTieredPriceCadenceQuarterly, NewSubscriptionPriceModelNewSubscriptionTieredPriceCadenceOneTime, NewSubscriptionPriceModelNewSubscriptionTieredPriceCadenceCustom:
		return true
	}
	return false
}

type NewSubscriptionPriceModelNewSubscriptionTieredPriceModelType string

const (
	NewSubscriptionPriceModelNewSubscriptionTieredPriceModelTypeTiered NewSubscriptionPriceModelNewSubscriptionTieredPriceModelType = "tiered"
)

func (r NewSubscriptionPriceModelNewSubscriptionTieredPriceModelType) IsKnown() bool {
	switch r {
	case NewSubscriptionPriceModelNewSubscriptionTieredPriceModelTypeTiered:
		return true
	}
	return false
}

type NewSubscriptionPriceModelNewSubscriptionTieredBpsPriceParam struct {
	// The cadence to bill for this price on.
	Cadence param.Field[NewSubscriptionPriceModelNewSubscriptionTieredBpsPriceCadence] `json:"cadence,required"`
	// The id of the item the price will be associated with.
	ItemID    param.Field[string]                                                          `json:"item_id,required"`
	ModelType param.Field[NewSubscriptionPriceModelNewSubscriptionTieredBpsPriceModelType] `json:"model_type,required"`
	// The name of the price.
	Name            param.Field[string]                    `json:"name,required"`
	TieredBpsConfig param.Field[TieredBpsConfigModelParam] `json:"tiered_bps_config,required"`
	// The id of the billable metric for the price. Only needed if the price is
	// usage-based.
	BillableMetricID param.Field[string] `json:"billable_metric_id"`
	// If the Price represents a fixed cost, the price will be billed in-advance if
	// this is true, and in-arrears if this is false.
	BilledInAdvance param.Field[bool] `json:"billed_in_advance"`
	// For custom cadence: specifies the duration of the billing period in days or
	// months.
	BillingCycleConfiguration param.Field[NewBillingCycleConfigurationModelParam] `json:"billing_cycle_configuration"`
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
	InvoicingCycleConfiguration param.Field[NewBillingCycleConfigurationModelParam] `json:"invoicing_cycle_configuration"`
	// User-specified key/value pairs for the resource. Individual keys can be removed
	// by setting the value to `null`, and the entire metadata mapping can be cleared
	// by setting `metadata` to `null`.
	Metadata param.Field[map[string]string] `json:"metadata"`
	// A transient ID that can be used to reference this price when adding adjustments
	// in the same API call.
	ReferenceID param.Field[string] `json:"reference_id"`
}

func (r NewSubscriptionPriceModelNewSubscriptionTieredBpsPriceParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r NewSubscriptionPriceModelNewSubscriptionTieredBpsPriceParam) implementsNewSubscriptionPriceModelUnionParam() {
}

// The cadence to bill for this price on.
type NewSubscriptionPriceModelNewSubscriptionTieredBpsPriceCadence string

const (
	NewSubscriptionPriceModelNewSubscriptionTieredBpsPriceCadenceAnnual     NewSubscriptionPriceModelNewSubscriptionTieredBpsPriceCadence = "annual"
	NewSubscriptionPriceModelNewSubscriptionTieredBpsPriceCadenceSemiAnnual NewSubscriptionPriceModelNewSubscriptionTieredBpsPriceCadence = "semi_annual"
	NewSubscriptionPriceModelNewSubscriptionTieredBpsPriceCadenceMonthly    NewSubscriptionPriceModelNewSubscriptionTieredBpsPriceCadence = "monthly"
	NewSubscriptionPriceModelNewSubscriptionTieredBpsPriceCadenceQuarterly  NewSubscriptionPriceModelNewSubscriptionTieredBpsPriceCadence = "quarterly"
	NewSubscriptionPriceModelNewSubscriptionTieredBpsPriceCadenceOneTime    NewSubscriptionPriceModelNewSubscriptionTieredBpsPriceCadence = "one_time"
	NewSubscriptionPriceModelNewSubscriptionTieredBpsPriceCadenceCustom     NewSubscriptionPriceModelNewSubscriptionTieredBpsPriceCadence = "custom"
)

func (r NewSubscriptionPriceModelNewSubscriptionTieredBpsPriceCadence) IsKnown() bool {
	switch r {
	case NewSubscriptionPriceModelNewSubscriptionTieredBpsPriceCadenceAnnual, NewSubscriptionPriceModelNewSubscriptionTieredBpsPriceCadenceSemiAnnual, NewSubscriptionPriceModelNewSubscriptionTieredBpsPriceCadenceMonthly, NewSubscriptionPriceModelNewSubscriptionTieredBpsPriceCadenceQuarterly, NewSubscriptionPriceModelNewSubscriptionTieredBpsPriceCadenceOneTime, NewSubscriptionPriceModelNewSubscriptionTieredBpsPriceCadenceCustom:
		return true
	}
	return false
}

type NewSubscriptionPriceModelNewSubscriptionTieredBpsPriceModelType string

const (
	NewSubscriptionPriceModelNewSubscriptionTieredBpsPriceModelTypeTieredBps NewSubscriptionPriceModelNewSubscriptionTieredBpsPriceModelType = "tiered_bps"
)

func (r NewSubscriptionPriceModelNewSubscriptionTieredBpsPriceModelType) IsKnown() bool {
	switch r {
	case NewSubscriptionPriceModelNewSubscriptionTieredBpsPriceModelTypeTieredBps:
		return true
	}
	return false
}

type NewSubscriptionPriceModelNewSubscriptionBpsPriceParam struct {
	BpsConfig param.Field[BpsConfigModelParam] `json:"bps_config,required"`
	// The cadence to bill for this price on.
	Cadence param.Field[NewSubscriptionPriceModelNewSubscriptionBpsPriceCadence] `json:"cadence,required"`
	// The id of the item the price will be associated with.
	ItemID    param.Field[string]                                                    `json:"item_id,required"`
	ModelType param.Field[NewSubscriptionPriceModelNewSubscriptionBpsPriceModelType] `json:"model_type,required"`
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
	BillingCycleConfiguration param.Field[NewBillingCycleConfigurationModelParam] `json:"billing_cycle_configuration"`
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
	InvoicingCycleConfiguration param.Field[NewBillingCycleConfigurationModelParam] `json:"invoicing_cycle_configuration"`
	// User-specified key/value pairs for the resource. Individual keys can be removed
	// by setting the value to `null`, and the entire metadata mapping can be cleared
	// by setting `metadata` to `null`.
	Metadata param.Field[map[string]string] `json:"metadata"`
	// A transient ID that can be used to reference this price when adding adjustments
	// in the same API call.
	ReferenceID param.Field[string] `json:"reference_id"`
}

func (r NewSubscriptionPriceModelNewSubscriptionBpsPriceParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r NewSubscriptionPriceModelNewSubscriptionBpsPriceParam) implementsNewSubscriptionPriceModelUnionParam() {
}

// The cadence to bill for this price on.
type NewSubscriptionPriceModelNewSubscriptionBpsPriceCadence string

const (
	NewSubscriptionPriceModelNewSubscriptionBpsPriceCadenceAnnual     NewSubscriptionPriceModelNewSubscriptionBpsPriceCadence = "annual"
	NewSubscriptionPriceModelNewSubscriptionBpsPriceCadenceSemiAnnual NewSubscriptionPriceModelNewSubscriptionBpsPriceCadence = "semi_annual"
	NewSubscriptionPriceModelNewSubscriptionBpsPriceCadenceMonthly    NewSubscriptionPriceModelNewSubscriptionBpsPriceCadence = "monthly"
	NewSubscriptionPriceModelNewSubscriptionBpsPriceCadenceQuarterly  NewSubscriptionPriceModelNewSubscriptionBpsPriceCadence = "quarterly"
	NewSubscriptionPriceModelNewSubscriptionBpsPriceCadenceOneTime    NewSubscriptionPriceModelNewSubscriptionBpsPriceCadence = "one_time"
	NewSubscriptionPriceModelNewSubscriptionBpsPriceCadenceCustom     NewSubscriptionPriceModelNewSubscriptionBpsPriceCadence = "custom"
)

func (r NewSubscriptionPriceModelNewSubscriptionBpsPriceCadence) IsKnown() bool {
	switch r {
	case NewSubscriptionPriceModelNewSubscriptionBpsPriceCadenceAnnual, NewSubscriptionPriceModelNewSubscriptionBpsPriceCadenceSemiAnnual, NewSubscriptionPriceModelNewSubscriptionBpsPriceCadenceMonthly, NewSubscriptionPriceModelNewSubscriptionBpsPriceCadenceQuarterly, NewSubscriptionPriceModelNewSubscriptionBpsPriceCadenceOneTime, NewSubscriptionPriceModelNewSubscriptionBpsPriceCadenceCustom:
		return true
	}
	return false
}

type NewSubscriptionPriceModelNewSubscriptionBpsPriceModelType string

const (
	NewSubscriptionPriceModelNewSubscriptionBpsPriceModelTypeBps NewSubscriptionPriceModelNewSubscriptionBpsPriceModelType = "bps"
)

func (r NewSubscriptionPriceModelNewSubscriptionBpsPriceModelType) IsKnown() bool {
	switch r {
	case NewSubscriptionPriceModelNewSubscriptionBpsPriceModelTypeBps:
		return true
	}
	return false
}

type NewSubscriptionPriceModelNewSubscriptionBulkBpsPriceParam struct {
	BulkBpsConfig param.Field[BulkBpsConfigModelParam] `json:"bulk_bps_config,required"`
	// The cadence to bill for this price on.
	Cadence param.Field[NewSubscriptionPriceModelNewSubscriptionBulkBpsPriceCadence] `json:"cadence,required"`
	// The id of the item the price will be associated with.
	ItemID    param.Field[string]                                                        `json:"item_id,required"`
	ModelType param.Field[NewSubscriptionPriceModelNewSubscriptionBulkBpsPriceModelType] `json:"model_type,required"`
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
	BillingCycleConfiguration param.Field[NewBillingCycleConfigurationModelParam] `json:"billing_cycle_configuration"`
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
	InvoicingCycleConfiguration param.Field[NewBillingCycleConfigurationModelParam] `json:"invoicing_cycle_configuration"`
	// User-specified key/value pairs for the resource. Individual keys can be removed
	// by setting the value to `null`, and the entire metadata mapping can be cleared
	// by setting `metadata` to `null`.
	Metadata param.Field[map[string]string] `json:"metadata"`
	// A transient ID that can be used to reference this price when adding adjustments
	// in the same API call.
	ReferenceID param.Field[string] `json:"reference_id"`
}

func (r NewSubscriptionPriceModelNewSubscriptionBulkBpsPriceParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r NewSubscriptionPriceModelNewSubscriptionBulkBpsPriceParam) implementsNewSubscriptionPriceModelUnionParam() {
}

// The cadence to bill for this price on.
type NewSubscriptionPriceModelNewSubscriptionBulkBpsPriceCadence string

const (
	NewSubscriptionPriceModelNewSubscriptionBulkBpsPriceCadenceAnnual     NewSubscriptionPriceModelNewSubscriptionBulkBpsPriceCadence = "annual"
	NewSubscriptionPriceModelNewSubscriptionBulkBpsPriceCadenceSemiAnnual NewSubscriptionPriceModelNewSubscriptionBulkBpsPriceCadence = "semi_annual"
	NewSubscriptionPriceModelNewSubscriptionBulkBpsPriceCadenceMonthly    NewSubscriptionPriceModelNewSubscriptionBulkBpsPriceCadence = "monthly"
	NewSubscriptionPriceModelNewSubscriptionBulkBpsPriceCadenceQuarterly  NewSubscriptionPriceModelNewSubscriptionBulkBpsPriceCadence = "quarterly"
	NewSubscriptionPriceModelNewSubscriptionBulkBpsPriceCadenceOneTime    NewSubscriptionPriceModelNewSubscriptionBulkBpsPriceCadence = "one_time"
	NewSubscriptionPriceModelNewSubscriptionBulkBpsPriceCadenceCustom     NewSubscriptionPriceModelNewSubscriptionBulkBpsPriceCadence = "custom"
)

func (r NewSubscriptionPriceModelNewSubscriptionBulkBpsPriceCadence) IsKnown() bool {
	switch r {
	case NewSubscriptionPriceModelNewSubscriptionBulkBpsPriceCadenceAnnual, NewSubscriptionPriceModelNewSubscriptionBulkBpsPriceCadenceSemiAnnual, NewSubscriptionPriceModelNewSubscriptionBulkBpsPriceCadenceMonthly, NewSubscriptionPriceModelNewSubscriptionBulkBpsPriceCadenceQuarterly, NewSubscriptionPriceModelNewSubscriptionBulkBpsPriceCadenceOneTime, NewSubscriptionPriceModelNewSubscriptionBulkBpsPriceCadenceCustom:
		return true
	}
	return false
}

type NewSubscriptionPriceModelNewSubscriptionBulkBpsPriceModelType string

const (
	NewSubscriptionPriceModelNewSubscriptionBulkBpsPriceModelTypeBulkBps NewSubscriptionPriceModelNewSubscriptionBulkBpsPriceModelType = "bulk_bps"
)

func (r NewSubscriptionPriceModelNewSubscriptionBulkBpsPriceModelType) IsKnown() bool {
	switch r {
	case NewSubscriptionPriceModelNewSubscriptionBulkBpsPriceModelTypeBulkBps:
		return true
	}
	return false
}

type NewSubscriptionPriceModelNewSubscriptionBulkPriceParam struct {
	BulkConfig param.Field[BulkConfigModelParam] `json:"bulk_config,required"`
	// The cadence to bill for this price on.
	Cadence param.Field[NewSubscriptionPriceModelNewSubscriptionBulkPriceCadence] `json:"cadence,required"`
	// The id of the item the price will be associated with.
	ItemID    param.Field[string]                                                     `json:"item_id,required"`
	ModelType param.Field[NewSubscriptionPriceModelNewSubscriptionBulkPriceModelType] `json:"model_type,required"`
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
	BillingCycleConfiguration param.Field[NewBillingCycleConfigurationModelParam] `json:"billing_cycle_configuration"`
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
	InvoicingCycleConfiguration param.Field[NewBillingCycleConfigurationModelParam] `json:"invoicing_cycle_configuration"`
	// User-specified key/value pairs for the resource. Individual keys can be removed
	// by setting the value to `null`, and the entire metadata mapping can be cleared
	// by setting `metadata` to `null`.
	Metadata param.Field[map[string]string] `json:"metadata"`
	// A transient ID that can be used to reference this price when adding adjustments
	// in the same API call.
	ReferenceID param.Field[string] `json:"reference_id"`
}

func (r NewSubscriptionPriceModelNewSubscriptionBulkPriceParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r NewSubscriptionPriceModelNewSubscriptionBulkPriceParam) implementsNewSubscriptionPriceModelUnionParam() {
}

// The cadence to bill for this price on.
type NewSubscriptionPriceModelNewSubscriptionBulkPriceCadence string

const (
	NewSubscriptionPriceModelNewSubscriptionBulkPriceCadenceAnnual     NewSubscriptionPriceModelNewSubscriptionBulkPriceCadence = "annual"
	NewSubscriptionPriceModelNewSubscriptionBulkPriceCadenceSemiAnnual NewSubscriptionPriceModelNewSubscriptionBulkPriceCadence = "semi_annual"
	NewSubscriptionPriceModelNewSubscriptionBulkPriceCadenceMonthly    NewSubscriptionPriceModelNewSubscriptionBulkPriceCadence = "monthly"
	NewSubscriptionPriceModelNewSubscriptionBulkPriceCadenceQuarterly  NewSubscriptionPriceModelNewSubscriptionBulkPriceCadence = "quarterly"
	NewSubscriptionPriceModelNewSubscriptionBulkPriceCadenceOneTime    NewSubscriptionPriceModelNewSubscriptionBulkPriceCadence = "one_time"
	NewSubscriptionPriceModelNewSubscriptionBulkPriceCadenceCustom     NewSubscriptionPriceModelNewSubscriptionBulkPriceCadence = "custom"
)

func (r NewSubscriptionPriceModelNewSubscriptionBulkPriceCadence) IsKnown() bool {
	switch r {
	case NewSubscriptionPriceModelNewSubscriptionBulkPriceCadenceAnnual, NewSubscriptionPriceModelNewSubscriptionBulkPriceCadenceSemiAnnual, NewSubscriptionPriceModelNewSubscriptionBulkPriceCadenceMonthly, NewSubscriptionPriceModelNewSubscriptionBulkPriceCadenceQuarterly, NewSubscriptionPriceModelNewSubscriptionBulkPriceCadenceOneTime, NewSubscriptionPriceModelNewSubscriptionBulkPriceCadenceCustom:
		return true
	}
	return false
}

type NewSubscriptionPriceModelNewSubscriptionBulkPriceModelType string

const (
	NewSubscriptionPriceModelNewSubscriptionBulkPriceModelTypeBulk NewSubscriptionPriceModelNewSubscriptionBulkPriceModelType = "bulk"
)

func (r NewSubscriptionPriceModelNewSubscriptionBulkPriceModelType) IsKnown() bool {
	switch r {
	case NewSubscriptionPriceModelNewSubscriptionBulkPriceModelTypeBulk:
		return true
	}
	return false
}

type NewSubscriptionPriceModelNewSubscriptionThresholdTotalAmountPriceParam struct {
	// The cadence to bill for this price on.
	Cadence param.Field[NewSubscriptionPriceModelNewSubscriptionThresholdTotalAmountPriceCadence] `json:"cadence,required"`
	// The id of the item the price will be associated with.
	ItemID    param.Field[string]                                                                     `json:"item_id,required"`
	ModelType param.Field[NewSubscriptionPriceModelNewSubscriptionThresholdTotalAmountPriceModelType] `json:"model_type,required"`
	// The name of the price.
	Name                       param.Field[string]                               `json:"name,required"`
	ThresholdTotalAmountConfig param.Field[CustomRatingFunctionConfigModelParam] `json:"threshold_total_amount_config,required"`
	// The id of the billable metric for the price. Only needed if the price is
	// usage-based.
	BillableMetricID param.Field[string] `json:"billable_metric_id"`
	// If the Price represents a fixed cost, the price will be billed in-advance if
	// this is true, and in-arrears if this is false.
	BilledInAdvance param.Field[bool] `json:"billed_in_advance"`
	// For custom cadence: specifies the duration of the billing period in days or
	// months.
	BillingCycleConfiguration param.Field[NewBillingCycleConfigurationModelParam] `json:"billing_cycle_configuration"`
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
	InvoicingCycleConfiguration param.Field[NewBillingCycleConfigurationModelParam] `json:"invoicing_cycle_configuration"`
	// User-specified key/value pairs for the resource. Individual keys can be removed
	// by setting the value to `null`, and the entire metadata mapping can be cleared
	// by setting `metadata` to `null`.
	Metadata param.Field[map[string]string] `json:"metadata"`
	// A transient ID that can be used to reference this price when adding adjustments
	// in the same API call.
	ReferenceID param.Field[string] `json:"reference_id"`
}

func (r NewSubscriptionPriceModelNewSubscriptionThresholdTotalAmountPriceParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r NewSubscriptionPriceModelNewSubscriptionThresholdTotalAmountPriceParam) implementsNewSubscriptionPriceModelUnionParam() {
}

// The cadence to bill for this price on.
type NewSubscriptionPriceModelNewSubscriptionThresholdTotalAmountPriceCadence string

const (
	NewSubscriptionPriceModelNewSubscriptionThresholdTotalAmountPriceCadenceAnnual     NewSubscriptionPriceModelNewSubscriptionThresholdTotalAmountPriceCadence = "annual"
	NewSubscriptionPriceModelNewSubscriptionThresholdTotalAmountPriceCadenceSemiAnnual NewSubscriptionPriceModelNewSubscriptionThresholdTotalAmountPriceCadence = "semi_annual"
	NewSubscriptionPriceModelNewSubscriptionThresholdTotalAmountPriceCadenceMonthly    NewSubscriptionPriceModelNewSubscriptionThresholdTotalAmountPriceCadence = "monthly"
	NewSubscriptionPriceModelNewSubscriptionThresholdTotalAmountPriceCadenceQuarterly  NewSubscriptionPriceModelNewSubscriptionThresholdTotalAmountPriceCadence = "quarterly"
	NewSubscriptionPriceModelNewSubscriptionThresholdTotalAmountPriceCadenceOneTime    NewSubscriptionPriceModelNewSubscriptionThresholdTotalAmountPriceCadence = "one_time"
	NewSubscriptionPriceModelNewSubscriptionThresholdTotalAmountPriceCadenceCustom     NewSubscriptionPriceModelNewSubscriptionThresholdTotalAmountPriceCadence = "custom"
)

func (r NewSubscriptionPriceModelNewSubscriptionThresholdTotalAmountPriceCadence) IsKnown() bool {
	switch r {
	case NewSubscriptionPriceModelNewSubscriptionThresholdTotalAmountPriceCadenceAnnual, NewSubscriptionPriceModelNewSubscriptionThresholdTotalAmountPriceCadenceSemiAnnual, NewSubscriptionPriceModelNewSubscriptionThresholdTotalAmountPriceCadenceMonthly, NewSubscriptionPriceModelNewSubscriptionThresholdTotalAmountPriceCadenceQuarterly, NewSubscriptionPriceModelNewSubscriptionThresholdTotalAmountPriceCadenceOneTime, NewSubscriptionPriceModelNewSubscriptionThresholdTotalAmountPriceCadenceCustom:
		return true
	}
	return false
}

type NewSubscriptionPriceModelNewSubscriptionThresholdTotalAmountPriceModelType string

const (
	NewSubscriptionPriceModelNewSubscriptionThresholdTotalAmountPriceModelTypeThresholdTotalAmount NewSubscriptionPriceModelNewSubscriptionThresholdTotalAmountPriceModelType = "threshold_total_amount"
)

func (r NewSubscriptionPriceModelNewSubscriptionThresholdTotalAmountPriceModelType) IsKnown() bool {
	switch r {
	case NewSubscriptionPriceModelNewSubscriptionThresholdTotalAmountPriceModelTypeThresholdTotalAmount:
		return true
	}
	return false
}

type NewSubscriptionPriceModelNewSubscriptionTieredPackagePriceParam struct {
	// The cadence to bill for this price on.
	Cadence param.Field[NewSubscriptionPriceModelNewSubscriptionTieredPackagePriceCadence] `json:"cadence,required"`
	// The id of the item the price will be associated with.
	ItemID    param.Field[string]                                                              `json:"item_id,required"`
	ModelType param.Field[NewSubscriptionPriceModelNewSubscriptionTieredPackagePriceModelType] `json:"model_type,required"`
	// The name of the price.
	Name                param.Field[string]                               `json:"name,required"`
	TieredPackageConfig param.Field[CustomRatingFunctionConfigModelParam] `json:"tiered_package_config,required"`
	// The id of the billable metric for the price. Only needed if the price is
	// usage-based.
	BillableMetricID param.Field[string] `json:"billable_metric_id"`
	// If the Price represents a fixed cost, the price will be billed in-advance if
	// this is true, and in-arrears if this is false.
	BilledInAdvance param.Field[bool] `json:"billed_in_advance"`
	// For custom cadence: specifies the duration of the billing period in days or
	// months.
	BillingCycleConfiguration param.Field[NewBillingCycleConfigurationModelParam] `json:"billing_cycle_configuration"`
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
	InvoicingCycleConfiguration param.Field[NewBillingCycleConfigurationModelParam] `json:"invoicing_cycle_configuration"`
	// User-specified key/value pairs for the resource. Individual keys can be removed
	// by setting the value to `null`, and the entire metadata mapping can be cleared
	// by setting `metadata` to `null`.
	Metadata param.Field[map[string]string] `json:"metadata"`
	// A transient ID that can be used to reference this price when adding adjustments
	// in the same API call.
	ReferenceID param.Field[string] `json:"reference_id"`
}

func (r NewSubscriptionPriceModelNewSubscriptionTieredPackagePriceParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r NewSubscriptionPriceModelNewSubscriptionTieredPackagePriceParam) implementsNewSubscriptionPriceModelUnionParam() {
}

// The cadence to bill for this price on.
type NewSubscriptionPriceModelNewSubscriptionTieredPackagePriceCadence string

const (
	NewSubscriptionPriceModelNewSubscriptionTieredPackagePriceCadenceAnnual     NewSubscriptionPriceModelNewSubscriptionTieredPackagePriceCadence = "annual"
	NewSubscriptionPriceModelNewSubscriptionTieredPackagePriceCadenceSemiAnnual NewSubscriptionPriceModelNewSubscriptionTieredPackagePriceCadence = "semi_annual"
	NewSubscriptionPriceModelNewSubscriptionTieredPackagePriceCadenceMonthly    NewSubscriptionPriceModelNewSubscriptionTieredPackagePriceCadence = "monthly"
	NewSubscriptionPriceModelNewSubscriptionTieredPackagePriceCadenceQuarterly  NewSubscriptionPriceModelNewSubscriptionTieredPackagePriceCadence = "quarterly"
	NewSubscriptionPriceModelNewSubscriptionTieredPackagePriceCadenceOneTime    NewSubscriptionPriceModelNewSubscriptionTieredPackagePriceCadence = "one_time"
	NewSubscriptionPriceModelNewSubscriptionTieredPackagePriceCadenceCustom     NewSubscriptionPriceModelNewSubscriptionTieredPackagePriceCadence = "custom"
)

func (r NewSubscriptionPriceModelNewSubscriptionTieredPackagePriceCadence) IsKnown() bool {
	switch r {
	case NewSubscriptionPriceModelNewSubscriptionTieredPackagePriceCadenceAnnual, NewSubscriptionPriceModelNewSubscriptionTieredPackagePriceCadenceSemiAnnual, NewSubscriptionPriceModelNewSubscriptionTieredPackagePriceCadenceMonthly, NewSubscriptionPriceModelNewSubscriptionTieredPackagePriceCadenceQuarterly, NewSubscriptionPriceModelNewSubscriptionTieredPackagePriceCadenceOneTime, NewSubscriptionPriceModelNewSubscriptionTieredPackagePriceCadenceCustom:
		return true
	}
	return false
}

type NewSubscriptionPriceModelNewSubscriptionTieredPackagePriceModelType string

const (
	NewSubscriptionPriceModelNewSubscriptionTieredPackagePriceModelTypeTieredPackage NewSubscriptionPriceModelNewSubscriptionTieredPackagePriceModelType = "tiered_package"
)

func (r NewSubscriptionPriceModelNewSubscriptionTieredPackagePriceModelType) IsKnown() bool {
	switch r {
	case NewSubscriptionPriceModelNewSubscriptionTieredPackagePriceModelTypeTieredPackage:
		return true
	}
	return false
}

type NewSubscriptionPriceModelNewSubscriptionTieredWithMinimumPriceParam struct {
	// The cadence to bill for this price on.
	Cadence param.Field[NewSubscriptionPriceModelNewSubscriptionTieredWithMinimumPriceCadence] `json:"cadence,required"`
	// The id of the item the price will be associated with.
	ItemID    param.Field[string]                                                                  `json:"item_id,required"`
	ModelType param.Field[NewSubscriptionPriceModelNewSubscriptionTieredWithMinimumPriceModelType] `json:"model_type,required"`
	// The name of the price.
	Name                    param.Field[string]                               `json:"name,required"`
	TieredWithMinimumConfig param.Field[CustomRatingFunctionConfigModelParam] `json:"tiered_with_minimum_config,required"`
	// The id of the billable metric for the price. Only needed if the price is
	// usage-based.
	BillableMetricID param.Field[string] `json:"billable_metric_id"`
	// If the Price represents a fixed cost, the price will be billed in-advance if
	// this is true, and in-arrears if this is false.
	BilledInAdvance param.Field[bool] `json:"billed_in_advance"`
	// For custom cadence: specifies the duration of the billing period in days or
	// months.
	BillingCycleConfiguration param.Field[NewBillingCycleConfigurationModelParam] `json:"billing_cycle_configuration"`
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
	InvoicingCycleConfiguration param.Field[NewBillingCycleConfigurationModelParam] `json:"invoicing_cycle_configuration"`
	// User-specified key/value pairs for the resource. Individual keys can be removed
	// by setting the value to `null`, and the entire metadata mapping can be cleared
	// by setting `metadata` to `null`.
	Metadata param.Field[map[string]string] `json:"metadata"`
	// A transient ID that can be used to reference this price when adding adjustments
	// in the same API call.
	ReferenceID param.Field[string] `json:"reference_id"`
}

func (r NewSubscriptionPriceModelNewSubscriptionTieredWithMinimumPriceParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r NewSubscriptionPriceModelNewSubscriptionTieredWithMinimumPriceParam) implementsNewSubscriptionPriceModelUnionParam() {
}

// The cadence to bill for this price on.
type NewSubscriptionPriceModelNewSubscriptionTieredWithMinimumPriceCadence string

const (
	NewSubscriptionPriceModelNewSubscriptionTieredWithMinimumPriceCadenceAnnual     NewSubscriptionPriceModelNewSubscriptionTieredWithMinimumPriceCadence = "annual"
	NewSubscriptionPriceModelNewSubscriptionTieredWithMinimumPriceCadenceSemiAnnual NewSubscriptionPriceModelNewSubscriptionTieredWithMinimumPriceCadence = "semi_annual"
	NewSubscriptionPriceModelNewSubscriptionTieredWithMinimumPriceCadenceMonthly    NewSubscriptionPriceModelNewSubscriptionTieredWithMinimumPriceCadence = "monthly"
	NewSubscriptionPriceModelNewSubscriptionTieredWithMinimumPriceCadenceQuarterly  NewSubscriptionPriceModelNewSubscriptionTieredWithMinimumPriceCadence = "quarterly"
	NewSubscriptionPriceModelNewSubscriptionTieredWithMinimumPriceCadenceOneTime    NewSubscriptionPriceModelNewSubscriptionTieredWithMinimumPriceCadence = "one_time"
	NewSubscriptionPriceModelNewSubscriptionTieredWithMinimumPriceCadenceCustom     NewSubscriptionPriceModelNewSubscriptionTieredWithMinimumPriceCadence = "custom"
)

func (r NewSubscriptionPriceModelNewSubscriptionTieredWithMinimumPriceCadence) IsKnown() bool {
	switch r {
	case NewSubscriptionPriceModelNewSubscriptionTieredWithMinimumPriceCadenceAnnual, NewSubscriptionPriceModelNewSubscriptionTieredWithMinimumPriceCadenceSemiAnnual, NewSubscriptionPriceModelNewSubscriptionTieredWithMinimumPriceCadenceMonthly, NewSubscriptionPriceModelNewSubscriptionTieredWithMinimumPriceCadenceQuarterly, NewSubscriptionPriceModelNewSubscriptionTieredWithMinimumPriceCadenceOneTime, NewSubscriptionPriceModelNewSubscriptionTieredWithMinimumPriceCadenceCustom:
		return true
	}
	return false
}

type NewSubscriptionPriceModelNewSubscriptionTieredWithMinimumPriceModelType string

const (
	NewSubscriptionPriceModelNewSubscriptionTieredWithMinimumPriceModelTypeTieredWithMinimum NewSubscriptionPriceModelNewSubscriptionTieredWithMinimumPriceModelType = "tiered_with_minimum"
)

func (r NewSubscriptionPriceModelNewSubscriptionTieredWithMinimumPriceModelType) IsKnown() bool {
	switch r {
	case NewSubscriptionPriceModelNewSubscriptionTieredWithMinimumPriceModelTypeTieredWithMinimum:
		return true
	}
	return false
}

type NewSubscriptionPriceModelNewSubscriptionUnitWithPercentPriceParam struct {
	// The cadence to bill for this price on.
	Cadence param.Field[NewSubscriptionPriceModelNewSubscriptionUnitWithPercentPriceCadence] `json:"cadence,required"`
	// The id of the item the price will be associated with.
	ItemID    param.Field[string]                                                                `json:"item_id,required"`
	ModelType param.Field[NewSubscriptionPriceModelNewSubscriptionUnitWithPercentPriceModelType] `json:"model_type,required"`
	// The name of the price.
	Name                  param.Field[string]                               `json:"name,required"`
	UnitWithPercentConfig param.Field[CustomRatingFunctionConfigModelParam] `json:"unit_with_percent_config,required"`
	// The id of the billable metric for the price. Only needed if the price is
	// usage-based.
	BillableMetricID param.Field[string] `json:"billable_metric_id"`
	// If the Price represents a fixed cost, the price will be billed in-advance if
	// this is true, and in-arrears if this is false.
	BilledInAdvance param.Field[bool] `json:"billed_in_advance"`
	// For custom cadence: specifies the duration of the billing period in days or
	// months.
	BillingCycleConfiguration param.Field[NewBillingCycleConfigurationModelParam] `json:"billing_cycle_configuration"`
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
	InvoicingCycleConfiguration param.Field[NewBillingCycleConfigurationModelParam] `json:"invoicing_cycle_configuration"`
	// User-specified key/value pairs for the resource. Individual keys can be removed
	// by setting the value to `null`, and the entire metadata mapping can be cleared
	// by setting `metadata` to `null`.
	Metadata param.Field[map[string]string] `json:"metadata"`
	// A transient ID that can be used to reference this price when adding adjustments
	// in the same API call.
	ReferenceID param.Field[string] `json:"reference_id"`
}

func (r NewSubscriptionPriceModelNewSubscriptionUnitWithPercentPriceParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r NewSubscriptionPriceModelNewSubscriptionUnitWithPercentPriceParam) implementsNewSubscriptionPriceModelUnionParam() {
}

// The cadence to bill for this price on.
type NewSubscriptionPriceModelNewSubscriptionUnitWithPercentPriceCadence string

const (
	NewSubscriptionPriceModelNewSubscriptionUnitWithPercentPriceCadenceAnnual     NewSubscriptionPriceModelNewSubscriptionUnitWithPercentPriceCadence = "annual"
	NewSubscriptionPriceModelNewSubscriptionUnitWithPercentPriceCadenceSemiAnnual NewSubscriptionPriceModelNewSubscriptionUnitWithPercentPriceCadence = "semi_annual"
	NewSubscriptionPriceModelNewSubscriptionUnitWithPercentPriceCadenceMonthly    NewSubscriptionPriceModelNewSubscriptionUnitWithPercentPriceCadence = "monthly"
	NewSubscriptionPriceModelNewSubscriptionUnitWithPercentPriceCadenceQuarterly  NewSubscriptionPriceModelNewSubscriptionUnitWithPercentPriceCadence = "quarterly"
	NewSubscriptionPriceModelNewSubscriptionUnitWithPercentPriceCadenceOneTime    NewSubscriptionPriceModelNewSubscriptionUnitWithPercentPriceCadence = "one_time"
	NewSubscriptionPriceModelNewSubscriptionUnitWithPercentPriceCadenceCustom     NewSubscriptionPriceModelNewSubscriptionUnitWithPercentPriceCadence = "custom"
)

func (r NewSubscriptionPriceModelNewSubscriptionUnitWithPercentPriceCadence) IsKnown() bool {
	switch r {
	case NewSubscriptionPriceModelNewSubscriptionUnitWithPercentPriceCadenceAnnual, NewSubscriptionPriceModelNewSubscriptionUnitWithPercentPriceCadenceSemiAnnual, NewSubscriptionPriceModelNewSubscriptionUnitWithPercentPriceCadenceMonthly, NewSubscriptionPriceModelNewSubscriptionUnitWithPercentPriceCadenceQuarterly, NewSubscriptionPriceModelNewSubscriptionUnitWithPercentPriceCadenceOneTime, NewSubscriptionPriceModelNewSubscriptionUnitWithPercentPriceCadenceCustom:
		return true
	}
	return false
}

type NewSubscriptionPriceModelNewSubscriptionUnitWithPercentPriceModelType string

const (
	NewSubscriptionPriceModelNewSubscriptionUnitWithPercentPriceModelTypeUnitWithPercent NewSubscriptionPriceModelNewSubscriptionUnitWithPercentPriceModelType = "unit_with_percent"
)

func (r NewSubscriptionPriceModelNewSubscriptionUnitWithPercentPriceModelType) IsKnown() bool {
	switch r {
	case NewSubscriptionPriceModelNewSubscriptionUnitWithPercentPriceModelTypeUnitWithPercent:
		return true
	}
	return false
}

type NewSubscriptionPriceModelNewSubscriptionPackageWithAllocationPriceParam struct {
	// The cadence to bill for this price on.
	Cadence param.Field[NewSubscriptionPriceModelNewSubscriptionPackageWithAllocationPriceCadence] `json:"cadence,required"`
	// The id of the item the price will be associated with.
	ItemID    param.Field[string]                                                                      `json:"item_id,required"`
	ModelType param.Field[NewSubscriptionPriceModelNewSubscriptionPackageWithAllocationPriceModelType] `json:"model_type,required"`
	// The name of the price.
	Name                        param.Field[string]                               `json:"name,required"`
	PackageWithAllocationConfig param.Field[CustomRatingFunctionConfigModelParam] `json:"package_with_allocation_config,required"`
	// The id of the billable metric for the price. Only needed if the price is
	// usage-based.
	BillableMetricID param.Field[string] `json:"billable_metric_id"`
	// If the Price represents a fixed cost, the price will be billed in-advance if
	// this is true, and in-arrears if this is false.
	BilledInAdvance param.Field[bool] `json:"billed_in_advance"`
	// For custom cadence: specifies the duration of the billing period in days or
	// months.
	BillingCycleConfiguration param.Field[NewBillingCycleConfigurationModelParam] `json:"billing_cycle_configuration"`
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
	InvoicingCycleConfiguration param.Field[NewBillingCycleConfigurationModelParam] `json:"invoicing_cycle_configuration"`
	// User-specified key/value pairs for the resource. Individual keys can be removed
	// by setting the value to `null`, and the entire metadata mapping can be cleared
	// by setting `metadata` to `null`.
	Metadata param.Field[map[string]string] `json:"metadata"`
	// A transient ID that can be used to reference this price when adding adjustments
	// in the same API call.
	ReferenceID param.Field[string] `json:"reference_id"`
}

func (r NewSubscriptionPriceModelNewSubscriptionPackageWithAllocationPriceParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r NewSubscriptionPriceModelNewSubscriptionPackageWithAllocationPriceParam) implementsNewSubscriptionPriceModelUnionParam() {
}

// The cadence to bill for this price on.
type NewSubscriptionPriceModelNewSubscriptionPackageWithAllocationPriceCadence string

const (
	NewSubscriptionPriceModelNewSubscriptionPackageWithAllocationPriceCadenceAnnual     NewSubscriptionPriceModelNewSubscriptionPackageWithAllocationPriceCadence = "annual"
	NewSubscriptionPriceModelNewSubscriptionPackageWithAllocationPriceCadenceSemiAnnual NewSubscriptionPriceModelNewSubscriptionPackageWithAllocationPriceCadence = "semi_annual"
	NewSubscriptionPriceModelNewSubscriptionPackageWithAllocationPriceCadenceMonthly    NewSubscriptionPriceModelNewSubscriptionPackageWithAllocationPriceCadence = "monthly"
	NewSubscriptionPriceModelNewSubscriptionPackageWithAllocationPriceCadenceQuarterly  NewSubscriptionPriceModelNewSubscriptionPackageWithAllocationPriceCadence = "quarterly"
	NewSubscriptionPriceModelNewSubscriptionPackageWithAllocationPriceCadenceOneTime    NewSubscriptionPriceModelNewSubscriptionPackageWithAllocationPriceCadence = "one_time"
	NewSubscriptionPriceModelNewSubscriptionPackageWithAllocationPriceCadenceCustom     NewSubscriptionPriceModelNewSubscriptionPackageWithAllocationPriceCadence = "custom"
)

func (r NewSubscriptionPriceModelNewSubscriptionPackageWithAllocationPriceCadence) IsKnown() bool {
	switch r {
	case NewSubscriptionPriceModelNewSubscriptionPackageWithAllocationPriceCadenceAnnual, NewSubscriptionPriceModelNewSubscriptionPackageWithAllocationPriceCadenceSemiAnnual, NewSubscriptionPriceModelNewSubscriptionPackageWithAllocationPriceCadenceMonthly, NewSubscriptionPriceModelNewSubscriptionPackageWithAllocationPriceCadenceQuarterly, NewSubscriptionPriceModelNewSubscriptionPackageWithAllocationPriceCadenceOneTime, NewSubscriptionPriceModelNewSubscriptionPackageWithAllocationPriceCadenceCustom:
		return true
	}
	return false
}

type NewSubscriptionPriceModelNewSubscriptionPackageWithAllocationPriceModelType string

const (
	NewSubscriptionPriceModelNewSubscriptionPackageWithAllocationPriceModelTypePackageWithAllocation NewSubscriptionPriceModelNewSubscriptionPackageWithAllocationPriceModelType = "package_with_allocation"
)

func (r NewSubscriptionPriceModelNewSubscriptionPackageWithAllocationPriceModelType) IsKnown() bool {
	switch r {
	case NewSubscriptionPriceModelNewSubscriptionPackageWithAllocationPriceModelTypePackageWithAllocation:
		return true
	}
	return false
}

type NewSubscriptionPriceModelNewSubscriptionTierWithProrationPriceParam struct {
	// The cadence to bill for this price on.
	Cadence param.Field[NewSubscriptionPriceModelNewSubscriptionTierWithProrationPriceCadence] `json:"cadence,required"`
	// The id of the item the price will be associated with.
	ItemID    param.Field[string]                                                                  `json:"item_id,required"`
	ModelType param.Field[NewSubscriptionPriceModelNewSubscriptionTierWithProrationPriceModelType] `json:"model_type,required"`
	// The name of the price.
	Name                      param.Field[string]                               `json:"name,required"`
	TieredWithProrationConfig param.Field[CustomRatingFunctionConfigModelParam] `json:"tiered_with_proration_config,required"`
	// The id of the billable metric for the price. Only needed if the price is
	// usage-based.
	BillableMetricID param.Field[string] `json:"billable_metric_id"`
	// If the Price represents a fixed cost, the price will be billed in-advance if
	// this is true, and in-arrears if this is false.
	BilledInAdvance param.Field[bool] `json:"billed_in_advance"`
	// For custom cadence: specifies the duration of the billing period in days or
	// months.
	BillingCycleConfiguration param.Field[NewBillingCycleConfigurationModelParam] `json:"billing_cycle_configuration"`
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
	InvoicingCycleConfiguration param.Field[NewBillingCycleConfigurationModelParam] `json:"invoicing_cycle_configuration"`
	// User-specified key/value pairs for the resource. Individual keys can be removed
	// by setting the value to `null`, and the entire metadata mapping can be cleared
	// by setting `metadata` to `null`.
	Metadata param.Field[map[string]string] `json:"metadata"`
	// A transient ID that can be used to reference this price when adding adjustments
	// in the same API call.
	ReferenceID param.Field[string] `json:"reference_id"`
}

func (r NewSubscriptionPriceModelNewSubscriptionTierWithProrationPriceParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r NewSubscriptionPriceModelNewSubscriptionTierWithProrationPriceParam) implementsNewSubscriptionPriceModelUnionParam() {
}

// The cadence to bill for this price on.
type NewSubscriptionPriceModelNewSubscriptionTierWithProrationPriceCadence string

const (
	NewSubscriptionPriceModelNewSubscriptionTierWithProrationPriceCadenceAnnual     NewSubscriptionPriceModelNewSubscriptionTierWithProrationPriceCadence = "annual"
	NewSubscriptionPriceModelNewSubscriptionTierWithProrationPriceCadenceSemiAnnual NewSubscriptionPriceModelNewSubscriptionTierWithProrationPriceCadence = "semi_annual"
	NewSubscriptionPriceModelNewSubscriptionTierWithProrationPriceCadenceMonthly    NewSubscriptionPriceModelNewSubscriptionTierWithProrationPriceCadence = "monthly"
	NewSubscriptionPriceModelNewSubscriptionTierWithProrationPriceCadenceQuarterly  NewSubscriptionPriceModelNewSubscriptionTierWithProrationPriceCadence = "quarterly"
	NewSubscriptionPriceModelNewSubscriptionTierWithProrationPriceCadenceOneTime    NewSubscriptionPriceModelNewSubscriptionTierWithProrationPriceCadence = "one_time"
	NewSubscriptionPriceModelNewSubscriptionTierWithProrationPriceCadenceCustom     NewSubscriptionPriceModelNewSubscriptionTierWithProrationPriceCadence = "custom"
)

func (r NewSubscriptionPriceModelNewSubscriptionTierWithProrationPriceCadence) IsKnown() bool {
	switch r {
	case NewSubscriptionPriceModelNewSubscriptionTierWithProrationPriceCadenceAnnual, NewSubscriptionPriceModelNewSubscriptionTierWithProrationPriceCadenceSemiAnnual, NewSubscriptionPriceModelNewSubscriptionTierWithProrationPriceCadenceMonthly, NewSubscriptionPriceModelNewSubscriptionTierWithProrationPriceCadenceQuarterly, NewSubscriptionPriceModelNewSubscriptionTierWithProrationPriceCadenceOneTime, NewSubscriptionPriceModelNewSubscriptionTierWithProrationPriceCadenceCustom:
		return true
	}
	return false
}

type NewSubscriptionPriceModelNewSubscriptionTierWithProrationPriceModelType string

const (
	NewSubscriptionPriceModelNewSubscriptionTierWithProrationPriceModelTypeTieredWithProration NewSubscriptionPriceModelNewSubscriptionTierWithProrationPriceModelType = "tiered_with_proration"
)

func (r NewSubscriptionPriceModelNewSubscriptionTierWithProrationPriceModelType) IsKnown() bool {
	switch r {
	case NewSubscriptionPriceModelNewSubscriptionTierWithProrationPriceModelTypeTieredWithProration:
		return true
	}
	return false
}

type NewSubscriptionPriceModelNewSubscriptionUnitWithProrationPriceParam struct {
	// The cadence to bill for this price on.
	Cadence param.Field[NewSubscriptionPriceModelNewSubscriptionUnitWithProrationPriceCadence] `json:"cadence,required"`
	// The id of the item the price will be associated with.
	ItemID    param.Field[string]                                                                  `json:"item_id,required"`
	ModelType param.Field[NewSubscriptionPriceModelNewSubscriptionUnitWithProrationPriceModelType] `json:"model_type,required"`
	// The name of the price.
	Name                    param.Field[string]                               `json:"name,required"`
	UnitWithProrationConfig param.Field[CustomRatingFunctionConfigModelParam] `json:"unit_with_proration_config,required"`
	// The id of the billable metric for the price. Only needed if the price is
	// usage-based.
	BillableMetricID param.Field[string] `json:"billable_metric_id"`
	// If the Price represents a fixed cost, the price will be billed in-advance if
	// this is true, and in-arrears if this is false.
	BilledInAdvance param.Field[bool] `json:"billed_in_advance"`
	// For custom cadence: specifies the duration of the billing period in days or
	// months.
	BillingCycleConfiguration param.Field[NewBillingCycleConfigurationModelParam] `json:"billing_cycle_configuration"`
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
	InvoicingCycleConfiguration param.Field[NewBillingCycleConfigurationModelParam] `json:"invoicing_cycle_configuration"`
	// User-specified key/value pairs for the resource. Individual keys can be removed
	// by setting the value to `null`, and the entire metadata mapping can be cleared
	// by setting `metadata` to `null`.
	Metadata param.Field[map[string]string] `json:"metadata"`
	// A transient ID that can be used to reference this price when adding adjustments
	// in the same API call.
	ReferenceID param.Field[string] `json:"reference_id"`
}

func (r NewSubscriptionPriceModelNewSubscriptionUnitWithProrationPriceParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r NewSubscriptionPriceModelNewSubscriptionUnitWithProrationPriceParam) implementsNewSubscriptionPriceModelUnionParam() {
}

// The cadence to bill for this price on.
type NewSubscriptionPriceModelNewSubscriptionUnitWithProrationPriceCadence string

const (
	NewSubscriptionPriceModelNewSubscriptionUnitWithProrationPriceCadenceAnnual     NewSubscriptionPriceModelNewSubscriptionUnitWithProrationPriceCadence = "annual"
	NewSubscriptionPriceModelNewSubscriptionUnitWithProrationPriceCadenceSemiAnnual NewSubscriptionPriceModelNewSubscriptionUnitWithProrationPriceCadence = "semi_annual"
	NewSubscriptionPriceModelNewSubscriptionUnitWithProrationPriceCadenceMonthly    NewSubscriptionPriceModelNewSubscriptionUnitWithProrationPriceCadence = "monthly"
	NewSubscriptionPriceModelNewSubscriptionUnitWithProrationPriceCadenceQuarterly  NewSubscriptionPriceModelNewSubscriptionUnitWithProrationPriceCadence = "quarterly"
	NewSubscriptionPriceModelNewSubscriptionUnitWithProrationPriceCadenceOneTime    NewSubscriptionPriceModelNewSubscriptionUnitWithProrationPriceCadence = "one_time"
	NewSubscriptionPriceModelNewSubscriptionUnitWithProrationPriceCadenceCustom     NewSubscriptionPriceModelNewSubscriptionUnitWithProrationPriceCadence = "custom"
)

func (r NewSubscriptionPriceModelNewSubscriptionUnitWithProrationPriceCadence) IsKnown() bool {
	switch r {
	case NewSubscriptionPriceModelNewSubscriptionUnitWithProrationPriceCadenceAnnual, NewSubscriptionPriceModelNewSubscriptionUnitWithProrationPriceCadenceSemiAnnual, NewSubscriptionPriceModelNewSubscriptionUnitWithProrationPriceCadenceMonthly, NewSubscriptionPriceModelNewSubscriptionUnitWithProrationPriceCadenceQuarterly, NewSubscriptionPriceModelNewSubscriptionUnitWithProrationPriceCadenceOneTime, NewSubscriptionPriceModelNewSubscriptionUnitWithProrationPriceCadenceCustom:
		return true
	}
	return false
}

type NewSubscriptionPriceModelNewSubscriptionUnitWithProrationPriceModelType string

const (
	NewSubscriptionPriceModelNewSubscriptionUnitWithProrationPriceModelTypeUnitWithProration NewSubscriptionPriceModelNewSubscriptionUnitWithProrationPriceModelType = "unit_with_proration"
)

func (r NewSubscriptionPriceModelNewSubscriptionUnitWithProrationPriceModelType) IsKnown() bool {
	switch r {
	case NewSubscriptionPriceModelNewSubscriptionUnitWithProrationPriceModelTypeUnitWithProration:
		return true
	}
	return false
}

type NewSubscriptionPriceModelNewSubscriptionGroupedAllocationPriceParam struct {
	// The cadence to bill for this price on.
	Cadence                 param.Field[NewSubscriptionPriceModelNewSubscriptionGroupedAllocationPriceCadence] `json:"cadence,required"`
	GroupedAllocationConfig param.Field[CustomRatingFunctionConfigModelParam]                                  `json:"grouped_allocation_config,required"`
	// The id of the item the price will be associated with.
	ItemID    param.Field[string]                                                                  `json:"item_id,required"`
	ModelType param.Field[NewSubscriptionPriceModelNewSubscriptionGroupedAllocationPriceModelType] `json:"model_type,required"`
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
	BillingCycleConfiguration param.Field[NewBillingCycleConfigurationModelParam] `json:"billing_cycle_configuration"`
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
	InvoicingCycleConfiguration param.Field[NewBillingCycleConfigurationModelParam] `json:"invoicing_cycle_configuration"`
	// User-specified key/value pairs for the resource. Individual keys can be removed
	// by setting the value to `null`, and the entire metadata mapping can be cleared
	// by setting `metadata` to `null`.
	Metadata param.Field[map[string]string] `json:"metadata"`
	// A transient ID that can be used to reference this price when adding adjustments
	// in the same API call.
	ReferenceID param.Field[string] `json:"reference_id"`
}

func (r NewSubscriptionPriceModelNewSubscriptionGroupedAllocationPriceParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r NewSubscriptionPriceModelNewSubscriptionGroupedAllocationPriceParam) implementsNewSubscriptionPriceModelUnionParam() {
}

// The cadence to bill for this price on.
type NewSubscriptionPriceModelNewSubscriptionGroupedAllocationPriceCadence string

const (
	NewSubscriptionPriceModelNewSubscriptionGroupedAllocationPriceCadenceAnnual     NewSubscriptionPriceModelNewSubscriptionGroupedAllocationPriceCadence = "annual"
	NewSubscriptionPriceModelNewSubscriptionGroupedAllocationPriceCadenceSemiAnnual NewSubscriptionPriceModelNewSubscriptionGroupedAllocationPriceCadence = "semi_annual"
	NewSubscriptionPriceModelNewSubscriptionGroupedAllocationPriceCadenceMonthly    NewSubscriptionPriceModelNewSubscriptionGroupedAllocationPriceCadence = "monthly"
	NewSubscriptionPriceModelNewSubscriptionGroupedAllocationPriceCadenceQuarterly  NewSubscriptionPriceModelNewSubscriptionGroupedAllocationPriceCadence = "quarterly"
	NewSubscriptionPriceModelNewSubscriptionGroupedAllocationPriceCadenceOneTime    NewSubscriptionPriceModelNewSubscriptionGroupedAllocationPriceCadence = "one_time"
	NewSubscriptionPriceModelNewSubscriptionGroupedAllocationPriceCadenceCustom     NewSubscriptionPriceModelNewSubscriptionGroupedAllocationPriceCadence = "custom"
)

func (r NewSubscriptionPriceModelNewSubscriptionGroupedAllocationPriceCadence) IsKnown() bool {
	switch r {
	case NewSubscriptionPriceModelNewSubscriptionGroupedAllocationPriceCadenceAnnual, NewSubscriptionPriceModelNewSubscriptionGroupedAllocationPriceCadenceSemiAnnual, NewSubscriptionPriceModelNewSubscriptionGroupedAllocationPriceCadenceMonthly, NewSubscriptionPriceModelNewSubscriptionGroupedAllocationPriceCadenceQuarterly, NewSubscriptionPriceModelNewSubscriptionGroupedAllocationPriceCadenceOneTime, NewSubscriptionPriceModelNewSubscriptionGroupedAllocationPriceCadenceCustom:
		return true
	}
	return false
}

type NewSubscriptionPriceModelNewSubscriptionGroupedAllocationPriceModelType string

const (
	NewSubscriptionPriceModelNewSubscriptionGroupedAllocationPriceModelTypeGroupedAllocation NewSubscriptionPriceModelNewSubscriptionGroupedAllocationPriceModelType = "grouped_allocation"
)

func (r NewSubscriptionPriceModelNewSubscriptionGroupedAllocationPriceModelType) IsKnown() bool {
	switch r {
	case NewSubscriptionPriceModelNewSubscriptionGroupedAllocationPriceModelTypeGroupedAllocation:
		return true
	}
	return false
}

type NewSubscriptionPriceModelNewSubscriptionGroupedWithProratedMinimumPriceParam struct {
	// The cadence to bill for this price on.
	Cadence                          param.Field[NewSubscriptionPriceModelNewSubscriptionGroupedWithProratedMinimumPriceCadence] `json:"cadence,required"`
	GroupedWithProratedMinimumConfig param.Field[CustomRatingFunctionConfigModelParam]                                           `json:"grouped_with_prorated_minimum_config,required"`
	// The id of the item the price will be associated with.
	ItemID    param.Field[string]                                                                           `json:"item_id,required"`
	ModelType param.Field[NewSubscriptionPriceModelNewSubscriptionGroupedWithProratedMinimumPriceModelType] `json:"model_type,required"`
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
	BillingCycleConfiguration param.Field[NewBillingCycleConfigurationModelParam] `json:"billing_cycle_configuration"`
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
	InvoicingCycleConfiguration param.Field[NewBillingCycleConfigurationModelParam] `json:"invoicing_cycle_configuration"`
	// User-specified key/value pairs for the resource. Individual keys can be removed
	// by setting the value to `null`, and the entire metadata mapping can be cleared
	// by setting `metadata` to `null`.
	Metadata param.Field[map[string]string] `json:"metadata"`
	// A transient ID that can be used to reference this price when adding adjustments
	// in the same API call.
	ReferenceID param.Field[string] `json:"reference_id"`
}

func (r NewSubscriptionPriceModelNewSubscriptionGroupedWithProratedMinimumPriceParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r NewSubscriptionPriceModelNewSubscriptionGroupedWithProratedMinimumPriceParam) implementsNewSubscriptionPriceModelUnionParam() {
}

// The cadence to bill for this price on.
type NewSubscriptionPriceModelNewSubscriptionGroupedWithProratedMinimumPriceCadence string

const (
	NewSubscriptionPriceModelNewSubscriptionGroupedWithProratedMinimumPriceCadenceAnnual     NewSubscriptionPriceModelNewSubscriptionGroupedWithProratedMinimumPriceCadence = "annual"
	NewSubscriptionPriceModelNewSubscriptionGroupedWithProratedMinimumPriceCadenceSemiAnnual NewSubscriptionPriceModelNewSubscriptionGroupedWithProratedMinimumPriceCadence = "semi_annual"
	NewSubscriptionPriceModelNewSubscriptionGroupedWithProratedMinimumPriceCadenceMonthly    NewSubscriptionPriceModelNewSubscriptionGroupedWithProratedMinimumPriceCadence = "monthly"
	NewSubscriptionPriceModelNewSubscriptionGroupedWithProratedMinimumPriceCadenceQuarterly  NewSubscriptionPriceModelNewSubscriptionGroupedWithProratedMinimumPriceCadence = "quarterly"
	NewSubscriptionPriceModelNewSubscriptionGroupedWithProratedMinimumPriceCadenceOneTime    NewSubscriptionPriceModelNewSubscriptionGroupedWithProratedMinimumPriceCadence = "one_time"
	NewSubscriptionPriceModelNewSubscriptionGroupedWithProratedMinimumPriceCadenceCustom     NewSubscriptionPriceModelNewSubscriptionGroupedWithProratedMinimumPriceCadence = "custom"
)

func (r NewSubscriptionPriceModelNewSubscriptionGroupedWithProratedMinimumPriceCadence) IsKnown() bool {
	switch r {
	case NewSubscriptionPriceModelNewSubscriptionGroupedWithProratedMinimumPriceCadenceAnnual, NewSubscriptionPriceModelNewSubscriptionGroupedWithProratedMinimumPriceCadenceSemiAnnual, NewSubscriptionPriceModelNewSubscriptionGroupedWithProratedMinimumPriceCadenceMonthly, NewSubscriptionPriceModelNewSubscriptionGroupedWithProratedMinimumPriceCadenceQuarterly, NewSubscriptionPriceModelNewSubscriptionGroupedWithProratedMinimumPriceCadenceOneTime, NewSubscriptionPriceModelNewSubscriptionGroupedWithProratedMinimumPriceCadenceCustom:
		return true
	}
	return false
}

type NewSubscriptionPriceModelNewSubscriptionGroupedWithProratedMinimumPriceModelType string

const (
	NewSubscriptionPriceModelNewSubscriptionGroupedWithProratedMinimumPriceModelTypeGroupedWithProratedMinimum NewSubscriptionPriceModelNewSubscriptionGroupedWithProratedMinimumPriceModelType = "grouped_with_prorated_minimum"
)

func (r NewSubscriptionPriceModelNewSubscriptionGroupedWithProratedMinimumPriceModelType) IsKnown() bool {
	switch r {
	case NewSubscriptionPriceModelNewSubscriptionGroupedWithProratedMinimumPriceModelTypeGroupedWithProratedMinimum:
		return true
	}
	return false
}

type NewSubscriptionPriceModelNewSubscriptionBulkWithProrationPriceParam struct {
	BulkWithProrationConfig param.Field[CustomRatingFunctionConfigModelParam] `json:"bulk_with_proration_config,required"`
	// The cadence to bill for this price on.
	Cadence param.Field[NewSubscriptionPriceModelNewSubscriptionBulkWithProrationPriceCadence] `json:"cadence,required"`
	// The id of the item the price will be associated with.
	ItemID    param.Field[string]                                                                  `json:"item_id,required"`
	ModelType param.Field[NewSubscriptionPriceModelNewSubscriptionBulkWithProrationPriceModelType] `json:"model_type,required"`
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
	BillingCycleConfiguration param.Field[NewBillingCycleConfigurationModelParam] `json:"billing_cycle_configuration"`
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
	InvoicingCycleConfiguration param.Field[NewBillingCycleConfigurationModelParam] `json:"invoicing_cycle_configuration"`
	// User-specified key/value pairs for the resource. Individual keys can be removed
	// by setting the value to `null`, and the entire metadata mapping can be cleared
	// by setting `metadata` to `null`.
	Metadata param.Field[map[string]string] `json:"metadata"`
	// A transient ID that can be used to reference this price when adding adjustments
	// in the same API call.
	ReferenceID param.Field[string] `json:"reference_id"`
}

func (r NewSubscriptionPriceModelNewSubscriptionBulkWithProrationPriceParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r NewSubscriptionPriceModelNewSubscriptionBulkWithProrationPriceParam) implementsNewSubscriptionPriceModelUnionParam() {
}

// The cadence to bill for this price on.
type NewSubscriptionPriceModelNewSubscriptionBulkWithProrationPriceCadence string

const (
	NewSubscriptionPriceModelNewSubscriptionBulkWithProrationPriceCadenceAnnual     NewSubscriptionPriceModelNewSubscriptionBulkWithProrationPriceCadence = "annual"
	NewSubscriptionPriceModelNewSubscriptionBulkWithProrationPriceCadenceSemiAnnual NewSubscriptionPriceModelNewSubscriptionBulkWithProrationPriceCadence = "semi_annual"
	NewSubscriptionPriceModelNewSubscriptionBulkWithProrationPriceCadenceMonthly    NewSubscriptionPriceModelNewSubscriptionBulkWithProrationPriceCadence = "monthly"
	NewSubscriptionPriceModelNewSubscriptionBulkWithProrationPriceCadenceQuarterly  NewSubscriptionPriceModelNewSubscriptionBulkWithProrationPriceCadence = "quarterly"
	NewSubscriptionPriceModelNewSubscriptionBulkWithProrationPriceCadenceOneTime    NewSubscriptionPriceModelNewSubscriptionBulkWithProrationPriceCadence = "one_time"
	NewSubscriptionPriceModelNewSubscriptionBulkWithProrationPriceCadenceCustom     NewSubscriptionPriceModelNewSubscriptionBulkWithProrationPriceCadence = "custom"
)

func (r NewSubscriptionPriceModelNewSubscriptionBulkWithProrationPriceCadence) IsKnown() bool {
	switch r {
	case NewSubscriptionPriceModelNewSubscriptionBulkWithProrationPriceCadenceAnnual, NewSubscriptionPriceModelNewSubscriptionBulkWithProrationPriceCadenceSemiAnnual, NewSubscriptionPriceModelNewSubscriptionBulkWithProrationPriceCadenceMonthly, NewSubscriptionPriceModelNewSubscriptionBulkWithProrationPriceCadenceQuarterly, NewSubscriptionPriceModelNewSubscriptionBulkWithProrationPriceCadenceOneTime, NewSubscriptionPriceModelNewSubscriptionBulkWithProrationPriceCadenceCustom:
		return true
	}
	return false
}

type NewSubscriptionPriceModelNewSubscriptionBulkWithProrationPriceModelType string

const (
	NewSubscriptionPriceModelNewSubscriptionBulkWithProrationPriceModelTypeBulkWithProration NewSubscriptionPriceModelNewSubscriptionBulkWithProrationPriceModelType = "bulk_with_proration"
)

func (r NewSubscriptionPriceModelNewSubscriptionBulkWithProrationPriceModelType) IsKnown() bool {
	switch r {
	case NewSubscriptionPriceModelNewSubscriptionBulkWithProrationPriceModelTypeBulkWithProration:
		return true
	}
	return false
}

type NewSubscriptionPriceModelNewSubscriptionScalableMatrixWithUnitPricingPriceParam struct {
	// The cadence to bill for this price on.
	Cadence param.Field[NewSubscriptionPriceModelNewSubscriptionScalableMatrixWithUnitPricingPriceCadence] `json:"cadence,required"`
	// The id of the item the price will be associated with.
	ItemID    param.Field[string]                                                                              `json:"item_id,required"`
	ModelType param.Field[NewSubscriptionPriceModelNewSubscriptionScalableMatrixWithUnitPricingPriceModelType] `json:"model_type,required"`
	// The name of the price.
	Name                                param.Field[string]                               `json:"name,required"`
	ScalableMatrixWithUnitPricingConfig param.Field[CustomRatingFunctionConfigModelParam] `json:"scalable_matrix_with_unit_pricing_config,required"`
	// The id of the billable metric for the price. Only needed if the price is
	// usage-based.
	BillableMetricID param.Field[string] `json:"billable_metric_id"`
	// If the Price represents a fixed cost, the price will be billed in-advance if
	// this is true, and in-arrears if this is false.
	BilledInAdvance param.Field[bool] `json:"billed_in_advance"`
	// For custom cadence: specifies the duration of the billing period in days or
	// months.
	BillingCycleConfiguration param.Field[NewBillingCycleConfigurationModelParam] `json:"billing_cycle_configuration"`
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
	InvoicingCycleConfiguration param.Field[NewBillingCycleConfigurationModelParam] `json:"invoicing_cycle_configuration"`
	// User-specified key/value pairs for the resource. Individual keys can be removed
	// by setting the value to `null`, and the entire metadata mapping can be cleared
	// by setting `metadata` to `null`.
	Metadata param.Field[map[string]string] `json:"metadata"`
	// A transient ID that can be used to reference this price when adding adjustments
	// in the same API call.
	ReferenceID param.Field[string] `json:"reference_id"`
}

func (r NewSubscriptionPriceModelNewSubscriptionScalableMatrixWithUnitPricingPriceParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r NewSubscriptionPriceModelNewSubscriptionScalableMatrixWithUnitPricingPriceParam) implementsNewSubscriptionPriceModelUnionParam() {
}

// The cadence to bill for this price on.
type NewSubscriptionPriceModelNewSubscriptionScalableMatrixWithUnitPricingPriceCadence string

const (
	NewSubscriptionPriceModelNewSubscriptionScalableMatrixWithUnitPricingPriceCadenceAnnual     NewSubscriptionPriceModelNewSubscriptionScalableMatrixWithUnitPricingPriceCadence = "annual"
	NewSubscriptionPriceModelNewSubscriptionScalableMatrixWithUnitPricingPriceCadenceSemiAnnual NewSubscriptionPriceModelNewSubscriptionScalableMatrixWithUnitPricingPriceCadence = "semi_annual"
	NewSubscriptionPriceModelNewSubscriptionScalableMatrixWithUnitPricingPriceCadenceMonthly    NewSubscriptionPriceModelNewSubscriptionScalableMatrixWithUnitPricingPriceCadence = "monthly"
	NewSubscriptionPriceModelNewSubscriptionScalableMatrixWithUnitPricingPriceCadenceQuarterly  NewSubscriptionPriceModelNewSubscriptionScalableMatrixWithUnitPricingPriceCadence = "quarterly"
	NewSubscriptionPriceModelNewSubscriptionScalableMatrixWithUnitPricingPriceCadenceOneTime    NewSubscriptionPriceModelNewSubscriptionScalableMatrixWithUnitPricingPriceCadence = "one_time"
	NewSubscriptionPriceModelNewSubscriptionScalableMatrixWithUnitPricingPriceCadenceCustom     NewSubscriptionPriceModelNewSubscriptionScalableMatrixWithUnitPricingPriceCadence = "custom"
)

func (r NewSubscriptionPriceModelNewSubscriptionScalableMatrixWithUnitPricingPriceCadence) IsKnown() bool {
	switch r {
	case NewSubscriptionPriceModelNewSubscriptionScalableMatrixWithUnitPricingPriceCadenceAnnual, NewSubscriptionPriceModelNewSubscriptionScalableMatrixWithUnitPricingPriceCadenceSemiAnnual, NewSubscriptionPriceModelNewSubscriptionScalableMatrixWithUnitPricingPriceCadenceMonthly, NewSubscriptionPriceModelNewSubscriptionScalableMatrixWithUnitPricingPriceCadenceQuarterly, NewSubscriptionPriceModelNewSubscriptionScalableMatrixWithUnitPricingPriceCadenceOneTime, NewSubscriptionPriceModelNewSubscriptionScalableMatrixWithUnitPricingPriceCadenceCustom:
		return true
	}
	return false
}

type NewSubscriptionPriceModelNewSubscriptionScalableMatrixWithUnitPricingPriceModelType string

const (
	NewSubscriptionPriceModelNewSubscriptionScalableMatrixWithUnitPricingPriceModelTypeScalableMatrixWithUnitPricing NewSubscriptionPriceModelNewSubscriptionScalableMatrixWithUnitPricingPriceModelType = "scalable_matrix_with_unit_pricing"
)

func (r NewSubscriptionPriceModelNewSubscriptionScalableMatrixWithUnitPricingPriceModelType) IsKnown() bool {
	switch r {
	case NewSubscriptionPriceModelNewSubscriptionScalableMatrixWithUnitPricingPriceModelTypeScalableMatrixWithUnitPricing:
		return true
	}
	return false
}

type NewSubscriptionPriceModelNewSubscriptionScalableMatrixWithTieredPricingPriceParam struct {
	// The cadence to bill for this price on.
	Cadence param.Field[NewSubscriptionPriceModelNewSubscriptionScalableMatrixWithTieredPricingPriceCadence] `json:"cadence,required"`
	// The id of the item the price will be associated with.
	ItemID    param.Field[string]                                                                                `json:"item_id,required"`
	ModelType param.Field[NewSubscriptionPriceModelNewSubscriptionScalableMatrixWithTieredPricingPriceModelType] `json:"model_type,required"`
	// The name of the price.
	Name                                  param.Field[string]                               `json:"name,required"`
	ScalableMatrixWithTieredPricingConfig param.Field[CustomRatingFunctionConfigModelParam] `json:"scalable_matrix_with_tiered_pricing_config,required"`
	// The id of the billable metric for the price. Only needed if the price is
	// usage-based.
	BillableMetricID param.Field[string] `json:"billable_metric_id"`
	// If the Price represents a fixed cost, the price will be billed in-advance if
	// this is true, and in-arrears if this is false.
	BilledInAdvance param.Field[bool] `json:"billed_in_advance"`
	// For custom cadence: specifies the duration of the billing period in days or
	// months.
	BillingCycleConfiguration param.Field[NewBillingCycleConfigurationModelParam] `json:"billing_cycle_configuration"`
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
	InvoicingCycleConfiguration param.Field[NewBillingCycleConfigurationModelParam] `json:"invoicing_cycle_configuration"`
	// User-specified key/value pairs for the resource. Individual keys can be removed
	// by setting the value to `null`, and the entire metadata mapping can be cleared
	// by setting `metadata` to `null`.
	Metadata param.Field[map[string]string] `json:"metadata"`
	// A transient ID that can be used to reference this price when adding adjustments
	// in the same API call.
	ReferenceID param.Field[string] `json:"reference_id"`
}

func (r NewSubscriptionPriceModelNewSubscriptionScalableMatrixWithTieredPricingPriceParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r NewSubscriptionPriceModelNewSubscriptionScalableMatrixWithTieredPricingPriceParam) implementsNewSubscriptionPriceModelUnionParam() {
}

// The cadence to bill for this price on.
type NewSubscriptionPriceModelNewSubscriptionScalableMatrixWithTieredPricingPriceCadence string

const (
	NewSubscriptionPriceModelNewSubscriptionScalableMatrixWithTieredPricingPriceCadenceAnnual     NewSubscriptionPriceModelNewSubscriptionScalableMatrixWithTieredPricingPriceCadence = "annual"
	NewSubscriptionPriceModelNewSubscriptionScalableMatrixWithTieredPricingPriceCadenceSemiAnnual NewSubscriptionPriceModelNewSubscriptionScalableMatrixWithTieredPricingPriceCadence = "semi_annual"
	NewSubscriptionPriceModelNewSubscriptionScalableMatrixWithTieredPricingPriceCadenceMonthly    NewSubscriptionPriceModelNewSubscriptionScalableMatrixWithTieredPricingPriceCadence = "monthly"
	NewSubscriptionPriceModelNewSubscriptionScalableMatrixWithTieredPricingPriceCadenceQuarterly  NewSubscriptionPriceModelNewSubscriptionScalableMatrixWithTieredPricingPriceCadence = "quarterly"
	NewSubscriptionPriceModelNewSubscriptionScalableMatrixWithTieredPricingPriceCadenceOneTime    NewSubscriptionPriceModelNewSubscriptionScalableMatrixWithTieredPricingPriceCadence = "one_time"
	NewSubscriptionPriceModelNewSubscriptionScalableMatrixWithTieredPricingPriceCadenceCustom     NewSubscriptionPriceModelNewSubscriptionScalableMatrixWithTieredPricingPriceCadence = "custom"
)

func (r NewSubscriptionPriceModelNewSubscriptionScalableMatrixWithTieredPricingPriceCadence) IsKnown() bool {
	switch r {
	case NewSubscriptionPriceModelNewSubscriptionScalableMatrixWithTieredPricingPriceCadenceAnnual, NewSubscriptionPriceModelNewSubscriptionScalableMatrixWithTieredPricingPriceCadenceSemiAnnual, NewSubscriptionPriceModelNewSubscriptionScalableMatrixWithTieredPricingPriceCadenceMonthly, NewSubscriptionPriceModelNewSubscriptionScalableMatrixWithTieredPricingPriceCadenceQuarterly, NewSubscriptionPriceModelNewSubscriptionScalableMatrixWithTieredPricingPriceCadenceOneTime, NewSubscriptionPriceModelNewSubscriptionScalableMatrixWithTieredPricingPriceCadenceCustom:
		return true
	}
	return false
}

type NewSubscriptionPriceModelNewSubscriptionScalableMatrixWithTieredPricingPriceModelType string

const (
	NewSubscriptionPriceModelNewSubscriptionScalableMatrixWithTieredPricingPriceModelTypeScalableMatrixWithTieredPricing NewSubscriptionPriceModelNewSubscriptionScalableMatrixWithTieredPricingPriceModelType = "scalable_matrix_with_tiered_pricing"
)

func (r NewSubscriptionPriceModelNewSubscriptionScalableMatrixWithTieredPricingPriceModelType) IsKnown() bool {
	switch r {
	case NewSubscriptionPriceModelNewSubscriptionScalableMatrixWithTieredPricingPriceModelTypeScalableMatrixWithTieredPricing:
		return true
	}
	return false
}

type NewSubscriptionPriceModelNewSubscriptionCumulativeGroupedBulkPriceParam struct {
	// The cadence to bill for this price on.
	Cadence                     param.Field[NewSubscriptionPriceModelNewSubscriptionCumulativeGroupedBulkPriceCadence] `json:"cadence,required"`
	CumulativeGroupedBulkConfig param.Field[CustomRatingFunctionConfigModelParam]                                      `json:"cumulative_grouped_bulk_config,required"`
	// The id of the item the price will be associated with.
	ItemID    param.Field[string]                                                                      `json:"item_id,required"`
	ModelType param.Field[NewSubscriptionPriceModelNewSubscriptionCumulativeGroupedBulkPriceModelType] `json:"model_type,required"`
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
	BillingCycleConfiguration param.Field[NewBillingCycleConfigurationModelParam] `json:"billing_cycle_configuration"`
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
	InvoicingCycleConfiguration param.Field[NewBillingCycleConfigurationModelParam] `json:"invoicing_cycle_configuration"`
	// User-specified key/value pairs for the resource. Individual keys can be removed
	// by setting the value to `null`, and the entire metadata mapping can be cleared
	// by setting `metadata` to `null`.
	Metadata param.Field[map[string]string] `json:"metadata"`
	// A transient ID that can be used to reference this price when adding adjustments
	// in the same API call.
	ReferenceID param.Field[string] `json:"reference_id"`
}

func (r NewSubscriptionPriceModelNewSubscriptionCumulativeGroupedBulkPriceParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r NewSubscriptionPriceModelNewSubscriptionCumulativeGroupedBulkPriceParam) implementsNewSubscriptionPriceModelUnionParam() {
}

// The cadence to bill for this price on.
type NewSubscriptionPriceModelNewSubscriptionCumulativeGroupedBulkPriceCadence string

const (
	NewSubscriptionPriceModelNewSubscriptionCumulativeGroupedBulkPriceCadenceAnnual     NewSubscriptionPriceModelNewSubscriptionCumulativeGroupedBulkPriceCadence = "annual"
	NewSubscriptionPriceModelNewSubscriptionCumulativeGroupedBulkPriceCadenceSemiAnnual NewSubscriptionPriceModelNewSubscriptionCumulativeGroupedBulkPriceCadence = "semi_annual"
	NewSubscriptionPriceModelNewSubscriptionCumulativeGroupedBulkPriceCadenceMonthly    NewSubscriptionPriceModelNewSubscriptionCumulativeGroupedBulkPriceCadence = "monthly"
	NewSubscriptionPriceModelNewSubscriptionCumulativeGroupedBulkPriceCadenceQuarterly  NewSubscriptionPriceModelNewSubscriptionCumulativeGroupedBulkPriceCadence = "quarterly"
	NewSubscriptionPriceModelNewSubscriptionCumulativeGroupedBulkPriceCadenceOneTime    NewSubscriptionPriceModelNewSubscriptionCumulativeGroupedBulkPriceCadence = "one_time"
	NewSubscriptionPriceModelNewSubscriptionCumulativeGroupedBulkPriceCadenceCustom     NewSubscriptionPriceModelNewSubscriptionCumulativeGroupedBulkPriceCadence = "custom"
)

func (r NewSubscriptionPriceModelNewSubscriptionCumulativeGroupedBulkPriceCadence) IsKnown() bool {
	switch r {
	case NewSubscriptionPriceModelNewSubscriptionCumulativeGroupedBulkPriceCadenceAnnual, NewSubscriptionPriceModelNewSubscriptionCumulativeGroupedBulkPriceCadenceSemiAnnual, NewSubscriptionPriceModelNewSubscriptionCumulativeGroupedBulkPriceCadenceMonthly, NewSubscriptionPriceModelNewSubscriptionCumulativeGroupedBulkPriceCadenceQuarterly, NewSubscriptionPriceModelNewSubscriptionCumulativeGroupedBulkPriceCadenceOneTime, NewSubscriptionPriceModelNewSubscriptionCumulativeGroupedBulkPriceCadenceCustom:
		return true
	}
	return false
}

type NewSubscriptionPriceModelNewSubscriptionCumulativeGroupedBulkPriceModelType string

const (
	NewSubscriptionPriceModelNewSubscriptionCumulativeGroupedBulkPriceModelTypeCumulativeGroupedBulk NewSubscriptionPriceModelNewSubscriptionCumulativeGroupedBulkPriceModelType = "cumulative_grouped_bulk"
)

func (r NewSubscriptionPriceModelNewSubscriptionCumulativeGroupedBulkPriceModelType) IsKnown() bool {
	switch r {
	case NewSubscriptionPriceModelNewSubscriptionCumulativeGroupedBulkPriceModelTypeCumulativeGroupedBulk:
		return true
	}
	return false
}

type NewSubscriptionPriceModelNewSubscriptionMaxGroupTieredPackagePriceParam struct {
	// The cadence to bill for this price on.
	Cadence param.Field[NewSubscriptionPriceModelNewSubscriptionMaxGroupTieredPackagePriceCadence] `json:"cadence,required"`
	// The id of the item the price will be associated with.
	ItemID                      param.Field[string]                                                                      `json:"item_id,required"`
	MaxGroupTieredPackageConfig param.Field[CustomRatingFunctionConfigModelParam]                                        `json:"max_group_tiered_package_config,required"`
	ModelType                   param.Field[NewSubscriptionPriceModelNewSubscriptionMaxGroupTieredPackagePriceModelType] `json:"model_type,required"`
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
	BillingCycleConfiguration param.Field[NewBillingCycleConfigurationModelParam] `json:"billing_cycle_configuration"`
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
	InvoicingCycleConfiguration param.Field[NewBillingCycleConfigurationModelParam] `json:"invoicing_cycle_configuration"`
	// User-specified key/value pairs for the resource. Individual keys can be removed
	// by setting the value to `null`, and the entire metadata mapping can be cleared
	// by setting `metadata` to `null`.
	Metadata param.Field[map[string]string] `json:"metadata"`
	// A transient ID that can be used to reference this price when adding adjustments
	// in the same API call.
	ReferenceID param.Field[string] `json:"reference_id"`
}

func (r NewSubscriptionPriceModelNewSubscriptionMaxGroupTieredPackagePriceParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r NewSubscriptionPriceModelNewSubscriptionMaxGroupTieredPackagePriceParam) implementsNewSubscriptionPriceModelUnionParam() {
}

// The cadence to bill for this price on.
type NewSubscriptionPriceModelNewSubscriptionMaxGroupTieredPackagePriceCadence string

const (
	NewSubscriptionPriceModelNewSubscriptionMaxGroupTieredPackagePriceCadenceAnnual     NewSubscriptionPriceModelNewSubscriptionMaxGroupTieredPackagePriceCadence = "annual"
	NewSubscriptionPriceModelNewSubscriptionMaxGroupTieredPackagePriceCadenceSemiAnnual NewSubscriptionPriceModelNewSubscriptionMaxGroupTieredPackagePriceCadence = "semi_annual"
	NewSubscriptionPriceModelNewSubscriptionMaxGroupTieredPackagePriceCadenceMonthly    NewSubscriptionPriceModelNewSubscriptionMaxGroupTieredPackagePriceCadence = "monthly"
	NewSubscriptionPriceModelNewSubscriptionMaxGroupTieredPackagePriceCadenceQuarterly  NewSubscriptionPriceModelNewSubscriptionMaxGroupTieredPackagePriceCadence = "quarterly"
	NewSubscriptionPriceModelNewSubscriptionMaxGroupTieredPackagePriceCadenceOneTime    NewSubscriptionPriceModelNewSubscriptionMaxGroupTieredPackagePriceCadence = "one_time"
	NewSubscriptionPriceModelNewSubscriptionMaxGroupTieredPackagePriceCadenceCustom     NewSubscriptionPriceModelNewSubscriptionMaxGroupTieredPackagePriceCadence = "custom"
)

func (r NewSubscriptionPriceModelNewSubscriptionMaxGroupTieredPackagePriceCadence) IsKnown() bool {
	switch r {
	case NewSubscriptionPriceModelNewSubscriptionMaxGroupTieredPackagePriceCadenceAnnual, NewSubscriptionPriceModelNewSubscriptionMaxGroupTieredPackagePriceCadenceSemiAnnual, NewSubscriptionPriceModelNewSubscriptionMaxGroupTieredPackagePriceCadenceMonthly, NewSubscriptionPriceModelNewSubscriptionMaxGroupTieredPackagePriceCadenceQuarterly, NewSubscriptionPriceModelNewSubscriptionMaxGroupTieredPackagePriceCadenceOneTime, NewSubscriptionPriceModelNewSubscriptionMaxGroupTieredPackagePriceCadenceCustom:
		return true
	}
	return false
}

type NewSubscriptionPriceModelNewSubscriptionMaxGroupTieredPackagePriceModelType string

const (
	NewSubscriptionPriceModelNewSubscriptionMaxGroupTieredPackagePriceModelTypeMaxGroupTieredPackage NewSubscriptionPriceModelNewSubscriptionMaxGroupTieredPackagePriceModelType = "max_group_tiered_package"
)

func (r NewSubscriptionPriceModelNewSubscriptionMaxGroupTieredPackagePriceModelType) IsKnown() bool {
	switch r {
	case NewSubscriptionPriceModelNewSubscriptionMaxGroupTieredPackagePriceModelTypeMaxGroupTieredPackage:
		return true
	}
	return false
}

type NewSubscriptionPriceModelNewSubscriptionGroupedWithMeteredMinimumPriceParam struct {
	// The cadence to bill for this price on.
	Cadence                         param.Field[NewSubscriptionPriceModelNewSubscriptionGroupedWithMeteredMinimumPriceCadence] `json:"cadence,required"`
	GroupedWithMeteredMinimumConfig param.Field[CustomRatingFunctionConfigModelParam]                                          `json:"grouped_with_metered_minimum_config,required"`
	// The id of the item the price will be associated with.
	ItemID    param.Field[string]                                                                          `json:"item_id,required"`
	ModelType param.Field[NewSubscriptionPriceModelNewSubscriptionGroupedWithMeteredMinimumPriceModelType] `json:"model_type,required"`
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
	BillingCycleConfiguration param.Field[NewBillingCycleConfigurationModelParam] `json:"billing_cycle_configuration"`
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
	InvoicingCycleConfiguration param.Field[NewBillingCycleConfigurationModelParam] `json:"invoicing_cycle_configuration"`
	// User-specified key/value pairs for the resource. Individual keys can be removed
	// by setting the value to `null`, and the entire metadata mapping can be cleared
	// by setting `metadata` to `null`.
	Metadata param.Field[map[string]string] `json:"metadata"`
	// A transient ID that can be used to reference this price when adding adjustments
	// in the same API call.
	ReferenceID param.Field[string] `json:"reference_id"`
}

func (r NewSubscriptionPriceModelNewSubscriptionGroupedWithMeteredMinimumPriceParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r NewSubscriptionPriceModelNewSubscriptionGroupedWithMeteredMinimumPriceParam) implementsNewSubscriptionPriceModelUnionParam() {
}

// The cadence to bill for this price on.
type NewSubscriptionPriceModelNewSubscriptionGroupedWithMeteredMinimumPriceCadence string

const (
	NewSubscriptionPriceModelNewSubscriptionGroupedWithMeteredMinimumPriceCadenceAnnual     NewSubscriptionPriceModelNewSubscriptionGroupedWithMeteredMinimumPriceCadence = "annual"
	NewSubscriptionPriceModelNewSubscriptionGroupedWithMeteredMinimumPriceCadenceSemiAnnual NewSubscriptionPriceModelNewSubscriptionGroupedWithMeteredMinimumPriceCadence = "semi_annual"
	NewSubscriptionPriceModelNewSubscriptionGroupedWithMeteredMinimumPriceCadenceMonthly    NewSubscriptionPriceModelNewSubscriptionGroupedWithMeteredMinimumPriceCadence = "monthly"
	NewSubscriptionPriceModelNewSubscriptionGroupedWithMeteredMinimumPriceCadenceQuarterly  NewSubscriptionPriceModelNewSubscriptionGroupedWithMeteredMinimumPriceCadence = "quarterly"
	NewSubscriptionPriceModelNewSubscriptionGroupedWithMeteredMinimumPriceCadenceOneTime    NewSubscriptionPriceModelNewSubscriptionGroupedWithMeteredMinimumPriceCadence = "one_time"
	NewSubscriptionPriceModelNewSubscriptionGroupedWithMeteredMinimumPriceCadenceCustom     NewSubscriptionPriceModelNewSubscriptionGroupedWithMeteredMinimumPriceCadence = "custom"
)

func (r NewSubscriptionPriceModelNewSubscriptionGroupedWithMeteredMinimumPriceCadence) IsKnown() bool {
	switch r {
	case NewSubscriptionPriceModelNewSubscriptionGroupedWithMeteredMinimumPriceCadenceAnnual, NewSubscriptionPriceModelNewSubscriptionGroupedWithMeteredMinimumPriceCadenceSemiAnnual, NewSubscriptionPriceModelNewSubscriptionGroupedWithMeteredMinimumPriceCadenceMonthly, NewSubscriptionPriceModelNewSubscriptionGroupedWithMeteredMinimumPriceCadenceQuarterly, NewSubscriptionPriceModelNewSubscriptionGroupedWithMeteredMinimumPriceCadenceOneTime, NewSubscriptionPriceModelNewSubscriptionGroupedWithMeteredMinimumPriceCadenceCustom:
		return true
	}
	return false
}

type NewSubscriptionPriceModelNewSubscriptionGroupedWithMeteredMinimumPriceModelType string

const (
	NewSubscriptionPriceModelNewSubscriptionGroupedWithMeteredMinimumPriceModelTypeGroupedWithMeteredMinimum NewSubscriptionPriceModelNewSubscriptionGroupedWithMeteredMinimumPriceModelType = "grouped_with_metered_minimum"
)

func (r NewSubscriptionPriceModelNewSubscriptionGroupedWithMeteredMinimumPriceModelType) IsKnown() bool {
	switch r {
	case NewSubscriptionPriceModelNewSubscriptionGroupedWithMeteredMinimumPriceModelTypeGroupedWithMeteredMinimum:
		return true
	}
	return false
}

type NewSubscriptionPriceModelNewSubscriptionMatrixWithDisplayNamePriceParam struct {
	// The cadence to bill for this price on.
	Cadence param.Field[NewSubscriptionPriceModelNewSubscriptionMatrixWithDisplayNamePriceCadence] `json:"cadence,required"`
	// The id of the item the price will be associated with.
	ItemID                      param.Field[string]                                                                      `json:"item_id,required"`
	MatrixWithDisplayNameConfig param.Field[CustomRatingFunctionConfigModelParam]                                        `json:"matrix_with_display_name_config,required"`
	ModelType                   param.Field[NewSubscriptionPriceModelNewSubscriptionMatrixWithDisplayNamePriceModelType] `json:"model_type,required"`
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
	BillingCycleConfiguration param.Field[NewBillingCycleConfigurationModelParam] `json:"billing_cycle_configuration"`
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
	InvoicingCycleConfiguration param.Field[NewBillingCycleConfigurationModelParam] `json:"invoicing_cycle_configuration"`
	// User-specified key/value pairs for the resource. Individual keys can be removed
	// by setting the value to `null`, and the entire metadata mapping can be cleared
	// by setting `metadata` to `null`.
	Metadata param.Field[map[string]string] `json:"metadata"`
	// A transient ID that can be used to reference this price when adding adjustments
	// in the same API call.
	ReferenceID param.Field[string] `json:"reference_id"`
}

func (r NewSubscriptionPriceModelNewSubscriptionMatrixWithDisplayNamePriceParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r NewSubscriptionPriceModelNewSubscriptionMatrixWithDisplayNamePriceParam) implementsNewSubscriptionPriceModelUnionParam() {
}

// The cadence to bill for this price on.
type NewSubscriptionPriceModelNewSubscriptionMatrixWithDisplayNamePriceCadence string

const (
	NewSubscriptionPriceModelNewSubscriptionMatrixWithDisplayNamePriceCadenceAnnual     NewSubscriptionPriceModelNewSubscriptionMatrixWithDisplayNamePriceCadence = "annual"
	NewSubscriptionPriceModelNewSubscriptionMatrixWithDisplayNamePriceCadenceSemiAnnual NewSubscriptionPriceModelNewSubscriptionMatrixWithDisplayNamePriceCadence = "semi_annual"
	NewSubscriptionPriceModelNewSubscriptionMatrixWithDisplayNamePriceCadenceMonthly    NewSubscriptionPriceModelNewSubscriptionMatrixWithDisplayNamePriceCadence = "monthly"
	NewSubscriptionPriceModelNewSubscriptionMatrixWithDisplayNamePriceCadenceQuarterly  NewSubscriptionPriceModelNewSubscriptionMatrixWithDisplayNamePriceCadence = "quarterly"
	NewSubscriptionPriceModelNewSubscriptionMatrixWithDisplayNamePriceCadenceOneTime    NewSubscriptionPriceModelNewSubscriptionMatrixWithDisplayNamePriceCadence = "one_time"
	NewSubscriptionPriceModelNewSubscriptionMatrixWithDisplayNamePriceCadenceCustom     NewSubscriptionPriceModelNewSubscriptionMatrixWithDisplayNamePriceCadence = "custom"
)

func (r NewSubscriptionPriceModelNewSubscriptionMatrixWithDisplayNamePriceCadence) IsKnown() bool {
	switch r {
	case NewSubscriptionPriceModelNewSubscriptionMatrixWithDisplayNamePriceCadenceAnnual, NewSubscriptionPriceModelNewSubscriptionMatrixWithDisplayNamePriceCadenceSemiAnnual, NewSubscriptionPriceModelNewSubscriptionMatrixWithDisplayNamePriceCadenceMonthly, NewSubscriptionPriceModelNewSubscriptionMatrixWithDisplayNamePriceCadenceQuarterly, NewSubscriptionPriceModelNewSubscriptionMatrixWithDisplayNamePriceCadenceOneTime, NewSubscriptionPriceModelNewSubscriptionMatrixWithDisplayNamePriceCadenceCustom:
		return true
	}
	return false
}

type NewSubscriptionPriceModelNewSubscriptionMatrixWithDisplayNamePriceModelType string

const (
	NewSubscriptionPriceModelNewSubscriptionMatrixWithDisplayNamePriceModelTypeMatrixWithDisplayName NewSubscriptionPriceModelNewSubscriptionMatrixWithDisplayNamePriceModelType = "matrix_with_display_name"
)

func (r NewSubscriptionPriceModelNewSubscriptionMatrixWithDisplayNamePriceModelType) IsKnown() bool {
	switch r {
	case NewSubscriptionPriceModelNewSubscriptionMatrixWithDisplayNamePriceModelTypeMatrixWithDisplayName:
		return true
	}
	return false
}

type NewSubscriptionPriceModelNewSubscriptionGroupedTieredPackagePriceParam struct {
	// The cadence to bill for this price on.
	Cadence                    param.Field[NewSubscriptionPriceModelNewSubscriptionGroupedTieredPackagePriceCadence] `json:"cadence,required"`
	GroupedTieredPackageConfig param.Field[CustomRatingFunctionConfigModelParam]                                     `json:"grouped_tiered_package_config,required"`
	// The id of the item the price will be associated with.
	ItemID    param.Field[string]                                                                     `json:"item_id,required"`
	ModelType param.Field[NewSubscriptionPriceModelNewSubscriptionGroupedTieredPackagePriceModelType] `json:"model_type,required"`
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
	BillingCycleConfiguration param.Field[NewBillingCycleConfigurationModelParam] `json:"billing_cycle_configuration"`
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
	InvoicingCycleConfiguration param.Field[NewBillingCycleConfigurationModelParam] `json:"invoicing_cycle_configuration"`
	// User-specified key/value pairs for the resource. Individual keys can be removed
	// by setting the value to `null`, and the entire metadata mapping can be cleared
	// by setting `metadata` to `null`.
	Metadata param.Field[map[string]string] `json:"metadata"`
	// A transient ID that can be used to reference this price when adding adjustments
	// in the same API call.
	ReferenceID param.Field[string] `json:"reference_id"`
}

func (r NewSubscriptionPriceModelNewSubscriptionGroupedTieredPackagePriceParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r NewSubscriptionPriceModelNewSubscriptionGroupedTieredPackagePriceParam) implementsNewSubscriptionPriceModelUnionParam() {
}

// The cadence to bill for this price on.
type NewSubscriptionPriceModelNewSubscriptionGroupedTieredPackagePriceCadence string

const (
	NewSubscriptionPriceModelNewSubscriptionGroupedTieredPackagePriceCadenceAnnual     NewSubscriptionPriceModelNewSubscriptionGroupedTieredPackagePriceCadence = "annual"
	NewSubscriptionPriceModelNewSubscriptionGroupedTieredPackagePriceCadenceSemiAnnual NewSubscriptionPriceModelNewSubscriptionGroupedTieredPackagePriceCadence = "semi_annual"
	NewSubscriptionPriceModelNewSubscriptionGroupedTieredPackagePriceCadenceMonthly    NewSubscriptionPriceModelNewSubscriptionGroupedTieredPackagePriceCadence = "monthly"
	NewSubscriptionPriceModelNewSubscriptionGroupedTieredPackagePriceCadenceQuarterly  NewSubscriptionPriceModelNewSubscriptionGroupedTieredPackagePriceCadence = "quarterly"
	NewSubscriptionPriceModelNewSubscriptionGroupedTieredPackagePriceCadenceOneTime    NewSubscriptionPriceModelNewSubscriptionGroupedTieredPackagePriceCadence = "one_time"
	NewSubscriptionPriceModelNewSubscriptionGroupedTieredPackagePriceCadenceCustom     NewSubscriptionPriceModelNewSubscriptionGroupedTieredPackagePriceCadence = "custom"
)

func (r NewSubscriptionPriceModelNewSubscriptionGroupedTieredPackagePriceCadence) IsKnown() bool {
	switch r {
	case NewSubscriptionPriceModelNewSubscriptionGroupedTieredPackagePriceCadenceAnnual, NewSubscriptionPriceModelNewSubscriptionGroupedTieredPackagePriceCadenceSemiAnnual, NewSubscriptionPriceModelNewSubscriptionGroupedTieredPackagePriceCadenceMonthly, NewSubscriptionPriceModelNewSubscriptionGroupedTieredPackagePriceCadenceQuarterly, NewSubscriptionPriceModelNewSubscriptionGroupedTieredPackagePriceCadenceOneTime, NewSubscriptionPriceModelNewSubscriptionGroupedTieredPackagePriceCadenceCustom:
		return true
	}
	return false
}

type NewSubscriptionPriceModelNewSubscriptionGroupedTieredPackagePriceModelType string

const (
	NewSubscriptionPriceModelNewSubscriptionGroupedTieredPackagePriceModelTypeGroupedTieredPackage NewSubscriptionPriceModelNewSubscriptionGroupedTieredPackagePriceModelType = "grouped_tiered_package"
)

func (r NewSubscriptionPriceModelNewSubscriptionGroupedTieredPackagePriceModelType) IsKnown() bool {
	switch r {
	case NewSubscriptionPriceModelNewSubscriptionGroupedTieredPackagePriceModelTypeGroupedTieredPackage:
		return true
	}
	return false
}

// The cadence to bill for this price on.
type NewSubscriptionPriceModelCadence string

const (
	NewSubscriptionPriceModelCadenceAnnual     NewSubscriptionPriceModelCadence = "annual"
	NewSubscriptionPriceModelCadenceSemiAnnual NewSubscriptionPriceModelCadence = "semi_annual"
	NewSubscriptionPriceModelCadenceMonthly    NewSubscriptionPriceModelCadence = "monthly"
	NewSubscriptionPriceModelCadenceQuarterly  NewSubscriptionPriceModelCadence = "quarterly"
	NewSubscriptionPriceModelCadenceOneTime    NewSubscriptionPriceModelCadence = "one_time"
	NewSubscriptionPriceModelCadenceCustom     NewSubscriptionPriceModelCadence = "custom"
)

func (r NewSubscriptionPriceModelCadence) IsKnown() bool {
	switch r {
	case NewSubscriptionPriceModelCadenceAnnual, NewSubscriptionPriceModelCadenceSemiAnnual, NewSubscriptionPriceModelCadenceMonthly, NewSubscriptionPriceModelCadenceQuarterly, NewSubscriptionPriceModelCadenceOneTime, NewSubscriptionPriceModelCadenceCustom:
		return true
	}
	return false
}

type NewSubscriptionPriceModelModelType string

const (
	NewSubscriptionPriceModelModelTypeUnit                            NewSubscriptionPriceModelModelType = "unit"
	NewSubscriptionPriceModelModelTypePackage                         NewSubscriptionPriceModelModelType = "package"
	NewSubscriptionPriceModelModelTypeMatrix                          NewSubscriptionPriceModelModelType = "matrix"
	NewSubscriptionPriceModelModelTypeTiered                          NewSubscriptionPriceModelModelType = "tiered"
	NewSubscriptionPriceModelModelTypeTieredBps                       NewSubscriptionPriceModelModelType = "tiered_bps"
	NewSubscriptionPriceModelModelTypeBps                             NewSubscriptionPriceModelModelType = "bps"
	NewSubscriptionPriceModelModelTypeBulkBps                         NewSubscriptionPriceModelModelType = "bulk_bps"
	NewSubscriptionPriceModelModelTypeBulk                            NewSubscriptionPriceModelModelType = "bulk"
	NewSubscriptionPriceModelModelTypeThresholdTotalAmount            NewSubscriptionPriceModelModelType = "threshold_total_amount"
	NewSubscriptionPriceModelModelTypeTieredPackage                   NewSubscriptionPriceModelModelType = "tiered_package"
	NewSubscriptionPriceModelModelTypeTieredWithMinimum               NewSubscriptionPriceModelModelType = "tiered_with_minimum"
	NewSubscriptionPriceModelModelTypeUnitWithPercent                 NewSubscriptionPriceModelModelType = "unit_with_percent"
	NewSubscriptionPriceModelModelTypePackageWithAllocation           NewSubscriptionPriceModelModelType = "package_with_allocation"
	NewSubscriptionPriceModelModelTypeTieredWithProration             NewSubscriptionPriceModelModelType = "tiered_with_proration"
	NewSubscriptionPriceModelModelTypeUnitWithProration               NewSubscriptionPriceModelModelType = "unit_with_proration"
	NewSubscriptionPriceModelModelTypeGroupedAllocation               NewSubscriptionPriceModelModelType = "grouped_allocation"
	NewSubscriptionPriceModelModelTypeGroupedWithProratedMinimum      NewSubscriptionPriceModelModelType = "grouped_with_prorated_minimum"
	NewSubscriptionPriceModelModelTypeBulkWithProration               NewSubscriptionPriceModelModelType = "bulk_with_proration"
	NewSubscriptionPriceModelModelTypeScalableMatrixWithUnitPricing   NewSubscriptionPriceModelModelType = "scalable_matrix_with_unit_pricing"
	NewSubscriptionPriceModelModelTypeScalableMatrixWithTieredPricing NewSubscriptionPriceModelModelType = "scalable_matrix_with_tiered_pricing"
	NewSubscriptionPriceModelModelTypeCumulativeGroupedBulk           NewSubscriptionPriceModelModelType = "cumulative_grouped_bulk"
	NewSubscriptionPriceModelModelTypeMaxGroupTieredPackage           NewSubscriptionPriceModelModelType = "max_group_tiered_package"
	NewSubscriptionPriceModelModelTypeGroupedWithMeteredMinimum       NewSubscriptionPriceModelModelType = "grouped_with_metered_minimum"
	NewSubscriptionPriceModelModelTypeMatrixWithDisplayName           NewSubscriptionPriceModelModelType = "matrix_with_display_name"
	NewSubscriptionPriceModelModelTypeGroupedTieredPackage            NewSubscriptionPriceModelModelType = "grouped_tiered_package"
)

func (r NewSubscriptionPriceModelModelType) IsKnown() bool {
	switch r {
	case NewSubscriptionPriceModelModelTypeUnit, NewSubscriptionPriceModelModelTypePackage, NewSubscriptionPriceModelModelTypeMatrix, NewSubscriptionPriceModelModelTypeTiered, NewSubscriptionPriceModelModelTypeTieredBps, NewSubscriptionPriceModelModelTypeBps, NewSubscriptionPriceModelModelTypeBulkBps, NewSubscriptionPriceModelModelTypeBulk, NewSubscriptionPriceModelModelTypeThresholdTotalAmount, NewSubscriptionPriceModelModelTypeTieredPackage, NewSubscriptionPriceModelModelTypeTieredWithMinimum, NewSubscriptionPriceModelModelTypeUnitWithPercent, NewSubscriptionPriceModelModelTypePackageWithAllocation, NewSubscriptionPriceModelModelTypeTieredWithProration, NewSubscriptionPriceModelModelTypeUnitWithProration, NewSubscriptionPriceModelModelTypeGroupedAllocation, NewSubscriptionPriceModelModelTypeGroupedWithProratedMinimum, NewSubscriptionPriceModelModelTypeBulkWithProration, NewSubscriptionPriceModelModelTypeScalableMatrixWithUnitPricing, NewSubscriptionPriceModelModelTypeScalableMatrixWithTieredPricing, NewSubscriptionPriceModelModelTypeCumulativeGroupedBulk, NewSubscriptionPriceModelModelTypeMaxGroupTieredPackage, NewSubscriptionPriceModelModelTypeGroupedWithMeteredMinimum, NewSubscriptionPriceModelModelTypeMatrixWithDisplayName, NewSubscriptionPriceModelModelTypeGroupedTieredPackage:
		return true
	}
	return false
}

type NewTaxConfigurationModelParam struct {
	TaxExempt        param.Field[bool]                                `json:"tax_exempt,required"`
	TaxProvider      param.Field[NewTaxConfigurationModelTaxProvider] `json:"tax_provider,required"`
	TaxExemptionCode param.Field[string]                              `json:"tax_exemption_code"`
}

func (r NewTaxConfigurationModelParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r NewTaxConfigurationModelParam) implementsNewTaxConfigurationModelUnionParam() {}

// Satisfied by [shared.NewTaxConfigurationModelNewAvalaraTaxConfigurationParam],
// [shared.NewTaxConfigurationModelNewTaxJarConfigurationParam],
// [NewTaxConfigurationModelParam].
type NewTaxConfigurationModelUnionParam interface {
	implementsNewTaxConfigurationModelUnionParam()
}

type NewTaxConfigurationModelNewAvalaraTaxConfigurationParam struct {
	TaxExempt        param.Field[bool]                                                          `json:"tax_exempt,required"`
	TaxProvider      param.Field[NewTaxConfigurationModelNewAvalaraTaxConfigurationTaxProvider] `json:"tax_provider,required"`
	TaxExemptionCode param.Field[string]                                                        `json:"tax_exemption_code"`
}

func (r NewTaxConfigurationModelNewAvalaraTaxConfigurationParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r NewTaxConfigurationModelNewAvalaraTaxConfigurationParam) implementsNewTaxConfigurationModelUnionParam() {
}

type NewTaxConfigurationModelNewAvalaraTaxConfigurationTaxProvider string

const (
	NewTaxConfigurationModelNewAvalaraTaxConfigurationTaxProviderAvalara NewTaxConfigurationModelNewAvalaraTaxConfigurationTaxProvider = "avalara"
)

func (r NewTaxConfigurationModelNewAvalaraTaxConfigurationTaxProvider) IsKnown() bool {
	switch r {
	case NewTaxConfigurationModelNewAvalaraTaxConfigurationTaxProviderAvalara:
		return true
	}
	return false
}

type NewTaxConfigurationModelNewTaxJarConfigurationParam struct {
	TaxExempt   param.Field[bool]                                                      `json:"tax_exempt,required"`
	TaxProvider param.Field[NewTaxConfigurationModelNewTaxJarConfigurationTaxProvider] `json:"tax_provider,required"`
}

func (r NewTaxConfigurationModelNewTaxJarConfigurationParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r NewTaxConfigurationModelNewTaxJarConfigurationParam) implementsNewTaxConfigurationModelUnionParam() {
}

type NewTaxConfigurationModelNewTaxJarConfigurationTaxProvider string

const (
	NewTaxConfigurationModelNewTaxJarConfigurationTaxProviderTaxjar NewTaxConfigurationModelNewTaxJarConfigurationTaxProvider = "taxjar"
)

func (r NewTaxConfigurationModelNewTaxJarConfigurationTaxProvider) IsKnown() bool {
	switch r {
	case NewTaxConfigurationModelNewTaxJarConfigurationTaxProviderTaxjar:
		return true
	}
	return false
}

type NewTaxConfigurationModelTaxProvider string

const (
	NewTaxConfigurationModelTaxProviderAvalara NewTaxConfigurationModelTaxProvider = "avalara"
	NewTaxConfigurationModelTaxProviderTaxjar  NewTaxConfigurationModelTaxProvider = "taxjar"
)

func (r NewTaxConfigurationModelTaxProvider) IsKnown() bool {
	switch r {
	case NewTaxConfigurationModelTaxProviderAvalara, NewTaxConfigurationModelTaxProviderTaxjar:
		return true
	}
	return false
}

type PackageConfigModel struct {
	// A currency amount to rate usage by
	PackageAmount string `json:"package_amount,required"`
	// An integer amount to represent package size. For example, 1000 here would divide
	// usage by 1000 before multiplying by package_amount in rating
	PackageSize int64                  `json:"package_size,required"`
	JSON        packageConfigModelJSON `json:"-"`
}

// packageConfigModelJSON contains the JSON metadata for the struct
// [PackageConfigModel]
type packageConfigModelJSON struct {
	PackageAmount apijson.Field
	PackageSize   apijson.Field
	raw           string
	ExtraFields   map[string]apijson.Field
}

func (r *PackageConfigModel) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r packageConfigModelJSON) RawJSON() string {
	return r.raw
}

type PackageConfigModelParam struct {
	// A currency amount to rate usage by
	PackageAmount param.Field[string] `json:"package_amount,required"`
	// An integer amount to represent package size. For example, 1000 here would divide
	// usage by 1000 before multiplying by package_amount in rating
	PackageSize param.Field[int64] `json:"package_size,required"`
}

func (r PackageConfigModelParam) MarshalJSON() (data []byte, err error) {
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

type PaymentAttemptModel struct {
	// The ID of the payment attempt.
	ID string `json:"id,required"`
	// The amount of the payment attempt.
	Amount string `json:"amount,required"`
	// The time at which the payment attempt was created.
	CreatedAt time.Time `json:"created_at,required" format:"date-time"`
	// The payment provider that attempted to collect the payment.
	PaymentProvider PaymentAttemptModelPaymentProvider `json:"payment_provider,required,nullable"`
	// The ID of the payment attempt in the payment provider.
	PaymentProviderID string `json:"payment_provider_id,required,nullable"`
	// Whether the payment attempt succeeded.
	Succeeded bool                    `json:"succeeded,required"`
	JSON      paymentAttemptModelJSON `json:"-"`
}

// paymentAttemptModelJSON contains the JSON metadata for the struct
// [PaymentAttemptModel]
type paymentAttemptModelJSON struct {
	ID                apijson.Field
	Amount            apijson.Field
	CreatedAt         apijson.Field
	PaymentProvider   apijson.Field
	PaymentProviderID apijson.Field
	Succeeded         apijson.Field
	raw               string
	ExtraFields       map[string]apijson.Field
}

func (r *PaymentAttemptModel) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r paymentAttemptModelJSON) RawJSON() string {
	return r.raw
}

// The payment provider that attempted to collect the payment.
type PaymentAttemptModelPaymentProvider string

const (
	PaymentAttemptModelPaymentProviderStripe PaymentAttemptModelPaymentProvider = "stripe"
)

func (r PaymentAttemptModelPaymentProvider) IsKnown() bool {
	switch r {
	case PaymentAttemptModelPaymentProviderStripe:
		return true
	}
	return false
}

type PercentageDiscount struct {
	// List of price_ids that this discount applies to. For plan/plan phase discounts,
	// this can be a subset of prices.
	AppliesToPriceIDs []string                       `json:"applies_to_price_ids,required"`
	DiscountType      PercentageDiscountDiscountType `json:"discount_type,required"`
	// Only available if discount_type is `percentage`. This is a number between 0
	// and 1.
	PercentageDiscount float64                `json:"percentage_discount,required"`
	Reason             string                 `json:"reason,nullable"`
	JSON               percentageDiscountJSON `json:"-"`
}

// percentageDiscountJSON contains the JSON metadata for the struct
// [PercentageDiscount]
type percentageDiscountJSON struct {
	AppliesToPriceIDs  apijson.Field
	DiscountType       apijson.Field
	PercentageDiscount apijson.Field
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

func (r PercentageDiscount) ImplementsCouponModelDiscount() {}

func (r PercentageDiscount) ImplementsDiscount() {}

func (r PercentageDiscount) ImplementsInvoiceLevelDiscount() {}

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
	// List of price_ids that this discount applies to. For plan/plan phase discounts,
	// this can be a subset of prices.
	AppliesToPriceIDs param.Field[[]string]                       `json:"applies_to_price_ids,required"`
	DiscountType      param.Field[PercentageDiscountDiscountType] `json:"discount_type,required"`
	// Only available if discount_type is `percentage`. This is a number between 0
	// and 1.
	PercentageDiscount param.Field[float64] `json:"percentage_discount,required"`
	Reason             param.Field[string]  `json:"reason"`
}

func (r PercentageDiscountParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r PercentageDiscountParam) ImplementsDiscountUnionParam() {}

type PercentageDiscountIntervalModel struct {
	// The price ids that this discount interval applies to.
	AppliesToPriceIDs []string `json:"applies_to_price_ids,required"`
	// The price interval ids that this discount interval applies to.
	AppliesToPriceIntervalIDs []string                                    `json:"applies_to_price_interval_ids,required"`
	DiscountType              PercentageDiscountIntervalModelDiscountType `json:"discount_type,required"`
	// The end date of the discount interval.
	EndDate time.Time `json:"end_date,required,nullable" format:"date-time"`
	// Only available if discount_type is `percentage`.This is a number between 0
	// and 1.
	PercentageDiscount float64 `json:"percentage_discount,required"`
	// The start date of the discount interval.
	StartDate time.Time                           `json:"start_date,required" format:"date-time"`
	JSON      percentageDiscountIntervalModelJSON `json:"-"`
}

// percentageDiscountIntervalModelJSON contains the JSON metadata for the struct
// [PercentageDiscountIntervalModel]
type percentageDiscountIntervalModelJSON struct {
	AppliesToPriceIDs         apijson.Field
	AppliesToPriceIntervalIDs apijson.Field
	DiscountType              apijson.Field
	EndDate                   apijson.Field
	PercentageDiscount        apijson.Field
	StartDate                 apijson.Field
	raw                       string
	ExtraFields               map[string]apijson.Field
}

func (r *PercentageDiscountIntervalModel) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r percentageDiscountIntervalModelJSON) RawJSON() string {
	return r.raw
}

func (r PercentageDiscountIntervalModel) ImplementsMutatedSubscriptionModelDiscountInterval() {}

func (r PercentageDiscountIntervalModel) ImplementsSubscriptionModelDiscountInterval() {}

type PercentageDiscountIntervalModelDiscountType string

const (
	PercentageDiscountIntervalModelDiscountTypePercentage PercentageDiscountIntervalModelDiscountType = "percentage"
)

func (r PercentageDiscountIntervalModelDiscountType) IsKnown() bool {
	switch r {
	case PercentageDiscountIntervalModelDiscountTypePercentage:
		return true
	}
	return false
}

type PlanMinifiedModel struct {
	ID string `json:"id,required,nullable"`
	// An optional user-defined ID for this plan resource, used throughout the system
	// as an alias for this Plan. Use this field to identify a plan by an existing
	// identifier in your system.
	ExternalPlanID string                `json:"external_plan_id,required,nullable"`
	Name           string                `json:"name,required,nullable"`
	JSON           planMinifiedModelJSON `json:"-"`
}

// planMinifiedModelJSON contains the JSON metadata for the struct
// [PlanMinifiedModel]
type planMinifiedModelJSON struct {
	ID             apijson.Field
	ExternalPlanID apijson.Field
	Name           apijson.Field
	raw            string
	ExtraFields    map[string]apijson.Field
}

func (r *PlanMinifiedModel) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r planMinifiedModelJSON) RawJSON() string {
	return r.raw
}

// The [Plan](/core-concepts#plan-and-price) resource represents a plan that can be
// subscribed to by a customer. Plans define the billing behavior of the
// subscription. You can see more about how to configure prices in the
// [Price resource](/reference/price).
type PlanModel struct {
	ID string `json:"id,required"`
	// Adjustments for this plan. If the plan has phases, this includes adjustments
	// across all phases of the plan.
	Adjustments []AdjustmentModel `json:"adjustments,required"`
	BasePlan    PlanMinifiedModel `json:"base_plan,required,nullable"`
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
	DefaultInvoiceMemo string   `json:"default_invoice_memo,required,nullable"`
	Description        string   `json:"description,required"`
	Discount           Discount `json:"discount,required,nullable"`
	// An optional user-defined ID for this plan resource, used throughout the system
	// as an alias for this Plan. Use this field to identify a plan by an existing
	// identifier in your system.
	ExternalPlanID string `json:"external_plan_id,required,nullable"`
	// An ISO 4217 currency string for which this plan is billed in. Matches `currency`
	// unless `currency` is a custom pricing unit.
	InvoicingCurrency string       `json:"invoicing_currency,required"`
	Maximum           MaximumModel `json:"maximum,required,nullable"`
	MaximumAmount     string       `json:"maximum_amount,required,nullable"`
	// User specified key-value pairs for the resource. If not present, this defaults
	// to an empty dictionary. Individual keys can be removed by setting the value to
	// `null`, and the entire metadata mapping can be cleared by setting `metadata` to
	// `null`.
	Metadata      map[string]string `json:"metadata,required"`
	Minimum       MinimumModel      `json:"minimum,required,nullable"`
	MinimumAmount string            `json:"minimum_amount,required,nullable"`
	Name          string            `json:"name,required"`
	// Determines the difference between the invoice issue date and the due date. A
	// value of "0" here signifies that invoices are due on issue, whereas a value of
	// "30" means that the customer has a month to pay the invoice before its overdue.
	// Note that individual subscriptions or invoices may set a different net terms
	// configuration.
	NetTerms   int64                `json:"net_terms,required,nullable"`
	PlanPhases []PlanModelPlanPhase `json:"plan_phases,required,nullable"`
	// Prices for this plan. If the plan has phases, this includes prices across all
	// phases of the plan.
	Prices      []PriceModel         `json:"prices,required"`
	Product     PlanModelProduct     `json:"product,required"`
	Status      PlanModelStatus      `json:"status,required"`
	TrialConfig PlanModelTrialConfig `json:"trial_config,required"`
	Version     int64                `json:"version,required"`
	JSON        planModelJSON        `json:"-"`
}

// planModelJSON contains the JSON metadata for the struct [PlanModel]
type planModelJSON struct {
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

func (r *PlanModel) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r planModelJSON) RawJSON() string {
	return r.raw
}

type PlanModelPlanPhase struct {
	ID          string   `json:"id,required"`
	Description string   `json:"description,required,nullable"`
	Discount    Discount `json:"discount,required,nullable"`
	// How many terms of length `duration_unit` this phase is active for. If null, this
	// phase is evergreen and active indefinitely
	Duration      int64                           `json:"duration,required,nullable"`
	DurationUnit  PlanModelPlanPhasesDurationUnit `json:"duration_unit,required,nullable"`
	Maximum       MaximumModel                    `json:"maximum,required,nullable"`
	MaximumAmount string                          `json:"maximum_amount,required,nullable"`
	Minimum       MinimumModel                    `json:"minimum,required,nullable"`
	MinimumAmount string                          `json:"minimum_amount,required,nullable"`
	Name          string                          `json:"name,required"`
	// Determines the ordering of the phase in a plan's lifecycle. 1 = first phase.
	Order int64                  `json:"order,required"`
	JSON  planModelPlanPhaseJSON `json:"-"`
}

// planModelPlanPhaseJSON contains the JSON metadata for the struct
// [PlanModelPlanPhase]
type planModelPlanPhaseJSON struct {
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

func (r *PlanModelPlanPhase) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r planModelPlanPhaseJSON) RawJSON() string {
	return r.raw
}

type PlanModelPlanPhasesDurationUnit string

const (
	PlanModelPlanPhasesDurationUnitDaily      PlanModelPlanPhasesDurationUnit = "daily"
	PlanModelPlanPhasesDurationUnitMonthly    PlanModelPlanPhasesDurationUnit = "monthly"
	PlanModelPlanPhasesDurationUnitQuarterly  PlanModelPlanPhasesDurationUnit = "quarterly"
	PlanModelPlanPhasesDurationUnitSemiAnnual PlanModelPlanPhasesDurationUnit = "semi_annual"
	PlanModelPlanPhasesDurationUnitAnnual     PlanModelPlanPhasesDurationUnit = "annual"
)

func (r PlanModelPlanPhasesDurationUnit) IsKnown() bool {
	switch r {
	case PlanModelPlanPhasesDurationUnitDaily, PlanModelPlanPhasesDurationUnitMonthly, PlanModelPlanPhasesDurationUnitQuarterly, PlanModelPlanPhasesDurationUnitSemiAnnual, PlanModelPlanPhasesDurationUnitAnnual:
		return true
	}
	return false
}

type PlanModelProduct struct {
	ID        string               `json:"id,required"`
	CreatedAt time.Time            `json:"created_at,required" format:"date-time"`
	Name      string               `json:"name,required"`
	JSON      planModelProductJSON `json:"-"`
}

// planModelProductJSON contains the JSON metadata for the struct
// [PlanModelProduct]
type planModelProductJSON struct {
	ID          apijson.Field
	CreatedAt   apijson.Field
	Name        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *PlanModelProduct) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r planModelProductJSON) RawJSON() string {
	return r.raw
}

type PlanModelStatus string

const (
	PlanModelStatusActive   PlanModelStatus = "active"
	PlanModelStatusArchived PlanModelStatus = "archived"
	PlanModelStatusDraft    PlanModelStatus = "draft"
)

func (r PlanModelStatus) IsKnown() bool {
	switch r {
	case PlanModelStatusActive, PlanModelStatusArchived, PlanModelStatusDraft:
		return true
	}
	return false
}

type PlanModelTrialConfig struct {
	TrialPeriod     int64                               `json:"trial_period,required,nullable"`
	TrialPeriodUnit PlanModelTrialConfigTrialPeriodUnit `json:"trial_period_unit,required"`
	JSON            planModelTrialConfigJSON            `json:"-"`
}

// planModelTrialConfigJSON contains the JSON metadata for the struct
// [PlanModelTrialConfig]
type planModelTrialConfigJSON struct {
	TrialPeriod     apijson.Field
	TrialPeriodUnit apijson.Field
	raw             string
	ExtraFields     map[string]apijson.Field
}

func (r *PlanModelTrialConfig) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r planModelTrialConfigJSON) RawJSON() string {
	return r.raw
}

type PlanModelTrialConfigTrialPeriodUnit string

const (
	PlanModelTrialConfigTrialPeriodUnitDays PlanModelTrialConfigTrialPeriodUnit = "days"
)

func (r PlanModelTrialConfigTrialPeriodUnit) IsKnown() bool {
	switch r {
	case PlanModelTrialConfigTrialPeriodUnitDays:
		return true
	}
	return false
}

type PriceIntervalFixedFeeQuantityTransitionModelParam struct {
	// The date that the fixed fee quantity transition should take effect.
	EffectiveDate param.Field[time.Time] `json:"effective_date,required" format:"date-time"`
	// The quantity of the fixed fee quantity transition.
	Quantity param.Field[int64] `json:"quantity,required"`
}

func (r PriceIntervalFixedFeeQuantityTransitionModelParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// The Price Interval resource represents a period of time for which a price will
// bill on a subscription. A subscription’s price intervals define its billing
// behavior.
type PriceIntervalModel struct {
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
	FixedFeeQuantityTransitions []PriceIntervalModelFixedFeeQuantityTransition `json:"fixed_fee_quantity_transitions,required,nullable"`
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
	Price PriceModel `json:"price,required"`
	// The start date of the price interval. This is the date that Orb starts billing
	// for this price.
	StartDate time.Time `json:"start_date,required" format:"date-time"`
	// A list of customer IDs whose usage events will be aggregated and billed under
	// this price interval.
	UsageCustomerIDs []string               `json:"usage_customer_ids,required,nullable"`
	JSON             priceIntervalModelJSON `json:"-"`
}

// priceIntervalModelJSON contains the JSON metadata for the struct
// [PriceIntervalModel]
type priceIntervalModelJSON struct {
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

func (r *PriceIntervalModel) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r priceIntervalModelJSON) RawJSON() string {
	return r.raw
}

type PriceIntervalModelFixedFeeQuantityTransition struct {
	EffectiveDate time.Time                                        `json:"effective_date,required" format:"date-time"`
	PriceID       string                                           `json:"price_id,required"`
	Quantity      int64                                            `json:"quantity,required"`
	JSON          priceIntervalModelFixedFeeQuantityTransitionJSON `json:"-"`
}

// priceIntervalModelFixedFeeQuantityTransitionJSON contains the JSON metadata for
// the struct [PriceIntervalModelFixedFeeQuantityTransition]
type priceIntervalModelFixedFeeQuantityTransitionJSON struct {
	EffectiveDate apijson.Field
	PriceID       apijson.Field
	Quantity      apijson.Field
	raw           string
	ExtraFields   map[string]apijson.Field
}

func (r *PriceIntervalModelFixedFeeQuantityTransition) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r priceIntervalModelFixedFeeQuantityTransitionJSON) RawJSON() string {
	return r.raw
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
type PriceModel struct {
	ID                          string                         `json:"id,required"`
	BillableMetric              BillableMetricTinyModel        `json:"billable_metric,required,nullable"`
	BillingCycleConfiguration   BillingCycleConfigurationModel `json:"billing_cycle_configuration,required"`
	Cadence                     PriceModelCadence              `json:"cadence,required"`
	ConversionRate              float64                        `json:"conversion_rate,required,nullable"`
	CreatedAt                   time.Time                      `json:"created_at,required" format:"date-time"`
	CreditAllocation            AllocationModel                `json:"credit_allocation,required,nullable"`
	Currency                    string                         `json:"currency,required"`
	Discount                    Discount                       `json:"discount,required,nullable"`
	ExternalPriceID             string                         `json:"external_price_id,required,nullable"`
	FixedPriceQuantity          float64                        `json:"fixed_price_quantity,required,nullable"`
	InvoicingCycleConfiguration BillingCycleConfigurationModel `json:"invoicing_cycle_configuration,required,nullable"`
	Item                        ItemSlimModel                  `json:"item,required"`
	Maximum                     MaximumModel                   `json:"maximum,required,nullable"`
	MaximumAmount               string                         `json:"maximum_amount,required,nullable"`
	// This field can have the runtime type of [map[string]string].
	Metadata                              interface{}                        `json:"metadata,required"`
	Minimum                               MinimumModel                       `json:"minimum,required,nullable"`
	MinimumAmount                         string                             `json:"minimum_amount,required,nullable"`
	ModelType                             PriceModelModelType                `json:"model_type,required"`
	Name                                  string                             `json:"name,required"`
	PlanPhaseOrder                        int64                              `json:"plan_phase_order,required,nullable"`
	PriceType                             PriceModelPriceType                `json:"price_type,required"`
	BpsConfig                             BpsConfigModel                     `json:"bps_config"`
	BulkBpsConfig                         BulkBpsConfigModel                 `json:"bulk_bps_config"`
	BulkConfig                            BulkConfigModel                    `json:"bulk_config"`
	BulkWithProrationConfig               CustomRatingFunctionConfigModel    `json:"bulk_with_proration_config"`
	CumulativeGroupedBulkConfig           CustomRatingFunctionConfigModel    `json:"cumulative_grouped_bulk_config"`
	DimensionalPriceConfiguration         DimensionalPriceConfigurationModel `json:"dimensional_price_configuration,nullable"`
	GroupedAllocationConfig               CustomRatingFunctionConfigModel    `json:"grouped_allocation_config"`
	GroupedTieredConfig                   CustomRatingFunctionConfigModel    `json:"grouped_tiered_config"`
	GroupedTieredPackageConfig            CustomRatingFunctionConfigModel    `json:"grouped_tiered_package_config"`
	GroupedWithMeteredMinimumConfig       CustomRatingFunctionConfigModel    `json:"grouped_with_metered_minimum_config"`
	GroupedWithProratedMinimumConfig      CustomRatingFunctionConfigModel    `json:"grouped_with_prorated_minimum_config"`
	MatrixConfig                          MatrixConfigModel                  `json:"matrix_config"`
	MatrixWithAllocationConfig            MatrixWithAllocationConfigModel    `json:"matrix_with_allocation_config"`
	MatrixWithDisplayNameConfig           CustomRatingFunctionConfigModel    `json:"matrix_with_display_name_config"`
	MaxGroupTieredPackageConfig           CustomRatingFunctionConfigModel    `json:"max_group_tiered_package_config"`
	PackageConfig                         PackageConfigModel                 `json:"package_config"`
	PackageWithAllocationConfig           CustomRatingFunctionConfigModel    `json:"package_with_allocation_config"`
	ScalableMatrixWithTieredPricingConfig CustomRatingFunctionConfigModel    `json:"scalable_matrix_with_tiered_pricing_config"`
	ScalableMatrixWithUnitPricingConfig   CustomRatingFunctionConfigModel    `json:"scalable_matrix_with_unit_pricing_config"`
	ThresholdTotalAmountConfig            CustomRatingFunctionConfigModel    `json:"threshold_total_amount_config"`
	TieredBpsConfig                       TieredBpsConfigModel               `json:"tiered_bps_config"`
	TieredConfig                          TieredConfigModel                  `json:"tiered_config"`
	TieredPackageConfig                   CustomRatingFunctionConfigModel    `json:"tiered_package_config"`
	TieredPackageWithMinimumConfig        CustomRatingFunctionConfigModel    `json:"tiered_package_with_minimum_config"`
	TieredWithMinimumConfig               CustomRatingFunctionConfigModel    `json:"tiered_with_minimum_config"`
	TieredWithProrationConfig             CustomRatingFunctionConfigModel    `json:"tiered_with_proration_config"`
	UnitConfig                            UnitConfigModel                    `json:"unit_config"`
	UnitWithPercentConfig                 CustomRatingFunctionConfigModel    `json:"unit_with_percent_config"`
	UnitWithProrationConfig               CustomRatingFunctionConfigModel    `json:"unit_with_proration_config"`
	JSON                                  priceModelJSON                     `json:"-"`
	union                                 PriceModelUnion
}

// priceModelJSON contains the JSON metadata for the struct [PriceModel]
type priceModelJSON struct {
	ID                                    apijson.Field
	BillableMetric                        apijson.Field
	BillingCycleConfiguration             apijson.Field
	Cadence                               apijson.Field
	ConversionRate                        apijson.Field
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
	BpsConfig                             apijson.Field
	BulkBpsConfig                         apijson.Field
	BulkConfig                            apijson.Field
	BulkWithProrationConfig               apijson.Field
	CumulativeGroupedBulkConfig           apijson.Field
	DimensionalPriceConfiguration         apijson.Field
	GroupedAllocationConfig               apijson.Field
	GroupedTieredConfig                   apijson.Field
	GroupedTieredPackageConfig            apijson.Field
	GroupedWithMeteredMinimumConfig       apijson.Field
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
	TieredBpsConfig                       apijson.Field
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

func (r priceModelJSON) RawJSON() string {
	return r.raw
}

func (r *PriceModel) UnmarshalJSON(data []byte) (err error) {
	*r = PriceModel{}
	err = apijson.UnmarshalRoot(data, &r.union)
	if err != nil {
		return err
	}
	return apijson.Port(r.union, &r)
}

// AsUnion returns a [PriceModelUnion] interface which you can cast to the specific
// types for more type safety.
//
// Possible runtime types of the union are [shared.PriceModelUnitPrice],
// [shared.PriceModelPackagePrice], [shared.PriceModelMatrixPrice],
// [shared.PriceModelTieredPrice], [shared.PriceModelTieredBpsPrice],
// [shared.PriceModelBpsPrice], [shared.PriceModelBulkBpsPrice],
// [shared.PriceModelBulkPrice], [shared.PriceModelThresholdTotalAmountPrice],
// [shared.PriceModelTieredPackagePrice], [shared.PriceModelGroupedTieredPrice],
// [shared.PriceModelTieredWithMinimumPrice],
// [shared.PriceModelTieredPackageWithMinimumPrice],
// [shared.PriceModelPackageWithAllocationPrice],
// [shared.PriceModelUnitWithPercentPrice],
// [shared.PriceModelMatrixWithAllocationPrice],
// [shared.PriceModelTieredWithProrationPrice],
// [shared.PriceModelUnitWithProrationPrice],
// [shared.PriceModelGroupedAllocationPrice],
// [shared.PriceModelGroupedWithProratedMinimumPrice],
// [shared.PriceModelGroupedWithMeteredMinimumPrice],
// [shared.PriceModelMatrixWithDisplayNamePrice],
// [shared.PriceModelBulkWithProrationPrice],
// [shared.PriceModelGroupedTieredPackagePrice],
// [shared.PriceModelMaxGroupTieredPackagePrice],
// [shared.PriceModelScalableMatrixWithUnitPricingPrice],
// [shared.PriceModelScalableMatrixWithTieredPricingPrice],
// [shared.PriceModelCumulativeGroupedBulkPrice].
func (r PriceModel) AsUnion() PriceModelUnion {
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
// Union satisfied by [shared.PriceModelUnitPrice],
// [shared.PriceModelPackagePrice], [shared.PriceModelMatrixPrice],
// [shared.PriceModelTieredPrice], [shared.PriceModelTieredBpsPrice],
// [shared.PriceModelBpsPrice], [shared.PriceModelBulkBpsPrice],
// [shared.PriceModelBulkPrice], [shared.PriceModelThresholdTotalAmountPrice],
// [shared.PriceModelTieredPackagePrice], [shared.PriceModelGroupedTieredPrice],
// [shared.PriceModelTieredWithMinimumPrice],
// [shared.PriceModelTieredPackageWithMinimumPrice],
// [shared.PriceModelPackageWithAllocationPrice],
// [shared.PriceModelUnitWithPercentPrice],
// [shared.PriceModelMatrixWithAllocationPrice],
// [shared.PriceModelTieredWithProrationPrice],
// [shared.PriceModelUnitWithProrationPrice],
// [shared.PriceModelGroupedAllocationPrice],
// [shared.PriceModelGroupedWithProratedMinimumPrice],
// [shared.PriceModelGroupedWithMeteredMinimumPrice],
// [shared.PriceModelMatrixWithDisplayNamePrice],
// [shared.PriceModelBulkWithProrationPrice],
// [shared.PriceModelGroupedTieredPackagePrice],
// [shared.PriceModelMaxGroupTieredPackagePrice],
// [shared.PriceModelScalableMatrixWithUnitPricingPrice],
// [shared.PriceModelScalableMatrixWithTieredPricingPrice] or
// [shared.PriceModelCumulativeGroupedBulkPrice].
type PriceModelUnion interface {
	implementsPriceModel()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*PriceModelUnion)(nil)).Elem(),
		"model_type",
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(PriceModelUnitPrice{}),
			DiscriminatorValue: "unit",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(PriceModelPackagePrice{}),
			DiscriminatorValue: "package",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(PriceModelMatrixPrice{}),
			DiscriminatorValue: "matrix",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(PriceModelTieredPrice{}),
			DiscriminatorValue: "tiered",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(PriceModelTieredBpsPrice{}),
			DiscriminatorValue: "tiered_bps",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(PriceModelBpsPrice{}),
			DiscriminatorValue: "bps",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(PriceModelBulkBpsPrice{}),
			DiscriminatorValue: "bulk_bps",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(PriceModelBulkPrice{}),
			DiscriminatorValue: "bulk",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(PriceModelThresholdTotalAmountPrice{}),
			DiscriminatorValue: "threshold_total_amount",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(PriceModelTieredPackagePrice{}),
			DiscriminatorValue: "tiered_package",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(PriceModelGroupedTieredPrice{}),
			DiscriminatorValue: "grouped_tiered",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(PriceModelTieredWithMinimumPrice{}),
			DiscriminatorValue: "tiered_with_minimum",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(PriceModelTieredPackageWithMinimumPrice{}),
			DiscriminatorValue: "tiered_package_with_minimum",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(PriceModelPackageWithAllocationPrice{}),
			DiscriminatorValue: "package_with_allocation",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(PriceModelUnitWithPercentPrice{}),
			DiscriminatorValue: "unit_with_percent",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(PriceModelMatrixWithAllocationPrice{}),
			DiscriminatorValue: "matrix_with_allocation",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(PriceModelTieredWithProrationPrice{}),
			DiscriminatorValue: "tiered_with_proration",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(PriceModelUnitWithProrationPrice{}),
			DiscriminatorValue: "unit_with_proration",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(PriceModelGroupedAllocationPrice{}),
			DiscriminatorValue: "grouped_allocation",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(PriceModelGroupedWithProratedMinimumPrice{}),
			DiscriminatorValue: "grouped_with_prorated_minimum",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(PriceModelGroupedWithMeteredMinimumPrice{}),
			DiscriminatorValue: "grouped_with_metered_minimum",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(PriceModelMatrixWithDisplayNamePrice{}),
			DiscriminatorValue: "matrix_with_display_name",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(PriceModelBulkWithProrationPrice{}),
			DiscriminatorValue: "bulk_with_proration",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(PriceModelGroupedTieredPackagePrice{}),
			DiscriminatorValue: "grouped_tiered_package",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(PriceModelMaxGroupTieredPackagePrice{}),
			DiscriminatorValue: "max_group_tiered_package",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(PriceModelScalableMatrixWithUnitPricingPrice{}),
			DiscriminatorValue: "scalable_matrix_with_unit_pricing",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(PriceModelScalableMatrixWithTieredPricingPrice{}),
			DiscriminatorValue: "scalable_matrix_with_tiered_pricing",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(PriceModelCumulativeGroupedBulkPrice{}),
			DiscriminatorValue: "cumulative_grouped_bulk",
		},
	)
}

type PriceModelUnitPrice struct {
	ID                          string                         `json:"id,required"`
	BillableMetric              BillableMetricTinyModel        `json:"billable_metric,required,nullable"`
	BillingCycleConfiguration   BillingCycleConfigurationModel `json:"billing_cycle_configuration,required"`
	Cadence                     PriceModelUnitPriceCadence     `json:"cadence,required"`
	ConversionRate              float64                        `json:"conversion_rate,required,nullable"`
	CreatedAt                   time.Time                      `json:"created_at,required" format:"date-time"`
	CreditAllocation            AllocationModel                `json:"credit_allocation,required,nullable"`
	Currency                    string                         `json:"currency,required"`
	Discount                    Discount                       `json:"discount,required,nullable"`
	ExternalPriceID             string                         `json:"external_price_id,required,nullable"`
	FixedPriceQuantity          float64                        `json:"fixed_price_quantity,required,nullable"`
	InvoicingCycleConfiguration BillingCycleConfigurationModel `json:"invoicing_cycle_configuration,required,nullable"`
	Item                        ItemSlimModel                  `json:"item,required"`
	Maximum                     MaximumModel                   `json:"maximum,required,nullable"`
	MaximumAmount               string                         `json:"maximum_amount,required,nullable"`
	// User specified key-value pairs for the resource. If not present, this defaults
	// to an empty dictionary. Individual keys can be removed by setting the value to
	// `null`, and the entire metadata mapping can be cleared by setting `metadata` to
	// `null`.
	Metadata                      map[string]string                  `json:"metadata,required"`
	Minimum                       MinimumModel                       `json:"minimum,required,nullable"`
	MinimumAmount                 string                             `json:"minimum_amount,required,nullable"`
	ModelType                     PriceModelUnitPriceModelType       `json:"model_type,required"`
	Name                          string                             `json:"name,required"`
	PlanPhaseOrder                int64                              `json:"plan_phase_order,required,nullable"`
	PriceType                     PriceModelUnitPricePriceType       `json:"price_type,required"`
	UnitConfig                    UnitConfigModel                    `json:"unit_config,required"`
	DimensionalPriceConfiguration DimensionalPriceConfigurationModel `json:"dimensional_price_configuration,nullable"`
	JSON                          priceModelUnitPriceJSON            `json:"-"`
}

// priceModelUnitPriceJSON contains the JSON metadata for the struct
// [PriceModelUnitPrice]
type priceModelUnitPriceJSON struct {
	ID                            apijson.Field
	BillableMetric                apijson.Field
	BillingCycleConfiguration     apijson.Field
	Cadence                       apijson.Field
	ConversionRate                apijson.Field
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
	UnitConfig                    apijson.Field
	DimensionalPriceConfiguration apijson.Field
	raw                           string
	ExtraFields                   map[string]apijson.Field
}

func (r *PriceModelUnitPrice) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r priceModelUnitPriceJSON) RawJSON() string {
	return r.raw
}

func (r PriceModelUnitPrice) implementsPriceModel() {}

type PriceModelUnitPriceCadence string

const (
	PriceModelUnitPriceCadenceOneTime    PriceModelUnitPriceCadence = "one_time"
	PriceModelUnitPriceCadenceMonthly    PriceModelUnitPriceCadence = "monthly"
	PriceModelUnitPriceCadenceQuarterly  PriceModelUnitPriceCadence = "quarterly"
	PriceModelUnitPriceCadenceSemiAnnual PriceModelUnitPriceCadence = "semi_annual"
	PriceModelUnitPriceCadenceAnnual     PriceModelUnitPriceCadence = "annual"
	PriceModelUnitPriceCadenceCustom     PriceModelUnitPriceCadence = "custom"
)

func (r PriceModelUnitPriceCadence) IsKnown() bool {
	switch r {
	case PriceModelUnitPriceCadenceOneTime, PriceModelUnitPriceCadenceMonthly, PriceModelUnitPriceCadenceQuarterly, PriceModelUnitPriceCadenceSemiAnnual, PriceModelUnitPriceCadenceAnnual, PriceModelUnitPriceCadenceCustom:
		return true
	}
	return false
}

type PriceModelUnitPriceModelType string

const (
	PriceModelUnitPriceModelTypeUnit PriceModelUnitPriceModelType = "unit"
)

func (r PriceModelUnitPriceModelType) IsKnown() bool {
	switch r {
	case PriceModelUnitPriceModelTypeUnit:
		return true
	}
	return false
}

type PriceModelUnitPricePriceType string

const (
	PriceModelUnitPricePriceTypeUsagePrice PriceModelUnitPricePriceType = "usage_price"
	PriceModelUnitPricePriceTypeFixedPrice PriceModelUnitPricePriceType = "fixed_price"
)

func (r PriceModelUnitPricePriceType) IsKnown() bool {
	switch r {
	case PriceModelUnitPricePriceTypeUsagePrice, PriceModelUnitPricePriceTypeFixedPrice:
		return true
	}
	return false
}

type PriceModelPackagePrice struct {
	ID                          string                         `json:"id,required"`
	BillableMetric              BillableMetricTinyModel        `json:"billable_metric,required,nullable"`
	BillingCycleConfiguration   BillingCycleConfigurationModel `json:"billing_cycle_configuration,required"`
	Cadence                     PriceModelPackagePriceCadence  `json:"cadence,required"`
	ConversionRate              float64                        `json:"conversion_rate,required,nullable"`
	CreatedAt                   time.Time                      `json:"created_at,required" format:"date-time"`
	CreditAllocation            AllocationModel                `json:"credit_allocation,required,nullable"`
	Currency                    string                         `json:"currency,required"`
	Discount                    Discount                       `json:"discount,required,nullable"`
	ExternalPriceID             string                         `json:"external_price_id,required,nullable"`
	FixedPriceQuantity          float64                        `json:"fixed_price_quantity,required,nullable"`
	InvoicingCycleConfiguration BillingCycleConfigurationModel `json:"invoicing_cycle_configuration,required,nullable"`
	Item                        ItemSlimModel                  `json:"item,required"`
	Maximum                     MaximumModel                   `json:"maximum,required,nullable"`
	MaximumAmount               string                         `json:"maximum_amount,required,nullable"`
	// User specified key-value pairs for the resource. If not present, this defaults
	// to an empty dictionary. Individual keys can be removed by setting the value to
	// `null`, and the entire metadata mapping can be cleared by setting `metadata` to
	// `null`.
	Metadata                      map[string]string                  `json:"metadata,required"`
	Minimum                       MinimumModel                       `json:"minimum,required,nullable"`
	MinimumAmount                 string                             `json:"minimum_amount,required,nullable"`
	ModelType                     PriceModelPackagePriceModelType    `json:"model_type,required"`
	Name                          string                             `json:"name,required"`
	PackageConfig                 PackageConfigModel                 `json:"package_config,required"`
	PlanPhaseOrder                int64                              `json:"plan_phase_order,required,nullable"`
	PriceType                     PriceModelPackagePricePriceType    `json:"price_type,required"`
	DimensionalPriceConfiguration DimensionalPriceConfigurationModel `json:"dimensional_price_configuration,nullable"`
	JSON                          priceModelPackagePriceJSON         `json:"-"`
}

// priceModelPackagePriceJSON contains the JSON metadata for the struct
// [PriceModelPackagePrice]
type priceModelPackagePriceJSON struct {
	ID                            apijson.Field
	BillableMetric                apijson.Field
	BillingCycleConfiguration     apijson.Field
	Cadence                       apijson.Field
	ConversionRate                apijson.Field
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
	DimensionalPriceConfiguration apijson.Field
	raw                           string
	ExtraFields                   map[string]apijson.Field
}

func (r *PriceModelPackagePrice) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r priceModelPackagePriceJSON) RawJSON() string {
	return r.raw
}

func (r PriceModelPackagePrice) implementsPriceModel() {}

type PriceModelPackagePriceCadence string

const (
	PriceModelPackagePriceCadenceOneTime    PriceModelPackagePriceCadence = "one_time"
	PriceModelPackagePriceCadenceMonthly    PriceModelPackagePriceCadence = "monthly"
	PriceModelPackagePriceCadenceQuarterly  PriceModelPackagePriceCadence = "quarterly"
	PriceModelPackagePriceCadenceSemiAnnual PriceModelPackagePriceCadence = "semi_annual"
	PriceModelPackagePriceCadenceAnnual     PriceModelPackagePriceCadence = "annual"
	PriceModelPackagePriceCadenceCustom     PriceModelPackagePriceCadence = "custom"
)

func (r PriceModelPackagePriceCadence) IsKnown() bool {
	switch r {
	case PriceModelPackagePriceCadenceOneTime, PriceModelPackagePriceCadenceMonthly, PriceModelPackagePriceCadenceQuarterly, PriceModelPackagePriceCadenceSemiAnnual, PriceModelPackagePriceCadenceAnnual, PriceModelPackagePriceCadenceCustom:
		return true
	}
	return false
}

type PriceModelPackagePriceModelType string

const (
	PriceModelPackagePriceModelTypePackage PriceModelPackagePriceModelType = "package"
)

func (r PriceModelPackagePriceModelType) IsKnown() bool {
	switch r {
	case PriceModelPackagePriceModelTypePackage:
		return true
	}
	return false
}

type PriceModelPackagePricePriceType string

const (
	PriceModelPackagePricePriceTypeUsagePrice PriceModelPackagePricePriceType = "usage_price"
	PriceModelPackagePricePriceTypeFixedPrice PriceModelPackagePricePriceType = "fixed_price"
)

func (r PriceModelPackagePricePriceType) IsKnown() bool {
	switch r {
	case PriceModelPackagePricePriceTypeUsagePrice, PriceModelPackagePricePriceTypeFixedPrice:
		return true
	}
	return false
}

type PriceModelMatrixPrice struct {
	ID                          string                         `json:"id,required"`
	BillableMetric              BillableMetricTinyModel        `json:"billable_metric,required,nullable"`
	BillingCycleConfiguration   BillingCycleConfigurationModel `json:"billing_cycle_configuration,required"`
	Cadence                     PriceModelMatrixPriceCadence   `json:"cadence,required"`
	ConversionRate              float64                        `json:"conversion_rate,required,nullable"`
	CreatedAt                   time.Time                      `json:"created_at,required" format:"date-time"`
	CreditAllocation            AllocationModel                `json:"credit_allocation,required,nullable"`
	Currency                    string                         `json:"currency,required"`
	Discount                    Discount                       `json:"discount,required,nullable"`
	ExternalPriceID             string                         `json:"external_price_id,required,nullable"`
	FixedPriceQuantity          float64                        `json:"fixed_price_quantity,required,nullable"`
	InvoicingCycleConfiguration BillingCycleConfigurationModel `json:"invoicing_cycle_configuration,required,nullable"`
	Item                        ItemSlimModel                  `json:"item,required"`
	MatrixConfig                MatrixConfigModel              `json:"matrix_config,required"`
	Maximum                     MaximumModel                   `json:"maximum,required,nullable"`
	MaximumAmount               string                         `json:"maximum_amount,required,nullable"`
	// User specified key-value pairs for the resource. If not present, this defaults
	// to an empty dictionary. Individual keys can be removed by setting the value to
	// `null`, and the entire metadata mapping can be cleared by setting `metadata` to
	// `null`.
	Metadata                      map[string]string                  `json:"metadata,required"`
	Minimum                       MinimumModel                       `json:"minimum,required,nullable"`
	MinimumAmount                 string                             `json:"minimum_amount,required,nullable"`
	ModelType                     PriceModelMatrixPriceModelType     `json:"model_type,required"`
	Name                          string                             `json:"name,required"`
	PlanPhaseOrder                int64                              `json:"plan_phase_order,required,nullable"`
	PriceType                     PriceModelMatrixPricePriceType     `json:"price_type,required"`
	DimensionalPriceConfiguration DimensionalPriceConfigurationModel `json:"dimensional_price_configuration,nullable"`
	JSON                          priceModelMatrixPriceJSON          `json:"-"`
}

// priceModelMatrixPriceJSON contains the JSON metadata for the struct
// [PriceModelMatrixPrice]
type priceModelMatrixPriceJSON struct {
	ID                            apijson.Field
	BillableMetric                apijson.Field
	BillingCycleConfiguration     apijson.Field
	Cadence                       apijson.Field
	ConversionRate                apijson.Field
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
	DimensionalPriceConfiguration apijson.Field
	raw                           string
	ExtraFields                   map[string]apijson.Field
}

func (r *PriceModelMatrixPrice) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r priceModelMatrixPriceJSON) RawJSON() string {
	return r.raw
}

func (r PriceModelMatrixPrice) implementsPriceModel() {}

type PriceModelMatrixPriceCadence string

const (
	PriceModelMatrixPriceCadenceOneTime    PriceModelMatrixPriceCadence = "one_time"
	PriceModelMatrixPriceCadenceMonthly    PriceModelMatrixPriceCadence = "monthly"
	PriceModelMatrixPriceCadenceQuarterly  PriceModelMatrixPriceCadence = "quarterly"
	PriceModelMatrixPriceCadenceSemiAnnual PriceModelMatrixPriceCadence = "semi_annual"
	PriceModelMatrixPriceCadenceAnnual     PriceModelMatrixPriceCadence = "annual"
	PriceModelMatrixPriceCadenceCustom     PriceModelMatrixPriceCadence = "custom"
)

func (r PriceModelMatrixPriceCadence) IsKnown() bool {
	switch r {
	case PriceModelMatrixPriceCadenceOneTime, PriceModelMatrixPriceCadenceMonthly, PriceModelMatrixPriceCadenceQuarterly, PriceModelMatrixPriceCadenceSemiAnnual, PriceModelMatrixPriceCadenceAnnual, PriceModelMatrixPriceCadenceCustom:
		return true
	}
	return false
}

type PriceModelMatrixPriceModelType string

const (
	PriceModelMatrixPriceModelTypeMatrix PriceModelMatrixPriceModelType = "matrix"
)

func (r PriceModelMatrixPriceModelType) IsKnown() bool {
	switch r {
	case PriceModelMatrixPriceModelTypeMatrix:
		return true
	}
	return false
}

type PriceModelMatrixPricePriceType string

const (
	PriceModelMatrixPricePriceTypeUsagePrice PriceModelMatrixPricePriceType = "usage_price"
	PriceModelMatrixPricePriceTypeFixedPrice PriceModelMatrixPricePriceType = "fixed_price"
)

func (r PriceModelMatrixPricePriceType) IsKnown() bool {
	switch r {
	case PriceModelMatrixPricePriceTypeUsagePrice, PriceModelMatrixPricePriceTypeFixedPrice:
		return true
	}
	return false
}

type PriceModelTieredPrice struct {
	ID                          string                         `json:"id,required"`
	BillableMetric              BillableMetricTinyModel        `json:"billable_metric,required,nullable"`
	BillingCycleConfiguration   BillingCycleConfigurationModel `json:"billing_cycle_configuration,required"`
	Cadence                     PriceModelTieredPriceCadence   `json:"cadence,required"`
	ConversionRate              float64                        `json:"conversion_rate,required,nullable"`
	CreatedAt                   time.Time                      `json:"created_at,required" format:"date-time"`
	CreditAllocation            AllocationModel                `json:"credit_allocation,required,nullable"`
	Currency                    string                         `json:"currency,required"`
	Discount                    Discount                       `json:"discount,required,nullable"`
	ExternalPriceID             string                         `json:"external_price_id,required,nullable"`
	FixedPriceQuantity          float64                        `json:"fixed_price_quantity,required,nullable"`
	InvoicingCycleConfiguration BillingCycleConfigurationModel `json:"invoicing_cycle_configuration,required,nullable"`
	Item                        ItemSlimModel                  `json:"item,required"`
	Maximum                     MaximumModel                   `json:"maximum,required,nullable"`
	MaximumAmount               string                         `json:"maximum_amount,required,nullable"`
	// User specified key-value pairs for the resource. If not present, this defaults
	// to an empty dictionary. Individual keys can be removed by setting the value to
	// `null`, and the entire metadata mapping can be cleared by setting `metadata` to
	// `null`.
	Metadata                      map[string]string                  `json:"metadata,required"`
	Minimum                       MinimumModel                       `json:"minimum,required,nullable"`
	MinimumAmount                 string                             `json:"minimum_amount,required,nullable"`
	ModelType                     PriceModelTieredPriceModelType     `json:"model_type,required"`
	Name                          string                             `json:"name,required"`
	PlanPhaseOrder                int64                              `json:"plan_phase_order,required,nullable"`
	PriceType                     PriceModelTieredPricePriceType     `json:"price_type,required"`
	TieredConfig                  TieredConfigModel                  `json:"tiered_config,required"`
	DimensionalPriceConfiguration DimensionalPriceConfigurationModel `json:"dimensional_price_configuration,nullable"`
	JSON                          priceModelTieredPriceJSON          `json:"-"`
}

// priceModelTieredPriceJSON contains the JSON metadata for the struct
// [PriceModelTieredPrice]
type priceModelTieredPriceJSON struct {
	ID                            apijson.Field
	BillableMetric                apijson.Field
	BillingCycleConfiguration     apijson.Field
	Cadence                       apijson.Field
	ConversionRate                apijson.Field
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
	TieredConfig                  apijson.Field
	DimensionalPriceConfiguration apijson.Field
	raw                           string
	ExtraFields                   map[string]apijson.Field
}

func (r *PriceModelTieredPrice) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r priceModelTieredPriceJSON) RawJSON() string {
	return r.raw
}

func (r PriceModelTieredPrice) implementsPriceModel() {}

type PriceModelTieredPriceCadence string

const (
	PriceModelTieredPriceCadenceOneTime    PriceModelTieredPriceCadence = "one_time"
	PriceModelTieredPriceCadenceMonthly    PriceModelTieredPriceCadence = "monthly"
	PriceModelTieredPriceCadenceQuarterly  PriceModelTieredPriceCadence = "quarterly"
	PriceModelTieredPriceCadenceSemiAnnual PriceModelTieredPriceCadence = "semi_annual"
	PriceModelTieredPriceCadenceAnnual     PriceModelTieredPriceCadence = "annual"
	PriceModelTieredPriceCadenceCustom     PriceModelTieredPriceCadence = "custom"
)

func (r PriceModelTieredPriceCadence) IsKnown() bool {
	switch r {
	case PriceModelTieredPriceCadenceOneTime, PriceModelTieredPriceCadenceMonthly, PriceModelTieredPriceCadenceQuarterly, PriceModelTieredPriceCadenceSemiAnnual, PriceModelTieredPriceCadenceAnnual, PriceModelTieredPriceCadenceCustom:
		return true
	}
	return false
}

type PriceModelTieredPriceModelType string

const (
	PriceModelTieredPriceModelTypeTiered PriceModelTieredPriceModelType = "tiered"
)

func (r PriceModelTieredPriceModelType) IsKnown() bool {
	switch r {
	case PriceModelTieredPriceModelTypeTiered:
		return true
	}
	return false
}

type PriceModelTieredPricePriceType string

const (
	PriceModelTieredPricePriceTypeUsagePrice PriceModelTieredPricePriceType = "usage_price"
	PriceModelTieredPricePriceTypeFixedPrice PriceModelTieredPricePriceType = "fixed_price"
)

func (r PriceModelTieredPricePriceType) IsKnown() bool {
	switch r {
	case PriceModelTieredPricePriceTypeUsagePrice, PriceModelTieredPricePriceTypeFixedPrice:
		return true
	}
	return false
}

type PriceModelTieredBpsPrice struct {
	ID                          string                          `json:"id,required"`
	BillableMetric              BillableMetricTinyModel         `json:"billable_metric,required,nullable"`
	BillingCycleConfiguration   BillingCycleConfigurationModel  `json:"billing_cycle_configuration,required"`
	Cadence                     PriceModelTieredBpsPriceCadence `json:"cadence,required"`
	ConversionRate              float64                         `json:"conversion_rate,required,nullable"`
	CreatedAt                   time.Time                       `json:"created_at,required" format:"date-time"`
	CreditAllocation            AllocationModel                 `json:"credit_allocation,required,nullable"`
	Currency                    string                          `json:"currency,required"`
	Discount                    Discount                        `json:"discount,required,nullable"`
	ExternalPriceID             string                          `json:"external_price_id,required,nullable"`
	FixedPriceQuantity          float64                         `json:"fixed_price_quantity,required,nullable"`
	InvoicingCycleConfiguration BillingCycleConfigurationModel  `json:"invoicing_cycle_configuration,required,nullable"`
	Item                        ItemSlimModel                   `json:"item,required"`
	Maximum                     MaximumModel                    `json:"maximum,required,nullable"`
	MaximumAmount               string                          `json:"maximum_amount,required,nullable"`
	// User specified key-value pairs for the resource. If not present, this defaults
	// to an empty dictionary. Individual keys can be removed by setting the value to
	// `null`, and the entire metadata mapping can be cleared by setting `metadata` to
	// `null`.
	Metadata                      map[string]string                  `json:"metadata,required"`
	Minimum                       MinimumModel                       `json:"minimum,required,nullable"`
	MinimumAmount                 string                             `json:"minimum_amount,required,nullable"`
	ModelType                     PriceModelTieredBpsPriceModelType  `json:"model_type,required"`
	Name                          string                             `json:"name,required"`
	PlanPhaseOrder                int64                              `json:"plan_phase_order,required,nullable"`
	PriceType                     PriceModelTieredBpsPricePriceType  `json:"price_type,required"`
	TieredBpsConfig               TieredBpsConfigModel               `json:"tiered_bps_config,required"`
	DimensionalPriceConfiguration DimensionalPriceConfigurationModel `json:"dimensional_price_configuration,nullable"`
	JSON                          priceModelTieredBpsPriceJSON       `json:"-"`
}

// priceModelTieredBpsPriceJSON contains the JSON metadata for the struct
// [PriceModelTieredBpsPrice]
type priceModelTieredBpsPriceJSON struct {
	ID                            apijson.Field
	BillableMetric                apijson.Field
	BillingCycleConfiguration     apijson.Field
	Cadence                       apijson.Field
	ConversionRate                apijson.Field
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
	TieredBpsConfig               apijson.Field
	DimensionalPriceConfiguration apijson.Field
	raw                           string
	ExtraFields                   map[string]apijson.Field
}

func (r *PriceModelTieredBpsPrice) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r priceModelTieredBpsPriceJSON) RawJSON() string {
	return r.raw
}

func (r PriceModelTieredBpsPrice) implementsPriceModel() {}

type PriceModelTieredBpsPriceCadence string

const (
	PriceModelTieredBpsPriceCadenceOneTime    PriceModelTieredBpsPriceCadence = "one_time"
	PriceModelTieredBpsPriceCadenceMonthly    PriceModelTieredBpsPriceCadence = "monthly"
	PriceModelTieredBpsPriceCadenceQuarterly  PriceModelTieredBpsPriceCadence = "quarterly"
	PriceModelTieredBpsPriceCadenceSemiAnnual PriceModelTieredBpsPriceCadence = "semi_annual"
	PriceModelTieredBpsPriceCadenceAnnual     PriceModelTieredBpsPriceCadence = "annual"
	PriceModelTieredBpsPriceCadenceCustom     PriceModelTieredBpsPriceCadence = "custom"
)

func (r PriceModelTieredBpsPriceCadence) IsKnown() bool {
	switch r {
	case PriceModelTieredBpsPriceCadenceOneTime, PriceModelTieredBpsPriceCadenceMonthly, PriceModelTieredBpsPriceCadenceQuarterly, PriceModelTieredBpsPriceCadenceSemiAnnual, PriceModelTieredBpsPriceCadenceAnnual, PriceModelTieredBpsPriceCadenceCustom:
		return true
	}
	return false
}

type PriceModelTieredBpsPriceModelType string

const (
	PriceModelTieredBpsPriceModelTypeTieredBps PriceModelTieredBpsPriceModelType = "tiered_bps"
)

func (r PriceModelTieredBpsPriceModelType) IsKnown() bool {
	switch r {
	case PriceModelTieredBpsPriceModelTypeTieredBps:
		return true
	}
	return false
}

type PriceModelTieredBpsPricePriceType string

const (
	PriceModelTieredBpsPricePriceTypeUsagePrice PriceModelTieredBpsPricePriceType = "usage_price"
	PriceModelTieredBpsPricePriceTypeFixedPrice PriceModelTieredBpsPricePriceType = "fixed_price"
)

func (r PriceModelTieredBpsPricePriceType) IsKnown() bool {
	switch r {
	case PriceModelTieredBpsPricePriceTypeUsagePrice, PriceModelTieredBpsPricePriceTypeFixedPrice:
		return true
	}
	return false
}

type PriceModelBpsPrice struct {
	ID                          string                         `json:"id,required"`
	BillableMetric              BillableMetricTinyModel        `json:"billable_metric,required,nullable"`
	BillingCycleConfiguration   BillingCycleConfigurationModel `json:"billing_cycle_configuration,required"`
	BpsConfig                   BpsConfigModel                 `json:"bps_config,required"`
	Cadence                     PriceModelBpsPriceCadence      `json:"cadence,required"`
	ConversionRate              float64                        `json:"conversion_rate,required,nullable"`
	CreatedAt                   time.Time                      `json:"created_at,required" format:"date-time"`
	CreditAllocation            AllocationModel                `json:"credit_allocation,required,nullable"`
	Currency                    string                         `json:"currency,required"`
	Discount                    Discount                       `json:"discount,required,nullable"`
	ExternalPriceID             string                         `json:"external_price_id,required,nullable"`
	FixedPriceQuantity          float64                        `json:"fixed_price_quantity,required,nullable"`
	InvoicingCycleConfiguration BillingCycleConfigurationModel `json:"invoicing_cycle_configuration,required,nullable"`
	Item                        ItemSlimModel                  `json:"item,required"`
	Maximum                     MaximumModel                   `json:"maximum,required,nullable"`
	MaximumAmount               string                         `json:"maximum_amount,required,nullable"`
	// User specified key-value pairs for the resource. If not present, this defaults
	// to an empty dictionary. Individual keys can be removed by setting the value to
	// `null`, and the entire metadata mapping can be cleared by setting `metadata` to
	// `null`.
	Metadata                      map[string]string                  `json:"metadata,required"`
	Minimum                       MinimumModel                       `json:"minimum,required,nullable"`
	MinimumAmount                 string                             `json:"minimum_amount,required,nullable"`
	ModelType                     PriceModelBpsPriceModelType        `json:"model_type,required"`
	Name                          string                             `json:"name,required"`
	PlanPhaseOrder                int64                              `json:"plan_phase_order,required,nullable"`
	PriceType                     PriceModelBpsPricePriceType        `json:"price_type,required"`
	DimensionalPriceConfiguration DimensionalPriceConfigurationModel `json:"dimensional_price_configuration,nullable"`
	JSON                          priceModelBpsPriceJSON             `json:"-"`
}

// priceModelBpsPriceJSON contains the JSON metadata for the struct
// [PriceModelBpsPrice]
type priceModelBpsPriceJSON struct {
	ID                            apijson.Field
	BillableMetric                apijson.Field
	BillingCycleConfiguration     apijson.Field
	BpsConfig                     apijson.Field
	Cadence                       apijson.Field
	ConversionRate                apijson.Field
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
	DimensionalPriceConfiguration apijson.Field
	raw                           string
	ExtraFields                   map[string]apijson.Field
}

func (r *PriceModelBpsPrice) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r priceModelBpsPriceJSON) RawJSON() string {
	return r.raw
}

func (r PriceModelBpsPrice) implementsPriceModel() {}

type PriceModelBpsPriceCadence string

const (
	PriceModelBpsPriceCadenceOneTime    PriceModelBpsPriceCadence = "one_time"
	PriceModelBpsPriceCadenceMonthly    PriceModelBpsPriceCadence = "monthly"
	PriceModelBpsPriceCadenceQuarterly  PriceModelBpsPriceCadence = "quarterly"
	PriceModelBpsPriceCadenceSemiAnnual PriceModelBpsPriceCadence = "semi_annual"
	PriceModelBpsPriceCadenceAnnual     PriceModelBpsPriceCadence = "annual"
	PriceModelBpsPriceCadenceCustom     PriceModelBpsPriceCadence = "custom"
)

func (r PriceModelBpsPriceCadence) IsKnown() bool {
	switch r {
	case PriceModelBpsPriceCadenceOneTime, PriceModelBpsPriceCadenceMonthly, PriceModelBpsPriceCadenceQuarterly, PriceModelBpsPriceCadenceSemiAnnual, PriceModelBpsPriceCadenceAnnual, PriceModelBpsPriceCadenceCustom:
		return true
	}
	return false
}

type PriceModelBpsPriceModelType string

const (
	PriceModelBpsPriceModelTypeBps PriceModelBpsPriceModelType = "bps"
)

func (r PriceModelBpsPriceModelType) IsKnown() bool {
	switch r {
	case PriceModelBpsPriceModelTypeBps:
		return true
	}
	return false
}

type PriceModelBpsPricePriceType string

const (
	PriceModelBpsPricePriceTypeUsagePrice PriceModelBpsPricePriceType = "usage_price"
	PriceModelBpsPricePriceTypeFixedPrice PriceModelBpsPricePriceType = "fixed_price"
)

func (r PriceModelBpsPricePriceType) IsKnown() bool {
	switch r {
	case PriceModelBpsPricePriceTypeUsagePrice, PriceModelBpsPricePriceTypeFixedPrice:
		return true
	}
	return false
}

type PriceModelBulkBpsPrice struct {
	ID                          string                         `json:"id,required"`
	BillableMetric              BillableMetricTinyModel        `json:"billable_metric,required,nullable"`
	BillingCycleConfiguration   BillingCycleConfigurationModel `json:"billing_cycle_configuration,required"`
	BulkBpsConfig               BulkBpsConfigModel             `json:"bulk_bps_config,required"`
	Cadence                     PriceModelBulkBpsPriceCadence  `json:"cadence,required"`
	ConversionRate              float64                        `json:"conversion_rate,required,nullable"`
	CreatedAt                   time.Time                      `json:"created_at,required" format:"date-time"`
	CreditAllocation            AllocationModel                `json:"credit_allocation,required,nullable"`
	Currency                    string                         `json:"currency,required"`
	Discount                    Discount                       `json:"discount,required,nullable"`
	ExternalPriceID             string                         `json:"external_price_id,required,nullable"`
	FixedPriceQuantity          float64                        `json:"fixed_price_quantity,required,nullable"`
	InvoicingCycleConfiguration BillingCycleConfigurationModel `json:"invoicing_cycle_configuration,required,nullable"`
	Item                        ItemSlimModel                  `json:"item,required"`
	Maximum                     MaximumModel                   `json:"maximum,required,nullable"`
	MaximumAmount               string                         `json:"maximum_amount,required,nullable"`
	// User specified key-value pairs for the resource. If not present, this defaults
	// to an empty dictionary. Individual keys can be removed by setting the value to
	// `null`, and the entire metadata mapping can be cleared by setting `metadata` to
	// `null`.
	Metadata                      map[string]string                  `json:"metadata,required"`
	Minimum                       MinimumModel                       `json:"minimum,required,nullable"`
	MinimumAmount                 string                             `json:"minimum_amount,required,nullable"`
	ModelType                     PriceModelBulkBpsPriceModelType    `json:"model_type,required"`
	Name                          string                             `json:"name,required"`
	PlanPhaseOrder                int64                              `json:"plan_phase_order,required,nullable"`
	PriceType                     PriceModelBulkBpsPricePriceType    `json:"price_type,required"`
	DimensionalPriceConfiguration DimensionalPriceConfigurationModel `json:"dimensional_price_configuration,nullable"`
	JSON                          priceModelBulkBpsPriceJSON         `json:"-"`
}

// priceModelBulkBpsPriceJSON contains the JSON metadata for the struct
// [PriceModelBulkBpsPrice]
type priceModelBulkBpsPriceJSON struct {
	ID                            apijson.Field
	BillableMetric                apijson.Field
	BillingCycleConfiguration     apijson.Field
	BulkBpsConfig                 apijson.Field
	Cadence                       apijson.Field
	ConversionRate                apijson.Field
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
	DimensionalPriceConfiguration apijson.Field
	raw                           string
	ExtraFields                   map[string]apijson.Field
}

func (r *PriceModelBulkBpsPrice) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r priceModelBulkBpsPriceJSON) RawJSON() string {
	return r.raw
}

func (r PriceModelBulkBpsPrice) implementsPriceModel() {}

type PriceModelBulkBpsPriceCadence string

const (
	PriceModelBulkBpsPriceCadenceOneTime    PriceModelBulkBpsPriceCadence = "one_time"
	PriceModelBulkBpsPriceCadenceMonthly    PriceModelBulkBpsPriceCadence = "monthly"
	PriceModelBulkBpsPriceCadenceQuarterly  PriceModelBulkBpsPriceCadence = "quarterly"
	PriceModelBulkBpsPriceCadenceSemiAnnual PriceModelBulkBpsPriceCadence = "semi_annual"
	PriceModelBulkBpsPriceCadenceAnnual     PriceModelBulkBpsPriceCadence = "annual"
	PriceModelBulkBpsPriceCadenceCustom     PriceModelBulkBpsPriceCadence = "custom"
)

func (r PriceModelBulkBpsPriceCadence) IsKnown() bool {
	switch r {
	case PriceModelBulkBpsPriceCadenceOneTime, PriceModelBulkBpsPriceCadenceMonthly, PriceModelBulkBpsPriceCadenceQuarterly, PriceModelBulkBpsPriceCadenceSemiAnnual, PriceModelBulkBpsPriceCadenceAnnual, PriceModelBulkBpsPriceCadenceCustom:
		return true
	}
	return false
}

type PriceModelBulkBpsPriceModelType string

const (
	PriceModelBulkBpsPriceModelTypeBulkBps PriceModelBulkBpsPriceModelType = "bulk_bps"
)

func (r PriceModelBulkBpsPriceModelType) IsKnown() bool {
	switch r {
	case PriceModelBulkBpsPriceModelTypeBulkBps:
		return true
	}
	return false
}

type PriceModelBulkBpsPricePriceType string

const (
	PriceModelBulkBpsPricePriceTypeUsagePrice PriceModelBulkBpsPricePriceType = "usage_price"
	PriceModelBulkBpsPricePriceTypeFixedPrice PriceModelBulkBpsPricePriceType = "fixed_price"
)

func (r PriceModelBulkBpsPricePriceType) IsKnown() bool {
	switch r {
	case PriceModelBulkBpsPricePriceTypeUsagePrice, PriceModelBulkBpsPricePriceTypeFixedPrice:
		return true
	}
	return false
}

type PriceModelBulkPrice struct {
	ID                          string                         `json:"id,required"`
	BillableMetric              BillableMetricTinyModel        `json:"billable_metric,required,nullable"`
	BillingCycleConfiguration   BillingCycleConfigurationModel `json:"billing_cycle_configuration,required"`
	BulkConfig                  BulkConfigModel                `json:"bulk_config,required"`
	Cadence                     PriceModelBulkPriceCadence     `json:"cadence,required"`
	ConversionRate              float64                        `json:"conversion_rate,required,nullable"`
	CreatedAt                   time.Time                      `json:"created_at,required" format:"date-time"`
	CreditAllocation            AllocationModel                `json:"credit_allocation,required,nullable"`
	Currency                    string                         `json:"currency,required"`
	Discount                    Discount                       `json:"discount,required,nullable"`
	ExternalPriceID             string                         `json:"external_price_id,required,nullable"`
	FixedPriceQuantity          float64                        `json:"fixed_price_quantity,required,nullable"`
	InvoicingCycleConfiguration BillingCycleConfigurationModel `json:"invoicing_cycle_configuration,required,nullable"`
	Item                        ItemSlimModel                  `json:"item,required"`
	Maximum                     MaximumModel                   `json:"maximum,required,nullable"`
	MaximumAmount               string                         `json:"maximum_amount,required,nullable"`
	// User specified key-value pairs for the resource. If not present, this defaults
	// to an empty dictionary. Individual keys can be removed by setting the value to
	// `null`, and the entire metadata mapping can be cleared by setting `metadata` to
	// `null`.
	Metadata                      map[string]string                  `json:"metadata,required"`
	Minimum                       MinimumModel                       `json:"minimum,required,nullable"`
	MinimumAmount                 string                             `json:"minimum_amount,required,nullable"`
	ModelType                     PriceModelBulkPriceModelType       `json:"model_type,required"`
	Name                          string                             `json:"name,required"`
	PlanPhaseOrder                int64                              `json:"plan_phase_order,required,nullable"`
	PriceType                     PriceModelBulkPricePriceType       `json:"price_type,required"`
	DimensionalPriceConfiguration DimensionalPriceConfigurationModel `json:"dimensional_price_configuration,nullable"`
	JSON                          priceModelBulkPriceJSON            `json:"-"`
}

// priceModelBulkPriceJSON contains the JSON metadata for the struct
// [PriceModelBulkPrice]
type priceModelBulkPriceJSON struct {
	ID                            apijson.Field
	BillableMetric                apijson.Field
	BillingCycleConfiguration     apijson.Field
	BulkConfig                    apijson.Field
	Cadence                       apijson.Field
	ConversionRate                apijson.Field
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
	DimensionalPriceConfiguration apijson.Field
	raw                           string
	ExtraFields                   map[string]apijson.Field
}

func (r *PriceModelBulkPrice) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r priceModelBulkPriceJSON) RawJSON() string {
	return r.raw
}

func (r PriceModelBulkPrice) implementsPriceModel() {}

type PriceModelBulkPriceCadence string

const (
	PriceModelBulkPriceCadenceOneTime    PriceModelBulkPriceCadence = "one_time"
	PriceModelBulkPriceCadenceMonthly    PriceModelBulkPriceCadence = "monthly"
	PriceModelBulkPriceCadenceQuarterly  PriceModelBulkPriceCadence = "quarterly"
	PriceModelBulkPriceCadenceSemiAnnual PriceModelBulkPriceCadence = "semi_annual"
	PriceModelBulkPriceCadenceAnnual     PriceModelBulkPriceCadence = "annual"
	PriceModelBulkPriceCadenceCustom     PriceModelBulkPriceCadence = "custom"
)

func (r PriceModelBulkPriceCadence) IsKnown() bool {
	switch r {
	case PriceModelBulkPriceCadenceOneTime, PriceModelBulkPriceCadenceMonthly, PriceModelBulkPriceCadenceQuarterly, PriceModelBulkPriceCadenceSemiAnnual, PriceModelBulkPriceCadenceAnnual, PriceModelBulkPriceCadenceCustom:
		return true
	}
	return false
}

type PriceModelBulkPriceModelType string

const (
	PriceModelBulkPriceModelTypeBulk PriceModelBulkPriceModelType = "bulk"
)

func (r PriceModelBulkPriceModelType) IsKnown() bool {
	switch r {
	case PriceModelBulkPriceModelTypeBulk:
		return true
	}
	return false
}

type PriceModelBulkPricePriceType string

const (
	PriceModelBulkPricePriceTypeUsagePrice PriceModelBulkPricePriceType = "usage_price"
	PriceModelBulkPricePriceTypeFixedPrice PriceModelBulkPricePriceType = "fixed_price"
)

func (r PriceModelBulkPricePriceType) IsKnown() bool {
	switch r {
	case PriceModelBulkPricePriceTypeUsagePrice, PriceModelBulkPricePriceTypeFixedPrice:
		return true
	}
	return false
}

type PriceModelThresholdTotalAmountPrice struct {
	ID                          string                                     `json:"id,required"`
	BillableMetric              BillableMetricTinyModel                    `json:"billable_metric,required,nullable"`
	BillingCycleConfiguration   BillingCycleConfigurationModel             `json:"billing_cycle_configuration,required"`
	Cadence                     PriceModelThresholdTotalAmountPriceCadence `json:"cadence,required"`
	ConversionRate              float64                                    `json:"conversion_rate,required,nullable"`
	CreatedAt                   time.Time                                  `json:"created_at,required" format:"date-time"`
	CreditAllocation            AllocationModel                            `json:"credit_allocation,required,nullable"`
	Currency                    string                                     `json:"currency,required"`
	Discount                    Discount                                   `json:"discount,required,nullable"`
	ExternalPriceID             string                                     `json:"external_price_id,required,nullable"`
	FixedPriceQuantity          float64                                    `json:"fixed_price_quantity,required,nullable"`
	InvoicingCycleConfiguration BillingCycleConfigurationModel             `json:"invoicing_cycle_configuration,required,nullable"`
	Item                        ItemSlimModel                              `json:"item,required"`
	Maximum                     MaximumModel                               `json:"maximum,required,nullable"`
	MaximumAmount               string                                     `json:"maximum_amount,required,nullable"`
	// User specified key-value pairs for the resource. If not present, this defaults
	// to an empty dictionary. Individual keys can be removed by setting the value to
	// `null`, and the entire metadata mapping can be cleared by setting `metadata` to
	// `null`.
	Metadata                      map[string]string                            `json:"metadata,required"`
	Minimum                       MinimumModel                                 `json:"minimum,required,nullable"`
	MinimumAmount                 string                                       `json:"minimum_amount,required,nullable"`
	ModelType                     PriceModelThresholdTotalAmountPriceModelType `json:"model_type,required"`
	Name                          string                                       `json:"name,required"`
	PlanPhaseOrder                int64                                        `json:"plan_phase_order,required,nullable"`
	PriceType                     PriceModelThresholdTotalAmountPricePriceType `json:"price_type,required"`
	ThresholdTotalAmountConfig    CustomRatingFunctionConfigModel              `json:"threshold_total_amount_config,required"`
	DimensionalPriceConfiguration DimensionalPriceConfigurationModel           `json:"dimensional_price_configuration,nullable"`
	JSON                          priceModelThresholdTotalAmountPriceJSON      `json:"-"`
}

// priceModelThresholdTotalAmountPriceJSON contains the JSON metadata for the
// struct [PriceModelThresholdTotalAmountPrice]
type priceModelThresholdTotalAmountPriceJSON struct {
	ID                            apijson.Field
	BillableMetric                apijson.Field
	BillingCycleConfiguration     apijson.Field
	Cadence                       apijson.Field
	ConversionRate                apijson.Field
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
	ThresholdTotalAmountConfig    apijson.Field
	DimensionalPriceConfiguration apijson.Field
	raw                           string
	ExtraFields                   map[string]apijson.Field
}

func (r *PriceModelThresholdTotalAmountPrice) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r priceModelThresholdTotalAmountPriceJSON) RawJSON() string {
	return r.raw
}

func (r PriceModelThresholdTotalAmountPrice) implementsPriceModel() {}

type PriceModelThresholdTotalAmountPriceCadence string

const (
	PriceModelThresholdTotalAmountPriceCadenceOneTime    PriceModelThresholdTotalAmountPriceCadence = "one_time"
	PriceModelThresholdTotalAmountPriceCadenceMonthly    PriceModelThresholdTotalAmountPriceCadence = "monthly"
	PriceModelThresholdTotalAmountPriceCadenceQuarterly  PriceModelThresholdTotalAmountPriceCadence = "quarterly"
	PriceModelThresholdTotalAmountPriceCadenceSemiAnnual PriceModelThresholdTotalAmountPriceCadence = "semi_annual"
	PriceModelThresholdTotalAmountPriceCadenceAnnual     PriceModelThresholdTotalAmountPriceCadence = "annual"
	PriceModelThresholdTotalAmountPriceCadenceCustom     PriceModelThresholdTotalAmountPriceCadence = "custom"
)

func (r PriceModelThresholdTotalAmountPriceCadence) IsKnown() bool {
	switch r {
	case PriceModelThresholdTotalAmountPriceCadenceOneTime, PriceModelThresholdTotalAmountPriceCadenceMonthly, PriceModelThresholdTotalAmountPriceCadenceQuarterly, PriceModelThresholdTotalAmountPriceCadenceSemiAnnual, PriceModelThresholdTotalAmountPriceCadenceAnnual, PriceModelThresholdTotalAmountPriceCadenceCustom:
		return true
	}
	return false
}

type PriceModelThresholdTotalAmountPriceModelType string

const (
	PriceModelThresholdTotalAmountPriceModelTypeThresholdTotalAmount PriceModelThresholdTotalAmountPriceModelType = "threshold_total_amount"
)

func (r PriceModelThresholdTotalAmountPriceModelType) IsKnown() bool {
	switch r {
	case PriceModelThresholdTotalAmountPriceModelTypeThresholdTotalAmount:
		return true
	}
	return false
}

type PriceModelThresholdTotalAmountPricePriceType string

const (
	PriceModelThresholdTotalAmountPricePriceTypeUsagePrice PriceModelThresholdTotalAmountPricePriceType = "usage_price"
	PriceModelThresholdTotalAmountPricePriceTypeFixedPrice PriceModelThresholdTotalAmountPricePriceType = "fixed_price"
)

func (r PriceModelThresholdTotalAmountPricePriceType) IsKnown() bool {
	switch r {
	case PriceModelThresholdTotalAmountPricePriceTypeUsagePrice, PriceModelThresholdTotalAmountPricePriceTypeFixedPrice:
		return true
	}
	return false
}

type PriceModelTieredPackagePrice struct {
	ID                          string                              `json:"id,required"`
	BillableMetric              BillableMetricTinyModel             `json:"billable_metric,required,nullable"`
	BillingCycleConfiguration   BillingCycleConfigurationModel      `json:"billing_cycle_configuration,required"`
	Cadence                     PriceModelTieredPackagePriceCadence `json:"cadence,required"`
	ConversionRate              float64                             `json:"conversion_rate,required,nullable"`
	CreatedAt                   time.Time                           `json:"created_at,required" format:"date-time"`
	CreditAllocation            AllocationModel                     `json:"credit_allocation,required,nullable"`
	Currency                    string                              `json:"currency,required"`
	Discount                    Discount                            `json:"discount,required,nullable"`
	ExternalPriceID             string                              `json:"external_price_id,required,nullable"`
	FixedPriceQuantity          float64                             `json:"fixed_price_quantity,required,nullable"`
	InvoicingCycleConfiguration BillingCycleConfigurationModel      `json:"invoicing_cycle_configuration,required,nullable"`
	Item                        ItemSlimModel                       `json:"item,required"`
	Maximum                     MaximumModel                        `json:"maximum,required,nullable"`
	MaximumAmount               string                              `json:"maximum_amount,required,nullable"`
	// User specified key-value pairs for the resource. If not present, this defaults
	// to an empty dictionary. Individual keys can be removed by setting the value to
	// `null`, and the entire metadata mapping can be cleared by setting `metadata` to
	// `null`.
	Metadata                      map[string]string                     `json:"metadata,required"`
	Minimum                       MinimumModel                          `json:"minimum,required,nullable"`
	MinimumAmount                 string                                `json:"minimum_amount,required,nullable"`
	ModelType                     PriceModelTieredPackagePriceModelType `json:"model_type,required"`
	Name                          string                                `json:"name,required"`
	PlanPhaseOrder                int64                                 `json:"plan_phase_order,required,nullable"`
	PriceType                     PriceModelTieredPackagePricePriceType `json:"price_type,required"`
	TieredPackageConfig           CustomRatingFunctionConfigModel       `json:"tiered_package_config,required"`
	DimensionalPriceConfiguration DimensionalPriceConfigurationModel    `json:"dimensional_price_configuration,nullable"`
	JSON                          priceModelTieredPackagePriceJSON      `json:"-"`
}

// priceModelTieredPackagePriceJSON contains the JSON metadata for the struct
// [PriceModelTieredPackagePrice]
type priceModelTieredPackagePriceJSON struct {
	ID                            apijson.Field
	BillableMetric                apijson.Field
	BillingCycleConfiguration     apijson.Field
	Cadence                       apijson.Field
	ConversionRate                apijson.Field
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
	TieredPackageConfig           apijson.Field
	DimensionalPriceConfiguration apijson.Field
	raw                           string
	ExtraFields                   map[string]apijson.Field
}

func (r *PriceModelTieredPackagePrice) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r priceModelTieredPackagePriceJSON) RawJSON() string {
	return r.raw
}

func (r PriceModelTieredPackagePrice) implementsPriceModel() {}

type PriceModelTieredPackagePriceCadence string

const (
	PriceModelTieredPackagePriceCadenceOneTime    PriceModelTieredPackagePriceCadence = "one_time"
	PriceModelTieredPackagePriceCadenceMonthly    PriceModelTieredPackagePriceCadence = "monthly"
	PriceModelTieredPackagePriceCadenceQuarterly  PriceModelTieredPackagePriceCadence = "quarterly"
	PriceModelTieredPackagePriceCadenceSemiAnnual PriceModelTieredPackagePriceCadence = "semi_annual"
	PriceModelTieredPackagePriceCadenceAnnual     PriceModelTieredPackagePriceCadence = "annual"
	PriceModelTieredPackagePriceCadenceCustom     PriceModelTieredPackagePriceCadence = "custom"
)

func (r PriceModelTieredPackagePriceCadence) IsKnown() bool {
	switch r {
	case PriceModelTieredPackagePriceCadenceOneTime, PriceModelTieredPackagePriceCadenceMonthly, PriceModelTieredPackagePriceCadenceQuarterly, PriceModelTieredPackagePriceCadenceSemiAnnual, PriceModelTieredPackagePriceCadenceAnnual, PriceModelTieredPackagePriceCadenceCustom:
		return true
	}
	return false
}

type PriceModelTieredPackagePriceModelType string

const (
	PriceModelTieredPackagePriceModelTypeTieredPackage PriceModelTieredPackagePriceModelType = "tiered_package"
)

func (r PriceModelTieredPackagePriceModelType) IsKnown() bool {
	switch r {
	case PriceModelTieredPackagePriceModelTypeTieredPackage:
		return true
	}
	return false
}

type PriceModelTieredPackagePricePriceType string

const (
	PriceModelTieredPackagePricePriceTypeUsagePrice PriceModelTieredPackagePricePriceType = "usage_price"
	PriceModelTieredPackagePricePriceTypeFixedPrice PriceModelTieredPackagePricePriceType = "fixed_price"
)

func (r PriceModelTieredPackagePricePriceType) IsKnown() bool {
	switch r {
	case PriceModelTieredPackagePricePriceTypeUsagePrice, PriceModelTieredPackagePricePriceTypeFixedPrice:
		return true
	}
	return false
}

type PriceModelGroupedTieredPrice struct {
	ID                          string                              `json:"id,required"`
	BillableMetric              BillableMetricTinyModel             `json:"billable_metric,required,nullable"`
	BillingCycleConfiguration   BillingCycleConfigurationModel      `json:"billing_cycle_configuration,required"`
	Cadence                     PriceModelGroupedTieredPriceCadence `json:"cadence,required"`
	ConversionRate              float64                             `json:"conversion_rate,required,nullable"`
	CreatedAt                   time.Time                           `json:"created_at,required" format:"date-time"`
	CreditAllocation            AllocationModel                     `json:"credit_allocation,required,nullable"`
	Currency                    string                              `json:"currency,required"`
	Discount                    Discount                            `json:"discount,required,nullable"`
	ExternalPriceID             string                              `json:"external_price_id,required,nullable"`
	FixedPriceQuantity          float64                             `json:"fixed_price_quantity,required,nullable"`
	GroupedTieredConfig         CustomRatingFunctionConfigModel     `json:"grouped_tiered_config,required"`
	InvoicingCycleConfiguration BillingCycleConfigurationModel      `json:"invoicing_cycle_configuration,required,nullable"`
	Item                        ItemSlimModel                       `json:"item,required"`
	Maximum                     MaximumModel                        `json:"maximum,required,nullable"`
	MaximumAmount               string                              `json:"maximum_amount,required,nullable"`
	// User specified key-value pairs for the resource. If not present, this defaults
	// to an empty dictionary. Individual keys can be removed by setting the value to
	// `null`, and the entire metadata mapping can be cleared by setting `metadata` to
	// `null`.
	Metadata                      map[string]string                     `json:"metadata,required"`
	Minimum                       MinimumModel                          `json:"minimum,required,nullable"`
	MinimumAmount                 string                                `json:"minimum_amount,required,nullable"`
	ModelType                     PriceModelGroupedTieredPriceModelType `json:"model_type,required"`
	Name                          string                                `json:"name,required"`
	PlanPhaseOrder                int64                                 `json:"plan_phase_order,required,nullable"`
	PriceType                     PriceModelGroupedTieredPricePriceType `json:"price_type,required"`
	DimensionalPriceConfiguration DimensionalPriceConfigurationModel    `json:"dimensional_price_configuration,nullable"`
	JSON                          priceModelGroupedTieredPriceJSON      `json:"-"`
}

// priceModelGroupedTieredPriceJSON contains the JSON metadata for the struct
// [PriceModelGroupedTieredPrice]
type priceModelGroupedTieredPriceJSON struct {
	ID                            apijson.Field
	BillableMetric                apijson.Field
	BillingCycleConfiguration     apijson.Field
	Cadence                       apijson.Field
	ConversionRate                apijson.Field
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
	DimensionalPriceConfiguration apijson.Field
	raw                           string
	ExtraFields                   map[string]apijson.Field
}

func (r *PriceModelGroupedTieredPrice) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r priceModelGroupedTieredPriceJSON) RawJSON() string {
	return r.raw
}

func (r PriceModelGroupedTieredPrice) implementsPriceModel() {}

type PriceModelGroupedTieredPriceCadence string

const (
	PriceModelGroupedTieredPriceCadenceOneTime    PriceModelGroupedTieredPriceCadence = "one_time"
	PriceModelGroupedTieredPriceCadenceMonthly    PriceModelGroupedTieredPriceCadence = "monthly"
	PriceModelGroupedTieredPriceCadenceQuarterly  PriceModelGroupedTieredPriceCadence = "quarterly"
	PriceModelGroupedTieredPriceCadenceSemiAnnual PriceModelGroupedTieredPriceCadence = "semi_annual"
	PriceModelGroupedTieredPriceCadenceAnnual     PriceModelGroupedTieredPriceCadence = "annual"
	PriceModelGroupedTieredPriceCadenceCustom     PriceModelGroupedTieredPriceCadence = "custom"
)

func (r PriceModelGroupedTieredPriceCadence) IsKnown() bool {
	switch r {
	case PriceModelGroupedTieredPriceCadenceOneTime, PriceModelGroupedTieredPriceCadenceMonthly, PriceModelGroupedTieredPriceCadenceQuarterly, PriceModelGroupedTieredPriceCadenceSemiAnnual, PriceModelGroupedTieredPriceCadenceAnnual, PriceModelGroupedTieredPriceCadenceCustom:
		return true
	}
	return false
}

type PriceModelGroupedTieredPriceModelType string

const (
	PriceModelGroupedTieredPriceModelTypeGroupedTiered PriceModelGroupedTieredPriceModelType = "grouped_tiered"
)

func (r PriceModelGroupedTieredPriceModelType) IsKnown() bool {
	switch r {
	case PriceModelGroupedTieredPriceModelTypeGroupedTiered:
		return true
	}
	return false
}

type PriceModelGroupedTieredPricePriceType string

const (
	PriceModelGroupedTieredPricePriceTypeUsagePrice PriceModelGroupedTieredPricePriceType = "usage_price"
	PriceModelGroupedTieredPricePriceTypeFixedPrice PriceModelGroupedTieredPricePriceType = "fixed_price"
)

func (r PriceModelGroupedTieredPricePriceType) IsKnown() bool {
	switch r {
	case PriceModelGroupedTieredPricePriceTypeUsagePrice, PriceModelGroupedTieredPricePriceTypeFixedPrice:
		return true
	}
	return false
}

type PriceModelTieredWithMinimumPrice struct {
	ID                          string                                  `json:"id,required"`
	BillableMetric              BillableMetricTinyModel                 `json:"billable_metric,required,nullable"`
	BillingCycleConfiguration   BillingCycleConfigurationModel          `json:"billing_cycle_configuration,required"`
	Cadence                     PriceModelTieredWithMinimumPriceCadence `json:"cadence,required"`
	ConversionRate              float64                                 `json:"conversion_rate,required,nullable"`
	CreatedAt                   time.Time                               `json:"created_at,required" format:"date-time"`
	CreditAllocation            AllocationModel                         `json:"credit_allocation,required,nullable"`
	Currency                    string                                  `json:"currency,required"`
	Discount                    Discount                                `json:"discount,required,nullable"`
	ExternalPriceID             string                                  `json:"external_price_id,required,nullable"`
	FixedPriceQuantity          float64                                 `json:"fixed_price_quantity,required,nullable"`
	InvoicingCycleConfiguration BillingCycleConfigurationModel          `json:"invoicing_cycle_configuration,required,nullable"`
	Item                        ItemSlimModel                           `json:"item,required"`
	Maximum                     MaximumModel                            `json:"maximum,required,nullable"`
	MaximumAmount               string                                  `json:"maximum_amount,required,nullable"`
	// User specified key-value pairs for the resource. If not present, this defaults
	// to an empty dictionary. Individual keys can be removed by setting the value to
	// `null`, and the entire metadata mapping can be cleared by setting `metadata` to
	// `null`.
	Metadata                      map[string]string                         `json:"metadata,required"`
	Minimum                       MinimumModel                              `json:"minimum,required,nullable"`
	MinimumAmount                 string                                    `json:"minimum_amount,required,nullable"`
	ModelType                     PriceModelTieredWithMinimumPriceModelType `json:"model_type,required"`
	Name                          string                                    `json:"name,required"`
	PlanPhaseOrder                int64                                     `json:"plan_phase_order,required,nullable"`
	PriceType                     PriceModelTieredWithMinimumPricePriceType `json:"price_type,required"`
	TieredWithMinimumConfig       CustomRatingFunctionConfigModel           `json:"tiered_with_minimum_config,required"`
	DimensionalPriceConfiguration DimensionalPriceConfigurationModel        `json:"dimensional_price_configuration,nullable"`
	JSON                          priceModelTieredWithMinimumPriceJSON      `json:"-"`
}

// priceModelTieredWithMinimumPriceJSON contains the JSON metadata for the struct
// [PriceModelTieredWithMinimumPrice]
type priceModelTieredWithMinimumPriceJSON struct {
	ID                            apijson.Field
	BillableMetric                apijson.Field
	BillingCycleConfiguration     apijson.Field
	Cadence                       apijson.Field
	ConversionRate                apijson.Field
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
	TieredWithMinimumConfig       apijson.Field
	DimensionalPriceConfiguration apijson.Field
	raw                           string
	ExtraFields                   map[string]apijson.Field
}

func (r *PriceModelTieredWithMinimumPrice) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r priceModelTieredWithMinimumPriceJSON) RawJSON() string {
	return r.raw
}

func (r PriceModelTieredWithMinimumPrice) implementsPriceModel() {}

type PriceModelTieredWithMinimumPriceCadence string

const (
	PriceModelTieredWithMinimumPriceCadenceOneTime    PriceModelTieredWithMinimumPriceCadence = "one_time"
	PriceModelTieredWithMinimumPriceCadenceMonthly    PriceModelTieredWithMinimumPriceCadence = "monthly"
	PriceModelTieredWithMinimumPriceCadenceQuarterly  PriceModelTieredWithMinimumPriceCadence = "quarterly"
	PriceModelTieredWithMinimumPriceCadenceSemiAnnual PriceModelTieredWithMinimumPriceCadence = "semi_annual"
	PriceModelTieredWithMinimumPriceCadenceAnnual     PriceModelTieredWithMinimumPriceCadence = "annual"
	PriceModelTieredWithMinimumPriceCadenceCustom     PriceModelTieredWithMinimumPriceCadence = "custom"
)

func (r PriceModelTieredWithMinimumPriceCadence) IsKnown() bool {
	switch r {
	case PriceModelTieredWithMinimumPriceCadenceOneTime, PriceModelTieredWithMinimumPriceCadenceMonthly, PriceModelTieredWithMinimumPriceCadenceQuarterly, PriceModelTieredWithMinimumPriceCadenceSemiAnnual, PriceModelTieredWithMinimumPriceCadenceAnnual, PriceModelTieredWithMinimumPriceCadenceCustom:
		return true
	}
	return false
}

type PriceModelTieredWithMinimumPriceModelType string

const (
	PriceModelTieredWithMinimumPriceModelTypeTieredWithMinimum PriceModelTieredWithMinimumPriceModelType = "tiered_with_minimum"
)

func (r PriceModelTieredWithMinimumPriceModelType) IsKnown() bool {
	switch r {
	case PriceModelTieredWithMinimumPriceModelTypeTieredWithMinimum:
		return true
	}
	return false
}

type PriceModelTieredWithMinimumPricePriceType string

const (
	PriceModelTieredWithMinimumPricePriceTypeUsagePrice PriceModelTieredWithMinimumPricePriceType = "usage_price"
	PriceModelTieredWithMinimumPricePriceTypeFixedPrice PriceModelTieredWithMinimumPricePriceType = "fixed_price"
)

func (r PriceModelTieredWithMinimumPricePriceType) IsKnown() bool {
	switch r {
	case PriceModelTieredWithMinimumPricePriceTypeUsagePrice, PriceModelTieredWithMinimumPricePriceTypeFixedPrice:
		return true
	}
	return false
}

type PriceModelTieredPackageWithMinimumPrice struct {
	ID                          string                                         `json:"id,required"`
	BillableMetric              BillableMetricTinyModel                        `json:"billable_metric,required,nullable"`
	BillingCycleConfiguration   BillingCycleConfigurationModel                 `json:"billing_cycle_configuration,required"`
	Cadence                     PriceModelTieredPackageWithMinimumPriceCadence `json:"cadence,required"`
	ConversionRate              float64                                        `json:"conversion_rate,required,nullable"`
	CreatedAt                   time.Time                                      `json:"created_at,required" format:"date-time"`
	CreditAllocation            AllocationModel                                `json:"credit_allocation,required,nullable"`
	Currency                    string                                         `json:"currency,required"`
	Discount                    Discount                                       `json:"discount,required,nullable"`
	ExternalPriceID             string                                         `json:"external_price_id,required,nullable"`
	FixedPriceQuantity          float64                                        `json:"fixed_price_quantity,required,nullable"`
	InvoicingCycleConfiguration BillingCycleConfigurationModel                 `json:"invoicing_cycle_configuration,required,nullable"`
	Item                        ItemSlimModel                                  `json:"item,required"`
	Maximum                     MaximumModel                                   `json:"maximum,required,nullable"`
	MaximumAmount               string                                         `json:"maximum_amount,required,nullable"`
	// User specified key-value pairs for the resource. If not present, this defaults
	// to an empty dictionary. Individual keys can be removed by setting the value to
	// `null`, and the entire metadata mapping can be cleared by setting `metadata` to
	// `null`.
	Metadata                       map[string]string                                `json:"metadata,required"`
	Minimum                        MinimumModel                                     `json:"minimum,required,nullable"`
	MinimumAmount                  string                                           `json:"minimum_amount,required,nullable"`
	ModelType                      PriceModelTieredPackageWithMinimumPriceModelType `json:"model_type,required"`
	Name                           string                                           `json:"name,required"`
	PlanPhaseOrder                 int64                                            `json:"plan_phase_order,required,nullable"`
	PriceType                      PriceModelTieredPackageWithMinimumPricePriceType `json:"price_type,required"`
	TieredPackageWithMinimumConfig CustomRatingFunctionConfigModel                  `json:"tiered_package_with_minimum_config,required"`
	DimensionalPriceConfiguration  DimensionalPriceConfigurationModel               `json:"dimensional_price_configuration,nullable"`
	JSON                           priceModelTieredPackageWithMinimumPriceJSON      `json:"-"`
}

// priceModelTieredPackageWithMinimumPriceJSON contains the JSON metadata for the
// struct [PriceModelTieredPackageWithMinimumPrice]
type priceModelTieredPackageWithMinimumPriceJSON struct {
	ID                             apijson.Field
	BillableMetric                 apijson.Field
	BillingCycleConfiguration      apijson.Field
	Cadence                        apijson.Field
	ConversionRate                 apijson.Field
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
	TieredPackageWithMinimumConfig apijson.Field
	DimensionalPriceConfiguration  apijson.Field
	raw                            string
	ExtraFields                    map[string]apijson.Field
}

func (r *PriceModelTieredPackageWithMinimumPrice) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r priceModelTieredPackageWithMinimumPriceJSON) RawJSON() string {
	return r.raw
}

func (r PriceModelTieredPackageWithMinimumPrice) implementsPriceModel() {}

type PriceModelTieredPackageWithMinimumPriceCadence string

const (
	PriceModelTieredPackageWithMinimumPriceCadenceOneTime    PriceModelTieredPackageWithMinimumPriceCadence = "one_time"
	PriceModelTieredPackageWithMinimumPriceCadenceMonthly    PriceModelTieredPackageWithMinimumPriceCadence = "monthly"
	PriceModelTieredPackageWithMinimumPriceCadenceQuarterly  PriceModelTieredPackageWithMinimumPriceCadence = "quarterly"
	PriceModelTieredPackageWithMinimumPriceCadenceSemiAnnual PriceModelTieredPackageWithMinimumPriceCadence = "semi_annual"
	PriceModelTieredPackageWithMinimumPriceCadenceAnnual     PriceModelTieredPackageWithMinimumPriceCadence = "annual"
	PriceModelTieredPackageWithMinimumPriceCadenceCustom     PriceModelTieredPackageWithMinimumPriceCadence = "custom"
)

func (r PriceModelTieredPackageWithMinimumPriceCadence) IsKnown() bool {
	switch r {
	case PriceModelTieredPackageWithMinimumPriceCadenceOneTime, PriceModelTieredPackageWithMinimumPriceCadenceMonthly, PriceModelTieredPackageWithMinimumPriceCadenceQuarterly, PriceModelTieredPackageWithMinimumPriceCadenceSemiAnnual, PriceModelTieredPackageWithMinimumPriceCadenceAnnual, PriceModelTieredPackageWithMinimumPriceCadenceCustom:
		return true
	}
	return false
}

type PriceModelTieredPackageWithMinimumPriceModelType string

const (
	PriceModelTieredPackageWithMinimumPriceModelTypeTieredPackageWithMinimum PriceModelTieredPackageWithMinimumPriceModelType = "tiered_package_with_minimum"
)

func (r PriceModelTieredPackageWithMinimumPriceModelType) IsKnown() bool {
	switch r {
	case PriceModelTieredPackageWithMinimumPriceModelTypeTieredPackageWithMinimum:
		return true
	}
	return false
}

type PriceModelTieredPackageWithMinimumPricePriceType string

const (
	PriceModelTieredPackageWithMinimumPricePriceTypeUsagePrice PriceModelTieredPackageWithMinimumPricePriceType = "usage_price"
	PriceModelTieredPackageWithMinimumPricePriceTypeFixedPrice PriceModelTieredPackageWithMinimumPricePriceType = "fixed_price"
)

func (r PriceModelTieredPackageWithMinimumPricePriceType) IsKnown() bool {
	switch r {
	case PriceModelTieredPackageWithMinimumPricePriceTypeUsagePrice, PriceModelTieredPackageWithMinimumPricePriceTypeFixedPrice:
		return true
	}
	return false
}

type PriceModelPackageWithAllocationPrice struct {
	ID                          string                                      `json:"id,required"`
	BillableMetric              BillableMetricTinyModel                     `json:"billable_metric,required,nullable"`
	BillingCycleConfiguration   BillingCycleConfigurationModel              `json:"billing_cycle_configuration,required"`
	Cadence                     PriceModelPackageWithAllocationPriceCadence `json:"cadence,required"`
	ConversionRate              float64                                     `json:"conversion_rate,required,nullable"`
	CreatedAt                   time.Time                                   `json:"created_at,required" format:"date-time"`
	CreditAllocation            AllocationModel                             `json:"credit_allocation,required,nullable"`
	Currency                    string                                      `json:"currency,required"`
	Discount                    Discount                                    `json:"discount,required,nullable"`
	ExternalPriceID             string                                      `json:"external_price_id,required,nullable"`
	FixedPriceQuantity          float64                                     `json:"fixed_price_quantity,required,nullable"`
	InvoicingCycleConfiguration BillingCycleConfigurationModel              `json:"invoicing_cycle_configuration,required,nullable"`
	Item                        ItemSlimModel                               `json:"item,required"`
	Maximum                     MaximumModel                                `json:"maximum,required,nullable"`
	MaximumAmount               string                                      `json:"maximum_amount,required,nullable"`
	// User specified key-value pairs for the resource. If not present, this defaults
	// to an empty dictionary. Individual keys can be removed by setting the value to
	// `null`, and the entire metadata mapping can be cleared by setting `metadata` to
	// `null`.
	Metadata                      map[string]string                             `json:"metadata,required"`
	Minimum                       MinimumModel                                  `json:"minimum,required,nullable"`
	MinimumAmount                 string                                        `json:"minimum_amount,required,nullable"`
	ModelType                     PriceModelPackageWithAllocationPriceModelType `json:"model_type,required"`
	Name                          string                                        `json:"name,required"`
	PackageWithAllocationConfig   CustomRatingFunctionConfigModel               `json:"package_with_allocation_config,required"`
	PlanPhaseOrder                int64                                         `json:"plan_phase_order,required,nullable"`
	PriceType                     PriceModelPackageWithAllocationPricePriceType `json:"price_type,required"`
	DimensionalPriceConfiguration DimensionalPriceConfigurationModel            `json:"dimensional_price_configuration,nullable"`
	JSON                          priceModelPackageWithAllocationPriceJSON      `json:"-"`
}

// priceModelPackageWithAllocationPriceJSON contains the JSON metadata for the
// struct [PriceModelPackageWithAllocationPrice]
type priceModelPackageWithAllocationPriceJSON struct {
	ID                            apijson.Field
	BillableMetric                apijson.Field
	BillingCycleConfiguration     apijson.Field
	Cadence                       apijson.Field
	ConversionRate                apijson.Field
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
	DimensionalPriceConfiguration apijson.Field
	raw                           string
	ExtraFields                   map[string]apijson.Field
}

func (r *PriceModelPackageWithAllocationPrice) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r priceModelPackageWithAllocationPriceJSON) RawJSON() string {
	return r.raw
}

func (r PriceModelPackageWithAllocationPrice) implementsPriceModel() {}

type PriceModelPackageWithAllocationPriceCadence string

const (
	PriceModelPackageWithAllocationPriceCadenceOneTime    PriceModelPackageWithAllocationPriceCadence = "one_time"
	PriceModelPackageWithAllocationPriceCadenceMonthly    PriceModelPackageWithAllocationPriceCadence = "monthly"
	PriceModelPackageWithAllocationPriceCadenceQuarterly  PriceModelPackageWithAllocationPriceCadence = "quarterly"
	PriceModelPackageWithAllocationPriceCadenceSemiAnnual PriceModelPackageWithAllocationPriceCadence = "semi_annual"
	PriceModelPackageWithAllocationPriceCadenceAnnual     PriceModelPackageWithAllocationPriceCadence = "annual"
	PriceModelPackageWithAllocationPriceCadenceCustom     PriceModelPackageWithAllocationPriceCadence = "custom"
)

func (r PriceModelPackageWithAllocationPriceCadence) IsKnown() bool {
	switch r {
	case PriceModelPackageWithAllocationPriceCadenceOneTime, PriceModelPackageWithAllocationPriceCadenceMonthly, PriceModelPackageWithAllocationPriceCadenceQuarterly, PriceModelPackageWithAllocationPriceCadenceSemiAnnual, PriceModelPackageWithAllocationPriceCadenceAnnual, PriceModelPackageWithAllocationPriceCadenceCustom:
		return true
	}
	return false
}

type PriceModelPackageWithAllocationPriceModelType string

const (
	PriceModelPackageWithAllocationPriceModelTypePackageWithAllocation PriceModelPackageWithAllocationPriceModelType = "package_with_allocation"
)

func (r PriceModelPackageWithAllocationPriceModelType) IsKnown() bool {
	switch r {
	case PriceModelPackageWithAllocationPriceModelTypePackageWithAllocation:
		return true
	}
	return false
}

type PriceModelPackageWithAllocationPricePriceType string

const (
	PriceModelPackageWithAllocationPricePriceTypeUsagePrice PriceModelPackageWithAllocationPricePriceType = "usage_price"
	PriceModelPackageWithAllocationPricePriceTypeFixedPrice PriceModelPackageWithAllocationPricePriceType = "fixed_price"
)

func (r PriceModelPackageWithAllocationPricePriceType) IsKnown() bool {
	switch r {
	case PriceModelPackageWithAllocationPricePriceTypeUsagePrice, PriceModelPackageWithAllocationPricePriceTypeFixedPrice:
		return true
	}
	return false
}

type PriceModelUnitWithPercentPrice struct {
	ID                          string                                `json:"id,required"`
	BillableMetric              BillableMetricTinyModel               `json:"billable_metric,required,nullable"`
	BillingCycleConfiguration   BillingCycleConfigurationModel        `json:"billing_cycle_configuration,required"`
	Cadence                     PriceModelUnitWithPercentPriceCadence `json:"cadence,required"`
	ConversionRate              float64                               `json:"conversion_rate,required,nullable"`
	CreatedAt                   time.Time                             `json:"created_at,required" format:"date-time"`
	CreditAllocation            AllocationModel                       `json:"credit_allocation,required,nullable"`
	Currency                    string                                `json:"currency,required"`
	Discount                    Discount                              `json:"discount,required,nullable"`
	ExternalPriceID             string                                `json:"external_price_id,required,nullable"`
	FixedPriceQuantity          float64                               `json:"fixed_price_quantity,required,nullable"`
	InvoicingCycleConfiguration BillingCycleConfigurationModel        `json:"invoicing_cycle_configuration,required,nullable"`
	Item                        ItemSlimModel                         `json:"item,required"`
	Maximum                     MaximumModel                          `json:"maximum,required,nullable"`
	MaximumAmount               string                                `json:"maximum_amount,required,nullable"`
	// User specified key-value pairs for the resource. If not present, this defaults
	// to an empty dictionary. Individual keys can be removed by setting the value to
	// `null`, and the entire metadata mapping can be cleared by setting `metadata` to
	// `null`.
	Metadata                      map[string]string                       `json:"metadata,required"`
	Minimum                       MinimumModel                            `json:"minimum,required,nullable"`
	MinimumAmount                 string                                  `json:"minimum_amount,required,nullable"`
	ModelType                     PriceModelUnitWithPercentPriceModelType `json:"model_type,required"`
	Name                          string                                  `json:"name,required"`
	PlanPhaseOrder                int64                                   `json:"plan_phase_order,required,nullable"`
	PriceType                     PriceModelUnitWithPercentPricePriceType `json:"price_type,required"`
	UnitWithPercentConfig         CustomRatingFunctionConfigModel         `json:"unit_with_percent_config,required"`
	DimensionalPriceConfiguration DimensionalPriceConfigurationModel      `json:"dimensional_price_configuration,nullable"`
	JSON                          priceModelUnitWithPercentPriceJSON      `json:"-"`
}

// priceModelUnitWithPercentPriceJSON contains the JSON metadata for the struct
// [PriceModelUnitWithPercentPrice]
type priceModelUnitWithPercentPriceJSON struct {
	ID                            apijson.Field
	BillableMetric                apijson.Field
	BillingCycleConfiguration     apijson.Field
	Cadence                       apijson.Field
	ConversionRate                apijson.Field
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
	UnitWithPercentConfig         apijson.Field
	DimensionalPriceConfiguration apijson.Field
	raw                           string
	ExtraFields                   map[string]apijson.Field
}

func (r *PriceModelUnitWithPercentPrice) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r priceModelUnitWithPercentPriceJSON) RawJSON() string {
	return r.raw
}

func (r PriceModelUnitWithPercentPrice) implementsPriceModel() {}

type PriceModelUnitWithPercentPriceCadence string

const (
	PriceModelUnitWithPercentPriceCadenceOneTime    PriceModelUnitWithPercentPriceCadence = "one_time"
	PriceModelUnitWithPercentPriceCadenceMonthly    PriceModelUnitWithPercentPriceCadence = "monthly"
	PriceModelUnitWithPercentPriceCadenceQuarterly  PriceModelUnitWithPercentPriceCadence = "quarterly"
	PriceModelUnitWithPercentPriceCadenceSemiAnnual PriceModelUnitWithPercentPriceCadence = "semi_annual"
	PriceModelUnitWithPercentPriceCadenceAnnual     PriceModelUnitWithPercentPriceCadence = "annual"
	PriceModelUnitWithPercentPriceCadenceCustom     PriceModelUnitWithPercentPriceCadence = "custom"
)

func (r PriceModelUnitWithPercentPriceCadence) IsKnown() bool {
	switch r {
	case PriceModelUnitWithPercentPriceCadenceOneTime, PriceModelUnitWithPercentPriceCadenceMonthly, PriceModelUnitWithPercentPriceCadenceQuarterly, PriceModelUnitWithPercentPriceCadenceSemiAnnual, PriceModelUnitWithPercentPriceCadenceAnnual, PriceModelUnitWithPercentPriceCadenceCustom:
		return true
	}
	return false
}

type PriceModelUnitWithPercentPriceModelType string

const (
	PriceModelUnitWithPercentPriceModelTypeUnitWithPercent PriceModelUnitWithPercentPriceModelType = "unit_with_percent"
)

func (r PriceModelUnitWithPercentPriceModelType) IsKnown() bool {
	switch r {
	case PriceModelUnitWithPercentPriceModelTypeUnitWithPercent:
		return true
	}
	return false
}

type PriceModelUnitWithPercentPricePriceType string

const (
	PriceModelUnitWithPercentPricePriceTypeUsagePrice PriceModelUnitWithPercentPricePriceType = "usage_price"
	PriceModelUnitWithPercentPricePriceTypeFixedPrice PriceModelUnitWithPercentPricePriceType = "fixed_price"
)

func (r PriceModelUnitWithPercentPricePriceType) IsKnown() bool {
	switch r {
	case PriceModelUnitWithPercentPricePriceTypeUsagePrice, PriceModelUnitWithPercentPricePriceTypeFixedPrice:
		return true
	}
	return false
}

type PriceModelMatrixWithAllocationPrice struct {
	ID                          string                                     `json:"id,required"`
	BillableMetric              BillableMetricTinyModel                    `json:"billable_metric,required,nullable"`
	BillingCycleConfiguration   BillingCycleConfigurationModel             `json:"billing_cycle_configuration,required"`
	Cadence                     PriceModelMatrixWithAllocationPriceCadence `json:"cadence,required"`
	ConversionRate              float64                                    `json:"conversion_rate,required,nullable"`
	CreatedAt                   time.Time                                  `json:"created_at,required" format:"date-time"`
	CreditAllocation            AllocationModel                            `json:"credit_allocation,required,nullable"`
	Currency                    string                                     `json:"currency,required"`
	Discount                    Discount                                   `json:"discount,required,nullable"`
	ExternalPriceID             string                                     `json:"external_price_id,required,nullable"`
	FixedPriceQuantity          float64                                    `json:"fixed_price_quantity,required,nullable"`
	InvoicingCycleConfiguration BillingCycleConfigurationModel             `json:"invoicing_cycle_configuration,required,nullable"`
	Item                        ItemSlimModel                              `json:"item,required"`
	MatrixWithAllocationConfig  MatrixWithAllocationConfigModel            `json:"matrix_with_allocation_config,required"`
	Maximum                     MaximumModel                               `json:"maximum,required,nullable"`
	MaximumAmount               string                                     `json:"maximum_amount,required,nullable"`
	// User specified key-value pairs for the resource. If not present, this defaults
	// to an empty dictionary. Individual keys can be removed by setting the value to
	// `null`, and the entire metadata mapping can be cleared by setting `metadata` to
	// `null`.
	Metadata                      map[string]string                            `json:"metadata,required"`
	Minimum                       MinimumModel                                 `json:"minimum,required,nullable"`
	MinimumAmount                 string                                       `json:"minimum_amount,required,nullable"`
	ModelType                     PriceModelMatrixWithAllocationPriceModelType `json:"model_type,required"`
	Name                          string                                       `json:"name,required"`
	PlanPhaseOrder                int64                                        `json:"plan_phase_order,required,nullable"`
	PriceType                     PriceModelMatrixWithAllocationPricePriceType `json:"price_type,required"`
	DimensionalPriceConfiguration DimensionalPriceConfigurationModel           `json:"dimensional_price_configuration,nullable"`
	JSON                          priceModelMatrixWithAllocationPriceJSON      `json:"-"`
}

// priceModelMatrixWithAllocationPriceJSON contains the JSON metadata for the
// struct [PriceModelMatrixWithAllocationPrice]
type priceModelMatrixWithAllocationPriceJSON struct {
	ID                            apijson.Field
	BillableMetric                apijson.Field
	BillingCycleConfiguration     apijson.Field
	Cadence                       apijson.Field
	ConversionRate                apijson.Field
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
	DimensionalPriceConfiguration apijson.Field
	raw                           string
	ExtraFields                   map[string]apijson.Field
}

func (r *PriceModelMatrixWithAllocationPrice) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r priceModelMatrixWithAllocationPriceJSON) RawJSON() string {
	return r.raw
}

func (r PriceModelMatrixWithAllocationPrice) implementsPriceModel() {}

type PriceModelMatrixWithAllocationPriceCadence string

const (
	PriceModelMatrixWithAllocationPriceCadenceOneTime    PriceModelMatrixWithAllocationPriceCadence = "one_time"
	PriceModelMatrixWithAllocationPriceCadenceMonthly    PriceModelMatrixWithAllocationPriceCadence = "monthly"
	PriceModelMatrixWithAllocationPriceCadenceQuarterly  PriceModelMatrixWithAllocationPriceCadence = "quarterly"
	PriceModelMatrixWithAllocationPriceCadenceSemiAnnual PriceModelMatrixWithAllocationPriceCadence = "semi_annual"
	PriceModelMatrixWithAllocationPriceCadenceAnnual     PriceModelMatrixWithAllocationPriceCadence = "annual"
	PriceModelMatrixWithAllocationPriceCadenceCustom     PriceModelMatrixWithAllocationPriceCadence = "custom"
)

func (r PriceModelMatrixWithAllocationPriceCadence) IsKnown() bool {
	switch r {
	case PriceModelMatrixWithAllocationPriceCadenceOneTime, PriceModelMatrixWithAllocationPriceCadenceMonthly, PriceModelMatrixWithAllocationPriceCadenceQuarterly, PriceModelMatrixWithAllocationPriceCadenceSemiAnnual, PriceModelMatrixWithAllocationPriceCadenceAnnual, PriceModelMatrixWithAllocationPriceCadenceCustom:
		return true
	}
	return false
}

type PriceModelMatrixWithAllocationPriceModelType string

const (
	PriceModelMatrixWithAllocationPriceModelTypeMatrixWithAllocation PriceModelMatrixWithAllocationPriceModelType = "matrix_with_allocation"
)

func (r PriceModelMatrixWithAllocationPriceModelType) IsKnown() bool {
	switch r {
	case PriceModelMatrixWithAllocationPriceModelTypeMatrixWithAllocation:
		return true
	}
	return false
}

type PriceModelMatrixWithAllocationPricePriceType string

const (
	PriceModelMatrixWithAllocationPricePriceTypeUsagePrice PriceModelMatrixWithAllocationPricePriceType = "usage_price"
	PriceModelMatrixWithAllocationPricePriceTypeFixedPrice PriceModelMatrixWithAllocationPricePriceType = "fixed_price"
)

func (r PriceModelMatrixWithAllocationPricePriceType) IsKnown() bool {
	switch r {
	case PriceModelMatrixWithAllocationPricePriceTypeUsagePrice, PriceModelMatrixWithAllocationPricePriceTypeFixedPrice:
		return true
	}
	return false
}

type PriceModelTieredWithProrationPrice struct {
	ID                          string                                    `json:"id,required"`
	BillableMetric              BillableMetricTinyModel                   `json:"billable_metric,required,nullable"`
	BillingCycleConfiguration   BillingCycleConfigurationModel            `json:"billing_cycle_configuration,required"`
	Cadence                     PriceModelTieredWithProrationPriceCadence `json:"cadence,required"`
	ConversionRate              float64                                   `json:"conversion_rate,required,nullable"`
	CreatedAt                   time.Time                                 `json:"created_at,required" format:"date-time"`
	CreditAllocation            AllocationModel                           `json:"credit_allocation,required,nullable"`
	Currency                    string                                    `json:"currency,required"`
	Discount                    Discount                                  `json:"discount,required,nullable"`
	ExternalPriceID             string                                    `json:"external_price_id,required,nullable"`
	FixedPriceQuantity          float64                                   `json:"fixed_price_quantity,required,nullable"`
	InvoicingCycleConfiguration BillingCycleConfigurationModel            `json:"invoicing_cycle_configuration,required,nullable"`
	Item                        ItemSlimModel                             `json:"item,required"`
	Maximum                     MaximumModel                              `json:"maximum,required,nullable"`
	MaximumAmount               string                                    `json:"maximum_amount,required,nullable"`
	// User specified key-value pairs for the resource. If not present, this defaults
	// to an empty dictionary. Individual keys can be removed by setting the value to
	// `null`, and the entire metadata mapping can be cleared by setting `metadata` to
	// `null`.
	Metadata                      map[string]string                           `json:"metadata,required"`
	Minimum                       MinimumModel                                `json:"minimum,required,nullable"`
	MinimumAmount                 string                                      `json:"minimum_amount,required,nullable"`
	ModelType                     PriceModelTieredWithProrationPriceModelType `json:"model_type,required"`
	Name                          string                                      `json:"name,required"`
	PlanPhaseOrder                int64                                       `json:"plan_phase_order,required,nullable"`
	PriceType                     PriceModelTieredWithProrationPricePriceType `json:"price_type,required"`
	TieredWithProrationConfig     CustomRatingFunctionConfigModel             `json:"tiered_with_proration_config,required"`
	DimensionalPriceConfiguration DimensionalPriceConfigurationModel          `json:"dimensional_price_configuration,nullable"`
	JSON                          priceModelTieredWithProrationPriceJSON      `json:"-"`
}

// priceModelTieredWithProrationPriceJSON contains the JSON metadata for the struct
// [PriceModelTieredWithProrationPrice]
type priceModelTieredWithProrationPriceJSON struct {
	ID                            apijson.Field
	BillableMetric                apijson.Field
	BillingCycleConfiguration     apijson.Field
	Cadence                       apijson.Field
	ConversionRate                apijson.Field
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
	TieredWithProrationConfig     apijson.Field
	DimensionalPriceConfiguration apijson.Field
	raw                           string
	ExtraFields                   map[string]apijson.Field
}

func (r *PriceModelTieredWithProrationPrice) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r priceModelTieredWithProrationPriceJSON) RawJSON() string {
	return r.raw
}

func (r PriceModelTieredWithProrationPrice) implementsPriceModel() {}

type PriceModelTieredWithProrationPriceCadence string

const (
	PriceModelTieredWithProrationPriceCadenceOneTime    PriceModelTieredWithProrationPriceCadence = "one_time"
	PriceModelTieredWithProrationPriceCadenceMonthly    PriceModelTieredWithProrationPriceCadence = "monthly"
	PriceModelTieredWithProrationPriceCadenceQuarterly  PriceModelTieredWithProrationPriceCadence = "quarterly"
	PriceModelTieredWithProrationPriceCadenceSemiAnnual PriceModelTieredWithProrationPriceCadence = "semi_annual"
	PriceModelTieredWithProrationPriceCadenceAnnual     PriceModelTieredWithProrationPriceCadence = "annual"
	PriceModelTieredWithProrationPriceCadenceCustom     PriceModelTieredWithProrationPriceCadence = "custom"
)

func (r PriceModelTieredWithProrationPriceCadence) IsKnown() bool {
	switch r {
	case PriceModelTieredWithProrationPriceCadenceOneTime, PriceModelTieredWithProrationPriceCadenceMonthly, PriceModelTieredWithProrationPriceCadenceQuarterly, PriceModelTieredWithProrationPriceCadenceSemiAnnual, PriceModelTieredWithProrationPriceCadenceAnnual, PriceModelTieredWithProrationPriceCadenceCustom:
		return true
	}
	return false
}

type PriceModelTieredWithProrationPriceModelType string

const (
	PriceModelTieredWithProrationPriceModelTypeTieredWithProration PriceModelTieredWithProrationPriceModelType = "tiered_with_proration"
)

func (r PriceModelTieredWithProrationPriceModelType) IsKnown() bool {
	switch r {
	case PriceModelTieredWithProrationPriceModelTypeTieredWithProration:
		return true
	}
	return false
}

type PriceModelTieredWithProrationPricePriceType string

const (
	PriceModelTieredWithProrationPricePriceTypeUsagePrice PriceModelTieredWithProrationPricePriceType = "usage_price"
	PriceModelTieredWithProrationPricePriceTypeFixedPrice PriceModelTieredWithProrationPricePriceType = "fixed_price"
)

func (r PriceModelTieredWithProrationPricePriceType) IsKnown() bool {
	switch r {
	case PriceModelTieredWithProrationPricePriceTypeUsagePrice, PriceModelTieredWithProrationPricePriceTypeFixedPrice:
		return true
	}
	return false
}

type PriceModelUnitWithProrationPrice struct {
	ID                          string                                  `json:"id,required"`
	BillableMetric              BillableMetricTinyModel                 `json:"billable_metric,required,nullable"`
	BillingCycleConfiguration   BillingCycleConfigurationModel          `json:"billing_cycle_configuration,required"`
	Cadence                     PriceModelUnitWithProrationPriceCadence `json:"cadence,required"`
	ConversionRate              float64                                 `json:"conversion_rate,required,nullable"`
	CreatedAt                   time.Time                               `json:"created_at,required" format:"date-time"`
	CreditAllocation            AllocationModel                         `json:"credit_allocation,required,nullable"`
	Currency                    string                                  `json:"currency,required"`
	Discount                    Discount                                `json:"discount,required,nullable"`
	ExternalPriceID             string                                  `json:"external_price_id,required,nullable"`
	FixedPriceQuantity          float64                                 `json:"fixed_price_quantity,required,nullable"`
	InvoicingCycleConfiguration BillingCycleConfigurationModel          `json:"invoicing_cycle_configuration,required,nullable"`
	Item                        ItemSlimModel                           `json:"item,required"`
	Maximum                     MaximumModel                            `json:"maximum,required,nullable"`
	MaximumAmount               string                                  `json:"maximum_amount,required,nullable"`
	// User specified key-value pairs for the resource. If not present, this defaults
	// to an empty dictionary. Individual keys can be removed by setting the value to
	// `null`, and the entire metadata mapping can be cleared by setting `metadata` to
	// `null`.
	Metadata                      map[string]string                         `json:"metadata,required"`
	Minimum                       MinimumModel                              `json:"minimum,required,nullable"`
	MinimumAmount                 string                                    `json:"minimum_amount,required,nullable"`
	ModelType                     PriceModelUnitWithProrationPriceModelType `json:"model_type,required"`
	Name                          string                                    `json:"name,required"`
	PlanPhaseOrder                int64                                     `json:"plan_phase_order,required,nullable"`
	PriceType                     PriceModelUnitWithProrationPricePriceType `json:"price_type,required"`
	UnitWithProrationConfig       CustomRatingFunctionConfigModel           `json:"unit_with_proration_config,required"`
	DimensionalPriceConfiguration DimensionalPriceConfigurationModel        `json:"dimensional_price_configuration,nullable"`
	JSON                          priceModelUnitWithProrationPriceJSON      `json:"-"`
}

// priceModelUnitWithProrationPriceJSON contains the JSON metadata for the struct
// [PriceModelUnitWithProrationPrice]
type priceModelUnitWithProrationPriceJSON struct {
	ID                            apijson.Field
	BillableMetric                apijson.Field
	BillingCycleConfiguration     apijson.Field
	Cadence                       apijson.Field
	ConversionRate                apijson.Field
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
	UnitWithProrationConfig       apijson.Field
	DimensionalPriceConfiguration apijson.Field
	raw                           string
	ExtraFields                   map[string]apijson.Field
}

func (r *PriceModelUnitWithProrationPrice) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r priceModelUnitWithProrationPriceJSON) RawJSON() string {
	return r.raw
}

func (r PriceModelUnitWithProrationPrice) implementsPriceModel() {}

type PriceModelUnitWithProrationPriceCadence string

const (
	PriceModelUnitWithProrationPriceCadenceOneTime    PriceModelUnitWithProrationPriceCadence = "one_time"
	PriceModelUnitWithProrationPriceCadenceMonthly    PriceModelUnitWithProrationPriceCadence = "monthly"
	PriceModelUnitWithProrationPriceCadenceQuarterly  PriceModelUnitWithProrationPriceCadence = "quarterly"
	PriceModelUnitWithProrationPriceCadenceSemiAnnual PriceModelUnitWithProrationPriceCadence = "semi_annual"
	PriceModelUnitWithProrationPriceCadenceAnnual     PriceModelUnitWithProrationPriceCadence = "annual"
	PriceModelUnitWithProrationPriceCadenceCustom     PriceModelUnitWithProrationPriceCadence = "custom"
)

func (r PriceModelUnitWithProrationPriceCadence) IsKnown() bool {
	switch r {
	case PriceModelUnitWithProrationPriceCadenceOneTime, PriceModelUnitWithProrationPriceCadenceMonthly, PriceModelUnitWithProrationPriceCadenceQuarterly, PriceModelUnitWithProrationPriceCadenceSemiAnnual, PriceModelUnitWithProrationPriceCadenceAnnual, PriceModelUnitWithProrationPriceCadenceCustom:
		return true
	}
	return false
}

type PriceModelUnitWithProrationPriceModelType string

const (
	PriceModelUnitWithProrationPriceModelTypeUnitWithProration PriceModelUnitWithProrationPriceModelType = "unit_with_proration"
)

func (r PriceModelUnitWithProrationPriceModelType) IsKnown() bool {
	switch r {
	case PriceModelUnitWithProrationPriceModelTypeUnitWithProration:
		return true
	}
	return false
}

type PriceModelUnitWithProrationPricePriceType string

const (
	PriceModelUnitWithProrationPricePriceTypeUsagePrice PriceModelUnitWithProrationPricePriceType = "usage_price"
	PriceModelUnitWithProrationPricePriceTypeFixedPrice PriceModelUnitWithProrationPricePriceType = "fixed_price"
)

func (r PriceModelUnitWithProrationPricePriceType) IsKnown() bool {
	switch r {
	case PriceModelUnitWithProrationPricePriceTypeUsagePrice, PriceModelUnitWithProrationPricePriceTypeFixedPrice:
		return true
	}
	return false
}

type PriceModelGroupedAllocationPrice struct {
	ID                          string                                  `json:"id,required"`
	BillableMetric              BillableMetricTinyModel                 `json:"billable_metric,required,nullable"`
	BillingCycleConfiguration   BillingCycleConfigurationModel          `json:"billing_cycle_configuration,required"`
	Cadence                     PriceModelGroupedAllocationPriceCadence `json:"cadence,required"`
	ConversionRate              float64                                 `json:"conversion_rate,required,nullable"`
	CreatedAt                   time.Time                               `json:"created_at,required" format:"date-time"`
	CreditAllocation            AllocationModel                         `json:"credit_allocation,required,nullable"`
	Currency                    string                                  `json:"currency,required"`
	Discount                    Discount                                `json:"discount,required,nullable"`
	ExternalPriceID             string                                  `json:"external_price_id,required,nullable"`
	FixedPriceQuantity          float64                                 `json:"fixed_price_quantity,required,nullable"`
	GroupedAllocationConfig     CustomRatingFunctionConfigModel         `json:"grouped_allocation_config,required"`
	InvoicingCycleConfiguration BillingCycleConfigurationModel          `json:"invoicing_cycle_configuration,required,nullable"`
	Item                        ItemSlimModel                           `json:"item,required"`
	Maximum                     MaximumModel                            `json:"maximum,required,nullable"`
	MaximumAmount               string                                  `json:"maximum_amount,required,nullable"`
	// User specified key-value pairs for the resource. If not present, this defaults
	// to an empty dictionary. Individual keys can be removed by setting the value to
	// `null`, and the entire metadata mapping can be cleared by setting `metadata` to
	// `null`.
	Metadata                      map[string]string                         `json:"metadata,required"`
	Minimum                       MinimumModel                              `json:"minimum,required,nullable"`
	MinimumAmount                 string                                    `json:"minimum_amount,required,nullable"`
	ModelType                     PriceModelGroupedAllocationPriceModelType `json:"model_type,required"`
	Name                          string                                    `json:"name,required"`
	PlanPhaseOrder                int64                                     `json:"plan_phase_order,required,nullable"`
	PriceType                     PriceModelGroupedAllocationPricePriceType `json:"price_type,required"`
	DimensionalPriceConfiguration DimensionalPriceConfigurationModel        `json:"dimensional_price_configuration,nullable"`
	JSON                          priceModelGroupedAllocationPriceJSON      `json:"-"`
}

// priceModelGroupedAllocationPriceJSON contains the JSON metadata for the struct
// [PriceModelGroupedAllocationPrice]
type priceModelGroupedAllocationPriceJSON struct {
	ID                            apijson.Field
	BillableMetric                apijson.Field
	BillingCycleConfiguration     apijson.Field
	Cadence                       apijson.Field
	ConversionRate                apijson.Field
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
	DimensionalPriceConfiguration apijson.Field
	raw                           string
	ExtraFields                   map[string]apijson.Field
}

func (r *PriceModelGroupedAllocationPrice) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r priceModelGroupedAllocationPriceJSON) RawJSON() string {
	return r.raw
}

func (r PriceModelGroupedAllocationPrice) implementsPriceModel() {}

type PriceModelGroupedAllocationPriceCadence string

const (
	PriceModelGroupedAllocationPriceCadenceOneTime    PriceModelGroupedAllocationPriceCadence = "one_time"
	PriceModelGroupedAllocationPriceCadenceMonthly    PriceModelGroupedAllocationPriceCadence = "monthly"
	PriceModelGroupedAllocationPriceCadenceQuarterly  PriceModelGroupedAllocationPriceCadence = "quarterly"
	PriceModelGroupedAllocationPriceCadenceSemiAnnual PriceModelGroupedAllocationPriceCadence = "semi_annual"
	PriceModelGroupedAllocationPriceCadenceAnnual     PriceModelGroupedAllocationPriceCadence = "annual"
	PriceModelGroupedAllocationPriceCadenceCustom     PriceModelGroupedAllocationPriceCadence = "custom"
)

func (r PriceModelGroupedAllocationPriceCadence) IsKnown() bool {
	switch r {
	case PriceModelGroupedAllocationPriceCadenceOneTime, PriceModelGroupedAllocationPriceCadenceMonthly, PriceModelGroupedAllocationPriceCadenceQuarterly, PriceModelGroupedAllocationPriceCadenceSemiAnnual, PriceModelGroupedAllocationPriceCadenceAnnual, PriceModelGroupedAllocationPriceCadenceCustom:
		return true
	}
	return false
}

type PriceModelGroupedAllocationPriceModelType string

const (
	PriceModelGroupedAllocationPriceModelTypeGroupedAllocation PriceModelGroupedAllocationPriceModelType = "grouped_allocation"
)

func (r PriceModelGroupedAllocationPriceModelType) IsKnown() bool {
	switch r {
	case PriceModelGroupedAllocationPriceModelTypeGroupedAllocation:
		return true
	}
	return false
}

type PriceModelGroupedAllocationPricePriceType string

const (
	PriceModelGroupedAllocationPricePriceTypeUsagePrice PriceModelGroupedAllocationPricePriceType = "usage_price"
	PriceModelGroupedAllocationPricePriceTypeFixedPrice PriceModelGroupedAllocationPricePriceType = "fixed_price"
)

func (r PriceModelGroupedAllocationPricePriceType) IsKnown() bool {
	switch r {
	case PriceModelGroupedAllocationPricePriceTypeUsagePrice, PriceModelGroupedAllocationPricePriceTypeFixedPrice:
		return true
	}
	return false
}

type PriceModelGroupedWithProratedMinimumPrice struct {
	ID                               string                                           `json:"id,required"`
	BillableMetric                   BillableMetricTinyModel                          `json:"billable_metric,required,nullable"`
	BillingCycleConfiguration        BillingCycleConfigurationModel                   `json:"billing_cycle_configuration,required"`
	Cadence                          PriceModelGroupedWithProratedMinimumPriceCadence `json:"cadence,required"`
	ConversionRate                   float64                                          `json:"conversion_rate,required,nullable"`
	CreatedAt                        time.Time                                        `json:"created_at,required" format:"date-time"`
	CreditAllocation                 AllocationModel                                  `json:"credit_allocation,required,nullable"`
	Currency                         string                                           `json:"currency,required"`
	Discount                         Discount                                         `json:"discount,required,nullable"`
	ExternalPriceID                  string                                           `json:"external_price_id,required,nullable"`
	FixedPriceQuantity               float64                                          `json:"fixed_price_quantity,required,nullable"`
	GroupedWithProratedMinimumConfig CustomRatingFunctionConfigModel                  `json:"grouped_with_prorated_minimum_config,required"`
	InvoicingCycleConfiguration      BillingCycleConfigurationModel                   `json:"invoicing_cycle_configuration,required,nullable"`
	Item                             ItemSlimModel                                    `json:"item,required"`
	Maximum                          MaximumModel                                     `json:"maximum,required,nullable"`
	MaximumAmount                    string                                           `json:"maximum_amount,required,nullable"`
	// User specified key-value pairs for the resource. If not present, this defaults
	// to an empty dictionary. Individual keys can be removed by setting the value to
	// `null`, and the entire metadata mapping can be cleared by setting `metadata` to
	// `null`.
	Metadata                      map[string]string                                  `json:"metadata,required"`
	Minimum                       MinimumModel                                       `json:"minimum,required,nullable"`
	MinimumAmount                 string                                             `json:"minimum_amount,required,nullable"`
	ModelType                     PriceModelGroupedWithProratedMinimumPriceModelType `json:"model_type,required"`
	Name                          string                                             `json:"name,required"`
	PlanPhaseOrder                int64                                              `json:"plan_phase_order,required,nullable"`
	PriceType                     PriceModelGroupedWithProratedMinimumPricePriceType `json:"price_type,required"`
	DimensionalPriceConfiguration DimensionalPriceConfigurationModel                 `json:"dimensional_price_configuration,nullable"`
	JSON                          priceModelGroupedWithProratedMinimumPriceJSON      `json:"-"`
}

// priceModelGroupedWithProratedMinimumPriceJSON contains the JSON metadata for the
// struct [PriceModelGroupedWithProratedMinimumPrice]
type priceModelGroupedWithProratedMinimumPriceJSON struct {
	ID                               apijson.Field
	BillableMetric                   apijson.Field
	BillingCycleConfiguration        apijson.Field
	Cadence                          apijson.Field
	ConversionRate                   apijson.Field
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
	DimensionalPriceConfiguration    apijson.Field
	raw                              string
	ExtraFields                      map[string]apijson.Field
}

func (r *PriceModelGroupedWithProratedMinimumPrice) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r priceModelGroupedWithProratedMinimumPriceJSON) RawJSON() string {
	return r.raw
}

func (r PriceModelGroupedWithProratedMinimumPrice) implementsPriceModel() {}

type PriceModelGroupedWithProratedMinimumPriceCadence string

const (
	PriceModelGroupedWithProratedMinimumPriceCadenceOneTime    PriceModelGroupedWithProratedMinimumPriceCadence = "one_time"
	PriceModelGroupedWithProratedMinimumPriceCadenceMonthly    PriceModelGroupedWithProratedMinimumPriceCadence = "monthly"
	PriceModelGroupedWithProratedMinimumPriceCadenceQuarterly  PriceModelGroupedWithProratedMinimumPriceCadence = "quarterly"
	PriceModelGroupedWithProratedMinimumPriceCadenceSemiAnnual PriceModelGroupedWithProratedMinimumPriceCadence = "semi_annual"
	PriceModelGroupedWithProratedMinimumPriceCadenceAnnual     PriceModelGroupedWithProratedMinimumPriceCadence = "annual"
	PriceModelGroupedWithProratedMinimumPriceCadenceCustom     PriceModelGroupedWithProratedMinimumPriceCadence = "custom"
)

func (r PriceModelGroupedWithProratedMinimumPriceCadence) IsKnown() bool {
	switch r {
	case PriceModelGroupedWithProratedMinimumPriceCadenceOneTime, PriceModelGroupedWithProratedMinimumPriceCadenceMonthly, PriceModelGroupedWithProratedMinimumPriceCadenceQuarterly, PriceModelGroupedWithProratedMinimumPriceCadenceSemiAnnual, PriceModelGroupedWithProratedMinimumPriceCadenceAnnual, PriceModelGroupedWithProratedMinimumPriceCadenceCustom:
		return true
	}
	return false
}

type PriceModelGroupedWithProratedMinimumPriceModelType string

const (
	PriceModelGroupedWithProratedMinimumPriceModelTypeGroupedWithProratedMinimum PriceModelGroupedWithProratedMinimumPriceModelType = "grouped_with_prorated_minimum"
)

func (r PriceModelGroupedWithProratedMinimumPriceModelType) IsKnown() bool {
	switch r {
	case PriceModelGroupedWithProratedMinimumPriceModelTypeGroupedWithProratedMinimum:
		return true
	}
	return false
}

type PriceModelGroupedWithProratedMinimumPricePriceType string

const (
	PriceModelGroupedWithProratedMinimumPricePriceTypeUsagePrice PriceModelGroupedWithProratedMinimumPricePriceType = "usage_price"
	PriceModelGroupedWithProratedMinimumPricePriceTypeFixedPrice PriceModelGroupedWithProratedMinimumPricePriceType = "fixed_price"
)

func (r PriceModelGroupedWithProratedMinimumPricePriceType) IsKnown() bool {
	switch r {
	case PriceModelGroupedWithProratedMinimumPricePriceTypeUsagePrice, PriceModelGroupedWithProratedMinimumPricePriceTypeFixedPrice:
		return true
	}
	return false
}

type PriceModelGroupedWithMeteredMinimumPrice struct {
	ID                              string                                          `json:"id,required"`
	BillableMetric                  BillableMetricTinyModel                         `json:"billable_metric,required,nullable"`
	BillingCycleConfiguration       BillingCycleConfigurationModel                  `json:"billing_cycle_configuration,required"`
	Cadence                         PriceModelGroupedWithMeteredMinimumPriceCadence `json:"cadence,required"`
	ConversionRate                  float64                                         `json:"conversion_rate,required,nullable"`
	CreatedAt                       time.Time                                       `json:"created_at,required" format:"date-time"`
	CreditAllocation                AllocationModel                                 `json:"credit_allocation,required,nullable"`
	Currency                        string                                          `json:"currency,required"`
	Discount                        Discount                                        `json:"discount,required,nullable"`
	ExternalPriceID                 string                                          `json:"external_price_id,required,nullable"`
	FixedPriceQuantity              float64                                         `json:"fixed_price_quantity,required,nullable"`
	GroupedWithMeteredMinimumConfig CustomRatingFunctionConfigModel                 `json:"grouped_with_metered_minimum_config,required"`
	InvoicingCycleConfiguration     BillingCycleConfigurationModel                  `json:"invoicing_cycle_configuration,required,nullable"`
	Item                            ItemSlimModel                                   `json:"item,required"`
	Maximum                         MaximumModel                                    `json:"maximum,required,nullable"`
	MaximumAmount                   string                                          `json:"maximum_amount,required,nullable"`
	// User specified key-value pairs for the resource. If not present, this defaults
	// to an empty dictionary. Individual keys can be removed by setting the value to
	// `null`, and the entire metadata mapping can be cleared by setting `metadata` to
	// `null`.
	Metadata                      map[string]string                                 `json:"metadata,required"`
	Minimum                       MinimumModel                                      `json:"minimum,required,nullable"`
	MinimumAmount                 string                                            `json:"minimum_amount,required,nullable"`
	ModelType                     PriceModelGroupedWithMeteredMinimumPriceModelType `json:"model_type,required"`
	Name                          string                                            `json:"name,required"`
	PlanPhaseOrder                int64                                             `json:"plan_phase_order,required,nullable"`
	PriceType                     PriceModelGroupedWithMeteredMinimumPricePriceType `json:"price_type,required"`
	DimensionalPriceConfiguration DimensionalPriceConfigurationModel                `json:"dimensional_price_configuration,nullable"`
	JSON                          priceModelGroupedWithMeteredMinimumPriceJSON      `json:"-"`
}

// priceModelGroupedWithMeteredMinimumPriceJSON contains the JSON metadata for the
// struct [PriceModelGroupedWithMeteredMinimumPrice]
type priceModelGroupedWithMeteredMinimumPriceJSON struct {
	ID                              apijson.Field
	BillableMetric                  apijson.Field
	BillingCycleConfiguration       apijson.Field
	Cadence                         apijson.Field
	ConversionRate                  apijson.Field
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
	DimensionalPriceConfiguration   apijson.Field
	raw                             string
	ExtraFields                     map[string]apijson.Field
}

func (r *PriceModelGroupedWithMeteredMinimumPrice) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r priceModelGroupedWithMeteredMinimumPriceJSON) RawJSON() string {
	return r.raw
}

func (r PriceModelGroupedWithMeteredMinimumPrice) implementsPriceModel() {}

type PriceModelGroupedWithMeteredMinimumPriceCadence string

const (
	PriceModelGroupedWithMeteredMinimumPriceCadenceOneTime    PriceModelGroupedWithMeteredMinimumPriceCadence = "one_time"
	PriceModelGroupedWithMeteredMinimumPriceCadenceMonthly    PriceModelGroupedWithMeteredMinimumPriceCadence = "monthly"
	PriceModelGroupedWithMeteredMinimumPriceCadenceQuarterly  PriceModelGroupedWithMeteredMinimumPriceCadence = "quarterly"
	PriceModelGroupedWithMeteredMinimumPriceCadenceSemiAnnual PriceModelGroupedWithMeteredMinimumPriceCadence = "semi_annual"
	PriceModelGroupedWithMeteredMinimumPriceCadenceAnnual     PriceModelGroupedWithMeteredMinimumPriceCadence = "annual"
	PriceModelGroupedWithMeteredMinimumPriceCadenceCustom     PriceModelGroupedWithMeteredMinimumPriceCadence = "custom"
)

func (r PriceModelGroupedWithMeteredMinimumPriceCadence) IsKnown() bool {
	switch r {
	case PriceModelGroupedWithMeteredMinimumPriceCadenceOneTime, PriceModelGroupedWithMeteredMinimumPriceCadenceMonthly, PriceModelGroupedWithMeteredMinimumPriceCadenceQuarterly, PriceModelGroupedWithMeteredMinimumPriceCadenceSemiAnnual, PriceModelGroupedWithMeteredMinimumPriceCadenceAnnual, PriceModelGroupedWithMeteredMinimumPriceCadenceCustom:
		return true
	}
	return false
}

type PriceModelGroupedWithMeteredMinimumPriceModelType string

const (
	PriceModelGroupedWithMeteredMinimumPriceModelTypeGroupedWithMeteredMinimum PriceModelGroupedWithMeteredMinimumPriceModelType = "grouped_with_metered_minimum"
)

func (r PriceModelGroupedWithMeteredMinimumPriceModelType) IsKnown() bool {
	switch r {
	case PriceModelGroupedWithMeteredMinimumPriceModelTypeGroupedWithMeteredMinimum:
		return true
	}
	return false
}

type PriceModelGroupedWithMeteredMinimumPricePriceType string

const (
	PriceModelGroupedWithMeteredMinimumPricePriceTypeUsagePrice PriceModelGroupedWithMeteredMinimumPricePriceType = "usage_price"
	PriceModelGroupedWithMeteredMinimumPricePriceTypeFixedPrice PriceModelGroupedWithMeteredMinimumPricePriceType = "fixed_price"
)

func (r PriceModelGroupedWithMeteredMinimumPricePriceType) IsKnown() bool {
	switch r {
	case PriceModelGroupedWithMeteredMinimumPricePriceTypeUsagePrice, PriceModelGroupedWithMeteredMinimumPricePriceTypeFixedPrice:
		return true
	}
	return false
}

type PriceModelMatrixWithDisplayNamePrice struct {
	ID                          string                                      `json:"id,required"`
	BillableMetric              BillableMetricTinyModel                     `json:"billable_metric,required,nullable"`
	BillingCycleConfiguration   BillingCycleConfigurationModel              `json:"billing_cycle_configuration,required"`
	Cadence                     PriceModelMatrixWithDisplayNamePriceCadence `json:"cadence,required"`
	ConversionRate              float64                                     `json:"conversion_rate,required,nullable"`
	CreatedAt                   time.Time                                   `json:"created_at,required" format:"date-time"`
	CreditAllocation            AllocationModel                             `json:"credit_allocation,required,nullable"`
	Currency                    string                                      `json:"currency,required"`
	Discount                    Discount                                    `json:"discount,required,nullable"`
	ExternalPriceID             string                                      `json:"external_price_id,required,nullable"`
	FixedPriceQuantity          float64                                     `json:"fixed_price_quantity,required,nullable"`
	InvoicingCycleConfiguration BillingCycleConfigurationModel              `json:"invoicing_cycle_configuration,required,nullable"`
	Item                        ItemSlimModel                               `json:"item,required"`
	MatrixWithDisplayNameConfig CustomRatingFunctionConfigModel             `json:"matrix_with_display_name_config,required"`
	Maximum                     MaximumModel                                `json:"maximum,required,nullable"`
	MaximumAmount               string                                      `json:"maximum_amount,required,nullable"`
	// User specified key-value pairs for the resource. If not present, this defaults
	// to an empty dictionary. Individual keys can be removed by setting the value to
	// `null`, and the entire metadata mapping can be cleared by setting `metadata` to
	// `null`.
	Metadata                      map[string]string                             `json:"metadata,required"`
	Minimum                       MinimumModel                                  `json:"minimum,required,nullable"`
	MinimumAmount                 string                                        `json:"minimum_amount,required,nullable"`
	ModelType                     PriceModelMatrixWithDisplayNamePriceModelType `json:"model_type,required"`
	Name                          string                                        `json:"name,required"`
	PlanPhaseOrder                int64                                         `json:"plan_phase_order,required,nullable"`
	PriceType                     PriceModelMatrixWithDisplayNamePricePriceType `json:"price_type,required"`
	DimensionalPriceConfiguration DimensionalPriceConfigurationModel            `json:"dimensional_price_configuration,nullable"`
	JSON                          priceModelMatrixWithDisplayNamePriceJSON      `json:"-"`
}

// priceModelMatrixWithDisplayNamePriceJSON contains the JSON metadata for the
// struct [PriceModelMatrixWithDisplayNamePrice]
type priceModelMatrixWithDisplayNamePriceJSON struct {
	ID                            apijson.Field
	BillableMetric                apijson.Field
	BillingCycleConfiguration     apijson.Field
	Cadence                       apijson.Field
	ConversionRate                apijson.Field
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
	DimensionalPriceConfiguration apijson.Field
	raw                           string
	ExtraFields                   map[string]apijson.Field
}

func (r *PriceModelMatrixWithDisplayNamePrice) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r priceModelMatrixWithDisplayNamePriceJSON) RawJSON() string {
	return r.raw
}

func (r PriceModelMatrixWithDisplayNamePrice) implementsPriceModel() {}

type PriceModelMatrixWithDisplayNamePriceCadence string

const (
	PriceModelMatrixWithDisplayNamePriceCadenceOneTime    PriceModelMatrixWithDisplayNamePriceCadence = "one_time"
	PriceModelMatrixWithDisplayNamePriceCadenceMonthly    PriceModelMatrixWithDisplayNamePriceCadence = "monthly"
	PriceModelMatrixWithDisplayNamePriceCadenceQuarterly  PriceModelMatrixWithDisplayNamePriceCadence = "quarterly"
	PriceModelMatrixWithDisplayNamePriceCadenceSemiAnnual PriceModelMatrixWithDisplayNamePriceCadence = "semi_annual"
	PriceModelMatrixWithDisplayNamePriceCadenceAnnual     PriceModelMatrixWithDisplayNamePriceCadence = "annual"
	PriceModelMatrixWithDisplayNamePriceCadenceCustom     PriceModelMatrixWithDisplayNamePriceCadence = "custom"
)

func (r PriceModelMatrixWithDisplayNamePriceCadence) IsKnown() bool {
	switch r {
	case PriceModelMatrixWithDisplayNamePriceCadenceOneTime, PriceModelMatrixWithDisplayNamePriceCadenceMonthly, PriceModelMatrixWithDisplayNamePriceCadenceQuarterly, PriceModelMatrixWithDisplayNamePriceCadenceSemiAnnual, PriceModelMatrixWithDisplayNamePriceCadenceAnnual, PriceModelMatrixWithDisplayNamePriceCadenceCustom:
		return true
	}
	return false
}

type PriceModelMatrixWithDisplayNamePriceModelType string

const (
	PriceModelMatrixWithDisplayNamePriceModelTypeMatrixWithDisplayName PriceModelMatrixWithDisplayNamePriceModelType = "matrix_with_display_name"
)

func (r PriceModelMatrixWithDisplayNamePriceModelType) IsKnown() bool {
	switch r {
	case PriceModelMatrixWithDisplayNamePriceModelTypeMatrixWithDisplayName:
		return true
	}
	return false
}

type PriceModelMatrixWithDisplayNamePricePriceType string

const (
	PriceModelMatrixWithDisplayNamePricePriceTypeUsagePrice PriceModelMatrixWithDisplayNamePricePriceType = "usage_price"
	PriceModelMatrixWithDisplayNamePricePriceTypeFixedPrice PriceModelMatrixWithDisplayNamePricePriceType = "fixed_price"
)

func (r PriceModelMatrixWithDisplayNamePricePriceType) IsKnown() bool {
	switch r {
	case PriceModelMatrixWithDisplayNamePricePriceTypeUsagePrice, PriceModelMatrixWithDisplayNamePricePriceTypeFixedPrice:
		return true
	}
	return false
}

type PriceModelBulkWithProrationPrice struct {
	ID                          string                                  `json:"id,required"`
	BillableMetric              BillableMetricTinyModel                 `json:"billable_metric,required,nullable"`
	BillingCycleConfiguration   BillingCycleConfigurationModel          `json:"billing_cycle_configuration,required"`
	BulkWithProrationConfig     CustomRatingFunctionConfigModel         `json:"bulk_with_proration_config,required"`
	Cadence                     PriceModelBulkWithProrationPriceCadence `json:"cadence,required"`
	ConversionRate              float64                                 `json:"conversion_rate,required,nullable"`
	CreatedAt                   time.Time                               `json:"created_at,required" format:"date-time"`
	CreditAllocation            AllocationModel                         `json:"credit_allocation,required,nullable"`
	Currency                    string                                  `json:"currency,required"`
	Discount                    Discount                                `json:"discount,required,nullable"`
	ExternalPriceID             string                                  `json:"external_price_id,required,nullable"`
	FixedPriceQuantity          float64                                 `json:"fixed_price_quantity,required,nullable"`
	InvoicingCycleConfiguration BillingCycleConfigurationModel          `json:"invoicing_cycle_configuration,required,nullable"`
	Item                        ItemSlimModel                           `json:"item,required"`
	Maximum                     MaximumModel                            `json:"maximum,required,nullable"`
	MaximumAmount               string                                  `json:"maximum_amount,required,nullable"`
	// User specified key-value pairs for the resource. If not present, this defaults
	// to an empty dictionary. Individual keys can be removed by setting the value to
	// `null`, and the entire metadata mapping can be cleared by setting `metadata` to
	// `null`.
	Metadata                      map[string]string                         `json:"metadata,required"`
	Minimum                       MinimumModel                              `json:"minimum,required,nullable"`
	MinimumAmount                 string                                    `json:"minimum_amount,required,nullable"`
	ModelType                     PriceModelBulkWithProrationPriceModelType `json:"model_type,required"`
	Name                          string                                    `json:"name,required"`
	PlanPhaseOrder                int64                                     `json:"plan_phase_order,required,nullable"`
	PriceType                     PriceModelBulkWithProrationPricePriceType `json:"price_type,required"`
	DimensionalPriceConfiguration DimensionalPriceConfigurationModel        `json:"dimensional_price_configuration,nullable"`
	JSON                          priceModelBulkWithProrationPriceJSON      `json:"-"`
}

// priceModelBulkWithProrationPriceJSON contains the JSON metadata for the struct
// [PriceModelBulkWithProrationPrice]
type priceModelBulkWithProrationPriceJSON struct {
	ID                            apijson.Field
	BillableMetric                apijson.Field
	BillingCycleConfiguration     apijson.Field
	BulkWithProrationConfig       apijson.Field
	Cadence                       apijson.Field
	ConversionRate                apijson.Field
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
	DimensionalPriceConfiguration apijson.Field
	raw                           string
	ExtraFields                   map[string]apijson.Field
}

func (r *PriceModelBulkWithProrationPrice) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r priceModelBulkWithProrationPriceJSON) RawJSON() string {
	return r.raw
}

func (r PriceModelBulkWithProrationPrice) implementsPriceModel() {}

type PriceModelBulkWithProrationPriceCadence string

const (
	PriceModelBulkWithProrationPriceCadenceOneTime    PriceModelBulkWithProrationPriceCadence = "one_time"
	PriceModelBulkWithProrationPriceCadenceMonthly    PriceModelBulkWithProrationPriceCadence = "monthly"
	PriceModelBulkWithProrationPriceCadenceQuarterly  PriceModelBulkWithProrationPriceCadence = "quarterly"
	PriceModelBulkWithProrationPriceCadenceSemiAnnual PriceModelBulkWithProrationPriceCadence = "semi_annual"
	PriceModelBulkWithProrationPriceCadenceAnnual     PriceModelBulkWithProrationPriceCadence = "annual"
	PriceModelBulkWithProrationPriceCadenceCustom     PriceModelBulkWithProrationPriceCadence = "custom"
)

func (r PriceModelBulkWithProrationPriceCadence) IsKnown() bool {
	switch r {
	case PriceModelBulkWithProrationPriceCadenceOneTime, PriceModelBulkWithProrationPriceCadenceMonthly, PriceModelBulkWithProrationPriceCadenceQuarterly, PriceModelBulkWithProrationPriceCadenceSemiAnnual, PriceModelBulkWithProrationPriceCadenceAnnual, PriceModelBulkWithProrationPriceCadenceCustom:
		return true
	}
	return false
}

type PriceModelBulkWithProrationPriceModelType string

const (
	PriceModelBulkWithProrationPriceModelTypeBulkWithProration PriceModelBulkWithProrationPriceModelType = "bulk_with_proration"
)

func (r PriceModelBulkWithProrationPriceModelType) IsKnown() bool {
	switch r {
	case PriceModelBulkWithProrationPriceModelTypeBulkWithProration:
		return true
	}
	return false
}

type PriceModelBulkWithProrationPricePriceType string

const (
	PriceModelBulkWithProrationPricePriceTypeUsagePrice PriceModelBulkWithProrationPricePriceType = "usage_price"
	PriceModelBulkWithProrationPricePriceTypeFixedPrice PriceModelBulkWithProrationPricePriceType = "fixed_price"
)

func (r PriceModelBulkWithProrationPricePriceType) IsKnown() bool {
	switch r {
	case PriceModelBulkWithProrationPricePriceTypeUsagePrice, PriceModelBulkWithProrationPricePriceTypeFixedPrice:
		return true
	}
	return false
}

type PriceModelGroupedTieredPackagePrice struct {
	ID                          string                                     `json:"id,required"`
	BillableMetric              BillableMetricTinyModel                    `json:"billable_metric,required,nullable"`
	BillingCycleConfiguration   BillingCycleConfigurationModel             `json:"billing_cycle_configuration,required"`
	Cadence                     PriceModelGroupedTieredPackagePriceCadence `json:"cadence,required"`
	ConversionRate              float64                                    `json:"conversion_rate,required,nullable"`
	CreatedAt                   time.Time                                  `json:"created_at,required" format:"date-time"`
	CreditAllocation            AllocationModel                            `json:"credit_allocation,required,nullable"`
	Currency                    string                                     `json:"currency,required"`
	Discount                    Discount                                   `json:"discount,required,nullable"`
	ExternalPriceID             string                                     `json:"external_price_id,required,nullable"`
	FixedPriceQuantity          float64                                    `json:"fixed_price_quantity,required,nullable"`
	GroupedTieredPackageConfig  CustomRatingFunctionConfigModel            `json:"grouped_tiered_package_config,required"`
	InvoicingCycleConfiguration BillingCycleConfigurationModel             `json:"invoicing_cycle_configuration,required,nullable"`
	Item                        ItemSlimModel                              `json:"item,required"`
	Maximum                     MaximumModel                               `json:"maximum,required,nullable"`
	MaximumAmount               string                                     `json:"maximum_amount,required,nullable"`
	// User specified key-value pairs for the resource. If not present, this defaults
	// to an empty dictionary. Individual keys can be removed by setting the value to
	// `null`, and the entire metadata mapping can be cleared by setting `metadata` to
	// `null`.
	Metadata                      map[string]string                            `json:"metadata,required"`
	Minimum                       MinimumModel                                 `json:"minimum,required,nullable"`
	MinimumAmount                 string                                       `json:"minimum_amount,required,nullable"`
	ModelType                     PriceModelGroupedTieredPackagePriceModelType `json:"model_type,required"`
	Name                          string                                       `json:"name,required"`
	PlanPhaseOrder                int64                                        `json:"plan_phase_order,required,nullable"`
	PriceType                     PriceModelGroupedTieredPackagePricePriceType `json:"price_type,required"`
	DimensionalPriceConfiguration DimensionalPriceConfigurationModel           `json:"dimensional_price_configuration,nullable"`
	JSON                          priceModelGroupedTieredPackagePriceJSON      `json:"-"`
}

// priceModelGroupedTieredPackagePriceJSON contains the JSON metadata for the
// struct [PriceModelGroupedTieredPackagePrice]
type priceModelGroupedTieredPackagePriceJSON struct {
	ID                            apijson.Field
	BillableMetric                apijson.Field
	BillingCycleConfiguration     apijson.Field
	Cadence                       apijson.Field
	ConversionRate                apijson.Field
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
	DimensionalPriceConfiguration apijson.Field
	raw                           string
	ExtraFields                   map[string]apijson.Field
}

func (r *PriceModelGroupedTieredPackagePrice) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r priceModelGroupedTieredPackagePriceJSON) RawJSON() string {
	return r.raw
}

func (r PriceModelGroupedTieredPackagePrice) implementsPriceModel() {}

type PriceModelGroupedTieredPackagePriceCadence string

const (
	PriceModelGroupedTieredPackagePriceCadenceOneTime    PriceModelGroupedTieredPackagePriceCadence = "one_time"
	PriceModelGroupedTieredPackagePriceCadenceMonthly    PriceModelGroupedTieredPackagePriceCadence = "monthly"
	PriceModelGroupedTieredPackagePriceCadenceQuarterly  PriceModelGroupedTieredPackagePriceCadence = "quarterly"
	PriceModelGroupedTieredPackagePriceCadenceSemiAnnual PriceModelGroupedTieredPackagePriceCadence = "semi_annual"
	PriceModelGroupedTieredPackagePriceCadenceAnnual     PriceModelGroupedTieredPackagePriceCadence = "annual"
	PriceModelGroupedTieredPackagePriceCadenceCustom     PriceModelGroupedTieredPackagePriceCadence = "custom"
)

func (r PriceModelGroupedTieredPackagePriceCadence) IsKnown() bool {
	switch r {
	case PriceModelGroupedTieredPackagePriceCadenceOneTime, PriceModelGroupedTieredPackagePriceCadenceMonthly, PriceModelGroupedTieredPackagePriceCadenceQuarterly, PriceModelGroupedTieredPackagePriceCadenceSemiAnnual, PriceModelGroupedTieredPackagePriceCadenceAnnual, PriceModelGroupedTieredPackagePriceCadenceCustom:
		return true
	}
	return false
}

type PriceModelGroupedTieredPackagePriceModelType string

const (
	PriceModelGroupedTieredPackagePriceModelTypeGroupedTieredPackage PriceModelGroupedTieredPackagePriceModelType = "grouped_tiered_package"
)

func (r PriceModelGroupedTieredPackagePriceModelType) IsKnown() bool {
	switch r {
	case PriceModelGroupedTieredPackagePriceModelTypeGroupedTieredPackage:
		return true
	}
	return false
}

type PriceModelGroupedTieredPackagePricePriceType string

const (
	PriceModelGroupedTieredPackagePricePriceTypeUsagePrice PriceModelGroupedTieredPackagePricePriceType = "usage_price"
	PriceModelGroupedTieredPackagePricePriceTypeFixedPrice PriceModelGroupedTieredPackagePricePriceType = "fixed_price"
)

func (r PriceModelGroupedTieredPackagePricePriceType) IsKnown() bool {
	switch r {
	case PriceModelGroupedTieredPackagePricePriceTypeUsagePrice, PriceModelGroupedTieredPackagePricePriceTypeFixedPrice:
		return true
	}
	return false
}

type PriceModelMaxGroupTieredPackagePrice struct {
	ID                          string                                      `json:"id,required"`
	BillableMetric              BillableMetricTinyModel                     `json:"billable_metric,required,nullable"`
	BillingCycleConfiguration   BillingCycleConfigurationModel              `json:"billing_cycle_configuration,required"`
	Cadence                     PriceModelMaxGroupTieredPackagePriceCadence `json:"cadence,required"`
	ConversionRate              float64                                     `json:"conversion_rate,required,nullable"`
	CreatedAt                   time.Time                                   `json:"created_at,required" format:"date-time"`
	CreditAllocation            AllocationModel                             `json:"credit_allocation,required,nullable"`
	Currency                    string                                      `json:"currency,required"`
	Discount                    Discount                                    `json:"discount,required,nullable"`
	ExternalPriceID             string                                      `json:"external_price_id,required,nullable"`
	FixedPriceQuantity          float64                                     `json:"fixed_price_quantity,required,nullable"`
	InvoicingCycleConfiguration BillingCycleConfigurationModel              `json:"invoicing_cycle_configuration,required,nullable"`
	Item                        ItemSlimModel                               `json:"item,required"`
	MaxGroupTieredPackageConfig CustomRatingFunctionConfigModel             `json:"max_group_tiered_package_config,required"`
	Maximum                     MaximumModel                                `json:"maximum,required,nullable"`
	MaximumAmount               string                                      `json:"maximum_amount,required,nullable"`
	// User specified key-value pairs for the resource. If not present, this defaults
	// to an empty dictionary. Individual keys can be removed by setting the value to
	// `null`, and the entire metadata mapping can be cleared by setting `metadata` to
	// `null`.
	Metadata                      map[string]string                             `json:"metadata,required"`
	Minimum                       MinimumModel                                  `json:"minimum,required,nullable"`
	MinimumAmount                 string                                        `json:"minimum_amount,required,nullable"`
	ModelType                     PriceModelMaxGroupTieredPackagePriceModelType `json:"model_type,required"`
	Name                          string                                        `json:"name,required"`
	PlanPhaseOrder                int64                                         `json:"plan_phase_order,required,nullable"`
	PriceType                     PriceModelMaxGroupTieredPackagePricePriceType `json:"price_type,required"`
	DimensionalPriceConfiguration DimensionalPriceConfigurationModel            `json:"dimensional_price_configuration,nullable"`
	JSON                          priceModelMaxGroupTieredPackagePriceJSON      `json:"-"`
}

// priceModelMaxGroupTieredPackagePriceJSON contains the JSON metadata for the
// struct [PriceModelMaxGroupTieredPackagePrice]
type priceModelMaxGroupTieredPackagePriceJSON struct {
	ID                            apijson.Field
	BillableMetric                apijson.Field
	BillingCycleConfiguration     apijson.Field
	Cadence                       apijson.Field
	ConversionRate                apijson.Field
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
	DimensionalPriceConfiguration apijson.Field
	raw                           string
	ExtraFields                   map[string]apijson.Field
}

func (r *PriceModelMaxGroupTieredPackagePrice) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r priceModelMaxGroupTieredPackagePriceJSON) RawJSON() string {
	return r.raw
}

func (r PriceModelMaxGroupTieredPackagePrice) implementsPriceModel() {}

type PriceModelMaxGroupTieredPackagePriceCadence string

const (
	PriceModelMaxGroupTieredPackagePriceCadenceOneTime    PriceModelMaxGroupTieredPackagePriceCadence = "one_time"
	PriceModelMaxGroupTieredPackagePriceCadenceMonthly    PriceModelMaxGroupTieredPackagePriceCadence = "monthly"
	PriceModelMaxGroupTieredPackagePriceCadenceQuarterly  PriceModelMaxGroupTieredPackagePriceCadence = "quarterly"
	PriceModelMaxGroupTieredPackagePriceCadenceSemiAnnual PriceModelMaxGroupTieredPackagePriceCadence = "semi_annual"
	PriceModelMaxGroupTieredPackagePriceCadenceAnnual     PriceModelMaxGroupTieredPackagePriceCadence = "annual"
	PriceModelMaxGroupTieredPackagePriceCadenceCustom     PriceModelMaxGroupTieredPackagePriceCadence = "custom"
)

func (r PriceModelMaxGroupTieredPackagePriceCadence) IsKnown() bool {
	switch r {
	case PriceModelMaxGroupTieredPackagePriceCadenceOneTime, PriceModelMaxGroupTieredPackagePriceCadenceMonthly, PriceModelMaxGroupTieredPackagePriceCadenceQuarterly, PriceModelMaxGroupTieredPackagePriceCadenceSemiAnnual, PriceModelMaxGroupTieredPackagePriceCadenceAnnual, PriceModelMaxGroupTieredPackagePriceCadenceCustom:
		return true
	}
	return false
}

type PriceModelMaxGroupTieredPackagePriceModelType string

const (
	PriceModelMaxGroupTieredPackagePriceModelTypeMaxGroupTieredPackage PriceModelMaxGroupTieredPackagePriceModelType = "max_group_tiered_package"
)

func (r PriceModelMaxGroupTieredPackagePriceModelType) IsKnown() bool {
	switch r {
	case PriceModelMaxGroupTieredPackagePriceModelTypeMaxGroupTieredPackage:
		return true
	}
	return false
}

type PriceModelMaxGroupTieredPackagePricePriceType string

const (
	PriceModelMaxGroupTieredPackagePricePriceTypeUsagePrice PriceModelMaxGroupTieredPackagePricePriceType = "usage_price"
	PriceModelMaxGroupTieredPackagePricePriceTypeFixedPrice PriceModelMaxGroupTieredPackagePricePriceType = "fixed_price"
)

func (r PriceModelMaxGroupTieredPackagePricePriceType) IsKnown() bool {
	switch r {
	case PriceModelMaxGroupTieredPackagePricePriceTypeUsagePrice, PriceModelMaxGroupTieredPackagePricePriceTypeFixedPrice:
		return true
	}
	return false
}

type PriceModelScalableMatrixWithUnitPricingPrice struct {
	ID                          string                                              `json:"id,required"`
	BillableMetric              BillableMetricTinyModel                             `json:"billable_metric,required,nullable"`
	BillingCycleConfiguration   BillingCycleConfigurationModel                      `json:"billing_cycle_configuration,required"`
	Cadence                     PriceModelScalableMatrixWithUnitPricingPriceCadence `json:"cadence,required"`
	ConversionRate              float64                                             `json:"conversion_rate,required,nullable"`
	CreatedAt                   time.Time                                           `json:"created_at,required" format:"date-time"`
	CreditAllocation            AllocationModel                                     `json:"credit_allocation,required,nullable"`
	Currency                    string                                              `json:"currency,required"`
	Discount                    Discount                                            `json:"discount,required,nullable"`
	ExternalPriceID             string                                              `json:"external_price_id,required,nullable"`
	FixedPriceQuantity          float64                                             `json:"fixed_price_quantity,required,nullable"`
	InvoicingCycleConfiguration BillingCycleConfigurationModel                      `json:"invoicing_cycle_configuration,required,nullable"`
	Item                        ItemSlimModel                                       `json:"item,required"`
	Maximum                     MaximumModel                                        `json:"maximum,required,nullable"`
	MaximumAmount               string                                              `json:"maximum_amount,required,nullable"`
	// User specified key-value pairs for the resource. If not present, this defaults
	// to an empty dictionary. Individual keys can be removed by setting the value to
	// `null`, and the entire metadata mapping can be cleared by setting `metadata` to
	// `null`.
	Metadata                            map[string]string                                     `json:"metadata,required"`
	Minimum                             MinimumModel                                          `json:"minimum,required,nullable"`
	MinimumAmount                       string                                                `json:"minimum_amount,required,nullable"`
	ModelType                           PriceModelScalableMatrixWithUnitPricingPriceModelType `json:"model_type,required"`
	Name                                string                                                `json:"name,required"`
	PlanPhaseOrder                      int64                                                 `json:"plan_phase_order,required,nullable"`
	PriceType                           PriceModelScalableMatrixWithUnitPricingPricePriceType `json:"price_type,required"`
	ScalableMatrixWithUnitPricingConfig CustomRatingFunctionConfigModel                       `json:"scalable_matrix_with_unit_pricing_config,required"`
	DimensionalPriceConfiguration       DimensionalPriceConfigurationModel                    `json:"dimensional_price_configuration,nullable"`
	JSON                                priceModelScalableMatrixWithUnitPricingPriceJSON      `json:"-"`
}

// priceModelScalableMatrixWithUnitPricingPriceJSON contains the JSON metadata for
// the struct [PriceModelScalableMatrixWithUnitPricingPrice]
type priceModelScalableMatrixWithUnitPricingPriceJSON struct {
	ID                                  apijson.Field
	BillableMetric                      apijson.Field
	BillingCycleConfiguration           apijson.Field
	Cadence                             apijson.Field
	ConversionRate                      apijson.Field
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
	ScalableMatrixWithUnitPricingConfig apijson.Field
	DimensionalPriceConfiguration       apijson.Field
	raw                                 string
	ExtraFields                         map[string]apijson.Field
}

func (r *PriceModelScalableMatrixWithUnitPricingPrice) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r priceModelScalableMatrixWithUnitPricingPriceJSON) RawJSON() string {
	return r.raw
}

func (r PriceModelScalableMatrixWithUnitPricingPrice) implementsPriceModel() {}

type PriceModelScalableMatrixWithUnitPricingPriceCadence string

const (
	PriceModelScalableMatrixWithUnitPricingPriceCadenceOneTime    PriceModelScalableMatrixWithUnitPricingPriceCadence = "one_time"
	PriceModelScalableMatrixWithUnitPricingPriceCadenceMonthly    PriceModelScalableMatrixWithUnitPricingPriceCadence = "monthly"
	PriceModelScalableMatrixWithUnitPricingPriceCadenceQuarterly  PriceModelScalableMatrixWithUnitPricingPriceCadence = "quarterly"
	PriceModelScalableMatrixWithUnitPricingPriceCadenceSemiAnnual PriceModelScalableMatrixWithUnitPricingPriceCadence = "semi_annual"
	PriceModelScalableMatrixWithUnitPricingPriceCadenceAnnual     PriceModelScalableMatrixWithUnitPricingPriceCadence = "annual"
	PriceModelScalableMatrixWithUnitPricingPriceCadenceCustom     PriceModelScalableMatrixWithUnitPricingPriceCadence = "custom"
)

func (r PriceModelScalableMatrixWithUnitPricingPriceCadence) IsKnown() bool {
	switch r {
	case PriceModelScalableMatrixWithUnitPricingPriceCadenceOneTime, PriceModelScalableMatrixWithUnitPricingPriceCadenceMonthly, PriceModelScalableMatrixWithUnitPricingPriceCadenceQuarterly, PriceModelScalableMatrixWithUnitPricingPriceCadenceSemiAnnual, PriceModelScalableMatrixWithUnitPricingPriceCadenceAnnual, PriceModelScalableMatrixWithUnitPricingPriceCadenceCustom:
		return true
	}
	return false
}

type PriceModelScalableMatrixWithUnitPricingPriceModelType string

const (
	PriceModelScalableMatrixWithUnitPricingPriceModelTypeScalableMatrixWithUnitPricing PriceModelScalableMatrixWithUnitPricingPriceModelType = "scalable_matrix_with_unit_pricing"
)

func (r PriceModelScalableMatrixWithUnitPricingPriceModelType) IsKnown() bool {
	switch r {
	case PriceModelScalableMatrixWithUnitPricingPriceModelTypeScalableMatrixWithUnitPricing:
		return true
	}
	return false
}

type PriceModelScalableMatrixWithUnitPricingPricePriceType string

const (
	PriceModelScalableMatrixWithUnitPricingPricePriceTypeUsagePrice PriceModelScalableMatrixWithUnitPricingPricePriceType = "usage_price"
	PriceModelScalableMatrixWithUnitPricingPricePriceTypeFixedPrice PriceModelScalableMatrixWithUnitPricingPricePriceType = "fixed_price"
)

func (r PriceModelScalableMatrixWithUnitPricingPricePriceType) IsKnown() bool {
	switch r {
	case PriceModelScalableMatrixWithUnitPricingPricePriceTypeUsagePrice, PriceModelScalableMatrixWithUnitPricingPricePriceTypeFixedPrice:
		return true
	}
	return false
}

type PriceModelScalableMatrixWithTieredPricingPrice struct {
	ID                          string                                                `json:"id,required"`
	BillableMetric              BillableMetricTinyModel                               `json:"billable_metric,required,nullable"`
	BillingCycleConfiguration   BillingCycleConfigurationModel                        `json:"billing_cycle_configuration,required"`
	Cadence                     PriceModelScalableMatrixWithTieredPricingPriceCadence `json:"cadence,required"`
	ConversionRate              float64                                               `json:"conversion_rate,required,nullable"`
	CreatedAt                   time.Time                                             `json:"created_at,required" format:"date-time"`
	CreditAllocation            AllocationModel                                       `json:"credit_allocation,required,nullable"`
	Currency                    string                                                `json:"currency,required"`
	Discount                    Discount                                              `json:"discount,required,nullable"`
	ExternalPriceID             string                                                `json:"external_price_id,required,nullable"`
	FixedPriceQuantity          float64                                               `json:"fixed_price_quantity,required,nullable"`
	InvoicingCycleConfiguration BillingCycleConfigurationModel                        `json:"invoicing_cycle_configuration,required,nullable"`
	Item                        ItemSlimModel                                         `json:"item,required"`
	Maximum                     MaximumModel                                          `json:"maximum,required,nullable"`
	MaximumAmount               string                                                `json:"maximum_amount,required,nullable"`
	// User specified key-value pairs for the resource. If not present, this defaults
	// to an empty dictionary. Individual keys can be removed by setting the value to
	// `null`, and the entire metadata mapping can be cleared by setting `metadata` to
	// `null`.
	Metadata                              map[string]string                                       `json:"metadata,required"`
	Minimum                               MinimumModel                                            `json:"minimum,required,nullable"`
	MinimumAmount                         string                                                  `json:"minimum_amount,required,nullable"`
	ModelType                             PriceModelScalableMatrixWithTieredPricingPriceModelType `json:"model_type,required"`
	Name                                  string                                                  `json:"name,required"`
	PlanPhaseOrder                        int64                                                   `json:"plan_phase_order,required,nullable"`
	PriceType                             PriceModelScalableMatrixWithTieredPricingPricePriceType `json:"price_type,required"`
	ScalableMatrixWithTieredPricingConfig CustomRatingFunctionConfigModel                         `json:"scalable_matrix_with_tiered_pricing_config,required"`
	DimensionalPriceConfiguration         DimensionalPriceConfigurationModel                      `json:"dimensional_price_configuration,nullable"`
	JSON                                  priceModelScalableMatrixWithTieredPricingPriceJSON      `json:"-"`
}

// priceModelScalableMatrixWithTieredPricingPriceJSON contains the JSON metadata
// for the struct [PriceModelScalableMatrixWithTieredPricingPrice]
type priceModelScalableMatrixWithTieredPricingPriceJSON struct {
	ID                                    apijson.Field
	BillableMetric                        apijson.Field
	BillingCycleConfiguration             apijson.Field
	Cadence                               apijson.Field
	ConversionRate                        apijson.Field
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
	ScalableMatrixWithTieredPricingConfig apijson.Field
	DimensionalPriceConfiguration         apijson.Field
	raw                                   string
	ExtraFields                           map[string]apijson.Field
}

func (r *PriceModelScalableMatrixWithTieredPricingPrice) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r priceModelScalableMatrixWithTieredPricingPriceJSON) RawJSON() string {
	return r.raw
}

func (r PriceModelScalableMatrixWithTieredPricingPrice) implementsPriceModel() {}

type PriceModelScalableMatrixWithTieredPricingPriceCadence string

const (
	PriceModelScalableMatrixWithTieredPricingPriceCadenceOneTime    PriceModelScalableMatrixWithTieredPricingPriceCadence = "one_time"
	PriceModelScalableMatrixWithTieredPricingPriceCadenceMonthly    PriceModelScalableMatrixWithTieredPricingPriceCadence = "monthly"
	PriceModelScalableMatrixWithTieredPricingPriceCadenceQuarterly  PriceModelScalableMatrixWithTieredPricingPriceCadence = "quarterly"
	PriceModelScalableMatrixWithTieredPricingPriceCadenceSemiAnnual PriceModelScalableMatrixWithTieredPricingPriceCadence = "semi_annual"
	PriceModelScalableMatrixWithTieredPricingPriceCadenceAnnual     PriceModelScalableMatrixWithTieredPricingPriceCadence = "annual"
	PriceModelScalableMatrixWithTieredPricingPriceCadenceCustom     PriceModelScalableMatrixWithTieredPricingPriceCadence = "custom"
)

func (r PriceModelScalableMatrixWithTieredPricingPriceCadence) IsKnown() bool {
	switch r {
	case PriceModelScalableMatrixWithTieredPricingPriceCadenceOneTime, PriceModelScalableMatrixWithTieredPricingPriceCadenceMonthly, PriceModelScalableMatrixWithTieredPricingPriceCadenceQuarterly, PriceModelScalableMatrixWithTieredPricingPriceCadenceSemiAnnual, PriceModelScalableMatrixWithTieredPricingPriceCadenceAnnual, PriceModelScalableMatrixWithTieredPricingPriceCadenceCustom:
		return true
	}
	return false
}

type PriceModelScalableMatrixWithTieredPricingPriceModelType string

const (
	PriceModelScalableMatrixWithTieredPricingPriceModelTypeScalableMatrixWithTieredPricing PriceModelScalableMatrixWithTieredPricingPriceModelType = "scalable_matrix_with_tiered_pricing"
)

func (r PriceModelScalableMatrixWithTieredPricingPriceModelType) IsKnown() bool {
	switch r {
	case PriceModelScalableMatrixWithTieredPricingPriceModelTypeScalableMatrixWithTieredPricing:
		return true
	}
	return false
}

type PriceModelScalableMatrixWithTieredPricingPricePriceType string

const (
	PriceModelScalableMatrixWithTieredPricingPricePriceTypeUsagePrice PriceModelScalableMatrixWithTieredPricingPricePriceType = "usage_price"
	PriceModelScalableMatrixWithTieredPricingPricePriceTypeFixedPrice PriceModelScalableMatrixWithTieredPricingPricePriceType = "fixed_price"
)

func (r PriceModelScalableMatrixWithTieredPricingPricePriceType) IsKnown() bool {
	switch r {
	case PriceModelScalableMatrixWithTieredPricingPricePriceTypeUsagePrice, PriceModelScalableMatrixWithTieredPricingPricePriceTypeFixedPrice:
		return true
	}
	return false
}

type PriceModelCumulativeGroupedBulkPrice struct {
	ID                          string                                      `json:"id,required"`
	BillableMetric              BillableMetricTinyModel                     `json:"billable_metric,required,nullable"`
	BillingCycleConfiguration   BillingCycleConfigurationModel              `json:"billing_cycle_configuration,required"`
	Cadence                     PriceModelCumulativeGroupedBulkPriceCadence `json:"cadence,required"`
	ConversionRate              float64                                     `json:"conversion_rate,required,nullable"`
	CreatedAt                   time.Time                                   `json:"created_at,required" format:"date-time"`
	CreditAllocation            AllocationModel                             `json:"credit_allocation,required,nullable"`
	CumulativeGroupedBulkConfig CustomRatingFunctionConfigModel             `json:"cumulative_grouped_bulk_config,required"`
	Currency                    string                                      `json:"currency,required"`
	Discount                    Discount                                    `json:"discount,required,nullable"`
	ExternalPriceID             string                                      `json:"external_price_id,required,nullable"`
	FixedPriceQuantity          float64                                     `json:"fixed_price_quantity,required,nullable"`
	InvoicingCycleConfiguration BillingCycleConfigurationModel              `json:"invoicing_cycle_configuration,required,nullable"`
	Item                        ItemSlimModel                               `json:"item,required"`
	Maximum                     MaximumModel                                `json:"maximum,required,nullable"`
	MaximumAmount               string                                      `json:"maximum_amount,required,nullable"`
	// User specified key-value pairs for the resource. If not present, this defaults
	// to an empty dictionary. Individual keys can be removed by setting the value to
	// `null`, and the entire metadata mapping can be cleared by setting `metadata` to
	// `null`.
	Metadata                      map[string]string                             `json:"metadata,required"`
	Minimum                       MinimumModel                                  `json:"minimum,required,nullable"`
	MinimumAmount                 string                                        `json:"minimum_amount,required,nullable"`
	ModelType                     PriceModelCumulativeGroupedBulkPriceModelType `json:"model_type,required"`
	Name                          string                                        `json:"name,required"`
	PlanPhaseOrder                int64                                         `json:"plan_phase_order,required,nullable"`
	PriceType                     PriceModelCumulativeGroupedBulkPricePriceType `json:"price_type,required"`
	DimensionalPriceConfiguration DimensionalPriceConfigurationModel            `json:"dimensional_price_configuration,nullable"`
	JSON                          priceModelCumulativeGroupedBulkPriceJSON      `json:"-"`
}

// priceModelCumulativeGroupedBulkPriceJSON contains the JSON metadata for the
// struct [PriceModelCumulativeGroupedBulkPrice]
type priceModelCumulativeGroupedBulkPriceJSON struct {
	ID                            apijson.Field
	BillableMetric                apijson.Field
	BillingCycleConfiguration     apijson.Field
	Cadence                       apijson.Field
	ConversionRate                apijson.Field
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
	DimensionalPriceConfiguration apijson.Field
	raw                           string
	ExtraFields                   map[string]apijson.Field
}

func (r *PriceModelCumulativeGroupedBulkPrice) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r priceModelCumulativeGroupedBulkPriceJSON) RawJSON() string {
	return r.raw
}

func (r PriceModelCumulativeGroupedBulkPrice) implementsPriceModel() {}

type PriceModelCumulativeGroupedBulkPriceCadence string

const (
	PriceModelCumulativeGroupedBulkPriceCadenceOneTime    PriceModelCumulativeGroupedBulkPriceCadence = "one_time"
	PriceModelCumulativeGroupedBulkPriceCadenceMonthly    PriceModelCumulativeGroupedBulkPriceCadence = "monthly"
	PriceModelCumulativeGroupedBulkPriceCadenceQuarterly  PriceModelCumulativeGroupedBulkPriceCadence = "quarterly"
	PriceModelCumulativeGroupedBulkPriceCadenceSemiAnnual PriceModelCumulativeGroupedBulkPriceCadence = "semi_annual"
	PriceModelCumulativeGroupedBulkPriceCadenceAnnual     PriceModelCumulativeGroupedBulkPriceCadence = "annual"
	PriceModelCumulativeGroupedBulkPriceCadenceCustom     PriceModelCumulativeGroupedBulkPriceCadence = "custom"
)

func (r PriceModelCumulativeGroupedBulkPriceCadence) IsKnown() bool {
	switch r {
	case PriceModelCumulativeGroupedBulkPriceCadenceOneTime, PriceModelCumulativeGroupedBulkPriceCadenceMonthly, PriceModelCumulativeGroupedBulkPriceCadenceQuarterly, PriceModelCumulativeGroupedBulkPriceCadenceSemiAnnual, PriceModelCumulativeGroupedBulkPriceCadenceAnnual, PriceModelCumulativeGroupedBulkPriceCadenceCustom:
		return true
	}
	return false
}

type PriceModelCumulativeGroupedBulkPriceModelType string

const (
	PriceModelCumulativeGroupedBulkPriceModelTypeCumulativeGroupedBulk PriceModelCumulativeGroupedBulkPriceModelType = "cumulative_grouped_bulk"
)

func (r PriceModelCumulativeGroupedBulkPriceModelType) IsKnown() bool {
	switch r {
	case PriceModelCumulativeGroupedBulkPriceModelTypeCumulativeGroupedBulk:
		return true
	}
	return false
}

type PriceModelCumulativeGroupedBulkPricePriceType string

const (
	PriceModelCumulativeGroupedBulkPricePriceTypeUsagePrice PriceModelCumulativeGroupedBulkPricePriceType = "usage_price"
	PriceModelCumulativeGroupedBulkPricePriceTypeFixedPrice PriceModelCumulativeGroupedBulkPricePriceType = "fixed_price"
)

func (r PriceModelCumulativeGroupedBulkPricePriceType) IsKnown() bool {
	switch r {
	case PriceModelCumulativeGroupedBulkPricePriceTypeUsagePrice, PriceModelCumulativeGroupedBulkPricePriceTypeFixedPrice:
		return true
	}
	return false
}

type PriceModelCadence string

const (
	PriceModelCadenceOneTime    PriceModelCadence = "one_time"
	PriceModelCadenceMonthly    PriceModelCadence = "monthly"
	PriceModelCadenceQuarterly  PriceModelCadence = "quarterly"
	PriceModelCadenceSemiAnnual PriceModelCadence = "semi_annual"
	PriceModelCadenceAnnual     PriceModelCadence = "annual"
	PriceModelCadenceCustom     PriceModelCadence = "custom"
)

func (r PriceModelCadence) IsKnown() bool {
	switch r {
	case PriceModelCadenceOneTime, PriceModelCadenceMonthly, PriceModelCadenceQuarterly, PriceModelCadenceSemiAnnual, PriceModelCadenceAnnual, PriceModelCadenceCustom:
		return true
	}
	return false
}

type PriceModelModelType string

const (
	PriceModelModelTypeUnit                            PriceModelModelType = "unit"
	PriceModelModelTypePackage                         PriceModelModelType = "package"
	PriceModelModelTypeMatrix                          PriceModelModelType = "matrix"
	PriceModelModelTypeTiered                          PriceModelModelType = "tiered"
	PriceModelModelTypeTieredBps                       PriceModelModelType = "tiered_bps"
	PriceModelModelTypeBps                             PriceModelModelType = "bps"
	PriceModelModelTypeBulkBps                         PriceModelModelType = "bulk_bps"
	PriceModelModelTypeBulk                            PriceModelModelType = "bulk"
	PriceModelModelTypeThresholdTotalAmount            PriceModelModelType = "threshold_total_amount"
	PriceModelModelTypeTieredPackage                   PriceModelModelType = "tiered_package"
	PriceModelModelTypeGroupedTiered                   PriceModelModelType = "grouped_tiered"
	PriceModelModelTypeTieredWithMinimum               PriceModelModelType = "tiered_with_minimum"
	PriceModelModelTypeTieredPackageWithMinimum        PriceModelModelType = "tiered_package_with_minimum"
	PriceModelModelTypePackageWithAllocation           PriceModelModelType = "package_with_allocation"
	PriceModelModelTypeUnitWithPercent                 PriceModelModelType = "unit_with_percent"
	PriceModelModelTypeMatrixWithAllocation            PriceModelModelType = "matrix_with_allocation"
	PriceModelModelTypeTieredWithProration             PriceModelModelType = "tiered_with_proration"
	PriceModelModelTypeUnitWithProration               PriceModelModelType = "unit_with_proration"
	PriceModelModelTypeGroupedAllocation               PriceModelModelType = "grouped_allocation"
	PriceModelModelTypeGroupedWithProratedMinimum      PriceModelModelType = "grouped_with_prorated_minimum"
	PriceModelModelTypeGroupedWithMeteredMinimum       PriceModelModelType = "grouped_with_metered_minimum"
	PriceModelModelTypeMatrixWithDisplayName           PriceModelModelType = "matrix_with_display_name"
	PriceModelModelTypeBulkWithProration               PriceModelModelType = "bulk_with_proration"
	PriceModelModelTypeGroupedTieredPackage            PriceModelModelType = "grouped_tiered_package"
	PriceModelModelTypeMaxGroupTieredPackage           PriceModelModelType = "max_group_tiered_package"
	PriceModelModelTypeScalableMatrixWithUnitPricing   PriceModelModelType = "scalable_matrix_with_unit_pricing"
	PriceModelModelTypeScalableMatrixWithTieredPricing PriceModelModelType = "scalable_matrix_with_tiered_pricing"
	PriceModelModelTypeCumulativeGroupedBulk           PriceModelModelType = "cumulative_grouped_bulk"
)

func (r PriceModelModelType) IsKnown() bool {
	switch r {
	case PriceModelModelTypeUnit, PriceModelModelTypePackage, PriceModelModelTypeMatrix, PriceModelModelTypeTiered, PriceModelModelTypeTieredBps, PriceModelModelTypeBps, PriceModelModelTypeBulkBps, PriceModelModelTypeBulk, PriceModelModelTypeThresholdTotalAmount, PriceModelModelTypeTieredPackage, PriceModelModelTypeGroupedTiered, PriceModelModelTypeTieredWithMinimum, PriceModelModelTypeTieredPackageWithMinimum, PriceModelModelTypePackageWithAllocation, PriceModelModelTypeUnitWithPercent, PriceModelModelTypeMatrixWithAllocation, PriceModelModelTypeTieredWithProration, PriceModelModelTypeUnitWithProration, PriceModelModelTypeGroupedAllocation, PriceModelModelTypeGroupedWithProratedMinimum, PriceModelModelTypeGroupedWithMeteredMinimum, PriceModelModelTypeMatrixWithDisplayName, PriceModelModelTypeBulkWithProration, PriceModelModelTypeGroupedTieredPackage, PriceModelModelTypeMaxGroupTieredPackage, PriceModelModelTypeScalableMatrixWithUnitPricing, PriceModelModelTypeScalableMatrixWithTieredPricing, PriceModelModelTypeCumulativeGroupedBulk:
		return true
	}
	return false
}

type PriceModelPriceType string

const (
	PriceModelPriceTypeUsagePrice PriceModelPriceType = "usage_price"
	PriceModelPriceTypeFixedPrice PriceModelPriceType = "fixed_price"
)

func (r PriceModelPriceType) IsKnown() bool {
	switch r {
	case PriceModelPriceTypeUsagePrice, PriceModelPriceTypeFixedPrice:
		return true
	}
	return false
}

type RemoveSubscriptionAdjustmentParams struct {
	// The id of the adjustment to remove on the subscription.
	AdjustmentID param.Field[string] `json:"adjustment_id,required"`
}

func (r RemoveSubscriptionAdjustmentParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type RemoveSubscriptionPriceParams struct {
	// The external price id of the price to remove on the subscription.
	ExternalPriceID param.Field[string] `json:"external_price_id"`
	// The id of the price to remove on the subscription.
	PriceID param.Field[string] `json:"price_id"`
}

func (r RemoveSubscriptionPriceParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type ReplaceSubscriptionAdjustmentParams struct {
	// The definition of a new adjustment to create and add to the subscription.
	Adjustment param.Field[NewAdjustmentModelUnionParam] `json:"adjustment,required"`
	// The id of the adjustment on the plan to replace in the subscription.
	ReplacesAdjustmentID param.Field[string] `json:"replaces_adjustment_id,required"`
}

func (r ReplaceSubscriptionAdjustmentParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type ReplaceSubscriptionPriceParams struct {
	// The id of the price on the plan to replace in the subscription.
	ReplacesPriceID param.Field[string] `json:"replaces_price_id,required"`
	// The definition of a new allocation price to create and add to the subscription.
	AllocationPrice param.Field[NewAllocationPriceModelParam] `json:"allocation_price"`
	// [DEPRECATED] Use add_adjustments instead. The subscription's discounts for the
	// replacement price.
	//
	// Deprecated: deprecated
	Discounts param.Field[[]DiscountOverrideModelParam] `json:"discounts"`
	// The external price id of the price to add to the subscription.
	ExternalPriceID param.Field[string] `json:"external_price_id"`
	// The new quantity of the price, if the price is a fixed price.
	FixedPriceQuantity param.Field[float64] `json:"fixed_price_quantity"`
	// [DEPRECATED] Use add_adjustments instead. The subscription's maximum amount for
	// the replacement price.
	//
	// Deprecated: deprecated
	MaximumAmount param.Field[string] `json:"maximum_amount"`
	// [DEPRECATED] Use add_adjustments instead. The subscription's minimum amount for
	// the replacement price.
	//
	// Deprecated: deprecated
	MinimumAmount param.Field[string] `json:"minimum_amount"`
	// The definition of a new price to create and add to the subscription.
	Price param.Field[NewSubscriptionPriceModelUnionParam] `json:"price"`
	// The id of the price to add to the subscription.
	PriceID param.Field[string] `json:"price_id"`
}

func (r ReplaceSubscriptionPriceParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type SubLineItemGroupingModel struct {
	Key string `json:"key,required"`
	// No value indicates the default group
	Value string                       `json:"value,required,nullable"`
	JSON  subLineItemGroupingModelJSON `json:"-"`
}

// subLineItemGroupingModelJSON contains the JSON metadata for the struct
// [SubLineItemGroupingModel]
type subLineItemGroupingModelJSON struct {
	Key         apijson.Field
	Value       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *SubLineItemGroupingModel) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r subLineItemGroupingModelJSON) RawJSON() string {
	return r.raw
}

type SubscriptionMinifiedModel struct {
	ID   string                        `json:"id,required"`
	JSON subscriptionMinifiedModelJSON `json:"-"`
}

// subscriptionMinifiedModelJSON contains the JSON metadata for the struct
// [SubscriptionMinifiedModel]
type subscriptionMinifiedModelJSON struct {
	ID          apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *SubscriptionMinifiedModel) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r subscriptionMinifiedModelJSON) RawJSON() string {
	return r.raw
}

// A [subscription](/core-concepts#subscription) represents the purchase of a plan
// by a customer.
//
// By default, subscriptions begin on the day that they're created and renew
// automatically for each billing cycle at the cadence that's configured in the
// plan definition.
//
// Subscriptions also default to **beginning of month alignment**, which means the
// first invoice issued for the subscription will have pro-rated charges between
// the `start_date` and the first of the following month. Subsequent billing
// periods will always start and end on a month boundary (e.g. subsequent month
// starts for monthly billing).
//
// Depending on the plan configuration, any _flat_ recurring fees will be billed
// either at the beginning (in-advance) or end (in-arrears) of each billing cycle.
// Plans default to **in-advance billing**. Usage-based fees are billed in arrears
// as usage is accumulated. In the normal course of events, you can expect an
// invoice to contain usage-based charges for the previous period, and a recurring
// fee for the following period.
type SubscriptionModel struct {
	ID string `json:"id,required"`
	// The current plan phase that is active, only if the subscription's plan has
	// phases.
	ActivePlanPhaseOrder int64 `json:"active_plan_phase_order,required,nullable"`
	// The adjustment intervals for this subscription.
	AdjustmentIntervals []AdjustmentIntervalModel `json:"adjustment_intervals,required"`
	// Determines whether issued invoices for this subscription will automatically be
	// charged with the saved payment method on the due date. This property defaults to
	// the plan's behavior. If null, defaults to the customer's setting.
	AutoCollection                  bool                                 `json:"auto_collection,required,nullable"`
	BillingCycleAnchorConfiguration BillingCycleAnchorConfigurationModel `json:"billing_cycle_anchor_configuration,required"`
	// The day of the month on which the billing cycle is anchored. If the maximum
	// number of days in a month is greater than this value, the last day of the month
	// is the billing cycle day (e.g. billing_cycle_day=31 for April means the billing
	// period begins on the 30th.
	BillingCycleDay int64     `json:"billing_cycle_day,required"`
	CreatedAt       time.Time `json:"created_at,required" format:"date-time"`
	// The end of the current billing period. This is an exclusive timestamp, such that
	// the instant returned is not part of the billing period. Set to null for
	// subscriptions that are not currently active.
	CurrentBillingPeriodEndDate time.Time `json:"current_billing_period_end_date,required,nullable" format:"date-time"`
	// The start date of the current billing period. This is an inclusive timestamp;
	// the instant returned is exactly the beginning of the billing period. Set to null
	// if the subscription is not currently active.
	CurrentBillingPeriodStartDate time.Time `json:"current_billing_period_start_date,required,nullable" format:"date-time"`
	// A customer is a buyer of your products, and the other party to the billing
	// relationship.
	//
	// In Orb, customers are assigned system generated identifiers automatically, but
	// it's often desirable to have these match existing identifiers in your system. To
	// avoid having to denormalize Orb ID information, you can pass in an
	// `external_customer_id` with your own identifier. See
	// [Customer ID Aliases](/events-and-metrics/customer-aliases) for further
	// information about how these aliases work in Orb.
	//
	// In addition to having an identifier in your system, a customer may exist in a
	// payment provider solution like Stripe. Use the `payment_provider_id` and the
	// `payment_provider` enum field to express this mapping.
	//
	// A customer also has a timezone (from the standard
	// [IANA timezone database](https://www.iana.org/time-zones)), which defaults to
	// your account's timezone. See [Timezone localization](/essentials/timezones) for
	// information on what this timezone parameter influences within Orb.
	Customer CustomerModel `json:"customer,required"`
	// Determines the default memo on this subscriptions' invoices. Note that if this
	// is not provided, it is determined by the plan configuration.
	DefaultInvoiceMemo string `json:"default_invoice_memo,required,nullable"`
	// The discount intervals for this subscription.
	DiscountIntervals []SubscriptionModelDiscountInterval `json:"discount_intervals,required"`
	// The date Orb stops billing for this subscription.
	EndDate                  time.Time                            `json:"end_date,required,nullable" format:"date-time"`
	FixedFeeQuantitySchedule []FixedFeeQuantityScheduleEntryModel `json:"fixed_fee_quantity_schedule,required"`
	InvoicingThreshold       string                               `json:"invoicing_threshold,required,nullable"`
	// The maximum intervals for this subscription.
	MaximumIntervals []MaximumIntervalModel `json:"maximum_intervals,required"`
	// User specified key-value pairs for the resource. If not present, this defaults
	// to an empty dictionary. Individual keys can be removed by setting the value to
	// `null`, and the entire metadata mapping can be cleared by setting `metadata` to
	// `null`.
	Metadata map[string]string `json:"metadata,required"`
	// The minimum intervals for this subscription.
	MinimumIntervals []MinimumIntervalModel `json:"minimum_intervals,required"`
	// Determines the difference between the invoice issue date for subscription
	// invoices as the date that they are due. A value of `0` here represents that the
	// invoice is due on issue, whereas a value of `30` represents that the customer
	// has a month to pay the invoice.
	NetTerms int64 `json:"net_terms,required"`
	// The [Plan](/core-concepts#plan-and-price) resource represents a plan that can be
	// subscribed to by a customer. Plans define the billing behavior of the
	// subscription. You can see more about how to configure prices in the
	// [Price resource](/reference/price).
	Plan PlanModel `json:"plan,required"`
	// The price intervals for this subscription.
	PriceIntervals []PriceIntervalModel  `json:"price_intervals,required"`
	RedeemedCoupon CouponRedemptionModel `json:"redeemed_coupon,required,nullable"`
	// The date Orb starts billing for this subscription.
	StartDate time.Time                  `json:"start_date,required" format:"date-time"`
	Status    SubscriptionModelStatus    `json:"status,required"`
	TrialInfo SubscriptionTrialInfoModel `json:"trial_info,required"`
	JSON      subscriptionModelJSON      `json:"-"`
}

// subscriptionModelJSON contains the JSON metadata for the struct
// [SubscriptionModel]
type subscriptionModelJSON struct {
	ID                              apijson.Field
	ActivePlanPhaseOrder            apijson.Field
	AdjustmentIntervals             apijson.Field
	AutoCollection                  apijson.Field
	BillingCycleAnchorConfiguration apijson.Field
	BillingCycleDay                 apijson.Field
	CreatedAt                       apijson.Field
	CurrentBillingPeriodEndDate     apijson.Field
	CurrentBillingPeriodStartDate   apijson.Field
	Customer                        apijson.Field
	DefaultInvoiceMemo              apijson.Field
	DiscountIntervals               apijson.Field
	EndDate                         apijson.Field
	FixedFeeQuantitySchedule        apijson.Field
	InvoicingThreshold              apijson.Field
	MaximumIntervals                apijson.Field
	Metadata                        apijson.Field
	MinimumIntervals                apijson.Field
	NetTerms                        apijson.Field
	Plan                            apijson.Field
	PriceIntervals                  apijson.Field
	RedeemedCoupon                  apijson.Field
	StartDate                       apijson.Field
	Status                          apijson.Field
	TrialInfo                       apijson.Field
	raw                             string
	ExtraFields                     map[string]apijson.Field
}

func (r *SubscriptionModel) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r subscriptionModelJSON) RawJSON() string {
	return r.raw
}

type SubscriptionModelDiscountInterval struct {
	// This field can have the runtime type of [[]string].
	AppliesToPriceIDs interface{} `json:"applies_to_price_ids,required"`
	// This field can have the runtime type of [[]string].
	AppliesToPriceIntervalIDs interface{}                                    `json:"applies_to_price_interval_ids,required"`
	DiscountType              SubscriptionModelDiscountIntervalsDiscountType `json:"discount_type,required"`
	// The end date of the discount interval.
	EndDate time.Time `json:"end_date,required,nullable" format:"date-time"`
	// The start date of the discount interval.
	StartDate time.Time `json:"start_date,required" format:"date-time"`
	// Only available if discount_type is `amount`.
	AmountDiscount string `json:"amount_discount"`
	// Only available if discount_type is `percentage`.This is a number between 0
	// and 1.
	PercentageDiscount float64 `json:"percentage_discount"`
	// Only available if discount_type is `usage`. Number of usage units that this
	// discount is for
	UsageDiscount float64                               `json:"usage_discount"`
	JSON          subscriptionModelDiscountIntervalJSON `json:"-"`
	union         SubscriptionModelDiscountIntervalsUnion
}

// subscriptionModelDiscountIntervalJSON contains the JSON metadata for the struct
// [SubscriptionModelDiscountInterval]
type subscriptionModelDiscountIntervalJSON struct {
	AppliesToPriceIDs         apijson.Field
	AppliesToPriceIntervalIDs apijson.Field
	DiscountType              apijson.Field
	EndDate                   apijson.Field
	StartDate                 apijson.Field
	AmountDiscount            apijson.Field
	PercentageDiscount        apijson.Field
	UsageDiscount             apijson.Field
	raw                       string
	ExtraFields               map[string]apijson.Field
}

func (r subscriptionModelDiscountIntervalJSON) RawJSON() string {
	return r.raw
}

func (r *SubscriptionModelDiscountInterval) UnmarshalJSON(data []byte) (err error) {
	*r = SubscriptionModelDiscountInterval{}
	err = apijson.UnmarshalRoot(data, &r.union)
	if err != nil {
		return err
	}
	return apijson.Port(r.union, &r)
}

// AsUnion returns a [SubscriptionModelDiscountIntervalsUnion] interface which you
// can cast to the specific types for more type safety.
//
// Possible runtime types of the union are [shared.AmountDiscountIntervalModel],
// [shared.PercentageDiscountIntervalModel], [shared.UsageDiscountIntervalModel].
func (r SubscriptionModelDiscountInterval) AsUnion() SubscriptionModelDiscountIntervalsUnion {
	return r.union
}

// Union satisfied by [shared.AmountDiscountIntervalModel],
// [shared.PercentageDiscountIntervalModel] or [shared.UsageDiscountIntervalModel].
type SubscriptionModelDiscountIntervalsUnion interface {
	ImplementsSubscriptionModelDiscountInterval()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*SubscriptionModelDiscountIntervalsUnion)(nil)).Elem(),
		"discount_type",
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(AmountDiscountIntervalModel{}),
			DiscriminatorValue: "amount",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(PercentageDiscountIntervalModel{}),
			DiscriminatorValue: "percentage",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(UsageDiscountIntervalModel{}),
			DiscriminatorValue: "usage",
		},
	)
}

type SubscriptionModelDiscountIntervalsDiscountType string

const (
	SubscriptionModelDiscountIntervalsDiscountTypeAmount     SubscriptionModelDiscountIntervalsDiscountType = "amount"
	SubscriptionModelDiscountIntervalsDiscountTypePercentage SubscriptionModelDiscountIntervalsDiscountType = "percentage"
	SubscriptionModelDiscountIntervalsDiscountTypeUsage      SubscriptionModelDiscountIntervalsDiscountType = "usage"
)

func (r SubscriptionModelDiscountIntervalsDiscountType) IsKnown() bool {
	switch r {
	case SubscriptionModelDiscountIntervalsDiscountTypeAmount, SubscriptionModelDiscountIntervalsDiscountTypePercentage, SubscriptionModelDiscountIntervalsDiscountTypeUsage:
		return true
	}
	return false
}

type SubscriptionModelStatus string

const (
	SubscriptionModelStatusActive   SubscriptionModelStatus = "active"
	SubscriptionModelStatusEnded    SubscriptionModelStatus = "ended"
	SubscriptionModelStatusUpcoming SubscriptionModelStatus = "upcoming"
)

func (r SubscriptionModelStatus) IsKnown() bool {
	switch r {
	case SubscriptionModelStatusActive, SubscriptionModelStatusEnded, SubscriptionModelStatusUpcoming:
		return true
	}
	return false
}

type SubscriptionTrialInfoModel struct {
	EndDate time.Time                      `json:"end_date,required,nullable" format:"date-time"`
	JSON    subscriptionTrialInfoModelJSON `json:"-"`
}

// subscriptionTrialInfoModelJSON contains the JSON metadata for the struct
// [SubscriptionTrialInfoModel]
type subscriptionTrialInfoModelJSON struct {
	EndDate     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *SubscriptionTrialInfoModel) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r subscriptionTrialInfoModelJSON) RawJSON() string {
	return r.raw
}

type SubscriptionsModel struct {
	Data               []SubscriptionModel    `json:"data,required"`
	PaginationMetadata PaginationMetadata     `json:"pagination_metadata,required"`
	JSON               subscriptionsModelJSON `json:"-"`
}

// subscriptionsModelJSON contains the JSON metadata for the struct
// [SubscriptionsModel]
type subscriptionsModelJSON struct {
	Data               apijson.Field
	PaginationMetadata apijson.Field
	raw                string
	ExtraFields        map[string]apijson.Field
}

func (r *SubscriptionsModel) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r subscriptionsModelJSON) RawJSON() string {
	return r.raw
}

type TaxAmountModel struct {
	// The amount of additional tax incurred by this tax rate.
	Amount string `json:"amount,required"`
	// The human-readable description of the applied tax rate.
	TaxRateDescription string `json:"tax_rate_description,required"`
	// The tax rate percentage, out of 100.
	TaxRatePercentage string             `json:"tax_rate_percentage,required,nullable"`
	JSON              taxAmountModelJSON `json:"-"`
}

// taxAmountModelJSON contains the JSON metadata for the struct [TaxAmountModel]
type taxAmountModelJSON struct {
	Amount             apijson.Field
	TaxRateDescription apijson.Field
	TaxRatePercentage  apijson.Field
	raw                string
	ExtraFields        map[string]apijson.Field
}

func (r *TaxAmountModel) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r taxAmountModelJSON) RawJSON() string {
	return r.raw
}

// Thresholds are used to define the conditions under which an alert will be
// triggered.
type ThresholdModel struct {
	// The value at which an alert will fire. For credit balance alerts, the alert will
	// fire at or below this value. For usage and cost alerts, the alert will fire at
	// or above this value.
	Value float64            `json:"value,required"`
	JSON  thresholdModelJSON `json:"-"`
}

// thresholdModelJSON contains the JSON metadata for the struct [ThresholdModel]
type thresholdModelJSON struct {
	Value       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ThresholdModel) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r thresholdModelJSON) RawJSON() string {
	return r.raw
}

// Thresholds are used to define the conditions under which an alert will be
// triggered.
type ThresholdModelParam struct {
	// The value at which an alert will fire. For credit balance alerts, the alert will
	// fire at or below this value. For usage and cost alerts, the alert will fire at
	// or above this value.
	Value param.Field[float64] `json:"value,required"`
}

func (r ThresholdModelParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type TieredBpsConfigModel struct {
	// Tiers for a Graduated BPS pricing model, where usage is bucketed into specified
	// tiers
	Tiers []TieredBpsConfigModelTier `json:"tiers,required"`
	JSON  tieredBpsConfigModelJSON   `json:"-"`
}

// tieredBpsConfigModelJSON contains the JSON metadata for the struct
// [TieredBpsConfigModel]
type tieredBpsConfigModelJSON struct {
	Tiers       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *TieredBpsConfigModel) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r tieredBpsConfigModelJSON) RawJSON() string {
	return r.raw
}

type TieredBpsConfigModelTier struct {
	// Per-event basis point rate
	Bps float64 `json:"bps,required"`
	// Inclusive tier starting value
	MinimumAmount string `json:"minimum_amount,required"`
	// Exclusive tier ending value
	MaximumAmount string `json:"maximum_amount,nullable"`
	// Per unit maximum to charge
	PerUnitMaximum string                       `json:"per_unit_maximum,nullable"`
	JSON           tieredBpsConfigModelTierJSON `json:"-"`
}

// tieredBpsConfigModelTierJSON contains the JSON metadata for the struct
// [TieredBpsConfigModelTier]
type tieredBpsConfigModelTierJSON struct {
	Bps            apijson.Field
	MinimumAmount  apijson.Field
	MaximumAmount  apijson.Field
	PerUnitMaximum apijson.Field
	raw            string
	ExtraFields    map[string]apijson.Field
}

func (r *TieredBpsConfigModelTier) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r tieredBpsConfigModelTierJSON) RawJSON() string {
	return r.raw
}

type TieredBpsConfigModelParam struct {
	// Tiers for a Graduated BPS pricing model, where usage is bucketed into specified
	// tiers
	Tiers param.Field[[]TieredBpsConfigModelTierParam] `json:"tiers,required"`
}

func (r TieredBpsConfigModelParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type TieredBpsConfigModelTierParam struct {
	// Per-event basis point rate
	Bps param.Field[float64] `json:"bps,required"`
	// Inclusive tier starting value
	MinimumAmount param.Field[string] `json:"minimum_amount,required"`
	// Exclusive tier ending value
	MaximumAmount param.Field[string] `json:"maximum_amount"`
	// Per unit maximum to charge
	PerUnitMaximum param.Field[string] `json:"per_unit_maximum"`
}

func (r TieredBpsConfigModelTierParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type TieredConfigModel struct {
	// Tiers for rating based on total usage quantities into the specified tier
	Tiers []TieredConfigModelTier `json:"tiers,required"`
	JSON  tieredConfigModelJSON   `json:"-"`
}

// tieredConfigModelJSON contains the JSON metadata for the struct
// [TieredConfigModel]
type tieredConfigModelJSON struct {
	Tiers       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *TieredConfigModel) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r tieredConfigModelJSON) RawJSON() string {
	return r.raw
}

type TieredConfigModelTier struct {
	// Inclusive tier starting value
	FirstUnit float64 `json:"first_unit,required"`
	// Amount per unit
	UnitAmount string `json:"unit_amount,required"`
	// Exclusive tier ending value. If null, this is treated as the last tier
	LastUnit float64                   `json:"last_unit,nullable"`
	JSON     tieredConfigModelTierJSON `json:"-"`
}

// tieredConfigModelTierJSON contains the JSON metadata for the struct
// [TieredConfigModelTier]
type tieredConfigModelTierJSON struct {
	FirstUnit   apijson.Field
	UnitAmount  apijson.Field
	LastUnit    apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *TieredConfigModelTier) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r tieredConfigModelTierJSON) RawJSON() string {
	return r.raw
}

type TieredConfigModelParam struct {
	// Tiers for rating based on total usage quantities into the specified tier
	Tiers param.Field[[]TieredConfigModelTierParam] `json:"tiers,required"`
}

func (r TieredConfigModelParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type TieredConfigModelTierParam struct {
	// Inclusive tier starting value
	FirstUnit param.Field[float64] `json:"first_unit,required"`
	// Amount per unit
	UnitAmount param.Field[string] `json:"unit_amount,required"`
	// Exclusive tier ending value. If null, this is treated as the last tier
	LastUnit param.Field[float64] `json:"last_unit"`
}

func (r TieredConfigModelTierParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type TopUpModel struct {
	ID string `json:"id,required"`
	// The amount to increment when the threshold is reached.
	Amount string `json:"amount,required"`
	// The currency or custom pricing unit to use for this top-up. If this is a
	// real-world currency, it must match the customer's invoicing currency.
	Currency string `json:"currency,required"`
	// Settings for invoices generated by triggered top-ups.
	InvoiceSettings TopUpModelInvoiceSettings `json:"invoice_settings,required"`
	// How much, in the customer's currency, to charge for each unit.
	PerUnitCostBasis string `json:"per_unit_cost_basis,required"`
	// The threshold at which to trigger the top-up. If the balance is at or below this
	// threshold, the top-up will be triggered.
	Threshold string `json:"threshold,required"`
	// The number of days or months after which the top-up expires. If unspecified, it
	// does not expire.
	ExpiresAfter int64 `json:"expires_after,nullable"`
	// The unit of expires_after.
	ExpiresAfterUnit TopUpModelExpiresAfterUnit `json:"expires_after_unit,nullable"`
	JSON             topUpModelJSON             `json:"-"`
}

// topUpModelJSON contains the JSON metadata for the struct [TopUpModel]
type topUpModelJSON struct {
	ID               apijson.Field
	Amount           apijson.Field
	Currency         apijson.Field
	InvoiceSettings  apijson.Field
	PerUnitCostBasis apijson.Field
	Threshold        apijson.Field
	ExpiresAfter     apijson.Field
	ExpiresAfterUnit apijson.Field
	raw              string
	ExtraFields      map[string]apijson.Field
}

func (r *TopUpModel) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r topUpModelJSON) RawJSON() string {
	return r.raw
}

// Settings for invoices generated by triggered top-ups.
type TopUpModelInvoiceSettings struct {
	// Whether the credits purchase invoice should auto collect with the customer's
	// saved payment method.
	AutoCollection bool `json:"auto_collection,required"`
	// The net terms determines the difference between the invoice date and the issue
	// date for the invoice. If you intend the invoice to be due on issue, set this
	// to 0.
	NetTerms int64 `json:"net_terms,required"`
	// An optional memo to display on the invoice.
	Memo string `json:"memo,nullable"`
	// If true, new credit blocks created by this top-up will require that the
	// corresponding invoice is paid before they can be drawn down from.
	RequireSuccessfulPayment bool                          `json:"require_successful_payment"`
	JSON                     topUpModelInvoiceSettingsJSON `json:"-"`
}

// topUpModelInvoiceSettingsJSON contains the JSON metadata for the struct
// [TopUpModelInvoiceSettings]
type topUpModelInvoiceSettingsJSON struct {
	AutoCollection           apijson.Field
	NetTerms                 apijson.Field
	Memo                     apijson.Field
	RequireSuccessfulPayment apijson.Field
	raw                      string
	ExtraFields              map[string]apijson.Field
}

func (r *TopUpModelInvoiceSettings) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r topUpModelInvoiceSettingsJSON) RawJSON() string {
	return r.raw
}

// The unit of expires_after.
type TopUpModelExpiresAfterUnit string

const (
	TopUpModelExpiresAfterUnitDay   TopUpModelExpiresAfterUnit = "day"
	TopUpModelExpiresAfterUnitMonth TopUpModelExpiresAfterUnit = "month"
)

func (r TopUpModelExpiresAfterUnit) IsKnown() bool {
	switch r {
	case TopUpModelExpiresAfterUnitDay, TopUpModelExpiresAfterUnitMonth:
		return true
	}
	return false
}

type TopUpsModel struct {
	Data               []TopUpModel       `json:"data,required"`
	PaginationMetadata PaginationMetadata `json:"pagination_metadata,required"`
	JSON               topUpsModelJSON    `json:"-"`
}

// topUpsModelJSON contains the JSON metadata for the struct [TopUpsModel]
type topUpsModelJSON struct {
	Data               apijson.Field
	PaginationMetadata apijson.Field
	raw                string
	ExtraFields        map[string]apijson.Field
}

func (r *TopUpsModel) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r topUpsModelJSON) RawJSON() string {
	return r.raw
}

type TrialDiscount struct {
	// List of price_ids that this discount applies to. For plan/plan phase discounts,
	// this can be a subset of prices.
	AppliesToPriceIDs []string                  `json:"applies_to_price_ids,required"`
	DiscountType      TrialDiscountDiscountType `json:"discount_type,required"`
	Reason            string                    `json:"reason,nullable"`
	// Only available if discount_type is `trial`
	TrialAmountDiscount string `json:"trial_amount_discount,nullable"`
	// Only available if discount_type is `trial`
	TrialPercentageDiscount float64           `json:"trial_percentage_discount,nullable"`
	JSON                    trialDiscountJSON `json:"-"`
}

// trialDiscountJSON contains the JSON metadata for the struct [TrialDiscount]
type trialDiscountJSON struct {
	AppliesToPriceIDs       apijson.Field
	DiscountType            apijson.Field
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
	// List of price_ids that this discount applies to. For plan/plan phase discounts,
	// this can be a subset of prices.
	AppliesToPriceIDs param.Field[[]string]                  `json:"applies_to_price_ids,required"`
	DiscountType      param.Field[TrialDiscountDiscountType] `json:"discount_type,required"`
	Reason            param.Field[string]                    `json:"reason"`
	// Only available if discount_type is `trial`
	TrialAmountDiscount param.Field[string] `json:"trial_amount_discount"`
	// Only available if discount_type is `trial`
	TrialPercentageDiscount param.Field[float64] `json:"trial_percentage_discount"`
}

func (r TrialDiscountParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r TrialDiscountParam) ImplementsDiscountUnionParam() {}

type UnitConfigModel struct {
	// Rate per unit of usage
	UnitAmount string              `json:"unit_amount,required"`
	JSON       unitConfigModelJSON `json:"-"`
}

// unitConfigModelJSON contains the JSON metadata for the struct [UnitConfigModel]
type unitConfigModelJSON struct {
	UnitAmount  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *UnitConfigModel) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r unitConfigModelJSON) RawJSON() string {
	return r.raw
}

type UnitConfigModelParam struct {
	// Rate per unit of usage
	UnitAmount param.Field[string] `json:"unit_amount,required"`
}

func (r UnitConfigModelParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type UsageDiscountIntervalModel struct {
	// The price ids that this discount interval applies to.
	AppliesToPriceIDs []string `json:"applies_to_price_ids,required"`
	// The price interval ids that this discount interval applies to.
	AppliesToPriceIntervalIDs []string                               `json:"applies_to_price_interval_ids,required"`
	DiscountType              UsageDiscountIntervalModelDiscountType `json:"discount_type,required"`
	// The end date of the discount interval.
	EndDate time.Time `json:"end_date,required,nullable" format:"date-time"`
	// The start date of the discount interval.
	StartDate time.Time `json:"start_date,required" format:"date-time"`
	// Only available if discount_type is `usage`. Number of usage units that this
	// discount is for
	UsageDiscount float64                        `json:"usage_discount,required"`
	JSON          usageDiscountIntervalModelJSON `json:"-"`
}

// usageDiscountIntervalModelJSON contains the JSON metadata for the struct
// [UsageDiscountIntervalModel]
type usageDiscountIntervalModelJSON struct {
	AppliesToPriceIDs         apijson.Field
	AppliesToPriceIntervalIDs apijson.Field
	DiscountType              apijson.Field
	EndDate                   apijson.Field
	StartDate                 apijson.Field
	UsageDiscount             apijson.Field
	raw                       string
	ExtraFields               map[string]apijson.Field
}

func (r *UsageDiscountIntervalModel) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r usageDiscountIntervalModelJSON) RawJSON() string {
	return r.raw
}

func (r UsageDiscountIntervalModel) ImplementsMutatedSubscriptionModelDiscountInterval() {}

func (r UsageDiscountIntervalModel) ImplementsSubscriptionModelDiscountInterval() {}

type UsageDiscountIntervalModelDiscountType string

const (
	UsageDiscountIntervalModelDiscountTypeUsage UsageDiscountIntervalModelDiscountType = "usage"
)

func (r UsageDiscountIntervalModelDiscountType) IsKnown() bool {
	switch r {
	case UsageDiscountIntervalModelDiscountTypeUsage:
		return true
	}
	return false
}

type UsageModel struct {
	Quantity       float64        `json:"quantity,required"`
	TimeframeEnd   time.Time      `json:"timeframe_end,required" format:"date-time"`
	TimeframeStart time.Time      `json:"timeframe_start,required" format:"date-time"`
	JSON           usageModelJSON `json:"-"`
}

// usageModelJSON contains the JSON metadata for the struct [UsageModel]
type usageModelJSON struct {
	Quantity       apijson.Field
	TimeframeEnd   apijson.Field
	TimeframeStart apijson.Field
	raw            string
	ExtraFields    map[string]apijson.Field
}

func (r *UsageModel) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r usageModelJSON) RawJSON() string {
	return r.raw
}
