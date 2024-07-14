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
	"github.com/orbcorp/orb-go/internal/pagination"
	"github.com/orbcorp/orb-go/internal/param"
	"github.com/orbcorp/orb-go/internal/requestconfig"
	"github.com/orbcorp/orb-go/option"
	"github.com/tidwall/gjson"
)

// CustomerCreditLedgerService contains methods and other services that help with
// interacting with the orb API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewCustomerCreditLedgerService] method instead.
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
// the block with the _lower cost basis_ first (e.g. trial credits with a $0 cost
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
func (r *CustomerCreditLedgerService) List(ctx context.Context, customerID string, query CustomerCreditLedgerListParams, opts ...option.RequestOption) (res *pagination.Page[CustomerCreditLedgerListResponse], err error) {
	var raw *http.Response
	opts = append(r.Options[:], opts...)
	opts = append([]option.RequestOption{option.WithResponseInto(&raw)}, opts...)
	if customerID == "" {
		err = errors.New("missing required customer_id parameter")
		return
	}
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
// the block with the _lower cost basis_ first (e.g. trial credits with a $0 cost
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
func (r *CustomerCreditLedgerService) ListAutoPaging(ctx context.Context, customerID string, query CustomerCreditLedgerListParams, opts ...option.RequestOption) *pagination.PageAutoPager[CustomerCreditLedgerListResponse] {
	return pagination.NewPageAutoPager(r.List(ctx, customerID, query, opts...))
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
	if customerID == "" {
		err = errors.New("missing required customer_id parameter")
		return
	}
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
func (r *CustomerCreditLedgerService) NewEntryByExternalID(ctx context.Context, externalCustomerID string, body CustomerCreditLedgerNewEntryByExternalIDParams, opts ...option.RequestOption) (res *CustomerCreditLedgerNewEntryByExternalIDResponse, err error) {
	opts = append(r.Options[:], opts...)
	if externalCustomerID == "" {
		err = errors.New("missing required external_customer_id parameter")
		return
	}
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
// the block with the _lower cost basis_ first (e.g. trial credits with a $0 cost
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
func (r *CustomerCreditLedgerService) ListByExternalID(ctx context.Context, externalCustomerID string, query CustomerCreditLedgerListByExternalIDParams, opts ...option.RequestOption) (res *pagination.Page[CustomerCreditLedgerListByExternalIDResponse], err error) {
	var raw *http.Response
	opts = append(r.Options[:], opts...)
	opts = append([]option.RequestOption{option.WithResponseInto(&raw)}, opts...)
	if externalCustomerID == "" {
		err = errors.New("missing required external_customer_id parameter")
		return
	}
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
// the block with the _lower cost basis_ first (e.g. trial credits with a $0 cost
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
func (r *CustomerCreditLedgerService) ListByExternalIDAutoPaging(ctx context.Context, externalCustomerID string, query CustomerCreditLedgerListByExternalIDParams, opts ...option.RequestOption) *pagination.PageAutoPager[CustomerCreditLedgerListByExternalIDResponse] {
	return pagination.NewPageAutoPager(r.ListByExternalID(ctx, externalCustomerID, query, opts...))
}

// The [Credit Ledger Entry resource](/guides/product-catalog/prepurchase) models
// prepaid credits within Orb.
type CustomerCreditLedgerListResponse struct {
	// This field can have the runtime type of [map[string]string].
	Metadata             interface{}                                 `json:"metadata"`
	ID                   string                                      `json:"id,required"`
	LedgerSequenceNumber int64                                       `json:"ledger_sequence_number,required"`
	EntryStatus          CustomerCreditLedgerListResponseEntryStatus `json:"entry_status,required"`
	// This field can have the runtime type of
	// [CustomerCreditLedgerListResponseIncrementLedgerEntryCustomer],
	// [CustomerCreditLedgerListResponseDecrementLedgerEntryCustomer],
	// [CustomerCreditLedgerListResponseExpirationChangeLedgerEntryCustomer],
	// [CustomerCreditLedgerListResponseCreditBlockExpiryLedgerEntryCustomer],
	// [CustomerCreditLedgerListResponseVoidLedgerEntryCustomer],
	// [CustomerCreditLedgerListResponseVoidInitiatedLedgerEntryCustomer],
	// [CustomerCreditLedgerListResponseAmendmentLedgerEntryCustomer].
	Customer        interface{} `json:"customer"`
	StartingBalance float64     `json:"starting_balance,required"`
	EndingBalance   float64     `json:"ending_balance,required"`
	Amount          float64     `json:"amount,required"`
	Currency        string      `json:"currency,required"`
	CreatedAt       time.Time   `json:"created_at,required" format:"date-time"`
	Description     string      `json:"description,required,nullable"`
	// This field can have the runtime type of
	// [CustomerCreditLedgerListResponseIncrementLedgerEntryCreditBlock],
	// [CustomerCreditLedgerListResponseDecrementLedgerEntryCreditBlock],
	// [CustomerCreditLedgerListResponseExpirationChangeLedgerEntryCreditBlock],
	// [CustomerCreditLedgerListResponseCreditBlockExpiryLedgerEntryCreditBlock],
	// [CustomerCreditLedgerListResponseVoidLedgerEntryCreditBlock],
	// [CustomerCreditLedgerListResponseVoidInitiatedLedgerEntryCreditBlock],
	// [CustomerCreditLedgerListResponseAmendmentLedgerEntryCreditBlock].
	CreditBlock        interface{}                               `json:"credit_block"`
	EntryType          CustomerCreditLedgerListResponseEntryType `json:"entry_type,required"`
	PriceID            string                                    `json:"price_id,nullable"`
	EventID            string                                    `json:"event_id,nullable"`
	InvoiceID          string                                    `json:"invoice_id,nullable"`
	NewBlockExpiryDate time.Time                                 `json:"new_block_expiry_date" format:"date-time"`
	VoidReason         string                                    `json:"void_reason,nullable"`
	VoidAmount         float64                                   `json:"void_amount"`
	JSON               customerCreditLedgerListResponseJSON      `json:"-"`
	union              CustomerCreditLedgerListResponseUnion
}

// customerCreditLedgerListResponseJSON contains the JSON metadata for the struct
// [CustomerCreditLedgerListResponse]
type customerCreditLedgerListResponseJSON struct {
	Metadata             apijson.Field
	ID                   apijson.Field
	LedgerSequenceNumber apijson.Field
	EntryStatus          apijson.Field
	Customer             apijson.Field
	StartingBalance      apijson.Field
	EndingBalance        apijson.Field
	Amount               apijson.Field
	Currency             apijson.Field
	CreatedAt            apijson.Field
	Description          apijson.Field
	CreditBlock          apijson.Field
	EntryType            apijson.Field
	PriceID              apijson.Field
	EventID              apijson.Field
	InvoiceID            apijson.Field
	NewBlockExpiryDate   apijson.Field
	VoidReason           apijson.Field
	VoidAmount           apijson.Field
	raw                  string
	ExtraFields          map[string]apijson.Field
}

func (r customerCreditLedgerListResponseJSON) RawJSON() string {
	return r.raw
}

func (r *CustomerCreditLedgerListResponse) UnmarshalJSON(data []byte) (err error) {
	*r = CustomerCreditLedgerListResponse{}
	err = apijson.UnmarshalRoot(data, &r.union)
	if err != nil {
		return err
	}
	return apijson.Port(r.union, &r)
}

// AsUnion returns a [CustomerCreditLedgerListResponseUnion] interface which you
// can cast to the specific types for more type safety.
//
// Possible runtime types of the union are
// [CustomerCreditLedgerListResponseIncrementLedgerEntry],
// [CustomerCreditLedgerListResponseDecrementLedgerEntry],
// [CustomerCreditLedgerListResponseExpirationChangeLedgerEntry],
// [CustomerCreditLedgerListResponseCreditBlockExpiryLedgerEntry],
// [CustomerCreditLedgerListResponseVoidLedgerEntry],
// [CustomerCreditLedgerListResponseVoidInitiatedLedgerEntry],
// [CustomerCreditLedgerListResponseAmendmentLedgerEntry].
func (r CustomerCreditLedgerListResponse) AsUnion() CustomerCreditLedgerListResponseUnion {
	return r.union
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
type CustomerCreditLedgerListResponseUnion interface {
	implementsCustomerCreditLedgerListResponse()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*CustomerCreditLedgerListResponseUnion)(nil)).Elem(),
		"entry_type",
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(CustomerCreditLedgerListResponseIncrementLedgerEntry{}),
			DiscriminatorValue: "increment",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(CustomerCreditLedgerListResponseDecrementLedgerEntry{}),
			DiscriminatorValue: "decrement",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(CustomerCreditLedgerListResponseExpirationChangeLedgerEntry{}),
			DiscriminatorValue: "expiration_change",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(CustomerCreditLedgerListResponseCreditBlockExpiryLedgerEntry{}),
			DiscriminatorValue: "credit_block_expiry",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(CustomerCreditLedgerListResponseVoidLedgerEntry{}),
			DiscriminatorValue: "void",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(CustomerCreditLedgerListResponseVoidInitiatedLedgerEntry{}),
			DiscriminatorValue: "void_initiated",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(CustomerCreditLedgerListResponseAmendmentLedgerEntry{}),
			DiscriminatorValue: "amendment",
		},
	)
}

type CustomerCreditLedgerListResponseIncrementLedgerEntry struct {
	ID                   string                                                          `json:"id,required"`
	Amount               float64                                                         `json:"amount,required"`
	CreatedAt            time.Time                                                       `json:"created_at,required" format:"date-time"`
	CreditBlock          CustomerCreditLedgerListResponseIncrementLedgerEntryCreditBlock `json:"credit_block,required"`
	Currency             string                                                          `json:"currency,required"`
	Customer             CustomerCreditLedgerListResponseIncrementLedgerEntryCustomer    `json:"customer,required"`
	Description          string                                                          `json:"description,required,nullable"`
	EndingBalance        float64                                                         `json:"ending_balance,required"`
	EntryStatus          CustomerCreditLedgerListResponseIncrementLedgerEntryEntryStatus `json:"entry_status,required"`
	EntryType            CustomerCreditLedgerListResponseIncrementLedgerEntryEntryType   `json:"entry_type,required"`
	LedgerSequenceNumber int64                                                           `json:"ledger_sequence_number,required"`
	// User specified key-value pairs for the resource. If not present, this defaults
	// to an empty dictionary. Individual keys can be removed by setting the value to
	// `null`, and the entire metadata mapping can be cleared by setting `metadata` to
	// `null`.
	Metadata        map[string]string                                        `json:"metadata,required"`
	StartingBalance float64                                                  `json:"starting_balance,required"`
	JSON            customerCreditLedgerListResponseIncrementLedgerEntryJSON `json:"-"`
}

// customerCreditLedgerListResponseIncrementLedgerEntryJSON contains the JSON
// metadata for the struct [CustomerCreditLedgerListResponseIncrementLedgerEntry]
type customerCreditLedgerListResponseIncrementLedgerEntryJSON struct {
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

func (r *CustomerCreditLedgerListResponseIncrementLedgerEntry) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r customerCreditLedgerListResponseIncrementLedgerEntryJSON) RawJSON() string {
	return r.raw
}

func (r CustomerCreditLedgerListResponseIncrementLedgerEntry) implementsCustomerCreditLedgerListResponse() {
}

type CustomerCreditLedgerListResponseIncrementLedgerEntryCreditBlock struct {
	ID               string                                                              `json:"id,required"`
	ExpiryDate       time.Time                                                           `json:"expiry_date,required,nullable" format:"date-time"`
	PerUnitCostBasis string                                                              `json:"per_unit_cost_basis,required,nullable"`
	JSON             customerCreditLedgerListResponseIncrementLedgerEntryCreditBlockJSON `json:"-"`
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

func (r customerCreditLedgerListResponseIncrementLedgerEntryCreditBlockJSON) RawJSON() string {
	return r.raw
}

type CustomerCreditLedgerListResponseIncrementLedgerEntryCustomer struct {
	ID                 string                                                           `json:"id,required"`
	ExternalCustomerID string                                                           `json:"external_customer_id,required,nullable"`
	JSON               customerCreditLedgerListResponseIncrementLedgerEntryCustomerJSON `json:"-"`
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

func (r customerCreditLedgerListResponseIncrementLedgerEntryCustomerJSON) RawJSON() string {
	return r.raw
}

type CustomerCreditLedgerListResponseIncrementLedgerEntryEntryStatus string

const (
	CustomerCreditLedgerListResponseIncrementLedgerEntryEntryStatusCommitted CustomerCreditLedgerListResponseIncrementLedgerEntryEntryStatus = "committed"
	CustomerCreditLedgerListResponseIncrementLedgerEntryEntryStatusPending   CustomerCreditLedgerListResponseIncrementLedgerEntryEntryStatus = "pending"
)

func (r CustomerCreditLedgerListResponseIncrementLedgerEntryEntryStatus) IsKnown() bool {
	switch r {
	case CustomerCreditLedgerListResponseIncrementLedgerEntryEntryStatusCommitted, CustomerCreditLedgerListResponseIncrementLedgerEntryEntryStatusPending:
		return true
	}
	return false
}

type CustomerCreditLedgerListResponseIncrementLedgerEntryEntryType string

const (
	CustomerCreditLedgerListResponseIncrementLedgerEntryEntryTypeIncrement CustomerCreditLedgerListResponseIncrementLedgerEntryEntryType = "increment"
)

func (r CustomerCreditLedgerListResponseIncrementLedgerEntryEntryType) IsKnown() bool {
	switch r {
	case CustomerCreditLedgerListResponseIncrementLedgerEntryEntryTypeIncrement:
		return true
	}
	return false
}

type CustomerCreditLedgerListResponseDecrementLedgerEntry struct {
	ID                   string                                                          `json:"id,required"`
	Amount               float64                                                         `json:"amount,required"`
	CreatedAt            time.Time                                                       `json:"created_at,required" format:"date-time"`
	CreditBlock          CustomerCreditLedgerListResponseDecrementLedgerEntryCreditBlock `json:"credit_block,required"`
	Currency             string                                                          `json:"currency,required"`
	Customer             CustomerCreditLedgerListResponseDecrementLedgerEntryCustomer    `json:"customer,required"`
	Description          string                                                          `json:"description,required,nullable"`
	EndingBalance        float64                                                         `json:"ending_balance,required"`
	EntryStatus          CustomerCreditLedgerListResponseDecrementLedgerEntryEntryStatus `json:"entry_status,required"`
	EntryType            CustomerCreditLedgerListResponseDecrementLedgerEntryEntryType   `json:"entry_type,required"`
	LedgerSequenceNumber int64                                                           `json:"ledger_sequence_number,required"`
	// User specified key-value pairs for the resource. If not present, this defaults
	// to an empty dictionary. Individual keys can be removed by setting the value to
	// `null`, and the entire metadata mapping can be cleared by setting `metadata` to
	// `null`.
	Metadata        map[string]string                                        `json:"metadata,required"`
	StartingBalance float64                                                  `json:"starting_balance,required"`
	EventID         string                                                   `json:"event_id,nullable"`
	InvoiceID       string                                                   `json:"invoice_id,nullable"`
	PriceID         string                                                   `json:"price_id,nullable"`
	JSON            customerCreditLedgerListResponseDecrementLedgerEntryJSON `json:"-"`
}

// customerCreditLedgerListResponseDecrementLedgerEntryJSON contains the JSON
// metadata for the struct [CustomerCreditLedgerListResponseDecrementLedgerEntry]
type customerCreditLedgerListResponseDecrementLedgerEntryJSON struct {
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

func (r *CustomerCreditLedgerListResponseDecrementLedgerEntry) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r customerCreditLedgerListResponseDecrementLedgerEntryJSON) RawJSON() string {
	return r.raw
}

func (r CustomerCreditLedgerListResponseDecrementLedgerEntry) implementsCustomerCreditLedgerListResponse() {
}

type CustomerCreditLedgerListResponseDecrementLedgerEntryCreditBlock struct {
	ID               string                                                              `json:"id,required"`
	ExpiryDate       time.Time                                                           `json:"expiry_date,required,nullable" format:"date-time"`
	PerUnitCostBasis string                                                              `json:"per_unit_cost_basis,required,nullable"`
	JSON             customerCreditLedgerListResponseDecrementLedgerEntryCreditBlockJSON `json:"-"`
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

func (r customerCreditLedgerListResponseDecrementLedgerEntryCreditBlockJSON) RawJSON() string {
	return r.raw
}

type CustomerCreditLedgerListResponseDecrementLedgerEntryCustomer struct {
	ID                 string                                                           `json:"id,required"`
	ExternalCustomerID string                                                           `json:"external_customer_id,required,nullable"`
	JSON               customerCreditLedgerListResponseDecrementLedgerEntryCustomerJSON `json:"-"`
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

func (r customerCreditLedgerListResponseDecrementLedgerEntryCustomerJSON) RawJSON() string {
	return r.raw
}

type CustomerCreditLedgerListResponseDecrementLedgerEntryEntryStatus string

const (
	CustomerCreditLedgerListResponseDecrementLedgerEntryEntryStatusCommitted CustomerCreditLedgerListResponseDecrementLedgerEntryEntryStatus = "committed"
	CustomerCreditLedgerListResponseDecrementLedgerEntryEntryStatusPending   CustomerCreditLedgerListResponseDecrementLedgerEntryEntryStatus = "pending"
)

func (r CustomerCreditLedgerListResponseDecrementLedgerEntryEntryStatus) IsKnown() bool {
	switch r {
	case CustomerCreditLedgerListResponseDecrementLedgerEntryEntryStatusCommitted, CustomerCreditLedgerListResponseDecrementLedgerEntryEntryStatusPending:
		return true
	}
	return false
}

type CustomerCreditLedgerListResponseDecrementLedgerEntryEntryType string

const (
	CustomerCreditLedgerListResponseDecrementLedgerEntryEntryTypeDecrement CustomerCreditLedgerListResponseDecrementLedgerEntryEntryType = "decrement"
)

func (r CustomerCreditLedgerListResponseDecrementLedgerEntryEntryType) IsKnown() bool {
	switch r {
	case CustomerCreditLedgerListResponseDecrementLedgerEntryEntryTypeDecrement:
		return true
	}
	return false
}

type CustomerCreditLedgerListResponseExpirationChangeLedgerEntry struct {
	ID                   string                                                                 `json:"id,required"`
	Amount               float64                                                                `json:"amount,required"`
	CreatedAt            time.Time                                                              `json:"created_at,required" format:"date-time"`
	CreditBlock          CustomerCreditLedgerListResponseExpirationChangeLedgerEntryCreditBlock `json:"credit_block,required"`
	Currency             string                                                                 `json:"currency,required"`
	Customer             CustomerCreditLedgerListResponseExpirationChangeLedgerEntryCustomer    `json:"customer,required"`
	Description          string                                                                 `json:"description,required,nullable"`
	EndingBalance        float64                                                                `json:"ending_balance,required"`
	EntryStatus          CustomerCreditLedgerListResponseExpirationChangeLedgerEntryEntryStatus `json:"entry_status,required"`
	EntryType            CustomerCreditLedgerListResponseExpirationChangeLedgerEntryEntryType   `json:"entry_type,required"`
	LedgerSequenceNumber int64                                                                  `json:"ledger_sequence_number,required"`
	// User specified key-value pairs for the resource. If not present, this defaults
	// to an empty dictionary. Individual keys can be removed by setting the value to
	// `null`, and the entire metadata mapping can be cleared by setting `metadata` to
	// `null`.
	Metadata           map[string]string                                               `json:"metadata,required"`
	NewBlockExpiryDate time.Time                                                       `json:"new_block_expiry_date,required" format:"date-time"`
	StartingBalance    float64                                                         `json:"starting_balance,required"`
	JSON               customerCreditLedgerListResponseExpirationChangeLedgerEntryJSON `json:"-"`
}

// customerCreditLedgerListResponseExpirationChangeLedgerEntryJSON contains the
// JSON metadata for the struct
// [CustomerCreditLedgerListResponseExpirationChangeLedgerEntry]
type customerCreditLedgerListResponseExpirationChangeLedgerEntryJSON struct {
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

func (r *CustomerCreditLedgerListResponseExpirationChangeLedgerEntry) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r customerCreditLedgerListResponseExpirationChangeLedgerEntryJSON) RawJSON() string {
	return r.raw
}

func (r CustomerCreditLedgerListResponseExpirationChangeLedgerEntry) implementsCustomerCreditLedgerListResponse() {
}

type CustomerCreditLedgerListResponseExpirationChangeLedgerEntryCreditBlock struct {
	ID               string                                                                     `json:"id,required"`
	ExpiryDate       time.Time                                                                  `json:"expiry_date,required,nullable" format:"date-time"`
	PerUnitCostBasis string                                                                     `json:"per_unit_cost_basis,required,nullable"`
	JSON             customerCreditLedgerListResponseExpirationChangeLedgerEntryCreditBlockJSON `json:"-"`
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

func (r customerCreditLedgerListResponseExpirationChangeLedgerEntryCreditBlockJSON) RawJSON() string {
	return r.raw
}

type CustomerCreditLedgerListResponseExpirationChangeLedgerEntryCustomer struct {
	ID                 string                                                                  `json:"id,required"`
	ExternalCustomerID string                                                                  `json:"external_customer_id,required,nullable"`
	JSON               customerCreditLedgerListResponseExpirationChangeLedgerEntryCustomerJSON `json:"-"`
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

func (r customerCreditLedgerListResponseExpirationChangeLedgerEntryCustomerJSON) RawJSON() string {
	return r.raw
}

type CustomerCreditLedgerListResponseExpirationChangeLedgerEntryEntryStatus string

const (
	CustomerCreditLedgerListResponseExpirationChangeLedgerEntryEntryStatusCommitted CustomerCreditLedgerListResponseExpirationChangeLedgerEntryEntryStatus = "committed"
	CustomerCreditLedgerListResponseExpirationChangeLedgerEntryEntryStatusPending   CustomerCreditLedgerListResponseExpirationChangeLedgerEntryEntryStatus = "pending"
)

func (r CustomerCreditLedgerListResponseExpirationChangeLedgerEntryEntryStatus) IsKnown() bool {
	switch r {
	case CustomerCreditLedgerListResponseExpirationChangeLedgerEntryEntryStatusCommitted, CustomerCreditLedgerListResponseExpirationChangeLedgerEntryEntryStatusPending:
		return true
	}
	return false
}

type CustomerCreditLedgerListResponseExpirationChangeLedgerEntryEntryType string

const (
	CustomerCreditLedgerListResponseExpirationChangeLedgerEntryEntryTypeExpirationChange CustomerCreditLedgerListResponseExpirationChangeLedgerEntryEntryType = "expiration_change"
)

func (r CustomerCreditLedgerListResponseExpirationChangeLedgerEntryEntryType) IsKnown() bool {
	switch r {
	case CustomerCreditLedgerListResponseExpirationChangeLedgerEntryEntryTypeExpirationChange:
		return true
	}
	return false
}

type CustomerCreditLedgerListResponseCreditBlockExpiryLedgerEntry struct {
	ID                   string                                                                  `json:"id,required"`
	Amount               float64                                                                 `json:"amount,required"`
	CreatedAt            time.Time                                                               `json:"created_at,required" format:"date-time"`
	CreditBlock          CustomerCreditLedgerListResponseCreditBlockExpiryLedgerEntryCreditBlock `json:"credit_block,required"`
	Currency             string                                                                  `json:"currency,required"`
	Customer             CustomerCreditLedgerListResponseCreditBlockExpiryLedgerEntryCustomer    `json:"customer,required"`
	Description          string                                                                  `json:"description,required,nullable"`
	EndingBalance        float64                                                                 `json:"ending_balance,required"`
	EntryStatus          CustomerCreditLedgerListResponseCreditBlockExpiryLedgerEntryEntryStatus `json:"entry_status,required"`
	EntryType            CustomerCreditLedgerListResponseCreditBlockExpiryLedgerEntryEntryType   `json:"entry_type,required"`
	LedgerSequenceNumber int64                                                                   `json:"ledger_sequence_number,required"`
	// User specified key-value pairs for the resource. If not present, this defaults
	// to an empty dictionary. Individual keys can be removed by setting the value to
	// `null`, and the entire metadata mapping can be cleared by setting `metadata` to
	// `null`.
	Metadata        map[string]string                                                `json:"metadata,required"`
	StartingBalance float64                                                          `json:"starting_balance,required"`
	JSON            customerCreditLedgerListResponseCreditBlockExpiryLedgerEntryJSON `json:"-"`
}

// customerCreditLedgerListResponseCreditBlockExpiryLedgerEntryJSON contains the
// JSON metadata for the struct
// [CustomerCreditLedgerListResponseCreditBlockExpiryLedgerEntry]
type customerCreditLedgerListResponseCreditBlockExpiryLedgerEntryJSON struct {
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

func (r *CustomerCreditLedgerListResponseCreditBlockExpiryLedgerEntry) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r customerCreditLedgerListResponseCreditBlockExpiryLedgerEntryJSON) RawJSON() string {
	return r.raw
}

func (r CustomerCreditLedgerListResponseCreditBlockExpiryLedgerEntry) implementsCustomerCreditLedgerListResponse() {
}

type CustomerCreditLedgerListResponseCreditBlockExpiryLedgerEntryCreditBlock struct {
	ID               string                                                                      `json:"id,required"`
	ExpiryDate       time.Time                                                                   `json:"expiry_date,required,nullable" format:"date-time"`
	PerUnitCostBasis string                                                                      `json:"per_unit_cost_basis,required,nullable"`
	JSON             customerCreditLedgerListResponseCreditBlockExpiryLedgerEntryCreditBlockJSON `json:"-"`
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

func (r customerCreditLedgerListResponseCreditBlockExpiryLedgerEntryCreditBlockJSON) RawJSON() string {
	return r.raw
}

type CustomerCreditLedgerListResponseCreditBlockExpiryLedgerEntryCustomer struct {
	ID                 string                                                                   `json:"id,required"`
	ExternalCustomerID string                                                                   `json:"external_customer_id,required,nullable"`
	JSON               customerCreditLedgerListResponseCreditBlockExpiryLedgerEntryCustomerJSON `json:"-"`
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

func (r customerCreditLedgerListResponseCreditBlockExpiryLedgerEntryCustomerJSON) RawJSON() string {
	return r.raw
}

type CustomerCreditLedgerListResponseCreditBlockExpiryLedgerEntryEntryStatus string

const (
	CustomerCreditLedgerListResponseCreditBlockExpiryLedgerEntryEntryStatusCommitted CustomerCreditLedgerListResponseCreditBlockExpiryLedgerEntryEntryStatus = "committed"
	CustomerCreditLedgerListResponseCreditBlockExpiryLedgerEntryEntryStatusPending   CustomerCreditLedgerListResponseCreditBlockExpiryLedgerEntryEntryStatus = "pending"
)

func (r CustomerCreditLedgerListResponseCreditBlockExpiryLedgerEntryEntryStatus) IsKnown() bool {
	switch r {
	case CustomerCreditLedgerListResponseCreditBlockExpiryLedgerEntryEntryStatusCommitted, CustomerCreditLedgerListResponseCreditBlockExpiryLedgerEntryEntryStatusPending:
		return true
	}
	return false
}

type CustomerCreditLedgerListResponseCreditBlockExpiryLedgerEntryEntryType string

const (
	CustomerCreditLedgerListResponseCreditBlockExpiryLedgerEntryEntryTypeCreditBlockExpiry CustomerCreditLedgerListResponseCreditBlockExpiryLedgerEntryEntryType = "credit_block_expiry"
)

func (r CustomerCreditLedgerListResponseCreditBlockExpiryLedgerEntryEntryType) IsKnown() bool {
	switch r {
	case CustomerCreditLedgerListResponseCreditBlockExpiryLedgerEntryEntryTypeCreditBlockExpiry:
		return true
	}
	return false
}

type CustomerCreditLedgerListResponseVoidLedgerEntry struct {
	ID                   string                                                     `json:"id,required"`
	Amount               float64                                                    `json:"amount,required"`
	CreatedAt            time.Time                                                  `json:"created_at,required" format:"date-time"`
	CreditBlock          CustomerCreditLedgerListResponseVoidLedgerEntryCreditBlock `json:"credit_block,required"`
	Currency             string                                                     `json:"currency,required"`
	Customer             CustomerCreditLedgerListResponseVoidLedgerEntryCustomer    `json:"customer,required"`
	Description          string                                                     `json:"description,required,nullable"`
	EndingBalance        float64                                                    `json:"ending_balance,required"`
	EntryStatus          CustomerCreditLedgerListResponseVoidLedgerEntryEntryStatus `json:"entry_status,required"`
	EntryType            CustomerCreditLedgerListResponseVoidLedgerEntryEntryType   `json:"entry_type,required"`
	LedgerSequenceNumber int64                                                      `json:"ledger_sequence_number,required"`
	// User specified key-value pairs for the resource. If not present, this defaults
	// to an empty dictionary. Individual keys can be removed by setting the value to
	// `null`, and the entire metadata mapping can be cleared by setting `metadata` to
	// `null`.
	Metadata        map[string]string                                   `json:"metadata,required"`
	StartingBalance float64                                             `json:"starting_balance,required"`
	VoidAmount      float64                                             `json:"void_amount,required"`
	VoidReason      string                                              `json:"void_reason,required,nullable"`
	JSON            customerCreditLedgerListResponseVoidLedgerEntryJSON `json:"-"`
}

// customerCreditLedgerListResponseVoidLedgerEntryJSON contains the JSON metadata
// for the struct [CustomerCreditLedgerListResponseVoidLedgerEntry]
type customerCreditLedgerListResponseVoidLedgerEntryJSON struct {
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

func (r *CustomerCreditLedgerListResponseVoidLedgerEntry) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r customerCreditLedgerListResponseVoidLedgerEntryJSON) RawJSON() string {
	return r.raw
}

func (r CustomerCreditLedgerListResponseVoidLedgerEntry) implementsCustomerCreditLedgerListResponse() {
}

type CustomerCreditLedgerListResponseVoidLedgerEntryCreditBlock struct {
	ID               string                                                         `json:"id,required"`
	ExpiryDate       time.Time                                                      `json:"expiry_date,required,nullable" format:"date-time"`
	PerUnitCostBasis string                                                         `json:"per_unit_cost_basis,required,nullable"`
	JSON             customerCreditLedgerListResponseVoidLedgerEntryCreditBlockJSON `json:"-"`
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

func (r customerCreditLedgerListResponseVoidLedgerEntryCreditBlockJSON) RawJSON() string {
	return r.raw
}

type CustomerCreditLedgerListResponseVoidLedgerEntryCustomer struct {
	ID                 string                                                      `json:"id,required"`
	ExternalCustomerID string                                                      `json:"external_customer_id,required,nullable"`
	JSON               customerCreditLedgerListResponseVoidLedgerEntryCustomerJSON `json:"-"`
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

func (r customerCreditLedgerListResponseVoidLedgerEntryCustomerJSON) RawJSON() string {
	return r.raw
}

type CustomerCreditLedgerListResponseVoidLedgerEntryEntryStatus string

const (
	CustomerCreditLedgerListResponseVoidLedgerEntryEntryStatusCommitted CustomerCreditLedgerListResponseVoidLedgerEntryEntryStatus = "committed"
	CustomerCreditLedgerListResponseVoidLedgerEntryEntryStatusPending   CustomerCreditLedgerListResponseVoidLedgerEntryEntryStatus = "pending"
)

func (r CustomerCreditLedgerListResponseVoidLedgerEntryEntryStatus) IsKnown() bool {
	switch r {
	case CustomerCreditLedgerListResponseVoidLedgerEntryEntryStatusCommitted, CustomerCreditLedgerListResponseVoidLedgerEntryEntryStatusPending:
		return true
	}
	return false
}

type CustomerCreditLedgerListResponseVoidLedgerEntryEntryType string

const (
	CustomerCreditLedgerListResponseVoidLedgerEntryEntryTypeVoid CustomerCreditLedgerListResponseVoidLedgerEntryEntryType = "void"
)

func (r CustomerCreditLedgerListResponseVoidLedgerEntryEntryType) IsKnown() bool {
	switch r {
	case CustomerCreditLedgerListResponseVoidLedgerEntryEntryTypeVoid:
		return true
	}
	return false
}

type CustomerCreditLedgerListResponseVoidInitiatedLedgerEntry struct {
	ID                   string                                                              `json:"id,required"`
	Amount               float64                                                             `json:"amount,required"`
	CreatedAt            time.Time                                                           `json:"created_at,required" format:"date-time"`
	CreditBlock          CustomerCreditLedgerListResponseVoidInitiatedLedgerEntryCreditBlock `json:"credit_block,required"`
	Currency             string                                                              `json:"currency,required"`
	Customer             CustomerCreditLedgerListResponseVoidInitiatedLedgerEntryCustomer    `json:"customer,required"`
	Description          string                                                              `json:"description,required,nullable"`
	EndingBalance        float64                                                             `json:"ending_balance,required"`
	EntryStatus          CustomerCreditLedgerListResponseVoidInitiatedLedgerEntryEntryStatus `json:"entry_status,required"`
	EntryType            CustomerCreditLedgerListResponseVoidInitiatedLedgerEntryEntryType   `json:"entry_type,required"`
	LedgerSequenceNumber int64                                                               `json:"ledger_sequence_number,required"`
	// User specified key-value pairs for the resource. If not present, this defaults
	// to an empty dictionary. Individual keys can be removed by setting the value to
	// `null`, and the entire metadata mapping can be cleared by setting `metadata` to
	// `null`.
	Metadata           map[string]string                                            `json:"metadata,required"`
	NewBlockExpiryDate time.Time                                                    `json:"new_block_expiry_date,required" format:"date-time"`
	StartingBalance    float64                                                      `json:"starting_balance,required"`
	VoidAmount         float64                                                      `json:"void_amount,required"`
	VoidReason         string                                                       `json:"void_reason,required,nullable"`
	JSON               customerCreditLedgerListResponseVoidInitiatedLedgerEntryJSON `json:"-"`
}

// customerCreditLedgerListResponseVoidInitiatedLedgerEntryJSON contains the JSON
// metadata for the struct
// [CustomerCreditLedgerListResponseVoidInitiatedLedgerEntry]
type customerCreditLedgerListResponseVoidInitiatedLedgerEntryJSON struct {
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

func (r *CustomerCreditLedgerListResponseVoidInitiatedLedgerEntry) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r customerCreditLedgerListResponseVoidInitiatedLedgerEntryJSON) RawJSON() string {
	return r.raw
}

func (r CustomerCreditLedgerListResponseVoidInitiatedLedgerEntry) implementsCustomerCreditLedgerListResponse() {
}

type CustomerCreditLedgerListResponseVoidInitiatedLedgerEntryCreditBlock struct {
	ID               string                                                                  `json:"id,required"`
	ExpiryDate       time.Time                                                               `json:"expiry_date,required,nullable" format:"date-time"`
	PerUnitCostBasis string                                                                  `json:"per_unit_cost_basis,required,nullable"`
	JSON             customerCreditLedgerListResponseVoidInitiatedLedgerEntryCreditBlockJSON `json:"-"`
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

func (r customerCreditLedgerListResponseVoidInitiatedLedgerEntryCreditBlockJSON) RawJSON() string {
	return r.raw
}

type CustomerCreditLedgerListResponseVoidInitiatedLedgerEntryCustomer struct {
	ID                 string                                                               `json:"id,required"`
	ExternalCustomerID string                                                               `json:"external_customer_id,required,nullable"`
	JSON               customerCreditLedgerListResponseVoidInitiatedLedgerEntryCustomerJSON `json:"-"`
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

func (r customerCreditLedgerListResponseVoidInitiatedLedgerEntryCustomerJSON) RawJSON() string {
	return r.raw
}

type CustomerCreditLedgerListResponseVoidInitiatedLedgerEntryEntryStatus string

const (
	CustomerCreditLedgerListResponseVoidInitiatedLedgerEntryEntryStatusCommitted CustomerCreditLedgerListResponseVoidInitiatedLedgerEntryEntryStatus = "committed"
	CustomerCreditLedgerListResponseVoidInitiatedLedgerEntryEntryStatusPending   CustomerCreditLedgerListResponseVoidInitiatedLedgerEntryEntryStatus = "pending"
)

func (r CustomerCreditLedgerListResponseVoidInitiatedLedgerEntryEntryStatus) IsKnown() bool {
	switch r {
	case CustomerCreditLedgerListResponseVoidInitiatedLedgerEntryEntryStatusCommitted, CustomerCreditLedgerListResponseVoidInitiatedLedgerEntryEntryStatusPending:
		return true
	}
	return false
}

type CustomerCreditLedgerListResponseVoidInitiatedLedgerEntryEntryType string

const (
	CustomerCreditLedgerListResponseVoidInitiatedLedgerEntryEntryTypeVoidInitiated CustomerCreditLedgerListResponseVoidInitiatedLedgerEntryEntryType = "void_initiated"
)

func (r CustomerCreditLedgerListResponseVoidInitiatedLedgerEntryEntryType) IsKnown() bool {
	switch r {
	case CustomerCreditLedgerListResponseVoidInitiatedLedgerEntryEntryTypeVoidInitiated:
		return true
	}
	return false
}

type CustomerCreditLedgerListResponseAmendmentLedgerEntry struct {
	ID                   string                                                          `json:"id,required"`
	Amount               float64                                                         `json:"amount,required"`
	CreatedAt            time.Time                                                       `json:"created_at,required" format:"date-time"`
	CreditBlock          CustomerCreditLedgerListResponseAmendmentLedgerEntryCreditBlock `json:"credit_block,required"`
	Currency             string                                                          `json:"currency,required"`
	Customer             CustomerCreditLedgerListResponseAmendmentLedgerEntryCustomer    `json:"customer,required"`
	Description          string                                                          `json:"description,required,nullable"`
	EndingBalance        float64                                                         `json:"ending_balance,required"`
	EntryStatus          CustomerCreditLedgerListResponseAmendmentLedgerEntryEntryStatus `json:"entry_status,required"`
	EntryType            CustomerCreditLedgerListResponseAmendmentLedgerEntryEntryType   `json:"entry_type,required"`
	LedgerSequenceNumber int64                                                           `json:"ledger_sequence_number,required"`
	// User specified key-value pairs for the resource. If not present, this defaults
	// to an empty dictionary. Individual keys can be removed by setting the value to
	// `null`, and the entire metadata mapping can be cleared by setting `metadata` to
	// `null`.
	Metadata        map[string]string                                        `json:"metadata,required"`
	StartingBalance float64                                                  `json:"starting_balance,required"`
	JSON            customerCreditLedgerListResponseAmendmentLedgerEntryJSON `json:"-"`
}

// customerCreditLedgerListResponseAmendmentLedgerEntryJSON contains the JSON
// metadata for the struct [CustomerCreditLedgerListResponseAmendmentLedgerEntry]
type customerCreditLedgerListResponseAmendmentLedgerEntryJSON struct {
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

func (r *CustomerCreditLedgerListResponseAmendmentLedgerEntry) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r customerCreditLedgerListResponseAmendmentLedgerEntryJSON) RawJSON() string {
	return r.raw
}

func (r CustomerCreditLedgerListResponseAmendmentLedgerEntry) implementsCustomerCreditLedgerListResponse() {
}

type CustomerCreditLedgerListResponseAmendmentLedgerEntryCreditBlock struct {
	ID               string                                                              `json:"id,required"`
	ExpiryDate       time.Time                                                           `json:"expiry_date,required,nullable" format:"date-time"`
	PerUnitCostBasis string                                                              `json:"per_unit_cost_basis,required,nullable"`
	JSON             customerCreditLedgerListResponseAmendmentLedgerEntryCreditBlockJSON `json:"-"`
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

func (r customerCreditLedgerListResponseAmendmentLedgerEntryCreditBlockJSON) RawJSON() string {
	return r.raw
}

type CustomerCreditLedgerListResponseAmendmentLedgerEntryCustomer struct {
	ID                 string                                                           `json:"id,required"`
	ExternalCustomerID string                                                           `json:"external_customer_id,required,nullable"`
	JSON               customerCreditLedgerListResponseAmendmentLedgerEntryCustomerJSON `json:"-"`
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

func (r customerCreditLedgerListResponseAmendmentLedgerEntryCustomerJSON) RawJSON() string {
	return r.raw
}

type CustomerCreditLedgerListResponseAmendmentLedgerEntryEntryStatus string

const (
	CustomerCreditLedgerListResponseAmendmentLedgerEntryEntryStatusCommitted CustomerCreditLedgerListResponseAmendmentLedgerEntryEntryStatus = "committed"
	CustomerCreditLedgerListResponseAmendmentLedgerEntryEntryStatusPending   CustomerCreditLedgerListResponseAmendmentLedgerEntryEntryStatus = "pending"
)

func (r CustomerCreditLedgerListResponseAmendmentLedgerEntryEntryStatus) IsKnown() bool {
	switch r {
	case CustomerCreditLedgerListResponseAmendmentLedgerEntryEntryStatusCommitted, CustomerCreditLedgerListResponseAmendmentLedgerEntryEntryStatusPending:
		return true
	}
	return false
}

type CustomerCreditLedgerListResponseAmendmentLedgerEntryEntryType string

const (
	CustomerCreditLedgerListResponseAmendmentLedgerEntryEntryTypeAmendment CustomerCreditLedgerListResponseAmendmentLedgerEntryEntryType = "amendment"
)

func (r CustomerCreditLedgerListResponseAmendmentLedgerEntryEntryType) IsKnown() bool {
	switch r {
	case CustomerCreditLedgerListResponseAmendmentLedgerEntryEntryTypeAmendment:
		return true
	}
	return false
}

type CustomerCreditLedgerListResponseEntryStatus string

const (
	CustomerCreditLedgerListResponseEntryStatusCommitted CustomerCreditLedgerListResponseEntryStatus = "committed"
	CustomerCreditLedgerListResponseEntryStatusPending   CustomerCreditLedgerListResponseEntryStatus = "pending"
)

func (r CustomerCreditLedgerListResponseEntryStatus) IsKnown() bool {
	switch r {
	case CustomerCreditLedgerListResponseEntryStatusCommitted, CustomerCreditLedgerListResponseEntryStatusPending:
		return true
	}
	return false
}

type CustomerCreditLedgerListResponseEntryType string

const (
	CustomerCreditLedgerListResponseEntryTypeIncrement         CustomerCreditLedgerListResponseEntryType = "increment"
	CustomerCreditLedgerListResponseEntryTypeDecrement         CustomerCreditLedgerListResponseEntryType = "decrement"
	CustomerCreditLedgerListResponseEntryTypeExpirationChange  CustomerCreditLedgerListResponseEntryType = "expiration_change"
	CustomerCreditLedgerListResponseEntryTypeCreditBlockExpiry CustomerCreditLedgerListResponseEntryType = "credit_block_expiry"
	CustomerCreditLedgerListResponseEntryTypeVoid              CustomerCreditLedgerListResponseEntryType = "void"
	CustomerCreditLedgerListResponseEntryTypeVoidInitiated     CustomerCreditLedgerListResponseEntryType = "void_initiated"
	CustomerCreditLedgerListResponseEntryTypeAmendment         CustomerCreditLedgerListResponseEntryType = "amendment"
)

func (r CustomerCreditLedgerListResponseEntryType) IsKnown() bool {
	switch r {
	case CustomerCreditLedgerListResponseEntryTypeIncrement, CustomerCreditLedgerListResponseEntryTypeDecrement, CustomerCreditLedgerListResponseEntryTypeExpirationChange, CustomerCreditLedgerListResponseEntryTypeCreditBlockExpiry, CustomerCreditLedgerListResponseEntryTypeVoid, CustomerCreditLedgerListResponseEntryTypeVoidInitiated, CustomerCreditLedgerListResponseEntryTypeAmendment:
		return true
	}
	return false
}

// The [Credit Ledger Entry resource](/guides/product-catalog/prepurchase) models
// prepaid credits within Orb.
type CustomerCreditLedgerNewEntryResponse struct {
	// This field can have the runtime type of [map[string]string].
	Metadata             interface{}                                     `json:"metadata"`
	ID                   string                                          `json:"id,required"`
	LedgerSequenceNumber int64                                           `json:"ledger_sequence_number,required"`
	EntryStatus          CustomerCreditLedgerNewEntryResponseEntryStatus `json:"entry_status,required"`
	// This field can have the runtime type of
	// [CustomerCreditLedgerNewEntryResponseIncrementLedgerEntryCustomer],
	// [CustomerCreditLedgerNewEntryResponseDecrementLedgerEntryCustomer],
	// [CustomerCreditLedgerNewEntryResponseExpirationChangeLedgerEntryCustomer],
	// [CustomerCreditLedgerNewEntryResponseCreditBlockExpiryLedgerEntryCustomer],
	// [CustomerCreditLedgerNewEntryResponseVoidLedgerEntryCustomer],
	// [CustomerCreditLedgerNewEntryResponseVoidInitiatedLedgerEntryCustomer],
	// [CustomerCreditLedgerNewEntryResponseAmendmentLedgerEntryCustomer].
	Customer        interface{} `json:"customer"`
	StartingBalance float64     `json:"starting_balance,required"`
	EndingBalance   float64     `json:"ending_balance,required"`
	Amount          float64     `json:"amount,required"`
	Currency        string      `json:"currency,required"`
	CreatedAt       time.Time   `json:"created_at,required" format:"date-time"`
	Description     string      `json:"description,required,nullable"`
	// This field can have the runtime type of
	// [CustomerCreditLedgerNewEntryResponseIncrementLedgerEntryCreditBlock],
	// [CustomerCreditLedgerNewEntryResponseDecrementLedgerEntryCreditBlock],
	// [CustomerCreditLedgerNewEntryResponseExpirationChangeLedgerEntryCreditBlock],
	// [CustomerCreditLedgerNewEntryResponseCreditBlockExpiryLedgerEntryCreditBlock],
	// [CustomerCreditLedgerNewEntryResponseVoidLedgerEntryCreditBlock],
	// [CustomerCreditLedgerNewEntryResponseVoidInitiatedLedgerEntryCreditBlock],
	// [CustomerCreditLedgerNewEntryResponseAmendmentLedgerEntryCreditBlock].
	CreditBlock        interface{}                                   `json:"credit_block"`
	EntryType          CustomerCreditLedgerNewEntryResponseEntryType `json:"entry_type,required"`
	PriceID            string                                        `json:"price_id,nullable"`
	EventID            string                                        `json:"event_id,nullable"`
	InvoiceID          string                                        `json:"invoice_id,nullable"`
	NewBlockExpiryDate time.Time                                     `json:"new_block_expiry_date" format:"date-time"`
	VoidReason         string                                        `json:"void_reason,nullable"`
	VoidAmount         float64                                       `json:"void_amount"`
	JSON               customerCreditLedgerNewEntryResponseJSON      `json:"-"`
	union              CustomerCreditLedgerNewEntryResponseUnion
}

// customerCreditLedgerNewEntryResponseJSON contains the JSON metadata for the
// struct [CustomerCreditLedgerNewEntryResponse]
type customerCreditLedgerNewEntryResponseJSON struct {
	Metadata             apijson.Field
	ID                   apijson.Field
	LedgerSequenceNumber apijson.Field
	EntryStatus          apijson.Field
	Customer             apijson.Field
	StartingBalance      apijson.Field
	EndingBalance        apijson.Field
	Amount               apijson.Field
	Currency             apijson.Field
	CreatedAt            apijson.Field
	Description          apijson.Field
	CreditBlock          apijson.Field
	EntryType            apijson.Field
	PriceID              apijson.Field
	EventID              apijson.Field
	InvoiceID            apijson.Field
	NewBlockExpiryDate   apijson.Field
	VoidReason           apijson.Field
	VoidAmount           apijson.Field
	raw                  string
	ExtraFields          map[string]apijson.Field
}

func (r customerCreditLedgerNewEntryResponseJSON) RawJSON() string {
	return r.raw
}

func (r *CustomerCreditLedgerNewEntryResponse) UnmarshalJSON(data []byte) (err error) {
	*r = CustomerCreditLedgerNewEntryResponse{}
	err = apijson.UnmarshalRoot(data, &r.union)
	if err != nil {
		return err
	}
	return apijson.Port(r.union, &r)
}

// AsUnion returns a [CustomerCreditLedgerNewEntryResponseUnion] interface which
// you can cast to the specific types for more type safety.
//
// Possible runtime types of the union are
// [CustomerCreditLedgerNewEntryResponseIncrementLedgerEntry],
// [CustomerCreditLedgerNewEntryResponseDecrementLedgerEntry],
// [CustomerCreditLedgerNewEntryResponseExpirationChangeLedgerEntry],
// [CustomerCreditLedgerNewEntryResponseCreditBlockExpiryLedgerEntry],
// [CustomerCreditLedgerNewEntryResponseVoidLedgerEntry],
// [CustomerCreditLedgerNewEntryResponseVoidInitiatedLedgerEntry],
// [CustomerCreditLedgerNewEntryResponseAmendmentLedgerEntry].
func (r CustomerCreditLedgerNewEntryResponse) AsUnion() CustomerCreditLedgerNewEntryResponseUnion {
	return r.union
}

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
type CustomerCreditLedgerNewEntryResponseUnion interface {
	implementsCustomerCreditLedgerNewEntryResponse()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*CustomerCreditLedgerNewEntryResponseUnion)(nil)).Elem(),
		"entry_type",
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(CustomerCreditLedgerNewEntryResponseIncrementLedgerEntry{}),
			DiscriminatorValue: "increment",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(CustomerCreditLedgerNewEntryResponseDecrementLedgerEntry{}),
			DiscriminatorValue: "decrement",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(CustomerCreditLedgerNewEntryResponseExpirationChangeLedgerEntry{}),
			DiscriminatorValue: "expiration_change",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(CustomerCreditLedgerNewEntryResponseCreditBlockExpiryLedgerEntry{}),
			DiscriminatorValue: "credit_block_expiry",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(CustomerCreditLedgerNewEntryResponseVoidLedgerEntry{}),
			DiscriminatorValue: "void",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(CustomerCreditLedgerNewEntryResponseVoidInitiatedLedgerEntry{}),
			DiscriminatorValue: "void_initiated",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(CustomerCreditLedgerNewEntryResponseAmendmentLedgerEntry{}),
			DiscriminatorValue: "amendment",
		},
	)
}

type CustomerCreditLedgerNewEntryResponseIncrementLedgerEntry struct {
	ID                   string                                                              `json:"id,required"`
	Amount               float64                                                             `json:"amount,required"`
	CreatedAt            time.Time                                                           `json:"created_at,required" format:"date-time"`
	CreditBlock          CustomerCreditLedgerNewEntryResponseIncrementLedgerEntryCreditBlock `json:"credit_block,required"`
	Currency             string                                                              `json:"currency,required"`
	Customer             CustomerCreditLedgerNewEntryResponseIncrementLedgerEntryCustomer    `json:"customer,required"`
	Description          string                                                              `json:"description,required,nullable"`
	EndingBalance        float64                                                             `json:"ending_balance,required"`
	EntryStatus          CustomerCreditLedgerNewEntryResponseIncrementLedgerEntryEntryStatus `json:"entry_status,required"`
	EntryType            CustomerCreditLedgerNewEntryResponseIncrementLedgerEntryEntryType   `json:"entry_type,required"`
	LedgerSequenceNumber int64                                                               `json:"ledger_sequence_number,required"`
	// User specified key-value pairs for the resource. If not present, this defaults
	// to an empty dictionary. Individual keys can be removed by setting the value to
	// `null`, and the entire metadata mapping can be cleared by setting `metadata` to
	// `null`.
	Metadata        map[string]string                                            `json:"metadata,required"`
	StartingBalance float64                                                      `json:"starting_balance,required"`
	JSON            customerCreditLedgerNewEntryResponseIncrementLedgerEntryJSON `json:"-"`
}

// customerCreditLedgerNewEntryResponseIncrementLedgerEntryJSON contains the JSON
// metadata for the struct
// [CustomerCreditLedgerNewEntryResponseIncrementLedgerEntry]
type customerCreditLedgerNewEntryResponseIncrementLedgerEntryJSON struct {
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

func (r *CustomerCreditLedgerNewEntryResponseIncrementLedgerEntry) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r customerCreditLedgerNewEntryResponseIncrementLedgerEntryJSON) RawJSON() string {
	return r.raw
}

func (r CustomerCreditLedgerNewEntryResponseIncrementLedgerEntry) implementsCustomerCreditLedgerNewEntryResponse() {
}

type CustomerCreditLedgerNewEntryResponseIncrementLedgerEntryCreditBlock struct {
	ID               string                                                                  `json:"id,required"`
	ExpiryDate       time.Time                                                               `json:"expiry_date,required,nullable" format:"date-time"`
	PerUnitCostBasis string                                                                  `json:"per_unit_cost_basis,required,nullable"`
	JSON             customerCreditLedgerNewEntryResponseIncrementLedgerEntryCreditBlockJSON `json:"-"`
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

func (r customerCreditLedgerNewEntryResponseIncrementLedgerEntryCreditBlockJSON) RawJSON() string {
	return r.raw
}

type CustomerCreditLedgerNewEntryResponseIncrementLedgerEntryCustomer struct {
	ID                 string                                                               `json:"id,required"`
	ExternalCustomerID string                                                               `json:"external_customer_id,required,nullable"`
	JSON               customerCreditLedgerNewEntryResponseIncrementLedgerEntryCustomerJSON `json:"-"`
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

func (r customerCreditLedgerNewEntryResponseIncrementLedgerEntryCustomerJSON) RawJSON() string {
	return r.raw
}

type CustomerCreditLedgerNewEntryResponseIncrementLedgerEntryEntryStatus string

const (
	CustomerCreditLedgerNewEntryResponseIncrementLedgerEntryEntryStatusCommitted CustomerCreditLedgerNewEntryResponseIncrementLedgerEntryEntryStatus = "committed"
	CustomerCreditLedgerNewEntryResponseIncrementLedgerEntryEntryStatusPending   CustomerCreditLedgerNewEntryResponseIncrementLedgerEntryEntryStatus = "pending"
)

func (r CustomerCreditLedgerNewEntryResponseIncrementLedgerEntryEntryStatus) IsKnown() bool {
	switch r {
	case CustomerCreditLedgerNewEntryResponseIncrementLedgerEntryEntryStatusCommitted, CustomerCreditLedgerNewEntryResponseIncrementLedgerEntryEntryStatusPending:
		return true
	}
	return false
}

type CustomerCreditLedgerNewEntryResponseIncrementLedgerEntryEntryType string

const (
	CustomerCreditLedgerNewEntryResponseIncrementLedgerEntryEntryTypeIncrement CustomerCreditLedgerNewEntryResponseIncrementLedgerEntryEntryType = "increment"
)

func (r CustomerCreditLedgerNewEntryResponseIncrementLedgerEntryEntryType) IsKnown() bool {
	switch r {
	case CustomerCreditLedgerNewEntryResponseIncrementLedgerEntryEntryTypeIncrement:
		return true
	}
	return false
}

type CustomerCreditLedgerNewEntryResponseDecrementLedgerEntry struct {
	ID                   string                                                              `json:"id,required"`
	Amount               float64                                                             `json:"amount,required"`
	CreatedAt            time.Time                                                           `json:"created_at,required" format:"date-time"`
	CreditBlock          CustomerCreditLedgerNewEntryResponseDecrementLedgerEntryCreditBlock `json:"credit_block,required"`
	Currency             string                                                              `json:"currency,required"`
	Customer             CustomerCreditLedgerNewEntryResponseDecrementLedgerEntryCustomer    `json:"customer,required"`
	Description          string                                                              `json:"description,required,nullable"`
	EndingBalance        float64                                                             `json:"ending_balance,required"`
	EntryStatus          CustomerCreditLedgerNewEntryResponseDecrementLedgerEntryEntryStatus `json:"entry_status,required"`
	EntryType            CustomerCreditLedgerNewEntryResponseDecrementLedgerEntryEntryType   `json:"entry_type,required"`
	LedgerSequenceNumber int64                                                               `json:"ledger_sequence_number,required"`
	// User specified key-value pairs for the resource. If not present, this defaults
	// to an empty dictionary. Individual keys can be removed by setting the value to
	// `null`, and the entire metadata mapping can be cleared by setting `metadata` to
	// `null`.
	Metadata        map[string]string                                            `json:"metadata,required"`
	StartingBalance float64                                                      `json:"starting_balance,required"`
	EventID         string                                                       `json:"event_id,nullable"`
	InvoiceID       string                                                       `json:"invoice_id,nullable"`
	PriceID         string                                                       `json:"price_id,nullable"`
	JSON            customerCreditLedgerNewEntryResponseDecrementLedgerEntryJSON `json:"-"`
}

// customerCreditLedgerNewEntryResponseDecrementLedgerEntryJSON contains the JSON
// metadata for the struct
// [CustomerCreditLedgerNewEntryResponseDecrementLedgerEntry]
type customerCreditLedgerNewEntryResponseDecrementLedgerEntryJSON struct {
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

func (r *CustomerCreditLedgerNewEntryResponseDecrementLedgerEntry) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r customerCreditLedgerNewEntryResponseDecrementLedgerEntryJSON) RawJSON() string {
	return r.raw
}

func (r CustomerCreditLedgerNewEntryResponseDecrementLedgerEntry) implementsCustomerCreditLedgerNewEntryResponse() {
}

type CustomerCreditLedgerNewEntryResponseDecrementLedgerEntryCreditBlock struct {
	ID               string                                                                  `json:"id,required"`
	ExpiryDate       time.Time                                                               `json:"expiry_date,required,nullable" format:"date-time"`
	PerUnitCostBasis string                                                                  `json:"per_unit_cost_basis,required,nullable"`
	JSON             customerCreditLedgerNewEntryResponseDecrementLedgerEntryCreditBlockJSON `json:"-"`
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

func (r customerCreditLedgerNewEntryResponseDecrementLedgerEntryCreditBlockJSON) RawJSON() string {
	return r.raw
}

type CustomerCreditLedgerNewEntryResponseDecrementLedgerEntryCustomer struct {
	ID                 string                                                               `json:"id,required"`
	ExternalCustomerID string                                                               `json:"external_customer_id,required,nullable"`
	JSON               customerCreditLedgerNewEntryResponseDecrementLedgerEntryCustomerJSON `json:"-"`
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

func (r customerCreditLedgerNewEntryResponseDecrementLedgerEntryCustomerJSON) RawJSON() string {
	return r.raw
}

type CustomerCreditLedgerNewEntryResponseDecrementLedgerEntryEntryStatus string

const (
	CustomerCreditLedgerNewEntryResponseDecrementLedgerEntryEntryStatusCommitted CustomerCreditLedgerNewEntryResponseDecrementLedgerEntryEntryStatus = "committed"
	CustomerCreditLedgerNewEntryResponseDecrementLedgerEntryEntryStatusPending   CustomerCreditLedgerNewEntryResponseDecrementLedgerEntryEntryStatus = "pending"
)

func (r CustomerCreditLedgerNewEntryResponseDecrementLedgerEntryEntryStatus) IsKnown() bool {
	switch r {
	case CustomerCreditLedgerNewEntryResponseDecrementLedgerEntryEntryStatusCommitted, CustomerCreditLedgerNewEntryResponseDecrementLedgerEntryEntryStatusPending:
		return true
	}
	return false
}

type CustomerCreditLedgerNewEntryResponseDecrementLedgerEntryEntryType string

const (
	CustomerCreditLedgerNewEntryResponseDecrementLedgerEntryEntryTypeDecrement CustomerCreditLedgerNewEntryResponseDecrementLedgerEntryEntryType = "decrement"
)

func (r CustomerCreditLedgerNewEntryResponseDecrementLedgerEntryEntryType) IsKnown() bool {
	switch r {
	case CustomerCreditLedgerNewEntryResponseDecrementLedgerEntryEntryTypeDecrement:
		return true
	}
	return false
}

type CustomerCreditLedgerNewEntryResponseExpirationChangeLedgerEntry struct {
	ID                   string                                                                     `json:"id,required"`
	Amount               float64                                                                    `json:"amount,required"`
	CreatedAt            time.Time                                                                  `json:"created_at,required" format:"date-time"`
	CreditBlock          CustomerCreditLedgerNewEntryResponseExpirationChangeLedgerEntryCreditBlock `json:"credit_block,required"`
	Currency             string                                                                     `json:"currency,required"`
	Customer             CustomerCreditLedgerNewEntryResponseExpirationChangeLedgerEntryCustomer    `json:"customer,required"`
	Description          string                                                                     `json:"description,required,nullable"`
	EndingBalance        float64                                                                    `json:"ending_balance,required"`
	EntryStatus          CustomerCreditLedgerNewEntryResponseExpirationChangeLedgerEntryEntryStatus `json:"entry_status,required"`
	EntryType            CustomerCreditLedgerNewEntryResponseExpirationChangeLedgerEntryEntryType   `json:"entry_type,required"`
	LedgerSequenceNumber int64                                                                      `json:"ledger_sequence_number,required"`
	// User specified key-value pairs for the resource. If not present, this defaults
	// to an empty dictionary. Individual keys can be removed by setting the value to
	// `null`, and the entire metadata mapping can be cleared by setting `metadata` to
	// `null`.
	Metadata           map[string]string                                                   `json:"metadata,required"`
	NewBlockExpiryDate time.Time                                                           `json:"new_block_expiry_date,required" format:"date-time"`
	StartingBalance    float64                                                             `json:"starting_balance,required"`
	JSON               customerCreditLedgerNewEntryResponseExpirationChangeLedgerEntryJSON `json:"-"`
}

// customerCreditLedgerNewEntryResponseExpirationChangeLedgerEntryJSON contains the
// JSON metadata for the struct
// [CustomerCreditLedgerNewEntryResponseExpirationChangeLedgerEntry]
type customerCreditLedgerNewEntryResponseExpirationChangeLedgerEntryJSON struct {
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

func (r *CustomerCreditLedgerNewEntryResponseExpirationChangeLedgerEntry) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r customerCreditLedgerNewEntryResponseExpirationChangeLedgerEntryJSON) RawJSON() string {
	return r.raw
}

func (r CustomerCreditLedgerNewEntryResponseExpirationChangeLedgerEntry) implementsCustomerCreditLedgerNewEntryResponse() {
}

type CustomerCreditLedgerNewEntryResponseExpirationChangeLedgerEntryCreditBlock struct {
	ID               string                                                                         `json:"id,required"`
	ExpiryDate       time.Time                                                                      `json:"expiry_date,required,nullable" format:"date-time"`
	PerUnitCostBasis string                                                                         `json:"per_unit_cost_basis,required,nullable"`
	JSON             customerCreditLedgerNewEntryResponseExpirationChangeLedgerEntryCreditBlockJSON `json:"-"`
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

func (r customerCreditLedgerNewEntryResponseExpirationChangeLedgerEntryCreditBlockJSON) RawJSON() string {
	return r.raw
}

type CustomerCreditLedgerNewEntryResponseExpirationChangeLedgerEntryCustomer struct {
	ID                 string                                                                      `json:"id,required"`
	ExternalCustomerID string                                                                      `json:"external_customer_id,required,nullable"`
	JSON               customerCreditLedgerNewEntryResponseExpirationChangeLedgerEntryCustomerJSON `json:"-"`
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

func (r customerCreditLedgerNewEntryResponseExpirationChangeLedgerEntryCustomerJSON) RawJSON() string {
	return r.raw
}

type CustomerCreditLedgerNewEntryResponseExpirationChangeLedgerEntryEntryStatus string

const (
	CustomerCreditLedgerNewEntryResponseExpirationChangeLedgerEntryEntryStatusCommitted CustomerCreditLedgerNewEntryResponseExpirationChangeLedgerEntryEntryStatus = "committed"
	CustomerCreditLedgerNewEntryResponseExpirationChangeLedgerEntryEntryStatusPending   CustomerCreditLedgerNewEntryResponseExpirationChangeLedgerEntryEntryStatus = "pending"
)

func (r CustomerCreditLedgerNewEntryResponseExpirationChangeLedgerEntryEntryStatus) IsKnown() bool {
	switch r {
	case CustomerCreditLedgerNewEntryResponseExpirationChangeLedgerEntryEntryStatusCommitted, CustomerCreditLedgerNewEntryResponseExpirationChangeLedgerEntryEntryStatusPending:
		return true
	}
	return false
}

type CustomerCreditLedgerNewEntryResponseExpirationChangeLedgerEntryEntryType string

const (
	CustomerCreditLedgerNewEntryResponseExpirationChangeLedgerEntryEntryTypeExpirationChange CustomerCreditLedgerNewEntryResponseExpirationChangeLedgerEntryEntryType = "expiration_change"
)

func (r CustomerCreditLedgerNewEntryResponseExpirationChangeLedgerEntryEntryType) IsKnown() bool {
	switch r {
	case CustomerCreditLedgerNewEntryResponseExpirationChangeLedgerEntryEntryTypeExpirationChange:
		return true
	}
	return false
}

type CustomerCreditLedgerNewEntryResponseCreditBlockExpiryLedgerEntry struct {
	ID                   string                                                                      `json:"id,required"`
	Amount               float64                                                                     `json:"amount,required"`
	CreatedAt            time.Time                                                                   `json:"created_at,required" format:"date-time"`
	CreditBlock          CustomerCreditLedgerNewEntryResponseCreditBlockExpiryLedgerEntryCreditBlock `json:"credit_block,required"`
	Currency             string                                                                      `json:"currency,required"`
	Customer             CustomerCreditLedgerNewEntryResponseCreditBlockExpiryLedgerEntryCustomer    `json:"customer,required"`
	Description          string                                                                      `json:"description,required,nullable"`
	EndingBalance        float64                                                                     `json:"ending_balance,required"`
	EntryStatus          CustomerCreditLedgerNewEntryResponseCreditBlockExpiryLedgerEntryEntryStatus `json:"entry_status,required"`
	EntryType            CustomerCreditLedgerNewEntryResponseCreditBlockExpiryLedgerEntryEntryType   `json:"entry_type,required"`
	LedgerSequenceNumber int64                                                                       `json:"ledger_sequence_number,required"`
	// User specified key-value pairs for the resource. If not present, this defaults
	// to an empty dictionary. Individual keys can be removed by setting the value to
	// `null`, and the entire metadata mapping can be cleared by setting `metadata` to
	// `null`.
	Metadata        map[string]string                                                    `json:"metadata,required"`
	StartingBalance float64                                                              `json:"starting_balance,required"`
	JSON            customerCreditLedgerNewEntryResponseCreditBlockExpiryLedgerEntryJSON `json:"-"`
}

// customerCreditLedgerNewEntryResponseCreditBlockExpiryLedgerEntryJSON contains
// the JSON metadata for the struct
// [CustomerCreditLedgerNewEntryResponseCreditBlockExpiryLedgerEntry]
type customerCreditLedgerNewEntryResponseCreditBlockExpiryLedgerEntryJSON struct {
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

func (r *CustomerCreditLedgerNewEntryResponseCreditBlockExpiryLedgerEntry) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r customerCreditLedgerNewEntryResponseCreditBlockExpiryLedgerEntryJSON) RawJSON() string {
	return r.raw
}

func (r CustomerCreditLedgerNewEntryResponseCreditBlockExpiryLedgerEntry) implementsCustomerCreditLedgerNewEntryResponse() {
}

type CustomerCreditLedgerNewEntryResponseCreditBlockExpiryLedgerEntryCreditBlock struct {
	ID               string                                                                          `json:"id,required"`
	ExpiryDate       time.Time                                                                       `json:"expiry_date,required,nullable" format:"date-time"`
	PerUnitCostBasis string                                                                          `json:"per_unit_cost_basis,required,nullable"`
	JSON             customerCreditLedgerNewEntryResponseCreditBlockExpiryLedgerEntryCreditBlockJSON `json:"-"`
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

func (r customerCreditLedgerNewEntryResponseCreditBlockExpiryLedgerEntryCreditBlockJSON) RawJSON() string {
	return r.raw
}

type CustomerCreditLedgerNewEntryResponseCreditBlockExpiryLedgerEntryCustomer struct {
	ID                 string                                                                       `json:"id,required"`
	ExternalCustomerID string                                                                       `json:"external_customer_id,required,nullable"`
	JSON               customerCreditLedgerNewEntryResponseCreditBlockExpiryLedgerEntryCustomerJSON `json:"-"`
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

func (r customerCreditLedgerNewEntryResponseCreditBlockExpiryLedgerEntryCustomerJSON) RawJSON() string {
	return r.raw
}

type CustomerCreditLedgerNewEntryResponseCreditBlockExpiryLedgerEntryEntryStatus string

const (
	CustomerCreditLedgerNewEntryResponseCreditBlockExpiryLedgerEntryEntryStatusCommitted CustomerCreditLedgerNewEntryResponseCreditBlockExpiryLedgerEntryEntryStatus = "committed"
	CustomerCreditLedgerNewEntryResponseCreditBlockExpiryLedgerEntryEntryStatusPending   CustomerCreditLedgerNewEntryResponseCreditBlockExpiryLedgerEntryEntryStatus = "pending"
)

func (r CustomerCreditLedgerNewEntryResponseCreditBlockExpiryLedgerEntryEntryStatus) IsKnown() bool {
	switch r {
	case CustomerCreditLedgerNewEntryResponseCreditBlockExpiryLedgerEntryEntryStatusCommitted, CustomerCreditLedgerNewEntryResponseCreditBlockExpiryLedgerEntryEntryStatusPending:
		return true
	}
	return false
}

type CustomerCreditLedgerNewEntryResponseCreditBlockExpiryLedgerEntryEntryType string

const (
	CustomerCreditLedgerNewEntryResponseCreditBlockExpiryLedgerEntryEntryTypeCreditBlockExpiry CustomerCreditLedgerNewEntryResponseCreditBlockExpiryLedgerEntryEntryType = "credit_block_expiry"
)

func (r CustomerCreditLedgerNewEntryResponseCreditBlockExpiryLedgerEntryEntryType) IsKnown() bool {
	switch r {
	case CustomerCreditLedgerNewEntryResponseCreditBlockExpiryLedgerEntryEntryTypeCreditBlockExpiry:
		return true
	}
	return false
}

type CustomerCreditLedgerNewEntryResponseVoidLedgerEntry struct {
	ID                   string                                                         `json:"id,required"`
	Amount               float64                                                        `json:"amount,required"`
	CreatedAt            time.Time                                                      `json:"created_at,required" format:"date-time"`
	CreditBlock          CustomerCreditLedgerNewEntryResponseVoidLedgerEntryCreditBlock `json:"credit_block,required"`
	Currency             string                                                         `json:"currency,required"`
	Customer             CustomerCreditLedgerNewEntryResponseVoidLedgerEntryCustomer    `json:"customer,required"`
	Description          string                                                         `json:"description,required,nullable"`
	EndingBalance        float64                                                        `json:"ending_balance,required"`
	EntryStatus          CustomerCreditLedgerNewEntryResponseVoidLedgerEntryEntryStatus `json:"entry_status,required"`
	EntryType            CustomerCreditLedgerNewEntryResponseVoidLedgerEntryEntryType   `json:"entry_type,required"`
	LedgerSequenceNumber int64                                                          `json:"ledger_sequence_number,required"`
	// User specified key-value pairs for the resource. If not present, this defaults
	// to an empty dictionary. Individual keys can be removed by setting the value to
	// `null`, and the entire metadata mapping can be cleared by setting `metadata` to
	// `null`.
	Metadata        map[string]string                                       `json:"metadata,required"`
	StartingBalance float64                                                 `json:"starting_balance,required"`
	VoidAmount      float64                                                 `json:"void_amount,required"`
	VoidReason      string                                                  `json:"void_reason,required,nullable"`
	JSON            customerCreditLedgerNewEntryResponseVoidLedgerEntryJSON `json:"-"`
}

// customerCreditLedgerNewEntryResponseVoidLedgerEntryJSON contains the JSON
// metadata for the struct [CustomerCreditLedgerNewEntryResponseVoidLedgerEntry]
type customerCreditLedgerNewEntryResponseVoidLedgerEntryJSON struct {
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

func (r *CustomerCreditLedgerNewEntryResponseVoidLedgerEntry) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r customerCreditLedgerNewEntryResponseVoidLedgerEntryJSON) RawJSON() string {
	return r.raw
}

func (r CustomerCreditLedgerNewEntryResponseVoidLedgerEntry) implementsCustomerCreditLedgerNewEntryResponse() {
}

type CustomerCreditLedgerNewEntryResponseVoidLedgerEntryCreditBlock struct {
	ID               string                                                             `json:"id,required"`
	ExpiryDate       time.Time                                                          `json:"expiry_date,required,nullable" format:"date-time"`
	PerUnitCostBasis string                                                             `json:"per_unit_cost_basis,required,nullable"`
	JSON             customerCreditLedgerNewEntryResponseVoidLedgerEntryCreditBlockJSON `json:"-"`
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

func (r customerCreditLedgerNewEntryResponseVoidLedgerEntryCreditBlockJSON) RawJSON() string {
	return r.raw
}

type CustomerCreditLedgerNewEntryResponseVoidLedgerEntryCustomer struct {
	ID                 string                                                          `json:"id,required"`
	ExternalCustomerID string                                                          `json:"external_customer_id,required,nullable"`
	JSON               customerCreditLedgerNewEntryResponseVoidLedgerEntryCustomerJSON `json:"-"`
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

func (r customerCreditLedgerNewEntryResponseVoidLedgerEntryCustomerJSON) RawJSON() string {
	return r.raw
}

type CustomerCreditLedgerNewEntryResponseVoidLedgerEntryEntryStatus string

const (
	CustomerCreditLedgerNewEntryResponseVoidLedgerEntryEntryStatusCommitted CustomerCreditLedgerNewEntryResponseVoidLedgerEntryEntryStatus = "committed"
	CustomerCreditLedgerNewEntryResponseVoidLedgerEntryEntryStatusPending   CustomerCreditLedgerNewEntryResponseVoidLedgerEntryEntryStatus = "pending"
)

func (r CustomerCreditLedgerNewEntryResponseVoidLedgerEntryEntryStatus) IsKnown() bool {
	switch r {
	case CustomerCreditLedgerNewEntryResponseVoidLedgerEntryEntryStatusCommitted, CustomerCreditLedgerNewEntryResponseVoidLedgerEntryEntryStatusPending:
		return true
	}
	return false
}

type CustomerCreditLedgerNewEntryResponseVoidLedgerEntryEntryType string

const (
	CustomerCreditLedgerNewEntryResponseVoidLedgerEntryEntryTypeVoid CustomerCreditLedgerNewEntryResponseVoidLedgerEntryEntryType = "void"
)

func (r CustomerCreditLedgerNewEntryResponseVoidLedgerEntryEntryType) IsKnown() bool {
	switch r {
	case CustomerCreditLedgerNewEntryResponseVoidLedgerEntryEntryTypeVoid:
		return true
	}
	return false
}

type CustomerCreditLedgerNewEntryResponseVoidInitiatedLedgerEntry struct {
	ID                   string                                                                  `json:"id,required"`
	Amount               float64                                                                 `json:"amount,required"`
	CreatedAt            time.Time                                                               `json:"created_at,required" format:"date-time"`
	CreditBlock          CustomerCreditLedgerNewEntryResponseVoidInitiatedLedgerEntryCreditBlock `json:"credit_block,required"`
	Currency             string                                                                  `json:"currency,required"`
	Customer             CustomerCreditLedgerNewEntryResponseVoidInitiatedLedgerEntryCustomer    `json:"customer,required"`
	Description          string                                                                  `json:"description,required,nullable"`
	EndingBalance        float64                                                                 `json:"ending_balance,required"`
	EntryStatus          CustomerCreditLedgerNewEntryResponseVoidInitiatedLedgerEntryEntryStatus `json:"entry_status,required"`
	EntryType            CustomerCreditLedgerNewEntryResponseVoidInitiatedLedgerEntryEntryType   `json:"entry_type,required"`
	LedgerSequenceNumber int64                                                                   `json:"ledger_sequence_number,required"`
	// User specified key-value pairs for the resource. If not present, this defaults
	// to an empty dictionary. Individual keys can be removed by setting the value to
	// `null`, and the entire metadata mapping can be cleared by setting `metadata` to
	// `null`.
	Metadata           map[string]string                                                `json:"metadata,required"`
	NewBlockExpiryDate time.Time                                                        `json:"new_block_expiry_date,required" format:"date-time"`
	StartingBalance    float64                                                          `json:"starting_balance,required"`
	VoidAmount         float64                                                          `json:"void_amount,required"`
	VoidReason         string                                                           `json:"void_reason,required,nullable"`
	JSON               customerCreditLedgerNewEntryResponseVoidInitiatedLedgerEntryJSON `json:"-"`
}

// customerCreditLedgerNewEntryResponseVoidInitiatedLedgerEntryJSON contains the
// JSON metadata for the struct
// [CustomerCreditLedgerNewEntryResponseVoidInitiatedLedgerEntry]
type customerCreditLedgerNewEntryResponseVoidInitiatedLedgerEntryJSON struct {
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

func (r *CustomerCreditLedgerNewEntryResponseVoidInitiatedLedgerEntry) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r customerCreditLedgerNewEntryResponseVoidInitiatedLedgerEntryJSON) RawJSON() string {
	return r.raw
}

func (r CustomerCreditLedgerNewEntryResponseVoidInitiatedLedgerEntry) implementsCustomerCreditLedgerNewEntryResponse() {
}

type CustomerCreditLedgerNewEntryResponseVoidInitiatedLedgerEntryCreditBlock struct {
	ID               string                                                                      `json:"id,required"`
	ExpiryDate       time.Time                                                                   `json:"expiry_date,required,nullable" format:"date-time"`
	PerUnitCostBasis string                                                                      `json:"per_unit_cost_basis,required,nullable"`
	JSON             customerCreditLedgerNewEntryResponseVoidInitiatedLedgerEntryCreditBlockJSON `json:"-"`
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

func (r customerCreditLedgerNewEntryResponseVoidInitiatedLedgerEntryCreditBlockJSON) RawJSON() string {
	return r.raw
}

type CustomerCreditLedgerNewEntryResponseVoidInitiatedLedgerEntryCustomer struct {
	ID                 string                                                                   `json:"id,required"`
	ExternalCustomerID string                                                                   `json:"external_customer_id,required,nullable"`
	JSON               customerCreditLedgerNewEntryResponseVoidInitiatedLedgerEntryCustomerJSON `json:"-"`
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

func (r customerCreditLedgerNewEntryResponseVoidInitiatedLedgerEntryCustomerJSON) RawJSON() string {
	return r.raw
}

type CustomerCreditLedgerNewEntryResponseVoidInitiatedLedgerEntryEntryStatus string

const (
	CustomerCreditLedgerNewEntryResponseVoidInitiatedLedgerEntryEntryStatusCommitted CustomerCreditLedgerNewEntryResponseVoidInitiatedLedgerEntryEntryStatus = "committed"
	CustomerCreditLedgerNewEntryResponseVoidInitiatedLedgerEntryEntryStatusPending   CustomerCreditLedgerNewEntryResponseVoidInitiatedLedgerEntryEntryStatus = "pending"
)

func (r CustomerCreditLedgerNewEntryResponseVoidInitiatedLedgerEntryEntryStatus) IsKnown() bool {
	switch r {
	case CustomerCreditLedgerNewEntryResponseVoidInitiatedLedgerEntryEntryStatusCommitted, CustomerCreditLedgerNewEntryResponseVoidInitiatedLedgerEntryEntryStatusPending:
		return true
	}
	return false
}

type CustomerCreditLedgerNewEntryResponseVoidInitiatedLedgerEntryEntryType string

const (
	CustomerCreditLedgerNewEntryResponseVoidInitiatedLedgerEntryEntryTypeVoidInitiated CustomerCreditLedgerNewEntryResponseVoidInitiatedLedgerEntryEntryType = "void_initiated"
)

func (r CustomerCreditLedgerNewEntryResponseVoidInitiatedLedgerEntryEntryType) IsKnown() bool {
	switch r {
	case CustomerCreditLedgerNewEntryResponseVoidInitiatedLedgerEntryEntryTypeVoidInitiated:
		return true
	}
	return false
}

type CustomerCreditLedgerNewEntryResponseAmendmentLedgerEntry struct {
	ID                   string                                                              `json:"id,required"`
	Amount               float64                                                             `json:"amount,required"`
	CreatedAt            time.Time                                                           `json:"created_at,required" format:"date-time"`
	CreditBlock          CustomerCreditLedgerNewEntryResponseAmendmentLedgerEntryCreditBlock `json:"credit_block,required"`
	Currency             string                                                              `json:"currency,required"`
	Customer             CustomerCreditLedgerNewEntryResponseAmendmentLedgerEntryCustomer    `json:"customer,required"`
	Description          string                                                              `json:"description,required,nullable"`
	EndingBalance        float64                                                             `json:"ending_balance,required"`
	EntryStatus          CustomerCreditLedgerNewEntryResponseAmendmentLedgerEntryEntryStatus `json:"entry_status,required"`
	EntryType            CustomerCreditLedgerNewEntryResponseAmendmentLedgerEntryEntryType   `json:"entry_type,required"`
	LedgerSequenceNumber int64                                                               `json:"ledger_sequence_number,required"`
	// User specified key-value pairs for the resource. If not present, this defaults
	// to an empty dictionary. Individual keys can be removed by setting the value to
	// `null`, and the entire metadata mapping can be cleared by setting `metadata` to
	// `null`.
	Metadata        map[string]string                                            `json:"metadata,required"`
	StartingBalance float64                                                      `json:"starting_balance,required"`
	JSON            customerCreditLedgerNewEntryResponseAmendmentLedgerEntryJSON `json:"-"`
}

// customerCreditLedgerNewEntryResponseAmendmentLedgerEntryJSON contains the JSON
// metadata for the struct
// [CustomerCreditLedgerNewEntryResponseAmendmentLedgerEntry]
type customerCreditLedgerNewEntryResponseAmendmentLedgerEntryJSON struct {
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

func (r *CustomerCreditLedgerNewEntryResponseAmendmentLedgerEntry) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r customerCreditLedgerNewEntryResponseAmendmentLedgerEntryJSON) RawJSON() string {
	return r.raw
}

func (r CustomerCreditLedgerNewEntryResponseAmendmentLedgerEntry) implementsCustomerCreditLedgerNewEntryResponse() {
}

type CustomerCreditLedgerNewEntryResponseAmendmentLedgerEntryCreditBlock struct {
	ID               string                                                                  `json:"id,required"`
	ExpiryDate       time.Time                                                               `json:"expiry_date,required,nullable" format:"date-time"`
	PerUnitCostBasis string                                                                  `json:"per_unit_cost_basis,required,nullable"`
	JSON             customerCreditLedgerNewEntryResponseAmendmentLedgerEntryCreditBlockJSON `json:"-"`
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

func (r customerCreditLedgerNewEntryResponseAmendmentLedgerEntryCreditBlockJSON) RawJSON() string {
	return r.raw
}

type CustomerCreditLedgerNewEntryResponseAmendmentLedgerEntryCustomer struct {
	ID                 string                                                               `json:"id,required"`
	ExternalCustomerID string                                                               `json:"external_customer_id,required,nullable"`
	JSON               customerCreditLedgerNewEntryResponseAmendmentLedgerEntryCustomerJSON `json:"-"`
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

func (r customerCreditLedgerNewEntryResponseAmendmentLedgerEntryCustomerJSON) RawJSON() string {
	return r.raw
}

type CustomerCreditLedgerNewEntryResponseAmendmentLedgerEntryEntryStatus string

const (
	CustomerCreditLedgerNewEntryResponseAmendmentLedgerEntryEntryStatusCommitted CustomerCreditLedgerNewEntryResponseAmendmentLedgerEntryEntryStatus = "committed"
	CustomerCreditLedgerNewEntryResponseAmendmentLedgerEntryEntryStatusPending   CustomerCreditLedgerNewEntryResponseAmendmentLedgerEntryEntryStatus = "pending"
)

func (r CustomerCreditLedgerNewEntryResponseAmendmentLedgerEntryEntryStatus) IsKnown() bool {
	switch r {
	case CustomerCreditLedgerNewEntryResponseAmendmentLedgerEntryEntryStatusCommitted, CustomerCreditLedgerNewEntryResponseAmendmentLedgerEntryEntryStatusPending:
		return true
	}
	return false
}

type CustomerCreditLedgerNewEntryResponseAmendmentLedgerEntryEntryType string

const (
	CustomerCreditLedgerNewEntryResponseAmendmentLedgerEntryEntryTypeAmendment CustomerCreditLedgerNewEntryResponseAmendmentLedgerEntryEntryType = "amendment"
)

func (r CustomerCreditLedgerNewEntryResponseAmendmentLedgerEntryEntryType) IsKnown() bool {
	switch r {
	case CustomerCreditLedgerNewEntryResponseAmendmentLedgerEntryEntryTypeAmendment:
		return true
	}
	return false
}

type CustomerCreditLedgerNewEntryResponseEntryStatus string

const (
	CustomerCreditLedgerNewEntryResponseEntryStatusCommitted CustomerCreditLedgerNewEntryResponseEntryStatus = "committed"
	CustomerCreditLedgerNewEntryResponseEntryStatusPending   CustomerCreditLedgerNewEntryResponseEntryStatus = "pending"
)

func (r CustomerCreditLedgerNewEntryResponseEntryStatus) IsKnown() bool {
	switch r {
	case CustomerCreditLedgerNewEntryResponseEntryStatusCommitted, CustomerCreditLedgerNewEntryResponseEntryStatusPending:
		return true
	}
	return false
}

type CustomerCreditLedgerNewEntryResponseEntryType string

const (
	CustomerCreditLedgerNewEntryResponseEntryTypeIncrement         CustomerCreditLedgerNewEntryResponseEntryType = "increment"
	CustomerCreditLedgerNewEntryResponseEntryTypeDecrement         CustomerCreditLedgerNewEntryResponseEntryType = "decrement"
	CustomerCreditLedgerNewEntryResponseEntryTypeExpirationChange  CustomerCreditLedgerNewEntryResponseEntryType = "expiration_change"
	CustomerCreditLedgerNewEntryResponseEntryTypeCreditBlockExpiry CustomerCreditLedgerNewEntryResponseEntryType = "credit_block_expiry"
	CustomerCreditLedgerNewEntryResponseEntryTypeVoid              CustomerCreditLedgerNewEntryResponseEntryType = "void"
	CustomerCreditLedgerNewEntryResponseEntryTypeVoidInitiated     CustomerCreditLedgerNewEntryResponseEntryType = "void_initiated"
	CustomerCreditLedgerNewEntryResponseEntryTypeAmendment         CustomerCreditLedgerNewEntryResponseEntryType = "amendment"
)

func (r CustomerCreditLedgerNewEntryResponseEntryType) IsKnown() bool {
	switch r {
	case CustomerCreditLedgerNewEntryResponseEntryTypeIncrement, CustomerCreditLedgerNewEntryResponseEntryTypeDecrement, CustomerCreditLedgerNewEntryResponseEntryTypeExpirationChange, CustomerCreditLedgerNewEntryResponseEntryTypeCreditBlockExpiry, CustomerCreditLedgerNewEntryResponseEntryTypeVoid, CustomerCreditLedgerNewEntryResponseEntryTypeVoidInitiated, CustomerCreditLedgerNewEntryResponseEntryTypeAmendment:
		return true
	}
	return false
}

// The [Credit Ledger Entry resource](/guides/product-catalog/prepurchase) models
// prepaid credits within Orb.
type CustomerCreditLedgerNewEntryByExternalIDResponse struct {
	// This field can have the runtime type of [map[string]string].
	Metadata             interface{}                                                 `json:"metadata"`
	ID                   string                                                      `json:"id,required"`
	LedgerSequenceNumber int64                                                       `json:"ledger_sequence_number,required"`
	EntryStatus          CustomerCreditLedgerNewEntryByExternalIDResponseEntryStatus `json:"entry_status,required"`
	// This field can have the runtime type of
	// [CustomerCreditLedgerNewEntryByExternalIDResponseIncrementLedgerEntryCustomer],
	// [CustomerCreditLedgerNewEntryByExternalIDResponseDecrementLedgerEntryCustomer],
	// [CustomerCreditLedgerNewEntryByExternalIDResponseExpirationChangeLedgerEntryCustomer],
	// [CustomerCreditLedgerNewEntryByExternalIDResponseCreditBlockExpiryLedgerEntryCustomer],
	// [CustomerCreditLedgerNewEntryByExternalIDResponseVoidLedgerEntryCustomer],
	// [CustomerCreditLedgerNewEntryByExternalIDResponseVoidInitiatedLedgerEntryCustomer],
	// [CustomerCreditLedgerNewEntryByExternalIDResponseAmendmentLedgerEntryCustomer].
	Customer        interface{} `json:"customer"`
	StartingBalance float64     `json:"starting_balance,required"`
	EndingBalance   float64     `json:"ending_balance,required"`
	Amount          float64     `json:"amount,required"`
	Currency        string      `json:"currency,required"`
	CreatedAt       time.Time   `json:"created_at,required" format:"date-time"`
	Description     string      `json:"description,required,nullable"`
	// This field can have the runtime type of
	// [CustomerCreditLedgerNewEntryByExternalIDResponseIncrementLedgerEntryCreditBlock],
	// [CustomerCreditLedgerNewEntryByExternalIDResponseDecrementLedgerEntryCreditBlock],
	// [CustomerCreditLedgerNewEntryByExternalIDResponseExpirationChangeLedgerEntryCreditBlock],
	// [CustomerCreditLedgerNewEntryByExternalIDResponseCreditBlockExpiryLedgerEntryCreditBlock],
	// [CustomerCreditLedgerNewEntryByExternalIDResponseVoidLedgerEntryCreditBlock],
	// [CustomerCreditLedgerNewEntryByExternalIDResponseVoidInitiatedLedgerEntryCreditBlock],
	// [CustomerCreditLedgerNewEntryByExternalIDResponseAmendmentLedgerEntryCreditBlock].
	CreditBlock        interface{}                                               `json:"credit_block"`
	EntryType          CustomerCreditLedgerNewEntryByExternalIDResponseEntryType `json:"entry_type,required"`
	PriceID            string                                                    `json:"price_id,nullable"`
	EventID            string                                                    `json:"event_id,nullable"`
	InvoiceID          string                                                    `json:"invoice_id,nullable"`
	NewBlockExpiryDate time.Time                                                 `json:"new_block_expiry_date" format:"date-time"`
	VoidReason         string                                                    `json:"void_reason,nullable"`
	VoidAmount         float64                                                   `json:"void_amount"`
	JSON               customerCreditLedgerNewEntryByExternalIDResponseJSON      `json:"-"`
	union              CustomerCreditLedgerNewEntryByExternalIDResponseUnion
}

// customerCreditLedgerNewEntryByExternalIDResponseJSON contains the JSON metadata
// for the struct [CustomerCreditLedgerNewEntryByExternalIDResponse]
type customerCreditLedgerNewEntryByExternalIDResponseJSON struct {
	Metadata             apijson.Field
	ID                   apijson.Field
	LedgerSequenceNumber apijson.Field
	EntryStatus          apijson.Field
	Customer             apijson.Field
	StartingBalance      apijson.Field
	EndingBalance        apijson.Field
	Amount               apijson.Field
	Currency             apijson.Field
	CreatedAt            apijson.Field
	Description          apijson.Field
	CreditBlock          apijson.Field
	EntryType            apijson.Field
	PriceID              apijson.Field
	EventID              apijson.Field
	InvoiceID            apijson.Field
	NewBlockExpiryDate   apijson.Field
	VoidReason           apijson.Field
	VoidAmount           apijson.Field
	raw                  string
	ExtraFields          map[string]apijson.Field
}

func (r customerCreditLedgerNewEntryByExternalIDResponseJSON) RawJSON() string {
	return r.raw
}

func (r *CustomerCreditLedgerNewEntryByExternalIDResponse) UnmarshalJSON(data []byte) (err error) {
	*r = CustomerCreditLedgerNewEntryByExternalIDResponse{}
	err = apijson.UnmarshalRoot(data, &r.union)
	if err != nil {
		return err
	}
	return apijson.Port(r.union, &r)
}

// AsUnion returns a [CustomerCreditLedgerNewEntryByExternalIDResponseUnion]
// interface which you can cast to the specific types for more type safety.
//
// Possible runtime types of the union are
// [CustomerCreditLedgerNewEntryByExternalIDResponseIncrementLedgerEntry],
// [CustomerCreditLedgerNewEntryByExternalIDResponseDecrementLedgerEntry],
// [CustomerCreditLedgerNewEntryByExternalIDResponseExpirationChangeLedgerEntry],
// [CustomerCreditLedgerNewEntryByExternalIDResponseCreditBlockExpiryLedgerEntry],
// [CustomerCreditLedgerNewEntryByExternalIDResponseVoidLedgerEntry],
// [CustomerCreditLedgerNewEntryByExternalIDResponseVoidInitiatedLedgerEntry],
// [CustomerCreditLedgerNewEntryByExternalIDResponseAmendmentLedgerEntry].
func (r CustomerCreditLedgerNewEntryByExternalIDResponse) AsUnion() CustomerCreditLedgerNewEntryByExternalIDResponseUnion {
	return r.union
}

// The [Credit Ledger Entry resource](/guides/product-catalog/prepurchase) models
// prepaid credits within Orb.
//
// Union satisfied by
// [CustomerCreditLedgerNewEntryByExternalIDResponseIncrementLedgerEntry],
// [CustomerCreditLedgerNewEntryByExternalIDResponseDecrementLedgerEntry],
// [CustomerCreditLedgerNewEntryByExternalIDResponseExpirationChangeLedgerEntry],
// [CustomerCreditLedgerNewEntryByExternalIDResponseCreditBlockExpiryLedgerEntry],
// [CustomerCreditLedgerNewEntryByExternalIDResponseVoidLedgerEntry],
// [CustomerCreditLedgerNewEntryByExternalIDResponseVoidInitiatedLedgerEntry] or
// [CustomerCreditLedgerNewEntryByExternalIDResponseAmendmentLedgerEntry].
type CustomerCreditLedgerNewEntryByExternalIDResponseUnion interface {
	implementsCustomerCreditLedgerNewEntryByExternalIDResponse()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*CustomerCreditLedgerNewEntryByExternalIDResponseUnion)(nil)).Elem(),
		"entry_type",
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(CustomerCreditLedgerNewEntryByExternalIDResponseIncrementLedgerEntry{}),
			DiscriminatorValue: "increment",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(CustomerCreditLedgerNewEntryByExternalIDResponseDecrementLedgerEntry{}),
			DiscriminatorValue: "decrement",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(CustomerCreditLedgerNewEntryByExternalIDResponseExpirationChangeLedgerEntry{}),
			DiscriminatorValue: "expiration_change",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(CustomerCreditLedgerNewEntryByExternalIDResponseCreditBlockExpiryLedgerEntry{}),
			DiscriminatorValue: "credit_block_expiry",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(CustomerCreditLedgerNewEntryByExternalIDResponseVoidLedgerEntry{}),
			DiscriminatorValue: "void",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(CustomerCreditLedgerNewEntryByExternalIDResponseVoidInitiatedLedgerEntry{}),
			DiscriminatorValue: "void_initiated",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(CustomerCreditLedgerNewEntryByExternalIDResponseAmendmentLedgerEntry{}),
			DiscriminatorValue: "amendment",
		},
	)
}

type CustomerCreditLedgerNewEntryByExternalIDResponseIncrementLedgerEntry struct {
	ID                   string                                                                          `json:"id,required"`
	Amount               float64                                                                         `json:"amount,required"`
	CreatedAt            time.Time                                                                       `json:"created_at,required" format:"date-time"`
	CreditBlock          CustomerCreditLedgerNewEntryByExternalIDResponseIncrementLedgerEntryCreditBlock `json:"credit_block,required"`
	Currency             string                                                                          `json:"currency,required"`
	Customer             CustomerCreditLedgerNewEntryByExternalIDResponseIncrementLedgerEntryCustomer    `json:"customer,required"`
	Description          string                                                                          `json:"description,required,nullable"`
	EndingBalance        float64                                                                         `json:"ending_balance,required"`
	EntryStatus          CustomerCreditLedgerNewEntryByExternalIDResponseIncrementLedgerEntryEntryStatus `json:"entry_status,required"`
	EntryType            CustomerCreditLedgerNewEntryByExternalIDResponseIncrementLedgerEntryEntryType   `json:"entry_type,required"`
	LedgerSequenceNumber int64                                                                           `json:"ledger_sequence_number,required"`
	// User specified key-value pairs for the resource. If not present, this defaults
	// to an empty dictionary. Individual keys can be removed by setting the value to
	// `null`, and the entire metadata mapping can be cleared by setting `metadata` to
	// `null`.
	Metadata        map[string]string                                                        `json:"metadata,required"`
	StartingBalance float64                                                                  `json:"starting_balance,required"`
	JSON            customerCreditLedgerNewEntryByExternalIDResponseIncrementLedgerEntryJSON `json:"-"`
}

// customerCreditLedgerNewEntryByExternalIDResponseIncrementLedgerEntryJSON
// contains the JSON metadata for the struct
// [CustomerCreditLedgerNewEntryByExternalIDResponseIncrementLedgerEntry]
type customerCreditLedgerNewEntryByExternalIDResponseIncrementLedgerEntryJSON struct {
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

func (r *CustomerCreditLedgerNewEntryByExternalIDResponseIncrementLedgerEntry) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r customerCreditLedgerNewEntryByExternalIDResponseIncrementLedgerEntryJSON) RawJSON() string {
	return r.raw
}

func (r CustomerCreditLedgerNewEntryByExternalIDResponseIncrementLedgerEntry) implementsCustomerCreditLedgerNewEntryByExternalIDResponse() {
}

type CustomerCreditLedgerNewEntryByExternalIDResponseIncrementLedgerEntryCreditBlock struct {
	ID               string                                                                              `json:"id,required"`
	ExpiryDate       time.Time                                                                           `json:"expiry_date,required,nullable" format:"date-time"`
	PerUnitCostBasis string                                                                              `json:"per_unit_cost_basis,required,nullable"`
	JSON             customerCreditLedgerNewEntryByExternalIDResponseIncrementLedgerEntryCreditBlockJSON `json:"-"`
}

// customerCreditLedgerNewEntryByExternalIDResponseIncrementLedgerEntryCreditBlockJSON
// contains the JSON metadata for the struct
// [CustomerCreditLedgerNewEntryByExternalIDResponseIncrementLedgerEntryCreditBlock]
type customerCreditLedgerNewEntryByExternalIDResponseIncrementLedgerEntryCreditBlockJSON struct {
	ID               apijson.Field
	ExpiryDate       apijson.Field
	PerUnitCostBasis apijson.Field
	raw              string
	ExtraFields      map[string]apijson.Field
}

func (r *CustomerCreditLedgerNewEntryByExternalIDResponseIncrementLedgerEntryCreditBlock) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r customerCreditLedgerNewEntryByExternalIDResponseIncrementLedgerEntryCreditBlockJSON) RawJSON() string {
	return r.raw
}

type CustomerCreditLedgerNewEntryByExternalIDResponseIncrementLedgerEntryCustomer struct {
	ID                 string                                                                           `json:"id,required"`
	ExternalCustomerID string                                                                           `json:"external_customer_id,required,nullable"`
	JSON               customerCreditLedgerNewEntryByExternalIDResponseIncrementLedgerEntryCustomerJSON `json:"-"`
}

// customerCreditLedgerNewEntryByExternalIDResponseIncrementLedgerEntryCustomerJSON
// contains the JSON metadata for the struct
// [CustomerCreditLedgerNewEntryByExternalIDResponseIncrementLedgerEntryCustomer]
type customerCreditLedgerNewEntryByExternalIDResponseIncrementLedgerEntryCustomerJSON struct {
	ID                 apijson.Field
	ExternalCustomerID apijson.Field
	raw                string
	ExtraFields        map[string]apijson.Field
}

func (r *CustomerCreditLedgerNewEntryByExternalIDResponseIncrementLedgerEntryCustomer) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r customerCreditLedgerNewEntryByExternalIDResponseIncrementLedgerEntryCustomerJSON) RawJSON() string {
	return r.raw
}

type CustomerCreditLedgerNewEntryByExternalIDResponseIncrementLedgerEntryEntryStatus string

const (
	CustomerCreditLedgerNewEntryByExternalIDResponseIncrementLedgerEntryEntryStatusCommitted CustomerCreditLedgerNewEntryByExternalIDResponseIncrementLedgerEntryEntryStatus = "committed"
	CustomerCreditLedgerNewEntryByExternalIDResponseIncrementLedgerEntryEntryStatusPending   CustomerCreditLedgerNewEntryByExternalIDResponseIncrementLedgerEntryEntryStatus = "pending"
)

func (r CustomerCreditLedgerNewEntryByExternalIDResponseIncrementLedgerEntryEntryStatus) IsKnown() bool {
	switch r {
	case CustomerCreditLedgerNewEntryByExternalIDResponseIncrementLedgerEntryEntryStatusCommitted, CustomerCreditLedgerNewEntryByExternalIDResponseIncrementLedgerEntryEntryStatusPending:
		return true
	}
	return false
}

type CustomerCreditLedgerNewEntryByExternalIDResponseIncrementLedgerEntryEntryType string

const (
	CustomerCreditLedgerNewEntryByExternalIDResponseIncrementLedgerEntryEntryTypeIncrement CustomerCreditLedgerNewEntryByExternalIDResponseIncrementLedgerEntryEntryType = "increment"
)

func (r CustomerCreditLedgerNewEntryByExternalIDResponseIncrementLedgerEntryEntryType) IsKnown() bool {
	switch r {
	case CustomerCreditLedgerNewEntryByExternalIDResponseIncrementLedgerEntryEntryTypeIncrement:
		return true
	}
	return false
}

type CustomerCreditLedgerNewEntryByExternalIDResponseDecrementLedgerEntry struct {
	ID                   string                                                                          `json:"id,required"`
	Amount               float64                                                                         `json:"amount,required"`
	CreatedAt            time.Time                                                                       `json:"created_at,required" format:"date-time"`
	CreditBlock          CustomerCreditLedgerNewEntryByExternalIDResponseDecrementLedgerEntryCreditBlock `json:"credit_block,required"`
	Currency             string                                                                          `json:"currency,required"`
	Customer             CustomerCreditLedgerNewEntryByExternalIDResponseDecrementLedgerEntryCustomer    `json:"customer,required"`
	Description          string                                                                          `json:"description,required,nullable"`
	EndingBalance        float64                                                                         `json:"ending_balance,required"`
	EntryStatus          CustomerCreditLedgerNewEntryByExternalIDResponseDecrementLedgerEntryEntryStatus `json:"entry_status,required"`
	EntryType            CustomerCreditLedgerNewEntryByExternalIDResponseDecrementLedgerEntryEntryType   `json:"entry_type,required"`
	LedgerSequenceNumber int64                                                                           `json:"ledger_sequence_number,required"`
	// User specified key-value pairs for the resource. If not present, this defaults
	// to an empty dictionary. Individual keys can be removed by setting the value to
	// `null`, and the entire metadata mapping can be cleared by setting `metadata` to
	// `null`.
	Metadata        map[string]string                                                        `json:"metadata,required"`
	StartingBalance float64                                                                  `json:"starting_balance,required"`
	EventID         string                                                                   `json:"event_id,nullable"`
	InvoiceID       string                                                                   `json:"invoice_id,nullable"`
	PriceID         string                                                                   `json:"price_id,nullable"`
	JSON            customerCreditLedgerNewEntryByExternalIDResponseDecrementLedgerEntryJSON `json:"-"`
}

// customerCreditLedgerNewEntryByExternalIDResponseDecrementLedgerEntryJSON
// contains the JSON metadata for the struct
// [CustomerCreditLedgerNewEntryByExternalIDResponseDecrementLedgerEntry]
type customerCreditLedgerNewEntryByExternalIDResponseDecrementLedgerEntryJSON struct {
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

func (r *CustomerCreditLedgerNewEntryByExternalIDResponseDecrementLedgerEntry) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r customerCreditLedgerNewEntryByExternalIDResponseDecrementLedgerEntryJSON) RawJSON() string {
	return r.raw
}

func (r CustomerCreditLedgerNewEntryByExternalIDResponseDecrementLedgerEntry) implementsCustomerCreditLedgerNewEntryByExternalIDResponse() {
}

type CustomerCreditLedgerNewEntryByExternalIDResponseDecrementLedgerEntryCreditBlock struct {
	ID               string                                                                              `json:"id,required"`
	ExpiryDate       time.Time                                                                           `json:"expiry_date,required,nullable" format:"date-time"`
	PerUnitCostBasis string                                                                              `json:"per_unit_cost_basis,required,nullable"`
	JSON             customerCreditLedgerNewEntryByExternalIDResponseDecrementLedgerEntryCreditBlockJSON `json:"-"`
}

// customerCreditLedgerNewEntryByExternalIDResponseDecrementLedgerEntryCreditBlockJSON
// contains the JSON metadata for the struct
// [CustomerCreditLedgerNewEntryByExternalIDResponseDecrementLedgerEntryCreditBlock]
type customerCreditLedgerNewEntryByExternalIDResponseDecrementLedgerEntryCreditBlockJSON struct {
	ID               apijson.Field
	ExpiryDate       apijson.Field
	PerUnitCostBasis apijson.Field
	raw              string
	ExtraFields      map[string]apijson.Field
}

func (r *CustomerCreditLedgerNewEntryByExternalIDResponseDecrementLedgerEntryCreditBlock) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r customerCreditLedgerNewEntryByExternalIDResponseDecrementLedgerEntryCreditBlockJSON) RawJSON() string {
	return r.raw
}

type CustomerCreditLedgerNewEntryByExternalIDResponseDecrementLedgerEntryCustomer struct {
	ID                 string                                                                           `json:"id,required"`
	ExternalCustomerID string                                                                           `json:"external_customer_id,required,nullable"`
	JSON               customerCreditLedgerNewEntryByExternalIDResponseDecrementLedgerEntryCustomerJSON `json:"-"`
}

// customerCreditLedgerNewEntryByExternalIDResponseDecrementLedgerEntryCustomerJSON
// contains the JSON metadata for the struct
// [CustomerCreditLedgerNewEntryByExternalIDResponseDecrementLedgerEntryCustomer]
type customerCreditLedgerNewEntryByExternalIDResponseDecrementLedgerEntryCustomerJSON struct {
	ID                 apijson.Field
	ExternalCustomerID apijson.Field
	raw                string
	ExtraFields        map[string]apijson.Field
}

func (r *CustomerCreditLedgerNewEntryByExternalIDResponseDecrementLedgerEntryCustomer) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r customerCreditLedgerNewEntryByExternalIDResponseDecrementLedgerEntryCustomerJSON) RawJSON() string {
	return r.raw
}

type CustomerCreditLedgerNewEntryByExternalIDResponseDecrementLedgerEntryEntryStatus string

const (
	CustomerCreditLedgerNewEntryByExternalIDResponseDecrementLedgerEntryEntryStatusCommitted CustomerCreditLedgerNewEntryByExternalIDResponseDecrementLedgerEntryEntryStatus = "committed"
	CustomerCreditLedgerNewEntryByExternalIDResponseDecrementLedgerEntryEntryStatusPending   CustomerCreditLedgerNewEntryByExternalIDResponseDecrementLedgerEntryEntryStatus = "pending"
)

func (r CustomerCreditLedgerNewEntryByExternalIDResponseDecrementLedgerEntryEntryStatus) IsKnown() bool {
	switch r {
	case CustomerCreditLedgerNewEntryByExternalIDResponseDecrementLedgerEntryEntryStatusCommitted, CustomerCreditLedgerNewEntryByExternalIDResponseDecrementLedgerEntryEntryStatusPending:
		return true
	}
	return false
}

type CustomerCreditLedgerNewEntryByExternalIDResponseDecrementLedgerEntryEntryType string

const (
	CustomerCreditLedgerNewEntryByExternalIDResponseDecrementLedgerEntryEntryTypeDecrement CustomerCreditLedgerNewEntryByExternalIDResponseDecrementLedgerEntryEntryType = "decrement"
)

func (r CustomerCreditLedgerNewEntryByExternalIDResponseDecrementLedgerEntryEntryType) IsKnown() bool {
	switch r {
	case CustomerCreditLedgerNewEntryByExternalIDResponseDecrementLedgerEntryEntryTypeDecrement:
		return true
	}
	return false
}

type CustomerCreditLedgerNewEntryByExternalIDResponseExpirationChangeLedgerEntry struct {
	ID                   string                                                                                 `json:"id,required"`
	Amount               float64                                                                                `json:"amount,required"`
	CreatedAt            time.Time                                                                              `json:"created_at,required" format:"date-time"`
	CreditBlock          CustomerCreditLedgerNewEntryByExternalIDResponseExpirationChangeLedgerEntryCreditBlock `json:"credit_block,required"`
	Currency             string                                                                                 `json:"currency,required"`
	Customer             CustomerCreditLedgerNewEntryByExternalIDResponseExpirationChangeLedgerEntryCustomer    `json:"customer,required"`
	Description          string                                                                                 `json:"description,required,nullable"`
	EndingBalance        float64                                                                                `json:"ending_balance,required"`
	EntryStatus          CustomerCreditLedgerNewEntryByExternalIDResponseExpirationChangeLedgerEntryEntryStatus `json:"entry_status,required"`
	EntryType            CustomerCreditLedgerNewEntryByExternalIDResponseExpirationChangeLedgerEntryEntryType   `json:"entry_type,required"`
	LedgerSequenceNumber int64                                                                                  `json:"ledger_sequence_number,required"`
	// User specified key-value pairs for the resource. If not present, this defaults
	// to an empty dictionary. Individual keys can be removed by setting the value to
	// `null`, and the entire metadata mapping can be cleared by setting `metadata` to
	// `null`.
	Metadata           map[string]string                                                               `json:"metadata,required"`
	NewBlockExpiryDate time.Time                                                                       `json:"new_block_expiry_date,required" format:"date-time"`
	StartingBalance    float64                                                                         `json:"starting_balance,required"`
	JSON               customerCreditLedgerNewEntryByExternalIDResponseExpirationChangeLedgerEntryJSON `json:"-"`
}

// customerCreditLedgerNewEntryByExternalIDResponseExpirationChangeLedgerEntryJSON
// contains the JSON metadata for the struct
// [CustomerCreditLedgerNewEntryByExternalIDResponseExpirationChangeLedgerEntry]
type customerCreditLedgerNewEntryByExternalIDResponseExpirationChangeLedgerEntryJSON struct {
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

func (r *CustomerCreditLedgerNewEntryByExternalIDResponseExpirationChangeLedgerEntry) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r customerCreditLedgerNewEntryByExternalIDResponseExpirationChangeLedgerEntryJSON) RawJSON() string {
	return r.raw
}

func (r CustomerCreditLedgerNewEntryByExternalIDResponseExpirationChangeLedgerEntry) implementsCustomerCreditLedgerNewEntryByExternalIDResponse() {
}

type CustomerCreditLedgerNewEntryByExternalIDResponseExpirationChangeLedgerEntryCreditBlock struct {
	ID               string                                                                                     `json:"id,required"`
	ExpiryDate       time.Time                                                                                  `json:"expiry_date,required,nullable" format:"date-time"`
	PerUnitCostBasis string                                                                                     `json:"per_unit_cost_basis,required,nullable"`
	JSON             customerCreditLedgerNewEntryByExternalIDResponseExpirationChangeLedgerEntryCreditBlockJSON `json:"-"`
}

// customerCreditLedgerNewEntryByExternalIDResponseExpirationChangeLedgerEntryCreditBlockJSON
// contains the JSON metadata for the struct
// [CustomerCreditLedgerNewEntryByExternalIDResponseExpirationChangeLedgerEntryCreditBlock]
type customerCreditLedgerNewEntryByExternalIDResponseExpirationChangeLedgerEntryCreditBlockJSON struct {
	ID               apijson.Field
	ExpiryDate       apijson.Field
	PerUnitCostBasis apijson.Field
	raw              string
	ExtraFields      map[string]apijson.Field
}

func (r *CustomerCreditLedgerNewEntryByExternalIDResponseExpirationChangeLedgerEntryCreditBlock) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r customerCreditLedgerNewEntryByExternalIDResponseExpirationChangeLedgerEntryCreditBlockJSON) RawJSON() string {
	return r.raw
}

type CustomerCreditLedgerNewEntryByExternalIDResponseExpirationChangeLedgerEntryCustomer struct {
	ID                 string                                                                                  `json:"id,required"`
	ExternalCustomerID string                                                                                  `json:"external_customer_id,required,nullable"`
	JSON               customerCreditLedgerNewEntryByExternalIDResponseExpirationChangeLedgerEntryCustomerJSON `json:"-"`
}

// customerCreditLedgerNewEntryByExternalIDResponseExpirationChangeLedgerEntryCustomerJSON
// contains the JSON metadata for the struct
// [CustomerCreditLedgerNewEntryByExternalIDResponseExpirationChangeLedgerEntryCustomer]
type customerCreditLedgerNewEntryByExternalIDResponseExpirationChangeLedgerEntryCustomerJSON struct {
	ID                 apijson.Field
	ExternalCustomerID apijson.Field
	raw                string
	ExtraFields        map[string]apijson.Field
}

func (r *CustomerCreditLedgerNewEntryByExternalIDResponseExpirationChangeLedgerEntryCustomer) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r customerCreditLedgerNewEntryByExternalIDResponseExpirationChangeLedgerEntryCustomerJSON) RawJSON() string {
	return r.raw
}

type CustomerCreditLedgerNewEntryByExternalIDResponseExpirationChangeLedgerEntryEntryStatus string

const (
	CustomerCreditLedgerNewEntryByExternalIDResponseExpirationChangeLedgerEntryEntryStatusCommitted CustomerCreditLedgerNewEntryByExternalIDResponseExpirationChangeLedgerEntryEntryStatus = "committed"
	CustomerCreditLedgerNewEntryByExternalIDResponseExpirationChangeLedgerEntryEntryStatusPending   CustomerCreditLedgerNewEntryByExternalIDResponseExpirationChangeLedgerEntryEntryStatus = "pending"
)

func (r CustomerCreditLedgerNewEntryByExternalIDResponseExpirationChangeLedgerEntryEntryStatus) IsKnown() bool {
	switch r {
	case CustomerCreditLedgerNewEntryByExternalIDResponseExpirationChangeLedgerEntryEntryStatusCommitted, CustomerCreditLedgerNewEntryByExternalIDResponseExpirationChangeLedgerEntryEntryStatusPending:
		return true
	}
	return false
}

type CustomerCreditLedgerNewEntryByExternalIDResponseExpirationChangeLedgerEntryEntryType string

const (
	CustomerCreditLedgerNewEntryByExternalIDResponseExpirationChangeLedgerEntryEntryTypeExpirationChange CustomerCreditLedgerNewEntryByExternalIDResponseExpirationChangeLedgerEntryEntryType = "expiration_change"
)

func (r CustomerCreditLedgerNewEntryByExternalIDResponseExpirationChangeLedgerEntryEntryType) IsKnown() bool {
	switch r {
	case CustomerCreditLedgerNewEntryByExternalIDResponseExpirationChangeLedgerEntryEntryTypeExpirationChange:
		return true
	}
	return false
}

type CustomerCreditLedgerNewEntryByExternalIDResponseCreditBlockExpiryLedgerEntry struct {
	ID                   string                                                                                  `json:"id,required"`
	Amount               float64                                                                                 `json:"amount,required"`
	CreatedAt            time.Time                                                                               `json:"created_at,required" format:"date-time"`
	CreditBlock          CustomerCreditLedgerNewEntryByExternalIDResponseCreditBlockExpiryLedgerEntryCreditBlock `json:"credit_block,required"`
	Currency             string                                                                                  `json:"currency,required"`
	Customer             CustomerCreditLedgerNewEntryByExternalIDResponseCreditBlockExpiryLedgerEntryCustomer    `json:"customer,required"`
	Description          string                                                                                  `json:"description,required,nullable"`
	EndingBalance        float64                                                                                 `json:"ending_balance,required"`
	EntryStatus          CustomerCreditLedgerNewEntryByExternalIDResponseCreditBlockExpiryLedgerEntryEntryStatus `json:"entry_status,required"`
	EntryType            CustomerCreditLedgerNewEntryByExternalIDResponseCreditBlockExpiryLedgerEntryEntryType   `json:"entry_type,required"`
	LedgerSequenceNumber int64                                                                                   `json:"ledger_sequence_number,required"`
	// User specified key-value pairs for the resource. If not present, this defaults
	// to an empty dictionary. Individual keys can be removed by setting the value to
	// `null`, and the entire metadata mapping can be cleared by setting `metadata` to
	// `null`.
	Metadata        map[string]string                                                                `json:"metadata,required"`
	StartingBalance float64                                                                          `json:"starting_balance,required"`
	JSON            customerCreditLedgerNewEntryByExternalIDResponseCreditBlockExpiryLedgerEntryJSON `json:"-"`
}

// customerCreditLedgerNewEntryByExternalIDResponseCreditBlockExpiryLedgerEntryJSON
// contains the JSON metadata for the struct
// [CustomerCreditLedgerNewEntryByExternalIDResponseCreditBlockExpiryLedgerEntry]
type customerCreditLedgerNewEntryByExternalIDResponseCreditBlockExpiryLedgerEntryJSON struct {
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

func (r *CustomerCreditLedgerNewEntryByExternalIDResponseCreditBlockExpiryLedgerEntry) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r customerCreditLedgerNewEntryByExternalIDResponseCreditBlockExpiryLedgerEntryJSON) RawJSON() string {
	return r.raw
}

func (r CustomerCreditLedgerNewEntryByExternalIDResponseCreditBlockExpiryLedgerEntry) implementsCustomerCreditLedgerNewEntryByExternalIDResponse() {
}

type CustomerCreditLedgerNewEntryByExternalIDResponseCreditBlockExpiryLedgerEntryCreditBlock struct {
	ID               string                                                                                      `json:"id,required"`
	ExpiryDate       time.Time                                                                                   `json:"expiry_date,required,nullable" format:"date-time"`
	PerUnitCostBasis string                                                                                      `json:"per_unit_cost_basis,required,nullable"`
	JSON             customerCreditLedgerNewEntryByExternalIDResponseCreditBlockExpiryLedgerEntryCreditBlockJSON `json:"-"`
}

// customerCreditLedgerNewEntryByExternalIDResponseCreditBlockExpiryLedgerEntryCreditBlockJSON
// contains the JSON metadata for the struct
// [CustomerCreditLedgerNewEntryByExternalIDResponseCreditBlockExpiryLedgerEntryCreditBlock]
type customerCreditLedgerNewEntryByExternalIDResponseCreditBlockExpiryLedgerEntryCreditBlockJSON struct {
	ID               apijson.Field
	ExpiryDate       apijson.Field
	PerUnitCostBasis apijson.Field
	raw              string
	ExtraFields      map[string]apijson.Field
}

func (r *CustomerCreditLedgerNewEntryByExternalIDResponseCreditBlockExpiryLedgerEntryCreditBlock) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r customerCreditLedgerNewEntryByExternalIDResponseCreditBlockExpiryLedgerEntryCreditBlockJSON) RawJSON() string {
	return r.raw
}

type CustomerCreditLedgerNewEntryByExternalIDResponseCreditBlockExpiryLedgerEntryCustomer struct {
	ID                 string                                                                                   `json:"id,required"`
	ExternalCustomerID string                                                                                   `json:"external_customer_id,required,nullable"`
	JSON               customerCreditLedgerNewEntryByExternalIDResponseCreditBlockExpiryLedgerEntryCustomerJSON `json:"-"`
}

// customerCreditLedgerNewEntryByExternalIDResponseCreditBlockExpiryLedgerEntryCustomerJSON
// contains the JSON metadata for the struct
// [CustomerCreditLedgerNewEntryByExternalIDResponseCreditBlockExpiryLedgerEntryCustomer]
type customerCreditLedgerNewEntryByExternalIDResponseCreditBlockExpiryLedgerEntryCustomerJSON struct {
	ID                 apijson.Field
	ExternalCustomerID apijson.Field
	raw                string
	ExtraFields        map[string]apijson.Field
}

func (r *CustomerCreditLedgerNewEntryByExternalIDResponseCreditBlockExpiryLedgerEntryCustomer) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r customerCreditLedgerNewEntryByExternalIDResponseCreditBlockExpiryLedgerEntryCustomerJSON) RawJSON() string {
	return r.raw
}

type CustomerCreditLedgerNewEntryByExternalIDResponseCreditBlockExpiryLedgerEntryEntryStatus string

const (
	CustomerCreditLedgerNewEntryByExternalIDResponseCreditBlockExpiryLedgerEntryEntryStatusCommitted CustomerCreditLedgerNewEntryByExternalIDResponseCreditBlockExpiryLedgerEntryEntryStatus = "committed"
	CustomerCreditLedgerNewEntryByExternalIDResponseCreditBlockExpiryLedgerEntryEntryStatusPending   CustomerCreditLedgerNewEntryByExternalIDResponseCreditBlockExpiryLedgerEntryEntryStatus = "pending"
)

func (r CustomerCreditLedgerNewEntryByExternalIDResponseCreditBlockExpiryLedgerEntryEntryStatus) IsKnown() bool {
	switch r {
	case CustomerCreditLedgerNewEntryByExternalIDResponseCreditBlockExpiryLedgerEntryEntryStatusCommitted, CustomerCreditLedgerNewEntryByExternalIDResponseCreditBlockExpiryLedgerEntryEntryStatusPending:
		return true
	}
	return false
}

type CustomerCreditLedgerNewEntryByExternalIDResponseCreditBlockExpiryLedgerEntryEntryType string

const (
	CustomerCreditLedgerNewEntryByExternalIDResponseCreditBlockExpiryLedgerEntryEntryTypeCreditBlockExpiry CustomerCreditLedgerNewEntryByExternalIDResponseCreditBlockExpiryLedgerEntryEntryType = "credit_block_expiry"
)

func (r CustomerCreditLedgerNewEntryByExternalIDResponseCreditBlockExpiryLedgerEntryEntryType) IsKnown() bool {
	switch r {
	case CustomerCreditLedgerNewEntryByExternalIDResponseCreditBlockExpiryLedgerEntryEntryTypeCreditBlockExpiry:
		return true
	}
	return false
}

type CustomerCreditLedgerNewEntryByExternalIDResponseVoidLedgerEntry struct {
	ID                   string                                                                     `json:"id,required"`
	Amount               float64                                                                    `json:"amount,required"`
	CreatedAt            time.Time                                                                  `json:"created_at,required" format:"date-time"`
	CreditBlock          CustomerCreditLedgerNewEntryByExternalIDResponseVoidLedgerEntryCreditBlock `json:"credit_block,required"`
	Currency             string                                                                     `json:"currency,required"`
	Customer             CustomerCreditLedgerNewEntryByExternalIDResponseVoidLedgerEntryCustomer    `json:"customer,required"`
	Description          string                                                                     `json:"description,required,nullable"`
	EndingBalance        float64                                                                    `json:"ending_balance,required"`
	EntryStatus          CustomerCreditLedgerNewEntryByExternalIDResponseVoidLedgerEntryEntryStatus `json:"entry_status,required"`
	EntryType            CustomerCreditLedgerNewEntryByExternalIDResponseVoidLedgerEntryEntryType   `json:"entry_type,required"`
	LedgerSequenceNumber int64                                                                      `json:"ledger_sequence_number,required"`
	// User specified key-value pairs for the resource. If not present, this defaults
	// to an empty dictionary. Individual keys can be removed by setting the value to
	// `null`, and the entire metadata mapping can be cleared by setting `metadata` to
	// `null`.
	Metadata        map[string]string                                                   `json:"metadata,required"`
	StartingBalance float64                                                             `json:"starting_balance,required"`
	VoidAmount      float64                                                             `json:"void_amount,required"`
	VoidReason      string                                                              `json:"void_reason,required,nullable"`
	JSON            customerCreditLedgerNewEntryByExternalIDResponseVoidLedgerEntryJSON `json:"-"`
}

// customerCreditLedgerNewEntryByExternalIDResponseVoidLedgerEntryJSON contains the
// JSON metadata for the struct
// [CustomerCreditLedgerNewEntryByExternalIDResponseVoidLedgerEntry]
type customerCreditLedgerNewEntryByExternalIDResponseVoidLedgerEntryJSON struct {
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

func (r *CustomerCreditLedgerNewEntryByExternalIDResponseVoidLedgerEntry) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r customerCreditLedgerNewEntryByExternalIDResponseVoidLedgerEntryJSON) RawJSON() string {
	return r.raw
}

func (r CustomerCreditLedgerNewEntryByExternalIDResponseVoidLedgerEntry) implementsCustomerCreditLedgerNewEntryByExternalIDResponse() {
}

type CustomerCreditLedgerNewEntryByExternalIDResponseVoidLedgerEntryCreditBlock struct {
	ID               string                                                                         `json:"id,required"`
	ExpiryDate       time.Time                                                                      `json:"expiry_date,required,nullable" format:"date-time"`
	PerUnitCostBasis string                                                                         `json:"per_unit_cost_basis,required,nullable"`
	JSON             customerCreditLedgerNewEntryByExternalIDResponseVoidLedgerEntryCreditBlockJSON `json:"-"`
}

// customerCreditLedgerNewEntryByExternalIDResponseVoidLedgerEntryCreditBlockJSON
// contains the JSON metadata for the struct
// [CustomerCreditLedgerNewEntryByExternalIDResponseVoidLedgerEntryCreditBlock]
type customerCreditLedgerNewEntryByExternalIDResponseVoidLedgerEntryCreditBlockJSON struct {
	ID               apijson.Field
	ExpiryDate       apijson.Field
	PerUnitCostBasis apijson.Field
	raw              string
	ExtraFields      map[string]apijson.Field
}

func (r *CustomerCreditLedgerNewEntryByExternalIDResponseVoidLedgerEntryCreditBlock) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r customerCreditLedgerNewEntryByExternalIDResponseVoidLedgerEntryCreditBlockJSON) RawJSON() string {
	return r.raw
}

type CustomerCreditLedgerNewEntryByExternalIDResponseVoidLedgerEntryCustomer struct {
	ID                 string                                                                      `json:"id,required"`
	ExternalCustomerID string                                                                      `json:"external_customer_id,required,nullable"`
	JSON               customerCreditLedgerNewEntryByExternalIDResponseVoidLedgerEntryCustomerJSON `json:"-"`
}

// customerCreditLedgerNewEntryByExternalIDResponseVoidLedgerEntryCustomerJSON
// contains the JSON metadata for the struct
// [CustomerCreditLedgerNewEntryByExternalIDResponseVoidLedgerEntryCustomer]
type customerCreditLedgerNewEntryByExternalIDResponseVoidLedgerEntryCustomerJSON struct {
	ID                 apijson.Field
	ExternalCustomerID apijson.Field
	raw                string
	ExtraFields        map[string]apijson.Field
}

func (r *CustomerCreditLedgerNewEntryByExternalIDResponseVoidLedgerEntryCustomer) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r customerCreditLedgerNewEntryByExternalIDResponseVoidLedgerEntryCustomerJSON) RawJSON() string {
	return r.raw
}

type CustomerCreditLedgerNewEntryByExternalIDResponseVoidLedgerEntryEntryStatus string

const (
	CustomerCreditLedgerNewEntryByExternalIDResponseVoidLedgerEntryEntryStatusCommitted CustomerCreditLedgerNewEntryByExternalIDResponseVoidLedgerEntryEntryStatus = "committed"
	CustomerCreditLedgerNewEntryByExternalIDResponseVoidLedgerEntryEntryStatusPending   CustomerCreditLedgerNewEntryByExternalIDResponseVoidLedgerEntryEntryStatus = "pending"
)

func (r CustomerCreditLedgerNewEntryByExternalIDResponseVoidLedgerEntryEntryStatus) IsKnown() bool {
	switch r {
	case CustomerCreditLedgerNewEntryByExternalIDResponseVoidLedgerEntryEntryStatusCommitted, CustomerCreditLedgerNewEntryByExternalIDResponseVoidLedgerEntryEntryStatusPending:
		return true
	}
	return false
}

type CustomerCreditLedgerNewEntryByExternalIDResponseVoidLedgerEntryEntryType string

const (
	CustomerCreditLedgerNewEntryByExternalIDResponseVoidLedgerEntryEntryTypeVoid CustomerCreditLedgerNewEntryByExternalIDResponseVoidLedgerEntryEntryType = "void"
)

func (r CustomerCreditLedgerNewEntryByExternalIDResponseVoidLedgerEntryEntryType) IsKnown() bool {
	switch r {
	case CustomerCreditLedgerNewEntryByExternalIDResponseVoidLedgerEntryEntryTypeVoid:
		return true
	}
	return false
}

type CustomerCreditLedgerNewEntryByExternalIDResponseVoidInitiatedLedgerEntry struct {
	ID                   string                                                                              `json:"id,required"`
	Amount               float64                                                                             `json:"amount,required"`
	CreatedAt            time.Time                                                                           `json:"created_at,required" format:"date-time"`
	CreditBlock          CustomerCreditLedgerNewEntryByExternalIDResponseVoidInitiatedLedgerEntryCreditBlock `json:"credit_block,required"`
	Currency             string                                                                              `json:"currency,required"`
	Customer             CustomerCreditLedgerNewEntryByExternalIDResponseVoidInitiatedLedgerEntryCustomer    `json:"customer,required"`
	Description          string                                                                              `json:"description,required,nullable"`
	EndingBalance        float64                                                                             `json:"ending_balance,required"`
	EntryStatus          CustomerCreditLedgerNewEntryByExternalIDResponseVoidInitiatedLedgerEntryEntryStatus `json:"entry_status,required"`
	EntryType            CustomerCreditLedgerNewEntryByExternalIDResponseVoidInitiatedLedgerEntryEntryType   `json:"entry_type,required"`
	LedgerSequenceNumber int64                                                                               `json:"ledger_sequence_number,required"`
	// User specified key-value pairs for the resource. If not present, this defaults
	// to an empty dictionary. Individual keys can be removed by setting the value to
	// `null`, and the entire metadata mapping can be cleared by setting `metadata` to
	// `null`.
	Metadata           map[string]string                                                            `json:"metadata,required"`
	NewBlockExpiryDate time.Time                                                                    `json:"new_block_expiry_date,required" format:"date-time"`
	StartingBalance    float64                                                                      `json:"starting_balance,required"`
	VoidAmount         float64                                                                      `json:"void_amount,required"`
	VoidReason         string                                                                       `json:"void_reason,required,nullable"`
	JSON               customerCreditLedgerNewEntryByExternalIDResponseVoidInitiatedLedgerEntryJSON `json:"-"`
}

// customerCreditLedgerNewEntryByExternalIDResponseVoidInitiatedLedgerEntryJSON
// contains the JSON metadata for the struct
// [CustomerCreditLedgerNewEntryByExternalIDResponseVoidInitiatedLedgerEntry]
type customerCreditLedgerNewEntryByExternalIDResponseVoidInitiatedLedgerEntryJSON struct {
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

func (r *CustomerCreditLedgerNewEntryByExternalIDResponseVoidInitiatedLedgerEntry) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r customerCreditLedgerNewEntryByExternalIDResponseVoidInitiatedLedgerEntryJSON) RawJSON() string {
	return r.raw
}

func (r CustomerCreditLedgerNewEntryByExternalIDResponseVoidInitiatedLedgerEntry) implementsCustomerCreditLedgerNewEntryByExternalIDResponse() {
}

type CustomerCreditLedgerNewEntryByExternalIDResponseVoidInitiatedLedgerEntryCreditBlock struct {
	ID               string                                                                                  `json:"id,required"`
	ExpiryDate       time.Time                                                                               `json:"expiry_date,required,nullable" format:"date-time"`
	PerUnitCostBasis string                                                                                  `json:"per_unit_cost_basis,required,nullable"`
	JSON             customerCreditLedgerNewEntryByExternalIDResponseVoidInitiatedLedgerEntryCreditBlockJSON `json:"-"`
}

// customerCreditLedgerNewEntryByExternalIDResponseVoidInitiatedLedgerEntryCreditBlockJSON
// contains the JSON metadata for the struct
// [CustomerCreditLedgerNewEntryByExternalIDResponseVoidInitiatedLedgerEntryCreditBlock]
type customerCreditLedgerNewEntryByExternalIDResponseVoidInitiatedLedgerEntryCreditBlockJSON struct {
	ID               apijson.Field
	ExpiryDate       apijson.Field
	PerUnitCostBasis apijson.Field
	raw              string
	ExtraFields      map[string]apijson.Field
}

func (r *CustomerCreditLedgerNewEntryByExternalIDResponseVoidInitiatedLedgerEntryCreditBlock) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r customerCreditLedgerNewEntryByExternalIDResponseVoidInitiatedLedgerEntryCreditBlockJSON) RawJSON() string {
	return r.raw
}

type CustomerCreditLedgerNewEntryByExternalIDResponseVoidInitiatedLedgerEntryCustomer struct {
	ID                 string                                                                               `json:"id,required"`
	ExternalCustomerID string                                                                               `json:"external_customer_id,required,nullable"`
	JSON               customerCreditLedgerNewEntryByExternalIDResponseVoidInitiatedLedgerEntryCustomerJSON `json:"-"`
}

// customerCreditLedgerNewEntryByExternalIDResponseVoidInitiatedLedgerEntryCustomerJSON
// contains the JSON metadata for the struct
// [CustomerCreditLedgerNewEntryByExternalIDResponseVoidInitiatedLedgerEntryCustomer]
type customerCreditLedgerNewEntryByExternalIDResponseVoidInitiatedLedgerEntryCustomerJSON struct {
	ID                 apijson.Field
	ExternalCustomerID apijson.Field
	raw                string
	ExtraFields        map[string]apijson.Field
}

func (r *CustomerCreditLedgerNewEntryByExternalIDResponseVoidInitiatedLedgerEntryCustomer) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r customerCreditLedgerNewEntryByExternalIDResponseVoidInitiatedLedgerEntryCustomerJSON) RawJSON() string {
	return r.raw
}

type CustomerCreditLedgerNewEntryByExternalIDResponseVoidInitiatedLedgerEntryEntryStatus string

const (
	CustomerCreditLedgerNewEntryByExternalIDResponseVoidInitiatedLedgerEntryEntryStatusCommitted CustomerCreditLedgerNewEntryByExternalIDResponseVoidInitiatedLedgerEntryEntryStatus = "committed"
	CustomerCreditLedgerNewEntryByExternalIDResponseVoidInitiatedLedgerEntryEntryStatusPending   CustomerCreditLedgerNewEntryByExternalIDResponseVoidInitiatedLedgerEntryEntryStatus = "pending"
)

func (r CustomerCreditLedgerNewEntryByExternalIDResponseVoidInitiatedLedgerEntryEntryStatus) IsKnown() bool {
	switch r {
	case CustomerCreditLedgerNewEntryByExternalIDResponseVoidInitiatedLedgerEntryEntryStatusCommitted, CustomerCreditLedgerNewEntryByExternalIDResponseVoidInitiatedLedgerEntryEntryStatusPending:
		return true
	}
	return false
}

type CustomerCreditLedgerNewEntryByExternalIDResponseVoidInitiatedLedgerEntryEntryType string

const (
	CustomerCreditLedgerNewEntryByExternalIDResponseVoidInitiatedLedgerEntryEntryTypeVoidInitiated CustomerCreditLedgerNewEntryByExternalIDResponseVoidInitiatedLedgerEntryEntryType = "void_initiated"
)

func (r CustomerCreditLedgerNewEntryByExternalIDResponseVoidInitiatedLedgerEntryEntryType) IsKnown() bool {
	switch r {
	case CustomerCreditLedgerNewEntryByExternalIDResponseVoidInitiatedLedgerEntryEntryTypeVoidInitiated:
		return true
	}
	return false
}

type CustomerCreditLedgerNewEntryByExternalIDResponseAmendmentLedgerEntry struct {
	ID                   string                                                                          `json:"id,required"`
	Amount               float64                                                                         `json:"amount,required"`
	CreatedAt            time.Time                                                                       `json:"created_at,required" format:"date-time"`
	CreditBlock          CustomerCreditLedgerNewEntryByExternalIDResponseAmendmentLedgerEntryCreditBlock `json:"credit_block,required"`
	Currency             string                                                                          `json:"currency,required"`
	Customer             CustomerCreditLedgerNewEntryByExternalIDResponseAmendmentLedgerEntryCustomer    `json:"customer,required"`
	Description          string                                                                          `json:"description,required,nullable"`
	EndingBalance        float64                                                                         `json:"ending_balance,required"`
	EntryStatus          CustomerCreditLedgerNewEntryByExternalIDResponseAmendmentLedgerEntryEntryStatus `json:"entry_status,required"`
	EntryType            CustomerCreditLedgerNewEntryByExternalIDResponseAmendmentLedgerEntryEntryType   `json:"entry_type,required"`
	LedgerSequenceNumber int64                                                                           `json:"ledger_sequence_number,required"`
	// User specified key-value pairs for the resource. If not present, this defaults
	// to an empty dictionary. Individual keys can be removed by setting the value to
	// `null`, and the entire metadata mapping can be cleared by setting `metadata` to
	// `null`.
	Metadata        map[string]string                                                        `json:"metadata,required"`
	StartingBalance float64                                                                  `json:"starting_balance,required"`
	JSON            customerCreditLedgerNewEntryByExternalIDResponseAmendmentLedgerEntryJSON `json:"-"`
}

// customerCreditLedgerNewEntryByExternalIDResponseAmendmentLedgerEntryJSON
// contains the JSON metadata for the struct
// [CustomerCreditLedgerNewEntryByExternalIDResponseAmendmentLedgerEntry]
type customerCreditLedgerNewEntryByExternalIDResponseAmendmentLedgerEntryJSON struct {
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

func (r *CustomerCreditLedgerNewEntryByExternalIDResponseAmendmentLedgerEntry) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r customerCreditLedgerNewEntryByExternalIDResponseAmendmentLedgerEntryJSON) RawJSON() string {
	return r.raw
}

func (r CustomerCreditLedgerNewEntryByExternalIDResponseAmendmentLedgerEntry) implementsCustomerCreditLedgerNewEntryByExternalIDResponse() {
}

type CustomerCreditLedgerNewEntryByExternalIDResponseAmendmentLedgerEntryCreditBlock struct {
	ID               string                                                                              `json:"id,required"`
	ExpiryDate       time.Time                                                                           `json:"expiry_date,required,nullable" format:"date-time"`
	PerUnitCostBasis string                                                                              `json:"per_unit_cost_basis,required,nullable"`
	JSON             customerCreditLedgerNewEntryByExternalIDResponseAmendmentLedgerEntryCreditBlockJSON `json:"-"`
}

// customerCreditLedgerNewEntryByExternalIDResponseAmendmentLedgerEntryCreditBlockJSON
// contains the JSON metadata for the struct
// [CustomerCreditLedgerNewEntryByExternalIDResponseAmendmentLedgerEntryCreditBlock]
type customerCreditLedgerNewEntryByExternalIDResponseAmendmentLedgerEntryCreditBlockJSON struct {
	ID               apijson.Field
	ExpiryDate       apijson.Field
	PerUnitCostBasis apijson.Field
	raw              string
	ExtraFields      map[string]apijson.Field
}

func (r *CustomerCreditLedgerNewEntryByExternalIDResponseAmendmentLedgerEntryCreditBlock) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r customerCreditLedgerNewEntryByExternalIDResponseAmendmentLedgerEntryCreditBlockJSON) RawJSON() string {
	return r.raw
}

type CustomerCreditLedgerNewEntryByExternalIDResponseAmendmentLedgerEntryCustomer struct {
	ID                 string                                                                           `json:"id,required"`
	ExternalCustomerID string                                                                           `json:"external_customer_id,required,nullable"`
	JSON               customerCreditLedgerNewEntryByExternalIDResponseAmendmentLedgerEntryCustomerJSON `json:"-"`
}

// customerCreditLedgerNewEntryByExternalIDResponseAmendmentLedgerEntryCustomerJSON
// contains the JSON metadata for the struct
// [CustomerCreditLedgerNewEntryByExternalIDResponseAmendmentLedgerEntryCustomer]
type customerCreditLedgerNewEntryByExternalIDResponseAmendmentLedgerEntryCustomerJSON struct {
	ID                 apijson.Field
	ExternalCustomerID apijson.Field
	raw                string
	ExtraFields        map[string]apijson.Field
}

func (r *CustomerCreditLedgerNewEntryByExternalIDResponseAmendmentLedgerEntryCustomer) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r customerCreditLedgerNewEntryByExternalIDResponseAmendmentLedgerEntryCustomerJSON) RawJSON() string {
	return r.raw
}

type CustomerCreditLedgerNewEntryByExternalIDResponseAmendmentLedgerEntryEntryStatus string

const (
	CustomerCreditLedgerNewEntryByExternalIDResponseAmendmentLedgerEntryEntryStatusCommitted CustomerCreditLedgerNewEntryByExternalIDResponseAmendmentLedgerEntryEntryStatus = "committed"
	CustomerCreditLedgerNewEntryByExternalIDResponseAmendmentLedgerEntryEntryStatusPending   CustomerCreditLedgerNewEntryByExternalIDResponseAmendmentLedgerEntryEntryStatus = "pending"
)

func (r CustomerCreditLedgerNewEntryByExternalIDResponseAmendmentLedgerEntryEntryStatus) IsKnown() bool {
	switch r {
	case CustomerCreditLedgerNewEntryByExternalIDResponseAmendmentLedgerEntryEntryStatusCommitted, CustomerCreditLedgerNewEntryByExternalIDResponseAmendmentLedgerEntryEntryStatusPending:
		return true
	}
	return false
}

type CustomerCreditLedgerNewEntryByExternalIDResponseAmendmentLedgerEntryEntryType string

const (
	CustomerCreditLedgerNewEntryByExternalIDResponseAmendmentLedgerEntryEntryTypeAmendment CustomerCreditLedgerNewEntryByExternalIDResponseAmendmentLedgerEntryEntryType = "amendment"
)

func (r CustomerCreditLedgerNewEntryByExternalIDResponseAmendmentLedgerEntryEntryType) IsKnown() bool {
	switch r {
	case CustomerCreditLedgerNewEntryByExternalIDResponseAmendmentLedgerEntryEntryTypeAmendment:
		return true
	}
	return false
}

type CustomerCreditLedgerNewEntryByExternalIDResponseEntryStatus string

const (
	CustomerCreditLedgerNewEntryByExternalIDResponseEntryStatusCommitted CustomerCreditLedgerNewEntryByExternalIDResponseEntryStatus = "committed"
	CustomerCreditLedgerNewEntryByExternalIDResponseEntryStatusPending   CustomerCreditLedgerNewEntryByExternalIDResponseEntryStatus = "pending"
)

func (r CustomerCreditLedgerNewEntryByExternalIDResponseEntryStatus) IsKnown() bool {
	switch r {
	case CustomerCreditLedgerNewEntryByExternalIDResponseEntryStatusCommitted, CustomerCreditLedgerNewEntryByExternalIDResponseEntryStatusPending:
		return true
	}
	return false
}

type CustomerCreditLedgerNewEntryByExternalIDResponseEntryType string

const (
	CustomerCreditLedgerNewEntryByExternalIDResponseEntryTypeIncrement         CustomerCreditLedgerNewEntryByExternalIDResponseEntryType = "increment"
	CustomerCreditLedgerNewEntryByExternalIDResponseEntryTypeDecrement         CustomerCreditLedgerNewEntryByExternalIDResponseEntryType = "decrement"
	CustomerCreditLedgerNewEntryByExternalIDResponseEntryTypeExpirationChange  CustomerCreditLedgerNewEntryByExternalIDResponseEntryType = "expiration_change"
	CustomerCreditLedgerNewEntryByExternalIDResponseEntryTypeCreditBlockExpiry CustomerCreditLedgerNewEntryByExternalIDResponseEntryType = "credit_block_expiry"
	CustomerCreditLedgerNewEntryByExternalIDResponseEntryTypeVoid              CustomerCreditLedgerNewEntryByExternalIDResponseEntryType = "void"
	CustomerCreditLedgerNewEntryByExternalIDResponseEntryTypeVoidInitiated     CustomerCreditLedgerNewEntryByExternalIDResponseEntryType = "void_initiated"
	CustomerCreditLedgerNewEntryByExternalIDResponseEntryTypeAmendment         CustomerCreditLedgerNewEntryByExternalIDResponseEntryType = "amendment"
)

func (r CustomerCreditLedgerNewEntryByExternalIDResponseEntryType) IsKnown() bool {
	switch r {
	case CustomerCreditLedgerNewEntryByExternalIDResponseEntryTypeIncrement, CustomerCreditLedgerNewEntryByExternalIDResponseEntryTypeDecrement, CustomerCreditLedgerNewEntryByExternalIDResponseEntryTypeExpirationChange, CustomerCreditLedgerNewEntryByExternalIDResponseEntryTypeCreditBlockExpiry, CustomerCreditLedgerNewEntryByExternalIDResponseEntryTypeVoid, CustomerCreditLedgerNewEntryByExternalIDResponseEntryTypeVoidInitiated, CustomerCreditLedgerNewEntryByExternalIDResponseEntryTypeAmendment:
		return true
	}
	return false
}

// The [Credit Ledger Entry resource](/guides/product-catalog/prepurchase) models
// prepaid credits within Orb.
type CustomerCreditLedgerListByExternalIDResponse struct {
	// This field can have the runtime type of [map[string]string].
	Metadata             interface{}                                             `json:"metadata"`
	ID                   string                                                  `json:"id,required"`
	LedgerSequenceNumber int64                                                   `json:"ledger_sequence_number,required"`
	EntryStatus          CustomerCreditLedgerListByExternalIDResponseEntryStatus `json:"entry_status,required"`
	// This field can have the runtime type of
	// [CustomerCreditLedgerListByExternalIDResponseIncrementLedgerEntryCustomer],
	// [CustomerCreditLedgerListByExternalIDResponseDecrementLedgerEntryCustomer],
	// [CustomerCreditLedgerListByExternalIDResponseExpirationChangeLedgerEntryCustomer],
	// [CustomerCreditLedgerListByExternalIDResponseCreditBlockExpiryLedgerEntryCustomer],
	// [CustomerCreditLedgerListByExternalIDResponseVoidLedgerEntryCustomer],
	// [CustomerCreditLedgerListByExternalIDResponseVoidInitiatedLedgerEntryCustomer],
	// [CustomerCreditLedgerListByExternalIDResponseAmendmentLedgerEntryCustomer].
	Customer        interface{} `json:"customer"`
	StartingBalance float64     `json:"starting_balance,required"`
	EndingBalance   float64     `json:"ending_balance,required"`
	Amount          float64     `json:"amount,required"`
	Currency        string      `json:"currency,required"`
	CreatedAt       time.Time   `json:"created_at,required" format:"date-time"`
	Description     string      `json:"description,required,nullable"`
	// This field can have the runtime type of
	// [CustomerCreditLedgerListByExternalIDResponseIncrementLedgerEntryCreditBlock],
	// [CustomerCreditLedgerListByExternalIDResponseDecrementLedgerEntryCreditBlock],
	// [CustomerCreditLedgerListByExternalIDResponseExpirationChangeLedgerEntryCreditBlock],
	// [CustomerCreditLedgerListByExternalIDResponseCreditBlockExpiryLedgerEntryCreditBlock],
	// [CustomerCreditLedgerListByExternalIDResponseVoidLedgerEntryCreditBlock],
	// [CustomerCreditLedgerListByExternalIDResponseVoidInitiatedLedgerEntryCreditBlock],
	// [CustomerCreditLedgerListByExternalIDResponseAmendmentLedgerEntryCreditBlock].
	CreditBlock        interface{}                                           `json:"credit_block"`
	EntryType          CustomerCreditLedgerListByExternalIDResponseEntryType `json:"entry_type,required"`
	PriceID            string                                                `json:"price_id,nullable"`
	EventID            string                                                `json:"event_id,nullable"`
	InvoiceID          string                                                `json:"invoice_id,nullable"`
	NewBlockExpiryDate time.Time                                             `json:"new_block_expiry_date" format:"date-time"`
	VoidReason         string                                                `json:"void_reason,nullable"`
	VoidAmount         float64                                               `json:"void_amount"`
	JSON               customerCreditLedgerListByExternalIDResponseJSON      `json:"-"`
	union              CustomerCreditLedgerListByExternalIDResponseUnion
}

// customerCreditLedgerListByExternalIDResponseJSON contains the JSON metadata for
// the struct [CustomerCreditLedgerListByExternalIDResponse]
type customerCreditLedgerListByExternalIDResponseJSON struct {
	Metadata             apijson.Field
	ID                   apijson.Field
	LedgerSequenceNumber apijson.Field
	EntryStatus          apijson.Field
	Customer             apijson.Field
	StartingBalance      apijson.Field
	EndingBalance        apijson.Field
	Amount               apijson.Field
	Currency             apijson.Field
	CreatedAt            apijson.Field
	Description          apijson.Field
	CreditBlock          apijson.Field
	EntryType            apijson.Field
	PriceID              apijson.Field
	EventID              apijson.Field
	InvoiceID            apijson.Field
	NewBlockExpiryDate   apijson.Field
	VoidReason           apijson.Field
	VoidAmount           apijson.Field
	raw                  string
	ExtraFields          map[string]apijson.Field
}

func (r customerCreditLedgerListByExternalIDResponseJSON) RawJSON() string {
	return r.raw
}

func (r *CustomerCreditLedgerListByExternalIDResponse) UnmarshalJSON(data []byte) (err error) {
	*r = CustomerCreditLedgerListByExternalIDResponse{}
	err = apijson.UnmarshalRoot(data, &r.union)
	if err != nil {
		return err
	}
	return apijson.Port(r.union, &r)
}

// AsUnion returns a [CustomerCreditLedgerListByExternalIDResponseUnion] interface
// which you can cast to the specific types for more type safety.
//
// Possible runtime types of the union are
// [CustomerCreditLedgerListByExternalIDResponseIncrementLedgerEntry],
// [CustomerCreditLedgerListByExternalIDResponseDecrementLedgerEntry],
// [CustomerCreditLedgerListByExternalIDResponseExpirationChangeLedgerEntry],
// [CustomerCreditLedgerListByExternalIDResponseCreditBlockExpiryLedgerEntry],
// [CustomerCreditLedgerListByExternalIDResponseVoidLedgerEntry],
// [CustomerCreditLedgerListByExternalIDResponseVoidInitiatedLedgerEntry],
// [CustomerCreditLedgerListByExternalIDResponseAmendmentLedgerEntry].
func (r CustomerCreditLedgerListByExternalIDResponse) AsUnion() CustomerCreditLedgerListByExternalIDResponseUnion {
	return r.union
}

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
type CustomerCreditLedgerListByExternalIDResponseUnion interface {
	implementsCustomerCreditLedgerListByExternalIDResponse()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*CustomerCreditLedgerListByExternalIDResponseUnion)(nil)).Elem(),
		"entry_type",
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(CustomerCreditLedgerListByExternalIDResponseIncrementLedgerEntry{}),
			DiscriminatorValue: "increment",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(CustomerCreditLedgerListByExternalIDResponseDecrementLedgerEntry{}),
			DiscriminatorValue: "decrement",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(CustomerCreditLedgerListByExternalIDResponseExpirationChangeLedgerEntry{}),
			DiscriminatorValue: "expiration_change",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(CustomerCreditLedgerListByExternalIDResponseCreditBlockExpiryLedgerEntry{}),
			DiscriminatorValue: "credit_block_expiry",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(CustomerCreditLedgerListByExternalIDResponseVoidLedgerEntry{}),
			DiscriminatorValue: "void",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(CustomerCreditLedgerListByExternalIDResponseVoidInitiatedLedgerEntry{}),
			DiscriminatorValue: "void_initiated",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(CustomerCreditLedgerListByExternalIDResponseAmendmentLedgerEntry{}),
			DiscriminatorValue: "amendment",
		},
	)
}

type CustomerCreditLedgerListByExternalIDResponseIncrementLedgerEntry struct {
	ID                   string                                                                      `json:"id,required"`
	Amount               float64                                                                     `json:"amount,required"`
	CreatedAt            time.Time                                                                   `json:"created_at,required" format:"date-time"`
	CreditBlock          CustomerCreditLedgerListByExternalIDResponseIncrementLedgerEntryCreditBlock `json:"credit_block,required"`
	Currency             string                                                                      `json:"currency,required"`
	Customer             CustomerCreditLedgerListByExternalIDResponseIncrementLedgerEntryCustomer    `json:"customer,required"`
	Description          string                                                                      `json:"description,required,nullable"`
	EndingBalance        float64                                                                     `json:"ending_balance,required"`
	EntryStatus          CustomerCreditLedgerListByExternalIDResponseIncrementLedgerEntryEntryStatus `json:"entry_status,required"`
	EntryType            CustomerCreditLedgerListByExternalIDResponseIncrementLedgerEntryEntryType   `json:"entry_type,required"`
	LedgerSequenceNumber int64                                                                       `json:"ledger_sequence_number,required"`
	// User specified key-value pairs for the resource. If not present, this defaults
	// to an empty dictionary. Individual keys can be removed by setting the value to
	// `null`, and the entire metadata mapping can be cleared by setting `metadata` to
	// `null`.
	Metadata        map[string]string                                                    `json:"metadata,required"`
	StartingBalance float64                                                              `json:"starting_balance,required"`
	JSON            customerCreditLedgerListByExternalIDResponseIncrementLedgerEntryJSON `json:"-"`
}

// customerCreditLedgerListByExternalIDResponseIncrementLedgerEntryJSON contains
// the JSON metadata for the struct
// [CustomerCreditLedgerListByExternalIDResponseIncrementLedgerEntry]
type customerCreditLedgerListByExternalIDResponseIncrementLedgerEntryJSON struct {
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

func (r *CustomerCreditLedgerListByExternalIDResponseIncrementLedgerEntry) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r customerCreditLedgerListByExternalIDResponseIncrementLedgerEntryJSON) RawJSON() string {
	return r.raw
}

func (r CustomerCreditLedgerListByExternalIDResponseIncrementLedgerEntry) implementsCustomerCreditLedgerListByExternalIDResponse() {
}

type CustomerCreditLedgerListByExternalIDResponseIncrementLedgerEntryCreditBlock struct {
	ID               string                                                                          `json:"id,required"`
	ExpiryDate       time.Time                                                                       `json:"expiry_date,required,nullable" format:"date-time"`
	PerUnitCostBasis string                                                                          `json:"per_unit_cost_basis,required,nullable"`
	JSON             customerCreditLedgerListByExternalIDResponseIncrementLedgerEntryCreditBlockJSON `json:"-"`
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

func (r customerCreditLedgerListByExternalIDResponseIncrementLedgerEntryCreditBlockJSON) RawJSON() string {
	return r.raw
}

type CustomerCreditLedgerListByExternalIDResponseIncrementLedgerEntryCustomer struct {
	ID                 string                                                                       `json:"id,required"`
	ExternalCustomerID string                                                                       `json:"external_customer_id,required,nullable"`
	JSON               customerCreditLedgerListByExternalIDResponseIncrementLedgerEntryCustomerJSON `json:"-"`
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

func (r customerCreditLedgerListByExternalIDResponseIncrementLedgerEntryCustomerJSON) RawJSON() string {
	return r.raw
}

type CustomerCreditLedgerListByExternalIDResponseIncrementLedgerEntryEntryStatus string

const (
	CustomerCreditLedgerListByExternalIDResponseIncrementLedgerEntryEntryStatusCommitted CustomerCreditLedgerListByExternalIDResponseIncrementLedgerEntryEntryStatus = "committed"
	CustomerCreditLedgerListByExternalIDResponseIncrementLedgerEntryEntryStatusPending   CustomerCreditLedgerListByExternalIDResponseIncrementLedgerEntryEntryStatus = "pending"
)

func (r CustomerCreditLedgerListByExternalIDResponseIncrementLedgerEntryEntryStatus) IsKnown() bool {
	switch r {
	case CustomerCreditLedgerListByExternalIDResponseIncrementLedgerEntryEntryStatusCommitted, CustomerCreditLedgerListByExternalIDResponseIncrementLedgerEntryEntryStatusPending:
		return true
	}
	return false
}

type CustomerCreditLedgerListByExternalIDResponseIncrementLedgerEntryEntryType string

const (
	CustomerCreditLedgerListByExternalIDResponseIncrementLedgerEntryEntryTypeIncrement CustomerCreditLedgerListByExternalIDResponseIncrementLedgerEntryEntryType = "increment"
)

func (r CustomerCreditLedgerListByExternalIDResponseIncrementLedgerEntryEntryType) IsKnown() bool {
	switch r {
	case CustomerCreditLedgerListByExternalIDResponseIncrementLedgerEntryEntryTypeIncrement:
		return true
	}
	return false
}

type CustomerCreditLedgerListByExternalIDResponseDecrementLedgerEntry struct {
	ID                   string                                                                      `json:"id,required"`
	Amount               float64                                                                     `json:"amount,required"`
	CreatedAt            time.Time                                                                   `json:"created_at,required" format:"date-time"`
	CreditBlock          CustomerCreditLedgerListByExternalIDResponseDecrementLedgerEntryCreditBlock `json:"credit_block,required"`
	Currency             string                                                                      `json:"currency,required"`
	Customer             CustomerCreditLedgerListByExternalIDResponseDecrementLedgerEntryCustomer    `json:"customer,required"`
	Description          string                                                                      `json:"description,required,nullable"`
	EndingBalance        float64                                                                     `json:"ending_balance,required"`
	EntryStatus          CustomerCreditLedgerListByExternalIDResponseDecrementLedgerEntryEntryStatus `json:"entry_status,required"`
	EntryType            CustomerCreditLedgerListByExternalIDResponseDecrementLedgerEntryEntryType   `json:"entry_type,required"`
	LedgerSequenceNumber int64                                                                       `json:"ledger_sequence_number,required"`
	// User specified key-value pairs for the resource. If not present, this defaults
	// to an empty dictionary. Individual keys can be removed by setting the value to
	// `null`, and the entire metadata mapping can be cleared by setting `metadata` to
	// `null`.
	Metadata        map[string]string                                                    `json:"metadata,required"`
	StartingBalance float64                                                              `json:"starting_balance,required"`
	EventID         string                                                               `json:"event_id,nullable"`
	InvoiceID       string                                                               `json:"invoice_id,nullable"`
	PriceID         string                                                               `json:"price_id,nullable"`
	JSON            customerCreditLedgerListByExternalIDResponseDecrementLedgerEntryJSON `json:"-"`
}

// customerCreditLedgerListByExternalIDResponseDecrementLedgerEntryJSON contains
// the JSON metadata for the struct
// [CustomerCreditLedgerListByExternalIDResponseDecrementLedgerEntry]
type customerCreditLedgerListByExternalIDResponseDecrementLedgerEntryJSON struct {
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

func (r *CustomerCreditLedgerListByExternalIDResponseDecrementLedgerEntry) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r customerCreditLedgerListByExternalIDResponseDecrementLedgerEntryJSON) RawJSON() string {
	return r.raw
}

func (r CustomerCreditLedgerListByExternalIDResponseDecrementLedgerEntry) implementsCustomerCreditLedgerListByExternalIDResponse() {
}

type CustomerCreditLedgerListByExternalIDResponseDecrementLedgerEntryCreditBlock struct {
	ID               string                                                                          `json:"id,required"`
	ExpiryDate       time.Time                                                                       `json:"expiry_date,required,nullable" format:"date-time"`
	PerUnitCostBasis string                                                                          `json:"per_unit_cost_basis,required,nullable"`
	JSON             customerCreditLedgerListByExternalIDResponseDecrementLedgerEntryCreditBlockJSON `json:"-"`
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

func (r customerCreditLedgerListByExternalIDResponseDecrementLedgerEntryCreditBlockJSON) RawJSON() string {
	return r.raw
}

type CustomerCreditLedgerListByExternalIDResponseDecrementLedgerEntryCustomer struct {
	ID                 string                                                                       `json:"id,required"`
	ExternalCustomerID string                                                                       `json:"external_customer_id,required,nullable"`
	JSON               customerCreditLedgerListByExternalIDResponseDecrementLedgerEntryCustomerJSON `json:"-"`
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

func (r customerCreditLedgerListByExternalIDResponseDecrementLedgerEntryCustomerJSON) RawJSON() string {
	return r.raw
}

type CustomerCreditLedgerListByExternalIDResponseDecrementLedgerEntryEntryStatus string

const (
	CustomerCreditLedgerListByExternalIDResponseDecrementLedgerEntryEntryStatusCommitted CustomerCreditLedgerListByExternalIDResponseDecrementLedgerEntryEntryStatus = "committed"
	CustomerCreditLedgerListByExternalIDResponseDecrementLedgerEntryEntryStatusPending   CustomerCreditLedgerListByExternalIDResponseDecrementLedgerEntryEntryStatus = "pending"
)

func (r CustomerCreditLedgerListByExternalIDResponseDecrementLedgerEntryEntryStatus) IsKnown() bool {
	switch r {
	case CustomerCreditLedgerListByExternalIDResponseDecrementLedgerEntryEntryStatusCommitted, CustomerCreditLedgerListByExternalIDResponseDecrementLedgerEntryEntryStatusPending:
		return true
	}
	return false
}

type CustomerCreditLedgerListByExternalIDResponseDecrementLedgerEntryEntryType string

const (
	CustomerCreditLedgerListByExternalIDResponseDecrementLedgerEntryEntryTypeDecrement CustomerCreditLedgerListByExternalIDResponseDecrementLedgerEntryEntryType = "decrement"
)

func (r CustomerCreditLedgerListByExternalIDResponseDecrementLedgerEntryEntryType) IsKnown() bool {
	switch r {
	case CustomerCreditLedgerListByExternalIDResponseDecrementLedgerEntryEntryTypeDecrement:
		return true
	}
	return false
}

type CustomerCreditLedgerListByExternalIDResponseExpirationChangeLedgerEntry struct {
	ID                   string                                                                             `json:"id,required"`
	Amount               float64                                                                            `json:"amount,required"`
	CreatedAt            time.Time                                                                          `json:"created_at,required" format:"date-time"`
	CreditBlock          CustomerCreditLedgerListByExternalIDResponseExpirationChangeLedgerEntryCreditBlock `json:"credit_block,required"`
	Currency             string                                                                             `json:"currency,required"`
	Customer             CustomerCreditLedgerListByExternalIDResponseExpirationChangeLedgerEntryCustomer    `json:"customer,required"`
	Description          string                                                                             `json:"description,required,nullable"`
	EndingBalance        float64                                                                            `json:"ending_balance,required"`
	EntryStatus          CustomerCreditLedgerListByExternalIDResponseExpirationChangeLedgerEntryEntryStatus `json:"entry_status,required"`
	EntryType            CustomerCreditLedgerListByExternalIDResponseExpirationChangeLedgerEntryEntryType   `json:"entry_type,required"`
	LedgerSequenceNumber int64                                                                              `json:"ledger_sequence_number,required"`
	// User specified key-value pairs for the resource. If not present, this defaults
	// to an empty dictionary. Individual keys can be removed by setting the value to
	// `null`, and the entire metadata mapping can be cleared by setting `metadata` to
	// `null`.
	Metadata           map[string]string                                                           `json:"metadata,required"`
	NewBlockExpiryDate time.Time                                                                   `json:"new_block_expiry_date,required" format:"date-time"`
	StartingBalance    float64                                                                     `json:"starting_balance,required"`
	JSON               customerCreditLedgerListByExternalIDResponseExpirationChangeLedgerEntryJSON `json:"-"`
}

// customerCreditLedgerListByExternalIDResponseExpirationChangeLedgerEntryJSON
// contains the JSON metadata for the struct
// [CustomerCreditLedgerListByExternalIDResponseExpirationChangeLedgerEntry]
type customerCreditLedgerListByExternalIDResponseExpirationChangeLedgerEntryJSON struct {
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

func (r *CustomerCreditLedgerListByExternalIDResponseExpirationChangeLedgerEntry) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r customerCreditLedgerListByExternalIDResponseExpirationChangeLedgerEntryJSON) RawJSON() string {
	return r.raw
}

func (r CustomerCreditLedgerListByExternalIDResponseExpirationChangeLedgerEntry) implementsCustomerCreditLedgerListByExternalIDResponse() {
}

type CustomerCreditLedgerListByExternalIDResponseExpirationChangeLedgerEntryCreditBlock struct {
	ID               string                                                                                 `json:"id,required"`
	ExpiryDate       time.Time                                                                              `json:"expiry_date,required,nullable" format:"date-time"`
	PerUnitCostBasis string                                                                                 `json:"per_unit_cost_basis,required,nullable"`
	JSON             customerCreditLedgerListByExternalIDResponseExpirationChangeLedgerEntryCreditBlockJSON `json:"-"`
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

func (r customerCreditLedgerListByExternalIDResponseExpirationChangeLedgerEntryCreditBlockJSON) RawJSON() string {
	return r.raw
}

type CustomerCreditLedgerListByExternalIDResponseExpirationChangeLedgerEntryCustomer struct {
	ID                 string                                                                              `json:"id,required"`
	ExternalCustomerID string                                                                              `json:"external_customer_id,required,nullable"`
	JSON               customerCreditLedgerListByExternalIDResponseExpirationChangeLedgerEntryCustomerJSON `json:"-"`
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

func (r customerCreditLedgerListByExternalIDResponseExpirationChangeLedgerEntryCustomerJSON) RawJSON() string {
	return r.raw
}

type CustomerCreditLedgerListByExternalIDResponseExpirationChangeLedgerEntryEntryStatus string

const (
	CustomerCreditLedgerListByExternalIDResponseExpirationChangeLedgerEntryEntryStatusCommitted CustomerCreditLedgerListByExternalIDResponseExpirationChangeLedgerEntryEntryStatus = "committed"
	CustomerCreditLedgerListByExternalIDResponseExpirationChangeLedgerEntryEntryStatusPending   CustomerCreditLedgerListByExternalIDResponseExpirationChangeLedgerEntryEntryStatus = "pending"
)

func (r CustomerCreditLedgerListByExternalIDResponseExpirationChangeLedgerEntryEntryStatus) IsKnown() bool {
	switch r {
	case CustomerCreditLedgerListByExternalIDResponseExpirationChangeLedgerEntryEntryStatusCommitted, CustomerCreditLedgerListByExternalIDResponseExpirationChangeLedgerEntryEntryStatusPending:
		return true
	}
	return false
}

type CustomerCreditLedgerListByExternalIDResponseExpirationChangeLedgerEntryEntryType string

const (
	CustomerCreditLedgerListByExternalIDResponseExpirationChangeLedgerEntryEntryTypeExpirationChange CustomerCreditLedgerListByExternalIDResponseExpirationChangeLedgerEntryEntryType = "expiration_change"
)

func (r CustomerCreditLedgerListByExternalIDResponseExpirationChangeLedgerEntryEntryType) IsKnown() bool {
	switch r {
	case CustomerCreditLedgerListByExternalIDResponseExpirationChangeLedgerEntryEntryTypeExpirationChange:
		return true
	}
	return false
}

type CustomerCreditLedgerListByExternalIDResponseCreditBlockExpiryLedgerEntry struct {
	ID                   string                                                                              `json:"id,required"`
	Amount               float64                                                                             `json:"amount,required"`
	CreatedAt            time.Time                                                                           `json:"created_at,required" format:"date-time"`
	CreditBlock          CustomerCreditLedgerListByExternalIDResponseCreditBlockExpiryLedgerEntryCreditBlock `json:"credit_block,required"`
	Currency             string                                                                              `json:"currency,required"`
	Customer             CustomerCreditLedgerListByExternalIDResponseCreditBlockExpiryLedgerEntryCustomer    `json:"customer,required"`
	Description          string                                                                              `json:"description,required,nullable"`
	EndingBalance        float64                                                                             `json:"ending_balance,required"`
	EntryStatus          CustomerCreditLedgerListByExternalIDResponseCreditBlockExpiryLedgerEntryEntryStatus `json:"entry_status,required"`
	EntryType            CustomerCreditLedgerListByExternalIDResponseCreditBlockExpiryLedgerEntryEntryType   `json:"entry_type,required"`
	LedgerSequenceNumber int64                                                                               `json:"ledger_sequence_number,required"`
	// User specified key-value pairs for the resource. If not present, this defaults
	// to an empty dictionary. Individual keys can be removed by setting the value to
	// `null`, and the entire metadata mapping can be cleared by setting `metadata` to
	// `null`.
	Metadata        map[string]string                                                            `json:"metadata,required"`
	StartingBalance float64                                                                      `json:"starting_balance,required"`
	JSON            customerCreditLedgerListByExternalIDResponseCreditBlockExpiryLedgerEntryJSON `json:"-"`
}

// customerCreditLedgerListByExternalIDResponseCreditBlockExpiryLedgerEntryJSON
// contains the JSON metadata for the struct
// [CustomerCreditLedgerListByExternalIDResponseCreditBlockExpiryLedgerEntry]
type customerCreditLedgerListByExternalIDResponseCreditBlockExpiryLedgerEntryJSON struct {
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

func (r *CustomerCreditLedgerListByExternalIDResponseCreditBlockExpiryLedgerEntry) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r customerCreditLedgerListByExternalIDResponseCreditBlockExpiryLedgerEntryJSON) RawJSON() string {
	return r.raw
}

func (r CustomerCreditLedgerListByExternalIDResponseCreditBlockExpiryLedgerEntry) implementsCustomerCreditLedgerListByExternalIDResponse() {
}

type CustomerCreditLedgerListByExternalIDResponseCreditBlockExpiryLedgerEntryCreditBlock struct {
	ID               string                                                                                  `json:"id,required"`
	ExpiryDate       time.Time                                                                               `json:"expiry_date,required,nullable" format:"date-time"`
	PerUnitCostBasis string                                                                                  `json:"per_unit_cost_basis,required,nullable"`
	JSON             customerCreditLedgerListByExternalIDResponseCreditBlockExpiryLedgerEntryCreditBlockJSON `json:"-"`
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

func (r customerCreditLedgerListByExternalIDResponseCreditBlockExpiryLedgerEntryCreditBlockJSON) RawJSON() string {
	return r.raw
}

type CustomerCreditLedgerListByExternalIDResponseCreditBlockExpiryLedgerEntryCustomer struct {
	ID                 string                                                                               `json:"id,required"`
	ExternalCustomerID string                                                                               `json:"external_customer_id,required,nullable"`
	JSON               customerCreditLedgerListByExternalIDResponseCreditBlockExpiryLedgerEntryCustomerJSON `json:"-"`
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

func (r customerCreditLedgerListByExternalIDResponseCreditBlockExpiryLedgerEntryCustomerJSON) RawJSON() string {
	return r.raw
}

type CustomerCreditLedgerListByExternalIDResponseCreditBlockExpiryLedgerEntryEntryStatus string

const (
	CustomerCreditLedgerListByExternalIDResponseCreditBlockExpiryLedgerEntryEntryStatusCommitted CustomerCreditLedgerListByExternalIDResponseCreditBlockExpiryLedgerEntryEntryStatus = "committed"
	CustomerCreditLedgerListByExternalIDResponseCreditBlockExpiryLedgerEntryEntryStatusPending   CustomerCreditLedgerListByExternalIDResponseCreditBlockExpiryLedgerEntryEntryStatus = "pending"
)

func (r CustomerCreditLedgerListByExternalIDResponseCreditBlockExpiryLedgerEntryEntryStatus) IsKnown() bool {
	switch r {
	case CustomerCreditLedgerListByExternalIDResponseCreditBlockExpiryLedgerEntryEntryStatusCommitted, CustomerCreditLedgerListByExternalIDResponseCreditBlockExpiryLedgerEntryEntryStatusPending:
		return true
	}
	return false
}

type CustomerCreditLedgerListByExternalIDResponseCreditBlockExpiryLedgerEntryEntryType string

const (
	CustomerCreditLedgerListByExternalIDResponseCreditBlockExpiryLedgerEntryEntryTypeCreditBlockExpiry CustomerCreditLedgerListByExternalIDResponseCreditBlockExpiryLedgerEntryEntryType = "credit_block_expiry"
)

func (r CustomerCreditLedgerListByExternalIDResponseCreditBlockExpiryLedgerEntryEntryType) IsKnown() bool {
	switch r {
	case CustomerCreditLedgerListByExternalIDResponseCreditBlockExpiryLedgerEntryEntryTypeCreditBlockExpiry:
		return true
	}
	return false
}

type CustomerCreditLedgerListByExternalIDResponseVoidLedgerEntry struct {
	ID                   string                                                                 `json:"id,required"`
	Amount               float64                                                                `json:"amount,required"`
	CreatedAt            time.Time                                                              `json:"created_at,required" format:"date-time"`
	CreditBlock          CustomerCreditLedgerListByExternalIDResponseVoidLedgerEntryCreditBlock `json:"credit_block,required"`
	Currency             string                                                                 `json:"currency,required"`
	Customer             CustomerCreditLedgerListByExternalIDResponseVoidLedgerEntryCustomer    `json:"customer,required"`
	Description          string                                                                 `json:"description,required,nullable"`
	EndingBalance        float64                                                                `json:"ending_balance,required"`
	EntryStatus          CustomerCreditLedgerListByExternalIDResponseVoidLedgerEntryEntryStatus `json:"entry_status,required"`
	EntryType            CustomerCreditLedgerListByExternalIDResponseVoidLedgerEntryEntryType   `json:"entry_type,required"`
	LedgerSequenceNumber int64                                                                  `json:"ledger_sequence_number,required"`
	// User specified key-value pairs for the resource. If not present, this defaults
	// to an empty dictionary. Individual keys can be removed by setting the value to
	// `null`, and the entire metadata mapping can be cleared by setting `metadata` to
	// `null`.
	Metadata        map[string]string                                               `json:"metadata,required"`
	StartingBalance float64                                                         `json:"starting_balance,required"`
	VoidAmount      float64                                                         `json:"void_amount,required"`
	VoidReason      string                                                          `json:"void_reason,required,nullable"`
	JSON            customerCreditLedgerListByExternalIDResponseVoidLedgerEntryJSON `json:"-"`
}

// customerCreditLedgerListByExternalIDResponseVoidLedgerEntryJSON contains the
// JSON metadata for the struct
// [CustomerCreditLedgerListByExternalIDResponseVoidLedgerEntry]
type customerCreditLedgerListByExternalIDResponseVoidLedgerEntryJSON struct {
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

func (r *CustomerCreditLedgerListByExternalIDResponseVoidLedgerEntry) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r customerCreditLedgerListByExternalIDResponseVoidLedgerEntryJSON) RawJSON() string {
	return r.raw
}

func (r CustomerCreditLedgerListByExternalIDResponseVoidLedgerEntry) implementsCustomerCreditLedgerListByExternalIDResponse() {
}

type CustomerCreditLedgerListByExternalIDResponseVoidLedgerEntryCreditBlock struct {
	ID               string                                                                     `json:"id,required"`
	ExpiryDate       time.Time                                                                  `json:"expiry_date,required,nullable" format:"date-time"`
	PerUnitCostBasis string                                                                     `json:"per_unit_cost_basis,required,nullable"`
	JSON             customerCreditLedgerListByExternalIDResponseVoidLedgerEntryCreditBlockJSON `json:"-"`
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

func (r customerCreditLedgerListByExternalIDResponseVoidLedgerEntryCreditBlockJSON) RawJSON() string {
	return r.raw
}

type CustomerCreditLedgerListByExternalIDResponseVoidLedgerEntryCustomer struct {
	ID                 string                                                                  `json:"id,required"`
	ExternalCustomerID string                                                                  `json:"external_customer_id,required,nullable"`
	JSON               customerCreditLedgerListByExternalIDResponseVoidLedgerEntryCustomerJSON `json:"-"`
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

func (r customerCreditLedgerListByExternalIDResponseVoidLedgerEntryCustomerJSON) RawJSON() string {
	return r.raw
}

type CustomerCreditLedgerListByExternalIDResponseVoidLedgerEntryEntryStatus string

const (
	CustomerCreditLedgerListByExternalIDResponseVoidLedgerEntryEntryStatusCommitted CustomerCreditLedgerListByExternalIDResponseVoidLedgerEntryEntryStatus = "committed"
	CustomerCreditLedgerListByExternalIDResponseVoidLedgerEntryEntryStatusPending   CustomerCreditLedgerListByExternalIDResponseVoidLedgerEntryEntryStatus = "pending"
)

func (r CustomerCreditLedgerListByExternalIDResponseVoidLedgerEntryEntryStatus) IsKnown() bool {
	switch r {
	case CustomerCreditLedgerListByExternalIDResponseVoidLedgerEntryEntryStatusCommitted, CustomerCreditLedgerListByExternalIDResponseVoidLedgerEntryEntryStatusPending:
		return true
	}
	return false
}

type CustomerCreditLedgerListByExternalIDResponseVoidLedgerEntryEntryType string

const (
	CustomerCreditLedgerListByExternalIDResponseVoidLedgerEntryEntryTypeVoid CustomerCreditLedgerListByExternalIDResponseVoidLedgerEntryEntryType = "void"
)

func (r CustomerCreditLedgerListByExternalIDResponseVoidLedgerEntryEntryType) IsKnown() bool {
	switch r {
	case CustomerCreditLedgerListByExternalIDResponseVoidLedgerEntryEntryTypeVoid:
		return true
	}
	return false
}

type CustomerCreditLedgerListByExternalIDResponseVoidInitiatedLedgerEntry struct {
	ID                   string                                                                          `json:"id,required"`
	Amount               float64                                                                         `json:"amount,required"`
	CreatedAt            time.Time                                                                       `json:"created_at,required" format:"date-time"`
	CreditBlock          CustomerCreditLedgerListByExternalIDResponseVoidInitiatedLedgerEntryCreditBlock `json:"credit_block,required"`
	Currency             string                                                                          `json:"currency,required"`
	Customer             CustomerCreditLedgerListByExternalIDResponseVoidInitiatedLedgerEntryCustomer    `json:"customer,required"`
	Description          string                                                                          `json:"description,required,nullable"`
	EndingBalance        float64                                                                         `json:"ending_balance,required"`
	EntryStatus          CustomerCreditLedgerListByExternalIDResponseVoidInitiatedLedgerEntryEntryStatus `json:"entry_status,required"`
	EntryType            CustomerCreditLedgerListByExternalIDResponseVoidInitiatedLedgerEntryEntryType   `json:"entry_type,required"`
	LedgerSequenceNumber int64                                                                           `json:"ledger_sequence_number,required"`
	// User specified key-value pairs for the resource. If not present, this defaults
	// to an empty dictionary. Individual keys can be removed by setting the value to
	// `null`, and the entire metadata mapping can be cleared by setting `metadata` to
	// `null`.
	Metadata           map[string]string                                                        `json:"metadata,required"`
	NewBlockExpiryDate time.Time                                                                `json:"new_block_expiry_date,required" format:"date-time"`
	StartingBalance    float64                                                                  `json:"starting_balance,required"`
	VoidAmount         float64                                                                  `json:"void_amount,required"`
	VoidReason         string                                                                   `json:"void_reason,required,nullable"`
	JSON               customerCreditLedgerListByExternalIDResponseVoidInitiatedLedgerEntryJSON `json:"-"`
}

// customerCreditLedgerListByExternalIDResponseVoidInitiatedLedgerEntryJSON
// contains the JSON metadata for the struct
// [CustomerCreditLedgerListByExternalIDResponseVoidInitiatedLedgerEntry]
type customerCreditLedgerListByExternalIDResponseVoidInitiatedLedgerEntryJSON struct {
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

func (r *CustomerCreditLedgerListByExternalIDResponseVoidInitiatedLedgerEntry) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r customerCreditLedgerListByExternalIDResponseVoidInitiatedLedgerEntryJSON) RawJSON() string {
	return r.raw
}

func (r CustomerCreditLedgerListByExternalIDResponseVoidInitiatedLedgerEntry) implementsCustomerCreditLedgerListByExternalIDResponse() {
}

type CustomerCreditLedgerListByExternalIDResponseVoidInitiatedLedgerEntryCreditBlock struct {
	ID               string                                                                              `json:"id,required"`
	ExpiryDate       time.Time                                                                           `json:"expiry_date,required,nullable" format:"date-time"`
	PerUnitCostBasis string                                                                              `json:"per_unit_cost_basis,required,nullable"`
	JSON             customerCreditLedgerListByExternalIDResponseVoidInitiatedLedgerEntryCreditBlockJSON `json:"-"`
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

func (r customerCreditLedgerListByExternalIDResponseVoidInitiatedLedgerEntryCreditBlockJSON) RawJSON() string {
	return r.raw
}

type CustomerCreditLedgerListByExternalIDResponseVoidInitiatedLedgerEntryCustomer struct {
	ID                 string                                                                           `json:"id,required"`
	ExternalCustomerID string                                                                           `json:"external_customer_id,required,nullable"`
	JSON               customerCreditLedgerListByExternalIDResponseVoidInitiatedLedgerEntryCustomerJSON `json:"-"`
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

func (r customerCreditLedgerListByExternalIDResponseVoidInitiatedLedgerEntryCustomerJSON) RawJSON() string {
	return r.raw
}

type CustomerCreditLedgerListByExternalIDResponseVoidInitiatedLedgerEntryEntryStatus string

const (
	CustomerCreditLedgerListByExternalIDResponseVoidInitiatedLedgerEntryEntryStatusCommitted CustomerCreditLedgerListByExternalIDResponseVoidInitiatedLedgerEntryEntryStatus = "committed"
	CustomerCreditLedgerListByExternalIDResponseVoidInitiatedLedgerEntryEntryStatusPending   CustomerCreditLedgerListByExternalIDResponseVoidInitiatedLedgerEntryEntryStatus = "pending"
)

func (r CustomerCreditLedgerListByExternalIDResponseVoidInitiatedLedgerEntryEntryStatus) IsKnown() bool {
	switch r {
	case CustomerCreditLedgerListByExternalIDResponseVoidInitiatedLedgerEntryEntryStatusCommitted, CustomerCreditLedgerListByExternalIDResponseVoidInitiatedLedgerEntryEntryStatusPending:
		return true
	}
	return false
}

type CustomerCreditLedgerListByExternalIDResponseVoidInitiatedLedgerEntryEntryType string

const (
	CustomerCreditLedgerListByExternalIDResponseVoidInitiatedLedgerEntryEntryTypeVoidInitiated CustomerCreditLedgerListByExternalIDResponseVoidInitiatedLedgerEntryEntryType = "void_initiated"
)

func (r CustomerCreditLedgerListByExternalIDResponseVoidInitiatedLedgerEntryEntryType) IsKnown() bool {
	switch r {
	case CustomerCreditLedgerListByExternalIDResponseVoidInitiatedLedgerEntryEntryTypeVoidInitiated:
		return true
	}
	return false
}

type CustomerCreditLedgerListByExternalIDResponseAmendmentLedgerEntry struct {
	ID                   string                                                                      `json:"id,required"`
	Amount               float64                                                                     `json:"amount,required"`
	CreatedAt            time.Time                                                                   `json:"created_at,required" format:"date-time"`
	CreditBlock          CustomerCreditLedgerListByExternalIDResponseAmendmentLedgerEntryCreditBlock `json:"credit_block,required"`
	Currency             string                                                                      `json:"currency,required"`
	Customer             CustomerCreditLedgerListByExternalIDResponseAmendmentLedgerEntryCustomer    `json:"customer,required"`
	Description          string                                                                      `json:"description,required,nullable"`
	EndingBalance        float64                                                                     `json:"ending_balance,required"`
	EntryStatus          CustomerCreditLedgerListByExternalIDResponseAmendmentLedgerEntryEntryStatus `json:"entry_status,required"`
	EntryType            CustomerCreditLedgerListByExternalIDResponseAmendmentLedgerEntryEntryType   `json:"entry_type,required"`
	LedgerSequenceNumber int64                                                                       `json:"ledger_sequence_number,required"`
	// User specified key-value pairs for the resource. If not present, this defaults
	// to an empty dictionary. Individual keys can be removed by setting the value to
	// `null`, and the entire metadata mapping can be cleared by setting `metadata` to
	// `null`.
	Metadata        map[string]string                                                    `json:"metadata,required"`
	StartingBalance float64                                                              `json:"starting_balance,required"`
	JSON            customerCreditLedgerListByExternalIDResponseAmendmentLedgerEntryJSON `json:"-"`
}

// customerCreditLedgerListByExternalIDResponseAmendmentLedgerEntryJSON contains
// the JSON metadata for the struct
// [CustomerCreditLedgerListByExternalIDResponseAmendmentLedgerEntry]
type customerCreditLedgerListByExternalIDResponseAmendmentLedgerEntryJSON struct {
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

func (r *CustomerCreditLedgerListByExternalIDResponseAmendmentLedgerEntry) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r customerCreditLedgerListByExternalIDResponseAmendmentLedgerEntryJSON) RawJSON() string {
	return r.raw
}

func (r CustomerCreditLedgerListByExternalIDResponseAmendmentLedgerEntry) implementsCustomerCreditLedgerListByExternalIDResponse() {
}

type CustomerCreditLedgerListByExternalIDResponseAmendmentLedgerEntryCreditBlock struct {
	ID               string                                                                          `json:"id,required"`
	ExpiryDate       time.Time                                                                       `json:"expiry_date,required,nullable" format:"date-time"`
	PerUnitCostBasis string                                                                          `json:"per_unit_cost_basis,required,nullable"`
	JSON             customerCreditLedgerListByExternalIDResponseAmendmentLedgerEntryCreditBlockJSON `json:"-"`
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

func (r customerCreditLedgerListByExternalIDResponseAmendmentLedgerEntryCreditBlockJSON) RawJSON() string {
	return r.raw
}

type CustomerCreditLedgerListByExternalIDResponseAmendmentLedgerEntryCustomer struct {
	ID                 string                                                                       `json:"id,required"`
	ExternalCustomerID string                                                                       `json:"external_customer_id,required,nullable"`
	JSON               customerCreditLedgerListByExternalIDResponseAmendmentLedgerEntryCustomerJSON `json:"-"`
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

func (r customerCreditLedgerListByExternalIDResponseAmendmentLedgerEntryCustomerJSON) RawJSON() string {
	return r.raw
}

type CustomerCreditLedgerListByExternalIDResponseAmendmentLedgerEntryEntryStatus string

const (
	CustomerCreditLedgerListByExternalIDResponseAmendmentLedgerEntryEntryStatusCommitted CustomerCreditLedgerListByExternalIDResponseAmendmentLedgerEntryEntryStatus = "committed"
	CustomerCreditLedgerListByExternalIDResponseAmendmentLedgerEntryEntryStatusPending   CustomerCreditLedgerListByExternalIDResponseAmendmentLedgerEntryEntryStatus = "pending"
)

func (r CustomerCreditLedgerListByExternalIDResponseAmendmentLedgerEntryEntryStatus) IsKnown() bool {
	switch r {
	case CustomerCreditLedgerListByExternalIDResponseAmendmentLedgerEntryEntryStatusCommitted, CustomerCreditLedgerListByExternalIDResponseAmendmentLedgerEntryEntryStatusPending:
		return true
	}
	return false
}

type CustomerCreditLedgerListByExternalIDResponseAmendmentLedgerEntryEntryType string

const (
	CustomerCreditLedgerListByExternalIDResponseAmendmentLedgerEntryEntryTypeAmendment CustomerCreditLedgerListByExternalIDResponseAmendmentLedgerEntryEntryType = "amendment"
)

func (r CustomerCreditLedgerListByExternalIDResponseAmendmentLedgerEntryEntryType) IsKnown() bool {
	switch r {
	case CustomerCreditLedgerListByExternalIDResponseAmendmentLedgerEntryEntryTypeAmendment:
		return true
	}
	return false
}

type CustomerCreditLedgerListByExternalIDResponseEntryStatus string

const (
	CustomerCreditLedgerListByExternalIDResponseEntryStatusCommitted CustomerCreditLedgerListByExternalIDResponseEntryStatus = "committed"
	CustomerCreditLedgerListByExternalIDResponseEntryStatusPending   CustomerCreditLedgerListByExternalIDResponseEntryStatus = "pending"
)

func (r CustomerCreditLedgerListByExternalIDResponseEntryStatus) IsKnown() bool {
	switch r {
	case CustomerCreditLedgerListByExternalIDResponseEntryStatusCommitted, CustomerCreditLedgerListByExternalIDResponseEntryStatusPending:
		return true
	}
	return false
}

type CustomerCreditLedgerListByExternalIDResponseEntryType string

const (
	CustomerCreditLedgerListByExternalIDResponseEntryTypeIncrement         CustomerCreditLedgerListByExternalIDResponseEntryType = "increment"
	CustomerCreditLedgerListByExternalIDResponseEntryTypeDecrement         CustomerCreditLedgerListByExternalIDResponseEntryType = "decrement"
	CustomerCreditLedgerListByExternalIDResponseEntryTypeExpirationChange  CustomerCreditLedgerListByExternalIDResponseEntryType = "expiration_change"
	CustomerCreditLedgerListByExternalIDResponseEntryTypeCreditBlockExpiry CustomerCreditLedgerListByExternalIDResponseEntryType = "credit_block_expiry"
	CustomerCreditLedgerListByExternalIDResponseEntryTypeVoid              CustomerCreditLedgerListByExternalIDResponseEntryType = "void"
	CustomerCreditLedgerListByExternalIDResponseEntryTypeVoidInitiated     CustomerCreditLedgerListByExternalIDResponseEntryType = "void_initiated"
	CustomerCreditLedgerListByExternalIDResponseEntryTypeAmendment         CustomerCreditLedgerListByExternalIDResponseEntryType = "amendment"
)

func (r CustomerCreditLedgerListByExternalIDResponseEntryType) IsKnown() bool {
	switch r {
	case CustomerCreditLedgerListByExternalIDResponseEntryTypeIncrement, CustomerCreditLedgerListByExternalIDResponseEntryTypeDecrement, CustomerCreditLedgerListByExternalIDResponseEntryTypeExpirationChange, CustomerCreditLedgerListByExternalIDResponseEntryTypeCreditBlockExpiry, CustomerCreditLedgerListByExternalIDResponseEntryTypeVoid, CustomerCreditLedgerListByExternalIDResponseEntryTypeVoidInitiated, CustomerCreditLedgerListByExternalIDResponseEntryTypeAmendment:
		return true
	}
	return false
}

type CustomerCreditLedgerListParams struct {
	CreatedAtGt  param.Field[time.Time] `query:"created_at[gt]" format:"date-time"`
	CreatedAtGte param.Field[time.Time] `query:"created_at[gte]" format:"date-time"`
	CreatedAtLt  param.Field[time.Time] `query:"created_at[lt]" format:"date-time"`
	CreatedAtLte param.Field[time.Time] `query:"created_at[lte]" format:"date-time"`
	// The ledger currency or custom pricing unit to use.
	Currency param.Field[string] `query:"currency"`
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
		ArrayFormat:  apiquery.ArrayQueryFormatBrackets,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type CustomerCreditLedgerListParamsEntryStatus string

const (
	CustomerCreditLedgerListParamsEntryStatusCommitted CustomerCreditLedgerListParamsEntryStatus = "committed"
	CustomerCreditLedgerListParamsEntryStatusPending   CustomerCreditLedgerListParamsEntryStatus = "pending"
)

func (r CustomerCreditLedgerListParamsEntryStatus) IsKnown() bool {
	switch r {
	case CustomerCreditLedgerListParamsEntryStatusCommitted, CustomerCreditLedgerListParamsEntryStatusPending:
		return true
	}
	return false
}

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

func (r CustomerCreditLedgerListParamsEntryType) IsKnown() bool {
	switch r {
	case CustomerCreditLedgerListParamsEntryTypeIncrement, CustomerCreditLedgerListParamsEntryTypeDecrement, CustomerCreditLedgerListParamsEntryTypeExpirationChange, CustomerCreditLedgerListParamsEntryTypeCreditBlockExpiry, CustomerCreditLedgerListParamsEntryTypeVoid, CustomerCreditLedgerListParamsEntryTypeVoidInitiated, CustomerCreditLedgerListParamsEntryTypeAmendment:
		return true
	}
	return false
}

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
	InvoiceSettings param.Field[CustomerCreditLedgerNewEntryParamsAddIncrementCreditLedgerEntryRequestParamsInvoiceSettings] `json:"invoice_settings"`
	// User-specified key/value pairs for the resource. Individual keys can be removed
	// by setting the value to `null`, and the entire metadata mapping can be cleared
	// by setting `metadata` to `null`.
	Metadata param.Field[map[string]string] `json:"metadata"`
	// Can only be specified when entry_type=increment. How much, in the customer's
	// currency, a customer paid for a single credit in this block
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

func (r CustomerCreditLedgerNewEntryParamsAddIncrementCreditLedgerEntryRequestParamsEntryType) IsKnown() bool {
	switch r {
	case CustomerCreditLedgerNewEntryParamsAddIncrementCreditLedgerEntryRequestParamsEntryTypeIncrement:
		return true
	}
	return false
}

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
	// If true, the new credit block will require that the corresponding invoice is
	// paid before it can be drawn down from.
	RequireSuccessfulPayment param.Field[bool] `json:"require_successful_payment"`
}

func (r CustomerCreditLedgerNewEntryParamsAddIncrementCreditLedgerEntryRequestParamsInvoiceSettings) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type CustomerCreditLedgerNewEntryParamsAddDecrementCreditLedgerEntryRequestParams struct {
	// The number of credits to effect. Note that this is required for increment,
	// decrement, void, or undo operations.
	Amount    param.Field[float64]                                                                               `json:"amount,required"`
	EntryType param.Field[CustomerCreditLedgerNewEntryParamsAddDecrementCreditLedgerEntryRequestParamsEntryType] `json:"entry_type,required"`
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

func (r CustomerCreditLedgerNewEntryParamsAddDecrementCreditLedgerEntryRequestParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (CustomerCreditLedgerNewEntryParamsAddDecrementCreditLedgerEntryRequestParams) ImplementsCustomerCreditLedgerNewEntryParams() {

}

type CustomerCreditLedgerNewEntryParamsAddDecrementCreditLedgerEntryRequestParamsEntryType string

const (
	CustomerCreditLedgerNewEntryParamsAddDecrementCreditLedgerEntryRequestParamsEntryTypeDecrement CustomerCreditLedgerNewEntryParamsAddDecrementCreditLedgerEntryRequestParamsEntryType = "decrement"
)

func (r CustomerCreditLedgerNewEntryParamsAddDecrementCreditLedgerEntryRequestParamsEntryType) IsKnown() bool {
	switch r {
	case CustomerCreditLedgerNewEntryParamsAddDecrementCreditLedgerEntryRequestParamsEntryTypeDecrement:
		return true
	}
	return false
}

type CustomerCreditLedgerNewEntryParamsAddExpirationChangeCreditLedgerEntryRequestParams struct {
	EntryType param.Field[CustomerCreditLedgerNewEntryParamsAddExpirationChangeCreditLedgerEntryRequestParamsEntryType] `json:"entry_type,required"`
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

func (r CustomerCreditLedgerNewEntryParamsAddExpirationChangeCreditLedgerEntryRequestParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (CustomerCreditLedgerNewEntryParamsAddExpirationChangeCreditLedgerEntryRequestParams) ImplementsCustomerCreditLedgerNewEntryParams() {

}

type CustomerCreditLedgerNewEntryParamsAddExpirationChangeCreditLedgerEntryRequestParamsEntryType string

const (
	CustomerCreditLedgerNewEntryParamsAddExpirationChangeCreditLedgerEntryRequestParamsEntryTypeExpirationChange CustomerCreditLedgerNewEntryParamsAddExpirationChangeCreditLedgerEntryRequestParamsEntryType = "expiration_change"
)

func (r CustomerCreditLedgerNewEntryParamsAddExpirationChangeCreditLedgerEntryRequestParamsEntryType) IsKnown() bool {
	switch r {
	case CustomerCreditLedgerNewEntryParamsAddExpirationChangeCreditLedgerEntryRequestParamsEntryTypeExpirationChange:
		return true
	}
	return false
}

type CustomerCreditLedgerNewEntryParamsAddVoidCreditLedgerEntryRequestParams struct {
	// The number of credits to effect. Note that this is required for increment,
	// decrement, void, or undo operations.
	Amount param.Field[float64] `json:"amount,required"`
	// The ID of the block to void.
	BlockID   param.Field[string]                                                                           `json:"block_id,required"`
	EntryType param.Field[CustomerCreditLedgerNewEntryParamsAddVoidCreditLedgerEntryRequestParamsEntryType] `json:"entry_type,required"`
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

func (r CustomerCreditLedgerNewEntryParamsAddVoidCreditLedgerEntryRequestParamsEntryType) IsKnown() bool {
	switch r {
	case CustomerCreditLedgerNewEntryParamsAddVoidCreditLedgerEntryRequestParamsEntryTypeVoid:
		return true
	}
	return false
}

// Can only be specified when `entry_type=void`. The reason for the void.
type CustomerCreditLedgerNewEntryParamsAddVoidCreditLedgerEntryRequestParamsVoidReason string

const (
	CustomerCreditLedgerNewEntryParamsAddVoidCreditLedgerEntryRequestParamsVoidReasonRefund CustomerCreditLedgerNewEntryParamsAddVoidCreditLedgerEntryRequestParamsVoidReason = "refund"
)

func (r CustomerCreditLedgerNewEntryParamsAddVoidCreditLedgerEntryRequestParamsVoidReason) IsKnown() bool {
	switch r {
	case CustomerCreditLedgerNewEntryParamsAddVoidCreditLedgerEntryRequestParamsVoidReasonRefund:
		return true
	}
	return false
}

type CustomerCreditLedgerNewEntryParamsAddAmendmentCreditLedgerEntryRequestParams struct {
	// The number of credits to effect. Note that this is required for increment,
	// decrement or void operations.
	Amount param.Field[float64] `json:"amount,required"`
	// The ID of the block to reverse a decrement from.
	BlockID   param.Field[string]                                                                                `json:"block_id,required"`
	EntryType param.Field[CustomerCreditLedgerNewEntryParamsAddAmendmentCreditLedgerEntryRequestParamsEntryType] `json:"entry_type,required"`
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

func (r CustomerCreditLedgerNewEntryParamsAddAmendmentCreditLedgerEntryRequestParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (CustomerCreditLedgerNewEntryParamsAddAmendmentCreditLedgerEntryRequestParams) ImplementsCustomerCreditLedgerNewEntryParams() {

}

type CustomerCreditLedgerNewEntryParamsAddAmendmentCreditLedgerEntryRequestParamsEntryType string

const (
	CustomerCreditLedgerNewEntryParamsAddAmendmentCreditLedgerEntryRequestParamsEntryTypeAmendment CustomerCreditLedgerNewEntryParamsAddAmendmentCreditLedgerEntryRequestParamsEntryType = "amendment"
)

func (r CustomerCreditLedgerNewEntryParamsAddAmendmentCreditLedgerEntryRequestParamsEntryType) IsKnown() bool {
	switch r {
	case CustomerCreditLedgerNewEntryParamsAddAmendmentCreditLedgerEntryRequestParamsEntryTypeAmendment:
		return true
	}
	return false
}

// This interface is a union satisfied by one of the following:
// [CustomerCreditLedgerNewEntryByExternalIDParamsAddIncrementCreditLedgerEntryRequestParams],
// [CustomerCreditLedgerNewEntryByExternalIDParamsAddDecrementCreditLedgerEntryRequestParams],
// [CustomerCreditLedgerNewEntryByExternalIDParamsAddExpirationChangeCreditLedgerEntryRequestParams],
// [CustomerCreditLedgerNewEntryByExternalIDParamsAddVoidCreditLedgerEntryRequestParams],
// [CustomerCreditLedgerNewEntryByExternalIDParamsAddAmendmentCreditLedgerEntryRequestParams].
type CustomerCreditLedgerNewEntryByExternalIDParams interface {
	ImplementsCustomerCreditLedgerNewEntryByExternalIDParams()
}

type CustomerCreditLedgerNewEntryByExternalIDParamsAddIncrementCreditLedgerEntryRequestParams struct {
	// The number of credits to effect. Note that this is required for increment,
	// decrement, void, or undo operations.
	Amount    param.Field[float64]                                                                                           `json:"amount,required"`
	EntryType param.Field[CustomerCreditLedgerNewEntryByExternalIDParamsAddIncrementCreditLedgerEntryRequestParamsEntryType] `json:"entry_type,required"`
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
	InvoiceSettings param.Field[CustomerCreditLedgerNewEntryByExternalIDParamsAddIncrementCreditLedgerEntryRequestParamsInvoiceSettings] `json:"invoice_settings"`
	// User-specified key/value pairs for the resource. Individual keys can be removed
	// by setting the value to `null`, and the entire metadata mapping can be cleared
	// by setting `metadata` to `null`.
	Metadata param.Field[map[string]string] `json:"metadata"`
	// Can only be specified when entry_type=increment. How much, in the customer's
	// currency, a customer paid for a single credit in this block
	PerUnitCostBasis param.Field[string] `json:"per_unit_cost_basis"`
}

func (r CustomerCreditLedgerNewEntryByExternalIDParamsAddIncrementCreditLedgerEntryRequestParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (CustomerCreditLedgerNewEntryByExternalIDParamsAddIncrementCreditLedgerEntryRequestParams) ImplementsCustomerCreditLedgerNewEntryByExternalIDParams() {

}

type CustomerCreditLedgerNewEntryByExternalIDParamsAddIncrementCreditLedgerEntryRequestParamsEntryType string

const (
	CustomerCreditLedgerNewEntryByExternalIDParamsAddIncrementCreditLedgerEntryRequestParamsEntryTypeIncrement CustomerCreditLedgerNewEntryByExternalIDParamsAddIncrementCreditLedgerEntryRequestParamsEntryType = "increment"
)

func (r CustomerCreditLedgerNewEntryByExternalIDParamsAddIncrementCreditLedgerEntryRequestParamsEntryType) IsKnown() bool {
	switch r {
	case CustomerCreditLedgerNewEntryByExternalIDParamsAddIncrementCreditLedgerEntryRequestParamsEntryTypeIncrement:
		return true
	}
	return false
}

// Passing `invoice_settings` automatically generates an invoice for the newly
// added credits. If `invoice_settings` is passed, you must specify
// per_unit_cost_basis, as the calculation of the invoice total is done on that
// basis.
type CustomerCreditLedgerNewEntryByExternalIDParamsAddIncrementCreditLedgerEntryRequestParamsInvoiceSettings struct {
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

func (r CustomerCreditLedgerNewEntryByExternalIDParamsAddIncrementCreditLedgerEntryRequestParamsInvoiceSettings) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type CustomerCreditLedgerNewEntryByExternalIDParamsAddDecrementCreditLedgerEntryRequestParams struct {
	// The number of credits to effect. Note that this is required for increment,
	// decrement, void, or undo operations.
	Amount    param.Field[float64]                                                                                           `json:"amount,required"`
	EntryType param.Field[CustomerCreditLedgerNewEntryByExternalIDParamsAddDecrementCreditLedgerEntryRequestParamsEntryType] `json:"entry_type,required"`
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

func (r CustomerCreditLedgerNewEntryByExternalIDParamsAddDecrementCreditLedgerEntryRequestParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (CustomerCreditLedgerNewEntryByExternalIDParamsAddDecrementCreditLedgerEntryRequestParams) ImplementsCustomerCreditLedgerNewEntryByExternalIDParams() {

}

type CustomerCreditLedgerNewEntryByExternalIDParamsAddDecrementCreditLedgerEntryRequestParamsEntryType string

const (
	CustomerCreditLedgerNewEntryByExternalIDParamsAddDecrementCreditLedgerEntryRequestParamsEntryTypeDecrement CustomerCreditLedgerNewEntryByExternalIDParamsAddDecrementCreditLedgerEntryRequestParamsEntryType = "decrement"
)

func (r CustomerCreditLedgerNewEntryByExternalIDParamsAddDecrementCreditLedgerEntryRequestParamsEntryType) IsKnown() bool {
	switch r {
	case CustomerCreditLedgerNewEntryByExternalIDParamsAddDecrementCreditLedgerEntryRequestParamsEntryTypeDecrement:
		return true
	}
	return false
}

type CustomerCreditLedgerNewEntryByExternalIDParamsAddExpirationChangeCreditLedgerEntryRequestParams struct {
	EntryType param.Field[CustomerCreditLedgerNewEntryByExternalIDParamsAddExpirationChangeCreditLedgerEntryRequestParamsEntryType] `json:"entry_type,required"`
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

func (r CustomerCreditLedgerNewEntryByExternalIDParamsAddExpirationChangeCreditLedgerEntryRequestParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (CustomerCreditLedgerNewEntryByExternalIDParamsAddExpirationChangeCreditLedgerEntryRequestParams) ImplementsCustomerCreditLedgerNewEntryByExternalIDParams() {

}

type CustomerCreditLedgerNewEntryByExternalIDParamsAddExpirationChangeCreditLedgerEntryRequestParamsEntryType string

const (
	CustomerCreditLedgerNewEntryByExternalIDParamsAddExpirationChangeCreditLedgerEntryRequestParamsEntryTypeExpirationChange CustomerCreditLedgerNewEntryByExternalIDParamsAddExpirationChangeCreditLedgerEntryRequestParamsEntryType = "expiration_change"
)

func (r CustomerCreditLedgerNewEntryByExternalIDParamsAddExpirationChangeCreditLedgerEntryRequestParamsEntryType) IsKnown() bool {
	switch r {
	case CustomerCreditLedgerNewEntryByExternalIDParamsAddExpirationChangeCreditLedgerEntryRequestParamsEntryTypeExpirationChange:
		return true
	}
	return false
}

type CustomerCreditLedgerNewEntryByExternalIDParamsAddVoidCreditLedgerEntryRequestParams struct {
	// The number of credits to effect. Note that this is required for increment,
	// decrement, void, or undo operations.
	Amount param.Field[float64] `json:"amount,required"`
	// The ID of the block to void.
	BlockID   param.Field[string]                                                                                       `json:"block_id,required"`
	EntryType param.Field[CustomerCreditLedgerNewEntryByExternalIDParamsAddVoidCreditLedgerEntryRequestParamsEntryType] `json:"entry_type,required"`
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
	VoidReason param.Field[CustomerCreditLedgerNewEntryByExternalIDParamsAddVoidCreditLedgerEntryRequestParamsVoidReason] `json:"void_reason"`
}

func (r CustomerCreditLedgerNewEntryByExternalIDParamsAddVoidCreditLedgerEntryRequestParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (CustomerCreditLedgerNewEntryByExternalIDParamsAddVoidCreditLedgerEntryRequestParams) ImplementsCustomerCreditLedgerNewEntryByExternalIDParams() {

}

type CustomerCreditLedgerNewEntryByExternalIDParamsAddVoidCreditLedgerEntryRequestParamsEntryType string

const (
	CustomerCreditLedgerNewEntryByExternalIDParamsAddVoidCreditLedgerEntryRequestParamsEntryTypeVoid CustomerCreditLedgerNewEntryByExternalIDParamsAddVoidCreditLedgerEntryRequestParamsEntryType = "void"
)

func (r CustomerCreditLedgerNewEntryByExternalIDParamsAddVoidCreditLedgerEntryRequestParamsEntryType) IsKnown() bool {
	switch r {
	case CustomerCreditLedgerNewEntryByExternalIDParamsAddVoidCreditLedgerEntryRequestParamsEntryTypeVoid:
		return true
	}
	return false
}

// Can only be specified when `entry_type=void`. The reason for the void.
type CustomerCreditLedgerNewEntryByExternalIDParamsAddVoidCreditLedgerEntryRequestParamsVoidReason string

const (
	CustomerCreditLedgerNewEntryByExternalIDParamsAddVoidCreditLedgerEntryRequestParamsVoidReasonRefund CustomerCreditLedgerNewEntryByExternalIDParamsAddVoidCreditLedgerEntryRequestParamsVoidReason = "refund"
)

func (r CustomerCreditLedgerNewEntryByExternalIDParamsAddVoidCreditLedgerEntryRequestParamsVoidReason) IsKnown() bool {
	switch r {
	case CustomerCreditLedgerNewEntryByExternalIDParamsAddVoidCreditLedgerEntryRequestParamsVoidReasonRefund:
		return true
	}
	return false
}

type CustomerCreditLedgerNewEntryByExternalIDParamsAddAmendmentCreditLedgerEntryRequestParams struct {
	// The number of credits to effect. Note that this is required for increment,
	// decrement or void operations.
	Amount param.Field[float64] `json:"amount,required"`
	// The ID of the block to reverse a decrement from.
	BlockID   param.Field[string]                                                                                            `json:"block_id,required"`
	EntryType param.Field[CustomerCreditLedgerNewEntryByExternalIDParamsAddAmendmentCreditLedgerEntryRequestParamsEntryType] `json:"entry_type,required"`
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

func (r CustomerCreditLedgerNewEntryByExternalIDParamsAddAmendmentCreditLedgerEntryRequestParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (CustomerCreditLedgerNewEntryByExternalIDParamsAddAmendmentCreditLedgerEntryRequestParams) ImplementsCustomerCreditLedgerNewEntryByExternalIDParams() {

}

type CustomerCreditLedgerNewEntryByExternalIDParamsAddAmendmentCreditLedgerEntryRequestParamsEntryType string

const (
	CustomerCreditLedgerNewEntryByExternalIDParamsAddAmendmentCreditLedgerEntryRequestParamsEntryTypeAmendment CustomerCreditLedgerNewEntryByExternalIDParamsAddAmendmentCreditLedgerEntryRequestParamsEntryType = "amendment"
)

func (r CustomerCreditLedgerNewEntryByExternalIDParamsAddAmendmentCreditLedgerEntryRequestParamsEntryType) IsKnown() bool {
	switch r {
	case CustomerCreditLedgerNewEntryByExternalIDParamsAddAmendmentCreditLedgerEntryRequestParamsEntryTypeAmendment:
		return true
	}
	return false
}

type CustomerCreditLedgerListByExternalIDParams struct {
	CreatedAtGt  param.Field[time.Time] `query:"created_at[gt]" format:"date-time"`
	CreatedAtGte param.Field[time.Time] `query:"created_at[gte]" format:"date-time"`
	CreatedAtLt  param.Field[time.Time] `query:"created_at[lt]" format:"date-time"`
	CreatedAtLte param.Field[time.Time] `query:"created_at[lte]" format:"date-time"`
	// The ledger currency or custom pricing unit to use.
	Currency param.Field[string] `query:"currency"`
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
		ArrayFormat:  apiquery.ArrayQueryFormatBrackets,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type CustomerCreditLedgerListByExternalIDParamsEntryStatus string

const (
	CustomerCreditLedgerListByExternalIDParamsEntryStatusCommitted CustomerCreditLedgerListByExternalIDParamsEntryStatus = "committed"
	CustomerCreditLedgerListByExternalIDParamsEntryStatusPending   CustomerCreditLedgerListByExternalIDParamsEntryStatus = "pending"
)

func (r CustomerCreditLedgerListByExternalIDParamsEntryStatus) IsKnown() bool {
	switch r {
	case CustomerCreditLedgerListByExternalIDParamsEntryStatusCommitted, CustomerCreditLedgerListByExternalIDParamsEntryStatusPending:
		return true
	}
	return false
}

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

func (r CustomerCreditLedgerListByExternalIDParamsEntryType) IsKnown() bool {
	switch r {
	case CustomerCreditLedgerListByExternalIDParamsEntryTypeIncrement, CustomerCreditLedgerListByExternalIDParamsEntryTypeDecrement, CustomerCreditLedgerListByExternalIDParamsEntryTypeExpirationChange, CustomerCreditLedgerListByExternalIDParamsEntryTypeCreditBlockExpiry, CustomerCreditLedgerListByExternalIDParamsEntryTypeVoid, CustomerCreditLedgerListByExternalIDParamsEntryTypeVoidInitiated, CustomerCreditLedgerListByExternalIDParamsEntryTypeAmendment:
		return true
	}
	return false
}
