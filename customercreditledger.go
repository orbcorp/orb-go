// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package orb

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"net/url"
	"time"

	"github.com/orbcorp/orb-go/internal/apijson"
	"github.com/orbcorp/orb-go/internal/apiquery"
	"github.com/orbcorp/orb-go/internal/param"
	"github.com/orbcorp/orb-go/internal/requestconfig"
	"github.com/orbcorp/orb-go/option"
	"github.com/orbcorp/orb-go/packages/pagination"
	"github.com/orbcorp/orb-go/shared"
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
// balance. This [paginated endpoint](/api-reference/pagination) lists these
// entries, starting from the most recent ledger entry.
//
// More details on using Orb's real-time credit feature are
// [here](/product-catalog/prepurchase).
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
func (r *CustomerCreditLedgerService) List(ctx context.Context, customerID string, query CustomerCreditLedgerListParams, opts ...option.RequestOption) (res *pagination.Page[shared.CreditLedgerEntryModel], err error) {
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
// balance. This [paginated endpoint](/api-reference/pagination) lists these
// entries, starting from the most recent ledger entry.
//
// More details on using Orb's real-time credit feature are
// [here](/product-catalog/prepurchase).
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
func (r *CustomerCreditLedgerService) ListAutoPaging(ctx context.Context, customerID string, query CustomerCreditLedgerListParams, opts ...option.RequestOption) *pagination.PageAutoPager[shared.CreditLedgerEntryModel] {
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
//     therefore returned in the
//     [View Credits Ledger](fetch-customer-credits-ledger) response as well as
//     serialized in the response to this request. In the case of deductions without
//     a specified block, multiple ledger entries may be created if the deduction
//     spans credit blocks.
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
func (r *CustomerCreditLedgerService) NewEntry(ctx context.Context, customerID string, body CustomerCreditLedgerNewEntryParams, opts ...option.RequestOption) (res *shared.CreditLedgerEntryModel, err error) {
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
//     therefore returned in the
//     [View Credits Ledger](fetch-customer-credits-ledger) response as well as
//     serialized in the response to this request. In the case of deductions without
//     a specified block, multiple ledger entries may be created if the deduction
//     spans credit blocks.
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
func (r *CustomerCreditLedgerService) NewEntryByExternalID(ctx context.Context, externalCustomerID string, body CustomerCreditLedgerNewEntryByExternalIDParams, opts ...option.RequestOption) (res *shared.CreditLedgerEntryModel, err error) {
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
// balance. This [paginated endpoint](/api-reference/pagination) lists these
// entries, starting from the most recent ledger entry.
//
// More details on using Orb's real-time credit feature are
// [here](/product-catalog/prepurchase).
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
func (r *CustomerCreditLedgerService) ListByExternalID(ctx context.Context, externalCustomerID string, query CustomerCreditLedgerListByExternalIDParams, opts ...option.RequestOption) (res *pagination.Page[shared.CreditLedgerEntryModel], err error) {
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
// balance. This [paginated endpoint](/api-reference/pagination) lists these
// entries, starting from the most recent ledger entry.
//
// More details on using Orb's real-time credit feature are
// [here](/product-catalog/prepurchase).
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
func (r *CustomerCreditLedgerService) ListByExternalIDAutoPaging(ctx context.Context, externalCustomerID string, query CustomerCreditLedgerListByExternalIDParams, opts ...option.RequestOption) *pagination.PageAutoPager[shared.CreditLedgerEntryModel] {
	return pagination.NewPageAutoPager(r.ListByExternalID(ctx, externalCustomerID, query, opts...))
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
