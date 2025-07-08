// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package shared

import (
	"time"
)

type UnionTime time.Time

func (UnionTime) ImplementsCustomerCreditLedgerNewEntryParamsAddIncrementCreditLedgerEntryRequestParamsInvoiceSettingsInvoiceDateUnion() {
}
func (UnionTime) ImplementsCustomerCreditLedgerNewEntryByExternalIDParamsAddIncrementCreditLedgerEntryRequestParamsInvoiceSettingsInvoiceDateUnion() {
}
func (UnionTime) ImplementsSubscriptionPriceIntervalsParamsAddStartDateUnion()             {}
func (UnionTime) ImplementsSubscriptionPriceIntervalsParamsAddEndDateUnion()               {}
func (UnionTime) ImplementsSubscriptionPriceIntervalsParamsAddAdjustmentsStartDateUnion()  {}
func (UnionTime) ImplementsSubscriptionPriceIntervalsParamsAddAdjustmentsEndDateUnion()    {}
func (UnionTime) ImplementsSubscriptionPriceIntervalsParamsEditEndDateUnion()              {}
func (UnionTime) ImplementsSubscriptionPriceIntervalsParamsEditStartDateUnion()            {}
func (UnionTime) ImplementsSubscriptionPriceIntervalsParamsEditAdjustmentsEndDateUnion()   {}
func (UnionTime) ImplementsSubscriptionPriceIntervalsParamsEditAdjustmentsStartDateUnion() {}
func (UnionTime) ImplementsSubscriptionUpdateTrialParamsTrialEndDateUnion()                {}

type UnionString string

func (UnionString) ImplementsEvaluatePriceGroupGroupingValuesUnion() {}

type UnionBool bool

func (UnionBool) ImplementsEvaluatePriceGroupGroupingValuesUnion() {}

type UnionFloat float64

func (UnionFloat) ImplementsEvaluatePriceGroupGroupingValuesUnion() {}
