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
type AmountDiscount = shared.AmountDiscount

// This is an alias to an internal type.
type AmountDiscountDiscountType = shared.AmountDiscountDiscountType

// This is an alias to an internal value.
const AmountDiscountDiscountTypeAmount = shared.AmountDiscountDiscountTypeAmount

// This is an alias to an internal type.
type AmountDiscountParam = shared.AmountDiscountParam

// This is an alias to an internal type.
type AmountDiscountInterval = shared.AmountDiscountInterval

// This is an alias to an internal type.
type AmountDiscountIntervalDiscountType = shared.AmountDiscountIntervalDiscountType

// This is an alias to an internal value.
const AmountDiscountIntervalDiscountTypeAmount = shared.AmountDiscountIntervalDiscountTypeAmount

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

// This is an alias to an internal type.
type BPSConfig = shared.BPSConfig

// This is an alias to an internal type.
type BPSConfigParam = shared.BPSConfigParam

// This is an alias to an internal type.
type BPSTier = shared.BPSTier

// This is an alias to an internal type.
type BPSTierParam = shared.BPSTierParam

// This is an alias to an internal type.
type BulkBPSConfig = shared.BulkBPSConfig

// This is an alias to an internal type.
type BulkBPSConfigParam = shared.BulkBPSConfigParam

// This is an alias to an internal type.
type BulkBPSTier = shared.BulkBPSTier

// This is an alias to an internal type.
type BulkBPSTierParam = shared.BulkBPSTierParam

// This is an alias to an internal type.
type BulkConfig = shared.BulkConfig

// This is an alias to an internal type.
type BulkConfigParam = shared.BulkConfigParam

// This is an alias to an internal type.
type BulkTier = shared.BulkTier

// This is an alias to an internal type.
type BulkTierParam = shared.BulkTierParam

// This is an alias to an internal type.
type ChangedSubscriptionResources = shared.ChangedSubscriptionResources

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

// This is an alias to an internal type.
type ItemSlim = shared.ItemSlim

// This is an alias to an internal type.
type MatrixConfig = shared.MatrixConfig

// This is an alias to an internal type.
type MatrixConfigParam = shared.MatrixConfigParam

// This is an alias to an internal type.
type MatrixSubLineItem = shared.MatrixSubLineItem

// This is an alias to an internal type.
type MatrixSubLineItemType = shared.MatrixSubLineItemType

// This is an alias to an internal value.
const MatrixSubLineItemTypeMatrix = shared.MatrixSubLineItemTypeMatrix

// This is an alias to an internal type.
type MatrixValue = shared.MatrixValue

// This is an alias to an internal type.
type MatrixValueParam = shared.MatrixValueParam

// This is an alias to an internal type.
type MatrixWithAllocationConfig = shared.MatrixWithAllocationConfig

// This is an alias to an internal type.
type MatrixWithAllocationConfigParam = shared.MatrixWithAllocationConfigParam

// This is an alias to an internal type.
type Maximum = shared.Maximum

// This is an alias to an internal type.
type MaximumInterval = shared.MaximumInterval

// This is an alias to an internal type.
type Minimum = shared.Minimum

// This is an alias to an internal type.
type MinimumInterval = shared.MinimumInterval

// This is an alias to an internal type.
type MonetaryAmountDiscountAdjustment = shared.MonetaryAmountDiscountAdjustment

// This is an alias to an internal type.
type MonetaryAmountDiscountAdjustmentAdjustmentType = shared.MonetaryAmountDiscountAdjustmentAdjustmentType

// This is an alias to an internal value.
const MonetaryAmountDiscountAdjustmentAdjustmentTypeAmountDiscount = shared.MonetaryAmountDiscountAdjustmentAdjustmentTypeAmountDiscount

// This is an alias to an internal type.
type MonetaryMaximumAdjustment = shared.MonetaryMaximumAdjustment

// This is an alias to an internal type.
type MonetaryMaximumAdjustmentAdjustmentType = shared.MonetaryMaximumAdjustmentAdjustmentType

// This is an alias to an internal value.
const MonetaryMaximumAdjustmentAdjustmentTypeMaximum = shared.MonetaryMaximumAdjustmentAdjustmentTypeMaximum

// This is an alias to an internal type.
type MonetaryMinimumAdjustment = shared.MonetaryMinimumAdjustment

// This is an alias to an internal type.
type MonetaryMinimumAdjustmentAdjustmentType = shared.MonetaryMinimumAdjustmentAdjustmentType

// This is an alias to an internal value.
const MonetaryMinimumAdjustmentAdjustmentTypeMinimum = shared.MonetaryMinimumAdjustmentAdjustmentTypeMinimum

// This is an alias to an internal type.
type MonetaryPercentageDiscountAdjustment = shared.MonetaryPercentageDiscountAdjustment

// This is an alias to an internal type.
type MonetaryPercentageDiscountAdjustmentAdjustmentType = shared.MonetaryPercentageDiscountAdjustmentAdjustmentType

// This is an alias to an internal value.
const MonetaryPercentageDiscountAdjustmentAdjustmentTypePercentageDiscount = shared.MonetaryPercentageDiscountAdjustmentAdjustmentTypePercentageDiscount

// This is an alias to an internal type.
type MonetaryUsageDiscountAdjustment = shared.MonetaryUsageDiscountAdjustment

// This is an alias to an internal type.
type MonetaryUsageDiscountAdjustmentAdjustmentType = shared.MonetaryUsageDiscountAdjustmentAdjustmentType

// This is an alias to an internal value.
const MonetaryUsageDiscountAdjustmentAdjustmentTypeUsageDiscount = shared.MonetaryUsageDiscountAdjustmentAdjustmentTypeUsageDiscount

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
type NewFloatingBPSPriceParam = shared.NewFloatingBPSPriceParam

// The cadence to bill for this price on.
//
// This is an alias to an internal type.
type NewFloatingBPSPriceCadence = shared.NewFloatingBPSPriceCadence

// This is an alias to an internal value.
const NewFloatingBPSPriceCadenceAnnual = shared.NewFloatingBPSPriceCadenceAnnual

// This is an alias to an internal value.
const NewFloatingBPSPriceCadenceSemiAnnual = shared.NewFloatingBPSPriceCadenceSemiAnnual

// This is an alias to an internal value.
const NewFloatingBPSPriceCadenceMonthly = shared.NewFloatingBPSPriceCadenceMonthly

// This is an alias to an internal value.
const NewFloatingBPSPriceCadenceQuarterly = shared.NewFloatingBPSPriceCadenceQuarterly

// This is an alias to an internal value.
const NewFloatingBPSPriceCadenceOneTime = shared.NewFloatingBPSPriceCadenceOneTime

// This is an alias to an internal value.
const NewFloatingBPSPriceCadenceCustom = shared.NewFloatingBPSPriceCadenceCustom

// This is an alias to an internal type.
type NewFloatingBPSPriceModelType = shared.NewFloatingBPSPriceModelType

// This is an alias to an internal value.
const NewFloatingBPSPriceModelTypeBPS = shared.NewFloatingBPSPriceModelTypeBPS

// The configuration for the rate of the price currency to the invoicing currency.
//
// This is an alias to an internal type.
type NewFloatingBPSPriceConversionRateConfigUnionParam = shared.NewFloatingBPSPriceConversionRateConfigUnionParam

// This is an alias to an internal type.
type NewFloatingBPSPriceConversionRateConfigConversionRateType = shared.NewFloatingBPSPriceConversionRateConfigConversionRateType

// This is an alias to an internal value.
const NewFloatingBPSPriceConversionRateConfigConversionRateTypeUnit = shared.NewFloatingBPSPriceConversionRateConfigConversionRateTypeUnit

// This is an alias to an internal value.
const NewFloatingBPSPriceConversionRateConfigConversionRateTypeTiered = shared.NewFloatingBPSPriceConversionRateConfigConversionRateTypeTiered

// This is an alias to an internal type.
type NewFloatingBulkBPSPriceParam = shared.NewFloatingBulkBPSPriceParam

// The cadence to bill for this price on.
//
// This is an alias to an internal type.
type NewFloatingBulkBPSPriceCadence = shared.NewFloatingBulkBPSPriceCadence

// This is an alias to an internal value.
const NewFloatingBulkBPSPriceCadenceAnnual = shared.NewFloatingBulkBPSPriceCadenceAnnual

// This is an alias to an internal value.
const NewFloatingBulkBPSPriceCadenceSemiAnnual = shared.NewFloatingBulkBPSPriceCadenceSemiAnnual

// This is an alias to an internal value.
const NewFloatingBulkBPSPriceCadenceMonthly = shared.NewFloatingBulkBPSPriceCadenceMonthly

// This is an alias to an internal value.
const NewFloatingBulkBPSPriceCadenceQuarterly = shared.NewFloatingBulkBPSPriceCadenceQuarterly

// This is an alias to an internal value.
const NewFloatingBulkBPSPriceCadenceOneTime = shared.NewFloatingBulkBPSPriceCadenceOneTime

// This is an alias to an internal value.
const NewFloatingBulkBPSPriceCadenceCustom = shared.NewFloatingBulkBPSPriceCadenceCustom

// This is an alias to an internal type.
type NewFloatingBulkBPSPriceModelType = shared.NewFloatingBulkBPSPriceModelType

// This is an alias to an internal value.
const NewFloatingBulkBPSPriceModelTypeBulkBPS = shared.NewFloatingBulkBPSPriceModelTypeBulkBPS

// The configuration for the rate of the price currency to the invoicing currency.
//
// This is an alias to an internal type.
type NewFloatingBulkBPSPriceConversionRateConfigUnionParam = shared.NewFloatingBulkBPSPriceConversionRateConfigUnionParam

// This is an alias to an internal type.
type NewFloatingBulkBPSPriceConversionRateConfigConversionRateType = shared.NewFloatingBulkBPSPriceConversionRateConfigConversionRateType

// This is an alias to an internal value.
const NewFloatingBulkBPSPriceConversionRateConfigConversionRateTypeUnit = shared.NewFloatingBulkBPSPriceConversionRateConfigConversionRateTypeUnit

// This is an alias to an internal value.
const NewFloatingBulkBPSPriceConversionRateConfigConversionRateTypeTiered = shared.NewFloatingBulkBPSPriceConversionRateConfigConversionRateTypeTiered

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

// This is an alias to an internal type.
type NewFloatingBulkPriceModelType = shared.NewFloatingBulkPriceModelType

// This is an alias to an internal value.
const NewFloatingBulkPriceModelTypeBulk = shared.NewFloatingBulkPriceModelTypeBulk

// The configuration for the rate of the price currency to the invoicing currency.
//
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

// This is an alias to an internal type.
type NewFloatingBulkWithProrationPriceModelType = shared.NewFloatingBulkWithProrationPriceModelType

// This is an alias to an internal value.
const NewFloatingBulkWithProrationPriceModelTypeBulkWithProration = shared.NewFloatingBulkWithProrationPriceModelTypeBulkWithProration

// The configuration for the rate of the price currency to the invoicing currency.
//
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

// This is an alias to an internal type.
type NewFloatingCumulativeGroupedBulkPriceModelType = shared.NewFloatingCumulativeGroupedBulkPriceModelType

// This is an alias to an internal value.
const NewFloatingCumulativeGroupedBulkPriceModelTypeCumulativeGroupedBulk = shared.NewFloatingCumulativeGroupedBulkPriceModelTypeCumulativeGroupedBulk

// The configuration for the rate of the price currency to the invoicing currency.
//
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

// This is an alias to an internal type.
type NewFloatingGroupedAllocationPriceModelType = shared.NewFloatingGroupedAllocationPriceModelType

// This is an alias to an internal value.
const NewFloatingGroupedAllocationPriceModelTypeGroupedAllocation = shared.NewFloatingGroupedAllocationPriceModelTypeGroupedAllocation

// The configuration for the rate of the price currency to the invoicing currency.
//
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

// This is an alias to an internal type.
type NewFloatingGroupedTieredPackagePriceModelType = shared.NewFloatingGroupedTieredPackagePriceModelType

// This is an alias to an internal value.
const NewFloatingGroupedTieredPackagePriceModelTypeGroupedTieredPackage = shared.NewFloatingGroupedTieredPackagePriceModelTypeGroupedTieredPackage

// The configuration for the rate of the price currency to the invoicing currency.
//
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

// This is an alias to an internal type.
type NewFloatingGroupedTieredPriceModelType = shared.NewFloatingGroupedTieredPriceModelType

// This is an alias to an internal value.
const NewFloatingGroupedTieredPriceModelTypeGroupedTiered = shared.NewFloatingGroupedTieredPriceModelTypeGroupedTiered

// The configuration for the rate of the price currency to the invoicing currency.
//
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

// This is an alias to an internal type.
type NewFloatingGroupedWithMeteredMinimumPriceModelType = shared.NewFloatingGroupedWithMeteredMinimumPriceModelType

// This is an alias to an internal value.
const NewFloatingGroupedWithMeteredMinimumPriceModelTypeGroupedWithMeteredMinimum = shared.NewFloatingGroupedWithMeteredMinimumPriceModelTypeGroupedWithMeteredMinimum

// The configuration for the rate of the price currency to the invoicing currency.
//
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

// This is an alias to an internal type.
type NewFloatingGroupedWithProratedMinimumPriceModelType = shared.NewFloatingGroupedWithProratedMinimumPriceModelType

// This is an alias to an internal value.
const NewFloatingGroupedWithProratedMinimumPriceModelTypeGroupedWithProratedMinimum = shared.NewFloatingGroupedWithProratedMinimumPriceModelTypeGroupedWithProratedMinimum

// The configuration for the rate of the price currency to the invoicing currency.
//
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

// This is an alias to an internal type.
type NewFloatingMatrixPriceModelType = shared.NewFloatingMatrixPriceModelType

// This is an alias to an internal value.
const NewFloatingMatrixPriceModelTypeMatrix = shared.NewFloatingMatrixPriceModelTypeMatrix

// The configuration for the rate of the price currency to the invoicing currency.
//
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

// This is an alias to an internal type.
type NewFloatingMatrixWithAllocationPriceModelType = shared.NewFloatingMatrixWithAllocationPriceModelType

// This is an alias to an internal value.
const NewFloatingMatrixWithAllocationPriceModelTypeMatrixWithAllocation = shared.NewFloatingMatrixWithAllocationPriceModelTypeMatrixWithAllocation

// The configuration for the rate of the price currency to the invoicing currency.
//
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

// This is an alias to an internal type.
type NewFloatingMatrixWithDisplayNamePriceModelType = shared.NewFloatingMatrixWithDisplayNamePriceModelType

// This is an alias to an internal value.
const NewFloatingMatrixWithDisplayNamePriceModelTypeMatrixWithDisplayName = shared.NewFloatingMatrixWithDisplayNamePriceModelTypeMatrixWithDisplayName

// The configuration for the rate of the price currency to the invoicing currency.
//
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

// This is an alias to an internal type.
type NewFloatingMaxGroupTieredPackagePriceModelType = shared.NewFloatingMaxGroupTieredPackagePriceModelType

// This is an alias to an internal value.
const NewFloatingMaxGroupTieredPackagePriceModelTypeMaxGroupTieredPackage = shared.NewFloatingMaxGroupTieredPackagePriceModelTypeMaxGroupTieredPackage

// The configuration for the rate of the price currency to the invoicing currency.
//
// This is an alias to an internal type.
type NewFloatingMaxGroupTieredPackagePriceConversionRateConfigUnionParam = shared.NewFloatingMaxGroupTieredPackagePriceConversionRateConfigUnionParam

// This is an alias to an internal type.
type NewFloatingMaxGroupTieredPackagePriceConversionRateConfigConversionRateType = shared.NewFloatingMaxGroupTieredPackagePriceConversionRateConfigConversionRateType

// This is an alias to an internal value.
const NewFloatingMaxGroupTieredPackagePriceConversionRateConfigConversionRateTypeUnit = shared.NewFloatingMaxGroupTieredPackagePriceConversionRateConfigConversionRateTypeUnit

// This is an alias to an internal value.
const NewFloatingMaxGroupTieredPackagePriceConversionRateConfigConversionRateTypeTiered = shared.NewFloatingMaxGroupTieredPackagePriceConversionRateConfigConversionRateTypeTiered

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

// This is an alias to an internal type.
type NewFloatingPackagePriceModelType = shared.NewFloatingPackagePriceModelType

// This is an alias to an internal value.
const NewFloatingPackagePriceModelTypePackage = shared.NewFloatingPackagePriceModelTypePackage

// The configuration for the rate of the price currency to the invoicing currency.
//
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

// This is an alias to an internal type.
type NewFloatingPackageWithAllocationPriceModelType = shared.NewFloatingPackageWithAllocationPriceModelType

// This is an alias to an internal value.
const NewFloatingPackageWithAllocationPriceModelTypePackageWithAllocation = shared.NewFloatingPackageWithAllocationPriceModelTypePackageWithAllocation

// The configuration for the rate of the price currency to the invoicing currency.
//
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

// This is an alias to an internal type.
type NewFloatingScalableMatrixWithTieredPricingPriceModelType = shared.NewFloatingScalableMatrixWithTieredPricingPriceModelType

// This is an alias to an internal value.
const NewFloatingScalableMatrixWithTieredPricingPriceModelTypeScalableMatrixWithTieredPricing = shared.NewFloatingScalableMatrixWithTieredPricingPriceModelTypeScalableMatrixWithTieredPricing

// The configuration for the rate of the price currency to the invoicing currency.
//
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

// This is an alias to an internal type.
type NewFloatingScalableMatrixWithUnitPricingPriceModelType = shared.NewFloatingScalableMatrixWithUnitPricingPriceModelType

// This is an alias to an internal value.
const NewFloatingScalableMatrixWithUnitPricingPriceModelTypeScalableMatrixWithUnitPricing = shared.NewFloatingScalableMatrixWithUnitPricingPriceModelTypeScalableMatrixWithUnitPricing

// The configuration for the rate of the price currency to the invoicing currency.
//
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

// This is an alias to an internal type.
type NewFloatingThresholdTotalAmountPriceModelType = shared.NewFloatingThresholdTotalAmountPriceModelType

// This is an alias to an internal value.
const NewFloatingThresholdTotalAmountPriceModelTypeThresholdTotalAmount = shared.NewFloatingThresholdTotalAmountPriceModelTypeThresholdTotalAmount

// The configuration for the rate of the price currency to the invoicing currency.
//
// This is an alias to an internal type.
type NewFloatingThresholdTotalAmountPriceConversionRateConfigUnionParam = shared.NewFloatingThresholdTotalAmountPriceConversionRateConfigUnionParam

// This is an alias to an internal type.
type NewFloatingThresholdTotalAmountPriceConversionRateConfigConversionRateType = shared.NewFloatingThresholdTotalAmountPriceConversionRateConfigConversionRateType

// This is an alias to an internal value.
const NewFloatingThresholdTotalAmountPriceConversionRateConfigConversionRateTypeUnit = shared.NewFloatingThresholdTotalAmountPriceConversionRateConfigConversionRateTypeUnit

// This is an alias to an internal value.
const NewFloatingThresholdTotalAmountPriceConversionRateConfigConversionRateTypeTiered = shared.NewFloatingThresholdTotalAmountPriceConversionRateConfigConversionRateTypeTiered

// This is an alias to an internal type.
type NewFloatingTieredBPSPriceParam = shared.NewFloatingTieredBPSPriceParam

// The cadence to bill for this price on.
//
// This is an alias to an internal type.
type NewFloatingTieredBPSPriceCadence = shared.NewFloatingTieredBPSPriceCadence

// This is an alias to an internal value.
const NewFloatingTieredBPSPriceCadenceAnnual = shared.NewFloatingTieredBPSPriceCadenceAnnual

// This is an alias to an internal value.
const NewFloatingTieredBPSPriceCadenceSemiAnnual = shared.NewFloatingTieredBPSPriceCadenceSemiAnnual

// This is an alias to an internal value.
const NewFloatingTieredBPSPriceCadenceMonthly = shared.NewFloatingTieredBPSPriceCadenceMonthly

// This is an alias to an internal value.
const NewFloatingTieredBPSPriceCadenceQuarterly = shared.NewFloatingTieredBPSPriceCadenceQuarterly

// This is an alias to an internal value.
const NewFloatingTieredBPSPriceCadenceOneTime = shared.NewFloatingTieredBPSPriceCadenceOneTime

// This is an alias to an internal value.
const NewFloatingTieredBPSPriceCadenceCustom = shared.NewFloatingTieredBPSPriceCadenceCustom

// This is an alias to an internal type.
type NewFloatingTieredBPSPriceModelType = shared.NewFloatingTieredBPSPriceModelType

// This is an alias to an internal value.
const NewFloatingTieredBPSPriceModelTypeTieredBPS = shared.NewFloatingTieredBPSPriceModelTypeTieredBPS

// The configuration for the rate of the price currency to the invoicing currency.
//
// This is an alias to an internal type.
type NewFloatingTieredBPSPriceConversionRateConfigUnionParam = shared.NewFloatingTieredBPSPriceConversionRateConfigUnionParam

// This is an alias to an internal type.
type NewFloatingTieredBPSPriceConversionRateConfigConversionRateType = shared.NewFloatingTieredBPSPriceConversionRateConfigConversionRateType

// This is an alias to an internal value.
const NewFloatingTieredBPSPriceConversionRateConfigConversionRateTypeUnit = shared.NewFloatingTieredBPSPriceConversionRateConfigConversionRateTypeUnit

// This is an alias to an internal value.
const NewFloatingTieredBPSPriceConversionRateConfigConversionRateTypeTiered = shared.NewFloatingTieredBPSPriceConversionRateConfigConversionRateTypeTiered

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

// This is an alias to an internal type.
type NewFloatingTieredPackagePriceModelType = shared.NewFloatingTieredPackagePriceModelType

// This is an alias to an internal value.
const NewFloatingTieredPackagePriceModelTypeTieredPackage = shared.NewFloatingTieredPackagePriceModelTypeTieredPackage

// The configuration for the rate of the price currency to the invoicing currency.
//
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

// This is an alias to an internal type.
type NewFloatingTieredPackageWithMinimumPriceModelType = shared.NewFloatingTieredPackageWithMinimumPriceModelType

// This is an alias to an internal value.
const NewFloatingTieredPackageWithMinimumPriceModelTypeTieredPackageWithMinimum = shared.NewFloatingTieredPackageWithMinimumPriceModelTypeTieredPackageWithMinimum

// The configuration for the rate of the price currency to the invoicing currency.
//
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

// This is an alias to an internal type.
type NewFloatingTieredPriceModelType = shared.NewFloatingTieredPriceModelType

// This is an alias to an internal value.
const NewFloatingTieredPriceModelTypeTiered = shared.NewFloatingTieredPriceModelTypeTiered

// The configuration for the rate of the price currency to the invoicing currency.
//
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

// This is an alias to an internal type.
type NewFloatingTieredWithMinimumPriceModelType = shared.NewFloatingTieredWithMinimumPriceModelType

// This is an alias to an internal value.
const NewFloatingTieredWithMinimumPriceModelTypeTieredWithMinimum = shared.NewFloatingTieredWithMinimumPriceModelTypeTieredWithMinimum

// The configuration for the rate of the price currency to the invoicing currency.
//
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

// This is an alias to an internal type.
type NewFloatingTieredWithProrationPriceModelType = shared.NewFloatingTieredWithProrationPriceModelType

// This is an alias to an internal value.
const NewFloatingTieredWithProrationPriceModelTypeTieredWithProration = shared.NewFloatingTieredWithProrationPriceModelTypeTieredWithProration

// The configuration for the rate of the price currency to the invoicing currency.
//
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

// This is an alias to an internal type.
type NewFloatingUnitPriceModelType = shared.NewFloatingUnitPriceModelType

// This is an alias to an internal value.
const NewFloatingUnitPriceModelTypeUnit = shared.NewFloatingUnitPriceModelTypeUnit

// The configuration for the rate of the price currency to the invoicing currency.
//
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

// This is an alias to an internal type.
type NewFloatingUnitWithPercentPriceModelType = shared.NewFloatingUnitWithPercentPriceModelType

// This is an alias to an internal value.
const NewFloatingUnitWithPercentPriceModelTypeUnitWithPercent = shared.NewFloatingUnitWithPercentPriceModelTypeUnitWithPercent

// The configuration for the rate of the price currency to the invoicing currency.
//
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

// This is an alias to an internal type.
type NewFloatingUnitWithProrationPriceModelType = shared.NewFloatingUnitWithProrationPriceModelType

// This is an alias to an internal value.
const NewFloatingUnitWithProrationPriceModelTypeUnitWithProration = shared.NewFloatingUnitWithProrationPriceModelTypeUnitWithProration

// The configuration for the rate of the price currency to the invoicing currency.
//
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
type NewPlanBPSPriceParam = shared.NewPlanBPSPriceParam

// The cadence to bill for this price on.
//
// This is an alias to an internal type.
type NewPlanBPSPriceCadence = shared.NewPlanBPSPriceCadence

// This is an alias to an internal value.
const NewPlanBPSPriceCadenceAnnual = shared.NewPlanBPSPriceCadenceAnnual

// This is an alias to an internal value.
const NewPlanBPSPriceCadenceSemiAnnual = shared.NewPlanBPSPriceCadenceSemiAnnual

// This is an alias to an internal value.
const NewPlanBPSPriceCadenceMonthly = shared.NewPlanBPSPriceCadenceMonthly

// This is an alias to an internal value.
const NewPlanBPSPriceCadenceQuarterly = shared.NewPlanBPSPriceCadenceQuarterly

// This is an alias to an internal value.
const NewPlanBPSPriceCadenceOneTime = shared.NewPlanBPSPriceCadenceOneTime

// This is an alias to an internal value.
const NewPlanBPSPriceCadenceCustom = shared.NewPlanBPSPriceCadenceCustom

// This is an alias to an internal type.
type NewPlanBPSPriceModelType = shared.NewPlanBPSPriceModelType

// This is an alias to an internal value.
const NewPlanBPSPriceModelTypeBPS = shared.NewPlanBPSPriceModelTypeBPS

// The configuration for the rate of the price currency to the invoicing currency.
//
// This is an alias to an internal type.
type NewPlanBPSPriceConversionRateConfigUnionParam = shared.NewPlanBPSPriceConversionRateConfigUnionParam

// This is an alias to an internal type.
type NewPlanBPSPriceConversionRateConfigConversionRateType = shared.NewPlanBPSPriceConversionRateConfigConversionRateType

// This is an alias to an internal value.
const NewPlanBPSPriceConversionRateConfigConversionRateTypeUnit = shared.NewPlanBPSPriceConversionRateConfigConversionRateTypeUnit

// This is an alias to an internal value.
const NewPlanBPSPriceConversionRateConfigConversionRateTypeTiered = shared.NewPlanBPSPriceConversionRateConfigConversionRateTypeTiered

// This is an alias to an internal type.
type NewPlanBulkBPSPriceParam = shared.NewPlanBulkBPSPriceParam

// The cadence to bill for this price on.
//
// This is an alias to an internal type.
type NewPlanBulkBPSPriceCadence = shared.NewPlanBulkBPSPriceCadence

// This is an alias to an internal value.
const NewPlanBulkBPSPriceCadenceAnnual = shared.NewPlanBulkBPSPriceCadenceAnnual

// This is an alias to an internal value.
const NewPlanBulkBPSPriceCadenceSemiAnnual = shared.NewPlanBulkBPSPriceCadenceSemiAnnual

// This is an alias to an internal value.
const NewPlanBulkBPSPriceCadenceMonthly = shared.NewPlanBulkBPSPriceCadenceMonthly

// This is an alias to an internal value.
const NewPlanBulkBPSPriceCadenceQuarterly = shared.NewPlanBulkBPSPriceCadenceQuarterly

// This is an alias to an internal value.
const NewPlanBulkBPSPriceCadenceOneTime = shared.NewPlanBulkBPSPriceCadenceOneTime

// This is an alias to an internal value.
const NewPlanBulkBPSPriceCadenceCustom = shared.NewPlanBulkBPSPriceCadenceCustom

// This is an alias to an internal type.
type NewPlanBulkBPSPriceModelType = shared.NewPlanBulkBPSPriceModelType

// This is an alias to an internal value.
const NewPlanBulkBPSPriceModelTypeBulkBPS = shared.NewPlanBulkBPSPriceModelTypeBulkBPS

// The configuration for the rate of the price currency to the invoicing currency.
//
// This is an alias to an internal type.
type NewPlanBulkBPSPriceConversionRateConfigUnionParam = shared.NewPlanBulkBPSPriceConversionRateConfigUnionParam

// This is an alias to an internal type.
type NewPlanBulkBPSPriceConversionRateConfigConversionRateType = shared.NewPlanBulkBPSPriceConversionRateConfigConversionRateType

// This is an alias to an internal value.
const NewPlanBulkBPSPriceConversionRateConfigConversionRateTypeUnit = shared.NewPlanBulkBPSPriceConversionRateConfigConversionRateTypeUnit

// This is an alias to an internal value.
const NewPlanBulkBPSPriceConversionRateConfigConversionRateTypeTiered = shared.NewPlanBulkBPSPriceConversionRateConfigConversionRateTypeTiered

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

// This is an alias to an internal type.
type NewPlanBulkPriceModelType = shared.NewPlanBulkPriceModelType

// This is an alias to an internal value.
const NewPlanBulkPriceModelTypeBulk = shared.NewPlanBulkPriceModelTypeBulk

// The configuration for the rate of the price currency to the invoicing currency.
//
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

// This is an alias to an internal type.
type NewPlanBulkWithProrationPriceModelType = shared.NewPlanBulkWithProrationPriceModelType

// This is an alias to an internal value.
const NewPlanBulkWithProrationPriceModelTypeBulkWithProration = shared.NewPlanBulkWithProrationPriceModelTypeBulkWithProration

// The configuration for the rate of the price currency to the invoicing currency.
//
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

// This is an alias to an internal type.
type NewPlanCumulativeGroupedBulkPriceModelType = shared.NewPlanCumulativeGroupedBulkPriceModelType

// This is an alias to an internal value.
const NewPlanCumulativeGroupedBulkPriceModelTypeCumulativeGroupedBulk = shared.NewPlanCumulativeGroupedBulkPriceModelTypeCumulativeGroupedBulk

// The configuration for the rate of the price currency to the invoicing currency.
//
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

// This is an alias to an internal type.
type NewPlanGroupedAllocationPriceModelType = shared.NewPlanGroupedAllocationPriceModelType

// This is an alias to an internal value.
const NewPlanGroupedAllocationPriceModelTypeGroupedAllocation = shared.NewPlanGroupedAllocationPriceModelTypeGroupedAllocation

// The configuration for the rate of the price currency to the invoicing currency.
//
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

// This is an alias to an internal type.
type NewPlanGroupedTieredPackagePriceModelType = shared.NewPlanGroupedTieredPackagePriceModelType

// This is an alias to an internal value.
const NewPlanGroupedTieredPackagePriceModelTypeGroupedTieredPackage = shared.NewPlanGroupedTieredPackagePriceModelTypeGroupedTieredPackage

// The configuration for the rate of the price currency to the invoicing currency.
//
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

// This is an alias to an internal type.
type NewPlanGroupedTieredPriceModelType = shared.NewPlanGroupedTieredPriceModelType

// This is an alias to an internal value.
const NewPlanGroupedTieredPriceModelTypeGroupedTiered = shared.NewPlanGroupedTieredPriceModelTypeGroupedTiered

// The configuration for the rate of the price currency to the invoicing currency.
//
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

// This is an alias to an internal type.
type NewPlanGroupedWithMeteredMinimumPriceModelType = shared.NewPlanGroupedWithMeteredMinimumPriceModelType

// This is an alias to an internal value.
const NewPlanGroupedWithMeteredMinimumPriceModelTypeGroupedWithMeteredMinimum = shared.NewPlanGroupedWithMeteredMinimumPriceModelTypeGroupedWithMeteredMinimum

// The configuration for the rate of the price currency to the invoicing currency.
//
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

// This is an alias to an internal type.
type NewPlanGroupedWithProratedMinimumPriceModelType = shared.NewPlanGroupedWithProratedMinimumPriceModelType

// This is an alias to an internal value.
const NewPlanGroupedWithProratedMinimumPriceModelTypeGroupedWithProratedMinimum = shared.NewPlanGroupedWithProratedMinimumPriceModelTypeGroupedWithProratedMinimum

// The configuration for the rate of the price currency to the invoicing currency.
//
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

// This is an alias to an internal type.
type NewPlanMatrixPriceModelType = shared.NewPlanMatrixPriceModelType

// This is an alias to an internal value.
const NewPlanMatrixPriceModelTypeMatrix = shared.NewPlanMatrixPriceModelTypeMatrix

// The configuration for the rate of the price currency to the invoicing currency.
//
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

// This is an alias to an internal type.
type NewPlanMatrixWithAllocationPriceModelType = shared.NewPlanMatrixWithAllocationPriceModelType

// This is an alias to an internal value.
const NewPlanMatrixWithAllocationPriceModelTypeMatrixWithAllocation = shared.NewPlanMatrixWithAllocationPriceModelTypeMatrixWithAllocation

// The configuration for the rate of the price currency to the invoicing currency.
//
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

// This is an alias to an internal type.
type NewPlanMatrixWithDisplayNamePriceModelType = shared.NewPlanMatrixWithDisplayNamePriceModelType

// This is an alias to an internal value.
const NewPlanMatrixWithDisplayNamePriceModelTypeMatrixWithDisplayName = shared.NewPlanMatrixWithDisplayNamePriceModelTypeMatrixWithDisplayName

// The configuration for the rate of the price currency to the invoicing currency.
//
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

// This is an alias to an internal type.
type NewPlanMaxGroupTieredPackagePriceModelType = shared.NewPlanMaxGroupTieredPackagePriceModelType

// This is an alias to an internal value.
const NewPlanMaxGroupTieredPackagePriceModelTypeMaxGroupTieredPackage = shared.NewPlanMaxGroupTieredPackagePriceModelTypeMaxGroupTieredPackage

// The configuration for the rate of the price currency to the invoicing currency.
//
// This is an alias to an internal type.
type NewPlanMaxGroupTieredPackagePriceConversionRateConfigUnionParam = shared.NewPlanMaxGroupTieredPackagePriceConversionRateConfigUnionParam

// This is an alias to an internal type.
type NewPlanMaxGroupTieredPackagePriceConversionRateConfigConversionRateType = shared.NewPlanMaxGroupTieredPackagePriceConversionRateConfigConversionRateType

// This is an alias to an internal value.
const NewPlanMaxGroupTieredPackagePriceConversionRateConfigConversionRateTypeUnit = shared.NewPlanMaxGroupTieredPackagePriceConversionRateConfigConversionRateTypeUnit

// This is an alias to an internal value.
const NewPlanMaxGroupTieredPackagePriceConversionRateConfigConversionRateTypeTiered = shared.NewPlanMaxGroupTieredPackagePriceConversionRateConfigConversionRateTypeTiered

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

// This is an alias to an internal type.
type NewPlanPackagePriceModelType = shared.NewPlanPackagePriceModelType

// This is an alias to an internal value.
const NewPlanPackagePriceModelTypePackage = shared.NewPlanPackagePriceModelTypePackage

// The configuration for the rate of the price currency to the invoicing currency.
//
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

// This is an alias to an internal type.
type NewPlanPackageWithAllocationPriceModelType = shared.NewPlanPackageWithAllocationPriceModelType

// This is an alias to an internal value.
const NewPlanPackageWithAllocationPriceModelTypePackageWithAllocation = shared.NewPlanPackageWithAllocationPriceModelTypePackageWithAllocation

// The configuration for the rate of the price currency to the invoicing currency.
//
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

// This is an alias to an internal type.
type NewPlanScalableMatrixWithTieredPricingPriceModelType = shared.NewPlanScalableMatrixWithTieredPricingPriceModelType

// This is an alias to an internal value.
const NewPlanScalableMatrixWithTieredPricingPriceModelTypeScalableMatrixWithTieredPricing = shared.NewPlanScalableMatrixWithTieredPricingPriceModelTypeScalableMatrixWithTieredPricing

// The configuration for the rate of the price currency to the invoicing currency.
//
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

// This is an alias to an internal type.
type NewPlanScalableMatrixWithUnitPricingPriceModelType = shared.NewPlanScalableMatrixWithUnitPricingPriceModelType

// This is an alias to an internal value.
const NewPlanScalableMatrixWithUnitPricingPriceModelTypeScalableMatrixWithUnitPricing = shared.NewPlanScalableMatrixWithUnitPricingPriceModelTypeScalableMatrixWithUnitPricing

// The configuration for the rate of the price currency to the invoicing currency.
//
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

// This is an alias to an internal type.
type NewPlanThresholdTotalAmountPriceModelType = shared.NewPlanThresholdTotalAmountPriceModelType

// This is an alias to an internal value.
const NewPlanThresholdTotalAmountPriceModelTypeThresholdTotalAmount = shared.NewPlanThresholdTotalAmountPriceModelTypeThresholdTotalAmount

// The configuration for the rate of the price currency to the invoicing currency.
//
// This is an alias to an internal type.
type NewPlanThresholdTotalAmountPriceConversionRateConfigUnionParam = shared.NewPlanThresholdTotalAmountPriceConversionRateConfigUnionParam

// This is an alias to an internal type.
type NewPlanThresholdTotalAmountPriceConversionRateConfigConversionRateType = shared.NewPlanThresholdTotalAmountPriceConversionRateConfigConversionRateType

// This is an alias to an internal value.
const NewPlanThresholdTotalAmountPriceConversionRateConfigConversionRateTypeUnit = shared.NewPlanThresholdTotalAmountPriceConversionRateConfigConversionRateTypeUnit

// This is an alias to an internal value.
const NewPlanThresholdTotalAmountPriceConversionRateConfigConversionRateTypeTiered = shared.NewPlanThresholdTotalAmountPriceConversionRateConfigConversionRateTypeTiered

// This is an alias to an internal type.
type NewPlanTierWithProrationPriceParam = shared.NewPlanTierWithProrationPriceParam

// The cadence to bill for this price on.
//
// This is an alias to an internal type.
type NewPlanTierWithProrationPriceCadence = shared.NewPlanTierWithProrationPriceCadence

// This is an alias to an internal value.
const NewPlanTierWithProrationPriceCadenceAnnual = shared.NewPlanTierWithProrationPriceCadenceAnnual

// This is an alias to an internal value.
const NewPlanTierWithProrationPriceCadenceSemiAnnual = shared.NewPlanTierWithProrationPriceCadenceSemiAnnual

// This is an alias to an internal value.
const NewPlanTierWithProrationPriceCadenceMonthly = shared.NewPlanTierWithProrationPriceCadenceMonthly

// This is an alias to an internal value.
const NewPlanTierWithProrationPriceCadenceQuarterly = shared.NewPlanTierWithProrationPriceCadenceQuarterly

// This is an alias to an internal value.
const NewPlanTierWithProrationPriceCadenceOneTime = shared.NewPlanTierWithProrationPriceCadenceOneTime

// This is an alias to an internal value.
const NewPlanTierWithProrationPriceCadenceCustom = shared.NewPlanTierWithProrationPriceCadenceCustom

// This is an alias to an internal type.
type NewPlanTierWithProrationPriceModelType = shared.NewPlanTierWithProrationPriceModelType

// This is an alias to an internal value.
const NewPlanTierWithProrationPriceModelTypeTieredWithProration = shared.NewPlanTierWithProrationPriceModelTypeTieredWithProration

// The configuration for the rate of the price currency to the invoicing currency.
//
// This is an alias to an internal type.
type NewPlanTierWithProrationPriceConversionRateConfigUnionParam = shared.NewPlanTierWithProrationPriceConversionRateConfigUnionParam

// This is an alias to an internal type.
type NewPlanTierWithProrationPriceConversionRateConfigConversionRateType = shared.NewPlanTierWithProrationPriceConversionRateConfigConversionRateType

// This is an alias to an internal value.
const NewPlanTierWithProrationPriceConversionRateConfigConversionRateTypeUnit = shared.NewPlanTierWithProrationPriceConversionRateConfigConversionRateTypeUnit

// This is an alias to an internal value.
const NewPlanTierWithProrationPriceConversionRateConfigConversionRateTypeTiered = shared.NewPlanTierWithProrationPriceConversionRateConfigConversionRateTypeTiered

// This is an alias to an internal type.
type NewPlanTieredBPSPriceParam = shared.NewPlanTieredBPSPriceParam

// The cadence to bill for this price on.
//
// This is an alias to an internal type.
type NewPlanTieredBPSPriceCadence = shared.NewPlanTieredBPSPriceCadence

// This is an alias to an internal value.
const NewPlanTieredBPSPriceCadenceAnnual = shared.NewPlanTieredBPSPriceCadenceAnnual

// This is an alias to an internal value.
const NewPlanTieredBPSPriceCadenceSemiAnnual = shared.NewPlanTieredBPSPriceCadenceSemiAnnual

// This is an alias to an internal value.
const NewPlanTieredBPSPriceCadenceMonthly = shared.NewPlanTieredBPSPriceCadenceMonthly

// This is an alias to an internal value.
const NewPlanTieredBPSPriceCadenceQuarterly = shared.NewPlanTieredBPSPriceCadenceQuarterly

// This is an alias to an internal value.
const NewPlanTieredBPSPriceCadenceOneTime = shared.NewPlanTieredBPSPriceCadenceOneTime

// This is an alias to an internal value.
const NewPlanTieredBPSPriceCadenceCustom = shared.NewPlanTieredBPSPriceCadenceCustom

// This is an alias to an internal type.
type NewPlanTieredBPSPriceModelType = shared.NewPlanTieredBPSPriceModelType

// This is an alias to an internal value.
const NewPlanTieredBPSPriceModelTypeTieredBPS = shared.NewPlanTieredBPSPriceModelTypeTieredBPS

// The configuration for the rate of the price currency to the invoicing currency.
//
// This is an alias to an internal type.
type NewPlanTieredBPSPriceConversionRateConfigUnionParam = shared.NewPlanTieredBPSPriceConversionRateConfigUnionParam

// This is an alias to an internal type.
type NewPlanTieredBPSPriceConversionRateConfigConversionRateType = shared.NewPlanTieredBPSPriceConversionRateConfigConversionRateType

// This is an alias to an internal value.
const NewPlanTieredBPSPriceConversionRateConfigConversionRateTypeUnit = shared.NewPlanTieredBPSPriceConversionRateConfigConversionRateTypeUnit

// This is an alias to an internal value.
const NewPlanTieredBPSPriceConversionRateConfigConversionRateTypeTiered = shared.NewPlanTieredBPSPriceConversionRateConfigConversionRateTypeTiered

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

// This is an alias to an internal type.
type NewPlanTieredPackagePriceModelType = shared.NewPlanTieredPackagePriceModelType

// This is an alias to an internal value.
const NewPlanTieredPackagePriceModelTypeTieredPackage = shared.NewPlanTieredPackagePriceModelTypeTieredPackage

// The configuration for the rate of the price currency to the invoicing currency.
//
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

// This is an alias to an internal type.
type NewPlanTieredPackageWithMinimumPriceModelType = shared.NewPlanTieredPackageWithMinimumPriceModelType

// This is an alias to an internal value.
const NewPlanTieredPackageWithMinimumPriceModelTypeTieredPackageWithMinimum = shared.NewPlanTieredPackageWithMinimumPriceModelTypeTieredPackageWithMinimum

// The configuration for the rate of the price currency to the invoicing currency.
//
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

// This is an alias to an internal type.
type NewPlanTieredPriceModelType = shared.NewPlanTieredPriceModelType

// This is an alias to an internal value.
const NewPlanTieredPriceModelTypeTiered = shared.NewPlanTieredPriceModelTypeTiered

// The configuration for the rate of the price currency to the invoicing currency.
//
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

// This is an alias to an internal type.
type NewPlanTieredWithMinimumPriceModelType = shared.NewPlanTieredWithMinimumPriceModelType

// This is an alias to an internal value.
const NewPlanTieredWithMinimumPriceModelTypeTieredWithMinimum = shared.NewPlanTieredWithMinimumPriceModelTypeTieredWithMinimum

// The configuration for the rate of the price currency to the invoicing currency.
//
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

// This is an alias to an internal type.
type NewPlanUnitPriceModelType = shared.NewPlanUnitPriceModelType

// This is an alias to an internal value.
const NewPlanUnitPriceModelTypeUnit = shared.NewPlanUnitPriceModelTypeUnit

// The configuration for the rate of the price currency to the invoicing currency.
//
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

// This is an alias to an internal type.
type NewPlanUnitWithPercentPriceModelType = shared.NewPlanUnitWithPercentPriceModelType

// This is an alias to an internal value.
const NewPlanUnitWithPercentPriceModelTypeUnitWithPercent = shared.NewPlanUnitWithPercentPriceModelTypeUnitWithPercent

// The configuration for the rate of the price currency to the invoicing currency.
//
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

// This is an alias to an internal type.
type NewPlanUnitWithProrationPriceModelType = shared.NewPlanUnitWithProrationPriceModelType

// This is an alias to an internal value.
const NewPlanUnitWithProrationPriceModelTypeUnitWithProration = shared.NewPlanUnitWithProrationPriceModelTypeUnitWithProration

// The configuration for the rate of the price currency to the invoicing currency.
//
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

// This is an alias to an internal type.
type PackageConfig = shared.PackageConfig

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
type PercentageDiscountParam = shared.PercentageDiscountParam

// This is an alias to an internal type.
type PercentageDiscountInterval = shared.PercentageDiscountInterval

// This is an alias to an internal type.
type PercentageDiscountIntervalDiscountType = shared.PercentageDiscountIntervalDiscountType

// This is an alias to an internal value.
const PercentageDiscountIntervalDiscountTypePercentage = shared.PercentageDiscountIntervalDiscountTypePercentage

// This is an alias to an internal type.
type PlanPhaseAmountDiscountAdjustment = shared.PlanPhaseAmountDiscountAdjustment

// This is an alias to an internal type.
type PlanPhaseAmountDiscountAdjustmentAdjustmentType = shared.PlanPhaseAmountDiscountAdjustmentAdjustmentType

// This is an alias to an internal value.
const PlanPhaseAmountDiscountAdjustmentAdjustmentTypeAmountDiscount = shared.PlanPhaseAmountDiscountAdjustmentAdjustmentTypeAmountDiscount

// This is an alias to an internal type.
type PlanPhaseMaximumAdjustment = shared.PlanPhaseMaximumAdjustment

// This is an alias to an internal type.
type PlanPhaseMaximumAdjustmentAdjustmentType = shared.PlanPhaseMaximumAdjustmentAdjustmentType

// This is an alias to an internal value.
const PlanPhaseMaximumAdjustmentAdjustmentTypeMaximum = shared.PlanPhaseMaximumAdjustmentAdjustmentTypeMaximum

// This is an alias to an internal type.
type PlanPhaseMinimumAdjustment = shared.PlanPhaseMinimumAdjustment

// This is an alias to an internal type.
type PlanPhaseMinimumAdjustmentAdjustmentType = shared.PlanPhaseMinimumAdjustmentAdjustmentType

// This is an alias to an internal value.
const PlanPhaseMinimumAdjustmentAdjustmentTypeMinimum = shared.PlanPhaseMinimumAdjustmentAdjustmentTypeMinimum

// This is an alias to an internal type.
type PlanPhasePercentageDiscountAdjustment = shared.PlanPhasePercentageDiscountAdjustment

// This is an alias to an internal type.
type PlanPhasePercentageDiscountAdjustmentAdjustmentType = shared.PlanPhasePercentageDiscountAdjustmentAdjustmentType

// This is an alias to an internal value.
const PlanPhasePercentageDiscountAdjustmentAdjustmentTypePercentageDiscount = shared.PlanPhasePercentageDiscountAdjustmentAdjustmentTypePercentageDiscount

// This is an alias to an internal type.
type PlanPhaseUsageDiscountAdjustment = shared.PlanPhaseUsageDiscountAdjustment

// This is an alias to an internal type.
type PlanPhaseUsageDiscountAdjustmentAdjustmentType = shared.PlanPhaseUsageDiscountAdjustmentAdjustmentType

// This is an alias to an internal value.
const PlanPhaseUsageDiscountAdjustmentAdjustmentTypeUsageDiscount = shared.PlanPhaseUsageDiscountAdjustmentAdjustmentTypeUsageDiscount

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
type PriceUnitPriceConversionRateConfig = shared.PriceUnitPriceConversionRateConfig

// This is an alias to an internal type.
type PriceUnitPriceConversionRateConfigConversionRateType = shared.PriceUnitPriceConversionRateConfigConversionRateType

// This is an alias to an internal value.
const PriceUnitPriceConversionRateConfigConversionRateTypeUnit = shared.PriceUnitPriceConversionRateConfigConversionRateTypeUnit

// This is an alias to an internal value.
const PriceUnitPriceConversionRateConfigConversionRateTypeTiered = shared.PriceUnitPriceConversionRateConfigConversionRateTypeTiered

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

// This is an alias to an internal type.
type PricePackagePrice = shared.PricePackagePrice

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
type PricePackagePriceConversionRateConfig = shared.PricePackagePriceConversionRateConfig

// This is an alias to an internal type.
type PricePackagePriceConversionRateConfigConversionRateType = shared.PricePackagePriceConversionRateConfigConversionRateType

// This is an alias to an internal value.
const PricePackagePriceConversionRateConfigConversionRateTypeUnit = shared.PricePackagePriceConversionRateConfigConversionRateTypeUnit

// This is an alias to an internal value.
const PricePackagePriceConversionRateConfigConversionRateTypeTiered = shared.PricePackagePriceConversionRateConfigConversionRateTypeTiered

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

// This is an alias to an internal type.
type PriceMatrixPrice = shared.PriceMatrixPrice

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
type PriceMatrixPriceConversionRateConfig = shared.PriceMatrixPriceConversionRateConfig

// This is an alias to an internal type.
type PriceMatrixPriceConversionRateConfigConversionRateType = shared.PriceMatrixPriceConversionRateConfigConversionRateType

// This is an alias to an internal value.
const PriceMatrixPriceConversionRateConfigConversionRateTypeUnit = shared.PriceMatrixPriceConversionRateConfigConversionRateTypeUnit

// This is an alias to an internal value.
const PriceMatrixPriceConversionRateConfigConversionRateTypeTiered = shared.PriceMatrixPriceConversionRateConfigConversionRateTypeTiered

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

// This is an alias to an internal type.
type PriceTieredPrice = shared.PriceTieredPrice

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
type PriceTieredPriceConversionRateConfig = shared.PriceTieredPriceConversionRateConfig

// This is an alias to an internal type.
type PriceTieredPriceConversionRateConfigConversionRateType = shared.PriceTieredPriceConversionRateConfigConversionRateType

// This is an alias to an internal value.
const PriceTieredPriceConversionRateConfigConversionRateTypeUnit = shared.PriceTieredPriceConversionRateConfigConversionRateTypeUnit

// This is an alias to an internal value.
const PriceTieredPriceConversionRateConfigConversionRateTypeTiered = shared.PriceTieredPriceConversionRateConfigConversionRateTypeTiered

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

// This is an alias to an internal type.
type PriceTieredBPSPrice = shared.PriceTieredBPSPrice

// This is an alias to an internal type.
type PriceTieredBPSPriceCadence = shared.PriceTieredBPSPriceCadence

// This is an alias to an internal value.
const PriceTieredBPSPriceCadenceOneTime = shared.PriceTieredBPSPriceCadenceOneTime

// This is an alias to an internal value.
const PriceTieredBPSPriceCadenceMonthly = shared.PriceTieredBPSPriceCadenceMonthly

// This is an alias to an internal value.
const PriceTieredBPSPriceCadenceQuarterly = shared.PriceTieredBPSPriceCadenceQuarterly

// This is an alias to an internal value.
const PriceTieredBPSPriceCadenceSemiAnnual = shared.PriceTieredBPSPriceCadenceSemiAnnual

// This is an alias to an internal value.
const PriceTieredBPSPriceCadenceAnnual = shared.PriceTieredBPSPriceCadenceAnnual

// This is an alias to an internal value.
const PriceTieredBPSPriceCadenceCustom = shared.PriceTieredBPSPriceCadenceCustom

// This is an alias to an internal type.
type PriceTieredBPSPriceConversionRateConfig = shared.PriceTieredBPSPriceConversionRateConfig

// This is an alias to an internal type.
type PriceTieredBPSPriceConversionRateConfigConversionRateType = shared.PriceTieredBPSPriceConversionRateConfigConversionRateType

// This is an alias to an internal value.
const PriceTieredBPSPriceConversionRateConfigConversionRateTypeUnit = shared.PriceTieredBPSPriceConversionRateConfigConversionRateTypeUnit

// This is an alias to an internal value.
const PriceTieredBPSPriceConversionRateConfigConversionRateTypeTiered = shared.PriceTieredBPSPriceConversionRateConfigConversionRateTypeTiered

// This is an alias to an internal type.
type PriceTieredBPSPriceModelType = shared.PriceTieredBPSPriceModelType

// This is an alias to an internal value.
const PriceTieredBPSPriceModelTypeTieredBPS = shared.PriceTieredBPSPriceModelTypeTieredBPS

// This is an alias to an internal type.
type PriceTieredBPSPricePriceType = shared.PriceTieredBPSPricePriceType

// This is an alias to an internal value.
const PriceTieredBPSPricePriceTypeUsagePrice = shared.PriceTieredBPSPricePriceTypeUsagePrice

// This is an alias to an internal value.
const PriceTieredBPSPricePriceTypeFixedPrice = shared.PriceTieredBPSPricePriceTypeFixedPrice

// This is an alias to an internal type.
type PriceBPSPrice = shared.PriceBPSPrice

// This is an alias to an internal type.
type PriceBPSPriceCadence = shared.PriceBPSPriceCadence

// This is an alias to an internal value.
const PriceBPSPriceCadenceOneTime = shared.PriceBPSPriceCadenceOneTime

// This is an alias to an internal value.
const PriceBPSPriceCadenceMonthly = shared.PriceBPSPriceCadenceMonthly

// This is an alias to an internal value.
const PriceBPSPriceCadenceQuarterly = shared.PriceBPSPriceCadenceQuarterly

// This is an alias to an internal value.
const PriceBPSPriceCadenceSemiAnnual = shared.PriceBPSPriceCadenceSemiAnnual

// This is an alias to an internal value.
const PriceBPSPriceCadenceAnnual = shared.PriceBPSPriceCadenceAnnual

// This is an alias to an internal value.
const PriceBPSPriceCadenceCustom = shared.PriceBPSPriceCadenceCustom

// This is an alias to an internal type.
type PriceBPSPriceConversionRateConfig = shared.PriceBPSPriceConversionRateConfig

// This is an alias to an internal type.
type PriceBPSPriceConversionRateConfigConversionRateType = shared.PriceBPSPriceConversionRateConfigConversionRateType

// This is an alias to an internal value.
const PriceBPSPriceConversionRateConfigConversionRateTypeUnit = shared.PriceBPSPriceConversionRateConfigConversionRateTypeUnit

// This is an alias to an internal value.
const PriceBPSPriceConversionRateConfigConversionRateTypeTiered = shared.PriceBPSPriceConversionRateConfigConversionRateTypeTiered

// This is an alias to an internal type.
type PriceBPSPriceModelType = shared.PriceBPSPriceModelType

// This is an alias to an internal value.
const PriceBPSPriceModelTypeBPS = shared.PriceBPSPriceModelTypeBPS

// This is an alias to an internal type.
type PriceBPSPricePriceType = shared.PriceBPSPricePriceType

// This is an alias to an internal value.
const PriceBPSPricePriceTypeUsagePrice = shared.PriceBPSPricePriceTypeUsagePrice

// This is an alias to an internal value.
const PriceBPSPricePriceTypeFixedPrice = shared.PriceBPSPricePriceTypeFixedPrice

// This is an alias to an internal type.
type PriceBulkBPSPrice = shared.PriceBulkBPSPrice

// This is an alias to an internal type.
type PriceBulkBPSPriceCadence = shared.PriceBulkBPSPriceCadence

// This is an alias to an internal value.
const PriceBulkBPSPriceCadenceOneTime = shared.PriceBulkBPSPriceCadenceOneTime

// This is an alias to an internal value.
const PriceBulkBPSPriceCadenceMonthly = shared.PriceBulkBPSPriceCadenceMonthly

// This is an alias to an internal value.
const PriceBulkBPSPriceCadenceQuarterly = shared.PriceBulkBPSPriceCadenceQuarterly

// This is an alias to an internal value.
const PriceBulkBPSPriceCadenceSemiAnnual = shared.PriceBulkBPSPriceCadenceSemiAnnual

// This is an alias to an internal value.
const PriceBulkBPSPriceCadenceAnnual = shared.PriceBulkBPSPriceCadenceAnnual

// This is an alias to an internal value.
const PriceBulkBPSPriceCadenceCustom = shared.PriceBulkBPSPriceCadenceCustom

// This is an alias to an internal type.
type PriceBulkBPSPriceConversionRateConfig = shared.PriceBulkBPSPriceConversionRateConfig

// This is an alias to an internal type.
type PriceBulkBPSPriceConversionRateConfigConversionRateType = shared.PriceBulkBPSPriceConversionRateConfigConversionRateType

// This is an alias to an internal value.
const PriceBulkBPSPriceConversionRateConfigConversionRateTypeUnit = shared.PriceBulkBPSPriceConversionRateConfigConversionRateTypeUnit

// This is an alias to an internal value.
const PriceBulkBPSPriceConversionRateConfigConversionRateTypeTiered = shared.PriceBulkBPSPriceConversionRateConfigConversionRateTypeTiered

// This is an alias to an internal type.
type PriceBulkBPSPriceModelType = shared.PriceBulkBPSPriceModelType

// This is an alias to an internal value.
const PriceBulkBPSPriceModelTypeBulkBPS = shared.PriceBulkBPSPriceModelTypeBulkBPS

// This is an alias to an internal type.
type PriceBulkBPSPricePriceType = shared.PriceBulkBPSPricePriceType

// This is an alias to an internal value.
const PriceBulkBPSPricePriceTypeUsagePrice = shared.PriceBulkBPSPricePriceTypeUsagePrice

// This is an alias to an internal value.
const PriceBulkBPSPricePriceTypeFixedPrice = shared.PriceBulkBPSPricePriceTypeFixedPrice

// This is an alias to an internal type.
type PriceBulkPrice = shared.PriceBulkPrice

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
type PriceBulkPriceConversionRateConfig = shared.PriceBulkPriceConversionRateConfig

// This is an alias to an internal type.
type PriceBulkPriceConversionRateConfigConversionRateType = shared.PriceBulkPriceConversionRateConfigConversionRateType

// This is an alias to an internal value.
const PriceBulkPriceConversionRateConfigConversionRateTypeUnit = shared.PriceBulkPriceConversionRateConfigConversionRateTypeUnit

// This is an alias to an internal value.
const PriceBulkPriceConversionRateConfigConversionRateTypeTiered = shared.PriceBulkPriceConversionRateConfigConversionRateTypeTiered

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

// This is an alias to an internal type.
type PriceThresholdTotalAmountPrice = shared.PriceThresholdTotalAmountPrice

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
type PriceThresholdTotalAmountPriceConversionRateConfig = shared.PriceThresholdTotalAmountPriceConversionRateConfig

// This is an alias to an internal type.
type PriceThresholdTotalAmountPriceConversionRateConfigConversionRateType = shared.PriceThresholdTotalAmountPriceConversionRateConfigConversionRateType

// This is an alias to an internal value.
const PriceThresholdTotalAmountPriceConversionRateConfigConversionRateTypeUnit = shared.PriceThresholdTotalAmountPriceConversionRateConfigConversionRateTypeUnit

// This is an alias to an internal value.
const PriceThresholdTotalAmountPriceConversionRateConfigConversionRateTypeTiered = shared.PriceThresholdTotalAmountPriceConversionRateConfigConversionRateTypeTiered

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

// This is an alias to an internal type.
type PriceTieredPackagePrice = shared.PriceTieredPackagePrice

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
type PriceTieredPackagePriceConversionRateConfig = shared.PriceTieredPackagePriceConversionRateConfig

// This is an alias to an internal type.
type PriceTieredPackagePriceConversionRateConfigConversionRateType = shared.PriceTieredPackagePriceConversionRateConfigConversionRateType

// This is an alias to an internal value.
const PriceTieredPackagePriceConversionRateConfigConversionRateTypeUnit = shared.PriceTieredPackagePriceConversionRateConfigConversionRateTypeUnit

// This is an alias to an internal value.
const PriceTieredPackagePriceConversionRateConfigConversionRateTypeTiered = shared.PriceTieredPackagePriceConversionRateConfigConversionRateTypeTiered

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

// This is an alias to an internal type.
type PriceGroupedTieredPrice = shared.PriceGroupedTieredPrice

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
type PriceGroupedTieredPriceConversionRateConfig = shared.PriceGroupedTieredPriceConversionRateConfig

// This is an alias to an internal type.
type PriceGroupedTieredPriceConversionRateConfigConversionRateType = shared.PriceGroupedTieredPriceConversionRateConfigConversionRateType

// This is an alias to an internal value.
const PriceGroupedTieredPriceConversionRateConfigConversionRateTypeUnit = shared.PriceGroupedTieredPriceConversionRateConfigConversionRateTypeUnit

// This is an alias to an internal value.
const PriceGroupedTieredPriceConversionRateConfigConversionRateTypeTiered = shared.PriceGroupedTieredPriceConversionRateConfigConversionRateTypeTiered

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

// This is an alias to an internal type.
type PriceTieredWithMinimumPrice = shared.PriceTieredWithMinimumPrice

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
type PriceTieredWithMinimumPriceConversionRateConfig = shared.PriceTieredWithMinimumPriceConversionRateConfig

// This is an alias to an internal type.
type PriceTieredWithMinimumPriceConversionRateConfigConversionRateType = shared.PriceTieredWithMinimumPriceConversionRateConfigConversionRateType

// This is an alias to an internal value.
const PriceTieredWithMinimumPriceConversionRateConfigConversionRateTypeUnit = shared.PriceTieredWithMinimumPriceConversionRateConfigConversionRateTypeUnit

// This is an alias to an internal value.
const PriceTieredWithMinimumPriceConversionRateConfigConversionRateTypeTiered = shared.PriceTieredWithMinimumPriceConversionRateConfigConversionRateTypeTiered

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

// This is an alias to an internal type.
type PriceTieredPackageWithMinimumPrice = shared.PriceTieredPackageWithMinimumPrice

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
type PriceTieredPackageWithMinimumPriceConversionRateConfig = shared.PriceTieredPackageWithMinimumPriceConversionRateConfig

// This is an alias to an internal type.
type PriceTieredPackageWithMinimumPriceConversionRateConfigConversionRateType = shared.PriceTieredPackageWithMinimumPriceConversionRateConfigConversionRateType

// This is an alias to an internal value.
const PriceTieredPackageWithMinimumPriceConversionRateConfigConversionRateTypeUnit = shared.PriceTieredPackageWithMinimumPriceConversionRateConfigConversionRateTypeUnit

// This is an alias to an internal value.
const PriceTieredPackageWithMinimumPriceConversionRateConfigConversionRateTypeTiered = shared.PriceTieredPackageWithMinimumPriceConversionRateConfigConversionRateTypeTiered

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

// This is an alias to an internal type.
type PricePackageWithAllocationPrice = shared.PricePackageWithAllocationPrice

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
type PricePackageWithAllocationPriceConversionRateConfig = shared.PricePackageWithAllocationPriceConversionRateConfig

// This is an alias to an internal type.
type PricePackageWithAllocationPriceConversionRateConfigConversionRateType = shared.PricePackageWithAllocationPriceConversionRateConfigConversionRateType

// This is an alias to an internal value.
const PricePackageWithAllocationPriceConversionRateConfigConversionRateTypeUnit = shared.PricePackageWithAllocationPriceConversionRateConfigConversionRateTypeUnit

// This is an alias to an internal value.
const PricePackageWithAllocationPriceConversionRateConfigConversionRateTypeTiered = shared.PricePackageWithAllocationPriceConversionRateConfigConversionRateTypeTiered

// This is an alias to an internal type.
type PricePackageWithAllocationPriceModelType = shared.PricePackageWithAllocationPriceModelType

// This is an alias to an internal value.
const PricePackageWithAllocationPriceModelTypePackageWithAllocation = shared.PricePackageWithAllocationPriceModelTypePackageWithAllocation

// This is an alias to an internal type.
type PricePackageWithAllocationPricePriceType = shared.PricePackageWithAllocationPricePriceType

// This is an alias to an internal value.
const PricePackageWithAllocationPricePriceTypeUsagePrice = shared.PricePackageWithAllocationPricePriceTypeUsagePrice

// This is an alias to an internal value.
const PricePackageWithAllocationPricePriceTypeFixedPrice = shared.PricePackageWithAllocationPricePriceTypeFixedPrice

// This is an alias to an internal type.
type PriceUnitWithPercentPrice = shared.PriceUnitWithPercentPrice

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
type PriceUnitWithPercentPriceConversionRateConfig = shared.PriceUnitWithPercentPriceConversionRateConfig

// This is an alias to an internal type.
type PriceUnitWithPercentPriceConversionRateConfigConversionRateType = shared.PriceUnitWithPercentPriceConversionRateConfigConversionRateType

// This is an alias to an internal value.
const PriceUnitWithPercentPriceConversionRateConfigConversionRateTypeUnit = shared.PriceUnitWithPercentPriceConversionRateConfigConversionRateTypeUnit

// This is an alias to an internal value.
const PriceUnitWithPercentPriceConversionRateConfigConversionRateTypeTiered = shared.PriceUnitWithPercentPriceConversionRateConfigConversionRateTypeTiered

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

// This is an alias to an internal type.
type PriceMatrixWithAllocationPrice = shared.PriceMatrixWithAllocationPrice

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
type PriceMatrixWithAllocationPriceConversionRateConfig = shared.PriceMatrixWithAllocationPriceConversionRateConfig

// This is an alias to an internal type.
type PriceMatrixWithAllocationPriceConversionRateConfigConversionRateType = shared.PriceMatrixWithAllocationPriceConversionRateConfigConversionRateType

// This is an alias to an internal value.
const PriceMatrixWithAllocationPriceConversionRateConfigConversionRateTypeUnit = shared.PriceMatrixWithAllocationPriceConversionRateConfigConversionRateTypeUnit

// This is an alias to an internal value.
const PriceMatrixWithAllocationPriceConversionRateConfigConversionRateTypeTiered = shared.PriceMatrixWithAllocationPriceConversionRateConfigConversionRateTypeTiered

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

// This is an alias to an internal type.
type PriceTieredWithProrationPrice = shared.PriceTieredWithProrationPrice

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
type PriceTieredWithProrationPriceConversionRateConfig = shared.PriceTieredWithProrationPriceConversionRateConfig

// This is an alias to an internal type.
type PriceTieredWithProrationPriceConversionRateConfigConversionRateType = shared.PriceTieredWithProrationPriceConversionRateConfigConversionRateType

// This is an alias to an internal value.
const PriceTieredWithProrationPriceConversionRateConfigConversionRateTypeUnit = shared.PriceTieredWithProrationPriceConversionRateConfigConversionRateTypeUnit

// This is an alias to an internal value.
const PriceTieredWithProrationPriceConversionRateConfigConversionRateTypeTiered = shared.PriceTieredWithProrationPriceConversionRateConfigConversionRateTypeTiered

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

// This is an alias to an internal type.
type PriceUnitWithProrationPrice = shared.PriceUnitWithProrationPrice

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
type PriceUnitWithProrationPriceConversionRateConfig = shared.PriceUnitWithProrationPriceConversionRateConfig

// This is an alias to an internal type.
type PriceUnitWithProrationPriceConversionRateConfigConversionRateType = shared.PriceUnitWithProrationPriceConversionRateConfigConversionRateType

// This is an alias to an internal value.
const PriceUnitWithProrationPriceConversionRateConfigConversionRateTypeUnit = shared.PriceUnitWithProrationPriceConversionRateConfigConversionRateTypeUnit

// This is an alias to an internal value.
const PriceUnitWithProrationPriceConversionRateConfigConversionRateTypeTiered = shared.PriceUnitWithProrationPriceConversionRateConfigConversionRateTypeTiered

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

// This is an alias to an internal type.
type PriceGroupedAllocationPrice = shared.PriceGroupedAllocationPrice

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
type PriceGroupedAllocationPriceConversionRateConfig = shared.PriceGroupedAllocationPriceConversionRateConfig

// This is an alias to an internal type.
type PriceGroupedAllocationPriceConversionRateConfigConversionRateType = shared.PriceGroupedAllocationPriceConversionRateConfigConversionRateType

// This is an alias to an internal value.
const PriceGroupedAllocationPriceConversionRateConfigConversionRateTypeUnit = shared.PriceGroupedAllocationPriceConversionRateConfigConversionRateTypeUnit

// This is an alias to an internal value.
const PriceGroupedAllocationPriceConversionRateConfigConversionRateTypeTiered = shared.PriceGroupedAllocationPriceConversionRateConfigConversionRateTypeTiered

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

// This is an alias to an internal type.
type PriceGroupedWithProratedMinimumPrice = shared.PriceGroupedWithProratedMinimumPrice

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
type PriceGroupedWithProratedMinimumPriceConversionRateConfig = shared.PriceGroupedWithProratedMinimumPriceConversionRateConfig

// This is an alias to an internal type.
type PriceGroupedWithProratedMinimumPriceConversionRateConfigConversionRateType = shared.PriceGroupedWithProratedMinimumPriceConversionRateConfigConversionRateType

// This is an alias to an internal value.
const PriceGroupedWithProratedMinimumPriceConversionRateConfigConversionRateTypeUnit = shared.PriceGroupedWithProratedMinimumPriceConversionRateConfigConversionRateTypeUnit

// This is an alias to an internal value.
const PriceGroupedWithProratedMinimumPriceConversionRateConfigConversionRateTypeTiered = shared.PriceGroupedWithProratedMinimumPriceConversionRateConfigConversionRateTypeTiered

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

// This is an alias to an internal type.
type PriceGroupedWithMeteredMinimumPrice = shared.PriceGroupedWithMeteredMinimumPrice

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
type PriceGroupedWithMeteredMinimumPriceConversionRateConfig = shared.PriceGroupedWithMeteredMinimumPriceConversionRateConfig

// This is an alias to an internal type.
type PriceGroupedWithMeteredMinimumPriceConversionRateConfigConversionRateType = shared.PriceGroupedWithMeteredMinimumPriceConversionRateConfigConversionRateType

// This is an alias to an internal value.
const PriceGroupedWithMeteredMinimumPriceConversionRateConfigConversionRateTypeUnit = shared.PriceGroupedWithMeteredMinimumPriceConversionRateConfigConversionRateTypeUnit

// This is an alias to an internal value.
const PriceGroupedWithMeteredMinimumPriceConversionRateConfigConversionRateTypeTiered = shared.PriceGroupedWithMeteredMinimumPriceConversionRateConfigConversionRateTypeTiered

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

// This is an alias to an internal type.
type PriceMatrixWithDisplayNamePrice = shared.PriceMatrixWithDisplayNamePrice

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
type PriceMatrixWithDisplayNamePriceConversionRateConfig = shared.PriceMatrixWithDisplayNamePriceConversionRateConfig

// This is an alias to an internal type.
type PriceMatrixWithDisplayNamePriceConversionRateConfigConversionRateType = shared.PriceMatrixWithDisplayNamePriceConversionRateConfigConversionRateType

// This is an alias to an internal value.
const PriceMatrixWithDisplayNamePriceConversionRateConfigConversionRateTypeUnit = shared.PriceMatrixWithDisplayNamePriceConversionRateConfigConversionRateTypeUnit

// This is an alias to an internal value.
const PriceMatrixWithDisplayNamePriceConversionRateConfigConversionRateTypeTiered = shared.PriceMatrixWithDisplayNamePriceConversionRateConfigConversionRateTypeTiered

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

// This is an alias to an internal type.
type PriceBulkWithProrationPrice = shared.PriceBulkWithProrationPrice

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
type PriceBulkWithProrationPriceConversionRateConfig = shared.PriceBulkWithProrationPriceConversionRateConfig

// This is an alias to an internal type.
type PriceBulkWithProrationPriceConversionRateConfigConversionRateType = shared.PriceBulkWithProrationPriceConversionRateConfigConversionRateType

// This is an alias to an internal value.
const PriceBulkWithProrationPriceConversionRateConfigConversionRateTypeUnit = shared.PriceBulkWithProrationPriceConversionRateConfigConversionRateTypeUnit

// This is an alias to an internal value.
const PriceBulkWithProrationPriceConversionRateConfigConversionRateTypeTiered = shared.PriceBulkWithProrationPriceConversionRateConfigConversionRateTypeTiered

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

// This is an alias to an internal type.
type PriceGroupedTieredPackagePrice = shared.PriceGroupedTieredPackagePrice

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
type PriceGroupedTieredPackagePriceConversionRateConfig = shared.PriceGroupedTieredPackagePriceConversionRateConfig

// This is an alias to an internal type.
type PriceGroupedTieredPackagePriceConversionRateConfigConversionRateType = shared.PriceGroupedTieredPackagePriceConversionRateConfigConversionRateType

// This is an alias to an internal value.
const PriceGroupedTieredPackagePriceConversionRateConfigConversionRateTypeUnit = shared.PriceGroupedTieredPackagePriceConversionRateConfigConversionRateTypeUnit

// This is an alias to an internal value.
const PriceGroupedTieredPackagePriceConversionRateConfigConversionRateTypeTiered = shared.PriceGroupedTieredPackagePriceConversionRateConfigConversionRateTypeTiered

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

// This is an alias to an internal type.
type PriceMaxGroupTieredPackagePrice = shared.PriceMaxGroupTieredPackagePrice

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
type PriceMaxGroupTieredPackagePriceConversionRateConfig = shared.PriceMaxGroupTieredPackagePriceConversionRateConfig

// This is an alias to an internal type.
type PriceMaxGroupTieredPackagePriceConversionRateConfigConversionRateType = shared.PriceMaxGroupTieredPackagePriceConversionRateConfigConversionRateType

// This is an alias to an internal value.
const PriceMaxGroupTieredPackagePriceConversionRateConfigConversionRateTypeUnit = shared.PriceMaxGroupTieredPackagePriceConversionRateConfigConversionRateTypeUnit

// This is an alias to an internal value.
const PriceMaxGroupTieredPackagePriceConversionRateConfigConversionRateTypeTiered = shared.PriceMaxGroupTieredPackagePriceConversionRateConfigConversionRateTypeTiered

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

// This is an alias to an internal type.
type PriceScalableMatrixWithUnitPricingPrice = shared.PriceScalableMatrixWithUnitPricingPrice

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
type PriceScalableMatrixWithUnitPricingPriceConversionRateConfig = shared.PriceScalableMatrixWithUnitPricingPriceConversionRateConfig

// This is an alias to an internal type.
type PriceScalableMatrixWithUnitPricingPriceConversionRateConfigConversionRateType = shared.PriceScalableMatrixWithUnitPricingPriceConversionRateConfigConversionRateType

// This is an alias to an internal value.
const PriceScalableMatrixWithUnitPricingPriceConversionRateConfigConversionRateTypeUnit = shared.PriceScalableMatrixWithUnitPricingPriceConversionRateConfigConversionRateTypeUnit

// This is an alias to an internal value.
const PriceScalableMatrixWithUnitPricingPriceConversionRateConfigConversionRateTypeTiered = shared.PriceScalableMatrixWithUnitPricingPriceConversionRateConfigConversionRateTypeTiered

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

// This is an alias to an internal type.
type PriceScalableMatrixWithTieredPricingPrice = shared.PriceScalableMatrixWithTieredPricingPrice

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
type PriceScalableMatrixWithTieredPricingPriceConversionRateConfig = shared.PriceScalableMatrixWithTieredPricingPriceConversionRateConfig

// This is an alias to an internal type.
type PriceScalableMatrixWithTieredPricingPriceConversionRateConfigConversionRateType = shared.PriceScalableMatrixWithTieredPricingPriceConversionRateConfigConversionRateType

// This is an alias to an internal value.
const PriceScalableMatrixWithTieredPricingPriceConversionRateConfigConversionRateTypeUnit = shared.PriceScalableMatrixWithTieredPricingPriceConversionRateConfigConversionRateTypeUnit

// This is an alias to an internal value.
const PriceScalableMatrixWithTieredPricingPriceConversionRateConfigConversionRateTypeTiered = shared.PriceScalableMatrixWithTieredPricingPriceConversionRateConfigConversionRateTypeTiered

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

// This is an alias to an internal type.
type PriceCumulativeGroupedBulkPrice = shared.PriceCumulativeGroupedBulkPrice

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
type PriceCumulativeGroupedBulkPriceConversionRateConfig = shared.PriceCumulativeGroupedBulkPriceConversionRateConfig

// This is an alias to an internal type.
type PriceCumulativeGroupedBulkPriceConversionRateConfigConversionRateType = shared.PriceCumulativeGroupedBulkPriceConversionRateConfigConversionRateType

// This is an alias to an internal value.
const PriceCumulativeGroupedBulkPriceConversionRateConfigConversionRateTypeUnit = shared.PriceCumulativeGroupedBulkPriceConversionRateConfigConversionRateTypeUnit

// This is an alias to an internal value.
const PriceCumulativeGroupedBulkPriceConversionRateConfigConversionRateTypeTiered = shared.PriceCumulativeGroupedBulkPriceConversionRateConfigConversionRateTypeTiered

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

// This is an alias to an internal type.
type PriceGroupedWithMinMaxThresholdsPrice = shared.PriceGroupedWithMinMaxThresholdsPrice

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
type PriceGroupedWithMinMaxThresholdsPriceConversionRateConfig = shared.PriceGroupedWithMinMaxThresholdsPriceConversionRateConfig

// This is an alias to an internal type.
type PriceGroupedWithMinMaxThresholdsPriceConversionRateConfigConversionRateType = shared.PriceGroupedWithMinMaxThresholdsPriceConversionRateConfigConversionRateType

// This is an alias to an internal value.
const PriceGroupedWithMinMaxThresholdsPriceConversionRateConfigConversionRateTypeUnit = shared.PriceGroupedWithMinMaxThresholdsPriceConversionRateConfigConversionRateTypeUnit

// This is an alias to an internal value.
const PriceGroupedWithMinMaxThresholdsPriceConversionRateConfigConversionRateTypeTiered = shared.PriceGroupedWithMinMaxThresholdsPriceConversionRateConfigConversionRateTypeTiered

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

// This is an alias to an internal type.
type PriceModelType = shared.PriceModelType

// This is an alias to an internal value.
const PriceModelTypeUnit = shared.PriceModelTypeUnit

// This is an alias to an internal value.
const PriceModelTypePackage = shared.PriceModelTypePackage

// This is an alias to an internal value.
const PriceModelTypeMatrix = shared.PriceModelTypeMatrix

// This is an alias to an internal value.
const PriceModelTypeTiered = shared.PriceModelTypeTiered

// This is an alias to an internal value.
const PriceModelTypeTieredBPS = shared.PriceModelTypeTieredBPS

// This is an alias to an internal value.
const PriceModelTypeBPS = shared.PriceModelTypeBPS

// This is an alias to an internal value.
const PriceModelTypeBulkBPS = shared.PriceModelTypeBulkBPS

// This is an alias to an internal value.
const PriceModelTypeBulk = shared.PriceModelTypeBulk

// This is an alias to an internal value.
const PriceModelTypeThresholdTotalAmount = shared.PriceModelTypeThresholdTotalAmount

// This is an alias to an internal value.
const PriceModelTypeTieredPackage = shared.PriceModelTypeTieredPackage

// This is an alias to an internal value.
const PriceModelTypeGroupedTiered = shared.PriceModelTypeGroupedTiered

// This is an alias to an internal value.
const PriceModelTypeTieredWithMinimum = shared.PriceModelTypeTieredWithMinimum

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
const PriceModelTypeGroupedWithProratedMinimum = shared.PriceModelTypeGroupedWithProratedMinimum

// This is an alias to an internal value.
const PriceModelTypeGroupedWithMeteredMinimum = shared.PriceModelTypeGroupedWithMeteredMinimum

// This is an alias to an internal value.
const PriceModelTypeMatrixWithDisplayName = shared.PriceModelTypeMatrixWithDisplayName

// This is an alias to an internal value.
const PriceModelTypeBulkWithProration = shared.PriceModelTypeBulkWithProration

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
const PriceModelTypeGroupedWithMinMaxThresholds = shared.PriceModelTypeGroupedWithMinMaxThresholds

// This is an alias to an internal type.
type PricePriceType = shared.PricePriceType

// This is an alias to an internal value.
const PricePriceTypeUsagePrice = shared.PricePriceTypeUsagePrice

// This is an alias to an internal value.
const PricePriceTypeFixedPrice = shared.PricePriceTypeFixedPrice

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

// This is an alias to an internal type.
type Tier = shared.Tier

// This is an alias to an internal type.
type TierParam = shared.TierParam

// This is an alias to an internal type.
type TierConfig = shared.TierConfig

// This is an alias to an internal type.
type TierSubLineItem = shared.TierSubLineItem

// This is an alias to an internal type.
type TierSubLineItemType = shared.TierSubLineItemType

// This is an alias to an internal value.
const TierSubLineItemTypeTier = shared.TierSubLineItemTypeTier

// This is an alias to an internal type.
type TieredBPSConfig = shared.TieredBPSConfig

// This is an alias to an internal type.
type TieredBPSConfigParam = shared.TieredBPSConfigParam

// This is an alias to an internal type.
type TieredConfig = shared.TieredConfig

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
type TransformPriceFilter = shared.TransformPriceFilter

// The property of the price to filter on.
//
// This is an alias to an internal type.
type TransformPriceFilterField = shared.TransformPriceFilterField

// This is an alias to an internal value.
const TransformPriceFilterFieldPriceID = shared.TransformPriceFilterFieldPriceID

// This is an alias to an internal value.
const TransformPriceFilterFieldItemID = shared.TransformPriceFilterFieldItemID

// This is an alias to an internal value.
const TransformPriceFilterFieldPriceType = shared.TransformPriceFilterFieldPriceType

// This is an alias to an internal value.
const TransformPriceFilterFieldCurrency = shared.TransformPriceFilterFieldCurrency

// This is an alias to an internal value.
const TransformPriceFilterFieldPricingUnitID = shared.TransformPriceFilterFieldPricingUnitID

// Should prices that match the filter be included or excluded.
//
// This is an alias to an internal type.
type TransformPriceFilterOperator = shared.TransformPriceFilterOperator

// This is an alias to an internal value.
const TransformPriceFilterOperatorIncludes = shared.TransformPriceFilterOperatorIncludes

// This is an alias to an internal value.
const TransformPriceFilterOperatorExcludes = shared.TransformPriceFilterOperatorExcludes

// This is an alias to an internal type.
type TransformPriceFilterParam = shared.TransformPriceFilterParam

// This is an alias to an internal type.
type TrialDiscount = shared.TrialDiscount

// This is an alias to an internal type.
type TrialDiscountDiscountType = shared.TrialDiscountDiscountType

// This is an alias to an internal value.
const TrialDiscountDiscountTypeTrial = shared.TrialDiscountDiscountTypeTrial

// This is an alias to an internal type.
type TrialDiscountParam = shared.TrialDiscountParam

// This is an alias to an internal type.
type UnitConfig = shared.UnitConfig

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
type UsageDiscountParam = shared.UsageDiscountParam

// This is an alias to an internal type.
type UsageDiscountInterval = shared.UsageDiscountInterval

// This is an alias to an internal type.
type UsageDiscountIntervalDiscountType = shared.UsageDiscountIntervalDiscountType

// This is an alias to an internal value.
const UsageDiscountIntervalDiscountTypeUsage = shared.UsageDiscountIntervalDiscountTypeUsage
