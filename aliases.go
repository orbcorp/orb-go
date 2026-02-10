// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package orb

import (
	"github.com/orbcorp/orb-go/internal/apierror"
	"github.com/orbcorp/orb-go/shared"
)

type Error = apierror.Error
type ErrorStatus = apierror.ErrorStatus

const ErrorStatus500 = apierror.ErrorStatus500
const ErrorStatus429 = apierror.ErrorStatus429
const ErrorStatus413 = apierror.ErrorStatus413
const ErrorStatus409 = apierror.ErrorStatus409
const ErrorStatus404 = apierror.ErrorStatus404
const ErrorStatus400 = apierror.ErrorStatus400
const ErrorStatus401 = apierror.ErrorStatus401

type ErrorType = apierror.ErrorType

const ErrorTypeOrbInternalServerError = apierror.ErrorTypeOrbInternalServerError
const ErrorTypeTooManyRequests = apierror.ErrorTypeTooManyRequests
const ErrorTypeResourceTooLarge = apierror.ErrorTypeResourceTooLarge
const ErrorTypeRequestTooLarge = apierror.ErrorTypeRequestTooLarge
const ErrorTypeResourceConflict = apierror.ErrorTypeResourceConflict
const ErrorTypeURLNotFound = apierror.ErrorTypeURLNotFound
const ErrorTypeResourceNotFound = apierror.ErrorTypeResourceNotFound
const ErrorTypeFeatureNotAvailable = apierror.ErrorTypeFeatureNotAvailable
const ErrorTypeOrbAuthenticationError = apierror.ErrorTypeOrbAuthenticationError
const ErrorTypeRequestValidationError = apierror.ErrorTypeRequestValidationError
const ErrorTypeDuplicateResourceCreation = apierror.ErrorTypeDuplicateResourceCreation
const ErrorTypeConstraintViolation = apierror.ErrorTypeConstraintViolation

// This is an alias to an internal type.
type Address = shared.Address

// This is an alias to an internal type.
type AdjustmentInterval = shared.AdjustmentInterval

// This is an alias to an internal type.
type AdjustmentIntervalAdjustment = shared.AdjustmentIntervalAdjustment

// This is an alias to an internal type.
type AdjustmentIntervalAdjustmentAdjustmentType = shared.AdjustmentIntervalAdjustmentAdjustmentType

// This is an alias to an internal value.
const AdjustmentIntervalAdjustmentAdjustmentTypeUsageDiscount = shared.AdjustmentIntervalAdjustmentAdjustmentTypeUsageDiscount

// This is an alias to an internal value.
const AdjustmentIntervalAdjustmentAdjustmentTypeAmountDiscount = shared.AdjustmentIntervalAdjustmentAdjustmentTypeAmountDiscount

// This is an alias to an internal value.
const AdjustmentIntervalAdjustmentAdjustmentTypePercentageDiscount = shared.AdjustmentIntervalAdjustmentAdjustmentTypePercentageDiscount

// This is an alias to an internal value.
const AdjustmentIntervalAdjustmentAdjustmentTypeMinimum = shared.AdjustmentIntervalAdjustmentAdjustmentTypeMinimum

// This is an alias to an internal value.
const AdjustmentIntervalAdjustmentAdjustmentTypeMaximum = shared.AdjustmentIntervalAdjustmentAdjustmentTypeMaximum

// This is an alias to an internal type.
type AggregatedCost = shared.AggregatedCost

// This is an alias to an internal type.
type Allocation = shared.Allocation

// This is an alias to an internal type.
type AllocationFilter = shared.AllocationFilter

// The property of the price to filter on.
//
// This is an alias to an internal type.
type AllocationFiltersField = shared.AllocationFiltersField

// This is an alias to an internal value.
const AllocationFiltersFieldPriceID = shared.AllocationFiltersFieldPriceID

// This is an alias to an internal value.
const AllocationFiltersFieldItemID = shared.AllocationFiltersFieldItemID

// This is an alias to an internal value.
const AllocationFiltersFieldPriceType = shared.AllocationFiltersFieldPriceType

// This is an alias to an internal value.
const AllocationFiltersFieldCurrency = shared.AllocationFiltersFieldCurrency

// This is an alias to an internal value.
const AllocationFiltersFieldPricingUnitID = shared.AllocationFiltersFieldPricingUnitID

// Should prices that match the filter be included or excluded.
//
// This is an alias to an internal type.
type AllocationFiltersOperator = shared.AllocationFiltersOperator

// This is an alias to an internal value.
const AllocationFiltersOperatorIncludes = shared.AllocationFiltersOperatorIncludes

// This is an alias to an internal value.
const AllocationFiltersOperatorExcludes = shared.AllocationFiltersOperatorExcludes

// This is an alias to an internal type.
type AmountDiscount = shared.AmountDiscount

// This is an alias to an internal type.
type AmountDiscountDiscountType = shared.AmountDiscountDiscountType

// This is an alias to an internal value.
const AmountDiscountDiscountTypeAmount = shared.AmountDiscountDiscountTypeAmount

// This is an alias to an internal type.
type AmountDiscountFilter = shared.AmountDiscountFilter

// The property of the price to filter on.
//
// This is an alias to an internal type.
type AmountDiscountFiltersField = shared.AmountDiscountFiltersField

// This is an alias to an internal value.
const AmountDiscountFiltersFieldPriceID = shared.AmountDiscountFiltersFieldPriceID

// This is an alias to an internal value.
const AmountDiscountFiltersFieldItemID = shared.AmountDiscountFiltersFieldItemID

// This is an alias to an internal value.
const AmountDiscountFiltersFieldPriceType = shared.AmountDiscountFiltersFieldPriceType

// This is an alias to an internal value.
const AmountDiscountFiltersFieldCurrency = shared.AmountDiscountFiltersFieldCurrency

// This is an alias to an internal value.
const AmountDiscountFiltersFieldPricingUnitID = shared.AmountDiscountFiltersFieldPricingUnitID

// Should prices that match the filter be included or excluded.
//
// This is an alias to an internal type.
type AmountDiscountFiltersOperator = shared.AmountDiscountFiltersOperator

// This is an alias to an internal value.
const AmountDiscountFiltersOperatorIncludes = shared.AmountDiscountFiltersOperatorIncludes

// This is an alias to an internal value.
const AmountDiscountFiltersOperatorExcludes = shared.AmountDiscountFiltersOperatorExcludes

// This is an alias to an internal type.
type AmountDiscountParam = shared.AmountDiscountParam

// This is an alias to an internal type.
type AmountDiscountFilterParam = shared.AmountDiscountFilterParam

// This is an alias to an internal type.
type AmountDiscountInterval = shared.AmountDiscountInterval

// This is an alias to an internal type.
type AmountDiscountIntervalDiscountType = shared.AmountDiscountIntervalDiscountType

// This is an alias to an internal value.
const AmountDiscountIntervalDiscountTypeAmount = shared.AmountDiscountIntervalDiscountTypeAmount

// This is an alias to an internal type.
type AmountDiscountIntervalFilter = shared.AmountDiscountIntervalFilter

// The property of the price to filter on.
//
// This is an alias to an internal type.
type AmountDiscountIntervalFiltersField = shared.AmountDiscountIntervalFiltersField

// This is an alias to an internal value.
const AmountDiscountIntervalFiltersFieldPriceID = shared.AmountDiscountIntervalFiltersFieldPriceID

// This is an alias to an internal value.
const AmountDiscountIntervalFiltersFieldItemID = shared.AmountDiscountIntervalFiltersFieldItemID

// This is an alias to an internal value.
const AmountDiscountIntervalFiltersFieldPriceType = shared.AmountDiscountIntervalFiltersFieldPriceType

// This is an alias to an internal value.
const AmountDiscountIntervalFiltersFieldCurrency = shared.AmountDiscountIntervalFiltersFieldCurrency

// This is an alias to an internal value.
const AmountDiscountIntervalFiltersFieldPricingUnitID = shared.AmountDiscountIntervalFiltersFieldPricingUnitID

// Should prices that match the filter be included or excluded.
//
// This is an alias to an internal type.
type AmountDiscountIntervalFiltersOperator = shared.AmountDiscountIntervalFiltersOperator

// This is an alias to an internal value.
const AmountDiscountIntervalFiltersOperatorIncludes = shared.AmountDiscountIntervalFiltersOperatorIncludes

// This is an alias to an internal value.
const AmountDiscountIntervalFiltersOperatorExcludes = shared.AmountDiscountIntervalFiltersOperatorExcludes

// This is an alias to an internal type.
type BillableMetricTiny = shared.BillableMetricTiny

// This is an alias to an internal type.
type BillingCycleAnchorConfiguration = shared.BillingCycleAnchorConfiguration

// This is an alias to an internal type.
type BillingCycleAnchorConfigurationParam = shared.BillingCycleAnchorConfigurationParam

// This is an alias to an internal type.
type BillingCycleConfiguration = shared.BillingCycleConfiguration

// This is an alias to an internal type.
type BillingCycleConfigurationDurationUnit = shared.BillingCycleConfigurationDurationUnit

// This is an alias to an internal value.
const BillingCycleConfigurationDurationUnitDay = shared.BillingCycleConfigurationDurationUnitDay

// This is an alias to an internal value.
const BillingCycleConfigurationDurationUnitMonth = shared.BillingCycleConfigurationDurationUnitMonth

// This is an alias to an internal type.
type BillingCycleRelativeDate = shared.BillingCycleRelativeDate

// This is an alias to an internal value.
const BillingCycleRelativeDateStartOfTerm = shared.BillingCycleRelativeDateStartOfTerm

// This is an alias to an internal value.
const BillingCycleRelativeDateEndOfTerm = shared.BillingCycleRelativeDateEndOfTerm

// Configuration for bulk pricing
//
// This is an alias to an internal type.
type BulkConfig = shared.BulkConfig

// Configuration for bulk pricing
//
// This is an alias to an internal type.
type BulkConfigParam = shared.BulkConfigParam

// Configuration for a single bulk pricing tier
//
// This is an alias to an internal type.
type BulkTier = shared.BulkTier

// Configuration for a single bulk pricing tier
//
// This is an alias to an internal type.
type BulkTierParam = shared.BulkTierParam

// This is an alias to an internal type.
type ChangedSubscriptionResources = shared.ChangedSubscriptionResources

// This is an alias to an internal type.
type ChangedSubscriptionResourcesCreatedInvoice = shared.ChangedSubscriptionResourcesCreatedInvoice

// This is an alias to an internal type.
type ChangedSubscriptionResourcesCreatedInvoicesAutoCollection = shared.ChangedSubscriptionResourcesCreatedInvoicesAutoCollection

// This is an alias to an internal type.
type ChangedSubscriptionResourcesCreatedInvoicesCreditNote = shared.ChangedSubscriptionResourcesCreatedInvoicesCreditNote

// This is an alias to an internal type.
type ChangedSubscriptionResourcesCreatedInvoicesCustomerBalanceTransaction = shared.ChangedSubscriptionResourcesCreatedInvoicesCustomerBalanceTransaction

// This is an alias to an internal type.
type ChangedSubscriptionResourcesCreatedInvoicesCustomerBalanceTransactionsAction = shared.ChangedSubscriptionResourcesCreatedInvoicesCustomerBalanceTransactionsAction

// This is an alias to an internal value.
const ChangedSubscriptionResourcesCreatedInvoicesCustomerBalanceTransactionsActionAppliedToInvoice = shared.ChangedSubscriptionResourcesCreatedInvoicesCustomerBalanceTransactionsActionAppliedToInvoice

// This is an alias to an internal value.
const ChangedSubscriptionResourcesCreatedInvoicesCustomerBalanceTransactionsActionManualAdjustment = shared.ChangedSubscriptionResourcesCreatedInvoicesCustomerBalanceTransactionsActionManualAdjustment

// This is an alias to an internal value.
const ChangedSubscriptionResourcesCreatedInvoicesCustomerBalanceTransactionsActionProratedRefund = shared.ChangedSubscriptionResourcesCreatedInvoicesCustomerBalanceTransactionsActionProratedRefund

// This is an alias to an internal value.
const ChangedSubscriptionResourcesCreatedInvoicesCustomerBalanceTransactionsActionRevertProratedRefund = shared.ChangedSubscriptionResourcesCreatedInvoicesCustomerBalanceTransactionsActionRevertProratedRefund

// This is an alias to an internal value.
const ChangedSubscriptionResourcesCreatedInvoicesCustomerBalanceTransactionsActionReturnFromVoiding = shared.ChangedSubscriptionResourcesCreatedInvoicesCustomerBalanceTransactionsActionReturnFromVoiding

// This is an alias to an internal value.
const ChangedSubscriptionResourcesCreatedInvoicesCustomerBalanceTransactionsActionCreditNoteApplied = shared.ChangedSubscriptionResourcesCreatedInvoicesCustomerBalanceTransactionsActionCreditNoteApplied

// This is an alias to an internal value.
const ChangedSubscriptionResourcesCreatedInvoicesCustomerBalanceTransactionsActionCreditNoteVoided = shared.ChangedSubscriptionResourcesCreatedInvoicesCustomerBalanceTransactionsActionCreditNoteVoided

// This is an alias to an internal value.
const ChangedSubscriptionResourcesCreatedInvoicesCustomerBalanceTransactionsActionOverpaymentRefund = shared.ChangedSubscriptionResourcesCreatedInvoicesCustomerBalanceTransactionsActionOverpaymentRefund

// This is an alias to an internal value.
const ChangedSubscriptionResourcesCreatedInvoicesCustomerBalanceTransactionsActionExternalPayment = shared.ChangedSubscriptionResourcesCreatedInvoicesCustomerBalanceTransactionsActionExternalPayment

// This is an alias to an internal value.
const ChangedSubscriptionResourcesCreatedInvoicesCustomerBalanceTransactionsActionSmallInvoiceCarryover = shared.ChangedSubscriptionResourcesCreatedInvoicesCustomerBalanceTransactionsActionSmallInvoiceCarryover

// This is an alias to an internal type.
type ChangedSubscriptionResourcesCreatedInvoicesCustomerBalanceTransactionsType = shared.ChangedSubscriptionResourcesCreatedInvoicesCustomerBalanceTransactionsType

// This is an alias to an internal value.
const ChangedSubscriptionResourcesCreatedInvoicesCustomerBalanceTransactionsTypeIncrement = shared.ChangedSubscriptionResourcesCreatedInvoicesCustomerBalanceTransactionsTypeIncrement

// This is an alias to an internal value.
const ChangedSubscriptionResourcesCreatedInvoicesCustomerBalanceTransactionsTypeDecrement = shared.ChangedSubscriptionResourcesCreatedInvoicesCustomerBalanceTransactionsTypeDecrement

// This is an alias to an internal type.
type ChangedSubscriptionResourcesCreatedInvoicesInvoiceSource = shared.ChangedSubscriptionResourcesCreatedInvoicesInvoiceSource

// This is an alias to an internal value.
const ChangedSubscriptionResourcesCreatedInvoicesInvoiceSourceSubscription = shared.ChangedSubscriptionResourcesCreatedInvoicesInvoiceSourceSubscription

// This is an alias to an internal value.
const ChangedSubscriptionResourcesCreatedInvoicesInvoiceSourcePartial = shared.ChangedSubscriptionResourcesCreatedInvoicesInvoiceSourcePartial

// This is an alias to an internal value.
const ChangedSubscriptionResourcesCreatedInvoicesInvoiceSourceOneOff = shared.ChangedSubscriptionResourcesCreatedInvoicesInvoiceSourceOneOff

// This is an alias to an internal type.
type ChangedSubscriptionResourcesCreatedInvoicesLineItem = shared.ChangedSubscriptionResourcesCreatedInvoicesLineItem

// This is an alias to an internal type.
type ChangedSubscriptionResourcesCreatedInvoicesLineItemsAdjustment = shared.ChangedSubscriptionResourcesCreatedInvoicesLineItemsAdjustment

// This is an alias to an internal type.
type ChangedSubscriptionResourcesCreatedInvoicesLineItemsAdjustmentsAdjustmentType = shared.ChangedSubscriptionResourcesCreatedInvoicesLineItemsAdjustmentsAdjustmentType

// This is an alias to an internal value.
const ChangedSubscriptionResourcesCreatedInvoicesLineItemsAdjustmentsAdjustmentTypeUsageDiscount = shared.ChangedSubscriptionResourcesCreatedInvoicesLineItemsAdjustmentsAdjustmentTypeUsageDiscount

// This is an alias to an internal value.
const ChangedSubscriptionResourcesCreatedInvoicesLineItemsAdjustmentsAdjustmentTypeAmountDiscount = shared.ChangedSubscriptionResourcesCreatedInvoicesLineItemsAdjustmentsAdjustmentTypeAmountDiscount

// This is an alias to an internal value.
const ChangedSubscriptionResourcesCreatedInvoicesLineItemsAdjustmentsAdjustmentTypePercentageDiscount = shared.ChangedSubscriptionResourcesCreatedInvoicesLineItemsAdjustmentsAdjustmentTypePercentageDiscount

// This is an alias to an internal value.
const ChangedSubscriptionResourcesCreatedInvoicesLineItemsAdjustmentsAdjustmentTypeMinimum = shared.ChangedSubscriptionResourcesCreatedInvoicesLineItemsAdjustmentsAdjustmentTypeMinimum

// This is an alias to an internal value.
const ChangedSubscriptionResourcesCreatedInvoicesLineItemsAdjustmentsAdjustmentTypeMaximum = shared.ChangedSubscriptionResourcesCreatedInvoicesLineItemsAdjustmentsAdjustmentTypeMaximum

// This is an alias to an internal type.
type ChangedSubscriptionResourcesCreatedInvoicesLineItemsSubLineItem = shared.ChangedSubscriptionResourcesCreatedInvoicesLineItemsSubLineItem

// This is an alias to an internal type.
type ChangedSubscriptionResourcesCreatedInvoicesLineItemsSubLineItemsType = shared.ChangedSubscriptionResourcesCreatedInvoicesLineItemsSubLineItemsType

// This is an alias to an internal value.
const ChangedSubscriptionResourcesCreatedInvoicesLineItemsSubLineItemsTypeMatrix = shared.ChangedSubscriptionResourcesCreatedInvoicesLineItemsSubLineItemsTypeMatrix

// This is an alias to an internal value.
const ChangedSubscriptionResourcesCreatedInvoicesLineItemsSubLineItemsTypeTier = shared.ChangedSubscriptionResourcesCreatedInvoicesLineItemsSubLineItemsTypeTier

// This is an alias to an internal value.
const ChangedSubscriptionResourcesCreatedInvoicesLineItemsSubLineItemsTypeNull = shared.ChangedSubscriptionResourcesCreatedInvoicesLineItemsSubLineItemsTypeNull

// This is an alias to an internal type.
type ChangedSubscriptionResourcesCreatedInvoicesPaymentAttempt = shared.ChangedSubscriptionResourcesCreatedInvoicesPaymentAttempt

// The payment provider that attempted to collect the payment.
//
// This is an alias to an internal type.
type ChangedSubscriptionResourcesCreatedInvoicesPaymentAttemptsPaymentProvider = shared.ChangedSubscriptionResourcesCreatedInvoicesPaymentAttemptsPaymentProvider

// This is an alias to an internal value.
const ChangedSubscriptionResourcesCreatedInvoicesPaymentAttemptsPaymentProviderStripe = shared.ChangedSubscriptionResourcesCreatedInvoicesPaymentAttemptsPaymentProviderStripe

// This is an alias to an internal type.
type ChangedSubscriptionResourcesCreatedInvoicesStatus = shared.ChangedSubscriptionResourcesCreatedInvoicesStatus

// This is an alias to an internal value.
const ChangedSubscriptionResourcesCreatedInvoicesStatusIssued = shared.ChangedSubscriptionResourcesCreatedInvoicesStatusIssued

// This is an alias to an internal value.
const ChangedSubscriptionResourcesCreatedInvoicesStatusPaid = shared.ChangedSubscriptionResourcesCreatedInvoicesStatusPaid

// This is an alias to an internal value.
const ChangedSubscriptionResourcesCreatedInvoicesStatusSynced = shared.ChangedSubscriptionResourcesCreatedInvoicesStatusSynced

// This is an alias to an internal value.
const ChangedSubscriptionResourcesCreatedInvoicesStatusVoid = shared.ChangedSubscriptionResourcesCreatedInvoicesStatusVoid

// This is an alias to an internal value.
const ChangedSubscriptionResourcesCreatedInvoicesStatusDraft = shared.ChangedSubscriptionResourcesCreatedInvoicesStatusDraft

// This is an alias to an internal type.
type ConversionRateTier = shared.ConversionRateTier

// This is an alias to an internal type.
type ConversionRateTierParam = shared.ConversionRateTierParam

// This is an alias to an internal type.
type ConversionRateTieredConfig = shared.ConversionRateTieredConfig

// This is an alias to an internal type.
type ConversionRateTieredConfigParam = shared.ConversionRateTieredConfigParam

// This is an alias to an internal type.
type ConversionRateUnitConfig = shared.ConversionRateUnitConfig

// This is an alias to an internal type.
type ConversionRateUnitConfigParam = shared.ConversionRateUnitConfigParam

// This is an alias to an internal type.
type CouponRedemption = shared.CouponRedemption

// The [Credit Note](/invoicing/credit-notes) resource represents a credit that has
// been applied to a particular invoice.
//
// This is an alias to an internal type.
type CreditNote = shared.CreditNote

// This is an alias to an internal type.
type CreditNoteLineItem = shared.CreditNoteLineItem

// This is an alias to an internal type.
type CreditNoteLineItemsDiscount = shared.CreditNoteLineItemsDiscount

// This is an alias to an internal type.
type CreditNoteLineItemsDiscountsDiscountType = shared.CreditNoteLineItemsDiscountsDiscountType

// This is an alias to an internal value.
const CreditNoteLineItemsDiscountsDiscountTypePercentage = shared.CreditNoteLineItemsDiscountsDiscountTypePercentage

// This is an alias to an internal value.
const CreditNoteLineItemsDiscountsDiscountTypeAmount = shared.CreditNoteLineItemsDiscountsDiscountTypeAmount

// The maximum amount applied on the original invoice
//
// This is an alias to an internal type.
type CreditNoteMaximumAmountAdjustment = shared.CreditNoteMaximumAmountAdjustment

// This is an alias to an internal type.
type CreditNoteMaximumAmountAdjustmentDiscountType = shared.CreditNoteMaximumAmountAdjustmentDiscountType

// This is an alias to an internal value.
const CreditNoteMaximumAmountAdjustmentDiscountTypePercentage = shared.CreditNoteMaximumAmountAdjustmentDiscountTypePercentage

// This is an alias to an internal type.
type CreditNoteMaximumAmountAdjustmentAppliesToPrice = shared.CreditNoteMaximumAmountAdjustmentAppliesToPrice

// This is an alias to an internal type.
type CreditNoteReason = shared.CreditNoteReason

// This is an alias to an internal value.
const CreditNoteReasonDuplicate = shared.CreditNoteReasonDuplicate

// This is an alias to an internal value.
const CreditNoteReasonFraudulent = shared.CreditNoteReasonFraudulent

// This is an alias to an internal value.
const CreditNoteReasonOrderChange = shared.CreditNoteReasonOrderChange

// This is an alias to an internal value.
const CreditNoteReasonProductUnsatisfactory = shared.CreditNoteReasonProductUnsatisfactory

// This is an alias to an internal type.
type CreditNoteType = shared.CreditNoteType

// This is an alias to an internal value.
const CreditNoteTypeRefund = shared.CreditNoteTypeRefund

// This is an alias to an internal value.
const CreditNoteTypeAdjustment = shared.CreditNoteTypeAdjustment

// This is an alias to an internal type.
type CreditNoteDiscount = shared.CreditNoteDiscount

// This is an alias to an internal type.
type CreditNoteDiscountsDiscountType = shared.CreditNoteDiscountsDiscountType

// This is an alias to an internal value.
const CreditNoteDiscountsDiscountTypePercentage = shared.CreditNoteDiscountsDiscountTypePercentage

// This is an alias to an internal type.
type CreditNoteDiscountsAppliesToPrice = shared.CreditNoteDiscountsAppliesToPrice

// This is an alias to an internal type.
type CreditNoteTiny = shared.CreditNoteTiny

// This is an alias to an internal type.
type CustomExpiration = shared.CustomExpiration

// This is an alias to an internal type.
type CustomExpirationDurationUnit = shared.CustomExpirationDurationUnit

// This is an alias to an internal value.
const CustomExpirationDurationUnitDay = shared.CustomExpirationDurationUnitDay

// This is an alias to an internal value.
const CustomExpirationDurationUnitMonth = shared.CustomExpirationDurationUnitMonth

// This is an alias to an internal type.
type CustomExpirationParam = shared.CustomExpirationParam

// This is an alias to an internal type.
type CustomerMinified = shared.CustomerMinified

// Tax IDs are commonly required to be displayed on customer invoices, which are
// added to the headers of invoices.
//
// ### Supported Tax ID Countries and Types
//
// | Country                | Type         | Description                                                                                             |
// | ---------------------- | ------------ | ------------------------------------------------------------------------------------------------------- |
// | Albania                | `al_tin`     | Albania Tax Identification Number                                                                       |
// | Andorra                | `ad_nrt`     | Andorran NRT Number                                                                                     |
// | Angola                 | `ao_tin`     | Angola Tax Identification Number                                                                        |
// | Argentina              | `ar_cuit`    | Argentinian Tax ID Number                                                                               |
// | Armenia                | `am_tin`     | Armenia Tax Identification Number                                                                       |
// | Aruba                  | `aw_tin`     | Aruba Tax Identification Number                                                                         |
// | Australia              | `au_abn`     | Australian Business Number (AU ABN)                                                                     |
// | Australia              | `au_arn`     | Australian Taxation Office Reference Number                                                             |
// | Austria                | `eu_vat`     | European VAT Number                                                                                     |
// | Azerbaijan             | `az_tin`     | Azerbaijan Tax Identification Number                                                                    |
// | Bahamas                | `bs_tin`     | Bahamas Tax Identification Number                                                                       |
// | Bahrain                | `bh_vat`     | Bahraini VAT Number                                                                                     |
// | Bangladesh             | `bd_bin`     | Bangladesh Business Identification Number                                                               |
// | Barbados               | `bb_tin`     | Barbados Tax Identification Number                                                                      |
// | Belarus                | `by_tin`     | Belarus TIN Number                                                                                      |
// | Belgium                | `eu_vat`     | European VAT Number                                                                                     |
// | Benin                  | `bj_ifu`     | Benin Tax Identification Number (Identifiant Fiscal Unique)                                             |
// | Bolivia                | `bo_tin`     | Bolivian Tax ID                                                                                         |
// | Bosnia and Herzegovina | `ba_tin`     | Bosnia and Herzegovina Tax Identification Number                                                        |
// | Brazil                 | `br_cnpj`    | Brazilian CNPJ Number                                                                                   |
// | Brazil                 | `br_cpf`     | Brazilian CPF Number                                                                                    |
// | Bulgaria               | `bg_uic`     | Bulgaria Unified Identification Code                                                                    |
// | Bulgaria               | `eu_vat`     | European VAT Number                                                                                     |
// | Burkina Faso           | `bf_ifu`     | Burkina Faso Tax Identification Number (Numéro d'Identifiant Fiscal Unique)                             |
// | Cambodia               | `kh_tin`     | Cambodia Tax Identification Number                                                                      |
// | Cameroon               | `cm_niu`     | Cameroon Tax Identification Number (Numéro d'Identifiant fiscal Unique)                                 |
// | Canada                 | `ca_bn`      | Canadian BN                                                                                             |
// | Canada                 | `ca_gst_hst` | Canadian GST/HST Number                                                                                 |
// | Canada                 | `ca_pst_bc`  | Canadian PST Number (British Columbia)                                                                  |
// | Canada                 | `ca_pst_mb`  | Canadian PST Number (Manitoba)                                                                          |
// | Canada                 | `ca_pst_sk`  | Canadian PST Number (Saskatchewan)                                                                      |
// | Canada                 | `ca_qst`     | Canadian QST Number (Québec)                                                                            |
// | Cape Verde             | `cv_nif`     | Cape Verde Tax Identification Number (Número de Identificação Fiscal)                                   |
// | Chile                  | `cl_tin`     | Chilean TIN                                                                                             |
// | China                  | `cn_tin`     | Chinese Tax ID                                                                                          |
// | Colombia               | `co_nit`     | Colombian NIT Number                                                                                    |
// | Congo-Kinshasa         | `cd_nif`     | Congo (DR) Tax Identification Number (Número de Identificação Fiscal)                                   |
// | Costa Rica             | `cr_tin`     | Costa Rican Tax ID                                                                                      |
// | Croatia                | `eu_vat`     | European VAT Number                                                                                     |
// | Croatia                | `hr_oib`     | Croatian Personal Identification Number (OIB)                                                           |
// | Cyprus                 | `eu_vat`     | European VAT Number                                                                                     |
// | Czech Republic         | `eu_vat`     | European VAT Number                                                                                     |
// | Denmark                | `eu_vat`     | European VAT Number                                                                                     |
// | Dominican Republic     | `do_rcn`     | Dominican RCN Number                                                                                    |
// | Ecuador                | `ec_ruc`     | Ecuadorian RUC Number                                                                                   |
// | Egypt                  | `eg_tin`     | Egyptian Tax Identification Number                                                                      |
// | El Salvador            | `sv_nit`     | El Salvadorian NIT Number                                                                               |
// | Estonia                | `eu_vat`     | European VAT Number                                                                                     |
// | Ethiopia               | `et_tin`     | Ethiopia Tax Identification Number                                                                      |
// | European Union         | `eu_oss_vat` | European One Stop Shop VAT Number for non-Union scheme                                                  |
// | Finland                | `eu_vat`     | European VAT Number                                                                                     |
// | France                 | `eu_vat`     | European VAT Number                                                                                     |
// | Georgia                | `ge_vat`     | Georgian VAT                                                                                            |
// | Germany                | `de_stn`     | German Tax Number (Steuernummer)                                                                        |
// | Germany                | `eu_vat`     | European VAT Number                                                                                     |
// | Greece                 | `eu_vat`     | European VAT Number                                                                                     |
// | Guinea                 | `gn_nif`     | Guinea Tax Identification Number (Número de Identificação Fiscal)                                       |
// | Hong Kong              | `hk_br`      | Hong Kong BR Number                                                                                     |
// | Hungary                | `eu_vat`     | European VAT Number                                                                                     |
// | Hungary                | `hu_tin`     | Hungary Tax Number (adószám)                                                                            |
// | Iceland                | `is_vat`     | Icelandic VAT                                                                                           |
// | India                  | `in_gst`     | Indian GST Number                                                                                       |
// | Indonesia              | `id_npwp`    | Indonesian NPWP Number                                                                                  |
// | Ireland                | `eu_vat`     | European VAT Number                                                                                     |
// | Israel                 | `il_vat`     | Israel VAT                                                                                              |
// | Italy                  | `eu_vat`     | European VAT Number                                                                                     |
// | Japan                  | `jp_cn`      | Japanese Corporate Number (_Hōjin Bangō_)                                                               |
// | Japan                  | `jp_rn`      | Japanese Registered Foreign Businesses' Registration Number (_Tōroku Kokugai Jigyōsha no Tōroku Bangō_) |
// | Japan                  | `jp_trn`     | Japanese Tax Registration Number (_Tōroku Bangō_)                                                       |
// | Kazakhstan             | `kz_bin`     | Kazakhstani Business Identification Number                                                              |
// | Kenya                  | `ke_pin`     | Kenya Revenue Authority Personal Identification Number                                                  |
// | Kyrgyzstan             | `kg_tin`     | Kyrgyzstan Tax Identification Number                                                                    |
// | Laos                   | `la_tin`     | Laos Tax Identification Number                                                                          |
// | Latvia                 | `eu_vat`     | European VAT Number                                                                                     |
// | Liechtenstein          | `li_uid`     | Liechtensteinian UID Number                                                                             |
// | Liechtenstein          | `li_vat`     | Liechtenstein VAT Number                                                                                |
// | Lithuania              | `eu_vat`     | European VAT Number                                                                                     |
// | Luxembourg             | `eu_vat`     | European VAT Number                                                                                     |
// | Malaysia               | `my_frp`     | Malaysian FRP Number                                                                                    |
// | Malaysia               | `my_itn`     | Malaysian ITN                                                                                           |
// | Malaysia               | `my_sst`     | Malaysian SST Number                                                                                    |
// | Malta                  | `eu_vat`     | European VAT Number                                                                                     |
// | Mauritania             | `mr_nif`     | Mauritania Tax Identification Number (Número de Identificação Fiscal)                                   |
// | Mexico                 | `mx_rfc`     | Mexican RFC Number                                                                                      |
// | Moldova                | `md_vat`     | Moldova VAT Number                                                                                      |
// | Montenegro             | `me_pib`     | Montenegro PIB Number                                                                                   |
// | Morocco                | `ma_vat`     | Morocco VAT Number                                                                                      |
// | Nepal                  | `np_pan`     | Nepal PAN Number                                                                                        |
// | Netherlands            | `eu_vat`     | European VAT Number                                                                                     |
// | New Zealand            | `nz_gst`     | New Zealand GST Number                                                                                  |
// | Nigeria                | `ng_tin`     | Nigerian Tax Identification Number                                                                      |
// | North Macedonia        | `mk_vat`     | North Macedonia VAT Number                                                                              |
// | Northern Ireland       | `eu_vat`     | Northern Ireland VAT Number                                                                             |
// | Norway                 | `no_vat`     | Norwegian VAT Number                                                                                    |
// | Norway                 | `no_voec`    | Norwegian VAT on e-commerce Number                                                                      |
// | Oman                   | `om_vat`     | Omani VAT Number                                                                                        |
// | Peru                   | `pe_ruc`     | Peruvian RUC Number                                                                                     |
// | Philippines            | `ph_tin`     | Philippines Tax Identification Number                                                                   |
// | Poland                 | `eu_vat`     | European VAT Number                                                                                     |
// | Poland                 | `pl_nip`     | Polish Tax ID Number                                                                                    |
// | Portugal               | `eu_vat`     | European VAT Number                                                                                     |
// | Romania                | `eu_vat`     | European VAT Number                                                                                     |
// | Romania                | `ro_tin`     | Romanian Tax ID Number                                                                                  |
// | Russia                 | `ru_inn`     | Russian INN                                                                                             |
// | Russia                 | `ru_kpp`     | Russian KPP                                                                                             |
// | Saudi Arabia           | `sa_vat`     | Saudi Arabia VAT                                                                                        |
// | Senegal                | `sn_ninea`   | Senegal NINEA Number                                                                                    |
// | Serbia                 | `rs_pib`     | Serbian PIB Number                                                                                      |
// | Singapore              | `sg_gst`     | Singaporean GST                                                                                         |
// | Singapore              | `sg_uen`     | Singaporean UEN                                                                                         |
// | Slovakia               | `eu_vat`     | European VAT Number                                                                                     |
// | Slovenia               | `eu_vat`     | European VAT Number                                                                                     |
// | Slovenia               | `si_tin`     | Slovenia Tax Number (davčna številka)                                                                   |
// | South Africa           | `za_vat`     | South African VAT Number                                                                                |
// | South Korea            | `kr_brn`     | Korean BRN                                                                                              |
// | Spain                  | `es_cif`     | Spanish NIF Number (previously Spanish CIF Number)                                                      |
// | Spain                  | `eu_vat`     | European VAT Number                                                                                     |
// | Suriname               | `sr_fin`     | Suriname FIN Number                                                                                     |
// | Sweden                 | `eu_vat`     | European VAT Number                                                                                     |
// | Switzerland            | `ch_uid`     | Switzerland UID Number                                                                                  |
// | Switzerland            | `ch_vat`     | Switzerland VAT Number                                                                                  |
// | Taiwan                 | `tw_vat`     | Taiwanese VAT                                                                                           |
// | Tajikistan             | `tj_tin`     | Tajikistan Tax Identification Number                                                                    |
// | Tanzania               | `tz_vat`     | Tanzania VAT Number                                                                                     |
// | Thailand               | `th_vat`     | Thai VAT                                                                                                |
// | Turkey                 | `tr_tin`     | Turkish Tax Identification Number                                                                       |
// | Uganda                 | `ug_tin`     | Uganda Tax Identification Number                                                                        |
// | Ukraine                | `ua_vat`     | Ukrainian VAT                                                                                           |
// | United Arab Emirates   | `ae_trn`     | United Arab Emirates TRN                                                                                |
// | United Kingdom         | `gb_vat`     | United Kingdom VAT Number                                                                               |
// | United States          | `us_ein`     | United States EIN                                                                                       |
// | Uruguay                | `uy_ruc`     | Uruguayan RUC Number                                                                                    |
// | Uzbekistan             | `uz_tin`     | Uzbekistan TIN Number                                                                                   |
// | Uzbekistan             | `uz_vat`     | Uzbekistan VAT Number                                                                                   |
// | Venezuela              | `ve_rif`     | Venezuelan RIF Number                                                                                   |
// | Vietnam                | `vn_tin`     | Vietnamese Tax ID Number                                                                                |
// | Zambia                 | `zm_tin`     | Zambia Tax Identification Number                                                                        |
// | Zimbabwe               | `zw_tin`     | Zimbabwe Tax Identification Number                                                                      |
//
// This is an alias to an internal type.
type CustomerTaxID = shared.CustomerTaxID

// This is an alias to an internal type.
type CustomerTaxIDCountry = shared.CustomerTaxIDCountry

// This is an alias to an internal value.
const CustomerTaxIDCountryAd = shared.CustomerTaxIDCountryAd

// This is an alias to an internal value.
const CustomerTaxIDCountryAe = shared.CustomerTaxIDCountryAe

// This is an alias to an internal value.
const CustomerTaxIDCountryAl = shared.CustomerTaxIDCountryAl

// This is an alias to an internal value.
const CustomerTaxIDCountryAm = shared.CustomerTaxIDCountryAm

// This is an alias to an internal value.
const CustomerTaxIDCountryAo = shared.CustomerTaxIDCountryAo

// This is an alias to an internal value.
const CustomerTaxIDCountryAr = shared.CustomerTaxIDCountryAr

// This is an alias to an internal value.
const CustomerTaxIDCountryAt = shared.CustomerTaxIDCountryAt

// This is an alias to an internal value.
const CustomerTaxIDCountryAu = shared.CustomerTaxIDCountryAu

// This is an alias to an internal value.
const CustomerTaxIDCountryAw = shared.CustomerTaxIDCountryAw

// This is an alias to an internal value.
const CustomerTaxIDCountryAz = shared.CustomerTaxIDCountryAz

// This is an alias to an internal value.
const CustomerTaxIDCountryBa = shared.CustomerTaxIDCountryBa

// This is an alias to an internal value.
const CustomerTaxIDCountryBb = shared.CustomerTaxIDCountryBb

// This is an alias to an internal value.
const CustomerTaxIDCountryBd = shared.CustomerTaxIDCountryBd

// This is an alias to an internal value.
const CustomerTaxIDCountryBe = shared.CustomerTaxIDCountryBe

// This is an alias to an internal value.
const CustomerTaxIDCountryBf = shared.CustomerTaxIDCountryBf

// This is an alias to an internal value.
const CustomerTaxIDCountryBg = shared.CustomerTaxIDCountryBg

// This is an alias to an internal value.
const CustomerTaxIDCountryBh = shared.CustomerTaxIDCountryBh

// This is an alias to an internal value.
const CustomerTaxIDCountryBj = shared.CustomerTaxIDCountryBj

// This is an alias to an internal value.
const CustomerTaxIDCountryBo = shared.CustomerTaxIDCountryBo

// This is an alias to an internal value.
const CustomerTaxIDCountryBr = shared.CustomerTaxIDCountryBr

// This is an alias to an internal value.
const CustomerTaxIDCountryBs = shared.CustomerTaxIDCountryBs

// This is an alias to an internal value.
const CustomerTaxIDCountryBy = shared.CustomerTaxIDCountryBy

// This is an alias to an internal value.
const CustomerTaxIDCountryCa = shared.CustomerTaxIDCountryCa

// This is an alias to an internal value.
const CustomerTaxIDCountryCd = shared.CustomerTaxIDCountryCd

// This is an alias to an internal value.
const CustomerTaxIDCountryCh = shared.CustomerTaxIDCountryCh

// This is an alias to an internal value.
const CustomerTaxIDCountryCl = shared.CustomerTaxIDCountryCl

// This is an alias to an internal value.
const CustomerTaxIDCountryCm = shared.CustomerTaxIDCountryCm

// This is an alias to an internal value.
const CustomerTaxIDCountryCn = shared.CustomerTaxIDCountryCn

// This is an alias to an internal value.
const CustomerTaxIDCountryCo = shared.CustomerTaxIDCountryCo

// This is an alias to an internal value.
const CustomerTaxIDCountryCr = shared.CustomerTaxIDCountryCr

// This is an alias to an internal value.
const CustomerTaxIDCountryCv = shared.CustomerTaxIDCountryCv

// This is an alias to an internal value.
const CustomerTaxIDCountryDe = shared.CustomerTaxIDCountryDe

// This is an alias to an internal value.
const CustomerTaxIDCountryCy = shared.CustomerTaxIDCountryCy

// This is an alias to an internal value.
const CustomerTaxIDCountryCz = shared.CustomerTaxIDCountryCz

// This is an alias to an internal value.
const CustomerTaxIDCountryDk = shared.CustomerTaxIDCountryDk

// This is an alias to an internal value.
const CustomerTaxIDCountryDo = shared.CustomerTaxIDCountryDo

// This is an alias to an internal value.
const CustomerTaxIDCountryEc = shared.CustomerTaxIDCountryEc

// This is an alias to an internal value.
const CustomerTaxIDCountryEe = shared.CustomerTaxIDCountryEe

// This is an alias to an internal value.
const CustomerTaxIDCountryEg = shared.CustomerTaxIDCountryEg

// This is an alias to an internal value.
const CustomerTaxIDCountryEs = shared.CustomerTaxIDCountryEs

// This is an alias to an internal value.
const CustomerTaxIDCountryEt = shared.CustomerTaxIDCountryEt

// This is an alias to an internal value.
const CustomerTaxIDCountryEu = shared.CustomerTaxIDCountryEu

// This is an alias to an internal value.
const CustomerTaxIDCountryFi = shared.CustomerTaxIDCountryFi

// This is an alias to an internal value.
const CustomerTaxIDCountryFr = shared.CustomerTaxIDCountryFr

// This is an alias to an internal value.
const CustomerTaxIDCountryGB = shared.CustomerTaxIDCountryGB

// This is an alias to an internal value.
const CustomerTaxIDCountryGe = shared.CustomerTaxIDCountryGe

// This is an alias to an internal value.
const CustomerTaxIDCountryGn = shared.CustomerTaxIDCountryGn

// This is an alias to an internal value.
const CustomerTaxIDCountryGr = shared.CustomerTaxIDCountryGr

// This is an alias to an internal value.
const CustomerTaxIDCountryHk = shared.CustomerTaxIDCountryHk

// This is an alias to an internal value.
const CustomerTaxIDCountryHr = shared.CustomerTaxIDCountryHr

// This is an alias to an internal value.
const CustomerTaxIDCountryHu = shared.CustomerTaxIDCountryHu

// This is an alias to an internal value.
const CustomerTaxIDCountryID = shared.CustomerTaxIDCountryID

// This is an alias to an internal value.
const CustomerTaxIDCountryIe = shared.CustomerTaxIDCountryIe

// This is an alias to an internal value.
const CustomerTaxIDCountryIl = shared.CustomerTaxIDCountryIl

// This is an alias to an internal value.
const CustomerTaxIDCountryIn = shared.CustomerTaxIDCountryIn

// This is an alias to an internal value.
const CustomerTaxIDCountryIs = shared.CustomerTaxIDCountryIs

// This is an alias to an internal value.
const CustomerTaxIDCountryIt = shared.CustomerTaxIDCountryIt

// This is an alias to an internal value.
const CustomerTaxIDCountryJp = shared.CustomerTaxIDCountryJp

// This is an alias to an internal value.
const CustomerTaxIDCountryKe = shared.CustomerTaxIDCountryKe

// This is an alias to an internal value.
const CustomerTaxIDCountryKg = shared.CustomerTaxIDCountryKg

// This is an alias to an internal value.
const CustomerTaxIDCountryKh = shared.CustomerTaxIDCountryKh

// This is an alias to an internal value.
const CustomerTaxIDCountryKr = shared.CustomerTaxIDCountryKr

// This is an alias to an internal value.
const CustomerTaxIDCountryKz = shared.CustomerTaxIDCountryKz

// This is an alias to an internal value.
const CustomerTaxIDCountryLa = shared.CustomerTaxIDCountryLa

// This is an alias to an internal value.
const CustomerTaxIDCountryLi = shared.CustomerTaxIDCountryLi

// This is an alias to an internal value.
const CustomerTaxIDCountryLt = shared.CustomerTaxIDCountryLt

// This is an alias to an internal value.
const CustomerTaxIDCountryLu = shared.CustomerTaxIDCountryLu

// This is an alias to an internal value.
const CustomerTaxIDCountryLv = shared.CustomerTaxIDCountryLv

// This is an alias to an internal value.
const CustomerTaxIDCountryMa = shared.CustomerTaxIDCountryMa

// This is an alias to an internal value.
const CustomerTaxIDCountryMd = shared.CustomerTaxIDCountryMd

// This is an alias to an internal value.
const CustomerTaxIDCountryMe = shared.CustomerTaxIDCountryMe

// This is an alias to an internal value.
const CustomerTaxIDCountryMk = shared.CustomerTaxIDCountryMk

// This is an alias to an internal value.
const CustomerTaxIDCountryMr = shared.CustomerTaxIDCountryMr

// This is an alias to an internal value.
const CustomerTaxIDCountryMt = shared.CustomerTaxIDCountryMt

// This is an alias to an internal value.
const CustomerTaxIDCountryMx = shared.CustomerTaxIDCountryMx

// This is an alias to an internal value.
const CustomerTaxIDCountryMy = shared.CustomerTaxIDCountryMy

// This is an alias to an internal value.
const CustomerTaxIDCountryNg = shared.CustomerTaxIDCountryNg

// This is an alias to an internal value.
const CustomerTaxIDCountryNl = shared.CustomerTaxIDCountryNl

// This is an alias to an internal value.
const CustomerTaxIDCountryNo = shared.CustomerTaxIDCountryNo

// This is an alias to an internal value.
const CustomerTaxIDCountryNp = shared.CustomerTaxIDCountryNp

// This is an alias to an internal value.
const CustomerTaxIDCountryNz = shared.CustomerTaxIDCountryNz

// This is an alias to an internal value.
const CustomerTaxIDCountryOm = shared.CustomerTaxIDCountryOm

// This is an alias to an internal value.
const CustomerTaxIDCountryPe = shared.CustomerTaxIDCountryPe

// This is an alias to an internal value.
const CustomerTaxIDCountryPh = shared.CustomerTaxIDCountryPh

// This is an alias to an internal value.
const CustomerTaxIDCountryPl = shared.CustomerTaxIDCountryPl

// This is an alias to an internal value.
const CustomerTaxIDCountryPt = shared.CustomerTaxIDCountryPt

// This is an alias to an internal value.
const CustomerTaxIDCountryRo = shared.CustomerTaxIDCountryRo

// This is an alias to an internal value.
const CustomerTaxIDCountryRs = shared.CustomerTaxIDCountryRs

// This is an alias to an internal value.
const CustomerTaxIDCountryRu = shared.CustomerTaxIDCountryRu

// This is an alias to an internal value.
const CustomerTaxIDCountrySa = shared.CustomerTaxIDCountrySa

// This is an alias to an internal value.
const CustomerTaxIDCountrySe = shared.CustomerTaxIDCountrySe

// This is an alias to an internal value.
const CustomerTaxIDCountrySg = shared.CustomerTaxIDCountrySg

// This is an alias to an internal value.
const CustomerTaxIDCountrySi = shared.CustomerTaxIDCountrySi

// This is an alias to an internal value.
const CustomerTaxIDCountrySk = shared.CustomerTaxIDCountrySk

// This is an alias to an internal value.
const CustomerTaxIDCountrySn = shared.CustomerTaxIDCountrySn

// This is an alias to an internal value.
const CustomerTaxIDCountrySr = shared.CustomerTaxIDCountrySr

// This is an alias to an internal value.
const CustomerTaxIDCountrySv = shared.CustomerTaxIDCountrySv

// This is an alias to an internal value.
const CustomerTaxIDCountryTh = shared.CustomerTaxIDCountryTh

// This is an alias to an internal value.
const CustomerTaxIDCountryTj = shared.CustomerTaxIDCountryTj

// This is an alias to an internal value.
const CustomerTaxIDCountryTr = shared.CustomerTaxIDCountryTr

// This is an alias to an internal value.
const CustomerTaxIDCountryTw = shared.CustomerTaxIDCountryTw

// This is an alias to an internal value.
const CustomerTaxIDCountryTz = shared.CustomerTaxIDCountryTz

// This is an alias to an internal value.
const CustomerTaxIDCountryUa = shared.CustomerTaxIDCountryUa

// This is an alias to an internal value.
const CustomerTaxIDCountryUg = shared.CustomerTaxIDCountryUg

// This is an alias to an internal value.
const CustomerTaxIDCountryUs = shared.CustomerTaxIDCountryUs

// This is an alias to an internal value.
const CustomerTaxIDCountryUy = shared.CustomerTaxIDCountryUy

// This is an alias to an internal value.
const CustomerTaxIDCountryUz = shared.CustomerTaxIDCountryUz

// This is an alias to an internal value.
const CustomerTaxIDCountryVe = shared.CustomerTaxIDCountryVe

// This is an alias to an internal value.
const CustomerTaxIDCountryVn = shared.CustomerTaxIDCountryVn

// This is an alias to an internal value.
const CustomerTaxIDCountryZa = shared.CustomerTaxIDCountryZa

// This is an alias to an internal value.
const CustomerTaxIDCountryZm = shared.CustomerTaxIDCountryZm

// This is an alias to an internal value.
const CustomerTaxIDCountryZw = shared.CustomerTaxIDCountryZw

// This is an alias to an internal type.
type CustomerTaxIDType = shared.CustomerTaxIDType

// This is an alias to an internal value.
const CustomerTaxIDTypeAdNrt = shared.CustomerTaxIDTypeAdNrt

// This is an alias to an internal value.
const CustomerTaxIDTypeAeTrn = shared.CustomerTaxIDTypeAeTrn

// This is an alias to an internal value.
const CustomerTaxIDTypeAlTin = shared.CustomerTaxIDTypeAlTin

// This is an alias to an internal value.
const CustomerTaxIDTypeAmTin = shared.CustomerTaxIDTypeAmTin

// This is an alias to an internal value.
const CustomerTaxIDTypeAoTin = shared.CustomerTaxIDTypeAoTin

// This is an alias to an internal value.
const CustomerTaxIDTypeArCuit = shared.CustomerTaxIDTypeArCuit

// This is an alias to an internal value.
const CustomerTaxIDTypeEuVat = shared.CustomerTaxIDTypeEuVat

// This is an alias to an internal value.
const CustomerTaxIDTypeAuAbn = shared.CustomerTaxIDTypeAuAbn

// This is an alias to an internal value.
const CustomerTaxIDTypeAuArn = shared.CustomerTaxIDTypeAuArn

// This is an alias to an internal value.
const CustomerTaxIDTypeAwTin = shared.CustomerTaxIDTypeAwTin

// This is an alias to an internal value.
const CustomerTaxIDTypeAzTin = shared.CustomerTaxIDTypeAzTin

// This is an alias to an internal value.
const CustomerTaxIDTypeBaTin = shared.CustomerTaxIDTypeBaTin

// This is an alias to an internal value.
const CustomerTaxIDTypeBbTin = shared.CustomerTaxIDTypeBbTin

// This is an alias to an internal value.
const CustomerTaxIDTypeBdBin = shared.CustomerTaxIDTypeBdBin

// This is an alias to an internal value.
const CustomerTaxIDTypeBfIfu = shared.CustomerTaxIDTypeBfIfu

// This is an alias to an internal value.
const CustomerTaxIDTypeBgUic = shared.CustomerTaxIDTypeBgUic

// This is an alias to an internal value.
const CustomerTaxIDTypeBhVat = shared.CustomerTaxIDTypeBhVat

// This is an alias to an internal value.
const CustomerTaxIDTypeBjIfu = shared.CustomerTaxIDTypeBjIfu

// This is an alias to an internal value.
const CustomerTaxIDTypeBoTin = shared.CustomerTaxIDTypeBoTin

// This is an alias to an internal value.
const CustomerTaxIDTypeBrCnpj = shared.CustomerTaxIDTypeBrCnpj

// This is an alias to an internal value.
const CustomerTaxIDTypeBrCpf = shared.CustomerTaxIDTypeBrCpf

// This is an alias to an internal value.
const CustomerTaxIDTypeBsTin = shared.CustomerTaxIDTypeBsTin

// This is an alias to an internal value.
const CustomerTaxIDTypeByTin = shared.CustomerTaxIDTypeByTin

// This is an alias to an internal value.
const CustomerTaxIDTypeCaBn = shared.CustomerTaxIDTypeCaBn

// This is an alias to an internal value.
const CustomerTaxIDTypeCaGstHst = shared.CustomerTaxIDTypeCaGstHst

// This is an alias to an internal value.
const CustomerTaxIDTypeCaPstBc = shared.CustomerTaxIDTypeCaPstBc

// This is an alias to an internal value.
const CustomerTaxIDTypeCaPstMB = shared.CustomerTaxIDTypeCaPstMB

// This is an alias to an internal value.
const CustomerTaxIDTypeCaPstSk = shared.CustomerTaxIDTypeCaPstSk

// This is an alias to an internal value.
const CustomerTaxIDTypeCaQst = shared.CustomerTaxIDTypeCaQst

// This is an alias to an internal value.
const CustomerTaxIDTypeCdNif = shared.CustomerTaxIDTypeCdNif

// This is an alias to an internal value.
const CustomerTaxIDTypeChUid = shared.CustomerTaxIDTypeChUid

// This is an alias to an internal value.
const CustomerTaxIDTypeChVat = shared.CustomerTaxIDTypeChVat

// This is an alias to an internal value.
const CustomerTaxIDTypeClTin = shared.CustomerTaxIDTypeClTin

// This is an alias to an internal value.
const CustomerTaxIDTypeCmNiu = shared.CustomerTaxIDTypeCmNiu

// This is an alias to an internal value.
const CustomerTaxIDTypeCnTin = shared.CustomerTaxIDTypeCnTin

// This is an alias to an internal value.
const CustomerTaxIDTypeCoNit = shared.CustomerTaxIDTypeCoNit

// This is an alias to an internal value.
const CustomerTaxIDTypeCrTin = shared.CustomerTaxIDTypeCrTin

// This is an alias to an internal value.
const CustomerTaxIDTypeCvNif = shared.CustomerTaxIDTypeCvNif

// This is an alias to an internal value.
const CustomerTaxIDTypeDeStn = shared.CustomerTaxIDTypeDeStn

// This is an alias to an internal value.
const CustomerTaxIDTypeDoRcn = shared.CustomerTaxIDTypeDoRcn

// This is an alias to an internal value.
const CustomerTaxIDTypeEcRuc = shared.CustomerTaxIDTypeEcRuc

// This is an alias to an internal value.
const CustomerTaxIDTypeEgTin = shared.CustomerTaxIDTypeEgTin

// This is an alias to an internal value.
const CustomerTaxIDTypeEsCif = shared.CustomerTaxIDTypeEsCif

// This is an alias to an internal value.
const CustomerTaxIDTypeEtTin = shared.CustomerTaxIDTypeEtTin

// This is an alias to an internal value.
const CustomerTaxIDTypeEuOssVat = shared.CustomerTaxIDTypeEuOssVat

// This is an alias to an internal value.
const CustomerTaxIDTypeGBVat = shared.CustomerTaxIDTypeGBVat

// This is an alias to an internal value.
const CustomerTaxIDTypeGeVat = shared.CustomerTaxIDTypeGeVat

// This is an alias to an internal value.
const CustomerTaxIDTypeGnNif = shared.CustomerTaxIDTypeGnNif

// This is an alias to an internal value.
const CustomerTaxIDTypeHkBr = shared.CustomerTaxIDTypeHkBr

// This is an alias to an internal value.
const CustomerTaxIDTypeHrOib = shared.CustomerTaxIDTypeHrOib

// This is an alias to an internal value.
const CustomerTaxIDTypeHuTin = shared.CustomerTaxIDTypeHuTin

// This is an alias to an internal value.
const CustomerTaxIDTypeIDNpwp = shared.CustomerTaxIDTypeIDNpwp

// This is an alias to an internal value.
const CustomerTaxIDTypeIlVat = shared.CustomerTaxIDTypeIlVat

// This is an alias to an internal value.
const CustomerTaxIDTypeInGst = shared.CustomerTaxIDTypeInGst

// This is an alias to an internal value.
const CustomerTaxIDTypeIsVat = shared.CustomerTaxIDTypeIsVat

// This is an alias to an internal value.
const CustomerTaxIDTypeJpCn = shared.CustomerTaxIDTypeJpCn

// This is an alias to an internal value.
const CustomerTaxIDTypeJpRn = shared.CustomerTaxIDTypeJpRn

// This is an alias to an internal value.
const CustomerTaxIDTypeJpTrn = shared.CustomerTaxIDTypeJpTrn

// This is an alias to an internal value.
const CustomerTaxIDTypeKePin = shared.CustomerTaxIDTypeKePin

// This is an alias to an internal value.
const CustomerTaxIDTypeKgTin = shared.CustomerTaxIDTypeKgTin

// This is an alias to an internal value.
const CustomerTaxIDTypeKhTin = shared.CustomerTaxIDTypeKhTin

// This is an alias to an internal value.
const CustomerTaxIDTypeKrBrn = shared.CustomerTaxIDTypeKrBrn

// This is an alias to an internal value.
const CustomerTaxIDTypeKzBin = shared.CustomerTaxIDTypeKzBin

// This is an alias to an internal value.
const CustomerTaxIDTypeLaTin = shared.CustomerTaxIDTypeLaTin

// This is an alias to an internal value.
const CustomerTaxIDTypeLiUid = shared.CustomerTaxIDTypeLiUid

// This is an alias to an internal value.
const CustomerTaxIDTypeLiVat = shared.CustomerTaxIDTypeLiVat

// This is an alias to an internal value.
const CustomerTaxIDTypeMaVat = shared.CustomerTaxIDTypeMaVat

// This is an alias to an internal value.
const CustomerTaxIDTypeMdVat = shared.CustomerTaxIDTypeMdVat

// This is an alias to an internal value.
const CustomerTaxIDTypeMePib = shared.CustomerTaxIDTypeMePib

// This is an alias to an internal value.
const CustomerTaxIDTypeMkVat = shared.CustomerTaxIDTypeMkVat

// This is an alias to an internal value.
const CustomerTaxIDTypeMrNif = shared.CustomerTaxIDTypeMrNif

// This is an alias to an internal value.
const CustomerTaxIDTypeMxRfc = shared.CustomerTaxIDTypeMxRfc

// This is an alias to an internal value.
const CustomerTaxIDTypeMyFrp = shared.CustomerTaxIDTypeMyFrp

// This is an alias to an internal value.
const CustomerTaxIDTypeMyItn = shared.CustomerTaxIDTypeMyItn

// This is an alias to an internal value.
const CustomerTaxIDTypeMySst = shared.CustomerTaxIDTypeMySst

// This is an alias to an internal value.
const CustomerTaxIDTypeNgTin = shared.CustomerTaxIDTypeNgTin

// This is an alias to an internal value.
const CustomerTaxIDTypeNoVat = shared.CustomerTaxIDTypeNoVat

// This is an alias to an internal value.
const CustomerTaxIDTypeNoVoec = shared.CustomerTaxIDTypeNoVoec

// This is an alias to an internal value.
const CustomerTaxIDTypeNpPan = shared.CustomerTaxIDTypeNpPan

// This is an alias to an internal value.
const CustomerTaxIDTypeNzGst = shared.CustomerTaxIDTypeNzGst

// This is an alias to an internal value.
const CustomerTaxIDTypeOmVat = shared.CustomerTaxIDTypeOmVat

// This is an alias to an internal value.
const CustomerTaxIDTypePeRuc = shared.CustomerTaxIDTypePeRuc

// This is an alias to an internal value.
const CustomerTaxIDTypePhTin = shared.CustomerTaxIDTypePhTin

// This is an alias to an internal value.
const CustomerTaxIDTypePlNip = shared.CustomerTaxIDTypePlNip

// This is an alias to an internal value.
const CustomerTaxIDTypeRoTin = shared.CustomerTaxIDTypeRoTin

// This is an alias to an internal value.
const CustomerTaxIDTypeRsPib = shared.CustomerTaxIDTypeRsPib

// This is an alias to an internal value.
const CustomerTaxIDTypeRuInn = shared.CustomerTaxIDTypeRuInn

// This is an alias to an internal value.
const CustomerTaxIDTypeRuKpp = shared.CustomerTaxIDTypeRuKpp

// This is an alias to an internal value.
const CustomerTaxIDTypeSaVat = shared.CustomerTaxIDTypeSaVat

// This is an alias to an internal value.
const CustomerTaxIDTypeSgGst = shared.CustomerTaxIDTypeSgGst

// This is an alias to an internal value.
const CustomerTaxIDTypeSgUen = shared.CustomerTaxIDTypeSgUen

// This is an alias to an internal value.
const CustomerTaxIDTypeSiTin = shared.CustomerTaxIDTypeSiTin

// This is an alias to an internal value.
const CustomerTaxIDTypeSnNinea = shared.CustomerTaxIDTypeSnNinea

// This is an alias to an internal value.
const CustomerTaxIDTypeSrFin = shared.CustomerTaxIDTypeSrFin

// This is an alias to an internal value.
const CustomerTaxIDTypeSvNit = shared.CustomerTaxIDTypeSvNit

// This is an alias to an internal value.
const CustomerTaxIDTypeThVat = shared.CustomerTaxIDTypeThVat

// This is an alias to an internal value.
const CustomerTaxIDTypeTjTin = shared.CustomerTaxIDTypeTjTin

// This is an alias to an internal value.
const CustomerTaxIDTypeTrTin = shared.CustomerTaxIDTypeTrTin

// This is an alias to an internal value.
const CustomerTaxIDTypeTwVat = shared.CustomerTaxIDTypeTwVat

// This is an alias to an internal value.
const CustomerTaxIDTypeTzVat = shared.CustomerTaxIDTypeTzVat

// This is an alias to an internal value.
const CustomerTaxIDTypeUaVat = shared.CustomerTaxIDTypeUaVat

// This is an alias to an internal value.
const CustomerTaxIDTypeUgTin = shared.CustomerTaxIDTypeUgTin

// This is an alias to an internal value.
const CustomerTaxIDTypeUsEin = shared.CustomerTaxIDTypeUsEin

// This is an alias to an internal value.
const CustomerTaxIDTypeUyRuc = shared.CustomerTaxIDTypeUyRuc

// This is an alias to an internal value.
const CustomerTaxIDTypeUzTin = shared.CustomerTaxIDTypeUzTin

// This is an alias to an internal value.
const CustomerTaxIDTypeUzVat = shared.CustomerTaxIDTypeUzVat

// This is an alias to an internal value.
const CustomerTaxIDTypeVeRif = shared.CustomerTaxIDTypeVeRif

// This is an alias to an internal value.
const CustomerTaxIDTypeVnTin = shared.CustomerTaxIDTypeVnTin

// This is an alias to an internal value.
const CustomerTaxIDTypeZaVat = shared.CustomerTaxIDTypeZaVat

// This is an alias to an internal value.
const CustomerTaxIDTypeZmTin = shared.CustomerTaxIDTypeZmTin

// This is an alias to an internal value.
const CustomerTaxIDTypeZwTin = shared.CustomerTaxIDTypeZwTin

// Tax IDs are commonly required to be displayed on customer invoices, which are
// added to the headers of invoices.
//
// ### Supported Tax ID Countries and Types
//
// | Country                | Type         | Description                                                                                             |
// | ---------------------- | ------------ | ------------------------------------------------------------------------------------------------------- |
// | Albania                | `al_tin`     | Albania Tax Identification Number                                                                       |
// | Andorra                | `ad_nrt`     | Andorran NRT Number                                                                                     |
// | Angola                 | `ao_tin`     | Angola Tax Identification Number                                                                        |
// | Argentina              | `ar_cuit`    | Argentinian Tax ID Number                                                                               |
// | Armenia                | `am_tin`     | Armenia Tax Identification Number                                                                       |
// | Aruba                  | `aw_tin`     | Aruba Tax Identification Number                                                                         |
// | Australia              | `au_abn`     | Australian Business Number (AU ABN)                                                                     |
// | Australia              | `au_arn`     | Australian Taxation Office Reference Number                                                             |
// | Austria                | `eu_vat`     | European VAT Number                                                                                     |
// | Azerbaijan             | `az_tin`     | Azerbaijan Tax Identification Number                                                                    |
// | Bahamas                | `bs_tin`     | Bahamas Tax Identification Number                                                                       |
// | Bahrain                | `bh_vat`     | Bahraini VAT Number                                                                                     |
// | Bangladesh             | `bd_bin`     | Bangladesh Business Identification Number                                                               |
// | Barbados               | `bb_tin`     | Barbados Tax Identification Number                                                                      |
// | Belarus                | `by_tin`     | Belarus TIN Number                                                                                      |
// | Belgium                | `eu_vat`     | European VAT Number                                                                                     |
// | Benin                  | `bj_ifu`     | Benin Tax Identification Number (Identifiant Fiscal Unique)                                             |
// | Bolivia                | `bo_tin`     | Bolivian Tax ID                                                                                         |
// | Bosnia and Herzegovina | `ba_tin`     | Bosnia and Herzegovina Tax Identification Number                                                        |
// | Brazil                 | `br_cnpj`    | Brazilian CNPJ Number                                                                                   |
// | Brazil                 | `br_cpf`     | Brazilian CPF Number                                                                                    |
// | Bulgaria               | `bg_uic`     | Bulgaria Unified Identification Code                                                                    |
// | Bulgaria               | `eu_vat`     | European VAT Number                                                                                     |
// | Burkina Faso           | `bf_ifu`     | Burkina Faso Tax Identification Number (Numéro d'Identifiant Fiscal Unique)                             |
// | Cambodia               | `kh_tin`     | Cambodia Tax Identification Number                                                                      |
// | Cameroon               | `cm_niu`     | Cameroon Tax Identification Number (Numéro d'Identifiant fiscal Unique)                                 |
// | Canada                 | `ca_bn`      | Canadian BN                                                                                             |
// | Canada                 | `ca_gst_hst` | Canadian GST/HST Number                                                                                 |
// | Canada                 | `ca_pst_bc`  | Canadian PST Number (British Columbia)                                                                  |
// | Canada                 | `ca_pst_mb`  | Canadian PST Number (Manitoba)                                                                          |
// | Canada                 | `ca_pst_sk`  | Canadian PST Number (Saskatchewan)                                                                      |
// | Canada                 | `ca_qst`     | Canadian QST Number (Québec)                                                                            |
// | Cape Verde             | `cv_nif`     | Cape Verde Tax Identification Number (Número de Identificação Fiscal)                                   |
// | Chile                  | `cl_tin`     | Chilean TIN                                                                                             |
// | China                  | `cn_tin`     | Chinese Tax ID                                                                                          |
// | Colombia               | `co_nit`     | Colombian NIT Number                                                                                    |
// | Congo-Kinshasa         | `cd_nif`     | Congo (DR) Tax Identification Number (Número de Identificação Fiscal)                                   |
// | Costa Rica             | `cr_tin`     | Costa Rican Tax ID                                                                                      |
// | Croatia                | `eu_vat`     | European VAT Number                                                                                     |
// | Croatia                | `hr_oib`     | Croatian Personal Identification Number (OIB)                                                           |
// | Cyprus                 | `eu_vat`     | European VAT Number                                                                                     |
// | Czech Republic         | `eu_vat`     | European VAT Number                                                                                     |
// | Denmark                | `eu_vat`     | European VAT Number                                                                                     |
// | Dominican Republic     | `do_rcn`     | Dominican RCN Number                                                                                    |
// | Ecuador                | `ec_ruc`     | Ecuadorian RUC Number                                                                                   |
// | Egypt                  | `eg_tin`     | Egyptian Tax Identification Number                                                                      |
// | El Salvador            | `sv_nit`     | El Salvadorian NIT Number                                                                               |
// | Estonia                | `eu_vat`     | European VAT Number                                                                                     |
// | Ethiopia               | `et_tin`     | Ethiopia Tax Identification Number                                                                      |
// | European Union         | `eu_oss_vat` | European One Stop Shop VAT Number for non-Union scheme                                                  |
// | Finland                | `eu_vat`     | European VAT Number                                                                                     |
// | France                 | `eu_vat`     | European VAT Number                                                                                     |
// | Georgia                | `ge_vat`     | Georgian VAT                                                                                            |
// | Germany                | `de_stn`     | German Tax Number (Steuernummer)                                                                        |
// | Germany                | `eu_vat`     | European VAT Number                                                                                     |
// | Greece                 | `eu_vat`     | European VAT Number                                                                                     |
// | Guinea                 | `gn_nif`     | Guinea Tax Identification Number (Número de Identificação Fiscal)                                       |
// | Hong Kong              | `hk_br`      | Hong Kong BR Number                                                                                     |
// | Hungary                | `eu_vat`     | European VAT Number                                                                                     |
// | Hungary                | `hu_tin`     | Hungary Tax Number (adószám)                                                                            |
// | Iceland                | `is_vat`     | Icelandic VAT                                                                                           |
// | India                  | `in_gst`     | Indian GST Number                                                                                       |
// | Indonesia              | `id_npwp`    | Indonesian NPWP Number                                                                                  |
// | Ireland                | `eu_vat`     | European VAT Number                                                                                     |
// | Israel                 | `il_vat`     | Israel VAT                                                                                              |
// | Italy                  | `eu_vat`     | European VAT Number                                                                                     |
// | Japan                  | `jp_cn`      | Japanese Corporate Number (_Hōjin Bangō_)                                                               |
// | Japan                  | `jp_rn`      | Japanese Registered Foreign Businesses' Registration Number (_Tōroku Kokugai Jigyōsha no Tōroku Bangō_) |
// | Japan                  | `jp_trn`     | Japanese Tax Registration Number (_Tōroku Bangō_)                                                       |
// | Kazakhstan             | `kz_bin`     | Kazakhstani Business Identification Number                                                              |
// | Kenya                  | `ke_pin`     | Kenya Revenue Authority Personal Identification Number                                                  |
// | Kyrgyzstan             | `kg_tin`     | Kyrgyzstan Tax Identification Number                                                                    |
// | Laos                   | `la_tin`     | Laos Tax Identification Number                                                                          |
// | Latvia                 | `eu_vat`     | European VAT Number                                                                                     |
// | Liechtenstein          | `li_uid`     | Liechtensteinian UID Number                                                                             |
// | Liechtenstein          | `li_vat`     | Liechtenstein VAT Number                                                                                |
// | Lithuania              | `eu_vat`     | European VAT Number                                                                                     |
// | Luxembourg             | `eu_vat`     | European VAT Number                                                                                     |
// | Malaysia               | `my_frp`     | Malaysian FRP Number                                                                                    |
// | Malaysia               | `my_itn`     | Malaysian ITN                                                                                           |
// | Malaysia               | `my_sst`     | Malaysian SST Number                                                                                    |
// | Malta                  | `eu_vat`     | European VAT Number                                                                                     |
// | Mauritania             | `mr_nif`     | Mauritania Tax Identification Number (Número de Identificação Fiscal)                                   |
// | Mexico                 | `mx_rfc`     | Mexican RFC Number                                                                                      |
// | Moldova                | `md_vat`     | Moldova VAT Number                                                                                      |
// | Montenegro             | `me_pib`     | Montenegro PIB Number                                                                                   |
// | Morocco                | `ma_vat`     | Morocco VAT Number                                                                                      |
// | Nepal                  | `np_pan`     | Nepal PAN Number                                                                                        |
// | Netherlands            | `eu_vat`     | European VAT Number                                                                                     |
// | New Zealand            | `nz_gst`     | New Zealand GST Number                                                                                  |
// | Nigeria                | `ng_tin`     | Nigerian Tax Identification Number                                                                      |
// | North Macedonia        | `mk_vat`     | North Macedonia VAT Number                                                                              |
// | Northern Ireland       | `eu_vat`     | Northern Ireland VAT Number                                                                             |
// | Norway                 | `no_vat`     | Norwegian VAT Number                                                                                    |
// | Norway                 | `no_voec`    | Norwegian VAT on e-commerce Number                                                                      |
// | Oman                   | `om_vat`     | Omani VAT Number                                                                                        |
// | Peru                   | `pe_ruc`     | Peruvian RUC Number                                                                                     |
// | Philippines            | `ph_tin`     | Philippines Tax Identification Number                                                                   |
// | Poland                 | `eu_vat`     | European VAT Number                                                                                     |
// | Poland                 | `pl_nip`     | Polish Tax ID Number                                                                                    |
// | Portugal               | `eu_vat`     | European VAT Number                                                                                     |
// | Romania                | `eu_vat`     | European VAT Number                                                                                     |
// | Romania                | `ro_tin`     | Romanian Tax ID Number                                                                                  |
// | Russia                 | `ru_inn`     | Russian INN                                                                                             |
// | Russia                 | `ru_kpp`     | Russian KPP                                                                                             |
// | Saudi Arabia           | `sa_vat`     | Saudi Arabia VAT                                                                                        |
// | Senegal                | `sn_ninea`   | Senegal NINEA Number                                                                                    |
// | Serbia                 | `rs_pib`     | Serbian PIB Number                                                                                      |
// | Singapore              | `sg_gst`     | Singaporean GST                                                                                         |
// | Singapore              | `sg_uen`     | Singaporean UEN                                                                                         |
// | Slovakia               | `eu_vat`     | European VAT Number                                                                                     |
// | Slovenia               | `eu_vat`     | European VAT Number                                                                                     |
// | Slovenia               | `si_tin`     | Slovenia Tax Number (davčna številka)                                                                   |
// | South Africa           | `za_vat`     | South African VAT Number                                                                                |
// | South Korea            | `kr_brn`     | Korean BRN                                                                                              |
// | Spain                  | `es_cif`     | Spanish NIF Number (previously Spanish CIF Number)                                                      |
// | Spain                  | `eu_vat`     | European VAT Number                                                                                     |
// | Suriname               | `sr_fin`     | Suriname FIN Number                                                                                     |
// | Sweden                 | `eu_vat`     | European VAT Number                                                                                     |
// | Switzerland            | `ch_uid`     | Switzerland UID Number                                                                                  |
// | Switzerland            | `ch_vat`     | Switzerland VAT Number                                                                                  |
// | Taiwan                 | `tw_vat`     | Taiwanese VAT                                                                                           |
// | Tajikistan             | `tj_tin`     | Tajikistan Tax Identification Number                                                                    |
// | Tanzania               | `tz_vat`     | Tanzania VAT Number                                                                                     |
// | Thailand               | `th_vat`     | Thai VAT                                                                                                |
// | Turkey                 | `tr_tin`     | Turkish Tax Identification Number                                                                       |
// | Uganda                 | `ug_tin`     | Uganda Tax Identification Number                                                                        |
// | Ukraine                | `ua_vat`     | Ukrainian VAT                                                                                           |
// | United Arab Emirates   | `ae_trn`     | United Arab Emirates TRN                                                                                |
// | United Kingdom         | `gb_vat`     | United Kingdom VAT Number                                                                               |
// | United States          | `us_ein`     | United States EIN                                                                                       |
// | Uruguay                | `uy_ruc`     | Uruguayan RUC Number                                                                                    |
// | Uzbekistan             | `uz_tin`     | Uzbekistan TIN Number                                                                                   |
// | Uzbekistan             | `uz_vat`     | Uzbekistan VAT Number                                                                                   |
// | Venezuela              | `ve_rif`     | Venezuelan RIF Number                                                                                   |
// | Vietnam                | `vn_tin`     | Vietnamese Tax ID Number                                                                                |
// | Zambia                 | `zm_tin`     | Zambia Tax Identification Number                                                                        |
// | Zimbabwe               | `zw_tin`     | Zimbabwe Tax Identification Number                                                                      |
//
// This is an alias to an internal type.
type CustomerTaxIDParam = shared.CustomerTaxIDParam

// This is an alias to an internal type.
type DimensionalPriceConfiguration = shared.DimensionalPriceConfiguration

// This is an alias to an internal type.
type Discount = shared.Discount

// This is an alias to an internal type.
type DiscountDiscountType = shared.DiscountDiscountType

// This is an alias to an internal value.
const DiscountDiscountTypePercentage = shared.DiscountDiscountTypePercentage

// This is an alias to an internal value.
const DiscountDiscountTypeTrial = shared.DiscountDiscountTypeTrial

// This is an alias to an internal value.
const DiscountDiscountTypeUsage = shared.DiscountDiscountTypeUsage

// This is an alias to an internal value.
const DiscountDiscountTypeAmount = shared.DiscountDiscountTypeAmount

// This is an alias to an internal type.
type DiscountUnionParam = shared.DiscountUnionParam

// This is an alias to an internal type.
type FixedFeeQuantityScheduleEntry = shared.FixedFeeQuantityScheduleEntry

// This is an alias to an internal type.
type FixedFeeQuantityTransition = shared.FixedFeeQuantityTransition

// An [`Invoice`](/core-concepts#invoice) is a fundamental billing entity,
// representing the request for payment for a single subscription. This includes a
// set of line items, which correspond to prices in the subscription's plan and can
// represent fixed recurring fees or usage-based fees. They are generated at the
// end of a billing period, or as the result of an action, such as a cancellation.
//
// This is an alias to an internal type.
type Invoice = shared.Invoice

// This is an alias to an internal type.
type InvoiceAutoCollection = shared.InvoiceAutoCollection

// This is an alias to an internal type.
type InvoiceCreditNote = shared.InvoiceCreditNote

// This is an alias to an internal type.
type InvoiceCustomerBalanceTransaction = shared.InvoiceCustomerBalanceTransaction

// This is an alias to an internal type.
type InvoiceCustomerBalanceTransactionsAction = shared.InvoiceCustomerBalanceTransactionsAction

// This is an alias to an internal value.
const InvoiceCustomerBalanceTransactionsActionAppliedToInvoice = shared.InvoiceCustomerBalanceTransactionsActionAppliedToInvoice

// This is an alias to an internal value.
const InvoiceCustomerBalanceTransactionsActionManualAdjustment = shared.InvoiceCustomerBalanceTransactionsActionManualAdjustment

// This is an alias to an internal value.
const InvoiceCustomerBalanceTransactionsActionProratedRefund = shared.InvoiceCustomerBalanceTransactionsActionProratedRefund

// This is an alias to an internal value.
const InvoiceCustomerBalanceTransactionsActionRevertProratedRefund = shared.InvoiceCustomerBalanceTransactionsActionRevertProratedRefund

// This is an alias to an internal value.
const InvoiceCustomerBalanceTransactionsActionReturnFromVoiding = shared.InvoiceCustomerBalanceTransactionsActionReturnFromVoiding

// This is an alias to an internal value.
const InvoiceCustomerBalanceTransactionsActionCreditNoteApplied = shared.InvoiceCustomerBalanceTransactionsActionCreditNoteApplied

// This is an alias to an internal value.
const InvoiceCustomerBalanceTransactionsActionCreditNoteVoided = shared.InvoiceCustomerBalanceTransactionsActionCreditNoteVoided

// This is an alias to an internal value.
const InvoiceCustomerBalanceTransactionsActionOverpaymentRefund = shared.InvoiceCustomerBalanceTransactionsActionOverpaymentRefund

// This is an alias to an internal value.
const InvoiceCustomerBalanceTransactionsActionExternalPayment = shared.InvoiceCustomerBalanceTransactionsActionExternalPayment

// This is an alias to an internal value.
const InvoiceCustomerBalanceTransactionsActionSmallInvoiceCarryover = shared.InvoiceCustomerBalanceTransactionsActionSmallInvoiceCarryover

// This is an alias to an internal type.
type InvoiceCustomerBalanceTransactionsType = shared.InvoiceCustomerBalanceTransactionsType

// This is an alias to an internal value.
const InvoiceCustomerBalanceTransactionsTypeIncrement = shared.InvoiceCustomerBalanceTransactionsTypeIncrement

// This is an alias to an internal value.
const InvoiceCustomerBalanceTransactionsTypeDecrement = shared.InvoiceCustomerBalanceTransactionsTypeDecrement

// This is an alias to an internal type.
type InvoiceInvoiceSource = shared.InvoiceInvoiceSource

// This is an alias to an internal value.
const InvoiceInvoiceSourceSubscription = shared.InvoiceInvoiceSourceSubscription

// This is an alias to an internal value.
const InvoiceInvoiceSourcePartial = shared.InvoiceInvoiceSourcePartial

// This is an alias to an internal value.
const InvoiceInvoiceSourceOneOff = shared.InvoiceInvoiceSourceOneOff

// This is an alias to an internal type.
type InvoiceLineItem = shared.InvoiceLineItem

// This is an alias to an internal type.
type InvoiceLineItemsAdjustment = shared.InvoiceLineItemsAdjustment

// This is an alias to an internal type.
type InvoiceLineItemsAdjustmentsAdjustmentType = shared.InvoiceLineItemsAdjustmentsAdjustmentType

// This is an alias to an internal value.
const InvoiceLineItemsAdjustmentsAdjustmentTypeUsageDiscount = shared.InvoiceLineItemsAdjustmentsAdjustmentTypeUsageDiscount

// This is an alias to an internal value.
const InvoiceLineItemsAdjustmentsAdjustmentTypeAmountDiscount = shared.InvoiceLineItemsAdjustmentsAdjustmentTypeAmountDiscount

// This is an alias to an internal value.
const InvoiceLineItemsAdjustmentsAdjustmentTypePercentageDiscount = shared.InvoiceLineItemsAdjustmentsAdjustmentTypePercentageDiscount

// This is an alias to an internal value.
const InvoiceLineItemsAdjustmentsAdjustmentTypeMinimum = shared.InvoiceLineItemsAdjustmentsAdjustmentTypeMinimum

// This is an alias to an internal value.
const InvoiceLineItemsAdjustmentsAdjustmentTypeMaximum = shared.InvoiceLineItemsAdjustmentsAdjustmentTypeMaximum

// This is an alias to an internal type.
type InvoiceLineItemsSubLineItem = shared.InvoiceLineItemsSubLineItem

// This is an alias to an internal type.
type InvoiceLineItemsSubLineItemsType = shared.InvoiceLineItemsSubLineItemsType

// This is an alias to an internal value.
const InvoiceLineItemsSubLineItemsTypeMatrix = shared.InvoiceLineItemsSubLineItemsTypeMatrix

// This is an alias to an internal value.
const InvoiceLineItemsSubLineItemsTypeTier = shared.InvoiceLineItemsSubLineItemsTypeTier

// This is an alias to an internal value.
const InvoiceLineItemsSubLineItemsTypeNull = shared.InvoiceLineItemsSubLineItemsTypeNull

// This is an alias to an internal type.
type InvoicePaymentAttempt = shared.InvoicePaymentAttempt

// The payment provider that attempted to collect the payment.
//
// This is an alias to an internal type.
type InvoicePaymentAttemptsPaymentProvider = shared.InvoicePaymentAttemptsPaymentProvider

// This is an alias to an internal value.
const InvoicePaymentAttemptsPaymentProviderStripe = shared.InvoicePaymentAttemptsPaymentProviderStripe

// This is an alias to an internal type.
type InvoiceStatus = shared.InvoiceStatus

// This is an alias to an internal value.
const InvoiceStatusIssued = shared.InvoiceStatusIssued

// This is an alias to an internal value.
const InvoiceStatusPaid = shared.InvoiceStatusPaid

// This is an alias to an internal value.
const InvoiceStatusSynced = shared.InvoiceStatusSynced

// This is an alias to an internal value.
const InvoiceStatusVoid = shared.InvoiceStatusVoid

// This is an alias to an internal value.
const InvoiceStatusDraft = shared.InvoiceStatusDraft

// This is an alias to an internal type.
type InvoiceLevelDiscount = shared.InvoiceLevelDiscount

// This is an alias to an internal type.
type InvoiceLevelDiscountDiscountType = shared.InvoiceLevelDiscountDiscountType

// This is an alias to an internal value.
const InvoiceLevelDiscountDiscountTypePercentage = shared.InvoiceLevelDiscountDiscountTypePercentage

// This is an alias to an internal value.
const InvoiceLevelDiscountDiscountTypeAmount = shared.InvoiceLevelDiscountDiscountTypeAmount

// This is an alias to an internal value.
const InvoiceLevelDiscountDiscountTypeTrial = shared.InvoiceLevelDiscountDiscountTypeTrial

// This is an alias to an internal type.
type InvoiceTiny = shared.InvoiceTiny

// A minimal representation of an Item containing only the essential identifying
// information.
//
// This is an alias to an internal type.
type ItemSlim = shared.ItemSlim

// Configuration for matrix pricing
//
// This is an alias to an internal type.
type MatrixConfig = shared.MatrixConfig

// Configuration for matrix pricing
//
// This is an alias to an internal type.
type MatrixConfigParam = shared.MatrixConfigParam

// This is an alias to an internal type.
type MatrixSubLineItem = shared.MatrixSubLineItem

// This is an alias to an internal type.
type MatrixSubLineItemType = shared.MatrixSubLineItemType

// This is an alias to an internal value.
const MatrixSubLineItemTypeMatrix = shared.MatrixSubLineItemTypeMatrix

// Configuration for a single matrix value
//
// This is an alias to an internal type.
type MatrixValue = shared.MatrixValue

// Configuration for a single matrix value
//
// This is an alias to an internal type.
type MatrixValueParam = shared.MatrixValueParam

// Configuration for matrix pricing with usage allocation
//
// This is an alias to an internal type.
type MatrixWithAllocationConfig = shared.MatrixWithAllocationConfig

// Configuration for a single matrix value
//
// This is an alias to an internal type.
type MatrixWithAllocationConfigMatrixValue = shared.MatrixWithAllocationConfigMatrixValue

// Configuration for matrix pricing with usage allocation
//
// This is an alias to an internal type.
type MatrixWithAllocationConfigParam = shared.MatrixWithAllocationConfigParam

// Configuration for a single matrix value
//
// This is an alias to an internal type.
type MatrixWithAllocationConfigMatrixValueParam = shared.MatrixWithAllocationConfigMatrixValueParam

// This is an alias to an internal type.
type Maximum = shared.Maximum

// This is an alias to an internal type.
type MaximumFilter = shared.MaximumFilter

// The property of the price to filter on.
//
// This is an alias to an internal type.
type MaximumFiltersField = shared.MaximumFiltersField

// This is an alias to an internal value.
const MaximumFiltersFieldPriceID = shared.MaximumFiltersFieldPriceID

// This is an alias to an internal value.
const MaximumFiltersFieldItemID = shared.MaximumFiltersFieldItemID

// This is an alias to an internal value.
const MaximumFiltersFieldPriceType = shared.MaximumFiltersFieldPriceType

// This is an alias to an internal value.
const MaximumFiltersFieldCurrency = shared.MaximumFiltersFieldCurrency

// This is an alias to an internal value.
const MaximumFiltersFieldPricingUnitID = shared.MaximumFiltersFieldPricingUnitID

// Should prices that match the filter be included or excluded.
//
// This is an alias to an internal type.
type MaximumFiltersOperator = shared.MaximumFiltersOperator

// This is an alias to an internal value.
const MaximumFiltersOperatorIncludes = shared.MaximumFiltersOperatorIncludes

// This is an alias to an internal value.
const MaximumFiltersOperatorExcludes = shared.MaximumFiltersOperatorExcludes

// This is an alias to an internal type.
type MaximumInterval = shared.MaximumInterval

// This is an alias to an internal type.
type MaximumIntervalFilter = shared.MaximumIntervalFilter

// The property of the price to filter on.
//
// This is an alias to an internal type.
type MaximumIntervalFiltersField = shared.MaximumIntervalFiltersField

// This is an alias to an internal value.
const MaximumIntervalFiltersFieldPriceID = shared.MaximumIntervalFiltersFieldPriceID

// This is an alias to an internal value.
const MaximumIntervalFiltersFieldItemID = shared.MaximumIntervalFiltersFieldItemID

// This is an alias to an internal value.
const MaximumIntervalFiltersFieldPriceType = shared.MaximumIntervalFiltersFieldPriceType

// This is an alias to an internal value.
const MaximumIntervalFiltersFieldCurrency = shared.MaximumIntervalFiltersFieldCurrency

// This is an alias to an internal value.
const MaximumIntervalFiltersFieldPricingUnitID = shared.MaximumIntervalFiltersFieldPricingUnitID

// Should prices that match the filter be included or excluded.
//
// This is an alias to an internal type.
type MaximumIntervalFiltersOperator = shared.MaximumIntervalFiltersOperator

// This is an alias to an internal value.
const MaximumIntervalFiltersOperatorIncludes = shared.MaximumIntervalFiltersOperatorIncludes

// This is an alias to an internal value.
const MaximumIntervalFiltersOperatorExcludes = shared.MaximumIntervalFiltersOperatorExcludes

// This is an alias to an internal type.
type Minimum = shared.Minimum

// This is an alias to an internal type.
type MinimumFilter = shared.MinimumFilter

// The property of the price to filter on.
//
// This is an alias to an internal type.
type MinimumFiltersField = shared.MinimumFiltersField

// This is an alias to an internal value.
const MinimumFiltersFieldPriceID = shared.MinimumFiltersFieldPriceID

// This is an alias to an internal value.
const MinimumFiltersFieldItemID = shared.MinimumFiltersFieldItemID

// This is an alias to an internal value.
const MinimumFiltersFieldPriceType = shared.MinimumFiltersFieldPriceType

// This is an alias to an internal value.
const MinimumFiltersFieldCurrency = shared.MinimumFiltersFieldCurrency

// This is an alias to an internal value.
const MinimumFiltersFieldPricingUnitID = shared.MinimumFiltersFieldPricingUnitID

// Should prices that match the filter be included or excluded.
//
// This is an alias to an internal type.
type MinimumFiltersOperator = shared.MinimumFiltersOperator

// This is an alias to an internal value.
const MinimumFiltersOperatorIncludes = shared.MinimumFiltersOperatorIncludes

// This is an alias to an internal value.
const MinimumFiltersOperatorExcludes = shared.MinimumFiltersOperatorExcludes

// This is an alias to an internal type.
type MinimumInterval = shared.MinimumInterval

// This is an alias to an internal type.
type MinimumIntervalFilter = shared.MinimumIntervalFilter

// The property of the price to filter on.
//
// This is an alias to an internal type.
type MinimumIntervalFiltersField = shared.MinimumIntervalFiltersField

// This is an alias to an internal value.
const MinimumIntervalFiltersFieldPriceID = shared.MinimumIntervalFiltersFieldPriceID

// This is an alias to an internal value.
const MinimumIntervalFiltersFieldItemID = shared.MinimumIntervalFiltersFieldItemID

// This is an alias to an internal value.
const MinimumIntervalFiltersFieldPriceType = shared.MinimumIntervalFiltersFieldPriceType

// This is an alias to an internal value.
const MinimumIntervalFiltersFieldCurrency = shared.MinimumIntervalFiltersFieldCurrency

// This is an alias to an internal value.
const MinimumIntervalFiltersFieldPricingUnitID = shared.MinimumIntervalFiltersFieldPricingUnitID

// Should prices that match the filter be included or excluded.
//
// This is an alias to an internal type.
type MinimumIntervalFiltersOperator = shared.MinimumIntervalFiltersOperator

// This is an alias to an internal value.
const MinimumIntervalFiltersOperatorIncludes = shared.MinimumIntervalFiltersOperatorIncludes

// This is an alias to an internal value.
const MinimumIntervalFiltersOperatorExcludes = shared.MinimumIntervalFiltersOperatorExcludes

// This is an alias to an internal type.
type MonetaryAmountDiscountAdjustment = shared.MonetaryAmountDiscountAdjustment

// This is an alias to an internal type.
type MonetaryAmountDiscountAdjustmentAdjustmentType = shared.MonetaryAmountDiscountAdjustmentAdjustmentType

// This is an alias to an internal value.
const MonetaryAmountDiscountAdjustmentAdjustmentTypeAmountDiscount = shared.MonetaryAmountDiscountAdjustmentAdjustmentTypeAmountDiscount

// This is an alias to an internal type.
type MonetaryAmountDiscountAdjustmentFilter = shared.MonetaryAmountDiscountAdjustmentFilter

// The property of the price to filter on.
//
// This is an alias to an internal type.
type MonetaryAmountDiscountAdjustmentFiltersField = shared.MonetaryAmountDiscountAdjustmentFiltersField

// This is an alias to an internal value.
const MonetaryAmountDiscountAdjustmentFiltersFieldPriceID = shared.MonetaryAmountDiscountAdjustmentFiltersFieldPriceID

// This is an alias to an internal value.
const MonetaryAmountDiscountAdjustmentFiltersFieldItemID = shared.MonetaryAmountDiscountAdjustmentFiltersFieldItemID

// This is an alias to an internal value.
const MonetaryAmountDiscountAdjustmentFiltersFieldPriceType = shared.MonetaryAmountDiscountAdjustmentFiltersFieldPriceType

// This is an alias to an internal value.
const MonetaryAmountDiscountAdjustmentFiltersFieldCurrency = shared.MonetaryAmountDiscountAdjustmentFiltersFieldCurrency

// This is an alias to an internal value.
const MonetaryAmountDiscountAdjustmentFiltersFieldPricingUnitID = shared.MonetaryAmountDiscountAdjustmentFiltersFieldPricingUnitID

// Should prices that match the filter be included or excluded.
//
// This is an alias to an internal type.
type MonetaryAmountDiscountAdjustmentFiltersOperator = shared.MonetaryAmountDiscountAdjustmentFiltersOperator

// This is an alias to an internal value.
const MonetaryAmountDiscountAdjustmentFiltersOperatorIncludes = shared.MonetaryAmountDiscountAdjustmentFiltersOperatorIncludes

// This is an alias to an internal value.
const MonetaryAmountDiscountAdjustmentFiltersOperatorExcludes = shared.MonetaryAmountDiscountAdjustmentFiltersOperatorExcludes

// This is an alias to an internal type.
type MonetaryMaximumAdjustment = shared.MonetaryMaximumAdjustment

// This is an alias to an internal type.
type MonetaryMaximumAdjustmentAdjustmentType = shared.MonetaryMaximumAdjustmentAdjustmentType

// This is an alias to an internal value.
const MonetaryMaximumAdjustmentAdjustmentTypeMaximum = shared.MonetaryMaximumAdjustmentAdjustmentTypeMaximum

// This is an alias to an internal type.
type MonetaryMaximumAdjustmentFilter = shared.MonetaryMaximumAdjustmentFilter

// The property of the price to filter on.
//
// This is an alias to an internal type.
type MonetaryMaximumAdjustmentFiltersField = shared.MonetaryMaximumAdjustmentFiltersField

// This is an alias to an internal value.
const MonetaryMaximumAdjustmentFiltersFieldPriceID = shared.MonetaryMaximumAdjustmentFiltersFieldPriceID

// This is an alias to an internal value.
const MonetaryMaximumAdjustmentFiltersFieldItemID = shared.MonetaryMaximumAdjustmentFiltersFieldItemID

// This is an alias to an internal value.
const MonetaryMaximumAdjustmentFiltersFieldPriceType = shared.MonetaryMaximumAdjustmentFiltersFieldPriceType

// This is an alias to an internal value.
const MonetaryMaximumAdjustmentFiltersFieldCurrency = shared.MonetaryMaximumAdjustmentFiltersFieldCurrency

// This is an alias to an internal value.
const MonetaryMaximumAdjustmentFiltersFieldPricingUnitID = shared.MonetaryMaximumAdjustmentFiltersFieldPricingUnitID

// Should prices that match the filter be included or excluded.
//
// This is an alias to an internal type.
type MonetaryMaximumAdjustmentFiltersOperator = shared.MonetaryMaximumAdjustmentFiltersOperator

// This is an alias to an internal value.
const MonetaryMaximumAdjustmentFiltersOperatorIncludes = shared.MonetaryMaximumAdjustmentFiltersOperatorIncludes

// This is an alias to an internal value.
const MonetaryMaximumAdjustmentFiltersOperatorExcludes = shared.MonetaryMaximumAdjustmentFiltersOperatorExcludes

// This is an alias to an internal type.
type MonetaryMinimumAdjustment = shared.MonetaryMinimumAdjustment

// This is an alias to an internal type.
type MonetaryMinimumAdjustmentAdjustmentType = shared.MonetaryMinimumAdjustmentAdjustmentType

// This is an alias to an internal value.
const MonetaryMinimumAdjustmentAdjustmentTypeMinimum = shared.MonetaryMinimumAdjustmentAdjustmentTypeMinimum

// This is an alias to an internal type.
type MonetaryMinimumAdjustmentFilter = shared.MonetaryMinimumAdjustmentFilter

// The property of the price to filter on.
//
// This is an alias to an internal type.
type MonetaryMinimumAdjustmentFiltersField = shared.MonetaryMinimumAdjustmentFiltersField

// This is an alias to an internal value.
const MonetaryMinimumAdjustmentFiltersFieldPriceID = shared.MonetaryMinimumAdjustmentFiltersFieldPriceID

// This is an alias to an internal value.
const MonetaryMinimumAdjustmentFiltersFieldItemID = shared.MonetaryMinimumAdjustmentFiltersFieldItemID

// This is an alias to an internal value.
const MonetaryMinimumAdjustmentFiltersFieldPriceType = shared.MonetaryMinimumAdjustmentFiltersFieldPriceType

// This is an alias to an internal value.
const MonetaryMinimumAdjustmentFiltersFieldCurrency = shared.MonetaryMinimumAdjustmentFiltersFieldCurrency

// This is an alias to an internal value.
const MonetaryMinimumAdjustmentFiltersFieldPricingUnitID = shared.MonetaryMinimumAdjustmentFiltersFieldPricingUnitID

// Should prices that match the filter be included or excluded.
//
// This is an alias to an internal type.
type MonetaryMinimumAdjustmentFiltersOperator = shared.MonetaryMinimumAdjustmentFiltersOperator

// This is an alias to an internal value.
const MonetaryMinimumAdjustmentFiltersOperatorIncludes = shared.MonetaryMinimumAdjustmentFiltersOperatorIncludes

// This is an alias to an internal value.
const MonetaryMinimumAdjustmentFiltersOperatorExcludes = shared.MonetaryMinimumAdjustmentFiltersOperatorExcludes

// This is an alias to an internal type.
type MonetaryPercentageDiscountAdjustment = shared.MonetaryPercentageDiscountAdjustment

// This is an alias to an internal type.
type MonetaryPercentageDiscountAdjustmentAdjustmentType = shared.MonetaryPercentageDiscountAdjustmentAdjustmentType

// This is an alias to an internal value.
const MonetaryPercentageDiscountAdjustmentAdjustmentTypePercentageDiscount = shared.MonetaryPercentageDiscountAdjustmentAdjustmentTypePercentageDiscount

// This is an alias to an internal type.
type MonetaryPercentageDiscountAdjustmentFilter = shared.MonetaryPercentageDiscountAdjustmentFilter

// The property of the price to filter on.
//
// This is an alias to an internal type.
type MonetaryPercentageDiscountAdjustmentFiltersField = shared.MonetaryPercentageDiscountAdjustmentFiltersField

// This is an alias to an internal value.
const MonetaryPercentageDiscountAdjustmentFiltersFieldPriceID = shared.MonetaryPercentageDiscountAdjustmentFiltersFieldPriceID

// This is an alias to an internal value.
const MonetaryPercentageDiscountAdjustmentFiltersFieldItemID = shared.MonetaryPercentageDiscountAdjustmentFiltersFieldItemID

// This is an alias to an internal value.
const MonetaryPercentageDiscountAdjustmentFiltersFieldPriceType = shared.MonetaryPercentageDiscountAdjustmentFiltersFieldPriceType

// This is an alias to an internal value.
const MonetaryPercentageDiscountAdjustmentFiltersFieldCurrency = shared.MonetaryPercentageDiscountAdjustmentFiltersFieldCurrency

// This is an alias to an internal value.
const MonetaryPercentageDiscountAdjustmentFiltersFieldPricingUnitID = shared.MonetaryPercentageDiscountAdjustmentFiltersFieldPricingUnitID

// Should prices that match the filter be included or excluded.
//
// This is an alias to an internal type.
type MonetaryPercentageDiscountAdjustmentFiltersOperator = shared.MonetaryPercentageDiscountAdjustmentFiltersOperator

// This is an alias to an internal value.
const MonetaryPercentageDiscountAdjustmentFiltersOperatorIncludes = shared.MonetaryPercentageDiscountAdjustmentFiltersOperatorIncludes

// This is an alias to an internal value.
const MonetaryPercentageDiscountAdjustmentFiltersOperatorExcludes = shared.MonetaryPercentageDiscountAdjustmentFiltersOperatorExcludes

// This is an alias to an internal type.
type MonetaryUsageDiscountAdjustment = shared.MonetaryUsageDiscountAdjustment

// This is an alias to an internal type.
type MonetaryUsageDiscountAdjustmentAdjustmentType = shared.MonetaryUsageDiscountAdjustmentAdjustmentType

// This is an alias to an internal value.
const MonetaryUsageDiscountAdjustmentAdjustmentTypeUsageDiscount = shared.MonetaryUsageDiscountAdjustmentAdjustmentTypeUsageDiscount

// This is an alias to an internal type.
type MonetaryUsageDiscountAdjustmentFilter = shared.MonetaryUsageDiscountAdjustmentFilter

// The property of the price to filter on.
//
// This is an alias to an internal type.
type MonetaryUsageDiscountAdjustmentFiltersField = shared.MonetaryUsageDiscountAdjustmentFiltersField

// This is an alias to an internal value.
const MonetaryUsageDiscountAdjustmentFiltersFieldPriceID = shared.MonetaryUsageDiscountAdjustmentFiltersFieldPriceID

// This is an alias to an internal value.
const MonetaryUsageDiscountAdjustmentFiltersFieldItemID = shared.MonetaryUsageDiscountAdjustmentFiltersFieldItemID

// This is an alias to an internal value.
const MonetaryUsageDiscountAdjustmentFiltersFieldPriceType = shared.MonetaryUsageDiscountAdjustmentFiltersFieldPriceType

// This is an alias to an internal value.
const MonetaryUsageDiscountAdjustmentFiltersFieldCurrency = shared.MonetaryUsageDiscountAdjustmentFiltersFieldCurrency

// This is an alias to an internal value.
const MonetaryUsageDiscountAdjustmentFiltersFieldPricingUnitID = shared.MonetaryUsageDiscountAdjustmentFiltersFieldPricingUnitID

// Should prices that match the filter be included or excluded.
//
// This is an alias to an internal type.
type MonetaryUsageDiscountAdjustmentFiltersOperator = shared.MonetaryUsageDiscountAdjustmentFiltersOperator

// This is an alias to an internal value.
const MonetaryUsageDiscountAdjustmentFiltersOperatorIncludes = shared.MonetaryUsageDiscountAdjustmentFiltersOperatorIncludes

// This is an alias to an internal value.
const MonetaryUsageDiscountAdjustmentFiltersOperatorExcludes = shared.MonetaryUsageDiscountAdjustmentFiltersOperatorExcludes

// This is an alias to an internal type.
type NewAllocationPriceParam = shared.NewAllocationPriceParam

// The cadence at which to allocate the amount to the customer.
//
// This is an alias to an internal type.
type NewAllocationPriceCadence = shared.NewAllocationPriceCadence

// This is an alias to an internal value.
const NewAllocationPriceCadenceOneTime = shared.NewAllocationPriceCadenceOneTime

// This is an alias to an internal value.
const NewAllocationPriceCadenceMonthly = shared.NewAllocationPriceCadenceMonthly

// This is an alias to an internal value.
const NewAllocationPriceCadenceQuarterly = shared.NewAllocationPriceCadenceQuarterly

// This is an alias to an internal value.
const NewAllocationPriceCadenceSemiAnnual = shared.NewAllocationPriceCadenceSemiAnnual

// This is an alias to an internal value.
const NewAllocationPriceCadenceAnnual = shared.NewAllocationPriceCadenceAnnual

// A PriceFilter that only allows item_id field for block filters.
//
// This is an alias to an internal type.
type NewAllocationPriceFilterParam = shared.NewAllocationPriceFilterParam

// The property of the price the block applies to. Only item_id is supported.
//
// This is an alias to an internal type.
type NewAllocationPriceFiltersField = shared.NewAllocationPriceFiltersField

// This is an alias to an internal value.
const NewAllocationPriceFiltersFieldItemID = shared.NewAllocationPriceFiltersFieldItemID

// Should prices that match the filter be included or excluded.
//
// This is an alias to an internal type.
type NewAllocationPriceFiltersOperator = shared.NewAllocationPriceFiltersOperator

// This is an alias to an internal value.
const NewAllocationPriceFiltersOperatorIncludes = shared.NewAllocationPriceFiltersOperatorIncludes

// This is an alias to an internal value.
const NewAllocationPriceFiltersOperatorExcludes = shared.NewAllocationPriceFiltersOperatorExcludes

// This is an alias to an internal type.
type NewAmountDiscountParam = shared.NewAmountDiscountParam

// This is an alias to an internal type.
type NewAmountDiscountAdjustmentType = shared.NewAmountDiscountAdjustmentType

// This is an alias to an internal value.
const NewAmountDiscountAdjustmentTypeAmountDiscount = shared.NewAmountDiscountAdjustmentTypeAmountDiscount

// If set, the adjustment will apply to every price on the subscription.
//
// This is an alias to an internal type.
type NewAmountDiscountAppliesToAll = shared.NewAmountDiscountAppliesToAll

// This is an alias to an internal value.
const NewAmountDiscountAppliesToAllTrue = shared.NewAmountDiscountAppliesToAllTrue

// This is an alias to an internal type.
type NewAmountDiscountFilterParam = shared.NewAmountDiscountFilterParam

// The property of the price to filter on.
//
// This is an alias to an internal type.
type NewAmountDiscountFiltersField = shared.NewAmountDiscountFiltersField

// This is an alias to an internal value.
const NewAmountDiscountFiltersFieldPriceID = shared.NewAmountDiscountFiltersFieldPriceID

// This is an alias to an internal value.
const NewAmountDiscountFiltersFieldItemID = shared.NewAmountDiscountFiltersFieldItemID

// This is an alias to an internal value.
const NewAmountDiscountFiltersFieldPriceType = shared.NewAmountDiscountFiltersFieldPriceType

// This is an alias to an internal value.
const NewAmountDiscountFiltersFieldCurrency = shared.NewAmountDiscountFiltersFieldCurrency

// This is an alias to an internal value.
const NewAmountDiscountFiltersFieldPricingUnitID = shared.NewAmountDiscountFiltersFieldPricingUnitID

// Should prices that match the filter be included or excluded.
//
// This is an alias to an internal type.
type NewAmountDiscountFiltersOperator = shared.NewAmountDiscountFiltersOperator

// This is an alias to an internal value.
const NewAmountDiscountFiltersOperatorIncludes = shared.NewAmountDiscountFiltersOperatorIncludes

// This is an alias to an internal value.
const NewAmountDiscountFiltersOperatorExcludes = shared.NewAmountDiscountFiltersOperatorExcludes

// If set, only prices of the specified type will have the adjustment applied.
//
// This is an alias to an internal type.
type NewAmountDiscountPriceType = shared.NewAmountDiscountPriceType

// This is an alias to an internal value.
const NewAmountDiscountPriceTypeUsage = shared.NewAmountDiscountPriceTypeUsage

// This is an alias to an internal value.
const NewAmountDiscountPriceTypeFixedInAdvance = shared.NewAmountDiscountPriceTypeFixedInAdvance

// This is an alias to an internal value.
const NewAmountDiscountPriceTypeFixedInArrears = shared.NewAmountDiscountPriceTypeFixedInArrears

// This is an alias to an internal value.
const NewAmountDiscountPriceTypeFixed = shared.NewAmountDiscountPriceTypeFixed

// This is an alias to an internal value.
const NewAmountDiscountPriceTypeInArrears = shared.NewAmountDiscountPriceTypeInArrears

// This is an alias to an internal type.
type NewBillingCycleConfigurationParam = shared.NewBillingCycleConfigurationParam

// The unit of billing period duration.
//
// This is an alias to an internal type.
type NewBillingCycleConfigurationDurationUnit = shared.NewBillingCycleConfigurationDurationUnit

// This is an alias to an internal value.
const NewBillingCycleConfigurationDurationUnitDay = shared.NewBillingCycleConfigurationDurationUnitDay

// This is an alias to an internal value.
const NewBillingCycleConfigurationDurationUnitMonth = shared.NewBillingCycleConfigurationDurationUnitMonth

// This is an alias to an internal type.
type NewDimensionalPriceConfigurationParam = shared.NewDimensionalPriceConfigurationParam

// This is an alias to an internal type.
type NewFloatingBulkPriceParam = shared.NewFloatingBulkPriceParam

// The cadence to bill for this price on.
//
// This is an alias to an internal type.
type NewFloatingBulkPriceCadence = shared.NewFloatingBulkPriceCadence

// This is an alias to an internal value.
const NewFloatingBulkPriceCadenceAnnual = shared.NewFloatingBulkPriceCadenceAnnual

// This is an alias to an internal value.
const NewFloatingBulkPriceCadenceSemiAnnual = shared.NewFloatingBulkPriceCadenceSemiAnnual

// This is an alias to an internal value.
const NewFloatingBulkPriceCadenceMonthly = shared.NewFloatingBulkPriceCadenceMonthly

// This is an alias to an internal value.
const NewFloatingBulkPriceCadenceQuarterly = shared.NewFloatingBulkPriceCadenceQuarterly

// This is an alias to an internal value.
const NewFloatingBulkPriceCadenceOneTime = shared.NewFloatingBulkPriceCadenceOneTime

// This is an alias to an internal value.
const NewFloatingBulkPriceCadenceCustom = shared.NewFloatingBulkPriceCadenceCustom

// The pricing model type
//
// This is an alias to an internal type.
type NewFloatingBulkPriceModelType = shared.NewFloatingBulkPriceModelType

// This is an alias to an internal value.
const NewFloatingBulkPriceModelTypeBulk = shared.NewFloatingBulkPriceModelTypeBulk

// This is an alias to an internal type.
type NewFloatingBulkPriceConversionRateConfigUnionParam = shared.NewFloatingBulkPriceConversionRateConfigUnionParam

// This is an alias to an internal type.
type NewFloatingBulkPriceConversionRateConfigConversionRateType = shared.NewFloatingBulkPriceConversionRateConfigConversionRateType

// This is an alias to an internal value.
const NewFloatingBulkPriceConversionRateConfigConversionRateTypeUnit = shared.NewFloatingBulkPriceConversionRateConfigConversionRateTypeUnit

// This is an alias to an internal value.
const NewFloatingBulkPriceConversionRateConfigConversionRateTypeTiered = shared.NewFloatingBulkPriceConversionRateConfigConversionRateTypeTiered

// This is an alias to an internal type.
type NewFloatingBulkWithProrationPriceParam = shared.NewFloatingBulkWithProrationPriceParam

// Configuration for bulk_with_proration pricing
//
// This is an alias to an internal type.
type NewFloatingBulkWithProrationPriceBulkWithProrationConfigParam = shared.NewFloatingBulkWithProrationPriceBulkWithProrationConfigParam

// Configuration for a single bulk pricing tier with proration
//
// This is an alias to an internal type.
type NewFloatingBulkWithProrationPriceBulkWithProrationConfigTierParam = shared.NewFloatingBulkWithProrationPriceBulkWithProrationConfigTierParam

// The cadence to bill for this price on.
//
// This is an alias to an internal type.
type NewFloatingBulkWithProrationPriceCadence = shared.NewFloatingBulkWithProrationPriceCadence

// This is an alias to an internal value.
const NewFloatingBulkWithProrationPriceCadenceAnnual = shared.NewFloatingBulkWithProrationPriceCadenceAnnual

// This is an alias to an internal value.
const NewFloatingBulkWithProrationPriceCadenceSemiAnnual = shared.NewFloatingBulkWithProrationPriceCadenceSemiAnnual

// This is an alias to an internal value.
const NewFloatingBulkWithProrationPriceCadenceMonthly = shared.NewFloatingBulkWithProrationPriceCadenceMonthly

// This is an alias to an internal value.
const NewFloatingBulkWithProrationPriceCadenceQuarterly = shared.NewFloatingBulkWithProrationPriceCadenceQuarterly

// This is an alias to an internal value.
const NewFloatingBulkWithProrationPriceCadenceOneTime = shared.NewFloatingBulkWithProrationPriceCadenceOneTime

// This is an alias to an internal value.
const NewFloatingBulkWithProrationPriceCadenceCustom = shared.NewFloatingBulkWithProrationPriceCadenceCustom

// The pricing model type
//
// This is an alias to an internal type.
type NewFloatingBulkWithProrationPriceModelType = shared.NewFloatingBulkWithProrationPriceModelType

// This is an alias to an internal value.
const NewFloatingBulkWithProrationPriceModelTypeBulkWithProration = shared.NewFloatingBulkWithProrationPriceModelTypeBulkWithProration

// This is an alias to an internal type.
type NewFloatingBulkWithProrationPriceConversionRateConfigUnionParam = shared.NewFloatingBulkWithProrationPriceConversionRateConfigUnionParam

// This is an alias to an internal type.
type NewFloatingBulkWithProrationPriceConversionRateConfigConversionRateType = shared.NewFloatingBulkWithProrationPriceConversionRateConfigConversionRateType

// This is an alias to an internal value.
const NewFloatingBulkWithProrationPriceConversionRateConfigConversionRateTypeUnit = shared.NewFloatingBulkWithProrationPriceConversionRateConfigConversionRateTypeUnit

// This is an alias to an internal value.
const NewFloatingBulkWithProrationPriceConversionRateConfigConversionRateTypeTiered = shared.NewFloatingBulkWithProrationPriceConversionRateConfigConversionRateTypeTiered

// This is an alias to an internal type.
type NewFloatingCumulativeGroupedBulkPriceParam = shared.NewFloatingCumulativeGroupedBulkPriceParam

// The cadence to bill for this price on.
//
// This is an alias to an internal type.
type NewFloatingCumulativeGroupedBulkPriceCadence = shared.NewFloatingCumulativeGroupedBulkPriceCadence

// This is an alias to an internal value.
const NewFloatingCumulativeGroupedBulkPriceCadenceAnnual = shared.NewFloatingCumulativeGroupedBulkPriceCadenceAnnual

// This is an alias to an internal value.
const NewFloatingCumulativeGroupedBulkPriceCadenceSemiAnnual = shared.NewFloatingCumulativeGroupedBulkPriceCadenceSemiAnnual

// This is an alias to an internal value.
const NewFloatingCumulativeGroupedBulkPriceCadenceMonthly = shared.NewFloatingCumulativeGroupedBulkPriceCadenceMonthly

// This is an alias to an internal value.
const NewFloatingCumulativeGroupedBulkPriceCadenceQuarterly = shared.NewFloatingCumulativeGroupedBulkPriceCadenceQuarterly

// This is an alias to an internal value.
const NewFloatingCumulativeGroupedBulkPriceCadenceOneTime = shared.NewFloatingCumulativeGroupedBulkPriceCadenceOneTime

// This is an alias to an internal value.
const NewFloatingCumulativeGroupedBulkPriceCadenceCustom = shared.NewFloatingCumulativeGroupedBulkPriceCadenceCustom

// Configuration for cumulative_grouped_bulk pricing
//
// This is an alias to an internal type.
type NewFloatingCumulativeGroupedBulkPriceCumulativeGroupedBulkConfigParam = shared.NewFloatingCumulativeGroupedBulkPriceCumulativeGroupedBulkConfigParam

// Configuration for a dimension value entry
//
// This is an alias to an internal type.
type NewFloatingCumulativeGroupedBulkPriceCumulativeGroupedBulkConfigDimensionValueParam = shared.NewFloatingCumulativeGroupedBulkPriceCumulativeGroupedBulkConfigDimensionValueParam

// The pricing model type
//
// This is an alias to an internal type.
type NewFloatingCumulativeGroupedBulkPriceModelType = shared.NewFloatingCumulativeGroupedBulkPriceModelType

// This is an alias to an internal value.
const NewFloatingCumulativeGroupedBulkPriceModelTypeCumulativeGroupedBulk = shared.NewFloatingCumulativeGroupedBulkPriceModelTypeCumulativeGroupedBulk

// This is an alias to an internal type.
type NewFloatingCumulativeGroupedBulkPriceConversionRateConfigUnionParam = shared.NewFloatingCumulativeGroupedBulkPriceConversionRateConfigUnionParam

// This is an alias to an internal type.
type NewFloatingCumulativeGroupedBulkPriceConversionRateConfigConversionRateType = shared.NewFloatingCumulativeGroupedBulkPriceConversionRateConfigConversionRateType

// This is an alias to an internal value.
const NewFloatingCumulativeGroupedBulkPriceConversionRateConfigConversionRateTypeUnit = shared.NewFloatingCumulativeGroupedBulkPriceConversionRateConfigConversionRateTypeUnit

// This is an alias to an internal value.
const NewFloatingCumulativeGroupedBulkPriceConversionRateConfigConversionRateTypeTiered = shared.NewFloatingCumulativeGroupedBulkPriceConversionRateConfigConversionRateTypeTiered

// This is an alias to an internal type.
type NewFloatingGroupedAllocationPriceParam = shared.NewFloatingGroupedAllocationPriceParam

// The cadence to bill for this price on.
//
// This is an alias to an internal type.
type NewFloatingGroupedAllocationPriceCadence = shared.NewFloatingGroupedAllocationPriceCadence

// This is an alias to an internal value.
const NewFloatingGroupedAllocationPriceCadenceAnnual = shared.NewFloatingGroupedAllocationPriceCadenceAnnual

// This is an alias to an internal value.
const NewFloatingGroupedAllocationPriceCadenceSemiAnnual = shared.NewFloatingGroupedAllocationPriceCadenceSemiAnnual

// This is an alias to an internal value.
const NewFloatingGroupedAllocationPriceCadenceMonthly = shared.NewFloatingGroupedAllocationPriceCadenceMonthly

// This is an alias to an internal value.
const NewFloatingGroupedAllocationPriceCadenceQuarterly = shared.NewFloatingGroupedAllocationPriceCadenceQuarterly

// This is an alias to an internal value.
const NewFloatingGroupedAllocationPriceCadenceOneTime = shared.NewFloatingGroupedAllocationPriceCadenceOneTime

// This is an alias to an internal value.
const NewFloatingGroupedAllocationPriceCadenceCustom = shared.NewFloatingGroupedAllocationPriceCadenceCustom

// Configuration for grouped_allocation pricing
//
// This is an alias to an internal type.
type NewFloatingGroupedAllocationPriceGroupedAllocationConfigParam = shared.NewFloatingGroupedAllocationPriceGroupedAllocationConfigParam

// The pricing model type
//
// This is an alias to an internal type.
type NewFloatingGroupedAllocationPriceModelType = shared.NewFloatingGroupedAllocationPriceModelType

// This is an alias to an internal value.
const NewFloatingGroupedAllocationPriceModelTypeGroupedAllocation = shared.NewFloatingGroupedAllocationPriceModelTypeGroupedAllocation

// This is an alias to an internal type.
type NewFloatingGroupedAllocationPriceConversionRateConfigUnionParam = shared.NewFloatingGroupedAllocationPriceConversionRateConfigUnionParam

// This is an alias to an internal type.
type NewFloatingGroupedAllocationPriceConversionRateConfigConversionRateType = shared.NewFloatingGroupedAllocationPriceConversionRateConfigConversionRateType

// This is an alias to an internal value.
const NewFloatingGroupedAllocationPriceConversionRateConfigConversionRateTypeUnit = shared.NewFloatingGroupedAllocationPriceConversionRateConfigConversionRateTypeUnit

// This is an alias to an internal value.
const NewFloatingGroupedAllocationPriceConversionRateConfigConversionRateTypeTiered = shared.NewFloatingGroupedAllocationPriceConversionRateConfigConversionRateTypeTiered

// This is an alias to an internal type.
type NewFloatingGroupedTieredPackagePriceParam = shared.NewFloatingGroupedTieredPackagePriceParam

// The cadence to bill for this price on.
//
// This is an alias to an internal type.
type NewFloatingGroupedTieredPackagePriceCadence = shared.NewFloatingGroupedTieredPackagePriceCadence

// This is an alias to an internal value.
const NewFloatingGroupedTieredPackagePriceCadenceAnnual = shared.NewFloatingGroupedTieredPackagePriceCadenceAnnual

// This is an alias to an internal value.
const NewFloatingGroupedTieredPackagePriceCadenceSemiAnnual = shared.NewFloatingGroupedTieredPackagePriceCadenceSemiAnnual

// This is an alias to an internal value.
const NewFloatingGroupedTieredPackagePriceCadenceMonthly = shared.NewFloatingGroupedTieredPackagePriceCadenceMonthly

// This is an alias to an internal value.
const NewFloatingGroupedTieredPackagePriceCadenceQuarterly = shared.NewFloatingGroupedTieredPackagePriceCadenceQuarterly

// This is an alias to an internal value.
const NewFloatingGroupedTieredPackagePriceCadenceOneTime = shared.NewFloatingGroupedTieredPackagePriceCadenceOneTime

// This is an alias to an internal value.
const NewFloatingGroupedTieredPackagePriceCadenceCustom = shared.NewFloatingGroupedTieredPackagePriceCadenceCustom

// Configuration for grouped_tiered_package pricing
//
// This is an alias to an internal type.
type NewFloatingGroupedTieredPackagePriceGroupedTieredPackageConfigParam = shared.NewFloatingGroupedTieredPackagePriceGroupedTieredPackageConfigParam

// Configuration for a single tier
//
// This is an alias to an internal type.
type NewFloatingGroupedTieredPackagePriceGroupedTieredPackageConfigTierParam = shared.NewFloatingGroupedTieredPackagePriceGroupedTieredPackageConfigTierParam

// The pricing model type
//
// This is an alias to an internal type.
type NewFloatingGroupedTieredPackagePriceModelType = shared.NewFloatingGroupedTieredPackagePriceModelType

// This is an alias to an internal value.
const NewFloatingGroupedTieredPackagePriceModelTypeGroupedTieredPackage = shared.NewFloatingGroupedTieredPackagePriceModelTypeGroupedTieredPackage

// This is an alias to an internal type.
type NewFloatingGroupedTieredPackagePriceConversionRateConfigUnionParam = shared.NewFloatingGroupedTieredPackagePriceConversionRateConfigUnionParam

// This is an alias to an internal type.
type NewFloatingGroupedTieredPackagePriceConversionRateConfigConversionRateType = shared.NewFloatingGroupedTieredPackagePriceConversionRateConfigConversionRateType

// This is an alias to an internal value.
const NewFloatingGroupedTieredPackagePriceConversionRateConfigConversionRateTypeUnit = shared.NewFloatingGroupedTieredPackagePriceConversionRateConfigConversionRateTypeUnit

// This is an alias to an internal value.
const NewFloatingGroupedTieredPackagePriceConversionRateConfigConversionRateTypeTiered = shared.NewFloatingGroupedTieredPackagePriceConversionRateConfigConversionRateTypeTiered

// This is an alias to an internal type.
type NewFloatingGroupedTieredPriceParam = shared.NewFloatingGroupedTieredPriceParam

// The cadence to bill for this price on.
//
// This is an alias to an internal type.
type NewFloatingGroupedTieredPriceCadence = shared.NewFloatingGroupedTieredPriceCadence

// This is an alias to an internal value.
const NewFloatingGroupedTieredPriceCadenceAnnual = shared.NewFloatingGroupedTieredPriceCadenceAnnual

// This is an alias to an internal value.
const NewFloatingGroupedTieredPriceCadenceSemiAnnual = shared.NewFloatingGroupedTieredPriceCadenceSemiAnnual

// This is an alias to an internal value.
const NewFloatingGroupedTieredPriceCadenceMonthly = shared.NewFloatingGroupedTieredPriceCadenceMonthly

// This is an alias to an internal value.
const NewFloatingGroupedTieredPriceCadenceQuarterly = shared.NewFloatingGroupedTieredPriceCadenceQuarterly

// This is an alias to an internal value.
const NewFloatingGroupedTieredPriceCadenceOneTime = shared.NewFloatingGroupedTieredPriceCadenceOneTime

// This is an alias to an internal value.
const NewFloatingGroupedTieredPriceCadenceCustom = shared.NewFloatingGroupedTieredPriceCadenceCustom

// Configuration for grouped_tiered pricing
//
// This is an alias to an internal type.
type NewFloatingGroupedTieredPriceGroupedTieredConfigParam = shared.NewFloatingGroupedTieredPriceGroupedTieredConfigParam

// Configuration for a single tier
//
// This is an alias to an internal type.
type NewFloatingGroupedTieredPriceGroupedTieredConfigTierParam = shared.NewFloatingGroupedTieredPriceGroupedTieredConfigTierParam

// The pricing model type
//
// This is an alias to an internal type.
type NewFloatingGroupedTieredPriceModelType = shared.NewFloatingGroupedTieredPriceModelType

// This is an alias to an internal value.
const NewFloatingGroupedTieredPriceModelTypeGroupedTiered = shared.NewFloatingGroupedTieredPriceModelTypeGroupedTiered

// This is an alias to an internal type.
type NewFloatingGroupedTieredPriceConversionRateConfigUnionParam = shared.NewFloatingGroupedTieredPriceConversionRateConfigUnionParam

// This is an alias to an internal type.
type NewFloatingGroupedTieredPriceConversionRateConfigConversionRateType = shared.NewFloatingGroupedTieredPriceConversionRateConfigConversionRateType

// This is an alias to an internal value.
const NewFloatingGroupedTieredPriceConversionRateConfigConversionRateTypeUnit = shared.NewFloatingGroupedTieredPriceConversionRateConfigConversionRateTypeUnit

// This is an alias to an internal value.
const NewFloatingGroupedTieredPriceConversionRateConfigConversionRateTypeTiered = shared.NewFloatingGroupedTieredPriceConversionRateConfigConversionRateTypeTiered

// This is an alias to an internal type.
type NewFloatingGroupedWithMeteredMinimumPriceParam = shared.NewFloatingGroupedWithMeteredMinimumPriceParam

// The cadence to bill for this price on.
//
// This is an alias to an internal type.
type NewFloatingGroupedWithMeteredMinimumPriceCadence = shared.NewFloatingGroupedWithMeteredMinimumPriceCadence

// This is an alias to an internal value.
const NewFloatingGroupedWithMeteredMinimumPriceCadenceAnnual = shared.NewFloatingGroupedWithMeteredMinimumPriceCadenceAnnual

// This is an alias to an internal value.
const NewFloatingGroupedWithMeteredMinimumPriceCadenceSemiAnnual = shared.NewFloatingGroupedWithMeteredMinimumPriceCadenceSemiAnnual

// This is an alias to an internal value.
const NewFloatingGroupedWithMeteredMinimumPriceCadenceMonthly = shared.NewFloatingGroupedWithMeteredMinimumPriceCadenceMonthly

// This is an alias to an internal value.
const NewFloatingGroupedWithMeteredMinimumPriceCadenceQuarterly = shared.NewFloatingGroupedWithMeteredMinimumPriceCadenceQuarterly

// This is an alias to an internal value.
const NewFloatingGroupedWithMeteredMinimumPriceCadenceOneTime = shared.NewFloatingGroupedWithMeteredMinimumPriceCadenceOneTime

// This is an alias to an internal value.
const NewFloatingGroupedWithMeteredMinimumPriceCadenceCustom = shared.NewFloatingGroupedWithMeteredMinimumPriceCadenceCustom

// Configuration for grouped_with_metered_minimum pricing
//
// This is an alias to an internal type.
type NewFloatingGroupedWithMeteredMinimumPriceGroupedWithMeteredMinimumConfigParam = shared.NewFloatingGroupedWithMeteredMinimumPriceGroupedWithMeteredMinimumConfigParam

// Configuration for a scaling factor
//
// This is an alias to an internal type.
type NewFloatingGroupedWithMeteredMinimumPriceGroupedWithMeteredMinimumConfigScalingFactorParam = shared.NewFloatingGroupedWithMeteredMinimumPriceGroupedWithMeteredMinimumConfigScalingFactorParam

// Configuration for a unit amount
//
// This is an alias to an internal type.
type NewFloatingGroupedWithMeteredMinimumPriceGroupedWithMeteredMinimumConfigUnitAmountParam = shared.NewFloatingGroupedWithMeteredMinimumPriceGroupedWithMeteredMinimumConfigUnitAmountParam

// The pricing model type
//
// This is an alias to an internal type.
type NewFloatingGroupedWithMeteredMinimumPriceModelType = shared.NewFloatingGroupedWithMeteredMinimumPriceModelType

// This is an alias to an internal value.
const NewFloatingGroupedWithMeteredMinimumPriceModelTypeGroupedWithMeteredMinimum = shared.NewFloatingGroupedWithMeteredMinimumPriceModelTypeGroupedWithMeteredMinimum

// This is an alias to an internal type.
type NewFloatingGroupedWithMeteredMinimumPriceConversionRateConfigUnionParam = shared.NewFloatingGroupedWithMeteredMinimumPriceConversionRateConfigUnionParam

// This is an alias to an internal type.
type NewFloatingGroupedWithMeteredMinimumPriceConversionRateConfigConversionRateType = shared.NewFloatingGroupedWithMeteredMinimumPriceConversionRateConfigConversionRateType

// This is an alias to an internal value.
const NewFloatingGroupedWithMeteredMinimumPriceConversionRateConfigConversionRateTypeUnit = shared.NewFloatingGroupedWithMeteredMinimumPriceConversionRateConfigConversionRateTypeUnit

// This is an alias to an internal value.
const NewFloatingGroupedWithMeteredMinimumPriceConversionRateConfigConversionRateTypeTiered = shared.NewFloatingGroupedWithMeteredMinimumPriceConversionRateConfigConversionRateTypeTiered

// This is an alias to an internal type.
type NewFloatingGroupedWithProratedMinimumPriceParam = shared.NewFloatingGroupedWithProratedMinimumPriceParam

// The cadence to bill for this price on.
//
// This is an alias to an internal type.
type NewFloatingGroupedWithProratedMinimumPriceCadence = shared.NewFloatingGroupedWithProratedMinimumPriceCadence

// This is an alias to an internal value.
const NewFloatingGroupedWithProratedMinimumPriceCadenceAnnual = shared.NewFloatingGroupedWithProratedMinimumPriceCadenceAnnual

// This is an alias to an internal value.
const NewFloatingGroupedWithProratedMinimumPriceCadenceSemiAnnual = shared.NewFloatingGroupedWithProratedMinimumPriceCadenceSemiAnnual

// This is an alias to an internal value.
const NewFloatingGroupedWithProratedMinimumPriceCadenceMonthly = shared.NewFloatingGroupedWithProratedMinimumPriceCadenceMonthly

// This is an alias to an internal value.
const NewFloatingGroupedWithProratedMinimumPriceCadenceQuarterly = shared.NewFloatingGroupedWithProratedMinimumPriceCadenceQuarterly

// This is an alias to an internal value.
const NewFloatingGroupedWithProratedMinimumPriceCadenceOneTime = shared.NewFloatingGroupedWithProratedMinimumPriceCadenceOneTime

// This is an alias to an internal value.
const NewFloatingGroupedWithProratedMinimumPriceCadenceCustom = shared.NewFloatingGroupedWithProratedMinimumPriceCadenceCustom

// Configuration for grouped_with_prorated_minimum pricing
//
// This is an alias to an internal type.
type NewFloatingGroupedWithProratedMinimumPriceGroupedWithProratedMinimumConfigParam = shared.NewFloatingGroupedWithProratedMinimumPriceGroupedWithProratedMinimumConfigParam

// The pricing model type
//
// This is an alias to an internal type.
type NewFloatingGroupedWithProratedMinimumPriceModelType = shared.NewFloatingGroupedWithProratedMinimumPriceModelType

// This is an alias to an internal value.
const NewFloatingGroupedWithProratedMinimumPriceModelTypeGroupedWithProratedMinimum = shared.NewFloatingGroupedWithProratedMinimumPriceModelTypeGroupedWithProratedMinimum

// This is an alias to an internal type.
type NewFloatingGroupedWithProratedMinimumPriceConversionRateConfigUnionParam = shared.NewFloatingGroupedWithProratedMinimumPriceConversionRateConfigUnionParam

// This is an alias to an internal type.
type NewFloatingGroupedWithProratedMinimumPriceConversionRateConfigConversionRateType = shared.NewFloatingGroupedWithProratedMinimumPriceConversionRateConfigConversionRateType

// This is an alias to an internal value.
const NewFloatingGroupedWithProratedMinimumPriceConversionRateConfigConversionRateTypeUnit = shared.NewFloatingGroupedWithProratedMinimumPriceConversionRateConfigConversionRateTypeUnit

// This is an alias to an internal value.
const NewFloatingGroupedWithProratedMinimumPriceConversionRateConfigConversionRateTypeTiered = shared.NewFloatingGroupedWithProratedMinimumPriceConversionRateConfigConversionRateTypeTiered

// This is an alias to an internal type.
type NewFloatingMatrixPriceParam = shared.NewFloatingMatrixPriceParam

// The cadence to bill for this price on.
//
// This is an alias to an internal type.
type NewFloatingMatrixPriceCadence = shared.NewFloatingMatrixPriceCadence

// This is an alias to an internal value.
const NewFloatingMatrixPriceCadenceAnnual = shared.NewFloatingMatrixPriceCadenceAnnual

// This is an alias to an internal value.
const NewFloatingMatrixPriceCadenceSemiAnnual = shared.NewFloatingMatrixPriceCadenceSemiAnnual

// This is an alias to an internal value.
const NewFloatingMatrixPriceCadenceMonthly = shared.NewFloatingMatrixPriceCadenceMonthly

// This is an alias to an internal value.
const NewFloatingMatrixPriceCadenceQuarterly = shared.NewFloatingMatrixPriceCadenceQuarterly

// This is an alias to an internal value.
const NewFloatingMatrixPriceCadenceOneTime = shared.NewFloatingMatrixPriceCadenceOneTime

// This is an alias to an internal value.
const NewFloatingMatrixPriceCadenceCustom = shared.NewFloatingMatrixPriceCadenceCustom

// The pricing model type
//
// This is an alias to an internal type.
type NewFloatingMatrixPriceModelType = shared.NewFloatingMatrixPriceModelType

// This is an alias to an internal value.
const NewFloatingMatrixPriceModelTypeMatrix = shared.NewFloatingMatrixPriceModelTypeMatrix

// This is an alias to an internal type.
type NewFloatingMatrixPriceConversionRateConfigUnionParam = shared.NewFloatingMatrixPriceConversionRateConfigUnionParam

// This is an alias to an internal type.
type NewFloatingMatrixPriceConversionRateConfigConversionRateType = shared.NewFloatingMatrixPriceConversionRateConfigConversionRateType

// This is an alias to an internal value.
const NewFloatingMatrixPriceConversionRateConfigConversionRateTypeUnit = shared.NewFloatingMatrixPriceConversionRateConfigConversionRateTypeUnit

// This is an alias to an internal value.
const NewFloatingMatrixPriceConversionRateConfigConversionRateTypeTiered = shared.NewFloatingMatrixPriceConversionRateConfigConversionRateTypeTiered

// This is an alias to an internal type.
type NewFloatingMatrixWithAllocationPriceParam = shared.NewFloatingMatrixWithAllocationPriceParam

// The cadence to bill for this price on.
//
// This is an alias to an internal type.
type NewFloatingMatrixWithAllocationPriceCadence = shared.NewFloatingMatrixWithAllocationPriceCadence

// This is an alias to an internal value.
const NewFloatingMatrixWithAllocationPriceCadenceAnnual = shared.NewFloatingMatrixWithAllocationPriceCadenceAnnual

// This is an alias to an internal value.
const NewFloatingMatrixWithAllocationPriceCadenceSemiAnnual = shared.NewFloatingMatrixWithAllocationPriceCadenceSemiAnnual

// This is an alias to an internal value.
const NewFloatingMatrixWithAllocationPriceCadenceMonthly = shared.NewFloatingMatrixWithAllocationPriceCadenceMonthly

// This is an alias to an internal value.
const NewFloatingMatrixWithAllocationPriceCadenceQuarterly = shared.NewFloatingMatrixWithAllocationPriceCadenceQuarterly

// This is an alias to an internal value.
const NewFloatingMatrixWithAllocationPriceCadenceOneTime = shared.NewFloatingMatrixWithAllocationPriceCadenceOneTime

// This is an alias to an internal value.
const NewFloatingMatrixWithAllocationPriceCadenceCustom = shared.NewFloatingMatrixWithAllocationPriceCadenceCustom

// The pricing model type
//
// This is an alias to an internal type.
type NewFloatingMatrixWithAllocationPriceModelType = shared.NewFloatingMatrixWithAllocationPriceModelType

// This is an alias to an internal value.
const NewFloatingMatrixWithAllocationPriceModelTypeMatrixWithAllocation = shared.NewFloatingMatrixWithAllocationPriceModelTypeMatrixWithAllocation

// This is an alias to an internal type.
type NewFloatingMatrixWithAllocationPriceConversionRateConfigUnionParam = shared.NewFloatingMatrixWithAllocationPriceConversionRateConfigUnionParam

// This is an alias to an internal type.
type NewFloatingMatrixWithAllocationPriceConversionRateConfigConversionRateType = shared.NewFloatingMatrixWithAllocationPriceConversionRateConfigConversionRateType

// This is an alias to an internal value.
const NewFloatingMatrixWithAllocationPriceConversionRateConfigConversionRateTypeUnit = shared.NewFloatingMatrixWithAllocationPriceConversionRateConfigConversionRateTypeUnit

// This is an alias to an internal value.
const NewFloatingMatrixWithAllocationPriceConversionRateConfigConversionRateTypeTiered = shared.NewFloatingMatrixWithAllocationPriceConversionRateConfigConversionRateTypeTiered

// This is an alias to an internal type.
type NewFloatingMatrixWithDisplayNamePriceParam = shared.NewFloatingMatrixWithDisplayNamePriceParam

// The cadence to bill for this price on.
//
// This is an alias to an internal type.
type NewFloatingMatrixWithDisplayNamePriceCadence = shared.NewFloatingMatrixWithDisplayNamePriceCadence

// This is an alias to an internal value.
const NewFloatingMatrixWithDisplayNamePriceCadenceAnnual = shared.NewFloatingMatrixWithDisplayNamePriceCadenceAnnual

// This is an alias to an internal value.
const NewFloatingMatrixWithDisplayNamePriceCadenceSemiAnnual = shared.NewFloatingMatrixWithDisplayNamePriceCadenceSemiAnnual

// This is an alias to an internal value.
const NewFloatingMatrixWithDisplayNamePriceCadenceMonthly = shared.NewFloatingMatrixWithDisplayNamePriceCadenceMonthly

// This is an alias to an internal value.
const NewFloatingMatrixWithDisplayNamePriceCadenceQuarterly = shared.NewFloatingMatrixWithDisplayNamePriceCadenceQuarterly

// This is an alias to an internal value.
const NewFloatingMatrixWithDisplayNamePriceCadenceOneTime = shared.NewFloatingMatrixWithDisplayNamePriceCadenceOneTime

// This is an alias to an internal value.
const NewFloatingMatrixWithDisplayNamePriceCadenceCustom = shared.NewFloatingMatrixWithDisplayNamePriceCadenceCustom

// Configuration for matrix_with_display_name pricing
//
// This is an alias to an internal type.
type NewFloatingMatrixWithDisplayNamePriceMatrixWithDisplayNameConfigParam = shared.NewFloatingMatrixWithDisplayNamePriceMatrixWithDisplayNameConfigParam

// Configuration for a unit amount item
//
// This is an alias to an internal type.
type NewFloatingMatrixWithDisplayNamePriceMatrixWithDisplayNameConfigUnitAmountParam = shared.NewFloatingMatrixWithDisplayNamePriceMatrixWithDisplayNameConfigUnitAmountParam

// The pricing model type
//
// This is an alias to an internal type.
type NewFloatingMatrixWithDisplayNamePriceModelType = shared.NewFloatingMatrixWithDisplayNamePriceModelType

// This is an alias to an internal value.
const NewFloatingMatrixWithDisplayNamePriceModelTypeMatrixWithDisplayName = shared.NewFloatingMatrixWithDisplayNamePriceModelTypeMatrixWithDisplayName

// This is an alias to an internal type.
type NewFloatingMatrixWithDisplayNamePriceConversionRateConfigUnionParam = shared.NewFloatingMatrixWithDisplayNamePriceConversionRateConfigUnionParam

// This is an alias to an internal type.
type NewFloatingMatrixWithDisplayNamePriceConversionRateConfigConversionRateType = shared.NewFloatingMatrixWithDisplayNamePriceConversionRateConfigConversionRateType

// This is an alias to an internal value.
const NewFloatingMatrixWithDisplayNamePriceConversionRateConfigConversionRateTypeUnit = shared.NewFloatingMatrixWithDisplayNamePriceConversionRateConfigConversionRateTypeUnit

// This is an alias to an internal value.
const NewFloatingMatrixWithDisplayNamePriceConversionRateConfigConversionRateTypeTiered = shared.NewFloatingMatrixWithDisplayNamePriceConversionRateConfigConversionRateTypeTiered

// This is an alias to an internal type.
type NewFloatingMaxGroupTieredPackagePriceParam = shared.NewFloatingMaxGroupTieredPackagePriceParam

// The cadence to bill for this price on.
//
// This is an alias to an internal type.
type NewFloatingMaxGroupTieredPackagePriceCadence = shared.NewFloatingMaxGroupTieredPackagePriceCadence

// This is an alias to an internal value.
const NewFloatingMaxGroupTieredPackagePriceCadenceAnnual = shared.NewFloatingMaxGroupTieredPackagePriceCadenceAnnual

// This is an alias to an internal value.
const NewFloatingMaxGroupTieredPackagePriceCadenceSemiAnnual = shared.NewFloatingMaxGroupTieredPackagePriceCadenceSemiAnnual

// This is an alias to an internal value.
const NewFloatingMaxGroupTieredPackagePriceCadenceMonthly = shared.NewFloatingMaxGroupTieredPackagePriceCadenceMonthly

// This is an alias to an internal value.
const NewFloatingMaxGroupTieredPackagePriceCadenceQuarterly = shared.NewFloatingMaxGroupTieredPackagePriceCadenceQuarterly

// This is an alias to an internal value.
const NewFloatingMaxGroupTieredPackagePriceCadenceOneTime = shared.NewFloatingMaxGroupTieredPackagePriceCadenceOneTime

// This is an alias to an internal value.
const NewFloatingMaxGroupTieredPackagePriceCadenceCustom = shared.NewFloatingMaxGroupTieredPackagePriceCadenceCustom

// Configuration for max_group_tiered_package pricing
//
// This is an alias to an internal type.
type NewFloatingMaxGroupTieredPackagePriceMaxGroupTieredPackageConfigParam = shared.NewFloatingMaxGroupTieredPackagePriceMaxGroupTieredPackageConfigParam

// Configuration for a single tier
//
// This is an alias to an internal type.
type NewFloatingMaxGroupTieredPackagePriceMaxGroupTieredPackageConfigTierParam = shared.NewFloatingMaxGroupTieredPackagePriceMaxGroupTieredPackageConfigTierParam

// The pricing model type
//
// This is an alias to an internal type.
type NewFloatingMaxGroupTieredPackagePriceModelType = shared.NewFloatingMaxGroupTieredPackagePriceModelType

// This is an alias to an internal value.
const NewFloatingMaxGroupTieredPackagePriceModelTypeMaxGroupTieredPackage = shared.NewFloatingMaxGroupTieredPackagePriceModelTypeMaxGroupTieredPackage

// This is an alias to an internal type.
type NewFloatingMaxGroupTieredPackagePriceConversionRateConfigUnionParam = shared.NewFloatingMaxGroupTieredPackagePriceConversionRateConfigUnionParam

// This is an alias to an internal type.
type NewFloatingMaxGroupTieredPackagePriceConversionRateConfigConversionRateType = shared.NewFloatingMaxGroupTieredPackagePriceConversionRateConfigConversionRateType

// This is an alias to an internal value.
const NewFloatingMaxGroupTieredPackagePriceConversionRateConfigConversionRateTypeUnit = shared.NewFloatingMaxGroupTieredPackagePriceConversionRateConfigConversionRateTypeUnit

// This is an alias to an internal value.
const NewFloatingMaxGroupTieredPackagePriceConversionRateConfigConversionRateTypeTiered = shared.NewFloatingMaxGroupTieredPackagePriceConversionRateConfigConversionRateTypeTiered

// This is an alias to an internal type.
type NewFloatingMinimumCompositePriceParam = shared.NewFloatingMinimumCompositePriceParam

// The cadence to bill for this price on.
//
// This is an alias to an internal type.
type NewFloatingMinimumCompositePriceCadence = shared.NewFloatingMinimumCompositePriceCadence

// This is an alias to an internal value.
const NewFloatingMinimumCompositePriceCadenceAnnual = shared.NewFloatingMinimumCompositePriceCadenceAnnual

// This is an alias to an internal value.
const NewFloatingMinimumCompositePriceCadenceSemiAnnual = shared.NewFloatingMinimumCompositePriceCadenceSemiAnnual

// This is an alias to an internal value.
const NewFloatingMinimumCompositePriceCadenceMonthly = shared.NewFloatingMinimumCompositePriceCadenceMonthly

// This is an alias to an internal value.
const NewFloatingMinimumCompositePriceCadenceQuarterly = shared.NewFloatingMinimumCompositePriceCadenceQuarterly

// This is an alias to an internal value.
const NewFloatingMinimumCompositePriceCadenceOneTime = shared.NewFloatingMinimumCompositePriceCadenceOneTime

// This is an alias to an internal value.
const NewFloatingMinimumCompositePriceCadenceCustom = shared.NewFloatingMinimumCompositePriceCadenceCustom

// Configuration for minimum_composite pricing
//
// This is an alias to an internal type.
type NewFloatingMinimumCompositePriceMinimumCompositeConfigParam = shared.NewFloatingMinimumCompositePriceMinimumCompositeConfigParam

// The pricing model type
//
// This is an alias to an internal type.
type NewFloatingMinimumCompositePriceModelType = shared.NewFloatingMinimumCompositePriceModelType

// This is an alias to an internal value.
const NewFloatingMinimumCompositePriceModelTypeMinimumComposite = shared.NewFloatingMinimumCompositePriceModelTypeMinimumComposite

// This is an alias to an internal type.
type NewFloatingMinimumCompositePriceConversionRateConfigUnionParam = shared.NewFloatingMinimumCompositePriceConversionRateConfigUnionParam

// This is an alias to an internal type.
type NewFloatingMinimumCompositePriceConversionRateConfigConversionRateType = shared.NewFloatingMinimumCompositePriceConversionRateConfigConversionRateType

// This is an alias to an internal value.
const NewFloatingMinimumCompositePriceConversionRateConfigConversionRateTypeUnit = shared.NewFloatingMinimumCompositePriceConversionRateConfigConversionRateTypeUnit

// This is an alias to an internal value.
const NewFloatingMinimumCompositePriceConversionRateConfigConversionRateTypeTiered = shared.NewFloatingMinimumCompositePriceConversionRateConfigConversionRateTypeTiered

// This is an alias to an internal type.
type NewFloatingPackagePriceParam = shared.NewFloatingPackagePriceParam

// The cadence to bill for this price on.
//
// This is an alias to an internal type.
type NewFloatingPackagePriceCadence = shared.NewFloatingPackagePriceCadence

// This is an alias to an internal value.
const NewFloatingPackagePriceCadenceAnnual = shared.NewFloatingPackagePriceCadenceAnnual

// This is an alias to an internal value.
const NewFloatingPackagePriceCadenceSemiAnnual = shared.NewFloatingPackagePriceCadenceSemiAnnual

// This is an alias to an internal value.
const NewFloatingPackagePriceCadenceMonthly = shared.NewFloatingPackagePriceCadenceMonthly

// This is an alias to an internal value.
const NewFloatingPackagePriceCadenceQuarterly = shared.NewFloatingPackagePriceCadenceQuarterly

// This is an alias to an internal value.
const NewFloatingPackagePriceCadenceOneTime = shared.NewFloatingPackagePriceCadenceOneTime

// This is an alias to an internal value.
const NewFloatingPackagePriceCadenceCustom = shared.NewFloatingPackagePriceCadenceCustom

// The pricing model type
//
// This is an alias to an internal type.
type NewFloatingPackagePriceModelType = shared.NewFloatingPackagePriceModelType

// This is an alias to an internal value.
const NewFloatingPackagePriceModelTypePackage = shared.NewFloatingPackagePriceModelTypePackage

// This is an alias to an internal type.
type NewFloatingPackagePriceConversionRateConfigUnionParam = shared.NewFloatingPackagePriceConversionRateConfigUnionParam

// This is an alias to an internal type.
type NewFloatingPackagePriceConversionRateConfigConversionRateType = shared.NewFloatingPackagePriceConversionRateConfigConversionRateType

// This is an alias to an internal value.
const NewFloatingPackagePriceConversionRateConfigConversionRateTypeUnit = shared.NewFloatingPackagePriceConversionRateConfigConversionRateTypeUnit

// This is an alias to an internal value.
const NewFloatingPackagePriceConversionRateConfigConversionRateTypeTiered = shared.NewFloatingPackagePriceConversionRateConfigConversionRateTypeTiered

// This is an alias to an internal type.
type NewFloatingPackageWithAllocationPriceParam = shared.NewFloatingPackageWithAllocationPriceParam

// The cadence to bill for this price on.
//
// This is an alias to an internal type.
type NewFloatingPackageWithAllocationPriceCadence = shared.NewFloatingPackageWithAllocationPriceCadence

// This is an alias to an internal value.
const NewFloatingPackageWithAllocationPriceCadenceAnnual = shared.NewFloatingPackageWithAllocationPriceCadenceAnnual

// This is an alias to an internal value.
const NewFloatingPackageWithAllocationPriceCadenceSemiAnnual = shared.NewFloatingPackageWithAllocationPriceCadenceSemiAnnual

// This is an alias to an internal value.
const NewFloatingPackageWithAllocationPriceCadenceMonthly = shared.NewFloatingPackageWithAllocationPriceCadenceMonthly

// This is an alias to an internal value.
const NewFloatingPackageWithAllocationPriceCadenceQuarterly = shared.NewFloatingPackageWithAllocationPriceCadenceQuarterly

// This is an alias to an internal value.
const NewFloatingPackageWithAllocationPriceCadenceOneTime = shared.NewFloatingPackageWithAllocationPriceCadenceOneTime

// This is an alias to an internal value.
const NewFloatingPackageWithAllocationPriceCadenceCustom = shared.NewFloatingPackageWithAllocationPriceCadenceCustom

// The pricing model type
//
// This is an alias to an internal type.
type NewFloatingPackageWithAllocationPriceModelType = shared.NewFloatingPackageWithAllocationPriceModelType

// This is an alias to an internal value.
const NewFloatingPackageWithAllocationPriceModelTypePackageWithAllocation = shared.NewFloatingPackageWithAllocationPriceModelTypePackageWithAllocation

// Configuration for package_with_allocation pricing
//
// This is an alias to an internal type.
type NewFloatingPackageWithAllocationPricePackageWithAllocationConfigParam = shared.NewFloatingPackageWithAllocationPricePackageWithAllocationConfigParam

// This is an alias to an internal type.
type NewFloatingPackageWithAllocationPriceConversionRateConfigUnionParam = shared.NewFloatingPackageWithAllocationPriceConversionRateConfigUnionParam

// This is an alias to an internal type.
type NewFloatingPackageWithAllocationPriceConversionRateConfigConversionRateType = shared.NewFloatingPackageWithAllocationPriceConversionRateConfigConversionRateType

// This is an alias to an internal value.
const NewFloatingPackageWithAllocationPriceConversionRateConfigConversionRateTypeUnit = shared.NewFloatingPackageWithAllocationPriceConversionRateConfigConversionRateTypeUnit

// This is an alias to an internal value.
const NewFloatingPackageWithAllocationPriceConversionRateConfigConversionRateTypeTiered = shared.NewFloatingPackageWithAllocationPriceConversionRateConfigConversionRateTypeTiered

// This is an alias to an internal type.
type NewFloatingScalableMatrixWithTieredPricingPriceParam = shared.NewFloatingScalableMatrixWithTieredPricingPriceParam

// The cadence to bill for this price on.
//
// This is an alias to an internal type.
type NewFloatingScalableMatrixWithTieredPricingPriceCadence = shared.NewFloatingScalableMatrixWithTieredPricingPriceCadence

// This is an alias to an internal value.
const NewFloatingScalableMatrixWithTieredPricingPriceCadenceAnnual = shared.NewFloatingScalableMatrixWithTieredPricingPriceCadenceAnnual

// This is an alias to an internal value.
const NewFloatingScalableMatrixWithTieredPricingPriceCadenceSemiAnnual = shared.NewFloatingScalableMatrixWithTieredPricingPriceCadenceSemiAnnual

// This is an alias to an internal value.
const NewFloatingScalableMatrixWithTieredPricingPriceCadenceMonthly = shared.NewFloatingScalableMatrixWithTieredPricingPriceCadenceMonthly

// This is an alias to an internal value.
const NewFloatingScalableMatrixWithTieredPricingPriceCadenceQuarterly = shared.NewFloatingScalableMatrixWithTieredPricingPriceCadenceQuarterly

// This is an alias to an internal value.
const NewFloatingScalableMatrixWithTieredPricingPriceCadenceOneTime = shared.NewFloatingScalableMatrixWithTieredPricingPriceCadenceOneTime

// This is an alias to an internal value.
const NewFloatingScalableMatrixWithTieredPricingPriceCadenceCustom = shared.NewFloatingScalableMatrixWithTieredPricingPriceCadenceCustom

// The pricing model type
//
// This is an alias to an internal type.
type NewFloatingScalableMatrixWithTieredPricingPriceModelType = shared.NewFloatingScalableMatrixWithTieredPricingPriceModelType

// This is an alias to an internal value.
const NewFloatingScalableMatrixWithTieredPricingPriceModelTypeScalableMatrixWithTieredPricing = shared.NewFloatingScalableMatrixWithTieredPricingPriceModelTypeScalableMatrixWithTieredPricing

// Configuration for scalable_matrix_with_tiered_pricing pricing
//
// This is an alias to an internal type.
type NewFloatingScalableMatrixWithTieredPricingPriceScalableMatrixWithTieredPricingConfigParam = shared.NewFloatingScalableMatrixWithTieredPricingPriceScalableMatrixWithTieredPricingConfigParam

// Configuration for a single matrix scaling factor
//
// This is an alias to an internal type.
type NewFloatingScalableMatrixWithTieredPricingPriceScalableMatrixWithTieredPricingConfigMatrixScalingFactorParam = shared.NewFloatingScalableMatrixWithTieredPricingPriceScalableMatrixWithTieredPricingConfigMatrixScalingFactorParam

// Configuration for a single tier entry with business logic
//
// This is an alias to an internal type.
type NewFloatingScalableMatrixWithTieredPricingPriceScalableMatrixWithTieredPricingConfigTierParam = shared.NewFloatingScalableMatrixWithTieredPricingPriceScalableMatrixWithTieredPricingConfigTierParam

// This is an alias to an internal type.
type NewFloatingScalableMatrixWithTieredPricingPriceConversionRateConfigUnionParam = shared.NewFloatingScalableMatrixWithTieredPricingPriceConversionRateConfigUnionParam

// This is an alias to an internal type.
type NewFloatingScalableMatrixWithTieredPricingPriceConversionRateConfigConversionRateType = shared.NewFloatingScalableMatrixWithTieredPricingPriceConversionRateConfigConversionRateType

// This is an alias to an internal value.
const NewFloatingScalableMatrixWithTieredPricingPriceConversionRateConfigConversionRateTypeUnit = shared.NewFloatingScalableMatrixWithTieredPricingPriceConversionRateConfigConversionRateTypeUnit

// This is an alias to an internal value.
const NewFloatingScalableMatrixWithTieredPricingPriceConversionRateConfigConversionRateTypeTiered = shared.NewFloatingScalableMatrixWithTieredPricingPriceConversionRateConfigConversionRateTypeTiered

// This is an alias to an internal type.
type NewFloatingScalableMatrixWithUnitPricingPriceParam = shared.NewFloatingScalableMatrixWithUnitPricingPriceParam

// The cadence to bill for this price on.
//
// This is an alias to an internal type.
type NewFloatingScalableMatrixWithUnitPricingPriceCadence = shared.NewFloatingScalableMatrixWithUnitPricingPriceCadence

// This is an alias to an internal value.
const NewFloatingScalableMatrixWithUnitPricingPriceCadenceAnnual = shared.NewFloatingScalableMatrixWithUnitPricingPriceCadenceAnnual

// This is an alias to an internal value.
const NewFloatingScalableMatrixWithUnitPricingPriceCadenceSemiAnnual = shared.NewFloatingScalableMatrixWithUnitPricingPriceCadenceSemiAnnual

// This is an alias to an internal value.
const NewFloatingScalableMatrixWithUnitPricingPriceCadenceMonthly = shared.NewFloatingScalableMatrixWithUnitPricingPriceCadenceMonthly

// This is an alias to an internal value.
const NewFloatingScalableMatrixWithUnitPricingPriceCadenceQuarterly = shared.NewFloatingScalableMatrixWithUnitPricingPriceCadenceQuarterly

// This is an alias to an internal value.
const NewFloatingScalableMatrixWithUnitPricingPriceCadenceOneTime = shared.NewFloatingScalableMatrixWithUnitPricingPriceCadenceOneTime

// This is an alias to an internal value.
const NewFloatingScalableMatrixWithUnitPricingPriceCadenceCustom = shared.NewFloatingScalableMatrixWithUnitPricingPriceCadenceCustom

// The pricing model type
//
// This is an alias to an internal type.
type NewFloatingScalableMatrixWithUnitPricingPriceModelType = shared.NewFloatingScalableMatrixWithUnitPricingPriceModelType

// This is an alias to an internal value.
const NewFloatingScalableMatrixWithUnitPricingPriceModelTypeScalableMatrixWithUnitPricing = shared.NewFloatingScalableMatrixWithUnitPricingPriceModelTypeScalableMatrixWithUnitPricing

// Configuration for scalable_matrix_with_unit_pricing pricing
//
// This is an alias to an internal type.
type NewFloatingScalableMatrixWithUnitPricingPriceScalableMatrixWithUnitPricingConfigParam = shared.NewFloatingScalableMatrixWithUnitPricingPriceScalableMatrixWithUnitPricingConfigParam

// Configuration for a single matrix scaling factor
//
// This is an alias to an internal type.
type NewFloatingScalableMatrixWithUnitPricingPriceScalableMatrixWithUnitPricingConfigMatrixScalingFactorParam = shared.NewFloatingScalableMatrixWithUnitPricingPriceScalableMatrixWithUnitPricingConfigMatrixScalingFactorParam

// This is an alias to an internal type.
type NewFloatingScalableMatrixWithUnitPricingPriceConversionRateConfigUnionParam = shared.NewFloatingScalableMatrixWithUnitPricingPriceConversionRateConfigUnionParam

// This is an alias to an internal type.
type NewFloatingScalableMatrixWithUnitPricingPriceConversionRateConfigConversionRateType = shared.NewFloatingScalableMatrixWithUnitPricingPriceConversionRateConfigConversionRateType

// This is an alias to an internal value.
const NewFloatingScalableMatrixWithUnitPricingPriceConversionRateConfigConversionRateTypeUnit = shared.NewFloatingScalableMatrixWithUnitPricingPriceConversionRateConfigConversionRateTypeUnit

// This is an alias to an internal value.
const NewFloatingScalableMatrixWithUnitPricingPriceConversionRateConfigConversionRateTypeTiered = shared.NewFloatingScalableMatrixWithUnitPricingPriceConversionRateConfigConversionRateTypeTiered

// This is an alias to an internal type.
type NewFloatingThresholdTotalAmountPriceParam = shared.NewFloatingThresholdTotalAmountPriceParam

// The cadence to bill for this price on.
//
// This is an alias to an internal type.
type NewFloatingThresholdTotalAmountPriceCadence = shared.NewFloatingThresholdTotalAmountPriceCadence

// This is an alias to an internal value.
const NewFloatingThresholdTotalAmountPriceCadenceAnnual = shared.NewFloatingThresholdTotalAmountPriceCadenceAnnual

// This is an alias to an internal value.
const NewFloatingThresholdTotalAmountPriceCadenceSemiAnnual = shared.NewFloatingThresholdTotalAmountPriceCadenceSemiAnnual

// This is an alias to an internal value.
const NewFloatingThresholdTotalAmountPriceCadenceMonthly = shared.NewFloatingThresholdTotalAmountPriceCadenceMonthly

// This is an alias to an internal value.
const NewFloatingThresholdTotalAmountPriceCadenceQuarterly = shared.NewFloatingThresholdTotalAmountPriceCadenceQuarterly

// This is an alias to an internal value.
const NewFloatingThresholdTotalAmountPriceCadenceOneTime = shared.NewFloatingThresholdTotalAmountPriceCadenceOneTime

// This is an alias to an internal value.
const NewFloatingThresholdTotalAmountPriceCadenceCustom = shared.NewFloatingThresholdTotalAmountPriceCadenceCustom

// The pricing model type
//
// This is an alias to an internal type.
type NewFloatingThresholdTotalAmountPriceModelType = shared.NewFloatingThresholdTotalAmountPriceModelType

// This is an alias to an internal value.
const NewFloatingThresholdTotalAmountPriceModelTypeThresholdTotalAmount = shared.NewFloatingThresholdTotalAmountPriceModelTypeThresholdTotalAmount

// Configuration for threshold_total_amount pricing
//
// This is an alias to an internal type.
type NewFloatingThresholdTotalAmountPriceThresholdTotalAmountConfigParam = shared.NewFloatingThresholdTotalAmountPriceThresholdTotalAmountConfigParam

// Configuration for a single threshold
//
// This is an alias to an internal type.
type NewFloatingThresholdTotalAmountPriceThresholdTotalAmountConfigConsumptionTableParam = shared.NewFloatingThresholdTotalAmountPriceThresholdTotalAmountConfigConsumptionTableParam

// This is an alias to an internal type.
type NewFloatingThresholdTotalAmountPriceConversionRateConfigUnionParam = shared.NewFloatingThresholdTotalAmountPriceConversionRateConfigUnionParam

// This is an alias to an internal type.
type NewFloatingThresholdTotalAmountPriceConversionRateConfigConversionRateType = shared.NewFloatingThresholdTotalAmountPriceConversionRateConfigConversionRateType

// This is an alias to an internal value.
const NewFloatingThresholdTotalAmountPriceConversionRateConfigConversionRateTypeUnit = shared.NewFloatingThresholdTotalAmountPriceConversionRateConfigConversionRateTypeUnit

// This is an alias to an internal value.
const NewFloatingThresholdTotalAmountPriceConversionRateConfigConversionRateTypeTiered = shared.NewFloatingThresholdTotalAmountPriceConversionRateConfigConversionRateTypeTiered

// This is an alias to an internal type.
type NewFloatingTieredPackagePriceParam = shared.NewFloatingTieredPackagePriceParam

// The cadence to bill for this price on.
//
// This is an alias to an internal type.
type NewFloatingTieredPackagePriceCadence = shared.NewFloatingTieredPackagePriceCadence

// This is an alias to an internal value.
const NewFloatingTieredPackagePriceCadenceAnnual = shared.NewFloatingTieredPackagePriceCadenceAnnual

// This is an alias to an internal value.
const NewFloatingTieredPackagePriceCadenceSemiAnnual = shared.NewFloatingTieredPackagePriceCadenceSemiAnnual

// This is an alias to an internal value.
const NewFloatingTieredPackagePriceCadenceMonthly = shared.NewFloatingTieredPackagePriceCadenceMonthly

// This is an alias to an internal value.
const NewFloatingTieredPackagePriceCadenceQuarterly = shared.NewFloatingTieredPackagePriceCadenceQuarterly

// This is an alias to an internal value.
const NewFloatingTieredPackagePriceCadenceOneTime = shared.NewFloatingTieredPackagePriceCadenceOneTime

// This is an alias to an internal value.
const NewFloatingTieredPackagePriceCadenceCustom = shared.NewFloatingTieredPackagePriceCadenceCustom

// The pricing model type
//
// This is an alias to an internal type.
type NewFloatingTieredPackagePriceModelType = shared.NewFloatingTieredPackagePriceModelType

// This is an alias to an internal value.
const NewFloatingTieredPackagePriceModelTypeTieredPackage = shared.NewFloatingTieredPackagePriceModelTypeTieredPackage

// Configuration for tiered_package pricing
//
// This is an alias to an internal type.
type NewFloatingTieredPackagePriceTieredPackageConfigParam = shared.NewFloatingTieredPackagePriceTieredPackageConfigParam

// Configuration for a single tier with business logic
//
// This is an alias to an internal type.
type NewFloatingTieredPackagePriceTieredPackageConfigTierParam = shared.NewFloatingTieredPackagePriceTieredPackageConfigTierParam

// This is an alias to an internal type.
type NewFloatingTieredPackagePriceConversionRateConfigUnionParam = shared.NewFloatingTieredPackagePriceConversionRateConfigUnionParam

// This is an alias to an internal type.
type NewFloatingTieredPackagePriceConversionRateConfigConversionRateType = shared.NewFloatingTieredPackagePriceConversionRateConfigConversionRateType

// This is an alias to an internal value.
const NewFloatingTieredPackagePriceConversionRateConfigConversionRateTypeUnit = shared.NewFloatingTieredPackagePriceConversionRateConfigConversionRateTypeUnit

// This is an alias to an internal value.
const NewFloatingTieredPackagePriceConversionRateConfigConversionRateTypeTiered = shared.NewFloatingTieredPackagePriceConversionRateConfigConversionRateTypeTiered

// This is an alias to an internal type.
type NewFloatingTieredPackageWithMinimumPriceParam = shared.NewFloatingTieredPackageWithMinimumPriceParam

// The cadence to bill for this price on.
//
// This is an alias to an internal type.
type NewFloatingTieredPackageWithMinimumPriceCadence = shared.NewFloatingTieredPackageWithMinimumPriceCadence

// This is an alias to an internal value.
const NewFloatingTieredPackageWithMinimumPriceCadenceAnnual = shared.NewFloatingTieredPackageWithMinimumPriceCadenceAnnual

// This is an alias to an internal value.
const NewFloatingTieredPackageWithMinimumPriceCadenceSemiAnnual = shared.NewFloatingTieredPackageWithMinimumPriceCadenceSemiAnnual

// This is an alias to an internal value.
const NewFloatingTieredPackageWithMinimumPriceCadenceMonthly = shared.NewFloatingTieredPackageWithMinimumPriceCadenceMonthly

// This is an alias to an internal value.
const NewFloatingTieredPackageWithMinimumPriceCadenceQuarterly = shared.NewFloatingTieredPackageWithMinimumPriceCadenceQuarterly

// This is an alias to an internal value.
const NewFloatingTieredPackageWithMinimumPriceCadenceOneTime = shared.NewFloatingTieredPackageWithMinimumPriceCadenceOneTime

// This is an alias to an internal value.
const NewFloatingTieredPackageWithMinimumPriceCadenceCustom = shared.NewFloatingTieredPackageWithMinimumPriceCadenceCustom

// The pricing model type
//
// This is an alias to an internal type.
type NewFloatingTieredPackageWithMinimumPriceModelType = shared.NewFloatingTieredPackageWithMinimumPriceModelType

// This is an alias to an internal value.
const NewFloatingTieredPackageWithMinimumPriceModelTypeTieredPackageWithMinimum = shared.NewFloatingTieredPackageWithMinimumPriceModelTypeTieredPackageWithMinimum

// Configuration for tiered_package_with_minimum pricing
//
// This is an alias to an internal type.
type NewFloatingTieredPackageWithMinimumPriceTieredPackageWithMinimumConfigParam = shared.NewFloatingTieredPackageWithMinimumPriceTieredPackageWithMinimumConfigParam

// Configuration for a single tier
//
// This is an alias to an internal type.
type NewFloatingTieredPackageWithMinimumPriceTieredPackageWithMinimumConfigTierParam = shared.NewFloatingTieredPackageWithMinimumPriceTieredPackageWithMinimumConfigTierParam

// This is an alias to an internal type.
type NewFloatingTieredPackageWithMinimumPriceConversionRateConfigUnionParam = shared.NewFloatingTieredPackageWithMinimumPriceConversionRateConfigUnionParam

// This is an alias to an internal type.
type NewFloatingTieredPackageWithMinimumPriceConversionRateConfigConversionRateType = shared.NewFloatingTieredPackageWithMinimumPriceConversionRateConfigConversionRateType

// This is an alias to an internal value.
const NewFloatingTieredPackageWithMinimumPriceConversionRateConfigConversionRateTypeUnit = shared.NewFloatingTieredPackageWithMinimumPriceConversionRateConfigConversionRateTypeUnit

// This is an alias to an internal value.
const NewFloatingTieredPackageWithMinimumPriceConversionRateConfigConversionRateTypeTiered = shared.NewFloatingTieredPackageWithMinimumPriceConversionRateConfigConversionRateTypeTiered

// This is an alias to an internal type.
type NewFloatingTieredPriceParam = shared.NewFloatingTieredPriceParam

// The cadence to bill for this price on.
//
// This is an alias to an internal type.
type NewFloatingTieredPriceCadence = shared.NewFloatingTieredPriceCadence

// This is an alias to an internal value.
const NewFloatingTieredPriceCadenceAnnual = shared.NewFloatingTieredPriceCadenceAnnual

// This is an alias to an internal value.
const NewFloatingTieredPriceCadenceSemiAnnual = shared.NewFloatingTieredPriceCadenceSemiAnnual

// This is an alias to an internal value.
const NewFloatingTieredPriceCadenceMonthly = shared.NewFloatingTieredPriceCadenceMonthly

// This is an alias to an internal value.
const NewFloatingTieredPriceCadenceQuarterly = shared.NewFloatingTieredPriceCadenceQuarterly

// This is an alias to an internal value.
const NewFloatingTieredPriceCadenceOneTime = shared.NewFloatingTieredPriceCadenceOneTime

// This is an alias to an internal value.
const NewFloatingTieredPriceCadenceCustom = shared.NewFloatingTieredPriceCadenceCustom

// The pricing model type
//
// This is an alias to an internal type.
type NewFloatingTieredPriceModelType = shared.NewFloatingTieredPriceModelType

// This is an alias to an internal value.
const NewFloatingTieredPriceModelTypeTiered = shared.NewFloatingTieredPriceModelTypeTiered

// This is an alias to an internal type.
type NewFloatingTieredPriceConversionRateConfigUnionParam = shared.NewFloatingTieredPriceConversionRateConfigUnionParam

// This is an alias to an internal type.
type NewFloatingTieredPriceConversionRateConfigConversionRateType = shared.NewFloatingTieredPriceConversionRateConfigConversionRateType

// This is an alias to an internal value.
const NewFloatingTieredPriceConversionRateConfigConversionRateTypeUnit = shared.NewFloatingTieredPriceConversionRateConfigConversionRateTypeUnit

// This is an alias to an internal value.
const NewFloatingTieredPriceConversionRateConfigConversionRateTypeTiered = shared.NewFloatingTieredPriceConversionRateConfigConversionRateTypeTiered

// This is an alias to an internal type.
type NewFloatingTieredWithMinimumPriceParam = shared.NewFloatingTieredWithMinimumPriceParam

// The cadence to bill for this price on.
//
// This is an alias to an internal type.
type NewFloatingTieredWithMinimumPriceCadence = shared.NewFloatingTieredWithMinimumPriceCadence

// This is an alias to an internal value.
const NewFloatingTieredWithMinimumPriceCadenceAnnual = shared.NewFloatingTieredWithMinimumPriceCadenceAnnual

// This is an alias to an internal value.
const NewFloatingTieredWithMinimumPriceCadenceSemiAnnual = shared.NewFloatingTieredWithMinimumPriceCadenceSemiAnnual

// This is an alias to an internal value.
const NewFloatingTieredWithMinimumPriceCadenceMonthly = shared.NewFloatingTieredWithMinimumPriceCadenceMonthly

// This is an alias to an internal value.
const NewFloatingTieredWithMinimumPriceCadenceQuarterly = shared.NewFloatingTieredWithMinimumPriceCadenceQuarterly

// This is an alias to an internal value.
const NewFloatingTieredWithMinimumPriceCadenceOneTime = shared.NewFloatingTieredWithMinimumPriceCadenceOneTime

// This is an alias to an internal value.
const NewFloatingTieredWithMinimumPriceCadenceCustom = shared.NewFloatingTieredWithMinimumPriceCadenceCustom

// The pricing model type
//
// This is an alias to an internal type.
type NewFloatingTieredWithMinimumPriceModelType = shared.NewFloatingTieredWithMinimumPriceModelType

// This is an alias to an internal value.
const NewFloatingTieredWithMinimumPriceModelTypeTieredWithMinimum = shared.NewFloatingTieredWithMinimumPriceModelTypeTieredWithMinimum

// Configuration for tiered_with_minimum pricing
//
// This is an alias to an internal type.
type NewFloatingTieredWithMinimumPriceTieredWithMinimumConfigParam = shared.NewFloatingTieredWithMinimumPriceTieredWithMinimumConfigParam

// Configuration for a single tier
//
// This is an alias to an internal type.
type NewFloatingTieredWithMinimumPriceTieredWithMinimumConfigTierParam = shared.NewFloatingTieredWithMinimumPriceTieredWithMinimumConfigTierParam

// This is an alias to an internal type.
type NewFloatingTieredWithMinimumPriceConversionRateConfigUnionParam = shared.NewFloatingTieredWithMinimumPriceConversionRateConfigUnionParam

// This is an alias to an internal type.
type NewFloatingTieredWithMinimumPriceConversionRateConfigConversionRateType = shared.NewFloatingTieredWithMinimumPriceConversionRateConfigConversionRateType

// This is an alias to an internal value.
const NewFloatingTieredWithMinimumPriceConversionRateConfigConversionRateTypeUnit = shared.NewFloatingTieredWithMinimumPriceConversionRateConfigConversionRateTypeUnit

// This is an alias to an internal value.
const NewFloatingTieredWithMinimumPriceConversionRateConfigConversionRateTypeTiered = shared.NewFloatingTieredWithMinimumPriceConversionRateConfigConversionRateTypeTiered

// This is an alias to an internal type.
type NewFloatingTieredWithProrationPriceParam = shared.NewFloatingTieredWithProrationPriceParam

// The cadence to bill for this price on.
//
// This is an alias to an internal type.
type NewFloatingTieredWithProrationPriceCadence = shared.NewFloatingTieredWithProrationPriceCadence

// This is an alias to an internal value.
const NewFloatingTieredWithProrationPriceCadenceAnnual = shared.NewFloatingTieredWithProrationPriceCadenceAnnual

// This is an alias to an internal value.
const NewFloatingTieredWithProrationPriceCadenceSemiAnnual = shared.NewFloatingTieredWithProrationPriceCadenceSemiAnnual

// This is an alias to an internal value.
const NewFloatingTieredWithProrationPriceCadenceMonthly = shared.NewFloatingTieredWithProrationPriceCadenceMonthly

// This is an alias to an internal value.
const NewFloatingTieredWithProrationPriceCadenceQuarterly = shared.NewFloatingTieredWithProrationPriceCadenceQuarterly

// This is an alias to an internal value.
const NewFloatingTieredWithProrationPriceCadenceOneTime = shared.NewFloatingTieredWithProrationPriceCadenceOneTime

// This is an alias to an internal value.
const NewFloatingTieredWithProrationPriceCadenceCustom = shared.NewFloatingTieredWithProrationPriceCadenceCustom

// The pricing model type
//
// This is an alias to an internal type.
type NewFloatingTieredWithProrationPriceModelType = shared.NewFloatingTieredWithProrationPriceModelType

// This is an alias to an internal value.
const NewFloatingTieredWithProrationPriceModelTypeTieredWithProration = shared.NewFloatingTieredWithProrationPriceModelTypeTieredWithProration

// Configuration for tiered_with_proration pricing
//
// This is an alias to an internal type.
type NewFloatingTieredWithProrationPriceTieredWithProrationConfigParam = shared.NewFloatingTieredWithProrationPriceTieredWithProrationConfigParam

// Configuration for a single tiered with proration tier
//
// This is an alias to an internal type.
type NewFloatingTieredWithProrationPriceTieredWithProrationConfigTierParam = shared.NewFloatingTieredWithProrationPriceTieredWithProrationConfigTierParam

// This is an alias to an internal type.
type NewFloatingTieredWithProrationPriceConversionRateConfigUnionParam = shared.NewFloatingTieredWithProrationPriceConversionRateConfigUnionParam

// This is an alias to an internal type.
type NewFloatingTieredWithProrationPriceConversionRateConfigConversionRateType = shared.NewFloatingTieredWithProrationPriceConversionRateConfigConversionRateType

// This is an alias to an internal value.
const NewFloatingTieredWithProrationPriceConversionRateConfigConversionRateTypeUnit = shared.NewFloatingTieredWithProrationPriceConversionRateConfigConversionRateTypeUnit

// This is an alias to an internal value.
const NewFloatingTieredWithProrationPriceConversionRateConfigConversionRateTypeTiered = shared.NewFloatingTieredWithProrationPriceConversionRateConfigConversionRateTypeTiered

// This is an alias to an internal type.
type NewFloatingUnitPriceParam = shared.NewFloatingUnitPriceParam

// The cadence to bill for this price on.
//
// This is an alias to an internal type.
type NewFloatingUnitPriceCadence = shared.NewFloatingUnitPriceCadence

// This is an alias to an internal value.
const NewFloatingUnitPriceCadenceAnnual = shared.NewFloatingUnitPriceCadenceAnnual

// This is an alias to an internal value.
const NewFloatingUnitPriceCadenceSemiAnnual = shared.NewFloatingUnitPriceCadenceSemiAnnual

// This is an alias to an internal value.
const NewFloatingUnitPriceCadenceMonthly = shared.NewFloatingUnitPriceCadenceMonthly

// This is an alias to an internal value.
const NewFloatingUnitPriceCadenceQuarterly = shared.NewFloatingUnitPriceCadenceQuarterly

// This is an alias to an internal value.
const NewFloatingUnitPriceCadenceOneTime = shared.NewFloatingUnitPriceCadenceOneTime

// This is an alias to an internal value.
const NewFloatingUnitPriceCadenceCustom = shared.NewFloatingUnitPriceCadenceCustom

// The pricing model type
//
// This is an alias to an internal type.
type NewFloatingUnitPriceModelType = shared.NewFloatingUnitPriceModelType

// This is an alias to an internal value.
const NewFloatingUnitPriceModelTypeUnit = shared.NewFloatingUnitPriceModelTypeUnit

// This is an alias to an internal type.
type NewFloatingUnitPriceConversionRateConfigUnionParam = shared.NewFloatingUnitPriceConversionRateConfigUnionParam

// This is an alias to an internal type.
type NewFloatingUnitPriceConversionRateConfigConversionRateType = shared.NewFloatingUnitPriceConversionRateConfigConversionRateType

// This is an alias to an internal value.
const NewFloatingUnitPriceConversionRateConfigConversionRateTypeUnit = shared.NewFloatingUnitPriceConversionRateConfigConversionRateTypeUnit

// This is an alias to an internal value.
const NewFloatingUnitPriceConversionRateConfigConversionRateTypeTiered = shared.NewFloatingUnitPriceConversionRateConfigConversionRateTypeTiered

// This is an alias to an internal type.
type NewFloatingUnitWithPercentPriceParam = shared.NewFloatingUnitWithPercentPriceParam

// The cadence to bill for this price on.
//
// This is an alias to an internal type.
type NewFloatingUnitWithPercentPriceCadence = shared.NewFloatingUnitWithPercentPriceCadence

// This is an alias to an internal value.
const NewFloatingUnitWithPercentPriceCadenceAnnual = shared.NewFloatingUnitWithPercentPriceCadenceAnnual

// This is an alias to an internal value.
const NewFloatingUnitWithPercentPriceCadenceSemiAnnual = shared.NewFloatingUnitWithPercentPriceCadenceSemiAnnual

// This is an alias to an internal value.
const NewFloatingUnitWithPercentPriceCadenceMonthly = shared.NewFloatingUnitWithPercentPriceCadenceMonthly

// This is an alias to an internal value.
const NewFloatingUnitWithPercentPriceCadenceQuarterly = shared.NewFloatingUnitWithPercentPriceCadenceQuarterly

// This is an alias to an internal value.
const NewFloatingUnitWithPercentPriceCadenceOneTime = shared.NewFloatingUnitWithPercentPriceCadenceOneTime

// This is an alias to an internal value.
const NewFloatingUnitWithPercentPriceCadenceCustom = shared.NewFloatingUnitWithPercentPriceCadenceCustom

// The pricing model type
//
// This is an alias to an internal type.
type NewFloatingUnitWithPercentPriceModelType = shared.NewFloatingUnitWithPercentPriceModelType

// This is an alias to an internal value.
const NewFloatingUnitWithPercentPriceModelTypeUnitWithPercent = shared.NewFloatingUnitWithPercentPriceModelTypeUnitWithPercent

// Configuration for unit_with_percent pricing
//
// This is an alias to an internal type.
type NewFloatingUnitWithPercentPriceUnitWithPercentConfigParam = shared.NewFloatingUnitWithPercentPriceUnitWithPercentConfigParam

// This is an alias to an internal type.
type NewFloatingUnitWithPercentPriceConversionRateConfigUnionParam = shared.NewFloatingUnitWithPercentPriceConversionRateConfigUnionParam

// This is an alias to an internal type.
type NewFloatingUnitWithPercentPriceConversionRateConfigConversionRateType = shared.NewFloatingUnitWithPercentPriceConversionRateConfigConversionRateType

// This is an alias to an internal value.
const NewFloatingUnitWithPercentPriceConversionRateConfigConversionRateTypeUnit = shared.NewFloatingUnitWithPercentPriceConversionRateConfigConversionRateTypeUnit

// This is an alias to an internal value.
const NewFloatingUnitWithPercentPriceConversionRateConfigConversionRateTypeTiered = shared.NewFloatingUnitWithPercentPriceConversionRateConfigConversionRateTypeTiered

// This is an alias to an internal type.
type NewFloatingUnitWithProrationPriceParam = shared.NewFloatingUnitWithProrationPriceParam

// The cadence to bill for this price on.
//
// This is an alias to an internal type.
type NewFloatingUnitWithProrationPriceCadence = shared.NewFloatingUnitWithProrationPriceCadence

// This is an alias to an internal value.
const NewFloatingUnitWithProrationPriceCadenceAnnual = shared.NewFloatingUnitWithProrationPriceCadenceAnnual

// This is an alias to an internal value.
const NewFloatingUnitWithProrationPriceCadenceSemiAnnual = shared.NewFloatingUnitWithProrationPriceCadenceSemiAnnual

// This is an alias to an internal value.
const NewFloatingUnitWithProrationPriceCadenceMonthly = shared.NewFloatingUnitWithProrationPriceCadenceMonthly

// This is an alias to an internal value.
const NewFloatingUnitWithProrationPriceCadenceQuarterly = shared.NewFloatingUnitWithProrationPriceCadenceQuarterly

// This is an alias to an internal value.
const NewFloatingUnitWithProrationPriceCadenceOneTime = shared.NewFloatingUnitWithProrationPriceCadenceOneTime

// This is an alias to an internal value.
const NewFloatingUnitWithProrationPriceCadenceCustom = shared.NewFloatingUnitWithProrationPriceCadenceCustom

// The pricing model type
//
// This is an alias to an internal type.
type NewFloatingUnitWithProrationPriceModelType = shared.NewFloatingUnitWithProrationPriceModelType

// This is an alias to an internal value.
const NewFloatingUnitWithProrationPriceModelTypeUnitWithProration = shared.NewFloatingUnitWithProrationPriceModelTypeUnitWithProration

// Configuration for unit_with_proration pricing
//
// This is an alias to an internal type.
type NewFloatingUnitWithProrationPriceUnitWithProrationConfigParam = shared.NewFloatingUnitWithProrationPriceUnitWithProrationConfigParam

// This is an alias to an internal type.
type NewFloatingUnitWithProrationPriceConversionRateConfigUnionParam = shared.NewFloatingUnitWithProrationPriceConversionRateConfigUnionParam

// This is an alias to an internal type.
type NewFloatingUnitWithProrationPriceConversionRateConfigConversionRateType = shared.NewFloatingUnitWithProrationPriceConversionRateConfigConversionRateType

// This is an alias to an internal value.
const NewFloatingUnitWithProrationPriceConversionRateConfigConversionRateTypeUnit = shared.NewFloatingUnitWithProrationPriceConversionRateConfigConversionRateTypeUnit

// This is an alias to an internal value.
const NewFloatingUnitWithProrationPriceConversionRateConfigConversionRateTypeTiered = shared.NewFloatingUnitWithProrationPriceConversionRateConfigConversionRateTypeTiered

// This is an alias to an internal type.
type NewMaximumParam = shared.NewMaximumParam

// This is an alias to an internal type.
type NewMaximumAdjustmentType = shared.NewMaximumAdjustmentType

// This is an alias to an internal value.
const NewMaximumAdjustmentTypeMaximum = shared.NewMaximumAdjustmentTypeMaximum

// If set, the adjustment will apply to every price on the subscription.
//
// This is an alias to an internal type.
type NewMaximumAppliesToAll = shared.NewMaximumAppliesToAll

// This is an alias to an internal value.
const NewMaximumAppliesToAllTrue = shared.NewMaximumAppliesToAllTrue

// This is an alias to an internal type.
type NewMaximumFilterParam = shared.NewMaximumFilterParam

// The property of the price to filter on.
//
// This is an alias to an internal type.
type NewMaximumFiltersField = shared.NewMaximumFiltersField

// This is an alias to an internal value.
const NewMaximumFiltersFieldPriceID = shared.NewMaximumFiltersFieldPriceID

// This is an alias to an internal value.
const NewMaximumFiltersFieldItemID = shared.NewMaximumFiltersFieldItemID

// This is an alias to an internal value.
const NewMaximumFiltersFieldPriceType = shared.NewMaximumFiltersFieldPriceType

// This is an alias to an internal value.
const NewMaximumFiltersFieldCurrency = shared.NewMaximumFiltersFieldCurrency

// This is an alias to an internal value.
const NewMaximumFiltersFieldPricingUnitID = shared.NewMaximumFiltersFieldPricingUnitID

// Should prices that match the filter be included or excluded.
//
// This is an alias to an internal type.
type NewMaximumFiltersOperator = shared.NewMaximumFiltersOperator

// This is an alias to an internal value.
const NewMaximumFiltersOperatorIncludes = shared.NewMaximumFiltersOperatorIncludes

// This is an alias to an internal value.
const NewMaximumFiltersOperatorExcludes = shared.NewMaximumFiltersOperatorExcludes

// If set, only prices of the specified type will have the adjustment applied.
//
// This is an alias to an internal type.
type NewMaximumPriceType = shared.NewMaximumPriceType

// This is an alias to an internal value.
const NewMaximumPriceTypeUsage = shared.NewMaximumPriceTypeUsage

// This is an alias to an internal value.
const NewMaximumPriceTypeFixedInAdvance = shared.NewMaximumPriceTypeFixedInAdvance

// This is an alias to an internal value.
const NewMaximumPriceTypeFixedInArrears = shared.NewMaximumPriceTypeFixedInArrears

// This is an alias to an internal value.
const NewMaximumPriceTypeFixed = shared.NewMaximumPriceTypeFixed

// This is an alias to an internal value.
const NewMaximumPriceTypeInArrears = shared.NewMaximumPriceTypeInArrears

// This is an alias to an internal type.
type NewMinimumParam = shared.NewMinimumParam

// This is an alias to an internal type.
type NewMinimumAdjustmentType = shared.NewMinimumAdjustmentType

// This is an alias to an internal value.
const NewMinimumAdjustmentTypeMinimum = shared.NewMinimumAdjustmentTypeMinimum

// If set, the adjustment will apply to every price on the subscription.
//
// This is an alias to an internal type.
type NewMinimumAppliesToAll = shared.NewMinimumAppliesToAll

// This is an alias to an internal value.
const NewMinimumAppliesToAllTrue = shared.NewMinimumAppliesToAllTrue

// This is an alias to an internal type.
type NewMinimumFilterParam = shared.NewMinimumFilterParam

// The property of the price to filter on.
//
// This is an alias to an internal type.
type NewMinimumFiltersField = shared.NewMinimumFiltersField

// This is an alias to an internal value.
const NewMinimumFiltersFieldPriceID = shared.NewMinimumFiltersFieldPriceID

// This is an alias to an internal value.
const NewMinimumFiltersFieldItemID = shared.NewMinimumFiltersFieldItemID

// This is an alias to an internal value.
const NewMinimumFiltersFieldPriceType = shared.NewMinimumFiltersFieldPriceType

// This is an alias to an internal value.
const NewMinimumFiltersFieldCurrency = shared.NewMinimumFiltersFieldCurrency

// This is an alias to an internal value.
const NewMinimumFiltersFieldPricingUnitID = shared.NewMinimumFiltersFieldPricingUnitID

// Should prices that match the filter be included or excluded.
//
// This is an alias to an internal type.
type NewMinimumFiltersOperator = shared.NewMinimumFiltersOperator

// This is an alias to an internal value.
const NewMinimumFiltersOperatorIncludes = shared.NewMinimumFiltersOperatorIncludes

// This is an alias to an internal value.
const NewMinimumFiltersOperatorExcludes = shared.NewMinimumFiltersOperatorExcludes

// If set, only prices of the specified type will have the adjustment applied.
//
// This is an alias to an internal type.
type NewMinimumPriceType = shared.NewMinimumPriceType

// This is an alias to an internal value.
const NewMinimumPriceTypeUsage = shared.NewMinimumPriceTypeUsage

// This is an alias to an internal value.
const NewMinimumPriceTypeFixedInAdvance = shared.NewMinimumPriceTypeFixedInAdvance

// This is an alias to an internal value.
const NewMinimumPriceTypeFixedInArrears = shared.NewMinimumPriceTypeFixedInArrears

// This is an alias to an internal value.
const NewMinimumPriceTypeFixed = shared.NewMinimumPriceTypeFixed

// This is an alias to an internal value.
const NewMinimumPriceTypeInArrears = shared.NewMinimumPriceTypeInArrears

// This is an alias to an internal type.
type NewPercentageDiscountParam = shared.NewPercentageDiscountParam

// This is an alias to an internal type.
type NewPercentageDiscountAdjustmentType = shared.NewPercentageDiscountAdjustmentType

// This is an alias to an internal value.
const NewPercentageDiscountAdjustmentTypePercentageDiscount = shared.NewPercentageDiscountAdjustmentTypePercentageDiscount

// If set, the adjustment will apply to every price on the subscription.
//
// This is an alias to an internal type.
type NewPercentageDiscountAppliesToAll = shared.NewPercentageDiscountAppliesToAll

// This is an alias to an internal value.
const NewPercentageDiscountAppliesToAllTrue = shared.NewPercentageDiscountAppliesToAllTrue

// This is an alias to an internal type.
type NewPercentageDiscountFilterParam = shared.NewPercentageDiscountFilterParam

// The property of the price to filter on.
//
// This is an alias to an internal type.
type NewPercentageDiscountFiltersField = shared.NewPercentageDiscountFiltersField

// This is an alias to an internal value.
const NewPercentageDiscountFiltersFieldPriceID = shared.NewPercentageDiscountFiltersFieldPriceID

// This is an alias to an internal value.
const NewPercentageDiscountFiltersFieldItemID = shared.NewPercentageDiscountFiltersFieldItemID

// This is an alias to an internal value.
const NewPercentageDiscountFiltersFieldPriceType = shared.NewPercentageDiscountFiltersFieldPriceType

// This is an alias to an internal value.
const NewPercentageDiscountFiltersFieldCurrency = shared.NewPercentageDiscountFiltersFieldCurrency

// This is an alias to an internal value.
const NewPercentageDiscountFiltersFieldPricingUnitID = shared.NewPercentageDiscountFiltersFieldPricingUnitID

// Should prices that match the filter be included or excluded.
//
// This is an alias to an internal type.
type NewPercentageDiscountFiltersOperator = shared.NewPercentageDiscountFiltersOperator

// This is an alias to an internal value.
const NewPercentageDiscountFiltersOperatorIncludes = shared.NewPercentageDiscountFiltersOperatorIncludes

// This is an alias to an internal value.
const NewPercentageDiscountFiltersOperatorExcludes = shared.NewPercentageDiscountFiltersOperatorExcludes

// If set, only prices of the specified type will have the adjustment applied.
//
// This is an alias to an internal type.
type NewPercentageDiscountPriceType = shared.NewPercentageDiscountPriceType

// This is an alias to an internal value.
const NewPercentageDiscountPriceTypeUsage = shared.NewPercentageDiscountPriceTypeUsage

// This is an alias to an internal value.
const NewPercentageDiscountPriceTypeFixedInAdvance = shared.NewPercentageDiscountPriceTypeFixedInAdvance

// This is an alias to an internal value.
const NewPercentageDiscountPriceTypeFixedInArrears = shared.NewPercentageDiscountPriceTypeFixedInArrears

// This is an alias to an internal value.
const NewPercentageDiscountPriceTypeFixed = shared.NewPercentageDiscountPriceTypeFixed

// This is an alias to an internal value.
const NewPercentageDiscountPriceTypeInArrears = shared.NewPercentageDiscountPriceTypeInArrears

// This is an alias to an internal type.
type NewPlanBulkPriceParam = shared.NewPlanBulkPriceParam

// The cadence to bill for this price on.
//
// This is an alias to an internal type.
type NewPlanBulkPriceCadence = shared.NewPlanBulkPriceCadence

// This is an alias to an internal value.
const NewPlanBulkPriceCadenceAnnual = shared.NewPlanBulkPriceCadenceAnnual

// This is an alias to an internal value.
const NewPlanBulkPriceCadenceSemiAnnual = shared.NewPlanBulkPriceCadenceSemiAnnual

// This is an alias to an internal value.
const NewPlanBulkPriceCadenceMonthly = shared.NewPlanBulkPriceCadenceMonthly

// This is an alias to an internal value.
const NewPlanBulkPriceCadenceQuarterly = shared.NewPlanBulkPriceCadenceQuarterly

// This is an alias to an internal value.
const NewPlanBulkPriceCadenceOneTime = shared.NewPlanBulkPriceCadenceOneTime

// This is an alias to an internal value.
const NewPlanBulkPriceCadenceCustom = shared.NewPlanBulkPriceCadenceCustom

// The pricing model type
//
// This is an alias to an internal type.
type NewPlanBulkPriceModelType = shared.NewPlanBulkPriceModelType

// This is an alias to an internal value.
const NewPlanBulkPriceModelTypeBulk = shared.NewPlanBulkPriceModelTypeBulk

// This is an alias to an internal type.
type NewPlanBulkPriceConversionRateConfigUnionParam = shared.NewPlanBulkPriceConversionRateConfigUnionParam

// This is an alias to an internal type.
type NewPlanBulkPriceConversionRateConfigConversionRateType = shared.NewPlanBulkPriceConversionRateConfigConversionRateType

// This is an alias to an internal value.
const NewPlanBulkPriceConversionRateConfigConversionRateTypeUnit = shared.NewPlanBulkPriceConversionRateConfigConversionRateTypeUnit

// This is an alias to an internal value.
const NewPlanBulkPriceConversionRateConfigConversionRateTypeTiered = shared.NewPlanBulkPriceConversionRateConfigConversionRateTypeTiered

// This is an alias to an internal type.
type NewPlanBulkWithProrationPriceParam = shared.NewPlanBulkWithProrationPriceParam

// Configuration for bulk_with_proration pricing
//
// This is an alias to an internal type.
type NewPlanBulkWithProrationPriceBulkWithProrationConfigParam = shared.NewPlanBulkWithProrationPriceBulkWithProrationConfigParam

// Configuration for a single bulk pricing tier with proration
//
// This is an alias to an internal type.
type NewPlanBulkWithProrationPriceBulkWithProrationConfigTierParam = shared.NewPlanBulkWithProrationPriceBulkWithProrationConfigTierParam

// The cadence to bill for this price on.
//
// This is an alias to an internal type.
type NewPlanBulkWithProrationPriceCadence = shared.NewPlanBulkWithProrationPriceCadence

// This is an alias to an internal value.
const NewPlanBulkWithProrationPriceCadenceAnnual = shared.NewPlanBulkWithProrationPriceCadenceAnnual

// This is an alias to an internal value.
const NewPlanBulkWithProrationPriceCadenceSemiAnnual = shared.NewPlanBulkWithProrationPriceCadenceSemiAnnual

// This is an alias to an internal value.
const NewPlanBulkWithProrationPriceCadenceMonthly = shared.NewPlanBulkWithProrationPriceCadenceMonthly

// This is an alias to an internal value.
const NewPlanBulkWithProrationPriceCadenceQuarterly = shared.NewPlanBulkWithProrationPriceCadenceQuarterly

// This is an alias to an internal value.
const NewPlanBulkWithProrationPriceCadenceOneTime = shared.NewPlanBulkWithProrationPriceCadenceOneTime

// This is an alias to an internal value.
const NewPlanBulkWithProrationPriceCadenceCustom = shared.NewPlanBulkWithProrationPriceCadenceCustom

// The pricing model type
//
// This is an alias to an internal type.
type NewPlanBulkWithProrationPriceModelType = shared.NewPlanBulkWithProrationPriceModelType

// This is an alias to an internal value.
const NewPlanBulkWithProrationPriceModelTypeBulkWithProration = shared.NewPlanBulkWithProrationPriceModelTypeBulkWithProration

// This is an alias to an internal type.
type NewPlanBulkWithProrationPriceConversionRateConfigUnionParam = shared.NewPlanBulkWithProrationPriceConversionRateConfigUnionParam

// This is an alias to an internal type.
type NewPlanBulkWithProrationPriceConversionRateConfigConversionRateType = shared.NewPlanBulkWithProrationPriceConversionRateConfigConversionRateType

// This is an alias to an internal value.
const NewPlanBulkWithProrationPriceConversionRateConfigConversionRateTypeUnit = shared.NewPlanBulkWithProrationPriceConversionRateConfigConversionRateTypeUnit

// This is an alias to an internal value.
const NewPlanBulkWithProrationPriceConversionRateConfigConversionRateTypeTiered = shared.NewPlanBulkWithProrationPriceConversionRateConfigConversionRateTypeTiered

// This is an alias to an internal type.
type NewPlanCumulativeGroupedBulkPriceParam = shared.NewPlanCumulativeGroupedBulkPriceParam

// The cadence to bill for this price on.
//
// This is an alias to an internal type.
type NewPlanCumulativeGroupedBulkPriceCadence = shared.NewPlanCumulativeGroupedBulkPriceCadence

// This is an alias to an internal value.
const NewPlanCumulativeGroupedBulkPriceCadenceAnnual = shared.NewPlanCumulativeGroupedBulkPriceCadenceAnnual

// This is an alias to an internal value.
const NewPlanCumulativeGroupedBulkPriceCadenceSemiAnnual = shared.NewPlanCumulativeGroupedBulkPriceCadenceSemiAnnual

// This is an alias to an internal value.
const NewPlanCumulativeGroupedBulkPriceCadenceMonthly = shared.NewPlanCumulativeGroupedBulkPriceCadenceMonthly

// This is an alias to an internal value.
const NewPlanCumulativeGroupedBulkPriceCadenceQuarterly = shared.NewPlanCumulativeGroupedBulkPriceCadenceQuarterly

// This is an alias to an internal value.
const NewPlanCumulativeGroupedBulkPriceCadenceOneTime = shared.NewPlanCumulativeGroupedBulkPriceCadenceOneTime

// This is an alias to an internal value.
const NewPlanCumulativeGroupedBulkPriceCadenceCustom = shared.NewPlanCumulativeGroupedBulkPriceCadenceCustom

// Configuration for cumulative_grouped_bulk pricing
//
// This is an alias to an internal type.
type NewPlanCumulativeGroupedBulkPriceCumulativeGroupedBulkConfigParam = shared.NewPlanCumulativeGroupedBulkPriceCumulativeGroupedBulkConfigParam

// Configuration for a dimension value entry
//
// This is an alias to an internal type.
type NewPlanCumulativeGroupedBulkPriceCumulativeGroupedBulkConfigDimensionValueParam = shared.NewPlanCumulativeGroupedBulkPriceCumulativeGroupedBulkConfigDimensionValueParam

// The pricing model type
//
// This is an alias to an internal type.
type NewPlanCumulativeGroupedBulkPriceModelType = shared.NewPlanCumulativeGroupedBulkPriceModelType

// This is an alias to an internal value.
const NewPlanCumulativeGroupedBulkPriceModelTypeCumulativeGroupedBulk = shared.NewPlanCumulativeGroupedBulkPriceModelTypeCumulativeGroupedBulk

// This is an alias to an internal type.
type NewPlanCumulativeGroupedBulkPriceConversionRateConfigUnionParam = shared.NewPlanCumulativeGroupedBulkPriceConversionRateConfigUnionParam

// This is an alias to an internal type.
type NewPlanCumulativeGroupedBulkPriceConversionRateConfigConversionRateType = shared.NewPlanCumulativeGroupedBulkPriceConversionRateConfigConversionRateType

// This is an alias to an internal value.
const NewPlanCumulativeGroupedBulkPriceConversionRateConfigConversionRateTypeUnit = shared.NewPlanCumulativeGroupedBulkPriceConversionRateConfigConversionRateTypeUnit

// This is an alias to an internal value.
const NewPlanCumulativeGroupedBulkPriceConversionRateConfigConversionRateTypeTiered = shared.NewPlanCumulativeGroupedBulkPriceConversionRateConfigConversionRateTypeTiered

// This is an alias to an internal type.
type NewPlanGroupedAllocationPriceParam = shared.NewPlanGroupedAllocationPriceParam

// The cadence to bill for this price on.
//
// This is an alias to an internal type.
type NewPlanGroupedAllocationPriceCadence = shared.NewPlanGroupedAllocationPriceCadence

// This is an alias to an internal value.
const NewPlanGroupedAllocationPriceCadenceAnnual = shared.NewPlanGroupedAllocationPriceCadenceAnnual

// This is an alias to an internal value.
const NewPlanGroupedAllocationPriceCadenceSemiAnnual = shared.NewPlanGroupedAllocationPriceCadenceSemiAnnual

// This is an alias to an internal value.
const NewPlanGroupedAllocationPriceCadenceMonthly = shared.NewPlanGroupedAllocationPriceCadenceMonthly

// This is an alias to an internal value.
const NewPlanGroupedAllocationPriceCadenceQuarterly = shared.NewPlanGroupedAllocationPriceCadenceQuarterly

// This is an alias to an internal value.
const NewPlanGroupedAllocationPriceCadenceOneTime = shared.NewPlanGroupedAllocationPriceCadenceOneTime

// This is an alias to an internal value.
const NewPlanGroupedAllocationPriceCadenceCustom = shared.NewPlanGroupedAllocationPriceCadenceCustom

// Configuration for grouped_allocation pricing
//
// This is an alias to an internal type.
type NewPlanGroupedAllocationPriceGroupedAllocationConfigParam = shared.NewPlanGroupedAllocationPriceGroupedAllocationConfigParam

// The pricing model type
//
// This is an alias to an internal type.
type NewPlanGroupedAllocationPriceModelType = shared.NewPlanGroupedAllocationPriceModelType

// This is an alias to an internal value.
const NewPlanGroupedAllocationPriceModelTypeGroupedAllocation = shared.NewPlanGroupedAllocationPriceModelTypeGroupedAllocation

// This is an alias to an internal type.
type NewPlanGroupedAllocationPriceConversionRateConfigUnionParam = shared.NewPlanGroupedAllocationPriceConversionRateConfigUnionParam

// This is an alias to an internal type.
type NewPlanGroupedAllocationPriceConversionRateConfigConversionRateType = shared.NewPlanGroupedAllocationPriceConversionRateConfigConversionRateType

// This is an alias to an internal value.
const NewPlanGroupedAllocationPriceConversionRateConfigConversionRateTypeUnit = shared.NewPlanGroupedAllocationPriceConversionRateConfigConversionRateTypeUnit

// This is an alias to an internal value.
const NewPlanGroupedAllocationPriceConversionRateConfigConversionRateTypeTiered = shared.NewPlanGroupedAllocationPriceConversionRateConfigConversionRateTypeTiered

// This is an alias to an internal type.
type NewPlanGroupedTieredPackagePriceParam = shared.NewPlanGroupedTieredPackagePriceParam

// The cadence to bill for this price on.
//
// This is an alias to an internal type.
type NewPlanGroupedTieredPackagePriceCadence = shared.NewPlanGroupedTieredPackagePriceCadence

// This is an alias to an internal value.
const NewPlanGroupedTieredPackagePriceCadenceAnnual = shared.NewPlanGroupedTieredPackagePriceCadenceAnnual

// This is an alias to an internal value.
const NewPlanGroupedTieredPackagePriceCadenceSemiAnnual = shared.NewPlanGroupedTieredPackagePriceCadenceSemiAnnual

// This is an alias to an internal value.
const NewPlanGroupedTieredPackagePriceCadenceMonthly = shared.NewPlanGroupedTieredPackagePriceCadenceMonthly

// This is an alias to an internal value.
const NewPlanGroupedTieredPackagePriceCadenceQuarterly = shared.NewPlanGroupedTieredPackagePriceCadenceQuarterly

// This is an alias to an internal value.
const NewPlanGroupedTieredPackagePriceCadenceOneTime = shared.NewPlanGroupedTieredPackagePriceCadenceOneTime

// This is an alias to an internal value.
const NewPlanGroupedTieredPackagePriceCadenceCustom = shared.NewPlanGroupedTieredPackagePriceCadenceCustom

// Configuration for grouped_tiered_package pricing
//
// This is an alias to an internal type.
type NewPlanGroupedTieredPackagePriceGroupedTieredPackageConfigParam = shared.NewPlanGroupedTieredPackagePriceGroupedTieredPackageConfigParam

// Configuration for a single tier
//
// This is an alias to an internal type.
type NewPlanGroupedTieredPackagePriceGroupedTieredPackageConfigTierParam = shared.NewPlanGroupedTieredPackagePriceGroupedTieredPackageConfigTierParam

// The pricing model type
//
// This is an alias to an internal type.
type NewPlanGroupedTieredPackagePriceModelType = shared.NewPlanGroupedTieredPackagePriceModelType

// This is an alias to an internal value.
const NewPlanGroupedTieredPackagePriceModelTypeGroupedTieredPackage = shared.NewPlanGroupedTieredPackagePriceModelTypeGroupedTieredPackage

// This is an alias to an internal type.
type NewPlanGroupedTieredPackagePriceConversionRateConfigUnionParam = shared.NewPlanGroupedTieredPackagePriceConversionRateConfigUnionParam

// This is an alias to an internal type.
type NewPlanGroupedTieredPackagePriceConversionRateConfigConversionRateType = shared.NewPlanGroupedTieredPackagePriceConversionRateConfigConversionRateType

// This is an alias to an internal value.
const NewPlanGroupedTieredPackagePriceConversionRateConfigConversionRateTypeUnit = shared.NewPlanGroupedTieredPackagePriceConversionRateConfigConversionRateTypeUnit

// This is an alias to an internal value.
const NewPlanGroupedTieredPackagePriceConversionRateConfigConversionRateTypeTiered = shared.NewPlanGroupedTieredPackagePriceConversionRateConfigConversionRateTypeTiered

// This is an alias to an internal type.
type NewPlanGroupedTieredPriceParam = shared.NewPlanGroupedTieredPriceParam

// The cadence to bill for this price on.
//
// This is an alias to an internal type.
type NewPlanGroupedTieredPriceCadence = shared.NewPlanGroupedTieredPriceCadence

// This is an alias to an internal value.
const NewPlanGroupedTieredPriceCadenceAnnual = shared.NewPlanGroupedTieredPriceCadenceAnnual

// This is an alias to an internal value.
const NewPlanGroupedTieredPriceCadenceSemiAnnual = shared.NewPlanGroupedTieredPriceCadenceSemiAnnual

// This is an alias to an internal value.
const NewPlanGroupedTieredPriceCadenceMonthly = shared.NewPlanGroupedTieredPriceCadenceMonthly

// This is an alias to an internal value.
const NewPlanGroupedTieredPriceCadenceQuarterly = shared.NewPlanGroupedTieredPriceCadenceQuarterly

// This is an alias to an internal value.
const NewPlanGroupedTieredPriceCadenceOneTime = shared.NewPlanGroupedTieredPriceCadenceOneTime

// This is an alias to an internal value.
const NewPlanGroupedTieredPriceCadenceCustom = shared.NewPlanGroupedTieredPriceCadenceCustom

// Configuration for grouped_tiered pricing
//
// This is an alias to an internal type.
type NewPlanGroupedTieredPriceGroupedTieredConfigParam = shared.NewPlanGroupedTieredPriceGroupedTieredConfigParam

// Configuration for a single tier
//
// This is an alias to an internal type.
type NewPlanGroupedTieredPriceGroupedTieredConfigTierParam = shared.NewPlanGroupedTieredPriceGroupedTieredConfigTierParam

// The pricing model type
//
// This is an alias to an internal type.
type NewPlanGroupedTieredPriceModelType = shared.NewPlanGroupedTieredPriceModelType

// This is an alias to an internal value.
const NewPlanGroupedTieredPriceModelTypeGroupedTiered = shared.NewPlanGroupedTieredPriceModelTypeGroupedTiered

// This is an alias to an internal type.
type NewPlanGroupedTieredPriceConversionRateConfigUnionParam = shared.NewPlanGroupedTieredPriceConversionRateConfigUnionParam

// This is an alias to an internal type.
type NewPlanGroupedTieredPriceConversionRateConfigConversionRateType = shared.NewPlanGroupedTieredPriceConversionRateConfigConversionRateType

// This is an alias to an internal value.
const NewPlanGroupedTieredPriceConversionRateConfigConversionRateTypeUnit = shared.NewPlanGroupedTieredPriceConversionRateConfigConversionRateTypeUnit

// This is an alias to an internal value.
const NewPlanGroupedTieredPriceConversionRateConfigConversionRateTypeTiered = shared.NewPlanGroupedTieredPriceConversionRateConfigConversionRateTypeTiered

// This is an alias to an internal type.
type NewPlanGroupedWithMeteredMinimumPriceParam = shared.NewPlanGroupedWithMeteredMinimumPriceParam

// The cadence to bill for this price on.
//
// This is an alias to an internal type.
type NewPlanGroupedWithMeteredMinimumPriceCadence = shared.NewPlanGroupedWithMeteredMinimumPriceCadence

// This is an alias to an internal value.
const NewPlanGroupedWithMeteredMinimumPriceCadenceAnnual = shared.NewPlanGroupedWithMeteredMinimumPriceCadenceAnnual

// This is an alias to an internal value.
const NewPlanGroupedWithMeteredMinimumPriceCadenceSemiAnnual = shared.NewPlanGroupedWithMeteredMinimumPriceCadenceSemiAnnual

// This is an alias to an internal value.
const NewPlanGroupedWithMeteredMinimumPriceCadenceMonthly = shared.NewPlanGroupedWithMeteredMinimumPriceCadenceMonthly

// This is an alias to an internal value.
const NewPlanGroupedWithMeteredMinimumPriceCadenceQuarterly = shared.NewPlanGroupedWithMeteredMinimumPriceCadenceQuarterly

// This is an alias to an internal value.
const NewPlanGroupedWithMeteredMinimumPriceCadenceOneTime = shared.NewPlanGroupedWithMeteredMinimumPriceCadenceOneTime

// This is an alias to an internal value.
const NewPlanGroupedWithMeteredMinimumPriceCadenceCustom = shared.NewPlanGroupedWithMeteredMinimumPriceCadenceCustom

// Configuration for grouped_with_metered_minimum pricing
//
// This is an alias to an internal type.
type NewPlanGroupedWithMeteredMinimumPriceGroupedWithMeteredMinimumConfigParam = shared.NewPlanGroupedWithMeteredMinimumPriceGroupedWithMeteredMinimumConfigParam

// Configuration for a scaling factor
//
// This is an alias to an internal type.
type NewPlanGroupedWithMeteredMinimumPriceGroupedWithMeteredMinimumConfigScalingFactorParam = shared.NewPlanGroupedWithMeteredMinimumPriceGroupedWithMeteredMinimumConfigScalingFactorParam

// Configuration for a unit amount
//
// This is an alias to an internal type.
type NewPlanGroupedWithMeteredMinimumPriceGroupedWithMeteredMinimumConfigUnitAmountParam = shared.NewPlanGroupedWithMeteredMinimumPriceGroupedWithMeteredMinimumConfigUnitAmountParam

// The pricing model type
//
// This is an alias to an internal type.
type NewPlanGroupedWithMeteredMinimumPriceModelType = shared.NewPlanGroupedWithMeteredMinimumPriceModelType

// This is an alias to an internal value.
const NewPlanGroupedWithMeteredMinimumPriceModelTypeGroupedWithMeteredMinimum = shared.NewPlanGroupedWithMeteredMinimumPriceModelTypeGroupedWithMeteredMinimum

// This is an alias to an internal type.
type NewPlanGroupedWithMeteredMinimumPriceConversionRateConfigUnionParam = shared.NewPlanGroupedWithMeteredMinimumPriceConversionRateConfigUnionParam

// This is an alias to an internal type.
type NewPlanGroupedWithMeteredMinimumPriceConversionRateConfigConversionRateType = shared.NewPlanGroupedWithMeteredMinimumPriceConversionRateConfigConversionRateType

// This is an alias to an internal value.
const NewPlanGroupedWithMeteredMinimumPriceConversionRateConfigConversionRateTypeUnit = shared.NewPlanGroupedWithMeteredMinimumPriceConversionRateConfigConversionRateTypeUnit

// This is an alias to an internal value.
const NewPlanGroupedWithMeteredMinimumPriceConversionRateConfigConversionRateTypeTiered = shared.NewPlanGroupedWithMeteredMinimumPriceConversionRateConfigConversionRateTypeTiered

// This is an alias to an internal type.
type NewPlanGroupedWithProratedMinimumPriceParam = shared.NewPlanGroupedWithProratedMinimumPriceParam

// The cadence to bill for this price on.
//
// This is an alias to an internal type.
type NewPlanGroupedWithProratedMinimumPriceCadence = shared.NewPlanGroupedWithProratedMinimumPriceCadence

// This is an alias to an internal value.
const NewPlanGroupedWithProratedMinimumPriceCadenceAnnual = shared.NewPlanGroupedWithProratedMinimumPriceCadenceAnnual

// This is an alias to an internal value.
const NewPlanGroupedWithProratedMinimumPriceCadenceSemiAnnual = shared.NewPlanGroupedWithProratedMinimumPriceCadenceSemiAnnual

// This is an alias to an internal value.
const NewPlanGroupedWithProratedMinimumPriceCadenceMonthly = shared.NewPlanGroupedWithProratedMinimumPriceCadenceMonthly

// This is an alias to an internal value.
const NewPlanGroupedWithProratedMinimumPriceCadenceQuarterly = shared.NewPlanGroupedWithProratedMinimumPriceCadenceQuarterly

// This is an alias to an internal value.
const NewPlanGroupedWithProratedMinimumPriceCadenceOneTime = shared.NewPlanGroupedWithProratedMinimumPriceCadenceOneTime

// This is an alias to an internal value.
const NewPlanGroupedWithProratedMinimumPriceCadenceCustom = shared.NewPlanGroupedWithProratedMinimumPriceCadenceCustom

// Configuration for grouped_with_prorated_minimum pricing
//
// This is an alias to an internal type.
type NewPlanGroupedWithProratedMinimumPriceGroupedWithProratedMinimumConfigParam = shared.NewPlanGroupedWithProratedMinimumPriceGroupedWithProratedMinimumConfigParam

// The pricing model type
//
// This is an alias to an internal type.
type NewPlanGroupedWithProratedMinimumPriceModelType = shared.NewPlanGroupedWithProratedMinimumPriceModelType

// This is an alias to an internal value.
const NewPlanGroupedWithProratedMinimumPriceModelTypeGroupedWithProratedMinimum = shared.NewPlanGroupedWithProratedMinimumPriceModelTypeGroupedWithProratedMinimum

// This is an alias to an internal type.
type NewPlanGroupedWithProratedMinimumPriceConversionRateConfigUnionParam = shared.NewPlanGroupedWithProratedMinimumPriceConversionRateConfigUnionParam

// This is an alias to an internal type.
type NewPlanGroupedWithProratedMinimumPriceConversionRateConfigConversionRateType = shared.NewPlanGroupedWithProratedMinimumPriceConversionRateConfigConversionRateType

// This is an alias to an internal value.
const NewPlanGroupedWithProratedMinimumPriceConversionRateConfigConversionRateTypeUnit = shared.NewPlanGroupedWithProratedMinimumPriceConversionRateConfigConversionRateTypeUnit

// This is an alias to an internal value.
const NewPlanGroupedWithProratedMinimumPriceConversionRateConfigConversionRateTypeTiered = shared.NewPlanGroupedWithProratedMinimumPriceConversionRateConfigConversionRateTypeTiered

// This is an alias to an internal type.
type NewPlanMatrixPriceParam = shared.NewPlanMatrixPriceParam

// The cadence to bill for this price on.
//
// This is an alias to an internal type.
type NewPlanMatrixPriceCadence = shared.NewPlanMatrixPriceCadence

// This is an alias to an internal value.
const NewPlanMatrixPriceCadenceAnnual = shared.NewPlanMatrixPriceCadenceAnnual

// This is an alias to an internal value.
const NewPlanMatrixPriceCadenceSemiAnnual = shared.NewPlanMatrixPriceCadenceSemiAnnual

// This is an alias to an internal value.
const NewPlanMatrixPriceCadenceMonthly = shared.NewPlanMatrixPriceCadenceMonthly

// This is an alias to an internal value.
const NewPlanMatrixPriceCadenceQuarterly = shared.NewPlanMatrixPriceCadenceQuarterly

// This is an alias to an internal value.
const NewPlanMatrixPriceCadenceOneTime = shared.NewPlanMatrixPriceCadenceOneTime

// This is an alias to an internal value.
const NewPlanMatrixPriceCadenceCustom = shared.NewPlanMatrixPriceCadenceCustom

// The pricing model type
//
// This is an alias to an internal type.
type NewPlanMatrixPriceModelType = shared.NewPlanMatrixPriceModelType

// This is an alias to an internal value.
const NewPlanMatrixPriceModelTypeMatrix = shared.NewPlanMatrixPriceModelTypeMatrix

// This is an alias to an internal type.
type NewPlanMatrixPriceConversionRateConfigUnionParam = shared.NewPlanMatrixPriceConversionRateConfigUnionParam

// This is an alias to an internal type.
type NewPlanMatrixPriceConversionRateConfigConversionRateType = shared.NewPlanMatrixPriceConversionRateConfigConversionRateType

// This is an alias to an internal value.
const NewPlanMatrixPriceConversionRateConfigConversionRateTypeUnit = shared.NewPlanMatrixPriceConversionRateConfigConversionRateTypeUnit

// This is an alias to an internal value.
const NewPlanMatrixPriceConversionRateConfigConversionRateTypeTiered = shared.NewPlanMatrixPriceConversionRateConfigConversionRateTypeTiered

// This is an alias to an internal type.
type NewPlanMatrixWithAllocationPriceParam = shared.NewPlanMatrixWithAllocationPriceParam

// The cadence to bill for this price on.
//
// This is an alias to an internal type.
type NewPlanMatrixWithAllocationPriceCadence = shared.NewPlanMatrixWithAllocationPriceCadence

// This is an alias to an internal value.
const NewPlanMatrixWithAllocationPriceCadenceAnnual = shared.NewPlanMatrixWithAllocationPriceCadenceAnnual

// This is an alias to an internal value.
const NewPlanMatrixWithAllocationPriceCadenceSemiAnnual = shared.NewPlanMatrixWithAllocationPriceCadenceSemiAnnual

// This is an alias to an internal value.
const NewPlanMatrixWithAllocationPriceCadenceMonthly = shared.NewPlanMatrixWithAllocationPriceCadenceMonthly

// This is an alias to an internal value.
const NewPlanMatrixWithAllocationPriceCadenceQuarterly = shared.NewPlanMatrixWithAllocationPriceCadenceQuarterly

// This is an alias to an internal value.
const NewPlanMatrixWithAllocationPriceCadenceOneTime = shared.NewPlanMatrixWithAllocationPriceCadenceOneTime

// This is an alias to an internal value.
const NewPlanMatrixWithAllocationPriceCadenceCustom = shared.NewPlanMatrixWithAllocationPriceCadenceCustom

// The pricing model type
//
// This is an alias to an internal type.
type NewPlanMatrixWithAllocationPriceModelType = shared.NewPlanMatrixWithAllocationPriceModelType

// This is an alias to an internal value.
const NewPlanMatrixWithAllocationPriceModelTypeMatrixWithAllocation = shared.NewPlanMatrixWithAllocationPriceModelTypeMatrixWithAllocation

// This is an alias to an internal type.
type NewPlanMatrixWithAllocationPriceConversionRateConfigUnionParam = shared.NewPlanMatrixWithAllocationPriceConversionRateConfigUnionParam

// This is an alias to an internal type.
type NewPlanMatrixWithAllocationPriceConversionRateConfigConversionRateType = shared.NewPlanMatrixWithAllocationPriceConversionRateConfigConversionRateType

// This is an alias to an internal value.
const NewPlanMatrixWithAllocationPriceConversionRateConfigConversionRateTypeUnit = shared.NewPlanMatrixWithAllocationPriceConversionRateConfigConversionRateTypeUnit

// This is an alias to an internal value.
const NewPlanMatrixWithAllocationPriceConversionRateConfigConversionRateTypeTiered = shared.NewPlanMatrixWithAllocationPriceConversionRateConfigConversionRateTypeTiered

// This is an alias to an internal type.
type NewPlanMatrixWithDisplayNamePriceParam = shared.NewPlanMatrixWithDisplayNamePriceParam

// The cadence to bill for this price on.
//
// This is an alias to an internal type.
type NewPlanMatrixWithDisplayNamePriceCadence = shared.NewPlanMatrixWithDisplayNamePriceCadence

// This is an alias to an internal value.
const NewPlanMatrixWithDisplayNamePriceCadenceAnnual = shared.NewPlanMatrixWithDisplayNamePriceCadenceAnnual

// This is an alias to an internal value.
const NewPlanMatrixWithDisplayNamePriceCadenceSemiAnnual = shared.NewPlanMatrixWithDisplayNamePriceCadenceSemiAnnual

// This is an alias to an internal value.
const NewPlanMatrixWithDisplayNamePriceCadenceMonthly = shared.NewPlanMatrixWithDisplayNamePriceCadenceMonthly

// This is an alias to an internal value.
const NewPlanMatrixWithDisplayNamePriceCadenceQuarterly = shared.NewPlanMatrixWithDisplayNamePriceCadenceQuarterly

// This is an alias to an internal value.
const NewPlanMatrixWithDisplayNamePriceCadenceOneTime = shared.NewPlanMatrixWithDisplayNamePriceCadenceOneTime

// This is an alias to an internal value.
const NewPlanMatrixWithDisplayNamePriceCadenceCustom = shared.NewPlanMatrixWithDisplayNamePriceCadenceCustom

// Configuration for matrix_with_display_name pricing
//
// This is an alias to an internal type.
type NewPlanMatrixWithDisplayNamePriceMatrixWithDisplayNameConfigParam = shared.NewPlanMatrixWithDisplayNamePriceMatrixWithDisplayNameConfigParam

// Configuration for a unit amount item
//
// This is an alias to an internal type.
type NewPlanMatrixWithDisplayNamePriceMatrixWithDisplayNameConfigUnitAmountParam = shared.NewPlanMatrixWithDisplayNamePriceMatrixWithDisplayNameConfigUnitAmountParam

// The pricing model type
//
// This is an alias to an internal type.
type NewPlanMatrixWithDisplayNamePriceModelType = shared.NewPlanMatrixWithDisplayNamePriceModelType

// This is an alias to an internal value.
const NewPlanMatrixWithDisplayNamePriceModelTypeMatrixWithDisplayName = shared.NewPlanMatrixWithDisplayNamePriceModelTypeMatrixWithDisplayName

// This is an alias to an internal type.
type NewPlanMatrixWithDisplayNamePriceConversionRateConfigUnionParam = shared.NewPlanMatrixWithDisplayNamePriceConversionRateConfigUnionParam

// This is an alias to an internal type.
type NewPlanMatrixWithDisplayNamePriceConversionRateConfigConversionRateType = shared.NewPlanMatrixWithDisplayNamePriceConversionRateConfigConversionRateType

// This is an alias to an internal value.
const NewPlanMatrixWithDisplayNamePriceConversionRateConfigConversionRateTypeUnit = shared.NewPlanMatrixWithDisplayNamePriceConversionRateConfigConversionRateTypeUnit

// This is an alias to an internal value.
const NewPlanMatrixWithDisplayNamePriceConversionRateConfigConversionRateTypeTiered = shared.NewPlanMatrixWithDisplayNamePriceConversionRateConfigConversionRateTypeTiered

// This is an alias to an internal type.
type NewPlanMaxGroupTieredPackagePriceParam = shared.NewPlanMaxGroupTieredPackagePriceParam

// The cadence to bill for this price on.
//
// This is an alias to an internal type.
type NewPlanMaxGroupTieredPackagePriceCadence = shared.NewPlanMaxGroupTieredPackagePriceCadence

// This is an alias to an internal value.
const NewPlanMaxGroupTieredPackagePriceCadenceAnnual = shared.NewPlanMaxGroupTieredPackagePriceCadenceAnnual

// This is an alias to an internal value.
const NewPlanMaxGroupTieredPackagePriceCadenceSemiAnnual = shared.NewPlanMaxGroupTieredPackagePriceCadenceSemiAnnual

// This is an alias to an internal value.
const NewPlanMaxGroupTieredPackagePriceCadenceMonthly = shared.NewPlanMaxGroupTieredPackagePriceCadenceMonthly

// This is an alias to an internal value.
const NewPlanMaxGroupTieredPackagePriceCadenceQuarterly = shared.NewPlanMaxGroupTieredPackagePriceCadenceQuarterly

// This is an alias to an internal value.
const NewPlanMaxGroupTieredPackagePriceCadenceOneTime = shared.NewPlanMaxGroupTieredPackagePriceCadenceOneTime

// This is an alias to an internal value.
const NewPlanMaxGroupTieredPackagePriceCadenceCustom = shared.NewPlanMaxGroupTieredPackagePriceCadenceCustom

// Configuration for max_group_tiered_package pricing
//
// This is an alias to an internal type.
type NewPlanMaxGroupTieredPackagePriceMaxGroupTieredPackageConfigParam = shared.NewPlanMaxGroupTieredPackagePriceMaxGroupTieredPackageConfigParam

// Configuration for a single tier
//
// This is an alias to an internal type.
type NewPlanMaxGroupTieredPackagePriceMaxGroupTieredPackageConfigTierParam = shared.NewPlanMaxGroupTieredPackagePriceMaxGroupTieredPackageConfigTierParam

// The pricing model type
//
// This is an alias to an internal type.
type NewPlanMaxGroupTieredPackagePriceModelType = shared.NewPlanMaxGroupTieredPackagePriceModelType

// This is an alias to an internal value.
const NewPlanMaxGroupTieredPackagePriceModelTypeMaxGroupTieredPackage = shared.NewPlanMaxGroupTieredPackagePriceModelTypeMaxGroupTieredPackage

// This is an alias to an internal type.
type NewPlanMaxGroupTieredPackagePriceConversionRateConfigUnionParam = shared.NewPlanMaxGroupTieredPackagePriceConversionRateConfigUnionParam

// This is an alias to an internal type.
type NewPlanMaxGroupTieredPackagePriceConversionRateConfigConversionRateType = shared.NewPlanMaxGroupTieredPackagePriceConversionRateConfigConversionRateType

// This is an alias to an internal value.
const NewPlanMaxGroupTieredPackagePriceConversionRateConfigConversionRateTypeUnit = shared.NewPlanMaxGroupTieredPackagePriceConversionRateConfigConversionRateTypeUnit

// This is an alias to an internal value.
const NewPlanMaxGroupTieredPackagePriceConversionRateConfigConversionRateTypeTiered = shared.NewPlanMaxGroupTieredPackagePriceConversionRateConfigConversionRateTypeTiered

// This is an alias to an internal type.
type NewPlanMinimumCompositePriceParam = shared.NewPlanMinimumCompositePriceParam

// The cadence to bill for this price on.
//
// This is an alias to an internal type.
type NewPlanMinimumCompositePriceCadence = shared.NewPlanMinimumCompositePriceCadence

// This is an alias to an internal value.
const NewPlanMinimumCompositePriceCadenceAnnual = shared.NewPlanMinimumCompositePriceCadenceAnnual

// This is an alias to an internal value.
const NewPlanMinimumCompositePriceCadenceSemiAnnual = shared.NewPlanMinimumCompositePriceCadenceSemiAnnual

// This is an alias to an internal value.
const NewPlanMinimumCompositePriceCadenceMonthly = shared.NewPlanMinimumCompositePriceCadenceMonthly

// This is an alias to an internal value.
const NewPlanMinimumCompositePriceCadenceQuarterly = shared.NewPlanMinimumCompositePriceCadenceQuarterly

// This is an alias to an internal value.
const NewPlanMinimumCompositePriceCadenceOneTime = shared.NewPlanMinimumCompositePriceCadenceOneTime

// This is an alias to an internal value.
const NewPlanMinimumCompositePriceCadenceCustom = shared.NewPlanMinimumCompositePriceCadenceCustom

// Configuration for minimum_composite pricing
//
// This is an alias to an internal type.
type NewPlanMinimumCompositePriceMinimumCompositeConfigParam = shared.NewPlanMinimumCompositePriceMinimumCompositeConfigParam

// The pricing model type
//
// This is an alias to an internal type.
type NewPlanMinimumCompositePriceModelType = shared.NewPlanMinimumCompositePriceModelType

// This is an alias to an internal value.
const NewPlanMinimumCompositePriceModelTypeMinimumComposite = shared.NewPlanMinimumCompositePriceModelTypeMinimumComposite

// This is an alias to an internal type.
type NewPlanMinimumCompositePriceConversionRateConfigUnionParam = shared.NewPlanMinimumCompositePriceConversionRateConfigUnionParam

// This is an alias to an internal type.
type NewPlanMinimumCompositePriceConversionRateConfigConversionRateType = shared.NewPlanMinimumCompositePriceConversionRateConfigConversionRateType

// This is an alias to an internal value.
const NewPlanMinimumCompositePriceConversionRateConfigConversionRateTypeUnit = shared.NewPlanMinimumCompositePriceConversionRateConfigConversionRateTypeUnit

// This is an alias to an internal value.
const NewPlanMinimumCompositePriceConversionRateConfigConversionRateTypeTiered = shared.NewPlanMinimumCompositePriceConversionRateConfigConversionRateTypeTiered

// This is an alias to an internal type.
type NewPlanPackagePriceParam = shared.NewPlanPackagePriceParam

// The cadence to bill for this price on.
//
// This is an alias to an internal type.
type NewPlanPackagePriceCadence = shared.NewPlanPackagePriceCadence

// This is an alias to an internal value.
const NewPlanPackagePriceCadenceAnnual = shared.NewPlanPackagePriceCadenceAnnual

// This is an alias to an internal value.
const NewPlanPackagePriceCadenceSemiAnnual = shared.NewPlanPackagePriceCadenceSemiAnnual

// This is an alias to an internal value.
const NewPlanPackagePriceCadenceMonthly = shared.NewPlanPackagePriceCadenceMonthly

// This is an alias to an internal value.
const NewPlanPackagePriceCadenceQuarterly = shared.NewPlanPackagePriceCadenceQuarterly

// This is an alias to an internal value.
const NewPlanPackagePriceCadenceOneTime = shared.NewPlanPackagePriceCadenceOneTime

// This is an alias to an internal value.
const NewPlanPackagePriceCadenceCustom = shared.NewPlanPackagePriceCadenceCustom

// The pricing model type
//
// This is an alias to an internal type.
type NewPlanPackagePriceModelType = shared.NewPlanPackagePriceModelType

// This is an alias to an internal value.
const NewPlanPackagePriceModelTypePackage = shared.NewPlanPackagePriceModelTypePackage

// This is an alias to an internal type.
type NewPlanPackagePriceConversionRateConfigUnionParam = shared.NewPlanPackagePriceConversionRateConfigUnionParam

// This is an alias to an internal type.
type NewPlanPackagePriceConversionRateConfigConversionRateType = shared.NewPlanPackagePriceConversionRateConfigConversionRateType

// This is an alias to an internal value.
const NewPlanPackagePriceConversionRateConfigConversionRateTypeUnit = shared.NewPlanPackagePriceConversionRateConfigConversionRateTypeUnit

// This is an alias to an internal value.
const NewPlanPackagePriceConversionRateConfigConversionRateTypeTiered = shared.NewPlanPackagePriceConversionRateConfigConversionRateTypeTiered

// This is an alias to an internal type.
type NewPlanPackageWithAllocationPriceParam = shared.NewPlanPackageWithAllocationPriceParam

// The cadence to bill for this price on.
//
// This is an alias to an internal type.
type NewPlanPackageWithAllocationPriceCadence = shared.NewPlanPackageWithAllocationPriceCadence

// This is an alias to an internal value.
const NewPlanPackageWithAllocationPriceCadenceAnnual = shared.NewPlanPackageWithAllocationPriceCadenceAnnual

// This is an alias to an internal value.
const NewPlanPackageWithAllocationPriceCadenceSemiAnnual = shared.NewPlanPackageWithAllocationPriceCadenceSemiAnnual

// This is an alias to an internal value.
const NewPlanPackageWithAllocationPriceCadenceMonthly = shared.NewPlanPackageWithAllocationPriceCadenceMonthly

// This is an alias to an internal value.
const NewPlanPackageWithAllocationPriceCadenceQuarterly = shared.NewPlanPackageWithAllocationPriceCadenceQuarterly

// This is an alias to an internal value.
const NewPlanPackageWithAllocationPriceCadenceOneTime = shared.NewPlanPackageWithAllocationPriceCadenceOneTime

// This is an alias to an internal value.
const NewPlanPackageWithAllocationPriceCadenceCustom = shared.NewPlanPackageWithAllocationPriceCadenceCustom

// The pricing model type
//
// This is an alias to an internal type.
type NewPlanPackageWithAllocationPriceModelType = shared.NewPlanPackageWithAllocationPriceModelType

// This is an alias to an internal value.
const NewPlanPackageWithAllocationPriceModelTypePackageWithAllocation = shared.NewPlanPackageWithAllocationPriceModelTypePackageWithAllocation

// Configuration for package_with_allocation pricing
//
// This is an alias to an internal type.
type NewPlanPackageWithAllocationPricePackageWithAllocationConfigParam = shared.NewPlanPackageWithAllocationPricePackageWithAllocationConfigParam

// This is an alias to an internal type.
type NewPlanPackageWithAllocationPriceConversionRateConfigUnionParam = shared.NewPlanPackageWithAllocationPriceConversionRateConfigUnionParam

// This is an alias to an internal type.
type NewPlanPackageWithAllocationPriceConversionRateConfigConversionRateType = shared.NewPlanPackageWithAllocationPriceConversionRateConfigConversionRateType

// This is an alias to an internal value.
const NewPlanPackageWithAllocationPriceConversionRateConfigConversionRateTypeUnit = shared.NewPlanPackageWithAllocationPriceConversionRateConfigConversionRateTypeUnit

// This is an alias to an internal value.
const NewPlanPackageWithAllocationPriceConversionRateConfigConversionRateTypeTiered = shared.NewPlanPackageWithAllocationPriceConversionRateConfigConversionRateTypeTiered

// This is an alias to an internal type.
type NewPlanScalableMatrixWithTieredPricingPriceParam = shared.NewPlanScalableMatrixWithTieredPricingPriceParam

// The cadence to bill for this price on.
//
// This is an alias to an internal type.
type NewPlanScalableMatrixWithTieredPricingPriceCadence = shared.NewPlanScalableMatrixWithTieredPricingPriceCadence

// This is an alias to an internal value.
const NewPlanScalableMatrixWithTieredPricingPriceCadenceAnnual = shared.NewPlanScalableMatrixWithTieredPricingPriceCadenceAnnual

// This is an alias to an internal value.
const NewPlanScalableMatrixWithTieredPricingPriceCadenceSemiAnnual = shared.NewPlanScalableMatrixWithTieredPricingPriceCadenceSemiAnnual

// This is an alias to an internal value.
const NewPlanScalableMatrixWithTieredPricingPriceCadenceMonthly = shared.NewPlanScalableMatrixWithTieredPricingPriceCadenceMonthly

// This is an alias to an internal value.
const NewPlanScalableMatrixWithTieredPricingPriceCadenceQuarterly = shared.NewPlanScalableMatrixWithTieredPricingPriceCadenceQuarterly

// This is an alias to an internal value.
const NewPlanScalableMatrixWithTieredPricingPriceCadenceOneTime = shared.NewPlanScalableMatrixWithTieredPricingPriceCadenceOneTime

// This is an alias to an internal value.
const NewPlanScalableMatrixWithTieredPricingPriceCadenceCustom = shared.NewPlanScalableMatrixWithTieredPricingPriceCadenceCustom

// The pricing model type
//
// This is an alias to an internal type.
type NewPlanScalableMatrixWithTieredPricingPriceModelType = shared.NewPlanScalableMatrixWithTieredPricingPriceModelType

// This is an alias to an internal value.
const NewPlanScalableMatrixWithTieredPricingPriceModelTypeScalableMatrixWithTieredPricing = shared.NewPlanScalableMatrixWithTieredPricingPriceModelTypeScalableMatrixWithTieredPricing

// Configuration for scalable_matrix_with_tiered_pricing pricing
//
// This is an alias to an internal type.
type NewPlanScalableMatrixWithTieredPricingPriceScalableMatrixWithTieredPricingConfigParam = shared.NewPlanScalableMatrixWithTieredPricingPriceScalableMatrixWithTieredPricingConfigParam

// Configuration for a single matrix scaling factor
//
// This is an alias to an internal type.
type NewPlanScalableMatrixWithTieredPricingPriceScalableMatrixWithTieredPricingConfigMatrixScalingFactorParam = shared.NewPlanScalableMatrixWithTieredPricingPriceScalableMatrixWithTieredPricingConfigMatrixScalingFactorParam

// Configuration for a single tier entry with business logic
//
// This is an alias to an internal type.
type NewPlanScalableMatrixWithTieredPricingPriceScalableMatrixWithTieredPricingConfigTierParam = shared.NewPlanScalableMatrixWithTieredPricingPriceScalableMatrixWithTieredPricingConfigTierParam

// This is an alias to an internal type.
type NewPlanScalableMatrixWithTieredPricingPriceConversionRateConfigUnionParam = shared.NewPlanScalableMatrixWithTieredPricingPriceConversionRateConfigUnionParam

// This is an alias to an internal type.
type NewPlanScalableMatrixWithTieredPricingPriceConversionRateConfigConversionRateType = shared.NewPlanScalableMatrixWithTieredPricingPriceConversionRateConfigConversionRateType

// This is an alias to an internal value.
const NewPlanScalableMatrixWithTieredPricingPriceConversionRateConfigConversionRateTypeUnit = shared.NewPlanScalableMatrixWithTieredPricingPriceConversionRateConfigConversionRateTypeUnit

// This is an alias to an internal value.
const NewPlanScalableMatrixWithTieredPricingPriceConversionRateConfigConversionRateTypeTiered = shared.NewPlanScalableMatrixWithTieredPricingPriceConversionRateConfigConversionRateTypeTiered

// This is an alias to an internal type.
type NewPlanScalableMatrixWithUnitPricingPriceParam = shared.NewPlanScalableMatrixWithUnitPricingPriceParam

// The cadence to bill for this price on.
//
// This is an alias to an internal type.
type NewPlanScalableMatrixWithUnitPricingPriceCadence = shared.NewPlanScalableMatrixWithUnitPricingPriceCadence

// This is an alias to an internal value.
const NewPlanScalableMatrixWithUnitPricingPriceCadenceAnnual = shared.NewPlanScalableMatrixWithUnitPricingPriceCadenceAnnual

// This is an alias to an internal value.
const NewPlanScalableMatrixWithUnitPricingPriceCadenceSemiAnnual = shared.NewPlanScalableMatrixWithUnitPricingPriceCadenceSemiAnnual

// This is an alias to an internal value.
const NewPlanScalableMatrixWithUnitPricingPriceCadenceMonthly = shared.NewPlanScalableMatrixWithUnitPricingPriceCadenceMonthly

// This is an alias to an internal value.
const NewPlanScalableMatrixWithUnitPricingPriceCadenceQuarterly = shared.NewPlanScalableMatrixWithUnitPricingPriceCadenceQuarterly

// This is an alias to an internal value.
const NewPlanScalableMatrixWithUnitPricingPriceCadenceOneTime = shared.NewPlanScalableMatrixWithUnitPricingPriceCadenceOneTime

// This is an alias to an internal value.
const NewPlanScalableMatrixWithUnitPricingPriceCadenceCustom = shared.NewPlanScalableMatrixWithUnitPricingPriceCadenceCustom

// The pricing model type
//
// This is an alias to an internal type.
type NewPlanScalableMatrixWithUnitPricingPriceModelType = shared.NewPlanScalableMatrixWithUnitPricingPriceModelType

// This is an alias to an internal value.
const NewPlanScalableMatrixWithUnitPricingPriceModelTypeScalableMatrixWithUnitPricing = shared.NewPlanScalableMatrixWithUnitPricingPriceModelTypeScalableMatrixWithUnitPricing

// Configuration for scalable_matrix_with_unit_pricing pricing
//
// This is an alias to an internal type.
type NewPlanScalableMatrixWithUnitPricingPriceScalableMatrixWithUnitPricingConfigParam = shared.NewPlanScalableMatrixWithUnitPricingPriceScalableMatrixWithUnitPricingConfigParam

// Configuration for a single matrix scaling factor
//
// This is an alias to an internal type.
type NewPlanScalableMatrixWithUnitPricingPriceScalableMatrixWithUnitPricingConfigMatrixScalingFactorParam = shared.NewPlanScalableMatrixWithUnitPricingPriceScalableMatrixWithUnitPricingConfigMatrixScalingFactorParam

// This is an alias to an internal type.
type NewPlanScalableMatrixWithUnitPricingPriceConversionRateConfigUnionParam = shared.NewPlanScalableMatrixWithUnitPricingPriceConversionRateConfigUnionParam

// This is an alias to an internal type.
type NewPlanScalableMatrixWithUnitPricingPriceConversionRateConfigConversionRateType = shared.NewPlanScalableMatrixWithUnitPricingPriceConversionRateConfigConversionRateType

// This is an alias to an internal value.
const NewPlanScalableMatrixWithUnitPricingPriceConversionRateConfigConversionRateTypeUnit = shared.NewPlanScalableMatrixWithUnitPricingPriceConversionRateConfigConversionRateTypeUnit

// This is an alias to an internal value.
const NewPlanScalableMatrixWithUnitPricingPriceConversionRateConfigConversionRateTypeTiered = shared.NewPlanScalableMatrixWithUnitPricingPriceConversionRateConfigConversionRateTypeTiered

// This is an alias to an internal type.
type NewPlanThresholdTotalAmountPriceParam = shared.NewPlanThresholdTotalAmountPriceParam

// The cadence to bill for this price on.
//
// This is an alias to an internal type.
type NewPlanThresholdTotalAmountPriceCadence = shared.NewPlanThresholdTotalAmountPriceCadence

// This is an alias to an internal value.
const NewPlanThresholdTotalAmountPriceCadenceAnnual = shared.NewPlanThresholdTotalAmountPriceCadenceAnnual

// This is an alias to an internal value.
const NewPlanThresholdTotalAmountPriceCadenceSemiAnnual = shared.NewPlanThresholdTotalAmountPriceCadenceSemiAnnual

// This is an alias to an internal value.
const NewPlanThresholdTotalAmountPriceCadenceMonthly = shared.NewPlanThresholdTotalAmountPriceCadenceMonthly

// This is an alias to an internal value.
const NewPlanThresholdTotalAmountPriceCadenceQuarterly = shared.NewPlanThresholdTotalAmountPriceCadenceQuarterly

// This is an alias to an internal value.
const NewPlanThresholdTotalAmountPriceCadenceOneTime = shared.NewPlanThresholdTotalAmountPriceCadenceOneTime

// This is an alias to an internal value.
const NewPlanThresholdTotalAmountPriceCadenceCustom = shared.NewPlanThresholdTotalAmountPriceCadenceCustom

// The pricing model type
//
// This is an alias to an internal type.
type NewPlanThresholdTotalAmountPriceModelType = shared.NewPlanThresholdTotalAmountPriceModelType

// This is an alias to an internal value.
const NewPlanThresholdTotalAmountPriceModelTypeThresholdTotalAmount = shared.NewPlanThresholdTotalAmountPriceModelTypeThresholdTotalAmount

// Configuration for threshold_total_amount pricing
//
// This is an alias to an internal type.
type NewPlanThresholdTotalAmountPriceThresholdTotalAmountConfigParam = shared.NewPlanThresholdTotalAmountPriceThresholdTotalAmountConfigParam

// Configuration for a single threshold
//
// This is an alias to an internal type.
type NewPlanThresholdTotalAmountPriceThresholdTotalAmountConfigConsumptionTableParam = shared.NewPlanThresholdTotalAmountPriceThresholdTotalAmountConfigConsumptionTableParam

// This is an alias to an internal type.
type NewPlanThresholdTotalAmountPriceConversionRateConfigUnionParam = shared.NewPlanThresholdTotalAmountPriceConversionRateConfigUnionParam

// This is an alias to an internal type.
type NewPlanThresholdTotalAmountPriceConversionRateConfigConversionRateType = shared.NewPlanThresholdTotalAmountPriceConversionRateConfigConversionRateType

// This is an alias to an internal value.
const NewPlanThresholdTotalAmountPriceConversionRateConfigConversionRateTypeUnit = shared.NewPlanThresholdTotalAmountPriceConversionRateConfigConversionRateTypeUnit

// This is an alias to an internal value.
const NewPlanThresholdTotalAmountPriceConversionRateConfigConversionRateTypeTiered = shared.NewPlanThresholdTotalAmountPriceConversionRateConfigConversionRateTypeTiered

// This is an alias to an internal type.
type NewPlanTieredPackagePriceParam = shared.NewPlanTieredPackagePriceParam

// The cadence to bill for this price on.
//
// This is an alias to an internal type.
type NewPlanTieredPackagePriceCadence = shared.NewPlanTieredPackagePriceCadence

// This is an alias to an internal value.
const NewPlanTieredPackagePriceCadenceAnnual = shared.NewPlanTieredPackagePriceCadenceAnnual

// This is an alias to an internal value.
const NewPlanTieredPackagePriceCadenceSemiAnnual = shared.NewPlanTieredPackagePriceCadenceSemiAnnual

// This is an alias to an internal value.
const NewPlanTieredPackagePriceCadenceMonthly = shared.NewPlanTieredPackagePriceCadenceMonthly

// This is an alias to an internal value.
const NewPlanTieredPackagePriceCadenceQuarterly = shared.NewPlanTieredPackagePriceCadenceQuarterly

// This is an alias to an internal value.
const NewPlanTieredPackagePriceCadenceOneTime = shared.NewPlanTieredPackagePriceCadenceOneTime

// This is an alias to an internal value.
const NewPlanTieredPackagePriceCadenceCustom = shared.NewPlanTieredPackagePriceCadenceCustom

// The pricing model type
//
// This is an alias to an internal type.
type NewPlanTieredPackagePriceModelType = shared.NewPlanTieredPackagePriceModelType

// This is an alias to an internal value.
const NewPlanTieredPackagePriceModelTypeTieredPackage = shared.NewPlanTieredPackagePriceModelTypeTieredPackage

// Configuration for tiered_package pricing
//
// This is an alias to an internal type.
type NewPlanTieredPackagePriceTieredPackageConfigParam = shared.NewPlanTieredPackagePriceTieredPackageConfigParam

// Configuration for a single tier with business logic
//
// This is an alias to an internal type.
type NewPlanTieredPackagePriceTieredPackageConfigTierParam = shared.NewPlanTieredPackagePriceTieredPackageConfigTierParam

// This is an alias to an internal type.
type NewPlanTieredPackagePriceConversionRateConfigUnionParam = shared.NewPlanTieredPackagePriceConversionRateConfigUnionParam

// This is an alias to an internal type.
type NewPlanTieredPackagePriceConversionRateConfigConversionRateType = shared.NewPlanTieredPackagePriceConversionRateConfigConversionRateType

// This is an alias to an internal value.
const NewPlanTieredPackagePriceConversionRateConfigConversionRateTypeUnit = shared.NewPlanTieredPackagePriceConversionRateConfigConversionRateTypeUnit

// This is an alias to an internal value.
const NewPlanTieredPackagePriceConversionRateConfigConversionRateTypeTiered = shared.NewPlanTieredPackagePriceConversionRateConfigConversionRateTypeTiered

// This is an alias to an internal type.
type NewPlanTieredPackageWithMinimumPriceParam = shared.NewPlanTieredPackageWithMinimumPriceParam

// The cadence to bill for this price on.
//
// This is an alias to an internal type.
type NewPlanTieredPackageWithMinimumPriceCadence = shared.NewPlanTieredPackageWithMinimumPriceCadence

// This is an alias to an internal value.
const NewPlanTieredPackageWithMinimumPriceCadenceAnnual = shared.NewPlanTieredPackageWithMinimumPriceCadenceAnnual

// This is an alias to an internal value.
const NewPlanTieredPackageWithMinimumPriceCadenceSemiAnnual = shared.NewPlanTieredPackageWithMinimumPriceCadenceSemiAnnual

// This is an alias to an internal value.
const NewPlanTieredPackageWithMinimumPriceCadenceMonthly = shared.NewPlanTieredPackageWithMinimumPriceCadenceMonthly

// This is an alias to an internal value.
const NewPlanTieredPackageWithMinimumPriceCadenceQuarterly = shared.NewPlanTieredPackageWithMinimumPriceCadenceQuarterly

// This is an alias to an internal value.
const NewPlanTieredPackageWithMinimumPriceCadenceOneTime = shared.NewPlanTieredPackageWithMinimumPriceCadenceOneTime

// This is an alias to an internal value.
const NewPlanTieredPackageWithMinimumPriceCadenceCustom = shared.NewPlanTieredPackageWithMinimumPriceCadenceCustom

// The pricing model type
//
// This is an alias to an internal type.
type NewPlanTieredPackageWithMinimumPriceModelType = shared.NewPlanTieredPackageWithMinimumPriceModelType

// This is an alias to an internal value.
const NewPlanTieredPackageWithMinimumPriceModelTypeTieredPackageWithMinimum = shared.NewPlanTieredPackageWithMinimumPriceModelTypeTieredPackageWithMinimum

// Configuration for tiered_package_with_minimum pricing
//
// This is an alias to an internal type.
type NewPlanTieredPackageWithMinimumPriceTieredPackageWithMinimumConfigParam = shared.NewPlanTieredPackageWithMinimumPriceTieredPackageWithMinimumConfigParam

// Configuration for a single tier
//
// This is an alias to an internal type.
type NewPlanTieredPackageWithMinimumPriceTieredPackageWithMinimumConfigTierParam = shared.NewPlanTieredPackageWithMinimumPriceTieredPackageWithMinimumConfigTierParam

// This is an alias to an internal type.
type NewPlanTieredPackageWithMinimumPriceConversionRateConfigUnionParam = shared.NewPlanTieredPackageWithMinimumPriceConversionRateConfigUnionParam

// This is an alias to an internal type.
type NewPlanTieredPackageWithMinimumPriceConversionRateConfigConversionRateType = shared.NewPlanTieredPackageWithMinimumPriceConversionRateConfigConversionRateType

// This is an alias to an internal value.
const NewPlanTieredPackageWithMinimumPriceConversionRateConfigConversionRateTypeUnit = shared.NewPlanTieredPackageWithMinimumPriceConversionRateConfigConversionRateTypeUnit

// This is an alias to an internal value.
const NewPlanTieredPackageWithMinimumPriceConversionRateConfigConversionRateTypeTiered = shared.NewPlanTieredPackageWithMinimumPriceConversionRateConfigConversionRateTypeTiered

// This is an alias to an internal type.
type NewPlanTieredPriceParam = shared.NewPlanTieredPriceParam

// The cadence to bill for this price on.
//
// This is an alias to an internal type.
type NewPlanTieredPriceCadence = shared.NewPlanTieredPriceCadence

// This is an alias to an internal value.
const NewPlanTieredPriceCadenceAnnual = shared.NewPlanTieredPriceCadenceAnnual

// This is an alias to an internal value.
const NewPlanTieredPriceCadenceSemiAnnual = shared.NewPlanTieredPriceCadenceSemiAnnual

// This is an alias to an internal value.
const NewPlanTieredPriceCadenceMonthly = shared.NewPlanTieredPriceCadenceMonthly

// This is an alias to an internal value.
const NewPlanTieredPriceCadenceQuarterly = shared.NewPlanTieredPriceCadenceQuarterly

// This is an alias to an internal value.
const NewPlanTieredPriceCadenceOneTime = shared.NewPlanTieredPriceCadenceOneTime

// This is an alias to an internal value.
const NewPlanTieredPriceCadenceCustom = shared.NewPlanTieredPriceCadenceCustom

// The pricing model type
//
// This is an alias to an internal type.
type NewPlanTieredPriceModelType = shared.NewPlanTieredPriceModelType

// This is an alias to an internal value.
const NewPlanTieredPriceModelTypeTiered = shared.NewPlanTieredPriceModelTypeTiered

// This is an alias to an internal type.
type NewPlanTieredPriceConversionRateConfigUnionParam = shared.NewPlanTieredPriceConversionRateConfigUnionParam

// This is an alias to an internal type.
type NewPlanTieredPriceConversionRateConfigConversionRateType = shared.NewPlanTieredPriceConversionRateConfigConversionRateType

// This is an alias to an internal value.
const NewPlanTieredPriceConversionRateConfigConversionRateTypeUnit = shared.NewPlanTieredPriceConversionRateConfigConversionRateTypeUnit

// This is an alias to an internal value.
const NewPlanTieredPriceConversionRateConfigConversionRateTypeTiered = shared.NewPlanTieredPriceConversionRateConfigConversionRateTypeTiered

// This is an alias to an internal type.
type NewPlanTieredWithMinimumPriceParam = shared.NewPlanTieredWithMinimumPriceParam

// The cadence to bill for this price on.
//
// This is an alias to an internal type.
type NewPlanTieredWithMinimumPriceCadence = shared.NewPlanTieredWithMinimumPriceCadence

// This is an alias to an internal value.
const NewPlanTieredWithMinimumPriceCadenceAnnual = shared.NewPlanTieredWithMinimumPriceCadenceAnnual

// This is an alias to an internal value.
const NewPlanTieredWithMinimumPriceCadenceSemiAnnual = shared.NewPlanTieredWithMinimumPriceCadenceSemiAnnual

// This is an alias to an internal value.
const NewPlanTieredWithMinimumPriceCadenceMonthly = shared.NewPlanTieredWithMinimumPriceCadenceMonthly

// This is an alias to an internal value.
const NewPlanTieredWithMinimumPriceCadenceQuarterly = shared.NewPlanTieredWithMinimumPriceCadenceQuarterly

// This is an alias to an internal value.
const NewPlanTieredWithMinimumPriceCadenceOneTime = shared.NewPlanTieredWithMinimumPriceCadenceOneTime

// This is an alias to an internal value.
const NewPlanTieredWithMinimumPriceCadenceCustom = shared.NewPlanTieredWithMinimumPriceCadenceCustom

// The pricing model type
//
// This is an alias to an internal type.
type NewPlanTieredWithMinimumPriceModelType = shared.NewPlanTieredWithMinimumPriceModelType

// This is an alias to an internal value.
const NewPlanTieredWithMinimumPriceModelTypeTieredWithMinimum = shared.NewPlanTieredWithMinimumPriceModelTypeTieredWithMinimum

// Configuration for tiered_with_minimum pricing
//
// This is an alias to an internal type.
type NewPlanTieredWithMinimumPriceTieredWithMinimumConfigParam = shared.NewPlanTieredWithMinimumPriceTieredWithMinimumConfigParam

// Configuration for a single tier
//
// This is an alias to an internal type.
type NewPlanTieredWithMinimumPriceTieredWithMinimumConfigTierParam = shared.NewPlanTieredWithMinimumPriceTieredWithMinimumConfigTierParam

// This is an alias to an internal type.
type NewPlanTieredWithMinimumPriceConversionRateConfigUnionParam = shared.NewPlanTieredWithMinimumPriceConversionRateConfigUnionParam

// This is an alias to an internal type.
type NewPlanTieredWithMinimumPriceConversionRateConfigConversionRateType = shared.NewPlanTieredWithMinimumPriceConversionRateConfigConversionRateType

// This is an alias to an internal value.
const NewPlanTieredWithMinimumPriceConversionRateConfigConversionRateTypeUnit = shared.NewPlanTieredWithMinimumPriceConversionRateConfigConversionRateTypeUnit

// This is an alias to an internal value.
const NewPlanTieredWithMinimumPriceConversionRateConfigConversionRateTypeTiered = shared.NewPlanTieredWithMinimumPriceConversionRateConfigConversionRateTypeTiered

// This is an alias to an internal type.
type NewPlanUnitPriceParam = shared.NewPlanUnitPriceParam

// The cadence to bill for this price on.
//
// This is an alias to an internal type.
type NewPlanUnitPriceCadence = shared.NewPlanUnitPriceCadence

// This is an alias to an internal value.
const NewPlanUnitPriceCadenceAnnual = shared.NewPlanUnitPriceCadenceAnnual

// This is an alias to an internal value.
const NewPlanUnitPriceCadenceSemiAnnual = shared.NewPlanUnitPriceCadenceSemiAnnual

// This is an alias to an internal value.
const NewPlanUnitPriceCadenceMonthly = shared.NewPlanUnitPriceCadenceMonthly

// This is an alias to an internal value.
const NewPlanUnitPriceCadenceQuarterly = shared.NewPlanUnitPriceCadenceQuarterly

// This is an alias to an internal value.
const NewPlanUnitPriceCadenceOneTime = shared.NewPlanUnitPriceCadenceOneTime

// This is an alias to an internal value.
const NewPlanUnitPriceCadenceCustom = shared.NewPlanUnitPriceCadenceCustom

// The pricing model type
//
// This is an alias to an internal type.
type NewPlanUnitPriceModelType = shared.NewPlanUnitPriceModelType

// This is an alias to an internal value.
const NewPlanUnitPriceModelTypeUnit = shared.NewPlanUnitPriceModelTypeUnit

// This is an alias to an internal type.
type NewPlanUnitPriceConversionRateConfigUnionParam = shared.NewPlanUnitPriceConversionRateConfigUnionParam

// This is an alias to an internal type.
type NewPlanUnitPriceConversionRateConfigConversionRateType = shared.NewPlanUnitPriceConversionRateConfigConversionRateType

// This is an alias to an internal value.
const NewPlanUnitPriceConversionRateConfigConversionRateTypeUnit = shared.NewPlanUnitPriceConversionRateConfigConversionRateTypeUnit

// This is an alias to an internal value.
const NewPlanUnitPriceConversionRateConfigConversionRateTypeTiered = shared.NewPlanUnitPriceConversionRateConfigConversionRateTypeTiered

// This is an alias to an internal type.
type NewPlanUnitWithPercentPriceParam = shared.NewPlanUnitWithPercentPriceParam

// The cadence to bill for this price on.
//
// This is an alias to an internal type.
type NewPlanUnitWithPercentPriceCadence = shared.NewPlanUnitWithPercentPriceCadence

// This is an alias to an internal value.
const NewPlanUnitWithPercentPriceCadenceAnnual = shared.NewPlanUnitWithPercentPriceCadenceAnnual

// This is an alias to an internal value.
const NewPlanUnitWithPercentPriceCadenceSemiAnnual = shared.NewPlanUnitWithPercentPriceCadenceSemiAnnual

// This is an alias to an internal value.
const NewPlanUnitWithPercentPriceCadenceMonthly = shared.NewPlanUnitWithPercentPriceCadenceMonthly

// This is an alias to an internal value.
const NewPlanUnitWithPercentPriceCadenceQuarterly = shared.NewPlanUnitWithPercentPriceCadenceQuarterly

// This is an alias to an internal value.
const NewPlanUnitWithPercentPriceCadenceOneTime = shared.NewPlanUnitWithPercentPriceCadenceOneTime

// This is an alias to an internal value.
const NewPlanUnitWithPercentPriceCadenceCustom = shared.NewPlanUnitWithPercentPriceCadenceCustom

// The pricing model type
//
// This is an alias to an internal type.
type NewPlanUnitWithPercentPriceModelType = shared.NewPlanUnitWithPercentPriceModelType

// This is an alias to an internal value.
const NewPlanUnitWithPercentPriceModelTypeUnitWithPercent = shared.NewPlanUnitWithPercentPriceModelTypeUnitWithPercent

// Configuration for unit_with_percent pricing
//
// This is an alias to an internal type.
type NewPlanUnitWithPercentPriceUnitWithPercentConfigParam = shared.NewPlanUnitWithPercentPriceUnitWithPercentConfigParam

// This is an alias to an internal type.
type NewPlanUnitWithPercentPriceConversionRateConfigUnionParam = shared.NewPlanUnitWithPercentPriceConversionRateConfigUnionParam

// This is an alias to an internal type.
type NewPlanUnitWithPercentPriceConversionRateConfigConversionRateType = shared.NewPlanUnitWithPercentPriceConversionRateConfigConversionRateType

// This is an alias to an internal value.
const NewPlanUnitWithPercentPriceConversionRateConfigConversionRateTypeUnit = shared.NewPlanUnitWithPercentPriceConversionRateConfigConversionRateTypeUnit

// This is an alias to an internal value.
const NewPlanUnitWithPercentPriceConversionRateConfigConversionRateTypeTiered = shared.NewPlanUnitWithPercentPriceConversionRateConfigConversionRateTypeTiered

// This is an alias to an internal type.
type NewPlanUnitWithProrationPriceParam = shared.NewPlanUnitWithProrationPriceParam

// The cadence to bill for this price on.
//
// This is an alias to an internal type.
type NewPlanUnitWithProrationPriceCadence = shared.NewPlanUnitWithProrationPriceCadence

// This is an alias to an internal value.
const NewPlanUnitWithProrationPriceCadenceAnnual = shared.NewPlanUnitWithProrationPriceCadenceAnnual

// This is an alias to an internal value.
const NewPlanUnitWithProrationPriceCadenceSemiAnnual = shared.NewPlanUnitWithProrationPriceCadenceSemiAnnual

// This is an alias to an internal value.
const NewPlanUnitWithProrationPriceCadenceMonthly = shared.NewPlanUnitWithProrationPriceCadenceMonthly

// This is an alias to an internal value.
const NewPlanUnitWithProrationPriceCadenceQuarterly = shared.NewPlanUnitWithProrationPriceCadenceQuarterly

// This is an alias to an internal value.
const NewPlanUnitWithProrationPriceCadenceOneTime = shared.NewPlanUnitWithProrationPriceCadenceOneTime

// This is an alias to an internal value.
const NewPlanUnitWithProrationPriceCadenceCustom = shared.NewPlanUnitWithProrationPriceCadenceCustom

// The pricing model type
//
// This is an alias to an internal type.
type NewPlanUnitWithProrationPriceModelType = shared.NewPlanUnitWithProrationPriceModelType

// This is an alias to an internal value.
const NewPlanUnitWithProrationPriceModelTypeUnitWithProration = shared.NewPlanUnitWithProrationPriceModelTypeUnitWithProration

// Configuration for unit_with_proration pricing
//
// This is an alias to an internal type.
type NewPlanUnitWithProrationPriceUnitWithProrationConfigParam = shared.NewPlanUnitWithProrationPriceUnitWithProrationConfigParam

// This is an alias to an internal type.
type NewPlanUnitWithProrationPriceConversionRateConfigUnionParam = shared.NewPlanUnitWithProrationPriceConversionRateConfigUnionParam

// This is an alias to an internal type.
type NewPlanUnitWithProrationPriceConversionRateConfigConversionRateType = shared.NewPlanUnitWithProrationPriceConversionRateConfigConversionRateType

// This is an alias to an internal value.
const NewPlanUnitWithProrationPriceConversionRateConfigConversionRateTypeUnit = shared.NewPlanUnitWithProrationPriceConversionRateConfigConversionRateTypeUnit

// This is an alias to an internal value.
const NewPlanUnitWithProrationPriceConversionRateConfigConversionRateTypeTiered = shared.NewPlanUnitWithProrationPriceConversionRateConfigConversionRateTypeTiered

// This is an alias to an internal type.
type NewUsageDiscountParam = shared.NewUsageDiscountParam

// This is an alias to an internal type.
type NewUsageDiscountAdjustmentType = shared.NewUsageDiscountAdjustmentType

// This is an alias to an internal value.
const NewUsageDiscountAdjustmentTypeUsageDiscount = shared.NewUsageDiscountAdjustmentTypeUsageDiscount

// If set, the adjustment will apply to every price on the subscription.
//
// This is an alias to an internal type.
type NewUsageDiscountAppliesToAll = shared.NewUsageDiscountAppliesToAll

// This is an alias to an internal value.
const NewUsageDiscountAppliesToAllTrue = shared.NewUsageDiscountAppliesToAllTrue

// This is an alias to an internal type.
type NewUsageDiscountFilterParam = shared.NewUsageDiscountFilterParam

// The property of the price to filter on.
//
// This is an alias to an internal type.
type NewUsageDiscountFiltersField = shared.NewUsageDiscountFiltersField

// This is an alias to an internal value.
const NewUsageDiscountFiltersFieldPriceID = shared.NewUsageDiscountFiltersFieldPriceID

// This is an alias to an internal value.
const NewUsageDiscountFiltersFieldItemID = shared.NewUsageDiscountFiltersFieldItemID

// This is an alias to an internal value.
const NewUsageDiscountFiltersFieldPriceType = shared.NewUsageDiscountFiltersFieldPriceType

// This is an alias to an internal value.
const NewUsageDiscountFiltersFieldCurrency = shared.NewUsageDiscountFiltersFieldCurrency

// This is an alias to an internal value.
const NewUsageDiscountFiltersFieldPricingUnitID = shared.NewUsageDiscountFiltersFieldPricingUnitID

// Should prices that match the filter be included or excluded.
//
// This is an alias to an internal type.
type NewUsageDiscountFiltersOperator = shared.NewUsageDiscountFiltersOperator

// This is an alias to an internal value.
const NewUsageDiscountFiltersOperatorIncludes = shared.NewUsageDiscountFiltersOperatorIncludes

// This is an alias to an internal value.
const NewUsageDiscountFiltersOperatorExcludes = shared.NewUsageDiscountFiltersOperatorExcludes

// If set, only prices of the specified type will have the adjustment applied.
//
// This is an alias to an internal type.
type NewUsageDiscountPriceType = shared.NewUsageDiscountPriceType

// This is an alias to an internal value.
const NewUsageDiscountPriceTypeUsage = shared.NewUsageDiscountPriceTypeUsage

// This is an alias to an internal value.
const NewUsageDiscountPriceTypeFixedInAdvance = shared.NewUsageDiscountPriceTypeFixedInAdvance

// This is an alias to an internal value.
const NewUsageDiscountPriceTypeFixedInArrears = shared.NewUsageDiscountPriceTypeFixedInArrears

// This is an alias to an internal value.
const NewUsageDiscountPriceTypeFixed = shared.NewUsageDiscountPriceTypeFixed

// This is an alias to an internal value.
const NewUsageDiscountPriceTypeInArrears = shared.NewUsageDiscountPriceTypeInArrears

// This is an alias to an internal type.
type OtherSubLineItem = shared.OtherSubLineItem

// This is an alias to an internal type.
type OtherSubLineItemType = shared.OtherSubLineItemType

// This is an alias to an internal value.
const OtherSubLineItemTypeNull = shared.OtherSubLineItemTypeNull

// Configuration for package pricing
//
// This is an alias to an internal type.
type PackageConfig = shared.PackageConfig

// Configuration for package pricing
//
// This is an alias to an internal type.
type PackageConfigParam = shared.PackageConfigParam

// This is an alias to an internal type.
type PaginationMetadata = shared.PaginationMetadata

// This is an alias to an internal type.
type PerPriceCost = shared.PerPriceCost

// This is an alias to an internal type.
type PercentageDiscount = shared.PercentageDiscount

// This is an alias to an internal type.
type PercentageDiscountDiscountType = shared.PercentageDiscountDiscountType

// This is an alias to an internal value.
const PercentageDiscountDiscountTypePercentage = shared.PercentageDiscountDiscountTypePercentage

// This is an alias to an internal type.
type PercentageDiscountFilter = shared.PercentageDiscountFilter

// The property of the price to filter on.
//
// This is an alias to an internal type.
type PercentageDiscountFiltersField = shared.PercentageDiscountFiltersField

// This is an alias to an internal value.
const PercentageDiscountFiltersFieldPriceID = shared.PercentageDiscountFiltersFieldPriceID

// This is an alias to an internal value.
const PercentageDiscountFiltersFieldItemID = shared.PercentageDiscountFiltersFieldItemID

// This is an alias to an internal value.
const PercentageDiscountFiltersFieldPriceType = shared.PercentageDiscountFiltersFieldPriceType

// This is an alias to an internal value.
const PercentageDiscountFiltersFieldCurrency = shared.PercentageDiscountFiltersFieldCurrency

// This is an alias to an internal value.
const PercentageDiscountFiltersFieldPricingUnitID = shared.PercentageDiscountFiltersFieldPricingUnitID

// Should prices that match the filter be included or excluded.
//
// This is an alias to an internal type.
type PercentageDiscountFiltersOperator = shared.PercentageDiscountFiltersOperator

// This is an alias to an internal value.
const PercentageDiscountFiltersOperatorIncludes = shared.PercentageDiscountFiltersOperatorIncludes

// This is an alias to an internal value.
const PercentageDiscountFiltersOperatorExcludes = shared.PercentageDiscountFiltersOperatorExcludes

// This is an alias to an internal type.
type PercentageDiscountParam = shared.PercentageDiscountParam

// This is an alias to an internal type.
type PercentageDiscountFilterParam = shared.PercentageDiscountFilterParam

// This is an alias to an internal type.
type PercentageDiscountInterval = shared.PercentageDiscountInterval

// This is an alias to an internal type.
type PercentageDiscountIntervalDiscountType = shared.PercentageDiscountIntervalDiscountType

// This is an alias to an internal value.
const PercentageDiscountIntervalDiscountTypePercentage = shared.PercentageDiscountIntervalDiscountTypePercentage

// This is an alias to an internal type.
type PercentageDiscountIntervalFilter = shared.PercentageDiscountIntervalFilter

// The property of the price to filter on.
//
// This is an alias to an internal type.
type PercentageDiscountIntervalFiltersField = shared.PercentageDiscountIntervalFiltersField

// This is an alias to an internal value.
const PercentageDiscountIntervalFiltersFieldPriceID = shared.PercentageDiscountIntervalFiltersFieldPriceID

// This is an alias to an internal value.
const PercentageDiscountIntervalFiltersFieldItemID = shared.PercentageDiscountIntervalFiltersFieldItemID

// This is an alias to an internal value.
const PercentageDiscountIntervalFiltersFieldPriceType = shared.PercentageDiscountIntervalFiltersFieldPriceType

// This is an alias to an internal value.
const PercentageDiscountIntervalFiltersFieldCurrency = shared.PercentageDiscountIntervalFiltersFieldCurrency

// This is an alias to an internal value.
const PercentageDiscountIntervalFiltersFieldPricingUnitID = shared.PercentageDiscountIntervalFiltersFieldPricingUnitID

// Should prices that match the filter be included or excluded.
//
// This is an alias to an internal type.
type PercentageDiscountIntervalFiltersOperator = shared.PercentageDiscountIntervalFiltersOperator

// This is an alias to an internal value.
const PercentageDiscountIntervalFiltersOperatorIncludes = shared.PercentageDiscountIntervalFiltersOperatorIncludes

// This is an alias to an internal value.
const PercentageDiscountIntervalFiltersOperatorExcludes = shared.PercentageDiscountIntervalFiltersOperatorExcludes

// This is an alias to an internal type.
type PlanPhaseAmountDiscountAdjustment = shared.PlanPhaseAmountDiscountAdjustment

// This is an alias to an internal type.
type PlanPhaseAmountDiscountAdjustmentAdjustmentType = shared.PlanPhaseAmountDiscountAdjustmentAdjustmentType

// This is an alias to an internal value.
const PlanPhaseAmountDiscountAdjustmentAdjustmentTypeAmountDiscount = shared.PlanPhaseAmountDiscountAdjustmentAdjustmentTypeAmountDiscount

// This is an alias to an internal type.
type PlanPhaseAmountDiscountAdjustmentFilter = shared.PlanPhaseAmountDiscountAdjustmentFilter

// The property of the price to filter on.
//
// This is an alias to an internal type.
type PlanPhaseAmountDiscountAdjustmentFiltersField = shared.PlanPhaseAmountDiscountAdjustmentFiltersField

// This is an alias to an internal value.
const PlanPhaseAmountDiscountAdjustmentFiltersFieldPriceID = shared.PlanPhaseAmountDiscountAdjustmentFiltersFieldPriceID

// This is an alias to an internal value.
const PlanPhaseAmountDiscountAdjustmentFiltersFieldItemID = shared.PlanPhaseAmountDiscountAdjustmentFiltersFieldItemID

// This is an alias to an internal value.
const PlanPhaseAmountDiscountAdjustmentFiltersFieldPriceType = shared.PlanPhaseAmountDiscountAdjustmentFiltersFieldPriceType

// This is an alias to an internal value.
const PlanPhaseAmountDiscountAdjustmentFiltersFieldCurrency = shared.PlanPhaseAmountDiscountAdjustmentFiltersFieldCurrency

// This is an alias to an internal value.
const PlanPhaseAmountDiscountAdjustmentFiltersFieldPricingUnitID = shared.PlanPhaseAmountDiscountAdjustmentFiltersFieldPricingUnitID

// Should prices that match the filter be included or excluded.
//
// This is an alias to an internal type.
type PlanPhaseAmountDiscountAdjustmentFiltersOperator = shared.PlanPhaseAmountDiscountAdjustmentFiltersOperator

// This is an alias to an internal value.
const PlanPhaseAmountDiscountAdjustmentFiltersOperatorIncludes = shared.PlanPhaseAmountDiscountAdjustmentFiltersOperatorIncludes

// This is an alias to an internal value.
const PlanPhaseAmountDiscountAdjustmentFiltersOperatorExcludes = shared.PlanPhaseAmountDiscountAdjustmentFiltersOperatorExcludes

// This is an alias to an internal type.
type PlanPhaseMaximumAdjustment = shared.PlanPhaseMaximumAdjustment

// This is an alias to an internal type.
type PlanPhaseMaximumAdjustmentAdjustmentType = shared.PlanPhaseMaximumAdjustmentAdjustmentType

// This is an alias to an internal value.
const PlanPhaseMaximumAdjustmentAdjustmentTypeMaximum = shared.PlanPhaseMaximumAdjustmentAdjustmentTypeMaximum

// This is an alias to an internal type.
type PlanPhaseMaximumAdjustmentFilter = shared.PlanPhaseMaximumAdjustmentFilter

// The property of the price to filter on.
//
// This is an alias to an internal type.
type PlanPhaseMaximumAdjustmentFiltersField = shared.PlanPhaseMaximumAdjustmentFiltersField

// This is an alias to an internal value.
const PlanPhaseMaximumAdjustmentFiltersFieldPriceID = shared.PlanPhaseMaximumAdjustmentFiltersFieldPriceID

// This is an alias to an internal value.
const PlanPhaseMaximumAdjustmentFiltersFieldItemID = shared.PlanPhaseMaximumAdjustmentFiltersFieldItemID

// This is an alias to an internal value.
const PlanPhaseMaximumAdjustmentFiltersFieldPriceType = shared.PlanPhaseMaximumAdjustmentFiltersFieldPriceType

// This is an alias to an internal value.
const PlanPhaseMaximumAdjustmentFiltersFieldCurrency = shared.PlanPhaseMaximumAdjustmentFiltersFieldCurrency

// This is an alias to an internal value.
const PlanPhaseMaximumAdjustmentFiltersFieldPricingUnitID = shared.PlanPhaseMaximumAdjustmentFiltersFieldPricingUnitID

// Should prices that match the filter be included or excluded.
//
// This is an alias to an internal type.
type PlanPhaseMaximumAdjustmentFiltersOperator = shared.PlanPhaseMaximumAdjustmentFiltersOperator

// This is an alias to an internal value.
const PlanPhaseMaximumAdjustmentFiltersOperatorIncludes = shared.PlanPhaseMaximumAdjustmentFiltersOperatorIncludes

// This is an alias to an internal value.
const PlanPhaseMaximumAdjustmentFiltersOperatorExcludes = shared.PlanPhaseMaximumAdjustmentFiltersOperatorExcludes

// This is an alias to an internal type.
type PlanPhaseMinimumAdjustment = shared.PlanPhaseMinimumAdjustment

// This is an alias to an internal type.
type PlanPhaseMinimumAdjustmentAdjustmentType = shared.PlanPhaseMinimumAdjustmentAdjustmentType

// This is an alias to an internal value.
const PlanPhaseMinimumAdjustmentAdjustmentTypeMinimum = shared.PlanPhaseMinimumAdjustmentAdjustmentTypeMinimum

// This is an alias to an internal type.
type PlanPhaseMinimumAdjustmentFilter = shared.PlanPhaseMinimumAdjustmentFilter

// The property of the price to filter on.
//
// This is an alias to an internal type.
type PlanPhaseMinimumAdjustmentFiltersField = shared.PlanPhaseMinimumAdjustmentFiltersField

// This is an alias to an internal value.
const PlanPhaseMinimumAdjustmentFiltersFieldPriceID = shared.PlanPhaseMinimumAdjustmentFiltersFieldPriceID

// This is an alias to an internal value.
const PlanPhaseMinimumAdjustmentFiltersFieldItemID = shared.PlanPhaseMinimumAdjustmentFiltersFieldItemID

// This is an alias to an internal value.
const PlanPhaseMinimumAdjustmentFiltersFieldPriceType = shared.PlanPhaseMinimumAdjustmentFiltersFieldPriceType

// This is an alias to an internal value.
const PlanPhaseMinimumAdjustmentFiltersFieldCurrency = shared.PlanPhaseMinimumAdjustmentFiltersFieldCurrency

// This is an alias to an internal value.
const PlanPhaseMinimumAdjustmentFiltersFieldPricingUnitID = shared.PlanPhaseMinimumAdjustmentFiltersFieldPricingUnitID

// Should prices that match the filter be included or excluded.
//
// This is an alias to an internal type.
type PlanPhaseMinimumAdjustmentFiltersOperator = shared.PlanPhaseMinimumAdjustmentFiltersOperator

// This is an alias to an internal value.
const PlanPhaseMinimumAdjustmentFiltersOperatorIncludes = shared.PlanPhaseMinimumAdjustmentFiltersOperatorIncludes

// This is an alias to an internal value.
const PlanPhaseMinimumAdjustmentFiltersOperatorExcludes = shared.PlanPhaseMinimumAdjustmentFiltersOperatorExcludes

// This is an alias to an internal type.
type PlanPhasePercentageDiscountAdjustment = shared.PlanPhasePercentageDiscountAdjustment

// This is an alias to an internal type.
type PlanPhasePercentageDiscountAdjustmentAdjustmentType = shared.PlanPhasePercentageDiscountAdjustmentAdjustmentType

// This is an alias to an internal value.
const PlanPhasePercentageDiscountAdjustmentAdjustmentTypePercentageDiscount = shared.PlanPhasePercentageDiscountAdjustmentAdjustmentTypePercentageDiscount

// This is an alias to an internal type.
type PlanPhasePercentageDiscountAdjustmentFilter = shared.PlanPhasePercentageDiscountAdjustmentFilter

// The property of the price to filter on.
//
// This is an alias to an internal type.
type PlanPhasePercentageDiscountAdjustmentFiltersField = shared.PlanPhasePercentageDiscountAdjustmentFiltersField

// This is an alias to an internal value.
const PlanPhasePercentageDiscountAdjustmentFiltersFieldPriceID = shared.PlanPhasePercentageDiscountAdjustmentFiltersFieldPriceID

// This is an alias to an internal value.
const PlanPhasePercentageDiscountAdjustmentFiltersFieldItemID = shared.PlanPhasePercentageDiscountAdjustmentFiltersFieldItemID

// This is an alias to an internal value.
const PlanPhasePercentageDiscountAdjustmentFiltersFieldPriceType = shared.PlanPhasePercentageDiscountAdjustmentFiltersFieldPriceType

// This is an alias to an internal value.
const PlanPhasePercentageDiscountAdjustmentFiltersFieldCurrency = shared.PlanPhasePercentageDiscountAdjustmentFiltersFieldCurrency

// This is an alias to an internal value.
const PlanPhasePercentageDiscountAdjustmentFiltersFieldPricingUnitID = shared.PlanPhasePercentageDiscountAdjustmentFiltersFieldPricingUnitID

// Should prices that match the filter be included or excluded.
//
// This is an alias to an internal type.
type PlanPhasePercentageDiscountAdjustmentFiltersOperator = shared.PlanPhasePercentageDiscountAdjustmentFiltersOperator

// This is an alias to an internal value.
const PlanPhasePercentageDiscountAdjustmentFiltersOperatorIncludes = shared.PlanPhasePercentageDiscountAdjustmentFiltersOperatorIncludes

// This is an alias to an internal value.
const PlanPhasePercentageDiscountAdjustmentFiltersOperatorExcludes = shared.PlanPhasePercentageDiscountAdjustmentFiltersOperatorExcludes

// This is an alias to an internal type.
type PlanPhaseUsageDiscountAdjustment = shared.PlanPhaseUsageDiscountAdjustment

// This is an alias to an internal type.
type PlanPhaseUsageDiscountAdjustmentAdjustmentType = shared.PlanPhaseUsageDiscountAdjustmentAdjustmentType

// This is an alias to an internal value.
const PlanPhaseUsageDiscountAdjustmentAdjustmentTypeUsageDiscount = shared.PlanPhaseUsageDiscountAdjustmentAdjustmentTypeUsageDiscount

// This is an alias to an internal type.
type PlanPhaseUsageDiscountAdjustmentFilter = shared.PlanPhaseUsageDiscountAdjustmentFilter

// The property of the price to filter on.
//
// This is an alias to an internal type.
type PlanPhaseUsageDiscountAdjustmentFiltersField = shared.PlanPhaseUsageDiscountAdjustmentFiltersField

// This is an alias to an internal value.
const PlanPhaseUsageDiscountAdjustmentFiltersFieldPriceID = shared.PlanPhaseUsageDiscountAdjustmentFiltersFieldPriceID

// This is an alias to an internal value.
const PlanPhaseUsageDiscountAdjustmentFiltersFieldItemID = shared.PlanPhaseUsageDiscountAdjustmentFiltersFieldItemID

// This is an alias to an internal value.
const PlanPhaseUsageDiscountAdjustmentFiltersFieldPriceType = shared.PlanPhaseUsageDiscountAdjustmentFiltersFieldPriceType

// This is an alias to an internal value.
const PlanPhaseUsageDiscountAdjustmentFiltersFieldCurrency = shared.PlanPhaseUsageDiscountAdjustmentFiltersFieldCurrency

// This is an alias to an internal value.
const PlanPhaseUsageDiscountAdjustmentFiltersFieldPricingUnitID = shared.PlanPhaseUsageDiscountAdjustmentFiltersFieldPricingUnitID

// Should prices that match the filter be included or excluded.
//
// This is an alias to an internal type.
type PlanPhaseUsageDiscountAdjustmentFiltersOperator = shared.PlanPhaseUsageDiscountAdjustmentFiltersOperator

// This is an alias to an internal value.
const PlanPhaseUsageDiscountAdjustmentFiltersOperatorIncludes = shared.PlanPhaseUsageDiscountAdjustmentFiltersOperatorIncludes

// This is an alias to an internal value.
const PlanPhaseUsageDiscountAdjustmentFiltersOperatorExcludes = shared.PlanPhaseUsageDiscountAdjustmentFiltersOperatorExcludes

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
// This is an alias to an internal type.
type Price = shared.Price

// This is an alias to an internal type.
type PriceUnitPrice = shared.PriceUnitPrice

// This is an alias to an internal type.
type PriceUnitPriceBillingMode = shared.PriceUnitPriceBillingMode

// This is an alias to an internal value.
const PriceUnitPriceBillingModeInAdvance = shared.PriceUnitPriceBillingModeInAdvance

// This is an alias to an internal value.
const PriceUnitPriceBillingModeInArrear = shared.PriceUnitPriceBillingModeInArrear

// This is an alias to an internal type.
type PriceUnitPriceCadence = shared.PriceUnitPriceCadence

// This is an alias to an internal value.
const PriceUnitPriceCadenceOneTime = shared.PriceUnitPriceCadenceOneTime

// This is an alias to an internal value.
const PriceUnitPriceCadenceMonthly = shared.PriceUnitPriceCadenceMonthly

// This is an alias to an internal value.
const PriceUnitPriceCadenceQuarterly = shared.PriceUnitPriceCadenceQuarterly

// This is an alias to an internal value.
const PriceUnitPriceCadenceSemiAnnual = shared.PriceUnitPriceCadenceSemiAnnual

// This is an alias to an internal value.
const PriceUnitPriceCadenceAnnual = shared.PriceUnitPriceCadenceAnnual

// This is an alias to an internal value.
const PriceUnitPriceCadenceCustom = shared.PriceUnitPriceCadenceCustom

// This is an alias to an internal type.
type PriceUnitPriceCompositePriceFilter = shared.PriceUnitPriceCompositePriceFilter

// The property of the price to filter on.
//
// This is an alias to an internal type.
type PriceUnitPriceCompositePriceFiltersField = shared.PriceUnitPriceCompositePriceFiltersField

// This is an alias to an internal value.
const PriceUnitPriceCompositePriceFiltersFieldPriceID = shared.PriceUnitPriceCompositePriceFiltersFieldPriceID

// This is an alias to an internal value.
const PriceUnitPriceCompositePriceFiltersFieldItemID = shared.PriceUnitPriceCompositePriceFiltersFieldItemID

// This is an alias to an internal value.
const PriceUnitPriceCompositePriceFiltersFieldPriceType = shared.PriceUnitPriceCompositePriceFiltersFieldPriceType

// This is an alias to an internal value.
const PriceUnitPriceCompositePriceFiltersFieldCurrency = shared.PriceUnitPriceCompositePriceFiltersFieldCurrency

// This is an alias to an internal value.
const PriceUnitPriceCompositePriceFiltersFieldPricingUnitID = shared.PriceUnitPriceCompositePriceFiltersFieldPricingUnitID

// Should prices that match the filter be included or excluded.
//
// This is an alias to an internal type.
type PriceUnitPriceCompositePriceFiltersOperator = shared.PriceUnitPriceCompositePriceFiltersOperator

// This is an alias to an internal value.
const PriceUnitPriceCompositePriceFiltersOperatorIncludes = shared.PriceUnitPriceCompositePriceFiltersOperatorIncludes

// This is an alias to an internal value.
const PriceUnitPriceCompositePriceFiltersOperatorExcludes = shared.PriceUnitPriceCompositePriceFiltersOperatorExcludes

// This is an alias to an internal type.
type PriceUnitPriceConversionRateConfig = shared.PriceUnitPriceConversionRateConfig

// This is an alias to an internal type.
type PriceUnitPriceConversionRateConfigConversionRateType = shared.PriceUnitPriceConversionRateConfigConversionRateType

// This is an alias to an internal value.
const PriceUnitPriceConversionRateConfigConversionRateTypeUnit = shared.PriceUnitPriceConversionRateConfigConversionRateTypeUnit

// This is an alias to an internal value.
const PriceUnitPriceConversionRateConfigConversionRateTypeTiered = shared.PriceUnitPriceConversionRateConfigConversionRateTypeTiered

// The pricing model type
//
// This is an alias to an internal type.
type PriceUnitPriceModelType = shared.PriceUnitPriceModelType

// This is an alias to an internal value.
const PriceUnitPriceModelTypeUnit = shared.PriceUnitPriceModelTypeUnit

// This is an alias to an internal type.
type PriceUnitPricePriceType = shared.PriceUnitPricePriceType

// This is an alias to an internal value.
const PriceUnitPricePriceTypeUsagePrice = shared.PriceUnitPricePriceTypeUsagePrice

// This is an alias to an internal value.
const PriceUnitPricePriceTypeFixedPrice = shared.PriceUnitPricePriceTypeFixedPrice

// This is an alias to an internal value.
const PriceUnitPricePriceTypeCompositePrice = shared.PriceUnitPricePriceTypeCompositePrice

// The LicenseType resource represents a type of license that can be assigned to
// users. License types are used during billing by grouping metrics on the
// configured grouping key.
//
// This is an alias to an internal type.
type PriceUnitPriceLicenseType = shared.PriceUnitPriceLicenseType

// This is an alias to an internal type.
type PriceTieredPrice = shared.PriceTieredPrice

// This is an alias to an internal type.
type PriceTieredPriceBillingMode = shared.PriceTieredPriceBillingMode

// This is an alias to an internal value.
const PriceTieredPriceBillingModeInAdvance = shared.PriceTieredPriceBillingModeInAdvance

// This is an alias to an internal value.
const PriceTieredPriceBillingModeInArrear = shared.PriceTieredPriceBillingModeInArrear

// This is an alias to an internal type.
type PriceTieredPriceCadence = shared.PriceTieredPriceCadence

// This is an alias to an internal value.
const PriceTieredPriceCadenceOneTime = shared.PriceTieredPriceCadenceOneTime

// This is an alias to an internal value.
const PriceTieredPriceCadenceMonthly = shared.PriceTieredPriceCadenceMonthly

// This is an alias to an internal value.
const PriceTieredPriceCadenceQuarterly = shared.PriceTieredPriceCadenceQuarterly

// This is an alias to an internal value.
const PriceTieredPriceCadenceSemiAnnual = shared.PriceTieredPriceCadenceSemiAnnual

// This is an alias to an internal value.
const PriceTieredPriceCadenceAnnual = shared.PriceTieredPriceCadenceAnnual

// This is an alias to an internal value.
const PriceTieredPriceCadenceCustom = shared.PriceTieredPriceCadenceCustom

// This is an alias to an internal type.
type PriceTieredPriceCompositePriceFilter = shared.PriceTieredPriceCompositePriceFilter

// The property of the price to filter on.
//
// This is an alias to an internal type.
type PriceTieredPriceCompositePriceFiltersField = shared.PriceTieredPriceCompositePriceFiltersField

// This is an alias to an internal value.
const PriceTieredPriceCompositePriceFiltersFieldPriceID = shared.PriceTieredPriceCompositePriceFiltersFieldPriceID

// This is an alias to an internal value.
const PriceTieredPriceCompositePriceFiltersFieldItemID = shared.PriceTieredPriceCompositePriceFiltersFieldItemID

// This is an alias to an internal value.
const PriceTieredPriceCompositePriceFiltersFieldPriceType = shared.PriceTieredPriceCompositePriceFiltersFieldPriceType

// This is an alias to an internal value.
const PriceTieredPriceCompositePriceFiltersFieldCurrency = shared.PriceTieredPriceCompositePriceFiltersFieldCurrency

// This is an alias to an internal value.
const PriceTieredPriceCompositePriceFiltersFieldPricingUnitID = shared.PriceTieredPriceCompositePriceFiltersFieldPricingUnitID

// Should prices that match the filter be included or excluded.
//
// This is an alias to an internal type.
type PriceTieredPriceCompositePriceFiltersOperator = shared.PriceTieredPriceCompositePriceFiltersOperator

// This is an alias to an internal value.
const PriceTieredPriceCompositePriceFiltersOperatorIncludes = shared.PriceTieredPriceCompositePriceFiltersOperatorIncludes

// This is an alias to an internal value.
const PriceTieredPriceCompositePriceFiltersOperatorExcludes = shared.PriceTieredPriceCompositePriceFiltersOperatorExcludes

// This is an alias to an internal type.
type PriceTieredPriceConversionRateConfig = shared.PriceTieredPriceConversionRateConfig

// This is an alias to an internal type.
type PriceTieredPriceConversionRateConfigConversionRateType = shared.PriceTieredPriceConversionRateConfigConversionRateType

// This is an alias to an internal value.
const PriceTieredPriceConversionRateConfigConversionRateTypeUnit = shared.PriceTieredPriceConversionRateConfigConversionRateTypeUnit

// This is an alias to an internal value.
const PriceTieredPriceConversionRateConfigConversionRateTypeTiered = shared.PriceTieredPriceConversionRateConfigConversionRateTypeTiered

// The pricing model type
//
// This is an alias to an internal type.
type PriceTieredPriceModelType = shared.PriceTieredPriceModelType

// This is an alias to an internal value.
const PriceTieredPriceModelTypeTiered = shared.PriceTieredPriceModelTypeTiered

// This is an alias to an internal type.
type PriceTieredPricePriceType = shared.PriceTieredPricePriceType

// This is an alias to an internal value.
const PriceTieredPricePriceTypeUsagePrice = shared.PriceTieredPricePriceTypeUsagePrice

// This is an alias to an internal value.
const PriceTieredPricePriceTypeFixedPrice = shared.PriceTieredPricePriceTypeFixedPrice

// This is an alias to an internal value.
const PriceTieredPricePriceTypeCompositePrice = shared.PriceTieredPricePriceTypeCompositePrice

// The LicenseType resource represents a type of license that can be assigned to
// users. License types are used during billing by grouping metrics on the
// configured grouping key.
//
// This is an alias to an internal type.
type PriceTieredPriceLicenseType = shared.PriceTieredPriceLicenseType

// This is an alias to an internal type.
type PriceBulkPrice = shared.PriceBulkPrice

// This is an alias to an internal type.
type PriceBulkPriceBillingMode = shared.PriceBulkPriceBillingMode

// This is an alias to an internal value.
const PriceBulkPriceBillingModeInAdvance = shared.PriceBulkPriceBillingModeInAdvance

// This is an alias to an internal value.
const PriceBulkPriceBillingModeInArrear = shared.PriceBulkPriceBillingModeInArrear

// This is an alias to an internal type.
type PriceBulkPriceCadence = shared.PriceBulkPriceCadence

// This is an alias to an internal value.
const PriceBulkPriceCadenceOneTime = shared.PriceBulkPriceCadenceOneTime

// This is an alias to an internal value.
const PriceBulkPriceCadenceMonthly = shared.PriceBulkPriceCadenceMonthly

// This is an alias to an internal value.
const PriceBulkPriceCadenceQuarterly = shared.PriceBulkPriceCadenceQuarterly

// This is an alias to an internal value.
const PriceBulkPriceCadenceSemiAnnual = shared.PriceBulkPriceCadenceSemiAnnual

// This is an alias to an internal value.
const PriceBulkPriceCadenceAnnual = shared.PriceBulkPriceCadenceAnnual

// This is an alias to an internal value.
const PriceBulkPriceCadenceCustom = shared.PriceBulkPriceCadenceCustom

// This is an alias to an internal type.
type PriceBulkPriceCompositePriceFilter = shared.PriceBulkPriceCompositePriceFilter

// The property of the price to filter on.
//
// This is an alias to an internal type.
type PriceBulkPriceCompositePriceFiltersField = shared.PriceBulkPriceCompositePriceFiltersField

// This is an alias to an internal value.
const PriceBulkPriceCompositePriceFiltersFieldPriceID = shared.PriceBulkPriceCompositePriceFiltersFieldPriceID

// This is an alias to an internal value.
const PriceBulkPriceCompositePriceFiltersFieldItemID = shared.PriceBulkPriceCompositePriceFiltersFieldItemID

// This is an alias to an internal value.
const PriceBulkPriceCompositePriceFiltersFieldPriceType = shared.PriceBulkPriceCompositePriceFiltersFieldPriceType

// This is an alias to an internal value.
const PriceBulkPriceCompositePriceFiltersFieldCurrency = shared.PriceBulkPriceCompositePriceFiltersFieldCurrency

// This is an alias to an internal value.
const PriceBulkPriceCompositePriceFiltersFieldPricingUnitID = shared.PriceBulkPriceCompositePriceFiltersFieldPricingUnitID

// Should prices that match the filter be included or excluded.
//
// This is an alias to an internal type.
type PriceBulkPriceCompositePriceFiltersOperator = shared.PriceBulkPriceCompositePriceFiltersOperator

// This is an alias to an internal value.
const PriceBulkPriceCompositePriceFiltersOperatorIncludes = shared.PriceBulkPriceCompositePriceFiltersOperatorIncludes

// This is an alias to an internal value.
const PriceBulkPriceCompositePriceFiltersOperatorExcludes = shared.PriceBulkPriceCompositePriceFiltersOperatorExcludes

// This is an alias to an internal type.
type PriceBulkPriceConversionRateConfig = shared.PriceBulkPriceConversionRateConfig

// This is an alias to an internal type.
type PriceBulkPriceConversionRateConfigConversionRateType = shared.PriceBulkPriceConversionRateConfigConversionRateType

// This is an alias to an internal value.
const PriceBulkPriceConversionRateConfigConversionRateTypeUnit = shared.PriceBulkPriceConversionRateConfigConversionRateTypeUnit

// This is an alias to an internal value.
const PriceBulkPriceConversionRateConfigConversionRateTypeTiered = shared.PriceBulkPriceConversionRateConfigConversionRateTypeTiered

// The pricing model type
//
// This is an alias to an internal type.
type PriceBulkPriceModelType = shared.PriceBulkPriceModelType

// This is an alias to an internal value.
const PriceBulkPriceModelTypeBulk = shared.PriceBulkPriceModelTypeBulk

// This is an alias to an internal type.
type PriceBulkPricePriceType = shared.PriceBulkPricePriceType

// This is an alias to an internal value.
const PriceBulkPricePriceTypeUsagePrice = shared.PriceBulkPricePriceTypeUsagePrice

// This is an alias to an internal value.
const PriceBulkPricePriceTypeFixedPrice = shared.PriceBulkPricePriceTypeFixedPrice

// This is an alias to an internal value.
const PriceBulkPricePriceTypeCompositePrice = shared.PriceBulkPricePriceTypeCompositePrice

// The LicenseType resource represents a type of license that can be assigned to
// users. License types are used during billing by grouping metrics on the
// configured grouping key.
//
// This is an alias to an internal type.
type PriceBulkPriceLicenseType = shared.PriceBulkPriceLicenseType

// This is an alias to an internal type.
type PriceBulkWithFiltersPrice = shared.PriceBulkWithFiltersPrice

// This is an alias to an internal type.
type PriceBulkWithFiltersPriceBillingMode = shared.PriceBulkWithFiltersPriceBillingMode

// This is an alias to an internal value.
const PriceBulkWithFiltersPriceBillingModeInAdvance = shared.PriceBulkWithFiltersPriceBillingModeInAdvance

// This is an alias to an internal value.
const PriceBulkWithFiltersPriceBillingModeInArrear = shared.PriceBulkWithFiltersPriceBillingModeInArrear

// Configuration for bulk_with_filters pricing
//
// This is an alias to an internal type.
type PriceBulkWithFiltersPriceBulkWithFiltersConfig = shared.PriceBulkWithFiltersPriceBulkWithFiltersConfig

// Configuration for a single property filter
//
// This is an alias to an internal type.
type PriceBulkWithFiltersPriceBulkWithFiltersConfigFilter = shared.PriceBulkWithFiltersPriceBulkWithFiltersConfigFilter

// Configuration for a single bulk pricing tier
//
// This is an alias to an internal type.
type PriceBulkWithFiltersPriceBulkWithFiltersConfigTier = shared.PriceBulkWithFiltersPriceBulkWithFiltersConfigTier

// This is an alias to an internal type.
type PriceBulkWithFiltersPriceCadence = shared.PriceBulkWithFiltersPriceCadence

// This is an alias to an internal value.
const PriceBulkWithFiltersPriceCadenceOneTime = shared.PriceBulkWithFiltersPriceCadenceOneTime

// This is an alias to an internal value.
const PriceBulkWithFiltersPriceCadenceMonthly = shared.PriceBulkWithFiltersPriceCadenceMonthly

// This is an alias to an internal value.
const PriceBulkWithFiltersPriceCadenceQuarterly = shared.PriceBulkWithFiltersPriceCadenceQuarterly

// This is an alias to an internal value.
const PriceBulkWithFiltersPriceCadenceSemiAnnual = shared.PriceBulkWithFiltersPriceCadenceSemiAnnual

// This is an alias to an internal value.
const PriceBulkWithFiltersPriceCadenceAnnual = shared.PriceBulkWithFiltersPriceCadenceAnnual

// This is an alias to an internal value.
const PriceBulkWithFiltersPriceCadenceCustom = shared.PriceBulkWithFiltersPriceCadenceCustom

// This is an alias to an internal type.
type PriceBulkWithFiltersPriceCompositePriceFilter = shared.PriceBulkWithFiltersPriceCompositePriceFilter

// The property of the price to filter on.
//
// This is an alias to an internal type.
type PriceBulkWithFiltersPriceCompositePriceFiltersField = shared.PriceBulkWithFiltersPriceCompositePriceFiltersField

// This is an alias to an internal value.
const PriceBulkWithFiltersPriceCompositePriceFiltersFieldPriceID = shared.PriceBulkWithFiltersPriceCompositePriceFiltersFieldPriceID

// This is an alias to an internal value.
const PriceBulkWithFiltersPriceCompositePriceFiltersFieldItemID = shared.PriceBulkWithFiltersPriceCompositePriceFiltersFieldItemID

// This is an alias to an internal value.
const PriceBulkWithFiltersPriceCompositePriceFiltersFieldPriceType = shared.PriceBulkWithFiltersPriceCompositePriceFiltersFieldPriceType

// This is an alias to an internal value.
const PriceBulkWithFiltersPriceCompositePriceFiltersFieldCurrency = shared.PriceBulkWithFiltersPriceCompositePriceFiltersFieldCurrency

// This is an alias to an internal value.
const PriceBulkWithFiltersPriceCompositePriceFiltersFieldPricingUnitID = shared.PriceBulkWithFiltersPriceCompositePriceFiltersFieldPricingUnitID

// Should prices that match the filter be included or excluded.
//
// This is an alias to an internal type.
type PriceBulkWithFiltersPriceCompositePriceFiltersOperator = shared.PriceBulkWithFiltersPriceCompositePriceFiltersOperator

// This is an alias to an internal value.
const PriceBulkWithFiltersPriceCompositePriceFiltersOperatorIncludes = shared.PriceBulkWithFiltersPriceCompositePriceFiltersOperatorIncludes

// This is an alias to an internal value.
const PriceBulkWithFiltersPriceCompositePriceFiltersOperatorExcludes = shared.PriceBulkWithFiltersPriceCompositePriceFiltersOperatorExcludes

// This is an alias to an internal type.
type PriceBulkWithFiltersPriceConversionRateConfig = shared.PriceBulkWithFiltersPriceConversionRateConfig

// This is an alias to an internal type.
type PriceBulkWithFiltersPriceConversionRateConfigConversionRateType = shared.PriceBulkWithFiltersPriceConversionRateConfigConversionRateType

// This is an alias to an internal value.
const PriceBulkWithFiltersPriceConversionRateConfigConversionRateTypeUnit = shared.PriceBulkWithFiltersPriceConversionRateConfigConversionRateTypeUnit

// This is an alias to an internal value.
const PriceBulkWithFiltersPriceConversionRateConfigConversionRateTypeTiered = shared.PriceBulkWithFiltersPriceConversionRateConfigConversionRateTypeTiered

// The pricing model type
//
// This is an alias to an internal type.
type PriceBulkWithFiltersPriceModelType = shared.PriceBulkWithFiltersPriceModelType

// This is an alias to an internal value.
const PriceBulkWithFiltersPriceModelTypeBulkWithFilters = shared.PriceBulkWithFiltersPriceModelTypeBulkWithFilters

// This is an alias to an internal type.
type PriceBulkWithFiltersPricePriceType = shared.PriceBulkWithFiltersPricePriceType

// This is an alias to an internal value.
const PriceBulkWithFiltersPricePriceTypeUsagePrice = shared.PriceBulkWithFiltersPricePriceTypeUsagePrice

// This is an alias to an internal value.
const PriceBulkWithFiltersPricePriceTypeFixedPrice = shared.PriceBulkWithFiltersPricePriceTypeFixedPrice

// This is an alias to an internal value.
const PriceBulkWithFiltersPricePriceTypeCompositePrice = shared.PriceBulkWithFiltersPricePriceTypeCompositePrice

// The LicenseType resource represents a type of license that can be assigned to
// users. License types are used during billing by grouping metrics on the
// configured grouping key.
//
// This is an alias to an internal type.
type PriceBulkWithFiltersPriceLicenseType = shared.PriceBulkWithFiltersPriceLicenseType

// This is an alias to an internal type.
type PricePackagePrice = shared.PricePackagePrice

// This is an alias to an internal type.
type PricePackagePriceBillingMode = shared.PricePackagePriceBillingMode

// This is an alias to an internal value.
const PricePackagePriceBillingModeInAdvance = shared.PricePackagePriceBillingModeInAdvance

// This is an alias to an internal value.
const PricePackagePriceBillingModeInArrear = shared.PricePackagePriceBillingModeInArrear

// This is an alias to an internal type.
type PricePackagePriceCadence = shared.PricePackagePriceCadence

// This is an alias to an internal value.
const PricePackagePriceCadenceOneTime = shared.PricePackagePriceCadenceOneTime

// This is an alias to an internal value.
const PricePackagePriceCadenceMonthly = shared.PricePackagePriceCadenceMonthly

// This is an alias to an internal value.
const PricePackagePriceCadenceQuarterly = shared.PricePackagePriceCadenceQuarterly

// This is an alias to an internal value.
const PricePackagePriceCadenceSemiAnnual = shared.PricePackagePriceCadenceSemiAnnual

// This is an alias to an internal value.
const PricePackagePriceCadenceAnnual = shared.PricePackagePriceCadenceAnnual

// This is an alias to an internal value.
const PricePackagePriceCadenceCustom = shared.PricePackagePriceCadenceCustom

// This is an alias to an internal type.
type PricePackagePriceCompositePriceFilter = shared.PricePackagePriceCompositePriceFilter

// The property of the price to filter on.
//
// This is an alias to an internal type.
type PricePackagePriceCompositePriceFiltersField = shared.PricePackagePriceCompositePriceFiltersField

// This is an alias to an internal value.
const PricePackagePriceCompositePriceFiltersFieldPriceID = shared.PricePackagePriceCompositePriceFiltersFieldPriceID

// This is an alias to an internal value.
const PricePackagePriceCompositePriceFiltersFieldItemID = shared.PricePackagePriceCompositePriceFiltersFieldItemID

// This is an alias to an internal value.
const PricePackagePriceCompositePriceFiltersFieldPriceType = shared.PricePackagePriceCompositePriceFiltersFieldPriceType

// This is an alias to an internal value.
const PricePackagePriceCompositePriceFiltersFieldCurrency = shared.PricePackagePriceCompositePriceFiltersFieldCurrency

// This is an alias to an internal value.
const PricePackagePriceCompositePriceFiltersFieldPricingUnitID = shared.PricePackagePriceCompositePriceFiltersFieldPricingUnitID

// Should prices that match the filter be included or excluded.
//
// This is an alias to an internal type.
type PricePackagePriceCompositePriceFiltersOperator = shared.PricePackagePriceCompositePriceFiltersOperator

// This is an alias to an internal value.
const PricePackagePriceCompositePriceFiltersOperatorIncludes = shared.PricePackagePriceCompositePriceFiltersOperatorIncludes

// This is an alias to an internal value.
const PricePackagePriceCompositePriceFiltersOperatorExcludes = shared.PricePackagePriceCompositePriceFiltersOperatorExcludes

// This is an alias to an internal type.
type PricePackagePriceConversionRateConfig = shared.PricePackagePriceConversionRateConfig

// This is an alias to an internal type.
type PricePackagePriceConversionRateConfigConversionRateType = shared.PricePackagePriceConversionRateConfigConversionRateType

// This is an alias to an internal value.
const PricePackagePriceConversionRateConfigConversionRateTypeUnit = shared.PricePackagePriceConversionRateConfigConversionRateTypeUnit

// This is an alias to an internal value.
const PricePackagePriceConversionRateConfigConversionRateTypeTiered = shared.PricePackagePriceConversionRateConfigConversionRateTypeTiered

// The pricing model type
//
// This is an alias to an internal type.
type PricePackagePriceModelType = shared.PricePackagePriceModelType

// This is an alias to an internal value.
const PricePackagePriceModelTypePackage = shared.PricePackagePriceModelTypePackage

// This is an alias to an internal type.
type PricePackagePricePriceType = shared.PricePackagePricePriceType

// This is an alias to an internal value.
const PricePackagePricePriceTypeUsagePrice = shared.PricePackagePricePriceTypeUsagePrice

// This is an alias to an internal value.
const PricePackagePricePriceTypeFixedPrice = shared.PricePackagePricePriceTypeFixedPrice

// This is an alias to an internal value.
const PricePackagePricePriceTypeCompositePrice = shared.PricePackagePricePriceTypeCompositePrice

// The LicenseType resource represents a type of license that can be assigned to
// users. License types are used during billing by grouping metrics on the
// configured grouping key.
//
// This is an alias to an internal type.
type PricePackagePriceLicenseType = shared.PricePackagePriceLicenseType

// This is an alias to an internal type.
type PriceMatrixPrice = shared.PriceMatrixPrice

// This is an alias to an internal type.
type PriceMatrixPriceBillingMode = shared.PriceMatrixPriceBillingMode

// This is an alias to an internal value.
const PriceMatrixPriceBillingModeInAdvance = shared.PriceMatrixPriceBillingModeInAdvance

// This is an alias to an internal value.
const PriceMatrixPriceBillingModeInArrear = shared.PriceMatrixPriceBillingModeInArrear

// This is an alias to an internal type.
type PriceMatrixPriceCadence = shared.PriceMatrixPriceCadence

// This is an alias to an internal value.
const PriceMatrixPriceCadenceOneTime = shared.PriceMatrixPriceCadenceOneTime

// This is an alias to an internal value.
const PriceMatrixPriceCadenceMonthly = shared.PriceMatrixPriceCadenceMonthly

// This is an alias to an internal value.
const PriceMatrixPriceCadenceQuarterly = shared.PriceMatrixPriceCadenceQuarterly

// This is an alias to an internal value.
const PriceMatrixPriceCadenceSemiAnnual = shared.PriceMatrixPriceCadenceSemiAnnual

// This is an alias to an internal value.
const PriceMatrixPriceCadenceAnnual = shared.PriceMatrixPriceCadenceAnnual

// This is an alias to an internal value.
const PriceMatrixPriceCadenceCustom = shared.PriceMatrixPriceCadenceCustom

// This is an alias to an internal type.
type PriceMatrixPriceCompositePriceFilter = shared.PriceMatrixPriceCompositePriceFilter

// The property of the price to filter on.
//
// This is an alias to an internal type.
type PriceMatrixPriceCompositePriceFiltersField = shared.PriceMatrixPriceCompositePriceFiltersField

// This is an alias to an internal value.
const PriceMatrixPriceCompositePriceFiltersFieldPriceID = shared.PriceMatrixPriceCompositePriceFiltersFieldPriceID

// This is an alias to an internal value.
const PriceMatrixPriceCompositePriceFiltersFieldItemID = shared.PriceMatrixPriceCompositePriceFiltersFieldItemID

// This is an alias to an internal value.
const PriceMatrixPriceCompositePriceFiltersFieldPriceType = shared.PriceMatrixPriceCompositePriceFiltersFieldPriceType

// This is an alias to an internal value.
const PriceMatrixPriceCompositePriceFiltersFieldCurrency = shared.PriceMatrixPriceCompositePriceFiltersFieldCurrency

// This is an alias to an internal value.
const PriceMatrixPriceCompositePriceFiltersFieldPricingUnitID = shared.PriceMatrixPriceCompositePriceFiltersFieldPricingUnitID

// Should prices that match the filter be included or excluded.
//
// This is an alias to an internal type.
type PriceMatrixPriceCompositePriceFiltersOperator = shared.PriceMatrixPriceCompositePriceFiltersOperator

// This is an alias to an internal value.
const PriceMatrixPriceCompositePriceFiltersOperatorIncludes = shared.PriceMatrixPriceCompositePriceFiltersOperatorIncludes

// This is an alias to an internal value.
const PriceMatrixPriceCompositePriceFiltersOperatorExcludes = shared.PriceMatrixPriceCompositePriceFiltersOperatorExcludes

// This is an alias to an internal type.
type PriceMatrixPriceConversionRateConfig = shared.PriceMatrixPriceConversionRateConfig

// This is an alias to an internal type.
type PriceMatrixPriceConversionRateConfigConversionRateType = shared.PriceMatrixPriceConversionRateConfigConversionRateType

// This is an alias to an internal value.
const PriceMatrixPriceConversionRateConfigConversionRateTypeUnit = shared.PriceMatrixPriceConversionRateConfigConversionRateTypeUnit

// This is an alias to an internal value.
const PriceMatrixPriceConversionRateConfigConversionRateTypeTiered = shared.PriceMatrixPriceConversionRateConfigConversionRateTypeTiered

// The pricing model type
//
// This is an alias to an internal type.
type PriceMatrixPriceModelType = shared.PriceMatrixPriceModelType

// This is an alias to an internal value.
const PriceMatrixPriceModelTypeMatrix = shared.PriceMatrixPriceModelTypeMatrix

// This is an alias to an internal type.
type PriceMatrixPricePriceType = shared.PriceMatrixPricePriceType

// This is an alias to an internal value.
const PriceMatrixPricePriceTypeUsagePrice = shared.PriceMatrixPricePriceTypeUsagePrice

// This is an alias to an internal value.
const PriceMatrixPricePriceTypeFixedPrice = shared.PriceMatrixPricePriceTypeFixedPrice

// This is an alias to an internal value.
const PriceMatrixPricePriceTypeCompositePrice = shared.PriceMatrixPricePriceTypeCompositePrice

// The LicenseType resource represents a type of license that can be assigned to
// users. License types are used during billing by grouping metrics on the
// configured grouping key.
//
// This is an alias to an internal type.
type PriceMatrixPriceLicenseType = shared.PriceMatrixPriceLicenseType

// This is an alias to an internal type.
type PriceThresholdTotalAmountPrice = shared.PriceThresholdTotalAmountPrice

// This is an alias to an internal type.
type PriceThresholdTotalAmountPriceBillingMode = shared.PriceThresholdTotalAmountPriceBillingMode

// This is an alias to an internal value.
const PriceThresholdTotalAmountPriceBillingModeInAdvance = shared.PriceThresholdTotalAmountPriceBillingModeInAdvance

// This is an alias to an internal value.
const PriceThresholdTotalAmountPriceBillingModeInArrear = shared.PriceThresholdTotalAmountPriceBillingModeInArrear

// This is an alias to an internal type.
type PriceThresholdTotalAmountPriceCadence = shared.PriceThresholdTotalAmountPriceCadence

// This is an alias to an internal value.
const PriceThresholdTotalAmountPriceCadenceOneTime = shared.PriceThresholdTotalAmountPriceCadenceOneTime

// This is an alias to an internal value.
const PriceThresholdTotalAmountPriceCadenceMonthly = shared.PriceThresholdTotalAmountPriceCadenceMonthly

// This is an alias to an internal value.
const PriceThresholdTotalAmountPriceCadenceQuarterly = shared.PriceThresholdTotalAmountPriceCadenceQuarterly

// This is an alias to an internal value.
const PriceThresholdTotalAmountPriceCadenceSemiAnnual = shared.PriceThresholdTotalAmountPriceCadenceSemiAnnual

// This is an alias to an internal value.
const PriceThresholdTotalAmountPriceCadenceAnnual = shared.PriceThresholdTotalAmountPriceCadenceAnnual

// This is an alias to an internal value.
const PriceThresholdTotalAmountPriceCadenceCustom = shared.PriceThresholdTotalAmountPriceCadenceCustom

// This is an alias to an internal type.
type PriceThresholdTotalAmountPriceCompositePriceFilter = shared.PriceThresholdTotalAmountPriceCompositePriceFilter

// The property of the price to filter on.
//
// This is an alias to an internal type.
type PriceThresholdTotalAmountPriceCompositePriceFiltersField = shared.PriceThresholdTotalAmountPriceCompositePriceFiltersField

// This is an alias to an internal value.
const PriceThresholdTotalAmountPriceCompositePriceFiltersFieldPriceID = shared.PriceThresholdTotalAmountPriceCompositePriceFiltersFieldPriceID

// This is an alias to an internal value.
const PriceThresholdTotalAmountPriceCompositePriceFiltersFieldItemID = shared.PriceThresholdTotalAmountPriceCompositePriceFiltersFieldItemID

// This is an alias to an internal value.
const PriceThresholdTotalAmountPriceCompositePriceFiltersFieldPriceType = shared.PriceThresholdTotalAmountPriceCompositePriceFiltersFieldPriceType

// This is an alias to an internal value.
const PriceThresholdTotalAmountPriceCompositePriceFiltersFieldCurrency = shared.PriceThresholdTotalAmountPriceCompositePriceFiltersFieldCurrency

// This is an alias to an internal value.
const PriceThresholdTotalAmountPriceCompositePriceFiltersFieldPricingUnitID = shared.PriceThresholdTotalAmountPriceCompositePriceFiltersFieldPricingUnitID

// Should prices that match the filter be included or excluded.
//
// This is an alias to an internal type.
type PriceThresholdTotalAmountPriceCompositePriceFiltersOperator = shared.PriceThresholdTotalAmountPriceCompositePriceFiltersOperator

// This is an alias to an internal value.
const PriceThresholdTotalAmountPriceCompositePriceFiltersOperatorIncludes = shared.PriceThresholdTotalAmountPriceCompositePriceFiltersOperatorIncludes

// This is an alias to an internal value.
const PriceThresholdTotalAmountPriceCompositePriceFiltersOperatorExcludes = shared.PriceThresholdTotalAmountPriceCompositePriceFiltersOperatorExcludes

// This is an alias to an internal type.
type PriceThresholdTotalAmountPriceConversionRateConfig = shared.PriceThresholdTotalAmountPriceConversionRateConfig

// This is an alias to an internal type.
type PriceThresholdTotalAmountPriceConversionRateConfigConversionRateType = shared.PriceThresholdTotalAmountPriceConversionRateConfigConversionRateType

// This is an alias to an internal value.
const PriceThresholdTotalAmountPriceConversionRateConfigConversionRateTypeUnit = shared.PriceThresholdTotalAmountPriceConversionRateConfigConversionRateTypeUnit

// This is an alias to an internal value.
const PriceThresholdTotalAmountPriceConversionRateConfigConversionRateTypeTiered = shared.PriceThresholdTotalAmountPriceConversionRateConfigConversionRateTypeTiered

// The pricing model type
//
// This is an alias to an internal type.
type PriceThresholdTotalAmountPriceModelType = shared.PriceThresholdTotalAmountPriceModelType

// This is an alias to an internal value.
const PriceThresholdTotalAmountPriceModelTypeThresholdTotalAmount = shared.PriceThresholdTotalAmountPriceModelTypeThresholdTotalAmount

// This is an alias to an internal type.
type PriceThresholdTotalAmountPricePriceType = shared.PriceThresholdTotalAmountPricePriceType

// This is an alias to an internal value.
const PriceThresholdTotalAmountPricePriceTypeUsagePrice = shared.PriceThresholdTotalAmountPricePriceTypeUsagePrice

// This is an alias to an internal value.
const PriceThresholdTotalAmountPricePriceTypeFixedPrice = shared.PriceThresholdTotalAmountPricePriceTypeFixedPrice

// This is an alias to an internal value.
const PriceThresholdTotalAmountPricePriceTypeCompositePrice = shared.PriceThresholdTotalAmountPricePriceTypeCompositePrice

// Configuration for threshold_total_amount pricing
//
// This is an alias to an internal type.
type PriceThresholdTotalAmountPriceThresholdTotalAmountConfig = shared.PriceThresholdTotalAmountPriceThresholdTotalAmountConfig

// Configuration for a single threshold
//
// This is an alias to an internal type.
type PriceThresholdTotalAmountPriceThresholdTotalAmountConfigConsumptionTable = shared.PriceThresholdTotalAmountPriceThresholdTotalAmountConfigConsumptionTable

// The LicenseType resource represents a type of license that can be assigned to
// users. License types are used during billing by grouping metrics on the
// configured grouping key.
//
// This is an alias to an internal type.
type PriceThresholdTotalAmountPriceLicenseType = shared.PriceThresholdTotalAmountPriceLicenseType

// This is an alias to an internal type.
type PriceTieredPackagePrice = shared.PriceTieredPackagePrice

// This is an alias to an internal type.
type PriceTieredPackagePriceBillingMode = shared.PriceTieredPackagePriceBillingMode

// This is an alias to an internal value.
const PriceTieredPackagePriceBillingModeInAdvance = shared.PriceTieredPackagePriceBillingModeInAdvance

// This is an alias to an internal value.
const PriceTieredPackagePriceBillingModeInArrear = shared.PriceTieredPackagePriceBillingModeInArrear

// This is an alias to an internal type.
type PriceTieredPackagePriceCadence = shared.PriceTieredPackagePriceCadence

// This is an alias to an internal value.
const PriceTieredPackagePriceCadenceOneTime = shared.PriceTieredPackagePriceCadenceOneTime

// This is an alias to an internal value.
const PriceTieredPackagePriceCadenceMonthly = shared.PriceTieredPackagePriceCadenceMonthly

// This is an alias to an internal value.
const PriceTieredPackagePriceCadenceQuarterly = shared.PriceTieredPackagePriceCadenceQuarterly

// This is an alias to an internal value.
const PriceTieredPackagePriceCadenceSemiAnnual = shared.PriceTieredPackagePriceCadenceSemiAnnual

// This is an alias to an internal value.
const PriceTieredPackagePriceCadenceAnnual = shared.PriceTieredPackagePriceCadenceAnnual

// This is an alias to an internal value.
const PriceTieredPackagePriceCadenceCustom = shared.PriceTieredPackagePriceCadenceCustom

// This is an alias to an internal type.
type PriceTieredPackagePriceCompositePriceFilter = shared.PriceTieredPackagePriceCompositePriceFilter

// The property of the price to filter on.
//
// This is an alias to an internal type.
type PriceTieredPackagePriceCompositePriceFiltersField = shared.PriceTieredPackagePriceCompositePriceFiltersField

// This is an alias to an internal value.
const PriceTieredPackagePriceCompositePriceFiltersFieldPriceID = shared.PriceTieredPackagePriceCompositePriceFiltersFieldPriceID

// This is an alias to an internal value.
const PriceTieredPackagePriceCompositePriceFiltersFieldItemID = shared.PriceTieredPackagePriceCompositePriceFiltersFieldItemID

// This is an alias to an internal value.
const PriceTieredPackagePriceCompositePriceFiltersFieldPriceType = shared.PriceTieredPackagePriceCompositePriceFiltersFieldPriceType

// This is an alias to an internal value.
const PriceTieredPackagePriceCompositePriceFiltersFieldCurrency = shared.PriceTieredPackagePriceCompositePriceFiltersFieldCurrency

// This is an alias to an internal value.
const PriceTieredPackagePriceCompositePriceFiltersFieldPricingUnitID = shared.PriceTieredPackagePriceCompositePriceFiltersFieldPricingUnitID

// Should prices that match the filter be included or excluded.
//
// This is an alias to an internal type.
type PriceTieredPackagePriceCompositePriceFiltersOperator = shared.PriceTieredPackagePriceCompositePriceFiltersOperator

// This is an alias to an internal value.
const PriceTieredPackagePriceCompositePriceFiltersOperatorIncludes = shared.PriceTieredPackagePriceCompositePriceFiltersOperatorIncludes

// This is an alias to an internal value.
const PriceTieredPackagePriceCompositePriceFiltersOperatorExcludes = shared.PriceTieredPackagePriceCompositePriceFiltersOperatorExcludes

// This is an alias to an internal type.
type PriceTieredPackagePriceConversionRateConfig = shared.PriceTieredPackagePriceConversionRateConfig

// This is an alias to an internal type.
type PriceTieredPackagePriceConversionRateConfigConversionRateType = shared.PriceTieredPackagePriceConversionRateConfigConversionRateType

// This is an alias to an internal value.
const PriceTieredPackagePriceConversionRateConfigConversionRateTypeUnit = shared.PriceTieredPackagePriceConversionRateConfigConversionRateTypeUnit

// This is an alias to an internal value.
const PriceTieredPackagePriceConversionRateConfigConversionRateTypeTiered = shared.PriceTieredPackagePriceConversionRateConfigConversionRateTypeTiered

// The pricing model type
//
// This is an alias to an internal type.
type PriceTieredPackagePriceModelType = shared.PriceTieredPackagePriceModelType

// This is an alias to an internal value.
const PriceTieredPackagePriceModelTypeTieredPackage = shared.PriceTieredPackagePriceModelTypeTieredPackage

// This is an alias to an internal type.
type PriceTieredPackagePricePriceType = shared.PriceTieredPackagePricePriceType

// This is an alias to an internal value.
const PriceTieredPackagePricePriceTypeUsagePrice = shared.PriceTieredPackagePricePriceTypeUsagePrice

// This is an alias to an internal value.
const PriceTieredPackagePricePriceTypeFixedPrice = shared.PriceTieredPackagePricePriceTypeFixedPrice

// This is an alias to an internal value.
const PriceTieredPackagePricePriceTypeCompositePrice = shared.PriceTieredPackagePricePriceTypeCompositePrice

// Configuration for tiered_package pricing
//
// This is an alias to an internal type.
type PriceTieredPackagePriceTieredPackageConfig = shared.PriceTieredPackagePriceTieredPackageConfig

// Configuration for a single tier with business logic
//
// This is an alias to an internal type.
type PriceTieredPackagePriceTieredPackageConfigTier = shared.PriceTieredPackagePriceTieredPackageConfigTier

// The LicenseType resource represents a type of license that can be assigned to
// users. License types are used during billing by grouping metrics on the
// configured grouping key.
//
// This is an alias to an internal type.
type PriceTieredPackagePriceLicenseType = shared.PriceTieredPackagePriceLicenseType

// This is an alias to an internal type.
type PriceTieredWithMinimumPrice = shared.PriceTieredWithMinimumPrice

// This is an alias to an internal type.
type PriceTieredWithMinimumPriceBillingMode = shared.PriceTieredWithMinimumPriceBillingMode

// This is an alias to an internal value.
const PriceTieredWithMinimumPriceBillingModeInAdvance = shared.PriceTieredWithMinimumPriceBillingModeInAdvance

// This is an alias to an internal value.
const PriceTieredWithMinimumPriceBillingModeInArrear = shared.PriceTieredWithMinimumPriceBillingModeInArrear

// This is an alias to an internal type.
type PriceTieredWithMinimumPriceCadence = shared.PriceTieredWithMinimumPriceCadence

// This is an alias to an internal value.
const PriceTieredWithMinimumPriceCadenceOneTime = shared.PriceTieredWithMinimumPriceCadenceOneTime

// This is an alias to an internal value.
const PriceTieredWithMinimumPriceCadenceMonthly = shared.PriceTieredWithMinimumPriceCadenceMonthly

// This is an alias to an internal value.
const PriceTieredWithMinimumPriceCadenceQuarterly = shared.PriceTieredWithMinimumPriceCadenceQuarterly

// This is an alias to an internal value.
const PriceTieredWithMinimumPriceCadenceSemiAnnual = shared.PriceTieredWithMinimumPriceCadenceSemiAnnual

// This is an alias to an internal value.
const PriceTieredWithMinimumPriceCadenceAnnual = shared.PriceTieredWithMinimumPriceCadenceAnnual

// This is an alias to an internal value.
const PriceTieredWithMinimumPriceCadenceCustom = shared.PriceTieredWithMinimumPriceCadenceCustom

// This is an alias to an internal type.
type PriceTieredWithMinimumPriceCompositePriceFilter = shared.PriceTieredWithMinimumPriceCompositePriceFilter

// The property of the price to filter on.
//
// This is an alias to an internal type.
type PriceTieredWithMinimumPriceCompositePriceFiltersField = shared.PriceTieredWithMinimumPriceCompositePriceFiltersField

// This is an alias to an internal value.
const PriceTieredWithMinimumPriceCompositePriceFiltersFieldPriceID = shared.PriceTieredWithMinimumPriceCompositePriceFiltersFieldPriceID

// This is an alias to an internal value.
const PriceTieredWithMinimumPriceCompositePriceFiltersFieldItemID = shared.PriceTieredWithMinimumPriceCompositePriceFiltersFieldItemID

// This is an alias to an internal value.
const PriceTieredWithMinimumPriceCompositePriceFiltersFieldPriceType = shared.PriceTieredWithMinimumPriceCompositePriceFiltersFieldPriceType

// This is an alias to an internal value.
const PriceTieredWithMinimumPriceCompositePriceFiltersFieldCurrency = shared.PriceTieredWithMinimumPriceCompositePriceFiltersFieldCurrency

// This is an alias to an internal value.
const PriceTieredWithMinimumPriceCompositePriceFiltersFieldPricingUnitID = shared.PriceTieredWithMinimumPriceCompositePriceFiltersFieldPricingUnitID

// Should prices that match the filter be included or excluded.
//
// This is an alias to an internal type.
type PriceTieredWithMinimumPriceCompositePriceFiltersOperator = shared.PriceTieredWithMinimumPriceCompositePriceFiltersOperator

// This is an alias to an internal value.
const PriceTieredWithMinimumPriceCompositePriceFiltersOperatorIncludes = shared.PriceTieredWithMinimumPriceCompositePriceFiltersOperatorIncludes

// This is an alias to an internal value.
const PriceTieredWithMinimumPriceCompositePriceFiltersOperatorExcludes = shared.PriceTieredWithMinimumPriceCompositePriceFiltersOperatorExcludes

// This is an alias to an internal type.
type PriceTieredWithMinimumPriceConversionRateConfig = shared.PriceTieredWithMinimumPriceConversionRateConfig

// This is an alias to an internal type.
type PriceTieredWithMinimumPriceConversionRateConfigConversionRateType = shared.PriceTieredWithMinimumPriceConversionRateConfigConversionRateType

// This is an alias to an internal value.
const PriceTieredWithMinimumPriceConversionRateConfigConversionRateTypeUnit = shared.PriceTieredWithMinimumPriceConversionRateConfigConversionRateTypeUnit

// This is an alias to an internal value.
const PriceTieredWithMinimumPriceConversionRateConfigConversionRateTypeTiered = shared.PriceTieredWithMinimumPriceConversionRateConfigConversionRateTypeTiered

// The pricing model type
//
// This is an alias to an internal type.
type PriceTieredWithMinimumPriceModelType = shared.PriceTieredWithMinimumPriceModelType

// This is an alias to an internal value.
const PriceTieredWithMinimumPriceModelTypeTieredWithMinimum = shared.PriceTieredWithMinimumPriceModelTypeTieredWithMinimum

// This is an alias to an internal type.
type PriceTieredWithMinimumPricePriceType = shared.PriceTieredWithMinimumPricePriceType

// This is an alias to an internal value.
const PriceTieredWithMinimumPricePriceTypeUsagePrice = shared.PriceTieredWithMinimumPricePriceTypeUsagePrice

// This is an alias to an internal value.
const PriceTieredWithMinimumPricePriceTypeFixedPrice = shared.PriceTieredWithMinimumPricePriceTypeFixedPrice

// This is an alias to an internal value.
const PriceTieredWithMinimumPricePriceTypeCompositePrice = shared.PriceTieredWithMinimumPricePriceTypeCompositePrice

// Configuration for tiered_with_minimum pricing
//
// This is an alias to an internal type.
type PriceTieredWithMinimumPriceTieredWithMinimumConfig = shared.PriceTieredWithMinimumPriceTieredWithMinimumConfig

// Configuration for a single tier
//
// This is an alias to an internal type.
type PriceTieredWithMinimumPriceTieredWithMinimumConfigTier = shared.PriceTieredWithMinimumPriceTieredWithMinimumConfigTier

// The LicenseType resource represents a type of license that can be assigned to
// users. License types are used during billing by grouping metrics on the
// configured grouping key.
//
// This is an alias to an internal type.
type PriceTieredWithMinimumPriceLicenseType = shared.PriceTieredWithMinimumPriceLicenseType

// This is an alias to an internal type.
type PriceGroupedTieredPrice = shared.PriceGroupedTieredPrice

// This is an alias to an internal type.
type PriceGroupedTieredPriceBillingMode = shared.PriceGroupedTieredPriceBillingMode

// This is an alias to an internal value.
const PriceGroupedTieredPriceBillingModeInAdvance = shared.PriceGroupedTieredPriceBillingModeInAdvance

// This is an alias to an internal value.
const PriceGroupedTieredPriceBillingModeInArrear = shared.PriceGroupedTieredPriceBillingModeInArrear

// This is an alias to an internal type.
type PriceGroupedTieredPriceCadence = shared.PriceGroupedTieredPriceCadence

// This is an alias to an internal value.
const PriceGroupedTieredPriceCadenceOneTime = shared.PriceGroupedTieredPriceCadenceOneTime

// This is an alias to an internal value.
const PriceGroupedTieredPriceCadenceMonthly = shared.PriceGroupedTieredPriceCadenceMonthly

// This is an alias to an internal value.
const PriceGroupedTieredPriceCadenceQuarterly = shared.PriceGroupedTieredPriceCadenceQuarterly

// This is an alias to an internal value.
const PriceGroupedTieredPriceCadenceSemiAnnual = shared.PriceGroupedTieredPriceCadenceSemiAnnual

// This is an alias to an internal value.
const PriceGroupedTieredPriceCadenceAnnual = shared.PriceGroupedTieredPriceCadenceAnnual

// This is an alias to an internal value.
const PriceGroupedTieredPriceCadenceCustom = shared.PriceGroupedTieredPriceCadenceCustom

// This is an alias to an internal type.
type PriceGroupedTieredPriceCompositePriceFilter = shared.PriceGroupedTieredPriceCompositePriceFilter

// The property of the price to filter on.
//
// This is an alias to an internal type.
type PriceGroupedTieredPriceCompositePriceFiltersField = shared.PriceGroupedTieredPriceCompositePriceFiltersField

// This is an alias to an internal value.
const PriceGroupedTieredPriceCompositePriceFiltersFieldPriceID = shared.PriceGroupedTieredPriceCompositePriceFiltersFieldPriceID

// This is an alias to an internal value.
const PriceGroupedTieredPriceCompositePriceFiltersFieldItemID = shared.PriceGroupedTieredPriceCompositePriceFiltersFieldItemID

// This is an alias to an internal value.
const PriceGroupedTieredPriceCompositePriceFiltersFieldPriceType = shared.PriceGroupedTieredPriceCompositePriceFiltersFieldPriceType

// This is an alias to an internal value.
const PriceGroupedTieredPriceCompositePriceFiltersFieldCurrency = shared.PriceGroupedTieredPriceCompositePriceFiltersFieldCurrency

// This is an alias to an internal value.
const PriceGroupedTieredPriceCompositePriceFiltersFieldPricingUnitID = shared.PriceGroupedTieredPriceCompositePriceFiltersFieldPricingUnitID

// Should prices that match the filter be included or excluded.
//
// This is an alias to an internal type.
type PriceGroupedTieredPriceCompositePriceFiltersOperator = shared.PriceGroupedTieredPriceCompositePriceFiltersOperator

// This is an alias to an internal value.
const PriceGroupedTieredPriceCompositePriceFiltersOperatorIncludes = shared.PriceGroupedTieredPriceCompositePriceFiltersOperatorIncludes

// This is an alias to an internal value.
const PriceGroupedTieredPriceCompositePriceFiltersOperatorExcludes = shared.PriceGroupedTieredPriceCompositePriceFiltersOperatorExcludes

// This is an alias to an internal type.
type PriceGroupedTieredPriceConversionRateConfig = shared.PriceGroupedTieredPriceConversionRateConfig

// This is an alias to an internal type.
type PriceGroupedTieredPriceConversionRateConfigConversionRateType = shared.PriceGroupedTieredPriceConversionRateConfigConversionRateType

// This is an alias to an internal value.
const PriceGroupedTieredPriceConversionRateConfigConversionRateTypeUnit = shared.PriceGroupedTieredPriceConversionRateConfigConversionRateTypeUnit

// This is an alias to an internal value.
const PriceGroupedTieredPriceConversionRateConfigConversionRateTypeTiered = shared.PriceGroupedTieredPriceConversionRateConfigConversionRateTypeTiered

// Configuration for grouped_tiered pricing
//
// This is an alias to an internal type.
type PriceGroupedTieredPriceGroupedTieredConfig = shared.PriceGroupedTieredPriceGroupedTieredConfig

// Configuration for a single tier
//
// This is an alias to an internal type.
type PriceGroupedTieredPriceGroupedTieredConfigTier = shared.PriceGroupedTieredPriceGroupedTieredConfigTier

// The pricing model type
//
// This is an alias to an internal type.
type PriceGroupedTieredPriceModelType = shared.PriceGroupedTieredPriceModelType

// This is an alias to an internal value.
const PriceGroupedTieredPriceModelTypeGroupedTiered = shared.PriceGroupedTieredPriceModelTypeGroupedTiered

// This is an alias to an internal type.
type PriceGroupedTieredPricePriceType = shared.PriceGroupedTieredPricePriceType

// This is an alias to an internal value.
const PriceGroupedTieredPricePriceTypeUsagePrice = shared.PriceGroupedTieredPricePriceTypeUsagePrice

// This is an alias to an internal value.
const PriceGroupedTieredPricePriceTypeFixedPrice = shared.PriceGroupedTieredPricePriceTypeFixedPrice

// This is an alias to an internal value.
const PriceGroupedTieredPricePriceTypeCompositePrice = shared.PriceGroupedTieredPricePriceTypeCompositePrice

// The LicenseType resource represents a type of license that can be assigned to
// users. License types are used during billing by grouping metrics on the
// configured grouping key.
//
// This is an alias to an internal type.
type PriceGroupedTieredPriceLicenseType = shared.PriceGroupedTieredPriceLicenseType

// This is an alias to an internal type.
type PriceTieredPackageWithMinimumPrice = shared.PriceTieredPackageWithMinimumPrice

// This is an alias to an internal type.
type PriceTieredPackageWithMinimumPriceBillingMode = shared.PriceTieredPackageWithMinimumPriceBillingMode

// This is an alias to an internal value.
const PriceTieredPackageWithMinimumPriceBillingModeInAdvance = shared.PriceTieredPackageWithMinimumPriceBillingModeInAdvance

// This is an alias to an internal value.
const PriceTieredPackageWithMinimumPriceBillingModeInArrear = shared.PriceTieredPackageWithMinimumPriceBillingModeInArrear

// This is an alias to an internal type.
type PriceTieredPackageWithMinimumPriceCadence = shared.PriceTieredPackageWithMinimumPriceCadence

// This is an alias to an internal value.
const PriceTieredPackageWithMinimumPriceCadenceOneTime = shared.PriceTieredPackageWithMinimumPriceCadenceOneTime

// This is an alias to an internal value.
const PriceTieredPackageWithMinimumPriceCadenceMonthly = shared.PriceTieredPackageWithMinimumPriceCadenceMonthly

// This is an alias to an internal value.
const PriceTieredPackageWithMinimumPriceCadenceQuarterly = shared.PriceTieredPackageWithMinimumPriceCadenceQuarterly

// This is an alias to an internal value.
const PriceTieredPackageWithMinimumPriceCadenceSemiAnnual = shared.PriceTieredPackageWithMinimumPriceCadenceSemiAnnual

// This is an alias to an internal value.
const PriceTieredPackageWithMinimumPriceCadenceAnnual = shared.PriceTieredPackageWithMinimumPriceCadenceAnnual

// This is an alias to an internal value.
const PriceTieredPackageWithMinimumPriceCadenceCustom = shared.PriceTieredPackageWithMinimumPriceCadenceCustom

// This is an alias to an internal type.
type PriceTieredPackageWithMinimumPriceCompositePriceFilter = shared.PriceTieredPackageWithMinimumPriceCompositePriceFilter

// The property of the price to filter on.
//
// This is an alias to an internal type.
type PriceTieredPackageWithMinimumPriceCompositePriceFiltersField = shared.PriceTieredPackageWithMinimumPriceCompositePriceFiltersField

// This is an alias to an internal value.
const PriceTieredPackageWithMinimumPriceCompositePriceFiltersFieldPriceID = shared.PriceTieredPackageWithMinimumPriceCompositePriceFiltersFieldPriceID

// This is an alias to an internal value.
const PriceTieredPackageWithMinimumPriceCompositePriceFiltersFieldItemID = shared.PriceTieredPackageWithMinimumPriceCompositePriceFiltersFieldItemID

// This is an alias to an internal value.
const PriceTieredPackageWithMinimumPriceCompositePriceFiltersFieldPriceType = shared.PriceTieredPackageWithMinimumPriceCompositePriceFiltersFieldPriceType

// This is an alias to an internal value.
const PriceTieredPackageWithMinimumPriceCompositePriceFiltersFieldCurrency = shared.PriceTieredPackageWithMinimumPriceCompositePriceFiltersFieldCurrency

// This is an alias to an internal value.
const PriceTieredPackageWithMinimumPriceCompositePriceFiltersFieldPricingUnitID = shared.PriceTieredPackageWithMinimumPriceCompositePriceFiltersFieldPricingUnitID

// Should prices that match the filter be included or excluded.
//
// This is an alias to an internal type.
type PriceTieredPackageWithMinimumPriceCompositePriceFiltersOperator = shared.PriceTieredPackageWithMinimumPriceCompositePriceFiltersOperator

// This is an alias to an internal value.
const PriceTieredPackageWithMinimumPriceCompositePriceFiltersOperatorIncludes = shared.PriceTieredPackageWithMinimumPriceCompositePriceFiltersOperatorIncludes

// This is an alias to an internal value.
const PriceTieredPackageWithMinimumPriceCompositePriceFiltersOperatorExcludes = shared.PriceTieredPackageWithMinimumPriceCompositePriceFiltersOperatorExcludes

// This is an alias to an internal type.
type PriceTieredPackageWithMinimumPriceConversionRateConfig = shared.PriceTieredPackageWithMinimumPriceConversionRateConfig

// This is an alias to an internal type.
type PriceTieredPackageWithMinimumPriceConversionRateConfigConversionRateType = shared.PriceTieredPackageWithMinimumPriceConversionRateConfigConversionRateType

// This is an alias to an internal value.
const PriceTieredPackageWithMinimumPriceConversionRateConfigConversionRateTypeUnit = shared.PriceTieredPackageWithMinimumPriceConversionRateConfigConversionRateTypeUnit

// This is an alias to an internal value.
const PriceTieredPackageWithMinimumPriceConversionRateConfigConversionRateTypeTiered = shared.PriceTieredPackageWithMinimumPriceConversionRateConfigConversionRateTypeTiered

// The pricing model type
//
// This is an alias to an internal type.
type PriceTieredPackageWithMinimumPriceModelType = shared.PriceTieredPackageWithMinimumPriceModelType

// This is an alias to an internal value.
const PriceTieredPackageWithMinimumPriceModelTypeTieredPackageWithMinimum = shared.PriceTieredPackageWithMinimumPriceModelTypeTieredPackageWithMinimum

// This is an alias to an internal type.
type PriceTieredPackageWithMinimumPricePriceType = shared.PriceTieredPackageWithMinimumPricePriceType

// This is an alias to an internal value.
const PriceTieredPackageWithMinimumPricePriceTypeUsagePrice = shared.PriceTieredPackageWithMinimumPricePriceTypeUsagePrice

// This is an alias to an internal value.
const PriceTieredPackageWithMinimumPricePriceTypeFixedPrice = shared.PriceTieredPackageWithMinimumPricePriceTypeFixedPrice

// This is an alias to an internal value.
const PriceTieredPackageWithMinimumPricePriceTypeCompositePrice = shared.PriceTieredPackageWithMinimumPricePriceTypeCompositePrice

// Configuration for tiered_package_with_minimum pricing
//
// This is an alias to an internal type.
type PriceTieredPackageWithMinimumPriceTieredPackageWithMinimumConfig = shared.PriceTieredPackageWithMinimumPriceTieredPackageWithMinimumConfig

// Configuration for a single tier
//
// This is an alias to an internal type.
type PriceTieredPackageWithMinimumPriceTieredPackageWithMinimumConfigTier = shared.PriceTieredPackageWithMinimumPriceTieredPackageWithMinimumConfigTier

// The LicenseType resource represents a type of license that can be assigned to
// users. License types are used during billing by grouping metrics on the
// configured grouping key.
//
// This is an alias to an internal type.
type PriceTieredPackageWithMinimumPriceLicenseType = shared.PriceTieredPackageWithMinimumPriceLicenseType

// This is an alias to an internal type.
type PricePackageWithAllocationPrice = shared.PricePackageWithAllocationPrice

// This is an alias to an internal type.
type PricePackageWithAllocationPriceBillingMode = shared.PricePackageWithAllocationPriceBillingMode

// This is an alias to an internal value.
const PricePackageWithAllocationPriceBillingModeInAdvance = shared.PricePackageWithAllocationPriceBillingModeInAdvance

// This is an alias to an internal value.
const PricePackageWithAllocationPriceBillingModeInArrear = shared.PricePackageWithAllocationPriceBillingModeInArrear

// This is an alias to an internal type.
type PricePackageWithAllocationPriceCadence = shared.PricePackageWithAllocationPriceCadence

// This is an alias to an internal value.
const PricePackageWithAllocationPriceCadenceOneTime = shared.PricePackageWithAllocationPriceCadenceOneTime

// This is an alias to an internal value.
const PricePackageWithAllocationPriceCadenceMonthly = shared.PricePackageWithAllocationPriceCadenceMonthly

// This is an alias to an internal value.
const PricePackageWithAllocationPriceCadenceQuarterly = shared.PricePackageWithAllocationPriceCadenceQuarterly

// This is an alias to an internal value.
const PricePackageWithAllocationPriceCadenceSemiAnnual = shared.PricePackageWithAllocationPriceCadenceSemiAnnual

// This is an alias to an internal value.
const PricePackageWithAllocationPriceCadenceAnnual = shared.PricePackageWithAllocationPriceCadenceAnnual

// This is an alias to an internal value.
const PricePackageWithAllocationPriceCadenceCustom = shared.PricePackageWithAllocationPriceCadenceCustom

// This is an alias to an internal type.
type PricePackageWithAllocationPriceCompositePriceFilter = shared.PricePackageWithAllocationPriceCompositePriceFilter

// The property of the price to filter on.
//
// This is an alias to an internal type.
type PricePackageWithAllocationPriceCompositePriceFiltersField = shared.PricePackageWithAllocationPriceCompositePriceFiltersField

// This is an alias to an internal value.
const PricePackageWithAllocationPriceCompositePriceFiltersFieldPriceID = shared.PricePackageWithAllocationPriceCompositePriceFiltersFieldPriceID

// This is an alias to an internal value.
const PricePackageWithAllocationPriceCompositePriceFiltersFieldItemID = shared.PricePackageWithAllocationPriceCompositePriceFiltersFieldItemID

// This is an alias to an internal value.
const PricePackageWithAllocationPriceCompositePriceFiltersFieldPriceType = shared.PricePackageWithAllocationPriceCompositePriceFiltersFieldPriceType

// This is an alias to an internal value.
const PricePackageWithAllocationPriceCompositePriceFiltersFieldCurrency = shared.PricePackageWithAllocationPriceCompositePriceFiltersFieldCurrency

// This is an alias to an internal value.
const PricePackageWithAllocationPriceCompositePriceFiltersFieldPricingUnitID = shared.PricePackageWithAllocationPriceCompositePriceFiltersFieldPricingUnitID

// Should prices that match the filter be included or excluded.
//
// This is an alias to an internal type.
type PricePackageWithAllocationPriceCompositePriceFiltersOperator = shared.PricePackageWithAllocationPriceCompositePriceFiltersOperator

// This is an alias to an internal value.
const PricePackageWithAllocationPriceCompositePriceFiltersOperatorIncludes = shared.PricePackageWithAllocationPriceCompositePriceFiltersOperatorIncludes

// This is an alias to an internal value.
const PricePackageWithAllocationPriceCompositePriceFiltersOperatorExcludes = shared.PricePackageWithAllocationPriceCompositePriceFiltersOperatorExcludes

// This is an alias to an internal type.
type PricePackageWithAllocationPriceConversionRateConfig = shared.PricePackageWithAllocationPriceConversionRateConfig

// This is an alias to an internal type.
type PricePackageWithAllocationPriceConversionRateConfigConversionRateType = shared.PricePackageWithAllocationPriceConversionRateConfigConversionRateType

// This is an alias to an internal value.
const PricePackageWithAllocationPriceConversionRateConfigConversionRateTypeUnit = shared.PricePackageWithAllocationPriceConversionRateConfigConversionRateTypeUnit

// This is an alias to an internal value.
const PricePackageWithAllocationPriceConversionRateConfigConversionRateTypeTiered = shared.PricePackageWithAllocationPriceConversionRateConfigConversionRateTypeTiered

// The pricing model type
//
// This is an alias to an internal type.
type PricePackageWithAllocationPriceModelType = shared.PricePackageWithAllocationPriceModelType

// This is an alias to an internal value.
const PricePackageWithAllocationPriceModelTypePackageWithAllocation = shared.PricePackageWithAllocationPriceModelTypePackageWithAllocation

// Configuration for package_with_allocation pricing
//
// This is an alias to an internal type.
type PricePackageWithAllocationPricePackageWithAllocationConfig = shared.PricePackageWithAllocationPricePackageWithAllocationConfig

// This is an alias to an internal type.
type PricePackageWithAllocationPricePriceType = shared.PricePackageWithAllocationPricePriceType

// This is an alias to an internal value.
const PricePackageWithAllocationPricePriceTypeUsagePrice = shared.PricePackageWithAllocationPricePriceTypeUsagePrice

// This is an alias to an internal value.
const PricePackageWithAllocationPricePriceTypeFixedPrice = shared.PricePackageWithAllocationPricePriceTypeFixedPrice

// This is an alias to an internal value.
const PricePackageWithAllocationPricePriceTypeCompositePrice = shared.PricePackageWithAllocationPricePriceTypeCompositePrice

// The LicenseType resource represents a type of license that can be assigned to
// users. License types are used during billing by grouping metrics on the
// configured grouping key.
//
// This is an alias to an internal type.
type PricePackageWithAllocationPriceLicenseType = shared.PricePackageWithAllocationPriceLicenseType

// This is an alias to an internal type.
type PriceUnitWithPercentPrice = shared.PriceUnitWithPercentPrice

// This is an alias to an internal type.
type PriceUnitWithPercentPriceBillingMode = shared.PriceUnitWithPercentPriceBillingMode

// This is an alias to an internal value.
const PriceUnitWithPercentPriceBillingModeInAdvance = shared.PriceUnitWithPercentPriceBillingModeInAdvance

// This is an alias to an internal value.
const PriceUnitWithPercentPriceBillingModeInArrear = shared.PriceUnitWithPercentPriceBillingModeInArrear

// This is an alias to an internal type.
type PriceUnitWithPercentPriceCadence = shared.PriceUnitWithPercentPriceCadence

// This is an alias to an internal value.
const PriceUnitWithPercentPriceCadenceOneTime = shared.PriceUnitWithPercentPriceCadenceOneTime

// This is an alias to an internal value.
const PriceUnitWithPercentPriceCadenceMonthly = shared.PriceUnitWithPercentPriceCadenceMonthly

// This is an alias to an internal value.
const PriceUnitWithPercentPriceCadenceQuarterly = shared.PriceUnitWithPercentPriceCadenceQuarterly

// This is an alias to an internal value.
const PriceUnitWithPercentPriceCadenceSemiAnnual = shared.PriceUnitWithPercentPriceCadenceSemiAnnual

// This is an alias to an internal value.
const PriceUnitWithPercentPriceCadenceAnnual = shared.PriceUnitWithPercentPriceCadenceAnnual

// This is an alias to an internal value.
const PriceUnitWithPercentPriceCadenceCustom = shared.PriceUnitWithPercentPriceCadenceCustom

// This is an alias to an internal type.
type PriceUnitWithPercentPriceCompositePriceFilter = shared.PriceUnitWithPercentPriceCompositePriceFilter

// The property of the price to filter on.
//
// This is an alias to an internal type.
type PriceUnitWithPercentPriceCompositePriceFiltersField = shared.PriceUnitWithPercentPriceCompositePriceFiltersField

// This is an alias to an internal value.
const PriceUnitWithPercentPriceCompositePriceFiltersFieldPriceID = shared.PriceUnitWithPercentPriceCompositePriceFiltersFieldPriceID

// This is an alias to an internal value.
const PriceUnitWithPercentPriceCompositePriceFiltersFieldItemID = shared.PriceUnitWithPercentPriceCompositePriceFiltersFieldItemID

// This is an alias to an internal value.
const PriceUnitWithPercentPriceCompositePriceFiltersFieldPriceType = shared.PriceUnitWithPercentPriceCompositePriceFiltersFieldPriceType

// This is an alias to an internal value.
const PriceUnitWithPercentPriceCompositePriceFiltersFieldCurrency = shared.PriceUnitWithPercentPriceCompositePriceFiltersFieldCurrency

// This is an alias to an internal value.
const PriceUnitWithPercentPriceCompositePriceFiltersFieldPricingUnitID = shared.PriceUnitWithPercentPriceCompositePriceFiltersFieldPricingUnitID

// Should prices that match the filter be included or excluded.
//
// This is an alias to an internal type.
type PriceUnitWithPercentPriceCompositePriceFiltersOperator = shared.PriceUnitWithPercentPriceCompositePriceFiltersOperator

// This is an alias to an internal value.
const PriceUnitWithPercentPriceCompositePriceFiltersOperatorIncludes = shared.PriceUnitWithPercentPriceCompositePriceFiltersOperatorIncludes

// This is an alias to an internal value.
const PriceUnitWithPercentPriceCompositePriceFiltersOperatorExcludes = shared.PriceUnitWithPercentPriceCompositePriceFiltersOperatorExcludes

// This is an alias to an internal type.
type PriceUnitWithPercentPriceConversionRateConfig = shared.PriceUnitWithPercentPriceConversionRateConfig

// This is an alias to an internal type.
type PriceUnitWithPercentPriceConversionRateConfigConversionRateType = shared.PriceUnitWithPercentPriceConversionRateConfigConversionRateType

// This is an alias to an internal value.
const PriceUnitWithPercentPriceConversionRateConfigConversionRateTypeUnit = shared.PriceUnitWithPercentPriceConversionRateConfigConversionRateTypeUnit

// This is an alias to an internal value.
const PriceUnitWithPercentPriceConversionRateConfigConversionRateTypeTiered = shared.PriceUnitWithPercentPriceConversionRateConfigConversionRateTypeTiered

// The pricing model type
//
// This is an alias to an internal type.
type PriceUnitWithPercentPriceModelType = shared.PriceUnitWithPercentPriceModelType

// This is an alias to an internal value.
const PriceUnitWithPercentPriceModelTypeUnitWithPercent = shared.PriceUnitWithPercentPriceModelTypeUnitWithPercent

// This is an alias to an internal type.
type PriceUnitWithPercentPricePriceType = shared.PriceUnitWithPercentPricePriceType

// This is an alias to an internal value.
const PriceUnitWithPercentPricePriceTypeUsagePrice = shared.PriceUnitWithPercentPricePriceTypeUsagePrice

// This is an alias to an internal value.
const PriceUnitWithPercentPricePriceTypeFixedPrice = shared.PriceUnitWithPercentPricePriceTypeFixedPrice

// This is an alias to an internal value.
const PriceUnitWithPercentPricePriceTypeCompositePrice = shared.PriceUnitWithPercentPricePriceTypeCompositePrice

// Configuration for unit_with_percent pricing
//
// This is an alias to an internal type.
type PriceUnitWithPercentPriceUnitWithPercentConfig = shared.PriceUnitWithPercentPriceUnitWithPercentConfig

// The LicenseType resource represents a type of license that can be assigned to
// users. License types are used during billing by grouping metrics on the
// configured grouping key.
//
// This is an alias to an internal type.
type PriceUnitWithPercentPriceLicenseType = shared.PriceUnitWithPercentPriceLicenseType

// This is an alias to an internal type.
type PriceMatrixWithAllocationPrice = shared.PriceMatrixWithAllocationPrice

// This is an alias to an internal type.
type PriceMatrixWithAllocationPriceBillingMode = shared.PriceMatrixWithAllocationPriceBillingMode

// This is an alias to an internal value.
const PriceMatrixWithAllocationPriceBillingModeInAdvance = shared.PriceMatrixWithAllocationPriceBillingModeInAdvance

// This is an alias to an internal value.
const PriceMatrixWithAllocationPriceBillingModeInArrear = shared.PriceMatrixWithAllocationPriceBillingModeInArrear

// This is an alias to an internal type.
type PriceMatrixWithAllocationPriceCadence = shared.PriceMatrixWithAllocationPriceCadence

// This is an alias to an internal value.
const PriceMatrixWithAllocationPriceCadenceOneTime = shared.PriceMatrixWithAllocationPriceCadenceOneTime

// This is an alias to an internal value.
const PriceMatrixWithAllocationPriceCadenceMonthly = shared.PriceMatrixWithAllocationPriceCadenceMonthly

// This is an alias to an internal value.
const PriceMatrixWithAllocationPriceCadenceQuarterly = shared.PriceMatrixWithAllocationPriceCadenceQuarterly

// This is an alias to an internal value.
const PriceMatrixWithAllocationPriceCadenceSemiAnnual = shared.PriceMatrixWithAllocationPriceCadenceSemiAnnual

// This is an alias to an internal value.
const PriceMatrixWithAllocationPriceCadenceAnnual = shared.PriceMatrixWithAllocationPriceCadenceAnnual

// This is an alias to an internal value.
const PriceMatrixWithAllocationPriceCadenceCustom = shared.PriceMatrixWithAllocationPriceCadenceCustom

// This is an alias to an internal type.
type PriceMatrixWithAllocationPriceCompositePriceFilter = shared.PriceMatrixWithAllocationPriceCompositePriceFilter

// The property of the price to filter on.
//
// This is an alias to an internal type.
type PriceMatrixWithAllocationPriceCompositePriceFiltersField = shared.PriceMatrixWithAllocationPriceCompositePriceFiltersField

// This is an alias to an internal value.
const PriceMatrixWithAllocationPriceCompositePriceFiltersFieldPriceID = shared.PriceMatrixWithAllocationPriceCompositePriceFiltersFieldPriceID

// This is an alias to an internal value.
const PriceMatrixWithAllocationPriceCompositePriceFiltersFieldItemID = shared.PriceMatrixWithAllocationPriceCompositePriceFiltersFieldItemID

// This is an alias to an internal value.
const PriceMatrixWithAllocationPriceCompositePriceFiltersFieldPriceType = shared.PriceMatrixWithAllocationPriceCompositePriceFiltersFieldPriceType

// This is an alias to an internal value.
const PriceMatrixWithAllocationPriceCompositePriceFiltersFieldCurrency = shared.PriceMatrixWithAllocationPriceCompositePriceFiltersFieldCurrency

// This is an alias to an internal value.
const PriceMatrixWithAllocationPriceCompositePriceFiltersFieldPricingUnitID = shared.PriceMatrixWithAllocationPriceCompositePriceFiltersFieldPricingUnitID

// Should prices that match the filter be included or excluded.
//
// This is an alias to an internal type.
type PriceMatrixWithAllocationPriceCompositePriceFiltersOperator = shared.PriceMatrixWithAllocationPriceCompositePriceFiltersOperator

// This is an alias to an internal value.
const PriceMatrixWithAllocationPriceCompositePriceFiltersOperatorIncludes = shared.PriceMatrixWithAllocationPriceCompositePriceFiltersOperatorIncludes

// This is an alias to an internal value.
const PriceMatrixWithAllocationPriceCompositePriceFiltersOperatorExcludes = shared.PriceMatrixWithAllocationPriceCompositePriceFiltersOperatorExcludes

// This is an alias to an internal type.
type PriceMatrixWithAllocationPriceConversionRateConfig = shared.PriceMatrixWithAllocationPriceConversionRateConfig

// This is an alias to an internal type.
type PriceMatrixWithAllocationPriceConversionRateConfigConversionRateType = shared.PriceMatrixWithAllocationPriceConversionRateConfigConversionRateType

// This is an alias to an internal value.
const PriceMatrixWithAllocationPriceConversionRateConfigConversionRateTypeUnit = shared.PriceMatrixWithAllocationPriceConversionRateConfigConversionRateTypeUnit

// This is an alias to an internal value.
const PriceMatrixWithAllocationPriceConversionRateConfigConversionRateTypeTiered = shared.PriceMatrixWithAllocationPriceConversionRateConfigConversionRateTypeTiered

// The pricing model type
//
// This is an alias to an internal type.
type PriceMatrixWithAllocationPriceModelType = shared.PriceMatrixWithAllocationPriceModelType

// This is an alias to an internal value.
const PriceMatrixWithAllocationPriceModelTypeMatrixWithAllocation = shared.PriceMatrixWithAllocationPriceModelTypeMatrixWithAllocation

// This is an alias to an internal type.
type PriceMatrixWithAllocationPricePriceType = shared.PriceMatrixWithAllocationPricePriceType

// This is an alias to an internal value.
const PriceMatrixWithAllocationPricePriceTypeUsagePrice = shared.PriceMatrixWithAllocationPricePriceTypeUsagePrice

// This is an alias to an internal value.
const PriceMatrixWithAllocationPricePriceTypeFixedPrice = shared.PriceMatrixWithAllocationPricePriceTypeFixedPrice

// This is an alias to an internal value.
const PriceMatrixWithAllocationPricePriceTypeCompositePrice = shared.PriceMatrixWithAllocationPricePriceTypeCompositePrice

// The LicenseType resource represents a type of license that can be assigned to
// users. License types are used during billing by grouping metrics on the
// configured grouping key.
//
// This is an alias to an internal type.
type PriceMatrixWithAllocationPriceLicenseType = shared.PriceMatrixWithAllocationPriceLicenseType

// This is an alias to an internal type.
type PriceTieredWithProrationPrice = shared.PriceTieredWithProrationPrice

// This is an alias to an internal type.
type PriceTieredWithProrationPriceBillingMode = shared.PriceTieredWithProrationPriceBillingMode

// This is an alias to an internal value.
const PriceTieredWithProrationPriceBillingModeInAdvance = shared.PriceTieredWithProrationPriceBillingModeInAdvance

// This is an alias to an internal value.
const PriceTieredWithProrationPriceBillingModeInArrear = shared.PriceTieredWithProrationPriceBillingModeInArrear

// This is an alias to an internal type.
type PriceTieredWithProrationPriceCadence = shared.PriceTieredWithProrationPriceCadence

// This is an alias to an internal value.
const PriceTieredWithProrationPriceCadenceOneTime = shared.PriceTieredWithProrationPriceCadenceOneTime

// This is an alias to an internal value.
const PriceTieredWithProrationPriceCadenceMonthly = shared.PriceTieredWithProrationPriceCadenceMonthly

// This is an alias to an internal value.
const PriceTieredWithProrationPriceCadenceQuarterly = shared.PriceTieredWithProrationPriceCadenceQuarterly

// This is an alias to an internal value.
const PriceTieredWithProrationPriceCadenceSemiAnnual = shared.PriceTieredWithProrationPriceCadenceSemiAnnual

// This is an alias to an internal value.
const PriceTieredWithProrationPriceCadenceAnnual = shared.PriceTieredWithProrationPriceCadenceAnnual

// This is an alias to an internal value.
const PriceTieredWithProrationPriceCadenceCustom = shared.PriceTieredWithProrationPriceCadenceCustom

// This is an alias to an internal type.
type PriceTieredWithProrationPriceCompositePriceFilter = shared.PriceTieredWithProrationPriceCompositePriceFilter

// The property of the price to filter on.
//
// This is an alias to an internal type.
type PriceTieredWithProrationPriceCompositePriceFiltersField = shared.PriceTieredWithProrationPriceCompositePriceFiltersField

// This is an alias to an internal value.
const PriceTieredWithProrationPriceCompositePriceFiltersFieldPriceID = shared.PriceTieredWithProrationPriceCompositePriceFiltersFieldPriceID

// This is an alias to an internal value.
const PriceTieredWithProrationPriceCompositePriceFiltersFieldItemID = shared.PriceTieredWithProrationPriceCompositePriceFiltersFieldItemID

// This is an alias to an internal value.
const PriceTieredWithProrationPriceCompositePriceFiltersFieldPriceType = shared.PriceTieredWithProrationPriceCompositePriceFiltersFieldPriceType

// This is an alias to an internal value.
const PriceTieredWithProrationPriceCompositePriceFiltersFieldCurrency = shared.PriceTieredWithProrationPriceCompositePriceFiltersFieldCurrency

// This is an alias to an internal value.
const PriceTieredWithProrationPriceCompositePriceFiltersFieldPricingUnitID = shared.PriceTieredWithProrationPriceCompositePriceFiltersFieldPricingUnitID

// Should prices that match the filter be included or excluded.
//
// This is an alias to an internal type.
type PriceTieredWithProrationPriceCompositePriceFiltersOperator = shared.PriceTieredWithProrationPriceCompositePriceFiltersOperator

// This is an alias to an internal value.
const PriceTieredWithProrationPriceCompositePriceFiltersOperatorIncludes = shared.PriceTieredWithProrationPriceCompositePriceFiltersOperatorIncludes

// This is an alias to an internal value.
const PriceTieredWithProrationPriceCompositePriceFiltersOperatorExcludes = shared.PriceTieredWithProrationPriceCompositePriceFiltersOperatorExcludes

// This is an alias to an internal type.
type PriceTieredWithProrationPriceConversionRateConfig = shared.PriceTieredWithProrationPriceConversionRateConfig

// This is an alias to an internal type.
type PriceTieredWithProrationPriceConversionRateConfigConversionRateType = shared.PriceTieredWithProrationPriceConversionRateConfigConversionRateType

// This is an alias to an internal value.
const PriceTieredWithProrationPriceConversionRateConfigConversionRateTypeUnit = shared.PriceTieredWithProrationPriceConversionRateConfigConversionRateTypeUnit

// This is an alias to an internal value.
const PriceTieredWithProrationPriceConversionRateConfigConversionRateTypeTiered = shared.PriceTieredWithProrationPriceConversionRateConfigConversionRateTypeTiered

// The pricing model type
//
// This is an alias to an internal type.
type PriceTieredWithProrationPriceModelType = shared.PriceTieredWithProrationPriceModelType

// This is an alias to an internal value.
const PriceTieredWithProrationPriceModelTypeTieredWithProration = shared.PriceTieredWithProrationPriceModelTypeTieredWithProration

// This is an alias to an internal type.
type PriceTieredWithProrationPricePriceType = shared.PriceTieredWithProrationPricePriceType

// This is an alias to an internal value.
const PriceTieredWithProrationPricePriceTypeUsagePrice = shared.PriceTieredWithProrationPricePriceTypeUsagePrice

// This is an alias to an internal value.
const PriceTieredWithProrationPricePriceTypeFixedPrice = shared.PriceTieredWithProrationPricePriceTypeFixedPrice

// This is an alias to an internal value.
const PriceTieredWithProrationPricePriceTypeCompositePrice = shared.PriceTieredWithProrationPricePriceTypeCompositePrice

// Configuration for tiered_with_proration pricing
//
// This is an alias to an internal type.
type PriceTieredWithProrationPriceTieredWithProrationConfig = shared.PriceTieredWithProrationPriceTieredWithProrationConfig

// Configuration for a single tiered with proration tier
//
// This is an alias to an internal type.
type PriceTieredWithProrationPriceTieredWithProrationConfigTier = shared.PriceTieredWithProrationPriceTieredWithProrationConfigTier

// The LicenseType resource represents a type of license that can be assigned to
// users. License types are used during billing by grouping metrics on the
// configured grouping key.
//
// This is an alias to an internal type.
type PriceTieredWithProrationPriceLicenseType = shared.PriceTieredWithProrationPriceLicenseType

// This is an alias to an internal type.
type PriceUnitWithProrationPrice = shared.PriceUnitWithProrationPrice

// This is an alias to an internal type.
type PriceUnitWithProrationPriceBillingMode = shared.PriceUnitWithProrationPriceBillingMode

// This is an alias to an internal value.
const PriceUnitWithProrationPriceBillingModeInAdvance = shared.PriceUnitWithProrationPriceBillingModeInAdvance

// This is an alias to an internal value.
const PriceUnitWithProrationPriceBillingModeInArrear = shared.PriceUnitWithProrationPriceBillingModeInArrear

// This is an alias to an internal type.
type PriceUnitWithProrationPriceCadence = shared.PriceUnitWithProrationPriceCadence

// This is an alias to an internal value.
const PriceUnitWithProrationPriceCadenceOneTime = shared.PriceUnitWithProrationPriceCadenceOneTime

// This is an alias to an internal value.
const PriceUnitWithProrationPriceCadenceMonthly = shared.PriceUnitWithProrationPriceCadenceMonthly

// This is an alias to an internal value.
const PriceUnitWithProrationPriceCadenceQuarterly = shared.PriceUnitWithProrationPriceCadenceQuarterly

// This is an alias to an internal value.
const PriceUnitWithProrationPriceCadenceSemiAnnual = shared.PriceUnitWithProrationPriceCadenceSemiAnnual

// This is an alias to an internal value.
const PriceUnitWithProrationPriceCadenceAnnual = shared.PriceUnitWithProrationPriceCadenceAnnual

// This is an alias to an internal value.
const PriceUnitWithProrationPriceCadenceCustom = shared.PriceUnitWithProrationPriceCadenceCustom

// This is an alias to an internal type.
type PriceUnitWithProrationPriceCompositePriceFilter = shared.PriceUnitWithProrationPriceCompositePriceFilter

// The property of the price to filter on.
//
// This is an alias to an internal type.
type PriceUnitWithProrationPriceCompositePriceFiltersField = shared.PriceUnitWithProrationPriceCompositePriceFiltersField

// This is an alias to an internal value.
const PriceUnitWithProrationPriceCompositePriceFiltersFieldPriceID = shared.PriceUnitWithProrationPriceCompositePriceFiltersFieldPriceID

// This is an alias to an internal value.
const PriceUnitWithProrationPriceCompositePriceFiltersFieldItemID = shared.PriceUnitWithProrationPriceCompositePriceFiltersFieldItemID

// This is an alias to an internal value.
const PriceUnitWithProrationPriceCompositePriceFiltersFieldPriceType = shared.PriceUnitWithProrationPriceCompositePriceFiltersFieldPriceType

// This is an alias to an internal value.
const PriceUnitWithProrationPriceCompositePriceFiltersFieldCurrency = shared.PriceUnitWithProrationPriceCompositePriceFiltersFieldCurrency

// This is an alias to an internal value.
const PriceUnitWithProrationPriceCompositePriceFiltersFieldPricingUnitID = shared.PriceUnitWithProrationPriceCompositePriceFiltersFieldPricingUnitID

// Should prices that match the filter be included or excluded.
//
// This is an alias to an internal type.
type PriceUnitWithProrationPriceCompositePriceFiltersOperator = shared.PriceUnitWithProrationPriceCompositePriceFiltersOperator

// This is an alias to an internal value.
const PriceUnitWithProrationPriceCompositePriceFiltersOperatorIncludes = shared.PriceUnitWithProrationPriceCompositePriceFiltersOperatorIncludes

// This is an alias to an internal value.
const PriceUnitWithProrationPriceCompositePriceFiltersOperatorExcludes = shared.PriceUnitWithProrationPriceCompositePriceFiltersOperatorExcludes

// This is an alias to an internal type.
type PriceUnitWithProrationPriceConversionRateConfig = shared.PriceUnitWithProrationPriceConversionRateConfig

// This is an alias to an internal type.
type PriceUnitWithProrationPriceConversionRateConfigConversionRateType = shared.PriceUnitWithProrationPriceConversionRateConfigConversionRateType

// This is an alias to an internal value.
const PriceUnitWithProrationPriceConversionRateConfigConversionRateTypeUnit = shared.PriceUnitWithProrationPriceConversionRateConfigConversionRateTypeUnit

// This is an alias to an internal value.
const PriceUnitWithProrationPriceConversionRateConfigConversionRateTypeTiered = shared.PriceUnitWithProrationPriceConversionRateConfigConversionRateTypeTiered

// The pricing model type
//
// This is an alias to an internal type.
type PriceUnitWithProrationPriceModelType = shared.PriceUnitWithProrationPriceModelType

// This is an alias to an internal value.
const PriceUnitWithProrationPriceModelTypeUnitWithProration = shared.PriceUnitWithProrationPriceModelTypeUnitWithProration

// This is an alias to an internal type.
type PriceUnitWithProrationPricePriceType = shared.PriceUnitWithProrationPricePriceType

// This is an alias to an internal value.
const PriceUnitWithProrationPricePriceTypeUsagePrice = shared.PriceUnitWithProrationPricePriceTypeUsagePrice

// This is an alias to an internal value.
const PriceUnitWithProrationPricePriceTypeFixedPrice = shared.PriceUnitWithProrationPricePriceTypeFixedPrice

// This is an alias to an internal value.
const PriceUnitWithProrationPricePriceTypeCompositePrice = shared.PriceUnitWithProrationPricePriceTypeCompositePrice

// Configuration for unit_with_proration pricing
//
// This is an alias to an internal type.
type PriceUnitWithProrationPriceUnitWithProrationConfig = shared.PriceUnitWithProrationPriceUnitWithProrationConfig

// The LicenseType resource represents a type of license that can be assigned to
// users. License types are used during billing by grouping metrics on the
// configured grouping key.
//
// This is an alias to an internal type.
type PriceUnitWithProrationPriceLicenseType = shared.PriceUnitWithProrationPriceLicenseType

// This is an alias to an internal type.
type PriceGroupedAllocationPrice = shared.PriceGroupedAllocationPrice

// This is an alias to an internal type.
type PriceGroupedAllocationPriceBillingMode = shared.PriceGroupedAllocationPriceBillingMode

// This is an alias to an internal value.
const PriceGroupedAllocationPriceBillingModeInAdvance = shared.PriceGroupedAllocationPriceBillingModeInAdvance

// This is an alias to an internal value.
const PriceGroupedAllocationPriceBillingModeInArrear = shared.PriceGroupedAllocationPriceBillingModeInArrear

// This is an alias to an internal type.
type PriceGroupedAllocationPriceCadence = shared.PriceGroupedAllocationPriceCadence

// This is an alias to an internal value.
const PriceGroupedAllocationPriceCadenceOneTime = shared.PriceGroupedAllocationPriceCadenceOneTime

// This is an alias to an internal value.
const PriceGroupedAllocationPriceCadenceMonthly = shared.PriceGroupedAllocationPriceCadenceMonthly

// This is an alias to an internal value.
const PriceGroupedAllocationPriceCadenceQuarterly = shared.PriceGroupedAllocationPriceCadenceQuarterly

// This is an alias to an internal value.
const PriceGroupedAllocationPriceCadenceSemiAnnual = shared.PriceGroupedAllocationPriceCadenceSemiAnnual

// This is an alias to an internal value.
const PriceGroupedAllocationPriceCadenceAnnual = shared.PriceGroupedAllocationPriceCadenceAnnual

// This is an alias to an internal value.
const PriceGroupedAllocationPriceCadenceCustom = shared.PriceGroupedAllocationPriceCadenceCustom

// This is an alias to an internal type.
type PriceGroupedAllocationPriceCompositePriceFilter = shared.PriceGroupedAllocationPriceCompositePriceFilter

// The property of the price to filter on.
//
// This is an alias to an internal type.
type PriceGroupedAllocationPriceCompositePriceFiltersField = shared.PriceGroupedAllocationPriceCompositePriceFiltersField

// This is an alias to an internal value.
const PriceGroupedAllocationPriceCompositePriceFiltersFieldPriceID = shared.PriceGroupedAllocationPriceCompositePriceFiltersFieldPriceID

// This is an alias to an internal value.
const PriceGroupedAllocationPriceCompositePriceFiltersFieldItemID = shared.PriceGroupedAllocationPriceCompositePriceFiltersFieldItemID

// This is an alias to an internal value.
const PriceGroupedAllocationPriceCompositePriceFiltersFieldPriceType = shared.PriceGroupedAllocationPriceCompositePriceFiltersFieldPriceType

// This is an alias to an internal value.
const PriceGroupedAllocationPriceCompositePriceFiltersFieldCurrency = shared.PriceGroupedAllocationPriceCompositePriceFiltersFieldCurrency

// This is an alias to an internal value.
const PriceGroupedAllocationPriceCompositePriceFiltersFieldPricingUnitID = shared.PriceGroupedAllocationPriceCompositePriceFiltersFieldPricingUnitID

// Should prices that match the filter be included or excluded.
//
// This is an alias to an internal type.
type PriceGroupedAllocationPriceCompositePriceFiltersOperator = shared.PriceGroupedAllocationPriceCompositePriceFiltersOperator

// This is an alias to an internal value.
const PriceGroupedAllocationPriceCompositePriceFiltersOperatorIncludes = shared.PriceGroupedAllocationPriceCompositePriceFiltersOperatorIncludes

// This is an alias to an internal value.
const PriceGroupedAllocationPriceCompositePriceFiltersOperatorExcludes = shared.PriceGroupedAllocationPriceCompositePriceFiltersOperatorExcludes

// This is an alias to an internal type.
type PriceGroupedAllocationPriceConversionRateConfig = shared.PriceGroupedAllocationPriceConversionRateConfig

// This is an alias to an internal type.
type PriceGroupedAllocationPriceConversionRateConfigConversionRateType = shared.PriceGroupedAllocationPriceConversionRateConfigConversionRateType

// This is an alias to an internal value.
const PriceGroupedAllocationPriceConversionRateConfigConversionRateTypeUnit = shared.PriceGroupedAllocationPriceConversionRateConfigConversionRateTypeUnit

// This is an alias to an internal value.
const PriceGroupedAllocationPriceConversionRateConfigConversionRateTypeTiered = shared.PriceGroupedAllocationPriceConversionRateConfigConversionRateTypeTiered

// Configuration for grouped_allocation pricing
//
// This is an alias to an internal type.
type PriceGroupedAllocationPriceGroupedAllocationConfig = shared.PriceGroupedAllocationPriceGroupedAllocationConfig

// The pricing model type
//
// This is an alias to an internal type.
type PriceGroupedAllocationPriceModelType = shared.PriceGroupedAllocationPriceModelType

// This is an alias to an internal value.
const PriceGroupedAllocationPriceModelTypeGroupedAllocation = shared.PriceGroupedAllocationPriceModelTypeGroupedAllocation

// This is an alias to an internal type.
type PriceGroupedAllocationPricePriceType = shared.PriceGroupedAllocationPricePriceType

// This is an alias to an internal value.
const PriceGroupedAllocationPricePriceTypeUsagePrice = shared.PriceGroupedAllocationPricePriceTypeUsagePrice

// This is an alias to an internal value.
const PriceGroupedAllocationPricePriceTypeFixedPrice = shared.PriceGroupedAllocationPricePriceTypeFixedPrice

// This is an alias to an internal value.
const PriceGroupedAllocationPricePriceTypeCompositePrice = shared.PriceGroupedAllocationPricePriceTypeCompositePrice

// The LicenseType resource represents a type of license that can be assigned to
// users. License types are used during billing by grouping metrics on the
// configured grouping key.
//
// This is an alias to an internal type.
type PriceGroupedAllocationPriceLicenseType = shared.PriceGroupedAllocationPriceLicenseType

// This is an alias to an internal type.
type PriceBulkWithProrationPrice = shared.PriceBulkWithProrationPrice

// This is an alias to an internal type.
type PriceBulkWithProrationPriceBillingMode = shared.PriceBulkWithProrationPriceBillingMode

// This is an alias to an internal value.
const PriceBulkWithProrationPriceBillingModeInAdvance = shared.PriceBulkWithProrationPriceBillingModeInAdvance

// This is an alias to an internal value.
const PriceBulkWithProrationPriceBillingModeInArrear = shared.PriceBulkWithProrationPriceBillingModeInArrear

// Configuration for bulk_with_proration pricing
//
// This is an alias to an internal type.
type PriceBulkWithProrationPriceBulkWithProrationConfig = shared.PriceBulkWithProrationPriceBulkWithProrationConfig

// Configuration for a single bulk pricing tier with proration
//
// This is an alias to an internal type.
type PriceBulkWithProrationPriceBulkWithProrationConfigTier = shared.PriceBulkWithProrationPriceBulkWithProrationConfigTier

// This is an alias to an internal type.
type PriceBulkWithProrationPriceCadence = shared.PriceBulkWithProrationPriceCadence

// This is an alias to an internal value.
const PriceBulkWithProrationPriceCadenceOneTime = shared.PriceBulkWithProrationPriceCadenceOneTime

// This is an alias to an internal value.
const PriceBulkWithProrationPriceCadenceMonthly = shared.PriceBulkWithProrationPriceCadenceMonthly

// This is an alias to an internal value.
const PriceBulkWithProrationPriceCadenceQuarterly = shared.PriceBulkWithProrationPriceCadenceQuarterly

// This is an alias to an internal value.
const PriceBulkWithProrationPriceCadenceSemiAnnual = shared.PriceBulkWithProrationPriceCadenceSemiAnnual

// This is an alias to an internal value.
const PriceBulkWithProrationPriceCadenceAnnual = shared.PriceBulkWithProrationPriceCadenceAnnual

// This is an alias to an internal value.
const PriceBulkWithProrationPriceCadenceCustom = shared.PriceBulkWithProrationPriceCadenceCustom

// This is an alias to an internal type.
type PriceBulkWithProrationPriceCompositePriceFilter = shared.PriceBulkWithProrationPriceCompositePriceFilter

// The property of the price to filter on.
//
// This is an alias to an internal type.
type PriceBulkWithProrationPriceCompositePriceFiltersField = shared.PriceBulkWithProrationPriceCompositePriceFiltersField

// This is an alias to an internal value.
const PriceBulkWithProrationPriceCompositePriceFiltersFieldPriceID = shared.PriceBulkWithProrationPriceCompositePriceFiltersFieldPriceID

// This is an alias to an internal value.
const PriceBulkWithProrationPriceCompositePriceFiltersFieldItemID = shared.PriceBulkWithProrationPriceCompositePriceFiltersFieldItemID

// This is an alias to an internal value.
const PriceBulkWithProrationPriceCompositePriceFiltersFieldPriceType = shared.PriceBulkWithProrationPriceCompositePriceFiltersFieldPriceType

// This is an alias to an internal value.
const PriceBulkWithProrationPriceCompositePriceFiltersFieldCurrency = shared.PriceBulkWithProrationPriceCompositePriceFiltersFieldCurrency

// This is an alias to an internal value.
const PriceBulkWithProrationPriceCompositePriceFiltersFieldPricingUnitID = shared.PriceBulkWithProrationPriceCompositePriceFiltersFieldPricingUnitID

// Should prices that match the filter be included or excluded.
//
// This is an alias to an internal type.
type PriceBulkWithProrationPriceCompositePriceFiltersOperator = shared.PriceBulkWithProrationPriceCompositePriceFiltersOperator

// This is an alias to an internal value.
const PriceBulkWithProrationPriceCompositePriceFiltersOperatorIncludes = shared.PriceBulkWithProrationPriceCompositePriceFiltersOperatorIncludes

// This is an alias to an internal value.
const PriceBulkWithProrationPriceCompositePriceFiltersOperatorExcludes = shared.PriceBulkWithProrationPriceCompositePriceFiltersOperatorExcludes

// This is an alias to an internal type.
type PriceBulkWithProrationPriceConversionRateConfig = shared.PriceBulkWithProrationPriceConversionRateConfig

// This is an alias to an internal type.
type PriceBulkWithProrationPriceConversionRateConfigConversionRateType = shared.PriceBulkWithProrationPriceConversionRateConfigConversionRateType

// This is an alias to an internal value.
const PriceBulkWithProrationPriceConversionRateConfigConversionRateTypeUnit = shared.PriceBulkWithProrationPriceConversionRateConfigConversionRateTypeUnit

// This is an alias to an internal value.
const PriceBulkWithProrationPriceConversionRateConfigConversionRateTypeTiered = shared.PriceBulkWithProrationPriceConversionRateConfigConversionRateTypeTiered

// The pricing model type
//
// This is an alias to an internal type.
type PriceBulkWithProrationPriceModelType = shared.PriceBulkWithProrationPriceModelType

// This is an alias to an internal value.
const PriceBulkWithProrationPriceModelTypeBulkWithProration = shared.PriceBulkWithProrationPriceModelTypeBulkWithProration

// This is an alias to an internal type.
type PriceBulkWithProrationPricePriceType = shared.PriceBulkWithProrationPricePriceType

// This is an alias to an internal value.
const PriceBulkWithProrationPricePriceTypeUsagePrice = shared.PriceBulkWithProrationPricePriceTypeUsagePrice

// This is an alias to an internal value.
const PriceBulkWithProrationPricePriceTypeFixedPrice = shared.PriceBulkWithProrationPricePriceTypeFixedPrice

// This is an alias to an internal value.
const PriceBulkWithProrationPricePriceTypeCompositePrice = shared.PriceBulkWithProrationPricePriceTypeCompositePrice

// The LicenseType resource represents a type of license that can be assigned to
// users. License types are used during billing by grouping metrics on the
// configured grouping key.
//
// This is an alias to an internal type.
type PriceBulkWithProrationPriceLicenseType = shared.PriceBulkWithProrationPriceLicenseType

// This is an alias to an internal type.
type PriceGroupedWithProratedMinimumPrice = shared.PriceGroupedWithProratedMinimumPrice

// This is an alias to an internal type.
type PriceGroupedWithProratedMinimumPriceBillingMode = shared.PriceGroupedWithProratedMinimumPriceBillingMode

// This is an alias to an internal value.
const PriceGroupedWithProratedMinimumPriceBillingModeInAdvance = shared.PriceGroupedWithProratedMinimumPriceBillingModeInAdvance

// This is an alias to an internal value.
const PriceGroupedWithProratedMinimumPriceBillingModeInArrear = shared.PriceGroupedWithProratedMinimumPriceBillingModeInArrear

// This is an alias to an internal type.
type PriceGroupedWithProratedMinimumPriceCadence = shared.PriceGroupedWithProratedMinimumPriceCadence

// This is an alias to an internal value.
const PriceGroupedWithProratedMinimumPriceCadenceOneTime = shared.PriceGroupedWithProratedMinimumPriceCadenceOneTime

// This is an alias to an internal value.
const PriceGroupedWithProratedMinimumPriceCadenceMonthly = shared.PriceGroupedWithProratedMinimumPriceCadenceMonthly

// This is an alias to an internal value.
const PriceGroupedWithProratedMinimumPriceCadenceQuarterly = shared.PriceGroupedWithProratedMinimumPriceCadenceQuarterly

// This is an alias to an internal value.
const PriceGroupedWithProratedMinimumPriceCadenceSemiAnnual = shared.PriceGroupedWithProratedMinimumPriceCadenceSemiAnnual

// This is an alias to an internal value.
const PriceGroupedWithProratedMinimumPriceCadenceAnnual = shared.PriceGroupedWithProratedMinimumPriceCadenceAnnual

// This is an alias to an internal value.
const PriceGroupedWithProratedMinimumPriceCadenceCustom = shared.PriceGroupedWithProratedMinimumPriceCadenceCustom

// This is an alias to an internal type.
type PriceGroupedWithProratedMinimumPriceCompositePriceFilter = shared.PriceGroupedWithProratedMinimumPriceCompositePriceFilter

// The property of the price to filter on.
//
// This is an alias to an internal type.
type PriceGroupedWithProratedMinimumPriceCompositePriceFiltersField = shared.PriceGroupedWithProratedMinimumPriceCompositePriceFiltersField

// This is an alias to an internal value.
const PriceGroupedWithProratedMinimumPriceCompositePriceFiltersFieldPriceID = shared.PriceGroupedWithProratedMinimumPriceCompositePriceFiltersFieldPriceID

// This is an alias to an internal value.
const PriceGroupedWithProratedMinimumPriceCompositePriceFiltersFieldItemID = shared.PriceGroupedWithProratedMinimumPriceCompositePriceFiltersFieldItemID

// This is an alias to an internal value.
const PriceGroupedWithProratedMinimumPriceCompositePriceFiltersFieldPriceType = shared.PriceGroupedWithProratedMinimumPriceCompositePriceFiltersFieldPriceType

// This is an alias to an internal value.
const PriceGroupedWithProratedMinimumPriceCompositePriceFiltersFieldCurrency = shared.PriceGroupedWithProratedMinimumPriceCompositePriceFiltersFieldCurrency

// This is an alias to an internal value.
const PriceGroupedWithProratedMinimumPriceCompositePriceFiltersFieldPricingUnitID = shared.PriceGroupedWithProratedMinimumPriceCompositePriceFiltersFieldPricingUnitID

// Should prices that match the filter be included or excluded.
//
// This is an alias to an internal type.
type PriceGroupedWithProratedMinimumPriceCompositePriceFiltersOperator = shared.PriceGroupedWithProratedMinimumPriceCompositePriceFiltersOperator

// This is an alias to an internal value.
const PriceGroupedWithProratedMinimumPriceCompositePriceFiltersOperatorIncludes = shared.PriceGroupedWithProratedMinimumPriceCompositePriceFiltersOperatorIncludes

// This is an alias to an internal value.
const PriceGroupedWithProratedMinimumPriceCompositePriceFiltersOperatorExcludes = shared.PriceGroupedWithProratedMinimumPriceCompositePriceFiltersOperatorExcludes

// This is an alias to an internal type.
type PriceGroupedWithProratedMinimumPriceConversionRateConfig = shared.PriceGroupedWithProratedMinimumPriceConversionRateConfig

// This is an alias to an internal type.
type PriceGroupedWithProratedMinimumPriceConversionRateConfigConversionRateType = shared.PriceGroupedWithProratedMinimumPriceConversionRateConfigConversionRateType

// This is an alias to an internal value.
const PriceGroupedWithProratedMinimumPriceConversionRateConfigConversionRateTypeUnit = shared.PriceGroupedWithProratedMinimumPriceConversionRateConfigConversionRateTypeUnit

// This is an alias to an internal value.
const PriceGroupedWithProratedMinimumPriceConversionRateConfigConversionRateTypeTiered = shared.PriceGroupedWithProratedMinimumPriceConversionRateConfigConversionRateTypeTiered

// Configuration for grouped_with_prorated_minimum pricing
//
// This is an alias to an internal type.
type PriceGroupedWithProratedMinimumPriceGroupedWithProratedMinimumConfig = shared.PriceGroupedWithProratedMinimumPriceGroupedWithProratedMinimumConfig

// The pricing model type
//
// This is an alias to an internal type.
type PriceGroupedWithProratedMinimumPriceModelType = shared.PriceGroupedWithProratedMinimumPriceModelType

// This is an alias to an internal value.
const PriceGroupedWithProratedMinimumPriceModelTypeGroupedWithProratedMinimum = shared.PriceGroupedWithProratedMinimumPriceModelTypeGroupedWithProratedMinimum

// This is an alias to an internal type.
type PriceGroupedWithProratedMinimumPricePriceType = shared.PriceGroupedWithProratedMinimumPricePriceType

// This is an alias to an internal value.
const PriceGroupedWithProratedMinimumPricePriceTypeUsagePrice = shared.PriceGroupedWithProratedMinimumPricePriceTypeUsagePrice

// This is an alias to an internal value.
const PriceGroupedWithProratedMinimumPricePriceTypeFixedPrice = shared.PriceGroupedWithProratedMinimumPricePriceTypeFixedPrice

// This is an alias to an internal value.
const PriceGroupedWithProratedMinimumPricePriceTypeCompositePrice = shared.PriceGroupedWithProratedMinimumPricePriceTypeCompositePrice

// The LicenseType resource represents a type of license that can be assigned to
// users. License types are used during billing by grouping metrics on the
// configured grouping key.
//
// This is an alias to an internal type.
type PriceGroupedWithProratedMinimumPriceLicenseType = shared.PriceGroupedWithProratedMinimumPriceLicenseType

// This is an alias to an internal type.
type PriceGroupedWithMeteredMinimumPrice = shared.PriceGroupedWithMeteredMinimumPrice

// This is an alias to an internal type.
type PriceGroupedWithMeteredMinimumPriceBillingMode = shared.PriceGroupedWithMeteredMinimumPriceBillingMode

// This is an alias to an internal value.
const PriceGroupedWithMeteredMinimumPriceBillingModeInAdvance = shared.PriceGroupedWithMeteredMinimumPriceBillingModeInAdvance

// This is an alias to an internal value.
const PriceGroupedWithMeteredMinimumPriceBillingModeInArrear = shared.PriceGroupedWithMeteredMinimumPriceBillingModeInArrear

// This is an alias to an internal type.
type PriceGroupedWithMeteredMinimumPriceCadence = shared.PriceGroupedWithMeteredMinimumPriceCadence

// This is an alias to an internal value.
const PriceGroupedWithMeteredMinimumPriceCadenceOneTime = shared.PriceGroupedWithMeteredMinimumPriceCadenceOneTime

// This is an alias to an internal value.
const PriceGroupedWithMeteredMinimumPriceCadenceMonthly = shared.PriceGroupedWithMeteredMinimumPriceCadenceMonthly

// This is an alias to an internal value.
const PriceGroupedWithMeteredMinimumPriceCadenceQuarterly = shared.PriceGroupedWithMeteredMinimumPriceCadenceQuarterly

// This is an alias to an internal value.
const PriceGroupedWithMeteredMinimumPriceCadenceSemiAnnual = shared.PriceGroupedWithMeteredMinimumPriceCadenceSemiAnnual

// This is an alias to an internal value.
const PriceGroupedWithMeteredMinimumPriceCadenceAnnual = shared.PriceGroupedWithMeteredMinimumPriceCadenceAnnual

// This is an alias to an internal value.
const PriceGroupedWithMeteredMinimumPriceCadenceCustom = shared.PriceGroupedWithMeteredMinimumPriceCadenceCustom

// This is an alias to an internal type.
type PriceGroupedWithMeteredMinimumPriceCompositePriceFilter = shared.PriceGroupedWithMeteredMinimumPriceCompositePriceFilter

// The property of the price to filter on.
//
// This is an alias to an internal type.
type PriceGroupedWithMeteredMinimumPriceCompositePriceFiltersField = shared.PriceGroupedWithMeteredMinimumPriceCompositePriceFiltersField

// This is an alias to an internal value.
const PriceGroupedWithMeteredMinimumPriceCompositePriceFiltersFieldPriceID = shared.PriceGroupedWithMeteredMinimumPriceCompositePriceFiltersFieldPriceID

// This is an alias to an internal value.
const PriceGroupedWithMeteredMinimumPriceCompositePriceFiltersFieldItemID = shared.PriceGroupedWithMeteredMinimumPriceCompositePriceFiltersFieldItemID

// This is an alias to an internal value.
const PriceGroupedWithMeteredMinimumPriceCompositePriceFiltersFieldPriceType = shared.PriceGroupedWithMeteredMinimumPriceCompositePriceFiltersFieldPriceType

// This is an alias to an internal value.
const PriceGroupedWithMeteredMinimumPriceCompositePriceFiltersFieldCurrency = shared.PriceGroupedWithMeteredMinimumPriceCompositePriceFiltersFieldCurrency

// This is an alias to an internal value.
const PriceGroupedWithMeteredMinimumPriceCompositePriceFiltersFieldPricingUnitID = shared.PriceGroupedWithMeteredMinimumPriceCompositePriceFiltersFieldPricingUnitID

// Should prices that match the filter be included or excluded.
//
// This is an alias to an internal type.
type PriceGroupedWithMeteredMinimumPriceCompositePriceFiltersOperator = shared.PriceGroupedWithMeteredMinimumPriceCompositePriceFiltersOperator

// This is an alias to an internal value.
const PriceGroupedWithMeteredMinimumPriceCompositePriceFiltersOperatorIncludes = shared.PriceGroupedWithMeteredMinimumPriceCompositePriceFiltersOperatorIncludes

// This is an alias to an internal value.
const PriceGroupedWithMeteredMinimumPriceCompositePriceFiltersOperatorExcludes = shared.PriceGroupedWithMeteredMinimumPriceCompositePriceFiltersOperatorExcludes

// This is an alias to an internal type.
type PriceGroupedWithMeteredMinimumPriceConversionRateConfig = shared.PriceGroupedWithMeteredMinimumPriceConversionRateConfig

// This is an alias to an internal type.
type PriceGroupedWithMeteredMinimumPriceConversionRateConfigConversionRateType = shared.PriceGroupedWithMeteredMinimumPriceConversionRateConfigConversionRateType

// This is an alias to an internal value.
const PriceGroupedWithMeteredMinimumPriceConversionRateConfigConversionRateTypeUnit = shared.PriceGroupedWithMeteredMinimumPriceConversionRateConfigConversionRateTypeUnit

// This is an alias to an internal value.
const PriceGroupedWithMeteredMinimumPriceConversionRateConfigConversionRateTypeTiered = shared.PriceGroupedWithMeteredMinimumPriceConversionRateConfigConversionRateTypeTiered

// Configuration for grouped_with_metered_minimum pricing
//
// This is an alias to an internal type.
type PriceGroupedWithMeteredMinimumPriceGroupedWithMeteredMinimumConfig = shared.PriceGroupedWithMeteredMinimumPriceGroupedWithMeteredMinimumConfig

// Configuration for a scaling factor
//
// This is an alias to an internal type.
type PriceGroupedWithMeteredMinimumPriceGroupedWithMeteredMinimumConfigScalingFactor = shared.PriceGroupedWithMeteredMinimumPriceGroupedWithMeteredMinimumConfigScalingFactor

// Configuration for a unit amount
//
// This is an alias to an internal type.
type PriceGroupedWithMeteredMinimumPriceGroupedWithMeteredMinimumConfigUnitAmount = shared.PriceGroupedWithMeteredMinimumPriceGroupedWithMeteredMinimumConfigUnitAmount

// The pricing model type
//
// This is an alias to an internal type.
type PriceGroupedWithMeteredMinimumPriceModelType = shared.PriceGroupedWithMeteredMinimumPriceModelType

// This is an alias to an internal value.
const PriceGroupedWithMeteredMinimumPriceModelTypeGroupedWithMeteredMinimum = shared.PriceGroupedWithMeteredMinimumPriceModelTypeGroupedWithMeteredMinimum

// This is an alias to an internal type.
type PriceGroupedWithMeteredMinimumPricePriceType = shared.PriceGroupedWithMeteredMinimumPricePriceType

// This is an alias to an internal value.
const PriceGroupedWithMeteredMinimumPricePriceTypeUsagePrice = shared.PriceGroupedWithMeteredMinimumPricePriceTypeUsagePrice

// This is an alias to an internal value.
const PriceGroupedWithMeteredMinimumPricePriceTypeFixedPrice = shared.PriceGroupedWithMeteredMinimumPricePriceTypeFixedPrice

// This is an alias to an internal value.
const PriceGroupedWithMeteredMinimumPricePriceTypeCompositePrice = shared.PriceGroupedWithMeteredMinimumPricePriceTypeCompositePrice

// The LicenseType resource represents a type of license that can be assigned to
// users. License types are used during billing by grouping metrics on the
// configured grouping key.
//
// This is an alias to an internal type.
type PriceGroupedWithMeteredMinimumPriceLicenseType = shared.PriceGroupedWithMeteredMinimumPriceLicenseType

// This is an alias to an internal type.
type PriceGroupedWithMinMaxThresholdsPrice = shared.PriceGroupedWithMinMaxThresholdsPrice

// This is an alias to an internal type.
type PriceGroupedWithMinMaxThresholdsPriceBillingMode = shared.PriceGroupedWithMinMaxThresholdsPriceBillingMode

// This is an alias to an internal value.
const PriceGroupedWithMinMaxThresholdsPriceBillingModeInAdvance = shared.PriceGroupedWithMinMaxThresholdsPriceBillingModeInAdvance

// This is an alias to an internal value.
const PriceGroupedWithMinMaxThresholdsPriceBillingModeInArrear = shared.PriceGroupedWithMinMaxThresholdsPriceBillingModeInArrear

// This is an alias to an internal type.
type PriceGroupedWithMinMaxThresholdsPriceCadence = shared.PriceGroupedWithMinMaxThresholdsPriceCadence

// This is an alias to an internal value.
const PriceGroupedWithMinMaxThresholdsPriceCadenceOneTime = shared.PriceGroupedWithMinMaxThresholdsPriceCadenceOneTime

// This is an alias to an internal value.
const PriceGroupedWithMinMaxThresholdsPriceCadenceMonthly = shared.PriceGroupedWithMinMaxThresholdsPriceCadenceMonthly

// This is an alias to an internal value.
const PriceGroupedWithMinMaxThresholdsPriceCadenceQuarterly = shared.PriceGroupedWithMinMaxThresholdsPriceCadenceQuarterly

// This is an alias to an internal value.
const PriceGroupedWithMinMaxThresholdsPriceCadenceSemiAnnual = shared.PriceGroupedWithMinMaxThresholdsPriceCadenceSemiAnnual

// This is an alias to an internal value.
const PriceGroupedWithMinMaxThresholdsPriceCadenceAnnual = shared.PriceGroupedWithMinMaxThresholdsPriceCadenceAnnual

// This is an alias to an internal value.
const PriceGroupedWithMinMaxThresholdsPriceCadenceCustom = shared.PriceGroupedWithMinMaxThresholdsPriceCadenceCustom

// This is an alias to an internal type.
type PriceGroupedWithMinMaxThresholdsPriceCompositePriceFilter = shared.PriceGroupedWithMinMaxThresholdsPriceCompositePriceFilter

// The property of the price to filter on.
//
// This is an alias to an internal type.
type PriceGroupedWithMinMaxThresholdsPriceCompositePriceFiltersField = shared.PriceGroupedWithMinMaxThresholdsPriceCompositePriceFiltersField

// This is an alias to an internal value.
const PriceGroupedWithMinMaxThresholdsPriceCompositePriceFiltersFieldPriceID = shared.PriceGroupedWithMinMaxThresholdsPriceCompositePriceFiltersFieldPriceID

// This is an alias to an internal value.
const PriceGroupedWithMinMaxThresholdsPriceCompositePriceFiltersFieldItemID = shared.PriceGroupedWithMinMaxThresholdsPriceCompositePriceFiltersFieldItemID

// This is an alias to an internal value.
const PriceGroupedWithMinMaxThresholdsPriceCompositePriceFiltersFieldPriceType = shared.PriceGroupedWithMinMaxThresholdsPriceCompositePriceFiltersFieldPriceType

// This is an alias to an internal value.
const PriceGroupedWithMinMaxThresholdsPriceCompositePriceFiltersFieldCurrency = shared.PriceGroupedWithMinMaxThresholdsPriceCompositePriceFiltersFieldCurrency

// This is an alias to an internal value.
const PriceGroupedWithMinMaxThresholdsPriceCompositePriceFiltersFieldPricingUnitID = shared.PriceGroupedWithMinMaxThresholdsPriceCompositePriceFiltersFieldPricingUnitID

// Should prices that match the filter be included or excluded.
//
// This is an alias to an internal type.
type PriceGroupedWithMinMaxThresholdsPriceCompositePriceFiltersOperator = shared.PriceGroupedWithMinMaxThresholdsPriceCompositePriceFiltersOperator

// This is an alias to an internal value.
const PriceGroupedWithMinMaxThresholdsPriceCompositePriceFiltersOperatorIncludes = shared.PriceGroupedWithMinMaxThresholdsPriceCompositePriceFiltersOperatorIncludes

// This is an alias to an internal value.
const PriceGroupedWithMinMaxThresholdsPriceCompositePriceFiltersOperatorExcludes = shared.PriceGroupedWithMinMaxThresholdsPriceCompositePriceFiltersOperatorExcludes

// This is an alias to an internal type.
type PriceGroupedWithMinMaxThresholdsPriceConversionRateConfig = shared.PriceGroupedWithMinMaxThresholdsPriceConversionRateConfig

// This is an alias to an internal type.
type PriceGroupedWithMinMaxThresholdsPriceConversionRateConfigConversionRateType = shared.PriceGroupedWithMinMaxThresholdsPriceConversionRateConfigConversionRateType

// This is an alias to an internal value.
const PriceGroupedWithMinMaxThresholdsPriceConversionRateConfigConversionRateTypeUnit = shared.PriceGroupedWithMinMaxThresholdsPriceConversionRateConfigConversionRateTypeUnit

// This is an alias to an internal value.
const PriceGroupedWithMinMaxThresholdsPriceConversionRateConfigConversionRateTypeTiered = shared.PriceGroupedWithMinMaxThresholdsPriceConversionRateConfigConversionRateTypeTiered

// Configuration for grouped_with_min_max_thresholds pricing
//
// This is an alias to an internal type.
type PriceGroupedWithMinMaxThresholdsPriceGroupedWithMinMaxThresholdsConfig = shared.PriceGroupedWithMinMaxThresholdsPriceGroupedWithMinMaxThresholdsConfig

// The pricing model type
//
// This is an alias to an internal type.
type PriceGroupedWithMinMaxThresholdsPriceModelType = shared.PriceGroupedWithMinMaxThresholdsPriceModelType

// This is an alias to an internal value.
const PriceGroupedWithMinMaxThresholdsPriceModelTypeGroupedWithMinMaxThresholds = shared.PriceGroupedWithMinMaxThresholdsPriceModelTypeGroupedWithMinMaxThresholds

// This is an alias to an internal type.
type PriceGroupedWithMinMaxThresholdsPricePriceType = shared.PriceGroupedWithMinMaxThresholdsPricePriceType

// This is an alias to an internal value.
const PriceGroupedWithMinMaxThresholdsPricePriceTypeUsagePrice = shared.PriceGroupedWithMinMaxThresholdsPricePriceTypeUsagePrice

// This is an alias to an internal value.
const PriceGroupedWithMinMaxThresholdsPricePriceTypeFixedPrice = shared.PriceGroupedWithMinMaxThresholdsPricePriceTypeFixedPrice

// This is an alias to an internal value.
const PriceGroupedWithMinMaxThresholdsPricePriceTypeCompositePrice = shared.PriceGroupedWithMinMaxThresholdsPricePriceTypeCompositePrice

// The LicenseType resource represents a type of license that can be assigned to
// users. License types are used during billing by grouping metrics on the
// configured grouping key.
//
// This is an alias to an internal type.
type PriceGroupedWithMinMaxThresholdsPriceLicenseType = shared.PriceGroupedWithMinMaxThresholdsPriceLicenseType

// This is an alias to an internal type.
type PriceMatrixWithDisplayNamePrice = shared.PriceMatrixWithDisplayNamePrice

// This is an alias to an internal type.
type PriceMatrixWithDisplayNamePriceBillingMode = shared.PriceMatrixWithDisplayNamePriceBillingMode

// This is an alias to an internal value.
const PriceMatrixWithDisplayNamePriceBillingModeInAdvance = shared.PriceMatrixWithDisplayNamePriceBillingModeInAdvance

// This is an alias to an internal value.
const PriceMatrixWithDisplayNamePriceBillingModeInArrear = shared.PriceMatrixWithDisplayNamePriceBillingModeInArrear

// This is an alias to an internal type.
type PriceMatrixWithDisplayNamePriceCadence = shared.PriceMatrixWithDisplayNamePriceCadence

// This is an alias to an internal value.
const PriceMatrixWithDisplayNamePriceCadenceOneTime = shared.PriceMatrixWithDisplayNamePriceCadenceOneTime

// This is an alias to an internal value.
const PriceMatrixWithDisplayNamePriceCadenceMonthly = shared.PriceMatrixWithDisplayNamePriceCadenceMonthly

// This is an alias to an internal value.
const PriceMatrixWithDisplayNamePriceCadenceQuarterly = shared.PriceMatrixWithDisplayNamePriceCadenceQuarterly

// This is an alias to an internal value.
const PriceMatrixWithDisplayNamePriceCadenceSemiAnnual = shared.PriceMatrixWithDisplayNamePriceCadenceSemiAnnual

// This is an alias to an internal value.
const PriceMatrixWithDisplayNamePriceCadenceAnnual = shared.PriceMatrixWithDisplayNamePriceCadenceAnnual

// This is an alias to an internal value.
const PriceMatrixWithDisplayNamePriceCadenceCustom = shared.PriceMatrixWithDisplayNamePriceCadenceCustom

// This is an alias to an internal type.
type PriceMatrixWithDisplayNamePriceCompositePriceFilter = shared.PriceMatrixWithDisplayNamePriceCompositePriceFilter

// The property of the price to filter on.
//
// This is an alias to an internal type.
type PriceMatrixWithDisplayNamePriceCompositePriceFiltersField = shared.PriceMatrixWithDisplayNamePriceCompositePriceFiltersField

// This is an alias to an internal value.
const PriceMatrixWithDisplayNamePriceCompositePriceFiltersFieldPriceID = shared.PriceMatrixWithDisplayNamePriceCompositePriceFiltersFieldPriceID

// This is an alias to an internal value.
const PriceMatrixWithDisplayNamePriceCompositePriceFiltersFieldItemID = shared.PriceMatrixWithDisplayNamePriceCompositePriceFiltersFieldItemID

// This is an alias to an internal value.
const PriceMatrixWithDisplayNamePriceCompositePriceFiltersFieldPriceType = shared.PriceMatrixWithDisplayNamePriceCompositePriceFiltersFieldPriceType

// This is an alias to an internal value.
const PriceMatrixWithDisplayNamePriceCompositePriceFiltersFieldCurrency = shared.PriceMatrixWithDisplayNamePriceCompositePriceFiltersFieldCurrency

// This is an alias to an internal value.
const PriceMatrixWithDisplayNamePriceCompositePriceFiltersFieldPricingUnitID = shared.PriceMatrixWithDisplayNamePriceCompositePriceFiltersFieldPricingUnitID

// Should prices that match the filter be included or excluded.
//
// This is an alias to an internal type.
type PriceMatrixWithDisplayNamePriceCompositePriceFiltersOperator = shared.PriceMatrixWithDisplayNamePriceCompositePriceFiltersOperator

// This is an alias to an internal value.
const PriceMatrixWithDisplayNamePriceCompositePriceFiltersOperatorIncludes = shared.PriceMatrixWithDisplayNamePriceCompositePriceFiltersOperatorIncludes

// This is an alias to an internal value.
const PriceMatrixWithDisplayNamePriceCompositePriceFiltersOperatorExcludes = shared.PriceMatrixWithDisplayNamePriceCompositePriceFiltersOperatorExcludes

// This is an alias to an internal type.
type PriceMatrixWithDisplayNamePriceConversionRateConfig = shared.PriceMatrixWithDisplayNamePriceConversionRateConfig

// This is an alias to an internal type.
type PriceMatrixWithDisplayNamePriceConversionRateConfigConversionRateType = shared.PriceMatrixWithDisplayNamePriceConversionRateConfigConversionRateType

// This is an alias to an internal value.
const PriceMatrixWithDisplayNamePriceConversionRateConfigConversionRateTypeUnit = shared.PriceMatrixWithDisplayNamePriceConversionRateConfigConversionRateTypeUnit

// This is an alias to an internal value.
const PriceMatrixWithDisplayNamePriceConversionRateConfigConversionRateTypeTiered = shared.PriceMatrixWithDisplayNamePriceConversionRateConfigConversionRateTypeTiered

// Configuration for matrix_with_display_name pricing
//
// This is an alias to an internal type.
type PriceMatrixWithDisplayNamePriceMatrixWithDisplayNameConfig = shared.PriceMatrixWithDisplayNamePriceMatrixWithDisplayNameConfig

// Configuration for a unit amount item
//
// This is an alias to an internal type.
type PriceMatrixWithDisplayNamePriceMatrixWithDisplayNameConfigUnitAmount = shared.PriceMatrixWithDisplayNamePriceMatrixWithDisplayNameConfigUnitAmount

// The pricing model type
//
// This is an alias to an internal type.
type PriceMatrixWithDisplayNamePriceModelType = shared.PriceMatrixWithDisplayNamePriceModelType

// This is an alias to an internal value.
const PriceMatrixWithDisplayNamePriceModelTypeMatrixWithDisplayName = shared.PriceMatrixWithDisplayNamePriceModelTypeMatrixWithDisplayName

// This is an alias to an internal type.
type PriceMatrixWithDisplayNamePricePriceType = shared.PriceMatrixWithDisplayNamePricePriceType

// This is an alias to an internal value.
const PriceMatrixWithDisplayNamePricePriceTypeUsagePrice = shared.PriceMatrixWithDisplayNamePricePriceTypeUsagePrice

// This is an alias to an internal value.
const PriceMatrixWithDisplayNamePricePriceTypeFixedPrice = shared.PriceMatrixWithDisplayNamePricePriceTypeFixedPrice

// This is an alias to an internal value.
const PriceMatrixWithDisplayNamePricePriceTypeCompositePrice = shared.PriceMatrixWithDisplayNamePricePriceTypeCompositePrice

// The LicenseType resource represents a type of license that can be assigned to
// users. License types are used during billing by grouping metrics on the
// configured grouping key.
//
// This is an alias to an internal type.
type PriceMatrixWithDisplayNamePriceLicenseType = shared.PriceMatrixWithDisplayNamePriceLicenseType

// This is an alias to an internal type.
type PriceGroupedTieredPackagePrice = shared.PriceGroupedTieredPackagePrice

// This is an alias to an internal type.
type PriceGroupedTieredPackagePriceBillingMode = shared.PriceGroupedTieredPackagePriceBillingMode

// This is an alias to an internal value.
const PriceGroupedTieredPackagePriceBillingModeInAdvance = shared.PriceGroupedTieredPackagePriceBillingModeInAdvance

// This is an alias to an internal value.
const PriceGroupedTieredPackagePriceBillingModeInArrear = shared.PriceGroupedTieredPackagePriceBillingModeInArrear

// This is an alias to an internal type.
type PriceGroupedTieredPackagePriceCadence = shared.PriceGroupedTieredPackagePriceCadence

// This is an alias to an internal value.
const PriceGroupedTieredPackagePriceCadenceOneTime = shared.PriceGroupedTieredPackagePriceCadenceOneTime

// This is an alias to an internal value.
const PriceGroupedTieredPackagePriceCadenceMonthly = shared.PriceGroupedTieredPackagePriceCadenceMonthly

// This is an alias to an internal value.
const PriceGroupedTieredPackagePriceCadenceQuarterly = shared.PriceGroupedTieredPackagePriceCadenceQuarterly

// This is an alias to an internal value.
const PriceGroupedTieredPackagePriceCadenceSemiAnnual = shared.PriceGroupedTieredPackagePriceCadenceSemiAnnual

// This is an alias to an internal value.
const PriceGroupedTieredPackagePriceCadenceAnnual = shared.PriceGroupedTieredPackagePriceCadenceAnnual

// This is an alias to an internal value.
const PriceGroupedTieredPackagePriceCadenceCustom = shared.PriceGroupedTieredPackagePriceCadenceCustom

// This is an alias to an internal type.
type PriceGroupedTieredPackagePriceCompositePriceFilter = shared.PriceGroupedTieredPackagePriceCompositePriceFilter

// The property of the price to filter on.
//
// This is an alias to an internal type.
type PriceGroupedTieredPackagePriceCompositePriceFiltersField = shared.PriceGroupedTieredPackagePriceCompositePriceFiltersField

// This is an alias to an internal value.
const PriceGroupedTieredPackagePriceCompositePriceFiltersFieldPriceID = shared.PriceGroupedTieredPackagePriceCompositePriceFiltersFieldPriceID

// This is an alias to an internal value.
const PriceGroupedTieredPackagePriceCompositePriceFiltersFieldItemID = shared.PriceGroupedTieredPackagePriceCompositePriceFiltersFieldItemID

// This is an alias to an internal value.
const PriceGroupedTieredPackagePriceCompositePriceFiltersFieldPriceType = shared.PriceGroupedTieredPackagePriceCompositePriceFiltersFieldPriceType

// This is an alias to an internal value.
const PriceGroupedTieredPackagePriceCompositePriceFiltersFieldCurrency = shared.PriceGroupedTieredPackagePriceCompositePriceFiltersFieldCurrency

// This is an alias to an internal value.
const PriceGroupedTieredPackagePriceCompositePriceFiltersFieldPricingUnitID = shared.PriceGroupedTieredPackagePriceCompositePriceFiltersFieldPricingUnitID

// Should prices that match the filter be included or excluded.
//
// This is an alias to an internal type.
type PriceGroupedTieredPackagePriceCompositePriceFiltersOperator = shared.PriceGroupedTieredPackagePriceCompositePriceFiltersOperator

// This is an alias to an internal value.
const PriceGroupedTieredPackagePriceCompositePriceFiltersOperatorIncludes = shared.PriceGroupedTieredPackagePriceCompositePriceFiltersOperatorIncludes

// This is an alias to an internal value.
const PriceGroupedTieredPackagePriceCompositePriceFiltersOperatorExcludes = shared.PriceGroupedTieredPackagePriceCompositePriceFiltersOperatorExcludes

// This is an alias to an internal type.
type PriceGroupedTieredPackagePriceConversionRateConfig = shared.PriceGroupedTieredPackagePriceConversionRateConfig

// This is an alias to an internal type.
type PriceGroupedTieredPackagePriceConversionRateConfigConversionRateType = shared.PriceGroupedTieredPackagePriceConversionRateConfigConversionRateType

// This is an alias to an internal value.
const PriceGroupedTieredPackagePriceConversionRateConfigConversionRateTypeUnit = shared.PriceGroupedTieredPackagePriceConversionRateConfigConversionRateTypeUnit

// This is an alias to an internal value.
const PriceGroupedTieredPackagePriceConversionRateConfigConversionRateTypeTiered = shared.PriceGroupedTieredPackagePriceConversionRateConfigConversionRateTypeTiered

// Configuration for grouped_tiered_package pricing
//
// This is an alias to an internal type.
type PriceGroupedTieredPackagePriceGroupedTieredPackageConfig = shared.PriceGroupedTieredPackagePriceGroupedTieredPackageConfig

// Configuration for a single tier
//
// This is an alias to an internal type.
type PriceGroupedTieredPackagePriceGroupedTieredPackageConfigTier = shared.PriceGroupedTieredPackagePriceGroupedTieredPackageConfigTier

// The pricing model type
//
// This is an alias to an internal type.
type PriceGroupedTieredPackagePriceModelType = shared.PriceGroupedTieredPackagePriceModelType

// This is an alias to an internal value.
const PriceGroupedTieredPackagePriceModelTypeGroupedTieredPackage = shared.PriceGroupedTieredPackagePriceModelTypeGroupedTieredPackage

// This is an alias to an internal type.
type PriceGroupedTieredPackagePricePriceType = shared.PriceGroupedTieredPackagePricePriceType

// This is an alias to an internal value.
const PriceGroupedTieredPackagePricePriceTypeUsagePrice = shared.PriceGroupedTieredPackagePricePriceTypeUsagePrice

// This is an alias to an internal value.
const PriceGroupedTieredPackagePricePriceTypeFixedPrice = shared.PriceGroupedTieredPackagePricePriceTypeFixedPrice

// This is an alias to an internal value.
const PriceGroupedTieredPackagePricePriceTypeCompositePrice = shared.PriceGroupedTieredPackagePricePriceTypeCompositePrice

// The LicenseType resource represents a type of license that can be assigned to
// users. License types are used during billing by grouping metrics on the
// configured grouping key.
//
// This is an alias to an internal type.
type PriceGroupedTieredPackagePriceLicenseType = shared.PriceGroupedTieredPackagePriceLicenseType

// This is an alias to an internal type.
type PriceMaxGroupTieredPackagePrice = shared.PriceMaxGroupTieredPackagePrice

// This is an alias to an internal type.
type PriceMaxGroupTieredPackagePriceBillingMode = shared.PriceMaxGroupTieredPackagePriceBillingMode

// This is an alias to an internal value.
const PriceMaxGroupTieredPackagePriceBillingModeInAdvance = shared.PriceMaxGroupTieredPackagePriceBillingModeInAdvance

// This is an alias to an internal value.
const PriceMaxGroupTieredPackagePriceBillingModeInArrear = shared.PriceMaxGroupTieredPackagePriceBillingModeInArrear

// This is an alias to an internal type.
type PriceMaxGroupTieredPackagePriceCadence = shared.PriceMaxGroupTieredPackagePriceCadence

// This is an alias to an internal value.
const PriceMaxGroupTieredPackagePriceCadenceOneTime = shared.PriceMaxGroupTieredPackagePriceCadenceOneTime

// This is an alias to an internal value.
const PriceMaxGroupTieredPackagePriceCadenceMonthly = shared.PriceMaxGroupTieredPackagePriceCadenceMonthly

// This is an alias to an internal value.
const PriceMaxGroupTieredPackagePriceCadenceQuarterly = shared.PriceMaxGroupTieredPackagePriceCadenceQuarterly

// This is an alias to an internal value.
const PriceMaxGroupTieredPackagePriceCadenceSemiAnnual = shared.PriceMaxGroupTieredPackagePriceCadenceSemiAnnual

// This is an alias to an internal value.
const PriceMaxGroupTieredPackagePriceCadenceAnnual = shared.PriceMaxGroupTieredPackagePriceCadenceAnnual

// This is an alias to an internal value.
const PriceMaxGroupTieredPackagePriceCadenceCustom = shared.PriceMaxGroupTieredPackagePriceCadenceCustom

// This is an alias to an internal type.
type PriceMaxGroupTieredPackagePriceCompositePriceFilter = shared.PriceMaxGroupTieredPackagePriceCompositePriceFilter

// The property of the price to filter on.
//
// This is an alias to an internal type.
type PriceMaxGroupTieredPackagePriceCompositePriceFiltersField = shared.PriceMaxGroupTieredPackagePriceCompositePriceFiltersField

// This is an alias to an internal value.
const PriceMaxGroupTieredPackagePriceCompositePriceFiltersFieldPriceID = shared.PriceMaxGroupTieredPackagePriceCompositePriceFiltersFieldPriceID

// This is an alias to an internal value.
const PriceMaxGroupTieredPackagePriceCompositePriceFiltersFieldItemID = shared.PriceMaxGroupTieredPackagePriceCompositePriceFiltersFieldItemID

// This is an alias to an internal value.
const PriceMaxGroupTieredPackagePriceCompositePriceFiltersFieldPriceType = shared.PriceMaxGroupTieredPackagePriceCompositePriceFiltersFieldPriceType

// This is an alias to an internal value.
const PriceMaxGroupTieredPackagePriceCompositePriceFiltersFieldCurrency = shared.PriceMaxGroupTieredPackagePriceCompositePriceFiltersFieldCurrency

// This is an alias to an internal value.
const PriceMaxGroupTieredPackagePriceCompositePriceFiltersFieldPricingUnitID = shared.PriceMaxGroupTieredPackagePriceCompositePriceFiltersFieldPricingUnitID

// Should prices that match the filter be included or excluded.
//
// This is an alias to an internal type.
type PriceMaxGroupTieredPackagePriceCompositePriceFiltersOperator = shared.PriceMaxGroupTieredPackagePriceCompositePriceFiltersOperator

// This is an alias to an internal value.
const PriceMaxGroupTieredPackagePriceCompositePriceFiltersOperatorIncludes = shared.PriceMaxGroupTieredPackagePriceCompositePriceFiltersOperatorIncludes

// This is an alias to an internal value.
const PriceMaxGroupTieredPackagePriceCompositePriceFiltersOperatorExcludes = shared.PriceMaxGroupTieredPackagePriceCompositePriceFiltersOperatorExcludes

// This is an alias to an internal type.
type PriceMaxGroupTieredPackagePriceConversionRateConfig = shared.PriceMaxGroupTieredPackagePriceConversionRateConfig

// This is an alias to an internal type.
type PriceMaxGroupTieredPackagePriceConversionRateConfigConversionRateType = shared.PriceMaxGroupTieredPackagePriceConversionRateConfigConversionRateType

// This is an alias to an internal value.
const PriceMaxGroupTieredPackagePriceConversionRateConfigConversionRateTypeUnit = shared.PriceMaxGroupTieredPackagePriceConversionRateConfigConversionRateTypeUnit

// This is an alias to an internal value.
const PriceMaxGroupTieredPackagePriceConversionRateConfigConversionRateTypeTiered = shared.PriceMaxGroupTieredPackagePriceConversionRateConfigConversionRateTypeTiered

// Configuration for max_group_tiered_package pricing
//
// This is an alias to an internal type.
type PriceMaxGroupTieredPackagePriceMaxGroupTieredPackageConfig = shared.PriceMaxGroupTieredPackagePriceMaxGroupTieredPackageConfig

// Configuration for a single tier
//
// This is an alias to an internal type.
type PriceMaxGroupTieredPackagePriceMaxGroupTieredPackageConfigTier = shared.PriceMaxGroupTieredPackagePriceMaxGroupTieredPackageConfigTier

// The pricing model type
//
// This is an alias to an internal type.
type PriceMaxGroupTieredPackagePriceModelType = shared.PriceMaxGroupTieredPackagePriceModelType

// This is an alias to an internal value.
const PriceMaxGroupTieredPackagePriceModelTypeMaxGroupTieredPackage = shared.PriceMaxGroupTieredPackagePriceModelTypeMaxGroupTieredPackage

// This is an alias to an internal type.
type PriceMaxGroupTieredPackagePricePriceType = shared.PriceMaxGroupTieredPackagePricePriceType

// This is an alias to an internal value.
const PriceMaxGroupTieredPackagePricePriceTypeUsagePrice = shared.PriceMaxGroupTieredPackagePricePriceTypeUsagePrice

// This is an alias to an internal value.
const PriceMaxGroupTieredPackagePricePriceTypeFixedPrice = shared.PriceMaxGroupTieredPackagePricePriceTypeFixedPrice

// This is an alias to an internal value.
const PriceMaxGroupTieredPackagePricePriceTypeCompositePrice = shared.PriceMaxGroupTieredPackagePricePriceTypeCompositePrice

// The LicenseType resource represents a type of license that can be assigned to
// users. License types are used during billing by grouping metrics on the
// configured grouping key.
//
// This is an alias to an internal type.
type PriceMaxGroupTieredPackagePriceLicenseType = shared.PriceMaxGroupTieredPackagePriceLicenseType

// This is an alias to an internal type.
type PriceScalableMatrixWithUnitPricingPrice = shared.PriceScalableMatrixWithUnitPricingPrice

// This is an alias to an internal type.
type PriceScalableMatrixWithUnitPricingPriceBillingMode = shared.PriceScalableMatrixWithUnitPricingPriceBillingMode

// This is an alias to an internal value.
const PriceScalableMatrixWithUnitPricingPriceBillingModeInAdvance = shared.PriceScalableMatrixWithUnitPricingPriceBillingModeInAdvance

// This is an alias to an internal value.
const PriceScalableMatrixWithUnitPricingPriceBillingModeInArrear = shared.PriceScalableMatrixWithUnitPricingPriceBillingModeInArrear

// This is an alias to an internal type.
type PriceScalableMatrixWithUnitPricingPriceCadence = shared.PriceScalableMatrixWithUnitPricingPriceCadence

// This is an alias to an internal value.
const PriceScalableMatrixWithUnitPricingPriceCadenceOneTime = shared.PriceScalableMatrixWithUnitPricingPriceCadenceOneTime

// This is an alias to an internal value.
const PriceScalableMatrixWithUnitPricingPriceCadenceMonthly = shared.PriceScalableMatrixWithUnitPricingPriceCadenceMonthly

// This is an alias to an internal value.
const PriceScalableMatrixWithUnitPricingPriceCadenceQuarterly = shared.PriceScalableMatrixWithUnitPricingPriceCadenceQuarterly

// This is an alias to an internal value.
const PriceScalableMatrixWithUnitPricingPriceCadenceSemiAnnual = shared.PriceScalableMatrixWithUnitPricingPriceCadenceSemiAnnual

// This is an alias to an internal value.
const PriceScalableMatrixWithUnitPricingPriceCadenceAnnual = shared.PriceScalableMatrixWithUnitPricingPriceCadenceAnnual

// This is an alias to an internal value.
const PriceScalableMatrixWithUnitPricingPriceCadenceCustom = shared.PriceScalableMatrixWithUnitPricingPriceCadenceCustom

// This is an alias to an internal type.
type PriceScalableMatrixWithUnitPricingPriceCompositePriceFilter = shared.PriceScalableMatrixWithUnitPricingPriceCompositePriceFilter

// The property of the price to filter on.
//
// This is an alias to an internal type.
type PriceScalableMatrixWithUnitPricingPriceCompositePriceFiltersField = shared.PriceScalableMatrixWithUnitPricingPriceCompositePriceFiltersField

// This is an alias to an internal value.
const PriceScalableMatrixWithUnitPricingPriceCompositePriceFiltersFieldPriceID = shared.PriceScalableMatrixWithUnitPricingPriceCompositePriceFiltersFieldPriceID

// This is an alias to an internal value.
const PriceScalableMatrixWithUnitPricingPriceCompositePriceFiltersFieldItemID = shared.PriceScalableMatrixWithUnitPricingPriceCompositePriceFiltersFieldItemID

// This is an alias to an internal value.
const PriceScalableMatrixWithUnitPricingPriceCompositePriceFiltersFieldPriceType = shared.PriceScalableMatrixWithUnitPricingPriceCompositePriceFiltersFieldPriceType

// This is an alias to an internal value.
const PriceScalableMatrixWithUnitPricingPriceCompositePriceFiltersFieldCurrency = shared.PriceScalableMatrixWithUnitPricingPriceCompositePriceFiltersFieldCurrency

// This is an alias to an internal value.
const PriceScalableMatrixWithUnitPricingPriceCompositePriceFiltersFieldPricingUnitID = shared.PriceScalableMatrixWithUnitPricingPriceCompositePriceFiltersFieldPricingUnitID

// Should prices that match the filter be included or excluded.
//
// This is an alias to an internal type.
type PriceScalableMatrixWithUnitPricingPriceCompositePriceFiltersOperator = shared.PriceScalableMatrixWithUnitPricingPriceCompositePriceFiltersOperator

// This is an alias to an internal value.
const PriceScalableMatrixWithUnitPricingPriceCompositePriceFiltersOperatorIncludes = shared.PriceScalableMatrixWithUnitPricingPriceCompositePriceFiltersOperatorIncludes

// This is an alias to an internal value.
const PriceScalableMatrixWithUnitPricingPriceCompositePriceFiltersOperatorExcludes = shared.PriceScalableMatrixWithUnitPricingPriceCompositePriceFiltersOperatorExcludes

// This is an alias to an internal type.
type PriceScalableMatrixWithUnitPricingPriceConversionRateConfig = shared.PriceScalableMatrixWithUnitPricingPriceConversionRateConfig

// This is an alias to an internal type.
type PriceScalableMatrixWithUnitPricingPriceConversionRateConfigConversionRateType = shared.PriceScalableMatrixWithUnitPricingPriceConversionRateConfigConversionRateType

// This is an alias to an internal value.
const PriceScalableMatrixWithUnitPricingPriceConversionRateConfigConversionRateTypeUnit = shared.PriceScalableMatrixWithUnitPricingPriceConversionRateConfigConversionRateTypeUnit

// This is an alias to an internal value.
const PriceScalableMatrixWithUnitPricingPriceConversionRateConfigConversionRateTypeTiered = shared.PriceScalableMatrixWithUnitPricingPriceConversionRateConfigConversionRateTypeTiered

// The pricing model type
//
// This is an alias to an internal type.
type PriceScalableMatrixWithUnitPricingPriceModelType = shared.PriceScalableMatrixWithUnitPricingPriceModelType

// This is an alias to an internal value.
const PriceScalableMatrixWithUnitPricingPriceModelTypeScalableMatrixWithUnitPricing = shared.PriceScalableMatrixWithUnitPricingPriceModelTypeScalableMatrixWithUnitPricing

// This is an alias to an internal type.
type PriceScalableMatrixWithUnitPricingPricePriceType = shared.PriceScalableMatrixWithUnitPricingPricePriceType

// This is an alias to an internal value.
const PriceScalableMatrixWithUnitPricingPricePriceTypeUsagePrice = shared.PriceScalableMatrixWithUnitPricingPricePriceTypeUsagePrice

// This is an alias to an internal value.
const PriceScalableMatrixWithUnitPricingPricePriceTypeFixedPrice = shared.PriceScalableMatrixWithUnitPricingPricePriceTypeFixedPrice

// This is an alias to an internal value.
const PriceScalableMatrixWithUnitPricingPricePriceTypeCompositePrice = shared.PriceScalableMatrixWithUnitPricingPricePriceTypeCompositePrice

// Configuration for scalable_matrix_with_unit_pricing pricing
//
// This is an alias to an internal type.
type PriceScalableMatrixWithUnitPricingPriceScalableMatrixWithUnitPricingConfig = shared.PriceScalableMatrixWithUnitPricingPriceScalableMatrixWithUnitPricingConfig

// Configuration for a single matrix scaling factor
//
// This is an alias to an internal type.
type PriceScalableMatrixWithUnitPricingPriceScalableMatrixWithUnitPricingConfigMatrixScalingFactor = shared.PriceScalableMatrixWithUnitPricingPriceScalableMatrixWithUnitPricingConfigMatrixScalingFactor

// The LicenseType resource represents a type of license that can be assigned to
// users. License types are used during billing by grouping metrics on the
// configured grouping key.
//
// This is an alias to an internal type.
type PriceScalableMatrixWithUnitPricingPriceLicenseType = shared.PriceScalableMatrixWithUnitPricingPriceLicenseType

// This is an alias to an internal type.
type PriceScalableMatrixWithTieredPricingPrice = shared.PriceScalableMatrixWithTieredPricingPrice

// This is an alias to an internal type.
type PriceScalableMatrixWithTieredPricingPriceBillingMode = shared.PriceScalableMatrixWithTieredPricingPriceBillingMode

// This is an alias to an internal value.
const PriceScalableMatrixWithTieredPricingPriceBillingModeInAdvance = shared.PriceScalableMatrixWithTieredPricingPriceBillingModeInAdvance

// This is an alias to an internal value.
const PriceScalableMatrixWithTieredPricingPriceBillingModeInArrear = shared.PriceScalableMatrixWithTieredPricingPriceBillingModeInArrear

// This is an alias to an internal type.
type PriceScalableMatrixWithTieredPricingPriceCadence = shared.PriceScalableMatrixWithTieredPricingPriceCadence

// This is an alias to an internal value.
const PriceScalableMatrixWithTieredPricingPriceCadenceOneTime = shared.PriceScalableMatrixWithTieredPricingPriceCadenceOneTime

// This is an alias to an internal value.
const PriceScalableMatrixWithTieredPricingPriceCadenceMonthly = shared.PriceScalableMatrixWithTieredPricingPriceCadenceMonthly

// This is an alias to an internal value.
const PriceScalableMatrixWithTieredPricingPriceCadenceQuarterly = shared.PriceScalableMatrixWithTieredPricingPriceCadenceQuarterly

// This is an alias to an internal value.
const PriceScalableMatrixWithTieredPricingPriceCadenceSemiAnnual = shared.PriceScalableMatrixWithTieredPricingPriceCadenceSemiAnnual

// This is an alias to an internal value.
const PriceScalableMatrixWithTieredPricingPriceCadenceAnnual = shared.PriceScalableMatrixWithTieredPricingPriceCadenceAnnual

// This is an alias to an internal value.
const PriceScalableMatrixWithTieredPricingPriceCadenceCustom = shared.PriceScalableMatrixWithTieredPricingPriceCadenceCustom

// This is an alias to an internal type.
type PriceScalableMatrixWithTieredPricingPriceCompositePriceFilter = shared.PriceScalableMatrixWithTieredPricingPriceCompositePriceFilter

// The property of the price to filter on.
//
// This is an alias to an internal type.
type PriceScalableMatrixWithTieredPricingPriceCompositePriceFiltersField = shared.PriceScalableMatrixWithTieredPricingPriceCompositePriceFiltersField

// This is an alias to an internal value.
const PriceScalableMatrixWithTieredPricingPriceCompositePriceFiltersFieldPriceID = shared.PriceScalableMatrixWithTieredPricingPriceCompositePriceFiltersFieldPriceID

// This is an alias to an internal value.
const PriceScalableMatrixWithTieredPricingPriceCompositePriceFiltersFieldItemID = shared.PriceScalableMatrixWithTieredPricingPriceCompositePriceFiltersFieldItemID

// This is an alias to an internal value.
const PriceScalableMatrixWithTieredPricingPriceCompositePriceFiltersFieldPriceType = shared.PriceScalableMatrixWithTieredPricingPriceCompositePriceFiltersFieldPriceType

// This is an alias to an internal value.
const PriceScalableMatrixWithTieredPricingPriceCompositePriceFiltersFieldCurrency = shared.PriceScalableMatrixWithTieredPricingPriceCompositePriceFiltersFieldCurrency

// This is an alias to an internal value.
const PriceScalableMatrixWithTieredPricingPriceCompositePriceFiltersFieldPricingUnitID = shared.PriceScalableMatrixWithTieredPricingPriceCompositePriceFiltersFieldPricingUnitID

// Should prices that match the filter be included or excluded.
//
// This is an alias to an internal type.
type PriceScalableMatrixWithTieredPricingPriceCompositePriceFiltersOperator = shared.PriceScalableMatrixWithTieredPricingPriceCompositePriceFiltersOperator

// This is an alias to an internal value.
const PriceScalableMatrixWithTieredPricingPriceCompositePriceFiltersOperatorIncludes = shared.PriceScalableMatrixWithTieredPricingPriceCompositePriceFiltersOperatorIncludes

// This is an alias to an internal value.
const PriceScalableMatrixWithTieredPricingPriceCompositePriceFiltersOperatorExcludes = shared.PriceScalableMatrixWithTieredPricingPriceCompositePriceFiltersOperatorExcludes

// This is an alias to an internal type.
type PriceScalableMatrixWithTieredPricingPriceConversionRateConfig = shared.PriceScalableMatrixWithTieredPricingPriceConversionRateConfig

// This is an alias to an internal type.
type PriceScalableMatrixWithTieredPricingPriceConversionRateConfigConversionRateType = shared.PriceScalableMatrixWithTieredPricingPriceConversionRateConfigConversionRateType

// This is an alias to an internal value.
const PriceScalableMatrixWithTieredPricingPriceConversionRateConfigConversionRateTypeUnit = shared.PriceScalableMatrixWithTieredPricingPriceConversionRateConfigConversionRateTypeUnit

// This is an alias to an internal value.
const PriceScalableMatrixWithTieredPricingPriceConversionRateConfigConversionRateTypeTiered = shared.PriceScalableMatrixWithTieredPricingPriceConversionRateConfigConversionRateTypeTiered

// The pricing model type
//
// This is an alias to an internal type.
type PriceScalableMatrixWithTieredPricingPriceModelType = shared.PriceScalableMatrixWithTieredPricingPriceModelType

// This is an alias to an internal value.
const PriceScalableMatrixWithTieredPricingPriceModelTypeScalableMatrixWithTieredPricing = shared.PriceScalableMatrixWithTieredPricingPriceModelTypeScalableMatrixWithTieredPricing

// This is an alias to an internal type.
type PriceScalableMatrixWithTieredPricingPricePriceType = shared.PriceScalableMatrixWithTieredPricingPricePriceType

// This is an alias to an internal value.
const PriceScalableMatrixWithTieredPricingPricePriceTypeUsagePrice = shared.PriceScalableMatrixWithTieredPricingPricePriceTypeUsagePrice

// This is an alias to an internal value.
const PriceScalableMatrixWithTieredPricingPricePriceTypeFixedPrice = shared.PriceScalableMatrixWithTieredPricingPricePriceTypeFixedPrice

// This is an alias to an internal value.
const PriceScalableMatrixWithTieredPricingPricePriceTypeCompositePrice = shared.PriceScalableMatrixWithTieredPricingPricePriceTypeCompositePrice

// Configuration for scalable_matrix_with_tiered_pricing pricing
//
// This is an alias to an internal type.
type PriceScalableMatrixWithTieredPricingPriceScalableMatrixWithTieredPricingConfig = shared.PriceScalableMatrixWithTieredPricingPriceScalableMatrixWithTieredPricingConfig

// Configuration for a single matrix scaling factor
//
// This is an alias to an internal type.
type PriceScalableMatrixWithTieredPricingPriceScalableMatrixWithTieredPricingConfigMatrixScalingFactor = shared.PriceScalableMatrixWithTieredPricingPriceScalableMatrixWithTieredPricingConfigMatrixScalingFactor

// Configuration for a single tier entry with business logic
//
// This is an alias to an internal type.
type PriceScalableMatrixWithTieredPricingPriceScalableMatrixWithTieredPricingConfigTier = shared.PriceScalableMatrixWithTieredPricingPriceScalableMatrixWithTieredPricingConfigTier

// The LicenseType resource represents a type of license that can be assigned to
// users. License types are used during billing by grouping metrics on the
// configured grouping key.
//
// This is an alias to an internal type.
type PriceScalableMatrixWithTieredPricingPriceLicenseType = shared.PriceScalableMatrixWithTieredPricingPriceLicenseType

// This is an alias to an internal type.
type PriceCumulativeGroupedBulkPrice = shared.PriceCumulativeGroupedBulkPrice

// This is an alias to an internal type.
type PriceCumulativeGroupedBulkPriceBillingMode = shared.PriceCumulativeGroupedBulkPriceBillingMode

// This is an alias to an internal value.
const PriceCumulativeGroupedBulkPriceBillingModeInAdvance = shared.PriceCumulativeGroupedBulkPriceBillingModeInAdvance

// This is an alias to an internal value.
const PriceCumulativeGroupedBulkPriceBillingModeInArrear = shared.PriceCumulativeGroupedBulkPriceBillingModeInArrear

// This is an alias to an internal type.
type PriceCumulativeGroupedBulkPriceCadence = shared.PriceCumulativeGroupedBulkPriceCadence

// This is an alias to an internal value.
const PriceCumulativeGroupedBulkPriceCadenceOneTime = shared.PriceCumulativeGroupedBulkPriceCadenceOneTime

// This is an alias to an internal value.
const PriceCumulativeGroupedBulkPriceCadenceMonthly = shared.PriceCumulativeGroupedBulkPriceCadenceMonthly

// This is an alias to an internal value.
const PriceCumulativeGroupedBulkPriceCadenceQuarterly = shared.PriceCumulativeGroupedBulkPriceCadenceQuarterly

// This is an alias to an internal value.
const PriceCumulativeGroupedBulkPriceCadenceSemiAnnual = shared.PriceCumulativeGroupedBulkPriceCadenceSemiAnnual

// This is an alias to an internal value.
const PriceCumulativeGroupedBulkPriceCadenceAnnual = shared.PriceCumulativeGroupedBulkPriceCadenceAnnual

// This is an alias to an internal value.
const PriceCumulativeGroupedBulkPriceCadenceCustom = shared.PriceCumulativeGroupedBulkPriceCadenceCustom

// This is an alias to an internal type.
type PriceCumulativeGroupedBulkPriceCompositePriceFilter = shared.PriceCumulativeGroupedBulkPriceCompositePriceFilter

// The property of the price to filter on.
//
// This is an alias to an internal type.
type PriceCumulativeGroupedBulkPriceCompositePriceFiltersField = shared.PriceCumulativeGroupedBulkPriceCompositePriceFiltersField

// This is an alias to an internal value.
const PriceCumulativeGroupedBulkPriceCompositePriceFiltersFieldPriceID = shared.PriceCumulativeGroupedBulkPriceCompositePriceFiltersFieldPriceID

// This is an alias to an internal value.
const PriceCumulativeGroupedBulkPriceCompositePriceFiltersFieldItemID = shared.PriceCumulativeGroupedBulkPriceCompositePriceFiltersFieldItemID

// This is an alias to an internal value.
const PriceCumulativeGroupedBulkPriceCompositePriceFiltersFieldPriceType = shared.PriceCumulativeGroupedBulkPriceCompositePriceFiltersFieldPriceType

// This is an alias to an internal value.
const PriceCumulativeGroupedBulkPriceCompositePriceFiltersFieldCurrency = shared.PriceCumulativeGroupedBulkPriceCompositePriceFiltersFieldCurrency

// This is an alias to an internal value.
const PriceCumulativeGroupedBulkPriceCompositePriceFiltersFieldPricingUnitID = shared.PriceCumulativeGroupedBulkPriceCompositePriceFiltersFieldPricingUnitID

// Should prices that match the filter be included or excluded.
//
// This is an alias to an internal type.
type PriceCumulativeGroupedBulkPriceCompositePriceFiltersOperator = shared.PriceCumulativeGroupedBulkPriceCompositePriceFiltersOperator

// This is an alias to an internal value.
const PriceCumulativeGroupedBulkPriceCompositePriceFiltersOperatorIncludes = shared.PriceCumulativeGroupedBulkPriceCompositePriceFiltersOperatorIncludes

// This is an alias to an internal value.
const PriceCumulativeGroupedBulkPriceCompositePriceFiltersOperatorExcludes = shared.PriceCumulativeGroupedBulkPriceCompositePriceFiltersOperatorExcludes

// This is an alias to an internal type.
type PriceCumulativeGroupedBulkPriceConversionRateConfig = shared.PriceCumulativeGroupedBulkPriceConversionRateConfig

// This is an alias to an internal type.
type PriceCumulativeGroupedBulkPriceConversionRateConfigConversionRateType = shared.PriceCumulativeGroupedBulkPriceConversionRateConfigConversionRateType

// This is an alias to an internal value.
const PriceCumulativeGroupedBulkPriceConversionRateConfigConversionRateTypeUnit = shared.PriceCumulativeGroupedBulkPriceConversionRateConfigConversionRateTypeUnit

// This is an alias to an internal value.
const PriceCumulativeGroupedBulkPriceConversionRateConfigConversionRateTypeTiered = shared.PriceCumulativeGroupedBulkPriceConversionRateConfigConversionRateTypeTiered

// Configuration for cumulative_grouped_bulk pricing
//
// This is an alias to an internal type.
type PriceCumulativeGroupedBulkPriceCumulativeGroupedBulkConfig = shared.PriceCumulativeGroupedBulkPriceCumulativeGroupedBulkConfig

// Configuration for a dimension value entry
//
// This is an alias to an internal type.
type PriceCumulativeGroupedBulkPriceCumulativeGroupedBulkConfigDimensionValue = shared.PriceCumulativeGroupedBulkPriceCumulativeGroupedBulkConfigDimensionValue

// The pricing model type
//
// This is an alias to an internal type.
type PriceCumulativeGroupedBulkPriceModelType = shared.PriceCumulativeGroupedBulkPriceModelType

// This is an alias to an internal value.
const PriceCumulativeGroupedBulkPriceModelTypeCumulativeGroupedBulk = shared.PriceCumulativeGroupedBulkPriceModelTypeCumulativeGroupedBulk

// This is an alias to an internal type.
type PriceCumulativeGroupedBulkPricePriceType = shared.PriceCumulativeGroupedBulkPricePriceType

// This is an alias to an internal value.
const PriceCumulativeGroupedBulkPricePriceTypeUsagePrice = shared.PriceCumulativeGroupedBulkPricePriceTypeUsagePrice

// This is an alias to an internal value.
const PriceCumulativeGroupedBulkPricePriceTypeFixedPrice = shared.PriceCumulativeGroupedBulkPricePriceTypeFixedPrice

// This is an alias to an internal value.
const PriceCumulativeGroupedBulkPricePriceTypeCompositePrice = shared.PriceCumulativeGroupedBulkPricePriceTypeCompositePrice

// The LicenseType resource represents a type of license that can be assigned to
// users. License types are used during billing by grouping metrics on the
// configured grouping key.
//
// This is an alias to an internal type.
type PriceCumulativeGroupedBulkPriceLicenseType = shared.PriceCumulativeGroupedBulkPriceLicenseType

// This is an alias to an internal type.
type PriceCumulativeGroupedAllocationPrice = shared.PriceCumulativeGroupedAllocationPrice

// This is an alias to an internal type.
type PriceCumulativeGroupedAllocationPriceBillingMode = shared.PriceCumulativeGroupedAllocationPriceBillingMode

// This is an alias to an internal value.
const PriceCumulativeGroupedAllocationPriceBillingModeInAdvance = shared.PriceCumulativeGroupedAllocationPriceBillingModeInAdvance

// This is an alias to an internal value.
const PriceCumulativeGroupedAllocationPriceBillingModeInArrear = shared.PriceCumulativeGroupedAllocationPriceBillingModeInArrear

// This is an alias to an internal type.
type PriceCumulativeGroupedAllocationPriceCadence = shared.PriceCumulativeGroupedAllocationPriceCadence

// This is an alias to an internal value.
const PriceCumulativeGroupedAllocationPriceCadenceOneTime = shared.PriceCumulativeGroupedAllocationPriceCadenceOneTime

// This is an alias to an internal value.
const PriceCumulativeGroupedAllocationPriceCadenceMonthly = shared.PriceCumulativeGroupedAllocationPriceCadenceMonthly

// This is an alias to an internal value.
const PriceCumulativeGroupedAllocationPriceCadenceQuarterly = shared.PriceCumulativeGroupedAllocationPriceCadenceQuarterly

// This is an alias to an internal value.
const PriceCumulativeGroupedAllocationPriceCadenceSemiAnnual = shared.PriceCumulativeGroupedAllocationPriceCadenceSemiAnnual

// This is an alias to an internal value.
const PriceCumulativeGroupedAllocationPriceCadenceAnnual = shared.PriceCumulativeGroupedAllocationPriceCadenceAnnual

// This is an alias to an internal value.
const PriceCumulativeGroupedAllocationPriceCadenceCustom = shared.PriceCumulativeGroupedAllocationPriceCadenceCustom

// This is an alias to an internal type.
type PriceCumulativeGroupedAllocationPriceCompositePriceFilter = shared.PriceCumulativeGroupedAllocationPriceCompositePriceFilter

// The property of the price to filter on.
//
// This is an alias to an internal type.
type PriceCumulativeGroupedAllocationPriceCompositePriceFiltersField = shared.PriceCumulativeGroupedAllocationPriceCompositePriceFiltersField

// This is an alias to an internal value.
const PriceCumulativeGroupedAllocationPriceCompositePriceFiltersFieldPriceID = shared.PriceCumulativeGroupedAllocationPriceCompositePriceFiltersFieldPriceID

// This is an alias to an internal value.
const PriceCumulativeGroupedAllocationPriceCompositePriceFiltersFieldItemID = shared.PriceCumulativeGroupedAllocationPriceCompositePriceFiltersFieldItemID

// This is an alias to an internal value.
const PriceCumulativeGroupedAllocationPriceCompositePriceFiltersFieldPriceType = shared.PriceCumulativeGroupedAllocationPriceCompositePriceFiltersFieldPriceType

// This is an alias to an internal value.
const PriceCumulativeGroupedAllocationPriceCompositePriceFiltersFieldCurrency = shared.PriceCumulativeGroupedAllocationPriceCompositePriceFiltersFieldCurrency

// This is an alias to an internal value.
const PriceCumulativeGroupedAllocationPriceCompositePriceFiltersFieldPricingUnitID = shared.PriceCumulativeGroupedAllocationPriceCompositePriceFiltersFieldPricingUnitID

// Should prices that match the filter be included or excluded.
//
// This is an alias to an internal type.
type PriceCumulativeGroupedAllocationPriceCompositePriceFiltersOperator = shared.PriceCumulativeGroupedAllocationPriceCompositePriceFiltersOperator

// This is an alias to an internal value.
const PriceCumulativeGroupedAllocationPriceCompositePriceFiltersOperatorIncludes = shared.PriceCumulativeGroupedAllocationPriceCompositePriceFiltersOperatorIncludes

// This is an alias to an internal value.
const PriceCumulativeGroupedAllocationPriceCompositePriceFiltersOperatorExcludes = shared.PriceCumulativeGroupedAllocationPriceCompositePriceFiltersOperatorExcludes

// This is an alias to an internal type.
type PriceCumulativeGroupedAllocationPriceConversionRateConfig = shared.PriceCumulativeGroupedAllocationPriceConversionRateConfig

// This is an alias to an internal type.
type PriceCumulativeGroupedAllocationPriceConversionRateConfigConversionRateType = shared.PriceCumulativeGroupedAllocationPriceConversionRateConfigConversionRateType

// This is an alias to an internal value.
const PriceCumulativeGroupedAllocationPriceConversionRateConfigConversionRateTypeUnit = shared.PriceCumulativeGroupedAllocationPriceConversionRateConfigConversionRateTypeUnit

// This is an alias to an internal value.
const PriceCumulativeGroupedAllocationPriceConversionRateConfigConversionRateTypeTiered = shared.PriceCumulativeGroupedAllocationPriceConversionRateConfigConversionRateTypeTiered

// Configuration for cumulative_grouped_allocation pricing
//
// This is an alias to an internal type.
type PriceCumulativeGroupedAllocationPriceCumulativeGroupedAllocationConfig = shared.PriceCumulativeGroupedAllocationPriceCumulativeGroupedAllocationConfig

// The pricing model type
//
// This is an alias to an internal type.
type PriceCumulativeGroupedAllocationPriceModelType = shared.PriceCumulativeGroupedAllocationPriceModelType

// This is an alias to an internal value.
const PriceCumulativeGroupedAllocationPriceModelTypeCumulativeGroupedAllocation = shared.PriceCumulativeGroupedAllocationPriceModelTypeCumulativeGroupedAllocation

// This is an alias to an internal type.
type PriceCumulativeGroupedAllocationPricePriceType = shared.PriceCumulativeGroupedAllocationPricePriceType

// This is an alias to an internal value.
const PriceCumulativeGroupedAllocationPricePriceTypeUsagePrice = shared.PriceCumulativeGroupedAllocationPricePriceTypeUsagePrice

// This is an alias to an internal value.
const PriceCumulativeGroupedAllocationPricePriceTypeFixedPrice = shared.PriceCumulativeGroupedAllocationPricePriceTypeFixedPrice

// This is an alias to an internal value.
const PriceCumulativeGroupedAllocationPricePriceTypeCompositePrice = shared.PriceCumulativeGroupedAllocationPricePriceTypeCompositePrice

// The LicenseType resource represents a type of license that can be assigned to
// users. License types are used during billing by grouping metrics on the
// configured grouping key.
//
// This is an alias to an internal type.
type PriceCumulativeGroupedAllocationPriceLicenseType = shared.PriceCumulativeGroupedAllocationPriceLicenseType

// This is an alias to an internal type.
type PriceMinimumCompositePrice = shared.PriceMinimumCompositePrice

// This is an alias to an internal type.
type PriceMinimumCompositePriceBillingMode = shared.PriceMinimumCompositePriceBillingMode

// This is an alias to an internal value.
const PriceMinimumCompositePriceBillingModeInAdvance = shared.PriceMinimumCompositePriceBillingModeInAdvance

// This is an alias to an internal value.
const PriceMinimumCompositePriceBillingModeInArrear = shared.PriceMinimumCompositePriceBillingModeInArrear

// This is an alias to an internal type.
type PriceMinimumCompositePriceCadence = shared.PriceMinimumCompositePriceCadence

// This is an alias to an internal value.
const PriceMinimumCompositePriceCadenceOneTime = shared.PriceMinimumCompositePriceCadenceOneTime

// This is an alias to an internal value.
const PriceMinimumCompositePriceCadenceMonthly = shared.PriceMinimumCompositePriceCadenceMonthly

// This is an alias to an internal value.
const PriceMinimumCompositePriceCadenceQuarterly = shared.PriceMinimumCompositePriceCadenceQuarterly

// This is an alias to an internal value.
const PriceMinimumCompositePriceCadenceSemiAnnual = shared.PriceMinimumCompositePriceCadenceSemiAnnual

// This is an alias to an internal value.
const PriceMinimumCompositePriceCadenceAnnual = shared.PriceMinimumCompositePriceCadenceAnnual

// This is an alias to an internal value.
const PriceMinimumCompositePriceCadenceCustom = shared.PriceMinimumCompositePriceCadenceCustom

// This is an alias to an internal type.
type PriceMinimumCompositePriceCompositePriceFilter = shared.PriceMinimumCompositePriceCompositePriceFilter

// The property of the price to filter on.
//
// This is an alias to an internal type.
type PriceMinimumCompositePriceCompositePriceFiltersField = shared.PriceMinimumCompositePriceCompositePriceFiltersField

// This is an alias to an internal value.
const PriceMinimumCompositePriceCompositePriceFiltersFieldPriceID = shared.PriceMinimumCompositePriceCompositePriceFiltersFieldPriceID

// This is an alias to an internal value.
const PriceMinimumCompositePriceCompositePriceFiltersFieldItemID = shared.PriceMinimumCompositePriceCompositePriceFiltersFieldItemID

// This is an alias to an internal value.
const PriceMinimumCompositePriceCompositePriceFiltersFieldPriceType = shared.PriceMinimumCompositePriceCompositePriceFiltersFieldPriceType

// This is an alias to an internal value.
const PriceMinimumCompositePriceCompositePriceFiltersFieldCurrency = shared.PriceMinimumCompositePriceCompositePriceFiltersFieldCurrency

// This is an alias to an internal value.
const PriceMinimumCompositePriceCompositePriceFiltersFieldPricingUnitID = shared.PriceMinimumCompositePriceCompositePriceFiltersFieldPricingUnitID

// Should prices that match the filter be included or excluded.
//
// This is an alias to an internal type.
type PriceMinimumCompositePriceCompositePriceFiltersOperator = shared.PriceMinimumCompositePriceCompositePriceFiltersOperator

// This is an alias to an internal value.
const PriceMinimumCompositePriceCompositePriceFiltersOperatorIncludes = shared.PriceMinimumCompositePriceCompositePriceFiltersOperatorIncludes

// This is an alias to an internal value.
const PriceMinimumCompositePriceCompositePriceFiltersOperatorExcludes = shared.PriceMinimumCompositePriceCompositePriceFiltersOperatorExcludes

// This is an alias to an internal type.
type PriceMinimumCompositePriceConversionRateConfig = shared.PriceMinimumCompositePriceConversionRateConfig

// This is an alias to an internal type.
type PriceMinimumCompositePriceConversionRateConfigConversionRateType = shared.PriceMinimumCompositePriceConversionRateConfigConversionRateType

// This is an alias to an internal value.
const PriceMinimumCompositePriceConversionRateConfigConversionRateTypeUnit = shared.PriceMinimumCompositePriceConversionRateConfigConversionRateTypeUnit

// This is an alias to an internal value.
const PriceMinimumCompositePriceConversionRateConfigConversionRateTypeTiered = shared.PriceMinimumCompositePriceConversionRateConfigConversionRateTypeTiered

// Configuration for minimum_composite pricing
//
// This is an alias to an internal type.
type PriceMinimumCompositePriceMinimumCompositeConfig = shared.PriceMinimumCompositePriceMinimumCompositeConfig

// The pricing model type
//
// This is an alias to an internal type.
type PriceMinimumCompositePriceModelType = shared.PriceMinimumCompositePriceModelType

// This is an alias to an internal value.
const PriceMinimumCompositePriceModelTypeMinimumComposite = shared.PriceMinimumCompositePriceModelTypeMinimumComposite

// This is an alias to an internal type.
type PriceMinimumCompositePricePriceType = shared.PriceMinimumCompositePricePriceType

// This is an alias to an internal value.
const PriceMinimumCompositePricePriceTypeUsagePrice = shared.PriceMinimumCompositePricePriceTypeUsagePrice

// This is an alias to an internal value.
const PriceMinimumCompositePricePriceTypeFixedPrice = shared.PriceMinimumCompositePricePriceTypeFixedPrice

// This is an alias to an internal value.
const PriceMinimumCompositePricePriceTypeCompositePrice = shared.PriceMinimumCompositePricePriceTypeCompositePrice

// The LicenseType resource represents a type of license that can be assigned to
// users. License types are used during billing by grouping metrics on the
// configured grouping key.
//
// This is an alias to an internal type.
type PriceMinimumCompositePriceLicenseType = shared.PriceMinimumCompositePriceLicenseType

// This is an alias to an internal type.
type PricePercentCompositePrice = shared.PricePercentCompositePrice

// This is an alias to an internal type.
type PricePercentCompositePriceBillingMode = shared.PricePercentCompositePriceBillingMode

// This is an alias to an internal value.
const PricePercentCompositePriceBillingModeInAdvance = shared.PricePercentCompositePriceBillingModeInAdvance

// This is an alias to an internal value.
const PricePercentCompositePriceBillingModeInArrear = shared.PricePercentCompositePriceBillingModeInArrear

// This is an alias to an internal type.
type PricePercentCompositePriceCadence = shared.PricePercentCompositePriceCadence

// This is an alias to an internal value.
const PricePercentCompositePriceCadenceOneTime = shared.PricePercentCompositePriceCadenceOneTime

// This is an alias to an internal value.
const PricePercentCompositePriceCadenceMonthly = shared.PricePercentCompositePriceCadenceMonthly

// This is an alias to an internal value.
const PricePercentCompositePriceCadenceQuarterly = shared.PricePercentCompositePriceCadenceQuarterly

// This is an alias to an internal value.
const PricePercentCompositePriceCadenceSemiAnnual = shared.PricePercentCompositePriceCadenceSemiAnnual

// This is an alias to an internal value.
const PricePercentCompositePriceCadenceAnnual = shared.PricePercentCompositePriceCadenceAnnual

// This is an alias to an internal value.
const PricePercentCompositePriceCadenceCustom = shared.PricePercentCompositePriceCadenceCustom

// This is an alias to an internal type.
type PricePercentCompositePriceCompositePriceFilter = shared.PricePercentCompositePriceCompositePriceFilter

// The property of the price to filter on.
//
// This is an alias to an internal type.
type PricePercentCompositePriceCompositePriceFiltersField = shared.PricePercentCompositePriceCompositePriceFiltersField

// This is an alias to an internal value.
const PricePercentCompositePriceCompositePriceFiltersFieldPriceID = shared.PricePercentCompositePriceCompositePriceFiltersFieldPriceID

// This is an alias to an internal value.
const PricePercentCompositePriceCompositePriceFiltersFieldItemID = shared.PricePercentCompositePriceCompositePriceFiltersFieldItemID

// This is an alias to an internal value.
const PricePercentCompositePriceCompositePriceFiltersFieldPriceType = shared.PricePercentCompositePriceCompositePriceFiltersFieldPriceType

// This is an alias to an internal value.
const PricePercentCompositePriceCompositePriceFiltersFieldCurrency = shared.PricePercentCompositePriceCompositePriceFiltersFieldCurrency

// This is an alias to an internal value.
const PricePercentCompositePriceCompositePriceFiltersFieldPricingUnitID = shared.PricePercentCompositePriceCompositePriceFiltersFieldPricingUnitID

// Should prices that match the filter be included or excluded.
//
// This is an alias to an internal type.
type PricePercentCompositePriceCompositePriceFiltersOperator = shared.PricePercentCompositePriceCompositePriceFiltersOperator

// This is an alias to an internal value.
const PricePercentCompositePriceCompositePriceFiltersOperatorIncludes = shared.PricePercentCompositePriceCompositePriceFiltersOperatorIncludes

// This is an alias to an internal value.
const PricePercentCompositePriceCompositePriceFiltersOperatorExcludes = shared.PricePercentCompositePriceCompositePriceFiltersOperatorExcludes

// This is an alias to an internal type.
type PricePercentCompositePriceConversionRateConfig = shared.PricePercentCompositePriceConversionRateConfig

// This is an alias to an internal type.
type PricePercentCompositePriceConversionRateConfigConversionRateType = shared.PricePercentCompositePriceConversionRateConfigConversionRateType

// This is an alias to an internal value.
const PricePercentCompositePriceConversionRateConfigConversionRateTypeUnit = shared.PricePercentCompositePriceConversionRateConfigConversionRateTypeUnit

// This is an alias to an internal value.
const PricePercentCompositePriceConversionRateConfigConversionRateTypeTiered = shared.PricePercentCompositePriceConversionRateConfigConversionRateTypeTiered

// The pricing model type
//
// This is an alias to an internal type.
type PricePercentCompositePriceModelType = shared.PricePercentCompositePriceModelType

// This is an alias to an internal value.
const PricePercentCompositePriceModelTypePercent = shared.PricePercentCompositePriceModelTypePercent

// Configuration for percent pricing
//
// This is an alias to an internal type.
type PricePercentCompositePricePercentConfig = shared.PricePercentCompositePricePercentConfig

// This is an alias to an internal type.
type PricePercentCompositePricePriceType = shared.PricePercentCompositePricePriceType

// This is an alias to an internal value.
const PricePercentCompositePricePriceTypeUsagePrice = shared.PricePercentCompositePricePriceTypeUsagePrice

// This is an alias to an internal value.
const PricePercentCompositePricePriceTypeFixedPrice = shared.PricePercentCompositePricePriceTypeFixedPrice

// This is an alias to an internal value.
const PricePercentCompositePricePriceTypeCompositePrice = shared.PricePercentCompositePricePriceTypeCompositePrice

// The LicenseType resource represents a type of license that can be assigned to
// users. License types are used during billing by grouping metrics on the
// configured grouping key.
//
// This is an alias to an internal type.
type PricePercentCompositePriceLicenseType = shared.PricePercentCompositePriceLicenseType

// This is an alias to an internal type.
type PriceEventOutputPrice = shared.PriceEventOutputPrice

// This is an alias to an internal type.
type PriceEventOutputPriceBillingMode = shared.PriceEventOutputPriceBillingMode

// This is an alias to an internal value.
const PriceEventOutputPriceBillingModeInAdvance = shared.PriceEventOutputPriceBillingModeInAdvance

// This is an alias to an internal value.
const PriceEventOutputPriceBillingModeInArrear = shared.PriceEventOutputPriceBillingModeInArrear

// This is an alias to an internal type.
type PriceEventOutputPriceCadence = shared.PriceEventOutputPriceCadence

// This is an alias to an internal value.
const PriceEventOutputPriceCadenceOneTime = shared.PriceEventOutputPriceCadenceOneTime

// This is an alias to an internal value.
const PriceEventOutputPriceCadenceMonthly = shared.PriceEventOutputPriceCadenceMonthly

// This is an alias to an internal value.
const PriceEventOutputPriceCadenceQuarterly = shared.PriceEventOutputPriceCadenceQuarterly

// This is an alias to an internal value.
const PriceEventOutputPriceCadenceSemiAnnual = shared.PriceEventOutputPriceCadenceSemiAnnual

// This is an alias to an internal value.
const PriceEventOutputPriceCadenceAnnual = shared.PriceEventOutputPriceCadenceAnnual

// This is an alias to an internal value.
const PriceEventOutputPriceCadenceCustom = shared.PriceEventOutputPriceCadenceCustom

// This is an alias to an internal type.
type PriceEventOutputPriceCompositePriceFilter = shared.PriceEventOutputPriceCompositePriceFilter

// The property of the price to filter on.
//
// This is an alias to an internal type.
type PriceEventOutputPriceCompositePriceFiltersField = shared.PriceEventOutputPriceCompositePriceFiltersField

// This is an alias to an internal value.
const PriceEventOutputPriceCompositePriceFiltersFieldPriceID = shared.PriceEventOutputPriceCompositePriceFiltersFieldPriceID

// This is an alias to an internal value.
const PriceEventOutputPriceCompositePriceFiltersFieldItemID = shared.PriceEventOutputPriceCompositePriceFiltersFieldItemID

// This is an alias to an internal value.
const PriceEventOutputPriceCompositePriceFiltersFieldPriceType = shared.PriceEventOutputPriceCompositePriceFiltersFieldPriceType

// This is an alias to an internal value.
const PriceEventOutputPriceCompositePriceFiltersFieldCurrency = shared.PriceEventOutputPriceCompositePriceFiltersFieldCurrency

// This is an alias to an internal value.
const PriceEventOutputPriceCompositePriceFiltersFieldPricingUnitID = shared.PriceEventOutputPriceCompositePriceFiltersFieldPricingUnitID

// Should prices that match the filter be included or excluded.
//
// This is an alias to an internal type.
type PriceEventOutputPriceCompositePriceFiltersOperator = shared.PriceEventOutputPriceCompositePriceFiltersOperator

// This is an alias to an internal value.
const PriceEventOutputPriceCompositePriceFiltersOperatorIncludes = shared.PriceEventOutputPriceCompositePriceFiltersOperatorIncludes

// This is an alias to an internal value.
const PriceEventOutputPriceCompositePriceFiltersOperatorExcludes = shared.PriceEventOutputPriceCompositePriceFiltersOperatorExcludes

// This is an alias to an internal type.
type PriceEventOutputPriceConversionRateConfig = shared.PriceEventOutputPriceConversionRateConfig

// This is an alias to an internal type.
type PriceEventOutputPriceConversionRateConfigConversionRateType = shared.PriceEventOutputPriceConversionRateConfigConversionRateType

// This is an alias to an internal value.
const PriceEventOutputPriceConversionRateConfigConversionRateTypeUnit = shared.PriceEventOutputPriceConversionRateConfigConversionRateTypeUnit

// This is an alias to an internal value.
const PriceEventOutputPriceConversionRateConfigConversionRateTypeTiered = shared.PriceEventOutputPriceConversionRateConfigConversionRateTypeTiered

// Configuration for event_output pricing
//
// This is an alias to an internal type.
type PriceEventOutputPriceEventOutputConfig = shared.PriceEventOutputPriceEventOutputConfig

// The pricing model type
//
// This is an alias to an internal type.
type PriceEventOutputPriceModelType = shared.PriceEventOutputPriceModelType

// This is an alias to an internal value.
const PriceEventOutputPriceModelTypeEventOutput = shared.PriceEventOutputPriceModelTypeEventOutput

// This is an alias to an internal type.
type PriceEventOutputPricePriceType = shared.PriceEventOutputPricePriceType

// This is an alias to an internal value.
const PriceEventOutputPricePriceTypeUsagePrice = shared.PriceEventOutputPricePriceTypeUsagePrice

// This is an alias to an internal value.
const PriceEventOutputPricePriceTypeFixedPrice = shared.PriceEventOutputPricePriceTypeFixedPrice

// This is an alias to an internal value.
const PriceEventOutputPricePriceTypeCompositePrice = shared.PriceEventOutputPricePriceTypeCompositePrice

// The LicenseType resource represents a type of license that can be assigned to
// users. License types are used during billing by grouping metrics on the
// configured grouping key.
//
// This is an alias to an internal type.
type PriceEventOutputPriceLicenseType = shared.PriceEventOutputPriceLicenseType

// This is an alias to an internal type.
type PriceBillingMode = shared.PriceBillingMode

// This is an alias to an internal value.
const PriceBillingModeInAdvance = shared.PriceBillingModeInAdvance

// This is an alias to an internal value.
const PriceBillingModeInArrear = shared.PriceBillingModeInArrear

// This is an alias to an internal type.
type PriceCadence = shared.PriceCadence

// This is an alias to an internal value.
const PriceCadenceOneTime = shared.PriceCadenceOneTime

// This is an alias to an internal value.
const PriceCadenceMonthly = shared.PriceCadenceMonthly

// This is an alias to an internal value.
const PriceCadenceQuarterly = shared.PriceCadenceQuarterly

// This is an alias to an internal value.
const PriceCadenceSemiAnnual = shared.PriceCadenceSemiAnnual

// This is an alias to an internal value.
const PriceCadenceAnnual = shared.PriceCadenceAnnual

// This is an alias to an internal value.
const PriceCadenceCustom = shared.PriceCadenceCustom

// The pricing model type
//
// This is an alias to an internal type.
type PriceModelType = shared.PriceModelType

// This is an alias to an internal value.
const PriceModelTypeUnit = shared.PriceModelTypeUnit

// This is an alias to an internal value.
const PriceModelTypeTiered = shared.PriceModelTypeTiered

// This is an alias to an internal value.
const PriceModelTypeBulk = shared.PriceModelTypeBulk

// This is an alias to an internal value.
const PriceModelTypeBulkWithFilters = shared.PriceModelTypeBulkWithFilters

// This is an alias to an internal value.
const PriceModelTypePackage = shared.PriceModelTypePackage

// This is an alias to an internal value.
const PriceModelTypeMatrix = shared.PriceModelTypeMatrix

// This is an alias to an internal value.
const PriceModelTypeThresholdTotalAmount = shared.PriceModelTypeThresholdTotalAmount

// This is an alias to an internal value.
const PriceModelTypeTieredPackage = shared.PriceModelTypeTieredPackage

// This is an alias to an internal value.
const PriceModelTypeTieredWithMinimum = shared.PriceModelTypeTieredWithMinimum

// This is an alias to an internal value.
const PriceModelTypeGroupedTiered = shared.PriceModelTypeGroupedTiered

// This is an alias to an internal value.
const PriceModelTypeTieredPackageWithMinimum = shared.PriceModelTypeTieredPackageWithMinimum

// This is an alias to an internal value.
const PriceModelTypePackageWithAllocation = shared.PriceModelTypePackageWithAllocation

// This is an alias to an internal value.
const PriceModelTypeUnitWithPercent = shared.PriceModelTypeUnitWithPercent

// This is an alias to an internal value.
const PriceModelTypeMatrixWithAllocation = shared.PriceModelTypeMatrixWithAllocation

// This is an alias to an internal value.
const PriceModelTypeTieredWithProration = shared.PriceModelTypeTieredWithProration

// This is an alias to an internal value.
const PriceModelTypeUnitWithProration = shared.PriceModelTypeUnitWithProration

// This is an alias to an internal value.
const PriceModelTypeGroupedAllocation = shared.PriceModelTypeGroupedAllocation

// This is an alias to an internal value.
const PriceModelTypeBulkWithProration = shared.PriceModelTypeBulkWithProration

// This is an alias to an internal value.
const PriceModelTypeGroupedWithProratedMinimum = shared.PriceModelTypeGroupedWithProratedMinimum

// This is an alias to an internal value.
const PriceModelTypeGroupedWithMeteredMinimum = shared.PriceModelTypeGroupedWithMeteredMinimum

// This is an alias to an internal value.
const PriceModelTypeGroupedWithMinMaxThresholds = shared.PriceModelTypeGroupedWithMinMaxThresholds

// This is an alias to an internal value.
const PriceModelTypeMatrixWithDisplayName = shared.PriceModelTypeMatrixWithDisplayName

// This is an alias to an internal value.
const PriceModelTypeGroupedTieredPackage = shared.PriceModelTypeGroupedTieredPackage

// This is an alias to an internal value.
const PriceModelTypeMaxGroupTieredPackage = shared.PriceModelTypeMaxGroupTieredPackage

// This is an alias to an internal value.
const PriceModelTypeScalableMatrixWithUnitPricing = shared.PriceModelTypeScalableMatrixWithUnitPricing

// This is an alias to an internal value.
const PriceModelTypeScalableMatrixWithTieredPricing = shared.PriceModelTypeScalableMatrixWithTieredPricing

// This is an alias to an internal value.
const PriceModelTypeCumulativeGroupedBulk = shared.PriceModelTypeCumulativeGroupedBulk

// This is an alias to an internal value.
const PriceModelTypeCumulativeGroupedAllocation = shared.PriceModelTypeCumulativeGroupedAllocation

// This is an alias to an internal value.
const PriceModelTypeMinimumComposite = shared.PriceModelTypeMinimumComposite

// This is an alias to an internal value.
const PriceModelTypePercent = shared.PriceModelTypePercent

// This is an alias to an internal value.
const PriceModelTypeEventOutput = shared.PriceModelTypeEventOutput

// This is an alias to an internal type.
type PricePriceType = shared.PricePriceType

// This is an alias to an internal value.
const PricePriceTypeUsagePrice = shared.PricePriceTypeUsagePrice

// This is an alias to an internal value.
const PricePriceTypeFixedPrice = shared.PricePriceTypeFixedPrice

// This is an alias to an internal value.
const PricePriceTypeCompositePrice = shared.PricePriceTypeCompositePrice

// The Price Interval resource represents a period of time for which a price will
// bill on a subscription. A subscription’s price intervals define its billing
// behavior.
//
// This is an alias to an internal type.
type PriceInterval = shared.PriceInterval

// This is an alias to an internal type.
type SubLineItemGrouping = shared.SubLineItemGrouping

// This is an alias to an internal type.
type SubLineItemMatrixConfig = shared.SubLineItemMatrixConfig

// This is an alias to an internal type.
type SubscriptionChangeMinified = shared.SubscriptionChangeMinified

// This is an alias to an internal type.
type SubscriptionMinified = shared.SubscriptionMinified

// This is an alias to an internal type.
type SubscriptionTrialInfo = shared.SubscriptionTrialInfo

// This is an alias to an internal type.
type TaxAmount = shared.TaxAmount

// Configuration for a single tier
//
// This is an alias to an internal type.
type Tier = shared.Tier

// Configuration for a single tier
//
// This is an alias to an internal type.
type TierParam = shared.TierParam

// This is an alias to an internal type.
type TierSubLineItem = shared.TierSubLineItem

// This is an alias to an internal type.
type TierSubLineItemTierConfig = shared.TierSubLineItemTierConfig

// This is an alias to an internal type.
type TierSubLineItemType = shared.TierSubLineItemType

// This is an alias to an internal value.
const TierSubLineItemTypeTier = shared.TierSubLineItemTypeTier

// Configuration for tiered pricing
//
// This is an alias to an internal type.
type TieredConfig = shared.TieredConfig

// Configuration for tiered pricing
//
// This is an alias to an internal type.
type TieredConfigParam = shared.TieredConfigParam

// This is an alias to an internal type.
type TieredConversionRateConfig = shared.TieredConversionRateConfig

// This is an alias to an internal type.
type TieredConversionRateConfigConversionRateType = shared.TieredConversionRateConfigConversionRateType

// This is an alias to an internal value.
const TieredConversionRateConfigConversionRateTypeTiered = shared.TieredConversionRateConfigConversionRateTypeTiered

// This is an alias to an internal type.
type TieredConversionRateConfigParam = shared.TieredConversionRateConfigParam

// This is an alias to an internal type.
type TrialDiscount = shared.TrialDiscount

// This is an alias to an internal type.
type TrialDiscountDiscountType = shared.TrialDiscountDiscountType

// This is an alias to an internal value.
const TrialDiscountDiscountTypeTrial = shared.TrialDiscountDiscountTypeTrial

// This is an alias to an internal type.
type TrialDiscountFilter = shared.TrialDiscountFilter

// The property of the price to filter on.
//
// This is an alias to an internal type.
type TrialDiscountFiltersField = shared.TrialDiscountFiltersField

// This is an alias to an internal value.
const TrialDiscountFiltersFieldPriceID = shared.TrialDiscountFiltersFieldPriceID

// This is an alias to an internal value.
const TrialDiscountFiltersFieldItemID = shared.TrialDiscountFiltersFieldItemID

// This is an alias to an internal value.
const TrialDiscountFiltersFieldPriceType = shared.TrialDiscountFiltersFieldPriceType

// This is an alias to an internal value.
const TrialDiscountFiltersFieldCurrency = shared.TrialDiscountFiltersFieldCurrency

// This is an alias to an internal value.
const TrialDiscountFiltersFieldPricingUnitID = shared.TrialDiscountFiltersFieldPricingUnitID

// Should prices that match the filter be included or excluded.
//
// This is an alias to an internal type.
type TrialDiscountFiltersOperator = shared.TrialDiscountFiltersOperator

// This is an alias to an internal value.
const TrialDiscountFiltersOperatorIncludes = shared.TrialDiscountFiltersOperatorIncludes

// This is an alias to an internal value.
const TrialDiscountFiltersOperatorExcludes = shared.TrialDiscountFiltersOperatorExcludes

// This is an alias to an internal type.
type TrialDiscountParam = shared.TrialDiscountParam

// This is an alias to an internal type.
type TrialDiscountFilterParam = shared.TrialDiscountFilterParam

// Configuration for unit pricing
//
// This is an alias to an internal type.
type UnitConfig = shared.UnitConfig

// Configuration for unit pricing
//
// This is an alias to an internal type.
type UnitConfigParam = shared.UnitConfigParam

// This is an alias to an internal type.
type UnitConversionRateConfig = shared.UnitConversionRateConfig

// This is an alias to an internal type.
type UnitConversionRateConfigConversionRateType = shared.UnitConversionRateConfigConversionRateType

// This is an alias to an internal value.
const UnitConversionRateConfigConversionRateTypeUnit = shared.UnitConversionRateConfigConversionRateTypeUnit

// This is an alias to an internal type.
type UnitConversionRateConfigParam = shared.UnitConversionRateConfigParam

// This is an alias to an internal type.
type UsageDiscount = shared.UsageDiscount

// This is an alias to an internal type.
type UsageDiscountDiscountType = shared.UsageDiscountDiscountType

// This is an alias to an internal value.
const UsageDiscountDiscountTypeUsage = shared.UsageDiscountDiscountTypeUsage

// This is an alias to an internal type.
type UsageDiscountFilter = shared.UsageDiscountFilter

// The property of the price to filter on.
//
// This is an alias to an internal type.
type UsageDiscountFiltersField = shared.UsageDiscountFiltersField

// This is an alias to an internal value.
const UsageDiscountFiltersFieldPriceID = shared.UsageDiscountFiltersFieldPriceID

// This is an alias to an internal value.
const UsageDiscountFiltersFieldItemID = shared.UsageDiscountFiltersFieldItemID

// This is an alias to an internal value.
const UsageDiscountFiltersFieldPriceType = shared.UsageDiscountFiltersFieldPriceType

// This is an alias to an internal value.
const UsageDiscountFiltersFieldCurrency = shared.UsageDiscountFiltersFieldCurrency

// This is an alias to an internal value.
const UsageDiscountFiltersFieldPricingUnitID = shared.UsageDiscountFiltersFieldPricingUnitID

// Should prices that match the filter be included or excluded.
//
// This is an alias to an internal type.
type UsageDiscountFiltersOperator = shared.UsageDiscountFiltersOperator

// This is an alias to an internal value.
const UsageDiscountFiltersOperatorIncludes = shared.UsageDiscountFiltersOperatorIncludes

// This is an alias to an internal value.
const UsageDiscountFiltersOperatorExcludes = shared.UsageDiscountFiltersOperatorExcludes

// This is an alias to an internal type.
type UsageDiscountParam = shared.UsageDiscountParam

// This is an alias to an internal type.
type UsageDiscountFilterParam = shared.UsageDiscountFilterParam

// This is an alias to an internal type.
type UsageDiscountInterval = shared.UsageDiscountInterval

// This is an alias to an internal type.
type UsageDiscountIntervalDiscountType = shared.UsageDiscountIntervalDiscountType

// This is an alias to an internal value.
const UsageDiscountIntervalDiscountTypeUsage = shared.UsageDiscountIntervalDiscountTypeUsage

// This is an alias to an internal type.
type UsageDiscountIntervalFilter = shared.UsageDiscountIntervalFilter

// The property of the price to filter on.
//
// This is an alias to an internal type.
type UsageDiscountIntervalFiltersField = shared.UsageDiscountIntervalFiltersField

// This is an alias to an internal value.
const UsageDiscountIntervalFiltersFieldPriceID = shared.UsageDiscountIntervalFiltersFieldPriceID

// This is an alias to an internal value.
const UsageDiscountIntervalFiltersFieldItemID = shared.UsageDiscountIntervalFiltersFieldItemID

// This is an alias to an internal value.
const UsageDiscountIntervalFiltersFieldPriceType = shared.UsageDiscountIntervalFiltersFieldPriceType

// This is an alias to an internal value.
const UsageDiscountIntervalFiltersFieldCurrency = shared.UsageDiscountIntervalFiltersFieldCurrency

// This is an alias to an internal value.
const UsageDiscountIntervalFiltersFieldPricingUnitID = shared.UsageDiscountIntervalFiltersFieldPricingUnitID

// Should prices that match the filter be included or excluded.
//
// This is an alias to an internal type.
type UsageDiscountIntervalFiltersOperator = shared.UsageDiscountIntervalFiltersOperator

// This is an alias to an internal value.
const UsageDiscountIntervalFiltersOperatorIncludes = shared.UsageDiscountIntervalFiltersOperatorIncludes

// This is an alias to an internal value.
const UsageDiscountIntervalFiltersOperatorExcludes = shared.UsageDiscountIntervalFiltersOperatorExcludes
