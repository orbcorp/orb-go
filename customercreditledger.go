// File generated from our OpenAPI spec by Stainless.

package orb

import (
	"context"
	"fmt"
	"net/http"
	"net/url"
	"reflect"
	"time"

	"github.com/orbcorp/orb-go/internal/apijson"
	"github.com/orbcorp/orb-go/internal/apiquery"
	"github.com/orbcorp/orb-go/internal/param"
	"github.com/orbcorp/orb-go/internal/requestconfig"
	"github.com/orbcorp/orb-go/internal/shared"
	"github.com/orbcorp/orb-go/option"
	"github.com/tidwall/gjson"
)

// CustomerCreditLedgerService contains methods and other services that help with
// interacting with the orb API. Note, unlike clients, this service does not read
// variables from the environment automatically. You should not instantiate this
// service directly, and instead use the [NewCustomerCreditLedgerService] method
// instead.
type CustomerCreditLedgerService struct {
	Options []option.RequestOption
}

// NewCustomerCreditLedgerService generates a new service that applies the given
// options to each request. These options are applied after the parent client's
// options (if there is one), and before any request-specific options.
func NewCustomerCreditLedgerService(opts ...option.RequestOption) (r *CustomerCreditLedgerService) {
	r = &CustomerCreditLedgerService{}
	r.Options = opts
	return
}

// The credits ledger provides _auditing_ functionality over Orb's credits system
// with a list of actions that have taken place to modify a customer's credit
// balance. This [paginated endpoint](../reference/pagination) lists these entries,
// starting from the most recent ledger entry.
//
// More details on using Orb's real-time credit feature are
// [here](../guides/product-catalog/prepurchase.md).
//
// There are four major types of modifications to credit balance, detailed below.
//
// ## Increment
//
// Credits (which optionally expire on a future date) can be added via the API
// ([Add Ledger Entry](create-ledger-entry)). The ledger entry for such an action
// will always contain the total eligible starting and ending balance for the
// customer at the time the entry was added to the ledger.
//
// ## Decrement
//
// Deductions can occur as a result of an API call to create a ledger entry (see
// [Add Ledger Entry](create-ledger-entry)), or automatically as a result of
// incurring usage. Both ledger entries present the `decrement` entry type.
//
// As usage for a customer is reported into Orb, credits may be deducted according
// to the customer's plan configuration. An automated deduction of this type will
// result in a ledger entry, also with a starting and ending balance. In order to
// provide better tracing capabilities for automatic deductions, Orb always
// associates each automatic deduction with the `event_id` at the time of
// ingestion, used to pinpoint _why_ credit deduction took place and to ensure that
// credits are never deducted without an associated usage event.
//
// By default, Orb uses an algorithm that automatically deducts from the _soonest
// expiring credit block_ first in order to ensure that all credits are utilized
// appropriately. As an example, if trial credits with an expiration date of 2
// weeks from now are present for a customer, they will be used before any
// deductions take place from a non-expiring credit block.
//
// If there are multiple blocks with the same expiration date, Orb will deduct from
// the block with the _lower cost basis_ first (ex. trial credits with a $0 cost
// basis before paid credits with a $5.00 cost basis).
//
// It's also possible for a single usage event's deduction to _span_ credit blocks.
// In this case, Orb will deduct from the next block, ending at the credit block
// which consists of unexpiring credits. Each of these deductions will lead to a
// _separate_ ledger entry, one per credit block that is deducted from. By default,
// the customer's total credit balance in Orb can be negative as a result of a
// decrement.
//
// ## Expiration change
//
// The expiry of credits can be changed as a result of the API (See
// [Add Ledger Entry](create-ledger-entry)). This will create a ledger entry that
// specifies the balance as well as the initial and target expiry dates.
//
// Note that for this entry type, `starting_balance` will equal `ending_balance`,
// and the `amount` represents the balance transferred. The credit block linked to
// the ledger entry is the source credit block from which there was an expiration
// change
//
// ## Credits expiry
//
// When a set of credits expire on pre-set expiration date, the customer's balance
// automatically reflects this change and adds an entry to the ledger indicating
// this event. Note that credit expiry should always happen close to a date
// boundary in the customer's timezone.
//
// ## Void initiated
//
// Credit blocks can be voided via the API. The `amount` on this entry corresponds
// to the number of credits that were remaining in the block at time of void.
// `void_reason` will be populated if the void is created with a reason.
//
// ## Void
//
// When a set of credits is voided, the customer's balance automatically reflects
// this change and adds an entry to the ledger indicating this event.
//
// ## Amendment
//
// When credits are added to a customer's balance as a result of a correction, this
// entry will be added to the ledger to indicate the adjustment of credits.
func (r *CustomerCreditLedgerService) List(ctx context.Context, customerID string, query CustomerCreditLedgerListParams, opts ...option.RequestOption) (res *shared.Page[CustomerCreditLedgerListResponse], err error) {
	var raw *http.Response
	opts = append(r.Options, opts...)
	opts = append([]option.RequestOption{option.WithResponseInto(&raw)}, opts...)
	path := fmt.Sprintf("customers/%s/credits/ledger", customerID)
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

// The credits ledger provides _auditing_ functionality over Orb's credits system
// with a list of actions that have taken place to modify a customer's credit
// balance. This [paginated endpoint](../reference/pagination) lists these entries,
// starting from the most recent ledger entry.
//
// More details on using Orb's real-time credit feature are
// [here](../guides/product-catalog/prepurchase.md).
//
// There are four major types of modifications to credit balance, detailed below.
//
// ## Increment
//
// Credits (which optionally expire on a future date) can be added via the API
// ([Add Ledger Entry](create-ledger-entry)). The ledger entry for such an action
// will always contain the total eligible starting and ending balance for the
// customer at the time the entry was added to the ledger.
//
// ## Decrement
//
// Deductions can occur as a result of an API call to create a ledger entry (see
// [Add Ledger Entry](create-ledger-entry)), or automatically as a result of
// incurring usage. Both ledger entries present the `decrement` entry type.
//
// As usage for a customer is reported into Orb, credits may be deducted according
// to the customer's plan configuration. An automated deduction of this type will
// result in a ledger entry, also with a starting and ending balance. In order to
// provide better tracing capabilities for automatic deductions, Orb always
// associates each automatic deduction with the `event_id` at the time of
// ingestion, used to pinpoint _why_ credit deduction took place and to ensure that
// credits are never deducted without an associated usage event.
//
// By default, Orb uses an algorithm that automatically deducts from the _soonest
// expiring credit block_ first in order to ensure that all credits are utilized
// appropriately. As an example, if trial credits with an expiration date of 2
// weeks from now are present for a customer, they will be used before any
// deductions take place from a non-expiring credit block.
//
// If there are multiple blocks with the same expiration date, Orb will deduct from
// the block with the _lower cost basis_ first (ex. trial credits with a $0 cost
// basis before paid credits with a $5.00 cost basis).
//
// It's also possible for a single usage event's deduction to _span_ credit blocks.
// In this case, Orb will deduct from the next block, ending at the credit block
// which consists of unexpiring credits. Each of these deductions will lead to a
// _separate_ ledger entry, one per credit block that is deducted from. By default,
// the customer's total credit balance in Orb can be negative as a result of a
// decrement.
//
// ## Expiration change
//
// The expiry of credits can be changed as a result of the API (See
// [Add Ledger Entry](create-ledger-entry)). This will create a ledger entry that
// specifies the balance as well as the initial and target expiry dates.
//
// Note that for this entry type, `starting_balance` will equal `ending_balance`,
// and the `amount` represents the balance transferred. The credit block linked to
// the ledger entry is the source credit block from which there was an expiration
// change
//
// ## Credits expiry
//
// When a set of credits expire on pre-set expiration date, the customer's balance
// automatically reflects this change and adds an entry to the ledger indicating
// this event. Note that credit expiry should always happen close to a date
// boundary in the customer's timezone.
//
// ## Void initiated
//
// Credit blocks can be voided via the API. The `amount` on this entry corresponds
// to the number of credits that were remaining in the block at time of void.
// `void_reason` will be populated if the void is created with a reason.
//
// ## Void
//
// When a set of credits is voided, the customer's balance automatically reflects
// this change and adds an entry to the ledger indicating this event.
//
// ## Amendment
//
// When credits are added to a customer's balance as a result of a correction, this
// entry will be added to the ledger to indicate the adjustment of credits.
func (r *CustomerCreditLedgerService) ListAutoPaging(ctx context.Context, customerID string, query CustomerCreditLedgerListParams, opts ...option.RequestOption) *shared.PageAutoPager[CustomerCreditLedgerListResponse] {
	return shared.NewPageAutoPager(r.List(ctx, customerID, query, opts...))
}

// This endpoint allows you to create a new ledger entry for a specified customer's
// balance. This can be used to increment balance, deduct credits, and change the
// expiry date of existing credits.
//
// ## Effects of adding a ledger entry
//
//  1. After calling this endpoint, [Fetch Credit Balance](fetch-customer-credits)
//     will return a credit block that represents the changes (i.e. balance changes
//     or transfers).
//  2. A ledger entry will be added to the credits ledger for this customer, and
//     therefore returned in the [View Credits Ledger](fetch-customer-credits)
//     response as well as serialized in the response to this request. In the case
//     of deductions without a specified block, multiple ledger entries may be
//     created if the deduction spans credit blocks.
//  3. If `invoice_settings` is specified, an invoice will be created that reflects
//     the cost of the credits (based on `amount` and `per_unit_cost_basis`).
//
// ## Adding credits
//
// Adding credits is done by creating an entry of type `increment`. This requires
// the caller to specify a number of credits as well as an optional expiry date in
// `YYYY-MM-DD` format. Orb also recommends specifying a description to assist with
// auditing. When adding credits, the caller can also specify a cost basis
// per-credit, to indicate how much in USD a customer paid for a single credit in a
// block. This can later be used for revenue recognition.
//
// The following snippet illustrates a sample request body to increment credits
// which will expire in January of 2022.
//
// ```json
//
//	{
//	  "entry_type": "increment",
//	  "amount": 100,
//	  "expiry_date": "2022-12-28",
//	  "per_unit_cost_basis": "0.20",
//	  "description": "Purchased 100 credits"
//	}
//
// ```
//
// Note that by default, Orb will always first increment any _negative_ balance in
// existing blocks before adding the remaining amount to the desired credit block.
//
// ### Invoicing for credits
//
// By default, Orb manipulates the credit ledger but does not charge for credits.
// However, if you pass `invoice_settings` in the body of this request, Orb will
// also generate a one-off invoice for the customer for the credits pre-purchase.
// Note that you _must_ provide the `per_unit_cost_basis`, since the total charges
// on the invoice are calculated by multiplying the cost basis with the number of
// credit units added.
//
// ## Deducting Credits
//
// Orb allows you to deduct credits from a customer by creating an entry of type
// `decrement`. Orb matches the algorithm for automatic deductions for determining
// which credit blocks to decrement from. In the case that the deduction leads to
// multiple ledger entries, the response from this endpoint will be the final
// deduction. Orb also optionally allows specifying a description to assist with
// auditing.
//
// The following snippet illustrates a sample request body to decrement credits.
//
// ```json
//
//	{
//	  "entry_type": "decrement",
//	  "amount": 20,
//	  "description": "Removing excess credits"
//	}
//
// ```
//
// ## Changing credits expiry
//
// If you'd like to change when existing credits expire, you should create a ledger
// entry of type `expiration_change`. For this entry, the required parameter
// `expiry_date` identifies the _originating_ block, and the required parameter
// `target_expiry_date` identifies when the transferred credits should now expire.
// A new credit block will be created with expiry date `target_expiry_date`, with
// the same cost basis data as the original credit block, if present.
//
// Note that the balance of the block with the given `expiry_date` must be at least
// equal to the desired transfer amount determined by the `amount` parameter.
//
// The following snippet illustrates a sample request body to extend the expiration
// date of credits by one year:
//
// ```json
//
//	{
//	  "entry_type": "expiration_change",
//	  "amount": 10,
//	  "expiry_date": "2022-12-28",
//	  "block_id": "UiUhFWeLHPrBY4Ad",
//	  "target_expiry_date": "2023-12-28",
//	  "description": "Extending credit validity"
//	}
//
// ```
//
// ## Voiding credits
//
// If you'd like to void a credit block, create a ledger entry of type `void`. For
// this entry, `block_id` is required to identify the block, and `amount` indicates
// how many credits to void, up to the block's initial balance. Pass in a
// `void_reason` of `refund` if the void is due to a refund.
//
// ## Amendment
//
// If you'd like to undo a decrement on a credit block, create a ledger entry of
// type `amendment`. For this entry, `block_id` is required to identify the block
// that was originally decremented from, and `amount` indicates how many credits to
// return to the customer, up to the block's initial balance.
func (r *CustomerCreditLedgerService) NewEntry(ctx context.Context, customerID string, body CustomerCreditLedgerNewEntryParams, opts ...option.RequestOption) (res *CustomerCreditLedgerNewEntryResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := fmt.Sprintf("customers/%s/credits/ledger_entry", customerID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return
}

// This endpoint allows you to create a new ledger entry for a specified customer's
// balance. This can be used to increment balance, deduct credits, and change the
// expiry date of existing credits.
//
// ## Effects of adding a ledger entry
//
//  1. After calling this endpoint, [Fetch Credit Balance](fetch-customer-credits)
//     will return a credit block that represents the changes (i.e. balance changes
//     or transfers).
//  2. A ledger entry will be added to the credits ledger for this customer, and
//     therefore returned in the [View Credits Ledger](fetch-customer-credits)
//     response as well as serialized in the response to this request. In the case
//     of deductions without a specified block, multiple ledger entries may be
//     created if the deduction spans credit blocks.
//  3. If `invoice_settings` is specified, an invoice will be created that reflects
//     the cost of the credits (based on `amount` and `per_unit_cost_basis`).
//
// ## Adding credits
//
// Adding credits is done by creating an entry of type `increment`. This requires
// the caller to specify a number of credits as well as an optional expiry date in
// `YYYY-MM-DD` format. Orb also recommends specifying a description to assist with
// auditing. When adding credits, the caller can also specify a cost basis
// per-credit, to indicate how much in USD a customer paid for a single credit in a
// block. This can later be used for revenue recognition.
//
// The following snippet illustrates a sample request body to increment credits
// which will expire in January of 2022.
//
// ```json
//
//	{
//	  "entry_type": "increment",
//	  "amount": 100,
//	  "expiry_date": "2022-12-28",
//	  "per_unit_cost_basis": "0.20",
//	  "description": "Purchased 100 credits"
//	}
//
// ```
//
// Note that by default, Orb will always first increment any _negative_ balance in
// existing blocks before adding the remaining amount to the desired credit block.
//
// ### Invoicing for credits
//
// By default, Orb manipulates the credit ledger but does not charge for credits.
// However, if you pass `invoice_settings` in the body of this request, Orb will
// also generate a one-off invoice for the customer for the credits pre-purchase.
// Note that you _must_ provide the `per_unit_cost_basis`, since the total charges
// on the invoice are calculated by multiplying the cost basis with the number of
// credit units added.
//
// ## Deducting Credits
//
// Orb allows you to deduct credits from a customer by creating an entry of type
// `decrement`. Orb matches the algorithm for automatic deductions for determining
// which credit blocks to decrement from. In the case that the deduction leads to
// multiple ledger entries, the response from this endpoint will be the final
// deduction. Orb also optionally allows specifying a description to assist with
// auditing.
//
// The following snippet illustrates a sample request body to decrement credits.
//
// ```json
//
//	{
//	  "entry_type": "decrement",
//	  "amount": 20,
//	  "description": "Removing excess credits"
//	}
//
// ```
//
// ## Changing credits expiry
//
// If you'd like to change when existing credits expire, you should create a ledger
// entry of type `expiration_change`. For this entry, the required parameter
// `expiry_date` identifies the _originating_ block, and the required parameter
// `target_expiry_date` identifies when the transferred credits should now expire.
// A new credit block will be created with expiry date `target_expiry_date`, with
// the same cost basis data as the original credit block, if present.
//
// Note that the balance of the block with the given `expiry_date` must be at least
// equal to the desired transfer amount determined by the `amount` parameter.
//
// The following snippet illustrates a sample request body to extend the expiration
// date of credits by one year:
//
// ```json
//
//	{
//	  "entry_type": "expiration_change",
//	  "amount": 10,
//	  "expiry_date": "2022-12-28",
//	  "block_id": "UiUhFWeLHPrBY4Ad",
//	  "target_expiry_date": "2023-12-28",
//	  "description": "Extending credit validity"
//	}
//
// ```
//
// ## Voiding credits
//
// If you'd like to void a credit block, create a ledger entry of type `void`. For
// this entry, `block_id` is required to identify the block, and `amount` indicates
// how many credits to void, up to the block's initial balance. Pass in a
// `void_reason` of `refund` if the void is due to a refund.
//
// ## Amendment
//
// If you'd like to undo a decrement on a credit block, create a ledger entry of
// type `amendment`. For this entry, `block_id` is required to identify the block
// that was originally decremented from, and `amount` indicates how many credits to
// return to the customer, up to the block's initial balance.
func (r *CustomerCreditLedgerService) NewEntryByExteralID(ctx context.Context, externalCustomerID string, body CustomerCreditLedgerNewEntryByExteralIDParams, opts ...option.RequestOption) (res *CustomerCreditLedgerNewEntryByExteralIDResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := fmt.Sprintf("customers/external_customer_id/%s/credits/ledger_entry", externalCustomerID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return
}

// The credits ledger provides _auditing_ functionality over Orb's credits system
// with a list of actions that have taken place to modify a customer's credit
// balance. This [paginated endpoint](../reference/pagination) lists these entries,
// starting from the most recent ledger entry.
//
// More details on using Orb's real-time credit feature are
// [here](../guides/product-catalog/prepurchase.md).
//
// There are four major types of modifications to credit balance, detailed below.
//
// ## Increment
//
// Credits (which optionally expire on a future date) can be added via the API
// ([Add Ledger Entry](create-ledger-entry)). The ledger entry for such an action
// will always contain the total eligible starting and ending balance for the
// customer at the time the entry was added to the ledger.
//
// ## Decrement
//
// Deductions can occur as a result of an API call to create a ledger entry (see
// [Add Ledger Entry](create-ledger-entry)), or automatically as a result of
// incurring usage. Both ledger entries present the `decrement` entry type.
//
// As usage for a customer is reported into Orb, credits may be deducted according
// to the customer's plan configuration. An automated deduction of this type will
// result in a ledger entry, also with a starting and ending balance. In order to
// provide better tracing capabilities for automatic deductions, Orb always
// associates each automatic deduction with the `event_id` at the time of
// ingestion, used to pinpoint _why_ credit deduction took place and to ensure that
// credits are never deducted without an associated usage event.
//
// By default, Orb uses an algorithm that automatically deducts from the _soonest
// expiring credit block_ first in order to ensure that all credits are utilized
// appropriately. As an example, if trial credits with an expiration date of 2
// weeks from now are present for a customer, they will be used before any
// deductions take place from a non-expiring credit block.
//
// If there are multiple blocks with the same expiration date, Orb will deduct from
// the block with the _lower cost basis_ first (ex. trial credits with a $0 cost
// basis before paid credits with a $5.00 cost basis).
//
// It's also possible for a single usage event's deduction to _span_ credit blocks.
// In this case, Orb will deduct from the next block, ending at the credit block
// which consists of unexpiring credits. Each of these deductions will lead to a
// _separate_ ledger entry, one per credit block that is deducted from. By default,
// the customer's total credit balance in Orb can be negative as a result of a
// decrement.
//
// ## Expiration change
//
// The expiry of credits can be changed as a result of the API (See
// [Add Ledger Entry](create-ledger-entry)). This will create a ledger entry that
// specifies the balance as well as the initial and target expiry dates.
//
// Note that for this entry type, `starting_balance` will equal `ending_balance`,
// and the `amount` represents the balance transferred. The credit block linked to
// the ledger entry is the source credit block from which there was an expiration
// change
//
// ## Credits expiry
//
// When a set of credits expire on pre-set expiration date, the customer's balance
// automatically reflects this change and adds an entry to the ledger indicating
// this event. Note that credit expiry should always happen close to a date
// boundary in the customer's timezone.
//
// ## Void initiated
//
// Credit blocks can be voided via the API. The `amount` on this entry corresponds
// to the number of credits that were remaining in the block at time of void.
// `void_reason` will be populated if the void is created with a reason.
//
// ## Void
//
// When a set of credits is voided, the customer's balance automatically reflects
// this change and adds an entry to the ledger indicating this event.
//
// ## Amendment
//
// When credits are added to a customer's balance as a result of a correction, this
// entry will be added to the ledger to indicate the adjustment of credits.
func (r *CustomerCreditLedgerService) ListByExternalID(ctx context.Context, externalCustomerID string, query CustomerCreditLedgerListByExternalIDParams, opts ...option.RequestOption) (res *shared.Page[CustomerCreditLedgerListByExternalIDResponse], err error) {
	var raw *http.Response
	opts = append(r.Options, opts...)
	opts = append([]option.RequestOption{option.WithResponseInto(&raw)}, opts...)
	path := fmt.Sprintf("customers/external_customer_id/%s/credits/ledger", externalCustomerID)
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

// The credits ledger provides _auditing_ functionality over Orb's credits system
// with a list of actions that have taken place to modify a customer's credit
// balance. This [paginated endpoint](../reference/pagination) lists these entries,
// starting from the most recent ledger entry.
//
// More details on using Orb's real-time credit feature are
// [here](../guides/product-catalog/prepurchase.md).
//
// There are four major types of modifications to credit balance, detailed below.
//
// ## Increment
//
// Credits (which optionally expire on a future date) can be added via the API
// ([Add Ledger Entry](create-ledger-entry)). The ledger entry for such an action
// will always contain the total eligible starting and ending balance for the
// customer at the time the entry was added to the ledger.
//
// ## Decrement
//
// Deductions can occur as a result of an API call to create a ledger entry (see
// [Add Ledger Entry](create-ledger-entry)), or automatically as a result of
// incurring usage. Both ledger entries present the `decrement` entry type.
//
// As usage for a customer is reported into Orb, credits may be deducted according
// to the customer's plan configuration. An automated deduction of this type will
// result in a ledger entry, also with a starting and ending balance. In order to
// provide better tracing capabilities for automatic deductions, Orb always
// associates each automatic deduction with the `event_id` at the time of
// ingestion, used to pinpoint _why_ credit deduction took place and to ensure that
// credits are never deducted without an associated usage event.
//
// By default, Orb uses an algorithm that automatically deducts from the _soonest
// expiring credit block_ first in order to ensure that all credits are utilized
// appropriately. As an example, if trial credits with an expiration date of 2
// weeks from now are present for a customer, they will be used before any
// deductions take place from a non-expiring credit block.
//
// If there are multiple blocks with the same expiration date, Orb will deduct from
// the block with the _lower cost basis_ first (ex. trial credits with a $0 cost
// basis before paid credits with a $5.00 cost basis).
//
// It's also possible for a single usage event's deduction to _span_ credit blocks.
// In this case, Orb will deduct from the next block, ending at the credit block
// which consists of unexpiring credits. Each of these deductions will lead to a
// _separate_ ledger entry, one per credit block that is deducted from. By default,
// the customer's total credit balance in Orb can be negative as a result of a
// decrement.
//
// ## Expiration change
//
// The expiry of credits can be changed as a result of the API (See
// [Add Ledger Entry](create-ledger-entry)). This will create a ledger entry that
// specifies the balance as well as the initial and target expiry dates.
//
// Note that for this entry type, `starting_balance` will equal `ending_balance`,
// and the `amount` represents the balance transferred. The credit block linked to
// the ledger entry is the source credit block from which there was an expiration
// change
//
// ## Credits expiry
//
// When a set of credits expire on pre-set expiration date, the customer's balance
// automatically reflects this change and adds an entry to the ledger indicating
// this event. Note that credit expiry should always happen close to a date
// boundary in the customer's timezone.
//
// ## Void initiated
//
// Credit blocks can be voided via the API. The `amount` on this entry corresponds
// to the number of credits that were remaining in the block at time of void.
// `void_reason` will be populated if the void is created with a reason.
//
// ## Void
//
// When a set of credits is voided, the customer's balance automatically reflects
// this change and adds an entry to the ledger indicating this event.
//
// ## Amendment
//
// When credits are added to a customer's balance as a result of a correction, this
// entry will be added to the ledger to indicate the adjustment of credits.
func (r *CustomerCreditLedgerService) ListByExternalIDAutoPaging(ctx context.Context, externalCustomerID string, query CustomerCreditLedgerListByExternalIDParams, opts ...option.RequestOption) *shared.PageAutoPager[CustomerCreditLedgerListByExternalIDResponse] {
	return shared.NewPageAutoPager(r.ListByExternalID(ctx, externalCustomerID, query, opts...))
}

// The [Credit Ledger Entry resource](/guides/product-catalog/prepurchase) models
// prepaid credits within Orb.
//
// Union satisfied by [CustomerCreditLedgerListResponseIncrementLedgerEntry],
// [CustomerCreditLedgerListResponseDecrementLedgerEntry],
// [CustomerCreditLedgerListResponseExpirationChangeLedgerEntry],
// [CustomerCreditLedgerListResponseCreditBlockExpiryLedgerEntry],
// [CustomerCreditLedgerListResponseVoidLedgerEntry],
// [CustomerCreditLedgerListResponseVoidInitiatedLedgerEntry] or
// [CustomerCreditLedgerListResponseAmendmentLedgerEntry].
type CustomerCreditLedgerListResponse interface {
	implementsCustomerCreditLedgerListResponse()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*CustomerCreditLedgerListResponse)(nil)).Elem(),
		"entry_type",
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			DiscriminatorValue: "\"increment\"",
			Type:               reflect.TypeOf(CustomerCreditLedgerListResponseIncrementLedgerEntry{}),
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			DiscriminatorValue: "\"decrement\"",
			Type:               reflect.TypeOf(CustomerCreditLedgerListResponseDecrementLedgerEntry{}),
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			DiscriminatorValue: "\"expiration_change\"",
			Type:               reflect.TypeOf(CustomerCreditLedgerListResponseExpirationChangeLedgerEntry{}),
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			DiscriminatorValue: "\"credit_block_expiry\"",
			Type:               reflect.TypeOf(CustomerCreditLedgerListResponseCreditBlockExpiryLedgerEntry{}),
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			DiscriminatorValue: "\"void\"",
			Type:               reflect.TypeOf(CustomerCreditLedgerListResponseVoidLedgerEntry{}),
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			DiscriminatorValue: "\"void_initiated\"",
			Type:               reflect.TypeOf(CustomerCreditLedgerListResponseVoidInitiatedLedgerEntry{}),
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			DiscriminatorValue: "\"amendment\"",
			Type:               reflect.TypeOf(CustomerCreditLedgerListResponseAmendmentLedgerEntry{}),
		},
	)
}

type CustomerCreditLedgerListResponseIncrementLedgerEntry struct {
	ID                   string                                                          `json:"id,required"`
	Amount               float64                                                         `json:"amount,required"`
	CreatedAt            time.Time                                                       `json:"created_at,required" format:"date-time"`
	CreditBlock          CustomerCreditLedgerListResponseIncrementLedgerEntryCreditBlock `json:"credit_block,required"`
	Customer             CustomerCreditLedgerListResponseIncrementLedgerEntryCustomer    `json:"customer,required"`
	Description          string                                                          `json:"description,required,nullable"`
	EndingBalance        float64                                                         `json:"ending_balance,required"`
	EntryStatus          CustomerCreditLedgerListResponseIncrementLedgerEntryEntryStatus `json:"entry_status,required"`
	EntryType            CustomerCreditLedgerListResponseIncrementLedgerEntryEntryType   `json:"entry_type,required"`
	LedgerSequenceNumber int64                                                           `json:"ledger_sequence_number,required"`
	Metadata             map[string]string                                               `json:"metadata,required"`
	StartingBalance      float64                                                         `json:"starting_balance,required"`
	JSON                 customerCreditLedgerListResponseIncrementLedgerEntryJSON
}

// customerCreditLedgerListResponseIncrementLedgerEntryJSON contains the JSON
// metadata for the struct [CustomerCreditLedgerListResponseIncrementLedgerEntry]
type customerCreditLedgerListResponseIncrementLedgerEntryJSON struct {
	ID                   apijson.Field
	Amount               apijson.Field
	CreatedAt            apijson.Field
	CreditBlock          apijson.Field
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

func (r *CustomerCreditLedgerListResponseIncrementLedgerEntry) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r CustomerCreditLedgerListResponseIncrementLedgerEntry) implementsCustomerCreditLedgerListResponse() {
}

type CustomerCreditLedgerListResponseIncrementLedgerEntryCreditBlock struct {
	ID               string    `json:"id,required"`
	ExpiryDate       time.Time `json:"expiry_date,required,nullable" format:"date-time"`
	PerUnitCostBasis string    `json:"per_unit_cost_basis,required,nullable"`
	JSON             customerCreditLedgerListResponseIncrementLedgerEntryCreditBlockJSON
}

// customerCreditLedgerListResponseIncrementLedgerEntryCreditBlockJSON contains the
// JSON metadata for the struct
// [CustomerCreditLedgerListResponseIncrementLedgerEntryCreditBlock]
type customerCreditLedgerListResponseIncrementLedgerEntryCreditBlockJSON struct {
	ID               apijson.Field
	ExpiryDate       apijson.Field
	PerUnitCostBasis apijson.Field
	raw              string
	ExtraFields      map[string]apijson.Field
}

func (r *CustomerCreditLedgerListResponseIncrementLedgerEntryCreditBlock) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type CustomerCreditLedgerListResponseIncrementLedgerEntryCustomer struct {
	ID                 string `json:"id,required"`
	ExternalCustomerID string `json:"external_customer_id,required,nullable"`
	JSON               customerCreditLedgerListResponseIncrementLedgerEntryCustomerJSON
}

// customerCreditLedgerListResponseIncrementLedgerEntryCustomerJSON contains the
// JSON metadata for the struct
// [CustomerCreditLedgerListResponseIncrementLedgerEntryCustomer]
type customerCreditLedgerListResponseIncrementLedgerEntryCustomerJSON struct {
	ID                 apijson.Field
	ExternalCustomerID apijson.Field
	raw                string
	ExtraFields        map[string]apijson.Field
}

func (r *CustomerCreditLedgerListResponseIncrementLedgerEntryCustomer) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type CustomerCreditLedgerListResponseIncrementLedgerEntryEntryStatus string

const (
	CustomerCreditLedgerListResponseIncrementLedgerEntryEntryStatusCommitted CustomerCreditLedgerListResponseIncrementLedgerEntryEntryStatus = "committed"
	CustomerCreditLedgerListResponseIncrementLedgerEntryEntryStatusPending   CustomerCreditLedgerListResponseIncrementLedgerEntryEntryStatus = "pending"
)

type CustomerCreditLedgerListResponseIncrementLedgerEntryEntryType string

const (
	CustomerCreditLedgerListResponseIncrementLedgerEntryEntryTypeIncrement CustomerCreditLedgerListResponseIncrementLedgerEntryEntryType = "increment"
)

type CustomerCreditLedgerListResponseDecrementLedgerEntry struct {
	ID                   string                                                          `json:"id,required"`
	Amount               float64                                                         `json:"amount,required"`
	CreatedAt            time.Time                                                       `json:"created_at,required" format:"date-time"`
	CreditBlock          CustomerCreditLedgerListResponseDecrementLedgerEntryCreditBlock `json:"credit_block,required"`
	Customer             CustomerCreditLedgerListResponseDecrementLedgerEntryCustomer    `json:"customer,required"`
	Description          string                                                          `json:"description,required,nullable"`
	EndingBalance        float64                                                         `json:"ending_balance,required"`
	EntryStatus          CustomerCreditLedgerListResponseDecrementLedgerEntryEntryStatus `json:"entry_status,required"`
	EntryType            CustomerCreditLedgerListResponseDecrementLedgerEntryEntryType   `json:"entry_type,required"`
	LedgerSequenceNumber int64                                                           `json:"ledger_sequence_number,required"`
	Metadata             map[string]string                                               `json:"metadata,required"`
	StartingBalance      float64                                                         `json:"starting_balance,required"`
	EventID              string                                                          `json:"event_id,nullable"`
	InvoiceID            string                                                          `json:"invoice_id,nullable"`
	PriceID              string                                                          `json:"price_id,nullable"`
	JSON                 customerCreditLedgerListResponseDecrementLedgerEntryJSON
}

// customerCreditLedgerListResponseDecrementLedgerEntryJSON contains the JSON
// metadata for the struct [CustomerCreditLedgerListResponseDecrementLedgerEntry]
type customerCreditLedgerListResponseDecrementLedgerEntryJSON struct {
	ID                   apijson.Field
	Amount               apijson.Field
	CreatedAt            apijson.Field
	CreditBlock          apijson.Field
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

func (r *CustomerCreditLedgerListResponseDecrementLedgerEntry) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r CustomerCreditLedgerListResponseDecrementLedgerEntry) implementsCustomerCreditLedgerListResponse() {
}

type CustomerCreditLedgerListResponseDecrementLedgerEntryCreditBlock struct {
	ID               string    `json:"id,required"`
	ExpiryDate       time.Time `json:"expiry_date,required,nullable" format:"date-time"`
	PerUnitCostBasis string    `json:"per_unit_cost_basis,required,nullable"`
	JSON             customerCreditLedgerListResponseDecrementLedgerEntryCreditBlockJSON
}

// customerCreditLedgerListResponseDecrementLedgerEntryCreditBlockJSON contains the
// JSON metadata for the struct
// [CustomerCreditLedgerListResponseDecrementLedgerEntryCreditBlock]
type customerCreditLedgerListResponseDecrementLedgerEntryCreditBlockJSON struct {
	ID               apijson.Field
	ExpiryDate       apijson.Field
	PerUnitCostBasis apijson.Field
	raw              string
	ExtraFields      map[string]apijson.Field
}

func (r *CustomerCreditLedgerListResponseDecrementLedgerEntryCreditBlock) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type CustomerCreditLedgerListResponseDecrementLedgerEntryCustomer struct {
	ID                 string `json:"id,required"`
	ExternalCustomerID string `json:"external_customer_id,required,nullable"`
	JSON               customerCreditLedgerListResponseDecrementLedgerEntryCustomerJSON
}

// customerCreditLedgerListResponseDecrementLedgerEntryCustomerJSON contains the
// JSON metadata for the struct
// [CustomerCreditLedgerListResponseDecrementLedgerEntryCustomer]
type customerCreditLedgerListResponseDecrementLedgerEntryCustomerJSON struct {
	ID                 apijson.Field
	ExternalCustomerID apijson.Field
	raw                string
	ExtraFields        map[string]apijson.Field
}

func (r *CustomerCreditLedgerListResponseDecrementLedgerEntryCustomer) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type CustomerCreditLedgerListResponseDecrementLedgerEntryEntryStatus string

const (
	CustomerCreditLedgerListResponseDecrementLedgerEntryEntryStatusCommitted CustomerCreditLedgerListResponseDecrementLedgerEntryEntryStatus = "committed"
	CustomerCreditLedgerListResponseDecrementLedgerEntryEntryStatusPending   CustomerCreditLedgerListResponseDecrementLedgerEntryEntryStatus = "pending"
)

type CustomerCreditLedgerListResponseDecrementLedgerEntryEntryType string

const (
	CustomerCreditLedgerListResponseDecrementLedgerEntryEntryTypeDecrement CustomerCreditLedgerListResponseDecrementLedgerEntryEntryType = "decrement"
)

type CustomerCreditLedgerListResponseExpirationChangeLedgerEntry struct {
	ID                   string                                                                 `json:"id,required"`
	Amount               float64                                                                `json:"amount,required"`
	CreatedAt            time.Time                                                              `json:"created_at,required" format:"date-time"`
	CreditBlock          CustomerCreditLedgerListResponseExpirationChangeLedgerEntryCreditBlock `json:"credit_block,required"`
	Customer             CustomerCreditLedgerListResponseExpirationChangeLedgerEntryCustomer    `json:"customer,required"`
	Description          string                                                                 `json:"description,required,nullable"`
	EndingBalance        float64                                                                `json:"ending_balance,required"`
	EntryStatus          CustomerCreditLedgerListResponseExpirationChangeLedgerEntryEntryStatus `json:"entry_status,required"`
	EntryType            CustomerCreditLedgerListResponseExpirationChangeLedgerEntryEntryType   `json:"entry_type,required"`
	LedgerSequenceNumber int64                                                                  `json:"ledger_sequence_number,required"`
	Metadata             map[string]string                                                      `json:"metadata,required"`
	NewBlockExpiryDate   time.Time                                                              `json:"new_block_expiry_date,required" format:"date-time"`
	StartingBalance      float64                                                                `json:"starting_balance,required"`
	JSON                 customerCreditLedgerListResponseExpirationChangeLedgerEntryJSON
}

// customerCreditLedgerListResponseExpirationChangeLedgerEntryJSON contains the
// JSON metadata for the struct
// [CustomerCreditLedgerListResponseExpirationChangeLedgerEntry]
type customerCreditLedgerListResponseExpirationChangeLedgerEntryJSON struct {
	ID                   apijson.Field
	Amount               apijson.Field
	CreatedAt            apijson.Field
	CreditBlock          apijson.Field
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

func (r *CustomerCreditLedgerListResponseExpirationChangeLedgerEntry) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r CustomerCreditLedgerListResponseExpirationChangeLedgerEntry) implementsCustomerCreditLedgerListResponse() {
}

type CustomerCreditLedgerListResponseExpirationChangeLedgerEntryCreditBlock struct {
	ID               string    `json:"id,required"`
	ExpiryDate       time.Time `json:"expiry_date,required,nullable" format:"date-time"`
	PerUnitCostBasis string    `json:"per_unit_cost_basis,required,nullable"`
	JSON             customerCreditLedgerListResponseExpirationChangeLedgerEntryCreditBlockJSON
}

// customerCreditLedgerListResponseExpirationChangeLedgerEntryCreditBlockJSON
// contains the JSON metadata for the struct
// [CustomerCreditLedgerListResponseExpirationChangeLedgerEntryCreditBlock]
type customerCreditLedgerListResponseExpirationChangeLedgerEntryCreditBlockJSON struct {
	ID               apijson.Field
	ExpiryDate       apijson.Field
	PerUnitCostBasis apijson.Field
	raw              string
	ExtraFields      map[string]apijson.Field
}

func (r *CustomerCreditLedgerListResponseExpirationChangeLedgerEntryCreditBlock) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type CustomerCreditLedgerListResponseExpirationChangeLedgerEntryCustomer struct {
	ID                 string `json:"id,required"`
	ExternalCustomerID string `json:"external_customer_id,required,nullable"`
	JSON               customerCreditLedgerListResponseExpirationChangeLedgerEntryCustomerJSON
}

// customerCreditLedgerListResponseExpirationChangeLedgerEntryCustomerJSON contains
// the JSON metadata for the struct
// [CustomerCreditLedgerListResponseExpirationChangeLedgerEntryCustomer]
type customerCreditLedgerListResponseExpirationChangeLedgerEntryCustomerJSON struct {
	ID                 apijson.Field
	ExternalCustomerID apijson.Field
	raw                string
	ExtraFields        map[string]apijson.Field
}

func (r *CustomerCreditLedgerListResponseExpirationChangeLedgerEntryCustomer) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type CustomerCreditLedgerListResponseExpirationChangeLedgerEntryEntryStatus string

const (
	CustomerCreditLedgerListResponseExpirationChangeLedgerEntryEntryStatusCommitted CustomerCreditLedgerListResponseExpirationChangeLedgerEntryEntryStatus = "committed"
	CustomerCreditLedgerListResponseExpirationChangeLedgerEntryEntryStatusPending   CustomerCreditLedgerListResponseExpirationChangeLedgerEntryEntryStatus = "pending"
)

type CustomerCreditLedgerListResponseExpirationChangeLedgerEntryEntryType string

const (
	CustomerCreditLedgerListResponseExpirationChangeLedgerEntryEntryTypeExpirationChange CustomerCreditLedgerListResponseExpirationChangeLedgerEntryEntryType = "expiration_change"
)

type CustomerCreditLedgerListResponseCreditBlockExpiryLedgerEntry struct {
	ID                   string                                                                  `json:"id,required"`
	Amount               float64                                                                 `json:"amount,required"`
	CreatedAt            time.Time                                                               `json:"created_at,required" format:"date-time"`
	CreditBlock          CustomerCreditLedgerListResponseCreditBlockExpiryLedgerEntryCreditBlock `json:"credit_block,required"`
	Customer             CustomerCreditLedgerListResponseCreditBlockExpiryLedgerEntryCustomer    `json:"customer,required"`
	Description          string                                                                  `json:"description,required,nullable"`
	EndingBalance        float64                                                                 `json:"ending_balance,required"`
	EntryStatus          CustomerCreditLedgerListResponseCreditBlockExpiryLedgerEntryEntryStatus `json:"entry_status,required"`
	EntryType            CustomerCreditLedgerListResponseCreditBlockExpiryLedgerEntryEntryType   `json:"entry_type,required"`
	LedgerSequenceNumber int64                                                                   `json:"ledger_sequence_number,required"`
	Metadata             map[string]string                                                       `json:"metadata,required"`
	StartingBalance      float64                                                                 `json:"starting_balance,required"`
	JSON                 customerCreditLedgerListResponseCreditBlockExpiryLedgerEntryJSON
}

// customerCreditLedgerListResponseCreditBlockExpiryLedgerEntryJSON contains the
// JSON metadata for the struct
// [CustomerCreditLedgerListResponseCreditBlockExpiryLedgerEntry]
type customerCreditLedgerListResponseCreditBlockExpiryLedgerEntryJSON struct {
	ID                   apijson.Field
	Amount               apijson.Field
	CreatedAt            apijson.Field
	CreditBlock          apijson.Field
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

func (r *CustomerCreditLedgerListResponseCreditBlockExpiryLedgerEntry) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r CustomerCreditLedgerListResponseCreditBlockExpiryLedgerEntry) implementsCustomerCreditLedgerListResponse() {
}

type CustomerCreditLedgerListResponseCreditBlockExpiryLedgerEntryCreditBlock struct {
	ID               string    `json:"id,required"`
	ExpiryDate       time.Time `json:"expiry_date,required,nullable" format:"date-time"`
	PerUnitCostBasis string    `json:"per_unit_cost_basis,required,nullable"`
	JSON             customerCreditLedgerListResponseCreditBlockExpiryLedgerEntryCreditBlockJSON
}

// customerCreditLedgerListResponseCreditBlockExpiryLedgerEntryCreditBlockJSON
// contains the JSON metadata for the struct
// [CustomerCreditLedgerListResponseCreditBlockExpiryLedgerEntryCreditBlock]
type customerCreditLedgerListResponseCreditBlockExpiryLedgerEntryCreditBlockJSON struct {
	ID               apijson.Field
	ExpiryDate       apijson.Field
	PerUnitCostBasis apijson.Field
	raw              string
	ExtraFields      map[string]apijson.Field
}

func (r *CustomerCreditLedgerListResponseCreditBlockExpiryLedgerEntryCreditBlock) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type CustomerCreditLedgerListResponseCreditBlockExpiryLedgerEntryCustomer struct {
	ID                 string `json:"id,required"`
	ExternalCustomerID string `json:"external_customer_id,required,nullable"`
	JSON               customerCreditLedgerListResponseCreditBlockExpiryLedgerEntryCustomerJSON
}

// customerCreditLedgerListResponseCreditBlockExpiryLedgerEntryCustomerJSON
// contains the JSON metadata for the struct
// [CustomerCreditLedgerListResponseCreditBlockExpiryLedgerEntryCustomer]
type customerCreditLedgerListResponseCreditBlockExpiryLedgerEntryCustomerJSON struct {
	ID                 apijson.Field
	ExternalCustomerID apijson.Field
	raw                string
	ExtraFields        map[string]apijson.Field
}

func (r *CustomerCreditLedgerListResponseCreditBlockExpiryLedgerEntryCustomer) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type CustomerCreditLedgerListResponseCreditBlockExpiryLedgerEntryEntryStatus string

const (
	CustomerCreditLedgerListResponseCreditBlockExpiryLedgerEntryEntryStatusCommitted CustomerCreditLedgerListResponseCreditBlockExpiryLedgerEntryEntryStatus = "committed"
	CustomerCreditLedgerListResponseCreditBlockExpiryLedgerEntryEntryStatusPending   CustomerCreditLedgerListResponseCreditBlockExpiryLedgerEntryEntryStatus = "pending"
)

type CustomerCreditLedgerListResponseCreditBlockExpiryLedgerEntryEntryType string

const (
	CustomerCreditLedgerListResponseCreditBlockExpiryLedgerEntryEntryTypeCreditBlockExpiry CustomerCreditLedgerListResponseCreditBlockExpiryLedgerEntryEntryType = "credit_block_expiry"
)

type CustomerCreditLedgerListResponseVoidLedgerEntry struct {
	ID                   string                                                     `json:"id,required"`
	Amount               float64                                                    `json:"amount,required"`
	CreatedAt            time.Time                                                  `json:"created_at,required" format:"date-time"`
	CreditBlock          CustomerCreditLedgerListResponseVoidLedgerEntryCreditBlock `json:"credit_block,required"`
	Customer             CustomerCreditLedgerListResponseVoidLedgerEntryCustomer    `json:"customer,required"`
	Description          string                                                     `json:"description,required,nullable"`
	EndingBalance        float64                                                    `json:"ending_balance,required"`
	EntryStatus          CustomerCreditLedgerListResponseVoidLedgerEntryEntryStatus `json:"entry_status,required"`
	EntryType            CustomerCreditLedgerListResponseVoidLedgerEntryEntryType   `json:"entry_type,required"`
	LedgerSequenceNumber int64                                                      `json:"ledger_sequence_number,required"`
	Metadata             map[string]string                                          `json:"metadata,required"`
	StartingBalance      float64                                                    `json:"starting_balance,required"`
	VoidAmount           float64                                                    `json:"void_amount,required"`
	VoidReason           string                                                     `json:"void_reason,required,nullable"`
	JSON                 customerCreditLedgerListResponseVoidLedgerEntryJSON
}

// customerCreditLedgerListResponseVoidLedgerEntryJSON contains the JSON metadata
// for the struct [CustomerCreditLedgerListResponseVoidLedgerEntry]
type customerCreditLedgerListResponseVoidLedgerEntryJSON struct {
	ID                   apijson.Field
	Amount               apijson.Field
	CreatedAt            apijson.Field
	CreditBlock          apijson.Field
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

func (r *CustomerCreditLedgerListResponseVoidLedgerEntry) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r CustomerCreditLedgerListResponseVoidLedgerEntry) implementsCustomerCreditLedgerListResponse() {
}

type CustomerCreditLedgerListResponseVoidLedgerEntryCreditBlock struct {
	ID               string    `json:"id,required"`
	ExpiryDate       time.Time `json:"expiry_date,required,nullable" format:"date-time"`
	PerUnitCostBasis string    `json:"per_unit_cost_basis,required,nullable"`
	JSON             customerCreditLedgerListResponseVoidLedgerEntryCreditBlockJSON
}

// customerCreditLedgerListResponseVoidLedgerEntryCreditBlockJSON contains the JSON
// metadata for the struct
// [CustomerCreditLedgerListResponseVoidLedgerEntryCreditBlock]
type customerCreditLedgerListResponseVoidLedgerEntryCreditBlockJSON struct {
	ID               apijson.Field
	ExpiryDate       apijson.Field
	PerUnitCostBasis apijson.Field
	raw              string
	ExtraFields      map[string]apijson.Field
}

func (r *CustomerCreditLedgerListResponseVoidLedgerEntryCreditBlock) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type CustomerCreditLedgerListResponseVoidLedgerEntryCustomer struct {
	ID                 string `json:"id,required"`
	ExternalCustomerID string `json:"external_customer_id,required,nullable"`
	JSON               customerCreditLedgerListResponseVoidLedgerEntryCustomerJSON
}

// customerCreditLedgerListResponseVoidLedgerEntryCustomerJSON contains the JSON
// metadata for the struct
// [CustomerCreditLedgerListResponseVoidLedgerEntryCustomer]
type customerCreditLedgerListResponseVoidLedgerEntryCustomerJSON struct {
	ID                 apijson.Field
	ExternalCustomerID apijson.Field
	raw                string
	ExtraFields        map[string]apijson.Field
}

func (r *CustomerCreditLedgerListResponseVoidLedgerEntryCustomer) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type CustomerCreditLedgerListResponseVoidLedgerEntryEntryStatus string

const (
	CustomerCreditLedgerListResponseVoidLedgerEntryEntryStatusCommitted CustomerCreditLedgerListResponseVoidLedgerEntryEntryStatus = "committed"
	CustomerCreditLedgerListResponseVoidLedgerEntryEntryStatusPending   CustomerCreditLedgerListResponseVoidLedgerEntryEntryStatus = "pending"
)

type CustomerCreditLedgerListResponseVoidLedgerEntryEntryType string

const (
	CustomerCreditLedgerListResponseVoidLedgerEntryEntryTypeVoid CustomerCreditLedgerListResponseVoidLedgerEntryEntryType = "void"
)

type CustomerCreditLedgerListResponseVoidInitiatedLedgerEntry struct {
	ID                   string                                                              `json:"id,required"`
	Amount               float64                                                             `json:"amount,required"`
	CreatedAt            time.Time                                                           `json:"created_at,required" format:"date-time"`
	CreditBlock          CustomerCreditLedgerListResponseVoidInitiatedLedgerEntryCreditBlock `json:"credit_block,required"`
	Customer             CustomerCreditLedgerListResponseVoidInitiatedLedgerEntryCustomer    `json:"customer,required"`
	Description          string                                                              `json:"description,required,nullable"`
	EndingBalance        float64                                                             `json:"ending_balance,required"`
	EntryStatus          CustomerCreditLedgerListResponseVoidInitiatedLedgerEntryEntryStatus `json:"entry_status,required"`
	EntryType            CustomerCreditLedgerListResponseVoidInitiatedLedgerEntryEntryType   `json:"entry_type,required"`
	LedgerSequenceNumber int64                                                               `json:"ledger_sequence_number,required"`
	Metadata             map[string]string                                                   `json:"metadata,required"`
	NewBlockExpiryDate   time.Time                                                           `json:"new_block_expiry_date,required" format:"date-time"`
	StartingBalance      float64                                                             `json:"starting_balance,required"`
	VoidAmount           float64                                                             `json:"void_amount,required"`
	VoidReason           string                                                              `json:"void_reason,required,nullable"`
	JSON                 customerCreditLedgerListResponseVoidInitiatedLedgerEntryJSON
}

// customerCreditLedgerListResponseVoidInitiatedLedgerEntryJSON contains the JSON
// metadata for the struct
// [CustomerCreditLedgerListResponseVoidInitiatedLedgerEntry]
type customerCreditLedgerListResponseVoidInitiatedLedgerEntryJSON struct {
	ID                   apijson.Field
	Amount               apijson.Field
	CreatedAt            apijson.Field
	CreditBlock          apijson.Field
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

func (r *CustomerCreditLedgerListResponseVoidInitiatedLedgerEntry) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r CustomerCreditLedgerListResponseVoidInitiatedLedgerEntry) implementsCustomerCreditLedgerListResponse() {
}

type CustomerCreditLedgerListResponseVoidInitiatedLedgerEntryCreditBlock struct {
	ID               string    `json:"id,required"`
	ExpiryDate       time.Time `json:"expiry_date,required,nullable" format:"date-time"`
	PerUnitCostBasis string    `json:"per_unit_cost_basis,required,nullable"`
	JSON             customerCreditLedgerListResponseVoidInitiatedLedgerEntryCreditBlockJSON
}

// customerCreditLedgerListResponseVoidInitiatedLedgerEntryCreditBlockJSON contains
// the JSON metadata for the struct
// [CustomerCreditLedgerListResponseVoidInitiatedLedgerEntryCreditBlock]
type customerCreditLedgerListResponseVoidInitiatedLedgerEntryCreditBlockJSON struct {
	ID               apijson.Field
	ExpiryDate       apijson.Field
	PerUnitCostBasis apijson.Field
	raw              string
	ExtraFields      map[string]apijson.Field
}

func (r *CustomerCreditLedgerListResponseVoidInitiatedLedgerEntryCreditBlock) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type CustomerCreditLedgerListResponseVoidInitiatedLedgerEntryCustomer struct {
	ID                 string `json:"id,required"`
	ExternalCustomerID string `json:"external_customer_id,required,nullable"`
	JSON               customerCreditLedgerListResponseVoidInitiatedLedgerEntryCustomerJSON
}

// customerCreditLedgerListResponseVoidInitiatedLedgerEntryCustomerJSON contains
// the JSON metadata for the struct
// [CustomerCreditLedgerListResponseVoidInitiatedLedgerEntryCustomer]
type customerCreditLedgerListResponseVoidInitiatedLedgerEntryCustomerJSON struct {
	ID                 apijson.Field
	ExternalCustomerID apijson.Field
	raw                string
	ExtraFields        map[string]apijson.Field
}

func (r *CustomerCreditLedgerListResponseVoidInitiatedLedgerEntryCustomer) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type CustomerCreditLedgerListResponseVoidInitiatedLedgerEntryEntryStatus string

const (
	CustomerCreditLedgerListResponseVoidInitiatedLedgerEntryEntryStatusCommitted CustomerCreditLedgerListResponseVoidInitiatedLedgerEntryEntryStatus = "committed"
	CustomerCreditLedgerListResponseVoidInitiatedLedgerEntryEntryStatusPending   CustomerCreditLedgerListResponseVoidInitiatedLedgerEntryEntryStatus = "pending"
)

type CustomerCreditLedgerListResponseVoidInitiatedLedgerEntryEntryType string

const (
	CustomerCreditLedgerListResponseVoidInitiatedLedgerEntryEntryTypeVoidInitiated CustomerCreditLedgerListResponseVoidInitiatedLedgerEntryEntryType = "void_initiated"
)

type CustomerCreditLedgerListResponseAmendmentLedgerEntry struct {
	ID                   string                                                          `json:"id,required"`
	Amount               float64                                                         `json:"amount,required"`
	CreatedAt            time.Time                                                       `json:"created_at,required" format:"date-time"`
	CreditBlock          CustomerCreditLedgerListResponseAmendmentLedgerEntryCreditBlock `json:"credit_block,required"`
	Customer             CustomerCreditLedgerListResponseAmendmentLedgerEntryCustomer    `json:"customer,required"`
	Description          string                                                          `json:"description,required,nullable"`
	EndingBalance        float64                                                         `json:"ending_balance,required"`
	EntryStatus          CustomerCreditLedgerListResponseAmendmentLedgerEntryEntryStatus `json:"entry_status,required"`
	EntryType            CustomerCreditLedgerListResponseAmendmentLedgerEntryEntryType   `json:"entry_type,required"`
	LedgerSequenceNumber int64                                                           `json:"ledger_sequence_number,required"`
	Metadata             map[string]string                                               `json:"metadata,required"`
	StartingBalance      float64                                                         `json:"starting_balance,required"`
	JSON                 customerCreditLedgerListResponseAmendmentLedgerEntryJSON
}

// customerCreditLedgerListResponseAmendmentLedgerEntryJSON contains the JSON
// metadata for the struct [CustomerCreditLedgerListResponseAmendmentLedgerEntry]
type customerCreditLedgerListResponseAmendmentLedgerEntryJSON struct {
	ID                   apijson.Field
	Amount               apijson.Field
	CreatedAt            apijson.Field
	CreditBlock          apijson.Field
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

func (r *CustomerCreditLedgerListResponseAmendmentLedgerEntry) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r CustomerCreditLedgerListResponseAmendmentLedgerEntry) implementsCustomerCreditLedgerListResponse() {
}

type CustomerCreditLedgerListResponseAmendmentLedgerEntryCreditBlock struct {
	ID               string    `json:"id,required"`
	ExpiryDate       time.Time `json:"expiry_date,required,nullable" format:"date-time"`
	PerUnitCostBasis string    `json:"per_unit_cost_basis,required,nullable"`
	JSON             customerCreditLedgerListResponseAmendmentLedgerEntryCreditBlockJSON
}

// customerCreditLedgerListResponseAmendmentLedgerEntryCreditBlockJSON contains the
// JSON metadata for the struct
// [CustomerCreditLedgerListResponseAmendmentLedgerEntryCreditBlock]
type customerCreditLedgerListResponseAmendmentLedgerEntryCreditBlockJSON struct {
	ID               apijson.Field
	ExpiryDate       apijson.Field
	PerUnitCostBasis apijson.Field
	raw              string
	ExtraFields      map[string]apijson.Field
}

func (r *CustomerCreditLedgerListResponseAmendmentLedgerEntryCreditBlock) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type CustomerCreditLedgerListResponseAmendmentLedgerEntryCustomer struct {
	ID                 string `json:"id,required"`
	ExternalCustomerID string `json:"external_customer_id,required,nullable"`
	JSON               customerCreditLedgerListResponseAmendmentLedgerEntryCustomerJSON
}

// customerCreditLedgerListResponseAmendmentLedgerEntryCustomerJSON contains the
// JSON metadata for the struct
// [CustomerCreditLedgerListResponseAmendmentLedgerEntryCustomer]
type customerCreditLedgerListResponseAmendmentLedgerEntryCustomerJSON struct {
	ID                 apijson.Field
	ExternalCustomerID apijson.Field
	raw                string
	ExtraFields        map[string]apijson.Field
}

func (r *CustomerCreditLedgerListResponseAmendmentLedgerEntryCustomer) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type CustomerCreditLedgerListResponseAmendmentLedgerEntryEntryStatus string

const (
	CustomerCreditLedgerListResponseAmendmentLedgerEntryEntryStatusCommitted CustomerCreditLedgerListResponseAmendmentLedgerEntryEntryStatus = "committed"
	CustomerCreditLedgerListResponseAmendmentLedgerEntryEntryStatusPending   CustomerCreditLedgerListResponseAmendmentLedgerEntryEntryStatus = "pending"
)

type CustomerCreditLedgerListResponseAmendmentLedgerEntryEntryType string

const (
	CustomerCreditLedgerListResponseAmendmentLedgerEntryEntryTypeAmendment CustomerCreditLedgerListResponseAmendmentLedgerEntryEntryType = "amendment"
)

// The [Credit Ledger Entry resource](/guides/product-catalog/prepurchase) models
// prepaid credits within Orb.
//
// Union satisfied by [CustomerCreditLedgerNewEntryResponseIncrementLedgerEntry],
// [CustomerCreditLedgerNewEntryResponseDecrementLedgerEntry],
// [CustomerCreditLedgerNewEntryResponseExpirationChangeLedgerEntry],
// [CustomerCreditLedgerNewEntryResponseCreditBlockExpiryLedgerEntry],
// [CustomerCreditLedgerNewEntryResponseVoidLedgerEntry],
// [CustomerCreditLedgerNewEntryResponseVoidInitiatedLedgerEntry] or
// [CustomerCreditLedgerNewEntryResponseAmendmentLedgerEntry].
type CustomerCreditLedgerNewEntryResponse interface {
	implementsCustomerCreditLedgerNewEntryResponse()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*CustomerCreditLedgerNewEntryResponse)(nil)).Elem(),
		"entry_type",
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			DiscriminatorValue: "\"increment\"",
			Type:               reflect.TypeOf(CustomerCreditLedgerNewEntryResponseIncrementLedgerEntry{}),
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			DiscriminatorValue: "\"decrement\"",
			Type:               reflect.TypeOf(CustomerCreditLedgerNewEntryResponseDecrementLedgerEntry{}),
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			DiscriminatorValue: "\"expiration_change\"",
			Type:               reflect.TypeOf(CustomerCreditLedgerNewEntryResponseExpirationChangeLedgerEntry{}),
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			DiscriminatorValue: "\"credit_block_expiry\"",
			Type:               reflect.TypeOf(CustomerCreditLedgerNewEntryResponseCreditBlockExpiryLedgerEntry{}),
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			DiscriminatorValue: "\"void\"",
			Type:               reflect.TypeOf(CustomerCreditLedgerNewEntryResponseVoidLedgerEntry{}),
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			DiscriminatorValue: "\"void_initiated\"",
			Type:               reflect.TypeOf(CustomerCreditLedgerNewEntryResponseVoidInitiatedLedgerEntry{}),
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			DiscriminatorValue: "\"amendment\"",
			Type:               reflect.TypeOf(CustomerCreditLedgerNewEntryResponseAmendmentLedgerEntry{}),
		},
	)
}

type CustomerCreditLedgerNewEntryResponseIncrementLedgerEntry struct {
	ID                   string                                                              `json:"id,required"`
	Amount               float64                                                             `json:"amount,required"`
	CreatedAt            time.Time                                                           `json:"created_at,required" format:"date-time"`
	CreditBlock          CustomerCreditLedgerNewEntryResponseIncrementLedgerEntryCreditBlock `json:"credit_block,required"`
	Customer             CustomerCreditLedgerNewEntryResponseIncrementLedgerEntryCustomer    `json:"customer,required"`
	Description          string                                                              `json:"description,required,nullable"`
	EndingBalance        float64                                                             `json:"ending_balance,required"`
	EntryStatus          CustomerCreditLedgerNewEntryResponseIncrementLedgerEntryEntryStatus `json:"entry_status,required"`
	EntryType            CustomerCreditLedgerNewEntryResponseIncrementLedgerEntryEntryType   `json:"entry_type,required"`
	LedgerSequenceNumber int64                                                               `json:"ledger_sequence_number,required"`
	Metadata             map[string]string                                                   `json:"metadata,required"`
	StartingBalance      float64                                                             `json:"starting_balance,required"`
	JSON                 customerCreditLedgerNewEntryResponseIncrementLedgerEntryJSON
}

// customerCreditLedgerNewEntryResponseIncrementLedgerEntryJSON contains the JSON
// metadata for the struct
// [CustomerCreditLedgerNewEntryResponseIncrementLedgerEntry]
type customerCreditLedgerNewEntryResponseIncrementLedgerEntryJSON struct {
	ID                   apijson.Field
	Amount               apijson.Field
	CreatedAt            apijson.Field
	CreditBlock          apijson.Field
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

func (r *CustomerCreditLedgerNewEntryResponseIncrementLedgerEntry) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r CustomerCreditLedgerNewEntryResponseIncrementLedgerEntry) implementsCustomerCreditLedgerNewEntryResponse() {
}

type CustomerCreditLedgerNewEntryResponseIncrementLedgerEntryCreditBlock struct {
	ID               string    `json:"id,required"`
	ExpiryDate       time.Time `json:"expiry_date,required,nullable" format:"date-time"`
	PerUnitCostBasis string    `json:"per_unit_cost_basis,required,nullable"`
	JSON             customerCreditLedgerNewEntryResponseIncrementLedgerEntryCreditBlockJSON
}

// customerCreditLedgerNewEntryResponseIncrementLedgerEntryCreditBlockJSON contains
// the JSON metadata for the struct
// [CustomerCreditLedgerNewEntryResponseIncrementLedgerEntryCreditBlock]
type customerCreditLedgerNewEntryResponseIncrementLedgerEntryCreditBlockJSON struct {
	ID               apijson.Field
	ExpiryDate       apijson.Field
	PerUnitCostBasis apijson.Field
	raw              string
	ExtraFields      map[string]apijson.Field
}

func (r *CustomerCreditLedgerNewEntryResponseIncrementLedgerEntryCreditBlock) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type CustomerCreditLedgerNewEntryResponseIncrementLedgerEntryCustomer struct {
	ID                 string `json:"id,required"`
	ExternalCustomerID string `json:"external_customer_id,required,nullable"`
	JSON               customerCreditLedgerNewEntryResponseIncrementLedgerEntryCustomerJSON
}

// customerCreditLedgerNewEntryResponseIncrementLedgerEntryCustomerJSON contains
// the JSON metadata for the struct
// [CustomerCreditLedgerNewEntryResponseIncrementLedgerEntryCustomer]
type customerCreditLedgerNewEntryResponseIncrementLedgerEntryCustomerJSON struct {
	ID                 apijson.Field
	ExternalCustomerID apijson.Field
	raw                string
	ExtraFields        map[string]apijson.Field
}

func (r *CustomerCreditLedgerNewEntryResponseIncrementLedgerEntryCustomer) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type CustomerCreditLedgerNewEntryResponseIncrementLedgerEntryEntryStatus string

const (
	CustomerCreditLedgerNewEntryResponseIncrementLedgerEntryEntryStatusCommitted CustomerCreditLedgerNewEntryResponseIncrementLedgerEntryEntryStatus = "committed"
	CustomerCreditLedgerNewEntryResponseIncrementLedgerEntryEntryStatusPending   CustomerCreditLedgerNewEntryResponseIncrementLedgerEntryEntryStatus = "pending"
)

type CustomerCreditLedgerNewEntryResponseIncrementLedgerEntryEntryType string

const (
	CustomerCreditLedgerNewEntryResponseIncrementLedgerEntryEntryTypeIncrement CustomerCreditLedgerNewEntryResponseIncrementLedgerEntryEntryType = "increment"
)

type CustomerCreditLedgerNewEntryResponseDecrementLedgerEntry struct {
	ID                   string                                                              `json:"id,required"`
	Amount               float64                                                             `json:"amount,required"`
	CreatedAt            time.Time                                                           `json:"created_at,required" format:"date-time"`
	CreditBlock          CustomerCreditLedgerNewEntryResponseDecrementLedgerEntryCreditBlock `json:"credit_block,required"`
	Customer             CustomerCreditLedgerNewEntryResponseDecrementLedgerEntryCustomer    `json:"customer,required"`
	Description          string                                                              `json:"description,required,nullable"`
	EndingBalance        float64                                                             `json:"ending_balance,required"`
	EntryStatus          CustomerCreditLedgerNewEntryResponseDecrementLedgerEntryEntryStatus `json:"entry_status,required"`
	EntryType            CustomerCreditLedgerNewEntryResponseDecrementLedgerEntryEntryType   `json:"entry_type,required"`
	LedgerSequenceNumber int64                                                               `json:"ledger_sequence_number,required"`
	Metadata             map[string]string                                                   `json:"metadata,required"`
	StartingBalance      float64                                                             `json:"starting_balance,required"`
	EventID              string                                                              `json:"event_id,nullable"`
	InvoiceID            string                                                              `json:"invoice_id,nullable"`
	PriceID              string                                                              `json:"price_id,nullable"`
	JSON                 customerCreditLedgerNewEntryResponseDecrementLedgerEntryJSON
}

// customerCreditLedgerNewEntryResponseDecrementLedgerEntryJSON contains the JSON
// metadata for the struct
// [CustomerCreditLedgerNewEntryResponseDecrementLedgerEntry]
type customerCreditLedgerNewEntryResponseDecrementLedgerEntryJSON struct {
	ID                   apijson.Field
	Amount               apijson.Field
	CreatedAt            apijson.Field
	CreditBlock          apijson.Field
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

func (r *CustomerCreditLedgerNewEntryResponseDecrementLedgerEntry) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r CustomerCreditLedgerNewEntryResponseDecrementLedgerEntry) implementsCustomerCreditLedgerNewEntryResponse() {
}

type CustomerCreditLedgerNewEntryResponseDecrementLedgerEntryCreditBlock struct {
	ID               string    `json:"id,required"`
	ExpiryDate       time.Time `json:"expiry_date,required,nullable" format:"date-time"`
	PerUnitCostBasis string    `json:"per_unit_cost_basis,required,nullable"`
	JSON             customerCreditLedgerNewEntryResponseDecrementLedgerEntryCreditBlockJSON
}

// customerCreditLedgerNewEntryResponseDecrementLedgerEntryCreditBlockJSON contains
// the JSON metadata for the struct
// [CustomerCreditLedgerNewEntryResponseDecrementLedgerEntryCreditBlock]
type customerCreditLedgerNewEntryResponseDecrementLedgerEntryCreditBlockJSON struct {
	ID               apijson.Field
	ExpiryDate       apijson.Field
	PerUnitCostBasis apijson.Field
	raw              string
	ExtraFields      map[string]apijson.Field
}

func (r *CustomerCreditLedgerNewEntryResponseDecrementLedgerEntryCreditBlock) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type CustomerCreditLedgerNewEntryResponseDecrementLedgerEntryCustomer struct {
	ID                 string `json:"id,required"`
	ExternalCustomerID string `json:"external_customer_id,required,nullable"`
	JSON               customerCreditLedgerNewEntryResponseDecrementLedgerEntryCustomerJSON
}

// customerCreditLedgerNewEntryResponseDecrementLedgerEntryCustomerJSON contains
// the JSON metadata for the struct
// [CustomerCreditLedgerNewEntryResponseDecrementLedgerEntryCustomer]
type customerCreditLedgerNewEntryResponseDecrementLedgerEntryCustomerJSON struct {
	ID                 apijson.Field
	ExternalCustomerID apijson.Field
	raw                string
	ExtraFields        map[string]apijson.Field
}

func (r *CustomerCreditLedgerNewEntryResponseDecrementLedgerEntryCustomer) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type CustomerCreditLedgerNewEntryResponseDecrementLedgerEntryEntryStatus string

const (
	CustomerCreditLedgerNewEntryResponseDecrementLedgerEntryEntryStatusCommitted CustomerCreditLedgerNewEntryResponseDecrementLedgerEntryEntryStatus = "committed"
	CustomerCreditLedgerNewEntryResponseDecrementLedgerEntryEntryStatusPending   CustomerCreditLedgerNewEntryResponseDecrementLedgerEntryEntryStatus = "pending"
)

type CustomerCreditLedgerNewEntryResponseDecrementLedgerEntryEntryType string

const (
	CustomerCreditLedgerNewEntryResponseDecrementLedgerEntryEntryTypeDecrement CustomerCreditLedgerNewEntryResponseDecrementLedgerEntryEntryType = "decrement"
)

type CustomerCreditLedgerNewEntryResponseExpirationChangeLedgerEntry struct {
	ID                   string                                                                     `json:"id,required"`
	Amount               float64                                                                    `json:"amount,required"`
	CreatedAt            time.Time                                                                  `json:"created_at,required" format:"date-time"`
	CreditBlock          CustomerCreditLedgerNewEntryResponseExpirationChangeLedgerEntryCreditBlock `json:"credit_block,required"`
	Customer             CustomerCreditLedgerNewEntryResponseExpirationChangeLedgerEntryCustomer    `json:"customer,required"`
	Description          string                                                                     `json:"description,required,nullable"`
	EndingBalance        float64                                                                    `json:"ending_balance,required"`
	EntryStatus          CustomerCreditLedgerNewEntryResponseExpirationChangeLedgerEntryEntryStatus `json:"entry_status,required"`
	EntryType            CustomerCreditLedgerNewEntryResponseExpirationChangeLedgerEntryEntryType   `json:"entry_type,required"`
	LedgerSequenceNumber int64                                                                      `json:"ledger_sequence_number,required"`
	Metadata             map[string]string                                                          `json:"metadata,required"`
	NewBlockExpiryDate   time.Time                                                                  `json:"new_block_expiry_date,required" format:"date-time"`
	StartingBalance      float64                                                                    `json:"starting_balance,required"`
	JSON                 customerCreditLedgerNewEntryResponseExpirationChangeLedgerEntryJSON
}

// customerCreditLedgerNewEntryResponseExpirationChangeLedgerEntryJSON contains the
// JSON metadata for the struct
// [CustomerCreditLedgerNewEntryResponseExpirationChangeLedgerEntry]
type customerCreditLedgerNewEntryResponseExpirationChangeLedgerEntryJSON struct {
	ID                   apijson.Field
	Amount               apijson.Field
	CreatedAt            apijson.Field
	CreditBlock          apijson.Field
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

func (r *CustomerCreditLedgerNewEntryResponseExpirationChangeLedgerEntry) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r CustomerCreditLedgerNewEntryResponseExpirationChangeLedgerEntry) implementsCustomerCreditLedgerNewEntryResponse() {
}

type CustomerCreditLedgerNewEntryResponseExpirationChangeLedgerEntryCreditBlock struct {
	ID               string    `json:"id,required"`
	ExpiryDate       time.Time `json:"expiry_date,required,nullable" format:"date-time"`
	PerUnitCostBasis string    `json:"per_unit_cost_basis,required,nullable"`
	JSON             customerCreditLedgerNewEntryResponseExpirationChangeLedgerEntryCreditBlockJSON
}

// customerCreditLedgerNewEntryResponseExpirationChangeLedgerEntryCreditBlockJSON
// contains the JSON metadata for the struct
// [CustomerCreditLedgerNewEntryResponseExpirationChangeLedgerEntryCreditBlock]
type customerCreditLedgerNewEntryResponseExpirationChangeLedgerEntryCreditBlockJSON struct {
	ID               apijson.Field
	ExpiryDate       apijson.Field
	PerUnitCostBasis apijson.Field
	raw              string
	ExtraFields      map[string]apijson.Field
}

func (r *CustomerCreditLedgerNewEntryResponseExpirationChangeLedgerEntryCreditBlock) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type CustomerCreditLedgerNewEntryResponseExpirationChangeLedgerEntryCustomer struct {
	ID                 string `json:"id,required"`
	ExternalCustomerID string `json:"external_customer_id,required,nullable"`
	JSON               customerCreditLedgerNewEntryResponseExpirationChangeLedgerEntryCustomerJSON
}

// customerCreditLedgerNewEntryResponseExpirationChangeLedgerEntryCustomerJSON
// contains the JSON metadata for the struct
// [CustomerCreditLedgerNewEntryResponseExpirationChangeLedgerEntryCustomer]
type customerCreditLedgerNewEntryResponseExpirationChangeLedgerEntryCustomerJSON struct {
	ID                 apijson.Field
	ExternalCustomerID apijson.Field
	raw                string
	ExtraFields        map[string]apijson.Field
}

func (r *CustomerCreditLedgerNewEntryResponseExpirationChangeLedgerEntryCustomer) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type CustomerCreditLedgerNewEntryResponseExpirationChangeLedgerEntryEntryStatus string

const (
	CustomerCreditLedgerNewEntryResponseExpirationChangeLedgerEntryEntryStatusCommitted CustomerCreditLedgerNewEntryResponseExpirationChangeLedgerEntryEntryStatus = "committed"
	CustomerCreditLedgerNewEntryResponseExpirationChangeLedgerEntryEntryStatusPending   CustomerCreditLedgerNewEntryResponseExpirationChangeLedgerEntryEntryStatus = "pending"
)

type CustomerCreditLedgerNewEntryResponseExpirationChangeLedgerEntryEntryType string

const (
	CustomerCreditLedgerNewEntryResponseExpirationChangeLedgerEntryEntryTypeExpirationChange CustomerCreditLedgerNewEntryResponseExpirationChangeLedgerEntryEntryType = "expiration_change"
)

type CustomerCreditLedgerNewEntryResponseCreditBlockExpiryLedgerEntry struct {
	ID                   string                                                                      `json:"id,required"`
	Amount               float64                                                                     `json:"amount,required"`
	CreatedAt            time.Time                                                                   `json:"created_at,required" format:"date-time"`
	CreditBlock          CustomerCreditLedgerNewEntryResponseCreditBlockExpiryLedgerEntryCreditBlock `json:"credit_block,required"`
	Customer             CustomerCreditLedgerNewEntryResponseCreditBlockExpiryLedgerEntryCustomer    `json:"customer,required"`
	Description          string                                                                      `json:"description,required,nullable"`
	EndingBalance        float64                                                                     `json:"ending_balance,required"`
	EntryStatus          CustomerCreditLedgerNewEntryResponseCreditBlockExpiryLedgerEntryEntryStatus `json:"entry_status,required"`
	EntryType            CustomerCreditLedgerNewEntryResponseCreditBlockExpiryLedgerEntryEntryType   `json:"entry_type,required"`
	LedgerSequenceNumber int64                                                                       `json:"ledger_sequence_number,required"`
	Metadata             map[string]string                                                           `json:"metadata,required"`
	StartingBalance      float64                                                                     `json:"starting_balance,required"`
	JSON                 customerCreditLedgerNewEntryResponseCreditBlockExpiryLedgerEntryJSON
}

// customerCreditLedgerNewEntryResponseCreditBlockExpiryLedgerEntryJSON contains
// the JSON metadata for the struct
// [CustomerCreditLedgerNewEntryResponseCreditBlockExpiryLedgerEntry]
type customerCreditLedgerNewEntryResponseCreditBlockExpiryLedgerEntryJSON struct {
	ID                   apijson.Field
	Amount               apijson.Field
	CreatedAt            apijson.Field
	CreditBlock          apijson.Field
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

func (r *CustomerCreditLedgerNewEntryResponseCreditBlockExpiryLedgerEntry) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r CustomerCreditLedgerNewEntryResponseCreditBlockExpiryLedgerEntry) implementsCustomerCreditLedgerNewEntryResponse() {
}

type CustomerCreditLedgerNewEntryResponseCreditBlockExpiryLedgerEntryCreditBlock struct {
	ID               string    `json:"id,required"`
	ExpiryDate       time.Time `json:"expiry_date,required,nullable" format:"date-time"`
	PerUnitCostBasis string    `json:"per_unit_cost_basis,required,nullable"`
	JSON             customerCreditLedgerNewEntryResponseCreditBlockExpiryLedgerEntryCreditBlockJSON
}

// customerCreditLedgerNewEntryResponseCreditBlockExpiryLedgerEntryCreditBlockJSON
// contains the JSON metadata for the struct
// [CustomerCreditLedgerNewEntryResponseCreditBlockExpiryLedgerEntryCreditBlock]
type customerCreditLedgerNewEntryResponseCreditBlockExpiryLedgerEntryCreditBlockJSON struct {
	ID               apijson.Field
	ExpiryDate       apijson.Field
	PerUnitCostBasis apijson.Field
	raw              string
	ExtraFields      map[string]apijson.Field
}

func (r *CustomerCreditLedgerNewEntryResponseCreditBlockExpiryLedgerEntryCreditBlock) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type CustomerCreditLedgerNewEntryResponseCreditBlockExpiryLedgerEntryCustomer struct {
	ID                 string `json:"id,required"`
	ExternalCustomerID string `json:"external_customer_id,required,nullable"`
	JSON               customerCreditLedgerNewEntryResponseCreditBlockExpiryLedgerEntryCustomerJSON
}

// customerCreditLedgerNewEntryResponseCreditBlockExpiryLedgerEntryCustomerJSON
// contains the JSON metadata for the struct
// [CustomerCreditLedgerNewEntryResponseCreditBlockExpiryLedgerEntryCustomer]
type customerCreditLedgerNewEntryResponseCreditBlockExpiryLedgerEntryCustomerJSON struct {
	ID                 apijson.Field
	ExternalCustomerID apijson.Field
	raw                string
	ExtraFields        map[string]apijson.Field
}

func (r *CustomerCreditLedgerNewEntryResponseCreditBlockExpiryLedgerEntryCustomer) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type CustomerCreditLedgerNewEntryResponseCreditBlockExpiryLedgerEntryEntryStatus string

const (
	CustomerCreditLedgerNewEntryResponseCreditBlockExpiryLedgerEntryEntryStatusCommitted CustomerCreditLedgerNewEntryResponseCreditBlockExpiryLedgerEntryEntryStatus = "committed"
	CustomerCreditLedgerNewEntryResponseCreditBlockExpiryLedgerEntryEntryStatusPending   CustomerCreditLedgerNewEntryResponseCreditBlockExpiryLedgerEntryEntryStatus = "pending"
)

type CustomerCreditLedgerNewEntryResponseCreditBlockExpiryLedgerEntryEntryType string

const (
	CustomerCreditLedgerNewEntryResponseCreditBlockExpiryLedgerEntryEntryTypeCreditBlockExpiry CustomerCreditLedgerNewEntryResponseCreditBlockExpiryLedgerEntryEntryType = "credit_block_expiry"
)

type CustomerCreditLedgerNewEntryResponseVoidLedgerEntry struct {
	ID                   string                                                         `json:"id,required"`
	Amount               float64                                                        `json:"amount,required"`
	CreatedAt            time.Time                                                      `json:"created_at,required" format:"date-time"`
	CreditBlock          CustomerCreditLedgerNewEntryResponseVoidLedgerEntryCreditBlock `json:"credit_block,required"`
	Customer             CustomerCreditLedgerNewEntryResponseVoidLedgerEntryCustomer    `json:"customer,required"`
	Description          string                                                         `json:"description,required,nullable"`
	EndingBalance        float64                                                        `json:"ending_balance,required"`
	EntryStatus          CustomerCreditLedgerNewEntryResponseVoidLedgerEntryEntryStatus `json:"entry_status,required"`
	EntryType            CustomerCreditLedgerNewEntryResponseVoidLedgerEntryEntryType   `json:"entry_type,required"`
	LedgerSequenceNumber int64                                                          `json:"ledger_sequence_number,required"`
	Metadata             map[string]string                                              `json:"metadata,required"`
	StartingBalance      float64                                                        `json:"starting_balance,required"`
	VoidAmount           float64                                                        `json:"void_amount,required"`
	VoidReason           string                                                         `json:"void_reason,required,nullable"`
	JSON                 customerCreditLedgerNewEntryResponseVoidLedgerEntryJSON
}

// customerCreditLedgerNewEntryResponseVoidLedgerEntryJSON contains the JSON
// metadata for the struct [CustomerCreditLedgerNewEntryResponseVoidLedgerEntry]
type customerCreditLedgerNewEntryResponseVoidLedgerEntryJSON struct {
	ID                   apijson.Field
	Amount               apijson.Field
	CreatedAt            apijson.Field
	CreditBlock          apijson.Field
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

func (r *CustomerCreditLedgerNewEntryResponseVoidLedgerEntry) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r CustomerCreditLedgerNewEntryResponseVoidLedgerEntry) implementsCustomerCreditLedgerNewEntryResponse() {
}

type CustomerCreditLedgerNewEntryResponseVoidLedgerEntryCreditBlock struct {
	ID               string    `json:"id,required"`
	ExpiryDate       time.Time `json:"expiry_date,required,nullable" format:"date-time"`
	PerUnitCostBasis string    `json:"per_unit_cost_basis,required,nullable"`
	JSON             customerCreditLedgerNewEntryResponseVoidLedgerEntryCreditBlockJSON
}

// customerCreditLedgerNewEntryResponseVoidLedgerEntryCreditBlockJSON contains the
// JSON metadata for the struct
// [CustomerCreditLedgerNewEntryResponseVoidLedgerEntryCreditBlock]
type customerCreditLedgerNewEntryResponseVoidLedgerEntryCreditBlockJSON struct {
	ID               apijson.Field
	ExpiryDate       apijson.Field
	PerUnitCostBasis apijson.Field
	raw              string
	ExtraFields      map[string]apijson.Field
}

func (r *CustomerCreditLedgerNewEntryResponseVoidLedgerEntryCreditBlock) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type CustomerCreditLedgerNewEntryResponseVoidLedgerEntryCustomer struct {
	ID                 string `json:"id,required"`
	ExternalCustomerID string `json:"external_customer_id,required,nullable"`
	JSON               customerCreditLedgerNewEntryResponseVoidLedgerEntryCustomerJSON
}

// customerCreditLedgerNewEntryResponseVoidLedgerEntryCustomerJSON contains the
// JSON metadata for the struct
// [CustomerCreditLedgerNewEntryResponseVoidLedgerEntryCustomer]
type customerCreditLedgerNewEntryResponseVoidLedgerEntryCustomerJSON struct {
	ID                 apijson.Field
	ExternalCustomerID apijson.Field
	raw                string
	ExtraFields        map[string]apijson.Field
}

func (r *CustomerCreditLedgerNewEntryResponseVoidLedgerEntryCustomer) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type CustomerCreditLedgerNewEntryResponseVoidLedgerEntryEntryStatus string

const (
	CustomerCreditLedgerNewEntryResponseVoidLedgerEntryEntryStatusCommitted CustomerCreditLedgerNewEntryResponseVoidLedgerEntryEntryStatus = "committed"
	CustomerCreditLedgerNewEntryResponseVoidLedgerEntryEntryStatusPending   CustomerCreditLedgerNewEntryResponseVoidLedgerEntryEntryStatus = "pending"
)

type CustomerCreditLedgerNewEntryResponseVoidLedgerEntryEntryType string

const (
	CustomerCreditLedgerNewEntryResponseVoidLedgerEntryEntryTypeVoid CustomerCreditLedgerNewEntryResponseVoidLedgerEntryEntryType = "void"
)

type CustomerCreditLedgerNewEntryResponseVoidInitiatedLedgerEntry struct {
	ID                   string                                                                  `json:"id,required"`
	Amount               float64                                                                 `json:"amount,required"`
	CreatedAt            time.Time                                                               `json:"created_at,required" format:"date-time"`
	CreditBlock          CustomerCreditLedgerNewEntryResponseVoidInitiatedLedgerEntryCreditBlock `json:"credit_block,required"`
	Customer             CustomerCreditLedgerNewEntryResponseVoidInitiatedLedgerEntryCustomer    `json:"customer,required"`
	Description          string                                                                  `json:"description,required,nullable"`
	EndingBalance        float64                                                                 `json:"ending_balance,required"`
	EntryStatus          CustomerCreditLedgerNewEntryResponseVoidInitiatedLedgerEntryEntryStatus `json:"entry_status,required"`
	EntryType            CustomerCreditLedgerNewEntryResponseVoidInitiatedLedgerEntryEntryType   `json:"entry_type,required"`
	LedgerSequenceNumber int64                                                                   `json:"ledger_sequence_number,required"`
	Metadata             map[string]string                                                       `json:"metadata,required"`
	NewBlockExpiryDate   time.Time                                                               `json:"new_block_expiry_date,required" format:"date-time"`
	StartingBalance      float64                                                                 `json:"starting_balance,required"`
	VoidAmount           float64                                                                 `json:"void_amount,required"`
	VoidReason           string                                                                  `json:"void_reason,required,nullable"`
	JSON                 customerCreditLedgerNewEntryResponseVoidInitiatedLedgerEntryJSON
}

// customerCreditLedgerNewEntryResponseVoidInitiatedLedgerEntryJSON contains the
// JSON metadata for the struct
// [CustomerCreditLedgerNewEntryResponseVoidInitiatedLedgerEntry]
type customerCreditLedgerNewEntryResponseVoidInitiatedLedgerEntryJSON struct {
	ID                   apijson.Field
	Amount               apijson.Field
	CreatedAt            apijson.Field
	CreditBlock          apijson.Field
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

func (r *CustomerCreditLedgerNewEntryResponseVoidInitiatedLedgerEntry) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r CustomerCreditLedgerNewEntryResponseVoidInitiatedLedgerEntry) implementsCustomerCreditLedgerNewEntryResponse() {
}

type CustomerCreditLedgerNewEntryResponseVoidInitiatedLedgerEntryCreditBlock struct {
	ID               string    `json:"id,required"`
	ExpiryDate       time.Time `json:"expiry_date,required,nullable" format:"date-time"`
	PerUnitCostBasis string    `json:"per_unit_cost_basis,required,nullable"`
	JSON             customerCreditLedgerNewEntryResponseVoidInitiatedLedgerEntryCreditBlockJSON
}

// customerCreditLedgerNewEntryResponseVoidInitiatedLedgerEntryCreditBlockJSON
// contains the JSON metadata for the struct
// [CustomerCreditLedgerNewEntryResponseVoidInitiatedLedgerEntryCreditBlock]
type customerCreditLedgerNewEntryResponseVoidInitiatedLedgerEntryCreditBlockJSON struct {
	ID               apijson.Field
	ExpiryDate       apijson.Field
	PerUnitCostBasis apijson.Field
	raw              string
	ExtraFields      map[string]apijson.Field
}

func (r *CustomerCreditLedgerNewEntryResponseVoidInitiatedLedgerEntryCreditBlock) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type CustomerCreditLedgerNewEntryResponseVoidInitiatedLedgerEntryCustomer struct {
	ID                 string `json:"id,required"`
	ExternalCustomerID string `json:"external_customer_id,required,nullable"`
	JSON               customerCreditLedgerNewEntryResponseVoidInitiatedLedgerEntryCustomerJSON
}

// customerCreditLedgerNewEntryResponseVoidInitiatedLedgerEntryCustomerJSON
// contains the JSON metadata for the struct
// [CustomerCreditLedgerNewEntryResponseVoidInitiatedLedgerEntryCustomer]
type customerCreditLedgerNewEntryResponseVoidInitiatedLedgerEntryCustomerJSON struct {
	ID                 apijson.Field
	ExternalCustomerID apijson.Field
	raw                string
	ExtraFields        map[string]apijson.Field
}

func (r *CustomerCreditLedgerNewEntryResponseVoidInitiatedLedgerEntryCustomer) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type CustomerCreditLedgerNewEntryResponseVoidInitiatedLedgerEntryEntryStatus string

const (
	CustomerCreditLedgerNewEntryResponseVoidInitiatedLedgerEntryEntryStatusCommitted CustomerCreditLedgerNewEntryResponseVoidInitiatedLedgerEntryEntryStatus = "committed"
	CustomerCreditLedgerNewEntryResponseVoidInitiatedLedgerEntryEntryStatusPending   CustomerCreditLedgerNewEntryResponseVoidInitiatedLedgerEntryEntryStatus = "pending"
)

type CustomerCreditLedgerNewEntryResponseVoidInitiatedLedgerEntryEntryType string

const (
	CustomerCreditLedgerNewEntryResponseVoidInitiatedLedgerEntryEntryTypeVoidInitiated CustomerCreditLedgerNewEntryResponseVoidInitiatedLedgerEntryEntryType = "void_initiated"
)

type CustomerCreditLedgerNewEntryResponseAmendmentLedgerEntry struct {
	ID                   string                                                              `json:"id,required"`
	Amount               float64                                                             `json:"amount,required"`
	CreatedAt            time.Time                                                           `json:"created_at,required" format:"date-time"`
	CreditBlock          CustomerCreditLedgerNewEntryResponseAmendmentLedgerEntryCreditBlock `json:"credit_block,required"`
	Customer             CustomerCreditLedgerNewEntryResponseAmendmentLedgerEntryCustomer    `json:"customer,required"`
	Description          string                                                              `json:"description,required,nullable"`
	EndingBalance        float64                                                             `json:"ending_balance,required"`
	EntryStatus          CustomerCreditLedgerNewEntryResponseAmendmentLedgerEntryEntryStatus `json:"entry_status,required"`
	EntryType            CustomerCreditLedgerNewEntryResponseAmendmentLedgerEntryEntryType   `json:"entry_type,required"`
	LedgerSequenceNumber int64                                                               `json:"ledger_sequence_number,required"`
	Metadata             map[string]string                                                   `json:"metadata,required"`
	StartingBalance      float64                                                             `json:"starting_balance,required"`
	JSON                 customerCreditLedgerNewEntryResponseAmendmentLedgerEntryJSON
}

// customerCreditLedgerNewEntryResponseAmendmentLedgerEntryJSON contains the JSON
// metadata for the struct
// [CustomerCreditLedgerNewEntryResponseAmendmentLedgerEntry]
type customerCreditLedgerNewEntryResponseAmendmentLedgerEntryJSON struct {
	ID                   apijson.Field
	Amount               apijson.Field
	CreatedAt            apijson.Field
	CreditBlock          apijson.Field
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

func (r *CustomerCreditLedgerNewEntryResponseAmendmentLedgerEntry) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r CustomerCreditLedgerNewEntryResponseAmendmentLedgerEntry) implementsCustomerCreditLedgerNewEntryResponse() {
}

type CustomerCreditLedgerNewEntryResponseAmendmentLedgerEntryCreditBlock struct {
	ID               string    `json:"id,required"`
	ExpiryDate       time.Time `json:"expiry_date,required,nullable" format:"date-time"`
	PerUnitCostBasis string    `json:"per_unit_cost_basis,required,nullable"`
	JSON             customerCreditLedgerNewEntryResponseAmendmentLedgerEntryCreditBlockJSON
}

// customerCreditLedgerNewEntryResponseAmendmentLedgerEntryCreditBlockJSON contains
// the JSON metadata for the struct
// [CustomerCreditLedgerNewEntryResponseAmendmentLedgerEntryCreditBlock]
type customerCreditLedgerNewEntryResponseAmendmentLedgerEntryCreditBlockJSON struct {
	ID               apijson.Field
	ExpiryDate       apijson.Field
	PerUnitCostBasis apijson.Field
	raw              string
	ExtraFields      map[string]apijson.Field
}

func (r *CustomerCreditLedgerNewEntryResponseAmendmentLedgerEntryCreditBlock) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type CustomerCreditLedgerNewEntryResponseAmendmentLedgerEntryCustomer struct {
	ID                 string `json:"id,required"`
	ExternalCustomerID string `json:"external_customer_id,required,nullable"`
	JSON               customerCreditLedgerNewEntryResponseAmendmentLedgerEntryCustomerJSON
}

// customerCreditLedgerNewEntryResponseAmendmentLedgerEntryCustomerJSON contains
// the JSON metadata for the struct
// [CustomerCreditLedgerNewEntryResponseAmendmentLedgerEntryCustomer]
type customerCreditLedgerNewEntryResponseAmendmentLedgerEntryCustomerJSON struct {
	ID                 apijson.Field
	ExternalCustomerID apijson.Field
	raw                string
	ExtraFields        map[string]apijson.Field
}

func (r *CustomerCreditLedgerNewEntryResponseAmendmentLedgerEntryCustomer) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type CustomerCreditLedgerNewEntryResponseAmendmentLedgerEntryEntryStatus string

const (
	CustomerCreditLedgerNewEntryResponseAmendmentLedgerEntryEntryStatusCommitted CustomerCreditLedgerNewEntryResponseAmendmentLedgerEntryEntryStatus = "committed"
	CustomerCreditLedgerNewEntryResponseAmendmentLedgerEntryEntryStatusPending   CustomerCreditLedgerNewEntryResponseAmendmentLedgerEntryEntryStatus = "pending"
)

type CustomerCreditLedgerNewEntryResponseAmendmentLedgerEntryEntryType string

const (
	CustomerCreditLedgerNewEntryResponseAmendmentLedgerEntryEntryTypeAmendment CustomerCreditLedgerNewEntryResponseAmendmentLedgerEntryEntryType = "amendment"
)

// The [Credit Ledger Entry resource](/guides/product-catalog/prepurchase) models
// prepaid credits within Orb.
//
// Union satisfied by
// [CustomerCreditLedgerNewEntryByExteralIDResponseIncrementLedgerEntry],
// [CustomerCreditLedgerNewEntryByExteralIDResponseDecrementLedgerEntry],
// [CustomerCreditLedgerNewEntryByExteralIDResponseExpirationChangeLedgerEntry],
// [CustomerCreditLedgerNewEntryByExteralIDResponseCreditBlockExpiryLedgerEntry],
// [CustomerCreditLedgerNewEntryByExteralIDResponseVoidLedgerEntry],
// [CustomerCreditLedgerNewEntryByExteralIDResponseVoidInitiatedLedgerEntry] or
// [CustomerCreditLedgerNewEntryByExteralIDResponseAmendmentLedgerEntry].
type CustomerCreditLedgerNewEntryByExteralIDResponse interface {
	implementsCustomerCreditLedgerNewEntryByExteralIDResponse()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*CustomerCreditLedgerNewEntryByExteralIDResponse)(nil)).Elem(),
		"entry_type",
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			DiscriminatorValue: "\"increment\"",
			Type:               reflect.TypeOf(CustomerCreditLedgerNewEntryByExteralIDResponseIncrementLedgerEntry{}),
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			DiscriminatorValue: "\"decrement\"",
			Type:               reflect.TypeOf(CustomerCreditLedgerNewEntryByExteralIDResponseDecrementLedgerEntry{}),
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			DiscriminatorValue: "\"expiration_change\"",
			Type:               reflect.TypeOf(CustomerCreditLedgerNewEntryByExteralIDResponseExpirationChangeLedgerEntry{}),
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			DiscriminatorValue: "\"credit_block_expiry\"",
			Type:               reflect.TypeOf(CustomerCreditLedgerNewEntryByExteralIDResponseCreditBlockExpiryLedgerEntry{}),
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			DiscriminatorValue: "\"void\"",
			Type:               reflect.TypeOf(CustomerCreditLedgerNewEntryByExteralIDResponseVoidLedgerEntry{}),
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			DiscriminatorValue: "\"void_initiated\"",
			Type:               reflect.TypeOf(CustomerCreditLedgerNewEntryByExteralIDResponseVoidInitiatedLedgerEntry{}),
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			DiscriminatorValue: "\"amendment\"",
			Type:               reflect.TypeOf(CustomerCreditLedgerNewEntryByExteralIDResponseAmendmentLedgerEntry{}),
		},
	)
}

type CustomerCreditLedgerNewEntryByExteralIDResponseIncrementLedgerEntry struct {
	ID                   string                                                                         `json:"id,required"`
	Amount               float64                                                                        `json:"amount,required"`
	CreatedAt            time.Time                                                                      `json:"created_at,required" format:"date-time"`
	CreditBlock          CustomerCreditLedgerNewEntryByExteralIDResponseIncrementLedgerEntryCreditBlock `json:"credit_block,required"`
	Customer             CustomerCreditLedgerNewEntryByExteralIDResponseIncrementLedgerEntryCustomer    `json:"customer,required"`
	Description          string                                                                         `json:"description,required,nullable"`
	EndingBalance        float64                                                                        `json:"ending_balance,required"`
	EntryStatus          CustomerCreditLedgerNewEntryByExteralIDResponseIncrementLedgerEntryEntryStatus `json:"entry_status,required"`
	EntryType            CustomerCreditLedgerNewEntryByExteralIDResponseIncrementLedgerEntryEntryType   `json:"entry_type,required"`
	LedgerSequenceNumber int64                                                                          `json:"ledger_sequence_number,required"`
	Metadata             map[string]string                                                              `json:"metadata,required"`
	StartingBalance      float64                                                                        `json:"starting_balance,required"`
	JSON                 customerCreditLedgerNewEntryByExteralIDResponseIncrementLedgerEntryJSON
}

// customerCreditLedgerNewEntryByExteralIDResponseIncrementLedgerEntryJSON contains
// the JSON metadata for the struct
// [CustomerCreditLedgerNewEntryByExteralIDResponseIncrementLedgerEntry]
type customerCreditLedgerNewEntryByExteralIDResponseIncrementLedgerEntryJSON struct {
	ID                   apijson.Field
	Amount               apijson.Field
	CreatedAt            apijson.Field
	CreditBlock          apijson.Field
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

func (r *CustomerCreditLedgerNewEntryByExteralIDResponseIncrementLedgerEntry) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r CustomerCreditLedgerNewEntryByExteralIDResponseIncrementLedgerEntry) implementsCustomerCreditLedgerNewEntryByExteralIDResponse() {
}

type CustomerCreditLedgerNewEntryByExteralIDResponseIncrementLedgerEntryCreditBlock struct {
	ID               string    `json:"id,required"`
	ExpiryDate       time.Time `json:"expiry_date,required,nullable" format:"date-time"`
	PerUnitCostBasis string    `json:"per_unit_cost_basis,required,nullable"`
	JSON             customerCreditLedgerNewEntryByExteralIDResponseIncrementLedgerEntryCreditBlockJSON
}

// customerCreditLedgerNewEntryByExteralIDResponseIncrementLedgerEntryCreditBlockJSON
// contains the JSON metadata for the struct
// [CustomerCreditLedgerNewEntryByExteralIDResponseIncrementLedgerEntryCreditBlock]
type customerCreditLedgerNewEntryByExteralIDResponseIncrementLedgerEntryCreditBlockJSON struct {
	ID               apijson.Field
	ExpiryDate       apijson.Field
	PerUnitCostBasis apijson.Field
	raw              string
	ExtraFields      map[string]apijson.Field
}

func (r *CustomerCreditLedgerNewEntryByExteralIDResponseIncrementLedgerEntryCreditBlock) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type CustomerCreditLedgerNewEntryByExteralIDResponseIncrementLedgerEntryCustomer struct {
	ID                 string `json:"id,required"`
	ExternalCustomerID string `json:"external_customer_id,required,nullable"`
	JSON               customerCreditLedgerNewEntryByExteralIDResponseIncrementLedgerEntryCustomerJSON
}

// customerCreditLedgerNewEntryByExteralIDResponseIncrementLedgerEntryCustomerJSON
// contains the JSON metadata for the struct
// [CustomerCreditLedgerNewEntryByExteralIDResponseIncrementLedgerEntryCustomer]
type customerCreditLedgerNewEntryByExteralIDResponseIncrementLedgerEntryCustomerJSON struct {
	ID                 apijson.Field
	ExternalCustomerID apijson.Field
	raw                string
	ExtraFields        map[string]apijson.Field
}

func (r *CustomerCreditLedgerNewEntryByExteralIDResponseIncrementLedgerEntryCustomer) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type CustomerCreditLedgerNewEntryByExteralIDResponseIncrementLedgerEntryEntryStatus string

const (
	CustomerCreditLedgerNewEntryByExteralIDResponseIncrementLedgerEntryEntryStatusCommitted CustomerCreditLedgerNewEntryByExteralIDResponseIncrementLedgerEntryEntryStatus = "committed"
	CustomerCreditLedgerNewEntryByExteralIDResponseIncrementLedgerEntryEntryStatusPending   CustomerCreditLedgerNewEntryByExteralIDResponseIncrementLedgerEntryEntryStatus = "pending"
)

type CustomerCreditLedgerNewEntryByExteralIDResponseIncrementLedgerEntryEntryType string

const (
	CustomerCreditLedgerNewEntryByExteralIDResponseIncrementLedgerEntryEntryTypeIncrement CustomerCreditLedgerNewEntryByExteralIDResponseIncrementLedgerEntryEntryType = "increment"
)

type CustomerCreditLedgerNewEntryByExteralIDResponseDecrementLedgerEntry struct {
	ID                   string                                                                         `json:"id,required"`
	Amount               float64                                                                        `json:"amount,required"`
	CreatedAt            time.Time                                                                      `json:"created_at,required" format:"date-time"`
	CreditBlock          CustomerCreditLedgerNewEntryByExteralIDResponseDecrementLedgerEntryCreditBlock `json:"credit_block,required"`
	Customer             CustomerCreditLedgerNewEntryByExteralIDResponseDecrementLedgerEntryCustomer    `json:"customer,required"`
	Description          string                                                                         `json:"description,required,nullable"`
	EndingBalance        float64                                                                        `json:"ending_balance,required"`
	EntryStatus          CustomerCreditLedgerNewEntryByExteralIDResponseDecrementLedgerEntryEntryStatus `json:"entry_status,required"`
	EntryType            CustomerCreditLedgerNewEntryByExteralIDResponseDecrementLedgerEntryEntryType   `json:"entry_type,required"`
	LedgerSequenceNumber int64                                                                          `json:"ledger_sequence_number,required"`
	Metadata             map[string]string                                                              `json:"metadata,required"`
	StartingBalance      float64                                                                        `json:"starting_balance,required"`
	EventID              string                                                                         `json:"event_id,nullable"`
	InvoiceID            string                                                                         `json:"invoice_id,nullable"`
	PriceID              string                                                                         `json:"price_id,nullable"`
	JSON                 customerCreditLedgerNewEntryByExteralIDResponseDecrementLedgerEntryJSON
}

// customerCreditLedgerNewEntryByExteralIDResponseDecrementLedgerEntryJSON contains
// the JSON metadata for the struct
// [CustomerCreditLedgerNewEntryByExteralIDResponseDecrementLedgerEntry]
type customerCreditLedgerNewEntryByExteralIDResponseDecrementLedgerEntryJSON struct {
	ID                   apijson.Field
	Amount               apijson.Field
	CreatedAt            apijson.Field
	CreditBlock          apijson.Field
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

func (r *CustomerCreditLedgerNewEntryByExteralIDResponseDecrementLedgerEntry) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r CustomerCreditLedgerNewEntryByExteralIDResponseDecrementLedgerEntry) implementsCustomerCreditLedgerNewEntryByExteralIDResponse() {
}

type CustomerCreditLedgerNewEntryByExteralIDResponseDecrementLedgerEntryCreditBlock struct {
	ID               string    `json:"id,required"`
	ExpiryDate       time.Time `json:"expiry_date,required,nullable" format:"date-time"`
	PerUnitCostBasis string    `json:"per_unit_cost_basis,required,nullable"`
	JSON             customerCreditLedgerNewEntryByExteralIDResponseDecrementLedgerEntryCreditBlockJSON
}

// customerCreditLedgerNewEntryByExteralIDResponseDecrementLedgerEntryCreditBlockJSON
// contains the JSON metadata for the struct
// [CustomerCreditLedgerNewEntryByExteralIDResponseDecrementLedgerEntryCreditBlock]
type customerCreditLedgerNewEntryByExteralIDResponseDecrementLedgerEntryCreditBlockJSON struct {
	ID               apijson.Field
	ExpiryDate       apijson.Field
	PerUnitCostBasis apijson.Field
	raw              string
	ExtraFields      map[string]apijson.Field
}

func (r *CustomerCreditLedgerNewEntryByExteralIDResponseDecrementLedgerEntryCreditBlock) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type CustomerCreditLedgerNewEntryByExteralIDResponseDecrementLedgerEntryCustomer struct {
	ID                 string `json:"id,required"`
	ExternalCustomerID string `json:"external_customer_id,required,nullable"`
	JSON               customerCreditLedgerNewEntryByExteralIDResponseDecrementLedgerEntryCustomerJSON
}

// customerCreditLedgerNewEntryByExteralIDResponseDecrementLedgerEntryCustomerJSON
// contains the JSON metadata for the struct
// [CustomerCreditLedgerNewEntryByExteralIDResponseDecrementLedgerEntryCustomer]
type customerCreditLedgerNewEntryByExteralIDResponseDecrementLedgerEntryCustomerJSON struct {
	ID                 apijson.Field
	ExternalCustomerID apijson.Field
	raw                string
	ExtraFields        map[string]apijson.Field
}

func (r *CustomerCreditLedgerNewEntryByExteralIDResponseDecrementLedgerEntryCustomer) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type CustomerCreditLedgerNewEntryByExteralIDResponseDecrementLedgerEntryEntryStatus string

const (
	CustomerCreditLedgerNewEntryByExteralIDResponseDecrementLedgerEntryEntryStatusCommitted CustomerCreditLedgerNewEntryByExteralIDResponseDecrementLedgerEntryEntryStatus = "committed"
	CustomerCreditLedgerNewEntryByExteralIDResponseDecrementLedgerEntryEntryStatusPending   CustomerCreditLedgerNewEntryByExteralIDResponseDecrementLedgerEntryEntryStatus = "pending"
)

type CustomerCreditLedgerNewEntryByExteralIDResponseDecrementLedgerEntryEntryType string

const (
	CustomerCreditLedgerNewEntryByExteralIDResponseDecrementLedgerEntryEntryTypeDecrement CustomerCreditLedgerNewEntryByExteralIDResponseDecrementLedgerEntryEntryType = "decrement"
)

type CustomerCreditLedgerNewEntryByExteralIDResponseExpirationChangeLedgerEntry struct {
	ID                   string                                                                                `json:"id,required"`
	Amount               float64                                                                               `json:"amount,required"`
	CreatedAt            time.Time                                                                             `json:"created_at,required" format:"date-time"`
	CreditBlock          CustomerCreditLedgerNewEntryByExteralIDResponseExpirationChangeLedgerEntryCreditBlock `json:"credit_block,required"`
	Customer             CustomerCreditLedgerNewEntryByExteralIDResponseExpirationChangeLedgerEntryCustomer    `json:"customer,required"`
	Description          string                                                                                `json:"description,required,nullable"`
	EndingBalance        float64                                                                               `json:"ending_balance,required"`
	EntryStatus          CustomerCreditLedgerNewEntryByExteralIDResponseExpirationChangeLedgerEntryEntryStatus `json:"entry_status,required"`
	EntryType            CustomerCreditLedgerNewEntryByExteralIDResponseExpirationChangeLedgerEntryEntryType   `json:"entry_type,required"`
	LedgerSequenceNumber int64                                                                                 `json:"ledger_sequence_number,required"`
	Metadata             map[string]string                                                                     `json:"metadata,required"`
	NewBlockExpiryDate   time.Time                                                                             `json:"new_block_expiry_date,required" format:"date-time"`
	StartingBalance      float64                                                                               `json:"starting_balance,required"`
	JSON                 customerCreditLedgerNewEntryByExteralIDResponseExpirationChangeLedgerEntryJSON
}

// customerCreditLedgerNewEntryByExteralIDResponseExpirationChangeLedgerEntryJSON
// contains the JSON metadata for the struct
// [CustomerCreditLedgerNewEntryByExteralIDResponseExpirationChangeLedgerEntry]
type customerCreditLedgerNewEntryByExteralIDResponseExpirationChangeLedgerEntryJSON struct {
	ID                   apijson.Field
	Amount               apijson.Field
	CreatedAt            apijson.Field
	CreditBlock          apijson.Field
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

func (r *CustomerCreditLedgerNewEntryByExteralIDResponseExpirationChangeLedgerEntry) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r CustomerCreditLedgerNewEntryByExteralIDResponseExpirationChangeLedgerEntry) implementsCustomerCreditLedgerNewEntryByExteralIDResponse() {
}

type CustomerCreditLedgerNewEntryByExteralIDResponseExpirationChangeLedgerEntryCreditBlock struct {
	ID               string    `json:"id,required"`
	ExpiryDate       time.Time `json:"expiry_date,required,nullable" format:"date-time"`
	PerUnitCostBasis string    `json:"per_unit_cost_basis,required,nullable"`
	JSON             customerCreditLedgerNewEntryByExteralIDResponseExpirationChangeLedgerEntryCreditBlockJSON
}

// customerCreditLedgerNewEntryByExteralIDResponseExpirationChangeLedgerEntryCreditBlockJSON
// contains the JSON metadata for the struct
// [CustomerCreditLedgerNewEntryByExteralIDResponseExpirationChangeLedgerEntryCreditBlock]
type customerCreditLedgerNewEntryByExteralIDResponseExpirationChangeLedgerEntryCreditBlockJSON struct {
	ID               apijson.Field
	ExpiryDate       apijson.Field
	PerUnitCostBasis apijson.Field
	raw              string
	ExtraFields      map[string]apijson.Field
}

func (r *CustomerCreditLedgerNewEntryByExteralIDResponseExpirationChangeLedgerEntryCreditBlock) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type CustomerCreditLedgerNewEntryByExteralIDResponseExpirationChangeLedgerEntryCustomer struct {
	ID                 string `json:"id,required"`
	ExternalCustomerID string `json:"external_customer_id,required,nullable"`
	JSON               customerCreditLedgerNewEntryByExteralIDResponseExpirationChangeLedgerEntryCustomerJSON
}

// customerCreditLedgerNewEntryByExteralIDResponseExpirationChangeLedgerEntryCustomerJSON
// contains the JSON metadata for the struct
// [CustomerCreditLedgerNewEntryByExteralIDResponseExpirationChangeLedgerEntryCustomer]
type customerCreditLedgerNewEntryByExteralIDResponseExpirationChangeLedgerEntryCustomerJSON struct {
	ID                 apijson.Field
	ExternalCustomerID apijson.Field
	raw                string
	ExtraFields        map[string]apijson.Field
}

func (r *CustomerCreditLedgerNewEntryByExteralIDResponseExpirationChangeLedgerEntryCustomer) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type CustomerCreditLedgerNewEntryByExteralIDResponseExpirationChangeLedgerEntryEntryStatus string

const (
	CustomerCreditLedgerNewEntryByExteralIDResponseExpirationChangeLedgerEntryEntryStatusCommitted CustomerCreditLedgerNewEntryByExteralIDResponseExpirationChangeLedgerEntryEntryStatus = "committed"
	CustomerCreditLedgerNewEntryByExteralIDResponseExpirationChangeLedgerEntryEntryStatusPending   CustomerCreditLedgerNewEntryByExteralIDResponseExpirationChangeLedgerEntryEntryStatus = "pending"
)

type CustomerCreditLedgerNewEntryByExteralIDResponseExpirationChangeLedgerEntryEntryType string

const (
	CustomerCreditLedgerNewEntryByExteralIDResponseExpirationChangeLedgerEntryEntryTypeExpirationChange CustomerCreditLedgerNewEntryByExteralIDResponseExpirationChangeLedgerEntryEntryType = "expiration_change"
)

type CustomerCreditLedgerNewEntryByExteralIDResponseCreditBlockExpiryLedgerEntry struct {
	ID                   string                                                                                 `json:"id,required"`
	Amount               float64                                                                                `json:"amount,required"`
	CreatedAt            time.Time                                                                              `json:"created_at,required" format:"date-time"`
	CreditBlock          CustomerCreditLedgerNewEntryByExteralIDResponseCreditBlockExpiryLedgerEntryCreditBlock `json:"credit_block,required"`
	Customer             CustomerCreditLedgerNewEntryByExteralIDResponseCreditBlockExpiryLedgerEntryCustomer    `json:"customer,required"`
	Description          string                                                                                 `json:"description,required,nullable"`
	EndingBalance        float64                                                                                `json:"ending_balance,required"`
	EntryStatus          CustomerCreditLedgerNewEntryByExteralIDResponseCreditBlockExpiryLedgerEntryEntryStatus `json:"entry_status,required"`
	EntryType            CustomerCreditLedgerNewEntryByExteralIDResponseCreditBlockExpiryLedgerEntryEntryType   `json:"entry_type,required"`
	LedgerSequenceNumber int64                                                                                  `json:"ledger_sequence_number,required"`
	Metadata             map[string]string                                                                      `json:"metadata,required"`
	StartingBalance      float64                                                                                `json:"starting_balance,required"`
	JSON                 customerCreditLedgerNewEntryByExteralIDResponseCreditBlockExpiryLedgerEntryJSON
}

// customerCreditLedgerNewEntryByExteralIDResponseCreditBlockExpiryLedgerEntryJSON
// contains the JSON metadata for the struct
// [CustomerCreditLedgerNewEntryByExteralIDResponseCreditBlockExpiryLedgerEntry]
type customerCreditLedgerNewEntryByExteralIDResponseCreditBlockExpiryLedgerEntryJSON struct {
	ID                   apijson.Field
	Amount               apijson.Field
	CreatedAt            apijson.Field
	CreditBlock          apijson.Field
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

func (r *CustomerCreditLedgerNewEntryByExteralIDResponseCreditBlockExpiryLedgerEntry) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r CustomerCreditLedgerNewEntryByExteralIDResponseCreditBlockExpiryLedgerEntry) implementsCustomerCreditLedgerNewEntryByExteralIDResponse() {
}

type CustomerCreditLedgerNewEntryByExteralIDResponseCreditBlockExpiryLedgerEntryCreditBlock struct {
	ID               string    `json:"id,required"`
	ExpiryDate       time.Time `json:"expiry_date,required,nullable" format:"date-time"`
	PerUnitCostBasis string    `json:"per_unit_cost_basis,required,nullable"`
	JSON             customerCreditLedgerNewEntryByExteralIDResponseCreditBlockExpiryLedgerEntryCreditBlockJSON
}

// customerCreditLedgerNewEntryByExteralIDResponseCreditBlockExpiryLedgerEntryCreditBlockJSON
// contains the JSON metadata for the struct
// [CustomerCreditLedgerNewEntryByExteralIDResponseCreditBlockExpiryLedgerEntryCreditBlock]
type customerCreditLedgerNewEntryByExteralIDResponseCreditBlockExpiryLedgerEntryCreditBlockJSON struct {
	ID               apijson.Field
	ExpiryDate       apijson.Field
	PerUnitCostBasis apijson.Field
	raw              string
	ExtraFields      map[string]apijson.Field
}

func (r *CustomerCreditLedgerNewEntryByExteralIDResponseCreditBlockExpiryLedgerEntryCreditBlock) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type CustomerCreditLedgerNewEntryByExteralIDResponseCreditBlockExpiryLedgerEntryCustomer struct {
	ID                 string `json:"id,required"`
	ExternalCustomerID string `json:"external_customer_id,required,nullable"`
	JSON               customerCreditLedgerNewEntryByExteralIDResponseCreditBlockExpiryLedgerEntryCustomerJSON
}

// customerCreditLedgerNewEntryByExteralIDResponseCreditBlockExpiryLedgerEntryCustomerJSON
// contains the JSON metadata for the struct
// [CustomerCreditLedgerNewEntryByExteralIDResponseCreditBlockExpiryLedgerEntryCustomer]
type customerCreditLedgerNewEntryByExteralIDResponseCreditBlockExpiryLedgerEntryCustomerJSON struct {
	ID                 apijson.Field
	ExternalCustomerID apijson.Field
	raw                string
	ExtraFields        map[string]apijson.Field
}

func (r *CustomerCreditLedgerNewEntryByExteralIDResponseCreditBlockExpiryLedgerEntryCustomer) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type CustomerCreditLedgerNewEntryByExteralIDResponseCreditBlockExpiryLedgerEntryEntryStatus string

const (
	CustomerCreditLedgerNewEntryByExteralIDResponseCreditBlockExpiryLedgerEntryEntryStatusCommitted CustomerCreditLedgerNewEntryByExteralIDResponseCreditBlockExpiryLedgerEntryEntryStatus = "committed"
	CustomerCreditLedgerNewEntryByExteralIDResponseCreditBlockExpiryLedgerEntryEntryStatusPending   CustomerCreditLedgerNewEntryByExteralIDResponseCreditBlockExpiryLedgerEntryEntryStatus = "pending"
)

type CustomerCreditLedgerNewEntryByExteralIDResponseCreditBlockExpiryLedgerEntryEntryType string

const (
	CustomerCreditLedgerNewEntryByExteralIDResponseCreditBlockExpiryLedgerEntryEntryTypeCreditBlockExpiry CustomerCreditLedgerNewEntryByExteralIDResponseCreditBlockExpiryLedgerEntryEntryType = "credit_block_expiry"
)

type CustomerCreditLedgerNewEntryByExteralIDResponseVoidLedgerEntry struct {
	ID                   string                                                                    `json:"id,required"`
	Amount               float64                                                                   `json:"amount,required"`
	CreatedAt            time.Time                                                                 `json:"created_at,required" format:"date-time"`
	CreditBlock          CustomerCreditLedgerNewEntryByExteralIDResponseVoidLedgerEntryCreditBlock `json:"credit_block,required"`
	Customer             CustomerCreditLedgerNewEntryByExteralIDResponseVoidLedgerEntryCustomer    `json:"customer,required"`
	Description          string                                                                    `json:"description,required,nullable"`
	EndingBalance        float64                                                                   `json:"ending_balance,required"`
	EntryStatus          CustomerCreditLedgerNewEntryByExteralIDResponseVoidLedgerEntryEntryStatus `json:"entry_status,required"`
	EntryType            CustomerCreditLedgerNewEntryByExteralIDResponseVoidLedgerEntryEntryType   `json:"entry_type,required"`
	LedgerSequenceNumber int64                                                                     `json:"ledger_sequence_number,required"`
	Metadata             map[string]string                                                         `json:"metadata,required"`
	StartingBalance      float64                                                                   `json:"starting_balance,required"`
	VoidAmount           float64                                                                   `json:"void_amount,required"`
	VoidReason           string                                                                    `json:"void_reason,required,nullable"`
	JSON                 customerCreditLedgerNewEntryByExteralIDResponseVoidLedgerEntryJSON
}

// customerCreditLedgerNewEntryByExteralIDResponseVoidLedgerEntryJSON contains the
// JSON metadata for the struct
// [CustomerCreditLedgerNewEntryByExteralIDResponseVoidLedgerEntry]
type customerCreditLedgerNewEntryByExteralIDResponseVoidLedgerEntryJSON struct {
	ID                   apijson.Field
	Amount               apijson.Field
	CreatedAt            apijson.Field
	CreditBlock          apijson.Field
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

func (r *CustomerCreditLedgerNewEntryByExteralIDResponseVoidLedgerEntry) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r CustomerCreditLedgerNewEntryByExteralIDResponseVoidLedgerEntry) implementsCustomerCreditLedgerNewEntryByExteralIDResponse() {
}

type CustomerCreditLedgerNewEntryByExteralIDResponseVoidLedgerEntryCreditBlock struct {
	ID               string    `json:"id,required"`
	ExpiryDate       time.Time `json:"expiry_date,required,nullable" format:"date-time"`
	PerUnitCostBasis string    `json:"per_unit_cost_basis,required,nullable"`
	JSON             customerCreditLedgerNewEntryByExteralIDResponseVoidLedgerEntryCreditBlockJSON
}

// customerCreditLedgerNewEntryByExteralIDResponseVoidLedgerEntryCreditBlockJSON
// contains the JSON metadata for the struct
// [CustomerCreditLedgerNewEntryByExteralIDResponseVoidLedgerEntryCreditBlock]
type customerCreditLedgerNewEntryByExteralIDResponseVoidLedgerEntryCreditBlockJSON struct {
	ID               apijson.Field
	ExpiryDate       apijson.Field
	PerUnitCostBasis apijson.Field
	raw              string
	ExtraFields      map[string]apijson.Field
}

func (r *CustomerCreditLedgerNewEntryByExteralIDResponseVoidLedgerEntryCreditBlock) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type CustomerCreditLedgerNewEntryByExteralIDResponseVoidLedgerEntryCustomer struct {
	ID                 string `json:"id,required"`
	ExternalCustomerID string `json:"external_customer_id,required,nullable"`
	JSON               customerCreditLedgerNewEntryByExteralIDResponseVoidLedgerEntryCustomerJSON
}

// customerCreditLedgerNewEntryByExteralIDResponseVoidLedgerEntryCustomerJSON
// contains the JSON metadata for the struct
// [CustomerCreditLedgerNewEntryByExteralIDResponseVoidLedgerEntryCustomer]
type customerCreditLedgerNewEntryByExteralIDResponseVoidLedgerEntryCustomerJSON struct {
	ID                 apijson.Field
	ExternalCustomerID apijson.Field
	raw                string
	ExtraFields        map[string]apijson.Field
}

func (r *CustomerCreditLedgerNewEntryByExteralIDResponseVoidLedgerEntryCustomer) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type CustomerCreditLedgerNewEntryByExteralIDResponseVoidLedgerEntryEntryStatus string

const (
	CustomerCreditLedgerNewEntryByExteralIDResponseVoidLedgerEntryEntryStatusCommitted CustomerCreditLedgerNewEntryByExteralIDResponseVoidLedgerEntryEntryStatus = "committed"
	CustomerCreditLedgerNewEntryByExteralIDResponseVoidLedgerEntryEntryStatusPending   CustomerCreditLedgerNewEntryByExteralIDResponseVoidLedgerEntryEntryStatus = "pending"
)

type CustomerCreditLedgerNewEntryByExteralIDResponseVoidLedgerEntryEntryType string

const (
	CustomerCreditLedgerNewEntryByExteralIDResponseVoidLedgerEntryEntryTypeVoid CustomerCreditLedgerNewEntryByExteralIDResponseVoidLedgerEntryEntryType = "void"
)

type CustomerCreditLedgerNewEntryByExteralIDResponseVoidInitiatedLedgerEntry struct {
	ID                   string                                                                             `json:"id,required"`
	Amount               float64                                                                            `json:"amount,required"`
	CreatedAt            time.Time                                                                          `json:"created_at,required" format:"date-time"`
	CreditBlock          CustomerCreditLedgerNewEntryByExteralIDResponseVoidInitiatedLedgerEntryCreditBlock `json:"credit_block,required"`
	Customer             CustomerCreditLedgerNewEntryByExteralIDResponseVoidInitiatedLedgerEntryCustomer    `json:"customer,required"`
	Description          string                                                                             `json:"description,required,nullable"`
	EndingBalance        float64                                                                            `json:"ending_balance,required"`
	EntryStatus          CustomerCreditLedgerNewEntryByExteralIDResponseVoidInitiatedLedgerEntryEntryStatus `json:"entry_status,required"`
	EntryType            CustomerCreditLedgerNewEntryByExteralIDResponseVoidInitiatedLedgerEntryEntryType   `json:"entry_type,required"`
	LedgerSequenceNumber int64                                                                              `json:"ledger_sequence_number,required"`
	Metadata             map[string]string                                                                  `json:"metadata,required"`
	NewBlockExpiryDate   time.Time                                                                          `json:"new_block_expiry_date,required" format:"date-time"`
	StartingBalance      float64                                                                            `json:"starting_balance,required"`
	VoidAmount           float64                                                                            `json:"void_amount,required"`
	VoidReason           string                                                                             `json:"void_reason,required,nullable"`
	JSON                 customerCreditLedgerNewEntryByExteralIDResponseVoidInitiatedLedgerEntryJSON
}

// customerCreditLedgerNewEntryByExteralIDResponseVoidInitiatedLedgerEntryJSON
// contains the JSON metadata for the struct
// [CustomerCreditLedgerNewEntryByExteralIDResponseVoidInitiatedLedgerEntry]
type customerCreditLedgerNewEntryByExteralIDResponseVoidInitiatedLedgerEntryJSON struct {
	ID                   apijson.Field
	Amount               apijson.Field
	CreatedAt            apijson.Field
	CreditBlock          apijson.Field
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

func (r *CustomerCreditLedgerNewEntryByExteralIDResponseVoidInitiatedLedgerEntry) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r CustomerCreditLedgerNewEntryByExteralIDResponseVoidInitiatedLedgerEntry) implementsCustomerCreditLedgerNewEntryByExteralIDResponse() {
}

type CustomerCreditLedgerNewEntryByExteralIDResponseVoidInitiatedLedgerEntryCreditBlock struct {
	ID               string    `json:"id,required"`
	ExpiryDate       time.Time `json:"expiry_date,required,nullable" format:"date-time"`
	PerUnitCostBasis string    `json:"per_unit_cost_basis,required,nullable"`
	JSON             customerCreditLedgerNewEntryByExteralIDResponseVoidInitiatedLedgerEntryCreditBlockJSON
}

// customerCreditLedgerNewEntryByExteralIDResponseVoidInitiatedLedgerEntryCreditBlockJSON
// contains the JSON metadata for the struct
// [CustomerCreditLedgerNewEntryByExteralIDResponseVoidInitiatedLedgerEntryCreditBlock]
type customerCreditLedgerNewEntryByExteralIDResponseVoidInitiatedLedgerEntryCreditBlockJSON struct {
	ID               apijson.Field
	ExpiryDate       apijson.Field
	PerUnitCostBasis apijson.Field
	raw              string
	ExtraFields      map[string]apijson.Field
}

func (r *CustomerCreditLedgerNewEntryByExteralIDResponseVoidInitiatedLedgerEntryCreditBlock) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type CustomerCreditLedgerNewEntryByExteralIDResponseVoidInitiatedLedgerEntryCustomer struct {
	ID                 string `json:"id,required"`
	ExternalCustomerID string `json:"external_customer_id,required,nullable"`
	JSON               customerCreditLedgerNewEntryByExteralIDResponseVoidInitiatedLedgerEntryCustomerJSON
}

// customerCreditLedgerNewEntryByExteralIDResponseVoidInitiatedLedgerEntryCustomerJSON
// contains the JSON metadata for the struct
// [CustomerCreditLedgerNewEntryByExteralIDResponseVoidInitiatedLedgerEntryCustomer]
type customerCreditLedgerNewEntryByExteralIDResponseVoidInitiatedLedgerEntryCustomerJSON struct {
	ID                 apijson.Field
	ExternalCustomerID apijson.Field
	raw                string
	ExtraFields        map[string]apijson.Field
}

func (r *CustomerCreditLedgerNewEntryByExteralIDResponseVoidInitiatedLedgerEntryCustomer) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type CustomerCreditLedgerNewEntryByExteralIDResponseVoidInitiatedLedgerEntryEntryStatus string

const (
	CustomerCreditLedgerNewEntryByExteralIDResponseVoidInitiatedLedgerEntryEntryStatusCommitted CustomerCreditLedgerNewEntryByExteralIDResponseVoidInitiatedLedgerEntryEntryStatus = "committed"
	CustomerCreditLedgerNewEntryByExteralIDResponseVoidInitiatedLedgerEntryEntryStatusPending   CustomerCreditLedgerNewEntryByExteralIDResponseVoidInitiatedLedgerEntryEntryStatus = "pending"
)

type CustomerCreditLedgerNewEntryByExteralIDResponseVoidInitiatedLedgerEntryEntryType string

const (
	CustomerCreditLedgerNewEntryByExteralIDResponseVoidInitiatedLedgerEntryEntryTypeVoidInitiated CustomerCreditLedgerNewEntryByExteralIDResponseVoidInitiatedLedgerEntryEntryType = "void_initiated"
)

type CustomerCreditLedgerNewEntryByExteralIDResponseAmendmentLedgerEntry struct {
	ID                   string                                                                         `json:"id,required"`
	Amount               float64                                                                        `json:"amount,required"`
	CreatedAt            time.Time                                                                      `json:"created_at,required" format:"date-time"`
	CreditBlock          CustomerCreditLedgerNewEntryByExteralIDResponseAmendmentLedgerEntryCreditBlock `json:"credit_block,required"`
	Customer             CustomerCreditLedgerNewEntryByExteralIDResponseAmendmentLedgerEntryCustomer    `json:"customer,required"`
	Description          string                                                                         `json:"description,required,nullable"`
	EndingBalance        float64                                                                        `json:"ending_balance,required"`
	EntryStatus          CustomerCreditLedgerNewEntryByExteralIDResponseAmendmentLedgerEntryEntryStatus `json:"entry_status,required"`
	EntryType            CustomerCreditLedgerNewEntryByExteralIDResponseAmendmentLedgerEntryEntryType   `json:"entry_type,required"`
	LedgerSequenceNumber int64                                                                          `json:"ledger_sequence_number,required"`
	Metadata             map[string]string                                                              `json:"metadata,required"`
	StartingBalance      float64                                                                        `json:"starting_balance,required"`
	JSON                 customerCreditLedgerNewEntryByExteralIDResponseAmendmentLedgerEntryJSON
}

// customerCreditLedgerNewEntryByExteralIDResponseAmendmentLedgerEntryJSON contains
// the JSON metadata for the struct
// [CustomerCreditLedgerNewEntryByExteralIDResponseAmendmentLedgerEntry]
type customerCreditLedgerNewEntryByExteralIDResponseAmendmentLedgerEntryJSON struct {
	ID                   apijson.Field
	Amount               apijson.Field
	CreatedAt            apijson.Field
	CreditBlock          apijson.Field
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

func (r *CustomerCreditLedgerNewEntryByExteralIDResponseAmendmentLedgerEntry) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r CustomerCreditLedgerNewEntryByExteralIDResponseAmendmentLedgerEntry) implementsCustomerCreditLedgerNewEntryByExteralIDResponse() {
}

type CustomerCreditLedgerNewEntryByExteralIDResponseAmendmentLedgerEntryCreditBlock struct {
	ID               string    `json:"id,required"`
	ExpiryDate       time.Time `json:"expiry_date,required,nullable" format:"date-time"`
	PerUnitCostBasis string    `json:"per_unit_cost_basis,required,nullable"`
	JSON             customerCreditLedgerNewEntryByExteralIDResponseAmendmentLedgerEntryCreditBlockJSON
}

// customerCreditLedgerNewEntryByExteralIDResponseAmendmentLedgerEntryCreditBlockJSON
// contains the JSON metadata for the struct
// [CustomerCreditLedgerNewEntryByExteralIDResponseAmendmentLedgerEntryCreditBlock]
type customerCreditLedgerNewEntryByExteralIDResponseAmendmentLedgerEntryCreditBlockJSON struct {
	ID               apijson.Field
	ExpiryDate       apijson.Field
	PerUnitCostBasis apijson.Field
	raw              string
	ExtraFields      map[string]apijson.Field
}

func (r *CustomerCreditLedgerNewEntryByExteralIDResponseAmendmentLedgerEntryCreditBlock) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type CustomerCreditLedgerNewEntryByExteralIDResponseAmendmentLedgerEntryCustomer struct {
	ID                 string `json:"id,required"`
	ExternalCustomerID string `json:"external_customer_id,required,nullable"`
	JSON               customerCreditLedgerNewEntryByExteralIDResponseAmendmentLedgerEntryCustomerJSON
}

// customerCreditLedgerNewEntryByExteralIDResponseAmendmentLedgerEntryCustomerJSON
// contains the JSON metadata for the struct
// [CustomerCreditLedgerNewEntryByExteralIDResponseAmendmentLedgerEntryCustomer]
type customerCreditLedgerNewEntryByExteralIDResponseAmendmentLedgerEntryCustomerJSON struct {
	ID                 apijson.Field
	ExternalCustomerID apijson.Field
	raw                string
	ExtraFields        map[string]apijson.Field
}

func (r *CustomerCreditLedgerNewEntryByExteralIDResponseAmendmentLedgerEntryCustomer) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type CustomerCreditLedgerNewEntryByExteralIDResponseAmendmentLedgerEntryEntryStatus string

const (
	CustomerCreditLedgerNewEntryByExteralIDResponseAmendmentLedgerEntryEntryStatusCommitted CustomerCreditLedgerNewEntryByExteralIDResponseAmendmentLedgerEntryEntryStatus = "committed"
	CustomerCreditLedgerNewEntryByExteralIDResponseAmendmentLedgerEntryEntryStatusPending   CustomerCreditLedgerNewEntryByExteralIDResponseAmendmentLedgerEntryEntryStatus = "pending"
)

type CustomerCreditLedgerNewEntryByExteralIDResponseAmendmentLedgerEntryEntryType string

const (
	CustomerCreditLedgerNewEntryByExteralIDResponseAmendmentLedgerEntryEntryTypeAmendment CustomerCreditLedgerNewEntryByExteralIDResponseAmendmentLedgerEntryEntryType = "amendment"
)

// The [Credit Ledger Entry resource](/guides/product-catalog/prepurchase) models
// prepaid credits within Orb.
//
// Union satisfied by
// [CustomerCreditLedgerListByExternalIDResponseIncrementLedgerEntry],
// [CustomerCreditLedgerListByExternalIDResponseDecrementLedgerEntry],
// [CustomerCreditLedgerListByExternalIDResponseExpirationChangeLedgerEntry],
// [CustomerCreditLedgerListByExternalIDResponseCreditBlockExpiryLedgerEntry],
// [CustomerCreditLedgerListByExternalIDResponseVoidLedgerEntry],
// [CustomerCreditLedgerListByExternalIDResponseVoidInitiatedLedgerEntry] or
// [CustomerCreditLedgerListByExternalIDResponseAmendmentLedgerEntry].
type CustomerCreditLedgerListByExternalIDResponse interface {
	implementsCustomerCreditLedgerListByExternalIDResponse()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*CustomerCreditLedgerListByExternalIDResponse)(nil)).Elem(),
		"entry_type",
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			DiscriminatorValue: "\"increment\"",
			Type:               reflect.TypeOf(CustomerCreditLedgerListByExternalIDResponseIncrementLedgerEntry{}),
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			DiscriminatorValue: "\"decrement\"",
			Type:               reflect.TypeOf(CustomerCreditLedgerListByExternalIDResponseDecrementLedgerEntry{}),
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			DiscriminatorValue: "\"expiration_change\"",
			Type:               reflect.TypeOf(CustomerCreditLedgerListByExternalIDResponseExpirationChangeLedgerEntry{}),
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			DiscriminatorValue: "\"credit_block_expiry\"",
			Type:               reflect.TypeOf(CustomerCreditLedgerListByExternalIDResponseCreditBlockExpiryLedgerEntry{}),
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			DiscriminatorValue: "\"void\"",
			Type:               reflect.TypeOf(CustomerCreditLedgerListByExternalIDResponseVoidLedgerEntry{}),
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			DiscriminatorValue: "\"void_initiated\"",
			Type:               reflect.TypeOf(CustomerCreditLedgerListByExternalIDResponseVoidInitiatedLedgerEntry{}),
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			DiscriminatorValue: "\"amendment\"",
			Type:               reflect.TypeOf(CustomerCreditLedgerListByExternalIDResponseAmendmentLedgerEntry{}),
		},
	)
}

type CustomerCreditLedgerListByExternalIDResponseIncrementLedgerEntry struct {
	ID                   string                                                                      `json:"id,required"`
	Amount               float64                                                                     `json:"amount,required"`
	CreatedAt            time.Time                                                                   `json:"created_at,required" format:"date-time"`
	CreditBlock          CustomerCreditLedgerListByExternalIDResponseIncrementLedgerEntryCreditBlock `json:"credit_block,required"`
	Customer             CustomerCreditLedgerListByExternalIDResponseIncrementLedgerEntryCustomer    `json:"customer,required"`
	Description          string                                                                      `json:"description,required,nullable"`
	EndingBalance        float64                                                                     `json:"ending_balance,required"`
	EntryStatus          CustomerCreditLedgerListByExternalIDResponseIncrementLedgerEntryEntryStatus `json:"entry_status,required"`
	EntryType            CustomerCreditLedgerListByExternalIDResponseIncrementLedgerEntryEntryType   `json:"entry_type,required"`
	LedgerSequenceNumber int64                                                                       `json:"ledger_sequence_number,required"`
	Metadata             map[string]string                                                           `json:"metadata,required"`
	StartingBalance      float64                                                                     `json:"starting_balance,required"`
	JSON                 customerCreditLedgerListByExternalIDResponseIncrementLedgerEntryJSON
}

// customerCreditLedgerListByExternalIDResponseIncrementLedgerEntryJSON contains
// the JSON metadata for the struct
// [CustomerCreditLedgerListByExternalIDResponseIncrementLedgerEntry]
type customerCreditLedgerListByExternalIDResponseIncrementLedgerEntryJSON struct {
	ID                   apijson.Field
	Amount               apijson.Field
	CreatedAt            apijson.Field
	CreditBlock          apijson.Field
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

func (r *CustomerCreditLedgerListByExternalIDResponseIncrementLedgerEntry) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r CustomerCreditLedgerListByExternalIDResponseIncrementLedgerEntry) implementsCustomerCreditLedgerListByExternalIDResponse() {
}

type CustomerCreditLedgerListByExternalIDResponseIncrementLedgerEntryCreditBlock struct {
	ID               string    `json:"id,required"`
	ExpiryDate       time.Time `json:"expiry_date,required,nullable" format:"date-time"`
	PerUnitCostBasis string    `json:"per_unit_cost_basis,required,nullable"`
	JSON             customerCreditLedgerListByExternalIDResponseIncrementLedgerEntryCreditBlockJSON
}

// customerCreditLedgerListByExternalIDResponseIncrementLedgerEntryCreditBlockJSON
// contains the JSON metadata for the struct
// [CustomerCreditLedgerListByExternalIDResponseIncrementLedgerEntryCreditBlock]
type customerCreditLedgerListByExternalIDResponseIncrementLedgerEntryCreditBlockJSON struct {
	ID               apijson.Field
	ExpiryDate       apijson.Field
	PerUnitCostBasis apijson.Field
	raw              string
	ExtraFields      map[string]apijson.Field
}

func (r *CustomerCreditLedgerListByExternalIDResponseIncrementLedgerEntryCreditBlock) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type CustomerCreditLedgerListByExternalIDResponseIncrementLedgerEntryCustomer struct {
	ID                 string `json:"id,required"`
	ExternalCustomerID string `json:"external_customer_id,required,nullable"`
	JSON               customerCreditLedgerListByExternalIDResponseIncrementLedgerEntryCustomerJSON
}

// customerCreditLedgerListByExternalIDResponseIncrementLedgerEntryCustomerJSON
// contains the JSON metadata for the struct
// [CustomerCreditLedgerListByExternalIDResponseIncrementLedgerEntryCustomer]
type customerCreditLedgerListByExternalIDResponseIncrementLedgerEntryCustomerJSON struct {
	ID                 apijson.Field
	ExternalCustomerID apijson.Field
	raw                string
	ExtraFields        map[string]apijson.Field
}

func (r *CustomerCreditLedgerListByExternalIDResponseIncrementLedgerEntryCustomer) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type CustomerCreditLedgerListByExternalIDResponseIncrementLedgerEntryEntryStatus string

const (
	CustomerCreditLedgerListByExternalIDResponseIncrementLedgerEntryEntryStatusCommitted CustomerCreditLedgerListByExternalIDResponseIncrementLedgerEntryEntryStatus = "committed"
	CustomerCreditLedgerListByExternalIDResponseIncrementLedgerEntryEntryStatusPending   CustomerCreditLedgerListByExternalIDResponseIncrementLedgerEntryEntryStatus = "pending"
)

type CustomerCreditLedgerListByExternalIDResponseIncrementLedgerEntryEntryType string

const (
	CustomerCreditLedgerListByExternalIDResponseIncrementLedgerEntryEntryTypeIncrement CustomerCreditLedgerListByExternalIDResponseIncrementLedgerEntryEntryType = "increment"
)

type CustomerCreditLedgerListByExternalIDResponseDecrementLedgerEntry struct {
	ID                   string                                                                      `json:"id,required"`
	Amount               float64                                                                     `json:"amount,required"`
	CreatedAt            time.Time                                                                   `json:"created_at,required" format:"date-time"`
	CreditBlock          CustomerCreditLedgerListByExternalIDResponseDecrementLedgerEntryCreditBlock `json:"credit_block,required"`
	Customer             CustomerCreditLedgerListByExternalIDResponseDecrementLedgerEntryCustomer    `json:"customer,required"`
	Description          string                                                                      `json:"description,required,nullable"`
	EndingBalance        float64                                                                     `json:"ending_balance,required"`
	EntryStatus          CustomerCreditLedgerListByExternalIDResponseDecrementLedgerEntryEntryStatus `json:"entry_status,required"`
	EntryType            CustomerCreditLedgerListByExternalIDResponseDecrementLedgerEntryEntryType   `json:"entry_type,required"`
	LedgerSequenceNumber int64                                                                       `json:"ledger_sequence_number,required"`
	Metadata             map[string]string                                                           `json:"metadata,required"`
	StartingBalance      float64                                                                     `json:"starting_balance,required"`
	EventID              string                                                                      `json:"event_id,nullable"`
	InvoiceID            string                                                                      `json:"invoice_id,nullable"`
	PriceID              string                                                                      `json:"price_id,nullable"`
	JSON                 customerCreditLedgerListByExternalIDResponseDecrementLedgerEntryJSON
}

// customerCreditLedgerListByExternalIDResponseDecrementLedgerEntryJSON contains
// the JSON metadata for the struct
// [CustomerCreditLedgerListByExternalIDResponseDecrementLedgerEntry]
type customerCreditLedgerListByExternalIDResponseDecrementLedgerEntryJSON struct {
	ID                   apijson.Field
	Amount               apijson.Field
	CreatedAt            apijson.Field
	CreditBlock          apijson.Field
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

func (r *CustomerCreditLedgerListByExternalIDResponseDecrementLedgerEntry) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r CustomerCreditLedgerListByExternalIDResponseDecrementLedgerEntry) implementsCustomerCreditLedgerListByExternalIDResponse() {
}

type CustomerCreditLedgerListByExternalIDResponseDecrementLedgerEntryCreditBlock struct {
	ID               string    `json:"id,required"`
	ExpiryDate       time.Time `json:"expiry_date,required,nullable" format:"date-time"`
	PerUnitCostBasis string    `json:"per_unit_cost_basis,required,nullable"`
	JSON             customerCreditLedgerListByExternalIDResponseDecrementLedgerEntryCreditBlockJSON
}

// customerCreditLedgerListByExternalIDResponseDecrementLedgerEntryCreditBlockJSON
// contains the JSON metadata for the struct
// [CustomerCreditLedgerListByExternalIDResponseDecrementLedgerEntryCreditBlock]
type customerCreditLedgerListByExternalIDResponseDecrementLedgerEntryCreditBlockJSON struct {
	ID               apijson.Field
	ExpiryDate       apijson.Field
	PerUnitCostBasis apijson.Field
	raw              string
	ExtraFields      map[string]apijson.Field
}

func (r *CustomerCreditLedgerListByExternalIDResponseDecrementLedgerEntryCreditBlock) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type CustomerCreditLedgerListByExternalIDResponseDecrementLedgerEntryCustomer struct {
	ID                 string `json:"id,required"`
	ExternalCustomerID string `json:"external_customer_id,required,nullable"`
	JSON               customerCreditLedgerListByExternalIDResponseDecrementLedgerEntryCustomerJSON
}

// customerCreditLedgerListByExternalIDResponseDecrementLedgerEntryCustomerJSON
// contains the JSON metadata for the struct
// [CustomerCreditLedgerListByExternalIDResponseDecrementLedgerEntryCustomer]
type customerCreditLedgerListByExternalIDResponseDecrementLedgerEntryCustomerJSON struct {
	ID                 apijson.Field
	ExternalCustomerID apijson.Field
	raw                string
	ExtraFields        map[string]apijson.Field
}

func (r *CustomerCreditLedgerListByExternalIDResponseDecrementLedgerEntryCustomer) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type CustomerCreditLedgerListByExternalIDResponseDecrementLedgerEntryEntryStatus string

const (
	CustomerCreditLedgerListByExternalIDResponseDecrementLedgerEntryEntryStatusCommitted CustomerCreditLedgerListByExternalIDResponseDecrementLedgerEntryEntryStatus = "committed"
	CustomerCreditLedgerListByExternalIDResponseDecrementLedgerEntryEntryStatusPending   CustomerCreditLedgerListByExternalIDResponseDecrementLedgerEntryEntryStatus = "pending"
)

type CustomerCreditLedgerListByExternalIDResponseDecrementLedgerEntryEntryType string

const (
	CustomerCreditLedgerListByExternalIDResponseDecrementLedgerEntryEntryTypeDecrement CustomerCreditLedgerListByExternalIDResponseDecrementLedgerEntryEntryType = "decrement"
)

type CustomerCreditLedgerListByExternalIDResponseExpirationChangeLedgerEntry struct {
	ID                   string                                                                             `json:"id,required"`
	Amount               float64                                                                            `json:"amount,required"`
	CreatedAt            time.Time                                                                          `json:"created_at,required" format:"date-time"`
	CreditBlock          CustomerCreditLedgerListByExternalIDResponseExpirationChangeLedgerEntryCreditBlock `json:"credit_block,required"`
	Customer             CustomerCreditLedgerListByExternalIDResponseExpirationChangeLedgerEntryCustomer    `json:"customer,required"`
	Description          string                                                                             `json:"description,required,nullable"`
	EndingBalance        float64                                                                            `json:"ending_balance,required"`
	EntryStatus          CustomerCreditLedgerListByExternalIDResponseExpirationChangeLedgerEntryEntryStatus `json:"entry_status,required"`
	EntryType            CustomerCreditLedgerListByExternalIDResponseExpirationChangeLedgerEntryEntryType   `json:"entry_type,required"`
	LedgerSequenceNumber int64                                                                              `json:"ledger_sequence_number,required"`
	Metadata             map[string]string                                                                  `json:"metadata,required"`
	NewBlockExpiryDate   time.Time                                                                          `json:"new_block_expiry_date,required" format:"date-time"`
	StartingBalance      float64                                                                            `json:"starting_balance,required"`
	JSON                 customerCreditLedgerListByExternalIDResponseExpirationChangeLedgerEntryJSON
}

// customerCreditLedgerListByExternalIDResponseExpirationChangeLedgerEntryJSON
// contains the JSON metadata for the struct
// [CustomerCreditLedgerListByExternalIDResponseExpirationChangeLedgerEntry]
type customerCreditLedgerListByExternalIDResponseExpirationChangeLedgerEntryJSON struct {
	ID                   apijson.Field
	Amount               apijson.Field
	CreatedAt            apijson.Field
	CreditBlock          apijson.Field
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

func (r *CustomerCreditLedgerListByExternalIDResponseExpirationChangeLedgerEntry) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r CustomerCreditLedgerListByExternalIDResponseExpirationChangeLedgerEntry) implementsCustomerCreditLedgerListByExternalIDResponse() {
}

type CustomerCreditLedgerListByExternalIDResponseExpirationChangeLedgerEntryCreditBlock struct {
	ID               string    `json:"id,required"`
	ExpiryDate       time.Time `json:"expiry_date,required,nullable" format:"date-time"`
	PerUnitCostBasis string    `json:"per_unit_cost_basis,required,nullable"`
	JSON             customerCreditLedgerListByExternalIDResponseExpirationChangeLedgerEntryCreditBlockJSON
}

// customerCreditLedgerListByExternalIDResponseExpirationChangeLedgerEntryCreditBlockJSON
// contains the JSON metadata for the struct
// [CustomerCreditLedgerListByExternalIDResponseExpirationChangeLedgerEntryCreditBlock]
type customerCreditLedgerListByExternalIDResponseExpirationChangeLedgerEntryCreditBlockJSON struct {
	ID               apijson.Field
	ExpiryDate       apijson.Field
	PerUnitCostBasis apijson.Field
	raw              string
	ExtraFields      map[string]apijson.Field
}

func (r *CustomerCreditLedgerListByExternalIDResponseExpirationChangeLedgerEntryCreditBlock) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type CustomerCreditLedgerListByExternalIDResponseExpirationChangeLedgerEntryCustomer struct {
	ID                 string `json:"id,required"`
	ExternalCustomerID string `json:"external_customer_id,required,nullable"`
	JSON               customerCreditLedgerListByExternalIDResponseExpirationChangeLedgerEntryCustomerJSON
}

// customerCreditLedgerListByExternalIDResponseExpirationChangeLedgerEntryCustomerJSON
// contains the JSON metadata for the struct
// [CustomerCreditLedgerListByExternalIDResponseExpirationChangeLedgerEntryCustomer]
type customerCreditLedgerListByExternalIDResponseExpirationChangeLedgerEntryCustomerJSON struct {
	ID                 apijson.Field
	ExternalCustomerID apijson.Field
	raw                string
	ExtraFields        map[string]apijson.Field
}

func (r *CustomerCreditLedgerListByExternalIDResponseExpirationChangeLedgerEntryCustomer) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type CustomerCreditLedgerListByExternalIDResponseExpirationChangeLedgerEntryEntryStatus string

const (
	CustomerCreditLedgerListByExternalIDResponseExpirationChangeLedgerEntryEntryStatusCommitted CustomerCreditLedgerListByExternalIDResponseExpirationChangeLedgerEntryEntryStatus = "committed"
	CustomerCreditLedgerListByExternalIDResponseExpirationChangeLedgerEntryEntryStatusPending   CustomerCreditLedgerListByExternalIDResponseExpirationChangeLedgerEntryEntryStatus = "pending"
)

type CustomerCreditLedgerListByExternalIDResponseExpirationChangeLedgerEntryEntryType string

const (
	CustomerCreditLedgerListByExternalIDResponseExpirationChangeLedgerEntryEntryTypeExpirationChange CustomerCreditLedgerListByExternalIDResponseExpirationChangeLedgerEntryEntryType = "expiration_change"
)

type CustomerCreditLedgerListByExternalIDResponseCreditBlockExpiryLedgerEntry struct {
	ID                   string                                                                              `json:"id,required"`
	Amount               float64                                                                             `json:"amount,required"`
	CreatedAt            time.Time                                                                           `json:"created_at,required" format:"date-time"`
	CreditBlock          CustomerCreditLedgerListByExternalIDResponseCreditBlockExpiryLedgerEntryCreditBlock `json:"credit_block,required"`
	Customer             CustomerCreditLedgerListByExternalIDResponseCreditBlockExpiryLedgerEntryCustomer    `json:"customer,required"`
	Description          string                                                                              `json:"description,required,nullable"`
	EndingBalance        float64                                                                             `json:"ending_balance,required"`
	EntryStatus          CustomerCreditLedgerListByExternalIDResponseCreditBlockExpiryLedgerEntryEntryStatus `json:"entry_status,required"`
	EntryType            CustomerCreditLedgerListByExternalIDResponseCreditBlockExpiryLedgerEntryEntryType   `json:"entry_type,required"`
	LedgerSequenceNumber int64                                                                               `json:"ledger_sequence_number,required"`
	Metadata             map[string]string                                                                   `json:"metadata,required"`
	StartingBalance      float64                                                                             `json:"starting_balance,required"`
	JSON                 customerCreditLedgerListByExternalIDResponseCreditBlockExpiryLedgerEntryJSON
}

// customerCreditLedgerListByExternalIDResponseCreditBlockExpiryLedgerEntryJSON
// contains the JSON metadata for the struct
// [CustomerCreditLedgerListByExternalIDResponseCreditBlockExpiryLedgerEntry]
type customerCreditLedgerListByExternalIDResponseCreditBlockExpiryLedgerEntryJSON struct {
	ID                   apijson.Field
	Amount               apijson.Field
	CreatedAt            apijson.Field
	CreditBlock          apijson.Field
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

func (r *CustomerCreditLedgerListByExternalIDResponseCreditBlockExpiryLedgerEntry) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r CustomerCreditLedgerListByExternalIDResponseCreditBlockExpiryLedgerEntry) implementsCustomerCreditLedgerListByExternalIDResponse() {
}

type CustomerCreditLedgerListByExternalIDResponseCreditBlockExpiryLedgerEntryCreditBlock struct {
	ID               string    `json:"id,required"`
	ExpiryDate       time.Time `json:"expiry_date,required,nullable" format:"date-time"`
	PerUnitCostBasis string    `json:"per_unit_cost_basis,required,nullable"`
	JSON             customerCreditLedgerListByExternalIDResponseCreditBlockExpiryLedgerEntryCreditBlockJSON
}

// customerCreditLedgerListByExternalIDResponseCreditBlockExpiryLedgerEntryCreditBlockJSON
// contains the JSON metadata for the struct
// [CustomerCreditLedgerListByExternalIDResponseCreditBlockExpiryLedgerEntryCreditBlock]
type customerCreditLedgerListByExternalIDResponseCreditBlockExpiryLedgerEntryCreditBlockJSON struct {
	ID               apijson.Field
	ExpiryDate       apijson.Field
	PerUnitCostBasis apijson.Field
	raw              string
	ExtraFields      map[string]apijson.Field
}

func (r *CustomerCreditLedgerListByExternalIDResponseCreditBlockExpiryLedgerEntryCreditBlock) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type CustomerCreditLedgerListByExternalIDResponseCreditBlockExpiryLedgerEntryCustomer struct {
	ID                 string `json:"id,required"`
	ExternalCustomerID string `json:"external_customer_id,required,nullable"`
	JSON               customerCreditLedgerListByExternalIDResponseCreditBlockExpiryLedgerEntryCustomerJSON
}

// customerCreditLedgerListByExternalIDResponseCreditBlockExpiryLedgerEntryCustomerJSON
// contains the JSON metadata for the struct
// [CustomerCreditLedgerListByExternalIDResponseCreditBlockExpiryLedgerEntryCustomer]
type customerCreditLedgerListByExternalIDResponseCreditBlockExpiryLedgerEntryCustomerJSON struct {
	ID                 apijson.Field
	ExternalCustomerID apijson.Field
	raw                string
	ExtraFields        map[string]apijson.Field
}

func (r *CustomerCreditLedgerListByExternalIDResponseCreditBlockExpiryLedgerEntryCustomer) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type CustomerCreditLedgerListByExternalIDResponseCreditBlockExpiryLedgerEntryEntryStatus string

const (
	CustomerCreditLedgerListByExternalIDResponseCreditBlockExpiryLedgerEntryEntryStatusCommitted CustomerCreditLedgerListByExternalIDResponseCreditBlockExpiryLedgerEntryEntryStatus = "committed"
	CustomerCreditLedgerListByExternalIDResponseCreditBlockExpiryLedgerEntryEntryStatusPending   CustomerCreditLedgerListByExternalIDResponseCreditBlockExpiryLedgerEntryEntryStatus = "pending"
)

type CustomerCreditLedgerListByExternalIDResponseCreditBlockExpiryLedgerEntryEntryType string

const (
	CustomerCreditLedgerListByExternalIDResponseCreditBlockExpiryLedgerEntryEntryTypeCreditBlockExpiry CustomerCreditLedgerListByExternalIDResponseCreditBlockExpiryLedgerEntryEntryType = "credit_block_expiry"
)

type CustomerCreditLedgerListByExternalIDResponseVoidLedgerEntry struct {
	ID                   string                                                                 `json:"id,required"`
	Amount               float64                                                                `json:"amount,required"`
	CreatedAt            time.Time                                                              `json:"created_at,required" format:"date-time"`
	CreditBlock          CustomerCreditLedgerListByExternalIDResponseVoidLedgerEntryCreditBlock `json:"credit_block,required"`
	Customer             CustomerCreditLedgerListByExternalIDResponseVoidLedgerEntryCustomer    `json:"customer,required"`
	Description          string                                                                 `json:"description,required,nullable"`
	EndingBalance        float64                                                                `json:"ending_balance,required"`
	EntryStatus          CustomerCreditLedgerListByExternalIDResponseVoidLedgerEntryEntryStatus `json:"entry_status,required"`
	EntryType            CustomerCreditLedgerListByExternalIDResponseVoidLedgerEntryEntryType   `json:"entry_type,required"`
	LedgerSequenceNumber int64                                                                  `json:"ledger_sequence_number,required"`
	Metadata             map[string]string                                                      `json:"metadata,required"`
	StartingBalance      float64                                                                `json:"starting_balance,required"`
	VoidAmount           float64                                                                `json:"void_amount,required"`
	VoidReason           string                                                                 `json:"void_reason,required,nullable"`
	JSON                 customerCreditLedgerListByExternalIDResponseVoidLedgerEntryJSON
}

// customerCreditLedgerListByExternalIDResponseVoidLedgerEntryJSON contains the
// JSON metadata for the struct
// [CustomerCreditLedgerListByExternalIDResponseVoidLedgerEntry]
type customerCreditLedgerListByExternalIDResponseVoidLedgerEntryJSON struct {
	ID                   apijson.Field
	Amount               apijson.Field
	CreatedAt            apijson.Field
	CreditBlock          apijson.Field
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

func (r *CustomerCreditLedgerListByExternalIDResponseVoidLedgerEntry) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r CustomerCreditLedgerListByExternalIDResponseVoidLedgerEntry) implementsCustomerCreditLedgerListByExternalIDResponse() {
}

type CustomerCreditLedgerListByExternalIDResponseVoidLedgerEntryCreditBlock struct {
	ID               string    `json:"id,required"`
	ExpiryDate       time.Time `json:"expiry_date,required,nullable" format:"date-time"`
	PerUnitCostBasis string    `json:"per_unit_cost_basis,required,nullable"`
	JSON             customerCreditLedgerListByExternalIDResponseVoidLedgerEntryCreditBlockJSON
}

// customerCreditLedgerListByExternalIDResponseVoidLedgerEntryCreditBlockJSON
// contains the JSON metadata for the struct
// [CustomerCreditLedgerListByExternalIDResponseVoidLedgerEntryCreditBlock]
type customerCreditLedgerListByExternalIDResponseVoidLedgerEntryCreditBlockJSON struct {
	ID               apijson.Field
	ExpiryDate       apijson.Field
	PerUnitCostBasis apijson.Field
	raw              string
	ExtraFields      map[string]apijson.Field
}

func (r *CustomerCreditLedgerListByExternalIDResponseVoidLedgerEntryCreditBlock) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type CustomerCreditLedgerListByExternalIDResponseVoidLedgerEntryCustomer struct {
	ID                 string `json:"id,required"`
	ExternalCustomerID string `json:"external_customer_id,required,nullable"`
	JSON               customerCreditLedgerListByExternalIDResponseVoidLedgerEntryCustomerJSON
}

// customerCreditLedgerListByExternalIDResponseVoidLedgerEntryCustomerJSON contains
// the JSON metadata for the struct
// [CustomerCreditLedgerListByExternalIDResponseVoidLedgerEntryCustomer]
type customerCreditLedgerListByExternalIDResponseVoidLedgerEntryCustomerJSON struct {
	ID                 apijson.Field
	ExternalCustomerID apijson.Field
	raw                string
	ExtraFields        map[string]apijson.Field
}

func (r *CustomerCreditLedgerListByExternalIDResponseVoidLedgerEntryCustomer) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type CustomerCreditLedgerListByExternalIDResponseVoidLedgerEntryEntryStatus string

const (
	CustomerCreditLedgerListByExternalIDResponseVoidLedgerEntryEntryStatusCommitted CustomerCreditLedgerListByExternalIDResponseVoidLedgerEntryEntryStatus = "committed"
	CustomerCreditLedgerListByExternalIDResponseVoidLedgerEntryEntryStatusPending   CustomerCreditLedgerListByExternalIDResponseVoidLedgerEntryEntryStatus = "pending"
)

type CustomerCreditLedgerListByExternalIDResponseVoidLedgerEntryEntryType string

const (
	CustomerCreditLedgerListByExternalIDResponseVoidLedgerEntryEntryTypeVoid CustomerCreditLedgerListByExternalIDResponseVoidLedgerEntryEntryType = "void"
)

type CustomerCreditLedgerListByExternalIDResponseVoidInitiatedLedgerEntry struct {
	ID                   string                                                                          `json:"id,required"`
	Amount               float64                                                                         `json:"amount,required"`
	CreatedAt            time.Time                                                                       `json:"created_at,required" format:"date-time"`
	CreditBlock          CustomerCreditLedgerListByExternalIDResponseVoidInitiatedLedgerEntryCreditBlock `json:"credit_block,required"`
	Customer             CustomerCreditLedgerListByExternalIDResponseVoidInitiatedLedgerEntryCustomer    `json:"customer,required"`
	Description          string                                                                          `json:"description,required,nullable"`
	EndingBalance        float64                                                                         `json:"ending_balance,required"`
	EntryStatus          CustomerCreditLedgerListByExternalIDResponseVoidInitiatedLedgerEntryEntryStatus `json:"entry_status,required"`
	EntryType            CustomerCreditLedgerListByExternalIDResponseVoidInitiatedLedgerEntryEntryType   `json:"entry_type,required"`
	LedgerSequenceNumber int64                                                                           `json:"ledger_sequence_number,required"`
	Metadata             map[string]string                                                               `json:"metadata,required"`
	NewBlockExpiryDate   time.Time                                                                       `json:"new_block_expiry_date,required" format:"date-time"`
	StartingBalance      float64                                                                         `json:"starting_balance,required"`
	VoidAmount           float64                                                                         `json:"void_amount,required"`
	VoidReason           string                                                                          `json:"void_reason,required,nullable"`
	JSON                 customerCreditLedgerListByExternalIDResponseVoidInitiatedLedgerEntryJSON
}

// customerCreditLedgerListByExternalIDResponseVoidInitiatedLedgerEntryJSON
// contains the JSON metadata for the struct
// [CustomerCreditLedgerListByExternalIDResponseVoidInitiatedLedgerEntry]
type customerCreditLedgerListByExternalIDResponseVoidInitiatedLedgerEntryJSON struct {
	ID                   apijson.Field
	Amount               apijson.Field
	CreatedAt            apijson.Field
	CreditBlock          apijson.Field
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

func (r *CustomerCreditLedgerListByExternalIDResponseVoidInitiatedLedgerEntry) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r CustomerCreditLedgerListByExternalIDResponseVoidInitiatedLedgerEntry) implementsCustomerCreditLedgerListByExternalIDResponse() {
}

type CustomerCreditLedgerListByExternalIDResponseVoidInitiatedLedgerEntryCreditBlock struct {
	ID               string    `json:"id,required"`
	ExpiryDate       time.Time `json:"expiry_date,required,nullable" format:"date-time"`
	PerUnitCostBasis string    `json:"per_unit_cost_basis,required,nullable"`
	JSON             customerCreditLedgerListByExternalIDResponseVoidInitiatedLedgerEntryCreditBlockJSON
}

// customerCreditLedgerListByExternalIDResponseVoidInitiatedLedgerEntryCreditBlockJSON
// contains the JSON metadata for the struct
// [CustomerCreditLedgerListByExternalIDResponseVoidInitiatedLedgerEntryCreditBlock]
type customerCreditLedgerListByExternalIDResponseVoidInitiatedLedgerEntryCreditBlockJSON struct {
	ID               apijson.Field
	ExpiryDate       apijson.Field
	PerUnitCostBasis apijson.Field
	raw              string
	ExtraFields      map[string]apijson.Field
}

func (r *CustomerCreditLedgerListByExternalIDResponseVoidInitiatedLedgerEntryCreditBlock) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type CustomerCreditLedgerListByExternalIDResponseVoidInitiatedLedgerEntryCustomer struct {
	ID                 string `json:"id,required"`
	ExternalCustomerID string `json:"external_customer_id,required,nullable"`
	JSON               customerCreditLedgerListByExternalIDResponseVoidInitiatedLedgerEntryCustomerJSON
}

// customerCreditLedgerListByExternalIDResponseVoidInitiatedLedgerEntryCustomerJSON
// contains the JSON metadata for the struct
// [CustomerCreditLedgerListByExternalIDResponseVoidInitiatedLedgerEntryCustomer]
type customerCreditLedgerListByExternalIDResponseVoidInitiatedLedgerEntryCustomerJSON struct {
	ID                 apijson.Field
	ExternalCustomerID apijson.Field
	raw                string
	ExtraFields        map[string]apijson.Field
}

func (r *CustomerCreditLedgerListByExternalIDResponseVoidInitiatedLedgerEntryCustomer) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type CustomerCreditLedgerListByExternalIDResponseVoidInitiatedLedgerEntryEntryStatus string

const (
	CustomerCreditLedgerListByExternalIDResponseVoidInitiatedLedgerEntryEntryStatusCommitted CustomerCreditLedgerListByExternalIDResponseVoidInitiatedLedgerEntryEntryStatus = "committed"
	CustomerCreditLedgerListByExternalIDResponseVoidInitiatedLedgerEntryEntryStatusPending   CustomerCreditLedgerListByExternalIDResponseVoidInitiatedLedgerEntryEntryStatus = "pending"
)

type CustomerCreditLedgerListByExternalIDResponseVoidInitiatedLedgerEntryEntryType string

const (
	CustomerCreditLedgerListByExternalIDResponseVoidInitiatedLedgerEntryEntryTypeVoidInitiated CustomerCreditLedgerListByExternalIDResponseVoidInitiatedLedgerEntryEntryType = "void_initiated"
)

type CustomerCreditLedgerListByExternalIDResponseAmendmentLedgerEntry struct {
	ID                   string                                                                      `json:"id,required"`
	Amount               float64                                                                     `json:"amount,required"`
	CreatedAt            time.Time                                                                   `json:"created_at,required" format:"date-time"`
	CreditBlock          CustomerCreditLedgerListByExternalIDResponseAmendmentLedgerEntryCreditBlock `json:"credit_block,required"`
	Customer             CustomerCreditLedgerListByExternalIDResponseAmendmentLedgerEntryCustomer    `json:"customer,required"`
	Description          string                                                                      `json:"description,required,nullable"`
	EndingBalance        float64                                                                     `json:"ending_balance,required"`
	EntryStatus          CustomerCreditLedgerListByExternalIDResponseAmendmentLedgerEntryEntryStatus `json:"entry_status,required"`
	EntryType            CustomerCreditLedgerListByExternalIDResponseAmendmentLedgerEntryEntryType   `json:"entry_type,required"`
	LedgerSequenceNumber int64                                                                       `json:"ledger_sequence_number,required"`
	Metadata             map[string]string                                                           `json:"metadata,required"`
	StartingBalance      float64                                                                     `json:"starting_balance,required"`
	JSON                 customerCreditLedgerListByExternalIDResponseAmendmentLedgerEntryJSON
}

// customerCreditLedgerListByExternalIDResponseAmendmentLedgerEntryJSON contains
// the JSON metadata for the struct
// [CustomerCreditLedgerListByExternalIDResponseAmendmentLedgerEntry]
type customerCreditLedgerListByExternalIDResponseAmendmentLedgerEntryJSON struct {
	ID                   apijson.Field
	Amount               apijson.Field
	CreatedAt            apijson.Field
	CreditBlock          apijson.Field
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

func (r *CustomerCreditLedgerListByExternalIDResponseAmendmentLedgerEntry) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r CustomerCreditLedgerListByExternalIDResponseAmendmentLedgerEntry) implementsCustomerCreditLedgerListByExternalIDResponse() {
}

type CustomerCreditLedgerListByExternalIDResponseAmendmentLedgerEntryCreditBlock struct {
	ID               string    `json:"id,required"`
	ExpiryDate       time.Time `json:"expiry_date,required,nullable" format:"date-time"`
	PerUnitCostBasis string    `json:"per_unit_cost_basis,required,nullable"`
	JSON             customerCreditLedgerListByExternalIDResponseAmendmentLedgerEntryCreditBlockJSON
}

// customerCreditLedgerListByExternalIDResponseAmendmentLedgerEntryCreditBlockJSON
// contains the JSON metadata for the struct
// [CustomerCreditLedgerListByExternalIDResponseAmendmentLedgerEntryCreditBlock]
type customerCreditLedgerListByExternalIDResponseAmendmentLedgerEntryCreditBlockJSON struct {
	ID               apijson.Field
	ExpiryDate       apijson.Field
	PerUnitCostBasis apijson.Field
	raw              string
	ExtraFields      map[string]apijson.Field
}

func (r *CustomerCreditLedgerListByExternalIDResponseAmendmentLedgerEntryCreditBlock) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type CustomerCreditLedgerListByExternalIDResponseAmendmentLedgerEntryCustomer struct {
	ID                 string `json:"id,required"`
	ExternalCustomerID string `json:"external_customer_id,required,nullable"`
	JSON               customerCreditLedgerListByExternalIDResponseAmendmentLedgerEntryCustomerJSON
}

// customerCreditLedgerListByExternalIDResponseAmendmentLedgerEntryCustomerJSON
// contains the JSON metadata for the struct
// [CustomerCreditLedgerListByExternalIDResponseAmendmentLedgerEntryCustomer]
type customerCreditLedgerListByExternalIDResponseAmendmentLedgerEntryCustomerJSON struct {
	ID                 apijson.Field
	ExternalCustomerID apijson.Field
	raw                string
	ExtraFields        map[string]apijson.Field
}

func (r *CustomerCreditLedgerListByExternalIDResponseAmendmentLedgerEntryCustomer) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type CustomerCreditLedgerListByExternalIDResponseAmendmentLedgerEntryEntryStatus string

const (
	CustomerCreditLedgerListByExternalIDResponseAmendmentLedgerEntryEntryStatusCommitted CustomerCreditLedgerListByExternalIDResponseAmendmentLedgerEntryEntryStatus = "committed"
	CustomerCreditLedgerListByExternalIDResponseAmendmentLedgerEntryEntryStatusPending   CustomerCreditLedgerListByExternalIDResponseAmendmentLedgerEntryEntryStatus = "pending"
)

type CustomerCreditLedgerListByExternalIDResponseAmendmentLedgerEntryEntryType string

const (
	CustomerCreditLedgerListByExternalIDResponseAmendmentLedgerEntryEntryTypeAmendment CustomerCreditLedgerListByExternalIDResponseAmendmentLedgerEntryEntryType = "amendment"
)

type CustomerCreditLedgerListParams struct {
	CreatedAtGt  param.Field[time.Time] `query:"created_at[gt]" format:"date-time"`
	CreatedAtGte param.Field[time.Time] `query:"created_at[gte]" format:"date-time"`
	CreatedAtLt  param.Field[time.Time] `query:"created_at[lt]" format:"date-time"`
	CreatedAtLte param.Field[time.Time] `query:"created_at[lte]" format:"date-time"`
	Currency     param.Field[string]    `query:"currency"`
	// Cursor for pagination. This can be populated by the `next_cursor` value returned
	// from the initial request.
	Cursor      param.Field[string]                                    `query:"cursor"`
	EntryStatus param.Field[CustomerCreditLedgerListParamsEntryStatus] `query:"entry_status"`
	EntryType   param.Field[CustomerCreditLedgerListParamsEntryType]   `query:"entry_type"`
	// The number of items to fetch. Defaults to 20.
	Limit         param.Field[int64]  `query:"limit"`
	MinimumAmount param.Field[string] `query:"minimum_amount"`
}

// URLQuery serializes [CustomerCreditLedgerListParams]'s query parameters as
// `url.Values`.
func (r CustomerCreditLedgerListParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type CustomerCreditLedgerListParamsEntryStatus string

const (
	CustomerCreditLedgerListParamsEntryStatusCommitted CustomerCreditLedgerListParamsEntryStatus = "committed"
	CustomerCreditLedgerListParamsEntryStatusPending   CustomerCreditLedgerListParamsEntryStatus = "pending"
)

type CustomerCreditLedgerListParamsEntryType string

const (
	CustomerCreditLedgerListParamsEntryTypeIncrement         CustomerCreditLedgerListParamsEntryType = "increment"
	CustomerCreditLedgerListParamsEntryTypeDecrement         CustomerCreditLedgerListParamsEntryType = "decrement"
	CustomerCreditLedgerListParamsEntryTypeExpirationChange  CustomerCreditLedgerListParamsEntryType = "expiration_change"
	CustomerCreditLedgerListParamsEntryTypeCreditBlockExpiry CustomerCreditLedgerListParamsEntryType = "credit_block_expiry"
	CustomerCreditLedgerListParamsEntryTypeVoid              CustomerCreditLedgerListParamsEntryType = "void"
	CustomerCreditLedgerListParamsEntryTypeVoidInitiated     CustomerCreditLedgerListParamsEntryType = "void_initiated"
	CustomerCreditLedgerListParamsEntryTypeAmendment         CustomerCreditLedgerListParamsEntryType = "amendment"
)

// This interface is a union satisfied by one of the following:
// [CustomerCreditLedgerNewEntryParamsAddIncrementCreditLedgerEntryRequestParams],
// [CustomerCreditLedgerNewEntryParamsAddDecrementCreditLedgerEntryRequestParams],
// [CustomerCreditLedgerNewEntryParamsAddExpirationChangeCreditLedgerEntryRequestParams],
// [CustomerCreditLedgerNewEntryParamsAddVoidCreditLedgerEntryRequestParams],
// [CustomerCreditLedgerNewEntryParamsAddAmendmentCreditLedgerEntryRequestParams].
type CustomerCreditLedgerNewEntryParams interface {
	ImplementsCustomerCreditLedgerNewEntryParams()
}

type CustomerCreditLedgerNewEntryParamsAddIncrementCreditLedgerEntryRequestParams struct {
	// The number of credits to effect. Note that this is required for increment,
	// decrement, void, or undo operations.
	Amount    param.Field[float64]                                                                               `json:"amount,required"`
	EntryType param.Field[CustomerCreditLedgerNewEntryParamsAddIncrementCreditLedgerEntryRequestParamsEntryType] `json:"entry_type,required"`
	// Optional metadata that can be specified when adding ledger results via the API.
	// For example, this can be used to note an increment refers to trial credits, or
	// for noting corrections as a result of an incident, etc.
	Description param.Field[string] `json:"description"`
	// A future date (specified in YYYY-MM-DD format) that denotes when this credit
	// balance should become available for use.
	EffectiveDate param.Field[time.Time] `json:"effective_date" format:"date"`
	// A future date (specified in YYYY-MM-DD format) that denotes when this credit
	// balance should expire.
	ExpiryDate param.Field[time.Time] `json:"expiry_date" format:"date"`
	// Passing `invoice_settings` automatically generates an invoice for the newly
	// added credits. If `invoice_settings` is passed, you must specify
	// per_unit_cost_basis, as the calculation of the invoice total is done on that
	// basis.
	InvoiceSettings param.Field[CustomerCreditLedgerNewEntryParamsAddIncrementCreditLedgerEntryRequestParamsInvoiceSettings] `json:"invoice_settings"`
	// User-specified key/value pairs for the ledger entry resource.
	Metadata param.Field[interface{}] `json:"metadata"`
	// Can only be specified when entry_type=increment. How much, in USD, a customer
	// paid for a single credit in this block
	PerUnitCostBasis param.Field[string] `json:"per_unit_cost_basis"`
}

func (r CustomerCreditLedgerNewEntryParamsAddIncrementCreditLedgerEntryRequestParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (CustomerCreditLedgerNewEntryParamsAddIncrementCreditLedgerEntryRequestParams) ImplementsCustomerCreditLedgerNewEntryParams() {

}

type CustomerCreditLedgerNewEntryParamsAddIncrementCreditLedgerEntryRequestParamsEntryType string

const (
	CustomerCreditLedgerNewEntryParamsAddIncrementCreditLedgerEntryRequestParamsEntryTypeIncrement CustomerCreditLedgerNewEntryParamsAddIncrementCreditLedgerEntryRequestParamsEntryType = "increment"
)

// Passing `invoice_settings` automatically generates an invoice for the newly
// added credits. If `invoice_settings` is passed, you must specify
// per_unit_cost_basis, as the calculation of the invoice total is done on that
// basis.
type CustomerCreditLedgerNewEntryParamsAddIncrementCreditLedgerEntryRequestParamsInvoiceSettings struct {
	// Whether the credits purchase invoice should auto collect with the customer's
	// saved payment method.
	AutoCollection param.Field[bool] `json:"auto_collection,required"`
	// The net terms determines the difference between the invoice date and the issue
	// date for the invoice. If you intend the invoice to be due on issue, set this
	// to 0.
	NetTerms param.Field[int64] `json:"net_terms,required"`
	// An optional memo to display on the invoice.
	Memo param.Field[string] `json:"memo"`
}

func (r CustomerCreditLedgerNewEntryParamsAddIncrementCreditLedgerEntryRequestParamsInvoiceSettings) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type CustomerCreditLedgerNewEntryParamsAddDecrementCreditLedgerEntryRequestParams struct {
	// The number of credits to effect. Note that this is required for increment,
	// decrement, void, or undo operations.
	Amount    param.Field[float64]                                                                               `json:"amount,required"`
	EntryType param.Field[CustomerCreditLedgerNewEntryParamsAddDecrementCreditLedgerEntryRequestParamsEntryType] `json:"entry_type,required"`
	// Optional metadata that can be specified when adding ledger results via the API.
	// For example, this can be used to note an increment refers to trial credits, or
	// for noting corrections as a result of an incident, etc.
	Description param.Field[string] `json:"description"`
	// User-specified key/value pairs for the ledger entry resource.
	Metadata param.Field[interface{}] `json:"metadata"`
}

func (r CustomerCreditLedgerNewEntryParamsAddDecrementCreditLedgerEntryRequestParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (CustomerCreditLedgerNewEntryParamsAddDecrementCreditLedgerEntryRequestParams) ImplementsCustomerCreditLedgerNewEntryParams() {

}

type CustomerCreditLedgerNewEntryParamsAddDecrementCreditLedgerEntryRequestParamsEntryType string

const (
	CustomerCreditLedgerNewEntryParamsAddDecrementCreditLedgerEntryRequestParamsEntryTypeDecrement CustomerCreditLedgerNewEntryParamsAddDecrementCreditLedgerEntryRequestParamsEntryType = "decrement"
)

type CustomerCreditLedgerNewEntryParamsAddExpirationChangeCreditLedgerEntryRequestParams struct {
	EntryType param.Field[CustomerCreditLedgerNewEntryParamsAddExpirationChangeCreditLedgerEntryRequestParamsEntryType] `json:"entry_type,required"`
	// A future date (specified in YYYY-MM-DD format) that identifies the origination
	// credit block to expire
	ExpiryDate param.Field[time.Time] `json:"expiry_date,required" format:"date"`
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
	// Optional metadata that can be specified when adding ledger results via the API.
	// For example, this can be used to note an increment refers to trial credits, or
	// for noting corrections as a result of an incident, etc.
	Description param.Field[string] `json:"description"`
	// User-specified key/value pairs for the ledger entry resource.
	Metadata param.Field[interface{}] `json:"metadata"`
}

func (r CustomerCreditLedgerNewEntryParamsAddExpirationChangeCreditLedgerEntryRequestParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (CustomerCreditLedgerNewEntryParamsAddExpirationChangeCreditLedgerEntryRequestParams) ImplementsCustomerCreditLedgerNewEntryParams() {

}

type CustomerCreditLedgerNewEntryParamsAddExpirationChangeCreditLedgerEntryRequestParamsEntryType string

const (
	CustomerCreditLedgerNewEntryParamsAddExpirationChangeCreditLedgerEntryRequestParamsEntryTypeExpirationChange CustomerCreditLedgerNewEntryParamsAddExpirationChangeCreditLedgerEntryRequestParamsEntryType = "expiration_change"
)

type CustomerCreditLedgerNewEntryParamsAddVoidCreditLedgerEntryRequestParams struct {
	// The number of credits to effect. Note that this is required for increment,
	// decrement, void, or undo operations.
	Amount param.Field[float64] `json:"amount,required"`
	// The ID of the block to void.
	BlockID   param.Field[string]                                                                           `json:"block_id,required"`
	EntryType param.Field[CustomerCreditLedgerNewEntryParamsAddVoidCreditLedgerEntryRequestParamsEntryType] `json:"entry_type,required"`
	// Optional metadata that can be specified when adding ledger results via the API.
	// For example, this can be used to note an increment refers to trial credits, or
	// for noting corrections as a result of an incident, etc.
	Description param.Field[string] `json:"description"`
	// User-specified key/value pairs for the ledger entry resource.
	Metadata param.Field[interface{}] `json:"metadata"`
	// Can only be specified when `entry_type=void`. The reason for the void.
	VoidReason param.Field[CustomerCreditLedgerNewEntryParamsAddVoidCreditLedgerEntryRequestParamsVoidReason] `json:"void_reason"`
}

func (r CustomerCreditLedgerNewEntryParamsAddVoidCreditLedgerEntryRequestParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (CustomerCreditLedgerNewEntryParamsAddVoidCreditLedgerEntryRequestParams) ImplementsCustomerCreditLedgerNewEntryParams() {

}

type CustomerCreditLedgerNewEntryParamsAddVoidCreditLedgerEntryRequestParamsEntryType string

const (
	CustomerCreditLedgerNewEntryParamsAddVoidCreditLedgerEntryRequestParamsEntryTypeVoid CustomerCreditLedgerNewEntryParamsAddVoidCreditLedgerEntryRequestParamsEntryType = "void"
)

// Can only be specified when `entry_type=void`. The reason for the void.
type CustomerCreditLedgerNewEntryParamsAddVoidCreditLedgerEntryRequestParamsVoidReason string

const (
	CustomerCreditLedgerNewEntryParamsAddVoidCreditLedgerEntryRequestParamsVoidReasonRefund CustomerCreditLedgerNewEntryParamsAddVoidCreditLedgerEntryRequestParamsVoidReason = "refund"
)

type CustomerCreditLedgerNewEntryParamsAddAmendmentCreditLedgerEntryRequestParams struct {
	// The number of credits to effect. Note that this is required for increment,
	// decrement or void operations.
	Amount param.Field[float64] `json:"amount,required"`
	// The ID of the block to reverse a decrement from.
	BlockID   param.Field[string]                                                                                `json:"block_id,required"`
	EntryType param.Field[CustomerCreditLedgerNewEntryParamsAddAmendmentCreditLedgerEntryRequestParamsEntryType] `json:"entry_type,required"`
	// Optional metadata that can be specified when adding ledger results via the API.
	// For example, this can be used to note an increment refers to trial credits, or
	// for noting corrections as a result of an incident, etc.
	Description param.Field[string] `json:"description"`
	// User-specified key/value pairs for the ledger entry resource.
	Metadata param.Field[interface{}] `json:"metadata"`
}

func (r CustomerCreditLedgerNewEntryParamsAddAmendmentCreditLedgerEntryRequestParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (CustomerCreditLedgerNewEntryParamsAddAmendmentCreditLedgerEntryRequestParams) ImplementsCustomerCreditLedgerNewEntryParams() {

}

type CustomerCreditLedgerNewEntryParamsAddAmendmentCreditLedgerEntryRequestParamsEntryType string

const (
	CustomerCreditLedgerNewEntryParamsAddAmendmentCreditLedgerEntryRequestParamsEntryTypeAmendment CustomerCreditLedgerNewEntryParamsAddAmendmentCreditLedgerEntryRequestParamsEntryType = "amendment"
)

// This interface is a union satisfied by one of the following:
// [CustomerCreditLedgerNewEntryByExteralIDParamsAddIncrementCreditLedgerEntryRequestParams],
// [CustomerCreditLedgerNewEntryByExteralIDParamsAddDecrementCreditLedgerEntryRequestParams],
// [CustomerCreditLedgerNewEntryByExteralIDParamsAddExpirationChangeCreditLedgerEntryRequestParams],
// [CustomerCreditLedgerNewEntryByExteralIDParamsAddVoidCreditLedgerEntryRequestParams],
// [CustomerCreditLedgerNewEntryByExteralIDParamsAddAmendmentCreditLedgerEntryRequestParams].
type CustomerCreditLedgerNewEntryByExteralIDParams interface {
	ImplementsCustomerCreditLedgerNewEntryByExteralIDParams()
}

type CustomerCreditLedgerNewEntryByExteralIDParamsAddIncrementCreditLedgerEntryRequestParams struct {
	// The number of credits to effect. Note that this is required for increment,
	// decrement, void, or undo operations.
	Amount    param.Field[float64]                                                                                          `json:"amount,required"`
	EntryType param.Field[CustomerCreditLedgerNewEntryByExteralIDParamsAddIncrementCreditLedgerEntryRequestParamsEntryType] `json:"entry_type,required"`
	// Optional metadata that can be specified when adding ledger results via the API.
	// For example, this can be used to note an increment refers to trial credits, or
	// for noting corrections as a result of an incident, etc.
	Description param.Field[string] `json:"description"`
	// A future date (specified in YYYY-MM-DD format) that denotes when this credit
	// balance should become available for use.
	EffectiveDate param.Field[time.Time] `json:"effective_date" format:"date"`
	// A future date (specified in YYYY-MM-DD format) that denotes when this credit
	// balance should expire.
	ExpiryDate param.Field[time.Time] `json:"expiry_date" format:"date"`
	// Passing `invoice_settings` automatically generates an invoice for the newly
	// added credits. If `invoice_settings` is passed, you must specify
	// per_unit_cost_basis, as the calculation of the invoice total is done on that
	// basis.
	InvoiceSettings param.Field[CustomerCreditLedgerNewEntryByExteralIDParamsAddIncrementCreditLedgerEntryRequestParamsInvoiceSettings] `json:"invoice_settings"`
	// User-specified key/value pairs for the ledger entry resource.
	Metadata param.Field[interface{}] `json:"metadata"`
	// Can only be specified when entry_type=increment. How much, in USD, a customer
	// paid for a single credit in this block
	PerUnitCostBasis param.Field[string] `json:"per_unit_cost_basis"`
}

func (r CustomerCreditLedgerNewEntryByExteralIDParamsAddIncrementCreditLedgerEntryRequestParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (CustomerCreditLedgerNewEntryByExteralIDParamsAddIncrementCreditLedgerEntryRequestParams) ImplementsCustomerCreditLedgerNewEntryByExteralIDParams() {

}

type CustomerCreditLedgerNewEntryByExteralIDParamsAddIncrementCreditLedgerEntryRequestParamsEntryType string

const (
	CustomerCreditLedgerNewEntryByExteralIDParamsAddIncrementCreditLedgerEntryRequestParamsEntryTypeIncrement CustomerCreditLedgerNewEntryByExteralIDParamsAddIncrementCreditLedgerEntryRequestParamsEntryType = "increment"
)

// Passing `invoice_settings` automatically generates an invoice for the newly
// added credits. If `invoice_settings` is passed, you must specify
// per_unit_cost_basis, as the calculation of the invoice total is done on that
// basis.
type CustomerCreditLedgerNewEntryByExteralIDParamsAddIncrementCreditLedgerEntryRequestParamsInvoiceSettings struct {
	// Whether the credits purchase invoice should auto collect with the customer's
	// saved payment method.
	AutoCollection param.Field[bool] `json:"auto_collection,required"`
	// The net terms determines the difference between the invoice date and the issue
	// date for the invoice. If you intend the invoice to be due on issue, set this
	// to 0.
	NetTerms param.Field[int64] `json:"net_terms,required"`
	// An optional memo to display on the invoice.
	Memo param.Field[string] `json:"memo"`
}

func (r CustomerCreditLedgerNewEntryByExteralIDParamsAddIncrementCreditLedgerEntryRequestParamsInvoiceSettings) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type CustomerCreditLedgerNewEntryByExteralIDParamsAddDecrementCreditLedgerEntryRequestParams struct {
	// The number of credits to effect. Note that this is required for increment,
	// decrement, void, or undo operations.
	Amount    param.Field[float64]                                                                                          `json:"amount,required"`
	EntryType param.Field[CustomerCreditLedgerNewEntryByExteralIDParamsAddDecrementCreditLedgerEntryRequestParamsEntryType] `json:"entry_type,required"`
	// Optional metadata that can be specified when adding ledger results via the API.
	// For example, this can be used to note an increment refers to trial credits, or
	// for noting corrections as a result of an incident, etc.
	Description param.Field[string] `json:"description"`
	// User-specified key/value pairs for the ledger entry resource.
	Metadata param.Field[interface{}] `json:"metadata"`
}

func (r CustomerCreditLedgerNewEntryByExteralIDParamsAddDecrementCreditLedgerEntryRequestParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (CustomerCreditLedgerNewEntryByExteralIDParamsAddDecrementCreditLedgerEntryRequestParams) ImplementsCustomerCreditLedgerNewEntryByExteralIDParams() {

}

type CustomerCreditLedgerNewEntryByExteralIDParamsAddDecrementCreditLedgerEntryRequestParamsEntryType string

const (
	CustomerCreditLedgerNewEntryByExteralIDParamsAddDecrementCreditLedgerEntryRequestParamsEntryTypeDecrement CustomerCreditLedgerNewEntryByExteralIDParamsAddDecrementCreditLedgerEntryRequestParamsEntryType = "decrement"
)

type CustomerCreditLedgerNewEntryByExteralIDParamsAddExpirationChangeCreditLedgerEntryRequestParams struct {
	EntryType param.Field[CustomerCreditLedgerNewEntryByExteralIDParamsAddExpirationChangeCreditLedgerEntryRequestParamsEntryType] `json:"entry_type,required"`
	// A future date (specified in YYYY-MM-DD format) that identifies the origination
	// credit block to expire
	ExpiryDate param.Field[time.Time] `json:"expiry_date,required" format:"date"`
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
	// Optional metadata that can be specified when adding ledger results via the API.
	// For example, this can be used to note an increment refers to trial credits, or
	// for noting corrections as a result of an incident, etc.
	Description param.Field[string] `json:"description"`
	// User-specified key/value pairs for the ledger entry resource.
	Metadata param.Field[interface{}] `json:"metadata"`
}

func (r CustomerCreditLedgerNewEntryByExteralIDParamsAddExpirationChangeCreditLedgerEntryRequestParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (CustomerCreditLedgerNewEntryByExteralIDParamsAddExpirationChangeCreditLedgerEntryRequestParams) ImplementsCustomerCreditLedgerNewEntryByExteralIDParams() {

}

type CustomerCreditLedgerNewEntryByExteralIDParamsAddExpirationChangeCreditLedgerEntryRequestParamsEntryType string

const (
	CustomerCreditLedgerNewEntryByExteralIDParamsAddExpirationChangeCreditLedgerEntryRequestParamsEntryTypeExpirationChange CustomerCreditLedgerNewEntryByExteralIDParamsAddExpirationChangeCreditLedgerEntryRequestParamsEntryType = "expiration_change"
)

type CustomerCreditLedgerNewEntryByExteralIDParamsAddVoidCreditLedgerEntryRequestParams struct {
	// The number of credits to effect. Note that this is required for increment,
	// decrement, void, or undo operations.
	Amount param.Field[float64] `json:"amount,required"`
	// The ID of the block to void.
	BlockID   param.Field[string]                                                                                      `json:"block_id,required"`
	EntryType param.Field[CustomerCreditLedgerNewEntryByExteralIDParamsAddVoidCreditLedgerEntryRequestParamsEntryType] `json:"entry_type,required"`
	// Optional metadata that can be specified when adding ledger results via the API.
	// For example, this can be used to note an increment refers to trial credits, or
	// for noting corrections as a result of an incident, etc.
	Description param.Field[string] `json:"description"`
	// User-specified key/value pairs for the ledger entry resource.
	Metadata param.Field[interface{}] `json:"metadata"`
	// Can only be specified when `entry_type=void`. The reason for the void.
	VoidReason param.Field[CustomerCreditLedgerNewEntryByExteralIDParamsAddVoidCreditLedgerEntryRequestParamsVoidReason] `json:"void_reason"`
}

func (r CustomerCreditLedgerNewEntryByExteralIDParamsAddVoidCreditLedgerEntryRequestParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (CustomerCreditLedgerNewEntryByExteralIDParamsAddVoidCreditLedgerEntryRequestParams) ImplementsCustomerCreditLedgerNewEntryByExteralIDParams() {

}

type CustomerCreditLedgerNewEntryByExteralIDParamsAddVoidCreditLedgerEntryRequestParamsEntryType string

const (
	CustomerCreditLedgerNewEntryByExteralIDParamsAddVoidCreditLedgerEntryRequestParamsEntryTypeVoid CustomerCreditLedgerNewEntryByExteralIDParamsAddVoidCreditLedgerEntryRequestParamsEntryType = "void"
)

// Can only be specified when `entry_type=void`. The reason for the void.
type CustomerCreditLedgerNewEntryByExteralIDParamsAddVoidCreditLedgerEntryRequestParamsVoidReason string

const (
	CustomerCreditLedgerNewEntryByExteralIDParamsAddVoidCreditLedgerEntryRequestParamsVoidReasonRefund CustomerCreditLedgerNewEntryByExteralIDParamsAddVoidCreditLedgerEntryRequestParamsVoidReason = "refund"
)

type CustomerCreditLedgerNewEntryByExteralIDParamsAddAmendmentCreditLedgerEntryRequestParams struct {
	// The number of credits to effect. Note that this is required for increment,
	// decrement or void operations.
	Amount param.Field[float64] `json:"amount,required"`
	// The ID of the block to reverse a decrement from.
	BlockID   param.Field[string]                                                                                           `json:"block_id,required"`
	EntryType param.Field[CustomerCreditLedgerNewEntryByExteralIDParamsAddAmendmentCreditLedgerEntryRequestParamsEntryType] `json:"entry_type,required"`
	// Optional metadata that can be specified when adding ledger results via the API.
	// For example, this can be used to note an increment refers to trial credits, or
	// for noting corrections as a result of an incident, etc.
	Description param.Field[string] `json:"description"`
	// User-specified key/value pairs for the ledger entry resource.
	Metadata param.Field[interface{}] `json:"metadata"`
}

func (r CustomerCreditLedgerNewEntryByExteralIDParamsAddAmendmentCreditLedgerEntryRequestParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (CustomerCreditLedgerNewEntryByExteralIDParamsAddAmendmentCreditLedgerEntryRequestParams) ImplementsCustomerCreditLedgerNewEntryByExteralIDParams() {

}

type CustomerCreditLedgerNewEntryByExteralIDParamsAddAmendmentCreditLedgerEntryRequestParamsEntryType string

const (
	CustomerCreditLedgerNewEntryByExteralIDParamsAddAmendmentCreditLedgerEntryRequestParamsEntryTypeAmendment CustomerCreditLedgerNewEntryByExteralIDParamsAddAmendmentCreditLedgerEntryRequestParamsEntryType = "amendment"
)

type CustomerCreditLedgerListByExternalIDParams struct {
	CreatedAtGt  param.Field[time.Time] `query:"created_at[gt]" format:"date-time"`
	CreatedAtGte param.Field[time.Time] `query:"created_at[gte]" format:"date-time"`
	CreatedAtLt  param.Field[time.Time] `query:"created_at[lt]" format:"date-time"`
	CreatedAtLte param.Field[time.Time] `query:"created_at[lte]" format:"date-time"`
	Currency     param.Field[string]    `query:"currency"`
	// Cursor for pagination. This can be populated by the `next_cursor` value returned
	// from the initial request.
	Cursor      param.Field[string]                                                `query:"cursor"`
	EntryStatus param.Field[CustomerCreditLedgerListByExternalIDParamsEntryStatus] `query:"entry_status"`
	EntryType   param.Field[CustomerCreditLedgerListByExternalIDParamsEntryType]   `query:"entry_type"`
	// The number of items to fetch. Defaults to 20.
	Limit         param.Field[int64]  `query:"limit"`
	MinimumAmount param.Field[string] `query:"minimum_amount"`
}

// URLQuery serializes [CustomerCreditLedgerListByExternalIDParams]'s query
// parameters as `url.Values`.
func (r CustomerCreditLedgerListByExternalIDParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type CustomerCreditLedgerListByExternalIDParamsEntryStatus string

const (
	CustomerCreditLedgerListByExternalIDParamsEntryStatusCommitted CustomerCreditLedgerListByExternalIDParamsEntryStatus = "committed"
	CustomerCreditLedgerListByExternalIDParamsEntryStatusPending   CustomerCreditLedgerListByExternalIDParamsEntryStatus = "pending"
)

type CustomerCreditLedgerListByExternalIDParamsEntryType string

const (
	CustomerCreditLedgerListByExternalIDParamsEntryTypeIncrement         CustomerCreditLedgerListByExternalIDParamsEntryType = "increment"
	CustomerCreditLedgerListByExternalIDParamsEntryTypeDecrement         CustomerCreditLedgerListByExternalIDParamsEntryType = "decrement"
	CustomerCreditLedgerListByExternalIDParamsEntryTypeExpirationChange  CustomerCreditLedgerListByExternalIDParamsEntryType = "expiration_change"
	CustomerCreditLedgerListByExternalIDParamsEntryTypeCreditBlockExpiry CustomerCreditLedgerListByExternalIDParamsEntryType = "credit_block_expiry"
	CustomerCreditLedgerListByExternalIDParamsEntryTypeVoid              CustomerCreditLedgerListByExternalIDParamsEntryType = "void"
	CustomerCreditLedgerListByExternalIDParamsEntryTypeVoidInitiated     CustomerCreditLedgerListByExternalIDParamsEntryType = "void_initiated"
	CustomerCreditLedgerListByExternalIDParamsEntryTypeAmendment         CustomerCreditLedgerListByExternalIDParamsEntryType = "amendment"
)
