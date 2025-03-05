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
type AddCreditLedgerEntryRequestUnionParam = shared.AddCreditLedgerEntryRequestUnionParam

// This is an alias to an internal type.
type AddCreditLedgerEntryRequestAddIncrementCreditLedgerEntryRequestParamsParam = shared.AddCreditLedgerEntryRequestAddIncrementCreditLedgerEntryRequestParamsParam

// This is an alias to an internal type.
type AddCreditLedgerEntryRequestAddIncrementCreditLedgerEntryRequestParamsEntryType = shared.AddCreditLedgerEntryRequestAddIncrementCreditLedgerEntryRequestParamsEntryType

// This is an alias to an internal value.
const AddCreditLedgerEntryRequestAddIncrementCreditLedgerEntryRequestParamsEntryTypeIncrement = shared.AddCreditLedgerEntryRequestAddIncrementCreditLedgerEntryRequestParamsEntryTypeIncrement

// Passing `invoice_settings` automatically generates an invoice for the newly
// added credits. If `invoice_settings` is passed, you must specify
// per_unit_cost_basis, as the calculation of the invoice total is done on that
// basis.
//
// This is an alias to an internal type.
type AddCreditLedgerEntryRequestAddIncrementCreditLedgerEntryRequestParamsInvoiceSettingsParam = shared.AddCreditLedgerEntryRequestAddIncrementCreditLedgerEntryRequestParamsInvoiceSettingsParam

// This is an alias to an internal type.
type AddCreditLedgerEntryRequestAddDecrementCreditLedgerEntryRequestParamsParam = shared.AddCreditLedgerEntryRequestAddDecrementCreditLedgerEntryRequestParamsParam

// This is an alias to an internal type.
type AddCreditLedgerEntryRequestAddDecrementCreditLedgerEntryRequestParamsEntryType = shared.AddCreditLedgerEntryRequestAddDecrementCreditLedgerEntryRequestParamsEntryType

// This is an alias to an internal value.
const AddCreditLedgerEntryRequestAddDecrementCreditLedgerEntryRequestParamsEntryTypeDecrement = shared.AddCreditLedgerEntryRequestAddDecrementCreditLedgerEntryRequestParamsEntryTypeDecrement

// This is an alias to an internal type.
type AddCreditLedgerEntryRequestAddExpirationChangeCreditLedgerEntryRequestParamsParam = shared.AddCreditLedgerEntryRequestAddExpirationChangeCreditLedgerEntryRequestParamsParam

// This is an alias to an internal type.
type AddCreditLedgerEntryRequestAddExpirationChangeCreditLedgerEntryRequestParamsEntryType = shared.AddCreditLedgerEntryRequestAddExpirationChangeCreditLedgerEntryRequestParamsEntryType

// This is an alias to an internal value.
const AddCreditLedgerEntryRequestAddExpirationChangeCreditLedgerEntryRequestParamsEntryTypeExpirationChange = shared.AddCreditLedgerEntryRequestAddExpirationChangeCreditLedgerEntryRequestParamsEntryTypeExpirationChange

// This is an alias to an internal type.
type AddCreditLedgerEntryRequestAddVoidCreditLedgerEntryRequestParamsParam = shared.AddCreditLedgerEntryRequestAddVoidCreditLedgerEntryRequestParamsParam

// This is an alias to an internal type.
type AddCreditLedgerEntryRequestAddVoidCreditLedgerEntryRequestParamsEntryType = shared.AddCreditLedgerEntryRequestAddVoidCreditLedgerEntryRequestParamsEntryType

// This is an alias to an internal value.
const AddCreditLedgerEntryRequestAddVoidCreditLedgerEntryRequestParamsEntryTypeVoid = shared.AddCreditLedgerEntryRequestAddVoidCreditLedgerEntryRequestParamsEntryTypeVoid

// Can only be specified when `entry_type=void`. The reason for the void.
//
// This is an alias to an internal type.
type AddCreditLedgerEntryRequestAddVoidCreditLedgerEntryRequestParamsVoidReason = shared.AddCreditLedgerEntryRequestAddVoidCreditLedgerEntryRequestParamsVoidReason

// This is an alias to an internal value.
const AddCreditLedgerEntryRequestAddVoidCreditLedgerEntryRequestParamsVoidReasonRefund = shared.AddCreditLedgerEntryRequestAddVoidCreditLedgerEntryRequestParamsVoidReasonRefund

// This is an alias to an internal type.
type AddCreditLedgerEntryRequestAddAmendmentCreditLedgerEntryRequestParamsParam = shared.AddCreditLedgerEntryRequestAddAmendmentCreditLedgerEntryRequestParamsParam

// This is an alias to an internal type.
type AddCreditLedgerEntryRequestAddAmendmentCreditLedgerEntryRequestParamsEntryType = shared.AddCreditLedgerEntryRequestAddAmendmentCreditLedgerEntryRequestParamsEntryType

// This is an alias to an internal value.
const AddCreditLedgerEntryRequestAddAmendmentCreditLedgerEntryRequestParamsEntryTypeAmendment = shared.AddCreditLedgerEntryRequestAddAmendmentCreditLedgerEntryRequestParamsEntryTypeAmendment

// This is an alias to an internal type.
type AddCreditLedgerEntryRequestEntryType = shared.AddCreditLedgerEntryRequestEntryType

// This is an alias to an internal value.
const AddCreditLedgerEntryRequestEntryTypeIncrement = shared.AddCreditLedgerEntryRequestEntryTypeIncrement

// This is an alias to an internal value.
const AddCreditLedgerEntryRequestEntryTypeDecrement = shared.AddCreditLedgerEntryRequestEntryTypeDecrement

// This is an alias to an internal value.
const AddCreditLedgerEntryRequestEntryTypeExpirationChange = shared.AddCreditLedgerEntryRequestEntryTypeExpirationChange

// This is an alias to an internal value.
const AddCreditLedgerEntryRequestEntryTypeVoid = shared.AddCreditLedgerEntryRequestEntryTypeVoid

// This is an alias to an internal value.
const AddCreditLedgerEntryRequestEntryTypeAmendment = shared.AddCreditLedgerEntryRequestEntryTypeAmendment

// Can only be specified when `entry_type=void`. The reason for the void.
//
// This is an alias to an internal type.
type AddCreditLedgerEntryRequestVoidReason = shared.AddCreditLedgerEntryRequestVoidReason

// This is an alias to an internal value.
const AddCreditLedgerEntryRequestVoidReasonRefund = shared.AddCreditLedgerEntryRequestVoidReasonRefund

// This is an alias to an internal type.
type AddSubscriptionAdjustmentParams = shared.AddSubscriptionAdjustmentParams

// The definition of a new adjustment to create and add to the subscription.
//
// This is an alias to an internal type.
type AddSubscriptionAdjustmentParamsAdjustmentUnion = shared.AddSubscriptionAdjustmentParamsAdjustmentUnion

// This is an alias to an internal type.
type AddSubscriptionAdjustmentParamsAdjustmentNewPercentageDiscount = shared.AddSubscriptionAdjustmentParamsAdjustmentNewPercentageDiscount

// This is an alias to an internal type.
type AddSubscriptionAdjustmentParamsAdjustmentNewPercentageDiscountAdjustmentType = shared.AddSubscriptionAdjustmentParamsAdjustmentNewPercentageDiscountAdjustmentType

// This is an alias to an internal value.
const AddSubscriptionAdjustmentParamsAdjustmentNewPercentageDiscountAdjustmentTypePercentageDiscount = shared.AddSubscriptionAdjustmentParamsAdjustmentNewPercentageDiscountAdjustmentTypePercentageDiscount

// This is an alias to an internal type.
type AddSubscriptionAdjustmentParamsAdjustmentNewUsageDiscount = shared.AddSubscriptionAdjustmentParamsAdjustmentNewUsageDiscount

// This is an alias to an internal type.
type AddSubscriptionAdjustmentParamsAdjustmentNewUsageDiscountAdjustmentType = shared.AddSubscriptionAdjustmentParamsAdjustmentNewUsageDiscountAdjustmentType

// This is an alias to an internal value.
const AddSubscriptionAdjustmentParamsAdjustmentNewUsageDiscountAdjustmentTypeUsageDiscount = shared.AddSubscriptionAdjustmentParamsAdjustmentNewUsageDiscountAdjustmentTypeUsageDiscount

// This is an alias to an internal type.
type AddSubscriptionAdjustmentParamsAdjustmentNewAmountDiscount = shared.AddSubscriptionAdjustmentParamsAdjustmentNewAmountDiscount

// This is an alias to an internal type.
type AddSubscriptionAdjustmentParamsAdjustmentNewAmountDiscountAdjustmentType = shared.AddSubscriptionAdjustmentParamsAdjustmentNewAmountDiscountAdjustmentType

// This is an alias to an internal value.
const AddSubscriptionAdjustmentParamsAdjustmentNewAmountDiscountAdjustmentTypeAmountDiscount = shared.AddSubscriptionAdjustmentParamsAdjustmentNewAmountDiscountAdjustmentTypeAmountDiscount

// This is an alias to an internal type.
type AddSubscriptionAdjustmentParamsAdjustmentNewMinimum = shared.AddSubscriptionAdjustmentParamsAdjustmentNewMinimum

// This is an alias to an internal type.
type AddSubscriptionAdjustmentParamsAdjustmentNewMinimumAdjustmentType = shared.AddSubscriptionAdjustmentParamsAdjustmentNewMinimumAdjustmentType

// This is an alias to an internal value.
const AddSubscriptionAdjustmentParamsAdjustmentNewMinimumAdjustmentTypeMinimum = shared.AddSubscriptionAdjustmentParamsAdjustmentNewMinimumAdjustmentTypeMinimum

// This is an alias to an internal type.
type AddSubscriptionAdjustmentParamsAdjustmentNewMaximum = shared.AddSubscriptionAdjustmentParamsAdjustmentNewMaximum

// This is an alias to an internal type.
type AddSubscriptionAdjustmentParamsAdjustmentNewMaximumAdjustmentType = shared.AddSubscriptionAdjustmentParamsAdjustmentNewMaximumAdjustmentType

// This is an alias to an internal value.
const AddSubscriptionAdjustmentParamsAdjustmentNewMaximumAdjustmentTypeMaximum = shared.AddSubscriptionAdjustmentParamsAdjustmentNewMaximumAdjustmentTypeMaximum

// This is an alias to an internal type.
type AddSubscriptionAdjustmentParamsAdjustmentAdjustmentType = shared.AddSubscriptionAdjustmentParamsAdjustmentAdjustmentType

// This is an alias to an internal value.
const AddSubscriptionAdjustmentParamsAdjustmentAdjustmentTypePercentageDiscount = shared.AddSubscriptionAdjustmentParamsAdjustmentAdjustmentTypePercentageDiscount

// This is an alias to an internal value.
const AddSubscriptionAdjustmentParamsAdjustmentAdjustmentTypeUsageDiscount = shared.AddSubscriptionAdjustmentParamsAdjustmentAdjustmentTypeUsageDiscount

// This is an alias to an internal value.
const AddSubscriptionAdjustmentParamsAdjustmentAdjustmentTypeAmountDiscount = shared.AddSubscriptionAdjustmentParamsAdjustmentAdjustmentTypeAmountDiscount

// This is an alias to an internal value.
const AddSubscriptionAdjustmentParamsAdjustmentAdjustmentTypeMinimum = shared.AddSubscriptionAdjustmentParamsAdjustmentAdjustmentTypeMinimum

// This is an alias to an internal value.
const AddSubscriptionAdjustmentParamsAdjustmentAdjustmentTypeMaximum = shared.AddSubscriptionAdjustmentParamsAdjustmentAdjustmentTypeMaximum

// This is an alias to an internal type.
type AddSubscriptionPriceParams = shared.AddSubscriptionPriceParams

// The definition of a new allocation price to create and add to the subscription.
//
// This is an alias to an internal type.
type AddSubscriptionPriceParamsAllocationPrice = shared.AddSubscriptionPriceParamsAllocationPrice

// The cadence at which to allocate the amount to the customer.
//
// This is an alias to an internal type.
type AddSubscriptionPriceParamsAllocationPriceCadence = shared.AddSubscriptionPriceParamsAllocationPriceCadence

// This is an alias to an internal value.
const AddSubscriptionPriceParamsAllocationPriceCadenceOneTime = shared.AddSubscriptionPriceParamsAllocationPriceCadenceOneTime

// This is an alias to an internal value.
const AddSubscriptionPriceParamsAllocationPriceCadenceMonthly = shared.AddSubscriptionPriceParamsAllocationPriceCadenceMonthly

// This is an alias to an internal value.
const AddSubscriptionPriceParamsAllocationPriceCadenceQuarterly = shared.AddSubscriptionPriceParamsAllocationPriceCadenceQuarterly

// This is an alias to an internal value.
const AddSubscriptionPriceParamsAllocationPriceCadenceSemiAnnual = shared.AddSubscriptionPriceParamsAllocationPriceCadenceSemiAnnual

// This is an alias to an internal value.
const AddSubscriptionPriceParamsAllocationPriceCadenceAnnual = shared.AddSubscriptionPriceParamsAllocationPriceCadenceAnnual

// This is an alias to an internal value.
const AddSubscriptionPriceParamsAllocationPriceCadenceCustom = shared.AddSubscriptionPriceParamsAllocationPriceCadenceCustom

// This is an alias to an internal type.
type AddSubscriptionPriceParamsDiscount = shared.AddSubscriptionPriceParamsDiscount

// This is an alias to an internal type.
type AddSubscriptionPriceParamsDiscountsDiscountType = shared.AddSubscriptionPriceParamsDiscountsDiscountType

// This is an alias to an internal value.
const AddSubscriptionPriceParamsDiscountsDiscountTypePercentage = shared.AddSubscriptionPriceParamsDiscountsDiscountTypePercentage

// This is an alias to an internal value.
const AddSubscriptionPriceParamsDiscountsDiscountTypeUsage = shared.AddSubscriptionPriceParamsDiscountsDiscountTypeUsage

// This is an alias to an internal value.
const AddSubscriptionPriceParamsDiscountsDiscountTypeAmount = shared.AddSubscriptionPriceParamsDiscountsDiscountTypeAmount

// The definition of a new price to create and add to the subscription.
//
// This is an alias to an internal type.
type AddSubscriptionPriceParamsPriceUnion = shared.AddSubscriptionPriceParamsPriceUnion

// This is an alias to an internal type.
type AddSubscriptionPriceParamsPriceNewSubscriptionUnitPrice = shared.AddSubscriptionPriceParamsPriceNewSubscriptionUnitPrice

// The cadence to bill for this price on.
//
// This is an alias to an internal type.
type AddSubscriptionPriceParamsPriceNewSubscriptionUnitPriceCadence = shared.AddSubscriptionPriceParamsPriceNewSubscriptionUnitPriceCadence

// This is an alias to an internal value.
const AddSubscriptionPriceParamsPriceNewSubscriptionUnitPriceCadenceAnnual = shared.AddSubscriptionPriceParamsPriceNewSubscriptionUnitPriceCadenceAnnual

// This is an alias to an internal value.
const AddSubscriptionPriceParamsPriceNewSubscriptionUnitPriceCadenceSemiAnnual = shared.AddSubscriptionPriceParamsPriceNewSubscriptionUnitPriceCadenceSemiAnnual

// This is an alias to an internal value.
const AddSubscriptionPriceParamsPriceNewSubscriptionUnitPriceCadenceMonthly = shared.AddSubscriptionPriceParamsPriceNewSubscriptionUnitPriceCadenceMonthly

// This is an alias to an internal value.
const AddSubscriptionPriceParamsPriceNewSubscriptionUnitPriceCadenceQuarterly = shared.AddSubscriptionPriceParamsPriceNewSubscriptionUnitPriceCadenceQuarterly

// This is an alias to an internal value.
const AddSubscriptionPriceParamsPriceNewSubscriptionUnitPriceCadenceOneTime = shared.AddSubscriptionPriceParamsPriceNewSubscriptionUnitPriceCadenceOneTime

// This is an alias to an internal value.
const AddSubscriptionPriceParamsPriceNewSubscriptionUnitPriceCadenceCustom = shared.AddSubscriptionPriceParamsPriceNewSubscriptionUnitPriceCadenceCustom

// This is an alias to an internal type.
type AddSubscriptionPriceParamsPriceNewSubscriptionUnitPriceModelType = shared.AddSubscriptionPriceParamsPriceNewSubscriptionUnitPriceModelType

// This is an alias to an internal value.
const AddSubscriptionPriceParamsPriceNewSubscriptionUnitPriceModelTypeUnit = shared.AddSubscriptionPriceParamsPriceNewSubscriptionUnitPriceModelTypeUnit

// This is an alias to an internal type.
type AddSubscriptionPriceParamsPriceNewSubscriptionUnitPriceUnitConfig = shared.AddSubscriptionPriceParamsPriceNewSubscriptionUnitPriceUnitConfig

// For custom cadence: specifies the duration of the billing period in days or
// months.
//
// This is an alias to an internal type.
type AddSubscriptionPriceParamsPriceNewSubscriptionUnitPriceBillingCycleConfiguration = shared.AddSubscriptionPriceParamsPriceNewSubscriptionUnitPriceBillingCycleConfiguration

// The unit of billing period duration.
//
// This is an alias to an internal type.
type AddSubscriptionPriceParamsPriceNewSubscriptionUnitPriceBillingCycleConfigurationDurationUnit = shared.AddSubscriptionPriceParamsPriceNewSubscriptionUnitPriceBillingCycleConfigurationDurationUnit

// This is an alias to an internal value.
const AddSubscriptionPriceParamsPriceNewSubscriptionUnitPriceBillingCycleConfigurationDurationUnitDay = shared.AddSubscriptionPriceParamsPriceNewSubscriptionUnitPriceBillingCycleConfigurationDurationUnitDay

// This is an alias to an internal value.
const AddSubscriptionPriceParamsPriceNewSubscriptionUnitPriceBillingCycleConfigurationDurationUnitMonth = shared.AddSubscriptionPriceParamsPriceNewSubscriptionUnitPriceBillingCycleConfigurationDurationUnitMonth

// Within each billing cycle, specifies the cadence at which invoices are produced.
// If unspecified, a single invoice is produced per billing cycle.
//
// This is an alias to an internal type.
type AddSubscriptionPriceParamsPriceNewSubscriptionUnitPriceInvoicingCycleConfiguration = shared.AddSubscriptionPriceParamsPriceNewSubscriptionUnitPriceInvoicingCycleConfiguration

// The unit of billing period duration.
//
// This is an alias to an internal type.
type AddSubscriptionPriceParamsPriceNewSubscriptionUnitPriceInvoicingCycleConfigurationDurationUnit = shared.AddSubscriptionPriceParamsPriceNewSubscriptionUnitPriceInvoicingCycleConfigurationDurationUnit

// This is an alias to an internal value.
const AddSubscriptionPriceParamsPriceNewSubscriptionUnitPriceInvoicingCycleConfigurationDurationUnitDay = shared.AddSubscriptionPriceParamsPriceNewSubscriptionUnitPriceInvoicingCycleConfigurationDurationUnitDay

// This is an alias to an internal value.
const AddSubscriptionPriceParamsPriceNewSubscriptionUnitPriceInvoicingCycleConfigurationDurationUnitMonth = shared.AddSubscriptionPriceParamsPriceNewSubscriptionUnitPriceInvoicingCycleConfigurationDurationUnitMonth

// This is an alias to an internal type.
type AddSubscriptionPriceParamsPriceNewSubscriptionPackagePrice = shared.AddSubscriptionPriceParamsPriceNewSubscriptionPackagePrice

// The cadence to bill for this price on.
//
// This is an alias to an internal type.
type AddSubscriptionPriceParamsPriceNewSubscriptionPackagePriceCadence = shared.AddSubscriptionPriceParamsPriceNewSubscriptionPackagePriceCadence

// This is an alias to an internal value.
const AddSubscriptionPriceParamsPriceNewSubscriptionPackagePriceCadenceAnnual = shared.AddSubscriptionPriceParamsPriceNewSubscriptionPackagePriceCadenceAnnual

// This is an alias to an internal value.
const AddSubscriptionPriceParamsPriceNewSubscriptionPackagePriceCadenceSemiAnnual = shared.AddSubscriptionPriceParamsPriceNewSubscriptionPackagePriceCadenceSemiAnnual

// This is an alias to an internal value.
const AddSubscriptionPriceParamsPriceNewSubscriptionPackagePriceCadenceMonthly = shared.AddSubscriptionPriceParamsPriceNewSubscriptionPackagePriceCadenceMonthly

// This is an alias to an internal value.
const AddSubscriptionPriceParamsPriceNewSubscriptionPackagePriceCadenceQuarterly = shared.AddSubscriptionPriceParamsPriceNewSubscriptionPackagePriceCadenceQuarterly

// This is an alias to an internal value.
const AddSubscriptionPriceParamsPriceNewSubscriptionPackagePriceCadenceOneTime = shared.AddSubscriptionPriceParamsPriceNewSubscriptionPackagePriceCadenceOneTime

// This is an alias to an internal value.
const AddSubscriptionPriceParamsPriceNewSubscriptionPackagePriceCadenceCustom = shared.AddSubscriptionPriceParamsPriceNewSubscriptionPackagePriceCadenceCustom

// This is an alias to an internal type.
type AddSubscriptionPriceParamsPriceNewSubscriptionPackagePriceModelType = shared.AddSubscriptionPriceParamsPriceNewSubscriptionPackagePriceModelType

// This is an alias to an internal value.
const AddSubscriptionPriceParamsPriceNewSubscriptionPackagePriceModelTypePackage = shared.AddSubscriptionPriceParamsPriceNewSubscriptionPackagePriceModelTypePackage

// This is an alias to an internal type.
type AddSubscriptionPriceParamsPriceNewSubscriptionPackagePricePackageConfig = shared.AddSubscriptionPriceParamsPriceNewSubscriptionPackagePricePackageConfig

// For custom cadence: specifies the duration of the billing period in days or
// months.
//
// This is an alias to an internal type.
type AddSubscriptionPriceParamsPriceNewSubscriptionPackagePriceBillingCycleConfiguration = shared.AddSubscriptionPriceParamsPriceNewSubscriptionPackagePriceBillingCycleConfiguration

// The unit of billing period duration.
//
// This is an alias to an internal type.
type AddSubscriptionPriceParamsPriceNewSubscriptionPackagePriceBillingCycleConfigurationDurationUnit = shared.AddSubscriptionPriceParamsPriceNewSubscriptionPackagePriceBillingCycleConfigurationDurationUnit

// This is an alias to an internal value.
const AddSubscriptionPriceParamsPriceNewSubscriptionPackagePriceBillingCycleConfigurationDurationUnitDay = shared.AddSubscriptionPriceParamsPriceNewSubscriptionPackagePriceBillingCycleConfigurationDurationUnitDay

// This is an alias to an internal value.
const AddSubscriptionPriceParamsPriceNewSubscriptionPackagePriceBillingCycleConfigurationDurationUnitMonth = shared.AddSubscriptionPriceParamsPriceNewSubscriptionPackagePriceBillingCycleConfigurationDurationUnitMonth

// Within each billing cycle, specifies the cadence at which invoices are produced.
// If unspecified, a single invoice is produced per billing cycle.
//
// This is an alias to an internal type.
type AddSubscriptionPriceParamsPriceNewSubscriptionPackagePriceInvoicingCycleConfiguration = shared.AddSubscriptionPriceParamsPriceNewSubscriptionPackagePriceInvoicingCycleConfiguration

// The unit of billing period duration.
//
// This is an alias to an internal type.
type AddSubscriptionPriceParamsPriceNewSubscriptionPackagePriceInvoicingCycleConfigurationDurationUnit = shared.AddSubscriptionPriceParamsPriceNewSubscriptionPackagePriceInvoicingCycleConfigurationDurationUnit

// This is an alias to an internal value.
const AddSubscriptionPriceParamsPriceNewSubscriptionPackagePriceInvoicingCycleConfigurationDurationUnitDay = shared.AddSubscriptionPriceParamsPriceNewSubscriptionPackagePriceInvoicingCycleConfigurationDurationUnitDay

// This is an alias to an internal value.
const AddSubscriptionPriceParamsPriceNewSubscriptionPackagePriceInvoicingCycleConfigurationDurationUnitMonth = shared.AddSubscriptionPriceParamsPriceNewSubscriptionPackagePriceInvoicingCycleConfigurationDurationUnitMonth

// This is an alias to an internal type.
type AddSubscriptionPriceParamsPriceNewSubscriptionMatrixPrice = shared.AddSubscriptionPriceParamsPriceNewSubscriptionMatrixPrice

// The cadence to bill for this price on.
//
// This is an alias to an internal type.
type AddSubscriptionPriceParamsPriceNewSubscriptionMatrixPriceCadence = shared.AddSubscriptionPriceParamsPriceNewSubscriptionMatrixPriceCadence

// This is an alias to an internal value.
const AddSubscriptionPriceParamsPriceNewSubscriptionMatrixPriceCadenceAnnual = shared.AddSubscriptionPriceParamsPriceNewSubscriptionMatrixPriceCadenceAnnual

// This is an alias to an internal value.
const AddSubscriptionPriceParamsPriceNewSubscriptionMatrixPriceCadenceSemiAnnual = shared.AddSubscriptionPriceParamsPriceNewSubscriptionMatrixPriceCadenceSemiAnnual

// This is an alias to an internal value.
const AddSubscriptionPriceParamsPriceNewSubscriptionMatrixPriceCadenceMonthly = shared.AddSubscriptionPriceParamsPriceNewSubscriptionMatrixPriceCadenceMonthly

// This is an alias to an internal value.
const AddSubscriptionPriceParamsPriceNewSubscriptionMatrixPriceCadenceQuarterly = shared.AddSubscriptionPriceParamsPriceNewSubscriptionMatrixPriceCadenceQuarterly

// This is an alias to an internal value.
const AddSubscriptionPriceParamsPriceNewSubscriptionMatrixPriceCadenceOneTime = shared.AddSubscriptionPriceParamsPriceNewSubscriptionMatrixPriceCadenceOneTime

// This is an alias to an internal value.
const AddSubscriptionPriceParamsPriceNewSubscriptionMatrixPriceCadenceCustom = shared.AddSubscriptionPriceParamsPriceNewSubscriptionMatrixPriceCadenceCustom

// This is an alias to an internal type.
type AddSubscriptionPriceParamsPriceNewSubscriptionMatrixPriceMatrixConfig = shared.AddSubscriptionPriceParamsPriceNewSubscriptionMatrixPriceMatrixConfig

// This is an alias to an internal type.
type AddSubscriptionPriceParamsPriceNewSubscriptionMatrixPriceMatrixConfigMatrixValue = shared.AddSubscriptionPriceParamsPriceNewSubscriptionMatrixPriceMatrixConfigMatrixValue

// This is an alias to an internal type.
type AddSubscriptionPriceParamsPriceNewSubscriptionMatrixPriceModelType = shared.AddSubscriptionPriceParamsPriceNewSubscriptionMatrixPriceModelType

// This is an alias to an internal value.
const AddSubscriptionPriceParamsPriceNewSubscriptionMatrixPriceModelTypeMatrix = shared.AddSubscriptionPriceParamsPriceNewSubscriptionMatrixPriceModelTypeMatrix

// For custom cadence: specifies the duration of the billing period in days or
// months.
//
// This is an alias to an internal type.
type AddSubscriptionPriceParamsPriceNewSubscriptionMatrixPriceBillingCycleConfiguration = shared.AddSubscriptionPriceParamsPriceNewSubscriptionMatrixPriceBillingCycleConfiguration

// The unit of billing period duration.
//
// This is an alias to an internal type.
type AddSubscriptionPriceParamsPriceNewSubscriptionMatrixPriceBillingCycleConfigurationDurationUnit = shared.AddSubscriptionPriceParamsPriceNewSubscriptionMatrixPriceBillingCycleConfigurationDurationUnit

// This is an alias to an internal value.
const AddSubscriptionPriceParamsPriceNewSubscriptionMatrixPriceBillingCycleConfigurationDurationUnitDay = shared.AddSubscriptionPriceParamsPriceNewSubscriptionMatrixPriceBillingCycleConfigurationDurationUnitDay

// This is an alias to an internal value.
const AddSubscriptionPriceParamsPriceNewSubscriptionMatrixPriceBillingCycleConfigurationDurationUnitMonth = shared.AddSubscriptionPriceParamsPriceNewSubscriptionMatrixPriceBillingCycleConfigurationDurationUnitMonth

// Within each billing cycle, specifies the cadence at which invoices are produced.
// If unspecified, a single invoice is produced per billing cycle.
//
// This is an alias to an internal type.
type AddSubscriptionPriceParamsPriceNewSubscriptionMatrixPriceInvoicingCycleConfiguration = shared.AddSubscriptionPriceParamsPriceNewSubscriptionMatrixPriceInvoicingCycleConfiguration

// The unit of billing period duration.
//
// This is an alias to an internal type.
type AddSubscriptionPriceParamsPriceNewSubscriptionMatrixPriceInvoicingCycleConfigurationDurationUnit = shared.AddSubscriptionPriceParamsPriceNewSubscriptionMatrixPriceInvoicingCycleConfigurationDurationUnit

// This is an alias to an internal value.
const AddSubscriptionPriceParamsPriceNewSubscriptionMatrixPriceInvoicingCycleConfigurationDurationUnitDay = shared.AddSubscriptionPriceParamsPriceNewSubscriptionMatrixPriceInvoicingCycleConfigurationDurationUnitDay

// This is an alias to an internal value.
const AddSubscriptionPriceParamsPriceNewSubscriptionMatrixPriceInvoicingCycleConfigurationDurationUnitMonth = shared.AddSubscriptionPriceParamsPriceNewSubscriptionMatrixPriceInvoicingCycleConfigurationDurationUnitMonth

// This is an alias to an internal type.
type AddSubscriptionPriceParamsPriceNewSubscriptionTieredPrice = shared.AddSubscriptionPriceParamsPriceNewSubscriptionTieredPrice

// The cadence to bill for this price on.
//
// This is an alias to an internal type.
type AddSubscriptionPriceParamsPriceNewSubscriptionTieredPriceCadence = shared.AddSubscriptionPriceParamsPriceNewSubscriptionTieredPriceCadence

// This is an alias to an internal value.
const AddSubscriptionPriceParamsPriceNewSubscriptionTieredPriceCadenceAnnual = shared.AddSubscriptionPriceParamsPriceNewSubscriptionTieredPriceCadenceAnnual

// This is an alias to an internal value.
const AddSubscriptionPriceParamsPriceNewSubscriptionTieredPriceCadenceSemiAnnual = shared.AddSubscriptionPriceParamsPriceNewSubscriptionTieredPriceCadenceSemiAnnual

// This is an alias to an internal value.
const AddSubscriptionPriceParamsPriceNewSubscriptionTieredPriceCadenceMonthly = shared.AddSubscriptionPriceParamsPriceNewSubscriptionTieredPriceCadenceMonthly

// This is an alias to an internal value.
const AddSubscriptionPriceParamsPriceNewSubscriptionTieredPriceCadenceQuarterly = shared.AddSubscriptionPriceParamsPriceNewSubscriptionTieredPriceCadenceQuarterly

// This is an alias to an internal value.
const AddSubscriptionPriceParamsPriceNewSubscriptionTieredPriceCadenceOneTime = shared.AddSubscriptionPriceParamsPriceNewSubscriptionTieredPriceCadenceOneTime

// This is an alias to an internal value.
const AddSubscriptionPriceParamsPriceNewSubscriptionTieredPriceCadenceCustom = shared.AddSubscriptionPriceParamsPriceNewSubscriptionTieredPriceCadenceCustom

// This is an alias to an internal type.
type AddSubscriptionPriceParamsPriceNewSubscriptionTieredPriceModelType = shared.AddSubscriptionPriceParamsPriceNewSubscriptionTieredPriceModelType

// This is an alias to an internal value.
const AddSubscriptionPriceParamsPriceNewSubscriptionTieredPriceModelTypeTiered = shared.AddSubscriptionPriceParamsPriceNewSubscriptionTieredPriceModelTypeTiered

// This is an alias to an internal type.
type AddSubscriptionPriceParamsPriceNewSubscriptionTieredPriceTieredConfig = shared.AddSubscriptionPriceParamsPriceNewSubscriptionTieredPriceTieredConfig

// This is an alias to an internal type.
type AddSubscriptionPriceParamsPriceNewSubscriptionTieredPriceTieredConfigTier = shared.AddSubscriptionPriceParamsPriceNewSubscriptionTieredPriceTieredConfigTier

// For custom cadence: specifies the duration of the billing period in days or
// months.
//
// This is an alias to an internal type.
type AddSubscriptionPriceParamsPriceNewSubscriptionTieredPriceBillingCycleConfiguration = shared.AddSubscriptionPriceParamsPriceNewSubscriptionTieredPriceBillingCycleConfiguration

// The unit of billing period duration.
//
// This is an alias to an internal type.
type AddSubscriptionPriceParamsPriceNewSubscriptionTieredPriceBillingCycleConfigurationDurationUnit = shared.AddSubscriptionPriceParamsPriceNewSubscriptionTieredPriceBillingCycleConfigurationDurationUnit

// This is an alias to an internal value.
const AddSubscriptionPriceParamsPriceNewSubscriptionTieredPriceBillingCycleConfigurationDurationUnitDay = shared.AddSubscriptionPriceParamsPriceNewSubscriptionTieredPriceBillingCycleConfigurationDurationUnitDay

// This is an alias to an internal value.
const AddSubscriptionPriceParamsPriceNewSubscriptionTieredPriceBillingCycleConfigurationDurationUnitMonth = shared.AddSubscriptionPriceParamsPriceNewSubscriptionTieredPriceBillingCycleConfigurationDurationUnitMonth

// Within each billing cycle, specifies the cadence at which invoices are produced.
// If unspecified, a single invoice is produced per billing cycle.
//
// This is an alias to an internal type.
type AddSubscriptionPriceParamsPriceNewSubscriptionTieredPriceInvoicingCycleConfiguration = shared.AddSubscriptionPriceParamsPriceNewSubscriptionTieredPriceInvoicingCycleConfiguration

// The unit of billing period duration.
//
// This is an alias to an internal type.
type AddSubscriptionPriceParamsPriceNewSubscriptionTieredPriceInvoicingCycleConfigurationDurationUnit = shared.AddSubscriptionPriceParamsPriceNewSubscriptionTieredPriceInvoicingCycleConfigurationDurationUnit

// This is an alias to an internal value.
const AddSubscriptionPriceParamsPriceNewSubscriptionTieredPriceInvoicingCycleConfigurationDurationUnitDay = shared.AddSubscriptionPriceParamsPriceNewSubscriptionTieredPriceInvoicingCycleConfigurationDurationUnitDay

// This is an alias to an internal value.
const AddSubscriptionPriceParamsPriceNewSubscriptionTieredPriceInvoicingCycleConfigurationDurationUnitMonth = shared.AddSubscriptionPriceParamsPriceNewSubscriptionTieredPriceInvoicingCycleConfigurationDurationUnitMonth

// This is an alias to an internal type.
type AddSubscriptionPriceParamsPriceNewSubscriptionTieredBpsPrice = shared.AddSubscriptionPriceParamsPriceNewSubscriptionTieredBpsPrice

// The cadence to bill for this price on.
//
// This is an alias to an internal type.
type AddSubscriptionPriceParamsPriceNewSubscriptionTieredBpsPriceCadence = shared.AddSubscriptionPriceParamsPriceNewSubscriptionTieredBpsPriceCadence

// This is an alias to an internal value.
const AddSubscriptionPriceParamsPriceNewSubscriptionTieredBpsPriceCadenceAnnual = shared.AddSubscriptionPriceParamsPriceNewSubscriptionTieredBpsPriceCadenceAnnual

// This is an alias to an internal value.
const AddSubscriptionPriceParamsPriceNewSubscriptionTieredBpsPriceCadenceSemiAnnual = shared.AddSubscriptionPriceParamsPriceNewSubscriptionTieredBpsPriceCadenceSemiAnnual

// This is an alias to an internal value.
const AddSubscriptionPriceParamsPriceNewSubscriptionTieredBpsPriceCadenceMonthly = shared.AddSubscriptionPriceParamsPriceNewSubscriptionTieredBpsPriceCadenceMonthly

// This is an alias to an internal value.
const AddSubscriptionPriceParamsPriceNewSubscriptionTieredBpsPriceCadenceQuarterly = shared.AddSubscriptionPriceParamsPriceNewSubscriptionTieredBpsPriceCadenceQuarterly

// This is an alias to an internal value.
const AddSubscriptionPriceParamsPriceNewSubscriptionTieredBpsPriceCadenceOneTime = shared.AddSubscriptionPriceParamsPriceNewSubscriptionTieredBpsPriceCadenceOneTime

// This is an alias to an internal value.
const AddSubscriptionPriceParamsPriceNewSubscriptionTieredBpsPriceCadenceCustom = shared.AddSubscriptionPriceParamsPriceNewSubscriptionTieredBpsPriceCadenceCustom

// This is an alias to an internal type.
type AddSubscriptionPriceParamsPriceNewSubscriptionTieredBpsPriceModelType = shared.AddSubscriptionPriceParamsPriceNewSubscriptionTieredBpsPriceModelType

// This is an alias to an internal value.
const AddSubscriptionPriceParamsPriceNewSubscriptionTieredBpsPriceModelTypeTieredBps = shared.AddSubscriptionPriceParamsPriceNewSubscriptionTieredBpsPriceModelTypeTieredBps

// This is an alias to an internal type.
type AddSubscriptionPriceParamsPriceNewSubscriptionTieredBpsPriceTieredBpsConfig = shared.AddSubscriptionPriceParamsPriceNewSubscriptionTieredBpsPriceTieredBpsConfig

// This is an alias to an internal type.
type AddSubscriptionPriceParamsPriceNewSubscriptionTieredBpsPriceTieredBpsConfigTier = shared.AddSubscriptionPriceParamsPriceNewSubscriptionTieredBpsPriceTieredBpsConfigTier

// For custom cadence: specifies the duration of the billing period in days or
// months.
//
// This is an alias to an internal type.
type AddSubscriptionPriceParamsPriceNewSubscriptionTieredBpsPriceBillingCycleConfiguration = shared.AddSubscriptionPriceParamsPriceNewSubscriptionTieredBpsPriceBillingCycleConfiguration

// The unit of billing period duration.
//
// This is an alias to an internal type.
type AddSubscriptionPriceParamsPriceNewSubscriptionTieredBpsPriceBillingCycleConfigurationDurationUnit = shared.AddSubscriptionPriceParamsPriceNewSubscriptionTieredBpsPriceBillingCycleConfigurationDurationUnit

// This is an alias to an internal value.
const AddSubscriptionPriceParamsPriceNewSubscriptionTieredBpsPriceBillingCycleConfigurationDurationUnitDay = shared.AddSubscriptionPriceParamsPriceNewSubscriptionTieredBpsPriceBillingCycleConfigurationDurationUnitDay

// This is an alias to an internal value.
const AddSubscriptionPriceParamsPriceNewSubscriptionTieredBpsPriceBillingCycleConfigurationDurationUnitMonth = shared.AddSubscriptionPriceParamsPriceNewSubscriptionTieredBpsPriceBillingCycleConfigurationDurationUnitMonth

// Within each billing cycle, specifies the cadence at which invoices are produced.
// If unspecified, a single invoice is produced per billing cycle.
//
// This is an alias to an internal type.
type AddSubscriptionPriceParamsPriceNewSubscriptionTieredBpsPriceInvoicingCycleConfiguration = shared.AddSubscriptionPriceParamsPriceNewSubscriptionTieredBpsPriceInvoicingCycleConfiguration

// The unit of billing period duration.
//
// This is an alias to an internal type.
type AddSubscriptionPriceParamsPriceNewSubscriptionTieredBpsPriceInvoicingCycleConfigurationDurationUnit = shared.AddSubscriptionPriceParamsPriceNewSubscriptionTieredBpsPriceInvoicingCycleConfigurationDurationUnit

// This is an alias to an internal value.
const AddSubscriptionPriceParamsPriceNewSubscriptionTieredBpsPriceInvoicingCycleConfigurationDurationUnitDay = shared.AddSubscriptionPriceParamsPriceNewSubscriptionTieredBpsPriceInvoicingCycleConfigurationDurationUnitDay

// This is an alias to an internal value.
const AddSubscriptionPriceParamsPriceNewSubscriptionTieredBpsPriceInvoicingCycleConfigurationDurationUnitMonth = shared.AddSubscriptionPriceParamsPriceNewSubscriptionTieredBpsPriceInvoicingCycleConfigurationDurationUnitMonth

// This is an alias to an internal type.
type AddSubscriptionPriceParamsPriceNewSubscriptionBpsPrice = shared.AddSubscriptionPriceParamsPriceNewSubscriptionBpsPrice

// This is an alias to an internal type.
type AddSubscriptionPriceParamsPriceNewSubscriptionBpsPriceBpsConfig = shared.AddSubscriptionPriceParamsPriceNewSubscriptionBpsPriceBpsConfig

// The cadence to bill for this price on.
//
// This is an alias to an internal type.
type AddSubscriptionPriceParamsPriceNewSubscriptionBpsPriceCadence = shared.AddSubscriptionPriceParamsPriceNewSubscriptionBpsPriceCadence

// This is an alias to an internal value.
const AddSubscriptionPriceParamsPriceNewSubscriptionBpsPriceCadenceAnnual = shared.AddSubscriptionPriceParamsPriceNewSubscriptionBpsPriceCadenceAnnual

// This is an alias to an internal value.
const AddSubscriptionPriceParamsPriceNewSubscriptionBpsPriceCadenceSemiAnnual = shared.AddSubscriptionPriceParamsPriceNewSubscriptionBpsPriceCadenceSemiAnnual

// This is an alias to an internal value.
const AddSubscriptionPriceParamsPriceNewSubscriptionBpsPriceCadenceMonthly = shared.AddSubscriptionPriceParamsPriceNewSubscriptionBpsPriceCadenceMonthly

// This is an alias to an internal value.
const AddSubscriptionPriceParamsPriceNewSubscriptionBpsPriceCadenceQuarterly = shared.AddSubscriptionPriceParamsPriceNewSubscriptionBpsPriceCadenceQuarterly

// This is an alias to an internal value.
const AddSubscriptionPriceParamsPriceNewSubscriptionBpsPriceCadenceOneTime = shared.AddSubscriptionPriceParamsPriceNewSubscriptionBpsPriceCadenceOneTime

// This is an alias to an internal value.
const AddSubscriptionPriceParamsPriceNewSubscriptionBpsPriceCadenceCustom = shared.AddSubscriptionPriceParamsPriceNewSubscriptionBpsPriceCadenceCustom

// This is an alias to an internal type.
type AddSubscriptionPriceParamsPriceNewSubscriptionBpsPriceModelType = shared.AddSubscriptionPriceParamsPriceNewSubscriptionBpsPriceModelType

// This is an alias to an internal value.
const AddSubscriptionPriceParamsPriceNewSubscriptionBpsPriceModelTypeBps = shared.AddSubscriptionPriceParamsPriceNewSubscriptionBpsPriceModelTypeBps

// For custom cadence: specifies the duration of the billing period in days or
// months.
//
// This is an alias to an internal type.
type AddSubscriptionPriceParamsPriceNewSubscriptionBpsPriceBillingCycleConfiguration = shared.AddSubscriptionPriceParamsPriceNewSubscriptionBpsPriceBillingCycleConfiguration

// The unit of billing period duration.
//
// This is an alias to an internal type.
type AddSubscriptionPriceParamsPriceNewSubscriptionBpsPriceBillingCycleConfigurationDurationUnit = shared.AddSubscriptionPriceParamsPriceNewSubscriptionBpsPriceBillingCycleConfigurationDurationUnit

// This is an alias to an internal value.
const AddSubscriptionPriceParamsPriceNewSubscriptionBpsPriceBillingCycleConfigurationDurationUnitDay = shared.AddSubscriptionPriceParamsPriceNewSubscriptionBpsPriceBillingCycleConfigurationDurationUnitDay

// This is an alias to an internal value.
const AddSubscriptionPriceParamsPriceNewSubscriptionBpsPriceBillingCycleConfigurationDurationUnitMonth = shared.AddSubscriptionPriceParamsPriceNewSubscriptionBpsPriceBillingCycleConfigurationDurationUnitMonth

// Within each billing cycle, specifies the cadence at which invoices are produced.
// If unspecified, a single invoice is produced per billing cycle.
//
// This is an alias to an internal type.
type AddSubscriptionPriceParamsPriceNewSubscriptionBpsPriceInvoicingCycleConfiguration = shared.AddSubscriptionPriceParamsPriceNewSubscriptionBpsPriceInvoicingCycleConfiguration

// The unit of billing period duration.
//
// This is an alias to an internal type.
type AddSubscriptionPriceParamsPriceNewSubscriptionBpsPriceInvoicingCycleConfigurationDurationUnit = shared.AddSubscriptionPriceParamsPriceNewSubscriptionBpsPriceInvoicingCycleConfigurationDurationUnit

// This is an alias to an internal value.
const AddSubscriptionPriceParamsPriceNewSubscriptionBpsPriceInvoicingCycleConfigurationDurationUnitDay = shared.AddSubscriptionPriceParamsPriceNewSubscriptionBpsPriceInvoicingCycleConfigurationDurationUnitDay

// This is an alias to an internal value.
const AddSubscriptionPriceParamsPriceNewSubscriptionBpsPriceInvoicingCycleConfigurationDurationUnitMonth = shared.AddSubscriptionPriceParamsPriceNewSubscriptionBpsPriceInvoicingCycleConfigurationDurationUnitMonth

// This is an alias to an internal type.
type AddSubscriptionPriceParamsPriceNewSubscriptionBulkBpsPrice = shared.AddSubscriptionPriceParamsPriceNewSubscriptionBulkBpsPrice

// This is an alias to an internal type.
type AddSubscriptionPriceParamsPriceNewSubscriptionBulkBpsPriceBulkBpsConfig = shared.AddSubscriptionPriceParamsPriceNewSubscriptionBulkBpsPriceBulkBpsConfig

// This is an alias to an internal type.
type AddSubscriptionPriceParamsPriceNewSubscriptionBulkBpsPriceBulkBpsConfigTier = shared.AddSubscriptionPriceParamsPriceNewSubscriptionBulkBpsPriceBulkBpsConfigTier

// The cadence to bill for this price on.
//
// This is an alias to an internal type.
type AddSubscriptionPriceParamsPriceNewSubscriptionBulkBpsPriceCadence = shared.AddSubscriptionPriceParamsPriceNewSubscriptionBulkBpsPriceCadence

// This is an alias to an internal value.
const AddSubscriptionPriceParamsPriceNewSubscriptionBulkBpsPriceCadenceAnnual = shared.AddSubscriptionPriceParamsPriceNewSubscriptionBulkBpsPriceCadenceAnnual

// This is an alias to an internal value.
const AddSubscriptionPriceParamsPriceNewSubscriptionBulkBpsPriceCadenceSemiAnnual = shared.AddSubscriptionPriceParamsPriceNewSubscriptionBulkBpsPriceCadenceSemiAnnual

// This is an alias to an internal value.
const AddSubscriptionPriceParamsPriceNewSubscriptionBulkBpsPriceCadenceMonthly = shared.AddSubscriptionPriceParamsPriceNewSubscriptionBulkBpsPriceCadenceMonthly

// This is an alias to an internal value.
const AddSubscriptionPriceParamsPriceNewSubscriptionBulkBpsPriceCadenceQuarterly = shared.AddSubscriptionPriceParamsPriceNewSubscriptionBulkBpsPriceCadenceQuarterly

// This is an alias to an internal value.
const AddSubscriptionPriceParamsPriceNewSubscriptionBulkBpsPriceCadenceOneTime = shared.AddSubscriptionPriceParamsPriceNewSubscriptionBulkBpsPriceCadenceOneTime

// This is an alias to an internal value.
const AddSubscriptionPriceParamsPriceNewSubscriptionBulkBpsPriceCadenceCustom = shared.AddSubscriptionPriceParamsPriceNewSubscriptionBulkBpsPriceCadenceCustom

// This is an alias to an internal type.
type AddSubscriptionPriceParamsPriceNewSubscriptionBulkBpsPriceModelType = shared.AddSubscriptionPriceParamsPriceNewSubscriptionBulkBpsPriceModelType

// This is an alias to an internal value.
const AddSubscriptionPriceParamsPriceNewSubscriptionBulkBpsPriceModelTypeBulkBps = shared.AddSubscriptionPriceParamsPriceNewSubscriptionBulkBpsPriceModelTypeBulkBps

// For custom cadence: specifies the duration of the billing period in days or
// months.
//
// This is an alias to an internal type.
type AddSubscriptionPriceParamsPriceNewSubscriptionBulkBpsPriceBillingCycleConfiguration = shared.AddSubscriptionPriceParamsPriceNewSubscriptionBulkBpsPriceBillingCycleConfiguration

// The unit of billing period duration.
//
// This is an alias to an internal type.
type AddSubscriptionPriceParamsPriceNewSubscriptionBulkBpsPriceBillingCycleConfigurationDurationUnit = shared.AddSubscriptionPriceParamsPriceNewSubscriptionBulkBpsPriceBillingCycleConfigurationDurationUnit

// This is an alias to an internal value.
const AddSubscriptionPriceParamsPriceNewSubscriptionBulkBpsPriceBillingCycleConfigurationDurationUnitDay = shared.AddSubscriptionPriceParamsPriceNewSubscriptionBulkBpsPriceBillingCycleConfigurationDurationUnitDay

// This is an alias to an internal value.
const AddSubscriptionPriceParamsPriceNewSubscriptionBulkBpsPriceBillingCycleConfigurationDurationUnitMonth = shared.AddSubscriptionPriceParamsPriceNewSubscriptionBulkBpsPriceBillingCycleConfigurationDurationUnitMonth

// Within each billing cycle, specifies the cadence at which invoices are produced.
// If unspecified, a single invoice is produced per billing cycle.
//
// This is an alias to an internal type.
type AddSubscriptionPriceParamsPriceNewSubscriptionBulkBpsPriceInvoicingCycleConfiguration = shared.AddSubscriptionPriceParamsPriceNewSubscriptionBulkBpsPriceInvoicingCycleConfiguration

// The unit of billing period duration.
//
// This is an alias to an internal type.
type AddSubscriptionPriceParamsPriceNewSubscriptionBulkBpsPriceInvoicingCycleConfigurationDurationUnit = shared.AddSubscriptionPriceParamsPriceNewSubscriptionBulkBpsPriceInvoicingCycleConfigurationDurationUnit

// This is an alias to an internal value.
const AddSubscriptionPriceParamsPriceNewSubscriptionBulkBpsPriceInvoicingCycleConfigurationDurationUnitDay = shared.AddSubscriptionPriceParamsPriceNewSubscriptionBulkBpsPriceInvoicingCycleConfigurationDurationUnitDay

// This is an alias to an internal value.
const AddSubscriptionPriceParamsPriceNewSubscriptionBulkBpsPriceInvoicingCycleConfigurationDurationUnitMonth = shared.AddSubscriptionPriceParamsPriceNewSubscriptionBulkBpsPriceInvoicingCycleConfigurationDurationUnitMonth

// This is an alias to an internal type.
type AddSubscriptionPriceParamsPriceNewSubscriptionBulkPrice = shared.AddSubscriptionPriceParamsPriceNewSubscriptionBulkPrice

// This is an alias to an internal type.
type AddSubscriptionPriceParamsPriceNewSubscriptionBulkPriceBulkConfig = shared.AddSubscriptionPriceParamsPriceNewSubscriptionBulkPriceBulkConfig

// This is an alias to an internal type.
type AddSubscriptionPriceParamsPriceNewSubscriptionBulkPriceBulkConfigTier = shared.AddSubscriptionPriceParamsPriceNewSubscriptionBulkPriceBulkConfigTier

// The cadence to bill for this price on.
//
// This is an alias to an internal type.
type AddSubscriptionPriceParamsPriceNewSubscriptionBulkPriceCadence = shared.AddSubscriptionPriceParamsPriceNewSubscriptionBulkPriceCadence

// This is an alias to an internal value.
const AddSubscriptionPriceParamsPriceNewSubscriptionBulkPriceCadenceAnnual = shared.AddSubscriptionPriceParamsPriceNewSubscriptionBulkPriceCadenceAnnual

// This is an alias to an internal value.
const AddSubscriptionPriceParamsPriceNewSubscriptionBulkPriceCadenceSemiAnnual = shared.AddSubscriptionPriceParamsPriceNewSubscriptionBulkPriceCadenceSemiAnnual

// This is an alias to an internal value.
const AddSubscriptionPriceParamsPriceNewSubscriptionBulkPriceCadenceMonthly = shared.AddSubscriptionPriceParamsPriceNewSubscriptionBulkPriceCadenceMonthly

// This is an alias to an internal value.
const AddSubscriptionPriceParamsPriceNewSubscriptionBulkPriceCadenceQuarterly = shared.AddSubscriptionPriceParamsPriceNewSubscriptionBulkPriceCadenceQuarterly

// This is an alias to an internal value.
const AddSubscriptionPriceParamsPriceNewSubscriptionBulkPriceCadenceOneTime = shared.AddSubscriptionPriceParamsPriceNewSubscriptionBulkPriceCadenceOneTime

// This is an alias to an internal value.
const AddSubscriptionPriceParamsPriceNewSubscriptionBulkPriceCadenceCustom = shared.AddSubscriptionPriceParamsPriceNewSubscriptionBulkPriceCadenceCustom

// This is an alias to an internal type.
type AddSubscriptionPriceParamsPriceNewSubscriptionBulkPriceModelType = shared.AddSubscriptionPriceParamsPriceNewSubscriptionBulkPriceModelType

// This is an alias to an internal value.
const AddSubscriptionPriceParamsPriceNewSubscriptionBulkPriceModelTypeBulk = shared.AddSubscriptionPriceParamsPriceNewSubscriptionBulkPriceModelTypeBulk

// For custom cadence: specifies the duration of the billing period in days or
// months.
//
// This is an alias to an internal type.
type AddSubscriptionPriceParamsPriceNewSubscriptionBulkPriceBillingCycleConfiguration = shared.AddSubscriptionPriceParamsPriceNewSubscriptionBulkPriceBillingCycleConfiguration

// The unit of billing period duration.
//
// This is an alias to an internal type.
type AddSubscriptionPriceParamsPriceNewSubscriptionBulkPriceBillingCycleConfigurationDurationUnit = shared.AddSubscriptionPriceParamsPriceNewSubscriptionBulkPriceBillingCycleConfigurationDurationUnit

// This is an alias to an internal value.
const AddSubscriptionPriceParamsPriceNewSubscriptionBulkPriceBillingCycleConfigurationDurationUnitDay = shared.AddSubscriptionPriceParamsPriceNewSubscriptionBulkPriceBillingCycleConfigurationDurationUnitDay

// This is an alias to an internal value.
const AddSubscriptionPriceParamsPriceNewSubscriptionBulkPriceBillingCycleConfigurationDurationUnitMonth = shared.AddSubscriptionPriceParamsPriceNewSubscriptionBulkPriceBillingCycleConfigurationDurationUnitMonth

// Within each billing cycle, specifies the cadence at which invoices are produced.
// If unspecified, a single invoice is produced per billing cycle.
//
// This is an alias to an internal type.
type AddSubscriptionPriceParamsPriceNewSubscriptionBulkPriceInvoicingCycleConfiguration = shared.AddSubscriptionPriceParamsPriceNewSubscriptionBulkPriceInvoicingCycleConfiguration

// The unit of billing period duration.
//
// This is an alias to an internal type.
type AddSubscriptionPriceParamsPriceNewSubscriptionBulkPriceInvoicingCycleConfigurationDurationUnit = shared.AddSubscriptionPriceParamsPriceNewSubscriptionBulkPriceInvoicingCycleConfigurationDurationUnit

// This is an alias to an internal value.
const AddSubscriptionPriceParamsPriceNewSubscriptionBulkPriceInvoicingCycleConfigurationDurationUnitDay = shared.AddSubscriptionPriceParamsPriceNewSubscriptionBulkPriceInvoicingCycleConfigurationDurationUnitDay

// This is an alias to an internal value.
const AddSubscriptionPriceParamsPriceNewSubscriptionBulkPriceInvoicingCycleConfigurationDurationUnitMonth = shared.AddSubscriptionPriceParamsPriceNewSubscriptionBulkPriceInvoicingCycleConfigurationDurationUnitMonth

// This is an alias to an internal type.
type AddSubscriptionPriceParamsPriceNewSubscriptionThresholdTotalAmountPrice = shared.AddSubscriptionPriceParamsPriceNewSubscriptionThresholdTotalAmountPrice

// The cadence to bill for this price on.
//
// This is an alias to an internal type.
type AddSubscriptionPriceParamsPriceNewSubscriptionThresholdTotalAmountPriceCadence = shared.AddSubscriptionPriceParamsPriceNewSubscriptionThresholdTotalAmountPriceCadence

// This is an alias to an internal value.
const AddSubscriptionPriceParamsPriceNewSubscriptionThresholdTotalAmountPriceCadenceAnnual = shared.AddSubscriptionPriceParamsPriceNewSubscriptionThresholdTotalAmountPriceCadenceAnnual

// This is an alias to an internal value.
const AddSubscriptionPriceParamsPriceNewSubscriptionThresholdTotalAmountPriceCadenceSemiAnnual = shared.AddSubscriptionPriceParamsPriceNewSubscriptionThresholdTotalAmountPriceCadenceSemiAnnual

// This is an alias to an internal value.
const AddSubscriptionPriceParamsPriceNewSubscriptionThresholdTotalAmountPriceCadenceMonthly = shared.AddSubscriptionPriceParamsPriceNewSubscriptionThresholdTotalAmountPriceCadenceMonthly

// This is an alias to an internal value.
const AddSubscriptionPriceParamsPriceNewSubscriptionThresholdTotalAmountPriceCadenceQuarterly = shared.AddSubscriptionPriceParamsPriceNewSubscriptionThresholdTotalAmountPriceCadenceQuarterly

// This is an alias to an internal value.
const AddSubscriptionPriceParamsPriceNewSubscriptionThresholdTotalAmountPriceCadenceOneTime = shared.AddSubscriptionPriceParamsPriceNewSubscriptionThresholdTotalAmountPriceCadenceOneTime

// This is an alias to an internal value.
const AddSubscriptionPriceParamsPriceNewSubscriptionThresholdTotalAmountPriceCadenceCustom = shared.AddSubscriptionPriceParamsPriceNewSubscriptionThresholdTotalAmountPriceCadenceCustom

// This is an alias to an internal type.
type AddSubscriptionPriceParamsPriceNewSubscriptionThresholdTotalAmountPriceModelType = shared.AddSubscriptionPriceParamsPriceNewSubscriptionThresholdTotalAmountPriceModelType

// This is an alias to an internal value.
const AddSubscriptionPriceParamsPriceNewSubscriptionThresholdTotalAmountPriceModelTypeThresholdTotalAmount = shared.AddSubscriptionPriceParamsPriceNewSubscriptionThresholdTotalAmountPriceModelTypeThresholdTotalAmount

// For custom cadence: specifies the duration of the billing period in days or
// months.
//
// This is an alias to an internal type.
type AddSubscriptionPriceParamsPriceNewSubscriptionThresholdTotalAmountPriceBillingCycleConfiguration = shared.AddSubscriptionPriceParamsPriceNewSubscriptionThresholdTotalAmountPriceBillingCycleConfiguration

// The unit of billing period duration.
//
// This is an alias to an internal type.
type AddSubscriptionPriceParamsPriceNewSubscriptionThresholdTotalAmountPriceBillingCycleConfigurationDurationUnit = shared.AddSubscriptionPriceParamsPriceNewSubscriptionThresholdTotalAmountPriceBillingCycleConfigurationDurationUnit

// This is an alias to an internal value.
const AddSubscriptionPriceParamsPriceNewSubscriptionThresholdTotalAmountPriceBillingCycleConfigurationDurationUnitDay = shared.AddSubscriptionPriceParamsPriceNewSubscriptionThresholdTotalAmountPriceBillingCycleConfigurationDurationUnitDay

// This is an alias to an internal value.
const AddSubscriptionPriceParamsPriceNewSubscriptionThresholdTotalAmountPriceBillingCycleConfigurationDurationUnitMonth = shared.AddSubscriptionPriceParamsPriceNewSubscriptionThresholdTotalAmountPriceBillingCycleConfigurationDurationUnitMonth

// Within each billing cycle, specifies the cadence at which invoices are produced.
// If unspecified, a single invoice is produced per billing cycle.
//
// This is an alias to an internal type.
type AddSubscriptionPriceParamsPriceNewSubscriptionThresholdTotalAmountPriceInvoicingCycleConfiguration = shared.AddSubscriptionPriceParamsPriceNewSubscriptionThresholdTotalAmountPriceInvoicingCycleConfiguration

// The unit of billing period duration.
//
// This is an alias to an internal type.
type AddSubscriptionPriceParamsPriceNewSubscriptionThresholdTotalAmountPriceInvoicingCycleConfigurationDurationUnit = shared.AddSubscriptionPriceParamsPriceNewSubscriptionThresholdTotalAmountPriceInvoicingCycleConfigurationDurationUnit

// This is an alias to an internal value.
const AddSubscriptionPriceParamsPriceNewSubscriptionThresholdTotalAmountPriceInvoicingCycleConfigurationDurationUnitDay = shared.AddSubscriptionPriceParamsPriceNewSubscriptionThresholdTotalAmountPriceInvoicingCycleConfigurationDurationUnitDay

// This is an alias to an internal value.
const AddSubscriptionPriceParamsPriceNewSubscriptionThresholdTotalAmountPriceInvoicingCycleConfigurationDurationUnitMonth = shared.AddSubscriptionPriceParamsPriceNewSubscriptionThresholdTotalAmountPriceInvoicingCycleConfigurationDurationUnitMonth

// This is an alias to an internal type.
type AddSubscriptionPriceParamsPriceNewSubscriptionTieredPackagePrice = shared.AddSubscriptionPriceParamsPriceNewSubscriptionTieredPackagePrice

// The cadence to bill for this price on.
//
// This is an alias to an internal type.
type AddSubscriptionPriceParamsPriceNewSubscriptionTieredPackagePriceCadence = shared.AddSubscriptionPriceParamsPriceNewSubscriptionTieredPackagePriceCadence

// This is an alias to an internal value.
const AddSubscriptionPriceParamsPriceNewSubscriptionTieredPackagePriceCadenceAnnual = shared.AddSubscriptionPriceParamsPriceNewSubscriptionTieredPackagePriceCadenceAnnual

// This is an alias to an internal value.
const AddSubscriptionPriceParamsPriceNewSubscriptionTieredPackagePriceCadenceSemiAnnual = shared.AddSubscriptionPriceParamsPriceNewSubscriptionTieredPackagePriceCadenceSemiAnnual

// This is an alias to an internal value.
const AddSubscriptionPriceParamsPriceNewSubscriptionTieredPackagePriceCadenceMonthly = shared.AddSubscriptionPriceParamsPriceNewSubscriptionTieredPackagePriceCadenceMonthly

// This is an alias to an internal value.
const AddSubscriptionPriceParamsPriceNewSubscriptionTieredPackagePriceCadenceQuarterly = shared.AddSubscriptionPriceParamsPriceNewSubscriptionTieredPackagePriceCadenceQuarterly

// This is an alias to an internal value.
const AddSubscriptionPriceParamsPriceNewSubscriptionTieredPackagePriceCadenceOneTime = shared.AddSubscriptionPriceParamsPriceNewSubscriptionTieredPackagePriceCadenceOneTime

// This is an alias to an internal value.
const AddSubscriptionPriceParamsPriceNewSubscriptionTieredPackagePriceCadenceCustom = shared.AddSubscriptionPriceParamsPriceNewSubscriptionTieredPackagePriceCadenceCustom

// This is an alias to an internal type.
type AddSubscriptionPriceParamsPriceNewSubscriptionTieredPackagePriceModelType = shared.AddSubscriptionPriceParamsPriceNewSubscriptionTieredPackagePriceModelType

// This is an alias to an internal value.
const AddSubscriptionPriceParamsPriceNewSubscriptionTieredPackagePriceModelTypeTieredPackage = shared.AddSubscriptionPriceParamsPriceNewSubscriptionTieredPackagePriceModelTypeTieredPackage

// For custom cadence: specifies the duration of the billing period in days or
// months.
//
// This is an alias to an internal type.
type AddSubscriptionPriceParamsPriceNewSubscriptionTieredPackagePriceBillingCycleConfiguration = shared.AddSubscriptionPriceParamsPriceNewSubscriptionTieredPackagePriceBillingCycleConfiguration

// The unit of billing period duration.
//
// This is an alias to an internal type.
type AddSubscriptionPriceParamsPriceNewSubscriptionTieredPackagePriceBillingCycleConfigurationDurationUnit = shared.AddSubscriptionPriceParamsPriceNewSubscriptionTieredPackagePriceBillingCycleConfigurationDurationUnit

// This is an alias to an internal value.
const AddSubscriptionPriceParamsPriceNewSubscriptionTieredPackagePriceBillingCycleConfigurationDurationUnitDay = shared.AddSubscriptionPriceParamsPriceNewSubscriptionTieredPackagePriceBillingCycleConfigurationDurationUnitDay

// This is an alias to an internal value.
const AddSubscriptionPriceParamsPriceNewSubscriptionTieredPackagePriceBillingCycleConfigurationDurationUnitMonth = shared.AddSubscriptionPriceParamsPriceNewSubscriptionTieredPackagePriceBillingCycleConfigurationDurationUnitMonth

// Within each billing cycle, specifies the cadence at which invoices are produced.
// If unspecified, a single invoice is produced per billing cycle.
//
// This is an alias to an internal type.
type AddSubscriptionPriceParamsPriceNewSubscriptionTieredPackagePriceInvoicingCycleConfiguration = shared.AddSubscriptionPriceParamsPriceNewSubscriptionTieredPackagePriceInvoicingCycleConfiguration

// The unit of billing period duration.
//
// This is an alias to an internal type.
type AddSubscriptionPriceParamsPriceNewSubscriptionTieredPackagePriceInvoicingCycleConfigurationDurationUnit = shared.AddSubscriptionPriceParamsPriceNewSubscriptionTieredPackagePriceInvoicingCycleConfigurationDurationUnit

// This is an alias to an internal value.
const AddSubscriptionPriceParamsPriceNewSubscriptionTieredPackagePriceInvoicingCycleConfigurationDurationUnitDay = shared.AddSubscriptionPriceParamsPriceNewSubscriptionTieredPackagePriceInvoicingCycleConfigurationDurationUnitDay

// This is an alias to an internal value.
const AddSubscriptionPriceParamsPriceNewSubscriptionTieredPackagePriceInvoicingCycleConfigurationDurationUnitMonth = shared.AddSubscriptionPriceParamsPriceNewSubscriptionTieredPackagePriceInvoicingCycleConfigurationDurationUnitMonth

// This is an alias to an internal type.
type AddSubscriptionPriceParamsPriceNewSubscriptionTieredWithMinimumPrice = shared.AddSubscriptionPriceParamsPriceNewSubscriptionTieredWithMinimumPrice

// The cadence to bill for this price on.
//
// This is an alias to an internal type.
type AddSubscriptionPriceParamsPriceNewSubscriptionTieredWithMinimumPriceCadence = shared.AddSubscriptionPriceParamsPriceNewSubscriptionTieredWithMinimumPriceCadence

// This is an alias to an internal value.
const AddSubscriptionPriceParamsPriceNewSubscriptionTieredWithMinimumPriceCadenceAnnual = shared.AddSubscriptionPriceParamsPriceNewSubscriptionTieredWithMinimumPriceCadenceAnnual

// This is an alias to an internal value.
const AddSubscriptionPriceParamsPriceNewSubscriptionTieredWithMinimumPriceCadenceSemiAnnual = shared.AddSubscriptionPriceParamsPriceNewSubscriptionTieredWithMinimumPriceCadenceSemiAnnual

// This is an alias to an internal value.
const AddSubscriptionPriceParamsPriceNewSubscriptionTieredWithMinimumPriceCadenceMonthly = shared.AddSubscriptionPriceParamsPriceNewSubscriptionTieredWithMinimumPriceCadenceMonthly

// This is an alias to an internal value.
const AddSubscriptionPriceParamsPriceNewSubscriptionTieredWithMinimumPriceCadenceQuarterly = shared.AddSubscriptionPriceParamsPriceNewSubscriptionTieredWithMinimumPriceCadenceQuarterly

// This is an alias to an internal value.
const AddSubscriptionPriceParamsPriceNewSubscriptionTieredWithMinimumPriceCadenceOneTime = shared.AddSubscriptionPriceParamsPriceNewSubscriptionTieredWithMinimumPriceCadenceOneTime

// This is an alias to an internal value.
const AddSubscriptionPriceParamsPriceNewSubscriptionTieredWithMinimumPriceCadenceCustom = shared.AddSubscriptionPriceParamsPriceNewSubscriptionTieredWithMinimumPriceCadenceCustom

// This is an alias to an internal type.
type AddSubscriptionPriceParamsPriceNewSubscriptionTieredWithMinimumPriceModelType = shared.AddSubscriptionPriceParamsPriceNewSubscriptionTieredWithMinimumPriceModelType

// This is an alias to an internal value.
const AddSubscriptionPriceParamsPriceNewSubscriptionTieredWithMinimumPriceModelTypeTieredWithMinimum = shared.AddSubscriptionPriceParamsPriceNewSubscriptionTieredWithMinimumPriceModelTypeTieredWithMinimum

// For custom cadence: specifies the duration of the billing period in days or
// months.
//
// This is an alias to an internal type.
type AddSubscriptionPriceParamsPriceNewSubscriptionTieredWithMinimumPriceBillingCycleConfiguration = shared.AddSubscriptionPriceParamsPriceNewSubscriptionTieredWithMinimumPriceBillingCycleConfiguration

// The unit of billing period duration.
//
// This is an alias to an internal type.
type AddSubscriptionPriceParamsPriceNewSubscriptionTieredWithMinimumPriceBillingCycleConfigurationDurationUnit = shared.AddSubscriptionPriceParamsPriceNewSubscriptionTieredWithMinimumPriceBillingCycleConfigurationDurationUnit

// This is an alias to an internal value.
const AddSubscriptionPriceParamsPriceNewSubscriptionTieredWithMinimumPriceBillingCycleConfigurationDurationUnitDay = shared.AddSubscriptionPriceParamsPriceNewSubscriptionTieredWithMinimumPriceBillingCycleConfigurationDurationUnitDay

// This is an alias to an internal value.
const AddSubscriptionPriceParamsPriceNewSubscriptionTieredWithMinimumPriceBillingCycleConfigurationDurationUnitMonth = shared.AddSubscriptionPriceParamsPriceNewSubscriptionTieredWithMinimumPriceBillingCycleConfigurationDurationUnitMonth

// Within each billing cycle, specifies the cadence at which invoices are produced.
// If unspecified, a single invoice is produced per billing cycle.
//
// This is an alias to an internal type.
type AddSubscriptionPriceParamsPriceNewSubscriptionTieredWithMinimumPriceInvoicingCycleConfiguration = shared.AddSubscriptionPriceParamsPriceNewSubscriptionTieredWithMinimumPriceInvoicingCycleConfiguration

// The unit of billing period duration.
//
// This is an alias to an internal type.
type AddSubscriptionPriceParamsPriceNewSubscriptionTieredWithMinimumPriceInvoicingCycleConfigurationDurationUnit = shared.AddSubscriptionPriceParamsPriceNewSubscriptionTieredWithMinimumPriceInvoicingCycleConfigurationDurationUnit

// This is an alias to an internal value.
const AddSubscriptionPriceParamsPriceNewSubscriptionTieredWithMinimumPriceInvoicingCycleConfigurationDurationUnitDay = shared.AddSubscriptionPriceParamsPriceNewSubscriptionTieredWithMinimumPriceInvoicingCycleConfigurationDurationUnitDay

// This is an alias to an internal value.
const AddSubscriptionPriceParamsPriceNewSubscriptionTieredWithMinimumPriceInvoicingCycleConfigurationDurationUnitMonth = shared.AddSubscriptionPriceParamsPriceNewSubscriptionTieredWithMinimumPriceInvoicingCycleConfigurationDurationUnitMonth

// This is an alias to an internal type.
type AddSubscriptionPriceParamsPriceNewSubscriptionUnitWithPercentPrice = shared.AddSubscriptionPriceParamsPriceNewSubscriptionUnitWithPercentPrice

// The cadence to bill for this price on.
//
// This is an alias to an internal type.
type AddSubscriptionPriceParamsPriceNewSubscriptionUnitWithPercentPriceCadence = shared.AddSubscriptionPriceParamsPriceNewSubscriptionUnitWithPercentPriceCadence

// This is an alias to an internal value.
const AddSubscriptionPriceParamsPriceNewSubscriptionUnitWithPercentPriceCadenceAnnual = shared.AddSubscriptionPriceParamsPriceNewSubscriptionUnitWithPercentPriceCadenceAnnual

// This is an alias to an internal value.
const AddSubscriptionPriceParamsPriceNewSubscriptionUnitWithPercentPriceCadenceSemiAnnual = shared.AddSubscriptionPriceParamsPriceNewSubscriptionUnitWithPercentPriceCadenceSemiAnnual

// This is an alias to an internal value.
const AddSubscriptionPriceParamsPriceNewSubscriptionUnitWithPercentPriceCadenceMonthly = shared.AddSubscriptionPriceParamsPriceNewSubscriptionUnitWithPercentPriceCadenceMonthly

// This is an alias to an internal value.
const AddSubscriptionPriceParamsPriceNewSubscriptionUnitWithPercentPriceCadenceQuarterly = shared.AddSubscriptionPriceParamsPriceNewSubscriptionUnitWithPercentPriceCadenceQuarterly

// This is an alias to an internal value.
const AddSubscriptionPriceParamsPriceNewSubscriptionUnitWithPercentPriceCadenceOneTime = shared.AddSubscriptionPriceParamsPriceNewSubscriptionUnitWithPercentPriceCadenceOneTime

// This is an alias to an internal value.
const AddSubscriptionPriceParamsPriceNewSubscriptionUnitWithPercentPriceCadenceCustom = shared.AddSubscriptionPriceParamsPriceNewSubscriptionUnitWithPercentPriceCadenceCustom

// This is an alias to an internal type.
type AddSubscriptionPriceParamsPriceNewSubscriptionUnitWithPercentPriceModelType = shared.AddSubscriptionPriceParamsPriceNewSubscriptionUnitWithPercentPriceModelType

// This is an alias to an internal value.
const AddSubscriptionPriceParamsPriceNewSubscriptionUnitWithPercentPriceModelTypeUnitWithPercent = shared.AddSubscriptionPriceParamsPriceNewSubscriptionUnitWithPercentPriceModelTypeUnitWithPercent

// For custom cadence: specifies the duration of the billing period in days or
// months.
//
// This is an alias to an internal type.
type AddSubscriptionPriceParamsPriceNewSubscriptionUnitWithPercentPriceBillingCycleConfiguration = shared.AddSubscriptionPriceParamsPriceNewSubscriptionUnitWithPercentPriceBillingCycleConfiguration

// The unit of billing period duration.
//
// This is an alias to an internal type.
type AddSubscriptionPriceParamsPriceNewSubscriptionUnitWithPercentPriceBillingCycleConfigurationDurationUnit = shared.AddSubscriptionPriceParamsPriceNewSubscriptionUnitWithPercentPriceBillingCycleConfigurationDurationUnit

// This is an alias to an internal value.
const AddSubscriptionPriceParamsPriceNewSubscriptionUnitWithPercentPriceBillingCycleConfigurationDurationUnitDay = shared.AddSubscriptionPriceParamsPriceNewSubscriptionUnitWithPercentPriceBillingCycleConfigurationDurationUnitDay

// This is an alias to an internal value.
const AddSubscriptionPriceParamsPriceNewSubscriptionUnitWithPercentPriceBillingCycleConfigurationDurationUnitMonth = shared.AddSubscriptionPriceParamsPriceNewSubscriptionUnitWithPercentPriceBillingCycleConfigurationDurationUnitMonth

// Within each billing cycle, specifies the cadence at which invoices are produced.
// If unspecified, a single invoice is produced per billing cycle.
//
// This is an alias to an internal type.
type AddSubscriptionPriceParamsPriceNewSubscriptionUnitWithPercentPriceInvoicingCycleConfiguration = shared.AddSubscriptionPriceParamsPriceNewSubscriptionUnitWithPercentPriceInvoicingCycleConfiguration

// The unit of billing period duration.
//
// This is an alias to an internal type.
type AddSubscriptionPriceParamsPriceNewSubscriptionUnitWithPercentPriceInvoicingCycleConfigurationDurationUnit = shared.AddSubscriptionPriceParamsPriceNewSubscriptionUnitWithPercentPriceInvoicingCycleConfigurationDurationUnit

// This is an alias to an internal value.
const AddSubscriptionPriceParamsPriceNewSubscriptionUnitWithPercentPriceInvoicingCycleConfigurationDurationUnitDay = shared.AddSubscriptionPriceParamsPriceNewSubscriptionUnitWithPercentPriceInvoicingCycleConfigurationDurationUnitDay

// This is an alias to an internal value.
const AddSubscriptionPriceParamsPriceNewSubscriptionUnitWithPercentPriceInvoicingCycleConfigurationDurationUnitMonth = shared.AddSubscriptionPriceParamsPriceNewSubscriptionUnitWithPercentPriceInvoicingCycleConfigurationDurationUnitMonth

// This is an alias to an internal type.
type AddSubscriptionPriceParamsPriceNewSubscriptionPackageWithAllocationPrice = shared.AddSubscriptionPriceParamsPriceNewSubscriptionPackageWithAllocationPrice

// The cadence to bill for this price on.
//
// This is an alias to an internal type.
type AddSubscriptionPriceParamsPriceNewSubscriptionPackageWithAllocationPriceCadence = shared.AddSubscriptionPriceParamsPriceNewSubscriptionPackageWithAllocationPriceCadence

// This is an alias to an internal value.
const AddSubscriptionPriceParamsPriceNewSubscriptionPackageWithAllocationPriceCadenceAnnual = shared.AddSubscriptionPriceParamsPriceNewSubscriptionPackageWithAllocationPriceCadenceAnnual

// This is an alias to an internal value.
const AddSubscriptionPriceParamsPriceNewSubscriptionPackageWithAllocationPriceCadenceSemiAnnual = shared.AddSubscriptionPriceParamsPriceNewSubscriptionPackageWithAllocationPriceCadenceSemiAnnual

// This is an alias to an internal value.
const AddSubscriptionPriceParamsPriceNewSubscriptionPackageWithAllocationPriceCadenceMonthly = shared.AddSubscriptionPriceParamsPriceNewSubscriptionPackageWithAllocationPriceCadenceMonthly

// This is an alias to an internal value.
const AddSubscriptionPriceParamsPriceNewSubscriptionPackageWithAllocationPriceCadenceQuarterly = shared.AddSubscriptionPriceParamsPriceNewSubscriptionPackageWithAllocationPriceCadenceQuarterly

// This is an alias to an internal value.
const AddSubscriptionPriceParamsPriceNewSubscriptionPackageWithAllocationPriceCadenceOneTime = shared.AddSubscriptionPriceParamsPriceNewSubscriptionPackageWithAllocationPriceCadenceOneTime

// This is an alias to an internal value.
const AddSubscriptionPriceParamsPriceNewSubscriptionPackageWithAllocationPriceCadenceCustom = shared.AddSubscriptionPriceParamsPriceNewSubscriptionPackageWithAllocationPriceCadenceCustom

// This is an alias to an internal type.
type AddSubscriptionPriceParamsPriceNewSubscriptionPackageWithAllocationPriceModelType = shared.AddSubscriptionPriceParamsPriceNewSubscriptionPackageWithAllocationPriceModelType

// This is an alias to an internal value.
const AddSubscriptionPriceParamsPriceNewSubscriptionPackageWithAllocationPriceModelTypePackageWithAllocation = shared.AddSubscriptionPriceParamsPriceNewSubscriptionPackageWithAllocationPriceModelTypePackageWithAllocation

// For custom cadence: specifies the duration of the billing period in days or
// months.
//
// This is an alias to an internal type.
type AddSubscriptionPriceParamsPriceNewSubscriptionPackageWithAllocationPriceBillingCycleConfiguration = shared.AddSubscriptionPriceParamsPriceNewSubscriptionPackageWithAllocationPriceBillingCycleConfiguration

// The unit of billing period duration.
//
// This is an alias to an internal type.
type AddSubscriptionPriceParamsPriceNewSubscriptionPackageWithAllocationPriceBillingCycleConfigurationDurationUnit = shared.AddSubscriptionPriceParamsPriceNewSubscriptionPackageWithAllocationPriceBillingCycleConfigurationDurationUnit

// This is an alias to an internal value.
const AddSubscriptionPriceParamsPriceNewSubscriptionPackageWithAllocationPriceBillingCycleConfigurationDurationUnitDay = shared.AddSubscriptionPriceParamsPriceNewSubscriptionPackageWithAllocationPriceBillingCycleConfigurationDurationUnitDay

// This is an alias to an internal value.
const AddSubscriptionPriceParamsPriceNewSubscriptionPackageWithAllocationPriceBillingCycleConfigurationDurationUnitMonth = shared.AddSubscriptionPriceParamsPriceNewSubscriptionPackageWithAllocationPriceBillingCycleConfigurationDurationUnitMonth

// Within each billing cycle, specifies the cadence at which invoices are produced.
// If unspecified, a single invoice is produced per billing cycle.
//
// This is an alias to an internal type.
type AddSubscriptionPriceParamsPriceNewSubscriptionPackageWithAllocationPriceInvoicingCycleConfiguration = shared.AddSubscriptionPriceParamsPriceNewSubscriptionPackageWithAllocationPriceInvoicingCycleConfiguration

// The unit of billing period duration.
//
// This is an alias to an internal type.
type AddSubscriptionPriceParamsPriceNewSubscriptionPackageWithAllocationPriceInvoicingCycleConfigurationDurationUnit = shared.AddSubscriptionPriceParamsPriceNewSubscriptionPackageWithAllocationPriceInvoicingCycleConfigurationDurationUnit

// This is an alias to an internal value.
const AddSubscriptionPriceParamsPriceNewSubscriptionPackageWithAllocationPriceInvoicingCycleConfigurationDurationUnitDay = shared.AddSubscriptionPriceParamsPriceNewSubscriptionPackageWithAllocationPriceInvoicingCycleConfigurationDurationUnitDay

// This is an alias to an internal value.
const AddSubscriptionPriceParamsPriceNewSubscriptionPackageWithAllocationPriceInvoicingCycleConfigurationDurationUnitMonth = shared.AddSubscriptionPriceParamsPriceNewSubscriptionPackageWithAllocationPriceInvoicingCycleConfigurationDurationUnitMonth

// This is an alias to an internal type.
type AddSubscriptionPriceParamsPriceNewSubscriptionTierWithProrationPrice = shared.AddSubscriptionPriceParamsPriceNewSubscriptionTierWithProrationPrice

// The cadence to bill for this price on.
//
// This is an alias to an internal type.
type AddSubscriptionPriceParamsPriceNewSubscriptionTierWithProrationPriceCadence = shared.AddSubscriptionPriceParamsPriceNewSubscriptionTierWithProrationPriceCadence

// This is an alias to an internal value.
const AddSubscriptionPriceParamsPriceNewSubscriptionTierWithProrationPriceCadenceAnnual = shared.AddSubscriptionPriceParamsPriceNewSubscriptionTierWithProrationPriceCadenceAnnual

// This is an alias to an internal value.
const AddSubscriptionPriceParamsPriceNewSubscriptionTierWithProrationPriceCadenceSemiAnnual = shared.AddSubscriptionPriceParamsPriceNewSubscriptionTierWithProrationPriceCadenceSemiAnnual

// This is an alias to an internal value.
const AddSubscriptionPriceParamsPriceNewSubscriptionTierWithProrationPriceCadenceMonthly = shared.AddSubscriptionPriceParamsPriceNewSubscriptionTierWithProrationPriceCadenceMonthly

// This is an alias to an internal value.
const AddSubscriptionPriceParamsPriceNewSubscriptionTierWithProrationPriceCadenceQuarterly = shared.AddSubscriptionPriceParamsPriceNewSubscriptionTierWithProrationPriceCadenceQuarterly

// This is an alias to an internal value.
const AddSubscriptionPriceParamsPriceNewSubscriptionTierWithProrationPriceCadenceOneTime = shared.AddSubscriptionPriceParamsPriceNewSubscriptionTierWithProrationPriceCadenceOneTime

// This is an alias to an internal value.
const AddSubscriptionPriceParamsPriceNewSubscriptionTierWithProrationPriceCadenceCustom = shared.AddSubscriptionPriceParamsPriceNewSubscriptionTierWithProrationPriceCadenceCustom

// This is an alias to an internal type.
type AddSubscriptionPriceParamsPriceNewSubscriptionTierWithProrationPriceModelType = shared.AddSubscriptionPriceParamsPriceNewSubscriptionTierWithProrationPriceModelType

// This is an alias to an internal value.
const AddSubscriptionPriceParamsPriceNewSubscriptionTierWithProrationPriceModelTypeTieredWithProration = shared.AddSubscriptionPriceParamsPriceNewSubscriptionTierWithProrationPriceModelTypeTieredWithProration

// For custom cadence: specifies the duration of the billing period in days or
// months.
//
// This is an alias to an internal type.
type AddSubscriptionPriceParamsPriceNewSubscriptionTierWithProrationPriceBillingCycleConfiguration = shared.AddSubscriptionPriceParamsPriceNewSubscriptionTierWithProrationPriceBillingCycleConfiguration

// The unit of billing period duration.
//
// This is an alias to an internal type.
type AddSubscriptionPriceParamsPriceNewSubscriptionTierWithProrationPriceBillingCycleConfigurationDurationUnit = shared.AddSubscriptionPriceParamsPriceNewSubscriptionTierWithProrationPriceBillingCycleConfigurationDurationUnit

// This is an alias to an internal value.
const AddSubscriptionPriceParamsPriceNewSubscriptionTierWithProrationPriceBillingCycleConfigurationDurationUnitDay = shared.AddSubscriptionPriceParamsPriceNewSubscriptionTierWithProrationPriceBillingCycleConfigurationDurationUnitDay

// This is an alias to an internal value.
const AddSubscriptionPriceParamsPriceNewSubscriptionTierWithProrationPriceBillingCycleConfigurationDurationUnitMonth = shared.AddSubscriptionPriceParamsPriceNewSubscriptionTierWithProrationPriceBillingCycleConfigurationDurationUnitMonth

// Within each billing cycle, specifies the cadence at which invoices are produced.
// If unspecified, a single invoice is produced per billing cycle.
//
// This is an alias to an internal type.
type AddSubscriptionPriceParamsPriceNewSubscriptionTierWithProrationPriceInvoicingCycleConfiguration = shared.AddSubscriptionPriceParamsPriceNewSubscriptionTierWithProrationPriceInvoicingCycleConfiguration

// The unit of billing period duration.
//
// This is an alias to an internal type.
type AddSubscriptionPriceParamsPriceNewSubscriptionTierWithProrationPriceInvoicingCycleConfigurationDurationUnit = shared.AddSubscriptionPriceParamsPriceNewSubscriptionTierWithProrationPriceInvoicingCycleConfigurationDurationUnit

// This is an alias to an internal value.
const AddSubscriptionPriceParamsPriceNewSubscriptionTierWithProrationPriceInvoicingCycleConfigurationDurationUnitDay = shared.AddSubscriptionPriceParamsPriceNewSubscriptionTierWithProrationPriceInvoicingCycleConfigurationDurationUnitDay

// This is an alias to an internal value.
const AddSubscriptionPriceParamsPriceNewSubscriptionTierWithProrationPriceInvoicingCycleConfigurationDurationUnitMonth = shared.AddSubscriptionPriceParamsPriceNewSubscriptionTierWithProrationPriceInvoicingCycleConfigurationDurationUnitMonth

// This is an alias to an internal type.
type AddSubscriptionPriceParamsPriceNewSubscriptionUnitWithProrationPrice = shared.AddSubscriptionPriceParamsPriceNewSubscriptionUnitWithProrationPrice

// The cadence to bill for this price on.
//
// This is an alias to an internal type.
type AddSubscriptionPriceParamsPriceNewSubscriptionUnitWithProrationPriceCadence = shared.AddSubscriptionPriceParamsPriceNewSubscriptionUnitWithProrationPriceCadence

// This is an alias to an internal value.
const AddSubscriptionPriceParamsPriceNewSubscriptionUnitWithProrationPriceCadenceAnnual = shared.AddSubscriptionPriceParamsPriceNewSubscriptionUnitWithProrationPriceCadenceAnnual

// This is an alias to an internal value.
const AddSubscriptionPriceParamsPriceNewSubscriptionUnitWithProrationPriceCadenceSemiAnnual = shared.AddSubscriptionPriceParamsPriceNewSubscriptionUnitWithProrationPriceCadenceSemiAnnual

// This is an alias to an internal value.
const AddSubscriptionPriceParamsPriceNewSubscriptionUnitWithProrationPriceCadenceMonthly = shared.AddSubscriptionPriceParamsPriceNewSubscriptionUnitWithProrationPriceCadenceMonthly

// This is an alias to an internal value.
const AddSubscriptionPriceParamsPriceNewSubscriptionUnitWithProrationPriceCadenceQuarterly = shared.AddSubscriptionPriceParamsPriceNewSubscriptionUnitWithProrationPriceCadenceQuarterly

// This is an alias to an internal value.
const AddSubscriptionPriceParamsPriceNewSubscriptionUnitWithProrationPriceCadenceOneTime = shared.AddSubscriptionPriceParamsPriceNewSubscriptionUnitWithProrationPriceCadenceOneTime

// This is an alias to an internal value.
const AddSubscriptionPriceParamsPriceNewSubscriptionUnitWithProrationPriceCadenceCustom = shared.AddSubscriptionPriceParamsPriceNewSubscriptionUnitWithProrationPriceCadenceCustom

// This is an alias to an internal type.
type AddSubscriptionPriceParamsPriceNewSubscriptionUnitWithProrationPriceModelType = shared.AddSubscriptionPriceParamsPriceNewSubscriptionUnitWithProrationPriceModelType

// This is an alias to an internal value.
const AddSubscriptionPriceParamsPriceNewSubscriptionUnitWithProrationPriceModelTypeUnitWithProration = shared.AddSubscriptionPriceParamsPriceNewSubscriptionUnitWithProrationPriceModelTypeUnitWithProration

// For custom cadence: specifies the duration of the billing period in days or
// months.
//
// This is an alias to an internal type.
type AddSubscriptionPriceParamsPriceNewSubscriptionUnitWithProrationPriceBillingCycleConfiguration = shared.AddSubscriptionPriceParamsPriceNewSubscriptionUnitWithProrationPriceBillingCycleConfiguration

// The unit of billing period duration.
//
// This is an alias to an internal type.
type AddSubscriptionPriceParamsPriceNewSubscriptionUnitWithProrationPriceBillingCycleConfigurationDurationUnit = shared.AddSubscriptionPriceParamsPriceNewSubscriptionUnitWithProrationPriceBillingCycleConfigurationDurationUnit

// This is an alias to an internal value.
const AddSubscriptionPriceParamsPriceNewSubscriptionUnitWithProrationPriceBillingCycleConfigurationDurationUnitDay = shared.AddSubscriptionPriceParamsPriceNewSubscriptionUnitWithProrationPriceBillingCycleConfigurationDurationUnitDay

// This is an alias to an internal value.
const AddSubscriptionPriceParamsPriceNewSubscriptionUnitWithProrationPriceBillingCycleConfigurationDurationUnitMonth = shared.AddSubscriptionPriceParamsPriceNewSubscriptionUnitWithProrationPriceBillingCycleConfigurationDurationUnitMonth

// Within each billing cycle, specifies the cadence at which invoices are produced.
// If unspecified, a single invoice is produced per billing cycle.
//
// This is an alias to an internal type.
type AddSubscriptionPriceParamsPriceNewSubscriptionUnitWithProrationPriceInvoicingCycleConfiguration = shared.AddSubscriptionPriceParamsPriceNewSubscriptionUnitWithProrationPriceInvoicingCycleConfiguration

// The unit of billing period duration.
//
// This is an alias to an internal type.
type AddSubscriptionPriceParamsPriceNewSubscriptionUnitWithProrationPriceInvoicingCycleConfigurationDurationUnit = shared.AddSubscriptionPriceParamsPriceNewSubscriptionUnitWithProrationPriceInvoicingCycleConfigurationDurationUnit

// This is an alias to an internal value.
const AddSubscriptionPriceParamsPriceNewSubscriptionUnitWithProrationPriceInvoicingCycleConfigurationDurationUnitDay = shared.AddSubscriptionPriceParamsPriceNewSubscriptionUnitWithProrationPriceInvoicingCycleConfigurationDurationUnitDay

// This is an alias to an internal value.
const AddSubscriptionPriceParamsPriceNewSubscriptionUnitWithProrationPriceInvoicingCycleConfigurationDurationUnitMonth = shared.AddSubscriptionPriceParamsPriceNewSubscriptionUnitWithProrationPriceInvoicingCycleConfigurationDurationUnitMonth

// This is an alias to an internal type.
type AddSubscriptionPriceParamsPriceNewSubscriptionGroupedAllocationPrice = shared.AddSubscriptionPriceParamsPriceNewSubscriptionGroupedAllocationPrice

// The cadence to bill for this price on.
//
// This is an alias to an internal type.
type AddSubscriptionPriceParamsPriceNewSubscriptionGroupedAllocationPriceCadence = shared.AddSubscriptionPriceParamsPriceNewSubscriptionGroupedAllocationPriceCadence

// This is an alias to an internal value.
const AddSubscriptionPriceParamsPriceNewSubscriptionGroupedAllocationPriceCadenceAnnual = shared.AddSubscriptionPriceParamsPriceNewSubscriptionGroupedAllocationPriceCadenceAnnual

// This is an alias to an internal value.
const AddSubscriptionPriceParamsPriceNewSubscriptionGroupedAllocationPriceCadenceSemiAnnual = shared.AddSubscriptionPriceParamsPriceNewSubscriptionGroupedAllocationPriceCadenceSemiAnnual

// This is an alias to an internal value.
const AddSubscriptionPriceParamsPriceNewSubscriptionGroupedAllocationPriceCadenceMonthly = shared.AddSubscriptionPriceParamsPriceNewSubscriptionGroupedAllocationPriceCadenceMonthly

// This is an alias to an internal value.
const AddSubscriptionPriceParamsPriceNewSubscriptionGroupedAllocationPriceCadenceQuarterly = shared.AddSubscriptionPriceParamsPriceNewSubscriptionGroupedAllocationPriceCadenceQuarterly

// This is an alias to an internal value.
const AddSubscriptionPriceParamsPriceNewSubscriptionGroupedAllocationPriceCadenceOneTime = shared.AddSubscriptionPriceParamsPriceNewSubscriptionGroupedAllocationPriceCadenceOneTime

// This is an alias to an internal value.
const AddSubscriptionPriceParamsPriceNewSubscriptionGroupedAllocationPriceCadenceCustom = shared.AddSubscriptionPriceParamsPriceNewSubscriptionGroupedAllocationPriceCadenceCustom

// This is an alias to an internal type.
type AddSubscriptionPriceParamsPriceNewSubscriptionGroupedAllocationPriceModelType = shared.AddSubscriptionPriceParamsPriceNewSubscriptionGroupedAllocationPriceModelType

// This is an alias to an internal value.
const AddSubscriptionPriceParamsPriceNewSubscriptionGroupedAllocationPriceModelTypeGroupedAllocation = shared.AddSubscriptionPriceParamsPriceNewSubscriptionGroupedAllocationPriceModelTypeGroupedAllocation

// For custom cadence: specifies the duration of the billing period in days or
// months.
//
// This is an alias to an internal type.
type AddSubscriptionPriceParamsPriceNewSubscriptionGroupedAllocationPriceBillingCycleConfiguration = shared.AddSubscriptionPriceParamsPriceNewSubscriptionGroupedAllocationPriceBillingCycleConfiguration

// The unit of billing period duration.
//
// This is an alias to an internal type.
type AddSubscriptionPriceParamsPriceNewSubscriptionGroupedAllocationPriceBillingCycleConfigurationDurationUnit = shared.AddSubscriptionPriceParamsPriceNewSubscriptionGroupedAllocationPriceBillingCycleConfigurationDurationUnit

// This is an alias to an internal value.
const AddSubscriptionPriceParamsPriceNewSubscriptionGroupedAllocationPriceBillingCycleConfigurationDurationUnitDay = shared.AddSubscriptionPriceParamsPriceNewSubscriptionGroupedAllocationPriceBillingCycleConfigurationDurationUnitDay

// This is an alias to an internal value.
const AddSubscriptionPriceParamsPriceNewSubscriptionGroupedAllocationPriceBillingCycleConfigurationDurationUnitMonth = shared.AddSubscriptionPriceParamsPriceNewSubscriptionGroupedAllocationPriceBillingCycleConfigurationDurationUnitMonth

// Within each billing cycle, specifies the cadence at which invoices are produced.
// If unspecified, a single invoice is produced per billing cycle.
//
// This is an alias to an internal type.
type AddSubscriptionPriceParamsPriceNewSubscriptionGroupedAllocationPriceInvoicingCycleConfiguration = shared.AddSubscriptionPriceParamsPriceNewSubscriptionGroupedAllocationPriceInvoicingCycleConfiguration

// The unit of billing period duration.
//
// This is an alias to an internal type.
type AddSubscriptionPriceParamsPriceNewSubscriptionGroupedAllocationPriceInvoicingCycleConfigurationDurationUnit = shared.AddSubscriptionPriceParamsPriceNewSubscriptionGroupedAllocationPriceInvoicingCycleConfigurationDurationUnit

// This is an alias to an internal value.
const AddSubscriptionPriceParamsPriceNewSubscriptionGroupedAllocationPriceInvoicingCycleConfigurationDurationUnitDay = shared.AddSubscriptionPriceParamsPriceNewSubscriptionGroupedAllocationPriceInvoicingCycleConfigurationDurationUnitDay

// This is an alias to an internal value.
const AddSubscriptionPriceParamsPriceNewSubscriptionGroupedAllocationPriceInvoicingCycleConfigurationDurationUnitMonth = shared.AddSubscriptionPriceParamsPriceNewSubscriptionGroupedAllocationPriceInvoicingCycleConfigurationDurationUnitMonth

// This is an alias to an internal type.
type AddSubscriptionPriceParamsPriceNewSubscriptionGroupedWithProratedMinimumPrice = shared.AddSubscriptionPriceParamsPriceNewSubscriptionGroupedWithProratedMinimumPrice

// The cadence to bill for this price on.
//
// This is an alias to an internal type.
type AddSubscriptionPriceParamsPriceNewSubscriptionGroupedWithProratedMinimumPriceCadence = shared.AddSubscriptionPriceParamsPriceNewSubscriptionGroupedWithProratedMinimumPriceCadence

// This is an alias to an internal value.
const AddSubscriptionPriceParamsPriceNewSubscriptionGroupedWithProratedMinimumPriceCadenceAnnual = shared.AddSubscriptionPriceParamsPriceNewSubscriptionGroupedWithProratedMinimumPriceCadenceAnnual

// This is an alias to an internal value.
const AddSubscriptionPriceParamsPriceNewSubscriptionGroupedWithProratedMinimumPriceCadenceSemiAnnual = shared.AddSubscriptionPriceParamsPriceNewSubscriptionGroupedWithProratedMinimumPriceCadenceSemiAnnual

// This is an alias to an internal value.
const AddSubscriptionPriceParamsPriceNewSubscriptionGroupedWithProratedMinimumPriceCadenceMonthly = shared.AddSubscriptionPriceParamsPriceNewSubscriptionGroupedWithProratedMinimumPriceCadenceMonthly

// This is an alias to an internal value.
const AddSubscriptionPriceParamsPriceNewSubscriptionGroupedWithProratedMinimumPriceCadenceQuarterly = shared.AddSubscriptionPriceParamsPriceNewSubscriptionGroupedWithProratedMinimumPriceCadenceQuarterly

// This is an alias to an internal value.
const AddSubscriptionPriceParamsPriceNewSubscriptionGroupedWithProratedMinimumPriceCadenceOneTime = shared.AddSubscriptionPriceParamsPriceNewSubscriptionGroupedWithProratedMinimumPriceCadenceOneTime

// This is an alias to an internal value.
const AddSubscriptionPriceParamsPriceNewSubscriptionGroupedWithProratedMinimumPriceCadenceCustom = shared.AddSubscriptionPriceParamsPriceNewSubscriptionGroupedWithProratedMinimumPriceCadenceCustom

// This is an alias to an internal type.
type AddSubscriptionPriceParamsPriceNewSubscriptionGroupedWithProratedMinimumPriceModelType = shared.AddSubscriptionPriceParamsPriceNewSubscriptionGroupedWithProratedMinimumPriceModelType

// This is an alias to an internal value.
const AddSubscriptionPriceParamsPriceNewSubscriptionGroupedWithProratedMinimumPriceModelTypeGroupedWithProratedMinimum = shared.AddSubscriptionPriceParamsPriceNewSubscriptionGroupedWithProratedMinimumPriceModelTypeGroupedWithProratedMinimum

// For custom cadence: specifies the duration of the billing period in days or
// months.
//
// This is an alias to an internal type.
type AddSubscriptionPriceParamsPriceNewSubscriptionGroupedWithProratedMinimumPriceBillingCycleConfiguration = shared.AddSubscriptionPriceParamsPriceNewSubscriptionGroupedWithProratedMinimumPriceBillingCycleConfiguration

// The unit of billing period duration.
//
// This is an alias to an internal type.
type AddSubscriptionPriceParamsPriceNewSubscriptionGroupedWithProratedMinimumPriceBillingCycleConfigurationDurationUnit = shared.AddSubscriptionPriceParamsPriceNewSubscriptionGroupedWithProratedMinimumPriceBillingCycleConfigurationDurationUnit

// This is an alias to an internal value.
const AddSubscriptionPriceParamsPriceNewSubscriptionGroupedWithProratedMinimumPriceBillingCycleConfigurationDurationUnitDay = shared.AddSubscriptionPriceParamsPriceNewSubscriptionGroupedWithProratedMinimumPriceBillingCycleConfigurationDurationUnitDay

// This is an alias to an internal value.
const AddSubscriptionPriceParamsPriceNewSubscriptionGroupedWithProratedMinimumPriceBillingCycleConfigurationDurationUnitMonth = shared.AddSubscriptionPriceParamsPriceNewSubscriptionGroupedWithProratedMinimumPriceBillingCycleConfigurationDurationUnitMonth

// Within each billing cycle, specifies the cadence at which invoices are produced.
// If unspecified, a single invoice is produced per billing cycle.
//
// This is an alias to an internal type.
type AddSubscriptionPriceParamsPriceNewSubscriptionGroupedWithProratedMinimumPriceInvoicingCycleConfiguration = shared.AddSubscriptionPriceParamsPriceNewSubscriptionGroupedWithProratedMinimumPriceInvoicingCycleConfiguration

// The unit of billing period duration.
//
// This is an alias to an internal type.
type AddSubscriptionPriceParamsPriceNewSubscriptionGroupedWithProratedMinimumPriceInvoicingCycleConfigurationDurationUnit = shared.AddSubscriptionPriceParamsPriceNewSubscriptionGroupedWithProratedMinimumPriceInvoicingCycleConfigurationDurationUnit

// This is an alias to an internal value.
const AddSubscriptionPriceParamsPriceNewSubscriptionGroupedWithProratedMinimumPriceInvoicingCycleConfigurationDurationUnitDay = shared.AddSubscriptionPriceParamsPriceNewSubscriptionGroupedWithProratedMinimumPriceInvoicingCycleConfigurationDurationUnitDay

// This is an alias to an internal value.
const AddSubscriptionPriceParamsPriceNewSubscriptionGroupedWithProratedMinimumPriceInvoicingCycleConfigurationDurationUnitMonth = shared.AddSubscriptionPriceParamsPriceNewSubscriptionGroupedWithProratedMinimumPriceInvoicingCycleConfigurationDurationUnitMonth

// This is an alias to an internal type.
type AddSubscriptionPriceParamsPriceNewSubscriptionBulkWithProrationPrice = shared.AddSubscriptionPriceParamsPriceNewSubscriptionBulkWithProrationPrice

// The cadence to bill for this price on.
//
// This is an alias to an internal type.
type AddSubscriptionPriceParamsPriceNewSubscriptionBulkWithProrationPriceCadence = shared.AddSubscriptionPriceParamsPriceNewSubscriptionBulkWithProrationPriceCadence

// This is an alias to an internal value.
const AddSubscriptionPriceParamsPriceNewSubscriptionBulkWithProrationPriceCadenceAnnual = shared.AddSubscriptionPriceParamsPriceNewSubscriptionBulkWithProrationPriceCadenceAnnual

// This is an alias to an internal value.
const AddSubscriptionPriceParamsPriceNewSubscriptionBulkWithProrationPriceCadenceSemiAnnual = shared.AddSubscriptionPriceParamsPriceNewSubscriptionBulkWithProrationPriceCadenceSemiAnnual

// This is an alias to an internal value.
const AddSubscriptionPriceParamsPriceNewSubscriptionBulkWithProrationPriceCadenceMonthly = shared.AddSubscriptionPriceParamsPriceNewSubscriptionBulkWithProrationPriceCadenceMonthly

// This is an alias to an internal value.
const AddSubscriptionPriceParamsPriceNewSubscriptionBulkWithProrationPriceCadenceQuarterly = shared.AddSubscriptionPriceParamsPriceNewSubscriptionBulkWithProrationPriceCadenceQuarterly

// This is an alias to an internal value.
const AddSubscriptionPriceParamsPriceNewSubscriptionBulkWithProrationPriceCadenceOneTime = shared.AddSubscriptionPriceParamsPriceNewSubscriptionBulkWithProrationPriceCadenceOneTime

// This is an alias to an internal value.
const AddSubscriptionPriceParamsPriceNewSubscriptionBulkWithProrationPriceCadenceCustom = shared.AddSubscriptionPriceParamsPriceNewSubscriptionBulkWithProrationPriceCadenceCustom

// This is an alias to an internal type.
type AddSubscriptionPriceParamsPriceNewSubscriptionBulkWithProrationPriceModelType = shared.AddSubscriptionPriceParamsPriceNewSubscriptionBulkWithProrationPriceModelType

// This is an alias to an internal value.
const AddSubscriptionPriceParamsPriceNewSubscriptionBulkWithProrationPriceModelTypeBulkWithProration = shared.AddSubscriptionPriceParamsPriceNewSubscriptionBulkWithProrationPriceModelTypeBulkWithProration

// For custom cadence: specifies the duration of the billing period in days or
// months.
//
// This is an alias to an internal type.
type AddSubscriptionPriceParamsPriceNewSubscriptionBulkWithProrationPriceBillingCycleConfiguration = shared.AddSubscriptionPriceParamsPriceNewSubscriptionBulkWithProrationPriceBillingCycleConfiguration

// The unit of billing period duration.
//
// This is an alias to an internal type.
type AddSubscriptionPriceParamsPriceNewSubscriptionBulkWithProrationPriceBillingCycleConfigurationDurationUnit = shared.AddSubscriptionPriceParamsPriceNewSubscriptionBulkWithProrationPriceBillingCycleConfigurationDurationUnit

// This is an alias to an internal value.
const AddSubscriptionPriceParamsPriceNewSubscriptionBulkWithProrationPriceBillingCycleConfigurationDurationUnitDay = shared.AddSubscriptionPriceParamsPriceNewSubscriptionBulkWithProrationPriceBillingCycleConfigurationDurationUnitDay

// This is an alias to an internal value.
const AddSubscriptionPriceParamsPriceNewSubscriptionBulkWithProrationPriceBillingCycleConfigurationDurationUnitMonth = shared.AddSubscriptionPriceParamsPriceNewSubscriptionBulkWithProrationPriceBillingCycleConfigurationDurationUnitMonth

// Within each billing cycle, specifies the cadence at which invoices are produced.
// If unspecified, a single invoice is produced per billing cycle.
//
// This is an alias to an internal type.
type AddSubscriptionPriceParamsPriceNewSubscriptionBulkWithProrationPriceInvoicingCycleConfiguration = shared.AddSubscriptionPriceParamsPriceNewSubscriptionBulkWithProrationPriceInvoicingCycleConfiguration

// The unit of billing period duration.
//
// This is an alias to an internal type.
type AddSubscriptionPriceParamsPriceNewSubscriptionBulkWithProrationPriceInvoicingCycleConfigurationDurationUnit = shared.AddSubscriptionPriceParamsPriceNewSubscriptionBulkWithProrationPriceInvoicingCycleConfigurationDurationUnit

// This is an alias to an internal value.
const AddSubscriptionPriceParamsPriceNewSubscriptionBulkWithProrationPriceInvoicingCycleConfigurationDurationUnitDay = shared.AddSubscriptionPriceParamsPriceNewSubscriptionBulkWithProrationPriceInvoicingCycleConfigurationDurationUnitDay

// This is an alias to an internal value.
const AddSubscriptionPriceParamsPriceNewSubscriptionBulkWithProrationPriceInvoicingCycleConfigurationDurationUnitMonth = shared.AddSubscriptionPriceParamsPriceNewSubscriptionBulkWithProrationPriceInvoicingCycleConfigurationDurationUnitMonth

// This is an alias to an internal type.
type AddSubscriptionPriceParamsPriceNewSubscriptionScalableMatrixWithUnitPricingPrice = shared.AddSubscriptionPriceParamsPriceNewSubscriptionScalableMatrixWithUnitPricingPrice

// The cadence to bill for this price on.
//
// This is an alias to an internal type.
type AddSubscriptionPriceParamsPriceNewSubscriptionScalableMatrixWithUnitPricingPriceCadence = shared.AddSubscriptionPriceParamsPriceNewSubscriptionScalableMatrixWithUnitPricingPriceCadence

// This is an alias to an internal value.
const AddSubscriptionPriceParamsPriceNewSubscriptionScalableMatrixWithUnitPricingPriceCadenceAnnual = shared.AddSubscriptionPriceParamsPriceNewSubscriptionScalableMatrixWithUnitPricingPriceCadenceAnnual

// This is an alias to an internal value.
const AddSubscriptionPriceParamsPriceNewSubscriptionScalableMatrixWithUnitPricingPriceCadenceSemiAnnual = shared.AddSubscriptionPriceParamsPriceNewSubscriptionScalableMatrixWithUnitPricingPriceCadenceSemiAnnual

// This is an alias to an internal value.
const AddSubscriptionPriceParamsPriceNewSubscriptionScalableMatrixWithUnitPricingPriceCadenceMonthly = shared.AddSubscriptionPriceParamsPriceNewSubscriptionScalableMatrixWithUnitPricingPriceCadenceMonthly

// This is an alias to an internal value.
const AddSubscriptionPriceParamsPriceNewSubscriptionScalableMatrixWithUnitPricingPriceCadenceQuarterly = shared.AddSubscriptionPriceParamsPriceNewSubscriptionScalableMatrixWithUnitPricingPriceCadenceQuarterly

// This is an alias to an internal value.
const AddSubscriptionPriceParamsPriceNewSubscriptionScalableMatrixWithUnitPricingPriceCadenceOneTime = shared.AddSubscriptionPriceParamsPriceNewSubscriptionScalableMatrixWithUnitPricingPriceCadenceOneTime

// This is an alias to an internal value.
const AddSubscriptionPriceParamsPriceNewSubscriptionScalableMatrixWithUnitPricingPriceCadenceCustom = shared.AddSubscriptionPriceParamsPriceNewSubscriptionScalableMatrixWithUnitPricingPriceCadenceCustom

// This is an alias to an internal type.
type AddSubscriptionPriceParamsPriceNewSubscriptionScalableMatrixWithUnitPricingPriceModelType = shared.AddSubscriptionPriceParamsPriceNewSubscriptionScalableMatrixWithUnitPricingPriceModelType

// This is an alias to an internal value.
const AddSubscriptionPriceParamsPriceNewSubscriptionScalableMatrixWithUnitPricingPriceModelTypeScalableMatrixWithUnitPricing = shared.AddSubscriptionPriceParamsPriceNewSubscriptionScalableMatrixWithUnitPricingPriceModelTypeScalableMatrixWithUnitPricing

// For custom cadence: specifies the duration of the billing period in days or
// months.
//
// This is an alias to an internal type.
type AddSubscriptionPriceParamsPriceNewSubscriptionScalableMatrixWithUnitPricingPriceBillingCycleConfiguration = shared.AddSubscriptionPriceParamsPriceNewSubscriptionScalableMatrixWithUnitPricingPriceBillingCycleConfiguration

// The unit of billing period duration.
//
// This is an alias to an internal type.
type AddSubscriptionPriceParamsPriceNewSubscriptionScalableMatrixWithUnitPricingPriceBillingCycleConfigurationDurationUnit = shared.AddSubscriptionPriceParamsPriceNewSubscriptionScalableMatrixWithUnitPricingPriceBillingCycleConfigurationDurationUnit

// This is an alias to an internal value.
const AddSubscriptionPriceParamsPriceNewSubscriptionScalableMatrixWithUnitPricingPriceBillingCycleConfigurationDurationUnitDay = shared.AddSubscriptionPriceParamsPriceNewSubscriptionScalableMatrixWithUnitPricingPriceBillingCycleConfigurationDurationUnitDay

// This is an alias to an internal value.
const AddSubscriptionPriceParamsPriceNewSubscriptionScalableMatrixWithUnitPricingPriceBillingCycleConfigurationDurationUnitMonth = shared.AddSubscriptionPriceParamsPriceNewSubscriptionScalableMatrixWithUnitPricingPriceBillingCycleConfigurationDurationUnitMonth

// Within each billing cycle, specifies the cadence at which invoices are produced.
// If unspecified, a single invoice is produced per billing cycle.
//
// This is an alias to an internal type.
type AddSubscriptionPriceParamsPriceNewSubscriptionScalableMatrixWithUnitPricingPriceInvoicingCycleConfiguration = shared.AddSubscriptionPriceParamsPriceNewSubscriptionScalableMatrixWithUnitPricingPriceInvoicingCycleConfiguration

// The unit of billing period duration.
//
// This is an alias to an internal type.
type AddSubscriptionPriceParamsPriceNewSubscriptionScalableMatrixWithUnitPricingPriceInvoicingCycleConfigurationDurationUnit = shared.AddSubscriptionPriceParamsPriceNewSubscriptionScalableMatrixWithUnitPricingPriceInvoicingCycleConfigurationDurationUnit

// This is an alias to an internal value.
const AddSubscriptionPriceParamsPriceNewSubscriptionScalableMatrixWithUnitPricingPriceInvoicingCycleConfigurationDurationUnitDay = shared.AddSubscriptionPriceParamsPriceNewSubscriptionScalableMatrixWithUnitPricingPriceInvoicingCycleConfigurationDurationUnitDay

// This is an alias to an internal value.
const AddSubscriptionPriceParamsPriceNewSubscriptionScalableMatrixWithUnitPricingPriceInvoicingCycleConfigurationDurationUnitMonth = shared.AddSubscriptionPriceParamsPriceNewSubscriptionScalableMatrixWithUnitPricingPriceInvoicingCycleConfigurationDurationUnitMonth

// This is an alias to an internal type.
type AddSubscriptionPriceParamsPriceNewSubscriptionScalableMatrixWithTieredPricingPrice = shared.AddSubscriptionPriceParamsPriceNewSubscriptionScalableMatrixWithTieredPricingPrice

// The cadence to bill for this price on.
//
// This is an alias to an internal type.
type AddSubscriptionPriceParamsPriceNewSubscriptionScalableMatrixWithTieredPricingPriceCadence = shared.AddSubscriptionPriceParamsPriceNewSubscriptionScalableMatrixWithTieredPricingPriceCadence

// This is an alias to an internal value.
const AddSubscriptionPriceParamsPriceNewSubscriptionScalableMatrixWithTieredPricingPriceCadenceAnnual = shared.AddSubscriptionPriceParamsPriceNewSubscriptionScalableMatrixWithTieredPricingPriceCadenceAnnual

// This is an alias to an internal value.
const AddSubscriptionPriceParamsPriceNewSubscriptionScalableMatrixWithTieredPricingPriceCadenceSemiAnnual = shared.AddSubscriptionPriceParamsPriceNewSubscriptionScalableMatrixWithTieredPricingPriceCadenceSemiAnnual

// This is an alias to an internal value.
const AddSubscriptionPriceParamsPriceNewSubscriptionScalableMatrixWithTieredPricingPriceCadenceMonthly = shared.AddSubscriptionPriceParamsPriceNewSubscriptionScalableMatrixWithTieredPricingPriceCadenceMonthly

// This is an alias to an internal value.
const AddSubscriptionPriceParamsPriceNewSubscriptionScalableMatrixWithTieredPricingPriceCadenceQuarterly = shared.AddSubscriptionPriceParamsPriceNewSubscriptionScalableMatrixWithTieredPricingPriceCadenceQuarterly

// This is an alias to an internal value.
const AddSubscriptionPriceParamsPriceNewSubscriptionScalableMatrixWithTieredPricingPriceCadenceOneTime = shared.AddSubscriptionPriceParamsPriceNewSubscriptionScalableMatrixWithTieredPricingPriceCadenceOneTime

// This is an alias to an internal value.
const AddSubscriptionPriceParamsPriceNewSubscriptionScalableMatrixWithTieredPricingPriceCadenceCustom = shared.AddSubscriptionPriceParamsPriceNewSubscriptionScalableMatrixWithTieredPricingPriceCadenceCustom

// This is an alias to an internal type.
type AddSubscriptionPriceParamsPriceNewSubscriptionScalableMatrixWithTieredPricingPriceModelType = shared.AddSubscriptionPriceParamsPriceNewSubscriptionScalableMatrixWithTieredPricingPriceModelType

// This is an alias to an internal value.
const AddSubscriptionPriceParamsPriceNewSubscriptionScalableMatrixWithTieredPricingPriceModelTypeScalableMatrixWithTieredPricing = shared.AddSubscriptionPriceParamsPriceNewSubscriptionScalableMatrixWithTieredPricingPriceModelTypeScalableMatrixWithTieredPricing

// For custom cadence: specifies the duration of the billing period in days or
// months.
//
// This is an alias to an internal type.
type AddSubscriptionPriceParamsPriceNewSubscriptionScalableMatrixWithTieredPricingPriceBillingCycleConfiguration = shared.AddSubscriptionPriceParamsPriceNewSubscriptionScalableMatrixWithTieredPricingPriceBillingCycleConfiguration

// The unit of billing period duration.
//
// This is an alias to an internal type.
type AddSubscriptionPriceParamsPriceNewSubscriptionScalableMatrixWithTieredPricingPriceBillingCycleConfigurationDurationUnit = shared.AddSubscriptionPriceParamsPriceNewSubscriptionScalableMatrixWithTieredPricingPriceBillingCycleConfigurationDurationUnit

// This is an alias to an internal value.
const AddSubscriptionPriceParamsPriceNewSubscriptionScalableMatrixWithTieredPricingPriceBillingCycleConfigurationDurationUnitDay = shared.AddSubscriptionPriceParamsPriceNewSubscriptionScalableMatrixWithTieredPricingPriceBillingCycleConfigurationDurationUnitDay

// This is an alias to an internal value.
const AddSubscriptionPriceParamsPriceNewSubscriptionScalableMatrixWithTieredPricingPriceBillingCycleConfigurationDurationUnitMonth = shared.AddSubscriptionPriceParamsPriceNewSubscriptionScalableMatrixWithTieredPricingPriceBillingCycleConfigurationDurationUnitMonth

// Within each billing cycle, specifies the cadence at which invoices are produced.
// If unspecified, a single invoice is produced per billing cycle.
//
// This is an alias to an internal type.
type AddSubscriptionPriceParamsPriceNewSubscriptionScalableMatrixWithTieredPricingPriceInvoicingCycleConfiguration = shared.AddSubscriptionPriceParamsPriceNewSubscriptionScalableMatrixWithTieredPricingPriceInvoicingCycleConfiguration

// The unit of billing period duration.
//
// This is an alias to an internal type.
type AddSubscriptionPriceParamsPriceNewSubscriptionScalableMatrixWithTieredPricingPriceInvoicingCycleConfigurationDurationUnit = shared.AddSubscriptionPriceParamsPriceNewSubscriptionScalableMatrixWithTieredPricingPriceInvoicingCycleConfigurationDurationUnit

// This is an alias to an internal value.
const AddSubscriptionPriceParamsPriceNewSubscriptionScalableMatrixWithTieredPricingPriceInvoicingCycleConfigurationDurationUnitDay = shared.AddSubscriptionPriceParamsPriceNewSubscriptionScalableMatrixWithTieredPricingPriceInvoicingCycleConfigurationDurationUnitDay

// This is an alias to an internal value.
const AddSubscriptionPriceParamsPriceNewSubscriptionScalableMatrixWithTieredPricingPriceInvoicingCycleConfigurationDurationUnitMonth = shared.AddSubscriptionPriceParamsPriceNewSubscriptionScalableMatrixWithTieredPricingPriceInvoicingCycleConfigurationDurationUnitMonth

// This is an alias to an internal type.
type AddSubscriptionPriceParamsPriceNewSubscriptionCumulativeGroupedBulkPrice = shared.AddSubscriptionPriceParamsPriceNewSubscriptionCumulativeGroupedBulkPrice

// The cadence to bill for this price on.
//
// This is an alias to an internal type.
type AddSubscriptionPriceParamsPriceNewSubscriptionCumulativeGroupedBulkPriceCadence = shared.AddSubscriptionPriceParamsPriceNewSubscriptionCumulativeGroupedBulkPriceCadence

// This is an alias to an internal value.
const AddSubscriptionPriceParamsPriceNewSubscriptionCumulativeGroupedBulkPriceCadenceAnnual = shared.AddSubscriptionPriceParamsPriceNewSubscriptionCumulativeGroupedBulkPriceCadenceAnnual

// This is an alias to an internal value.
const AddSubscriptionPriceParamsPriceNewSubscriptionCumulativeGroupedBulkPriceCadenceSemiAnnual = shared.AddSubscriptionPriceParamsPriceNewSubscriptionCumulativeGroupedBulkPriceCadenceSemiAnnual

// This is an alias to an internal value.
const AddSubscriptionPriceParamsPriceNewSubscriptionCumulativeGroupedBulkPriceCadenceMonthly = shared.AddSubscriptionPriceParamsPriceNewSubscriptionCumulativeGroupedBulkPriceCadenceMonthly

// This is an alias to an internal value.
const AddSubscriptionPriceParamsPriceNewSubscriptionCumulativeGroupedBulkPriceCadenceQuarterly = shared.AddSubscriptionPriceParamsPriceNewSubscriptionCumulativeGroupedBulkPriceCadenceQuarterly

// This is an alias to an internal value.
const AddSubscriptionPriceParamsPriceNewSubscriptionCumulativeGroupedBulkPriceCadenceOneTime = shared.AddSubscriptionPriceParamsPriceNewSubscriptionCumulativeGroupedBulkPriceCadenceOneTime

// This is an alias to an internal value.
const AddSubscriptionPriceParamsPriceNewSubscriptionCumulativeGroupedBulkPriceCadenceCustom = shared.AddSubscriptionPriceParamsPriceNewSubscriptionCumulativeGroupedBulkPriceCadenceCustom

// This is an alias to an internal type.
type AddSubscriptionPriceParamsPriceNewSubscriptionCumulativeGroupedBulkPriceModelType = shared.AddSubscriptionPriceParamsPriceNewSubscriptionCumulativeGroupedBulkPriceModelType

// This is an alias to an internal value.
const AddSubscriptionPriceParamsPriceNewSubscriptionCumulativeGroupedBulkPriceModelTypeCumulativeGroupedBulk = shared.AddSubscriptionPriceParamsPriceNewSubscriptionCumulativeGroupedBulkPriceModelTypeCumulativeGroupedBulk

// For custom cadence: specifies the duration of the billing period in days or
// months.
//
// This is an alias to an internal type.
type AddSubscriptionPriceParamsPriceNewSubscriptionCumulativeGroupedBulkPriceBillingCycleConfiguration = shared.AddSubscriptionPriceParamsPriceNewSubscriptionCumulativeGroupedBulkPriceBillingCycleConfiguration

// The unit of billing period duration.
//
// This is an alias to an internal type.
type AddSubscriptionPriceParamsPriceNewSubscriptionCumulativeGroupedBulkPriceBillingCycleConfigurationDurationUnit = shared.AddSubscriptionPriceParamsPriceNewSubscriptionCumulativeGroupedBulkPriceBillingCycleConfigurationDurationUnit

// This is an alias to an internal value.
const AddSubscriptionPriceParamsPriceNewSubscriptionCumulativeGroupedBulkPriceBillingCycleConfigurationDurationUnitDay = shared.AddSubscriptionPriceParamsPriceNewSubscriptionCumulativeGroupedBulkPriceBillingCycleConfigurationDurationUnitDay

// This is an alias to an internal value.
const AddSubscriptionPriceParamsPriceNewSubscriptionCumulativeGroupedBulkPriceBillingCycleConfigurationDurationUnitMonth = shared.AddSubscriptionPriceParamsPriceNewSubscriptionCumulativeGroupedBulkPriceBillingCycleConfigurationDurationUnitMonth

// Within each billing cycle, specifies the cadence at which invoices are produced.
// If unspecified, a single invoice is produced per billing cycle.
//
// This is an alias to an internal type.
type AddSubscriptionPriceParamsPriceNewSubscriptionCumulativeGroupedBulkPriceInvoicingCycleConfiguration = shared.AddSubscriptionPriceParamsPriceNewSubscriptionCumulativeGroupedBulkPriceInvoicingCycleConfiguration

// The unit of billing period duration.
//
// This is an alias to an internal type.
type AddSubscriptionPriceParamsPriceNewSubscriptionCumulativeGroupedBulkPriceInvoicingCycleConfigurationDurationUnit = shared.AddSubscriptionPriceParamsPriceNewSubscriptionCumulativeGroupedBulkPriceInvoicingCycleConfigurationDurationUnit

// This is an alias to an internal value.
const AddSubscriptionPriceParamsPriceNewSubscriptionCumulativeGroupedBulkPriceInvoicingCycleConfigurationDurationUnitDay = shared.AddSubscriptionPriceParamsPriceNewSubscriptionCumulativeGroupedBulkPriceInvoicingCycleConfigurationDurationUnitDay

// This is an alias to an internal value.
const AddSubscriptionPriceParamsPriceNewSubscriptionCumulativeGroupedBulkPriceInvoicingCycleConfigurationDurationUnitMonth = shared.AddSubscriptionPriceParamsPriceNewSubscriptionCumulativeGroupedBulkPriceInvoicingCycleConfigurationDurationUnitMonth

// This is an alias to an internal type.
type AddSubscriptionPriceParamsPriceNewSubscriptionMaxGroupTieredPackagePrice = shared.AddSubscriptionPriceParamsPriceNewSubscriptionMaxGroupTieredPackagePrice

// The cadence to bill for this price on.
//
// This is an alias to an internal type.
type AddSubscriptionPriceParamsPriceNewSubscriptionMaxGroupTieredPackagePriceCadence = shared.AddSubscriptionPriceParamsPriceNewSubscriptionMaxGroupTieredPackagePriceCadence

// This is an alias to an internal value.
const AddSubscriptionPriceParamsPriceNewSubscriptionMaxGroupTieredPackagePriceCadenceAnnual = shared.AddSubscriptionPriceParamsPriceNewSubscriptionMaxGroupTieredPackagePriceCadenceAnnual

// This is an alias to an internal value.
const AddSubscriptionPriceParamsPriceNewSubscriptionMaxGroupTieredPackagePriceCadenceSemiAnnual = shared.AddSubscriptionPriceParamsPriceNewSubscriptionMaxGroupTieredPackagePriceCadenceSemiAnnual

// This is an alias to an internal value.
const AddSubscriptionPriceParamsPriceNewSubscriptionMaxGroupTieredPackagePriceCadenceMonthly = shared.AddSubscriptionPriceParamsPriceNewSubscriptionMaxGroupTieredPackagePriceCadenceMonthly

// This is an alias to an internal value.
const AddSubscriptionPriceParamsPriceNewSubscriptionMaxGroupTieredPackagePriceCadenceQuarterly = shared.AddSubscriptionPriceParamsPriceNewSubscriptionMaxGroupTieredPackagePriceCadenceQuarterly

// This is an alias to an internal value.
const AddSubscriptionPriceParamsPriceNewSubscriptionMaxGroupTieredPackagePriceCadenceOneTime = shared.AddSubscriptionPriceParamsPriceNewSubscriptionMaxGroupTieredPackagePriceCadenceOneTime

// This is an alias to an internal value.
const AddSubscriptionPriceParamsPriceNewSubscriptionMaxGroupTieredPackagePriceCadenceCustom = shared.AddSubscriptionPriceParamsPriceNewSubscriptionMaxGroupTieredPackagePriceCadenceCustom

// This is an alias to an internal type.
type AddSubscriptionPriceParamsPriceNewSubscriptionMaxGroupTieredPackagePriceModelType = shared.AddSubscriptionPriceParamsPriceNewSubscriptionMaxGroupTieredPackagePriceModelType

// This is an alias to an internal value.
const AddSubscriptionPriceParamsPriceNewSubscriptionMaxGroupTieredPackagePriceModelTypeMaxGroupTieredPackage = shared.AddSubscriptionPriceParamsPriceNewSubscriptionMaxGroupTieredPackagePriceModelTypeMaxGroupTieredPackage

// For custom cadence: specifies the duration of the billing period in days or
// months.
//
// This is an alias to an internal type.
type AddSubscriptionPriceParamsPriceNewSubscriptionMaxGroupTieredPackagePriceBillingCycleConfiguration = shared.AddSubscriptionPriceParamsPriceNewSubscriptionMaxGroupTieredPackagePriceBillingCycleConfiguration

// The unit of billing period duration.
//
// This is an alias to an internal type.
type AddSubscriptionPriceParamsPriceNewSubscriptionMaxGroupTieredPackagePriceBillingCycleConfigurationDurationUnit = shared.AddSubscriptionPriceParamsPriceNewSubscriptionMaxGroupTieredPackagePriceBillingCycleConfigurationDurationUnit

// This is an alias to an internal value.
const AddSubscriptionPriceParamsPriceNewSubscriptionMaxGroupTieredPackagePriceBillingCycleConfigurationDurationUnitDay = shared.AddSubscriptionPriceParamsPriceNewSubscriptionMaxGroupTieredPackagePriceBillingCycleConfigurationDurationUnitDay

// This is an alias to an internal value.
const AddSubscriptionPriceParamsPriceNewSubscriptionMaxGroupTieredPackagePriceBillingCycleConfigurationDurationUnitMonth = shared.AddSubscriptionPriceParamsPriceNewSubscriptionMaxGroupTieredPackagePriceBillingCycleConfigurationDurationUnitMonth

// Within each billing cycle, specifies the cadence at which invoices are produced.
// If unspecified, a single invoice is produced per billing cycle.
//
// This is an alias to an internal type.
type AddSubscriptionPriceParamsPriceNewSubscriptionMaxGroupTieredPackagePriceInvoicingCycleConfiguration = shared.AddSubscriptionPriceParamsPriceNewSubscriptionMaxGroupTieredPackagePriceInvoicingCycleConfiguration

// The unit of billing period duration.
//
// This is an alias to an internal type.
type AddSubscriptionPriceParamsPriceNewSubscriptionMaxGroupTieredPackagePriceInvoicingCycleConfigurationDurationUnit = shared.AddSubscriptionPriceParamsPriceNewSubscriptionMaxGroupTieredPackagePriceInvoicingCycleConfigurationDurationUnit

// This is an alias to an internal value.
const AddSubscriptionPriceParamsPriceNewSubscriptionMaxGroupTieredPackagePriceInvoicingCycleConfigurationDurationUnitDay = shared.AddSubscriptionPriceParamsPriceNewSubscriptionMaxGroupTieredPackagePriceInvoicingCycleConfigurationDurationUnitDay

// This is an alias to an internal value.
const AddSubscriptionPriceParamsPriceNewSubscriptionMaxGroupTieredPackagePriceInvoicingCycleConfigurationDurationUnitMonth = shared.AddSubscriptionPriceParamsPriceNewSubscriptionMaxGroupTieredPackagePriceInvoicingCycleConfigurationDurationUnitMonth

// This is an alias to an internal type.
type AddSubscriptionPriceParamsPriceNewSubscriptionGroupedWithMeteredMinimumPrice = shared.AddSubscriptionPriceParamsPriceNewSubscriptionGroupedWithMeteredMinimumPrice

// The cadence to bill for this price on.
//
// This is an alias to an internal type.
type AddSubscriptionPriceParamsPriceNewSubscriptionGroupedWithMeteredMinimumPriceCadence = shared.AddSubscriptionPriceParamsPriceNewSubscriptionGroupedWithMeteredMinimumPriceCadence

// This is an alias to an internal value.
const AddSubscriptionPriceParamsPriceNewSubscriptionGroupedWithMeteredMinimumPriceCadenceAnnual = shared.AddSubscriptionPriceParamsPriceNewSubscriptionGroupedWithMeteredMinimumPriceCadenceAnnual

// This is an alias to an internal value.
const AddSubscriptionPriceParamsPriceNewSubscriptionGroupedWithMeteredMinimumPriceCadenceSemiAnnual = shared.AddSubscriptionPriceParamsPriceNewSubscriptionGroupedWithMeteredMinimumPriceCadenceSemiAnnual

// This is an alias to an internal value.
const AddSubscriptionPriceParamsPriceNewSubscriptionGroupedWithMeteredMinimumPriceCadenceMonthly = shared.AddSubscriptionPriceParamsPriceNewSubscriptionGroupedWithMeteredMinimumPriceCadenceMonthly

// This is an alias to an internal value.
const AddSubscriptionPriceParamsPriceNewSubscriptionGroupedWithMeteredMinimumPriceCadenceQuarterly = shared.AddSubscriptionPriceParamsPriceNewSubscriptionGroupedWithMeteredMinimumPriceCadenceQuarterly

// This is an alias to an internal value.
const AddSubscriptionPriceParamsPriceNewSubscriptionGroupedWithMeteredMinimumPriceCadenceOneTime = shared.AddSubscriptionPriceParamsPriceNewSubscriptionGroupedWithMeteredMinimumPriceCadenceOneTime

// This is an alias to an internal value.
const AddSubscriptionPriceParamsPriceNewSubscriptionGroupedWithMeteredMinimumPriceCadenceCustom = shared.AddSubscriptionPriceParamsPriceNewSubscriptionGroupedWithMeteredMinimumPriceCadenceCustom

// This is an alias to an internal type.
type AddSubscriptionPriceParamsPriceNewSubscriptionGroupedWithMeteredMinimumPriceModelType = shared.AddSubscriptionPriceParamsPriceNewSubscriptionGroupedWithMeteredMinimumPriceModelType

// This is an alias to an internal value.
const AddSubscriptionPriceParamsPriceNewSubscriptionGroupedWithMeteredMinimumPriceModelTypeGroupedWithMeteredMinimum = shared.AddSubscriptionPriceParamsPriceNewSubscriptionGroupedWithMeteredMinimumPriceModelTypeGroupedWithMeteredMinimum

// For custom cadence: specifies the duration of the billing period in days or
// months.
//
// This is an alias to an internal type.
type AddSubscriptionPriceParamsPriceNewSubscriptionGroupedWithMeteredMinimumPriceBillingCycleConfiguration = shared.AddSubscriptionPriceParamsPriceNewSubscriptionGroupedWithMeteredMinimumPriceBillingCycleConfiguration

// The unit of billing period duration.
//
// This is an alias to an internal type.
type AddSubscriptionPriceParamsPriceNewSubscriptionGroupedWithMeteredMinimumPriceBillingCycleConfigurationDurationUnit = shared.AddSubscriptionPriceParamsPriceNewSubscriptionGroupedWithMeteredMinimumPriceBillingCycleConfigurationDurationUnit

// This is an alias to an internal value.
const AddSubscriptionPriceParamsPriceNewSubscriptionGroupedWithMeteredMinimumPriceBillingCycleConfigurationDurationUnitDay = shared.AddSubscriptionPriceParamsPriceNewSubscriptionGroupedWithMeteredMinimumPriceBillingCycleConfigurationDurationUnitDay

// This is an alias to an internal value.
const AddSubscriptionPriceParamsPriceNewSubscriptionGroupedWithMeteredMinimumPriceBillingCycleConfigurationDurationUnitMonth = shared.AddSubscriptionPriceParamsPriceNewSubscriptionGroupedWithMeteredMinimumPriceBillingCycleConfigurationDurationUnitMonth

// Within each billing cycle, specifies the cadence at which invoices are produced.
// If unspecified, a single invoice is produced per billing cycle.
//
// This is an alias to an internal type.
type AddSubscriptionPriceParamsPriceNewSubscriptionGroupedWithMeteredMinimumPriceInvoicingCycleConfiguration = shared.AddSubscriptionPriceParamsPriceNewSubscriptionGroupedWithMeteredMinimumPriceInvoicingCycleConfiguration

// The unit of billing period duration.
//
// This is an alias to an internal type.
type AddSubscriptionPriceParamsPriceNewSubscriptionGroupedWithMeteredMinimumPriceInvoicingCycleConfigurationDurationUnit = shared.AddSubscriptionPriceParamsPriceNewSubscriptionGroupedWithMeteredMinimumPriceInvoicingCycleConfigurationDurationUnit

// This is an alias to an internal value.
const AddSubscriptionPriceParamsPriceNewSubscriptionGroupedWithMeteredMinimumPriceInvoicingCycleConfigurationDurationUnitDay = shared.AddSubscriptionPriceParamsPriceNewSubscriptionGroupedWithMeteredMinimumPriceInvoicingCycleConfigurationDurationUnitDay

// This is an alias to an internal value.
const AddSubscriptionPriceParamsPriceNewSubscriptionGroupedWithMeteredMinimumPriceInvoicingCycleConfigurationDurationUnitMonth = shared.AddSubscriptionPriceParamsPriceNewSubscriptionGroupedWithMeteredMinimumPriceInvoicingCycleConfigurationDurationUnitMonth

// This is an alias to an internal type.
type AddSubscriptionPriceParamsPriceNewSubscriptionMatrixWithDisplayNamePrice = shared.AddSubscriptionPriceParamsPriceNewSubscriptionMatrixWithDisplayNamePrice

// The cadence to bill for this price on.
//
// This is an alias to an internal type.
type AddSubscriptionPriceParamsPriceNewSubscriptionMatrixWithDisplayNamePriceCadence = shared.AddSubscriptionPriceParamsPriceNewSubscriptionMatrixWithDisplayNamePriceCadence

// This is an alias to an internal value.
const AddSubscriptionPriceParamsPriceNewSubscriptionMatrixWithDisplayNamePriceCadenceAnnual = shared.AddSubscriptionPriceParamsPriceNewSubscriptionMatrixWithDisplayNamePriceCadenceAnnual

// This is an alias to an internal value.
const AddSubscriptionPriceParamsPriceNewSubscriptionMatrixWithDisplayNamePriceCadenceSemiAnnual = shared.AddSubscriptionPriceParamsPriceNewSubscriptionMatrixWithDisplayNamePriceCadenceSemiAnnual

// This is an alias to an internal value.
const AddSubscriptionPriceParamsPriceNewSubscriptionMatrixWithDisplayNamePriceCadenceMonthly = shared.AddSubscriptionPriceParamsPriceNewSubscriptionMatrixWithDisplayNamePriceCadenceMonthly

// This is an alias to an internal value.
const AddSubscriptionPriceParamsPriceNewSubscriptionMatrixWithDisplayNamePriceCadenceQuarterly = shared.AddSubscriptionPriceParamsPriceNewSubscriptionMatrixWithDisplayNamePriceCadenceQuarterly

// This is an alias to an internal value.
const AddSubscriptionPriceParamsPriceNewSubscriptionMatrixWithDisplayNamePriceCadenceOneTime = shared.AddSubscriptionPriceParamsPriceNewSubscriptionMatrixWithDisplayNamePriceCadenceOneTime

// This is an alias to an internal value.
const AddSubscriptionPriceParamsPriceNewSubscriptionMatrixWithDisplayNamePriceCadenceCustom = shared.AddSubscriptionPriceParamsPriceNewSubscriptionMatrixWithDisplayNamePriceCadenceCustom

// This is an alias to an internal type.
type AddSubscriptionPriceParamsPriceNewSubscriptionMatrixWithDisplayNamePriceModelType = shared.AddSubscriptionPriceParamsPriceNewSubscriptionMatrixWithDisplayNamePriceModelType

// This is an alias to an internal value.
const AddSubscriptionPriceParamsPriceNewSubscriptionMatrixWithDisplayNamePriceModelTypeMatrixWithDisplayName = shared.AddSubscriptionPriceParamsPriceNewSubscriptionMatrixWithDisplayNamePriceModelTypeMatrixWithDisplayName

// For custom cadence: specifies the duration of the billing period in days or
// months.
//
// This is an alias to an internal type.
type AddSubscriptionPriceParamsPriceNewSubscriptionMatrixWithDisplayNamePriceBillingCycleConfiguration = shared.AddSubscriptionPriceParamsPriceNewSubscriptionMatrixWithDisplayNamePriceBillingCycleConfiguration

// The unit of billing period duration.
//
// This is an alias to an internal type.
type AddSubscriptionPriceParamsPriceNewSubscriptionMatrixWithDisplayNamePriceBillingCycleConfigurationDurationUnit = shared.AddSubscriptionPriceParamsPriceNewSubscriptionMatrixWithDisplayNamePriceBillingCycleConfigurationDurationUnit

// This is an alias to an internal value.
const AddSubscriptionPriceParamsPriceNewSubscriptionMatrixWithDisplayNamePriceBillingCycleConfigurationDurationUnitDay = shared.AddSubscriptionPriceParamsPriceNewSubscriptionMatrixWithDisplayNamePriceBillingCycleConfigurationDurationUnitDay

// This is an alias to an internal value.
const AddSubscriptionPriceParamsPriceNewSubscriptionMatrixWithDisplayNamePriceBillingCycleConfigurationDurationUnitMonth = shared.AddSubscriptionPriceParamsPriceNewSubscriptionMatrixWithDisplayNamePriceBillingCycleConfigurationDurationUnitMonth

// Within each billing cycle, specifies the cadence at which invoices are produced.
// If unspecified, a single invoice is produced per billing cycle.
//
// This is an alias to an internal type.
type AddSubscriptionPriceParamsPriceNewSubscriptionMatrixWithDisplayNamePriceInvoicingCycleConfiguration = shared.AddSubscriptionPriceParamsPriceNewSubscriptionMatrixWithDisplayNamePriceInvoicingCycleConfiguration

// The unit of billing period duration.
//
// This is an alias to an internal type.
type AddSubscriptionPriceParamsPriceNewSubscriptionMatrixWithDisplayNamePriceInvoicingCycleConfigurationDurationUnit = shared.AddSubscriptionPriceParamsPriceNewSubscriptionMatrixWithDisplayNamePriceInvoicingCycleConfigurationDurationUnit

// This is an alias to an internal value.
const AddSubscriptionPriceParamsPriceNewSubscriptionMatrixWithDisplayNamePriceInvoicingCycleConfigurationDurationUnitDay = shared.AddSubscriptionPriceParamsPriceNewSubscriptionMatrixWithDisplayNamePriceInvoicingCycleConfigurationDurationUnitDay

// This is an alias to an internal value.
const AddSubscriptionPriceParamsPriceNewSubscriptionMatrixWithDisplayNamePriceInvoicingCycleConfigurationDurationUnitMonth = shared.AddSubscriptionPriceParamsPriceNewSubscriptionMatrixWithDisplayNamePriceInvoicingCycleConfigurationDurationUnitMonth

// This is an alias to an internal type.
type AddSubscriptionPriceParamsPriceNewSubscriptionGroupedTieredPackagePrice = shared.AddSubscriptionPriceParamsPriceNewSubscriptionGroupedTieredPackagePrice

// The cadence to bill for this price on.
//
// This is an alias to an internal type.
type AddSubscriptionPriceParamsPriceNewSubscriptionGroupedTieredPackagePriceCadence = shared.AddSubscriptionPriceParamsPriceNewSubscriptionGroupedTieredPackagePriceCadence

// This is an alias to an internal value.
const AddSubscriptionPriceParamsPriceNewSubscriptionGroupedTieredPackagePriceCadenceAnnual = shared.AddSubscriptionPriceParamsPriceNewSubscriptionGroupedTieredPackagePriceCadenceAnnual

// This is an alias to an internal value.
const AddSubscriptionPriceParamsPriceNewSubscriptionGroupedTieredPackagePriceCadenceSemiAnnual = shared.AddSubscriptionPriceParamsPriceNewSubscriptionGroupedTieredPackagePriceCadenceSemiAnnual

// This is an alias to an internal value.
const AddSubscriptionPriceParamsPriceNewSubscriptionGroupedTieredPackagePriceCadenceMonthly = shared.AddSubscriptionPriceParamsPriceNewSubscriptionGroupedTieredPackagePriceCadenceMonthly

// This is an alias to an internal value.
const AddSubscriptionPriceParamsPriceNewSubscriptionGroupedTieredPackagePriceCadenceQuarterly = shared.AddSubscriptionPriceParamsPriceNewSubscriptionGroupedTieredPackagePriceCadenceQuarterly

// This is an alias to an internal value.
const AddSubscriptionPriceParamsPriceNewSubscriptionGroupedTieredPackagePriceCadenceOneTime = shared.AddSubscriptionPriceParamsPriceNewSubscriptionGroupedTieredPackagePriceCadenceOneTime

// This is an alias to an internal value.
const AddSubscriptionPriceParamsPriceNewSubscriptionGroupedTieredPackagePriceCadenceCustom = shared.AddSubscriptionPriceParamsPriceNewSubscriptionGroupedTieredPackagePriceCadenceCustom

// This is an alias to an internal type.
type AddSubscriptionPriceParamsPriceNewSubscriptionGroupedTieredPackagePriceModelType = shared.AddSubscriptionPriceParamsPriceNewSubscriptionGroupedTieredPackagePriceModelType

// This is an alias to an internal value.
const AddSubscriptionPriceParamsPriceNewSubscriptionGroupedTieredPackagePriceModelTypeGroupedTieredPackage = shared.AddSubscriptionPriceParamsPriceNewSubscriptionGroupedTieredPackagePriceModelTypeGroupedTieredPackage

// For custom cadence: specifies the duration of the billing period in days or
// months.
//
// This is an alias to an internal type.
type AddSubscriptionPriceParamsPriceNewSubscriptionGroupedTieredPackagePriceBillingCycleConfiguration = shared.AddSubscriptionPriceParamsPriceNewSubscriptionGroupedTieredPackagePriceBillingCycleConfiguration

// The unit of billing period duration.
//
// This is an alias to an internal type.
type AddSubscriptionPriceParamsPriceNewSubscriptionGroupedTieredPackagePriceBillingCycleConfigurationDurationUnit = shared.AddSubscriptionPriceParamsPriceNewSubscriptionGroupedTieredPackagePriceBillingCycleConfigurationDurationUnit

// This is an alias to an internal value.
const AddSubscriptionPriceParamsPriceNewSubscriptionGroupedTieredPackagePriceBillingCycleConfigurationDurationUnitDay = shared.AddSubscriptionPriceParamsPriceNewSubscriptionGroupedTieredPackagePriceBillingCycleConfigurationDurationUnitDay

// This is an alias to an internal value.
const AddSubscriptionPriceParamsPriceNewSubscriptionGroupedTieredPackagePriceBillingCycleConfigurationDurationUnitMonth = shared.AddSubscriptionPriceParamsPriceNewSubscriptionGroupedTieredPackagePriceBillingCycleConfigurationDurationUnitMonth

// Within each billing cycle, specifies the cadence at which invoices are produced.
// If unspecified, a single invoice is produced per billing cycle.
//
// This is an alias to an internal type.
type AddSubscriptionPriceParamsPriceNewSubscriptionGroupedTieredPackagePriceInvoicingCycleConfiguration = shared.AddSubscriptionPriceParamsPriceNewSubscriptionGroupedTieredPackagePriceInvoicingCycleConfiguration

// The unit of billing period duration.
//
// This is an alias to an internal type.
type AddSubscriptionPriceParamsPriceNewSubscriptionGroupedTieredPackagePriceInvoicingCycleConfigurationDurationUnit = shared.AddSubscriptionPriceParamsPriceNewSubscriptionGroupedTieredPackagePriceInvoicingCycleConfigurationDurationUnit

// This is an alias to an internal value.
const AddSubscriptionPriceParamsPriceNewSubscriptionGroupedTieredPackagePriceInvoicingCycleConfigurationDurationUnitDay = shared.AddSubscriptionPriceParamsPriceNewSubscriptionGroupedTieredPackagePriceInvoicingCycleConfigurationDurationUnitDay

// This is an alias to an internal value.
const AddSubscriptionPriceParamsPriceNewSubscriptionGroupedTieredPackagePriceInvoicingCycleConfigurationDurationUnitMonth = shared.AddSubscriptionPriceParamsPriceNewSubscriptionGroupedTieredPackagePriceInvoicingCycleConfigurationDurationUnitMonth

// The cadence to bill for this price on.
//
// This is an alias to an internal type.
type AddSubscriptionPriceParamsPriceCadence = shared.AddSubscriptionPriceParamsPriceCadence

// This is an alias to an internal value.
const AddSubscriptionPriceParamsPriceCadenceAnnual = shared.AddSubscriptionPriceParamsPriceCadenceAnnual

// This is an alias to an internal value.
const AddSubscriptionPriceParamsPriceCadenceSemiAnnual = shared.AddSubscriptionPriceParamsPriceCadenceSemiAnnual

// This is an alias to an internal value.
const AddSubscriptionPriceParamsPriceCadenceMonthly = shared.AddSubscriptionPriceParamsPriceCadenceMonthly

// This is an alias to an internal value.
const AddSubscriptionPriceParamsPriceCadenceQuarterly = shared.AddSubscriptionPriceParamsPriceCadenceQuarterly

// This is an alias to an internal value.
const AddSubscriptionPriceParamsPriceCadenceOneTime = shared.AddSubscriptionPriceParamsPriceCadenceOneTime

// This is an alias to an internal value.
const AddSubscriptionPriceParamsPriceCadenceCustom = shared.AddSubscriptionPriceParamsPriceCadenceCustom

// This is an alias to an internal type.
type AddSubscriptionPriceParamsPriceModelType = shared.AddSubscriptionPriceParamsPriceModelType

// This is an alias to an internal value.
const AddSubscriptionPriceParamsPriceModelTypeUnit = shared.AddSubscriptionPriceParamsPriceModelTypeUnit

// This is an alias to an internal value.
const AddSubscriptionPriceParamsPriceModelTypePackage = shared.AddSubscriptionPriceParamsPriceModelTypePackage

// This is an alias to an internal value.
const AddSubscriptionPriceParamsPriceModelTypeMatrix = shared.AddSubscriptionPriceParamsPriceModelTypeMatrix

// This is an alias to an internal value.
const AddSubscriptionPriceParamsPriceModelTypeTiered = shared.AddSubscriptionPriceParamsPriceModelTypeTiered

// This is an alias to an internal value.
const AddSubscriptionPriceParamsPriceModelTypeTieredBps = shared.AddSubscriptionPriceParamsPriceModelTypeTieredBps

// This is an alias to an internal value.
const AddSubscriptionPriceParamsPriceModelTypeBps = shared.AddSubscriptionPriceParamsPriceModelTypeBps

// This is an alias to an internal value.
const AddSubscriptionPriceParamsPriceModelTypeBulkBps = shared.AddSubscriptionPriceParamsPriceModelTypeBulkBps

// This is an alias to an internal value.
const AddSubscriptionPriceParamsPriceModelTypeBulk = shared.AddSubscriptionPriceParamsPriceModelTypeBulk

// This is an alias to an internal value.
const AddSubscriptionPriceParamsPriceModelTypeThresholdTotalAmount = shared.AddSubscriptionPriceParamsPriceModelTypeThresholdTotalAmount

// This is an alias to an internal value.
const AddSubscriptionPriceParamsPriceModelTypeTieredPackage = shared.AddSubscriptionPriceParamsPriceModelTypeTieredPackage

// This is an alias to an internal value.
const AddSubscriptionPriceParamsPriceModelTypeTieredWithMinimum = shared.AddSubscriptionPriceParamsPriceModelTypeTieredWithMinimum

// This is an alias to an internal value.
const AddSubscriptionPriceParamsPriceModelTypeUnitWithPercent = shared.AddSubscriptionPriceParamsPriceModelTypeUnitWithPercent

// This is an alias to an internal value.
const AddSubscriptionPriceParamsPriceModelTypePackageWithAllocation = shared.AddSubscriptionPriceParamsPriceModelTypePackageWithAllocation

// This is an alias to an internal value.
const AddSubscriptionPriceParamsPriceModelTypeTieredWithProration = shared.AddSubscriptionPriceParamsPriceModelTypeTieredWithProration

// This is an alias to an internal value.
const AddSubscriptionPriceParamsPriceModelTypeUnitWithProration = shared.AddSubscriptionPriceParamsPriceModelTypeUnitWithProration

// This is an alias to an internal value.
const AddSubscriptionPriceParamsPriceModelTypeGroupedAllocation = shared.AddSubscriptionPriceParamsPriceModelTypeGroupedAllocation

// This is an alias to an internal value.
const AddSubscriptionPriceParamsPriceModelTypeGroupedWithProratedMinimum = shared.AddSubscriptionPriceParamsPriceModelTypeGroupedWithProratedMinimum

// This is an alias to an internal value.
const AddSubscriptionPriceParamsPriceModelTypeBulkWithProration = shared.AddSubscriptionPriceParamsPriceModelTypeBulkWithProration

// This is an alias to an internal value.
const AddSubscriptionPriceParamsPriceModelTypeScalableMatrixWithUnitPricing = shared.AddSubscriptionPriceParamsPriceModelTypeScalableMatrixWithUnitPricing

// This is an alias to an internal value.
const AddSubscriptionPriceParamsPriceModelTypeScalableMatrixWithTieredPricing = shared.AddSubscriptionPriceParamsPriceModelTypeScalableMatrixWithTieredPricing

// This is an alias to an internal value.
const AddSubscriptionPriceParamsPriceModelTypeCumulativeGroupedBulk = shared.AddSubscriptionPriceParamsPriceModelTypeCumulativeGroupedBulk

// This is an alias to an internal value.
const AddSubscriptionPriceParamsPriceModelTypeMaxGroupTieredPackage = shared.AddSubscriptionPriceParamsPriceModelTypeMaxGroupTieredPackage

// This is an alias to an internal value.
const AddSubscriptionPriceParamsPriceModelTypeGroupedWithMeteredMinimum = shared.AddSubscriptionPriceParamsPriceModelTypeGroupedWithMeteredMinimum

// This is an alias to an internal value.
const AddSubscriptionPriceParamsPriceModelTypeMatrixWithDisplayName = shared.AddSubscriptionPriceParamsPriceModelTypeMatrixWithDisplayName

// This is an alias to an internal value.
const AddSubscriptionPriceParamsPriceModelTypeGroupedTieredPackage = shared.AddSubscriptionPriceParamsPriceModelTypeGroupedTieredPackage

// This is an alias to an internal type.
type AmountDiscount = shared.AmountDiscount

// This is an alias to an internal type.
type AmountDiscountDiscountType = shared.AmountDiscountDiscountType

// This is an alias to an internal value.
const AmountDiscountDiscountTypeAmount = shared.AmountDiscountDiscountTypeAmount

// This is an alias to an internal type.
type AmountDiscountParam = shared.AmountDiscountParam

// This is an alias to an internal type.
type BillingCycleRelativeDate = shared.BillingCycleRelativeDate

// This is an alias to an internal value.
const BillingCycleRelativeDateStartOfTerm = shared.BillingCycleRelativeDateStartOfTerm

// This is an alias to an internal value.
const BillingCycleRelativeDateEndOfTerm = shared.BillingCycleRelativeDateEndOfTerm

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
type PaginationMetadata = shared.PaginationMetadata

// This is an alias to an internal type.
type PercentageDiscount = shared.PercentageDiscount

// This is an alias to an internal type.
type PercentageDiscountDiscountType = shared.PercentageDiscountDiscountType

// This is an alias to an internal value.
const PercentageDiscountDiscountTypePercentage = shared.PercentageDiscountDiscountTypePercentage

// This is an alias to an internal type.
type PercentageDiscountParam = shared.PercentageDiscountParam

// This is an alias to an internal type.
type RemoveSubscriptionAdjustmentParams = shared.RemoveSubscriptionAdjustmentParams

// This is an alias to an internal type.
type RemoveSubscriptionPriceParams = shared.RemoveSubscriptionPriceParams

// This is an alias to an internal type.
type ReplaceSubscriptionAdjustmentParams = shared.ReplaceSubscriptionAdjustmentParams

// The definition of a new adjustment to create and add to the subscription.
//
// This is an alias to an internal type.
type ReplaceSubscriptionAdjustmentParamsAdjustmentUnion = shared.ReplaceSubscriptionAdjustmentParamsAdjustmentUnion

// This is an alias to an internal type.
type ReplaceSubscriptionAdjustmentParamsAdjustmentNewPercentageDiscount = shared.ReplaceSubscriptionAdjustmentParamsAdjustmentNewPercentageDiscount

// This is an alias to an internal type.
type ReplaceSubscriptionAdjustmentParamsAdjustmentNewPercentageDiscountAdjustmentType = shared.ReplaceSubscriptionAdjustmentParamsAdjustmentNewPercentageDiscountAdjustmentType

// This is an alias to an internal value.
const ReplaceSubscriptionAdjustmentParamsAdjustmentNewPercentageDiscountAdjustmentTypePercentageDiscount = shared.ReplaceSubscriptionAdjustmentParamsAdjustmentNewPercentageDiscountAdjustmentTypePercentageDiscount

// This is an alias to an internal type.
type ReplaceSubscriptionAdjustmentParamsAdjustmentNewUsageDiscount = shared.ReplaceSubscriptionAdjustmentParamsAdjustmentNewUsageDiscount

// This is an alias to an internal type.
type ReplaceSubscriptionAdjustmentParamsAdjustmentNewUsageDiscountAdjustmentType = shared.ReplaceSubscriptionAdjustmentParamsAdjustmentNewUsageDiscountAdjustmentType

// This is an alias to an internal value.
const ReplaceSubscriptionAdjustmentParamsAdjustmentNewUsageDiscountAdjustmentTypeUsageDiscount = shared.ReplaceSubscriptionAdjustmentParamsAdjustmentNewUsageDiscountAdjustmentTypeUsageDiscount

// This is an alias to an internal type.
type ReplaceSubscriptionAdjustmentParamsAdjustmentNewAmountDiscount = shared.ReplaceSubscriptionAdjustmentParamsAdjustmentNewAmountDiscount

// This is an alias to an internal type.
type ReplaceSubscriptionAdjustmentParamsAdjustmentNewAmountDiscountAdjustmentType = shared.ReplaceSubscriptionAdjustmentParamsAdjustmentNewAmountDiscountAdjustmentType

// This is an alias to an internal value.
const ReplaceSubscriptionAdjustmentParamsAdjustmentNewAmountDiscountAdjustmentTypeAmountDiscount = shared.ReplaceSubscriptionAdjustmentParamsAdjustmentNewAmountDiscountAdjustmentTypeAmountDiscount

// This is an alias to an internal type.
type ReplaceSubscriptionAdjustmentParamsAdjustmentNewMinimum = shared.ReplaceSubscriptionAdjustmentParamsAdjustmentNewMinimum

// This is an alias to an internal type.
type ReplaceSubscriptionAdjustmentParamsAdjustmentNewMinimumAdjustmentType = shared.ReplaceSubscriptionAdjustmentParamsAdjustmentNewMinimumAdjustmentType

// This is an alias to an internal value.
const ReplaceSubscriptionAdjustmentParamsAdjustmentNewMinimumAdjustmentTypeMinimum = shared.ReplaceSubscriptionAdjustmentParamsAdjustmentNewMinimumAdjustmentTypeMinimum

// This is an alias to an internal type.
type ReplaceSubscriptionAdjustmentParamsAdjustmentNewMaximum = shared.ReplaceSubscriptionAdjustmentParamsAdjustmentNewMaximum

// This is an alias to an internal type.
type ReplaceSubscriptionAdjustmentParamsAdjustmentNewMaximumAdjustmentType = shared.ReplaceSubscriptionAdjustmentParamsAdjustmentNewMaximumAdjustmentType

// This is an alias to an internal value.
const ReplaceSubscriptionAdjustmentParamsAdjustmentNewMaximumAdjustmentTypeMaximum = shared.ReplaceSubscriptionAdjustmentParamsAdjustmentNewMaximumAdjustmentTypeMaximum

// This is an alias to an internal type.
type ReplaceSubscriptionAdjustmentParamsAdjustmentAdjustmentType = shared.ReplaceSubscriptionAdjustmentParamsAdjustmentAdjustmentType

// This is an alias to an internal value.
const ReplaceSubscriptionAdjustmentParamsAdjustmentAdjustmentTypePercentageDiscount = shared.ReplaceSubscriptionAdjustmentParamsAdjustmentAdjustmentTypePercentageDiscount

// This is an alias to an internal value.
const ReplaceSubscriptionAdjustmentParamsAdjustmentAdjustmentTypeUsageDiscount = shared.ReplaceSubscriptionAdjustmentParamsAdjustmentAdjustmentTypeUsageDiscount

// This is an alias to an internal value.
const ReplaceSubscriptionAdjustmentParamsAdjustmentAdjustmentTypeAmountDiscount = shared.ReplaceSubscriptionAdjustmentParamsAdjustmentAdjustmentTypeAmountDiscount

// This is an alias to an internal value.
const ReplaceSubscriptionAdjustmentParamsAdjustmentAdjustmentTypeMinimum = shared.ReplaceSubscriptionAdjustmentParamsAdjustmentAdjustmentTypeMinimum

// This is an alias to an internal value.
const ReplaceSubscriptionAdjustmentParamsAdjustmentAdjustmentTypeMaximum = shared.ReplaceSubscriptionAdjustmentParamsAdjustmentAdjustmentTypeMaximum

// This is an alias to an internal type.
type ReplaceSubscriptionPriceParams = shared.ReplaceSubscriptionPriceParams

// The definition of a new allocation price to create and add to the subscription.
//
// This is an alias to an internal type.
type ReplaceSubscriptionPriceParamsAllocationPrice = shared.ReplaceSubscriptionPriceParamsAllocationPrice

// The cadence at which to allocate the amount to the customer.
//
// This is an alias to an internal type.
type ReplaceSubscriptionPriceParamsAllocationPriceCadence = shared.ReplaceSubscriptionPriceParamsAllocationPriceCadence

// This is an alias to an internal value.
const ReplaceSubscriptionPriceParamsAllocationPriceCadenceOneTime = shared.ReplaceSubscriptionPriceParamsAllocationPriceCadenceOneTime

// This is an alias to an internal value.
const ReplaceSubscriptionPriceParamsAllocationPriceCadenceMonthly = shared.ReplaceSubscriptionPriceParamsAllocationPriceCadenceMonthly

// This is an alias to an internal value.
const ReplaceSubscriptionPriceParamsAllocationPriceCadenceQuarterly = shared.ReplaceSubscriptionPriceParamsAllocationPriceCadenceQuarterly

// This is an alias to an internal value.
const ReplaceSubscriptionPriceParamsAllocationPriceCadenceSemiAnnual = shared.ReplaceSubscriptionPriceParamsAllocationPriceCadenceSemiAnnual

// This is an alias to an internal value.
const ReplaceSubscriptionPriceParamsAllocationPriceCadenceAnnual = shared.ReplaceSubscriptionPriceParamsAllocationPriceCadenceAnnual

// This is an alias to an internal value.
const ReplaceSubscriptionPriceParamsAllocationPriceCadenceCustom = shared.ReplaceSubscriptionPriceParamsAllocationPriceCadenceCustom

// This is an alias to an internal type.
type ReplaceSubscriptionPriceParamsDiscount = shared.ReplaceSubscriptionPriceParamsDiscount

// This is an alias to an internal type.
type ReplaceSubscriptionPriceParamsDiscountsDiscountType = shared.ReplaceSubscriptionPriceParamsDiscountsDiscountType

// This is an alias to an internal value.
const ReplaceSubscriptionPriceParamsDiscountsDiscountTypePercentage = shared.ReplaceSubscriptionPriceParamsDiscountsDiscountTypePercentage

// This is an alias to an internal value.
const ReplaceSubscriptionPriceParamsDiscountsDiscountTypeUsage = shared.ReplaceSubscriptionPriceParamsDiscountsDiscountTypeUsage

// This is an alias to an internal value.
const ReplaceSubscriptionPriceParamsDiscountsDiscountTypeAmount = shared.ReplaceSubscriptionPriceParamsDiscountsDiscountTypeAmount

// The definition of a new price to create and add to the subscription.
//
// This is an alias to an internal type.
type ReplaceSubscriptionPriceParamsPriceUnion = shared.ReplaceSubscriptionPriceParamsPriceUnion

// This is an alias to an internal type.
type ReplaceSubscriptionPriceParamsPriceNewSubscriptionUnitPrice = shared.ReplaceSubscriptionPriceParamsPriceNewSubscriptionUnitPrice

// The cadence to bill for this price on.
//
// This is an alias to an internal type.
type ReplaceSubscriptionPriceParamsPriceNewSubscriptionUnitPriceCadence = shared.ReplaceSubscriptionPriceParamsPriceNewSubscriptionUnitPriceCadence

// This is an alias to an internal value.
const ReplaceSubscriptionPriceParamsPriceNewSubscriptionUnitPriceCadenceAnnual = shared.ReplaceSubscriptionPriceParamsPriceNewSubscriptionUnitPriceCadenceAnnual

// This is an alias to an internal value.
const ReplaceSubscriptionPriceParamsPriceNewSubscriptionUnitPriceCadenceSemiAnnual = shared.ReplaceSubscriptionPriceParamsPriceNewSubscriptionUnitPriceCadenceSemiAnnual

// This is an alias to an internal value.
const ReplaceSubscriptionPriceParamsPriceNewSubscriptionUnitPriceCadenceMonthly = shared.ReplaceSubscriptionPriceParamsPriceNewSubscriptionUnitPriceCadenceMonthly

// This is an alias to an internal value.
const ReplaceSubscriptionPriceParamsPriceNewSubscriptionUnitPriceCadenceQuarterly = shared.ReplaceSubscriptionPriceParamsPriceNewSubscriptionUnitPriceCadenceQuarterly

// This is an alias to an internal value.
const ReplaceSubscriptionPriceParamsPriceNewSubscriptionUnitPriceCadenceOneTime = shared.ReplaceSubscriptionPriceParamsPriceNewSubscriptionUnitPriceCadenceOneTime

// This is an alias to an internal value.
const ReplaceSubscriptionPriceParamsPriceNewSubscriptionUnitPriceCadenceCustom = shared.ReplaceSubscriptionPriceParamsPriceNewSubscriptionUnitPriceCadenceCustom

// This is an alias to an internal type.
type ReplaceSubscriptionPriceParamsPriceNewSubscriptionUnitPriceModelType = shared.ReplaceSubscriptionPriceParamsPriceNewSubscriptionUnitPriceModelType

// This is an alias to an internal value.
const ReplaceSubscriptionPriceParamsPriceNewSubscriptionUnitPriceModelTypeUnit = shared.ReplaceSubscriptionPriceParamsPriceNewSubscriptionUnitPriceModelTypeUnit

// This is an alias to an internal type.
type ReplaceSubscriptionPriceParamsPriceNewSubscriptionUnitPriceUnitConfig = shared.ReplaceSubscriptionPriceParamsPriceNewSubscriptionUnitPriceUnitConfig

// For custom cadence: specifies the duration of the billing period in days or
// months.
//
// This is an alias to an internal type.
type ReplaceSubscriptionPriceParamsPriceNewSubscriptionUnitPriceBillingCycleConfiguration = shared.ReplaceSubscriptionPriceParamsPriceNewSubscriptionUnitPriceBillingCycleConfiguration

// The unit of billing period duration.
//
// This is an alias to an internal type.
type ReplaceSubscriptionPriceParamsPriceNewSubscriptionUnitPriceBillingCycleConfigurationDurationUnit = shared.ReplaceSubscriptionPriceParamsPriceNewSubscriptionUnitPriceBillingCycleConfigurationDurationUnit

// This is an alias to an internal value.
const ReplaceSubscriptionPriceParamsPriceNewSubscriptionUnitPriceBillingCycleConfigurationDurationUnitDay = shared.ReplaceSubscriptionPriceParamsPriceNewSubscriptionUnitPriceBillingCycleConfigurationDurationUnitDay

// This is an alias to an internal value.
const ReplaceSubscriptionPriceParamsPriceNewSubscriptionUnitPriceBillingCycleConfigurationDurationUnitMonth = shared.ReplaceSubscriptionPriceParamsPriceNewSubscriptionUnitPriceBillingCycleConfigurationDurationUnitMonth

// Within each billing cycle, specifies the cadence at which invoices are produced.
// If unspecified, a single invoice is produced per billing cycle.
//
// This is an alias to an internal type.
type ReplaceSubscriptionPriceParamsPriceNewSubscriptionUnitPriceInvoicingCycleConfiguration = shared.ReplaceSubscriptionPriceParamsPriceNewSubscriptionUnitPriceInvoicingCycleConfiguration

// The unit of billing period duration.
//
// This is an alias to an internal type.
type ReplaceSubscriptionPriceParamsPriceNewSubscriptionUnitPriceInvoicingCycleConfigurationDurationUnit = shared.ReplaceSubscriptionPriceParamsPriceNewSubscriptionUnitPriceInvoicingCycleConfigurationDurationUnit

// This is an alias to an internal value.
const ReplaceSubscriptionPriceParamsPriceNewSubscriptionUnitPriceInvoicingCycleConfigurationDurationUnitDay = shared.ReplaceSubscriptionPriceParamsPriceNewSubscriptionUnitPriceInvoicingCycleConfigurationDurationUnitDay

// This is an alias to an internal value.
const ReplaceSubscriptionPriceParamsPriceNewSubscriptionUnitPriceInvoicingCycleConfigurationDurationUnitMonth = shared.ReplaceSubscriptionPriceParamsPriceNewSubscriptionUnitPriceInvoicingCycleConfigurationDurationUnitMonth

// This is an alias to an internal type.
type ReplaceSubscriptionPriceParamsPriceNewSubscriptionPackagePrice = shared.ReplaceSubscriptionPriceParamsPriceNewSubscriptionPackagePrice

// The cadence to bill for this price on.
//
// This is an alias to an internal type.
type ReplaceSubscriptionPriceParamsPriceNewSubscriptionPackagePriceCadence = shared.ReplaceSubscriptionPriceParamsPriceNewSubscriptionPackagePriceCadence

// This is an alias to an internal value.
const ReplaceSubscriptionPriceParamsPriceNewSubscriptionPackagePriceCadenceAnnual = shared.ReplaceSubscriptionPriceParamsPriceNewSubscriptionPackagePriceCadenceAnnual

// This is an alias to an internal value.
const ReplaceSubscriptionPriceParamsPriceNewSubscriptionPackagePriceCadenceSemiAnnual = shared.ReplaceSubscriptionPriceParamsPriceNewSubscriptionPackagePriceCadenceSemiAnnual

// This is an alias to an internal value.
const ReplaceSubscriptionPriceParamsPriceNewSubscriptionPackagePriceCadenceMonthly = shared.ReplaceSubscriptionPriceParamsPriceNewSubscriptionPackagePriceCadenceMonthly

// This is an alias to an internal value.
const ReplaceSubscriptionPriceParamsPriceNewSubscriptionPackagePriceCadenceQuarterly = shared.ReplaceSubscriptionPriceParamsPriceNewSubscriptionPackagePriceCadenceQuarterly

// This is an alias to an internal value.
const ReplaceSubscriptionPriceParamsPriceNewSubscriptionPackagePriceCadenceOneTime = shared.ReplaceSubscriptionPriceParamsPriceNewSubscriptionPackagePriceCadenceOneTime

// This is an alias to an internal value.
const ReplaceSubscriptionPriceParamsPriceNewSubscriptionPackagePriceCadenceCustom = shared.ReplaceSubscriptionPriceParamsPriceNewSubscriptionPackagePriceCadenceCustom

// This is an alias to an internal type.
type ReplaceSubscriptionPriceParamsPriceNewSubscriptionPackagePriceModelType = shared.ReplaceSubscriptionPriceParamsPriceNewSubscriptionPackagePriceModelType

// This is an alias to an internal value.
const ReplaceSubscriptionPriceParamsPriceNewSubscriptionPackagePriceModelTypePackage = shared.ReplaceSubscriptionPriceParamsPriceNewSubscriptionPackagePriceModelTypePackage

// This is an alias to an internal type.
type ReplaceSubscriptionPriceParamsPriceNewSubscriptionPackagePricePackageConfig = shared.ReplaceSubscriptionPriceParamsPriceNewSubscriptionPackagePricePackageConfig

// For custom cadence: specifies the duration of the billing period in days or
// months.
//
// This is an alias to an internal type.
type ReplaceSubscriptionPriceParamsPriceNewSubscriptionPackagePriceBillingCycleConfiguration = shared.ReplaceSubscriptionPriceParamsPriceNewSubscriptionPackagePriceBillingCycleConfiguration

// The unit of billing period duration.
//
// This is an alias to an internal type.
type ReplaceSubscriptionPriceParamsPriceNewSubscriptionPackagePriceBillingCycleConfigurationDurationUnit = shared.ReplaceSubscriptionPriceParamsPriceNewSubscriptionPackagePriceBillingCycleConfigurationDurationUnit

// This is an alias to an internal value.
const ReplaceSubscriptionPriceParamsPriceNewSubscriptionPackagePriceBillingCycleConfigurationDurationUnitDay = shared.ReplaceSubscriptionPriceParamsPriceNewSubscriptionPackagePriceBillingCycleConfigurationDurationUnitDay

// This is an alias to an internal value.
const ReplaceSubscriptionPriceParamsPriceNewSubscriptionPackagePriceBillingCycleConfigurationDurationUnitMonth = shared.ReplaceSubscriptionPriceParamsPriceNewSubscriptionPackagePriceBillingCycleConfigurationDurationUnitMonth

// Within each billing cycle, specifies the cadence at which invoices are produced.
// If unspecified, a single invoice is produced per billing cycle.
//
// This is an alias to an internal type.
type ReplaceSubscriptionPriceParamsPriceNewSubscriptionPackagePriceInvoicingCycleConfiguration = shared.ReplaceSubscriptionPriceParamsPriceNewSubscriptionPackagePriceInvoicingCycleConfiguration

// The unit of billing period duration.
//
// This is an alias to an internal type.
type ReplaceSubscriptionPriceParamsPriceNewSubscriptionPackagePriceInvoicingCycleConfigurationDurationUnit = shared.ReplaceSubscriptionPriceParamsPriceNewSubscriptionPackagePriceInvoicingCycleConfigurationDurationUnit

// This is an alias to an internal value.
const ReplaceSubscriptionPriceParamsPriceNewSubscriptionPackagePriceInvoicingCycleConfigurationDurationUnitDay = shared.ReplaceSubscriptionPriceParamsPriceNewSubscriptionPackagePriceInvoicingCycleConfigurationDurationUnitDay

// This is an alias to an internal value.
const ReplaceSubscriptionPriceParamsPriceNewSubscriptionPackagePriceInvoicingCycleConfigurationDurationUnitMonth = shared.ReplaceSubscriptionPriceParamsPriceNewSubscriptionPackagePriceInvoicingCycleConfigurationDurationUnitMonth

// This is an alias to an internal type.
type ReplaceSubscriptionPriceParamsPriceNewSubscriptionMatrixPrice = shared.ReplaceSubscriptionPriceParamsPriceNewSubscriptionMatrixPrice

// The cadence to bill for this price on.
//
// This is an alias to an internal type.
type ReplaceSubscriptionPriceParamsPriceNewSubscriptionMatrixPriceCadence = shared.ReplaceSubscriptionPriceParamsPriceNewSubscriptionMatrixPriceCadence

// This is an alias to an internal value.
const ReplaceSubscriptionPriceParamsPriceNewSubscriptionMatrixPriceCadenceAnnual = shared.ReplaceSubscriptionPriceParamsPriceNewSubscriptionMatrixPriceCadenceAnnual

// This is an alias to an internal value.
const ReplaceSubscriptionPriceParamsPriceNewSubscriptionMatrixPriceCadenceSemiAnnual = shared.ReplaceSubscriptionPriceParamsPriceNewSubscriptionMatrixPriceCadenceSemiAnnual

// This is an alias to an internal value.
const ReplaceSubscriptionPriceParamsPriceNewSubscriptionMatrixPriceCadenceMonthly = shared.ReplaceSubscriptionPriceParamsPriceNewSubscriptionMatrixPriceCadenceMonthly

// This is an alias to an internal value.
const ReplaceSubscriptionPriceParamsPriceNewSubscriptionMatrixPriceCadenceQuarterly = shared.ReplaceSubscriptionPriceParamsPriceNewSubscriptionMatrixPriceCadenceQuarterly

// This is an alias to an internal value.
const ReplaceSubscriptionPriceParamsPriceNewSubscriptionMatrixPriceCadenceOneTime = shared.ReplaceSubscriptionPriceParamsPriceNewSubscriptionMatrixPriceCadenceOneTime

// This is an alias to an internal value.
const ReplaceSubscriptionPriceParamsPriceNewSubscriptionMatrixPriceCadenceCustom = shared.ReplaceSubscriptionPriceParamsPriceNewSubscriptionMatrixPriceCadenceCustom

// This is an alias to an internal type.
type ReplaceSubscriptionPriceParamsPriceNewSubscriptionMatrixPriceMatrixConfig = shared.ReplaceSubscriptionPriceParamsPriceNewSubscriptionMatrixPriceMatrixConfig

// This is an alias to an internal type.
type ReplaceSubscriptionPriceParamsPriceNewSubscriptionMatrixPriceMatrixConfigMatrixValue = shared.ReplaceSubscriptionPriceParamsPriceNewSubscriptionMatrixPriceMatrixConfigMatrixValue

// This is an alias to an internal type.
type ReplaceSubscriptionPriceParamsPriceNewSubscriptionMatrixPriceModelType = shared.ReplaceSubscriptionPriceParamsPriceNewSubscriptionMatrixPriceModelType

// This is an alias to an internal value.
const ReplaceSubscriptionPriceParamsPriceNewSubscriptionMatrixPriceModelTypeMatrix = shared.ReplaceSubscriptionPriceParamsPriceNewSubscriptionMatrixPriceModelTypeMatrix

// For custom cadence: specifies the duration of the billing period in days or
// months.
//
// This is an alias to an internal type.
type ReplaceSubscriptionPriceParamsPriceNewSubscriptionMatrixPriceBillingCycleConfiguration = shared.ReplaceSubscriptionPriceParamsPriceNewSubscriptionMatrixPriceBillingCycleConfiguration

// The unit of billing period duration.
//
// This is an alias to an internal type.
type ReplaceSubscriptionPriceParamsPriceNewSubscriptionMatrixPriceBillingCycleConfigurationDurationUnit = shared.ReplaceSubscriptionPriceParamsPriceNewSubscriptionMatrixPriceBillingCycleConfigurationDurationUnit

// This is an alias to an internal value.
const ReplaceSubscriptionPriceParamsPriceNewSubscriptionMatrixPriceBillingCycleConfigurationDurationUnitDay = shared.ReplaceSubscriptionPriceParamsPriceNewSubscriptionMatrixPriceBillingCycleConfigurationDurationUnitDay

// This is an alias to an internal value.
const ReplaceSubscriptionPriceParamsPriceNewSubscriptionMatrixPriceBillingCycleConfigurationDurationUnitMonth = shared.ReplaceSubscriptionPriceParamsPriceNewSubscriptionMatrixPriceBillingCycleConfigurationDurationUnitMonth

// Within each billing cycle, specifies the cadence at which invoices are produced.
// If unspecified, a single invoice is produced per billing cycle.
//
// This is an alias to an internal type.
type ReplaceSubscriptionPriceParamsPriceNewSubscriptionMatrixPriceInvoicingCycleConfiguration = shared.ReplaceSubscriptionPriceParamsPriceNewSubscriptionMatrixPriceInvoicingCycleConfiguration

// The unit of billing period duration.
//
// This is an alias to an internal type.
type ReplaceSubscriptionPriceParamsPriceNewSubscriptionMatrixPriceInvoicingCycleConfigurationDurationUnit = shared.ReplaceSubscriptionPriceParamsPriceNewSubscriptionMatrixPriceInvoicingCycleConfigurationDurationUnit

// This is an alias to an internal value.
const ReplaceSubscriptionPriceParamsPriceNewSubscriptionMatrixPriceInvoicingCycleConfigurationDurationUnitDay = shared.ReplaceSubscriptionPriceParamsPriceNewSubscriptionMatrixPriceInvoicingCycleConfigurationDurationUnitDay

// This is an alias to an internal value.
const ReplaceSubscriptionPriceParamsPriceNewSubscriptionMatrixPriceInvoicingCycleConfigurationDurationUnitMonth = shared.ReplaceSubscriptionPriceParamsPriceNewSubscriptionMatrixPriceInvoicingCycleConfigurationDurationUnitMonth

// This is an alias to an internal type.
type ReplaceSubscriptionPriceParamsPriceNewSubscriptionTieredPrice = shared.ReplaceSubscriptionPriceParamsPriceNewSubscriptionTieredPrice

// The cadence to bill for this price on.
//
// This is an alias to an internal type.
type ReplaceSubscriptionPriceParamsPriceNewSubscriptionTieredPriceCadence = shared.ReplaceSubscriptionPriceParamsPriceNewSubscriptionTieredPriceCadence

// This is an alias to an internal value.
const ReplaceSubscriptionPriceParamsPriceNewSubscriptionTieredPriceCadenceAnnual = shared.ReplaceSubscriptionPriceParamsPriceNewSubscriptionTieredPriceCadenceAnnual

// This is an alias to an internal value.
const ReplaceSubscriptionPriceParamsPriceNewSubscriptionTieredPriceCadenceSemiAnnual = shared.ReplaceSubscriptionPriceParamsPriceNewSubscriptionTieredPriceCadenceSemiAnnual

// This is an alias to an internal value.
const ReplaceSubscriptionPriceParamsPriceNewSubscriptionTieredPriceCadenceMonthly = shared.ReplaceSubscriptionPriceParamsPriceNewSubscriptionTieredPriceCadenceMonthly

// This is an alias to an internal value.
const ReplaceSubscriptionPriceParamsPriceNewSubscriptionTieredPriceCadenceQuarterly = shared.ReplaceSubscriptionPriceParamsPriceNewSubscriptionTieredPriceCadenceQuarterly

// This is an alias to an internal value.
const ReplaceSubscriptionPriceParamsPriceNewSubscriptionTieredPriceCadenceOneTime = shared.ReplaceSubscriptionPriceParamsPriceNewSubscriptionTieredPriceCadenceOneTime

// This is an alias to an internal value.
const ReplaceSubscriptionPriceParamsPriceNewSubscriptionTieredPriceCadenceCustom = shared.ReplaceSubscriptionPriceParamsPriceNewSubscriptionTieredPriceCadenceCustom

// This is an alias to an internal type.
type ReplaceSubscriptionPriceParamsPriceNewSubscriptionTieredPriceModelType = shared.ReplaceSubscriptionPriceParamsPriceNewSubscriptionTieredPriceModelType

// This is an alias to an internal value.
const ReplaceSubscriptionPriceParamsPriceNewSubscriptionTieredPriceModelTypeTiered = shared.ReplaceSubscriptionPriceParamsPriceNewSubscriptionTieredPriceModelTypeTiered

// This is an alias to an internal type.
type ReplaceSubscriptionPriceParamsPriceNewSubscriptionTieredPriceTieredConfig = shared.ReplaceSubscriptionPriceParamsPriceNewSubscriptionTieredPriceTieredConfig

// This is an alias to an internal type.
type ReplaceSubscriptionPriceParamsPriceNewSubscriptionTieredPriceTieredConfigTier = shared.ReplaceSubscriptionPriceParamsPriceNewSubscriptionTieredPriceTieredConfigTier

// For custom cadence: specifies the duration of the billing period in days or
// months.
//
// This is an alias to an internal type.
type ReplaceSubscriptionPriceParamsPriceNewSubscriptionTieredPriceBillingCycleConfiguration = shared.ReplaceSubscriptionPriceParamsPriceNewSubscriptionTieredPriceBillingCycleConfiguration

// The unit of billing period duration.
//
// This is an alias to an internal type.
type ReplaceSubscriptionPriceParamsPriceNewSubscriptionTieredPriceBillingCycleConfigurationDurationUnit = shared.ReplaceSubscriptionPriceParamsPriceNewSubscriptionTieredPriceBillingCycleConfigurationDurationUnit

// This is an alias to an internal value.
const ReplaceSubscriptionPriceParamsPriceNewSubscriptionTieredPriceBillingCycleConfigurationDurationUnitDay = shared.ReplaceSubscriptionPriceParamsPriceNewSubscriptionTieredPriceBillingCycleConfigurationDurationUnitDay

// This is an alias to an internal value.
const ReplaceSubscriptionPriceParamsPriceNewSubscriptionTieredPriceBillingCycleConfigurationDurationUnitMonth = shared.ReplaceSubscriptionPriceParamsPriceNewSubscriptionTieredPriceBillingCycleConfigurationDurationUnitMonth

// Within each billing cycle, specifies the cadence at which invoices are produced.
// If unspecified, a single invoice is produced per billing cycle.
//
// This is an alias to an internal type.
type ReplaceSubscriptionPriceParamsPriceNewSubscriptionTieredPriceInvoicingCycleConfiguration = shared.ReplaceSubscriptionPriceParamsPriceNewSubscriptionTieredPriceInvoicingCycleConfiguration

// The unit of billing period duration.
//
// This is an alias to an internal type.
type ReplaceSubscriptionPriceParamsPriceNewSubscriptionTieredPriceInvoicingCycleConfigurationDurationUnit = shared.ReplaceSubscriptionPriceParamsPriceNewSubscriptionTieredPriceInvoicingCycleConfigurationDurationUnit

// This is an alias to an internal value.
const ReplaceSubscriptionPriceParamsPriceNewSubscriptionTieredPriceInvoicingCycleConfigurationDurationUnitDay = shared.ReplaceSubscriptionPriceParamsPriceNewSubscriptionTieredPriceInvoicingCycleConfigurationDurationUnitDay

// This is an alias to an internal value.
const ReplaceSubscriptionPriceParamsPriceNewSubscriptionTieredPriceInvoicingCycleConfigurationDurationUnitMonth = shared.ReplaceSubscriptionPriceParamsPriceNewSubscriptionTieredPriceInvoicingCycleConfigurationDurationUnitMonth

// This is an alias to an internal type.
type ReplaceSubscriptionPriceParamsPriceNewSubscriptionTieredBpsPrice = shared.ReplaceSubscriptionPriceParamsPriceNewSubscriptionTieredBpsPrice

// The cadence to bill for this price on.
//
// This is an alias to an internal type.
type ReplaceSubscriptionPriceParamsPriceNewSubscriptionTieredBpsPriceCadence = shared.ReplaceSubscriptionPriceParamsPriceNewSubscriptionTieredBpsPriceCadence

// This is an alias to an internal value.
const ReplaceSubscriptionPriceParamsPriceNewSubscriptionTieredBpsPriceCadenceAnnual = shared.ReplaceSubscriptionPriceParamsPriceNewSubscriptionTieredBpsPriceCadenceAnnual

// This is an alias to an internal value.
const ReplaceSubscriptionPriceParamsPriceNewSubscriptionTieredBpsPriceCadenceSemiAnnual = shared.ReplaceSubscriptionPriceParamsPriceNewSubscriptionTieredBpsPriceCadenceSemiAnnual

// This is an alias to an internal value.
const ReplaceSubscriptionPriceParamsPriceNewSubscriptionTieredBpsPriceCadenceMonthly = shared.ReplaceSubscriptionPriceParamsPriceNewSubscriptionTieredBpsPriceCadenceMonthly

// This is an alias to an internal value.
const ReplaceSubscriptionPriceParamsPriceNewSubscriptionTieredBpsPriceCadenceQuarterly = shared.ReplaceSubscriptionPriceParamsPriceNewSubscriptionTieredBpsPriceCadenceQuarterly

// This is an alias to an internal value.
const ReplaceSubscriptionPriceParamsPriceNewSubscriptionTieredBpsPriceCadenceOneTime = shared.ReplaceSubscriptionPriceParamsPriceNewSubscriptionTieredBpsPriceCadenceOneTime

// This is an alias to an internal value.
const ReplaceSubscriptionPriceParamsPriceNewSubscriptionTieredBpsPriceCadenceCustom = shared.ReplaceSubscriptionPriceParamsPriceNewSubscriptionTieredBpsPriceCadenceCustom

// This is an alias to an internal type.
type ReplaceSubscriptionPriceParamsPriceNewSubscriptionTieredBpsPriceModelType = shared.ReplaceSubscriptionPriceParamsPriceNewSubscriptionTieredBpsPriceModelType

// This is an alias to an internal value.
const ReplaceSubscriptionPriceParamsPriceNewSubscriptionTieredBpsPriceModelTypeTieredBps = shared.ReplaceSubscriptionPriceParamsPriceNewSubscriptionTieredBpsPriceModelTypeTieredBps

// This is an alias to an internal type.
type ReplaceSubscriptionPriceParamsPriceNewSubscriptionTieredBpsPriceTieredBpsConfig = shared.ReplaceSubscriptionPriceParamsPriceNewSubscriptionTieredBpsPriceTieredBpsConfig

// This is an alias to an internal type.
type ReplaceSubscriptionPriceParamsPriceNewSubscriptionTieredBpsPriceTieredBpsConfigTier = shared.ReplaceSubscriptionPriceParamsPriceNewSubscriptionTieredBpsPriceTieredBpsConfigTier

// For custom cadence: specifies the duration of the billing period in days or
// months.
//
// This is an alias to an internal type.
type ReplaceSubscriptionPriceParamsPriceNewSubscriptionTieredBpsPriceBillingCycleConfiguration = shared.ReplaceSubscriptionPriceParamsPriceNewSubscriptionTieredBpsPriceBillingCycleConfiguration

// The unit of billing period duration.
//
// This is an alias to an internal type.
type ReplaceSubscriptionPriceParamsPriceNewSubscriptionTieredBpsPriceBillingCycleConfigurationDurationUnit = shared.ReplaceSubscriptionPriceParamsPriceNewSubscriptionTieredBpsPriceBillingCycleConfigurationDurationUnit

// This is an alias to an internal value.
const ReplaceSubscriptionPriceParamsPriceNewSubscriptionTieredBpsPriceBillingCycleConfigurationDurationUnitDay = shared.ReplaceSubscriptionPriceParamsPriceNewSubscriptionTieredBpsPriceBillingCycleConfigurationDurationUnitDay

// This is an alias to an internal value.
const ReplaceSubscriptionPriceParamsPriceNewSubscriptionTieredBpsPriceBillingCycleConfigurationDurationUnitMonth = shared.ReplaceSubscriptionPriceParamsPriceNewSubscriptionTieredBpsPriceBillingCycleConfigurationDurationUnitMonth

// Within each billing cycle, specifies the cadence at which invoices are produced.
// If unspecified, a single invoice is produced per billing cycle.
//
// This is an alias to an internal type.
type ReplaceSubscriptionPriceParamsPriceNewSubscriptionTieredBpsPriceInvoicingCycleConfiguration = shared.ReplaceSubscriptionPriceParamsPriceNewSubscriptionTieredBpsPriceInvoicingCycleConfiguration

// The unit of billing period duration.
//
// This is an alias to an internal type.
type ReplaceSubscriptionPriceParamsPriceNewSubscriptionTieredBpsPriceInvoicingCycleConfigurationDurationUnit = shared.ReplaceSubscriptionPriceParamsPriceNewSubscriptionTieredBpsPriceInvoicingCycleConfigurationDurationUnit

// This is an alias to an internal value.
const ReplaceSubscriptionPriceParamsPriceNewSubscriptionTieredBpsPriceInvoicingCycleConfigurationDurationUnitDay = shared.ReplaceSubscriptionPriceParamsPriceNewSubscriptionTieredBpsPriceInvoicingCycleConfigurationDurationUnitDay

// This is an alias to an internal value.
const ReplaceSubscriptionPriceParamsPriceNewSubscriptionTieredBpsPriceInvoicingCycleConfigurationDurationUnitMonth = shared.ReplaceSubscriptionPriceParamsPriceNewSubscriptionTieredBpsPriceInvoicingCycleConfigurationDurationUnitMonth

// This is an alias to an internal type.
type ReplaceSubscriptionPriceParamsPriceNewSubscriptionBpsPrice = shared.ReplaceSubscriptionPriceParamsPriceNewSubscriptionBpsPrice

// This is an alias to an internal type.
type ReplaceSubscriptionPriceParamsPriceNewSubscriptionBpsPriceBpsConfig = shared.ReplaceSubscriptionPriceParamsPriceNewSubscriptionBpsPriceBpsConfig

// The cadence to bill for this price on.
//
// This is an alias to an internal type.
type ReplaceSubscriptionPriceParamsPriceNewSubscriptionBpsPriceCadence = shared.ReplaceSubscriptionPriceParamsPriceNewSubscriptionBpsPriceCadence

// This is an alias to an internal value.
const ReplaceSubscriptionPriceParamsPriceNewSubscriptionBpsPriceCadenceAnnual = shared.ReplaceSubscriptionPriceParamsPriceNewSubscriptionBpsPriceCadenceAnnual

// This is an alias to an internal value.
const ReplaceSubscriptionPriceParamsPriceNewSubscriptionBpsPriceCadenceSemiAnnual = shared.ReplaceSubscriptionPriceParamsPriceNewSubscriptionBpsPriceCadenceSemiAnnual

// This is an alias to an internal value.
const ReplaceSubscriptionPriceParamsPriceNewSubscriptionBpsPriceCadenceMonthly = shared.ReplaceSubscriptionPriceParamsPriceNewSubscriptionBpsPriceCadenceMonthly

// This is an alias to an internal value.
const ReplaceSubscriptionPriceParamsPriceNewSubscriptionBpsPriceCadenceQuarterly = shared.ReplaceSubscriptionPriceParamsPriceNewSubscriptionBpsPriceCadenceQuarterly

// This is an alias to an internal value.
const ReplaceSubscriptionPriceParamsPriceNewSubscriptionBpsPriceCadenceOneTime = shared.ReplaceSubscriptionPriceParamsPriceNewSubscriptionBpsPriceCadenceOneTime

// This is an alias to an internal value.
const ReplaceSubscriptionPriceParamsPriceNewSubscriptionBpsPriceCadenceCustom = shared.ReplaceSubscriptionPriceParamsPriceNewSubscriptionBpsPriceCadenceCustom

// This is an alias to an internal type.
type ReplaceSubscriptionPriceParamsPriceNewSubscriptionBpsPriceModelType = shared.ReplaceSubscriptionPriceParamsPriceNewSubscriptionBpsPriceModelType

// This is an alias to an internal value.
const ReplaceSubscriptionPriceParamsPriceNewSubscriptionBpsPriceModelTypeBps = shared.ReplaceSubscriptionPriceParamsPriceNewSubscriptionBpsPriceModelTypeBps

// For custom cadence: specifies the duration of the billing period in days or
// months.
//
// This is an alias to an internal type.
type ReplaceSubscriptionPriceParamsPriceNewSubscriptionBpsPriceBillingCycleConfiguration = shared.ReplaceSubscriptionPriceParamsPriceNewSubscriptionBpsPriceBillingCycleConfiguration

// The unit of billing period duration.
//
// This is an alias to an internal type.
type ReplaceSubscriptionPriceParamsPriceNewSubscriptionBpsPriceBillingCycleConfigurationDurationUnit = shared.ReplaceSubscriptionPriceParamsPriceNewSubscriptionBpsPriceBillingCycleConfigurationDurationUnit

// This is an alias to an internal value.
const ReplaceSubscriptionPriceParamsPriceNewSubscriptionBpsPriceBillingCycleConfigurationDurationUnitDay = shared.ReplaceSubscriptionPriceParamsPriceNewSubscriptionBpsPriceBillingCycleConfigurationDurationUnitDay

// This is an alias to an internal value.
const ReplaceSubscriptionPriceParamsPriceNewSubscriptionBpsPriceBillingCycleConfigurationDurationUnitMonth = shared.ReplaceSubscriptionPriceParamsPriceNewSubscriptionBpsPriceBillingCycleConfigurationDurationUnitMonth

// Within each billing cycle, specifies the cadence at which invoices are produced.
// If unspecified, a single invoice is produced per billing cycle.
//
// This is an alias to an internal type.
type ReplaceSubscriptionPriceParamsPriceNewSubscriptionBpsPriceInvoicingCycleConfiguration = shared.ReplaceSubscriptionPriceParamsPriceNewSubscriptionBpsPriceInvoicingCycleConfiguration

// The unit of billing period duration.
//
// This is an alias to an internal type.
type ReplaceSubscriptionPriceParamsPriceNewSubscriptionBpsPriceInvoicingCycleConfigurationDurationUnit = shared.ReplaceSubscriptionPriceParamsPriceNewSubscriptionBpsPriceInvoicingCycleConfigurationDurationUnit

// This is an alias to an internal value.
const ReplaceSubscriptionPriceParamsPriceNewSubscriptionBpsPriceInvoicingCycleConfigurationDurationUnitDay = shared.ReplaceSubscriptionPriceParamsPriceNewSubscriptionBpsPriceInvoicingCycleConfigurationDurationUnitDay

// This is an alias to an internal value.
const ReplaceSubscriptionPriceParamsPriceNewSubscriptionBpsPriceInvoicingCycleConfigurationDurationUnitMonth = shared.ReplaceSubscriptionPriceParamsPriceNewSubscriptionBpsPriceInvoicingCycleConfigurationDurationUnitMonth

// This is an alias to an internal type.
type ReplaceSubscriptionPriceParamsPriceNewSubscriptionBulkBpsPrice = shared.ReplaceSubscriptionPriceParamsPriceNewSubscriptionBulkBpsPrice

// This is an alias to an internal type.
type ReplaceSubscriptionPriceParamsPriceNewSubscriptionBulkBpsPriceBulkBpsConfig = shared.ReplaceSubscriptionPriceParamsPriceNewSubscriptionBulkBpsPriceBulkBpsConfig

// This is an alias to an internal type.
type ReplaceSubscriptionPriceParamsPriceNewSubscriptionBulkBpsPriceBulkBpsConfigTier = shared.ReplaceSubscriptionPriceParamsPriceNewSubscriptionBulkBpsPriceBulkBpsConfigTier

// The cadence to bill for this price on.
//
// This is an alias to an internal type.
type ReplaceSubscriptionPriceParamsPriceNewSubscriptionBulkBpsPriceCadence = shared.ReplaceSubscriptionPriceParamsPriceNewSubscriptionBulkBpsPriceCadence

// This is an alias to an internal value.
const ReplaceSubscriptionPriceParamsPriceNewSubscriptionBulkBpsPriceCadenceAnnual = shared.ReplaceSubscriptionPriceParamsPriceNewSubscriptionBulkBpsPriceCadenceAnnual

// This is an alias to an internal value.
const ReplaceSubscriptionPriceParamsPriceNewSubscriptionBulkBpsPriceCadenceSemiAnnual = shared.ReplaceSubscriptionPriceParamsPriceNewSubscriptionBulkBpsPriceCadenceSemiAnnual

// This is an alias to an internal value.
const ReplaceSubscriptionPriceParamsPriceNewSubscriptionBulkBpsPriceCadenceMonthly = shared.ReplaceSubscriptionPriceParamsPriceNewSubscriptionBulkBpsPriceCadenceMonthly

// This is an alias to an internal value.
const ReplaceSubscriptionPriceParamsPriceNewSubscriptionBulkBpsPriceCadenceQuarterly = shared.ReplaceSubscriptionPriceParamsPriceNewSubscriptionBulkBpsPriceCadenceQuarterly

// This is an alias to an internal value.
const ReplaceSubscriptionPriceParamsPriceNewSubscriptionBulkBpsPriceCadenceOneTime = shared.ReplaceSubscriptionPriceParamsPriceNewSubscriptionBulkBpsPriceCadenceOneTime

// This is an alias to an internal value.
const ReplaceSubscriptionPriceParamsPriceNewSubscriptionBulkBpsPriceCadenceCustom = shared.ReplaceSubscriptionPriceParamsPriceNewSubscriptionBulkBpsPriceCadenceCustom

// This is an alias to an internal type.
type ReplaceSubscriptionPriceParamsPriceNewSubscriptionBulkBpsPriceModelType = shared.ReplaceSubscriptionPriceParamsPriceNewSubscriptionBulkBpsPriceModelType

// This is an alias to an internal value.
const ReplaceSubscriptionPriceParamsPriceNewSubscriptionBulkBpsPriceModelTypeBulkBps = shared.ReplaceSubscriptionPriceParamsPriceNewSubscriptionBulkBpsPriceModelTypeBulkBps

// For custom cadence: specifies the duration of the billing period in days or
// months.
//
// This is an alias to an internal type.
type ReplaceSubscriptionPriceParamsPriceNewSubscriptionBulkBpsPriceBillingCycleConfiguration = shared.ReplaceSubscriptionPriceParamsPriceNewSubscriptionBulkBpsPriceBillingCycleConfiguration

// The unit of billing period duration.
//
// This is an alias to an internal type.
type ReplaceSubscriptionPriceParamsPriceNewSubscriptionBulkBpsPriceBillingCycleConfigurationDurationUnit = shared.ReplaceSubscriptionPriceParamsPriceNewSubscriptionBulkBpsPriceBillingCycleConfigurationDurationUnit

// This is an alias to an internal value.
const ReplaceSubscriptionPriceParamsPriceNewSubscriptionBulkBpsPriceBillingCycleConfigurationDurationUnitDay = shared.ReplaceSubscriptionPriceParamsPriceNewSubscriptionBulkBpsPriceBillingCycleConfigurationDurationUnitDay

// This is an alias to an internal value.
const ReplaceSubscriptionPriceParamsPriceNewSubscriptionBulkBpsPriceBillingCycleConfigurationDurationUnitMonth = shared.ReplaceSubscriptionPriceParamsPriceNewSubscriptionBulkBpsPriceBillingCycleConfigurationDurationUnitMonth

// Within each billing cycle, specifies the cadence at which invoices are produced.
// If unspecified, a single invoice is produced per billing cycle.
//
// This is an alias to an internal type.
type ReplaceSubscriptionPriceParamsPriceNewSubscriptionBulkBpsPriceInvoicingCycleConfiguration = shared.ReplaceSubscriptionPriceParamsPriceNewSubscriptionBulkBpsPriceInvoicingCycleConfiguration

// The unit of billing period duration.
//
// This is an alias to an internal type.
type ReplaceSubscriptionPriceParamsPriceNewSubscriptionBulkBpsPriceInvoicingCycleConfigurationDurationUnit = shared.ReplaceSubscriptionPriceParamsPriceNewSubscriptionBulkBpsPriceInvoicingCycleConfigurationDurationUnit

// This is an alias to an internal value.
const ReplaceSubscriptionPriceParamsPriceNewSubscriptionBulkBpsPriceInvoicingCycleConfigurationDurationUnitDay = shared.ReplaceSubscriptionPriceParamsPriceNewSubscriptionBulkBpsPriceInvoicingCycleConfigurationDurationUnitDay

// This is an alias to an internal value.
const ReplaceSubscriptionPriceParamsPriceNewSubscriptionBulkBpsPriceInvoicingCycleConfigurationDurationUnitMonth = shared.ReplaceSubscriptionPriceParamsPriceNewSubscriptionBulkBpsPriceInvoicingCycleConfigurationDurationUnitMonth

// This is an alias to an internal type.
type ReplaceSubscriptionPriceParamsPriceNewSubscriptionBulkPrice = shared.ReplaceSubscriptionPriceParamsPriceNewSubscriptionBulkPrice

// This is an alias to an internal type.
type ReplaceSubscriptionPriceParamsPriceNewSubscriptionBulkPriceBulkConfig = shared.ReplaceSubscriptionPriceParamsPriceNewSubscriptionBulkPriceBulkConfig

// This is an alias to an internal type.
type ReplaceSubscriptionPriceParamsPriceNewSubscriptionBulkPriceBulkConfigTier = shared.ReplaceSubscriptionPriceParamsPriceNewSubscriptionBulkPriceBulkConfigTier

// The cadence to bill for this price on.
//
// This is an alias to an internal type.
type ReplaceSubscriptionPriceParamsPriceNewSubscriptionBulkPriceCadence = shared.ReplaceSubscriptionPriceParamsPriceNewSubscriptionBulkPriceCadence

// This is an alias to an internal value.
const ReplaceSubscriptionPriceParamsPriceNewSubscriptionBulkPriceCadenceAnnual = shared.ReplaceSubscriptionPriceParamsPriceNewSubscriptionBulkPriceCadenceAnnual

// This is an alias to an internal value.
const ReplaceSubscriptionPriceParamsPriceNewSubscriptionBulkPriceCadenceSemiAnnual = shared.ReplaceSubscriptionPriceParamsPriceNewSubscriptionBulkPriceCadenceSemiAnnual

// This is an alias to an internal value.
const ReplaceSubscriptionPriceParamsPriceNewSubscriptionBulkPriceCadenceMonthly = shared.ReplaceSubscriptionPriceParamsPriceNewSubscriptionBulkPriceCadenceMonthly

// This is an alias to an internal value.
const ReplaceSubscriptionPriceParamsPriceNewSubscriptionBulkPriceCadenceQuarterly = shared.ReplaceSubscriptionPriceParamsPriceNewSubscriptionBulkPriceCadenceQuarterly

// This is an alias to an internal value.
const ReplaceSubscriptionPriceParamsPriceNewSubscriptionBulkPriceCadenceOneTime = shared.ReplaceSubscriptionPriceParamsPriceNewSubscriptionBulkPriceCadenceOneTime

// This is an alias to an internal value.
const ReplaceSubscriptionPriceParamsPriceNewSubscriptionBulkPriceCadenceCustom = shared.ReplaceSubscriptionPriceParamsPriceNewSubscriptionBulkPriceCadenceCustom

// This is an alias to an internal type.
type ReplaceSubscriptionPriceParamsPriceNewSubscriptionBulkPriceModelType = shared.ReplaceSubscriptionPriceParamsPriceNewSubscriptionBulkPriceModelType

// This is an alias to an internal value.
const ReplaceSubscriptionPriceParamsPriceNewSubscriptionBulkPriceModelTypeBulk = shared.ReplaceSubscriptionPriceParamsPriceNewSubscriptionBulkPriceModelTypeBulk

// For custom cadence: specifies the duration of the billing period in days or
// months.
//
// This is an alias to an internal type.
type ReplaceSubscriptionPriceParamsPriceNewSubscriptionBulkPriceBillingCycleConfiguration = shared.ReplaceSubscriptionPriceParamsPriceNewSubscriptionBulkPriceBillingCycleConfiguration

// The unit of billing period duration.
//
// This is an alias to an internal type.
type ReplaceSubscriptionPriceParamsPriceNewSubscriptionBulkPriceBillingCycleConfigurationDurationUnit = shared.ReplaceSubscriptionPriceParamsPriceNewSubscriptionBulkPriceBillingCycleConfigurationDurationUnit

// This is an alias to an internal value.
const ReplaceSubscriptionPriceParamsPriceNewSubscriptionBulkPriceBillingCycleConfigurationDurationUnitDay = shared.ReplaceSubscriptionPriceParamsPriceNewSubscriptionBulkPriceBillingCycleConfigurationDurationUnitDay

// This is an alias to an internal value.
const ReplaceSubscriptionPriceParamsPriceNewSubscriptionBulkPriceBillingCycleConfigurationDurationUnitMonth = shared.ReplaceSubscriptionPriceParamsPriceNewSubscriptionBulkPriceBillingCycleConfigurationDurationUnitMonth

// Within each billing cycle, specifies the cadence at which invoices are produced.
// If unspecified, a single invoice is produced per billing cycle.
//
// This is an alias to an internal type.
type ReplaceSubscriptionPriceParamsPriceNewSubscriptionBulkPriceInvoicingCycleConfiguration = shared.ReplaceSubscriptionPriceParamsPriceNewSubscriptionBulkPriceInvoicingCycleConfiguration

// The unit of billing period duration.
//
// This is an alias to an internal type.
type ReplaceSubscriptionPriceParamsPriceNewSubscriptionBulkPriceInvoicingCycleConfigurationDurationUnit = shared.ReplaceSubscriptionPriceParamsPriceNewSubscriptionBulkPriceInvoicingCycleConfigurationDurationUnit

// This is an alias to an internal value.
const ReplaceSubscriptionPriceParamsPriceNewSubscriptionBulkPriceInvoicingCycleConfigurationDurationUnitDay = shared.ReplaceSubscriptionPriceParamsPriceNewSubscriptionBulkPriceInvoicingCycleConfigurationDurationUnitDay

// This is an alias to an internal value.
const ReplaceSubscriptionPriceParamsPriceNewSubscriptionBulkPriceInvoicingCycleConfigurationDurationUnitMonth = shared.ReplaceSubscriptionPriceParamsPriceNewSubscriptionBulkPriceInvoicingCycleConfigurationDurationUnitMonth

// This is an alias to an internal type.
type ReplaceSubscriptionPriceParamsPriceNewSubscriptionThresholdTotalAmountPrice = shared.ReplaceSubscriptionPriceParamsPriceNewSubscriptionThresholdTotalAmountPrice

// The cadence to bill for this price on.
//
// This is an alias to an internal type.
type ReplaceSubscriptionPriceParamsPriceNewSubscriptionThresholdTotalAmountPriceCadence = shared.ReplaceSubscriptionPriceParamsPriceNewSubscriptionThresholdTotalAmountPriceCadence

// This is an alias to an internal value.
const ReplaceSubscriptionPriceParamsPriceNewSubscriptionThresholdTotalAmountPriceCadenceAnnual = shared.ReplaceSubscriptionPriceParamsPriceNewSubscriptionThresholdTotalAmountPriceCadenceAnnual

// This is an alias to an internal value.
const ReplaceSubscriptionPriceParamsPriceNewSubscriptionThresholdTotalAmountPriceCadenceSemiAnnual = shared.ReplaceSubscriptionPriceParamsPriceNewSubscriptionThresholdTotalAmountPriceCadenceSemiAnnual

// This is an alias to an internal value.
const ReplaceSubscriptionPriceParamsPriceNewSubscriptionThresholdTotalAmountPriceCadenceMonthly = shared.ReplaceSubscriptionPriceParamsPriceNewSubscriptionThresholdTotalAmountPriceCadenceMonthly

// This is an alias to an internal value.
const ReplaceSubscriptionPriceParamsPriceNewSubscriptionThresholdTotalAmountPriceCadenceQuarterly = shared.ReplaceSubscriptionPriceParamsPriceNewSubscriptionThresholdTotalAmountPriceCadenceQuarterly

// This is an alias to an internal value.
const ReplaceSubscriptionPriceParamsPriceNewSubscriptionThresholdTotalAmountPriceCadenceOneTime = shared.ReplaceSubscriptionPriceParamsPriceNewSubscriptionThresholdTotalAmountPriceCadenceOneTime

// This is an alias to an internal value.
const ReplaceSubscriptionPriceParamsPriceNewSubscriptionThresholdTotalAmountPriceCadenceCustom = shared.ReplaceSubscriptionPriceParamsPriceNewSubscriptionThresholdTotalAmountPriceCadenceCustom

// This is an alias to an internal type.
type ReplaceSubscriptionPriceParamsPriceNewSubscriptionThresholdTotalAmountPriceModelType = shared.ReplaceSubscriptionPriceParamsPriceNewSubscriptionThresholdTotalAmountPriceModelType

// This is an alias to an internal value.
const ReplaceSubscriptionPriceParamsPriceNewSubscriptionThresholdTotalAmountPriceModelTypeThresholdTotalAmount = shared.ReplaceSubscriptionPriceParamsPriceNewSubscriptionThresholdTotalAmountPriceModelTypeThresholdTotalAmount

// For custom cadence: specifies the duration of the billing period in days or
// months.
//
// This is an alias to an internal type.
type ReplaceSubscriptionPriceParamsPriceNewSubscriptionThresholdTotalAmountPriceBillingCycleConfiguration = shared.ReplaceSubscriptionPriceParamsPriceNewSubscriptionThresholdTotalAmountPriceBillingCycleConfiguration

// The unit of billing period duration.
//
// This is an alias to an internal type.
type ReplaceSubscriptionPriceParamsPriceNewSubscriptionThresholdTotalAmountPriceBillingCycleConfigurationDurationUnit = shared.ReplaceSubscriptionPriceParamsPriceNewSubscriptionThresholdTotalAmountPriceBillingCycleConfigurationDurationUnit

// This is an alias to an internal value.
const ReplaceSubscriptionPriceParamsPriceNewSubscriptionThresholdTotalAmountPriceBillingCycleConfigurationDurationUnitDay = shared.ReplaceSubscriptionPriceParamsPriceNewSubscriptionThresholdTotalAmountPriceBillingCycleConfigurationDurationUnitDay

// This is an alias to an internal value.
const ReplaceSubscriptionPriceParamsPriceNewSubscriptionThresholdTotalAmountPriceBillingCycleConfigurationDurationUnitMonth = shared.ReplaceSubscriptionPriceParamsPriceNewSubscriptionThresholdTotalAmountPriceBillingCycleConfigurationDurationUnitMonth

// Within each billing cycle, specifies the cadence at which invoices are produced.
// If unspecified, a single invoice is produced per billing cycle.
//
// This is an alias to an internal type.
type ReplaceSubscriptionPriceParamsPriceNewSubscriptionThresholdTotalAmountPriceInvoicingCycleConfiguration = shared.ReplaceSubscriptionPriceParamsPriceNewSubscriptionThresholdTotalAmountPriceInvoicingCycleConfiguration

// The unit of billing period duration.
//
// This is an alias to an internal type.
type ReplaceSubscriptionPriceParamsPriceNewSubscriptionThresholdTotalAmountPriceInvoicingCycleConfigurationDurationUnit = shared.ReplaceSubscriptionPriceParamsPriceNewSubscriptionThresholdTotalAmountPriceInvoicingCycleConfigurationDurationUnit

// This is an alias to an internal value.
const ReplaceSubscriptionPriceParamsPriceNewSubscriptionThresholdTotalAmountPriceInvoicingCycleConfigurationDurationUnitDay = shared.ReplaceSubscriptionPriceParamsPriceNewSubscriptionThresholdTotalAmountPriceInvoicingCycleConfigurationDurationUnitDay

// This is an alias to an internal value.
const ReplaceSubscriptionPriceParamsPriceNewSubscriptionThresholdTotalAmountPriceInvoicingCycleConfigurationDurationUnitMonth = shared.ReplaceSubscriptionPriceParamsPriceNewSubscriptionThresholdTotalAmountPriceInvoicingCycleConfigurationDurationUnitMonth

// This is an alias to an internal type.
type ReplaceSubscriptionPriceParamsPriceNewSubscriptionTieredPackagePrice = shared.ReplaceSubscriptionPriceParamsPriceNewSubscriptionTieredPackagePrice

// The cadence to bill for this price on.
//
// This is an alias to an internal type.
type ReplaceSubscriptionPriceParamsPriceNewSubscriptionTieredPackagePriceCadence = shared.ReplaceSubscriptionPriceParamsPriceNewSubscriptionTieredPackagePriceCadence

// This is an alias to an internal value.
const ReplaceSubscriptionPriceParamsPriceNewSubscriptionTieredPackagePriceCadenceAnnual = shared.ReplaceSubscriptionPriceParamsPriceNewSubscriptionTieredPackagePriceCadenceAnnual

// This is an alias to an internal value.
const ReplaceSubscriptionPriceParamsPriceNewSubscriptionTieredPackagePriceCadenceSemiAnnual = shared.ReplaceSubscriptionPriceParamsPriceNewSubscriptionTieredPackagePriceCadenceSemiAnnual

// This is an alias to an internal value.
const ReplaceSubscriptionPriceParamsPriceNewSubscriptionTieredPackagePriceCadenceMonthly = shared.ReplaceSubscriptionPriceParamsPriceNewSubscriptionTieredPackagePriceCadenceMonthly

// This is an alias to an internal value.
const ReplaceSubscriptionPriceParamsPriceNewSubscriptionTieredPackagePriceCadenceQuarterly = shared.ReplaceSubscriptionPriceParamsPriceNewSubscriptionTieredPackagePriceCadenceQuarterly

// This is an alias to an internal value.
const ReplaceSubscriptionPriceParamsPriceNewSubscriptionTieredPackagePriceCadenceOneTime = shared.ReplaceSubscriptionPriceParamsPriceNewSubscriptionTieredPackagePriceCadenceOneTime

// This is an alias to an internal value.
const ReplaceSubscriptionPriceParamsPriceNewSubscriptionTieredPackagePriceCadenceCustom = shared.ReplaceSubscriptionPriceParamsPriceNewSubscriptionTieredPackagePriceCadenceCustom

// This is an alias to an internal type.
type ReplaceSubscriptionPriceParamsPriceNewSubscriptionTieredPackagePriceModelType = shared.ReplaceSubscriptionPriceParamsPriceNewSubscriptionTieredPackagePriceModelType

// This is an alias to an internal value.
const ReplaceSubscriptionPriceParamsPriceNewSubscriptionTieredPackagePriceModelTypeTieredPackage = shared.ReplaceSubscriptionPriceParamsPriceNewSubscriptionTieredPackagePriceModelTypeTieredPackage

// For custom cadence: specifies the duration of the billing period in days or
// months.
//
// This is an alias to an internal type.
type ReplaceSubscriptionPriceParamsPriceNewSubscriptionTieredPackagePriceBillingCycleConfiguration = shared.ReplaceSubscriptionPriceParamsPriceNewSubscriptionTieredPackagePriceBillingCycleConfiguration

// The unit of billing period duration.
//
// This is an alias to an internal type.
type ReplaceSubscriptionPriceParamsPriceNewSubscriptionTieredPackagePriceBillingCycleConfigurationDurationUnit = shared.ReplaceSubscriptionPriceParamsPriceNewSubscriptionTieredPackagePriceBillingCycleConfigurationDurationUnit

// This is an alias to an internal value.
const ReplaceSubscriptionPriceParamsPriceNewSubscriptionTieredPackagePriceBillingCycleConfigurationDurationUnitDay = shared.ReplaceSubscriptionPriceParamsPriceNewSubscriptionTieredPackagePriceBillingCycleConfigurationDurationUnitDay

// This is an alias to an internal value.
const ReplaceSubscriptionPriceParamsPriceNewSubscriptionTieredPackagePriceBillingCycleConfigurationDurationUnitMonth = shared.ReplaceSubscriptionPriceParamsPriceNewSubscriptionTieredPackagePriceBillingCycleConfigurationDurationUnitMonth

// Within each billing cycle, specifies the cadence at which invoices are produced.
// If unspecified, a single invoice is produced per billing cycle.
//
// This is an alias to an internal type.
type ReplaceSubscriptionPriceParamsPriceNewSubscriptionTieredPackagePriceInvoicingCycleConfiguration = shared.ReplaceSubscriptionPriceParamsPriceNewSubscriptionTieredPackagePriceInvoicingCycleConfiguration

// The unit of billing period duration.
//
// This is an alias to an internal type.
type ReplaceSubscriptionPriceParamsPriceNewSubscriptionTieredPackagePriceInvoicingCycleConfigurationDurationUnit = shared.ReplaceSubscriptionPriceParamsPriceNewSubscriptionTieredPackagePriceInvoicingCycleConfigurationDurationUnit

// This is an alias to an internal value.
const ReplaceSubscriptionPriceParamsPriceNewSubscriptionTieredPackagePriceInvoicingCycleConfigurationDurationUnitDay = shared.ReplaceSubscriptionPriceParamsPriceNewSubscriptionTieredPackagePriceInvoicingCycleConfigurationDurationUnitDay

// This is an alias to an internal value.
const ReplaceSubscriptionPriceParamsPriceNewSubscriptionTieredPackagePriceInvoicingCycleConfigurationDurationUnitMonth = shared.ReplaceSubscriptionPriceParamsPriceNewSubscriptionTieredPackagePriceInvoicingCycleConfigurationDurationUnitMonth

// This is an alias to an internal type.
type ReplaceSubscriptionPriceParamsPriceNewSubscriptionTieredWithMinimumPrice = shared.ReplaceSubscriptionPriceParamsPriceNewSubscriptionTieredWithMinimumPrice

// The cadence to bill for this price on.
//
// This is an alias to an internal type.
type ReplaceSubscriptionPriceParamsPriceNewSubscriptionTieredWithMinimumPriceCadence = shared.ReplaceSubscriptionPriceParamsPriceNewSubscriptionTieredWithMinimumPriceCadence

// This is an alias to an internal value.
const ReplaceSubscriptionPriceParamsPriceNewSubscriptionTieredWithMinimumPriceCadenceAnnual = shared.ReplaceSubscriptionPriceParamsPriceNewSubscriptionTieredWithMinimumPriceCadenceAnnual

// This is an alias to an internal value.
const ReplaceSubscriptionPriceParamsPriceNewSubscriptionTieredWithMinimumPriceCadenceSemiAnnual = shared.ReplaceSubscriptionPriceParamsPriceNewSubscriptionTieredWithMinimumPriceCadenceSemiAnnual

// This is an alias to an internal value.
const ReplaceSubscriptionPriceParamsPriceNewSubscriptionTieredWithMinimumPriceCadenceMonthly = shared.ReplaceSubscriptionPriceParamsPriceNewSubscriptionTieredWithMinimumPriceCadenceMonthly

// This is an alias to an internal value.
const ReplaceSubscriptionPriceParamsPriceNewSubscriptionTieredWithMinimumPriceCadenceQuarterly = shared.ReplaceSubscriptionPriceParamsPriceNewSubscriptionTieredWithMinimumPriceCadenceQuarterly

// This is an alias to an internal value.
const ReplaceSubscriptionPriceParamsPriceNewSubscriptionTieredWithMinimumPriceCadenceOneTime = shared.ReplaceSubscriptionPriceParamsPriceNewSubscriptionTieredWithMinimumPriceCadenceOneTime

// This is an alias to an internal value.
const ReplaceSubscriptionPriceParamsPriceNewSubscriptionTieredWithMinimumPriceCadenceCustom = shared.ReplaceSubscriptionPriceParamsPriceNewSubscriptionTieredWithMinimumPriceCadenceCustom

// This is an alias to an internal type.
type ReplaceSubscriptionPriceParamsPriceNewSubscriptionTieredWithMinimumPriceModelType = shared.ReplaceSubscriptionPriceParamsPriceNewSubscriptionTieredWithMinimumPriceModelType

// This is an alias to an internal value.
const ReplaceSubscriptionPriceParamsPriceNewSubscriptionTieredWithMinimumPriceModelTypeTieredWithMinimum = shared.ReplaceSubscriptionPriceParamsPriceNewSubscriptionTieredWithMinimumPriceModelTypeTieredWithMinimum

// For custom cadence: specifies the duration of the billing period in days or
// months.
//
// This is an alias to an internal type.
type ReplaceSubscriptionPriceParamsPriceNewSubscriptionTieredWithMinimumPriceBillingCycleConfiguration = shared.ReplaceSubscriptionPriceParamsPriceNewSubscriptionTieredWithMinimumPriceBillingCycleConfiguration

// The unit of billing period duration.
//
// This is an alias to an internal type.
type ReplaceSubscriptionPriceParamsPriceNewSubscriptionTieredWithMinimumPriceBillingCycleConfigurationDurationUnit = shared.ReplaceSubscriptionPriceParamsPriceNewSubscriptionTieredWithMinimumPriceBillingCycleConfigurationDurationUnit

// This is an alias to an internal value.
const ReplaceSubscriptionPriceParamsPriceNewSubscriptionTieredWithMinimumPriceBillingCycleConfigurationDurationUnitDay = shared.ReplaceSubscriptionPriceParamsPriceNewSubscriptionTieredWithMinimumPriceBillingCycleConfigurationDurationUnitDay

// This is an alias to an internal value.
const ReplaceSubscriptionPriceParamsPriceNewSubscriptionTieredWithMinimumPriceBillingCycleConfigurationDurationUnitMonth = shared.ReplaceSubscriptionPriceParamsPriceNewSubscriptionTieredWithMinimumPriceBillingCycleConfigurationDurationUnitMonth

// Within each billing cycle, specifies the cadence at which invoices are produced.
// If unspecified, a single invoice is produced per billing cycle.
//
// This is an alias to an internal type.
type ReplaceSubscriptionPriceParamsPriceNewSubscriptionTieredWithMinimumPriceInvoicingCycleConfiguration = shared.ReplaceSubscriptionPriceParamsPriceNewSubscriptionTieredWithMinimumPriceInvoicingCycleConfiguration

// The unit of billing period duration.
//
// This is an alias to an internal type.
type ReplaceSubscriptionPriceParamsPriceNewSubscriptionTieredWithMinimumPriceInvoicingCycleConfigurationDurationUnit = shared.ReplaceSubscriptionPriceParamsPriceNewSubscriptionTieredWithMinimumPriceInvoicingCycleConfigurationDurationUnit

// This is an alias to an internal value.
const ReplaceSubscriptionPriceParamsPriceNewSubscriptionTieredWithMinimumPriceInvoicingCycleConfigurationDurationUnitDay = shared.ReplaceSubscriptionPriceParamsPriceNewSubscriptionTieredWithMinimumPriceInvoicingCycleConfigurationDurationUnitDay

// This is an alias to an internal value.
const ReplaceSubscriptionPriceParamsPriceNewSubscriptionTieredWithMinimumPriceInvoicingCycleConfigurationDurationUnitMonth = shared.ReplaceSubscriptionPriceParamsPriceNewSubscriptionTieredWithMinimumPriceInvoicingCycleConfigurationDurationUnitMonth

// This is an alias to an internal type.
type ReplaceSubscriptionPriceParamsPriceNewSubscriptionUnitWithPercentPrice = shared.ReplaceSubscriptionPriceParamsPriceNewSubscriptionUnitWithPercentPrice

// The cadence to bill for this price on.
//
// This is an alias to an internal type.
type ReplaceSubscriptionPriceParamsPriceNewSubscriptionUnitWithPercentPriceCadence = shared.ReplaceSubscriptionPriceParamsPriceNewSubscriptionUnitWithPercentPriceCadence

// This is an alias to an internal value.
const ReplaceSubscriptionPriceParamsPriceNewSubscriptionUnitWithPercentPriceCadenceAnnual = shared.ReplaceSubscriptionPriceParamsPriceNewSubscriptionUnitWithPercentPriceCadenceAnnual

// This is an alias to an internal value.
const ReplaceSubscriptionPriceParamsPriceNewSubscriptionUnitWithPercentPriceCadenceSemiAnnual = shared.ReplaceSubscriptionPriceParamsPriceNewSubscriptionUnitWithPercentPriceCadenceSemiAnnual

// This is an alias to an internal value.
const ReplaceSubscriptionPriceParamsPriceNewSubscriptionUnitWithPercentPriceCadenceMonthly = shared.ReplaceSubscriptionPriceParamsPriceNewSubscriptionUnitWithPercentPriceCadenceMonthly

// This is an alias to an internal value.
const ReplaceSubscriptionPriceParamsPriceNewSubscriptionUnitWithPercentPriceCadenceQuarterly = shared.ReplaceSubscriptionPriceParamsPriceNewSubscriptionUnitWithPercentPriceCadenceQuarterly

// This is an alias to an internal value.
const ReplaceSubscriptionPriceParamsPriceNewSubscriptionUnitWithPercentPriceCadenceOneTime = shared.ReplaceSubscriptionPriceParamsPriceNewSubscriptionUnitWithPercentPriceCadenceOneTime

// This is an alias to an internal value.
const ReplaceSubscriptionPriceParamsPriceNewSubscriptionUnitWithPercentPriceCadenceCustom = shared.ReplaceSubscriptionPriceParamsPriceNewSubscriptionUnitWithPercentPriceCadenceCustom

// This is an alias to an internal type.
type ReplaceSubscriptionPriceParamsPriceNewSubscriptionUnitWithPercentPriceModelType = shared.ReplaceSubscriptionPriceParamsPriceNewSubscriptionUnitWithPercentPriceModelType

// This is an alias to an internal value.
const ReplaceSubscriptionPriceParamsPriceNewSubscriptionUnitWithPercentPriceModelTypeUnitWithPercent = shared.ReplaceSubscriptionPriceParamsPriceNewSubscriptionUnitWithPercentPriceModelTypeUnitWithPercent

// For custom cadence: specifies the duration of the billing period in days or
// months.
//
// This is an alias to an internal type.
type ReplaceSubscriptionPriceParamsPriceNewSubscriptionUnitWithPercentPriceBillingCycleConfiguration = shared.ReplaceSubscriptionPriceParamsPriceNewSubscriptionUnitWithPercentPriceBillingCycleConfiguration

// The unit of billing period duration.
//
// This is an alias to an internal type.
type ReplaceSubscriptionPriceParamsPriceNewSubscriptionUnitWithPercentPriceBillingCycleConfigurationDurationUnit = shared.ReplaceSubscriptionPriceParamsPriceNewSubscriptionUnitWithPercentPriceBillingCycleConfigurationDurationUnit

// This is an alias to an internal value.
const ReplaceSubscriptionPriceParamsPriceNewSubscriptionUnitWithPercentPriceBillingCycleConfigurationDurationUnitDay = shared.ReplaceSubscriptionPriceParamsPriceNewSubscriptionUnitWithPercentPriceBillingCycleConfigurationDurationUnitDay

// This is an alias to an internal value.
const ReplaceSubscriptionPriceParamsPriceNewSubscriptionUnitWithPercentPriceBillingCycleConfigurationDurationUnitMonth = shared.ReplaceSubscriptionPriceParamsPriceNewSubscriptionUnitWithPercentPriceBillingCycleConfigurationDurationUnitMonth

// Within each billing cycle, specifies the cadence at which invoices are produced.
// If unspecified, a single invoice is produced per billing cycle.
//
// This is an alias to an internal type.
type ReplaceSubscriptionPriceParamsPriceNewSubscriptionUnitWithPercentPriceInvoicingCycleConfiguration = shared.ReplaceSubscriptionPriceParamsPriceNewSubscriptionUnitWithPercentPriceInvoicingCycleConfiguration

// The unit of billing period duration.
//
// This is an alias to an internal type.
type ReplaceSubscriptionPriceParamsPriceNewSubscriptionUnitWithPercentPriceInvoicingCycleConfigurationDurationUnit = shared.ReplaceSubscriptionPriceParamsPriceNewSubscriptionUnitWithPercentPriceInvoicingCycleConfigurationDurationUnit

// This is an alias to an internal value.
const ReplaceSubscriptionPriceParamsPriceNewSubscriptionUnitWithPercentPriceInvoicingCycleConfigurationDurationUnitDay = shared.ReplaceSubscriptionPriceParamsPriceNewSubscriptionUnitWithPercentPriceInvoicingCycleConfigurationDurationUnitDay

// This is an alias to an internal value.
const ReplaceSubscriptionPriceParamsPriceNewSubscriptionUnitWithPercentPriceInvoicingCycleConfigurationDurationUnitMonth = shared.ReplaceSubscriptionPriceParamsPriceNewSubscriptionUnitWithPercentPriceInvoicingCycleConfigurationDurationUnitMonth

// This is an alias to an internal type.
type ReplaceSubscriptionPriceParamsPriceNewSubscriptionPackageWithAllocationPrice = shared.ReplaceSubscriptionPriceParamsPriceNewSubscriptionPackageWithAllocationPrice

// The cadence to bill for this price on.
//
// This is an alias to an internal type.
type ReplaceSubscriptionPriceParamsPriceNewSubscriptionPackageWithAllocationPriceCadence = shared.ReplaceSubscriptionPriceParamsPriceNewSubscriptionPackageWithAllocationPriceCadence

// This is an alias to an internal value.
const ReplaceSubscriptionPriceParamsPriceNewSubscriptionPackageWithAllocationPriceCadenceAnnual = shared.ReplaceSubscriptionPriceParamsPriceNewSubscriptionPackageWithAllocationPriceCadenceAnnual

// This is an alias to an internal value.
const ReplaceSubscriptionPriceParamsPriceNewSubscriptionPackageWithAllocationPriceCadenceSemiAnnual = shared.ReplaceSubscriptionPriceParamsPriceNewSubscriptionPackageWithAllocationPriceCadenceSemiAnnual

// This is an alias to an internal value.
const ReplaceSubscriptionPriceParamsPriceNewSubscriptionPackageWithAllocationPriceCadenceMonthly = shared.ReplaceSubscriptionPriceParamsPriceNewSubscriptionPackageWithAllocationPriceCadenceMonthly

// This is an alias to an internal value.
const ReplaceSubscriptionPriceParamsPriceNewSubscriptionPackageWithAllocationPriceCadenceQuarterly = shared.ReplaceSubscriptionPriceParamsPriceNewSubscriptionPackageWithAllocationPriceCadenceQuarterly

// This is an alias to an internal value.
const ReplaceSubscriptionPriceParamsPriceNewSubscriptionPackageWithAllocationPriceCadenceOneTime = shared.ReplaceSubscriptionPriceParamsPriceNewSubscriptionPackageWithAllocationPriceCadenceOneTime

// This is an alias to an internal value.
const ReplaceSubscriptionPriceParamsPriceNewSubscriptionPackageWithAllocationPriceCadenceCustom = shared.ReplaceSubscriptionPriceParamsPriceNewSubscriptionPackageWithAllocationPriceCadenceCustom

// This is an alias to an internal type.
type ReplaceSubscriptionPriceParamsPriceNewSubscriptionPackageWithAllocationPriceModelType = shared.ReplaceSubscriptionPriceParamsPriceNewSubscriptionPackageWithAllocationPriceModelType

// This is an alias to an internal value.
const ReplaceSubscriptionPriceParamsPriceNewSubscriptionPackageWithAllocationPriceModelTypePackageWithAllocation = shared.ReplaceSubscriptionPriceParamsPriceNewSubscriptionPackageWithAllocationPriceModelTypePackageWithAllocation

// For custom cadence: specifies the duration of the billing period in days or
// months.
//
// This is an alias to an internal type.
type ReplaceSubscriptionPriceParamsPriceNewSubscriptionPackageWithAllocationPriceBillingCycleConfiguration = shared.ReplaceSubscriptionPriceParamsPriceNewSubscriptionPackageWithAllocationPriceBillingCycleConfiguration

// The unit of billing period duration.
//
// This is an alias to an internal type.
type ReplaceSubscriptionPriceParamsPriceNewSubscriptionPackageWithAllocationPriceBillingCycleConfigurationDurationUnit = shared.ReplaceSubscriptionPriceParamsPriceNewSubscriptionPackageWithAllocationPriceBillingCycleConfigurationDurationUnit

// This is an alias to an internal value.
const ReplaceSubscriptionPriceParamsPriceNewSubscriptionPackageWithAllocationPriceBillingCycleConfigurationDurationUnitDay = shared.ReplaceSubscriptionPriceParamsPriceNewSubscriptionPackageWithAllocationPriceBillingCycleConfigurationDurationUnitDay

// This is an alias to an internal value.
const ReplaceSubscriptionPriceParamsPriceNewSubscriptionPackageWithAllocationPriceBillingCycleConfigurationDurationUnitMonth = shared.ReplaceSubscriptionPriceParamsPriceNewSubscriptionPackageWithAllocationPriceBillingCycleConfigurationDurationUnitMonth

// Within each billing cycle, specifies the cadence at which invoices are produced.
// If unspecified, a single invoice is produced per billing cycle.
//
// This is an alias to an internal type.
type ReplaceSubscriptionPriceParamsPriceNewSubscriptionPackageWithAllocationPriceInvoicingCycleConfiguration = shared.ReplaceSubscriptionPriceParamsPriceNewSubscriptionPackageWithAllocationPriceInvoicingCycleConfiguration

// The unit of billing period duration.
//
// This is an alias to an internal type.
type ReplaceSubscriptionPriceParamsPriceNewSubscriptionPackageWithAllocationPriceInvoicingCycleConfigurationDurationUnit = shared.ReplaceSubscriptionPriceParamsPriceNewSubscriptionPackageWithAllocationPriceInvoicingCycleConfigurationDurationUnit

// This is an alias to an internal value.
const ReplaceSubscriptionPriceParamsPriceNewSubscriptionPackageWithAllocationPriceInvoicingCycleConfigurationDurationUnitDay = shared.ReplaceSubscriptionPriceParamsPriceNewSubscriptionPackageWithAllocationPriceInvoicingCycleConfigurationDurationUnitDay

// This is an alias to an internal value.
const ReplaceSubscriptionPriceParamsPriceNewSubscriptionPackageWithAllocationPriceInvoicingCycleConfigurationDurationUnitMonth = shared.ReplaceSubscriptionPriceParamsPriceNewSubscriptionPackageWithAllocationPriceInvoicingCycleConfigurationDurationUnitMonth

// This is an alias to an internal type.
type ReplaceSubscriptionPriceParamsPriceNewSubscriptionTierWithProrationPrice = shared.ReplaceSubscriptionPriceParamsPriceNewSubscriptionTierWithProrationPrice

// The cadence to bill for this price on.
//
// This is an alias to an internal type.
type ReplaceSubscriptionPriceParamsPriceNewSubscriptionTierWithProrationPriceCadence = shared.ReplaceSubscriptionPriceParamsPriceNewSubscriptionTierWithProrationPriceCadence

// This is an alias to an internal value.
const ReplaceSubscriptionPriceParamsPriceNewSubscriptionTierWithProrationPriceCadenceAnnual = shared.ReplaceSubscriptionPriceParamsPriceNewSubscriptionTierWithProrationPriceCadenceAnnual

// This is an alias to an internal value.
const ReplaceSubscriptionPriceParamsPriceNewSubscriptionTierWithProrationPriceCadenceSemiAnnual = shared.ReplaceSubscriptionPriceParamsPriceNewSubscriptionTierWithProrationPriceCadenceSemiAnnual

// This is an alias to an internal value.
const ReplaceSubscriptionPriceParamsPriceNewSubscriptionTierWithProrationPriceCadenceMonthly = shared.ReplaceSubscriptionPriceParamsPriceNewSubscriptionTierWithProrationPriceCadenceMonthly

// This is an alias to an internal value.
const ReplaceSubscriptionPriceParamsPriceNewSubscriptionTierWithProrationPriceCadenceQuarterly = shared.ReplaceSubscriptionPriceParamsPriceNewSubscriptionTierWithProrationPriceCadenceQuarterly

// This is an alias to an internal value.
const ReplaceSubscriptionPriceParamsPriceNewSubscriptionTierWithProrationPriceCadenceOneTime = shared.ReplaceSubscriptionPriceParamsPriceNewSubscriptionTierWithProrationPriceCadenceOneTime

// This is an alias to an internal value.
const ReplaceSubscriptionPriceParamsPriceNewSubscriptionTierWithProrationPriceCadenceCustom = shared.ReplaceSubscriptionPriceParamsPriceNewSubscriptionTierWithProrationPriceCadenceCustom

// This is an alias to an internal type.
type ReplaceSubscriptionPriceParamsPriceNewSubscriptionTierWithProrationPriceModelType = shared.ReplaceSubscriptionPriceParamsPriceNewSubscriptionTierWithProrationPriceModelType

// This is an alias to an internal value.
const ReplaceSubscriptionPriceParamsPriceNewSubscriptionTierWithProrationPriceModelTypeTieredWithProration = shared.ReplaceSubscriptionPriceParamsPriceNewSubscriptionTierWithProrationPriceModelTypeTieredWithProration

// For custom cadence: specifies the duration of the billing period in days or
// months.
//
// This is an alias to an internal type.
type ReplaceSubscriptionPriceParamsPriceNewSubscriptionTierWithProrationPriceBillingCycleConfiguration = shared.ReplaceSubscriptionPriceParamsPriceNewSubscriptionTierWithProrationPriceBillingCycleConfiguration

// The unit of billing period duration.
//
// This is an alias to an internal type.
type ReplaceSubscriptionPriceParamsPriceNewSubscriptionTierWithProrationPriceBillingCycleConfigurationDurationUnit = shared.ReplaceSubscriptionPriceParamsPriceNewSubscriptionTierWithProrationPriceBillingCycleConfigurationDurationUnit

// This is an alias to an internal value.
const ReplaceSubscriptionPriceParamsPriceNewSubscriptionTierWithProrationPriceBillingCycleConfigurationDurationUnitDay = shared.ReplaceSubscriptionPriceParamsPriceNewSubscriptionTierWithProrationPriceBillingCycleConfigurationDurationUnitDay

// This is an alias to an internal value.
const ReplaceSubscriptionPriceParamsPriceNewSubscriptionTierWithProrationPriceBillingCycleConfigurationDurationUnitMonth = shared.ReplaceSubscriptionPriceParamsPriceNewSubscriptionTierWithProrationPriceBillingCycleConfigurationDurationUnitMonth

// Within each billing cycle, specifies the cadence at which invoices are produced.
// If unspecified, a single invoice is produced per billing cycle.
//
// This is an alias to an internal type.
type ReplaceSubscriptionPriceParamsPriceNewSubscriptionTierWithProrationPriceInvoicingCycleConfiguration = shared.ReplaceSubscriptionPriceParamsPriceNewSubscriptionTierWithProrationPriceInvoicingCycleConfiguration

// The unit of billing period duration.
//
// This is an alias to an internal type.
type ReplaceSubscriptionPriceParamsPriceNewSubscriptionTierWithProrationPriceInvoicingCycleConfigurationDurationUnit = shared.ReplaceSubscriptionPriceParamsPriceNewSubscriptionTierWithProrationPriceInvoicingCycleConfigurationDurationUnit

// This is an alias to an internal value.
const ReplaceSubscriptionPriceParamsPriceNewSubscriptionTierWithProrationPriceInvoicingCycleConfigurationDurationUnitDay = shared.ReplaceSubscriptionPriceParamsPriceNewSubscriptionTierWithProrationPriceInvoicingCycleConfigurationDurationUnitDay

// This is an alias to an internal value.
const ReplaceSubscriptionPriceParamsPriceNewSubscriptionTierWithProrationPriceInvoicingCycleConfigurationDurationUnitMonth = shared.ReplaceSubscriptionPriceParamsPriceNewSubscriptionTierWithProrationPriceInvoicingCycleConfigurationDurationUnitMonth

// This is an alias to an internal type.
type ReplaceSubscriptionPriceParamsPriceNewSubscriptionUnitWithProrationPrice = shared.ReplaceSubscriptionPriceParamsPriceNewSubscriptionUnitWithProrationPrice

// The cadence to bill for this price on.
//
// This is an alias to an internal type.
type ReplaceSubscriptionPriceParamsPriceNewSubscriptionUnitWithProrationPriceCadence = shared.ReplaceSubscriptionPriceParamsPriceNewSubscriptionUnitWithProrationPriceCadence

// This is an alias to an internal value.
const ReplaceSubscriptionPriceParamsPriceNewSubscriptionUnitWithProrationPriceCadenceAnnual = shared.ReplaceSubscriptionPriceParamsPriceNewSubscriptionUnitWithProrationPriceCadenceAnnual

// This is an alias to an internal value.
const ReplaceSubscriptionPriceParamsPriceNewSubscriptionUnitWithProrationPriceCadenceSemiAnnual = shared.ReplaceSubscriptionPriceParamsPriceNewSubscriptionUnitWithProrationPriceCadenceSemiAnnual

// This is an alias to an internal value.
const ReplaceSubscriptionPriceParamsPriceNewSubscriptionUnitWithProrationPriceCadenceMonthly = shared.ReplaceSubscriptionPriceParamsPriceNewSubscriptionUnitWithProrationPriceCadenceMonthly

// This is an alias to an internal value.
const ReplaceSubscriptionPriceParamsPriceNewSubscriptionUnitWithProrationPriceCadenceQuarterly = shared.ReplaceSubscriptionPriceParamsPriceNewSubscriptionUnitWithProrationPriceCadenceQuarterly

// This is an alias to an internal value.
const ReplaceSubscriptionPriceParamsPriceNewSubscriptionUnitWithProrationPriceCadenceOneTime = shared.ReplaceSubscriptionPriceParamsPriceNewSubscriptionUnitWithProrationPriceCadenceOneTime

// This is an alias to an internal value.
const ReplaceSubscriptionPriceParamsPriceNewSubscriptionUnitWithProrationPriceCadenceCustom = shared.ReplaceSubscriptionPriceParamsPriceNewSubscriptionUnitWithProrationPriceCadenceCustom

// This is an alias to an internal type.
type ReplaceSubscriptionPriceParamsPriceNewSubscriptionUnitWithProrationPriceModelType = shared.ReplaceSubscriptionPriceParamsPriceNewSubscriptionUnitWithProrationPriceModelType

// This is an alias to an internal value.
const ReplaceSubscriptionPriceParamsPriceNewSubscriptionUnitWithProrationPriceModelTypeUnitWithProration = shared.ReplaceSubscriptionPriceParamsPriceNewSubscriptionUnitWithProrationPriceModelTypeUnitWithProration

// For custom cadence: specifies the duration of the billing period in days or
// months.
//
// This is an alias to an internal type.
type ReplaceSubscriptionPriceParamsPriceNewSubscriptionUnitWithProrationPriceBillingCycleConfiguration = shared.ReplaceSubscriptionPriceParamsPriceNewSubscriptionUnitWithProrationPriceBillingCycleConfiguration

// The unit of billing period duration.
//
// This is an alias to an internal type.
type ReplaceSubscriptionPriceParamsPriceNewSubscriptionUnitWithProrationPriceBillingCycleConfigurationDurationUnit = shared.ReplaceSubscriptionPriceParamsPriceNewSubscriptionUnitWithProrationPriceBillingCycleConfigurationDurationUnit

// This is an alias to an internal value.
const ReplaceSubscriptionPriceParamsPriceNewSubscriptionUnitWithProrationPriceBillingCycleConfigurationDurationUnitDay = shared.ReplaceSubscriptionPriceParamsPriceNewSubscriptionUnitWithProrationPriceBillingCycleConfigurationDurationUnitDay

// This is an alias to an internal value.
const ReplaceSubscriptionPriceParamsPriceNewSubscriptionUnitWithProrationPriceBillingCycleConfigurationDurationUnitMonth = shared.ReplaceSubscriptionPriceParamsPriceNewSubscriptionUnitWithProrationPriceBillingCycleConfigurationDurationUnitMonth

// Within each billing cycle, specifies the cadence at which invoices are produced.
// If unspecified, a single invoice is produced per billing cycle.
//
// This is an alias to an internal type.
type ReplaceSubscriptionPriceParamsPriceNewSubscriptionUnitWithProrationPriceInvoicingCycleConfiguration = shared.ReplaceSubscriptionPriceParamsPriceNewSubscriptionUnitWithProrationPriceInvoicingCycleConfiguration

// The unit of billing period duration.
//
// This is an alias to an internal type.
type ReplaceSubscriptionPriceParamsPriceNewSubscriptionUnitWithProrationPriceInvoicingCycleConfigurationDurationUnit = shared.ReplaceSubscriptionPriceParamsPriceNewSubscriptionUnitWithProrationPriceInvoicingCycleConfigurationDurationUnit

// This is an alias to an internal value.
const ReplaceSubscriptionPriceParamsPriceNewSubscriptionUnitWithProrationPriceInvoicingCycleConfigurationDurationUnitDay = shared.ReplaceSubscriptionPriceParamsPriceNewSubscriptionUnitWithProrationPriceInvoicingCycleConfigurationDurationUnitDay

// This is an alias to an internal value.
const ReplaceSubscriptionPriceParamsPriceNewSubscriptionUnitWithProrationPriceInvoicingCycleConfigurationDurationUnitMonth = shared.ReplaceSubscriptionPriceParamsPriceNewSubscriptionUnitWithProrationPriceInvoicingCycleConfigurationDurationUnitMonth

// This is an alias to an internal type.
type ReplaceSubscriptionPriceParamsPriceNewSubscriptionGroupedAllocationPrice = shared.ReplaceSubscriptionPriceParamsPriceNewSubscriptionGroupedAllocationPrice

// The cadence to bill for this price on.
//
// This is an alias to an internal type.
type ReplaceSubscriptionPriceParamsPriceNewSubscriptionGroupedAllocationPriceCadence = shared.ReplaceSubscriptionPriceParamsPriceNewSubscriptionGroupedAllocationPriceCadence

// This is an alias to an internal value.
const ReplaceSubscriptionPriceParamsPriceNewSubscriptionGroupedAllocationPriceCadenceAnnual = shared.ReplaceSubscriptionPriceParamsPriceNewSubscriptionGroupedAllocationPriceCadenceAnnual

// This is an alias to an internal value.
const ReplaceSubscriptionPriceParamsPriceNewSubscriptionGroupedAllocationPriceCadenceSemiAnnual = shared.ReplaceSubscriptionPriceParamsPriceNewSubscriptionGroupedAllocationPriceCadenceSemiAnnual

// This is an alias to an internal value.
const ReplaceSubscriptionPriceParamsPriceNewSubscriptionGroupedAllocationPriceCadenceMonthly = shared.ReplaceSubscriptionPriceParamsPriceNewSubscriptionGroupedAllocationPriceCadenceMonthly

// This is an alias to an internal value.
const ReplaceSubscriptionPriceParamsPriceNewSubscriptionGroupedAllocationPriceCadenceQuarterly = shared.ReplaceSubscriptionPriceParamsPriceNewSubscriptionGroupedAllocationPriceCadenceQuarterly

// This is an alias to an internal value.
const ReplaceSubscriptionPriceParamsPriceNewSubscriptionGroupedAllocationPriceCadenceOneTime = shared.ReplaceSubscriptionPriceParamsPriceNewSubscriptionGroupedAllocationPriceCadenceOneTime

// This is an alias to an internal value.
const ReplaceSubscriptionPriceParamsPriceNewSubscriptionGroupedAllocationPriceCadenceCustom = shared.ReplaceSubscriptionPriceParamsPriceNewSubscriptionGroupedAllocationPriceCadenceCustom

// This is an alias to an internal type.
type ReplaceSubscriptionPriceParamsPriceNewSubscriptionGroupedAllocationPriceModelType = shared.ReplaceSubscriptionPriceParamsPriceNewSubscriptionGroupedAllocationPriceModelType

// This is an alias to an internal value.
const ReplaceSubscriptionPriceParamsPriceNewSubscriptionGroupedAllocationPriceModelTypeGroupedAllocation = shared.ReplaceSubscriptionPriceParamsPriceNewSubscriptionGroupedAllocationPriceModelTypeGroupedAllocation

// For custom cadence: specifies the duration of the billing period in days or
// months.
//
// This is an alias to an internal type.
type ReplaceSubscriptionPriceParamsPriceNewSubscriptionGroupedAllocationPriceBillingCycleConfiguration = shared.ReplaceSubscriptionPriceParamsPriceNewSubscriptionGroupedAllocationPriceBillingCycleConfiguration

// The unit of billing period duration.
//
// This is an alias to an internal type.
type ReplaceSubscriptionPriceParamsPriceNewSubscriptionGroupedAllocationPriceBillingCycleConfigurationDurationUnit = shared.ReplaceSubscriptionPriceParamsPriceNewSubscriptionGroupedAllocationPriceBillingCycleConfigurationDurationUnit

// This is an alias to an internal value.
const ReplaceSubscriptionPriceParamsPriceNewSubscriptionGroupedAllocationPriceBillingCycleConfigurationDurationUnitDay = shared.ReplaceSubscriptionPriceParamsPriceNewSubscriptionGroupedAllocationPriceBillingCycleConfigurationDurationUnitDay

// This is an alias to an internal value.
const ReplaceSubscriptionPriceParamsPriceNewSubscriptionGroupedAllocationPriceBillingCycleConfigurationDurationUnitMonth = shared.ReplaceSubscriptionPriceParamsPriceNewSubscriptionGroupedAllocationPriceBillingCycleConfigurationDurationUnitMonth

// Within each billing cycle, specifies the cadence at which invoices are produced.
// If unspecified, a single invoice is produced per billing cycle.
//
// This is an alias to an internal type.
type ReplaceSubscriptionPriceParamsPriceNewSubscriptionGroupedAllocationPriceInvoicingCycleConfiguration = shared.ReplaceSubscriptionPriceParamsPriceNewSubscriptionGroupedAllocationPriceInvoicingCycleConfiguration

// The unit of billing period duration.
//
// This is an alias to an internal type.
type ReplaceSubscriptionPriceParamsPriceNewSubscriptionGroupedAllocationPriceInvoicingCycleConfigurationDurationUnit = shared.ReplaceSubscriptionPriceParamsPriceNewSubscriptionGroupedAllocationPriceInvoicingCycleConfigurationDurationUnit

// This is an alias to an internal value.
const ReplaceSubscriptionPriceParamsPriceNewSubscriptionGroupedAllocationPriceInvoicingCycleConfigurationDurationUnitDay = shared.ReplaceSubscriptionPriceParamsPriceNewSubscriptionGroupedAllocationPriceInvoicingCycleConfigurationDurationUnitDay

// This is an alias to an internal value.
const ReplaceSubscriptionPriceParamsPriceNewSubscriptionGroupedAllocationPriceInvoicingCycleConfigurationDurationUnitMonth = shared.ReplaceSubscriptionPriceParamsPriceNewSubscriptionGroupedAllocationPriceInvoicingCycleConfigurationDurationUnitMonth

// This is an alias to an internal type.
type ReplaceSubscriptionPriceParamsPriceNewSubscriptionGroupedWithProratedMinimumPrice = shared.ReplaceSubscriptionPriceParamsPriceNewSubscriptionGroupedWithProratedMinimumPrice

// The cadence to bill for this price on.
//
// This is an alias to an internal type.
type ReplaceSubscriptionPriceParamsPriceNewSubscriptionGroupedWithProratedMinimumPriceCadence = shared.ReplaceSubscriptionPriceParamsPriceNewSubscriptionGroupedWithProratedMinimumPriceCadence

// This is an alias to an internal value.
const ReplaceSubscriptionPriceParamsPriceNewSubscriptionGroupedWithProratedMinimumPriceCadenceAnnual = shared.ReplaceSubscriptionPriceParamsPriceNewSubscriptionGroupedWithProratedMinimumPriceCadenceAnnual

// This is an alias to an internal value.
const ReplaceSubscriptionPriceParamsPriceNewSubscriptionGroupedWithProratedMinimumPriceCadenceSemiAnnual = shared.ReplaceSubscriptionPriceParamsPriceNewSubscriptionGroupedWithProratedMinimumPriceCadenceSemiAnnual

// This is an alias to an internal value.
const ReplaceSubscriptionPriceParamsPriceNewSubscriptionGroupedWithProratedMinimumPriceCadenceMonthly = shared.ReplaceSubscriptionPriceParamsPriceNewSubscriptionGroupedWithProratedMinimumPriceCadenceMonthly

// This is an alias to an internal value.
const ReplaceSubscriptionPriceParamsPriceNewSubscriptionGroupedWithProratedMinimumPriceCadenceQuarterly = shared.ReplaceSubscriptionPriceParamsPriceNewSubscriptionGroupedWithProratedMinimumPriceCadenceQuarterly

// This is an alias to an internal value.
const ReplaceSubscriptionPriceParamsPriceNewSubscriptionGroupedWithProratedMinimumPriceCadenceOneTime = shared.ReplaceSubscriptionPriceParamsPriceNewSubscriptionGroupedWithProratedMinimumPriceCadenceOneTime

// This is an alias to an internal value.
const ReplaceSubscriptionPriceParamsPriceNewSubscriptionGroupedWithProratedMinimumPriceCadenceCustom = shared.ReplaceSubscriptionPriceParamsPriceNewSubscriptionGroupedWithProratedMinimumPriceCadenceCustom

// This is an alias to an internal type.
type ReplaceSubscriptionPriceParamsPriceNewSubscriptionGroupedWithProratedMinimumPriceModelType = shared.ReplaceSubscriptionPriceParamsPriceNewSubscriptionGroupedWithProratedMinimumPriceModelType

// This is an alias to an internal value.
const ReplaceSubscriptionPriceParamsPriceNewSubscriptionGroupedWithProratedMinimumPriceModelTypeGroupedWithProratedMinimum = shared.ReplaceSubscriptionPriceParamsPriceNewSubscriptionGroupedWithProratedMinimumPriceModelTypeGroupedWithProratedMinimum

// For custom cadence: specifies the duration of the billing period in days or
// months.
//
// This is an alias to an internal type.
type ReplaceSubscriptionPriceParamsPriceNewSubscriptionGroupedWithProratedMinimumPriceBillingCycleConfiguration = shared.ReplaceSubscriptionPriceParamsPriceNewSubscriptionGroupedWithProratedMinimumPriceBillingCycleConfiguration

// The unit of billing period duration.
//
// This is an alias to an internal type.
type ReplaceSubscriptionPriceParamsPriceNewSubscriptionGroupedWithProratedMinimumPriceBillingCycleConfigurationDurationUnit = shared.ReplaceSubscriptionPriceParamsPriceNewSubscriptionGroupedWithProratedMinimumPriceBillingCycleConfigurationDurationUnit

// This is an alias to an internal value.
const ReplaceSubscriptionPriceParamsPriceNewSubscriptionGroupedWithProratedMinimumPriceBillingCycleConfigurationDurationUnitDay = shared.ReplaceSubscriptionPriceParamsPriceNewSubscriptionGroupedWithProratedMinimumPriceBillingCycleConfigurationDurationUnitDay

// This is an alias to an internal value.
const ReplaceSubscriptionPriceParamsPriceNewSubscriptionGroupedWithProratedMinimumPriceBillingCycleConfigurationDurationUnitMonth = shared.ReplaceSubscriptionPriceParamsPriceNewSubscriptionGroupedWithProratedMinimumPriceBillingCycleConfigurationDurationUnitMonth

// Within each billing cycle, specifies the cadence at which invoices are produced.
// If unspecified, a single invoice is produced per billing cycle.
//
// This is an alias to an internal type.
type ReplaceSubscriptionPriceParamsPriceNewSubscriptionGroupedWithProratedMinimumPriceInvoicingCycleConfiguration = shared.ReplaceSubscriptionPriceParamsPriceNewSubscriptionGroupedWithProratedMinimumPriceInvoicingCycleConfiguration

// The unit of billing period duration.
//
// This is an alias to an internal type.
type ReplaceSubscriptionPriceParamsPriceNewSubscriptionGroupedWithProratedMinimumPriceInvoicingCycleConfigurationDurationUnit = shared.ReplaceSubscriptionPriceParamsPriceNewSubscriptionGroupedWithProratedMinimumPriceInvoicingCycleConfigurationDurationUnit

// This is an alias to an internal value.
const ReplaceSubscriptionPriceParamsPriceNewSubscriptionGroupedWithProratedMinimumPriceInvoicingCycleConfigurationDurationUnitDay = shared.ReplaceSubscriptionPriceParamsPriceNewSubscriptionGroupedWithProratedMinimumPriceInvoicingCycleConfigurationDurationUnitDay

// This is an alias to an internal value.
const ReplaceSubscriptionPriceParamsPriceNewSubscriptionGroupedWithProratedMinimumPriceInvoicingCycleConfigurationDurationUnitMonth = shared.ReplaceSubscriptionPriceParamsPriceNewSubscriptionGroupedWithProratedMinimumPriceInvoicingCycleConfigurationDurationUnitMonth

// This is an alias to an internal type.
type ReplaceSubscriptionPriceParamsPriceNewSubscriptionBulkWithProrationPrice = shared.ReplaceSubscriptionPriceParamsPriceNewSubscriptionBulkWithProrationPrice

// The cadence to bill for this price on.
//
// This is an alias to an internal type.
type ReplaceSubscriptionPriceParamsPriceNewSubscriptionBulkWithProrationPriceCadence = shared.ReplaceSubscriptionPriceParamsPriceNewSubscriptionBulkWithProrationPriceCadence

// This is an alias to an internal value.
const ReplaceSubscriptionPriceParamsPriceNewSubscriptionBulkWithProrationPriceCadenceAnnual = shared.ReplaceSubscriptionPriceParamsPriceNewSubscriptionBulkWithProrationPriceCadenceAnnual

// This is an alias to an internal value.
const ReplaceSubscriptionPriceParamsPriceNewSubscriptionBulkWithProrationPriceCadenceSemiAnnual = shared.ReplaceSubscriptionPriceParamsPriceNewSubscriptionBulkWithProrationPriceCadenceSemiAnnual

// This is an alias to an internal value.
const ReplaceSubscriptionPriceParamsPriceNewSubscriptionBulkWithProrationPriceCadenceMonthly = shared.ReplaceSubscriptionPriceParamsPriceNewSubscriptionBulkWithProrationPriceCadenceMonthly

// This is an alias to an internal value.
const ReplaceSubscriptionPriceParamsPriceNewSubscriptionBulkWithProrationPriceCadenceQuarterly = shared.ReplaceSubscriptionPriceParamsPriceNewSubscriptionBulkWithProrationPriceCadenceQuarterly

// This is an alias to an internal value.
const ReplaceSubscriptionPriceParamsPriceNewSubscriptionBulkWithProrationPriceCadenceOneTime = shared.ReplaceSubscriptionPriceParamsPriceNewSubscriptionBulkWithProrationPriceCadenceOneTime

// This is an alias to an internal value.
const ReplaceSubscriptionPriceParamsPriceNewSubscriptionBulkWithProrationPriceCadenceCustom = shared.ReplaceSubscriptionPriceParamsPriceNewSubscriptionBulkWithProrationPriceCadenceCustom

// This is an alias to an internal type.
type ReplaceSubscriptionPriceParamsPriceNewSubscriptionBulkWithProrationPriceModelType = shared.ReplaceSubscriptionPriceParamsPriceNewSubscriptionBulkWithProrationPriceModelType

// This is an alias to an internal value.
const ReplaceSubscriptionPriceParamsPriceNewSubscriptionBulkWithProrationPriceModelTypeBulkWithProration = shared.ReplaceSubscriptionPriceParamsPriceNewSubscriptionBulkWithProrationPriceModelTypeBulkWithProration

// For custom cadence: specifies the duration of the billing period in days or
// months.
//
// This is an alias to an internal type.
type ReplaceSubscriptionPriceParamsPriceNewSubscriptionBulkWithProrationPriceBillingCycleConfiguration = shared.ReplaceSubscriptionPriceParamsPriceNewSubscriptionBulkWithProrationPriceBillingCycleConfiguration

// The unit of billing period duration.
//
// This is an alias to an internal type.
type ReplaceSubscriptionPriceParamsPriceNewSubscriptionBulkWithProrationPriceBillingCycleConfigurationDurationUnit = shared.ReplaceSubscriptionPriceParamsPriceNewSubscriptionBulkWithProrationPriceBillingCycleConfigurationDurationUnit

// This is an alias to an internal value.
const ReplaceSubscriptionPriceParamsPriceNewSubscriptionBulkWithProrationPriceBillingCycleConfigurationDurationUnitDay = shared.ReplaceSubscriptionPriceParamsPriceNewSubscriptionBulkWithProrationPriceBillingCycleConfigurationDurationUnitDay

// This is an alias to an internal value.
const ReplaceSubscriptionPriceParamsPriceNewSubscriptionBulkWithProrationPriceBillingCycleConfigurationDurationUnitMonth = shared.ReplaceSubscriptionPriceParamsPriceNewSubscriptionBulkWithProrationPriceBillingCycleConfigurationDurationUnitMonth

// Within each billing cycle, specifies the cadence at which invoices are produced.
// If unspecified, a single invoice is produced per billing cycle.
//
// This is an alias to an internal type.
type ReplaceSubscriptionPriceParamsPriceNewSubscriptionBulkWithProrationPriceInvoicingCycleConfiguration = shared.ReplaceSubscriptionPriceParamsPriceNewSubscriptionBulkWithProrationPriceInvoicingCycleConfiguration

// The unit of billing period duration.
//
// This is an alias to an internal type.
type ReplaceSubscriptionPriceParamsPriceNewSubscriptionBulkWithProrationPriceInvoicingCycleConfigurationDurationUnit = shared.ReplaceSubscriptionPriceParamsPriceNewSubscriptionBulkWithProrationPriceInvoicingCycleConfigurationDurationUnit

// This is an alias to an internal value.
const ReplaceSubscriptionPriceParamsPriceNewSubscriptionBulkWithProrationPriceInvoicingCycleConfigurationDurationUnitDay = shared.ReplaceSubscriptionPriceParamsPriceNewSubscriptionBulkWithProrationPriceInvoicingCycleConfigurationDurationUnitDay

// This is an alias to an internal value.
const ReplaceSubscriptionPriceParamsPriceNewSubscriptionBulkWithProrationPriceInvoicingCycleConfigurationDurationUnitMonth = shared.ReplaceSubscriptionPriceParamsPriceNewSubscriptionBulkWithProrationPriceInvoicingCycleConfigurationDurationUnitMonth

// This is an alias to an internal type.
type ReplaceSubscriptionPriceParamsPriceNewSubscriptionScalableMatrixWithUnitPricingPrice = shared.ReplaceSubscriptionPriceParamsPriceNewSubscriptionScalableMatrixWithUnitPricingPrice

// The cadence to bill for this price on.
//
// This is an alias to an internal type.
type ReplaceSubscriptionPriceParamsPriceNewSubscriptionScalableMatrixWithUnitPricingPriceCadence = shared.ReplaceSubscriptionPriceParamsPriceNewSubscriptionScalableMatrixWithUnitPricingPriceCadence

// This is an alias to an internal value.
const ReplaceSubscriptionPriceParamsPriceNewSubscriptionScalableMatrixWithUnitPricingPriceCadenceAnnual = shared.ReplaceSubscriptionPriceParamsPriceNewSubscriptionScalableMatrixWithUnitPricingPriceCadenceAnnual

// This is an alias to an internal value.
const ReplaceSubscriptionPriceParamsPriceNewSubscriptionScalableMatrixWithUnitPricingPriceCadenceSemiAnnual = shared.ReplaceSubscriptionPriceParamsPriceNewSubscriptionScalableMatrixWithUnitPricingPriceCadenceSemiAnnual

// This is an alias to an internal value.
const ReplaceSubscriptionPriceParamsPriceNewSubscriptionScalableMatrixWithUnitPricingPriceCadenceMonthly = shared.ReplaceSubscriptionPriceParamsPriceNewSubscriptionScalableMatrixWithUnitPricingPriceCadenceMonthly

// This is an alias to an internal value.
const ReplaceSubscriptionPriceParamsPriceNewSubscriptionScalableMatrixWithUnitPricingPriceCadenceQuarterly = shared.ReplaceSubscriptionPriceParamsPriceNewSubscriptionScalableMatrixWithUnitPricingPriceCadenceQuarterly

// This is an alias to an internal value.
const ReplaceSubscriptionPriceParamsPriceNewSubscriptionScalableMatrixWithUnitPricingPriceCadenceOneTime = shared.ReplaceSubscriptionPriceParamsPriceNewSubscriptionScalableMatrixWithUnitPricingPriceCadenceOneTime

// This is an alias to an internal value.
const ReplaceSubscriptionPriceParamsPriceNewSubscriptionScalableMatrixWithUnitPricingPriceCadenceCustom = shared.ReplaceSubscriptionPriceParamsPriceNewSubscriptionScalableMatrixWithUnitPricingPriceCadenceCustom

// This is an alias to an internal type.
type ReplaceSubscriptionPriceParamsPriceNewSubscriptionScalableMatrixWithUnitPricingPriceModelType = shared.ReplaceSubscriptionPriceParamsPriceNewSubscriptionScalableMatrixWithUnitPricingPriceModelType

// This is an alias to an internal value.
const ReplaceSubscriptionPriceParamsPriceNewSubscriptionScalableMatrixWithUnitPricingPriceModelTypeScalableMatrixWithUnitPricing = shared.ReplaceSubscriptionPriceParamsPriceNewSubscriptionScalableMatrixWithUnitPricingPriceModelTypeScalableMatrixWithUnitPricing

// For custom cadence: specifies the duration of the billing period in days or
// months.
//
// This is an alias to an internal type.
type ReplaceSubscriptionPriceParamsPriceNewSubscriptionScalableMatrixWithUnitPricingPriceBillingCycleConfiguration = shared.ReplaceSubscriptionPriceParamsPriceNewSubscriptionScalableMatrixWithUnitPricingPriceBillingCycleConfiguration

// The unit of billing period duration.
//
// This is an alias to an internal type.
type ReplaceSubscriptionPriceParamsPriceNewSubscriptionScalableMatrixWithUnitPricingPriceBillingCycleConfigurationDurationUnit = shared.ReplaceSubscriptionPriceParamsPriceNewSubscriptionScalableMatrixWithUnitPricingPriceBillingCycleConfigurationDurationUnit

// This is an alias to an internal value.
const ReplaceSubscriptionPriceParamsPriceNewSubscriptionScalableMatrixWithUnitPricingPriceBillingCycleConfigurationDurationUnitDay = shared.ReplaceSubscriptionPriceParamsPriceNewSubscriptionScalableMatrixWithUnitPricingPriceBillingCycleConfigurationDurationUnitDay

// This is an alias to an internal value.
const ReplaceSubscriptionPriceParamsPriceNewSubscriptionScalableMatrixWithUnitPricingPriceBillingCycleConfigurationDurationUnitMonth = shared.ReplaceSubscriptionPriceParamsPriceNewSubscriptionScalableMatrixWithUnitPricingPriceBillingCycleConfigurationDurationUnitMonth

// Within each billing cycle, specifies the cadence at which invoices are produced.
// If unspecified, a single invoice is produced per billing cycle.
//
// This is an alias to an internal type.
type ReplaceSubscriptionPriceParamsPriceNewSubscriptionScalableMatrixWithUnitPricingPriceInvoicingCycleConfiguration = shared.ReplaceSubscriptionPriceParamsPriceNewSubscriptionScalableMatrixWithUnitPricingPriceInvoicingCycleConfiguration

// The unit of billing period duration.
//
// This is an alias to an internal type.
type ReplaceSubscriptionPriceParamsPriceNewSubscriptionScalableMatrixWithUnitPricingPriceInvoicingCycleConfigurationDurationUnit = shared.ReplaceSubscriptionPriceParamsPriceNewSubscriptionScalableMatrixWithUnitPricingPriceInvoicingCycleConfigurationDurationUnit

// This is an alias to an internal value.
const ReplaceSubscriptionPriceParamsPriceNewSubscriptionScalableMatrixWithUnitPricingPriceInvoicingCycleConfigurationDurationUnitDay = shared.ReplaceSubscriptionPriceParamsPriceNewSubscriptionScalableMatrixWithUnitPricingPriceInvoicingCycleConfigurationDurationUnitDay

// This is an alias to an internal value.
const ReplaceSubscriptionPriceParamsPriceNewSubscriptionScalableMatrixWithUnitPricingPriceInvoicingCycleConfigurationDurationUnitMonth = shared.ReplaceSubscriptionPriceParamsPriceNewSubscriptionScalableMatrixWithUnitPricingPriceInvoicingCycleConfigurationDurationUnitMonth

// This is an alias to an internal type.
type ReplaceSubscriptionPriceParamsPriceNewSubscriptionScalableMatrixWithTieredPricingPrice = shared.ReplaceSubscriptionPriceParamsPriceNewSubscriptionScalableMatrixWithTieredPricingPrice

// The cadence to bill for this price on.
//
// This is an alias to an internal type.
type ReplaceSubscriptionPriceParamsPriceNewSubscriptionScalableMatrixWithTieredPricingPriceCadence = shared.ReplaceSubscriptionPriceParamsPriceNewSubscriptionScalableMatrixWithTieredPricingPriceCadence

// This is an alias to an internal value.
const ReplaceSubscriptionPriceParamsPriceNewSubscriptionScalableMatrixWithTieredPricingPriceCadenceAnnual = shared.ReplaceSubscriptionPriceParamsPriceNewSubscriptionScalableMatrixWithTieredPricingPriceCadenceAnnual

// This is an alias to an internal value.
const ReplaceSubscriptionPriceParamsPriceNewSubscriptionScalableMatrixWithTieredPricingPriceCadenceSemiAnnual = shared.ReplaceSubscriptionPriceParamsPriceNewSubscriptionScalableMatrixWithTieredPricingPriceCadenceSemiAnnual

// This is an alias to an internal value.
const ReplaceSubscriptionPriceParamsPriceNewSubscriptionScalableMatrixWithTieredPricingPriceCadenceMonthly = shared.ReplaceSubscriptionPriceParamsPriceNewSubscriptionScalableMatrixWithTieredPricingPriceCadenceMonthly

// This is an alias to an internal value.
const ReplaceSubscriptionPriceParamsPriceNewSubscriptionScalableMatrixWithTieredPricingPriceCadenceQuarterly = shared.ReplaceSubscriptionPriceParamsPriceNewSubscriptionScalableMatrixWithTieredPricingPriceCadenceQuarterly

// This is an alias to an internal value.
const ReplaceSubscriptionPriceParamsPriceNewSubscriptionScalableMatrixWithTieredPricingPriceCadenceOneTime = shared.ReplaceSubscriptionPriceParamsPriceNewSubscriptionScalableMatrixWithTieredPricingPriceCadenceOneTime

// This is an alias to an internal value.
const ReplaceSubscriptionPriceParamsPriceNewSubscriptionScalableMatrixWithTieredPricingPriceCadenceCustom = shared.ReplaceSubscriptionPriceParamsPriceNewSubscriptionScalableMatrixWithTieredPricingPriceCadenceCustom

// This is an alias to an internal type.
type ReplaceSubscriptionPriceParamsPriceNewSubscriptionScalableMatrixWithTieredPricingPriceModelType = shared.ReplaceSubscriptionPriceParamsPriceNewSubscriptionScalableMatrixWithTieredPricingPriceModelType

// This is an alias to an internal value.
const ReplaceSubscriptionPriceParamsPriceNewSubscriptionScalableMatrixWithTieredPricingPriceModelTypeScalableMatrixWithTieredPricing = shared.ReplaceSubscriptionPriceParamsPriceNewSubscriptionScalableMatrixWithTieredPricingPriceModelTypeScalableMatrixWithTieredPricing

// For custom cadence: specifies the duration of the billing period in days or
// months.
//
// This is an alias to an internal type.
type ReplaceSubscriptionPriceParamsPriceNewSubscriptionScalableMatrixWithTieredPricingPriceBillingCycleConfiguration = shared.ReplaceSubscriptionPriceParamsPriceNewSubscriptionScalableMatrixWithTieredPricingPriceBillingCycleConfiguration

// The unit of billing period duration.
//
// This is an alias to an internal type.
type ReplaceSubscriptionPriceParamsPriceNewSubscriptionScalableMatrixWithTieredPricingPriceBillingCycleConfigurationDurationUnit = shared.ReplaceSubscriptionPriceParamsPriceNewSubscriptionScalableMatrixWithTieredPricingPriceBillingCycleConfigurationDurationUnit

// This is an alias to an internal value.
const ReplaceSubscriptionPriceParamsPriceNewSubscriptionScalableMatrixWithTieredPricingPriceBillingCycleConfigurationDurationUnitDay = shared.ReplaceSubscriptionPriceParamsPriceNewSubscriptionScalableMatrixWithTieredPricingPriceBillingCycleConfigurationDurationUnitDay

// This is an alias to an internal value.
const ReplaceSubscriptionPriceParamsPriceNewSubscriptionScalableMatrixWithTieredPricingPriceBillingCycleConfigurationDurationUnitMonth = shared.ReplaceSubscriptionPriceParamsPriceNewSubscriptionScalableMatrixWithTieredPricingPriceBillingCycleConfigurationDurationUnitMonth

// Within each billing cycle, specifies the cadence at which invoices are produced.
// If unspecified, a single invoice is produced per billing cycle.
//
// This is an alias to an internal type.
type ReplaceSubscriptionPriceParamsPriceNewSubscriptionScalableMatrixWithTieredPricingPriceInvoicingCycleConfiguration = shared.ReplaceSubscriptionPriceParamsPriceNewSubscriptionScalableMatrixWithTieredPricingPriceInvoicingCycleConfiguration

// The unit of billing period duration.
//
// This is an alias to an internal type.
type ReplaceSubscriptionPriceParamsPriceNewSubscriptionScalableMatrixWithTieredPricingPriceInvoicingCycleConfigurationDurationUnit = shared.ReplaceSubscriptionPriceParamsPriceNewSubscriptionScalableMatrixWithTieredPricingPriceInvoicingCycleConfigurationDurationUnit

// This is an alias to an internal value.
const ReplaceSubscriptionPriceParamsPriceNewSubscriptionScalableMatrixWithTieredPricingPriceInvoicingCycleConfigurationDurationUnitDay = shared.ReplaceSubscriptionPriceParamsPriceNewSubscriptionScalableMatrixWithTieredPricingPriceInvoicingCycleConfigurationDurationUnitDay

// This is an alias to an internal value.
const ReplaceSubscriptionPriceParamsPriceNewSubscriptionScalableMatrixWithTieredPricingPriceInvoicingCycleConfigurationDurationUnitMonth = shared.ReplaceSubscriptionPriceParamsPriceNewSubscriptionScalableMatrixWithTieredPricingPriceInvoicingCycleConfigurationDurationUnitMonth

// This is an alias to an internal type.
type ReplaceSubscriptionPriceParamsPriceNewSubscriptionCumulativeGroupedBulkPrice = shared.ReplaceSubscriptionPriceParamsPriceNewSubscriptionCumulativeGroupedBulkPrice

// The cadence to bill for this price on.
//
// This is an alias to an internal type.
type ReplaceSubscriptionPriceParamsPriceNewSubscriptionCumulativeGroupedBulkPriceCadence = shared.ReplaceSubscriptionPriceParamsPriceNewSubscriptionCumulativeGroupedBulkPriceCadence

// This is an alias to an internal value.
const ReplaceSubscriptionPriceParamsPriceNewSubscriptionCumulativeGroupedBulkPriceCadenceAnnual = shared.ReplaceSubscriptionPriceParamsPriceNewSubscriptionCumulativeGroupedBulkPriceCadenceAnnual

// This is an alias to an internal value.
const ReplaceSubscriptionPriceParamsPriceNewSubscriptionCumulativeGroupedBulkPriceCadenceSemiAnnual = shared.ReplaceSubscriptionPriceParamsPriceNewSubscriptionCumulativeGroupedBulkPriceCadenceSemiAnnual

// This is an alias to an internal value.
const ReplaceSubscriptionPriceParamsPriceNewSubscriptionCumulativeGroupedBulkPriceCadenceMonthly = shared.ReplaceSubscriptionPriceParamsPriceNewSubscriptionCumulativeGroupedBulkPriceCadenceMonthly

// This is an alias to an internal value.
const ReplaceSubscriptionPriceParamsPriceNewSubscriptionCumulativeGroupedBulkPriceCadenceQuarterly = shared.ReplaceSubscriptionPriceParamsPriceNewSubscriptionCumulativeGroupedBulkPriceCadenceQuarterly

// This is an alias to an internal value.
const ReplaceSubscriptionPriceParamsPriceNewSubscriptionCumulativeGroupedBulkPriceCadenceOneTime = shared.ReplaceSubscriptionPriceParamsPriceNewSubscriptionCumulativeGroupedBulkPriceCadenceOneTime

// This is an alias to an internal value.
const ReplaceSubscriptionPriceParamsPriceNewSubscriptionCumulativeGroupedBulkPriceCadenceCustom = shared.ReplaceSubscriptionPriceParamsPriceNewSubscriptionCumulativeGroupedBulkPriceCadenceCustom

// This is an alias to an internal type.
type ReplaceSubscriptionPriceParamsPriceNewSubscriptionCumulativeGroupedBulkPriceModelType = shared.ReplaceSubscriptionPriceParamsPriceNewSubscriptionCumulativeGroupedBulkPriceModelType

// This is an alias to an internal value.
const ReplaceSubscriptionPriceParamsPriceNewSubscriptionCumulativeGroupedBulkPriceModelTypeCumulativeGroupedBulk = shared.ReplaceSubscriptionPriceParamsPriceNewSubscriptionCumulativeGroupedBulkPriceModelTypeCumulativeGroupedBulk

// For custom cadence: specifies the duration of the billing period in days or
// months.
//
// This is an alias to an internal type.
type ReplaceSubscriptionPriceParamsPriceNewSubscriptionCumulativeGroupedBulkPriceBillingCycleConfiguration = shared.ReplaceSubscriptionPriceParamsPriceNewSubscriptionCumulativeGroupedBulkPriceBillingCycleConfiguration

// The unit of billing period duration.
//
// This is an alias to an internal type.
type ReplaceSubscriptionPriceParamsPriceNewSubscriptionCumulativeGroupedBulkPriceBillingCycleConfigurationDurationUnit = shared.ReplaceSubscriptionPriceParamsPriceNewSubscriptionCumulativeGroupedBulkPriceBillingCycleConfigurationDurationUnit

// This is an alias to an internal value.
const ReplaceSubscriptionPriceParamsPriceNewSubscriptionCumulativeGroupedBulkPriceBillingCycleConfigurationDurationUnitDay = shared.ReplaceSubscriptionPriceParamsPriceNewSubscriptionCumulativeGroupedBulkPriceBillingCycleConfigurationDurationUnitDay

// This is an alias to an internal value.
const ReplaceSubscriptionPriceParamsPriceNewSubscriptionCumulativeGroupedBulkPriceBillingCycleConfigurationDurationUnitMonth = shared.ReplaceSubscriptionPriceParamsPriceNewSubscriptionCumulativeGroupedBulkPriceBillingCycleConfigurationDurationUnitMonth

// Within each billing cycle, specifies the cadence at which invoices are produced.
// If unspecified, a single invoice is produced per billing cycle.
//
// This is an alias to an internal type.
type ReplaceSubscriptionPriceParamsPriceNewSubscriptionCumulativeGroupedBulkPriceInvoicingCycleConfiguration = shared.ReplaceSubscriptionPriceParamsPriceNewSubscriptionCumulativeGroupedBulkPriceInvoicingCycleConfiguration

// The unit of billing period duration.
//
// This is an alias to an internal type.
type ReplaceSubscriptionPriceParamsPriceNewSubscriptionCumulativeGroupedBulkPriceInvoicingCycleConfigurationDurationUnit = shared.ReplaceSubscriptionPriceParamsPriceNewSubscriptionCumulativeGroupedBulkPriceInvoicingCycleConfigurationDurationUnit

// This is an alias to an internal value.
const ReplaceSubscriptionPriceParamsPriceNewSubscriptionCumulativeGroupedBulkPriceInvoicingCycleConfigurationDurationUnitDay = shared.ReplaceSubscriptionPriceParamsPriceNewSubscriptionCumulativeGroupedBulkPriceInvoicingCycleConfigurationDurationUnitDay

// This is an alias to an internal value.
const ReplaceSubscriptionPriceParamsPriceNewSubscriptionCumulativeGroupedBulkPriceInvoicingCycleConfigurationDurationUnitMonth = shared.ReplaceSubscriptionPriceParamsPriceNewSubscriptionCumulativeGroupedBulkPriceInvoicingCycleConfigurationDurationUnitMonth

// This is an alias to an internal type.
type ReplaceSubscriptionPriceParamsPriceNewSubscriptionMaxGroupTieredPackagePrice = shared.ReplaceSubscriptionPriceParamsPriceNewSubscriptionMaxGroupTieredPackagePrice

// The cadence to bill for this price on.
//
// This is an alias to an internal type.
type ReplaceSubscriptionPriceParamsPriceNewSubscriptionMaxGroupTieredPackagePriceCadence = shared.ReplaceSubscriptionPriceParamsPriceNewSubscriptionMaxGroupTieredPackagePriceCadence

// This is an alias to an internal value.
const ReplaceSubscriptionPriceParamsPriceNewSubscriptionMaxGroupTieredPackagePriceCadenceAnnual = shared.ReplaceSubscriptionPriceParamsPriceNewSubscriptionMaxGroupTieredPackagePriceCadenceAnnual

// This is an alias to an internal value.
const ReplaceSubscriptionPriceParamsPriceNewSubscriptionMaxGroupTieredPackagePriceCadenceSemiAnnual = shared.ReplaceSubscriptionPriceParamsPriceNewSubscriptionMaxGroupTieredPackagePriceCadenceSemiAnnual

// This is an alias to an internal value.
const ReplaceSubscriptionPriceParamsPriceNewSubscriptionMaxGroupTieredPackagePriceCadenceMonthly = shared.ReplaceSubscriptionPriceParamsPriceNewSubscriptionMaxGroupTieredPackagePriceCadenceMonthly

// This is an alias to an internal value.
const ReplaceSubscriptionPriceParamsPriceNewSubscriptionMaxGroupTieredPackagePriceCadenceQuarterly = shared.ReplaceSubscriptionPriceParamsPriceNewSubscriptionMaxGroupTieredPackagePriceCadenceQuarterly

// This is an alias to an internal value.
const ReplaceSubscriptionPriceParamsPriceNewSubscriptionMaxGroupTieredPackagePriceCadenceOneTime = shared.ReplaceSubscriptionPriceParamsPriceNewSubscriptionMaxGroupTieredPackagePriceCadenceOneTime

// This is an alias to an internal value.
const ReplaceSubscriptionPriceParamsPriceNewSubscriptionMaxGroupTieredPackagePriceCadenceCustom = shared.ReplaceSubscriptionPriceParamsPriceNewSubscriptionMaxGroupTieredPackagePriceCadenceCustom

// This is an alias to an internal type.
type ReplaceSubscriptionPriceParamsPriceNewSubscriptionMaxGroupTieredPackagePriceModelType = shared.ReplaceSubscriptionPriceParamsPriceNewSubscriptionMaxGroupTieredPackagePriceModelType

// This is an alias to an internal value.
const ReplaceSubscriptionPriceParamsPriceNewSubscriptionMaxGroupTieredPackagePriceModelTypeMaxGroupTieredPackage = shared.ReplaceSubscriptionPriceParamsPriceNewSubscriptionMaxGroupTieredPackagePriceModelTypeMaxGroupTieredPackage

// For custom cadence: specifies the duration of the billing period in days or
// months.
//
// This is an alias to an internal type.
type ReplaceSubscriptionPriceParamsPriceNewSubscriptionMaxGroupTieredPackagePriceBillingCycleConfiguration = shared.ReplaceSubscriptionPriceParamsPriceNewSubscriptionMaxGroupTieredPackagePriceBillingCycleConfiguration

// The unit of billing period duration.
//
// This is an alias to an internal type.
type ReplaceSubscriptionPriceParamsPriceNewSubscriptionMaxGroupTieredPackagePriceBillingCycleConfigurationDurationUnit = shared.ReplaceSubscriptionPriceParamsPriceNewSubscriptionMaxGroupTieredPackagePriceBillingCycleConfigurationDurationUnit

// This is an alias to an internal value.
const ReplaceSubscriptionPriceParamsPriceNewSubscriptionMaxGroupTieredPackagePriceBillingCycleConfigurationDurationUnitDay = shared.ReplaceSubscriptionPriceParamsPriceNewSubscriptionMaxGroupTieredPackagePriceBillingCycleConfigurationDurationUnitDay

// This is an alias to an internal value.
const ReplaceSubscriptionPriceParamsPriceNewSubscriptionMaxGroupTieredPackagePriceBillingCycleConfigurationDurationUnitMonth = shared.ReplaceSubscriptionPriceParamsPriceNewSubscriptionMaxGroupTieredPackagePriceBillingCycleConfigurationDurationUnitMonth

// Within each billing cycle, specifies the cadence at which invoices are produced.
// If unspecified, a single invoice is produced per billing cycle.
//
// This is an alias to an internal type.
type ReplaceSubscriptionPriceParamsPriceNewSubscriptionMaxGroupTieredPackagePriceInvoicingCycleConfiguration = shared.ReplaceSubscriptionPriceParamsPriceNewSubscriptionMaxGroupTieredPackagePriceInvoicingCycleConfiguration

// The unit of billing period duration.
//
// This is an alias to an internal type.
type ReplaceSubscriptionPriceParamsPriceNewSubscriptionMaxGroupTieredPackagePriceInvoicingCycleConfigurationDurationUnit = shared.ReplaceSubscriptionPriceParamsPriceNewSubscriptionMaxGroupTieredPackagePriceInvoicingCycleConfigurationDurationUnit

// This is an alias to an internal value.
const ReplaceSubscriptionPriceParamsPriceNewSubscriptionMaxGroupTieredPackagePriceInvoicingCycleConfigurationDurationUnitDay = shared.ReplaceSubscriptionPriceParamsPriceNewSubscriptionMaxGroupTieredPackagePriceInvoicingCycleConfigurationDurationUnitDay

// This is an alias to an internal value.
const ReplaceSubscriptionPriceParamsPriceNewSubscriptionMaxGroupTieredPackagePriceInvoicingCycleConfigurationDurationUnitMonth = shared.ReplaceSubscriptionPriceParamsPriceNewSubscriptionMaxGroupTieredPackagePriceInvoicingCycleConfigurationDurationUnitMonth

// This is an alias to an internal type.
type ReplaceSubscriptionPriceParamsPriceNewSubscriptionGroupedWithMeteredMinimumPrice = shared.ReplaceSubscriptionPriceParamsPriceNewSubscriptionGroupedWithMeteredMinimumPrice

// The cadence to bill for this price on.
//
// This is an alias to an internal type.
type ReplaceSubscriptionPriceParamsPriceNewSubscriptionGroupedWithMeteredMinimumPriceCadence = shared.ReplaceSubscriptionPriceParamsPriceNewSubscriptionGroupedWithMeteredMinimumPriceCadence

// This is an alias to an internal value.
const ReplaceSubscriptionPriceParamsPriceNewSubscriptionGroupedWithMeteredMinimumPriceCadenceAnnual = shared.ReplaceSubscriptionPriceParamsPriceNewSubscriptionGroupedWithMeteredMinimumPriceCadenceAnnual

// This is an alias to an internal value.
const ReplaceSubscriptionPriceParamsPriceNewSubscriptionGroupedWithMeteredMinimumPriceCadenceSemiAnnual = shared.ReplaceSubscriptionPriceParamsPriceNewSubscriptionGroupedWithMeteredMinimumPriceCadenceSemiAnnual

// This is an alias to an internal value.
const ReplaceSubscriptionPriceParamsPriceNewSubscriptionGroupedWithMeteredMinimumPriceCadenceMonthly = shared.ReplaceSubscriptionPriceParamsPriceNewSubscriptionGroupedWithMeteredMinimumPriceCadenceMonthly

// This is an alias to an internal value.
const ReplaceSubscriptionPriceParamsPriceNewSubscriptionGroupedWithMeteredMinimumPriceCadenceQuarterly = shared.ReplaceSubscriptionPriceParamsPriceNewSubscriptionGroupedWithMeteredMinimumPriceCadenceQuarterly

// This is an alias to an internal value.
const ReplaceSubscriptionPriceParamsPriceNewSubscriptionGroupedWithMeteredMinimumPriceCadenceOneTime = shared.ReplaceSubscriptionPriceParamsPriceNewSubscriptionGroupedWithMeteredMinimumPriceCadenceOneTime

// This is an alias to an internal value.
const ReplaceSubscriptionPriceParamsPriceNewSubscriptionGroupedWithMeteredMinimumPriceCadenceCustom = shared.ReplaceSubscriptionPriceParamsPriceNewSubscriptionGroupedWithMeteredMinimumPriceCadenceCustom

// This is an alias to an internal type.
type ReplaceSubscriptionPriceParamsPriceNewSubscriptionGroupedWithMeteredMinimumPriceModelType = shared.ReplaceSubscriptionPriceParamsPriceNewSubscriptionGroupedWithMeteredMinimumPriceModelType

// This is an alias to an internal value.
const ReplaceSubscriptionPriceParamsPriceNewSubscriptionGroupedWithMeteredMinimumPriceModelTypeGroupedWithMeteredMinimum = shared.ReplaceSubscriptionPriceParamsPriceNewSubscriptionGroupedWithMeteredMinimumPriceModelTypeGroupedWithMeteredMinimum

// For custom cadence: specifies the duration of the billing period in days or
// months.
//
// This is an alias to an internal type.
type ReplaceSubscriptionPriceParamsPriceNewSubscriptionGroupedWithMeteredMinimumPriceBillingCycleConfiguration = shared.ReplaceSubscriptionPriceParamsPriceNewSubscriptionGroupedWithMeteredMinimumPriceBillingCycleConfiguration

// The unit of billing period duration.
//
// This is an alias to an internal type.
type ReplaceSubscriptionPriceParamsPriceNewSubscriptionGroupedWithMeteredMinimumPriceBillingCycleConfigurationDurationUnit = shared.ReplaceSubscriptionPriceParamsPriceNewSubscriptionGroupedWithMeteredMinimumPriceBillingCycleConfigurationDurationUnit

// This is an alias to an internal value.
const ReplaceSubscriptionPriceParamsPriceNewSubscriptionGroupedWithMeteredMinimumPriceBillingCycleConfigurationDurationUnitDay = shared.ReplaceSubscriptionPriceParamsPriceNewSubscriptionGroupedWithMeteredMinimumPriceBillingCycleConfigurationDurationUnitDay

// This is an alias to an internal value.
const ReplaceSubscriptionPriceParamsPriceNewSubscriptionGroupedWithMeteredMinimumPriceBillingCycleConfigurationDurationUnitMonth = shared.ReplaceSubscriptionPriceParamsPriceNewSubscriptionGroupedWithMeteredMinimumPriceBillingCycleConfigurationDurationUnitMonth

// Within each billing cycle, specifies the cadence at which invoices are produced.
// If unspecified, a single invoice is produced per billing cycle.
//
// This is an alias to an internal type.
type ReplaceSubscriptionPriceParamsPriceNewSubscriptionGroupedWithMeteredMinimumPriceInvoicingCycleConfiguration = shared.ReplaceSubscriptionPriceParamsPriceNewSubscriptionGroupedWithMeteredMinimumPriceInvoicingCycleConfiguration

// The unit of billing period duration.
//
// This is an alias to an internal type.
type ReplaceSubscriptionPriceParamsPriceNewSubscriptionGroupedWithMeteredMinimumPriceInvoicingCycleConfigurationDurationUnit = shared.ReplaceSubscriptionPriceParamsPriceNewSubscriptionGroupedWithMeteredMinimumPriceInvoicingCycleConfigurationDurationUnit

// This is an alias to an internal value.
const ReplaceSubscriptionPriceParamsPriceNewSubscriptionGroupedWithMeteredMinimumPriceInvoicingCycleConfigurationDurationUnitDay = shared.ReplaceSubscriptionPriceParamsPriceNewSubscriptionGroupedWithMeteredMinimumPriceInvoicingCycleConfigurationDurationUnitDay

// This is an alias to an internal value.
const ReplaceSubscriptionPriceParamsPriceNewSubscriptionGroupedWithMeteredMinimumPriceInvoicingCycleConfigurationDurationUnitMonth = shared.ReplaceSubscriptionPriceParamsPriceNewSubscriptionGroupedWithMeteredMinimumPriceInvoicingCycleConfigurationDurationUnitMonth

// This is an alias to an internal type.
type ReplaceSubscriptionPriceParamsPriceNewSubscriptionMatrixWithDisplayNamePrice = shared.ReplaceSubscriptionPriceParamsPriceNewSubscriptionMatrixWithDisplayNamePrice

// The cadence to bill for this price on.
//
// This is an alias to an internal type.
type ReplaceSubscriptionPriceParamsPriceNewSubscriptionMatrixWithDisplayNamePriceCadence = shared.ReplaceSubscriptionPriceParamsPriceNewSubscriptionMatrixWithDisplayNamePriceCadence

// This is an alias to an internal value.
const ReplaceSubscriptionPriceParamsPriceNewSubscriptionMatrixWithDisplayNamePriceCadenceAnnual = shared.ReplaceSubscriptionPriceParamsPriceNewSubscriptionMatrixWithDisplayNamePriceCadenceAnnual

// This is an alias to an internal value.
const ReplaceSubscriptionPriceParamsPriceNewSubscriptionMatrixWithDisplayNamePriceCadenceSemiAnnual = shared.ReplaceSubscriptionPriceParamsPriceNewSubscriptionMatrixWithDisplayNamePriceCadenceSemiAnnual

// This is an alias to an internal value.
const ReplaceSubscriptionPriceParamsPriceNewSubscriptionMatrixWithDisplayNamePriceCadenceMonthly = shared.ReplaceSubscriptionPriceParamsPriceNewSubscriptionMatrixWithDisplayNamePriceCadenceMonthly

// This is an alias to an internal value.
const ReplaceSubscriptionPriceParamsPriceNewSubscriptionMatrixWithDisplayNamePriceCadenceQuarterly = shared.ReplaceSubscriptionPriceParamsPriceNewSubscriptionMatrixWithDisplayNamePriceCadenceQuarterly

// This is an alias to an internal value.
const ReplaceSubscriptionPriceParamsPriceNewSubscriptionMatrixWithDisplayNamePriceCadenceOneTime = shared.ReplaceSubscriptionPriceParamsPriceNewSubscriptionMatrixWithDisplayNamePriceCadenceOneTime

// This is an alias to an internal value.
const ReplaceSubscriptionPriceParamsPriceNewSubscriptionMatrixWithDisplayNamePriceCadenceCustom = shared.ReplaceSubscriptionPriceParamsPriceNewSubscriptionMatrixWithDisplayNamePriceCadenceCustom

// This is an alias to an internal type.
type ReplaceSubscriptionPriceParamsPriceNewSubscriptionMatrixWithDisplayNamePriceModelType = shared.ReplaceSubscriptionPriceParamsPriceNewSubscriptionMatrixWithDisplayNamePriceModelType

// This is an alias to an internal value.
const ReplaceSubscriptionPriceParamsPriceNewSubscriptionMatrixWithDisplayNamePriceModelTypeMatrixWithDisplayName = shared.ReplaceSubscriptionPriceParamsPriceNewSubscriptionMatrixWithDisplayNamePriceModelTypeMatrixWithDisplayName

// For custom cadence: specifies the duration of the billing period in days or
// months.
//
// This is an alias to an internal type.
type ReplaceSubscriptionPriceParamsPriceNewSubscriptionMatrixWithDisplayNamePriceBillingCycleConfiguration = shared.ReplaceSubscriptionPriceParamsPriceNewSubscriptionMatrixWithDisplayNamePriceBillingCycleConfiguration

// The unit of billing period duration.
//
// This is an alias to an internal type.
type ReplaceSubscriptionPriceParamsPriceNewSubscriptionMatrixWithDisplayNamePriceBillingCycleConfigurationDurationUnit = shared.ReplaceSubscriptionPriceParamsPriceNewSubscriptionMatrixWithDisplayNamePriceBillingCycleConfigurationDurationUnit

// This is an alias to an internal value.
const ReplaceSubscriptionPriceParamsPriceNewSubscriptionMatrixWithDisplayNamePriceBillingCycleConfigurationDurationUnitDay = shared.ReplaceSubscriptionPriceParamsPriceNewSubscriptionMatrixWithDisplayNamePriceBillingCycleConfigurationDurationUnitDay

// This is an alias to an internal value.
const ReplaceSubscriptionPriceParamsPriceNewSubscriptionMatrixWithDisplayNamePriceBillingCycleConfigurationDurationUnitMonth = shared.ReplaceSubscriptionPriceParamsPriceNewSubscriptionMatrixWithDisplayNamePriceBillingCycleConfigurationDurationUnitMonth

// Within each billing cycle, specifies the cadence at which invoices are produced.
// If unspecified, a single invoice is produced per billing cycle.
//
// This is an alias to an internal type.
type ReplaceSubscriptionPriceParamsPriceNewSubscriptionMatrixWithDisplayNamePriceInvoicingCycleConfiguration = shared.ReplaceSubscriptionPriceParamsPriceNewSubscriptionMatrixWithDisplayNamePriceInvoicingCycleConfiguration

// The unit of billing period duration.
//
// This is an alias to an internal type.
type ReplaceSubscriptionPriceParamsPriceNewSubscriptionMatrixWithDisplayNamePriceInvoicingCycleConfigurationDurationUnit = shared.ReplaceSubscriptionPriceParamsPriceNewSubscriptionMatrixWithDisplayNamePriceInvoicingCycleConfigurationDurationUnit

// This is an alias to an internal value.
const ReplaceSubscriptionPriceParamsPriceNewSubscriptionMatrixWithDisplayNamePriceInvoicingCycleConfigurationDurationUnitDay = shared.ReplaceSubscriptionPriceParamsPriceNewSubscriptionMatrixWithDisplayNamePriceInvoicingCycleConfigurationDurationUnitDay

// This is an alias to an internal value.
const ReplaceSubscriptionPriceParamsPriceNewSubscriptionMatrixWithDisplayNamePriceInvoicingCycleConfigurationDurationUnitMonth = shared.ReplaceSubscriptionPriceParamsPriceNewSubscriptionMatrixWithDisplayNamePriceInvoicingCycleConfigurationDurationUnitMonth

// This is an alias to an internal type.
type ReplaceSubscriptionPriceParamsPriceNewSubscriptionGroupedTieredPackagePrice = shared.ReplaceSubscriptionPriceParamsPriceNewSubscriptionGroupedTieredPackagePrice

// The cadence to bill for this price on.
//
// This is an alias to an internal type.
type ReplaceSubscriptionPriceParamsPriceNewSubscriptionGroupedTieredPackagePriceCadence = shared.ReplaceSubscriptionPriceParamsPriceNewSubscriptionGroupedTieredPackagePriceCadence

// This is an alias to an internal value.
const ReplaceSubscriptionPriceParamsPriceNewSubscriptionGroupedTieredPackagePriceCadenceAnnual = shared.ReplaceSubscriptionPriceParamsPriceNewSubscriptionGroupedTieredPackagePriceCadenceAnnual

// This is an alias to an internal value.
const ReplaceSubscriptionPriceParamsPriceNewSubscriptionGroupedTieredPackagePriceCadenceSemiAnnual = shared.ReplaceSubscriptionPriceParamsPriceNewSubscriptionGroupedTieredPackagePriceCadenceSemiAnnual

// This is an alias to an internal value.
const ReplaceSubscriptionPriceParamsPriceNewSubscriptionGroupedTieredPackagePriceCadenceMonthly = shared.ReplaceSubscriptionPriceParamsPriceNewSubscriptionGroupedTieredPackagePriceCadenceMonthly

// This is an alias to an internal value.
const ReplaceSubscriptionPriceParamsPriceNewSubscriptionGroupedTieredPackagePriceCadenceQuarterly = shared.ReplaceSubscriptionPriceParamsPriceNewSubscriptionGroupedTieredPackagePriceCadenceQuarterly

// This is an alias to an internal value.
const ReplaceSubscriptionPriceParamsPriceNewSubscriptionGroupedTieredPackagePriceCadenceOneTime = shared.ReplaceSubscriptionPriceParamsPriceNewSubscriptionGroupedTieredPackagePriceCadenceOneTime

// This is an alias to an internal value.
const ReplaceSubscriptionPriceParamsPriceNewSubscriptionGroupedTieredPackagePriceCadenceCustom = shared.ReplaceSubscriptionPriceParamsPriceNewSubscriptionGroupedTieredPackagePriceCadenceCustom

// This is an alias to an internal type.
type ReplaceSubscriptionPriceParamsPriceNewSubscriptionGroupedTieredPackagePriceModelType = shared.ReplaceSubscriptionPriceParamsPriceNewSubscriptionGroupedTieredPackagePriceModelType

// This is an alias to an internal value.
const ReplaceSubscriptionPriceParamsPriceNewSubscriptionGroupedTieredPackagePriceModelTypeGroupedTieredPackage = shared.ReplaceSubscriptionPriceParamsPriceNewSubscriptionGroupedTieredPackagePriceModelTypeGroupedTieredPackage

// For custom cadence: specifies the duration of the billing period in days or
// months.
//
// This is an alias to an internal type.
type ReplaceSubscriptionPriceParamsPriceNewSubscriptionGroupedTieredPackagePriceBillingCycleConfiguration = shared.ReplaceSubscriptionPriceParamsPriceNewSubscriptionGroupedTieredPackagePriceBillingCycleConfiguration

// The unit of billing period duration.
//
// This is an alias to an internal type.
type ReplaceSubscriptionPriceParamsPriceNewSubscriptionGroupedTieredPackagePriceBillingCycleConfigurationDurationUnit = shared.ReplaceSubscriptionPriceParamsPriceNewSubscriptionGroupedTieredPackagePriceBillingCycleConfigurationDurationUnit

// This is an alias to an internal value.
const ReplaceSubscriptionPriceParamsPriceNewSubscriptionGroupedTieredPackagePriceBillingCycleConfigurationDurationUnitDay = shared.ReplaceSubscriptionPriceParamsPriceNewSubscriptionGroupedTieredPackagePriceBillingCycleConfigurationDurationUnitDay

// This is an alias to an internal value.
const ReplaceSubscriptionPriceParamsPriceNewSubscriptionGroupedTieredPackagePriceBillingCycleConfigurationDurationUnitMonth = shared.ReplaceSubscriptionPriceParamsPriceNewSubscriptionGroupedTieredPackagePriceBillingCycleConfigurationDurationUnitMonth

// Within each billing cycle, specifies the cadence at which invoices are produced.
// If unspecified, a single invoice is produced per billing cycle.
//
// This is an alias to an internal type.
type ReplaceSubscriptionPriceParamsPriceNewSubscriptionGroupedTieredPackagePriceInvoicingCycleConfiguration = shared.ReplaceSubscriptionPriceParamsPriceNewSubscriptionGroupedTieredPackagePriceInvoicingCycleConfiguration

// The unit of billing period duration.
//
// This is an alias to an internal type.
type ReplaceSubscriptionPriceParamsPriceNewSubscriptionGroupedTieredPackagePriceInvoicingCycleConfigurationDurationUnit = shared.ReplaceSubscriptionPriceParamsPriceNewSubscriptionGroupedTieredPackagePriceInvoicingCycleConfigurationDurationUnit

// This is an alias to an internal value.
const ReplaceSubscriptionPriceParamsPriceNewSubscriptionGroupedTieredPackagePriceInvoicingCycleConfigurationDurationUnitDay = shared.ReplaceSubscriptionPriceParamsPriceNewSubscriptionGroupedTieredPackagePriceInvoicingCycleConfigurationDurationUnitDay

// This is an alias to an internal value.
const ReplaceSubscriptionPriceParamsPriceNewSubscriptionGroupedTieredPackagePriceInvoicingCycleConfigurationDurationUnitMonth = shared.ReplaceSubscriptionPriceParamsPriceNewSubscriptionGroupedTieredPackagePriceInvoicingCycleConfigurationDurationUnitMonth

// The cadence to bill for this price on.
//
// This is an alias to an internal type.
type ReplaceSubscriptionPriceParamsPriceCadence = shared.ReplaceSubscriptionPriceParamsPriceCadence

// This is an alias to an internal value.
const ReplaceSubscriptionPriceParamsPriceCadenceAnnual = shared.ReplaceSubscriptionPriceParamsPriceCadenceAnnual

// This is an alias to an internal value.
const ReplaceSubscriptionPriceParamsPriceCadenceSemiAnnual = shared.ReplaceSubscriptionPriceParamsPriceCadenceSemiAnnual

// This is an alias to an internal value.
const ReplaceSubscriptionPriceParamsPriceCadenceMonthly = shared.ReplaceSubscriptionPriceParamsPriceCadenceMonthly

// This is an alias to an internal value.
const ReplaceSubscriptionPriceParamsPriceCadenceQuarterly = shared.ReplaceSubscriptionPriceParamsPriceCadenceQuarterly

// This is an alias to an internal value.
const ReplaceSubscriptionPriceParamsPriceCadenceOneTime = shared.ReplaceSubscriptionPriceParamsPriceCadenceOneTime

// This is an alias to an internal value.
const ReplaceSubscriptionPriceParamsPriceCadenceCustom = shared.ReplaceSubscriptionPriceParamsPriceCadenceCustom

// This is an alias to an internal type.
type ReplaceSubscriptionPriceParamsPriceModelType = shared.ReplaceSubscriptionPriceParamsPriceModelType

// This is an alias to an internal value.
const ReplaceSubscriptionPriceParamsPriceModelTypeUnit = shared.ReplaceSubscriptionPriceParamsPriceModelTypeUnit

// This is an alias to an internal value.
const ReplaceSubscriptionPriceParamsPriceModelTypePackage = shared.ReplaceSubscriptionPriceParamsPriceModelTypePackage

// This is an alias to an internal value.
const ReplaceSubscriptionPriceParamsPriceModelTypeMatrix = shared.ReplaceSubscriptionPriceParamsPriceModelTypeMatrix

// This is an alias to an internal value.
const ReplaceSubscriptionPriceParamsPriceModelTypeTiered = shared.ReplaceSubscriptionPriceParamsPriceModelTypeTiered

// This is an alias to an internal value.
const ReplaceSubscriptionPriceParamsPriceModelTypeTieredBps = shared.ReplaceSubscriptionPriceParamsPriceModelTypeTieredBps

// This is an alias to an internal value.
const ReplaceSubscriptionPriceParamsPriceModelTypeBps = shared.ReplaceSubscriptionPriceParamsPriceModelTypeBps

// This is an alias to an internal value.
const ReplaceSubscriptionPriceParamsPriceModelTypeBulkBps = shared.ReplaceSubscriptionPriceParamsPriceModelTypeBulkBps

// This is an alias to an internal value.
const ReplaceSubscriptionPriceParamsPriceModelTypeBulk = shared.ReplaceSubscriptionPriceParamsPriceModelTypeBulk

// This is an alias to an internal value.
const ReplaceSubscriptionPriceParamsPriceModelTypeThresholdTotalAmount = shared.ReplaceSubscriptionPriceParamsPriceModelTypeThresholdTotalAmount

// This is an alias to an internal value.
const ReplaceSubscriptionPriceParamsPriceModelTypeTieredPackage = shared.ReplaceSubscriptionPriceParamsPriceModelTypeTieredPackage

// This is an alias to an internal value.
const ReplaceSubscriptionPriceParamsPriceModelTypeTieredWithMinimum = shared.ReplaceSubscriptionPriceParamsPriceModelTypeTieredWithMinimum

// This is an alias to an internal value.
const ReplaceSubscriptionPriceParamsPriceModelTypeUnitWithPercent = shared.ReplaceSubscriptionPriceParamsPriceModelTypeUnitWithPercent

// This is an alias to an internal value.
const ReplaceSubscriptionPriceParamsPriceModelTypePackageWithAllocation = shared.ReplaceSubscriptionPriceParamsPriceModelTypePackageWithAllocation

// This is an alias to an internal value.
const ReplaceSubscriptionPriceParamsPriceModelTypeTieredWithProration = shared.ReplaceSubscriptionPriceParamsPriceModelTypeTieredWithProration

// This is an alias to an internal value.
const ReplaceSubscriptionPriceParamsPriceModelTypeUnitWithProration = shared.ReplaceSubscriptionPriceParamsPriceModelTypeUnitWithProration

// This is an alias to an internal value.
const ReplaceSubscriptionPriceParamsPriceModelTypeGroupedAllocation = shared.ReplaceSubscriptionPriceParamsPriceModelTypeGroupedAllocation

// This is an alias to an internal value.
const ReplaceSubscriptionPriceParamsPriceModelTypeGroupedWithProratedMinimum = shared.ReplaceSubscriptionPriceParamsPriceModelTypeGroupedWithProratedMinimum

// This is an alias to an internal value.
const ReplaceSubscriptionPriceParamsPriceModelTypeBulkWithProration = shared.ReplaceSubscriptionPriceParamsPriceModelTypeBulkWithProration

// This is an alias to an internal value.
const ReplaceSubscriptionPriceParamsPriceModelTypeScalableMatrixWithUnitPricing = shared.ReplaceSubscriptionPriceParamsPriceModelTypeScalableMatrixWithUnitPricing

// This is an alias to an internal value.
const ReplaceSubscriptionPriceParamsPriceModelTypeScalableMatrixWithTieredPricing = shared.ReplaceSubscriptionPriceParamsPriceModelTypeScalableMatrixWithTieredPricing

// This is an alias to an internal value.
const ReplaceSubscriptionPriceParamsPriceModelTypeCumulativeGroupedBulk = shared.ReplaceSubscriptionPriceParamsPriceModelTypeCumulativeGroupedBulk

// This is an alias to an internal value.
const ReplaceSubscriptionPriceParamsPriceModelTypeMaxGroupTieredPackage = shared.ReplaceSubscriptionPriceParamsPriceModelTypeMaxGroupTieredPackage

// This is an alias to an internal value.
const ReplaceSubscriptionPriceParamsPriceModelTypeGroupedWithMeteredMinimum = shared.ReplaceSubscriptionPriceParamsPriceModelTypeGroupedWithMeteredMinimum

// This is an alias to an internal value.
const ReplaceSubscriptionPriceParamsPriceModelTypeMatrixWithDisplayName = shared.ReplaceSubscriptionPriceParamsPriceModelTypeMatrixWithDisplayName

// This is an alias to an internal value.
const ReplaceSubscriptionPriceParamsPriceModelTypeGroupedTieredPackage = shared.ReplaceSubscriptionPriceParamsPriceModelTypeGroupedTieredPackage

// This is an alias to an internal type.
type TrialDiscount = shared.TrialDiscount

// This is an alias to an internal type.
type TrialDiscountDiscountType = shared.TrialDiscountDiscountType

// This is an alias to an internal value.
const TrialDiscountDiscountTypeTrial = shared.TrialDiscountDiscountTypeTrial

// This is an alias to an internal type.
type TrialDiscountParam = shared.TrialDiscountParam

// This is an alias to an internal type.
type UsageDiscount = shared.UsageDiscount

// This is an alias to an internal type.
type UsageDiscountDiscountType = shared.UsageDiscountDiscountType

// This is an alias to an internal value.
const UsageDiscountDiscountTypeUsage = shared.UsageDiscountDiscountTypeUsage

// This is an alias to an internal type.
type UsageDiscountParam = shared.UsageDiscountParam
